/*
Spotify Web API

Testing CategoriesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/viqueen/spotify-api"
)

func Test_openapi_CategoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CategoriesAPIService GetACategoriesPlaylists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string

		resp, httpRes, err := apiClient.CategoriesAPI.GetACategoriesPlaylists(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService GetACategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string

		resp, httpRes, err := apiClient.CategoriesAPI.GetACategory(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService GetCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CategoriesAPI.GetCategories(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}