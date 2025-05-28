# CommissionTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to [**NullableCommissionTaskOperationType**](CommissionTaskOperationType.md) | Task Type | [optional] 

## Methods

### NewCommissionTask

`func NewCommissionTask() *CommissionTask`

NewCommissionTask instantiates a new CommissionTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionTaskWithDefaults

`func NewCommissionTaskWithDefaults() *CommissionTask`

NewCommissionTaskWithDefaults instantiates a new CommissionTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *CommissionTask) GetOperation() CommissionTaskOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CommissionTask) GetOperationOk() (*CommissionTaskOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CommissionTask) SetOperation(v CommissionTaskOperationType)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *CommissionTask) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *CommissionTask) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *CommissionTask) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


