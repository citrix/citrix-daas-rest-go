# \MachineCatalogsAPIs 

All URIs are relative to *https://api-us.cloud.com/cvad/manage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MachineCatalogsAddMachineCatalogMachine**](MachineCatalogsAPIs .md#MachineCatalogsAddMachineCatalogMachine) | **Post** /MachineCatalogs/{nameOrId}/Machines | Add machine to the machine catalog.
[**MachineCatalogsAddMachineCatalogMachineAccount**](MachineCatalogsAPIs .md#MachineCatalogsAddMachineCatalogMachineAccount) | **Post** /MachineCatalogs/{nameOrId}/MachineAccounts | Add a machine account to the machine catalog.
[**MachineCatalogsAddMachineCatalogTags**](MachineCatalogsAPIs .md#MachineCatalogsAddMachineCatalogTags) | **Post** /MachineCatalogs/{nameOrId}/Tags/{tagNameOrId} | Add a tag to a machine catalog.
[**MachineCatalogsCheckMachineCatalogExists**](MachineCatalogsAPIs .md#MachineCatalogsCheckMachineCatalogExists) | **Head** /MachineCatalogs/{name} | Check for the existence of a machine catalog by name.
[**MachineCatalogsCheckRemotePCEnrollmentScopeExists**](MachineCatalogsAPIs .md#MachineCatalogsCheckRemotePCEnrollmentScopeExists) | **Head** /MachineCatalogs/{nameOrId}/RemotePCEnrollmentScopes/{name} | Check for the existence of a RemotePC enrollment scope by name.
[**MachineCatalogsCreateMachineCatalog**](MachineCatalogsAPIs .md#MachineCatalogsCreateMachineCatalog) | **Post** /MachineCatalogs | Create a machine catalog.
[**MachineCatalogsCreateMachineCatalogUpgradeSchedule**](MachineCatalogsAPIs .md#MachineCatalogsCreateMachineCatalogUpgradeSchedule) | **Post** /MachineCatalogs/{nameOrId}/UpgradeSchedule | Setup a new VDA upgrade schedule for a machine catalog.
[**MachineCatalogsDeleteMachineCatalog**](MachineCatalogsAPIs .md#MachineCatalogsDeleteMachineCatalog) | **Delete** /MachineCatalogs/{nameOrId} | Delete a machine catalog.
[**MachineCatalogsDoMachineCatalogSearch**](MachineCatalogsAPIs .md#MachineCatalogsDoMachineCatalogSearch) | **Post** /Machinecatalogs/$search | Perform an advanced search for machinecatalogs.
[**MachineCatalogsGetMachineCatalog**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalog) | **Get** /MachineCatalogs/{nameOrId} | Get details about a single machine catalog.
[**MachineCatalogsGetMachineCatalogDeliveryGroupAssociations**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogDeliveryGroupAssociations) | **Get** /MachineCatalogs/{nameOrId}/DeliveryGroupAssociations | Get delivery group associations of a machine catalog.
[**MachineCatalogsGetMachineCatalogLastMasterImage**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogLastMasterImage) | **Get** /MachineCatalogs/{nameOrId}/LastMasterImage | Get the last master VM images used by the machine catalog.
[**MachineCatalogsGetMachineCatalogMachineAccounts**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogMachineAccounts) | **Get** /MachineCatalogs/{nameOrId}/MachineAccounts | Get all machine accounts associated with a machine catalog.
[**MachineCatalogsGetMachineCatalogMachines**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogMachines) | **Get** /MachineCatalogs/{nameOrId}/Machines | GET the machines of a catalog
[**MachineCatalogsGetMachineCatalogMasterImageHistory**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogMasterImageHistory) | **Get** /MachineCatalogs/{nameOrId}/MasterImageHistory | Get the history of master VM images used by the machine catalog.
[**MachineCatalogsGetMachineCatalogStartMenuApplications**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogStartMenuApplications) | **Get** /MachineCatalogs/{nameOrId}/StartMenuApplications | Get the start menu applications from a machine in the machine catalog.
[**MachineCatalogsGetMachineCatalogTags**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogTags) | **Get** /MachineCatalogs/{nameOrId}/Tags | Get tags associated with a machine catalog.
[**MachineCatalogsGetMachineCatalogTestReport**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogTestReport) | **Get** /MachineCatalogs/{nameOrId}/TestReport | Get the most recent test report of a machine catalog.
[**MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures) | **Get** /MachineCatalogs/{nameOrId}/VDAComponentsAndFeatures | Get the components and features of VDAs associated with a machine catalog.
[**MachineCatalogsGetMachineCatalogs**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogs) | **Get** /MachineCatalogs | Get all machine catalogs.
[**MachineCatalogsGetMachineCatalogsAdministrators**](MachineCatalogsAPIs .md#MachineCatalogsGetMachineCatalogsAdministrators) | **Get** /MachineCatalogs/{nameOrId}/Administrators | Get administrators who can administer a machine catalog.
[**MachineCatalogsRebootMachineCatalogMachines**](MachineCatalogsAPIs .md#MachineCatalogsRebootMachineCatalogMachines) | **Post** /MachineCatalogs/{nameOrId}/$RebootMachines | Reboot all machines in a machine catalog.
[**MachineCatalogsRemoveMachineCatalogMachine**](MachineCatalogsAPIs .md#MachineCatalogsRemoveMachineCatalogMachine) | **Delete** /MachineCatalogs/{nameOrId}/Machines/{machine} | Remove a machine from the machine catalog.
[**MachineCatalogsRemoveMachineCatalogMachineAccount**](MachineCatalogsAPIs .md#MachineCatalogsRemoveMachineCatalogMachineAccount) | **Delete** /MachineCatalogs/{nameOrId}/MachineAccounts/{machine} | Remove a machine account from the machine catalog.
[**MachineCatalogsRemoveMachineCatalogTags**](MachineCatalogsAPIs .md#MachineCatalogsRemoveMachineCatalogTags) | **Delete** /MachineCatalogs/{nameOrId}/Tags/{tagNameOrId} | Remove a tag from a machine catalog.
[**MachineCatalogsRemoveMachineCatalogUpgradeSchedule**](MachineCatalogsAPIs .md#MachineCatalogsRemoveMachineCatalogUpgradeSchedule) | **Delete** /MachineCatalogs/{nameOrId}/UpgradeSchedule | Remove a pending VDA upgrade schedule for a machine catalog.
[**MachineCatalogsRollbackMachineCatalogProvisioningScheme**](MachineCatalogsAPIs .md#MachineCatalogsRollbackMachineCatalogProvisioningScheme) | **Post** /MachineCatalogs/{nameOrId}/$RollbackProvisioningScheme | Rollback the machine catalog&#39;s provisioning scheme.
[**MachineCatalogsSetMachineCatalogTags**](MachineCatalogsAPIs .md#MachineCatalogsSetMachineCatalogTags) | **Put** /MachineCatalogs/{nameOrId}/Tags | Set tags associated with a machine catalog.
[**MachineCatalogsTestMachineCatalog**](MachineCatalogsAPIs .md#MachineCatalogsTestMachineCatalog) | **Post** /MachineCatalogs/{nameOrId}/$test | Run tests on a machine catalog and create a test report.
[**MachineCatalogsTestMachineCatalogExists**](MachineCatalogsAPIs .md#MachineCatalogsTestMachineCatalogExists) | **Post** /MachineCatalogs/$checkCatalogName | Test the existence of a machine catalog by name.
[**MachineCatalogsTestRemotePCEnrollmentScopeExists**](MachineCatalogsAPIs .md#MachineCatalogsTestRemotePCEnrollmentScopeExists) | **Post** /MachineCatalogs/{nameOrId}/RemotePCEnrollmentScopes/$checkScopeName | Check for the existence of a RemotePC enrollment scope by name in the request body. Use HttpPost and take name in body as this name may be very long, and cause url too long exception.
[**MachineCatalogsUpdateMachineCatalog**](MachineCatalogsAPIs .md#MachineCatalogsUpdateMachineCatalog) | **Patch** /MachineCatalogs/{nameOrId} | Update a machine catalog.
[**MachineCatalogsUpdateMachineCatalogMachineAccount**](MachineCatalogsAPIs .md#MachineCatalogsUpdateMachineCatalogMachineAccount) | **Patch** /MachineCatalogs/{nameOrId}/MachineAccounts/{machine} | Update a machine account within the machine catalog.
[**MachineCatalogsUpdateMachineCatalogProvisioningScheme**](MachineCatalogsAPIs .md#MachineCatalogsUpdateMachineCatalogProvisioningScheme) | **Post** /MachineCatalogs/{nameOrId}/$UpdateProvisioningScheme | Update the machine catalog&#39;s provisioning scheme.
[**MachineCatalogsUpdateMachineCatalogUpgradeSchedule**](MachineCatalogsAPIs .md#MachineCatalogsUpdateMachineCatalogUpgradeSchedule) | **Patch** /MachineCatalogs/{nameOrId}/UpgradeSchedule | Reschedule a pending VDA upgrade schedule for a machine catalog.
[**MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection**](MachineCatalogsAPIs .md#MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection) | **Post** /MachineCatalogs/{nameOrId}/VDAComponentsAndFeatures/$validate | Validate the VDA component selection at catalog level that can be Included/Excluded and features that need enable during an VDA upgrade using VDA Upgrade Service.



## MachineCatalogsAddMachineCatalogMachine

> JobResponseModel MachineCatalogsAddMachineCatalogMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AddMachineToMachineCatalogDetailRequestModel(addMachineToMachineCatalogDetailRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    addMachineToMachineCatalogDetailRequestModel := *openapiclient.NewAddMachineToMachineCatalogDetailRequestModel() // AddMachineToMachineCatalogDetailRequestModel | Details of the machine to add to the machine catalog.
    async := true // bool | If `true`, the machine will be added into the machine catalog as a background task. The task will have JobType AddMachineCatalogMachine. When the task is complete it will redirect to GetMachineCatalog. The job's Parameters will contain properties: * _NameOrId_ - Name or ID of the machine being added. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsAddMachineCatalogMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AddMachineToMachineCatalogDetailRequestModel(addMachineToMachineCatalogDetailRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsAddMachineCatalogMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsAddMachineCatalogMachine`: JobResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsAddMachineCatalogMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to add machine to.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsAddMachineCatalogMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **addMachineToMachineCatalogDetailRequestModel** | [**AddMachineToMachineCatalogDetailRequestModel**](AddMachineToMachineCatalogDetailRequestModel.md) | Details of the machine to add to the machine catalog. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine will be added into the machine catalog as a background task. The task will have JobType AddMachineCatalogMachine. When the task is complete it will redirect to GetMachineCatalog. The job&#39;s Parameters will contain properties: * _NameOrId_ - Name or ID of the machine being added. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsAddMachineCatalogMachineAccount

> ProvisioningSchemeMachineAccountResponseModel MachineCatalogsAddMachineCatalogMachineAccount(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).MachineAccountRequestModel(machineAccountRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    machineAccountRequestModel := *openapiclient.NewMachineAccountRequestModel("For example, MyDomain\MyComputer") // MachineAccountRequestModel | Machine account to add to the machine catalog.  The machine must not be a member of any other machine catalog.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsAddMachineCatalogMachineAccount(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).MachineAccountRequestModel(machineAccountRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsAddMachineCatalogMachineAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsAddMachineCatalogMachineAccount`: ProvisioningSchemeMachineAccountResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsAddMachineCatalogMachineAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to add the machine account to.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsAddMachineCatalogMachineAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **machineAccountRequestModel** | [**MachineAccountRequestModel**](MachineAccountRequestModel.md) | Machine account to add to the machine catalog.  The machine must not be a member of any other machine catalog. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsAddMachineCatalogTags

> TagResponseModelCollection MachineCatalogsAddMachineCatalogTags(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsAddMachineCatalogTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsAddMachineCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsAddMachineCatalogTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsAddMachineCatalogTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**tagNameOrId** | **string** | Name or ID of the tag to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsAddMachineCatalogTagsRequest struct via the builder pattern


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


## MachineCatalogsCheckMachineCatalogExists

> MachineCatalogsCheckMachineCatalogExists(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsCheckMachineCatalogExists(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsCheckMachineCatalogExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsCheckMachineCatalogExistsRequest struct via the builder pattern


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


## MachineCatalogsCheckRemotePCEnrollmentScopeExists

> MachineCatalogsCheckRemotePCEnrollmentScopeExists(ctx, nameOrId, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsCheckRemotePCEnrollmentScopeExists(context.Background(), nameOrId, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsCheckRemotePCEnrollmentScopeExists``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsCheckRemotePCEnrollmentScopeExistsRequest struct via the builder pattern


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


## MachineCatalogsCreateMachineCatalog

> MachineCatalogDetailResponseModel MachineCatalogsCreateMachineCatalog(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateMachineCatalogRequestModel(createMachineCatalogRequestModel).Async(async).SourceNameOrId(sourceNameOrId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    createMachineCatalogRequestModel := *openapiclient.NewCreateMachineCatalogRequestModel("Name of the machine catalog", openapiclient.ProvisioningType("Unknown"), openapiclient.SessionSupport("Unknown")) // CreateMachineCatalogRequestModel | Details about the machine catalog to create.
    async := true // bool | If `true`, the machine catalog (and associated objects) will be created as a background task. The task will have JobType CreateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job's Parameters will contain properties: * _Name_ - Name of the machine catalog being created. (optional) (default to false)
    sourceNameOrId := "sourceNameOrId_example" // string | The machine catalog name or Id to recreate from. It is only used for catalog recreation scenario. Here specify name or Id of the failed catalog to recreate. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsCreateMachineCatalog(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateMachineCatalogRequestModel(createMachineCatalogRequestModel).Async(async).SourceNameOrId(sourceNameOrId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsCreateMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsCreateMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsCreateMachineCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsCreateMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **createMachineCatalogRequestModel** | [**CreateMachineCatalogRequestModel**](CreateMachineCatalogRequestModel.md) | Details about the machine catalog to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine catalog (and associated objects) will be created as a background task. The task will have JobType CreateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job&#39;s Parameters will contain properties: * _Name_ - Name of the machine catalog being created. | [default to false]
 **sourceNameOrId** | **string** | The machine catalog name or Id to recreate from. It is only used for catalog recreation scenario. Here specify name or Id of the failed catalog to recreate. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsCreateMachineCatalogUpgradeSchedule

> MachineCatalogsCreateMachineCatalogUpgradeSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateUpgradeScheduleRequestModel(createUpgradeScheduleRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    createUpgradeScheduleRequestModel := *openapiclient.NewCreateUpgradeScheduleRequestModel(int32(123)) // CreateUpgradeScheduleRequestModel | Details about the VDA upgrade schedule to create.
    async := true // bool | If `true`, the VDA upgrade schedule will be created as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsCreateMachineCatalogUpgradeSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateUpgradeScheduleRequestModel(createUpgradeScheduleRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsCreateMachineCatalogUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsCreateMachineCatalogUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **createUpgradeScheduleRequestModel** | [**CreateUpgradeScheduleRequestModel**](CreateUpgradeScheduleRequestModel.md) | Details about the VDA upgrade schedule to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the VDA upgrade schedule will be created as a background task. | [default to false]
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


## MachineCatalogsDeleteMachineCatalog

> MachineCatalogsDeleteMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeleteVm(deleteVm).DeleteAccount(deleteAccount).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    deleteVm := true // bool | Specifies whether to delete virtual machines that were provisioned in the machine catalog. Only relevant when ProvisioningType is MCS; otherwise, ignored.  Optional; default is `true`. (optional) (default to true)
    deleteAccount := openapiclient.MachineCatalogs_DeleteMachineCatalog_deleteAccount_parameter{MachineAccountDeleteOption: openapiclient.MachineAccountDeleteOption("Unknown")} // MachineCatalogsDeleteMachineCatalogDeleteAccountParameter | Specifies what to do with AD machine accounts that are associated with the machine catalog.  Optional; default is . (optional) (default to None)
    async := true // bool | If `true`, the machine catalog (and associated objects) will be deleted as a background task. The task will have JobType DeleteMachineCatalog. When the task is complete it will redirect to GetMachineCatalogs. The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog being deleted, * _Name_ - Name of the machine catalog being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsDeleteMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeleteVm(deleteVm).DeleteAccount(deleteAccount).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsDeleteMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to update.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsDeleteMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **deleteVm** | **bool** | Specifies whether to delete virtual machines that were provisioned in the machine catalog. Only relevant when ProvisioningType is MCS; otherwise, ignored.  Optional; default is &#x60;true&#x60;. | [default to true]
 **deleteAccount** | [**MachineCatalogsDeleteMachineCatalogDeleteAccountParameter**](MachineCatalogsDeleteMachineCatalogDeleteAccountParameter.md) | Specifies what to do with AD machine accounts that are associated with the machine catalog.  Optional; default is . | [default to None]
 **async** | **bool** | If &#x60;true&#x60;, the machine catalog (and associated objects) will be deleted as a background task. The task will have JobType DeleteMachineCatalog. When the task is complete it will redirect to GetMachineCatalogs. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog being deleted, * _Name_ - Name of the machine catalog being deleted. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
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


## MachineCatalogsDoMachineCatalogSearch

> CatalogSearchResponseModelCollection MachineCatalogsDoMachineCatalogSearch(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CatalogSearchRequestModel(catalogSearchRequestModel).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    catalogSearchRequestModel := *openapiclient.NewCatalogSearchRequestModel() // CatalogSearchRequestModel | Specifies the advanced search parameters.
    async := true // bool | If `true`, the catalogs (and associated objects) will be added as a background task. The task will have JobType SearchMachineCatalogs. (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsDoMachineCatalogSearch(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CatalogSearchRequestModel(catalogSearchRequestModel).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsDoMachineCatalogSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsDoMachineCatalogSearch`: CatalogSearchResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsDoMachineCatalogSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsDoMachineCatalogSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **catalogSearchRequestModel** | [**CatalogSearchRequestModel**](CatalogSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **async** | **bool** | If &#x60;true&#x60;, the catalogs (and associated objects) will be added as a background task. The task will have JobType SearchMachineCatalogs. | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsGetMachineCatalog

> MachineCatalogDetailResponseModel MachineCatalogsGetMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    fields := "Name,FullName,Id" // string | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **fields** | **string** | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. | 
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


## MachineCatalogsGetMachineCatalogDeliveryGroupAssociations

> DeliveryGroupAssociationResponseModelCollection MachineCatalogsGetMachineCatalogDeliveryGroupAssociations(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of delivery group associations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogDeliveryGroupAssociations(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogDeliveryGroupAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogDeliveryGroupAssociations`: DeliveryGroupAssociationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogDeliveryGroupAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogDeliveryGroupAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of delivery group associations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsGetMachineCatalogLastMasterImage

> VMImageResponseModel MachineCatalogsGetMachineCatalogLastMasterImage(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogLastMasterImage(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogLastMasterImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogLastMasterImage`: VMImageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogLastMasterImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogLastMasterImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsGetMachineCatalogMachineAccounts

> ProvisioningSchemeMachineAccountResponseModelCollection MachineCatalogsGetMachineCatalogMachineAccounts(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Async(async).AccountState(accountState).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of machine accounts returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the machine accounts will be queried as a background task. The task will have JobType GetMachineCatalogMachineAccounts. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    accountState := openapiclient.MachineCatalogs_GetMachineCatalogMachineAccounts_accountState_parameter{ProvisioningSchemeMachineAccountState: openapiclient.ProvisioningSchemeMachineAccountState("Unknown")} // MachineCatalogsGetMachineCatalogMachineAccountsAccountStateParameter | The state of accounts for query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogMachineAccounts(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Async(async).AccountState(accountState).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogMachineAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogMachineAccounts`: ProvisioningSchemeMachineAccountResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogMachineAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to get machine accounts for.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogMachineAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of machine accounts returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine accounts will be queried as a background task. The task will have JobType GetMachineCatalogMachineAccounts. When the task is complete it will redirect to GetJobResults. | [default to false]
 **accountState** | [**MachineCatalogsGetMachineCatalogMachineAccountsAccountStateParameter**](MachineCatalogsGetMachineCatalogMachineAccountsAccountStateParameter.md) | The state of accounts for query. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsGetMachineCatalogMachines

> MachineResponseModelCollection MachineCatalogsGetMachineCatalogMachines(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    fields := "Uid,AgentVersion,AllocationType" // string | Optional parameters, only the specified properties in the fields are required. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogMachines`: MachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **fields** | **string** | Optional parameters, only the specified properties in the fields are required. | 
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


## MachineCatalogsGetMachineCatalogMasterImageHistory

> VMImageResponseModelCollection MachineCatalogsGetMachineCatalogMasterImageHistory(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogMasterImageHistory(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogMasterImageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogMasterImageHistory`: VMImageResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogMasterImageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogMasterImageHistoryRequest struct via the builder pattern


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

[**VMImageResponseModelCollection**](VMImageResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsGetMachineCatalogStartMenuApplications

> StartMenuApplicationResponseModelCollection MachineCatalogsGetMachineCatalogStartMenuApplications(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Machines(machines).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    machines := []string{"Inner_example"} // []string | The machines to get start menu applications (optional)
    async := true // bool | If `true`, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to \"JobsControllerTP.GetJobResults(string)\". The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog from which start menu applications are being obtained, * _Name_ - Name of the machine catalog from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogStartMenuApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Machines(machines).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogStartMenuApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogStartMenuApplications`: StartMenuApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogStartMenuApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogStartMenuApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **machines** | **[]string** | The machines to get start menu applications | 
 **async** | **bool** | If &#x60;true&#x60;, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to \&quot;JobsControllerTP.GetJobResults(string)\&quot;. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog from which start menu applications are being obtained, * _Name_ - Name of the machine catalog from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. | [default to false]
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


## MachineCatalogsGetMachineCatalogTags

> TagResponseModelCollection MachineCatalogsGetMachineCatalogTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    fields := "Name,Id,Description" // string | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **fields** | **string** | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. | 
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


## MachineCatalogsGetMachineCatalogTestReport

> TestReportResponseModel MachineCatalogsGetMachineCatalogTestReport(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogTestReport(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogTestReportRequest struct via the builder pattern


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


## MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures

> MachineCatalogVusComponentResponseModel MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures`: MachineCatalogVusComponentResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogVDAComponentsAndFeaturesRequest struct via the builder pattern


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

[**MachineCatalogVusComponentResponseModel**](MachineCatalogVusComponentResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsGetMachineCatalogs

> MachineCatalogResponseModelCollection MachineCatalogsGetMachineCatalogs(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    adminFolder := "adminFolder_example" // string | Admin folder path or Id. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Name,FullName,Id" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogs(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogs`: MachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **adminFolder** | **string** | Admin folder path or Id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server | 
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


## MachineCatalogsGetMachineCatalogsAdministrators

> AdministratorResponseModelCollection MachineCatalogsGetMachineCatalogsAdministrators(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogsAdministrators(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogsAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogsAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsGetMachineCatalogsAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogsAdministratorsRequest struct via the builder pattern


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

[**AdministratorResponseModelCollection**](AdministratorResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsRebootMachineCatalogMachines

> MachineCatalogsRebootMachineCatalogMachines(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootMachinesRequestModel(rebootMachinesRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    rebootMachinesRequestModel := *openapiclient.NewRebootMachinesRequestModel() // RebootMachinesRequestModel | Reboot options.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsRebootMachineCatalogMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootMachinesRequestModel(rebootMachinesRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsRebootMachineCatalogMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsRebootMachineCatalogMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **rebootMachinesRequestModel** | [**RebootMachinesRequestModel**](RebootMachinesRequestModel.md) | Reboot options. | 
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


## MachineCatalogsRemoveMachineCatalogMachine

> MachineCatalogsRemoveMachineCatalogMachine(ctx, nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsRemoveMachineCatalogMachine(context.Background(), nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsRemoveMachineCatalogMachine``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsRemoveMachineCatalogMachineRequest struct via the builder pattern


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


## MachineCatalogsRemoveMachineCatalogMachineAccount

> MachineCatalogsRemoveMachineCatalogMachineAccount(ctx, nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeleteAccount(deleteAccount).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    deleteAccount := openapiclient.MachineCatalogs_DeleteMachineCatalog_deleteAccount_parameter{MachineAccountDeleteOption: openapiclient.MachineAccountDeleteOption("Unknown")} // MachineCatalogsDeleteMachineCatalogDeleteAccountParameter | Specifies what to do with the machine account in AD.  Optional; default is None. (optional) (default to None)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsRemoveMachineCatalogMachineAccount(context.Background(), nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).DeleteAccount(deleteAccount).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsRemoveMachineCatalogMachineAccount``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsRemoveMachineCatalogMachineAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **deleteAccount** | [**MachineCatalogsDeleteMachineCatalogDeleteAccountParameter**](MachineCatalogsDeleteMachineCatalogDeleteAccountParameter.md) | Specifies what to do with the machine account in AD.  Optional; default is None. | [default to None]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
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


## MachineCatalogsRemoveMachineCatalogTags

> MachineCatalogsRemoveMachineCatalogTags(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsRemoveMachineCatalogTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsRemoveMachineCatalogTags``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsRemoveMachineCatalogTagsRequest struct via the builder pattern


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


## MachineCatalogsRemoveMachineCatalogUpgradeSchedule

> MachineCatalogsRemoveMachineCatalogUpgradeSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Force(force).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    force := true // bool | If `true`, VDA upgrade schedules existing with all machines within the catalog will be removed except those upgrade status in progressing. If `false`, only the perform removing when catalog level schedule not started. (optional) (default to false)
    async := true // bool | If `true`, the VDA upgrade schedule will be removed as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsRemoveMachineCatalogUpgradeSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Force(force).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsRemoveMachineCatalogUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsRemoveMachineCatalogUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **force** | **bool** | If &#x60;true&#x60;, VDA upgrade schedules existing with all machines within the catalog will be removed except those upgrade status in progressing. If &#x60;false&#x60;, only the perform removing when catalog level schedule not started. | [default to false]
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


## MachineCatalogsRollbackMachineCatalogProvisioningScheme

> JobResponseModel MachineCatalogsRollbackMachineCatalogProvisioningScheme(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootMachinesRequestModel(rebootMachinesRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    rebootMachinesRequestModel := *openapiclient.NewRebootMachinesRequestModel() // RebootMachinesRequestModel | Details about the image update to be performed.
    async := true // bool | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the `async` parameter is `true`, the only difference is that instead of returning the initial job status in the payload, the response is instead `202` with the `Location` header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsRollbackMachineCatalogProvisioningScheme(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RebootMachinesRequestModel(rebootMachinesRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsRollbackMachineCatalogProvisioningScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsRollbackMachineCatalogProvisioningScheme`: JobResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsRollbackMachineCatalogProvisioningScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsRollbackMachineCatalogProvisioningSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **rebootMachinesRequestModel** | [**RebootMachinesRequestModel**](RebootMachinesRequestModel.md) | Details about the image update to be performed. | 
 **async** | **bool** | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the &#x60;async&#x60; parameter is &#x60;true&#x60;, the only difference is that instead of returning the initial job status in the payload, the response is instead &#x60;202&#x60; with the &#x60;Location&#x60; header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsSetMachineCatalogTags

> MachineCatalogsSetMachineCatalogTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagsRequestModel(tagsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    tagsRequestModel := *openapiclient.NewTagsRequestModel() // TagsRequestModel | Details of the tags to set on the machine catalog.
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsSetMachineCatalogTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagsRequestModel(tagsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsSetMachineCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsSetMachineCatalogTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **tagsRequestModel** | [**TagsRequestModel**](TagsRequestModel.md) | Details of the tags to set on the machine catalog. | 
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


## MachineCatalogsTestMachineCatalog

> MachineCatalogTestResponseModel MachineCatalogsTestMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestMachineCatalog. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog being tested, * _Name_ - Name of the machine catalog being tested. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsTestMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsTestMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTestMachineCatalog`: MachineCatalogTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsTestMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to test.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTestMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestMachineCatalog. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog being tested, * _Name_ - Name of the machine catalog being tested. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsTestMachineCatalogExists

> MachineCatalogsTestMachineCatalogExists(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CatalogNameCheckRequestModel(catalogNameCheckRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    catalogNameCheckRequestModel := *openapiclient.NewCatalogNameCheckRequestModel() // CatalogNameCheckRequestModel | Name of the machine catalog.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsTestMachineCatalogExists(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CatalogNameCheckRequestModel(catalogNameCheckRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsTestMachineCatalogExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTestMachineCatalogExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **catalogNameCheckRequestModel** | [**CatalogNameCheckRequestModel**](CatalogNameCheckRequestModel.md) | Name of the machine catalog. | 
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


## MachineCatalogsTestRemotePCEnrollmentScopeExists

> MachineCatalogsTestRemotePCEnrollmentScopeExists(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RemotePCEnrollmentScopeCheckModel(remotePCEnrollmentScopeCheckModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    remotePCEnrollmentScopeCheckModel := *openapiclient.NewRemotePCEnrollmentScopeCheckModel("Name_example") // RemotePCEnrollmentScopeCheckModel | Request model that contains distinguished name of an OU containing machines allowed to enroll as remote PCs.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsTestRemotePCEnrollmentScopeExists(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).RemotePCEnrollmentScopeCheckModel(remotePCEnrollmentScopeCheckModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsTestRemotePCEnrollmentScopeExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTestRemotePCEnrollmentScopeExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **remotePCEnrollmentScopeCheckModel** | [**RemotePCEnrollmentScopeCheckModel**](RemotePCEnrollmentScopeCheckModel.md) | Request model that contains distinguished name of an OU containing machines allowed to enroll as remote PCs. | 
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


## MachineCatalogsUpdateMachineCatalog

> MachineCatalogDetailResponseModel MachineCatalogsUpdateMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateMachineCatalogRequestModel(updateMachineCatalogRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    updateMachineCatalogRequestModel := *openapiclient.NewUpdateMachineCatalogRequestModel() // UpdateMachineCatalogRequestModel | Properties of the machine catalog to update.
    async := true // bool | If `true`, the machine catalog (and associated objects) will be updated as a background task. The task will have JobType UpdateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog being updated, * _Name_ - Name of the machine catalog being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateMachineCatalogRequestModel(updateMachineCatalogRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsUpdateMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to update.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsUpdateMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **updateMachineCatalogRequestModel** | [**UpdateMachineCatalogRequestModel**](UpdateMachineCatalogRequestModel.md) | Properties of the machine catalog to update. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine catalog (and associated objects) will be updated as a background task. The task will have JobType UpdateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog being updated, * _Name_ - Name of the machine catalog being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsUpdateMachineCatalogMachineAccount

> ProvisioningSchemeMachineAccountResponseModel MachineCatalogsUpdateMachineCatalogMachineAccount(ctx, nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateMachineAccountRequestModel(updateMachineAccountRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    updateMachineAccountRequestModel := *openapiclient.NewUpdateMachineAccountRequestModel() // UpdateMachineAccountRequestModel | New details for the machine account record.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalogMachineAccount(context.Background(), nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateMachineAccountRequestModel(updateMachineAccountRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalogMachineAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsUpdateMachineCatalogMachineAccount`: ProvisioningSchemeMachineAccountResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalogMachineAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog containing the machine account.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**machine** | **string** | Machine account to update.  May be specified by SamName, UPN, or SID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsUpdateMachineCatalogMachineAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **updateMachineAccountRequestModel** | [**UpdateMachineAccountRequestModel**](UpdateMachineAccountRequestModel.md) | New details for the machine account record. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsUpdateMachineCatalogProvisioningScheme

> JobResponseModel MachineCatalogsUpdateMachineCatalogProvisioningScheme(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateMachineCatalogProvisioningSchemeRequestModel(updateMachineCatalogProvisioningSchemeRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    updateMachineCatalogProvisioningSchemeRequestModel := *openapiclient.NewUpdateMachineCatalogProvisioningSchemeRequestModel() // UpdateMachineCatalogProvisioningSchemeRequestModel | Details about the image update to be performed.
    async := true // bool | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the `async` parameter is `true`, the only difference is that instead of returning the initial job status in the payload, the response is instead `202` with the `Location` header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated, * _MasterImageXDPath_ - Path to the new image specified for the provisioning scheme, in XenApp & XenDesktop path format. * _MasterImageFullPath_ - Path to the new image specified for the provisioning scheme, relative to the root of the API. * _MasterImageRelativePath_ - Path to the new image specified for the provisioning scheme, relative to the resource pool. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalogProvisioningScheme(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateMachineCatalogProvisioningSchemeRequestModel(updateMachineCatalogProvisioningSchemeRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalogProvisioningScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsUpdateMachineCatalogProvisioningScheme`: JobResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalogProvisioningScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsUpdateMachineCatalogProvisioningSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **updateMachineCatalogProvisioningSchemeRequestModel** | [**UpdateMachineCatalogProvisioningSchemeRequestModel**](UpdateMachineCatalogProvisioningSchemeRequestModel.md) | Details about the image update to be performed. | 
 **async** | **bool** | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the &#x60;async&#x60; parameter is &#x60;true&#x60;, the only difference is that instead of returning the initial job status in the payload, the response is instead &#x60;202&#x60; with the &#x60;Location&#x60; header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated, * _MasterImageXDPath_ - Path to the new image specified for the provisioning scheme, in XenApp &amp; XenDesktop path format. * _MasterImageFullPath_ - Path to the new image specified for the provisioning scheme, relative to the root of the API. * _MasterImageRelativePath_ - Path to the new image specified for the provisioning scheme, relative to the resource pool. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## MachineCatalogsUpdateMachineCatalogUpgradeSchedule

> MachineCatalogsUpdateMachineCatalogUpgradeSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateUpgradeScheduleRequestModel(updateUpgradeScheduleRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    updateUpgradeScheduleRequestModel := *openapiclient.NewUpdateUpgradeScheduleRequestModel(time.Now(), int32(123)) // UpdateUpgradeScheduleRequestModel | Details about the updated VDA upgrade schedule.
    async := true // bool | If `true`, the VDA upgrade schedule will be updated as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalogUpgradeSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateUpgradeScheduleRequestModel(updateUpgradeScheduleRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsUpdateMachineCatalogUpgradeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsUpdateMachineCatalogUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **updateUpgradeScheduleRequestModel** | [**UpdateUpgradeScheduleRequestModel**](UpdateUpgradeScheduleRequestModel.md) | Details about the updated VDA upgrade schedule. | 
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


## MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection

> VDAComponentsSelectionValidationResponseModel MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).VDAComponentsSelectionValidationRequestModel(vDAComponentsSelectionValidationRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    vDAComponentsSelectionValidationRequestModel := *openapiclient.NewVDAComponentsSelectionValidationRequestModel() // VDAComponentsSelectionValidationRequestModel | More details attached to this validation.
    async := true // bool | If async execute. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIs .MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).VDAComponentsSelectionValidationRequestModel(vDAComponentsSelectionValidationRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIs .MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection`: VDAComponentsSelectionValidationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIs .MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **vDAComponentsSelectionValidationRequestModel** | [**VDAComponentsSelectionValidationRequestModel**](VDAComponentsSelectionValidationRequestModel.md) | More details attached to this validation. | 
 **async** | **bool** | If async execute. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

