# ModelingResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets Name | [optional] 
**User** | Pointer to **NullableString** | Gets or sets User | [optional] 
**Machine** | Pointer to **NullableString** | Gets or sets Machine | [optional] 
**Result** | Pointer to [**ModelingResultContract**](ModelingResultContract.md) |  | [optional] 

## Methods

### NewModelingResponseContract

`func NewModelingResponseContract() *ModelingResponseContract`

NewModelingResponseContract instantiates a new ModelingResponseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelingResponseContractWithDefaults

`func NewModelingResponseContractWithDefaults() *ModelingResponseContract`

NewModelingResponseContractWithDefaults instantiates a new ModelingResponseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelingResponseContract) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelingResponseContract) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelingResponseContract) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelingResponseContract) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelingResponseContract) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelingResponseContract) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUser

`func (o *ModelingResponseContract) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelingResponseContract) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelingResponseContract) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelingResponseContract) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ModelingResponseContract) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ModelingResponseContract) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetMachine

`func (o *ModelingResponseContract) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *ModelingResponseContract) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *ModelingResponseContract) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *ModelingResponseContract) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### SetMachineNil

`func (o *ModelingResponseContract) SetMachineNil(b bool)`

 SetMachineNil sets the value for Machine to be an explicit nil

### UnsetMachine
`func (o *ModelingResponseContract) UnsetMachine()`

UnsetMachine ensures that no value is present for Machine, not even an explicit nil
### GetResult

`func (o *ModelingResponseContract) GetResult() ModelingResultContract`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelingResponseContract) GetResultOk() (*ModelingResultContract, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelingResponseContract) SetResult(v ModelingResultContract)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelingResponseContract) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


