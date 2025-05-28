# CatalogsForBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogId** | **string** |  | 
**AzureSubscriptionId** | **string** |  | 
**Machines** | [**[]MachineForBackup**](MachineForBackup.md) |  | 
**BackupType** | [**SnapshotType**](SnapshotType.md) |  | 

## Methods

### NewCatalogsForBackup

`func NewCatalogsForBackup(catalogId string, azureSubscriptionId string, machines []MachineForBackup, backupType SnapshotType, ) *CatalogsForBackup`

NewCatalogsForBackup instantiates a new CatalogsForBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogsForBackupWithDefaults

`func NewCatalogsForBackupWithDefaults() *CatalogsForBackup`

NewCatalogsForBackupWithDefaults instantiates a new CatalogsForBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogId

`func (o *CatalogsForBackup) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *CatalogsForBackup) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *CatalogsForBackup) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.


### GetAzureSubscriptionId

`func (o *CatalogsForBackup) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *CatalogsForBackup) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *CatalogsForBackup) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### GetMachines

`func (o *CatalogsForBackup) GetMachines() []MachineForBackup`

GetMachines returns the Machines field if non-nil, zero value otherwise.

### GetMachinesOk

`func (o *CatalogsForBackup) GetMachinesOk() (*[]MachineForBackup, bool)`

GetMachinesOk returns a tuple with the Machines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachines

`func (o *CatalogsForBackup) SetMachines(v []MachineForBackup)`

SetMachines sets Machines field to given value.


### GetBackupType

`func (o *CatalogsForBackup) GetBackupType() SnapshotType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *CatalogsForBackup) GetBackupTypeOk() (*SnapshotType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *CatalogsForBackup) SetBackupType(v SnapshotType)`

SetBackupType sets BackupType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


