# LicensingUserResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **NullableString** | Gets or sets the account name | [optional] 
**AccountSid** | Pointer to **NullableString** | Gets or sets the account SID  | [optional] 
**PermissionLevel** | Pointer to [**LicensingPermissionLevel**](LicensingPermissionLevel.md) |  | [optional] 
**IsGroup** | Pointer to **bool** | Gets or sets a value indicating whether it is an individual account or a group | [optional] 

## Methods

### NewLicensingUserResponseModel

`func NewLicensingUserResponseModel() *LicensingUserResponseModel`

NewLicensingUserResponseModel instantiates a new LicensingUserResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingUserResponseModelWithDefaults

`func NewLicensingUserResponseModelWithDefaults() *LicensingUserResponseModel`

NewLicensingUserResponseModelWithDefaults instantiates a new LicensingUserResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LicensingUserResponseModel) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LicensingUserResponseModel) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LicensingUserResponseModel) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LicensingUserResponseModel) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *LicensingUserResponseModel) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *LicensingUserResponseModel) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAccountSid

`func (o *LicensingUserResponseModel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *LicensingUserResponseModel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *LicensingUserResponseModel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *LicensingUserResponseModel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *LicensingUserResponseModel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *LicensingUserResponseModel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetPermissionLevel

`func (o *LicensingUserResponseModel) GetPermissionLevel() LicensingPermissionLevel`

GetPermissionLevel returns the PermissionLevel field if non-nil, zero value otherwise.

### GetPermissionLevelOk

`func (o *LicensingUserResponseModel) GetPermissionLevelOk() (*LicensingPermissionLevel, bool)`

GetPermissionLevelOk returns a tuple with the PermissionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionLevel

`func (o *LicensingUserResponseModel) SetPermissionLevel(v LicensingPermissionLevel)`

SetPermissionLevel sets PermissionLevel field to given value.

### HasPermissionLevel

`func (o *LicensingUserResponseModel) HasPermissionLevel() bool`

HasPermissionLevel returns a boolean if a field has been set.

### GetIsGroup

`func (o *LicensingUserResponseModel) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *LicensingUserResponseModel) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *LicensingUserResponseModel) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *LicensingUserResponseModel) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


