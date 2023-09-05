# \AdminTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminTPCheckRoleNameExists**](AdminTPApi.md#AdminTPCheckRoleNameExists) | **Head** /techpreview/{customerid}/{siteid}/Admin/Roles/{name} | Check for the existence of role by name.
[**AdminTPCheckScopeNameExists**](AdminTPApi.md#AdminTPCheckScopeNameExists) | **Head** /techpreview/{customerid}/{siteid}/Admin/Scopes/{name} | Check for the existence of scope by name.
[**AdminTPCreateAdminAdministrator**](AdminTPApi.md#AdminTPCreateAdminAdministrator) | **Post** /techpreview/{customerid}/{siteid}/Admin/Administrators | Create a new administrator.
[**AdminTPCreateAdminRole**](AdminTPApi.md#AdminTPCreateAdminRole) | **Post** /techpreview/{customerid}/{siteid}/Admin/Roles | Create a new admin role.
[**AdminTPCreateAdminScope**](AdminTPApi.md#AdminTPCreateAdminScope) | **Post** /techpreview/{customerid}/{siteid}/Admin/Scopes | Create a new admin scope.
[**AdminTPDeleteAdminAdministrator**](AdminTPApi.md#AdminTPDeleteAdminAdministrator) | **Delete** /techpreview/{customerid}/{siteid}/Admin/Administrators/{nameOrId} | Delete an administrator.
[**AdminTPDeleteAdminRole**](AdminTPApi.md#AdminTPDeleteAdminRole) | **Delete** /techpreview/{customerid}/{siteid}/Admin/Roles/{nameOrId} | Delete an admin role.
[**AdminTPDeleteAdminScope**](AdminTPApi.md#AdminTPDeleteAdminScope) | **Delete** /techpreview/{customerid}/{siteid}/Admin/Scopes/{nameOrId} | Delete an admin scope.
[**AdminTPGetAdminAdministrator**](AdminTPApi.md#AdminTPGetAdminAdministrator) | **Get** /techpreview/{customerid}/{siteid}/Admin/Administrators/{nameOrId} | Get a single administrator
[**AdminTPGetAdminAdministrators**](AdminTPApi.md#AdminTPGetAdminAdministrators) | **Get** /techpreview/{customerid}/{siteid}/Admin/Administrators | Get administrators.
[**AdminTPGetAdminEffectiveRights**](AdminTPApi.md#AdminTPGetAdminEffectiveRights) | **Get** /techpreview/{customerid}/{siteid}/Admin/EffectiveRights | Get the effective rights of the current user.  This is the union of all rights of the enabled administrators that the current user matches, taking into account group membership.
[**AdminTPGetAdminRole**](AdminTPApi.md#AdminTPGetAdminRole) | **Get** /techpreview/{customerid}/{siteid}/Admin/Roles/{nameOrId} | Get details about a single admin role.
[**AdminTPGetAdminRoles**](AdminTPApi.md#AdminTPGetAdminRoles) | **Get** /techpreview/{customerid}/{siteid}/Admin/Roles | Get admin roles.
[**AdminTPGetAdminScope**](AdminTPApi.md#AdminTPGetAdminScope) | **Get** /techpreview/{customerid}/{siteid}/Admin/Scopes/{nameOrId} | Get details about a single admin scope.
[**AdminTPGetAdminScopedObjects**](AdminTPApi.md#AdminTPGetAdminScopedObjects) | **Get** /techpreview/{customerid}/{siteid}/Admin/Scopes/{nameOrId}/ScopedObjects | Get the objects in an admin scope.
[**AdminTPGetAdminScopes**](AdminTPApi.md#AdminTPGetAdminScopes) | **Get** /techpreview/{customerid}/{siteid}/Admin/Scopes | Get admin scopes.
[**AdminTPGetAdministratorNamePreview**](AdminTPApi.md#AdminTPGetAdministratorNamePreview) | **Get** /techpreview/{customerid}/{siteid}/Admin/Administrators/{name}/NamePreview | Get preview report of the administrator user name.
[**AdminTPGetAdministratorReport**](AdminTPApi.md#AdminTPGetAdministratorReport) | **Get** /techpreview/{customerid}/{siteid}/Admin/Administrators/{nameOrId}/Report | Get report of the administrator.
[**AdminTPGetDeleteAdministratorConsequence**](AdminTPApi.md#AdminTPGetDeleteAdministratorConsequence) | **Get** /techpreview/{customerid}/{siteid}/Admin/Administrators/{nameOrId}/PreviewDeleteConsequence | Preview the consequence of deleting an administrator.
[**AdminTPGetPredefinedPermissionGroups**](AdminTPApi.md#AdminTPGetPredefinedPermissionGroups) | **Get** /techpreview/{customerid}/{siteid}/Admin/PermissionGroups | Get all permission groups.
[**AdminTPGetPredefinedPermissions**](AdminTPApi.md#AdminTPGetPredefinedPermissions) | **Get** /techpreview/{customerid}/{siteid}/Admin/Permissions | Get all predefined permissions.
[**AdminTPGetPredefinedPermissionsForGroups**](AdminTPApi.md#AdminTPGetPredefinedPermissionsForGroups) | **Get** /techpreview/{customerid}/{siteid}/Admin/PermissionGroups/{id}/Permissions | Get all permissions for a permission group.
[**AdminTPPreviewAdministratorReport**](AdminTPApi.md#AdminTPPreviewAdministratorReport) | **Post** /techpreview/{customerid}/{siteid}/Admin/Administrators/$PreviewReport | Get preview report of the administrator.
[**AdminTPUpdateAdminAdministrator**](AdminTPApi.md#AdminTPUpdateAdminAdministrator) | **Patch** /techpreview/{customerid}/{siteid}/Admin/Administrators/{nameOrId} | Update an administrator.
[**AdminTPUpdateAdminRole**](AdminTPApi.md#AdminTPUpdateAdminRole) | **Patch** /techpreview/{customerid}/{siteid}/Admin/Roles/{nameOrId} | Update an admin role.
[**AdminTPUpdateAdminScope**](AdminTPApi.md#AdminTPUpdateAdminScope) | **Patch** /techpreview/{customerid}/{siteid}/Admin/Scopes/{nameOrId} | Update an admin scope.



## AdminTPCheckRoleNameExists

> AdminTPCheckRoleNameExists(ctx, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of role by name.



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
    name := "name_example" // string | Name of the role.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPCheckRoleNameExists(context.Background(), name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPCheckRoleNameExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPCheckRoleNameExistsRequest struct via the builder pattern


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


## AdminTPCheckScopeNameExists

> AdminTPCheckScopeNameExists(ctx, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of scope by name.



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
    name := "name_example" // string | Name of the scope.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPCheckScopeNameExists(context.Background(), name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPCheckScopeNameExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the scope. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPCheckScopeNameExistsRequest struct via the builder pattern


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


## AdminTPCreateAdminAdministrator

> AdminTPCreateAdminAdministrator(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a new administrator.



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
    request := *openapiclient.NewCreateAdminAdministratorRequestModel() // CreateAdminAdministratorRequestModel | Details of the admin to create.
    async := true // bool | If `true`, the admin will be created as a background task.  The task will have JobType CreateAdmin.  When the task (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPCreateAdminAdministrator(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPCreateAdminAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPCreateAdminAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateAdminAdministratorRequestModel**](CreateAdminAdministratorRequestModel.md) | Details of the admin to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin will be created as a background task.  The task will have JobType CreateAdmin.  When the task | [default to false]
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


## AdminTPCreateAdminRole

> AdminTPCreateAdminRole(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a new admin role.



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
    request := *openapiclient.NewCreateAdminRoleRequestModel("Name_example") // CreateAdminRoleRequestModel | Details of the admin role to create.
    async := true // bool | If `true`, the admin role will be created as a background task.  The task will have JobType CreateAdminRole.  When the task is complete it will redirect to GetAdminRole. The job Parameters will contain properties:  * _Name_ - Name of the admin role being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPCreateAdminRole(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPCreateAdminRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPCreateAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateAdminRoleRequestModel**](CreateAdminRoleRequestModel.md) | Details of the admin role to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin role will be created as a background task.  The task will have JobType CreateAdminRole.  When the task is complete it will redirect to GetAdminRole. The job Parameters will contain properties:  * _Name_ - Name of the admin role being created. | [default to false]
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


## AdminTPCreateAdminScope

> AdminTPCreateAdminScope(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a new admin scope.



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
    request := *openapiclient.NewCreateAdminScopeRequestModel("Name_example") // CreateAdminScopeRequestModel | Details of the admin scope to create.
    async := true // bool | If `true`, the admin scope will be created as a background task. The task will have JobType CreateAdminScope. When the task is complete it will redirect to GetAdminScopes. The job Parameters will contain properties:  * _Name_ - Name of the admin scope being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPCreateAdminScope(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPCreateAdminScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPCreateAdminScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateAdminScopeRequestModel**](CreateAdminScopeRequestModel.md) | Details of the admin scope to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin scope will be created as a background task. The task will have JobType CreateAdminScope. When the task is complete it will redirect to GetAdminScopes. The job Parameters will contain properties:  * _Name_ - Name of the admin scope being created. | [default to false]
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


## AdminTPDeleteAdminAdministrator

> AdminTPDeleteAdminAdministrator(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an administrator.

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
    nameOrId := "nameOrId_example" // string | Name or Id of the admin to delete.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the admin will be deleted as a background task. The task will have JobType DeleteAdmin. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPDeleteAdminAdministrator(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPDeleteAdminAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the admin to delete. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPDeleteAdminAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the admin will be deleted as a background task. The task will have JobType DeleteAdmin. | [default to false]
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


## AdminTPDeleteAdminRole

> AdminTPDeleteAdminRole(ctx, nameOrId, customerid, siteid).Force(force).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an admin role.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the admin role to delete.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    force := true // bool | Allow removal of roles that are still in use. (optional) (default to false)
    async := true // bool | If `true`, the admin role will be deleted as a background task.  The task will have JobType DeleteAdminRole.  When * _NameOrId_ - Name or ID of the admin role being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPDeleteAdminRole(context.Background(), nameOrId, customerid, siteid).Force(force).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPDeleteAdminRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin role to delete. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPDeleteAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | Allow removal of roles that are still in use. | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, the admin role will be deleted as a background task.  The task will have JobType DeleteAdminRole.  When * _NameOrId_ - Name or ID of the admin role being deleted. | [default to false]
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


## AdminTPDeleteAdminScope

> AdminTPDeleteAdminScope(ctx, nameOrId, customerid, siteid).Force(force).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an admin scope.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the admin scope to delete.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    force := true // bool | Allow removal of scopes that are still in use. (optional) (default to false)
    async := true // bool | If `true`, the admin scope will be deleted as a background task. The task will have JobType DeleteAdminScope. When the task is complete it will redirect to GetAdminScopes. The job Parameters will contain properties:  * _Name_ - Name of the admin scope being deleted. * _Id_ - ID of the admin scope being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPDeleteAdminScope(context.Background(), nameOrId, customerid, siteid).Force(force).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPDeleteAdminScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin scope to delete. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPDeleteAdminScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | Allow removal of scopes that are still in use. | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, the admin scope will be deleted as a background task. The task will have JobType DeleteAdminScope. When the task is complete it will redirect to GetAdminScopes. The job Parameters will contain properties:  * _Name_ - Name of the admin scope being deleted. * _Id_ - ID of the admin scope being deleted. | [default to false]
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


## AdminTPGetAdminAdministrator

> AdministratorResponseModel AdminTPGetAdminAdministrator(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get a single administrator

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
    nameOrId := "nameOrId_example" // string | Name or id of the administrator.              May be specified as:                           * The SID of the admin user or group.  Note: in directory types              other than Active Directory, the SID is a              computed property, and is not related to any representation of that              user within Active Directory.  However it can still be useful if the              user already has a generated SID; for example, if copying users from              one object to another.                           * `Domain\\User` format.  This implies the directory type              Active Directory. If the Domain\\User is not              unique across AD Forests, the call will fail with an ambiguous name              error, status code 400.                           * `Forest\\Domain\\User` format.  This implies the directory type              Active Directory. This is the preferred form              of specifying an Active Directory user by name, as the name is              guaranteed to be unambiguous.              
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdminAdministrator(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdminAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdminAdministrator`: AdministratorResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdminAdministrator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or id of the administrator.              May be specified as:                           * The SID of the admin user or group.  Note: in directory types              other than Active Directory, the SID is a              computed property, and is not related to any representation of that              user within Active Directory.  However it can still be useful if the              user already has a generated SID; for example, if copying users from              one object to another.                           * &#x60;Domain\\User&#x60; format.  This implies the directory type              Active Directory. If the Domain\\User is not              unique across AD Forests, the call will fail with an ambiguous name              error, status code 400.                           * &#x60;Forest\\Domain\\User&#x60; format.  This implies the directory type              Active Directory. This is the preferred form              of specifying an Active Directory user by name, as the name is              guaranteed to be unambiguous.               | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdminAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdministratorResponseModel**](AdministratorResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetAdminAdministrators

> AdministratorResponseModelCollection AdminTPGetAdminAdministrators(ctx, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get administrators.



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
    limit := int32(56) // int32 | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the administrators will be fetched as a background task. The task will have the JobTypeGetAdminAdministrators (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdminAdministrators(context.Background(), customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdminAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdminAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdminAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdminAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the administrators will be fetched as a background task. The task will have the JobTypeGetAdminAdministrators | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdministratorResponseModelCollection**](AdministratorResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetAdminEffectiveRights

> AdministratorRightResponseModelCollection AdminTPGetAdminEffectiveRights(ctx, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).IfNoneMatch(ifNoneMatch).XActionName(xActionName).Execute()

Get the effective rights of the current user.  This is the union of all rights of the enabled administrators that the current user matches, taking into account group membership.

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
    limit := int32(56) // int32 | The max number of admin rights returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Optional ETag response header that was returned on the previous query. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdminEffectiveRights(context.Background(), customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).IfNoneMatch(ifNoneMatch).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdminEffectiveRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdminEffectiveRights`: AdministratorRightResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdminEffectiveRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdminEffectiveRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of admin rights returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **ifNoneMatch** | **string** | Optional ETag response header that was returned on the previous query. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdministratorRightResponseModelCollection**](AdministratorRightResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetAdminRole

> RoleResponseModel AdminTPGetAdminRole(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details about a single admin role.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the admin role.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdminRole(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdminRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdminRole`: RoleResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdminRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin role. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**RoleResponseModel**](RoleResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetAdminRoles

> RoleResponseModelCollection AdminTPGetAdminRoles(ctx, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get admin roles.



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
    limit := int32(56) // int32 | The max number of admin roles returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdminRoles(context.Background(), customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdminRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdminRoles`: RoleResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdminRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdminRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of admin roles returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**RoleResponseModelCollection**](RoleResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetAdminScope

> ScopeResponseModel AdminTPGetAdminScope(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details about a single admin scope.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the admin scope.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdminScope(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdminScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdminScope`: ScopeResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdminScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin scope. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdminScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ScopeResponseModel**](ScopeResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetAdminScopedObjects

> ScopedObjectResponseModelCollection AdminTPGetAdminScopedObjects(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the objects in an admin scope.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the admin scope.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of objects returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdminScopedObjects(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdminScopedObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdminScopedObjects`: ScopedObjectResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdminScopedObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin scope. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdminScopedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of objects returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ScopedObjectResponseModelCollection**](ScopedObjectResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetAdminScopes

> ScopeResponseModelCollection AdminTPGetAdminScopes(ctx, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get admin scopes.



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
    limit := int32(56) // int32 | The max number of admin scopes returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdminScopes(context.Background(), customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdminScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdminScopes`: ScopeResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdminScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdminScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of admin scopes returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ScopeResponseModelCollection**](ScopeResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetAdministratorNamePreview

> AdministratorNamePreviewResponseModel AdminTPGetAdministratorNamePreview(ctx, name, customerid, siteid).IgnoreFormatCheck(ignoreFormatCheck).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get preview report of the administrator user name.

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
    name := "name_example" // string | Example: domain\\username or domain\\group
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    ignoreFormatCheck := true // bool | Ignore name format check. If true,             will only check whether AD account available and not conflicting to existing administrators.             Note, invalid name will still be reported, only report error message will be changed. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdministratorNamePreview(context.Background(), name, customerid, siteid).IgnoreFormatCheck(ignoreFormatCheck).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdministratorNamePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdministratorNamePreview`: AdministratorNamePreviewResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdministratorNamePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Example: domain\\username or domain\\group | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdministratorNamePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ignoreFormatCheck** | **bool** | Ignore name format check. If true,             will only check whether AD account available and not conflicting to existing administrators.             Note, invalid name will still be reported, only report error message will be changed. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdministratorNamePreviewResponseModel**](AdministratorNamePreviewResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetAdministratorReport

> AdministratorReportResponseModel AdminTPGetAdministratorReport(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get report of the administrator.

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
    nameOrId := "nameOrId_example" // string | Name or Id of the admin to report.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetAdministratorReport(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetAdministratorReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetAdministratorReport`: AdministratorReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetAdministratorReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the admin to report. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetAdministratorReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdministratorReportResponseModel**](AdministratorReportResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetDeleteAdministratorConsequence

> AdministratorDeleteConcequenceResponseModel AdminTPGetDeleteAdministratorConsequence(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Preview the consequence of deleting an administrator.

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
    nameOrId := "nameOrId_example" // string | Name or Id of the admin to delete.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetDeleteAdministratorConsequence(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetDeleteAdministratorConsequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetDeleteAdministratorConsequence`: AdministratorDeleteConcequenceResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetDeleteAdministratorConsequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the admin to delete. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetDeleteAdministratorConsequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdministratorDeleteConcequenceResponseModel**](AdministratorDeleteConcequenceResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetPredefinedPermissionGroups

> PredefinedPermissionGroupResponseModelCollection AdminTPGetPredefinedPermissionGroups(ctx, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all permission groups.



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
    limit := int32(56) // int32 | The max number of predefined permission groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetPredefinedPermissionGroups(context.Background(), customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetPredefinedPermissionGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetPredefinedPermissionGroups`: PredefinedPermissionGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetPredefinedPermissionGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetPredefinedPermissionGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of predefined permission groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PredefinedPermissionGroupResponseModelCollection**](PredefinedPermissionGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetPredefinedPermissions

> PredefinedPermissionResponseModelCollection AdminTPGetPredefinedPermissions(ctx, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all predefined permissions.

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
    limit := int32(56) // int32 | The max number of predefined permissions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetPredefinedPermissions(context.Background(), customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetPredefinedPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetPredefinedPermissions`: PredefinedPermissionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetPredefinedPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetPredefinedPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of predefined permissions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PredefinedPermissionResponseModelCollection**](PredefinedPermissionResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPGetPredefinedPermissionsForGroups

> PredefinedPermissionResponseModelCollection AdminTPGetPredefinedPermissionsForGroups(ctx, id, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all permissions for a permission group.



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
    id := "id_example" // string | ID of the admin permission group to query.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of predefined permissions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPGetPredefinedPermissionsForGroups(context.Background(), id, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPGetPredefinedPermissionsForGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPGetPredefinedPermissionsForGroups`: PredefinedPermissionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPGetPredefinedPermissionsForGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the admin permission group to query. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPGetPredefinedPermissionsForGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of predefined permissions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PredefinedPermissionResponseModelCollection**](PredefinedPermissionResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPPreviewAdministratorReport

> AdministratorReportResponseModel AdminTPPreviewAdministratorReport(ctx, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get preview report of the administrator.

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
    request := *openapiclient.NewCreateAdminAdministratorRequestModel() // CreateAdminAdministratorRequestModel | The administrator creation request to preview.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminTPApi.AdminTPPreviewAdministratorReport(context.Background(), customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPPreviewAdministratorReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTPPreviewAdministratorReport`: AdministratorReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminTPApi.AdminTPPreviewAdministratorReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPPreviewAdministratorReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateAdminAdministratorRequestModel**](CreateAdminAdministratorRequestModel.md) | The administrator creation request to preview. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdministratorReportResponseModel**](AdministratorReportResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTPUpdateAdminAdministrator

> AdminTPUpdateAdminAdministrator(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an administrator.



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
    nameOrId := "nameOrId_example" // string | Name of the admin to update.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewUpdateAdminAdministratorRequestModel() // UpdateAdminAdministratorRequestModel | Updated details of the admin.
    async := true // bool | If `true`, the admin will be updated as a background task.  The task will have JobType UpdateAdmin.  When the task (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPUpdateAdminAdministrator(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPUpdateAdminAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name of the admin to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPUpdateAdminAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateAdminAdministratorRequestModel**](UpdateAdminAdministratorRequestModel.md) | Updated details of the admin. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin will be updated as a background task.  The task will have JobType UpdateAdmin.  When the task | [default to false]
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


## AdminTPUpdateAdminRole

> AdminTPUpdateAdminRole(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an admin role.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the admin role to update.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewEditAdminRoleRequestModel() // EditAdminRoleRequestModel | Updated details of the admin role.
    async := true // bool | If `true`, the admin scope will be updated as a background task. The task will have JobType UpdateAdminRole. When the task is complete it will redirect to GetAdminRole. The job Parameters will contain properties: * _NameOrId_ - Name or ID of the admin role being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPUpdateAdminRole(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPUpdateAdminRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin role to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPUpdateAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**EditAdminRoleRequestModel**](EditAdminRoleRequestModel.md) | Updated details of the admin role. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin scope will be updated as a background task. The task will have JobType UpdateAdminRole. When the task is complete it will redirect to GetAdminRole. The job Parameters will contain properties: * _NameOrId_ - Name or ID of the admin role being updated. | [default to false]
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


## AdminTPUpdateAdminScope

> AdminTPUpdateAdminScope(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an admin scope.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the admin scope to update.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewEditAdminScopeRequestModel() // EditAdminScopeRequestModel | Updated details of the admin scope.
    async := true // bool | If `true`, the admin scope will be updated as a background task. The task will have JobType UpdateAdminScope. When the task is complete it will redirect to GetAdminScope. The job Parameters will contain properties:  * _Name_ - Name of the admin scope being updated. * _Id_ - ID of the admin scope being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminTPApi.AdminTPUpdateAdminScope(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTPApi.AdminTPUpdateAdminScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin scope to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTPUpdateAdminScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**EditAdminScopeRequestModel**](EditAdminScopeRequestModel.md) | Updated details of the admin scope. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin scope will be updated as a background task. The task will have JobType UpdateAdminScope. When the task is complete it will redirect to GetAdminScope. The job Parameters will contain properties:  * _Name_ - Name of the admin scope being updated. * _Id_ - ID of the admin scope being updated. | [default to false]
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

