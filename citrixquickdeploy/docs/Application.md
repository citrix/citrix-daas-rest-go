# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of app, generally a GUID | 
**Name** | **string** | Display name of app | 
**ApplicationPath** | **string** | Path to execute the application | 
**WorkingDirectory** | Pointer to **string** | Working directory of the app at launch | [optional] 
**Description** | Pointer to **string** | Application description that show up in Workspace | [optional] 
**Compressedb64Icon** | Pointer to **string** | The compressed b64 icon used by UI | [optional] 
**CommandLineParams** | Pointer to **string** | Extra parameters to provide the application when it is launched | [optional] 

## Methods

### NewApplication

`func NewApplication(id string, name string, applicationPath string, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Application) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.


### GetApplicationPath

`func (o *Application) GetApplicationPath() string`

GetApplicationPath returns the ApplicationPath field if non-nil, zero value otherwise.

### GetApplicationPathOk

`func (o *Application) GetApplicationPathOk() (*string, bool)`

GetApplicationPathOk returns a tuple with the ApplicationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPath

`func (o *Application) SetApplicationPath(v string)`

SetApplicationPath sets ApplicationPath field to given value.


### GetWorkingDirectory

`func (o *Application) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *Application) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *Application) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *Application) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCompressedb64Icon

`func (o *Application) GetCompressedb64Icon() string`

GetCompressedb64Icon returns the Compressedb64Icon field if non-nil, zero value otherwise.

### GetCompressedb64IconOk

`func (o *Application) GetCompressedb64IconOk() (*string, bool)`

GetCompressedb64IconOk returns a tuple with the Compressedb64Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedb64Icon

`func (o *Application) SetCompressedb64Icon(v string)`

SetCompressedb64Icon sets Compressedb64Icon field to given value.

### HasCompressedb64Icon

`func (o *Application) HasCompressedb64Icon() bool`

HasCompressedb64Icon returns a boolean if a field has been set.

### GetCommandLineParams

`func (o *Application) GetCommandLineParams() string`

GetCommandLineParams returns the CommandLineParams field if non-nil, zero value otherwise.

### GetCommandLineParamsOk

`func (o *Application) GetCommandLineParamsOk() (*string, bool)`

GetCommandLineParamsOk returns a tuple with the CommandLineParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineParams

`func (o *Application) SetCommandLineParams(v string)`

SetCommandLineParams sets CommandLineParams field to given value.

### HasCommandLineParams

`func (o *Application) HasCommandLineParams() bool`

HasCommandLineParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


