# CreateAdminScopeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the scope. Name is globally unique. | 
**Description** | Pointer to **string** | The description of the admin scope. | [optional] 
**IsTenantScope** | Pointer to **bool** | Indicates if it is tenant scope. | [optional] 
**ScopedObjects** | Pointer to [**[]ScopedObjectRequestModel**](ScopedObjectRequestModel.md) | Scoped objects associated with the scope object. | [optional] 

## Methods

### NewCreateAdminScopeRequestModel

`func NewCreateAdminScopeRequestModel(name string, ) *CreateAdminScopeRequestModel`

NewCreateAdminScopeRequestModel instantiates a new CreateAdminScopeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdminScopeRequestModelWithDefaults

`func NewCreateAdminScopeRequestModelWithDefaults() *CreateAdminScopeRequestModel`

NewCreateAdminScopeRequestModelWithDefaults instantiates a new CreateAdminScopeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAdminScopeRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAdminScopeRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAdminScopeRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAdminScopeRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAdminScopeRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAdminScopeRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAdminScopeRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsTenantScope

`func (o *CreateAdminScopeRequestModel) GetIsTenantScope() bool`

GetIsTenantScope returns the IsTenantScope field if non-nil, zero value otherwise.

### GetIsTenantScopeOk

`func (o *CreateAdminScopeRequestModel) GetIsTenantScopeOk() (*bool, bool)`

GetIsTenantScopeOk returns a tuple with the IsTenantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTenantScope

`func (o *CreateAdminScopeRequestModel) SetIsTenantScope(v bool)`

SetIsTenantScope sets IsTenantScope field to given value.

### HasIsTenantScope

`func (o *CreateAdminScopeRequestModel) HasIsTenantScope() bool`

HasIsTenantScope returns a boolean if a field has been set.

### GetScopedObjects

`func (o *CreateAdminScopeRequestModel) GetScopedObjects() []ScopedObjectRequestModel`

GetScopedObjects returns the ScopedObjects field if non-nil, zero value otherwise.

### GetScopedObjectsOk

`func (o *CreateAdminScopeRequestModel) GetScopedObjectsOk() (*[]ScopedObjectRequestModel, bool)`

GetScopedObjectsOk returns a tuple with the ScopedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedObjects

`func (o *CreateAdminScopeRequestModel) SetScopedObjects(v []ScopedObjectRequestModel)`

SetScopedObjects sets ScopedObjects field to given value.

### HasScopedObjects

`func (o *CreateAdminScopeRequestModel) HasScopedObjects() bool`

HasScopedObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


