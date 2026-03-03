# DeleteAdJoinedObjectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccount** | Pointer to **NullableString** | Service account to perform delete with | [optional] 
**ServiceAccountPassword** | Pointer to **NullableString** | Customer&#39;s domain password | [optional] 

## Methods

### NewDeleteAdJoinedObjectModel

`func NewDeleteAdJoinedObjectModel() *DeleteAdJoinedObjectModel`

NewDeleteAdJoinedObjectModel instantiates a new DeleteAdJoinedObjectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAdJoinedObjectModelWithDefaults

`func NewDeleteAdJoinedObjectModelWithDefaults() *DeleteAdJoinedObjectModel`

NewDeleteAdJoinedObjectModelWithDefaults instantiates a new DeleteAdJoinedObjectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccount

`func (o *DeleteAdJoinedObjectModel) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *DeleteAdJoinedObjectModel) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *DeleteAdJoinedObjectModel) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *DeleteAdJoinedObjectModel) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### SetServiceAccountNil

`func (o *DeleteAdJoinedObjectModel) SetServiceAccountNil(b bool)`

 SetServiceAccountNil sets the value for ServiceAccount to be an explicit nil

### UnsetServiceAccount
`func (o *DeleteAdJoinedObjectModel) UnsetServiceAccount()`

UnsetServiceAccount ensures that no value is present for ServiceAccount, not even an explicit nil
### GetServiceAccountPassword

`func (o *DeleteAdJoinedObjectModel) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *DeleteAdJoinedObjectModel) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *DeleteAdJoinedObjectModel) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.

### HasServiceAccountPassword

`func (o *DeleteAdJoinedObjectModel) HasServiceAccountPassword() bool`

HasServiceAccountPassword returns a boolean if a field has been set.

### SetServiceAccountPasswordNil

`func (o *DeleteAdJoinedObjectModel) SetServiceAccountPasswordNil(b bool)`

 SetServiceAccountPasswordNil sets the value for ServiceAccountPassword to be an explicit nil

### UnsetServiceAccountPassword
`func (o *DeleteAdJoinedObjectModel) UnsetServiceAccountPassword()`

UnsetServiceAccountPassword ensures that no value is present for ServiceAccountPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


