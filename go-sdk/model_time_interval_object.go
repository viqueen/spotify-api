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

// checks if the TimeIntervalObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeIntervalObject{}

// TimeIntervalObject struct for TimeIntervalObject
type TimeIntervalObject struct {
	// The starting point (in seconds) of the time interval.
	Start *float32 `json:"start,omitempty"`
	// The duration (in seconds) of the time interval.
	Duration *float32 `json:"duration,omitempty"`
	// The confidence, from 0.0 to 1.0, of the reliability of the interval.
	Confidence *float32 `json:"confidence,omitempty"`
}

// NewTimeIntervalObject instantiates a new TimeIntervalObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeIntervalObject() *TimeIntervalObject {
	this := TimeIntervalObject{}
	return &this
}

// NewTimeIntervalObjectWithDefaults instantiates a new TimeIntervalObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeIntervalObjectWithDefaults() *TimeIntervalObject {
	this := TimeIntervalObject{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *TimeIntervalObject) GetStart() float32 {
	if o == nil || IsNil(o.Start) {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeIntervalObject) GetStartOk() (*float32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *TimeIntervalObject) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *TimeIntervalObject) SetStart(v float32) {
	o.Start = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *TimeIntervalObject) GetDuration() float32 {
	if o == nil || IsNil(o.Duration) {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeIntervalObject) GetDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *TimeIntervalObject) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *TimeIntervalObject) SetDuration(v float32) {
	o.Duration = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *TimeIntervalObject) GetConfidence() float32 {
	if o == nil || IsNil(o.Confidence) {
		var ret float32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeIntervalObject) GetConfidenceOk() (*float32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *TimeIntervalObject) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given float32 and assigns it to the Confidence field.
func (o *TimeIntervalObject) SetConfidence(v float32) {
	o.Confidence = &v
}

func (o TimeIntervalObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeIntervalObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	return toSerialize, nil
}

type NullableTimeIntervalObject struct {
	value *TimeIntervalObject
	isSet bool
}

func (v NullableTimeIntervalObject) Get() *TimeIntervalObject {
	return v.value
}

func (v *NullableTimeIntervalObject) Set(val *TimeIntervalObject) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeIntervalObject) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeIntervalObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeIntervalObject(val *TimeIntervalObject) *NullableTimeIntervalObject {
	return &NullableTimeIntervalObject{value: val, isSet: true}
}

func (v NullableTimeIntervalObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeIntervalObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

