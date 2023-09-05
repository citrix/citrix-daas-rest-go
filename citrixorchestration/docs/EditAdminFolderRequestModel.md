# EditAdminFolderRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the admin folder.  If not set, will not be changed.  The name must be unique within the folder&#39;s parent folder. | [optional] 
**Parent** | Pointer to **string** | Parent folder.  If not set, will not be changed.  Can be set to either Id or Path.  To move the application to the root folder, specify the empty string (\&quot;\&quot;).  If specified as a path, and the path does not exist, it will be created. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of Admin Folder. Set the value of the NameValueStringPairModel is null will not update MetadataMap Set the value of the NameValueStringPairModel empty will be remove this metadata. Not existing Name and Value NameValueStringPairModel object will be added. The same Name but different value object will be updated. | [optional] 

## Methods

### NewEditAdminFolderRequestModel

`func NewEditAdminFolderRequestModel() *EditAdminFolderRequestModel`

NewEditAdminFolderRequestModel instantiates a new EditAdminFolderRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAdminFolderRequestModelWithDefaults

`func NewEditAdminFolderRequestModelWithDefaults() *EditAdminFolderRequestModel`

NewEditAdminFolderRequestModelWithDefaults instantiates a new EditAdminFolderRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditAdminFolderRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditAdminFolderRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditAdminFolderRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditAdminFolderRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *EditAdminFolderRequestModel) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EditAdminFolderRequestModel) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EditAdminFolderRequestModel) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EditAdminFolderRequestModel) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetMetadata

`func (o *EditAdminFolderRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EditAdminFolderRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EditAdminFolderRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EditAdminFolderRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


