# HypervisorDetailAzureResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | Application ID of the service principal used to access the Azure APIs. | 
**SubscriptionId** | **string** | Azure subscription ID. | 
**ActiveDirectoryId** | **string** | Azure active directory ID. | 
**Environment** | [**AzureEnvironment**](AzureEnvironment.md) |  | 
**ManagementEndpoint** | **string** | Azure management endpoint. | 
**AuthenticationAuthority** | **string** | Azure authentication authority. | 
**StorageSuffix** | **string** | Azure storage suffix. | 

## Methods

### NewHypervisorDetailAzureResponseModelAllOf

`func NewHypervisorDetailAzureResponseModelAllOf(applicationId string, subscriptionId string, activeDirectoryId string, environment AzureEnvironment, managementEndpoint string, authenticationAuthority string, storageSuffix string, ) *HypervisorDetailAzureResponseModelAllOf`

NewHypervisorDetailAzureResponseModelAllOf instantiates a new HypervisorDetailAzureResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorDetailAzureResponseModelAllOfWithDefaults

`func NewHypervisorDetailAzureResponseModelAllOfWithDefaults() *HypervisorDetailAzureResponseModelAllOf`

NewHypervisorDetailAzureResponseModelAllOfWithDefaults instantiates a new HypervisorDetailAzureResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *HypervisorDetailAzureResponseModelAllOf) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *HypervisorDetailAzureResponseModelAllOf) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *HypervisorDetailAzureResponseModelAllOf) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetSubscriptionId

`func (o *HypervisorDetailAzureResponseModelAllOf) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *HypervisorDetailAzureResponseModelAllOf) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *HypervisorDetailAzureResponseModelAllOf) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetActiveDirectoryId

`func (o *HypervisorDetailAzureResponseModelAllOf) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *HypervisorDetailAzureResponseModelAllOf) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *HypervisorDetailAzureResponseModelAllOf) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.


### GetEnvironment

`func (o *HypervisorDetailAzureResponseModelAllOf) GetEnvironment() AzureEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *HypervisorDetailAzureResponseModelAllOf) GetEnvironmentOk() (*AzureEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *HypervisorDetailAzureResponseModelAllOf) SetEnvironment(v AzureEnvironment)`

SetEnvironment sets Environment field to given value.


### GetManagementEndpoint

`func (o *HypervisorDetailAzureResponseModelAllOf) GetManagementEndpoint() string`

GetManagementEndpoint returns the ManagementEndpoint field if non-nil, zero value otherwise.

### GetManagementEndpointOk

`func (o *HypervisorDetailAzureResponseModelAllOf) GetManagementEndpointOk() (*string, bool)`

GetManagementEndpointOk returns a tuple with the ManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEndpoint

`func (o *HypervisorDetailAzureResponseModelAllOf) SetManagementEndpoint(v string)`

SetManagementEndpoint sets ManagementEndpoint field to given value.


### GetAuthenticationAuthority

`func (o *HypervisorDetailAzureResponseModelAllOf) GetAuthenticationAuthority() string`

GetAuthenticationAuthority returns the AuthenticationAuthority field if non-nil, zero value otherwise.

### GetAuthenticationAuthorityOk

`func (o *HypervisorDetailAzureResponseModelAllOf) GetAuthenticationAuthorityOk() (*string, bool)`

GetAuthenticationAuthorityOk returns a tuple with the AuthenticationAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAuthority

`func (o *HypervisorDetailAzureResponseModelAllOf) SetAuthenticationAuthority(v string)`

SetAuthenticationAuthority sets AuthenticationAuthority field to given value.


### GetStorageSuffix

`func (o *HypervisorDetailAzureResponseModelAllOf) GetStorageSuffix() string`

GetStorageSuffix returns the StorageSuffix field if non-nil, zero value otherwise.

### GetStorageSuffixOk

`func (o *HypervisorDetailAzureResponseModelAllOf) GetStorageSuffixOk() (*string, bool)`

GetStorageSuffixOk returns a tuple with the StorageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSuffix

`func (o *HypervisorDetailAzureResponseModelAllOf) SetStorageSuffix(v string)`

SetStorageSuffix sets StorageSuffix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


