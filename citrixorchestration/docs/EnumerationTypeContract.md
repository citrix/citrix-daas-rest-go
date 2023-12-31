# EnumerationTypeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **NullableString** | Name of the type. | [optional] 
**Members** | Pointer to [**[]EnumerationMemberContract**](EnumerationMemberContract.md) | Members of the type. | [optional] 

## Methods

### NewEnumerationTypeContract

`func NewEnumerationTypeContract() *EnumerationTypeContract`

NewEnumerationTypeContract instantiates a new EnumerationTypeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumerationTypeContractWithDefaults

`func NewEnumerationTypeContractWithDefaults() *EnumerationTypeContract`

NewEnumerationTypeContractWithDefaults instantiates a new EnumerationTypeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *EnumerationTypeContract) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *EnumerationTypeContract) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *EnumerationTypeContract) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *EnumerationTypeContract) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *EnumerationTypeContract) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *EnumerationTypeContract) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetMembers

`func (o *EnumerationTypeContract) GetMembers() []EnumerationMemberContract`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *EnumerationTypeContract) GetMembersOk() (*[]EnumerationMemberContract, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *EnumerationTypeContract) SetMembers(v []EnumerationMemberContract)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *EnumerationTypeContract) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *EnumerationTypeContract) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *EnumerationTypeContract) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


