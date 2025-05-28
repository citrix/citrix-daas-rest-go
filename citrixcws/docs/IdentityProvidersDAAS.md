# \IdentityProvidersDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerIdentityProvidersAllStatusGet**](IdentityProvidersDAAS.md#CustomerIdentityProvidersAllStatusGet) | **Get** /{customer}/identityProviders/all/status | 
[**CustomerIdentityProvidersClientGet**](IdentityProvidersDAAS.md#CustomerIdentityProvidersClientGet) | **Get** /{customer}/identityProviders/client | 
[**CustomerIdentityProvidersConfigurationSamlIdGet**](IdentityProvidersDAAS.md#CustomerIdentityProvidersConfigurationSamlIdGet) | **Get** /{customer}/identityProviders/configuration/saml/{id} | 
[**CustomerIdentityProvidersConfigureGooglePost**](IdentityProvidersDAAS.md#CustomerIdentityProvidersConfigureGooglePost) | **Post** /{customer}/identityProviders/configure/google | 
[**CustomerIdentityProvidersConfigureOktaPost**](IdentityProvidersDAAS.md#CustomerIdentityProvidersConfigureOktaPost) | **Post** /{customer}/identityProviders/configure/okta | 
[**CustomerIdentityProvidersGet**](IdentityProvidersDAAS.md#CustomerIdentityProvidersGet) | **Get** /{customer}/identityProviders | 
[**CustomerIdentityProvidersIdPost**](IdentityProvidersDAAS.md#CustomerIdentityProvidersIdPost) | **Post** /{customer}/identityProviders/{id} | 
[**CustomerIdentityProvidersIdentityProviderIdPut**](IdentityProvidersDAAS.md#CustomerIdentityProvidersIdentityProviderIdPut) | **Put** /{customer}/identityProviders/{identityProvider}/{id} | 
[**CustomerIdentityProvidersIdpTypeGet**](IdentityProvidersDAAS.md#CustomerIdentityProvidersIdpTypeGet) | **Get** /{customer}/identityProviders/{idpType} | 
[**CustomerIdentityProvidersMultiIdentityProvidersPost**](IdentityProvidersDAAS.md#CustomerIdentityProvidersMultiIdentityProvidersPost) | **Post** /{customer}/identityProviders/multiIdentityProviders | 
[**CustomerIdentityProvidersPost**](IdentityProvidersDAAS.md#CustomerIdentityProvidersPost) | **Post** /{customer}/identityProviders | 
[**CustomerIdentityProvidersTypeIdDelete**](IdentityProvidersDAAS.md#CustomerIdentityProvidersTypeIdDelete) | **Delete** /{customer}/identityProviders/{type}/{id} | 
[**CustomerIdentityProvidersValidateClientsClientIdGet**](IdentityProvidersDAAS.md#CustomerIdentityProvidersValidateClientsClientIdGet) | **Get** /{customer}/identityProviders/validate/clients/{clientId} | 
[**CustomerIdentityProvidersValidateFqdnGet**](IdentityProvidersDAAS.md#CustomerIdentityProvidersValidateFqdnGet) | **Get** /{customer}/identityProviders/validate/fqdn | 
[**CustomerIdentityProvidersValidateSamlPost**](IdentityProvidersDAAS.md#CustomerIdentityProvidersValidateSamlPost) | **Post** /{customer}/identityProviders/validate/saml | 
[**CustomerIdentityprovidersIdentityProviderCreateconnectPost**](IdentityProvidersDAAS.md#CustomerIdentityprovidersIdentityProviderCreateconnectPost) | **Post** /{customer}/identityproviders/{identityProvider}/$createconnect | 



## CustomerIdentityProvidersAllStatusGet

> IdentityProvidersStatus CustomerIdentityProvidersAllStatusGet(ctx, customer).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	customer := "customer_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersAllStatusGet(context.Background(), customer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersAllStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersAllStatusGet`: IdentityProvidersStatus
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersAllStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersAllStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityProvidersStatus**](IdentityProvidersStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersClientGet

> CipClientResultModel CustomerIdentityProvidersClientGet(ctx, customer).Fqdn(fqdn).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	customer := "customer_example" // string | 
	fqdn := "fqdn_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersClientGet(context.Background(), customer).Fqdn(fqdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersClientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersClientGet`: CipClientResultModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersClientGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersClientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fqdn** | **string** |  | 

### Return type

[**CipClientResultModel**](CipClientResultModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersConfigurationSamlIdGet

> SamlConfigModel CustomerIdentityProvidersConfigurationSamlIdGet(ctx, id, customer).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	id := "id_example" // string | 
	customer := "customer_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersConfigurationSamlIdGet(context.Background(), id, customer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersConfigurationSamlIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersConfigurationSamlIdGet`: SamlConfigModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersConfigurationSamlIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersConfigurationSamlIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SamlConfigModel**](SamlConfigModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersConfigureGooglePost

> GoogleResultModel CustomerIdentityProvidersConfigureGooglePost(ctx, customer).GoogleConnectionModel(googleConnectionModel).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	customer := "customer_example" // string | 
	googleConnectionModel := *openapiclient.NewGoogleConnectionModel("GoogleClientEmail_example", "GooglePrivateKey_example", "GoogleImpersonatedUser_example") // GoogleConnectionModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersConfigureGooglePost(context.Background(), customer).GoogleConnectionModel(googleConnectionModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersConfigureGooglePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersConfigureGooglePost`: GoogleResultModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersConfigureGooglePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersConfigureGooglePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleConnectionModel** | [**GoogleConnectionModel**](GoogleConnectionModel.md) |  | 

### Return type

[**GoogleResultModel**](GoogleResultModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersConfigureOktaPost

> OktaResultModel CustomerIdentityProvidersConfigureOktaPost(ctx, customer).OktaConnectionModel(oktaConnectionModel).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	customer := "customer_example" // string | 
	oktaConnectionModel := *openapiclient.NewOktaConnectionModel() // OktaConnectionModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersConfigureOktaPost(context.Background(), customer).OktaConnectionModel(oktaConnectionModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersConfigureOktaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersConfigureOktaPost`: OktaResultModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersConfigureOktaPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersConfigureOktaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaConnectionModel** | [**OktaConnectionModel**](OktaConnectionModel.md) |  | 

### Return type

[**OktaResultModel**](OktaResultModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersGet

> []string CustomerIdentityProvidersGet(ctx, customer).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	customer := "customer_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersGet(context.Background(), customer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersIdPost

> IdpStatusModel CustomerIdentityProvidersIdPost(ctx, id, customer).IdpInstanceConnectModel(idpInstanceConnectModel).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	id := "id_example" // string | 
	customer := "customer_example" // string | 
	idpInstanceConnectModel := *openapiclient.NewIdpInstanceConnectModel("IdentityProviderType_example") // IdpInstanceConnectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersIdPost(context.Background(), id, customer).IdpInstanceConnectModel(idpInstanceConnectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersIdPost`: IdpStatusModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idpInstanceConnectModel** | [**IdpInstanceConnectModel**](IdpInstanceConnectModel.md) |  | 

### Return type

[**IdpStatusModel**](IdpStatusModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersIdentityProviderIdPut

> IdpUpdateModel CustomerIdentityProvidersIdentityProviderIdPut(ctx, identityProvider, id, customer).IdpUpdateModel(idpUpdateModel).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	identityProvider := "identityProvider_example" // string | 
	id := "id_example" // string | 
	customer := "customer_example" // string | 
	idpUpdateModel := *openapiclient.NewIdpUpdateModel() // IdpUpdateModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersIdentityProviderIdPut(context.Background(), identityProvider, id, customer).IdpUpdateModel(idpUpdateModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersIdentityProviderIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersIdentityProviderIdPut`: IdpUpdateModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersIdentityProviderIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProvider** | **string** |  | 
**id** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersIdentityProviderIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **idpUpdateModel** | [**IdpUpdateModel**](IdpUpdateModel.md) |  | 

### Return type

[**IdpUpdateModel**](IdpUpdateModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersIdpTypeGet

> IdpStatusesModel CustomerIdentityProvidersIdpTypeGet(ctx, idpType, customer).Limit(limit).ContinuationToken(continuationToken).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	idpType := "idpType_example" // string | 
	customer := "customer_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 30)
	continuationToken := "continuationToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersIdpTypeGet(context.Background(), idpType, customer).Limit(limit).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersIdpTypeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersIdpTypeGet`: IdpStatusesModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersIdpTypeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpType** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersIdpTypeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | [default to 30]
 **continuationToken** | **string** |  | 

### Return type

[**IdpStatusesModel**](IdpStatusesModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersMultiIdentityProvidersPost

> IdpStatusModel CustomerIdentityProvidersMultiIdentityProvidersPost(ctx, customer).IdpCreateModel(idpCreateModel).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	customer := "customer_example" // string | 
	idpCreateModel := *openapiclient.NewIdpCreateModel("IdentityProviderId_example", "IdentityProviderNickname_example") // IdpCreateModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersMultiIdentityProvidersPost(context.Background(), customer).IdpCreateModel(idpCreateModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersMultiIdentityProvidersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersMultiIdentityProvidersPost`: IdpStatusModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersMultiIdentityProvidersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersMultiIdentityProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idpCreateModel** | [**IdpCreateModel**](IdpCreateModel.md) |  | 

### Return type

[**IdpStatusModel**](IdpStatusModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersPost

> IdpStatusModel CustomerIdentityProvidersPost(ctx, customer).IdpConnectModel(idpConnectModel).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	customer := "customer_example" // string | 
	idpConnectModel := *openapiclient.NewIdpConnectModel("IdentityProviderId_example") // IdpConnectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersPost(context.Background(), customer).IdpConnectModel(idpConnectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersPost`: IdpStatusModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idpConnectModel** | [**IdpConnectModel**](IdpConnectModel.md) |  | 

### Return type

[**IdpStatusModel**](IdpStatusModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersTypeIdDelete

> IdpStatusModel CustomerIdentityProvidersTypeIdDelete(ctx, type_, id, customer).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	type_ := "type__example" // string | 
	id := "id_example" // string | 
	customer := "customer_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersTypeIdDelete(context.Background(), type_, id, customer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersTypeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersTypeIdDelete`: IdpStatusModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersTypeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**id** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersTypeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IdpStatusModel**](IdpStatusModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersValidateClientsClientIdGet

> bool CustomerIdentityProvidersValidateClientsClientIdGet(ctx, clientId, customer).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	clientId := "clientId_example" // string | 
	customer := "customer_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersValidateClientsClientIdGet(context.Background(), clientId, customer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersValidateClientsClientIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersValidateClientsClientIdGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersValidateClientsClientIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersValidateClientsClientIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersValidateFqdnGet

> bool CustomerIdentityProvidersValidateFqdnGet(ctx, customer).Fqdn(fqdn).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	customer := "customer_example" // string | 
	fqdn := "fqdn_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersValidateFqdnGet(context.Background(), customer).Fqdn(fqdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersValidateFqdnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersValidateFqdnGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersValidateFqdnGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersValidateFqdnGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fqdn** | **string** |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityProvidersValidateSamlPost

> SamlCertInfoModel CustomerIdentityProvidersValidateSamlPost(ctx, customer).FileName(fileName).CertFile(certFile).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	customer := "customer_example" // string | 
	fileName := "fileName_example" // string |  (optional)
	certFile := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityProvidersValidateSamlPost(context.Background(), customer).FileName(fileName).CertFile(certFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityProvidersValidateSamlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityProvidersValidateSamlPost`: SamlCertInfoModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityProvidersValidateSamlPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityProvidersValidateSamlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileName** | **string** |  | 
 **certFile** | ***os.File** |  | 

### Return type

[**SamlCertInfoModel**](SamlCertInfoModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerIdentityprovidersIdentityProviderCreateconnectPost

> IdpStatusModel CustomerIdentityprovidersIdentityProviderCreateconnectPost(ctx, identityProvider, customer).IdpInstanceCreateConnectModel(idpInstanceCreateConnectModel).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
	identityProvider := "identityProvider_example" // string | 
	customer := "customer_example" // string | 
	idpInstanceCreateConnectModel := *openapiclient.NewIdpInstanceCreateConnectModel("IdentityProviderType_example", "IdentityProviderNickname_example") // IdpInstanceCreateConnectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersDAAS.CustomerIdentityprovidersIdentityProviderCreateconnectPost(context.Background(), identityProvider, customer).IdpInstanceCreateConnectModel(idpInstanceCreateConnectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersDAAS.CustomerIdentityprovidersIdentityProviderCreateconnectPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerIdentityprovidersIdentityProviderCreateconnectPost`: IdpStatusModel
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersDAAS.CustomerIdentityprovidersIdentityProviderCreateconnectPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProvider** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerIdentityprovidersIdentityProviderCreateconnectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idpInstanceCreateConnectModel** | [**IdpInstanceCreateConnectModel**](IdpInstanceCreateConnectModel.md) |  | 

### Return type

[**IdpStatusModel**](IdpStatusModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

