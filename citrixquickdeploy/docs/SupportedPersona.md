# SupportedPersona

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | the Id of the Persona | [optional] 
**Name** | **string** | The name of the Persona | 
**VmSku** | **string** | Internal only: The VM SKU string for the persona | 
**SessionSupport** | [**SessionSupport**](SessionSupport.md) | Quantity of sessions supported per-machine. | 
**SessionsPerVm** | Pointer to **int32** | Number of sessions per vm. Value is 1 for Single Session Personas | [optional] 
**GpuMemoryInMB** | Pointer to **NullableInt32** | The GPU memory in MB | [optional] 
**EstimatedCredits** | **int32** | The cost per user per month for the persona | 
**DiskSku** | Pointer to **NullableString** | The Disk SKU for the persona | [optional] 
**DiskSize** | Pointer to **int32** | The Disk size associated with the persona | [optional] 
**IsDeprecated** | Pointer to **bool** | Specifies if the persona is deprecated | [optional] 

## Methods

### NewSupportedPersona

`func NewSupportedPersona(name string, vmSku string, sessionSupport SessionSupport, estimatedCredits int32, ) *SupportedPersona`

NewSupportedPersona instantiates a new SupportedPersona object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedPersonaWithDefaults

`func NewSupportedPersonaWithDefaults() *SupportedPersona`

NewSupportedPersonaWithDefaults instantiates a new SupportedPersona object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportedPersona) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportedPersona) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportedPersona) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportedPersona) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SupportedPersona) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SupportedPersona) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *SupportedPersona) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupportedPersona) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupportedPersona) SetName(v string)`

SetName sets Name field to given value.


### GetVmSku

`func (o *SupportedPersona) GetVmSku() string`

GetVmSku returns the VmSku field if non-nil, zero value otherwise.

### GetVmSkuOk

`func (o *SupportedPersona) GetVmSkuOk() (*string, bool)`

GetVmSkuOk returns a tuple with the VmSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSku

`func (o *SupportedPersona) SetVmSku(v string)`

SetVmSku sets VmSku field to given value.


### GetSessionSupport

`func (o *SupportedPersona) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *SupportedPersona) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *SupportedPersona) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSessionsPerVm

`func (o *SupportedPersona) GetSessionsPerVm() int32`

GetSessionsPerVm returns the SessionsPerVm field if non-nil, zero value otherwise.

### GetSessionsPerVmOk

`func (o *SupportedPersona) GetSessionsPerVmOk() (*int32, bool)`

GetSessionsPerVmOk returns a tuple with the SessionsPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsPerVm

`func (o *SupportedPersona) SetSessionsPerVm(v int32)`

SetSessionsPerVm sets SessionsPerVm field to given value.

### HasSessionsPerVm

`func (o *SupportedPersona) HasSessionsPerVm() bool`

HasSessionsPerVm returns a boolean if a field has been set.

### GetGpuMemoryInMB

`func (o *SupportedPersona) GetGpuMemoryInMB() int32`

GetGpuMemoryInMB returns the GpuMemoryInMB field if non-nil, zero value otherwise.

### GetGpuMemoryInMBOk

`func (o *SupportedPersona) GetGpuMemoryInMBOk() (*int32, bool)`

GetGpuMemoryInMBOk returns a tuple with the GpuMemoryInMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMemoryInMB

`func (o *SupportedPersona) SetGpuMemoryInMB(v int32)`

SetGpuMemoryInMB sets GpuMemoryInMB field to given value.

### HasGpuMemoryInMB

`func (o *SupportedPersona) HasGpuMemoryInMB() bool`

HasGpuMemoryInMB returns a boolean if a field has been set.

### SetGpuMemoryInMBNil

`func (o *SupportedPersona) SetGpuMemoryInMBNil(b bool)`

 SetGpuMemoryInMBNil sets the value for GpuMemoryInMB to be an explicit nil

### UnsetGpuMemoryInMB
`func (o *SupportedPersona) UnsetGpuMemoryInMB()`

UnsetGpuMemoryInMB ensures that no value is present for GpuMemoryInMB, not even an explicit nil
### GetEstimatedCredits

`func (o *SupportedPersona) GetEstimatedCredits() int32`

GetEstimatedCredits returns the EstimatedCredits field if non-nil, zero value otherwise.

### GetEstimatedCreditsOk

`func (o *SupportedPersona) GetEstimatedCreditsOk() (*int32, bool)`

GetEstimatedCreditsOk returns a tuple with the EstimatedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCredits

`func (o *SupportedPersona) SetEstimatedCredits(v int32)`

SetEstimatedCredits sets EstimatedCredits field to given value.


### GetDiskSku

`func (o *SupportedPersona) GetDiskSku() string`

GetDiskSku returns the DiskSku field if non-nil, zero value otherwise.

### GetDiskSkuOk

`func (o *SupportedPersona) GetDiskSkuOk() (*string, bool)`

GetDiskSkuOk returns a tuple with the DiskSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSku

`func (o *SupportedPersona) SetDiskSku(v string)`

SetDiskSku sets DiskSku field to given value.

### HasDiskSku

`func (o *SupportedPersona) HasDiskSku() bool`

HasDiskSku returns a boolean if a field has been set.

### SetDiskSkuNil

`func (o *SupportedPersona) SetDiskSkuNil(b bool)`

 SetDiskSkuNil sets the value for DiskSku to be an explicit nil

### UnsetDiskSku
`func (o *SupportedPersona) UnsetDiskSku()`

UnsetDiskSku ensures that no value is present for DiskSku, not even an explicit nil
### GetDiskSize

`func (o *SupportedPersona) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *SupportedPersona) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *SupportedPersona) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *SupportedPersona) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetIsDeprecated

`func (o *SupportedPersona) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *SupportedPersona) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *SupportedPersona) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *SupportedPersona) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


