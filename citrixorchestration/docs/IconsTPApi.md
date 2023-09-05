# \IconsTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IconsTPAddIcon**](IconsTPApi.md#IconsTPAddIcon) | **Post** /techpreview/{customerid}/{siteid}/Icons | add a customized icon.
[**IconsTPGetAppVServerPackageApplicationIcon**](IconsTPApi.md#IconsTPGetAppVServerPackageApplicationIcon) | **Post** /techpreview/{customerid}/{siteid}/Icons/AppVServerPackageApplication | Get the icon for a single App-V application within an App-V package on a server.
[**IconsTPGetIcon**](IconsTPApi.md#IconsTPGetIcon) | **Get** /techpreview/{customerid}/{siteid}/Icons/{id} | Get a single icon from the site.
[**IconsTPGetIcons**](IconsTPApi.md#IconsTPGetIcons) | **Get** /techpreview/{customerid}/{siteid}/Icons | Get all icons in the site.
[**IconsTPRemoveIcon**](IconsTPApi.md#IconsTPRemoveIcon) | **Delete** /techpreview/{customerid}/{siteid}/Icons/{id} | Remove a customized icon.



## IconsTPAddIcon

> IconResponseModel IconsTPAddIcon(ctx, customerid, siteid).RequestModel(requestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

add a customized icon.

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
    requestModel := *openapiclient.NewAddIconRequestModel("RawData_example") // AddIconRequestModel | request body containing icon data
    async := true // bool | If `true`, the icon will be added as a background task. The task will have JobType AddIcon. When the task is complete it will redirect to GetIcon.              (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IconsTPApi.IconsTPAddIcon(context.Background(), customerid, siteid).RequestModel(requestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsTPApi.IconsTPAddIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IconsTPAddIcon`: IconResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IconsTPApi.IconsTPAddIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIconsTPAddIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestModel** | [**AddIconRequestModel**](AddIconRequestModel.md) | request body containing icon data | 
 **async** | **bool** | If &#x60;true&#x60;, the icon will be added as a background task. The task will have JobType AddIcon. When the task is complete it will redirect to GetIcon.              | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IconResponseModel**](IconResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IconsTPGetAppVServerPackageApplicationIcon

> IconDataResponseModel IconsTPGetAppVServerPackageApplicationIcon(ctx, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the icon for a single App-V application within an App-V package on a server.



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
    request := *openapiclient.NewAppvServerPackageApplicationIconRequestModel("Server_example", "PackageId_example", "ApplicationId_example") // AppvServerPackageApplicationIconRequestModel | App-V server package application icon request model.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IconsTPApi.IconsTPGetAppVServerPackageApplicationIcon(context.Background(), customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsTPApi.IconsTPGetAppVServerPackageApplicationIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IconsTPGetAppVServerPackageApplicationIcon`: IconDataResponseModel
    fmt.Fprintf(os.Stdout, "Response from `IconsTPApi.IconsTPGetAppVServerPackageApplicationIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIconsTPGetAppVServerPackageApplicationIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**AppvServerPackageApplicationIconRequestModel**](AppvServerPackageApplicationIconRequestModel.md) | App-V server package application icon request model. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IconDataResponseModel**](IconDataResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IconsTPGetIcon

> map[string]map[string]interface{} IconsTPGetIcon(ctx, id, customerid, siteid).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get a single icon from the site.



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
    id := "id_example" // string | ID of the icon.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    iconFormat := "iconFormat_example" // string | Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`   where: * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.  example: `\"image/png;32x32x24\"`   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IconsTPApi.IconsTPGetIcon(context.Background(), id, customerid, siteid).IconFormat(iconFormat).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsTPApi.IconsTPGetIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IconsTPGetIcon`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IconsTPApi.IconsTPGetIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the icon. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIconsTPGetIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **iconFormat** | **string** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;   where: * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.  example: &#x60;\&quot;image/png;32x32x24\&quot;&#x60;   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IconsTPGetIcons

> IconResponseModelCollection IconsTPGetIcons(ctx, customerid, siteid).IconFormat(iconFormat).BuiltIn(builtIn).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all icons in the site.



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
    iconFormat := "iconFormat_example" // string | Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`   where:   * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.   example: `\"image/png;32x32x24\"`   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. (optional)
    builtIn := true // bool | If specified as `true`, only built-in icons will be returned.  If specified as `false`, only user-created icons will be returned.  If not specified, all icons will be returned. (optional)
    async := true // bool | If async execute.  (optional) (default to false)
    limit := int32(56) // int32 | The max number of icons returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IconsTPApi.IconsTPGetIcons(context.Background(), customerid, siteid).IconFormat(iconFormat).BuiltIn(builtIn).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsTPApi.IconsTPGetIcons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IconsTPGetIcons`: IconResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `IconsTPApi.IconsTPGetIcons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIconsTPGetIconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iconFormat** | **string** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;   where:   * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.   example: &#x60;\&quot;image/png;32x32x24\&quot;&#x60;   Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. | 
 **builtIn** | **bool** | If specified as &#x60;true&#x60;, only built-in icons will be returned.  If specified as &#x60;false&#x60;, only user-created icons will be returned.  If not specified, all icons will be returned. | 
 **async** | **bool** | If async execute.  | [default to false]
 **limit** | **int32** | The max number of icons returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**IconResponseModelCollection**](IconResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IconsTPRemoveIcon

> IconsTPRemoveIcon(ctx, id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Remove a customized icon.

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
    id := int32(56) // int32 | id of the icon to remove
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the icon will be removed as a background task. The task will have JobType RemoveIcon. When the task is complete it will redirect to GetIcons.              (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IconsTPApi.IconsTPRemoveIcon(context.Background(), id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IconsTPApi.IconsTPRemoveIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id of the icon to remove | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIconsTPRemoveIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the icon will be removed as a background task. The task will have JobType RemoveIcon. When the task is complete it will redirect to GetIcons.              | [default to false]
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

