/*
 *
 *     Copyright (C) 2018 ConnectUS
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

import {AuthUtil} from "../auth/auth-util";
import * as errors from "../routes/errors";
import {AccountModel} from "../interfaces/internal/account";
import * as servers from "../server";
import IOrganization from "../interfaces/internal/organization";
import IValidations from "../interfaces/internal/organization";
import {ExperienceModel} from "../interfaces/internal/experience";
import * as mongoose from "mongoose";
import IUser from "../interfaces/internal/user";
import IExperienceAPI from "../interfaces/api/experience";
import IAddressAPI from "../interfaces/api/address";
import IValidationsAPI from "../interfaces/api/organization";
import {sendError} from "../routes/errors";

export class ExperiencesUtil {

    /*
     * REST API handlers
     */

    // Get experiences of personal user (username found from token)
    public static getPersonalExperiences(req, res) {
        if (req.accountType != "User") return sendError(res, 400, errors.badRequest + " (Incorrect account type! User account type required.)", 4000);
        req.params.id = req.account.username;
        this.getExperiences(req, res);
    }

    // Get all experiences of any user
    public static getExperiences(req, res) {
        AccountModel.findOne({username: req.params.id, type: "User"}, function (err, user: IUser) { // TODO CASE INSENSITIVE QUERY
            if (err) {
                if (servers.DEBUG) console.error(err);
                return sendError(res, 500, errors.internalServerError, 4001);
            }
            if (!user) {
                return sendError(res, 404, errors.notFound + " (User not found? Is this the correct account type?)", 4002);
            }
            let object: Array<IExperienceAPI> = [];
            user.experiences.forEach((element) => {
                object.push(new IExperienceAPI(new IAddressAPI(element.location), element._id, element.name, element.organization, element.opportunity, element.description, element.when, element.is_verified, element.created_at, element.hours));
            });
            res.status(200).send(object);
        });
    }

    private static getDefaultExperience(experience, id) {
        return new ExperienceModel({
            _id: id,
            schema_version: 0,
            location: experience.location,
            name: experience.name,
            organization: experience.organization,
            opportunity: experience.opportunity,
            description: experience.description,
            when: {
              begin: experience.begin,
              end: experience.end
            },
            hours: experience.hours,
            is_verified: false,
            created_at: (new Date).getTime()
        });
    }

    public static updateExperience(req, res) { //TODO REDO THE CODE (MULTIPLE SAVES IN PARALLEL), SAVE AND TRANSFER APPROVAL FROM ORGANIZATION IF ORGANIZATION REMAINS THE SAME
        this.deleteExperience(req, res, false, () => {
            if (!res.headersSent) this.createExperience(req, res, req.params.id, true);
        });
    }

    public static createExperience(req, res, id, save: boolean) {
        if (req.accountType != "User") return sendError(res, 400, errors.badRequest + " (Incorrect account type! User account type required.)", 4000);

        let failed = false;

        // default experience object

        let exp = this.getDefaultExperience(req.body.experience, id);

        req.account.experiences.push(exp); // add to user's experiences array

        // verifications for data

        if (req.body.experience.opportunity != undefined && req.body.experience.opportunity != "") {
            // TODO OPPORTUNITY
        }

        if (req.body.experience.organization != undefined && req.body.experience.organization != "") { // check if there is an associated organization on the site (for validations)
            // add experience to organization pending validations list
            // TODO NOTIFICATION ON PENDING VALIDATION
            // TODO DUPLICATE HEADERS SENT WHEN SAVING FAILURE
            AccountModel.findOne({username: req.body.experience.organization, type: "Organization"}, function (err, org: IOrganization) {
                if (err) {
                    failed = true;
                    if (servers.DEBUG) console.error(err);
                    return sendError(res, 500, errors.internalServerError, 4001);
                }
                if (!org) {
                    failed = true;
                    return sendError(res, 404, errors.badRequest + " (Organization not found.)", 4002);
                }

                org.experience_validations.push(new IValidations(req.account.username, exp._id)); // add validation entry to organization

                org.save(function (err) {
                    if (err) {
                        if (servers.DEBUG) console.error(err);
                        return sendError(res, 500, errors.internalServerError, 4001);
                    }
                    if (!failed && save) ExperiencesUtil.saveAccountMongo(req, res);
                }); // save to db
            });// TODO CASE INSENSITIVE LOOKUPS
        } else {
            // finish adding experience to database
            if (save) ExperiencesUtil.saveAccountMongo(req, res);
        }
    }

    /*
     * Save the account
     */

    private static saveAccountMongo(req, res, call?) {
        if (call) call();
        req.account.save(function (err) {
            if (err) {
                if (servers.DEBUG) console.error(err);
                return sendError(res, 500, errors.internalServerError, 4001);
            }
            if (!res.headersSent) return res.status(200).send({message: errors.ok}); // send ok header if all is good
        }); // save to db
    }

    // TODO BUG IF ORGANIZATION IS REMOVED
    public static deleteExperience(req, res, save: boolean, callback?) {
        if (req.accountType != "User") return sendError(res, 400, errors.badRequest + " (Incorrect account type! User account type required.)", 4000);

        let experience, index = -1;

        for (let i = 0; i < req.account.experiences.length; i++) { // find experience to delete in array
            if (req.account.experiences[i]._id == req.params.id) {
                experience = req.account.experiences[i];
                index = i;
                break;
            }
        }

        if (index < 0) return sendError(res, 404, errors.notFound + " (Experience not found with supplied ID)", 4002);
        let func = () => req.account.experiences.splice(index, 1); // remove experience from array (save later)

        if (experience.opportunity != undefined && experience.opportunity != "") {
            //TODO OPPORTUNITY
        }
        if (experience.organization != undefined && experience.organization != "") { // remove pending requests for experience
            AccountModel.findOne({username: experience.organization, type: "Organization"}, function (err, org: IOrganization) {
                if (err) {
                    if (servers.DEBUG) console.error(err);
                    return sendError(res, 500, errors.internalServerError, 4001);
                }
                if (!org) {
                    if (save) ExperiencesUtil.saveAccountMongo(req, res, func);
                    if (callback) callback();
                }

                let index = -1; // get index of experience validation request

                for (let i = 0; i < org.experience_validations.length; i++) { // remove all entries with the same id and user (duplicates as well)
                    if (req.account.username == org.experience_validations[i].user_id && experience._id == org.experience_validations[i].experience_id) {
                        index = i;
                        org.experience_validations.splice(i, 1); // remove from array
                        i--;
                    }
                }

                if (index > -1) { // if it exists
                    org.save(function (err) { // save to db
                        if (err) {
                            if (servers.DEBUG) console.error(err);
                            return sendError(res, 500, errors.internalServerError, 4001);
                        }
                        if (save) ExperiencesUtil.saveAccountMongo(req, res, func);
                        if (callback) callback();
                    }); // save to db
                } else { // if it doesn't exist
                    if (save) ExperiencesUtil.saveAccountMongo(req, res, func);
                    if (callback) callback();
                }
            });
        } else {
            if (save) ExperiencesUtil.saveAccountMongo(req, res, func);
            if (callback) callback();
        }
    }

    public static getExperienceValidations(req, res) {
        if (req.accountType != "Organization") return sendError(res, 400, errors.badRequest + " (Incorrect account type! Organization account type required.)", 4000);

        let object: Array<IValidationsAPI> = [];
        req.account.experience_validations.forEach((element) => {
            object.push(new IValidationsAPI(element.user_id, element.experience_id));
        });
        res.status(200).send(object);
    }

    // Approve or don't approve an experience validation request
    public static async reviewExperienceValidations(req, res) {
        if (req.accountType != "Organization") return sendError(res, 400, errors.badRequest + " (Incorrect account type! Organization account type required.)", 4000);
        if (!req.body.approved) return res.status(400).send({message: errors.badRequest + " (Bad query; no \"approved\" field)"}); //TODO REMOVE
        let found = false, accepted = req.body.approved;

        // Remove the experience validation request from the organization object
        for (let i = 0; i < req.account.experience_validations.length; i++) {
            if (req.account.experience_validations[i].user_id == req.params.user && req.account.experience_validations[i].experience_id == req.params.id) {
                req.account.experience_validations.splice(i, 1);
                i--;
                found = true;
            }
        }
        if (!found) return sendError(res, 404, errors.notFound + " (Experience validation request not found)", 4002);

        try {
            await req.account.save(); // save the organization object

            // Update user experience object with approval
            // @ts-ignore
            let user: IUser = await AccountModel.findOne({username: req.params.user, type: "User"}).exec();
            if (!user) {
                return sendError(res, 400, errors.badRequest + " (User not found.)", 4003);
            }
            found = false;
            for (let i = 0; i < user.experiences.length; i++) {
                if (user.experiences[i]._id == req.params.id) {
                    console.log(accepted); //TODO
                    if (accepted) user.experiences[i].is_verified = accepted; // verify experience object if approved
                    else {
                        user.experiences.splice(i, 1); // delete experience object if not approved
                        console.log("delete " + i); //TODO bug with removal (not working)
                        i--;
                    }
                    found = true;
                }
            }
            if (!found) return sendError(res, 400, errors.internalServerError + " (Experience not found in user object)", 4004);
            await user.save(); // save user object
            res.status(200).send({message: errors.ok});
        } catch (err) {
            if (servers.DEBUG) console.error(err);
            return sendError(res, 500, errors.internalServerError, 4001);
        }
    }
}