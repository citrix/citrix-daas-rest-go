# HypervisorDetailOciResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | The API key used to authenticate with the AWS APIs. | 
**Region** | **string** | The AWS region which the hypervisor represents. | 
**MaximumConcurrentProvisioningOperations** | Pointer to **NullableInt32** | AWS maximum concurrent provisioning operations. | [optional] 
**ApplicationId** | **string** | Application ID of the service principal used to access the Azure APIs. | 
**SubscriptionId** | **string** | Azure subscription ID. | 
**ActiveDirectoryId** | **string** | Azure active directory ID. | 
**Environment** | [**AzureEnvironment**](AzureEnvironment.md) |  | 
**ManagementEndpoint** | **string** | Azure management endpoint. | 
**AuthenticationAuthority** | **string** | Azure authentication authority. | 
**StorageSuffix** | **string** | Azure storage suffix. | 
**ServiceAccountId** | **string** | The service account ID used to access the Oracle Cloud Infrastructure APIs. | 
**SccmWakeUpProxy** | **bool** | Indicates whether Microsoft System Center Configuration Manager 2012 SP1 Wake-up Proxy is used for power management. | 
**WakeOnLanPackets** | [**WakeOnLanTransmission**](WakeOnLanTransmission.md) |  | 
**SslThumbprints** | Pointer to **[]string** | SSL thumbprints considered acceptable for the SSL certificate presented by the hypervisor. | [optional] 
**UserName** | **string** | The user name for the credentials used to communicate with the hypervisor. | 
**TenancyOcid** | **string** | Oracle Cloud Infrastructure tenancy to connect to. Required. | 
**OciRegion** | **string** | The region of Oracle Cloud Infrastructure. | 
**OciEnvironment** | **string** | The environment of Oracle Cloud Infrastructure, Commercial or Government. | 
**Fingerprint** | **string** | The fingerprint of Oracle Cloud Infrastructure API key. | 
**Id** | Pointer to **NullableString** | Id of the resource. | [optional] 
**Name** | Pointer to **NullableString** | Name of the resource. | [optional] 
**XDPath** | Pointer to **NullableString** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**Addresses** | **[]string** | Addresses that can be used to contact the required hypervisor. All the addresses are considered equivalent, that is, all of the addresses provide access to the same virtual machines, snapshots, network, and storage. | 
**InMaintenanceMode** | **bool** | Indicates whether the hypervisor is in maintenance mode, which disables all communication between XenApp &amp; XenDesktop and the Hypervisor. | 
**Unentitled** | Pointer to **bool** | Indicates whether is UnEntitled. | [optional] 
**PluginId** | **string** | The class name for the Citrix Managed Machine library that is used to access the hypervisor. | 
**Scopes** | [**[]ScopeResponseModel**](ScopeResponseModel.md) | The list of administrative scopes that the connection is a part of. The scopes control which administrators are able to work with the connection. | 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the hypervisor is assigned to.  If &#x60;null&#x60;, the hypervisor is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**UsesCloudInfrastructure** | **bool** | Indicates whether the hypervisor uses cloud infrastructure. | 
**Zone** | [**RefResponseModel**](RefResponseModel.md) |  | 
**Fault** | Pointer to [**HypervisorFaultResponseModel**](HypervisorFaultResponseModel.md) |  | [optional] 
**CustomProperties** | Pointer to **NullableString** | CustomProperties of hypervisor connection | [optional] 
**Uid** | Pointer to **NullableInt32** | The broker id. | [optional] 
**IsVirtual** | Pointer to **bool** | If this connection is virtual. | [optional] 
**Capabilities** | **[]string** | List of capabilities supported by the hypervisor. | 
**ConfigurationObjectCapabilities** | [**[]HypervisorConfigurationObjectCapability**](HypervisorConfigurationObjectCapability.md) | List of configuration capabilities supported by each type of object on the hypervisor. | 
**PluginRevision** | **string** | Revision of the plugin in use. | 
**MaxAbsoluteActiveActions** | **int32** | Maximum number of actions that can execute in parallel on the hypervisor. | 
**MaxAbsoluteNewActionsPerMinute** | **int32** | Maximum number of actions that can be started on the hypervisor per-minute. | 
**MaxPowerActionsPercentageOfMachines** | **int32** | Maximum percentage of machines on the hypervisor which can have their power state changed simultaneously. | 
**ConnectionOptions** | **string** | Connection options to use for the hypervisor. | 
**SupportsLocalStorageCaching** | **bool** | Indicates whether the hypervisor supports local storage caching. | 
**SupportsPvsVms** | **bool** | Indicates whether the hypervisor supports PVS VMs. | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata for hypervisor connections. | [optional] 

## Methods

### NewHypervisorDetailOciResponseModel

`func NewHypervisorDetailOciResponseModel(apiKey string, region string, applicationId string, subscriptionId string, activeDirectoryId string, environment AzureEnvironment, managementEndpoint string, authenticationAuthority string, storageSuffix string, serviceAccountId string, sccmWakeUpProxy bool, wakeOnLanPackets WakeOnLanTransmission, userName string, tenancyOcid string, ociRegion string, ociEnvironment string, fingerprint string, connectionType HypervisorConnectionType, addresses []string, inMaintenanceMode bool, pluginId string, scopes []ScopeResponseModel, usesCloudInfrastructure bool, zone RefResponseModel, capabilities []string, configurationObjectCapabilities []HypervisorConfigurationObjectCapability, pluginRevision string, maxAbsoluteActiveActions int32, maxAbsoluteNewActionsPerMinute int32, maxPowerActionsPercentageOfMachines int32, connectionOptions string, supportsLocalStorageCaching bool, supportsPvsVms bool, ) *HypervisorDetailOciResponseModel`

NewHypervisorDetailOciResponseModel instantiates a new HypervisorDetailOciResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorDetailOciResponseModelWithDefaults

`func NewHypervisorDetailOciResponseModelWithDefaults() *HypervisorDetailOciResponseModel`

NewHypervisorDetailOciResponseModelWithDefaults instantiates a new HypervisorDetailOciResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *HypervisorDetailOciResponseModel) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HypervisorDetailOciResponseModel) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HypervisorDetailOciResponseModel) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetRegion

`func (o *HypervisorDetailOciResponseModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorDetailOciResponseModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorDetailOciResponseModel) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetMaximumConcurrentProvisioningOperations

`func (o *HypervisorDetailOciResponseModel) GetMaximumConcurrentProvisioningOperations() int32`

GetMaximumConcurrentProvisioningOperations returns the MaximumConcurrentProvisioningOperations field if non-nil, zero value otherwise.

### GetMaximumConcurrentProvisioningOperationsOk

`func (o *HypervisorDetailOciResponseModel) GetMaximumConcurrentProvisioningOperationsOk() (*int32, bool)`

GetMaximumConcurrentProvisioningOperationsOk returns a tuple with the MaximumConcurrentProvisioningOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentProvisioningOperations

`func (o *HypervisorDetailOciResponseModel) SetMaximumConcurrentProvisioningOperations(v int32)`

SetMaximumConcurrentProvisioningOperations sets MaximumConcurrentProvisioningOperations field to given value.

### HasMaximumConcurrentProvisioningOperations

`func (o *HypervisorDetailOciResponseModel) HasMaximumConcurrentProvisioningOperations() bool`

HasMaximumConcurrentProvisioningOperations returns a boolean if a field has been set.

### SetMaximumConcurrentProvisioningOperationsNil

`func (o *HypervisorDetailOciResponseModel) SetMaximumConcurrentProvisioningOperationsNil(b bool)`

 SetMaximumConcurrentProvisioningOperationsNil sets the value for MaximumConcurrentProvisioningOperations to be an explicit nil

### UnsetMaximumConcurrentProvisioningOperations
`func (o *HypervisorDetailOciResponseModel) UnsetMaximumConcurrentProvisioningOperations()`

UnsetMaximumConcurrentProvisioningOperations ensures that no value is present for MaximumConcurrentProvisioningOperations, not even an explicit nil
### GetApplicationId

`func (o *HypervisorDetailOciResponseModel) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *HypervisorDetailOciResponseModel) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *HypervisorDetailOciResponseModel) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetSubscriptionId

`func (o *HypervisorDetailOciResponseModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *HypervisorDetailOciResponseModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *HypervisorDetailOciResponseModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetActiveDirectoryId

`func (o *HypervisorDetailOciResponseModel) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *HypervisorDetailOciResponseModel) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *HypervisorDetailOciResponseModel) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.


### GetEnvironment

`func (o *HypervisorDetailOciResponseModel) GetEnvironment() AzureEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *HypervisorDetailOciResponseModel) GetEnvironmentOk() (*AzureEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *HypervisorDetailOciResponseModel) SetEnvironment(v AzureEnvironment)`

SetEnvironment sets Environment field to given value.


### GetManagementEndpoint

`func (o *HypervisorDetailOciResponseModel) GetManagementEndpoint() string`

GetManagementEndpoint returns the ManagementEndpoint field if non-nil, zero value otherwise.

### GetManagementEndpointOk

`func (o *HypervisorDetailOciResponseModel) GetManagementEndpointOk() (*string, bool)`

GetManagementEndpointOk returns a tuple with the ManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEndpoint

`func (o *HypervisorDetailOciResponseModel) SetManagementEndpoint(v string)`

SetManagementEndpoint sets ManagementEndpoint field to given value.


### GetAuthenticationAuthority

`func (o *HypervisorDetailOciResponseModel) GetAuthenticationAuthority() string`

GetAuthenticationAuthority returns the AuthenticationAuthority field if non-nil, zero value otherwise.

### GetAuthenticationAuthorityOk

`func (o *HypervisorDetailOciResponseModel) GetAuthenticationAuthorityOk() (*string, bool)`

GetAuthenticationAuthorityOk returns a tuple with the AuthenticationAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAuthority

`func (o *HypervisorDetailOciResponseModel) SetAuthenticationAuthority(v string)`

SetAuthenticationAuthority sets AuthenticationAuthority field to given value.


### GetStorageSuffix

`func (o *HypervisorDetailOciResponseModel) GetStorageSuffix() string`

GetStorageSuffix returns the StorageSuffix field if non-nil, zero value otherwise.

### GetStorageSuffixOk

`func (o *HypervisorDetailOciResponseModel) GetStorageSuffixOk() (*string, bool)`

GetStorageSuffixOk returns a tuple with the StorageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSuffix

`func (o *HypervisorDetailOciResponseModel) SetStorageSuffix(v string)`

SetStorageSuffix sets StorageSuffix field to given value.


### GetServiceAccountId

`func (o *HypervisorDetailOciResponseModel) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *HypervisorDetailOciResponseModel) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *HypervisorDetailOciResponseModel) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetSccmWakeUpProxy

`func (o *HypervisorDetailOciResponseModel) GetSccmWakeUpProxy() bool`

GetSccmWakeUpProxy returns the SccmWakeUpProxy field if non-nil, zero value otherwise.

### GetSccmWakeUpProxyOk

`func (o *HypervisorDetailOciResponseModel) GetSccmWakeUpProxyOk() (*bool, bool)`

GetSccmWakeUpProxyOk returns a tuple with the SccmWakeUpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmWakeUpProxy

`func (o *HypervisorDetailOciResponseModel) SetSccmWakeUpProxy(v bool)`

SetSccmWakeUpProxy sets SccmWakeUpProxy field to given value.


### GetWakeOnLanPackets

`func (o *HypervisorDetailOciResponseModel) GetWakeOnLanPackets() WakeOnLanTransmission`

GetWakeOnLanPackets returns the WakeOnLanPackets field if non-nil, zero value otherwise.

### GetWakeOnLanPacketsOk

`func (o *HypervisorDetailOciResponseModel) GetWakeOnLanPacketsOk() (*WakeOnLanTransmission, bool)`

GetWakeOnLanPacketsOk returns a tuple with the WakeOnLanPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanPackets

`func (o *HypervisorDetailOciResponseModel) SetWakeOnLanPackets(v WakeOnLanTransmission)`

SetWakeOnLanPackets sets WakeOnLanPackets field to given value.


### GetSslThumbprints

`func (o *HypervisorDetailOciResponseModel) GetSslThumbprints() []string`

GetSslThumbprints returns the SslThumbprints field if non-nil, zero value otherwise.

### GetSslThumbprintsOk

`func (o *HypervisorDetailOciResponseModel) GetSslThumbprintsOk() (*[]string, bool)`

GetSslThumbprintsOk returns a tuple with the SslThumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprints

`func (o *HypervisorDetailOciResponseModel) SetSslThumbprints(v []string)`

SetSslThumbprints sets SslThumbprints field to given value.

### HasSslThumbprints

`func (o *HypervisorDetailOciResponseModel) HasSslThumbprints() bool`

HasSslThumbprints returns a boolean if a field has been set.

### SetSslThumbprintsNil

`func (o *HypervisorDetailOciResponseModel) SetSslThumbprintsNil(b bool)`

 SetSslThumbprintsNil sets the value for SslThumbprints to be an explicit nil

### UnsetSslThumbprints
`func (o *HypervisorDetailOciResponseModel) UnsetSslThumbprints()`

UnsetSslThumbprints ensures that no value is present for SslThumbprints, not even an explicit nil
### GetUserName

`func (o *HypervisorDetailOciResponseModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HypervisorDetailOciResponseModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HypervisorDetailOciResponseModel) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetTenancyOcid

`func (o *HypervisorDetailOciResponseModel) GetTenancyOcid() string`

GetTenancyOcid returns the TenancyOcid field if non-nil, zero value otherwise.

### GetTenancyOcidOk

`func (o *HypervisorDetailOciResponseModel) GetTenancyOcidOk() (*string, bool)`

GetTenancyOcidOk returns a tuple with the TenancyOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyOcid

`func (o *HypervisorDetailOciResponseModel) SetTenancyOcid(v string)`

SetTenancyOcid sets TenancyOcid field to given value.


### GetOciRegion

`func (o *HypervisorDetailOciResponseModel) GetOciRegion() string`

GetOciRegion returns the OciRegion field if non-nil, zero value otherwise.

### GetOciRegionOk

`func (o *HypervisorDetailOciResponseModel) GetOciRegionOk() (*string, bool)`

GetOciRegionOk returns a tuple with the OciRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciRegion

`func (o *HypervisorDetailOciResponseModel) SetOciRegion(v string)`

SetOciRegion sets OciRegion field to given value.


### GetOciEnvironment

`func (o *HypervisorDetailOciResponseModel) GetOciEnvironment() string`

GetOciEnvironment returns the OciEnvironment field if non-nil, zero value otherwise.

### GetOciEnvironmentOk

`func (o *HypervisorDetailOciResponseModel) GetOciEnvironmentOk() (*string, bool)`

GetOciEnvironmentOk returns a tuple with the OciEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciEnvironment

`func (o *HypervisorDetailOciResponseModel) SetOciEnvironment(v string)`

SetOciEnvironment sets OciEnvironment field to given value.


### GetFingerprint

`func (o *HypervisorDetailOciResponseModel) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *HypervisorDetailOciResponseModel) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *HypervisorDetailOciResponseModel) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetId

`func (o *HypervisorDetailOciResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorDetailOciResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorDetailOciResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorDetailOciResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorDetailOciResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorDetailOciResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorDetailOciResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorDetailOciResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorDetailOciResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorDetailOciResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorDetailOciResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorDetailOciResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorDetailOciResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorDetailOciResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorDetailOciResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorDetailOciResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorDetailOciResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorDetailOciResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetConnectionType

`func (o *HypervisorDetailOciResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorDetailOciResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorDetailOciResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetAddresses

`func (o *HypervisorDetailOciResponseModel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorDetailOciResponseModel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorDetailOciResponseModel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetInMaintenanceMode

`func (o *HypervisorDetailOciResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *HypervisorDetailOciResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *HypervisorDetailOciResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetUnentitled

`func (o *HypervisorDetailOciResponseModel) GetUnentitled() bool`

GetUnentitled returns the Unentitled field if non-nil, zero value otherwise.

### GetUnentitledOk

`func (o *HypervisorDetailOciResponseModel) GetUnentitledOk() (*bool, bool)`

GetUnentitledOk returns a tuple with the Unentitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnentitled

`func (o *HypervisorDetailOciResponseModel) SetUnentitled(v bool)`

SetUnentitled sets Unentitled field to given value.

### HasUnentitled

`func (o *HypervisorDetailOciResponseModel) HasUnentitled() bool`

HasUnentitled returns a boolean if a field has been set.

### GetPluginId

`func (o *HypervisorDetailOciResponseModel) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorDetailOciResponseModel) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorDetailOciResponseModel) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.


### GetScopes

`func (o *HypervisorDetailOciResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorDetailOciResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorDetailOciResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *HypervisorDetailOciResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorDetailOciResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorDetailOciResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorDetailOciResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *HypervisorDetailOciResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *HypervisorDetailOciResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetUsesCloudInfrastructure

`func (o *HypervisorDetailOciResponseModel) GetUsesCloudInfrastructure() bool`

GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field if non-nil, zero value otherwise.

### GetUsesCloudInfrastructureOk

`func (o *HypervisorDetailOciResponseModel) GetUsesCloudInfrastructureOk() (*bool, bool)`

GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCloudInfrastructure

`func (o *HypervisorDetailOciResponseModel) SetUsesCloudInfrastructure(v bool)`

SetUsesCloudInfrastructure sets UsesCloudInfrastructure field to given value.


### GetZone

`func (o *HypervisorDetailOciResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorDetailOciResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorDetailOciResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.


### GetFault

`func (o *HypervisorDetailOciResponseModel) GetFault() HypervisorFaultResponseModel`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *HypervisorDetailOciResponseModel) GetFaultOk() (*HypervisorFaultResponseModel, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *HypervisorDetailOciResponseModel) SetFault(v HypervisorFaultResponseModel)`

SetFault sets Fault field to given value.

### HasFault

`func (o *HypervisorDetailOciResponseModel) HasFault() bool`

HasFault returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorDetailOciResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorDetailOciResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorDetailOciResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorDetailOciResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *HypervisorDetailOciResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *HypervisorDetailOciResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetUid

`func (o *HypervisorDetailOciResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *HypervisorDetailOciResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *HypervisorDetailOciResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *HypervisorDetailOciResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *HypervisorDetailOciResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *HypervisorDetailOciResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetIsVirtual

`func (o *HypervisorDetailOciResponseModel) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *HypervisorDetailOciResponseModel) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *HypervisorDetailOciResponseModel) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *HypervisorDetailOciResponseModel) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetCapabilities

`func (o *HypervisorDetailOciResponseModel) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *HypervisorDetailOciResponseModel) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *HypervisorDetailOciResponseModel) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetConfigurationObjectCapabilities

`func (o *HypervisorDetailOciResponseModel) GetConfigurationObjectCapabilities() []HypervisorConfigurationObjectCapability`

GetConfigurationObjectCapabilities returns the ConfigurationObjectCapabilities field if non-nil, zero value otherwise.

### GetConfigurationObjectCapabilitiesOk

`func (o *HypervisorDetailOciResponseModel) GetConfigurationObjectCapabilitiesOk() (*[]HypervisorConfigurationObjectCapability, bool)`

GetConfigurationObjectCapabilitiesOk returns a tuple with the ConfigurationObjectCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationObjectCapabilities

`func (o *HypervisorDetailOciResponseModel) SetConfigurationObjectCapabilities(v []HypervisorConfigurationObjectCapability)`

SetConfigurationObjectCapabilities sets ConfigurationObjectCapabilities field to given value.


### GetPluginRevision

`func (o *HypervisorDetailOciResponseModel) GetPluginRevision() string`

GetPluginRevision returns the PluginRevision field if non-nil, zero value otherwise.

### GetPluginRevisionOk

`func (o *HypervisorDetailOciResponseModel) GetPluginRevisionOk() (*string, bool)`

GetPluginRevisionOk returns a tuple with the PluginRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginRevision

`func (o *HypervisorDetailOciResponseModel) SetPluginRevision(v string)`

SetPluginRevision sets PluginRevision field to given value.


### GetMaxAbsoluteActiveActions

`func (o *HypervisorDetailOciResponseModel) GetMaxAbsoluteActiveActions() int32`

GetMaxAbsoluteActiveActions returns the MaxAbsoluteActiveActions field if non-nil, zero value otherwise.

### GetMaxAbsoluteActiveActionsOk

`func (o *HypervisorDetailOciResponseModel) GetMaxAbsoluteActiveActionsOk() (*int32, bool)`

GetMaxAbsoluteActiveActionsOk returns a tuple with the MaxAbsoluteActiveActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteActiveActions

`func (o *HypervisorDetailOciResponseModel) SetMaxAbsoluteActiveActions(v int32)`

SetMaxAbsoluteActiveActions sets MaxAbsoluteActiveActions field to given value.


### GetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorDetailOciResponseModel) GetMaxAbsoluteNewActionsPerMinute() int32`

GetMaxAbsoluteNewActionsPerMinute returns the MaxAbsoluteNewActionsPerMinute field if non-nil, zero value otherwise.

### GetMaxAbsoluteNewActionsPerMinuteOk

`func (o *HypervisorDetailOciResponseModel) GetMaxAbsoluteNewActionsPerMinuteOk() (*int32, bool)`

GetMaxAbsoluteNewActionsPerMinuteOk returns a tuple with the MaxAbsoluteNewActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorDetailOciResponseModel) SetMaxAbsoluteNewActionsPerMinute(v int32)`

SetMaxAbsoluteNewActionsPerMinute sets MaxAbsoluteNewActionsPerMinute field to given value.


### GetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorDetailOciResponseModel) GetMaxPowerActionsPercentageOfMachines() int32`

GetMaxPowerActionsPercentageOfMachines returns the MaxPowerActionsPercentageOfMachines field if non-nil, zero value otherwise.

### GetMaxPowerActionsPercentageOfMachinesOk

`func (o *HypervisorDetailOciResponseModel) GetMaxPowerActionsPercentageOfMachinesOk() (*int32, bool)`

GetMaxPowerActionsPercentageOfMachinesOk returns a tuple with the MaxPowerActionsPercentageOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorDetailOciResponseModel) SetMaxPowerActionsPercentageOfMachines(v int32)`

SetMaxPowerActionsPercentageOfMachines sets MaxPowerActionsPercentageOfMachines field to given value.


### GetConnectionOptions

`func (o *HypervisorDetailOciResponseModel) GetConnectionOptions() string`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *HypervisorDetailOciResponseModel) GetConnectionOptionsOk() (*string, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *HypervisorDetailOciResponseModel) SetConnectionOptions(v string)`

SetConnectionOptions sets ConnectionOptions field to given value.


### GetSupportsLocalStorageCaching

`func (o *HypervisorDetailOciResponseModel) GetSupportsLocalStorageCaching() bool`

GetSupportsLocalStorageCaching returns the SupportsLocalStorageCaching field if non-nil, zero value otherwise.

### GetSupportsLocalStorageCachingOk

`func (o *HypervisorDetailOciResponseModel) GetSupportsLocalStorageCachingOk() (*bool, bool)`

GetSupportsLocalStorageCachingOk returns a tuple with the SupportsLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLocalStorageCaching

`func (o *HypervisorDetailOciResponseModel) SetSupportsLocalStorageCaching(v bool)`

SetSupportsLocalStorageCaching sets SupportsLocalStorageCaching field to given value.


### GetSupportsPvsVms

`func (o *HypervisorDetailOciResponseModel) GetSupportsPvsVms() bool`

GetSupportsPvsVms returns the SupportsPvsVms field if non-nil, zero value otherwise.

### GetSupportsPvsVmsOk

`func (o *HypervisorDetailOciResponseModel) GetSupportsPvsVmsOk() (*bool, bool)`

GetSupportsPvsVmsOk returns a tuple with the SupportsPvsVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPvsVms

`func (o *HypervisorDetailOciResponseModel) SetSupportsPvsVms(v bool)`

SetSupportsPvsVms sets SupportsPvsVms field to given value.


### GetMetadata

`func (o *HypervisorDetailOciResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorDetailOciResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorDetailOciResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorDetailOciResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HypervisorDetailOciResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HypervisorDetailOciResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


