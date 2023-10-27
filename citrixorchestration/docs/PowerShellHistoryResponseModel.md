# PowerShellHistoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | The title of the PowerShell script history | [optional] 
**ExecutionTime** | Pointer to **NullableString** | The execution time of the PowerShell | [optional] 
**Status** | Pointer to [**PowerShellExecutionStatus**](PowerShellExecutionStatus.md) |  | [optional] 
**TotalCommands** | Pointer to **int32** | The total commands in the execution | [optional] 
**Commands** | Pointer to [**[]PowerShellExecutedCommandModel**](PowerShellExecutedCommandModel.md) | The executed PowerShell commands | [optional] 

## Methods

### NewPowerShellHistoryResponseModel

`func NewPowerShellHistoryResponseModel() *PowerShellHistoryResponseModel`

NewPowerShellHistoryResponseModel instantiates a new PowerShellHistoryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerShellHistoryResponseModelWithDefaults

`func NewPowerShellHistoryResponseModelWithDefaults() *PowerShellHistoryResponseModel`

NewPowerShellHistoryResponseModelWithDefaults instantiates a new PowerShellHistoryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PowerShellHistoryResponseModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PowerShellHistoryResponseModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PowerShellHistoryResponseModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PowerShellHistoryResponseModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PowerShellHistoryResponseModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PowerShellHistoryResponseModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetExecutionTime

`func (o *PowerShellHistoryResponseModel) GetExecutionTime() string`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *PowerShellHistoryResponseModel) GetExecutionTimeOk() (*string, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *PowerShellHistoryResponseModel) SetExecutionTime(v string)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *PowerShellHistoryResponseModel) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### SetExecutionTimeNil

`func (o *PowerShellHistoryResponseModel) SetExecutionTimeNil(b bool)`

 SetExecutionTimeNil sets the value for ExecutionTime to be an explicit nil

### UnsetExecutionTime
`func (o *PowerShellHistoryResponseModel) UnsetExecutionTime()`

UnsetExecutionTime ensures that no value is present for ExecutionTime, not even an explicit nil
### GetStatus

`func (o *PowerShellHistoryResponseModel) GetStatus() PowerShellExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PowerShellHistoryResponseModel) GetStatusOk() (*PowerShellExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PowerShellHistoryResponseModel) SetStatus(v PowerShellExecutionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PowerShellHistoryResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalCommands

`func (o *PowerShellHistoryResponseModel) GetTotalCommands() int32`

GetTotalCommands returns the TotalCommands field if non-nil, zero value otherwise.

### GetTotalCommandsOk

`func (o *PowerShellHistoryResponseModel) GetTotalCommandsOk() (*int32, bool)`

GetTotalCommandsOk returns a tuple with the TotalCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommands

`func (o *PowerShellHistoryResponseModel) SetTotalCommands(v int32)`

SetTotalCommands sets TotalCommands field to given value.

### HasTotalCommands

`func (o *PowerShellHistoryResponseModel) HasTotalCommands() bool`

HasTotalCommands returns a boolean if a field has been set.

### GetCommands

`func (o *PowerShellHistoryResponseModel) GetCommands() []PowerShellExecutedCommandModel`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *PowerShellHistoryResponseModel) GetCommandsOk() (*[]PowerShellExecutedCommandModel, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *PowerShellHistoryResponseModel) SetCommands(v []PowerShellExecutedCommandModel)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *PowerShellHistoryResponseModel) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### SetCommandsNil

`func (o *PowerShellHistoryResponseModel) SetCommandsNil(b bool)`

 SetCommandsNil sets the value for Commands to be an explicit nil

### UnsetCommands
`func (o *PowerShellHistoryResponseModel) UnsetCommands()`

UnsetCommands ensures that no value is present for Commands, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


