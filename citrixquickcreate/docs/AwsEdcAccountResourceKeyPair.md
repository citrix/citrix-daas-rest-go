# AwsEdcAccountResourceKeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | Pointer to **NullableString** | Name of Key Pari | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Vpc Tags | [optional] 

## Methods

### NewAwsEdcAccountResourceKeyPair

`func NewAwsEdcAccountResourceKeyPair() *AwsEdcAccountResourceKeyPair`

NewAwsEdcAccountResourceKeyPair instantiates a new AwsEdcAccountResourceKeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcAccountResourceKeyPairWithDefaults

`func NewAwsEdcAccountResourceKeyPairWithDefaults() *AwsEdcAccountResourceKeyPair`

NewAwsEdcAccountResourceKeyPairWithDefaults instantiates a new AwsEdcAccountResourceKeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *AwsEdcAccountResourceKeyPair) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *AwsEdcAccountResourceKeyPair) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *AwsEdcAccountResourceKeyPair) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *AwsEdcAccountResourceKeyPair) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *AwsEdcAccountResourceKeyPair) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *AwsEdcAccountResourceKeyPair) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetTags

`func (o *AwsEdcAccountResourceKeyPair) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwsEdcAccountResourceKeyPair) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwsEdcAccountResourceKeyPair) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwsEdcAccountResourceKeyPair) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AwsEdcAccountResourceKeyPair) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AwsEdcAccountResourceKeyPair) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


