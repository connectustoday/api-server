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

import IPoint, { PointSchema } from './point';
import * as mongoose from "mongoose";
import { Document, Schema, Model, model } from "mongoose";

export default interface IAddress extends mongoose.Document {
    schema_version: number;
    street?: string;
    city?: string;
    province?: string;
    country?: string;
    postal_code?: string;
    apt_number?: string;
    geojson?: IPoint;
}

export const AddressSchema = new mongoose.Schema({
    schema_version: { type: Number, required: true },
    street: { type: String },
    city: { type: String },
    province: { type: String },
    country: { type: String },
    postal_code: { type: String },
    apt_number: { type: String },
    geojson: { type: PointSchema }
});

export const AddressModel: Model<IAddress> = model<IAddress>("AddressModel", AddressSchema);