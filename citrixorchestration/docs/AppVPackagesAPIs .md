# \AppVPackagesAPIs 

All URIs are relative to *https://api-us.cloud.com/cvad/manage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppVPackagesGetAppVPackage**](AppVPackagesAPIs .md#AppVPackagesGetAppVPackage) | **Get** /AppVPackages/{id} | Get the details for a single App-V package within the site
[**AppVPackagesGetAppVPackageApplication**](AppVPackagesAPIs .md#AppVPackagesGetAppVPackageApplication) | **Get** /AppVPackages/{id}/Applications/{appId} | Get details for a single App-V application within an App-V package
[**AppVPackagesGetAppVPackageApplicationIcon**](AppVPackagesAPIs .md#AppVPackagesGetAppVPackageApplicationIcon) | **Get** /AppVPackages/{id}/Applications/{appId}/Icon | Get the icon for a single App-V application within an App-V package within the site.
[**AppVPackagesGetAppVPackageApplications**](AppVPackagesAPIs .md#AppVPackagesGetAppVPackageApplications) | **Get** /AppVPackages/{id}/Applications | Get App-V applications within an App-V package
[**AppVPackagesGetAppVPackageBrokerApplications**](AppVPackagesAPIs .md#AppVPackagesGetAppVPackageBrokerApplications) | **Get** /AppVPackages/{id}/BrokerApplications | Get Broker applications delivered from the App-V package
[**AppVPackagesGetAppVPackageDeliveryGroups**](AppVPackagesAPIs .md#AppVPackagesGetAppVPackageDeliveryGroups) | **Get** /AppVPackages/{id}/DeliveryGroups | Get delivery groups which contain applications in the App-V package
[**AppVPackagesGetAppVPackageIcon**](AppVPackagesAPIs .md#AppVPackagesGetAppVPackageIcon) | **Get** /AppVPackages/{id}/Icon | Get the icon for a single App-V package within the site
[**AppVPackagesGetAppVPackages**](AppVPackagesAPIs .md#AppVPackagesGetAppVPackages) | **Get** /AppVPackages | Get the App-V packages configured in the site
[**AppVPackagesImportAppVPackages**](AppVPackagesAPIs .md#AppVPackagesImportAppVPackages) | **Post** /AppVPackages | Import App-V packages to the site
[**AppVPackagesRemoveAppVPackage**](AppVPackagesAPIs .md#AppVPackagesRemoveAppVPackage) | **Delete** /AppVPackages/{id} | Remove a single App-V package within the site



## AppVPackagesGetAppVPackage

> AppVPackageResponseModel AppVPackagesGetAppVPackage(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).LibraryUid(libraryUid).VersionId(versionId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the details for a single App-V package within the site



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    id := "id_example" // string | ID of the App-V package.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    libraryUid := int32(56) // int32 | ID of the library where the package is present. (optional) (default to 0)
    versionId := "versionId_example" // string | Package version guid. If not specified, return the first             package with id. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIs .AppVPackagesGetAppVPackage(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).LibraryUid(libraryUid).VersionId(versionId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesGetAppVPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackage`: AppVPackageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIs .AppVPackagesGetAppVPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetAppVPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **libraryUid** | **int32** | ID of the library where the package is present. | [default to 0]
 **versionId** | **string** | Package version guid. If not specified, return the first             package with id. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AppVPackageResponseModel**](AppVPackageResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVPackagesGetAppVPackageApplication

> AppVApplicationDetailResponseModel AppVPackagesGetAppVPackageApplication(ctx, id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get details for a single App-V application within an App-V package



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    id := "id_example" // string | ID of the App-V package.
    appId := "appId_example" // string | ID of the App-V application within the package.
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
    resp, r, err := apiClient.AppVPackagesAPIs .AppVPackagesGetAppVPackageApplication(context.Background(), id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesGetAppVPackageApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageApplication`: AppVApplicationDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIs .AppVPackagesGetAppVPackageApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**appId** | **string** | ID of the App-V application within the package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetAppVPackageApplicationRequest struct via the builder pattern


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

[**AppVApplicationDetailResponseModel**](AppVApplicationDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVPackagesGetAppVPackageApplicationIcon

> IconDataResponseModel AppVPackagesGetAppVPackageApplicationIcon(ctx, id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the icon for a single App-V application within an App-V package within the site.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    id := "id_example" // string | ID of the App-V package.
    appId := "appId_example" // string | ID of the App-V application within the package.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    iconFormat := "iconFormat_example" // string | Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`  where:   * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIs .AppVPackagesGetAppVPackageApplicationIcon(context.Background(), id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesGetAppVPackageApplicationIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageApplicationIcon`: IconDataResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIs .AppVPackagesGetAppVPackageApplicationIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**appId** | **string** | ID of the App-V application within the package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetAppVPackageApplicationIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **iconFormat** | **string** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;  where:   * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**IconDataResponseModel**](IconDataResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVPackagesGetAppVPackageApplications

> AppVApplicationResponseModelCollection AppVPackagesGetAppVPackageApplications(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get App-V applications within an App-V package



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    id := "id_example" // string | ID of the App-V package.
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
    resp, r, err := apiClient.AppVPackagesAPIs .AppVPackagesGetAppVPackageApplications(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesGetAppVPackageApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageApplications`: AppVApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIs .AppVPackagesGetAppVPackageApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetAppVPackageApplicationsRequest struct via the builder pattern


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

[**AppVApplicationResponseModelCollection**](AppVApplicationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVPackagesGetAppVPackageBrokerApplications

> ApplicationResponseModelCollection AppVPackagesGetAppVPackageBrokerApplications(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get Broker applications delivered from the App-V package



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    id := "id_example" // string | ID of the App-V package.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    libraryUid := int32(56) // int32 | ID of the library where the package is present. (optional) (default to 0)
    versionId := "versionId_example" // string | Package version Id. (optional)
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Id,Uid,ApplicationFolder" // string | The required fields. (optional)
    async := true // bool | If `true`, the applications will be fetched as a background task. The task will have JobType GetAppVPackageBrokerApplications. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIs .AppVPackagesGetAppVPackageBrokerApplications(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesGetAppVPackageBrokerApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageBrokerApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIs .AppVPackagesGetAppVPackageBrokerApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetAppVPackageBrokerApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **libraryUid** | **int32** | ID of the library where the package is present. | [default to 0]
 **versionId** | **string** | Package version Id. | 
 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | The required fields. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications will be fetched as a background task. The task will have JobType GetAppVPackageBrokerApplications. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## AppVPackagesGetAppVPackageDeliveryGroups

> DeliveryGroupResponseModelCollection AppVPackagesGetAppVPackageDeliveryGroups(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get delivery groups which contain applications in the App-V package



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    id := "id_example" // string | ID of the App-V package.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    libraryUid := int32(56) // int32 | ID of the library where the package is present. (optional) (default to 0)
    versionId := "versionId_example" // string | Package version Id. If not specified, all delivery groups             that associated with id will be fetched. (optional)
    limit := int32(56) // int32 | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Id,Uid,UserManagement" // string | The required fields. (optional)
    async := true // bool | If `true`, the delivery groups will be fetched as a background task. The task will have JobType GetAppVPackageDelveryGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIs .AppVPackagesGetAppVPackageDeliveryGroups(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesGetAppVPackageDeliveryGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageDeliveryGroups`: DeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIs .AppVPackagesGetAppVPackageDeliveryGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetAppVPackageDeliveryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **libraryUid** | **int32** | ID of the library where the package is present. | [default to 0]
 **versionId** | **string** | Package version Id. If not specified, all delivery groups             that associated with id will be fetched. | 
 **limit** | **int32** | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | The required fields. | 
 **async** | **bool** | If &#x60;true&#x60;, the delivery groups will be fetched as a background task. The task will have JobType GetAppVPackageDelveryGroups. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## AppVPackagesGetAppVPackageIcon

> IconDataResponseModel AppVPackagesGetAppVPackageIcon(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the icon for a single App-V package within the site



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    id := "id_example" // string | ID of the App-V package.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    iconFormat := "iconFormat_example" // string | Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`  where:   * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIs .AppVPackagesGetAppVPackageIcon(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesGetAppVPackageIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageIcon`: IconDataResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIs .AppVPackagesGetAppVPackageIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetAppVPackageIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **iconFormat** | **string** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;  where:   * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**IconDataResponseModel**](IconDataResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVPackagesGetAppVPackages

> AppVPackageResponseModelCollection AppVPackagesGetAppVPackages(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the App-V packages configured in the site



### Example

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
    async := true // bool | If `true`, the packages will be fetched as a background task. The task will have JobType GetAppVPackages. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIs .AppVPackagesGetAppVPackages(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesGetAppVPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackages`: AppVPackageResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIs .AppVPackagesGetAppVPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetAppVPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the packages will be fetched as a background task. The task will have JobType GetAppVPackages. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AppVPackageResponseModelCollection**](AppVPackageResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVPackagesImportAppVPackages

> AppVPackagesImportAppVPackages(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AppVPackageRequestModel(appVPackageRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Import App-V packages to the site



### Example

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
    appVPackageRequestModel := *openapiclient.NewAppVPackageRequestModel() // AppVPackageRequestModel | Request model that includes packages path.
    async := true // bool | If `true`, the packages will be imported as a background task. The task will have JobType ImportAppVPackages. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVPackagesAPIs .AppVPackagesImportAppVPackages(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AppVPackageRequestModel(appVPackageRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesImportAppVPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesImportAppVPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **appVPackageRequestModel** | [**AppVPackageRequestModel**](AppVPackageRequestModel.md) | Request model that includes packages path. | 
 **async** | **bool** | If &#x60;true&#x60;, the packages will be imported as a background task. The task will have JobType ImportAppVPackages. | [default to false]
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


## AppVPackagesRemoveAppVPackage

> AppVPackagesRemoveAppVPackage(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).VersionId(versionId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Remove a single App-V package within the site



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    id := "id_example" // string | ID of the App-V package.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    versionId := "versionId_example" // string | Package version guid.             If not specified, the first package with the id will be removed. (optional)
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVPackages. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVPackagesAPIs .AppVPackagesRemoveAppVPackage(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).VersionId(versionId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIs .AppVPackagesRemoveAppVPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesRemoveAppVPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **versionId** | **string** | Package version guid.             If not specified, the first package with the id will be removed. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVPackages. | [default to false]
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

