# CreateAdminRoleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the role. Name is globally unique. | 
**Description** | Pointer to **NullableString** | The description of the admin role. | [optional] 
**CanLaunchManage** | Pointer to **NullableBool** | Indicates whether the role has access to the Manage tab in Citrix Cloud. | [optional] 
**CanLaunchMonitor** | Pointer to **NullableBool** | Indicates whether the role has access to the Monitor tab in Citrix Cloud. | [optional] 
**Permissions** | Pointer to **[]string** | List of permissions granted by the role. At least one permission is required. | [optional] 

## Methods

### NewCreateAdminRoleRequestModel

`func NewCreateAdminRoleRequestModel(name string, ) *CreateAdminRoleRequestModel`

NewCreateAdminRoleRequestModel instantiates a new CreateAdminRoleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdminRoleRequestModelWithDefaults

`func NewCreateAdminRoleRequestModelWithDefaults() *CreateAdminRoleRequestModel`

NewCreateAdminRoleRequestModelWithDefaults instantiates a new CreateAdminRoleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAdminRoleRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAdminRoleRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAdminRoleRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAdminRoleRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAdminRoleRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAdminRoleRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAdminRoleRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateAdminRoleRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateAdminRoleRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCanLaunchManage

`func (o *CreateAdminRoleRequestModel) GetCanLaunchManage() bool`

GetCanLaunchManage returns the CanLaunchManage field if non-nil, zero value otherwise.

### GetCanLaunchManageOk

`func (o *CreateAdminRoleRequestModel) GetCanLaunchManageOk() (*bool, bool)`

GetCanLaunchManageOk returns a tuple with the CanLaunchManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchManage

`func (o *CreateAdminRoleRequestModel) SetCanLaunchManage(v bool)`

SetCanLaunchManage sets CanLaunchManage field to given value.

### HasCanLaunchManage

`func (o *CreateAdminRoleRequestModel) HasCanLaunchManage() bool`

HasCanLaunchManage returns a boolean if a field has been set.

### SetCanLaunchManageNil

`func (o *CreateAdminRoleRequestModel) SetCanLaunchManageNil(b bool)`

 SetCanLaunchManageNil sets the value for CanLaunchManage to be an explicit nil

### UnsetCanLaunchManage
`func (o *CreateAdminRoleRequestModel) UnsetCanLaunchManage()`

UnsetCanLaunchManage ensures that no value is present for CanLaunchManage, not even an explicit nil
### GetCanLaunchMonitor

`func (o *CreateAdminRoleRequestModel) GetCanLaunchMonitor() bool`

GetCanLaunchMonitor returns the CanLaunchMonitor field if non-nil, zero value otherwise.

### GetCanLaunchMonitorOk

`func (o *CreateAdminRoleRequestModel) GetCanLaunchMonitorOk() (*bool, bool)`

GetCanLaunchMonitorOk returns a tuple with the CanLaunchMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchMonitor

`func (o *CreateAdminRoleRequestModel) SetCanLaunchMonitor(v bool)`

SetCanLaunchMonitor sets CanLaunchMonitor field to given value.

### HasCanLaunchMonitor

`func (o *CreateAdminRoleRequestModel) HasCanLaunchMonitor() bool`

HasCanLaunchMonitor returns a boolean if a field has been set.

### SetCanLaunchMonitorNil

`func (o *CreateAdminRoleRequestModel) SetCanLaunchMonitorNil(b bool)`

 SetCanLaunchMonitorNil sets the value for CanLaunchMonitor to be an explicit nil

### UnsetCanLaunchMonitor
`func (o *CreateAdminRoleRequestModel) UnsetCanLaunchMonitor()`

UnsetCanLaunchMonitor ensures that no value is present for CanLaunchMonitor, not even an explicit nil
### GetPermissions

`func (o *CreateAdminRoleRequestModel) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateAdminRoleRequestModel) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateAdminRoleRequestModel) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateAdminRoleRequestModel) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CreateAdminRoleRequestModel) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CreateAdminRoleRequestModel) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


