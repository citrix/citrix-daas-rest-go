# HypervisorPluginResponseModel

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

### NewHypervisorPluginResponseModel

`func NewHypervisorPluginResponseModel() *HypervisorPluginResponseModel`

NewHypervisorPluginResponseModel instantiates a new HypervisorPluginResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorPluginResponseModelWithDefaults

`func NewHypervisorPluginResponseModelWithDefaults() *HypervisorPluginResponseModel`

NewHypervisorPluginResponseModelWithDefaults instantiates a new HypervisorPluginResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *HypervisorPluginResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorPluginResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorPluginResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *HypervisorPluginResponseModel) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetDisplayName

`func (o *HypervisorPluginResponseModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HypervisorPluginResponseModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HypervisorPluginResponseModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HypervisorPluginResponseModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPluginFactoryName

`func (o *HypervisorPluginResponseModel) GetPluginFactoryName() string`

GetPluginFactoryName returns the PluginFactoryName field if non-nil, zero value otherwise.

### GetPluginFactoryNameOk

`func (o *HypervisorPluginResponseModel) GetPluginFactoryNameOk() (*string, bool)`

GetPluginFactoryNameOk returns a tuple with the PluginFactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginFactoryName

`func (o *HypervisorPluginResponseModel) SetPluginFactoryName(v string)`

SetPluginFactoryName sets PluginFactoryName field to given value.

### HasPluginFactoryName

`func (o *HypervisorPluginResponseModel) HasPluginFactoryName() bool`

HasPluginFactoryName returns a boolean if a field has been set.

### GetUsesCloudInfrastructure

`func (o *HypervisorPluginResponseModel) GetUsesCloudInfrastructure() bool`

GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field if non-nil, zero value otherwise.

### GetUsesCloudInfrastructureOk

`func (o *HypervisorPluginResponseModel) GetUsesCloudInfrastructureOk() (*bool, bool)`

GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCloudInfrastructure

`func (o *HypervisorPluginResponseModel) SetUsesCloudInfrastructure(v bool)`

SetUsesCloudInfrastructure sets UsesCloudInfrastructure field to given value.

### HasUsesCloudInfrastructure

`func (o *HypervisorPluginResponseModel) HasUsesCloudInfrastructure() bool`

HasUsesCloudInfrastructure returns a boolean if a field has been set.

### GetEnvironments

`func (o *HypervisorPluginResponseModel) GetEnvironments() []HypervisorAzureEnvironmentResponseModel`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *HypervisorPluginResponseModel) GetEnvironmentsOk() (*[]HypervisorAzureEnvironmentResponseModel, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *HypervisorPluginResponseModel) SetEnvironments(v []HypervisorAzureEnvironmentResponseModel)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *HypervisorPluginResponseModel) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetCitrixVerified

`func (o *HypervisorPluginResponseModel) GetCitrixVerified() bool`

GetCitrixVerified returns the CitrixVerified field if non-nil, zero value otherwise.

### GetCitrixVerifiedOk

`func (o *HypervisorPluginResponseModel) GetCitrixVerifiedOk() (*bool, bool)`

GetCitrixVerifiedOk returns a tuple with the CitrixVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixVerified

`func (o *HypervisorPluginResponseModel) SetCitrixVerified(v bool)`

SetCitrixVerified sets CitrixVerified field to given value.

### HasCitrixVerified

`func (o *HypervisorPluginResponseModel) HasCitrixVerified() bool`

HasCitrixVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


