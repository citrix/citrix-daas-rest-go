# PatchMachines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) | The type of provider associated with the account | 
**InMaintenanceMode** | **bool** | The type of provider associated with the account | 
**MachineIds** | **[]string** | The list of machine ids to be updated | 

## Methods

### NewPatchMachines

`func NewPatchMachines(accountType AccountType, inMaintenanceMode bool, machineIds []string, ) *PatchMachines`

NewPatchMachines instantiates a new PatchMachines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchMachinesWithDefaults

`func NewPatchMachinesWithDefaults() *PatchMachines`

NewPatchMachinesWithDefaults instantiates a new PatchMachines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *PatchMachines) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PatchMachines) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PatchMachines) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetInMaintenanceMode

`func (o *PatchMachines) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *PatchMachines) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *PatchMachines) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetMachineIds

`func (o *PatchMachines) GetMachineIds() []string`

GetMachineIds returns the MachineIds field if non-nil, zero value otherwise.

### GetMachineIdsOk

`func (o *PatchMachines) GetMachineIdsOk() (*[]string, bool)`

GetMachineIdsOk returns a tuple with the MachineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineIds

`func (o *PatchMachines) SetMachineIds(v []string)`

SetMachineIds sets MachineIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


