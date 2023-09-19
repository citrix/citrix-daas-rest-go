# ImageSchemeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name associated with the image scheme. | [optional] 
**Id** | Pointer to **string** | The Id of the image scheme. | [optional] 
**ResourcePool** | Pointer to [**HypervisorResourcePoolRefResponseModel**](HypervisorResourcePoolRefResponseModel.md) |  | [optional] 
**HostingUnitName** | Pointer to **NullableString** | The name of HostingUnit where image preparing happens | [optional] 
**HostingUnitId** | Pointer to **string** | The ID of HostingUnit where image preparing happens | [optional] 
**CpuCount** | Pointer to **int32** | The number of processors that VMs will be created with when using this scheme. | [optional] 
**MemoryMB** | Pointer to **int32** | The maximum amount of memory that VMs will be created with when using this scheme. | [optional] 
**ServiceOffering** | Pointer to **NullableString** | Service offering used by Cloud provisioned VMs. | [optional] 
**NetworkMaps** | Pointer to [**[]NetworkMapResponseModel**](NetworkMapResponseModel.md) | Networks used by image scheme. | [optional] 
**MachineProfile** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image scheme that are specific to the target hosting infrastructure. | [optional] 
**CustomPropertiesInString** | Pointer to **NullableString** | The properties of the image scheme that are specific to the target hosting infrastructure in string format. | [optional] 

## Methods

### NewImageSchemeResponseModel

`func NewImageSchemeResponseModel() *ImageSchemeResponseModel`

NewImageSchemeResponseModel instantiates a new ImageSchemeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageSchemeResponseModelWithDefaults

`func NewImageSchemeResponseModelWithDefaults() *ImageSchemeResponseModel`

NewImageSchemeResponseModelWithDefaults instantiates a new ImageSchemeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageSchemeResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageSchemeResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageSchemeResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageSchemeResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImageSchemeResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImageSchemeResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### HasId

`func (o *ImageSchemeResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourcePool

`func (o *ImageSchemeResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ImageSchemeResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ImageSchemeResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ImageSchemeResponseModel) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetHostingUnitName

`func (o *ImageSchemeResponseModel) GetHostingUnitName() string`

GetHostingUnitName returns the HostingUnitName field if non-nil, zero value otherwise.

### GetHostingUnitNameOk

`func (o *ImageSchemeResponseModel) GetHostingUnitNameOk() (*string, bool)`

GetHostingUnitNameOk returns a tuple with the HostingUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnitName

`func (o *ImageSchemeResponseModel) SetHostingUnitName(v string)`

SetHostingUnitName sets HostingUnitName field to given value.

### HasHostingUnitName

`func (o *ImageSchemeResponseModel) HasHostingUnitName() bool`

HasHostingUnitName returns a boolean if a field has been set.

### SetHostingUnitNameNil

`func (o *ImageSchemeResponseModel) SetHostingUnitNameNil(b bool)`

 SetHostingUnitNameNil sets the value for HostingUnitName to be an explicit nil

### UnsetHostingUnitName
`func (o *ImageSchemeResponseModel) UnsetHostingUnitName()`

UnsetHostingUnitName ensures that no value is present for HostingUnitName, not even an explicit nil
### GetHostingUnitId

`func (o *ImageSchemeResponseModel) GetHostingUnitId() string`

GetHostingUnitId returns the HostingUnitId field if non-nil, zero value otherwise.

### GetHostingUnitIdOk

`func (o *ImageSchemeResponseModel) GetHostingUnitIdOk() (*string, bool)`

GetHostingUnitIdOk returns a tuple with the HostingUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnitId

`func (o *ImageSchemeResponseModel) SetHostingUnitId(v string)`

SetHostingUnitId sets HostingUnitId field to given value.

### HasHostingUnitId

`func (o *ImageSchemeResponseModel) HasHostingUnitId() bool`

HasHostingUnitId returns a boolean if a field has been set.

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


