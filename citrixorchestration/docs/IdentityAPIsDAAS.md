# \IdentityAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityAddAzureAdSecurityGroupMember**](IdentityAPIsDAAS.md#IdentityAddAzureAdSecurityGroupMember) | **Put** /Identity/AzureADTenants/{azureAdTenantId}/AzureADSecurityGroups/{groupId}/members/{refGroupId} | Add a security group to an assigned security as it&#39;s member
[**IdentityCreateMachine**](IdentityAPIsDAAS.md#IdentityCreateMachine) | **Post** /Identity/Machines | Create machine identity
[**IdentityCreateServiceAccount**](IdentityAPIsDAAS.md#IdentityCreateServiceAccount) | **Post** /Identity/ServiceAccount | Create a service account.
[**IdentityDeleteMachine**](IdentityAPIsDAAS.md#IdentityDeleteMachine) | **Delete** /Identity/Machines/{machine} | Delete a machine identity
[**IdentityDeleteServiceAccount**](IdentityAPIsDAAS.md#IdentityDeleteServiceAccount) | **Delete** /Identity/ServiceAccount/{serviceAccountUid} | Delete an existing service account.
[**IdentityGetAzureADSecurityGroups**](IdentityAPIsDAAS.md#IdentityGetAzureADSecurityGroups) | **Get** /Identity/AzureADTenants/{azureAdTenantId}/AzureADSecurityGroups/{groupId} | Get Azure AD security group by group id.
[**IdentityGetAzureAdSecurityGroupMember**](IdentityAPIsDAAS.md#IdentityGetAzureAdSecurityGroupMember) | **Get** /Identity/AzureADTenants/{azureAdTenantId}/AzureADSecurityGroups/{groupId}/members | Retrieves all the group type of members of a specific group
[**IdentityGetContainer**](IdentityAPIsDAAS.md#IdentityGetContainer) | **Get** /Identity/Containers/{container} | Get a single container (e.g. OU)
[**IdentityGetContainers**](IdentityAPIsDAAS.md#IdentityGetContainers) | **Get** /Identity/Containers | Get containers from identity provider.
[**IdentityGetDirectories**](IdentityAPIsDAAS.md#IdentityGetDirectories) | **Get** /Identity/Directories | Get list of directories from all identity providers
[**IdentityGetDomain**](IdentityAPIsDAAS.md#IdentityGetDomain) | **Get** /Identity/Domains/{domain} | Get a single domain from the identity provider
[**IdentityGetDomainAzureADCustomDomain**](IdentityAPIsDAAS.md#IdentityGetDomainAzureADCustomDomain) | **Get** /Identity/Domains/{domain}/AzureADCustomDomain | Gets the Azure AD custom domain with the specified domain name.
[**IdentityGetDomains**](IdentityAPIsDAAS.md#IdentityGetDomains) | **Get** /Identity/Domains | Get list of domains from the identity provider
[**IdentityGetForest**](IdentityAPIsDAAS.md#IdentityGetForest) | **Get** /Identity/Forests/{forest} | Get information about a single forest
[**IdentityGetForests**](IdentityAPIsDAAS.md#IdentityGetForests) | **Get** /Identity/Forests | Get the list of forests from the identity provider
[**IdentityGetMachine**](IdentityAPIsDAAS.md#IdentityGetMachine) | **Get** /Identity/Machines/{machine} | Get a single machine account from identity provider
[**IdentityGetMachines**](IdentityAPIsDAAS.md#IdentityGetMachines) | **Get** /Identity/Machines | Get machine accounts from identity provider
[**IdentityGetServiceAccount**](IdentityAPIsDAAS.md#IdentityGetServiceAccount) | **Get** /Identity/ServiceAccount/{serviceAccountUid} | Get a specific service account.
[**IdentityGetServiceAccounts**](IdentityAPIsDAAS.md#IdentityGetServiceAccounts) | **Get** /Identity/ServiceAccounts | Get all service accounts.
[**IdentityGetSite**](IdentityAPIsDAAS.md#IdentityGetSite) | **Get** /Identity/Sites/{site} | Get information about a single site
[**IdentityGetSites**](IdentityAPIsDAAS.md#IdentityGetSites) | **Get** /Identity/Sites | Get list of sites from the identity provider
[**IdentityGetUser**](IdentityAPIsDAAS.md#IdentityGetUser) | **Get** /Identity/Users/{userOrGroup} | Get a single user or group
[**IdentityGetUsers**](IdentityAPIsDAAS.md#IdentityGetUsers) | **Get** /Identity/Users | Get users from identity provider
[**IdentityQueryAzureAdSecurityGroups**](IdentityAPIsDAAS.md#IdentityQueryAzureAdSecurityGroups) | **Get** /Identity/AzureADTenants/{azureAdTenantId}/AzureADSecurityGroups | Query AzureAD security group by user&#39;s input.
[**IdentitySetServiceAccount**](IdentityAPIsDAAS.md#IdentitySetServiceAccount) | **Patch** /Identity/ServiceAccount/{serviceAccountUid} | Set properties of an existing service account.
[**IdentityUpdateMachine**](IdentityAPIsDAAS.md#IdentityUpdateMachine) | **Patch** /Identity/Machines/{machine} | Update a machine identity
[**IdentityUpdateUser**](IdentityAPIsDAAS.md#IdentityUpdateUser) | **Patch** /Identity/Users/{user} | Update a user identity
[**IdentityValidateUser**](IdentityAPIsDAAS.md#IdentityValidateUser) | **Post** /Identity/Users/$validate | Validate a user&#39;s credentials



## IdentityAddAzureAdSecurityGroupMember

> IdentityAddAzureAdSecurityGroupMember(ctx, azureAdTenantId, groupId, refGroupId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    azureAdTenantId := "azureAdTenantId_example" // string | The Azure tenant id.
    groupId := "groupId_example" // string | The parent assigned security group object id
    refGroupId := "refGroupId_example" // string | The nested devices security group object id
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityAPIsDAAS.IdentityAddAzureAdSecurityGroupMember(context.Background(), azureAdTenantId, groupId, refGroupId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityAddAzureAdSecurityGroupMember``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityAddAzureAdSecurityGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 



 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
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


## IdentityCreateMachine

> IdentityMachineResponseModel IdentityCreateMachine(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).IdentityCreateMachineRequestModel(identityCreateMachineRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    identityCreateMachineRequestModel := *openapiclient.NewIdentityCreateMachineRequestModel("For example, test.local", "Name of the machine identity", "For example, password") // IdentityCreateMachineRequestModel | Details of the machine identity to create.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the machine identity will be created as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityCreateMachine(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).IdentityCreateMachineRequestModel(identityCreateMachineRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityCreateMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityCreateMachine`: IdentityMachineResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityCreateMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityCreateMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **identityCreateMachineRequestModel** | [**IdentityCreateMachineRequestModel**](IdentityCreateMachineRequestModel.md) | Details of the machine identity to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine identity will be created as a background task. | [default to false]

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


## IdentityCreateServiceAccount

> ServiceAccountResponseModel IdentityCreateServiceAccount(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateServiceAccountRequestModel(createServiceAccountRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create a service account.

### Example

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
    createServiceAccountRequestModel := *openapiclient.NewCreateServiceAccountRequestModel("Valid types: AzureAD, ActiveDirectory", "Domain name if 'IdentityProviderType' is ActiveDirectory. Azure AD tenant id if 'IdentityProviderType' is AzureAD", "Domain user like "cvad\\joe" if 'IdentityProviderType' is ActiveDirectory"", "SecretExpiryTime_example", "AccountSecret_example") // CreateServiceAccountRequestModel | The object of CreateServiceAccountRequestModel used to create a service account. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityCreateServiceAccount(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateServiceAccountRequestModel(createServiceAccountRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityCreateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityCreateServiceAccount`: ServiceAccountResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityCreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createServiceAccountRequestModel** | [**CreateServiceAccountRequestModel**](CreateServiceAccountRequestModel.md) | The object of CreateServiceAccountRequestModel used to create a service account.  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**ServiceAccountResponseModel**](ServiceAccountResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityDeleteMachine

> IdentityDeleteMachine(ctx, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).Forest(forest).Domain(domain).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DirectoryServerHint(directoryServerHint).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    machine := "machine_example" // string | Machine identity to delete.
    forest := "forest_example" // string | Forest to delete machine account from.  Required.
    domain := "domain_example" // string | Domain to delete machine account from.  Required.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    async := true // bool | If `true`, the machine identity will be deleted as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityAPIsDAAS.IdentityDeleteMachine(context.Background(), machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).Forest(forest).Domain(domain).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DirectoryServerHint(directoryServerHint).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityDeleteMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine identity to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityDeleteMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **forest** | **string** | Forest to delete machine account from.  Required. | 
 **domain** | **string** | Domain to delete machine account from.  Required. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine identity will be deleted as a background task. | [default to false]

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


## IdentityDeleteServiceAccount

> IdentityDeleteServiceAccount(ctx, serviceAccountUid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete an existing service account.

### Example

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
    serviceAccountUid := "serviceAccountUid_example" // string | The ServiceAccountUid of an existing service account. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityAPIsDAAS.IdentityDeleteServiceAccount(context.Background(), serviceAccountUid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityDeleteServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountUid** | **string** | The ServiceAccountUid of an existing service account.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityDeleteServiceAccountRequest struct via the builder pattern


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


## IdentityGetAzureADSecurityGroups

> AzureADSecurityGroupResponseModel IdentityGetAzureADSecurityGroups(ctx, azureAdTenantId, groupId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    azureAdTenantId := "azureAdTenantId_example" // string | AzureAD tenantId
    groupId := "groupId_example" // string | AzureAD security group's objectId
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetAzureADSecurityGroups(context.Background(), azureAdTenantId, groupId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetAzureADSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetAzureADSecurityGroups`: AzureADSecurityGroupResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetAzureADSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAdTenantId** | **string** | AzureAD tenantId | 
**groupId** | **string** | AzureAD security group&#39;s objectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetAzureADSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## IdentityGetAzureAdSecurityGroupMember

> AzureADSecurityGroupResponseModelCollection IdentityGetAzureAdSecurityGroupMember(ctx, azureAdTenantId, groupId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Type_(type_).MaxCount(maxCount).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    azureAdTenantId := "azureAdTenantId_example" // string | The Azure tenant id.
    groupId := "groupId_example" // string | The security group object id.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    type_ := "type__example" // string | Only type=group is support now. (optional) (default to "group")
    maxCount := int32(56) // int32 | The max return records number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetAzureAdSecurityGroupMember(context.Background(), azureAdTenantId, groupId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Type_(type_).MaxCount(maxCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetAzureAdSecurityGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetAzureAdSecurityGroupMember`: AzureADSecurityGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetAzureAdSecurityGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAdTenantId** | **string** | The Azure tenant id. | 
**groupId** | **string** | The security group object id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetAzureAdSecurityGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **type_** | **string** | Only type&#x3D;group is support now. | [default to &quot;group&quot;]
 **maxCount** | **int32** | The max return records number. | 

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


## IdentityGetContainer

> IdentityContainerResponseModel IdentityGetContainer(ctx, container).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    container := "container_example" // string | Container to get.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    forest := "forest_example" // string | Forest to get container from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get container from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 2147483647)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetContainer(context.Background(), container).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetContainer`: IdentityContainerResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**container** | **string** | Container to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetContainerRequest struct via the builder pattern


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
 **forest** | **string** | Forest to get container from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get container from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 2147483647]

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


## IdentityGetContainers

> IdentityContainerResponseModelCollection IdentityGetContainers(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).Container(container).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).ContainerType(containerType).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
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
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 2147483647)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetContainers(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).Container(container).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).ContainerType(containerType).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetContainers`: IdentityContainerResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetContainersRequest struct via the builder pattern


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
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 2147483647]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 

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


## IdentityGetDirectories

> IdentityDirectoryResponseModelCollection IdentityGetDirectories(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    limit := int32(56) // int32 | The max number of items returned by this query. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetDirectories(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetDirectories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetDirectories`: IdentityDirectoryResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetDirectories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of items returned by this query. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 

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


## IdentityGetDomain

> IdentityDomainResponseModel IdentityGetDomain(ctx, domain).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    domain := "domain_example" // string | Domain to get details from.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    forest := "forest_example" // string | Forest to get domain from.  If not specified, all forests are queried, which may take a long time. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties. If not specified, all properties other than IdentityDomainProperty.Controllers IdentityDomainProperty.PrimaryController will be retrieved. (optional) (default to 2147483647)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetDomain(context.Background(), domain).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetDomain`: IdentityDomainResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain to get details from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetDomainRequest struct via the builder pattern


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
 **forest** | **string** | Forest to get domain from.  If not specified, all forests are queried, which may take a long time. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties. If not specified, all properties other than IdentityDomainProperty.Controllers IdentityDomainProperty.PrimaryController will be retrieved. | [default to 2147483647]

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


## IdentityGetDomainAzureADCustomDomain

> AzureADCustomDomainResponseModel IdentityGetDomainAzureADCustomDomain(ctx, domain).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    domain := "domain_example" // string | Domain name.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetDomainAzureADCustomDomain(context.Background(), domain).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetDomainAzureADCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetDomainAzureADCustomDomain`: AzureADCustomDomainResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetDomainAzureADCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetDomainAzureADCustomDomainRequest struct via the builder pattern


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

[**AzureADCustomDomainResponseModel**](AzureADCustomDomainResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGetDomains

> IdentityDomainResponseModelCollection IdentityGetDomains(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).ParentDomain(parentDomain).Recursive(recursive).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    forest := "forest_example" // string | Forest to get domains from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := []string{"Inner_example"} // []string | Specific domain(s) to filter the results to.  If not specified, all matching domains are returned. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    parentDomain := "parentDomain_example" // string | Parent domain name to search.  Default is the root domain of the forest. (optional)
    recursive := true // bool | Specifies whether the search is recursive. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties. If not specified, all properties other than IdentityDomainProperty.Controllers IdentityDomainProperty.PrimaryController will be retrieved. (optional) (default to 2147483647)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)
    async := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetDomains(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).ParentDomain(parentDomain).Recursive(recursive).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetDomains`: IdentityDomainResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetDomainsRequest struct via the builder pattern


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
 **forest** | **string** | Forest to get domains from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **[]string** | Specific domain(s) to filter the results to.  If not specified, all matching domains are returned. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **parentDomain** | **string** | Parent domain name to search.  Default is the root domain of the forest. | 
 **recursive** | **bool** | Specifies whether the search is recursive. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties. If not specified, all properties other than IdentityDomainProperty.Controllers IdentityDomainProperty.PrimaryController will be retrieved. | [default to 2147483647]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 
 **async** | **bool** |  | [default to false]

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


## IdentityGetForest

> IdentityForestResponseModel IdentityGetForest(ctx, forest).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).PropertiesToRetrieve(propertiesToRetrieve).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    forest := "forest_example" // string | forest
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 2147483647)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetForest(context.Background(), forest).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).PropertiesToRetrieve(propertiesToRetrieve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetForest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetForest`: IdentityForestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetForest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forest** | **string** | forest | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetForestRequest struct via the builder pattern


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
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 2147483647]

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


## IdentityGetForests

> IdentityForestResponseModelCollection IdentityGetForests(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    forest := []string{"Inner_example"} // []string | Specific forest(s) to filter the results to.  If not specified, all forests are returned. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 2147483647)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetForests(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetForests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetForests`: IdentityForestResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetForests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetForestsRequest struct via the builder pattern


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
 **forest** | **[]string** | Specific forest(s) to filter the results to.  If not specified, all forests are returned. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 2147483647]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 

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


## IdentityGetMachine

> IdentityMachineResponseModel IdentityGetMachine(ctx, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    machine := "machine_example" // string | Machine to get.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    forest := "forest_example" // string | Forest to get machine account from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get machine account from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties other than IPAddress will be retrieved. (optional) (default to 2147483647)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetMachine(context.Background(), machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetMachine`: IdentityMachineResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetMachineRequest struct via the builder pattern


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
 **forest** | **string** | Forest to get machine account from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get machine account from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties other than IPAddress will be retrieved. | [default to 2147483647]

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


## IdentityGetMachines

> IdentityMachineResponseModelCollection IdentityGetMachines(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).Machine(machine).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    forest := "forest_example" // string | Forest to get machine accounts from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get machine accounts from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    machine := []string{"Inner_example"} // []string | Specific machine(s) to filter the results to.  If not specified, all matching machines are returned.  If set, `recursive` parameter is ignored and is implied `true`. (optional)
    startsWith := "startsWith_example" // string | Search for machine accounts that start with a string.  This parameter is exclusive with `contains`, and `matches`. (optional)
    contains := "contains_example" // string | Search for machine accounts that contain a string.  This parameter is exclusive with `startsWith`, and `matches`. (optional)
    matches := "matches_example" // string | Search for machine accounts that match a string.  This parameter is exclusive with `startsWith`, and `contains`. (optional)
    parentOU := "parentOU_example" // string | The parent OU to search.  If not specified, will search from the root OU. (optional)
    recursive := true // bool | Indicates whether the search should be recursive. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties other than IPAddress will be retrieved. (optional) (default to 2147483647)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetMachines(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).Machine(machine).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetMachines`: IdentityMachineResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetMachinesRequest struct via the builder pattern


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
 **forest** | **string** | Forest to get machine accounts from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get machine accounts from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **machine** | **[]string** | Specific machine(s) to filter the results to.  If not specified, all matching machines are returned.  If set, &#x60;recursive&#x60; parameter is ignored and is implied &#x60;true&#x60;. | 
 **startsWith** | **string** | Search for machine accounts that start with a string.  This parameter is exclusive with &#x60;contains&#x60;, and &#x60;matches&#x60;. | 
 **contains** | **string** | Search for machine accounts that contain a string.  This parameter is exclusive with &#x60;startsWith&#x60;, and &#x60;matches&#x60;. | 
 **matches** | **string** | Search for machine accounts that match a string.  This parameter is exclusive with &#x60;startsWith&#x60;, and &#x60;contains&#x60;. | 
 **parentOU** | **string** | The parent OU to search.  If not specified, will search from the root OU. | 
 **recursive** | **bool** | Indicates whether the search should be recursive. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties other than IPAddress will be retrieved. | [default to 2147483647]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 

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


## IdentityGetServiceAccount

> ServiceAccountResponseModel IdentityGetServiceAccount(ctx, serviceAccountUid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get a specific service account.

### Example

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
    serviceAccountUid := "serviceAccountUid_example" // string | The ServiceAccountUid of a specific service account. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetServiceAccount(context.Background(), serviceAccountUid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetServiceAccount`: ServiceAccountResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountUid** | **string** | The ServiceAccountUid of a specific service account.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetServiceAccountRequest struct via the builder pattern


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

[**ServiceAccountResponseModel**](ServiceAccountResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGetServiceAccounts

> ServiceAccountResponseModelCollection IdentityGetServiceAccounts(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get all service accounts.

### Example

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
    async := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetServiceAccounts(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetServiceAccounts`: ServiceAccountResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** |  | [default to false]

### Return type

[**ServiceAccountResponseModelCollection**](ServiceAccountResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGetSite

> IdentitySiteResponseModel IdentityGetSite(ctx, site).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).PropertiesToRetrieve(propertiesToRetrieve).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    site := "site_example" // string | The site to get details for.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    forest := "forest_example" // string | Forest to get the site from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get the site from.  If not specified, all domains in the forest(s) are queried until the site is found. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 2147483647)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetSite(context.Background(), site).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).PropertiesToRetrieve(propertiesToRetrieve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetSite`: IdentitySiteResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string** | The site to get details for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetSiteRequest struct via the builder pattern


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
 **forest** | **string** | Forest to get the site from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get the site from.  If not specified, all domains in the forest(s) are queried until the site is found. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 2147483647]

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


## IdentityGetSites

> IdentitySiteResponseModelCollection IdentityGetSites(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).HostNameOrIP(hostNameOrIP).Site(site).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    forest := "forest_example" // string | Forest to get sites from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get sites from.  If not specified, all domains in the forest(s) are queried. (optional)
    hostNameOrIP := "hostNameOrIP_example" // string | Hostname or IP to get sites from. (optional)
    site := []string{"Inner_example"} // []string | Specific site(s) to filter the results to.  If not specified, all matching sites are returned. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 2147483647)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetSites(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).HostNameOrIP(hostNameOrIP).Site(site).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetSites`: IdentitySiteResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetSitesRequest struct via the builder pattern


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
 **forest** | **string** | Forest to get sites from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get sites from.  If not specified, all domains in the forest(s) are queried. | 
 **hostNameOrIP** | **string** | Hostname or IP to get sites from. | 
 **site** | **[]string** | Specific site(s) to filter the results to.  If not specified, all matching sites are returned. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 2147483647]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 

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


## IdentityGetUser

> IdentityUserResponseModel IdentityGetUser(ctx, userOrGroup).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userOrGroup := "userOrGroup_example" // string | Identity of the user or group to get.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    forest := "forest_example" // string | Forest to get users from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get users from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 2147483647)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetUser(context.Background(), userOrGroup).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Forest(forest).Domain(domain).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetUser`: IdentityUserResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userOrGroup** | **string** | Identity of the user or group to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetUserRequest struct via the builder pattern


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
 **forest** | **string** | Forest to get users from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get users from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 2147483647]

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


## IdentityGetUsers

> IdentityUserResponseModelCollection IdentityGetUsers(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Provider(provider).Forest(forest).Domain(domain).Tenant(tenant).User(user).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).UserType(userType).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    provider := openapiclient.IdentityProviderType("Unknown") // IdentityProviderType | Provider to get users from. (optional)
    forest := "forest_example" // string | Forest to get users from.  If not specified, all forests are queried, which may take a long time. (optional)
    domain := "domain_example" // string | Domain to get users from.  If not specified, all domains in the forest(s) are queried, which may take a long time. (optional)
    tenant := "tenant_example" // string | Tenant to get users from. (optional)
    user := []string{"Inner_example"} // []string | Specific user(s) to filter the results to. If not specified, all matching users are returned.  If set, `recursive` parameter is ignored and is implied `true`. (optional)
    startsWith := "startsWith_example" // string | Search for users that start with a string.  This parameter is exclusive with `contains`, and `matches`. (optional)
    contains := "contains_example" // string | Search for users that contain a string.  This parameter is exclusive with `startsWith`, and `matches`. (optional)
    matches := "matches_example" // string | Search for users that match a string.  This parameter is exclusive with `startsWith`, and `contains`. (optional)
    parentOU := "parentOU_example" // string | The parent OU to search.  If not specified, will search from the root OU. (optional)
    recursive := true // bool | Indicates whether the search should be recursive. (optional)
    userType := openapiclient.IdentityUserType("Unknown") // IdentityUserType | Indicates user type(s) that should be retrieved.  If not specified, all user types will be searched. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    propertiesToRetrieve := int32(56) // int32 | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. (optional) (default to 2147483647)
    limit := int32(56) // int32 | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. (optional)
    continuationToken := "continuationToken_example" // string | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. (optional)
    async := true // bool | If `true`, the get users will be executed as a background task. The task will have JobType GetUsersIdentity. When the task is complete it will redirect to GetJobResults. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityGetUsers(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).Provider(provider).Forest(forest).Domain(domain).Tenant(tenant).User(user).StartsWith(startsWith).Contains(contains).Matches(matches).ParentOU(parentOU).Recursive(recursive).UserType(userType).DirectoryServerHint(directoryServerHint).PropertiesToRetrieve(propertiesToRetrieve).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetUsers`: IdentityUserResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityGetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetUsersRequest struct via the builder pattern


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
 **provider** | [**IdentityProviderType**](IdentityProviderType.md) | Provider to get users from. | 
 **forest** | **string** | Forest to get users from.  If not specified, all forests are queried, which may take a long time. | 
 **domain** | **string** | Domain to get users from.  If not specified, all domains in the forest(s) are queried, which may take a long time. | 
 **tenant** | **string** | Tenant to get users from. | 
 **user** | **[]string** | Specific user(s) to filter the results to. If not specified, all matching users are returned.  If set, &#x60;recursive&#x60; parameter is ignored and is implied &#x60;true&#x60;. | 
 **startsWith** | **string** | Search for users that start with a string.  This parameter is exclusive with &#x60;contains&#x60;, and &#x60;matches&#x60;. | 
 **contains** | **string** | Search for users that contain a string.  This parameter is exclusive with &#x60;startsWith&#x60;, and &#x60;matches&#x60;. | 
 **matches** | **string** | Search for users that match a string.  This parameter is exclusive with &#x60;startsWith&#x60;, and &#x60;contains&#x60;. | 
 **parentOU** | **string** | The parent OU to search.  If not specified, will search from the root OU. | 
 **recursive** | **bool** | Indicates whether the search should be recursive. | 
 **userType** | [**IdentityUserType**](IdentityUserType.md) | Indicates user type(s) that should be retrieved.  If not specified, all user types will be searched. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **propertiesToRetrieve** | **int32** | Properties to retrieve.  This should be specified as an integer representing the OR-ed together values of the properties.  If not specified, all properties will be retrieved. | [default to 2147483647]
 **limit** | **int32** | Maximum number of items to return.  If more items are available, a continuation token will be returned.  If not specified, all items will be returned. | 
 **continuationToken** | **string** | If specified, a previous query will be continued.  The caller must specify the same query parameters and admin credentials as the initial query or else the behavior is undefined. | 
 **async** | **bool** | If &#x60;true&#x60;, the get users will be executed as a background task. The task will have JobType GetUsersIdentity. When the task is complete it will redirect to GetJobResults. | [default to false]

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


## IdentityQueryAzureAdSecurityGroups

> AzureADSecurityGroupResponseModelCollection IdentityQueryAzureAdSecurityGroups(ctx, azureAdTenantId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Name(name).Assigned(assigned).MaxCount(maxCount).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    azureAdTenantId := "azureAdTenantId_example" // string | The specific azure tenant id.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAccessToken := "xAccessToken_example" // string | Access token to use when performing the operation. If specified, must be in a format matching that of the standard 'Authorization' request header; UTF8-encoded, then base64-encoded, then the \"Bearer\" scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ== (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    name := "name_example" // string | Specific the group display name. (optional)
    assigned := true // bool | When name is empty, assigned only support the value of true. (optional) (default to true)
    maxCount := int32(56) // int32 | The max return count,default is 300. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityQueryAzureAdSecurityGroups(context.Background(), azureAdTenantId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAccessToken(xAccessToken).Accept(accept).CitrixLocale(citrixLocale).Name(name).Assigned(assigned).MaxCount(maxCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityQueryAzureAdSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityQueryAzureAdSecurityGroups`: AzureADSecurityGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityQueryAzureAdSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAdTenantId** | **string** | The specific azure tenant id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityQueryAzureAdSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAccessToken** | **string** | Access token to use when performing the operation. If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; UTF8-encoded, then base64-encoded, then the \&quot;Bearer\&quot; scheme prepended.  Example: Bearer bGljaGVuZy5saW5AY2l0cml4LmNvbQ&#x3D;&#x3D; | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **name** | **string** | Specific the group display name. | 
 **assigned** | **bool** | When name is empty, assigned only support the value of true. | [default to true]
 **maxCount** | **int32** | The max return count,default is 300. | 

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


## IdentitySetServiceAccount

> UpdateServiceAccountRequestModel IdentitySetServiceAccount(ctx, serviceAccountUid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateServiceAccountRequestModel(updateServiceAccountRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Set properties of an existing service account.

### Example

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
    serviceAccountUid := "serviceAccountUid_example" // string | The ServiceAccountUid of service account. 
    updateServiceAccountRequestModel := *openapiclient.NewUpdateServiceAccountRequestModel() // UpdateServiceAccountRequestModel | The object of UpdateServiceAccountRequestModel used to set an existing service account. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentitySetServiceAccount(context.Background(), serviceAccountUid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateServiceAccountRequestModel(updateServiceAccountRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentitySetServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentitySetServiceAccount`: UpdateServiceAccountRequestModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentitySetServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountUid** | **string** | The ServiceAccountUid of service account.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentitySetServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **updateServiceAccountRequestModel** | [**UpdateServiceAccountRequestModel**](UpdateServiceAccountRequestModel.md) | The object of UpdateServiceAccountRequestModel used to set an existing service account.  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**UpdateServiceAccountRequestModel**](UpdateServiceAccountRequestModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityUpdateMachine

> IdentityMachineResponseModel IdentityUpdateMachine(ctx, machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).Forest(forest).Domain(domain).IdentityUpdateMachineRequestModel(identityUpdateMachineRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DirectoryServerHint(directoryServerHint).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    machine := "machine_example" // string | Machine identity to update.
    forest := "forest_example" // string | Forest to update machine account within.  Required.
    domain := "domain_example" // string | Domain to update machine account within.  Required.
    identityUpdateMachineRequestModel := *openapiclient.NewIdentityUpdateMachineRequestModel() // IdentityUpdateMachineRequestModel | Properties of the machine identity that should be updated.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    async := true // bool | If `true`, the machine identity will be updated as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityUpdateMachine(context.Background(), machine).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).Forest(forest).Domain(domain).IdentityUpdateMachineRequestModel(identityUpdateMachineRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DirectoryServerHint(directoryServerHint).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityUpdateMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityUpdateMachine`: IdentityMachineResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityUpdateMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**machine** | **string** | Machine identity to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityUpdateMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **forest** | **string** | Forest to update machine account within.  Required. | 
 **domain** | **string** | Domain to update machine account within.  Required. | 
 **identityUpdateMachineRequestModel** | [**IdentityUpdateMachineRequestModel**](IdentityUpdateMachineRequestModel.md) | Properties of the machine identity that should be updated. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **async** | **bool** | If &#x60;true&#x60;, the machine identity will be updated as a background task. | [default to false]

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


## IdentityUpdateUser

> IdentityUserResponseModel IdentityUpdateUser(ctx, user).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).Forest(forest).Domain(domain).IdentityUpdateUserRequestModel(identityUpdateUserRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DirectoryServerHint(directoryServerHint).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    user := "user_example" // string | User identity to update.
    forest := "forest_example" // string | Forest to update user account within.  Required.
    domain := "domain_example" // string | Domain to update user account within.  Required.
    identityUpdateUserRequestModel := *openapiclient.NewIdentityUpdateUserRequestModel() // IdentityUpdateUserRequestModel | Properties of the user identity that should be updated.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    directoryServerHint := "directoryServerHint_example" // string | Hint to inform the system of a directory server which is most likely to successfully perform the operation. (optional)
    async := true // bool | If `true`, the user identity will be updated as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityUpdateUser(context.Background(), user).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).Forest(forest).Domain(domain).IdentityUpdateUserRequestModel(identityUpdateUserRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).Accept(accept).CitrixLocale(citrixLocale).DirectoryServerHint(directoryServerHint).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityUpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityUpdateUser`: IdentityUserResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityUpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | User identity to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **forest** | **string** | Forest to update user account within.  Required. | 
 **domain** | **string** | Domain to update user account within.  Required. | 
 **identityUpdateUserRequestModel** | [**IdentityUpdateUserRequestModel**](IdentityUpdateUserRequestModel.md) | Properties of the user identity that should be updated. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **directoryServerHint** | **string** | Hint to inform the system of a directory server which is most likely to successfully perform the operation. | 
 **async** | **bool** | If &#x60;true&#x60;, the user identity will be updated as a background task. | [default to false]

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


## IdentityValidateUser

> IdentityUserResponseModel IdentityValidateUser(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).IdentityValidateUserRequestModel(identityValidateUserRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    identityValidateUserRequestModel := *openapiclient.NewIdentityValidateUserRequestModel("For example, administrator", "For example, password") // IdentityValidateUserRequestModel | Details of the user credentials to validate.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAPIsDAAS.IdentityValidateUser(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).IdentityValidateUserRequestModel(identityValidateUserRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPIsDAAS.IdentityValidateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityValidateUser`: IdentityUserResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IdentityAPIsDAAS.IdentityValidateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityValidateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **identityValidateUserRequestModel** | [**IdentityValidateUserRequestModel**](IdentityValidateUserRequestModel.md) | Details of the user credentials to validate. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

