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
import IAccount, { AccountModel } from './account';
import IOrganizationProfile, { OrganizationProfileSchema } from './organization-profile';
import { Document, Schema, Model, model } from "mongoose";
import { AccountSettingsSchema } from "./account-settings";
import { OrganizationSettingsSchema } from "./organization-settings";

export default interface IOrganization extends IAccount {
    preferred_name: string;
    is_verified: boolean;
    opportunities?: Array<string>;
    org_info: IOrganizationProfile;
    experience_validations?: Array<[string, string]>
}

let expSchema = new Schema({
    user_id: String,
    experience_id: Number
}, { _id: false });

export const OrganizationSchema = new Schema({
    preferred_name: { type: String, required: true, index: true },
    is_verified: { type: Boolean, required: true },
    opportunities: { type: [String] },
    org_info: { type: OrganizationProfileSchema, required: true },
    experience_validations: { type: [expSchema] }
});
//export const OrganizationModel: Model<IOrganization> = model<IOrganization>("OrganizationModel", OrganizationSchema);
export const OrganizationModel = AccountModel.discriminator("Organization", OrganizationSchema);