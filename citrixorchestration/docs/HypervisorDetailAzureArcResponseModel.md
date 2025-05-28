# HypervisorDetailAzureArcResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | The API key used to authenticate with the AWS APIs. | 
**Region** | **string** | The AWS region which the hypervisor represents. | 
**MaximumConcurrentProvisioningOperations** | Pointer to **NullableInt32** | AWS maximum concurrent provisioning operations. | [optional] 
**ApplicationId** | **string** | Application ID of the service principal used to access the Azure Arc APIs. | 
**SubscriptionId** | **string** | Azure subscription ID. | 
**ActiveDirectoryId** | **string** | Azure active directory ID. | 
**Environment** | [**AzureEnvironment**](AzureEnvironment.md) |  | 
**ManagementEndpoint** | **string** | Azure management endpoint. | 
**AuthenticationAuthority** | **string** | Azure authentication authority. | 
**StorageSuffix** | **string** | Azure storage suffix. | 
**ServiceAccountId** | **string** | The service account ID used to access the Google Cloud APIs. | 
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

### NewHypervisorDetailAzureArcResponseModel

`func NewHypervisorDetailAzureArcResponseModel(apiKey string, region string, applicationId string, subscriptionId string, activeDirectoryId string, environment AzureEnvironment, managementEndpoint string, authenticationAuthority string, storageSuffix string, serviceAccountId string, sccmWakeUpProxy bool, wakeOnLanPackets WakeOnLanTransmission, userName string, tenancyOcid string, ociRegion string, ociEnvironment string, fingerprint string, connectionType HypervisorConnectionType, addresses []string, inMaintenanceMode bool, pluginId string, scopes []ScopeResponseModel, usesCloudInfrastructure bool, zone RefResponseModel, capabilities []string, configurationObjectCapabilities []HypervisorConfigurationObjectCapability, pluginRevision string, maxAbsoluteActiveActions int32, maxAbsoluteNewActionsPerMinute int32, maxPowerActionsPercentageOfMachines int32, connectionOptions string, supportsLocalStorageCaching bool, supportsPvsVms bool, ) *HypervisorDetailAzureArcResponseModel`

NewHypervisorDetailAzureArcResponseModel instantiates a new HypervisorDetailAzureArcResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorDetailAzureArcResponseModelWithDefaults

`func NewHypervisorDetailAzureArcResponseModelWithDefaults() *HypervisorDetailAzureArcResponseModel`

NewHypervisorDetailAzureArcResponseModelWithDefaults instantiates a new HypervisorDetailAzureArcResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *HypervisorDetailAzureArcResponseModel) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HypervisorDetailAzureArcResponseModel) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HypervisorDetailAzureArcResponseModel) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetRegion

`func (o *HypervisorDetailAzureArcResponseModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorDetailAzureArcResponseModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorDetailAzureArcResponseModel) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetMaximumConcurrentProvisioningOperations

`func (o *HypervisorDetailAzureArcResponseModel) GetMaximumConcurrentProvisioningOperations() int32`

GetMaximumConcurrentProvisioningOperations returns the MaximumConcurrentProvisioningOperations field if non-nil, zero value otherwise.

### GetMaximumConcurrentProvisioningOperationsOk

`func (o *HypervisorDetailAzureArcResponseModel) GetMaximumConcurrentProvisioningOperationsOk() (*int32, bool)`

GetMaximumConcurrentProvisioningOperationsOk returns a tuple with the MaximumConcurrentProvisioningOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentProvisioningOperations

`func (o *HypervisorDetailAzureArcResponseModel) SetMaximumConcurrentProvisioningOperations(v int32)`

SetMaximumConcurrentProvisioningOperations sets MaximumConcurrentProvisioningOperations field to given value.

### HasMaximumConcurrentProvisioningOperations

`func (o *HypervisorDetailAzureArcResponseModel) HasMaximumConcurrentProvisioningOperations() bool`

HasMaximumConcurrentProvisioningOperations returns a boolean if a field has been set.

### SetMaximumConcurrentProvisioningOperationsNil

`func (o *HypervisorDetailAzureArcResponseModel) SetMaximumConcurrentProvisioningOperationsNil(b bool)`

 SetMaximumConcurrentProvisioningOperationsNil sets the value for MaximumConcurrentProvisioningOperations to be an explicit nil

### UnsetMaximumConcurrentProvisioningOperations
`func (o *HypervisorDetailAzureArcResponseModel) UnsetMaximumConcurrentProvisioningOperations()`

UnsetMaximumConcurrentProvisioningOperations ensures that no value is present for MaximumConcurrentProvisioningOperations, not even an explicit nil
### GetApplicationId

`func (o *HypervisorDetailAzureArcResponseModel) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *HypervisorDetailAzureArcResponseModel) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *HypervisorDetailAzureArcResponseModel) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetSubscriptionId

`func (o *HypervisorDetailAzureArcResponseModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *HypervisorDetailAzureArcResponseModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *HypervisorDetailAzureArcResponseModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetActiveDirectoryId

`func (o *HypervisorDetailAzureArcResponseModel) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *HypervisorDetailAzureArcResponseModel) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *HypervisorDetailAzureArcResponseModel) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.


### GetEnvironment

`func (o *HypervisorDetailAzureArcResponseModel) GetEnvironment() AzureEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *HypervisorDetailAzureArcResponseModel) GetEnvironmentOk() (*AzureEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *HypervisorDetailAzureArcResponseModel) SetEnvironment(v AzureEnvironment)`

SetEnvironment sets Environment field to given value.


### GetManagementEndpoint

`func (o *HypervisorDetailAzureArcResponseModel) GetManagementEndpoint() string`

GetManagementEndpoint returns the ManagementEndpoint field if non-nil, zero value otherwise.

### GetManagementEndpointOk

`func (o *HypervisorDetailAzureArcResponseModel) GetManagementEndpointOk() (*string, bool)`

GetManagementEndpointOk returns a tuple with the ManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEndpoint

`func (o *HypervisorDetailAzureArcResponseModel) SetManagementEndpoint(v string)`

SetManagementEndpoint sets ManagementEndpoint field to given value.


### GetAuthenticationAuthority

`func (o *HypervisorDetailAzureArcResponseModel) GetAuthenticationAuthority() string`

GetAuthenticationAuthority returns the AuthenticationAuthority field if non-nil, zero value otherwise.

### GetAuthenticationAuthorityOk

`func (o *HypervisorDetailAzureArcResponseModel) GetAuthenticationAuthorityOk() (*string, bool)`

GetAuthenticationAuthorityOk returns a tuple with the AuthenticationAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAuthority

`func (o *HypervisorDetailAzureArcResponseModel) SetAuthenticationAuthority(v string)`

SetAuthenticationAuthority sets AuthenticationAuthority field to given value.


### GetStorageSuffix

`func (o *HypervisorDetailAzureArcResponseModel) GetStorageSuffix() string`

GetStorageSuffix returns the StorageSuffix field if non-nil, zero value otherwise.

### GetStorageSuffixOk

`func (o *HypervisorDetailAzureArcResponseModel) GetStorageSuffixOk() (*string, bool)`

GetStorageSuffixOk returns a tuple with the StorageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSuffix

`func (o *HypervisorDetailAzureArcResponseModel) SetStorageSuffix(v string)`

SetStorageSuffix sets StorageSuffix field to given value.


### GetServiceAccountId

`func (o *HypervisorDetailAzureArcResponseModel) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *HypervisorDetailAzureArcResponseModel) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *HypervisorDetailAzureArcResponseModel) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetSccmWakeUpProxy

`func (o *HypervisorDetailAzureArcResponseModel) GetSccmWakeUpProxy() bool`

GetSccmWakeUpProxy returns the SccmWakeUpProxy field if non-nil, zero value otherwise.

### GetSccmWakeUpProxyOk

`func (o *HypervisorDetailAzureArcResponseModel) GetSccmWakeUpProxyOk() (*bool, bool)`

GetSccmWakeUpProxyOk returns a tuple with the SccmWakeUpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmWakeUpProxy

`func (o *HypervisorDetailAzureArcResponseModel) SetSccmWakeUpProxy(v bool)`

SetSccmWakeUpProxy sets SccmWakeUpProxy field to given value.


### GetWakeOnLanPackets

`func (o *HypervisorDetailAzureArcResponseModel) GetWakeOnLanPackets() WakeOnLanTransmission`

GetWakeOnLanPackets returns the WakeOnLanPackets field if non-nil, zero value otherwise.

### GetWakeOnLanPacketsOk

`func (o *HypervisorDetailAzureArcResponseModel) GetWakeOnLanPacketsOk() (*WakeOnLanTransmission, bool)`

GetWakeOnLanPacketsOk returns a tuple with the WakeOnLanPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanPackets

`func (o *HypervisorDetailAzureArcResponseModel) SetWakeOnLanPackets(v WakeOnLanTransmission)`

SetWakeOnLanPackets sets WakeOnLanPackets field to given value.


### GetSslThumbprints

`func (o *HypervisorDetailAzureArcResponseModel) GetSslThumbprints() []string`

GetSslThumbprints returns the SslThumbprints field if non-nil, zero value otherwise.

### GetSslThumbprintsOk

`func (o *HypervisorDetailAzureArcResponseModel) GetSslThumbprintsOk() (*[]string, bool)`

GetSslThumbprintsOk returns a tuple with the SslThumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprints

`func (o *HypervisorDetailAzureArcResponseModel) SetSslThumbprints(v []string)`

SetSslThumbprints sets SslThumbprints field to given value.

### HasSslThumbprints

`func (o *HypervisorDetailAzureArcResponseModel) HasSslThumbprints() bool`

HasSslThumbprints returns a boolean if a field has been set.

### SetSslThumbprintsNil

`func (o *HypervisorDetailAzureArcResponseModel) SetSslThumbprintsNil(b bool)`

 SetSslThumbprintsNil sets the value for SslThumbprints to be an explicit nil

### UnsetSslThumbprints
`func (o *HypervisorDetailAzureArcResponseModel) UnsetSslThumbprints()`

UnsetSslThumbprints ensures that no value is present for SslThumbprints, not even an explicit nil
### GetUserName

`func (o *HypervisorDetailAzureArcResponseModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HypervisorDetailAzureArcResponseModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HypervisorDetailAzureArcResponseModel) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetTenancyOcid

`func (o *HypervisorDetailAzureArcResponseModel) GetTenancyOcid() string`

GetTenancyOcid returns the TenancyOcid field if non-nil, zero value otherwise.

### GetTenancyOcidOk

`func (o *HypervisorDetailAzureArcResponseModel) GetTenancyOcidOk() (*string, bool)`

GetTenancyOcidOk returns a tuple with the TenancyOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyOcid

`func (o *HypervisorDetailAzureArcResponseModel) SetTenancyOcid(v string)`

SetTenancyOcid sets TenancyOcid field to given value.


### GetOciRegion

`func (o *HypervisorDetailAzureArcResponseModel) GetOciRegion() string`

GetOciRegion returns the OciRegion field if non-nil, zero value otherwise.

### GetOciRegionOk

`func (o *HypervisorDetailAzureArcResponseModel) GetOciRegionOk() (*string, bool)`

GetOciRegionOk returns a tuple with the OciRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciRegion

`func (o *HypervisorDetailAzureArcResponseModel) SetOciRegion(v string)`

SetOciRegion sets OciRegion field to given value.


### GetOciEnvironment

`func (o *HypervisorDetailAzureArcResponseModel) GetOciEnvironment() string`

GetOciEnvironment returns the OciEnvironment field if non-nil, zero value otherwise.

### GetOciEnvironmentOk

`func (o *HypervisorDetailAzureArcResponseModel) GetOciEnvironmentOk() (*string, bool)`

GetOciEnvironmentOk returns a tuple with the OciEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciEnvironment

`func (o *HypervisorDetailAzureArcResponseModel) SetOciEnvironment(v string)`

SetOciEnvironment sets OciEnvironment field to given value.


### GetFingerprint

`func (o *HypervisorDetailAzureArcResponseModel) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *HypervisorDetailAzureArcResponseModel) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *HypervisorDetailAzureArcResponseModel) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetId

`func (o *HypervisorDetailAzureArcResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorDetailAzureArcResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorDetailAzureArcResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorDetailAzureArcResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorDetailAzureArcResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorDetailAzureArcResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorDetailAzureArcResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorDetailAzureArcResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorDetailAzureArcResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorDetailAzureArcResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorDetailAzureArcResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorDetailAzureArcResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorDetailAzureArcResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorDetailAzureArcResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorDetailAzureArcResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorDetailAzureArcResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorDetailAzureArcResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorDetailAzureArcResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetConnectionType

`func (o *HypervisorDetailAzureArcResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorDetailAzureArcResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorDetailAzureArcResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetAddresses

`func (o *HypervisorDetailAzureArcResponseModel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorDetailAzureArcResponseModel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorDetailAzureArcResponseModel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetInMaintenanceMode

`func (o *HypervisorDetailAzureArcResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *HypervisorDetailAzureArcResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *HypervisorDetailAzureArcResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetUnentitled

`func (o *HypervisorDetailAzureArcResponseModel) GetUnentitled() bool`

GetUnentitled returns the Unentitled field if non-nil, zero value otherwise.

### GetUnentitledOk

`func (o *HypervisorDetailAzureArcResponseModel) GetUnentitledOk() (*bool, bool)`

GetUnentitledOk returns a tuple with the Unentitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnentitled

`func (o *HypervisorDetailAzureArcResponseModel) SetUnentitled(v bool)`

SetUnentitled sets Unentitled field to given value.

### HasUnentitled

`func (o *HypervisorDetailAzureArcResponseModel) HasUnentitled() bool`

HasUnentitled returns a boolean if a field has been set.

### GetPluginId

`func (o *HypervisorDetailAzureArcResponseModel) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorDetailAzureArcResponseModel) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorDetailAzureArcResponseModel) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.


### GetScopes

`func (o *HypervisorDetailAzureArcResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorDetailAzureArcResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorDetailAzureArcResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *HypervisorDetailAzureArcResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorDetailAzureArcResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorDetailAzureArcResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorDetailAzureArcResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *HypervisorDetailAzureArcResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *HypervisorDetailAzureArcResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetUsesCloudInfrastructure

`func (o *HypervisorDetailAzureArcResponseModel) GetUsesCloudInfrastructure() bool`

GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field if non-nil, zero value otherwise.

### GetUsesCloudInfrastructureOk

`func (o *HypervisorDetailAzureArcResponseModel) GetUsesCloudInfrastructureOk() (*bool, bool)`

GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCloudInfrastructure

`func (o *HypervisorDetailAzureArcResponseModel) SetUsesCloudInfrastructure(v bool)`

SetUsesCloudInfrastructure sets UsesCloudInfrastructure field to given value.


### GetZone

`func (o *HypervisorDetailAzureArcResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorDetailAzureArcResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorDetailAzureArcResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.


### GetFault

`func (o *HypervisorDetailAzureArcResponseModel) GetFault() HypervisorFaultResponseModel`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *HypervisorDetailAzureArcResponseModel) GetFaultOk() (*HypervisorFaultResponseModel, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *HypervisorDetailAzureArcResponseModel) SetFault(v HypervisorFaultResponseModel)`

SetFault sets Fault field to given value.

### HasFault

`func (o *HypervisorDetailAzureArcResponseModel) HasFault() bool`

HasFault returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorDetailAzureArcResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorDetailAzureArcResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorDetailAzureArcResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorDetailAzureArcResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *HypervisorDetailAzureArcResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *HypervisorDetailAzureArcResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetUid

`func (o *HypervisorDetailAzureArcResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *HypervisorDetailAzureArcResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *HypervisorDetailAzureArcResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *HypervisorDetailAzureArcResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *HypervisorDetailAzureArcResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *HypervisorDetailAzureArcResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetIsVirtual

`func (o *HypervisorDetailAzureArcResponseModel) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *HypervisorDetailAzureArcResponseModel) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *HypervisorDetailAzureArcResponseModel) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *HypervisorDetailAzureArcResponseModel) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetCapabilities

`func (o *HypervisorDetailAzureArcResponseModel) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *HypervisorDetailAzureArcResponseModel) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *HypervisorDetailAzureArcResponseModel) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetConfigurationObjectCapabilities

`func (o *HypervisorDetailAzureArcResponseModel) GetConfigurationObjectCapabilities() []HypervisorConfigurationObjectCapability`

GetConfigurationObjectCapabilities returns the ConfigurationObjectCapabilities field if non-nil, zero value otherwise.

### GetConfigurationObjectCapabilitiesOk

`func (o *HypervisorDetailAzureArcResponseModel) GetConfigurationObjectCapabilitiesOk() (*[]HypervisorConfigurationObjectCapability, bool)`

GetConfigurationObjectCapabilitiesOk returns a tuple with the ConfigurationObjectCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationObjectCapabilities

`func (o *HypervisorDetailAzureArcResponseModel) SetConfigurationObjectCapabilities(v []HypervisorConfigurationObjectCapability)`

SetConfigurationObjectCapabilities sets ConfigurationObjectCapabilities field to given value.


### GetPluginRevision

`func (o *HypervisorDetailAzureArcResponseModel) GetPluginRevision() string`

GetPluginRevision returns the PluginRevision field if non-nil, zero value otherwise.

### GetPluginRevisionOk

`func (o *HypervisorDetailAzureArcResponseModel) GetPluginRevisionOk() (*string, bool)`

GetPluginRevisionOk returns a tuple with the PluginRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginRevision

`func (o *HypervisorDetailAzureArcResponseModel) SetPluginRevision(v string)`

SetPluginRevision sets PluginRevision field to given value.


### GetMaxAbsoluteActiveActions

`func (o *HypervisorDetailAzureArcResponseModel) GetMaxAbsoluteActiveActions() int32`

GetMaxAbsoluteActiveActions returns the MaxAbsoluteActiveActions field if non-nil, zero value otherwise.

### GetMaxAbsoluteActiveActionsOk

`func (o *HypervisorDetailAzureArcResponseModel) GetMaxAbsoluteActiveActionsOk() (*int32, bool)`

GetMaxAbsoluteActiveActionsOk returns a tuple with the MaxAbsoluteActiveActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteActiveActions

`func (o *HypervisorDetailAzureArcResponseModel) SetMaxAbsoluteActiveActions(v int32)`

SetMaxAbsoluteActiveActions sets MaxAbsoluteActiveActions field to given value.


### GetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorDetailAzureArcResponseModel) GetMaxAbsoluteNewActionsPerMinute() int32`

GetMaxAbsoluteNewActionsPerMinute returns the MaxAbsoluteNewActionsPerMinute field if non-nil, zero value otherwise.

### GetMaxAbsoluteNewActionsPerMinuteOk

`func (o *HypervisorDetailAzureArcResponseModel) GetMaxAbsoluteNewActionsPerMinuteOk() (*int32, bool)`

GetMaxAbsoluteNewActionsPerMinuteOk returns a tuple with the MaxAbsoluteNewActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorDetailAzureArcResponseModel) SetMaxAbsoluteNewActionsPerMinute(v int32)`

SetMaxAbsoluteNewActionsPerMinute sets MaxAbsoluteNewActionsPerMinute field to given value.


### GetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorDetailAzureArcResponseModel) GetMaxPowerActionsPercentageOfMachines() int32`

GetMaxPowerActionsPercentageOfMachines returns the MaxPowerActionsPercentageOfMachines field if non-nil, zero value otherwise.

### GetMaxPowerActionsPercentageOfMachinesOk

`func (o *HypervisorDetailAzureArcResponseModel) GetMaxPowerActionsPercentageOfMachinesOk() (*int32, bool)`

GetMaxPowerActionsPercentageOfMachinesOk returns a tuple with the MaxPowerActionsPercentageOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorDetailAzureArcResponseModel) SetMaxPowerActionsPercentageOfMachines(v int32)`

SetMaxPowerActionsPercentageOfMachines sets MaxPowerActionsPercentageOfMachines field to given value.


### GetConnectionOptions

`func (o *HypervisorDetailAzureArcResponseModel) GetConnectionOptions() string`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *HypervisorDetailAzureArcResponseModel) GetConnectionOptionsOk() (*string, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *HypervisorDetailAzureArcResponseModel) SetConnectionOptions(v string)`

SetConnectionOptions sets ConnectionOptions field to given value.


### GetSupportsLocalStorageCaching

`func (o *HypervisorDetailAzureArcResponseModel) GetSupportsLocalStorageCaching() bool`

GetSupportsLocalStorageCaching returns the SupportsLocalStorageCaching field if non-nil, zero value otherwise.

### GetSupportsLocalStorageCachingOk

`func (o *HypervisorDetailAzureArcResponseModel) GetSupportsLocalStorageCachingOk() (*bool, bool)`

GetSupportsLocalStorageCachingOk returns a tuple with the SupportsLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLocalStorageCaching

`func (o *HypervisorDetailAzureArcResponseModel) SetSupportsLocalStorageCaching(v bool)`

SetSupportsLocalStorageCaching sets SupportsLocalStorageCaching field to given value.


### GetSupportsPvsVms

`func (o *HypervisorDetailAzureArcResponseModel) GetSupportsPvsVms() bool`

GetSupportsPvsVms returns the SupportsPvsVms field if non-nil, zero value otherwise.

### GetSupportsPvsVmsOk

`func (o *HypervisorDetailAzureArcResponseModel) GetSupportsPvsVmsOk() (*bool, bool)`

GetSupportsPvsVmsOk returns a tuple with the SupportsPvsVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPvsVms

`func (o *HypervisorDetailAzureArcResponseModel) SetSupportsPvsVms(v bool)`

SetSupportsPvsVms sets SupportsPvsVms field to given value.


### GetMetadata

`func (o *HypervisorDetailAzureArcResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorDetailAzureArcResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorDetailAzureArcResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorDetailAzureArcResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HypervisorDetailAzureArcResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HypervisorDetailAzureArcResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


