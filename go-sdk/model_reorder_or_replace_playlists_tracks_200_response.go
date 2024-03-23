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

// checks if the ReorderOrReplacePlaylistsTracks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReorderOrReplacePlaylistsTracks200Response{}

// ReorderOrReplacePlaylistsTracks200Response struct for ReorderOrReplacePlaylistsTracks200Response
type ReorderOrReplacePlaylistsTracks200Response struct {
	SnapshotId *string `json:"snapshot_id,omitempty"`
}

// NewReorderOrReplacePlaylistsTracks200Response instantiates a new ReorderOrReplacePlaylistsTracks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReorderOrReplacePlaylistsTracks200Response() *ReorderOrReplacePlaylistsTracks200Response {
	this := ReorderOrReplacePlaylistsTracks200Response{}
	return &this
}

// NewReorderOrReplacePlaylistsTracks200ResponseWithDefaults instantiates a new ReorderOrReplacePlaylistsTracks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReorderOrReplacePlaylistsTracks200ResponseWithDefaults() *ReorderOrReplacePlaylistsTracks200Response {
	this := ReorderOrReplacePlaylistsTracks200Response{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *ReorderOrReplacePlaylistsTracks200Response) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReorderOrReplacePlaylistsTracks200Response) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ReorderOrReplacePlaylistsTracks200Response) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *ReorderOrReplacePlaylistsTracks200Response) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

func (o ReorderOrReplacePlaylistsTracks200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReorderOrReplacePlaylistsTracks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	return toSerialize, nil
}

type NullableReorderOrReplacePlaylistsTracks200Response struct {
	value *ReorderOrReplacePlaylistsTracks200Response
	isSet bool
}

func (v NullableReorderOrReplacePlaylistsTracks200Response) Get() *ReorderOrReplacePlaylistsTracks200Response {
	return v.value
}

func (v *NullableReorderOrReplacePlaylistsTracks200Response) Set(val *ReorderOrReplacePlaylistsTracks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReorderOrReplacePlaylistsTracks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReorderOrReplacePlaylistsTracks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReorderOrReplacePlaylistsTracks200Response(val *ReorderOrReplacePlaylistsTracks200Response) *NullableReorderOrReplacePlaylistsTracks200Response {
	return &NullableReorderOrReplacePlaylistsTracks200Response{value: val, isSet: true}
}

func (v NullableReorderOrReplacePlaylistsTracks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReorderOrReplacePlaylistsTracks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

