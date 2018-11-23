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

import * as nodemailer from 'nodemailer';
import * as servers from "../server";

export class Mailer {

    public static mailer: Mailer;
    // Use SMTP for now

    public transporter;
    public sender: string; // email to send from

    constructor(username: string, password: string, host: string, port: number, sender: string) {
        this.transporter = nodemailer.createTransport({
            pool: true,
            host: host,
            port: port,
            secure: false,
            auth: {
                user: username,
                pass: password
            }
        });
        this.transporter.verify ((err, success) => {
           if (err) {
               return console.log(err);
           }
           console.log("Verified SMTP configuration!");
        });
    }

    // @ts-ignore
    public async sendMail(recipient: string, subject: string, text: string, html: string) {
        let mail = {
            from: this.sender,
            to: recipient,
            subject: subject,
            text: text,
            html: html
        };
        try {
            await this.transporter.sendMail(mail);
        } catch (err) {
            throw err;
        }
        if (servers.DEBUG) console.log("Sent mail " + mail.subject);
        return;
    }

}