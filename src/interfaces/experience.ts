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
import IAddress, {AddressSchema} from './address';

export default interface IExperience {
    schema_version: number;
    location?: IAddress;
    id: string;
    organization?: string;
    opportunity?: string;
    description?: string;
    when?: [string, string];
    is_verified: boolean;
    created_at: number;
}

export const ExperienceSchema = new mongoose.Schema({
    schema_version: {type: Number, required: true},
    location: {type: AddressSchema},
    id: {type: String, required: true},
    organization: {type: String},
    opportunity: {type: String},
    description: {type: String},
    when: {type: {begin: String, end: String}},
    is_verified: {type: Boolean, required: true},
    created_at: {type: Number, required: true}
});