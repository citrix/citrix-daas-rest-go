# Tenancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewTenancy

`func NewTenancy() *Tenancy`

NewTenancy instantiates a new Tenancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenancyWithDefaults

`func NewTenancyWithDefaults() *Tenancy`

NewTenancyWithDefaults instantiates a new Tenancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Tenancy) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Tenancy) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Tenancy) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Tenancy) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Tenancy) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Tenancy) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


