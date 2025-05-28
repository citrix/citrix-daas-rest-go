# CapacityJobOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**CatalogJobState**](CatalogJobState.md) | Current state of the Capacity job | [optional] 
**SubState** | Pointer to [**CatalogJobSubState**](CatalogJobSubState.md) | Current sub state of the Capacity job | [optional] 
**Status** | Pointer to **string** | Status message related to the job | [optional] 
**Error** | Pointer to **string** | Error that occured in job processing | [optional] 
**TransactionId** | Pointer to **string** | Transaction Id for the CapacityUpdate | [optional] 

## Methods

### NewCapacityJobOverview

`func NewCapacityJobOverview() *CapacityJobOverview`

NewCapacityJobOverview instantiates a new CapacityJobOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityJobOverviewWithDefaults

`func NewCapacityJobOverviewWithDefaults() *CapacityJobOverview`

NewCapacityJobOverviewWithDefaults instantiates a new CapacityJobOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CapacityJobOverview) GetState() CatalogJobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CapacityJobOverview) GetStateOk() (*CatalogJobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CapacityJobOverview) SetState(v CatalogJobState)`

SetState sets State field to given value.

### HasState

`func (o *CapacityJobOverview) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubState

`func (o *CapacityJobOverview) GetSubState() CatalogJobSubState`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *CapacityJobOverview) GetSubStateOk() (*CatalogJobSubState, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *CapacityJobOverview) SetSubState(v CatalogJobSubState)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *CapacityJobOverview) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### GetStatus

`func (o *CapacityJobOverview) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CapacityJobOverview) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CapacityJobOverview) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CapacityJobOverview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *CapacityJobOverview) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CapacityJobOverview) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CapacityJobOverview) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CapacityJobOverview) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTransactionId

`func (o *CapacityJobOverview) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CapacityJobOverview) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CapacityJobOverview) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CapacityJobOverview) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


