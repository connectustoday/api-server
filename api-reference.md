# API Reference

The ConnectUS API server uses a REST API for interactions with other applications (`LIST`, `GET`, `POST`, `DELETE`). Clients interface with the API by querying a URL 
(`https://connectus.today/api/[version]/[function]`) with one of the operations specified.

## API Requirements

There are several requirements for queries against the API.
* If the query needs user authorization, an the header `x-access-token` must be included with the token.

https://www.npmjs.com/package/oauth2-server
https://oauth2-server.readthedocs.io/en/latest/misc/migrating-v2-to-v3.html
https://docs.mongodb.com/manual/tutorial/geospatial-tutorial/
https://github.com/expressjs/multer

## Error Handling

When checking for errors, check if the `error` field in the JSON object exists. If there is an error, the JSON object will not have any other fields. Otherwise, the server will return with code `200` if the query was successful.
</br>
error.code = error code
</br>
error.message = description of error

## Resource Types

### Account

### User

### Organization

### Experience

| Field | Type | Description |
|-------|:----:|-------------|
| `location` | `Address` | Location that the event had taken place. |
| `id` | string | ID of the `Experience`. The IDs are specific to the user. |
| `name` | string | The name of the `Experience`. |
| `organization` | string | `Organization` ID if the experience is being tied to an `Organization` on the site, or the email if it is using email verification. |
| `opportunity` | string | `Opportunity` ID if the experience is being tied to a specific `Organization` on the site. |
| `description` | string | A user-defined description of the experience. Another description might be provided from an `Opportunity` if it is tied to one. |
| `when` | 2 strings (array of 2) | When the `Experience` took place (ex. Sept. 2015 - Aug. 2016)
| `is_verified` | bool | Whether or not this `Experience` has been verified by the `Organization` specified. If no organization is specified, it will not show as verified. |
| `email_verify` | bool | Whether or not the experience is using email verification instead of account verification for the organization. |
| `created_at` | number | Timestamp of when the `Experience` was created. |
| `hours` | number | Number of hours gained from the `Experience`. |

### Address

| Field | Type |
|-------|:----:|
| `street` | string |
| `city` | string |
| `province` | string |
| `country` | string |
| `postal_code` | string |
| `apt_number` | string |
| `geojson` | Point |

## Endpoints

### Authentication

#### Login

`POST /v1/auth/login`

Form Data:

| Field | Type | Description |
|-------|:----:|-------------|
| `username` | string | Username of the account. |
| `password` | string | Password of the account. |

Returns (if successful):

| Field | Type | Description |
|-------|:----:|-------------|
| `token` | string | The authentication token for the account. |

Error Codes:

| Error Code | Message | HTTP Code |
|-------------------|---------------|------------------|
| 3100 | Internal server error. | 500 |
| 3101 | Invalid login. | 400 |
| 3102 | Email not verified. | 400 |
| 3103 | Bad request. | 400 |

#### Register

`POST /v1/auth/register`

Form Data (for both Users and Organizations): 

| Field | Type | Description |
|-------|:----:|-------------|
| `username` | string | Username of the account. |
| `email` | string | Email of the account. |
| `password` | string | Password of the account. |
| `type` | string | Type of the account (organization, user) |

User specific form data fields:

| Field | Type | Description |
|-------|:----:|-------------|
| `first_name` | string | First name of the user. |
| `birthday` | string | Birthday of the user. |

Organization specific form data fields:

| Field | Type | Description |
|-------|:----:|-------------|
| `is_nonprofit` | bool | Whether or not the organization is a non profit. |
| `preferred_name` | string | Preferred name of the organization. |

Returns (if successful):

HTTP Code 200 (successful).

Error Codes:

| Error Code | Message | HTTP Code |
|-------------------|---------------|------------------|
| 3200 | Invalid account type. | 500 |
| 3201 | Username already taken. | 400 |
| 3203 | Internal server error registering the account. | 500 |
| 3204 | Internal server error sending the verification email. | 500 |
| 3205 | There was a problem reading the request. | 500 |
| 3206 | Bad request. | 500 |

#### Verify Email

`POST /v1/auth/verify-email/:token`

Note: This endpoint does not need to be implemented by your client, since it is called directly when user's attempt to verify their email addresses.

---

### Accounts

#### Search for accounts 

`GET /v1/accounts/search`

Not implemented

#### Fetch Account

`GET /v1/accounts/:id`

Returns an Account object.

#### Get Account's Profile

`GET /v1/accounts/:id/profile`

Not implemented

#### Get Account's Connections

`GET /v1/accounts/:id/connections`

Not implemented

#### Get Account's Posts

`GET /v1/accounts/:id/posts`

Not implemented

#### Get User's Experiences

`GET /v1/accounts/:id/experiences`

Returns an array of `Experience`s.

Error Codes:

| Error Code | Message | HTTP Code |
|-------------------|---------------|------------------|
| 3000 | No token provided. | 401 |
| 3001 | Failed to authenticate token. | 401 |
| 3002 | Internal server error when finding account. | 500 |
| 3003 | Account not found. | 401 |
| 3004 | Email not verified. | 401 |
| 4001 | Internal server error. | 500 |
| 4002 | User not found, is this the correct account type? | 404 |

#### Get Organization's Opportunities

`GET /v1/accounts/:id/opportunities`

Not implemented

`POST /v1/accounts/:id/request_connection`

`POST /v1/accounts/:id/accept_connection`

`POST /v1/accounts/:id/block`

`POST /v1/accounts/:id/unblock`

### Personal Accounts

`GET /v1/notifications`

`POST /v1/notification/clear`

`POST /v1/notification/dismiss`

`GET /v1/settings`

`GET /v1/profile`

`GET /v1/connection-requests`

---

### Experiences

#### Get the current authenticated user's experiences

`GET /v1/experiences`

This query requires authentication.<br/>
This query only applies to Users.

Returns an array of `Experience`s.

Error Codes:

| Error Code | Message | HTTP Code |
|-------------------|---------------|------------------|
| 3000 | No token provided. | 401 |
| 3001 | Failed to authenticate token. | 401 |
| 3002 | Internal server error when finding account. | 500 |
| 3003 | Account not found. | 401 |
| 3004 | Email not verified. | 401 |
| 4000 | Incorrect account type, user account required. | 400 |
| 4001 | Internal server error. | 500 |
| 4002 | User not found, is this the correct account type? | 404 |

#### Create an experience

`POST /v1/experiences`

This query requires authentication.<br/>
This query only applies to Users.

Form Data:

| Field | Type | Description |
|-------|:----:|-------------|
| `location` | `Address` | Location that the event had taken place. |
| `name` | string | The name of the `Experience`. |
| `organization` | string | `Organization` username if the experience is being tied to an `Organization` on the site, or email if it is being tied to an email verification. |
| `opportunity` | string | `Opportunity` ID if the experience is being tied to a specific `Organization` on the site. |
| `description` | string | A user-defined description of the experience. Another description might be provided from an `Opportunity` if it is tied to one. |
| `when.begin` | string | When the `Experience` started (ex. Sept. 2015) |
| `when.end` | string | When the `Experience` ended (ex. Aug. 2016) |
| `hours` | int | Amount of hours gained from the `experience` |
| `email_verify` | bool | Whether or not the experience is being bound to email. |

Note: the `when` field is a json object storing the fields `begin` and `end`.

Error Codes:

| Error Code | Message | HTTP Code |
|-------------------|---------------|------------------|
| 3000 | No token provided. | 401 |
| 3001 | Failed to authenticate token. | 401 |
| 3002 | Internal server error when finding account. | 500 |
| 3003 | Account not found. | 401 |
| 3004 | Email not verified. | 401 |
| 4000 | Incorrect account type, user account required. | 400 |
| 4001 | Internal server error. | 500 |
| 4002 | Organization not found. | 404 |
| 4003 | Issue sending verification email. | 500 |
 
#### Replace (update) an experience

`PUT /v1/experiences/:id`

  This query requires authentication.
  This query only applies to Users.

Form Data:

| Field | Type | Description |
|-------|:----:|-------------|
| `location` | `Address` | Location that the event had taken place. |
| `name` | string | The name of the `Experience`. |
| `organization` | string | `Organization` username if the experience is being tied to an `Organization` on the site. |
| `opportunity` | string | `Opportunity` ID if the experience is being tied to a specific `Organization` on the site. |
| `description` | string | A user-defined description of the experience. Another description might be provided from an `Opportunity` if it is tied to one. |
| `when.begin` | string | When the `Experience` started (ex. Sept. 2015) |
| `when.end` | string | When the `Experience ended (ex. Aug. 2016) |
| `hours` | int | Amount of hours gained from the `experience` |

Note: the `when` field is a json object storing the fields `begin` and `end`. <br/>
Extra Note: This will set the `is_verified` field to false.

#### Delete an experience

`DELETE /v1/experiences/:id`

This query requires authentication.

Error Codes:

| Error Code | Message | HTTP Code |
|-------------------|---------------|------------------|
| 3000 | No token provided. | 401 |
| 3001 | Failed to authenticate token. | 401 |
| 3002 | Internal server error when finding account. | 500 |
| 3003 | Account not found. | 401 |
| 3004 | Email not verified. | 401 |
| 4000 | Incorrect account type, user account required. | 400 |
| 4001 | Internal server error. | 500 |
| 4002 | Experience not found with supplied ID. | 404 |

#### Get validation requests

`GET /v1/experiences/validations`

This query requires authentication.<br/>
This query only applies to Organizations.

Returns an array of experience validations:

| Field | Type | Description |
|-------|:----:|-------------|
| `user_id` | string | Username of the user that is requesting theexperience validation. |
| `experience_id` | string | The id of the `Experience`. |

Error Codes:

| Error Code | Message | HTTP Code |
|-------------------|---------------|------------------|
| 3000 | No token provided. | 401 |
| 3001 | Failed to authenticate token. | 401 |
| 3002 | Internal server error when finding account. | 500 |
| 3003 | Account not found. | 401 |
| 3004 | Email not verified. | 401 |
| 4000 | Incorrect account type, organization account required. | 400 |

#### Approve or don't approve validation

`POST /v1/experiences/validations/:user/:id`

This query requires authentication.<br/>
This query only applies to Organizations.

Form Data:

| Field | Type | Description |
|-------|:----:|-------------|
| `approve` | bool | Whether or not to approve the validation. |

Error Codes:

| Error Code | Message | HTTP Code |
|-------------------|---------------|------------------|
| 3000 | No token provided. | 401 |
| 3001 | Failed to authenticate token. | 401 |
| 3002 | Internal server error when finding account. | 500 |
| 3003 | Account not found. | 401 |
| 3004 | Email not verified. | 401 |
| 4000 | Incorrect account type, organization account required. | 400 |
| 4001 | Internal server error. | 500 |
| 4002 | Experience validation request not found. | 404 |
| 4003 | User not found. | 400 |
| 4004 | Experience not found in user object. | 400 |
 
#### Approve Validation (From email instead of account)

`POST /v1/experiences/email_approve/:token`

Note: This endpoint does not need to be implemented by your client, since it is accessed directly by the organization.
