# API Reference

The ConnectUS API server uses a REST API for interactions with other applications (`LIST`, `GET`, `POST`, `DELETE`). Clients interface with the API by querying a URL 
(`https://connectus.today/api/[version]/[function]`) with one of the operations specified.

## API Requirements

There are several requirements for queries against the API.
* An API key is required for each query (using the `api_key` parameter).
* If the query needs user authorization, an `oauth2_token` parameter must be included.


