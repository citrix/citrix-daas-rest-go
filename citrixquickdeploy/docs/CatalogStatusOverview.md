# CatalogStatusOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**CatalogOverallState**](CatalogOverallState.md) |  | 
**SubState** | Pointer to [**CatalogOverallSubState**](CatalogOverallSubState.md) |  | [optional] 
**Warnings** | Pointer to [**[]CatalogWarning**](CatalogWarning.md) |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**CatalogErrorDetails** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**ExtraInfo** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **int32** | Estimated total time for the job to finish | [optional] 

## Methods

### NewCatalogStatusOverview

`func NewCatalogStatusOverview(state CatalogOverallState, ) *CatalogStatusOverview`

NewCatalogStatusOverview instantiates a new CatalogStatusOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogStatusOverviewWithDefaults

`func NewCatalogStatusOverviewWithDefaults() *CatalogStatusOverview`

NewCatalogStatusOverviewWithDefaults instantiates a new CatalogStatusOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CatalogStatusOverview) GetState() CatalogOverallState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CatalogStatusOverview) GetStateOk() (*CatalogOverallState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CatalogStatusOverview) SetState(v CatalogOverallState)`

SetState sets State field to given value.


### GetSubState

`func (o *CatalogStatusOverview) GetSubState() CatalogOverallSubState`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *CatalogStatusOverview) GetSubStateOk() (*CatalogOverallSubState, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *CatalogStatusOverview) SetSubState(v CatalogOverallSubState)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *CatalogStatusOverview) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### GetWarnings

`func (o *CatalogStatusOverview) GetWarnings() []CatalogWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CatalogStatusOverview) GetWarningsOk() (*[]CatalogWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CatalogStatusOverview) SetWarnings(v []CatalogWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CatalogStatusOverview) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetStatusMessage

`func (o *CatalogStatusOverview) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *CatalogStatusOverview) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *CatalogStatusOverview) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *CatalogStatusOverview) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetCatalogErrorDetails

`func (o *CatalogStatusOverview) GetCatalogErrorDetails() string`

GetCatalogErrorDetails returns the CatalogErrorDetails field if non-nil, zero value otherwise.

### GetCatalogErrorDetailsOk

`func (o *CatalogStatusOverview) GetCatalogErrorDetailsOk() (*string, bool)`

GetCatalogErrorDetailsOk returns a tuple with the CatalogErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogErrorDetails

`func (o *CatalogStatusOverview) SetCatalogErrorDetails(v string)`

SetCatalogErrorDetails sets CatalogErrorDetails field to given value.

### HasCatalogErrorDetails

`func (o *CatalogStatusOverview) HasCatalogErrorDetails() bool`

HasCatalogErrorDetails returns a boolean if a field has been set.

### GetTransactionId

`func (o *CatalogStatusOverview) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CatalogStatusOverview) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CatalogStatusOverview) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CatalogStatusOverview) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetExtraInfo

`func (o *CatalogStatusOverview) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *CatalogStatusOverview) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *CatalogStatusOverview) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *CatalogStatusOverview) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetStartedAt

`func (o *CatalogStatusOverview) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CatalogStatusOverview) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CatalogStatusOverview) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CatalogStatusOverview) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEstimatedTimeInMinute

`func (o *CatalogStatusOverview) GetEstimatedTimeInMinute() int32`

GetEstimatedTimeInMinute returns the EstimatedTimeInMinute field if non-nil, zero value otherwise.

### GetEstimatedTimeInMinuteOk

`func (o *CatalogStatusOverview) GetEstimatedTimeInMinuteOk() (*int32, bool)`

GetEstimatedTimeInMinuteOk returns a tuple with the EstimatedTimeInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeInMinute

`func (o *CatalogStatusOverview) SetEstimatedTimeInMinute(v int32)`

SetEstimatedTimeInMinute sets EstimatedTimeInMinute field to given value.

### HasEstimatedTimeInMinute

`func (o *CatalogStatusOverview) HasEstimatedTimeInMinute() bool`

HasEstimatedTimeInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


