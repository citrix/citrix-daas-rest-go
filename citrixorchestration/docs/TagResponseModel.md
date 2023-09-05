# TagResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the tag. | 
**Id** | **string** | Id of the tag. Used to be: Guid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | 
**Description** | Pointer to **string** | Description of the tag. | [optional] 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED&#x60; DEPRECATED.  Use Id. | [optional] 
**IsUsedByGpo** | Pointer to **bool** | Used by Gpo or not. | [optional] 
**NumMachines** | **int32** | Number of machines the tag is applied to, within the context of the query. | 
**NumApplications** | **int32** | Number of applications the tag is applied to, within the context of the query. | 
**NumApplicationGroups** | **int32** | Number of application groups the tag is applied to, within the context of the query. | 
**NumMachineCatalogs** | **int32** | Number of machine catalogs the tag is applied to, within the context of the query. | 
**NumDeliveryGroups** | **int32** | Number of delivery groups the tag is applied to, within the context of the query. | 
**NumAutoscale** | Pointer to **int32** | Number of delivery groups using this tag for Autoscale, within the context of the query. | [optional] 
**NumUnknownObjects** | **int32** | The number of objects of all types that are tagged with this tag and that are *not* visible to the calling delegated administrator. | 
**ScopeReferences** | Pointer to [**[]ScopeReferenceModel**](ScopeReferenceModel.md) | The list of the delegated admin scopes to which the tag belongs. | [optional] 

## Methods

### NewTagResponseModel

`func NewTagResponseModel(name string, id string, numMachines int32, numApplications int32, numApplicationGroups int32, numMachineCatalogs int32, numDeliveryGroups int32, numUnknownObjects int32, ) *TagResponseModel`

NewTagResponseModel instantiates a new TagResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagResponseModelWithDefaults

`func NewTagResponseModelWithDefaults() *TagResponseModel`

NewTagResponseModelWithDefaults instantiates a new TagResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *TagResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *TagResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUid

`func (o *TagResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *TagResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *TagResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *TagResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetIsUsedByGpo

`func (o *TagResponseModel) GetIsUsedByGpo() bool`

GetIsUsedByGpo returns the IsUsedByGpo field if non-nil, zero value otherwise.

### GetIsUsedByGpoOk

`func (o *TagResponseModel) GetIsUsedByGpoOk() (*bool, bool)`

GetIsUsedByGpoOk returns a tuple with the IsUsedByGpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedByGpo

`func (o *TagResponseModel) SetIsUsedByGpo(v bool)`

SetIsUsedByGpo sets IsUsedByGpo field to given value.

### HasIsUsedByGpo

`func (o *TagResponseModel) HasIsUsedByGpo() bool`

HasIsUsedByGpo returns a boolean if a field has been set.

### GetNumMachines

`func (o *TagResponseModel) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *TagResponseModel) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *TagResponseModel) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.


### GetNumApplications

`func (o *TagResponseModel) GetNumApplications() int32`

GetNumApplications returns the NumApplications field if non-nil, zero value otherwise.

### GetNumApplicationsOk

`func (o *TagResponseModel) GetNumApplicationsOk() (*int32, bool)`

GetNumApplicationsOk returns a tuple with the NumApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApplications

`func (o *TagResponseModel) SetNumApplications(v int32)`

SetNumApplications sets NumApplications field to given value.


### GetNumApplicationGroups

`func (o *TagResponseModel) GetNumApplicationGroups() int32`

GetNumApplicationGroups returns the NumApplicationGroups field if non-nil, zero value otherwise.

### GetNumApplicationGroupsOk

`func (o *TagResponseModel) GetNumApplicationGroupsOk() (*int32, bool)`

GetNumApplicationGroupsOk returns a tuple with the NumApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApplicationGroups

`func (o *TagResponseModel) SetNumApplicationGroups(v int32)`

SetNumApplicationGroups sets NumApplicationGroups field to given value.


### GetNumMachineCatalogs

`func (o *TagResponseModel) GetNumMachineCatalogs() int32`

GetNumMachineCatalogs returns the NumMachineCatalogs field if non-nil, zero value otherwise.

### GetNumMachineCatalogsOk

`func (o *TagResponseModel) GetNumMachineCatalogsOk() (*int32, bool)`

GetNumMachineCatalogsOk returns a tuple with the NumMachineCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachineCatalogs

`func (o *TagResponseModel) SetNumMachineCatalogs(v int32)`

SetNumMachineCatalogs sets NumMachineCatalogs field to given value.


### GetNumDeliveryGroups

`func (o *TagResponseModel) GetNumDeliveryGroups() int32`

GetNumDeliveryGroups returns the NumDeliveryGroups field if non-nil, zero value otherwise.

### GetNumDeliveryGroupsOk

`func (o *TagResponseModel) GetNumDeliveryGroupsOk() (*int32, bool)`

GetNumDeliveryGroupsOk returns a tuple with the NumDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeliveryGroups

`func (o *TagResponseModel) SetNumDeliveryGroups(v int32)`

SetNumDeliveryGroups sets NumDeliveryGroups field to given value.


### GetNumAutoscale

`func (o *TagResponseModel) GetNumAutoscale() int32`

GetNumAutoscale returns the NumAutoscale field if non-nil, zero value otherwise.

### GetNumAutoscaleOk

`func (o *TagResponseModel) GetNumAutoscaleOk() (*int32, bool)`

GetNumAutoscaleOk returns a tuple with the NumAutoscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAutoscale

`func (o *TagResponseModel) SetNumAutoscale(v int32)`

SetNumAutoscale sets NumAutoscale field to given value.

### HasNumAutoscale

`func (o *TagResponseModel) HasNumAutoscale() bool`

HasNumAutoscale returns a boolean if a field has been set.

### GetNumUnknownObjects

`func (o *TagResponseModel) GetNumUnknownObjects() int32`

GetNumUnknownObjects returns the NumUnknownObjects field if non-nil, zero value otherwise.

### GetNumUnknownObjectsOk

`func (o *TagResponseModel) GetNumUnknownObjectsOk() (*int32, bool)`

GetNumUnknownObjectsOk returns a tuple with the NumUnknownObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnknownObjects

`func (o *TagResponseModel) SetNumUnknownObjects(v int32)`

SetNumUnknownObjects sets NumUnknownObjects field to given value.


### GetScopeReferences

`func (o *TagResponseModel) GetScopeReferences() []ScopeReferenceModel`

GetScopeReferences returns the ScopeReferences field if non-nil, zero value otherwise.

### GetScopeReferencesOk

`func (o *TagResponseModel) GetScopeReferencesOk() (*[]ScopeReferenceModel, bool)`

GetScopeReferencesOk returns a tuple with the ScopeReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeReferences

`func (o *TagResponseModel) SetScopeReferences(v []ScopeReferenceModel)`

SetScopeReferences sets ScopeReferences field to given value.

### HasScopeReferences

`func (o *TagResponseModel) HasScopeReferences() bool`

HasScopeReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


