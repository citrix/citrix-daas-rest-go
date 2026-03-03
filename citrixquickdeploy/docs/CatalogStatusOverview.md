# CatalogStatusOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**CatalogOverallState**](CatalogOverallState.md) |  | 
**SubState** | Pointer to [**NullableCatalogOverallSubState**](CatalogOverallSubState.md) |  | [optional] 
**Warnings** | Pointer to [**[]CatalogWarning**](CatalogWarning.md) |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**CatalogErrorDetails** | Pointer to **NullableString** |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 
**ExtraInfo** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **NullableInt32** | Estimated total time for the job to finish | [optional] 

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

### SetSubStateNil

`func (o *CatalogStatusOverview) SetSubStateNil(b bool)`

 SetSubStateNil sets the value for SubState to be an explicit nil

### UnsetSubState
`func (o *CatalogStatusOverview) UnsetSubState()`

UnsetSubState ensures that no value is present for SubState, not even an explicit nil
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

### SetWarningsNil

`func (o *CatalogStatusOverview) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *CatalogStatusOverview) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
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

### SetStatusMessageNil

`func (o *CatalogStatusOverview) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *CatalogStatusOverview) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
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

### SetCatalogErrorDetailsNil

`func (o *CatalogStatusOverview) SetCatalogErrorDetailsNil(b bool)`

 SetCatalogErrorDetailsNil sets the value for CatalogErrorDetails to be an explicit nil

### UnsetCatalogErrorDetails
`func (o *CatalogStatusOverview) UnsetCatalogErrorDetails()`

UnsetCatalogErrorDetails ensures that no value is present for CatalogErrorDetails, not even an explicit nil
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

### SetTransactionIdNil

`func (o *CatalogStatusOverview) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *CatalogStatusOverview) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
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

### SetExtraInfoNil

`func (o *CatalogStatusOverview) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *CatalogStatusOverview) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil
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

### SetStartedAtNil

`func (o *CatalogStatusOverview) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *CatalogStatusOverview) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
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

### SetEstimatedTimeInMinuteNil

`func (o *CatalogStatusOverview) SetEstimatedTimeInMinuteNil(b bool)`

 SetEstimatedTimeInMinuteNil sets the value for EstimatedTimeInMinute to be an explicit nil

### UnsetEstimatedTimeInMinute
`func (o *CatalogStatusOverview) UnsetEstimatedTimeInMinute()`

UnsetEstimatedTimeInMinute ensures that no value is present for EstimatedTimeInMinute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


