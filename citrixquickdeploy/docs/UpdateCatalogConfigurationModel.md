# UpdateCatalogConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreUsersConfigured** | Pointer to **NullableBool** | Flag that indicates whether user has been assigned to the catalog | [optional] 
**AreAppsConfigured** | Pointer to **NullableBool** | Indicates if apps or desktops have been configured for the catalog | [optional] 
**Region** | Pointer to **NullableString** | Region where the catalog is being deployed | [optional] 

## Methods

### NewUpdateCatalogConfigurationModel

`func NewUpdateCatalogConfigurationModel() *UpdateCatalogConfigurationModel`

NewUpdateCatalogConfigurationModel instantiates a new UpdateCatalogConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCatalogConfigurationModelWithDefaults

`func NewUpdateCatalogConfigurationModelWithDefaults() *UpdateCatalogConfigurationModel`

NewUpdateCatalogConfigurationModelWithDefaults instantiates a new UpdateCatalogConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreUsersConfigured

`func (o *UpdateCatalogConfigurationModel) GetAreUsersConfigured() bool`

GetAreUsersConfigured returns the AreUsersConfigured field if non-nil, zero value otherwise.

### GetAreUsersConfiguredOk

`func (o *UpdateCatalogConfigurationModel) GetAreUsersConfiguredOk() (*bool, bool)`

GetAreUsersConfiguredOk returns a tuple with the AreUsersConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreUsersConfigured

`func (o *UpdateCatalogConfigurationModel) SetAreUsersConfigured(v bool)`

SetAreUsersConfigured sets AreUsersConfigured field to given value.

### HasAreUsersConfigured

`func (o *UpdateCatalogConfigurationModel) HasAreUsersConfigured() bool`

HasAreUsersConfigured returns a boolean if a field has been set.

### SetAreUsersConfiguredNil

`func (o *UpdateCatalogConfigurationModel) SetAreUsersConfiguredNil(b bool)`

 SetAreUsersConfiguredNil sets the value for AreUsersConfigured to be an explicit nil

### UnsetAreUsersConfigured
`func (o *UpdateCatalogConfigurationModel) UnsetAreUsersConfigured()`

UnsetAreUsersConfigured ensures that no value is present for AreUsersConfigured, not even an explicit nil
### GetAreAppsConfigured

`func (o *UpdateCatalogConfigurationModel) GetAreAppsConfigured() bool`

GetAreAppsConfigured returns the AreAppsConfigured field if non-nil, zero value otherwise.

### GetAreAppsConfiguredOk

`func (o *UpdateCatalogConfigurationModel) GetAreAppsConfiguredOk() (*bool, bool)`

GetAreAppsConfiguredOk returns a tuple with the AreAppsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreAppsConfigured

`func (o *UpdateCatalogConfigurationModel) SetAreAppsConfigured(v bool)`

SetAreAppsConfigured sets AreAppsConfigured field to given value.

### HasAreAppsConfigured

`func (o *UpdateCatalogConfigurationModel) HasAreAppsConfigured() bool`

HasAreAppsConfigured returns a boolean if a field has been set.

### SetAreAppsConfiguredNil

`func (o *UpdateCatalogConfigurationModel) SetAreAppsConfiguredNil(b bool)`

 SetAreAppsConfiguredNil sets the value for AreAppsConfigured to be an explicit nil

### UnsetAreAppsConfigured
`func (o *UpdateCatalogConfigurationModel) UnsetAreAppsConfigured()`

UnsetAreAppsConfigured ensures that no value is present for AreAppsConfigured, not even an explicit nil
### GetRegion

`func (o *UpdateCatalogConfigurationModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpdateCatalogConfigurationModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpdateCatalogConfigurationModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UpdateCatalogConfigurationModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *UpdateCatalogConfigurationModel) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *UpdateCatalogConfigurationModel) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


