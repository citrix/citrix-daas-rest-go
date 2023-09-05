# AppSearchFilterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**AppSearchProperty**](AppSearchProperty.md) |  | 
**Value** | Pointer to **string** | Value to match. | [optional] 
**Operator** | [**SearchOperator**](SearchOperator.md) |  | 

## Methods

### NewAppSearchFilterRequestModel

`func NewAppSearchFilterRequestModel(property AppSearchProperty, operator SearchOperator, ) *AppSearchFilterRequestModel`

NewAppSearchFilterRequestModel instantiates a new AppSearchFilterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSearchFilterRequestModelWithDefaults

`func NewAppSearchFilterRequestModelWithDefaults() *AppSearchFilterRequestModel`

NewAppSearchFilterRequestModelWithDefaults instantiates a new AppSearchFilterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *AppSearchFilterRequestModel) GetProperty() AppSearchProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AppSearchFilterRequestModel) GetPropertyOk() (*AppSearchProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AppSearchFilterRequestModel) SetProperty(v AppSearchProperty)`

SetProperty sets Property field to given value.


### GetValue

`func (o *AppSearchFilterRequestModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppSearchFilterRequestModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppSearchFilterRequestModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AppSearchFilterRequestModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *AppSearchFilterRequestModel) GetOperator() SearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AppSearchFilterRequestModel) GetOperatorOk() (*SearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AppSearchFilterRequestModel) SetOperator(v SearchOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


