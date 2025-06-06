/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing ActionsAPIsDAASService

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

func Test_citrixorchestration_ActionsAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ActionsAPIsDAASService ActionsCancelAction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var catalogNameOrId string
		var actionId string

		httpRes, err := apiClient.ActionsAPIsDAAS.ActionsCancelAction(context.Background(), catalogNameOrId, actionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionsAPIsDAASService ActionsDeleteAction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var catalogNameOrId string

		httpRes, err := apiClient.ActionsAPIsDAAS.ActionsDeleteAction(context.Background(), catalogNameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionsAPIsDAASService ActionsDeleteActions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ActionsAPIsDAAS.ActionsDeleteActions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionsAPIsDAASService ActionsGetAction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var catalogNameOrId string

		resp, httpRes, err := apiClient.ActionsAPIsDAAS.ActionsGetAction(context.Background(), catalogNameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionsAPIsDAASService ActionsGetActionById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var catalogNameOrId string
		var actionId string

		resp, httpRes, err := apiClient.ActionsAPIsDAAS.ActionsGetActionById(context.Background(), catalogNameOrId, actionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionsAPIsDAASService ActionsGetActions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ActionsAPIsDAAS.ActionsGetActions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
