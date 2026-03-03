# CatalogMachineEventSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfMachinesAdded** | Pointer to **int32** | Number of machines added to catalog within the queried period. | [optional] 
**NumberOfMachinesChanged** | Pointer to **int32** | Change of number of machines in catalog within the queried period.  If total number of machines were less than the start of the queried period, this will be negative.  If there were no events invovled in the queried period, this will be 0. | [optional] 
**NumberOfMachines** | Pointer to **int32** | Total number of machines at the end of the queried period. | [optional] 
**VmSku** | Pointer to **NullableString** | The VM SKU string for the catalog | [optional] 
**DiskSku** | Pointer to [**NullableCatalogCapacityStorageType**](CatalogCapacityStorageType.md) | The Disk SKU string for the catalog | [optional] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) | Quantity of sessions supported per-machine. | [optional] 
**SessionsPerVm** | Pointer to **int32** | Number of sessions per vm. Value is 1 for Single Session Personas | [optional] 
**Region** | Pointer to **NullableString** | Region of the catalog | [optional] 
**PersonaId** | Pointer to **NullableString** | The ID of the Persona associated with the catalog | [optional] 
**LastRecordedEventDate** | Pointer to **time.Time** | The timestamp of the last event recorded in the queried period. | [optional] 

## Methods

### NewCatalogMachineEventSummary

`func NewCatalogMachineEventSummary() *CatalogMachineEventSummary`

NewCatalogMachineEventSummary instantiates a new CatalogMachineEventSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogMachineEventSummaryWithDefaults

`func NewCatalogMachineEventSummaryWithDefaults() *CatalogMachineEventSummary`

NewCatalogMachineEventSummaryWithDefaults instantiates a new CatalogMachineEventSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfMachinesAdded

`func (o *CatalogMachineEventSummary) GetNumberOfMachinesAdded() int32`

GetNumberOfMachinesAdded returns the NumberOfMachinesAdded field if non-nil, zero value otherwise.

### GetNumberOfMachinesAddedOk

`func (o *CatalogMachineEventSummary) GetNumberOfMachinesAddedOk() (*int32, bool)`

GetNumberOfMachinesAddedOk returns a tuple with the NumberOfMachinesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfMachinesAdded

`func (o *CatalogMachineEventSummary) SetNumberOfMachinesAdded(v int32)`

SetNumberOfMachinesAdded sets NumberOfMachinesAdded field to given value.

### HasNumberOfMachinesAdded

`func (o *CatalogMachineEventSummary) HasNumberOfMachinesAdded() bool`

HasNumberOfMachinesAdded returns a boolean if a field has been set.

### GetNumberOfMachinesChanged

`func (o *CatalogMachineEventSummary) GetNumberOfMachinesChanged() int32`

GetNumberOfMachinesChanged returns the NumberOfMachinesChanged field if non-nil, zero value otherwise.

### GetNumberOfMachinesChangedOk

`func (o *CatalogMachineEventSummary) GetNumberOfMachinesChangedOk() (*int32, bool)`

GetNumberOfMachinesChangedOk returns a tuple with the NumberOfMachinesChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfMachinesChanged

`func (o *CatalogMachineEventSummary) SetNumberOfMachinesChanged(v int32)`

SetNumberOfMachinesChanged sets NumberOfMachinesChanged field to given value.

### HasNumberOfMachinesChanged

`func (o *CatalogMachineEventSummary) HasNumberOfMachinesChanged() bool`

HasNumberOfMachinesChanged returns a boolean if a field has been set.

### GetNumberOfMachines

`func (o *CatalogMachineEventSummary) GetNumberOfMachines() int32`

GetNumberOfMachines returns the NumberOfMachines field if non-nil, zero value otherwise.

### GetNumberOfMachinesOk

`func (o *CatalogMachineEventSummary) GetNumberOfMachinesOk() (*int32, bool)`

GetNumberOfMachinesOk returns a tuple with the NumberOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfMachines

`func (o *CatalogMachineEventSummary) SetNumberOfMachines(v int32)`

SetNumberOfMachines sets NumberOfMachines field to given value.

### HasNumberOfMachines

`func (o *CatalogMachineEventSummary) HasNumberOfMachines() bool`

HasNumberOfMachines returns a boolean if a field has been set.

### GetVmSku

`func (o *CatalogMachineEventSummary) GetVmSku() string`

GetVmSku returns the VmSku field if non-nil, zero value otherwise.

### GetVmSkuOk

`func (o *CatalogMachineEventSummary) GetVmSkuOk() (*string, bool)`

GetVmSkuOk returns a tuple with the VmSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSku

`func (o *CatalogMachineEventSummary) SetVmSku(v string)`

SetVmSku sets VmSku field to given value.

### HasVmSku

`func (o *CatalogMachineEventSummary) HasVmSku() bool`

HasVmSku returns a boolean if a field has been set.

### SetVmSkuNil

`func (o *CatalogMachineEventSummary) SetVmSkuNil(b bool)`

 SetVmSkuNil sets the value for VmSku to be an explicit nil

### UnsetVmSku
`func (o *CatalogMachineEventSummary) UnsetVmSku()`

UnsetVmSku ensures that no value is present for VmSku, not even an explicit nil
### GetDiskSku

`func (o *CatalogMachineEventSummary) GetDiskSku() CatalogCapacityStorageType`

GetDiskSku returns the DiskSku field if non-nil, zero value otherwise.

### GetDiskSkuOk

`func (o *CatalogMachineEventSummary) GetDiskSkuOk() (*CatalogCapacityStorageType, bool)`

GetDiskSkuOk returns a tuple with the DiskSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSku

`func (o *CatalogMachineEventSummary) SetDiskSku(v CatalogCapacityStorageType)`

SetDiskSku sets DiskSku field to given value.

### HasDiskSku

`func (o *CatalogMachineEventSummary) HasDiskSku() bool`

HasDiskSku returns a boolean if a field has been set.

### SetDiskSkuNil

`func (o *CatalogMachineEventSummary) SetDiskSkuNil(b bool)`

 SetDiskSkuNil sets the value for DiskSku to be an explicit nil

### UnsetDiskSku
`func (o *CatalogMachineEventSummary) UnsetDiskSku()`

UnsetDiskSku ensures that no value is present for DiskSku, not even an explicit nil
### GetSessionSupport

`func (o *CatalogMachineEventSummary) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *CatalogMachineEventSummary) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *CatalogMachineEventSummary) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *CatalogMachineEventSummary) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### GetSessionsPerVm

`func (o *CatalogMachineEventSummary) GetSessionsPerVm() int32`

GetSessionsPerVm returns the SessionsPerVm field if non-nil, zero value otherwise.

### GetSessionsPerVmOk

`func (o *CatalogMachineEventSummary) GetSessionsPerVmOk() (*int32, bool)`

GetSessionsPerVmOk returns a tuple with the SessionsPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsPerVm

`func (o *CatalogMachineEventSummary) SetSessionsPerVm(v int32)`

SetSessionsPerVm sets SessionsPerVm field to given value.

### HasSessionsPerVm

`func (o *CatalogMachineEventSummary) HasSessionsPerVm() bool`

HasSessionsPerVm returns a boolean if a field has been set.

### GetRegion

`func (o *CatalogMachineEventSummary) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CatalogMachineEventSummary) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CatalogMachineEventSummary) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CatalogMachineEventSummary) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CatalogMachineEventSummary) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CatalogMachineEventSummary) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPersonaId

`func (o *CatalogMachineEventSummary) GetPersonaId() string`

GetPersonaId returns the PersonaId field if non-nil, zero value otherwise.

### GetPersonaIdOk

`func (o *CatalogMachineEventSummary) GetPersonaIdOk() (*string, bool)`

GetPersonaIdOk returns a tuple with the PersonaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaId

`func (o *CatalogMachineEventSummary) SetPersonaId(v string)`

SetPersonaId sets PersonaId field to given value.

### HasPersonaId

`func (o *CatalogMachineEventSummary) HasPersonaId() bool`

HasPersonaId returns a boolean if a field has been set.

### SetPersonaIdNil

`func (o *CatalogMachineEventSummary) SetPersonaIdNil(b bool)`

 SetPersonaIdNil sets the value for PersonaId to be an explicit nil

### UnsetPersonaId
`func (o *CatalogMachineEventSummary) UnsetPersonaId()`

UnsetPersonaId ensures that no value is present for PersonaId, not even an explicit nil
### GetLastRecordedEventDate

`func (o *CatalogMachineEventSummary) GetLastRecordedEventDate() time.Time`

GetLastRecordedEventDate returns the LastRecordedEventDate field if non-nil, zero value otherwise.

### GetLastRecordedEventDateOk

`func (o *CatalogMachineEventSummary) GetLastRecordedEventDateOk() (*time.Time, bool)`

GetLastRecordedEventDateOk returns a tuple with the LastRecordedEventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecordedEventDate

`func (o *CatalogMachineEventSummary) SetLastRecordedEventDate(v time.Time)`

SetLastRecordedEventDate sets LastRecordedEventDate field to given value.

### HasLastRecordedEventDate

`func (o *CatalogMachineEventSummary) HasLastRecordedEventDate() bool`

HasLastRecordedEventDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


