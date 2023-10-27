# BackupRestoreStatusRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**BackupRestoreActions**](BackupRestoreActions.md) |  | 
**ExecutionId** | **string** | Execution Id | 
**Duration** | **int32** | Current/Total duration in seconds | 
**RestApiDuration** | Pointer to **int32** | Current/Total Rest API duration in seconds | [optional] 
**Status** | [**BackupRestoreStatus**](BackupRestoreStatus.md) |  | 
**BackupName** | Pointer to **NullableString** | Backup Name | [optional] 
**BackupSize** | Pointer to **int64** | Backup Size | [optional] 
**ScheduleName** | Pointer to **NullableString** | Schedule Name | [optional] 
**MemberCount** | Pointer to **int32** | Member Count | [optional] 
**StartTime** | **time.Time** | Action Start UTC Time              | 
**CurrentTime** | **time.Time** | Action Current UTC Time | 
**RestoreProgress** | Pointer to [**BackupRestoreRestoreSingleMemberModel**](BackupRestoreRestoreSingleMemberModel.md) |  | [optional] 
**ActivityCount** | Pointer to **int32** | Activity count for stall testing | [optional] 
**CurrentComponent** | Pointer to **NullableString** | Current component being restored | [optional] 
**CurrentMemberName** | Pointer to **NullableString** | Current component member being restored | [optional] 
**ComponentProgress** | Pointer to [**BackupRestorePlaybookActionState**](BackupRestorePlaybookActionState.md) |  | [optional] 
**CheckMode** | Pointer to **bool** | Check Mode | [optional] 
**PreImportCheckMode** | Pointer to **bool** | Pre-Import Check Mode | [optional] 
**AdministratrName** | Pointer to **NullableString** | AdministratorName | [optional] 
**IsBackupRetry** | Pointer to **bool** | Is this backup a retry after a failed backup | [optional] 
**BackupDetails** | Pointer to **map[string]string** | Backup Status Details | [optional] 
**RestoreDetails** | Pointer to [**[]BackupRestoreRestoreSingleMemberModel**](BackupRestoreRestoreSingleMemberModel.md) | Restore Status Details | [optional] 
**SimpleResults** | Pointer to **[]string** | Simple Results (such as Get backed up member names) | [optional] 
**RestoreType** | Pointer to [**BackupRestoreRestoreTypes**](BackupRestoreRestoreTypes.md) |  | [optional] 

## Methods

### NewBackupRestoreStatusRequestModel

`func NewBackupRestoreStatusRequestModel(action BackupRestoreActions, executionId string, duration int32, status BackupRestoreStatus, startTime time.Time, currentTime time.Time, ) *BackupRestoreStatusRequestModel`

NewBackupRestoreStatusRequestModel instantiates a new BackupRestoreStatusRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreStatusRequestModelWithDefaults

`func NewBackupRestoreStatusRequestModelWithDefaults() *BackupRestoreStatusRequestModel`

NewBackupRestoreStatusRequestModelWithDefaults instantiates a new BackupRestoreStatusRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BackupRestoreStatusRequestModel) GetAction() BackupRestoreActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BackupRestoreStatusRequestModel) GetActionOk() (*BackupRestoreActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BackupRestoreStatusRequestModel) SetAction(v BackupRestoreActions)`

SetAction sets Action field to given value.


### GetExecutionId

`func (o *BackupRestoreStatusRequestModel) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BackupRestoreStatusRequestModel) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BackupRestoreStatusRequestModel) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetDuration

`func (o *BackupRestoreStatusRequestModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BackupRestoreStatusRequestModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BackupRestoreStatusRequestModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetRestApiDuration

`func (o *BackupRestoreStatusRequestModel) GetRestApiDuration() int32`

GetRestApiDuration returns the RestApiDuration field if non-nil, zero value otherwise.

### GetRestApiDurationOk

`func (o *BackupRestoreStatusRequestModel) GetRestApiDurationOk() (*int32, bool)`

GetRestApiDurationOk returns a tuple with the RestApiDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestApiDuration

`func (o *BackupRestoreStatusRequestModel) SetRestApiDuration(v int32)`

SetRestApiDuration sets RestApiDuration field to given value.

### HasRestApiDuration

`func (o *BackupRestoreStatusRequestModel) HasRestApiDuration() bool`

HasRestApiDuration returns a boolean if a field has been set.

### GetStatus

`func (o *BackupRestoreStatusRequestModel) GetStatus() BackupRestoreStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupRestoreStatusRequestModel) GetStatusOk() (*BackupRestoreStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupRestoreStatusRequestModel) SetStatus(v BackupRestoreStatus)`

SetStatus sets Status field to given value.


### GetBackupName

`func (o *BackupRestoreStatusRequestModel) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *BackupRestoreStatusRequestModel) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *BackupRestoreStatusRequestModel) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *BackupRestoreStatusRequestModel) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### SetBackupNameNil

`func (o *BackupRestoreStatusRequestModel) SetBackupNameNil(b bool)`

 SetBackupNameNil sets the value for BackupName to be an explicit nil

### UnsetBackupName
`func (o *BackupRestoreStatusRequestModel) UnsetBackupName()`

UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
### GetBackupSize

`func (o *BackupRestoreStatusRequestModel) GetBackupSize() int64`

GetBackupSize returns the BackupSize field if non-nil, zero value otherwise.

### GetBackupSizeOk

`func (o *BackupRestoreStatusRequestModel) GetBackupSizeOk() (*int64, bool)`

GetBackupSizeOk returns a tuple with the BackupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSize

`func (o *BackupRestoreStatusRequestModel) SetBackupSize(v int64)`

SetBackupSize sets BackupSize field to given value.

### HasBackupSize

`func (o *BackupRestoreStatusRequestModel) HasBackupSize() bool`

HasBackupSize returns a boolean if a field has been set.

### GetScheduleName

`func (o *BackupRestoreStatusRequestModel) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *BackupRestoreStatusRequestModel) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *BackupRestoreStatusRequestModel) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *BackupRestoreStatusRequestModel) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### SetScheduleNameNil

`func (o *BackupRestoreStatusRequestModel) SetScheduleNameNil(b bool)`

 SetScheduleNameNil sets the value for ScheduleName to be an explicit nil

### UnsetScheduleName
`func (o *BackupRestoreStatusRequestModel) UnsetScheduleName()`

UnsetScheduleName ensures that no value is present for ScheduleName, not even an explicit nil
### GetMemberCount

`func (o *BackupRestoreStatusRequestModel) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *BackupRestoreStatusRequestModel) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *BackupRestoreStatusRequestModel) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *BackupRestoreStatusRequestModel) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetStartTime

`func (o *BackupRestoreStatusRequestModel) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupRestoreStatusRequestModel) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupRestoreStatusRequestModel) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetCurrentTime

`func (o *BackupRestoreStatusRequestModel) GetCurrentTime() time.Time`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *BackupRestoreStatusRequestModel) GetCurrentTimeOk() (*time.Time, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *BackupRestoreStatusRequestModel) SetCurrentTime(v time.Time)`

SetCurrentTime sets CurrentTime field to given value.


### GetRestoreProgress

`func (o *BackupRestoreStatusRequestModel) GetRestoreProgress() BackupRestoreRestoreSingleMemberModel`

GetRestoreProgress returns the RestoreProgress field if non-nil, zero value otherwise.

### GetRestoreProgressOk

`func (o *BackupRestoreStatusRequestModel) GetRestoreProgressOk() (*BackupRestoreRestoreSingleMemberModel, bool)`

GetRestoreProgressOk returns a tuple with the RestoreProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreProgress

`func (o *BackupRestoreStatusRequestModel) SetRestoreProgress(v BackupRestoreRestoreSingleMemberModel)`

SetRestoreProgress sets RestoreProgress field to given value.

### HasRestoreProgress

`func (o *BackupRestoreStatusRequestModel) HasRestoreProgress() bool`

HasRestoreProgress returns a boolean if a field has been set.

### GetActivityCount

`func (o *BackupRestoreStatusRequestModel) GetActivityCount() int32`

GetActivityCount returns the ActivityCount field if non-nil, zero value otherwise.

### GetActivityCountOk

`func (o *BackupRestoreStatusRequestModel) GetActivityCountOk() (*int32, bool)`

GetActivityCountOk returns a tuple with the ActivityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCount

`func (o *BackupRestoreStatusRequestModel) SetActivityCount(v int32)`

SetActivityCount sets ActivityCount field to given value.

### HasActivityCount

`func (o *BackupRestoreStatusRequestModel) HasActivityCount() bool`

HasActivityCount returns a boolean if a field has been set.

### GetCurrentComponent

`func (o *BackupRestoreStatusRequestModel) GetCurrentComponent() string`

GetCurrentComponent returns the CurrentComponent field if non-nil, zero value otherwise.

### GetCurrentComponentOk

`func (o *BackupRestoreStatusRequestModel) GetCurrentComponentOk() (*string, bool)`

GetCurrentComponentOk returns a tuple with the CurrentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentComponent

`func (o *BackupRestoreStatusRequestModel) SetCurrentComponent(v string)`

SetCurrentComponent sets CurrentComponent field to given value.

### HasCurrentComponent

`func (o *BackupRestoreStatusRequestModel) HasCurrentComponent() bool`

HasCurrentComponent returns a boolean if a field has been set.

### SetCurrentComponentNil

`func (o *BackupRestoreStatusRequestModel) SetCurrentComponentNil(b bool)`

 SetCurrentComponentNil sets the value for CurrentComponent to be an explicit nil

### UnsetCurrentComponent
`func (o *BackupRestoreStatusRequestModel) UnsetCurrentComponent()`

UnsetCurrentComponent ensures that no value is present for CurrentComponent, not even an explicit nil
### GetCurrentMemberName

`func (o *BackupRestoreStatusRequestModel) GetCurrentMemberName() string`

GetCurrentMemberName returns the CurrentMemberName field if non-nil, zero value otherwise.

### GetCurrentMemberNameOk

`func (o *BackupRestoreStatusRequestModel) GetCurrentMemberNameOk() (*string, bool)`

GetCurrentMemberNameOk returns a tuple with the CurrentMemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMemberName

`func (o *BackupRestoreStatusRequestModel) SetCurrentMemberName(v string)`

SetCurrentMemberName sets CurrentMemberName field to given value.

### HasCurrentMemberName

`func (o *BackupRestoreStatusRequestModel) HasCurrentMemberName() bool`

HasCurrentMemberName returns a boolean if a field has been set.

### SetCurrentMemberNameNil

`func (o *BackupRestoreStatusRequestModel) SetCurrentMemberNameNil(b bool)`

 SetCurrentMemberNameNil sets the value for CurrentMemberName to be an explicit nil

### UnsetCurrentMemberName
`func (o *BackupRestoreStatusRequestModel) UnsetCurrentMemberName()`

UnsetCurrentMemberName ensures that no value is present for CurrentMemberName, not even an explicit nil
### GetComponentProgress

`func (o *BackupRestoreStatusRequestModel) GetComponentProgress() BackupRestorePlaybookActionState`

GetComponentProgress returns the ComponentProgress field if non-nil, zero value otherwise.

### GetComponentProgressOk

`func (o *BackupRestoreStatusRequestModel) GetComponentProgressOk() (*BackupRestorePlaybookActionState, bool)`

GetComponentProgressOk returns a tuple with the ComponentProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentProgress

`func (o *BackupRestoreStatusRequestModel) SetComponentProgress(v BackupRestorePlaybookActionState)`

SetComponentProgress sets ComponentProgress field to given value.

### HasComponentProgress

`func (o *BackupRestoreStatusRequestModel) HasComponentProgress() bool`

HasComponentProgress returns a boolean if a field has been set.

### GetCheckMode

`func (o *BackupRestoreStatusRequestModel) GetCheckMode() bool`

GetCheckMode returns the CheckMode field if non-nil, zero value otherwise.

### GetCheckModeOk

`func (o *BackupRestoreStatusRequestModel) GetCheckModeOk() (*bool, bool)`

GetCheckModeOk returns a tuple with the CheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMode

`func (o *BackupRestoreStatusRequestModel) SetCheckMode(v bool)`

SetCheckMode sets CheckMode field to given value.

### HasCheckMode

`func (o *BackupRestoreStatusRequestModel) HasCheckMode() bool`

HasCheckMode returns a boolean if a field has been set.

### GetPreImportCheckMode

`func (o *BackupRestoreStatusRequestModel) GetPreImportCheckMode() bool`

GetPreImportCheckMode returns the PreImportCheckMode field if non-nil, zero value otherwise.

### GetPreImportCheckModeOk

`func (o *BackupRestoreStatusRequestModel) GetPreImportCheckModeOk() (*bool, bool)`

GetPreImportCheckModeOk returns a tuple with the PreImportCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreImportCheckMode

`func (o *BackupRestoreStatusRequestModel) SetPreImportCheckMode(v bool)`

SetPreImportCheckMode sets PreImportCheckMode field to given value.

### HasPreImportCheckMode

`func (o *BackupRestoreStatusRequestModel) HasPreImportCheckMode() bool`

HasPreImportCheckMode returns a boolean if a field has been set.

### GetAdministratrName

`func (o *BackupRestoreStatusRequestModel) GetAdministratrName() string`

GetAdministratrName returns the AdministratrName field if non-nil, zero value otherwise.

### GetAdministratrNameOk

`func (o *BackupRestoreStatusRequestModel) GetAdministratrNameOk() (*string, bool)`

GetAdministratrNameOk returns a tuple with the AdministratrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratrName

`func (o *BackupRestoreStatusRequestModel) SetAdministratrName(v string)`

SetAdministratrName sets AdministratrName field to given value.

### HasAdministratrName

`func (o *BackupRestoreStatusRequestModel) HasAdministratrName() bool`

HasAdministratrName returns a boolean if a field has been set.

### SetAdministratrNameNil

`func (o *BackupRestoreStatusRequestModel) SetAdministratrNameNil(b bool)`

 SetAdministratrNameNil sets the value for AdministratrName to be an explicit nil

### UnsetAdministratrName
`func (o *BackupRestoreStatusRequestModel) UnsetAdministratrName()`

UnsetAdministratrName ensures that no value is present for AdministratrName, not even an explicit nil
### GetIsBackupRetry

`func (o *BackupRestoreStatusRequestModel) GetIsBackupRetry() bool`

GetIsBackupRetry returns the IsBackupRetry field if non-nil, zero value otherwise.

### GetIsBackupRetryOk

`func (o *BackupRestoreStatusRequestModel) GetIsBackupRetryOk() (*bool, bool)`

GetIsBackupRetryOk returns a tuple with the IsBackupRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupRetry

`func (o *BackupRestoreStatusRequestModel) SetIsBackupRetry(v bool)`

SetIsBackupRetry sets IsBackupRetry field to given value.

### HasIsBackupRetry

`func (o *BackupRestoreStatusRequestModel) HasIsBackupRetry() bool`

HasIsBackupRetry returns a boolean if a field has been set.

### GetBackupDetails

`func (o *BackupRestoreStatusRequestModel) GetBackupDetails() map[string]string`

GetBackupDetails returns the BackupDetails field if non-nil, zero value otherwise.

### GetBackupDetailsOk

`func (o *BackupRestoreStatusRequestModel) GetBackupDetailsOk() (*map[string]string, bool)`

GetBackupDetailsOk returns a tuple with the BackupDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDetails

`func (o *BackupRestoreStatusRequestModel) SetBackupDetails(v map[string]string)`

SetBackupDetails sets BackupDetails field to given value.

### HasBackupDetails

`func (o *BackupRestoreStatusRequestModel) HasBackupDetails() bool`

HasBackupDetails returns a boolean if a field has been set.

### SetBackupDetailsNil

`func (o *BackupRestoreStatusRequestModel) SetBackupDetailsNil(b bool)`

 SetBackupDetailsNil sets the value for BackupDetails to be an explicit nil

### UnsetBackupDetails
`func (o *BackupRestoreStatusRequestModel) UnsetBackupDetails()`

UnsetBackupDetails ensures that no value is present for BackupDetails, not even an explicit nil
### GetRestoreDetails

`func (o *BackupRestoreStatusRequestModel) GetRestoreDetails() []BackupRestoreRestoreSingleMemberModel`

GetRestoreDetails returns the RestoreDetails field if non-nil, zero value otherwise.

### GetRestoreDetailsOk

`func (o *BackupRestoreStatusRequestModel) GetRestoreDetailsOk() (*[]BackupRestoreRestoreSingleMemberModel, bool)`

GetRestoreDetailsOk returns a tuple with the RestoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDetails

`func (o *BackupRestoreStatusRequestModel) SetRestoreDetails(v []BackupRestoreRestoreSingleMemberModel)`

SetRestoreDetails sets RestoreDetails field to given value.

### HasRestoreDetails

`func (o *BackupRestoreStatusRequestModel) HasRestoreDetails() bool`

HasRestoreDetails returns a boolean if a field has been set.

### SetRestoreDetailsNil

`func (o *BackupRestoreStatusRequestModel) SetRestoreDetailsNil(b bool)`

 SetRestoreDetailsNil sets the value for RestoreDetails to be an explicit nil

### UnsetRestoreDetails
`func (o *BackupRestoreStatusRequestModel) UnsetRestoreDetails()`

UnsetRestoreDetails ensures that no value is present for RestoreDetails, not even an explicit nil
### GetSimpleResults

`func (o *BackupRestoreStatusRequestModel) GetSimpleResults() []string`

GetSimpleResults returns the SimpleResults field if non-nil, zero value otherwise.

### GetSimpleResultsOk

`func (o *BackupRestoreStatusRequestModel) GetSimpleResultsOk() (*[]string, bool)`

GetSimpleResultsOk returns a tuple with the SimpleResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleResults

`func (o *BackupRestoreStatusRequestModel) SetSimpleResults(v []string)`

SetSimpleResults sets SimpleResults field to given value.

### HasSimpleResults

`func (o *BackupRestoreStatusRequestModel) HasSimpleResults() bool`

HasSimpleResults returns a boolean if a field has been set.

### SetSimpleResultsNil

`func (o *BackupRestoreStatusRequestModel) SetSimpleResultsNil(b bool)`

 SetSimpleResultsNil sets the value for SimpleResults to be an explicit nil

### UnsetSimpleResults
`func (o *BackupRestoreStatusRequestModel) UnsetSimpleResults()`

UnsetSimpleResults ensures that no value is present for SimpleResults, not even an explicit nil
### GetRestoreType

`func (o *BackupRestoreStatusRequestModel) GetRestoreType() BackupRestoreRestoreTypes`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *BackupRestoreStatusRequestModel) GetRestoreTypeOk() (*BackupRestoreRestoreTypes, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *BackupRestoreStatusRequestModel) SetRestoreType(v BackupRestoreRestoreTypes)`

SetRestoreType sets RestoreType field to given value.

### HasRestoreType

`func (o *BackupRestoreStatusRequestModel) HasRestoreType() bool`

HasRestoreType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


