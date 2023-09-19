# EditApplicationFolderRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the application folder.  If not set, will not be changed.  The name must be unique within the folder&#39;s parent folder. | [optional] 
**Parent** | Pointer to **NullableString** | Parent folder.  If not set, will not be changed.  Can be set to either Id or Path.  To move the application to the root folder, specify the empty string (\&quot;\&quot;).  If specified as a path, and the path does not exist, it will be created. | [optional] 

## Methods

### NewEditApplicationFolderRequestModel

`func NewEditApplicationFolderRequestModel() *EditApplicationFolderRequestModel`

NewEditApplicationFolderRequestModel instantiates a new EditApplicationFolderRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditApplicationFolderRequestModelWithDefaults

`func NewEditApplicationFolderRequestModelWithDefaults() *EditApplicationFolderRequestModel`

NewEditApplicationFolderRequestModelWithDefaults instantiates a new EditApplicationFolderRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditApplicationFolderRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditApplicationFolderRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditApplicationFolderRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditApplicationFolderRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditApplicationFolderRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditApplicationFolderRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParent

`func (o *EditApplicationFolderRequestModel) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EditApplicationFolderRequestModel) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EditApplicationFolderRequestModel) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EditApplicationFolderRequestModel) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *EditApplicationFolderRequestModel) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *EditApplicationFolderRequestModel) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


