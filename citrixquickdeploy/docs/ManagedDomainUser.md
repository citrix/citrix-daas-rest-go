# ManagedDomainUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] [readonly] 
**LastName** | Pointer to **NullableString** |  | [optional] [readonly] 
**IsAssigned** | Pointer to **bool** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**CustomUser** | Pointer to **bool** |  | [optional] [readonly] 
**State** | Pointer to [**DomainUserState**](DomainUserState.md) |  | [optional] 

## Methods

### NewManagedDomainUser

`func NewManagedDomainUser() *ManagedDomainUser`

NewManagedDomainUser instantiates a new ManagedDomainUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDomainUserWithDefaults

`func NewManagedDomainUserWithDefaults() *ManagedDomainUser`

NewManagedDomainUserWithDefaults instantiates a new ManagedDomainUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ManagedDomainUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ManagedDomainUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ManagedDomainUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ManagedDomainUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ManagedDomainUser) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ManagedDomainUser) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetFirstName

`func (o *ManagedDomainUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ManagedDomainUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ManagedDomainUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ManagedDomainUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ManagedDomainUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ManagedDomainUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ManagedDomainUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ManagedDomainUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ManagedDomainUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ManagedDomainUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ManagedDomainUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ManagedDomainUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetIsAssigned

`func (o *ManagedDomainUser) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *ManagedDomainUser) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *ManagedDomainUser) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *ManagedDomainUser) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetIsExternal

`func (o *ManagedDomainUser) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *ManagedDomainUser) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *ManagedDomainUser) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *ManagedDomainUser) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetCustomUser

`func (o *ManagedDomainUser) GetCustomUser() bool`

GetCustomUser returns the CustomUser field if non-nil, zero value otherwise.

### GetCustomUserOk

`func (o *ManagedDomainUser) GetCustomUserOk() (*bool, bool)`

GetCustomUserOk returns a tuple with the CustomUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUser

`func (o *ManagedDomainUser) SetCustomUser(v bool)`

SetCustomUser sets CustomUser field to given value.

### HasCustomUser

`func (o *ManagedDomainUser) HasCustomUser() bool`

HasCustomUser returns a boolean if a field has been set.

### GetState

`func (o *ManagedDomainUser) GetState() DomainUserState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagedDomainUser) GetStateOk() (*DomainUserState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagedDomainUser) SetState(v DomainUserState)`

SetState sets State field to given value.

### HasState

`func (o *ManagedDomainUser) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


