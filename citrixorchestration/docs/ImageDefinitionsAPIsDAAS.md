# \ImageDefinitionsAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImageDefinitionsCheckImageDefinitionExist**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsCheckImageDefinitionExist) | **Head** /ImageDefinitions/{name} | Check for the existence of an image definition by name.
[**ImageDefinitionsCreateImageDefinition**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsCreateImageDefinition) | **Post** /ImageDefinitions | Create an image definition.
[**ImageDefinitionsCreateImageVersion**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsCreateImageVersion) | **Post** /ImageDefinitions/{nameOrId}/ImageVersions | Create an image version.
[**ImageDefinitionsDeleteImageDefinition**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsDeleteImageDefinition) | **Delete** /ImageDefinitions/{nameOrId} | Delete an image definition.
[**ImageDefinitionsDeleteImageDefinitionImageVersion**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsDeleteImageDefinitionImageVersion) | **Delete** /ImageDefinitions/{nameOrId}/ImageVersions/{versionNumberOrId} | Delete an image version.
[**ImageDefinitionsDoImageDefinitionAndImageVersionSearch**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsDoImageDefinitionAndImageVersionSearch) | **Post** /ImageDefinitionsAndImageVersions/$search | Perform an advanced search for image definitions and image versions.
[**ImageDefinitionsGetImageDefinition**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsGetImageDefinition) | **Get** /ImageDefinitions/{nameOrId} | Get details about a single image definition.
[**ImageDefinitionsGetImageDefinitionImageVersion**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsGetImageDefinitionImageVersion) | **Get** /ImageDefinitions/{nameOrId}/ImageVersions/{versionNumberOrId} | Get details about a single image version.
[**ImageDefinitionsGetImageDefinitionImageVersions**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsGetImageDefinitionImageVersions) | **Get** /ImageDefinitions/{nameOrId}/ImageVersions | Get all image versions associated with an image definition.
[**ImageDefinitionsGetImageDefinitions**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsGetImageDefinitions) | **Get** /ImageDefinitions | Get all image definitions.
[**ImageDefinitionsGetImageVersionProvisioningSchemes**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsGetImageVersionProvisioningSchemes) | **Get** /ImageDefinitions/{nameOrId}/ImageVersions/{versionNumberOrId}/ProvisioningSchemes | Get provisioning schemes associated with an image version.
[**ImageDefinitionsSetImageDefinitionImageVersion**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsSetImageDefinitionImageVersion) | **Put** /ImageDefinitions/{nameOrId}/ImageVersions/{versionNumberOrId} | Set properties associated with an image version.
[**ImageDefinitionsUpdateImageDefinition**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsUpdateImageDefinition) | **Patch** /ImageDefinitions/{nameOrId} | Update an image definition.
[**ImageDefinitionsUpdateImageDefinitionImageVersion**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsUpdateImageDefinitionImageVersion) | **Patch** /ImageDefinitions/{nameOrId}/ImageVersions/{versionNumberOrId} | Update an image version.
[**ImageDefinitionsUpdateImageVersionResourcePools**](ImageDefinitionsAPIsDAAS.md#ImageDefinitionsUpdateImageVersionResourcePools) | **Put** /ImageDefinitions/{nameOrId}/ImageVersions/{versionNumberOrId}/$UpdateResourcePools | Update resource pools associated with an image version.



## ImageDefinitionsCheckImageDefinitionExist

> ImageDefinitionsCheckImageDefinitionExist(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    name := "name_example" // string | Name of the image definition.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsCheckImageDefinitionExist(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsCheckImageDefinitionExist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the image definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsCheckImageDefinitionExistRequest struct via the builder pattern


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


## ImageDefinitionsCreateImageDefinition

> ImageDefinitionResponseModel ImageDefinitionsCreateImageDefinition(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateImageDefinitionRequestModel(createImageDefinitionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Create an image definition.

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
    createImageDefinitionRequestModel := *openapiclient.NewCreateImageDefinitionRequestModel("Name_example", openapiclient.OsType("Unknown"), openapiclient.SessionSupport("Unknown")) // CreateImageDefinitionRequestModel | Details about the image definition to create.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | Create image definition async (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsCreateImageDefinition(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateImageDefinitionRequestModel(createImageDefinitionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsCreateImageDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsCreateImageDefinition`: ImageDefinitionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsCreateImageDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsCreateImageDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createImageDefinitionRequestModel** | [**CreateImageDefinitionRequestModel**](CreateImageDefinitionRequestModel.md) | Details about the image definition to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | Create image definition async | [default to false]

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


## ImageDefinitionsCreateImageVersion

> ImageVersionResponseModel ImageDefinitionsCreateImageVersion(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateImageVersionRequestModel(createImageVersionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Create an image version.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the image definition.
    createImageVersionRequestModel := *openapiclient.NewCreateImageVersionRequestModel(*openapiclient.NewCreateImageSchemeRequestModel(), "For example, XDHyp:\path\to\MasterImage", "ResourcePool_example") // CreateImageVersionRequestModel | Details about the image version to create.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the image version (and associated objects) will be created as a background task. The task will have JobType CreateImageVersion. When the task is complete it will redirect to GetImageDefinitionImageVersion. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsCreateImageVersion(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateImageVersionRequestModel(createImageVersionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsCreateImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsCreateImageVersion`: ImageVersionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsCreateImageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the image definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsCreateImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **createImageVersionRequestModel** | [**CreateImageVersionRequestModel**](CreateImageVersionRequestModel.md) | Details about the image version to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the image version (and associated objects) will be created as a background task. The task will have JobType CreateImageVersion. When the task is complete it will redirect to GetImageDefinitionImageVersion. | [default to false]

### Return type

[**ImageVersionResponseModel**](ImageVersionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDefinitionsDeleteImageDefinition

> ImageDefinitionsDeleteImageDefinition(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Delete an image definition.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the image definition to delete.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | Async call (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsDeleteImageDefinition(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsDeleteImageDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the image definition to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsDeleteImageDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | Async call | 

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


## ImageDefinitionsDeleteImageDefinitionImageVersion

> ImageDefinitionsDeleteImageDefinitionImageVersion(ctx, nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).PurgeDBOnly(purgeDBOnly).Async(async).Execute()

Delete an image version.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the image definition.
    versionNumberOrId := "versionNumberOrId_example" // string | Number or ID of the image version to delete.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    purgeDBOnly := true // bool | If this option is specified, this command will only remove the image version data from the Citrix site database.  However, the disk images created in the image version still remain in the hypervisor.   Optional; default is `false`. (optional) (default to false)
    async := true // bool | If `true`, the image version (and associated objects) will be deleted as a background task. The task will have JobType DeleteImageVersion. When the task is complete it will redirect to GetImageDefinitionImageVersion. The job's Parameters will contain properties: * _Id_ - ID of the image version being deleted, * _Number_ - Number of the image version being deleted. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsDeleteImageDefinitionImageVersion(context.Background(), nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).PurgeDBOnly(purgeDBOnly).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsDeleteImageDefinitionImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the image definition. | 
**versionNumberOrId** | **string** | Number or ID of the image version to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsDeleteImageDefinitionImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **purgeDBOnly** | **bool** | If this option is specified, this command will only remove the image version data from the Citrix site database.  However, the disk images created in the image version still remain in the hypervisor.   Optional; default is &#x60;false&#x60;. | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, the image version (and associated objects) will be deleted as a background task. The task will have JobType DeleteImageVersion. When the task is complete it will redirect to GetImageDefinitionImageVersion. The job&#39;s Parameters will contain properties: * _Id_ - ID of the image version being deleted, * _Number_ - Number of the image version being deleted. | [default to false]

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


## ImageDefinitionsDoImageDefinitionAndImageVersionSearch

> ImageDefinitionsAndImageVersionsResponseModelCollection ImageDefinitionsDoImageDefinitionAndImageVersionSearch(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).ImageDefinitionAndImageVersionSearchRequestModel(imageDefinitionAndImageVersionSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Perform an advanced search for image definitions and image versions.



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
    imageDefinitionAndImageVersionSearchRequestModel := *openapiclient.NewImageDefinitionAndImageVersionSearchRequestModel() // ImageDefinitionAndImageVersionSearchRequestModel | Specifies the advanced search parameters.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the search operation will be executed as a background task. The task will have JobType DoImageDefinitionAndImageVersionSearch. When the task is complete it will redirect to GetJobResults. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsDoImageDefinitionAndImageVersionSearch(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).ImageDefinitionAndImageVersionSearchRequestModel(imageDefinitionAndImageVersionSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsDoImageDefinitionAndImageVersionSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsDoImageDefinitionAndImageVersionSearch`: ImageDefinitionsAndImageVersionsResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsDoImageDefinitionAndImageVersionSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsDoImageDefinitionAndImageVersionSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **imageDefinitionAndImageVersionSearchRequestModel** | [**ImageDefinitionAndImageVersionSearchRequestModel**](ImageDefinitionAndImageVersionSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the search operation will be executed as a background task. The task will have JobType DoImageDefinitionAndImageVersionSearch. When the task is complete it will redirect to GetJobResults. | [default to false]

### Return type

[**ImageDefinitionsAndImageVersionsResponseModelCollection**](ImageDefinitionsAndImageVersionsResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDefinitionsGetImageDefinition

> ImageDefinitionResponseModel ImageDefinitionsGetImageDefinition(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get details about a single image definition.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the image definition.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinition(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsGetImageDefinition`: ImageDefinitionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the image definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsGetImageDefinitionRequest struct via the builder pattern


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

[**ImageDefinitionResponseModel**](ImageDefinitionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDefinitionsGetImageDefinitionImageVersion

> ImageVersionResponseModel ImageDefinitionsGetImageDefinitionImageVersion(ctx, nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get details about a single image version.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the image definition.
    versionNumberOrId := "versionNumberOrId_example" // string | Number or ID of the image version.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitionImageVersion(context.Background(), nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitionImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsGetImageDefinitionImageVersion`: ImageVersionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitionImageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the image definition. | 
**versionNumberOrId** | **string** | Number or ID of the image version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsGetImageDefinitionImageVersionRequest struct via the builder pattern


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

[**ImageVersionResponseModel**](ImageVersionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDefinitionsGetImageDefinitionImageVersions

> ImageVersionResponseModelCollection ImageDefinitionsGetImageDefinitionImageVersions(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

Get all image versions associated with an image definition.

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
    nameOrId := "nameOrId_example" // string | Name or Id of the image definition.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    limit := int32(56) // int32 | The max number of image versions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitionImageVersions(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitionImageVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsGetImageDefinitionImageVersions`: ImageVersionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitionImageVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the image definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsGetImageDefinitionImageVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of image versions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 

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


## ImageDefinitionsGetImageDefinitions

> ImageDefinitionResponseModelCollection ImageDefinitionsGetImageDefinitions(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

Get all image definitions.

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
    limit := int32(56) // int32 | The max number of image definitions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitions(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsGetImageDefinitions`: ImageDefinitionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsGetImageDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of image definitions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 

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


## ImageDefinitionsGetImageVersionProvisioningSchemes

> ImageVersionProvisioningSchemeRefResponseModelCollection ImageDefinitionsGetImageVersionProvisioningSchemes(ctx, nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get provisioning schemes associated with an image version.

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
    nameOrId := "nameOrId_example" // string | Name or ID of image definition.
    versionNumberOrId := "versionNumberOrId_example" // string | Number or ID of image version.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageVersionProvisioningSchemes(context.Background(), nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageVersionProvisioningSchemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsGetImageVersionProvisioningSchemes`: ImageVersionProvisioningSchemeRefResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsGetImageVersionProvisioningSchemes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of image definition. | 
**versionNumberOrId** | **string** | Number or ID of image version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsGetImageVersionProvisioningSchemesRequest struct via the builder pattern


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

[**ImageVersionProvisioningSchemeRefResponseModelCollection**](ImageVersionProvisioningSchemeRefResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDefinitionsSetImageDefinitionImageVersion

> ImageDefinitionsSetImageDefinitionImageVersion(ctx, nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateImageVersionRequestModel(createImageVersionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Set properties associated with an image version.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the image definition.
    versionNumberOrId := "versionNumberOrId_example" // string | Number or ID of the image version to update.
    createImageVersionRequestModel := *openapiclient.NewCreateImageVersionRequestModel(*openapiclient.NewCreateImageSchemeRequestModel(), "For example, XDHyp:\path\to\MasterImage", "ResourcePool_example") // CreateImageVersionRequestModel | Details about the image version to update.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the image version (and associated objects) will be updated as a background task. The task will have JobType SetImageVersion. When the task is complete it will redirect to GetImageDefinitionImageVersion. The job's Parameters will contain properties: * _Id_ - ID of the image version being updated. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsSetImageDefinitionImageVersion(context.Background(), nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateImageVersionRequestModel(createImageVersionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsSetImageDefinitionImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the image definition. | 
**versionNumberOrId** | **string** | Number or ID of the image version to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsSetImageDefinitionImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **createImageVersionRequestModel** | [**CreateImageVersionRequestModel**](CreateImageVersionRequestModel.md) | Details about the image version to update. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the image version (and associated objects) will be updated as a background task. The task will have JobType SetImageVersion. When the task is complete it will redirect to GetImageDefinitionImageVersion. The job&#39;s Parameters will contain properties: * _Id_ - ID of the image version being updated. | [default to false]

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


## ImageDefinitionsUpdateImageDefinition

> ImageDefinitionResponseModel ImageDefinitionsUpdateImageDefinition(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateImageDefinitionRequestModel(updateImageDefinitionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update an image definition.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the image definition to update.
    updateImageDefinitionRequestModel := *openapiclient.NewUpdateImageDefinitionRequestModel() // UpdateImageDefinitionRequestModel | Properties of the image definition to update.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageDefinition(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateImageDefinitionRequestModel(updateImageDefinitionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsUpdateImageDefinition`: ImageDefinitionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the image definition to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsUpdateImageDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **updateImageDefinitionRequestModel** | [**UpdateImageDefinitionRequestModel**](UpdateImageDefinitionRequestModel.md) | Properties of the image definition to update. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## ImageDefinitionsUpdateImageDefinitionImageVersion

> ImageVersionResponseModel ImageDefinitionsUpdateImageDefinitionImageVersion(ctx, nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateImageVersionRequestModel(updateImageVersionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update an image version.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the image definition.
    versionNumberOrId := "versionNumberOrId_example" // string | Number or ID of the image version to update.
    updateImageVersionRequestModel := *openapiclient.NewUpdateImageVersionRequestModel() // UpdateImageVersionRequestModel | Properties of the image version to update.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageDefinitionImageVersion(context.Background(), nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateImageVersionRequestModel(updateImageVersionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageDefinitionImageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageDefinitionsUpdateImageDefinitionImageVersion`: ImageVersionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageDefinitionImageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the image definition. | 
**versionNumberOrId** | **string** | Number or ID of the image version to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsUpdateImageDefinitionImageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **updateImageVersionRequestModel** | [**UpdateImageVersionRequestModel**](UpdateImageVersionRequestModel.md) | Properties of the image version to update. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**ImageVersionResponseModel**](ImageVersionResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageDefinitionsUpdateImageVersionResourcePools

> ImageDefinitionsUpdateImageVersionResourcePools(ctx, nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateImageVersionResourcePoolsRequestModel(updateImageVersionResourcePoolsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Update resource pools associated with an image version.



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
    nameOrId := "nameOrId_example" // string | Name or ID of image definition.
    versionNumberOrId := "versionNumberOrId_example" // string | 
    updateImageVersionResourcePoolsRequestModel := *openapiclient.NewUpdateImageVersionResourcePoolsRequestModel() // UpdateImageVersionResourcePoolsRequestModel | Details of the resource pools to add to the image version.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the resource pools will be added into the image version as a background task. The task will have JobType UpdateImageVersionResourcePools. When the task is complete it will redirect to GetImageDefinitionImageVersion. The job's Parameters will contain properties: * _Id_ - ID of the image version being updated. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageVersionResourcePools(context.Background(), nameOrId, versionNumberOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateImageVersionResourcePoolsRequestModel(updateImageVersionResourcePoolsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageDefinitionsAPIsDAAS.ImageDefinitionsUpdateImageVersionResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of image definition. | 
**versionNumberOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageDefinitionsUpdateImageVersionResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **updateImageVersionResourcePoolsRequestModel** | [**UpdateImageVersionResourcePoolsRequestModel**](UpdateImageVersionResourcePoolsRequestModel.md) | Details of the resource pools to add to the image version. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the resource pools will be added into the image version as a background task. The task will have JobType UpdateImageVersionResourcePools. When the task is complete it will redirect to GetImageDefinitionImageVersion. The job&#39;s Parameters will contain properties: * _Id_ - ID of the image version being updated. | [default to false]

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

