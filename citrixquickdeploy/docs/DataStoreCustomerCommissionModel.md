# DataStoreCustomerCommissionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommissionState** | Pointer to [**CommissionState**](CommissionState.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**CommissionStateTransitionTimeStamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDataStoreCustomerCommissionModel

`func NewDataStoreCustomerCommissionModel() *DataStoreCustomerCommissionModel`

NewDataStoreCustomerCommissionModel instantiates a new DataStoreCustomerCommissionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreCustomerCommissionModelWithDefaults

`func NewDataStoreCustomerCommissionModelWithDefaults() *DataStoreCustomerCommissionModel`

NewDataStoreCustomerCommissionModelWithDefaults instantiates a new DataStoreCustomerCommissionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommissionState

`func (o *DataStoreCustomerCommissionModel) GetCommissionState() CommissionState`

GetCommissionState returns the CommissionState field if non-nil, zero value otherwise.

### GetCommissionStateOk

`func (o *DataStoreCustomerCommissionModel) GetCommissionStateOk() (*CommissionState, bool)`

GetCommissionStateOk returns a tuple with the CommissionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionState

`func (o *DataStoreCustomerCommissionModel) SetCommissionState(v CommissionState)`

SetCommissionState sets CommissionState field to given value.

### HasCommissionState

`func (o *DataStoreCustomerCommissionModel) HasCommissionState() bool`

HasCommissionState returns a boolean if a field has been set.

### GetError

`func (o *DataStoreCustomerCommissionModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DataStoreCustomerCommissionModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DataStoreCustomerCommissionModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DataStoreCustomerCommissionModel) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTransactionId

`func (o *DataStoreCustomerCommissionModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DataStoreCustomerCommissionModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DataStoreCustomerCommissionModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DataStoreCustomerCommissionModel) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetCommissionStateTransitionTimeStamp

`func (o *DataStoreCustomerCommissionModel) GetCommissionStateTransitionTimeStamp() time.Time`

GetCommissionStateTransitionTimeStamp returns the CommissionStateTransitionTimeStamp field if non-nil, zero value otherwise.

### GetCommissionStateTransitionTimeStampOk

`func (o *DataStoreCustomerCommissionModel) GetCommissionStateTransitionTimeStampOk() (*time.Time, bool)`

GetCommissionStateTransitionTimeStampOk returns a tuple with the CommissionStateTransitionTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionStateTransitionTimeStamp

`func (o *DataStoreCustomerCommissionModel) SetCommissionStateTransitionTimeStamp(v time.Time)`

SetCommissionStateTransitionTimeStamp sets CommissionStateTransitionTimeStamp field to given value.

### HasCommissionStateTransitionTimeStamp

`func (o *DataStoreCustomerCommissionModel) HasCommissionStateTransitionTimeStamp() bool`

HasCommissionStateTransitionTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


