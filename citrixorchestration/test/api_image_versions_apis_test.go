/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing ImageVersionsAPIsDAASService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package citrixorchestration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func Test_citrixorchestration_ImageVersionsAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImageVersionsAPIsDAASService ImageVersionsUpdateImageVersionResourcePools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ImageVersionsAPIsDAAS.ImageVersionsUpdateImageVersionResourcePools(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}