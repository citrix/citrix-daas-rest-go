# RebootCycleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebootDurationMinutes** | Pointer to **NullableInt32** | Approximate maximum number of minutes over which the reboot cycle runs. | [optional] 
**IgnoreMaintenanceMode** | Pointer to **NullableBool** | Indicates whether the reboot cycle is active for maintained machines. | [optional] 
**RestrictToTag** | Pointer to **NullableString** | If set, the reboot cycle only applies to machines in the delivery group with the specified tag. | [optional] 
**WarningTitle** | Pointer to **NullableString** | The window title used when showing the warning message in user sessions on a machine to reboot. | [optional] 
**WarningMessage** | Pointer to **NullableString** | Warning message displayed in user sessions on a machine to reboot. | [optional] 
**WarningDurationMinutes** | Pointer to **NullableInt32** | Time prior to the start of a machine reboot at which a warning message is displayed to all the users on the machine. | [optional] 
**WarningRepeatIntervalMinutes** | Pointer to **NullableInt32** | Time to wait after the previous reboot warning before displaying the warning message in all user sessions on that machine again. | [optional] 

## Methods

### NewRebootCycleRequestModel

`func NewRebootCycleRequestModel() *RebootCycleRequestModel`

NewRebootCycleRequestModel instantiates a new RebootCycleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootCycleRequestModelWithDefaults

`func NewRebootCycleRequestModelWithDefaults() *RebootCycleRequestModel`

NewRebootCycleRequestModelWithDefaults instantiates a new RebootCycleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRebootDurationMinutes

`func (o *RebootCycleRequestModel) GetRebootDurationMinutes() int32`

GetRebootDurationMinutes returns the RebootDurationMinutes field if non-nil, zero value otherwise.

### GetRebootDurationMinutesOk

`func (o *RebootCycleRequestModel) GetRebootDurationMinutesOk() (*int32, bool)`

GetRebootDurationMinutesOk returns a tuple with the RebootDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDurationMinutes

`func (o *RebootCycleRequestModel) SetRebootDurationMinutes(v int32)`

SetRebootDurationMinutes sets RebootDurationMinutes field to given value.

### HasRebootDurationMinutes

`func (o *RebootCycleRequestModel) HasRebootDurationMinutes() bool`

HasRebootDurationMinutes returns a boolean if a field has been set.

### SetRebootDurationMinutesNil

`func (o *RebootCycleRequestModel) SetRebootDurationMinutesNil(b bool)`

 SetRebootDurationMinutesNil sets the value for RebootDurationMinutes to be an explicit nil

### UnsetRebootDurationMinutes
`func (o *RebootCycleRequestModel) UnsetRebootDurationMinutes()`

UnsetRebootDurationMinutes ensures that no value is present for RebootDurationMinutes, not even an explicit nil
### GetIgnoreMaintenanceMode

`func (o *RebootCycleRequestModel) GetIgnoreMaintenanceMode() bool`

GetIgnoreMaintenanceMode returns the IgnoreMaintenanceMode field if non-nil, zero value otherwise.

### GetIgnoreMaintenanceModeOk

`func (o *RebootCycleRequestModel) GetIgnoreMaintenanceModeOk() (*bool, bool)`

GetIgnoreMaintenanceModeOk returns a tuple with the IgnoreMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMaintenanceMode

`func (o *RebootCycleRequestModel) SetIgnoreMaintenanceMode(v bool)`

SetIgnoreMaintenanceMode sets IgnoreMaintenanceMode field to given value.

### HasIgnoreMaintenanceMode

`func (o *RebootCycleRequestModel) HasIgnoreMaintenanceMode() bool`

HasIgnoreMaintenanceMode returns a boolean if a field has been set.

### SetIgnoreMaintenanceModeNil

`func (o *RebootCycleRequestModel) SetIgnoreMaintenanceModeNil(b bool)`

 SetIgnoreMaintenanceModeNil sets the value for IgnoreMaintenanceMode to be an explicit nil

### UnsetIgnoreMaintenanceMode
`func (o *RebootCycleRequestModel) UnsetIgnoreMaintenanceMode()`

UnsetIgnoreMaintenanceMode ensures that no value is present for IgnoreMaintenanceMode, not even an explicit nil
### GetRestrictToTag

`func (o *RebootCycleRequestModel) GetRestrictToTag() string`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *RebootCycleRequestModel) GetRestrictToTagOk() (*string, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *RebootCycleRequestModel) SetRestrictToTag(v string)`

SetRestrictToTag sets RestrictToTag field to given value.

### HasRestrictToTag

`func (o *RebootCycleRequestModel) HasRestrictToTag() bool`

HasRestrictToTag returns a boolean if a field has been set.

### SetRestrictToTagNil

`func (o *RebootCycleRequestModel) SetRestrictToTagNil(b bool)`

 SetRestrictToTagNil sets the value for RestrictToTag to be an explicit nil

### UnsetRestrictToTag
`func (o *RebootCycleRequestModel) UnsetRestrictToTag()`

UnsetRestrictToTag ensures that no value is present for RestrictToTag, not even an explicit nil
### GetWarningTitle

`func (o *RebootCycleRequestModel) GetWarningTitle() string`

GetWarningTitle returns the WarningTitle field if non-nil, zero value otherwise.

### GetWarningTitleOk

`func (o *RebootCycleRequestModel) GetWarningTitleOk() (*string, bool)`

GetWarningTitleOk returns a tuple with the WarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningTitle

`func (o *RebootCycleRequestModel) SetWarningTitle(v string)`

SetWarningTitle sets WarningTitle field to given value.

### HasWarningTitle

`func (o *RebootCycleRequestModel) HasWarningTitle() bool`

HasWarningTitle returns a boolean if a field has been set.

### SetWarningTitleNil

`func (o *RebootCycleRequestModel) SetWarningTitleNil(b bool)`

 SetWarningTitleNil sets the value for WarningTitle to be an explicit nil

### UnsetWarningTitle
`func (o *RebootCycleRequestModel) UnsetWarningTitle()`

UnsetWarningTitle ensures that no value is present for WarningTitle, not even an explicit nil
### GetWarningMessage

`func (o *RebootCycleRequestModel) GetWarningMessage() string`

GetWarningMessage returns the WarningMessage field if non-nil, zero value otherwise.

### GetWarningMessageOk

`func (o *RebootCycleRequestModel) GetWarningMessageOk() (*string, bool)`

GetWarningMessageOk returns a tuple with the WarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessage

`func (o *RebootCycleRequestModel) SetWarningMessage(v string)`

SetWarningMessage sets WarningMessage field to given value.

### HasWarningMessage

`func (o *RebootCycleRequestModel) HasWarningMessage() bool`

HasWarningMessage returns a boolean if a field has been set.

### SetWarningMessageNil

`func (o *RebootCycleRequestModel) SetWarningMessageNil(b bool)`

 SetWarningMessageNil sets the value for WarningMessage to be an explicit nil

### UnsetWarningMessage
`func (o *RebootCycleRequestModel) UnsetWarningMessage()`

UnsetWarningMessage ensures that no value is present for WarningMessage, not even an explicit nil
### GetWarningDurationMinutes

`func (o *RebootCycleRequestModel) GetWarningDurationMinutes() int32`

GetWarningDurationMinutes returns the WarningDurationMinutes field if non-nil, zero value otherwise.

### GetWarningDurationMinutesOk

`func (o *RebootCycleRequestModel) GetWarningDurationMinutesOk() (*int32, bool)`

GetWarningDurationMinutesOk returns a tuple with the WarningDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningDurationMinutes

`func (o *RebootCycleRequestModel) SetWarningDurationMinutes(v int32)`

SetWarningDurationMinutes sets WarningDurationMinutes field to given value.

### HasWarningDurationMinutes

`func (o *RebootCycleRequestModel) HasWarningDurationMinutes() bool`

HasWarningDurationMinutes returns a boolean if a field has been set.

### SetWarningDurationMinutesNil

`func (o *RebootCycleRequestModel) SetWarningDurationMinutesNil(b bool)`

 SetWarningDurationMinutesNil sets the value for WarningDurationMinutes to be an explicit nil

### UnsetWarningDurationMinutes
`func (o *RebootCycleRequestModel) UnsetWarningDurationMinutes()`

UnsetWarningDurationMinutes ensures that no value is present for WarningDurationMinutes, not even an explicit nil
### GetWarningRepeatIntervalMinutes

`func (o *RebootCycleRequestModel) GetWarningRepeatIntervalMinutes() int32`

GetWarningRepeatIntervalMinutes returns the WarningRepeatIntervalMinutes field if non-nil, zero value otherwise.

### GetWarningRepeatIntervalMinutesOk

`func (o *RebootCycleRequestModel) GetWarningRepeatIntervalMinutesOk() (*int32, bool)`

GetWarningRepeatIntervalMinutesOk returns a tuple with the WarningRepeatIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningRepeatIntervalMinutes

`func (o *RebootCycleRequestModel) SetWarningRepeatIntervalMinutes(v int32)`

SetWarningRepeatIntervalMinutes sets WarningRepeatIntervalMinutes field to given value.

### HasWarningRepeatIntervalMinutes

`func (o *RebootCycleRequestModel) HasWarningRepeatIntervalMinutes() bool`

HasWarningRepeatIntervalMinutes returns a boolean if a field has been set.

### SetWarningRepeatIntervalMinutesNil

`func (o *RebootCycleRequestModel) SetWarningRepeatIntervalMinutesNil(b bool)`

 SetWarningRepeatIntervalMinutesNil sets the value for WarningRepeatIntervalMinutes to be an explicit nil

### UnsetWarningRepeatIntervalMinutes
`func (o *RebootCycleRequestModel) UnsetWarningRepeatIntervalMinutes()`

UnsetWarningRepeatIntervalMinutes ensures that no value is present for WarningRepeatIntervalMinutes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


