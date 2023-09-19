# EditAdminScopeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the scope. Name is globally unique. | [optional] 
**Description** | Pointer to **NullableString** | The description of the admin scope. | [optional] 
**ScopedObjects** | Pointer to [**[]ScopedObjectRequestModel**](ScopedObjectRequestModel.md) | Scoped objects associated with the scope object. | [optional] 

## Methods

### NewEditAdminScopeRequestModel

`func NewEditAdminScopeRequestModel() *EditAdminScopeRequestModel`

NewEditAdminScopeRequestModel instantiates a new EditAdminScopeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAdminScopeRequestModelWithDefaults

`func NewEditAdminScopeRequestModelWithDefaults() *EditAdminScopeRequestModel`

NewEditAdminScopeRequestModelWithDefaults instantiates a new EditAdminScopeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditAdminScopeRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditAdminScopeRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditAdminScopeRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditAdminScopeRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditAdminScopeRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditAdminScopeRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *EditAdminScopeRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditAdminScopeRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditAdminScopeRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditAdminScopeRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EditAdminScopeRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditAdminScopeRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScopedObjects

`func (o *EditAdminScopeRequestModel) GetScopedObjects() []ScopedObjectRequestModel`

GetScopedObjects returns the ScopedObjects field if non-nil, zero value otherwise.

### GetScopedObjectsOk

`func (o *EditAdminScopeRequestModel) GetScopedObjectsOk() (*[]ScopedObjectRequestModel, bool)`

GetScopedObjectsOk returns a tuple with the ScopedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedObjects

`func (o *EditAdminScopeRequestModel) SetScopedObjects(v []ScopedObjectRequestModel)`

SetScopedObjects sets ScopedObjects field to given value.

### HasScopedObjects

`func (o *EditAdminScopeRequestModel) HasScopedObjects() bool`

HasScopedObjects returns a boolean if a field has been set.

### SetScopedObjectsNil

`func (o *EditAdminScopeRequestModel) SetScopedObjectsNil(b bool)`

 SetScopedObjectsNil sets the value for ScopedObjects to be an explicit nil

### UnsetScopedObjects
`func (o *EditAdminScopeRequestModel) UnsetScopedObjects()`

UnsetScopedObjects ensures that no value is present for ScopedObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


