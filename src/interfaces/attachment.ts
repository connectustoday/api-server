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
import { Document, Schema, Model, model} from "mongoose";

export default interface IAttachment {
    schema_version: number;
    type: string;
    url: string;
    description?: string;
}

export const AttachmentSchema = new mongoose.Schema({
    schema_version: {type: Number, required: true},
    type: {type: String, required: true},
    url: {type: String, required: true, index: true},
    description: {type: String}
});
export const AttachmentModel: Model<IAttachment> = model<IAttachment>("AttachmentModel", AttachmentSchema);