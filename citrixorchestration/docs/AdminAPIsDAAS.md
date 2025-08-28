# \AdminAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminCheckRoleNameExists**](AdminAPIsDAAS.md#AdminCheckRoleNameExists) | **Head** /Admin/Roles/{name} | Check for the existence of role by name.
[**AdminCheckScopeNameExists**](AdminAPIsDAAS.md#AdminCheckScopeNameExists) | **Head** /Admin/Scopes/{name} | Check for the existence of scope by name.
[**AdminCreateAdminAdministrator**](AdminAPIsDAAS.md#AdminCreateAdminAdministrator) | **Post** /Admin/Administrators | Create a new administrator.
[**AdminCreateAdminRole**](AdminAPIsDAAS.md#AdminCreateAdminRole) | **Post** /Admin/Roles | Create a new admin role.
[**AdminCreateAdminScope**](AdminAPIsDAAS.md#AdminCreateAdminScope) | **Post** /Admin/Scopes | Create a new admin scope.
[**AdminDeleteAdminAdministrator**](AdminAPIsDAAS.md#AdminDeleteAdminAdministrator) | **Delete** /Admin/Administrators/{nameOrId} | Delete an administrator.
[**AdminDeleteAdminRole**](AdminAPIsDAAS.md#AdminDeleteAdminRole) | **Delete** /Admin/Roles/{nameOrId} | Delete an admin role.
[**AdminDeleteAdminScope**](AdminAPIsDAAS.md#AdminDeleteAdminScope) | **Delete** /Admin/Scopes/{nameOrId} | Delete an admin scope.
[**AdminGetAdminAdministrator**](AdminAPIsDAAS.md#AdminGetAdminAdministrator) | **Get** /Admin/Administrators/{nameOrId} | Get a single administrator
[**AdminGetAdminAdministrators**](AdminAPIsDAAS.md#AdminGetAdminAdministrators) | **Get** /Admin/Administrators | Get administrators.
[**AdminGetAdminEffectiveRights**](AdminAPIsDAAS.md#AdminGetAdminEffectiveRights) | **Get** /Admin/EffectiveRights | Get the effective rights of the current user.  This is the union of all rights of the enabled administrators that the current user matches, taking into account group membership.
[**AdminGetAdminRole**](AdminAPIsDAAS.md#AdminGetAdminRole) | **Get** /Admin/Roles/{nameOrId} | Get details about a single admin role.
[**AdminGetAdminRoles**](AdminAPIsDAAS.md#AdminGetAdminRoles) | **Get** /Admin/Roles | Get admin roles.
[**AdminGetAdminScope**](AdminAPIsDAAS.md#AdminGetAdminScope) | **Get** /Admin/Scopes/{nameOrId} | Get details about a single admin scope.
[**AdminGetAdminScopedObjects**](AdminAPIsDAAS.md#AdminGetAdminScopedObjects) | **Get** /Admin/Scopes/{nameOrId}/ScopedObjects | Get the objects in an admin scope.
[**AdminGetAdminScopes**](AdminAPIsDAAS.md#AdminGetAdminScopes) | **Get** /Admin/Scopes | Get admin scopes.
[**AdminGetAdministratorNamePreview**](AdminAPIsDAAS.md#AdminGetAdministratorNamePreview) | **Get** /Admin/Administrators/{name}/NamePreview | Get preview report of the administrator user name.
[**AdminGetAdministratorReport**](AdminAPIsDAAS.md#AdminGetAdministratorReport) | **Get** /Admin/Administrators/{nameOrId}/Report | Get report of the administrator.
[**AdminGetDeleteAdministratorConsequence**](AdminAPIsDAAS.md#AdminGetDeleteAdministratorConsequence) | **Get** /Admin/Administrators/{nameOrId}/PreviewDeleteConsequence | Preview the consequence of deleting an administrator.
[**AdminGetPredefinedPermissionGroups**](AdminAPIsDAAS.md#AdminGetPredefinedPermissionGroups) | **Get** /Admin/PermissionGroups | Get all permission groups.
[**AdminGetPredefinedPermissions**](AdminAPIsDAAS.md#AdminGetPredefinedPermissions) | **Get** /Admin/Permissions | Get all predefined permissions.
[**AdminGetPredefinedPermissionsForGroups**](AdminAPIsDAAS.md#AdminGetPredefinedPermissionsForGroups) | **Get** /Admin/PermissionGroups/{id}/Permissions | Get all permissions for a permission group.
[**AdminPreviewAdministratorReport**](AdminAPIsDAAS.md#AdminPreviewAdministratorReport) | **Post** /Admin/Administrators/$PreviewReport | Get preview report of the administrator.
[**AdminUpdateAdminAdministrator**](AdminAPIsDAAS.md#AdminUpdateAdminAdministrator) | **Patch** /Admin/Administrators/{nameOrId} | Update an administrator.
[**AdminUpdateAdminRole**](AdminAPIsDAAS.md#AdminUpdateAdminRole) | **Patch** /Admin/Roles/{nameOrId} | Update an admin role.
[**AdminUpdateAdminScope**](AdminAPIsDAAS.md#AdminUpdateAdminScope) | **Patch** /Admin/Scopes/{nameOrId} | Update an admin scope.



## AdminCheckRoleNameExists

> AdminCheckRoleNameExists(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	name := "name_example" // string | Name of the role.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminCheckRoleNameExists(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminCheckRoleNameExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCheckRoleNameExistsRequest struct via the builder pattern


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


## AdminCheckScopeNameExists

> AdminCheckScopeNameExists(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	name := "name_example" // string | Name of the scope.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminCheckScopeNameExists(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminCheckScopeNameExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCheckScopeNameExistsRequest struct via the builder pattern


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


## AdminCreateAdminAdministrator

> AdminCreateAdminAdministrator(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminAdministratorRequestModel(createAdminAdministratorRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	createAdminAdministratorRequestModel := *openapiclient.NewCreateAdminAdministratorRequestModel() // CreateAdminAdministratorRequestModel | Details of the admin to create.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin will be created as a background task.  The task will have JobType CreateAdmin.  When the task (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminCreateAdminAdministrator(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminAdministratorRequestModel(createAdminAdministratorRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminCreateAdminAdministrator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateAdminAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createAdminAdministratorRequestModel** | [**CreateAdminAdministratorRequestModel**](CreateAdminAdministratorRequestModel.md) | Details of the admin to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin will be created as a background task.  The task will have JobType CreateAdmin.  When the task | [default to false]

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


## AdminCreateAdminRole

> AdminCreateAdminRole(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminRoleRequestModel(createAdminRoleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	createAdminRoleRequestModel := *openapiclient.NewCreateAdminRoleRequestModel("Name_example") // CreateAdminRoleRequestModel | Details of the admin role to create.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin role will be created as a background task.  The task will have JobType CreateAdminRole.  When the task is complete it will redirect to GetAdminRole. The job Parameters will contain properties:              * _Name_ - Name of the admin role being created. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminCreateAdminRole(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminRoleRequestModel(createAdminRoleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminCreateAdminRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createAdminRoleRequestModel** | [**CreateAdminRoleRequestModel**](CreateAdminRoleRequestModel.md) | Details of the admin role to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin role will be created as a background task.  The task will have JobType CreateAdminRole.  When the task is complete it will redirect to GetAdminRole. The job Parameters will contain properties:              * _Name_ - Name of the admin role being created. | [default to false]

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


## AdminCreateAdminScope

> AdminCreateAdminScope(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminScopeRequestModel(createAdminScopeRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	createAdminScopeRequestModel := *openapiclient.NewCreateAdminScopeRequestModel("Name_example") // CreateAdminScopeRequestModel | Details of the admin scope to create.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin scope will be created as a background task. The task will have JobType CreateAdminScope. When the task is complete it will redirect to GetAdminScopes. The job Parameters will contain properties:              * _Name_ - Name of the admin scope being created. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminCreateAdminScope(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminScopeRequestModel(createAdminScopeRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminCreateAdminScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateAdminScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createAdminScopeRequestModel** | [**CreateAdminScopeRequestModel**](CreateAdminScopeRequestModel.md) | Details of the admin scope to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin scope will be created as a background task. The task will have JobType CreateAdminScope. When the task is complete it will redirect to GetAdminScopes. The job Parameters will contain properties:              * _Name_ - Name of the admin scope being created. | [default to false]

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


## AdminDeleteAdminAdministrator

> AdminDeleteAdminAdministrator(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or Id of the admin to delete.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin will be deleted as a background task. The task will have JobType DeleteAdmin. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminDeleteAdminAdministrator(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminDeleteAdminAdministrator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the admin to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteAdminAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin will be deleted as a background task. The task will have JobType DeleteAdmin. | [default to false]

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


## AdminDeleteAdminRole

> AdminDeleteAdminRole(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or ID of the admin role to delete.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	force := true // bool | Allow removal of roles that are still in use. (optional) (default to false)
	async := true // bool | If `true`, the admin role will be deleted as a background task.  The task will have JobType DeleteAdminRole.  When * _NameOrId_ - Name or ID of the admin role being deleted. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminDeleteAdminRole(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminDeleteAdminRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin role to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **force** | **bool** | Allow removal of roles that are still in use. | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, the admin role will be deleted as a background task.  The task will have JobType DeleteAdminRole.  When * _NameOrId_ - Name or ID of the admin role being deleted. | [default to false]

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


## AdminDeleteAdminScope

> AdminDeleteAdminScope(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or ID of the admin scope to delete.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	force := true // bool | Allow removal of scopes that are still in use. (optional) (default to false)
	async := true // bool | If `true`, the admin scope will be deleted as a background task. The task will have JobType DeleteAdminScope. When the task is complete it will redirect to GetAdminScopes. The job Parameters will contain properties:              * _Name_ - Name of the admin scope being deleted. * _Id_ - ID of the admin scope being deleted. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminDeleteAdminScope(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Force(force).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminDeleteAdminScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin scope to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteAdminScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **force** | **bool** | Allow removal of scopes that are still in use. | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, the admin scope will be deleted as a background task. The task will have JobType DeleteAdminScope. When the task is complete it will redirect to GetAdminScopes. The job Parameters will contain properties:              * _Name_ - Name of the admin scope being deleted. * _Id_ - ID of the admin scope being deleted. | [default to false]

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


## AdminGetAdminAdministrator

> AdministratorResponseModel AdminGetAdminAdministrator(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or id of the administrator.              May be specified as:                           * The SID of the admin user or group.  Note: in directory types              other than Active Directory, the SID is a              computed property, and is not related to any representation of that              user within Active Directory.  However it can still be useful if the              user already has a generated SID; for example, if copying users from              one object to another.                           * `Domain\\User` format.  This implies the directory type              Active Directory. If the Domain\\User is not              unique across AD Forests, the call will fail with an ambiguous name              error, status code 400.                           * `Forest\\Domain\\User` format.  This implies the directory type              Active Directory. This is the preferred form              of specifying an Active Directory user by name, as the name is              guaranteed to be unambiguous.              
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdminAdministrator(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdminAdministrator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdminAdministrator`: AdministratorResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdminAdministrator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or id of the administrator.              May be specified as:                           * The SID of the admin user or group.  Note: in directory types              other than Active Directory, the SID is a              computed property, and is not related to any representation of that              user within Active Directory.  However it can still be useful if the              user already has a generated SID; for example, if copying users from              one object to another.                           * &#x60;Domain\\User&#x60; format.  This implies the directory type              Active Directory. If the Domain\\User is not              unique across AD Forests, the call will fail with an ambiguous name              error, status code 400.                           * &#x60;Forest\\Domain\\User&#x60; format.  This implies the directory type              Active Directory. This is the preferred form              of specifying an Active Directory user by name, as the name is              guaranteed to be unambiguous.               | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdminAdministratorRequest struct via the builder pattern


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

[**AdministratorResponseModel**](AdministratorResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGetAdminAdministrators

> AdministratorResponseModelCollection AdminGetAdminAdministrators(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
	async := true // bool | If `true`, the administrators will be fetched as a background task. The task will have the JobTypeGetAdminAdministrators (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdminAdministrators(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdminAdministrators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdminAdministrators`: AdministratorResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdminAdministrators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdminAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the administrators will be fetched as a background task. The task will have the JobTypeGetAdminAdministrators | [default to false]

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


## AdminGetAdminEffectiveRights

> AdministratorRightResponseModelCollection AdminGetAdminEffectiveRights(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).IfNoneMatch(ifNoneMatch).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Optional ETag response header that was returned on the previous query. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of admin rights returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdminEffectiveRights(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).IfNoneMatch(ifNoneMatch).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdminEffectiveRights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdminEffectiveRights`: AdministratorRightResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdminEffectiveRights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdminEffectiveRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **ifNoneMatch** | **string** | Optional ETag response header that was returned on the previous query. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of admin rights returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 

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


## AdminGetAdminRole

> RoleResponseModel AdminGetAdminRole(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or ID of the admin role.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdminRole(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdminRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdminRole`: RoleResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdminRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdminRoleRequest struct via the builder pattern


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

[**RoleResponseModel**](RoleResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGetAdminRoles

> RoleResponseModelCollection AdminGetAdminRoles(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of admin roles returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdminRoles(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdminRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdminRoles`: RoleResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdminRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdminRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of admin roles returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 

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


## AdminGetAdminScope

> ScopeResponseModel AdminGetAdminScope(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or ID of the admin scope.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdminScope(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdminScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdminScope`: ScopeResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdminScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdminScopeRequest struct via the builder pattern


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

[**ScopeResponseModel**](ScopeResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGetAdminScopedObjects

> ScopedObjectResponseModelCollection AdminGetAdminScopedObjects(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or ID of the admin scope.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of objects returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdminScopedObjects(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdminScopedObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdminScopedObjects`: ScopedObjectResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdminScopedObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdminScopedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of objects returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 

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


## AdminGetAdminScopes

> ScopeResponseModelCollection AdminGetAdminScopes(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of admin scopes returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdminScopes(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdminScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdminScopes`: ScopeResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdminScopes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdminScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of admin scopes returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 

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


## AdminGetAdministratorNamePreview

> AdministratorNamePreviewResponseModel AdminGetAdministratorNamePreview(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).IgnoreFormatCheck(ignoreFormatCheck).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	name := "name_example" // string | Example: domain\\username or domain\\group
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	ignoreFormatCheck := true // bool | Ignore name format check. If true,             will only check whether AD account available and not conflicting to existing administrators.             Note, invalid name will still be reported, only report error message will be changed. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdministratorNamePreview(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).IgnoreFormatCheck(ignoreFormatCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdministratorNamePreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdministratorNamePreview`: AdministratorNamePreviewResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdministratorNamePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Example: domain\\username or domain\\group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdministratorNamePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **ignoreFormatCheck** | **bool** | Ignore name format check. If true,             will only check whether AD account available and not conflicting to existing administrators.             Note, invalid name will still be reported, only report error message will be changed. | [default to false]

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


## AdminGetAdministratorReport

> AdministratorReportResponseModel AdminGetAdministratorReport(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or Id of the admin to report.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetAdministratorReport(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetAdministratorReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdministratorReport`: AdministratorReportResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetAdministratorReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the admin to report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdministratorReportRequest struct via the builder pattern


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

[**AdministratorReportResponseModel**](AdministratorReportResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGetDeleteAdministratorConsequence

> AdministratorDeleteConcequenceResponseModel AdminGetDeleteAdministratorConsequence(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or Id of the admin to delete.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetDeleteAdministratorConsequence(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetDeleteAdministratorConsequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetDeleteAdministratorConsequence`: AdministratorDeleteConcequenceResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetDeleteAdministratorConsequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the admin to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetDeleteAdministratorConsequenceRequest struct via the builder pattern


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

[**AdministratorDeleteConcequenceResponseModel**](AdministratorDeleteConcequenceResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGetPredefinedPermissionGroups

> PredefinedPermissionGroupResponseModelCollection AdminGetPredefinedPermissionGroups(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of predefined permission groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetPredefinedPermissionGroups(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetPredefinedPermissionGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetPredefinedPermissionGroups`: PredefinedPermissionGroupResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetPredefinedPermissionGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetPredefinedPermissionGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of predefined permission groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 

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


## AdminGetPredefinedPermissions

> PredefinedPermissionResponseModelCollection AdminGetPredefinedPermissions(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of predefined permissions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetPredefinedPermissions(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetPredefinedPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetPredefinedPermissions`: PredefinedPermissionResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetPredefinedPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetPredefinedPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of predefined permissions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 

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


## AdminGetPredefinedPermissionsForGroups

> PredefinedPermissionResponseModelCollection AdminGetPredefinedPermissionsForGroups(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	id := "id_example" // string | ID of the admin permission group to query.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of predefined permissions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminGetPredefinedPermissionsForGroups(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminGetPredefinedPermissionsForGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetPredefinedPermissionsForGroups`: PredefinedPermissionResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminGetPredefinedPermissionsForGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the admin permission group to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetPredefinedPermissionsForGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of predefined permissions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 

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


## AdminPreviewAdministratorReport

> AdministratorReportResponseModel AdminPreviewAdministratorReport(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminAdministratorRequestModel(createAdminAdministratorRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	createAdminAdministratorRequestModel := *openapiclient.NewCreateAdminAdministratorRequestModel() // CreateAdminAdministratorRequestModel | The administrator creation request to preview.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPIsDAAS.AdminPreviewAdministratorReport(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminAdministratorRequestModel(createAdminAdministratorRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminPreviewAdministratorReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminPreviewAdministratorReport`: AdministratorReportResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminAPIsDAAS.AdminPreviewAdministratorReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminPreviewAdministratorReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createAdminAdministratorRequestModel** | [**CreateAdminAdministratorRequestModel**](CreateAdminAdministratorRequestModel.md) | The administrator creation request to preview. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## AdminUpdateAdminAdministrator

> AdminUpdateAdminAdministrator(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateAdminAdministratorRequestModel(updateAdminAdministratorRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name of the admin to update.
	updateAdminAdministratorRequestModel := *openapiclient.NewUpdateAdminAdministratorRequestModel() // UpdateAdminAdministratorRequestModel | Updated details of the admin.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin will be updated as a background task.  The task will have JobType UpdateAdmin.  When the task (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminUpdateAdminAdministrator(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateAdminAdministratorRequestModel(updateAdminAdministratorRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminUpdateAdminAdministrator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name of the admin to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateAdminAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **updateAdminAdministratorRequestModel** | [**UpdateAdminAdministratorRequestModel**](UpdateAdminAdministratorRequestModel.md) | Updated details of the admin. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin will be updated as a background task.  The task will have JobType UpdateAdmin.  When the task | [default to false]

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


## AdminUpdateAdminRole

> AdminUpdateAdminRole(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditAdminRoleRequestModel(editAdminRoleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or ID of the admin role to update.
	editAdminRoleRequestModel := *openapiclient.NewEditAdminRoleRequestModel() // EditAdminRoleRequestModel | Updated details of the admin role.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin scope will be updated as a background task. The task will have JobType UpdateAdminRole. When the task is complete it will redirect to GetAdminRole. The job Parameters will contain properties: * _NameOrId_ - Name or ID of the admin role being updated. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminUpdateAdminRole(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditAdminRoleRequestModel(editAdminRoleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminUpdateAdminRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin role to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **editAdminRoleRequestModel** | [**EditAdminRoleRequestModel**](EditAdminRoleRequestModel.md) | Updated details of the admin role. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin scope will be updated as a background task. The task will have JobType UpdateAdminRole. When the task is complete it will redirect to GetAdminRole. The job Parameters will contain properties: * _NameOrId_ - Name or ID of the admin role being updated. | [default to false]

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


## AdminUpdateAdminScope

> AdminUpdateAdminScope(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditAdminScopeRequestModel(editAdminScopeRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or ID of the admin scope to update.
	editAdminScopeRequestModel := *openapiclient.NewEditAdminScopeRequestModel() // EditAdminScopeRequestModel | Updated details of the admin scope.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin scope will be updated as a background task. The task will have JobType UpdateAdminScope. When the task is complete it will redirect to GetAdminScope. The job Parameters will contain properties:              * _Name_ - Name of the admin scope being updated. * _Id_ - ID of the admin scope being updated. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPIsDAAS.AdminUpdateAdminScope(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditAdminScopeRequestModel(editAdminScopeRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPIsDAAS.AdminUpdateAdminScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the admin scope to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateAdminScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **editAdminScopeRequestModel** | [**EditAdminScopeRequestModel**](EditAdminScopeRequestModel.md) | Updated details of the admin scope. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin scope will be updated as a background task. The task will have JobType UpdateAdminScope. When the task is complete it will redirect to GetAdminScope. The job Parameters will contain properties:              * _Name_ - Name of the admin scope being updated. * _Id_ - ID of the admin scope being updated. | [default to false]

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

