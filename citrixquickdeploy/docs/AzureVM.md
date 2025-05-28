# AzureVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerState** | Pointer to **string** |  | [optional] 
**MachineSid** | Pointer to **string** |  | [optional] 
**MachineUid** | Pointer to **int32** |  | [optional] 
**ResourceGroupName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to [**CatalogResourceType**](CatalogResourceType.md) |  | [optional] 

## Methods

### NewAzureVM

`func NewAzureVM() *AzureVM`

NewAzureVM instantiates a new AzureVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVMWithDefaults

`func NewAzureVMWithDefaults() *AzureVM`

NewAzureVMWithDefaults instantiates a new AzureVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerState

`func (o *AzureVM) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *AzureVM) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *AzureVM) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *AzureVM) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetMachineSid

`func (o *AzureVM) GetMachineSid() string`

GetMachineSid returns the MachineSid field if non-nil, zero value otherwise.

### GetMachineSidOk

`func (o *AzureVM) GetMachineSidOk() (*string, bool)`

GetMachineSidOk returns a tuple with the MachineSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineSid

`func (o *AzureVM) SetMachineSid(v string)`

SetMachineSid sets MachineSid field to given value.

### HasMachineSid

`func (o *AzureVM) HasMachineSid() bool`

HasMachineSid returns a boolean if a field has been set.

### GetMachineUid

`func (o *AzureVM) GetMachineUid() int32`

GetMachineUid returns the MachineUid field if non-nil, zero value otherwise.

### GetMachineUidOk

`func (o *AzureVM) GetMachineUidOk() (*int32, bool)`

GetMachineUidOk returns a tuple with the MachineUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineUid

`func (o *AzureVM) SetMachineUid(v int32)`

SetMachineUid sets MachineUid field to given value.

### HasMachineUid

`func (o *AzureVM) HasMachineUid() bool`

HasMachineUid returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *AzureVM) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureVM) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureVM) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureVM) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetName

`func (o *AzureVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureVM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceType

`func (o *AzureVM) GetResourceType() CatalogResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AzureVM) GetResourceTypeOk() (*CatalogResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AzureVM) SetResourceType(v CatalogResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AzureVM) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


