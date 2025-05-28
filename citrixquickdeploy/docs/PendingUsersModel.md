# PendingUsersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProvider** | Pointer to [**IdentityProvider**](IdentityProvider.md) | Identity Provider used to contain information for users | [optional] 
**Users** | Pointer to [**[]PendingUser**](PendingUser.md) |  | [optional] 

## Methods

### NewPendingUsersModel

`func NewPendingUsersModel() *PendingUsersModel`

NewPendingUsersModel instantiates a new PendingUsersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingUsersModelWithDefaults

`func NewPendingUsersModelWithDefaults() *PendingUsersModel`

NewPendingUsersModelWithDefaults instantiates a new PendingUsersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityProvider

`func (o *PendingUsersModel) GetIdentityProvider() IdentityProvider`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *PendingUsersModel) GetIdentityProviderOk() (*IdentityProvider, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *PendingUsersModel) SetIdentityProvider(v IdentityProvider)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *PendingUsersModel) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetUsers

`func (o *PendingUsersModel) GetUsers() []PendingUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PendingUsersModel) GetUsersOk() (*[]PendingUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PendingUsersModel) SetUsers(v []PendingUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PendingUsersModel) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


