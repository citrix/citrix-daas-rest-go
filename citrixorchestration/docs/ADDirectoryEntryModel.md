# ADDirectoryEntryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the object | [optional] 
**Path** | Pointer to **string** | Path of the object | [optional] 
**Forest** | Pointer to **string** | Forest of the object | [optional] 
**Tag** | Pointer to **map[string]interface{}** | Cookie for use by the IActiveDirectoryService implementation | [optional] 
**CanContainComputer** | Pointer to **bool** | True if this object can directly contain computer accounts | [optional] 
**CanIndirectlyContainComputer** | Pointer to **bool** | True if this object can indirectly contain computer accounts | [optional] 
**Parent** | Pointer to [**ADDirectoryEntryModelParent**](ADDirectoryEntryModelParent.md) |  | [optional] 
**Children** | Pointer to [**[]ADDirectoryEntryModel**](ADDirectoryEntryModel.md) | Child entries of this object | [optional] 

## Methods

### NewADDirectoryEntryModel

`func NewADDirectoryEntryModel() *ADDirectoryEntryModel`

NewADDirectoryEntryModel instantiates a new ADDirectoryEntryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDirectoryEntryModelWithDefaults

`func NewADDirectoryEntryModelWithDefaults() *ADDirectoryEntryModel`

NewADDirectoryEntryModelWithDefaults instantiates a new ADDirectoryEntryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ADDirectoryEntryModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ADDirectoryEntryModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ADDirectoryEntryModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ADDirectoryEntryModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ADDirectoryEntryModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ADDirectoryEntryModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ADDirectoryEntryModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ADDirectoryEntryModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetForest

`func (o *ADDirectoryEntryModel) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *ADDirectoryEntryModel) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *ADDirectoryEntryModel) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *ADDirectoryEntryModel) HasForest() bool`

HasForest returns a boolean if a field has been set.

### GetTag

`func (o *ADDirectoryEntryModel) GetTag() map[string]interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ADDirectoryEntryModel) GetTagOk() (*map[string]interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ADDirectoryEntryModel) SetTag(v map[string]interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ADDirectoryEntryModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetCanContainComputer

`func (o *ADDirectoryEntryModel) GetCanContainComputer() bool`

GetCanContainComputer returns the CanContainComputer field if non-nil, zero value otherwise.

### GetCanContainComputerOk

`func (o *ADDirectoryEntryModel) GetCanContainComputerOk() (*bool, bool)`

GetCanContainComputerOk returns a tuple with the CanContainComputer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanContainComputer

`func (o *ADDirectoryEntryModel) SetCanContainComputer(v bool)`

SetCanContainComputer sets CanContainComputer field to given value.

### HasCanContainComputer

`func (o *ADDirectoryEntryModel) HasCanContainComputer() bool`

HasCanContainComputer returns a boolean if a field has been set.

### GetCanIndirectlyContainComputer

`func (o *ADDirectoryEntryModel) GetCanIndirectlyContainComputer() bool`

GetCanIndirectlyContainComputer returns the CanIndirectlyContainComputer field if non-nil, zero value otherwise.

### GetCanIndirectlyContainComputerOk

`func (o *ADDirectoryEntryModel) GetCanIndirectlyContainComputerOk() (*bool, bool)`

GetCanIndirectlyContainComputerOk returns a tuple with the CanIndirectlyContainComputer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanIndirectlyContainComputer

`func (o *ADDirectoryEntryModel) SetCanIndirectlyContainComputer(v bool)`

SetCanIndirectlyContainComputer sets CanIndirectlyContainComputer field to given value.

### HasCanIndirectlyContainComputer

`func (o *ADDirectoryEntryModel) HasCanIndirectlyContainComputer() bool`

HasCanIndirectlyContainComputer returns a boolean if a field has been set.

### GetParent

`func (o *ADDirectoryEntryModel) GetParent() ADDirectoryEntryModelParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ADDirectoryEntryModel) GetParentOk() (*ADDirectoryEntryModelParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ADDirectoryEntryModel) SetParent(v ADDirectoryEntryModelParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ADDirectoryEntryModel) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetChildren

`func (o *ADDirectoryEntryModel) GetChildren() []ADDirectoryEntryModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ADDirectoryEntryModel) GetChildrenOk() (*[]ADDirectoryEntryModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ADDirectoryEntryModel) SetChildren(v []ADDirectoryEntryModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ADDirectoryEntryModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


