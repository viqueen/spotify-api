# AlbumObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlbumType** | **string** | The type of the album.  | 
**TotalTracks** | **int32** | The number of tracks in the album. | 
**AvailableMarkets** | **[]string** | The markets in which the album is available: [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _**NOTE**: an album is considered available in a market when at least 1 of its tracks is available in that market._  | 
**ExternalUrls** | [**ExternalUrlObject**](ExternalUrlObject.md) | Known external URLs for this album.  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the album.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the album.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the album in various sizes, widest first.  | 
**Name** | **string** | The name of the album. In case of an album takedown, the value may be an empty string.  | 
**ReleaseDate** | **string** | The date the album was first released.  | 
**ReleaseDatePrecision** | **string** | The precision with which &#x60;release_date&#x60; value is known.  | 
**Restrictions** | Pointer to [**AlbumRestrictionObject**](AlbumRestrictionObject.md) | Included in the response when a content restriction is applied.  | [optional] 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the album.  | 
**Artists** | [**[]SimplifiedArtistObject**](SimplifiedArtistObject.md) | The artists of the album. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist.  | 
**Tracks** | [**PagingSimplifiedTrackObject**](PagingSimplifiedTrackObject.md) | The tracks of the album.  | 
**Copyrights** | [**[]CopyrightObject**](CopyrightObject.md) | The copyright statements of the album.  | 
**ExternalIds** | [**ExternalIdObject**](ExternalIdObject.md) | Known external IDs for the album.  | 
**Genres** | **[]string** | A list of the genres the album is associated with. If not yet classified, the array is empty.  | 
**Label** | **string** | The label associated with the album.  | 
**Popularity** | **int32** | The popularity of the album. The value will be between 0 and 100, with 100 being the most popular.  | 

## Methods

### NewAlbumObject

`func NewAlbumObject(albumType string, totalTracks int32, availableMarkets []string, externalUrls ExternalUrlObject, href string, id string, images []ImageObject, name string, releaseDate string, releaseDatePrecision string, type_ string, uri string, artists []SimplifiedArtistObject, tracks PagingSimplifiedTrackObject, copyrights []CopyrightObject, externalIds ExternalIdObject, genres []string, label string, popularity int32, ) *AlbumObject`

NewAlbumObject instantiates a new AlbumObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumObjectWithDefaults

`func NewAlbumObjectWithDefaults() *AlbumObject`

NewAlbumObjectWithDefaults instantiates a new AlbumObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbumType

`func (o *AlbumObject) GetAlbumType() string`

GetAlbumType returns the AlbumType field if non-nil, zero value otherwise.

### GetAlbumTypeOk

`func (o *AlbumObject) GetAlbumTypeOk() (*string, bool)`

GetAlbumTypeOk returns a tuple with the AlbumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumType

`func (o *AlbumObject) SetAlbumType(v string)`

SetAlbumType sets AlbumType field to given value.


### GetTotalTracks

`func (o *AlbumObject) GetTotalTracks() int32`

GetTotalTracks returns the TotalTracks field if non-nil, zero value otherwise.

### GetTotalTracksOk

`func (o *AlbumObject) GetTotalTracksOk() (*int32, bool)`

GetTotalTracksOk returns a tuple with the TotalTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTracks

`func (o *AlbumObject) SetTotalTracks(v int32)`

SetTotalTracks sets TotalTracks field to given value.


### GetAvailableMarkets

`func (o *AlbumObject) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *AlbumObject) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *AlbumObject) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.


### GetExternalUrls

`func (o *AlbumObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *AlbumObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *AlbumObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *AlbumObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AlbumObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AlbumObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *AlbumObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlbumObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlbumObject) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *AlbumObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AlbumObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AlbumObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetName

`func (o *AlbumObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlbumObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlbumObject) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *AlbumObject) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *AlbumObject) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *AlbumObject) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *AlbumObject) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *AlbumObject) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *AlbumObject) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetRestrictions

`func (o *AlbumObject) GetRestrictions() AlbumRestrictionObject`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *AlbumObject) GetRestrictionsOk() (*AlbumRestrictionObject, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *AlbumObject) SetRestrictions(v AlbumRestrictionObject)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *AlbumObject) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetType

`func (o *AlbumObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlbumObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlbumObject) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *AlbumObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AlbumObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AlbumObject) SetUri(v string)`

SetUri sets Uri field to given value.


### GetArtists

`func (o *AlbumObject) GetArtists() []SimplifiedArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *AlbumObject) GetArtistsOk() (*[]SimplifiedArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *AlbumObject) SetArtists(v []SimplifiedArtistObject)`

SetArtists sets Artists field to given value.


### GetTracks

`func (o *AlbumObject) GetTracks() PagingSimplifiedTrackObject`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *AlbumObject) GetTracksOk() (*PagingSimplifiedTrackObject, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *AlbumObject) SetTracks(v PagingSimplifiedTrackObject)`

SetTracks sets Tracks field to given value.


### GetCopyrights

`func (o *AlbumObject) GetCopyrights() []CopyrightObject`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *AlbumObject) GetCopyrightsOk() (*[]CopyrightObject, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *AlbumObject) SetCopyrights(v []CopyrightObject)`

SetCopyrights sets Copyrights field to given value.


### GetExternalIds

`func (o *AlbumObject) GetExternalIds() ExternalIdObject`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *AlbumObject) GetExternalIdsOk() (*ExternalIdObject, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *AlbumObject) SetExternalIds(v ExternalIdObject)`

SetExternalIds sets ExternalIds field to given value.


### GetGenres

`func (o *AlbumObject) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *AlbumObject) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *AlbumObject) SetGenres(v []string)`

SetGenres sets Genres field to given value.


### GetLabel

`func (o *AlbumObject) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AlbumObject) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AlbumObject) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPopularity

`func (o *AlbumObject) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *AlbumObject) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *AlbumObject) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


