/*
Spotify Web API

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TrackObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackObject{}

// TrackObject struct for TrackObject
type TrackObject struct {
	// The album on which the track appears. The album object includes a link in `href` to full information about the album. 
	Album *SimplifiedAlbumObject `json:"album,omitempty"`
	// The artists who performed the track. Each artist object includes a link in `href` to more detailed information about the artist. 
	Artists []ArtistObject `json:"artists,omitempty"`
	// A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. 
	AvailableMarkets []string `json:"available_markets,omitempty"`
	// The disc number (usually `1` unless the album consists of more than one disc). 
	DiscNumber *int32 `json:"disc_number,omitempty"`
	// The track length in milliseconds. 
	DurationMs *int32 `json:"duration_ms,omitempty"`
	// Whether or not the track has explicit lyrics ( `true` = yes it does; `false` = no it does not OR unknown). 
	Explicit *bool `json:"explicit,omitempty"`
	// Known external IDs for the track. 
	ExternalIds *ExternalIdObject `json:"external_ids,omitempty"`
	// Known external URLs for this track. 
	ExternalUrls *ExternalUrlObject `json:"external_urls,omitempty"`
	// A link to the Web API endpoint providing full details of the track. 
	Href *string `json:"href,omitempty"`
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track. 
	Id *string `json:"id,omitempty"`
	// Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied. If `true`, the track is playable in the given market. Otherwise `false`. 
	IsPlayable *bool `json:"is_playable,omitempty"`
	// Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied, and the requested track has been replaced with different track. The track in the `linked_from` object contains information about the originally requested track. 
	LinkedFrom map[string]interface{} `json:"linked_from,omitempty"`
	// Included in the response when a content restriction is applied. 
	Restrictions *TrackRestrictionObject `json:"restrictions,omitempty"`
	// The name of the track. 
	Name *string `json:"name,omitempty"`
	// The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.<br/>The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.<br/>Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. _**Note**: the popularity value may lag actual popularity by a few days: the value is not updated in real time._ 
	Popularity *int32 `json:"popularity,omitempty"`
	// A link to a 30 second preview (MP3 format) of the track. Can be `null` 
	PreviewUrl NullableString `json:"preview_url,omitempty"`
	// The number of the track. If an album has several discs, the track number is the number on the specified disc. 
	TrackNumber *int32 `json:"track_number,omitempty"`
	// The object type: \"track\". 
	Type *string `json:"type,omitempty"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track. 
	Uri *string `json:"uri,omitempty"`
	// Whether or not the track is from a local file. 
	IsLocal *bool `json:"is_local,omitempty"`
}

// NewTrackObject instantiates a new TrackObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackObject() *TrackObject {
	this := TrackObject{}
	return &this
}

// NewTrackObjectWithDefaults instantiates a new TrackObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackObjectWithDefaults() *TrackObject {
	this := TrackObject{}
	return &this
}

// GetAlbum returns the Album field value if set, zero value otherwise.
func (o *TrackObject) GetAlbum() SimplifiedAlbumObject {
	if o == nil || IsNil(o.Album) {
		var ret SimplifiedAlbumObject
		return ret
	}
	return *o.Album
}

// GetAlbumOk returns a tuple with the Album field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetAlbumOk() (*SimplifiedAlbumObject, bool) {
	if o == nil || IsNil(o.Album) {
		return nil, false
	}
	return o.Album, true
}

// HasAlbum returns a boolean if a field has been set.
func (o *TrackObject) HasAlbum() bool {
	if o != nil && !IsNil(o.Album) {
		return true
	}

	return false
}

// SetAlbum gets a reference to the given SimplifiedAlbumObject and assigns it to the Album field.
func (o *TrackObject) SetAlbum(v SimplifiedAlbumObject) {
	o.Album = &v
}

// GetArtists returns the Artists field value if set, zero value otherwise.
func (o *TrackObject) GetArtists() []ArtistObject {
	if o == nil || IsNil(o.Artists) {
		var ret []ArtistObject
		return ret
	}
	return o.Artists
}

// GetArtistsOk returns a tuple with the Artists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetArtistsOk() ([]ArtistObject, bool) {
	if o == nil || IsNil(o.Artists) {
		return nil, false
	}
	return o.Artists, true
}

// HasArtists returns a boolean if a field has been set.
func (o *TrackObject) HasArtists() bool {
	if o != nil && !IsNil(o.Artists) {
		return true
	}

	return false
}

// SetArtists gets a reference to the given []ArtistObject and assigns it to the Artists field.
func (o *TrackObject) SetArtists(v []ArtistObject) {
	o.Artists = v
}

// GetAvailableMarkets returns the AvailableMarkets field value if set, zero value otherwise.
func (o *TrackObject) GetAvailableMarkets() []string {
	if o == nil || IsNil(o.AvailableMarkets) {
		var ret []string
		return ret
	}
	return o.AvailableMarkets
}

// GetAvailableMarketsOk returns a tuple with the AvailableMarkets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetAvailableMarketsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableMarkets) {
		return nil, false
	}
	return o.AvailableMarkets, true
}

// HasAvailableMarkets returns a boolean if a field has been set.
func (o *TrackObject) HasAvailableMarkets() bool {
	if o != nil && !IsNil(o.AvailableMarkets) {
		return true
	}

	return false
}

// SetAvailableMarkets gets a reference to the given []string and assigns it to the AvailableMarkets field.
func (o *TrackObject) SetAvailableMarkets(v []string) {
	o.AvailableMarkets = v
}

// GetDiscNumber returns the DiscNumber field value if set, zero value otherwise.
func (o *TrackObject) GetDiscNumber() int32 {
	if o == nil || IsNil(o.DiscNumber) {
		var ret int32
		return ret
	}
	return *o.DiscNumber
}

// GetDiscNumberOk returns a tuple with the DiscNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetDiscNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscNumber) {
		return nil, false
	}
	return o.DiscNumber, true
}

// HasDiscNumber returns a boolean if a field has been set.
func (o *TrackObject) HasDiscNumber() bool {
	if o != nil && !IsNil(o.DiscNumber) {
		return true
	}

	return false
}

// SetDiscNumber gets a reference to the given int32 and assigns it to the DiscNumber field.
func (o *TrackObject) SetDiscNumber(v int32) {
	o.DiscNumber = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *TrackObject) GetDurationMs() int32 {
	if o == nil || IsNil(o.DurationMs) {
		var ret int32
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetDurationMsOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationMs) {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *TrackObject) HasDurationMs() bool {
	if o != nil && !IsNil(o.DurationMs) {
		return true
	}

	return false
}

// SetDurationMs gets a reference to the given int32 and assigns it to the DurationMs field.
func (o *TrackObject) SetDurationMs(v int32) {
	o.DurationMs = &v
}

// GetExplicit returns the Explicit field value if set, zero value otherwise.
func (o *TrackObject) GetExplicit() bool {
	if o == nil || IsNil(o.Explicit) {
		var ret bool
		return ret
	}
	return *o.Explicit
}

// GetExplicitOk returns a tuple with the Explicit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetExplicitOk() (*bool, bool) {
	if o == nil || IsNil(o.Explicit) {
		return nil, false
	}
	return o.Explicit, true
}

// HasExplicit returns a boolean if a field has been set.
func (o *TrackObject) HasExplicit() bool {
	if o != nil && !IsNil(o.Explicit) {
		return true
	}

	return false
}

// SetExplicit gets a reference to the given bool and assigns it to the Explicit field.
func (o *TrackObject) SetExplicit(v bool) {
	o.Explicit = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *TrackObject) GetExternalIds() ExternalIdObject {
	if o == nil || IsNil(o.ExternalIds) {
		var ret ExternalIdObject
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetExternalIdsOk() (*ExternalIdObject, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *TrackObject) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given ExternalIdObject and assigns it to the ExternalIds field.
func (o *TrackObject) SetExternalIds(v ExternalIdObject) {
	o.ExternalIds = &v
}

// GetExternalUrls returns the ExternalUrls field value if set, zero value otherwise.
func (o *TrackObject) GetExternalUrls() ExternalUrlObject {
	if o == nil || IsNil(o.ExternalUrls) {
		var ret ExternalUrlObject
		return ret
	}
	return *o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetExternalUrlsOk() (*ExternalUrlObject, bool) {
	if o == nil || IsNil(o.ExternalUrls) {
		return nil, false
	}
	return o.ExternalUrls, true
}

// HasExternalUrls returns a boolean if a field has been set.
func (o *TrackObject) HasExternalUrls() bool {
	if o != nil && !IsNil(o.ExternalUrls) {
		return true
	}

	return false
}

// SetExternalUrls gets a reference to the given ExternalUrlObject and assigns it to the ExternalUrls field.
func (o *TrackObject) SetExternalUrls(v ExternalUrlObject) {
	o.ExternalUrls = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *TrackObject) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *TrackObject) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *TrackObject) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrackObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrackObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TrackObject) SetId(v string) {
	o.Id = &v
}

// GetIsPlayable returns the IsPlayable field value if set, zero value otherwise.
func (o *TrackObject) GetIsPlayable() bool {
	if o == nil || IsNil(o.IsPlayable) {
		var ret bool
		return ret
	}
	return *o.IsPlayable
}

// GetIsPlayableOk returns a tuple with the IsPlayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetIsPlayableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPlayable) {
		return nil, false
	}
	return o.IsPlayable, true
}

// HasIsPlayable returns a boolean if a field has been set.
func (o *TrackObject) HasIsPlayable() bool {
	if o != nil && !IsNil(o.IsPlayable) {
		return true
	}

	return false
}

// SetIsPlayable gets a reference to the given bool and assigns it to the IsPlayable field.
func (o *TrackObject) SetIsPlayable(v bool) {
	o.IsPlayable = &v
}

// GetLinkedFrom returns the LinkedFrom field value if set, zero value otherwise.
func (o *TrackObject) GetLinkedFrom() map[string]interface{} {
	if o == nil || IsNil(o.LinkedFrom) {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedFrom
}

// GetLinkedFromOk returns a tuple with the LinkedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetLinkedFromOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LinkedFrom) {
		return map[string]interface{}{}, false
	}
	return o.LinkedFrom, true
}

// HasLinkedFrom returns a boolean if a field has been set.
func (o *TrackObject) HasLinkedFrom() bool {
	if o != nil && !IsNil(o.LinkedFrom) {
		return true
	}

	return false
}

// SetLinkedFrom gets a reference to the given map[string]interface{} and assigns it to the LinkedFrom field.
func (o *TrackObject) SetLinkedFrom(v map[string]interface{}) {
	o.LinkedFrom = v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *TrackObject) GetRestrictions() TrackRestrictionObject {
	if o == nil || IsNil(o.Restrictions) {
		var ret TrackRestrictionObject
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetRestrictionsOk() (*TrackRestrictionObject, bool) {
	if o == nil || IsNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *TrackObject) HasRestrictions() bool {
	if o != nil && !IsNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given TrackRestrictionObject and assigns it to the Restrictions field.
func (o *TrackObject) SetRestrictions(v TrackRestrictionObject) {
	o.Restrictions = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TrackObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TrackObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TrackObject) SetName(v string) {
	o.Name = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *TrackObject) GetPopularity() int32 {
	if o == nil || IsNil(o.Popularity) {
		var ret int32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetPopularityOk() (*int32, bool) {
	if o == nil || IsNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *TrackObject) HasPopularity() bool {
	if o != nil && !IsNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given int32 and assigns it to the Popularity field.
func (o *TrackObject) SetPopularity(v int32) {
	o.Popularity = &v
}

// GetPreviewUrl returns the PreviewUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackObject) GetPreviewUrl() string {
	if o == nil || IsNil(o.PreviewUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PreviewUrl.Get()
}

// GetPreviewUrlOk returns a tuple with the PreviewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackObject) GetPreviewUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviewUrl.Get(), o.PreviewUrl.IsSet()
}

// HasPreviewUrl returns a boolean if a field has been set.
func (o *TrackObject) HasPreviewUrl() bool {
	if o != nil && o.PreviewUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviewUrl gets a reference to the given NullableString and assigns it to the PreviewUrl field.
func (o *TrackObject) SetPreviewUrl(v string) {
	o.PreviewUrl.Set(&v)
}
// SetPreviewUrlNil sets the value for PreviewUrl to be an explicit nil
func (o *TrackObject) SetPreviewUrlNil() {
	o.PreviewUrl.Set(nil)
}

// UnsetPreviewUrl ensures that no value is present for PreviewUrl, not even an explicit nil
func (o *TrackObject) UnsetPreviewUrl() {
	o.PreviewUrl.Unset()
}

// GetTrackNumber returns the TrackNumber field value if set, zero value otherwise.
func (o *TrackObject) GetTrackNumber() int32 {
	if o == nil || IsNil(o.TrackNumber) {
		var ret int32
		return ret
	}
	return *o.TrackNumber
}

// GetTrackNumberOk returns a tuple with the TrackNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetTrackNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.TrackNumber) {
		return nil, false
	}
	return o.TrackNumber, true
}

// HasTrackNumber returns a boolean if a field has been set.
func (o *TrackObject) HasTrackNumber() bool {
	if o != nil && !IsNil(o.TrackNumber) {
		return true
	}

	return false
}

// SetTrackNumber gets a reference to the given int32 and assigns it to the TrackNumber field.
func (o *TrackObject) SetTrackNumber(v int32) {
	o.TrackNumber = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TrackObject) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TrackObject) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TrackObject) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *TrackObject) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *TrackObject) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *TrackObject) SetUri(v string) {
	o.Uri = &v
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *TrackObject) GetIsLocal() bool {
	if o == nil || IsNil(o.IsLocal) {
		var ret bool
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObject) GetIsLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocal) {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *TrackObject) HasIsLocal() bool {
	if o != nil && !IsNil(o.IsLocal) {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given bool and assigns it to the IsLocal field.
func (o *TrackObject) SetIsLocal(v bool) {
	o.IsLocal = &v
}

func (o TrackObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Album) {
		toSerialize["album"] = o.Album
	}
	if !IsNil(o.Artists) {
		toSerialize["artists"] = o.Artists
	}
	if !IsNil(o.AvailableMarkets) {
		toSerialize["available_markets"] = o.AvailableMarkets
	}
	if !IsNil(o.DiscNumber) {
		toSerialize["disc_number"] = o.DiscNumber
	}
	if !IsNil(o.DurationMs) {
		toSerialize["duration_ms"] = o.DurationMs
	}
	if !IsNil(o.Explicit) {
		toSerialize["explicit"] = o.Explicit
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["external_ids"] = o.ExternalIds
	}
	if !IsNil(o.ExternalUrls) {
		toSerialize["external_urls"] = o.ExternalUrls
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsPlayable) {
		toSerialize["is_playable"] = o.IsPlayable
	}
	if !IsNil(o.LinkedFrom) {
		toSerialize["linked_from"] = o.LinkedFrom
	}
	if !IsNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if o.PreviewUrl.IsSet() {
		toSerialize["preview_url"] = o.PreviewUrl.Get()
	}
	if !IsNil(o.TrackNumber) {
		toSerialize["track_number"] = o.TrackNumber
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.IsLocal) {
		toSerialize["is_local"] = o.IsLocal
	}
	return toSerialize, nil
}

type NullableTrackObject struct {
	value *TrackObject
	isSet bool
}

func (v NullableTrackObject) Get() *TrackObject {
	return v.value
}

func (v *NullableTrackObject) Set(val *TrackObject) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackObject) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackObject(val *TrackObject) *NullableTrackObject {
	return &NullableTrackObject{value: val, isSet: true}
}

func (v NullableTrackObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

