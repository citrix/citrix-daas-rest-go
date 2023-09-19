# CreateAdminAdministratorRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **NullableString** | The administrator to add. | [optional] 
**Rights** | Pointer to [**[]AdminRightRequestModel**](AdminRightRequestModel.md) | Rights associated with the administrator. It is ok if there is no rights specified. | [optional] 
**Enabled** | Pointer to **NullableBool** | Indicates whether the administrator is enabled.  Disabled administrators cannot administer the site unless they are a member of a different user group which is granted access by a different administrator record. | [optional] [default to true]

## Methods

### NewCreateAdminAdministratorRequestModel

`func NewCreateAdminAdministratorRequestModel() *CreateAdminAdministratorRequestModel`

NewCreateAdminAdministratorRequestModel instantiates a new CreateAdminAdministratorRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdminAdministratorRequestModelWithDefaults

`func NewCreateAdminAdministratorRequestModelWithDefaults() *CreateAdminAdministratorRequestModel`

NewCreateAdminAdministratorRequestModelWithDefaults instantiates a new CreateAdminAdministratorRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *CreateAdminAdministratorRequestModel) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateAdminAdministratorRequestModel) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateAdminAdministratorRequestModel) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateAdminAdministratorRequestModel) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *CreateAdminAdministratorRequestModel) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CreateAdminAdministratorRequestModel) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetRights

`func (o *CreateAdminAdministratorRequestModel) GetRights() []AdminRightRequestModel`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *CreateAdminAdministratorRequestModel) GetRightsOk() (*[]AdminRightRequestModel, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *CreateAdminAdministratorRequestModel) SetRights(v []AdminRightRequestModel)`

SetRights sets Rights field to given value.

### HasRights

`func (o *CreateAdminAdministratorRequestModel) HasRights() bool`

HasRights returns a boolean if a field has been set.

### SetRightsNil

`func (o *CreateAdminAdministratorRequestModel) SetRightsNil(b bool)`

 SetRightsNil sets the value for Rights to be an explicit nil

### UnsetRights
`func (o *CreateAdminAdministratorRequestModel) UnsetRights()`

UnsetRights ensures that no value is present for Rights, not even an explicit nil
### GetEnabled

`func (o *CreateAdminAdministratorRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAdminAdministratorRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAdminAdministratorRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateAdminAdministratorRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *CreateAdminAdministratorRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *CreateAdminAdministratorRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


