# HypervisorAzureConnectionDetailRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | Application ID of the service principal used to access the Azure APIs.  Required. | 
**ApplicationSecret** | **string** | The Application Secret of the service principal used to access the Azure APIs.  Required. Must be specified in the format indicated by ApplicationSecretFormat. | 
**ApplicationSecretFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**SubscriptionId** | **string** | Azure subscription ID.  Required. | 
**ActiveDirectoryId** | **string** | Azure active directory ID.  Required. | 
**Environment** | Pointer to [**AzureEnvironment**](AzureEnvironment.md) |  | [optional] 
**ManagementEndpoint** | Pointer to **string** | Azure management endpoint.  Required if Environment is Custom. Optional otherwise. | [optional] 
**AuthenticationAuthority** | Pointer to **string** | Azure authentication authority.  Required if Environment is Custom. Optional otherwise. | [optional] 
**StorageSuffix** | Pointer to **string** | Azure storage suffix.  Required if Environment is Custom. Optional otherwise. | [optional] 
**CustomProperties** | Pointer to **string** | The properties of host connection that are specific to the target hosting infrastructure. | [optional] 

## Methods

### NewHypervisorAzureConnectionDetailRequestModelAllOf

`func NewHypervisorAzureConnectionDetailRequestModelAllOf(applicationId string, applicationSecret string, subscriptionId string, activeDirectoryId string, ) *HypervisorAzureConnectionDetailRequestModelAllOf`

NewHypervisorAzureConnectionDetailRequestModelAllOf instantiates a new HypervisorAzureConnectionDetailRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorAzureConnectionDetailRequestModelAllOfWithDefaults

`func NewHypervisorAzureConnectionDetailRequestModelAllOfWithDefaults() *HypervisorAzureConnectionDetailRequestModelAllOf`

NewHypervisorAzureConnectionDetailRequestModelAllOfWithDefaults instantiates a new HypervisorAzureConnectionDetailRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationSecret

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetApplicationSecret() string`

GetApplicationSecret returns the ApplicationSecret field if non-nil, zero value otherwise.

### GetApplicationSecretOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetApplicationSecretOk() (*string, bool)`

GetApplicationSecretOk returns a tuple with the ApplicationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecret

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetApplicationSecret(v string)`

SetApplicationSecret sets ApplicationSecret field to given value.


### GetApplicationSecretFormat

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetApplicationSecretFormat() IdentityPasswordFormat`

GetApplicationSecretFormat returns the ApplicationSecretFormat field if non-nil, zero value otherwise.

### GetApplicationSecretFormatOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetApplicationSecretFormatOk() (*IdentityPasswordFormat, bool)`

GetApplicationSecretFormatOk returns a tuple with the ApplicationSecretFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecretFormat

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetApplicationSecretFormat(v IdentityPasswordFormat)`

SetApplicationSecretFormat sets ApplicationSecretFormat field to given value.

### HasApplicationSecretFormat

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) HasApplicationSecretFormat() bool`

HasApplicationSecretFormat returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetActiveDirectoryId

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.


### GetEnvironment

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetEnvironment() AzureEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetEnvironmentOk() (*AzureEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetEnvironment(v AzureEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetManagementEndpoint

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetManagementEndpoint() string`

GetManagementEndpoint returns the ManagementEndpoint field if non-nil, zero value otherwise.

### GetManagementEndpointOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetManagementEndpointOk() (*string, bool)`

GetManagementEndpointOk returns a tuple with the ManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEndpoint

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetManagementEndpoint(v string)`

SetManagementEndpoint sets ManagementEndpoint field to given value.

### HasManagementEndpoint

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) HasManagementEndpoint() bool`

HasManagementEndpoint returns a boolean if a field has been set.

### GetAuthenticationAuthority

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetAuthenticationAuthority() string`

GetAuthenticationAuthority returns the AuthenticationAuthority field if non-nil, zero value otherwise.

### GetAuthenticationAuthorityOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetAuthenticationAuthorityOk() (*string, bool)`

GetAuthenticationAuthorityOk returns a tuple with the AuthenticationAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAuthority

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetAuthenticationAuthority(v string)`

SetAuthenticationAuthority sets AuthenticationAuthority field to given value.

### HasAuthenticationAuthority

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) HasAuthenticationAuthority() bool`

HasAuthenticationAuthority returns a boolean if a field has been set.

### GetStorageSuffix

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetStorageSuffix() string`

GetStorageSuffix returns the StorageSuffix field if non-nil, zero value otherwise.

### GetStorageSuffixOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetStorageSuffixOk() (*string, bool)`

GetStorageSuffixOk returns a tuple with the StorageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSuffix

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetStorageSuffix(v string)`

SetStorageSuffix sets StorageSuffix field to given value.

### HasStorageSuffix

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) HasStorageSuffix() bool`

HasStorageSuffix returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorAzureConnectionDetailRequestModelAllOf) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


