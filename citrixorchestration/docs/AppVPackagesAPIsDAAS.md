# \AppVPackagesAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppVPackagesGetAppVPackage**](AppVPackagesAPIsDAAS.md#AppVPackagesGetAppVPackage) | **Get** /AppVPackages/{id} | Get the details for a single App-V package within the site
[**AppVPackagesGetAppVPackageApplication**](AppVPackagesAPIsDAAS.md#AppVPackagesGetAppVPackageApplication) | **Get** /AppVPackages/{id}/Applications/{appId} | Get details for a single App-V application within an App-V package
[**AppVPackagesGetAppVPackageApplicationIcon**](AppVPackagesAPIsDAAS.md#AppVPackagesGetAppVPackageApplicationIcon) | **Get** /AppVPackages/{id}/Applications/{appId}/Icon | Get the icon for a single App-V application within an App-V package within the site.
[**AppVPackagesGetAppVPackageApplications**](AppVPackagesAPIsDAAS.md#AppVPackagesGetAppVPackageApplications) | **Get** /AppVPackages/{id}/Applications | Get App-V applications within an App-V package
[**AppVPackagesGetAppVPackageBrokerApplications**](AppVPackagesAPIsDAAS.md#AppVPackagesGetAppVPackageBrokerApplications) | **Get** /AppVPackages/{id}/BrokerApplications | Get Broker applications delivered from the App-V package
[**AppVPackagesGetAppVPackageDeliveryGroups**](AppVPackagesAPIsDAAS.md#AppVPackagesGetAppVPackageDeliveryGroups) | **Get** /AppVPackages/{id}/DeliveryGroups | Get delivery groups which contain applications in the App-V package
[**AppVPackagesGetAppVPackageIcon**](AppVPackagesAPIsDAAS.md#AppVPackagesGetAppVPackageIcon) | **Get** /AppVPackages/{id}/Icon | Get the icon for a single App-V package within the site
[**AppVPackagesGetAppVPackages**](AppVPackagesAPIsDAAS.md#AppVPackagesGetAppVPackages) | **Get** /AppVPackages | Get the App-V packages configured in the site
[**AppVPackagesGetPackageApplicationFileTypes**](AppVPackagesAPIsDAAS.md#AppVPackagesGetPackageApplicationFileTypes) | **Get** /AppVPackages/{id}/Applications/{appId}/FileTypes | Get the fileTypes for an application within a package within the site.
[**AppVPackagesImportAppVPackages**](AppVPackagesAPIsDAAS.md#AppVPackagesImportAppVPackages) | **Post** /AppVPackages | Import App-V packages to the site
[**AppVPackagesRemoveAppVPackage**](AppVPackagesAPIsDAAS.md#AppVPackagesRemoveAppVPackage) | **Delete** /AppVPackages/{id} | Remove a single App-V package within the site



## AppVPackagesGetAppVPackage

> AppVPackageResponseModel AppVPackagesGetAppVPackage(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).LibraryUid(libraryUid).VersionId(versionId).Fields(fields).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    id := "id_example" // string | ID of the App-V package.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    libraryUid := int32(56) // int32 | ID of the library where the package is present. (optional) (default to 0)
    versionId := "versionId_example" // string | Package version guid. If not specified, return the first             package with id. (optional)
    fields := "AppVApplications,Description,Exists" // string | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma','.              (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackage(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).LibraryUid(libraryUid).VersionId(versionId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackage`: AppVPackageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackage`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **libraryUid** | **int32** | ID of the library where the package is present. | [default to 0]
 **versionId** | **string** | Package version guid. If not specified, return the first             package with id. | 
 **fields** | **string** | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma&#39;,&#39;.              | 

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

> AppVApplicationDetailResponseModel AppVPackagesGetAppVPackageApplication(ctx, id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    id := "id_example" // string | ID of the App-V package.
    appId := "appId_example" // string | ID of the App-V application within the package.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageApplication(context.Background(), id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageApplication`: AppVApplicationDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageApplication`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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

> IconDataResponseModel AppVPackagesGetAppVPackageApplicationIcon(ctx, id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).IconFormat(iconFormat).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    id := "id_example" // string | ID of the App-V package.
    appId := "appId_example" // string | ID of the App-V application within the package.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    iconFormat := "iconFormat_example" // string | Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`  where:   * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageApplicationIcon(context.Background(), id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).IconFormat(iconFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageApplicationIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageApplicationIcon`: IconDataResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageApplicationIcon`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **iconFormat** | **string** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;  where:   * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. | 

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

> AppVApplicationResponseModelCollection AppVPackagesGetAppVPackageApplications(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    id := "id_example" // string | ID of the App-V package.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageApplications(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageApplications`: AppVApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageApplications`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
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

> ApplicationResponseModelCollection AppVPackagesGetAppVPackageBrokerApplications(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    id := "id_example" // string | ID of the App-V package.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    libraryUid := int32(56) // int32 | ID of the library where the package is present. (optional) (default to 0)
    versionId := "versionId_example" // string | Package version Id. (optional)
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Id,Uid,ApplicationFolder" // string | The required fields. (optional)
    async := true // bool | If `true`, the applications will be fetched as a background task. The task will have JobType GetAppVPackageBrokerApplications. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageBrokerApplications(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageBrokerApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageBrokerApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageBrokerApplications`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **libraryUid** | **int32** | ID of the library where the package is present. | [default to 0]
 **versionId** | **string** | Package version Id. | 
 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | The required fields. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications will be fetched as a background task. The task will have JobType GetAppVPackageBrokerApplications. | [default to false]

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

> DeliveryGroupResponseModelCollection AppVPackagesGetAppVPackageDeliveryGroups(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    id := "id_example" // string | ID of the App-V package.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    libraryUid := int32(56) // int32 | ID of the library where the package is present. (optional) (default to 0)
    versionId := "versionId_example" // string | Package version Id. If not specified, all delivery groups             that associated with id will be fetched. (optional)
    limit := int32(56) // int32 | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Id,Uid,UserManagement" // string | The required fields. (optional)
    async := true // bool | If `true`, the delivery groups will be fetched as a background task. The task will have JobType GetAppVPackageDelveryGroups. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageDeliveryGroups(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageDeliveryGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageDeliveryGroups`: DeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageDeliveryGroups`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **libraryUid** | **int32** | ID of the library where the package is present. | [default to 0]
 **versionId** | **string** | Package version Id. If not specified, all delivery groups             that associated with id will be fetched. | 
 **limit** | **int32** | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | The required fields. | 
 **async** | **bool** | If &#x60;true&#x60;, the delivery groups will be fetched as a background task. The task will have JobType GetAppVPackageDelveryGroups. | [default to false]

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

> IconDataResponseModel AppVPackagesGetAppVPackageIcon(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).IconFormat(iconFormat).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    id := "id_example" // string | ID of the App-V package.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    iconFormat := "iconFormat_example" // string | Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`  where:   * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageIcon(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).IconFormat(iconFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackageIcon`: IconDataResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackageIcon`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **iconFormat** | **string** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;  where:   * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. | 

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

> AppVPackageResponseModelCollection AppVPackagesGetAppVPackages(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Fields(fields).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the packages will be fetched as a background task. The task will have JobType GetAppVPackages. (optional) (default to false)
    fields := "AppVApplications,Description,Exists" // string | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma','.              (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackages(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetAppVPackages`: AppVPackageResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIsDAAS.AppVPackagesGetAppVPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetAppVPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the packages will be fetched as a background task. The task will have JobType GetAppVPackages. | [default to false]
 **fields** | **string** | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma&#39;,&#39;.              | 

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


## AppVPackagesGetPackageApplicationFileTypes

> FtaResponseModelCollection AppVPackagesGetPackageApplicationFileTypes(ctx, id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).LibraryUid(libraryUid).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get the fileTypes for an application within a package within the site.



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
    id := "id_example" // string | ID of the package.
    appId := "appId_example" // string | Identifier of the application within the package.
    libraryUid := int32(56) // int32 | ID of the library where the package is present.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesGetPackageApplicationFileTypes(context.Background(), id, appId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).LibraryUid(libraryUid).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesGetPackageApplicationFileTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesGetPackageApplicationFileTypes`: FtaResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesAPIsDAAS.AppVPackagesGetPackageApplicationFileTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the package. | 
**appId** | **string** | Identifier of the application within the package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesGetPackageApplicationFileTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **libraryUid** | **int32** | ID of the library where the package is present. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** |  | [default to false]

### Return type

[**FtaResponseModelCollection**](FtaResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVPackagesImportAppVPackages

> AppVPackagesImportAppVPackages(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).AppVPackageRequestModel(appVPackageRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    appVPackageRequestModel := *openapiclient.NewAppVPackageRequestModel() // AppVPackageRequestModel | Request model that includes packages path.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the packages will be imported as a background task. The task will have JobType ImportAppVPackages. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesImportAppVPackages(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).AppVPackageRequestModel(appVPackageRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesImportAppVPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesImportAppVPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **appVPackageRequestModel** | [**AppVPackageRequestModel**](AppVPackageRequestModel.md) | Request model that includes packages path. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the packages will be imported as a background task. The task will have JobType ImportAppVPackages. | [default to false]

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

> AppVPackagesRemoveAppVPackage(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).VersionId(versionId).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    id := "id_example" // string | ID of the App-V package.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    versionId := "versionId_example" // string | Package version guid.             If not specified, the first package with the id will be removed. (optional)
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVPackages. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVPackagesAPIsDAAS.AppVPackagesRemoveAppVPackage(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).VersionId(versionId).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesAPIsDAAS.AppVPackagesRemoveAppVPackage``: %v\n", err)
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **versionId** | **string** | Package version guid.             If not specified, the first package with the id will be removed. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVPackages. | [default to false]

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

