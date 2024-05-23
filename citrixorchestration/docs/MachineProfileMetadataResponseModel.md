# MachineProfileMetadataResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineSize** | Pointer to **NullableString** | VmMetadata machine size | [optional] 
**OsDiskCaching** | Pointer to **NullableString** | VmMetadata OS disk cache | [optional] 
**Tags** | Pointer to **NullableString** | VmMetadata tags | [optional] 
**BootDiagnostics** | Pointer to **NullableBool** | VmMetadata boot diagnostics configuration | [optional] 
**AcceleratedNetwork** | Pointer to **NullableBool** | VmMetadata accelerated network configuration | [optional] 
**SupportsHibernation** | Pointer to **NullableBool** | VmMetadata hibernation configuration. | [optional] 
**ZoneName** | Pointer to **NullableString** | VmMetadata zone name | [optional] 
**Labels** | Pointer to **NullableString** | VmMetadata labels | [optional] 

## Methods

### NewMachineProfileMetadataResponseModel

`func NewMachineProfileMetadataResponseModel() *MachineProfileMetadataResponseModel`

NewMachineProfileMetadataResponseModel instantiates a new MachineProfileMetadataResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineProfileMetadataResponseModelWithDefaults

`func NewMachineProfileMetadataResponseModelWithDefaults() *MachineProfileMetadataResponseModel`

NewMachineProfileMetadataResponseModelWithDefaults instantiates a new MachineProfileMetadataResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineSize

`func (o *MachineProfileMetadataResponseModel) GetMachineSize() string`

GetMachineSize returns the MachineSize field if non-nil, zero value otherwise.

### GetMachineSizeOk

`func (o *MachineProfileMetadataResponseModel) GetMachineSizeOk() (*string, bool)`

GetMachineSizeOk returns a tuple with the MachineSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineSize

`func (o *MachineProfileMetadataResponseModel) SetMachineSize(v string)`

SetMachineSize sets MachineSize field to given value.

### HasMachineSize

`func (o *MachineProfileMetadataResponseModel) HasMachineSize() bool`

HasMachineSize returns a boolean if a field has been set.

### SetMachineSizeNil

`func (o *MachineProfileMetadataResponseModel) SetMachineSizeNil(b bool)`

 SetMachineSizeNil sets the value for MachineSize to be an explicit nil

### UnsetMachineSize
`func (o *MachineProfileMetadataResponseModel) UnsetMachineSize()`

UnsetMachineSize ensures that no value is present for MachineSize, not even an explicit nil
### GetOsDiskCaching

`func (o *MachineProfileMetadataResponseModel) GetOsDiskCaching() string`

GetOsDiskCaching returns the OsDiskCaching field if non-nil, zero value otherwise.

### GetOsDiskCachingOk

`func (o *MachineProfileMetadataResponseModel) GetOsDiskCachingOk() (*string, bool)`

GetOsDiskCachingOk returns a tuple with the OsDiskCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDiskCaching

`func (o *MachineProfileMetadataResponseModel) SetOsDiskCaching(v string)`

SetOsDiskCaching sets OsDiskCaching field to given value.

### HasOsDiskCaching

`func (o *MachineProfileMetadataResponseModel) HasOsDiskCaching() bool`

HasOsDiskCaching returns a boolean if a field has been set.

### SetOsDiskCachingNil

`func (o *MachineProfileMetadataResponseModel) SetOsDiskCachingNil(b bool)`

 SetOsDiskCachingNil sets the value for OsDiskCaching to be an explicit nil

### UnsetOsDiskCaching
`func (o *MachineProfileMetadataResponseModel) UnsetOsDiskCaching()`

UnsetOsDiskCaching ensures that no value is present for OsDiskCaching, not even an explicit nil
### GetTags

`func (o *MachineProfileMetadataResponseModel) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MachineProfileMetadataResponseModel) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MachineProfileMetadataResponseModel) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MachineProfileMetadataResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MachineProfileMetadataResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MachineProfileMetadataResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetBootDiagnostics

`func (o *MachineProfileMetadataResponseModel) GetBootDiagnostics() bool`

GetBootDiagnostics returns the BootDiagnostics field if non-nil, zero value otherwise.

### GetBootDiagnosticsOk

`func (o *MachineProfileMetadataResponseModel) GetBootDiagnosticsOk() (*bool, bool)`

GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiagnostics

`func (o *MachineProfileMetadataResponseModel) SetBootDiagnostics(v bool)`

SetBootDiagnostics sets BootDiagnostics field to given value.

### HasBootDiagnostics

`func (o *MachineProfileMetadataResponseModel) HasBootDiagnostics() bool`

HasBootDiagnostics returns a boolean if a field has been set.

### SetBootDiagnosticsNil

`func (o *MachineProfileMetadataResponseModel) SetBootDiagnosticsNil(b bool)`

 SetBootDiagnosticsNil sets the value for BootDiagnostics to be an explicit nil

### UnsetBootDiagnostics
`func (o *MachineProfileMetadataResponseModel) UnsetBootDiagnostics()`

UnsetBootDiagnostics ensures that no value is present for BootDiagnostics, not even an explicit nil
### GetAcceleratedNetwork

`func (o *MachineProfileMetadataResponseModel) GetAcceleratedNetwork() bool`

GetAcceleratedNetwork returns the AcceleratedNetwork field if non-nil, zero value otherwise.

### GetAcceleratedNetworkOk

`func (o *MachineProfileMetadataResponseModel) GetAcceleratedNetworkOk() (*bool, bool)`

GetAcceleratedNetworkOk returns a tuple with the AcceleratedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratedNetwork

`func (o *MachineProfileMetadataResponseModel) SetAcceleratedNetwork(v bool)`

SetAcceleratedNetwork sets AcceleratedNetwork field to given value.

### HasAcceleratedNetwork

`func (o *MachineProfileMetadataResponseModel) HasAcceleratedNetwork() bool`

HasAcceleratedNetwork returns a boolean if a field has been set.

### SetAcceleratedNetworkNil

`func (o *MachineProfileMetadataResponseModel) SetAcceleratedNetworkNil(b bool)`

 SetAcceleratedNetworkNil sets the value for AcceleratedNetwork to be an explicit nil

### UnsetAcceleratedNetwork
`func (o *MachineProfileMetadataResponseModel) UnsetAcceleratedNetwork()`

UnsetAcceleratedNetwork ensures that no value is present for AcceleratedNetwork, not even an explicit nil
### GetSupportsHibernation

`func (o *MachineProfileMetadataResponseModel) GetSupportsHibernation() bool`

GetSupportsHibernation returns the SupportsHibernation field if non-nil, zero value otherwise.

### GetSupportsHibernationOk

`func (o *MachineProfileMetadataResponseModel) GetSupportsHibernationOk() (*bool, bool)`

GetSupportsHibernationOk returns a tuple with the SupportsHibernation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsHibernation

`func (o *MachineProfileMetadataResponseModel) SetSupportsHibernation(v bool)`

SetSupportsHibernation sets SupportsHibernation field to given value.

### HasSupportsHibernation

`func (o *MachineProfileMetadataResponseModel) HasSupportsHibernation() bool`

HasSupportsHibernation returns a boolean if a field has been set.

### SetSupportsHibernationNil

`func (o *MachineProfileMetadataResponseModel) SetSupportsHibernationNil(b bool)`

 SetSupportsHibernationNil sets the value for SupportsHibernation to be an explicit nil

### UnsetSupportsHibernation
`func (o *MachineProfileMetadataResponseModel) UnsetSupportsHibernation()`

UnsetSupportsHibernation ensures that no value is present for SupportsHibernation, not even an explicit nil
### GetZoneName

`func (o *MachineProfileMetadataResponseModel) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *MachineProfileMetadataResponseModel) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *MachineProfileMetadataResponseModel) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *MachineProfileMetadataResponseModel) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### SetZoneNameNil

`func (o *MachineProfileMetadataResponseModel) SetZoneNameNil(b bool)`

 SetZoneNameNil sets the value for ZoneName to be an explicit nil

### UnsetZoneName
`func (o *MachineProfileMetadataResponseModel) UnsetZoneName()`

UnsetZoneName ensures that no value is present for ZoneName, not even an explicit nil
### GetLabels

`func (o *MachineProfileMetadataResponseModel) GetLabels() string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MachineProfileMetadataResponseModel) GetLabelsOk() (*string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MachineProfileMetadataResponseModel) SetLabels(v string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MachineProfileMetadataResponseModel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *MachineProfileMetadataResponseModel) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *MachineProfileMetadataResponseModel) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


