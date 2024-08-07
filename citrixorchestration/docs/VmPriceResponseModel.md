# VmPriceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineSize** | Pointer to **NullableString** | The size of the virtual machine. | [optional] 
**OsType** | Pointer to [**OsType**](OsType.md) |  | [optional] 
**ByoLicense** | Pointer to **bool** | Whether bring your own license is enabled. | [optional] 
**Location** | Pointer to **NullableString** | Specifies the location of the virtual machine. | [optional] 
**InstanceType** | Pointer to [**VMInstanceType**](VMInstanceType.md) |  | [optional] 
**ResourcePrices** | Pointer to [**[]ResourcePriceResponseModel**](ResourcePriceResponseModel.md) | Gets or sets pricing information of resources. | [optional] 

## Methods

### NewVmPriceResponseModel

`func NewVmPriceResponseModel() *VmPriceResponseModel`

NewVmPriceResponseModel instantiates a new VmPriceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmPriceResponseModelWithDefaults

`func NewVmPriceResponseModelWithDefaults() *VmPriceResponseModel`

NewVmPriceResponseModelWithDefaults instantiates a new VmPriceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineSize

`func (o *VmPriceResponseModel) GetMachineSize() string`

GetMachineSize returns the MachineSize field if non-nil, zero value otherwise.

### GetMachineSizeOk

`func (o *VmPriceResponseModel) GetMachineSizeOk() (*string, bool)`

GetMachineSizeOk returns a tuple with the MachineSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineSize

`func (o *VmPriceResponseModel) SetMachineSize(v string)`

SetMachineSize sets MachineSize field to given value.

### HasMachineSize

`func (o *VmPriceResponseModel) HasMachineSize() bool`

HasMachineSize returns a boolean if a field has been set.

### SetMachineSizeNil

`func (o *VmPriceResponseModel) SetMachineSizeNil(b bool)`

 SetMachineSizeNil sets the value for MachineSize to be an explicit nil

### UnsetMachineSize
`func (o *VmPriceResponseModel) UnsetMachineSize()`

UnsetMachineSize ensures that no value is present for MachineSize, not even an explicit nil
### GetOsType

`func (o *VmPriceResponseModel) GetOsType() OsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *VmPriceResponseModel) GetOsTypeOk() (*OsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *VmPriceResponseModel) SetOsType(v OsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *VmPriceResponseModel) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetByoLicense

`func (o *VmPriceResponseModel) GetByoLicense() bool`

GetByoLicense returns the ByoLicense field if non-nil, zero value otherwise.

### GetByoLicenseOk

`func (o *VmPriceResponseModel) GetByoLicenseOk() (*bool, bool)`

GetByoLicenseOk returns a tuple with the ByoLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoLicense

`func (o *VmPriceResponseModel) SetByoLicense(v bool)`

SetByoLicense sets ByoLicense field to given value.

### HasByoLicense

`func (o *VmPriceResponseModel) HasByoLicense() bool`

HasByoLicense returns a boolean if a field has been set.

### GetLocation

`func (o *VmPriceResponseModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VmPriceResponseModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VmPriceResponseModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VmPriceResponseModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *VmPriceResponseModel) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *VmPriceResponseModel) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetInstanceType

`func (o *VmPriceResponseModel) GetInstanceType() VMInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *VmPriceResponseModel) GetInstanceTypeOk() (*VMInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *VmPriceResponseModel) SetInstanceType(v VMInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *VmPriceResponseModel) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetResourcePrices

`func (o *VmPriceResponseModel) GetResourcePrices() []ResourcePriceResponseModel`

GetResourcePrices returns the ResourcePrices field if non-nil, zero value otherwise.

### GetResourcePricesOk

`func (o *VmPriceResponseModel) GetResourcePricesOk() (*[]ResourcePriceResponseModel, bool)`

GetResourcePricesOk returns a tuple with the ResourcePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePrices

`func (o *VmPriceResponseModel) SetResourcePrices(v []ResourcePriceResponseModel)`

SetResourcePrices sets ResourcePrices field to given value.

### HasResourcePrices

`func (o *VmPriceResponseModel) HasResourcePrices() bool`

HasResourcePrices returns a boolean if a field has been set.

### SetResourcePricesNil

`func (o *VmPriceResponseModel) SetResourcePricesNil(b bool)`

 SetResourcePricesNil sets the value for ResourcePrices to be an explicit nil

### UnsetResourcePrices
`func (o *VmPriceResponseModel) UnsetResourcePrices()`

UnsetResourcePrices ensures that no value is present for ResourcePrices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


