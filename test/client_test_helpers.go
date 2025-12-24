// Copyright © 2025. Citrix Systems, Inc.

package test

import (
	"context"

	"github.com/citrix/citrix-daas-rest-go/citrixorchestration"
	citrixclient "github.com/citrix/citrix-daas-rest-go/client"
)

// DaaSTestContext creates a context with zero-delay retry configuration for fast test execution.
// This is intended for use in unit tests to avoid long retry delays.
func DaaSTestContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, citrixclient.RetryBaseDelayKey, 0)
	ctx = context.WithValue(ctx, citrixclient.MaxRetriesKey, 3)
	ctx = context.WithValue(ctx, citrixclient.JobPollIntervalSecondsKey, 0)
	ctx = context.WithValue(ctx, citrixclient.JobLongPollIntervalSecondsKey, 0)
	return ctx
}

// NewTestDaaSClient creates a CitrixDaasClient with a mock APIClient for testing.
// Returns the client and the mock client (for assertion) as a tuple.
//
// Example usage:
//
//	citrixClient, mockClient := test.NewTestDaaSClient()
//	defer mockClient.AssertExpectations(t)
//
//	// Access specific mock APIs as needed
//	mockJobsAPI := citrixorchestration.GetMockJobsAPIsDAAS(mockClient.APIClient)
//	mockJobsAPI.On("JobsGetJobExecute", mock.Anything).Return(...)
func NewTestDaaSClient() (*citrixclient.CitrixDaasClient, *citrixorchestration.MockAPIClient) {
	mockClient := citrixorchestration.NewMockAPIClient()

	citrixClient := &citrixclient.CitrixDaasClient{
		ApiClient:    mockClient.APIClient,
		ClientConfig: &citrixclient.ClientConfiguration{CustomerId: "test-customer", SiteId: "test-site"},
		AuthToken:    &citrixclient.AuthTokenModel{Token: "test-token"},
	}

	return citrixClient, mockClient
}
