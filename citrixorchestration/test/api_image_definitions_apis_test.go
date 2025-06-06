/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing ImageDefinitionsAPIsDAASService

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

func Test_citrixorchestration_ImageDefinitionsAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsCheckImageDefinitionExist", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsCheckImageDefinitionExist(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsCreateImageDefinition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsCreateImageDefinition(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsCreateImageVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsCreateImageVersion(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsDeleteImageDefinition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsDeleteImageDefinition(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsDeleteImageDefinitionImageVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string
		var versionNumberOrId string

		httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsDeleteImageDefinitionImageVersion(context.Background(), nameOrId, versionNumberOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsDoImageDefinitionAndImageVersionSearch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsDoImageDefinitionAndImageVersionSearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsGetImageDefinition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinition(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsGetImageDefinitionImageVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string
		var versionNumberOrId string

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitionImageVersion(context.Background(), nameOrId, versionNumberOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsGetImageDefinitionImageVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitionImageVersions(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsGetImageDefinitions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsGetImageVersionProvisioningSchemes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string
		var versionNumberOrId string

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageVersionProvisioningSchemes(context.Background(), nameOrId, versionNumberOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsSetImageDefinitionImageVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string
		var versionNumberOrId string

		httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsSetImageDefinitionImageVersion(context.Background(), nameOrId, versionNumberOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsUpdateImageDefinition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageDefinition(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsUpdateImageDefinitionImageVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string
		var versionNumberOrId string

		resp, httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageDefinitionImageVersion(context.Background(), nameOrId, versionNumberOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageDefinitionsAPIsDAASService ImageDefinitionsUpdateImageVersionResourcePools", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string
		var versionNumberOrId string

		httpRes, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageVersionResourcePools(context.Background(), nameOrId, versionNumberOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
