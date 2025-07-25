# Persona

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Id of the persona | [optional] 
**Name** | Pointer to **string** | Name of the persona | [optional] 
**AzureVmSku** | Pointer to [**AzureVMSize**](AzureVMSize.md) | The VM SKU for the persona | [optional] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) | Quantity of sessions supported per-machine. | [optional] 
**SessionsPerVm** | Pointer to **int32** | Number of sessions per vm. Value is 1 for Single Session Personas | [optional] 
**GpuMemoryInMB** | Pointer to **int32** | The GPU memory in MB | [optional] 
**EstimatedCredits** | Pointer to **int32** | The cost per user per month for the persona | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


