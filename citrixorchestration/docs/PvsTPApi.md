# \PvsTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PvsTPGetPvsCollections**](PvsTPApi.md#PvsTPGetPvsCollections) | **Get** /techpreview/{customerid}/{siteid}/Pvs/{serverAddress}/Collections | Get PVS collections within a site.
[**PvsTPGetPvsMachinesForCatalog**](PvsTPApi.md#PvsTPGetPvsMachinesForCatalog) | **Get** /techpreview/{customerid}/{siteid}/Pvs/{serverAddress}/MachinesForCatalog | Get PVS devices within a collection.
[**PvsTPGetPvsSites**](PvsTPApi.md#PvsTPGetPvsSites) | **Get** /techpreview/{customerid}/{siteid}/Pvs/{serverAddress}/Sites | Get PVS sites.



## PvsTPGetPvsCollections

> PvsCollectionResponseModelCollection PvsTPGetPvsCollections(ctx, serverAddress, customerid, siteid).SiteNameOrId(siteNameOrId).ResourceLocationNameOrId(resourceLocationNameOrId).Forest(forest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get PVS collections within a site.



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
    serverAddress := "serverAddress_example" // string | Name or IP of the PVS server.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    siteNameOrId := []string{"Inner_example"} // []string | Name or ID of the PVS site.  Optional.  If not specified, collections from all sites are returned. (optional)
    resourceLocationNameOrId := "resourceLocationNameOrId_example" // string | Name or ID of the resource location or zone through which to communicate to the PVS server. Optional; however, if this is not specified then communication to the PVS server may randomly fail based on network firewall rules between resource locations. Not used for on-premises deployments. (optional)
    forest := "forest_example" // string | Active Directory forest of the PVS server.  May be different than the AD forest of the machines managed by PVS.   Optional; however, if this is not specified then communication to the PVS server may randomly fail if the site is connected to untrusted forests, and communication is attempted through a forest that is not within the trust scope of the PVS server's AD forest.   Not used for on-premises deployments. (optional)
    async := true // bool | If `true`, the get pvs collections will be executed as a background task. The task will have JobType GetPvsCollections. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PvsTPApi.PvsTPGetPvsCollections(context.Background(), serverAddress, customerid, siteid).SiteNameOrId(siteNameOrId).ResourceLocationNameOrId(resourceLocationNameOrId).Forest(forest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PvsTPApi.PvsTPGetPvsCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PvsTPGetPvsCollections`: PvsCollectionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `PvsTPApi.PvsTPGetPvsCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverAddress** | **string** | Name or IP of the PVS server. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPvsTPGetPvsCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **siteNameOrId** | **[]string** | Name or ID of the PVS site.  Optional.  If not specified, collections from all sites are returned. | 
 **resourceLocationNameOrId** | **string** | Name or ID of the resource location or zone through which to communicate to the PVS server. Optional; however, if this is not specified then communication to the PVS server may randomly fail based on network firewall rules between resource locations. Not used for on-premises deployments. | 
 **forest** | **string** | Active Directory forest of the PVS server.  May be different than the AD forest of the machines managed by PVS.   Optional; however, if this is not specified then communication to the PVS server may randomly fail if the site is connected to untrusted forests, and communication is attempted through a forest that is not within the trust scope of the PVS server&#39;s AD forest.   Not used for on-premises deployments. | 
 **async** | **bool** | If &#x60;true&#x60;, the get pvs collections will be executed as a background task. The task will have JobType GetPvsCollections. When the task is complete it will redirect to GetJobResults. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PvsCollectionResponseModelCollection**](PvsCollectionResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvsTPGetPvsMachinesForCatalog

> PvsDeviceResponseModelCollection PvsTPGetPvsMachinesForCatalog(ctx, serverAddress, customerid, siteid).SiteNameOrId(siteNameOrId).CollectionNameOrId(collectionNameOrId).VirtualOnly(virtualOnly).ResourceLocationNameOrId(resourceLocationNameOrId).Forest(forest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get PVS devices within a collection.



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
    serverAddress := "serverAddress_example" // string | Name or IP of the PVS server.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    siteNameOrId := []string{"Inner_example"} // []string | Name or ID of the PVS site.  Optional; if not specified, devices from all sites are returned. (optional)
    collectionNameOrId := []string{"Inner_example"} // []string | Name or ID of the PVS collection.  Optional; if not specified, devices from all collections are returned. (optional)
    virtualOnly := true // bool | If `true` then the results will be limited to only machines which can be located on hypervisors connected to the site, and each returned device will include the  and  through which the machine may be power-managed.   If `false` or not specified then all devices are returned, and  and  properties will not be resolved. (optional) (default to false)
    resourceLocationNameOrId := "resourceLocationNameOrId_example" // string | Name or ID of the resource location or zone through which to communicate to the PVS server. Optional; however, if this is not specified then communication to the PVS server may randomly fail based on network firewall rules between resource locations. Not used for on-premises deployments. (optional)
    forest := "forest_example" // string | Active Directory forest of the PVS server.  May be different than the AD forest of the machines managed by PVS.   Optional; however, if this is not specified then communication to the PVS server may randomly fail if the site is connected to untrusted forests, and communication is attempted through a forest that is not within the trust scope of the PVS server's AD forest.   Not used for on-premises deployments. (optional)
    async := true // bool | If `true`, the get pvs machines will be executed as a background task. The task will have JobType GetPvsMachinesForCatalog. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PvsTPApi.PvsTPGetPvsMachinesForCatalog(context.Background(), serverAddress, customerid, siteid).SiteNameOrId(siteNameOrId).CollectionNameOrId(collectionNameOrId).VirtualOnly(virtualOnly).ResourceLocationNameOrId(resourceLocationNameOrId).Forest(forest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PvsTPApi.PvsTPGetPvsMachinesForCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PvsTPGetPvsMachinesForCatalog`: PvsDeviceResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `PvsTPApi.PvsTPGetPvsMachinesForCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverAddress** | **string** | Name or IP of the PVS server. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPvsTPGetPvsMachinesForCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **siteNameOrId** | **[]string** | Name or ID of the PVS site.  Optional; if not specified, devices from all sites are returned. | 
 **collectionNameOrId** | **[]string** | Name or ID of the PVS collection.  Optional; if not specified, devices from all collections are returned. | 
 **virtualOnly** | **bool** | If &#x60;true&#x60; then the results will be limited to only machines which can be located on hypervisors connected to the site, and each returned device will include the  and  through which the machine may be power-managed.   If &#x60;false&#x60; or not specified then all devices are returned, and  and  properties will not be resolved. | [default to false]
 **resourceLocationNameOrId** | **string** | Name or ID of the resource location or zone through which to communicate to the PVS server. Optional; however, if this is not specified then communication to the PVS server may randomly fail based on network firewall rules between resource locations. Not used for on-premises deployments. | 
 **forest** | **string** | Active Directory forest of the PVS server.  May be different than the AD forest of the machines managed by PVS.   Optional; however, if this is not specified then communication to the PVS server may randomly fail if the site is connected to untrusted forests, and communication is attempted through a forest that is not within the trust scope of the PVS server&#39;s AD forest.   Not used for on-premises deployments. | 
 **async** | **bool** | If &#x60;true&#x60;, the get pvs machines will be executed as a background task. The task will have JobType GetPvsMachinesForCatalog. When the task is complete it will redirect to GetJobResults. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PvsDeviceResponseModelCollection**](PvsDeviceResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvsTPGetPvsSites

> PvsSiteResponseModelCollection PvsTPGetPvsSites(ctx, serverAddress, customerid, siteid).ResourceLocationNameOrId(resourceLocationNameOrId).Forest(forest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Get PVS sites.



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
    serverAddress := "serverAddress_example" // string | Name or IP of the PVS server.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    resourceLocationNameOrId := "resourceLocationNameOrId_example" // string | Name or ID of the resource location or zone through which to communicate to the PVS server. Optional; however, if this is not specified then communication to the PVS server may randomly fail based on network firewall rules between resource locations. Not used for on-premises deployments. (optional)
    forest := "forest_example" // string | Active Directory forest of the PVS server.  May be different than the AD forest of the machines managed by PVS.   Optional; however, if this is not specified then communication to the PVS server may randomly fail if the site is connected to untrusted forests, and communication is attempted through a forest that is not within the trust scope of the PVS server's AD forest.   Not used for on-premises deployments. (optional)
    async := true // bool | If `true`, the get pvs sites will be executed as a background task. The task will have JobType GetPvsSites. When the task is complete it will redirect to GetJobResults. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PvsTPApi.PvsTPGetPvsSites(context.Background(), serverAddress, customerid, siteid).ResourceLocationNameOrId(resourceLocationNameOrId).Forest(forest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PvsTPApi.PvsTPGetPvsSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PvsTPGetPvsSites`: PvsSiteResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `PvsTPApi.PvsTPGetPvsSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverAddress** | **string** | Name or IP of the PVS server. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPvsTPGetPvsSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceLocationNameOrId** | **string** | Name or ID of the resource location or zone through which to communicate to the PVS server. Optional; however, if this is not specified then communication to the PVS server may randomly fail based on network firewall rules between resource locations. Not used for on-premises deployments. | 
 **forest** | **string** | Active Directory forest of the PVS server.  May be different than the AD forest of the machines managed by PVS.   Optional; however, if this is not specified then communication to the PVS server may randomly fail if the site is connected to untrusted forests, and communication is attempted through a forest that is not within the trust scope of the PVS server&#39;s AD forest.   Not used for on-premises deployments. | 
 **async** | **bool** | If &#x60;true&#x60;, the get pvs sites will be executed as a background task. The task will have JobType GetPvsSites. When the task is complete it will redirect to GetJobResults. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PvsSiteResponseModelCollection**](PvsSiteResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

