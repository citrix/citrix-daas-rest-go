# CitrixManagedCatalogConfigDeployModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddCatalog** | Pointer to [**AddCitrixManagedCatalogModel**](AddCitrixManagedCatalogModel.md) |  | [optional] 
**AddCatalogOnPremConnectivity** | Pointer to [**CatalogOnPremConnectivityModel**](CatalogOnPremConnectivityModel.md) |  | [optional] 
**AddCatalogDomain** | Pointer to [**CatalogDomainModel**](CatalogDomainModel.md) |  | [optional] 
**AddCatalogResourceLocation** | Pointer to [**CatalogResourceLocationConfiguration**](CatalogResourceLocationConfiguration.md) |  | [optional] 
**AddCatalogImage** | Pointer to [**CatalogTemplateImageModel**](CatalogTemplateImageModel.md) |  | [optional] 
**AddCatalogCapacity** | Pointer to [**CatalogCapacitySettingsModel**](CatalogCapacitySettingsModel.md) |  | [optional] 
**DeploySecrets** | Pointer to [**DeploySecretsModel**](DeploySecretsModel.md) | Only needed for vnet peered/domain joined catalogs | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**ManagedSubscriptionId** | Pointer to **string** | ID of the managed subscription the peering should be added to. | [optional] 
**SetDefault** | Pointer to **bool** |  | [optional] 
**PendingCitrixManagedUsers** | Pointer to [**[]NewAzureADUser**](NewAzureADUser.md) |  | [optional] 
**PendingUsers** | Pointer to [**PendingUsersModel**](PendingUsersModel.md) |  | [optional] 
**WbcConfig** | Pointer to [**WbcConfig**](WbcConfig.md) | Write back cache (MCS IO optimization) config | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


