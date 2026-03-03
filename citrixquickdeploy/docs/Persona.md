# Persona

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The Id of the persona | [optional] 
**Name** | Pointer to **NullableString** | Name of the persona | [optional] 
**Description** | Pointer to **NullableString** | Description of the persona | [optional] 
**AzureVmSku** | Pointer to [**NullableAzureVMSize**](AzureVMSize.md) | The VM SKU for the persona | [optional] [readonly] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) | Quantity of sessions supported per-machine. | [optional] 
**SessionsPerVm** | Pointer to **int32** | Number of sessions per vm. Value is 1 for Single Session Personas | [optional] 
**GpuMemoryInMB** | Pointer to **NullableInt32** | The GPU memory in MB | [optional] 
**EstimatedCredits** | Pointer to **int32** | The cost per user per month for the persona | [optional] 
**DiskSku** | Pointer to **NullableString** | The Disk SKU for the persona | [optional] 
**DiskSize** | Pointer to **int32** | The Disk size associated with the persona | [optional] 
**AllocationType** | Pointer to [**CatalogAllocationType**](CatalogAllocationType.md) | The allocation type for the persona | [optional] 
**AddOns** | Pointer to [**[]SupportedPersonaAddOn**](SupportedPersonaAddOn.md) | The add-ons supported by the persona | [optional] 

## Methods

### NewPersona

`func NewPersona() *Persona`

NewPersona instantiates a new Persona object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonaWithDefaults

`func NewPersonaWithDefaults() *Persona`

NewPersonaWithDefaults instantiates a new Persona object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Persona) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Persona) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Persona) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Persona) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Persona) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Persona) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Persona) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Persona) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Persona) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Persona) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Persona) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Persona) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Persona) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Persona) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Persona) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Persona) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Persona) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Persona) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAzureVmSku

`func (o *Persona) GetAzureVmSku() AzureVMSize`

GetAzureVmSku returns the AzureVmSku field if non-nil, zero value otherwise.

### GetAzureVmSkuOk

`func (o *Persona) GetAzureVmSkuOk() (*AzureVMSize, bool)`

GetAzureVmSkuOk returns a tuple with the AzureVmSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVmSku

`func (o *Persona) SetAzureVmSku(v AzureVMSize)`

SetAzureVmSku sets AzureVmSku field to given value.

### HasAzureVmSku

`func (o *Persona) HasAzureVmSku() bool`

HasAzureVmSku returns a boolean if a field has been set.

### SetAzureVmSkuNil

`func (o *Persona) SetAzureVmSkuNil(b bool)`

 SetAzureVmSkuNil sets the value for AzureVmSku to be an explicit nil

### UnsetAzureVmSku
`func (o *Persona) UnsetAzureVmSku()`

UnsetAzureVmSku ensures that no value is present for AzureVmSku, not even an explicit nil
### GetSessionSupport

`func (o *Persona) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *Persona) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *Persona) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *Persona) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### GetSessionsPerVm

`func (o *Persona) GetSessionsPerVm() int32`

GetSessionsPerVm returns the SessionsPerVm field if non-nil, zero value otherwise.

### GetSessionsPerVmOk

`func (o *Persona) GetSessionsPerVmOk() (*int32, bool)`

GetSessionsPerVmOk returns a tuple with the SessionsPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsPerVm

`func (o *Persona) SetSessionsPerVm(v int32)`

SetSessionsPerVm sets SessionsPerVm field to given value.

### HasSessionsPerVm

`func (o *Persona) HasSessionsPerVm() bool`

HasSessionsPerVm returns a boolean if a field has been set.

### GetGpuMemoryInMB

`func (o *Persona) GetGpuMemoryInMB() int32`

GetGpuMemoryInMB returns the GpuMemoryInMB field if non-nil, zero value otherwise.

### GetGpuMemoryInMBOk

`func (o *Persona) GetGpuMemoryInMBOk() (*int32, bool)`

GetGpuMemoryInMBOk returns a tuple with the GpuMemoryInMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMemoryInMB

`func (o *Persona) SetGpuMemoryInMB(v int32)`

SetGpuMemoryInMB sets GpuMemoryInMB field to given value.

### HasGpuMemoryInMB

`func (o *Persona) HasGpuMemoryInMB() bool`

HasGpuMemoryInMB returns a boolean if a field has been set.

### SetGpuMemoryInMBNil

`func (o *Persona) SetGpuMemoryInMBNil(b bool)`

 SetGpuMemoryInMBNil sets the value for GpuMemoryInMB to be an explicit nil

### UnsetGpuMemoryInMB
`func (o *Persona) UnsetGpuMemoryInMB()`

UnsetGpuMemoryInMB ensures that no value is present for GpuMemoryInMB, not even an explicit nil
### GetEstimatedCredits

`func (o *Persona) GetEstimatedCredits() int32`

GetEstimatedCredits returns the EstimatedCredits field if non-nil, zero value otherwise.

### GetEstimatedCreditsOk

`func (o *Persona) GetEstimatedCreditsOk() (*int32, bool)`

GetEstimatedCreditsOk returns a tuple with the EstimatedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCredits

`func (o *Persona) SetEstimatedCredits(v int32)`

SetEstimatedCredits sets EstimatedCredits field to given value.

### HasEstimatedCredits

`func (o *Persona) HasEstimatedCredits() bool`

HasEstimatedCredits returns a boolean if a field has been set.

### GetDiskSku

`func (o *Persona) GetDiskSku() string`

GetDiskSku returns the DiskSku field if non-nil, zero value otherwise.

### GetDiskSkuOk

`func (o *Persona) GetDiskSkuOk() (*string, bool)`

GetDiskSkuOk returns a tuple with the DiskSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSku

`func (o *Persona) SetDiskSku(v string)`

SetDiskSku sets DiskSku field to given value.

### HasDiskSku

`func (o *Persona) HasDiskSku() bool`

HasDiskSku returns a boolean if a field has been set.

### SetDiskSkuNil

`func (o *Persona) SetDiskSkuNil(b bool)`

 SetDiskSkuNil sets the value for DiskSku to be an explicit nil

### UnsetDiskSku
`func (o *Persona) UnsetDiskSku()`

UnsetDiskSku ensures that no value is present for DiskSku, not even an explicit nil
### GetDiskSize

`func (o *Persona) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *Persona) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *Persona) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *Persona) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetAllocationType

`func (o *Persona) GetAllocationType() CatalogAllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *Persona) GetAllocationTypeOk() (*CatalogAllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *Persona) SetAllocationType(v CatalogAllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *Persona) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetAddOns

`func (o *Persona) GetAddOns() []SupportedPersonaAddOn`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *Persona) GetAddOnsOk() (*[]SupportedPersonaAddOn, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *Persona) SetAddOns(v []SupportedPersonaAddOn)`

SetAddOns sets AddOns field to given value.

### HasAddOns

`func (o *Persona) HasAddOns() bool`

HasAddOns returns a boolean if a field has been set.

### SetAddOnsNil

`func (o *Persona) SetAddOnsNil(b bool)`

 SetAddOnsNil sets the value for AddOns to be an explicit nil

### UnsetAddOns
`func (o *Persona) UnsetAddOns()`

UnsetAddOns ensures that no value is present for AddOns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


