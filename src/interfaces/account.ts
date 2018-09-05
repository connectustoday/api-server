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
import INotification from './notification';
import IAccountSettings from './account-settings';

export default class IAccount extends Typegoose {

    @prop({ required: true })
    schema_version: number;

    @prop({ required: true })
    id: string;

    @prop({ required: true })
    username: string;

    @prop({ required: true })
    email: string;

    @prop({ required: true })
    password: string;

    @prop()
    oauth_token?: string;

    @prop()
    oauth_service?: string;

    @prop({ required: true })
    is_email_verified: boolean;

    @prop({ required: true })
    last_login: number;

    @prop({ required: true })
    notifications: Array<INotification>;

    @prop({ required: true })
    avatar: string;

    @prop({ required: true })
    header: string;

    @prop({ required: true })
    created_at: number;

    @prop()
    pending_connections?: Array<string>;

    @prop()
    requested_connections?: Array<string>;

    @prop()
    posts?: Array<string>;

    @prop()
    liked?: Array<[string, number]>;

    @prop()
    shared?: Array<[string, number]>;

    @prop({ required: true })
    settings: IAccountSettings;

    @prop()
    admin_note?: string;

}
