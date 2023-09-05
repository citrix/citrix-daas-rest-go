# \GroupPolicyApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupPolicyCreateFilter**](GroupPolicyApi.md#GroupPolicyCreateFilter) | **Post** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/filters | Create a new policy filter
[**GroupPolicyCreatePolicy**](GroupPolicyApi.md#GroupPolicyCreatePolicy) | **Post** /techpreview/{customerid}/{siteid}/groupPolicy/policies | Create a policy and initialize it with the data provided.
[**GroupPolicyCreateTemplate**](GroupPolicyApi.md#GroupPolicyCreateTemplate) | **Post** /techpreview/{customerid}/{siteid}/groupPolicy/templates | Create a template
[**GroupPolicyDeleteFilter**](GroupPolicyApi.md#GroupPolicyDeleteFilter) | **Delete** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/filters/{filterName} | Delete a filter
[**GroupPolicyDeletePolicy**](GroupPolicyApi.md#GroupPolicyDeletePolicy) | **Delete** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName} | Delete a policy.
[**GroupPolicyDeletePolicySetting**](GroupPolicyApi.md#GroupPolicyDeletePolicySetting) | **Delete** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/settings/{settingName} | Unselect a setting from a policy
[**GroupPolicyDeleteTemplate**](GroupPolicyApi.md#GroupPolicyDeleteTemplate) | **Delete** /techpreview/{customerid}/{siteid}/groupPolicy/templates/{templateName} | Delete a template
[**GroupPolicyDeleteTemplateSetting**](GroupPolicyApi.md#GroupPolicyDeleteTemplateSetting) | **Delete** /techpreview/{customerid}/{siteid}/groupPolicy/templates/{templateName}/settings/{settingName} | Unselect a setting from a template
[**GroupPolicyGetComparison**](GroupPolicyApi.md#GroupPolicyGetComparison) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/compare | Compare policies or templates. Templates are currently not supported.
[**GroupPolicyGetFilterDefinitions**](GroupPolicyApi.md#GroupPolicyGetFilterDefinitions) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/filters | Get all filter definitions.
[**GroupPolicyGetModeling**](GroupPolicyApi.md#GroupPolicyGetModeling) | **Post** /techpreview/{customerid}/{siteid}/groupPolicy/modeling | GET method for Policy Modeling.
[**GroupPolicyGetSettingDefinitions**](GroupPolicyApi.md#GroupPolicyGetSettingDefinitions) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/settings | Retrieve all setting definitions.
[**GroupPolicyPrioritizePolicies**](GroupPolicyApi.md#GroupPolicyPrioritizePolicies) | **Post** /techpreview/{customerid}/{siteid}/groupPolicy/policyPriorities | Specify new priority order for all existing policies.
[**GroupPolicyReplacePolicy**](GroupPolicyApi.md#GroupPolicyReplacePolicy) | **Put** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName} | Replace an existing policy. The entire policy is replaced, including the name. Priority is ignored and will be set to the priority of the existing policy.
[**GroupPolicyRetrieveAllData**](GroupPolicyApi.md#GroupPolicyRetrieveAllData) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/alldata | Get all policies, templates, setting definitions and filter definitions
[**GroupPolicyRetrieveFilter**](GroupPolicyApi.md#GroupPolicyRetrieveFilter) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/filters/{filterName} | Retrieve a filter
[**GroupPolicyRetrieveFilters**](GroupPolicyApi.md#GroupPolicyRetrieveFilters) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/filters | Get filters of a policy
[**GroupPolicyRetrievePolicies**](GroupPolicyApi.md#GroupPolicyRetrievePolicies) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/policies | Erase the local cache and retrieve all policies from the database. Regardless if the load succeeds or fails, the local cache is always cleared.
[**GroupPolicyRetrievePoliciesTemplates**](GroupPolicyApi.md#GroupPolicyRetrievePoliciesTemplates) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/policiestemplates | Get all policies and templates
[**GroupPolicyRetrievePolicy**](GroupPolicyApi.md#GroupPolicyRetrievePolicy) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName} | Retrieve an existing policy.
[**GroupPolicyRetrievePolicySetting**](GroupPolicyApi.md#GroupPolicyRetrievePolicySetting) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/settings/{settingName} | Retrieve an existing selected policy setting
[**GroupPolicyRetrievePolicySettings**](GroupPolicyApi.md#GroupPolicyRetrievePolicySettings) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/settings | Get settings of a policy
[**GroupPolicyRetrieveTemplate**](GroupPolicyApi.md#GroupPolicyRetrieveTemplate) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/templates/{templateName} | Get one specified template
[**GroupPolicyRetrieveTemplateSetting**](GroupPolicyApi.md#GroupPolicyRetrieveTemplateSetting) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/templates/{templateName}/settings/{settingName} | Retrieve an existing selected template setting
[**GroupPolicyRetrieveTemplateSettings**](GroupPolicyApi.md#GroupPolicyRetrieveTemplateSettings) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/templates/{templateName}/settings | Get settings of a template
[**GroupPolicyRetrieveTemplates**](GroupPolicyApi.md#GroupPolicyRetrieveTemplates) | **Get** /techpreview/{customerid}/{siteid}/groupPolicy/templates | Get all templates
[**GroupPolicySelectPolicySetting**](GroupPolicyApi.md#GroupPolicySelectPolicySetting) | **Post** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/settings | Select a setting and configure it for the specified policy.
[**GroupPolicySelectTemplateSetting**](GroupPolicyApi.md#GroupPolicySelectTemplateSetting) | **Post** /techpreview/{customerid}/{siteid}/groupPolicy/templates/{templateName}/settings | Select a setting and configure it for the specified template.
[**GroupPolicyUpdateFilter**](GroupPolicyApi.md#GroupPolicyUpdateFilter) | **Post** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/filters/{filterName} | Update a filter
[**GroupPolicyUpdatePolicy**](GroupPolicyApi.md#GroupPolicyUpdatePolicy) | **Patch** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName} | Update an existing policy. If a property is null, the property is not updated. If a collection is not null, the entire collection is replaced. This API can be used to rename a policy and raise or lower a policy priority.
[**GroupPolicyUpdatePolicySetting**](GroupPolicyApi.md#GroupPolicyUpdatePolicySetting) | **Patch** /techpreview/{customerid}/{siteid}/groupPolicy/policies/{policyName}/settings/{settingName} | Update a selected policy setting.
[**GroupPolicyUpdateTemplate**](GroupPolicyApi.md#GroupPolicyUpdateTemplate) | **Post** /techpreview/{customerid}/{siteid}/groupPolicy/templates/{templateName} | Update a template
[**GroupPolicyUpdateTemplateSetting**](GroupPolicyApi.md#GroupPolicyUpdateTemplateSetting) | **Patch** /techpreview/{customerid}/{siteid}/groupPolicy/templates/{templateName}/settings/{settingName} | Update a selected template setting.



## GroupPolicyCreateFilter

> FilterResponseContract GroupPolicyCreateFilter(ctx, customerid, siteid, policyName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a new policy filter

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The policy name
    request := *openapiclient.NewFilterRequestContract() // FilterRequestContract | New filter data
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyCreateFilter(context.Background(), customerid, siteid, policyName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyCreateFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyCreateFilter`: FilterResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyCreateFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The policy name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyCreateFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**FilterRequestContract**](FilterRequestContract.md) | New filter data | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**FilterResponseContract**](FilterResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyCreatePolicy

> PolicyResponseContract GroupPolicyCreatePolicy(ctx, customerid, siteid).Locale(locale).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a policy and initialize it with the data provided.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string | Locale for display strings
    request := *openapiclient.NewPolicyRequestContract() // PolicyRequestContract | The policy model
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyCreatePolicy(context.Background(), customerid, siteid).Locale(locale).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyCreatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyCreatePolicy`: PolicyResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyCreatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | Locale for display strings | 
 **request** | [**PolicyRequestContract**](PolicyRequestContract.md) | The policy model | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PolicyResponseContract**](PolicyResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyCreateTemplate

> TemplateResponseContract GroupPolicyCreateTemplate(ctx, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a template

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    request := *openapiclient.NewTemplateRequestContract() // TemplateRequestContract | The template model
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyCreateTemplate(context.Background(), customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyCreateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyCreateTemplate`: TemplateResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyCreateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**TemplateRequestContract**](TemplateRequestContract.md) | The template model | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**TemplateResponseContract**](TemplateResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyDeleteFilter

> bool GroupPolicyDeleteFilter(ctx, customerid, siteid, policyName, filterName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a filter

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The policy name
    filterName := "filterName_example" // string | The filter name
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyDeleteFilter(context.Background(), customerid, siteid, policyName, filterName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyDeleteFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyDeleteFilter`: bool
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyDeleteFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The policy name | 
**filterName** | **string** | The filter name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyDeleteFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyDeletePolicy

> bool GroupPolicyDeletePolicy(ctx, customerid, siteid, policyName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a policy.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The policy to be deleted
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyDeletePolicy(context.Background(), customerid, siteid, policyName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyDeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyDeletePolicy`: bool
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyDeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The policy to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyDeletePolicySetting

> bool GroupPolicyDeletePolicySetting(ctx, customerid, siteid, policyName, settingName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Unselect a setting from a policy

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The policy name
    settingName := "settingName_example" // string | The setting name
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyDeletePolicySetting(context.Background(), customerid, siteid, policyName, settingName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyDeletePolicySetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyDeletePolicySetting`: bool
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyDeletePolicySetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The policy name | 
**settingName** | **string** | The setting name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyDeletePolicySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyDeleteTemplate

> bool GroupPolicyDeleteTemplate(ctx, customerid, siteid, templateName).ForceDelete(forceDelete).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a template

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    templateName := "templateName_example" // string | The template name
    forceDelete := true // bool | True to force this delete even if the database has been changed by another user
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyDeleteTemplate(context.Background(), customerid, siteid, templateName).ForceDelete(forceDelete).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyDeleteTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyDeleteTemplate`: bool
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyDeleteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**templateName** | **string** | The template name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forceDelete** | **bool** | True to force this delete even if the database has been changed by another user | 
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


## GroupPolicyDeleteTemplateSetting

> bool GroupPolicyDeleteTemplateSetting(ctx, customerid, siteid, templateName, settingName).ForceWrite(forceWrite).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Unselect a setting from a template

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    templateName := "templateName_example" // string | The template name
    settingName := "settingName_example" // string | The setting name
    forceWrite := true // bool | Force write to database
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyDeleteTemplateSetting(context.Background(), customerid, siteid, templateName, settingName).ForceWrite(forceWrite).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyDeleteTemplateSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyDeleteTemplateSetting`: bool
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyDeleteTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**templateName** | **string** | The template name | 
**settingName** | **string** | The setting name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyDeleteTemplateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **forceWrite** | **bool** | Force write to database | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyGetComparison

> ComparisonResponseContract GroupPolicyGetComparison(ctx, customerid, siteid).Locale(locale).Targets(targets).WithDefaults(withDefaults).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Compare policies or templates. Templates are currently not supported.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string | Locale for returned strings
    targets := "targets_example" // string | Comma-separated list of policies, if null, compare all
    withDefaults := true // bool | Include defaults in comparison
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyGetComparison(context.Background(), customerid, siteid).Locale(locale).Targets(targets).WithDefaults(withDefaults).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyGetComparison``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyGetComparison`: ComparisonResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyGetComparison`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyGetComparisonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | Locale for returned strings | 
 **targets** | **string** | Comma-separated list of policies, if null, compare all | 
 **withDefaults** | **bool** | Include defaults in comparison | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ComparisonResponseContract**](ComparisonResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyGetFilterDefinitions

> []FilterDefinitionContract GroupPolicyGetFilterDefinitions(ctx, customerid, siteid).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string |  (optional) (default to "en-US")
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyGetFilterDefinitions(context.Background(), customerid, siteid).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyGetFilterDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyGetFilterDefinitions`: []FilterDefinitionContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyGetFilterDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyGetFilterDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** |  | [default to &quot;en-US&quot;]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**[]FilterDefinitionContract**](FilterDefinitionContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyGetModeling

> []ModelingResponseContract GroupPolicyGetModeling(ctx, customerid, siteid).Request(request).PolicySetGuid(policySetGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

GET method for Policy Modeling.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    request := *openapiclient.NewModelingRequestContract() // ModelingRequestContract | Modeling request
    policySetGuid := "policySetGuid_example" // string | Model the specified policy set (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyGetModeling(context.Background(), customerid, siteid).Request(request).PolicySetGuid(policySetGuid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyGetModeling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyGetModeling`: []ModelingResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyGetModeling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyGetModelingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**ModelingRequestContract**](ModelingRequestContract.md) | Modeling request | 
 **policySetGuid** | **string** | Model the specified policy set | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**[]ModelingResponseContract**](ModelingResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyGetSettingDefinitions

> []SettingDefinitionContract GroupPolicyGetSettingDefinitions(ctx, customerid, siteid).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Retrieve all setting definitions.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string | The locale for the returned strings
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyGetSettingDefinitions(context.Background(), customerid, siteid).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyGetSettingDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyGetSettingDefinitions`: []SettingDefinitionContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyGetSettingDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyGetSettingDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | The locale for the returned strings | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**[]SettingDefinitionContract**](SettingDefinitionContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyPrioritizePolicies

> bool GroupPolicyPrioritizePolicies(ctx, customerid, siteid).PolicyNames(policyNames).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Specify new priority order for all existing policies.

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    policyNames := []string{"Property_example"} // []string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyPrioritizePolicies(context.Background(), customerid, siteid).PolicyNames(policyNames).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyPrioritizePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyPrioritizePolicies`: bool
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyPrioritizePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyPrioritizePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyNames** | **[]string** |  | 
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


## GroupPolicyReplacePolicy

> PolicyResponseContract GroupPolicyReplacePolicy(ctx, customerid, siteid, policyName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Replace an existing policy. The entire policy is replaced, including the name. Priority is ignored and will be set to the priority of the existing policy.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The name of an existing policy
    request := *openapiclient.NewPolicyRequestContract() // PolicyRequestContract | Input data
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyReplacePolicy(context.Background(), customerid, siteid, policyName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyReplacePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyReplacePolicy`: PolicyResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyReplacePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The name of an existing policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyReplacePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**PolicyRequestContract**](PolicyRequestContract.md) | Input data | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PolicyResponseContract**](PolicyResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrieveAllData

> AllResponseContract GroupPolicyRetrieveAllData(ctx, customerid, siteid).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all policies, templates, setting definitions and filter definitions

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string | Target locale
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrieveAllData(context.Background(), customerid, siteid).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrieveAllData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrieveAllData`: AllResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrieveAllData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrieveAllDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | Target locale | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AllResponseContract**](AllResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrieveFilter

> FilterResponseContract GroupPolicyRetrieveFilter(ctx, customerid, siteid, policyName, filterName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Retrieve a filter

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The policy name
    filterName := "filterName_example" // string | The filter name
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrieveFilter(context.Background(), customerid, siteid, policyName, filterName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrieveFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrieveFilter`: FilterResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrieveFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The policy name | 
**filterName** | **string** | The filter name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrieveFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**FilterResponseContract**](FilterResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrieveFilters

> []FilterResponseContract GroupPolicyRetrieveFilters(ctx, customerid, siteid, policyName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get filters of a policy

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The policy name
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrieveFilters(context.Background(), customerid, siteid, policyName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrieveFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrieveFilters`: []FilterResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrieveFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The policy name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrieveFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**[]FilterResponseContract**](FilterResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrievePolicies

> []PolicyResponseContract GroupPolicyRetrievePolicies(ctx, customerid, siteid).Locale(locale).WithSettings(withSettings).WithFilters(withFilters).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Erase the local cache and retrieve all policies from the database. Regardless if the load succeeds or fails, the local cache is always cleared.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string | The locale for the returned strings
    withSettings := true // bool | True if settings are included in the result
    withFilters := true // bool | True if filters are included in the result
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrievePolicies(context.Background(), customerid, siteid).Locale(locale).WithSettings(withSettings).WithFilters(withFilters).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrievePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrievePolicies`: []PolicyResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrievePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrievePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | The locale for the returned strings | 
 **withSettings** | **bool** | True if settings are included in the result | 
 **withFilters** | **bool** | True if filters are included in the result | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**[]PolicyResponseContract**](PolicyResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrievePoliciesTemplates

> PolicyTemplateResponseContract GroupPolicyRetrievePoliciesTemplates(ctx, customerid, siteid).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all policies and templates

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string | The locale
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrievePoliciesTemplates(context.Background(), customerid, siteid).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrievePoliciesTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrievePoliciesTemplates`: PolicyTemplateResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrievePoliciesTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrievePoliciesTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | The locale | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PolicyTemplateResponseContract**](PolicyTemplateResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrievePolicy

> PolicyResponseContract GroupPolicyRetrievePolicy(ctx, customerid, siteid, policyName).Locale(locale).WithSettings(withSettings).WithFilters(withFilters).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Retrieve an existing policy.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The policy name
    locale := "locale_example" // string | Locale for display strings
    withSettings := true // bool | True if settings are included in the result
    withFilters := true // bool | True if filters are included in the result
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrievePolicy(context.Background(), customerid, siteid, policyName).Locale(locale).WithSettings(withSettings).WithFilters(withFilters).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrievePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrievePolicy`: PolicyResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrievePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The policy name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrievePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **locale** | **string** | Locale for display strings | 
 **withSettings** | **bool** | True if settings are included in the result | 
 **withFilters** | **bool** | True if filters are included in the result | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PolicyResponseContract**](PolicyResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrievePolicySetting

> SettingResponseContract GroupPolicyRetrievePolicySetting(ctx, customerid, siteid, policyName, settingName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Retrieve an existing selected policy setting

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | Name of the policy
    settingName := "settingName_example" // string | The setting name
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrievePolicySetting(context.Background(), customerid, siteid, policyName, settingName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrievePolicySetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrievePolicySetting`: SettingResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrievePolicySetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | Name of the policy | 
**settingName** | **string** | The setting name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrievePolicySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingResponseContract**](SettingResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrievePolicySettings

> []SettingResponseContract GroupPolicyRetrievePolicySettings(ctx, customerid, siteid, policyName).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get settings of a policy

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string | The locale for the returned strings
    policyName := "policyName_example" // string | Name of policy for which settings are retrieved
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrievePolicySettings(context.Background(), customerid, siteid, policyName).Locale(locale).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrievePolicySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrievePolicySettings`: []SettingResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrievePolicySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | Name of policy for which settings are retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrievePolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | The locale for the returned strings | 

 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**[]SettingResponseContract**](SettingResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrieveTemplate

> TemplateResponseContract GroupPolicyRetrieveTemplate(ctx, customerid, siteid, templateName).Locale(locale).WithSettings(withSettings).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get one specified template

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    templateName := "templateName_example" // string | The template name
    locale := "locale_example" // string | The locale for strings
    withSettings := true // bool | Retrieve template settings
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrieveTemplate(context.Background(), customerid, siteid, templateName).Locale(locale).WithSettings(withSettings).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrieveTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrieveTemplate`: TemplateResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrieveTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**templateName** | **string** | The template name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrieveTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **locale** | **string** | The locale for strings | 
 **withSettings** | **bool** | Retrieve template settings | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**TemplateResponseContract**](TemplateResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrieveTemplateSetting

> SettingResponseContract GroupPolicyRetrieveTemplateSetting(ctx, customerid, siteid, templateName, settingName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Retrieve an existing selected template setting

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    templateName := "templateName_example" // string | Name of the template
    settingName := "settingName_example" // string | The setting name
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrieveTemplateSetting(context.Background(), customerid, siteid, templateName, settingName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrieveTemplateSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrieveTemplateSetting`: SettingResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrieveTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**templateName** | **string** | Name of the template | 
**settingName** | **string** | The setting name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrieveTemplateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingResponseContract**](SettingResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrieveTemplateSettings

> []SettingResponseContract GroupPolicyRetrieveTemplateSettings(ctx, customerid, siteid, templateName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get settings of a template

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    templateName := "templateName_example" // string | Name of template for which settings are retrieved
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrieveTemplateSettings(context.Background(), customerid, siteid, templateName).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrieveTemplateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrieveTemplateSettings`: []SettingResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrieveTemplateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**templateName** | **string** | Name of template for which settings are retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrieveTemplateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**[]SettingResponseContract**](SettingResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyRetrieveTemplates

> []TemplateResponseContract GroupPolicyRetrieveTemplates(ctx, customerid, siteid).Locale(locale).WithSettings(withSettings).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all templates

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string | The locale for strings
    withSettings := true // bool | Retrieve template settings
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyRetrieveTemplates(context.Background(), customerid, siteid).Locale(locale).WithSettings(withSettings).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyRetrieveTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyRetrieveTemplates`: []TemplateResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyRetrieveTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyRetrieveTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | The locale for strings | 
 **withSettings** | **bool** | Retrieve template settings | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**[]TemplateResponseContract**](TemplateResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicySelectPolicySetting

> SettingResponseContract GroupPolicySelectPolicySetting(ctx, customerid, siteid, policyName).Locale(locale).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Select a setting and configure it for the specified policy.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    locale := "locale_example" // string | Locale for displayed strings
    policyName := "policyName_example" // string | Name of the policy
    request := *openapiclient.NewSettingRequestContract() // SettingRequestContract | Setting to be selected
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicySelectPolicySetting(context.Background(), customerid, siteid, policyName).Locale(locale).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicySelectPolicySetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicySelectPolicySetting`: SettingResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicySelectPolicySetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | Name of the policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicySelectPolicySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** | Locale for displayed strings | 

 **request** | [**SettingRequestContract**](SettingRequestContract.md) | Setting to be selected | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingResponseContract**](SettingResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicySelectTemplateSetting

> SettingResponseContract GroupPolicySelectTemplateSetting(ctx, customerid, siteid, templateName).ForceWrite(forceWrite).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Select a setting and configure it for the specified template.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    templateName := "templateName_example" // string | Name of the template
    forceWrite := true // bool | Force write to database
    request := *openapiclient.NewSettingRequestContract() // SettingRequestContract | Setting to be selected
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicySelectTemplateSetting(context.Background(), customerid, siteid, templateName).ForceWrite(forceWrite).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicySelectTemplateSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicySelectTemplateSetting`: SettingResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicySelectTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**templateName** | **string** | Name of the template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicySelectTemplateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forceWrite** | **bool** | Force write to database | 
 **request** | [**SettingRequestContract**](SettingRequestContract.md) | Setting to be selected | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingResponseContract**](SettingResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyUpdateFilter

> FilterResponseContract GroupPolicyUpdateFilter(ctx, customerid, siteid, policyName, filterName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a filter

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The policy name
    filterName := "filterName_example" // string | The filter name
    request := *openapiclient.NewFilterRequestContract() // FilterRequestContract | New filter data
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyUpdateFilter(context.Background(), customerid, siteid, policyName, filterName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyUpdateFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyUpdateFilter`: FilterResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyUpdateFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The policy name | 
**filterName** | **string** | The filter name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyUpdateFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**FilterRequestContract**](FilterRequestContract.md) | New filter data | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**FilterResponseContract**](FilterResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyUpdatePolicy

> PolicyResponseContract GroupPolicyUpdatePolicy(ctx, customerid, siteid, policyName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an existing policy. If a property is null, the property is not updated. If a collection is not null, the entire collection is replaced. This API can be used to rename a policy and raise or lower a policy priority.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The name of an existing policy
    request := *openapiclient.NewPolicyRequestContract() // PolicyRequestContract | The new policy data
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyUpdatePolicy(context.Background(), customerid, siteid, policyName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyUpdatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyUpdatePolicy`: PolicyResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyUpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The name of an existing policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**PolicyRequestContract**](PolicyRequestContract.md) | The new policy data | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PolicyResponseContract**](PolicyResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyUpdatePolicySetting

> SettingResponseContract GroupPolicyUpdatePolicySetting(ctx, customerid, siteid, policyName, settingName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a selected policy setting.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    policyName := "policyName_example" // string | The policy name
    settingName := "settingName_example" // string | The setting name
    request := *openapiclient.NewSettingRequestContract() // SettingRequestContract | New setting data
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyUpdatePolicySetting(context.Background(), customerid, siteid, policyName, settingName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyUpdatePolicySetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyUpdatePolicySetting`: SettingResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyUpdatePolicySetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**policyName** | **string** | The policy name | 
**settingName** | **string** | The setting name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyUpdatePolicySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**SettingRequestContract**](SettingRequestContract.md) | New setting data | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingResponseContract**](SettingResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyUpdateTemplate

> TemplateResponseContract GroupPolicyUpdateTemplate(ctx, customerid, siteid, templateName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a template

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    templateName := "templateName_example" // string | The template name
    request := *openapiclient.NewTemplateRequestContract() // TemplateRequestContract | New template data
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyUpdateTemplate(context.Background(), customerid, siteid, templateName).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyUpdateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyUpdateTemplate`: TemplateResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyUpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**templateName** | **string** | The template name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**TemplateRequestContract**](TemplateRequestContract.md) | New template data | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**TemplateResponseContract**](TemplateResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPolicyUpdateTemplateSetting

> SettingResponseContract GroupPolicyUpdateTemplateSetting(ctx, customerid, siteid, templateName, settingName).ForceWrite(forceWrite).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a selected template setting.

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
    customerid := "customerid_example" // string | The customer id of the customer making this rest call
    siteid := "siteid_example" // string | The site id of the customer's site making this rest call
    templateName := "templateName_example" // string | The template name
    settingName := "settingName_example" // string | The setting name
    forceWrite := true // bool | Force write to database
    request := *openapiclient.NewSettingRequestContract() // SettingRequestContract | New setting data
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPolicyApi.GroupPolicyUpdateTemplateSetting(context.Background(), customerid, siteid, templateName, settingName).ForceWrite(forceWrite).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPolicyApi.GroupPolicyUpdateTemplateSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupPolicyUpdateTemplateSetting`: SettingResponseContract
    fmt.Fprintf(os.Stdout, "Response from `GroupPolicyApi.GroupPolicyUpdateTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer id of the customer making this rest call | 
**siteid** | **string** | The site id of the customer&#39;s site making this rest call | 
**templateName** | **string** | The template name | 
**settingName** | **string** | The setting name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPolicyUpdateTemplateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **forceWrite** | **bool** | Force write to database | 
 **request** | [**SettingRequestContract**](SettingRequestContract.md) | New setting data | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SettingResponseContract**](SettingResponseContract.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

