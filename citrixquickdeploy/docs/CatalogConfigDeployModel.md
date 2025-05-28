# CatalogConfigDeployModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddCatalog** | Pointer to [**AddCatalogModel**](AddCatalogModel.md) |  | [optional] 
**AddAzureConfiguration** | Pointer to [**CatalogAzureConfigurationModel**](CatalogAzureConfigurationModel.md) |  | [optional] 
**AddCatalogDomain** | Pointer to [**CatalogDomainModel**](CatalogDomainModel.md) |  | [optional] 
**AddCatalogResourceLocation** | Pointer to [**CatalogResourceLocationConfiguration**](CatalogResourceLocationConfiguration.md) |  | [optional] 
**AddCatalogImage** | Pointer to [**CatalogTemplateImageModel**](CatalogTemplateImageModel.md) |  | [optional] 
**AddCatalogCapacity** | Pointer to [**CatalogCapacitySettingsModel**](CatalogCapacitySettingsModel.md) |  | [optional] 
**DeploySecrets** | Pointer to [**DeploySecretsModel**](DeploySecretsModel.md) |  | [optional] 
**AddSecureBrowser** | Pointer to [**CatalogSecureBrowserModel**](CatalogSecureBrowserModel.md) |  | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**WbcConfig** | Pointer to [**WbcConfig**](WbcConfig.md) |  | [optional] 

## Methods

### NewCatalogConfigDeployModel

`func NewCatalogConfigDeployModel() *CatalogConfigDeployModel`

NewCatalogConfigDeployModel instantiates a new CatalogConfigDeployModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogConfigDeployModelWithDefaults

`func NewCatalogConfigDeployModelWithDefaults() *CatalogConfigDeployModel`

NewCatalogConfigDeployModelWithDefaults instantiates a new CatalogConfigDeployModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddCatalog

`func (o *CatalogConfigDeployModel) GetAddCatalog() AddCatalogModel`

GetAddCatalog returns the AddCatalog field if non-nil, zero value otherwise.

### GetAddCatalogOk

`func (o *CatalogConfigDeployModel) GetAddCatalogOk() (*AddCatalogModel, bool)`

GetAddCatalogOk returns a tuple with the AddCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalog

`func (o *CatalogConfigDeployModel) SetAddCatalog(v AddCatalogModel)`

SetAddCatalog sets AddCatalog field to given value.

### HasAddCatalog

`func (o *CatalogConfigDeployModel) HasAddCatalog() bool`

HasAddCatalog returns a boolean if a field has been set.

### GetAddAzureConfiguration

`func (o *CatalogConfigDeployModel) GetAddAzureConfiguration() CatalogAzureConfigurationModel`

GetAddAzureConfiguration returns the AddAzureConfiguration field if non-nil, zero value otherwise.

### GetAddAzureConfigurationOk

`func (o *CatalogConfigDeployModel) GetAddAzureConfigurationOk() (*CatalogAzureConfigurationModel, bool)`

GetAddAzureConfigurationOk returns a tuple with the AddAzureConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAzureConfiguration

`func (o *CatalogConfigDeployModel) SetAddAzureConfiguration(v CatalogAzureConfigurationModel)`

SetAddAzureConfiguration sets AddAzureConfiguration field to given value.

### HasAddAzureConfiguration

`func (o *CatalogConfigDeployModel) HasAddAzureConfiguration() bool`

HasAddAzureConfiguration returns a boolean if a field has been set.

### GetAddCatalogDomain

`func (o *CatalogConfigDeployModel) GetAddCatalogDomain() CatalogDomainModel`

GetAddCatalogDomain returns the AddCatalogDomain field if non-nil, zero value otherwise.

### GetAddCatalogDomainOk

`func (o *CatalogConfigDeployModel) GetAddCatalogDomainOk() (*CatalogDomainModel, bool)`

GetAddCatalogDomainOk returns a tuple with the AddCatalogDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalogDomain

`func (o *CatalogConfigDeployModel) SetAddCatalogDomain(v CatalogDomainModel)`

SetAddCatalogDomain sets AddCatalogDomain field to given value.

### HasAddCatalogDomain

`func (o *CatalogConfigDeployModel) HasAddCatalogDomain() bool`

HasAddCatalogDomain returns a boolean if a field has been set.

### GetAddCatalogResourceLocation

`func (o *CatalogConfigDeployModel) GetAddCatalogResourceLocation() CatalogResourceLocationConfiguration`

GetAddCatalogResourceLocation returns the AddCatalogResourceLocation field if non-nil, zero value otherwise.

### GetAddCatalogResourceLocationOk

`func (o *CatalogConfigDeployModel) GetAddCatalogResourceLocationOk() (*CatalogResourceLocationConfiguration, bool)`

GetAddCatalogResourceLocationOk returns a tuple with the AddCatalogResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalogResourceLocation

`func (o *CatalogConfigDeployModel) SetAddCatalogResourceLocation(v CatalogResourceLocationConfiguration)`

SetAddCatalogResourceLocation sets AddCatalogResourceLocation field to given value.

### HasAddCatalogResourceLocation

`func (o *CatalogConfigDeployModel) HasAddCatalogResourceLocation() bool`

HasAddCatalogResourceLocation returns a boolean if a field has been set.

### GetAddCatalogImage

`func (o *CatalogConfigDeployModel) GetAddCatalogImage() CatalogTemplateImageModel`

GetAddCatalogImage returns the AddCatalogImage field if non-nil, zero value otherwise.

### GetAddCatalogImageOk

`func (o *CatalogConfigDeployModel) GetAddCatalogImageOk() (*CatalogTemplateImageModel, bool)`

GetAddCatalogImageOk returns a tuple with the AddCatalogImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalogImage

`func (o *CatalogConfigDeployModel) SetAddCatalogImage(v CatalogTemplateImageModel)`

SetAddCatalogImage sets AddCatalogImage field to given value.

### HasAddCatalogImage

`func (o *CatalogConfigDeployModel) HasAddCatalogImage() bool`

HasAddCatalogImage returns a boolean if a field has been set.

### GetAddCatalogCapacity

`func (o *CatalogConfigDeployModel) GetAddCatalogCapacity() CatalogCapacitySettingsModel`

GetAddCatalogCapacity returns the AddCatalogCapacity field if non-nil, zero value otherwise.

### GetAddCatalogCapacityOk

`func (o *CatalogConfigDeployModel) GetAddCatalogCapacityOk() (*CatalogCapacitySettingsModel, bool)`

GetAddCatalogCapacityOk returns a tuple with the AddCatalogCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCatalogCapacity

`func (o *CatalogConfigDeployModel) SetAddCatalogCapacity(v CatalogCapacitySettingsModel)`

SetAddCatalogCapacity sets AddCatalogCapacity field to given value.

### HasAddCatalogCapacity

`func (o *CatalogConfigDeployModel) HasAddCatalogCapacity() bool`

HasAddCatalogCapacity returns a boolean if a field has been set.

### GetDeploySecrets

`func (o *CatalogConfigDeployModel) GetDeploySecrets() DeploySecretsModel`

GetDeploySecrets returns the DeploySecrets field if non-nil, zero value otherwise.

### GetDeploySecretsOk

`func (o *CatalogConfigDeployModel) GetDeploySecretsOk() (*DeploySecretsModel, bool)`

GetDeploySecretsOk returns a tuple with the DeploySecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploySecrets

`func (o *CatalogConfigDeployModel) SetDeploySecrets(v DeploySecretsModel)`

SetDeploySecrets sets DeploySecrets field to given value.

### HasDeploySecrets

`func (o *CatalogConfigDeployModel) HasDeploySecrets() bool`

HasDeploySecrets returns a boolean if a field has been set.

### GetAddSecureBrowser

`func (o *CatalogConfigDeployModel) GetAddSecureBrowser() CatalogSecureBrowserModel`

GetAddSecureBrowser returns the AddSecureBrowser field if non-nil, zero value otherwise.

### GetAddSecureBrowserOk

`func (o *CatalogConfigDeployModel) GetAddSecureBrowserOk() (*CatalogSecureBrowserModel, bool)`

GetAddSecureBrowserOk returns a tuple with the AddSecureBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSecureBrowser

`func (o *CatalogConfigDeployModel) SetAddSecureBrowser(v CatalogSecureBrowserModel)`

SetAddSecureBrowser sets AddSecureBrowser field to given value.

### HasAddSecureBrowser

`func (o *CatalogConfigDeployModel) HasAddSecureBrowser() bool`

HasAddSecureBrowser returns a boolean if a field has been set.

### GetCspCustomerId

`func (o *CatalogConfigDeployModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *CatalogConfigDeployModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *CatalogConfigDeployModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *CatalogConfigDeployModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### GetCspSiteId

`func (o *CatalogConfigDeployModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *CatalogConfigDeployModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *CatalogConfigDeployModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *CatalogConfigDeployModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.

### GetWbcConfig

`func (o *CatalogConfigDeployModel) GetWbcConfig() WbcConfig`

GetWbcConfig returns the WbcConfig field if non-nil, zero value otherwise.

### GetWbcConfigOk

`func (o *CatalogConfigDeployModel) GetWbcConfigOk() (*WbcConfig, bool)`

GetWbcConfigOk returns a tuple with the WbcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbcConfig

`func (o *CatalogConfigDeployModel) SetWbcConfig(v WbcConfig)`

SetWbcConfig sets WbcConfig field to given value.

### HasWbcConfig

`func (o *CatalogConfigDeployModel) HasWbcConfig() bool`

HasWbcConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


