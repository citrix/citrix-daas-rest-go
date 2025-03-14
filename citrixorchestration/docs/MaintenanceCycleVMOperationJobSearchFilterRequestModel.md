# MaintenanceCycleVMOperationJobSearchFilterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**MaintenanceCycleVMOperationJobSearchProperty**](MaintenanceCycleVMOperationJobSearchProperty.md) |  | 
**Value** | Pointer to **NullableString** | Value to match. | [optional] 
**Operator** | [**SearchOperator**](SearchOperator.md) |  | 

## Methods

### NewMaintenanceCycleVMOperationJobSearchFilterRequestModel

`func NewMaintenanceCycleVMOperationJobSearchFilterRequestModel(property MaintenanceCycleVMOperationJobSearchProperty, operator SearchOperator, ) *MaintenanceCycleVMOperationJobSearchFilterRequestModel`

NewMaintenanceCycleVMOperationJobSearchFilterRequestModel instantiates a new MaintenanceCycleVMOperationJobSearchFilterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceCycleVMOperationJobSearchFilterRequestModelWithDefaults

`func NewMaintenanceCycleVMOperationJobSearchFilterRequestModelWithDefaults() *MaintenanceCycleVMOperationJobSearchFilterRequestModel`

NewMaintenanceCycleVMOperationJobSearchFilterRequestModelWithDefaults instantiates a new MaintenanceCycleVMOperationJobSearchFilterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) GetProperty() MaintenanceCycleVMOperationJobSearchProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) GetPropertyOk() (*MaintenanceCycleVMOperationJobSearchProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) SetProperty(v MaintenanceCycleVMOperationJobSearchProperty)`

SetProperty sets Property field to given value.


### GetValue

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetOperator

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) GetOperator() SearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) GetOperatorOk() (*SearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MaintenanceCycleVMOperationJobSearchFilterRequestModel) SetOperator(v SearchOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


