// Organization Interface
import { IEntity } from "./entity";

export interface IOrganization extends IEntity {
    name?: string;
    metainfo?: object;
    orgverified?: string;
    intrests?: string[];
    opportunities?: object[];
    settings?: object;

}
