# StudioSettingsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MultiSites** | Pointer to [**[]MultiSiteModel**](MultiSiteModel.md) | Multiple sites configuration | [optional] 

## Methods

### NewStudioSettingsResponseModel

`func NewStudioSettingsResponseModel() *StudioSettingsResponseModel`

NewStudioSettingsResponseModel instantiates a new StudioSettingsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioSettingsResponseModelWithDefaults

`func NewStudioSettingsResponseModelWithDefaults() *StudioSettingsResponseModel`

NewStudioSettingsResponseModelWithDefaults instantiates a new StudioSettingsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultiSites

`func (o *StudioSettingsResponseModel) GetMultiSites() []MultiSiteModel`

GetMultiSites returns the MultiSites field if non-nil, zero value otherwise.

### GetMultiSitesOk

`func (o *StudioSettingsResponseModel) GetMultiSitesOk() (*[]MultiSiteModel, bool)`

GetMultiSitesOk returns a tuple with the MultiSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSites

`func (o *StudioSettingsResponseModel) SetMultiSites(v []MultiSiteModel)`

SetMultiSites sets MultiSites field to given value.

### HasMultiSites

`func (o *StudioSettingsResponseModel) HasMultiSites() bool`

HasMultiSites returns a boolean if a field has been set.

### SetMultiSitesNil

`func (o *StudioSettingsResponseModel) SetMultiSitesNil(b bool)`

 SetMultiSitesNil sets the value for MultiSites to be an explicit nil

### UnsetMultiSites
`func (o *StudioSettingsResponseModel) UnsetMultiSites()`

UnsetMultiSites ensures that no value is present for MultiSites, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


