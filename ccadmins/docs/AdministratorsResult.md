# AdministratorsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationToken** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]AdministratorResult**](AdministratorResult.md) |  | [optional] 

## Methods

### NewAdministratorsResult

`func NewAdministratorsResult() *AdministratorsResult`

NewAdministratorsResult instantiates a new AdministratorsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorsResultWithDefaults

`func NewAdministratorsResultWithDefaults() *AdministratorsResult`

NewAdministratorsResultWithDefaults instantiates a new AdministratorsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationToken

`func (o *AdministratorsResult) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *AdministratorsResult) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *AdministratorsResult) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *AdministratorsResult) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *AdministratorsResult) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *AdministratorsResult) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
### GetItems

`func (o *AdministratorsResult) GetItems() []AdministratorResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AdministratorsResult) GetItemsOk() (*[]AdministratorResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AdministratorsResult) SetItems(v []AdministratorResult)`

SetItems sets Items field to given value.

### HasItems

`func (o *AdministratorsResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *AdministratorsResult) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AdministratorsResult) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


