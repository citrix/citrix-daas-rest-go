/*
Citrix.CloudServices.Administrators.Api

Testing PingAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ccadmins

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccadmins"
)

func Test_ccadmins_PingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PingAPIService CustomerPingGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customer string

		resp, httpRes, err := apiClient.PingAPI.CustomerPingGet(context.Background(), customer).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}