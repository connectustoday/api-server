// Express App ConnectUS
import * as express from "express";
import * as bodyParser from "body-parser";
import { Routes } from "./routes/routes"
import { AuthRoutes } from "./routes/auth/authRoutes"

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

    private config(): void{
        // support application/json type post data
        this.app.use(bodyParser.json());

        //support application/x-www-form-urlencoded post data
        this.app.use(bodyParser.urlencoded({ extended: false }));
    }

}

export default new App().app;