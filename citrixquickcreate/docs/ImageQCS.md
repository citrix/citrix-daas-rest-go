# \ImageQCS

All URIs are relative to *https://api.cloud.com/quickcreateservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyImageAsync**](ImageQCS.md#CopyImageAsync) | **Post** /{customerId}/accounts/{accountId}/images/{imageId}/$copy | Makes a copy of a workspace image
[**GetImageAsync**](ImageQCS.md#GetImageAsync) | **Get** /{customerId}/accounts/{accountId}/images/{imageId} | Gets image from workspace
[**GetImagesAsync**](ImageQCS.md#GetImagesAsync) | **Get** /{customerId}/accounts/{accountId}/images | Gets images from workspace
[**ImportImageAsync**](ImageQCS.md#ImportImageAsync) | **Post** /{customerId}/accounts/{accountId}/images/$import | Imports image to workspace
[**RemoveImageAsync**](ImageQCS.md#RemoveImageAsync) | **Delete** /{customerId}/accounts/{accountId}/images/{imageId} | Removes image to workspace



## CopyImageAsync

> AwsEdcImage CopyImageAsync(ctx, customerId, accountId, imageId).ImageName(imageName).CitrixTransactionId(citrixTransactionId).Execute()

Makes a copy of a workspace image

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickcreate"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    accountId := "accountId_example" // string | ID of the account
    imageId := "imageId_example" // string | ID of the image
    imageName := "imageName_example" // string | Optional new image name (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageQCS.CopyImageAsync(context.Background(), customerId, accountId, imageId).ImageName(imageName).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageQCS.CopyImageAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyImageAsync`: AwsEdcImage
    fmt.Fprintf(os.Stdout, "Response from `ImageQCS.CopyImageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of the account | 
**imageId** | **string** | ID of the image | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyImageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **imageName** | **string** | Optional new image name | 
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**AwsEdcImage**](AwsEdcImage.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageAsync

> AwsEdcImage GetImageAsync(ctx, customerId, accountId, imageId).CitrixTransactionId(citrixTransactionId).Execute()

Gets image from workspace

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickcreate"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    accountId := "accountId_example" // string | ID of account
    imageId := "imageId_example" // string | Id of image
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageQCS.GetImageAsync(context.Background(), customerId, accountId, imageId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageQCS.GetImageAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageAsync`: AwsEdcImage
    fmt.Fprintf(os.Stdout, "Response from `ImageQCS.GetImageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of account | 
**imageId** | **string** | Id of image | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**AwsEdcImage**](AwsEdcImage.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagesAsync

> Images GetImagesAsync(ctx, customerId, accountId).CitrixTransactionId(citrixTransactionId).Execute()

Gets images from workspace

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickcreate"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    accountId := "accountId_example" // string | ID of account
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageQCS.GetImagesAsync(context.Background(), customerId, accountId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageQCS.GetImagesAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImagesAsync`: Images
    fmt.Fprintf(os.Stdout, "Response from `ImageQCS.GetImagesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**Images**](Images.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImageAsync

> AwsEdcImage ImportImageAsync(ctx, customerId, accountId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Imports image to workspace

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickcreate"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    accountId := "accountId_example" // string | ID of hypervisor
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := ImportAwsEdcImage(987) // ImportAwsEdcImage | Import image configuration (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageQCS.ImportImageAsync(context.Background(), customerId, accountId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageQCS.ImportImageAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageAsync`: AwsEdcImage
    fmt.Fprintf(os.Stdout, "Response from `ImageQCS.ImportImageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of hypervisor | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportImageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **ImportAwsEdcImage** | Import image configuration | 

### Return type

[**AwsEdcImage**](AwsEdcImage.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveImageAsync

> RemoveImageAsync(ctx, customerId, accountId, imageId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).Execute()

Removes image to workspace

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickcreate"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    accountId := "accountId_example" // string | ID of hypervisor
    imageId := "imageId_example" // string | ID of image
    forceDelete := true // bool | Force delete an image (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImageQCS.RemoveImageAsync(context.Background(), customerId, accountId, imageId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageQCS.RemoveImageAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of hypervisor | 
**imageId** | **string** | ID of image | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveImageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forceDelete** | **bool** | Force delete an image | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

 (empty response body)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

