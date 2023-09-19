# ActionMachineRemovalDetailsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdAccountAction** | Pointer to **NullableString** | Active directory account action. | [optional] 
**SuccessfulAccounts** | Pointer to **[]string** | Successful accounts. | [optional] 
**SuccessfulMachines** | Pointer to **[]string** | Successful machines.  | [optional] 
**FailedMachines** | Pointer to [**[]ActionFailedMachineOrAccountResponseModel**](ActionFailedMachineOrAccountResponseModel.md) | Failed machines and Action errors. | [optional] 
**FailedAccounts** | Pointer to [**[]ActionFailedMachineOrAccountResponseModel**](ActionFailedMachineOrAccountResponseModel.md) | Failed accounts and action errors. | [optional] 
**VMAction** | Pointer to **NullableString** | The virtual machine actions.  | [optional] 

## Methods

### NewActionMachineRemovalDetailsResponseModel

`func NewActionMachineRemovalDetailsResponseModel() *ActionMachineRemovalDetailsResponseModel`

NewActionMachineRemovalDetailsResponseModel instantiates a new ActionMachineRemovalDetailsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionMachineRemovalDetailsResponseModelWithDefaults

`func NewActionMachineRemovalDetailsResponseModelWithDefaults() *ActionMachineRemovalDetailsResponseModel`

NewActionMachineRemovalDetailsResponseModelWithDefaults instantiates a new ActionMachineRemovalDetailsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdAccountAction

`func (o *ActionMachineRemovalDetailsResponseModel) GetAdAccountAction() string`

GetAdAccountAction returns the AdAccountAction field if non-nil, zero value otherwise.

### GetAdAccountActionOk

`func (o *ActionMachineRemovalDetailsResponseModel) GetAdAccountActionOk() (*string, bool)`

GetAdAccountActionOk returns a tuple with the AdAccountAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountAction

`func (o *ActionMachineRemovalDetailsResponseModel) SetAdAccountAction(v string)`

SetAdAccountAction sets AdAccountAction field to given value.

### HasAdAccountAction

`func (o *ActionMachineRemovalDetailsResponseModel) HasAdAccountAction() bool`

HasAdAccountAction returns a boolean if a field has been set.

### SetAdAccountActionNil

`func (o *ActionMachineRemovalDetailsResponseModel) SetAdAccountActionNil(b bool)`

 SetAdAccountActionNil sets the value for AdAccountAction to be an explicit nil

### UnsetAdAccountAction
`func (o *ActionMachineRemovalDetailsResponseModel) UnsetAdAccountAction()`

UnsetAdAccountAction ensures that no value is present for AdAccountAction, not even an explicit nil
### GetSuccessfulAccounts

`func (o *ActionMachineRemovalDetailsResponseModel) GetSuccessfulAccounts() []string`

GetSuccessfulAccounts returns the SuccessfulAccounts field if non-nil, zero value otherwise.

### GetSuccessfulAccountsOk

`func (o *ActionMachineRemovalDetailsResponseModel) GetSuccessfulAccountsOk() (*[]string, bool)`

GetSuccessfulAccountsOk returns a tuple with the SuccessfulAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulAccounts

`func (o *ActionMachineRemovalDetailsResponseModel) SetSuccessfulAccounts(v []string)`

SetSuccessfulAccounts sets SuccessfulAccounts field to given value.

### HasSuccessfulAccounts

`func (o *ActionMachineRemovalDetailsResponseModel) HasSuccessfulAccounts() bool`

HasSuccessfulAccounts returns a boolean if a field has been set.

### SetSuccessfulAccountsNil

`func (o *ActionMachineRemovalDetailsResponseModel) SetSuccessfulAccountsNil(b bool)`

 SetSuccessfulAccountsNil sets the value for SuccessfulAccounts to be an explicit nil

### UnsetSuccessfulAccounts
`func (o *ActionMachineRemovalDetailsResponseModel) UnsetSuccessfulAccounts()`

UnsetSuccessfulAccounts ensures that no value is present for SuccessfulAccounts, not even an explicit nil
### GetSuccessfulMachines

`func (o *ActionMachineRemovalDetailsResponseModel) GetSuccessfulMachines() []string`

GetSuccessfulMachines returns the SuccessfulMachines field if non-nil, zero value otherwise.

### GetSuccessfulMachinesOk

`func (o *ActionMachineRemovalDetailsResponseModel) GetSuccessfulMachinesOk() (*[]string, bool)`

GetSuccessfulMachinesOk returns a tuple with the SuccessfulMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulMachines

`func (o *ActionMachineRemovalDetailsResponseModel) SetSuccessfulMachines(v []string)`

SetSuccessfulMachines sets SuccessfulMachines field to given value.

### HasSuccessfulMachines

`func (o *ActionMachineRemovalDetailsResponseModel) HasSuccessfulMachines() bool`

HasSuccessfulMachines returns a boolean if a field has been set.

### SetSuccessfulMachinesNil

`func (o *ActionMachineRemovalDetailsResponseModel) SetSuccessfulMachinesNil(b bool)`

 SetSuccessfulMachinesNil sets the value for SuccessfulMachines to be an explicit nil

### UnsetSuccessfulMachines
`func (o *ActionMachineRemovalDetailsResponseModel) UnsetSuccessfulMachines()`

UnsetSuccessfulMachines ensures that no value is present for SuccessfulMachines, not even an explicit nil
### GetFailedMachines

`func (o *ActionMachineRemovalDetailsResponseModel) GetFailedMachines() []ActionFailedMachineOrAccountResponseModel`

GetFailedMachines returns the FailedMachines field if non-nil, zero value otherwise.

### GetFailedMachinesOk

`func (o *ActionMachineRemovalDetailsResponseModel) GetFailedMachinesOk() (*[]ActionFailedMachineOrAccountResponseModel, bool)`

GetFailedMachinesOk returns a tuple with the FailedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMachines

`func (o *ActionMachineRemovalDetailsResponseModel) SetFailedMachines(v []ActionFailedMachineOrAccountResponseModel)`

SetFailedMachines sets FailedMachines field to given value.

### HasFailedMachines

`func (o *ActionMachineRemovalDetailsResponseModel) HasFailedMachines() bool`

HasFailedMachines returns a boolean if a field has been set.

### SetFailedMachinesNil

`func (o *ActionMachineRemovalDetailsResponseModel) SetFailedMachinesNil(b bool)`

 SetFailedMachinesNil sets the value for FailedMachines to be an explicit nil

### UnsetFailedMachines
`func (o *ActionMachineRemovalDetailsResponseModel) UnsetFailedMachines()`

UnsetFailedMachines ensures that no value is present for FailedMachines, not even an explicit nil
### GetFailedAccounts

`func (o *ActionMachineRemovalDetailsResponseModel) GetFailedAccounts() []ActionFailedMachineOrAccountResponseModel`

GetFailedAccounts returns the FailedAccounts field if non-nil, zero value otherwise.

### GetFailedAccountsOk

`func (o *ActionMachineRemovalDetailsResponseModel) GetFailedAccountsOk() (*[]ActionFailedMachineOrAccountResponseModel, bool)`

GetFailedAccountsOk returns a tuple with the FailedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAccounts

`func (o *ActionMachineRemovalDetailsResponseModel) SetFailedAccounts(v []ActionFailedMachineOrAccountResponseModel)`

SetFailedAccounts sets FailedAccounts field to given value.

### HasFailedAccounts

`func (o *ActionMachineRemovalDetailsResponseModel) HasFailedAccounts() bool`

HasFailedAccounts returns a boolean if a field has been set.

### SetFailedAccountsNil

`func (o *ActionMachineRemovalDetailsResponseModel) SetFailedAccountsNil(b bool)`

 SetFailedAccountsNil sets the value for FailedAccounts to be an explicit nil

### UnsetFailedAccounts
`func (o *ActionMachineRemovalDetailsResponseModel) UnsetFailedAccounts()`

UnsetFailedAccounts ensures that no value is present for FailedAccounts, not even an explicit nil
### GetVMAction

`func (o *ActionMachineRemovalDetailsResponseModel) GetVMAction() string`

GetVMAction returns the VMAction field if non-nil, zero value otherwise.

### GetVMActionOk

`func (o *ActionMachineRemovalDetailsResponseModel) GetVMActionOk() (*string, bool)`

GetVMActionOk returns a tuple with the VMAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMAction

`func (o *ActionMachineRemovalDetailsResponseModel) SetVMAction(v string)`

SetVMAction sets VMAction field to given value.

### HasVMAction

`func (o *ActionMachineRemovalDetailsResponseModel) HasVMAction() bool`

HasVMAction returns a boolean if a field has been set.

### SetVMActionNil

`func (o *ActionMachineRemovalDetailsResponseModel) SetVMActionNil(b bool)`

 SetVMActionNil sets the value for VMAction to be an explicit nil

### UnsetVMAction
`func (o *ActionMachineRemovalDetailsResponseModel) UnsetVMAction()`

UnsetVMAction ensures that no value is present for VMAction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


