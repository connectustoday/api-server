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
import * as jwt from "jsonwebtoken";
import IUser from "../interfaces/internal/user";
import IExperienceAPI from "../interfaces/api/experience";
import IAddressAPI from "../interfaces/api/address";
import IValidationsAPI from "../interfaces/api/validations";
import {sendError} from "../routes/errors";
import {Mailer} from "../mail/mailer";

export class ExperiencesUtil {

    /*
     * REST API handlers
     */

    // Get experiences of personal user (username found from token)
    public static async getPersonalExperiences(req, res) {
        if (req.accountType != "User") return sendError(res, 400, errors.badRequest + " (Incorrect account type! User account type required.)", 4000);
        req.params.id = req.account.username;
        await this.getExperiences(req, res);
    }

    // Get all experiences of any user
    // @ts-ignore
    public static async getExperiences(req, res) {
        let user: IUser;
        try {
            // @ts-ignore
            user = await AccountModel.findOne({username: req.params.id, type: "User"}); // TODO CASE INSENSITIVE QUERY
        } catch (err) {
            if (servers.DEBUG) console.error(err);
            return sendError(res, 500, errors.internalServerError, 4001);
        }
        if (!user) {
            return sendError(res, 404, errors.notFound + " (User not found? Is this the correct account type?)", 4002);
        }
        let object: Array<IExperienceAPI> = [];
        user.experiences.forEach((element) => {
            object.push(new IExperienceAPI(new IAddressAPI(element.location), element._id, element.name, element.organization, element.opportunity, element.description, JSON.parse(element.when), element.is_verified, element.email_verify, element.created_at, element.hours));
        });
        res.status(200).send(object);
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
            when: JSON.stringify({
              begin: experience.when.begin,
              end: experience.when.end
            }),
            hours: experience.hours,
            is_verified: false,
            email_verify: experience.email_verify,
            created_at: (new Date).getTime()
        });
    }

    public static async updateExperience(req, res) { //TODO REDO THE CODE (MULTIPLE SAVES IN PARALLEL), SAVE AND TRANSFER APPROVAL FROM ORGANIZATION IF ORGANIZATION REMAINS THE SAME
        await this.deleteExperience(req, res, false);
        if (!res.headersSent) await this.createExperience(req, res, req.params.id, true);
    }

    public static async createExperience(req, res, id, save: boolean) {
        if (req.accountType != "User") return sendError(res, 400, errors.badRequest + " (Incorrect account type! User account type required.)", 4000);

        // default experience object

        let exp = this.getDefaultExperience(req.body, id);

        // verifications for data

        if (req.body.opportunity != undefined && req.body.opportunity != "") {
            // TODO OPPORTUNITY
        }

        if (req.body.organization != undefined && req.body.organization != "") { // if the organization field is not empty
            if (req.body.email_verify) { // send email request to organization not on site (for validations)

                exp.emailjwt = jwt.sign({ username: req.account.username, ms: Date.now() }, servers.APPROVAL_VERIFY_SECRET, {
                    expiresIn: 604800 // 1 week
                }); // create token for validation
                let verifyLink = servers.API_DOMAIN + "/v1/experiences/email_approve/" + exp.emailjwt; // generate verification link with code

                try {
                    await Mailer.mailer.sendMail(req.body.organization, "Volunteer or Work Experience Validation Request",
                        "A user from ConnectUS is requesting validation for their experience! You can approve the submission here: " + verifyLink,
                        "validate_experience_without_account", {
                            verifyLink: verifyLink,
                            website: servers.SITE_DOMAIN,
                            email: req.account.email,
                            userName: req.account.username,
                            fullName: (req.account.first_name ? req.account.first_name + " " : "") + (req.account.middle_name ? req.account.middle_name + " " : "") + (req.account.last_name ? req.account.last_name + " " : ""),
                            expName: exp.name,
                            expHours: exp.hours,
                            expStart: JSON.parse(exp.when).begin,
                            expEnd: JSON.parse(exp.when).end,
                            expDesc: exp.description,
                            random: Math.random()
                        }); // send validation request by email
                    console.log("string:" + JSON.stringify(exp));
                } catch (err) {
                    if (servers.DEBUG) console.error(err);
                    return sendError(res, 500, errors.internalServerError + " (Issue sending mail)", 4003);
                }

            } else { // check if there is an associated organization on the site (for validations)

                // add experience to organization pending validations list
                // TODO NOTIFICATION ON PENDING VALIDATION
                // TODO DUPLICATE HEADERS SENT WHEN SAVING FAILURE

                let org: IOrganization;
                try {
                    // @ts-ignore
                    org = await AccountModel.findOne({username: req.body.organization, type: "Organization"});// TODO CASE INSENSITIVE LOOKUPS
                } catch (err) {
                    if (servers.DEBUG) console.error(err);
                    return sendError(res, 500, errors.internalServerError, 4001);
                }
                if (!org) {
                    return sendError(res, 404, errors.badRequest + " (Organization not found.)", 4002);
                }

                org.experience_validations.push(new IValidations(req.account.username, exp._id)); // add validation entry to organization

                try {
                    await org.save(); // save to db
                } catch (err) {
                    if (servers.DEBUG) console.error(err);
                    return sendError(res, 500, errors.internalServerError, 4001);
                }

            }
        }

        req.account.experiences.push(exp); // add to user's experiences array
        // finish adding experience to database
        if (save) await ExperiencesUtil.saveAccountMongo(req, res);
    }

    /*
     * Save the account
     */

    private static async saveAccountMongo(req, res, call?) {
        if (call) call();
        try {
            await req.account.save(); // save to db
        } catch (err) {
            if (servers.DEBUG) console.error(err);
            return sendError(res, 500, errors.internalServerError, 4001);
        }
        if (!res.headersSent) return res.status(200).send({message: errors.ok}); // send ok header if all is good
    }

    // TODO BUG IF ORGANIZATION IS REMOVED ALREADY
    public static async deleteExperience(req, res, save: boolean) {
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
        if (experience.organization != undefined && experience.organization != "" && !experience.email_verify) { // remove pending requests for experience
            let org: IOrganization;

            try {
                // @ts-ignore
                org = await AccountModel.findOne({username: experience.organization, type: "Organization"}); // get organization from db
            } catch (err) {
                if (servers.DEBUG) console.error(err);
                return sendError(res, 500, errors.internalServerError, 4001);
            }

            if (!org) {
                if (save) return await ExperiencesUtil.saveAccountMongo(req, res, func);
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
                try {
                    await org.save(); // save to db
                } catch (err) {
                    if (servers.DEBUG) console.error(err);
                    return sendError(res, 500, errors.internalServerError, 4001);
                }
            }

        }
        if (save) return await ExperiencesUtil.saveAccountMongo(req, res, func);
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

    public static async emailApproveExperienceValidation(req, res) {
        res.set('Content-Type', 'text/html');
        jwt.verify(req.params.token, servers.APPROVAL_VERIFY_SECRET, async function (err, decoded) {
            if (err) return res.status(404).send("Invalid approval link. Perhaps it has expired?");

            let user: IUser;
            try {
                // @ts-ignore
                user = await AccountModel.findOne({username: decoded.username}, {password: 0}); // TODO switch to id
            } catch (err) {
                if (servers.DEBUG) console.error(err);
                return res.status(500).send("Internal server error. Problem finding account.");
            }

            if (!user) return res.status(500).send("Account not found. Perhaps the user has been removed?");

            let found = false;
            for (let i = 0; i < user.experiences.length; i++) {
                if (user.experiences[i].emailjwt && jwt.verify(user.experiences[i].emailjwt, servers.APPROVAL_VERIFY_SECRET).ms == decoded.ms) { // compare timestamp
                    found = true;
                    user.experiences[i].is_verified = true; // verify experience
                    user.experiences[i].emailjwt = undefined;
                    break;
                }
            }
            if (!found) return res.status(500).send("Could not find experience to validate. Perhaps the user has removed it, or the experience was already validated?");

            try {
                await user.save();
            } catch (err) {
                if (servers.DEBUG) console.error(err);
                return res.status(500).send("Internal server error.");
            }
            return res.status(200).send("You have successfully approved the request! Sign up for ConnectUS to approve and manage validations directly from the site...<script>setTimeout(()=>{window.location.replace('" + servers.SITE_DOMAIN + "/auth/login.php')}, 5000)</script>")
        });
    }
}