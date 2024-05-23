# \SettingsControllerDAAS

All URIs are relative to *https://wsaca.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSettingsApiUsingDELETE**](SettingsControllerDAAS.md#DeleteSettingsApiUsingDELETE) | **Delete** /aca/settings/app/{app}/url/{url} | Delete Record For App And Service-url.
[**DeleteSettingsForChannel**](SettingsControllerDAAS.md#DeleteSettingsForChannel) | **Delete** /aca/settings/app/{app}/url/{url}/channel/{channelName} | Delete Settings Record For The Specified Customer Id, Product, Service And Channel.
[**GetAllSettingsApiUsingGET**](SettingsControllerDAAS.md#GetAllSettingsApiUsingGET) | **Get** /aca/settings | Retrieve All Settings Records For The Specified Customer Id.
[**GetSettingsApiUsingGET**](SettingsControllerDAAS.md#GetSettingsApiUsingGET) | **Get** /aca/settings/app/{app}/url/{url} | Retrieve Settings Record For App And Service-URL.
[**PostSettingsApiUsingPOST**](SettingsControllerDAAS.md#PostSettingsApiUsingPOST) | **Post** /aca/settings/app/{app}/url | Create A New Settings Record.
[**PutSettingsApiUsingPUT**](SettingsControllerDAAS.md#PutSettingsApiUsingPUT) | **Put** /aca/settings/app/{app}/url/{url} | Update Settings Record For An App And Service-URL.
[**RetrieveAllChannelSettings**](SettingsControllerDAAS.md#RetrieveAllChannelSettings) | **Get** /aca/settings/app/{app}/url/{url}/channel | Retrieve All Channels Settings Records For The Specified Customer Id and url
[**RetrieveSettingsForChannel**](SettingsControllerDAAS.md#RetrieveSettingsForChannel) | **Get** /aca/settings/app/{app}/url/{url}/channel/{channelName} | Retrieve Settings Record For The Specified Customer Id, Product, Service And Channel.



## DeleteSettingsApiUsingDELETE

> DeleteSettingsApiUsingDELETE(ctx, app, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete Record For App And Service-url.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | Application Name.
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsControllerDAAS.DeleteSettingsApiUsingDELETE(context.Background(), app, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsControllerDAAS.DeleteSettingsApiUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | Application Name. | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingsApiUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

 (empty response body)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSettingsForChannel

> DeleteSettingsForChannel(ctx, app, channelName, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete Settings Record For The Specified Customer Id, Product, Service And Channel.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	app := "app_example" // string | App Name. ['workspace']
	channelName := "channelName_example" // string | Channel Name
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsControllerDAAS.DeleteSettingsForChannel(context.Background(), app, channelName, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsControllerDAAS.DeleteSettingsForChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**channelName** | **string** | Channel Name | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingsForChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

 (empty response body)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSettingsApiUsingGET

> GetAllSettingResponse GetAllSettingsApiUsingGET(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve All Settings Records For The Specified Customer Id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsControllerDAAS.GetAllSettingsApiUsingGET(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsControllerDAAS.GetAllSettingsApiUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSettingsApiUsingGET`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsControllerDAAS.GetAllSettingsApiUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSettingsApiUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsApiUsingGET

> GetAllSettingResponse GetSettingsApiUsingGET(ctx, app, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve Settings Record For App And Service-URL.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['Workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsControllerDAAS.GetSettingsApiUsingGET(context.Background(), app, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsControllerDAAS.GetSettingsApiUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsApiUsingGET`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsControllerDAAS.GetSettingsApiUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;Workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsApiUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsApiUsingPOST

> GetAllSettingResponse PostSettingsApiUsingPOST(ctx, app).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).SettingsRecord(settingsRecord).CitrixTransactionId(citrixTransactionId).Execute()

Create A New Settings Record.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | Application Name.
	settingsRecord := *openapiclient.NewSettingsRecordModel() // SettingsRecordModel | Settings record to be created.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsControllerDAAS.PostSettingsApiUsingPOST(context.Background(), app).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).SettingsRecord(settingsRecord).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsControllerDAAS.PostSettingsApiUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSettingsApiUsingPOST`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsControllerDAAS.PostSettingsApiUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | Application Name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsApiUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **settingsRecord** | [**SettingsRecordModel**](SettingsRecordModel.md) | Settings record to be created. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSettingsApiUsingPUT

> PutSettingsApiUsingPUT(ctx, app, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).SettingsRecord(settingsRecord).CitrixTransactionId(citrixTransactionId).Execute()

Update Settings Record For An App And Service-URL.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['Workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	settingsRecord := *openapiclient.NewSettingsRecordModel() // SettingsRecordModel | Settings record to be created.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsControllerDAAS.PutSettingsApiUsingPUT(context.Background(), app, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).SettingsRecord(settingsRecord).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsControllerDAAS.PutSettingsApiUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;Workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSettingsApiUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **settingsRecord** | [**SettingsRecordModel**](SettingsRecordModel.md) | Settings record to be created. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

 (empty response body)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllChannelSettings

> GetAllSettingResponse RetrieveAllChannelSettings(ctx, app, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve All Channels Settings Records For The Specified Customer Id and url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	app := "app_example" // string | App Name. ['workspace']
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsControllerDAAS.RetrieveAllChannelSettings(context.Background(), app, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsControllerDAAS.RetrieveAllChannelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAllChannelSettings`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsControllerDAAS.RetrieveAllChannelSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllChannelSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSettingsForChannel

> GetAllSettingResponse RetrieveSettingsForChannel(ctx, app, channelName, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve Settings Record For The Specified Customer Id, Product, Service And Channel.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	app := "app_example" // string | App Name. ['workspace']
	channelName := "channelName_example" // string | Channel Name
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsControllerDAAS.RetrieveSettingsForChannel(context.Background(), app, channelName, url).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsControllerDAAS.RetrieveSettingsForChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveSettingsForChannel`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsControllerDAAS.RetrieveSettingsForChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**channelName** | **string** | Channel Name | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSettingsForChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

