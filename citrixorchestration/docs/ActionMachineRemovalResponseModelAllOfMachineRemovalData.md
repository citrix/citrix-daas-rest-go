# ActionMachineRemovalResponseModelAllOfMachineRemovalData

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

### NewActionMachineRemovalResponseModelAllOfMachineRemovalData

`func NewActionMachineRemovalResponseModelAllOfMachineRemovalData() *ActionMachineRemovalResponseModelAllOfMachineRemovalData`

NewActionMachineRemovalResponseModelAllOfMachineRemovalData instantiates a new ActionMachineRemovalResponseModelAllOfMachineRemovalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionMachineRemovalResponseModelAllOfMachineRemovalDataWithDefaults

`func NewActionMachineRemovalResponseModelAllOfMachineRemovalDataWithDefaults() *ActionMachineRemovalResponseModelAllOfMachineRemovalData`

NewActionMachineRemovalResponseModelAllOfMachineRemovalDataWithDefaults instantiates a new ActionMachineRemovalResponseModelAllOfMachineRemovalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdAccountAction

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetAdAccountAction() string`

GetAdAccountAction returns the AdAccountAction field if non-nil, zero value otherwise.

### GetAdAccountActionOk

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetAdAccountActionOk() (*string, bool)`

GetAdAccountActionOk returns a tuple with the AdAccountAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountAction

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) SetAdAccountAction(v string)`

SetAdAccountAction sets AdAccountAction field to given value.

### HasAdAccountAction

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) HasAdAccountAction() bool`

HasAdAccountAction returns a boolean if a field has been set.

### GetSuccessfulAccounts

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetSuccessfulAccounts() []string`

GetSuccessfulAccounts returns the SuccessfulAccounts field if non-nil, zero value otherwise.

### GetSuccessfulAccountsOk

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetSuccessfulAccountsOk() (*[]string, bool)`

GetSuccessfulAccountsOk returns a tuple with the SuccessfulAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulAccounts

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) SetSuccessfulAccounts(v []string)`

SetSuccessfulAccounts sets SuccessfulAccounts field to given value.

### HasSuccessfulAccounts

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) HasSuccessfulAccounts() bool`

HasSuccessfulAccounts returns a boolean if a field has been set.

### GetSuccessfulMachines

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetSuccessfulMachines() []string`

GetSuccessfulMachines returns the SuccessfulMachines field if non-nil, zero value otherwise.

### GetSuccessfulMachinesOk

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetSuccessfulMachinesOk() (*[]string, bool)`

GetSuccessfulMachinesOk returns a tuple with the SuccessfulMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulMachines

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) SetSuccessfulMachines(v []string)`

SetSuccessfulMachines sets SuccessfulMachines field to given value.

### HasSuccessfulMachines

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) HasSuccessfulMachines() bool`

HasSuccessfulMachines returns a boolean if a field has been set.

### GetFailedMachines

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetFailedMachines() []ActionFailedMachineOrAccountResponseModel`

GetFailedMachines returns the FailedMachines field if non-nil, zero value otherwise.

### GetFailedMachinesOk

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetFailedMachinesOk() (*[]ActionFailedMachineOrAccountResponseModel, bool)`

GetFailedMachinesOk returns a tuple with the FailedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMachines

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) SetFailedMachines(v []ActionFailedMachineOrAccountResponseModel)`

SetFailedMachines sets FailedMachines field to given value.

### HasFailedMachines

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) HasFailedMachines() bool`

HasFailedMachines returns a boolean if a field has been set.

### GetFailedAccounts

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetFailedAccounts() []ActionFailedMachineOrAccountResponseModel`

GetFailedAccounts returns the FailedAccounts field if non-nil, zero value otherwise.

### GetFailedAccountsOk

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetFailedAccountsOk() (*[]ActionFailedMachineOrAccountResponseModel, bool)`

GetFailedAccountsOk returns a tuple with the FailedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAccounts

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) SetFailedAccounts(v []ActionFailedMachineOrAccountResponseModel)`

SetFailedAccounts sets FailedAccounts field to given value.

### HasFailedAccounts

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) HasFailedAccounts() bool`

HasFailedAccounts returns a boolean if a field has been set.

### GetVMAction

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetVMAction() string`

GetVMAction returns the VMAction field if non-nil, zero value otherwise.

### GetVMActionOk

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) GetVMActionOk() (*string, bool)`

GetVMActionOk returns a tuple with the VMAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMAction

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) SetVMAction(v string)`

SetVMAction sets VMAction field to given value.

### HasVMAction

`func (o *ActionMachineRemovalResponseModelAllOfMachineRemovalData) HasVMAction() bool`

HasVMAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


