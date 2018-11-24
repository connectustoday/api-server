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

import app from "./app";
import * as adapter from "./mongo/adapter";
import {Mailer} from "./mail/mailer";

function getArg(env: string, def: string): string {
    return (env == null) || (env == undefined) ? def : env;
}

console.log("Starting ConnectUS API Server...");

// MongoDB runtime flags
export const DB_PORT: string = getArg(process.env.DB_PORT, "27017"),
    DB_ADDRESS: string = getArg(process.env.DB_ADDRESS, "localhost"),
    DB_NAME: string = getArg(process.env.DB_NAME, "database"),
    SECRET: string = getArg(process.env.SECRET, "defaultsecret"),
    REGISTER_VERIFY_SECRET: string = getArg(process.env.REGISTER_VERIFY_SECRET, "defaultsecret"),
    APPROVAL_VERIFY_SECRET: string = getArg(process.env.APPROVAL_VERIFY_SECRET, "defaultsecret"),
    TOKEN_EXPIRY: number = parseInt(getArg(process.env.TOKEN_EXPIRY, "86400"), 10),
    // Mail Options
    MAIL_USERNAME: string = getArg(process.env.MAIL_USERNAME, "user"),
    MAIL_PASSWORD: string = getArg(process.env.MAIL_PASSWORD, "pass"),
    MAIL_SENDER: string = getArg(process.env.MAIL_SENDER, "user@host.com"),
    SMTP_HOST: string = getArg(process.env.SMTP_HOST, "host.com"),
    SMTP_PORT: number = parseInt(getArg(process.env.SMTP_PORT, "587"), 10),
    API_DOMAIN: string = getArg(process.env.API_DOMAIN, "localhost:3000"),
    SITE_DOMAIN: string = getArg(process.env.SITE_DOMAIN, "localhost:3000"),
    DEBUG: boolean = JSON.parse(getArg(process.env.DEBUG, "false"));

const PORT: string = getArg(process.env.API_PORT, "3000");

adapter.connectDB();  // MongoDB connection
Mailer.mailer = new Mailer(MAIL_USERNAME, MAIL_PASSWORD, SMTP_HOST, SMTP_PORT, MAIL_SENDER); // Initialize mailer
//Mailer.mailer.sendMail("devinhanlin@gmail.com", "This is a test mail", "This is plain text", "<strong> This is strong html text.</strong>");

app.listen(parseInt(PORT, 10), () => {
    console.log("ConnectUS API Server listening on port " + PORT);
});