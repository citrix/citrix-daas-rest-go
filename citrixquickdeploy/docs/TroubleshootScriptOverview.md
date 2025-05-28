# TroubleshootScriptOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TroubleshootId** | Pointer to **string** | The Id of the troubleshoot script | [optional] 
**ScriptName** | Pointer to **string** | The name of the troubleshoot script | [optional] 
**State** | Pointer to [**TroubleshootState**](TroubleshootState.md) | The current state of the script | [optional] 
**ErrorMsg** | Pointer to **string** | Any error that occured during execution of the troubleshoot script | [optional] 
**StartedAt** | Pointer to **time.Time** | The datetime when the job started | [optional] 
**EndedAt** | Pointer to **time.Time** | The datetime when the job ended | [optional] 
**MachineName** | Pointer to **string** | The name of the machine that the script runs on | [optional] 
**SasUrl** | Pointer to **string** | The sas url for downloading the output of the script | [optional] 
**TransactionId** | Pointer to **string** | The transaction id | [optional] 

## Methods

### NewTroubleshootScriptOverview

`func NewTroubleshootScriptOverview() *TroubleshootScriptOverview`

NewTroubleshootScriptOverview instantiates a new TroubleshootScriptOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTroubleshootScriptOverviewWithDefaults

`func NewTroubleshootScriptOverviewWithDefaults() *TroubleshootScriptOverview`

NewTroubleshootScriptOverviewWithDefaults instantiates a new TroubleshootScriptOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTroubleshootId

`func (o *TroubleshootScriptOverview) GetTroubleshootId() string`

GetTroubleshootId returns the TroubleshootId field if non-nil, zero value otherwise.

### GetTroubleshootIdOk

`func (o *TroubleshootScriptOverview) GetTroubleshootIdOk() (*string, bool)`

GetTroubleshootIdOk returns a tuple with the TroubleshootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroubleshootId

`func (o *TroubleshootScriptOverview) SetTroubleshootId(v string)`

SetTroubleshootId sets TroubleshootId field to given value.

### HasTroubleshootId

`func (o *TroubleshootScriptOverview) HasTroubleshootId() bool`

HasTroubleshootId returns a boolean if a field has been set.

### GetScriptName

`func (o *TroubleshootScriptOverview) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *TroubleshootScriptOverview) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *TroubleshootScriptOverview) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *TroubleshootScriptOverview) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.

### GetState

`func (o *TroubleshootScriptOverview) GetState() TroubleshootState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TroubleshootScriptOverview) GetStateOk() (*TroubleshootState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TroubleshootScriptOverview) SetState(v TroubleshootState)`

SetState sets State field to given value.

### HasState

`func (o *TroubleshootScriptOverview) HasState() bool`

HasState returns a boolean if a field has been set.

### GetErrorMsg

`func (o *TroubleshootScriptOverview) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *TroubleshootScriptOverview) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *TroubleshootScriptOverview) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *TroubleshootScriptOverview) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetStartedAt

`func (o *TroubleshootScriptOverview) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TroubleshootScriptOverview) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TroubleshootScriptOverview) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TroubleshootScriptOverview) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *TroubleshootScriptOverview) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *TroubleshootScriptOverview) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *TroubleshootScriptOverview) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *TroubleshootScriptOverview) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetMachineName

`func (o *TroubleshootScriptOverview) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *TroubleshootScriptOverview) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *TroubleshootScriptOverview) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *TroubleshootScriptOverview) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetSasUrl

`func (o *TroubleshootScriptOverview) GetSasUrl() string`

GetSasUrl returns the SasUrl field if non-nil, zero value otherwise.

### GetSasUrlOk

`func (o *TroubleshootScriptOverview) GetSasUrlOk() (*string, bool)`

GetSasUrlOk returns a tuple with the SasUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasUrl

`func (o *TroubleshootScriptOverview) SetSasUrl(v string)`

SetSasUrl sets SasUrl field to given value.

### HasSasUrl

`func (o *TroubleshootScriptOverview) HasSasUrl() bool`

HasSasUrl returns a boolean if a field has been set.

### GetTransactionId

`func (o *TroubleshootScriptOverview) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TroubleshootScriptOverview) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TroubleshootScriptOverview) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TroubleshootScriptOverview) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


