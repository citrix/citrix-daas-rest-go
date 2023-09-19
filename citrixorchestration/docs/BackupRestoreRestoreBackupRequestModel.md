# BackupRestoreRestoreBackupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupName** | **string** | Name of backup to restore  | 
**HistoryUid** | **int32** | Hisotry Uid | 
**Component** | [**BckRstrAutoConfigComponents**](BckRstrAutoConfigComponents.md) |  | 
**Filters** | Pointer to **NullableString** | Comma delimted component member names to restore; may include wildcards | [optional] 
**WithPrereq** | Pointer to **bool** | Restore component prerequisites as necessary | [optional] 
**Checkmode** | Pointer to **bool** | Determine what should be restored but to not do the actual restore; report only              | [optional] 
**Notes** | Pointer to **NullableString** | Admin entered notes | [optional] 
**RestoreType** | Pointer to [**BackupRestoreRestoreTypes**](BackupRestoreRestoreTypes.md) |  | [optional] 

## Methods

### NewBackupRestoreRestoreBackupRequestModel

`func NewBackupRestoreRestoreBackupRequestModel(backupName string, historyUid int32, component BckRstrAutoConfigComponents, ) *BackupRestoreRestoreBackupRequestModel`

NewBackupRestoreRestoreBackupRequestModel instantiates a new BackupRestoreRestoreBackupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreRestoreBackupRequestModelWithDefaults

`func NewBackupRestoreRestoreBackupRequestModelWithDefaults() *BackupRestoreRestoreBackupRequestModel`

NewBackupRestoreRestoreBackupRequestModelWithDefaults instantiates a new BackupRestoreRestoreBackupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupName

`func (o *BackupRestoreRestoreBackupRequestModel) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *BackupRestoreRestoreBackupRequestModel) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *BackupRestoreRestoreBackupRequestModel) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.


### GetHistoryUid

`func (o *BackupRestoreRestoreBackupRequestModel) GetHistoryUid() int32`

GetHistoryUid returns the HistoryUid field if non-nil, zero value otherwise.

### GetHistoryUidOk

`func (o *BackupRestoreRestoreBackupRequestModel) GetHistoryUidOk() (*int32, bool)`

GetHistoryUidOk returns a tuple with the HistoryUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryUid

`func (o *BackupRestoreRestoreBackupRequestModel) SetHistoryUid(v int32)`

SetHistoryUid sets HistoryUid field to given value.


### GetComponent

`func (o *BackupRestoreRestoreBackupRequestModel) GetComponent() BckRstrAutoConfigComponents`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *BackupRestoreRestoreBackupRequestModel) GetComponentOk() (*BckRstrAutoConfigComponents, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *BackupRestoreRestoreBackupRequestModel) SetComponent(v BckRstrAutoConfigComponents)`

SetComponent sets Component field to given value.


### GetFilters

`func (o *BackupRestoreRestoreBackupRequestModel) GetFilters() string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BackupRestoreRestoreBackupRequestModel) GetFiltersOk() (*string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BackupRestoreRestoreBackupRequestModel) SetFilters(v string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *BackupRestoreRestoreBackupRequestModel) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *BackupRestoreRestoreBackupRequestModel) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *BackupRestoreRestoreBackupRequestModel) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetWithPrereq

`func (o *BackupRestoreRestoreBackupRequestModel) GetWithPrereq() bool`

GetWithPrereq returns the WithPrereq field if non-nil, zero value otherwise.

### GetWithPrereqOk

`func (o *BackupRestoreRestoreBackupRequestModel) GetWithPrereqOk() (*bool, bool)`

GetWithPrereqOk returns a tuple with the WithPrereq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPrereq

`func (o *BackupRestoreRestoreBackupRequestModel) SetWithPrereq(v bool)`

SetWithPrereq sets WithPrereq field to given value.

### HasWithPrereq

`func (o *BackupRestoreRestoreBackupRequestModel) HasWithPrereq() bool`

HasWithPrereq returns a boolean if a field has been set.

### GetCheckmode

`func (o *BackupRestoreRestoreBackupRequestModel) GetCheckmode() bool`

GetCheckmode returns the Checkmode field if non-nil, zero value otherwise.

### GetCheckmodeOk

`func (o *BackupRestoreRestoreBackupRequestModel) GetCheckmodeOk() (*bool, bool)`

GetCheckmodeOk returns a tuple with the Checkmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckmode

`func (o *BackupRestoreRestoreBackupRequestModel) SetCheckmode(v bool)`

SetCheckmode sets Checkmode field to given value.

### HasCheckmode

`func (o *BackupRestoreRestoreBackupRequestModel) HasCheckmode() bool`

HasCheckmode returns a boolean if a field has been set.

### GetNotes

`func (o *BackupRestoreRestoreBackupRequestModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BackupRestoreRestoreBackupRequestModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BackupRestoreRestoreBackupRequestModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BackupRestoreRestoreBackupRequestModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *BackupRestoreRestoreBackupRequestModel) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *BackupRestoreRestoreBackupRequestModel) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetRestoreType

`func (o *BackupRestoreRestoreBackupRequestModel) GetRestoreType() BackupRestoreRestoreTypes`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *BackupRestoreRestoreBackupRequestModel) GetRestoreTypeOk() (*BackupRestoreRestoreTypes, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *BackupRestoreRestoreBackupRequestModel) SetRestoreType(v BackupRestoreRestoreTypes)`

SetRestoreType sets RestoreType field to given value.

### HasRestoreType

`func (o *BackupRestoreRestoreBackupRequestModel) HasRestoreType() bool`

HasRestoreType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


