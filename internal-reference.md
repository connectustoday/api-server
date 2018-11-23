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
| `avatar` | string | URL to the profile picture of the account. |
| `header` | string | URL to the header image of the account. |
| `created_at` | int64 | Timestamp for when the account was created. |
| `pending_connections` | string array | Pending connection requests from other `Account`s. Array of IDs. | 
| `requested_connections` | string array | Pending connection requests to other `Account`s. Array of IDs. |
| `posts` | string array | List of `Post` IDs that the `Account` has posted. |
| `liked` | [string, int64]  array | List of `Post` IDs that the `Account` has liked, with the timestamp of when it was liked. |
| `shared` | [string, int64] array | List of `Post` IDs that the `Account` has shared, with the timestamp of when it was shared. |
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
| `gender` | string | Male, female, or other. |
| `personal_info` | `UserProfile` | The personal information of an user. |
| `experiences` | `Experience` array | List of `Experience`s or "roles" that the user has. |

### UserSettings (extends `AccountSettings`)
Schema representing settings specific to `User`s. This is stored within the `Account` object.

| Field | Type | Description |
|-------|:----:|-------------|
| `is_full_name_visible` | bool | Whether or not the user allows others to see its full name. |
| `blocked_users` | string array | Array of `Account` IDs for blocked users. |

### UserProfile
`UserProfile` represents the personal information of an user. This is stored within the `User` 
object.

| Field | Type | Description |
|-------|:----:|-------------|
| `schema_version` | int32 | The schema version. |
| `interests` | string array | Tags that the user is interested in. |
| `biography` | string | Biography of the user. |
| `education` | undecided | undecided |
| `quote` | string | User defined quote. |
| `current_residence` | string | City that the user lives in. |
| `certifications` | undecided | undecided |

### Organization (extends `Account`)
Schema representing registered organizations on the site. It inherits fields from `Account`.

| Field | Type | Description |
|-------|:----:|-------------|
| `preferred_name` | string | Preferred name of the organization that shows up on their profile. |
| `is_verified` | bool | If the organization has been verified to exist, and the account belongs to the real organization. |
| `opportunities` | string array | List of `Opportunity` IDs that the organization has created. |
| `org_info` | `OrganizationProfile` | The organization's public information. |

### OrganizationProfile
`OrganizationProfile` represents organization specific information that is shown on their profile. 
This is stored within the `Organization` object.

| Field | Type | Description |
|-------|:----:|-------------|
| `schema_version` | int32 | The schema version. |
| `mission` | string | The organization's stated mission. |
| `quote` | string | The organization's specified quote. |
| `address` | `Address` | The organization's headquarters location. |
| `affiliated_orgs` | string array | IDs of other organizations this organization is affiliated with. |
| `interests` | string array | Tags that the organization is interested in. |
| `experience_validations` | array of 2 string pairs (object) | List of User IDs and personal Experience IDs. Represent users seeking validation for their Experience from the organization. |

### OrganizationSettings (extends `AccountSettings`)
Settings specific to `Organization`s. This is stored within the `Account` object.

| Field | Type | Description |
|-------|:----:|-------------|
| `is_nonprofit` | bool | Whether or not the organization is non-profit. |


### Post
Represents posts and replies that `Account`s create on the website.

| Field | Type | Description |
|-------|:----:|-------------|
| `schema_version` | int32 | The schema version. |
| `id` | string | The unique ID of the post. |
| `account` | string | ID of the `Account` that posted it. |
| `content` | string | Text content of the post. |
| `created_at` | int64 | Timestamp of when the post was created. |
| `reply_to` | string | ID of the `Post` that was replied to, if applicable. |
| `multimedia` | `Attachment` | Links to other media (images/videos/articles) that are attached to the post. Displayed separately.|
| `tags` | string array | Tags that the post is categorized under. |
| `likes_count` | int32 | Number of likes for the post. | 
| `comments_count` | int32 | Number of comments for the post. |
| `shares_count` | int32 | Number of shares for the post. |
| `likes` | [string, int64] array | Array of `Account` IDs and timestamps representing accounts that have liked the post, and when. |
| `comments` | string array | Array of `Post` IDs that `Account`s have posted on the post. |
| `shares` | [string, int64] array | Array of `Account` IDs representing accounts that have shared the post, and the timestamp of when. |
| `visibility` | string | Visibility of the post on the account's profile. |

### Attachment

| Field | Type | Description |
|-------|:----:|-------------|
| `schema_version` | int32 | The schema version. |
| `type` | string | The type of the attachment (image, news post, video, etc.) | 
| `url` | string | URL of the attachment. |
| `description` | string | Description of the attachment, if applicable. |

### Opportunity
An opportunity object represents an opportunity that an organization has for people to sign up 
for. This can be for an event (Canada Day), or simply a shift for a job (food bank). Organizations 
can choose whether or not they want to use the built-in mechanism for signing up users. They will 
have a choice to allow users to either sign up for a shift directly on the site, or to let the 
organization know which users have expressed their interest for the position. 

| Field | Type | Description |
|-------|:----:|-------------|
| `schema_version` | int32 | The schema version. |
| `organization` | string | ID of the `Organization` that created it. |
| `id` | string | ID of the `Opportunity`. |
| `name` | string | Name of the `Opportunity`. |
| `description` | string | Description of the `Opportunity`. |
| `address` | `Address` | Address of where the `Opportunity` takes place. |
| `is_signups_enabled` | bool | Whether or not the website will handle signups for each shift, or if it will simply only display interested users. |
| `number_of_people_needed` | int32 | Amount of people needed for the opportunity; only enabled if signups are done on the site. |
| `tags` | string array | Categories that this opportunity falls under (ex. #foodbank, #richmondhill) |
| `interested_users` | string array | The IDs of the users that have confirmed their interest for the opportunity. If signups are enabled, it will also show their interest for whatever shift they chose. This is only accessible to organization. |
| `shift_times` | array of 2 strings pairs (object) | Shift times of the opportunity, if the signups for it are done on the website. |
| `method_of_contact` | string | If signups are not enabled, a method of contact (email, messaging) is provided to allow users to contact the organization for more information. |
| `created_at` | int64 | Timestamp of when the `Opportunity` was created. |

### Experience
An experience represents an experience that a user has had. This can include jobs (a shift at 
McDonald’s), volunteer shifts, and a role that a person has in an organization (club president for 
DECA at Richmond Hill High). It is optional to get it verified by the organization. Experiences 
can be tied to existing `Opportunity` objects, to show that the user has participated in that 
specific opportunity. `Experience`s are stored under the corresponding user’s schema.

| Field | Type | Description |
|-------|:----:|-------------|
| `schema_version` | int32 | The schema version. |
| `location` | `Address` | Location that the event had taken place. |
| `id` | string | ID of the `Experience`. The IDs are specific to the user. |
| `organization` | string | `Organization` ID if the experience is being tied to an `Organization` on the site. |
| `opportunity` | string | `Opportunity` ID if the experience is being tied to a specific `Organization` on the site. |
| `description` | string | A user-defined description of the experience. Another description might be provided from an `Opportunity` if it is tied to one. |
| `when` | 2 strings (tuple) | When the `Experience` took place (ex. Sept. 2015 - Aug. 2016)
| `is_verified` | bool | Whether or not this `Experience` has been verified by the `Organization` specified. If no organization is specified, it will not show as verified. |
| `email_bound` | bool | Whether or not this experience is being verified by email (instead of account). |
| `created_at` | int64 | Timestamp of when the `Experience` was created. |

### Notification

IDs are specific to a user.

| Field | Type | Description |
|-------|:----:|-------------|
| `id` | int64 | ID of the notification. Specific to the user. |
| `created_at` | int64 | Timestamp of when the notification was created. |
| `type` | string | The notification type.
| `content` | string | The content of the notification. |
| `account` | string | The account that caused the notification, if applicable. |

### Address

May be replaced by MongoDB Geospatial objects.

| Field | Type |
|-------|:----:|
| `street` | string |
| `city` | string |
| `province` | string |
| `country` | string |
| `postal_code` | string |
| `apt_number` | string |
| `geojson` | GeoJSON point |

___

### Useful links
* [https://docs.mongodb.com/manual/indexes/]
