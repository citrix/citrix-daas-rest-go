# EnumerationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **NullableString** | Name of the type. | [optional] 
**Members** | Pointer to [**[]EnumerationMember**](EnumerationMember.md) | Members of the type. | [optional] 

## Methods

### NewEnumerationType

`func NewEnumerationType() *EnumerationType`

NewEnumerationType instantiates a new EnumerationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumerationTypeWithDefaults

`func NewEnumerationTypeWithDefaults() *EnumerationType`

NewEnumerationTypeWithDefaults instantiates a new EnumerationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *EnumerationType) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *EnumerationType) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *EnumerationType) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *EnumerationType) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *EnumerationType) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *EnumerationType) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetMembers

`func (o *EnumerationType) GetMembers() []EnumerationMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *EnumerationType) GetMembersOk() (*[]EnumerationMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *EnumerationType) SetMembers(v []EnumerationMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *EnumerationType) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *EnumerationType) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *EnumerationType) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


