# CreateAdminFolderRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the folder to create. May be omitted if Path is used; otherwise required.  If both  and  are specified, a folder with the specified name will be created as a child of the specified path. | 
**Path** | Pointer to **NullableString** | Path where the folder should be created. Exclusive with ParentId.  One or the other property may be used but not both.  Path to where the folder should be created. Any folders in the path which don&#39;t exist will be created.  If not specified, default is the root folder path. | [optional] 
**ParentId** | Pointer to **NullableString** | Parent folder where the folder should be created. Exclusive with Path.  One or the other property may be used but not both.  If not specified, default is the root folder path. | [optional] 
**ObjectIdentifiers** | [**[]AdminFolderObjectIdentifier**](AdminFolderObjectIdentifier.md) | The object identifiers of the new admin folder. To indicated what objects this folder will include. | 

## Methods

### NewCreateAdminFolderRequestModel

`func NewCreateAdminFolderRequestModel(name string, objectIdentifiers []AdminFolderObjectIdentifier, ) *CreateAdminFolderRequestModel`

NewCreateAdminFolderRequestModel instantiates a new CreateAdminFolderRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdminFolderRequestModelWithDefaults

`func NewCreateAdminFolderRequestModelWithDefaults() *CreateAdminFolderRequestModel`

NewCreateAdminFolderRequestModelWithDefaults instantiates a new CreateAdminFolderRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAdminFolderRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAdminFolderRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAdminFolderRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *CreateAdminFolderRequestModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateAdminFolderRequestModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateAdminFolderRequestModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateAdminFolderRequestModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *CreateAdminFolderRequestModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CreateAdminFolderRequestModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetParentId

`func (o *CreateAdminFolderRequestModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateAdminFolderRequestModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateAdminFolderRequestModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateAdminFolderRequestModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *CreateAdminFolderRequestModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *CreateAdminFolderRequestModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetObjectIdentifiers

`func (o *CreateAdminFolderRequestModel) GetObjectIdentifiers() []AdminFolderObjectIdentifier`

GetObjectIdentifiers returns the ObjectIdentifiers field if non-nil, zero value otherwise.

### GetObjectIdentifiersOk

`func (o *CreateAdminFolderRequestModel) GetObjectIdentifiersOk() (*[]AdminFolderObjectIdentifier, bool)`

GetObjectIdentifiersOk returns a tuple with the ObjectIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIdentifiers

`func (o *CreateAdminFolderRequestModel) SetObjectIdentifiers(v []AdminFolderObjectIdentifier)`

SetObjectIdentifiers sets ObjectIdentifiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


