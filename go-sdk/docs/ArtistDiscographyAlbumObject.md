# ArtistDiscographyAlbumObject

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
**AlbumGroup** | **string** | This field describes the relationship between the artist and the album.  | 

## Methods

### NewArtistDiscographyAlbumObject

`func NewArtistDiscographyAlbumObject(albumType string, totalTracks int32, availableMarkets []string, externalUrls ExternalUrlObject, href string, id string, images []ImageObject, name string, releaseDate string, releaseDatePrecision string, type_ string, uri string, artists []SimplifiedArtistObject, albumGroup string, ) *ArtistDiscographyAlbumObject`

NewArtistDiscographyAlbumObject instantiates a new ArtistDiscographyAlbumObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtistDiscographyAlbumObjectWithDefaults

`func NewArtistDiscographyAlbumObjectWithDefaults() *ArtistDiscographyAlbumObject`

NewArtistDiscographyAlbumObjectWithDefaults instantiates a new ArtistDiscographyAlbumObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbumType

`func (o *ArtistDiscographyAlbumObject) GetAlbumType() string`

GetAlbumType returns the AlbumType field if non-nil, zero value otherwise.

### GetAlbumTypeOk

`func (o *ArtistDiscographyAlbumObject) GetAlbumTypeOk() (*string, bool)`

GetAlbumTypeOk returns a tuple with the AlbumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumType

`func (o *ArtistDiscographyAlbumObject) SetAlbumType(v string)`

SetAlbumType sets AlbumType field to given value.


### GetTotalTracks

`func (o *ArtistDiscographyAlbumObject) GetTotalTracks() int32`

GetTotalTracks returns the TotalTracks field if non-nil, zero value otherwise.

### GetTotalTracksOk

`func (o *ArtistDiscographyAlbumObject) GetTotalTracksOk() (*int32, bool)`

GetTotalTracksOk returns a tuple with the TotalTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTracks

`func (o *ArtistDiscographyAlbumObject) SetTotalTracks(v int32)`

SetTotalTracks sets TotalTracks field to given value.


### GetAvailableMarkets

`func (o *ArtistDiscographyAlbumObject) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *ArtistDiscographyAlbumObject) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *ArtistDiscographyAlbumObject) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.


### GetExternalUrls

`func (o *ArtistDiscographyAlbumObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *ArtistDiscographyAlbumObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *ArtistDiscographyAlbumObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *ArtistDiscographyAlbumObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArtistDiscographyAlbumObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArtistDiscographyAlbumObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *ArtistDiscographyAlbumObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtistDiscographyAlbumObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtistDiscographyAlbumObject) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *ArtistDiscographyAlbumObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ArtistDiscographyAlbumObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ArtistDiscographyAlbumObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetName

`func (o *ArtistDiscographyAlbumObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtistDiscographyAlbumObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtistDiscographyAlbumObject) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *ArtistDiscographyAlbumObject) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ArtistDiscographyAlbumObject) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ArtistDiscographyAlbumObject) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *ArtistDiscographyAlbumObject) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *ArtistDiscographyAlbumObject) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *ArtistDiscographyAlbumObject) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetRestrictions

`func (o *ArtistDiscographyAlbumObject) GetRestrictions() AlbumRestrictionObject`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *ArtistDiscographyAlbumObject) GetRestrictionsOk() (*AlbumRestrictionObject, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *ArtistDiscographyAlbumObject) SetRestrictions(v AlbumRestrictionObject)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *ArtistDiscographyAlbumObject) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetType

`func (o *ArtistDiscographyAlbumObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtistDiscographyAlbumObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtistDiscographyAlbumObject) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *ArtistDiscographyAlbumObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ArtistDiscographyAlbumObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ArtistDiscographyAlbumObject) SetUri(v string)`

SetUri sets Uri field to given value.


### GetArtists

`func (o *ArtistDiscographyAlbumObject) GetArtists() []SimplifiedArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *ArtistDiscographyAlbumObject) GetArtistsOk() (*[]SimplifiedArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *ArtistDiscographyAlbumObject) SetArtists(v []SimplifiedArtistObject)`

SetArtists sets Artists field to given value.


### GetAlbumGroup

`func (o *ArtistDiscographyAlbumObject) GetAlbumGroup() string`

GetAlbumGroup returns the AlbumGroup field if non-nil, zero value otherwise.

### GetAlbumGroupOk

`func (o *ArtistDiscographyAlbumObject) GetAlbumGroupOk() (*string, bool)`

GetAlbumGroupOk returns a tuple with the AlbumGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumGroup

`func (o *ArtistDiscographyAlbumObject) SetAlbumGroup(v string)`

SetAlbumGroup sets AlbumGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

