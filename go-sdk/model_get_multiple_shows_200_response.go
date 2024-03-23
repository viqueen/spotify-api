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

// checks if the GetMultipleShows200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMultipleShows200Response{}

// GetMultipleShows200Response struct for GetMultipleShows200Response
type GetMultipleShows200Response struct {
	Shows []SimplifiedShowObject `json:"shows"`
}

type _GetMultipleShows200Response GetMultipleShows200Response

// NewGetMultipleShows200Response instantiates a new GetMultipleShows200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMultipleShows200Response(shows []SimplifiedShowObject) *GetMultipleShows200Response {
	this := GetMultipleShows200Response{}
	this.Shows = shows
	return &this
}

// NewGetMultipleShows200ResponseWithDefaults instantiates a new GetMultipleShows200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMultipleShows200ResponseWithDefaults() *GetMultipleShows200Response {
	this := GetMultipleShows200Response{}
	return &this
}

// GetShows returns the Shows field value
func (o *GetMultipleShows200Response) GetShows() []SimplifiedShowObject {
	if o == nil {
		var ret []SimplifiedShowObject
		return ret
	}

	return o.Shows
}

// GetShowsOk returns a tuple with the Shows field value
// and a boolean to check if the value has been set.
func (o *GetMultipleShows200Response) GetShowsOk() ([]SimplifiedShowObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shows, true
}

// SetShows sets field value
func (o *GetMultipleShows200Response) SetShows(v []SimplifiedShowObject) {
	o.Shows = v
}

func (o GetMultipleShows200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMultipleShows200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shows"] = o.Shows
	return toSerialize, nil
}

func (o *GetMultipleShows200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shows",
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

	varGetMultipleShows200Response := _GetMultipleShows200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMultipleShows200Response)

	if err != nil {
		return err
	}

	*o = GetMultipleShows200Response(varGetMultipleShows200Response)

	return err
}

type NullableGetMultipleShows200Response struct {
	value *GetMultipleShows200Response
	isSet bool
}

func (v NullableGetMultipleShows200Response) Get() *GetMultipleShows200Response {
	return v.value
}

func (v *NullableGetMultipleShows200Response) Set(val *GetMultipleShows200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMultipleShows200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMultipleShows200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMultipleShows200Response(val *GetMultipleShows200Response) *NullableGetMultipleShows200Response {
	return &NullableGetMultipleShows200Response{value: val, isSet: true}
}

func (v NullableGetMultipleShows200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMultipleShows200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


