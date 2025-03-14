# HypervisorGCPConnectionDetailRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the hypervisor to create.  Required. | 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**Scopes** | Pointer to **[]string** | Administrative scopes which the newly created hypervisor will be a part of. | [optional] 
**Tenants** | Pointer to **[]string** | Tenant(s) to associate the hypervisor with. | [optional] 
**MaxAbsoluteActiveActions** | Pointer to **NullableInt32** | Maximum number of actions that can execute in parallel on the hypervisor.  Optional; if not specified, an appropriate default is selected based on the ConnectionType. | [optional] 
**MaxAbsoluteNewActionsPerMinute** | Pointer to **NullableInt32** | Maximum number of actions that can be started on the hypervisor per-minute.  Optional; if not specified, an appropriate default is selected based on the ConnectionType. | [optional] 
**MaxPowerActionsPercentageOfMachines** | Pointer to **NullableInt32** | Maximum percentage of machines on the hypervisor which can have their power state changed simultaneously.  Optional; if not specified, an appropriate default is selected based on the ConnectionType. | [optional] 
**ConnectionOptions** | Pointer to **NullableString** | Connection options to use for the hypervisor.  Optional. | [optional] 
**Zone** | Pointer to **NullableString** | Zone the hypervisor is associated with.  Optional. If not specified, the hypervisor is associated with the primary zone. See PrimaryZone. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata to use for the hypervisor. Optional. | [optional] 
**ApiKey** | Pointer to **string** | The API key used to authenticate with the AWS APIs.  Required. | [optional] 
**SecretKey** | Pointer to **string** | The secret key used to authenticate with the AWS APIs.  Required. Must be specified in the format indicated by SecretKeyFormat. | [optional] 
**SecretKeyFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**Region** | Pointer to **NullableString** | AWS region to connect to.  Optional.  If not specified, will connect to the global AWS APIs.  This can be used to discover the regions available within AWS.  Access to all other AWS resources requires the region to be set explicitly. | [optional] 
**Address** | Pointer to **NullableString** | Custom AWS Address. | [optional] 
**CustomProperties** | Pointer to **NullableString** | The properties of host connection that are specific to the target hosting infrastructure. | [optional] 
**ApplicationId** | Pointer to **NullableString** | Application ID of the service principal used to access the Azure APIs. Required for AuthenticationMode in CustomProperties is AppClientSecret or UserAssignedManagedIdentity. | [optional] 
**ApplicationSecret** | Pointer to **NullableString** | The Application Secret of the service principal used to access the Azure APIs. Required for AuthenticationMode in CustomProperties is AppClientSecret. Must be specified in the format indicated by ApplicationSecretFormat. | [optional] 
**ApplicationSecretFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | Azure subscription ID.  Required. | [optional] 
**ActiveDirectoryId** | Pointer to **string** | Azure active directory ID.  Required. | [optional] 
**Environment** | Pointer to [**AzureEnvironment**](AzureEnvironment.md) |  | [optional] 
**ManagementEndpoint** | Pointer to **NullableString** | Azure management endpoint.  Required if Environment is Custom. Optional otherwise. | [optional] 
**AuthenticationAuthority** | Pointer to **NullableString** | Azure authentication authority.  Required if Environment is Custom. Optional otherwise. | [optional] 
**StorageSuffix** | Pointer to **NullableString** | Azure storage suffix.  Required if Environment is Custom. Optional otherwise. | [optional] 
**ServiceAccountId** | **string** | The service account ID used to access the Google Cloud APIs. Required. | 
**ServiceAccountCredentials** | **string** | the JSON-encoded service account credentials used to access the Google Cloud APIs.  Required. Must be specified in the format indicated by ServiceAccountCredentialsFormat. | 
**ServiceAccountCredentialsFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**UserName** | Pointer to **string** | Hypervisor user name.  Required. | [optional] 
**Password** | Pointer to **string** | Hypervisor password.  Required. Must be specified in the format indicated by PasswordFormat. | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**Addresses** | Pointer to **[]string** | Hypervisor address(es).  At least one is required. | [optional] 
**PluginId** | Pointer to **NullableString** | Custom hypervisor plugin ID.  Required if the ConnectionType is Custom.  Otherwise must not be specified. | [optional] 
**SslThumbprints** | Pointer to **[]string** | SSL certificate thumbprints to consider acceptable for this connection.  Optional.  If not specified, and the hypervisor uses SSL for its connection, the SSL certificate&#39;s root certification authority and any intermediate certificates must be trusted. | [optional] 
**SccmWakeUpProxy** | Pointer to **NullableBool** | Indicates whether Microsoft System Center Configuration Manager 2012 SP1 Wake-up Proxy is used for power management.  Optional; default is &#x60;false&#x60;. | [optional] 
**WakeOnLanPackets** | Pointer to [**WakeOnLanTransmission**](WakeOnLanTransmission.md) |  | [optional] 
**TenancyOcid** | Pointer to **string** | Oracle Cloud Infrastructure tenancy to connect to. Required. | [optional] 
**ServiceAccountCredential** | Pointer to **string** | The private key string to access the Oracle Cloud Infrastructure APIs. Required. Must be specified in the format indicated by ServiceAccountCredentialFormat. | [optional] 
**ServiceAccountCredentialFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**ServiceAccountFingerprint** | Pointer to **string** | The fingerprint of the public key associate with the ServiceAccountCredential. | [optional] 
**OciEnvironment** | Pointer to [**OciEnvironment**](OciEnvironment.md) |  | [optional] 

## Methods

### NewHypervisorGCPConnectionDetailRequestModel

`func NewHypervisorGCPConnectionDetailRequestModel(name string, connectionType HypervisorConnectionType, serviceAccountId string, serviceAccountCredentials string, ) *HypervisorGCPConnectionDetailRequestModel`

NewHypervisorGCPConnectionDetailRequestModel instantiates a new HypervisorGCPConnectionDetailRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorGCPConnectionDetailRequestModelWithDefaults

`func NewHypervisorGCPConnectionDetailRequestModelWithDefaults() *HypervisorGCPConnectionDetailRequestModel`

NewHypervisorGCPConnectionDetailRequestModelWithDefaults instantiates a new HypervisorGCPConnectionDetailRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HypervisorGCPConnectionDetailRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorGCPConnectionDetailRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetConnectionType

`func (o *HypervisorGCPConnectionDetailRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorGCPConnectionDetailRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetScopes

`func (o *HypervisorGCPConnectionDetailRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorGCPConnectionDetailRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *HypervisorGCPConnectionDetailRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *HypervisorGCPConnectionDetailRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorGCPConnectionDetailRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorGCPConnectionDetailRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetMaxAbsoluteActiveActions

`func (o *HypervisorGCPConnectionDetailRequestModel) GetMaxAbsoluteActiveActions() int32`

GetMaxAbsoluteActiveActions returns the MaxAbsoluteActiveActions field if non-nil, zero value otherwise.

### GetMaxAbsoluteActiveActionsOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetMaxAbsoluteActiveActionsOk() (*int32, bool)`

GetMaxAbsoluteActiveActionsOk returns a tuple with the MaxAbsoluteActiveActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteActiveActions

`func (o *HypervisorGCPConnectionDetailRequestModel) SetMaxAbsoluteActiveActions(v int32)`

SetMaxAbsoluteActiveActions sets MaxAbsoluteActiveActions field to given value.

### HasMaxAbsoluteActiveActions

`func (o *HypervisorGCPConnectionDetailRequestModel) HasMaxAbsoluteActiveActions() bool`

HasMaxAbsoluteActiveActions returns a boolean if a field has been set.

### SetMaxAbsoluteActiveActionsNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetMaxAbsoluteActiveActionsNil(b bool)`

 SetMaxAbsoluteActiveActionsNil sets the value for MaxAbsoluteActiveActions to be an explicit nil

### UnsetMaxAbsoluteActiveActions
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetMaxAbsoluteActiveActions()`

UnsetMaxAbsoluteActiveActions ensures that no value is present for MaxAbsoluteActiveActions, not even an explicit nil
### GetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorGCPConnectionDetailRequestModel) GetMaxAbsoluteNewActionsPerMinute() int32`

GetMaxAbsoluteNewActionsPerMinute returns the MaxAbsoluteNewActionsPerMinute field if non-nil, zero value otherwise.

### GetMaxAbsoluteNewActionsPerMinuteOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetMaxAbsoluteNewActionsPerMinuteOk() (*int32, bool)`

GetMaxAbsoluteNewActionsPerMinuteOk returns a tuple with the MaxAbsoluteNewActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorGCPConnectionDetailRequestModel) SetMaxAbsoluteNewActionsPerMinute(v int32)`

SetMaxAbsoluteNewActionsPerMinute sets MaxAbsoluteNewActionsPerMinute field to given value.

### HasMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorGCPConnectionDetailRequestModel) HasMaxAbsoluteNewActionsPerMinute() bool`

HasMaxAbsoluteNewActionsPerMinute returns a boolean if a field has been set.

### SetMaxAbsoluteNewActionsPerMinuteNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetMaxAbsoluteNewActionsPerMinuteNil(b bool)`

 SetMaxAbsoluteNewActionsPerMinuteNil sets the value for MaxAbsoluteNewActionsPerMinute to be an explicit nil

### UnsetMaxAbsoluteNewActionsPerMinute
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetMaxAbsoluteNewActionsPerMinute()`

UnsetMaxAbsoluteNewActionsPerMinute ensures that no value is present for MaxAbsoluteNewActionsPerMinute, not even an explicit nil
### GetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorGCPConnectionDetailRequestModel) GetMaxPowerActionsPercentageOfMachines() int32`

GetMaxPowerActionsPercentageOfMachines returns the MaxPowerActionsPercentageOfMachines field if non-nil, zero value otherwise.

### GetMaxPowerActionsPercentageOfMachinesOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetMaxPowerActionsPercentageOfMachinesOk() (*int32, bool)`

GetMaxPowerActionsPercentageOfMachinesOk returns a tuple with the MaxPowerActionsPercentageOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorGCPConnectionDetailRequestModel) SetMaxPowerActionsPercentageOfMachines(v int32)`

SetMaxPowerActionsPercentageOfMachines sets MaxPowerActionsPercentageOfMachines field to given value.

### HasMaxPowerActionsPercentageOfMachines

`func (o *HypervisorGCPConnectionDetailRequestModel) HasMaxPowerActionsPercentageOfMachines() bool`

HasMaxPowerActionsPercentageOfMachines returns a boolean if a field has been set.

### SetMaxPowerActionsPercentageOfMachinesNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetMaxPowerActionsPercentageOfMachinesNil(b bool)`

 SetMaxPowerActionsPercentageOfMachinesNil sets the value for MaxPowerActionsPercentageOfMachines to be an explicit nil

### UnsetMaxPowerActionsPercentageOfMachines
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetMaxPowerActionsPercentageOfMachines()`

UnsetMaxPowerActionsPercentageOfMachines ensures that no value is present for MaxPowerActionsPercentageOfMachines, not even an explicit nil
### GetConnectionOptions

`func (o *HypervisorGCPConnectionDetailRequestModel) GetConnectionOptions() string`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetConnectionOptionsOk() (*string, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *HypervisorGCPConnectionDetailRequestModel) SetConnectionOptions(v string)`

SetConnectionOptions sets ConnectionOptions field to given value.

### HasConnectionOptions

`func (o *HypervisorGCPConnectionDetailRequestModel) HasConnectionOptions() bool`

HasConnectionOptions returns a boolean if a field has been set.

### SetConnectionOptionsNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetConnectionOptionsNil(b bool)`

 SetConnectionOptionsNil sets the value for ConnectionOptions to be an explicit nil

### UnsetConnectionOptions
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetConnectionOptions()`

UnsetConnectionOptions ensures that no value is present for ConnectionOptions, not even an explicit nil
### GetZone

`func (o *HypervisorGCPConnectionDetailRequestModel) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorGCPConnectionDetailRequestModel) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *HypervisorGCPConnectionDetailRequestModel) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetMetadata

`func (o *HypervisorGCPConnectionDetailRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorGCPConnectionDetailRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorGCPConnectionDetailRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetApiKey

`func (o *HypervisorGCPConnectionDetailRequestModel) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HypervisorGCPConnectionDetailRequestModel) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *HypervisorGCPConnectionDetailRequestModel) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *HypervisorGCPConnectionDetailRequestModel) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *HypervisorGCPConnectionDetailRequestModel) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSecretKeyFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSecretKeyFormat() IdentityPasswordFormat`

GetSecretKeyFormat returns the SecretKeyFormat field if non-nil, zero value otherwise.

### GetSecretKeyFormatOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSecretKeyFormatOk() (*IdentityPasswordFormat, bool)`

GetSecretKeyFormatOk returns a tuple with the SecretKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) SetSecretKeyFormat(v IdentityPasswordFormat)`

SetSecretKeyFormat sets SecretKeyFormat field to given value.

### HasSecretKeyFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) HasSecretKeyFormat() bool`

HasSecretKeyFormat returns a boolean if a field has been set.

### GetRegion

`func (o *HypervisorGCPConnectionDetailRequestModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorGCPConnectionDetailRequestModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HypervisorGCPConnectionDetailRequestModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetAddress

`func (o *HypervisorGCPConnectionDetailRequestModel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HypervisorGCPConnectionDetailRequestModel) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *HypervisorGCPConnectionDetailRequestModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCustomProperties

`func (o *HypervisorGCPConnectionDetailRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorGCPConnectionDetailRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorGCPConnectionDetailRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetApplicationId

`func (o *HypervisorGCPConnectionDetailRequestModel) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *HypervisorGCPConnectionDetailRequestModel) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *HypervisorGCPConnectionDetailRequestModel) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetApplicationSecret

`func (o *HypervisorGCPConnectionDetailRequestModel) GetApplicationSecret() string`

GetApplicationSecret returns the ApplicationSecret field if non-nil, zero value otherwise.

### GetApplicationSecretOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetApplicationSecretOk() (*string, bool)`

GetApplicationSecretOk returns a tuple with the ApplicationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecret

`func (o *HypervisorGCPConnectionDetailRequestModel) SetApplicationSecret(v string)`

SetApplicationSecret sets ApplicationSecret field to given value.

### HasApplicationSecret

`func (o *HypervisorGCPConnectionDetailRequestModel) HasApplicationSecret() bool`

HasApplicationSecret returns a boolean if a field has been set.

### SetApplicationSecretNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetApplicationSecretNil(b bool)`

 SetApplicationSecretNil sets the value for ApplicationSecret to be an explicit nil

### UnsetApplicationSecret
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetApplicationSecret()`

UnsetApplicationSecret ensures that no value is present for ApplicationSecret, not even an explicit nil
### GetApplicationSecretFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) GetApplicationSecretFormat() IdentityPasswordFormat`

GetApplicationSecretFormat returns the ApplicationSecretFormat field if non-nil, zero value otherwise.

### GetApplicationSecretFormatOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetApplicationSecretFormatOk() (*IdentityPasswordFormat, bool)`

GetApplicationSecretFormatOk returns a tuple with the ApplicationSecretFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecretFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) SetApplicationSecretFormat(v IdentityPasswordFormat)`

SetApplicationSecretFormat sets ApplicationSecretFormat field to given value.

### HasApplicationSecretFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) HasApplicationSecretFormat() bool`

HasApplicationSecretFormat returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *HypervisorGCPConnectionDetailRequestModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *HypervisorGCPConnectionDetailRequestModel) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetActiveDirectoryId

`func (o *HypervisorGCPConnectionDetailRequestModel) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *HypervisorGCPConnectionDetailRequestModel) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.

### HasActiveDirectoryId

`func (o *HypervisorGCPConnectionDetailRequestModel) HasActiveDirectoryId() bool`

HasActiveDirectoryId returns a boolean if a field has been set.

### GetEnvironment

`func (o *HypervisorGCPConnectionDetailRequestModel) GetEnvironment() AzureEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetEnvironmentOk() (*AzureEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *HypervisorGCPConnectionDetailRequestModel) SetEnvironment(v AzureEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *HypervisorGCPConnectionDetailRequestModel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetManagementEndpoint

`func (o *HypervisorGCPConnectionDetailRequestModel) GetManagementEndpoint() string`

GetManagementEndpoint returns the ManagementEndpoint field if non-nil, zero value otherwise.

### GetManagementEndpointOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetManagementEndpointOk() (*string, bool)`

GetManagementEndpointOk returns a tuple with the ManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEndpoint

`func (o *HypervisorGCPConnectionDetailRequestModel) SetManagementEndpoint(v string)`

SetManagementEndpoint sets ManagementEndpoint field to given value.

### HasManagementEndpoint

`func (o *HypervisorGCPConnectionDetailRequestModel) HasManagementEndpoint() bool`

HasManagementEndpoint returns a boolean if a field has been set.

### SetManagementEndpointNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetManagementEndpointNil(b bool)`

 SetManagementEndpointNil sets the value for ManagementEndpoint to be an explicit nil

### UnsetManagementEndpoint
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetManagementEndpoint()`

UnsetManagementEndpoint ensures that no value is present for ManagementEndpoint, not even an explicit nil
### GetAuthenticationAuthority

`func (o *HypervisorGCPConnectionDetailRequestModel) GetAuthenticationAuthority() string`

GetAuthenticationAuthority returns the AuthenticationAuthority field if non-nil, zero value otherwise.

### GetAuthenticationAuthorityOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetAuthenticationAuthorityOk() (*string, bool)`

GetAuthenticationAuthorityOk returns a tuple with the AuthenticationAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAuthority

`func (o *HypervisorGCPConnectionDetailRequestModel) SetAuthenticationAuthority(v string)`

SetAuthenticationAuthority sets AuthenticationAuthority field to given value.

### HasAuthenticationAuthority

`func (o *HypervisorGCPConnectionDetailRequestModel) HasAuthenticationAuthority() bool`

HasAuthenticationAuthority returns a boolean if a field has been set.

### SetAuthenticationAuthorityNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetAuthenticationAuthorityNil(b bool)`

 SetAuthenticationAuthorityNil sets the value for AuthenticationAuthority to be an explicit nil

### UnsetAuthenticationAuthority
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetAuthenticationAuthority()`

UnsetAuthenticationAuthority ensures that no value is present for AuthenticationAuthority, not even an explicit nil
### GetStorageSuffix

`func (o *HypervisorGCPConnectionDetailRequestModel) GetStorageSuffix() string`

GetStorageSuffix returns the StorageSuffix field if non-nil, zero value otherwise.

### GetStorageSuffixOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetStorageSuffixOk() (*string, bool)`

GetStorageSuffixOk returns a tuple with the StorageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSuffix

`func (o *HypervisorGCPConnectionDetailRequestModel) SetStorageSuffix(v string)`

SetStorageSuffix sets StorageSuffix field to given value.

### HasStorageSuffix

`func (o *HypervisorGCPConnectionDetailRequestModel) HasStorageSuffix() bool`

HasStorageSuffix returns a boolean if a field has been set.

### SetStorageSuffixNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetStorageSuffixNil(b bool)`

 SetStorageSuffixNil sets the value for StorageSuffix to be an explicit nil

### UnsetStorageSuffix
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetStorageSuffix()`

UnsetStorageSuffix ensures that no value is present for StorageSuffix, not even an explicit nil
### GetServiceAccountId

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *HypervisorGCPConnectionDetailRequestModel) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetServiceAccountCredentials

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountCredentials() string`

GetServiceAccountCredentials returns the ServiceAccountCredentials field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountCredentialsOk() (*string, bool)`

GetServiceAccountCredentialsOk returns a tuple with the ServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentials

`func (o *HypervisorGCPConnectionDetailRequestModel) SetServiceAccountCredentials(v string)`

SetServiceAccountCredentials sets ServiceAccountCredentials field to given value.


### GetServiceAccountCredentialsFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountCredentialsFormat() IdentityPasswordFormat`

GetServiceAccountCredentialsFormat returns the ServiceAccountCredentialsFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsFormatOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountCredentialsFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialsFormatOk returns a tuple with the ServiceAccountCredentialsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialsFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) SetServiceAccountCredentialsFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialsFormat sets ServiceAccountCredentialsFormat field to given value.

### HasServiceAccountCredentialsFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) HasServiceAccountCredentialsFormat() bool`

HasServiceAccountCredentialsFormat returns a boolean if a field has been set.

### GetUserName

`func (o *HypervisorGCPConnectionDetailRequestModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HypervisorGCPConnectionDetailRequestModel) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *HypervisorGCPConnectionDetailRequestModel) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *HypervisorGCPConnectionDetailRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HypervisorGCPConnectionDetailRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HypervisorGCPConnectionDetailRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.

### GetAddresses

`func (o *HypervisorGCPConnectionDetailRequestModel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorGCPConnectionDetailRequestModel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *HypervisorGCPConnectionDetailRequestModel) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetPluginId

`func (o *HypervisorGCPConnectionDetailRequestModel) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorGCPConnectionDetailRequestModel) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *HypervisorGCPConnectionDetailRequestModel) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### SetPluginIdNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetPluginIdNil(b bool)`

 SetPluginIdNil sets the value for PluginId to be an explicit nil

### UnsetPluginId
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetPluginId()`

UnsetPluginId ensures that no value is present for PluginId, not even an explicit nil
### GetSslThumbprints

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSslThumbprints() []string`

GetSslThumbprints returns the SslThumbprints field if non-nil, zero value otherwise.

### GetSslThumbprintsOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSslThumbprintsOk() (*[]string, bool)`

GetSslThumbprintsOk returns a tuple with the SslThumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprints

`func (o *HypervisorGCPConnectionDetailRequestModel) SetSslThumbprints(v []string)`

SetSslThumbprints sets SslThumbprints field to given value.

### HasSslThumbprints

`func (o *HypervisorGCPConnectionDetailRequestModel) HasSslThumbprints() bool`

HasSslThumbprints returns a boolean if a field has been set.

### SetSslThumbprintsNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetSslThumbprintsNil(b bool)`

 SetSslThumbprintsNil sets the value for SslThumbprints to be an explicit nil

### UnsetSslThumbprints
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetSslThumbprints()`

UnsetSslThumbprints ensures that no value is present for SslThumbprints, not even an explicit nil
### GetSccmWakeUpProxy

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSccmWakeUpProxy() bool`

GetSccmWakeUpProxy returns the SccmWakeUpProxy field if non-nil, zero value otherwise.

### GetSccmWakeUpProxyOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetSccmWakeUpProxyOk() (*bool, bool)`

GetSccmWakeUpProxyOk returns a tuple with the SccmWakeUpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmWakeUpProxy

`func (o *HypervisorGCPConnectionDetailRequestModel) SetSccmWakeUpProxy(v bool)`

SetSccmWakeUpProxy sets SccmWakeUpProxy field to given value.

### HasSccmWakeUpProxy

`func (o *HypervisorGCPConnectionDetailRequestModel) HasSccmWakeUpProxy() bool`

HasSccmWakeUpProxy returns a boolean if a field has been set.

### SetSccmWakeUpProxyNil

`func (o *HypervisorGCPConnectionDetailRequestModel) SetSccmWakeUpProxyNil(b bool)`

 SetSccmWakeUpProxyNil sets the value for SccmWakeUpProxy to be an explicit nil

### UnsetSccmWakeUpProxy
`func (o *HypervisorGCPConnectionDetailRequestModel) UnsetSccmWakeUpProxy()`

UnsetSccmWakeUpProxy ensures that no value is present for SccmWakeUpProxy, not even an explicit nil
### GetWakeOnLanPackets

`func (o *HypervisorGCPConnectionDetailRequestModel) GetWakeOnLanPackets() WakeOnLanTransmission`

GetWakeOnLanPackets returns the WakeOnLanPackets field if non-nil, zero value otherwise.

### GetWakeOnLanPacketsOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetWakeOnLanPacketsOk() (*WakeOnLanTransmission, bool)`

GetWakeOnLanPacketsOk returns a tuple with the WakeOnLanPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanPackets

`func (o *HypervisorGCPConnectionDetailRequestModel) SetWakeOnLanPackets(v WakeOnLanTransmission)`

SetWakeOnLanPackets sets WakeOnLanPackets field to given value.

### HasWakeOnLanPackets

`func (o *HypervisorGCPConnectionDetailRequestModel) HasWakeOnLanPackets() bool`

HasWakeOnLanPackets returns a boolean if a field has been set.

### GetTenancyOcid

`func (o *HypervisorGCPConnectionDetailRequestModel) GetTenancyOcid() string`

GetTenancyOcid returns the TenancyOcid field if non-nil, zero value otherwise.

### GetTenancyOcidOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetTenancyOcidOk() (*string, bool)`

GetTenancyOcidOk returns a tuple with the TenancyOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyOcid

`func (o *HypervisorGCPConnectionDetailRequestModel) SetTenancyOcid(v string)`

SetTenancyOcid sets TenancyOcid field to given value.

### HasTenancyOcid

`func (o *HypervisorGCPConnectionDetailRequestModel) HasTenancyOcid() bool`

HasTenancyOcid returns a boolean if a field has been set.

### GetServiceAccountCredential

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountCredential() string`

GetServiceAccountCredential returns the ServiceAccountCredential field if non-nil, zero value otherwise.

### GetServiceAccountCredentialOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountCredentialOk() (*string, bool)`

GetServiceAccountCredentialOk returns a tuple with the ServiceAccountCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredential

`func (o *HypervisorGCPConnectionDetailRequestModel) SetServiceAccountCredential(v string)`

SetServiceAccountCredential sets ServiceAccountCredential field to given value.

### HasServiceAccountCredential

`func (o *HypervisorGCPConnectionDetailRequestModel) HasServiceAccountCredential() bool`

HasServiceAccountCredential returns a boolean if a field has been set.

### GetServiceAccountCredentialFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountCredentialFormat() IdentityPasswordFormat`

GetServiceAccountCredentialFormat returns the ServiceAccountCredentialFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialFormatOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountCredentialFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialFormatOk returns a tuple with the ServiceAccountCredentialFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) SetServiceAccountCredentialFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialFormat sets ServiceAccountCredentialFormat field to given value.

### HasServiceAccountCredentialFormat

`func (o *HypervisorGCPConnectionDetailRequestModel) HasServiceAccountCredentialFormat() bool`

HasServiceAccountCredentialFormat returns a boolean if a field has been set.

### GetServiceAccountFingerprint

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountFingerprint() string`

GetServiceAccountFingerprint returns the ServiceAccountFingerprint field if non-nil, zero value otherwise.

### GetServiceAccountFingerprintOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetServiceAccountFingerprintOk() (*string, bool)`

GetServiceAccountFingerprintOk returns a tuple with the ServiceAccountFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountFingerprint

`func (o *HypervisorGCPConnectionDetailRequestModel) SetServiceAccountFingerprint(v string)`

SetServiceAccountFingerprint sets ServiceAccountFingerprint field to given value.

### HasServiceAccountFingerprint

`func (o *HypervisorGCPConnectionDetailRequestModel) HasServiceAccountFingerprint() bool`

HasServiceAccountFingerprint returns a boolean if a field has been set.

### GetOciEnvironment

`func (o *HypervisorGCPConnectionDetailRequestModel) GetOciEnvironment() OciEnvironment`

GetOciEnvironment returns the OciEnvironment field if non-nil, zero value otherwise.

### GetOciEnvironmentOk

`func (o *HypervisorGCPConnectionDetailRequestModel) GetOciEnvironmentOk() (*OciEnvironment, bool)`

GetOciEnvironmentOk returns a tuple with the OciEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciEnvironment

`func (o *HypervisorGCPConnectionDetailRequestModel) SetOciEnvironment(v OciEnvironment)`

SetOciEnvironment sets OciEnvironment field to given value.

### HasOciEnvironment

`func (o *HypervisorGCPConnectionDetailRequestModel) HasOciEnvironment() bool`

HasOciEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


