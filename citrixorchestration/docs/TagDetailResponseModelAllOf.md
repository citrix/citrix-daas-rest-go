# TagDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictedDesktops** | [**[]RefResponseModel**](RefResponseModel.md) | Desktops which are restricted to be used with machines that are tagged with this tag. | 
**RestrictedApplicationGroups** | [**[]RefResponseModel**](RefResponseModel.md) | Application groups which are restricted such that the applications contained within will be launched only on machines that are tagged with this tag. | 
**RestrictedDeliveryGroups** | [**[]RefResponseModel**](RefResponseModel.md) | Delivery groups with reboot schedules that are restricted to machines that are tagged with this tag. | 
**RestrictedRebootSchedules** | [**[]RefResponseModel**](RefResponseModel.md) | Reboot schedules that are restricted to machines that are tagged with this tag. Only reboot schedules that are visible to the calling delegated administrator are counted. | 
**RestrictedAutoscale** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | Delivery groups that are using this tag for autoscale | [optional] 
**NumUnknownRestrictedObjects** | **int32** | The number of objects of all types that have this tag restriction and that are *not* visible to the calling delegated administrator. | 

## Methods

### NewTagDetailResponseModelAllOf

`func NewTagDetailResponseModelAllOf(restrictedDesktops []RefResponseModel, restrictedApplicationGroups []RefResponseModel, restrictedDeliveryGroups []RefResponseModel, restrictedRebootSchedules []RefResponseModel, numUnknownRestrictedObjects int32, ) *TagDetailResponseModelAllOf`

NewTagDetailResponseModelAllOf instantiates a new TagDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagDetailResponseModelAllOfWithDefaults

`func NewTagDetailResponseModelAllOfWithDefaults() *TagDetailResponseModelAllOf`

NewTagDetailResponseModelAllOfWithDefaults instantiates a new TagDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictedDesktops

`func (o *TagDetailResponseModelAllOf) GetRestrictedDesktops() []RefResponseModel`

GetRestrictedDesktops returns the RestrictedDesktops field if non-nil, zero value otherwise.

### GetRestrictedDesktopsOk

`func (o *TagDetailResponseModelAllOf) GetRestrictedDesktopsOk() (*[]RefResponseModel, bool)`

GetRestrictedDesktopsOk returns a tuple with the RestrictedDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDesktops

`func (o *TagDetailResponseModelAllOf) SetRestrictedDesktops(v []RefResponseModel)`

SetRestrictedDesktops sets RestrictedDesktops field to given value.


### GetRestrictedApplicationGroups

`func (o *TagDetailResponseModelAllOf) GetRestrictedApplicationGroups() []RefResponseModel`

GetRestrictedApplicationGroups returns the RestrictedApplicationGroups field if non-nil, zero value otherwise.

### GetRestrictedApplicationGroupsOk

`func (o *TagDetailResponseModelAllOf) GetRestrictedApplicationGroupsOk() (*[]RefResponseModel, bool)`

GetRestrictedApplicationGroupsOk returns a tuple with the RestrictedApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedApplicationGroups

`func (o *TagDetailResponseModelAllOf) SetRestrictedApplicationGroups(v []RefResponseModel)`

SetRestrictedApplicationGroups sets RestrictedApplicationGroups field to given value.


### GetRestrictedDeliveryGroups

`func (o *TagDetailResponseModelAllOf) GetRestrictedDeliveryGroups() []RefResponseModel`

GetRestrictedDeliveryGroups returns the RestrictedDeliveryGroups field if non-nil, zero value otherwise.

### GetRestrictedDeliveryGroupsOk

`func (o *TagDetailResponseModelAllOf) GetRestrictedDeliveryGroupsOk() (*[]RefResponseModel, bool)`

GetRestrictedDeliveryGroupsOk returns a tuple with the RestrictedDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDeliveryGroups

`func (o *TagDetailResponseModelAllOf) SetRestrictedDeliveryGroups(v []RefResponseModel)`

SetRestrictedDeliveryGroups sets RestrictedDeliveryGroups field to given value.


### GetRestrictedRebootSchedules

`func (o *TagDetailResponseModelAllOf) GetRestrictedRebootSchedules() []RefResponseModel`

GetRestrictedRebootSchedules returns the RestrictedRebootSchedules field if non-nil, zero value otherwise.

### GetRestrictedRebootSchedulesOk

`func (o *TagDetailResponseModelAllOf) GetRestrictedRebootSchedulesOk() (*[]RefResponseModel, bool)`

GetRestrictedRebootSchedulesOk returns a tuple with the RestrictedRebootSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedRebootSchedules

`func (o *TagDetailResponseModelAllOf) SetRestrictedRebootSchedules(v []RefResponseModel)`

SetRestrictedRebootSchedules sets RestrictedRebootSchedules field to given value.


### GetRestrictedAutoscale

`func (o *TagDetailResponseModelAllOf) GetRestrictedAutoscale() []RefResponseModel`

GetRestrictedAutoscale returns the RestrictedAutoscale field if non-nil, zero value otherwise.

### GetRestrictedAutoscaleOk

`func (o *TagDetailResponseModelAllOf) GetRestrictedAutoscaleOk() (*[]RefResponseModel, bool)`

GetRestrictedAutoscaleOk returns a tuple with the RestrictedAutoscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedAutoscale

`func (o *TagDetailResponseModelAllOf) SetRestrictedAutoscale(v []RefResponseModel)`

SetRestrictedAutoscale sets RestrictedAutoscale field to given value.

### HasRestrictedAutoscale

`func (o *TagDetailResponseModelAllOf) HasRestrictedAutoscale() bool`

HasRestrictedAutoscale returns a boolean if a field has been set.

### GetNumUnknownRestrictedObjects

`func (o *TagDetailResponseModelAllOf) GetNumUnknownRestrictedObjects() int32`

GetNumUnknownRestrictedObjects returns the NumUnknownRestrictedObjects field if non-nil, zero value otherwise.

### GetNumUnknownRestrictedObjectsOk

`func (o *TagDetailResponseModelAllOf) GetNumUnknownRestrictedObjectsOk() (*int32, bool)`

GetNumUnknownRestrictedObjectsOk returns a tuple with the NumUnknownRestrictedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnknownRestrictedObjects

`func (o *TagDetailResponseModelAllOf) SetNumUnknownRestrictedObjects(v int32)`

SetNumUnknownRestrictedObjects sets NumUnknownRestrictedObjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


