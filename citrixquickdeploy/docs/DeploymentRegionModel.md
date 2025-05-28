# DeploymentRegionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the region | [optional] 
**Name** | Pointer to **string** | Display name of the region | [optional] 
**NonDomainJoinedCatalogsSupported** | Pointer to **bool** | Indicates if non-domain joined catalogs are supported in this region | [optional] 
**IsDefault** | Pointer to **bool** | Indicates that the region is the default to use | [optional] 
**IsHidden** | Pointer to **bool** | Inidicates if the regino is hidden from the UI | [optional] 

## Methods

### NewDeploymentRegionModel

`func NewDeploymentRegionModel() *DeploymentRegionModel`

NewDeploymentRegionModel instantiates a new DeploymentRegionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRegionModelWithDefaults

`func NewDeploymentRegionModelWithDefaults() *DeploymentRegionModel`

NewDeploymentRegionModelWithDefaults instantiates a new DeploymentRegionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentRegionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentRegionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentRegionModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentRegionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeploymentRegionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentRegionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentRegionModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentRegionModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNonDomainJoinedCatalogsSupported

`func (o *DeploymentRegionModel) GetNonDomainJoinedCatalogsSupported() bool`

GetNonDomainJoinedCatalogsSupported returns the NonDomainJoinedCatalogsSupported field if non-nil, zero value otherwise.

### GetNonDomainJoinedCatalogsSupportedOk

`func (o *DeploymentRegionModel) GetNonDomainJoinedCatalogsSupportedOk() (*bool, bool)`

GetNonDomainJoinedCatalogsSupportedOk returns a tuple with the NonDomainJoinedCatalogsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDomainJoinedCatalogsSupported

`func (o *DeploymentRegionModel) SetNonDomainJoinedCatalogsSupported(v bool)`

SetNonDomainJoinedCatalogsSupported sets NonDomainJoinedCatalogsSupported field to given value.

### HasNonDomainJoinedCatalogsSupported

`func (o *DeploymentRegionModel) HasNonDomainJoinedCatalogsSupported() bool`

HasNonDomainJoinedCatalogsSupported returns a boolean if a field has been set.

### GetIsDefault

`func (o *DeploymentRegionModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *DeploymentRegionModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *DeploymentRegionModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *DeploymentRegionModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsHidden

`func (o *DeploymentRegionModel) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *DeploymentRegionModel) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *DeploymentRegionModel) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *DeploymentRegionModel) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


