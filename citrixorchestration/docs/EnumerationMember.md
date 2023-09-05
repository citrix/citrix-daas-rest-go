# EnumerationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Member name | [optional] 
**Value** | **int32** | Integer value of the member. | 
**DisplayName** | Pointer to **string** | localized name of the member. | [optional] 

## Methods

### NewEnumerationMember

`func NewEnumerationMember(value int32, ) *EnumerationMember`

NewEnumerationMember instantiates a new EnumerationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumerationMemberWithDefaults

`func NewEnumerationMemberWithDefaults() *EnumerationMember`

NewEnumerationMemberWithDefaults instantiates a new EnumerationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnumerationMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumerationMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumerationMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnumerationMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *EnumerationMember) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnumerationMember) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnumerationMember) SetValue(v int32)`

SetValue sets Value field to given value.


### GetDisplayName

`func (o *EnumerationMember) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnumerationMember) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnumerationMember) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EnumerationMember) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


