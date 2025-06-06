/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the HypervisorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorResponseModel{}

// HypervisorResponseModel struct for HypervisorResponseModel
type HypervisorResponseModel struct {
	// Id of the resource.
	Id NullableString `json:"Id,omitempty"`
	// Name of the resource.
	Name NullableString `json:"Name,omitempty"`
	// XenApp & XenDesktop path to the resource on the hypervisor.  An example value is: `XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot` or `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}`
	XDPath         NullableString           `json:"XDPath,omitempty"`
	ConnectionType HypervisorConnectionType `json:"ConnectionType"`
	// Addresses that can be used to contact the required hypervisor. All the addresses are considered equivalent, that is, all of the addresses provide access to the same virtual machines, snapshots, network, and storage.
	Addresses []string `json:"Addresses"`
	// Indicates whether the hypervisor is in maintenance mode, which disables all communication between XenApp & XenDesktop and the Hypervisor.
	InMaintenanceMode bool `json:"InMaintenanceMode"`
	// Indicates whether is UnEntitled.
	Unentitled *bool `json:"Unentitled,omitempty"`
	// The class name for the Citrix Managed Machine library that is used to access the hypervisor.
	PluginId string `json:"PluginId"`
	// The list of administrative scopes that the connection is a part of. The scopes control which administrators are able to work with the connection.
	Scopes []ScopeResponseModel `json:"Scopes"`
	// The tenant(s) that the hypervisor is assigned to.  If `null`, the hypervisor is not assigned to tenants, and may be used by any tenant, including future added tenants.
	Tenants []RefResponseModel `json:"Tenants,omitempty"`
	// Indicates whether the hypervisor uses cloud infrastructure.
	UsesCloudInfrastructure bool                          `json:"UsesCloudInfrastructure"`
	Zone                    RefResponseModel              `json:"Zone"`
	Fault                   *HypervisorFaultResponseModel `json:"Fault,omitempty"`
	// CustomProperties of hypervisor connection
	CustomProperties NullableString `json:"CustomProperties,omitempty"`
	// The broker id.
	Uid NullableInt32 `json:"Uid,omitempty"`
	// If this connection is virtual.
	IsVirtual *bool `json:"IsVirtual,omitempty"`
}

// NewHypervisorResponseModel instantiates a new HypervisorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorResponseModel(connectionType HypervisorConnectionType, addresses []string, inMaintenanceMode bool, pluginId string, scopes []ScopeResponseModel, usesCloudInfrastructure bool, zone RefResponseModel) *HypervisorResponseModel {
	this := HypervisorResponseModel{}
	this.ConnectionType = connectionType
	this.Addresses = addresses
	this.InMaintenanceMode = inMaintenanceMode
	this.PluginId = pluginId
	this.Scopes = scopes
	this.UsesCloudInfrastructure = usesCloudInfrastructure
	this.Zone = zone
	return &this
}

// NewHypervisorResponseModelWithDefaults instantiates a new HypervisorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorResponseModelWithDefaults() *HypervisorResponseModel {
	this := HypervisorResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResponseModel) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorResponseModel) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HypervisorResponseModel) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *HypervisorResponseModel) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HypervisorResponseModel) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HypervisorResponseModel) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *HypervisorResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HypervisorResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetXDPath returns the XDPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResponseModel) GetXDPath() string {
	if o == nil || IsNil(o.XDPath.Get()) {
		var ret string
		return ret
	}
	return *o.XDPath.Get()
}

// GetXDPathOk returns a tuple with the XDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResponseModel) GetXDPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.XDPath.Get(), o.XDPath.IsSet()
}

// HasXDPath returns a boolean if a field has been set.
func (o *HypervisorResponseModel) HasXDPath() bool {
	if o != nil && o.XDPath.IsSet() {
		return true
	}

	return false
}

// SetXDPath gets a reference to the given NullableString and assigns it to the XDPath field.
func (o *HypervisorResponseModel) SetXDPath(v string) {
	o.XDPath.Set(&v)
}

// SetXDPathNil sets the value for XDPath to be an explicit nil
func (o *HypervisorResponseModel) SetXDPathNil() {
	o.XDPath.Set(nil)
}

// UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
func (o *HypervisorResponseModel) UnsetXDPath() {
	o.XDPath.Unset()
}

// GetConnectionType returns the ConnectionType field value
func (o *HypervisorResponseModel) GetConnectionType() HypervisorConnectionType {
	if o == nil {
		var ret HypervisorConnectionType
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *HypervisorResponseModel) SetConnectionType(v HypervisorConnectionType) {
	o.ConnectionType = v
}

// GetAddresses returns the Addresses field value
func (o *HypervisorResponseModel) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *HypervisorResponseModel) SetAddresses(v []string) {
	o.Addresses = v
}

// GetInMaintenanceMode returns the InMaintenanceMode field value
func (o *HypervisorResponseModel) GetInMaintenanceMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InMaintenanceMode
}

// GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field value
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetInMaintenanceModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InMaintenanceMode, true
}

// SetInMaintenanceMode sets field value
func (o *HypervisorResponseModel) SetInMaintenanceMode(v bool) {
	o.InMaintenanceMode = v
}

// GetUnentitled returns the Unentitled field value if set, zero value otherwise.
func (o *HypervisorResponseModel) GetUnentitled() bool {
	if o == nil || IsNil(o.Unentitled) {
		var ret bool
		return ret
	}
	return *o.Unentitled
}

// GetUnentitledOk returns a tuple with the Unentitled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetUnentitledOk() (*bool, bool) {
	if o == nil || IsNil(o.Unentitled) {
		return nil, false
	}
	return o.Unentitled, true
}

// HasUnentitled returns a boolean if a field has been set.
func (o *HypervisorResponseModel) HasUnentitled() bool {
	if o != nil && !IsNil(o.Unentitled) {
		return true
	}

	return false
}

// SetUnentitled gets a reference to the given bool and assigns it to the Unentitled field.
func (o *HypervisorResponseModel) SetUnentitled(v bool) {
	o.Unentitled = &v
}

// GetPluginId returns the PluginId field value
func (o *HypervisorResponseModel) GetPluginId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetPluginIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginId, true
}

// SetPluginId sets field value
func (o *HypervisorResponseModel) SetPluginId(v string) {
	o.PluginId = v
}

// GetScopes returns the Scopes field value
func (o *HypervisorResponseModel) GetScopes() []ScopeResponseModel {
	if o == nil {
		var ret []ScopeResponseModel
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetScopesOk() ([]ScopeResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *HypervisorResponseModel) SetScopes(v []ScopeResponseModel) {
	o.Scopes = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResponseModel) GetTenants() []RefResponseModel {
	if o == nil {
		var ret []RefResponseModel
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResponseModel) GetTenantsOk() ([]RefResponseModel, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *HypervisorResponseModel) HasTenants() bool {
	if o != nil && IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []RefResponseModel and assigns it to the Tenants field.
func (o *HypervisorResponseModel) SetTenants(v []RefResponseModel) {
	o.Tenants = v
}

// GetUsesCloudInfrastructure returns the UsesCloudInfrastructure field value
func (o *HypervisorResponseModel) GetUsesCloudInfrastructure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UsesCloudInfrastructure
}

// GetUsesCloudInfrastructureOk returns a tuple with the UsesCloudInfrastructure field value
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetUsesCloudInfrastructureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsesCloudInfrastructure, true
}

// SetUsesCloudInfrastructure sets field value
func (o *HypervisorResponseModel) SetUsesCloudInfrastructure(v bool) {
	o.UsesCloudInfrastructure = v
}

// GetZone returns the Zone field value
func (o *HypervisorResponseModel) GetZone() RefResponseModel {
	if o == nil {
		var ret RefResponseModel
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetZoneOk() (*RefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *HypervisorResponseModel) SetZone(v RefResponseModel) {
	o.Zone = v
}

// GetFault returns the Fault field value if set, zero value otherwise.
func (o *HypervisorResponseModel) GetFault() HypervisorFaultResponseModel {
	if o == nil || IsNil(o.Fault) {
		var ret HypervisorFaultResponseModel
		return ret
	}
	return *o.Fault
}

// GetFaultOk returns a tuple with the Fault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetFaultOk() (*HypervisorFaultResponseModel, bool) {
	if o == nil || IsNil(o.Fault) {
		return nil, false
	}
	return o.Fault, true
}

// HasFault returns a boolean if a field has been set.
func (o *HypervisorResponseModel) HasFault() bool {
	if o != nil && !IsNil(o.Fault) {
		return true
	}

	return false
}

// SetFault gets a reference to the given HypervisorFaultResponseModel and assigns it to the Fault field.
func (o *HypervisorResponseModel) SetFault(v HypervisorFaultResponseModel) {
	o.Fault = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResponseModel) GetCustomProperties() string {
	if o == nil || IsNil(o.CustomProperties.Get()) {
		var ret string
		return ret
	}
	return *o.CustomProperties.Get()
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResponseModel) GetCustomPropertiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomProperties.Get(), o.CustomProperties.IsSet()
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *HypervisorResponseModel) HasCustomProperties() bool {
	if o != nil && o.CustomProperties.IsSet() {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given NullableString and assigns it to the CustomProperties field.
func (o *HypervisorResponseModel) SetCustomProperties(v string) {
	o.CustomProperties.Set(&v)
}

// SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil
func (o *HypervisorResponseModel) SetCustomPropertiesNil() {
	o.CustomProperties.Set(nil)
}

// UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
func (o *HypervisorResponseModel) UnsetCustomProperties() {
	o.CustomProperties.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResponseModel) GetUid() int32 {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret int32
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResponseModel) GetUidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *HypervisorResponseModel) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableInt32 and assigns it to the Uid field.
func (o *HypervisorResponseModel) SetUid(v int32) {
	o.Uid.Set(&v)
}

// SetUidNil sets the value for Uid to be an explicit nil
func (o *HypervisorResponseModel) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *HypervisorResponseModel) UnsetUid() {
	o.Uid.Unset()
}

// GetIsVirtual returns the IsVirtual field value if set, zero value otherwise.
func (o *HypervisorResponseModel) GetIsVirtual() bool {
	if o == nil || IsNil(o.IsVirtual) {
		var ret bool
		return ret
	}
	return *o.IsVirtual
}

// GetIsVirtualOk returns a tuple with the IsVirtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorResponseModel) GetIsVirtualOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVirtual) {
		return nil, false
	}
	return o.IsVirtual, true
}

// HasIsVirtual returns a boolean if a field has been set.
func (o *HypervisorResponseModel) HasIsVirtual() bool {
	if o != nil && !IsNil(o.IsVirtual) {
		return true
	}

	return false
}

// SetIsVirtual gets a reference to the given bool and assigns it to the IsVirtual field.
func (o *HypervisorResponseModel) SetIsVirtual(v bool) {
	o.IsVirtual = &v
}

func (o HypervisorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.XDPath.IsSet() {
		toSerialize["XDPath"] = o.XDPath.Get()
	}
	toSerialize["ConnectionType"] = o.ConnectionType
	toSerialize["Addresses"] = o.Addresses
	toSerialize["InMaintenanceMode"] = o.InMaintenanceMode
	if !IsNil(o.Unentitled) {
		toSerialize["Unentitled"] = o.Unentitled
	}
	toSerialize["PluginId"] = o.PluginId
	toSerialize["Scopes"] = o.Scopes
	if o.Tenants != nil {
		toSerialize["Tenants"] = o.Tenants
	}
	toSerialize["UsesCloudInfrastructure"] = o.UsesCloudInfrastructure
	toSerialize["Zone"] = o.Zone
	if !IsNil(o.Fault) {
		toSerialize["Fault"] = o.Fault
	}
	if o.CustomProperties.IsSet() {
		toSerialize["CustomProperties"] = o.CustomProperties.Get()
	}
	if o.Uid.IsSet() {
		toSerialize["Uid"] = o.Uid.Get()
	}
	if !IsNil(o.IsVirtual) {
		toSerialize["IsVirtual"] = o.IsVirtual
	}
	return toSerialize, nil
}

type NullableHypervisorResponseModel struct {
	value *HypervisorResponseModel
	isSet bool
}

func (v NullableHypervisorResponseModel) Get() *HypervisorResponseModel {
	return v.value
}

func (v *NullableHypervisorResponseModel) Set(val *HypervisorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorResponseModel(val *HypervisorResponseModel) *NullableHypervisorResponseModel {
	return &NullableHypervisorResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
