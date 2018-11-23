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


import IAddressAPI from "./address";

export default class IExperienceAPI {
    location: IAddressAPI;
    id: string;
    name: string;
    organization: string;
    opportunity: string;
    description: string;
    when: [string, string];
    is_verified: boolean;
    email_verify: boolean;
    created_at: number;
    hours: number;
    public constructor(location: IAddressAPI, id: string, name: string, organization: string, opportunity: string, description: string, when: [string, string], is_verified: boolean, email_verify: boolean, created_at: number, hours: number) {
        this.location = location;
        this.id = id;
        this.name = name;
        this.organization = organization;
        this.opportunity = opportunity;
        this.description = description;
        this.when = when;
        this.is_verified = is_verified;
        this.email_verify = email_verify;
        this.created_at = created_at;
        this.hours = hours;
    }
}