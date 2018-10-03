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
import IAddress, { AddressSchema } from './address';
import { Document, Schema, Model, model } from "mongoose";

export default interface IOrganizationProfile extends mongoose.Document {
    schema_version: string;
    mission?: string;
    quote?: string;
    address?: IAddress;
    affiliated_orgs?: Array<string>;
    interests?: Array<string>;
}

export const OrganizationProfileSchema = new mongoose.Schema({
    schema_version: { type: String, required: true },
    mission: { type: String },
    quote: { type: String },
    address: { type: AddressSchema },
    affiliated_orgs: { type: [String] },
    interests: { type: [String] }
});
export const OrganizationProfileModel: Model<IOrganizationProfile> = model<IOrganizationProfile>("OrganizationProfileModel", OrganizationProfileSchema);