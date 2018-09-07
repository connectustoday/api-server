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

// ConnectUS Server

import app from "./src/app";

function getArg(env: string, def: string): string {
    return (env == null) || (env == undefined) ? def : env;
}

console.log("Starting ConnectUS API Server...");

// MongoDB runtime flags
export const DB_PORT: string = getArg(process.env.DB_PORT, "3000"),
    DB_ADDRESS: string = getArg(process.env.DB_ADDRESS, "localhost"),
    DB_NAME: string = getArg(process.env.DB_NAME, "database");

const PORT: string = getArg(process.env.API_PORT, "3000");

app.listen(PORT, () => {
    console.log('ConnectUS API Server listening on port ' + PORT);
});