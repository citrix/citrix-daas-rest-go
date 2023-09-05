# AdministratorRightResponseModelRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of role. Needs to be made globally unique (perhaps add site id) | 
**Name** | **string** | Name of role. | 
**Description** | Pointer to **string** | Description of role. | [optional] 
**IsBuiltIn** | **bool** | Indicates whether the role is built-in. | 
**CanLaunchManage** | Pointer to **bool** | Indicate that if the mangement page could be launch on xdconsole | [optional] 
**CanLaunchMonitor** | Pointer to **bool** | Indicate that if the monitor page could be launch on xdconsole | [optional] 
**Permissions** | Pointer to [**[]DelegatedAdminPermissionResponseModel**](DelegatedAdminPermissionResponseModel.md) | List of permissions granted by the role. | [optional] 

## Methods

### NewAdministratorRightResponseModelRole

`func NewAdministratorRightResponseModelRole(id string, name string, isBuiltIn bool, ) *AdministratorRightResponseModelRole`

NewAdministratorRightResponseModelRole instantiates a new AdministratorRightResponseModelRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorRightResponseModelRoleWithDefaults

`func NewAdministratorRightResponseModelRoleWithDefaults() *AdministratorRightResponseModelRole`

NewAdministratorRightResponseModelRoleWithDefaults instantiates a new AdministratorRightResponseModelRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdministratorRightResponseModelRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdministratorRightResponseModelRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdministratorRightResponseModelRole) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AdministratorRightResponseModelRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdministratorRightResponseModelRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdministratorRightResponseModelRole) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AdministratorRightResponseModelRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdministratorRightResponseModelRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdministratorRightResponseModelRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdministratorRightResponseModelRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsBuiltIn

`func (o *AdministratorRightResponseModelRole) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *AdministratorRightResponseModelRole) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *AdministratorRightResponseModelRole) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.


### GetCanLaunchManage

`func (o *AdministratorRightResponseModelRole) GetCanLaunchManage() bool`

GetCanLaunchManage returns the CanLaunchManage field if non-nil, zero value otherwise.

### GetCanLaunchManageOk

`func (o *AdministratorRightResponseModelRole) GetCanLaunchManageOk() (*bool, bool)`

GetCanLaunchManageOk returns a tuple with the CanLaunchManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchManage

`func (o *AdministratorRightResponseModelRole) SetCanLaunchManage(v bool)`

SetCanLaunchManage sets CanLaunchManage field to given value.

### HasCanLaunchManage

`func (o *AdministratorRightResponseModelRole) HasCanLaunchManage() bool`

HasCanLaunchManage returns a boolean if a field has been set.

### GetCanLaunchMonitor

`func (o *AdministratorRightResponseModelRole) GetCanLaunchMonitor() bool`

GetCanLaunchMonitor returns the CanLaunchMonitor field if non-nil, zero value otherwise.

### GetCanLaunchMonitorOk

`func (o *AdministratorRightResponseModelRole) GetCanLaunchMonitorOk() (*bool, bool)`

GetCanLaunchMonitorOk returns a tuple with the CanLaunchMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchMonitor

`func (o *AdministratorRightResponseModelRole) SetCanLaunchMonitor(v bool)`

SetCanLaunchMonitor sets CanLaunchMonitor field to given value.

### HasCanLaunchMonitor

`func (o *AdministratorRightResponseModelRole) HasCanLaunchMonitor() bool`

HasCanLaunchMonitor returns a boolean if a field has been set.

### GetPermissions

`func (o *AdministratorRightResponseModelRole) GetPermissions() []DelegatedAdminPermissionResponseModel`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AdministratorRightResponseModelRole) GetPermissionsOk() (*[]DelegatedAdminPermissionResponseModel, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AdministratorRightResponseModelRole) SetPermissions(v []DelegatedAdminPermissionResponseModel)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AdministratorRightResponseModelRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


