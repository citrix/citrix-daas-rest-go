# \MasterImageCMD

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTemplateImage**](MasterImageCMD.md#AddTemplateImage) | **Post** /{customerId}/{siteId}/images | Add a new master image to the XenApp Essential customer&#39;s account, which is linked to a VHD image inside the customer&#39;s storage account.
[**CancelCustomerImageUrl**](MasterImageCMD.md#CancelCustomerImageUrl) | **Delete** /{customerId}/{siteId}/images/{imageId}/imgUrl | Cancel url of customer image
[**CreateCustomerImageUrl**](MasterImageCMD.md#CreateCustomerImageUrl) | **Post** /{customerId}/{siteId}/images/{imageId}/imgUrl | Returns url of customer image
[**DeleteTemplateImage**](MasterImageCMD.md#DeleteTemplateImage) | **Delete** /{customerId}/{siteId}/images/{imageId} | Delete the specified master image
[**GetAzureSasUrlExpiryTime**](MasterImageCMD.md#GetAzureSasUrlExpiryTime) | **Get** /{customerId}/{siteId}/images/{imageId}/azureSasUrlExpiryTime | Returns the expiry time of the Azure SAS URL
[**GetImages**](MasterImageCMD.md#GetImages) | **Get** /{customerId}/{siteId}/images | Returns all the master images the customer has linked to their account
[**GetTemplateImage**](MasterImageCMD.md#GetTemplateImage) | **Get** /{customerId}/{siteId}/images/{imageId} | Returns details of the specified master image
[**ImportTemplateImage**](MasterImageCMD.md#ImportTemplateImage) | **Post** /{customerId}/{siteId}/images/$import | Add a new master image to the DaaS customer&#39;s account, which is linked to a VHD image inside the customer&#39;s storage account (Image import feature for DaaS)
[**UpdateTemplateImage**](MasterImageCMD.md#UpdateTemplateImage) | **Patch** /{customerId}/{siteId}/images/{imageId} | Updates template image configuration (name, notes, allowed ips)



## AddTemplateImage

> string AddTemplateImage(ctx, customerId, siteId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Add a new master image to the XenApp Essential customer's account, which is linked to a VHD image inside the customer's storage account.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | The site ID of the customer
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := *openapiclient.NewAddTemplateImageModel("Name_example", "SubscriptionId_example", "ResourceGroup_example", "StorageAccount_example") // AddTemplateImageModel | Configuration of the master image to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MasterImageCMD.AddTemplateImage(context.Background(), customerId, siteId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterImageCMD.AddTemplateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTemplateImage`: string
    fmt.Fprintf(os.Stdout, "Response from `MasterImageCMD.AddTemplateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | The site ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTemplateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | [**AddTemplateImageModel**](AddTemplateImageModel.md) | Configuration of the master image to add | 

### Return type

**string**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelCustomerImageUrl

> CancelCustomerImageUrl(ctx, customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Execute()

Cancel url of customer image

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | The site ID of the customer
    imageId := "imageId_example" // string | The master image to view details of
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MasterImageCMD.CancelCustomerImageUrl(context.Background(), customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterImageCMD.CancelCustomerImageUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | The site ID of the customer | 
**imageId** | **string** | The master image to view details of | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCustomerImageUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## CreateCustomerImageUrl

> TemplateImageUrl CreateCustomerImageUrl(ctx, customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Returns url of customer image

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | The site ID of the customer
    imageId := "imageId_example" // string | The master image to view details of
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := *openapiclient.NewCustomerImgUrlModel() // CustomerImgUrlModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MasterImageCMD.CreateCustomerImageUrl(context.Background(), customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterImageCMD.CreateCustomerImageUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerImageUrl`: TemplateImageUrl
    fmt.Fprintf(os.Stdout, "Response from `MasterImageCMD.CreateCustomerImageUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | The site ID of the customer | 
**imageId** | **string** | The master image to view details of | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerImageUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | [**CustomerImgUrlModel**](CustomerImgUrlModel.md) |  | 

### Return type

[**TemplateImageUrl**](TemplateImageUrl.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplateImage

> DeleteTemplateImage(ctx, customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Execute()

Delete the specified master image

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | The site ID of the customer
    imageId := "imageId_example" // string | The master image to update
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MasterImageCMD.DeleteTemplateImage(context.Background(), customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterImageCMD.DeleteTemplateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | The site ID of the customer | 
**imageId** | **string** | The master image to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetAzureSasUrlExpiryTime

> TemplateImageUrl GetAzureSasUrlExpiryTime(ctx, customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Execute()

Returns the expiry time of the Azure SAS URL

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | The site ID of the customer
    imageId := "imageId_example" // string | The master image to view details of
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MasterImageCMD.GetAzureSasUrlExpiryTime(context.Background(), customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterImageCMD.GetAzureSasUrlExpiryTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureSasUrlExpiryTime`: TemplateImageUrl
    fmt.Fprintf(os.Stdout, "Response from `MasterImageCMD.GetAzureSasUrlExpiryTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | The site ID of the customer | 
**imageId** | **string** | The master image to view details of | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureSasUrlExpiryTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**TemplateImageUrl**](TemplateImageUrl.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImages

> CustomerTemplateImageOverviewsModel GetImages(ctx, customerId, siteId).CspCustomerId(cspCustomerId).CspSiteId(cspSiteId).AzureSubscriptions(azureSubscriptions).IncludeCitrixPrepared(includeCitrixPrepared).IncludeCustomerPrepared(includeCustomerPrepared).CitrixTransactionId(citrixTransactionId).Execute()

Returns all the master images the customer has linked to their account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | The site ID of the customer
    cspCustomerId := "cspCustomerId_example" // string | Name of tenant customer ID if partner-tenant relationship exists otherwise null (optional)
    cspSiteId := "cspSiteId_example" // string | Name of tenant site ID if partner-tenant relationship exists otherwise null (optional)
    azureSubscriptions := []string{"Inner_example"} // []string | Optional: The IDs of the Azure Subcrption we want customer images from (If not specified we get all) (optional)
    includeCitrixPrepared := true // bool | Indicates if Citrix prepared images should be included (optional) (default to true)
    includeCustomerPrepared := true // bool | Indicates if customer prepared images should be included (optional) (default to true)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MasterImageCMD.GetImages(context.Background(), customerId, siteId).CspCustomerId(cspCustomerId).CspSiteId(cspSiteId).AzureSubscriptions(azureSubscriptions).IncludeCitrixPrepared(includeCitrixPrepared).IncludeCustomerPrepared(includeCustomerPrepared).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterImageCMD.GetImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImages`: CustomerTemplateImageOverviewsModel
    fmt.Fprintf(os.Stdout, "Response from `MasterImageCMD.GetImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | The site ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cspCustomerId** | **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | 
 **cspSiteId** | **string** | Name of tenant site ID if partner-tenant relationship exists otherwise null | 
 **azureSubscriptions** | **[]string** | Optional: The IDs of the Azure Subcrption we want customer images from (If not specified we get all) | 
 **includeCitrixPrepared** | **bool** | Indicates if Citrix prepared images should be included | [default to true]
 **includeCustomerPrepared** | **bool** | Indicates if customer prepared images should be included | [default to true]
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**CustomerTemplateImageOverviewsModel**](CustomerTemplateImageOverviewsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateImage

> TemplateImageDetails GetTemplateImage(ctx, customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Execute()

Returns details of the specified master image

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | The site ID of the customer
    imageId := "imageId_example" // string | The master image to view details of
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MasterImageCMD.GetTemplateImage(context.Background(), customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterImageCMD.GetTemplateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateImage`: TemplateImageDetails
    fmt.Fprintf(os.Stdout, "Response from `MasterImageCMD.GetTemplateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | The site ID of the customer | 
**imageId** | **string** | The master image to view details of | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**TemplateImageDetails**](TemplateImageDetails.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportTemplateImage

> TemplateImageOverview ImportTemplateImage(ctx, customerId, siteId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Add a new master image to the DaaS customer's account, which is linked to a VHD image inside the customer's storage account (Image import feature for DaaS)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | The site ID of the customer
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := *openapiclient.NewImportTemplateImageModel("Name_example", "VhdUri_example") // ImportTemplateImageModel | Configuration of the master image to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MasterImageCMD.ImportTemplateImage(context.Background(), customerId, siteId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterImageCMD.ImportTemplateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportTemplateImage`: TemplateImageOverview
    fmt.Fprintf(os.Stdout, "Response from `MasterImageCMD.ImportTemplateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | The site ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportTemplateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | [**ImportTemplateImageModel**](ImportTemplateImageModel.md) | Configuration of the master image to add | 

### Return type

[**TemplateImageOverview**](TemplateImageOverview.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplateImage

> UpdateTemplateImage(ctx, customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Updates template image configuration (name, notes, allowed ips)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | 
    imageId := "imageId_example" // string | The template image to update
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := *openapiclient.NewUpdateTemplateImageModel() // UpdateTemplateImageModel | The updated configuration of image (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MasterImageCMD.UpdateTemplateImage(context.Background(), customerId, siteId, imageId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasterImageCMD.UpdateTemplateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** |  | 
**imageId** | **string** | The template image to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | [**UpdateTemplateImageModel**](UpdateTemplateImageModel.md) | The updated configuration of image | 

### Return type

 (empty response body)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

