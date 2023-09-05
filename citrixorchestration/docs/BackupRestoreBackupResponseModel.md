# BackupRestoreBackupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupName** | Pointer to **string** | Backup File name | [optional] 
**Notes** | Pointer to **string** | Notes | [optional] 
**Pinned** | Pointer to **bool** | Pinned &#x3D; true, unpinned &#x3D; false | [optional] 
**DateTime** | Pointer to **time.Time** | Backup date and time | [optional] 
**HistoryUid** | Pointer to **int32** | Hisotry Uid | [optional] 
**RelatedUid** | Pointer to **int32** | Related history UID | [optional] 
**RelatedDate** | Pointer to **time.Time** | Related history Date/Time | [optional] 
**RelatedIsCheckMode** | Pointer to **bool** | Related Restore was run in check mode | [optional] 
**Duration** | Pointer to **int32** | Duration in seconds              | [optional] 
**AdministratorName** | Pointer to **string** | Administrator Name | [optional] 
**Details** | Pointer to **map[string]string** | Details | [optional] 

## Methods

### NewBackupRestoreBackupResponseModel

`func NewBackupRestoreBackupResponseModel() *BackupRestoreBackupResponseModel`

NewBackupRestoreBackupResponseModel instantiates a new BackupRestoreBackupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreBackupResponseModelWithDefaults

`func NewBackupRestoreBackupResponseModelWithDefaults() *BackupRestoreBackupResponseModel`

NewBackupRestoreBackupResponseModelWithDefaults instantiates a new BackupRestoreBackupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupName

`func (o *BackupRestoreBackupResponseModel) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *BackupRestoreBackupResponseModel) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *BackupRestoreBackupResponseModel) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *BackupRestoreBackupResponseModel) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### GetNotes

`func (o *BackupRestoreBackupResponseModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BackupRestoreBackupResponseModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BackupRestoreBackupResponseModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BackupRestoreBackupResponseModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPinned

`func (o *BackupRestoreBackupResponseModel) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *BackupRestoreBackupResponseModel) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *BackupRestoreBackupResponseModel) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *BackupRestoreBackupResponseModel) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetDateTime

`func (o *BackupRestoreBackupResponseModel) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *BackupRestoreBackupResponseModel) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *BackupRestoreBackupResponseModel) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *BackupRestoreBackupResponseModel) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetHistoryUid

`func (o *BackupRestoreBackupResponseModel) GetHistoryUid() int32`

GetHistoryUid returns the HistoryUid field if non-nil, zero value otherwise.

### GetHistoryUidOk

`func (o *BackupRestoreBackupResponseModel) GetHistoryUidOk() (*int32, bool)`

GetHistoryUidOk returns a tuple with the HistoryUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryUid

`func (o *BackupRestoreBackupResponseModel) SetHistoryUid(v int32)`

SetHistoryUid sets HistoryUid field to given value.

### HasHistoryUid

`func (o *BackupRestoreBackupResponseModel) HasHistoryUid() bool`

HasHistoryUid returns a boolean if a field has been set.

### GetRelatedUid

`func (o *BackupRestoreBackupResponseModel) GetRelatedUid() int32`

GetRelatedUid returns the RelatedUid field if non-nil, zero value otherwise.

### GetRelatedUidOk

`func (o *BackupRestoreBackupResponseModel) GetRelatedUidOk() (*int32, bool)`

GetRelatedUidOk returns a tuple with the RelatedUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedUid

`func (o *BackupRestoreBackupResponseModel) SetRelatedUid(v int32)`

SetRelatedUid sets RelatedUid field to given value.

### HasRelatedUid

`func (o *BackupRestoreBackupResponseModel) HasRelatedUid() bool`

HasRelatedUid returns a boolean if a field has been set.

### GetRelatedDate

`func (o *BackupRestoreBackupResponseModel) GetRelatedDate() time.Time`

GetRelatedDate returns the RelatedDate field if non-nil, zero value otherwise.

### GetRelatedDateOk

`func (o *BackupRestoreBackupResponseModel) GetRelatedDateOk() (*time.Time, bool)`

GetRelatedDateOk returns a tuple with the RelatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDate

`func (o *BackupRestoreBackupResponseModel) SetRelatedDate(v time.Time)`

SetRelatedDate sets RelatedDate field to given value.

### HasRelatedDate

`func (o *BackupRestoreBackupResponseModel) HasRelatedDate() bool`

HasRelatedDate returns a boolean if a field has been set.

### GetRelatedIsCheckMode

`func (o *BackupRestoreBackupResponseModel) GetRelatedIsCheckMode() bool`

GetRelatedIsCheckMode returns the RelatedIsCheckMode field if non-nil, zero value otherwise.

### GetRelatedIsCheckModeOk

`func (o *BackupRestoreBackupResponseModel) GetRelatedIsCheckModeOk() (*bool, bool)`

GetRelatedIsCheckModeOk returns a tuple with the RelatedIsCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIsCheckMode

`func (o *BackupRestoreBackupResponseModel) SetRelatedIsCheckMode(v bool)`

SetRelatedIsCheckMode sets RelatedIsCheckMode field to given value.

### HasRelatedIsCheckMode

`func (o *BackupRestoreBackupResponseModel) HasRelatedIsCheckMode() bool`

HasRelatedIsCheckMode returns a boolean if a field has been set.

### GetDuration

`func (o *BackupRestoreBackupResponseModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BackupRestoreBackupResponseModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BackupRestoreBackupResponseModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BackupRestoreBackupResponseModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAdministratorName

`func (o *BackupRestoreBackupResponseModel) GetAdministratorName() string`

GetAdministratorName returns the AdministratorName field if non-nil, zero value otherwise.

### GetAdministratorNameOk

`func (o *BackupRestoreBackupResponseModel) GetAdministratorNameOk() (*string, bool)`

GetAdministratorNameOk returns a tuple with the AdministratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorName

`func (o *BackupRestoreBackupResponseModel) SetAdministratorName(v string)`

SetAdministratorName sets AdministratorName field to given value.

### HasAdministratorName

`func (o *BackupRestoreBackupResponseModel) HasAdministratorName() bool`

HasAdministratorName returns a boolean if a field has been set.

### GetDetails

`func (o *BackupRestoreBackupResponseModel) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BackupRestoreBackupResponseModel) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BackupRestoreBackupResponseModel) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BackupRestoreBackupResponseModel) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


