# ImageVersionDetailResponseModelAllOfImageScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name associated with the image scheme. | [optional] 
**Id** | Pointer to **string** | The Id of the image scheme. | [optional] 
**ResourcePool** | Pointer to [**HypervisorResourcePoolRefResponseModel**](HypervisorResourcePoolRefResponseModel.md) |  | [optional] 
**HostingUnitName** | Pointer to **string** | The name of HostingUnit where image preparing happens | [optional] 
**HostingUnitId** | Pointer to **string** | The ID of HostingUnit where image preparing happens | [optional] 
**CpuCount** | Pointer to **int32** | The number of processors that VMs will be created with when using this scheme. | [optional] 
**MemoryMB** | Pointer to **int32** | The maximum amount of memory that VMs will be created with when using this scheme. | [optional] 
**ServiceOffering** | Pointer to **string** | Service offering used by Cloud provisioned VMs. | [optional] 
**NetworkMaps** | Pointer to [**[]NetworkMapResponseModel**](NetworkMapResponseModel.md) | Networks used by image scheme. | [optional] 
**MachineProfile** | Pointer to [**ImageSchemeResponseModelMachineProfile**](ImageSchemeResponseModelMachineProfile.md) |  | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image scheme that are specific to the target hosting infrastructure. | [optional] 
**CustomPropertiesInString** | Pointer to **string** | The properties of the image scheme that are specific to the target hosting infrastructure in string format. | [optional] 

## Methods

### NewImageVersionDetailResponseModelAllOfImageScheme

`func NewImageVersionDetailResponseModelAllOfImageScheme() *ImageVersionDetailResponseModelAllOfImageScheme`

NewImageVersionDetailResponseModelAllOfImageScheme instantiates a new ImageVersionDetailResponseModelAllOfImageScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionDetailResponseModelAllOfImageSchemeWithDefaults

`func NewImageVersionDetailResponseModelAllOfImageSchemeWithDefaults() *ImageVersionDetailResponseModelAllOfImageScheme`

NewImageVersionDetailResponseModelAllOfImageSchemeWithDefaults instantiates a new ImageVersionDetailResponseModelAllOfImageScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourcePool

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetResourcePool() HypervisorResourcePoolRefResponseModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetResourcePool(v HypervisorResourcePoolRefResponseModel)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetHostingUnitName

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetHostingUnitName() string`

GetHostingUnitName returns the HostingUnitName field if non-nil, zero value otherwise.

### GetHostingUnitNameOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetHostingUnitNameOk() (*string, bool)`

GetHostingUnitNameOk returns a tuple with the HostingUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnitName

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetHostingUnitName(v string)`

SetHostingUnitName sets HostingUnitName field to given value.

### HasHostingUnitName

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasHostingUnitName() bool`

HasHostingUnitName returns a boolean if a field has been set.

### GetHostingUnitId

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetHostingUnitId() string`

GetHostingUnitId returns the HostingUnitId field if non-nil, zero value otherwise.

### GetHostingUnitIdOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetHostingUnitIdOk() (*string, bool)`

GetHostingUnitIdOk returns a tuple with the HostingUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnitId

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetHostingUnitId(v string)`

SetHostingUnitId sets HostingUnitId field to given value.

### HasHostingUnitId

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasHostingUnitId() bool`

HasHostingUnitId returns a boolean if a field has been set.

### GetCpuCount

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetMemoryMB

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetServiceOffering

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetServiceOffering() string`

GetServiceOffering returns the ServiceOffering field if non-nil, zero value otherwise.

### GetServiceOfferingOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetServiceOfferingOk() (*string, bool)`

GetServiceOfferingOk returns a tuple with the ServiceOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOffering

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetServiceOffering(v string)`

SetServiceOffering sets ServiceOffering field to given value.

### HasServiceOffering

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasServiceOffering() bool`

HasServiceOffering returns a boolean if a field has been set.

### GetNetworkMaps

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetNetworkMaps() []NetworkMapResponseModel`

GetNetworkMaps returns the NetworkMaps field if non-nil, zero value otherwise.

### GetNetworkMapsOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetNetworkMapsOk() (*[]NetworkMapResponseModel, bool)`

GetNetworkMapsOk returns a tuple with the NetworkMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMaps

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetNetworkMaps(v []NetworkMapResponseModel)`

SetNetworkMaps sets NetworkMaps field to given value.

### HasNetworkMaps

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasNetworkMaps() bool`

HasNetworkMaps returns a boolean if a field has been set.

### GetMachineProfile

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetMachineProfile() ImageSchemeResponseModelMachineProfile`

GetMachineProfile returns the MachineProfile field if non-nil, zero value otherwise.

### GetMachineProfileOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetMachineProfileOk() (*ImageSchemeResponseModelMachineProfile, bool)`

GetMachineProfileOk returns a tuple with the MachineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfile

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetMachineProfile(v ImageSchemeResponseModelMachineProfile)`

SetMachineProfile sets MachineProfile field to given value.

### HasMachineProfile

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasMachineProfile() bool`

HasMachineProfile returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetCustomPropertiesInString

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetCustomPropertiesInString() string`

GetCustomPropertiesInString returns the CustomPropertiesInString field if non-nil, zero value otherwise.

### GetCustomPropertiesInStringOk

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) GetCustomPropertiesInStringOk() (*string, bool)`

GetCustomPropertiesInStringOk returns a tuple with the CustomPropertiesInString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPropertiesInString

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) SetCustomPropertiesInString(v string)`

SetCustomPropertiesInString sets CustomPropertiesInString field to given value.

### HasCustomPropertiesInString

`func (o *ImageVersionDetailResponseModelAllOfImageScheme) HasCustomPropertiesInString() bool`

HasCustomPropertiesInString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


