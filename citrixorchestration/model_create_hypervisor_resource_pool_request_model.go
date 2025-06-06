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

// checks if the CreateHypervisorResourcePoolRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateHypervisorResourcePoolRequestModel{}

// CreateHypervisorResourcePoolRequestModel Create a hypervisor resource pool.  Note that in previous versions of the API, a \"resource pool\" was known as a \"hosting unit\".
type CreateHypervisorResourcePoolRequestModel struct {
	// Name of the resource pool to create.  Required.
	Name string `json:"Name" validate:"regexp=(.*)*"`
	// Indicates whether VMs created by Virtual Apps & Desktops provisioning operations should be tagged.  Tagged VMs are filtered out of queries by default. Optional.  Default is `true`.
	VmTagging NullableBool `json:"VmTagging,omitempty"`
	// Path to the GPU type resource(s) that are available for provisioning operations in this resource pool.  Optional.  Not supported by all hypervisor types.
	GpuTypes       []string                 `json:"GpuTypes,omitempty"`
	ConnectionType HypervisorConnectionType `json:"ConnectionType"`
	// Metadata of the resource pool. Optional.
	Metadata []NameValueStringPairModel `json:"Metadata,omitempty"`
	// AWS Virtual Private Cloud (VPC) resource which the resource pool is connected to.  Required.
	VirtualPrivateCloud *string "json:\"VirtualPrivateCloud,omitempty\" validate:\"regexp=(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})Connections(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})HostingUnits(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)hypervisors(\\\\\\\\{1,2}|\\/{1,2}).*|^[^;:#\\\\*\\\\?=<>\\\\|\\\\[\\\\]\\\\(\\\\)\\\"'\\\\{\\\\}`]*.*\""
	// Path to the availability zone resource to use for provisioning operations in this resource pool.  Required.
	AvailabilityZone *string "json:\"AvailabilityZone,omitempty\" validate:\"regexp=(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})Connections(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})HostingUnits(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)hypervisors(\\\\\\\\{1,2}|\\/{1,2}).*|^[^;:#\\\\*\\\\?=<>\\\\|\\\\[\\\\]\\\\(\\\\)\\\"'\\\\{\\\\}`]*.*\""
	// Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required.
	Networks []string `json:"Networks,omitempty"`
	// Azure region which the resource pool is connected to.  Required.
	Region *string "json:\"Region,omitempty\" validate:\"regexp=(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})Connections(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})HostingUnits(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)hypervisors(\\\\\\\\{1,2}|\\/{1,2}).*|^[^;:#\\\\*\\\\?=<>\\\\|\\\\[\\\\]\\\\(\\\\)\\\"'\\\\{\\\\}`]*.*\""
	// Azure virtual network which the resource pool is connected to. Required.
	VirtualNetwork *string "json:\"VirtualNetwork,omitempty\" validate:\"regexp=(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})Connections(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})HostingUnits(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)hypervisors(\\\\\\\\{1,2}|\\/{1,2}).*|^[^;:#\\\\*\\\\?=<>\\\\|\\\\[\\\\]\\\\(\\\\)\\\"'\\\\{\\\\}`]*.*\""
	// Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required.
	Subnets []string `json:"Subnets,omitempty"`
	// Root path of the resources on the hypervisor which should be included in the resource pool.  Required.
	RootPath NullableString "json:\"RootPath,omitempty\" validate:\"regexp=(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})Connections(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)XDHyp:(\\\\\\\\{1,2}|\\/{1,2})HostingUnits(\\\\\\\\{1,2}|\\/{1,2}).*|(?i)hypervisors(\\\\\\\\{1,2}|\\/{1,2}).*|^[^;:#\\\\*\\\\?=<>\\\\|\\\\[\\\\]\\\\(\\\\)\\\"'\\\\{\\\\}`]*.*\""
	// Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to `false`.
	UseLocalStorageCaching NullableBool `json:"UseLocalStorageCaching,omitempty"`
	// Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types.
	Storage []string `json:"Storage,omitempty"`
	// Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types.
	PersonalvDiskStorage []string `json:"PersonalvDiskStorage,omitempty"`
	// Path to the storage resource(s) that are used for temporary operations in this resource pool.  Required for some hypervisor types.
	TemporaryStorage []string `json:"TemporaryStorage,omitempty"`
	// Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom.
	CustomProperties   NullableString      `json:"CustomProperties,omitempty" validate:"regexp=((.|\\\\n)*)"`
	StorageBalanceType *StorageBalanceType `json:"StorageBalanceType,omitempty"`
}

// NewCreateHypervisorResourcePoolRequestModel instantiates a new CreateHypervisorResourcePoolRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateHypervisorResourcePoolRequestModel(name string, connectionType HypervisorConnectionType) *CreateHypervisorResourcePoolRequestModel {
	this := CreateHypervisorResourcePoolRequestModel{}
	this.Name = name
	var vmTagging bool = true
	this.VmTagging = *NewNullableBool(&vmTagging)
	this.ConnectionType = connectionType
	var useLocalStorageCaching bool = false
	this.UseLocalStorageCaching = *NewNullableBool(&useLocalStorageCaching)
	return &this
}

// NewCreateHypervisorResourcePoolRequestModelWithDefaults instantiates a new CreateHypervisorResourcePoolRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateHypervisorResourcePoolRequestModelWithDefaults() *CreateHypervisorResourcePoolRequestModel {
	this := CreateHypervisorResourcePoolRequestModel{}
	var vmTagging bool = true
	this.VmTagging = *NewNullableBool(&vmTagging)
	var useLocalStorageCaching bool = false
	this.UseLocalStorageCaching = *NewNullableBool(&useLocalStorageCaching)
	return &this
}

// GetName returns the Name field value
func (o *CreateHypervisorResourcePoolRequestModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateHypervisorResourcePoolRequestModel) SetName(v string) {
	o.Name = v
}

// GetVmTagging returns the VmTagging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateHypervisorResourcePoolRequestModel) GetVmTagging() bool {
	if o == nil || IsNil(o.VmTagging.Get()) {
		var ret bool
		return ret
	}
	return *o.VmTagging.Get()
}

// GetVmTaggingOk returns a tuple with the VmTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHypervisorResourcePoolRequestModel) GetVmTaggingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmTagging.Get(), o.VmTagging.IsSet()
}

// HasVmTagging returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasVmTagging() bool {
	if o != nil && o.VmTagging.IsSet() {
		return true
	}

	return false
}

// SetVmTagging gets a reference to the given NullableBool and assigns it to the VmTagging field.
func (o *CreateHypervisorResourcePoolRequestModel) SetVmTagging(v bool) {
	o.VmTagging.Set(&v)
}

// SetVmTaggingNil sets the value for VmTagging to be an explicit nil
func (o *CreateHypervisorResourcePoolRequestModel) SetVmTaggingNil() {
	o.VmTagging.Set(nil)
}

// UnsetVmTagging ensures that no value is present for VmTagging, not even an explicit nil
func (o *CreateHypervisorResourcePoolRequestModel) UnsetVmTagging() {
	o.VmTagging.Unset()
}

// GetGpuTypes returns the GpuTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateHypervisorResourcePoolRequestModel) GetGpuTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GpuTypes
}

// GetGpuTypesOk returns a tuple with the GpuTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHypervisorResourcePoolRequestModel) GetGpuTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.GpuTypes) {
		return nil, false
	}
	return o.GpuTypes, true
}

// HasGpuTypes returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasGpuTypes() bool {
	if o != nil && IsNil(o.GpuTypes) {
		return true
	}

	return false
}

// SetGpuTypes gets a reference to the given []string and assigns it to the GpuTypes field.
func (o *CreateHypervisorResourcePoolRequestModel) SetGpuTypes(v []string) {
	o.GpuTypes = v
}

// GetConnectionType returns the ConnectionType field value
func (o *CreateHypervisorResourcePoolRequestModel) GetConnectionType() HypervisorConnectionType {
	if o == nil {
		var ret HypervisorConnectionType
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *CreateHypervisorResourcePoolRequestModel) SetConnectionType(v HypervisorConnectionType) {
	o.ConnectionType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateHypervisorResourcePoolRequestModel) GetMetadata() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHypervisorResourcePoolRequestModel) GetMetadataOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []NameValueStringPairModel and assigns it to the Metadata field.
func (o *CreateHypervisorResourcePoolRequestModel) SetMetadata(v []NameValueStringPairModel) {
	o.Metadata = v
}

// GetVirtualPrivateCloud returns the VirtualPrivateCloud field value if set, zero value otherwise.
func (o *CreateHypervisorResourcePoolRequestModel) GetVirtualPrivateCloud() string {
	if o == nil || IsNil(o.VirtualPrivateCloud) {
		var ret string
		return ret
	}
	return *o.VirtualPrivateCloud
}

// GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolRequestModel) GetVirtualPrivateCloudOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualPrivateCloud) {
		return nil, false
	}
	return o.VirtualPrivateCloud, true
}

// HasVirtualPrivateCloud returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasVirtualPrivateCloud() bool {
	if o != nil && !IsNil(o.VirtualPrivateCloud) {
		return true
	}

	return false
}

// SetVirtualPrivateCloud gets a reference to the given string and assigns it to the VirtualPrivateCloud field.
func (o *CreateHypervisorResourcePoolRequestModel) SetVirtualPrivateCloud(v string) {
	o.VirtualPrivateCloud = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *CreateHypervisorResourcePoolRequestModel) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolRequestModel) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *CreateHypervisorResourcePoolRequestModel) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *CreateHypervisorResourcePoolRequestModel) GetNetworks() []string {
	if o == nil || IsNil(o.Networks) {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolRequestModel) GetNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *CreateHypervisorResourcePoolRequestModel) SetNetworks(v []string) {
	o.Networks = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CreateHypervisorResourcePoolRequestModel) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolRequestModel) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateHypervisorResourcePoolRequestModel) SetRegion(v string) {
	o.Region = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *CreateHypervisorResourcePoolRequestModel) GetVirtualNetwork() string {
	if o == nil || IsNil(o.VirtualNetwork) {
		var ret string
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolRequestModel) GetVirtualNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualNetwork) {
		return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasVirtualNetwork() bool {
	if o != nil && !IsNil(o.VirtualNetwork) {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given string and assigns it to the VirtualNetwork field.
func (o *CreateHypervisorResourcePoolRequestModel) SetVirtualNetwork(v string) {
	o.VirtualNetwork = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *CreateHypervisorResourcePoolRequestModel) GetSubnets() []string {
	if o == nil || IsNil(o.Subnets) {
		var ret []string
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolRequestModel) GetSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subnets) {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasSubnets() bool {
	if o != nil && !IsNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *CreateHypervisorResourcePoolRequestModel) SetSubnets(v []string) {
	o.Subnets = v
}

// GetRootPath returns the RootPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateHypervisorResourcePoolRequestModel) GetRootPath() string {
	if o == nil || IsNil(o.RootPath.Get()) {
		var ret string
		return ret
	}
	return *o.RootPath.Get()
}

// GetRootPathOk returns a tuple with the RootPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHypervisorResourcePoolRequestModel) GetRootPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootPath.Get(), o.RootPath.IsSet()
}

// HasRootPath returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasRootPath() bool {
	if o != nil && o.RootPath.IsSet() {
		return true
	}

	return false
}

// SetRootPath gets a reference to the given NullableString and assigns it to the RootPath field.
func (o *CreateHypervisorResourcePoolRequestModel) SetRootPath(v string) {
	o.RootPath.Set(&v)
}

// SetRootPathNil sets the value for RootPath to be an explicit nil
func (o *CreateHypervisorResourcePoolRequestModel) SetRootPathNil() {
	o.RootPath.Set(nil)
}

// UnsetRootPath ensures that no value is present for RootPath, not even an explicit nil
func (o *CreateHypervisorResourcePoolRequestModel) UnsetRootPath() {
	o.RootPath.Unset()
}

// GetUseLocalStorageCaching returns the UseLocalStorageCaching field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateHypervisorResourcePoolRequestModel) GetUseLocalStorageCaching() bool {
	if o == nil || IsNil(o.UseLocalStorageCaching.Get()) {
		var ret bool
		return ret
	}
	return *o.UseLocalStorageCaching.Get()
}

// GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHypervisorResourcePoolRequestModel) GetUseLocalStorageCachingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseLocalStorageCaching.Get(), o.UseLocalStorageCaching.IsSet()
}

// HasUseLocalStorageCaching returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasUseLocalStorageCaching() bool {
	if o != nil && o.UseLocalStorageCaching.IsSet() {
		return true
	}

	return false
}

// SetUseLocalStorageCaching gets a reference to the given NullableBool and assigns it to the UseLocalStorageCaching field.
func (o *CreateHypervisorResourcePoolRequestModel) SetUseLocalStorageCaching(v bool) {
	o.UseLocalStorageCaching.Set(&v)
}

// SetUseLocalStorageCachingNil sets the value for UseLocalStorageCaching to be an explicit nil
func (o *CreateHypervisorResourcePoolRequestModel) SetUseLocalStorageCachingNil() {
	o.UseLocalStorageCaching.Set(nil)
}

// UnsetUseLocalStorageCaching ensures that no value is present for UseLocalStorageCaching, not even an explicit nil
func (o *CreateHypervisorResourcePoolRequestModel) UnsetUseLocalStorageCaching() {
	o.UseLocalStorageCaching.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateHypervisorResourcePoolRequestModel) GetStorage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHypervisorResourcePoolRequestModel) GetStorageOk() ([]string, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasStorage() bool {
	if o != nil && IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []string and assigns it to the Storage field.
func (o *CreateHypervisorResourcePoolRequestModel) SetStorage(v []string) {
	o.Storage = v
}

// GetPersonalvDiskStorage returns the PersonalvDiskStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateHypervisorResourcePoolRequestModel) GetPersonalvDiskStorage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PersonalvDiskStorage
}

// GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHypervisorResourcePoolRequestModel) GetPersonalvDiskStorageOk() ([]string, bool) {
	if o == nil || IsNil(o.PersonalvDiskStorage) {
		return nil, false
	}
	return o.PersonalvDiskStorage, true
}

// HasPersonalvDiskStorage returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasPersonalvDiskStorage() bool {
	if o != nil && IsNil(o.PersonalvDiskStorage) {
		return true
	}

	return false
}

// SetPersonalvDiskStorage gets a reference to the given []string and assigns it to the PersonalvDiskStorage field.
func (o *CreateHypervisorResourcePoolRequestModel) SetPersonalvDiskStorage(v []string) {
	o.PersonalvDiskStorage = v
}

// GetTemporaryStorage returns the TemporaryStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateHypervisorResourcePoolRequestModel) GetTemporaryStorage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TemporaryStorage
}

// GetTemporaryStorageOk returns a tuple with the TemporaryStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHypervisorResourcePoolRequestModel) GetTemporaryStorageOk() ([]string, bool) {
	if o == nil || IsNil(o.TemporaryStorage) {
		return nil, false
	}
	return o.TemporaryStorage, true
}

// HasTemporaryStorage returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasTemporaryStorage() bool {
	if o != nil && IsNil(o.TemporaryStorage) {
		return true
	}

	return false
}

// SetTemporaryStorage gets a reference to the given []string and assigns it to the TemporaryStorage field.
func (o *CreateHypervisorResourcePoolRequestModel) SetTemporaryStorage(v []string) {
	o.TemporaryStorage = v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateHypervisorResourcePoolRequestModel) GetCustomProperties() string {
	if o == nil || IsNil(o.CustomProperties.Get()) {
		var ret string
		return ret
	}
	return *o.CustomProperties.Get()
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHypervisorResourcePoolRequestModel) GetCustomPropertiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomProperties.Get(), o.CustomProperties.IsSet()
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasCustomProperties() bool {
	if o != nil && o.CustomProperties.IsSet() {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given NullableString and assigns it to the CustomProperties field.
func (o *CreateHypervisorResourcePoolRequestModel) SetCustomProperties(v string) {
	o.CustomProperties.Set(&v)
}

// SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil
func (o *CreateHypervisorResourcePoolRequestModel) SetCustomPropertiesNil() {
	o.CustomProperties.Set(nil)
}

// UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
func (o *CreateHypervisorResourcePoolRequestModel) UnsetCustomProperties() {
	o.CustomProperties.Unset()
}

// GetStorageBalanceType returns the StorageBalanceType field value if set, zero value otherwise.
func (o *CreateHypervisorResourcePoolRequestModel) GetStorageBalanceType() StorageBalanceType {
	if o == nil || IsNil(o.StorageBalanceType) {
		var ret StorageBalanceType
		return ret
	}
	return *o.StorageBalanceType
}

// GetStorageBalanceTypeOk returns a tuple with the StorageBalanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHypervisorResourcePoolRequestModel) GetStorageBalanceTypeOk() (*StorageBalanceType, bool) {
	if o == nil || IsNil(o.StorageBalanceType) {
		return nil, false
	}
	return o.StorageBalanceType, true
}

// HasStorageBalanceType returns a boolean if a field has been set.
func (o *CreateHypervisorResourcePoolRequestModel) HasStorageBalanceType() bool {
	if o != nil && !IsNil(o.StorageBalanceType) {
		return true
	}

	return false
}

// SetStorageBalanceType gets a reference to the given StorageBalanceType and assigns it to the StorageBalanceType field.
func (o *CreateHypervisorResourcePoolRequestModel) SetStorageBalanceType(v StorageBalanceType) {
	o.StorageBalanceType = &v
}

func (o CreateHypervisorResourcePoolRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateHypervisorResourcePoolRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	if o.VmTagging.IsSet() {
		toSerialize["VmTagging"] = o.VmTagging.Get()
	}
	if o.GpuTypes != nil {
		toSerialize["GpuTypes"] = o.GpuTypes
	}
	toSerialize["ConnectionType"] = o.ConnectionType
	if o.Metadata != nil {
		toSerialize["Metadata"] = o.Metadata
	}
	if !IsNil(o.VirtualPrivateCloud) {
		toSerialize["VirtualPrivateCloud"] = o.VirtualPrivateCloud
	}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["AvailabilityZone"] = o.AvailabilityZone
	}
	if !IsNil(o.Networks) {
		toSerialize["Networks"] = o.Networks
	}
	if !IsNil(o.Region) {
		toSerialize["Region"] = o.Region
	}
	if !IsNil(o.VirtualNetwork) {
		toSerialize["VirtualNetwork"] = o.VirtualNetwork
	}
	if !IsNil(o.Subnets) {
		toSerialize["Subnets"] = o.Subnets
	}
	if o.RootPath.IsSet() {
		toSerialize["RootPath"] = o.RootPath.Get()
	}
	if o.UseLocalStorageCaching.IsSet() {
		toSerialize["UseLocalStorageCaching"] = o.UseLocalStorageCaching.Get()
	}
	if o.Storage != nil {
		toSerialize["Storage"] = o.Storage
	}
	if o.PersonalvDiskStorage != nil {
		toSerialize["PersonalvDiskStorage"] = o.PersonalvDiskStorage
	}
	if o.TemporaryStorage != nil {
		toSerialize["TemporaryStorage"] = o.TemporaryStorage
	}
	if o.CustomProperties.IsSet() {
		toSerialize["CustomProperties"] = o.CustomProperties.Get()
	}
	if !IsNil(o.StorageBalanceType) {
		toSerialize["StorageBalanceType"] = o.StorageBalanceType
	}
	return toSerialize, nil
}

type NullableCreateHypervisorResourcePoolRequestModel struct {
	value *CreateHypervisorResourcePoolRequestModel
	isSet bool
}

func (v NullableCreateHypervisorResourcePoolRequestModel) Get() *CreateHypervisorResourcePoolRequestModel {
	return v.value
}

func (v *NullableCreateHypervisorResourcePoolRequestModel) Set(val *CreateHypervisorResourcePoolRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateHypervisorResourcePoolRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateHypervisorResourcePoolRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateHypervisorResourcePoolRequestModel(val *CreateHypervisorResourcePoolRequestModel) *NullableCreateHypervisorResourcePoolRequestModel {
	return &NullableCreateHypervisorResourcePoolRequestModel{value: val, isSet: true}
}

func (v NullableCreateHypervisorResourcePoolRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateHypervisorResourcePoolRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
