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
import * as server from "../server";
import { Mongoose } from "mongoose";

/*
 * Mongo connection
 */

export function connectDB(): void {

    let url: string = "mongodb://" + server.DB_ADDRESS + ":" + server.DB_PORT + "/" + server.DB_NAME;

    console.log("Connecting to MongoDB at " + url);
    let promise: Promise<Mongoose> = mongoose.connect(url, { useNewUrlParser: true });
    mongoose.connection.on("error", console.error.bind(console, "MongoDB connection error:"));
}