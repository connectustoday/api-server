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

import INotification, { NotificationSchema } from './notification';
import IAccountSettings, { AccountSettingsSchema } from './account-settings';
import * as mongoose from "mongoose";
import { Document, Schema, Model, model } from "mongoose";
import * as passportLocalMongoose from "passport-local-mongoose";

export default interface IAccount extends Document {
    schema_version: number;
    //    id: string;
    username: string;
    email: string;
    password: string;
    oauth_token?: string;
    oauth_service?: string;
    is_email_verified: boolean;
    last_login: number;
    notifications: Array<INotification>;
    avatar: string;
    header: string;
    created_at: number;
    pending_connections?: Array<string>;
    requested_connections?: Array<string>;
    posts?: Array<string>;
    liked?: Array<[string, number]>;
    shared?: Array<[string, number]>;
    settings: IAccountSettings;
    type: string; // new
    admin_note?: string;
}

let comSchema = new Schema({ posts: String, when: Number }, { _id: false });

export const AccountSchema = new Schema({
    schema_version: { type: String, required: true },
    //    id: {type: String, required: true, index: true, unique: true},
    username: { type: String, required: true, index: true, unique: true, trim: true, collation: { locale: 'en', strength: 2 } },
    email: { type: String, required: true, index: true },
    password: { type: String, required: true },
    oauth_token: { type: String },
    oauth_service: { type: String },
    is_email_verified: { type: Boolean, required: true },
    last_login: { type: Number, required: true },
    notifications: { type: [NotificationSchema], required: true },
    avatar: { type: String, required: true, index: true },
    header: { type: String, required: true },
    created_at: { type: Number, required: true },
    pending_connections: { type: [String] },
    requested_connections: { type: [String] },
    posts: { type: [String] },
    liked: { type: [comSchema] },
    shared: { type: [comSchema] },
    settings: { type: AccountSettingsSchema, required: true },
    admin_note: { type: String, index: true }
}, { discriminatorKey: 'type', id: false });

AccountSchema.plugin(passportLocalMongoose);
//AccountSchema.set('autoIndex', false);

export const AccountModel: Model<IAccount> = model<IAccount>("AccountModel", AccountSchema);