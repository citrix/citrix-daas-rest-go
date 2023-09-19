# EnumerationMemberContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Member name | [optional] 
**Value** | Pointer to **int32** | Integer value of the member. | [optional] 
**DisplayName** | Pointer to **NullableString** | Translated names of the member. | [optional] 

## Methods

### NewEnumerationMemberContract

`func NewEnumerationMemberContract() *EnumerationMemberContract`

NewEnumerationMemberContract instantiates a new EnumerationMemberContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumerationMemberContractWithDefaults

`func NewEnumerationMemberContractWithDefaults() *EnumerationMemberContract`

NewEnumerationMemberContractWithDefaults instantiates a new EnumerationMemberContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnumerationMemberContract) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumerationMemberContract) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumerationMemberContract) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnumerationMemberContract) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnumerationMemberContract) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnumerationMemberContract) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *EnumerationMemberContract) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnumerationMemberContract) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnumerationMemberContract) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnumerationMemberContract) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDisplayName

`func (o *EnumerationMemberContract) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnumerationMemberContract) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnumerationMemberContract) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EnumerationMemberContract) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *EnumerationMemberContract) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *EnumerationMemberContract) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


