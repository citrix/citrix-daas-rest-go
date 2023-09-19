# \DeliveryGroupsAPIs 

All URIs are relative to *https://api-us.cloud.com/cvad/manage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliveryGroupsAddDeliveryGroupMachineCatalog**](DeliveryGroupsAPIs .md#DeliveryGroupsAddDeliveryGroupMachineCatalog) | **Post** /DeliveryGroups/{nameOrId}/MachineCatalogs | Associate a machine catalog with a RemotePC delivery group.
[**DeliveryGroupsAddDeliveryGroupTags**](DeliveryGroupsAPIs .md#DeliveryGroupsAddDeliveryGroupTags) | **Post** /DeliveryGroups/{nameOrId}/Tags/{tagNameOrId} | Add a tag to a delivery group.
[**DeliveryGroupsCheckDeliveryGroupExists**](DeliveryGroupsAPIs .md#DeliveryGroupsCheckDeliveryGroupExists) | **Head** /DeliveryGroups/{name} | Check for the existence of a delivery group by name.
[**DeliveryGroupsCheckRebootScheduleNameExists**](DeliveryGroupsAPIs .md#DeliveryGroupsCheckRebootScheduleNameExists) | **Head** /DeliveryGroups/{nameOrId}/RebootSchedules/{name} | Check for the existence of a reboot schedule by name.
[**DeliveryGroupsCreateDeliveryGroup**](DeliveryGroupsAPIs .md#DeliveryGroupsCreateDeliveryGroup) | **Post** /DeliveryGroups | Create a delivery group.
[**DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsAPIs .md#DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy) | **Post** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies | Create an advanced access policy associated with a delivery group.
[**DeliveryGroupsCreateDeliveryGroupPowerTimeScheme**](DeliveryGroupsAPIs .md#DeliveryGroupsCreateDeliveryGroupPowerTimeScheme) | **Post** /DeliveryGroups/{nameOrId}/PowerTimeSchemes | Create a power time scheme for a delivery group.
[**DeliveryGroupsCreateDeliveryGroupRebootSchedule**](DeliveryGroupsAPIs .md#DeliveryGroupsCreateDeliveryGroupRebootSchedule) | **Post** /DeliveryGroups/{nameOrId}/RebootSchedules | Create a reboot schedule for a delivery group.
[**DeliveryGroupsDeleteDeliveryGroup**](DeliveryGroupsAPIs .md#DeliveryGroupsDeleteDeliveryGroup) | **Delete** /DeliveryGroups/{nameOrId} | Delete a delivery group.
[**DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsAPIs .md#DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy) | **Delete** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies/{policyId} | Delete an advanced access policy associated with a delivery group.
[**DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme**](DeliveryGroupsAPIs .md#DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme) | **Delete** /DeliveryGroups/{nameOrId}/PowerTimeSchemes/{schemeNameOrId} | Delete a power time scheme from a delivery group.
[**DeliveryGroupsDeleteDeliveryGroupRebootSchedule**](DeliveryGroupsAPIs .md#DeliveryGroupsDeleteDeliveryGroupRebootSchedule) | **Delete** /DeliveryGroups/{nameOrId}/RebootSchedules/{scheduleNameOrId} | Delete a reboot schedule from a delivery group.
[**DeliveryGroupsDoAddApplications**](DeliveryGroupsAPIs .md#DeliveryGroupsDoAddApplications) | **Post** /DeliveryGroups/{nameOrId}/Applications | Add applications and/or application groups to a delivery group.
[**DeliveryGroupsDoAddMachines**](DeliveryGroupsAPIs .md#DeliveryGroupsDoAddMachines) | **Post** /DeliveryGroups/{nameOrId}/Machines | Add machines to a delivery group.
[**DeliveryGroupsDoDeliveryGroupSearch**](DeliveryGroupsAPIs .md#DeliveryGroupsDoDeliveryGroupSearch) | **Post** /DeliveryGroups/$search | Perform an advanced search for deliveryGroups.
[**DeliveryGroupsDoRemoveApplicationGroups**](DeliveryGroupsAPIs .md#DeliveryGroupsDoRemoveApplicationGroups) | **Delete** /DeliveryGroups/{nameOrId}/ApplicationGroups/{appGroupNameOrId} | Remove an application group from a delivery group.
[**DeliveryGroupsDoRemoveApplications**](DeliveryGroupsAPIs .md#DeliveryGroupsDoRemoveApplications) | **Delete** /DeliveryGroups/{nameOrId}/Applications/{appNameOrId} | Remove an application from a delivery group.
[**DeliveryGroupsDoRemoveMachines**](DeliveryGroupsAPIs .md#DeliveryGroupsDoRemoveMachines) | **Delete** /DeliveryGroups/{nameOrId}/Machines/{machineNameOrId} | Remove a machine from a delivery group.
[**DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines) | **Get** /DeliveryGroups/{nameOrId}/ProjectedAutoscaleMachines | Gets the projected number of machines that Autoscale will keep powered on over the specified period
[**DeliveryGroupsGetDeliveryGroup**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroup) | **Get** /DeliveryGroups/{nameOrId} | Get details of a single delivery group.
[**DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies) | **Get** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies | Get the advanced access policies associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy) | **Get** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies/{policyId} | Get the details of a single advanced access policy associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupApplications**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupApplications) | **Get** /DeliveryGroups/{nameOrId}/Applications | Get the applications associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupMachines**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupMachines) | **Get** /DeliveryGroups/{nameOrId}/Machines | Get the machines associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupPowerTimeScheme**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupPowerTimeScheme) | **Get** /DeliveryGroups/{nameOrId}/PowerTimeSchemes/{schemeNameOrId} | Get the details about a single power time scheme associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupPowerTimeSchemes**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupPowerTimeSchemes) | **Get** /DeliveryGroups/{nameOrId}/PowerTimeSchemes | Get the power time schemes associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupRebootSchedule**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupRebootSchedule) | **Get** /DeliveryGroups/{nameOrId}/RebootSchedules/{scheduleNameOrId} | Get the details about a single reboot schedule associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupRebootSchedules**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupRebootSchedules) | **Get** /DeliveryGroups/{nameOrId}/RebootSchedules | Get the reboot schedules associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupStartMenuApplications**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupStartMenuApplications) | **Get** /DeliveryGroups/{nameOrId}/StartMenuApplications | Get the start menu applications from a machine in the delivery group.
[**DeliveryGroupsGetDeliveryGroupTags**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupTags) | **Get** /DeliveryGroups/{nameOrId}/Tags | Get tags associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupTestReport**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupTestReport) | **Get** /DeliveryGroups/{nameOrId}/TestReport | Get the most recent test report of a delivery group.
[**DeliveryGroupsGetDeliveryGroups**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroups) | **Get** /DeliveryGroups | Get all delivery groups.
[**DeliveryGroupsGetDeliveryGroupsAdministrators**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupsAdministrators) | **Get** /DeliveryGroups/{nameOrId}/Administrators | Get administrators who can administer a delivery group.
[**DeliveryGroupsGetDeliveryGroupsApplicationGroups**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupsApplicationGroups) | **Get** /DeliveryGroups/{nameOrId}/ApplicationGroups | Get the application groups associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupsDesktops**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupsDesktops) | **Get** /DeliveryGroups/{nameOrId}/Desktops | Get the published desktops associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails) | **Get** /DeliveryGroups/{nameOrId}/MachineCatalogs/{id} | Get the details of machine catalog associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupsMachineCatalogs**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupsMachineCatalogs) | **Get** /DeliveryGroups/{nameOrId}/MachineCatalogs | Get machine catalogs associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupsUsage**](DeliveryGroupsAPIs .md#DeliveryGroupsGetDeliveryGroupsUsage) | **Get** /DeliveryGroups/{nameOrId}/Usage | Get usage data about a delivery group.
[**DeliveryGroupsPatchDeliveryGroup**](DeliveryGroupsAPIs .md#DeliveryGroupsPatchDeliveryGroup) | **Patch** /DeliveryGroups/{nameOrId} | Modify a delivery group.
[**DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsAPIs .md#DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy) | **Patch** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies/{policyId} | Update an advanced access policy associated with a delivery group.
[**DeliveryGroupsPatchDeliveryGroupPowerTimeScheme**](DeliveryGroupsAPIs .md#DeliveryGroupsPatchDeliveryGroupPowerTimeScheme) | **Patch** /DeliveryGroups/{nameOrId}/PowerTimeSchemes/{schemeNameOrId} | Update a power time scheme for a delivery group.
[**DeliveryGroupsPatchDeliveryGroupRebootSchedule**](DeliveryGroupsAPIs .md#DeliveryGroupsPatchDeliveryGroupRebootSchedule) | **Patch** /DeliveryGroups/{nameOrId}/RebootSchedules/{scheduleNameOrId} | Update a reboot schedule for a delivery group.
[**DeliveryGroupsRemoveDeliveryGroupMachineCatalog**](DeliveryGroupsAPIs .md#DeliveryGroupsRemoveDeliveryGroupMachineCatalog) | **Delete** /DeliveryGroups/{nameOrId}/MachineCatalogs/{catalogNameOrId} | Disassociate a machine catalog from a RemotePC delivery group.
[**DeliveryGroupsRemoveDeliveryGroupTags**](DeliveryGroupsAPIs .md#DeliveryGroupsRemoveDeliveryGroupTags) | **Delete** /DeliveryGroups/{nameOrId}/Tags/{tagNameOrId} | Remove a tag from a delivery group.
[**DeliveryGroupsSetDeliveryGroupTags**](DeliveryGroupsAPIs .md#DeliveryGroupsSetDeliveryGroupTags) | **Put** /DeliveryGroups/{nameOrId}/Tags | Set tags associated with a delivery group.
[**DeliveryGroupsStartDeliveryGroupRebootCycle**](DeliveryGroupsAPIs .md#DeliveryGroupsStartDeliveryGroupRebootCycle) | **Post** /DeliveryGroups/{nameOrId}/$StartRebootCycle | Start a reboot cycle for machines in a delivery group.
[**DeliveryGroupsTestDeliveryGroup**](DeliveryGroupsAPIs .md#DeliveryGroupsTestDeliveryGroup) | **Post** /DeliveryGroups/{nameOrId}/$test | Run tests on a delivery group and create a test report.
[**DeliveryGroupsTestDeliveryGroupExists**](DeliveryGroupsAPIs .md#DeliveryGroupsTestDeliveryGroupExists) | **Post** /DeliveryGroups/$checkName | Test the existence of a delivery group by name.



## DeliveryGroupsAddDeliveryGroupMachineCatalog

> DeliveryGroupDetailResponseModel DeliveryGroupsAddDeliveryGroupMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).PriorityRefRequestModel(priorityRefRequestModel).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Associate a machine catalog with a RemotePC delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    priorityRefRequestModel := *openapiclient.NewPriorityRefRequestModel() // PriorityRefRequestModel | The machine catalog to associate.  The  property must refer to a machine catalog with  equal to `true`.  It may refer to the machine catalog by  or by .   The machine catalog must be available to the same , or a superset of  to whom the delivery group is available.  The Priority property determines which delivery group Remote PC automation will move a qualifying unconfigured machine into when it registers. Priority also determines which delivery group a machine will be published to when a user is assigned to the machine by Remote PC automation. Numerically lower values indicate a higher priority.  If a priority value is not supplied, then the delivery group association is automatically assigned a lower priority than any existing associations.  If a priority value is specified that conflicts with an existing association's priority value, then the new association is inserted with that value and existing associations are renumbered upwards to accommodate it.
    detailResponseRequired := true // bool | Identify if the api need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. (optional) (default to true)
    async := true // bool | If `true`, the association will be updated as a background task. The task will have JobType UpdateDeliveryGroup. When the task is complete it will redirect to GetDeliveryGroup. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group being updated. * _Name_ - Name of the delivery group being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsAddDeliveryGroupMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).PriorityRefRequestModel(priorityRefRequestModel).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsAddDeliveryGroupMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsAddDeliveryGroupMachineCatalog`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsAddDeliveryGroupMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsAddDeliveryGroupMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **priorityRefRequestModel** | [**PriorityRefRequestModel**](PriorityRefRequestModel.md) | The machine catalog to associate.  The  property must refer to a machine catalog with  equal to &#x60;true&#x60;.  It may refer to the machine catalog by  or by .   The machine catalog must be available to the same , or a superset of  to whom the delivery group is available.  The Priority property determines which delivery group Remote PC automation will move a qualifying unconfigured machine into when it registers. Priority also determines which delivery group a machine will be published to when a user is assigned to the machine by Remote PC automation. Numerically lower values indicate a higher priority.  If a priority value is not supplied, then the delivery group association is automatically assigned a lower priority than any existing associations.  If a priority value is specified that conflicts with an existing association&#39;s priority value, then the new association is inserted with that value and existing associations are renumbered upwards to accommodate it. | 
 **detailResponseRequired** | **bool** | Identify if the api need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the association will be updated as a background task. The task will have JobType UpdateDeliveryGroup. When the task is complete it will redirect to GetDeliveryGroup. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group being updated. * _Name_ - Name of the delivery group being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DeliveryGroupDetailResponseModel**](DeliveryGroupDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsAddDeliveryGroupTags

> TagResponseModelCollection DeliveryGroupsAddDeliveryGroupTags(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Add a tag to a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag to add.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsAddDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsAddDeliveryGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsAddDeliveryGroupTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsAddDeliveryGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**tagNameOrId** | **string** | Name or ID of the tag to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsAddDeliveryGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## DeliveryGroupsCheckDeliveryGroupExists

> DeliveryGroupsCheckDeliveryGroupExists(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Check for the existence of a delivery group by name.

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
    name := "name_example" // string | Name of the delivery group.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsCheckDeliveryGroupExists(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsCheckDeliveryGroupExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the delivery group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsCheckDeliveryGroupExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
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


## DeliveryGroupsCheckRebootScheduleNameExists

> DeliveryGroupsCheckRebootScheduleNameExists(ctx, nameOrId, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Check for the existence of a reboot schedule by name.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    name := "name_example" // string | Name of the reboot schedule.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsCheckRebootScheduleNameExists(context.Background(), nameOrId, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsCheckRebootScheduleNameExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**name** | **string** | Name of the reboot schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsCheckRebootScheduleNameExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
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


## DeliveryGroupsCreateDeliveryGroup

> DeliveryGroupDetailResponseModel DeliveryGroupsCreateDeliveryGroup(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateDeliveryGroupRequestModel(createDeliveryGroupRequestModel).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create a delivery group.

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
    createDeliveryGroupRequestModel := *openapiclient.NewCreateDeliveryGroupRequestModel([]openapiclient.DeliveryGroupAddMachinesRequestModel{*openapiclient.NewDeliveryGroupAddMachinesRequestModel("f254399c-6b4c-46fd-86e0-c03ade9135a2")}, "Name of the delivery group") // CreateDeliveryGroupRequestModel | The request
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. (optional) (default to true)
    async := true // bool | Async request (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroup(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateDeliveryGroupRequestModel(createDeliveryGroupRequestModel).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsCreateDeliveryGroup`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsCreateDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **createDeliveryGroupRequestModel** | [**CreateDeliveryGroupRequestModel**](CreateDeliveryGroupRequestModel.md) | The request | 
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. | [default to true]
 **async** | **bool** | Async request | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DeliveryGroupDetailResponseModel**](DeliveryGroupDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy

> AdvancedAccessPolicyResponseModel DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdvancedAccessPolicyRequestModel(advancedAccessPolicyRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create an advanced access policy associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    advancedAccessPolicyRequestModel := *openapiclient.NewAdvancedAccessPolicyRequestModel() // AdvancedAccessPolicyRequestModel | Details of the advanced access policy to create.
    async := true // bool | If `true`, the advanced access policy will be created as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdvancedAccessPolicyRequestModel(advancedAccessPolicyRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy`: AdvancedAccessPolicyResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **advancedAccessPolicyRequestModel** | [**AdvancedAccessPolicyRequestModel**](AdvancedAccessPolicyRequestModel.md) | Details of the advanced access policy to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the advanced access policy will be created as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AdvancedAccessPolicyResponseModel**](AdvancedAccessPolicyResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsCreateDeliveryGroupPowerTimeScheme

> PowerTimeSchemeResponseModel DeliveryGroupsCreateDeliveryGroupPowerTimeScheme(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).PowerTimeSchemeRequestModel(powerTimeSchemeRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create a power time scheme for a delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    powerTimeSchemeRequestModel := *openapiclient.NewPowerTimeSchemeRequestModel() // PowerTimeSchemeRequestModel | Details of the power time scheme to create.
    async := true // bool | If `true`, the power time scheme will be created as a background task. The task will have JobType CreatePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeScheme. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group for which a power time scheme is being created, * _Name_ - Name of the delivery group for which a power time scheme is being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroupPowerTimeScheme(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).PowerTimeSchemeRequestModel(powerTimeSchemeRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsCreateDeliveryGroupPowerTimeScheme`: PowerTimeSchemeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroupPowerTimeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsCreateDeliveryGroupPowerTimeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **powerTimeSchemeRequestModel** | [**PowerTimeSchemeRequestModel**](PowerTimeSchemeRequestModel.md) | Details of the power time scheme to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the power time scheme will be created as a background task. The task will have JobType CreatePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeScheme. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group for which a power time scheme is being created, * _Name_ - Name of the delivery group for which a power time scheme is being created. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**PowerTimeSchemeResponseModel**](PowerTimeSchemeResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsCreateDeliveryGroupRebootSchedule

> RebootScheduleResponseModel DeliveryGroupsCreateDeliveryGroupRebootSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootScheduleRequestModel(rebootScheduleRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create a reboot schedule for a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    rebootScheduleRequestModel := *openapiclient.NewRebootScheduleRequestModel() // RebootScheduleRequestModel | Details of the reboot schedule to create.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroupRebootSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootScheduleRequestModel(rebootScheduleRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsCreateDeliveryGroupRebootSchedule`: RebootScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsCreateDeliveryGroupRebootSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsCreateDeliveryGroupRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **rebootScheduleRequestModel** | [**RebootScheduleRequestModel**](RebootScheduleRequestModel.md) | Details of the reboot schedule to create. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**RebootScheduleResponseModel**](RebootScheduleResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsDeleteDeliveryGroup

> DeliveryGroupsDeleteDeliveryGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Force(force).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete a delivery group.

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
    nameOrId := "nameOrId_example" // string | Delivery group name or ID. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | Async or not (optional) (default to false)
    force := true // bool | Bypass user assignment check (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDeleteDeliveryGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Force(force).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDeleteDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Delivery group name or ID. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDeleteDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | Async or not | [default to false]
 **force** | **bool** | Bypass user assignment check | [default to false]
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


## DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy

> DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy(ctx, nameOrId, policyId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete an advanced access policy associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    policyId := "policyId_example" // string | ID of the advanced access policy.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the advanced access policy will be deleted as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**policyId** | **string** | ID of the advanced access policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the advanced access policy will be deleted as a background task. | [default to false]
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


## DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme

> DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme(ctx, nameOrId, schemeNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete a power time scheme from a delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group.If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    schemeNameOrId := "schemeNameOrId_example" // string | Name or ID of the power time scheme.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the power time scheme will be deleted as a background task. The task will have JobType DeletePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeSchemes. The job's Parameters will contain properties:  *_Id_ - ID of the delivery group for which a power time scheme is being deleted, *_Name_ - Name of the delivery group for which a power time scheme is being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group.If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**schemeNameOrId** | **string** | Name or ID of the power time scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDeleteDeliveryGroupPowerTimeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the power time scheme will be deleted as a background task. The task will have JobType DeletePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeSchemes. The job&#39;s Parameters will contain properties:  *_Id_ - ID of the delivery group for which a power time scheme is being deleted, *_Name_ - Name of the delivery group for which a power time scheme is being deleted. | [default to false]
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


## DeliveryGroupsDeleteDeliveryGroupRebootSchedule

> DeliveryGroupsDeleteDeliveryGroupRebootSchedule(ctx, nameOrId, scheduleNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete a reboot schedule from a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    scheduleNameOrId := "scheduleNameOrId_example" // string | Name or ID of the reboot schedule.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDeleteDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDeleteDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**scheduleNameOrId** | **string** | Name or ID of the reboot schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDeleteDeliveryGroupRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
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


## DeliveryGroupsDoAddApplications

> DeliveryGroupsDoAddApplications(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeliveryGroupAddApplicationsRequestModel(deliveryGroupAddApplicationsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Add applications and/or application groups to a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    deliveryGroupAddApplicationsRequestModel := *openapiclient.NewDeliveryGroupAddApplicationsRequestModel() // DeliveryGroupAddApplicationsRequestModel | Details about the applications and application groups to add.
    async := true // bool | If `true`, the applications will be added to the delivery group as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDoAddApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeliveryGroupAddApplicationsRequestModel(deliveryGroupAddApplicationsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDoAddApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDoAddApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **deliveryGroupAddApplicationsRequestModel** | [**DeliveryGroupAddApplicationsRequestModel**](DeliveryGroupAddApplicationsRequestModel.md) | Details about the applications and application groups to add. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications will be added to the delivery group as a background task. | [default to false]
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


## DeliveryGroupsDoAddMachines

> DeliveryGroupDetailResponseModel DeliveryGroupsDoAddMachines(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeliveryGroupAddMachinesRequestModel(deliveryGroupAddMachinesRequestModel).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Add machines to a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group to add machines to.             If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    deliveryGroupAddMachinesRequestModel := *openapiclient.NewDeliveryGroupAddMachinesRequestModel("f254399c-6b4c-46fd-86e0-c03ade9135a2") // DeliveryGroupAddMachinesRequestModel | Details of the machines to add.
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. (optional) (default to true)
    async := true // bool | If `true`, the machines will be added to the delivery group (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDoAddMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeliveryGroupAddMachinesRequestModel(deliveryGroupAddMachinesRequestModel).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDoAddMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsDoAddMachines`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsDoAddMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group to add machines to.             If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDoAddMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **deliveryGroupAddMachinesRequestModel** | [**DeliveryGroupAddMachinesRequestModel**](DeliveryGroupAddMachinesRequestModel.md) | Details of the machines to add. | 
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the machines will be added to the delivery group | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DeliveryGroupDetailResponseModel**](DeliveryGroupDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsDoDeliveryGroupSearch

> DeliveryGroupResponseModelCollection DeliveryGroupsDoDeliveryGroupSearch(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeliveryGroupSearchRequestModel(deliveryGroupSearchRequestModel).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Perform an advanced search for deliveryGroups.



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
    deliveryGroupSearchRequestModel := *openapiclient.NewDeliveryGroupSearchRequestModel() // DeliveryGroupSearchRequestModel | Specifies the advanced search parameters.
    async := true // bool | If `true`, the delivery groups (and associated objects) will be added as a background task. The task will have JobType DoDeliveryGroupSearch. (optional) (default to false)
    limit := int32(56) // int32 | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDoDeliveryGroupSearch(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeliveryGroupSearchRequestModel(deliveryGroupSearchRequestModel).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDoDeliveryGroupSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsDoDeliveryGroupSearch`: DeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsDoDeliveryGroupSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDoDeliveryGroupSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **deliveryGroupSearchRequestModel** | [**DeliveryGroupSearchRequestModel**](DeliveryGroupSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **async** | **bool** | If &#x60;true&#x60;, the delivery groups (and associated objects) will be added as a background task. The task will have JobType DoDeliveryGroupSearch. | [default to false]
 **limit** | **int32** | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DeliveryGroupResponseModelCollection**](DeliveryGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsDoRemoveApplicationGroups

> DeliveryGroupsDoRemoveApplicationGroups(ctx, nameOrId, appGroupNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Remove an application group from a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    appGroupNameOrId := "appGroupNameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the application groups will be removed to the delivery group as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDoRemoveApplicationGroups(context.Background(), nameOrId, appGroupNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDoRemoveApplicationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**appGroupNameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDoRemoveApplicationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the application groups will be removed to the delivery group as a background task. | [default to false]
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


## DeliveryGroupsDoRemoveApplications

> DeliveryGroupsDoRemoveApplications(ctx, nameOrId, appNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Remove an application from a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    appNameOrId := "appNameOrId_example" // string | Name or ID of the application.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the applications will be removed to the delivery group as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDoRemoveApplications(context.Background(), nameOrId, appNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDoRemoveApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**appNameOrId** | **string** | Name or ID of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDoRemoveApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications will be removed to the delivery group as a background task. | [default to false]
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


## DeliveryGroupsDoRemoveMachines

> DeliveryGroupsDoRemoveMachines(ctx, nameOrId, machineNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Remove a machine from a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    machineNameOrId := "machineNameOrId_example" // string | Name or ID of the machine.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the machines will be removed from the delivery group as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsDoRemoveMachines(context.Background(), nameOrId, machineNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsDoRemoveMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**machineNameOrId** | **string** | Name or ID of the machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsDoRemoveMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the machines will be removed from the delivery group as a background task. | [default to false]
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


## DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines

> ProjectedAutoscaleMachinesResponseModel DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).FromDate(fromDate).NumberOfDays(numberOfDays).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Gets the projected number of machines that Autoscale will keep powered on over the specified period

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    fromDate := time.Now() // time.Time | Gets projected Autoscale machines for the period starting at the specified date (YYYY-MM-DD) that is based on the time zone of the delivery group. (optional)
    numberOfDays := int32(56) // int32 | Gets projected Autoscale machines for the period consisting of the specified number of days. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).FromDate(fromDate).NumberOfDays(numberOfDays).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines`: ProjectedAutoscaleMachinesResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **fromDate** | **time.Time** | Gets projected Autoscale machines for the period starting at the specified date (YYYY-MM-DD) that is based on the time zone of the delivery group. | 
 **numberOfDays** | **int32** | Gets projected Autoscale machines for the period consisting of the specified number of days. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**ProjectedAutoscaleMachinesResponseModel**](ProjectedAutoscaleMachinesResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroup

> DeliveryGroupDetailResponseModel DeliveryGroupsGetDeliveryGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get details of a single delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    fields := "Id,Uid,UserManagement" // string | Optional parameter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroup`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **fields** | **string** | Optional parameter, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
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


## DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies

> AdvancedAccessPolicyResponseModelCollection DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the advanced access policies associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of advanced access policies returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies`: AdvancedAccessPolicyResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupAdvancedAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of advanced access policies returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AdvancedAccessPolicyResponseModelCollection**](AdvancedAccessPolicyResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy

> AdvancedAccessPolicyResponseModel DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy(ctx, nameOrId, policyId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the details of a single advanced access policy associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    policyId := "policyId_example" // string | ID of the advanced access policy.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy`: AdvancedAccessPolicyResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**policyId** | **string** | ID of the advanced access policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupAdvancedAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AdvancedAccessPolicyResponseModel**](AdvancedAccessPolicyResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupApplications

> ApplicationResponseModelCollection DeliveryGroupsGetDeliveryGroupApplications(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the applications associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool |  (optional) (default to false)
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** |  | [default to false]
 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## DeliveryGroupsGetDeliveryGroupMachines

> MachineResponseModelCollection DeliveryGroupsGetDeliveryGroupMachines(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the machines associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupMachines`: MachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## DeliveryGroupsGetDeliveryGroupPowerTimeScheme

> PowerTimeSchemeResponseModel DeliveryGroupsGetDeliveryGroupPowerTimeScheme(ctx, nameOrId, schemeNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the details about a single power time scheme associated with a delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    schemeNameOrId := "schemeNameOrId_example" // string | Name or ID of the power time scheme.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupPowerTimeScheme`: PowerTimeSchemeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupPowerTimeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**schemeNameOrId** | **string** | Name or ID of the power time scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupPowerTimeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**PowerTimeSchemeResponseModel**](PowerTimeSchemeResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupPowerTimeSchemes

> PowerTimeSchemeResponseModelCollection DeliveryGroupsGetDeliveryGroupPowerTimeSchemes(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the power time schemes associated with a delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of power time schemes returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupPowerTimeSchemes(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupPowerTimeSchemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupPowerTimeSchemes`: PowerTimeSchemeResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupPowerTimeSchemes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupPowerTimeSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of power time schemes returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**PowerTimeSchemeResponseModelCollection**](PowerTimeSchemeResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupRebootSchedule

> RebootScheduleResponseModel DeliveryGroupsGetDeliveryGroupRebootSchedule(ctx, nameOrId, scheduleNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the details about a single reboot schedule associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    scheduleNameOrId := "scheduleNameOrId_example" // string | Name or ID of the reboot schedule.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupRebootSchedule`: RebootScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupRebootSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**scheduleNameOrId** | **string** | Name or ID of the reboot schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**RebootScheduleResponseModel**](RebootScheduleResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupRebootSchedules

> RebootScheduleResponseModelCollection DeliveryGroupsGetDeliveryGroupRebootSchedules(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the reboot schedules associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of reboot schedules returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupRebootSchedules(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupRebootSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupRebootSchedules`: RebootScheduleResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupRebootSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupRebootSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of reboot schedules returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**RebootScheduleResponseModelCollection**](RebootScheduleResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupStartMenuApplications

> StartMenuApplicationResponseModelCollection DeliveryGroupsGetDeliveryGroupStartMenuApplications(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagRestriction(tagRestriction).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the start menu applications from a machine in the delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    tagRestriction := "tagRestriction_example" // string | The tag name for restriction. (optional)
    async := true // bool | If `true`, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group from which start menu applications are being obtained, * _Name_ - Name of the delivery group from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupStartMenuApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagRestriction(tagRestriction).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupStartMenuApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupStartMenuApplications`: StartMenuApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupStartMenuApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupStartMenuApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **tagRestriction** | **string** | The tag name for restriction. | 
 **async** | **bool** | If &#x60;true&#x60;, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group from which start menu applications are being obtained, * _Name_ - Name of the delivery group from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
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


## DeliveryGroupsGetDeliveryGroupTags

> TagResponseModelCollection DeliveryGroupsGetDeliveryGroupTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get tags associated with a delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    fields := "Name,Id,Description" // string | Optional parameter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **fields** | **string** | Optional parameter, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## DeliveryGroupsGetDeliveryGroupTestReport

> TestReportResponseModel DeliveryGroupsGetDeliveryGroupTestReport(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the most recent test report of a delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupTestReport(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupTestReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**TestReportResponseModel**](TestReportResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroups

> DeliveryGroupResponseModelCollection DeliveryGroupsGetDeliveryGroups(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get all delivery groups.

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
    adminFolder := "adminFolder_example" // string | Optional folder path (URL-encoded) or ID.  If not specified, all delivery groups will be returned from all folders. (optional)
    async := true // bool |  (optional) (default to false)
    limit := int32(56) // int32 | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroups(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroups`: DeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **adminFolder** | **string** | Optional folder path (URL-encoded) or ID.  If not specified, all delivery groups will be returned from all folders. | 
 **async** | **bool** |  | [default to false]
 **limit** | **int32** | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DeliveryGroupResponseModelCollection**](DeliveryGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupsAdministrators

> AdministratorResponseModelCollection DeliveryGroupsGetDeliveryGroupsAdministrators(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get administrators who can administer a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsAdministrators(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupsAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
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


## DeliveryGroupsGetDeliveryGroupsApplicationGroups

> ApplicationGroupResponseModelCollection DeliveryGroupsGetDeliveryGroupsApplicationGroups(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the application groups associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of application groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsApplicationGroups(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsApplicationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsApplicationGroups`: ApplicationGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsApplicationGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupsApplicationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of application groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**ApplicationGroupResponseModelCollection**](ApplicationGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupsDesktops

> DesktopResponseModelCollection DeliveryGroupsGetDeliveryGroupsDesktops(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the published desktops associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool |  (optional) (default to false)
    limit := int32(56) // int32 | The max number of desktops returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsDesktops(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsDesktops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsDesktops`: DesktopResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsDesktops`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupsDesktopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** |  | [default to false]
 **limit** | **int32** | The max number of desktops returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DesktopResponseModelCollection**](DesktopResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails

> MachineCatalogDetailResponseModel DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails(ctx, nameOrId, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the details of machine catalog associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    id := "id_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {delivery group folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails(context.Background(), nameOrId, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**id** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {delivery group folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupsMachineCatalogDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
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


## DeliveryGroupsGetDeliveryGroupsMachineCatalogs

> MachineCatalogResponseModelCollection DeliveryGroupsGetDeliveryGroupsMachineCatalogs(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get machine catalogs associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Name,FullName,Id" // string | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsMachineCatalogs(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsMachineCatalogs`: MachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupsMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**MachineCatalogResponseModelCollection**](MachineCatalogResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsGetDeliveryGroupsUsage

> DesktopUsageResponseModelCollection DeliveryGroupsGetDeliveryGroupsUsage(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).ContinuationToken(continuationToken).Limit(limit).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get usage data about a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    limit := int32(56) // int32 | The max number of usage data items returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsUsage(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).ContinuationToken(continuationToken).Limit(limit).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsUsage`: DesktopUsageResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsGetDeliveryGroupsUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsGetDeliveryGroupsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **limit** | **int32** | The max number of usage data items returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DesktopUsageResponseModelCollection**](DesktopUsageResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsPatchDeliveryGroup

> DeliveryGroupsPatchDeliveryGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).EditDeliveryGroupRequestModel(editDeliveryGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Modify a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group to update. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    editDeliveryGroupRequestModel := *openapiclient.NewEditDeliveryGroupRequestModel() // EditDeliveryGroupRequestModel | Details about the delivery group modifications requested.
    async := true // bool | If `true`, the delivery group (and associated objects) will be updated as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).EditDeliveryGroupRequestModel(editDeliveryGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group to update. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsPatchDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **editDeliveryGroupRequestModel** | [**EditDeliveryGroupRequestModel**](EditDeliveryGroupRequestModel.md) | Details about the delivery group modifications requested. | 
 **async** | **bool** | If &#x60;true&#x60;, the delivery group (and associated objects) will be updated as a background task. | [default to false]
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


## DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy

> AdvancedAccessPolicyResponseModel DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy(ctx, nameOrId, policyId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdvancedAccessPolicyRequestModel(advancedAccessPolicyRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update an advanced access policy associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    policyId := "policyId_example" // string | ID of the advanced access policy.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    advancedAccessPolicyRequestModel := *openapiclient.NewAdvancedAccessPolicyRequestModel() // AdvancedAccessPolicyRequestModel | Details of the advanced access policy to update.
    async := true // bool | If `true`, the advanced access policy will be updated as a background task.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdvancedAccessPolicyRequestModel(advancedAccessPolicyRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy`: AdvancedAccessPolicyResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**policyId** | **string** | ID of the advanced access policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **advancedAccessPolicyRequestModel** | [**AdvancedAccessPolicyRequestModel**](AdvancedAccessPolicyRequestModel.md) | Details of the advanced access policy to update. | 
 **async** | **bool** | If &#x60;true&#x60;, the advanced access policy will be updated as a background task.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AdvancedAccessPolicyResponseModel**](AdvancedAccessPolicyResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsPatchDeliveryGroupPowerTimeScheme

> PowerTimeSchemeResponseModel DeliveryGroupsPatchDeliveryGroupPowerTimeScheme(ctx, nameOrId, schemeNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).PowerTimeSchemeRequestModel(powerTimeSchemeRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update a power time scheme for a delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    schemeNameOrId := "schemeNameOrId_example" // string | Name or ID of the power time scheme to update.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    powerTimeSchemeRequestModel := *openapiclient.NewPowerTimeSchemeRequestModel() // PowerTimeSchemeRequestModel | Details of the update request.
    async := true // bool | If `true`, the power time scheme will be updated as a background task. The task will have JobType UpdatePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeScheme. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group for which a power time scheme is being updated, * _Name_ - Name of the delivery group for which a power time scheme is being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).PowerTimeSchemeRequestModel(powerTimeSchemeRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsPatchDeliveryGroupPowerTimeScheme`: PowerTimeSchemeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroupPowerTimeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**schemeNameOrId** | **string** | Name or ID of the power time scheme to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsPatchDeliveryGroupPowerTimeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **powerTimeSchemeRequestModel** | [**PowerTimeSchemeRequestModel**](PowerTimeSchemeRequestModel.md) | Details of the update request. | 
 **async** | **bool** | If &#x60;true&#x60;, the power time scheme will be updated as a background task. The task will have JobType UpdatePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeScheme. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group for which a power time scheme is being updated, * _Name_ - Name of the delivery group for which a power time scheme is being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**PowerTimeSchemeResponseModel**](PowerTimeSchemeResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsPatchDeliveryGroupRebootSchedule

> RebootScheduleResponseModel DeliveryGroupsPatchDeliveryGroupRebootSchedule(ctx, nameOrId, scheduleNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootScheduleRequestModel(rebootScheduleRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update a reboot schedule for a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    scheduleNameOrId := "scheduleNameOrId_example" // string | Name or ID of the reboot schedule to update.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    rebootScheduleRequestModel := *openapiclient.NewRebootScheduleRequestModel() // RebootScheduleRequestModel | Details of the update request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootScheduleRequestModel(rebootScheduleRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsPatchDeliveryGroupRebootSchedule`: RebootScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsPatchDeliveryGroupRebootSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**scheduleNameOrId** | **string** | Name or ID of the reboot schedule to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsPatchDeliveryGroupRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **rebootScheduleRequestModel** | [**RebootScheduleRequestModel**](RebootScheduleRequestModel.md) | Details of the update request. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**RebootScheduleResponseModel**](RebootScheduleResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsRemoveDeliveryGroupMachineCatalog

> DeliveryGroupsRemoveDeliveryGroupMachineCatalog(ctx, nameOrId, catalogNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Disassociate a machine catalog from a RemotePC delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    catalogNameOrId := "catalogNameOrId_example" // string | Name or ID of the machine catalog to associate. If the ctalog is present in a catalog folder,             specify the name in this format: {delivery group folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    detailResponseRequired := true // bool | Identify if the api need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. (optional) (default to true)
    async := true // bool | If `true`, the association will be updated as a background task. The task will have JobType UpdateDeliveryGroup. When the task is complete it will redirect to GetDeliveryGroup. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group being updated. * _Name_ - Name of the delivery group being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsRemoveDeliveryGroupMachineCatalog(context.Background(), nameOrId, catalogNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsRemoveDeliveryGroupMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**catalogNameOrId** | **string** | Name or ID of the machine catalog to associate. If the ctalog is present in a catalog folder,             specify the name in this format: {delivery group folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsRemoveDeliveryGroupMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **detailResponseRequired** | **bool** | Identify if the api need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the association will be updated as a background task. The task will have JobType UpdateDeliveryGroup. When the task is complete it will redirect to GetDeliveryGroup. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group being updated. * _Name_ - Name of the delivery group being updated. | [default to false]
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


## DeliveryGroupsRemoveDeliveryGroupTags

> DeliveryGroupsRemoveDeliveryGroupTags(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Remove a tag from a delivery group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsRemoveDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsRemoveDeliveryGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**tagNameOrId** | **string** | Name or ID of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsRemoveDeliveryGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. | [default to false]
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


## DeliveryGroupsSetDeliveryGroupTags

> DeliveryGroupsSetDeliveryGroupTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagsRequestModel(tagsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Set tags associated with a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    tagsRequestModel := *openapiclient.NewTagsRequestModel() // TagsRequestModel | Details of the tags to set on the delivery group.
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsSetDeliveryGroupTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagsRequestModel(tagsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsSetDeliveryGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsSetDeliveryGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **tagsRequestModel** | [**TagsRequestModel**](TagsRequestModel.md) | Details of the tags to set on the delivery group. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. | [default to false]
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


## DeliveryGroupsStartDeliveryGroupRebootCycle

> DeliveryGroupsStartDeliveryGroupRebootCycle(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootCycleRequestModel(rebootCycleRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Start a reboot cycle for machines in a delivery group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    rebootCycleRequestModel := *openapiclient.NewRebootCycleRequestModel() // RebootCycleRequestModel | The request.
    async := true // bool | If `true`, reboot cycle will be started as a background task. The task will have JobType StartDeliveryGroupRebootCycle. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsStartDeliveryGroupRebootCycle(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootCycleRequestModel(rebootCycleRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsStartDeliveryGroupRebootCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsStartDeliveryGroupRebootCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **rebootCycleRequestModel** | [**RebootCycleRequestModel**](RebootCycleRequestModel.md) | The request. | 
 **async** | **bool** | If &#x60;true&#x60;, reboot cycle will be started as a background task. The task will have JobType StartDeliveryGroupRebootCycle. | [default to false]
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


## DeliveryGroupsTestDeliveryGroup

> DeliveryGroupTestResponseModel DeliveryGroupsTestDeliveryGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Run tests on a delivery group and create a test report.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery group to test. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the tests will run as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsTestDeliveryGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsTestDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTestDeliveryGroup`: DeliveryGroupTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIs .DeliveryGroupsTestDeliveryGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group to test. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTestDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**DeliveryGroupTestResponseModel**](DeliveryGroupTestResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryGroupsTestDeliveryGroupExists

> DeliveryGroupsTestDeliveryGroupExists(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DesktopGroupNameCheckRequestModel(desktopGroupNameCheckRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Test the existence of a delivery group by name.

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
    desktopGroupNameCheckRequestModel := *openapiclient.NewDesktopGroupNameCheckRequestModel() // DesktopGroupNameCheckRequestModel | Name of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsAPIs .DeliveryGroupsTestDeliveryGroupExists(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DesktopGroupNameCheckRequestModel(desktopGroupNameCheckRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIs .DeliveryGroupsTestDeliveryGroupExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTestDeliveryGroupExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **desktopGroupNameCheckRequestModel** | [**DesktopGroupNameCheckRequestModel**](DesktopGroupNameCheckRequestModel.md) | Name of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
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

