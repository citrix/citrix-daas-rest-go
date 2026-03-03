# AzureVMBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroupName** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to [**CatalogResourceType**](CatalogResourceType.md) |  | [optional] 
**NetworkInterfaceName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAzureVMBase

`func NewAzureVMBase() *AzureVMBase`

NewAzureVMBase instantiates a new AzureVMBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVMBaseWithDefaults

`func NewAzureVMBaseWithDefaults() *AzureVMBase`

NewAzureVMBaseWithDefaults instantiates a new AzureVMBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroupName

`func (o *AzureVMBase) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureVMBase) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureVMBase) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureVMBase) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *AzureVMBase) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *AzureVMBase) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetName

`func (o *AzureVMBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVMBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVMBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureVMBase) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureVMBase) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureVMBase) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResourceType

`func (o *AzureVMBase) GetResourceType() CatalogResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AzureVMBase) GetResourceTypeOk() (*CatalogResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AzureVMBase) SetResourceType(v CatalogResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AzureVMBase) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetNetworkInterfaceName

`func (o *AzureVMBase) GetNetworkInterfaceName() string`

GetNetworkInterfaceName returns the NetworkInterfaceName field if non-nil, zero value otherwise.

### GetNetworkInterfaceNameOk

`func (o *AzureVMBase) GetNetworkInterfaceNameOk() (*string, bool)`

GetNetworkInterfaceNameOk returns a tuple with the NetworkInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceName

`func (o *AzureVMBase) SetNetworkInterfaceName(v string)`

SetNetworkInterfaceName sets NetworkInterfaceName field to given value.

### HasNetworkInterfaceName

`func (o *AzureVMBase) HasNetworkInterfaceName() bool`

HasNetworkInterfaceName returns a boolean if a field has been set.

### SetNetworkInterfaceNameNil

`func (o *AzureVMBase) SetNetworkInterfaceNameNil(b bool)`

 SetNetworkInterfaceNameNil sets the value for NetworkInterfaceName to be an explicit nil

### UnsetNetworkInterfaceName
`func (o *AzureVMBase) UnsetNetworkInterfaceName()`

UnsetNetworkInterfaceName ensures that no value is present for NetworkInterfaceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


