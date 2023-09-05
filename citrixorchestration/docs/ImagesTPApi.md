# \ImagesTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesTPAddImageVersion**](ImagesTPApi.md#ImagesTPAddImageVersion) | **Post** /techpreview/{customerid}/{siteid}/Images/Definitions/{nameOrId}/Versions | Add an image version to a specific image definition
[**ImagesTPCheckImageDefinitionExists**](ImagesTPApi.md#ImagesTPCheckImageDefinitionExists) | **Head** /techpreview/{customerid}/{siteid}/Images/Definitions/{name} | Check for the existence of an image definition by name.
[**ImagesTPGetCatalogsByImageVersion**](ImagesTPApi.md#ImagesTPGetCatalogsByImageVersion) | **Get** /techpreview/{customerid}/{siteid}/Images/Versions/{versionId}/Catalogs | Get catalogs use the specific image version.
[**ImagesTPGetImageDefinition**](ImagesTPApi.md#ImagesTPGetImageDefinition) | **Get** /techpreview/{customerid}/{siteid}/Images/Definitions/{nameOrId} | Get a specified image definition
[**ImagesTPGetImageDefinitionImageVersion**](ImagesTPApi.md#ImagesTPGetImageDefinitionImageVersion) | **Get** /techpreview/{customerid}/{siteid}/Images/Definitions/{nameOrId}/Versions/{versionNameOrId} | Get details of a specified image version
[**ImagesTPGetImageDefinitions**](ImagesTPApi.md#ImagesTPGetImageDefinitions) | **Get** /techpreview/{customerid}/{siteid}/Images/Definitions | Get all defined image definitions
[**ImagesTPGetImageVersion**](ImagesTPApi.md#ImagesTPGetImageVersion) | **Get** /techpreview/{customerid}/{siteid}/Images/Versions/{id} | Get details of an image version specified by Id
[**ImagesTPGetImageVersionsByDefinition**](ImagesTPApi.md#ImagesTPGetImageVersionsByDefinition) | **Get** /techpreview/{customerid}/{siteid}/Images/Definitions/{nameOrId}/Versions | Get all image versions associated with a image definition
[**ImagesTPNewImageDefinitionAndVersion**](ImagesTPApi.md#ImagesTPNewImageDefinitionAndVersion) | **Post** /techpreview/{customerid}/{siteid}/Images/Definitions | Create image definition
[**ImagesTPPatchImageDefinition**](ImagesTPApi.md#ImagesTPPatchImageDefinition) | **Patch** /techpreview/{customerid}/{siteid}/Images/Definitions/{nameOrId} | Update a specified image definition
[**ImagesTPPatchImageDefinitionImageVersion**](ImagesTPApi.md#ImagesTPPatchImageDefinitionImageVersion) | **Patch** /techpreview/{customerid}/{siteid}/Images/Definitions/{nameOrId}/Versions/{versionNameOrId} | Update an image version
[**ImagesTPPatchImageVersion**](ImagesTPApi.md#ImagesTPPatchImageVersion) | **Patch** /techpreview/{customerid}/{siteid}/Images/Versions/{id} | Update an image version specified by id
[**ImagesTPRemoveImageDefinition**](ImagesTPApi.md#ImagesTPRemoveImageDefinition) | **Delete** /techpreview/{customerid}/{siteid}/Images/Definitions/{nameOrId} | Delete a specified image definition
[**ImagesTPRemoveImageDefinitionImageVersion**](ImagesTPApi.md#ImagesTPRemoveImageDefinitionImageVersion) | **Delete** /techpreview/{customerid}/{siteid}/Images/Definitions/{nameOrId}/Versions/{versionNameOrId} | Delete a specified image versions
[**ImagesTPRemoveImageVersion**](ImagesTPApi.md#ImagesTPRemoveImageVersion) | **Delete** /techpreview/{customerid}/{siteid}/Images/Versions/{id} | Delete an image versions specified by id



## ImagesTPAddImageVersion

> ImageVersionDetailResponseModel ImagesTPAddImageVersion(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Add an image version to a specific image definition

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
    nameOrId := "nameOrId_example" // string | Name or id of image definition
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewCreateImageVersionRequestModel() // CreateImageVersionRequestModel | CreateImageVersionRequestModel
    async := true // bool | If `true`, this will be queried as a background task. The task will have JobType AddImageVersion. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPAddImageVersion(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPAddImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPAddImageVersion`: ImageVersionDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPAddImageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or id of image definition | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPAddImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**CreateImageVersionRequestModel**](CreateImageVersionRequestModel.md) | CreateImageVersionRequestModel | 
 **async** | **bool** | If &#x60;true&#x60;, this will be queried as a background task. The task will have JobType AddImageVersion. When the task is complete it will redirect to GetJobResults. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageVersionDetailResponseModel**](ImageVersionDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPCheckImageDefinitionExists

> ImagesTPCheckImageDefinitionExists(ctx, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of an image definition by name.

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
    name := "name_example" // string | Name of the image definition.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesTPApi.ImagesTPCheckImageDefinitionExists(context.Background(), name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPCheckImageDefinitionExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the image definition. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPCheckImageDefinitionExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ImagesTPGetCatalogsByImageVersion

> ImageVersionCatalogResponseModelCollection ImagesTPGetCatalogsByImageVersion(ctx, versionId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get catalogs use the specific image version.

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
    versionId := "versionId_example" // string | Image version uid
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of items returned by this query. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPGetCatalogsByImageVersion(context.Background(), versionId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPGetCatalogsByImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPGetCatalogsByImageVersion`: ImageVersionCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPGetCatalogsByImageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | Image version uid | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPGetCatalogsByImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of items returned by this query. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageVersionCatalogResponseModelCollection**](ImageVersionCatalogResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPGetImageDefinition

> ImageDefinitionResponseModel ImagesTPGetImageDefinition(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get a specified image definition

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
    nameOrId := "nameOrId_example" // string | The name or Id of an image definition.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPGetImageDefinition(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPGetImageDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPGetImageDefinition`: ImageDefinitionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPGetImageDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The name or Id of an image definition. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPGetImageDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageDefinitionResponseModel**](ImageDefinitionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPGetImageDefinitionImageVersion

> ImageVersionDetailResponseModel ImagesTPGetImageDefinitionImageVersion(ctx, nameOrId, versionNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details of a specified image version

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
    nameOrId := "nameOrId_example" // string | Name or id of image definition
    versionNameOrId := "versionNameOrId_example" // string | Name or id of image version
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPGetImageDefinitionImageVersion(context.Background(), nameOrId, versionNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPGetImageDefinitionImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPGetImageDefinitionImageVersion`: ImageVersionDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPGetImageDefinitionImageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or id of image definition | 
**versionNameOrId** | **string** | Name or id of image version | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPGetImageDefinitionImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageVersionDetailResponseModel**](ImageVersionDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPGetImageDefinitions

> ImageDefinitionResponseModelCollection ImagesTPGetImageDefinitions(ctx, customerid, siteid).ResourcePool(resourcePool).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all defined image definitions

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
    resourcePool := "resourcePool_example" // string | The name or Id of a resource pool. (optional)
    limit := int32(56) // int32 | The max number of items returned by this query. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPGetImageDefinitions(context.Background(), customerid, siteid).ResourcePool(resourcePool).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPGetImageDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPGetImageDefinitions`: ImageDefinitionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPGetImageDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPGetImageDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourcePool** | **string** | The name or Id of a resource pool. | 
 **limit** | **int32** | The max number of items returned by this query. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageDefinitionResponseModelCollection**](ImageDefinitionResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPGetImageVersion

> ImageVersionDetailResponseModel ImagesTPGetImageVersion(ctx, id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details of an image version specified by Id

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
    id := "id_example" // string | Id of image version
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPGetImageVersion(context.Background(), id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPGetImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPGetImageVersion`: ImageVersionDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPGetImageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of image version | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPGetImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageVersionDetailResponseModel**](ImageVersionDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPGetImageVersionsByDefinition

> ImageVersionResponseModelCollection ImagesTPGetImageVersionsByDefinition(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all image versions associated with a image definition

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
    nameOrId := "nameOrId_example" // string | Name or Id of the image definition, currently only Id is supported.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of items returned by this query. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPGetImageVersionsByDefinition(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPGetImageVersionsByDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPGetImageVersionsByDefinition`: ImageVersionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPGetImageVersionsByDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the image definition, currently only Id is supported. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPGetImageVersionsByDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of items returned by this query. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageVersionResponseModelCollection**](ImageVersionResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPNewImageDefinitionAndVersion

> ImageVersionDetailResponseModel ImagesTPNewImageDefinitionAndVersion(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create image definition

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
    request := *openapiclient.NewCreateImageDefinitionRequestModel() // CreateImageDefinitionRequestModel | CreateImageDefinitionRequestModel
    async := true // bool | If `true`, this will be queried as a background task. The task will have JobType NewImageDefinition. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPNewImageDefinitionAndVersion(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPNewImageDefinitionAndVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPNewImageDefinitionAndVersion`: ImageVersionDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPNewImageDefinitionAndVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPNewImageDefinitionAndVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateImageDefinitionRequestModel**](CreateImageDefinitionRequestModel.md) | CreateImageDefinitionRequestModel | 
 **async** | **bool** | If &#x60;true&#x60;, this will be queried as a background task. The task will have JobType NewImageDefinition. When the task is complete it will redirect to GetJobResults. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageVersionDetailResponseModel**](ImageVersionDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPPatchImageDefinition

> ImageDefinitionResponseModel ImagesTPPatchImageDefinition(ctx, nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a specified image definition

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
    nameOrId := "nameOrId_example" // string | Name or id of an image definition
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewUpdateImageDefinitionRequestModel() // UpdateImageDefinitionRequestModel | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPPatchImageDefinition(context.Background(), nameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPPatchImageDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPPatchImageDefinition`: ImageDefinitionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPPatchImageDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or id of an image definition | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPPatchImageDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateImageDefinitionRequestModel**](UpdateImageDefinitionRequestModel.md) |  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageDefinitionResponseModel**](ImageDefinitionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPPatchImageDefinitionImageVersion

> ImageVersionDetailResponseModel ImagesTPPatchImageDefinitionImageVersion(ctx, nameOrId, versionNameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an image version

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
    nameOrId := "nameOrId_example" // string | Name or id of image definition
    versionNameOrId := "versionNameOrId_example" // string | Name or id of image version
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewUpdateImageVersionRequestModel() // UpdateImageVersionRequestModel | Request model to update image version
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPPatchImageDefinitionImageVersion(context.Background(), nameOrId, versionNameOrId, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPPatchImageDefinitionImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPPatchImageDefinitionImageVersion`: ImageVersionDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPPatchImageDefinitionImageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or id of image definition | 
**versionNameOrId** | **string** | Name or id of image version | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPPatchImageDefinitionImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**UpdateImageVersionRequestModel**](UpdateImageVersionRequestModel.md) | Request model to update image version | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageVersionDetailResponseModel**](ImageVersionDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPPatchImageVersion

> ImageVersionDetailResponseModel ImagesTPPatchImageVersion(ctx, id, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an image version specified by id

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
    id := "id_example" // string | Id of image version
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewUpdateImageVersionRequestModel() // UpdateImageVersionRequestModel | Request model to update image version
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesTPApi.ImagesTPPatchImageVersion(context.Background(), id, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPPatchImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTPPatchImageVersion`: ImageVersionDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImagesTPApi.ImagesTPPatchImageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of image version | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPPatchImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateImageVersionRequestModel**](UpdateImageVersionRequestModel.md) | Request model to update image version | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ImageVersionDetailResponseModel**](ImageVersionDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesTPRemoveImageDefinition

> ImagesTPRemoveImageDefinition(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a specified image definition

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
    nameOrId := "nameOrId_example" // string | Name or id of image definition
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, this will be queried as a background task. The task will have JobType AddImageVersion. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesTPApi.ImagesTPRemoveImageDefinition(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPRemoveImageDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or id of image definition | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPRemoveImageDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, this will be queried as a background task. The task will have JobType AddImageVersion. When the task is complete it will redirect to GetJobResults. | [default to false]
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


## ImagesTPRemoveImageDefinitionImageVersion

> ImagesTPRemoveImageDefinitionImageVersion(ctx, nameOrId, versionNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a specified image versions

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
    nameOrId := "nameOrId_example" // string | Name or id of image definition
    versionNameOrId := "versionNameOrId_example" // string | Name or id of image version
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesTPApi.ImagesTPRemoveImageDefinitionImageVersion(context.Background(), nameOrId, versionNameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPRemoveImageDefinitionImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or id of image definition | 
**versionNameOrId** | **string** | Name or id of image version | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPRemoveImageDefinitionImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## ImagesTPRemoveImageVersion

> ImagesTPRemoveImageVersion(ctx, id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an image versions specified by id

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
    id := "id_example" // string | Id of image version
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesTPApi.ImagesTPRemoveImageVersion(context.Background(), id, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesTPApi.ImagesTPRemoveImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of image version | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesTPRemoveImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

