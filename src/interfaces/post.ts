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
import IAttachment from './attachment';

export default class IPost extends Typegoose {

    @prop({ required: true })
    schema_version: number;

    @prop({ required: true })
    id: string;

    @prop({ required: true })
    account: string;

    @prop({ required: true })
    content: string;

    @prop({ required: true })
    created_at: number;

    @prop()
    reply_to?: string;

    @prop()
    multimedia?: IAttachment;

    @prop()
    tags?: Array<string>;

    @prop({ required: true })
    likes_count: number;

    @prop({ required: true })
    comments_count: number;

    @prop({ required: true })
    shares_count: number;

    @prop()
    likes?: Array<string>;

    @prop()
    comments?: Array<string>;

    @prop()
    shares?: Array<string>;

    @prop({ required: true })
    visibility: string;

}
