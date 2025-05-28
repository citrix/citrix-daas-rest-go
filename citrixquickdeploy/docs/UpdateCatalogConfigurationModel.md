# UpdateCatalogConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreUsersConfigured** | Pointer to **bool** | Flag that indicates whether user has been assigned to the catalog | [optional] 
**AreAppsConfigured** | Pointer to **bool** | Indicates if apps or desktops have been configured for the catalog | [optional] 
**Region** | Pointer to **string** | Region where the catalog is being deployed | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


