# ApplicationFolderResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the application folder. Used to be: Uuid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | 
**Uid** | Pointer to **NullableInt32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED. Use Id. | [optional] 
**Children** | [**[]RefResponseModel**](RefResponseModel.md) | Child folders. | 
**Name** | **string** | Name of the folder. | 
**Parent** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**Path** | **string** | Full path to the folder. Used to be: FullName | 
**TotalApplications** | **int32** | Total number of applications in the folder. | 
**TotalMachineCatalogs** | Pointer to **int32** | Total number of machine catalogs in the folder. | [optional] 
**TotalDesktopGroups** | Pointer to **int32** | Total number of delivery groups in the folder. | [optional] 
**TotalApplicationGroups** | Pointer to **int32** | Total number of application groups in the folder. | [optional] 

## Methods

### NewApplicationFolderResponseModel

`func NewApplicationFolderResponseModel(id string, children []RefResponseModel, name string, path string, totalApplications int32, ) *ApplicationFolderResponseModel`

NewApplicationFolderResponseModel instantiates a new ApplicationFolderResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFolderResponseModelWithDefaults

`func NewApplicationFolderResponseModelWithDefaults() *ApplicationFolderResponseModel`

NewApplicationFolderResponseModelWithDefaults instantiates a new ApplicationFolderResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationFolderResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationFolderResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationFolderResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *ApplicationFolderResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApplicationFolderResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApplicationFolderResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApplicationFolderResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ApplicationFolderResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ApplicationFolderResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetChildren

`func (o *ApplicationFolderResponseModel) GetChildren() []RefResponseModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ApplicationFolderResponseModel) GetChildrenOk() (*[]RefResponseModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ApplicationFolderResponseModel) SetChildren(v []RefResponseModel)`

SetChildren sets Children field to given value.


### GetName

`func (o *ApplicationFolderResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationFolderResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationFolderResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *ApplicationFolderResponseModel) GetParent() RefResponseModel`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ApplicationFolderResponseModel) GetParentOk() (*RefResponseModel, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ApplicationFolderResponseModel) SetParent(v RefResponseModel)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ApplicationFolderResponseModel) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPath

`func (o *ApplicationFolderResponseModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApplicationFolderResponseModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApplicationFolderResponseModel) SetPath(v string)`

SetPath sets Path field to given value.


### GetTotalApplications

`func (o *ApplicationFolderResponseModel) GetTotalApplications() int32`

GetTotalApplications returns the TotalApplications field if non-nil, zero value otherwise.

### GetTotalApplicationsOk

`func (o *ApplicationFolderResponseModel) GetTotalApplicationsOk() (*int32, bool)`

GetTotalApplicationsOk returns a tuple with the TotalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplications

`func (o *ApplicationFolderResponseModel) SetTotalApplications(v int32)`

SetTotalApplications sets TotalApplications field to given value.


### GetTotalMachineCatalogs

`func (o *ApplicationFolderResponseModel) GetTotalMachineCatalogs() int32`

GetTotalMachineCatalogs returns the TotalMachineCatalogs field if non-nil, zero value otherwise.

### GetTotalMachineCatalogsOk

`func (o *ApplicationFolderResponseModel) GetTotalMachineCatalogsOk() (*int32, bool)`

GetTotalMachineCatalogsOk returns a tuple with the TotalMachineCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachineCatalogs

`func (o *ApplicationFolderResponseModel) SetTotalMachineCatalogs(v int32)`

SetTotalMachineCatalogs sets TotalMachineCatalogs field to given value.

### HasTotalMachineCatalogs

`func (o *ApplicationFolderResponseModel) HasTotalMachineCatalogs() bool`

HasTotalMachineCatalogs returns a boolean if a field has been set.

### GetTotalDesktopGroups

`func (o *ApplicationFolderResponseModel) GetTotalDesktopGroups() int32`

GetTotalDesktopGroups returns the TotalDesktopGroups field if non-nil, zero value otherwise.

### GetTotalDesktopGroupsOk

`func (o *ApplicationFolderResponseModel) GetTotalDesktopGroupsOk() (*int32, bool)`

GetTotalDesktopGroupsOk returns a tuple with the TotalDesktopGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDesktopGroups

`func (o *ApplicationFolderResponseModel) SetTotalDesktopGroups(v int32)`

SetTotalDesktopGroups sets TotalDesktopGroups field to given value.

### HasTotalDesktopGroups

`func (o *ApplicationFolderResponseModel) HasTotalDesktopGroups() bool`

HasTotalDesktopGroups returns a boolean if a field has been set.

### GetTotalApplicationGroups

`func (o *ApplicationFolderResponseModel) GetTotalApplicationGroups() int32`

GetTotalApplicationGroups returns the TotalApplicationGroups field if non-nil, zero value otherwise.

### GetTotalApplicationGroupsOk

`func (o *ApplicationFolderResponseModel) GetTotalApplicationGroupsOk() (*int32, bool)`

GetTotalApplicationGroupsOk returns a tuple with the TotalApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplicationGroups

`func (o *ApplicationFolderResponseModel) SetTotalApplicationGroups(v int32)`

SetTotalApplicationGroups sets TotalApplicationGroups field to given value.

### HasTotalApplicationGroups

`func (o *ApplicationFolderResponseModel) HasTotalApplicationGroups() bool`

HasTotalApplicationGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


