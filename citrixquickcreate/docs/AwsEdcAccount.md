# AwsEdcAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByolAccount** | Pointer to **bool** | Indicates if the associated AWS EDC account has BYOL support enabled | [optional] 
**ByolCidrRange** | Pointer to **NullableString** | The CIDR range that is associated with this account | [optional] 
**AwsAccount** | Pointer to **NullableString** | ID of the AWS Account | [optional] 
**AwsRegion** | Pointer to **NullableString** | The AWS region the account is associated with | [optional] 
**AwsEdcMissingRolePermissions** | Pointer to [**[]AwsEdcMissingRolePermissions**](AwsEdcMissingRolePermissions.md) | Indicates the list of missing permissions in the role associated with the account | [optional] 
**AwsByolFeatureEnabled** | Pointer to **NullableBool** | Indicates if the associated AWS EDC account has BYOL feature is enabled | [optional] 
**AwsAuthType** | Pointer to [**AwsEdcAccountAuthType**](AwsEdcAccountAuthType.md) |  | [optional] 

## Methods

### NewAwsEdcAccount

`func NewAwsEdcAccount() *AwsEdcAccount`

NewAwsEdcAccount instantiates a new AwsEdcAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcAccountWithDefaults

`func NewAwsEdcAccountWithDefaults() *AwsEdcAccount`

NewAwsEdcAccountWithDefaults instantiates a new AwsEdcAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByolAccount

`func (o *AwsEdcAccount) GetByolAccount() bool`

GetByolAccount returns the ByolAccount field if non-nil, zero value otherwise.

### GetByolAccountOk

`func (o *AwsEdcAccount) GetByolAccountOk() (*bool, bool)`

GetByolAccountOk returns a tuple with the ByolAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByolAccount

`func (o *AwsEdcAccount) SetByolAccount(v bool)`

SetByolAccount sets ByolAccount field to given value.

### HasByolAccount

`func (o *AwsEdcAccount) HasByolAccount() bool`

HasByolAccount returns a boolean if a field has been set.

### GetByolCidrRange

`func (o *AwsEdcAccount) GetByolCidrRange() string`

GetByolCidrRange returns the ByolCidrRange field if non-nil, zero value otherwise.

### GetByolCidrRangeOk

`func (o *AwsEdcAccount) GetByolCidrRangeOk() (*string, bool)`

GetByolCidrRangeOk returns a tuple with the ByolCidrRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByolCidrRange

`func (o *AwsEdcAccount) SetByolCidrRange(v string)`

SetByolCidrRange sets ByolCidrRange field to given value.

### HasByolCidrRange

`func (o *AwsEdcAccount) HasByolCidrRange() bool`

HasByolCidrRange returns a boolean if a field has been set.

### SetByolCidrRangeNil

`func (o *AwsEdcAccount) SetByolCidrRangeNil(b bool)`

 SetByolCidrRangeNil sets the value for ByolCidrRange to be an explicit nil

### UnsetByolCidrRange
`func (o *AwsEdcAccount) UnsetByolCidrRange()`

UnsetByolCidrRange ensures that no value is present for ByolCidrRange, not even an explicit nil
### GetAwsAccount

`func (o *AwsEdcAccount) GetAwsAccount() string`

GetAwsAccount returns the AwsAccount field if non-nil, zero value otherwise.

### GetAwsAccountOk

`func (o *AwsEdcAccount) GetAwsAccountOk() (*string, bool)`

GetAwsAccountOk returns a tuple with the AwsAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccount

`func (o *AwsEdcAccount) SetAwsAccount(v string)`

SetAwsAccount sets AwsAccount field to given value.

### HasAwsAccount

`func (o *AwsEdcAccount) HasAwsAccount() bool`

HasAwsAccount returns a boolean if a field has been set.

### SetAwsAccountNil

`func (o *AwsEdcAccount) SetAwsAccountNil(b bool)`

 SetAwsAccountNil sets the value for AwsAccount to be an explicit nil

### UnsetAwsAccount
`func (o *AwsEdcAccount) UnsetAwsAccount()`

UnsetAwsAccount ensures that no value is present for AwsAccount, not even an explicit nil
### GetAwsRegion

`func (o *AwsEdcAccount) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AwsEdcAccount) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AwsEdcAccount) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AwsEdcAccount) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *AwsEdcAccount) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AwsEdcAccount) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetAwsEdcMissingRolePermissions

`func (o *AwsEdcAccount) GetAwsEdcMissingRolePermissions() []AwsEdcMissingRolePermissions`

GetAwsEdcMissingRolePermissions returns the AwsEdcMissingRolePermissions field if non-nil, zero value otherwise.

### GetAwsEdcMissingRolePermissionsOk

`func (o *AwsEdcAccount) GetAwsEdcMissingRolePermissionsOk() (*[]AwsEdcMissingRolePermissions, bool)`

GetAwsEdcMissingRolePermissionsOk returns a tuple with the AwsEdcMissingRolePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEdcMissingRolePermissions

`func (o *AwsEdcAccount) SetAwsEdcMissingRolePermissions(v []AwsEdcMissingRolePermissions)`

SetAwsEdcMissingRolePermissions sets AwsEdcMissingRolePermissions field to given value.

### HasAwsEdcMissingRolePermissions

`func (o *AwsEdcAccount) HasAwsEdcMissingRolePermissions() bool`

HasAwsEdcMissingRolePermissions returns a boolean if a field has been set.

### SetAwsEdcMissingRolePermissionsNil

`func (o *AwsEdcAccount) SetAwsEdcMissingRolePermissionsNil(b bool)`

 SetAwsEdcMissingRolePermissionsNil sets the value for AwsEdcMissingRolePermissions to be an explicit nil

### UnsetAwsEdcMissingRolePermissions
`func (o *AwsEdcAccount) UnsetAwsEdcMissingRolePermissions()`

UnsetAwsEdcMissingRolePermissions ensures that no value is present for AwsEdcMissingRolePermissions, not even an explicit nil
### GetAwsByolFeatureEnabled

`func (o *AwsEdcAccount) GetAwsByolFeatureEnabled() bool`

GetAwsByolFeatureEnabled returns the AwsByolFeatureEnabled field if non-nil, zero value otherwise.

### GetAwsByolFeatureEnabledOk

`func (o *AwsEdcAccount) GetAwsByolFeatureEnabledOk() (*bool, bool)`

GetAwsByolFeatureEnabledOk returns a tuple with the AwsByolFeatureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsByolFeatureEnabled

`func (o *AwsEdcAccount) SetAwsByolFeatureEnabled(v bool)`

SetAwsByolFeatureEnabled sets AwsByolFeatureEnabled field to given value.

### HasAwsByolFeatureEnabled

`func (o *AwsEdcAccount) HasAwsByolFeatureEnabled() bool`

HasAwsByolFeatureEnabled returns a boolean if a field has been set.

### SetAwsByolFeatureEnabledNil

`func (o *AwsEdcAccount) SetAwsByolFeatureEnabledNil(b bool)`

 SetAwsByolFeatureEnabledNil sets the value for AwsByolFeatureEnabled to be an explicit nil

### UnsetAwsByolFeatureEnabled
`func (o *AwsEdcAccount) UnsetAwsByolFeatureEnabled()`

UnsetAwsByolFeatureEnabled ensures that no value is present for AwsByolFeatureEnabled, not even an explicit nil
### GetAwsAuthType

`func (o *AwsEdcAccount) GetAwsAuthType() AwsEdcAccountAuthType`

GetAwsAuthType returns the AwsAuthType field if non-nil, zero value otherwise.

### GetAwsAuthTypeOk

`func (o *AwsEdcAccount) GetAwsAuthTypeOk() (*AwsEdcAccountAuthType, bool)`

GetAwsAuthTypeOk returns a tuple with the AwsAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAuthType

`func (o *AwsEdcAccount) SetAwsAuthType(v AwsEdcAccountAuthType)`

SetAwsAuthType sets AwsAuthType field to given value.

### HasAwsAuthType

`func (o *AwsEdcAccount) HasAwsAuthType() bool`

HasAwsAuthType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


