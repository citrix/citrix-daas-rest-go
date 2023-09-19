# \DeliveryGroupsAPIsDAAS

All URIs are relative to *https://api-us.cloud.com/cvad/manage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliveryGroupsAddDeliveryGroupMachineCatalog**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsAddDeliveryGroupMachineCatalog) | **Post** /DeliveryGroups/{nameOrId}/MachineCatalogs | Associate a machine catalog with a RemotePC delivery group.
[**DeliveryGroupsAddDeliveryGroupTags**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsAddDeliveryGroupTags) | **Post** /DeliveryGroups/{nameOrId}/Tags/{tagNameOrId} | Add a tag to a delivery group.
[**DeliveryGroupsCheckDeliveryGroupExists**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsCheckDeliveryGroupExists) | **Head** /DeliveryGroups/{name} | Check for the existence of a delivery group by name.
[**DeliveryGroupsCheckRebootScheduleNameExists**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsCheckRebootScheduleNameExists) | **Head** /DeliveryGroups/{nameOrId}/RebootSchedules/{name} | Check for the existence of a reboot schedule by name.
[**DeliveryGroupsCreateDeliveryGroup**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsCreateDeliveryGroup) | **Post** /DeliveryGroups | Create a delivery group.
[**DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy) | **Post** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies | Create an advanced access policy associated with a delivery group.
[**DeliveryGroupsCreateDeliveryGroupPowerTimeScheme**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsCreateDeliveryGroupPowerTimeScheme) | **Post** /DeliveryGroups/{nameOrId}/PowerTimeSchemes | Create a power time scheme for a delivery group.
[**DeliveryGroupsCreateDeliveryGroupRebootSchedule**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsCreateDeliveryGroupRebootSchedule) | **Post** /DeliveryGroups/{nameOrId}/RebootSchedules | Create a reboot schedule for a delivery group.
[**DeliveryGroupsDeleteDeliveryGroup**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDeleteDeliveryGroup) | **Delete** /DeliveryGroups/{nameOrId} | Delete a delivery group.
[**DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy) | **Delete** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies/{policyId} | Delete an advanced access policy associated with a delivery group.
[**DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme) | **Delete** /DeliveryGroups/{nameOrId}/PowerTimeSchemes/{schemeNameOrId} | Delete a power time scheme from a delivery group.
[**DeliveryGroupsDeleteDeliveryGroupRebootSchedule**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDeleteDeliveryGroupRebootSchedule) | **Delete** /DeliveryGroups/{nameOrId}/RebootSchedules/{scheduleNameOrId} | Delete a reboot schedule from a delivery group.
[**DeliveryGroupsDoAddApplications**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDoAddApplications) | **Post** /DeliveryGroups/{nameOrId}/Applications | Add applications and/or application groups to a delivery group.
[**DeliveryGroupsDoAddMachines**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDoAddMachines) | **Post** /DeliveryGroups/{nameOrId}/Machines | Add machines to a delivery group.
[**DeliveryGroupsDoDeliveryGroupSearch**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDoDeliveryGroupSearch) | **Post** /DeliveryGroups/$search | Perform an advanced search for deliveryGroups.
[**DeliveryGroupsDoRemoveApplicationGroups**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDoRemoveApplicationGroups) | **Delete** /DeliveryGroups/{nameOrId}/ApplicationGroups/{appGroupNameOrId} | Remove an application group from a delivery group.
[**DeliveryGroupsDoRemoveApplications**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDoRemoveApplications) | **Delete** /DeliveryGroups/{nameOrId}/Applications/{appNameOrId} | Remove an application from a delivery group.
[**DeliveryGroupsDoRemoveMachines**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsDoRemoveMachines) | **Delete** /DeliveryGroups/{nameOrId}/Machines/{machineNameOrId} | Remove a machine from a delivery group.
[**DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines) | **Get** /DeliveryGroups/{nameOrId}/ProjectedAutoscaleMachines | Gets the projected number of machines that Autoscale will keep powered on over the specified period
[**DeliveryGroupsGetDeliveryGroup**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroup) | **Get** /DeliveryGroups/{nameOrId} | Get details of a single delivery group.
[**DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies) | **Get** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies | Get the advanced access policies associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy) | **Get** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies/{policyId} | Get the details of a single advanced access policy associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupApplications**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupApplications) | **Get** /DeliveryGroups/{nameOrId}/Applications | Get the applications associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupMachines**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupMachines) | **Get** /DeliveryGroups/{nameOrId}/Machines | Get the machines associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupPowerTimeScheme**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupPowerTimeScheme) | **Get** /DeliveryGroups/{nameOrId}/PowerTimeSchemes/{schemeNameOrId} | Get the details about a single power time scheme associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupPowerTimeSchemes**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupPowerTimeSchemes) | **Get** /DeliveryGroups/{nameOrId}/PowerTimeSchemes | Get the power time schemes associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupRebootSchedule**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupRebootSchedule) | **Get** /DeliveryGroups/{nameOrId}/RebootSchedules/{scheduleNameOrId} | Get the details about a single reboot schedule associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupRebootSchedules**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupRebootSchedules) | **Get** /DeliveryGroups/{nameOrId}/RebootSchedules | Get the reboot schedules associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupStartMenuApplications**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupStartMenuApplications) | **Get** /DeliveryGroups/{nameOrId}/StartMenuApplications | Get the start menu applications from a machine in the delivery group.
[**DeliveryGroupsGetDeliveryGroupTags**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupTags) | **Get** /DeliveryGroups/{nameOrId}/Tags | Get tags associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupTestReport**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupTestReport) | **Get** /DeliveryGroups/{nameOrId}/TestReport | Get the most recent test report of a delivery group.
[**DeliveryGroupsGetDeliveryGroups**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroups) | **Get** /DeliveryGroups | Get all delivery groups.
[**DeliveryGroupsGetDeliveryGroupsAdministrators**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupsAdministrators) | **Get** /DeliveryGroups/{nameOrId}/Administrators | Get administrators who can administer a delivery group.
[**DeliveryGroupsGetDeliveryGroupsApplicationGroups**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupsApplicationGroups) | **Get** /DeliveryGroups/{nameOrId}/ApplicationGroups | Get the application groups associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupsDesktops**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupsDesktops) | **Get** /DeliveryGroups/{nameOrId}/Desktops | Get the published desktops associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails) | **Get** /DeliveryGroups/{nameOrId}/MachineCatalogs/{id} | Get the details of machine catalog associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupsMachineCatalogs**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupsMachineCatalogs) | **Get** /DeliveryGroups/{nameOrId}/MachineCatalogs | Get machine catalogs associated with a delivery group.
[**DeliveryGroupsGetDeliveryGroupsUsage**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsGetDeliveryGroupsUsage) | **Get** /DeliveryGroups/{nameOrId}/Usage | Get usage data about a delivery group.
[**DeliveryGroupsPatchDeliveryGroup**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsPatchDeliveryGroup) | **Patch** /DeliveryGroups/{nameOrId} | Modify a delivery group.
[**DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy) | **Patch** /DeliveryGroups/{nameOrId}/AdvancedAccessPolicies/{policyId} | Update an advanced access policy associated with a delivery group.
[**DeliveryGroupsPatchDeliveryGroupPowerTimeScheme**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsPatchDeliveryGroupPowerTimeScheme) | **Patch** /DeliveryGroups/{nameOrId}/PowerTimeSchemes/{schemeNameOrId} | Update a power time scheme for a delivery group.
[**DeliveryGroupsPatchDeliveryGroupRebootSchedule**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsPatchDeliveryGroupRebootSchedule) | **Patch** /DeliveryGroups/{nameOrId}/RebootSchedules/{scheduleNameOrId} | Update a reboot schedule for a delivery group.
[**DeliveryGroupsRemoveDeliveryGroupMachineCatalog**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsRemoveDeliveryGroupMachineCatalog) | **Delete** /DeliveryGroups/{nameOrId}/MachineCatalogs/{catalogNameOrId} | Disassociate a machine catalog from a RemotePC delivery group.
[**DeliveryGroupsRemoveDeliveryGroupTags**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsRemoveDeliveryGroupTags) | **Delete** /DeliveryGroups/{nameOrId}/Tags/{tagNameOrId} | Remove a tag from a delivery group.
[**DeliveryGroupsSetDeliveryGroupTags**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsSetDeliveryGroupTags) | **Put** /DeliveryGroups/{nameOrId}/Tags | Set tags associated with a delivery group.
[**DeliveryGroupsStartDeliveryGroupRebootCycle**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsStartDeliveryGroupRebootCycle) | **Post** /DeliveryGroups/{nameOrId}/$StartRebootCycle | Start a reboot cycle for machines in a delivery group.
[**DeliveryGroupsTestDeliveryGroup**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsTestDeliveryGroup) | **Post** /DeliveryGroups/{nameOrId}/$test | Run tests on a delivery group and create a test report.
[**DeliveryGroupsTestDeliveryGroupExists**](DeliveryGroupsAPIsDAAS.md#DeliveryGroupsTestDeliveryGroupExists) | **Post** /DeliveryGroups/$checkName | Test the existence of a delivery group by name.



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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsAddDeliveryGroupMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).PriorityRefRequestModel(priorityRefRequestModel).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsAddDeliveryGroupMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsAddDeliveryGroupMachineCatalog`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsAddDeliveryGroupMachineCatalog`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsAddDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsAddDeliveryGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsAddDeliveryGroupTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsAddDeliveryGroupTags`: %v\n", resp)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCheckDeliveryGroupExists(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsCheckDeliveryGroupExists``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCheckRebootScheduleNameExists(context.Background(), nameOrId, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsCheckRebootScheduleNameExists``: %v\n", err)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroup(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateDeliveryGroupRequestModel(createDeliveryGroupRequestModel).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsCreateDeliveryGroup`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroup`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdvancedAccessPolicyRequestModel(advancedAccessPolicyRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy`: AdvancedAccessPolicyResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupPowerTimeScheme(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).PowerTimeSchemeRequestModel(powerTimeSchemeRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsCreateDeliveryGroupPowerTimeScheme`: PowerTimeSchemeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupPowerTimeScheme`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupRebootSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootScheduleRequestModel(rebootScheduleRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsCreateDeliveryGroupRebootSchedule`: RebootScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupRebootSchedule`: %v\n", resp)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Force(force).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroup``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroupRebootSchedule``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoAddApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeliveryGroupAddApplicationsRequestModel(deliveryGroupAddApplicationsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDoAddApplications``: %v\n", err)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoAddMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeliveryGroupAddMachinesRequestModel(deliveryGroupAddMachinesRequestModel).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDoAddMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsDoAddMachines`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsDoAddMachines`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoDeliveryGroupSearch(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeliveryGroupSearchRequestModel(deliveryGroupSearchRequestModel).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDoDeliveryGroupSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsDoDeliveryGroupSearch`: DeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsDoDeliveryGroupSearch`: %v\n", resp)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoRemoveApplicationGroups(context.Background(), nameOrId, appGroupNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDoRemoveApplicationGroups``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoRemoveApplications(context.Background(), nameOrId, appNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDoRemoveApplications``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoRemoveMachines(context.Background(), nameOrId, machineNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsDoRemoveMachines``: %v\n", err)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).FromDate(fromDate).NumberOfDays(numberOfDays).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines`: ProjectedAutoscaleMachinesResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroup`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroup`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies`: AdvancedAccessPolicyResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy`: AdvancedAccessPolicyResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupApplications`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupMachines`: MachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupMachines`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupPowerTimeScheme`: PowerTimeSchemeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupPowerTimeScheme`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupPowerTimeSchemes(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupPowerTimeSchemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupPowerTimeSchemes`: PowerTimeSchemeResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupPowerTimeSchemes`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupRebootSchedule`: RebootScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupRebootSchedule`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupRebootSchedules(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupRebootSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupRebootSchedules`: RebootScheduleResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupRebootSchedules`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupStartMenuApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagRestriction(tagRestriction).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupStartMenuApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupStartMenuApplications`: StartMenuApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupStartMenuApplications`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupTags`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupTestReport(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupTestReport`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroups(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroups`: DeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroups`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsAdministrators(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsAdministrators`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsApplicationGroups(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsApplicationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsApplicationGroups`: ApplicationGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsApplicationGroups`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsDesktops(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsDesktops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsDesktops`: DesktopResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsDesktops`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails(context.Background(), nameOrId, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsMachineCatalogs(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsMachineCatalogs`: MachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsMachineCatalogs`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsUsage(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).ContinuationToken(continuationToken).Limit(limit).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsGetDeliveryGroupsUsage`: DesktopUsageResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsUsage`: %v\n", resp)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).EditDeliveryGroupRequestModel(editDeliveryGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroup``: %v\n", err)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdvancedAccessPolicyRequestModel(advancedAccessPolicyRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy`: AdvancedAccessPolicyResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).PowerTimeSchemeRequestModel(powerTimeSchemeRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsPatchDeliveryGroupPowerTimeScheme`: PowerTimeSchemeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupPowerTimeScheme`: %v\n", resp)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootScheduleRequestModel(rebootScheduleRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsPatchDeliveryGroupRebootSchedule`: RebootScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupRebootSchedule`: %v\n", resp)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsRemoveDeliveryGroupMachineCatalog(context.Background(), nameOrId, catalogNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsRemoveDeliveryGroupMachineCatalog``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsRemoveDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsRemoveDeliveryGroupTags``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsSetDeliveryGroupTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagsRequestModel(tagsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsSetDeliveryGroupTags``: %v\n", err)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsStartDeliveryGroupRebootCycle(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootCycleRequestModel(rebootCycleRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsStartDeliveryGroupRebootCycle``: %v\n", err)
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
    resp, r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsTestDeliveryGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsTestDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTestDeliveryGroup`: DeliveryGroupTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsAPIsDAAS.DeliveryGroupsTestDeliveryGroup`: %v\n", resp)
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
    r, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsTestDeliveryGroupExists(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DesktopGroupNameCheckRequestModel(desktopGroupNameCheckRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsAPIsDAAS.DeliveryGroupsTestDeliveryGroupExists``: %v\n", err)
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

