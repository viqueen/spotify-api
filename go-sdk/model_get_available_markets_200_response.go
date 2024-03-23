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

// checks if the GetAvailableMarkets200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAvailableMarkets200Response{}

// GetAvailableMarkets200Response struct for GetAvailableMarkets200Response
type GetAvailableMarkets200Response struct {
	Markets []string `json:"markets,omitempty"`
}

// NewGetAvailableMarkets200Response instantiates a new GetAvailableMarkets200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAvailableMarkets200Response() *GetAvailableMarkets200Response {
	this := GetAvailableMarkets200Response{}
	return &this
}

// NewGetAvailableMarkets200ResponseWithDefaults instantiates a new GetAvailableMarkets200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAvailableMarkets200ResponseWithDefaults() *GetAvailableMarkets200Response {
	this := GetAvailableMarkets200Response{}
	return &this
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *GetAvailableMarkets200Response) GetMarkets() []string {
	if o == nil || IsNil(o.Markets) {
		var ret []string
		return ret
	}
	return o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAvailableMarkets200Response) GetMarketsOk() ([]string, bool) {
	if o == nil || IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *GetAvailableMarkets200Response) HasMarkets() bool {
	if o != nil && !IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given []string and assigns it to the Markets field.
func (o *GetAvailableMarkets200Response) SetMarkets(v []string) {
	o.Markets = v
}

func (o GetAvailableMarkets200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAvailableMarkets200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Markets) {
		toSerialize["markets"] = o.Markets
	}
	return toSerialize, nil
}

type NullableGetAvailableMarkets200Response struct {
	value *GetAvailableMarkets200Response
	isSet bool
}

func (v NullableGetAvailableMarkets200Response) Get() *GetAvailableMarkets200Response {
	return v.value
}

func (v *NullableGetAvailableMarkets200Response) Set(val *GetAvailableMarkets200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAvailableMarkets200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAvailableMarkets200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAvailableMarkets200Response(val *GetAvailableMarkets200Response) *NullableGetAvailableMarkets200Response {
	return &NullableGetAvailableMarkets200Response{value: val, isSet: true}
}

func (v NullableGetAvailableMarkets200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAvailableMarkets200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


