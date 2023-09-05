# CreateApplicationRequestModelInstalledAppProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLineArguments** | Pointer to **string** | Specifies the command-line arguments to use when launching the executable. Environment variables can be used. | [optional] 
**CommandLineExecutable** | **string** | Specifies the name of the executable file to launch. The full path need not be provided if it&#39;s already in the path. Environment variables can also be used. | 
**WorkingDirectory** | Pointer to **string** | Specifies which working directory the executable is launched from. Environment variables can be used.  | [optional] 

## Methods

### NewCreateApplicationRequestModelInstalledAppProperties

`func NewCreateApplicationRequestModelInstalledAppProperties(commandLineExecutable string, ) *CreateApplicationRequestModelInstalledAppProperties`

NewCreateApplicationRequestModelInstalledAppProperties instantiates a new CreateApplicationRequestModelInstalledAppProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationRequestModelInstalledAppPropertiesWithDefaults

`func NewCreateApplicationRequestModelInstalledAppPropertiesWithDefaults() *CreateApplicationRequestModelInstalledAppProperties`

NewCreateApplicationRequestModelInstalledAppPropertiesWithDefaults instantiates a new CreateApplicationRequestModelInstalledAppProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandLineArguments

`func (o *CreateApplicationRequestModelInstalledAppProperties) GetCommandLineArguments() string`

GetCommandLineArguments returns the CommandLineArguments field if non-nil, zero value otherwise.

### GetCommandLineArgumentsOk

`func (o *CreateApplicationRequestModelInstalledAppProperties) GetCommandLineArgumentsOk() (*string, bool)`

GetCommandLineArgumentsOk returns a tuple with the CommandLineArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArguments

`func (o *CreateApplicationRequestModelInstalledAppProperties) SetCommandLineArguments(v string)`

SetCommandLineArguments sets CommandLineArguments field to given value.

### HasCommandLineArguments

`func (o *CreateApplicationRequestModelInstalledAppProperties) HasCommandLineArguments() bool`

HasCommandLineArguments returns a boolean if a field has been set.

### GetCommandLineExecutable

`func (o *CreateApplicationRequestModelInstalledAppProperties) GetCommandLineExecutable() string`

GetCommandLineExecutable returns the CommandLineExecutable field if non-nil, zero value otherwise.

### GetCommandLineExecutableOk

`func (o *CreateApplicationRequestModelInstalledAppProperties) GetCommandLineExecutableOk() (*string, bool)`

GetCommandLineExecutableOk returns a tuple with the CommandLineExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineExecutable

`func (o *CreateApplicationRequestModelInstalledAppProperties) SetCommandLineExecutable(v string)`

SetCommandLineExecutable sets CommandLineExecutable field to given value.


### GetWorkingDirectory

`func (o *CreateApplicationRequestModelInstalledAppProperties) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *CreateApplicationRequestModelInstalledAppProperties) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *CreateApplicationRequestModelInstalledAppProperties) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *CreateApplicationRequestModelInstalledAppProperties) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


