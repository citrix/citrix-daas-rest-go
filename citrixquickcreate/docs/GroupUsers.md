# GroupUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]CoreUser**](CoreUser.md) |  | [optional] 

## Methods

### NewGroupUsers

`func NewGroupUsers() *GroupUsers`

NewGroupUsers instantiates a new GroupUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUsersWithDefaults

`func NewGroupUsersWithDefaults() *GroupUsers`

NewGroupUsersWithDefaults instantiates a new GroupUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GroupUsers) GetUsers() []CoreUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupUsers) GetUsersOk() (*[]CoreUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupUsers) SetUsers(v []CoreUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupUsers) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *GroupUsers) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *GroupUsers) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


