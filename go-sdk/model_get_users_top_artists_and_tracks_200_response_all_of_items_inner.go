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

// GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner - struct for GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner
type GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner struct {
	ArtistObject *ArtistObject
	TrackObject *TrackObject
}

// ArtistObjectAsGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner is a convenience function that returns ArtistObject wrapped in GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner
func ArtistObjectAsGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner(v *ArtistObject) GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner {
	return GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner{
		ArtistObject: v,
	}
}

// TrackObjectAsGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner is a convenience function that returns TrackObject wrapped in GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner
func TrackObjectAsGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner(v *TrackObject) GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner {
	return GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner{
		TrackObject: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArtistObject
	err = newStrictDecoder(data).Decode(&dst.ArtistObject)
	if err == nil {
		jsonArtistObject, _ := json.Marshal(dst.ArtistObject)
		if string(jsonArtistObject) == "{}" { // empty struct
			dst.ArtistObject = nil
		} else {
			match++
		}
	} else {
		dst.ArtistObject = nil
	}

	// try to unmarshal data into TrackObject
	err = newStrictDecoder(data).Decode(&dst.TrackObject)
	if err == nil {
		jsonTrackObject, _ := json.Marshal(dst.TrackObject)
		if string(jsonTrackObject) == "{}" { // empty struct
			dst.TrackObject = nil
		} else {
			match++
		}
	} else {
		dst.TrackObject = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArtistObject = nil
		dst.TrackObject = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) MarshalJSON() ([]byte, error) {
	if src.ArtistObject != nil {
		return json.Marshal(&src.ArtistObject)
	}

	if src.TrackObject != nil {
		return json.Marshal(&src.TrackObject)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArtistObject != nil {
		return obj.ArtistObject
	}

	if obj.TrackObject != nil {
		return obj.TrackObject
	}

	// all schemas are nil
	return nil
}

type NullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner struct {
	value *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner
	isSet bool
}

func (v NullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) Get() *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner {
	return v.value
}

func (v *NullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) Set(val *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner(val *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) *NullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner {
	return &NullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner{value: val, isSet: true}
}

func (v NullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


