# FtuConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFtuComplete** | Pointer to **bool** | Indicates if FTU wizard has been completed by the customer | [optional] 
**Connectivity** | Pointer to [**OnPremConnectivity**](OnPremConnectivity.md) | The tye of on-prem connectivity | [optional] 
**IsOnPremConnectionConfigured** | Pointer to **bool** | Indicates if at least one on-prem connection has been configured by the customer | [optional] 
**Authentication** | Pointer to [**AuthenticationMethod**](AuthenticationMethod.md) | Indicates the prefered FtuAuthMethod method for Workspace users | [optional] 
**CreatedWithAzureQuickDeploy** | Pointer to **bool** | Indicates whenver AzureQuickDeploy UI should be on/off | [optional] 
**StaleData** | Pointer to **bool** |  | [optional] 

## Methods

### NewFtuConfigurationModel

`func NewFtuConfigurationModel() *FtuConfigurationModel`

NewFtuConfigurationModel instantiates a new FtuConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFtuConfigurationModelWithDefaults

`func NewFtuConfigurationModelWithDefaults() *FtuConfigurationModel`

NewFtuConfigurationModelWithDefaults instantiates a new FtuConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFtuComplete

`func (o *FtuConfigurationModel) GetIsFtuComplete() bool`

GetIsFtuComplete returns the IsFtuComplete field if non-nil, zero value otherwise.

### GetIsFtuCompleteOk

`func (o *FtuConfigurationModel) GetIsFtuCompleteOk() (*bool, bool)`

GetIsFtuCompleteOk returns a tuple with the IsFtuComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFtuComplete

`func (o *FtuConfigurationModel) SetIsFtuComplete(v bool)`

SetIsFtuComplete sets IsFtuComplete field to given value.

### HasIsFtuComplete

`func (o *FtuConfigurationModel) HasIsFtuComplete() bool`

HasIsFtuComplete returns a boolean if a field has been set.

### GetConnectivity

`func (o *FtuConfigurationModel) GetConnectivity() OnPremConnectivity`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *FtuConfigurationModel) GetConnectivityOk() (*OnPremConnectivity, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *FtuConfigurationModel) SetConnectivity(v OnPremConnectivity)`

SetConnectivity sets Connectivity field to given value.

### HasConnectivity

`func (o *FtuConfigurationModel) HasConnectivity() bool`

HasConnectivity returns a boolean if a field has been set.

### GetIsOnPremConnectionConfigured

`func (o *FtuConfigurationModel) GetIsOnPremConnectionConfigured() bool`

GetIsOnPremConnectionConfigured returns the IsOnPremConnectionConfigured field if non-nil, zero value otherwise.

### GetIsOnPremConnectionConfiguredOk

`func (o *FtuConfigurationModel) GetIsOnPremConnectionConfiguredOk() (*bool, bool)`

GetIsOnPremConnectionConfiguredOk returns a tuple with the IsOnPremConnectionConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnPremConnectionConfigured

`func (o *FtuConfigurationModel) SetIsOnPremConnectionConfigured(v bool)`

SetIsOnPremConnectionConfigured sets IsOnPremConnectionConfigured field to given value.

### HasIsOnPremConnectionConfigured

`func (o *FtuConfigurationModel) HasIsOnPremConnectionConfigured() bool`

HasIsOnPremConnectionConfigured returns a boolean if a field has been set.

### GetAuthentication

`func (o *FtuConfigurationModel) GetAuthentication() AuthenticationMethod`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *FtuConfigurationModel) GetAuthenticationOk() (*AuthenticationMethod, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *FtuConfigurationModel) SetAuthentication(v AuthenticationMethod)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *FtuConfigurationModel) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetCreatedWithAzureQuickDeploy

`func (o *FtuConfigurationModel) GetCreatedWithAzureQuickDeploy() bool`

GetCreatedWithAzureQuickDeploy returns the CreatedWithAzureQuickDeploy field if non-nil, zero value otherwise.

### GetCreatedWithAzureQuickDeployOk

`func (o *FtuConfigurationModel) GetCreatedWithAzureQuickDeployOk() (*bool, bool)`

GetCreatedWithAzureQuickDeployOk returns a tuple with the CreatedWithAzureQuickDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedWithAzureQuickDeploy

`func (o *FtuConfigurationModel) SetCreatedWithAzureQuickDeploy(v bool)`

SetCreatedWithAzureQuickDeploy sets CreatedWithAzureQuickDeploy field to given value.

### HasCreatedWithAzureQuickDeploy

`func (o *FtuConfigurationModel) HasCreatedWithAzureQuickDeploy() bool`

HasCreatedWithAzureQuickDeploy returns a boolean if a field has been set.

### GetStaleData

`func (o *FtuConfigurationModel) GetStaleData() bool`

GetStaleData returns the StaleData field if non-nil, zero value otherwise.

### GetStaleDataOk

`func (o *FtuConfigurationModel) GetStaleDataOk() (*bool, bool)`

GetStaleDataOk returns a tuple with the StaleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleData

`func (o *FtuConfigurationModel) SetStaleData(v bool)`

SetStaleData sets StaleData field to given value.

### HasStaleData

`func (o *FtuConfigurationModel) HasStaleData() bool`

HasStaleData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


