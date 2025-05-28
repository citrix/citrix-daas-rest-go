# \ProfilesDAAS

All URIs are relative to *https://wsaca.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupsInProfile**](ProfilesDAAS.md#AddGroupsInProfile) | **Patch** /aca/profile/{profileId}/groups$add | Add groups in existing Profile Record For given Profile ID.
[**CreateProfile**](ProfilesDAAS.md#CreateProfile) | **Post** /aca/profile | Create Profile Record For The Specified Customer Id.
[**DeleteProfile**](ProfilesDAAS.md#DeleteProfile) | **Delete** /aca/profile/{id} | Delete Profile Record For The Specified Customer Id and Profile Id.
[**GetAllProfile**](ProfilesDAAS.md#GetAllProfile) | **Get** /aca/profile | Retrieve all profile records for the given customer ID
[**GetById**](ProfilesDAAS.md#GetById) | **Get** /aca/profile/{profileId} | Retrieve all profile records for the given customer ID &amp; profile ID
[**GetByUrl**](ProfilesDAAS.md#GetByUrl) | **Get** /aca/profile/url/{url} | Retrieve all profile records for the given customer ID &amp; Service URL
[**RemoveGroupsInProfile**](ProfilesDAAS.md#RemoveGroupsInProfile) | **Patch** /aca/profile/{profileId}/groups$remove | Remove groups in existing Profile Record For given Profile ID.
[**UpdatePriorityInProfile**](ProfilesDAAS.md#UpdatePriorityInProfile) | **Patch** /aca/profile/{profileId}/priority$replace | Update Priority in existing Profile Record For given Profile ID.
[**Updateprofile**](ProfilesDAAS.md#Updateprofile) | **Put** /aca/profile/{profileId} | Update Profile Record For The Specified Customer Id, Profile ID.



## AddGroupsInProfile

> ProfileResponseModel AddGroupsInProfile(ctx, profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfilePatchGroupRequestModel(profilePatchGroupRequestModel).Execute()

Add groups in existing Profile Record For given Profile ID.

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
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Profile ID for which the profile record needs to be updated
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)
	profilePatchGroupRequestModel := *openapiclient.NewProfilePatchGroupRequestModel() // ProfilePatchGroupRequestModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesDAAS.AddGroupsInProfile(context.Background(), profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfilePatchGroupRequestModel(profilePatchGroupRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesDAAS.AddGroupsInProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupsInProfile`: ProfileResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ProfilesDAAS.AddGroupsInProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Profile ID for which the profile record needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupsInProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 
 **profilePatchGroupRequestModel** | [**ProfilePatchGroupRequestModel**](ProfilePatchGroupRequestModel.md) |  | 

### Return type

[**ProfileResponseModel**](ProfileResponseModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProfile

> ProfileResponseModel CreateProfile(ctx).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfileRequestModel(profileRequestModel).Execute()

Create Profile Record For The Specified Customer Id.

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
	profileRequestModel := *openapiclient.NewProfileRequestModel() // ProfileRequestModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesDAAS.CreateProfile(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfileRequestModel(profileRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesDAAS.CreateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProfile`: ProfileResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ProfilesDAAS.CreateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 
 **profileRequestModel** | [**ProfileRequestModel**](ProfileRequestModel.md) |  | 

### Return type

[**ProfileResponseModel**](ProfileResponseModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfile

> string DeleteProfile(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete Profile Record For The Specified Customer Id and Profile Id.

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Profile ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Citrix-TransactionId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesDAAS.DeleteProfile(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesDAAS.DeleteProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProfile`: string
	fmt.Fprintf(os.Stdout, "Response from `ProfilesDAAS.DeleteProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Citrix-TransactionId | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProfile

> GetAllProfileResponseModel GetAllProfile(ctx).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve all profile records for the given customer ID

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
	resp, r, err := apiClient.ProfilesDAAS.GetAllProfile(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesDAAS.GetAllProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProfile`: GetAllProfileResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ProfilesDAAS.GetAllProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllProfileResponseModel**](GetAllProfileResponseModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetById

> ProfileResponseModel GetById(ctx, profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve all profile records for the given customer ID & profile ID

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
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | profileId
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesDAAS.GetById(context.Background(), profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesDAAS.GetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetById`: ProfileResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ProfilesDAAS.GetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | profileId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**ProfileResponseModel**](ProfileResponseModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByUrl

> ProfileResponseModel GetByUrl(ctx, url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve all profile records for the given customer ID & Service URL

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
	url := "url_example" // string | url
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesDAAS.GetByUrl(context.Background(), url).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesDAAS.GetByUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByUrl`: ProfileResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ProfilesDAAS.GetByUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**url** | **string** | url | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**ProfileResponseModel**](ProfileResponseModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroupsInProfile

> ProfileResponseModel RemoveGroupsInProfile(ctx, profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfilePatchGroupRequestModel(profilePatchGroupRequestModel).Execute()

Remove groups in existing Profile Record For given Profile ID.

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
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Profile ID for which the profile record needs to be updated
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)
	profilePatchGroupRequestModel := *openapiclient.NewProfilePatchGroupRequestModel() // ProfilePatchGroupRequestModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesDAAS.RemoveGroupsInProfile(context.Background(), profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfilePatchGroupRequestModel(profilePatchGroupRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesDAAS.RemoveGroupsInProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveGroupsInProfile`: ProfileResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ProfilesDAAS.RemoveGroupsInProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Profile ID for which the profile record needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupsInProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 
 **profilePatchGroupRequestModel** | [**ProfilePatchGroupRequestModel**](ProfilePatchGroupRequestModel.md) |  | 

### Return type

[**ProfileResponseModel**](ProfileResponseModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePriorityInProfile

> ProfileResponseModel UpdatePriorityInProfile(ctx, profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfilePatchPriorityRequestModel(profilePatchPriorityRequestModel).Execute()

Update Priority in existing Profile Record For given Profile ID.

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
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Profile ID for which the profile record needs to be updated
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)
	profilePatchPriorityRequestModel := *openapiclient.NewProfilePatchPriorityRequestModel() // ProfilePatchPriorityRequestModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesDAAS.UpdatePriorityInProfile(context.Background(), profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfilePatchPriorityRequestModel(profilePatchPriorityRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesDAAS.UpdatePriorityInProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePriorityInProfile`: ProfileResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ProfilesDAAS.UpdatePriorityInProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Profile ID for which the profile record needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriorityInProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 
 **profilePatchPriorityRequestModel** | [**ProfilePatchPriorityRequestModel**](ProfilePatchPriorityRequestModel.md) |  | 

### Return type

[**ProfileResponseModel**](ProfileResponseModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Updateprofile

> interface Updateprofile(ctx, profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfileRequestModel(profileRequestModel).Execute()

Update Profile Record For The Specified Customer Id, Profile ID.

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
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Profile ID for which the profile record needs to be updated
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)
	profileRequestModel := *openapiclient.NewProfileRequestModel() // ProfileRequestModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesDAAS.Updateprofile(context.Background(), profileId).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).ProfileRequestModel(profileRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesDAAS.Updateprofile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Updateprofile`: interface
	fmt.Fprintf(os.Stdout, "Response from `ProfilesDAAS.Updateprofile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Profile ID for which the profile record needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateprofileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 
 **profileRequestModel** | [**ProfileRequestModel**](ProfileRequestModel.md) |  | 

### Return type

[**interface**](interface.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

