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
| `is_email_verified` | bool | Whether or not the `Account` has verified their email (to activate the account). |
| `last_login` | int64 | Last login time; stored as a Unix timestamp |
| `notifications` | `Notification` array | Stores the list of `Notification`s that the entity has. |
| `account_created` | int64 | Timestamp for when the account was created. |
| `pending_connections` | string array | Pending connection requests from other `Account`s. Array of IDs. | 
| `requested_connections` | string array | Pending connection requests to other `Account`s. Array of IDs. |
| `posts` | string array | List of `Post` IDs that the `Account` has posted.
| `settings` | `AccountSettings` | The settings for the `Account`. Will be inherited schema (`UserSettings` or `OrganizationSettings`) |
| `admin_note` | string | A note that an admin can leave about an account. |

### AccountSettings
Schema representing shared account settings between `User`s and `Organization`s. This is stored 
within the `Account` object.

| Field | Type | Description |
|-------|:----:|-------------|
| `allow_messages_from_unknown` | bool | Whether or not to allow messages from `Account`s that are not connected with this `Account`.
| `email_notifications` | bool | Whether or not the user allows emails regarding notifications. |

### User (extends `Account`)
Schema representing registered users on the site. It inherits the fields from `Account`.

| Field | Type | Description |
|-------|:----:|-------------|
| `first_name` | string | The first name of the user. | 
| `middle_name` | string | The middle name of the user, if applicable. (Otherwise will be blank) |
| `last_name` | string | The last name of the user. |
| `birthday` | string | The birthdate of the user (MM/DD/YYYY) |
| `personal_info` | `UserInfo` | The personal information of an user. |

### UserSettings (extends `AccountSettings`)
Schema representing settings specific to `User`s. This is stored within the `Account` object.

| Field | Type | Description |
|-------|:----:|-------------|
| `is_full_name_visible` | bool | Whether or not the user allows others to see its full name. |
| `blocked_users` | string array | Array of `Account` IDs for blocked users. |

### UserInfo
`UserInfo` represents the personal information of an user. This is stored within the `User` 
object.

| Field | Type | Description |
|-------|:----:|-------------|
| 

### Organization (extends `Account`)
Schema representing registered organizations on the site. It inherits fields from `Account`.

| Field | Type | Description |
|-------|:----:|-------------|
| `preferred_name` | string | Preferred name of the organization that shows up on their profile. |
| `is_verified` | bool | If the organization has been verified to exist, and the account belongs to the real organization. |
| `opportunities` | string array | List of `Opportunity` IDs that the organization has created. |
| `org_info` | `OrgInfo` | The organization's public information. |

### OrganizationInfo
`OrganizationInfo` represents organization specific information that is shown on their profile. 
This is stored within the `Organization` object.

| Field | Type | Description |
|-------|:----:|-------------|
| `mission` | string | The organization's stated mission. |
| `quote` | string | The organization's specified quote. |
| `address` | `Address` | The organization's headquarters location. |
| `logo` | string | Link to the organization's logo image. |
| `affiliated_orgs` | string array | IDs of other organizations this organization is affiliated with. |
| `interests` | string array | Tags that the organization is interested in. |

### OrganizationSettings (extends `AccountSettings`)
Settings specific to `Organization`s. This is stored within the `Account` object.

| Field | Type | Description |
|-------|:----:|-------------|
| `is_nonprofit` | bool | Whether or not the organization is non-profit. |


### Post
Represents posts that `Account`s create on the website.

| Field | Type | Description |
|-------|:----:|-------------|
| `post_id` | string | The unique ID of the post. |
| `account_id` | string | ID of the `Account` that posted it. |
| `content` | string | Text content of the post. |
| `time_created` | int64 | Timestamp of when the post was created. |
| `multimedia` | string array | Links to other media (images/videos/articles) that are attached to the post. Displayed separately.|
| `tags` | string array | Tags that the post is categorized under. |

### Opportunity
An opportunity object represents an opportunity that an organization has for people to sign up 
for. This can be for an event (Canada Day), or simply a shift for a job (food bank). Organizations 
can choose whether or not they want to use the built-in mechanism for signing up users. They will 
have a choice to allow users to either sign up for a shift directly on the site, or to let the 
organization know which users have expressed their interest for the position. 

| Field | Type | Description |
|-------|:----:|-------------|
| `organization_id` | string | ID of the `Organization` that created it. |
| `opportunity_id` | string | ID of the `Opportunity`. |
| `name` | string | Name of the `Opportunity`. |
| `description` | string | Description of the `Opportunity`. |
| `address` | `Address` | Address of where the `Opportunity` takes place. |
| `is_signups_enabled` | bool | Whether or not the website will handle signups for each shift, or if it will simply only display interested users. |
| `number_of_people_needed` | int32 | Amount of people needed for the opportunity; only enbaled if signups are done on the site. |
| `tags` | string array | Categories that this opportunity falls under (ex. #foodbank, #richmondhill) |

### Address

| Field | Type |
|-------|:----:|
| street | string |
| city | string |
| province | string |
| postal_code | string |
| apt_number | string |
