# MachineCatalogUpgradeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleStatus** | Pointer to [**VdaUpgradeScheduleStatus**](VdaUpgradeScheduleStatus.md) |  | [optional] 
**LastStateChangeTimeUtc** | Pointer to **time.Time** | UTC time when this upgrade status object changed status for the last time. | [optional] 
**TotalMachines** | Pointer to **int32** | Count of machines with an upgrade schedule in this catalog. This does not always equal to the machine count of the catalog, as there might be machines which joined the catalog after a catalog level upgrade is scheduled, so do not have an upgrade status. | [optional] 
**SuccessCount** | Pointer to **int32** | Count of machines whose last upgrade is in success state. | [optional] 
**ValidationFailureCount** | Pointer to **int32** | Count of machines whose last upgrade failed during package validation. | [optional] 
**InProgressCount** | Pointer to **int32** | Count of machines whose upgrade is in progress. | [optional] 
**UpgradeFailureCount** | Pointer to **int32** | Count of machines whose last upgrade failed during package installation. | [optional] 
**ScheduledTimeUtc** | Pointer to **time.Time** | UTC time when this VDA upgrade was scheduled to start. | [optional] 
**DurationInHours** | Pointer to **int32** | Timeout duration in hours, of the current VDA upgrade schedule. | [optional] 
**TargetPackageVersion** | Pointer to **NullableString** | Target package version of the current VDA upgrade schedule. | [optional] 
**CancelledUpgradeCount** | Pointer to **int32** | Count of machines whose last upgrade canceled during package installation. | [optional] 
**WaitingToUpgradeCount** | Pointer to **int32** | Count of machines who is waiting to upgrade. | [optional] 
**AvailableForUpgradeCount** | Pointer to **NullableInt32** | Count of machines who is available for upgrade. | [optional] 

## Methods

### NewMachineCatalogUpgradeDetail

`func NewMachineCatalogUpgradeDetail() *MachineCatalogUpgradeDetail`

NewMachineCatalogUpgradeDetail instantiates a new MachineCatalogUpgradeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogUpgradeDetailWithDefaults

`func NewMachineCatalogUpgradeDetailWithDefaults() *MachineCatalogUpgradeDetail`

NewMachineCatalogUpgradeDetailWithDefaults instantiates a new MachineCatalogUpgradeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleStatus

`func (o *MachineCatalogUpgradeDetail) GetScheduleStatus() VdaUpgradeScheduleStatus`

GetScheduleStatus returns the ScheduleStatus field if non-nil, zero value otherwise.

### GetScheduleStatusOk

`func (o *MachineCatalogUpgradeDetail) GetScheduleStatusOk() (*VdaUpgradeScheduleStatus, bool)`

GetScheduleStatusOk returns a tuple with the ScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStatus

`func (o *MachineCatalogUpgradeDetail) SetScheduleStatus(v VdaUpgradeScheduleStatus)`

SetScheduleStatus sets ScheduleStatus field to given value.

### HasScheduleStatus

`func (o *MachineCatalogUpgradeDetail) HasScheduleStatus() bool`

HasScheduleStatus returns a boolean if a field has been set.

### GetLastStateChangeTimeUtc

`func (o *MachineCatalogUpgradeDetail) GetLastStateChangeTimeUtc() time.Time`

GetLastStateChangeTimeUtc returns the LastStateChangeTimeUtc field if non-nil, zero value otherwise.

### GetLastStateChangeTimeUtcOk

`func (o *MachineCatalogUpgradeDetail) GetLastStateChangeTimeUtcOk() (*time.Time, bool)`

GetLastStateChangeTimeUtcOk returns a tuple with the LastStateChangeTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateChangeTimeUtc

`func (o *MachineCatalogUpgradeDetail) SetLastStateChangeTimeUtc(v time.Time)`

SetLastStateChangeTimeUtc sets LastStateChangeTimeUtc field to given value.

### HasLastStateChangeTimeUtc

`func (o *MachineCatalogUpgradeDetail) HasLastStateChangeTimeUtc() bool`

HasLastStateChangeTimeUtc returns a boolean if a field has been set.

### GetTotalMachines

`func (o *MachineCatalogUpgradeDetail) GetTotalMachines() int32`

GetTotalMachines returns the TotalMachines field if non-nil, zero value otherwise.

### GetTotalMachinesOk

`func (o *MachineCatalogUpgradeDetail) GetTotalMachinesOk() (*int32, bool)`

GetTotalMachinesOk returns a tuple with the TotalMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachines

`func (o *MachineCatalogUpgradeDetail) SetTotalMachines(v int32)`

SetTotalMachines sets TotalMachines field to given value.

### HasTotalMachines

`func (o *MachineCatalogUpgradeDetail) HasTotalMachines() bool`

HasTotalMachines returns a boolean if a field has been set.

### GetSuccessCount

`func (o *MachineCatalogUpgradeDetail) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *MachineCatalogUpgradeDetail) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *MachineCatalogUpgradeDetail) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *MachineCatalogUpgradeDetail) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetValidationFailureCount

`func (o *MachineCatalogUpgradeDetail) GetValidationFailureCount() int32`

GetValidationFailureCount returns the ValidationFailureCount field if non-nil, zero value otherwise.

### GetValidationFailureCountOk

`func (o *MachineCatalogUpgradeDetail) GetValidationFailureCountOk() (*int32, bool)`

GetValidationFailureCountOk returns a tuple with the ValidationFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFailureCount

`func (o *MachineCatalogUpgradeDetail) SetValidationFailureCount(v int32)`

SetValidationFailureCount sets ValidationFailureCount field to given value.

### HasValidationFailureCount

`func (o *MachineCatalogUpgradeDetail) HasValidationFailureCount() bool`

HasValidationFailureCount returns a boolean if a field has been set.

### GetInProgressCount

`func (o *MachineCatalogUpgradeDetail) GetInProgressCount() int32`

GetInProgressCount returns the InProgressCount field if non-nil, zero value otherwise.

### GetInProgressCountOk

`func (o *MachineCatalogUpgradeDetail) GetInProgressCountOk() (*int32, bool)`

GetInProgressCountOk returns a tuple with the InProgressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgressCount

`func (o *MachineCatalogUpgradeDetail) SetInProgressCount(v int32)`

SetInProgressCount sets InProgressCount field to given value.

### HasInProgressCount

`func (o *MachineCatalogUpgradeDetail) HasInProgressCount() bool`

HasInProgressCount returns a boolean if a field has been set.

### GetUpgradeFailureCount

`func (o *MachineCatalogUpgradeDetail) GetUpgradeFailureCount() int32`

GetUpgradeFailureCount returns the UpgradeFailureCount field if non-nil, zero value otherwise.

### GetUpgradeFailureCountOk

`func (o *MachineCatalogUpgradeDetail) GetUpgradeFailureCountOk() (*int32, bool)`

GetUpgradeFailureCountOk returns a tuple with the UpgradeFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeFailureCount

`func (o *MachineCatalogUpgradeDetail) SetUpgradeFailureCount(v int32)`

SetUpgradeFailureCount sets UpgradeFailureCount field to given value.

### HasUpgradeFailureCount

`func (o *MachineCatalogUpgradeDetail) HasUpgradeFailureCount() bool`

HasUpgradeFailureCount returns a boolean if a field has been set.

### GetScheduledTimeUtc

`func (o *MachineCatalogUpgradeDetail) GetScheduledTimeUtc() time.Time`

GetScheduledTimeUtc returns the ScheduledTimeUtc field if non-nil, zero value otherwise.

### GetScheduledTimeUtcOk

`func (o *MachineCatalogUpgradeDetail) GetScheduledTimeUtcOk() (*time.Time, bool)`

GetScheduledTimeUtcOk returns a tuple with the ScheduledTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTimeUtc

`func (o *MachineCatalogUpgradeDetail) SetScheduledTimeUtc(v time.Time)`

SetScheduledTimeUtc sets ScheduledTimeUtc field to given value.

### HasScheduledTimeUtc

`func (o *MachineCatalogUpgradeDetail) HasScheduledTimeUtc() bool`

HasScheduledTimeUtc returns a boolean if a field has been set.

### GetDurationInHours

`func (o *MachineCatalogUpgradeDetail) GetDurationInHours() int32`

GetDurationInHours returns the DurationInHours field if non-nil, zero value otherwise.

### GetDurationInHoursOk

`func (o *MachineCatalogUpgradeDetail) GetDurationInHoursOk() (*int32, bool)`

GetDurationInHoursOk returns a tuple with the DurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInHours

`func (o *MachineCatalogUpgradeDetail) SetDurationInHours(v int32)`

SetDurationInHours sets DurationInHours field to given value.

### HasDurationInHours

`func (o *MachineCatalogUpgradeDetail) HasDurationInHours() bool`

HasDurationInHours returns a boolean if a field has been set.

### GetTargetPackageVersion

`func (o *MachineCatalogUpgradeDetail) GetTargetPackageVersion() string`

GetTargetPackageVersion returns the TargetPackageVersion field if non-nil, zero value otherwise.

### GetTargetPackageVersionOk

`func (o *MachineCatalogUpgradeDetail) GetTargetPackageVersionOk() (*string, bool)`

GetTargetPackageVersionOk returns a tuple with the TargetPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPackageVersion

`func (o *MachineCatalogUpgradeDetail) SetTargetPackageVersion(v string)`

SetTargetPackageVersion sets TargetPackageVersion field to given value.

### HasTargetPackageVersion

`func (o *MachineCatalogUpgradeDetail) HasTargetPackageVersion() bool`

HasTargetPackageVersion returns a boolean if a field has been set.

### SetTargetPackageVersionNil

`func (o *MachineCatalogUpgradeDetail) SetTargetPackageVersionNil(b bool)`

 SetTargetPackageVersionNil sets the value for TargetPackageVersion to be an explicit nil

### UnsetTargetPackageVersion
`func (o *MachineCatalogUpgradeDetail) UnsetTargetPackageVersion()`

UnsetTargetPackageVersion ensures that no value is present for TargetPackageVersion, not even an explicit nil
### GetCancelledUpgradeCount

`func (o *MachineCatalogUpgradeDetail) GetCancelledUpgradeCount() int32`

GetCancelledUpgradeCount returns the CancelledUpgradeCount field if non-nil, zero value otherwise.

### GetCancelledUpgradeCountOk

`func (o *MachineCatalogUpgradeDetail) GetCancelledUpgradeCountOk() (*int32, bool)`

GetCancelledUpgradeCountOk returns a tuple with the CancelledUpgradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledUpgradeCount

`func (o *MachineCatalogUpgradeDetail) SetCancelledUpgradeCount(v int32)`

SetCancelledUpgradeCount sets CancelledUpgradeCount field to given value.

### HasCancelledUpgradeCount

`func (o *MachineCatalogUpgradeDetail) HasCancelledUpgradeCount() bool`

HasCancelledUpgradeCount returns a boolean if a field has been set.

### GetWaitingToUpgradeCount

`func (o *MachineCatalogUpgradeDetail) GetWaitingToUpgradeCount() int32`

GetWaitingToUpgradeCount returns the WaitingToUpgradeCount field if non-nil, zero value otherwise.

### GetWaitingToUpgradeCountOk

`func (o *MachineCatalogUpgradeDetail) GetWaitingToUpgradeCountOk() (*int32, bool)`

GetWaitingToUpgradeCountOk returns a tuple with the WaitingToUpgradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingToUpgradeCount

`func (o *MachineCatalogUpgradeDetail) SetWaitingToUpgradeCount(v int32)`

SetWaitingToUpgradeCount sets WaitingToUpgradeCount field to given value.

### HasWaitingToUpgradeCount

`func (o *MachineCatalogUpgradeDetail) HasWaitingToUpgradeCount() bool`

HasWaitingToUpgradeCount returns a boolean if a field has been set.

### GetAvailableForUpgradeCount

`func (o *MachineCatalogUpgradeDetail) GetAvailableForUpgradeCount() int32`

GetAvailableForUpgradeCount returns the AvailableForUpgradeCount field if non-nil, zero value otherwise.

### GetAvailableForUpgradeCountOk

`func (o *MachineCatalogUpgradeDetail) GetAvailableForUpgradeCountOk() (*int32, bool)`

GetAvailableForUpgradeCountOk returns a tuple with the AvailableForUpgradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForUpgradeCount

`func (o *MachineCatalogUpgradeDetail) SetAvailableForUpgradeCount(v int32)`

SetAvailableForUpgradeCount sets AvailableForUpgradeCount field to given value.

### HasAvailableForUpgradeCount

`func (o *MachineCatalogUpgradeDetail) HasAvailableForUpgradeCount() bool`

HasAvailableForUpgradeCount returns a boolean if a field has been set.

### SetAvailableForUpgradeCountNil

`func (o *MachineCatalogUpgradeDetail) SetAvailableForUpgradeCountNil(b bool)`

 SetAvailableForUpgradeCountNil sets the value for AvailableForUpgradeCount to be an explicit nil

### UnsetAvailableForUpgradeCount
`func (o *MachineCatalogUpgradeDetail) UnsetAvailableForUpgradeCount()`

UnsetAvailableForUpgradeCount ensures that no value is present for AvailableForUpgradeCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


