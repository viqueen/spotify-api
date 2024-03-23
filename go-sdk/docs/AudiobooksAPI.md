# \AudiobooksAPI

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUsersSavedAudiobooks**](AudiobooksAPI.md#CheckUsersSavedAudiobooks) | **Get** /me/audiobooks/contains | Check User&#39;s Saved Audiobooks 
[**GetAnAudiobook**](AudiobooksAPI.md#GetAnAudiobook) | **Get** /audiobooks/{id} | Get an Audiobook 
[**GetAudiobookChapters**](AudiobooksAPI.md#GetAudiobookChapters) | **Get** /audiobooks/{id}/chapters | Get Audiobook Chapters 
[**GetMultipleAudiobooks**](AudiobooksAPI.md#GetMultipleAudiobooks) | **Get** /audiobooks | Get Several Audiobooks 
[**GetUsersSavedAudiobooks**](AudiobooksAPI.md#GetUsersSavedAudiobooks) | **Get** /me/audiobooks | Get User&#39;s Saved Audiobooks 
[**RemoveAudiobooksUser**](AudiobooksAPI.md#RemoveAudiobooksUser) | **Delete** /me/audiobooks | Remove User&#39;s Saved Audiobooks 
[**SaveAudiobooksUser**](AudiobooksAPI.md#SaveAudiobooksUser) | **Put** /me/audiobooks | Save Audiobooks for Current User 



## CheckUsersSavedAudiobooks

> []bool CheckUsersSavedAudiobooks(ctx).Ids(ids).Execute()

Check User's Saved Audiobooks 



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
	ids := "18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ,7iHfbu1YPACw6oZPAFJtqe" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudiobooksAPI.CheckUsersSavedAudiobooks(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiobooksAPI.CheckUsersSavedAudiobooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsersSavedAudiobooks`: []bool
	fmt.Fprintf(os.Stdout, "Response from `AudiobooksAPI.CheckUsersSavedAudiobooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsersSavedAudiobooksRequest struct via the builder pattern


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


## GetAnAudiobook

> AudiobookObject GetAnAudiobook(ctx, id).Market(market).Execute()

Get an Audiobook 



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
	id := "7iHfbu1YPACw6oZPAFJtqe" // string | 
	market := "ES" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudiobooksAPI.GetAnAudiobook(context.Background(), id).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiobooksAPI.GetAnAudiobook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnAudiobook`: AudiobookObject
	fmt.Fprintf(os.Stdout, "Response from `AudiobooksAPI.GetAnAudiobook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnAudiobookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 

### Return type

[**AudiobookObject**](AudiobookObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudiobookChapters

> PagingSimplifiedChapterObject GetAudiobookChapters(ctx, id).Market(market).Limit(limit).Offset(offset).Execute()

Get Audiobook Chapters 



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
	id := "7iHfbu1YPACw6oZPAFJtqe" // string | 
	market := "ES" // string |  (optional)
	limit := int32(10) // int32 |  (optional) (default to 20)
	offset := int32(5) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudiobooksAPI.GetAudiobookChapters(context.Background(), id).Market(market).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiobooksAPI.GetAudiobookChapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudiobookChapters`: PagingSimplifiedChapterObject
	fmt.Fprintf(os.Stdout, "Response from `AudiobooksAPI.GetAudiobookChapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudiobookChaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSimplifiedChapterObject**](PagingSimplifiedChapterObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleAudiobooks

> GetMultipleAudiobooks200Response GetMultipleAudiobooks(ctx).Ids(ids).Market(market).Execute()

Get Several Audiobooks 



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
	ids := "18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ,7iHfbu1YPACw6oZPAFJtqe" // string | 
	market := "ES" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudiobooksAPI.GetMultipleAudiobooks(context.Background()).Ids(ids).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiobooksAPI.GetMultipleAudiobooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipleAudiobooks`: GetMultipleAudiobooks200Response
	fmt.Fprintf(os.Stdout, "Response from `AudiobooksAPI.GetMultipleAudiobooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleAudiobooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **market** | **string** |  | 

### Return type

[**GetMultipleAudiobooks200Response**](GetMultipleAudiobooks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersSavedAudiobooks

> PagingSimplifiedAudiobookObject GetUsersSavedAudiobooks(ctx).Limit(limit).Offset(offset).Execute()

Get User's Saved Audiobooks 



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
	resp, r, err := apiClient.AudiobooksAPI.GetUsersSavedAudiobooks(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiobooksAPI.GetUsersSavedAudiobooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersSavedAudiobooks`: PagingSimplifiedAudiobookObject
	fmt.Fprintf(os.Stdout, "Response from `AudiobooksAPI.GetUsersSavedAudiobooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSavedAudiobooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSimplifiedAudiobookObject**](PagingSimplifiedAudiobookObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAudiobooksUser

> RemoveAudiobooksUser(ctx).Ids(ids).Execute()

Remove User's Saved Audiobooks 



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
	ids := "18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ,7iHfbu1YPACw6oZPAFJtqe" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AudiobooksAPI.RemoveAudiobooksUser(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiobooksAPI.RemoveAudiobooksUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAudiobooksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAudiobooksUser

> SaveAudiobooksUser(ctx).Ids(ids).Execute()

Save Audiobooks for Current User 



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
	ids := "18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ,7iHfbu1YPACw6oZPAFJtqe" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AudiobooksAPI.SaveAudiobooksUser(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudiobooksAPI.SaveAudiobooksUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAudiobooksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

