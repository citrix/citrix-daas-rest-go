# RebootMachinesRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebootDuration** | Pointer to **NullableInt32** | Approximate maximum duration over which the reboot cycle runs, in minutes. | [optional] [default to -1]
**WarningDuration** | Pointer to **NullableInt32** | Time in minutes prior to a machine reboot at which a warning message is displayed in all user sessions on that machine. | [optional] [default to 0]
**WarningTitle** | Pointer to **NullableString** | The window title used when showing the warning message in user sessions on a machine scheduled for reboot. | [optional] 
**WarningMessage** | Pointer to **NullableString** | Warning message displayed in user sessions on a machine scheduled for a reboot.  The optional pattern &#39;%m%&#39; is replaced by the number of minutes until the reboot. | [optional] 
**WarningRepeatInterval** | Pointer to **NullableInt32** | Number of minutes to wait before showing the reboot warning message again. | [optional] [default to 0]

## Methods

### NewRebootMachinesRequestModel

`func NewRebootMachinesRequestModel() *RebootMachinesRequestModel`

NewRebootMachinesRequestModel instantiates a new RebootMachinesRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootMachinesRequestModelWithDefaults

`func NewRebootMachinesRequestModelWithDefaults() *RebootMachinesRequestModel`

NewRebootMachinesRequestModelWithDefaults instantiates a new RebootMachinesRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRebootDuration

`func (o *RebootMachinesRequestModel) GetRebootDuration() int32`

GetRebootDuration returns the RebootDuration field if non-nil, zero value otherwise.

### GetRebootDurationOk

`func (o *RebootMachinesRequestModel) GetRebootDurationOk() (*int32, bool)`

GetRebootDurationOk returns a tuple with the RebootDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDuration

`func (o *RebootMachinesRequestModel) SetRebootDuration(v int32)`

SetRebootDuration sets RebootDuration field to given value.

### HasRebootDuration

`func (o *RebootMachinesRequestModel) HasRebootDuration() bool`

HasRebootDuration returns a boolean if a field has been set.

### SetRebootDurationNil

`func (o *RebootMachinesRequestModel) SetRebootDurationNil(b bool)`

 SetRebootDurationNil sets the value for RebootDuration to be an explicit nil

### UnsetRebootDuration
`func (o *RebootMachinesRequestModel) UnsetRebootDuration()`

UnsetRebootDuration ensures that no value is present for RebootDuration, not even an explicit nil
### GetWarningDuration

`func (o *RebootMachinesRequestModel) GetWarningDuration() int32`

GetWarningDuration returns the WarningDuration field if non-nil, zero value otherwise.

### GetWarningDurationOk

`func (o *RebootMachinesRequestModel) GetWarningDurationOk() (*int32, bool)`

GetWarningDurationOk returns a tuple with the WarningDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningDuration

`func (o *RebootMachinesRequestModel) SetWarningDuration(v int32)`

SetWarningDuration sets WarningDuration field to given value.

### HasWarningDuration

`func (o *RebootMachinesRequestModel) HasWarningDuration() bool`

HasWarningDuration returns a boolean if a field has been set.

### SetWarningDurationNil

`func (o *RebootMachinesRequestModel) SetWarningDurationNil(b bool)`

 SetWarningDurationNil sets the value for WarningDuration to be an explicit nil

### UnsetWarningDuration
`func (o *RebootMachinesRequestModel) UnsetWarningDuration()`

UnsetWarningDuration ensures that no value is present for WarningDuration, not even an explicit nil
### GetWarningTitle

`func (o *RebootMachinesRequestModel) GetWarningTitle() string`

GetWarningTitle returns the WarningTitle field if non-nil, zero value otherwise.

### GetWarningTitleOk

`func (o *RebootMachinesRequestModel) GetWarningTitleOk() (*string, bool)`

GetWarningTitleOk returns a tuple with the WarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningTitle

`func (o *RebootMachinesRequestModel) SetWarningTitle(v string)`

SetWarningTitle sets WarningTitle field to given value.

### HasWarningTitle

`func (o *RebootMachinesRequestModel) HasWarningTitle() bool`

HasWarningTitle returns a boolean if a field has been set.

### SetWarningTitleNil

`func (o *RebootMachinesRequestModel) SetWarningTitleNil(b bool)`

 SetWarningTitleNil sets the value for WarningTitle to be an explicit nil

### UnsetWarningTitle
`func (o *RebootMachinesRequestModel) UnsetWarningTitle()`

UnsetWarningTitle ensures that no value is present for WarningTitle, not even an explicit nil
### GetWarningMessage

`func (o *RebootMachinesRequestModel) GetWarningMessage() string`

GetWarningMessage returns the WarningMessage field if non-nil, zero value otherwise.

### GetWarningMessageOk

`func (o *RebootMachinesRequestModel) GetWarningMessageOk() (*string, bool)`

GetWarningMessageOk returns a tuple with the WarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessage

`func (o *RebootMachinesRequestModel) SetWarningMessage(v string)`

SetWarningMessage sets WarningMessage field to given value.

### HasWarningMessage

`func (o *RebootMachinesRequestModel) HasWarningMessage() bool`

HasWarningMessage returns a boolean if a field has been set.

### SetWarningMessageNil

`func (o *RebootMachinesRequestModel) SetWarningMessageNil(b bool)`

 SetWarningMessageNil sets the value for WarningMessage to be an explicit nil

### UnsetWarningMessage
`func (o *RebootMachinesRequestModel) UnsetWarningMessage()`

UnsetWarningMessage ensures that no value is present for WarningMessage, not even an explicit nil
### GetWarningRepeatInterval

`func (o *RebootMachinesRequestModel) GetWarningRepeatInterval() int32`

GetWarningRepeatInterval returns the WarningRepeatInterval field if non-nil, zero value otherwise.

### GetWarningRepeatIntervalOk

`func (o *RebootMachinesRequestModel) GetWarningRepeatIntervalOk() (*int32, bool)`

GetWarningRepeatIntervalOk returns a tuple with the WarningRepeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningRepeatInterval

`func (o *RebootMachinesRequestModel) SetWarningRepeatInterval(v int32)`

SetWarningRepeatInterval sets WarningRepeatInterval field to given value.

### HasWarningRepeatInterval

`func (o *RebootMachinesRequestModel) HasWarningRepeatInterval() bool`

HasWarningRepeatInterval returns a boolean if a field has been set.

### SetWarningRepeatIntervalNil

`func (o *RebootMachinesRequestModel) SetWarningRepeatIntervalNil(b bool)`

 SetWarningRepeatIntervalNil sets the value for WarningRepeatInterval to be an explicit nil

### UnsetWarningRepeatInterval
`func (o *RebootMachinesRequestModel) UnsetWarningRepeatInterval()`

UnsetWarningRepeatInterval ensures that no value is present for WarningRepeatInterval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


