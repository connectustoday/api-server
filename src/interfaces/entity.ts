// Entity Object Interface

export interface IEntity {
    id?: string;
    username?: string;
    email?: string;
    password?: string;
    token?: string;
    verified?: boolean;
    logintime?: number;
    notifs?: object;

}
