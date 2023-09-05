# \IdentityTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityTPAddAzureAdSecurityGroupMember**](IdentityTPApi.md#IdentityTPAddAzureAdSecurityGroupMember) | **Put** /techpreview/{customerid}/{siteid}/Identity/AzureADTenants/{azureAdTenantId}/AzureADSecurityGroups/{groupId}/members/{refGroupId} | Add a security group to an assigned security as it&#39;s member
[**IdentityTPCreateMachine**](IdentityTPApi.md#IdentityTPCreateMachine) | **Post** /techpreview/{customerid}/{siteid}/Identity/Machines | Create machine identity
[**IdentityTPDeleteMachine**](IdentityTPApi.md#IdentityTPDeleteMachine) | **Delete** /techpreview/{customerid}/{siteid}/Identity/Machines/{machine} | Delete a machine identity
[**IdentityTPGetAzureADSecurityGroups**](IdentityTPApi.md#IdentityTPGetAzureADSecurityGroups) | **Get** /techpreview/{customerid}/{siteid}/Identity/AzureADTenants/{azureAdTenantId}/AzureADSecurityGroups/{groupId} | Get Azure AD security group by group id.
[**IdentityTPGetAzureAdSecurityGroupMember**](IdentityTPApi.md#IdentityTPGetAzureAdSecurityGroupMember) | **Get** /techpreview/{customerid}/{siteid}/Identity/AzureADTenants/{azureAdTenantId}/AzureADSecurityGroups/{groupId}/members | Retrieves all the group type of members of a specific group
[**IdentityTPGetContainer**](IdentityTPApi.md#IdentityTPGetContainer) | **Get** /techpreview/{customerid}/{siteid}/Identity/Containers/{container} | Get a single container (e.g. OU)
[**IdentityTPGetContainers**](IdentityTPApi.md#IdentityTPGetContainers) | **Get** /techpreview/{customerid}/{siteid}/Identity/Containers | Get containers from identity provider.
[**IdentityTPGetDirectories**](IdentityTPApi.md#IdentityTPGetDirectories) | **Get** /techpreview/{customerid}/{siteid}/Identity/Directories | Get list of directories from all identity providers
[**IdentityTPGetDomain**](IdentityTPApi.md#IdentityTPGetDomain) | **Get** /techpreview/{customerid}/{siteid}/Identity/Domains/{domain} | Get a single domain from the identity provider
[**IdentityTPGetDomainAzureADCustomDomain**](IdentityTPApi.md#IdentityTPGetDomainAzureADCustomDomain) | **Get** /techpreview/{customerid}/{siteid}/Identity/Domains/{domain}/AzureADCustomDomain | Gets the Azure AD custom domain with the specified domain name.
[**IdentityTPGetDomains**](IdentityTPApi.md#IdentityTPGetDomains) | **Get** /techpreview/{customerid}/{siteid}/Identity/Domains | Get list of domains from the identity provider
[**IdentityTPGetForest**](IdentityTPApi.md#IdentityTPGetForest) | **Get** /techpreview/{customerid}/{siteid}/Identity/Forests/{forest} | Get information about a single forest
[**IdentityTPGetForests**](IdentityTPApi.md#IdentityTPGetForests) | **Get** /techpreview/{customerid}/{siteid}/Identity/Forests | Get the list of forests from the identity provider
[**IdentityTPGetMachine**](IdentityTPApi.md#IdentityTPGetMachine) | **Get** /techpreview/{customerid}/{siteid}/Identity/Machines/{machine} | Get a single machine account from identity provider
[**IdentityTPGetMachines**](IdentityTPApi.md#IdentityTPGetMachines) | **Get** /techpreview/{customerid}/{siteid}/Identity/Machines | Get machine accounts from identity provider
[**IdentityTPGetSite**](IdentityTPApi.md#IdentityTPGetSite) | **Get** /techpreview/{customerid}/{siteid}/Identity/Sites/{site} | Get information about a single site
[**IdentityTPGetSites**](IdentityTPApi.md#IdentityTPGetSites) | **Get** /techpreview/{customerid}/{siteid}/Identity/Sites | Get list of sites from the identity provider
[**IdentityTPGetUser**](IdentityTPApi.md#IdentityTPGetUser) | **Get** /techpreview/{customerid}/{siteid}/Identity/Users/{userOrGroup} | Get a single user or group
[**IdentityTPGetUsers**](IdentityTPApi.md#IdentityTPGetUsers) | **Get** /techpreview/{customerid}/{siteid}/Identity/Users | Get users from identity provider
[**IdentityTPQueryAzureAdSecurityGroups**](IdentityTPApi.md#IdentityTPQueryAzureAdSecurityGroups) | **Get** /techpreview/{customerid}/{siteid}/Identity/AzureADTenants/{azureAdTenantId}/AzureADSecurityGroups | Query AzureAD security group by user&#39;s input.
[**IdentityTPUpdateMachine**](IdentityTPApi.md#IdentityTPUpdateMachine) | **Patch** /techpreview/{customerid}/{siteid}/Identity/Machines/{machine} | Update a machine identity
[**IdentityTPUpdateUser**](IdentityTPApi.md#IdentityTPUpdateUser) | **Patch** /techpreview/{customerid}/{siteid}/Identity/Users/{user} | Update a user identity
[**IdentityTPValidateUser**](IdentityTPApi.md#IdentityTPValidateUser) | **Post** /techpreview/{customerid}/{siteid}/Identity/Users/$validate | Validate a user&#39;s credentials



## IdentityTPAddAzureAdSecurityGroupMember

> IdentityTPAddAzureAdSecurityGroupMember(ctx, azureAdTenantId, groupId, refGroupId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()

Add a security group to an assigned security as it's member

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    azureAdTenantId := "azureAdTenantId_example" // string | The Azure tenant id.
    groupId := "groupId_example" // string | The parent assigned security group object id
    refGroupId := "refGroupId_example" // string | The nested devices security group object id
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityTPApi.IdentityTPAddAzureAdSecurityGroupMember(context.Background(), azureAdTenantId, groupId, refGroupId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPAddAzureAdSecurityGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAdTenantId** | **string** | The Azure tenant id. | 
**groupId** | **string** | The parent assigned security group object id | 
**refGroupId** | **string** | The nested devices security group object id | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPAddAzureAdSecurityGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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


## IdentityTPCreateMachine

> IdentityMachineResponseModel IdentityTPCreateMachine(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Create machine identity



### Example

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
    request := *openapiclient.NewIdentityCreateMachineRequestModel("Forest_example", "Name_example", "AccountPassword_example") // IdentityCreateMachineRequestModel | Details of the machine identity to create.
    async := true // bool | If `true`, the machine identity will be created as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPCreateMachine(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPCreateMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPCreateMachine`: IdentityMachineResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPCreateMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPCreateMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**IdentityCreateMachineRequestModel**](IdentityCreateMachineRequestModel.md) | Details of the machine identity to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine identity will be created as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityMachineResponseModel**](IdentityMachineResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPDeleteMachine

> IdentityTPDeleteMachine(ctx, machine, customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Delete a machine identity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    machine := "machine_example" // string | Machine identity to delete.
    forest := "forest_example" // string | Forest to delete machine account from.  Required.
    domain := "domain_example" // string | Domain to delete machine account from.  Required.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    async := true // bool | If `true`, the machine identity will be deleted as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityTPApi.IdentityTPDeleteMachine(context.Background(), machine, customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPDeleteMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine identity to delete. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPDeleteMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forest** | **string** | Forest to delete machine account from.  Required. | 
 **domain** | **string** | Domain to delete machine account from.  Required. | 


 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine identity will be deleted as a background task. | [default to false]
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


## IdentityTPGetAzureADSecurityGroups

> AzureADSecurityGroupResponseModel IdentityTPGetAzureADSecurityGroups(ctx, azureAdTenantId, groupId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()

Get Azure AD security group by group id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    azureAdTenantId := "azureAdTenantId_example" // string | AzureAD tenantId
    groupId := "groupId_example" // string | AzureAD security group's objectId
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetAzureADSecurityGroups(context.Background(), azureAdTenantId, groupId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetAzureADSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetAzureADSecurityGroups`: AzureADSecurityGroupResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetAzureADSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAdTenantId** | **string** | AzureAD tenantId | 
**groupId** | **string** | AzureAD security group&#39;s objectId | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetAzureADSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AzureADSecurityGroupResponseModel**](AzureADSecurityGroupResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetAzureAdSecurityGroupMember

> AzureADSecurityGroupResponseModelCollection IdentityTPGetAzureAdSecurityGroupMember(ctx, azureAdTenantId, groupId, customerid, siteid).Type_(type_).MaxCount(maxCount).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()

Retrieves all the group type of members of a specific group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    azureAdTenantId := "azureAdTenantId_example" // string | The Azure tenant id.
    groupId := "groupId_example" // string | The security group object id.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    type_ := "type__example" // string | Only type=group is support now. (optional) (default to "group")
    maxCount := int32(56) // int32 | The max return records number. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetAzureAdSecurityGroupMember(context.Background(), azureAdTenantId, groupId, customerid, siteid).Type_(type_).MaxCount(maxCount).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetAzureAdSecurityGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetAzureAdSecurityGroupMember`: AzureADSecurityGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetAzureAdSecurityGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAdTenantId** | **string** | The Azure tenant id. | 
**groupId** | **string** | The security group object id. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetAzureAdSecurityGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **type_** | **string** | Only type&#x3D;group is support now. | [default to &quot;group&quot;]
 **maxCount** | **int32** | The max return records number. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AzureADSecurityGroupResponseModelCollection**](AzureADSecurityGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetContainer

> IdentityContainerResponseModel IdentityTPGetContainer(ctx, container, customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get a single container (e.g. OU)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    container := "container_example" // string | Container to get.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    forest := "forest_example" // string | Forest to get container from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get container from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 4294967295)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetContainer(context.Background(), container, customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetContainer`: IdentityContainerResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**container** | **string** | Container to get. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forest** | **string** | Forest to get container from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get container from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 4294967295]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityContainerResponseModel**](IdentityContainerResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetContainers

> IdentityContainerResponseModelCollection IdentityTPGetContainers(ctx, customerid, siteid).Forest(forest).Domain(domain).Container(container).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).ContainerType(containerType).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get containers from identity provider.



### Example

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
    forest := "forest_example" // string | Forest to get containers from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get containers from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    container := []string{"Inner_example"} // []string | Specific container(s) to filter the results to.  If not specified, all matching containers are returned.  If set, `recursive` parameter is ignored and is implied `true`. (optional)
    startsWith := "startsWith_example" // string | Search for containers that start with a string.  If not specified, all matching containers are returned. (optional)
    contains := "contains_example" // string | Search for containers that contain a string.  If not specified, all matching containers are returned. (optional)
    matches := "matches_example" // string | Search for containers that match a string.  If not specified, all matching containers are returned. (optional)
    parentOU := "parentOU_example" // string | The parent OU to search.  If not specified, will search from the root OU. (optional)
    recursive := true // bool | Indicates whether the search should be recursive.  Default is `false`. (optional)
    containerType := []openapiclient.IdentityContainerType{openapiclient.IdentityContainerType("Unknown")} // []IdentityContainerType | Indicates container type(s) that should be retrieved.  If not specified, all container types will be searched. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 4294967295)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetContainers(context.Background(), customerid, siteid).Forest(forest).Domain(domain).Container(container).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).ContainerType(containerType).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetContainers`: IdentityContainerResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forest** | **string** | Forest to get containers from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get containers from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **container** | **[]string** | Specific container(s) to filter the results to.  If not specified, all matching containers are returned.  If set, &#x60;recursive&#x60; parameter is ignored and is implied &#x60;true&#x60;. | 
 **startsWith** | **string** | Search for containers that start with a string.  If not specified, all matching containers are returned. | 
 **contains** | **string** | Search for containers that contain a string.  If not specified, all matching containers are returned. | 
 **matches** | **string** | Search for containers that match a string.  If not specified, all matching containers are returned. | 
 **parentOU** | **string** | The parent OU to search.  If not specified, will search from the root OU. | 
 **recursive** | **bool** | Indicates whether the search should be recursive.  Default is &#x60;false&#x60;. | 
 **containerType** | [**[]IdentityContainerType**](IdentityContainerType.md) | Indicates container type(s) that should be retrieved.  If not specified, all container types will be searched. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 4294967295]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityContainerResponseModelCollection**](IdentityContainerResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetDirectories

> IdentityDirectoryResponseModelCollection IdentityTPGetDirectories(ctx, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get list of directories from all identity providers



### Example

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
    limit := int32(56) // int32 | The max number of items returned by this query. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetDirectories(context.Background(), customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetDirectories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetDirectories`: IdentityDirectoryResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetDirectories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items returned by this query. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityDirectoryResponseModelCollection**](IdentityDirectoryResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetDomain

> IdentityDomainResponseModel IdentityTPGetDomain(ctx, domain, customerid, siteid).Forest(forest).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get a single domain from the identity provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    domain := "domain_example" // string | Domain to get details from.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    forest := "forest_example" // string | Forest to get domain from.  If not specified, all forests are queried, which may take a long time. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties. If not specified, all properties other than IdentityDomainProperty.Controllers IdentityDomainProperty.PrimaryController will be retrieved. (optional) (default to 4294966527)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetDomain(context.Background(), domain, customerid, siteid).Forest(forest).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetDomain`: IdentityDomainResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain to get details from. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forest** | **string** | Forest to get domain from.  If not specified, all forests are queried, which may take a long time. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties. If not specified, all properties other than IdentityDomainProperty.Controllers IdentityDomainProperty.PrimaryController will be retrieved. | [default to 4294966527]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityDomainResponseModel**](IdentityDomainResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetDomainAzureADCustomDomain

> AzureADCustomDomainResponseModel IdentityTPGetDomainAzureADCustomDomain(ctx, domain, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Gets the Azure AD custom domain with the specified domain name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    domain := "domain_example" // string | Domain name.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetDomainAzureADCustomDomain(context.Background(), domain, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetDomainAzureADCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetDomainAzureADCustomDomain`: AzureADCustomDomainResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetDomainAzureADCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetDomainAzureADCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AzureADCustomDomainResponseModel**](AzureADCustomDomainResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetDomains

> IdentityDomainResponseModelCollection IdentityTPGetDomains(ctx, customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).ParentDomain(parentDomain).Recursive(recursive).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get list of domains from the identity provider



### Example

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
    forest := "forest_example" // string | Forest to get domains from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := []string{"Inner_example"} // []string | Specific domain(s) to filter the results to.  If not specified, all matching domains are returned. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    parentDomain := "parentDomain_example" // string | Parent domain name to search.  Default is the root domain of the forest. (optional)
    recursive := true // bool | Specifies whether the search is recursive. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties. If not specified, all properties other than IdentityDomainProperty.Controllers IdentityDomainProperty.PrimaryController will be retrieved. (optional) (default to 4294966527)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)
    async := true // bool |  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetDomains(context.Background(), customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).ParentDomain(parentDomain).Recursive(recursive).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetDomains`: IdentityDomainResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forest** | **string** | Forest to get domains from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **[]string** | Specific domain(s) to filter the results to.  If not specified, all matching domains are returned. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **parentDomain** | **string** | Parent domain name to search.  Default is the root domain of the forest. | 
 **recursive** | **bool** | Specifies whether the search is recursive. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties. If not specified, all properties other than IdentityDomainProperty.Controllers IdentityDomainProperty.PrimaryController will be retrieved. | [default to 4294966527]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 
 **async** | **bool** |  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityDomainResponseModelCollection**](IdentityDomainResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetForest

> IdentityForestResponseModel IdentityTPGetForest(ctx, forest, customerid, siteid).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get information about a single forest



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    forest := "forest_example" // string | forest
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 4294967295)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetForest(context.Background(), forest, customerid, siteid).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetForest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetForest`: IdentityForestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetForest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forest** | **string** | forest | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetForestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 4294967295]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityForestResponseModel**](IdentityForestResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetForests

> IdentityForestResponseModelCollection IdentityTPGetForests(ctx, customerid, siteid).Forest(forest).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get the list of forests from the identity provider



### Example

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
    forest := []string{"Inner_example"} // []string | Specific forest(s) to filter the results to.  If not specified, all forests are returned. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 4294967295)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetForests(context.Background(), customerid, siteid).Forest(forest).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetForests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetForests`: IdentityForestResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetForests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetForestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forest** | **[]string** | Specific forest(s) to filter the results to.  If not specified, all forests are returned. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 4294967295]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityForestResponseModelCollection**](IdentityForestResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetMachine

> IdentityMachineResponseModel IdentityTPGetMachine(ctx, machine, customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get a single machine account from identity provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    machine := "machine_example" // string | Machine to get.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    forest := "forest_example" // string | Forest to get machine account from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get machine account from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties other than IPAddress will be retrieved. (optional) (default to 4294967279)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetMachine(context.Background(), machine, customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetMachine`: IdentityMachineResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine to get. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forest** | **string** | Forest to get machine account from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get machine account from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties other than IPAddress will be retrieved. | [default to 4294967279]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityMachineResponseModel**](IdentityMachineResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetMachines

> IdentityMachineResponseModelCollection IdentityTPGetMachines(ctx, customerid, siteid).Forest(forest).Domain(domain).Machine(machine).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get machine accounts from identity provider



### Example

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
    forest := "forest_example" // string | Forest to get machine accounts from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get machine accounts from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    machine := []string{"Inner_example"} // []string | Specific machine(s) to filter the results to.  If not specified, all matching machines are returned.  If set, `recursive` parameter is ignored and is implied `true`. (optional)
    startsWith := "startsWith_example" // string | Search for machine accounts that start with a string.  This parameter is exclusive with `contains`, and `matches`. (optional)
    contains := "contains_example" // string | Search for machine accounts that contain a string.  This parameter is exclusive with `startsWith`, and `matches`. (optional)
    matches := "matches_example" // string | Search for machine accounts that match a string.  This parameter is exclusive with `startsWith`, and `contains`. (optional)
    parentOU := "parentOU_example" // string | The parent OU to search.  If not specified, will search from the root OU. (optional)
    recursive := true // bool | Indicates whether the search should be recursive. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties other than IPAddress will be retrieved. (optional) (default to 4294967279)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetMachines(context.Background(), customerid, siteid).Forest(forest).Domain(domain).Machine(machine).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetMachines`: IdentityMachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forest** | **string** | Forest to get machine accounts from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get machine accounts from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **machine** | **[]string** | Specific machine(s) to filter the results to.  If not specified, all matching machines are returned.  If set, &#x60;recursive&#x60; parameter is ignored and is implied &#x60;true&#x60;. | 
 **startsWith** | **string** | Search for machine accounts that start with a string.  This parameter is exclusive with &#x60;contains&#x60;, and &#x60;matches&#x60;. | 
 **contains** | **string** | Search for machine accounts that contain a string.  This parameter is exclusive with &#x60;startsWith&#x60;, and &#x60;matches&#x60;. | 
 **matches** | **string** | Search for machine accounts that match a string.  This parameter is exclusive with &#x60;startsWith&#x60;, and &#x60;contains&#x60;. | 
 **parentOU** | **string** | The parent OU to search.  If not specified, will search from the root OU. | 
 **recursive** | **bool** | Indicates whether the search should be recursive. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties other than IPAddress will be retrieved. | [default to 4294967279]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityMachineResponseModelCollection**](IdentityMachineResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetSite

> IdentitySiteResponseModel IdentityTPGetSite(ctx, site, customerid, siteid).Forest(forest).Domain(domain).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get information about a single site



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    site := "site_example" // string | The site to get details for.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    forest := "forest_example" // string | Forest to get the site from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get the site from.  If not specified, all domains in the forest(s) are queried until the site is found. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 4294967295)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetSite(context.Background(), site, customerid, siteid).Forest(forest).Domain(domain).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetSite`: IdentitySiteResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string** | The site to get details for. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forest** | **string** | Forest to get the site from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get the site from.  If not specified, all domains in the forest(s) are queried until the site is found. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 4294967295]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentitySiteResponseModel**](IdentitySiteResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetSites

> IdentitySiteResponseModelCollection IdentityTPGetSites(ctx, customerid, siteid).Forest(forest).Domain(domain).HostNameOrIP(hostNameOrIP).Site(site).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get list of sites from the identity provider



### Example

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
    forest := "forest_example" // string | Forest to get sites from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get sites from.  If not specified, all domains in the forest(s) are queried. (optional)
    hostNameOrIP := "hostNameOrIP_example" // string | Hostname or IP to get sites from. (optional)
    site := []string{"Inner_example"} // []string | Specific site(s) to filter the results to.  If not specified, all matching sites are returned. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 4294967295)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetSites(context.Background(), customerid, siteid).Forest(forest).Domain(domain).HostNameOrIP(hostNameOrIP).Site(site).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetSites`: IdentitySiteResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forest** | **string** | Forest to get sites from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get sites from.  If not specified, all domains in the forest(s) are queried. | 
 **hostNameOrIP** | **string** | Hostname or IP to get sites from. | 
 **site** | **[]string** | Specific site(s) to filter the results to.  If not specified, all matching sites are returned. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 4294967295]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentitySiteResponseModelCollection**](IdentitySiteResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetUser

> IdentityUserResponseModel IdentityTPGetUser(ctx, userOrGroup, customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get a single user or group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    userOrGroup := "userOrGroup_example" // string | Identity of the user or group to get.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    forest := "forest_example" // string | Forest to get users from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get users from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 4294967295)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetUser(context.Background(), userOrGroup, customerid, siteid).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetUser`: IdentityUserResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userOrGroup** | **string** | Identity of the user or group to get. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forest** | **string** | Forest to get users from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get users from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 4294967295]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityUserResponseModel**](IdentityUserResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPGetUsers

> IdentityUserResponseModelCollection IdentityTPGetUsers(ctx, customerid, siteid).Provider(provider).Forest(forest).Domain(domain).Tenant(tenant).User(user).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).UserType(userType).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get users from identity provider



### Example

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
    provider := "provider_example" // string | Provider to get users from. (optional)
    forest := "forest_example" // string | Forest to get users from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get users from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    tenant := "tenant_example" // string | Tenant to get users from. (optional)
    user := []string{"Inner_example"} // []string | Specific user(s) to filter the results to. If not specified, all matching users are returned.  If set, `recursive` parameter is ignored and is implied `true`. (optional)
    startsWith := "startsWith_example" // string | Search for users that start with a string.  This parameter is exclusive with `contains`, and `matches`. (optional)
    contains := "contains_example" // string | Search for users that contain a string.  This parameter is exclusive with `startsWith`, and `matches`. (optional)
    matches := "matches_example" // string | Search for users that match a string.  This parameter is exclusive with `startsWith`, and `contains`. (optional)
    parentOU := "parentOU_example" // string | The parent OU to search.  If not specified, will search from the root OU. (optional)
    recursive := true // bool | Indicates whether the search should be recursive. (optional)
    userType := "userType_example" // string | Indicates user type(s) that should be retrieved.  If not specified, all user types will be searched. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 4294967295)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPGetUsers(context.Background(), customerid, siteid).Provider(provider).Forest(forest).Domain(domain).Tenant(tenant).User(user).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).UserType(userType).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPGetUsers`: IdentityUserResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPGetUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provider** | **string** | Provider to get users from. | 
 **forest** | **string** | Forest to get users from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get users from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **tenant** | **string** | Tenant to get users from. | 
 **user** | **[]string** | Specific user(s) to filter the results to. If not specified, all matching users are returned.  If set, &#x60;recursive&#x60; parameter is ignored and is implied &#x60;true&#x60;. | 
 **startsWith** | **string** | Search for users that start with a string.  This parameter is exclusive with &#x60;contains&#x60;, and &#x60;matches&#x60;. | 
 **contains** | **string** | Search for users that contain a string.  This parameter is exclusive with &#x60;startsWith&#x60;, and &#x60;matches&#x60;. | 
 **matches** | **string** | Search for users that match a string.  This parameter is exclusive with &#x60;startsWith&#x60;, and &#x60;contains&#x60;. | 
 **parentOU** | **string** | The parent OU to search.  If not specified, will search from the root OU. | 
 **recursive** | **bool** | Indicates whether the search should be recursive. | 
 **userType** | **string** | Indicates user type(s) that should be retrieved.  If not specified, all user types will be searched. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 4294967295]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityUserResponseModelCollection**](IdentityUserResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPQueryAzureAdSecurityGroups

> AzureADSecurityGroupResponseModelCollection IdentityTPQueryAzureAdSecurityGroups(ctx, azureAdTenantId, customerid, siteid).Name(name).Assigned(assigned).MaxCount(maxCount).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()

Query AzureAD security group by user's input.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    azureAdTenantId := "azureAdTenantId_example" // string | The specific azure tenant id.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    name := "name_example" // string | Specific the group display name. (optional)
    assigned := true // bool | When name is empty, assigned only support the value of true. (optional) (default to true)
    maxCount := int32(56) // int32 | The max return count,default is 300. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPQueryAzureAdSecurityGroups(context.Background(), azureAdTenantId, customerid, siteid).Name(name).Assigned(assigned).MaxCount(maxCount).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPQueryAzureAdSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPQueryAzureAdSecurityGroups`: AzureADSecurityGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPQueryAzureAdSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAdTenantId** | **string** | The specific azure tenant id. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPQueryAzureAdSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **name** | **string** | Specific the group display name. | 
 **assigned** | **bool** | When name is empty, assigned only support the value of true. | [default to true]
 **maxCount** | **int32** | The max return count,default is 300. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AzureADSecurityGroupResponseModelCollection**](AzureADSecurityGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPUpdateMachine

> IdentityMachineResponseModel IdentityTPUpdateMachine(ctx, machine, customerid, siteid).Forest(forest).Domain(domain).Request(request).DirectoryServerHint(directoryServerHint).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Update a machine identity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    machine := "machine_example" // string | Machine identity to update.
    forest := "forest_example" // string | Forest to update machine account within.  Required.
    domain := "domain_example" // string | Domain to update machine account within.  Required.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewIdentityUpdateMachineRequestModel() // IdentityUpdateMachineRequestModel | Properties of the machine identity that should be updated.
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    async := true // bool | If `true`, the machine identity will be updated as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPUpdateMachine(context.Background(), machine, customerid, siteid).Forest(forest).Domain(domain).Request(request).DirectoryServerHint(directoryServerHint).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPUpdateMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPUpdateMachine`: IdentityMachineResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPUpdateMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine identity to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPUpdateMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forest** | **string** | Forest to update machine account within.  Required. | 
 **domain** | **string** | Domain to update machine account within.  Required. | 


 **request** | [**IdentityUpdateMachineRequestModel**](IdentityUpdateMachineRequestModel.md) | Properties of the machine identity that should be updated. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine identity will be updated as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityMachineResponseModel**](IdentityMachineResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPUpdateUser

> IdentityUserResponseModel IdentityTPUpdateUser(ctx, user, customerid, siteid).Forest(forest).Domain(domain).Request(request).DirectoryServerHint(directoryServerHint).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Update a user identity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    user := "user_example" // string | User identity to update.
    forest := "forest_example" // string | Forest to update user account within.  Required.
    domain := "domain_example" // string | Domain to update user account within.  Required.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewIdentityUpdateUserRequestModel() // IdentityUpdateUserRequestModel | Properties of the user identity that should be updated.
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    async := true // bool | If `true`, the user identity will be updated as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPUpdateUser(context.Background(), user, customerid, siteid).Forest(forest).Domain(domain).Request(request).DirectoryServerHint(directoryServerHint).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPUpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPUpdateUser`: IdentityUserResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPUpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | User identity to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forest** | **string** | Forest to update user account within.  Required. | 
 **domain** | **string** | Domain to update user account within.  Required. | 


 **request** | [**IdentityUpdateUserRequestModel**](IdentityUpdateUserRequestModel.md) | Properties of the user identity that should be updated. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **async** | **bool** | If &#x60;true&#x60;, the user identity will be updated as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityUserResponseModel**](IdentityUserResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityTPValidateUser

> IdentityUserResponseModel IdentityTPValidateUser(ctx, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Validate a user's credentials



### Example

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
    request := *openapiclient.NewIdentityValidateUserRequestModel("UserName_example", "Password_example") // IdentityValidateUserRequestModel | Details of the user credentials to validate.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityTPApi.IdentityTPValidateUser(context.Background(), customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityTPApi.IdentityTPValidateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityTPValidateUser`: IdentityUserResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityTPApi.IdentityTPValidateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityTPValidateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**IdentityValidateUserRequestModel**](IdentityValidateUserRequestModel.md) | Details of the user credentials to validate. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IdentityUserResponseModel**](IdentityUserResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

