# MachineCatalogVusComponentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** | Uid of the catalog. | [optional] 
**Id** | **string** | Id of the machine catalog. | 
**Name** | **string** | Name of the catalog. | 
**InstalledComponents** | Pointer to [**[]VDAComponentModel**](VDAComponentModel.md) | Installed VDA components of the catalog. | [optional] 
**AvailableComponents** | Pointer to [**[]VDAComponentModel**](VDAComponentModel.md) | Available VDA components of the catalog. | [optional] 
**AvailableFeatures** | Pointer to [**[]VDAFeatureModel**](VDAFeatureModel.md) | Available VDA features of the catalog. | [optional] 

## Methods

### NewMachineCatalogVusComponentResponseModel

`func NewMachineCatalogVusComponentResponseModel(id string, name string, ) *MachineCatalogVusComponentResponseModel`

NewMachineCatalogVusComponentResponseModel instantiates a new MachineCatalogVusComponentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogVusComponentResponseModelWithDefaults

`func NewMachineCatalogVusComponentResponseModelWithDefaults() *MachineCatalogVusComponentResponseModel`

NewMachineCatalogVusComponentResponseModelWithDefaults instantiates a new MachineCatalogVusComponentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *MachineCatalogVusComponentResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MachineCatalogVusComponentResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MachineCatalogVusComponentResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MachineCatalogVusComponentResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetId

`func (o *MachineCatalogVusComponentResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineCatalogVusComponentResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineCatalogVusComponentResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MachineCatalogVusComponentResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineCatalogVusComponentResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineCatalogVusComponentResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetInstalledComponents

`func (o *MachineCatalogVusComponentResponseModel) GetInstalledComponents() []VDAComponentModel`

GetInstalledComponents returns the InstalledComponents field if non-nil, zero value otherwise.

### GetInstalledComponentsOk

`func (o *MachineCatalogVusComponentResponseModel) GetInstalledComponentsOk() (*[]VDAComponentModel, bool)`

GetInstalledComponentsOk returns a tuple with the InstalledComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledComponents

`func (o *MachineCatalogVusComponentResponseModel) SetInstalledComponents(v []VDAComponentModel)`

SetInstalledComponents sets InstalledComponents field to given value.

### HasInstalledComponents

`func (o *MachineCatalogVusComponentResponseModel) HasInstalledComponents() bool`

HasInstalledComponents returns a boolean if a field has been set.

### GetAvailableComponents

`func (o *MachineCatalogVusComponentResponseModel) GetAvailableComponents() []VDAComponentModel`

GetAvailableComponents returns the AvailableComponents field if non-nil, zero value otherwise.

### GetAvailableComponentsOk

`func (o *MachineCatalogVusComponentResponseModel) GetAvailableComponentsOk() (*[]VDAComponentModel, bool)`

GetAvailableComponentsOk returns a tuple with the AvailableComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableComponents

`func (o *MachineCatalogVusComponentResponseModel) SetAvailableComponents(v []VDAComponentModel)`

SetAvailableComponents sets AvailableComponents field to given value.

### HasAvailableComponents

`func (o *MachineCatalogVusComponentResponseModel) HasAvailableComponents() bool`

HasAvailableComponents returns a boolean if a field has been set.

### GetAvailableFeatures

`func (o *MachineCatalogVusComponentResponseModel) GetAvailableFeatures() []VDAFeatureModel`

GetAvailableFeatures returns the AvailableFeatures field if non-nil, zero value otherwise.

### GetAvailableFeaturesOk

`func (o *MachineCatalogVusComponentResponseModel) GetAvailableFeaturesOk() (*[]VDAFeatureModel, bool)`

GetAvailableFeaturesOk returns a tuple with the AvailableFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFeatures

`func (o *MachineCatalogVusComponentResponseModel) SetAvailableFeatures(v []VDAFeatureModel)`

SetAvailableFeatures sets AvailableFeatures field to given value.

### HasAvailableFeatures

`func (o *MachineCatalogVusComponentResponseModel) HasAvailableFeatures() bool`

HasAvailableFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


