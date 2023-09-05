# DelegatedAdminPermissionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the permission. Needs to be made globally unique (perhaps add site id) | 
**Name** | **string** | Name of the permission. | 
**Description** | Pointer to **string** | Description of the permission. | [optional] 
**IsReadOnly** | **bool** | Indicates whether the permission is restricted to read-only operations. Was: ReadOnly | 
**Operations** | **[]string** | List of operations that the permission allows. | 

## Methods

### NewDelegatedAdminPermissionResponseModel

`func NewDelegatedAdminPermissionResponseModel(id string, name string, isReadOnly bool, operations []string, ) *DelegatedAdminPermissionResponseModel`

NewDelegatedAdminPermissionResponseModel instantiates a new DelegatedAdminPermissionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedAdminPermissionResponseModelWithDefaults

`func NewDelegatedAdminPermissionResponseModelWithDefaults() *DelegatedAdminPermissionResponseModel`

NewDelegatedAdminPermissionResponseModelWithDefaults instantiates a new DelegatedAdminPermissionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DelegatedAdminPermissionResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegatedAdminPermissionResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegatedAdminPermissionResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DelegatedAdminPermissionResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DelegatedAdminPermissionResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DelegatedAdminPermissionResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DelegatedAdminPermissionResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelegatedAdminPermissionResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelegatedAdminPermissionResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelegatedAdminPermissionResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *DelegatedAdminPermissionResponseModel) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *DelegatedAdminPermissionResponseModel) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *DelegatedAdminPermissionResponseModel) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.


### GetOperations

`func (o *DelegatedAdminPermissionResponseModel) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *DelegatedAdminPermissionResponseModel) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *DelegatedAdminPermissionResponseModel) SetOperations(v []string)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


