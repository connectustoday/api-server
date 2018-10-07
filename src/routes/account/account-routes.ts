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

import * as errors from "../errors";
import express = require("express");
import {ExperiencesUtil} from "../../experiences/experiences-util";

// REST API layout inspired by the mastodon API

export class AccountRoutes {
    public static routes(app: express.Application, prefix: string): void {
        app.get(prefix, (req, res) => res.send(errors.badRequest));

        /*
         * Global Account Routes
         */

        app.get(prefix + "/search", (req, res) => {

        });

        // Fetch an Account's basic information (Global)

        app.get(prefix + "/:id", (req, res) => {

        });

        app.get(prefix + "/:id/profile", (req, res) => {

        });

        app.get(prefix + "/:id/connections", (req, res) => {

        });

        app.get(prefix + "/:id/posts", (req, res) => {

        });

        app.get(prefix + "/:id/experiences", (req, res) => ExperiencesUtil.getExperiences(req, res));

        app.get(prefix + "/:id/opportunities", (req, res) => {

        });

        app.post(prefix + "/:id/request_connection", (req, res) => {

        });

        app.post(prefix + "/:id/accept_connection", (req, res) => {

        });

        app.post(prefix + "/:id/block", (req, res) => {

        });

        app.post(prefix + "/:id/unblock", (req, res) => {

        });

    }
}