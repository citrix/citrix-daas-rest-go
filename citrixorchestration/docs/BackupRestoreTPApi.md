# \BackupRestoreTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupRestoreTPCreateBackupSchedule**](BackupRestoreTPApi.md#BackupRestoreTPCreateBackupSchedule) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/Schedules | Create backup schedule
[**BackupRestoreTPDeleteBackup**](BackupRestoreTPApi.md#BackupRestoreTPDeleteBackup) | **Delete** /techpreview/{customerid}/{siteid}/BackupRestore/{backupName} | Delete a backup
[**BackupRestoreTPDeleteBackupSchedule**](BackupRestoreTPApi.md#BackupRestoreTPDeleteBackupSchedule) | **Delete** /techpreview/{customerid}/{siteid}/BackupRestore/Schedule/{uid} | Delete backup schedule
[**BackupRestoreTPDisableBackupSchedule**](BackupRestoreTPApi.md#BackupRestoreTPDisableBackupSchedule) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/Schedules/Disable/{uid} | Disable backup schedule
[**BackupRestoreTPDoBackup**](BackupRestoreTPApi.md#BackupRestoreTPDoBackup) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/DoBackup | Do backup
[**BackupRestoreTPEnableBackupSchedule**](BackupRestoreTPApi.md#BackupRestoreTPEnableBackupSchedule) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/Schedules/Enable/{uid} | Enable backup schedule
[**BackupRestoreTPGetAllBackupHistory**](BackupRestoreTPApi.md#BackupRestoreTPGetAllBackupHistory) | **Get** /techpreview/{customerid}/{siteid}/BackupRestore/History | Get all backup history
[**BackupRestoreTPGetBackupHistoryForSpecificBackup**](BackupRestoreTPApi.md#BackupRestoreTPGetBackupHistoryForSpecificBackup) | **Get** /techpreview/{customerid}/{siteid}/BackupRestore/History/All/{backupName} | Get all backup history for specific backup
[**BackupRestoreTPGetBackupHistoryForSpecificHistory**](BackupRestoreTPApi.md#BackupRestoreTPGetBackupHistoryForSpecificHistory) | **Get** /techpreview/{customerid}/{siteid}/BackupRestore/History/{uid} | Get backup history for a single specific backup
[**BackupRestoreTPGetBackupSchedule**](BackupRestoreTPApi.md#BackupRestoreTPGetBackupSchedule) | **Get** /techpreview/{customerid}/{siteid}/BackupRestore/Schedules/{Uid} | Get single backup schedule
[**BackupRestoreTPGetBackupSchedules**](BackupRestoreTPApi.md#BackupRestoreTPGetBackupSchedules) | **Get** /techpreview/{customerid}/{siteid}/BackupRestore/Schedules | Get backup schedules
[**BackupRestoreTPGetBackups**](BackupRestoreTPApi.md#BackupRestoreTPGetBackups) | **Get** /techpreview/{customerid}/{siteid}/BackupRestore | Get backups
[**BackupRestoreTPGetComponentMembers**](BackupRestoreTPApi.md#BackupRestoreTPGetComponentMembers) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/Members/{backupName} | Get backup Members
[**BackupRestoreTPGetCurrentActionStatus**](BackupRestoreTPApi.md#BackupRestoreTPGetCurrentActionStatus) | **Get** /techpreview/{customerid}/{siteid}/BackupRestore/Status | Get backup / restore status
[**BackupRestoreTPGetInformation**](BackupRestoreTPApi.md#BackupRestoreTPGetInformation) | **Get** /techpreview/{customerid}/{siteid}/BackupRestore/Information | Get backup / restore information
[**BackupRestoreTPModifyBackupSchedule**](BackupRestoreTPApi.md#BackupRestoreTPModifyBackupSchedule) | **Patch** /techpreview/{customerid}/{siteid}/BackupRestore/Schedules/{uid} | Modify backup schedules
[**BackupRestoreTPPinBackup**](BackupRestoreTPApi.md#BackupRestoreTPPinBackup) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/{backupName}/PinBackup | Pin a backup
[**BackupRestoreTPRemoveBackupHistory**](BackupRestoreTPApi.md#BackupRestoreTPRemoveBackupHistory) | **Delete** /techpreview/{customerid}/{siteid}/BackupRestore/History | Remove history entries falling within supplied date range
[**BackupRestoreTPRestoreBackup**](BackupRestoreTPApi.md#BackupRestoreTPRestoreBackup) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/Restore/{backupName} | Restore a backup
[**BackupRestoreTPSetCurrentActionStatus**](BackupRestoreTPApi.md#BackupRestoreTPSetCurrentActionStatus) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/Status | Set backup / restore status
[**BackupRestoreTPSetInformation**](BackupRestoreTPApi.md#BackupRestoreTPSetInformation) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/Information | Set backup / restore information
[**BackupRestoreTPUnpinBackup**](BackupRestoreTPApi.md#BackupRestoreTPUnpinBackup) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/{backupName}/UnpinBackup | Unpin a backup
[**BackupRestoreTPUpdateBckRstrHistoryNotes**](BackupRestoreTPApi.md#BackupRestoreTPUpdateBckRstrHistoryNotes) | **Post** /techpreview/{customerid}/{siteid}/BackupRestore/History/{uid} | Update notes in history entry



## BackupRestoreTPCreateBackupSchedule

> BackupRestoreScheduleResponseModel BackupRestoreTPCreateBackupSchedule(ctx, customerid, siteid).Schedule(schedule).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create backup schedule



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
    schedule := *openapiclient.NewBackupRestoreScheduleRequestModel("Name_example", "StartDate_example", int32(123), false, openapiclient.BackupRestoreScheduleFrequency("Daily"), "StartTime_example") // BackupRestoreScheduleRequestModel | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPCreateBackupSchedule(context.Background(), customerid, siteid).Schedule(schedule).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPCreateBackupSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPCreateBackupSchedule`: BackupRestoreScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPCreateBackupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPCreateBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **schedule** | [**BackupRestoreScheduleRequestModel**](BackupRestoreScheduleRequestModel.md) |  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreScheduleResponseModel**](BackupRestoreScheduleResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPDeleteBackup

> bool BackupRestoreTPDeleteBackup(ctx, customerid, siteid, backupName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a backup



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
    backupName := "backupName_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPDeleteBackup(context.Background(), customerid, siteid, backupName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPDeleteBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPDeleteBackup`: bool
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPDeleteBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 
**backupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPDeleteBackupSchedule

> BackupRestoreTPDeleteBackupSchedule(ctx, customerid, siteid, uid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete backup schedule



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
    uid := int32(56) // int32 | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPDeleteBackupSchedule(context.Background(), customerid, siteid, uid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPDeleteBackupSchedule``: %v\n", err)
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
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPDeleteBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## BackupRestoreTPDisableBackupSchedule

> BackupRestoreTPDisableBackupSchedule(ctx, customerid, siteid, uid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Disable backup schedule



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
    uid := int32(56) // int32 | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPDisableBackupSchedule(context.Background(), customerid, siteid, uid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPDisableBackupSchedule``: %v\n", err)
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
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPDisableBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## BackupRestoreTPDoBackup

> BackupDoBackupResponseModel BackupRestoreTPDoBackup(ctx, customerid, siteid).BackupRequest(backupRequest).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Do backup



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
    backupRequest := *openapiclient.NewBackupRestoreDoBackupRequestModel() // BackupRestoreDoBackupRequestModel | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPDoBackup(context.Background(), customerid, siteid).BackupRequest(backupRequest).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPDoBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPDoBackup`: BackupDoBackupResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPDoBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPDoBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backupRequest** | [**BackupRestoreDoBackupRequestModel**](BackupRestoreDoBackupRequestModel.md) |  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupDoBackupResponseModel**](BackupDoBackupResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPEnableBackupSchedule

> BackupRestoreTPEnableBackupSchedule(ctx, customerid, siteid, uid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Enable backup schedule



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
    uid := int32(56) // int32 | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPEnableBackupSchedule(context.Background(), customerid, siteid, uid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPEnableBackupSchedule``: %v\n", err)
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
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPEnableBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## BackupRestoreTPGetAllBackupHistory

> BackupRestoreHistoryResponseModelCollection BackupRestoreTPGetAllBackupHistory(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all backup history

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
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPGetAllBackupHistory(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPGetAllBackupHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPGetAllBackupHistory`: BackupRestoreHistoryResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPGetAllBackupHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPGetAllBackupHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreHistoryResponseModelCollection**](BackupRestoreHistoryResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPGetBackupHistoryForSpecificBackup

> BackupRestoreHistoryResponseModelCollection BackupRestoreTPGetBackupHistoryForSpecificBackup(ctx, customerid, siteid, backupName).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all backup history for specific backup

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
    backupName := "backupName_example" // string | 
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPGetBackupHistoryForSpecificBackup(context.Background(), customerid, siteid, backupName).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPGetBackupHistoryForSpecificBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPGetBackupHistoryForSpecificBackup`: BackupRestoreHistoryResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPGetBackupHistoryForSpecificBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 
**backupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPGetBackupHistoryForSpecificBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreHistoryResponseModelCollection**](BackupRestoreHistoryResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPGetBackupHistoryForSpecificHistory

> BackupRestoreHistoryInformation BackupRestoreTPGetBackupHistoryForSpecificHistory(ctx, customerid, siteid, uid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get backup history for a single specific backup

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
    uid := "uid_example" // string | 
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPGetBackupHistoryForSpecificHistory(context.Background(), customerid, siteid, uid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPGetBackupHistoryForSpecificHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPGetBackupHistoryForSpecificHistory`: BackupRestoreHistoryInformation
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPGetBackupHistoryForSpecificHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPGetBackupHistoryForSpecificHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreHistoryInformation**](BackupRestoreHistoryInformation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPGetBackupSchedule

> BackupRestoreScheduleModel BackupRestoreTPGetBackupSchedule(ctx, customerid, siteid, uid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get single backup schedule



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
    uid := int32(56) // int32 | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPGetBackupSchedule(context.Background(), customerid, siteid, uid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPGetBackupSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPGetBackupSchedule`: BackupRestoreScheduleModel
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPGetBackupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPGetBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreScheduleModel**](BackupRestoreScheduleModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPGetBackupSchedules

> BackupRestoreScheduleModelCollection BackupRestoreTPGetBackupSchedules(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get backup schedules



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
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPGetBackupSchedules(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPGetBackupSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPGetBackupSchedules`: BackupRestoreScheduleModelCollection
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPGetBackupSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPGetBackupSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreScheduleModelCollection**](BackupRestoreScheduleModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPGetBackups

> BackupResponseModelCollection BackupRestoreTPGetBackups(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get backups

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
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPGetBackups(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPGetBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPGetBackups`: BackupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPGetBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPGetBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupResponseModelCollection**](BackupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPGetComponentMembers

> BackupMembersResponseModelCollection BackupRestoreTPGetComponentMembers(ctx, customerid, siteid, backupName).GetMembersRequest(getMembersRequest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get backup Members



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
    backupName := "backupName_example" // string | Name of the backp to restore
    getMembersRequest := *openapiclient.NewBackupRestoreGetComponentMemberRequestModel(openapiclient.BckRstrAutoConfigComponents("Tags")) // BackupRestoreGetComponentMemberRequestModel | 
    async := true // bool |  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPGetComponentMembers(context.Background(), customerid, siteid, backupName).GetMembersRequest(getMembersRequest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPGetComponentMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPGetComponentMembers`: BackupMembersResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPGetComponentMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 
**backupName** | **string** | Name of the backp to restore | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPGetComponentMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **getMembersRequest** | [**BackupRestoreGetComponentMemberRequestModel**](BackupRestoreGetComponentMemberRequestModel.md) |  | 
 **async** | **bool** |  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupMembersResponseModelCollection**](BackupMembersResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPGetCurrentActionStatus

> BackupRestoreStatusResponseModel BackupRestoreTPGetCurrentActionStatus(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get backup / restore status



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
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPGetCurrentActionStatus(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPGetCurrentActionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPGetCurrentActionStatus`: BackupRestoreStatusResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPGetCurrentActionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPGetCurrentActionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreStatusResponseModel**](BackupRestoreStatusResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPGetInformation

> BackupRestoreInformationResponseModel BackupRestoreTPGetInformation(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get backup / restore information



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
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPGetInformation(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPGetInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPGetInformation`: BackupRestoreInformationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPGetInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPGetInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreInformationResponseModel**](BackupRestoreInformationResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPModifyBackupSchedule

> BackupRestoreScheduleResponseModel BackupRestoreTPModifyBackupSchedule(ctx, customerid, siteid, uid).Schedule(schedule).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Modify backup schedules



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
    uid := int32(56) // int32 | 
    schedule := *openapiclient.NewBackupRestoreScheduleRequestModel("Name_example", "StartDate_example", int32(123), false, openapiclient.BackupRestoreScheduleFrequency("Daily"), "StartTime_example") // BackupRestoreScheduleRequestModel | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPModifyBackupSchedule(context.Background(), customerid, siteid, uid).Schedule(schedule).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPModifyBackupSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPModifyBackupSchedule`: BackupRestoreScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPModifyBackupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPModifyBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **schedule** | [**BackupRestoreScheduleRequestModel**](BackupRestoreScheduleRequestModel.md) |  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreScheduleResponseModel**](BackupRestoreScheduleResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPPinBackup

> bool BackupRestoreTPPinBackup(ctx, customerid, siteid, backupName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Pin a backup



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
    backupName := "backupName_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPPinBackup(context.Background(), customerid, siteid, backupName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPPinBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPPinBackup`: bool
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPPinBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 
**backupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPPinBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPRemoveBackupHistory

> BackupRestoreTPRemoveBackupHistory(ctx, customerid, siteid).DateRange(dateRange).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Remove history entries falling within supplied date range

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    dateRange := *openapiclient.NewBackupRestoreRemoveHistoryRequestModel(time.Now(), time.Now()) // BackupRestoreRemoveHistoryRequestModel | Start and end date of history entries to remove
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPRemoveBackupHistory(context.Background(), customerid, siteid).DateRange(dateRange).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPRemoveBackupHistory``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBackupRestoreTPRemoveBackupHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dateRange** | [**BackupRestoreRemoveHistoryRequestModel**](BackupRestoreRemoveHistoryRequestModel.md) | Start and end date of history entries to remove | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## BackupRestoreTPRestoreBackup

> BackupRestoreRestoreBackupResponseModel BackupRestoreTPRestoreBackup(ctx, customerid, siteid, backupName).RestoreRequest(restoreRequest).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Restore a backup



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
    backupName := "backupName_example" // string | 
    restoreRequest := *openapiclient.NewBackupRestoreRestoreBackupRequestModel("BackupName_example", int32(123), openapiclient.BckRstrAutoConfigComponents("Tags")) // BackupRestoreRestoreBackupRequestModel | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPRestoreBackup(context.Background(), customerid, siteid, backupName).RestoreRequest(restoreRequest).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPRestoreBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPRestoreBackup`: BackupRestoreRestoreBackupResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPRestoreBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 
**backupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restoreRequest** | [**BackupRestoreRestoreBackupRequestModel**](BackupRestoreRestoreBackupRequestModel.md) |  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreRestoreBackupResponseModel**](BackupRestoreRestoreBackupResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPSetCurrentActionStatus

> BackupRestoreStatusResponseModel BackupRestoreTPSetCurrentActionStatus(ctx, customerid, siteid).CallersStatus(callersStatus).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Set backup / restore status



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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    callersStatus := *openapiclient.NewBackupRestoreStatusRequestModel(openapiclient.BackupRestoreActions("Backup"), "ExecutionId_example", int32(123), openapiclient.BackupRestoreStatus("Unknown"), time.Now(), time.Now()) // BackupRestoreStatusRequestModel | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPSetCurrentActionStatus(context.Background(), customerid, siteid).CallersStatus(callersStatus).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPSetCurrentActionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPSetCurrentActionStatus`: BackupRestoreStatusResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPSetCurrentActionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPSetCurrentActionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **callersStatus** | [**BackupRestoreStatusRequestModel**](BackupRestoreStatusRequestModel.md) |  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BackupRestoreStatusResponseModel**](BackupRestoreStatusResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPSetInformation

> BackupRestoreTPSetInformation(ctx, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Set backup / restore information



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
    request := *openapiclient.NewBackupRestoreInformationRequestModel("TimeZone_example", int32(123), false) // BackupRestoreInformationRequestModel | Request containing BackupRestoreInformationRequestModel values
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPSetInformation(context.Background(), customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPSetInformation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBackupRestoreTPSetInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**BackupRestoreInformationRequestModel**](BackupRestoreInformationRequestModel.md) | Request containing BackupRestoreInformationRequestModel values | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## BackupRestoreTPUnpinBackup

> bool BackupRestoreTPUnpinBackup(ctx, customerid, siteid, backupName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Unpin a backup



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
    backupName := "backupName_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPUnpinBackup(context.Background(), customerid, siteid, backupName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPUnpinBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPUnpinBackup`: bool
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPUnpinBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 
**backupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPUnpinBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreTPUpdateBckRstrHistoryNotes

> bool BackupRestoreTPUpdateBckRstrHistoryNotes(ctx, uid, customerid, siteid).Notes(notes).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update notes in history entry

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
    uid := int32(56) // int32 | History entry unique identifier
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    notes := *openapiclient.NewBackupRestoreUpdateNotesRequestModel() // BackupRestoreUpdateNotesRequestModel | new notes
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupRestoreTPApi.BackupRestoreTPUpdateBckRstrHistoryNotes(context.Background(), uid, customerid, siteid).Notes(notes).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreTPApi.BackupRestoreTPUpdateBckRstrHistoryNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupRestoreTPUpdateBckRstrHistoryNotes`: bool
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreTPApi.BackupRestoreTPUpdateBckRstrHistoryNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | History entry unique identifier | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreTPUpdateBckRstrHistoryNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **notes** | [**BackupRestoreUpdateNotesRequestModel**](BackupRestoreUpdateNotesRequestModel.md) | new notes | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

