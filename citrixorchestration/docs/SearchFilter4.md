# SearchFilter4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**FilterProperty**](FilterProperty.md) |  | 
**Value** | **string** | Value to match. For boolean properties this must be &#x60;\&quot;true\&quot;&#x60; or &#x60;\&quot;false\&quot;&#x60;. For string properties and enum values, this is case insensitive. | 
**Operator** | [**SearchOperator**](SearchOperator.md) |  | 

## Methods

### NewSearchFilter4

`func NewSearchFilter4(property FilterProperty, value string, operator SearchOperator, ) *SearchFilter4`

NewSearchFilter4 instantiates a new SearchFilter4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFilter4WithDefaults

`func NewSearchFilter4WithDefaults() *SearchFilter4`

NewSearchFilter4WithDefaults instantiates a new SearchFilter4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *SearchFilter4) GetProperty() FilterProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchFilter4) GetPropertyOk() (*FilterProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchFilter4) SetProperty(v FilterProperty)`

SetProperty sets Property field to given value.


### GetValue

`func (o *SearchFilter4) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchFilter4) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchFilter4) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *SearchFilter4) GetOperator() SearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SearchFilter4) GetOperatorOk() (*SearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SearchFilter4) SetOperator(v SearchOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


