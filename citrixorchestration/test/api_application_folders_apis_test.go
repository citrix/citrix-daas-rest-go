/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing ApplicationFoldersAPIsDAASService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package citrixorchestration

import (
	"context"
	"testing"

	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_citrixorchestration_ApplicationFoldersAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationFoldersAPIsDAASService ApplicationFoldersCheckApplicationFolderPathExists", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var path string

		httpRes, err := apiClient.ApplicationFoldersAPIsDAAS.ApplicationFoldersCheckApplicationFolderPathExists(context.Background(), path).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFoldersAPIsDAASService ApplicationFoldersCreateApplicationFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationFoldersAPIsDAAS.ApplicationFoldersCreateApplicationFolder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFoldersAPIsDAASService ApplicationFoldersDeleteApplicationFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var pathOrId string

		httpRes, err := apiClient.ApplicationFoldersAPIsDAAS.ApplicationFoldersDeleteApplicationFolder(context.Background(), pathOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFoldersAPIsDAASService ApplicationFoldersGetApplicationFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var pathOrId string

		resp, httpRes, err := apiClient.ApplicationFoldersAPIsDAAS.ApplicationFoldersGetApplicationFolder(context.Background(), pathOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFoldersAPIsDAASService ApplicationFoldersGetApplicationFolderApplications", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var pathOrId string

		resp, httpRes, err := apiClient.ApplicationFoldersAPIsDAAS.ApplicationFoldersGetApplicationFolderApplications(context.Background(), pathOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFoldersAPIsDAASService ApplicationFoldersGetApplicationFolders", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationFoldersAPIsDAAS.ApplicationFoldersGetApplicationFolders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFoldersAPIsDAASService ApplicationFoldersUpdateApplicationFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var pathOrId string

		resp, httpRes, err := apiClient.ApplicationFoldersAPIsDAAS.ApplicationFoldersUpdateApplicationFolder(context.Background(), pathOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
