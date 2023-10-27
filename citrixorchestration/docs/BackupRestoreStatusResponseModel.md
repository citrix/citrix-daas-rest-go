# BackupRestoreStatusResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**BackupRestoreActions**](BackupRestoreActions.md) |  | 
**ExecutionId** | **string** | Execution ID of operation | 
**Duration** | **int32** | Current/Total duration in seconds | 
**RestApiDuration** | **int32** | Current/Total Rest API duration in seconds | 
**Status** | [**BackupRestoreStatus**](BackupRestoreStatus.md) |  | 
**BackupName** | Pointer to **NullableString** | Backup name being created by backup operation or backup name being restored from | [optional] 
**ScheduleName** | Pointer to **NullableString** | Schedule Name | [optional] 
**Notes** | Pointer to **NullableString** | Backup name being created by backup operation or backup name being restored from | [optional] 
**ContinueAction** | Pointer to **bool** | Continue Action: True &#x3D; Continue Action, False &#x3D; Stop action | [optional] 
**StartTime** | Pointer to **time.Time** | Action Start UTC Time              | [optional] 
**CurrentTime** | Pointer to **time.Time** | Action Current UTC Time | [optional] 
**ActivityCount** | Pointer to **int32** | Activity count for stall testing | [optional] 
**CurrentComponent** | Pointer to [**BckRstrAutoConfigComponents**](BckRstrAutoConfigComponents.md) |  | [optional] 
**CurrentMemberName** | Pointer to **NullableString** | Current component member being restored | [optional] 
**MemberCount** | Pointer to **int32** | Member Count | [optional] 
**ComponentProgress** | Pointer to [**BackupRestorePlaybookActionState**](BackupRestorePlaybookActionState.md) |  | [optional] 
**CheckMode** | Pointer to **bool** | Check Mode restore from backup | [optional] 
**PreImportCheckMode** | Pointer to **bool** | Pre-Import Check Mode | [optional] 
**Filters** | Pointer to **NullableString** | Filters used by restore from backup | [optional] 
**AdministratrName** | Pointer to **NullableString** | AdministratorName | [optional] 
**IsBackupRetry** | Pointer to **bool** | Is this a backup retry after a failed backup | [optional] 
**BackupDetails** | Pointer to **map[string]string** | Backup Status Details | [optional] 
**RestoreDetails** | Pointer to [**[]BackupRestoreRestoreSingleMemberModel**](BackupRestoreRestoreSingleMemberModel.md) | Restore Status Details | [optional] 
**SimpleResults** | Pointer to **[]string** | Simple Results (such as Get backed up member names) | [optional] 
**RestoreType** | Pointer to [**BackupRestoreRestoreTypes**](BackupRestoreRestoreTypes.md) |  | [optional] 
**Fixups** | Pointer to [**[]BackupRestoreFixupModel**](BackupRestoreFixupModel.md) | Fixups  | [optional] 

## Methods

### NewBackupRestoreStatusResponseModel

`func NewBackupRestoreStatusResponseModel(action BackupRestoreActions, executionId string, duration int32, restApiDuration int32, status BackupRestoreStatus, ) *BackupRestoreStatusResponseModel`

NewBackupRestoreStatusResponseModel instantiates a new BackupRestoreStatusResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreStatusResponseModelWithDefaults

`func NewBackupRestoreStatusResponseModelWithDefaults() *BackupRestoreStatusResponseModel`

NewBackupRestoreStatusResponseModelWithDefaults instantiates a new BackupRestoreStatusResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BackupRestoreStatusResponseModel) GetAction() BackupRestoreActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BackupRestoreStatusResponseModel) GetActionOk() (*BackupRestoreActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BackupRestoreStatusResponseModel) SetAction(v BackupRestoreActions)`

SetAction sets Action field to given value.


### GetExecutionId

`func (o *BackupRestoreStatusResponseModel) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BackupRestoreStatusResponseModel) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BackupRestoreStatusResponseModel) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetDuration

`func (o *BackupRestoreStatusResponseModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BackupRestoreStatusResponseModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BackupRestoreStatusResponseModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetRestApiDuration

`func (o *BackupRestoreStatusResponseModel) GetRestApiDuration() int32`

GetRestApiDuration returns the RestApiDuration field if non-nil, zero value otherwise.

### GetRestApiDurationOk

`func (o *BackupRestoreStatusResponseModel) GetRestApiDurationOk() (*int32, bool)`

GetRestApiDurationOk returns a tuple with the RestApiDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestApiDuration

`func (o *BackupRestoreStatusResponseModel) SetRestApiDuration(v int32)`

SetRestApiDuration sets RestApiDuration field to given value.


### GetStatus

`func (o *BackupRestoreStatusResponseModel) GetStatus() BackupRestoreStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupRestoreStatusResponseModel) GetStatusOk() (*BackupRestoreStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupRestoreStatusResponseModel) SetStatus(v BackupRestoreStatus)`

SetStatus sets Status field to given value.


### GetBackupName

`func (o *BackupRestoreStatusResponseModel) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *BackupRestoreStatusResponseModel) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *BackupRestoreStatusResponseModel) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *BackupRestoreStatusResponseModel) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### SetBackupNameNil

`func (o *BackupRestoreStatusResponseModel) SetBackupNameNil(b bool)`

 SetBackupNameNil sets the value for BackupName to be an explicit nil

### UnsetBackupName
`func (o *BackupRestoreStatusResponseModel) UnsetBackupName()`

UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
### GetScheduleName

`func (o *BackupRestoreStatusResponseModel) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *BackupRestoreStatusResponseModel) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *BackupRestoreStatusResponseModel) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *BackupRestoreStatusResponseModel) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### SetScheduleNameNil

`func (o *BackupRestoreStatusResponseModel) SetScheduleNameNil(b bool)`

 SetScheduleNameNil sets the value for ScheduleName to be an explicit nil

### UnsetScheduleName
`func (o *BackupRestoreStatusResponseModel) UnsetScheduleName()`

UnsetScheduleName ensures that no value is present for ScheduleName, not even an explicit nil
### GetNotes

`func (o *BackupRestoreStatusResponseModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BackupRestoreStatusResponseModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BackupRestoreStatusResponseModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BackupRestoreStatusResponseModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *BackupRestoreStatusResponseModel) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *BackupRestoreStatusResponseModel) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetContinueAction

`func (o *BackupRestoreStatusResponseModel) GetContinueAction() bool`

GetContinueAction returns the ContinueAction field if non-nil, zero value otherwise.

### GetContinueActionOk

`func (o *BackupRestoreStatusResponseModel) GetContinueActionOk() (*bool, bool)`

GetContinueActionOk returns a tuple with the ContinueAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueAction

`func (o *BackupRestoreStatusResponseModel) SetContinueAction(v bool)`

SetContinueAction sets ContinueAction field to given value.

### HasContinueAction

`func (o *BackupRestoreStatusResponseModel) HasContinueAction() bool`

HasContinueAction returns a boolean if a field has been set.

### GetStartTime

`func (o *BackupRestoreStatusResponseModel) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupRestoreStatusResponseModel) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupRestoreStatusResponseModel) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupRestoreStatusResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCurrentTime

`func (o *BackupRestoreStatusResponseModel) GetCurrentTime() time.Time`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *BackupRestoreStatusResponseModel) GetCurrentTimeOk() (*time.Time, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *BackupRestoreStatusResponseModel) SetCurrentTime(v time.Time)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *BackupRestoreStatusResponseModel) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetActivityCount

`func (o *BackupRestoreStatusResponseModel) GetActivityCount() int32`

GetActivityCount returns the ActivityCount field if non-nil, zero value otherwise.

### GetActivityCountOk

`func (o *BackupRestoreStatusResponseModel) GetActivityCountOk() (*int32, bool)`

GetActivityCountOk returns a tuple with the ActivityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCount

`func (o *BackupRestoreStatusResponseModel) SetActivityCount(v int32)`

SetActivityCount sets ActivityCount field to given value.

### HasActivityCount

`func (o *BackupRestoreStatusResponseModel) HasActivityCount() bool`

HasActivityCount returns a boolean if a field has been set.

### GetCurrentComponent

`func (o *BackupRestoreStatusResponseModel) GetCurrentComponent() BckRstrAutoConfigComponents`

GetCurrentComponent returns the CurrentComponent field if non-nil, zero value otherwise.

### GetCurrentComponentOk

`func (o *BackupRestoreStatusResponseModel) GetCurrentComponentOk() (*BckRstrAutoConfigComponents, bool)`

GetCurrentComponentOk returns a tuple with the CurrentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentComponent

`func (o *BackupRestoreStatusResponseModel) SetCurrentComponent(v BckRstrAutoConfigComponents)`

SetCurrentComponent sets CurrentComponent field to given value.

### HasCurrentComponent

`func (o *BackupRestoreStatusResponseModel) HasCurrentComponent() bool`

HasCurrentComponent returns a boolean if a field has been set.

### GetCurrentMemberName

`func (o *BackupRestoreStatusResponseModel) GetCurrentMemberName() string`

GetCurrentMemberName returns the CurrentMemberName field if non-nil, zero value otherwise.

### GetCurrentMemberNameOk

`func (o *BackupRestoreStatusResponseModel) GetCurrentMemberNameOk() (*string, bool)`

GetCurrentMemberNameOk returns a tuple with the CurrentMemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMemberName

`func (o *BackupRestoreStatusResponseModel) SetCurrentMemberName(v string)`

SetCurrentMemberName sets CurrentMemberName field to given value.

### HasCurrentMemberName

`func (o *BackupRestoreStatusResponseModel) HasCurrentMemberName() bool`

HasCurrentMemberName returns a boolean if a field has been set.

### SetCurrentMemberNameNil

`func (o *BackupRestoreStatusResponseModel) SetCurrentMemberNameNil(b bool)`

 SetCurrentMemberNameNil sets the value for CurrentMemberName to be an explicit nil

### UnsetCurrentMemberName
`func (o *BackupRestoreStatusResponseModel) UnsetCurrentMemberName()`

UnsetCurrentMemberName ensures that no value is present for CurrentMemberName, not even an explicit nil
### GetMemberCount

`func (o *BackupRestoreStatusResponseModel) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *BackupRestoreStatusResponseModel) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *BackupRestoreStatusResponseModel) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *BackupRestoreStatusResponseModel) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetComponentProgress

`func (o *BackupRestoreStatusResponseModel) GetComponentProgress() BackupRestorePlaybookActionState`

GetComponentProgress returns the ComponentProgress field if non-nil, zero value otherwise.

### GetComponentProgressOk

`func (o *BackupRestoreStatusResponseModel) GetComponentProgressOk() (*BackupRestorePlaybookActionState, bool)`

GetComponentProgressOk returns a tuple with the ComponentProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentProgress

`func (o *BackupRestoreStatusResponseModel) SetComponentProgress(v BackupRestorePlaybookActionState)`

SetComponentProgress sets ComponentProgress field to given value.

### HasComponentProgress

`func (o *BackupRestoreStatusResponseModel) HasComponentProgress() bool`

HasComponentProgress returns a boolean if a field has been set.

### GetCheckMode

`func (o *BackupRestoreStatusResponseModel) GetCheckMode() bool`

GetCheckMode returns the CheckMode field if non-nil, zero value otherwise.

### GetCheckModeOk

`func (o *BackupRestoreStatusResponseModel) GetCheckModeOk() (*bool, bool)`

GetCheckModeOk returns a tuple with the CheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMode

`func (o *BackupRestoreStatusResponseModel) SetCheckMode(v bool)`

SetCheckMode sets CheckMode field to given value.

### HasCheckMode

`func (o *BackupRestoreStatusResponseModel) HasCheckMode() bool`

HasCheckMode returns a boolean if a field has been set.

### GetPreImportCheckMode

`func (o *BackupRestoreStatusResponseModel) GetPreImportCheckMode() bool`

GetPreImportCheckMode returns the PreImportCheckMode field if non-nil, zero value otherwise.

### GetPreImportCheckModeOk

`func (o *BackupRestoreStatusResponseModel) GetPreImportCheckModeOk() (*bool, bool)`

GetPreImportCheckModeOk returns a tuple with the PreImportCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreImportCheckMode

`func (o *BackupRestoreStatusResponseModel) SetPreImportCheckMode(v bool)`

SetPreImportCheckMode sets PreImportCheckMode field to given value.

### HasPreImportCheckMode

`func (o *BackupRestoreStatusResponseModel) HasPreImportCheckMode() bool`

HasPreImportCheckMode returns a boolean if a field has been set.

### GetFilters

`func (o *BackupRestoreStatusResponseModel) GetFilters() string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BackupRestoreStatusResponseModel) GetFiltersOk() (*string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BackupRestoreStatusResponseModel) SetFilters(v string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *BackupRestoreStatusResponseModel) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *BackupRestoreStatusResponseModel) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *BackupRestoreStatusResponseModel) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetAdministratrName

`func (o *BackupRestoreStatusResponseModel) GetAdministratrName() string`

GetAdministratrName returns the AdministratrName field if non-nil, zero value otherwise.

### GetAdministratrNameOk

`func (o *BackupRestoreStatusResponseModel) GetAdministratrNameOk() (*string, bool)`

GetAdministratrNameOk returns a tuple with the AdministratrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratrName

`func (o *BackupRestoreStatusResponseModel) SetAdministratrName(v string)`

SetAdministratrName sets AdministratrName field to given value.

### HasAdministratrName

`func (o *BackupRestoreStatusResponseModel) HasAdministratrName() bool`

HasAdministratrName returns a boolean if a field has been set.

### SetAdministratrNameNil

`func (o *BackupRestoreStatusResponseModel) SetAdministratrNameNil(b bool)`

 SetAdministratrNameNil sets the value for AdministratrName to be an explicit nil

### UnsetAdministratrName
`func (o *BackupRestoreStatusResponseModel) UnsetAdministratrName()`

UnsetAdministratrName ensures that no value is present for AdministratrName, not even an explicit nil
### GetIsBackupRetry

`func (o *BackupRestoreStatusResponseModel) GetIsBackupRetry() bool`

GetIsBackupRetry returns the IsBackupRetry field if non-nil, zero value otherwise.

### GetIsBackupRetryOk

`func (o *BackupRestoreStatusResponseModel) GetIsBackupRetryOk() (*bool, bool)`

GetIsBackupRetryOk returns a tuple with the IsBackupRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupRetry

`func (o *BackupRestoreStatusResponseModel) SetIsBackupRetry(v bool)`

SetIsBackupRetry sets IsBackupRetry field to given value.

### HasIsBackupRetry

`func (o *BackupRestoreStatusResponseModel) HasIsBackupRetry() bool`

HasIsBackupRetry returns a boolean if a field has been set.

### GetBackupDetails

`func (o *BackupRestoreStatusResponseModel) GetBackupDetails() map[string]string`

GetBackupDetails returns the BackupDetails field if non-nil, zero value otherwise.

### GetBackupDetailsOk

`func (o *BackupRestoreStatusResponseModel) GetBackupDetailsOk() (*map[string]string, bool)`

GetBackupDetailsOk returns a tuple with the BackupDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDetails

`func (o *BackupRestoreStatusResponseModel) SetBackupDetails(v map[string]string)`

SetBackupDetails sets BackupDetails field to given value.

### HasBackupDetails

`func (o *BackupRestoreStatusResponseModel) HasBackupDetails() bool`

HasBackupDetails returns a boolean if a field has been set.

### SetBackupDetailsNil

`func (o *BackupRestoreStatusResponseModel) SetBackupDetailsNil(b bool)`

 SetBackupDetailsNil sets the value for BackupDetails to be an explicit nil

### UnsetBackupDetails
`func (o *BackupRestoreStatusResponseModel) UnsetBackupDetails()`

UnsetBackupDetails ensures that no value is present for BackupDetails, not even an explicit nil
### GetRestoreDetails

`func (o *BackupRestoreStatusResponseModel) GetRestoreDetails() []BackupRestoreRestoreSingleMemberModel`

GetRestoreDetails returns the RestoreDetails field if non-nil, zero value otherwise.

### GetRestoreDetailsOk

`func (o *BackupRestoreStatusResponseModel) GetRestoreDetailsOk() (*[]BackupRestoreRestoreSingleMemberModel, bool)`

GetRestoreDetailsOk returns a tuple with the RestoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDetails

`func (o *BackupRestoreStatusResponseModel) SetRestoreDetails(v []BackupRestoreRestoreSingleMemberModel)`

SetRestoreDetails sets RestoreDetails field to given value.

### HasRestoreDetails

`func (o *BackupRestoreStatusResponseModel) HasRestoreDetails() bool`

HasRestoreDetails returns a boolean if a field has been set.

### SetRestoreDetailsNil

`func (o *BackupRestoreStatusResponseModel) SetRestoreDetailsNil(b bool)`

 SetRestoreDetailsNil sets the value for RestoreDetails to be an explicit nil

### UnsetRestoreDetails
`func (o *BackupRestoreStatusResponseModel) UnsetRestoreDetails()`

UnsetRestoreDetails ensures that no value is present for RestoreDetails, not even an explicit nil
### GetSimpleResults

`func (o *BackupRestoreStatusResponseModel) GetSimpleResults() []string`

GetSimpleResults returns the SimpleResults field if non-nil, zero value otherwise.

### GetSimpleResultsOk

`func (o *BackupRestoreStatusResponseModel) GetSimpleResultsOk() (*[]string, bool)`

GetSimpleResultsOk returns a tuple with the SimpleResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleResults

`func (o *BackupRestoreStatusResponseModel) SetSimpleResults(v []string)`

SetSimpleResults sets SimpleResults field to given value.

### HasSimpleResults

`func (o *BackupRestoreStatusResponseModel) HasSimpleResults() bool`

HasSimpleResults returns a boolean if a field has been set.

### SetSimpleResultsNil

`func (o *BackupRestoreStatusResponseModel) SetSimpleResultsNil(b bool)`

 SetSimpleResultsNil sets the value for SimpleResults to be an explicit nil

### UnsetSimpleResults
`func (o *BackupRestoreStatusResponseModel) UnsetSimpleResults()`

UnsetSimpleResults ensures that no value is present for SimpleResults, not even an explicit nil
### GetRestoreType

`func (o *BackupRestoreStatusResponseModel) GetRestoreType() BackupRestoreRestoreTypes`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *BackupRestoreStatusResponseModel) GetRestoreTypeOk() (*BackupRestoreRestoreTypes, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *BackupRestoreStatusResponseModel) SetRestoreType(v BackupRestoreRestoreTypes)`

SetRestoreType sets RestoreType field to given value.

### HasRestoreType

`func (o *BackupRestoreStatusResponseModel) HasRestoreType() bool`

HasRestoreType returns a boolean if a field has been set.

### GetFixups

`func (o *BackupRestoreStatusResponseModel) GetFixups() []BackupRestoreFixupModel`

GetFixups returns the Fixups field if non-nil, zero value otherwise.

### GetFixupsOk

`func (o *BackupRestoreStatusResponseModel) GetFixupsOk() (*[]BackupRestoreFixupModel, bool)`

GetFixupsOk returns a tuple with the Fixups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixups

`func (o *BackupRestoreStatusResponseModel) SetFixups(v []BackupRestoreFixupModel)`

SetFixups sets Fixups field to given value.

### HasFixups

`func (o *BackupRestoreStatusResponseModel) HasFixups() bool`

HasFixups returns a boolean if a field has been set.

### SetFixupsNil

`func (o *BackupRestoreStatusResponseModel) SetFixupsNil(b bool)`

 SetFixupsNil sets the value for Fixups to be an explicit nil

### UnsetFixups
`func (o *BackupRestoreStatusResponseModel) UnsetFixups()`

UnsetFixups ensures that no value is present for Fixups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


