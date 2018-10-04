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

export default class IAccount {
//    id: string;
    username: string;
    email: string;
    avatar: string;
    header: string;
    created_at: number;
    type: string;
    posts_count: number;
    liked_count: number;
    shared_count: number;

    public constructor(username: string, email: string, avatar: string, header: string, created_at: number, type: string, posts_count: number, liked_count: number, shared_count: number) {
        this.username = username;
        this.email = email;
        this.avatar = avatar;
        this.header = header;
        this.created_at = created_at;
        this.type = type;
        this.posts_count = posts_count;
        this.liked_count = liked_count;
        this.shared_count = shared_count;
    }
}