/*
Global App Config Admin

Testing SettingsControllerDAASService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package globalappconfiguration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func Test_globalappconfiguration_SettingsControllerDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SettingsControllerDAASService DeleteSettingsApiUsingDELETE", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var url string

		httpRes, err := apiClient.SettingsControllerDAAS.DeleteSettingsApiUsingDELETE(context.Background(), app, url).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsControllerDAASService DeleteSettingsForChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var channelName string
		var url string

		httpRes, err := apiClient.SettingsControllerDAAS.DeleteSettingsForChannel(context.Background(), app, channelName, url).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsControllerDAASService GetAllSettingsApiUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SettingsControllerDAAS.GetAllSettingsApiUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsControllerDAASService GetSettingsApiUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var url string

		resp, httpRes, err := apiClient.SettingsControllerDAAS.GetSettingsApiUsingGET(context.Background(), app, url).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsControllerDAASService PostSettingsApiUsingPOST", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string

		resp, httpRes, err := apiClient.SettingsControllerDAAS.PostSettingsApiUsingPOST(context.Background(), app).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsControllerDAASService PutSettingsApiUsingPUT", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var url string

		httpRes, err := apiClient.SettingsControllerDAAS.PutSettingsApiUsingPUT(context.Background(), app, url).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsControllerDAASService RetrieveAllChannelSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var url string

		resp, httpRes, err := apiClient.SettingsControllerDAAS.RetrieveAllChannelSettings(context.Background(), app, url).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsControllerDAASService RetrieveSettingsForChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var channelName string
		var url string

		resp, httpRes, err := apiClient.SettingsControllerDAAS.RetrieveSettingsForChannel(context.Background(), app, channelName, url).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
