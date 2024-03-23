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

// checks if the QueueObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueObject{}

// QueueObject struct for QueueObject
type QueueObject struct {
	CurrentlyPlaying *QueueObjectCurrentlyPlaying `json:"currently_playing,omitempty"`
	// The tracks or episodes in the queue. Can be empty.
	Queue []QueueObjectQueueInner `json:"queue,omitempty"`
}

// NewQueueObject instantiates a new QueueObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueObject() *QueueObject {
	this := QueueObject{}
	return &this
}

// NewQueueObjectWithDefaults instantiates a new QueueObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueObjectWithDefaults() *QueueObject {
	this := QueueObject{}
	return &this
}

// GetCurrentlyPlaying returns the CurrentlyPlaying field value if set, zero value otherwise.
func (o *QueueObject) GetCurrentlyPlaying() QueueObjectCurrentlyPlaying {
	if o == nil || IsNil(o.CurrentlyPlaying) {
		var ret QueueObjectCurrentlyPlaying
		return ret
	}
	return *o.CurrentlyPlaying
}

// GetCurrentlyPlayingOk returns a tuple with the CurrentlyPlaying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueObject) GetCurrentlyPlayingOk() (*QueueObjectCurrentlyPlaying, bool) {
	if o == nil || IsNil(o.CurrentlyPlaying) {
		return nil, false
	}
	return o.CurrentlyPlaying, true
}

// HasCurrentlyPlaying returns a boolean if a field has been set.
func (o *QueueObject) HasCurrentlyPlaying() bool {
	if o != nil && !IsNil(o.CurrentlyPlaying) {
		return true
	}

	return false
}

// SetCurrentlyPlaying gets a reference to the given QueueObjectCurrentlyPlaying and assigns it to the CurrentlyPlaying field.
func (o *QueueObject) SetCurrentlyPlaying(v QueueObjectCurrentlyPlaying) {
	o.CurrentlyPlaying = &v
}

// GetQueue returns the Queue field value if set, zero value otherwise.
func (o *QueueObject) GetQueue() []QueueObjectQueueInner {
	if o == nil || IsNil(o.Queue) {
		var ret []QueueObjectQueueInner
		return ret
	}
	return o.Queue
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueObject) GetQueueOk() ([]QueueObjectQueueInner, bool) {
	if o == nil || IsNil(o.Queue) {
		return nil, false
	}
	return o.Queue, true
}

// HasQueue returns a boolean if a field has been set.
func (o *QueueObject) HasQueue() bool {
	if o != nil && !IsNil(o.Queue) {
		return true
	}

	return false
}

// SetQueue gets a reference to the given []QueueObjectQueueInner and assigns it to the Queue field.
func (o *QueueObject) SetQueue(v []QueueObjectQueueInner) {
	o.Queue = v
}

func (o QueueObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentlyPlaying) {
		toSerialize["currently_playing"] = o.CurrentlyPlaying
	}
	if !IsNil(o.Queue) {
		toSerialize["queue"] = o.Queue
	}
	return toSerialize, nil
}

type NullableQueueObject struct {
	value *QueueObject
	isSet bool
}

func (v NullableQueueObject) Get() *QueueObject {
	return v.value
}

func (v *NullableQueueObject) Set(val *QueueObject) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueObject) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueObject(val *QueueObject) *NullableQueueObject {
	return &NullableQueueObject{value: val, isSet: true}
}

func (v NullableQueueObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

