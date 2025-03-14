# HypervisorOpenShiftConnectionDetailRequestModel

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
**ServiceAccountId** | Pointer to **string** | The service account ID used to access the Google Cloud APIs. Required. | [optional] 
**ServiceAccountCredentials** | Pointer to **string** | the JSON-encoded service account credentials used to access the Google Cloud APIs.  Required. Must be specified in the format indicated by ServiceAccountCredentialsFormat. | [optional] 
**ServiceAccountCredentialsFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**UserName** | **string** | Hypervisor user name.  Required. | 
**Password** | **string** | Hypervisor password.  Required. Must be specified in the format indicated by PasswordFormat. | 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**Addresses** | **[]string** | Hypervisor address(es).  At least one is required. | 
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

### NewHypervisorOpenShiftConnectionDetailRequestModel

`func NewHypervisorOpenShiftConnectionDetailRequestModel(name string, connectionType HypervisorConnectionType, userName string, password string, addresses []string, ) *HypervisorOpenShiftConnectionDetailRequestModel`

NewHypervisorOpenShiftConnectionDetailRequestModel instantiates a new HypervisorOpenShiftConnectionDetailRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorOpenShiftConnectionDetailRequestModelWithDefaults

`func NewHypervisorOpenShiftConnectionDetailRequestModelWithDefaults() *HypervisorOpenShiftConnectionDetailRequestModel`

NewHypervisorOpenShiftConnectionDetailRequestModelWithDefaults instantiates a new HypervisorOpenShiftConnectionDetailRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetConnectionType

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetScopes

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetMaxAbsoluteActiveActions

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetMaxAbsoluteActiveActions() int32`

GetMaxAbsoluteActiveActions returns the MaxAbsoluteActiveActions field if non-nil, zero value otherwise.

### GetMaxAbsoluteActiveActionsOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetMaxAbsoluteActiveActionsOk() (*int32, bool)`

GetMaxAbsoluteActiveActionsOk returns a tuple with the MaxAbsoluteActiveActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteActiveActions

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetMaxAbsoluteActiveActions(v int32)`

SetMaxAbsoluteActiveActions sets MaxAbsoluteActiveActions field to given value.

### HasMaxAbsoluteActiveActions

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasMaxAbsoluteActiveActions() bool`

HasMaxAbsoluteActiveActions returns a boolean if a field has been set.

### SetMaxAbsoluteActiveActionsNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetMaxAbsoluteActiveActionsNil(b bool)`

 SetMaxAbsoluteActiveActionsNil sets the value for MaxAbsoluteActiveActions to be an explicit nil

### UnsetMaxAbsoluteActiveActions
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetMaxAbsoluteActiveActions()`

UnsetMaxAbsoluteActiveActions ensures that no value is present for MaxAbsoluteActiveActions, not even an explicit nil
### GetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetMaxAbsoluteNewActionsPerMinute() int32`

GetMaxAbsoluteNewActionsPerMinute returns the MaxAbsoluteNewActionsPerMinute field if non-nil, zero value otherwise.

### GetMaxAbsoluteNewActionsPerMinuteOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetMaxAbsoluteNewActionsPerMinuteOk() (*int32, bool)`

GetMaxAbsoluteNewActionsPerMinuteOk returns a tuple with the MaxAbsoluteNewActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetMaxAbsoluteNewActionsPerMinute(v int32)`

SetMaxAbsoluteNewActionsPerMinute sets MaxAbsoluteNewActionsPerMinute field to given value.

### HasMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasMaxAbsoluteNewActionsPerMinute() bool`

HasMaxAbsoluteNewActionsPerMinute returns a boolean if a field has been set.

### SetMaxAbsoluteNewActionsPerMinuteNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetMaxAbsoluteNewActionsPerMinuteNil(b bool)`

 SetMaxAbsoluteNewActionsPerMinuteNil sets the value for MaxAbsoluteNewActionsPerMinute to be an explicit nil

### UnsetMaxAbsoluteNewActionsPerMinute
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetMaxAbsoluteNewActionsPerMinute()`

UnsetMaxAbsoluteNewActionsPerMinute ensures that no value is present for MaxAbsoluteNewActionsPerMinute, not even an explicit nil
### GetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetMaxPowerActionsPercentageOfMachines() int32`

GetMaxPowerActionsPercentageOfMachines returns the MaxPowerActionsPercentageOfMachines field if non-nil, zero value otherwise.

### GetMaxPowerActionsPercentageOfMachinesOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetMaxPowerActionsPercentageOfMachinesOk() (*int32, bool)`

GetMaxPowerActionsPercentageOfMachinesOk returns a tuple with the MaxPowerActionsPercentageOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetMaxPowerActionsPercentageOfMachines(v int32)`

SetMaxPowerActionsPercentageOfMachines sets MaxPowerActionsPercentageOfMachines field to given value.

### HasMaxPowerActionsPercentageOfMachines

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasMaxPowerActionsPercentageOfMachines() bool`

HasMaxPowerActionsPercentageOfMachines returns a boolean if a field has been set.

### SetMaxPowerActionsPercentageOfMachinesNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetMaxPowerActionsPercentageOfMachinesNil(b bool)`

 SetMaxPowerActionsPercentageOfMachinesNil sets the value for MaxPowerActionsPercentageOfMachines to be an explicit nil

### UnsetMaxPowerActionsPercentageOfMachines
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetMaxPowerActionsPercentageOfMachines()`

UnsetMaxPowerActionsPercentageOfMachines ensures that no value is present for MaxPowerActionsPercentageOfMachines, not even an explicit nil
### GetConnectionOptions

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetConnectionOptions() string`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetConnectionOptionsOk() (*string, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetConnectionOptions(v string)`

SetConnectionOptions sets ConnectionOptions field to given value.

### HasConnectionOptions

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasConnectionOptions() bool`

HasConnectionOptions returns a boolean if a field has been set.

### SetConnectionOptionsNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetConnectionOptionsNil(b bool)`

 SetConnectionOptionsNil sets the value for ConnectionOptions to be an explicit nil

### UnsetConnectionOptions
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetConnectionOptions()`

UnsetConnectionOptions ensures that no value is present for ConnectionOptions, not even an explicit nil
### GetZone

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetMetadata

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetApiKey

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSecretKeyFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSecretKeyFormat() IdentityPasswordFormat`

GetSecretKeyFormat returns the SecretKeyFormat field if non-nil, zero value otherwise.

### GetSecretKeyFormatOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSecretKeyFormatOk() (*IdentityPasswordFormat, bool)`

GetSecretKeyFormatOk returns a tuple with the SecretKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetSecretKeyFormat(v IdentityPasswordFormat)`

SetSecretKeyFormat sets SecretKeyFormat field to given value.

### HasSecretKeyFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasSecretKeyFormat() bool`

HasSecretKeyFormat returns a boolean if a field has been set.

### GetRegion

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetAddress

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCustomProperties

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetApplicationId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetApplicationSecret

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetApplicationSecret() string`

GetApplicationSecret returns the ApplicationSecret field if non-nil, zero value otherwise.

### GetApplicationSecretOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetApplicationSecretOk() (*string, bool)`

GetApplicationSecretOk returns a tuple with the ApplicationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecret

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetApplicationSecret(v string)`

SetApplicationSecret sets ApplicationSecret field to given value.

### HasApplicationSecret

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasApplicationSecret() bool`

HasApplicationSecret returns a boolean if a field has been set.

### SetApplicationSecretNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetApplicationSecretNil(b bool)`

 SetApplicationSecretNil sets the value for ApplicationSecret to be an explicit nil

### UnsetApplicationSecret
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetApplicationSecret()`

UnsetApplicationSecret ensures that no value is present for ApplicationSecret, not even an explicit nil
### GetApplicationSecretFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetApplicationSecretFormat() IdentityPasswordFormat`

GetApplicationSecretFormat returns the ApplicationSecretFormat field if non-nil, zero value otherwise.

### GetApplicationSecretFormatOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetApplicationSecretFormatOk() (*IdentityPasswordFormat, bool)`

GetApplicationSecretFormatOk returns a tuple with the ApplicationSecretFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecretFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetApplicationSecretFormat(v IdentityPasswordFormat)`

SetApplicationSecretFormat sets ApplicationSecretFormat field to given value.

### HasApplicationSecretFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasApplicationSecretFormat() bool`

HasApplicationSecretFormat returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetActiveDirectoryId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.

### HasActiveDirectoryId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasActiveDirectoryId() bool`

HasActiveDirectoryId returns a boolean if a field has been set.

### GetEnvironment

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetEnvironment() AzureEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetEnvironmentOk() (*AzureEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetEnvironment(v AzureEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetManagementEndpoint

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetManagementEndpoint() string`

GetManagementEndpoint returns the ManagementEndpoint field if non-nil, zero value otherwise.

### GetManagementEndpointOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetManagementEndpointOk() (*string, bool)`

GetManagementEndpointOk returns a tuple with the ManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEndpoint

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetManagementEndpoint(v string)`

SetManagementEndpoint sets ManagementEndpoint field to given value.

### HasManagementEndpoint

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasManagementEndpoint() bool`

HasManagementEndpoint returns a boolean if a field has been set.

### SetManagementEndpointNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetManagementEndpointNil(b bool)`

 SetManagementEndpointNil sets the value for ManagementEndpoint to be an explicit nil

### UnsetManagementEndpoint
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetManagementEndpoint()`

UnsetManagementEndpoint ensures that no value is present for ManagementEndpoint, not even an explicit nil
### GetAuthenticationAuthority

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetAuthenticationAuthority() string`

GetAuthenticationAuthority returns the AuthenticationAuthority field if non-nil, zero value otherwise.

### GetAuthenticationAuthorityOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetAuthenticationAuthorityOk() (*string, bool)`

GetAuthenticationAuthorityOk returns a tuple with the AuthenticationAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAuthority

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetAuthenticationAuthority(v string)`

SetAuthenticationAuthority sets AuthenticationAuthority field to given value.

### HasAuthenticationAuthority

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasAuthenticationAuthority() bool`

HasAuthenticationAuthority returns a boolean if a field has been set.

### SetAuthenticationAuthorityNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetAuthenticationAuthorityNil(b bool)`

 SetAuthenticationAuthorityNil sets the value for AuthenticationAuthority to be an explicit nil

### UnsetAuthenticationAuthority
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetAuthenticationAuthority()`

UnsetAuthenticationAuthority ensures that no value is present for AuthenticationAuthority, not even an explicit nil
### GetStorageSuffix

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetStorageSuffix() string`

GetStorageSuffix returns the StorageSuffix field if non-nil, zero value otherwise.

### GetStorageSuffixOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetStorageSuffixOk() (*string, bool)`

GetStorageSuffixOk returns a tuple with the StorageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSuffix

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetStorageSuffix(v string)`

SetStorageSuffix sets StorageSuffix field to given value.

### HasStorageSuffix

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasStorageSuffix() bool`

HasStorageSuffix returns a boolean if a field has been set.

### SetStorageSuffixNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetStorageSuffixNil(b bool)`

 SetStorageSuffixNil sets the value for StorageSuffix to be an explicit nil

### UnsetStorageSuffix
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetStorageSuffix()`

UnsetStorageSuffix ensures that no value is present for StorageSuffix, not even an explicit nil
### GetServiceAccountId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetServiceAccountCredentials

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountCredentials() string`

GetServiceAccountCredentials returns the ServiceAccountCredentials field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountCredentialsOk() (*string, bool)`

GetServiceAccountCredentialsOk returns a tuple with the ServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentials

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetServiceAccountCredentials(v string)`

SetServiceAccountCredentials sets ServiceAccountCredentials field to given value.

### HasServiceAccountCredentials

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasServiceAccountCredentials() bool`

HasServiceAccountCredentials returns a boolean if a field has been set.

### GetServiceAccountCredentialsFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountCredentialsFormat() IdentityPasswordFormat`

GetServiceAccountCredentialsFormat returns the ServiceAccountCredentialsFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsFormatOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountCredentialsFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialsFormatOk returns a tuple with the ServiceAccountCredentialsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialsFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetServiceAccountCredentialsFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialsFormat sets ServiceAccountCredentialsFormat field to given value.

### HasServiceAccountCredentialsFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasServiceAccountCredentialsFormat() bool`

HasServiceAccountCredentialsFormat returns a boolean if a field has been set.

### GetUserName

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPassword

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPasswordFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.

### GetAddresses

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetPluginId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### SetPluginIdNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetPluginIdNil(b bool)`

 SetPluginIdNil sets the value for PluginId to be an explicit nil

### UnsetPluginId
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetPluginId()`

UnsetPluginId ensures that no value is present for PluginId, not even an explicit nil
### GetSslThumbprints

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSslThumbprints() []string`

GetSslThumbprints returns the SslThumbprints field if non-nil, zero value otherwise.

### GetSslThumbprintsOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSslThumbprintsOk() (*[]string, bool)`

GetSslThumbprintsOk returns a tuple with the SslThumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprints

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetSslThumbprints(v []string)`

SetSslThumbprints sets SslThumbprints field to given value.

### HasSslThumbprints

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasSslThumbprints() bool`

HasSslThumbprints returns a boolean if a field has been set.

### SetSslThumbprintsNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetSslThumbprintsNil(b bool)`

 SetSslThumbprintsNil sets the value for SslThumbprints to be an explicit nil

### UnsetSslThumbprints
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetSslThumbprints()`

UnsetSslThumbprints ensures that no value is present for SslThumbprints, not even an explicit nil
### GetSccmWakeUpProxy

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSccmWakeUpProxy() bool`

GetSccmWakeUpProxy returns the SccmWakeUpProxy field if non-nil, zero value otherwise.

### GetSccmWakeUpProxyOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetSccmWakeUpProxyOk() (*bool, bool)`

GetSccmWakeUpProxyOk returns a tuple with the SccmWakeUpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmWakeUpProxy

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetSccmWakeUpProxy(v bool)`

SetSccmWakeUpProxy sets SccmWakeUpProxy field to given value.

### HasSccmWakeUpProxy

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasSccmWakeUpProxy() bool`

HasSccmWakeUpProxy returns a boolean if a field has been set.

### SetSccmWakeUpProxyNil

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetSccmWakeUpProxyNil(b bool)`

 SetSccmWakeUpProxyNil sets the value for SccmWakeUpProxy to be an explicit nil

### UnsetSccmWakeUpProxy
`func (o *HypervisorOpenShiftConnectionDetailRequestModel) UnsetSccmWakeUpProxy()`

UnsetSccmWakeUpProxy ensures that no value is present for SccmWakeUpProxy, not even an explicit nil
### GetWakeOnLanPackets

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetWakeOnLanPackets() WakeOnLanTransmission`

GetWakeOnLanPackets returns the WakeOnLanPackets field if non-nil, zero value otherwise.

### GetWakeOnLanPacketsOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetWakeOnLanPacketsOk() (*WakeOnLanTransmission, bool)`

GetWakeOnLanPacketsOk returns a tuple with the WakeOnLanPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanPackets

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetWakeOnLanPackets(v WakeOnLanTransmission)`

SetWakeOnLanPackets sets WakeOnLanPackets field to given value.

### HasWakeOnLanPackets

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasWakeOnLanPackets() bool`

HasWakeOnLanPackets returns a boolean if a field has been set.

### GetTenancyOcid

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetTenancyOcid() string`

GetTenancyOcid returns the TenancyOcid field if non-nil, zero value otherwise.

### GetTenancyOcidOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetTenancyOcidOk() (*string, bool)`

GetTenancyOcidOk returns a tuple with the TenancyOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyOcid

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetTenancyOcid(v string)`

SetTenancyOcid sets TenancyOcid field to given value.

### HasTenancyOcid

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasTenancyOcid() bool`

HasTenancyOcid returns a boolean if a field has been set.

### GetServiceAccountCredential

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountCredential() string`

GetServiceAccountCredential returns the ServiceAccountCredential field if non-nil, zero value otherwise.

### GetServiceAccountCredentialOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountCredentialOk() (*string, bool)`

GetServiceAccountCredentialOk returns a tuple with the ServiceAccountCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredential

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetServiceAccountCredential(v string)`

SetServiceAccountCredential sets ServiceAccountCredential field to given value.

### HasServiceAccountCredential

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasServiceAccountCredential() bool`

HasServiceAccountCredential returns a boolean if a field has been set.

### GetServiceAccountCredentialFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountCredentialFormat() IdentityPasswordFormat`

GetServiceAccountCredentialFormat returns the ServiceAccountCredentialFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialFormatOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountCredentialFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialFormatOk returns a tuple with the ServiceAccountCredentialFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetServiceAccountCredentialFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialFormat sets ServiceAccountCredentialFormat field to given value.

### HasServiceAccountCredentialFormat

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasServiceAccountCredentialFormat() bool`

HasServiceAccountCredentialFormat returns a boolean if a field has been set.

### GetServiceAccountFingerprint

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountFingerprint() string`

GetServiceAccountFingerprint returns the ServiceAccountFingerprint field if non-nil, zero value otherwise.

### GetServiceAccountFingerprintOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetServiceAccountFingerprintOk() (*string, bool)`

GetServiceAccountFingerprintOk returns a tuple with the ServiceAccountFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountFingerprint

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetServiceAccountFingerprint(v string)`

SetServiceAccountFingerprint sets ServiceAccountFingerprint field to given value.

### HasServiceAccountFingerprint

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasServiceAccountFingerprint() bool`

HasServiceAccountFingerprint returns a boolean if a field has been set.

### GetOciEnvironment

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetOciEnvironment() OciEnvironment`

GetOciEnvironment returns the OciEnvironment field if non-nil, zero value otherwise.

### GetOciEnvironmentOk

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) GetOciEnvironmentOk() (*OciEnvironment, bool)`

GetOciEnvironmentOk returns a tuple with the OciEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciEnvironment

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) SetOciEnvironment(v OciEnvironment)`

SetOciEnvironment sets OciEnvironment field to given value.

### HasOciEnvironment

`func (o *HypervisorOpenShiftConnectionDetailRequestModel) HasOciEnvironment() bool`

HasOciEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


