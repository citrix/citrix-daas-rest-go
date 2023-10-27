# SearchFilter3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**SettingProperty**](SettingProperty.md) |  | 
**Value** | **string** | Value to match. For boolean properties this must be &#x60;\&quot;true\&quot;&#x60; or &#x60;\&quot;false\&quot;&#x60;. For string properties and enum values, this is case insensitive. | 
**Operator** | [**SearchOperator**](SearchOperator.md) |  | 

## Methods

### NewSearchFilter3

`func NewSearchFilter3(property SettingProperty, value string, operator SearchOperator, ) *SearchFilter3`

NewSearchFilter3 instantiates a new SearchFilter3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFilter3WithDefaults

`func NewSearchFilter3WithDefaults() *SearchFilter3`

NewSearchFilter3WithDefaults instantiates a new SearchFilter3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *SearchFilter3) GetProperty() SettingProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchFilter3) GetPropertyOk() (*SettingProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchFilter3) SetProperty(v SettingProperty)`

SetProperty sets Property field to given value.


### GetValue

`func (o *SearchFilter3) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchFilter3) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchFilter3) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *SearchFilter3) GetOperator() SearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SearchFilter3) GetOperatorOk() (*SearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SearchFilter3) SetOperator(v SearchOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


