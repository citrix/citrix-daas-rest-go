# ModelingResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets or sets Name | [optional] 
**User** | Pointer to **string** | Gets or sets User | [optional] 
**Machine** | Pointer to **string** | Gets or sets Machine | [optional] 
**Result** | Pointer to [**ModelingResponseContractResult**](ModelingResponseContractResult.md) |  | [optional] 

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

### GetResult

`func (o *ModelingResponseContract) GetResult() ModelingResponseContractResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelingResponseContract) GetResultOk() (*ModelingResponseContractResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelingResponseContract) SetResult(v ModelingResponseContractResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelingResponseContract) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


