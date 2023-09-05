# RoleResponseModel

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

### NewRoleResponseModel

`func NewRoleResponseModel(id string, name string, isBuiltIn bool, ) *RoleResponseModel`

NewRoleResponseModel instantiates a new RoleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleResponseModelWithDefaults

`func NewRoleResponseModelWithDefaults() *RoleResponseModel`

NewRoleResponseModelWithDefaults instantiates a new RoleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RoleResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RoleResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsBuiltIn

`func (o *RoleResponseModel) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *RoleResponseModel) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *RoleResponseModel) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.


### GetCanLaunchManage

`func (o *RoleResponseModel) GetCanLaunchManage() bool`

GetCanLaunchManage returns the CanLaunchManage field if non-nil, zero value otherwise.

### GetCanLaunchManageOk

`func (o *RoleResponseModel) GetCanLaunchManageOk() (*bool, bool)`

GetCanLaunchManageOk returns a tuple with the CanLaunchManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchManage

`func (o *RoleResponseModel) SetCanLaunchManage(v bool)`

SetCanLaunchManage sets CanLaunchManage field to given value.

### HasCanLaunchManage

`func (o *RoleResponseModel) HasCanLaunchManage() bool`

HasCanLaunchManage returns a boolean if a field has been set.

### GetCanLaunchMonitor

`func (o *RoleResponseModel) GetCanLaunchMonitor() bool`

GetCanLaunchMonitor returns the CanLaunchMonitor field if non-nil, zero value otherwise.

### GetCanLaunchMonitorOk

`func (o *RoleResponseModel) GetCanLaunchMonitorOk() (*bool, bool)`

GetCanLaunchMonitorOk returns a tuple with the CanLaunchMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchMonitor

`func (o *RoleResponseModel) SetCanLaunchMonitor(v bool)`

SetCanLaunchMonitor sets CanLaunchMonitor field to given value.

### HasCanLaunchMonitor

`func (o *RoleResponseModel) HasCanLaunchMonitor() bool`

HasCanLaunchMonitor returns a boolean if a field has been set.

### GetPermissions

`func (o *RoleResponseModel) GetPermissions() []DelegatedAdminPermissionResponseModel`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleResponseModel) GetPermissionsOk() (*[]DelegatedAdminPermissionResponseModel, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleResponseModel) SetPermissions(v []DelegatedAdminPermissionResponseModel)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RoleResponseModel) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


