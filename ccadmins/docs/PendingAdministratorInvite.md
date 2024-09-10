# PendingAdministratorInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RequestorEmail** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Access** | Pointer to [**AdministratorAccessModel**](AdministratorAccessModel.md) |  | [optional] 

## Methods

### NewPendingAdministratorInvite

`func NewPendingAdministratorInvite(email string, ) *PendingAdministratorInvite`

NewPendingAdministratorInvite instantiates a new PendingAdministratorInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingAdministratorInviteWithDefaults

`func NewPendingAdministratorInviteWithDefaults() *PendingAdministratorInvite`

NewPendingAdministratorInviteWithDefaults instantiates a new PendingAdministratorInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PendingAdministratorInvite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PendingAdministratorInvite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PendingAdministratorInvite) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRequestorEmail

`func (o *PendingAdministratorInvite) GetRequestorEmail() string`

GetRequestorEmail returns the RequestorEmail field if non-nil, zero value otherwise.

### GetRequestorEmailOk

`func (o *PendingAdministratorInvite) GetRequestorEmailOk() (*string, bool)`

GetRequestorEmailOk returns a tuple with the RequestorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorEmail

`func (o *PendingAdministratorInvite) SetRequestorEmail(v string)`

SetRequestorEmail sets RequestorEmail field to given value.

### HasRequestorEmail

`func (o *PendingAdministratorInvite) HasRequestorEmail() bool`

HasRequestorEmail returns a boolean if a field has been set.

### SetRequestorEmailNil

`func (o *PendingAdministratorInvite) SetRequestorEmailNil(b bool)`

 SetRequestorEmailNil sets the value for RequestorEmail to be an explicit nil

### UnsetRequestorEmail
`func (o *PendingAdministratorInvite) UnsetRequestorEmail()`

UnsetRequestorEmail ensures that no value is present for RequestorEmail, not even an explicit nil
### GetFirstName

`func (o *PendingAdministratorInvite) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PendingAdministratorInvite) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PendingAdministratorInvite) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PendingAdministratorInvite) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *PendingAdministratorInvite) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *PendingAdministratorInvite) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *PendingAdministratorInvite) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PendingAdministratorInvite) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PendingAdministratorInvite) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PendingAdministratorInvite) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *PendingAdministratorInvite) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *PendingAdministratorInvite) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAccess

`func (o *PendingAdministratorInvite) GetAccess() AdministratorAccessModel`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *PendingAdministratorInvite) GetAccessOk() (*AdministratorAccessModel, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *PendingAdministratorInvite) SetAccess(v AdministratorAccessModel)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *PendingAdministratorInvite) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


