# AddManagedEncryptionKeyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**Region** | **string** |  | 
**Name** | **string** |  | 
**KeyOption** | [**ManagedEncryptionKeyOption**](ManagedEncryptionKeyOption.md) | Enum to describe option to either generate or upload a key | 
**KeyUpload** | Pointer to **NullableString** |  | [optional] 
**RotationPolicy** | Pointer to [**NullableManagedEncryptionKeyRotationPolicyModel**](ManagedEncryptionKeyRotationPolicyModel.md) | Optional rotation policy configuration for the managed encryption key | [optional] 

## Methods

### NewAddManagedEncryptionKeyModel

`func NewAddManagedEncryptionKeyModel(subscriptionId string, region string, name string, keyOption ManagedEncryptionKeyOption, ) *AddManagedEncryptionKeyModel`

NewAddManagedEncryptionKeyModel instantiates a new AddManagedEncryptionKeyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddManagedEncryptionKeyModelWithDefaults

`func NewAddManagedEncryptionKeyModelWithDefaults() *AddManagedEncryptionKeyModel`

NewAddManagedEncryptionKeyModelWithDefaults instantiates a new AddManagedEncryptionKeyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AddManagedEncryptionKeyModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AddManagedEncryptionKeyModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AddManagedEncryptionKeyModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetRegion

`func (o *AddManagedEncryptionKeyModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddManagedEncryptionKeyModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddManagedEncryptionKeyModel) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetName

`func (o *AddManagedEncryptionKeyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddManagedEncryptionKeyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddManagedEncryptionKeyModel) SetName(v string)`

SetName sets Name field to given value.


### GetKeyOption

`func (o *AddManagedEncryptionKeyModel) GetKeyOption() ManagedEncryptionKeyOption`

GetKeyOption returns the KeyOption field if non-nil, zero value otherwise.

### GetKeyOptionOk

`func (o *AddManagedEncryptionKeyModel) GetKeyOptionOk() (*ManagedEncryptionKeyOption, bool)`

GetKeyOptionOk returns a tuple with the KeyOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyOption

`func (o *AddManagedEncryptionKeyModel) SetKeyOption(v ManagedEncryptionKeyOption)`

SetKeyOption sets KeyOption field to given value.


### GetKeyUpload

`func (o *AddManagedEncryptionKeyModel) GetKeyUpload() string`

GetKeyUpload returns the KeyUpload field if non-nil, zero value otherwise.

### GetKeyUploadOk

`func (o *AddManagedEncryptionKeyModel) GetKeyUploadOk() (*string, bool)`

GetKeyUploadOk returns a tuple with the KeyUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUpload

`func (o *AddManagedEncryptionKeyModel) SetKeyUpload(v string)`

SetKeyUpload sets KeyUpload field to given value.

### HasKeyUpload

`func (o *AddManagedEncryptionKeyModel) HasKeyUpload() bool`

HasKeyUpload returns a boolean if a field has been set.

### SetKeyUploadNil

`func (o *AddManagedEncryptionKeyModel) SetKeyUploadNil(b bool)`

 SetKeyUploadNil sets the value for KeyUpload to be an explicit nil

### UnsetKeyUpload
`func (o *AddManagedEncryptionKeyModel) UnsetKeyUpload()`

UnsetKeyUpload ensures that no value is present for KeyUpload, not even an explicit nil
### GetRotationPolicy

`func (o *AddManagedEncryptionKeyModel) GetRotationPolicy() ManagedEncryptionKeyRotationPolicyModel`

GetRotationPolicy returns the RotationPolicy field if non-nil, zero value otherwise.

### GetRotationPolicyOk

`func (o *AddManagedEncryptionKeyModel) GetRotationPolicyOk() (*ManagedEncryptionKeyRotationPolicyModel, bool)`

GetRotationPolicyOk returns a tuple with the RotationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPolicy

`func (o *AddManagedEncryptionKeyModel) SetRotationPolicy(v ManagedEncryptionKeyRotationPolicyModel)`

SetRotationPolicy sets RotationPolicy field to given value.

### HasRotationPolicy

`func (o *AddManagedEncryptionKeyModel) HasRotationPolicy() bool`

HasRotationPolicy returns a boolean if a field has been set.

### SetRotationPolicyNil

`func (o *AddManagedEncryptionKeyModel) SetRotationPolicyNil(b bool)`

 SetRotationPolicyNil sets the value for RotationPolicy to be an explicit nil

### UnsetRotationPolicy
`func (o *AddManagedEncryptionKeyModel) UnsetRotationPolicy()`

UnsetRotationPolicy ensures that no value is present for RotationPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


