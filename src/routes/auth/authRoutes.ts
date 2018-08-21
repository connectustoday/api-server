// Authentication API Routes
import * as express from "express";

export class AuthRoutes {
    public routes(app): void {

        app.get('/auth', (req, res) => res.send('Authentication Endpoint - Invalid Request'));

        /* ~~ Registration Endpoint ~~ */
        app.get('/auth/register', (req, res) => res.send('Invalid Request')); // Client should not GET this endpoint.

        app.post('/auth/register', function (req, res) {
            res.send("Stub")
        });

        /* ~~ Sign In Endpoint ~~ */
        app.get('/auth/signin', (req, res) => res.send('Invalid Request')); // Client should not GET this endpoint.

        app.post('/auth/signin', function (req, res) {
            res.send("Stub")
        });
    }
}