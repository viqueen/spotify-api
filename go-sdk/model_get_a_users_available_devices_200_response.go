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

// checks if the GetAUsersAvailableDevices200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAUsersAvailableDevices200Response{}

// GetAUsersAvailableDevices200Response struct for GetAUsersAvailableDevices200Response
type GetAUsersAvailableDevices200Response struct {
	Devices []DeviceObject `json:"devices"`
}

type _GetAUsersAvailableDevices200Response GetAUsersAvailableDevices200Response

// NewGetAUsersAvailableDevices200Response instantiates a new GetAUsersAvailableDevices200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAUsersAvailableDevices200Response(devices []DeviceObject) *GetAUsersAvailableDevices200Response {
	this := GetAUsersAvailableDevices200Response{}
	this.Devices = devices
	return &this
}

// NewGetAUsersAvailableDevices200ResponseWithDefaults instantiates a new GetAUsersAvailableDevices200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAUsersAvailableDevices200ResponseWithDefaults() *GetAUsersAvailableDevices200Response {
	this := GetAUsersAvailableDevices200Response{}
	return &this
}

// GetDevices returns the Devices field value
func (o *GetAUsersAvailableDevices200Response) GetDevices() []DeviceObject {
	if o == nil {
		var ret []DeviceObject
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *GetAUsersAvailableDevices200Response) GetDevicesOk() ([]DeviceObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *GetAUsersAvailableDevices200Response) SetDevices(v []DeviceObject) {
	o.Devices = v
}

func (o GetAUsersAvailableDevices200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAUsersAvailableDevices200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devices"] = o.Devices
	return toSerialize, nil
}

func (o *GetAUsersAvailableDevices200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"devices",
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

	varGetAUsersAvailableDevices200Response := _GetAUsersAvailableDevices200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAUsersAvailableDevices200Response)

	if err != nil {
		return err
	}

	*o = GetAUsersAvailableDevices200Response(varGetAUsersAvailableDevices200Response)

	return err
}

type NullableGetAUsersAvailableDevices200Response struct {
	value *GetAUsersAvailableDevices200Response
	isSet bool
}

func (v NullableGetAUsersAvailableDevices200Response) Get() *GetAUsersAvailableDevices200Response {
	return v.value
}

func (v *NullableGetAUsersAvailableDevices200Response) Set(val *GetAUsersAvailableDevices200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAUsersAvailableDevices200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAUsersAvailableDevices200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAUsersAvailableDevices200Response(val *GetAUsersAvailableDevices200Response) *NullableGetAUsersAvailableDevices200Response {
	return &NullableGetAUsersAvailableDevices200Response{value: val, isSet: true}
}

func (v NullableGetAUsersAvailableDevices200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAUsersAvailableDevices200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

