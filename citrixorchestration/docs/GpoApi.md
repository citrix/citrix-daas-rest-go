# \GpoApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GpoComparePolicies**](GpoApi.md#GpoComparePolicies) | **Post** /techpreview/{customerId}/{siteId}/gpo/comparePolicies | Compare policies.
[**GpoCreateGpoFilter**](GpoApi.md#GpoCreateGpoFilter) | **Post** /techpreview/{customerId}/{siteId}/gpo/filters | Create a filter in a policy.
[**GpoCreateGpoPolicy**](GpoApi.md#GpoCreateGpoPolicy) | **Post** /techpreview/{customerId}/{siteId}/gpo/policies | Create a new policy.
[**GpoCreateGpoPolicySet**](GpoApi.md#GpoCreateGpoPolicySet) | **Post** /techpreview/{customerId}/{siteId}/gpo/policySets | Create a new GPO policy set.
[**GpoCreateGpoSetting**](GpoApi.md#GpoCreateGpoSetting) | **Post** /techpreview/{customerId}/{siteId}/gpo/settings | Create a setting in a policy.
[**GpoDeleteGpoFilter**](GpoApi.md#GpoDeleteGpoFilter) | **Delete** /techpreview/{customerId}/{siteId}/gpo/filters/{filterGuid} | Delete an existing filter.
[**GpoDeleteGpoPolicy**](GpoApi.md#GpoDeleteGpoPolicy) | **Delete** /techpreview/{customerId}/{siteId}/gpo/policies/{policyGuid} | Delete an existing GPO policy.
[**GpoDeleteGpoPolicySet**](GpoApi.md#GpoDeleteGpoPolicySet) | **Delete** /techpreview/{customerId}/{siteId}/gpo/policySets/{policySetGuid} | Delete an existing GPO policy set. Policies in the policy set are deleted if a policy set is deleted.
[**GpoDeleteGpoSetting**](GpoApi.md#GpoDeleteGpoSetting) | **Delete** /techpreview/{customerId}/{siteId}/gpo/settings/{settingGuid} | Delete a setting.
[**GpoGetFilterDefinitions**](GpoApi.md#GpoGetFilterDefinitions) | **Get** /techpreview/{customerId}/{siteId}/gpo/filterDefinitions | Generate all filter definitions.
[**GpoGetSettingDefinitions**](GpoApi.md#GpoGetSettingDefinitions) | **Get** /techpreview/{customerId}/{siteId}/gpo/settingDefinitions | Get setting definitions. If isLean is set to true, only basic session information is returned. EnumType, VdaVersions, VersionDetails, and Explanation are not retrieved. If limit is set to -1 or a number larger than the number of settings available, all entries are retrieved. If limit is set to a positive integer smaller than the number of settings available, the specified number of settings are retrieved.
[**GpoGetSettingFullDetail**](GpoApi.md#GpoGetSettingFullDetail) | **Get** /techpreview/{customerId}/{siteId}/gpo/settingFullDetail | Get full detail of a setting.
[**GpoRankGpoPolicies**](GpoApi.md#GpoRankGpoPolicies) | **Post** /techpreview/{customerId}/{siteId}/gpo/policyPriorities | Specify new priority order for all existing policies in a policy set.
[**GpoReadGpoFilter**](GpoApi.md#GpoReadGpoFilter) | **Get** /techpreview/{customerId}/{siteId}/gpo/filters/{filterGuid} | Read a specific filter.
[**GpoReadGpoFilters**](GpoApi.md#GpoReadGpoFilters) | **Get** /techpreview/{customerId}/{siteId}/gpo/filters | Read filters defined in a policy.
[**GpoReadGpoPolicies**](GpoApi.md#GpoReadGpoPolicies) | **Get** /techpreview/{customerId}/{siteId}/gpo/policies | Read all policies defined in a policy set.
[**GpoReadGpoPolicy**](GpoApi.md#GpoReadGpoPolicy) | **Get** /techpreview/{customerId}/{siteId}/gpo/policies/{policyGuid} | Read a policy.
[**GpoReadGpoPolicySet**](GpoApi.md#GpoReadGpoPolicySet) | **Get** /techpreview/{customerId}/{siteId}/gpo/policySets/{policySetGuid} | Read a GPO policy set.
[**GpoReadGpoPolicySets**](GpoApi.md#GpoReadGpoPolicySets) | **Get** /techpreview/{customerId}/{siteId}/gpo/policySets | Get all GPO policy sets in the site.
[**GpoReadGpoSetting**](GpoApi.md#GpoReadGpoSetting) | **Get** /techpreview/{customerId}/{siteId}/gpo/settings/{settingGuid} | Read a specific setting.
[**GpoReadGpoSettings**](GpoApi.md#GpoReadGpoSettings) | **Get** /techpreview/{customerId}/{siteId}/gpo/settings | Read settings defined in a policy.
[**GpoUpdateGpoFilter**](GpoApi.md#GpoUpdateGpoFilter) | **Patch** /techpreview/{customerId}/{siteId}/gpo/filters/{filterGuid} | Update an existing filter.
[**GpoUpdateGpoPolicy**](GpoApi.md#GpoUpdateGpoPolicy) | **Patch** /techpreview/{customerId}/{siteId}/gpo/policies/{policyGuid} | Update a policy. Only the policy body is updated.
[**GpoUpdateGpoPolicySet**](GpoApi.md#GpoUpdateGpoPolicySet) | **Patch** /techpreview/{customerId}/{siteId}/gpo/policySets/{policySetGuid} | Update an existing GPO policy set.
[**GpoUpdateGpoSetting**](GpoApi.md#GpoUpdateGpoSetting) | **Patch** /techpreview/{customerId}/{siteId}/gpo/settings/{settingGuid} | Update a setting.
[**GpoUpdatePolicySetBlob**](GpoApi.md#GpoUpdatePolicySetBlob) | **Put** /techpreview/{customerId}/{siteId}/gpo/policySets/{policySetGuid} | Update the internal data blob for a policy set to the latest data in the set.



## GpoComparePolicies

> ComparisonResponseContract GpoComparePolicies(ctx, customerId, siteId).WithDefaults(withDefaults).Targets(targets).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Compare policies.

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
    customerId := "customerId_example" // string | The customer id of the customer making this rest call
    siteId := "siteId_example" // string | The site id of the customer's site making this rest call
    withDefaults := true // bool | Include defaults in comparison
    targets := []string{"Property_example"} // []string | Guids of policies
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoComparePolicies(context.Background(), customerId, siteId).WithDefaults(withDefaults).Targets(targets).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoComparePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoComparePolicies`: ComparisonResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoComparePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer id of the customer making this rest call | 
**siteId** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoComparePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withDefaults** | **bool** | Include defaults in comparison | 
 **targets** | **[]string** | Guids of policies | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ComparisonResponseContract**](ComparisonResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoCreateGpoFilter

> FilterResponse GpoCreateGpoFilter(ctx, customerId, siteId).PolicyGuid(policyGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a filter in a policy.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policyGuid := "policyGuid_example" // string | The GUID of the policy to which the new filter belongs
    request := *openapiclient.NewFilterRequest(false, false) // FilterRequest | Filter data
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoCreateGpoFilter(context.Background(), customerId, siteId).PolicyGuid(policyGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoCreateGpoFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoCreateGpoFilter`: FilterResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoCreateGpoFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoCreateGpoFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyGuid** | **string** | The GUID of the policy to which the new filter belongs | 
 **request** | [**FilterRequest**](FilterRequest.md) | Filter data | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**FilterResponse**](FilterResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoCreateGpoPolicy

> PolicyResponse GpoCreateGpoPolicy(ctx, customerId, siteId).PolicySetGuid(policySetGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a new policy.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policySetGuid := "policySetGuid_example" // string | GUID of the policy set in which the new policy is created
    request := *openapiclient.NewPolicyRequest(false) // PolicyRequest | Data for the new policy
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoCreateGpoPolicy(context.Background(), customerId, siteId).PolicySetGuid(policySetGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoCreateGpoPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoCreateGpoPolicy`: PolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoCreateGpoPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoCreateGpoPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policySetGuid** | **string** | GUID of the policy set in which the new policy is created | 
 **request** | [**PolicyRequest**](PolicyRequest.md) | Data for the new policy | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PolicyResponse**](PolicyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoCreateGpoPolicySet

> PolicySetResponse GpoCreateGpoPolicySet(ctx, customerId, siteId).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a new GPO policy set.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    request := *openapiclient.NewPolicySetRequest() // PolicySetRequest | Parameters for the new policy set
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoCreateGpoPolicySet(context.Background(), customerId, siteId).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoCreateGpoPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoCreateGpoPolicySet`: PolicySetResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoCreateGpoPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoCreateGpoPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**PolicySetRequest**](PolicySetRequest.md) | Parameters for the new policy set | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PolicySetResponse**](PolicySetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoCreateGpoSetting

> SettingResponse GpoCreateGpoSetting(ctx, customerId, siteId).PolicyGuid(policyGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a setting in a policy.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policyGuid := "policyGuid_example" // string | GUID of the policy to which the setting belongs
    request := *openapiclient.NewSettingRequest(false) // SettingRequest | Data for the new setting
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoCreateGpoSetting(context.Background(), customerId, siteId).PolicyGuid(policyGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoCreateGpoSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoCreateGpoSetting`: SettingResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoCreateGpoSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoCreateGpoSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyGuid** | **string** | GUID of the policy to which the setting belongs | 
 **request** | [**SettingRequest**](SettingRequest.md) | Data for the new setting | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingResponse**](SettingResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoDeleteGpoFilter

> GpoDeleteGpoFilter(ctx, customerId, siteId, filterGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an existing filter.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    filterGuid := "filterGuid_example" // string | The GUID of the filter to be deleted
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GpoApi.GpoDeleteGpoFilter(context.Background(), customerId, siteId, filterGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoDeleteGpoFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**filterGuid** | **string** | The GUID of the filter to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoDeleteGpoFilterRequest struct via the builder pattern


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


## GpoDeleteGpoPolicy

> GpoDeleteGpoPolicy(ctx, customerId, siteId, policyGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an existing GPO policy.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policyGuid := "policyGuid_example" // string | GUID of the policy to be deleted
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GpoApi.GpoDeleteGpoPolicy(context.Background(), customerId, siteId, policyGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoDeleteGpoPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**policyGuid** | **string** | GUID of the policy to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoDeleteGpoPolicyRequest struct via the builder pattern


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


## GpoDeleteGpoPolicySet

> GpoDeleteGpoPolicySet(ctx, customerId, siteId, policySetGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an existing GPO policy set. Policies in the policy set are deleted if a policy set is deleted.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policySetGuid := "policySetGuid_example" // string | The GUID of the policy set to be deleted
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GpoApi.GpoDeleteGpoPolicySet(context.Background(), customerId, siteId, policySetGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoDeleteGpoPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**policySetGuid** | **string** | The GUID of the policy set to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoDeleteGpoPolicySetRequest struct via the builder pattern


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


## GpoDeleteGpoSetting

> GpoDeleteGpoSetting(ctx, customerId, siteId, settingGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a setting.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    settingGuid := "settingGuid_example" // string | GUID of the setting to be deleted
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GpoApi.GpoDeleteGpoSetting(context.Background(), customerId, siteId, settingGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoDeleteGpoSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**settingGuid** | **string** | GUID of the setting to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoDeleteGpoSettingRequest struct via the builder pattern


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


## GpoGetFilterDefinitions

> CollectionEnvelopeOfFilterDefinition GpoGetFilterDefinitions(ctx, customerId, siteId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Generate all filter definitions.

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
    customerId := "customerId_example" // string | 
    siteId := "siteId_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoGetFilterDefinitions(context.Background(), customerId, siteId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoGetFilterDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoGetFilterDefinitions`: CollectionEnvelopeOfFilterDefinition
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoGetFilterDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoGetFilterDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**CollectionEnvelopeOfFilterDefinition**](CollectionEnvelopeOfFilterDefinition.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoGetSettingDefinitions

> SettingDefinitionEnvelope GpoGetSettingDefinitions(ctx, customerId, siteId).IsLean(isLean).Limit(limit).IsAscending(isAscending).NamePattern(namePattern).IsUserSetting(isUserSetting).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get setting definitions. If isLean is set to true, only basic session information is returned. EnumType, VdaVersions, VersionDetails, and Explanation are not retrieved. If limit is set to -1 or a number larger than the number of settings available, all entries are retrieved. If limit is set to a positive integer smaller than the number of settings available, the specified number of settings are retrieved.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    isLean := true // bool | Get lean parts of setting definitions, the default is set to true (optional) (default to true)
    limit := int32(56) // int32 | Specify the number of entries to retrieve, the default is all entries (optional) (default to -1)
    isAscending := true // bool | Specify sort order, default is true (optional) (default to true)
    namePattern := "namePattern_example" // string | Specify a regular expression to match the internal setting name. The default is match all names.              (optional)
    isUserSetting := true // bool | Specify the target of applying the settings. If it's set to true, only user settings are retrieved. If it's set to false, only computer settings are retrieved. If not specified, both kinds of settings are retrieved. The default is to retrieve both kinds of settings.              (optional)
    continuationToken := "continuationToken_example" // string | Continuation token from a previous retrieval (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoGetSettingDefinitions(context.Background(), customerId, siteId).IsLean(isLean).Limit(limit).IsAscending(isAscending).NamePattern(namePattern).IsUserSetting(isUserSetting).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoGetSettingDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoGetSettingDefinitions`: SettingDefinitionEnvelope
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoGetSettingDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoGetSettingDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isLean** | **bool** | Get lean parts of setting definitions, the default is set to true | [default to true]
 **limit** | **int32** | Specify the number of entries to retrieve, the default is all entries | [default to -1]
 **isAscending** | **bool** | Specify sort order, default is true | [default to true]
 **namePattern** | **string** | Specify a regular expression to match the internal setting name. The default is match all names.              | 
 **isUserSetting** | **bool** | Specify the target of applying the settings. If it&#39;s set to true, only user settings are retrieved. If it&#39;s set to false, only computer settings are retrieved. If not specified, both kinds of settings are retrieved. The default is to retrieve both kinds of settings.              | 
 **continuationToken** | **string** | Continuation token from a previous retrieval | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingDefinitionEnvelope**](SettingDefinitionEnvelope.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoGetSettingFullDetail

> SettingDefinition GpoGetSettingFullDetail(ctx, customerId, siteId).SettingName(settingName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get full detail of a setting.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    settingName := "settingName_example" // string | Setting name
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoGetSettingFullDetail(context.Background(), customerId, siteId).SettingName(settingName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoGetSettingFullDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoGetSettingFullDetail`: SettingDefinition
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoGetSettingFullDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoGetSettingFullDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **settingName** | **string** | Setting name | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingDefinition**](SettingDefinition.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoRankGpoPolicies

> bool GpoRankGpoPolicies(ctx, customerId, siteId).PolicySetGuid(policySetGuid).PolicyGuids(policyGuids).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Specify new priority order for all existing policies in a policy set.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policySetGuid := "policySetGuid_example" // string | Guid of the policy set that contains the policies
    policyGuids := []string{"Property_example"} // []string | Guids of the policies to be re-prioritized
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoRankGpoPolicies(context.Background(), customerId, siteId).PolicySetGuid(policySetGuid).PolicyGuids(policyGuids).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoRankGpoPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoRankGpoPolicies`: bool
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoRankGpoPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoRankGpoPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policySetGuid** | **string** | Guid of the policy set that contains the policies | 
 **policyGuids** | **[]string** | Guids of the policies to be re-prioritized | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoReadGpoFilter

> FilterResponse GpoReadGpoFilter(ctx, customerId, siteId, filterGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Read a specific filter.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    filterGuid := "filterGuid_example" // string | The GUID of the filter to be read
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoReadGpoFilter(context.Background(), customerId, siteId, filterGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoReadGpoFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoReadGpoFilter`: FilterResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoReadGpoFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**filterGuid** | **string** | The GUID of the filter to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**FilterResponse**](FilterResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoReadGpoFilters

> CollectionEnvelopeOfFilterResponse GpoReadGpoFilters(ctx, customerId, siteId).PolicyGuid(policyGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Read filters defined in a policy.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policyGuid := "policyGuid_example" // string | The GUID of the policy from which filters are read
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoReadGpoFilters(context.Background(), customerId, siteId).PolicyGuid(policyGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoReadGpoFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoReadGpoFilters`: CollectionEnvelopeOfFilterResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoReadGpoFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyGuid** | **string** | The GUID of the policy from which filters are read | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**CollectionEnvelopeOfFilterResponse**](CollectionEnvelopeOfFilterResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoReadGpoPolicies

> CollectionEnvelopeOfPolicyResponse GpoReadGpoPolicies(ctx, customerId, siteId).PolicySetGuid(policySetGuid).WithSettings(withSettings).WithFilters(withFilters).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Read all policies defined in a policy set.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policySetGuid := "policySetGuid_example" // string | The GUID of the policy set from which policies are read
    withSettings := true // bool | If set to true, settings in the policy are read (optional)
    withFilters := true // bool | If set to true, filters in the policy are read (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoReadGpoPolicies(context.Background(), customerId, siteId).PolicySetGuid(policySetGuid).WithSettings(withSettings).WithFilters(withFilters).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoReadGpoPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoReadGpoPolicies`: CollectionEnvelopeOfPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoReadGpoPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policySetGuid** | **string** | The GUID of the policy set from which policies are read | 
 **withSettings** | **bool** | If set to true, settings in the policy are read | 
 **withFilters** | **bool** | If set to true, filters in the policy are read | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**CollectionEnvelopeOfPolicyResponse**](CollectionEnvelopeOfPolicyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoReadGpoPolicy

> PolicyResponse GpoReadGpoPolicy(ctx, customerId, siteId, policyGuid).WithSettings(withSettings).WithFilters(withFilters).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Read a policy.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policyGuid := "policyGuid_example" // string | GUID of the policy to be read
    withSettings := true // bool | If set to true, read policy settings (optional)
    withFilters := true // bool | If set to true, read policy filters (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoReadGpoPolicy(context.Background(), customerId, siteId, policyGuid).WithSettings(withSettings).WithFilters(withFilters).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoReadGpoPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoReadGpoPolicy`: PolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoReadGpoPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**policyGuid** | **string** | GUID of the policy to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **withSettings** | **bool** | If set to true, read policy settings | 
 **withFilters** | **bool** | If set to true, read policy filters | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PolicyResponse**](PolicyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoReadGpoPolicySet

> PolicySetResponse GpoReadGpoPolicySet(ctx, customerId, siteId, policySetGuid).WithPolicies(withPolicies).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Read a GPO policy set.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policySetGuid := "policySetGuid_example" // string | GUID of the policy set to read
    withPolicies := true // bool | If set to true, read the policies in the policy set (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoReadGpoPolicySet(context.Background(), customerId, siteId, policySetGuid).WithPolicies(withPolicies).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoReadGpoPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoReadGpoPolicySet`: PolicySetResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoReadGpoPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**policySetGuid** | **string** | GUID of the policy set to read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **withPolicies** | **bool** | If set to true, read the policies in the policy set | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PolicySetResponse**](PolicySetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoReadGpoPolicySets

> CollectionEnvelopeOfPolicySetResponse GpoReadGpoPolicySets(ctx, customerId, siteId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all GPO policy sets in the site.

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
    customerId := "customerId_example" // string | 
    siteId := "siteId_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoReadGpoPolicySets(context.Background(), customerId, siteId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoReadGpoPolicySets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoReadGpoPolicySets`: CollectionEnvelopeOfPolicySetResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoReadGpoPolicySets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoPolicySetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**CollectionEnvelopeOfPolicySetResponse**](CollectionEnvelopeOfPolicySetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoReadGpoSetting

> SettingResponse GpoReadGpoSetting(ctx, customerId, siteId, settingGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Read a specific setting.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    settingGuid := "settingGuid_example" // string | GUID of the setting to be read
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoReadGpoSetting(context.Background(), customerId, siteId, settingGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoReadGpoSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoReadGpoSetting`: SettingResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoReadGpoSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**settingGuid** | **string** | GUID of the setting to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingResponse**](SettingResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoReadGpoSettings

> CollectionEnvelopeOfSettingResponse GpoReadGpoSettings(ctx, customerId, siteId).PolicyGuid(policyGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Read settings defined in a policy.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policyGuid := "policyGuid_example" // string | GUID of the policy from which settings are read
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GpoApi.GpoReadGpoSettings(context.Background(), customerId, siteId).PolicyGuid(policyGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoReadGpoSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GpoReadGpoSettings`: CollectionEnvelopeOfSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `GpoApi.GpoReadGpoSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyGuid** | **string** | GUID of the policy from which settings are read | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**CollectionEnvelopeOfSettingResponse**](CollectionEnvelopeOfSettingResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoUpdateGpoFilter

> GpoUpdateGpoFilter(ctx, customerId, siteId, filterGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an existing filter.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    filterGuid := "filterGuid_example" // string | The GUID of the filter to be updated
    request := *openapiclient.NewFilterRequest(false, false) // FilterRequest | Filter data. The filter type is ignored
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GpoApi.GpoUpdateGpoFilter(context.Background(), customerId, siteId, filterGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoUpdateGpoFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**filterGuid** | **string** | The GUID of the filter to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdateGpoFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**FilterRequest**](FilterRequest.md) | Filter data. The filter type is ignored | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## GpoUpdateGpoPolicy

> GpoUpdateGpoPolicy(ctx, customerId, siteId, policyGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a policy. Only the policy body is updated.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policyGuid := "policyGuid_example" // string | GUID of the policy to be updated
    request := *openapiclient.NewPolicyBodyRequest(false) // PolicyBodyRequest | New policy data
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GpoApi.GpoUpdateGpoPolicy(context.Background(), customerId, siteId, policyGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoUpdateGpoPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**policyGuid** | **string** | GUID of the policy to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdateGpoPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**PolicyBodyRequest**](PolicyBodyRequest.md) | New policy data | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## GpoUpdateGpoPolicySet

> GpoUpdateGpoPolicySet(ctx, customerId, siteId, policySetGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an existing GPO policy set.



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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policySetGuid := "policySetGuid_example" // string | The GUID of the policy set to update
    request := *openapiclient.NewPolicySetRequest() // PolicySetRequest | New description of the policy set
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GpoApi.GpoUpdateGpoPolicySet(context.Background(), customerId, siteId, policySetGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoUpdateGpoPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**policySetGuid** | **string** | The GUID of the policy set to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdateGpoPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**PolicySetRequest**](PolicySetRequest.md) | New description of the policy set | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## GpoUpdateGpoSetting

> GpoUpdateGpoSetting(ctx, customerId, siteId, settingGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a setting.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    settingGuid := "settingGuid_example" // string | GUID of the setting to be updated
    request := *openapiclient.NewSettingRequest(false) // SettingRequest | New setting value
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GpoApi.GpoUpdateGpoSetting(context.Background(), customerId, siteId, settingGuid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoUpdateGpoSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**settingGuid** | **string** | GUID of the setting to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdateGpoSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**SettingRequest**](SettingRequest.md) | New setting value | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## GpoUpdatePolicySetBlob

> GpoUpdatePolicySetBlob(ctx, customerId, siteId, policySetGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update the internal data blob for a policy set to the latest data in the set.

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
    customerId := "customerId_example" // string | The customer ID
    siteId := "siteId_example" // string | The site ID
    policySetGuid := "policySetGuid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GpoApi.GpoUpdatePolicySetBlob(context.Background(), customerId, siteId, policySetGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GpoApi.GpoUpdatePolicySetBlob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**siteId** | **string** | The site ID | 
**policySetGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdatePolicySetBlobRequest struct via the builder pattern


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

