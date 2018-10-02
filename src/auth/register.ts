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

import {AccountModel} from "../interfaces";
import * as errors from "../routes/errors";
import * as server from "../server";
import * as bcrypt from "bcryptjs";
import * as jwt from "jsonwebtoken";

export function registerRequest(req, res) {
    let hashedPassword = bcrypt.hashSync(req.body.password, 8);

    AccountModel.create({ // Default user
        schema_version: "1",
        username: req.body.name,
        email: req.body.email,
        password: hashedPassword
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