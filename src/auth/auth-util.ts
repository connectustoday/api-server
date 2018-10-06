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

import * as server from "../server";
import {AccountModel} from "../interfaces/internal";
import * as errors from "../routes/errors";
import * as jwt from "jsonwebtoken";

export class AuthUtil {

    // Verify token middleware
    // Returns the account type, or null if the query failed.
    public static verifyToken(req, res, next): string {
        let token = req.headers["x-access-token"];
        if (!token) {
            res.status(401).send({ auth: false, message: "No token provided." });
            return null;
        }

        jwt.verify(token, server.SECRET, function (err, decoded) {
            if (err) return res.status(500).send({ auth: false, message: "Failed to authenticate token." });

            req.decodedToken = decoded;
            next();
        });
        return null;
    }

    // Verify account token
    // Sets the request's accountType field to the account's type
    // Sets the request's decodedToken field to the decoded token
    // Sets the request's account field to the account object
    public static verifyAccount(req, res, next): string {
        let token = req.headers["x-access-token"];
        if (!token) {
            res.status(401).send({ auth: false, message: "No token provided." });
            return null;
        }

        jwt.verify(token, server.SECRET, function (err, decoded) {
            if (err) return res.status(500).send({ auth: false, message: "Failed to authenticate token." });

            AccountModel.findOne({ username: decoded.username }, { password: 0 }, function (err, user) { //TODO switch to id
                if (err) return res.status(500).send(errors.internalServerError + " (Problem finding account)");
                if (!user) return res.status(404).send(errors.notFound + " (Account not found)");

                req.account = user;
                req.decodedToken = decoded;
                req.accountType = user.type;
                next();
            });
        });
    }
}