# StartMenuApplicationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLineArguments** | Pointer to **NullableString** | Command line arguments. | [optional] 
**CommandLineExecutable** | **string** | Command line executable. | 
**Description** | Pointer to **NullableString** | Description. | [optional] 
**DisplayName** | **string** | Display name. | 
**ShortcutPath** | **string** | Path to the shortcut within the start menu. | 
**WorkingDirectory** | Pointer to **NullableString** | Working directory. | [optional] 
**Machine** | [**RefResponseModel**](RefResponseModel.md) |  | 

## Methods

### NewStartMenuApplicationResponseModel

`func NewStartMenuApplicationResponseModel(commandLineExecutable string, displayName string, shortcutPath string, machine RefResponseModel, ) *StartMenuApplicationResponseModel`

NewStartMenuApplicationResponseModel instantiates a new StartMenuApplicationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartMenuApplicationResponseModelWithDefaults

`func NewStartMenuApplicationResponseModelWithDefaults() *StartMenuApplicationResponseModel`

NewStartMenuApplicationResponseModelWithDefaults instantiates a new StartMenuApplicationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandLineArguments

`func (o *StartMenuApplicationResponseModel) GetCommandLineArguments() string`

GetCommandLineArguments returns the CommandLineArguments field if non-nil, zero value otherwise.

### GetCommandLineArgumentsOk

`func (o *StartMenuApplicationResponseModel) GetCommandLineArgumentsOk() (*string, bool)`

GetCommandLineArgumentsOk returns a tuple with the CommandLineArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArguments

`func (o *StartMenuApplicationResponseModel) SetCommandLineArguments(v string)`

SetCommandLineArguments sets CommandLineArguments field to given value.

### HasCommandLineArguments

`func (o *StartMenuApplicationResponseModel) HasCommandLineArguments() bool`

HasCommandLineArguments returns a boolean if a field has been set.

### SetCommandLineArgumentsNil

`func (o *StartMenuApplicationResponseModel) SetCommandLineArgumentsNil(b bool)`

 SetCommandLineArgumentsNil sets the value for CommandLineArguments to be an explicit nil

### UnsetCommandLineArguments
`func (o *StartMenuApplicationResponseModel) UnsetCommandLineArguments()`

UnsetCommandLineArguments ensures that no value is present for CommandLineArguments, not even an explicit nil
### GetCommandLineExecutable

`func (o *StartMenuApplicationResponseModel) GetCommandLineExecutable() string`

GetCommandLineExecutable returns the CommandLineExecutable field if non-nil, zero value otherwise.

### GetCommandLineExecutableOk

`func (o *StartMenuApplicationResponseModel) GetCommandLineExecutableOk() (*string, bool)`

GetCommandLineExecutableOk returns a tuple with the CommandLineExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineExecutable

`func (o *StartMenuApplicationResponseModel) SetCommandLineExecutable(v string)`

SetCommandLineExecutable sets CommandLineExecutable field to given value.


### GetDescription

`func (o *StartMenuApplicationResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StartMenuApplicationResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StartMenuApplicationResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StartMenuApplicationResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StartMenuApplicationResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StartMenuApplicationResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *StartMenuApplicationResponseModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *StartMenuApplicationResponseModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *StartMenuApplicationResponseModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetShortcutPath

`func (o *StartMenuApplicationResponseModel) GetShortcutPath() string`

GetShortcutPath returns the ShortcutPath field if non-nil, zero value otherwise.

### GetShortcutPathOk

`func (o *StartMenuApplicationResponseModel) GetShortcutPathOk() (*string, bool)`

GetShortcutPathOk returns a tuple with the ShortcutPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutPath

`func (o *StartMenuApplicationResponseModel) SetShortcutPath(v string)`

SetShortcutPath sets ShortcutPath field to given value.


### GetWorkingDirectory

`func (o *StartMenuApplicationResponseModel) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *StartMenuApplicationResponseModel) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *StartMenuApplicationResponseModel) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *StartMenuApplicationResponseModel) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### SetWorkingDirectoryNil

`func (o *StartMenuApplicationResponseModel) SetWorkingDirectoryNil(b bool)`

 SetWorkingDirectoryNil sets the value for WorkingDirectory to be an explicit nil

### UnsetWorkingDirectory
`func (o *StartMenuApplicationResponseModel) UnsetWorkingDirectory()`

UnsetWorkingDirectory ensures that no value is present for WorkingDirectory, not even an explicit nil
### GetMachine

`func (o *StartMenuApplicationResponseModel) GetMachine() RefResponseModel`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *StartMenuApplicationResponseModel) GetMachineOk() (*RefResponseModel, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *StartMenuApplicationResponseModel) SetMachine(v RefResponseModel)`

SetMachine sets Machine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


