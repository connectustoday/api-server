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
import * as errors from  "../errors";
import * as passport from "passport";
import * as localStrategy from "passport-local";

export class AuthRoutes {
    public routes(app): void {

        app.get('/v1/auth', (req, res) => res.send(errors.badRequest));

        /* ~~ Registration Endpoint ~~ */
        app.get('/v1/auth/register', (req, res) => res.send(errors.methodNotAllowed)); // Client should not GET this endpoint.

        app.post('/v1/auth/register', function (req, res) {
            let user = req.body;
            res.send(user.username); //testing

        });

        /* ~~ Sign In Endpoint ~~ */
        app.get('/v1/auth/login', (req, res) => res.send(errors.methodNotAllowed)); // Client should not GET this endpoint.
        app.post('/v1/auth/login', passport.authenticate('local'), {successRedirect: '/', failureRedirect: '/login'});
    }
}