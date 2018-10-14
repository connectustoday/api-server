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
import IAddress, {AddressSchema} from "./address";
import {Document, Schema, Model, model} from "mongoose";
import {AccountSchema} from "./account";

export default interface IExperience extends Document {
    schema_version: number;
    location?: IAddress;
    name: string;
    organization?: string;
    opportunity?: string;
    description?: string;
    when?: [string, string];
    is_verified: boolean;
    created_at: number;
    hours: number;
}

export const ExperienceSchema = new Schema({
    schema_version: {type: Number, required: true},
    location: {type: AddressSchema},
    name: {type: String},
    organization: {type: String, index: true},
    opportunity: {type: String, index: true},
    description: {type: String},
    when: {type: {begin: String, end: String}},
    is_verified: {type: Boolean, required: true},
    created_at: {type: Number, required: true},
    hours: {type:Number}
});

//ExperienceSchema.set('autoIndex', false);

export const ExperienceModel: Model<IExperience> = model<IExperience>("ExperienceModel", ExperienceSchema);