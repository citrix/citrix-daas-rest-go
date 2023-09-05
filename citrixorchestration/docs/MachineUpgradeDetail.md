# MachineUpgradeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Uuid of this upgrade status object. | [optional] 
**ScheduleStatus** | Pointer to [**VdaUpgradeMachineScheduleStatus**](VdaUpgradeMachineScheduleStatus.md) |  | [optional] 
**StatusMessage** | Pointer to **string** | Status message resulted from the action of this upgrade status object. | [optional] 
**LastStateChangeTimeUtc** | Pointer to **string** | UTC time when this upgrade status object changed status for the last time. | [optional] 
**ScheduledTimeUtc** | Pointer to **string** | UTC time when this VDA upgrade is scheduled to start. | [optional] 
**DurationInHours** | Pointer to **int32** | Timeout duration in hours, of the current VDA upgrade schdeule. | [optional] 
**TargetPackageVersion** | Pointer to **string** | Target package version of the current VDA upgrade schdeule. | [optional] 

## Methods

### NewMachineUpgradeDetail

`func NewMachineUpgradeDetail() *MachineUpgradeDetail`

NewMachineUpgradeDetail instantiates a new MachineUpgradeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineUpgradeDetailWithDefaults

`func NewMachineUpgradeDetailWithDefaults() *MachineUpgradeDetail`

NewMachineUpgradeDetailWithDefaults instantiates a new MachineUpgradeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *MachineUpgradeDetail) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MachineUpgradeDetail) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MachineUpgradeDetail) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MachineUpgradeDetail) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetScheduleStatus

`func (o *MachineUpgradeDetail) GetScheduleStatus() VdaUpgradeMachineScheduleStatus`

GetScheduleStatus returns the ScheduleStatus field if non-nil, zero value otherwise.

### GetScheduleStatusOk

`func (o *MachineUpgradeDetail) GetScheduleStatusOk() (*VdaUpgradeMachineScheduleStatus, bool)`

GetScheduleStatusOk returns a tuple with the ScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStatus

`func (o *MachineUpgradeDetail) SetScheduleStatus(v VdaUpgradeMachineScheduleStatus)`

SetScheduleStatus sets ScheduleStatus field to given value.

### HasScheduleStatus

`func (o *MachineUpgradeDetail) HasScheduleStatus() bool`

HasScheduleStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *MachineUpgradeDetail) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *MachineUpgradeDetail) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *MachineUpgradeDetail) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *MachineUpgradeDetail) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastStateChangeTimeUtc

`func (o *MachineUpgradeDetail) GetLastStateChangeTimeUtc() string`

GetLastStateChangeTimeUtc returns the LastStateChangeTimeUtc field if non-nil, zero value otherwise.

### GetLastStateChangeTimeUtcOk

`func (o *MachineUpgradeDetail) GetLastStateChangeTimeUtcOk() (*string, bool)`

GetLastStateChangeTimeUtcOk returns a tuple with the LastStateChangeTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateChangeTimeUtc

`func (o *MachineUpgradeDetail) SetLastStateChangeTimeUtc(v string)`

SetLastStateChangeTimeUtc sets LastStateChangeTimeUtc field to given value.

### HasLastStateChangeTimeUtc

`func (o *MachineUpgradeDetail) HasLastStateChangeTimeUtc() bool`

HasLastStateChangeTimeUtc returns a boolean if a field has been set.

### GetScheduledTimeUtc

`func (o *MachineUpgradeDetail) GetScheduledTimeUtc() string`

GetScheduledTimeUtc returns the ScheduledTimeUtc field if non-nil, zero value otherwise.

### GetScheduledTimeUtcOk

`func (o *MachineUpgradeDetail) GetScheduledTimeUtcOk() (*string, bool)`

GetScheduledTimeUtcOk returns a tuple with the ScheduledTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTimeUtc

`func (o *MachineUpgradeDetail) SetScheduledTimeUtc(v string)`

SetScheduledTimeUtc sets ScheduledTimeUtc field to given value.

### HasScheduledTimeUtc

`func (o *MachineUpgradeDetail) HasScheduledTimeUtc() bool`

HasScheduledTimeUtc returns a boolean if a field has been set.

### GetDurationInHours

`func (o *MachineUpgradeDetail) GetDurationInHours() int32`

GetDurationInHours returns the DurationInHours field if non-nil, zero value otherwise.

### GetDurationInHoursOk

`func (o *MachineUpgradeDetail) GetDurationInHoursOk() (*int32, bool)`

GetDurationInHoursOk returns a tuple with the DurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInHours

`func (o *MachineUpgradeDetail) SetDurationInHours(v int32)`

SetDurationInHours sets DurationInHours field to given value.

### HasDurationInHours

`func (o *MachineUpgradeDetail) HasDurationInHours() bool`

HasDurationInHours returns a boolean if a field has been set.

### GetTargetPackageVersion

`func (o *MachineUpgradeDetail) GetTargetPackageVersion() string`

GetTargetPackageVersion returns the TargetPackageVersion field if non-nil, zero value otherwise.

### GetTargetPackageVersionOk

`func (o *MachineUpgradeDetail) GetTargetPackageVersionOk() (*string, bool)`

GetTargetPackageVersionOk returns a tuple with the TargetPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPackageVersion

`func (o *MachineUpgradeDetail) SetTargetPackageVersion(v string)`

SetTargetPackageVersion sets TargetPackageVersion field to given value.

### HasTargetPackageVersion

`func (o *MachineUpgradeDetail) HasTargetPackageVersion() bool`

HasTargetPackageVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


