# \MachineCatalogsTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MachineCatalogsTPAddMachineCatalogMachine**](MachineCatalogsTPApi.md#MachineCatalogsTPAddMachineCatalogMachine) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/Machines | Add machine to the machine catalog.
[**MachineCatalogsTPAddMachineCatalogMachineAccount**](MachineCatalogsTPApi.md#MachineCatalogsTPAddMachineCatalogMachineAccount) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/MachineAccounts | Add a machine account to the machine catalog.
[**MachineCatalogsTPAddMachineCatalogTags**](MachineCatalogsTPApi.md#MachineCatalogsTPAddMachineCatalogTags) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/Tags/{tagNameOrId} | Add a tag to a machine catalog.
[**MachineCatalogsTPCheckMachineCatalogExists**](MachineCatalogsTPApi.md#MachineCatalogsTPCheckMachineCatalogExists) | **Head** /techpreview/{customerid}/{siteid}/MachineCatalogs/{name} | Check for the existence of a machine catalog by name.
[**MachineCatalogsTPCheckRemotePCEnrollmentScopeExists**](MachineCatalogsTPApi.md#MachineCatalogsTPCheckRemotePCEnrollmentScopeExists) | **Head** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/RemotePCEnrollmentScopes/{name} | Check for the existence of a RemotePC enrollment scope by name.
[**MachineCatalogsTPCreateMachineCatalog**](MachineCatalogsTPApi.md#MachineCatalogsTPCreateMachineCatalog) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs | Create a machine catalog.
[**MachineCatalogsTPCreateMachineCatalogUpgradeSchedule**](MachineCatalogsTPApi.md#MachineCatalogsTPCreateMachineCatalogUpgradeSchedule) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/UpgradeSchedule | Setup a new VDA upgrade schedule for a machine catalog.
[**MachineCatalogsTPDeleteMachineCatalog**](MachineCatalogsTPApi.md#MachineCatalogsTPDeleteMachineCatalog) | **Delete** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId} | Delete a machine catalog.
[**MachineCatalogsTPDoMachineCatalogSearch**](MachineCatalogsTPApi.md#MachineCatalogsTPDoMachineCatalogSearch) | **Post** /techpreview/{customerid}/{siteid}/Machinecatalogs/$search | Perform an advanced search for machinecatalogs.
[**MachineCatalogsTPGetMachineCatalog**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalog) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId} | Get details about a single machine catalog.
[**MachineCatalogsTPGetMachineCatalogDeliveryGroupAssociations**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogDeliveryGroupAssociations) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/DeliveryGroupAssociations | Get delivery group associations of a machine catalog.
[**MachineCatalogsTPGetMachineCatalogLastMasterImage**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogLastMasterImage) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/LastMasterImage | Get the last master VM images used by the machine catalog.
[**MachineCatalogsTPGetMachineCatalogMachineAccounts**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogMachineAccounts) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/MachineAccounts | Get all machine accounts associated with a machine catalog.
[**MachineCatalogsTPGetMachineCatalogMachines**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogMachines) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/Machines | GET the machines of a catalog
[**MachineCatalogsTPGetMachineCatalogMasterImageHistory**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogMasterImageHistory) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/MasterImageHistory | Get the history of master VM images used by the machine catalog.
[**MachineCatalogsTPGetMachineCatalogStartMenuApplications**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogStartMenuApplications) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/StartMenuApplications | Get the start menu applications from a machine in the machine catalog.
[**MachineCatalogsTPGetMachineCatalogTags**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogTags) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/Tags | Get tags associated with a machine catalog.
[**MachineCatalogsTPGetMachineCatalogTestReport**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogTestReport) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/TestReport | Get the most recent test report of a machine catalog.
[**MachineCatalogsTPGetMachineCatalogVDAComponentsAndFeatures**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogVDAComponentsAndFeatures) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/VDAComponentsAndFeatures | Get the components and features of VDAs associated with a machine catalog.
[**MachineCatalogsTPGetMachineCatalogs**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogs) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs | Get all machine catalogs.
[**MachineCatalogsTPGetMachineCatalogsAdministrators**](MachineCatalogsTPApi.md#MachineCatalogsTPGetMachineCatalogsAdministrators) | **Get** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/Administrators | Get administrators who can administer a machine catalog.
[**MachineCatalogsTPRebootMachineCatalogMachines**](MachineCatalogsTPApi.md#MachineCatalogsTPRebootMachineCatalogMachines) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/$RebootMachines | Reboot all machines in a machine catalog.
[**MachineCatalogsTPRemoveMachineCatalogMachine**](MachineCatalogsTPApi.md#MachineCatalogsTPRemoveMachineCatalogMachine) | **Delete** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/Machines/{machine} | Remove a machine from the machine catalog.
[**MachineCatalogsTPRemoveMachineCatalogMachineAccount**](MachineCatalogsTPApi.md#MachineCatalogsTPRemoveMachineCatalogMachineAccount) | **Delete** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/MachineAccounts/{machine} | Remove a machine account from the machine catalog.
[**MachineCatalogsTPRemoveMachineCatalogTags**](MachineCatalogsTPApi.md#MachineCatalogsTPRemoveMachineCatalogTags) | **Delete** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/Tags/{tagNameOrId} | Remove a tag from a machine catalog.
[**MachineCatalogsTPRemoveMachineCatalogUpgradeSchedule**](MachineCatalogsTPApi.md#MachineCatalogsTPRemoveMachineCatalogUpgradeSchedule) | **Delete** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/UpgradeSchedule | Remove a pending VDA upgrade schedule for a machine catalog.
[**MachineCatalogsTPRollbackMachineCatalogProvisioningScheme**](MachineCatalogsTPApi.md#MachineCatalogsTPRollbackMachineCatalogProvisioningScheme) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/$RollbackProvisioningScheme | Rollback the machine catalog&#39;s provisioning scheme.
[**MachineCatalogsTPSetMachineCatalogTags**](MachineCatalogsTPApi.md#MachineCatalogsTPSetMachineCatalogTags) | **Put** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/Tags | Set tags associated with a machine catalog.
[**MachineCatalogsTPTestMachineCatalog**](MachineCatalogsTPApi.md#MachineCatalogsTPTestMachineCatalog) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/$test | Run tests on a machine catalog and create a test report.
[**MachineCatalogsTPTestMachineCatalogExists**](MachineCatalogsTPApi.md#MachineCatalogsTPTestMachineCatalogExists) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/$checkCatalogName | Test the existence of a machine catalog by name.
[**MachineCatalogsTPTestRemotePCEnrollmentScopeExists**](MachineCatalogsTPApi.md#MachineCatalogsTPTestRemotePCEnrollmentScopeExists) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/RemotePCEnrollmentScopes/$checkScopeName | Check for the existence of a RemotePC enrollment scope by name in the request body. Use HttpPost and take name in body as this name may be very long, and cause url too long exception.
[**MachineCatalogsTPUpdateMachineCatalog**](MachineCatalogsTPApi.md#MachineCatalogsTPUpdateMachineCatalog) | **Patch** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId} | Update a machine catalog.
[**MachineCatalogsTPUpdateMachineCatalogExistingMachines**](MachineCatalogsTPApi.md#MachineCatalogsTPUpdateMachineCatalogExistingMachines) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/$RequestProvVmUpdate | Apply current provisioning scheme to existing machines
[**MachineCatalogsTPUpdateMachineCatalogMachineAccount**](MachineCatalogsTPApi.md#MachineCatalogsTPUpdateMachineCatalogMachineAccount) | **Patch** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/MachineAccounts/{machine} | Update a machine account within the machine catalog.
[**MachineCatalogsTPUpdateMachineCatalogProvisioningScheme**](MachineCatalogsTPApi.md#MachineCatalogsTPUpdateMachineCatalogProvisioningScheme) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/$UpdateProvisioningScheme | Update the machine catalog&#39;s provisioning scheme.
[**MachineCatalogsTPUpdateMachineCatalogUpgradeSchedule**](MachineCatalogsTPApi.md#MachineCatalogsTPUpdateMachineCatalogUpgradeSchedule) | **Patch** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/UpgradeSchedule | Reschedule a pending VDA upgrade schedule for a machine catalog.
[**MachineCatalogsTPValidateMachineCatalogVDAComponentsAndFeaturesSelection**](MachineCatalogsTPApi.md#MachineCatalogsTPValidateMachineCatalogVDAComponentsAndFeaturesSelection) | **Post** /techpreview/{customerid}/{siteid}/MachineCatalogs/{nameOrId}/VDAComponentsAndFeatures/$validate | Validate the VDA component selection at catalog level that can be Included/Excluded and features that need enable during an VDA upgrade using VDA Upgrade Service.



## MachineCatalogsTPAddMachineCatalogMachine

> JobResponseModel MachineCatalogsTPAddMachineCatalogMachine(ctx, nameOrId, customerid, siteid).Machine(machine).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Execute()

Add machine to the machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to add machine to.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    machine := *openapiclient.NewAddMachineToMachineCatalogDetailRequestModel() // AddMachineToMachineCatalogDetailRequestModel | Details of the machine to add to the machine catalog.
    async := true // bool | If `true`, the machine will be added into the machine catalog as a background task. The task will have JobType AddMachineCatalogMachine. When the task is complete it will redirect to GetMachineCatalog. The job's Parameters will contain properties: * _NameOrId_ - Name or ID of the machine being added. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPAddMachineCatalogMachine(context.Background(), nameOrId, customerid, siteid).Machine(machine).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPAddMachineCatalogMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPAddMachineCatalogMachine`: JobResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPAddMachineCatalogMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to add machine to.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPAddMachineCatalogMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **machine** | [**AddMachineToMachineCatalogDetailRequestModel**](AddMachineToMachineCatalogDetailRequestModel.md) | Details of the machine to add to the machine catalog. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine will be added into the machine catalog as a background task. The task will have JobType AddMachineCatalogMachine. When the task is complete it will redirect to GetMachineCatalog. The job&#39;s Parameters will contain properties: * _NameOrId_ - Name or ID of the machine being added. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**JobResponseModel**](JobResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPAddMachineCatalogMachineAccount

> ProvisioningSchemeMachineAccountResponseModel MachineCatalogsTPAddMachineCatalogMachineAccount(ctx, nameOrId, customerid, siteid).Machine(machine).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Add a machine account to the machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to add the machine account to.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    machine := *openapiclient.NewMachineAccountRequestModel("ADAccountName_example") // MachineAccountRequestModel | Machine account to add to the machine catalog.  The machine must not be a member of any other machine catalog.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPAddMachineCatalogMachineAccount(context.Background(), nameOrId, customerid, siteid).Machine(machine).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPAddMachineCatalogMachineAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPAddMachineCatalogMachineAccount`: ProvisioningSchemeMachineAccountResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPAddMachineCatalogMachineAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to add the machine account to.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPAddMachineCatalogMachineAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **machine** | [**MachineAccountRequestModel**](MachineAccountRequestModel.md) | Machine account to add to the machine catalog.  The machine must not be a member of any other machine catalog. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ProvisioningSchemeMachineAccountResponseModel**](ProvisioningSchemeMachineAccountResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPAddMachineCatalogTags

> TagResponseModelCollection MachineCatalogsTPAddMachineCatalogTags(ctx, nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Add a tag to a machine catalog.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag to add.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPAddMachineCatalogTags(context.Background(), nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPAddMachineCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPAddMachineCatalogTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPAddMachineCatalogTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**tagNameOrId** | **string** | Name or ID of the tag to add. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPAddMachineCatalogTagsRequest struct via the builder pattern


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


## MachineCatalogsTPCheckMachineCatalogExists

> MachineCatalogsTPCheckMachineCatalogExists(ctx, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of a machine catalog by name.



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
    name := "name_example" // string | Name of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPCheckMachineCatalogExists(context.Background(), name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPCheckMachineCatalogExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPCheckMachineCatalogExistsRequest struct via the builder pattern


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


## MachineCatalogsTPCheckRemotePCEnrollmentScopeExists

> MachineCatalogsTPCheckRemotePCEnrollmentScopeExists(ctx, nameOrId, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of a RemotePC enrollment scope by name.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    name := "name_example" // string | Distinguished name of an OU containing machines allowed to enroll as remote PCs.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPCheckRemotePCEnrollmentScopeExists(context.Background(), nameOrId, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPCheckRemotePCEnrollmentScopeExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**name** | **string** | Distinguished name of an OU containing machines allowed to enroll as remote PCs. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPCheckRemotePCEnrollmentScopeExistsRequest struct via the builder pattern


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


## MachineCatalogsTPCreateMachineCatalog

> MachineCatalogDetailResponseModel MachineCatalogsTPCreateMachineCatalog(ctx, customerid, siteid).Request(request).Async(async).SourceNameOrId(sourceNameOrId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Execute()

Create a machine catalog.



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
    request := *openapiclient.NewCreateMachineCatalogRequestModel("Name_example", openapiclient.ProvisioningType("Unknown"), openapiclient.SessionSupport("Unknown")) // CreateMachineCatalogRequestModel | Details about the machine catalog to create.
    async := true // bool | If `true`, the machine catalog (and associated objects) will be created as a background task. The task will have JobType CreateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job's Parameters will contain properties: * _Name_ - Name of the machine catalog being created. (optional) (default to false)
    sourceNameOrId := "sourceNameOrId_example" // string | The machine catalog name or Id to recreate from. It is only used for catalog recreation scenario. Here specify name or Id of the failed catalog to recreate. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPCreateMachineCatalog(context.Background(), customerid, siteid).Request(request).Async(async).SourceNameOrId(sourceNameOrId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPCreateMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPCreateMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPCreateMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPCreateMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateMachineCatalogRequestModel**](CreateMachineCatalogRequestModel.md) | Details about the machine catalog to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine catalog (and associated objects) will be created as a background task. The task will have JobType CreateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job&#39;s Parameters will contain properties: * _Name_ - Name of the machine catalog being created. | [default to false]
 **sourceNameOrId** | **string** | The machine catalog name or Id to recreate from. It is only used for catalog recreation scenario. Here specify name or Id of the failed catalog to recreate. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**MachineCatalogDetailResponseModel**](MachineCatalogDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPCreateMachineCatalogUpgradeSchedule

> MachineCatalogsTPCreateMachineCatalogUpgradeSchedule(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Setup a new VDA upgrade schedule for a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewCreateUpgradeScheduleRequestModel(int32(123)) // CreateUpgradeScheduleRequestModel | Details about the VDA upgrade schedule to create.
    async := true // bool | If `true`, the VDA upgrade schedule will be created as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPCreateMachineCatalogUpgradeSchedule(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPCreateMachineCatalogUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPCreateMachineCatalogUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**CreateUpgradeScheduleRequestModel**](CreateUpgradeScheduleRequestModel.md) | Details about the VDA upgrade schedule to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedule will be created as a background task. | [default to false]
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


## MachineCatalogsTPDeleteMachineCatalog

> MachineCatalogsTPDeleteMachineCatalog(ctx, nameOrId, customerid, siteid).DeleteVm(deleteVm).DeleteAccount(deleteAccount).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Execute()

Delete a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to update.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    deleteVm := true // bool | Specifies whether to delete virtual machines that were provisioned in the machine catalog. Only relevant when ProvisioningType is MCS; otherwise, ignored.  Optional; default is `true`. (optional) (default to true)
    deleteAccount := "deleteAccount_example" // string | Specifies what to do with AD machine accounts that are associated with the machine catalog.  Optional; default is . (optional) (default to "None")
    async := true // bool | If `true`, the machine catalog (and associated objects) will be deleted as a background task. The task will have JobType DeleteMachineCatalog. When the task is complete it will redirect to GetMachineCatalogs. The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog being deleted, * _Name_ - Name of the machine catalog being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPDeleteMachineCatalog(context.Background(), nameOrId, customerid, siteid).DeleteVm(deleteVm).DeleteAccount(deleteAccount).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPDeleteMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to update.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPDeleteMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deleteVm** | **bool** | Specifies whether to delete virtual machines that were provisioned in the machine catalog. Only relevant when ProvisioningType is MCS; otherwise, ignored.  Optional; default is &#x60;true&#x60;. | [default to true]
 **deleteAccount** | **string** | Specifies what to do with AD machine accounts that are associated with the machine catalog.  Optional; default is . | [default to &quot;None&quot;]
 **async** | **bool** | If &#x60;true&#x60;, the machine catalog (and associated objects) will be deleted as a background task. The task will have JobType DeleteMachineCatalog. When the task is complete it will redirect to GetMachineCatalogs. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog being deleted, * _Name_ - Name of the machine catalog being deleted. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
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


## MachineCatalogsTPDoMachineCatalogSearch

> CatalogSearchResponseModelCollection MachineCatalogsTPDoMachineCatalogSearch(ctx, customerid, siteid).SearchRequest(searchRequest).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Perform an advanced search for machinecatalogs.



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
    searchRequest := *openapiclient.NewCatalogSearchRequestModel() // CatalogSearchRequestModel | Specifies the advanced search parameters.
    async := true // bool | If `true`, the catalogs (and associated objects) will be added as a background task. The task will have JobType SearchMachineCatalogs. (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPDoMachineCatalogSearch(context.Background(), customerid, siteid).SearchRequest(searchRequest).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPDoMachineCatalogSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPDoMachineCatalogSearch`: CatalogSearchResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPDoMachineCatalogSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPDoMachineCatalogSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchRequest** | [**CatalogSearchRequestModel**](CatalogSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **async** | **bool** | If &#x60;true&#x60;, the catalogs (and associated objects) will be added as a background task. The task will have JobType SearchMachineCatalogs. | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**CatalogSearchResponseModelCollection**](CatalogSearchResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPGetMachineCatalog

> MachineCatalogDetailResponseModel MachineCatalogsTPGetMachineCatalog(ctx, nameOrId, customerid, siteid).Async(async).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details about a single machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    fields := "fields_example" // string | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalog(context.Background(), nameOrId, customerid, siteid).Async(async).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **fields** | **string** | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. | 
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


## MachineCatalogsTPGetMachineCatalogDeliveryGroupAssociations

> DeliveryGroupAssociationResponseModelCollection MachineCatalogsTPGetMachineCatalogDeliveryGroupAssociations(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get delivery group associations of a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of delivery group associations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogDeliveryGroupAssociations(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogDeliveryGroupAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogDeliveryGroupAssociations`: DeliveryGroupAssociationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogDeliveryGroupAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogDeliveryGroupAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of delivery group associations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**DeliveryGroupAssociationResponseModelCollection**](DeliveryGroupAssociationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPGetMachineCatalogLastMasterImage

> VMImageResponseModel MachineCatalogsTPGetMachineCatalogLastMasterImage(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the last master VM images used by the machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogLastMasterImage(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogLastMasterImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogLastMasterImage`: VMImageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogLastMasterImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogLastMasterImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**VMImageResponseModel**](VMImageResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPGetMachineCatalogMachineAccounts

> ProvisioningSchemeMachineAccountResponseModelCollection MachineCatalogsTPGetMachineCatalogMachineAccounts(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).AccountState(accountState).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all machine accounts associated with a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to get machine accounts for.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of machine accounts returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the machine accounts will be queried as a background task. The task will have JobType GetMachineCatalogMachineAccounts. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    accountState := "accountState_example" // string | The state of accounts for query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogMachineAccounts(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).AccountState(accountState).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogMachineAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogMachineAccounts`: ProvisioningSchemeMachineAccountResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogMachineAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to get machine accounts for.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogMachineAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of machine accounts returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine accounts will be queried as a background task. The task will have JobType GetMachineCatalogMachineAccounts. When the task is complete it will redirect to GetJobResults. | [default to false]
 **accountState** | **string** | The state of accounts for query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ProvisioningSchemeMachineAccountResponseModelCollection**](ProvisioningSchemeMachineAccountResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPGetMachineCatalogMachines

> MachineResponseModelCollection MachineCatalogsTPGetMachineCatalogMachines(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

GET the machines of a catalog

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    fields := "fields_example" // string | Optional parameters, only the specified properties in the fields are required. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogMachines(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogMachines`: MachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **fields** | **string** | Optional parameters, only the specified properties in the fields are required. | 
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


## MachineCatalogsTPGetMachineCatalogMasterImageHistory

> VMImageResponseModelCollection MachineCatalogsTPGetMachineCatalogMasterImageHistory(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the history of master VM images used by the machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogMasterImageHistory(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogMasterImageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogMasterImageHistory`: VMImageResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogMasterImageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogMasterImageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**VMImageResponseModelCollection**](VMImageResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPGetMachineCatalogStartMenuApplications

> StartMenuApplicationResponseModelCollection MachineCatalogsTPGetMachineCatalogStartMenuApplications(ctx, nameOrId, customerid, siteid).Machines(machines).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the start menu applications from a machine in the machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    machines := []string{"Inner_example"} // []string | The machines to get start menu applications (optional)
    async := true // bool | If `true`, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to \"JobsControllerTP.GetJobResults(string)\". The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog from which start menu applications are being obtained, * _Name_ - Name of the machine catalog from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogStartMenuApplications(context.Background(), nameOrId, customerid, siteid).Machines(machines).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogStartMenuApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogStartMenuApplications`: StartMenuApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogStartMenuApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogStartMenuApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **machines** | **[]string** | The machines to get start menu applications | 
 **async** | **bool** | If &#x60;true&#x60;, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to \&quot;JobsControllerTP.GetJobResults(string)\&quot;. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog from which start menu applications are being obtained, * _Name_ - Name of the machine catalog from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. | [default to false]
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


## MachineCatalogsTPGetMachineCatalogTags

> TagResponseModelCollection MachineCatalogsTPGetMachineCatalogTags(ctx, nameOrId, customerid, siteid).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get tags associated with a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    fields := "fields_example" // string | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogTags(context.Background(), nameOrId, customerid, siteid).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **string** | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. | 
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


## MachineCatalogsTPGetMachineCatalogTestReport

> TestReportResponseModel MachineCatalogsTPGetMachineCatalogTestReport(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the most recent test report of a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogTestReport(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogTestReportRequest struct via the builder pattern


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


## MachineCatalogsTPGetMachineCatalogVDAComponentsAndFeatures

> MachineCatalogVusComponentResponseModel MachineCatalogsTPGetMachineCatalogVDAComponentsAndFeatures(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the components and features of VDAs associated with a machine catalog.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogVDAComponentsAndFeatures(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogVDAComponentsAndFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogVDAComponentsAndFeatures`: MachineCatalogVusComponentResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogVDAComponentsAndFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogVDAComponentsAndFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**MachineCatalogVusComponentResponseModel**](MachineCatalogVusComponentResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPGetMachineCatalogs

> MachineCatalogResponseModelCollection MachineCatalogsTPGetMachineCatalogs(ctx, customerid, siteid).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all machine catalogs.



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
    adminFolder := "adminFolder_example" // string | Admin folder path or Id. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogs(context.Background(), customerid, siteid).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogs`: MachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adminFolder** | **string** | Admin folder path or Id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server | 
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


## MachineCatalogsTPGetMachineCatalogsAdministrators

> AdministratorResponseModelCollection MachineCatalogsTPGetMachineCatalogsAdministrators(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get administrators who can administer a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogsAdministrators(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogsAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPGetMachineCatalogsAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPGetMachineCatalogsAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPGetMachineCatalogsAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## MachineCatalogsTPRebootMachineCatalogMachines

> MachineCatalogsTPRebootMachineCatalogMachines(ctx, nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Reboot all machines in a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewRebootMachinesRequestModel() // RebootMachinesRequestModel | Reboot options.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPRebootMachineCatalogMachines(context.Background(), nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPRebootMachineCatalogMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPRebootMachineCatalogMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**RebootMachinesRequestModel**](RebootMachinesRequestModel.md) | Reboot options. | 
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


## MachineCatalogsTPRemoveMachineCatalogMachine

> MachineCatalogsTPRemoveMachineCatalogMachine(ctx, nameOrId, machine, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Remove a machine from the machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to remove the machine from.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    machine := "machine_example" // string | Machine to remove from the catalog.  May be specified by name, ID, or SID.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPRemoveMachineCatalogMachine(context.Background(), nameOrId, machine, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPRemoveMachineCatalogMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to remove the machine from.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**machine** | **string** | Machine to remove from the catalog.  May be specified by name, ID, or SID. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPRemoveMachineCatalogMachineRequest struct via the builder pattern


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


## MachineCatalogsTPRemoveMachineCatalogMachineAccount

> MachineCatalogsTPRemoveMachineCatalogMachineAccount(ctx, nameOrId, machine, customerid, siteid).DeleteAccount(deleteAccount).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Remove a machine account from the machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to remove the machine account from.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    machine := "machine_example" // string | Machine account to remove.  May be specified by SamName, UPN, or SID.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    deleteAccount := "deleteAccount_example" // string | Specifies what to do with the machine account in AD.  Optional; default is None. (optional) (default to "None")
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPRemoveMachineCatalogMachineAccount(context.Background(), nameOrId, machine, customerid, siteid).DeleteAccount(deleteAccount).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPRemoveMachineCatalogMachineAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to remove the machine account from.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**machine** | **string** | Machine account to remove.  May be specified by SamName, UPN, or SID. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPRemoveMachineCatalogMachineAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **deleteAccount** | **string** | Specifies what to do with the machine account in AD.  Optional; default is None. | [default to &quot;None&quot;]
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


## MachineCatalogsTPRemoveMachineCatalogTags

> MachineCatalogsTPRemoveMachineCatalogTags(ctx, nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Remove a tag from a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPRemoveMachineCatalogTags(context.Background(), nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPRemoveMachineCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**tagNameOrId** | **string** | Name or ID of the tag. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPRemoveMachineCatalogTagsRequest struct via the builder pattern


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


## MachineCatalogsTPRemoveMachineCatalogUpgradeSchedule

> MachineCatalogsTPRemoveMachineCatalogUpgradeSchedule(ctx, nameOrId, customerid, siteid).Force(force).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Remove a pending VDA upgrade schedule for a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    force := true // bool | If `true`, VDA upgrade schedules existing with all machines within the catalog will be removed except those upgrade status in progressing. If `false`, only the perform removing when catalog level schedule not started. (optional) (default to false)
    async := true // bool | If `true`, the VDA upgrade schedule will be removed as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPRemoveMachineCatalogUpgradeSchedule(context.Background(), nameOrId, customerid, siteid).Force(force).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPRemoveMachineCatalogUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPRemoveMachineCatalogUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | If &#x60;true&#x60;, VDA upgrade schedules existing with all machines within the catalog will be removed except those upgrade status in progressing. If &#x60;false&#x60;, only the perform removing when catalog level schedule not started. | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedule will be removed as a background task. | [default to false]
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


## MachineCatalogsTPRollbackMachineCatalogProvisioningScheme

> JobResponseModel MachineCatalogsTPRollbackMachineCatalogProvisioningScheme(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Rollback the machine catalog's provisioning scheme.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewRebootMachinesRequestModel() // RebootMachinesRequestModel | Details about the image update to be performed.
    async := true // bool | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the `async` parameter is `true`, the only difference is that instead of returning the initial job status in the payload, the response is instead `202` with the `Location` header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPRollbackMachineCatalogProvisioningScheme(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPRollbackMachineCatalogProvisioningScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPRollbackMachineCatalogProvisioningScheme`: JobResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPRollbackMachineCatalogProvisioningScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPRollbackMachineCatalogProvisioningSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**RebootMachinesRequestModel**](RebootMachinesRequestModel.md) | Details about the image update to be performed. | 
 **async** | **bool** | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the &#x60;async&#x60; parameter is &#x60;true&#x60;, the only difference is that instead of returning the initial job status in the payload, the response is instead &#x60;202&#x60; with the &#x60;Location&#x60; header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**JobResponseModel**](JobResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPSetMachineCatalogTags

> MachineCatalogsTPSetMachineCatalogTags(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Set tags associated with a machine catalog.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewTagsRequestModel() // TagsRequestModel | Details of the tags to set on the machine catalog.
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPSetMachineCatalogTags(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPSetMachineCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPSetMachineCatalogTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**TagsRequestModel**](TagsRequestModel.md) | Details of the tags to set on the machine catalog. | 
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


## MachineCatalogsTPTestMachineCatalog

> MachineCatalogTestResponseModel MachineCatalogsTPTestMachineCatalog(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Run tests on a machine catalog and create a test report.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to test.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestMachineCatalog. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog being tested, * _Name_ - Name of the machine catalog being tested. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPTestMachineCatalog(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPTestMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPTestMachineCatalog`: MachineCatalogTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPTestMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to test.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPTestMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestMachineCatalog. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog being tested, * _Name_ - Name of the machine catalog being tested. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**MachineCatalogTestResponseModel**](MachineCatalogTestResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPTestMachineCatalogExists

> MachineCatalogsTPTestMachineCatalogExists(ctx, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Test the existence of a machine catalog by name.



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
    request := *openapiclient.NewCatalogNameCheckRequestModel() // CatalogNameCheckRequestModel | Name of the machine catalog.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPTestMachineCatalogExists(context.Background(), customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPTestMachineCatalogExists``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMachineCatalogsTPTestMachineCatalogExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CatalogNameCheckRequestModel**](CatalogNameCheckRequestModel.md) | Name of the machine catalog. | 
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


## MachineCatalogsTPTestRemotePCEnrollmentScopeExists

> MachineCatalogsTPTestRemotePCEnrollmentScopeExists(ctx, nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of a RemotePC enrollment scope by name in the request body. Use HttpPost and take name in body as this name may be very long, and cause url too long exception.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewRemotePCEnrollmentScopeCheckModel("Name_example") // RemotePCEnrollmentScopeCheckModel | Request model that contains distinguished name of an OU containing machines allowed to enroll as remote PCs.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPTestRemotePCEnrollmentScopeExists(context.Background(), nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPTestRemotePCEnrollmentScopeExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPTestRemotePCEnrollmentScopeExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**RemotePCEnrollmentScopeCheckModel**](RemotePCEnrollmentScopeCheckModel.md) | Request model that contains distinguished name of an OU containing machines allowed to enroll as remote PCs. | 
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


## MachineCatalogsTPUpdateMachineCatalog

> MachineCatalogDetailResponseModel MachineCatalogsTPUpdateMachineCatalog(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()

Update a machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to update.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewUpdateMachineCatalogRequestModel() // UpdateMachineCatalogRequestModel | Properties of the machine catalog to update.
    async := true // bool | If `true`, the machine catalog (and associated objects) will be updated as a background task. The task will have JobType UpdateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog being updated, * _Name_ - Name of the machine catalog being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalog(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPUpdateMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to update.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPUpdateMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateMachineCatalogRequestModel**](UpdateMachineCatalogRequestModel.md) | Properties of the machine catalog to update. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine catalog (and associated objects) will be updated as a background task. The task will have JobType UpdateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog being updated, * _Name_ - Name of the machine catalog being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**MachineCatalogDetailResponseModel**](MachineCatalogDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPUpdateMachineCatalogExistingMachines

> MachineCatalogsTPUpdateMachineCatalogExistingMachines(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Apply current provisioning scheme to existing machines



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogExistingMachines(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogExistingMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPUpdateMachineCatalogExistingMachinesRequest struct via the builder pattern


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


## MachineCatalogsTPUpdateMachineCatalogMachineAccount

> ProvisioningSchemeMachineAccountResponseModel MachineCatalogsTPUpdateMachineCatalogMachineAccount(ctx, nameOrId, machine, customerid, siteid).Details(details).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Update a machine account within the machine catalog.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog containing the machine account.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    machine := "machine_example" // string | Machine account to update.  May be specified by SamName, UPN, or SID.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    details := *openapiclient.NewUpdateMachineAccountRequestModel() // UpdateMachineAccountRequestModel | New details for the machine account record.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogMachineAccount(context.Background(), nameOrId, machine, customerid, siteid).Details(details).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogMachineAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPUpdateMachineCatalogMachineAccount`: ProvisioningSchemeMachineAccountResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogMachineAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog containing the machine account.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**machine** | **string** | Machine account to update.  May be specified by SamName, UPN, or SID. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPUpdateMachineCatalogMachineAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **details** | [**UpdateMachineAccountRequestModel**](UpdateMachineAccountRequestModel.md) | New details for the machine account record. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ProvisioningSchemeMachineAccountResponseModel**](ProvisioningSchemeMachineAccountResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPUpdateMachineCatalogProvisioningScheme

> JobResponseModel MachineCatalogsTPUpdateMachineCatalogProvisioningScheme(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update the machine catalog's provisioning scheme.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewUpdateMachineCatalogProvisioningSchemeRequestModel() // UpdateMachineCatalogProvisioningSchemeRequestModel | Details about the image update to be performed.
    async := true // bool | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the `async` parameter is `true`, the only difference is that instead of returning the initial job status in the payload, the response is instead `202` with the `Location` header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated, * _MasterImageXDPath_ - Path to the new image specified for the provisioning scheme, in XenApp & XenDesktop path format. * _MasterImageFullPath_ - Path to the new image specified for the provisioning scheme, relative to the root of the API. * _MasterImageRelativePath_ - Path to the new image specified for the provisioning scheme, relative to the resource pool. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogProvisioningScheme(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogProvisioningScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPUpdateMachineCatalogProvisioningScheme`: JobResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogProvisioningScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPUpdateMachineCatalogProvisioningSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateMachineCatalogProvisioningSchemeRequestModel**](UpdateMachineCatalogProvisioningSchemeRequestModel.md) | Details about the image update to be performed. | 
 **async** | **bool** | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the &#x60;async&#x60; parameter is &#x60;true&#x60;, the only difference is that instead of returning the initial job status in the payload, the response is instead &#x60;202&#x60; with the &#x60;Location&#x60; header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated, * _MasterImageXDPath_ - Path to the new image specified for the provisioning scheme, in XenApp &amp; XenDesktop path format. * _MasterImageFullPath_ - Path to the new image specified for the provisioning scheme, relative to the root of the API. * _MasterImageRelativePath_ - Path to the new image specified for the provisioning scheme, relative to the resource pool. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**JobResponseModel**](JobResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTPUpdateMachineCatalogUpgradeSchedule

> MachineCatalogsTPUpdateMachineCatalogUpgradeSchedule(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Reschedule a pending VDA upgrade schedule for a machine catalog.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewUpdateUpgradeScheduleRequestModel(time.Now(), int32(123)) // UpdateUpgradeScheduleRequestModel | Details about the updated VDA upgrade schedule.
    async := true // bool | If `true`, the VDA upgrade schedule will be updated as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogUpgradeSchedule(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPUpdateMachineCatalogUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPUpdateMachineCatalogUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateUpgradeScheduleRequestModel**](UpdateUpgradeScheduleRequestModel.md) | Details about the updated VDA upgrade schedule. | 
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedule will be updated as a background task. | [default to false]
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


## MachineCatalogsTPValidateMachineCatalogVDAComponentsAndFeaturesSelection

> VDAComponentsSelectionValidationResponseModel MachineCatalogsTPValidateMachineCatalogVDAComponentsAndFeaturesSelection(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Validate the VDA component selection at catalog level that can be Included/Excluded and features that need enable during an VDA upgrade using VDA Upgrade Service.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewVDAComponentsSelectionValidationRequestModel() // VDAComponentsSelectionValidationRequestModel | More details attached to this validation.
    async := true // bool | If async execute. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsTPApi.MachineCatalogsTPValidateMachineCatalogVDAComponentsAndFeaturesSelection(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsTPApi.MachineCatalogsTPValidateMachineCatalogVDAComponentsAndFeaturesSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTPValidateMachineCatalogVDAComponentsAndFeaturesSelection`: VDAComponentsSelectionValidationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsTPApi.MachineCatalogsTPValidateMachineCatalogVDAComponentsAndFeaturesSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTPValidateMachineCatalogVDAComponentsAndFeaturesSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**VDAComponentsSelectionValidationRequestModel**](VDAComponentsSelectionValidationRequestModel.md) | More details attached to this validation. | 
 **async** | **bool** | If async execute. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**VDAComponentsSelectionValidationResponseModel**](VDAComponentsSelectionValidationResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

