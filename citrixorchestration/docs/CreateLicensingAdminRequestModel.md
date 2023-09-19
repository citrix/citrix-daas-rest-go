# CreateLicensingAdminRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSamName** | **string** | The same name of the account | 
**AccountSid** | Pointer to **NullableString** | The sid of the account | [optional] 
**Permissions** | Pointer to [**LicensingPermissionLevel**](LicensingPermissionLevel.md) |  | [optional] 
**IsGroup** | Pointer to **bool** | Whether this account is a group account | [optional] 

## Methods

### NewCreateLicensingAdminRequestModel

`func NewCreateLicensingAdminRequestModel(accountSamName string, ) *CreateLicensingAdminRequestModel`

NewCreateLicensingAdminRequestModel instantiates a new CreateLicensingAdminRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLicensingAdminRequestModelWithDefaults

`func NewCreateLicensingAdminRequestModelWithDefaults() *CreateLicensingAdminRequestModel`

NewCreateLicensingAdminRequestModelWithDefaults instantiates a new CreateLicensingAdminRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSamName

`func (o *CreateLicensingAdminRequestModel) GetAccountSamName() string`

GetAccountSamName returns the AccountSamName field if non-nil, zero value otherwise.

### GetAccountSamNameOk

`func (o *CreateLicensingAdminRequestModel) GetAccountSamNameOk() (*string, bool)`

GetAccountSamNameOk returns a tuple with the AccountSamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSamName

`func (o *CreateLicensingAdminRequestModel) SetAccountSamName(v string)`

SetAccountSamName sets AccountSamName field to given value.


### GetAccountSid

`func (o *CreateLicensingAdminRequestModel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *CreateLicensingAdminRequestModel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *CreateLicensingAdminRequestModel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *CreateLicensingAdminRequestModel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *CreateLicensingAdminRequestModel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *CreateLicensingAdminRequestModel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetPermissions

`func (o *CreateLicensingAdminRequestModel) GetPermissions() LicensingPermissionLevel`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateLicensingAdminRequestModel) GetPermissionsOk() (*LicensingPermissionLevel, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateLicensingAdminRequestModel) SetPermissions(v LicensingPermissionLevel)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateLicensingAdminRequestModel) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetIsGroup

`func (o *CreateLicensingAdminRequestModel) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *CreateLicensingAdminRequestModel) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *CreateLicensingAdminRequestModel) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *CreateLicensingAdminRequestModel) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


