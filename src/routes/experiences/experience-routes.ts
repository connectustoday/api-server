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
import {AuthUtil} from "../../auth/auth-util";

export class ExperienceRoutes {
    public static routes(app: express.Application, prefix: string): void {

        /*
         * Experience Routes
         */

        // Get current user's experiences
        app.get(prefix, AuthUtil.verifyAccount, (req, res) => ExperiencesUtil.getExperiences(req, res));

        // Create experience
        // Use IExperienceAPI object as "experience" field
        app.post(prefix, AuthUtil.verifyAccount, (req, res) => ExperiencesUtil.createExperience(req, res));

        // Update experience
        app.put(prefix + "/:id", AuthUtil.verifyAccount, (req, res) => {});

        // Delete experience
        app.delete(prefix + "/:id", (req, res) => ExperiencesUtil.deleteExperience(req, res));

        // List pending experience validations (for organization)
        app.get(prefix + "/validations", AuthUtil.verifyAccount, (req, res) => ExperiencesUtil.getExperienceValidations(req, res));

        // Approve or don't approve validation (for organization)
        app.post(prefix + "/validations/:id", AuthUtil.verifyAccount, (req, res) => {

        });
    }
}