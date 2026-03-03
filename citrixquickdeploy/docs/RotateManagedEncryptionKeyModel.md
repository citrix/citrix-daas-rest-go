# RotateManagedEncryptionKeyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyOption** | [**ManagedEncryptionKeyOption**](ManagedEncryptionKeyOption.md) | Key option for rotation - Generate new key or Import existing key | 
**KeyUpload** | Pointer to **NullableString** | Key to upload (required only when KeyOption is Import) | [optional] 

## Methods

### NewRotateManagedEncryptionKeyModel

`func NewRotateManagedEncryptionKeyModel(keyOption ManagedEncryptionKeyOption, ) *RotateManagedEncryptionKeyModel`

NewRotateManagedEncryptionKeyModel instantiates a new RotateManagedEncryptionKeyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateManagedEncryptionKeyModelWithDefaults

`func NewRotateManagedEncryptionKeyModelWithDefaults() *RotateManagedEncryptionKeyModel`

NewRotateManagedEncryptionKeyModelWithDefaults instantiates a new RotateManagedEncryptionKeyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyOption

`func (o *RotateManagedEncryptionKeyModel) GetKeyOption() ManagedEncryptionKeyOption`

GetKeyOption returns the KeyOption field if non-nil, zero value otherwise.

### GetKeyOptionOk

`func (o *RotateManagedEncryptionKeyModel) GetKeyOptionOk() (*ManagedEncryptionKeyOption, bool)`

GetKeyOptionOk returns a tuple with the KeyOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyOption

`func (o *RotateManagedEncryptionKeyModel) SetKeyOption(v ManagedEncryptionKeyOption)`

SetKeyOption sets KeyOption field to given value.


### GetKeyUpload

`func (o *RotateManagedEncryptionKeyModel) GetKeyUpload() string`

GetKeyUpload returns the KeyUpload field if non-nil, zero value otherwise.

### GetKeyUploadOk

`func (o *RotateManagedEncryptionKeyModel) GetKeyUploadOk() (*string, bool)`

GetKeyUploadOk returns a tuple with the KeyUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUpload

`func (o *RotateManagedEncryptionKeyModel) SetKeyUpload(v string)`

SetKeyUpload sets KeyUpload field to given value.

### HasKeyUpload

`func (o *RotateManagedEncryptionKeyModel) HasKeyUpload() bool`

HasKeyUpload returns a boolean if a field has been set.

### SetKeyUploadNil

`func (o *RotateManagedEncryptionKeyModel) SetKeyUploadNil(b bool)`

 SetKeyUploadNil sets the value for KeyUpload to be an explicit nil

### UnsetKeyUpload
`func (o *RotateManagedEncryptionKeyModel) UnsetKeyUpload()`

UnsetKeyUpload ensures that no value is present for KeyUpload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


