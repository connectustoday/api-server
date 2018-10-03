
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

import express = require("express");
import * as errors from "../errors";

export class PersonalAccountRoutes {
    public static routes(app: express.Application, prefix: string): void {

        /*
        * Current User Routes
        */

        app.get(prefix + "/notifications", (req, res) => {

        });

        app.post(prefix + "/notification/clear", (req, res) => {

        });

        app.post(prefix + "/notification/dismiss", (req, res) => {

        });

        app.get(prefix + "/settings", (req, res) => {

        });

        app.post(prefix + "/settings", (req, res) => {

        });

        app.get(prefix + "/profile", (req, res) => {

        });

        app.post(prefix + "/profile", (req, res) => {

        });

        app.get(prefix + "/connection-requests", (req, res) => {

        })

    }
}