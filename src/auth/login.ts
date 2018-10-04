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

import {AccountModel} from "../interfaces/internal";
import * as errors from "../routes/errors";
import * as server from "../server";
import * as jwt from "jsonwebtoken";
import * as bcrypt from "bcryptjs";

export function login(req, res) {
    AccountModel.findOne({ username: req.body.username }, function (err, user) {
        if (err) return res.status(500).send(errors.internalServerError);
        if (!user) return res.status(404).send(errors.notFound + " (No user found.)");

        let passwordIsValid: boolean = bcrypt.compareSync(req.body.password, user.password);
        if (!passwordIsValid) {
            return res.status(401).send({ auth: false, token: null });
        }

        let token = jwt.sign({ username: user.username }, server.SECRET, {
            expiresIn: 86400 //TODO CONFIGURABLE
        });
        res.status(200).send({ auth: true, token: token });
    });
}