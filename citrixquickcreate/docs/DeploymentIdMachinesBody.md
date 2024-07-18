# DeploymentIdMachinesBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**InMaintenanceMode** | **bool** | The type of provider associated with the account | 
**MachineIds** | **[]string** | The list of machine ids to be updated | 

## Methods

### NewDeploymentIdMachinesBody

`func NewDeploymentIdMachinesBody(accountType AccountType, inMaintenanceMode bool, machineIds []string, ) *DeploymentIdMachinesBody`

NewDeploymentIdMachinesBody instantiates a new DeploymentIdMachinesBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentIdMachinesBodyWithDefaults

`func NewDeploymentIdMachinesBodyWithDefaults() *DeploymentIdMachinesBody`

NewDeploymentIdMachinesBodyWithDefaults instantiates a new DeploymentIdMachinesBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *DeploymentIdMachinesBody) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *DeploymentIdMachinesBody) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *DeploymentIdMachinesBody) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetInMaintenanceMode

`func (o *DeploymentIdMachinesBody) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *DeploymentIdMachinesBody) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *DeploymentIdMachinesBody) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetMachineIds

`func (o *DeploymentIdMachinesBody) GetMachineIds() []string`

GetMachineIds returns the MachineIds field if non-nil, zero value otherwise.

### GetMachineIdsOk

`func (o *DeploymentIdMachinesBody) GetMachineIdsOk() (*[]string, bool)`

GetMachineIdsOk returns a tuple with the MachineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineIds

`func (o *DeploymentIdMachinesBody) SetMachineIds(v []string)`

SetMachineIds sets MachineIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


