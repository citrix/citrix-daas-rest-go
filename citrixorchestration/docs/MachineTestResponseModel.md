# MachineTestResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Machine** | Pointer to [**MachineTestResponseModelMachine**](MachineTestResponseModelMachine.md) |  | [optional] 
**Status** | Pointer to [**CloudHealthCheckStatus**](CloudHealthCheckStatus.md) |  | [optional] 
**CommandResponse** | Pointer to **string** | CommandResponse. | [optional] 
**ErrorMessage** | Pointer to **string** | ErrorMessage. | [optional] 
**CommandName** | Pointer to **string** | CommandName. | [optional] 
**CategoryName** | Pointer to **string** | CategoryName. | [optional] 
**CreatedBy** | Pointer to **string** | The owner of the test | [optional] 

## Methods

### NewMachineTestResponseModel

`func NewMachineTestResponseModel() *MachineTestResponseModel`

NewMachineTestResponseModel instantiates a new MachineTestResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineTestResponseModelWithDefaults

`func NewMachineTestResponseModelWithDefaults() *MachineTestResponseModel`

NewMachineTestResponseModelWithDefaults instantiates a new MachineTestResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachine

`func (o *MachineTestResponseModel) GetMachine() MachineTestResponseModelMachine`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *MachineTestResponseModel) GetMachineOk() (*MachineTestResponseModelMachine, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *MachineTestResponseModel) SetMachine(v MachineTestResponseModelMachine)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *MachineTestResponseModel) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetStatus

`func (o *MachineTestResponseModel) GetStatus() CloudHealthCheckStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MachineTestResponseModel) GetStatusOk() (*CloudHealthCheckStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MachineTestResponseModel) SetStatus(v CloudHealthCheckStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MachineTestResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCommandResponse

`func (o *MachineTestResponseModel) GetCommandResponse() string`

GetCommandResponse returns the CommandResponse field if non-nil, zero value otherwise.

### GetCommandResponseOk

`func (o *MachineTestResponseModel) GetCommandResponseOk() (*string, bool)`

GetCommandResponseOk returns a tuple with the CommandResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandResponse

`func (o *MachineTestResponseModel) SetCommandResponse(v string)`

SetCommandResponse sets CommandResponse field to given value.

### HasCommandResponse

`func (o *MachineTestResponseModel) HasCommandResponse() bool`

HasCommandResponse returns a boolean if a field has been set.

### GetErrorMessage

`func (o *MachineTestResponseModel) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MachineTestResponseModel) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MachineTestResponseModel) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *MachineTestResponseModel) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCommandName

`func (o *MachineTestResponseModel) GetCommandName() string`

GetCommandName returns the CommandName field if non-nil, zero value otherwise.

### GetCommandNameOk

`func (o *MachineTestResponseModel) GetCommandNameOk() (*string, bool)`

GetCommandNameOk returns a tuple with the CommandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandName

`func (o *MachineTestResponseModel) SetCommandName(v string)`

SetCommandName sets CommandName field to given value.

### HasCommandName

`func (o *MachineTestResponseModel) HasCommandName() bool`

HasCommandName returns a boolean if a field has been set.

### GetCategoryName

`func (o *MachineTestResponseModel) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *MachineTestResponseModel) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *MachineTestResponseModel) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *MachineTestResponseModel) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MachineTestResponseModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MachineTestResponseModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MachineTestResponseModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MachineTestResponseModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


