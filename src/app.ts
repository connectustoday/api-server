// Express App ConnectUS
import * as express from "express";
import * as bodyParser from "body-parser";
import { Routes } from "./routes/routes"
import {AuthRoutes} from "./routes/auth/auth-routes"

class App {

    public app: express.Application;
    public routes: Routes = new Routes();
    public authRoutes: AuthRoutes = new AuthRoutes();

    constructor() {
        this.app = express();
        this.config();

        // API Endpoints
        this.routes.routes(this.app);
        this.authRoutes.routes(this.app);
    }

    private config(): void {

        // CHANGE THIS IN PRODUCTION! CORS policy
        this.app.use(function (req, res, next) {
            res.header("Access-Control-Allow-Origin", "*");
            res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
            next();
        });

        // support application/json type post data
        this.app.use(bodyParser.json());

    }

}

export default new App().app;
