# MachineHostingResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostedMachineId** | Pointer to **NullableString** | Unique ID within the hosting unit of the target managed machine. | [optional] 
**HostedMachineName** | Pointer to **NullableString** | The friendly name of a hosted machine as used by its hypervisor. This is not necessarily the DNS name of the machine. | [optional] 
**HostingServerName** | Pointer to **NullableString** | DNS name of the hypervisor that is hosting the machine. | [optional] 
**LastHostingUpdateTime** | Pointer to **NullableString** | Time of last update to any hosting data (such as power states) for this machine reported by the hypervisor connection. | [optional] 
**FormattedLastHostingUpdateTime** | Pointer to **NullableString** | Formatted time of last update to any hosting data (such as power states) for this machine reported by the hypervisor connection. RFC 3339 compatible format. | [optional] 
**HypervisorConnection** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**ImageOutOfDate** | Pointer to **bool** | Indicates whether the machine image matches the latest image specified for the machine catalog. | [optional] 

## Methods

### NewMachineHostingResponseModel

`func NewMachineHostingResponseModel() *MachineHostingResponseModel`

NewMachineHostingResponseModel instantiates a new MachineHostingResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineHostingResponseModelWithDefaults

`func NewMachineHostingResponseModelWithDefaults() *MachineHostingResponseModel`

NewMachineHostingResponseModelWithDefaults instantiates a new MachineHostingResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostedMachineId

`func (o *MachineHostingResponseModel) GetHostedMachineId() string`

GetHostedMachineId returns the HostedMachineId field if non-nil, zero value otherwise.

### GetHostedMachineIdOk

`func (o *MachineHostingResponseModel) GetHostedMachineIdOk() (*string, bool)`

GetHostedMachineIdOk returns a tuple with the HostedMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineId

`func (o *MachineHostingResponseModel) SetHostedMachineId(v string)`

SetHostedMachineId sets HostedMachineId field to given value.

### HasHostedMachineId

`func (o *MachineHostingResponseModel) HasHostedMachineId() bool`

HasHostedMachineId returns a boolean if a field has been set.

### SetHostedMachineIdNil

`func (o *MachineHostingResponseModel) SetHostedMachineIdNil(b bool)`

 SetHostedMachineIdNil sets the value for HostedMachineId to be an explicit nil

### UnsetHostedMachineId
`func (o *MachineHostingResponseModel) UnsetHostedMachineId()`

UnsetHostedMachineId ensures that no value is present for HostedMachineId, not even an explicit nil
### GetHostedMachineName

`func (o *MachineHostingResponseModel) GetHostedMachineName() string`

GetHostedMachineName returns the HostedMachineName field if non-nil, zero value otherwise.

### GetHostedMachineNameOk

`func (o *MachineHostingResponseModel) GetHostedMachineNameOk() (*string, bool)`

GetHostedMachineNameOk returns a tuple with the HostedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineName

`func (o *MachineHostingResponseModel) SetHostedMachineName(v string)`

SetHostedMachineName sets HostedMachineName field to given value.

### HasHostedMachineName

`func (o *MachineHostingResponseModel) HasHostedMachineName() bool`

HasHostedMachineName returns a boolean if a field has been set.

### SetHostedMachineNameNil

`func (o *MachineHostingResponseModel) SetHostedMachineNameNil(b bool)`

 SetHostedMachineNameNil sets the value for HostedMachineName to be an explicit nil

### UnsetHostedMachineName
`func (o *MachineHostingResponseModel) UnsetHostedMachineName()`

UnsetHostedMachineName ensures that no value is present for HostedMachineName, not even an explicit nil
### GetHostingServerName

`func (o *MachineHostingResponseModel) GetHostingServerName() string`

GetHostingServerName returns the HostingServerName field if non-nil, zero value otherwise.

### GetHostingServerNameOk

`func (o *MachineHostingResponseModel) GetHostingServerNameOk() (*string, bool)`

GetHostingServerNameOk returns a tuple with the HostingServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingServerName

`func (o *MachineHostingResponseModel) SetHostingServerName(v string)`

SetHostingServerName sets HostingServerName field to given value.

### HasHostingServerName

`func (o *MachineHostingResponseModel) HasHostingServerName() bool`

HasHostingServerName returns a boolean if a field has been set.

### SetHostingServerNameNil

`func (o *MachineHostingResponseModel) SetHostingServerNameNil(b bool)`

 SetHostingServerNameNil sets the value for HostingServerName to be an explicit nil

### UnsetHostingServerName
`func (o *MachineHostingResponseModel) UnsetHostingServerName()`

UnsetHostingServerName ensures that no value is present for HostingServerName, not even an explicit nil
### GetLastHostingUpdateTime

`func (o *MachineHostingResponseModel) GetLastHostingUpdateTime() string`

GetLastHostingUpdateTime returns the LastHostingUpdateTime field if non-nil, zero value otherwise.

### GetLastHostingUpdateTimeOk

`func (o *MachineHostingResponseModel) GetLastHostingUpdateTimeOk() (*string, bool)`

GetLastHostingUpdateTimeOk returns a tuple with the LastHostingUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHostingUpdateTime

`func (o *MachineHostingResponseModel) SetLastHostingUpdateTime(v string)`

SetLastHostingUpdateTime sets LastHostingUpdateTime field to given value.

### HasLastHostingUpdateTime

`func (o *MachineHostingResponseModel) HasLastHostingUpdateTime() bool`

HasLastHostingUpdateTime returns a boolean if a field has been set.

### SetLastHostingUpdateTimeNil

`func (o *MachineHostingResponseModel) SetLastHostingUpdateTimeNil(b bool)`

 SetLastHostingUpdateTimeNil sets the value for LastHostingUpdateTime to be an explicit nil

### UnsetLastHostingUpdateTime
`func (o *MachineHostingResponseModel) UnsetLastHostingUpdateTime()`

UnsetLastHostingUpdateTime ensures that no value is present for LastHostingUpdateTime, not even an explicit nil
### GetFormattedLastHostingUpdateTime

`func (o *MachineHostingResponseModel) GetFormattedLastHostingUpdateTime() string`

GetFormattedLastHostingUpdateTime returns the FormattedLastHostingUpdateTime field if non-nil, zero value otherwise.

### GetFormattedLastHostingUpdateTimeOk

`func (o *MachineHostingResponseModel) GetFormattedLastHostingUpdateTimeOk() (*string, bool)`

GetFormattedLastHostingUpdateTimeOk returns a tuple with the FormattedLastHostingUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastHostingUpdateTime

`func (o *MachineHostingResponseModel) SetFormattedLastHostingUpdateTime(v string)`

SetFormattedLastHostingUpdateTime sets FormattedLastHostingUpdateTime field to given value.

### HasFormattedLastHostingUpdateTime

`func (o *MachineHostingResponseModel) HasFormattedLastHostingUpdateTime() bool`

HasFormattedLastHostingUpdateTime returns a boolean if a field has been set.

### SetFormattedLastHostingUpdateTimeNil

`func (o *MachineHostingResponseModel) SetFormattedLastHostingUpdateTimeNil(b bool)`

 SetFormattedLastHostingUpdateTimeNil sets the value for FormattedLastHostingUpdateTime to be an explicit nil

### UnsetFormattedLastHostingUpdateTime
`func (o *MachineHostingResponseModel) UnsetFormattedLastHostingUpdateTime()`

UnsetFormattedLastHostingUpdateTime ensures that no value is present for FormattedLastHostingUpdateTime, not even an explicit nil
### GetHypervisorConnection

`func (o *MachineHostingResponseModel) GetHypervisorConnection() RefResponseModel`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *MachineHostingResponseModel) GetHypervisorConnectionOk() (*RefResponseModel, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *MachineHostingResponseModel) SetHypervisorConnection(v RefResponseModel)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *MachineHostingResponseModel) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### GetImageOutOfDate

`func (o *MachineHostingResponseModel) GetImageOutOfDate() bool`

GetImageOutOfDate returns the ImageOutOfDate field if non-nil, zero value otherwise.

### GetImageOutOfDateOk

`func (o *MachineHostingResponseModel) GetImageOutOfDateOk() (*bool, bool)`

GetImageOutOfDateOk returns a tuple with the ImageOutOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOutOfDate

`func (o *MachineHostingResponseModel) SetImageOutOfDate(v bool)`

SetImageOutOfDate sets ImageOutOfDate field to given value.

### HasImageOutOfDate

`func (o *MachineHostingResponseModel) HasImageOutOfDate() bool`

HasImageOutOfDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


