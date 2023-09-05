# HypervisorVmValidationResponseModelHypervisorConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**XDPath** | Pointer to **string** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**Addresses** | **[]string** | Addresses that can be used to contact the required hypervisor. All the addresses are considered equivalent, that is, all of the addresses provide access to the same virtual machines, snapshots, network, and storage. | 
**InMaintenanceMode** | **bool** | Indicates whether the hypervisor is in maintenance mode, which disables all communication between XenApp &amp; XenDesktop and the Hypervisor. | 
**Unentitled** | Pointer to **bool** | Indicates whether is UnEntitled. | [optional] 
**PluginId** | **string** | The class name for the Citrix Managed Machine library that is used to access the hypervisor. | 
**Scopes** | [**[]ScopeResponseModel**](ScopeResponseModel.md) | The list of administrative scopes that the connection is a part of. The scopes control which administrators are able to work with the connection. | 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the hypervisor is assigned to.  If &#x60;null&#x60;, the hypervisor is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**UsesCloudInfrastructure** | **bool** | Indicates whether the hypervisor uses cloud infrastructure. | 
**Zone** | [**HypervisorResponseModelAllOfZone**](HypervisorResponseModelAllOfZone.md) |  | 
**Fault** | Pointer to [**HypervisorResponseModelAllOfFault**](HypervisorResponseModelAllOfFault.md) |  | [optional] 
**CustomProperties** | Pointer to **string** | CustomProperties of hypervisor connection | [optional] 
**Uid** | Pointer to **int32** | The broker id. | [optional] 
**IsVirtual** | Pointer to **bool** | If this connection is virtual. | [optional] 

## Methods

### NewHypervisorVmValidationResponseModelHypervisorConnection

`func NewHypervisorVmValidationResponseModelHypervisorConnection(connectionType HypervisorConnectionType, addresses []string, inMaintenanceMode bool, pluginId string, scopes []ScopeResponseModel, usesCloudInfrastructure bool, zone HypervisorResponseModelAllOfZone, ) *HypervisorVmValidationResponseModelHypervisorConnection`

NewHypervisorVmValidationResponseModelHypervisorConnection instantiates a new HypervisorVmValidationResponseModelHypervisorConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorVmValidationResponseModelHypervisorConnectionWithDefaults

`func NewHypervisorVmValidationResponseModelHypervisorConnectionWithDefaults() *HypervisorVmValidationResponseModelHypervisorConnection`

NewHypervisorVmValidationResponseModelHypervisorConnectionWithDefaults instantiates a new HypervisorVmValidationResponseModelHypervisorConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetConnectionType

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetAddresses

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetInMaintenanceMode

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetUnentitled

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetUnentitled() bool`

GetUnentitled returns the Unentitled field if non-nil, zero value otherwise.

### GetUnentitledOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetUnentitledOk() (*bool, bool)`

GetUnentitledOk returns a tuple with the Unentitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnentitled

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetUnentitled(v bool)`

SetUnentitled sets Unentitled field to given value.

### HasUnentitled

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) HasUnentitled() bool`

HasUnentitled returns a boolean if a field has been set.

### GetPluginId

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.


### GetScopes

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetUsesCloudInfrastructure

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetUsesCloudInfrastructure() bool`

GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field if non-nil, zero value otherwise.

### GetUsesCloudInfrastructureOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetUsesCloudInfrastructureOk() (*bool, bool)`

GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCloudInfrastructure

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetUsesCloudInfrastructure(v bool)`

SetUsesCloudInfrastructure sets UsesCloudInfrastructure field to given value.


### GetZone

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetZone() HypervisorResponseModelAllOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetZoneOk() (*HypervisorResponseModelAllOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetZone(v HypervisorResponseModelAllOfZone)`

SetZone sets Zone field to given value.


### GetFault

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetFault() HypervisorResponseModelAllOfFault`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetFaultOk() (*HypervisorResponseModelAllOfFault, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetFault(v HypervisorResponseModelAllOfFault)`

SetFault sets Fault field to given value.

### HasFault

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) HasFault() bool`

HasFault returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetUid

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetIsVirtual

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *HypervisorVmValidationResponseModelHypervisorConnection) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


