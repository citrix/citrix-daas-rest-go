# \UpgradeSchedulesAPIs 

All URIs are relative to *https://api-us.cloud.com/cvad/manage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpgradeSchedulesCreateUpgradeScheduleForMachines**](UpgradeSchedulesAPIs .md#UpgradeSchedulesCreateUpgradeScheduleForMachines) | **Post** /UpgradeSchedules | Setup a new VDA upgrade schedule for a list of machines.
[**UpgradeSchedulesRemoveUpgradeSchedule**](UpgradeSchedulesAPIs .md#UpgradeSchedulesRemoveUpgradeSchedule) | **Delete** /UpgradeSchedules/{id} | Cancel a VDA upgrade schedule by id.
[**UpgradeSchedulesUpdateUpgradeSchedule**](UpgradeSchedulesAPIs .md#UpgradeSchedulesUpdateUpgradeSchedule) | **Patch** /UpgradeSchedules/{id} | Reschedule a VDA upgrade schedule by id.



## UpgradeSchedulesCreateUpgradeScheduleForMachines

> UpgradeSchedulesCreateUpgradeScheduleForMachines(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateUpgradeScheduleForMachinesRequestModel(createUpgradeScheduleForMachinesRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Setup a new VDA upgrade schedule for a list of machines.

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    createUpgradeScheduleForMachinesRequestModel := *openapiclient.NewCreateUpgradeScheduleForMachinesRequestModel([]string{"MachineNameOrIds_example"}, int32(123)) // CreateUpgradeScheduleForMachinesRequestModel | Details about the VDA upgrade schedules to create.
    async := true // bool | If `true`, the VDA upgrade schedules will be created as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UpgradeSchedulesAPIs .UpgradeSchedulesCreateUpgradeScheduleForMachines(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateUpgradeScheduleForMachinesRequestModel(createUpgradeScheduleForMachinesRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpgradeSchedulesAPIs .UpgradeSchedulesCreateUpgradeScheduleForMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSchedulesCreateUpgradeScheduleForMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **createUpgradeScheduleForMachinesRequestModel** | [**CreateUpgradeScheduleForMachinesRequestModel**](CreateUpgradeScheduleForMachinesRequestModel.md) | Details about the VDA upgrade schedules to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedules will be created as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## UpgradeSchedulesRemoveUpgradeSchedule

> UpgradeSchedulesRemoveUpgradeSchedule(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Cancel a VDA upgrade schedule by id.

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
    id := "id_example" // string | ID of the VDA upgrade schedule.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the VDA upgrade schedule will be removed as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UpgradeSchedulesAPIs .UpgradeSchedulesRemoveUpgradeSchedule(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpgradeSchedulesAPIs .UpgradeSchedulesRemoveUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the VDA upgrade schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSchedulesRemoveUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedule will be removed as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## UpgradeSchedulesUpdateUpgradeSchedule

> UpgradeSchedulesUpdateUpgradeSchedule(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateUpgradeScheduleRequestModel(updateUpgradeScheduleRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Reschedule a VDA upgrade schedule by id.

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
    id := "id_example" // string | ID of the VDA upgrade schedule.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    updateUpgradeScheduleRequestModel := *openapiclient.NewUpdateUpgradeScheduleRequestModel(time.Now(), int32(123)) // UpdateUpgradeScheduleRequestModel | 
    async := true // bool | If `true`, the VDA upgrade schedule will be updated as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UpgradeSchedulesAPIs .UpgradeSchedulesUpdateUpgradeSchedule(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateUpgradeScheduleRequestModel(updateUpgradeScheduleRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpgradeSchedulesAPIs .UpgradeSchedulesUpdateUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the VDA upgrade schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSchedulesUpdateUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **updateUpgradeScheduleRequestModel** | [**UpdateUpgradeScheduleRequestModel**](UpdateUpgradeScheduleRequestModel.md) |  | 
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedule will be updated as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

