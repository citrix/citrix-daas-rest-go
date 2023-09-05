# HypervisorResourcePoolPatchPreviewResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworksCanBeDeleted** | Pointer to **[]string** | These networks are not referenced and can be deleted normally. | [optional] 
**NetworksToBeDeleted** | Pointer to **[]string** | After comparing with the present hosting unit, these networks should be deleted. | [optional] 
**DeleteImpactedCatalogs** | Pointer to **map[string]map[string]int32** | Details of the referenced network and catalogs. | [optional] 

## Methods

### NewHypervisorResourcePoolPatchPreviewResponseModel

`func NewHypervisorResourcePoolPatchPreviewResponseModel() *HypervisorResourcePoolPatchPreviewResponseModel`

NewHypervisorResourcePoolPatchPreviewResponseModel instantiates a new HypervisorResourcePoolPatchPreviewResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolPatchPreviewResponseModelWithDefaults

`func NewHypervisorResourcePoolPatchPreviewResponseModelWithDefaults() *HypervisorResourcePoolPatchPreviewResponseModel`

NewHypervisorResourcePoolPatchPreviewResponseModelWithDefaults instantiates a new HypervisorResourcePoolPatchPreviewResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworksCanBeDeleted

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) GetNetworksCanBeDeleted() []string`

GetNetworksCanBeDeleted returns the NetworksCanBeDeleted field if non-nil, zero value otherwise.

### GetNetworksCanBeDeletedOk

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) GetNetworksCanBeDeletedOk() (*[]string, bool)`

GetNetworksCanBeDeletedOk returns a tuple with the NetworksCanBeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworksCanBeDeleted

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) SetNetworksCanBeDeleted(v []string)`

SetNetworksCanBeDeleted sets NetworksCanBeDeleted field to given value.

### HasNetworksCanBeDeleted

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) HasNetworksCanBeDeleted() bool`

HasNetworksCanBeDeleted returns a boolean if a field has been set.

### GetNetworksToBeDeleted

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) GetNetworksToBeDeleted() []string`

GetNetworksToBeDeleted returns the NetworksToBeDeleted field if non-nil, zero value otherwise.

### GetNetworksToBeDeletedOk

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) GetNetworksToBeDeletedOk() (*[]string, bool)`

GetNetworksToBeDeletedOk returns a tuple with the NetworksToBeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworksToBeDeleted

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) SetNetworksToBeDeleted(v []string)`

SetNetworksToBeDeleted sets NetworksToBeDeleted field to given value.

### HasNetworksToBeDeleted

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) HasNetworksToBeDeleted() bool`

HasNetworksToBeDeleted returns a boolean if a field has been set.

### GetDeleteImpactedCatalogs

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) GetDeleteImpactedCatalogs() map[string]map[string]int32`

GetDeleteImpactedCatalogs returns the DeleteImpactedCatalogs field if non-nil, zero value otherwise.

### GetDeleteImpactedCatalogsOk

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) GetDeleteImpactedCatalogsOk() (*map[string]map[string]int32, bool)`

GetDeleteImpactedCatalogsOk returns a tuple with the DeleteImpactedCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteImpactedCatalogs

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) SetDeleteImpactedCatalogs(v map[string]map[string]int32)`

SetDeleteImpactedCatalogs sets DeleteImpactedCatalogs field to given value.

### HasDeleteImpactedCatalogs

`func (o *HypervisorResourcePoolPatchPreviewResponseModel) HasDeleteImpactedCatalogs() bool`

HasDeleteImpactedCatalogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


