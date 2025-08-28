# \PvsStreamingAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PvsStreamingGetPvsStreamingSites**](PvsStreamingAPIsDAAS.md#PvsStreamingGetPvsStreamingSites) | **Get** /PvsStreaming/Sites | Get the list of PVS sites.
[**PvsStreamingGetPvsStreamingStores**](PvsStreamingAPIsDAAS.md#PvsStreamingGetPvsStreamingStores) | **Get** /PvsStreaming/Stores/{farmId} | Get the list of PVS stores.
[**PvsStreamingGetPvsStreamingVDisks**](PvsStreamingAPIsDAAS.md#PvsStreamingGetPvsStreamingVDisks) | **Get** /PvsStreaming/vDisks | Get the list of PVS vDisks.
[**PvsStreamingTestDeviceCollectionExists**](PvsStreamingAPIsDAAS.md#PvsStreamingTestDeviceCollectionExists) | **Post** /PvsStreaming/$checkDeviceCollectionName | Test the existence of a device collection by name.



## PvsStreamingGetPvsStreamingSites

> PvsStreamingSiteResponseModelCollection PvsStreamingGetPvsStreamingSites(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).FarmId(farmId).Async(async).Execute()

Get the list of PVS sites.

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
	farmId := "farmId_example" // string | PVS farm id. (optional)
	async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingSites(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).FarmId(farmId).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PvsStreamingGetPvsStreamingSites`: PvsStreamingSiteResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPvsStreamingGetPvsStreamingSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **farmId** | **string** | PVS farm id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]

### Return type

[**PvsStreamingSiteResponseModelCollection**](PvsStreamingSiteResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvsStreamingGetPvsStreamingStores

> PvsStreamingStoreResponseModelCollection PvsStreamingGetPvsStreamingStores(ctx, farmId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).PvsSiteId(pvsSiteId).Async(async).Execute()

Get the list of PVS stores.

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
	farmId := "farmId_example" // string | PVS farm id.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	pvsSiteId := "pvsSiteId_example" // string | PVS site id. (optional)
	async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingStores(context.Background(), farmId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).PvsSiteId(pvsSiteId).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PvsStreamingGetPvsStreamingStores`: PvsStreamingStoreResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingStores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**farmId** | **string** | PVS farm id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPvsStreamingGetPvsStreamingStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **pvsSiteId** | **string** | PVS site id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]

### Return type

[**PvsStreamingStoreResponseModelCollection**](PvsStreamingStoreResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvsStreamingGetPvsStreamingVDisks

> PvsStreamingVDiskResponseModelCollection PvsStreamingGetPvsStreamingVDisks(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).FarmId(farmId).PvsSiteId(pvsSiteId).StoreId(storeId).DiskLocatorId(diskLocatorId).Async(async).Execute()

Get the list of PVS vDisks.

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
	farmId := "farmId_example" // string | PVS farm id. (optional)
	pvsSiteId := "pvsSiteId_example" // string | PVS site id. (optional)
	storeId := "storeId_example" // string | PVS store id. (optional)
	diskLocatorId := "diskLocatorId_example" // string | PVS vDisk id. (optional)
	async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingVDisks(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).FarmId(farmId).PvsSiteId(pvsSiteId).StoreId(storeId).DiskLocatorId(diskLocatorId).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingVDisks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PvsStreamingGetPvsStreamingVDisks`: PvsStreamingVDiskResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingVDisks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPvsStreamingGetPvsStreamingVDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **farmId** | **string** | PVS farm id. | 
 **pvsSiteId** | **string** | PVS site id. | 
 **storeId** | **string** | PVS store id. | 
 **diskLocatorId** | **string** | PVS vDisk id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]

### Return type

[**PvsStreamingVDiskResponseModelCollection**](PvsStreamingVDiskResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvsStreamingTestDeviceCollectionExists

> PvsStreamingTestDeviceCollectionExists(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).DeviceCollectionNameCheckRequestModel(deviceCollectionNameCheckRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Test the existence of a device collection by name.



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
	deviceCollectionNameCheckRequestModel := *openapiclient.NewDeviceCollectionNameCheckRequestModel() // DeviceCollectionNameCheckRequestModel | Name of the device collection and PVS site id.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PvsStreamingAPIsDAAS.PvsStreamingTestDeviceCollectionExists(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).DeviceCollectionNameCheckRequestModel(deviceCollectionNameCheckRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PvsStreamingAPIsDAAS.PvsStreamingTestDeviceCollectionExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPvsStreamingTestDeviceCollectionExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **deviceCollectionNameCheckRequestModel** | [**DeviceCollectionNameCheckRequestModel**](DeviceCollectionNameCheckRequestModel.md) | Name of the device collection and PVS site id. | 
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

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

