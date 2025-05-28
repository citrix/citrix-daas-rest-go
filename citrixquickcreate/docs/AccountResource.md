# AccountResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) | The type of account the resource is associated with | 
**AccountId** | Pointer to **NullableString** | Account Id | [optional] 
**ResourceId** | Pointer to **NullableString** | Resource Id | [optional] 

## Methods

### NewAccountResource

`func NewAccountResource(accountType AccountType, ) *AccountResource`

NewAccountResource instantiates a new AccountResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResourceWithDefaults

`func NewAccountResourceWithDefaults() *AccountResource`

NewAccountResourceWithDefaults instantiates a new AccountResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *AccountResource) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountResource) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountResource) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetAccountId

`func (o *AccountResource) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountResource) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountResource) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountResource) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *AccountResource) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *AccountResource) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetResourceId

`func (o *AccountResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AccountResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AccountResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AccountResource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AccountResource) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AccountResource) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


