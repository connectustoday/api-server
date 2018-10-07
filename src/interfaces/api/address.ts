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


import IPointAPI from "./point";
import IAddress from "../internal/address";

export default class IAddressAPI {
    street?: string;
    city?: string;
    province?: string;
    country?: string;
    postal_code?: string;
    apt_number?: string;
    geojson?: IPointAPI;

    public constructor(address: IAddress) {
        if (address == undefined) return;
        this.street = address.street;
        this.city = address.city;
        this.province = address.province;
        this.country = address.country;
        this.postal_code = address.postal_code;
        this.apt_number = address.apt_number;
        this.geojson = IPointAPI.constructPointFromInternal(address.geojson);
    }
}