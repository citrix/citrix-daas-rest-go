# CitrixManagedCatalogConfigDeployModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddCatalog** | Pointer to [**NullableAddCitrixManagedCatalogModel**](AddCitrixManagedCatalogModel.md) |  | [optional] 
**AddCatalogOnPremConnectivity** | Pointer to [**NullableCatalogOnPremConnectivityModel**](CatalogOnPremConnectivityModel.md) |  | [optional] 
**AddCatalogDomain** | Pointer to [**NullableCatalogDomainModel**](CatalogDomainModel.md) |  | [optional] 
**AddCatalogResourceLocation** | Pointer to [**NullableCatalogResourceLocationConfiguration**](CatalogResourceLocationConfiguration.md) |  | [optional] 
**AddCatalogImage** | Pointer to [**NullableCatalogTemplateImageModel**](CatalogTemplateImageModel.md) |  | [optional] 
**AddCatalogCapacity** | Pointer to [**NullableCatalogCapacitySettingsModel**](CatalogCapacitySettingsModel.md) |  | [optional] 
**DeploySecrets** | Pointer to [**NullableDeploySecretsModel**](DeploySecretsModel.md) | Only needed for vnet peered/domain joined catalogs | [optional] 
**CspCustomerId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**ManagedSubscriptionId** | Pointer to **NullableString** | ID of the managed subscription the peering should be added to. | [optional] 
**SetDefault** | Pointer to **bool** |  | [optional] 
**PendingCitrixManagedUsers** | Pointer to [**[]NewAzureADUser**](NewAzureADUser.md) |  | [optional] 
**PendingUsers** | Pointer to [**NullablePendingUsersModel**](PendingUsersModel.md) |  | [optional] 
**WbcConfig** | Pointer to [**NullableWbcConfig**](WbcConfig.md) | Write back cache (MCS IO optimization) config | [optional] 

## Methods

### NewCitrixManagedCatalogConfigDeployModel

`func NewCitrixManagedCatalogConfigDeployModel() *CitrixManagedCatalogConfigDeployModel`

NewCitrixManagedCatalogConfigDeployModel instantiates a new CitrixManagedCatalogConfigDeployModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixManagedCatalogConfigDeployModelWithDefaults

`func NewCitrixManagedCatalogConfigDeployModelWithDefaults() *CitrixManagedCatalogConfigDeployModel`

NewCitrixManagedCatalogConfigDeployModelWithDefaults instantiates a new CitrixManagedCatalogConfigDeployModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddCatalog

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalog() AddCitrixManagedCatalogModel`

GetAddCatalog returns the AddCatalog field if non-nil, zero value otherwise.

### GetAddCatalogOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogOk() (*AddCitrixManagedCatalogModel, bool)`

GetAddCatalogOk returns a tuple with the AddCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalog

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalog(v AddCitrixManagedCatalogModel)`

SetAddCatalog sets AddCatalog field to given value.

### HasAddCatalog

`func (o *CitrixManagedCatalogConfigDeployModel) HasAddCatalog() bool`

HasAddCatalog returns a boolean if a field has been set.

### SetAddCatalogNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogNil(b bool)`

 SetAddCatalogNil sets the value for AddCatalog to be an explicit nil

### UnsetAddCatalog
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetAddCatalog()`

UnsetAddCatalog ensures that no value is present for AddCatalog, not even an explicit nil
### GetAddCatalogOnPremConnectivity

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogOnPremConnectivity() CatalogOnPremConnectivityModel`

GetAddCatalogOnPremConnectivity returns the AddCatalogOnPremConnectivity field if non-nil, zero value otherwise.

### GetAddCatalogOnPremConnectivityOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogOnPremConnectivityOk() (*CatalogOnPremConnectivityModel, bool)`

GetAddCatalogOnPremConnectivityOk returns a tuple with the AddCatalogOnPremConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalogOnPremConnectivity

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogOnPremConnectivity(v CatalogOnPremConnectivityModel)`

SetAddCatalogOnPremConnectivity sets AddCatalogOnPremConnectivity field to given value.

### HasAddCatalogOnPremConnectivity

`func (o *CitrixManagedCatalogConfigDeployModel) HasAddCatalogOnPremConnectivity() bool`

HasAddCatalogOnPremConnectivity returns a boolean if a field has been set.

### SetAddCatalogOnPremConnectivityNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogOnPremConnectivityNil(b bool)`

 SetAddCatalogOnPremConnectivityNil sets the value for AddCatalogOnPremConnectivity to be an explicit nil

### UnsetAddCatalogOnPremConnectivity
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetAddCatalogOnPremConnectivity()`

UnsetAddCatalogOnPremConnectivity ensures that no value is present for AddCatalogOnPremConnectivity, not even an explicit nil
### GetAddCatalogDomain

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogDomain() CatalogDomainModel`

GetAddCatalogDomain returns the AddCatalogDomain field if non-nil, zero value otherwise.

### GetAddCatalogDomainOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogDomainOk() (*CatalogDomainModel, bool)`

GetAddCatalogDomainOk returns a tuple with the AddCatalogDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalogDomain

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogDomain(v CatalogDomainModel)`

SetAddCatalogDomain sets AddCatalogDomain field to given value.

### HasAddCatalogDomain

`func (o *CitrixManagedCatalogConfigDeployModel) HasAddCatalogDomain() bool`

HasAddCatalogDomain returns a boolean if a field has been set.

### SetAddCatalogDomainNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogDomainNil(b bool)`

 SetAddCatalogDomainNil sets the value for AddCatalogDomain to be an explicit nil

### UnsetAddCatalogDomain
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetAddCatalogDomain()`

UnsetAddCatalogDomain ensures that no value is present for AddCatalogDomain, not even an explicit nil
### GetAddCatalogResourceLocation

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogResourceLocation() CatalogResourceLocationConfiguration`

GetAddCatalogResourceLocation returns the AddCatalogResourceLocation field if non-nil, zero value otherwise.

### GetAddCatalogResourceLocationOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogResourceLocationOk() (*CatalogResourceLocationConfiguration, bool)`

GetAddCatalogResourceLocationOk returns a tuple with the AddCatalogResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalogResourceLocation

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogResourceLocation(v CatalogResourceLocationConfiguration)`

SetAddCatalogResourceLocation sets AddCatalogResourceLocation field to given value.

### HasAddCatalogResourceLocation

`func (o *CitrixManagedCatalogConfigDeployModel) HasAddCatalogResourceLocation() bool`

HasAddCatalogResourceLocation returns a boolean if a field has been set.

### SetAddCatalogResourceLocationNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogResourceLocationNil(b bool)`

 SetAddCatalogResourceLocationNil sets the value for AddCatalogResourceLocation to be an explicit nil

### UnsetAddCatalogResourceLocation
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetAddCatalogResourceLocation()`

UnsetAddCatalogResourceLocation ensures that no value is present for AddCatalogResourceLocation, not even an explicit nil
### GetAddCatalogImage

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogImage() CatalogTemplateImageModel`

GetAddCatalogImage returns the AddCatalogImage field if non-nil, zero value otherwise.

### GetAddCatalogImageOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogImageOk() (*CatalogTemplateImageModel, bool)`

GetAddCatalogImageOk returns a tuple with the AddCatalogImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalogImage

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogImage(v CatalogTemplateImageModel)`

SetAddCatalogImage sets AddCatalogImage field to given value.

### HasAddCatalogImage

`func (o *CitrixManagedCatalogConfigDeployModel) HasAddCatalogImage() bool`

HasAddCatalogImage returns a boolean if a field has been set.

### SetAddCatalogImageNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogImageNil(b bool)`

 SetAddCatalogImageNil sets the value for AddCatalogImage to be an explicit nil

### UnsetAddCatalogImage
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetAddCatalogImage()`

UnsetAddCatalogImage ensures that no value is present for AddCatalogImage, not even an explicit nil
### GetAddCatalogCapacity

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogCapacity() CatalogCapacitySettingsModel`

GetAddCatalogCapacity returns the AddCatalogCapacity field if non-nil, zero value otherwise.

### GetAddCatalogCapacityOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetAddCatalogCapacityOk() (*CatalogCapacitySettingsModel, bool)`

GetAddCatalogCapacityOk returns a tuple with the AddCatalogCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalogCapacity

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogCapacity(v CatalogCapacitySettingsModel)`

SetAddCatalogCapacity sets AddCatalogCapacity field to given value.

### HasAddCatalogCapacity

`func (o *CitrixManagedCatalogConfigDeployModel) HasAddCatalogCapacity() bool`

HasAddCatalogCapacity returns a boolean if a field has been set.

### SetAddCatalogCapacityNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetAddCatalogCapacityNil(b bool)`

 SetAddCatalogCapacityNil sets the value for AddCatalogCapacity to be an explicit nil

### UnsetAddCatalogCapacity
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetAddCatalogCapacity()`

UnsetAddCatalogCapacity ensures that no value is present for AddCatalogCapacity, not even an explicit nil
### GetDeploySecrets

`func (o *CitrixManagedCatalogConfigDeployModel) GetDeploySecrets() DeploySecretsModel`

GetDeploySecrets returns the DeploySecrets field if non-nil, zero value otherwise.

### GetDeploySecretsOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetDeploySecretsOk() (*DeploySecretsModel, bool)`

GetDeploySecretsOk returns a tuple with the DeploySecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploySecrets

`func (o *CitrixManagedCatalogConfigDeployModel) SetDeploySecrets(v DeploySecretsModel)`

SetDeploySecrets sets DeploySecrets field to given value.

### HasDeploySecrets

`func (o *CitrixManagedCatalogConfigDeployModel) HasDeploySecrets() bool`

HasDeploySecrets returns a boolean if a field has been set.

### SetDeploySecretsNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetDeploySecretsNil(b bool)`

 SetDeploySecretsNil sets the value for DeploySecrets to be an explicit nil

### UnsetDeploySecrets
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetDeploySecrets()`

UnsetDeploySecrets ensures that no value is present for DeploySecrets, not even an explicit nil
### GetCspCustomerId

`func (o *CitrixManagedCatalogConfigDeployModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *CitrixManagedCatalogConfigDeployModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *CitrixManagedCatalogConfigDeployModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### SetCspCustomerIdNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetCspCustomerIdNil(b bool)`

 SetCspCustomerIdNil sets the value for CspCustomerId to be an explicit nil

### UnsetCspCustomerId
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetCspCustomerId()`

UnsetCspCustomerId ensures that no value is present for CspCustomerId, not even an explicit nil
### GetCspSiteId

`func (o *CitrixManagedCatalogConfigDeployModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *CitrixManagedCatalogConfigDeployModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *CitrixManagedCatalogConfigDeployModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.

### SetCspSiteIdNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetCspSiteIdNil(b bool)`

 SetCspSiteIdNil sets the value for CspSiteId to be an explicit nil

### UnsetCspSiteId
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetCspSiteId()`

UnsetCspSiteId ensures that no value is present for CspSiteId, not even an explicit nil
### GetManagedSubscriptionId

`func (o *CitrixManagedCatalogConfigDeployModel) GetManagedSubscriptionId() string`

GetManagedSubscriptionId returns the ManagedSubscriptionId field if non-nil, zero value otherwise.

### GetManagedSubscriptionIdOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetManagedSubscriptionIdOk() (*string, bool)`

GetManagedSubscriptionIdOk returns a tuple with the ManagedSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSubscriptionId

`func (o *CitrixManagedCatalogConfigDeployModel) SetManagedSubscriptionId(v string)`

SetManagedSubscriptionId sets ManagedSubscriptionId field to given value.

### HasManagedSubscriptionId

`func (o *CitrixManagedCatalogConfigDeployModel) HasManagedSubscriptionId() bool`

HasManagedSubscriptionId returns a boolean if a field has been set.

### SetManagedSubscriptionIdNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetManagedSubscriptionIdNil(b bool)`

 SetManagedSubscriptionIdNil sets the value for ManagedSubscriptionId to be an explicit nil

### UnsetManagedSubscriptionId
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetManagedSubscriptionId()`

UnsetManagedSubscriptionId ensures that no value is present for ManagedSubscriptionId, not even an explicit nil
### GetSetDefault

`func (o *CitrixManagedCatalogConfigDeployModel) GetSetDefault() bool`

GetSetDefault returns the SetDefault field if non-nil, zero value otherwise.

### GetSetDefaultOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetSetDefaultOk() (*bool, bool)`

GetSetDefaultOk returns a tuple with the SetDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDefault

`func (o *CitrixManagedCatalogConfigDeployModel) SetSetDefault(v bool)`

SetSetDefault sets SetDefault field to given value.

### HasSetDefault

`func (o *CitrixManagedCatalogConfigDeployModel) HasSetDefault() bool`

HasSetDefault returns a boolean if a field has been set.

### GetPendingCitrixManagedUsers

`func (o *CitrixManagedCatalogConfigDeployModel) GetPendingCitrixManagedUsers() []NewAzureADUser`

GetPendingCitrixManagedUsers returns the PendingCitrixManagedUsers field if non-nil, zero value otherwise.

### GetPendingCitrixManagedUsersOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetPendingCitrixManagedUsersOk() (*[]NewAzureADUser, bool)`

GetPendingCitrixManagedUsersOk returns a tuple with the PendingCitrixManagedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCitrixManagedUsers

`func (o *CitrixManagedCatalogConfigDeployModel) SetPendingCitrixManagedUsers(v []NewAzureADUser)`

SetPendingCitrixManagedUsers sets PendingCitrixManagedUsers field to given value.

### HasPendingCitrixManagedUsers

`func (o *CitrixManagedCatalogConfigDeployModel) HasPendingCitrixManagedUsers() bool`

HasPendingCitrixManagedUsers returns a boolean if a field has been set.

### SetPendingCitrixManagedUsersNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetPendingCitrixManagedUsersNil(b bool)`

 SetPendingCitrixManagedUsersNil sets the value for PendingCitrixManagedUsers to be an explicit nil

### UnsetPendingCitrixManagedUsers
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetPendingCitrixManagedUsers()`

UnsetPendingCitrixManagedUsers ensures that no value is present for PendingCitrixManagedUsers, not even an explicit nil
### GetPendingUsers

`func (o *CitrixManagedCatalogConfigDeployModel) GetPendingUsers() PendingUsersModel`

GetPendingUsers returns the PendingUsers field if non-nil, zero value otherwise.

### GetPendingUsersOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetPendingUsersOk() (*PendingUsersModel, bool)`

GetPendingUsersOk returns a tuple with the PendingUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUsers

`func (o *CitrixManagedCatalogConfigDeployModel) SetPendingUsers(v PendingUsersModel)`

SetPendingUsers sets PendingUsers field to given value.

### HasPendingUsers

`func (o *CitrixManagedCatalogConfigDeployModel) HasPendingUsers() bool`

HasPendingUsers returns a boolean if a field has been set.

### SetPendingUsersNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetPendingUsersNil(b bool)`

 SetPendingUsersNil sets the value for PendingUsers to be an explicit nil

### UnsetPendingUsers
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetPendingUsers()`

UnsetPendingUsers ensures that no value is present for PendingUsers, not even an explicit nil
### GetWbcConfig

`func (o *CitrixManagedCatalogConfigDeployModel) GetWbcConfig() WbcConfig`

GetWbcConfig returns the WbcConfig field if non-nil, zero value otherwise.

### GetWbcConfigOk

`func (o *CitrixManagedCatalogConfigDeployModel) GetWbcConfigOk() (*WbcConfig, bool)`

GetWbcConfigOk returns a tuple with the WbcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbcConfig

`func (o *CitrixManagedCatalogConfigDeployModel) SetWbcConfig(v WbcConfig)`

SetWbcConfig sets WbcConfig field to given value.

### HasWbcConfig

`func (o *CitrixManagedCatalogConfigDeployModel) HasWbcConfig() bool`

HasWbcConfig returns a boolean if a field has been set.

### SetWbcConfigNil

`func (o *CitrixManagedCatalogConfigDeployModel) SetWbcConfigNil(b bool)`

 SetWbcConfigNil sets the value for WbcConfig to be an explicit nil

### UnsetWbcConfig
`func (o *CitrixManagedCatalogConfigDeployModel) UnsetWbcConfig()`

UnsetWbcConfig ensures that no value is present for WbcConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


