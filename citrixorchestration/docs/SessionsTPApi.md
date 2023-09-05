# \SessionsTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionsTPDisconnectSession**](SessionsTPApi.md#SessionsTPDisconnectSession) | **Post** /techpreview/{customerid}/{siteid}/Sessions/{id}/$disconnect | Disconnect a session.
[**SessionsTPDoSessionSearch**](SessionsTPApi.md#SessionsTPDoSessionSearch) | **Post** /techpreview/{customerid}/{siteid}/Sessions/$search | Perform an advanced search for sessions.
[**SessionsTPGetSession**](SessionsTPApi.md#SessionsTPGetSession) | **Get** /techpreview/{customerid}/{siteid}/Sessions/{id} | Get details of a single session.
[**SessionsTPGetSessionApplications**](SessionsTPApi.md#SessionsTPGetSessionApplications) | **Get** /techpreview/{customerid}/{siteid}/Sessions/{id}/Applications | Get the list of applications running within a session.
[**SessionsTPGetSessionMachine**](SessionsTPApi.md#SessionsTPGetSessionMachine) | **Get** /techpreview/{customerid}/{siteid}/Sessions/{id}/Machine | Get the details of the machine on which a session is running.
[**SessionsTPGetSessionRecordingStatus**](SessionsTPApi.md#SessionsTPGetSessionRecordingStatus) | **Get** /techpreview/{customerid}/{siteid}/Sessions/RecordingStatus | get session recording status of a session.
[**SessionsTPGetSessions**](SessionsTPApi.md#SessionsTPGetSessions) | **Get** /techpreview/{customerid}/{siteid}/Sessions | Get all sessions in the site.
[**SessionsTPHideSession**](SessionsTPApi.md#SessionsTPHideSession) | **Post** /techpreview/{customerid}/{siteid}/Sessions/{id}/$hide | Hide a session
[**SessionsTPLogoffSession**](SessionsTPApi.md#SessionsTPLogoffSession) | **Post** /techpreview/{customerid}/{siteid}/Sessions/{id}/$logoff | Logoff a session.
[**SessionsTPSendSessionMessage**](SessionsTPApi.md#SessionsTPSendSessionMessage) | **Post** /techpreview/{customerid}/{siteid}/Sessions/{id}/$sendMessage | Send a message to a session.
[**SessionsTPStartSessionRecording**](SessionsTPApi.md#SessionsTPStartSessionRecording) | **Post** /techpreview/{customerid}/{siteid}/Sessions/{id}/$startRecording | start recording for a session.
[**SessionsTPStopSessionRecording**](SessionsTPApi.md#SessionsTPStopSessionRecording) | **Post** /techpreview/{customerid}/{siteid}/Sessions/{id}/$stopRecording | stop recording for a session.
[**SessionsTPUnhideSession**](SessionsTPApi.md#SessionsTPUnhideSession) | **Post** /techpreview/{customerid}/{siteid}/Sessions/{id}/$unhide | Unhide a session



## SessionsTPDisconnectSession

> SessionResponseModel SessionsTPDisconnectSession(ctx, id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Disconnect a session.



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
    id := "id_example" // string | ID of the session to disconnect.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the session will be disconnected as a background task. The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the session being disconnected. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPDisconnectSession(context.Background(), id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPDisconnectSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPDisconnectSession`: SessionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPDisconnectSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session to disconnect. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPDisconnectSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the session will be disconnected as a background task. The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the session being disconnected. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SessionResponseModel**](SessionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPDoSessionSearch

> SessionResponseModelCollection SessionsTPDoSessionSearch(ctx, customerid, siteid).SearchRequest(searchRequest).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Skip(skip).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Perform an advanced search for sessions.



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
    searchRequest := *openapiclient.NewMachineAndSessionSearchRequestModel() // MachineAndSessionSearchRequestModel | Specifies the advanced search parameters.
    limit := int32(56) // int32 | The max number of sessions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server (optional)
    async := true // bool | If `true`, the search sessions will be executed as a background task. The task will have JobType DoSessionSearch. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    skip := int32(56) // int32 | The number of sessions to skip. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPDoSessionSearch(context.Background(), customerid, siteid).SearchRequest(searchRequest).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Skip(skip).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPDoSessionSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPDoSessionSearch`: SessionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPDoSessionSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPDoSessionSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchRequest** | [**MachineAndSessionSearchRequestModel**](MachineAndSessionSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **limit** | **int32** | The max number of sessions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server | 
 **async** | **bool** | If &#x60;true&#x60;, the search sessions will be executed as a background task. The task will have JobType DoSessionSearch. When the task is complete it will redirect to GetJobResults. | [default to false]
 **skip** | **int32** | The number of sessions to skip. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SessionResponseModelCollection**](SessionResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPGetSession

> SessionResponseModel SessionsTPGetSession(ctx, id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details of a single session.



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
    id := "id_example" // string | ID of the session.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPGetSession(context.Background(), id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPGetSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPGetSession`: SessionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPGetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SessionResponseModel**](SessionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPGetSessionApplications

> ApplicationResponseModelCollection SessionsTPGetSessionApplications(ctx, id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the list of applications running within a session.



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
    id := "id_example" // string | ID of the session.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the get applications will be executed as a background task. The task will have JobType GetSessionApplications. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the session being queried. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPGetSessionApplications(context.Background(), id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPGetSessionApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPGetSessionApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPGetSessionApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPGetSessionApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the get applications will be executed as a background task. The task will have JobType GetSessionApplications. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the session being queried. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationResponseModelCollection**](ApplicationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPGetSessionMachine

> MachineDetailResponseModel SessionsTPGetSessionMachine(ctx, id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the details of the machine on which a session is running.



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
    id := "id_example" // string | ID of the session.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPGetSessionMachine(context.Background(), id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPGetSessionMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPGetSessionMachine`: MachineDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPGetSessionMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPGetSessionMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**MachineDetailResponseModel**](MachineDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPGetSessionRecordingStatus

> SessionRecordingStatusResponseModel SessionsTPGetSessionRecordingStatus(ctx, customerid, siteid).Id(id).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

get session recording status of a session.



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
    id := "id_example" // string | ID of the session.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPGetSessionRecordingStatus(context.Background(), customerid, siteid).Id(id).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPGetSessionRecordingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPGetSessionRecordingStatus`: SessionRecordingStatusResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPGetSessionRecordingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPGetSessionRecordingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | ID of the session. | 


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SessionRecordingStatusResponseModel**](SessionRecordingStatusResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPGetSessions

> SessionResponseModelCollection SessionsTPGetSessions(ctx, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all sessions in the site.



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
    limit := int32(56) // int32 | The max number of sessions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the get sessions will be executed as a background task. The task will have JobType GetSessions. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPGetSessions(context.Background(), customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPGetSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPGetSessions`: SessionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPGetSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of sessions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the get sessions will be executed as a background task. The task will have JobType GetSessions. When the task is complete it will redirect to GetJobResults. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SessionResponseModelCollection**](SessionResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPHideSession

> SessionsTPHideSession(ctx, id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Hide a session

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
    id := "id_example" // string | ID of the session to hide or unhide.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the session will be hidden as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SessionsTPApi.SessionsTPHideSession(context.Background(), id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPHideSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session to hide or unhide. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPHideSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the session will be hidden as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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


## SessionsTPLogoffSession

> SessionsTPLogoffSession(ctx, id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Logoff a session.



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
    id := "id_example" // string | ID of the session to log off.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the session will be logged off as a background task. The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the session being logged off. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SessionsTPApi.SessionsTPLogoffSession(context.Background(), id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPLogoffSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session to log off. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPLogoffSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the session will be logged off as a background task. The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the session being logged off. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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


## SessionsTPSendSessionMessage

> SessionResponseModel SessionsTPSendSessionMessage(ctx, id, customerid, siteid).Message(message).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Send a message to a session.



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
    id := "id_example" // string | ID of the session to send a message to.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    message := *openapiclient.NewSendMessageRequestModel() // SendMessageRequestModel | Details of the message to send.
    async := true // bool | If `true`, the message will be sent as a background task. The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the session to which a message is being sent. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPSendSessionMessage(context.Background(), id, customerid, siteid).Message(message).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPSendSessionMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPSendSessionMessage`: SessionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPSendSessionMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session to send a message to. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPSendSessionMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **message** | [**SendMessageRequestModel**](SendMessageRequestModel.md) | Details of the message to send. | 
 **async** | **bool** | If &#x60;true&#x60;, the message will be sent as a background task. The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the session to which a message is being sent. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SessionResponseModel**](SessionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPStartSessionRecording

> SessionResponseModel SessionsTPStartSessionRecording(ctx, id, customerid, siteid).NotifyUser(notifyUser).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

start recording for a session.



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
    id := "id_example" // string | ID of the session.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    notifyUser := true // bool | Whether to notify session user or not. (optional) (default to true)
    async := true // bool | If `true`, the session will be recorded as a background task. The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the session. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPStartSessionRecording(context.Background(), id, customerid, siteid).NotifyUser(notifyUser).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPStartSessionRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPStartSessionRecording`: SessionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPStartSessionRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPStartSessionRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **notifyUser** | **bool** | Whether to notify session user or not. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the session will be recorded as a background task. The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the session. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SessionResponseModel**](SessionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPStopSessionRecording

> SessionResponseModel SessionsTPStopSessionRecording(ctx, id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

stop recording for a session.



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
    id := "id_example" // string | ID of the session.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the session will be recorded as a background task. The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the session. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsTPApi.SessionsTPStopSessionRecording(context.Background(), id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPStopSessionRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsTPStopSessionRecording`: SessionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SessionsTPApi.SessionsTPStopSessionRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPStopSessionRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the session will be recorded as a background task. The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the session. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SessionResponseModel**](SessionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsTPUnhideSession

> SessionsTPUnhideSession(ctx, id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Unhide a session

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
    id := "id_example" // string | ID of the session to hide or unhide.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the session will be hidden as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SessionsTPApi.SessionsTPUnhideSession(context.Background(), id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsTPApi.SessionsTPUnhideSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session to hide or unhide. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsTPUnhideSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the session will be hidden as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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

