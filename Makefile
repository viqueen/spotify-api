sync-schema:
	curl https://developer.spotify.com/reference/web-api/open-api-schema.yaml -o schema.yaml

go-sdk:
	docker run --rm \
		--volume "$(PWD):/local" \
		--volume "$(PWD)/../spotify-api-go-sdk:/local/go-sdk" \
		openapitools/openapi-generator-cli generate \
		--input-spec /local/schema.yaml \
		--skip-validate-spec \
		--generator-name go \
		--git-user-id viqueen \
		--git-repo-id spotify-api-go-sdk \
		--output /local/go-sdk