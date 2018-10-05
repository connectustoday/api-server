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

// Authentication API Routes

// API Version 1.0
import * as server from "../../server";
import * as errors from "../errors";
import * as bcrypt from "bcryptjs";
import * as register from "../../auth/register";
import * as login from "../../auth/login";
import * as jwt from "jsonwebtoken";
import {AccountModel} from "../../interfaces/internal/account";
import express = require("express");
import {AuthUtil} from "../../auth/auth-util";

export class AuthRoutes {
    public static routes(app: express.Application, prefix: string): void {

        app.get(prefix, (req, res) => res.send(errors.badRequest));

        /*
         * Register Endpoint Required Fields
         * - username
         * - email
         * - password
         * - type ("organization", "user")
         * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
         * User Required Fields
         * - first_name
         * - birthday
         * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
         * Organization Required Fields
         * - is_nonprofit
         * - preferred_name
         * ------------------------------------
         * Returns 200 + auth=true + token if SUCCESSFUL
         * TODO EMAIL VERIFICATION
         */

        app.post(prefix + "/register", (req, res) => register.registerRequest(req, res));
        app.get(prefix + "/register", (req, res) => res.send(errors.methodNotAllowed));

        /*
         * Test utility to check if logged in
         */

        app.get(prefix + "/me", AuthUtil.verifyAccount, (req, res) => {
            // @ts-ignore
            let decoded = req.decodedToken;

            AccountModel.findOne({username: decoded.username}, {password: 0}, function (err, user) { //TODO switch to id
                if (err) return res.status(500).send(errors.internalServerError + " (Problem finding user)");
                if (!user) return res.status(404).send(errors.notFound + " (User not found)");

                res.status(200).send(user);
            });
        });

        /*
        * Register Endpoint Required Fields
        * - username
        * - password
        * Returns 200 + auth=true + token if SUCCESSFUL
        * TODO EMAIL LOGIN (RATHER THAN USERNAME)
        */

        app.get(prefix + "/login", (req, res) => res.send(errors.methodNotAllowed));
        app.post(prefix + "/login", (req, res) => login.login(req, res));
    }
}