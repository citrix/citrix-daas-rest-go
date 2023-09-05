# \DeliveryGroupsTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliveryGroupsTPAddDeliveryGroupMachineCatalog**](DeliveryGroupsTPApi.md#DeliveryGroupsTPAddDeliveryGroupMachineCatalog) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/MachineCatalogs | Associate a machine catalog with a RemotePC delivery group.
[**DeliveryGroupsTPAddDeliveryGroupTags**](DeliveryGroupsTPApi.md#DeliveryGroupsTPAddDeliveryGroupTags) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Tags/{tagNameOrId} | Add a tag to a delivery group.
[**DeliveryGroupsTPCheckDeliveryGroupExists**](DeliveryGroupsTPApi.md#DeliveryGroupsTPCheckDeliveryGroupExists) | **Head** /techpreview/{customerid}/{siteid}/DeliveryGroups/{name} | Check for the existence of a delivery group by name.
[**DeliveryGroupsTPCheckRebootScheduleNameExists**](DeliveryGroupsTPApi.md#DeliveryGroupsTPCheckRebootScheduleNameExists) | **Head** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/RebootSchedules/{name} | Check for the existence of a reboot schedule by name.
[**DeliveryGroupsTPCreateDeliveryGroup**](DeliveryGroupsTPApi.md#DeliveryGroupsTPCreateDeliveryGroup) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups | Create a delivery group.
[**DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsTPApi.md#DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/AdvancedAccessPolicies | Create an advanced access policy associated with a delivery group.
[**DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme**](DeliveryGroupsTPApi.md#DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/PowerTimeSchemes | Create a power time scheme for a delivery group.
[**DeliveryGroupsTPCreateDeliveryGroupRebootSchedule**](DeliveryGroupsTPApi.md#DeliveryGroupsTPCreateDeliveryGroupRebootSchedule) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/RebootSchedules | Create a reboot schedule for a delivery group.
[**DeliveryGroupsTPDeleteDeliveryGroup**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDeleteDeliveryGroup) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId} | Delete a delivery group.
[**DeliveryGroupsTPDeleteDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDeleteDeliveryGroupAdvancedAccessPolicy) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/AdvancedAccessPolicies/{policyId} | Delete an advanced access policy associated with a delivery group.
[**DeliveryGroupsTPDeleteDeliveryGroupPowerTimeScheme**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDeleteDeliveryGroupPowerTimeScheme) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/PowerTimeSchemes/{schemeNameOrId} | Delete a power time scheme from a delivery group.
[**DeliveryGroupsTPDeleteDeliveryGroupRebootSchedule**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDeleteDeliveryGroupRebootSchedule) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/RebootSchedules/{scheduleNameOrId} | Delete a reboot schedule from a delivery group.
[**DeliveryGroupsTPDoAddApplications**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDoAddApplications) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Applications | Add applications and/or application groups to a delivery group.
[**DeliveryGroupsTPDoAddMachines**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDoAddMachines) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Machines | Add machines to a delivery group.
[**DeliveryGroupsTPDoDeliveryGroupSearch**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDoDeliveryGroupSearch) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/$search | Perform an advanced search for deliveryGroups.
[**DeliveryGroupsTPDoRemoveApplicationGroups**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDoRemoveApplicationGroups) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/ApplicationGroups/{appGroupNameOrId} | Remove an application group from a delivery group.
[**DeliveryGroupsTPDoRemoveApplications**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDoRemoveApplications) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Applications/{appNameOrId} | Remove an application from a delivery group.
[**DeliveryGroupsTPDoRemoveMachines**](DeliveryGroupsTPApi.md#DeliveryGroupsTPDoRemoveMachines) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Machines/{machineNameOrId} | Remove a machine from a delivery group.
[**DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/ProjectedAutoscaleMachines | Gets the projected number of machines that Autoscale will keep powered on over the specified period
[**DeliveryGroupsTPGetDeliveryGroup**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroup) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId} | Get details of a single delivery group.
[**DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/AdvancedAccessPolicies | Get the advanced access policies associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/AdvancedAccessPolicies/{policyId} | Get the details of a single advanced access policy associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupApplications**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupApplications) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Applications | Get the applications associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupMachines**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupMachines) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Machines | Get the machines associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/PowerTimeSchemes/{schemeNameOrId} | Get the details about a single power time scheme associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/PowerTimeSchemes | Get the power time schemes associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupRebootSchedule**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupRebootSchedule) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/RebootSchedules/{scheduleNameOrId} | Get the details about a single reboot schedule associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupRebootSchedules**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupRebootSchedules) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/RebootSchedules | Get the reboot schedules associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupStartMenuApplications**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupStartMenuApplications) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/StartMenuApplications | Get the start menu applications from a machine in the delivery group.
[**DeliveryGroupsTPGetDeliveryGroupTags**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupTags) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Tags | Get tags associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupTestReport**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupTestReport) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/TestReport | Get the most recent test report of a delivery group.
[**DeliveryGroupsTPGetDeliveryGroups**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroups) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups | Get all delivery groups.
[**DeliveryGroupsTPGetDeliveryGroupsAdministrators**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupsAdministrators) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Administrators | Get administrators who can administer a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupsApplicationGroups**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupsApplicationGroups) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/ApplicationGroups | Get the application groups associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupsDesktops**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupsDesktops) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Desktops | Get the published desktops associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/MachineCatalogs/{id} | Get the details of machine catalog associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/MachineCatalogs | Get machine catalogs associated with a delivery group.
[**DeliveryGroupsTPGetDeliveryGroupsUsage**](DeliveryGroupsTPApi.md#DeliveryGroupsTPGetDeliveryGroupsUsage) | **Get** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Usage | Get usage data about a delivery group.
[**DeliveryGroupsTPPatchDeliveryGroup**](DeliveryGroupsTPApi.md#DeliveryGroupsTPPatchDeliveryGroup) | **Patch** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId} | Modify a delivery group.
[**DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy**](DeliveryGroupsTPApi.md#DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy) | **Patch** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/AdvancedAccessPolicies/{policyId} | Update an advanced access policy associated with a delivery group.
[**DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme**](DeliveryGroupsTPApi.md#DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme) | **Patch** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/PowerTimeSchemes/{schemeNameOrId} | Update a power time scheme for a delivery group.
[**DeliveryGroupsTPPatchDeliveryGroupRebootSchedule**](DeliveryGroupsTPApi.md#DeliveryGroupsTPPatchDeliveryGroupRebootSchedule) | **Patch** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/RebootSchedules/{scheduleNameOrId} | Update a reboot schedule for a delivery group.
[**DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog**](DeliveryGroupsTPApi.md#DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/MachineCatalogs/{catalogNameOrId} | Disassociate a machine catalog from a RemotePC delivery group.
[**DeliveryGroupsTPRemoveDeliveryGroupTags**](DeliveryGroupsTPApi.md#DeliveryGroupsTPRemoveDeliveryGroupTags) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Tags/{tagNameOrId} | Remove a tag from a delivery group.
[**DeliveryGroupsTPSetDeliveryGroupTags**](DeliveryGroupsTPApi.md#DeliveryGroupsTPSetDeliveryGroupTags) | **Put** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/Tags | Set tags associated with a delivery group.
[**DeliveryGroupsTPStartDeliveryGroupRebootCycle**](DeliveryGroupsTPApi.md#DeliveryGroupsTPStartDeliveryGroupRebootCycle) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/$StartRebootCycle | Start a reboot cycle for machines in a delivery group.
[**DeliveryGroupsTPTestDeliveryGroup**](DeliveryGroupsTPApi.md#DeliveryGroupsTPTestDeliveryGroup) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/{nameOrId}/$test | Run tests on a delivery group and create a test report.
[**DeliveryGroupsTPTestDeliveryGroupExists**](DeliveryGroupsTPApi.md#DeliveryGroupsTPTestDeliveryGroupExists) | **Post** /techpreview/{customerid}/{siteid}/DeliveryGroups/$checkName | Test the existence of a delivery group by name.



## DeliveryGroupsTPAddDeliveryGroupMachineCatalog

> DeliveryGroupDetailResponseModel DeliveryGroupsTPAddDeliveryGroupMachineCatalog(ctx, nameOrId, customerid, siteid).Catalog(catalog).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    catalog := *openapiclient.NewPriorityRefRequestModel() // PriorityRefRequestModel | The machine catalog to associate.  The  property must refer to a machine catalog with  equal to `true`.  It may refer to the machine catalog by  or by .   The machine catalog must be available to the same , or a superset of  to whom the delivery group is available.  The Priority property determines which delivery group Remote PC automation will move a qualifying unconfigured machine into when it registers. Priority also determines which delivery group a machine will be published to when a user is assigned to the machine by Remote PC automation. Numerically lower values indicate a higher priority.  If a priority value is not supplied, then the delivery group association is automatically assigned a lower priority than any existing associations.  If a priority value is specified that conflicts with an existing association's priority value, then the new association is inserted with that value and existing associations are renumbered upwards to accommodate it.
    detailResponseRequired := true // bool | Identify if the api need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. (optional) (default to true)
    async := true // bool | If `true`, the association will be updated as a background task. The task will have JobType UpdateDeliveryGroup. When the task is complete it will redirect to GetDeliveryGroup. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group being updated. * _Name_ - Name of the delivery group being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPAddDeliveryGroupMachineCatalog(context.Background(), nameOrId, customerid, siteid).Catalog(catalog).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPAddDeliveryGroupMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPAddDeliveryGroupMachineCatalog`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPAddDeliveryGroupMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPAddDeliveryGroupMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **catalog** | [**PriorityRefRequestModel**](PriorityRefRequestModel.md) | The machine catalog to associate.  The  property must refer to a machine catalog with  equal to &#x60;true&#x60;.  It may refer to the machine catalog by  or by .   The machine catalog must be available to the same , or a superset of  to whom the delivery group is available.  The Priority property determines which delivery group Remote PC automation will move a qualifying unconfigured machine into when it registers. Priority also determines which delivery group a machine will be published to when a user is assigned to the machine by Remote PC automation. Numerically lower values indicate a higher priority.  If a priority value is not supplied, then the delivery group association is automatically assigned a lower priority than any existing associations.  If a priority value is specified that conflicts with an existing association&#39;s priority value, then the new association is inserted with that value and existing associations are renumbered upwards to accommodate it. | 
 **detailResponseRequired** | **bool** | Identify if the api need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the association will be updated as a background task. The task will have JobType UpdateDeliveryGroup. When the task is complete it will redirect to GetDeliveryGroup. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group being updated. * _Name_ - Name of the delivery group being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPAddDeliveryGroupTags

> TagResponseModelCollection DeliveryGroupsTPAddDeliveryGroupTags(ctx, nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPAddDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPAddDeliveryGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPAddDeliveryGroupTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPAddDeliveryGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**tagNameOrId** | **string** | Name or ID of the tag to add. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPAddDeliveryGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPCheckDeliveryGroupExists

> DeliveryGroupsTPCheckDeliveryGroupExists(ctx, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCheckDeliveryGroupExists(context.Background(), name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPCheckDeliveryGroupExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the delivery group. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPCheckDeliveryGroupExistsRequest struct via the builder pattern


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


## DeliveryGroupsTPCheckRebootScheduleNameExists

> DeliveryGroupsTPCheckRebootScheduleNameExists(ctx, nameOrId, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCheckRebootScheduleNameExists(context.Background(), nameOrId, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPCheckRebootScheduleNameExists``: %v\n", err)
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
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPCheckRebootScheduleNameExistsRequest struct via the builder pattern


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


## DeliveryGroupsTPCreateDeliveryGroup

> DeliveryGroupDetailResponseModel DeliveryGroupsTPCreateDeliveryGroup(ctx, customerid, siteid).Request(request).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewCreateDeliveryGroupRequestModel([]openapiclient.DeliveryGroupAddMachinesRequestModel{*openapiclient.NewDeliveryGroupAddMachinesRequestModel("MachineCatalog_example")}, "Name_example") // CreateDeliveryGroupRequestModel | The request
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. (optional) (default to true)
    async := true // bool | Async request (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroup(context.Background(), customerid, siteid).Request(request).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPCreateDeliveryGroup`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPCreateDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateDeliveryGroupRequestModel**](CreateDeliveryGroupRequestModel.md) | The request | 
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. | [default to true]
 **async** | **bool** | Async request | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy

> AdvancedAccessPolicyResponseModel DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewAdvancedAccessPolicyRequestModel() // AdvancedAccessPolicyRequestModel | Details of the advanced access policy to create.
    async := true // bool | If `true`, the advanced access policy will be created as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy`: AdvancedAccessPolicyResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**AdvancedAccessPolicyRequestModel**](AdvancedAccessPolicyRequestModel.md) | Details of the advanced access policy to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the advanced access policy will be created as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme

> PowerTimeSchemeResponseModel DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewPowerTimeSchemeRequestModel() // PowerTimeSchemeRequestModel | Details of the power time scheme to create.
    async := true // bool | If `true`, the power time scheme will be created as a background task. The task will have JobType CreatePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeScheme. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group for which a power time scheme is being created, * _Name_ - Name of the delivery group for which a power time scheme is being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme`: PowerTimeSchemeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPCreateDeliveryGroupPowerTimeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**PowerTimeSchemeRequestModel**](PowerTimeSchemeRequestModel.md) | Details of the power time scheme to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the power time scheme will be created as a background task. The task will have JobType CreatePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeScheme. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group for which a power time scheme is being created, * _Name_ - Name of the delivery group for which a power time scheme is being created. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPCreateDeliveryGroupRebootSchedule

> RebootScheduleResponseModel DeliveryGroupsTPCreateDeliveryGroupRebootSchedule(ctx, nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewRebootScheduleRequestModel() // RebootScheduleRequestModel | Details of the reboot schedule to create.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupRebootSchedule(context.Background(), nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPCreateDeliveryGroupRebootSchedule`: RebootScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupRebootSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPCreateDeliveryGroupRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**RebootScheduleRequestModel**](RebootScheduleRequestModel.md) | Details of the reboot schedule to create. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPDeleteDeliveryGroup

> DeliveryGroupsTPDeleteDeliveryGroup(ctx, nameOrId, customerid, siteid).Async(async).Force(force).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | Async or not (optional) (default to false)
    force := true // bool | Bypass user assignment check (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroup(context.Background(), nameOrId, customerid, siteid).Async(async).Force(force).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Delivery group name or ID. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDeleteDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | Async or not | [default to false]
 **force** | **bool** | Bypass user assignment check | [default to false]
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


## DeliveryGroupsTPDeleteDeliveryGroupAdvancedAccessPolicy

> DeliveryGroupsTPDeleteDeliveryGroupAdvancedAccessPolicy(ctx, nameOrId, policyId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the advanced access policy will be deleted as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
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
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDeleteDeliveryGroupAdvancedAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the advanced access policy will be deleted as a background task. | [default to false]
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


## DeliveryGroupsTPDeleteDeliveryGroupPowerTimeScheme

> DeliveryGroupsTPDeleteDeliveryGroupPowerTimeScheme(ctx, nameOrId, schemeNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the power time scheme will be deleted as a background task. The task will have JobType DeletePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeSchemes. The job's Parameters will contain properties:  *_Id_ - ID of the delivery group for which a power time scheme is being deleted, *_Name_ - Name of the delivery group for which a power time scheme is being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroupPowerTimeScheme``: %v\n", err)
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
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDeleteDeliveryGroupPowerTimeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the power time scheme will be deleted as a background task. The task will have JobType DeletePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeSchemes. The job&#39;s Parameters will contain properties:  *_Id_ - ID of the delivery group for which a power time scheme is being deleted, *_Name_ - Name of the delivery group for which a power time scheme is being deleted. | [default to false]
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


## DeliveryGroupsTPDeleteDeliveryGroupRebootSchedule

> DeliveryGroupsTPDeleteDeliveryGroupRebootSchedule(ctx, nameOrId, scheduleNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroupRebootSchedule``: %v\n", err)
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
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDeleteDeliveryGroupRebootScheduleRequest struct via the builder pattern


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


## DeliveryGroupsTPDoAddApplications

> DeliveryGroupsTPDoAddApplications(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewDeliveryGroupAddApplicationsRequestModel() // DeliveryGroupAddApplicationsRequestModel | Details about the applications and application groups to add.
    async := true // bool | If `true`, the applications will be added to the delivery group as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoAddApplications(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDoAddApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDoAddApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**DeliveryGroupAddApplicationsRequestModel**](DeliveryGroupAddApplicationsRequestModel.md) | Details about the applications and application groups to add. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications will be added to the delivery group as a background task. | [default to false]
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


## DeliveryGroupsTPDoAddMachines

> DeliveryGroupDetailResponseModel DeliveryGroupsTPDoAddMachines(ctx, nameOrId, customerid, siteid).Request(request).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewDeliveryGroupAddMachinesRequestModel("MachineCatalog_example") // DeliveryGroupAddMachinesRequestModel | Details of the machines to add.
    detailResponseRequired := true // bool | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. (optional) (default to true)
    async := true // bool | If `true`, the machines will be added to the delivery group (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoAddMachines(context.Background(), nameOrId, customerid, siteid).Request(request).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDoAddMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPDoAddMachines`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPDoAddMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group to add machines to.             If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDoAddMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**DeliveryGroupAddMachinesRequestModel**](DeliveryGroupAddMachinesRequestModel.md) | Details of the machines to add. | 
 **detailResponseRequired** | **bool** | Identify if the API need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the machines will be added to the delivery group | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPDoDeliveryGroupSearch

> DeliveryGroupResponseModelCollection DeliveryGroupsTPDoDeliveryGroupSearch(ctx, customerid, siteid).SearchRequest(searchRequest).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    searchRequest := *openapiclient.NewDeliveryGroupSearchRequestModel() // DeliveryGroupSearchRequestModel | Specifies the advanced search parameters.
    async := true // bool | If `true`, the delivery groups (and associated objects) will be added as a background task. The task will have JobType DoDeliveryGroupSearch. (optional) (default to false)
    limit := int32(56) // int32 | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoDeliveryGroupSearch(context.Background(), customerid, siteid).SearchRequest(searchRequest).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDoDeliveryGroupSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPDoDeliveryGroupSearch`: DeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPDoDeliveryGroupSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDoDeliveryGroupSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchRequest** | [**DeliveryGroupSearchRequestModel**](DeliveryGroupSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **async** | **bool** | If &#x60;true&#x60;, the delivery groups (and associated objects) will be added as a background task. The task will have JobType DoDeliveryGroupSearch. | [default to false]
 **limit** | **int32** | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPDoRemoveApplicationGroups

> DeliveryGroupsTPDoRemoveApplicationGroups(ctx, nameOrId, appGroupNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the application groups will be removed to the delivery group as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoRemoveApplicationGroups(context.Background(), nameOrId, appGroupNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDoRemoveApplicationGroups``: %v\n", err)
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
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDoRemoveApplicationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the application groups will be removed to the delivery group as a background task. | [default to false]
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


## DeliveryGroupsTPDoRemoveApplications

> DeliveryGroupsTPDoRemoveApplications(ctx, nameOrId, appNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the applications will be removed to the delivery group as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoRemoveApplications(context.Background(), nameOrId, appNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDoRemoveApplications``: %v\n", err)
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
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDoRemoveApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the applications will be removed to the delivery group as a background task. | [default to false]
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


## DeliveryGroupsTPDoRemoveMachines

> DeliveryGroupsTPDoRemoveMachines(ctx, nameOrId, machineNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the machines will be removed from the delivery group as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoRemoveMachines(context.Background(), nameOrId, machineNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPDoRemoveMachines``: %v\n", err)
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
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPDoRemoveMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the machines will be removed from the delivery group as a background task. | [default to false]
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


## DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines

> ProjectedAutoscaleMachinesResponseModel DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines(ctx, nameOrId, customerid, siteid).FromDate(fromDate).NumberOfDays(numberOfDays).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    fromDate := time.Now() // time.Time | Gets projected Autoscale machines for the period starting at the specified date (YYYY-MM-DD) that is based on the time zone of the delivery group. (optional)
    numberOfDays := int32(56) // int32 | Gets projected Autoscale machines for the period consisting of the specified number of days. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines(context.Background(), nameOrId, customerid, siteid).FromDate(fromDate).NumberOfDays(numberOfDays).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines`: ProjectedAutoscaleMachinesResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromDate** | **time.Time** | Gets projected Autoscale machines for the period starting at the specified date (YYYY-MM-DD) that is based on the time zone of the delivery group. | 
 **numberOfDays** | **int32** | Gets projected Autoscale machines for the period consisting of the specified number of days. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroup

> DeliveryGroupDetailResponseModel DeliveryGroupsTPGetDeliveryGroup(ctx, nameOrId, customerid, siteid).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    fields := "fields_example" // string | Optional parameter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroup(context.Background(), nameOrId, customerid, siteid).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroup`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **string** | Optional parameter, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies

> AdvancedAccessPolicyResponseModelCollection DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of advanced access policies returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies`: AdvancedAccessPolicyResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupAdvancedAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of advanced access policies returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy

> AdvancedAccessPolicyResponseModel DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy(ctx, nameOrId, policyId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy`: AdvancedAccessPolicyResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**policyId** | **string** | ID of the advanced access policy. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupApplications

> ApplicationResponseModelCollection DeliveryGroupsTPGetDeliveryGroupApplications(ctx, nameOrId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool |  (optional) (default to false)
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupApplications(context.Background(), nameOrId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** |  | [default to false]
 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
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


## DeliveryGroupsTPGetDeliveryGroupMachines

> MachineResponseModelCollection DeliveryGroupsTPGetDeliveryGroupMachines(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupMachines(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupMachines`: MachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme

> PowerTimeSchemeResponseModel DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme(ctx, nameOrId, schemeNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme`: PowerTimeSchemeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**schemeNameOrId** | **string** | Name or ID of the power time scheme. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupPowerTimeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes

> PowerTimeSchemeResponseModelCollection DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of power time schemes returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes`: PowerTimeSchemeResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupPowerTimeSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of power time schemes returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupRebootSchedule

> RebootScheduleResponseModel DeliveryGroupsTPGetDeliveryGroupRebootSchedule(ctx, nameOrId, scheduleNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupRebootSchedule`: RebootScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupRebootSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**scheduleNameOrId** | **string** | Name or ID of the reboot schedule. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupRebootSchedules

> RebootScheduleResponseModelCollection DeliveryGroupsTPGetDeliveryGroupRebootSchedules(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of reboot schedules returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupRebootSchedules(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupRebootSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupRebootSchedules`: RebootScheduleResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupRebootSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupRebootSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of reboot schedules returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupStartMenuApplications

> StartMenuApplicationResponseModelCollection DeliveryGroupsTPGetDeliveryGroupStartMenuApplications(ctx, nameOrId, customerid, siteid).TagRestriction(tagRestriction).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    tagRestriction := "tagRestriction_example" // string | The tag name for restriction. (optional)
    async := true // bool | If `true`, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group from which start menu applications are being obtained, * _Name_ - Name of the delivery group from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupStartMenuApplications(context.Background(), nameOrId, customerid, siteid).TagRestriction(tagRestriction).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupStartMenuApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupStartMenuApplications`: StartMenuApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupStartMenuApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupStartMenuApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tagRestriction** | **string** | The tag name for restriction. | 
 **async** | **bool** | If &#x60;true&#x60;, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group from which start menu applications are being obtained, * _Name_ - Name of the delivery group from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupTags

> TagResponseModelCollection DeliveryGroupsTPGetDeliveryGroupTags(ctx, nameOrId, customerid, siteid).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    fields := "fields_example" // string | Optional parameter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupTags(context.Background(), nameOrId, customerid, siteid).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **string** | Optional parameter, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupTestReport

> TestReportResponseModel DeliveryGroupsTPGetDeliveryGroupTestReport(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupTestReport(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupTestReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroups

> DeliveryGroupResponseModelCollection DeliveryGroupsTPGetDeliveryGroups(ctx, customerid, siteid).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    adminFolder := "adminFolder_example" // string | Optional folder path (URL-encoded) or ID.  If not specified, all delivery groups will be returned from all folders. (optional)
    async := true // bool |  (optional) (default to false)
    limit := int32(56) // int32 | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroups(context.Background(), customerid, siteid).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroups`: DeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adminFolder** | **string** | Optional folder path (URL-encoded) or ID.  If not specified, all delivery groups will be returned from all folders. | 
 **async** | **bool** |  | [default to false]
 **limit** | **int32** | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupsAdministrators

> AdministratorResponseModelCollection DeliveryGroupsTPGetDeliveryGroupsAdministrators(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsAdministrators(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupsAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupsAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupsApplicationGroups

> ApplicationGroupResponseModelCollection DeliveryGroupsTPGetDeliveryGroupsApplicationGroups(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of application groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsApplicationGroups(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsApplicationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupsApplicationGroups`: ApplicationGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsApplicationGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupsApplicationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of application groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupsDesktops

> DesktopResponseModelCollection DeliveryGroupsTPGetDeliveryGroupsDesktops(ctx, nameOrId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool |  (optional) (default to false)
    limit := int32(56) // int32 | The max number of desktops returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsDesktops(context.Background(), nameOrId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsDesktops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupsDesktops`: DesktopResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsDesktops`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupsDesktopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** |  | [default to false]
 **limit** | **int32** | The max number of desktops returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails

> MachineCatalogDetailResponseModel DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails(ctx, nameOrId, id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails(context.Background(), nameOrId, id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**id** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {delivery group folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs

> MachineCatalogResponseModelCollection DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs`: MachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupsMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPGetDeliveryGroupsUsage

> DesktopUsageResponseModelCollection DeliveryGroupsTPGetDeliveryGroupsUsage(ctx, nameOrId, customerid, siteid).ContinuationToken(continuationToken).Limit(limit).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    limit := int32(56) // int32 | The max number of usage data items returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsUsage(context.Background(), nameOrId, customerid, siteid).ContinuationToken(continuationToken).Limit(limit).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPGetDeliveryGroupsUsage`: DesktopUsageResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPGetDeliveryGroupsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **limit** | **int32** | The max number of usage data items returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPPatchDeliveryGroup

> DeliveryGroupsTPPatchDeliveryGroup(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewEditDeliveryGroupRequestModel() // EditDeliveryGroupRequestModel | Details about the delivery group modifications requested.
    async := true // bool | If `true`, the delivery group (and associated objects) will be updated as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroup(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group to update. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPPatchDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**EditDeliveryGroupRequestModel**](EditDeliveryGroupRequestModel.md) | Details about the delivery group modifications requested. | 
 **async** | **bool** | If &#x60;true&#x60;, the delivery group (and associated objects) will be updated as a background task. | [default to false]
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


## DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy

> AdvancedAccessPolicyResponseModel DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy(ctx, nameOrId, policyId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewAdvancedAccessPolicyRequestModel() // AdvancedAccessPolicyRequestModel | Details of the advanced access policy to update.
    async := true // bool | If `true`, the advanced access policy will be updated as a background task.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy`: AdvancedAccessPolicyResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**policyId** | **string** | ID of the advanced access policy. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**AdvancedAccessPolicyRequestModel**](AdvancedAccessPolicyRequestModel.md) | Details of the advanced access policy to update. | 
 **async** | **bool** | If &#x60;true&#x60;, the advanced access policy will be updated as a background task.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme

> PowerTimeSchemeResponseModel DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme(ctx, nameOrId, schemeNameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewPowerTimeSchemeRequestModel() // PowerTimeSchemeRequestModel | Details of the update request.
    async := true // bool | If `true`, the power time scheme will be updated as a background task. The task will have JobType UpdatePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeScheme. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group for which a power time scheme is being updated, * _Name_ - Name of the delivery group for which a power time scheme is being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme`: PowerTimeSchemeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**schemeNameOrId** | **string** | Name or ID of the power time scheme to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPPatchDeliveryGroupPowerTimeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**PowerTimeSchemeRequestModel**](PowerTimeSchemeRequestModel.md) | Details of the update request. | 
 **async** | **bool** | If &#x60;true&#x60;, the power time scheme will be updated as a background task. The task will have JobType UpdatePowerTimeScheme. When the task is complete it will redirect to GetDeliveryGroupPowerTimeScheme. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group for which a power time scheme is being updated, * _Name_ - Name of the delivery group for which a power time scheme is being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPPatchDeliveryGroupRebootSchedule

> RebootScheduleResponseModel DeliveryGroupsTPPatchDeliveryGroupRebootSchedule(ctx, nameOrId, scheduleNameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewRebootScheduleRequestModel() // RebootScheduleRequestModel | Details of the update request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupRebootSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPPatchDeliveryGroupRebootSchedule`: RebootScheduleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupRebootSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**scheduleNameOrId** | **string** | Name or ID of the reboot schedule to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPPatchDeliveryGroupRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**RebootScheduleRequestModel**](RebootScheduleRequestModel.md) | Details of the update request. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog

> DeliveryGroupDetailResponseModel DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog(ctx, nameOrId, catalogNameOrId, customerid, siteid).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    detailResponseRequired := true // bool | Identify if the api need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. (optional) (default to true)
    async := true // bool | If `true`, the association will be updated as a background task. The task will have JobType UpdateDeliveryGroup. When the task is complete it will redirect to GetDeliveryGroup. The job's Parameters will contain properties:  * _Id_ - ID of the delivery group being updated. * _Name_ - Name of the delivery group being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog(context.Background(), nameOrId, catalogNameOrId, customerid, siteid).DetailResponseRequired(detailResponseRequired).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog`: DeliveryGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**catalogNameOrId** | **string** | Name or ID of the machine catalog to associate. If the ctalog is present in a catalog folder,             specify the name in this format: {delivery group folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPRemoveDeliveryGroupMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **detailResponseRequired** | **bool** | Identify if the api need to return detail info in response model.             Set to false will improve performance with return basic delivery group info. | [default to true]
 **async** | **bool** | If &#x60;true&#x60;, the association will be updated as a background task. The task will have JobType UpdateDeliveryGroup. When the task is complete it will redirect to GetDeliveryGroup. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the delivery group being updated. * _Name_ - Name of the delivery group being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPRemoveDeliveryGroupTags

> DeliveryGroupsTPRemoveDeliveryGroupTags(ctx, nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPRemoveDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPRemoveDeliveryGroupTags``: %v\n", err)
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
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPRemoveDeliveryGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. | [default to false]
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


## DeliveryGroupsTPSetDeliveryGroupTags

> DeliveryGroupsTPSetDeliveryGroupTags(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewTagsRequestModel() // TagsRequestModel | Details of the tags to set on the delivery group.
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPSetDeliveryGroupTags(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPSetDeliveryGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPSetDeliveryGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**TagsRequestModel**](TagsRequestModel.md) | Details of the tags to set on the delivery group. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. | [default to false]
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


## DeliveryGroupsTPStartDeliveryGroupRebootCycle

> DeliveryGroupsTPStartDeliveryGroupRebootCycle(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewRebootCycleRequestModel() // RebootCycleRequestModel | The request.
    async := true // bool | If `true`, reboot cycle will be started as a background task. The task will have JobType StartDeliveryGroupRebootCycle. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPStartDeliveryGroupRebootCycle(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPStartDeliveryGroupRebootCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPStartDeliveryGroupRebootCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**RebootCycleRequestModel**](RebootCycleRequestModel.md) | The request. | 
 **async** | **bool** | If &#x60;true&#x60;, reboot cycle will be started as a background task. The task will have JobType StartDeliveryGroupRebootCycle. | [default to false]
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


## DeliveryGroupsTPTestDeliveryGroup

> DeliveryGroupTestResponseModel DeliveryGroupsTPTestDeliveryGroup(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tests will run as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPTestDeliveryGroup(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPTestDeliveryGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryGroupsTPTestDeliveryGroup`: DeliveryGroupTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryGroupsTPApi.DeliveryGroupsTPTestDeliveryGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery group to test. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryGroupsTPTestDeliveryGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## DeliveryGroupsTPTestDeliveryGroupExists

> DeliveryGroupsTPTestDeliveryGroupExists(ctx, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewDesktopGroupNameCheckRequestModel() // DesktopGroupNameCheckRequestModel | Name of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPTestDeliveryGroupExists(context.Background(), customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryGroupsTPApi.DeliveryGroupsTPTestDeliveryGroupExists``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeliveryGroupsTPTestDeliveryGroupExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**DesktopGroupNameCheckRequestModel**](DesktopGroupNameCheckRequestModel.md) | Name of the delivery group. If the delivery group is present in a delivery group folder,             specify the name in this format: {delivery group folder path plus delivery group name}.             For example, FolderName1|FolderName2|DeliveryGroupName. | 
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

