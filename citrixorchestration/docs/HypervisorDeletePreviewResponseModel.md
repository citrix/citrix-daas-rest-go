# HypervisorDeletePreviewResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanBeDeleted** | Pointer to **bool** | If the hypervisor can be deleted. | [optional] 
**TotalImpactedCatalogCount** | Pointer to **int32** | Totally impacted catalogs count. | [optional] 
**ProvisioningSchemeReferences** | Pointer to **int32** | Provisioning scheme references count. | [optional] 
**TaskReferences** | Pointer to **int32** | Task references count. | [optional] 
**UnknownMachines** | Pointer to **int32** | Unknown machines count. | [optional] 
**HostsToBeDeleted** | Pointer to **[]string** | Hosting units to be deleted. | [optional] 
**ImpactedCatalogs** | Pointer to **map[string]int32** | Impacted catalogs dictionary. | [optional] 

## Methods

### NewHypervisorDeletePreviewResponseModel

`func NewHypervisorDeletePreviewResponseModel() *HypervisorDeletePreviewResponseModel`

NewHypervisorDeletePreviewResponseModel instantiates a new HypervisorDeletePreviewResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorDeletePreviewResponseModelWithDefaults

`func NewHypervisorDeletePreviewResponseModelWithDefaults() *HypervisorDeletePreviewResponseModel`

NewHypervisorDeletePreviewResponseModelWithDefaults instantiates a new HypervisorDeletePreviewResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanBeDeleted

`func (o *HypervisorDeletePreviewResponseModel) GetCanBeDeleted() bool`

GetCanBeDeleted returns the CanBeDeleted field if non-nil, zero value otherwise.

### GetCanBeDeletedOk

`func (o *HypervisorDeletePreviewResponseModel) GetCanBeDeletedOk() (*bool, bool)`

GetCanBeDeletedOk returns a tuple with the CanBeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeDeleted

`func (o *HypervisorDeletePreviewResponseModel) SetCanBeDeleted(v bool)`

SetCanBeDeleted sets CanBeDeleted field to given value.

### HasCanBeDeleted

`func (o *HypervisorDeletePreviewResponseModel) HasCanBeDeleted() bool`

HasCanBeDeleted returns a boolean if a field has been set.

### GetTotalImpactedCatalogCount

`func (o *HypervisorDeletePreviewResponseModel) GetTotalImpactedCatalogCount() int32`

GetTotalImpactedCatalogCount returns the TotalImpactedCatalogCount field if non-nil, zero value otherwise.

### GetTotalImpactedCatalogCountOk

`func (o *HypervisorDeletePreviewResponseModel) GetTotalImpactedCatalogCountOk() (*int32, bool)`

GetTotalImpactedCatalogCountOk returns a tuple with the TotalImpactedCatalogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalImpactedCatalogCount

`func (o *HypervisorDeletePreviewResponseModel) SetTotalImpactedCatalogCount(v int32)`

SetTotalImpactedCatalogCount sets TotalImpactedCatalogCount field to given value.

### HasTotalImpactedCatalogCount

`func (o *HypervisorDeletePreviewResponseModel) HasTotalImpactedCatalogCount() bool`

HasTotalImpactedCatalogCount returns a boolean if a field has been set.

### GetProvisioningSchemeReferences

`func (o *HypervisorDeletePreviewResponseModel) GetProvisioningSchemeReferences() int32`

GetProvisioningSchemeReferences returns the ProvisioningSchemeReferences field if non-nil, zero value otherwise.

### GetProvisioningSchemeReferencesOk

`func (o *HypervisorDeletePreviewResponseModel) GetProvisioningSchemeReferencesOk() (*int32, bool)`

GetProvisioningSchemeReferencesOk returns a tuple with the ProvisioningSchemeReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeReferences

`func (o *HypervisorDeletePreviewResponseModel) SetProvisioningSchemeReferences(v int32)`

SetProvisioningSchemeReferences sets ProvisioningSchemeReferences field to given value.

### HasProvisioningSchemeReferences

`func (o *HypervisorDeletePreviewResponseModel) HasProvisioningSchemeReferences() bool`

HasProvisioningSchemeReferences returns a boolean if a field has been set.

### GetTaskReferences

`func (o *HypervisorDeletePreviewResponseModel) GetTaskReferences() int32`

GetTaskReferences returns the TaskReferences field if non-nil, zero value otherwise.

### GetTaskReferencesOk

`func (o *HypervisorDeletePreviewResponseModel) GetTaskReferencesOk() (*int32, bool)`

GetTaskReferencesOk returns a tuple with the TaskReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReferences

`func (o *HypervisorDeletePreviewResponseModel) SetTaskReferences(v int32)`

SetTaskReferences sets TaskReferences field to given value.

### HasTaskReferences

`func (o *HypervisorDeletePreviewResponseModel) HasTaskReferences() bool`

HasTaskReferences returns a boolean if a field has been set.

### GetUnknownMachines

`func (o *HypervisorDeletePreviewResponseModel) GetUnknownMachines() int32`

GetUnknownMachines returns the UnknownMachines field if non-nil, zero value otherwise.

### GetUnknownMachinesOk

`func (o *HypervisorDeletePreviewResponseModel) GetUnknownMachinesOk() (*int32, bool)`

GetUnknownMachinesOk returns a tuple with the UnknownMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownMachines

`func (o *HypervisorDeletePreviewResponseModel) SetUnknownMachines(v int32)`

SetUnknownMachines sets UnknownMachines field to given value.

### HasUnknownMachines

`func (o *HypervisorDeletePreviewResponseModel) HasUnknownMachines() bool`

HasUnknownMachines returns a boolean if a field has been set.

### GetHostsToBeDeleted

`func (o *HypervisorDeletePreviewResponseModel) GetHostsToBeDeleted() []string`

GetHostsToBeDeleted returns the HostsToBeDeleted field if non-nil, zero value otherwise.

### GetHostsToBeDeletedOk

`func (o *HypervisorDeletePreviewResponseModel) GetHostsToBeDeletedOk() (*[]string, bool)`

GetHostsToBeDeletedOk returns a tuple with the HostsToBeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsToBeDeleted

`func (o *HypervisorDeletePreviewResponseModel) SetHostsToBeDeleted(v []string)`

SetHostsToBeDeleted sets HostsToBeDeleted field to given value.

### HasHostsToBeDeleted

`func (o *HypervisorDeletePreviewResponseModel) HasHostsToBeDeleted() bool`

HasHostsToBeDeleted returns a boolean if a field has been set.

### GetImpactedCatalogs

`func (o *HypervisorDeletePreviewResponseModel) GetImpactedCatalogs() map[string]int32`

GetImpactedCatalogs returns the ImpactedCatalogs field if non-nil, zero value otherwise.

### GetImpactedCatalogsOk

`func (o *HypervisorDeletePreviewResponseModel) GetImpactedCatalogsOk() (*map[string]int32, bool)`

GetImpactedCatalogsOk returns a tuple with the ImpactedCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedCatalogs

`func (o *HypervisorDeletePreviewResponseModel) SetImpactedCatalogs(v map[string]int32)`

SetImpactedCatalogs sets ImpactedCatalogs field to given value.

### HasImpactedCatalogs

`func (o *HypervisorDeletePreviewResponseModel) HasImpactedCatalogs() bool`

HasImpactedCatalogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


