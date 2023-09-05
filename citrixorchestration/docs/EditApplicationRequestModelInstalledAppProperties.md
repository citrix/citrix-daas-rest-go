# EditApplicationRequestModelInstalledAppProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLineArguments** | Pointer to **string** | The command-line arguments to use when launching the executable. Environment variables can be used. If not specified, the value will not be changed. | [optional] 
**CommandLineExecutable** | Pointer to **string** | The name of the executable file to launch. The full path need not be provided if it&#39;s already in the path. Environment variables can also be used. If not specified, the value will not be changed. | [optional] 
**WorkingDirectory** | Pointer to **string** | The working directory which the executable is launched from. Environment variables can be used.  If not specified, the value will not be changed. | [optional] 

## Methods

### NewEditApplicationRequestModelInstalledAppProperties

`func NewEditApplicationRequestModelInstalledAppProperties() *EditApplicationRequestModelInstalledAppProperties`

NewEditApplicationRequestModelInstalledAppProperties instantiates a new EditApplicationRequestModelInstalledAppProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditApplicationRequestModelInstalledAppPropertiesWithDefaults

`func NewEditApplicationRequestModelInstalledAppPropertiesWithDefaults() *EditApplicationRequestModelInstalledAppProperties`

NewEditApplicationRequestModelInstalledAppPropertiesWithDefaults instantiates a new EditApplicationRequestModelInstalledAppProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandLineArguments

`func (o *EditApplicationRequestModelInstalledAppProperties) GetCommandLineArguments() string`

GetCommandLineArguments returns the CommandLineArguments field if non-nil, zero value otherwise.

### GetCommandLineArgumentsOk

`func (o *EditApplicationRequestModelInstalledAppProperties) GetCommandLineArgumentsOk() (*string, bool)`

GetCommandLineArgumentsOk returns a tuple with the CommandLineArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArguments

`func (o *EditApplicationRequestModelInstalledAppProperties) SetCommandLineArguments(v string)`

SetCommandLineArguments sets CommandLineArguments field to given value.

### HasCommandLineArguments

`func (o *EditApplicationRequestModelInstalledAppProperties) HasCommandLineArguments() bool`

HasCommandLineArguments returns a boolean if a field has been set.

### GetCommandLineExecutable

`func (o *EditApplicationRequestModelInstalledAppProperties) GetCommandLineExecutable() string`

GetCommandLineExecutable returns the CommandLineExecutable field if non-nil, zero value otherwise.

### GetCommandLineExecutableOk

`func (o *EditApplicationRequestModelInstalledAppProperties) GetCommandLineExecutableOk() (*string, bool)`

GetCommandLineExecutableOk returns a tuple with the CommandLineExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineExecutable

`func (o *EditApplicationRequestModelInstalledAppProperties) SetCommandLineExecutable(v string)`

SetCommandLineExecutable sets CommandLineExecutable field to given value.

### HasCommandLineExecutable

`func (o *EditApplicationRequestModelInstalledAppProperties) HasCommandLineExecutable() bool`

HasCommandLineExecutable returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *EditApplicationRequestModelInstalledAppProperties) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *EditApplicationRequestModelInstalledAppProperties) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *EditApplicationRequestModelInstalledAppProperties) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *EditApplicationRequestModelInstalledAppProperties) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


