# TagDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the tag. | 
**Id** | **string** | Id of the tag. Used to be: Guid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | 
**Description** | Pointer to **NullableString** | Description of the tag. | [optional] 
**Uid** | Pointer to **NullableInt32** | &#x60;DEPRECATED&#x60; DEPRECATED.  Use Id. | [optional] 
**IsUsedByGpo** | Pointer to **bool** | Used by Gpo or not. | [optional] 
**NumMachines** | **int32** | Number of machines the tag is applied to, within the context of the query. | 
**NumApplications** | **int32** | Number of applications the tag is applied to, within the context of the query. | 
**NumApplicationGroups** | **int32** | Number of application groups the tag is applied to, within the context of the query. | 
**NumMachineCatalogs** | **int32** | Number of machine catalogs the tag is applied to, within the context of the query. | 
**NumDeliveryGroups** | **int32** | Number of delivery groups the tag is applied to, within the context of the query. | 
**NumAutoscale** | Pointer to **int32** | Number of delivery groups using this tag for Autoscale, within the context of the query. | [optional] 
**NumUnknownObjects** | **int32** | The number of objects of all types that are tagged with this tag and that are *not* visible to the calling delegated administrator. | 
**ScopeReferences** | Pointer to [**[]ScopeReferenceModel**](ScopeReferenceModel.md) | The list of the delegated admin scopes to which the tag belongs. | [optional] 
**RestrictedDesktops** | [**[]RefResponseModel**](RefResponseModel.md) | Desktops which are restricted to be used with machines that are tagged with this tag. | 
**RestrictedApplicationGroups** | [**[]RefResponseModel**](RefResponseModel.md) | Application groups which are restricted such that the applications contained within will be launched only on machines that are tagged with this tag. | 
**RestrictedDeliveryGroups** | [**[]RefResponseModel**](RefResponseModel.md) | Delivery groups with reboot schedules that are restricted to machines that are tagged with this tag. | 
**RestrictedRebootSchedules** | [**[]RefResponseModel**](RefResponseModel.md) | Reboot schedules that are restricted to machines that are tagged with this tag. Only reboot schedules that are visible to the calling delegated administrator are counted. | 
**RestrictedAutoscale** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | Delivery groups that are using this tag for autoscale | [optional] 
**NumUnknownRestrictedObjects** | **int32** | The number of objects of all types that have this tag restriction and that are *not* visible to the calling delegated administrator. | 

## Methods

### NewTagDetailResponseModel

`func NewTagDetailResponseModel(name string, id string, numMachines int32, numApplications int32, numApplicationGroups int32, numMachineCatalogs int32, numDeliveryGroups int32, numUnknownObjects int32, restrictedDesktops []RefResponseModel, restrictedApplicationGroups []RefResponseModel, restrictedDeliveryGroups []RefResponseModel, restrictedRebootSchedules []RefResponseModel, numUnknownRestrictedObjects int32, ) *TagDetailResponseModel`

NewTagDetailResponseModel instantiates a new TagDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagDetailResponseModelWithDefaults

`func NewTagDetailResponseModelWithDefaults() *TagDetailResponseModel`

NewTagDetailResponseModelWithDefaults instantiates a new TagDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *TagDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *TagDetailResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagDetailResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagDetailResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagDetailResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TagDetailResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TagDetailResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUid

`func (o *TagDetailResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *TagDetailResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *TagDetailResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *TagDetailResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *TagDetailResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *TagDetailResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetIsUsedByGpo

`func (o *TagDetailResponseModel) GetIsUsedByGpo() bool`

GetIsUsedByGpo returns the IsUsedByGpo field if non-nil, zero value otherwise.

### GetIsUsedByGpoOk

`func (o *TagDetailResponseModel) GetIsUsedByGpoOk() (*bool, bool)`

GetIsUsedByGpoOk returns a tuple with the IsUsedByGpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedByGpo

`func (o *TagDetailResponseModel) SetIsUsedByGpo(v bool)`

SetIsUsedByGpo sets IsUsedByGpo field to given value.

### HasIsUsedByGpo

`func (o *TagDetailResponseModel) HasIsUsedByGpo() bool`

HasIsUsedByGpo returns a boolean if a field has been set.

### GetNumMachines

`func (o *TagDetailResponseModel) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *TagDetailResponseModel) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *TagDetailResponseModel) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.


### GetNumApplications

`func (o *TagDetailResponseModel) GetNumApplications() int32`

GetNumApplications returns the NumApplications field if non-nil, zero value otherwise.

### GetNumApplicationsOk

`func (o *TagDetailResponseModel) GetNumApplicationsOk() (*int32, bool)`

GetNumApplicationsOk returns a tuple with the NumApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApplications

`func (o *TagDetailResponseModel) SetNumApplications(v int32)`

SetNumApplications sets NumApplications field to given value.


### GetNumApplicationGroups

`func (o *TagDetailResponseModel) GetNumApplicationGroups() int32`

GetNumApplicationGroups returns the NumApplicationGroups field if non-nil, zero value otherwise.

### GetNumApplicationGroupsOk

`func (o *TagDetailResponseModel) GetNumApplicationGroupsOk() (*int32, bool)`

GetNumApplicationGroupsOk returns a tuple with the NumApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApplicationGroups

`func (o *TagDetailResponseModel) SetNumApplicationGroups(v int32)`

SetNumApplicationGroups sets NumApplicationGroups field to given value.


### GetNumMachineCatalogs

`func (o *TagDetailResponseModel) GetNumMachineCatalogs() int32`

GetNumMachineCatalogs returns the NumMachineCatalogs field if non-nil, zero value otherwise.

### GetNumMachineCatalogsOk

`func (o *TagDetailResponseModel) GetNumMachineCatalogsOk() (*int32, bool)`

GetNumMachineCatalogsOk returns a tuple with the NumMachineCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachineCatalogs

`func (o *TagDetailResponseModel) SetNumMachineCatalogs(v int32)`

SetNumMachineCatalogs sets NumMachineCatalogs field to given value.


### GetNumDeliveryGroups

`func (o *TagDetailResponseModel) GetNumDeliveryGroups() int32`

GetNumDeliveryGroups returns the NumDeliveryGroups field if non-nil, zero value otherwise.

### GetNumDeliveryGroupsOk

`func (o *TagDetailResponseModel) GetNumDeliveryGroupsOk() (*int32, bool)`

GetNumDeliveryGroupsOk returns a tuple with the NumDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeliveryGroups

`func (o *TagDetailResponseModel) SetNumDeliveryGroups(v int32)`

SetNumDeliveryGroups sets NumDeliveryGroups field to given value.


### GetNumAutoscale

`func (o *TagDetailResponseModel) GetNumAutoscale() int32`

GetNumAutoscale returns the NumAutoscale field if non-nil, zero value otherwise.

### GetNumAutoscaleOk

`func (o *TagDetailResponseModel) GetNumAutoscaleOk() (*int32, bool)`

GetNumAutoscaleOk returns a tuple with the NumAutoscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAutoscale

`func (o *TagDetailResponseModel) SetNumAutoscale(v int32)`

SetNumAutoscale sets NumAutoscale field to given value.

### HasNumAutoscale

`func (o *TagDetailResponseModel) HasNumAutoscale() bool`

HasNumAutoscale returns a boolean if a field has been set.

### GetNumUnknownObjects

`func (o *TagDetailResponseModel) GetNumUnknownObjects() int32`

GetNumUnknownObjects returns the NumUnknownObjects field if non-nil, zero value otherwise.

### GetNumUnknownObjectsOk

`func (o *TagDetailResponseModel) GetNumUnknownObjectsOk() (*int32, bool)`

GetNumUnknownObjectsOk returns a tuple with the NumUnknownObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnknownObjects

`func (o *TagDetailResponseModel) SetNumUnknownObjects(v int32)`

SetNumUnknownObjects sets NumUnknownObjects field to given value.


### GetScopeReferences

`func (o *TagDetailResponseModel) GetScopeReferences() []ScopeReferenceModel`

GetScopeReferences returns the ScopeReferences field if non-nil, zero value otherwise.

### GetScopeReferencesOk

`func (o *TagDetailResponseModel) GetScopeReferencesOk() (*[]ScopeReferenceModel, bool)`

GetScopeReferencesOk returns a tuple with the ScopeReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeReferences

`func (o *TagDetailResponseModel) SetScopeReferences(v []ScopeReferenceModel)`

SetScopeReferences sets ScopeReferences field to given value.

### HasScopeReferences

`func (o *TagDetailResponseModel) HasScopeReferences() bool`

HasScopeReferences returns a boolean if a field has been set.

### SetScopeReferencesNil

`func (o *TagDetailResponseModel) SetScopeReferencesNil(b bool)`

 SetScopeReferencesNil sets the value for ScopeReferences to be an explicit nil

### UnsetScopeReferences
`func (o *TagDetailResponseModel) UnsetScopeReferences()`

UnsetScopeReferences ensures that no value is present for ScopeReferences, not even an explicit nil
### GetRestrictedDesktops

`func (o *TagDetailResponseModel) GetRestrictedDesktops() []RefResponseModel`

GetRestrictedDesktops returns the RestrictedDesktops field if non-nil, zero value otherwise.

### GetRestrictedDesktopsOk

`func (o *TagDetailResponseModel) GetRestrictedDesktopsOk() (*[]RefResponseModel, bool)`

GetRestrictedDesktopsOk returns a tuple with the RestrictedDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDesktops

`func (o *TagDetailResponseModel) SetRestrictedDesktops(v []RefResponseModel)`

SetRestrictedDesktops sets RestrictedDesktops field to given value.


### GetRestrictedApplicationGroups

`func (o *TagDetailResponseModel) GetRestrictedApplicationGroups() []RefResponseModel`

GetRestrictedApplicationGroups returns the RestrictedApplicationGroups field if non-nil, zero value otherwise.

### GetRestrictedApplicationGroupsOk

`func (o *TagDetailResponseModel) GetRestrictedApplicationGroupsOk() (*[]RefResponseModel, bool)`

GetRestrictedApplicationGroupsOk returns a tuple with the RestrictedApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedApplicationGroups

`func (o *TagDetailResponseModel) SetRestrictedApplicationGroups(v []RefResponseModel)`

SetRestrictedApplicationGroups sets RestrictedApplicationGroups field to given value.


### GetRestrictedDeliveryGroups

`func (o *TagDetailResponseModel) GetRestrictedDeliveryGroups() []RefResponseModel`

GetRestrictedDeliveryGroups returns the RestrictedDeliveryGroups field if non-nil, zero value otherwise.

### GetRestrictedDeliveryGroupsOk

`func (o *TagDetailResponseModel) GetRestrictedDeliveryGroupsOk() (*[]RefResponseModel, bool)`

GetRestrictedDeliveryGroupsOk returns a tuple with the RestrictedDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDeliveryGroups

`func (o *TagDetailResponseModel) SetRestrictedDeliveryGroups(v []RefResponseModel)`

SetRestrictedDeliveryGroups sets RestrictedDeliveryGroups field to given value.


### GetRestrictedRebootSchedules

`func (o *TagDetailResponseModel) GetRestrictedRebootSchedules() []RefResponseModel`

GetRestrictedRebootSchedules returns the RestrictedRebootSchedules field if non-nil, zero value otherwise.

### GetRestrictedRebootSchedulesOk

`func (o *TagDetailResponseModel) GetRestrictedRebootSchedulesOk() (*[]RefResponseModel, bool)`

GetRestrictedRebootSchedulesOk returns a tuple with the RestrictedRebootSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedRebootSchedules

`func (o *TagDetailResponseModel) SetRestrictedRebootSchedules(v []RefResponseModel)`

SetRestrictedRebootSchedules sets RestrictedRebootSchedules field to given value.


### GetRestrictedAutoscale

`func (o *TagDetailResponseModel) GetRestrictedAutoscale() []RefResponseModel`

GetRestrictedAutoscale returns the RestrictedAutoscale field if non-nil, zero value otherwise.

### GetRestrictedAutoscaleOk

`func (o *TagDetailResponseModel) GetRestrictedAutoscaleOk() (*[]RefResponseModel, bool)`

GetRestrictedAutoscaleOk returns a tuple with the RestrictedAutoscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedAutoscale

`func (o *TagDetailResponseModel) SetRestrictedAutoscale(v []RefResponseModel)`

SetRestrictedAutoscale sets RestrictedAutoscale field to given value.

### HasRestrictedAutoscale

`func (o *TagDetailResponseModel) HasRestrictedAutoscale() bool`

HasRestrictedAutoscale returns a boolean if a field has been set.

### SetRestrictedAutoscaleNil

`func (o *TagDetailResponseModel) SetRestrictedAutoscaleNil(b bool)`

 SetRestrictedAutoscaleNil sets the value for RestrictedAutoscale to be an explicit nil

### UnsetRestrictedAutoscale
`func (o *TagDetailResponseModel) UnsetRestrictedAutoscale()`

UnsetRestrictedAutoscale ensures that no value is present for RestrictedAutoscale, not even an explicit nil
### GetNumUnknownRestrictedObjects

`func (o *TagDetailResponseModel) GetNumUnknownRestrictedObjects() int32`

GetNumUnknownRestrictedObjects returns the NumUnknownRestrictedObjects field if non-nil, zero value otherwise.

### GetNumUnknownRestrictedObjectsOk

`func (o *TagDetailResponseModel) GetNumUnknownRestrictedObjectsOk() (*int32, bool)`

GetNumUnknownRestrictedObjectsOk returns a tuple with the NumUnknownRestrictedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnknownRestrictedObjects

`func (o *TagDetailResponseModel) SetNumUnknownRestrictedObjects(v int32)`

SetNumUnknownRestrictedObjects sets NumUnknownRestrictedObjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


