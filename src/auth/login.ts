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

import {AccountModel} from "../interfaces/internal/account";
import * as errors from "../routes/errors";
import * as server from "../server";
import * as jwt from "jsonwebtoken";
import * as bcrypt from "bcryptjs";
import {sendError} from "../routes/errors";

export function login(req, res) {
    AccountModel.findOne({ username: req.body.username }, function (err, user) {
        if (err) return sendError(res, 500, errors.internalServerError, 3100);
        if (!user) return sendError(res, 401, "Invalid login.", 3101); // if user is not found

        if (!bcrypt.compareSync(req.body.password, user.password)) { // if password isn't valid
            return sendError(res, 401, "Invalid login.", 3101);
        }

        if (!user.is_email_verified) return sendError(res, 401, errors.unauthorized + " (Email not verified.)", 3102);

        let token = jwt.sign({ username: user.username }, server.SECRET, {
            expiresIn: server.TOKEN_EXPIRY
        });
        res.status(200).send({ token: token });
    });
}