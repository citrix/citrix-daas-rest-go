# DeploymentRegionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DeploymentRegionModel**](DeploymentRegionModel.md) | List of availabile regions | [optional] [readonly] 
**Regions** | Pointer to [**[]DeploymentRegionModel**](DeploymentRegionModel.md) | Alias of Items property for backward compatibility | [optional] 

## Methods

### NewDeploymentRegionsModel

`func NewDeploymentRegionsModel() *DeploymentRegionsModel`

NewDeploymentRegionsModel instantiates a new DeploymentRegionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRegionsModelWithDefaults

`func NewDeploymentRegionsModelWithDefaults() *DeploymentRegionsModel`

NewDeploymentRegionsModelWithDefaults instantiates a new DeploymentRegionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DeploymentRegionsModel) GetItems() []DeploymentRegionModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeploymentRegionsModel) GetItemsOk() (*[]DeploymentRegionModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeploymentRegionsModel) SetItems(v []DeploymentRegionModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *DeploymentRegionsModel) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRegions

`func (o *DeploymentRegionsModel) GetRegions() []DeploymentRegionModel`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *DeploymentRegionsModel) GetRegionsOk() (*[]DeploymentRegionModel, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *DeploymentRegionsModel) SetRegions(v []DeploymentRegionModel)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *DeploymentRegionsModel) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


