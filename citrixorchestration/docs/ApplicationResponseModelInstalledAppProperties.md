# ApplicationResponseModelInstalledAppProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLineArguments** | **string** | The command-line arguments to use when launching the executable. Environment variables can be used. | 
**CommandLineExecutable** | **string** | The name of the executable file to launch. The full path need not be provided if it&#39;s already in the path. Environment variables can also be used. | 
**WorkingDirectory** | **string** | The working directory which the executable is launched from. Environment variables can be used.  | 

## Methods

### NewApplicationResponseModelInstalledAppProperties

`func NewApplicationResponseModelInstalledAppProperties(commandLineArguments string, commandLineExecutable string, workingDirectory string, ) *ApplicationResponseModelInstalledAppProperties`

NewApplicationResponseModelInstalledAppProperties instantiates a new ApplicationResponseModelInstalledAppProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseModelInstalledAppPropertiesWithDefaults

`func NewApplicationResponseModelInstalledAppPropertiesWithDefaults() *ApplicationResponseModelInstalledAppProperties`

NewApplicationResponseModelInstalledAppPropertiesWithDefaults instantiates a new ApplicationResponseModelInstalledAppProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandLineArguments

`func (o *ApplicationResponseModelInstalledAppProperties) GetCommandLineArguments() string`

GetCommandLineArguments returns the CommandLineArguments field if non-nil, zero value otherwise.

### GetCommandLineArgumentsOk

`func (o *ApplicationResponseModelInstalledAppProperties) GetCommandLineArgumentsOk() (*string, bool)`

GetCommandLineArgumentsOk returns a tuple with the CommandLineArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArguments

`func (o *ApplicationResponseModelInstalledAppProperties) SetCommandLineArguments(v string)`

SetCommandLineArguments sets CommandLineArguments field to given value.


### GetCommandLineExecutable

`func (o *ApplicationResponseModelInstalledAppProperties) GetCommandLineExecutable() string`

GetCommandLineExecutable returns the CommandLineExecutable field if non-nil, zero value otherwise.

### GetCommandLineExecutableOk

`func (o *ApplicationResponseModelInstalledAppProperties) GetCommandLineExecutableOk() (*string, bool)`

GetCommandLineExecutableOk returns a tuple with the CommandLineExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineExecutable

`func (o *ApplicationResponseModelInstalledAppProperties) SetCommandLineExecutable(v string)`

SetCommandLineExecutable sets CommandLineExecutable field to given value.


### GetWorkingDirectory

`func (o *ApplicationResponseModelInstalledAppProperties) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *ApplicationResponseModelInstalledAppProperties) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *ApplicationResponseModelInstalledAppProperties) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


