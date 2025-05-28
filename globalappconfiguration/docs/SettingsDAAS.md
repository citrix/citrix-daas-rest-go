# \SettingsDAAS

All URIs are relative to *https://wsaca.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSettings**](SettingsDAAS.md#CreateSettings) | **Post** /aca/settings/app/{app}/url | Create Settings Record For The Specified Customer Id For A Specific Product.
[**DeleteSettings**](SettingsDAAS.md#DeleteSettings) | **Delete** /aca/settings/app/{app}/url/{url} | Delete Settings Record For The Specified Customer Id, Product And Service.
[**DeleteSettingsForChannel**](SettingsDAAS.md#DeleteSettingsForChannel) | **Delete** /aca/settings/app/{app}/url/{url}/channel/{channelName} | Delete Settings Record For The Specified Customer Id, Product, Service And Channel.
[**DeleteSettingsForProfile**](SettingsDAAS.md#DeleteSettingsForProfile) | **Delete** /aca/settings/app/{app}/url/{url}/profile/{profileId} | Delete Settings Record For The Specified Customer Id, Product, Service And Profile.
[**RetrieveAllChannelSettings**](SettingsDAAS.md#RetrieveAllChannelSettings) | **Get** /aca/settings/app/{app}/url/{url}/channel | Retrieve All Channels Settings Records For The Specified Customer Id and url
[**RetrieveAllProfileSettings**](SettingsDAAS.md#RetrieveAllProfileSettings) | **Get** /aca/settings/app/{app}/url/{url}/profile | Retrieve All Profile Settings Records For The Specified Customer Id and url
[**RetrieveAllSettings**](SettingsDAAS.md#RetrieveAllSettings) | **Get** /aca/settings | Retrieve Settings Records For The Specified Customer Id.
[**RetrieveSettingsForChannel**](SettingsDAAS.md#RetrieveSettingsForChannel) | **Get** /aca/settings/app/{app}/url/{url}/channel/{channelName} | Retrieve Settings Record For The Specified Customer Id, Product, Service And Channel.
[**RetrieveSettingsForProfile**](SettingsDAAS.md#RetrieveSettingsForProfile) | **Get** /aca/settings/app/{app}/url/{url}/profile/{profileId} | Retrieve Settings Record For The Specified Customer Id, Product, Service And Profile ID.
[**RetrieveSettingsForURL**](SettingsDAAS.md#RetrieveSettingsForURL) | **Get** /aca/settings/app/{app}/url/{url} | Retrieve Settings Record For The Specified Customer Id, Product And Service.
[**UpdateSettings**](SettingsDAAS.md#UpdateSettings) | **Put** /aca/settings/app/{app}/url/{url} | Update Settings Record For The Specified Customer Id, Product And Service.



## CreateSettings

> GetAllSettingResponse CreateSettings(ctx, app).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).SettingsRecordModel(settingsRecordModel).Execute()

Create Settings Record For The Specified Customer Id For A Specific Product.

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)
	settingsRecordModel := *openapiclient.NewSettingsRecordModel() // SettingsRecordModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.CreateSettings(context.Background(), app).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).SettingsRecordModel(settingsRecordModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.CreateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSettings`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.CreateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 
 **settingsRecordModel** | [**SettingsRecordModel**](SettingsRecordModel.md) |  | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSettings

> string DeleteSettings(ctx, app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete Settings Record For The Specified Customer Id, Product And Service.

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.DeleteSettings(context.Background(), app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.DeleteSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSettings`: string
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.DeleteSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSettingsForChannel

> string DeleteSettingsForChannel(ctx, app, url, channelName).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	channelName := "channelName_example" // string | Channel Name
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.DeleteSettingsForChannel(context.Background(), app, url, channelName).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.DeleteSettingsForChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSettingsForChannel`: string
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.DeleteSettingsForChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 
**channelName** | **string** | Channel Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingsForChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 



 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSettingsForProfile

> string DeleteSettingsForProfile(ctx, app, url, profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete Settings Record For The Specified Customer Id, Product, Service And Profile.

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Profile ID
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.DeleteSettingsForProfile(context.Background(), app, url, profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.DeleteSettingsForProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSettingsForProfile`: string
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.DeleteSettingsForProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingsForProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 



 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllChannelSettings

> GetAllSettingResponse RetrieveAllChannelSettings(ctx, app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.RetrieveAllChannelSettings(context.Background(), app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.RetrieveAllChannelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAllChannelSettings`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.RetrieveAllChannelSettings`: %v\n", resp)
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
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllProfileSettings

> GetAllSettingResponse RetrieveAllProfileSettings(ctx, app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve All Profile Settings Records For The Specified Customer Id and url

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.RetrieveAllProfileSettings(context.Background(), app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.RetrieveAllProfileSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAllProfileSettings`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.RetrieveAllProfileSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllProfileSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllSettings

> GetAllSettingResponse RetrieveAllSettings(ctx).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve Settings Records For The Specified Customer Id.

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.RetrieveAllSettings(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.RetrieveAllSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAllSettings`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.RetrieveAllSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSettingsForChannel

> GetAllSettingResponse RetrieveSettingsForChannel(ctx, app, url, channelName).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	channelName := "channelName_example" // string | Channel Name
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.RetrieveSettingsForChannel(context.Background(), app, url, channelName).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.RetrieveSettingsForChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveSettingsForChannel`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.RetrieveSettingsForChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 
**channelName** | **string** | Channel Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSettingsForChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 



 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSettingsForProfile

> GetAllSettingResponse RetrieveSettingsForProfile(ctx, app, url, profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve Settings Record For The Specified Customer Id, Product, Service And Profile ID.

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Profile ID
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.RetrieveSettingsForProfile(context.Background(), app, url, profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.RetrieveSettingsForProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveSettingsForProfile`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.RetrieveSettingsForProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSettingsForProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 



 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSettingsForURL

> GetAllSettingResponse RetrieveSettingsForURL(ctx, app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve Settings Record For The Specified Customer Id, Product And Service.

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.RetrieveSettingsForURL(context.Background(), app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.RetrieveSettingsForURL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveSettingsForURL`: GetAllSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.RetrieveSettingsForURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSettingsForURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllSettingResponse**](GetAllSettingResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettings

> string UpdateSettings(ctx, app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).SettingsRecordModel(settingsRecordModel).Execute()

Update Settings Record For The Specified Customer Id, Product And Service.

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
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	url := "url_example" // string | Service URL for settings (based 64 url encoded)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)
	settingsRecordModel := *openapiclient.NewSettingsRecordModel() // SettingsRecordModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsDAAS.UpdateSettings(context.Background(), app, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).SettingsRecordModel(settingsRecordModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsDAAS.UpdateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSettings`: string
	fmt.Fprintf(os.Stdout, "Response from `SettingsDAAS.UpdateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**url** | **string** | Service URL for settings (based 64 url encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 
 **settingsRecordModel** | [**SettingsRecordModel**](SettingsRecordModel.md) |  | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

