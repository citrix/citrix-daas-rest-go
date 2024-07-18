# \SitesAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesCheckObjectNameExists**](SitesAPIsDAAS.md#SitesCheckObjectNameExists) | **Post** /$checkObjectNameExists/{objectType}/{nameOrPath} | Check for the existence of an object.
[**SitesGetMultipleRemotePCAssignments**](SitesAPIsDAAS.md#SitesGetMultipleRemotePCAssignments) | **Get** /Sites/{nameOrId}/MultipleRemotePCAssignments | Get multi-user auto-assignment for Remote PC Access.
[**SitesGetSessionsTrend**](SitesAPIsDAAS.md#SitesGetSessionsTrend) | **Get** /Sites/{nameOrId}/SessionsTrend | Get the sessions trend
[**SitesGetSite**](SitesAPIsDAAS.md#SitesGetSite) | **Get** /Sites/{nameOrId} | Get the details about a single site.
[**SitesGetSiteMisconfigurationReport**](SitesAPIsDAAS.md#SitesGetSiteMisconfigurationReport) | **Get** /Sites/{nameOrId}/MisconfigurationReport | Get the misconfiguration report.
[**SitesGetSiteSettings**](SitesAPIsDAAS.md#SitesGetSiteSettings) | **Get** /Sites/{nameOrId}/Settings | Get the settings for the site.
[**SitesGetSiteStatus**](SitesAPIsDAAS.md#SitesGetSiteStatus) | **Get** /Sites/{nameOrId}/Status | Get the status of a site.
[**SitesGetSiteTestReport**](SitesAPIsDAAS.md#SitesGetSiteTestReport) | **Get** /Sites/{nameOrId}/TestReport | Get the most recent test report.
[**SitesGetSites**](SitesAPIsDAAS.md#SitesGetSites) | **Get** /Sites | Get the list of sites that are available to the customer and visible to the admin.
[**SitesGetUpgradePackageVersions**](SitesAPIsDAAS.md#SitesGetUpgradePackageVersions) | **Get** /Sites/{nameOrId}/UpgradePackageVersions | Get the latest released VDA upgrade package versions in the site.
[**SitesPatchMultipleRemotePCAssignments**](SitesAPIsDAAS.md#SitesPatchMultipleRemotePCAssignments) | **Patch** /Sites/{nameOrId}/MultipleRemotePCAssignments | Update multi-user auto-assignment for Remote PC Access.
[**SitesPatchSiteSettings**](SitesAPIsDAAS.md#SitesPatchSiteSettings) | **Patch** /Sites/{nameOrId}/Settings | Update the broker site settings.
[**SitesTestSite**](SitesAPIsDAAS.md#SitesTestSite) | **Post** /Sites/{nameOrId}/$test | Run tests on a site and create a test report.



## SitesCheckObjectNameExists

> bool SitesCheckObjectNameExists(ctx, objectType, nameOrPath).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Check for the existence of an object.



### Example

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
    objectType := openapiclient.ObjectType("Unknown") // ObjectType | Type of the object to check.
    nameOrPath := "nameOrPath_example" // string | Name or path of the object to check. Currently the value should get rid of '\\\\' and replace it using '|'.             For instance, if the value is \"DomainA\\NameB\", the param will be \"DomainA|NameB\".
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesCheckObjectNameExists(context.Background(), objectType, nameOrPath).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesCheckObjectNameExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesCheckObjectNameExists`: bool
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesCheckObjectNameExists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | [**ObjectType**](.md) | Type of the object to check. | 
**nameOrPath** | **string** | Name or path of the object to check. Currently the value should get rid of &#39;\\\\&#39; and replace it using &#39;|&#39;.             For instance, if the value is \&quot;DomainA\\NameB\&quot;, the param will be \&quot;DomainA|NameB\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesCheckObjectNameExistsRequest struct via the builder pattern


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

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetMultipleRemotePCAssignments

> bool SitesGetMultipleRemotePCAssignments(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesGetMultipleRemotePCAssignments(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesGetMultipleRemotePCAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetMultipleRemotePCAssignments`: bool
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesGetMultipleRemotePCAssignments`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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


## SitesGetSessionsTrend

> SessionsTrendResponseModel SitesGetSessionsTrend(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).StartDate(startDate).EndDate(endDate).IntervalLength(intervalLength).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | The site name or ID.
    startDate := "startDate_example" // string | The start date of sessions trend to query, for example '2021-11-01T12:00:00'.
    endDate := "endDate_example" // string | The end date of sessions trend to query, for example '2021-11-08T12:00:00'.
    intervalLength := int32(56) // int32 | The minutes interval to query.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesGetSessionsTrend(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).StartDate(startDate).EndDate(endDate).IntervalLength(intervalLength).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesGetSessionsTrend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSessionsTrend`: SessionsTrendResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesGetSessionsTrend`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **startDate** | **string** | The start date of sessions trend to query, for example &#39;2021-11-01T12:00:00&#39;. | 
 **endDate** | **string** | The end date of sessions trend to query, for example &#39;2021-11-08T12:00:00&#39;. | 
 **intervalLength** | **int32** | The minutes interval to query. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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

> SiteDetailResponseModel SitesGetSite(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesGetSite(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesGetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSite`: SiteDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesGetSite`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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


## SitesGetSiteMisconfigurationReport

> string SitesGetSiteMisconfigurationReport(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the misconfiguration report.



### Example

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
    nameOrId := "nameOrId_example" // string | ID of the site.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesGetSiteMisconfigurationReport(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesGetSiteMisconfigurationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSiteMisconfigurationReport`: string
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesGetSiteMisconfigurationReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSiteMisconfigurationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetSiteSettings

> SiteSettingsResponseModel SitesGetSiteSettings(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesGetSiteSettings(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesGetSiteSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSiteSettings`: SiteSettingsResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesGetSiteSettings`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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


## SitesGetSiteStatus

> SiteStatusResponseModel SitesGetSiteStatus(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the site status query will be executed as a background task. The task will have JobType GetSiteStatus. When the task is complete it will redirect to GetJobResults. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesGetSiteStatus(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesGetSiteStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSiteStatus`: SiteStatusResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesGetSiteStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSiteStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the site status query will be executed as a background task. The task will have JobType GetSiteStatus. When the task is complete it will redirect to GetJobResults. | [default to false]

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


## SitesGetSiteTestReport

> TestReportResponseModel SitesGetSiteTestReport(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesGetSiteTestReport(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesGetSiteTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSiteTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesGetSiteTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSiteTestReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

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


## SitesGetSites

> SiteResponseModelCollection SitesGetSites(ctx).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesGetSites(context.Background()).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesGetSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetSites`: SiteResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesGetSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## SitesGetUpgradePackageVersions

> UpgradePackageVersionResponseModelCollection SitesGetUpgradePackageVersions(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesGetUpgradePackageVersions(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesGetUpgradePackageVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetUpgradePackageVersions`: UpgradePackageVersionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesGetUpgradePackageVersions`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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

> SitesPatchMultipleRemotePCAssignments(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Allow(allow).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    allow := true // bool | specify if allow multi-user auto-assignment for Remote PC Access.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesAPIsDAAS.SitesPatchMultipleRemotePCAssignments(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Allow(allow).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesPatchMultipleRemotePCAssignments``: %v\n", err)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **allow** | **bool** | specify if allow multi-user auto-assignment for Remote PC Access. | 
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


## SitesPatchSiteSettings

> SitesPatchSiteSettings(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).EditSiteSettingsRequestModel(editSiteSettingsRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    editSiteSettingsRequestModel := *openapiclient.NewEditSiteSettingsRequestModel() // EditSiteSettingsRequestModel | site settings request model.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesAPIsDAAS.SitesPatchSiteSettings(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).EditSiteSettingsRequestModel(editSiteSettingsRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesPatchSiteSettings``: %v\n", err)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **editSiteSettingsRequestModel** | [**EditSiteSettingsRequestModel**](EditSiteSettingsRequestModel.md) | site settings request model. | 
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


## SitesTestSite

> SiteTestResponseModel SitesTestSite(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    nameOrId := "nameOrId_example" // string | Name or ID of the site.
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestSite. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the site being tested. * _Name_ - Name of the site being tested. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPIsDAAS.SitesTestSite(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPIsDAAS.SitesTestSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesTestSite`: SiteTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `SitesAPIsDAAS.SitesTestSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesTestSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 

 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestSite. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the site being tested. * _Name_ - Name of the site being tested. | [default to false]

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

