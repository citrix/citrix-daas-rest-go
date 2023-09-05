# HypervisorsAndResourcePoolsResponseModel

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
**DisplayConnectionType** | **string** | Hypervisor connection connection type display name. | 
**IsBroken** | **bool** | Indicates whether the hypervisor connection is broken. | 
**Errors** | **[]string** | Hypervisor connection broken errors. | 
**ErrorState** | **string** | Current hypervisor connection error status. | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of hypervisor connections. | [optional] 
**ResourcePools** | Pointer to [**[]HypervisorBaseResponseModel**](HypervisorBaseResponseModel.md) | Hypervisor resource pools ref response list. | [optional] 

## Methods

### NewHypervisorsAndResourcePoolsResponseModel

`func NewHypervisorsAndResourcePoolsResponseModel(connectionType HypervisorConnectionType, addresses []string, inMaintenanceMode bool, pluginId string, scopes []ScopeResponseModel, usesCloudInfrastructure bool, zone HypervisorResponseModelAllOfZone, displayConnectionType string, isBroken bool, errors []string, errorState string, ) *HypervisorsAndResourcePoolsResponseModel`

NewHypervisorsAndResourcePoolsResponseModel instantiates a new HypervisorsAndResourcePoolsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorsAndResourcePoolsResponseModelWithDefaults

`func NewHypervisorsAndResourcePoolsResponseModelWithDefaults() *HypervisorsAndResourcePoolsResponseModel`

NewHypervisorsAndResourcePoolsResponseModelWithDefaults instantiates a new HypervisorsAndResourcePoolsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorsAndResourcePoolsResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorsAndResourcePoolsResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorsAndResourcePoolsResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorsAndResourcePoolsResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorsAndResourcePoolsResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorsAndResourcePoolsResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorsAndResourcePoolsResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorsAndResourcePoolsResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorsAndResourcePoolsResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetConnectionType

`func (o *HypervisorsAndResourcePoolsResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorsAndResourcePoolsResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetAddresses

`func (o *HypervisorsAndResourcePoolsResponseModel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorsAndResourcePoolsResponseModel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetInMaintenanceMode

`func (o *HypervisorsAndResourcePoolsResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *HypervisorsAndResourcePoolsResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetUnentitled

`func (o *HypervisorsAndResourcePoolsResponseModel) GetUnentitled() bool`

GetUnentitled returns the Unentitled field if non-nil, zero value otherwise.

### GetUnentitledOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetUnentitledOk() (*bool, bool)`

GetUnentitledOk returns a tuple with the Unentitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnentitled

`func (o *HypervisorsAndResourcePoolsResponseModel) SetUnentitled(v bool)`

SetUnentitled sets Unentitled field to given value.

### HasUnentitled

`func (o *HypervisorsAndResourcePoolsResponseModel) HasUnentitled() bool`

HasUnentitled returns a boolean if a field has been set.

### GetPluginId

`func (o *HypervisorsAndResourcePoolsResponseModel) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorsAndResourcePoolsResponseModel) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.


### GetScopes

`func (o *HypervisorsAndResourcePoolsResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorsAndResourcePoolsResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *HypervisorsAndResourcePoolsResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorsAndResourcePoolsResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorsAndResourcePoolsResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetUsesCloudInfrastructure

`func (o *HypervisorsAndResourcePoolsResponseModel) GetUsesCloudInfrastructure() bool`

GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field if non-nil, zero value otherwise.

### GetUsesCloudInfrastructureOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetUsesCloudInfrastructureOk() (*bool, bool)`

GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCloudInfrastructure

`func (o *HypervisorsAndResourcePoolsResponseModel) SetUsesCloudInfrastructure(v bool)`

SetUsesCloudInfrastructure sets UsesCloudInfrastructure field to given value.


### GetZone

`func (o *HypervisorsAndResourcePoolsResponseModel) GetZone() HypervisorResponseModelAllOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetZoneOk() (*HypervisorResponseModelAllOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorsAndResourcePoolsResponseModel) SetZone(v HypervisorResponseModelAllOfZone)`

SetZone sets Zone field to given value.


### GetFault

`func (o *HypervisorsAndResourcePoolsResponseModel) GetFault() HypervisorResponseModelAllOfFault`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetFaultOk() (*HypervisorResponseModelAllOfFault, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *HypervisorsAndResourcePoolsResponseModel) SetFault(v HypervisorResponseModelAllOfFault)`

SetFault sets Fault field to given value.

### HasFault

`func (o *HypervisorsAndResourcePoolsResponseModel) HasFault() bool`

HasFault returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorsAndResourcePoolsResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorsAndResourcePoolsResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorsAndResourcePoolsResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetUid

`func (o *HypervisorsAndResourcePoolsResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *HypervisorsAndResourcePoolsResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *HypervisorsAndResourcePoolsResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetIsVirtual

`func (o *HypervisorsAndResourcePoolsResponseModel) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *HypervisorsAndResourcePoolsResponseModel) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *HypervisorsAndResourcePoolsResponseModel) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetDisplayConnectionType

`func (o *HypervisorsAndResourcePoolsResponseModel) GetDisplayConnectionType() string`

GetDisplayConnectionType returns the DisplayConnectionType field if non-nil, zero value otherwise.

### GetDisplayConnectionTypeOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetDisplayConnectionTypeOk() (*string, bool)`

GetDisplayConnectionTypeOk returns a tuple with the DisplayConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayConnectionType

`func (o *HypervisorsAndResourcePoolsResponseModel) SetDisplayConnectionType(v string)`

SetDisplayConnectionType sets DisplayConnectionType field to given value.


### GetIsBroken

`func (o *HypervisorsAndResourcePoolsResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *HypervisorsAndResourcePoolsResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.


### GetErrors

`func (o *HypervisorsAndResourcePoolsResponseModel) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HypervisorsAndResourcePoolsResponseModel) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### GetErrorState

`func (o *HypervisorsAndResourcePoolsResponseModel) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *HypervisorsAndResourcePoolsResponseModel) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.


### GetMetadata

`func (o *HypervisorsAndResourcePoolsResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorsAndResourcePoolsResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorsAndResourcePoolsResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResourcePools

`func (o *HypervisorsAndResourcePoolsResponseModel) GetResourcePools() []HypervisorBaseResponseModel`

GetResourcePools returns the ResourcePools field if non-nil, zero value otherwise.

### GetResourcePoolsOk

`func (o *HypervisorsAndResourcePoolsResponseModel) GetResourcePoolsOk() (*[]HypervisorBaseResponseModel, bool)`

GetResourcePoolsOk returns a tuple with the ResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePools

`func (o *HypervisorsAndResourcePoolsResponseModel) SetResourcePools(v []HypervisorBaseResponseModel)`

SetResourcePools sets ResourcePools field to given value.

### HasResourcePools

`func (o *HypervisorsAndResourcePoolsResponseModel) HasResourcePools() bool`

HasResourcePools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


