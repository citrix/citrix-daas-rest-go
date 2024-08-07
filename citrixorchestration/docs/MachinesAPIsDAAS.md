# \MachinesAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MachinesAddMachineTags**](MachinesAPIsDAAS.md#MachinesAddMachineTags) | **Post** /Machines/{nameOrId}/Tags/{tagNameOrId} | Add a tag to a machine.
[**MachinesCheckMachineExists**](MachinesAPIsDAAS.md#MachinesCheckMachineExists) | **Head** /Machines/{name} | Check for the existence of a machine by name.
[**MachinesCreateMachineUpgradeSchedule**](MachinesAPIsDAAS.md#MachinesCreateMachineUpgradeSchedule) | **Post** /Machines/{nameOrId}/UpgradeSchedule | Setup a new VDA upgrade schedule for a machine.
[**MachinesDisconnectSessions**](MachinesAPIsDAAS.md#MachinesDisconnectSessions) | **Post** /Machines/{nameOrId}/$disconnect | Disconnect all sessions on a machine.
[**MachinesDoMachineSearch**](MachinesAPIsDAAS.md#MachinesDoMachineSearch) | **Post** /Machines/$search | Perform an advanced search for machines.
[**MachinesGetMachine**](MachinesAPIsDAAS.md#MachinesGetMachine) | **Get** /Machines/{nameOrId} | Get details of a single machine which belongs to this site, or have registered but are not yet configured in this site
[**MachinesGetMachineApplications**](MachinesAPIsDAAS.md#MachinesGetMachineApplications) | **Get** /Machines/{nameOrId}/Applications | Get the list of applications on a machine.
[**MachinesGetMachineDeliveryGroup**](MachinesAPIsDAAS.md#MachinesGetMachineDeliveryGroup) | **Get** /Machines/{nameOrId}/DeliveryGroup | Get the delivery group for a machine.
[**MachinesGetMachineDesktop**](MachinesAPIsDAAS.md#MachinesGetMachineDesktop) | **Get** /Machines/{nameOrId}/Desktop | Get the desktop associated with the machine, if any.
[**MachinesGetMachineMachineCatalog**](MachinesAPIsDAAS.md#MachinesGetMachineMachineCatalog) | **Get** /Machines/{nameOrId}/MachineCatalog | Get the machine catalog for a machine.
[**MachinesGetMachinePowerActionSchedules**](MachinesAPIsDAAS.md#MachinesGetMachinePowerActionSchedules) | **Get** /Machines/{nameOrId}/PowerActionSchedules | Get the power action schedules associated with a machine.
[**MachinesGetMachineSessions**](MachinesAPIsDAAS.md#MachinesGetMachineSessions) | **Get** /Machines/{nameOrId}/Sessions | Get the list of sessions running on a machine.
[**MachinesGetMachineStartMenuShortcutIcon**](MachinesAPIsDAAS.md#MachinesGetMachineStartMenuShortcutIcon) | **Get** /Machines/{nameOrId}/StartMenuShortcutIcon | Get a start menu shortcut icon from the machine.
[**MachinesGetMachineStartMenuShortcuts**](MachinesAPIsDAAS.md#MachinesGetMachineStartMenuShortcuts) | **Get** /Machines/{nameOrId}/StartMenuShortcuts | Get start menu shortcuts from the machine.
[**MachinesGetMachineTags**](MachinesAPIsDAAS.md#MachinesGetMachineTags) | **Get** /Machines/{nameOrId}/Tags | Get tags associated with a machine.
[**MachinesGetMachineUpgradeVersions**](MachinesAPIsDAAS.md#MachinesGetMachineUpgradeVersions) | **Get** /Machines/{nameOrId}/UpgradeVersions | Get available upgrade versions for a machine.
[**MachinesGetMachines**](MachinesAPIsDAAS.md#MachinesGetMachines) | **Get** /Machines | Get all machines in the site.
[**MachinesGetMachinesAdministrators**](MachinesAPIsDAAS.md#MachinesGetMachinesAdministrators) | **Get** /Machines/{nameOrId}/Administrators | Get administrators who can administer a machine
[**MachinesGetMachinesV2**](MachinesAPIsDAAS.md#MachinesGetMachinesV2) | **Get** /MachinesV2 | The V2 version of get all machines in the site.
[**MachinesGetTestMachineReport**](MachinesAPIsDAAS.md#MachinesGetTestMachineReport) | **Get** /Machines/{nameOrId}/TestReports/{reportId} | Get Cloud Health Check Report on a VDA machine.
[**MachinesImportFileTypes**](MachinesAPIsDAAS.md#MachinesImportFileTypes) | **Post** /Machines/{nameOrId}/FileTypes/$import | Import file type associations from the machine.
[**MachinesLogoffMachineSessions**](MachinesAPIsDAAS.md#MachinesLogoffMachineSessions) | **Post** /Machines/{nameOrId}/$logoff | Logoff all sessions on a machine.
[**MachinesRebootMachine**](MachinesAPIsDAAS.md#MachinesRebootMachine) | **Post** /Machines/{nameOrId}/$reboot | Reboot a machine.
[**MachinesRemoveMachine**](MachinesAPIsDAAS.md#MachinesRemoveMachine) | **Delete** /Machines/{nameOrId} | Delete a machine from the site.
[**MachinesRemoveMachineTags**](MachinesAPIsDAAS.md#MachinesRemoveMachineTags) | **Delete** /Machines/{nameOrId}/Tags/{tagNameOrId} | Remove a tag from a machine.
[**MachinesRemoveMachineUpgradeSchedule**](MachinesAPIsDAAS.md#MachinesRemoveMachineUpgradeSchedule) | **Delete** /Machines/{nameOrId}/UpgradeSchedule | Remove a pending VDA upgrade schedule for a machine.
[**MachinesRemovePowerActionSchedule**](MachinesAPIsDAAS.md#MachinesRemovePowerActionSchedule) | **Delete** /Machines/{nameOrId}/PowerActionSchedules | Remove pending power action schedules for a machine.
[**MachinesResetProvVMDisk**](MachinesAPIsDAAS.md#MachinesResetProvVMDisk) | **Post** /Machines/{nameOrId}/$ResetProvVMDisk | Resets the OS disk of persistent VMs to the current master image.
[**MachinesResumeMachine**](MachinesAPIsDAAS.md#MachinesResumeMachine) | **Post** /Machines/{nameOrId}/$resume | Resume a machine.
[**MachinesSetMachineTags**](MachinesAPIsDAAS.md#MachinesSetMachineTags) | **Put** /Machines/{nameOrId}/Tags | Set tags associated with a machine.
[**MachinesShutdownMachine**](MachinesAPIsDAAS.md#MachinesShutdownMachine) | **Post** /Machines/{nameOrId}/$shutdown | Shut down a machine.
[**MachinesStartMachine**](MachinesAPIsDAAS.md#MachinesStartMachine) | **Post** /Machines/{nameOrId}/$start | Start a machine.
[**MachinesSuspendMachine**](MachinesAPIsDAAS.md#MachinesSuspendMachine) | **Post** /Machines/{nameOrId}/$suspend | Suspend a machine.
[**MachinesTestMachine**](MachinesAPIsDAAS.md#MachinesTestMachine) | **Post** /Machines/{nameOrId}/$test | Run Cloud Health Check on a VDA machine.
[**MachinesUpdateMachineCatalogMachine**](MachinesAPIsDAAS.md#MachinesUpdateMachineCatalogMachine) | **Patch** /Machines/{nameOrId} | Update a machine.
[**MachinesUpdateMachineUpgradeSchedule**](MachinesAPIsDAAS.md#MachinesUpdateMachineUpgradeSchedule) | **Patch** /Machines/{nameOrId}/UpgradeSchedule | Reschedule a pending VDA upgrade schedule for a machine.



## MachinesAddMachineTags

> TagResponseModelCollection MachinesAddMachineTags(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Add a tag to a machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag to add.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesAddMachineTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesAddMachineTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesAddMachineTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesAddMachineTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 
**tagNameOrId** | **string** | Name or ID of the tag to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesAddMachineTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. | [default to false]

### Return type

[**TagResponseModelCollection**](TagResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesCheckMachineExists

> MachinesCheckMachineExists(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Check for the existence of a machine by name.



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
    name := "name_example" // string | Name of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'.             For instance, if a MachineName is \"DomainA\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesCheckMachineExists(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesCheckMachineExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;.             For instance, if a MachineName is \&quot;DomainA\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesCheckMachineExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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


## MachinesCreateMachineUpgradeSchedule

> MachinesCreateMachineUpgradeSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateUpgradeScheduleRequestModel(createUpgradeScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Setup a new VDA upgrade schedule for a machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    createUpgradeScheduleRequestModel := *openapiclient.NewCreateUpgradeScheduleRequestModel(int32(123)) // CreateUpgradeScheduleRequestModel | Details about the VDA upgrade schedule to create.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the VDA upgrade schedule will be created as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesCreateMachineUpgradeSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateUpgradeScheduleRequestModel(createUpgradeScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesCreateMachineUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesCreateMachineUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **createUpgradeScheduleRequestModel** | [**CreateUpgradeScheduleRequestModel**](CreateUpgradeScheduleRequestModel.md) | Details about the VDA upgrade schedule to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedule will be created as a background task. | [default to false]

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


## MachinesDisconnectSessions

> MachineDetailResponseModel MachinesDisconnectSessions(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Execute()

Disconnect all sessions on a machine.



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
    nameOrId := "nameOrId_example" // string | Machine to disconnect sessions from. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. (optional) (default to true)
    async := true // bool | If `true`, the session will be disconnected as a background task. The task will have JobType DisconnectSessions. When the task is complete it will redirect to GetMachine. The job's Parameters will contain properties: * _Id_ - ID of the machine on which sessions are being disconnected. * _Name_ - Name of the machine on which sessions are being disconnected. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesDisconnectSessions(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesDisconnectSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesDisconnectSessions`: MachineDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesDisconnectSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to disconnect sessions from. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesDisconnectSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the session will be disconnected as a background task. The task will have JobType DisconnectSessions. When the task is complete it will redirect to GetMachine. The job&#39;s Parameters will contain properties: * _Id_ - ID of the machine on which sessions are being disconnected. * _Name_ - Name of the machine on which sessions are being disconnected. | [default to false]

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


## MachinesDoMachineSearch

> MachineResponseModelCollection MachinesDoMachineSearch(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).MachineAndSessionSearchRequestModel(machineAndSessionSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XTimeZone(xTimeZone).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Skip(skip).Execute()

Perform an advanced search for machines.



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
    machineAndSessionSearchRequestModel := *openapiclient.NewMachineAndSessionSearchRequestModel() // MachineAndSessionSearchRequestModel | Specifies the advanced search parameters.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xTimeZone := "xTimeZone_example" // string | Time zone of the client. If specified, must be a valid Windows Id or Utc Offset from IANA (https://www.iana.org/time-zones) time zones.  Example: UTC or +00:00 (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    limit := int32(56) // int32 | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server (optional)
    async := true // bool | If `true`, the search machines will be executed as a background task. The task will have JobType DoMachineSearch. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    skip := int32(56) // int32 | The number of machines to skip. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesDoMachineSearch(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).MachineAndSessionSearchRequestModel(machineAndSessionSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XTimeZone(xTimeZone).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesDoMachineSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesDoMachineSearch`: MachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesDoMachineSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachinesDoMachineSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **machineAndSessionSearchRequestModel** | [**MachineAndSessionSearchRequestModel**](MachineAndSessionSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xTimeZone** | **string** | Time zone of the client. If specified, must be a valid Windows Id or Utc Offset from IANA (https://www.iana.org/time-zones) time zones.  Example: UTC or +00:00 | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server | 
 **async** | **bool** | If &#x60;true&#x60;, the search machines will be executed as a background task. The task will have JobType DoMachineSearch. When the task is complete it will redirect to GetJobResults. | [default to false]
 **skip** | **int32** | The number of machines to skip. | 

### Return type

[**MachineResponseModelCollection**](MachineResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachine

> MachineDetailResponseModel MachinesGetMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()

Get details of a single machine which belongs to this site, or have registered but are not yet configured in this site



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    fields := "Uid,AgentVersion,AllocationType" // string | The requested fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachine`: MachineDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **fields** | **string** | The requested fields. | 

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


## MachinesGetMachineApplications

> MachineApplicationResponseModelCollection MachinesGetMachineApplications(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the list of applications on a machine.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachineApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachineApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachineApplications`: MachineApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachineApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**MachineApplicationResponseModelCollection**](MachineApplicationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachineDeliveryGroup

> DeliveryGroupDetailResponseModel MachinesGetMachineDeliveryGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the delivery group for a machine.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachineDeliveryGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachineDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachineDeliveryGroup`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachineDeliveryGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DeliveryGroupDetailResponseModel**](DeliveryGroupDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachineDesktop

> DesktopResponseModel MachinesGetMachineDesktop(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the desktop associated with the machine, if any.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachineDesktop(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachineDesktop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachineDesktop`: DesktopResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachineDesktop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineDesktopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DesktopResponseModel**](DesktopResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachineMachineCatalog

> MachineCatalogDetailResponseModel MachinesGetMachineMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the machine catalog for a machine.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachineMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachineMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachineMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachineMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**MachineCatalogDetailResponseModel**](MachineCatalogDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachinePowerActionSchedules

> PowerActionScheduleResponseModelCollection MachinesGetMachinePowerActionSchedules(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the power action schedules associated with a machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachinePowerActionSchedules(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachinePowerActionSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachinePowerActionSchedules`: PowerActionScheduleResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachinePowerActionSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachinePowerActionSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**PowerActionScheduleResponseModelCollection**](PowerActionScheduleResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachineSessions

> SessionResponseModelCollection MachinesGetMachineSessions(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()

Get the list of sessions running on a machine.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    fields := "Id,Uid,ApplicationsInUse" // string | The requested session fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachineSessions(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachineSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachineSessions`: SessionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachineSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **fields** | **string** | The requested session fields. | 

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


## MachinesGetMachineStartMenuShortcutIcon

> *os.File MachinesGetMachineStartMenuShortcutIcon(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).ShortcutPath(shortcutPath).IconFormat(iconFormat).Async(async).Execute()

Get a start menu shortcut icon from the machine.



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
    nameOrId := "nameOrId_example" // string | Machine to get the shortcut icon from. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    shortcutPath := "shortcutPath_example" // string | Path to the start menu shortcut. (optional)
    iconFormat := "iconFormat_example" // string | Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`   where:   * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which many clients cannot display properly. (optional)
    async := true // bool | If `true`, the start menu shortcut icon will be queried as a background task. The task will have JobType GetMachineStartMenuShortcutIcon. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the machine from which start menu icon are being obtained, * _Name_ - Name of the machine from which start menu icon are being obtained. * _ShortcutPath_ - Path to the start menu shortcut. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachineStartMenuShortcutIcon(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).ShortcutPath(shortcutPath).IconFormat(iconFormat).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachineStartMenuShortcutIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachineStartMenuShortcutIcon`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachineStartMenuShortcutIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to get the shortcut icon from. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineStartMenuShortcutIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **shortcutPath** | **string** | Path to the start menu shortcut. | 
 **iconFormat** | **string** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;   where:   * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which many clients cannot display properly. | 
 **async** | **bool** | If &#x60;true&#x60;, the start menu shortcut icon will be queried as a background task. The task will have JobType GetMachineStartMenuShortcutIcon. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine from which start menu icon are being obtained, * _Name_ - Name of the machine from which start menu icon are being obtained. * _ShortcutPath_ - Path to the start menu shortcut. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachineStartMenuShortcuts

> StartMenuApplicationResponseModelCollection MachinesGetMachineStartMenuShortcuts(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get start menu shortcuts from the machine.



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
    nameOrId := "nameOrId_example" // string | Machine to get the shortcuts from. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachineStartMenuShortcuts(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachineStartMenuShortcuts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachineStartMenuShortcuts`: StartMenuApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachineStartMenuShortcuts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to get the shortcuts from. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineStartMenuShortcutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**StartMenuApplicationResponseModelCollection**](StartMenuApplicationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachineTags

> TagResponseModelCollection MachinesGetMachineTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()

Get tags associated with a machine.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    fields := "Name,Id,Description" // string | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma','. return all if not specified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachineTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachineTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachineTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachineTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **fields** | **string** | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma&#39;,&#39;. return all if not specified. | 

### Return type

[**TagResponseModelCollection**](TagResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachineUpgradeVersions

> []string MachinesGetMachineUpgradeVersions(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get available upgrade versions for a machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachineUpgradeVersions(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachineUpgradeVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachineUpgradeVersions`: []string
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachineUpgradeVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachineUpgradeVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachines

> MachineResponseModelCollection MachinesGetMachines(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XTimeZone(xTimeZone).Accept(accept).CitrixLocale(citrixLocale).SessionSupport(sessionSupport).Configured(configured).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Execute()

Get all machines in the site.



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
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xTimeZone := "xTimeZone_example" // string | Time zone of the client. If specified, must be a valid Windows Id or Utc Offset from IANA (https://www.iana.org/time-zones) time zones.  Example: UTC or +00:00 (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    sessionSupport := openapiclient.SessionSupport("Unknown") // SessionSupport | Optionally limit the results to machines that are either single or multi-session capable. If not specified, all types of machines are returned. (optional)
    configured := true // bool | Optionally limit the results to machines that are either configured or not. If not specified, only configured machines are returned. (optional) (default to true)
    limit := int32(56) // int32 | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the get machines will be executed as a background task. The task will have JobType GetMachines. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    fields := "Uid,AgentVersion,AllocationType" // string | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma','.              (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachines(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XTimeZone(xTimeZone).Accept(accept).CitrixLocale(citrixLocale).SessionSupport(sessionSupport).Configured(configured).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachines`: MachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xTimeZone** | **string** | Time zone of the client. If specified, must be a valid Windows Id or Utc Offset from IANA (https://www.iana.org/time-zones) time zones.  Example: UTC or +00:00 | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **sessionSupport** | [**SessionSupport**](SessionSupport.md) | Optionally limit the results to machines that are either single or multi-session capable. If not specified, all types of machines are returned. | 
 **configured** | **bool** | Optionally limit the results to machines that are either configured or not. If not specified, only configured machines are returned. | [default to true]
 **limit** | **int32** | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the get machines will be executed as a background task. The task will have JobType GetMachines. When the task is complete it will redirect to GetJobResults. | [default to false]
 **fields** | **string** | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma&#39;,&#39;.              | 

### Return type

[**MachineResponseModelCollection**](MachineResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachinesAdministrators

> AdministratorResponseModelCollection MachinesGetMachinesAdministrators(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get administrators who can administer a machine



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
    nameOrId := "nameOrId_example" // string | SamName, UPN, or SID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachinesAdministrators(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachinesAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachinesAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachinesAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | SamName, UPN, or SID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachinesAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AdministratorResponseModelCollection**](AdministratorResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetMachinesV2

> MachineResponseModelCollection MachinesGetMachinesV2(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XTimeZone(xTimeZone).Accept(accept).CitrixLocale(citrixLocale).SessionSupport(sessionSupport).Configured(configured).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Execute()

The V2 version of get all machines in the site.



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
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xTimeZone := "xTimeZone_example" // string | Time zone of the client. If specified, must be a valid Windows Id or Utc Offset from IANA (https://www.iana.org/time-zones) time zones.  Example: UTC or +00:00 (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    sessionSupport := openapiclient.SessionSupport("Unknown") // SessionSupport | Optionally limit the results to machines that are either single or multi-session capable. If not specified, all types of machines are returned. (optional)
    configured := true // bool | Optionally limit the results to machines that are either configured or not. If not specified, only configured machines are returned. (optional) (default to true)
    limit := int32(56) // int32 | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the get machines will be executed as a background task. The task will have JobType GetMachines. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    fields := "Uid,AgentVersion,AllocationType" // string | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma','.              (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetMachinesV2(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XTimeZone(xTimeZone).Accept(accept).CitrixLocale(citrixLocale).SessionSupport(sessionSupport).Configured(configured).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetMachinesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetMachinesV2`: MachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetMachinesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetMachinesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xTimeZone** | **string** | Time zone of the client. If specified, must be a valid Windows Id or Utc Offset from IANA (https://www.iana.org/time-zones) time zones.  Example: UTC or +00:00 | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **sessionSupport** | [**SessionSupport**](SessionSupport.md) | Optionally limit the results to machines that are either single or multi-session capable. If not specified, all types of machines are returned. | 
 **configured** | **bool** | Optionally limit the results to machines that are either configured or not. If not specified, only configured machines are returned. | [default to true]
 **limit** | **int32** | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the get machines will be executed as a background task. The task will have JobType GetMachines. When the task is complete it will redirect to GetJobResults. | [default to false]
 **fields** | **string** | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma&#39;,&#39;.              | 

### Return type

[**MachineResponseModelCollection**](MachineResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesGetTestMachineReport

> MachineTestResponseModel MachinesGetTestMachineReport(ctx, nameOrId, reportId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get Cloud Health Check Report on a VDA machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine to test. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    reportId := "reportId_example" // string | ID of the Cloud Health Check Report.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesGetTestMachineReport(context.Background(), nameOrId, reportId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesGetTestMachineReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesGetTestMachineReport`: MachineTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesGetTestMachineReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine to test. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 
**reportId** | **string** | ID of the Cloud Health Check Report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesGetTestMachineReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**MachineTestResponseModel**](MachineTestResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesImportFileTypes

> FtaResponseModelCollection MachinesImportFileTypes(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Import file type associations from the machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the file types will be imported as a background task. The task will have JobType ImportFTA. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties: * _MachineId_ - ID of the machine * _MachineName_ - Name of the machine (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesImportFileTypes(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesImportFileTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesImportFileTypes`: FtaResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesImportFileTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesImportFileTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the file types will be imported as a background task. The task will have JobType ImportFTA. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties: * _MachineId_ - ID of the machine * _MachineName_ - Name of the machine | [default to false]

### Return type

[**FtaResponseModelCollection**](FtaResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesLogoffMachineSessions

> MachineDetailResponseModel MachinesLogoffMachineSessions(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Execute()

Logoff all sessions on a machine.



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
    nameOrId := "nameOrId_example" // string | Machine to log sessions off from. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. (optional) (default to true)
    async := true // bool | If `true`, the sessions will be logged off as a background task. The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the machine which sessions are being logged off from, * _Name_ - Name of the machine which sessions are being logged off from. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesLogoffMachineSessions(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesLogoffMachineSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesLogoffMachineSessions`: MachineDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesLogoffMachineSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to log sessions off from. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesLogoffMachineSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the sessions will be logged off as a background task. The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the machine which sessions are being logged off from, * _Name_ - Name of the machine which sessions are being logged off from. | [default to false]

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


## MachinesRebootMachine

> MachineDetailResponseModel MachinesRebootMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).DetailResponseRequired(detailResponseRequired).Async(async).Execute()

Reboot a machine.



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
    nameOrId := "nameOrId_example" // string | Machine to reboot. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    force := true // bool | If `true`, the machine will be forced to reboot without waiting for user activity to stop.  This may cause loss of user data if work is not saved. (optional) (default to false)
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. (optional) (default to true)
    async := true // bool | If `true`, the machine will be rebooted as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType RebootMachines. When the task is complete it will redirect to GetMachine. The job's Parameters will contain properties: * _Id_ - ID of the machine being rebooted, * _Name_ - Name of the machine being rebooted. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesRebootMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).DetailResponseRequired(detailResponseRequired).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesRebootMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesRebootMachine`: MachineDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesRebootMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to reboot. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesRebootMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **force** | **bool** | If &#x60;true&#x60;, the machine will be forced to reboot without waiting for user activity to stop.  This may cause loss of user data if work is not saved. | [default to false]
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the machine will be rebooted as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType RebootMachines. When the task is complete it will redirect to GetMachine. The job&#39;s Parameters will contain properties: * _Id_ - ID of the machine being rebooted, * _Name_ - Name of the machine being rebooted. | [default to false]

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


## MachinesRemoveMachine

> MachinesRemoveMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DeleteVm(deleteVm).PurgeDBOnly(purgeDBOnly).DeleteAccount(deleteAccount).Async(async).Execute()

Delete a machine from the site.

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
    nameOrId := "nameOrId_example" // string | Machine to remove.  May be specified by SamName, UPN, or SID. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    deleteVm := true // bool | Specifies whether to delete the virtual machines, or just the machine record in the site database. Only relevant when ProvisioningType Optional; default is `true`. (optional) (default to true)
    purgeDBOnly := true // bool | Specifies whether to only delete virtual machines record that were provisioned in the machine catalog from database, not access hypervisor to delete virtual machines, this is recommended only when hypervisor cannot be accessed.  Optional; default is `false`. (optional) (default to false)
    deleteAccount := openapiclient.MachineAccountDeleteOption("Unknown") // MachineAccountDeleteOption | Specifies what to do with the machine account in AD.  Optional; default is Leave. Only relevant when `deleteVm` is `true`; otherwise, ignored. (optional) (default to "Leave")
    async := true // bool | If `true`, the machine will be deleted as a background task. The task will have JobType DeleteMachine. When the task is complete it will redirect to The job's Parameters will contain properties: * _Id_ - ID of the machine being deleted, * _Name_ - Name of the machine being deleted. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesRemoveMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DeleteVm(deleteVm).PurgeDBOnly(purgeDBOnly).DeleteAccount(deleteAccount).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesRemoveMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to remove.  May be specified by SamName, UPN, or SID. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesRemoveMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **deleteVm** | **bool** | Specifies whether to delete the virtual machines, or just the machine record in the site database. Only relevant when ProvisioningType Optional; default is &#x60;true&#x60;. | [default to true]
 **purgeDBOnly** | **bool** | Specifies whether to only delete virtual machines record that were provisioned in the machine catalog from database, not access hypervisor to delete virtual machines, this is recommended only when hypervisor cannot be accessed.  Optional; default is &#x60;false&#x60;. | [default to false]
 **deleteAccount** | [**MachineAccountDeleteOption**](MachineAccountDeleteOption.md) | Specifies what to do with the machine account in AD.  Optional; default is Leave. Only relevant when &#x60;deleteVm&#x60; is &#x60;true&#x60;; otherwise, ignored. | [default to &quot;Leave&quot;]
 **async** | **bool** | If &#x60;true&#x60;, the machine will be deleted as a background task. The task will have JobType DeleteMachine. When the task is complete it will redirect to The job&#39;s Parameters will contain properties: * _Id_ - ID of the machine being deleted, * _Name_ - Name of the machine being deleted. | [default to false]

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


## MachinesRemoveMachineTags

> MachinesRemoveMachineTags(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Remove a tag from a machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesRemoveMachineTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesRemoveMachineTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 
**tagNameOrId** | **string** | Name or ID of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesRemoveMachineTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. | [default to false]

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


## MachinesRemoveMachineUpgradeSchedule

> MachinesRemoveMachineUpgradeSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).Async(async).Execute()

Remove a pending VDA upgrade schedule for a machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    force := true // bool |  (optional) (default to false)
    async := true // bool | If `true`, the VDA upgrade schedule will be removed as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesRemoveMachineUpgradeSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesRemoveMachineUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesRemoveMachineUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **force** | **bool** |  | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedule will be removed as a background task. | [default to false]

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


## MachinesRemovePowerActionSchedule

> MachinesRemovePowerActionSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Remove pending power action schedules for a machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the power action schedule will be removed as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesRemovePowerActionSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesRemovePowerActionSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesRemovePowerActionScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the power action schedule will be removed as a background task. | [default to false]

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


## MachinesResetProvVMDisk

> MachinesResetProvVMDisk(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Resets the OS disk of persistent VMs to the current master image.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. May be specified by name, ID, or SID. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the machine will be deleted as a background task. The task will have JobType ResetVMDisk. When the task is complete it will redirect to The job's Parameters will contain properties: * _Id_ - ID of the machine being deleted, * _Name_ - Name of the machine being deleted. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesResetProvVMDisk(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesResetProvVMDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. May be specified by name, ID, or SID. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesResetProvVMDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine will be deleted as a background task. The task will have JobType ResetVMDisk. When the task is complete it will redirect to The job&#39;s Parameters will contain properties: * _Id_ - ID of the machine being deleted, * _Name_ - Name of the machine being deleted. | [default to false]

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


## MachinesResumeMachine

> MachineDetailResponseModel MachinesResumeMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Execute()

Resume a machine.



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
    nameOrId := "nameOrId_example" // string | Machine to resume. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. (optional) (default to true)
    async := true // bool | If `true`, the machine will be resumed as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType ResumeMachines. When the task is complete it will redirect to GetMachine. The job's Parameters will contain properties: * _Id_ - ID of the machine being resumed, * _Name_ - Name of the machine being resumed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesResumeMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesResumeMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesResumeMachine`: MachineDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesResumeMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to resume. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesResumeMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the machine will be resumed as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType ResumeMachines. When the task is complete it will redirect to GetMachine. The job&#39;s Parameters will contain properties: * _Id_ - ID of the machine being resumed, * _Name_ - Name of the machine being resumed. | [default to false]

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


## MachinesSetMachineTags

> MachinesSetMachineTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).TagsRequestModel(tagsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Set tags associated with a machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    tagsRequestModel := *openapiclient.NewTagsRequestModel() // TagsRequestModel | Details of the tags to set on the machine.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesSetMachineTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).TagsRequestModel(tagsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesSetMachineTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesSetMachineTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **tagsRequestModel** | [**TagsRequestModel**](TagsRequestModel.md) | Details of the tags to set on the machine. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. | [default to false]

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


## MachinesShutdownMachine

> MachineDetailResponseModel MachinesShutdownMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).DetailResponseRequired(detailResponseRequired).Async(async).Delay(delay).Execute()

Shut down a machine.



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
    nameOrId := "nameOrId_example" // string | Machine to shut down. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    force := true // bool | If `true`, the machine will be forced to shut down without waiting for user activity to stop.  This may cause loss of user data if work is not saved. (optional) (default to false)
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. (optional) (default to true)
    async := true // bool | If `true`, the machine will be shut down as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType ShutDownMachines. When the task is complete it will redirect to GetMachine. The job's Parameters will contain properties: * _Id_ - ID of the machine being shut down, * _Name_ - Name of the machine being shut down. (optional) (default to false)
    delay := int32(56) // int32 | Shutdown a machine after the specified number of seconds delay, If `null`, the machine will be shut down without delay. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesShutdownMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).DetailResponseRequired(detailResponseRequired).Async(async).Delay(delay).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesShutdownMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesShutdownMachine`: MachineDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesShutdownMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to shut down. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesShutdownMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **force** | **bool** | If &#x60;true&#x60;, the machine will be forced to shut down without waiting for user activity to stop.  This may cause loss of user data if work is not saved. | [default to false]
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the machine will be shut down as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType ShutDownMachines. When the task is complete it will redirect to GetMachine. The job&#39;s Parameters will contain properties: * _Id_ - ID of the machine being shut down, * _Name_ - Name of the machine being shut down. | [default to false]
 **delay** | **int32** | Shutdown a machine after the specified number of seconds delay, If &#x60;null&#x60;, the machine will be shut down without delay. | 

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


## MachinesStartMachine

> MachineDetailResponseModel MachinesStartMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Execute()

Start a machine.



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
    nameOrId := "nameOrId_example" // string | Machine to start. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. (optional) (default to true)
    async := true // bool | If `true`, the machine will be started as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType StartMachines. When the task is complete it will redirect to GetMachine. The job's Parameters will contain properties: * _Id_ - ID of the machine being started, * _Name_ - Name of the machine being started. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesStartMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesStartMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesStartMachine`: MachineDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesStartMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to start. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesStartMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the machine will be started as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType StartMachines. When the task is complete it will redirect to GetMachine. The job&#39;s Parameters will contain properties: * _Id_ - ID of the machine being started, * _Name_ - Name of the machine being started. | [default to false]

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


## MachinesSuspendMachine

> MachineDetailResponseModel MachinesSuspendMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Delay(delay).Execute()

Suspend a machine.



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
    nameOrId := "nameOrId_example" // string | Machine to suspend. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. (optional) (default to true)
    async := true // bool | If `true`, the machine will be suspended as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType SuspendMachines. When the task is complete it will redirect to GetMachine. The job's Parameters will contain properties: * _Id_ - ID of the machine being suspended, * _Name_ - Name of the machine being suspended. (optional) (default to false)
    delay := int32(56) // int32 | Suspend a machine after the specified number of seconds delay, If `null`, the machine will be suspended without delay. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesSuspendMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).DetailResponseRequired(detailResponseRequired).Async(async).Delay(delay).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesSuspendMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesSuspendMachine`: MachineDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesSuspendMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to suspend. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesSuspendMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic machine model. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the machine will be suspended as a background task. This is recommended as this operation may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType SuspendMachines. When the task is complete it will redirect to GetMachine. The job&#39;s Parameters will contain properties: * _Id_ - ID of the machine being suspended, * _Name_ - Name of the machine being suspended. | [default to false]
 **delay** | **int32** | Suspend a machine after the specified number of seconds delay, If &#x60;null&#x60;, the machine will be suspended without delay. | 

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


## MachinesTestMachine

> MachineTestResponseModel MachinesTestMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Run Cloud Health Check on a VDA machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine to test. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachinesAPIsDAAS.MachinesTestMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesTestMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachinesTestMachine`: MachineTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachinesAPIsDAAS.MachinesTestMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine to test. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesTestMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**MachineTestResponseModel**](MachineTestResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesUpdateMachineCatalogMachine

> MachinesUpdateMachineCatalogMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMachineRequestModel(updateMachineRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Update a machine.



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
    nameOrId := "nameOrId_example" // string | Machine to update.  May be specified by name, ID, or SID. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    updateMachineRequestModel := *openapiclient.NewUpdateMachineRequestModel() // UpdateMachineRequestModel | Details of the machine to update.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tag will be created as a background task. The task will have JobType UpdateMachineCatalogMachine. When the task is complete it will redirect to \"GetMachine(string)\". The job's Parameters will contain properties:  * _Id_ - ID of the machine being updated, * _Name_ - Name of the machine being updated. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesUpdateMachineCatalogMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMachineRequestModel(updateMachineRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesUpdateMachineCatalogMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Machine to update.  May be specified by name, ID, or SID. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesUpdateMachineCatalogMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **updateMachineRequestModel** | [**UpdateMachineRequestModel**](UpdateMachineRequestModel.md) | Details of the machine to update. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tag will be created as a background task. The task will have JobType UpdateMachineCatalogMachine. When the task is complete it will redirect to \&quot;GetMachine(string)\&quot;. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine being updated, * _Name_ - Name of the machine being updated. | [default to false]

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


## MachinesUpdateMachineUpgradeSchedule

> MachinesUpdateMachineUpgradeSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateUpgradeScheduleRequestModel(updateUpgradeScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Reschedule a pending VDA upgrade schedule for a machine.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine. If param is Name, currently it should get rid of '\\\\' and replace it using '|'. For instance, if a MachineName is \"DomainA\\\\NameB\", the param will be \"DomainA|NameB\".
    updateUpgradeScheduleRequestModel := *openapiclient.NewUpdateUpgradeScheduleRequestModel(time.Now(), int32(123)) // UpdateUpgradeScheduleRequestModel | Details about the updated VDA upgrade schedule.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the VDA upgrade schedule will be updated as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachinesAPIsDAAS.MachinesUpdateMachineUpgradeSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateUpgradeScheduleRequestModel(updateUpgradeScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPIsDAAS.MachinesUpdateMachineUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;. For instance, if a MachineName is \&quot;DomainA\\\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesUpdateMachineUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **updateUpgradeScheduleRequestModel** | [**UpdateUpgradeScheduleRequestModel**](UpdateUpgradeScheduleRequestModel.md) | Details about the updated VDA upgrade schedule. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedule will be updated as a background task. | [default to false]

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

