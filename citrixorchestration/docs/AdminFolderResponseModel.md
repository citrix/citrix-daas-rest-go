# AdminFolderResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the admin folder. Used to be: Uuid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **NullableInt32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED. Use Id. | [optional] 
**Children** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | Child folders. | [optional] 
**Name** | Pointer to **NullableString** | Name of the folder. | [optional] 
**Parent** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**Path** | Pointer to **NullableString** | Full path to the folder. Used to be: FullName | [optional] 
**TotalApplications** | Pointer to **int32** | Total number of applications in the folder. | [optional] 
**TotalMachineCatalogs** | Pointer to **int32** | Total number of catalogs in the folder. | [optional] 
**TotalApplicationGroups** | Pointer to **int32** | Total number of application groups in the folder. | [optional] 
**TotalDesktopGroups** | Pointer to **int32** | Total number of desktop groups in the folder. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata map of the admin folder. | [optional] 

## Methods

### NewAdminFolderResponseModel

`func NewAdminFolderResponseModel() *AdminFolderResponseModel`

NewAdminFolderResponseModel instantiates a new AdminFolderResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminFolderResponseModelWithDefaults

`func NewAdminFolderResponseModelWithDefaults() *AdminFolderResponseModel`

NewAdminFolderResponseModelWithDefaults instantiates a new AdminFolderResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminFolderResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminFolderResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminFolderResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdminFolderResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AdminFolderResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AdminFolderResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *AdminFolderResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AdminFolderResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AdminFolderResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AdminFolderResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *AdminFolderResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *AdminFolderResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetChildren

`func (o *AdminFolderResponseModel) GetChildren() []RefResponseModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AdminFolderResponseModel) GetChildrenOk() (*[]RefResponseModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AdminFolderResponseModel) SetChildren(v []RefResponseModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AdminFolderResponseModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *AdminFolderResponseModel) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *AdminFolderResponseModel) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetName

`func (o *AdminFolderResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminFolderResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminFolderResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminFolderResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdminFolderResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminFolderResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParent

`func (o *AdminFolderResponseModel) GetParent() RefResponseModel`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AdminFolderResponseModel) GetParentOk() (*RefResponseModel, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AdminFolderResponseModel) SetParent(v RefResponseModel)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AdminFolderResponseModel) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPath

`func (o *AdminFolderResponseModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AdminFolderResponseModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AdminFolderResponseModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AdminFolderResponseModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *AdminFolderResponseModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *AdminFolderResponseModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetTotalApplications

`func (o *AdminFolderResponseModel) GetTotalApplications() int32`

GetTotalApplications returns the TotalApplications field if non-nil, zero value otherwise.

### GetTotalApplicationsOk

`func (o *AdminFolderResponseModel) GetTotalApplicationsOk() (*int32, bool)`

GetTotalApplicationsOk returns a tuple with the TotalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplications

`func (o *AdminFolderResponseModel) SetTotalApplications(v int32)`

SetTotalApplications sets TotalApplications field to given value.

### HasTotalApplications

`func (o *AdminFolderResponseModel) HasTotalApplications() bool`

HasTotalApplications returns a boolean if a field has been set.

### GetTotalMachineCatalogs

`func (o *AdminFolderResponseModel) GetTotalMachineCatalogs() int32`

GetTotalMachineCatalogs returns the TotalMachineCatalogs field if non-nil, zero value otherwise.

### GetTotalMachineCatalogsOk

`func (o *AdminFolderResponseModel) GetTotalMachineCatalogsOk() (*int32, bool)`

GetTotalMachineCatalogsOk returns a tuple with the TotalMachineCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachineCatalogs

`func (o *AdminFolderResponseModel) SetTotalMachineCatalogs(v int32)`

SetTotalMachineCatalogs sets TotalMachineCatalogs field to given value.

### HasTotalMachineCatalogs

`func (o *AdminFolderResponseModel) HasTotalMachineCatalogs() bool`

HasTotalMachineCatalogs returns a boolean if a field has been set.

### GetTotalApplicationGroups

`func (o *AdminFolderResponseModel) GetTotalApplicationGroups() int32`

GetTotalApplicationGroups returns the TotalApplicationGroups field if non-nil, zero value otherwise.

### GetTotalApplicationGroupsOk

`func (o *AdminFolderResponseModel) GetTotalApplicationGroupsOk() (*int32, bool)`

GetTotalApplicationGroupsOk returns a tuple with the TotalApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplicationGroups

`func (o *AdminFolderResponseModel) SetTotalApplicationGroups(v int32)`

SetTotalApplicationGroups sets TotalApplicationGroups field to given value.

### HasTotalApplicationGroups

`func (o *AdminFolderResponseModel) HasTotalApplicationGroups() bool`

HasTotalApplicationGroups returns a boolean if a field has been set.

### GetTotalDesktopGroups

`func (o *AdminFolderResponseModel) GetTotalDesktopGroups() int32`

GetTotalDesktopGroups returns the TotalDesktopGroups field if non-nil, zero value otherwise.

### GetTotalDesktopGroupsOk

`func (o *AdminFolderResponseModel) GetTotalDesktopGroupsOk() (*int32, bool)`

GetTotalDesktopGroupsOk returns a tuple with the TotalDesktopGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDesktopGroups

`func (o *AdminFolderResponseModel) SetTotalDesktopGroups(v int32)`

SetTotalDesktopGroups sets TotalDesktopGroups field to given value.

### HasTotalDesktopGroups

`func (o *AdminFolderResponseModel) HasTotalDesktopGroups() bool`

HasTotalDesktopGroups returns a boolean if a field has been set.

### GetMetadata

`func (o *AdminFolderResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AdminFolderResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AdminFolderResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AdminFolderResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AdminFolderResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AdminFolderResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


