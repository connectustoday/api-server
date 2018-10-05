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

export class ExperiencesUtil {
    public static getExperiences(req, res): any {
        let accType = req.accountType;
        if (accType != "user") res.status(400).send({message: errors.badRequest + " (Incorrect account type! User account type required.)"});

    }

    public static createExperience(req, res) {
        let accType = req.accountType;
        if (accType != "user") res.status(400).send({message: errors.badRequest + " (Incorrect account type! User account type required.)"});

    }

    public static deleteExperience(req, res) {
        let accType = req.accountType;
        if (accType != "user") res.status(400).send({message: errors.badRequest + " (Incorrect account type! User account type required.)"});

    }

    public static getExperienceValidations(req, res) {

    }

    public static reviewExperienceValidations(req: Request, res) {

    }
}