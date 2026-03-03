# WorkspaceManagedAADModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceAADName** | Pointer to **NullableString** |  | [optional] 
**DirectoryName** | Pointer to **NullableString** |  | [optional] 
**CitrixManaged** | Pointer to **bool** |  | [optional] 
**IsLinked** | Pointer to **bool** |  | [optional] 
**IsProvisioned** | Pointer to **bool** |  | [optional] 
**IdpInstanceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkspaceManagedAADModel

`func NewWorkspaceManagedAADModel() *WorkspaceManagedAADModel`

NewWorkspaceManagedAADModel instantiates a new WorkspaceManagedAADModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceManagedAADModelWithDefaults

`func NewWorkspaceManagedAADModelWithDefaults() *WorkspaceManagedAADModel`

NewWorkspaceManagedAADModelWithDefaults instantiates a new WorkspaceManagedAADModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceAADName

`func (o *WorkspaceManagedAADModel) GetWorkspaceAADName() string`

GetWorkspaceAADName returns the WorkspaceAADName field if non-nil, zero value otherwise.

### GetWorkspaceAADNameOk

`func (o *WorkspaceManagedAADModel) GetWorkspaceAADNameOk() (*string, bool)`

GetWorkspaceAADNameOk returns a tuple with the WorkspaceAADName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAADName

`func (o *WorkspaceManagedAADModel) SetWorkspaceAADName(v string)`

SetWorkspaceAADName sets WorkspaceAADName field to given value.

### HasWorkspaceAADName

`func (o *WorkspaceManagedAADModel) HasWorkspaceAADName() bool`

HasWorkspaceAADName returns a boolean if a field has been set.

### SetWorkspaceAADNameNil

`func (o *WorkspaceManagedAADModel) SetWorkspaceAADNameNil(b bool)`

 SetWorkspaceAADNameNil sets the value for WorkspaceAADName to be an explicit nil

### UnsetWorkspaceAADName
`func (o *WorkspaceManagedAADModel) UnsetWorkspaceAADName()`

UnsetWorkspaceAADName ensures that no value is present for WorkspaceAADName, not even an explicit nil
### GetDirectoryName

`func (o *WorkspaceManagedAADModel) GetDirectoryName() string`

GetDirectoryName returns the DirectoryName field if non-nil, zero value otherwise.

### GetDirectoryNameOk

`func (o *WorkspaceManagedAADModel) GetDirectoryNameOk() (*string, bool)`

GetDirectoryNameOk returns a tuple with the DirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryName

`func (o *WorkspaceManagedAADModel) SetDirectoryName(v string)`

SetDirectoryName sets DirectoryName field to given value.

### HasDirectoryName

`func (o *WorkspaceManagedAADModel) HasDirectoryName() bool`

HasDirectoryName returns a boolean if a field has been set.

### SetDirectoryNameNil

`func (o *WorkspaceManagedAADModel) SetDirectoryNameNil(b bool)`

 SetDirectoryNameNil sets the value for DirectoryName to be an explicit nil

### UnsetDirectoryName
`func (o *WorkspaceManagedAADModel) UnsetDirectoryName()`

UnsetDirectoryName ensures that no value is present for DirectoryName, not even an explicit nil
### GetCitrixManaged

`func (o *WorkspaceManagedAADModel) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *WorkspaceManagedAADModel) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *WorkspaceManagedAADModel) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.

### HasCitrixManaged

`func (o *WorkspaceManagedAADModel) HasCitrixManaged() bool`

HasCitrixManaged returns a boolean if a field has been set.

### GetIsLinked

`func (o *WorkspaceManagedAADModel) GetIsLinked() bool`

GetIsLinked returns the IsLinked field if non-nil, zero value otherwise.

### GetIsLinkedOk

`func (o *WorkspaceManagedAADModel) GetIsLinkedOk() (*bool, bool)`

GetIsLinkedOk returns a tuple with the IsLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinked

`func (o *WorkspaceManagedAADModel) SetIsLinked(v bool)`

SetIsLinked sets IsLinked field to given value.

### HasIsLinked

`func (o *WorkspaceManagedAADModel) HasIsLinked() bool`

HasIsLinked returns a boolean if a field has been set.

### GetIsProvisioned

`func (o *WorkspaceManagedAADModel) GetIsProvisioned() bool`

GetIsProvisioned returns the IsProvisioned field if non-nil, zero value otherwise.

### GetIsProvisionedOk

`func (o *WorkspaceManagedAADModel) GetIsProvisionedOk() (*bool, bool)`

GetIsProvisionedOk returns a tuple with the IsProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvisioned

`func (o *WorkspaceManagedAADModel) SetIsProvisioned(v bool)`

SetIsProvisioned sets IsProvisioned field to given value.

### HasIsProvisioned

`func (o *WorkspaceManagedAADModel) HasIsProvisioned() bool`

HasIsProvisioned returns a boolean if a field has been set.

### GetIdpInstanceId

`func (o *WorkspaceManagedAADModel) GetIdpInstanceId() string`

GetIdpInstanceId returns the IdpInstanceId field if non-nil, zero value otherwise.

### GetIdpInstanceIdOk

`func (o *WorkspaceManagedAADModel) GetIdpInstanceIdOk() (*string, bool)`

GetIdpInstanceIdOk returns a tuple with the IdpInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpInstanceId

`func (o *WorkspaceManagedAADModel) SetIdpInstanceId(v string)`

SetIdpInstanceId sets IdpInstanceId field to given value.

### HasIdpInstanceId

`func (o *WorkspaceManagedAADModel) HasIdpInstanceId() bool`

HasIdpInstanceId returns a boolean if a field has been set.

### SetIdpInstanceIdNil

`func (o *WorkspaceManagedAADModel) SetIdpInstanceIdNil(b bool)`

 SetIdpInstanceIdNil sets the value for IdpInstanceId to be an explicit nil

### UnsetIdpInstanceId
`func (o *WorkspaceManagedAADModel) UnsetIdpInstanceId()`

UnsetIdpInstanceId ensures that no value is present for IdpInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


