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

// checks if the GetMultipleArtists200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMultipleArtists200Response{}

// GetMultipleArtists200Response struct for GetMultipleArtists200Response
type GetMultipleArtists200Response struct {
	Artists []ArtistObject `json:"artists"`
}

type _GetMultipleArtists200Response GetMultipleArtists200Response

// NewGetMultipleArtists200Response instantiates a new GetMultipleArtists200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMultipleArtists200Response(artists []ArtistObject) *GetMultipleArtists200Response {
	this := GetMultipleArtists200Response{}
	this.Artists = artists
	return &this
}

// NewGetMultipleArtists200ResponseWithDefaults instantiates a new GetMultipleArtists200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMultipleArtists200ResponseWithDefaults() *GetMultipleArtists200Response {
	this := GetMultipleArtists200Response{}
	return &this
}

// GetArtists returns the Artists field value
func (o *GetMultipleArtists200Response) GetArtists() []ArtistObject {
	if o == nil {
		var ret []ArtistObject
		return ret
	}

	return o.Artists
}

// GetArtistsOk returns a tuple with the Artists field value
// and a boolean to check if the value has been set.
func (o *GetMultipleArtists200Response) GetArtistsOk() ([]ArtistObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Artists, true
}

// SetArtists sets field value
func (o *GetMultipleArtists200Response) SetArtists(v []ArtistObject) {
	o.Artists = v
}

func (o GetMultipleArtists200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMultipleArtists200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artists"] = o.Artists
	return toSerialize, nil
}

func (o *GetMultipleArtists200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"artists",
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

	varGetMultipleArtists200Response := _GetMultipleArtists200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMultipleArtists200Response)

	if err != nil {
		return err
	}

	*o = GetMultipleArtists200Response(varGetMultipleArtists200Response)

	return err
}

type NullableGetMultipleArtists200Response struct {
	value *GetMultipleArtists200Response
	isSet bool
}

func (v NullableGetMultipleArtists200Response) Get() *GetMultipleArtists200Response {
	return v.value
}

func (v *NullableGetMultipleArtists200Response) Set(val *GetMultipleArtists200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMultipleArtists200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMultipleArtists200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMultipleArtists200Response(val *GetMultipleArtists200Response) *NullableGetMultipleArtists200Response {
	return &NullableGetMultipleArtists200Response{value: val, isSet: true}
}

func (v NullableGetMultipleArtists200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMultipleArtists200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


