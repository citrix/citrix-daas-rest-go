# PowerShellExecutedCommandModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **NullableString** | The command | [optional] 
**Errors** | Pointer to [**[]PowerShellCommandErrorModel**](PowerShellCommandErrorModel.md) | The errors when the command reported | [optional] 

## Methods

### NewPowerShellExecutedCommandModel

`func NewPowerShellExecutedCommandModel() *PowerShellExecutedCommandModel`

NewPowerShellExecutedCommandModel instantiates a new PowerShellExecutedCommandModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerShellExecutedCommandModelWithDefaults

`func NewPowerShellExecutedCommandModelWithDefaults() *PowerShellExecutedCommandModel`

NewPowerShellExecutedCommandModelWithDefaults instantiates a new PowerShellExecutedCommandModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *PowerShellExecutedCommandModel) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PowerShellExecutedCommandModel) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PowerShellExecutedCommandModel) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *PowerShellExecutedCommandModel) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *PowerShellExecutedCommandModel) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *PowerShellExecutedCommandModel) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetErrors

`func (o *PowerShellExecutedCommandModel) GetErrors() []PowerShellCommandErrorModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PowerShellExecutedCommandModel) GetErrorsOk() (*[]PowerShellCommandErrorModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PowerShellExecutedCommandModel) SetErrors(v []PowerShellCommandErrorModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PowerShellExecutedCommandModel) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *PowerShellExecutedCommandModel) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *PowerShellExecutedCommandModel) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


