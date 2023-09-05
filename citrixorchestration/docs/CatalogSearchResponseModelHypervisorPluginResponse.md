# CatalogSearchResponseModelHypervisorPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Hypervisor display name.  | [optional] 
**PluginFactoryName** | Pointer to **string** | Plugin factory name. | [optional] 
**UsesCloudInfrastructure** | Pointer to **bool** | If use Cloud Infrastructure. | [optional] 
**Environments** | Pointer to [**[]HypervisorAzureEnvironmentResponseModel**](HypervisorAzureEnvironmentResponseModel.md) | This property only used when ConnectionType is AzureRM This property provides all supported Azure environments. | [optional] 
**CitrixVerified** | Pointer to **bool** | If the plugin is verified by Citrix. | [optional] 

## Methods

### NewCatalogSearchResponseModelHypervisorPluginResponse

`func NewCatalogSearchResponseModelHypervisorPluginResponse() *CatalogSearchResponseModelHypervisorPluginResponse`

NewCatalogSearchResponseModelHypervisorPluginResponse instantiates a new CatalogSearchResponseModelHypervisorPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogSearchResponseModelHypervisorPluginResponseWithDefaults

`func NewCatalogSearchResponseModelHypervisorPluginResponseWithDefaults() *CatalogSearchResponseModelHypervisorPluginResponse`

NewCatalogSearchResponseModelHypervisorPluginResponseWithDefaults instantiates a new CatalogSearchResponseModelHypervisorPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPluginFactoryName

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetPluginFactoryName() string`

GetPluginFactoryName returns the PluginFactoryName field if non-nil, zero value otherwise.

### GetPluginFactoryNameOk

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetPluginFactoryNameOk() (*string, bool)`

GetPluginFactoryNameOk returns a tuple with the PluginFactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginFactoryName

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) SetPluginFactoryName(v string)`

SetPluginFactoryName sets PluginFactoryName field to given value.

### HasPluginFactoryName

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) HasPluginFactoryName() bool`

HasPluginFactoryName returns a boolean if a field has been set.

### GetUsesCloudInfrastructure

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetUsesCloudInfrastructure() bool`

GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field if non-nil, zero value otherwise.

### GetUsesCloudInfrastructureOk

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetUsesCloudInfrastructureOk() (*bool, bool)`

GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCloudInfrastructure

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) SetUsesCloudInfrastructure(v bool)`

SetUsesCloudInfrastructure sets UsesCloudInfrastructure field to given value.

### HasUsesCloudInfrastructure

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) HasUsesCloudInfrastructure() bool`

HasUsesCloudInfrastructure returns a boolean if a field has been set.

### GetEnvironments

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetEnvironments() []HypervisorAzureEnvironmentResponseModel`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetEnvironmentsOk() (*[]HypervisorAzureEnvironmentResponseModel, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) SetEnvironments(v []HypervisorAzureEnvironmentResponseModel)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetCitrixVerified

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetCitrixVerified() bool`

GetCitrixVerified returns the CitrixVerified field if non-nil, zero value otherwise.

### GetCitrixVerifiedOk

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) GetCitrixVerifiedOk() (*bool, bool)`

GetCitrixVerifiedOk returns a tuple with the CitrixVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixVerified

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) SetCitrixVerified(v bool)`

SetCitrixVerified sets CitrixVerified field to given value.

### HasCitrixVerified

`func (o *CatalogSearchResponseModelHypervisorPluginResponse) HasCitrixVerified() bool`

HasCitrixVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


