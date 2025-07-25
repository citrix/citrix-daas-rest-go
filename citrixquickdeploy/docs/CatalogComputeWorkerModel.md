# CatalogComputeWorkerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsePremiumStorage** | Pointer to **bool** | Indicates if Premium Storage will be used | [optional] 
**StorageType** | Pointer to [**CatalogCapacityStorageType**](CatalogCapacityStorageType.md) | Indicates type of Storage that will be used | [optional] 
**UseAzureHUB** | Pointer to **bool** | Indicates if the catalog VMs should be deployed with Azure HUB license | [optional] 
**MaxUsersPerVM** | Pointer to **int32** | Number of concurrent users per VM | [optional] 
**InstanceTypeId** | Pointer to **string** | Type of VM to create | [optional] 
**InstanceName** | Pointer to **string** | Name of the Azure VM Instance to use for the catalog | [optional] 
**BackupVmConfiguration** | Pointer to [**[]BackupVmConfiguration**](BackupVmConfiguration.md) | List of backup VM configurations | [optional] 
**UseManagedDisks** | Pointer to **bool** | Use managed disks for VMs in the catalog | [optional] 

## Methods

### NewCatalogComputeWorkerModel

`func NewCatalogComputeWorkerModel() *CatalogComputeWorkerModel`

NewCatalogComputeWorkerModel instantiates a new CatalogComputeWorkerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogComputeWorkerModelWithDefaults

`func NewCatalogComputeWorkerModelWithDefaults() *CatalogComputeWorkerModel`

NewCatalogComputeWorkerModelWithDefaults instantiates a new CatalogComputeWorkerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsePremiumStorage

`func (o *CatalogComputeWorkerModel) GetUsePremiumStorage() bool`

GetUsePremiumStorage returns the UsePremiumStorage field if non-nil, zero value otherwise.

### GetUsePremiumStorageOk

`func (o *CatalogComputeWorkerModel) GetUsePremiumStorageOk() (*bool, bool)`

GetUsePremiumStorageOk returns a tuple with the UsePremiumStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePremiumStorage

`func (o *CatalogComputeWorkerModel) SetUsePremiumStorage(v bool)`

SetUsePremiumStorage sets UsePremiumStorage field to given value.

### HasUsePremiumStorage

`func (o *CatalogComputeWorkerModel) HasUsePremiumStorage() bool`

HasUsePremiumStorage returns a boolean if a field has been set.

### GetStorageType

`func (o *CatalogComputeWorkerModel) GetStorageType() CatalogCapacityStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *CatalogComputeWorkerModel) GetStorageTypeOk() (*CatalogCapacityStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *CatalogComputeWorkerModel) SetStorageType(v CatalogCapacityStorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *CatalogComputeWorkerModel) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetUseAzureHUB

`func (o *CatalogComputeWorkerModel) GetUseAzureHUB() bool`

GetUseAzureHUB returns the UseAzureHUB field if non-nil, zero value otherwise.

### GetUseAzureHUBOk

`func (o *CatalogComputeWorkerModel) GetUseAzureHUBOk() (*bool, bool)`

GetUseAzureHUBOk returns a tuple with the UseAzureHUB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAzureHUB

`func (o *CatalogComputeWorkerModel) SetUseAzureHUB(v bool)`

SetUseAzureHUB sets UseAzureHUB field to given value.

### HasUseAzureHUB

`func (o *CatalogComputeWorkerModel) HasUseAzureHUB() bool`

HasUseAzureHUB returns a boolean if a field has been set.

### GetMaxUsersPerVM

`func (o *CatalogComputeWorkerModel) GetMaxUsersPerVM() int32`

GetMaxUsersPerVM returns the MaxUsersPerVM field if non-nil, zero value otherwise.

### GetMaxUsersPerVMOk

`func (o *CatalogComputeWorkerModel) GetMaxUsersPerVMOk() (*int32, bool)`

GetMaxUsersPerVMOk returns a tuple with the MaxUsersPerVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsersPerVM

`func (o *CatalogComputeWorkerModel) SetMaxUsersPerVM(v int32)`

SetMaxUsersPerVM sets MaxUsersPerVM field to given value.

### HasMaxUsersPerVM

`func (o *CatalogComputeWorkerModel) HasMaxUsersPerVM() bool`

HasMaxUsersPerVM returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *CatalogComputeWorkerModel) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *CatalogComputeWorkerModel) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *CatalogComputeWorkerModel) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *CatalogComputeWorkerModel) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### GetInstanceName

`func (o *CatalogComputeWorkerModel) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *CatalogComputeWorkerModel) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *CatalogComputeWorkerModel) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *CatalogComputeWorkerModel) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetBackupVmConfiguration

`func (o *CatalogComputeWorkerModel) GetBackupVmConfiguration() []BackupVmConfiguration`

GetBackupVmConfiguration returns the BackupVmConfiguration field if non-nil, zero value otherwise.

### GetBackupVmConfigurationOk

`func (o *CatalogComputeWorkerModel) GetBackupVmConfigurationOk() (*[]BackupVmConfiguration, bool)`

GetBackupVmConfigurationOk returns a tuple with the BackupVmConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupVmConfiguration

`func (o *CatalogComputeWorkerModel) SetBackupVmConfiguration(v []BackupVmConfiguration)`

SetBackupVmConfiguration sets BackupVmConfiguration field to given value.

### HasBackupVmConfiguration

`func (o *CatalogComputeWorkerModel) HasBackupVmConfiguration() bool`

HasBackupVmConfiguration returns a boolean if a field has been set.

### GetUseManagedDisks

`func (o *CatalogComputeWorkerModel) GetUseManagedDisks() bool`

GetUseManagedDisks returns the UseManagedDisks field if non-nil, zero value otherwise.

### GetUseManagedDisksOk

`func (o *CatalogComputeWorkerModel) GetUseManagedDisksOk() (*bool, bool)`

GetUseManagedDisksOk returns a tuple with the UseManagedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseManagedDisks

`func (o *CatalogComputeWorkerModel) SetUseManagedDisks(v bool)`

SetUseManagedDisks sets UseManagedDisks field to given value.

### HasUseManagedDisks

`func (o *CatalogComputeWorkerModel) HasUseManagedDisks() bool`

HasUseManagedDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


