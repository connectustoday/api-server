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

import * as mongoose from "mongoose";
import IAccount, {AccountModel} from './account';
import IUserProfile, {UserProfileSchema} from './user-profile';
import IExperience, {ExperienceSchema} from './experience';
import { Schema, Model, model} from "mongoose";
import {AccountSettingsSchema} from "./account-settings";
import {UserSettingsSchema} from "./user-settings";

export default interface IUser extends IAccount {
    first_name: string;
    middle_name?: string;
    last_name?: string;
    birthday: string;
    gender?: string;
    personal_info: IUserProfile;
    experiences?: Array<IExperience>;
}
export const UserSchema = new Schema({
    first_name: {type: String, required: true, index: true}, //TODO disable if private
    middle_name: {type: String},
    last_name: {type: String},
    birthday: {type: String, required: true},
    gender: {type: String},
    personal_info: {type: UserProfileSchema, required: true},
    experiences: {type: ExperienceSchema}
});
//export const UserModel: Model<IUser> = model<IUser>("UserModel", UserSchema);
export const UserModel = AccountModel.discriminator("User", UserSchema);