# \ConfigLogTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigLogTPDeleteLogs**](ConfigLogTPApi.md#ConfigLogTPDeleteLogs) | **Delete** /techpreview/{customerid}/{siteid}/ConfigLog/Logs | Delete logging event logs. This on-prem only API.
[**ConfigLogTPDoConfigLogSearch**](ConfigLogTPApi.md#ConfigLogTPDoConfigLogSearch) | **Post** /techpreview/{customerid}/{siteid}/ConfigLog/Operations/$search | Perform an advanced search for configuration log entries.
[**ConfigLogTPExportReportData**](ConfigLogTPApi.md#ConfigLogTPExportReportData) | **Post** /techpreview/{customerid}/{siteid}/ConfigLog/ExportReport | Export logging report data in the format of csv/html/both.
[**ConfigLogTPGetFirstLogDate**](ConfigLogTPApi.md#ConfigLogTPGetFirstLogDate) | **Get** /techpreview/{customerid}/{siteid}/ConfigLog/GetFirstLogDate | Get first log date
[**ConfigLogTPGetLogSite**](ConfigLogTPApi.md#ConfigLogTPGetLogSite) | **Get** /techpreview/{customerid}/{siteid}/ConfigLog/LoggingSite | Get logging site details. This on-prem only API.
[**ConfigLogTPGetLowLevelOperations**](ConfigLogTPApi.md#ConfigLogTPGetLowLevelOperations) | **Get** /techpreview/{customerid}/{siteid}/ConfigLog/Operations/{id}/LowLevelOperations | Get the detailed low level operations of a high level operation.
[**ConfigLogTPGetOperationById**](ConfigLogTPApi.md#ConfigLogTPGetOperationById) | **Get** /techpreview/{customerid}/{siteid}/ConfigLog/Operations/{id} | Get a single operation log by the id.
[**ConfigLogTPGetOperations**](ConfigLogTPApi.md#ConfigLogTPGetOperations) | **Get** /techpreview/{customerid}/{siteid}/ConfigLog/Operations | Get configuration log operations.
[**ConfigLogTPSetLogSite**](ConfigLogTPApi.md#ConfigLogTPSetLogSite) | **Patch** /techpreview/{customerid}/{siteid}/ConfigLog/LoggingSite | Update logging site info. This on-prem only API.



## ConfigLogTPDeleteLogs

> ConfigLogTPDeleteLogs(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Delete logging event logs. This on-prem only API.

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigLogTPApi.ConfigLogTPDeleteLogs(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigLogTPApi.ConfigLogTPDeleteLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigLogTPDeleteLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## ConfigLogTPDoConfigLogSearch

> LogOperationResponseModelCollection ConfigLogTPDoConfigLogSearch(ctx, customerid, siteid).SearchRequest(searchRequest).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Perform an advanced search for configuration log entries.



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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    searchRequest := *openapiclient.NewLogOperationSearchRequestModel() // LogOperationSearchRequestModel | Specifies the advanced search parameters.
    limit := int32(56) // int32 | The max number of configuration log entries returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the search operations will be executed as a background task. The task will have JobType DoOperationSearch. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigLogTPApi.ConfigLogTPDoConfigLogSearch(context.Background(), customerid, siteid).SearchRequest(searchRequest).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigLogTPApi.ConfigLogTPDoConfigLogSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigLogTPDoConfigLogSearch`: LogOperationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ConfigLogTPApi.ConfigLogTPDoConfigLogSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigLogTPDoConfigLogSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchRequest** | [**LogOperationSearchRequestModel**](LogOperationSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **limit** | **int32** | The max number of configuration log entries returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the search operations will be executed as a background task. The task will have JobType DoOperationSearch. When the task is complete it will redirect to GetJobResults. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LogOperationResponseModelCollection**](LogOperationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigLogTPExportReportData

> CustomReportDataResponseModel ConfigLogTPExportReportData(ctx, customerid, siteid).CustomReportInfoRequestModel(customReportInfoRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Export logging report data in the format of csv/html/both.

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    customReportInfoRequestModel := *openapiclient.NewCustomReportInfoRequestModel(openapiclient.CustomReportFormatType("Html"), openapiclient.CustomReportDateRangeType("All")) // CustomReportInfoRequestModel | CustomReportInfoRequest CustomReportInfoRequestModel
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigLogTPApi.ConfigLogTPExportReportData(context.Background(), customerid, siteid).CustomReportInfoRequestModel(customReportInfoRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigLogTPApi.ConfigLogTPExportReportData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigLogTPExportReportData`: CustomReportDataResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigLogTPApi.ConfigLogTPExportReportData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigLogTPExportReportDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customReportInfoRequestModel** | [**CustomReportInfoRequestModel**](CustomReportInfoRequestModel.md) | CustomReportInfoRequest CustomReportInfoRequestModel | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**CustomReportDataResponseModel**](CustomReportDataResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigLogTPGetFirstLogDate

> string ConfigLogTPGetFirstLogDate(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get first log date

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigLogTPApi.ConfigLogTPGetFirstLogDate(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigLogTPApi.ConfigLogTPGetFirstLogDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigLogTPGetFirstLogDate`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigLogTPApi.ConfigLogTPGetFirstLogDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigLogTPGetFirstLogDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigLogTPGetLogSite

> LogSiteResponseModel ConfigLogTPGetLogSite(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get logging site details. This on-prem only API.

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigLogTPApi.ConfigLogTPGetLogSite(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigLogTPApi.ConfigLogTPGetLogSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigLogTPGetLogSite`: LogSiteResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigLogTPApi.ConfigLogTPGetLogSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigLogTPGetLogSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LogSiteResponseModel**](LogSiteResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigLogTPGetLowLevelOperations

> LowLevelOperationResponseModelCollection ConfigLogTPGetLowLevelOperations(ctx, id, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the detailed low level operations of a high level operation.

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
    id := "id_example" // string | id of specified high level operation
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of low level operations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigLogTPApi.ConfigLogTPGetLowLevelOperations(context.Background(), id, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigLogTPApi.ConfigLogTPGetLowLevelOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigLogTPGetLowLevelOperations`: LowLevelOperationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ConfigLogTPApi.ConfigLogTPGetLowLevelOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of specified high level operation | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigLogTPGetLowLevelOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of low level operations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LowLevelOperationResponseModelCollection**](LowLevelOperationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigLogTPGetOperationById

> LogOperationResponseModel ConfigLogTPGetOperationById(ctx, id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get a single operation log by the id.

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
    id := "id_example" // string | operation id
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigLogTPApi.ConfigLogTPGetOperationById(context.Background(), id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigLogTPApi.ConfigLogTPGetOperationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigLogTPGetOperationById`: LogOperationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigLogTPApi.ConfigLogTPGetOperationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | operation id | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigLogTPGetOperationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LogOperationResponseModel**](LogOperationResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigLogTPGetOperations

> LogOperationResponseModelCollection ConfigLogTPGetOperations(ctx, customerid, siteid).Days(days).SearchDateOption(searchDateOption).Text(text).Limit(limit).TimeDiff(timeDiff).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get configuration log operations.



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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    days := int32(56) // int32 | Optional; number of days of history to retrieve. Note : This parameter is exclusive with parameter searchDateOption. If neither is specified, all records will be returned.              (optional)
    searchDateOption := "searchDateOption_example" // string | Optional; Specific time filters for searching operations Note : This parameter is exclusive with parameter days. If neither is specified, all records will be returned.              (optional)
    text := "text_example" // string | Optional; text pattern to match.  If specified, results will include only entries where Text contains the given case-insensitive pattern.  The pattern may include the wildcard \"*\". If not specified, all records will be returned.              (optional)
    limit := int32(56) // int32 | Optional; the max number of operations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings.              (optional)
    timeDiff := int32(56) // int32 | The time difference in seconds between browser/local time and UTC time. If not specified, it will default to 0. e.g. UTC+8, timeDiff=28800; UTC-5, timeDiff=-18000              (optional) (default to 0)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results.              (optional)
    async := true // bool | If 'true', the get operations will be executed as a background task. The task wil have JobTypeGetOperations. When the task is complete it will redirect to GetJobResults.              (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigLogTPApi.ConfigLogTPGetOperations(context.Background(), customerid, siteid).Days(days).SearchDateOption(searchDateOption).Text(text).Limit(limit).TimeDiff(timeDiff).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigLogTPApi.ConfigLogTPGetOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigLogTPGetOperations`: LogOperationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ConfigLogTPApi.ConfigLogTPGetOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigLogTPGetOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **days** | **int32** | Optional; number of days of history to retrieve. Note : This parameter is exclusive with parameter searchDateOption. If neither is specified, all records will be returned.              | 
 **searchDateOption** | **string** | Optional; Specific time filters for searching operations Note : This parameter is exclusive with parameter days. If neither is specified, all records will be returned.              | 
 **text** | **string** | Optional; text pattern to match.  If specified, results will include only entries where Text contains the given case-insensitive pattern.  The pattern may include the wildcard \&quot;*\&quot;. If not specified, all records will be returned.              | 
 **limit** | **int32** | Optional; the max number of operations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings.              | 
 **timeDiff** | **int32** | The time difference in seconds between browser/local time and UTC time. If not specified, it will default to 0. e.g. UTC+8, timeDiff&#x3D;28800; UTC-5, timeDiff&#x3D;-18000              | [default to 0]
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results.              | 
 **async** | **bool** | If &#39;true&#39;, the get operations will be executed as a background task. The task wil have JobTypeGetOperations. When the task is complete it will redirect to GetJobResults.              | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LogOperationResponseModelCollection**](LogOperationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigLogTPSetLogSite

> LogSiteResponseModel ConfigLogTPSetLogSite(ctx, customerid, siteid).RequestModel(requestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update logging site info. This on-prem only API.

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    requestModel := *openapiclient.NewLogSiteRequestModel() // LogSiteRequestModel | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigLogTPApi.ConfigLogTPSetLogSite(context.Background(), customerid, siteid).RequestModel(requestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigLogTPApi.ConfigLogTPSetLogSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigLogTPSetLogSite`: LogSiteResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigLogTPApi.ConfigLogTPSetLogSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigLogTPSetLogSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestModel** | [**LogSiteRequestModel**](LogSiteRequestModel.md) |  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LogSiteResponseModel**](LogSiteResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

