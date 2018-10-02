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

import {UserModel} from "../interfaces/user";
import {OrganizationModel} from "../interfaces/organization";
import * as errors from "../routes/errors";
import * as server from "../server";
import * as bcrypt from "bcryptjs";
import * as jwt from "jsonwebtoken";

export function registerRequest(req, res) {
    if (req.body.type == "organization") {
        return registerOrganizationRequest(req, res);
    } else if (req.body.type == "user") {
        return registerUserRequest(req, res);
    } else {
        return res.status(400).send(errors.internalServerError + " (Invalid account type)");
    }
}

export function registerUserRequest(req, res) {
    let hashedPassword = bcrypt.hashSync(req.body.password, 8);

    UserModel.create({ // Default user TODO ASSIGN ID
        schema_version: "0",
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
            allow_messages_from_unknown: true,
            email_notifications: true,
            is_full_name_visible: false,
            blocked_users: []
        },
        first_name: req.body.first_name,
        birthday: req.body.birthday,
        personal_info: {
            schema_version: "0"
        }
    }, function (err, user) {
        if (err) {
            console.error(err);
            return res.status(500).send(errors.internalServerError + " (There was a problem registering the user.)");
        }

        let token = jwt.sign({id: user._id}, server.SECRET, {
            expiresIn: 86400 //TODO token expiry
        });
        res.status(200).send({auth: true, token: token});
    });
}

export function registerOrganizationRequest(req, res) {
    let hashedPassword = bcrypt.hashSync(req.body.password, 8);

    OrganizationModel.create({ // Default user TODO ASSIGN ID
        schema_version: "0",
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
            is_nonprofit: req.body.is_nonprofit
        },
        preferred_name: req.body.first_name,
        is_verified: false,
        org_info: {
            schema_version: "0"
        }
    }, function (err, user) {
        if (err) {
            console.error(err);
            return res.status(500).send(errors.internalServerError + " (There was a problem registering the user.)");
        }

        let token = jwt.sign({id: user._id}, server.SECRET, {
            expiresIn: 86400 //TODO token expiry
        });
        res.status(200).send({auth: true, token: token});
    });
}
