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

import {AuthRoutes} from "./auth/auth-routes";

export class Routes {
    public static authRoutes: AuthRoutes = new AuthRoutes();
    static routes(app): void {
        app.get('/', (req, res) => res.send('ConnectUS Backend API Server - Working!')); // Default (root) route.

        // TODO USE BEST PRACTICES: https://www.owasp.org/index.php/OWASP_Cheat_Sheet_Series

        AuthRoutes.routes(app, "/v1/auth");

        app.get('/v1/')
    }
}
