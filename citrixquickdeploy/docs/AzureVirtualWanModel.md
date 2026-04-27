# AzureVirtualWanModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Gets or sets the Azure resource ID of the Virtual WAN. | [optional] 
**Name** | Pointer to **NullableString** | Gets or sets the name of the Virtual WAN. | [optional] 
**ResourceGroup** | Pointer to **NullableString** | Gets or sets the resource group containing the Virtual WAN. | [optional] 
**Location** | Pointer to **NullableString** | Gets or sets the Azure region where the Virtual WAN is located. | [optional] 
**WanType** | Pointer to **NullableString** | Gets or sets the type of Virtual WAN (e.g. Standard, Basic). | [optional] 
**ProvisioningState** | Pointer to **NullableString** | Gets or sets the provisioning state of the Virtual WAN. | [optional] 

## Methods

### NewAzureVirtualWanModel

`func NewAzureVirtualWanModel() *AzureVirtualWanModel`

NewAzureVirtualWanModel instantiates a new AzureVirtualWanModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVirtualWanModelWithDefaults

`func NewAzureVirtualWanModelWithDefaults() *AzureVirtualWanModel`

NewAzureVirtualWanModelWithDefaults instantiates a new AzureVirtualWanModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureVirtualWanModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureVirtualWanModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureVirtualWanModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureVirtualWanModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AzureVirtualWanModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AzureVirtualWanModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *AzureVirtualWanModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVirtualWanModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVirtualWanModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureVirtualWanModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureVirtualWanModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureVirtualWanModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResourceGroup

`func (o *AzureVirtualWanModel) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureVirtualWanModel) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureVirtualWanModel) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AzureVirtualWanModel) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### SetResourceGroupNil

`func (o *AzureVirtualWanModel) SetResourceGroupNil(b bool)`

 SetResourceGroupNil sets the value for ResourceGroup to be an explicit nil

### UnsetResourceGroup
`func (o *AzureVirtualWanModel) UnsetResourceGroup()`

UnsetResourceGroup ensures that no value is present for ResourceGroup, not even an explicit nil
### GetLocation

`func (o *AzureVirtualWanModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureVirtualWanModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureVirtualWanModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureVirtualWanModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AzureVirtualWanModel) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AzureVirtualWanModel) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetWanType

`func (o *AzureVirtualWanModel) GetWanType() string`

GetWanType returns the WanType field if non-nil, zero value otherwise.

### GetWanTypeOk

`func (o *AzureVirtualWanModel) GetWanTypeOk() (*string, bool)`

GetWanTypeOk returns a tuple with the WanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanType

`func (o *AzureVirtualWanModel) SetWanType(v string)`

SetWanType sets WanType field to given value.

### HasWanType

`func (o *AzureVirtualWanModel) HasWanType() bool`

HasWanType returns a boolean if a field has been set.

### SetWanTypeNil

`func (o *AzureVirtualWanModel) SetWanTypeNil(b bool)`

 SetWanTypeNil sets the value for WanType to be an explicit nil

### UnsetWanType
`func (o *AzureVirtualWanModel) UnsetWanType()`

UnsetWanType ensures that no value is present for WanType, not even an explicit nil
### GetProvisioningState

`func (o *AzureVirtualWanModel) GetProvisioningState() string`

GetProvisioningState returns the ProvisioningState field if non-nil, zero value otherwise.

### GetProvisioningStateOk

`func (o *AzureVirtualWanModel) GetProvisioningStateOk() (*string, bool)`

GetProvisioningStateOk returns a tuple with the ProvisioningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningState

`func (o *AzureVirtualWanModel) SetProvisioningState(v string)`

SetProvisioningState sets ProvisioningState field to given value.

### HasProvisioningState

`func (o *AzureVirtualWanModel) HasProvisioningState() bool`

HasProvisioningState returns a boolean if a field has been set.

### SetProvisioningStateNil

`func (o *AzureVirtualWanModel) SetProvisioningStateNil(b bool)`

 SetProvisioningStateNil sets the value for ProvisioningState to be an explicit nil

### UnsetProvisioningState
`func (o *AzureVirtualWanModel) UnsetProvisioningState()`

UnsetProvisioningState ensures that no value is present for ProvisioningState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


