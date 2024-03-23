## spotify-api

We use openapi-generator to generate the client code for the Spotify API. 

* [Spotify API](https://developer.spotify.com/documentation/web-api/reference/)
* [OpenAPI Generator](https://openapi-generator.tech/)

### Generate the client code

* go-sdk

```bash
make sync-schema
make go-sdk
```