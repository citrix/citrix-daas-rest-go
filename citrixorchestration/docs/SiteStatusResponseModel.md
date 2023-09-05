# SiteStatusResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteState** | [**SiteState**](SiteState.md) |  | 
**CurrentUserCanPerformUpgrade** | Pointer to **bool** | If the current administrator can perform upgrade. | [optional] 
**UpgradedControllers** | Pointer to **[]string** | These controllers have been upgraded and will function normally at the latest version.              | [optional] 
**ControllersRequiringBinaryUpgrade** | Pointer to **[]string** | These controllers require their binaries to be upgraded to bring them inline with the rest of the site. | [optional] 
**LatestControllerVersion** | Pointer to **string** | The highest of all controller versions in the site. All controllers should be upgraded to this version. | [optional] 
**DatabaseServersRequiringUpgrade** | Pointer to **[]string** | These database servers must be upgraded to be compatible with the site controllers. | [optional] 
**DataStoresRequiringUpgrade** | Pointer to **[]string** | The database versions for these data stores are older than what is required by the most up-to-date controllers. | [optional] 

## Methods

### NewSiteStatusResponseModel

`func NewSiteStatusResponseModel(siteState SiteState, ) *SiteStatusResponseModel`

NewSiteStatusResponseModel instantiates a new SiteStatusResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteStatusResponseModelWithDefaults

`func NewSiteStatusResponseModelWithDefaults() *SiteStatusResponseModel`

NewSiteStatusResponseModelWithDefaults instantiates a new SiteStatusResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteState

`func (o *SiteStatusResponseModel) GetSiteState() SiteState`

GetSiteState returns the SiteState field if non-nil, zero value otherwise.

### GetSiteStateOk

`func (o *SiteStatusResponseModel) GetSiteStateOk() (*SiteState, bool)`

GetSiteStateOk returns a tuple with the SiteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteState

`func (o *SiteStatusResponseModel) SetSiteState(v SiteState)`

SetSiteState sets SiteState field to given value.


### GetCurrentUserCanPerformUpgrade

`func (o *SiteStatusResponseModel) GetCurrentUserCanPerformUpgrade() bool`

GetCurrentUserCanPerformUpgrade returns the CurrentUserCanPerformUpgrade field if non-nil, zero value otherwise.

### GetCurrentUserCanPerformUpgradeOk

`func (o *SiteStatusResponseModel) GetCurrentUserCanPerformUpgradeOk() (*bool, bool)`

GetCurrentUserCanPerformUpgradeOk returns a tuple with the CurrentUserCanPerformUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserCanPerformUpgrade

`func (o *SiteStatusResponseModel) SetCurrentUserCanPerformUpgrade(v bool)`

SetCurrentUserCanPerformUpgrade sets CurrentUserCanPerformUpgrade field to given value.

### HasCurrentUserCanPerformUpgrade

`func (o *SiteStatusResponseModel) HasCurrentUserCanPerformUpgrade() bool`

HasCurrentUserCanPerformUpgrade returns a boolean if a field has been set.

### GetUpgradedControllers

`func (o *SiteStatusResponseModel) GetUpgradedControllers() []string`

GetUpgradedControllers returns the UpgradedControllers field if non-nil, zero value otherwise.

### GetUpgradedControllersOk

`func (o *SiteStatusResponseModel) GetUpgradedControllersOk() (*[]string, bool)`

GetUpgradedControllersOk returns a tuple with the UpgradedControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedControllers

`func (o *SiteStatusResponseModel) SetUpgradedControllers(v []string)`

SetUpgradedControllers sets UpgradedControllers field to given value.

### HasUpgradedControllers

`func (o *SiteStatusResponseModel) HasUpgradedControllers() bool`

HasUpgradedControllers returns a boolean if a field has been set.

### GetControllersRequiringBinaryUpgrade

`func (o *SiteStatusResponseModel) GetControllersRequiringBinaryUpgrade() []string`

GetControllersRequiringBinaryUpgrade returns the ControllersRequiringBinaryUpgrade field if non-nil, zero value otherwise.

### GetControllersRequiringBinaryUpgradeOk

`func (o *SiteStatusResponseModel) GetControllersRequiringBinaryUpgradeOk() (*[]string, bool)`

GetControllersRequiringBinaryUpgradeOk returns a tuple with the ControllersRequiringBinaryUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllersRequiringBinaryUpgrade

`func (o *SiteStatusResponseModel) SetControllersRequiringBinaryUpgrade(v []string)`

SetControllersRequiringBinaryUpgrade sets ControllersRequiringBinaryUpgrade field to given value.

### HasControllersRequiringBinaryUpgrade

`func (o *SiteStatusResponseModel) HasControllersRequiringBinaryUpgrade() bool`

HasControllersRequiringBinaryUpgrade returns a boolean if a field has been set.

### GetLatestControllerVersion

`func (o *SiteStatusResponseModel) GetLatestControllerVersion() string`

GetLatestControllerVersion returns the LatestControllerVersion field if non-nil, zero value otherwise.

### GetLatestControllerVersionOk

`func (o *SiteStatusResponseModel) GetLatestControllerVersionOk() (*string, bool)`

GetLatestControllerVersionOk returns a tuple with the LatestControllerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestControllerVersion

`func (o *SiteStatusResponseModel) SetLatestControllerVersion(v string)`

SetLatestControllerVersion sets LatestControllerVersion field to given value.

### HasLatestControllerVersion

`func (o *SiteStatusResponseModel) HasLatestControllerVersion() bool`

HasLatestControllerVersion returns a boolean if a field has been set.

### GetDatabaseServersRequiringUpgrade

`func (o *SiteStatusResponseModel) GetDatabaseServersRequiringUpgrade() []string`

GetDatabaseServersRequiringUpgrade returns the DatabaseServersRequiringUpgrade field if non-nil, zero value otherwise.

### GetDatabaseServersRequiringUpgradeOk

`func (o *SiteStatusResponseModel) GetDatabaseServersRequiringUpgradeOk() (*[]string, bool)`

GetDatabaseServersRequiringUpgradeOk returns a tuple with the DatabaseServersRequiringUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseServersRequiringUpgrade

`func (o *SiteStatusResponseModel) SetDatabaseServersRequiringUpgrade(v []string)`

SetDatabaseServersRequiringUpgrade sets DatabaseServersRequiringUpgrade field to given value.

### HasDatabaseServersRequiringUpgrade

`func (o *SiteStatusResponseModel) HasDatabaseServersRequiringUpgrade() bool`

HasDatabaseServersRequiringUpgrade returns a boolean if a field has been set.

### GetDataStoresRequiringUpgrade

`func (o *SiteStatusResponseModel) GetDataStoresRequiringUpgrade() []string`

GetDataStoresRequiringUpgrade returns the DataStoresRequiringUpgrade field if non-nil, zero value otherwise.

### GetDataStoresRequiringUpgradeOk

`func (o *SiteStatusResponseModel) GetDataStoresRequiringUpgradeOk() (*[]string, bool)`

GetDataStoresRequiringUpgradeOk returns a tuple with the DataStoresRequiringUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoresRequiringUpgrade

`func (o *SiteStatusResponseModel) SetDataStoresRequiringUpgrade(v []string)`

SetDataStoresRequiringUpgrade sets DataStoresRequiringUpgrade field to given value.

### HasDataStoresRequiringUpgrade

`func (o *SiteStatusResponseModel) HasDataStoresRequiringUpgrade() bool`

HasDataStoresRequiringUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


