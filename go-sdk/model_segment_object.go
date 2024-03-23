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

// checks if the SegmentObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentObject{}

// SegmentObject struct for SegmentObject
type SegmentObject struct {
	// The starting point (in seconds) of the segment.
	Start *float32 `json:"start,omitempty"`
	// The duration (in seconds) of the segment.
	Duration *float32 `json:"duration,omitempty"`
	// The confidence, from 0.0 to 1.0, of the reliability of the segmentation. Segments of the song which are difficult to logically segment (e.g: noise) may correspond to low values in this field. 
	Confidence *float32 `json:"confidence,omitempty"`
	// The onset loudness of the segment in decibels (dB). Combined with `loudness_max` and `loudness_max_time`, these components can be used to describe the \"attack\" of the segment.
	LoudnessStart *float32 `json:"loudness_start,omitempty"`
	// The peak loudness of the segment in decibels (dB). Combined with `loudness_start` and `loudness_max_time`, these components can be used to describe the \"attack\" of the segment.
	LoudnessMax *float32 `json:"loudness_max,omitempty"`
	// The segment-relative offset of the segment peak loudness in seconds. Combined with `loudness_start` and `loudness_max`, these components can be used to desctibe the \"attack\" of the segment.
	LoudnessMaxTime *float32 `json:"loudness_max_time,omitempty"`
	// The offset loudness of the segment in decibels (dB). This value should be equivalent to the loudness_start of the following segment.
	LoudnessEnd *float32 `json:"loudness_end,omitempty"`
	// Pitch content is given by a “chroma” vector, corresponding to the 12 pitch classes C, C#, D to B, with values ranging from 0 to 1 that describe the relative dominance of every pitch in the chromatic scale. For example a C Major chord would likely be represented by large values of C, E and G (i.e. classes 0, 4, and 7).  Vectors are normalized to 1 by their strongest dimension, therefore noisy sounds are likely represented by values that are all close to 1, while pure tones are described by one value at 1 (the pitch) and others near 0. As can be seen below, the 12 vector indices are a combination of low-power spectrum values at their respective pitch frequencies. ![pitch vector](https://developer.spotify.com/assets/audio/Pitch_vector.png) 
	Pitches []float32 `json:"pitches,omitempty"`
	// Timbre is the quality of a musical note or sound that distinguishes different types of musical instruments, or voices. It is a complex notion also referred to as sound color, texture, or tone quality, and is derived from the shape of a segment’s spectro-temporal surface, independently of pitch and loudness. The timbre feature is a vector that includes 12 unbounded values roughly centered around 0. Those values are high level abstractions of the spectral surface, ordered by degree of importance.  For completeness however, the first dimension represents the average loudness of the segment; second emphasizes brightness; third is more closely correlated to the flatness of a sound; fourth to sounds with a stronger attack; etc. See an image below representing the 12 basis functions (i.e. template segments). ![timbre basis functions](https://developer.spotify.com/assets/audio/Timbre_basis_functions.png)  The actual timbre of the segment is best described as a linear combination of these 12 basis functions weighted by the coefficient values: timbre = c1 x b1 + c2 x b2 + ... + c12 x b12, where c1 to c12 represent the 12 coefficients and b1 to b12 the 12 basis functions as displayed below. Timbre vectors are best used in comparison with each other. 
	Timbre []float32 `json:"timbre,omitempty"`
}

// NewSegmentObject instantiates a new SegmentObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentObject() *SegmentObject {
	this := SegmentObject{}
	return &this
}

// NewSegmentObjectWithDefaults instantiates a new SegmentObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentObjectWithDefaults() *SegmentObject {
	this := SegmentObject{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SegmentObject) GetStart() float32 {
	if o == nil || IsNil(o.Start) {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentObject) GetStartOk() (*float32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SegmentObject) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *SegmentObject) SetStart(v float32) {
	o.Start = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SegmentObject) GetDuration() float32 {
	if o == nil || IsNil(o.Duration) {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentObject) GetDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SegmentObject) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *SegmentObject) SetDuration(v float32) {
	o.Duration = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *SegmentObject) GetConfidence() float32 {
	if o == nil || IsNil(o.Confidence) {
		var ret float32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentObject) GetConfidenceOk() (*float32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *SegmentObject) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given float32 and assigns it to the Confidence field.
func (o *SegmentObject) SetConfidence(v float32) {
	o.Confidence = &v
}

// GetLoudnessStart returns the LoudnessStart field value if set, zero value otherwise.
func (o *SegmentObject) GetLoudnessStart() float32 {
	if o == nil || IsNil(o.LoudnessStart) {
		var ret float32
		return ret
	}
	return *o.LoudnessStart
}

// GetLoudnessStartOk returns a tuple with the LoudnessStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentObject) GetLoudnessStartOk() (*float32, bool) {
	if o == nil || IsNil(o.LoudnessStart) {
		return nil, false
	}
	return o.LoudnessStart, true
}

// HasLoudnessStart returns a boolean if a field has been set.
func (o *SegmentObject) HasLoudnessStart() bool {
	if o != nil && !IsNil(o.LoudnessStart) {
		return true
	}

	return false
}

// SetLoudnessStart gets a reference to the given float32 and assigns it to the LoudnessStart field.
func (o *SegmentObject) SetLoudnessStart(v float32) {
	o.LoudnessStart = &v
}

// GetLoudnessMax returns the LoudnessMax field value if set, zero value otherwise.
func (o *SegmentObject) GetLoudnessMax() float32 {
	if o == nil || IsNil(o.LoudnessMax) {
		var ret float32
		return ret
	}
	return *o.LoudnessMax
}

// GetLoudnessMaxOk returns a tuple with the LoudnessMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentObject) GetLoudnessMaxOk() (*float32, bool) {
	if o == nil || IsNil(o.LoudnessMax) {
		return nil, false
	}
	return o.LoudnessMax, true
}

// HasLoudnessMax returns a boolean if a field has been set.
func (o *SegmentObject) HasLoudnessMax() bool {
	if o != nil && !IsNil(o.LoudnessMax) {
		return true
	}

	return false
}

// SetLoudnessMax gets a reference to the given float32 and assigns it to the LoudnessMax field.
func (o *SegmentObject) SetLoudnessMax(v float32) {
	o.LoudnessMax = &v
}

// GetLoudnessMaxTime returns the LoudnessMaxTime field value if set, zero value otherwise.
func (o *SegmentObject) GetLoudnessMaxTime() float32 {
	if o == nil || IsNil(o.LoudnessMaxTime) {
		var ret float32
		return ret
	}
	return *o.LoudnessMaxTime
}

// GetLoudnessMaxTimeOk returns a tuple with the LoudnessMaxTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentObject) GetLoudnessMaxTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.LoudnessMaxTime) {
		return nil, false
	}
	return o.LoudnessMaxTime, true
}

// HasLoudnessMaxTime returns a boolean if a field has been set.
func (o *SegmentObject) HasLoudnessMaxTime() bool {
	if o != nil && !IsNil(o.LoudnessMaxTime) {
		return true
	}

	return false
}

// SetLoudnessMaxTime gets a reference to the given float32 and assigns it to the LoudnessMaxTime field.
func (o *SegmentObject) SetLoudnessMaxTime(v float32) {
	o.LoudnessMaxTime = &v
}

// GetLoudnessEnd returns the LoudnessEnd field value if set, zero value otherwise.
func (o *SegmentObject) GetLoudnessEnd() float32 {
	if o == nil || IsNil(o.LoudnessEnd) {
		var ret float32
		return ret
	}
	return *o.LoudnessEnd
}

// GetLoudnessEndOk returns a tuple with the LoudnessEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentObject) GetLoudnessEndOk() (*float32, bool) {
	if o == nil || IsNil(o.LoudnessEnd) {
		return nil, false
	}
	return o.LoudnessEnd, true
}

// HasLoudnessEnd returns a boolean if a field has been set.
func (o *SegmentObject) HasLoudnessEnd() bool {
	if o != nil && !IsNil(o.LoudnessEnd) {
		return true
	}

	return false
}

// SetLoudnessEnd gets a reference to the given float32 and assigns it to the LoudnessEnd field.
func (o *SegmentObject) SetLoudnessEnd(v float32) {
	o.LoudnessEnd = &v
}

// GetPitches returns the Pitches field value if set, zero value otherwise.
func (o *SegmentObject) GetPitches() []float32 {
	if o == nil || IsNil(o.Pitches) {
		var ret []float32
		return ret
	}
	return o.Pitches
}

// GetPitchesOk returns a tuple with the Pitches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentObject) GetPitchesOk() ([]float32, bool) {
	if o == nil || IsNil(o.Pitches) {
		return nil, false
	}
	return o.Pitches, true
}

// HasPitches returns a boolean if a field has been set.
func (o *SegmentObject) HasPitches() bool {
	if o != nil && !IsNil(o.Pitches) {
		return true
	}

	return false
}

// SetPitches gets a reference to the given []float32 and assigns it to the Pitches field.
func (o *SegmentObject) SetPitches(v []float32) {
	o.Pitches = v
}

// GetTimbre returns the Timbre field value if set, zero value otherwise.
func (o *SegmentObject) GetTimbre() []float32 {
	if o == nil || IsNil(o.Timbre) {
		var ret []float32
		return ret
	}
	return o.Timbre
}

// GetTimbreOk returns a tuple with the Timbre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentObject) GetTimbreOk() ([]float32, bool) {
	if o == nil || IsNil(o.Timbre) {
		return nil, false
	}
	return o.Timbre, true
}

// HasTimbre returns a boolean if a field has been set.
func (o *SegmentObject) HasTimbre() bool {
	if o != nil && !IsNil(o.Timbre) {
		return true
	}

	return false
}

// SetTimbre gets a reference to the given []float32 and assigns it to the Timbre field.
func (o *SegmentObject) SetTimbre(v []float32) {
	o.Timbre = v
}

func (o SegmentObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentObject) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LoudnessStart) {
		toSerialize["loudness_start"] = o.LoudnessStart
	}
	if !IsNil(o.LoudnessMax) {
		toSerialize["loudness_max"] = o.LoudnessMax
	}
	if !IsNil(o.LoudnessMaxTime) {
		toSerialize["loudness_max_time"] = o.LoudnessMaxTime
	}
	if !IsNil(o.LoudnessEnd) {
		toSerialize["loudness_end"] = o.LoudnessEnd
	}
	if !IsNil(o.Pitches) {
		toSerialize["pitches"] = o.Pitches
	}
	if !IsNil(o.Timbre) {
		toSerialize["timbre"] = o.Timbre
	}
	return toSerialize, nil
}

type NullableSegmentObject struct {
	value *SegmentObject
	isSet bool
}

func (v NullableSegmentObject) Get() *SegmentObject {
	return v.value
}

func (v *NullableSegmentObject) Set(val *SegmentObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentObject(val *SegmentObject) *NullableSegmentObject {
	return &NullableSegmentObject{value: val, isSet: true}
}

func (v NullableSegmentObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

