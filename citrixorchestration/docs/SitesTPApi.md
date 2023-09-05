# \SitesTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesTPGetCeipParticipation**](SitesTPApi.md#SitesTPGetCeipParticipation) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/CustomerExperienceImprovementProgram | Get the Customer Experience Improvement Program Participation for the site.
[**SitesTPGetDnsResolution**](SitesTPApi.md#SitesTPGetDnsResolution) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/DnsResolution | Get the DNS Resolution for the site.
[**SitesTPGetLocalAppAccessConfig**](SitesTPApi.md#SitesTPGetLocalAppAccessConfig) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/LocalAppAccessConfig | Get Local App Access Config.
[**SitesTPGetMultipleRemotePCAssignments**](SitesTPApi.md#SitesTPGetMultipleRemotePCAssignments) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/MultipleRemotePCAssignments | Get multi-user auto-assignment for Remote PC Access.
[**SitesTPGetNewSiteSecurityKey**](SitesTPApi.md#SitesTPGetNewSiteSecurityKey) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/NewSecurityKey | Get a new security key for a site.
[**SitesTPGetSessionsTrend**](SitesTPApi.md#SitesTPGetSessionsTrend) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/SessionsTrend | Get the sessions trend
[**SitesTPGetSite**](SitesTPApi.md#SitesTPGetSite) | **Get** /techpreview/{customerid}/Sites/{nameOrId} | Get the details about a single site.
[**SitesTPGetSiteLicense**](SitesTPApi.md#SitesTPGetSiteLicense) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/License | Get the license in use for a site.
[**SitesTPGetSiteLoadBalancingOption**](SitesTPApi.md#SitesTPGetSiteLoadBalancingOption) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/LoadBalancingOption | Get the load balancing option for the site.
[**SitesTPGetSiteSecurityKeyConfiguration**](SitesTPApi.md#SitesTPGetSiteSecurityKeyConfiguration) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/SecurityKeyConfiguration | Get the security key configuration for a site.
[**SitesTPGetSiteSecurityKeyManagementConfig**](SitesTPApi.md#SitesTPGetSiteSecurityKeyManagementConfig) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/SecurityKeyManagementConfig | Get security key management config for a site.
[**SitesTPGetSiteSettings**](SitesTPApi.md#SitesTPGetSiteSettings) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/Settings | Get the settings for the site.
[**SitesTPGetSiteStatus**](SitesTPApi.md#SitesTPGetSiteStatus) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/Status | Get the status of a site.
[**SitesTPGetSiteTestReport**](SitesTPApi.md#SitesTPGetSiteTestReport) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/TestReport | Get the most recent test report.
[**SitesTPGetSiteValidLicenses**](SitesTPApi.md#SitesTPGetSiteValidLicenses) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/ValidLicenses | Get valid licenses for a site.
[**SitesTPGetSites**](SitesTPApi.md#SitesTPGetSites) | **Get** /techpreview/{customerid}/Sites | Get the list of sites that are available to the customer and visible to the admin.
[**SitesTPGetUpgradePackageVersions**](SitesTPApi.md#SitesTPGetUpgradePackageVersions) | **Get** /techpreview/{customerid}/Sites/{nameOrId}/UpgradePackageVersions | Get the latest released VDA upgrade package versions in the site.
[**SitesTPPatchCeipParticipation**](SitesTPApi.md#SitesTPPatchCeipParticipation) | **Patch** /techpreview/{customerid}/Sites/{nameOrId}/CustomerExperienceImprovementProgram | Update the Customer Experience Improvement Program Participation for the site.
[**SitesTPPatchDnsResolution**](SitesTPApi.md#SitesTPPatchDnsResolution) | **Patch** /techpreview/{customerid}/Sites/{nameOrId}/DnsResolution | Update the broker site DNS resolution for the site.
[**SitesTPPatchMultipleRemotePCAssignments**](SitesTPApi.md#SitesTPPatchMultipleRemotePCAssignments) | **Patch** /techpreview/{customerid}/Sites/{nameOrId}/MultipleRemotePCAssignments | Update multi-user auto-assignment for Remote PC Access.
[**SitesTPPatchSiteLicense**](SitesTPApi.md#SitesTPPatchSiteLicense) | **Patch** /techpreview/{customerid}/Sites/{nameOrId}/License | Update the license in use for a site.
[**SitesTPPatchSiteLoadBalancingOption**](SitesTPApi.md#SitesTPPatchSiteLoadBalancingOption) | **Patch** /techpreview/{customerid}/Sites/{nameOrId}/LoadBalancingOption | Update the load balancing option for the site.
[**SitesTPPatchSiteSecurityKeyConfiguration**](SitesTPApi.md#SitesTPPatchSiteSecurityKeyConfiguration) | **Patch** /techpreview/{customerid}/Sites/{nameOrId}/SecurityKeyConfiguration | Update the security key configuration for a site.
[**SitesTPPatchSiteSettings**](SitesTPApi.md#SitesTPPatchSiteSettings) | **Patch** /techpreview/{customerid}/Sites/{nameOrId}/Settings | Update the broker site settings.
[**SitesTPTestSite**](SitesTPApi.md#SitesTPTestSite) | **Post** /techpreview/{customerid}/Sites/{nameOrId}/$test | Run tests on a site and create a test report.



## SitesTPGetCeipParticipation

> bool SitesTPGetCeipParticipation(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the Customer Experience Improvement Program Participation for the site.

### Example

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetCeipParticipation(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetCeipParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetCeipParticipation`: bool
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetCeipParticipation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetCeipParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetDnsResolution

> bool SitesTPGetDnsResolution(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the DNS Resolution for the site.

### Example

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetDnsResolution(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetDnsResolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetDnsResolution`: bool
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetDnsResolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetDnsResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetLocalAppAccessConfig

> NameValueStringPairModel SitesTPGetLocalAppAccessConfig(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetLocalAppAccessConfig(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetLocalAppAccessConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetLocalAppAccessConfig`: NameValueStringPairModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetLocalAppAccessConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetLocalAppAccessConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetMultipleRemotePCAssignments

> bool SitesTPGetMultipleRemotePCAssignments(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetMultipleRemotePCAssignments(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetMultipleRemotePCAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetMultipleRemotePCAssignments`: bool
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetMultipleRemotePCAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetMultipleRemotePCAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetNewSiteSecurityKey

> NameValueStringPairModel SitesTPGetNewSiteSecurityKey(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetNewSiteSecurityKey(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetNewSiteSecurityKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetNewSiteSecurityKey`: NameValueStringPairModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetNewSiteSecurityKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetNewSiteSecurityKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetSessionsTrend

> SessionsTrendResponseModel SitesTPGetSessionsTrend(ctx, customerid, nameOrId).StartDate(startDate).EndDate(endDate).IntervalLength(intervalLength).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | The customer Id.
    nameOrId := "nameOrId_example" // string | The site name or ID.
    startDate := "startDate_example" // string | The start date of sessions trend to query, for example '2021-11-01T12:00:00'.
    endDate := "endDate_example" // string | The end date of sessions trend to query, for example '2021-11-08T12:00:00'.
    intervalLength := int32(56) // int32 | The minutes interval to query.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSessionsTrend(context.Background(), customerid, nameOrId).StartDate(startDate).EndDate(endDate).IntervalLength(intervalLength).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSessionsTrend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSessionsTrend`: SessionsTrendResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSessionsTrend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer Id. | 
**nameOrId** | **string** | The site name or ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSessionsTrendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **string** | The start date of sessions trend to query, for example &#39;2021-11-01T12:00:00&#39;. | 
 **endDate** | **string** | The end date of sessions trend to query, for example &#39;2021-11-08T12:00:00&#39;. | 
 **intervalLength** | **int32** | The minutes interval to query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetSite

> SiteDetailResponseModel SitesTPGetSite(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSite(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSite`: SiteDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetSiteLicense

> LicenseResponseModel SitesTPGetSiteLicense(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSiteLicense(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSiteLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSiteLicense`: LicenseResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSiteLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSiteLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetSiteLoadBalancingOption

> bool SitesTPGetSiteLoadBalancingOption(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the load balancing option for the site.

### Example

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSiteLoadBalancingOption(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSiteLoadBalancingOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSiteLoadBalancingOption`: bool
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSiteLoadBalancingOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSiteLoadBalancingOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetSiteSecurityKeyConfiguration

> SecurityKeyConfigurationResponseModel SitesTPGetSiteSecurityKeyConfiguration(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSiteSecurityKeyConfiguration(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSiteSecurityKeyConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSiteSecurityKeyConfiguration`: SecurityKeyConfigurationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSiteSecurityKeyConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSiteSecurityKeyConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetSiteSecurityKeyManagementConfig

> NameValueStringPairModel SitesTPGetSiteSecurityKeyManagementConfig(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSiteSecurityKeyManagementConfig(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSiteSecurityKeyManagementConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSiteSecurityKeyManagementConfig`: NameValueStringPairModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSiteSecurityKeyManagementConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSiteSecurityKeyManagementConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetSiteSettings

> SiteSettingsResponseModel SitesTPGetSiteSettings(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSiteSettings(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSiteSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSiteSettings`: SiteSettingsResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSiteSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSiteSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetSiteStatus

> SiteStatusResponseModel SitesTPGetSiteStatus(ctx, nameOrId, customerid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the status of a site.

### Example

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
    customerid := "customerid_example" // string | 
    async := true // bool | If `true`, the site status query will be executed as a background task. The task will have JobType GetSiteStatus. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSiteStatus(context.Background(), nameOrId, customerid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSiteStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSiteStatus`: SiteStatusResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSiteStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSiteStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, the site status query will be executed as a background task. The task will have JobType GetSiteStatus. When the task is complete it will redirect to GetJobResults. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SiteStatusResponseModel**](SiteStatusResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesTPGetSiteTestReport

> TestReportResponseModel SitesTPGetSiteTestReport(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the most recent test report.



### Example

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSiteTestReport(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSiteTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSiteTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSiteTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSiteTestReportRequest struct via the builder pattern


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


## SitesTPGetSiteValidLicenses

> LicenseResponseModelCollection SitesTPGetSiteValidLicenses(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSiteValidLicenses(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSiteValidLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSiteValidLicenses`: LicenseResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSiteValidLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSiteValidLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPGetSites

> SiteResponseModelCollection SitesTPGetSites(ctx, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the list of sites that are available to the customer and visible to the admin.

### Example

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
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetSites(context.Background(), customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetSites`: SiteResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SiteResponseModelCollection**](SiteResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesTPGetUpgradePackageVersions

> UpgradePackageVersionResponseModelCollection SitesTPGetUpgradePackageVersions(ctx, nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPGetUpgradePackageVersions(context.Background(), nameOrId, customerid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPGetUpgradePackageVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPGetUpgradePackageVersions`: UpgradePackageVersionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPGetUpgradePackageVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPGetUpgradePackageVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## SitesTPPatchCeipParticipation

> SitesTPPatchCeipParticipation(ctx, nameOrId, customerid).IsParticipating(isParticipating).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update the Customer Experience Improvement Program Participation for the site.

### Example

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
    isParticipating := true // bool | specify if participate Customer Experience Improvement Program.
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesTPApi.SitesTPPatchCeipParticipation(context.Background(), nameOrId, customerid).IsParticipating(isParticipating).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPPatchCeipParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPPatchCeipParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isParticipating** | **bool** | specify if participate Customer Experience Improvement Program. | 

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


## SitesTPPatchDnsResolution

> SitesTPPatchDnsResolution(ctx, nameOrId, customerid).Enable(enable).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update the broker site DNS resolution for the site.

### Example

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
    enable := true // bool | specify if enable DNS Resolution.
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesTPApi.SitesTPPatchDnsResolution(context.Background(), nameOrId, customerid).Enable(enable).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPPatchDnsResolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPPatchDnsResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enable** | **bool** | specify if enable DNS Resolution. | 

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


## SitesTPPatchMultipleRemotePCAssignments

> SitesTPPatchMultipleRemotePCAssignments(ctx, nameOrId, customerid).Allow(allow).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesTPApi.SitesTPPatchMultipleRemotePCAssignments(context.Background(), nameOrId, customerid).Allow(allow).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPPatchMultipleRemotePCAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPPatchMultipleRemotePCAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allow** | **bool** | specify if allow multi-user auto-assignment for Remote PC Access. | 

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


## SitesTPPatchSiteLicense

> SitesTPPatchSiteLicense(ctx, nameOrId, customerid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    request := *openapiclient.NewLicenseRequestModel() // LicenseRequestModel | The license update request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesTPApi.SitesTPPatchSiteLicense(context.Background(), nameOrId, customerid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPPatchSiteLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPPatchSiteLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**LicenseRequestModel**](LicenseRequestModel.md) | The license update request. | 
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


## SitesTPPatchSiteLoadBalancingOption

> SitesTPPatchSiteLoadBalancingOption(ctx, nameOrId, customerid).EnableVerticalLoadBalancing(enableVerticalLoadBalancing).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update the load balancing option for the site.



### Example

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
    enableVerticalLoadBalancing := true // bool | specify if using vertical load balancing
    customerid := "customerid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesTPApi.SitesTPPatchSiteLoadBalancingOption(context.Background(), nameOrId, customerid).EnableVerticalLoadBalancing(enableVerticalLoadBalancing).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPPatchSiteLoadBalancingOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPPatchSiteLoadBalancingOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableVerticalLoadBalancing** | **bool** | specify if using vertical load balancing | 

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


## SitesTPPatchSiteSecurityKeyConfiguration

> SitesTPPatchSiteSecurityKeyConfiguration(ctx, nameOrId, customerid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    request := *openapiclient.NewSecurityKeyConfigurationRequestModel() // SecurityKeyConfigurationRequestModel | Site security key configuration update request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesTPApi.SitesTPPatchSiteSecurityKeyConfiguration(context.Background(), nameOrId, customerid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPPatchSiteSecurityKeyConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPPatchSiteSecurityKeyConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**SecurityKeyConfigurationRequestModel**](SecurityKeyConfigurationRequestModel.md) | Site security key configuration update request. | 
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


## SitesTPPatchSiteSettings

> SitesTPPatchSiteSettings(ctx, nameOrId, customerid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    request := *openapiclient.NewEditSiteSettingsRequestModel() // EditSiteSettingsRequestModel | site settings request model.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesTPApi.SitesTPPatchSiteSettings(context.Background(), nameOrId, customerid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPPatchSiteSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPPatchSiteSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**EditSiteSettingsRequestModel**](EditSiteSettingsRequestModel.md) | site settings request model. | 
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


## SitesTPTestSite

> SiteTestResponseModel SitesTPTestSite(ctx, nameOrId, customerid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Run tests on a site and create a test report.



### Example

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
    customerid := "customerid_example" // string | 
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestSite. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the site being tested. * _Name_ - Name of the site being tested. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesTPApi.SitesTPTestSite(context.Background(), nameOrId, customerid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesTPApi.SitesTPTestSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTPTestSite`: SiteTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesTPApi.SitesTPTestSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 
**customerid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTPTestSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestSite. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the site being tested. * _Name_ - Name of the site being tested. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SiteTestResponseModel**](SiteTestResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

