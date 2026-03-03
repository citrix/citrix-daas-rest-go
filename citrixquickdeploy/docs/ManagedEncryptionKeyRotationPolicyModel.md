# ManagedEncryptionKeyRotationPolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Gets or sets a value indicating whether key rotation is enabled | 
**RotationType** | Pointer to [**NullableManagedEncryptionKeyRotationPolicyType**](ManagedEncryptionKeyRotationPolicyType.md) | Gets or sets the rotation type  Required when Enabled is true | [optional] 
**ExpiryTime** | Pointer to **NullableString** | Gets or sets the expiry time in ISO 8601 duration format (e.g. \&quot;P2Y\&quot; for 2 years)  Used to set expiration date on newly rotated key.  Required when RotationType is TimeBeforeExpiry | [optional] 
**RotationTime** | Pointer to **NullableString** | Gets or sets the key rotation interval in ISO 8601 duration format (e.g. \&quot;P18M\&quot; for 18 months).  Required when Enabled is true. Minimum value is 7 days from creation or 7 days from expiration time. | [optional] 

## Methods

### NewManagedEncryptionKeyRotationPolicyModel

`func NewManagedEncryptionKeyRotationPolicyModel(enabled bool, ) *ManagedEncryptionKeyRotationPolicyModel`

NewManagedEncryptionKeyRotationPolicyModel instantiates a new ManagedEncryptionKeyRotationPolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedEncryptionKeyRotationPolicyModelWithDefaults

`func NewManagedEncryptionKeyRotationPolicyModelWithDefaults() *ManagedEncryptionKeyRotationPolicyModel`

NewManagedEncryptionKeyRotationPolicyModelWithDefaults instantiates a new ManagedEncryptionKeyRotationPolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ManagedEncryptionKeyRotationPolicyModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManagedEncryptionKeyRotationPolicyModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManagedEncryptionKeyRotationPolicyModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRotationType

`func (o *ManagedEncryptionKeyRotationPolicyModel) GetRotationType() ManagedEncryptionKeyRotationPolicyType`

GetRotationType returns the RotationType field if non-nil, zero value otherwise.

### GetRotationTypeOk

`func (o *ManagedEncryptionKeyRotationPolicyModel) GetRotationTypeOk() (*ManagedEncryptionKeyRotationPolicyType, bool)`

GetRotationTypeOk returns a tuple with the RotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationType

`func (o *ManagedEncryptionKeyRotationPolicyModel) SetRotationType(v ManagedEncryptionKeyRotationPolicyType)`

SetRotationType sets RotationType field to given value.

### HasRotationType

`func (o *ManagedEncryptionKeyRotationPolicyModel) HasRotationType() bool`

HasRotationType returns a boolean if a field has been set.

### SetRotationTypeNil

`func (o *ManagedEncryptionKeyRotationPolicyModel) SetRotationTypeNil(b bool)`

 SetRotationTypeNil sets the value for RotationType to be an explicit nil

### UnsetRotationType
`func (o *ManagedEncryptionKeyRotationPolicyModel) UnsetRotationType()`

UnsetRotationType ensures that no value is present for RotationType, not even an explicit nil
### GetExpiryTime

`func (o *ManagedEncryptionKeyRotationPolicyModel) GetExpiryTime() string`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ManagedEncryptionKeyRotationPolicyModel) GetExpiryTimeOk() (*string, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ManagedEncryptionKeyRotationPolicyModel) SetExpiryTime(v string)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ManagedEncryptionKeyRotationPolicyModel) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### SetExpiryTimeNil

`func (o *ManagedEncryptionKeyRotationPolicyModel) SetExpiryTimeNil(b bool)`

 SetExpiryTimeNil sets the value for ExpiryTime to be an explicit nil

### UnsetExpiryTime
`func (o *ManagedEncryptionKeyRotationPolicyModel) UnsetExpiryTime()`

UnsetExpiryTime ensures that no value is present for ExpiryTime, not even an explicit nil
### GetRotationTime

`func (o *ManagedEncryptionKeyRotationPolicyModel) GetRotationTime() string`

GetRotationTime returns the RotationTime field if non-nil, zero value otherwise.

### GetRotationTimeOk

`func (o *ManagedEncryptionKeyRotationPolicyModel) GetRotationTimeOk() (*string, bool)`

GetRotationTimeOk returns a tuple with the RotationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationTime

`func (o *ManagedEncryptionKeyRotationPolicyModel) SetRotationTime(v string)`

SetRotationTime sets RotationTime field to given value.

### HasRotationTime

`func (o *ManagedEncryptionKeyRotationPolicyModel) HasRotationTime() bool`

HasRotationTime returns a boolean if a field has been set.

### SetRotationTimeNil

`func (o *ManagedEncryptionKeyRotationPolicyModel) SetRotationTimeNil(b bool)`

 SetRotationTimeNil sets the value for RotationTime to be an explicit nil

### UnsetRotationTime
`func (o *ManagedEncryptionKeyRotationPolicyModel) UnsetRotationTime()`

UnsetRotationTime ensures that no value is present for RotationTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


