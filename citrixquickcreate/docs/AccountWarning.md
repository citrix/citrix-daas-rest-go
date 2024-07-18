# AccountWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warning** | Pointer to **NullableString** | The warning message | [optional] 
**IsError** | Pointer to **bool** | Indicates if the warning should be treated as an error | [optional] 
**IsDismissible** | Pointer to **bool** | Indicates if the warning message can be cleared out by the user | [optional] 

## Methods

### NewAccountWarning

`func NewAccountWarning() *AccountWarning`

NewAccountWarning instantiates a new AccountWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWarningWithDefaults

`func NewAccountWarningWithDefaults() *AccountWarning`

NewAccountWarningWithDefaults instantiates a new AccountWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarning

`func (o *AccountWarning) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *AccountWarning) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *AccountWarning) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *AccountWarning) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### SetWarningNil

`func (o *AccountWarning) SetWarningNil(b bool)`

 SetWarningNil sets the value for Warning to be an explicit nil

### UnsetWarning
`func (o *AccountWarning) UnsetWarning()`

UnsetWarning ensures that no value is present for Warning, not even an explicit nil
### GetIsError

`func (o *AccountWarning) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *AccountWarning) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *AccountWarning) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *AccountWarning) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetIsDismissible

`func (o *AccountWarning) GetIsDismissible() bool`

GetIsDismissible returns the IsDismissible field if non-nil, zero value otherwise.

### GetIsDismissibleOk

`func (o *AccountWarning) GetIsDismissibleOk() (*bool, bool)`

GetIsDismissibleOk returns a tuple with the IsDismissible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDismissible

`func (o *AccountWarning) SetIsDismissible(v bool)`

SetIsDismissible sets IsDismissible field to given value.

### HasIsDismissible

`func (o *AccountWarning) HasIsDismissible() bool`

HasIsDismissible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


