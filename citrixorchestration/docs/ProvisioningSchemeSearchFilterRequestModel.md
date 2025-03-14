# ProvisioningSchemeSearchFilterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**ProvisioningSchemeSearchProperty**](ProvisioningSchemeSearchProperty.md) |  | 
**Value** | Pointer to **NullableString** | Value to match. | [optional] 
**Operator** | [**SearchOperator**](SearchOperator.md) |  | 

## Methods

### NewProvisioningSchemeSearchFilterRequestModel

`func NewProvisioningSchemeSearchFilterRequestModel(property ProvisioningSchemeSearchProperty, operator SearchOperator, ) *ProvisioningSchemeSearchFilterRequestModel`

NewProvisioningSchemeSearchFilterRequestModel instantiates a new ProvisioningSchemeSearchFilterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeSearchFilterRequestModelWithDefaults

`func NewProvisioningSchemeSearchFilterRequestModelWithDefaults() *ProvisioningSchemeSearchFilterRequestModel`

NewProvisioningSchemeSearchFilterRequestModelWithDefaults instantiates a new ProvisioningSchemeSearchFilterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ProvisioningSchemeSearchFilterRequestModel) GetProperty() ProvisioningSchemeSearchProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ProvisioningSchemeSearchFilterRequestModel) GetPropertyOk() (*ProvisioningSchemeSearchProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ProvisioningSchemeSearchFilterRequestModel) SetProperty(v ProvisioningSchemeSearchProperty)`

SetProperty sets Property field to given value.


### GetValue

`func (o *ProvisioningSchemeSearchFilterRequestModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProvisioningSchemeSearchFilterRequestModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProvisioningSchemeSearchFilterRequestModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProvisioningSchemeSearchFilterRequestModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProvisioningSchemeSearchFilterRequestModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProvisioningSchemeSearchFilterRequestModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetOperator

`func (o *ProvisioningSchemeSearchFilterRequestModel) GetOperator() SearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ProvisioningSchemeSearchFilterRequestModel) GetOperatorOk() (*SearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ProvisioningSchemeSearchFilterRequestModel) SetOperator(v SearchOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


