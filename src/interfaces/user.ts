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
import IAccount from './account';
import IUserProfile from './user-profile';
import IExperience from './experience';

export default class IUser extends IAccount {
    @prop({ required: true })
    first_name: string;

    @prop()
    middle_name?: string;

    @prop()
    last_name?: string;

    @prop({ required: true })
    birthday: string;

    @prop()
    gender?: string;

    @prop({ required: true })
    personal_info: IUserProfile;

    @prop()
    experiences?: Array<IExperience>;
}
