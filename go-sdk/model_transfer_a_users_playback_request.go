/*
Spotify Web API

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the TransferAUsersPlaybackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferAUsersPlaybackRequest{}

// TransferAUsersPlaybackRequest struct for TransferAUsersPlaybackRequest
type TransferAUsersPlaybackRequest struct {
	// A JSON array containing the ID of the device on which playback should be started/transferred.<br/>For example:`{device_ids:[\"74ASZWbe4lXaubB36ztrGX\"]}`<br/>_**Note**: Although an array is accepted, only a single device_id is currently supported. Supplying more than one will return `400 Bad Request`_ 
	DeviceIds []string `json:"device_ids"`
	// **true**: ensure playback happens on new device.<br/>**false** or not provided: keep the current playback state. 
	Play map[string]interface{} `json:"play,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferAUsersPlaybackRequest TransferAUsersPlaybackRequest

// NewTransferAUsersPlaybackRequest instantiates a new TransferAUsersPlaybackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAUsersPlaybackRequest(deviceIds []string) *TransferAUsersPlaybackRequest {
	this := TransferAUsersPlaybackRequest{}
	this.DeviceIds = deviceIds
	return &this
}

// NewTransferAUsersPlaybackRequestWithDefaults instantiates a new TransferAUsersPlaybackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAUsersPlaybackRequestWithDefaults() *TransferAUsersPlaybackRequest {
	this := TransferAUsersPlaybackRequest{}
	return &this
}

// GetDeviceIds returns the DeviceIds field value
func (o *TransferAUsersPlaybackRequest) GetDeviceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value
// and a boolean to check if the value has been set.
func (o *TransferAUsersPlaybackRequest) GetDeviceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceIds, true
}

// SetDeviceIds sets field value
func (o *TransferAUsersPlaybackRequest) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetPlay returns the Play field value if set, zero value otherwise.
func (o *TransferAUsersPlaybackRequest) GetPlay() map[string]interface{} {
	if o == nil || IsNil(o.Play) {
		var ret map[string]interface{}
		return ret
	}
	return o.Play
}

// GetPlayOk returns a tuple with the Play field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAUsersPlaybackRequest) GetPlayOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Play) {
		return map[string]interface{}{}, false
	}
	return o.Play, true
}

// HasPlay returns a boolean if a field has been set.
func (o *TransferAUsersPlaybackRequest) HasPlay() bool {
	if o != nil && !IsNil(o.Play) {
		return true
	}

	return false
}

// SetPlay gets a reference to the given map[string]interface{} and assigns it to the Play field.
func (o *TransferAUsersPlaybackRequest) SetPlay(v map[string]interface{}) {
	o.Play = v
}

func (o TransferAUsersPlaybackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferAUsersPlaybackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device_ids"] = o.DeviceIds
	if !IsNil(o.Play) {
		toSerialize["play"] = o.Play
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransferAUsersPlaybackRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device_ids",
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

	varTransferAUsersPlaybackRequest := _TransferAUsersPlaybackRequest{}

	err = json.Unmarshal(data, &varTransferAUsersPlaybackRequest)

	if err != nil {
		return err
	}

	*o = TransferAUsersPlaybackRequest(varTransferAUsersPlaybackRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_ids")
		delete(additionalProperties, "play")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferAUsersPlaybackRequest struct {
	value *TransferAUsersPlaybackRequest
	isSet bool
}

func (v NullableTransferAUsersPlaybackRequest) Get() *TransferAUsersPlaybackRequest {
	return v.value
}

func (v *NullableTransferAUsersPlaybackRequest) Set(val *TransferAUsersPlaybackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAUsersPlaybackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAUsersPlaybackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAUsersPlaybackRequest(val *TransferAUsersPlaybackRequest) *NullableTransferAUsersPlaybackRequest {
	return &NullableTransferAUsersPlaybackRequest{value: val, isSet: true}
}

func (v NullableTransferAUsersPlaybackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAUsersPlaybackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


