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
import IExperienceAPI from "../interfaces/api/experience";
import IExperience from "../interfaces/internal/experience";
import {AccountModel} from "../interfaces/internal/account";
import IUser from "../interfaces/internal/user";
import * as servers from "../server";
import IOrganization from "../interfaces/internal/organization";
import IValidations from "../interfaces/internal/organization";
import {ExperienceModel} from "../interfaces/internal/experience";
import {Promise} from "mongoose";
import * as ClassTransformer from "class-transformer";

export class ExperiencesUtil {

    /*
    *{
		"location": "hi",
		"id": "",
		"name": "lol!",
		"organization": "",
		"opportunity": "",
		"description": "this is a desc",
		"when": ["hi", "hi"],
		"is_verified": true,
		"created_at": 0
	}
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

        if (!req.account.hasOwnProperty('experiences')) req.account.experiences = [];

        let promise = Promise.resolve(), failed: boolean = false;
        let newExpID = 0; // set the id to the latest largest id plus one
        if (req.account.experiences.length != 0) newExpID = req.account.experiences[req.account.experiences.length-1].id + 1;

        // verifications for data

        if (req.body.experience.hasOwnProperty('opportunity') && req.body.experience.opportunity != "") {
            // TODO OPPORTUNITY
        }

        if (req.body.experience.hasOwnProperty('organization') && req.body.experience.organization != "") { // check if there is an associated organization on the site (for validations)
            promise = AccountModel.count({username: req.body.experience.organization, type: "Organization"}, function (err, count) {
                if (err) {
                    failed = true;
                    if (servers.DEBUG) console.error(err);
                    return res.status(500).send({message: errors.internalServerError});
                }
                if (count <= 0) {
                    failed = true;
                    return res.status(400).send({message: errors.badRequest + " (Organization not found.)"})
                }

                // add to organization pending validations list
                // TODO NOTIFICATION ON PENDING VALIDATION
                AccountModel.findOne({username: req.body.experience.organization, type: "Organization"}, function (err, org: IOrganization) {
                    if (err) {
                        failed = true;
                        if (servers.DEBUG) console.error(err);
                        return res.status(500).send({message: errors.internalServerError});
                    }
                    org.experience_validations.push(new IValidations(req.account.username, newExpID)); // add validation entry to organization
                    org.save(); // save to db
                });
            }); // TODO CASE INSENSITIVE LOOKUPS
        }

        // finish adding experience to database

        promise.then(() => {
                if (failed) return;
                // cast to experienceapi object and then internal experience object

                // default experience object
                let exp = new ExperienceModel({
                    schema_version: 0,
                    location: req.body.experience.location,
                    id: newExpID,
                    name: req.body.experience.name,
                    organization: req.body.experience.organization,
                    opportunity: req.body.experience.opportunity,
                    description: req.body.experience.description,
                    is_verified: false,
                    created_at: (new Date).getTime()
                });

                req.account.experiences.push(exp); // add to user's experience
                req.account.save(); // save to db
                res.status(200).send({message: errors.ok});
            }
        );
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