# MaintenanceCycleVMOperationJobResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceCycleOperationInformation** | Pointer to **NullableString** | The operation information of maintenance cycle vm operation job | [optional] 
**JobId** | Pointer to **NullableString** | The job id of maintenance cycle vm operation job.  | [optional] 
**MachineName** | Pointer to **NullableString** | The machine name.  | [optional] 
**MaintenanceCycleId** | Pointer to **string** | The maintenance cycle id.  | [optional] 
**MaintenanceStartTime** | Pointer to **NullableTime** | The maintenance cycle vm operation job start time in UTC. | [optional] 
**MaintenanceEndTime** | Pointer to **NullableTime** | The maintenance cycle vm operation job end time in UTC. | [optional] 
**MaintenanceStatus** | Pointer to [**MaintenanceCycleStatus**](MaintenanceCycleStatus.md) |  | [optional] 
**Operation** | Pointer to [**MaintenanceCycleOperation**](MaintenanceCycleOperation.md) |  | [optional] 
**VirtualMachineSid** | Pointer to **NullableString** | Virtual machine sid. | [optional] 

## Methods

### NewMaintenanceCycleVMOperationJobResponseModel

`func NewMaintenanceCycleVMOperationJobResponseModel() *MaintenanceCycleVMOperationJobResponseModel`

NewMaintenanceCycleVMOperationJobResponseModel instantiates a new MaintenanceCycleVMOperationJobResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceCycleVMOperationJobResponseModelWithDefaults

`func NewMaintenanceCycleVMOperationJobResponseModelWithDefaults() *MaintenanceCycleVMOperationJobResponseModel`

NewMaintenanceCycleVMOperationJobResponseModelWithDefaults instantiates a new MaintenanceCycleVMOperationJobResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceCycleOperationInformation

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceCycleOperationInformation() string`

GetMaintenanceCycleOperationInformation returns the MaintenanceCycleOperationInformation field if non-nil, zero value otherwise.

### GetMaintenanceCycleOperationInformationOk

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceCycleOperationInformationOk() (*string, bool)`

GetMaintenanceCycleOperationInformationOk returns a tuple with the MaintenanceCycleOperationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleOperationInformation

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMaintenanceCycleOperationInformation(v string)`

SetMaintenanceCycleOperationInformation sets MaintenanceCycleOperationInformation field to given value.

### HasMaintenanceCycleOperationInformation

`func (o *MaintenanceCycleVMOperationJobResponseModel) HasMaintenanceCycleOperationInformation() bool`

HasMaintenanceCycleOperationInformation returns a boolean if a field has been set.

### SetMaintenanceCycleOperationInformationNil

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMaintenanceCycleOperationInformationNil(b bool)`

 SetMaintenanceCycleOperationInformationNil sets the value for MaintenanceCycleOperationInformation to be an explicit nil

### UnsetMaintenanceCycleOperationInformation
`func (o *MaintenanceCycleVMOperationJobResponseModel) UnsetMaintenanceCycleOperationInformation()`

UnsetMaintenanceCycleOperationInformation ensures that no value is present for MaintenanceCycleOperationInformation, not even an explicit nil
### GetJobId

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *MaintenanceCycleVMOperationJobResponseModel) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *MaintenanceCycleVMOperationJobResponseModel) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetMachineName

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *MaintenanceCycleVMOperationJobResponseModel) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### SetMachineNameNil

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMachineNameNil(b bool)`

 SetMachineNameNil sets the value for MachineName to be an explicit nil

### UnsetMachineName
`func (o *MaintenanceCycleVMOperationJobResponseModel) UnsetMachineName()`

UnsetMachineName ensures that no value is present for MachineName, not even an explicit nil
### GetMaintenanceCycleId

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceCycleId() string`

GetMaintenanceCycleId returns the MaintenanceCycleId field if non-nil, zero value otherwise.

### GetMaintenanceCycleIdOk

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceCycleIdOk() (*string, bool)`

GetMaintenanceCycleIdOk returns a tuple with the MaintenanceCycleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleId

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMaintenanceCycleId(v string)`

SetMaintenanceCycleId sets MaintenanceCycleId field to given value.

### HasMaintenanceCycleId

`func (o *MaintenanceCycleVMOperationJobResponseModel) HasMaintenanceCycleId() bool`

HasMaintenanceCycleId returns a boolean if a field has been set.

### GetMaintenanceStartTime

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceStartTime() time.Time`

GetMaintenanceStartTime returns the MaintenanceStartTime field if non-nil, zero value otherwise.

### GetMaintenanceStartTimeOk

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceStartTimeOk() (*time.Time, bool)`

GetMaintenanceStartTimeOk returns a tuple with the MaintenanceStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceStartTime

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMaintenanceStartTime(v time.Time)`

SetMaintenanceStartTime sets MaintenanceStartTime field to given value.

### HasMaintenanceStartTime

`func (o *MaintenanceCycleVMOperationJobResponseModel) HasMaintenanceStartTime() bool`

HasMaintenanceStartTime returns a boolean if a field has been set.

### SetMaintenanceStartTimeNil

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMaintenanceStartTimeNil(b bool)`

 SetMaintenanceStartTimeNil sets the value for MaintenanceStartTime to be an explicit nil

### UnsetMaintenanceStartTime
`func (o *MaintenanceCycleVMOperationJobResponseModel) UnsetMaintenanceStartTime()`

UnsetMaintenanceStartTime ensures that no value is present for MaintenanceStartTime, not even an explicit nil
### GetMaintenanceEndTime

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceEndTime() time.Time`

GetMaintenanceEndTime returns the MaintenanceEndTime field if non-nil, zero value otherwise.

### GetMaintenanceEndTimeOk

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceEndTimeOk() (*time.Time, bool)`

GetMaintenanceEndTimeOk returns a tuple with the MaintenanceEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceEndTime

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMaintenanceEndTime(v time.Time)`

SetMaintenanceEndTime sets MaintenanceEndTime field to given value.

### HasMaintenanceEndTime

`func (o *MaintenanceCycleVMOperationJobResponseModel) HasMaintenanceEndTime() bool`

HasMaintenanceEndTime returns a boolean if a field has been set.

### SetMaintenanceEndTimeNil

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMaintenanceEndTimeNil(b bool)`

 SetMaintenanceEndTimeNil sets the value for MaintenanceEndTime to be an explicit nil

### UnsetMaintenanceEndTime
`func (o *MaintenanceCycleVMOperationJobResponseModel) UnsetMaintenanceEndTime()`

UnsetMaintenanceEndTime ensures that no value is present for MaintenanceEndTime, not even an explicit nil
### GetMaintenanceStatus

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceStatus() MaintenanceCycleStatus`

GetMaintenanceStatus returns the MaintenanceStatus field if non-nil, zero value otherwise.

### GetMaintenanceStatusOk

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetMaintenanceStatusOk() (*MaintenanceCycleStatus, bool)`

GetMaintenanceStatusOk returns a tuple with the MaintenanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceStatus

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetMaintenanceStatus(v MaintenanceCycleStatus)`

SetMaintenanceStatus sets MaintenanceStatus field to given value.

### HasMaintenanceStatus

`func (o *MaintenanceCycleVMOperationJobResponseModel) HasMaintenanceStatus() bool`

HasMaintenanceStatus returns a boolean if a field has been set.

### GetOperation

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetOperation() MaintenanceCycleOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetOperationOk() (*MaintenanceCycleOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetOperation(v MaintenanceCycleOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *MaintenanceCycleVMOperationJobResponseModel) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetVirtualMachineSid

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetVirtualMachineSid() string`

GetVirtualMachineSid returns the VirtualMachineSid field if non-nil, zero value otherwise.

### GetVirtualMachineSidOk

`func (o *MaintenanceCycleVMOperationJobResponseModel) GetVirtualMachineSidOk() (*string, bool)`

GetVirtualMachineSidOk returns a tuple with the VirtualMachineSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineSid

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetVirtualMachineSid(v string)`

SetVirtualMachineSid sets VirtualMachineSid field to given value.

### HasVirtualMachineSid

`func (o *MaintenanceCycleVMOperationJobResponseModel) HasVirtualMachineSid() bool`

HasVirtualMachineSid returns a boolean if a field has been set.

### SetVirtualMachineSidNil

`func (o *MaintenanceCycleVMOperationJobResponseModel) SetVirtualMachineSidNil(b bool)`

 SetVirtualMachineSidNil sets the value for VirtualMachineSid to be an explicit nil

### UnsetVirtualMachineSid
`func (o *MaintenanceCycleVMOperationJobResponseModel) UnsetVirtualMachineSid()`

UnsetVirtualMachineSid ensures that no value is present for VirtualMachineSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


