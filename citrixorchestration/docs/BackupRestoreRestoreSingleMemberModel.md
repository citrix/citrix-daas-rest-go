# BackupRestoreRestoreSingleMemberModel

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

### NewBackupRestoreRestoreSingleMemberModel

`func NewBackupRestoreRestoreSingleMemberModel() *BackupRestoreRestoreSingleMemberModel`

NewBackupRestoreRestoreSingleMemberModel instantiates a new BackupRestoreRestoreSingleMemberModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreRestoreSingleMemberModelWithDefaults

`func NewBackupRestoreRestoreSingleMemberModelWithDefaults() *BackupRestoreRestoreSingleMemberModel`

NewBackupRestoreRestoreSingleMemberModelWithDefaults instantiates a new BackupRestoreRestoreSingleMemberModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *BackupRestoreRestoreSingleMemberModel) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *BackupRestoreRestoreSingleMemberModel) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *BackupRestoreRestoreSingleMemberModel) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetMemberName

`func (o *BackupRestoreRestoreSingleMemberModel) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *BackupRestoreRestoreSingleMemberModel) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *BackupRestoreRestoreSingleMemberModel) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### GetPlaybookActionState

`func (o *BackupRestoreRestoreSingleMemberModel) GetPlaybookActionState() BackupRestorePlaybookActionState`

GetPlaybookActionState returns the PlaybookActionState field if non-nil, zero value otherwise.

### GetPlaybookActionStateOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetPlaybookActionStateOk() (*BackupRestorePlaybookActionState, bool)`

GetPlaybookActionStateOk returns a tuple with the PlaybookActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookActionState

`func (o *BackupRestoreRestoreSingleMemberModel) SetPlaybookActionState(v BackupRestorePlaybookActionState)`

SetPlaybookActionState sets PlaybookActionState field to given value.

### HasPlaybookActionState

`func (o *BackupRestoreRestoreSingleMemberModel) HasPlaybookActionState() bool`

HasPlaybookActionState returns a boolean if a field has been set.

### GetResult

`func (o *BackupRestoreRestoreSingleMemberModel) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BackupRestoreRestoreSingleMemberModel) SetResult(v bool)`

SetResult sets Result field to given value.

### HasResult

`func (o *BackupRestoreRestoreSingleMemberModel) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDetails

`func (o *BackupRestoreRestoreSingleMemberModel) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BackupRestoreRestoreSingleMemberModel) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BackupRestoreRestoreSingleMemberModel) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTransactionId

`func (o *BackupRestoreRestoreSingleMemberModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BackupRestoreRestoreSingleMemberModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BackupRestoreRestoreSingleMemberModel) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetDateTime

`func (o *BackupRestoreRestoreSingleMemberModel) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *BackupRestoreRestoreSingleMemberModel) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *BackupRestoreRestoreSingleMemberModel) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetMemberCount

`func (o *BackupRestoreRestoreSingleMemberModel) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *BackupRestoreRestoreSingleMemberModel) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *BackupRestoreRestoreSingleMemberModel) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetCurrentMemberIndex

`func (o *BackupRestoreRestoreSingleMemberModel) GetCurrentMemberIndex() int32`

GetCurrentMemberIndex returns the CurrentMemberIndex field if non-nil, zero value otherwise.

### GetCurrentMemberIndexOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetCurrentMemberIndexOk() (*int32, bool)`

GetCurrentMemberIndexOk returns a tuple with the CurrentMemberIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMemberIndex

`func (o *BackupRestoreRestoreSingleMemberModel) SetCurrentMemberIndex(v int32)`

SetCurrentMemberIndex sets CurrentMemberIndex field to given value.

### HasCurrentMemberIndex

`func (o *BackupRestoreRestoreSingleMemberModel) HasCurrentMemberIndex() bool`

HasCurrentMemberIndex returns a boolean if a field has been set.

### GetCheckMode

`func (o *BackupRestoreRestoreSingleMemberModel) GetCheckMode() bool`

GetCheckMode returns the CheckMode field if non-nil, zero value otherwise.

### GetCheckModeOk

`func (o *BackupRestoreRestoreSingleMemberModel) GetCheckModeOk() (*bool, bool)`

GetCheckModeOk returns a tuple with the CheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMode

`func (o *BackupRestoreRestoreSingleMemberModel) SetCheckMode(v bool)`

SetCheckMode sets CheckMode field to given value.

### HasCheckMode

`func (o *BackupRestoreRestoreSingleMemberModel) HasCheckMode() bool`

HasCheckMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


