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
import * as register from "../../auth/register";
import * as login from "../../auth/login";
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

        /*
         * Test utility to check if logged in
         */

        // @ts-ignore
        if (server.DEBUG) app.get(prefix + "/me", AuthUtil.verifyAccount, (req, res) => res.status(200).send(req.account));

        /*
        * Login Endpoint Required Fields
        * - username
        * - password
        * Returns 200 + auth=true + token if SUCCESSFUL
        * TODO EMAIL LOGIN (RATHER THAN USERNAME)
        */

        app.post(prefix + "/login", (req, res) => login.login(req, res));
    }
}
