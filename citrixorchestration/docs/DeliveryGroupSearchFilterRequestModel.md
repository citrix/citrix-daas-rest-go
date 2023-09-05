# DeliveryGroupSearchFilterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**DeliveryGroupSearchProperty**](DeliveryGroupSearchProperty.md) |  | 
**Value** | Pointer to **string** | Value to match. | [optional] 
**Operator** | [**SearchOperator**](SearchOperator.md) |  | 

## Methods

### NewDeliveryGroupSearchFilterRequestModel

`func NewDeliveryGroupSearchFilterRequestModel(property DeliveryGroupSearchProperty, operator SearchOperator, ) *DeliveryGroupSearchFilterRequestModel`

NewDeliveryGroupSearchFilterRequestModel instantiates a new DeliveryGroupSearchFilterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupSearchFilterRequestModelWithDefaults

`func NewDeliveryGroupSearchFilterRequestModelWithDefaults() *DeliveryGroupSearchFilterRequestModel`

NewDeliveryGroupSearchFilterRequestModelWithDefaults instantiates a new DeliveryGroupSearchFilterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *DeliveryGroupSearchFilterRequestModel) GetProperty() DeliveryGroupSearchProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *DeliveryGroupSearchFilterRequestModel) GetPropertyOk() (*DeliveryGroupSearchProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *DeliveryGroupSearchFilterRequestModel) SetProperty(v DeliveryGroupSearchProperty)`

SetProperty sets Property field to given value.


### GetValue

`func (o *DeliveryGroupSearchFilterRequestModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeliveryGroupSearchFilterRequestModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeliveryGroupSearchFilterRequestModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeliveryGroupSearchFilterRequestModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *DeliveryGroupSearchFilterRequestModel) GetOperator() SearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DeliveryGroupSearchFilterRequestModel) GetOperatorOk() (*SearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DeliveryGroupSearchFilterRequestModel) SetOperator(v SearchOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


