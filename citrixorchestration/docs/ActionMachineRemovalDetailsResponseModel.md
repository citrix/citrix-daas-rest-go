# ActionMachineRemovalDetailsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdAccountAction** | Pointer to **string** | Active directory account action. | [optional] 
**SuccessfulAccounts** | Pointer to **[]string** | Successful accounts. | [optional] 
**SuccessfulMachines** | Pointer to **[]string** | Successful machines.  | [optional] 
**FailedMachines** | Pointer to [**[]ActionFailedMachineOrAccountResponseModel**](ActionFailedMachineOrAccountResponseModel.md) | Failed machines and Action errors. | [optional] 
**FailedAccounts** | Pointer to [**[]ActionFailedMachineOrAccountResponseModel**](ActionFailedMachineOrAccountResponseModel.md) | Failed accounts and action errors. | [optional] 
**VMAction** | Pointer to **string** | The virtual machine actions.  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


