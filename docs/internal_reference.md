# Internal Reference (v1)
This page documents the internal workings of the API server.

## Data Types (Represented in MongoDB)

### Account
A superschema of users and organizations. It represents any sort of registered account on the 
site. 
They share the same pool of IDs and similar behaviour. 

| Field | Type | Description |
|-------|:----:|-------------|
| `schema_version` | int32 | The schema version of the `Account`. |
| `id`  | string | The unique ID of the `Account`. |
| `username` | string | The unique user of the account, can be identified by @ (ex. @bayviewss) |
| `email` | string | Account's email; used for sign-in and notifications. |
| `password` | string | The hash for the entity's password. |
| `oauth_token` | string | API token for OAuth (Google, Facebook sign-in) |
| `email_verified` | bool | Whether or not the `Account` has verified their email (to activate the account). |
| `last_login` | int64 | Last login time; stored as a Unix timestamp |
| `notifications` | `Notification` array | Stores the list of `Notification`s that the entity has. |
| `date_account_created` | int64 | Timestamp for when the account was created. |
| `pending_connections` | string array | Pending connection requests from other `Account`s. Array of IDs. | 
| `requested_connections` | string array | Pending connection requests to other `Account`s. Array of IDs. |
| `settings` | `AccountSettings` | The settings for the `Account`. Will be inherited schema (`UserSettings` or `OrganizationSettings`) |

### AccountSettings
Schema representing shared account settings between `User`s and `Organization`s.

| Field | Type | Description |
|-------|:----:|-------------|
| `allow_messages_from_unknown` | bool | Whether or not to allow messages from `Account`s that are not connected with this `Account`.

### User (extends `Account`)
Schema representing registered users on the site. It inherits the fields from `Account`.

| Field | Type | Description |
|-------|:----:|-------------|
| `first_name` | string | The first name of the user. | 
| `middle_name` | string | The middle name of the user, if applicable. (Otherwise will be blank) |
| `last_name` | string | The last name of the user. |
| `birthday` | string | The birthdate of the user (MM/DD/YYYY) |
| `personal_info` | `UserPersonalInfo` | The personal information of an user. |

### UserSettings (extends `AccountSettings`)
Schema representing settings specific to `User`s.

| Field | Type | Description |
|------ |:----:|-------------|
| `is_full_name_visible` | bool | Whether or not the user allows others to see its full name. |
| `email_notifications` | bool | Whether or not the user allows emails regarding notifications. |
| `blocked_users` | string array | Array of `Account` IDs for blocked users. |

### UserPersonalInfo

### Organization (extends `Account`)

