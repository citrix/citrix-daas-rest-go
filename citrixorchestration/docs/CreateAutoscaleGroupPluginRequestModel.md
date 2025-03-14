# CreateAutoscaleGroupPluginRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the group plugin. | 
**Type** | [**AutoscalePluginType**](AutoscalePluginType.md) |  | 
**Priority** | Pointer to **NullableInt32** | The priority of this group plugin with respect to other group plugins associated with the same desktop group. Priority 1 is the highest priority. | [optional] 
**Enabled** | Pointer to **NullableBool** | Specifies whether the group plugin is enabled. | [optional] 
**Description** | Pointer to **NullableString** | Optional description of the group plugin. | [optional] 
**AutoApprove** | Pointer to **NullableBool** | If True, indicates that any schedule override produced by the plugin is automatically accepted as the schedule for the target date. Otherwise the schedule is only accepted if the target date is on or before the date specified in ApprovedUtil. If AutoApprove is False and the ApprovedUtil date is not set, or is in the past, then any schedule override produced is discarded without being used. | [optional] 
**Dates** | Pointer to **[]string** | Date range for the autoscale holiday plugin template. | [optional] 

## Methods

### NewCreateAutoscaleGroupPluginRequestModel

`func NewCreateAutoscaleGroupPluginRequestModel(name string, type_ AutoscalePluginType, ) *CreateAutoscaleGroupPluginRequestModel`

NewCreateAutoscaleGroupPluginRequestModel instantiates a new CreateAutoscaleGroupPluginRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoscaleGroupPluginRequestModelWithDefaults

`func NewCreateAutoscaleGroupPluginRequestModelWithDefaults() *CreateAutoscaleGroupPluginRequestModel`

NewCreateAutoscaleGroupPluginRequestModelWithDefaults instantiates a new CreateAutoscaleGroupPluginRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAutoscaleGroupPluginRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAutoscaleGroupPluginRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAutoscaleGroupPluginRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateAutoscaleGroupPluginRequestModel) GetType() AutoscalePluginType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAutoscaleGroupPluginRequestModel) GetTypeOk() (*AutoscalePluginType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAutoscaleGroupPluginRequestModel) SetType(v AutoscalePluginType)`

SetType sets Type field to given value.


### GetPriority

`func (o *CreateAutoscaleGroupPluginRequestModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateAutoscaleGroupPluginRequestModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateAutoscaleGroupPluginRequestModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateAutoscaleGroupPluginRequestModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CreateAutoscaleGroupPluginRequestModel) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CreateAutoscaleGroupPluginRequestModel) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetEnabled

`func (o *CreateAutoscaleGroupPluginRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAutoscaleGroupPluginRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAutoscaleGroupPluginRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateAutoscaleGroupPluginRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *CreateAutoscaleGroupPluginRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *CreateAutoscaleGroupPluginRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetDescription

`func (o *CreateAutoscaleGroupPluginRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAutoscaleGroupPluginRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAutoscaleGroupPluginRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAutoscaleGroupPluginRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateAutoscaleGroupPluginRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateAutoscaleGroupPluginRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAutoApprove

`func (o *CreateAutoscaleGroupPluginRequestModel) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *CreateAutoscaleGroupPluginRequestModel) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *CreateAutoscaleGroupPluginRequestModel) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.

### HasAutoApprove

`func (o *CreateAutoscaleGroupPluginRequestModel) HasAutoApprove() bool`

HasAutoApprove returns a boolean if a field has been set.

### SetAutoApproveNil

`func (o *CreateAutoscaleGroupPluginRequestModel) SetAutoApproveNil(b bool)`

 SetAutoApproveNil sets the value for AutoApprove to be an explicit nil

### UnsetAutoApprove
`func (o *CreateAutoscaleGroupPluginRequestModel) UnsetAutoApprove()`

UnsetAutoApprove ensures that no value is present for AutoApprove, not even an explicit nil
### GetDates

`func (o *CreateAutoscaleGroupPluginRequestModel) GetDates() []string`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *CreateAutoscaleGroupPluginRequestModel) GetDatesOk() (*[]string, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *CreateAutoscaleGroupPluginRequestModel) SetDates(v []string)`

SetDates sets Dates field to given value.

### HasDates

`func (o *CreateAutoscaleGroupPluginRequestModel) HasDates() bool`

HasDates returns a boolean if a field has been set.

### SetDatesNil

`func (o *CreateAutoscaleGroupPluginRequestModel) SetDatesNil(b bool)`

 SetDatesNil sets the value for Dates to be an explicit nil

### UnsetDates
`func (o *CreateAutoscaleGroupPluginRequestModel) UnsetDates()`

UnsetDates ensures that no value is present for Dates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


