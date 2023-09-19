# UpdateAdminAdministratorRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rights** | Pointer to [**[]AdminRightRequestModel**](AdminRightRequestModel.md) | Rights associated with the administrator. | [optional] 
**Enabled** | Pointer to **NullableBool** | Indicates whether the administrator is enabled.  Disabled administrators cannot administer the site unless they are a member of a different user group which is granted access by a different administrator record. | [optional] 

## Methods

### NewUpdateAdminAdministratorRequestModel

`func NewUpdateAdminAdministratorRequestModel() *UpdateAdminAdministratorRequestModel`

NewUpdateAdminAdministratorRequestModel instantiates a new UpdateAdminAdministratorRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdminAdministratorRequestModelWithDefaults

`func NewUpdateAdminAdministratorRequestModelWithDefaults() *UpdateAdminAdministratorRequestModel`

NewUpdateAdminAdministratorRequestModelWithDefaults instantiates a new UpdateAdminAdministratorRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRights

`func (o *UpdateAdminAdministratorRequestModel) GetRights() []AdminRightRequestModel`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *UpdateAdminAdministratorRequestModel) GetRightsOk() (*[]AdminRightRequestModel, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *UpdateAdminAdministratorRequestModel) SetRights(v []AdminRightRequestModel)`

SetRights sets Rights field to given value.

### HasRights

`func (o *UpdateAdminAdministratorRequestModel) HasRights() bool`

HasRights returns a boolean if a field has been set.

### SetRightsNil

`func (o *UpdateAdminAdministratorRequestModel) SetRightsNil(b bool)`

 SetRightsNil sets the value for Rights to be an explicit nil

### UnsetRights
`func (o *UpdateAdminAdministratorRequestModel) UnsetRights()`

UnsetRights ensures that no value is present for Rights, not even an explicit nil
### GetEnabled

`func (o *UpdateAdminAdministratorRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAdminAdministratorRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAdminAdministratorRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAdminAdministratorRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UpdateAdminAdministratorRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateAdminAdministratorRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


