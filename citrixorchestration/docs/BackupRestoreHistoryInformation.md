# BackupRestoreHistoryInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** | Unique Identifier              | [optional] 
**Action** | [**BackupRestoreActions**](BackupRestoreActions.md) |  | 
**HistoryId** | **string** | History entry ID | 
**Successful** | **bool** | True if successful, false if not successful | 
**Duration** | Pointer to **int32** | Duration in seconds | [optional] 
**TimeZoneOffset** | Pointer to **int32** | TimeZone offset from UTC | [optional] 
**DateTime** | Pointer to **time.Time** | Date and time the action was started | [optional] 
**Notes** | Pointer to **string** | Notes about the action | [optional] 
**RestoreType** | Pointer to [**BackupRestoreRestoreTypes**](BackupRestoreRestoreTypes.md) |  | [optional] 
**Filters** | Pointer to **string** | Filters used (applicable only when the action was restore) | [optional] 
**WithPrerequisites** | Pointer to **bool** | With Prerequisites (applicable only when the action was restore) | [optional] 
**CheckMode** | Pointer to **bool** | Check mode (applicable only when the action was restore) | [optional] 
**ScheduleName** | Pointer to **string** | Name of schedule to perform backup (applicable only when the action was backup) | [optional] 
**Component** | Pointer to [**BckRstrAutoConfigComponents**](BckRstrAutoConfigComponents.md) |  | [optional] 
**ExecutionId** | Pointer to **string** | Execution Id | [optional] 
**BackupName** | Pointer to **string** | Backup name | [optional] 
**BackupFileSpec** | Pointer to **string** | Backup File Specification | [optional] 
**RelatedUid** | Pointer to **int32** | Related History UID for restore with checkmode set to true | [optional] 
**RelatedDate** | Pointer to **time.Time** | Related History date for restore with checkmode set to true | [optional] 
**RelatedIsCheckMode** | Pointer to **bool** | Related is run as check mode | [optional] 
**Pinned** | Pointer to **bool** | Backup is pinned | [optional] 
**AdministratorName** | Pointer to **string** | Administrator Name | [optional] 
**BackupDetails** | Pointer to **map[string]string** | Backup Status Details | [optional] 
**RestoreDetails** | Pointer to [**[]BackupRestoreRestoreSingleMemberModel**](BackupRestoreRestoreSingleMemberModel.md) | Restore Status Details | [optional] 
**SimpleResults** | Pointer to **[]string** | Simple Results (such as Get backed up member names) | [optional] 

## Methods

### NewBackupRestoreHistoryInformation

`func NewBackupRestoreHistoryInformation(action BackupRestoreActions, historyId string, successful bool, ) *BackupRestoreHistoryInformation`

NewBackupRestoreHistoryInformation instantiates a new BackupRestoreHistoryInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreHistoryInformationWithDefaults

`func NewBackupRestoreHistoryInformationWithDefaults() *BackupRestoreHistoryInformation`

NewBackupRestoreHistoryInformationWithDefaults instantiates a new BackupRestoreHistoryInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *BackupRestoreHistoryInformation) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *BackupRestoreHistoryInformation) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *BackupRestoreHistoryInformation) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *BackupRestoreHistoryInformation) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetAction

`func (o *BackupRestoreHistoryInformation) GetAction() BackupRestoreActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BackupRestoreHistoryInformation) GetActionOk() (*BackupRestoreActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BackupRestoreHistoryInformation) SetAction(v BackupRestoreActions)`

SetAction sets Action field to given value.


### GetHistoryId

`func (o *BackupRestoreHistoryInformation) GetHistoryId() string`

GetHistoryId returns the HistoryId field if non-nil, zero value otherwise.

### GetHistoryIdOk

`func (o *BackupRestoreHistoryInformation) GetHistoryIdOk() (*string, bool)`

GetHistoryIdOk returns a tuple with the HistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryId

`func (o *BackupRestoreHistoryInformation) SetHistoryId(v string)`

SetHistoryId sets HistoryId field to given value.


### GetSuccessful

`func (o *BackupRestoreHistoryInformation) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *BackupRestoreHistoryInformation) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *BackupRestoreHistoryInformation) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.


### GetDuration

`func (o *BackupRestoreHistoryInformation) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BackupRestoreHistoryInformation) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BackupRestoreHistoryInformation) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BackupRestoreHistoryInformation) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTimeZoneOffset

`func (o *BackupRestoreHistoryInformation) GetTimeZoneOffset() int32`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *BackupRestoreHistoryInformation) GetTimeZoneOffsetOk() (*int32, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *BackupRestoreHistoryInformation) SetTimeZoneOffset(v int32)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.

### HasTimeZoneOffset

`func (o *BackupRestoreHistoryInformation) HasTimeZoneOffset() bool`

HasTimeZoneOffset returns a boolean if a field has been set.

### GetDateTime

`func (o *BackupRestoreHistoryInformation) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *BackupRestoreHistoryInformation) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *BackupRestoreHistoryInformation) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *BackupRestoreHistoryInformation) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetNotes

`func (o *BackupRestoreHistoryInformation) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BackupRestoreHistoryInformation) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BackupRestoreHistoryInformation) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BackupRestoreHistoryInformation) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRestoreType

`func (o *BackupRestoreHistoryInformation) GetRestoreType() BackupRestoreRestoreTypes`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *BackupRestoreHistoryInformation) GetRestoreTypeOk() (*BackupRestoreRestoreTypes, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *BackupRestoreHistoryInformation) SetRestoreType(v BackupRestoreRestoreTypes)`

SetRestoreType sets RestoreType field to given value.

### HasRestoreType

`func (o *BackupRestoreHistoryInformation) HasRestoreType() bool`

HasRestoreType returns a boolean if a field has been set.

### GetFilters

`func (o *BackupRestoreHistoryInformation) GetFilters() string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BackupRestoreHistoryInformation) GetFiltersOk() (*string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BackupRestoreHistoryInformation) SetFilters(v string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *BackupRestoreHistoryInformation) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetWithPrerequisites

`func (o *BackupRestoreHistoryInformation) GetWithPrerequisites() bool`

GetWithPrerequisites returns the WithPrerequisites field if non-nil, zero value otherwise.

### GetWithPrerequisitesOk

`func (o *BackupRestoreHistoryInformation) GetWithPrerequisitesOk() (*bool, bool)`

GetWithPrerequisitesOk returns a tuple with the WithPrerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPrerequisites

`func (o *BackupRestoreHistoryInformation) SetWithPrerequisites(v bool)`

SetWithPrerequisites sets WithPrerequisites field to given value.

### HasWithPrerequisites

`func (o *BackupRestoreHistoryInformation) HasWithPrerequisites() bool`

HasWithPrerequisites returns a boolean if a field has been set.

### GetCheckMode

`func (o *BackupRestoreHistoryInformation) GetCheckMode() bool`

GetCheckMode returns the CheckMode field if non-nil, zero value otherwise.

### GetCheckModeOk

`func (o *BackupRestoreHistoryInformation) GetCheckModeOk() (*bool, bool)`

GetCheckModeOk returns a tuple with the CheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMode

`func (o *BackupRestoreHistoryInformation) SetCheckMode(v bool)`

SetCheckMode sets CheckMode field to given value.

### HasCheckMode

`func (o *BackupRestoreHistoryInformation) HasCheckMode() bool`

HasCheckMode returns a boolean if a field has been set.

### GetScheduleName

`func (o *BackupRestoreHistoryInformation) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *BackupRestoreHistoryInformation) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *BackupRestoreHistoryInformation) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *BackupRestoreHistoryInformation) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetComponent

`func (o *BackupRestoreHistoryInformation) GetComponent() BckRstrAutoConfigComponents`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *BackupRestoreHistoryInformation) GetComponentOk() (*BckRstrAutoConfigComponents, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *BackupRestoreHistoryInformation) SetComponent(v BckRstrAutoConfigComponents)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *BackupRestoreHistoryInformation) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetExecutionId

`func (o *BackupRestoreHistoryInformation) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BackupRestoreHistoryInformation) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BackupRestoreHistoryInformation) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BackupRestoreHistoryInformation) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetBackupName

`func (o *BackupRestoreHistoryInformation) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *BackupRestoreHistoryInformation) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *BackupRestoreHistoryInformation) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *BackupRestoreHistoryInformation) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### GetBackupFileSpec

`func (o *BackupRestoreHistoryInformation) GetBackupFileSpec() string`

GetBackupFileSpec returns the BackupFileSpec field if non-nil, zero value otherwise.

### GetBackupFileSpecOk

`func (o *BackupRestoreHistoryInformation) GetBackupFileSpecOk() (*string, bool)`

GetBackupFileSpecOk returns a tuple with the BackupFileSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFileSpec

`func (o *BackupRestoreHistoryInformation) SetBackupFileSpec(v string)`

SetBackupFileSpec sets BackupFileSpec field to given value.

### HasBackupFileSpec

`func (o *BackupRestoreHistoryInformation) HasBackupFileSpec() bool`

HasBackupFileSpec returns a boolean if a field has been set.

### GetRelatedUid

`func (o *BackupRestoreHistoryInformation) GetRelatedUid() int32`

GetRelatedUid returns the RelatedUid field if non-nil, zero value otherwise.

### GetRelatedUidOk

`func (o *BackupRestoreHistoryInformation) GetRelatedUidOk() (*int32, bool)`

GetRelatedUidOk returns a tuple with the RelatedUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedUid

`func (o *BackupRestoreHistoryInformation) SetRelatedUid(v int32)`

SetRelatedUid sets RelatedUid field to given value.

### HasRelatedUid

`func (o *BackupRestoreHistoryInformation) HasRelatedUid() bool`

HasRelatedUid returns a boolean if a field has been set.

### GetRelatedDate

`func (o *BackupRestoreHistoryInformation) GetRelatedDate() time.Time`

GetRelatedDate returns the RelatedDate field if non-nil, zero value otherwise.

### GetRelatedDateOk

`func (o *BackupRestoreHistoryInformation) GetRelatedDateOk() (*time.Time, bool)`

GetRelatedDateOk returns a tuple with the RelatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDate

`func (o *BackupRestoreHistoryInformation) SetRelatedDate(v time.Time)`

SetRelatedDate sets RelatedDate field to given value.

### HasRelatedDate

`func (o *BackupRestoreHistoryInformation) HasRelatedDate() bool`

HasRelatedDate returns a boolean if a field has been set.

### GetRelatedIsCheckMode

`func (o *BackupRestoreHistoryInformation) GetRelatedIsCheckMode() bool`

GetRelatedIsCheckMode returns the RelatedIsCheckMode field if non-nil, zero value otherwise.

### GetRelatedIsCheckModeOk

`func (o *BackupRestoreHistoryInformation) GetRelatedIsCheckModeOk() (*bool, bool)`

GetRelatedIsCheckModeOk returns a tuple with the RelatedIsCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIsCheckMode

`func (o *BackupRestoreHistoryInformation) SetRelatedIsCheckMode(v bool)`

SetRelatedIsCheckMode sets RelatedIsCheckMode field to given value.

### HasRelatedIsCheckMode

`func (o *BackupRestoreHistoryInformation) HasRelatedIsCheckMode() bool`

HasRelatedIsCheckMode returns a boolean if a field has been set.

### GetPinned

`func (o *BackupRestoreHistoryInformation) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *BackupRestoreHistoryInformation) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *BackupRestoreHistoryInformation) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *BackupRestoreHistoryInformation) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetAdministratorName

`func (o *BackupRestoreHistoryInformation) GetAdministratorName() string`

GetAdministratorName returns the AdministratorName field if non-nil, zero value otherwise.

### GetAdministratorNameOk

`func (o *BackupRestoreHistoryInformation) GetAdministratorNameOk() (*string, bool)`

GetAdministratorNameOk returns a tuple with the AdministratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorName

`func (o *BackupRestoreHistoryInformation) SetAdministratorName(v string)`

SetAdministratorName sets AdministratorName field to given value.

### HasAdministratorName

`func (o *BackupRestoreHistoryInformation) HasAdministratorName() bool`

HasAdministratorName returns a boolean if a field has been set.

### GetBackupDetails

`func (o *BackupRestoreHistoryInformation) GetBackupDetails() map[string]string`

GetBackupDetails returns the BackupDetails field if non-nil, zero value otherwise.

### GetBackupDetailsOk

`func (o *BackupRestoreHistoryInformation) GetBackupDetailsOk() (*map[string]string, bool)`

GetBackupDetailsOk returns a tuple with the BackupDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDetails

`func (o *BackupRestoreHistoryInformation) SetBackupDetails(v map[string]string)`

SetBackupDetails sets BackupDetails field to given value.

### HasBackupDetails

`func (o *BackupRestoreHistoryInformation) HasBackupDetails() bool`

HasBackupDetails returns a boolean if a field has been set.

### GetRestoreDetails

`func (o *BackupRestoreHistoryInformation) GetRestoreDetails() []BackupRestoreRestoreSingleMemberModel`

GetRestoreDetails returns the RestoreDetails field if non-nil, zero value otherwise.

### GetRestoreDetailsOk

`func (o *BackupRestoreHistoryInformation) GetRestoreDetailsOk() (*[]BackupRestoreRestoreSingleMemberModel, bool)`

GetRestoreDetailsOk returns a tuple with the RestoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDetails

`func (o *BackupRestoreHistoryInformation) SetRestoreDetails(v []BackupRestoreRestoreSingleMemberModel)`

SetRestoreDetails sets RestoreDetails field to given value.

### HasRestoreDetails

`func (o *BackupRestoreHistoryInformation) HasRestoreDetails() bool`

HasRestoreDetails returns a boolean if a field has been set.

### GetSimpleResults

`func (o *BackupRestoreHistoryInformation) GetSimpleResults() []string`

GetSimpleResults returns the SimpleResults field if non-nil, zero value otherwise.

### GetSimpleResultsOk

`func (o *BackupRestoreHistoryInformation) GetSimpleResultsOk() (*[]string, bool)`

GetSimpleResultsOk returns a tuple with the SimpleResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleResults

`func (o *BackupRestoreHistoryInformation) SetSimpleResults(v []string)`

SetSimpleResults sets SimpleResults field to given value.

### HasSimpleResults

`func (o *BackupRestoreHistoryInformation) HasSimpleResults() bool`

HasSimpleResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


