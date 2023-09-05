# BackupRestoreGetComponentMemberRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | [**BckRstrAutoConfigComponents**](BckRstrAutoConfigComponents.md) |  | 
**Filters** | Pointer to **string** | Comma delimted component member names to restore; may include wildcards | [optional] 

## Methods

### NewBackupRestoreGetComponentMemberRequestModel

`func NewBackupRestoreGetComponentMemberRequestModel(component BckRstrAutoConfigComponents, ) *BackupRestoreGetComponentMemberRequestModel`

NewBackupRestoreGetComponentMemberRequestModel instantiates a new BackupRestoreGetComponentMemberRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreGetComponentMemberRequestModelWithDefaults

`func NewBackupRestoreGetComponentMemberRequestModelWithDefaults() *BackupRestoreGetComponentMemberRequestModel`

NewBackupRestoreGetComponentMemberRequestModelWithDefaults instantiates a new BackupRestoreGetComponentMemberRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *BackupRestoreGetComponentMemberRequestModel) GetComponent() BckRstrAutoConfigComponents`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *BackupRestoreGetComponentMemberRequestModel) GetComponentOk() (*BckRstrAutoConfigComponents, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *BackupRestoreGetComponentMemberRequestModel) SetComponent(v BckRstrAutoConfigComponents)`

SetComponent sets Component field to given value.


### GetFilters

`func (o *BackupRestoreGetComponentMemberRequestModel) GetFilters() string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BackupRestoreGetComponentMemberRequestModel) GetFiltersOk() (*string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BackupRestoreGetComponentMemberRequestModel) SetFilters(v string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *BackupRestoreGetComponentMemberRequestModel) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


