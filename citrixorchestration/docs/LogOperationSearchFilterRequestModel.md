# LogOperationSearchFilterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**LogOperationSearchProperty**](LogOperationSearchProperty.md) |  | 
**Value** | **string** | Value to match. | 
**Operator** | [**SearchOperator**](SearchOperator.md) |  | 

## Methods

### NewLogOperationSearchFilterRequestModel

`func NewLogOperationSearchFilterRequestModel(property LogOperationSearchProperty, value string, operator SearchOperator, ) *LogOperationSearchFilterRequestModel`

NewLogOperationSearchFilterRequestModel instantiates a new LogOperationSearchFilterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogOperationSearchFilterRequestModelWithDefaults

`func NewLogOperationSearchFilterRequestModelWithDefaults() *LogOperationSearchFilterRequestModel`

NewLogOperationSearchFilterRequestModelWithDefaults instantiates a new LogOperationSearchFilterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *LogOperationSearchFilterRequestModel) GetProperty() LogOperationSearchProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *LogOperationSearchFilterRequestModel) GetPropertyOk() (*LogOperationSearchProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *LogOperationSearchFilterRequestModel) SetProperty(v LogOperationSearchProperty)`

SetProperty sets Property field to given value.


### GetValue

`func (o *LogOperationSearchFilterRequestModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LogOperationSearchFilterRequestModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LogOperationSearchFilterRequestModel) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *LogOperationSearchFilterRequestModel) GetOperator() SearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LogOperationSearchFilterRequestModel) GetOperatorOk() (*SearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LogOperationSearchFilterRequestModel) SetOperator(v SearchOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


