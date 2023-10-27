# SearchFilter2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**PolicyProperty**](PolicyProperty.md) |  | 
**Value** | **string** | Value to match. For boolean properties this must be &#x60;\&quot;true\&quot;&#x60; or &#x60;\&quot;false\&quot;&#x60;. For string properties and enum values, this is case insensitive. | 
**Operator** | [**SearchOperator**](SearchOperator.md) |  | 

## Methods

### NewSearchFilter2

`func NewSearchFilter2(property PolicyProperty, value string, operator SearchOperator, ) *SearchFilter2`

NewSearchFilter2 instantiates a new SearchFilter2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFilter2WithDefaults

`func NewSearchFilter2WithDefaults() *SearchFilter2`

NewSearchFilter2WithDefaults instantiates a new SearchFilter2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *SearchFilter2) GetProperty() PolicyProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchFilter2) GetPropertyOk() (*PolicyProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchFilter2) SetProperty(v PolicyProperty)`

SetProperty sets Property field to given value.


### GetValue

`func (o *SearchFilter2) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchFilter2) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchFilter2) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *SearchFilter2) GetOperator() SearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SearchFilter2) GetOperatorOk() (*SearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SearchFilter2) SetOperator(v SearchOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


