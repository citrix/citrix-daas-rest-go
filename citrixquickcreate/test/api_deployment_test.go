/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

Testing DeploymentQCSService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package citrixquickcreate

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickcreate"
)

func Test_citrixquickcreate_DeploymentQCSService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeploymentQCSService AddMachineAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string

		resp, httpRes, err := apiClient.DeploymentQCS.AddMachineAsync(context.Background(), customerId, deploymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService DeleteDeploymentAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string

		resp, httpRes, err := apiClient.DeploymentQCS.DeleteDeploymentAsync(context.Background(), customerId, deploymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService GetDeploymentAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string

		resp, httpRes, err := apiClient.DeploymentQCS.GetDeploymentAsync(context.Background(), customerId, deploymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService GetDeploymentsAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.DeploymentQCS.GetDeploymentsAsync(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService InitiateDeleteDeploymentAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string

		resp, httpRes, err := apiClient.DeploymentQCS.InitiateDeleteDeploymentAsync(context.Background(), customerId, deploymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService InitiateDeploymentAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.DeploymentQCS.InitiateDeploymentAsync(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService InitiateRemoveMachineAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string
		var machineId string

		resp, httpRes, err := apiClient.DeploymentQCS.InitiateRemoveMachineAsync(context.Background(), customerId, deploymentId, machineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService PatchMachinesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string

		httpRes, err := apiClient.DeploymentQCS.PatchMachinesAsync(context.Background(), customerId, deploymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService RemoveMachineAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string
		var machineId string

		resp, httpRes, err := apiClient.DeploymentQCS.RemoveMachineAsync(context.Background(), customerId, deploymentId, machineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService RemoveMachinesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string

		resp, httpRes, err := apiClient.DeploymentQCS.RemoveMachinesAsync(context.Background(), customerId, deploymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService RestartMachineAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string
		var machineId string

		httpRes, err := apiClient.DeploymentQCS.RestartMachineAsync(context.Background(), customerId, deploymentId, machineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService SaveAsImageAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string
		var machineId string

		resp, httpRes, err := apiClient.DeploymentQCS.SaveAsImageAsync(context.Background(), customerId, deploymentId, machineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService UpdateDeploymentImageAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string

		resp, httpRes, err := apiClient.DeploymentQCS.UpdateDeploymentImageAsync(context.Background(), customerId, deploymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService UpdateDeploymentPropertiesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string

		resp, httpRes, err := apiClient.DeploymentQCS.UpdateDeploymentPropertiesAsync(context.Background(), customerId, deploymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentQCSService UpdateMachineAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var deploymentId string
		var machineId string

		resp, httpRes, err := apiClient.DeploymentQCS.UpdateMachineAsync(context.Background(), customerId, deploymentId, machineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}