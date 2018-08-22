// User Interface
import { IEntity } from "./entity";

export interface IUser extends IEntity {
    firstname?: string;
    lastname?: string;
    metainfo?: object;
    connections?: string[];
    intrests?: string[];
    experiences?: object[];
    settings?: object;

}
