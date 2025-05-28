# UserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Identity of user-level AD object | [optional] [readonly] 
**Sid** | **string** | SID of user-level AD object | [readonly] 
**Name** | **string** | Distingushed name of user-level AD object | [readonly] 
**Type** | **string** | Type of user-level AD object | [readonly] 
**Description** | Pointer to **string** | Description of user-level AD object | [optional] 
**SiteId** | **int64** | Identity of site to which this user-level AD object belongs | [readonly] 
**Enabled** | Pointer to **bool** | If this user-level AD objectis enabled | [optional] 
**Priority** | Pointer to **int64** | Priority of user-level AD object | [optional] 

## Methods

### NewUserModel

`func NewUserModel(sid string, name string, type_ string, siteId int64, ) *UserModel`

NewUserModel instantiates a new UserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelWithDefaults

`func NewUserModelWithDefaults() *UserModel`

NewUserModelWithDefaults instantiates a new UserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSid

`func (o *UserModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *UserModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *UserModel) SetSid(v string)`

SetSid sets Sid field to given value.


### GetName

`func (o *UserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UserModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserModel) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *UserModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSiteId

`func (o *UserModel) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UserModel) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UserModel) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.


### GetEnabled

`func (o *UserModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *UserModel) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UserModel) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UserModel) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UserModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


