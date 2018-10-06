import express = require("express");

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

export class OpportunityRoutes {
    public static routes(app: express.Application, prefix: string): void {

        /*
         * Opportunity Routes
         */

        // List opportunities the current organization has created
        app.get(prefix, (req, res) => {

        });

        // Create opportunity
        app.post(prefix, (req, res) => {

        });

        // Get opportunity from id
        app.get(prefix + "/:id", (req, res) => {

        });

        // Update opportunity at id
        app.put(prefix + "/:id", (req, res) => {

        });

        // Delete opportunity
        app.delete(prefix + "/:id", (req, res) => {

        });

        // TODO signups
    }
}