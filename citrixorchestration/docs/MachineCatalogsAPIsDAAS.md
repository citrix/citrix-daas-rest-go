# \MachineCatalogsAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MachineCatalogsAddMachineCatalogMachine**](MachineCatalogsAPIsDAAS.md#MachineCatalogsAddMachineCatalogMachine) | **Post** /MachineCatalogs/{nameOrId}/Machines | Add machine to the machine catalog.
[**MachineCatalogsAddMachineCatalogMachineAccount**](MachineCatalogsAPIsDAAS.md#MachineCatalogsAddMachineCatalogMachineAccount) | **Post** /MachineCatalogs/{nameOrId}/MachineAccounts | Add a machine account to the machine catalog.
[**MachineCatalogsAddMachineCatalogTags**](MachineCatalogsAPIsDAAS.md#MachineCatalogsAddMachineCatalogTags) | **Post** /MachineCatalogs/{nameOrId}/Tags/{tagNameOrId} | Add a tag to a machine catalog.
[**MachineCatalogsCheckMachineCatalogExists**](MachineCatalogsAPIsDAAS.md#MachineCatalogsCheckMachineCatalogExists) | **Head** /MachineCatalogs/{name} | Check for the existence of a machine catalog by name.
[**MachineCatalogsCheckRemotePCEnrollmentScopeExists**](MachineCatalogsAPIsDAAS.md#MachineCatalogsCheckRemotePCEnrollmentScopeExists) | **Head** /MachineCatalogs/{nameOrId}/RemotePCEnrollmentScopes/{name} | Check for the existence of a RemotePC enrollment scope by name.
[**MachineCatalogsCreateMachineCatalog**](MachineCatalogsAPIsDAAS.md#MachineCatalogsCreateMachineCatalog) | **Post** /MachineCatalogs | Create a machine catalog.
[**MachineCatalogsCreateMachineCatalogUpgradeSchedule**](MachineCatalogsAPIsDAAS.md#MachineCatalogsCreateMachineCatalogUpgradeSchedule) | **Post** /MachineCatalogs/{nameOrId}/UpgradeSchedule | Setup a new VDA upgrade schedule for a machine catalog.
[**MachineCatalogsDeleteMachineCatalog**](MachineCatalogsAPIsDAAS.md#MachineCatalogsDeleteMachineCatalog) | **Delete** /MachineCatalogs/{nameOrId} | Delete a machine catalog.
[**MachineCatalogsDoMachineCatalogSearch**](MachineCatalogsAPIsDAAS.md#MachineCatalogsDoMachineCatalogSearch) | **Post** /Machinecatalogs/$search | Perform an advanced search for machinecatalogs.
[**MachineCatalogsExport**](MachineCatalogsAPIsDAAS.md#MachineCatalogsExport) | **Post** /MachineCatalogs/{nameOrId}/$export | Exports a MCS machine catalog.
[**MachineCatalogsExportMachineProfile**](MachineCatalogsAPIsDAAS.md#MachineCatalogsExportMachineProfile) | **Post** /MachineCatalogs/{nameOrId}/$exportMachineProfile | Exports a MCS machine catalog&#39;s machine profile.
[**MachineCatalogsGetMachineCatalog**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalog) | **Get** /MachineCatalogs/{nameOrId} | Get details about a single machine catalog.
[**MachineCatalogsGetMachineCatalogDeliveryGroupAssociations**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogDeliveryGroupAssociations) | **Get** /MachineCatalogs/{nameOrId}/DeliveryGroupAssociations | Get delivery group associations of a machine catalog.
[**MachineCatalogsGetMachineCatalogEnrollments**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogEnrollments) | **Get** /MachineCatalogs/{nameOrId}/Enrollments | Get the list of enrollments in the machine catalog.
[**MachineCatalogsGetMachineCatalogLastMasterImage**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogLastMasterImage) | **Get** /MachineCatalogs/{nameOrId}/LastMasterImage | Get the last master VM images used by the machine catalog.
[**MachineCatalogsGetMachineCatalogMachineAccounts**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogMachineAccounts) | **Get** /MachineCatalogs/{nameOrId}/MachineAccounts | Get all machine accounts associated with a machine catalog.
[**MachineCatalogsGetMachineCatalogMachines**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogMachines) | **Get** /MachineCatalogs/{nameOrId}/Machines | GET the machines of a catalog
[**MachineCatalogsGetMachineCatalogMasterImageHistory**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogMasterImageHistory) | **Get** /MachineCatalogs/{nameOrId}/MasterImageHistory | Get the history of master VM images used by the machine catalog.
[**MachineCatalogsGetMachineCatalogStartMenuApplications**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogStartMenuApplications) | **Get** /MachineCatalogs/{nameOrId}/StartMenuApplications | Get the start menu applications from a machine in the machine catalog.
[**MachineCatalogsGetMachineCatalogTags**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogTags) | **Get** /MachineCatalogs/{nameOrId}/Tags | Get tags associated with a machine catalog.
[**MachineCatalogsGetMachineCatalogTestReport**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogTestReport) | **Get** /MachineCatalogs/{nameOrId}/TestReport | Get the most recent test report of a machine catalog.
[**MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures) | **Get** /MachineCatalogs/{nameOrId}/VDAComponentsAndFeatures | Get the components and features of VDAs associated with a machine catalog.
[**MachineCatalogsGetMachineCatalogVDAUpgradeVersions**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogVDAUpgradeVersions) | **Get** /MachineCatalogs/{nameOrId}/VDAUpgradeVersions | Get the available VDA upgrade versions associated with a machine catalog.
[**MachineCatalogsGetMachineCatalogs**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogs) | **Get** /MachineCatalogs | Get all machine catalogs.
[**MachineCatalogsGetMachineCatalogsAdministrators**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogsAdministrators) | **Get** /MachineCatalogs/{nameOrId}/Administrators | Get administrators who can administer a machine catalog.
[**MachineCatalogsGetMachineCatalogsV2**](MachineCatalogsAPIsDAAS.md#MachineCatalogsGetMachineCatalogsV2) | **Get** /MachineCatalogsV2 | The V2  version of get all machine catalogs.
[**MachineCatalogsImportProvisionedVirtualMachines**](MachineCatalogsAPIsDAAS.md#MachineCatalogsImportProvisionedVirtualMachines) | **Post** /MachineCatalogs/{nameOrId}/$importProvisionedVirtualMachines | Import existing provisioned virtual machines.
[**MachineCatalogsRebootMachineCatalogMachines**](MachineCatalogsAPIsDAAS.md#MachineCatalogsRebootMachineCatalogMachines) | **Post** /MachineCatalogs/{nameOrId}/$RebootMachines | Reboot all machines in a machine catalog.
[**MachineCatalogsRemoveMachineCatalogMachine**](MachineCatalogsAPIsDAAS.md#MachineCatalogsRemoveMachineCatalogMachine) | **Delete** /MachineCatalogs/{nameOrId}/Machines/{machine} | Remove a machine from the machine catalog.
[**MachineCatalogsRemoveMachineCatalogMachineAccount**](MachineCatalogsAPIsDAAS.md#MachineCatalogsRemoveMachineCatalogMachineAccount) | **Delete** /MachineCatalogs/{nameOrId}/MachineAccounts/{machine} | Remove a machine account from the machine catalog.
[**MachineCatalogsRemoveMachineCatalogTags**](MachineCatalogsAPIsDAAS.md#MachineCatalogsRemoveMachineCatalogTags) | **Delete** /MachineCatalogs/{nameOrId}/Tags/{tagNameOrId} | Remove a tag from a machine catalog.
[**MachineCatalogsRemoveMachineCatalogUpgradeSchedule**](MachineCatalogsAPIsDAAS.md#MachineCatalogsRemoveMachineCatalogUpgradeSchedule) | **Delete** /MachineCatalogs/{nameOrId}/UpgradeSchedule | Remove a pending VDA upgrade schedule for a machine catalog.
[**MachineCatalogsRepairMachineCatalogMachineAccounts**](MachineCatalogsAPIsDAAS.md#MachineCatalogsRepairMachineCatalogMachineAccounts) | **Post** /MachineCatalogs/{nameOrId}/$repairMachineAccounts | Repair all machine accounts in a machine catalog.
[**MachineCatalogsRollbackMachineCatalogProvisioningScheme**](MachineCatalogsAPIsDAAS.md#MachineCatalogsRollbackMachineCatalogProvisioningScheme) | **Post** /MachineCatalogs/{nameOrId}/$RollbackProvisioningScheme | Rollback the machine catalog&#39;s provisioning scheme.
[**MachineCatalogsSetMachineCatalogTags**](MachineCatalogsAPIsDAAS.md#MachineCatalogsSetMachineCatalogTags) | **Put** /MachineCatalogs/{nameOrId}/Tags | Set tags associated with a machine catalog.
[**MachineCatalogsTestMachineCatalog**](MachineCatalogsAPIsDAAS.md#MachineCatalogsTestMachineCatalog) | **Post** /MachineCatalogs/{nameOrId}/$test | Run tests on a machine catalog and create a test report.
[**MachineCatalogsTestMachineCatalogExists**](MachineCatalogsAPIsDAAS.md#MachineCatalogsTestMachineCatalogExists) | **Post** /MachineCatalogs/$checkCatalogName | Test the existence of a machine catalog by name.
[**MachineCatalogsTestRemotePCEnrollmentScopeExists**](MachineCatalogsAPIsDAAS.md#MachineCatalogsTestRemotePCEnrollmentScopeExists) | **Post** /MachineCatalogs/{nameOrId}/RemotePCEnrollmentScopes/$checkScopeName | Check for the existence of a RemotePC enrollment scope by name in the request body. Use HttpPost and take name in body as this name may be very long, and cause url too long exception.
[**MachineCatalogsUpdateMachineCatalog**](MachineCatalogsAPIsDAAS.md#MachineCatalogsUpdateMachineCatalog) | **Patch** /MachineCatalogs/{nameOrId} | Update a machine catalog.
[**MachineCatalogsUpdateMachineCatalogMachineAccount**](MachineCatalogsAPIsDAAS.md#MachineCatalogsUpdateMachineCatalogMachineAccount) | **Patch** /MachineCatalogs/{nameOrId}/MachineAccounts/{machine} | Update a machine account within the machine catalog.
[**MachineCatalogsUpdateMachineCatalogProvisioningScheme**](MachineCatalogsAPIsDAAS.md#MachineCatalogsUpdateMachineCatalogProvisioningScheme) | **Post** /MachineCatalogs/{nameOrId}/$UpdateProvisioningScheme | Update the machine catalog&#39;s provisioning scheme.
[**MachineCatalogsUpdateMachineCatalogUpgradeSchedule**](MachineCatalogsAPIsDAAS.md#MachineCatalogsUpdateMachineCatalogUpgradeSchedule) | **Patch** /MachineCatalogs/{nameOrId}/UpgradeSchedule | Reschedule a pending VDA upgrade schedule for a machine catalog.
[**MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection**](MachineCatalogsAPIsDAAS.md#MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection) | **Post** /MachineCatalogs/{nameOrId}/VDAComponentsAndFeatures/$validate | Validate the VDA component selection at catalog level that can be Included/Excluded and features that need enable during an VDA upgrade using VDA Upgrade Service.



## MachineCatalogsAddMachineCatalogMachine

> JobResponseModel MachineCatalogsAddMachineCatalogMachine(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).AddMachineToMachineCatalogDetailRequestModel(addMachineToMachineCatalogDetailRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to add machine to.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    addMachineToMachineCatalogDetailRequestModel := *openapiclient.NewAddMachineToMachineCatalogDetailRequestModel() // AddMachineToMachineCatalogDetailRequestModel | Details of the machine to add to the machine catalog.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the machine will be added into the machine catalog as a background task. The task will have JobType AddMachineCatalogMachine. When the task is complete it will redirect to GetMachineCatalog. The job's Parameters will contain properties: * _NameOrId_ - Name or ID of the machine being added. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsAddMachineCatalogMachine(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).AddMachineToMachineCatalogDetailRequestModel(addMachineToMachineCatalogDetailRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsAddMachineCatalogMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsAddMachineCatalogMachine`: JobResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsAddMachineCatalogMachine`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **addMachineToMachineCatalogDetailRequestModel** | [**AddMachineToMachineCatalogDetailRequestModel**](AddMachineToMachineCatalogDetailRequestModel.md) | Details of the machine to add to the machine catalog. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine will be added into the machine catalog as a background task. The task will have JobType AddMachineCatalogMachine. When the task is complete it will redirect to GetMachineCatalog. The job&#39;s Parameters will contain properties: * _NameOrId_ - Name or ID of the machine being added. | [default to false]

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

> ProvisioningSchemeMachineAccountResponseModel MachineCatalogsAddMachineCatalogMachineAccount(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).MachineAccountRequestModel(machineAccountRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to add the machine account to.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    machineAccountRequestModel := *openapiclient.NewMachineAccountRequestModel("For example, MyDomain\MyComputer") // MachineAccountRequestModel | Machine account to add to the machine catalog.  The machine must not be a member of any other machine catalog.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsAddMachineCatalogMachineAccount(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).MachineAccountRequestModel(machineAccountRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsAddMachineCatalogMachineAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsAddMachineCatalogMachineAccount`: ProvisioningSchemeMachineAccountResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsAddMachineCatalogMachineAccount`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **machineAccountRequestModel** | [**MachineAccountRequestModel**](MachineAccountRequestModel.md) | Machine account to add to the machine catalog.  The machine must not be a member of any other machine catalog. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
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

> TagResponseModelCollection MachineCatalogsAddMachineCatalogTags(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag to add.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsAddMachineCatalogTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsAddMachineCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsAddMachineCatalogTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsAddMachineCatalogTags`: %v\n", resp)
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


## MachineCatalogsCheckMachineCatalogExists

> MachineCatalogsCheckMachineCatalogExists(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    name := "name_example" // string | Name of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsCheckMachineCatalogExists(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsCheckMachineCatalogExists``: %v\n", err)
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


## MachineCatalogsCheckRemotePCEnrollmentScopeExists

> MachineCatalogsCheckRemotePCEnrollmentScopeExists(ctx, nameOrId, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    name := "name_example" // string | Distinguished name of an OU containing machines allowed to enroll as remote PCs.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsCheckRemotePCEnrollmentScopeExists(context.Background(), nameOrId, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsCheckRemotePCEnrollmentScopeExists``: %v\n", err)
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


## MachineCatalogsCreateMachineCatalog

> MachineCatalogDetailResponseModel MachineCatalogsCreateMachineCatalog(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateMachineCatalogRequestModel(createMachineCatalogRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Async(async).SourceNameOrId(sourceNameOrId).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    createMachineCatalogRequestModel := *openapiclient.NewCreateMachineCatalogRequestModel("Name of the machine catalog", openapiclient.ProvisioningType("Unknown"), openapiclient.SessionSupport("Unknown")) // CreateMachineCatalogRequestModel | Details about the machine catalog to create.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the machine catalog (and associated objects) will be created as a background task. The task will have JobType CreateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job's Parameters will contain properties: * _Name_ - Name of the machine catalog being created. (optional) (default to false)
    sourceNameOrId := "sourceNameOrId_example" // string | The machine catalog name or Id to recreate from. It is only used for catalog recreation scenario. Here specify name or Id of the failed catalog to recreate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsCreateMachineCatalog(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateMachineCatalogRequestModel(createMachineCatalogRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Async(async).SourceNameOrId(sourceNameOrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsCreateMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsCreateMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsCreateMachineCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsCreateMachineCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createMachineCatalogRequestModel** | [**CreateMachineCatalogRequestModel**](CreateMachineCatalogRequestModel.md) | Details about the machine catalog to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine catalog (and associated objects) will be created as a background task. The task will have JobType CreateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job&#39;s Parameters will contain properties: * _Name_ - Name of the machine catalog being created. | [default to false]
 **sourceNameOrId** | **string** | The machine catalog name or Id to recreate from. It is only used for catalog recreation scenario. Here specify name or Id of the failed catalog to recreate. | 

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

> MachineCatalogsCreateMachineCatalogUpgradeSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateUpgradeScheduleRequestModel(createUpgradeScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    createUpgradeScheduleRequestModel := *openapiclient.NewCreateUpgradeScheduleRequestModel(int32(123)) // CreateUpgradeScheduleRequestModel | Details about the VDA upgrade schedule to create.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the VDA upgrade schedule will be created as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsCreateMachineCatalogUpgradeSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateUpgradeScheduleRequestModel(createUpgradeScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsCreateMachineCatalogUpgradeSchedule``: %v\n", err)
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


## MachineCatalogsDeleteMachineCatalog

> MachineCatalogsDeleteMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).DeleteVm(deleteVm).DeleteAccount(deleteAccount).PurgeDBOnly(purgeDBOnly).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to update.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    deleteVm := true // bool | Specifies whether to delete virtual machines that were provisioned in the machine catalog. Only relevant when ProvisioningType is MCS; otherwise, ignored.  Optional; default is `true`. (optional) (default to true)
    deleteAccount := openapiclient.MachineAccountDeleteOption("Unknown") // MachineAccountDeleteOption | Specifies what to do with AD machine accounts that are associated with the machine catalog.  Optional; default is . (optional) (default to "None")
    purgeDBOnly := true // bool | Specifies whether to only delete virtual machines record that were provisioned in the machine catalog from database, not access hypervisor to delete virtual machines, this is recommended only when hypervisor cannot be accessed.  Optional; default is `false`. (optional) (default to false)
    async := true // bool | If `true`, the machine catalog (and associated objects) will be deleted as a background task. The task will have JobType DeleteMachineCatalog. When the task is complete it will redirect to GetMachineCatalogs. The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog being deleted, * _Name_ - Name of the machine catalog being deleted. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsDeleteMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).DeleteVm(deleteVm).DeleteAccount(deleteAccount).PurgeDBOnly(purgeDBOnly).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsDeleteMachineCatalog``: %v\n", err)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **deleteVm** | **bool** | Specifies whether to delete virtual machines that were provisioned in the machine catalog. Only relevant when ProvisioningType is MCS; otherwise, ignored.  Optional; default is &#x60;true&#x60;. | [default to true]
 **deleteAccount** | [**MachineAccountDeleteOption**](MachineAccountDeleteOption.md) | Specifies what to do with AD machine accounts that are associated with the machine catalog.  Optional; default is . | [default to &quot;None&quot;]
 **purgeDBOnly** | **bool** | Specifies whether to only delete virtual machines record that were provisioned in the machine catalog from database, not access hypervisor to delete virtual machines, this is recommended only when hypervisor cannot be accessed.  Optional; default is &#x60;false&#x60;. | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, the machine catalog (and associated objects) will be deleted as a background task. The task will have JobType DeleteMachineCatalog. When the task is complete it will redirect to GetMachineCatalogs. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog being deleted, * _Name_ - Name of the machine catalog being deleted. | [default to false]

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

> CatalogSearchResponseModelCollection MachineCatalogsDoMachineCatalogSearch(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CatalogSearchRequestModel(catalogSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    catalogSearchRequestModel := *openapiclient.NewCatalogSearchRequestModel() // CatalogSearchRequestModel | Specifies the advanced search parameters.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the catalogs (and associated objects) will be added as a background task. The task will have JobType SearchMachineCatalogs. (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsDoMachineCatalogSearch(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CatalogSearchRequestModel(catalogSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsDoMachineCatalogSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsDoMachineCatalogSearch`: CatalogSearchResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsDoMachineCatalogSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsDoMachineCatalogSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **catalogSearchRequestModel** | [**CatalogSearchRequestModel**](CatalogSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the catalogs (and associated objects) will be added as a background task. The task will have JobType SearchMachineCatalogs. | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 

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


## MachineCatalogsExport

> ExportMachineCatalogResponseModel MachineCatalogsExport(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()

Exports a MCS machine catalog.

### Example

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
    nameOrId := "nameOrId_example" // string | The machine catalog
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    fields := "fields_example" // string | The required fields (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsExport(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsExport`: ExportMachineCatalogResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The machine catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **fields** | **string** | The required fields | 

### Return type

[**ExportMachineCatalogResponseModel**](ExportMachineCatalogResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsExportMachineProfile

> ExportMachineProfileResponseModel MachineCatalogsExportMachineProfile(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Exports a MCS machine catalog's machine profile.

### Example

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
    nameOrId := "nameOrId_example" // string | The name or id of machine catalog.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsExportMachineProfile(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsExportMachineProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsExportMachineProfile`: ExportMachineProfileResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsExportMachineProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The name or id of machine catalog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsExportMachineProfileRequest struct via the builder pattern


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

[**ExportMachineProfileResponseModel**](ExportMachineProfileResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsGetMachineCatalog

> MachineCatalogDetailResponseModel MachineCatalogsGetMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Fields(fields).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    fields := "Name,FullName,Id" // string | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalog`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **fields** | **string** | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. | 

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

> DeliveryGroupAssociationResponseModelCollection MachineCatalogsGetMachineCatalogDeliveryGroupAssociations(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    limit := int32(56) // int32 | The max number of delivery group associations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogDeliveryGroupAssociations(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogDeliveryGroupAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogDeliveryGroupAssociations`: DeliveryGroupAssociationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogDeliveryGroupAssociations`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of delivery group associations returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]

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


## MachineCatalogsGetMachineCatalogEnrollments

> EnrollmentResponseModelCollection MachineCatalogsGetMachineCatalogEnrollments(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the list of enrollments in the machine catalog.

### Example

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogEnrollments(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogEnrollments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogEnrollments`: EnrollmentResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogEnrollments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogEnrollmentsRequest struct via the builder pattern


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

[**EnrollmentResponseModelCollection**](EnrollmentResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsGetMachineCatalogLastMasterImage

> VMImageResponseModel MachineCatalogsGetMachineCatalogLastMasterImage(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogLastMasterImage(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogLastMasterImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogLastMasterImage`: VMImageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogLastMasterImage`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]

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

> ProvisioningSchemeMachineAccountResponseModelCollection MachineCatalogsGetMachineCatalogMachineAccounts(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).AccountState(accountState).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to get machine accounts for.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    limit := int32(56) // int32 | The max number of machine accounts returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the machine accounts will be queried as a background task. The task will have JobType GetMachineCatalogMachineAccounts. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    accountState := openapiclient.ProvisioningSchemeMachineAccountState("Unknown") // ProvisioningSchemeMachineAccountState | The state of accounts for query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogMachineAccounts(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).AccountState(accountState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogMachineAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogMachineAccounts`: ProvisioningSchemeMachineAccountResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogMachineAccounts`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of machine accounts returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine accounts will be queried as a background task. The task will have JobType GetMachineCatalogMachineAccounts. When the task is complete it will redirect to GetJobResults. | [default to false]
 **accountState** | [**ProvisioningSchemeMachineAccountState**](ProvisioningSchemeMachineAccountState.md) | The state of accounts for query. | 

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

> MachineResponseModelCollection MachineCatalogsGetMachineCatalogMachines(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    limit := int32(56) // int32 | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    fields := "Uid,AgentVersion,AllocationType" // string | Optional parameters, only the specified properties in the fields are required. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogMachines`: MachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogMachines`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of machines returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **fields** | **string** | Optional parameters, only the specified properties in the fields are required. | 

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

> VMImageResponseModelCollection MachineCatalogsGetMachineCatalogMasterImageHistory(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogMasterImageHistory(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogMasterImageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogMasterImageHistory`: VMImageResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogMasterImageHistory`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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

> StartMenuApplicationResponseModelCollection MachineCatalogsGetMachineCatalogStartMenuApplications(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Machines(machines).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    machines := []string{"Inner_example"} // []string | The machines to get start menu applications (optional)
    async := true // bool | If `true`, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to \"JobsControllerTP.GetJobResults(string)\". The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog from which start menu applications are being obtained, * _Name_ - Name of the machine catalog from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogStartMenuApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Machines(machines).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogStartMenuApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogStartMenuApplications`: StartMenuApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogStartMenuApplications`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **machines** | **[]string** | The machines to get start menu applications | 
 **async** | **bool** | If &#x60;true&#x60;, the start menu applications will be queried as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType GetStartMenuApplications. When the task is complete it will redirect to \&quot;JobsControllerTP.GetJobResults(string)\&quot;. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog from which start menu applications are being obtained, * _Name_ - Name of the machine catalog from which start menu applications are being obtained. * _MachineId_ - ID of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the start menu applications are being obtained; will be present in Parameters only after a machine is selected. | [default to false]

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

> TagResponseModelCollection MachineCatalogsGetMachineCatalogTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    fields := "Name,Id,Description" // string | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogTags`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **fields** | **string** | Optional parameters, removing unspecified properties that otherwise would have been sent by the server. | 

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

> TestReportResponseModel MachineCatalogsGetMachineCatalogTestReport(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogTestReport(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogTestReport`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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

> MachineCatalogVusComponentResponseModel MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).UpgradeVersion(upgradeVersion).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    upgradeVersion := "upgradeVersion_example" // string | The version of the VDA to upgrade to. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).UpgradeVersion(upgradeVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures`: MachineCatalogVusComponentResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogVDAComponentsAndFeatures`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **upgradeVersion** | **string** | The version of the VDA to upgrade to. | 

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


## MachineCatalogsGetMachineCatalogVDAUpgradeVersions

> []string MachineCatalogsGetMachineCatalogVDAUpgradeVersions(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the available VDA upgrade versions associated with a machine catalog.

### Example

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogVDAUpgradeVersions(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogVDAUpgradeVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogVDAUpgradeVersions`: []string
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogVDAUpgradeVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogVDAUpgradeVersionsRequest struct via the builder pattern


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


## MachineCatalogsGetMachineCatalogs

> MachineCatalogResponseModelCollection MachineCatalogsGetMachineCatalogs(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    adminFolder := "adminFolder_example" // string | Admin folder path or Id. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Name,FullName,Id" // string | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma','.              (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogs(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogs`: MachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **adminFolder** | **string** | Admin folder path or Id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma&#39;,&#39;.              | 

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

> AdministratorResponseModelCollection MachineCatalogsGetMachineCatalogsAdministrators(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogsAdministrators(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogsAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogsAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogsAdministrators`: %v\n", resp)
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


## MachineCatalogsGetMachineCatalogsV2

> MachineCatalogResponseModelCollection MachineCatalogsGetMachineCatalogsV2(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Execute()

The V2  version of get all machine catalogs.



### Example

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
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    adminFolder := "adminFolder_example" // string | Admin folder path or Id. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Name,FullName,Id" // string | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma','.              (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogsV2(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).AdminFolder(adminFolder).Async(async).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsGetMachineCatalogsV2`: MachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsGetMachineCatalogsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsGetMachineCatalogsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **adminFolder** | **string** | Admin folder path or Id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma&#39;,&#39;.              | 

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


## MachineCatalogsImportProvisionedVirtualMachines

> ImportProvisionedVirtualMachineResponseModelCollection MachineCatalogsImportProvisionedVirtualMachines(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).ImportProvisionedVirtualMachinesRequestModel(importProvisionedVirtualMachinesRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Import existing provisioned virtual machines.

### Example

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
    nameOrId := "nameOrId_example" // string | The machineCatalog name or Id. 
    importProvisionedVirtualMachinesRequestModel := *openapiclient.NewImportProvisionedVirtualMachinesRequestModel() // ImportProvisionedVirtualMachinesRequestModel | The existing provisioned virtual machines data. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the machine accounts repair will be performed as a background task.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsImportProvisionedVirtualMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).ImportProvisionedVirtualMachinesRequestModel(importProvisionedVirtualMachinesRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsImportProvisionedVirtualMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsImportProvisionedVirtualMachines`: ImportProvisionedVirtualMachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsImportProvisionedVirtualMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The machineCatalog name or Id.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsImportProvisionedVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **importProvisionedVirtualMachinesRequestModel** | [**ImportProvisionedVirtualMachinesRequestModel**](ImportProvisionedVirtualMachinesRequestModel.md) | The existing provisioned virtual machines data.  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine accounts repair will be performed as a background task.  | [default to false]

### Return type

[**ImportProvisionedVirtualMachineResponseModelCollection**](ImportProvisionedVirtualMachineResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsRebootMachineCatalogMachines

> MachineCatalogsRebootMachineCatalogMachines(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RebootMachinesRequestModel(rebootMachinesRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    rebootMachinesRequestModel := *openapiclient.NewRebootMachinesRequestModel() // RebootMachinesRequestModel | Reboot options.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsRebootMachineCatalogMachines(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RebootMachinesRequestModel(rebootMachinesRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsRebootMachineCatalogMachines``: %v\n", err)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **rebootMachinesRequestModel** | [**RebootMachinesRequestModel**](RebootMachinesRequestModel.md) | Reboot options. | 
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

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsRemoveMachineCatalogMachine

> MachineCatalogsRemoveMachineCatalogMachine(ctx, nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to remove the machine from.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    machine := "machine_example" // string | Machine to remove from the catalog. May be specified by name, ID, or SID. If param is Name, currently it should get rid of '\\\\' and replace it using '|'.             For instance, if a MachineName is \"DomainA\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsRemoveMachineCatalogMachine(context.Background(), nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsRemoveMachineCatalogMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog to remove the machine from.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName. | 
**machine** | **string** | Machine to remove from the catalog. May be specified by name, ID, or SID. If param is Name, currently it should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;.             For instance, if a MachineName is \&quot;DomainA\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsRemoveMachineCatalogMachineRequest struct via the builder pattern


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


## MachineCatalogsRemoveMachineCatalogMachineAccount

> MachineCatalogsRemoveMachineCatalogMachineAccount(ctx, nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DeleteAccount(deleteAccount).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to remove the machine account from.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    machine := "machine_example" // string | Machine account to remove.  May be specified by SamName, UPN, or SID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    deleteAccount := openapiclient.MachineAccountDeleteOption("Unknown") // MachineAccountDeleteOption | Specifies what to do with the machine account in AD.  Optional; default is None. (optional) (default to "None")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsRemoveMachineCatalogMachineAccount(context.Background(), nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DeleteAccount(deleteAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsRemoveMachineCatalogMachineAccount``: %v\n", err)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **deleteAccount** | [**MachineAccountDeleteOption**](MachineAccountDeleteOption.md) | Specifies what to do with the machine account in AD.  Optional; default is None. | [default to &quot;None&quot;]

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

> MachineCatalogsRemoveMachineCatalogTags(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsRemoveMachineCatalogTags(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsRemoveMachineCatalogTags``: %v\n", err)
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


## MachineCatalogsRemoveMachineCatalogUpgradeSchedule

> MachineCatalogsRemoveMachineCatalogUpgradeSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    force := true // bool | If `true`, VDA upgrade schedules existing with all machines within the catalog will be removed except those upgrade status in progressing. If `false`, only the perform removing when catalog level schedule not started. (optional) (default to false)
    async := true // bool | If `true`, the VDA upgrade schedule will be removed as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsRemoveMachineCatalogUpgradeSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsRemoveMachineCatalogUpgradeSchedule``: %v\n", err)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **force** | **bool** | If &#x60;true&#x60;, VDA upgrade schedules existing with all machines within the catalog will be removed except those upgrade status in progressing. If &#x60;false&#x60;, only the perform removing when catalog level schedule not started. | [default to false]
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


## MachineCatalogsRepairMachineCatalogMachineAccounts

> MachineCatalogsRepairMachineCatalogMachineAccounts(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RepairOption(repairOption).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Repair all machine accounts in a machine catalog.



### Example

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
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog.
    repairOption := openapiclient.MachineAccountRepairOption("Unknown") // MachineAccountRepairOption | The repair option of the machine accounts
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the machine accounts repair will be performed as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsRepairMachineCatalogMachineAccounts(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RepairOption(repairOption).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsRepairMachineCatalogMachineAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the machine catalog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsRepairMachineCatalogMachineAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **repairOption** | [**MachineAccountRepairOption**](MachineAccountRepairOption.md) | The repair option of the machine accounts | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine accounts repair will be performed as a background task. | [default to false]

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

> JobResponseModel MachineCatalogsRollbackMachineCatalogProvisioningScheme(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RebootMachinesRequestModel(rebootMachinesRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    rebootMachinesRequestModel := *openapiclient.NewRebootMachinesRequestModel() // RebootMachinesRequestModel | Details about the image update to be performed.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the `async` parameter is `true`, the only difference is that instead of returning the initial job status in the payload, the response is instead `202` with the `Location` header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsRollbackMachineCatalogProvisioningScheme(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RebootMachinesRequestModel(rebootMachinesRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsRollbackMachineCatalogProvisioningScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsRollbackMachineCatalogProvisioningScheme`: JobResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsRollbackMachineCatalogProvisioningScheme`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **rebootMachinesRequestModel** | [**RebootMachinesRequestModel**](RebootMachinesRequestModel.md) | Details about the image update to be performed. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the &#x60;async&#x60; parameter is &#x60;true&#x60;, the only difference is that instead of returning the initial job status in the payload, the response is instead &#x60;202&#x60; with the &#x60;Location&#x60; header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated. | [default to false]

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

> MachineCatalogsSetMachineCatalogTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).TagsRequestModel(tagsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    tagsRequestModel := *openapiclient.NewTagsRequestModel() // TagsRequestModel | Details of the tags to set on the machine catalog.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tags will be modified as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsSetMachineCatalogTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).TagsRequestModel(tagsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsSetMachineCatalogTags``: %v\n", err)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **tagsRequestModel** | [**TagsRequestModel**](TagsRequestModel.md) | Details of the tags to set on the machine catalog. | 
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


## MachineCatalogsTestMachineCatalog

> MachineCatalogTestResponseModel MachineCatalogsTestMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to test.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestMachineCatalog. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog being tested, * _Name_ - Name of the machine catalog being tested. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsTestMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsTestMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsTestMachineCatalog`: MachineCatalogTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsTestMachineCatalog`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestMachineCatalog. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog being tested, * _Name_ - Name of the machine catalog being tested. | [default to false]

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

> MachineCatalogsTestMachineCatalogExists(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CatalogNameCheckRequestModel(catalogNameCheckRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    catalogNameCheckRequestModel := *openapiclient.NewCatalogNameCheckRequestModel() // CatalogNameCheckRequestModel | Name of the machine catalog.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsTestMachineCatalogExists(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CatalogNameCheckRequestModel(catalogNameCheckRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsTestMachineCatalogExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMachineCatalogsTestMachineCatalogExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **catalogNameCheckRequestModel** | [**CatalogNameCheckRequestModel**](CatalogNameCheckRequestModel.md) | Name of the machine catalog. | 
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

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsTestRemotePCEnrollmentScopeExists

> MachineCatalogsTestRemotePCEnrollmentScopeExists(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RemotePCEnrollmentScopeCheckModel(remotePCEnrollmentScopeCheckModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    remotePCEnrollmentScopeCheckModel := *openapiclient.NewRemotePCEnrollmentScopeCheckModel("Name_example") // RemotePCEnrollmentScopeCheckModel | Request model that contains distinguished name of an OU containing machines allowed to enroll as remote PCs.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsTestRemotePCEnrollmentScopeExists(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RemotePCEnrollmentScopeCheckModel(remotePCEnrollmentScopeCheckModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsTestRemotePCEnrollmentScopeExists``: %v\n", err)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **remotePCEnrollmentScopeCheckModel** | [**RemotePCEnrollmentScopeCheckModel**](RemotePCEnrollmentScopeCheckModel.md) | Request model that contains distinguished name of an OU containing machines allowed to enroll as remote PCs. | 
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

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineCatalogsUpdateMachineCatalog

> MachineCatalogDetailResponseModel MachineCatalogsUpdateMachineCatalog(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMachineCatalogRequestModel(updateMachineCatalogRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog to update.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    updateMachineCatalogRequestModel := *openapiclient.NewUpdateMachineCatalogRequestModel() // UpdateMachineCatalogRequestModel | Properties of the machine catalog to update.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the machine catalog (and associated objects) will be updated as a background task. The task will have JobType UpdateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job's Parameters will contain properties:  * _Id_ - ID of the machine catalog being updated, * _Name_ - Name of the machine catalog being updated. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalog(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMachineCatalogRequestModel(updateMachineCatalogRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsUpdateMachineCatalog`: MachineCatalogDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalog`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **updateMachineCatalogRequestModel** | [**UpdateMachineCatalogRequestModel**](UpdateMachineCatalogRequestModel.md) | Properties of the machine catalog to update. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine catalog (and associated objects) will be updated as a background task. The task will have JobType UpdateMachineCatalog. When the task is complete it will redirect to GetMachineCatalog. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the machine catalog being updated, * _Name_ - Name of the machine catalog being updated. | [default to false]

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

> ProvisioningSchemeMachineAccountResponseModel MachineCatalogsUpdateMachineCatalogMachineAccount(ctx, nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMachineAccountRequestModel(updateMachineAccountRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog containing the machine account.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    machine := "machine_example" // string | Machine account to update.  May be specified by SamName, UPN, or SID.
    updateMachineAccountRequestModel := *openapiclient.NewUpdateMachineAccountRequestModel() // UpdateMachineAccountRequestModel | New details for the machine account record.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalogMachineAccount(context.Background(), nameOrId, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMachineAccountRequestModel(updateMachineAccountRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalogMachineAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsUpdateMachineCatalogMachineAccount`: ProvisioningSchemeMachineAccountResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalogMachineAccount`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **updateMachineAccountRequestModel** | [**UpdateMachineAccountRequestModel**](UpdateMachineAccountRequestModel.md) | New details for the machine account record. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
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

> JobResponseModel MachineCatalogsUpdateMachineCatalogProvisioningScheme(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMachineCatalogProvisioningSchemeRequestModel(updateMachineCatalogProvisioningSchemeRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    updateMachineCatalogProvisioningSchemeRequestModel := *openapiclient.NewUpdateMachineCatalogProvisioningSchemeRequestModel() // UpdateMachineCatalogProvisioningSchemeRequestModel | Details about the image update to be performed.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the `async` parameter is `true`, the only difference is that instead of returning the initial job status in the payload, the response is instead `202` with the `Location` header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job's  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated, * _MasterImageXDPath_ - Path to the new image specified for the provisioning scheme, in XenApp & XenDesktop path format. * _MasterImageFullPath_ - Path to the new image specified for the provisioning scheme, relative to the root of the API. * _MasterImageRelativePath_ - Path to the new image specified for the provisioning scheme, relative to the resource pool. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalogProvisioningScheme(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateMachineCatalogProvisioningSchemeRequestModel(updateMachineCatalogProvisioningSchemeRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalogProvisioningScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsUpdateMachineCatalogProvisioningScheme`: JobResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalogProvisioningScheme`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **updateMachineCatalogProvisioningSchemeRequestModel** | [**UpdateMachineCatalogProvisioningSchemeRequestModel**](UpdateMachineCatalogProvisioningSchemeRequestModel.md) | Details about the image update to be performed. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | Due to the long-running nature of the update operation, it cannot complete synchronously, and always initiates an asynchronous job to run to completion.  If the &#x60;async&#x60; parameter is &#x60;true&#x60;, the only difference is that instead of returning the initial job status in the payload, the response is instead &#x60;202&#x60; with the &#x60;Location&#x60; header pointing to the job status API.  This may be desirable for clients that wish to unify their logic for handling asynchronous jobs.  The task will have JobType . When the task is complete it will redirect to . The job&#39;s  will contain properties:   * _Id_ - ID of the machine catalog for which the provisioning scheme is being updated, * _Name_ - Name of the machine catalog for which the provisioning scheme is being updated, * _MasterImageXDPath_ - Path to the new image specified for the provisioning scheme, in XenApp &amp; XenDesktop path format. * _MasterImageFullPath_ - Path to the new image specified for the provisioning scheme, relative to the root of the API. * _MasterImageRelativePath_ - Path to the new image specified for the provisioning scheme, relative to the resource pool. | [default to false]

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

> MachineCatalogsUpdateMachineCatalogUpgradeSchedule(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateUpgradeScheduleRequestModel(updateUpgradeScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog.             If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    updateUpgradeScheduleRequestModel := *openapiclient.NewUpdateUpgradeScheduleRequestModel(time.Now(), int32(123)) // UpdateUpgradeScheduleRequestModel | Details about the updated VDA upgrade schedule.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the VDA upgrade schedule will be updated as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalogUpgradeSchedule(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateUpgradeScheduleRequestModel(updateUpgradeScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsUpdateMachineCatalogUpgradeSchedule``: %v\n", err)
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


## MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection

> VDAComponentsSelectionValidationResponseModel MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).VDAComponentsSelectionValidationRequestModel(vDAComponentsSelectionValidationRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the machine catalog. If the catalog is present in a catalog folder,             specify the name in this format: {catalog folder path plus catalog name}.             For example, FolderName1|FolderName2|CatalogName.
    vDAComponentsSelectionValidationRequestModel := *openapiclient.NewVDAComponentsSelectionValidationRequestModel() // VDAComponentsSelectionValidationRequestModel | More details attached to this validation.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If async execute. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineCatalogsAPIsDAAS.MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).VDAComponentsSelectionValidationRequestModel(vDAComponentsSelectionValidationRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineCatalogsAPIsDAAS.MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection`: VDAComponentsSelectionValidationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineCatalogsAPIsDAAS.MachineCatalogsValidateMachineCatalogVDAComponentsAndFeaturesSelection`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **vDAComponentsSelectionValidationRequestModel** | [**VDAComponentsSelectionValidationRequestModel**](VDAComponentsSelectionValidationRequestModel.md) | More details attached to this validation. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If async execute. | [default to false]

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

