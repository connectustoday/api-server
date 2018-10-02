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
import * as server from "../../server";
import * as errors from  "../errors";
import * as passport from "passport";
import * as LocalStrategy from "passport-local";
import * as jwt from "jsonwebtoken";
import * as bcrypt from "bcryptjs";
import {AccountModel} from '../../interfaces/account';

export class AuthRoutes {
    public routes(app): void {
        //passport.use(new LocalStrategy(AccountModel.createStrategy()));

        app.get('/v1/auth', (req, res) => res.send(errors.badRequest));

        app.get('/v1/auth/register', (req, res) => res.send(errors.methodNotAllowed));

        app.post('/v1/auth/register', function(req, res) {
            let hashedPassword = bcrypt.hashSync(req.body.password, 8);

            AccountModel.create({
                schema_version: "1",
                username: req.body.name,
                email: req.body.email,
                password: hashedPassword
            }, function (err, user) {
                if (err) return res.status(500).send("There was a problem registering the user.");

                var token = jwt.sign({id: user._id}, server.SECRET, {
                    expiresIn: 86400 //TODO token expiry
                });
                res.status(200).send({auth: true, token: token});
            });
        });

        app.get('/v1/auth/me', function (req, res) {
            var token = req.headers['x-access-token'];
            if (!token) return res.status(401).send({ auth: false, message: 'No token provided.' });

            jwt.verify(token, server.SECRET, function(err, decoded) {
                if (err) return res.status(500).send({ auth: false, message: 'Failed to authenticate token.' });

                AccountModel.findById(decoded.id, {password: 0}, function (err, user) { //TODO switch to id
                    if (err) return res.status(500).send(errors.internalServerError + " (Problem finding user)");
                    if (!user) return res.status(404).send(errors.notFound + " (User not found)");

                    res.status(200).send(user);
                });
            });
        });

        app.get('/v1/auth/login', (req, res) => res.send(errors.methodNotAllowed));
        app.post('/v1/auth/login', function (req, res) {
            AccountModel.findOne({email: req.body.email}, function (err, user) {
                if (err) return res.status(500).send(errors.internalServerError);
                if (!user) return res.status(404).send(errors.notFound + ' (No user found.)');

                let passwordIsValid: boolean = bcrypt.compareSync(req.body.password, user.password);
                if (!passwordIsValid) {
                    return res.status(401).send({ auth: false, token: null });
                }
                let token = jwt.sign({ id: user._id }, server.SECRET, {
                    expiresIn: 86400 //TODO CONFIGURABLE
                });
                res.status(200).send({ auth: true, token: token });
            });
        });
    }
}