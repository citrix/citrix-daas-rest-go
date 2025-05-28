# CatalogAzureConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureSubscriptionId** | **string** | ID of the azure subscription the catalog is associated with | 
**AzureResourceGroup** | **string** | Name of the resource group where all objects are going to be located in | 
**AzureVNet** | **string** | Name of the VNet all machines will be connected to | 
**AzureSubnet** | **string** | Address range of the machines in the catalog | 
**AzureVNetRegion** | Pointer to **string** | Region of the Azure VNet | [optional] 

## Methods

### NewCatalogAzureConfigurationModel

`func NewCatalogAzureConfigurationModel(azureSubscriptionId string, azureResourceGroup string, azureVNet string, azureSubnet string, ) *CatalogAzureConfigurationModel`

NewCatalogAzureConfigurationModel instantiates a new CatalogAzureConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogAzureConfigurationModelWithDefaults

`func NewCatalogAzureConfigurationModelWithDefaults() *CatalogAzureConfigurationModel`

NewCatalogAzureConfigurationModelWithDefaults instantiates a new CatalogAzureConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureSubscriptionId

`func (o *CatalogAzureConfigurationModel) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *CatalogAzureConfigurationModel) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *CatalogAzureConfigurationModel) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### GetAzureResourceGroup

`func (o *CatalogAzureConfigurationModel) GetAzureResourceGroup() string`

GetAzureResourceGroup returns the AzureResourceGroup field if non-nil, zero value otherwise.

### GetAzureResourceGroupOk

`func (o *CatalogAzureConfigurationModel) GetAzureResourceGroupOk() (*string, bool)`

GetAzureResourceGroupOk returns a tuple with the AzureResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroup

`func (o *CatalogAzureConfigurationModel) SetAzureResourceGroup(v string)`

SetAzureResourceGroup sets AzureResourceGroup field to given value.


### GetAzureVNet

`func (o *CatalogAzureConfigurationModel) GetAzureVNet() string`

GetAzureVNet returns the AzureVNet field if non-nil, zero value otherwise.

### GetAzureVNetOk

`func (o *CatalogAzureConfigurationModel) GetAzureVNetOk() (*string, bool)`

GetAzureVNetOk returns a tuple with the AzureVNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVNet

`func (o *CatalogAzureConfigurationModel) SetAzureVNet(v string)`

SetAzureVNet sets AzureVNet field to given value.


### GetAzureSubnet

`func (o *CatalogAzureConfigurationModel) GetAzureSubnet() string`

GetAzureSubnet returns the AzureSubnet field if non-nil, zero value otherwise.

### GetAzureSubnetOk

`func (o *CatalogAzureConfigurationModel) GetAzureSubnetOk() (*string, bool)`

GetAzureSubnetOk returns a tuple with the AzureSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubnet

`func (o *CatalogAzureConfigurationModel) SetAzureSubnet(v string)`

SetAzureSubnet sets AzureSubnet field to given value.


### GetAzureVNetRegion

`func (o *CatalogAzureConfigurationModel) GetAzureVNetRegion() string`

GetAzureVNetRegion returns the AzureVNetRegion field if non-nil, zero value otherwise.

### GetAzureVNetRegionOk

`func (o *CatalogAzureConfigurationModel) GetAzureVNetRegionOk() (*string, bool)`

GetAzureVNetRegionOk returns a tuple with the AzureVNetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVNetRegion

`func (o *CatalogAzureConfigurationModel) SetAzureVNetRegion(v string)`

SetAzureVNetRegion sets AzureVNetRegion field to given value.

### HasAzureVNetRegion

`func (o *CatalogAzureConfigurationModel) HasAzureVNetRegion() bool`

HasAzureVNetRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


