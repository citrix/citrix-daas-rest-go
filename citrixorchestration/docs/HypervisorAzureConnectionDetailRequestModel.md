# HypervisorAzureConnectionDetailRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the hypervisor to create.  Required. | 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**Scopes** | Pointer to **[]string** | Administrative scopes which the newly created hypervisor will be a part of. | [optional] 
**Tenants** | Pointer to **[]string** | Tenant(s) to associate the hypervisor with. | [optional] 
**MaxAbsoluteActiveActions** | Pointer to **int32** | Maximum number of actions that can execute in parallel on the hypervisor.  Optional; if not specified, an appropriate default is selected based on the ConnectionType. | [optional] 
**MaxAbsoluteNewActionsPerMinute** | Pointer to **int32** | Maximum number of actions that can be started on the hypervisor per-minute.  Optional; if not specified, an appropriate default is selected based on the ConnectionType. | [optional] 
**MaxPowerActionsPercentageOfMachines** | Pointer to **int32** | Maximum percentage of machines on the hypervisor which can have their power state changed simultaneously.  Optional; if not specified, an appropriate default is selected based on the ConnectionType. | [optional] 
**ConnectionOptions** | Pointer to **string** | Connection options to use for the hypervisor.  Optional. | [optional] 
**Zone** | Pointer to **string** | Zone the hypervisor is associated with.  Optional. If not specified, the hypervisor is associated with the primary zone. See PrimaryZone. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata to use for the hypervisor. Optional. | [optional] 
**ApiKey** | Pointer to **string** | The API key used to authenticate with the AWS APIs.  Required. | [optional] 
**SecretKey** | Pointer to **string** | The secret key used to authenticate with the AWS APIs.  Required. Must be specified in the format indicated by SecretKeyFormat. | [optional] 
**SecretKeyFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**Region** | Pointer to **string** | AWS region to connect to.  Optional.  If not specified, will connect to the global AWS APIs.  This can be used to discover the regions available within AWS.  Access to all other AWS resources requires the region to be set explicitly. | [optional] 
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
**ServiceAccountId** | Pointer to **string** | The service account ID used to access the Google Cloud APIs. Required. | [optional] 
**ServiceAccountCredentials** | Pointer to **string** | the JSON-encoded service account credentials used to access the Google Cloud APIs.  Required. Must be specified in the format indicated by ServiceAccountCredentialsFormat. | [optional] 
**ServiceAccountCredentialsFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**UserName** | Pointer to **string** | Hypervisor user name.  Required. | [optional] 
**Password** | Pointer to **string** | Hypervisor password.  Required. Must be specified in the format indicated by PasswordFormat. | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**Addresses** | Pointer to **[]string** | Hypervisor address(es).  At least one is required. | [optional] 
**PluginId** | Pointer to **string** | Custom hypervisor plugin ID.  Required if the ConnectionType is Custom.  Otherwise must not be specified. | [optional] 
**SslThumbprints** | Pointer to **[]string** | SSL certificate thumbprints to consider acceptable for this connection.  Optional.  If not specified, and the hypervisor uses SSL for its connection, the SSL certificate&#39;s root certification authority and any intermediate certificates must be trusted. | [optional] 
**SccmWakeUpProxy** | Pointer to **bool** | Indicates whether Microsoft System Center Configuration Manager 2012 SP1 Wake-up Proxy is used for power management.  Optional; default is &#x60;false&#x60;. | [optional] 
**WakeOnLanPackets** | Pointer to [**WakeOnLanTransmission**](WakeOnLanTransmission.md) |  | [optional] 
**TenancyOcid** | Pointer to **string** | Oracle Cloud Infrastructure tenancy to connect to. Required. | [optional] 
**ServiceAccountCredential** | Pointer to **string** | The private key string to access the Oracle Cloud Infrastructure APIs. Required. Must be specified in the format indicated by ServiceAccountCredentialFormat. | [optional] 
**ServiceAccountCredentialFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**ServiceAccountFingerprint** | Pointer to **string** | The fingerprint of the public key associate with the ServiceAccountCredential. | [optional] 
**OciEnvironment** | Pointer to [**OciEnvironment**](OciEnvironment.md) |  | [optional] 

## Methods

### NewHypervisorAzureConnectionDetailRequestModel

`func NewHypervisorAzureConnectionDetailRequestModel(name string, connectionType HypervisorConnectionType, applicationId string, applicationSecret string, subscriptionId string, activeDirectoryId string, ) *HypervisorAzureConnectionDetailRequestModel`

NewHypervisorAzureConnectionDetailRequestModel instantiates a new HypervisorAzureConnectionDetailRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorAzureConnectionDetailRequestModelWithDefaults

`func NewHypervisorAzureConnectionDetailRequestModelWithDefaults() *HypervisorAzureConnectionDetailRequestModel`

NewHypervisorAzureConnectionDetailRequestModelWithDefaults instantiates a new HypervisorAzureConnectionDetailRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HypervisorAzureConnectionDetailRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorAzureConnectionDetailRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetConnectionType

`func (o *HypervisorAzureConnectionDetailRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorAzureConnectionDetailRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetScopes

`func (o *HypervisorAzureConnectionDetailRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorAzureConnectionDetailRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *HypervisorAzureConnectionDetailRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTenants

`func (o *HypervisorAzureConnectionDetailRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorAzureConnectionDetailRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorAzureConnectionDetailRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetMaxAbsoluteActiveActions

`func (o *HypervisorAzureConnectionDetailRequestModel) GetMaxAbsoluteActiveActions() int32`

GetMaxAbsoluteActiveActions returns the MaxAbsoluteActiveActions field if non-nil, zero value otherwise.

### GetMaxAbsoluteActiveActionsOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetMaxAbsoluteActiveActionsOk() (*int32, bool)`

GetMaxAbsoluteActiveActionsOk returns a tuple with the MaxAbsoluteActiveActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteActiveActions

`func (o *HypervisorAzureConnectionDetailRequestModel) SetMaxAbsoluteActiveActions(v int32)`

SetMaxAbsoluteActiveActions sets MaxAbsoluteActiveActions field to given value.

### HasMaxAbsoluteActiveActions

`func (o *HypervisorAzureConnectionDetailRequestModel) HasMaxAbsoluteActiveActions() bool`

HasMaxAbsoluteActiveActions returns a boolean if a field has been set.

### GetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorAzureConnectionDetailRequestModel) GetMaxAbsoluteNewActionsPerMinute() int32`

GetMaxAbsoluteNewActionsPerMinute returns the MaxAbsoluteNewActionsPerMinute field if non-nil, zero value otherwise.

### GetMaxAbsoluteNewActionsPerMinuteOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetMaxAbsoluteNewActionsPerMinuteOk() (*int32, bool)`

GetMaxAbsoluteNewActionsPerMinuteOk returns a tuple with the MaxAbsoluteNewActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorAzureConnectionDetailRequestModel) SetMaxAbsoluteNewActionsPerMinute(v int32)`

SetMaxAbsoluteNewActionsPerMinute sets MaxAbsoluteNewActionsPerMinute field to given value.

### HasMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorAzureConnectionDetailRequestModel) HasMaxAbsoluteNewActionsPerMinute() bool`

HasMaxAbsoluteNewActionsPerMinute returns a boolean if a field has been set.

### GetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorAzureConnectionDetailRequestModel) GetMaxPowerActionsPercentageOfMachines() int32`

GetMaxPowerActionsPercentageOfMachines returns the MaxPowerActionsPercentageOfMachines field if non-nil, zero value otherwise.

### GetMaxPowerActionsPercentageOfMachinesOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetMaxPowerActionsPercentageOfMachinesOk() (*int32, bool)`

GetMaxPowerActionsPercentageOfMachinesOk returns a tuple with the MaxPowerActionsPercentageOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorAzureConnectionDetailRequestModel) SetMaxPowerActionsPercentageOfMachines(v int32)`

SetMaxPowerActionsPercentageOfMachines sets MaxPowerActionsPercentageOfMachines field to given value.

### HasMaxPowerActionsPercentageOfMachines

`func (o *HypervisorAzureConnectionDetailRequestModel) HasMaxPowerActionsPercentageOfMachines() bool`

HasMaxPowerActionsPercentageOfMachines returns a boolean if a field has been set.

### GetConnectionOptions

`func (o *HypervisorAzureConnectionDetailRequestModel) GetConnectionOptions() string`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetConnectionOptionsOk() (*string, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *HypervisorAzureConnectionDetailRequestModel) SetConnectionOptions(v string)`

SetConnectionOptions sets ConnectionOptions field to given value.

### HasConnectionOptions

`func (o *HypervisorAzureConnectionDetailRequestModel) HasConnectionOptions() bool`

HasConnectionOptions returns a boolean if a field has been set.

### GetZone

`func (o *HypervisorAzureConnectionDetailRequestModel) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorAzureConnectionDetailRequestModel) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *HypervisorAzureConnectionDetailRequestModel) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetMetadata

`func (o *HypervisorAzureConnectionDetailRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorAzureConnectionDetailRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorAzureConnectionDetailRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetApiKey

`func (o *HypervisorAzureConnectionDetailRequestModel) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HypervisorAzureConnectionDetailRequestModel) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *HypervisorAzureConnectionDetailRequestModel) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *HypervisorAzureConnectionDetailRequestModel) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *HypervisorAzureConnectionDetailRequestModel) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSecretKeyFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSecretKeyFormat() IdentityPasswordFormat`

GetSecretKeyFormat returns the SecretKeyFormat field if non-nil, zero value otherwise.

### GetSecretKeyFormatOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSecretKeyFormatOk() (*IdentityPasswordFormat, bool)`

GetSecretKeyFormatOk returns a tuple with the SecretKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) SetSecretKeyFormat(v IdentityPasswordFormat)`

SetSecretKeyFormat sets SecretKeyFormat field to given value.

### HasSecretKeyFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) HasSecretKeyFormat() bool`

HasSecretKeyFormat returns a boolean if a field has been set.

### GetRegion

`func (o *HypervisorAzureConnectionDetailRequestModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorAzureConnectionDetailRequestModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HypervisorAzureConnectionDetailRequestModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetApplicationId

`func (o *HypervisorAzureConnectionDetailRequestModel) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *HypervisorAzureConnectionDetailRequestModel) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationSecret

`func (o *HypervisorAzureConnectionDetailRequestModel) GetApplicationSecret() string`

GetApplicationSecret returns the ApplicationSecret field if non-nil, zero value otherwise.

### GetApplicationSecretOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetApplicationSecretOk() (*string, bool)`

GetApplicationSecretOk returns a tuple with the ApplicationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecret

`func (o *HypervisorAzureConnectionDetailRequestModel) SetApplicationSecret(v string)`

SetApplicationSecret sets ApplicationSecret field to given value.


### GetApplicationSecretFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) GetApplicationSecretFormat() IdentityPasswordFormat`

GetApplicationSecretFormat returns the ApplicationSecretFormat field if non-nil, zero value otherwise.

### GetApplicationSecretFormatOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetApplicationSecretFormatOk() (*IdentityPasswordFormat, bool)`

GetApplicationSecretFormatOk returns a tuple with the ApplicationSecretFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecretFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) SetApplicationSecretFormat(v IdentityPasswordFormat)`

SetApplicationSecretFormat sets ApplicationSecretFormat field to given value.

### HasApplicationSecretFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) HasApplicationSecretFormat() bool`

HasApplicationSecretFormat returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *HypervisorAzureConnectionDetailRequestModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetActiveDirectoryId

`func (o *HypervisorAzureConnectionDetailRequestModel) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *HypervisorAzureConnectionDetailRequestModel) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.


### GetEnvironment

`func (o *HypervisorAzureConnectionDetailRequestModel) GetEnvironment() AzureEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetEnvironmentOk() (*AzureEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *HypervisorAzureConnectionDetailRequestModel) SetEnvironment(v AzureEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *HypervisorAzureConnectionDetailRequestModel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetManagementEndpoint

`func (o *HypervisorAzureConnectionDetailRequestModel) GetManagementEndpoint() string`

GetManagementEndpoint returns the ManagementEndpoint field if non-nil, zero value otherwise.

### GetManagementEndpointOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetManagementEndpointOk() (*string, bool)`

GetManagementEndpointOk returns a tuple with the ManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEndpoint

`func (o *HypervisorAzureConnectionDetailRequestModel) SetManagementEndpoint(v string)`

SetManagementEndpoint sets ManagementEndpoint field to given value.

### HasManagementEndpoint

`func (o *HypervisorAzureConnectionDetailRequestModel) HasManagementEndpoint() bool`

HasManagementEndpoint returns a boolean if a field has been set.

### GetAuthenticationAuthority

`func (o *HypervisorAzureConnectionDetailRequestModel) GetAuthenticationAuthority() string`

GetAuthenticationAuthority returns the AuthenticationAuthority field if non-nil, zero value otherwise.

### GetAuthenticationAuthorityOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetAuthenticationAuthorityOk() (*string, bool)`

GetAuthenticationAuthorityOk returns a tuple with the AuthenticationAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAuthority

`func (o *HypervisorAzureConnectionDetailRequestModel) SetAuthenticationAuthority(v string)`

SetAuthenticationAuthority sets AuthenticationAuthority field to given value.

### HasAuthenticationAuthority

`func (o *HypervisorAzureConnectionDetailRequestModel) HasAuthenticationAuthority() bool`

HasAuthenticationAuthority returns a boolean if a field has been set.

### GetStorageSuffix

`func (o *HypervisorAzureConnectionDetailRequestModel) GetStorageSuffix() string`

GetStorageSuffix returns the StorageSuffix field if non-nil, zero value otherwise.

### GetStorageSuffixOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetStorageSuffixOk() (*string, bool)`

GetStorageSuffixOk returns a tuple with the StorageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSuffix

`func (o *HypervisorAzureConnectionDetailRequestModel) SetStorageSuffix(v string)`

SetStorageSuffix sets StorageSuffix field to given value.

### HasStorageSuffix

`func (o *HypervisorAzureConnectionDetailRequestModel) HasStorageSuffix() bool`

HasStorageSuffix returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorAzureConnectionDetailRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorAzureConnectionDetailRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorAzureConnectionDetailRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *HypervisorAzureConnectionDetailRequestModel) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *HypervisorAzureConnectionDetailRequestModel) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetServiceAccountCredentials

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountCredentials() string`

GetServiceAccountCredentials returns the ServiceAccountCredentials field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountCredentialsOk() (*string, bool)`

GetServiceAccountCredentialsOk returns a tuple with the ServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentials

`func (o *HypervisorAzureConnectionDetailRequestModel) SetServiceAccountCredentials(v string)`

SetServiceAccountCredentials sets ServiceAccountCredentials field to given value.

### HasServiceAccountCredentials

`func (o *HypervisorAzureConnectionDetailRequestModel) HasServiceAccountCredentials() bool`

HasServiceAccountCredentials returns a boolean if a field has been set.

### GetServiceAccountCredentialsFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountCredentialsFormat() IdentityPasswordFormat`

GetServiceAccountCredentialsFormat returns the ServiceAccountCredentialsFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsFormatOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountCredentialsFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialsFormatOk returns a tuple with the ServiceAccountCredentialsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialsFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) SetServiceAccountCredentialsFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialsFormat sets ServiceAccountCredentialsFormat field to given value.

### HasServiceAccountCredentialsFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) HasServiceAccountCredentialsFormat() bool`

HasServiceAccountCredentialsFormat returns a boolean if a field has been set.

### GetUserName

`func (o *HypervisorAzureConnectionDetailRequestModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HypervisorAzureConnectionDetailRequestModel) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *HypervisorAzureConnectionDetailRequestModel) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *HypervisorAzureConnectionDetailRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HypervisorAzureConnectionDetailRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HypervisorAzureConnectionDetailRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.

### GetAddresses

`func (o *HypervisorAzureConnectionDetailRequestModel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorAzureConnectionDetailRequestModel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *HypervisorAzureConnectionDetailRequestModel) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetPluginId

`func (o *HypervisorAzureConnectionDetailRequestModel) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorAzureConnectionDetailRequestModel) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *HypervisorAzureConnectionDetailRequestModel) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetSslThumbprints

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSslThumbprints() []string`

GetSslThumbprints returns the SslThumbprints field if non-nil, zero value otherwise.

### GetSslThumbprintsOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSslThumbprintsOk() (*[]string, bool)`

GetSslThumbprintsOk returns a tuple with the SslThumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprints

`func (o *HypervisorAzureConnectionDetailRequestModel) SetSslThumbprints(v []string)`

SetSslThumbprints sets SslThumbprints field to given value.

### HasSslThumbprints

`func (o *HypervisorAzureConnectionDetailRequestModel) HasSslThumbprints() bool`

HasSslThumbprints returns a boolean if a field has been set.

### GetSccmWakeUpProxy

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSccmWakeUpProxy() bool`

GetSccmWakeUpProxy returns the SccmWakeUpProxy field if non-nil, zero value otherwise.

### GetSccmWakeUpProxyOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetSccmWakeUpProxyOk() (*bool, bool)`

GetSccmWakeUpProxyOk returns a tuple with the SccmWakeUpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmWakeUpProxy

`func (o *HypervisorAzureConnectionDetailRequestModel) SetSccmWakeUpProxy(v bool)`

SetSccmWakeUpProxy sets SccmWakeUpProxy field to given value.

### HasSccmWakeUpProxy

`func (o *HypervisorAzureConnectionDetailRequestModel) HasSccmWakeUpProxy() bool`

HasSccmWakeUpProxy returns a boolean if a field has been set.

### GetWakeOnLanPackets

`func (o *HypervisorAzureConnectionDetailRequestModel) GetWakeOnLanPackets() WakeOnLanTransmission`

GetWakeOnLanPackets returns the WakeOnLanPackets field if non-nil, zero value otherwise.

### GetWakeOnLanPacketsOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetWakeOnLanPacketsOk() (*WakeOnLanTransmission, bool)`

GetWakeOnLanPacketsOk returns a tuple with the WakeOnLanPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanPackets

`func (o *HypervisorAzureConnectionDetailRequestModel) SetWakeOnLanPackets(v WakeOnLanTransmission)`

SetWakeOnLanPackets sets WakeOnLanPackets field to given value.

### HasWakeOnLanPackets

`func (o *HypervisorAzureConnectionDetailRequestModel) HasWakeOnLanPackets() bool`

HasWakeOnLanPackets returns a boolean if a field has been set.

### GetTenancyOcid

`func (o *HypervisorAzureConnectionDetailRequestModel) GetTenancyOcid() string`

GetTenancyOcid returns the TenancyOcid field if non-nil, zero value otherwise.

### GetTenancyOcidOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetTenancyOcidOk() (*string, bool)`

GetTenancyOcidOk returns a tuple with the TenancyOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyOcid

`func (o *HypervisorAzureConnectionDetailRequestModel) SetTenancyOcid(v string)`

SetTenancyOcid sets TenancyOcid field to given value.

### HasTenancyOcid

`func (o *HypervisorAzureConnectionDetailRequestModel) HasTenancyOcid() bool`

HasTenancyOcid returns a boolean if a field has been set.

### GetServiceAccountCredential

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountCredential() string`

GetServiceAccountCredential returns the ServiceAccountCredential field if non-nil, zero value otherwise.

### GetServiceAccountCredentialOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountCredentialOk() (*string, bool)`

GetServiceAccountCredentialOk returns a tuple with the ServiceAccountCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredential

`func (o *HypervisorAzureConnectionDetailRequestModel) SetServiceAccountCredential(v string)`

SetServiceAccountCredential sets ServiceAccountCredential field to given value.

### HasServiceAccountCredential

`func (o *HypervisorAzureConnectionDetailRequestModel) HasServiceAccountCredential() bool`

HasServiceAccountCredential returns a boolean if a field has been set.

### GetServiceAccountCredentialFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountCredentialFormat() IdentityPasswordFormat`

GetServiceAccountCredentialFormat returns the ServiceAccountCredentialFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialFormatOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountCredentialFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialFormatOk returns a tuple with the ServiceAccountCredentialFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) SetServiceAccountCredentialFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialFormat sets ServiceAccountCredentialFormat field to given value.

### HasServiceAccountCredentialFormat

`func (o *HypervisorAzureConnectionDetailRequestModel) HasServiceAccountCredentialFormat() bool`

HasServiceAccountCredentialFormat returns a boolean if a field has been set.

### GetServiceAccountFingerprint

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountFingerprint() string`

GetServiceAccountFingerprint returns the ServiceAccountFingerprint field if non-nil, zero value otherwise.

### GetServiceAccountFingerprintOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetServiceAccountFingerprintOk() (*string, bool)`

GetServiceAccountFingerprintOk returns a tuple with the ServiceAccountFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountFingerprint

`func (o *HypervisorAzureConnectionDetailRequestModel) SetServiceAccountFingerprint(v string)`

SetServiceAccountFingerprint sets ServiceAccountFingerprint field to given value.

### HasServiceAccountFingerprint

`func (o *HypervisorAzureConnectionDetailRequestModel) HasServiceAccountFingerprint() bool`

HasServiceAccountFingerprint returns a boolean if a field has been set.

### GetOciEnvironment

`func (o *HypervisorAzureConnectionDetailRequestModel) GetOciEnvironment() OciEnvironment`

GetOciEnvironment returns the OciEnvironment field if non-nil, zero value otherwise.

### GetOciEnvironmentOk

`func (o *HypervisorAzureConnectionDetailRequestModel) GetOciEnvironmentOk() (*OciEnvironment, bool)`

GetOciEnvironmentOk returns a tuple with the OciEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciEnvironment

`func (o *HypervisorAzureConnectionDetailRequestModel) SetOciEnvironment(v OciEnvironment)`

SetOciEnvironment sets OciEnvironment field to given value.

### HasOciEnvironment

`func (o *HypervisorAzureConnectionDetailRequestModel) HasOciEnvironment() bool`

HasOciEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


