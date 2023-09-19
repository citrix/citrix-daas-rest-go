# ActionMachineCreationDetailsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdAccountAction** | Pointer to **NullableString** | Active directory account action. | [optional] 
**SuccessfulAccounts** | Pointer to **[]string** | Successful accounts. | [optional] 
**SuccessfulMachines** | Pointer to **[]string** | Successful machines.  | [optional] 
**FailedMachines** | Pointer to [**[]ActionFailedMachineOrAccountResponseModel**](ActionFailedMachineOrAccountResponseModel.md) | Failed machines and Action errors. | [optional] 
**FailedAccounts** | Pointer to [**[]ActionFailedMachineOrAccountResponseModel**](ActionFailedMachineOrAccountResponseModel.md) | Failed accounts and action errors. | [optional] 
**MachineRequested** | Pointer to **int32** | The machine requested | [optional] 

## Methods

### NewActionMachineCreationDetailsResponseModel

`func NewActionMachineCreationDetailsResponseModel() *ActionMachineCreationDetailsResponseModel`

NewActionMachineCreationDetailsResponseModel instantiates a new ActionMachineCreationDetailsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionMachineCreationDetailsResponseModelWithDefaults

`func NewActionMachineCreationDetailsResponseModelWithDefaults() *ActionMachineCreationDetailsResponseModel`

NewActionMachineCreationDetailsResponseModelWithDefaults instantiates a new ActionMachineCreationDetailsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdAccountAction

`func (o *ActionMachineCreationDetailsResponseModel) GetAdAccountAction() string`

GetAdAccountAction returns the AdAccountAction field if non-nil, zero value otherwise.

### GetAdAccountActionOk

`func (o *ActionMachineCreationDetailsResponseModel) GetAdAccountActionOk() (*string, bool)`

GetAdAccountActionOk returns a tuple with the AdAccountAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountAction

`func (o *ActionMachineCreationDetailsResponseModel) SetAdAccountAction(v string)`

SetAdAccountAction sets AdAccountAction field to given value.

### HasAdAccountAction

`func (o *ActionMachineCreationDetailsResponseModel) HasAdAccountAction() bool`

HasAdAccountAction returns a boolean if a field has been set.

### SetAdAccountActionNil

`func (o *ActionMachineCreationDetailsResponseModel) SetAdAccountActionNil(b bool)`

 SetAdAccountActionNil sets the value for AdAccountAction to be an explicit nil

### UnsetAdAccountAction
`func (o *ActionMachineCreationDetailsResponseModel) UnsetAdAccountAction()`

UnsetAdAccountAction ensures that no value is present for AdAccountAction, not even an explicit nil
### GetSuccessfulAccounts

`func (o *ActionMachineCreationDetailsResponseModel) GetSuccessfulAccounts() []string`

GetSuccessfulAccounts returns the SuccessfulAccounts field if non-nil, zero value otherwise.

### GetSuccessfulAccountsOk

`func (o *ActionMachineCreationDetailsResponseModel) GetSuccessfulAccountsOk() (*[]string, bool)`

GetSuccessfulAccountsOk returns a tuple with the SuccessfulAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulAccounts

`func (o *ActionMachineCreationDetailsResponseModel) SetSuccessfulAccounts(v []string)`

SetSuccessfulAccounts sets SuccessfulAccounts field to given value.

### HasSuccessfulAccounts

`func (o *ActionMachineCreationDetailsResponseModel) HasSuccessfulAccounts() bool`

HasSuccessfulAccounts returns a boolean if a field has been set.

### SetSuccessfulAccountsNil

`func (o *ActionMachineCreationDetailsResponseModel) SetSuccessfulAccountsNil(b bool)`

 SetSuccessfulAccountsNil sets the value for SuccessfulAccounts to be an explicit nil

### UnsetSuccessfulAccounts
`func (o *ActionMachineCreationDetailsResponseModel) UnsetSuccessfulAccounts()`

UnsetSuccessfulAccounts ensures that no value is present for SuccessfulAccounts, not even an explicit nil
### GetSuccessfulMachines

`func (o *ActionMachineCreationDetailsResponseModel) GetSuccessfulMachines() []string`

GetSuccessfulMachines returns the SuccessfulMachines field if non-nil, zero value otherwise.

### GetSuccessfulMachinesOk

`func (o *ActionMachineCreationDetailsResponseModel) GetSuccessfulMachinesOk() (*[]string, bool)`

GetSuccessfulMachinesOk returns a tuple with the SuccessfulMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulMachines

`func (o *ActionMachineCreationDetailsResponseModel) SetSuccessfulMachines(v []string)`

SetSuccessfulMachines sets SuccessfulMachines field to given value.

### HasSuccessfulMachines

`func (o *ActionMachineCreationDetailsResponseModel) HasSuccessfulMachines() bool`

HasSuccessfulMachines returns a boolean if a field has been set.

### SetSuccessfulMachinesNil

`func (o *ActionMachineCreationDetailsResponseModel) SetSuccessfulMachinesNil(b bool)`

 SetSuccessfulMachinesNil sets the value for SuccessfulMachines to be an explicit nil

### UnsetSuccessfulMachines
`func (o *ActionMachineCreationDetailsResponseModel) UnsetSuccessfulMachines()`

UnsetSuccessfulMachines ensures that no value is present for SuccessfulMachines, not even an explicit nil
### GetFailedMachines

`func (o *ActionMachineCreationDetailsResponseModel) GetFailedMachines() []ActionFailedMachineOrAccountResponseModel`

GetFailedMachines returns the FailedMachines field if non-nil, zero value otherwise.

### GetFailedMachinesOk

`func (o *ActionMachineCreationDetailsResponseModel) GetFailedMachinesOk() (*[]ActionFailedMachineOrAccountResponseModel, bool)`

GetFailedMachinesOk returns a tuple with the FailedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMachines

`func (o *ActionMachineCreationDetailsResponseModel) SetFailedMachines(v []ActionFailedMachineOrAccountResponseModel)`

SetFailedMachines sets FailedMachines field to given value.

### HasFailedMachines

`func (o *ActionMachineCreationDetailsResponseModel) HasFailedMachines() bool`

HasFailedMachines returns a boolean if a field has been set.

### SetFailedMachinesNil

`func (o *ActionMachineCreationDetailsResponseModel) SetFailedMachinesNil(b bool)`

 SetFailedMachinesNil sets the value for FailedMachines to be an explicit nil

### UnsetFailedMachines
`func (o *ActionMachineCreationDetailsResponseModel) UnsetFailedMachines()`

UnsetFailedMachines ensures that no value is present for FailedMachines, not even an explicit nil
### GetFailedAccounts

`func (o *ActionMachineCreationDetailsResponseModel) GetFailedAccounts() []ActionFailedMachineOrAccountResponseModel`

GetFailedAccounts returns the FailedAccounts field if non-nil, zero value otherwise.

### GetFailedAccountsOk

`func (o *ActionMachineCreationDetailsResponseModel) GetFailedAccountsOk() (*[]ActionFailedMachineOrAccountResponseModel, bool)`

GetFailedAccountsOk returns a tuple with the FailedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAccounts

`func (o *ActionMachineCreationDetailsResponseModel) SetFailedAccounts(v []ActionFailedMachineOrAccountResponseModel)`

SetFailedAccounts sets FailedAccounts field to given value.

### HasFailedAccounts

`func (o *ActionMachineCreationDetailsResponseModel) HasFailedAccounts() bool`

HasFailedAccounts returns a boolean if a field has been set.

### SetFailedAccountsNil

`func (o *ActionMachineCreationDetailsResponseModel) SetFailedAccountsNil(b bool)`

 SetFailedAccountsNil sets the value for FailedAccounts to be an explicit nil

### UnsetFailedAccounts
`func (o *ActionMachineCreationDetailsResponseModel) UnsetFailedAccounts()`

UnsetFailedAccounts ensures that no value is present for FailedAccounts, not even an explicit nil
### GetMachineRequested

`func (o *ActionMachineCreationDetailsResponseModel) GetMachineRequested() int32`

GetMachineRequested returns the MachineRequested field if non-nil, zero value otherwise.

### GetMachineRequestedOk

`func (o *ActionMachineCreationDetailsResponseModel) GetMachineRequestedOk() (*int32, bool)`

GetMachineRequestedOk returns a tuple with the MachineRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineRequested

`func (o *ActionMachineCreationDetailsResponseModel) SetMachineRequested(v int32)`

SetMachineRequested sets MachineRequested field to given value.

### HasMachineRequested

`func (o *ActionMachineCreationDetailsResponseModel) HasMachineRequested() bool`

HasMachineRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


