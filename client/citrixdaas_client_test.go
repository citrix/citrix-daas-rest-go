// Copyright © 2025. Citrix Systems, Inc.

package citrixclient_test

// Tests for utility functions in citrixdaas_client.go

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/citrix/citrix-daas-rest-go/citrixorchestration"
	citrixclient "github.com/citrix/citrix-daas-rest-go/client"
	test "github.com/citrix/citrix-daas-rest-go/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestExecuteWithRetry(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		expectError    bool
		expectedZoneId string
	}{
		{
			name:           "Success on First Attempt",
			statusCode:     200,
			expectError:    false,
			expectedZoneId: "zone-1",
		},
		{
			name:           "Success After Retries",
			statusCode:     200,
			expectError:    false,
			expectedZoneId: "zone-1",
		},
		{
			name:        "Permanent Failure",
			statusCode:  404,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			citrixClient, mockClient := test.NewTestDaaSClient()
			defer mockClient.AssertExpectations(t)

			mockAPI := citrixorchestration.GetMockZonesAPIsDAAS(mockClient.APIClient)
			ctx := test.DaaSTestContext()

			if tt.expectError {
				mockAPI.On("ZonesGetZoneExecute", mock.Anything).
					Return((*citrixorchestration.ZoneDetailResponseModel)(nil), &http.Response{StatusCode: tt.statusCode}, fmt.Errorf("zone not found")).Once()

				request := mockAPI.ZonesGetZone(ctx, "non-existent-zone")
				result, httpResp, err := citrixclient.ExecuteWithRetryContext[*citrixorchestration.ZoneDetailResponseModel](ctx, request, citrixClient)

				assert.Error(t, err)
				assert.Nil(t, result)
				assert.NotNil(t, httpResp)
				assert.Equal(t, tt.statusCode, httpResp.StatusCode)
			} else {
				expectedZone := &citrixorchestration.ZoneDetailResponseModel{
					Id:   tt.expectedZoneId,
					Name: "Test Zone",
				}
				mockAPI.On("ZonesGetZoneExecute", mock.Anything).
					Return(expectedZone, &http.Response{StatusCode: tt.statusCode}, nil).Once()

				request := mockAPI.ZonesGetZone(ctx, tt.expectedZoneId)
				result, httpResp, err := citrixclient.ExecuteWithRetryContext[*citrixorchestration.ZoneDetailResponseModel](ctx, request, citrixClient)

				require.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, tt.statusCode, httpResp.StatusCode)
				assert.Equal(t, tt.expectedZoneId, result.Id)
			}
		})
	}
}

func TestAddRequestData(t *testing.T) {
	mockClient := citrixorchestration.NewMockAPIClient()
	defer mockClient.AssertExpectations(t)
	mockAPI := citrixorchestration.GetMockZonesAPIsDAAS(mockClient.APIClient)

	customerId := "test-customer-123"
	siteId := "test-site-456"
	userAgent := "test-agent"
	authToken := "Bearer test-token"

	client := &citrixclient.CitrixDaasClient{
		ApiClient: mockClient.APIClient,
		ClientConfig: &citrixclient.ClientConfiguration{
			CustomerId:    customerId,
			SiteId:        siteId,
			UserAgent:     userAgent,
			Accept:        "application/json",
			Authorization: authToken,
		},
	}

	request := mockAPI.ZonesGetZone(test.DaaSTestContext(), "zone-1")
	enhancedRequest := citrixclient.AddRequestData(request, client)
	// Wrap in Mock to access getter methods for verification
	mockRequest := citrixorchestration.MockApiZonesGetZoneRequest{ApiZonesGetZoneRequest: enhancedRequest}

	assert.NotNil(t, enhancedRequest)
	assert.Equal(t, "zone-1", mockRequest.GetNameOrId())
	assert.NotNil(t, mockRequest.GetCitrixCustomerId())
	assert.Equal(t, customerId, *mockRequest.GetCitrixCustomerId())
	assert.NotNil(t, mockRequest.GetCitrixInstanceId())
	assert.Equal(t, siteId, *mockRequest.GetCitrixInstanceId())
	assert.NotNil(t, mockRequest.GetUserAgent())
	assert.Equal(t, userAgent, *mockRequest.GetUserAgent())
	assert.NotNil(t, mockRequest.GetAuthorization())
	assert.Equal(t, authToken, *mockRequest.GetAuthorization())
	assert.NotNil(t, mockRequest.GetAccept())
	assert.Equal(t, "application/json", *mockRequest.GetAccept())
}

func TestRetryOperationWithExponentialBackOff_Success(t *testing.T) {
	tests := []struct {
		name              string
		failuresBeforeOk  int
		expectedCallCount int
	}{
		{
			name:              "Success on First Attempt",
			failuresBeforeOk:  0,
			expectedCallCount: 1,
		},
		{
			name:              "Success After Retries",
			failuresBeforeOk:  2,
			expectedCallCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			expectedResult := "success"

			operation := func() (string, *http.Response, error) {
				callCount++
				// Simulate retryable failures (500) before succeeding
				if callCount <= tt.failuresBeforeOk {
					return "", &http.Response{StatusCode: 500, Body: http.NoBody}, fmt.Errorf("temporary error")
				}
				return expectedResult, &http.Response{StatusCode: 200, Body: http.NoBody}, nil
			}

			result, httpResp, err := citrixclient.RetryOperationWithExponentialBackOff(operation, 0, 3)

			require.NoError(t, err)
			assert.Equal(t, expectedResult, result)
			assert.Equal(t, 200, httpResp.StatusCode)
			assert.Equal(t, tt.expectedCallCount, callCount)
		})
	}
}

func TestRetryOperationWithExponentialBackOff_Failure(t *testing.T) {
	tests := []struct {
		name              string
		statusCode        int
		expectedCallCount int
	}{
		{
			name:              "All Retries Fail",
			statusCode:        503,
			expectedCallCount: 3,
		},
		{
			name:              "Non-Retryable Status Code",
			statusCode:        404, // 404 errors should not trigger retries
			expectedCallCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			expectedError := fmt.Errorf("error message")

			operation := func() (string, *http.Response, error) {
				callCount++
				return "", &http.Response{StatusCode: tt.statusCode, Body: http.NoBody}, expectedError
			}

			result, httpResp, err := citrixclient.RetryOperationWithExponentialBackOff(operation, 0, 3)

			assert.Error(t, err)
			assert.Equal(t, expectedError, err)
			assert.Empty(t, result)
			assert.Equal(t, tt.statusCode, httpResp.StatusCode)
			assert.Equal(t, tt.expectedCallCount, callCount)
		})
	}
}

func TestRetryOperationWithExponentialBackOffDefault(t *testing.T) {
	callCount := 0
	expectedResult := 42

	operation := func() (int, *http.Response, error) {
		callCount++
		return expectedResult, &http.Response{StatusCode: 200}, nil
	}

	result, httpResp, err := citrixclient.RetryOperationWithExponentialBackOffDefault(test.DaaSTestContext(), operation)

	require.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, 200, httpResp.StatusCode)
	assert.Equal(t, 1, callCount)
}

func TestGetJobIdFromHttpResponse(t *testing.T) {
	tests := []struct {
		name       string
		location   string
		expectedId string
	}{
		{
			name:       "Extracts Job ID from Location Header",
			location:   "https://api.cloud.com/jobs/12345-67890-abcdef",
			expectedId: "12345-67890-abcdef",
		},
		{
			name:       "Handles Complex Path",
			location:   "https://api.cloud.com/orchestration/api/customer-123/site-456/jobs/job-789",
			expectedId: "job-789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpResponse := http.Response{
				Header: http.Header{
					"Location": []string{tt.location},
				},
			}

			jobId := citrixclient.GetJobIdFromHttpResponse(httpResponse)

			assert.Equal(t, tt.expectedId, jobId)
		})
	}
}

func TestGetTransactionIdFromHttpResponse(t *testing.T) {
	tests := []struct {
		name       string
		response   *http.Response
		expectedId string
	}{
		{
			name: "Extracts Transaction ID from Header",
			response: &http.Response{
				StatusCode: 200,
				Header: http.Header{
					"Citrix-Transactionid": []string{"txn-12345-abcdef"},
				},
			},
			expectedId: "txn-12345-abcdef",
		},
		{
			name: "Returns Empty String When No Transaction ID",
			response: &http.Response{
				StatusCode: 200,
				Header:     http.Header{},
			},
			expectedId: "",
		},
		{
			name:       "Returns Error Message When Response is Nil",
			response:   nil,
			expectedId: "failed before request was sent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txnId := citrixclient.GetTransactionIdFromHttpResponse(tt.response)
			assert.Equal(t, tt.expectedId, txnId)
		})
	}
}

func TestGetBatchRequestItemRelativeUrl(t *testing.T) {
	tests := []struct {
		name        string
		onPremises  bool
		apiGateway  bool
		expectedUrl string
	}{
		{
			name:        "Cloud Customer With API Gateway",
			onPremises:  false,
			apiGateway:  true,
			expectedUrl: "zones/zone-1",
		},
		{
			name:        "On-Premises Customer",
			onPremises:  true,
			apiGateway:  false,
			expectedUrl: "customer-123/site-456/zones/zone-1",
		},
		{
			name:        "Cloud Customer Without API Gateway",
			onPremises:  false,
			apiGateway:  false,
			expectedUrl: "customer-123/site-456/zones/zone-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &citrixclient.CitrixDaasClient{
				AuthConfig: &citrixclient.AuthenticationConfiguration{
					OnPremises: tt.onPremises,
					ApiGateway: tt.apiGateway,
				},
				ClientConfig: &citrixclient.ClientConfiguration{
					CustomerId: "customer-123",
					SiteId:     "site-456",
				},
			}

			relativeUrl := "zones/zone-1"
			result := client.GetBatchRequestItemRelativeUrl(relativeUrl)

			assert.Equal(t, tt.expectedUrl, result)
		})
	}
}

func TestWaitForJob_Success(t *testing.T) {
	tests := []struct {
		name          string
		initialStatus citrixorchestration.JobStatus
		pollCount     int
	}{
		{
			name:          "Job Completes Immediately",
			initialStatus: citrixorchestration.JOBSTATUS_COMPLETE,
			pollCount:     1,
		},
		{
			name:          "Job Completes After Polling",
			initialStatus: citrixorchestration.JOBSTATUS_IN_PROGRESS,
			pollCount:     3,
		},
		{
			name:          "Job Status NOT_STARTED Then Completes",
			initialStatus: citrixorchestration.JOBSTATUS_NOT_STARTED,
			pollCount:     2,
		},
		{
			name:          "Job Status UNKNOWN Then Completes",
			initialStatus: citrixorchestration.JOBSTATUS_UNKNOWN,
			pollCount:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			citrixClient, mockClient := test.NewTestDaaSClient()
			defer mockClient.AssertExpectations(t)

			mockJobsAPI := citrixorchestration.GetMockJobsAPIsDAAS(mockClient.APIClient)
			jobId := "job-456"

			// Simulate multiple polls if pollCount > 1, final poll returns COMPLETE
			if tt.pollCount > 1 {
				inProgressJob := &citrixorchestration.JobResponseModel{
					Id:     jobId,
					Status: tt.initialStatus,
				}
				mockJobsAPI.On("JobsGetJobExecute", mock.Anything).
					Return(inProgressJob, &http.Response{StatusCode: 200}, nil).
					Times(tt.pollCount - 1)
			}

			completedJob := &citrixorchestration.JobResponseModel{
				Id:     jobId,
				Status: citrixorchestration.JOBSTATUS_COMPLETE,
			}
			mockJobsAPI.On("JobsGetJobExecute", mock.Anything).
				Return(completedJob, &http.Response{StatusCode: 200}, nil).Once()

			result, err := citrixClient.WaitForJob(test.DaaSTestContext(), jobId, 5)

			require.NoError(t, err)
			require.NotNil(t, result)
			assert.Equal(t, jobId, result.Id)
			assert.Equal(t, citrixorchestration.JOBSTATUS_COMPLETE, result.Status)
		})
	}
}

func TestWaitForJob_Failure(t *testing.T) {
	t.Run("Job Fails", func(t *testing.T) {
		citrixClient, mockClient := test.NewTestDaaSClient()
		defer mockClient.AssertExpectations(t)

		mockJobsAPI := citrixorchestration.GetMockJobsAPIsDAAS(mockClient.APIClient)

		jobId := "job-789"
		errorString := "Job failed due to an error"
		failedJob := &citrixorchestration.JobResponseModel{
			Id:     jobId,
			Status: citrixorchestration.JOBSTATUS_FAILED,
		}
		failedJob.ErrorString.Set(&errorString)

		mockJobsAPI.On("JobsGetJobExecute", mock.Anything).Return(failedJob, &http.Response{StatusCode: 200}, nil).Once()

		result, err := citrixClient.WaitForJob(test.DaaSTestContext(), jobId, 5)

		// WaitForJob returns the job model even when job status is FAILED (no error returned)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, jobId, result.Id)
		assert.Equal(t, citrixorchestration.JOBSTATUS_FAILED, result.Status)
		assert.Equal(t, errorString, *result.ErrorString.Get())
	})

	t.Run("Job Poll Returns Error", func(t *testing.T) {
		citrixClient, mockClient := test.NewTestDaaSClient()
		defer mockClient.AssertExpectations(t)

		mockJobsAPI := citrixorchestration.GetMockJobsAPIsDAAS(mockClient.APIClient)

		jobId := "job-error"
		expectedError := fmt.Errorf("failed to get job status")

		mockJobsAPI.On("JobsGetJobExecute", mock.Anything).Return((*citrixorchestration.JobResponseModel)(nil), &http.Response{StatusCode: 500}, expectedError).Once()

		result, err := citrixClient.WaitForJob(test.DaaSTestContext(), jobId, 5)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, result)
	})

	t.Run("Job Times Out", func(t *testing.T) {
		citrixClient, mockClient := test.NewTestDaaSClient()
		defer mockClient.AssertExpectations(t)

		mockJobsAPI := citrixorchestration.GetMockJobsAPIsDAAS(mockClient.APIClient)

		jobId := "job-timeout"
		inProgressJob := &citrixorchestration.JobResponseModel{
			Id:     jobId,
			Status: citrixorchestration.JOBSTATUS_IN_PROGRESS,
		}

		mockJobsAPI.On("JobsGetJobExecute", mock.Anything).Return(inProgressJob, &http.Response{StatusCode: 200}, nil).Maybe()

		cancelledJob := &citrixorchestration.JobResponseModel{
			Id: jobId,
		}
		mockJobsAPI.On("JobsCancelJobExecute", mock.Anything).Return(cancelledJob, &http.Response{StatusCode: 200}, nil).Maybe()

		// Timeout of 0 causes immediate timeout
		result, err := citrixClient.WaitForJob(test.DaaSTestContext(), jobId, 0)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "timed out")
		assert.NotNil(t, result)
		assert.Equal(t, jobId, result.Id)
	})
}
