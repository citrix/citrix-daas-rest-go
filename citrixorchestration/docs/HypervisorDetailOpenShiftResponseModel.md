# HypervisorDetailOpenShiftResponseModel

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
**SslThumbprints** | Pointer to **[]string** | SSL thumbprints considered acceptable for the SSL certificate presented by the hypervisor. | [optional] 
**UserName** | **string** | The user name for the credentials used to communicate with the hypervisor. | 

## Methods

### NewHypervisorDetailOpenShiftResponseModel

`func NewHypervisorDetailOpenShiftResponseModel(connectionType HypervisorConnectionType, addresses []string, inMaintenanceMode bool, pluginId string, scopes []ScopeResponseModel, usesCloudInfrastructure bool, zone RefResponseModel, capabilities []string, configurationObjectCapabilities []HypervisorConfigurationObjectCapability, pluginRevision string, maxAbsoluteActiveActions int32, maxAbsoluteNewActionsPerMinute int32, maxPowerActionsPercentageOfMachines int32, connectionOptions string, supportsLocalStorageCaching bool, supportsPvsVms bool, userName string, ) *HypervisorDetailOpenShiftResponseModel`

NewHypervisorDetailOpenShiftResponseModel instantiates a new HypervisorDetailOpenShiftResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorDetailOpenShiftResponseModelWithDefaults

`func NewHypervisorDetailOpenShiftResponseModelWithDefaults() *HypervisorDetailOpenShiftResponseModel`

NewHypervisorDetailOpenShiftResponseModelWithDefaults instantiates a new HypervisorDetailOpenShiftResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorDetailOpenShiftResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorDetailOpenShiftResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorDetailOpenShiftResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorDetailOpenShiftResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorDetailOpenShiftResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorDetailOpenShiftResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorDetailOpenShiftResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorDetailOpenShiftResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorDetailOpenShiftResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorDetailOpenShiftResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorDetailOpenShiftResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorDetailOpenShiftResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorDetailOpenShiftResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorDetailOpenShiftResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorDetailOpenShiftResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetConnectionType

`func (o *HypervisorDetailOpenShiftResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorDetailOpenShiftResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetAddresses

`func (o *HypervisorDetailOpenShiftResponseModel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorDetailOpenShiftResponseModel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetInMaintenanceMode

`func (o *HypervisorDetailOpenShiftResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *HypervisorDetailOpenShiftResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetUnentitled

`func (o *HypervisorDetailOpenShiftResponseModel) GetUnentitled() bool`

GetUnentitled returns the Unentitled field if non-nil, zero value otherwise.

### GetUnentitledOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetUnentitledOk() (*bool, bool)`

GetUnentitledOk returns a tuple with the Unentitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnentitled

`func (o *HypervisorDetailOpenShiftResponseModel) SetUnentitled(v bool)`

SetUnentitled sets Unentitled field to given value.

### HasUnentitled

`func (o *HypervisorDetailOpenShiftResponseModel) HasUnentitled() bool`

HasUnentitled returns a boolean if a field has been set.

### GetPluginId

`func (o *HypervisorDetailOpenShiftResponseModel) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorDetailOpenShiftResponseModel) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.


### GetScopes

`func (o *HypervisorDetailOpenShiftResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorDetailOpenShiftResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *HypervisorDetailOpenShiftResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorDetailOpenShiftResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorDetailOpenShiftResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *HypervisorDetailOpenShiftResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *HypervisorDetailOpenShiftResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetUsesCloudInfrastructure

`func (o *HypervisorDetailOpenShiftResponseModel) GetUsesCloudInfrastructure() bool`

GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field if non-nil, zero value otherwise.

### GetUsesCloudInfrastructureOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetUsesCloudInfrastructureOk() (*bool, bool)`

GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCloudInfrastructure

`func (o *HypervisorDetailOpenShiftResponseModel) SetUsesCloudInfrastructure(v bool)`

SetUsesCloudInfrastructure sets UsesCloudInfrastructure field to given value.


### GetZone

`func (o *HypervisorDetailOpenShiftResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorDetailOpenShiftResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.


### GetFault

`func (o *HypervisorDetailOpenShiftResponseModel) GetFault() HypervisorFaultResponseModel`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetFaultOk() (*HypervisorFaultResponseModel, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *HypervisorDetailOpenShiftResponseModel) SetFault(v HypervisorFaultResponseModel)`

SetFault sets Fault field to given value.

### HasFault

`func (o *HypervisorDetailOpenShiftResponseModel) HasFault() bool`

HasFault returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorDetailOpenShiftResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorDetailOpenShiftResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorDetailOpenShiftResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *HypervisorDetailOpenShiftResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *HypervisorDetailOpenShiftResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetUid

`func (o *HypervisorDetailOpenShiftResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *HypervisorDetailOpenShiftResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *HypervisorDetailOpenShiftResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *HypervisorDetailOpenShiftResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *HypervisorDetailOpenShiftResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetIsVirtual

`func (o *HypervisorDetailOpenShiftResponseModel) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *HypervisorDetailOpenShiftResponseModel) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *HypervisorDetailOpenShiftResponseModel) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetCapabilities

`func (o *HypervisorDetailOpenShiftResponseModel) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *HypervisorDetailOpenShiftResponseModel) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetConfigurationObjectCapabilities

`func (o *HypervisorDetailOpenShiftResponseModel) GetConfigurationObjectCapabilities() []HypervisorConfigurationObjectCapability`

GetConfigurationObjectCapabilities returns the ConfigurationObjectCapabilities field if non-nil, zero value otherwise.

### GetConfigurationObjectCapabilitiesOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetConfigurationObjectCapabilitiesOk() (*[]HypervisorConfigurationObjectCapability, bool)`

GetConfigurationObjectCapabilitiesOk returns a tuple with the ConfigurationObjectCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationObjectCapabilities

`func (o *HypervisorDetailOpenShiftResponseModel) SetConfigurationObjectCapabilities(v []HypervisorConfigurationObjectCapability)`

SetConfigurationObjectCapabilities sets ConfigurationObjectCapabilities field to given value.


### GetPluginRevision

`func (o *HypervisorDetailOpenShiftResponseModel) GetPluginRevision() string`

GetPluginRevision returns the PluginRevision field if non-nil, zero value otherwise.

### GetPluginRevisionOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetPluginRevisionOk() (*string, bool)`

GetPluginRevisionOk returns a tuple with the PluginRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginRevision

`func (o *HypervisorDetailOpenShiftResponseModel) SetPluginRevision(v string)`

SetPluginRevision sets PluginRevision field to given value.


### GetMaxAbsoluteActiveActions

`func (o *HypervisorDetailOpenShiftResponseModel) GetMaxAbsoluteActiveActions() int32`

GetMaxAbsoluteActiveActions returns the MaxAbsoluteActiveActions field if non-nil, zero value otherwise.

### GetMaxAbsoluteActiveActionsOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetMaxAbsoluteActiveActionsOk() (*int32, bool)`

GetMaxAbsoluteActiveActionsOk returns a tuple with the MaxAbsoluteActiveActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteActiveActions

`func (o *HypervisorDetailOpenShiftResponseModel) SetMaxAbsoluteActiveActions(v int32)`

SetMaxAbsoluteActiveActions sets MaxAbsoluteActiveActions field to given value.


### GetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorDetailOpenShiftResponseModel) GetMaxAbsoluteNewActionsPerMinute() int32`

GetMaxAbsoluteNewActionsPerMinute returns the MaxAbsoluteNewActionsPerMinute field if non-nil, zero value otherwise.

### GetMaxAbsoluteNewActionsPerMinuteOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetMaxAbsoluteNewActionsPerMinuteOk() (*int32, bool)`

GetMaxAbsoluteNewActionsPerMinuteOk returns a tuple with the MaxAbsoluteNewActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteNewActionsPerMinute

`func (o *HypervisorDetailOpenShiftResponseModel) SetMaxAbsoluteNewActionsPerMinute(v int32)`

SetMaxAbsoluteNewActionsPerMinute sets MaxAbsoluteNewActionsPerMinute field to given value.


### GetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorDetailOpenShiftResponseModel) GetMaxPowerActionsPercentageOfMachines() int32`

GetMaxPowerActionsPercentageOfMachines returns the MaxPowerActionsPercentageOfMachines field if non-nil, zero value otherwise.

### GetMaxPowerActionsPercentageOfMachinesOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetMaxPowerActionsPercentageOfMachinesOk() (*int32, bool)`

GetMaxPowerActionsPercentageOfMachinesOk returns a tuple with the MaxPowerActionsPercentageOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPowerActionsPercentageOfMachines

`func (o *HypervisorDetailOpenShiftResponseModel) SetMaxPowerActionsPercentageOfMachines(v int32)`

SetMaxPowerActionsPercentageOfMachines sets MaxPowerActionsPercentageOfMachines field to given value.


### GetConnectionOptions

`func (o *HypervisorDetailOpenShiftResponseModel) GetConnectionOptions() string`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetConnectionOptionsOk() (*string, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *HypervisorDetailOpenShiftResponseModel) SetConnectionOptions(v string)`

SetConnectionOptions sets ConnectionOptions field to given value.


### GetSupportsLocalStorageCaching

`func (o *HypervisorDetailOpenShiftResponseModel) GetSupportsLocalStorageCaching() bool`

GetSupportsLocalStorageCaching returns the SupportsLocalStorageCaching field if non-nil, zero value otherwise.

### GetSupportsLocalStorageCachingOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetSupportsLocalStorageCachingOk() (*bool, bool)`

GetSupportsLocalStorageCachingOk returns a tuple with the SupportsLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLocalStorageCaching

`func (o *HypervisorDetailOpenShiftResponseModel) SetSupportsLocalStorageCaching(v bool)`

SetSupportsLocalStorageCaching sets SupportsLocalStorageCaching field to given value.


### GetSupportsPvsVms

`func (o *HypervisorDetailOpenShiftResponseModel) GetSupportsPvsVms() bool`

GetSupportsPvsVms returns the SupportsPvsVms field if non-nil, zero value otherwise.

### GetSupportsPvsVmsOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetSupportsPvsVmsOk() (*bool, bool)`

GetSupportsPvsVmsOk returns a tuple with the SupportsPvsVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPvsVms

`func (o *HypervisorDetailOpenShiftResponseModel) SetSupportsPvsVms(v bool)`

SetSupportsPvsVms sets SupportsPvsVms field to given value.


### GetMetadata

`func (o *HypervisorDetailOpenShiftResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorDetailOpenShiftResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorDetailOpenShiftResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HypervisorDetailOpenShiftResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HypervisorDetailOpenShiftResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSslThumbprints

`func (o *HypervisorDetailOpenShiftResponseModel) GetSslThumbprints() []string`

GetSslThumbprints returns the SslThumbprints field if non-nil, zero value otherwise.

### GetSslThumbprintsOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetSslThumbprintsOk() (*[]string, bool)`

GetSslThumbprintsOk returns a tuple with the SslThumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprints

`func (o *HypervisorDetailOpenShiftResponseModel) SetSslThumbprints(v []string)`

SetSslThumbprints sets SslThumbprints field to given value.

### HasSslThumbprints

`func (o *HypervisorDetailOpenShiftResponseModel) HasSslThumbprints() bool`

HasSslThumbprints returns a boolean if a field has been set.

### SetSslThumbprintsNil

`func (o *HypervisorDetailOpenShiftResponseModel) SetSslThumbprintsNil(b bool)`

 SetSslThumbprintsNil sets the value for SslThumbprints to be an explicit nil

### UnsetSslThumbprints
`func (o *HypervisorDetailOpenShiftResponseModel) UnsetSslThumbprints()`

UnsetSslThumbprints ensures that no value is present for SslThumbprints, not even an explicit nil
### GetUserName

`func (o *HypervisorDetailOpenShiftResponseModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HypervisorDetailOpenShiftResponseModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HypervisorDetailOpenShiftResponseModel) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


