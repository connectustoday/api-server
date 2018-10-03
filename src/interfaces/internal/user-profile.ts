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
import { Document, Schema, Model, model } from "mongoose";

export default interface IUserProfile extends mongoose.Document {
    schema_version: string;
    interests?: Array<string>;
    biography?: string;
    education?: string; //TODO
    quote?: string;
    current_residence?: string;
    certifications?: string; //TODO
}
export const UserProfileSchema = new mongoose.Schema({
    schema_version: { type: String, required: true },
    interests: { type: [String] },
    biography: { type: String },
    education: { type: String },
    quote: { type: String },
    current_residence: { type: String },
    certifications: { type: String }
});
export const UserProfileModel: Model<IUserProfile> = model<IUserProfile>("UserProfileModel", UserProfileSchema);