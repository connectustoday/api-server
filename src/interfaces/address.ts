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
import IPoint from './point';

export default class IAddress {

    @prop({ required: true })
    schema_version: number;

    @prop()
    street?: string;

    @prop()
    city?: string;

    @prop()
    province?: string;

    @prop()
    country?: string;

    @prop()
    postal_code?: string;

    @prop()
    apt_number?: string;

    @prop()
    geojson?: IPoint;
}
