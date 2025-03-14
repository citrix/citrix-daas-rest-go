# AutoscaleGroupPluginModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** | The unique ID of the desktop group associated with this group plugin. | [optional] 
**Name** | Pointer to **NullableString** | The name of the group plugin. | [optional] 
**Priority** | Pointer to **int32** | The priority of this group plugin with respect to other group plugins associated with the same desktop group. Priority 1 is the highest priority. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether the group plugin is enabled. | [optional] 
**Description** | Pointer to **NullableString** | Optional description of the group plugin. | [optional] 
**DesktopGroupUid** | Pointer to **int32** | The unique ID of the desktop group associated with this group plugin. | [optional] 
**Type** | Pointer to [**AutoscalePluginType**](AutoscalePluginType.md) |  | [optional] 
**AutoApprove** | Pointer to **bool** | If True, indicates that any schedule override produced by the plugin is automatically accepted as the schedule for the target date. Otherwise the schedule is only accepted if the target date is on or before the date specified in ApprovedUtil. If AutoApprove is False and the ApprovedUtil date is not set, or is in the past, then any schedule override produced is discarded without being used. | [optional] 
**ApprovedUntil** | Pointer to **NullableString** | If AutoApprove is False, specifies the date until which any schedule override produced by the plugin is automatically accepted as the schedule for the target date. See AutoApprove for more information. | [optional] 
**Dates** | Pointer to **[]string** | Date range for the autoscale holiday plugin template. | [optional] 

## Methods

### NewAutoscaleGroupPluginModel

`func NewAutoscaleGroupPluginModel() *AutoscaleGroupPluginModel`

NewAutoscaleGroupPluginModel instantiates a new AutoscaleGroupPluginModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscaleGroupPluginModelWithDefaults

`func NewAutoscaleGroupPluginModelWithDefaults() *AutoscaleGroupPluginModel`

NewAutoscaleGroupPluginModelWithDefaults instantiates a new AutoscaleGroupPluginModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *AutoscaleGroupPluginModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AutoscaleGroupPluginModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AutoscaleGroupPluginModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AutoscaleGroupPluginModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *AutoscaleGroupPluginModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoscaleGroupPluginModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoscaleGroupPluginModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoscaleGroupPluginModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoscaleGroupPluginModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoscaleGroupPluginModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPriority

`func (o *AutoscaleGroupPluginModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AutoscaleGroupPluginModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AutoscaleGroupPluginModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AutoscaleGroupPluginModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEnabled

`func (o *AutoscaleGroupPluginModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoscaleGroupPluginModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoscaleGroupPluginModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoscaleGroupPluginModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *AutoscaleGroupPluginModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoscaleGroupPluginModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoscaleGroupPluginModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoscaleGroupPluginModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoscaleGroupPluginModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoscaleGroupPluginModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesktopGroupUid

`func (o *AutoscaleGroupPluginModel) GetDesktopGroupUid() int32`

GetDesktopGroupUid returns the DesktopGroupUid field if non-nil, zero value otherwise.

### GetDesktopGroupUidOk

`func (o *AutoscaleGroupPluginModel) GetDesktopGroupUidOk() (*int32, bool)`

GetDesktopGroupUidOk returns a tuple with the DesktopGroupUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopGroupUid

`func (o *AutoscaleGroupPluginModel) SetDesktopGroupUid(v int32)`

SetDesktopGroupUid sets DesktopGroupUid field to given value.

### HasDesktopGroupUid

`func (o *AutoscaleGroupPluginModel) HasDesktopGroupUid() bool`

HasDesktopGroupUid returns a boolean if a field has been set.

### GetType

`func (o *AutoscaleGroupPluginModel) GetType() AutoscalePluginType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoscaleGroupPluginModel) GetTypeOk() (*AutoscalePluginType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoscaleGroupPluginModel) SetType(v AutoscalePluginType)`

SetType sets Type field to given value.

### HasType

`func (o *AutoscaleGroupPluginModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAutoApprove

`func (o *AutoscaleGroupPluginModel) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *AutoscaleGroupPluginModel) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *AutoscaleGroupPluginModel) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.

### HasAutoApprove

`func (o *AutoscaleGroupPluginModel) HasAutoApprove() bool`

HasAutoApprove returns a boolean if a field has been set.

### GetApprovedUntil

`func (o *AutoscaleGroupPluginModel) GetApprovedUntil() string`

GetApprovedUntil returns the ApprovedUntil field if non-nil, zero value otherwise.

### GetApprovedUntilOk

`func (o *AutoscaleGroupPluginModel) GetApprovedUntilOk() (*string, bool)`

GetApprovedUntilOk returns a tuple with the ApprovedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedUntil

`func (o *AutoscaleGroupPluginModel) SetApprovedUntil(v string)`

SetApprovedUntil sets ApprovedUntil field to given value.

### HasApprovedUntil

`func (o *AutoscaleGroupPluginModel) HasApprovedUntil() bool`

HasApprovedUntil returns a boolean if a field has been set.

### SetApprovedUntilNil

`func (o *AutoscaleGroupPluginModel) SetApprovedUntilNil(b bool)`

 SetApprovedUntilNil sets the value for ApprovedUntil to be an explicit nil

### UnsetApprovedUntil
`func (o *AutoscaleGroupPluginModel) UnsetApprovedUntil()`

UnsetApprovedUntil ensures that no value is present for ApprovedUntil, not even an explicit nil
### GetDates

`func (o *AutoscaleGroupPluginModel) GetDates() []string`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *AutoscaleGroupPluginModel) GetDatesOk() (*[]string, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *AutoscaleGroupPluginModel) SetDates(v []string)`

SetDates sets Dates field to given value.

### HasDates

`func (o *AutoscaleGroupPluginModel) HasDates() bool`

HasDates returns a boolean if a field has been set.

### SetDatesNil

`func (o *AutoscaleGroupPluginModel) SetDatesNil(b bool)`

 SetDatesNil sets the value for Dates to be an explicit nil

### UnsetDates
`func (o *AutoscaleGroupPluginModel) UnsetDates()`

UnsetDates ensures that no value is present for Dates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


