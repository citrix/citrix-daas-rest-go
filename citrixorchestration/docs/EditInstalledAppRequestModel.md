# EditInstalledAppRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLineArguments** | Pointer to **NullableString** | The command-line arguments to use when launching the executable. Environment variables can be used. If not specified, the value will not be changed. | [optional] 
**CommandLineExecutable** | Pointer to **NullableString** | The name of the executable file to launch. The full path need not be provided if it&#39;s already in the path. Environment variables can also be used. If not specified, the value will not be changed. | [optional] 
**WorkingDirectory** | Pointer to **NullableString** | The working directory which the executable is launched from. Environment variables can be used.  If not specified, the value will not be changed. | [optional] 

## Methods

### NewEditInstalledAppRequestModel

`func NewEditInstalledAppRequestModel() *EditInstalledAppRequestModel`

NewEditInstalledAppRequestModel instantiates a new EditInstalledAppRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditInstalledAppRequestModelWithDefaults

`func NewEditInstalledAppRequestModelWithDefaults() *EditInstalledAppRequestModel`

NewEditInstalledAppRequestModelWithDefaults instantiates a new EditInstalledAppRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandLineArguments

`func (o *EditInstalledAppRequestModel) GetCommandLineArguments() string`

GetCommandLineArguments returns the CommandLineArguments field if non-nil, zero value otherwise.

### GetCommandLineArgumentsOk

`func (o *EditInstalledAppRequestModel) GetCommandLineArgumentsOk() (*string, bool)`

GetCommandLineArgumentsOk returns a tuple with the CommandLineArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArguments

`func (o *EditInstalledAppRequestModel) SetCommandLineArguments(v string)`

SetCommandLineArguments sets CommandLineArguments field to given value.

### HasCommandLineArguments

`func (o *EditInstalledAppRequestModel) HasCommandLineArguments() bool`

HasCommandLineArguments returns a boolean if a field has been set.

### SetCommandLineArgumentsNil

`func (o *EditInstalledAppRequestModel) SetCommandLineArgumentsNil(b bool)`

 SetCommandLineArgumentsNil sets the value for CommandLineArguments to be an explicit nil

### UnsetCommandLineArguments
`func (o *EditInstalledAppRequestModel) UnsetCommandLineArguments()`

UnsetCommandLineArguments ensures that no value is present for CommandLineArguments, not even an explicit nil
### GetCommandLineExecutable

`func (o *EditInstalledAppRequestModel) GetCommandLineExecutable() string`

GetCommandLineExecutable returns the CommandLineExecutable field if non-nil, zero value otherwise.

### GetCommandLineExecutableOk

`func (o *EditInstalledAppRequestModel) GetCommandLineExecutableOk() (*string, bool)`

GetCommandLineExecutableOk returns a tuple with the CommandLineExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineExecutable

`func (o *EditInstalledAppRequestModel) SetCommandLineExecutable(v string)`

SetCommandLineExecutable sets CommandLineExecutable field to given value.

### HasCommandLineExecutable

`func (o *EditInstalledAppRequestModel) HasCommandLineExecutable() bool`

HasCommandLineExecutable returns a boolean if a field has been set.

### SetCommandLineExecutableNil

`func (o *EditInstalledAppRequestModel) SetCommandLineExecutableNil(b bool)`

 SetCommandLineExecutableNil sets the value for CommandLineExecutable to be an explicit nil

### UnsetCommandLineExecutable
`func (o *EditInstalledAppRequestModel) UnsetCommandLineExecutable()`

UnsetCommandLineExecutable ensures that no value is present for CommandLineExecutable, not even an explicit nil
### GetWorkingDirectory

`func (o *EditInstalledAppRequestModel) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *EditInstalledAppRequestModel) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *EditInstalledAppRequestModel) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *EditInstalledAppRequestModel) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### SetWorkingDirectoryNil

`func (o *EditInstalledAppRequestModel) SetWorkingDirectoryNil(b bool)`

 SetWorkingDirectoryNil sets the value for WorkingDirectory to be an explicit nil

### UnsetWorkingDirectory
`func (o *EditInstalledAppRequestModel) UnsetWorkingDirectory()`

UnsetWorkingDirectory ensures that no value is present for WorkingDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


