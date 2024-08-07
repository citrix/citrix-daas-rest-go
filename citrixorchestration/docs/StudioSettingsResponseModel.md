# StudioSettingsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MultiSites** | Pointer to [**[]MultiSiteModel**](MultiSiteModel.md) | Multiple sites configuration | [optional] 
**DefaultDomain** | Pointer to **NullableString** | The default domain to be used in the login page. | [optional] 

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
### GetDefaultDomain

`func (o *StudioSettingsResponseModel) GetDefaultDomain() string`

GetDefaultDomain returns the DefaultDomain field if non-nil, zero value otherwise.

### GetDefaultDomainOk

`func (o *StudioSettingsResponseModel) GetDefaultDomainOk() (*string, bool)`

GetDefaultDomainOk returns a tuple with the DefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomain

`func (o *StudioSettingsResponseModel) SetDefaultDomain(v string)`

SetDefaultDomain sets DefaultDomain field to given value.

### HasDefaultDomain

`func (o *StudioSettingsResponseModel) HasDefaultDomain() bool`

HasDefaultDomain returns a boolean if a field has been set.

### SetDefaultDomainNil

`func (o *StudioSettingsResponseModel) SetDefaultDomainNil(b bool)`

 SetDefaultDomainNil sets the value for DefaultDomain to be an explicit nil

### UnsetDefaultDomain
`func (o *StudioSettingsResponseModel) UnsetDefaultDomain()`

UnsetDefaultDomain ensures that no value is present for DefaultDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


