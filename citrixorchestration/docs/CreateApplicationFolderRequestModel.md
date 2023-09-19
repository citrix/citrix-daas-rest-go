# CreateApplicationFolderRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the folder to create. May be omitted if Path is used; otherwise required.  If both  and  are specified, a folder with the specified name will be created as a child of the specified path. | [optional] 
**Path** | Pointer to **NullableString** | Path where the folder should be created. Exclusive with ParentId.  One or the other property may be used but not both.  Path to where the folder should be created. Any folders in the path which don&#39;t exist will be created.  If not specified, default is the root folder path. | [optional] 
**ParentId** | Pointer to **NullableString** | Parent folder where the folder should be created. Exclusive with Path.  One or the other property may be used but not both.  If not specified, default is the root folder path. | [optional] 

## Methods

### NewCreateApplicationFolderRequestModel

`func NewCreateApplicationFolderRequestModel() *CreateApplicationFolderRequestModel`

NewCreateApplicationFolderRequestModel instantiates a new CreateApplicationFolderRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationFolderRequestModelWithDefaults

`func NewCreateApplicationFolderRequestModelWithDefaults() *CreateApplicationFolderRequestModel`

NewCreateApplicationFolderRequestModelWithDefaults instantiates a new CreateApplicationFolderRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateApplicationFolderRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApplicationFolderRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApplicationFolderRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateApplicationFolderRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateApplicationFolderRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateApplicationFolderRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *CreateApplicationFolderRequestModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateApplicationFolderRequestModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateApplicationFolderRequestModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateApplicationFolderRequestModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *CreateApplicationFolderRequestModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CreateApplicationFolderRequestModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetParentId

`func (o *CreateApplicationFolderRequestModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateApplicationFolderRequestModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateApplicationFolderRequestModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateApplicationFolderRequestModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *CreateApplicationFolderRequestModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *CreateApplicationFolderRequestModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


