import * as express from "express";

export class Routes {       
    public routes(app): void {          
        app.get('/', (req, res) => res.send('ConnectUS Backend API Server - Working!')); // Default (root) route.       
    }
}