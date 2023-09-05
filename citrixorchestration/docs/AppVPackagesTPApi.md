# \AppVPackagesTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppVPackagesTPGetAppVPackage**](AppVPackagesTPApi.md#AppVPackagesTPGetAppVPackage) | **Get** /techpreview/{customerid}/{siteid}/AppVPackages/{id} | Get the details for a single App-V package within the site
[**AppVPackagesTPGetAppVPackageApplication**](AppVPackagesTPApi.md#AppVPackagesTPGetAppVPackageApplication) | **Get** /techpreview/{customerid}/{siteid}/AppVPackages/{id}/Applications/{appId} | Get details for a single App-V application within an App-V package
[**AppVPackagesTPGetAppVPackageApplicationIcon**](AppVPackagesTPApi.md#AppVPackagesTPGetAppVPackageApplicationIcon) | **Get** /techpreview/{customerid}/{siteid}/AppVPackages/{id}/Applications/{appId}/Icon | Get the icon for a single App-V application within an App-V package within the site.
[**AppVPackagesTPGetAppVPackageApplications**](AppVPackagesTPApi.md#AppVPackagesTPGetAppVPackageApplications) | **Get** /techpreview/{customerid}/{siteid}/AppVPackages/{id}/Applications | Get App-V applications within an App-V package
[**AppVPackagesTPGetAppVPackageBrokerApplications**](AppVPackagesTPApi.md#AppVPackagesTPGetAppVPackageBrokerApplications) | **Get** /techpreview/{customerid}/{siteid}/AppVPackages/{id}/BrokerApplications | Get Broker applications delivered from the App-V package
[**AppVPackagesTPGetAppVPackageDeliveryGroups**](AppVPackagesTPApi.md#AppVPackagesTPGetAppVPackageDeliveryGroups) | **Get** /techpreview/{customerid}/{siteid}/AppVPackages/{id}/DeliveryGroups | Get delivery groups which contain applications in the App-V package
[**AppVPackagesTPGetAppVPackageIcon**](AppVPackagesTPApi.md#AppVPackagesTPGetAppVPackageIcon) | **Get** /techpreview/{customerid}/{siteid}/AppVPackages/{id}/Icon | Get the icon for a single App-V package within the site
[**AppVPackagesTPGetAppVPackages**](AppVPackagesTPApi.md#AppVPackagesTPGetAppVPackages) | **Get** /techpreview/{customerid}/{siteid}/AppVPackages | Get the App-V packages configured in the site
[**AppVPackagesTPImportAppVPackages**](AppVPackagesTPApi.md#AppVPackagesTPImportAppVPackages) | **Post** /techpreview/{customerid}/{siteid}/AppVPackages | Import App-V packages to the site
[**AppVPackagesTPRemoveAppVPackage**](AppVPackagesTPApi.md#AppVPackagesTPRemoveAppVPackage) | **Delete** /techpreview/{customerid}/{siteid}/AppVPackages/{id} | Remove a single App-V package within the site



## AppVPackagesTPGetAppVPackage

> AppVPackageResponseModel AppVPackagesTPGetAppVPackage(ctx, id, customerid, siteid).LibraryUid(libraryUid).VersionId(versionId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    libraryUid := int32(56) // int32 | ID of the library where the package is present. (optional) (default to 0)
    versionId := "versionId_example" // string | Package version guid. If not specified, return the first             package with id. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPGetAppVPackage(context.Background(), id, customerid, siteid).LibraryUid(libraryUid).VersionId(versionId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPGetAppVPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesTPGetAppVPackage`: AppVPackageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesTPApi.AppVPackagesTPGetAppVPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesTPGetAppVPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **libraryUid** | **int32** | ID of the library where the package is present. | [default to 0]
 **versionId** | **string** | Package version guid. If not specified, return the first             package with id. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVPackagesTPGetAppVPackageApplication

> AppVApplicationDetailResponseModel AppVPackagesTPGetAppVPackageApplication(ctx, id, appId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPGetAppVPackageApplication(context.Background(), id, appId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesTPGetAppVPackageApplication`: AppVApplicationDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**appId** | **string** | ID of the App-V application within the package. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesTPGetAppVPackageApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVPackagesTPGetAppVPackageApplicationIcon

> IconDataResponseModel AppVPackagesTPGetAppVPackageApplicationIcon(ctx, id, appId, customerid, siteid).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    iconFormat := "iconFormat_example" // string | Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`  where:   * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPGetAppVPackageApplicationIcon(context.Background(), id, appId, customerid, siteid).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageApplicationIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesTPGetAppVPackageApplicationIcon`: IconDataResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageApplicationIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**appId** | **string** | ID of the App-V application within the package. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesTPGetAppVPackageApplicationIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **iconFormat** | **string** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;  where:   * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVPackagesTPGetAppVPackageApplications

> AppVApplicationResponseModelCollection AppVPackagesTPGetAppVPackageApplications(ctx, id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPGetAppVPackageApplications(context.Background(), id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesTPGetAppVPackageApplications`: AppVApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesTPGetAppVPackageApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVPackagesTPGetAppVPackageBrokerApplications

> ApplicationResponseModelCollection AppVPackagesTPGetAppVPackageBrokerApplications(ctx, id, customerid, siteid).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    libraryUid := int32(56) // int32 | ID of the library where the package is present. (optional) (default to 0)
    versionId := "versionId_example" // string | Package version Id. (optional)
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | The required fields. (optional)
    async := true // bool | If `true`, the applications will be fetched as a background task. The task will have JobType GetAppVPackageBrokerApplications. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPGetAppVPackageBrokerApplications(context.Background(), id, customerid, siteid).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageBrokerApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesTPGetAppVPackageBrokerApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageBrokerApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesTPGetAppVPackageBrokerApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **libraryUid** | **int32** | ID of the library where the package is present. | [default to 0]
 **versionId** | **string** | Package version Id. | 
 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | The required fields. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications will be fetched as a background task. The task will have JobType GetAppVPackageBrokerApplications. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVPackagesTPGetAppVPackageDeliveryGroups

> DeliveryGroupResponseModelCollection AppVPackagesTPGetAppVPackageDeliveryGroups(ctx, id, customerid, siteid).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    libraryUid := int32(56) // int32 | ID of the library where the package is present. (optional) (default to 0)
    versionId := "versionId_example" // string | Package version Id. If not specified, all delivery groups             that associated with id will be fetched. (optional)
    limit := int32(56) // int32 | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | The required fields. (optional)
    async := true // bool | If `true`, the delivery groups will be fetched as a background task. The task will have JobType GetAppVPackageDelveryGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPGetAppVPackageDeliveryGroups(context.Background(), id, customerid, siteid).LibraryUid(libraryUid).VersionId(versionId).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageDeliveryGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesTPGetAppVPackageDeliveryGroups`: DeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageDeliveryGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesTPGetAppVPackageDeliveryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **libraryUid** | **int32** | ID of the library where the package is present. | [default to 0]
 **versionId** | **string** | Package version Id. If not specified, all delivery groups             that associated with id will be fetched. | 
 **limit** | **int32** | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | The required fields. | 
 **async** | **bool** | If &#x60;true&#x60;, the delivery groups will be fetched as a background task. The task will have JobType GetAppVPackageDelveryGroups. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVPackagesTPGetAppVPackageIcon

> IconDataResponseModel AppVPackagesTPGetAppVPackageIcon(ctx, id, customerid, siteid).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    iconFormat := "iconFormat_example" // string | Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`  where:   * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPGetAppVPackageIcon(context.Background(), id, customerid, siteid).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesTPGetAppVPackageIcon`: IconDataResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesTPApi.AppVPackagesTPGetAppVPackageIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesTPGetAppVPackageIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **iconFormat** | **string** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;  where:   * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVPackagesTPGetAppVPackages

> AppVPackageResponseModelCollection AppVPackagesTPGetAppVPackages(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the packages will be fetched as a background task. The task will have JobType GetAppVPackages. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPGetAppVPackages(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPGetAppVPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVPackagesTPGetAppVPackages`: AppVPackageResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVPackagesTPApi.AppVPackagesTPGetAppVPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesTPGetAppVPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, the packages will be fetched as a background task. The task will have JobType GetAppVPackages. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVPackagesTPImportAppVPackages

> AppVPackagesTPImportAppVPackages(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewAppVPackageRequestModel() // AppVPackageRequestModel | Request model that includes packages path.
    async := true // bool | If `true`, the packages will be imported as a background task. The task will have JobType ImportAppVPackages. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPImportAppVPackages(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPImportAppVPackages``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppVPackagesTPImportAppVPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**AppVPackageRequestModel**](AppVPackageRequestModel.md) | Request model that includes packages path. | 
 **async** | **bool** | If &#x60;true&#x60;, the packages will be imported as a background task. The task will have JobType ImportAppVPackages. | [default to false]
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


## AppVPackagesTPRemoveAppVPackage

> AppVPackagesTPRemoveAppVPackage(ctx, id, customerid, siteid).VersionId(versionId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    versionId := "versionId_example" // string | Package version guid.             If not specified, the first package with the id will be removed. (optional)
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVPackages. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVPackagesTPApi.AppVPackagesTPRemoveAppVPackage(context.Background(), id, customerid, siteid).VersionId(versionId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVPackagesTPApi.AppVPackagesTPRemoveAppVPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the App-V package. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVPackagesTPRemoveAppVPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **versionId** | **string** | Package version guid.             If not specified, the first package with the id will be removed. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVPackages. | [default to false]
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

