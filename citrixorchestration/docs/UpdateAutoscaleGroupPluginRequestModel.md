# UpdateAutoscaleGroupPluginRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the group plugin. | [optional] 
**Priority** | Pointer to **NullableInt32** | The priority of this group plugin with respect to other group plugins associated with the same desktop group. Priority 1 is the highest priority. | [optional] 
**Enabled** | Pointer to **NullableBool** | Specifies whether the group plugin is enabled. | [optional] 
**Description** | Pointer to **NullableString** | Optional description of the group plugin. | [optional] 
**AutoApprove** | Pointer to **NullableBool** | If True, indicates that any schedule override produced by the plugin is automatically accepted as the schedule for the target date. Otherwise the schedule is only accepted if the target date is on or before the date specified in ApprovedUtil. If AutoApprove is False and the ApprovedUtil date is not set, or is in the past, then any schedule override produced is discarded without being used. | [optional] 
**Dates** | Pointer to **[]string** | Date range for the autoscale holiday plugin template. | [optional] 
**ApprovedUntil** | Pointer to **NullableString** | If AutoApprove is False, specifies the date until which any schedule override produced by the plugin is automatically accepted as the schedule for the target date. | [optional] 

## Methods

### NewUpdateAutoscaleGroupPluginRequestModel

`func NewUpdateAutoscaleGroupPluginRequestModel() *UpdateAutoscaleGroupPluginRequestModel`

NewUpdateAutoscaleGroupPluginRequestModel instantiates a new UpdateAutoscaleGroupPluginRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAutoscaleGroupPluginRequestModelWithDefaults

`func NewUpdateAutoscaleGroupPluginRequestModelWithDefaults() *UpdateAutoscaleGroupPluginRequestModel`

NewUpdateAutoscaleGroupPluginRequestModelWithDefaults instantiates a new UpdateAutoscaleGroupPluginRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAutoscaleGroupPluginRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPriority

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateAutoscaleGroupPluginRequestModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetEnabled

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAutoscaleGroupPluginRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetDescription

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAutoscaleGroupPluginRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAutoApprove

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.

### HasAutoApprove

`func (o *UpdateAutoscaleGroupPluginRequestModel) HasAutoApprove() bool`

HasAutoApprove returns a boolean if a field has been set.

### SetAutoApproveNil

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetAutoApproveNil(b bool)`

 SetAutoApproveNil sets the value for AutoApprove to be an explicit nil

### UnsetAutoApprove
`func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetAutoApprove()`

UnsetAutoApprove ensures that no value is present for AutoApprove, not even an explicit nil
### GetDates

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetDates() []string`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetDatesOk() (*[]string, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetDates(v []string)`

SetDates sets Dates field to given value.

### HasDates

`func (o *UpdateAutoscaleGroupPluginRequestModel) HasDates() bool`

HasDates returns a boolean if a field has been set.

### SetDatesNil

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetDatesNil(b bool)`

 SetDatesNil sets the value for Dates to be an explicit nil

### UnsetDates
`func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetDates()`

UnsetDates ensures that no value is present for Dates, not even an explicit nil
### GetApprovedUntil

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetApprovedUntil() string`

GetApprovedUntil returns the ApprovedUntil field if non-nil, zero value otherwise.

### GetApprovedUntilOk

`func (o *UpdateAutoscaleGroupPluginRequestModel) GetApprovedUntilOk() (*string, bool)`

GetApprovedUntilOk returns a tuple with the ApprovedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedUntil

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetApprovedUntil(v string)`

SetApprovedUntil sets ApprovedUntil field to given value.

### HasApprovedUntil

`func (o *UpdateAutoscaleGroupPluginRequestModel) HasApprovedUntil() bool`

HasApprovedUntil returns a boolean if a field has been set.

### SetApprovedUntilNil

`func (o *UpdateAutoscaleGroupPluginRequestModel) SetApprovedUntilNil(b bool)`

 SetApprovedUntilNil sets the value for ApprovedUntil to be an explicit nil

### UnsetApprovedUntil
`func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetApprovedUntil()`

UnsetApprovedUntil ensures that no value is present for ApprovedUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


