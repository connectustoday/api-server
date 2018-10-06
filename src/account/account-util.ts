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

export class AccountUtil {
    public static getAccount(req, res) {
        
    }

    // verifies if the username does not exist in the database
    public static verifyUniqueUsername(username: string, callback) { //TODO CASE INSENSITIVE
        AccountModel.count({username: username}, function (err, count) {
            if (err) console.log(err);
            callback(count <= 0);
        });
    }
}