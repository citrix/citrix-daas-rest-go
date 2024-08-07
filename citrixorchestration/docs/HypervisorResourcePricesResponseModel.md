# HypervisorResourcePricesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmPrices** | Pointer to [**[]VmPriceResponseModel**](VmPriceResponseModel.md) | The pricing information for the virtual machines settings. | [optional] 
**DiskPrices** | Pointer to [**[]DiskPriceResponseModel**](DiskPriceResponseModel.md) | The pricing information for the disk settings. | [optional] 
**WarningMessage** | Pointer to **NullableString** | The warning message. | [optional] 

## Methods

### NewHypervisorResourcePricesResponseModel

`func NewHypervisorResourcePricesResponseModel() *HypervisorResourcePricesResponseModel`

NewHypervisorResourcePricesResponseModel instantiates a new HypervisorResourcePricesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePricesResponseModelWithDefaults

`func NewHypervisorResourcePricesResponseModelWithDefaults() *HypervisorResourcePricesResponseModel`

NewHypervisorResourcePricesResponseModelWithDefaults instantiates a new HypervisorResourcePricesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmPrices

`func (o *HypervisorResourcePricesResponseModel) GetVmPrices() []VmPriceResponseModel`

GetVmPrices returns the VmPrices field if non-nil, zero value otherwise.

### GetVmPricesOk

`func (o *HypervisorResourcePricesResponseModel) GetVmPricesOk() (*[]VmPriceResponseModel, bool)`

GetVmPricesOk returns a tuple with the VmPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPrices

`func (o *HypervisorResourcePricesResponseModel) SetVmPrices(v []VmPriceResponseModel)`

SetVmPrices sets VmPrices field to given value.

### HasVmPrices

`func (o *HypervisorResourcePricesResponseModel) HasVmPrices() bool`

HasVmPrices returns a boolean if a field has been set.

### SetVmPricesNil

`func (o *HypervisorResourcePricesResponseModel) SetVmPricesNil(b bool)`

 SetVmPricesNil sets the value for VmPrices to be an explicit nil

### UnsetVmPrices
`func (o *HypervisorResourcePricesResponseModel) UnsetVmPrices()`

UnsetVmPrices ensures that no value is present for VmPrices, not even an explicit nil
### GetDiskPrices

`func (o *HypervisorResourcePricesResponseModel) GetDiskPrices() []DiskPriceResponseModel`

GetDiskPrices returns the DiskPrices field if non-nil, zero value otherwise.

### GetDiskPricesOk

`func (o *HypervisorResourcePricesResponseModel) GetDiskPricesOk() (*[]DiskPriceResponseModel, bool)`

GetDiskPricesOk returns a tuple with the DiskPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPrices

`func (o *HypervisorResourcePricesResponseModel) SetDiskPrices(v []DiskPriceResponseModel)`

SetDiskPrices sets DiskPrices field to given value.

### HasDiskPrices

`func (o *HypervisorResourcePricesResponseModel) HasDiskPrices() bool`

HasDiskPrices returns a boolean if a field has been set.

### SetDiskPricesNil

`func (o *HypervisorResourcePricesResponseModel) SetDiskPricesNil(b bool)`

 SetDiskPricesNil sets the value for DiskPrices to be an explicit nil

### UnsetDiskPrices
`func (o *HypervisorResourcePricesResponseModel) UnsetDiskPrices()`

UnsetDiskPrices ensures that no value is present for DiskPrices, not even an explicit nil
### GetWarningMessage

`func (o *HypervisorResourcePricesResponseModel) GetWarningMessage() string`

GetWarningMessage returns the WarningMessage field if non-nil, zero value otherwise.

### GetWarningMessageOk

`func (o *HypervisorResourcePricesResponseModel) GetWarningMessageOk() (*string, bool)`

GetWarningMessageOk returns a tuple with the WarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessage

`func (o *HypervisorResourcePricesResponseModel) SetWarningMessage(v string)`

SetWarningMessage sets WarningMessage field to given value.

### HasWarningMessage

`func (o *HypervisorResourcePricesResponseModel) HasWarningMessage() bool`

HasWarningMessage returns a boolean if a field has been set.

### SetWarningMessageNil

`func (o *HypervisorResourcePricesResponseModel) SetWarningMessageNil(b bool)`

 SetWarningMessageNil sets the value for WarningMessage to be an explicit nil

### UnsetWarningMessage
`func (o *HypervisorResourcePricesResponseModel) UnsetWarningMessage()`

UnsetWarningMessage ensures that no value is present for WarningMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


