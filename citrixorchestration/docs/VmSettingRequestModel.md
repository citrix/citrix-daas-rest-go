# VmSettingRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineSize** | Pointer to **NullableString** | The size of the virtual machine. | [optional] 
**OsType** | Pointer to [**OsType**](OsType.md) |  | [optional] 
**ByoLicense** | Pointer to **bool** | Whether bring your own license is enabled. | [optional] 
**Location** | Pointer to **NullableString** | Specifies the location of the virtual machine. | [optional] 
**InstanceType** | Pointer to [**VMInstanceType**](VMInstanceType.md) |  | [optional] 

## Methods

### NewVmSettingRequestModel

`func NewVmSettingRequestModel() *VmSettingRequestModel`

NewVmSettingRequestModel instantiates a new VmSettingRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmSettingRequestModelWithDefaults

`func NewVmSettingRequestModelWithDefaults() *VmSettingRequestModel`

NewVmSettingRequestModelWithDefaults instantiates a new VmSettingRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineSize

`func (o *VmSettingRequestModel) GetMachineSize() string`

GetMachineSize returns the MachineSize field if non-nil, zero value otherwise.

### GetMachineSizeOk

`func (o *VmSettingRequestModel) GetMachineSizeOk() (*string, bool)`

GetMachineSizeOk returns a tuple with the MachineSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineSize

`func (o *VmSettingRequestModel) SetMachineSize(v string)`

SetMachineSize sets MachineSize field to given value.

### HasMachineSize

`func (o *VmSettingRequestModel) HasMachineSize() bool`

HasMachineSize returns a boolean if a field has been set.

### SetMachineSizeNil

`func (o *VmSettingRequestModel) SetMachineSizeNil(b bool)`

 SetMachineSizeNil sets the value for MachineSize to be an explicit nil

### UnsetMachineSize
`func (o *VmSettingRequestModel) UnsetMachineSize()`

UnsetMachineSize ensures that no value is present for MachineSize, not even an explicit nil
### GetOsType

`func (o *VmSettingRequestModel) GetOsType() OsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *VmSettingRequestModel) GetOsTypeOk() (*OsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *VmSettingRequestModel) SetOsType(v OsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *VmSettingRequestModel) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetByoLicense

`func (o *VmSettingRequestModel) GetByoLicense() bool`

GetByoLicense returns the ByoLicense field if non-nil, zero value otherwise.

### GetByoLicenseOk

`func (o *VmSettingRequestModel) GetByoLicenseOk() (*bool, bool)`

GetByoLicenseOk returns a tuple with the ByoLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoLicense

`func (o *VmSettingRequestModel) SetByoLicense(v bool)`

SetByoLicense sets ByoLicense field to given value.

### HasByoLicense

`func (o *VmSettingRequestModel) HasByoLicense() bool`

HasByoLicense returns a boolean if a field has been set.

### GetLocation

`func (o *VmSettingRequestModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VmSettingRequestModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VmSettingRequestModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VmSettingRequestModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *VmSettingRequestModel) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *VmSettingRequestModel) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetInstanceType

`func (o *VmSettingRequestModel) GetInstanceType() VMInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *VmSettingRequestModel) GetInstanceTypeOk() (*VMInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *VmSettingRequestModel) SetInstanceType(v VMInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *VmSettingRequestModel) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


