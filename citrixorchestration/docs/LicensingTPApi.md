# \LicensingTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LicensingTPAddLicense**](LicensingTPApi.md#LicensingTPAddLicense) | **Post** /techpreview/{customerid}/{siteid}/Licensing/AddLicense | Add license to license server.
[**LicensingTPAllocateLicense**](LicensingTPApi.md#LicensingTPAllocateLicense) | **Post** /techpreview/{customerid}/{siteid}/Licensing/AllocateLicense | Allocate license by the License Access Code
[**LicensingTPCheckLicenseCertificate**](LicensingTPApi.md#LicensingTPCheckLicenseCertificate) | **Get** /techpreview/{customerid}/{siteid}/Licensing/CheckLicenseCertificate | Check the certificate of the license server
[**LicensingTPCreateLicenseAdministrator**](LicensingTPApi.md#LicensingTPCreateLicenseAdministrator) | **Post** /techpreview/{customerid}/{siteid}/Licensing/LicenseAdministrators | Create a new license administrator
[**LicensingTPDeleteLicenseAdministrator**](LicensingTPApi.md#LicensingTPDeleteLicenseAdministrator) | **Delete** /techpreview/{customerid}/{siteid}/Licensing/LicenseAdministrators/{nameOrId} | Delete a license administrator
[**LicensingTPEditLicenseProductEdition**](LicensingTPApi.md#LicensingTPEditLicenseProductEdition) | **Patch** /techpreview/{customerid}/{siteid}/Licensing/LicenseProductEdition | Update the product license edition
[**LicensingTPGetLicenseAdministrators**](LicensingTPApi.md#LicensingTPGetLicenseAdministrators) | **Get** /techpreview/{customerid}/{siteid}/Licensing/LicenseAdministrators | Get the license administrators
[**LicensingTPGetLicenseCertificate**](LicensingTPApi.md#LicensingTPGetLicenseCertificate) | **Get** /techpreview/{customerid}/{siteid}/Licensing/LicenseCertificate | Get the certificate of the license server
[**LicensingTPGetLicenseEntitlements**](LicensingTPApi.md#LicensingTPGetLicenseEntitlements) | **Post** /techpreview/{customerid}/{siteid}/Licensing/LicenseEntitlements | Get license entitlements by the License Access Code
[**LicensingTPGetLicensePermission**](LicensingTPApi.md#LicensingTPGetLicensePermission) | **Get** /techpreview/{customerid}/{siteid}/Licensing/LicensePermission | Get the permission to the license server
[**LicensingTPGetLicenseProductEdition**](LicensingTPApi.md#LicensingTPGetLicenseProductEdition) | **Get** /techpreview/{customerid}/{siteid}/Licensing/LicenseProductEdition | Get License product edition
[**LicensingTPGetLicenseSummary**](LicensingTPApi.md#LicensingTPGetLicenseSummary) | **Get** /techpreview/{customerid}/{siteid}/Licensing/LicenseOverview | Get the license overview
[**LicensingTPGetLicensingAlert**](LicensingTPApi.md#LicensingTPGetLicensingAlert) | **Get** /techpreview/{customerid}/{siteid}/Licensing/LicenseAlert | Get alert for license.
[**LicensingTPSetLicenseServer**](LicensingTPApi.md#LicensingTPSetLicenseServer) | **Patch** /techpreview/{customerid}/{siteid}/Licensing/LicenseServer | Change license server
[**LicensingTPUpdateLicenseAdministrator**](LicensingTPApi.md#LicensingTPUpdateLicenseAdministrator) | **Patch** /techpreview/{customerid}/{siteid}/Licensing/LicenseAdministrators/{nameOrId} | Update the license administrator permission



## LicensingTPAddLicense

> AddLicenseResponseModel LicensingTPAddLicense(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Add license to license server.

### Example

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
    request := *openapiclient.NewAddLicenseRequestModel("LicenseFile_example", "FileName_example", false) // AddLicenseRequestModel | Request model to add license file.
    async := true // bool | If `true`, Adding license will be executed as a background task. The task will have JobType AddLicense. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingTPApi.LicensingTPAddLicense(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPAddLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicensingTPAddLicense`: AddLicenseResponseModel
    fmt.Fprintf(os.Stdout, "Response from `LicensingTPApi.LicensingTPAddLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPAddLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**AddLicenseRequestModel**](AddLicenseRequestModel.md) | Request model to add license file. | 
 **async** | **bool** | If &#x60;true&#x60;, Adding license will be executed as a background task. The task will have JobType AddLicense. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AddLicenseResponseModel**](AddLicenseResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LicensingTPAllocateLicense

> LicensingTPAllocateLicense(ctx, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Allocate license by the License Access Code

### Example

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
    model := *openapiclient.NewLicenseEntitlementRequestModel("LicenseAccessCode_example", "Id_example") // LicenseEntitlementRequestModel | The request model to allocate license
    async := true // bool | If `true`, allocating license will be executed as a background task. The task will have JobType AllocateLicense. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LicensingTPApi.LicensingTPAllocateLicense(context.Background(), customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPAllocateLicense``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLicensingTPAllocateLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**LicenseEntitlementRequestModel**](LicenseEntitlementRequestModel.md) | The request model to allocate license | 
 **async** | **bool** | If &#x60;true&#x60;, allocating license will be executed as a background task. The task will have JobType AllocateLicense. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
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


## LicensingTPCheckLicenseCertificate

> LicensingTPCheckLicenseCertificate(ctx, customerid, siteid).AdminServer(adminServer).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check the certificate of the license server

### Example

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
    adminServer := "adminServer_example" // string |  (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LicensingTPApi.LicensingTPCheckLicenseCertificate(context.Background(), customerid, siteid).AdminServer(adminServer).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPCheckLicenseCertificate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLicensingTPCheckLicenseCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adminServer** | **string** |  | 
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


## LicensingTPCreateLicenseAdministrator

> LicensingTPCreateLicenseAdministrator(ctx, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Create a new license administrator

### Example

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
    request := *openapiclient.NewCreateLicensingAdminRequestModel("AccountSamName_example") // CreateLicensingAdminRequestModel | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LicensingTPApi.LicensingTPCreateLicenseAdministrator(context.Background(), customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPCreateLicenseAdministrator``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLicensingTPCreateLicenseAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateLicensingAdminRequestModel**](CreateLicensingAdminRequestModel.md) |  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
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


## LicensingTPDeleteLicenseAdministrator

> LicensingTPDeleteLicenseAdministrator(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Delete a license administrator

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | 
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LicensingTPApi.LicensingTPDeleteLicenseAdministrator(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPDeleteLicenseAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPDeleteLicenseAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## LicensingTPEditLicenseProductEdition

> LicensingTPEditLicenseProductEdition(ctx, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update the product license edition

### Example

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
    request := *openapiclient.NewEditLicenseProductEditionRequestModel() // EditLicenseProductEditionRequestModel | Request model to update License edition.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LicensingTPApi.LicensingTPEditLicenseProductEdition(context.Background(), customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPEditLicenseProductEdition``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLicensingTPEditLicenseProductEditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**EditLicenseProductEditionRequestModel**](EditLicenseProductEditionRequestModel.md) | Request model to update License edition. | 
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


## LicensingTPGetLicenseAdministrators

> LicensingUserResponseModelCollection LicensingTPGetLicenseAdministrators(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get the license administrators

### Example

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
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingTPApi.LicensingTPGetLicenseAdministrators(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPGetLicenseAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicensingTPGetLicenseAdministrators`: LicensingUserResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `LicensingTPApi.LicensingTPGetLicenseAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPGetLicenseAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LicensingUserResponseModelCollection**](LicensingUserResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LicensingTPGetLicenseCertificate

> LicenseCertificateResponseModel LicensingTPGetLicenseCertificate(ctx, customerid, siteid).AdminServer(adminServer).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the certificate of the license server

### Example

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
    adminServer := "adminServer_example" // string |  (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingTPApi.LicensingTPGetLicenseCertificate(context.Background(), customerid, siteid).AdminServer(adminServer).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPGetLicenseCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicensingTPGetLicenseCertificate`: LicenseCertificateResponseModel
    fmt.Fprintf(os.Stdout, "Response from `LicensingTPApi.LicensingTPGetLicenseCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPGetLicenseCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adminServer** | **string** |  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LicenseCertificateResponseModel**](LicenseCertificateResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LicensingTPGetLicenseEntitlements

> LicenseEntitlementResponseModelCollection LicensingTPGetLicenseEntitlements(ctx, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get license entitlements by the License Access Code

### Example

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
    model := *openapiclient.NewLicenseAccessCodeRequestModel("LicenseAccessCode_example") // LicenseAccessCodeRequestModel | The request model to get license entitlement
    async := true // bool | If `true`, getting licensing entitlements will be executed as a background task. The task will have JobType GetLicenseEntitlement. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingTPApi.LicensingTPGetLicenseEntitlements(context.Background(), customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPGetLicenseEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicensingTPGetLicenseEntitlements`: LicenseEntitlementResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `LicensingTPApi.LicensingTPGetLicenseEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPGetLicenseEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**LicenseAccessCodeRequestModel**](LicenseAccessCodeRequestModel.md) | The request model to get license entitlement | 
 **async** | **bool** | If &#x60;true&#x60;, getting licensing entitlements will be executed as a background task. The task will have JobType GetLicenseEntitlement. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LicenseEntitlementResponseModelCollection**](LicenseEntitlementResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LicensingTPGetLicensePermission

> LicensePermissionResponseModel LicensingTPGetLicensePermission(ctx, customerid, siteid).OnConnect(onConnect).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get the permission to the license server

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    onConnect := true // bool | If true, AuthorizationFailed results in None being returned; otherwise an exception is thrown.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, to get the permission to the license server will be created as a background task.             The task will have JobType GetLicensePermission.              (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingTPApi.LicensingTPGetLicensePermission(context.Background(), customerid, siteid).OnConnect(onConnect).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPGetLicensePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicensingTPGetLicensePermission`: LicensePermissionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `LicensingTPApi.LicensingTPGetLicensePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPGetLicensePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onConnect** | **bool** | If true, AuthorizationFailed results in None being returned; otherwise an exception is thrown. | 


 **async** | **bool** | If &#x60;true&#x60;, to get the permission to the license server will be created as a background task.             The task will have JobType GetLicensePermission.              | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LicensePermissionResponseModel**](LicensePermissionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LicensingTPGetLicenseProductEdition

> LicenseProductEditionResponseModel LicensingTPGetLicenseProductEdition(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get License product edition

### Example

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
    resp, r, err := apiClient.LicensingTPApi.LicensingTPGetLicenseProductEdition(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPGetLicenseProductEdition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicensingTPGetLicenseProductEdition`: LicenseProductEditionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `LicensingTPApi.LicensingTPGetLicenseProductEdition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPGetLicenseProductEditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LicenseProductEditionResponseModel**](LicenseProductEditionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LicensingTPGetLicenseSummary

> LicenseSummaryResponseModel LicensingTPGetLicenseSummary(ctx, customerid, siteid).IncludeInventory(includeInventory).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the license overview

### Example

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
    includeInventory := true // bool | If True, license inventories will be returned in the response. (optional) (default to false)
    async := true // bool | If `true`, to get the license overview information will be created as a background task.             The task will have JobType GetLicenseSummary.              (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingTPApi.LicensingTPGetLicenseSummary(context.Background(), customerid, siteid).IncludeInventory(includeInventory).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPGetLicenseSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicensingTPGetLicenseSummary`: LicenseSummaryResponseModel
    fmt.Fprintf(os.Stdout, "Response from `LicensingTPApi.LicensingTPGetLicenseSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPGetLicenseSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeInventory** | **bool** | If True, license inventories will be returned in the response. | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, to get the license overview information will be created as a background task.             The task will have JobType GetLicenseSummary.              | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LicenseSummaryResponseModel**](LicenseSummaryResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LicensingTPGetLicensingAlert

> LicensingAlertModel LicensingTPGetLicensingAlert(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get alert for license.

### Example

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
    async := true // bool |  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingTPApi.LicensingTPGetLicensingAlert(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPGetLicensingAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicensingTPGetLicensingAlert`: LicensingAlertModel
    fmt.Fprintf(os.Stdout, "Response from `LicensingTPApi.LicensingTPGetLicensingAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPGetLicensingAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** |  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**LicensingAlertModel**](LicensingAlertModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LicensingTPSetLicenseServer

> LicensingTPSetLicenseServer(ctx, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Change license server

### Example

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
    model := *openapiclient.NewEditLicenseServerRequestModel("LicenseServer_example") // EditLicenseServerRequestModel | Request model to set license server.
    async := true // bool | If `true`, setting license server  will be executed as a background task. The task will have JobType SetLicenseServer. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LicensingTPApi.LicensingTPSetLicenseServer(context.Background(), customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPSetLicenseServer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLicensingTPSetLicenseServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**EditLicenseServerRequestModel**](EditLicenseServerRequestModel.md) | Request model to set license server. | 
 **async** | **bool** | If &#x60;true&#x60;, setting license server  will be executed as a background task. The task will have JobType SetLicenseServer. | [default to false]
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


## LicensingTPUpdateLicenseAdministrator

> LicensingTPUpdateLicenseAdministrator(ctx, nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Update the license administrator permission

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | SamName or SID of the administrator account.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewUpdateLicensingAdminRequestModel() // UpdateLicensingAdminRequestModel | The update request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LicensingTPApi.LicensingTPUpdateLicenseAdministrator(context.Background(), nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingTPApi.LicensingTPUpdateLicenseAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | SamName or SID of the administrator account. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLicensingTPUpdateLicenseAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateLicensingAdminRequestModel**](UpdateLicensingAdminRequestModel.md) | The update request. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
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

