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

// Authentication API Routes

// API Version 1
import * as express from "express";

export class AuthRoutes {
    public routes(app): void {

        app.get('/v1/auth', (req, res) => res.send('Authentication Endpoint - Invalid Request'));

        /* ~~ Registration Endpoint ~~ */
        app.get('/v1/auth/register', (req, res) => res.send('Invalid Request')); // Client should not GET this endpoint.

        app.post('/v1/auth/register', function (req, res) {
            var user = req.body;
            res.send(user.username); //testing

        });

        /* ~~ Sign In Endpoint ~~ */
        app.get('/v1/auth/signin', (req, res) => res.send('Invalid Request')); // Client should not GET this endpoint.

        app.post('/v1/auth/signin', function (req, res) {
            res.send("Stub");
            // retrieve the JSON using body-parser (req.body)
        });
    }
}