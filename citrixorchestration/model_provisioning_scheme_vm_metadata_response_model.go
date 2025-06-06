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

// checks if the ProvisioningSchemeVmMetadataResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningSchemeVmMetadataResponseModel{}

// ProvisioningSchemeVmMetadataResponseModel Provisioning Scheme VMMetadata response model
type ProvisioningSchemeVmMetadataResponseModel struct {
	// VmMetadata machine size
	MachineSize NullableString `json:"MachineSize,omitempty"`
	// VmMetadata OS disk cache
	OsDiskCaching NullableString `json:"OsDiskCaching,omitempty"`
	// VmMetadata tags
	Tags NullableString `json:"Tags,omitempty"`
	// VmMetadata boot diagnostics configuration
	BootDiagnostics NullableBool `json:"BootDiagnostics,omitempty"`
	// VmMetadata accelerated network configuration
	AcceleratedNetwork NullableBool `json:"AcceleratedNetwork,omitempty"`
	// VmMetadata hibernation configuration.
	SupportsHibernation NullableBool `json:"SupportsHibernation,omitempty"`
	// VmMetadata security type configuration
	SecurityType NullableString `json:"SecurityType,omitempty"`
	// VmMetadata disk security type configuration
	DiskSecurityType NullableString `json:"DiskSecurityType,omitempty"`
	// VmMetadata confidential VM disk encryption set id configuration
	ConfidentialVmDiskEncryptionSetId NullableString `json:"ConfidentialVmDiskEncryptionSetId,omitempty"`
	// VmMetadata VirtualMachine.Resources.Properties.SecurityProfile.UefiSettings.secureBootEnabled
	EnableSecureBoot NullableBool `json:"EnableSecureBoot,omitempty"`
	// VmMetadata VirtualMachine.Resources.Properties.SecurityProfile.UefiSettings.vTpmEnabled
	EnableVTPM NullableBool `json:"EnableVTPM,omitempty"`
	// VmMetadata VirtualMachine.Resources.Properties.SecurityProfile.EncryptionAtHost
	EncryptionAtHost NullableBool `json:"EncryptionAtHost,omitempty"`
	// VmMetadata Labels
	Labels NullableString `json:"Labels,omitempty"`
	// VmMetadata zone name
	ZoneName NullableString `json:"ZoneName,omitempty"`
	// VmMetadata storage type
	StorageType NullableString `json:"StorageType,omitempty"`
	// VmMetadata encryption key
	EncryptionKeyId NullableString `json:"EncryptionKeyId,omitempty"`
}

// NewProvisioningSchemeVmMetadataResponseModel instantiates a new ProvisioningSchemeVmMetadataResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningSchemeVmMetadataResponseModel() *ProvisioningSchemeVmMetadataResponseModel {
	this := ProvisioningSchemeVmMetadataResponseModel{}
	return &this
}

// NewProvisioningSchemeVmMetadataResponseModelWithDefaults instantiates a new ProvisioningSchemeVmMetadataResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningSchemeVmMetadataResponseModelWithDefaults() *ProvisioningSchemeVmMetadataResponseModel {
	this := ProvisioningSchemeVmMetadataResponseModel{}
	return &this
}

// GetMachineSize returns the MachineSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetMachineSize() string {
	if o == nil || IsNil(o.MachineSize.Get()) {
		var ret string
		return ret
	}
	return *o.MachineSize.Get()
}

// GetMachineSizeOk returns a tuple with the MachineSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetMachineSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineSize.Get(), o.MachineSize.IsSet()
}

// HasMachineSize returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasMachineSize() bool {
	if o != nil && o.MachineSize.IsSet() {
		return true
	}

	return false
}

// SetMachineSize gets a reference to the given NullableString and assigns it to the MachineSize field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetMachineSize(v string) {
	o.MachineSize.Set(&v)
}

// SetMachineSizeNil sets the value for MachineSize to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetMachineSizeNil() {
	o.MachineSize.Set(nil)
}

// UnsetMachineSize ensures that no value is present for MachineSize, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetMachineSize() {
	o.MachineSize.Unset()
}

// GetOsDiskCaching returns the OsDiskCaching field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetOsDiskCaching() string {
	if o == nil || IsNil(o.OsDiskCaching.Get()) {
		var ret string
		return ret
	}
	return *o.OsDiskCaching.Get()
}

// GetOsDiskCachingOk returns a tuple with the OsDiskCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetOsDiskCachingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsDiskCaching.Get(), o.OsDiskCaching.IsSet()
}

// HasOsDiskCaching returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasOsDiskCaching() bool {
	if o != nil && o.OsDiskCaching.IsSet() {
		return true
	}

	return false
}

// SetOsDiskCaching gets a reference to the given NullableString and assigns it to the OsDiskCaching field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetOsDiskCaching(v string) {
	o.OsDiskCaching.Set(&v)
}

// SetOsDiskCachingNil sets the value for OsDiskCaching to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetOsDiskCachingNil() {
	o.OsDiskCaching.Set(nil)
}

// UnsetOsDiskCaching ensures that no value is present for OsDiskCaching, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetOsDiskCaching() {
	o.OsDiskCaching.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetTags() {
	o.Tags.Unset()
}

// GetBootDiagnostics returns the BootDiagnostics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetBootDiagnostics() bool {
	if o == nil || IsNil(o.BootDiagnostics.Get()) {
		var ret bool
		return ret
	}
	return *o.BootDiagnostics.Get()
}

// GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetBootDiagnosticsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BootDiagnostics.Get(), o.BootDiagnostics.IsSet()
}

// HasBootDiagnostics returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasBootDiagnostics() bool {
	if o != nil && o.BootDiagnostics.IsSet() {
		return true
	}

	return false
}

// SetBootDiagnostics gets a reference to the given NullableBool and assigns it to the BootDiagnostics field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetBootDiagnostics(v bool) {
	o.BootDiagnostics.Set(&v)
}

// SetBootDiagnosticsNil sets the value for BootDiagnostics to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetBootDiagnosticsNil() {
	o.BootDiagnostics.Set(nil)
}

// UnsetBootDiagnostics ensures that no value is present for BootDiagnostics, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetBootDiagnostics() {
	o.BootDiagnostics.Unset()
}

// GetAcceleratedNetwork returns the AcceleratedNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetAcceleratedNetwork() bool {
	if o == nil || IsNil(o.AcceleratedNetwork.Get()) {
		var ret bool
		return ret
	}
	return *o.AcceleratedNetwork.Get()
}

// GetAcceleratedNetworkOk returns a tuple with the AcceleratedNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetAcceleratedNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceleratedNetwork.Get(), o.AcceleratedNetwork.IsSet()
}

// HasAcceleratedNetwork returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasAcceleratedNetwork() bool {
	if o != nil && o.AcceleratedNetwork.IsSet() {
		return true
	}

	return false
}

// SetAcceleratedNetwork gets a reference to the given NullableBool and assigns it to the AcceleratedNetwork field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetAcceleratedNetwork(v bool) {
	o.AcceleratedNetwork.Set(&v)
}

// SetAcceleratedNetworkNil sets the value for AcceleratedNetwork to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetAcceleratedNetworkNil() {
	o.AcceleratedNetwork.Set(nil)
}

// UnsetAcceleratedNetwork ensures that no value is present for AcceleratedNetwork, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetAcceleratedNetwork() {
	o.AcceleratedNetwork.Unset()
}

// GetSupportsHibernation returns the SupportsHibernation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetSupportsHibernation() bool {
	if o == nil || IsNil(o.SupportsHibernation.Get()) {
		var ret bool
		return ret
	}
	return *o.SupportsHibernation.Get()
}

// GetSupportsHibernationOk returns a tuple with the SupportsHibernation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetSupportsHibernationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportsHibernation.Get(), o.SupportsHibernation.IsSet()
}

// HasSupportsHibernation returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasSupportsHibernation() bool {
	if o != nil && o.SupportsHibernation.IsSet() {
		return true
	}

	return false
}

// SetSupportsHibernation gets a reference to the given NullableBool and assigns it to the SupportsHibernation field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetSupportsHibernation(v bool) {
	o.SupportsHibernation.Set(&v)
}

// SetSupportsHibernationNil sets the value for SupportsHibernation to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetSupportsHibernationNil() {
	o.SupportsHibernation.Set(nil)
}

// UnsetSupportsHibernation ensures that no value is present for SupportsHibernation, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetSupportsHibernation() {
	o.SupportsHibernation.Unset()
}

// GetSecurityType returns the SecurityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetSecurityType() string {
	if o == nil || IsNil(o.SecurityType.Get()) {
		var ret string
		return ret
	}
	return *o.SecurityType.Get()
}

// GetSecurityTypeOk returns a tuple with the SecurityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetSecurityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecurityType.Get(), o.SecurityType.IsSet()
}

// HasSecurityType returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasSecurityType() bool {
	if o != nil && o.SecurityType.IsSet() {
		return true
	}

	return false
}

// SetSecurityType gets a reference to the given NullableString and assigns it to the SecurityType field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetSecurityType(v string) {
	o.SecurityType.Set(&v)
}

// SetSecurityTypeNil sets the value for SecurityType to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetSecurityTypeNil() {
	o.SecurityType.Set(nil)
}

// UnsetSecurityType ensures that no value is present for SecurityType, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetSecurityType() {
	o.SecurityType.Unset()
}

// GetDiskSecurityType returns the DiskSecurityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetDiskSecurityType() string {
	if o == nil || IsNil(o.DiskSecurityType.Get()) {
		var ret string
		return ret
	}
	return *o.DiskSecurityType.Get()
}

// GetDiskSecurityTypeOk returns a tuple with the DiskSecurityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetDiskSecurityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskSecurityType.Get(), o.DiskSecurityType.IsSet()
}

// HasDiskSecurityType returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasDiskSecurityType() bool {
	if o != nil && o.DiskSecurityType.IsSet() {
		return true
	}

	return false
}

// SetDiskSecurityType gets a reference to the given NullableString and assigns it to the DiskSecurityType field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetDiskSecurityType(v string) {
	o.DiskSecurityType.Set(&v)
}

// SetDiskSecurityTypeNil sets the value for DiskSecurityType to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetDiskSecurityTypeNil() {
	o.DiskSecurityType.Set(nil)
}

// UnsetDiskSecurityType ensures that no value is present for DiskSecurityType, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetDiskSecurityType() {
	o.DiskSecurityType.Unset()
}

// GetConfidentialVmDiskEncryptionSetId returns the ConfidentialVmDiskEncryptionSetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetConfidentialVmDiskEncryptionSetId() string {
	if o == nil || IsNil(o.ConfidentialVmDiskEncryptionSetId.Get()) {
		var ret string
		return ret
	}
	return *o.ConfidentialVmDiskEncryptionSetId.Get()
}

// GetConfidentialVmDiskEncryptionSetIdOk returns a tuple with the ConfidentialVmDiskEncryptionSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetConfidentialVmDiskEncryptionSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfidentialVmDiskEncryptionSetId.Get(), o.ConfidentialVmDiskEncryptionSetId.IsSet()
}

// HasConfidentialVmDiskEncryptionSetId returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasConfidentialVmDiskEncryptionSetId() bool {
	if o != nil && o.ConfidentialVmDiskEncryptionSetId.IsSet() {
		return true
	}

	return false
}

// SetConfidentialVmDiskEncryptionSetId gets a reference to the given NullableString and assigns it to the ConfidentialVmDiskEncryptionSetId field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetConfidentialVmDiskEncryptionSetId(v string) {
	o.ConfidentialVmDiskEncryptionSetId.Set(&v)
}

// SetConfidentialVmDiskEncryptionSetIdNil sets the value for ConfidentialVmDiskEncryptionSetId to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetConfidentialVmDiskEncryptionSetIdNil() {
	o.ConfidentialVmDiskEncryptionSetId.Set(nil)
}

// UnsetConfidentialVmDiskEncryptionSetId ensures that no value is present for ConfidentialVmDiskEncryptionSetId, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetConfidentialVmDiskEncryptionSetId() {
	o.ConfidentialVmDiskEncryptionSetId.Unset()
}

// GetEnableSecureBoot returns the EnableSecureBoot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetEnableSecureBoot() bool {
	if o == nil || IsNil(o.EnableSecureBoot.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableSecureBoot.Get()
}

// GetEnableSecureBootOk returns a tuple with the EnableSecureBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetEnableSecureBootOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableSecureBoot.Get(), o.EnableSecureBoot.IsSet()
}

// HasEnableSecureBoot returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasEnableSecureBoot() bool {
	if o != nil && o.EnableSecureBoot.IsSet() {
		return true
	}

	return false
}

// SetEnableSecureBoot gets a reference to the given NullableBool and assigns it to the EnableSecureBoot field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetEnableSecureBoot(v bool) {
	o.EnableSecureBoot.Set(&v)
}

// SetEnableSecureBootNil sets the value for EnableSecureBoot to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetEnableSecureBootNil() {
	o.EnableSecureBoot.Set(nil)
}

// UnsetEnableSecureBoot ensures that no value is present for EnableSecureBoot, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetEnableSecureBoot() {
	o.EnableSecureBoot.Unset()
}

// GetEnableVTPM returns the EnableVTPM field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetEnableVTPM() bool {
	if o == nil || IsNil(o.EnableVTPM.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableVTPM.Get()
}

// GetEnableVTPMOk returns a tuple with the EnableVTPM field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetEnableVTPMOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableVTPM.Get(), o.EnableVTPM.IsSet()
}

// HasEnableVTPM returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasEnableVTPM() bool {
	if o != nil && o.EnableVTPM.IsSet() {
		return true
	}

	return false
}

// SetEnableVTPM gets a reference to the given NullableBool and assigns it to the EnableVTPM field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetEnableVTPM(v bool) {
	o.EnableVTPM.Set(&v)
}

// SetEnableVTPMNil sets the value for EnableVTPM to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetEnableVTPMNil() {
	o.EnableVTPM.Set(nil)
}

// UnsetEnableVTPM ensures that no value is present for EnableVTPM, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetEnableVTPM() {
	o.EnableVTPM.Unset()
}

// GetEncryptionAtHost returns the EncryptionAtHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetEncryptionAtHost() bool {
	if o == nil || IsNil(o.EncryptionAtHost.Get()) {
		var ret bool
		return ret
	}
	return *o.EncryptionAtHost.Get()
}

// GetEncryptionAtHostOk returns a tuple with the EncryptionAtHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetEncryptionAtHostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptionAtHost.Get(), o.EncryptionAtHost.IsSet()
}

// HasEncryptionAtHost returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasEncryptionAtHost() bool {
	if o != nil && o.EncryptionAtHost.IsSet() {
		return true
	}

	return false
}

// SetEncryptionAtHost gets a reference to the given NullableBool and assigns it to the EncryptionAtHost field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetEncryptionAtHost(v bool) {
	o.EncryptionAtHost.Set(&v)
}

// SetEncryptionAtHostNil sets the value for EncryptionAtHost to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetEncryptionAtHostNil() {
	o.EncryptionAtHost.Set(nil)
}

// UnsetEncryptionAtHost ensures that no value is present for EncryptionAtHost, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetEncryptionAtHost() {
	o.EncryptionAtHost.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetLabels() string {
	if o == nil || IsNil(o.Labels.Get()) {
		var ret string
		return ret
	}
	return *o.Labels.Get()
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetLabelsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels.Get(), o.Labels.IsSet()
}

// HasLabels returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasLabels() bool {
	if o != nil && o.Labels.IsSet() {
		return true
	}

	return false
}

// SetLabels gets a reference to the given NullableString and assigns it to the Labels field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetLabels(v string) {
	o.Labels.Set(&v)
}

// SetLabelsNil sets the value for Labels to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetLabelsNil() {
	o.Labels.Set(nil)
}

// UnsetLabels ensures that no value is present for Labels, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetLabels() {
	o.Labels.Unset()
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetZoneName() string {
	if o == nil || IsNil(o.ZoneName.Get()) {
		var ret string
		return ret
	}
	return *o.ZoneName.Get()
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetZoneNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZoneName.Get(), o.ZoneName.IsSet()
}

// HasZoneName returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasZoneName() bool {
	if o != nil && o.ZoneName.IsSet() {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given NullableString and assigns it to the ZoneName field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetZoneName(v string) {
	o.ZoneName.Set(&v)
}

// SetZoneNameNil sets the value for ZoneName to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetZoneNameNil() {
	o.ZoneName.Set(nil)
}

// UnsetZoneName ensures that no value is present for ZoneName, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetZoneName() {
	o.ZoneName.Unset()
}

// GetStorageType returns the StorageType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetStorageType() string {
	if o == nil || IsNil(o.StorageType.Get()) {
		var ret string
		return ret
	}
	return *o.StorageType.Get()
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetStorageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageType.Get(), o.StorageType.IsSet()
}

// HasStorageType returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasStorageType() bool {
	if o != nil && o.StorageType.IsSet() {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given NullableString and assigns it to the StorageType field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetStorageType(v string) {
	o.StorageType.Set(&v)
}

// SetStorageTypeNil sets the value for StorageType to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetStorageTypeNil() {
	o.StorageType.Set(nil)
}

// UnsetStorageType ensures that no value is present for StorageType, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetStorageType() {
	o.StorageType.Unset()
}

// GetEncryptionKeyId returns the EncryptionKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeVmMetadataResponseModel) GetEncryptionKeyId() string {
	if o == nil || IsNil(o.EncryptionKeyId.Get()) {
		var ret string
		return ret
	}
	return *o.EncryptionKeyId.Get()
}

// GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeVmMetadataResponseModel) GetEncryptionKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptionKeyId.Get(), o.EncryptionKeyId.IsSet()
}

// HasEncryptionKeyId returns a boolean if a field has been set.
func (o *ProvisioningSchemeVmMetadataResponseModel) HasEncryptionKeyId() bool {
	if o != nil && o.EncryptionKeyId.IsSet() {
		return true
	}

	return false
}

// SetEncryptionKeyId gets a reference to the given NullableString and assigns it to the EncryptionKeyId field.
func (o *ProvisioningSchemeVmMetadataResponseModel) SetEncryptionKeyId(v string) {
	o.EncryptionKeyId.Set(&v)
}

// SetEncryptionKeyIdNil sets the value for EncryptionKeyId to be an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) SetEncryptionKeyIdNil() {
	o.EncryptionKeyId.Set(nil)
}

// UnsetEncryptionKeyId ensures that no value is present for EncryptionKeyId, not even an explicit nil
func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetEncryptionKeyId() {
	o.EncryptionKeyId.Unset()
}

func (o ProvisioningSchemeVmMetadataResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningSchemeVmMetadataResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MachineSize.IsSet() {
		toSerialize["MachineSize"] = o.MachineSize.Get()
	}
	if o.OsDiskCaching.IsSet() {
		toSerialize["OsDiskCaching"] = o.OsDiskCaching.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["Tags"] = o.Tags.Get()
	}
	if o.BootDiagnostics.IsSet() {
		toSerialize["BootDiagnostics"] = o.BootDiagnostics.Get()
	}
	if o.AcceleratedNetwork.IsSet() {
		toSerialize["AcceleratedNetwork"] = o.AcceleratedNetwork.Get()
	}
	if o.SupportsHibernation.IsSet() {
		toSerialize["SupportsHibernation"] = o.SupportsHibernation.Get()
	}
	if o.SecurityType.IsSet() {
		toSerialize["SecurityType"] = o.SecurityType.Get()
	}
	if o.DiskSecurityType.IsSet() {
		toSerialize["DiskSecurityType"] = o.DiskSecurityType.Get()
	}
	if o.ConfidentialVmDiskEncryptionSetId.IsSet() {
		toSerialize["ConfidentialVmDiskEncryptionSetId"] = o.ConfidentialVmDiskEncryptionSetId.Get()
	}
	if o.EnableSecureBoot.IsSet() {
		toSerialize["EnableSecureBoot"] = o.EnableSecureBoot.Get()
	}
	if o.EnableVTPM.IsSet() {
		toSerialize["EnableVTPM"] = o.EnableVTPM.Get()
	}
	if o.EncryptionAtHost.IsSet() {
		toSerialize["EncryptionAtHost"] = o.EncryptionAtHost.Get()
	}
	if o.Labels.IsSet() {
		toSerialize["Labels"] = o.Labels.Get()
	}
	if o.ZoneName.IsSet() {
		toSerialize["ZoneName"] = o.ZoneName.Get()
	}
	if o.StorageType.IsSet() {
		toSerialize["StorageType"] = o.StorageType.Get()
	}
	if o.EncryptionKeyId.IsSet() {
		toSerialize["EncryptionKeyId"] = o.EncryptionKeyId.Get()
	}
	return toSerialize, nil
}

type NullableProvisioningSchemeVmMetadataResponseModel struct {
	value *ProvisioningSchemeVmMetadataResponseModel
	isSet bool
}

func (v NullableProvisioningSchemeVmMetadataResponseModel) Get() *ProvisioningSchemeVmMetadataResponseModel {
	return v.value
}

func (v *NullableProvisioningSchemeVmMetadataResponseModel) Set(val *ProvisioningSchemeVmMetadataResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningSchemeVmMetadataResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningSchemeVmMetadataResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningSchemeVmMetadataResponseModel(val *ProvisioningSchemeVmMetadataResponseModel) *NullableProvisioningSchemeVmMetadataResponseModel {
	return &NullableProvisioningSchemeVmMetadataResponseModel{value: val, isSet: true}
}

func (v NullableProvisioningSchemeVmMetadataResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningSchemeVmMetadataResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
