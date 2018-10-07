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
import {Promise} from "mongoose";
import * as ClassTransformer from "class-transformer";
import * as mongoose from "mongoose";
import IUser from "../interfaces/internal/user";
import IExperienceAPI from "../interfaces/api/experience";
import IAddressAPI from "../interfaces/api/address";

export class ExperiencesUtil {

    /*
     * REST API handlers
     */

    public static getPersonalExperiences(req, res) {
        req.params.id = req.account.username;
        this.getExperiences(req, res);
    }

    public static getExperiences(req, res) {
        AccountModel.findOne({username: req.params.id, type: "User"}, function (err, user: IUser) {
            if (err) {
                if (servers.DEBUG) console.error(err);
                return res.status(500).send({message: errors.internalServerError});
            }
            let object: Array<IExperienceAPI> = [];
            user.experiences.forEach((element) => {
                object.push(new IExperienceAPI(new IAddressAPI(element.location), element.name, element.organization, element.opportunity, element.description, element.when, element.is_verified, element.created_at));
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
            is_verified: false,
            created_at: (new Date).getTime()
        });
    }

    public static updateExperience(req, res) {
        this.deleteExperience(req, res);
        this.createExperience(req, res, req.params.id);
    }

    public static createExperience(req, res, id) {
        if (req.accountType != "User") return res.status(400).send({message: errors.badRequest + " (Incorrect account type! User account type required.)"});

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
                    return res.status(500).send({message: errors.internalServerError});
                }
                if (!org) {
                    failed = true;
                    return res.status(400).send({message: errors.badRequest + " (Organization not found.)"})
                }

                org.experience_validations.push(new IValidations(req.account.username, exp._id)); // add validation entry to organization

                org.save(function (err) {
                    if (err) {
                        if (servers.DEBUG) console.error(err);
                        return res.status(500).send({message: errors.internalServerError});
                    }
                    if (!failed) ExperiencesUtil.saveExperienceMongo(req, res);
                }); // save to db
            });// TODO CASE INSENSITIVE LOOKUPS
        } else {
            // finish adding experience to database
            ExperiencesUtil.saveExperienceMongo(req, res);
        }
    }

    /*
     * Save the account
     */

    private static saveExperienceMongo(req, res) {
        req.account.save(function (err) {
            if (err) {
                if (servers.DEBUG) console.error(err);
                return res.status(500).send({message: errors.internalServerError});
            }
            if (!res.headersSent) return res.status(200).send({message: errors.ok}); // send ok header if all is good
        }); // save to db
    }

    public static deleteExperience(req, res) {
        if (req.accountType != "User") return res.status(400).send({message: errors.badRequest + " (Incorrect account type! User account type required.)"});

        let experience = JSON.parse(JSON.stringify(req.account.experiences[req.params.id])), index = req.account.experiences.indexOf(req.params.id);

        if (index < 0) return res.status(404).send({message: errors.notFound + " (Experience not found with suppplied ID)"});
        req.account.experiences.splice(index, 1); // remove experience from array (save later)

        if (experience.opportunity != undefined && experience.opportunity != "") {
            //TODO OPPORTUNITY
        }
        if (experience.organization != undefined && experience.organization != "") { // remove pending requests for experience
            AccountModel.findOne({username: experience.organization, type: "Organization"}, function(err, org: IOrganization) {
                if (err) {
                    if (servers.DEBUG) console.error(err);
                    return res.status(500).send({message: errors.internalServerError});
                }
                if (!org) {
                    return res.status(400).send({message: errors.badRequest + " (Organization not found.)"})
                }

                let index = org.experience_validations.indexOf(new IValidations(req.account.username, experience._id)); // get index of experience validation request

                if (index > -1) { // if it exists
                    org.experience_validations.splice(index, 1); // remove from array
                    org.save(function (err) {
                        if (err) {
                            if (servers.DEBUG) console.error(err);
                            return res.status(500).send({message: errors.internalServerError});
                        }
                        ExperiencesUtil.saveExperienceMongo(req, res);
                    }); // save to db
                } else { // if it doesn't exist
                    ExperiencesUtil.saveExperienceMongo(req, res);
                }
            });
        } else {
            ExperiencesUtil.saveExperienceMongo(req, res);
        }
    }

    public static getExperienceValidations(req, res) {

    }

    public static reviewExperienceValidations(req: Request, res) {

    }
}