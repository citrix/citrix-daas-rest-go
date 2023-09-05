# BackupRestoreStatusRequestModelRestoreProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | Pointer to **string** | Component Name | [optional] 
**MemberName** | Pointer to **string** | Member Name | [optional] 
**PlaybookActionState** | Pointer to [**BackupRestorePlaybookActionState**](BackupRestorePlaybookActionState.md) |  | [optional] 
**Result** | Pointer to **bool** | Result | [optional] 
**Details** | Pointer to **string** | Details or Error Message | [optional] 
**TransactionId** | Pointer to **string** | Transaction Id | [optional] 
**DateTime** | Pointer to **time.Time** | Date Time stamps | [optional] 
**MemberCount** | Pointer to **int32** | Member Count | [optional] 
**CurrentMemberIndex** | Pointer to **int32** |  | [optional] 
**CheckMode** | Pointer to **bool** | Check Mode | [optional] 

## Methods

### NewBackupRestoreStatusRequestModelRestoreProgress

`func NewBackupRestoreStatusRequestModelRestoreProgress() *BackupRestoreStatusRequestModelRestoreProgress`

NewBackupRestoreStatusRequestModelRestoreProgress instantiates a new BackupRestoreStatusRequestModelRestoreProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreStatusRequestModelRestoreProgressWithDefaults

`func NewBackupRestoreStatusRequestModelRestoreProgressWithDefaults() *BackupRestoreStatusRequestModelRestoreProgress`

NewBackupRestoreStatusRequestModelRestoreProgressWithDefaults instantiates a new BackupRestoreStatusRequestModelRestoreProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetMemberName

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### GetPlaybookActionState

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetPlaybookActionState() BackupRestorePlaybookActionState`

GetPlaybookActionState returns the PlaybookActionState field if non-nil, zero value otherwise.

### GetPlaybookActionStateOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetPlaybookActionStateOk() (*BackupRestorePlaybookActionState, bool)`

GetPlaybookActionStateOk returns a tuple with the PlaybookActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookActionState

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetPlaybookActionState(v BackupRestorePlaybookActionState)`

SetPlaybookActionState sets PlaybookActionState field to given value.

### HasPlaybookActionState

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasPlaybookActionState() bool`

HasPlaybookActionState returns a boolean if a field has been set.

### GetResult

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetResult(v bool)`

SetResult sets Result field to given value.

### HasResult

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDetails

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTransactionId

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetDateTime

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetMemberCount

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetCurrentMemberIndex

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetCurrentMemberIndex() int32`

GetCurrentMemberIndex returns the CurrentMemberIndex field if non-nil, zero value otherwise.

### GetCurrentMemberIndexOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetCurrentMemberIndexOk() (*int32, bool)`

GetCurrentMemberIndexOk returns a tuple with the CurrentMemberIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMemberIndex

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetCurrentMemberIndex(v int32)`

SetCurrentMemberIndex sets CurrentMemberIndex field to given value.

### HasCurrentMemberIndex

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasCurrentMemberIndex() bool`

HasCurrentMemberIndex returns a boolean if a field has been set.

### GetCheckMode

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetCheckMode() bool`

GetCheckMode returns the CheckMode field if non-nil, zero value otherwise.

### GetCheckModeOk

`func (o *BackupRestoreStatusRequestModelRestoreProgress) GetCheckModeOk() (*bool, bool)`

GetCheckModeOk returns a tuple with the CheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMode

`func (o *BackupRestoreStatusRequestModelRestoreProgress) SetCheckMode(v bool)`

SetCheckMode sets CheckMode field to given value.

### HasCheckMode

`func (o *BackupRestoreStatusRequestModelRestoreProgress) HasCheckMode() bool`

HasCheckMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


