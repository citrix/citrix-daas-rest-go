# Unit Testing with Mock Clients

This guide explains how to write unit tests for the Citrix DaaS REST Go client using the auto-generated mock implementations.

## Overview

The OpenAPI code generator creates mock implementations for all API services, making it easy to write comprehensive unit tests without requiring actual API connections. Each generated API service includes:

- A mock implementation of the service interface
- Mock request wrapper structs with getter methods for argument matching
- Convenience setup functions to streamline test creation

## Quick Start

### Basic Test Setup

```go
import (
    "context"
    "net/http"
    "testing"
    
    "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
)

// Function under test that uses the API client
func GetZoneById(api citrixorchestration.ZonesAPIsDAAS, zoneId string) (*citrixorchestration.ZoneDetailResponseModel, error) {
    resp, _, err := api.ZonesGetZone(context.Background(), zoneId).Execute()
    return resp, err
}

func TestZoneOperation(t *testing.T) {
    // Setup: Create mock client and extract the specific API mock
    mockClient := citrixorchestration.NewMockAPIClient()
    defer mockClient.AssertExpectations(t)
    mockAPI := citrixorchestration.GetMockZonesAPIsDAAS(mockClient.APIClient)
    
    // Configure: Set expectations on the mock
    mockZone := &citrixorchestration.ZoneDetailResponseModel{
        Id:        "test-zone",
        Name:      "Test Zone",
        IsPrimary: true,
    }
    mockAPI.On("ZonesGetZoneExecute", mock.Anything).
        Return(mockZone, &http.Response{StatusCode: 200}, nil)
    
    // Execute: Call the function under test
    resp, err := GetZoneById(mockAPI, "test-zone")
    
    // Assert: Verify results
    require.Nil(t, err)
    require.NotNil(t, resp)
    assert.Equal(t, "test-zone", resp.Id)
    assert.Equal(t, "Test Zone", resp.Name)
}
```

## Key Patterns

### 1. Success Scenarios

Test successful API responses by setting up mock expectations with valid return values.

```go
mockAPI.On("ZonesGetZoneExecute", mock.Anything).
    Return(mockZone, &http.Response{StatusCode: 200}, nil)
```

See `Test ZonesGetZone Success` in [example_test.go](./example_test.go) for details.

### 2. Error Handling

Verify your code properly handles API errors.

```go
mockAPI.On("ZonesGetZoneExecute", mock.Anything).
    Return(nil, &http.Response{StatusCode: 500}, expectedError)
```

See `Test ZonesGetZone Error` in [example_test.go](./example_test.go) for details.

### 3. Sequential Calls

Test scenarios requiring multiple calls with different responses using `.Once()`.

```go
mockAPI.On("ZonesGetZoneExecute", mock.Anything).
    Return(successResponse, &http.Response{StatusCode: 200}, nil).Once()
mockAPI.On("ZonesGetZoneExecute", mock.Anything).
    Return(nil, &http.Response{StatusCode: 500}, error).Once()
```

See `Test ZonesGetZone Multiple Calls Same Zone` in [example_test.go](./example_test.go) for details.

### 4. Argument Matching

Use `mock.MatchedBy()` with mock request wrapper structs to match based on specific request parameters.

```go
mockAPI.On("ZonesGetZoneExecute", mock.MatchedBy(func(req citrixorchestration.ApiZonesGetZoneRequest) bool {
    mockReq := citrixorchestration.MockApiZonesGetZoneRequest{ApiZonesGetZoneRequest: req}
    return mockReq.GetNameOrId() == "zone-1"
})).Return(zone1Response, &http.Response{StatusCode: 200}, nil)
```

The mock wrapper structs provide getter methods for all request fields, enabling precise argument matching.

See `Test ZonesGetZone Different Responses By Input` in [example_test.go](./example_test.go) for details.

### 5. Call Count Verification

Verify methods are called the expected number of times.

```go
mockAPI.On("ZonesGetZoneExecute", mock.Anything).
    Return(mockZone, &http.Response{StatusCode: 200}, nil).Times(3)

// ... make 3 calls ...

mockAPI.AssertNumberOfCalls(t, "ZonesGetZoneExecute", 3)
```

See `Test ZonesGetZone Verify Call Count` in [example_test.go](./example_test.go) for details.

### 6. Retry Logic

Simulate transient failures and successful retries.

```go
// First attempt: failure
mockAPI.On("ZonesGetZoneExecute", mock.Anything).
    Return(nil, &http.Response{StatusCode: 503}, error).Once()

// Second attempt: success
mockAPI.On("ZonesGetZoneExecute", mock.Anything).
    Return(mockZone, &http.Response{StatusCode: 200}, nil).Once()
```

See `Test ZonesGetZone Retry Logic Simulation` in [example_test.go](./example_test.go) for details.

### 7. Custom Assertions

Use `.Run()` to execute custom logic during mock calls.

```go
callCount := 0
mockAPI.On("ZonesGetZoneExecute", mock.Anything).
    Run(func(args mock.Arguments) {
        callCount++
        req := args.Get(0).(citrixorchestration.ApiZonesGetZoneRequest)
        // Perform custom validations on req
    }).
    Return(mockZone, &http.Response{StatusCode: 200}, nil)
```

See `Test ZonesGetZone With Custom Assertion` in [example_test.go](./example_test.go) for details.

### 8. Optional Calls

Use `.Maybe()` for expectations that may or may not be called.

```go
mockAPI.On("ZonesGetZoneExecute", mock.Anything).
    Return(mockZone, &http.Response{StatusCode: 200}, nil).Maybe()
```

See `Test ZonesGetZone Without Specific Expectations` in [example_test.go](./example_test.go) for details.

## Available Mock Helpers

Each API service has a convenience getter function to extract the mock from the APIClient:

- `citrixorchestration.GetMockAboutAPIsDAAS(client)`
- `citrixorchestration.GetMockZonesAPIsDAAS(client)`
- `citrixorchestration.GetMockJobsAPIsDAAS(client)`
- ... and more for all API services

These functions take an `*APIClient` and return the typed mock service, eliminating verbose type assertions.

**Pattern:**
```go
mockClient := citrixorchestration.NewMockAPIClient()
mockZonesAPI := citrixorchestration.GetMockZonesAPIsDAAS(mockClient.APIClient)
```

**Instead of:**
```go
mockZonesAPI := mockClient.APIClient.ZonesAPIsDAAS.(*citrixorchestration.MockZonesAPIsDAAS)
```

## Test Helper Functions

### NewTestDaaSClient

When testing functions that require a full `CitrixDaasClient` instance (like `ExecuteWithRetry` or `WaitForJob`), use the `citrixclient.NewTestDaaSClient()` helper to eliminate boilerplate. This helper automatically creates all necessary mocks internally.

**Example usage:**

```go
import citrixclient "github.com/citrix/citrix-daas-rest-go/client"

func TestExecuteWithRetry(t *testing.T) {
    // Setup - creates a fully configured test client with all mocks!
    citrixClient, mockClient := citrixclient.NewTestDaaSClient()
    defer mockClient.AssertExpectations(t)
    
    // Access the specific mock API you need using GetMock helper
    mockZonesAPI := citrixorchestration.GetMockZonesAPIsDAAS(mockClient.APIClient)
    
    // Configure mock expectations
    mockZonesAPI.On("ZonesGetZoneExecute", mock.Anything).
        Return(expectedZone, &http.Response{StatusCode: 200}, nil).Once()
    
    // Execute with retry logic
    request := mockZonesAPI.ZonesGetZone(context.Background(), "zone-1")
    result, httpResp, err := citrixclient.ExecuteWithRetry[*citrixorchestration.ZoneDetailResponseModel](request, citrixClient)
    
    // Assert
    require.NoError(t, err)
    assert.Equal(t, "zone-1", result.Id)
}
```

**NewTestDaaSClient signature:**
```go
func NewTestDaaSClient() (*citrixclient.CitrixDaasClient, *citrixorchestration.MockAPIClient)
```

This helper creates:
- A fully configured `CitrixDaasClient` with all mock API services initialized
- A `MockAPIClient` with all mock services ready to use
    - Use `GetMock*` helpers to access specific mocks, e.g.
    - `mockZonesAPI := citrixorchestration.GetMockZonesAPIsDAAS(mockClient.APIClient)`
- Default test configuration (CustomerId: "test-customer", SiteId: "test-site")
- Default test auth token

**For tests requiring custom configuration**, create the client manually instead:

```go
client := &citrixclient.CitrixDaasClient{
    ApiClient: mockClient.APIClient,
    ClientConfig: &citrixclient.ClientConfiguration{
        CustomerId: "custom-id",
        SiteId:     "custom-site",
        // ... other custom fields
    },
}
```

See `TestExecuteWithRetry` and `TestWaitForJob` in [citrixdaas_client_test.go](../client/citrixdaas_client_test.go) for complete usage examples.

## Mock Request Wrappers

For argument matching, mock wrapper structs are available for all request types:

- `citrixorchestration.MockApiZonesGetZoneRequest`
- `citrixorchestration.MockApiZonesCreateZoneRequest`
- ... and more for all API operations

Usage:
```go
mockRequest := citrixorchestration.MockApiZonesGetZoneRequest{ApiZonesGetZoneRequest: request}
```

Each wrapper provides getter methods for all request parameters, including:
- Path parameters (e.g., `GetNameOrId()`)
- Header parameters (e.g., `GetCitrixCustomerId()`, `GetAuthorization()`)
- Context (e.g., `GetCtx()`)

## Best Practices

1. **Use `defer mockClient.AssertExpectations(t)`** immediately after setup to ensure all mock expectations are verified, even if the test fails early. Use the client-level assertion (not the individual API service) in case the test later expands to use other services.

2. **Call the setup function in each test** (e.g., `SetupMockZonesAPIsDAAS()`) to create a fresh mock instance and avoid interference when multiple tests run together.

3. **Use argument matching** when you need different responses for different inputs in the same test.

4. **Keep tests focused** - each test should verify one specific behavior or scenario.

## Additional Resources

- Complete examples: [example_test.go](./example_test.go)
- Testify documentation: https://github.com/stretchr/testify
- Mock package documentation: https://pkg.go.dev/github.com/stretchr/testify/mock
