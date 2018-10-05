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


import IAddress from "./address";
import {create} from "domain";

export default class IExperience {
    location: IAddress;
    id: string;
    organization: string;
    opportunity: string;
    description: string;
    when: [string, string];
    is_verified: boolean;
    created_at: number;
    public constructor(location: IAddress, id: string, organization: string, opportunity: string, description: string, when: [string, string], is_verified: boolean, created_at: number) {
        this.location = location;
        this.id = id;
        this.organization = organization;
        this.opportunity = opportunity;
        this.description = description;
        this.when = when;
        this.is_verified = is_verified;
        this.created_at = created_at;
    }
}