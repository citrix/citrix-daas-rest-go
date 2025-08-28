# \GpoDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GpoComparePolicies**](GpoDAAS.md#GpoComparePolicies) | **Post** /gpo/comparePolicies | Compare policies. The policies can be from different policy sets.
[**GpoCopyGpoPolicies**](GpoDAAS.md#GpoCopyGpoPolicies) | **Post** /gpo/policies/$copy | Copy some policies of a policy set to other policy sets.
[**GpoCopyGpoPolicySet**](GpoDAAS.md#GpoCopyGpoPolicySet) | **Post** /gpo/policySets/{policySetGuid} | Create a new GPO policy set by copying an existing policy set. The policies and settings in each policy are copied to the new policy set. Optionally, the filters in the policies may be copied. Regardless of the type of the source policy set, the resulting policy set is always of type DeliveryGroupPolicies.
[**GpoCreateGpoFilter**](GpoDAAS.md#GpoCreateGpoFilter) | **Post** /gpo/filters | Create a filter in a policy. Filters cannot be created in a policy in a policy set of type SiteTemplates or CustomTemplates.
[**GpoCreateGpoPolicy**](GpoDAAS.md#GpoCreateGpoPolicy) | **Post** /gpo/policies | Create a new policy. Policies cannot be created in the policy set of type SiteTemplates.
[**GpoCreateGpoPolicySet**](GpoDAAS.md#GpoCreateGpoPolicySet) | **Post** /gpo/policySets | Create a new GPO policy set. Only a policy set of type DeliveryGroupPolicies can be created.
[**GpoCreateGpoSetting**](GpoDAAS.md#GpoCreateGpoSetting) | **Post** /gpo/settings | Create a setting in a policy. Settings cannot be created in the policy set of type SiteTemplates.
[**GpoDeleteGpoFilter**](GpoDAAS.md#GpoDeleteGpoFilter) | **Delete** /gpo/filters/{filterGuid} | Delete an existing filter.
[**GpoDeleteGpoPolicy**](GpoDAAS.md#GpoDeleteGpoPolicy) | **Delete** /gpo/policies/{policyGuid} | Delete an existing GPO policy. A policy in the policy set of type SiteTemplates cannot be deleted. The Unfiltered policy in the policy set of type SitePolicies cannot be deleted.
[**GpoDeleteGpoPolicySet**](GpoDAAS.md#GpoDeleteGpoPolicySet) | **Delete** /gpo/policySets/{policySetGuid} | Delete an existing GPO policy set. Only policy sets of type DeliveryGroupPolicies can be deleted. Policies in the policy set are deleted if a policy set is deleted.
[**GpoDeleteGpoSetting**](GpoDAAS.md#GpoDeleteGpoSetting) | **Delete** /gpo/settings/{settingGuid} | Delete a setting. Settings in the policy set of type SiteTemplates cannot be deleted.
[**GpoDisableGpoPolicies**](GpoDAAS.md#GpoDisableGpoPolicies) | **Post** /gpo/policies/$disable | Disable some policies of a policy set.
[**GpoEnableGpoPolicies**](GpoDAAS.md#GpoEnableGpoPolicies) | **Post** /gpo/policies/$enable | Enable some policies of a policy set.
[**GpoGetFilterDefinitions**](GpoDAAS.md#GpoGetFilterDefinitions) | **Get** /gpo/filterDefinitions | Get all filter definitions.
[**GpoGetSettingDefinitions**](GpoDAAS.md#GpoGetSettingDefinitions) | **Get** /gpo/settingDefinitions | Get setting definitions. If isLean is set to true, only basic session information is returned. EnumType, VdaVersions, VersionDetails, and Explanation are not retrieved. If limit is set to -1 or a number larger than the number of settings available, all entries are retrieved. If limit is set to a positive integer smaller than the number of settings available, the specified number of settings are retrieved.
[**GpoGetSettingFullDetail**](GpoDAAS.md#GpoGetSettingFullDetail) | **Get** /gpo/settingFullDetail | Get full detail of a setting definition.
[**GpoMoveGpoPolicies**](GpoDAAS.md#GpoMoveGpoPolicies) | **Post** /gpo/policies/$move | Move some policies of a policy set to another policy set.
[**GpoRankGpoPolicies**](GpoDAAS.md#GpoRankGpoPolicies) | **Post** /gpo/policyPriorities | Specify new priority order for all existing policies in a policy set. All the policies in the policy set must be specified, even if the priorities of only some of the policies are changed.
[**GpoReadGpoFilter**](GpoDAAS.md#GpoReadGpoFilter) | **Get** /gpo/filters/{filterGuid} | Read a specific filter.
[**GpoReadGpoFilters**](GpoDAAS.md#GpoReadGpoFilters) | **Get** /gpo/filters | Read filters defined in a policy. A policy in a policy set of type SiteTemplates or CustomTemplates does not have filters.
[**GpoReadGpoPolicies**](GpoDAAS.md#GpoReadGpoPolicies) | **Get** /gpo/policies | Read all policies defined in a policy set. Policy templates don&#39;t have filters.
[**GpoReadGpoPolicy**](GpoDAAS.md#GpoReadGpoPolicy) | **Get** /gpo/policies/{policyGuid} | Read a policy. A policy template doesn&#39;t have filters.
[**GpoReadGpoPolicySet**](GpoDAAS.md#GpoReadGpoPolicySet) | **Get** /gpo/policySets/{policySetGuid} | Read a GPO policy set.
[**GpoReadGpoPolicySets**](GpoDAAS.md#GpoReadGpoPolicySets) | **Get** /gpo/policySets | Get all GPO policy sets in the site.
[**GpoReadGpoSetting**](GpoDAAS.md#GpoReadGpoSetting) | **Get** /gpo/settings/{settingGuid} | Read a specific setting.
[**GpoReadGpoSettings**](GpoDAAS.md#GpoReadGpoSettings) | **Get** /gpo/settings | Read settings defined in a policy.
[**GpoRemoveGpoPolicies**](GpoDAAS.md#GpoRemoveGpoPolicies) | **Post** /gpo/policies/$remove | Remove some policies of a policy set.
[**GpoRunSimulation**](GpoDAAS.md#GpoRunSimulation) | **Post** /gpo/simulation | Simulate policy application.
[**GpoRunValidation**](GpoDAAS.md#GpoRunValidation) | **Get** /gpo/validation | Check the site policies to ensure that they can be successfully converted to the new GPO objects. If an error exists in the value of a setting or filter, the error must be fixed before the policy data can be converted to new GPO setting, filter, and policy objects. The validation is done only on the site policies.
[**GpoSearchFilters**](GpoDAAS.md#GpoSearchFilters) | **Post** /gpo/filters/$search | Perform an advanced search for GPO filters.
[**GpoSearchPolicies**](GpoDAAS.md#GpoSearchPolicies) | **Post** /gpo/policies/$search | Perform an advanced search for GPO policies.
[**GpoSearchPolicySets**](GpoDAAS.md#GpoSearchPolicySets) | **Post** /gpo/policySets/$search | Perform an advanced search for GPO policy sets.
[**GpoSearchSettings**](GpoDAAS.md#GpoSearchSettings) | **Post** /gpo/settings/$search | Perform an advanced search for GPO settings.
[**GpoUpdateGpoFilter**](GpoDAAS.md#GpoUpdateGpoFilter) | **Patch** /gpo/filters/{filterGuid} | Update an existing filter.
[**GpoUpdateGpoPolicy**](GpoDAAS.md#GpoUpdateGpoPolicy) | **Patch** /gpo/policies/{policyGuid} | Update a policy. Only the policy body is updated.
[**GpoUpdateGpoPolicySet**](GpoDAAS.md#GpoUpdateGpoPolicySet) | **Patch** /gpo/policySets/{policySetGuid} | Update an existing GPO policy set.
[**GpoUpdateGpoSetting**](GpoDAAS.md#GpoUpdateGpoSetting) | **Patch** /gpo/settings/{settingGuid} | Update a setting. Settings in the policy set of type SiteTemplates cannot be updated.
[**GpoUpdatePolicySetBlob**](GpoDAAS.md#GpoUpdatePolicySetBlob) | **Put** /gpo/policySets/{policySetGuid} | Force serialization of policy set. The data of a policy set is serialized into a byte stream before it is sent to VDAs. The serialization is done automatically in the background at 5 minute intervals. A change made to the data in a policy set may not be in the serialized data for up to 5 minutes. This call tells the background thread to serialize the data immediately. No serialization is done if there have been no changes to the policy set data since the last time the data was serialized.



## GpoComparePolicies

> ComparisonResponseContract GpoComparePolicies(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).WithDefaults(withDefaults).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Compare policies. The policies can be from different policy sets.

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
	withDefaults := true // bool | Include defaults in comparison
	requestBody := []string{"Property_example"} // []string | GUIDs of policies
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoComparePolicies(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).WithDefaults(withDefaults).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoComparePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoComparePolicies`: ComparisonResponseContract
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoComparePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoComparePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **withDefaults** | **bool** | Include defaults in comparison | 
 **requestBody** | **[]string** | GUIDs of policies | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## GpoCopyGpoPolicies

> GpoCopyGpoPolicies(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).WithFilters(withFilters).CopyPoliciesRequest(copyPoliciesRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Copy some policies of a policy set to other policy sets.

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
	withFilters := true // bool | Specify if policy filters should be included
	copyPoliciesRequest := *openapiclient.NewCopyPoliciesRequest() // CopyPoliciesRequest | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoCopyGpoPolicies(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).WithFilters(withFilters).CopyPoliciesRequest(copyPoliciesRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoCopyGpoPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoCopyGpoPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **withFilters** | **bool** | Specify if policy filters should be included | 
 **copyPoliciesRequest** | [**CopyPoliciesRequest**](CopyPoliciesRequest.md) |  | 
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


## GpoCopyGpoPolicySet

> PolicySetResponse GpoCopyGpoPolicySet(ctx, policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).Name(name).WithFilters(withFilters).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create a new GPO policy set by copying an existing policy set. The policies and settings in each policy are copied to the new policy set. Optionally, the filters in the policies may be copied. Regardless of the type of the source policy set, the resulting policy set is always of type DeliveryGroupPolicies.

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
	policySetGuid := "policySetGuid_example" // string | GUID of the existing policy set
	name := "name_example" // string | The name of the new policy set
	withFilters := true // bool | Indicate if filters in the policies should be copied
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoCopyGpoPolicySet(context.Background(), policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).Name(name).WithFilters(withFilters).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoCopyGpoPolicySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoCopyGpoPolicySet`: PolicySetResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoCopyGpoPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policySetGuid** | **string** | GUID of the existing policy set | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoCopyGpoPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **name** | **string** | The name of the new policy set | 
 **withFilters** | **bool** | Indicate if filters in the policies should be copied | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## GpoCreateGpoFilter

> FilterResponse GpoCreateGpoFilter(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyGuid(policyGuid).FilterRequest(filterRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create a filter in a policy. Filters cannot be created in a policy in a policy set of type SiteTemplates or CustomTemplates.

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
	policyGuid := "policyGuid_example" // string | The GUID of the policy to which the new filter belongs
	filterRequest := *openapiclient.NewFilterRequest() // FilterRequest | Filter data
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoCreateGpoFilter(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyGuid(policyGuid).FilterRequest(filterRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoCreateGpoFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoCreateGpoFilter`: FilterResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoCreateGpoFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoCreateGpoFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policyGuid** | **string** | The GUID of the policy to which the new filter belongs | 
 **filterRequest** | [**FilterRequest**](FilterRequest.md) | Filter data | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

> PolicyResponse GpoCreateGpoPolicy(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetGuid(policySetGuid).PolicyRequest(policyRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create a new policy. Policies cannot be created in the policy set of type SiteTemplates.

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
	policySetGuid := "policySetGuid_example" // string | GUID of the policy set in which the new policy is created
	policyRequest := *openapiclient.NewPolicyRequest() // PolicyRequest | Data for the new policy
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoCreateGpoPolicy(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetGuid(policySetGuid).PolicyRequest(policyRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoCreateGpoPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoCreateGpoPolicy`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoCreateGpoPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoCreateGpoPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policySetGuid** | **string** | GUID of the policy set in which the new policy is created | 
 **policyRequest** | [**PolicyRequest**](PolicyRequest.md) | Data for the new policy | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

> PolicySetResponse GpoCreateGpoPolicySet(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetRequest(policySetRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create a new GPO policy set. Only a policy set of type DeliveryGroupPolicies can be created.

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
	policySetRequest := *openapiclient.NewPolicySetRequest() // PolicySetRequest | Parameters for the new policy set
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoCreateGpoPolicySet(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetRequest(policySetRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoCreateGpoPolicySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoCreateGpoPolicySet`: PolicySetResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoCreateGpoPolicySet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoCreateGpoPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policySetRequest** | [**PolicySetRequest**](PolicySetRequest.md) | Parameters for the new policy set | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

> SettingResponse GpoCreateGpoSetting(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyGuid(policyGuid).SettingRequest(settingRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create a setting in a policy. Settings cannot be created in the policy set of type SiteTemplates.

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
	policyGuid := "policyGuid_example" // string | GUID of the policy to which the setting belongs
	settingRequest := *openapiclient.NewSettingRequest() // SettingRequest | Data for the new setting
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoCreateGpoSetting(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyGuid(policyGuid).SettingRequest(settingRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoCreateGpoSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoCreateGpoSetting`: SettingResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoCreateGpoSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoCreateGpoSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policyGuid** | **string** | GUID of the policy to which the setting belongs | 
 **settingRequest** | [**SettingRequest**](SettingRequest.md) | Data for the new setting | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

> GpoDeleteGpoFilter(ctx, filterGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	filterGuid := "filterGuid_example" // string | The GUID of the filter to be deleted
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoDeleteGpoFilter(context.Background(), filterGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoDeleteGpoFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterGuid** | **string** | The GUID of the filter to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoDeleteGpoFilterRequest struct via the builder pattern


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


## GpoDeleteGpoPolicy

> GpoDeleteGpoPolicy(ctx, policyGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete an existing GPO policy. A policy in the policy set of type SiteTemplates cannot be deleted. The Unfiltered policy in the policy set of type SitePolicies cannot be deleted.

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
	policyGuid := "policyGuid_example" // string | GUID of the policy to be deleted
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoDeleteGpoPolicy(context.Background(), policyGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoDeleteGpoPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyGuid** | **string** | GUID of the policy to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoDeleteGpoPolicyRequest struct via the builder pattern


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


## GpoDeleteGpoPolicySet

> GpoDeleteGpoPolicySet(ctx, policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete an existing GPO policy set. Only policy sets of type DeliveryGroupPolicies can be deleted. Policies in the policy set are deleted if a policy set is deleted.

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
	policySetGuid := "policySetGuid_example" // string | The GUID of the policy set to be deleted
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoDeleteGpoPolicySet(context.Background(), policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoDeleteGpoPolicySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policySetGuid** | **string** | The GUID of the policy set to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoDeleteGpoPolicySetRequest struct via the builder pattern


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


## GpoDeleteGpoSetting

> GpoDeleteGpoSetting(ctx, settingGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete a setting. Settings in the policy set of type SiteTemplates cannot be deleted.

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
	settingGuid := "settingGuid_example" // string | GUID of the setting to be deleted
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoDeleteGpoSetting(context.Background(), settingGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoDeleteGpoSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingGuid** | **string** | GUID of the setting to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoDeleteGpoSettingRequest struct via the builder pattern


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


## GpoDisableGpoPolicies

> GpoDisableGpoPolicies(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Disable some policies of a policy set.

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
	requestBody := []string{"Property_example"} // []string | GUIDs of the policies to be disabled
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoDisableGpoPolicies(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoDisableGpoPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoDisableGpoPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **requestBody** | **[]string** | GUIDs of the policies to be disabled | 
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


## GpoEnableGpoPolicies

> GpoEnableGpoPolicies(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Enable some policies of a policy set.

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
	requestBody := []string{"Property_example"} // []string | GUIDs of the policies to be enabled
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoEnableGpoPolicies(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoEnableGpoPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoEnableGpoPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **requestBody** | **[]string** | GUIDs of the policies to be enabled | 
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


## GpoGetFilterDefinitions

> CollectionEnvelopeOfFilterDefinition GpoGetFilterDefinitions(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get all filter definitions.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoGetFilterDefinitions(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoGetFilterDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoGetFilterDefinitions`: CollectionEnvelopeOfFilterDefinition
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoGetFilterDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoGetFilterDefinitionsRequest struct via the builder pattern


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

> SettingDefinitionEnvelope GpoGetSettingDefinitions(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).IsLean(isLean).Limit(limit).IsAscending(isAscending).NamePattern(namePattern).IsUserSetting(isUserSetting).ContinuationToken(continuationToken).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	isLean := true // bool | Get lean parts of setting definitions, the default is set to true (optional) (default to true)
	limit := int32(56) // int32 | Specify the number of entries to retrieve, the default is all entries (optional) (default to -1)
	isAscending := true // bool | Specify sort order, default is true (optional) (default to true)
	namePattern := "namePattern_example" // string | Specify a regular expression to match the internal setting name. The default is match all names.              (optional)
	isUserSetting := true // bool | Specify the target of applying the settings. If it's set to true, only user settings are retrieved. If it's set to false, only computer settings are retrieved. If not specified, both kinds of settings are retrieved. The default is to retrieve both kinds of settings.              (optional)
	continuationToken := "continuationToken_example" // string | Continuation token from a previous retrieval (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoGetSettingDefinitions(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).IsLean(isLean).Limit(limit).IsAscending(isAscending).NamePattern(namePattern).IsUserSetting(isUserSetting).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoGetSettingDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoGetSettingDefinitions`: SettingDefinitionEnvelope
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoGetSettingDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoGetSettingDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **isLean** | **bool** | Get lean parts of setting definitions, the default is set to true | [default to true]
 **limit** | **int32** | Specify the number of entries to retrieve, the default is all entries | [default to -1]
 **isAscending** | **bool** | Specify sort order, default is true | [default to true]
 **namePattern** | **string** | Specify a regular expression to match the internal setting name. The default is match all names.              | 
 **isUserSetting** | **bool** | Specify the target of applying the settings. If it&#39;s set to true, only user settings are retrieved. If it&#39;s set to false, only computer settings are retrieved. If not specified, both kinds of settings are retrieved. The default is to retrieve both kinds of settings.              | 
 **continuationToken** | **string** | Continuation token from a previous retrieval | 

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

> SettingDefinition GpoGetSettingFullDetail(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).SettingName(settingName).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get full detail of a setting definition.

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
	settingName := "settingName_example" // string | The internal name of the setting
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoGetSettingFullDetail(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).SettingName(settingName).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoGetSettingFullDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoGetSettingFullDetail`: SettingDefinition
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoGetSettingFullDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoGetSettingFullDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **settingName** | **string** | The internal name of the setting | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## GpoMoveGpoPolicies

> GpoMoveGpoPolicies(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).ToPolicySet(toPolicySet).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Move some policies of a policy set to another policy set.

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
	toPolicySet := "toPolicySet_example" // string | GUID of the destination policy set
	requestBody := []string{"Property_example"} // []string | GUIDs of the policies to be removed
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoMoveGpoPolicies(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).ToPolicySet(toPolicySet).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoMoveGpoPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoMoveGpoPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **toPolicySet** | **string** | GUID of the destination policy set | 
 **requestBody** | **[]string** | GUIDs of the policies to be removed | 
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


## GpoRankGpoPolicies

> bool GpoRankGpoPolicies(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetGuid(policySetGuid).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Specify new priority order for all existing policies in a policy set. All the policies in the policy set must be specified, even if the priorities of only some of the policies are changed.

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
	policySetGuid := "policySetGuid_example" // string | Guid of the policy set that contains the policies
	requestBody := []string{"Property_example"} // []string | GUIDs of the policies to be re-prioritized
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoRankGpoPolicies(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetGuid(policySetGuid).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoRankGpoPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoRankGpoPolicies`: bool
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoRankGpoPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoRankGpoPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policySetGuid** | **string** | Guid of the policy set that contains the policies | 
 **requestBody** | **[]string** | GUIDs of the policies to be re-prioritized | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

> FilterResponse GpoReadGpoFilter(ctx, filterGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	filterGuid := "filterGuid_example" // string | The GUID of the filter to be read
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoReadGpoFilter(context.Background(), filterGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoReadGpoFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoReadGpoFilter`: FilterResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoReadGpoFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterGuid** | **string** | The GUID of the filter to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoFilterRequest struct via the builder pattern


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

> CollectionEnvelopeOfFilterResponse GpoReadGpoFilters(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyGuid(policyGuid).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Read filters defined in a policy. A policy in a policy set of type SiteTemplates or CustomTemplates does not have filters.

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
	policyGuid := "policyGuid_example" // string | The GUID of the policy from which filters are read
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoReadGpoFilters(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyGuid(policyGuid).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoReadGpoFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoReadGpoFilters`: CollectionEnvelopeOfFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoReadGpoFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policyGuid** | **string** | The GUID of the policy from which filters are read | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

> CollectionEnvelopeOfPolicyResponse GpoReadGpoPolicies(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetGuid(policySetGuid).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).WithSettings(withSettings).WithFilters(withFilters).Execute()

Read all policies defined in a policy set. Policy templates don't have filters.

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
	policySetGuid := "policySetGuid_example" // string | The GUID of the policy set from which policies are read
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	withSettings := true // bool | If set to true, settings in the policy are read (optional)
	withFilters := true // bool | If set to true, filters in the policy are read (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoReadGpoPolicies(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetGuid(policySetGuid).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).WithSettings(withSettings).WithFilters(withFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoReadGpoPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoReadGpoPolicies`: CollectionEnvelopeOfPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoReadGpoPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policySetGuid** | **string** | The GUID of the policy set from which policies are read | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **withSettings** | **bool** | If set to true, settings in the policy are read | 
 **withFilters** | **bool** | If set to true, filters in the policy are read | 

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

> PolicyResponse GpoReadGpoPolicy(ctx, policyGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).WithSettings(withSettings).WithFilters(withFilters).Execute()

Read a policy. A policy template doesn't have filters.

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
	policyGuid := "policyGuid_example" // string | GUID of the policy to be read
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	withSettings := true // bool | If set to true, read policy settings (optional)
	withFilters := true // bool | If set to true, read policy filters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoReadGpoPolicy(context.Background(), policyGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).WithSettings(withSettings).WithFilters(withFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoReadGpoPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoReadGpoPolicy`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoReadGpoPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyGuid** | **string** | GUID of the policy to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **withSettings** | **bool** | If set to true, read policy settings | 
 **withFilters** | **bool** | If set to true, read policy filters | 

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

> PolicySetResponse GpoReadGpoPolicySet(ctx, policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).WithPolicies(withPolicies).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	policySetGuid := "policySetGuid_example" // string | GUID of the policy set to read
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	withPolicies := true // bool | If set to true, read the policies in the policy set (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoReadGpoPolicySet(context.Background(), policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).WithPolicies(withPolicies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoReadGpoPolicySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoReadGpoPolicySet`: PolicySetResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoReadGpoPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policySetGuid** | **string** | GUID of the policy set to read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **withPolicies** | **bool** | If set to true, read the policies in the policy set | 

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

> CollectionEnvelopeOfPolicySetResponse GpoReadGpoPolicySets(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoReadGpoPolicySets(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoReadGpoPolicySets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoReadGpoPolicySets`: CollectionEnvelopeOfPolicySetResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoReadGpoPolicySets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoPolicySetsRequest struct via the builder pattern


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

> SettingResponse GpoReadGpoSetting(ctx, settingGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	settingGuid := "settingGuid_example" // string | GUID of the setting to be read
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoReadGpoSetting(context.Background(), settingGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoReadGpoSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoReadGpoSetting`: SettingResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoReadGpoSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingGuid** | **string** | GUID of the setting to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoSettingRequest struct via the builder pattern


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

> CollectionEnvelopeOfSettingResponse GpoReadGpoSettings(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyGuid(policyGuid).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	policyGuid := "policyGuid_example" // string | GUID of the policy from which settings are read
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoReadGpoSettings(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyGuid(policyGuid).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoReadGpoSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoReadGpoSettings`: CollectionEnvelopeOfSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoReadGpoSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoReadGpoSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policyGuid** | **string** | GUID of the policy from which settings are read | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## GpoRemoveGpoPolicies

> GpoRemoveGpoPolicies(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Remove some policies of a policy set.

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
	requestBody := []string{"Property_example"} // []string | GUIDs of the policies to be removed
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoRemoveGpoPolicies(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).RequestBody(requestBody).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoRemoveGpoPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoRemoveGpoPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **requestBody** | **[]string** | GUIDs of the policies to be removed | 
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


## GpoRunSimulation

> []SimulationResponseContract GpoRunSimulation(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetGuid(policySetGuid).SimulationRequestContract(simulationRequestContract).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Simulate policy application.

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
	policySetGuid := "policySetGuid_example" // string | The GUID of the policy set, on which this simulation is run
	simulationRequestContract := *openapiclient.NewSimulationRequestContract() // SimulationRequestContract | Modeling request
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoRunSimulation(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetGuid(policySetGuid).SimulationRequestContract(simulationRequestContract).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoRunSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoRunSimulation`: []SimulationResponseContract
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoRunSimulation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoRunSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policySetGuid** | **string** | The GUID of the policy set, on which this simulation is run | 
 **simulationRequestContract** | [**SimulationRequestContract**](SimulationRequestContract.md) | Modeling request | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**[]SimulationResponseContract**](SimulationResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoRunValidation

> []GpoTestPolicyData GpoRunValidation(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).All(all).Execute()

Check the site policies to ensure that they can be successfully converted to the new GPO objects. If an error exists in the value of a setting or filter, the error must be fixed before the policy data can be converted to new GPO setting, filter, and policy objects. The validation is done only on the site policies.

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
	all := true // bool | If true, all settings and filters in the policies are returned, otherwise only settings and filters with errors are returned. Policies are always included in the result, regardless of the value of this parameter. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoRunValidation(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoRunValidation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoRunValidation`: []GpoTestPolicyData
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoRunValidation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoRunValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **all** | **bool** | If true, all settings and filters in the policies are returned, otherwise only settings and filters with errors are returned. Policies are always included in the result, regardless of the value of this parameter. | [default to false]

### Return type

[**[]GpoTestPolicyData**](GpoTestPolicyData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoSearchFilters

> CollectionEnvelopeOfFilterResponse GpoSearchFilters(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).FilterSearch(filterSearch).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()

Perform an advanced search for GPO filters.

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
	filterSearch := *openapiclient.NewFilterSearch() // FilterSearch | Specifies the advanced search parameters.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | Maximum number of objects, which must not exceed 1000, to return. If not specified, the default value is 250. (optional)
	continuationToken := "continuationToken_example" // string | If a previous call didn't return all the records expected, pass the continuation token returned from the previous response to this current query to continue retrieving the remaining objects. (optional)
	async := true // bool | If `true`, the search operations will be executed as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoSearchFilters(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).FilterSearch(filterSearch).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoSearchFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoSearchFilters`: CollectionEnvelopeOfFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoSearchFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoSearchFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **filterSearch** | [**FilterSearch**](FilterSearch.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | Maximum number of objects, which must not exceed 1000, to return. If not specified, the default value is 250. | 
 **continuationToken** | **string** | If a previous call didn&#39;t return all the records expected, pass the continuation token returned from the previous response to this current query to continue retrieving the remaining objects. | 
 **async** | **bool** | If &#x60;true&#x60;, the search operations will be executed as a background task. | [default to false]

### Return type

[**CollectionEnvelopeOfFilterResponse**](CollectionEnvelopeOfFilterResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoSearchPolicies

> CollectionEnvelopeOfPolicyResponse GpoSearchPolicies(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySearch(policySearch).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()

Perform an advanced search for GPO policies.

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
	policySearch := *openapiclient.NewPolicySearch() // PolicySearch | Specifies the advanced search parameters.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | Maximum number of objects, which must not exceed 1000, to return. If not specified, the default value is 250. (optional)
	continuationToken := "continuationToken_example" // string | If a previous call didn't return all the records expected, pass the continuation token returned from the previous response to this current query to continue retrieving the remaining objects. (optional)
	async := true // bool | If `true`, the search operations will be executed as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoSearchPolicies(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySearch(policySearch).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoSearchPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoSearchPolicies`: CollectionEnvelopeOfPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoSearchPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoSearchPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policySearch** | [**PolicySearch**](PolicySearch.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | Maximum number of objects, which must not exceed 1000, to return. If not specified, the default value is 250. | 
 **continuationToken** | **string** | If a previous call didn&#39;t return all the records expected, pass the continuation token returned from the previous response to this current query to continue retrieving the remaining objects. | 
 **async** | **bool** | If &#x60;true&#x60;, the search operations will be executed as a background task. | [default to false]

### Return type

[**CollectionEnvelopeOfPolicyResponse**](CollectionEnvelopeOfPolicyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoSearchPolicySets

> CollectionEnvelopeOfPolicySetResponse GpoSearchPolicySets(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetSearch(policySetSearch).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()

Perform an advanced search for GPO policy sets.

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
	policySetSearch := *openapiclient.NewPolicySetSearch() // PolicySetSearch | Specifies the advanced search parameters.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | Maximum number of objects, which must not exceed 1000, to return. If not specified, the default value is 250. (optional)
	continuationToken := "continuationToken_example" // string | If a previous call didn't return all the records expected, pass the continuation token returned from the previous response to this current query to continue retrieving the remaining objects. (optional)
	async := true // bool | If `true`, the search operations will be executed as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoSearchPolicySets(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetSearch(policySetSearch).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoSearchPolicySets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoSearchPolicySets`: CollectionEnvelopeOfPolicySetResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoSearchPolicySets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoSearchPolicySetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **policySetSearch** | [**PolicySetSearch**](PolicySetSearch.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | Maximum number of objects, which must not exceed 1000, to return. If not specified, the default value is 250. | 
 **continuationToken** | **string** | If a previous call didn&#39;t return all the records expected, pass the continuation token returned from the previous response to this current query to continue retrieving the remaining objects. | 
 **async** | **bool** | If &#x60;true&#x60;, the search operations will be executed as a background task. | [default to false]

### Return type

[**CollectionEnvelopeOfPolicySetResponse**](CollectionEnvelopeOfPolicySetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoSearchSettings

> CollectionEnvelopeOfSettingResponse GpoSearchSettings(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).SettingSearch(settingSearch).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()

Perform an advanced search for GPO settings.

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
	settingSearch := *openapiclient.NewSettingSearch() // SettingSearch | Specifies the advanced search parameters.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | Maximum number of objects, which must not exceed 1000, to return. If not specified, the default value is 250. (optional)
	continuationToken := "continuationToken_example" // string | If a previous call didn't return all the records expected, pass the continuation token returned from the previous response to this current query to continue retrieving the remaining objects. (optional)
	async := true // bool | If `true`, the search operations will be executed as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GpoDAAS.GpoSearchSettings(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).SettingSearch(settingSearch).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoSearchSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpoSearchSettings`: CollectionEnvelopeOfSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `GpoDAAS.GpoSearchSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGpoSearchSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **settingSearch** | [**SettingSearch**](SettingSearch.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | Maximum number of objects, which must not exceed 1000, to return. If not specified, the default value is 250. | 
 **continuationToken** | **string** | If a previous call didn&#39;t return all the records expected, pass the continuation token returned from the previous response to this current query to continue retrieving the remaining objects. | 
 **async** | **bool** | If &#x60;true&#x60;, the search operations will be executed as a background task. | [default to false]

### Return type

[**CollectionEnvelopeOfSettingResponse**](CollectionEnvelopeOfSettingResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GpoUpdateGpoFilter

> GpoUpdateGpoFilter(ctx, filterGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).FilterRequest(filterRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	filterGuid := "filterGuid_example" // string | The GUID of the filter to be updated
	filterRequest := *openapiclient.NewFilterRequest() // FilterRequest | Filter data. The filter type is ignored
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoUpdateGpoFilter(context.Background(), filterGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).FilterRequest(filterRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoUpdateGpoFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterGuid** | **string** | The GUID of the filter to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdateGpoFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **filterRequest** | [**FilterRequest**](FilterRequest.md) | Filter data. The filter type is ignored | 
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


## GpoUpdateGpoPolicy

> GpoUpdateGpoPolicy(ctx, policyGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyBodyRequest(policyBodyRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	policyGuid := "policyGuid_example" // string | GUID of the policy to be updated
	policyBodyRequest := *openapiclient.NewPolicyBodyRequest() // PolicyBodyRequest | New policy data
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoUpdateGpoPolicy(context.Background(), policyGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicyBodyRequest(policyBodyRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoUpdateGpoPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyGuid** | **string** | GUID of the policy to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdateGpoPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **policyBodyRequest** | [**PolicyBodyRequest**](PolicyBodyRequest.md) | New policy data | 
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


## GpoUpdateGpoPolicySet

> GpoUpdateGpoPolicySet(ctx, policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetRequest(policySetRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	policySetGuid := "policySetGuid_example" // string | The GUID of the policy set to update
	policySetRequest := *openapiclient.NewPolicySetRequest() // PolicySetRequest | New description of the policy set
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoUpdateGpoPolicySet(context.Background(), policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).PolicySetRequest(policySetRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoUpdateGpoPolicySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policySetGuid** | **string** | The GUID of the policy set to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdateGpoPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **policySetRequest** | [**PolicySetRequest**](PolicySetRequest.md) | New description of the policy set | 
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


## GpoUpdateGpoSetting

> GpoUpdateGpoSetting(ctx, settingGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).SettingRequest(settingRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update a setting. Settings in the policy set of type SiteTemplates cannot be updated.

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
	settingGuid := "settingGuid_example" // string | GUID of the setting to be updated
	settingRequest := *openapiclient.NewSettingRequest() // SettingRequest | New setting value
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoUpdateGpoSetting(context.Background(), settingGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).SettingRequest(settingRequest).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoUpdateGpoSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingGuid** | **string** | GUID of the setting to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdateGpoSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **settingRequest** | [**SettingRequest**](SettingRequest.md) | New setting value | 
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


## GpoUpdatePolicySetBlob

> GpoUpdatePolicySetBlob(ctx, policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Force serialization of policy set. The data of a policy set is serialized into a byte stream before it is sent to VDAs. The serialization is done automatically in the background at 5 minute intervals. A change made to the data in a policy set may not be in the serialized data for up to 5 minutes. This call tells the background thread to serialize the data immediately. No serialization is done if there have been no changes to the policy set data since the last time the data was serialized.

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
	policySetGuid := "policySetGuid_example" // string | The GUID of the policy set whose data is to be serialized
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GpoDAAS.GpoUpdatePolicySetBlob(context.Background(), policySetGuid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpoDAAS.GpoUpdatePolicySetBlob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policySetGuid** | **string** | The GUID of the policy set whose data is to be serialized | 

### Other Parameters

Other parameters are passed through a pointer to a apiGpoUpdatePolicySetBlobRequest struct via the builder pattern


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

