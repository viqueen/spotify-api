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

// checks if the StartAUsersPlaybackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartAUsersPlaybackRequest{}

// StartAUsersPlaybackRequest struct for StartAUsersPlaybackRequest
type StartAUsersPlaybackRequest struct {
	// Optional. Spotify URI of the context to play. Valid contexts are albums, artists & playlists. `{context_uri:\"spotify:album:1Je1IMUlBXcx1Fz0WE7oPT\"}` 
	ContextUri map[string]interface{} `json:"context_uri,omitempty"`
	// Optional. A JSON array of the Spotify track URIs to play. For example: `{\"uris\": [\"spotify:track:4iV5W9uYEdYUVa79Axb7Rh\", \"spotify:track:1301WleyT98MSxVHPZCA6M\"]}` 
	Uris []string `json:"uris,omitempty"`
	// Optional. Indicates from where in the context playback should start. Only available when context_uri corresponds to an album or playlist object \"position\" is zero based and can’t be negative. Example: `\"offset\": {\"position\": 5}` \"uri\" is a string representing the uri of the item to start at. Example: `\"offset\": {\"uri\": \"spotify:track:1301WleyT98MSxVHPZCA6M\"}` 
	Offset map[string]interface{} `json:"offset,omitempty"`
	// integer
	PositionMs map[string]interface{} `json:"position_ms,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StartAUsersPlaybackRequest StartAUsersPlaybackRequest

// NewStartAUsersPlaybackRequest instantiates a new StartAUsersPlaybackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartAUsersPlaybackRequest() *StartAUsersPlaybackRequest {
	this := StartAUsersPlaybackRequest{}
	return &this
}

// NewStartAUsersPlaybackRequestWithDefaults instantiates a new StartAUsersPlaybackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartAUsersPlaybackRequestWithDefaults() *StartAUsersPlaybackRequest {
	this := StartAUsersPlaybackRequest{}
	return &this
}

// GetContextUri returns the ContextUri field value if set, zero value otherwise.
func (o *StartAUsersPlaybackRequest) GetContextUri() map[string]interface{} {
	if o == nil || IsNil(o.ContextUri) {
		var ret map[string]interface{}
		return ret
	}
	return o.ContextUri
}

// GetContextUriOk returns a tuple with the ContextUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAUsersPlaybackRequest) GetContextUriOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ContextUri) {
		return map[string]interface{}{}, false
	}
	return o.ContextUri, true
}

// HasContextUri returns a boolean if a field has been set.
func (o *StartAUsersPlaybackRequest) HasContextUri() bool {
	if o != nil && !IsNil(o.ContextUri) {
		return true
	}

	return false
}

// SetContextUri gets a reference to the given map[string]interface{} and assigns it to the ContextUri field.
func (o *StartAUsersPlaybackRequest) SetContextUri(v map[string]interface{}) {
	o.ContextUri = v
}

// GetUris returns the Uris field value if set, zero value otherwise.
func (o *StartAUsersPlaybackRequest) GetUris() []string {
	if o == nil || IsNil(o.Uris) {
		var ret []string
		return ret
	}
	return o.Uris
}

// GetUrisOk returns a tuple with the Uris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAUsersPlaybackRequest) GetUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.Uris) {
		return nil, false
	}
	return o.Uris, true
}

// HasUris returns a boolean if a field has been set.
func (o *StartAUsersPlaybackRequest) HasUris() bool {
	if o != nil && !IsNil(o.Uris) {
		return true
	}

	return false
}

// SetUris gets a reference to the given []string and assigns it to the Uris field.
func (o *StartAUsersPlaybackRequest) SetUris(v []string) {
	o.Uris = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *StartAUsersPlaybackRequest) GetOffset() map[string]interface{} {
	if o == nil || IsNil(o.Offset) {
		var ret map[string]interface{}
		return ret
	}
	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAUsersPlaybackRequest) GetOffsetOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Offset) {
		return map[string]interface{}{}, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *StartAUsersPlaybackRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given map[string]interface{} and assigns it to the Offset field.
func (o *StartAUsersPlaybackRequest) SetOffset(v map[string]interface{}) {
	o.Offset = v
}

// GetPositionMs returns the PositionMs field value if set, zero value otherwise.
func (o *StartAUsersPlaybackRequest) GetPositionMs() map[string]interface{} {
	if o == nil || IsNil(o.PositionMs) {
		var ret map[string]interface{}
		return ret
	}
	return o.PositionMs
}

// GetPositionMsOk returns a tuple with the PositionMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAUsersPlaybackRequest) GetPositionMsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PositionMs) {
		return map[string]interface{}{}, false
	}
	return o.PositionMs, true
}

// HasPositionMs returns a boolean if a field has been set.
func (o *StartAUsersPlaybackRequest) HasPositionMs() bool {
	if o != nil && !IsNil(o.PositionMs) {
		return true
	}

	return false
}

// SetPositionMs gets a reference to the given map[string]interface{} and assigns it to the PositionMs field.
func (o *StartAUsersPlaybackRequest) SetPositionMs(v map[string]interface{}) {
	o.PositionMs = v
}

func (o StartAUsersPlaybackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartAUsersPlaybackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContextUri) {
		toSerialize["context_uri"] = o.ContextUri
	}
	if !IsNil(o.Uris) {
		toSerialize["uris"] = o.Uris
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.PositionMs) {
		toSerialize["position_ms"] = o.PositionMs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StartAUsersPlaybackRequest) UnmarshalJSON(data []byte) (err error) {
	varStartAUsersPlaybackRequest := _StartAUsersPlaybackRequest{}

	err = json.Unmarshal(data, &varStartAUsersPlaybackRequest)

	if err != nil {
		return err
	}

	*o = StartAUsersPlaybackRequest(varStartAUsersPlaybackRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "context_uri")
		delete(additionalProperties, "uris")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "position_ms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStartAUsersPlaybackRequest struct {
	value *StartAUsersPlaybackRequest
	isSet bool
}

func (v NullableStartAUsersPlaybackRequest) Get() *StartAUsersPlaybackRequest {
	return v.value
}

func (v *NullableStartAUsersPlaybackRequest) Set(val *StartAUsersPlaybackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStartAUsersPlaybackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStartAUsersPlaybackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartAUsersPlaybackRequest(val *StartAUsersPlaybackRequest) *NullableStartAUsersPlaybackRequest {
	return &NullableStartAUsersPlaybackRequest{value: val, isSet: true}
}

func (v NullableStartAUsersPlaybackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartAUsersPlaybackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


