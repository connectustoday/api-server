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
import IAttachment, { AttachmentSchema } from './attachment';
import { Document, Schema, Model, model } from "mongoose";

export default interface IPost extends mongoose.Document {
    schema_version: number;
    id: string;
    account: string;
    content: string;
    created_at: number;
    reply_to?: string;
    multimedia?: IAttachment;
    tags?: Array<string>;
    likes_count: number;
    comments_count: number;
    shares_count: number;
    likes?: Array<string>;
    comments?: Array<string>;
    shares?: Array<string>;
    visibility: string;
}
export const PostSchema = new mongoose.Schema({
    schema_version: { type: Number, required: true },
    id: { type: String, required: true, index: true },
    account: { type: String, required: true },
    content: { type: String, required: true },
    created_at: { type: Number, required: true },
    reply_to: { type: String },
    multimedia: { type: AttachmentSchema },
    tags: { type: [String], index: true }, // TODO
    likes_count: { type: Number, required: true },
    comments_count: { type: Number, required: true },
    shares_count: { type: Number, required: true },
    likes: { type: [String] },
    comments: { type: [String] },
    shares: { type: [String] },
    visibility: { type: String, required: true }
});
export const PostModel: Model<IPost> = model<IPost>("PostModel", PostSchema);