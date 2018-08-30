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