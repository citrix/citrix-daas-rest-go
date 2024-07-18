# Accounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AwsEdcAccount**](AwsEdcAccount.md) | Enumerable of Account | [optional] 

## Methods

### NewAccounts

`func NewAccounts() *Accounts`

NewAccounts instantiates a new Accounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsWithDefaults

`func NewAccountsWithDefaults() *Accounts`

NewAccountsWithDefaults instantiates a new Accounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Accounts) GetItems() []AwsEdcAccount`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Accounts) GetItemsOk() (*[]AwsEdcAccount, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Accounts) SetItems(v []AwsEdcAccount)`

SetItems sets Items field to given value.

### HasItems

`func (o *Accounts) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *Accounts) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *Accounts) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


