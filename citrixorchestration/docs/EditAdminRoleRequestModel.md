# EditAdminRoleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | The description of the admin role. | [optional] 
**Name** | Pointer to **NullableString** | The name of the role. Name is globally unique. | [optional] 
**CanLaunchManage** | Pointer to **NullableBool** | Indicate that if the mangement page could be launch on xdconsole | [optional] 
**CanLaunchMonitor** | Pointer to **NullableBool** | Indicate that if the monitor page could be launch on xdconsole | [optional] 
**Permissions** | Pointer to **[]string** | List of permissions granted by the role. At least one permission is required. | [optional] 

## Methods

### NewEditAdminRoleRequestModel

`func NewEditAdminRoleRequestModel() *EditAdminRoleRequestModel`

NewEditAdminRoleRequestModel instantiates a new EditAdminRoleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAdminRoleRequestModelWithDefaults

`func NewEditAdminRoleRequestModelWithDefaults() *EditAdminRoleRequestModel`

NewEditAdminRoleRequestModelWithDefaults instantiates a new EditAdminRoleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EditAdminRoleRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditAdminRoleRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditAdminRoleRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditAdminRoleRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EditAdminRoleRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditAdminRoleRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *EditAdminRoleRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditAdminRoleRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditAdminRoleRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditAdminRoleRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditAdminRoleRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditAdminRoleRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCanLaunchManage

`func (o *EditAdminRoleRequestModel) GetCanLaunchManage() bool`

GetCanLaunchManage returns the CanLaunchManage field if non-nil, zero value otherwise.

### GetCanLaunchManageOk

`func (o *EditAdminRoleRequestModel) GetCanLaunchManageOk() (*bool, bool)`

GetCanLaunchManageOk returns a tuple with the CanLaunchManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchManage

`func (o *EditAdminRoleRequestModel) SetCanLaunchManage(v bool)`

SetCanLaunchManage sets CanLaunchManage field to given value.

### HasCanLaunchManage

`func (o *EditAdminRoleRequestModel) HasCanLaunchManage() bool`

HasCanLaunchManage returns a boolean if a field has been set.

### SetCanLaunchManageNil

`func (o *EditAdminRoleRequestModel) SetCanLaunchManageNil(b bool)`

 SetCanLaunchManageNil sets the value for CanLaunchManage to be an explicit nil

### UnsetCanLaunchManage
`func (o *EditAdminRoleRequestModel) UnsetCanLaunchManage()`

UnsetCanLaunchManage ensures that no value is present for CanLaunchManage, not even an explicit nil
### GetCanLaunchMonitor

`func (o *EditAdminRoleRequestModel) GetCanLaunchMonitor() bool`

GetCanLaunchMonitor returns the CanLaunchMonitor field if non-nil, zero value otherwise.

### GetCanLaunchMonitorOk

`func (o *EditAdminRoleRequestModel) GetCanLaunchMonitorOk() (*bool, bool)`

GetCanLaunchMonitorOk returns a tuple with the CanLaunchMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchMonitor

`func (o *EditAdminRoleRequestModel) SetCanLaunchMonitor(v bool)`

SetCanLaunchMonitor sets CanLaunchMonitor field to given value.

### HasCanLaunchMonitor

`func (o *EditAdminRoleRequestModel) HasCanLaunchMonitor() bool`

HasCanLaunchMonitor returns a boolean if a field has been set.

### SetCanLaunchMonitorNil

`func (o *EditAdminRoleRequestModel) SetCanLaunchMonitorNil(b bool)`

 SetCanLaunchMonitorNil sets the value for CanLaunchMonitor to be an explicit nil

### UnsetCanLaunchMonitor
`func (o *EditAdminRoleRequestModel) UnsetCanLaunchMonitor()`

UnsetCanLaunchMonitor ensures that no value is present for CanLaunchMonitor, not even an explicit nil
### GetPermissions

`func (o *EditAdminRoleRequestModel) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EditAdminRoleRequestModel) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EditAdminRoleRequestModel) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *EditAdminRoleRequestModel) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *EditAdminRoleRequestModel) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *EditAdminRoleRequestModel) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


