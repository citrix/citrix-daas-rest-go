# ProvisioningSchemeVmMetadataResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineSize** | Pointer to **NullableString** | VmMetadata machine size | [optional] 
**OsDiskCaching** | Pointer to **NullableString** | VmMetadata OS disk cache | [optional] 
**Tags** | Pointer to **NullableString** | VmMetadata tags | [optional] 
**BootDiagnostics** | Pointer to **NullableBool** | VmMetadata boot diagnostics configuration | [optional] 
**AcceleratedNetwork** | Pointer to **NullableBool** | VmMetadata accelerated network configuration | [optional] 
**SupportsHibernation** | Pointer to **NullableBool** | VmMetadata hibernation configuration. | [optional] 
**SecurityType** | Pointer to **NullableString** | VmMetadata security type configuration | [optional] 
**DiskSecurityType** | Pointer to **NullableString** | VmMetadata disk security type configuration | [optional] 
**ConfidentialVmDiskEncryptionSetId** | Pointer to **NullableString** | VmMetadata confidential VM disk encryption set id configuration | [optional] 
**Labels** | Pointer to **NullableString** | VmMetadata Labels | [optional] 
**ZoneName** | Pointer to **NullableString** | VmMetadata zone name | [optional] 

## Methods

### NewProvisioningSchemeVmMetadataResponseModel

`func NewProvisioningSchemeVmMetadataResponseModel() *ProvisioningSchemeVmMetadataResponseModel`

NewProvisioningSchemeVmMetadataResponseModel instantiates a new ProvisioningSchemeVmMetadataResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeVmMetadataResponseModelWithDefaults

`func NewProvisioningSchemeVmMetadataResponseModelWithDefaults() *ProvisioningSchemeVmMetadataResponseModel`

NewProvisioningSchemeVmMetadataResponseModelWithDefaults instantiates a new ProvisioningSchemeVmMetadataResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineSize

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetMachineSize() string`

GetMachineSize returns the MachineSize field if non-nil, zero value otherwise.

### GetMachineSizeOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetMachineSizeOk() (*string, bool)`

GetMachineSizeOk returns a tuple with the MachineSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineSize

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetMachineSize(v string)`

SetMachineSize sets MachineSize field to given value.

### HasMachineSize

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasMachineSize() bool`

HasMachineSize returns a boolean if a field has been set.

### SetMachineSizeNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetMachineSizeNil(b bool)`

 SetMachineSizeNil sets the value for MachineSize to be an explicit nil

### UnsetMachineSize
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetMachineSize()`

UnsetMachineSize ensures that no value is present for MachineSize, not even an explicit nil
### GetOsDiskCaching

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetOsDiskCaching() string`

GetOsDiskCaching returns the OsDiskCaching field if non-nil, zero value otherwise.

### GetOsDiskCachingOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetOsDiskCachingOk() (*string, bool)`

GetOsDiskCachingOk returns a tuple with the OsDiskCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDiskCaching

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetOsDiskCaching(v string)`

SetOsDiskCaching sets OsDiskCaching field to given value.

### HasOsDiskCaching

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasOsDiskCaching() bool`

HasOsDiskCaching returns a boolean if a field has been set.

### SetOsDiskCachingNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetOsDiskCachingNil(b bool)`

 SetOsDiskCachingNil sets the value for OsDiskCaching to be an explicit nil

### UnsetOsDiskCaching
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetOsDiskCaching()`

UnsetOsDiskCaching ensures that no value is present for OsDiskCaching, not even an explicit nil
### GetTags

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetBootDiagnostics

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetBootDiagnostics() bool`

GetBootDiagnostics returns the BootDiagnostics field if non-nil, zero value otherwise.

### GetBootDiagnosticsOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetBootDiagnosticsOk() (*bool, bool)`

GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiagnostics

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetBootDiagnostics(v bool)`

SetBootDiagnostics sets BootDiagnostics field to given value.

### HasBootDiagnostics

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasBootDiagnostics() bool`

HasBootDiagnostics returns a boolean if a field has been set.

### SetBootDiagnosticsNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetBootDiagnosticsNil(b bool)`

 SetBootDiagnosticsNil sets the value for BootDiagnostics to be an explicit nil

### UnsetBootDiagnostics
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetBootDiagnostics()`

UnsetBootDiagnostics ensures that no value is present for BootDiagnostics, not even an explicit nil
### GetAcceleratedNetwork

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetAcceleratedNetwork() bool`

GetAcceleratedNetwork returns the AcceleratedNetwork field if non-nil, zero value otherwise.

### GetAcceleratedNetworkOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetAcceleratedNetworkOk() (*bool, bool)`

GetAcceleratedNetworkOk returns a tuple with the AcceleratedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratedNetwork

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetAcceleratedNetwork(v bool)`

SetAcceleratedNetwork sets AcceleratedNetwork field to given value.

### HasAcceleratedNetwork

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasAcceleratedNetwork() bool`

HasAcceleratedNetwork returns a boolean if a field has been set.

### SetAcceleratedNetworkNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetAcceleratedNetworkNil(b bool)`

 SetAcceleratedNetworkNil sets the value for AcceleratedNetwork to be an explicit nil

### UnsetAcceleratedNetwork
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetAcceleratedNetwork()`

UnsetAcceleratedNetwork ensures that no value is present for AcceleratedNetwork, not even an explicit nil
### GetSupportsHibernation

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetSupportsHibernation() bool`

GetSupportsHibernation returns the SupportsHibernation field if non-nil, zero value otherwise.

### GetSupportsHibernationOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetSupportsHibernationOk() (*bool, bool)`

GetSupportsHibernationOk returns a tuple with the SupportsHibernation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsHibernation

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetSupportsHibernation(v bool)`

SetSupportsHibernation sets SupportsHibernation field to given value.

### HasSupportsHibernation

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasSupportsHibernation() bool`

HasSupportsHibernation returns a boolean if a field has been set.

### SetSupportsHibernationNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetSupportsHibernationNil(b bool)`

 SetSupportsHibernationNil sets the value for SupportsHibernation to be an explicit nil

### UnsetSupportsHibernation
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetSupportsHibernation()`

UnsetSupportsHibernation ensures that no value is present for SupportsHibernation, not even an explicit nil
### GetSecurityType

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetSecurityType() string`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetSecurityTypeOk() (*string, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetSecurityType(v string)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### SetSecurityTypeNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetSecurityTypeNil(b bool)`

 SetSecurityTypeNil sets the value for SecurityType to be an explicit nil

### UnsetSecurityType
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetSecurityType()`

UnsetSecurityType ensures that no value is present for SecurityType, not even an explicit nil
### GetDiskSecurityType

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetDiskSecurityType() string`

GetDiskSecurityType returns the DiskSecurityType field if non-nil, zero value otherwise.

### GetDiskSecurityTypeOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetDiskSecurityTypeOk() (*string, bool)`

GetDiskSecurityTypeOk returns a tuple with the DiskSecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSecurityType

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetDiskSecurityType(v string)`

SetDiskSecurityType sets DiskSecurityType field to given value.

### HasDiskSecurityType

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasDiskSecurityType() bool`

HasDiskSecurityType returns a boolean if a field has been set.

### SetDiskSecurityTypeNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetDiskSecurityTypeNil(b bool)`

 SetDiskSecurityTypeNil sets the value for DiskSecurityType to be an explicit nil

### UnsetDiskSecurityType
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetDiskSecurityType()`

UnsetDiskSecurityType ensures that no value is present for DiskSecurityType, not even an explicit nil
### GetConfidentialVmDiskEncryptionSetId

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetConfidentialVmDiskEncryptionSetId() string`

GetConfidentialVmDiskEncryptionSetId returns the ConfidentialVmDiskEncryptionSetId field if non-nil, zero value otherwise.

### GetConfidentialVmDiskEncryptionSetIdOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetConfidentialVmDiskEncryptionSetIdOk() (*string, bool)`

GetConfidentialVmDiskEncryptionSetIdOk returns a tuple with the ConfidentialVmDiskEncryptionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialVmDiskEncryptionSetId

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetConfidentialVmDiskEncryptionSetId(v string)`

SetConfidentialVmDiskEncryptionSetId sets ConfidentialVmDiskEncryptionSetId field to given value.

### HasConfidentialVmDiskEncryptionSetId

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasConfidentialVmDiskEncryptionSetId() bool`

HasConfidentialVmDiskEncryptionSetId returns a boolean if a field has been set.

### SetConfidentialVmDiskEncryptionSetIdNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetConfidentialVmDiskEncryptionSetIdNil(b bool)`

 SetConfidentialVmDiskEncryptionSetIdNil sets the value for ConfidentialVmDiskEncryptionSetId to be an explicit nil

### UnsetConfidentialVmDiskEncryptionSetId
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetConfidentialVmDiskEncryptionSetId()`

UnsetConfidentialVmDiskEncryptionSetId ensures that no value is present for ConfidentialVmDiskEncryptionSetId, not even an explicit nil
### GetLabels

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetLabels() string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetLabelsOk() (*string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetLabels(v string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetZoneName

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ProvisioningSchemeVmMetadataResponseModel) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ProvisioningSchemeVmMetadataResponseModel) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### SetZoneNameNil

`func (o *ProvisioningSchemeVmMetadataResponseModel) SetZoneNameNil(b bool)`

 SetZoneNameNil sets the value for ZoneName to be an explicit nil

### UnsetZoneName
`func (o *ProvisioningSchemeVmMetadataResponseModel) UnsetZoneName()`

UnsetZoneName ensures that no value is present for ZoneName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


