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

import {UserModel} from "../interfaces/internal/user";
import {OrganizationModel} from "../interfaces/internal/organization";
import * as errors from "../routes/errors";
import * as server from "../server";
import * as bcrypt from "bcryptjs";
import * as jwt from "jsonwebtoken";
import {AccountUtil} from "../account/account-util";
import {sendError} from "../routes/errors";
import {Mailer} from "../mail/mailer";

export function registerRequest(req, res) {
    if (req.body.type == "organization") {
        return registerOrganizationRequest(req, res);
    } else if (req.body.type == "user") {
        return registerUserRequest(req, res);
    } else {
        return sendError(res, 400, errors.badRequest + " (Invalid account type)", 3200);
    }
}

// @ts-ignore
export async function registerUserRequest(req, res): Promise<void> {
    let hashedPassword = bcrypt.hashSync(req.body.password, 8);

    AccountUtil.verifyUniqueUsername(req.body.username, async function (isUnique: boolean) {
        if (!isUnique) return sendError(res, 500, errors.internalServerError + " (Username taken)", 3201);

        const defUser = new UserModel({ // Default user
            type: "User",
            schema_version: 0,
            username: req.body.username,
            email: req.body.email,
            password: hashedPassword,
            is_email_verified: false,
            last_login: 0.0,
            notifications: [],
            avatar: "https://pbs.twimg.com/profile_images/1017516299143041024/fLFdcGsl_400x400.jpg", //TODO default images
            header: "https://pbs.twimg.com/profile_images/1017516299143041024/fLFdcGsl_400x400.jpg",
            created_at: (new Date).getTime(),
            settings: { //TODO NOT DISTINGUISHING BETWEEN USER SETTINGS AND ACCOUNT SETTINGS
                type: "UserSettings",
                allow_messages_from_unknown: true,
                email_notifications: true,
                is_full_name_visible: false,
                blocked_users: [],
            },
            first_name: req.body.first_name,
            birthday: req.body.birthday,
            personal_info: {
                schema_version: 0
            }
        });

        try {
            await sendVerificationEmail(req.body.username, req.body.email, res);
        } catch (err) {
            console.error("Problem sending mail: " + err);
            return sendError(res, 500, errors.internalServerError + " (There was a problem sending the verification email. Please ask a website administrator for help.)", 3204);
        }

        try {
            await defUser.save();
        } catch (err) {
            if (server.DEBUG) console.error(err);
            return sendError(res, 500, errors.internalServerError + " (There was a problem registering the account.)", 3203);
        }
        res.status(200).send();
    });
}

// @ts-ignore
export async function registerOrganizationRequest(req, res): Promise<void> {
    let hashedPassword = bcrypt.hashSync(req.body.password, 8);

    AccountUtil.verifyUniqueUsername(req.body.username, async function (isUnique: boolean) {
        if (!isUnique) return sendError(res, 500, errors.internalServerError + " (Username taken)", 3201);

        const defOrganization = new OrganizationModel({ // Default organization
            type: "Organization",
            schema_version: 0,
            username: req.body.username,
            email: req.body.email,
            password: hashedPassword,
            is_email_verified: false,
            last_login: 0,
            notifications: [],
            avatar: "https://pbs.twimg.com/profile_images/1017516299143041024/fLFdcGsl_400x400.jpg", //TODO default images
            header: "https://pbs.twimg.com/profile_images/1017516299143041024/fLFdcGsl_400x400.jpg",
            created_at: (new Date).getTime(),
            settings: {
                is_nonprofit: req.body.is_nonprofit,
                allow_messages_from_unknown: true,
                email_notifications: true,
                type: "OrganizationSettings"
            },
            preferred_name: req.body.preferred_name,
            is_verified: false,
            org_info: {
                schema_version: 0
            }
        });

        try {
            await sendVerificationEmail(req.body.username, req.body.email, res);
        } catch (err) {
            console.error("Problem sending mail: " + err);
            return sendError(res, 500, errors.internalServerError + " (There was a problem sending the verification email. Please ask a website administrator for help.)", 3204);
        }

        try {
            await defOrganization.save();
        } catch (err) {
            if (server.DEBUG) console.error(err);
            return sendError(res, 500, errors.internalServerError + " (There was a problem registering the account.)", 3203);
        }
        res.status(200).send();
    });
}

// @ts-ignore
async function sendVerificationEmail(username: string, email: string, res) {
    let token = jwt.sign({ username: username }, server.SECRET, {
        expiresIn: 172800 // 2 days
    });
    let verifyLink: string = server.API_DOMAIN + "/v1/auth/verify-email/" + token;
    try {
        await Mailer.mailer.sendMail(email, "ConnectUS Account Signup Verification Code", "Thanks for signing up! To finish the setup process, please visit " + verifyLink + ".",
            "<strong>Thank you for signing up for ConnectUS!</strong> </br> In order to activate your account, please visit the link below: </br><a href='" + verifyLink + "'>" + verifyLink + "</a>")
    } catch (err) {
        throw err;
    }
    return;
}

export async function verifyEmailRequest(req, res) {
    jwt.verify(req.params.token, server.SECRET, function (err, decoded) {
        if (err) res.status(500).send("Failed to verify token. Please contact support");
    });
    //TODO
}