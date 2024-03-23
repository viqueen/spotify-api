/*
Spotify Web API

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RemoveTracksPlaylistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveTracksPlaylistRequest{}

// RemoveTracksPlaylistRequest struct for RemoveTracksPlaylistRequest
type RemoveTracksPlaylistRequest struct {
	// An array of objects containing [Spotify URIs](/documentation/web-api/concepts/spotify-uris-ids) of the tracks or episodes to remove. For example: `{ \"tracks\": [{ \"uri\": \"spotify:track:4iV5W9uYEdYUVa79Axb7Rh\" },{ \"uri\": \"spotify:track:1301WleyT98MSxVHPZCA6M\" }] }`. A maximum of 100 objects can be sent at once. 
	Tracks []RemoveTracksPlaylistRequestTracksInner `json:"tracks"`
	// The playlist's snapshot ID against which you want to make the changes. The API will validate that the specified items exist and in the specified positions and make the changes, even if more recent changes have been made to the playlist. 
	SnapshotId *string `json:"snapshot_id,omitempty"`
}

type _RemoveTracksPlaylistRequest RemoveTracksPlaylistRequest

// NewRemoveTracksPlaylistRequest instantiates a new RemoveTracksPlaylistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveTracksPlaylistRequest(tracks []RemoveTracksPlaylistRequestTracksInner) *RemoveTracksPlaylistRequest {
	this := RemoveTracksPlaylistRequest{}
	this.Tracks = tracks
	return &this
}

// NewRemoveTracksPlaylistRequestWithDefaults instantiates a new RemoveTracksPlaylistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveTracksPlaylistRequestWithDefaults() *RemoveTracksPlaylistRequest {
	this := RemoveTracksPlaylistRequest{}
	return &this
}

// GetTracks returns the Tracks field value
func (o *RemoveTracksPlaylistRequest) GetTracks() []RemoveTracksPlaylistRequestTracksInner {
	if o == nil {
		var ret []RemoveTracksPlaylistRequestTracksInner
		return ret
	}

	return o.Tracks
}

// GetTracksOk returns a tuple with the Tracks field value
// and a boolean to check if the value has been set.
func (o *RemoveTracksPlaylistRequest) GetTracksOk() ([]RemoveTracksPlaylistRequestTracksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tracks, true
}

// SetTracks sets field value
func (o *RemoveTracksPlaylistRequest) SetTracks(v []RemoveTracksPlaylistRequestTracksInner) {
	o.Tracks = v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *RemoveTracksPlaylistRequest) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveTracksPlaylistRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *RemoveTracksPlaylistRequest) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *RemoveTracksPlaylistRequest) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

func (o RemoveTracksPlaylistRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveTracksPlaylistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tracks"] = o.Tracks
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	return toSerialize, nil
}

func (o *RemoveTracksPlaylistRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tracks",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRemoveTracksPlaylistRequest := _RemoveTracksPlaylistRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRemoveTracksPlaylistRequest)

	if err != nil {
		return err
	}

	*o = RemoveTracksPlaylistRequest(varRemoveTracksPlaylistRequest)

	return err
}

type NullableRemoveTracksPlaylistRequest struct {
	value *RemoveTracksPlaylistRequest
	isSet bool
}

func (v NullableRemoveTracksPlaylistRequest) Get() *RemoveTracksPlaylistRequest {
	return v.value
}

func (v *NullableRemoveTracksPlaylistRequest) Set(val *RemoveTracksPlaylistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveTracksPlaylistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveTracksPlaylistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveTracksPlaylistRequest(val *RemoveTracksPlaylistRequest) *NullableRemoveTracksPlaylistRequest {
	return &NullableRemoveTracksPlaylistRequest{value: val, isSet: true}
}

func (v NullableRemoveTracksPlaylistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveTracksPlaylistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


