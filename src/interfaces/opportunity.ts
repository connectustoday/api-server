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

import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';
import IAddress from './address';

export default class IOpportunity extends Typegoose {

    @prop({ required: true })
    schema_version: number;

    @prop({ required: true })
    id: string;

    @prop({ required: true })
    organization: string;

    @prop({ required: true })
    name: string;

    @prop()
    description?: string;

    @prop()
    address?: IAddress;

    @prop({ required: true })
    is_signups_enabled: boolean;

    @prop()
    number_of_people_needed?: number;

    @prop()
    tags?: Array<string>;

    @prop()
    interested_users: Array<string>; //TODO

    @prop()
    shift_times: Array<[string, string]>;

    @prop()
    method_of_contact?: string;

    @prop({ required: true })
    created_at: number;
}
