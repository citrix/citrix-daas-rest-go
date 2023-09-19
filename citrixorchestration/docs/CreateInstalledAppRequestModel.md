# CreateInstalledAppRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLineArguments** | Pointer to **NullableString** | Specifies the command-line arguments to use when launching the executable. Environment variables can be used. | [optional] 
**CommandLineExecutable** | **string** | Specifies the name of the executable file to launch. The full path need not be provided if it&#39;s already in the path. Environment variables can also be used. | 
**WorkingDirectory** | Pointer to **NullableString** | Specifies which working directory the executable is launched from. Environment variables can be used.  | [optional] 

## Methods

### NewCreateInstalledAppRequestModel

`func NewCreateInstalledAppRequestModel(commandLineExecutable string, ) *CreateInstalledAppRequestModel`

NewCreateInstalledAppRequestModel instantiates a new CreateInstalledAppRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstalledAppRequestModelWithDefaults

`func NewCreateInstalledAppRequestModelWithDefaults() *CreateInstalledAppRequestModel`

NewCreateInstalledAppRequestModelWithDefaults instantiates a new CreateInstalledAppRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandLineArguments

`func (o *CreateInstalledAppRequestModel) GetCommandLineArguments() string`

GetCommandLineArguments returns the CommandLineArguments field if non-nil, zero value otherwise.

### GetCommandLineArgumentsOk

`func (o *CreateInstalledAppRequestModel) GetCommandLineArgumentsOk() (*string, bool)`

GetCommandLineArgumentsOk returns a tuple with the CommandLineArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArguments

`func (o *CreateInstalledAppRequestModel) SetCommandLineArguments(v string)`

SetCommandLineArguments sets CommandLineArguments field to given value.

### HasCommandLineArguments

`func (o *CreateInstalledAppRequestModel) HasCommandLineArguments() bool`

HasCommandLineArguments returns a boolean if a field has been set.

### SetCommandLineArgumentsNil

`func (o *CreateInstalledAppRequestModel) SetCommandLineArgumentsNil(b bool)`

 SetCommandLineArgumentsNil sets the value for CommandLineArguments to be an explicit nil

### UnsetCommandLineArguments
`func (o *CreateInstalledAppRequestModel) UnsetCommandLineArguments()`

UnsetCommandLineArguments ensures that no value is present for CommandLineArguments, not even an explicit nil
### GetCommandLineExecutable

`func (o *CreateInstalledAppRequestModel) GetCommandLineExecutable() string`

GetCommandLineExecutable returns the CommandLineExecutable field if non-nil, zero value otherwise.

### GetCommandLineExecutableOk

`func (o *CreateInstalledAppRequestModel) GetCommandLineExecutableOk() (*string, bool)`

GetCommandLineExecutableOk returns a tuple with the CommandLineExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineExecutable

`func (o *CreateInstalledAppRequestModel) SetCommandLineExecutable(v string)`

SetCommandLineExecutable sets CommandLineExecutable field to given value.


### GetWorkingDirectory

`func (o *CreateInstalledAppRequestModel) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *CreateInstalledAppRequestModel) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *CreateInstalledAppRequestModel) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *CreateInstalledAppRequestModel) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### SetWorkingDirectoryNil

`func (o *CreateInstalledAppRequestModel) SetWorkingDirectoryNil(b bool)`

 SetWorkingDirectoryNil sets the value for WorkingDirectory to be an explicit nil

### UnsetWorkingDirectory
`func (o *CreateInstalledAppRequestModel) UnsetWorkingDirectory()`

UnsetWorkingDirectory ensures that no value is present for WorkingDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


