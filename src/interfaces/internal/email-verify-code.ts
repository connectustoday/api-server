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

import {Document, model, Model, Schema} from "mongoose";

export default interface IEmailVerifyCode extends Document {
    code: string,
    username: string
}

export const EmailVerifyCodeSchema = new Schema({
    code: { type: String, required: true, index: true, unique: true },
    username: { type: String, required: true },
    createdAt: { type: Date, expires: '2d', default: Date.now }
});


//AccountSchema.set('autoIndex', false);

export const EmailVerifyCodeModel: Model<IEmailVerifyCode> =  model<IEmailVerifyCode>("EmailVerifyCodeModel", EmailVerifyCodeSchema);
