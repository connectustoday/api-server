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
import IAccountSettings, { AccountSettingsModel } from './account-settings';
import { Document, Schema, Model, model } from "mongoose";

export default interface IUserSettings extends IAccountSettings {
    is_full_name_visible: boolean;
    blocked_users: Array<string>;
}
export const UserSettingsSchema = AccountSettingsModel.discriminator("UserSettings", new mongoose.Schema({
    is_full_name_visible: { type: Boolean, required: true },
    blocked_users: { type: [String], required: true }
})).schema;
//export const UserSettingsModel: Model<IUserSettings> = model<IUserSettings>("UserSettingsModel", UserSettingsSchema);
export const UserSettingsModel: Model<IUserSettings> = AccountSettingsModel.discriminator("UserSettings", UserSettingsSchema);