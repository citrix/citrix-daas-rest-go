# UpdateApplicationConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of app | 
**ApplicationPath** | **string** | Path to execute the application | 
**WorkingDirectory** | Pointer to **string** | Working directory of the app at launch | [optional] 
**Description** | Pointer to **string** | Application description that show up in Workspace | [optional] 
**CommandLineParams** | Pointer to **string** | Extra parameters to provide the application when it is launched | [optional] 
**Base64Icon** | Pointer to **string** | The raw app icon represented as a base64 string | [optional] 

## Methods

### NewUpdateApplicationConfigurationModel

`func NewUpdateApplicationConfigurationModel(name string, applicationPath string, ) *UpdateApplicationConfigurationModel`

NewUpdateApplicationConfigurationModel instantiates a new UpdateApplicationConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationConfigurationModelWithDefaults

`func NewUpdateApplicationConfigurationModelWithDefaults() *UpdateApplicationConfigurationModel`

NewUpdateApplicationConfigurationModelWithDefaults instantiates a new UpdateApplicationConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateApplicationConfigurationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApplicationConfigurationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApplicationConfigurationModel) SetName(v string)`

SetName sets Name field to given value.


### GetApplicationPath

`func (o *UpdateApplicationConfigurationModel) GetApplicationPath() string`

GetApplicationPath returns the ApplicationPath field if non-nil, zero value otherwise.

### GetApplicationPathOk

`func (o *UpdateApplicationConfigurationModel) GetApplicationPathOk() (*string, bool)`

GetApplicationPathOk returns a tuple with the ApplicationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPath

`func (o *UpdateApplicationConfigurationModel) SetApplicationPath(v string)`

SetApplicationPath sets ApplicationPath field to given value.


### GetWorkingDirectory

`func (o *UpdateApplicationConfigurationModel) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *UpdateApplicationConfigurationModel) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *UpdateApplicationConfigurationModel) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *UpdateApplicationConfigurationModel) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateApplicationConfigurationModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApplicationConfigurationModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateApplicationConfigurationModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateApplicationConfigurationModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCommandLineParams

`func (o *UpdateApplicationConfigurationModel) GetCommandLineParams() string`

GetCommandLineParams returns the CommandLineParams field if non-nil, zero value otherwise.

### GetCommandLineParamsOk

`func (o *UpdateApplicationConfigurationModel) GetCommandLineParamsOk() (*string, bool)`

GetCommandLineParamsOk returns a tuple with the CommandLineParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineParams

`func (o *UpdateApplicationConfigurationModel) SetCommandLineParams(v string)`

SetCommandLineParams sets CommandLineParams field to given value.

### HasCommandLineParams

`func (o *UpdateApplicationConfigurationModel) HasCommandLineParams() bool`

HasCommandLineParams returns a boolean if a field has been set.

### GetBase64Icon

`func (o *UpdateApplicationConfigurationModel) GetBase64Icon() string`

GetBase64Icon returns the Base64Icon field if non-nil, zero value otherwise.

### GetBase64IconOk

`func (o *UpdateApplicationConfigurationModel) GetBase64IconOk() (*string, bool)`

GetBase64IconOk returns a tuple with the Base64Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Icon

`func (o *UpdateApplicationConfigurationModel) SetBase64Icon(v string)`

SetBase64Icon sets Base64Icon field to given value.

### HasBase64Icon

`func (o *UpdateApplicationConfigurationModel) HasBase64Icon() bool`

HasBase64Icon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


