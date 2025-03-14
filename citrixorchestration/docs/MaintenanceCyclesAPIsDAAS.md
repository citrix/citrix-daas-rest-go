# \MaintenanceCyclesAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaintenanceCyclesCancelMaintenanceCycle**](MaintenanceCyclesAPIsDAAS.md#MaintenanceCyclesCancelMaintenanceCycle) | **Post** /maintenanceCycles/{nameOrId}/$cancel | Cancel a maintenance cycle for a machine catalog.
[**MaintenanceCyclesCreateMaintenanceCycle**](MaintenanceCyclesAPIsDAAS.md#MaintenanceCyclesCreateMaintenanceCycle) | **Post** /maintenanceCycles | Create a maintenance cycle for a machine catalog.
[**MaintenanceCyclesDoMaintenanceCycleSearch**](MaintenanceCyclesAPIsDAAS.md#MaintenanceCyclesDoMaintenanceCycleSearch) | **Post** /maintenanceCycles/$search | Perform an advanced search for maintenance cycles.
[**MaintenanceCyclesDoMaintenanceCycleVMOperationJobSearch**](MaintenanceCyclesAPIsDAAS.md#MaintenanceCyclesDoMaintenanceCycleVMOperationJobSearch) | **Post** /maintenanceCycles/operationJobs/$search | Perform an advanced search for maintenance cycle vm operation jobs.
[**MaintenanceCyclesGetMaintenanceCycle**](MaintenanceCyclesAPIsDAAS.md#MaintenanceCyclesGetMaintenanceCycle) | **Get** /maintenanceCycles/{nameOrId} | Get a maintenance cycle for a machine catalog.
[**MaintenanceCyclesRemoveMaintenanceCycle**](MaintenanceCyclesAPIsDAAS.md#MaintenanceCyclesRemoveMaintenanceCycle) | **Delete** /maintenanceCycles/{nameOrId} | Remove a maintenance cycle for a machine catalog.
[**MaintenanceCyclesRestartMaintenanceCycle**](MaintenanceCyclesAPIsDAAS.md#MaintenanceCyclesRestartMaintenanceCycle) | **Post** /maintenanceCycles/{nameOrId}/$restart | Restart a maintenance cycle for a machine catalog.
[**MaintenanceCyclesUpdateMaintenanceCycle**](MaintenanceCyclesAPIsDAAS.md#MaintenanceCyclesUpdateMaintenanceCycle) | **Patch** /maintenanceCycles | Update a maintenance cycle for a machine catalog.



## MaintenanceCyclesCancelMaintenanceCycle

> MaintenanceCyclesCancelMaintenanceCycle(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Cancel a maintenance cycle for a machine catalog.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the maintenance cycle will be removed within a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MaintenanceCyclesAPIsDAAS.MaintenanceCyclesCancelMaintenanceCycle(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesCancelMaintenanceCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceCyclesCancelMaintenanceCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the maintenance cycle will be removed within a background task. | [default to false]

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceCyclesCreateMaintenanceCycle

> MaintenanceCyclesCreateMaintenanceCycle(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateMaintenanceCycleRequestModel(createMaintenanceCycleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Create a maintenance cycle for a machine catalog.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    createMaintenanceCycleRequestModel := *openapiclient.NewCreateMaintenanceCycleRequestModel(int32(123), openapiclient.MaintenanceCycleOperation2("Unknown")) // CreateMaintenanceCycleRequestModel | Details to create the maintenance cycle.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the maintenance cycle will be create within a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MaintenanceCyclesAPIsDAAS.MaintenanceCyclesCreateMaintenanceCycle(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateMaintenanceCycleRequestModel(createMaintenanceCycleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesCreateMaintenanceCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceCyclesCreateMaintenanceCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createMaintenanceCycleRequestModel** | [**CreateMaintenanceCycleRequestModel**](CreateMaintenanceCycleRequestModel.md) | Details to create the maintenance cycle. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the maintenance cycle will be create within a background task. | [default to false]

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceCyclesDoMaintenanceCycleSearch

> MaintenanceCycleResponseModelCollection MaintenanceCyclesDoMaintenanceCycleSearch(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).MaintenanceCycleSearchRequestModel(maintenanceCycleSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).Skip(skip).ContinuationToken(continuationToken).Fields(fields).Execute()

Perform an advanced search for maintenance cycles.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    maintenanceCycleSearchRequestModel := *openapiclient.NewMaintenanceCycleSearchRequestModel() // MaintenanceCycleSearchRequestModel | Specifies the advanced search parameters.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the maintenance cycles (and associated objects) will be added as a background task. The task will have JobType SearchMaintenanceCycles. (optional) (default to false)
    limit := int32(56) // int32 | The max number of maintenance cycles returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    skip := int32(56) // int32 |  (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceCyclesAPIsDAAS.MaintenanceCyclesDoMaintenanceCycleSearch(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).MaintenanceCycleSearchRequestModel(maintenanceCycleSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).Skip(skip).ContinuationToken(continuationToken).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesDoMaintenanceCycleSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MaintenanceCyclesDoMaintenanceCycleSearch`: MaintenanceCycleResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesDoMaintenanceCycleSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceCyclesDoMaintenanceCycleSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **maintenanceCycleSearchRequestModel** | [**MaintenanceCycleSearchRequestModel**](MaintenanceCycleSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the maintenance cycles (and associated objects) will be added as a background task. The task will have JobType SearchMaintenanceCycles. | [default to false]
 **limit** | **int32** | The max number of maintenance cycles returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **skip** | **int32** |  | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 

### Return type

[**MaintenanceCycleResponseModelCollection**](MaintenanceCycleResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceCyclesDoMaintenanceCycleVMOperationJobSearch

> MaintenanceCycleVMOperationJobResponseModelCollection MaintenanceCyclesDoMaintenanceCycleVMOperationJobSearch(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).MaintenanceCycleVMOperationJobSearchRequestModel(maintenanceCycleVMOperationJobSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).Skip(skip).ContinuationToken(continuationToken).Fields(fields).Execute()

Perform an advanced search for maintenance cycle vm operation jobs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    maintenanceCycleVMOperationJobSearchRequestModel := *openapiclient.NewMaintenanceCycleVMOperationJobSearchRequestModel() // MaintenanceCycleVMOperationJobSearchRequestModel | Specifies the advanced search parameters.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the maintenance cycle vm operation jobs (and associated objects) will be added as a background task. The task will have JobType SearchMaintenanceCycleVMOperationJobs. (optional) (default to false)
    limit := int32(56) // int32 | The max number of maintenance cycle vm operation jobs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    skip := int32(56) // int32 |  (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceCyclesAPIsDAAS.MaintenanceCyclesDoMaintenanceCycleVMOperationJobSearch(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).MaintenanceCycleVMOperationJobSearchRequestModel(maintenanceCycleVMOperationJobSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).Skip(skip).ContinuationToken(continuationToken).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesDoMaintenanceCycleVMOperationJobSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MaintenanceCyclesDoMaintenanceCycleVMOperationJobSearch`: MaintenanceCycleVMOperationJobResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesDoMaintenanceCycleVMOperationJobSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceCyclesDoMaintenanceCycleVMOperationJobSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **maintenanceCycleVMOperationJobSearchRequestModel** | [**MaintenanceCycleVMOperationJobSearchRequestModel**](MaintenanceCycleVMOperationJobSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the maintenance cycle vm operation jobs (and associated objects) will be added as a background task. The task will have JobType SearchMaintenanceCycleVMOperationJobs. | [default to false]
 **limit** | **int32** | The max number of maintenance cycle vm operation jobs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **skip** | **int32** |  | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 

### Return type

[**MaintenanceCycleVMOperationJobResponseModelCollection**](MaintenanceCycleVMOperationJobResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceCyclesGetMaintenanceCycle

> MaintenanceCycleDetailsResponseModel MaintenanceCyclesGetMaintenanceCycle(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get a maintenance cycle for a machine catalog.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | The Id or name of the maintenance cycle to get. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the maintenance cycle will be get within a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceCyclesAPIsDAAS.MaintenanceCyclesGetMaintenanceCycle(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesGetMaintenanceCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MaintenanceCyclesGetMaintenanceCycle`: MaintenanceCycleDetailsResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesGetMaintenanceCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The Id or name of the maintenance cycle to get.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceCyclesGetMaintenanceCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the maintenance cycle will be get within a background task. | [default to false]

### Return type

[**MaintenanceCycleDetailsResponseModel**](MaintenanceCycleDetailsResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceCyclesRemoveMaintenanceCycle

> MaintenanceCyclesRemoveMaintenanceCycle(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Remove a maintenance cycle for a machine catalog.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the maintenance cycle will be removed within a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MaintenanceCyclesAPIsDAAS.MaintenanceCyclesRemoveMaintenanceCycle(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesRemoveMaintenanceCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceCyclesRemoveMaintenanceCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the maintenance cycle will be removed within a background task. | [default to false]

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceCyclesRestartMaintenanceCycle

> MaintenanceCyclesRestartMaintenanceCycle(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RestartMaintenanceCycleRequestModel(restartMaintenanceCycleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Restart a maintenance cycle for a machine catalog.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | The Id or name of the maintenance cycle to restart. 
    restartMaintenanceCycleRequestModel := *openapiclient.NewRestartMaintenanceCycleRequestModel(time.Now(), int32(123)) // RestartMaintenanceCycleRequestModel | Details to restart the maintenance cycle.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the maintenance cycle will be restarted within a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MaintenanceCyclesAPIsDAAS.MaintenanceCyclesRestartMaintenanceCycle(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RestartMaintenanceCycleRequestModel(restartMaintenanceCycleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesRestartMaintenanceCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The Id or name of the maintenance cycle to restart.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceCyclesRestartMaintenanceCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **restartMaintenanceCycleRequestModel** | [**RestartMaintenanceCycleRequestModel**](RestartMaintenanceCycleRequestModel.md) | Details to restart the maintenance cycle. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the maintenance cycle will be restarted within a background task. | [default to false]

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceCyclesUpdateMaintenanceCycle

> MaintenanceCyclesUpdateMaintenanceCycle(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMaintenanceCycleRequestModel(updateMaintenanceCycleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Update a maintenance cycle for a machine catalog.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    updateMaintenanceCycleRequestModel := *openapiclient.NewUpdateMaintenanceCycleRequestModel("MaintenanceCycleUid_example", int32(123)) // UpdateMaintenanceCycleRequestModel | Details to create the maintenance cycle.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the maintenance cycle will be create within a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MaintenanceCyclesAPIsDAAS.MaintenanceCyclesUpdateMaintenanceCycle(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMaintenanceCycleRequestModel(updateMaintenanceCycleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceCyclesAPIsDAAS.MaintenanceCyclesUpdateMaintenanceCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceCyclesUpdateMaintenanceCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **updateMaintenanceCycleRequestModel** | [**UpdateMaintenanceCycleRequestModel**](UpdateMaintenanceCycleRequestModel.md) | Details to create the maintenance cycle. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the maintenance cycle will be create within a background task. | [default to false]

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

