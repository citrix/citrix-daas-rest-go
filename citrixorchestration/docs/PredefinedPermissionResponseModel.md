# PredefinedPermissionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the permission. It is usually missed (null or empty). | [optional] 
**GroupId** | **string** | The group id of the permission. | 
**GroupName** | **string** | The group name of the permission. | 
**Id** | **string** | The Id of the permission. It is globally unique. | 
**Name** | **string** | The name of the permission. It is a friendly name of permission. | 
**Operations** | **[]string** | The list of operations of the permission. | 
**IsReadOnly** | **bool** | Whether any of the operations of the permission may change object status. | 

## Methods

### NewPredefinedPermissionResponseModel

`func NewPredefinedPermissionResponseModel(groupId string, groupName string, id string, name string, operations []string, isReadOnly bool, ) *PredefinedPermissionResponseModel`

NewPredefinedPermissionResponseModel instantiates a new PredefinedPermissionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredefinedPermissionResponseModelWithDefaults

`func NewPredefinedPermissionResponseModelWithDefaults() *PredefinedPermissionResponseModel`

NewPredefinedPermissionResponseModelWithDefaults instantiates a new PredefinedPermissionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PredefinedPermissionResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PredefinedPermissionResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PredefinedPermissionResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PredefinedPermissionResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupId

`func (o *PredefinedPermissionResponseModel) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PredefinedPermissionResponseModel) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PredefinedPermissionResponseModel) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *PredefinedPermissionResponseModel) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PredefinedPermissionResponseModel) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PredefinedPermissionResponseModel) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetId

`func (o *PredefinedPermissionResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PredefinedPermissionResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PredefinedPermissionResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PredefinedPermissionResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PredefinedPermissionResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PredefinedPermissionResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetOperations

`func (o *PredefinedPermissionResponseModel) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *PredefinedPermissionResponseModel) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *PredefinedPermissionResponseModel) SetOperations(v []string)`

SetOperations sets Operations field to given value.


### GetIsReadOnly

`func (o *PredefinedPermissionResponseModel) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *PredefinedPermissionResponseModel) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *PredefinedPermissionResponseModel) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


