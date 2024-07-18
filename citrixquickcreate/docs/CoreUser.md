# CoreUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** | AAD User ID for which we want to assign/revoke license | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | AAD UPN for provided User | [optional] 

## Methods

### NewCoreUser

`func NewCoreUser() *CoreUser`

NewCoreUser instantiates a new CoreUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreUserWithDefaults

`func NewCoreUserWithDefaults() *CoreUser`

NewCoreUserWithDefaults instantiates a new CoreUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CoreUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CoreUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CoreUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CoreUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CoreUser) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CoreUser) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserPrincipalName

`func (o *CoreUser) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *CoreUser) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *CoreUser) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *CoreUser) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *CoreUser) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *CoreUser) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


