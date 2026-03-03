# SnapshotsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SnapshotModel**](SnapshotModel.md) |  | [optional] 
**LastRestoreTime** | Pointer to **NullableTime** |  | [optional] 
**LastBackupTime** | Pointer to **NullableTime** |  | [optional] 
**LastRestoredSnapshotName** | Pointer to **NullableString** |  | [optional] 
**LastRestoredSnapshot** | Pointer to [**NullableSnapshotModel**](SnapshotModel.md) |  | [optional] 

## Methods

### NewSnapshotsModel

`func NewSnapshotsModel() *SnapshotsModel`

NewSnapshotsModel instantiates a new SnapshotsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotsModelWithDefaults

`func NewSnapshotsModelWithDefaults() *SnapshotsModel`

NewSnapshotsModelWithDefaults instantiates a new SnapshotsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SnapshotsModel) GetItems() []SnapshotModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SnapshotsModel) GetItemsOk() (*[]SnapshotModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SnapshotsModel) SetItems(v []SnapshotModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *SnapshotsModel) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *SnapshotsModel) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *SnapshotsModel) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetLastRestoreTime

`func (o *SnapshotsModel) GetLastRestoreTime() time.Time`

GetLastRestoreTime returns the LastRestoreTime field if non-nil, zero value otherwise.

### GetLastRestoreTimeOk

`func (o *SnapshotsModel) GetLastRestoreTimeOk() (*time.Time, bool)`

GetLastRestoreTimeOk returns a tuple with the LastRestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRestoreTime

`func (o *SnapshotsModel) SetLastRestoreTime(v time.Time)`

SetLastRestoreTime sets LastRestoreTime field to given value.

### HasLastRestoreTime

`func (o *SnapshotsModel) HasLastRestoreTime() bool`

HasLastRestoreTime returns a boolean if a field has been set.

### SetLastRestoreTimeNil

`func (o *SnapshotsModel) SetLastRestoreTimeNil(b bool)`

 SetLastRestoreTimeNil sets the value for LastRestoreTime to be an explicit nil

### UnsetLastRestoreTime
`func (o *SnapshotsModel) UnsetLastRestoreTime()`

UnsetLastRestoreTime ensures that no value is present for LastRestoreTime, not even an explicit nil
### GetLastBackupTime

`func (o *SnapshotsModel) GetLastBackupTime() time.Time`

GetLastBackupTime returns the LastBackupTime field if non-nil, zero value otherwise.

### GetLastBackupTimeOk

`func (o *SnapshotsModel) GetLastBackupTimeOk() (*time.Time, bool)`

GetLastBackupTimeOk returns a tuple with the LastBackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTime

`func (o *SnapshotsModel) SetLastBackupTime(v time.Time)`

SetLastBackupTime sets LastBackupTime field to given value.

### HasLastBackupTime

`func (o *SnapshotsModel) HasLastBackupTime() bool`

HasLastBackupTime returns a boolean if a field has been set.

### SetLastBackupTimeNil

`func (o *SnapshotsModel) SetLastBackupTimeNil(b bool)`

 SetLastBackupTimeNil sets the value for LastBackupTime to be an explicit nil

### UnsetLastBackupTime
`func (o *SnapshotsModel) UnsetLastBackupTime()`

UnsetLastBackupTime ensures that no value is present for LastBackupTime, not even an explicit nil
### GetLastRestoredSnapshotName

`func (o *SnapshotsModel) GetLastRestoredSnapshotName() string`

GetLastRestoredSnapshotName returns the LastRestoredSnapshotName field if non-nil, zero value otherwise.

### GetLastRestoredSnapshotNameOk

`func (o *SnapshotsModel) GetLastRestoredSnapshotNameOk() (*string, bool)`

GetLastRestoredSnapshotNameOk returns a tuple with the LastRestoredSnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRestoredSnapshotName

`func (o *SnapshotsModel) SetLastRestoredSnapshotName(v string)`

SetLastRestoredSnapshotName sets LastRestoredSnapshotName field to given value.

### HasLastRestoredSnapshotName

`func (o *SnapshotsModel) HasLastRestoredSnapshotName() bool`

HasLastRestoredSnapshotName returns a boolean if a field has been set.

### SetLastRestoredSnapshotNameNil

`func (o *SnapshotsModel) SetLastRestoredSnapshotNameNil(b bool)`

 SetLastRestoredSnapshotNameNil sets the value for LastRestoredSnapshotName to be an explicit nil

### UnsetLastRestoredSnapshotName
`func (o *SnapshotsModel) UnsetLastRestoredSnapshotName()`

UnsetLastRestoredSnapshotName ensures that no value is present for LastRestoredSnapshotName, not even an explicit nil
### GetLastRestoredSnapshot

`func (o *SnapshotsModel) GetLastRestoredSnapshot() SnapshotModel`

GetLastRestoredSnapshot returns the LastRestoredSnapshot field if non-nil, zero value otherwise.

### GetLastRestoredSnapshotOk

`func (o *SnapshotsModel) GetLastRestoredSnapshotOk() (*SnapshotModel, bool)`

GetLastRestoredSnapshotOk returns a tuple with the LastRestoredSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRestoredSnapshot

`func (o *SnapshotsModel) SetLastRestoredSnapshot(v SnapshotModel)`

SetLastRestoredSnapshot sets LastRestoredSnapshot field to given value.

### HasLastRestoredSnapshot

`func (o *SnapshotsModel) HasLastRestoredSnapshot() bool`

HasLastRestoredSnapshot returns a boolean if a field has been set.

### SetLastRestoredSnapshotNil

`func (o *SnapshotsModel) SetLastRestoredSnapshotNil(b bool)`

 SetLastRestoredSnapshotNil sets the value for LastRestoredSnapshot to be an explicit nil

### UnsetLastRestoredSnapshot
`func (o *SnapshotsModel) UnsetLastRestoredSnapshot()`

UnsetLastRestoredSnapshot ensures that no value is present for LastRestoredSnapshot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


