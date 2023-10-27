# BackupRestoreFixupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Fixup Identifier | [optional] 
**Severity** | Pointer to [**FixupSeverity**](FixupSeverity.md) |  | [optional] 
**LocalizedFixup** | Pointer to **NullableString** | List of parameters used in fixup string formatting | [optional] 

## Methods

### NewBackupRestoreFixupModel

`func NewBackupRestoreFixupModel() *BackupRestoreFixupModel`

NewBackupRestoreFixupModel instantiates a new BackupRestoreFixupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreFixupModelWithDefaults

`func NewBackupRestoreFixupModelWithDefaults() *BackupRestoreFixupModel`

NewBackupRestoreFixupModelWithDefaults instantiates a new BackupRestoreFixupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupRestoreFixupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRestoreFixupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRestoreFixupModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupRestoreFixupModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BackupRestoreFixupModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BackupRestoreFixupModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSeverity

`func (o *BackupRestoreFixupModel) GetSeverity() FixupSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *BackupRestoreFixupModel) GetSeverityOk() (*FixupSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *BackupRestoreFixupModel) SetSeverity(v FixupSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *BackupRestoreFixupModel) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetLocalizedFixup

`func (o *BackupRestoreFixupModel) GetLocalizedFixup() string`

GetLocalizedFixup returns the LocalizedFixup field if non-nil, zero value otherwise.

### GetLocalizedFixupOk

`func (o *BackupRestoreFixupModel) GetLocalizedFixupOk() (*string, bool)`

GetLocalizedFixupOk returns a tuple with the LocalizedFixup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedFixup

`func (o *BackupRestoreFixupModel) SetLocalizedFixup(v string)`

SetLocalizedFixup sets LocalizedFixup field to given value.

### HasLocalizedFixup

`func (o *BackupRestoreFixupModel) HasLocalizedFixup() bool`

HasLocalizedFixup returns a boolean if a field has been set.

### SetLocalizedFixupNil

`func (o *BackupRestoreFixupModel) SetLocalizedFixupNil(b bool)`

 SetLocalizedFixupNil sets the value for LocalizedFixup to be an explicit nil

### UnsetLocalizedFixup
`func (o *BackupRestoreFixupModel) UnsetLocalizedFixup()`

UnsetLocalizedFixup ensures that no value is present for LocalizedFixup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


