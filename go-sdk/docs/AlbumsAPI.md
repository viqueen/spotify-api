# \AlbumsAPI

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUsersSavedAlbums**](AlbumsAPI.md#CheckUsersSavedAlbums) | **Get** /me/albums/contains | Check User&#39;s Saved Albums 
[**GetAnAlbum**](AlbumsAPI.md#GetAnAlbum) | **Get** /albums/{id} | Get Album 
[**GetAnAlbumsTracks**](AlbumsAPI.md#GetAnAlbumsTracks) | **Get** /albums/{id}/tracks | Get Album Tracks 
[**GetAnArtistsAlbums**](AlbumsAPI.md#GetAnArtistsAlbums) | **Get** /artists/{id}/albums | Get Artist&#39;s Albums 
[**GetMultipleAlbums**](AlbumsAPI.md#GetMultipleAlbums) | **Get** /albums | Get Several Albums 
[**GetNewReleases**](AlbumsAPI.md#GetNewReleases) | **Get** /browse/new-releases | Get New Releases 
[**GetUsersSavedAlbums**](AlbumsAPI.md#GetUsersSavedAlbums) | **Get** /me/albums | Get User&#39;s Saved Albums 
[**RemoveAlbumsUser**](AlbumsAPI.md#RemoveAlbumsUser) | **Delete** /me/albums | Remove Users&#39; Saved Albums 
[**SaveAlbumsUser**](AlbumsAPI.md#SaveAlbumsUser) | **Put** /me/albums | Save Albums for Current User 



## CheckUsersSavedAlbums

> []bool CheckUsersSavedAlbums(ctx).Ids(ids).Execute()

Check User's Saved Albums 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/viqueen/spotify-api"
)

func main() {
	ids := "382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlbumsAPI.CheckUsersSavedAlbums(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumsAPI.CheckUsersSavedAlbums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsersSavedAlbums`: []bool
	fmt.Fprintf(os.Stdout, "Response from `AlbumsAPI.CheckUsersSavedAlbums`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsersSavedAlbumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

### Return type

**[]bool**

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnAlbum

> AlbumObject GetAnAlbum(ctx, id).Market(market).Execute()

Get Album 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/viqueen/spotify-api"
)

func main() {
	id := "4aawyAB9vmqN3uQ7FjRGTy" // string | 
	market := "ES" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlbumsAPI.GetAnAlbum(context.Background(), id).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumsAPI.GetAnAlbum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnAlbum`: AlbumObject
	fmt.Fprintf(os.Stdout, "Response from `AlbumsAPI.GetAnAlbum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnAlbumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 

### Return type

[**AlbumObject**](AlbumObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnAlbumsTracks

> PagingSimplifiedTrackObject GetAnAlbumsTracks(ctx, id).Market(market).Limit(limit).Offset(offset).Execute()

Get Album Tracks 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/viqueen/spotify-api"
)

func main() {
	id := "4aawyAB9vmqN3uQ7FjRGTy" // string | 
	market := "ES" // string |  (optional)
	limit := int32(10) // int32 |  (optional) (default to 20)
	offset := int32(5) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlbumsAPI.GetAnAlbumsTracks(context.Background(), id).Market(market).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumsAPI.GetAnAlbumsTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnAlbumsTracks`: PagingSimplifiedTrackObject
	fmt.Fprintf(os.Stdout, "Response from `AlbumsAPI.GetAnAlbumsTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnAlbumsTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSimplifiedTrackObject**](PagingSimplifiedTrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnArtistsAlbums

> PagingArtistDiscographyAlbumObject GetAnArtistsAlbums(ctx, id).IncludeGroups(includeGroups).Market(market).Limit(limit).Offset(offset).Execute()

Get Artist's Albums 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/viqueen/spotify-api"
)

func main() {
	id := "0TnOYISbd1XYRBk9myaseg" // string | 
	includeGroups := "single,appears_on" // string |  (optional)
	market := "ES" // string |  (optional)
	limit := int32(10) // int32 |  (optional) (default to 20)
	offset := int32(5) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlbumsAPI.GetAnArtistsAlbums(context.Background(), id).IncludeGroups(includeGroups).Market(market).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumsAPI.GetAnArtistsAlbums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnArtistsAlbums`: PagingArtistDiscographyAlbumObject
	fmt.Fprintf(os.Stdout, "Response from `AlbumsAPI.GetAnArtistsAlbums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnArtistsAlbumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeGroups** | **string** |  | 
 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingArtistDiscographyAlbumObject**](PagingArtistDiscographyAlbumObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleAlbums

> GetMultipleAlbums200Response GetMultipleAlbums(ctx).Ids(ids).Market(market).Execute()

Get Several Albums 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/viqueen/spotify-api"
)

func main() {
	ids := "382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc" // string | 
	market := "ES" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlbumsAPI.GetMultipleAlbums(context.Background()).Ids(ids).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumsAPI.GetMultipleAlbums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipleAlbums`: GetMultipleAlbums200Response
	fmt.Fprintf(os.Stdout, "Response from `AlbumsAPI.GetMultipleAlbums`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleAlbumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **market** | **string** |  | 

### Return type

[**GetMultipleAlbums200Response**](GetMultipleAlbums200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewReleases

> GetNewReleases200Response GetNewReleases(ctx).Limit(limit).Offset(offset).Execute()

Get New Releases 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/viqueen/spotify-api"
)

func main() {
	limit := int32(10) // int32 |  (optional) (default to 20)
	offset := int32(5) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlbumsAPI.GetNewReleases(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumsAPI.GetNewReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewReleases`: GetNewReleases200Response
	fmt.Fprintf(os.Stdout, "Response from `AlbumsAPI.GetNewReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNewReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**GetNewReleases200Response**](GetNewReleases200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersSavedAlbums

> PagingSavedAlbumObject GetUsersSavedAlbums(ctx).Limit(limit).Offset(offset).Market(market).Execute()

Get User's Saved Albums 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/viqueen/spotify-api"
)

func main() {
	limit := int32(10) // int32 |  (optional) (default to 20)
	offset := int32(5) // int32 |  (optional) (default to 0)
	market := "ES" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlbumsAPI.GetUsersSavedAlbums(context.Background()).Limit(limit).Offset(offset).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumsAPI.GetUsersSavedAlbums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersSavedAlbums`: PagingSavedAlbumObject
	fmt.Fprintf(os.Stdout, "Response from `AlbumsAPI.GetUsersSavedAlbums`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSavedAlbumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **market** | **string** |  | 

### Return type

[**PagingSavedAlbumObject**](PagingSavedAlbumObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAlbumsUser

> RemoveAlbumsUser(ctx).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()

Remove Users' Saved Albums 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/viqueen/spotify-api"
)

func main() {
	ids := "382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc" // string | 
	saveAlbumsUserRequest := *openapiclient.NewSaveAlbumsUserRequest() // SaveAlbumsUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlbumsAPI.RemoveAlbumsUser(context.Background()).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumsAPI.RemoveAlbumsUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAlbumsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **saveAlbumsUserRequest** | [**SaveAlbumsUserRequest**](SaveAlbumsUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAlbumsUser

> SaveAlbumsUser(ctx).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()

Save Albums for Current User 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/viqueen/spotify-api"
)

func main() {
	ids := "382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc" // string | 
	saveAlbumsUserRequest := *openapiclient.NewSaveAlbumsUserRequest() // SaveAlbumsUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlbumsAPI.SaveAlbumsUser(context.Background()).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumsAPI.SaveAlbumsUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAlbumsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **saveAlbumsUserRequest** | [**SaveAlbumsUserRequest**](SaveAlbumsUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

