# \SitesAPIs 

All URIs are relative to *https://api-us.cloud.com/cvad/manage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesGetLocalAppAccessConfig**](SitesAPIs .md#SitesGetLocalAppAccessConfig) | **Get** /Sites/{nameOrId}/LocalAppAccessConfig | Get Local App Access Config.
[**SitesGetMultipleRemotePCAssignments**](SitesAPIs .md#SitesGetMultipleRemotePCAssignments) | **Get** /Sites/{nameOrId}/MultipleRemotePCAssignments | Get multi-user auto-assignment for Remote PC Access.
[**SitesGetNewSiteSecurityKey**](SitesAPIs .md#SitesGetNewSiteSecurityKey) | **Get** /Sites/{nameOrId}/NewSecurityKey | Get a new security key for a site.
[**SitesGetSessionsTrend**](SitesAPIs .md#SitesGetSessionsTrend) | **Get** /Sites/{nameOrId}/SessionsTrend | Get the sessions trend
[**SitesGetSite**](SitesAPIs .md#SitesGetSite) | **Get** /Sites/{nameOrId} | Get the details about a single site.
[**SitesGetSiteLicense**](SitesAPIs .md#SitesGetSiteLicense) | **Get** /Sites/{nameOrId}/License | Get the license in use for a site.
[**SitesGetSiteSecurityKeyConfiguration**](SitesAPIs .md#SitesGetSiteSecurityKeyConfiguration) | **Get** /Sites/{nameOrId}/SecurityKeyConfiguration | Get the security key configuration for a site.
[**SitesGetSiteSecurityKeyManagementConfig**](SitesAPIs .md#SitesGetSiteSecurityKeyManagementConfig) | **Get** /Sites/{nameOrId}/SecurityKeyManagementConfig | Get security key management config for a site.
[**SitesGetSiteSettings**](SitesAPIs .md#SitesGetSiteSettings) | **Get** /Sites/{nameOrId}/Settings | Get the settings for the site.
[**SitesGetSiteValidLicenses**](SitesAPIs .md#SitesGetSiteValidLicenses) | **Get** /Sites/{nameOrId}/ValidLicenses | Get valid licenses for a site.
[**SitesGetUpgradePackageVersions**](SitesAPIs .md#SitesGetUpgradePackageVersions) | **Get** /Sites/{nameOrId}/UpgradePackageVersions | Get the latest released VDA upgrade package versions in the site.
[**SitesPatchMultipleRemotePCAssignments**](SitesAPIs .md#SitesPatchMultipleRemotePCAssignments) | **Patch** /Sites/{nameOrId}/MultipleRemotePCAssignments | Update multi-user auto-assignment for Remote PC Access.
[**SitesPatchSiteLicense**](SitesAPIs .md#SitesPatchSiteLicense) | **Patch** /Sites/{nameOrId}/License | Update the license in use for a site.
[**SitesPatchSiteSecurityKeyConfiguration**](SitesAPIs .md#SitesPatchSiteSecurityKeyConfiguration) | **Patch** /Sites/{nameOrId}/SecurityKeyConfiguration | Update the security key configuration for a site.
[**SitesPatchSiteSettings**](SitesAPIs .md#SitesPatchSiteSettings) | **Patch** /Sites/{nameOrId}/Settings | Update the broker site settings.



## SitesGetLocalAppAccessConfig

> NameValueStringPairModel SitesGetLocalAppAccessConfig(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get Local App Access Config.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetLocalAppAccessConfig(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetLocalAppAccessConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetLocalAppAccessConfig`: NameValueStringPairModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetLocalAppAccessConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetLocalAppAccessConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**NameValueStringPairModel**](NameValueStringPairModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetMultipleRemotePCAssignments

> bool SitesGetMultipleRemotePCAssignments(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get multi-user auto-assignment for Remote PC Access.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetMultipleRemotePCAssignments(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetMultipleRemotePCAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetMultipleRemotePCAssignments`: bool
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetMultipleRemotePCAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetMultipleRemotePCAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetNewSiteSecurityKey

> NameValueStringPairModel SitesGetNewSiteSecurityKey(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get a new security key for a site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetNewSiteSecurityKey(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetNewSiteSecurityKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetNewSiteSecurityKey`: NameValueStringPairModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetNewSiteSecurityKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetNewSiteSecurityKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**NameValueStringPairModel**](NameValueStringPairModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetSessionsTrend

> SessionsTrendResponseModel SitesGetSessionsTrend(ctx, nameOrId).StartDate(startDate).EndDate(endDate).IntervalLength(intervalLength).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the sessions trend

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | The site name or ID.
    startDate := "startDate_example" // string | The start date of sessions trend to query, for example '2021-11-01T12:00:00'.
    endDate := "endDate_example" // string | The end date of sessions trend to query, for example '2021-11-08T12:00:00'.
    intervalLength := int32(56) // int32 | The minutes interval to query.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetSessionsTrend(context.Background(), nameOrId).StartDate(startDate).EndDate(endDate).IntervalLength(intervalLength).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetSessionsTrend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSessionsTrend`: SessionsTrendResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetSessionsTrend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The site name or ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSessionsTrendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | The start date of sessions trend to query, for example &#39;2021-11-01T12:00:00&#39;. | 
 **endDate** | **string** | The end date of sessions trend to query, for example &#39;2021-11-08T12:00:00&#39;. | 
 **intervalLength** | **int32** | The minutes interval to query. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**SessionsTrendResponseModel**](SessionsTrendResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetSite

> SiteDetailResponseModel SitesGetSite(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the details about a single site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetSite(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSite`: SiteDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**SiteDetailResponseModel**](SiteDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetSiteLicense

> LicenseResponseModel SitesGetSiteLicense(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the license in use for a site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetSiteLicense(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetSiteLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSiteLicense`: LicenseResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetSiteLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSiteLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**LicenseResponseModel**](LicenseResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetSiteSecurityKeyConfiguration

> SecurityKeyConfigurationResponseModel SitesGetSiteSecurityKeyConfiguration(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the security key configuration for a site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetSiteSecurityKeyConfiguration(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetSiteSecurityKeyConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSiteSecurityKeyConfiguration`: SecurityKeyConfigurationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetSiteSecurityKeyConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSiteSecurityKeyConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**SecurityKeyConfigurationResponseModel**](SecurityKeyConfigurationResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetSiteSecurityKeyManagementConfig

> NameValueStringPairModel SitesGetSiteSecurityKeyManagementConfig(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get security key management config for a site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetSiteSecurityKeyManagementConfig(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetSiteSecurityKeyManagementConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSiteSecurityKeyManagementConfig`: NameValueStringPairModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetSiteSecurityKeyManagementConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSiteSecurityKeyManagementConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**NameValueStringPairModel**](NameValueStringPairModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetSiteSettings

> SiteSettingsResponseModel SitesGetSiteSettings(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the settings for the site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetSiteSettings(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetSiteSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSiteSettings`: SiteSettingsResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetSiteSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSiteSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**SiteSettingsResponseModel**](SiteSettingsResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetSiteValidLicenses

> LicenseResponseModelCollection SitesGetSiteValidLicenses(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get valid licenses for a site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetSiteValidLicenses(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetSiteValidLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSiteValidLicenses`: LicenseResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetSiteValidLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSiteValidLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**LicenseResponseModelCollection**](LicenseResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetUpgradePackageVersions

> UpgradePackageVersionResponseModelCollection SitesGetUpgradePackageVersions(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the latest released VDA upgrade package versions in the site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIs .SitesGetUpgradePackageVersions(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesGetUpgradePackageVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetUpgradePackageVersions`: UpgradePackageVersionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIs .SitesGetUpgradePackageVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetUpgradePackageVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**UpgradePackageVersionResponseModelCollection**](UpgradePackageVersionResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesPatchMultipleRemotePCAssignments

> SitesPatchMultipleRemotePCAssignments(ctx, nameOrId).Allow(allow).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update multi-user auto-assignment for Remote PC Access.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    allow := true // bool | specify if allow multi-user auto-assignment for Remote PC Access.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesAPIs .SitesPatchMultipleRemotePCAssignments(context.Background(), nameOrId).Allow(allow).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesPatchMultipleRemotePCAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesPatchMultipleRemotePCAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allow** | **bool** | specify if allow multi-user auto-assignment for Remote PC Access. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
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


## SitesPatchSiteLicense

> SitesPatchSiteLicense(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).LicenseRequestModel(licenseRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update the license in use for a site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    licenseRequestModel := *openapiclient.NewLicenseRequestModel() // LicenseRequestModel | The license update request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesAPIs .SitesPatchSiteLicense(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).LicenseRequestModel(licenseRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesPatchSiteLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesPatchSiteLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **licenseRequestModel** | [**LicenseRequestModel**](LicenseRequestModel.md) | The license update request. | 
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


## SitesPatchSiteSecurityKeyConfiguration

> SitesPatchSiteSecurityKeyConfiguration(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).SecurityKeyConfigurationRequestModel(securityKeyConfigurationRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update the security key configuration for a site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    securityKeyConfigurationRequestModel := *openapiclient.NewSecurityKeyConfigurationRequestModel() // SecurityKeyConfigurationRequestModel | Site security key configuration update request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesAPIs .SitesPatchSiteSecurityKeyConfiguration(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).SecurityKeyConfigurationRequestModel(securityKeyConfigurationRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesPatchSiteSecurityKeyConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesPatchSiteSecurityKeyConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **securityKeyConfigurationRequestModel** | [**SecurityKeyConfigurationRequestModel**](SecurityKeyConfigurationRequestModel.md) | Site security key configuration update request. | 
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


## SitesPatchSiteSettings

> SitesPatchSiteSettings(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).EditSiteSettingsRequestModel(editSiteSettingsRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update the broker site settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    editSiteSettingsRequestModel := *openapiclient.NewEditSiteSettingsRequestModel() // EditSiteSettingsRequestModel | site settings request model.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesAPIs .SitesPatchSiteSettings(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).EditSiteSettingsRequestModel(editSiteSettingsRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIs .SitesPatchSiteSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesPatchSiteSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **editSiteSettingsRequestModel** | [**EditSiteSettingsRequestModel**](EditSiteSettingsRequestModel.md) | site settings request model. | 
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

