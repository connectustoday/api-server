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

export class ExperiencesUtil {

    /*
     * REST API handlers
     */

    public static getExperiences(req, res): any {
        let accType = req.accountType;
        if (accType != "User") return res.status(400).send({message: errors.badRequest + " (Incorrect account type! User account type required.)"});

    }

    public static createExperience(req, res) {
        let accType = req.accountType;
        if (accType != "User") return res.status(400).send({message: errors.badRequest + " (Incorrect account type! User account type required.)"});

        //if (!(req.experience instanceof IExperienceAPI)) return res.status(400).send({message: errors.badRequest + " (Malformed Experience object)"});
        //if (req.account.experiences == undefined) req.account.experiences = [];

        let failed = false, addOrganizationRequest = false;

        // verifications for data

        if (req.body.experience.opportunity != undefined && req.body.experience.opportunity != "") {
            // TODO OPPORTUNITY
        }

        if (req.body.experience.organization != undefined && req.body.experience.organization != "") { // check if there is an associated organization on the site (for validations)
            AccountModel.count({username: req.body.experience.organization, type: "Organization"}, function (err, count) {
                if (err) {
                    failed = true;
                    if (servers.DEBUG) console.error(err);
                    return res.status(500).send({message: errors.internalServerError});
                }
                if (count <= 0) {
                    failed = true;
                    return res.status(400).send({message: errors.badRequest + " (Organization not found.)"})
                }

                addOrganizationRequest = true;
                ExperiencesUtil.createExperienceMongo(req, res, failed, addOrganizationRequest);
            }); // TODO CASE INSENSITIVE LOOKUPS
        } else {
            // finish adding experience to database
            ExperiencesUtil.createExperienceMongo(req, res, failed, addOrganizationRequest);
        }
    }

    private static createExperienceMongo(req, res, failed: boolean, addOrganizationRequest: boolean) {
        let exp;

        if (failed) return;

        // default experience object
        exp = new ExperienceModel({
            schema_version: 0,
            location: req.body.experience.location,
            name: req.body.experience.name,
            organization: req.body.experience.organization,
            opportunity: req.body.experience.opportunity,
            description: req.body.experience.description,
            is_verified: false,
            created_at: (new Date).getTime()
        });

        req.account.experiences.push(exp); // add to user's experience
        req.account.save(function (err) {
            if (err) {
                if (servers.DEBUG) console.error(err);
                return res.status(500).send({message: errors.internalServerError});
            }
        }); // save to db
        console.log(addOrganizationRequest);
        if (addOrganizationRequest) {
            // add experience to organization pending validations list
            // TODO NOTIFICATION ON PENDING VALIDATION
            AccountModel.findOne({username: req.body.experience.organization, type: "Organization"}, function (err, org: IOrganization) {
                if (err) {
                    failed = true;
                    if (servers.DEBUG) console.error(err);
                    return res.status(500).send({message: errors.internalServerError});
                }
                org.experience_validations.push(new IValidations(req.account.username, req.account.experiences[req.account.experiences.length - 1]._id)); // add validation entry to organization
                org.save(function (err) {
                    if (err) {
                        if (servers.DEBUG) console.error(err);
                        return res.status(500).send({message: errors.internalServerError});
                    }
                }); // save to db
            });
        }
        if (!res.headersSent && !failed) return res.status(200).send({message: errors.ok}); // send ok header if all is good
    }

    public static deleteExperience(req, res) {
        let accType = req.accountType;
        if (accType != "User") return res.status(400).send({message: errors.badRequest + " (Incorrect account type! User account type required.)"});

    }

    public static getExperienceValidations(req, res) {

    }

    public static reviewExperienceValidations(req: Request, res) {

    }
}