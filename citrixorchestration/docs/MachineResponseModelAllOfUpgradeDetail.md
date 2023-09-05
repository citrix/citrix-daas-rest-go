# MachineResponseModelAllOfUpgradeDetail

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

### NewMachineResponseModelAllOfUpgradeDetail

`func NewMachineResponseModelAllOfUpgradeDetail() *MachineResponseModelAllOfUpgradeDetail`

NewMachineResponseModelAllOfUpgradeDetail instantiates a new MachineResponseModelAllOfUpgradeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineResponseModelAllOfUpgradeDetailWithDefaults

`func NewMachineResponseModelAllOfUpgradeDetailWithDefaults() *MachineResponseModelAllOfUpgradeDetail`

NewMachineResponseModelAllOfUpgradeDetailWithDefaults instantiates a new MachineResponseModelAllOfUpgradeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *MachineResponseModelAllOfUpgradeDetail) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MachineResponseModelAllOfUpgradeDetail) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MachineResponseModelAllOfUpgradeDetail) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MachineResponseModelAllOfUpgradeDetail) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetScheduleStatus

`func (o *MachineResponseModelAllOfUpgradeDetail) GetScheduleStatus() VdaUpgradeMachineScheduleStatus`

GetScheduleStatus returns the ScheduleStatus field if non-nil, zero value otherwise.

### GetScheduleStatusOk

`func (o *MachineResponseModelAllOfUpgradeDetail) GetScheduleStatusOk() (*VdaUpgradeMachineScheduleStatus, bool)`

GetScheduleStatusOk returns a tuple with the ScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStatus

`func (o *MachineResponseModelAllOfUpgradeDetail) SetScheduleStatus(v VdaUpgradeMachineScheduleStatus)`

SetScheduleStatus sets ScheduleStatus field to given value.

### HasScheduleStatus

`func (o *MachineResponseModelAllOfUpgradeDetail) HasScheduleStatus() bool`

HasScheduleStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *MachineResponseModelAllOfUpgradeDetail) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *MachineResponseModelAllOfUpgradeDetail) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *MachineResponseModelAllOfUpgradeDetail) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *MachineResponseModelAllOfUpgradeDetail) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastStateChangeTimeUtc

`func (o *MachineResponseModelAllOfUpgradeDetail) GetLastStateChangeTimeUtc() string`

GetLastStateChangeTimeUtc returns the LastStateChangeTimeUtc field if non-nil, zero value otherwise.

### GetLastStateChangeTimeUtcOk

`func (o *MachineResponseModelAllOfUpgradeDetail) GetLastStateChangeTimeUtcOk() (*string, bool)`

GetLastStateChangeTimeUtcOk returns a tuple with the LastStateChangeTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateChangeTimeUtc

`func (o *MachineResponseModelAllOfUpgradeDetail) SetLastStateChangeTimeUtc(v string)`

SetLastStateChangeTimeUtc sets LastStateChangeTimeUtc field to given value.

### HasLastStateChangeTimeUtc

`func (o *MachineResponseModelAllOfUpgradeDetail) HasLastStateChangeTimeUtc() bool`

HasLastStateChangeTimeUtc returns a boolean if a field has been set.

### GetScheduledTimeUtc

`func (o *MachineResponseModelAllOfUpgradeDetail) GetScheduledTimeUtc() string`

GetScheduledTimeUtc returns the ScheduledTimeUtc field if non-nil, zero value otherwise.

### GetScheduledTimeUtcOk

`func (o *MachineResponseModelAllOfUpgradeDetail) GetScheduledTimeUtcOk() (*string, bool)`

GetScheduledTimeUtcOk returns a tuple with the ScheduledTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTimeUtc

`func (o *MachineResponseModelAllOfUpgradeDetail) SetScheduledTimeUtc(v string)`

SetScheduledTimeUtc sets ScheduledTimeUtc field to given value.

### HasScheduledTimeUtc

`func (o *MachineResponseModelAllOfUpgradeDetail) HasScheduledTimeUtc() bool`

HasScheduledTimeUtc returns a boolean if a field has been set.

### GetDurationInHours

`func (o *MachineResponseModelAllOfUpgradeDetail) GetDurationInHours() int32`

GetDurationInHours returns the DurationInHours field if non-nil, zero value otherwise.

### GetDurationInHoursOk

`func (o *MachineResponseModelAllOfUpgradeDetail) GetDurationInHoursOk() (*int32, bool)`

GetDurationInHoursOk returns a tuple with the DurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInHours

`func (o *MachineResponseModelAllOfUpgradeDetail) SetDurationInHours(v int32)`

SetDurationInHours sets DurationInHours field to given value.

### HasDurationInHours

`func (o *MachineResponseModelAllOfUpgradeDetail) HasDurationInHours() bool`

HasDurationInHours returns a boolean if a field has been set.

### GetTargetPackageVersion

`func (o *MachineResponseModelAllOfUpgradeDetail) GetTargetPackageVersion() string`

GetTargetPackageVersion returns the TargetPackageVersion field if non-nil, zero value otherwise.

### GetTargetPackageVersionOk

`func (o *MachineResponseModelAllOfUpgradeDetail) GetTargetPackageVersionOk() (*string, bool)`

GetTargetPackageVersionOk returns a tuple with the TargetPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPackageVersion

`func (o *MachineResponseModelAllOfUpgradeDetail) SetTargetPackageVersion(v string)`

SetTargetPackageVersion sets TargetPackageVersion field to given value.

### HasTargetPackageVersion

`func (o *MachineResponseModelAllOfUpgradeDetail) HasTargetPackageVersion() bool`

HasTargetPackageVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


