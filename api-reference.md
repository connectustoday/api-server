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

When checking for errors, check if the error field in the JSON exists.

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
| `organization` | string | `Organization` ID if the experience is being tied to an `Organization` on the site. |
| `opportunity` | string | `Opportunity` ID if the experience is being tied to a specific `Organization` on the site. |
| `description` | string | A user-defined description of the experience. Another description might be provided from an `Opportunity` if it is tied to one. |
| `when` | 2 strings (array of 2) | When the `Experience` took place (ex. Sept. 2015 - Aug. 2016)
| `is_verified` | bool | Whether  or not this `Experience` has been verified by the `Organization` specified. If no organization is specified, it will not show as verified. |
| `created_at` | int64 | Timestamp of when the `Experience` was created. |

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

### Experiences

#### Get the current authenticated user's experiences

`GET /v1/experiences`

  This query requires authentication.
  This query only applies to Users.

Returns an array of `Experience`s.

#### Create an experience

`POST /v1/experiences`

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

Note: the `when` field is a json object storing the fields `begin` and `end`.

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

  Note: the `when` field is a json object storing the fields `begin` and `end`.
  Extra Note: This will set the `is_verified` field to false.

#### Delete an experience

`DELETE /v1/experiences/:id`

This query requires authentication.

#### Get validation requests

`GET /v1/experiences/validations`

  This query requires authentication.
  This query only applies to Organizations.

Returns an array of experience validations:

| Field | Type | Description |
|-------|:----:|-------------|
| `user_id` | string | Username of the user that is requesting theexperience validation. |
| `experience_id` | string | The id of the `Experience`. |

#### Approve or don't approve validation

`GET /v1/experiences/validations/:user/:id`

  This query requires authentication.
  This query only applies to Organizations.

Form Data:

| Field | Type | Description |
|-------|:----:|-------------|
| `approve` | bool | Whether or not to approve the validation. |
