# HypervisorDetailAWCResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**ApiKey** | **string** | The API key used to authenticate with the AWS APIs. | 
**Region** | **string** | The AWS region which the hypervisor represents. | 
**MaximumConcurrentProvisioningOperations** | Pointer to **NullableInt32** | AWS maximum concurrent provisioning operations. | [optional] 

## Methods

### NewHypervisorDetailAWCResponseModel

`func NewHypervisorDetailAWCResponseModel(connectionType HypervisorConnectionType, addresses []string, inMaintenanceMode bool, pluginId string, scopes []ScopeResponseModel, usesCloudInfrastructure bool, zone RefResponseModel, capabilities []string, configurationObjectCapabilities []HypervisorConfigurationObjectCapability, pluginRevision string, maxAbsoluteActiveActions int32, maxAbsoluteNewActionsPerMinute int32, maxPowerActionsPercentageOfMachines int32, connectionOptions string, supportsLocalStorageCaching bool, supportsPvsVms bool, apiKey string, region string, ) *HypervisorDetailAWCResponseModel`

NewHypervisorDetailAWCResponseModel instantiates a new HypervisorDetailAWCResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorDetailAWCResponseModelWithDefaults

`func NewHypervisorDetailAWCResponseModelWithDefaults() *HypervisorDetailAWCResponseModel`

NewHypervisorDetailAWCResponseModelWithDefaults instantiates a new HypervisorDetailAWCResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorDetailAWCResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorDetailAWCResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorDetailAWCResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorDetailAWCResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorDetailAWCResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorDetailAWCResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorDetailAWCResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorDetailAWCResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorDetailAWCResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorDetailAWCResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorDetailAWCResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorDetailAWCResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorDetailAWCResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorDetailAWCResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorDetailAWCResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorDetailAWCResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorDetailAWCResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorDetailAWCResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetConnectionType

`func (o *HypervisorDetailAWCResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorDetailAWCResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorDetailAWCResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetAddresses

`func (o *HypervisorDetailAWCResponseModel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorDetailAWCResponseModel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorDetailAWCResponseModel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetInMaintenanceMode

`func (o *HypervisorDetailAWCResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *HypervisorDetailAWCResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *HypervisorDetailAWCResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetUnentitled

`func (o *HypervisorDetailAWCResponseModel) GetUnentitled() bool`

GetUnentitled returns the Unentitled field if non-nil, zero value otherwise.

### GetUnentitledOk

`func (o *HypervisorDetailAWCResponseModel) GetUnentitledOk() (*bool, bool)`

GetUnentitledOk returns a tuple with the Unentitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnentitled

`func (o *HypervisorDetailAWCResponseModel) SetUnentitled(v bool)`

SetUnentitled sets Unentitled field to given value.

### HasUnentitled

`func (o *HypervisorDetailAWCResponseModel) HasUnentitled() bool`

HasUnentitled returns a boolean if a field has been set.

### GetPluginId

`func (o *HypervisorDetailAWCResponseModel) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorDetailAWCResponseModel) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorDetailAWCResponseModel) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.


### GetScopes

`func (o *HypervisorDetailAWCResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorDetailAWCResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorDetailAWCResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *HypervisorDetailAWCResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorDetailAWCResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorDetailAWCResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorDetailAWCResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *HypervisorDetailAWCResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *HypervisorDetailAWCResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetUsesCloudInfrastructure

`func (o *HypervisorDetailAWCResponseModel) GetUsesCloudInfrastructure() bool`

GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field if non-nil, zero value otherwise.

### GetUsesCloudInfrastructureOk

`func (o *HypervisorDetailAWCResponseModel) GetUsesCloudInfrastructureOk() (*bool, bool)`

GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCloudInfrastructure

`func (o *HypervisorDetailAWCResponseModel) SetUsesCloudInfrastructure(v bool)`

SetUsesCloudInfrastructure sets UsesCloudInfrastructure field to given value.


### GetZone

`func (o *HypervisorDetailAWCResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorDetailAWCResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorDetailAWCResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.


### GetFault

`func (o *HypervisorDetailAWCResponseModel) GetFault() HypervisorFaultResponseModel`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *HypervisorDetailAWCResponseModel) GetFaultOk() (*HypervisorFaultResponseModel, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *HypervisorDetailAWCResponseModel) SetFault(v HypervisorFaultResponseModel)`

SetFault sets Fault field to given value.

### HasFault

`func (o *HypervisorDetailAWCResponseModel) HasFault() bool`

HasFault returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorDetailAWCResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorDetailAWCResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorDetailAWCResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorDetailAWCResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *HypervisorDetailAWCResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *HypervisorDetailAWCResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetUid

`func (o *HypervisorDetailAWCResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *HypervisorDetailAWCResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *HypervisorDetailAWCResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *HypervisorDetailAWCResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *HypervisorDetailAWCResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *HypervisorDetailAWCResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetIsVirtual

`func (o *HypervisorDetailAWCResponseModel) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *HypervisorDetailAWCResponseModel) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *HypervisorDetailAWCResponseModel) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *HypervisorDetailAWCResponseModel) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetCapabilities

`func (o *HypervisorDetailAWCResponseModel) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *HypervisorDetailAWCResponseModel) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *HypervisorDetailAWCResponseModel) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetConfigurationObjectCapabilities

`func (o *HypervisorDetailAWCResponseModel) GetConfigurationObjectCapabilities() []HypervisorConfigurationObjectCapability`

GetConfigurationObjectCapabilities returns the ConfigurationObjectCapabilities field if non-nil, zero value otherwise.

### GetConfigurationObjectCapabilitiesOk

`func (o *HypervisorDetailAWCResponseModel) GetConfigurationObjectCapabilitiesOk() (*[]HypervisorConfigurationObjectCapability, bool)`

GetConfigurationObjectCapabilitiesOk returns a tuple with the ConfigurationObjectCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationObjectCapabilities

`func (o *HypervisorDetailAWCResponseModel) SetConfigurationObjectCapabilities(v []HypervisorConfigurationObjectCapability)`

SetConfigurationObjectCapabilities sets ConfigurationObjectCapabilities field to given value.


### GetPluginRevision

`func (o *HypervisorDetailAWCResponseModel) GetPluginRevision() string`

GetPluginRevision returns the PluginRevision field if non-nil, zero value otherwise.

### GetPluginRevisionOk

`func (o *HypervisorDetailAWCResponseModel) GetPluginRevisionOk() (*string, bool)`

GetPluginRevisionOk returns a tuple with the PluginRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginRevision

`func (o *HypervisorDetailAWCResponseModel) SetPluginRevision(v string)`

SetPluginRevision sets PluginRevision field to given value.


### GetMaxAbsoluteActiveActions

`func (o *HypervisorDetailAWCResponseModel) GetMaxAbsoluteActiveActions() int32`

GetMaxAbsoluteActiveActions returns the MaxAbsoluteActiveActions field if non-nil, zero value otherwise.

### GetMaxAbsoluteActiveActionsOk

`func (o *HypervisorDetailAWCResponseModel) GetMaxAbsoluteActiveActionsOk() (*int32, bool)`

GetMaxAbsoluteActiveActionsOk returns a tuple with the MaxAbsoluteActiveActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteActiveActions

`func (o *HypervisorDetailAWCResponseModel) SetMaxAbsoluteActiveActions(v int32)`

SetMaxAbsoluteActiveActions sets MaxAbsoluteActiveActions field to given value.


### GetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorDetailAWCResponseModel) GetMaxAbsoluteNewActionsPerMinute() int32`

GetMaxAbsoluteNewActionsPerMinute returns the MaxAbsoluteNewActionsPerMinute field if non-nil, zero value otherwise.

### GetMaxAbsoluteNewActionsPerMinuteOk

`func (o *HypervisorDetailAWCResponseModel) GetMaxAbsoluteNewActionsPerMinuteOk() (*int32, bool)`

GetMaxAbsoluteNewActionsPerMinuteOk returns a tuple with the MaxAbsoluteNewActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorDetailAWCResponseModel) SetMaxAbsoluteNewActionsPerMinute(v int32)`

SetMaxAbsoluteNewActionsPerMinute sets MaxAbsoluteNewActionsPerMinute field to given value.


### GetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorDetailAWCResponseModel) GetMaxPowerActionsPercentageOfMachines() int32`

GetMaxPowerActionsPercentageOfMachines returns the MaxPowerActionsPercentageOfMachines field if non-nil, zero value otherwise.

### GetMaxPowerActionsPercentageOfMachinesOk

`func (o *HypervisorDetailAWCResponseModel) GetMaxPowerActionsPercentageOfMachinesOk() (*int32, bool)`

GetMaxPowerActionsPercentageOfMachinesOk returns a tuple with the MaxPowerActionsPercentageOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorDetailAWCResponseModel) SetMaxPowerActionsPercentageOfMachines(v int32)`

SetMaxPowerActionsPercentageOfMachines sets MaxPowerActionsPercentageOfMachines field to given value.


### GetConnectionOptions

`func (o *HypervisorDetailAWCResponseModel) GetConnectionOptions() string`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *HypervisorDetailAWCResponseModel) GetConnectionOptionsOk() (*string, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *HypervisorDetailAWCResponseModel) SetConnectionOptions(v string)`

SetConnectionOptions sets ConnectionOptions field to given value.


### GetSupportsLocalStorageCaching

`func (o *HypervisorDetailAWCResponseModel) GetSupportsLocalStorageCaching() bool`

GetSupportsLocalStorageCaching returns the SupportsLocalStorageCaching field if non-nil, zero value otherwise.

### GetSupportsLocalStorageCachingOk

`func (o *HypervisorDetailAWCResponseModel) GetSupportsLocalStorageCachingOk() (*bool, bool)`

GetSupportsLocalStorageCachingOk returns a tuple with the SupportsLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLocalStorageCaching

`func (o *HypervisorDetailAWCResponseModel) SetSupportsLocalStorageCaching(v bool)`

SetSupportsLocalStorageCaching sets SupportsLocalStorageCaching field to given value.


### GetSupportsPvsVms

`func (o *HypervisorDetailAWCResponseModel) GetSupportsPvsVms() bool`

GetSupportsPvsVms returns the SupportsPvsVms field if non-nil, zero value otherwise.

### GetSupportsPvsVmsOk

`func (o *HypervisorDetailAWCResponseModel) GetSupportsPvsVmsOk() (*bool, bool)`

GetSupportsPvsVmsOk returns a tuple with the SupportsPvsVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPvsVms

`func (o *HypervisorDetailAWCResponseModel) SetSupportsPvsVms(v bool)`

SetSupportsPvsVms sets SupportsPvsVms field to given value.


### GetMetadata

`func (o *HypervisorDetailAWCResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorDetailAWCResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorDetailAWCResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorDetailAWCResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HypervisorDetailAWCResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HypervisorDetailAWCResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetApiKey

`func (o *HypervisorDetailAWCResponseModel) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HypervisorDetailAWCResponseModel) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HypervisorDetailAWCResponseModel) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetRegion

`func (o *HypervisorDetailAWCResponseModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorDetailAWCResponseModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorDetailAWCResponseModel) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetMaximumConcurrentProvisioningOperations

`func (o *HypervisorDetailAWCResponseModel) GetMaximumConcurrentProvisioningOperations() int32`

GetMaximumConcurrentProvisioningOperations returns the MaximumConcurrentProvisioningOperations field if non-nil, zero value otherwise.

### GetMaximumConcurrentProvisioningOperationsOk

`func (o *HypervisorDetailAWCResponseModel) GetMaximumConcurrentProvisioningOperationsOk() (*int32, bool)`

GetMaximumConcurrentProvisioningOperationsOk returns a tuple with the MaximumConcurrentProvisioningOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentProvisioningOperations

`func (o *HypervisorDetailAWCResponseModel) SetMaximumConcurrentProvisioningOperations(v int32)`

SetMaximumConcurrentProvisioningOperations sets MaximumConcurrentProvisioningOperations field to given value.

### HasMaximumConcurrentProvisioningOperations

`func (o *HypervisorDetailAWCResponseModel) HasMaximumConcurrentProvisioningOperations() bool`

HasMaximumConcurrentProvisioningOperations returns a boolean if a field has been set.

### SetMaximumConcurrentProvisioningOperationsNil

`func (o *HypervisorDetailAWCResponseModel) SetMaximumConcurrentProvisioningOperationsNil(b bool)`

 SetMaximumConcurrentProvisioningOperationsNil sets the value for MaximumConcurrentProvisioningOperations to be an explicit nil

### UnsetMaximumConcurrentProvisioningOperations
`func (o *HypervisorDetailAWCResponseModel) UnsetMaximumConcurrentProvisioningOperations()`

UnsetMaximumConcurrentProvisioningOperations ensures that no value is present for MaximumConcurrentProvisioningOperations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


