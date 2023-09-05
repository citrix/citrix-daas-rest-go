# MachineCatalogDetailResponseModelAllOfUpgradeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleStatus** | Pointer to [**VdaUpgradeScheduleStatus**](VdaUpgradeScheduleStatus.md) |  | [optional] 
**LastStateChangeTimeUtc** | Pointer to **time.Time** | UTC time when this upgrade status object changed status for the last time. | [optional] 
**TotalMachines** | Pointer to **int32** | Count of machines with an upgrade schedule in this catalog. This does not always equal to the machine count of the catalog, as there might be machines which joined the catalog after a catalog level upgrade is schdeuled, so do not have an upgrade status. | [optional] 
**SuccessCount** | Pointer to **int32** | Count of machines whose last upgrade is in success state. | [optional] 
**ValidationFailureCount** | Pointer to **int32** | Count of machines whose last upgrade failed during package validation. | [optional] 
**InProgressCount** | Pointer to **int32** | Count of machines whose upgrade is in progress. | [optional] 
**UpgradeFailureCount** | Pointer to **int32** | Count of machines whose last upgrade failed during package installtion. | [optional] 
**ScheduledTimeUtc** | Pointer to **time.Time** | UTC time when this VDA upgrade was scheduled to start. | [optional] 
**DurationInHours** | Pointer to **int32** | Timeout duration in hours, of the current VDA upgrade schdeule. | [optional] 
**TargetPackageVersion** | Pointer to **string** | Target package version of the current VDA upgrade schdeule. | [optional] 
**CancelledUpgradeCount** | Pointer to **int32** | Count of machines whose last upgrade canceled during package installtion. | [optional] 
**WaitingToUpgradeCount** | Pointer to **int32** | Count of machines who is waiting to upgrade. | [optional] 

## Methods

### NewMachineCatalogDetailResponseModelAllOfUpgradeDetail

`func NewMachineCatalogDetailResponseModelAllOfUpgradeDetail() *MachineCatalogDetailResponseModelAllOfUpgradeDetail`

NewMachineCatalogDetailResponseModelAllOfUpgradeDetail instantiates a new MachineCatalogDetailResponseModelAllOfUpgradeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogDetailResponseModelAllOfUpgradeDetailWithDefaults

`func NewMachineCatalogDetailResponseModelAllOfUpgradeDetailWithDefaults() *MachineCatalogDetailResponseModelAllOfUpgradeDetail`

NewMachineCatalogDetailResponseModelAllOfUpgradeDetailWithDefaults instantiates a new MachineCatalogDetailResponseModelAllOfUpgradeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleStatus

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetScheduleStatus() VdaUpgradeScheduleStatus`

GetScheduleStatus returns the ScheduleStatus field if non-nil, zero value otherwise.

### GetScheduleStatusOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetScheduleStatusOk() (*VdaUpgradeScheduleStatus, bool)`

GetScheduleStatusOk returns a tuple with the ScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStatus

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetScheduleStatus(v VdaUpgradeScheduleStatus)`

SetScheduleStatus sets ScheduleStatus field to given value.

### HasScheduleStatus

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasScheduleStatus() bool`

HasScheduleStatus returns a boolean if a field has been set.

### GetLastStateChangeTimeUtc

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetLastStateChangeTimeUtc() time.Time`

GetLastStateChangeTimeUtc returns the LastStateChangeTimeUtc field if non-nil, zero value otherwise.

### GetLastStateChangeTimeUtcOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetLastStateChangeTimeUtcOk() (*time.Time, bool)`

GetLastStateChangeTimeUtcOk returns a tuple with the LastStateChangeTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateChangeTimeUtc

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetLastStateChangeTimeUtc(v time.Time)`

SetLastStateChangeTimeUtc sets LastStateChangeTimeUtc field to given value.

### HasLastStateChangeTimeUtc

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasLastStateChangeTimeUtc() bool`

HasLastStateChangeTimeUtc returns a boolean if a field has been set.

### GetTotalMachines

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetTotalMachines() int32`

GetTotalMachines returns the TotalMachines field if non-nil, zero value otherwise.

### GetTotalMachinesOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetTotalMachinesOk() (*int32, bool)`

GetTotalMachinesOk returns a tuple with the TotalMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachines

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetTotalMachines(v int32)`

SetTotalMachines sets TotalMachines field to given value.

### HasTotalMachines

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasTotalMachines() bool`

HasTotalMachines returns a boolean if a field has been set.

### GetSuccessCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetValidationFailureCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetValidationFailureCount() int32`

GetValidationFailureCount returns the ValidationFailureCount field if non-nil, zero value otherwise.

### GetValidationFailureCountOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetValidationFailureCountOk() (*int32, bool)`

GetValidationFailureCountOk returns a tuple with the ValidationFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFailureCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetValidationFailureCount(v int32)`

SetValidationFailureCount sets ValidationFailureCount field to given value.

### HasValidationFailureCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasValidationFailureCount() bool`

HasValidationFailureCount returns a boolean if a field has been set.

### GetInProgressCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetInProgressCount() int32`

GetInProgressCount returns the InProgressCount field if non-nil, zero value otherwise.

### GetInProgressCountOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetInProgressCountOk() (*int32, bool)`

GetInProgressCountOk returns a tuple with the InProgressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgressCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetInProgressCount(v int32)`

SetInProgressCount sets InProgressCount field to given value.

### HasInProgressCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasInProgressCount() bool`

HasInProgressCount returns a boolean if a field has been set.

### GetUpgradeFailureCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetUpgradeFailureCount() int32`

GetUpgradeFailureCount returns the UpgradeFailureCount field if non-nil, zero value otherwise.

### GetUpgradeFailureCountOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetUpgradeFailureCountOk() (*int32, bool)`

GetUpgradeFailureCountOk returns a tuple with the UpgradeFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeFailureCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetUpgradeFailureCount(v int32)`

SetUpgradeFailureCount sets UpgradeFailureCount field to given value.

### HasUpgradeFailureCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasUpgradeFailureCount() bool`

HasUpgradeFailureCount returns a boolean if a field has been set.

### GetScheduledTimeUtc

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetScheduledTimeUtc() time.Time`

GetScheduledTimeUtc returns the ScheduledTimeUtc field if non-nil, zero value otherwise.

### GetScheduledTimeUtcOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetScheduledTimeUtcOk() (*time.Time, bool)`

GetScheduledTimeUtcOk returns a tuple with the ScheduledTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTimeUtc

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetScheduledTimeUtc(v time.Time)`

SetScheduledTimeUtc sets ScheduledTimeUtc field to given value.

### HasScheduledTimeUtc

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasScheduledTimeUtc() bool`

HasScheduledTimeUtc returns a boolean if a field has been set.

### GetDurationInHours

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetDurationInHours() int32`

GetDurationInHours returns the DurationInHours field if non-nil, zero value otherwise.

### GetDurationInHoursOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetDurationInHoursOk() (*int32, bool)`

GetDurationInHoursOk returns a tuple with the DurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInHours

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetDurationInHours(v int32)`

SetDurationInHours sets DurationInHours field to given value.

### HasDurationInHours

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasDurationInHours() bool`

HasDurationInHours returns a boolean if a field has been set.

### GetTargetPackageVersion

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetTargetPackageVersion() string`

GetTargetPackageVersion returns the TargetPackageVersion field if non-nil, zero value otherwise.

### GetTargetPackageVersionOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetTargetPackageVersionOk() (*string, bool)`

GetTargetPackageVersionOk returns a tuple with the TargetPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPackageVersion

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetTargetPackageVersion(v string)`

SetTargetPackageVersion sets TargetPackageVersion field to given value.

### HasTargetPackageVersion

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasTargetPackageVersion() bool`

HasTargetPackageVersion returns a boolean if a field has been set.

### GetCancelledUpgradeCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetCancelledUpgradeCount() int32`

GetCancelledUpgradeCount returns the CancelledUpgradeCount field if non-nil, zero value otherwise.

### GetCancelledUpgradeCountOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetCancelledUpgradeCountOk() (*int32, bool)`

GetCancelledUpgradeCountOk returns a tuple with the CancelledUpgradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledUpgradeCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetCancelledUpgradeCount(v int32)`

SetCancelledUpgradeCount sets CancelledUpgradeCount field to given value.

### HasCancelledUpgradeCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasCancelledUpgradeCount() bool`

HasCancelledUpgradeCount returns a boolean if a field has been set.

### GetWaitingToUpgradeCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetWaitingToUpgradeCount() int32`

GetWaitingToUpgradeCount returns the WaitingToUpgradeCount field if non-nil, zero value otherwise.

### GetWaitingToUpgradeCountOk

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) GetWaitingToUpgradeCountOk() (*int32, bool)`

GetWaitingToUpgradeCountOk returns a tuple with the WaitingToUpgradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingToUpgradeCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) SetWaitingToUpgradeCount(v int32)`

SetWaitingToUpgradeCount sets WaitingToUpgradeCount field to given value.

### HasWaitingToUpgradeCount

`func (o *MachineCatalogDetailResponseModelAllOfUpgradeDetail) HasWaitingToUpgradeCount() bool`

HasWaitingToUpgradeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


