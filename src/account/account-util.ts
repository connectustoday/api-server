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
import {sendError} from "../routes/errors";
import * as errors from "../routes/errors";
import IAccountAPI from "../interfaces/api/account";

export class AccountUtil {

    // verifies if the username does not exist in the database
    public static verifyUniqueUsername(username: string, callback) { //TODO CASE INSENSITIVE
        AccountModel.count({username: username}, function (err, count) {
            if (err) console.log(err);
            callback(count <= 0);
        });
    }

    // Account API Route Logic

    public static getAccount(req, res) {
        AccountModel.findOne({ username: req.params.id }, { password: 0 }, function (err, user) { //TODO switch to id
            if (err) return sendError(res, 500, errors.internalServerError + " (Problem finding account)", 3002);
            if (!user) return sendError(res, 404, errors.notFound + " (Account not found)", 3003);

            try {
                let acc = new IAccountAPI(user.username, user.email, user.avatar, user.header, user.created_at, user.type, user.posts.length, user.liked.length, user.shared.length);
                res.status(200).send(acc);
            } catch (error) {
                //TODO ERROR
            }
        });
    }

    public static getAccountProfile(req: any, res: any) {

    }
}