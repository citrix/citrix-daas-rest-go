# AddManagedEncryptionKeyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**Region** | **string** |  | 
**Name** | **string** |  | 
**KeyOption** | [**ManagedEncryptionKeyOption**](ManagedEncryptionKeyOption.md) | Enum to describe option to either generate or upload a key | 
**KeyUpload** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


