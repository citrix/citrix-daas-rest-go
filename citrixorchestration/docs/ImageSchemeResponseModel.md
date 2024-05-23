# ImageSchemeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Id of the image scheme. | 
**CpuCount** | Pointer to **int32** | The number of processors that VMs will be created with when using this scheme. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image scheme that are specific to the target hosting infrastructure. | [optional] 
**CustomPropertiesInString** | Pointer to **NullableString** | The properties of the image scheme that are specific to the target hosting infrastructure in string format. | [optional] 
**MachineProfile** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**MemoryMB** | Pointer to **int32** | The maximum amount of memory that VMs will be created with when using this scheme. | [optional] 
**NetworkMaps** | Pointer to [**[]NetworkMapResponseModel**](NetworkMapResponseModel.md) | Networks used by image scheme. | [optional] 
**ServiceOffering** | Pointer to **NullableString** | Service offering used by Cloud provisioned VMs. | [optional] 

## Methods

### NewImageSchemeResponseModel

`func NewImageSchemeResponseModel(id string, ) *ImageSchemeResponseModel`

NewImageSchemeResponseModel instantiates a new ImageSchemeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageSchemeResponseModelWithDefaults

`func NewImageSchemeResponseModelWithDefaults() *ImageSchemeResponseModel`

NewImageSchemeResponseModelWithDefaults instantiates a new ImageSchemeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageSchemeResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageSchemeResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageSchemeResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetCpuCount

`func (o *ImageSchemeResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ImageSchemeResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ImageSchemeResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *ImageSchemeResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ImageSchemeResponseModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ImageSchemeResponseModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ImageSchemeResponseModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ImageSchemeResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *ImageSchemeResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *ImageSchemeResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetCustomPropertiesInString

`func (o *ImageSchemeResponseModel) GetCustomPropertiesInString() string`

GetCustomPropertiesInString returns the CustomPropertiesInString field if non-nil, zero value otherwise.

### GetCustomPropertiesInStringOk

`func (o *ImageSchemeResponseModel) GetCustomPropertiesInStringOk() (*string, bool)`

GetCustomPropertiesInStringOk returns a tuple with the CustomPropertiesInString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPropertiesInString

`func (o *ImageSchemeResponseModel) SetCustomPropertiesInString(v string)`

SetCustomPropertiesInString sets CustomPropertiesInString field to given value.

### HasCustomPropertiesInString

`func (o *ImageSchemeResponseModel) HasCustomPropertiesInString() bool`

HasCustomPropertiesInString returns a boolean if a field has been set.

### SetCustomPropertiesInStringNil

`func (o *ImageSchemeResponseModel) SetCustomPropertiesInStringNil(b bool)`

 SetCustomPropertiesInStringNil sets the value for CustomPropertiesInString to be an explicit nil

### UnsetCustomPropertiesInString
`func (o *ImageSchemeResponseModel) UnsetCustomPropertiesInString()`

UnsetCustomPropertiesInString ensures that no value is present for CustomPropertiesInString, not even an explicit nil
### GetMachineProfile

`func (o *ImageSchemeResponseModel) GetMachineProfile() HypervisorResourceRefResponseModel`

GetMachineProfile returns the MachineProfile field if non-nil, zero value otherwise.

### GetMachineProfileOk

`func (o *ImageSchemeResponseModel) GetMachineProfileOk() (*HypervisorResourceRefResponseModel, bool)`

GetMachineProfileOk returns a tuple with the MachineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfile

`func (o *ImageSchemeResponseModel) SetMachineProfile(v HypervisorResourceRefResponseModel)`

SetMachineProfile sets MachineProfile field to given value.

### HasMachineProfile

`func (o *ImageSchemeResponseModel) HasMachineProfile() bool`

HasMachineProfile returns a boolean if a field has been set.

### GetMemoryMB

`func (o *ImageSchemeResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *ImageSchemeResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *ImageSchemeResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *ImageSchemeResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetNetworkMaps

`func (o *ImageSchemeResponseModel) GetNetworkMaps() []NetworkMapResponseModel`

GetNetworkMaps returns the NetworkMaps field if non-nil, zero value otherwise.

### GetNetworkMapsOk

`func (o *ImageSchemeResponseModel) GetNetworkMapsOk() (*[]NetworkMapResponseModel, bool)`

GetNetworkMapsOk returns a tuple with the NetworkMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMaps

`func (o *ImageSchemeResponseModel) SetNetworkMaps(v []NetworkMapResponseModel)`

SetNetworkMaps sets NetworkMaps field to given value.

### HasNetworkMaps

`func (o *ImageSchemeResponseModel) HasNetworkMaps() bool`

HasNetworkMaps returns a boolean if a field has been set.

### SetNetworkMapsNil

`func (o *ImageSchemeResponseModel) SetNetworkMapsNil(b bool)`

 SetNetworkMapsNil sets the value for NetworkMaps to be an explicit nil

### UnsetNetworkMaps
`func (o *ImageSchemeResponseModel) UnsetNetworkMaps()`

UnsetNetworkMaps ensures that no value is present for NetworkMaps, not even an explicit nil
### GetServiceOffering

`func (o *ImageSchemeResponseModel) GetServiceOffering() string`

GetServiceOffering returns the ServiceOffering field if non-nil, zero value otherwise.

### GetServiceOfferingOk

`func (o *ImageSchemeResponseModel) GetServiceOfferingOk() (*string, bool)`

GetServiceOfferingOk returns a tuple with the ServiceOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOffering

`func (o *ImageSchemeResponseModel) SetServiceOffering(v string)`

SetServiceOffering sets ServiceOffering field to given value.

### HasServiceOffering

`func (o *ImageSchemeResponseModel) HasServiceOffering() bool`

HasServiceOffering returns a boolean if a field has been set.

### SetServiceOfferingNil

`func (o *ImageSchemeResponseModel) SetServiceOfferingNil(b bool)`

 SetServiceOfferingNil sets the value for ServiceOffering to be an explicit nil

### UnsetServiceOffering
`func (o *ImageSchemeResponseModel) UnsetServiceOffering()`

UnsetServiceOffering ensures that no value is present for ServiceOffering, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


