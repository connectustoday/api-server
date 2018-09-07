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

export default interface IOpportunity {
    schema_version: number;
    id: string;
    organization: string;
    name: string;
    description?: string;
    address?: IAddress;
    is_signups_enabled: boolean;
    number_of_people_needed?: number;
    tags?: Array<string>;
    interested_users: Array<string>; //TODO
    shift_times: Array<[string, string]>;
    method_of_contact?: string;
    created_at: number;
}

let shiftSchema = new mongoose.Schema({begin: String, end: String}, {_id: false});

export const OpportunitySchema = new mongoose.Schema({
    schema_version: {type: Number, required: true},
    id: {type: String, required: true, index: true},
    organization: {type: String, required: true, index: true},
    name: {type: String, required: true, index: true},
    description: {type: String},
    address: {type: AddressSchema},
    is_signups_enabled: {type: Boolean, required: true},
    number_of_people_needed: {type: Number},
    tags: {type: [String]},
    interested_users: {type: [String]},
    shift_times: {type: shiftSchema},
    method_of_contact: {type: String},
    created_at: {type: Number, required: true}
});