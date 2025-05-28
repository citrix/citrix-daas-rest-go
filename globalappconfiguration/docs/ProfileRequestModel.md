# ProfileRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]GroupsModel**](GroupsModel.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewProfileRequestModel

`func NewProfileRequestModel() *ProfileRequestModel`

NewProfileRequestModel instantiates a new ProfileRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileRequestModelWithDefaults

`func NewProfileRequestModelWithDefaults() *ProfileRequestModel`

NewProfileRequestModelWithDefaults instantiates a new ProfileRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProfileRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProfileRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProfileRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProfileRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroups

`func (o *ProfileRequestModel) GetGroups() []GroupsModel`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProfileRequestModel) GetGroupsOk() (*[]GroupsModel, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProfileRequestModel) SetGroups(v []GroupsModel)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ProfileRequestModel) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetName

`func (o *ProfileRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfileRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfileRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProfileRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *ProfileRequestModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProfileRequestModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProfileRequestModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProfileRequestModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUrl

`func (o *ProfileRequestModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProfileRequestModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProfileRequestModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProfileRequestModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


