/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing AppVIsolationGroupsAPIsDAASService

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

func Test_citrixorchestration_AppVIsolationGroupsAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppVIsolationGroupsAPIsDAASService AppVIsolationGroupsCreateAppVIsolationGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsCreateAppVIsolationGroup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppVIsolationGroupsAPIsDAASService AppVIsolationGroupsDeleteAppVIsolationGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		httpRes, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsDeleteAppVIsolationGroup(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppVIsolationGroupsAPIsDAASService AppVIsolationGroupsGetAppVIsolationGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroup(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppVIsolationGroupsAPIsDAASService AppVIsolationGroupsGetAppVIsolationGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppVIsolationGroupsAPIsDAASService AppVIsolationGroupsUpdateAppVIsolationGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		httpRes, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsUpdateAppVIsolationGroup(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
