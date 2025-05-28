# PublishedApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**AppsAndDesktopsProvisionState**](AppsAndDesktopsProvisionState.md) | Current state in publishing the application | 
**Status** | Pointer to **string** | Status message related to application state | [optional] 
**Id** | **string** | Unique identifier of app, generally a GUID | 
**Name** | **string** | Display name of app | 
**ApplicationPath** | **string** | Path to execute the application | 
**WorkingDirectory** | Pointer to **string** | Working directory of the app at launch | [optional] 
**Description** | Pointer to **string** | Application description that show up in Workspace | [optional] 
**Compressedb64Icon** | Pointer to **string** | The compressed b64 icon used by UI | [optional] 
**CommandLineParams** | Pointer to **string** | Extra parameters to provide the application when it is launched | [optional] 

## Methods

### NewPublishedApplication

`func NewPublishedApplication(state AppsAndDesktopsProvisionState, id string, name string, applicationPath string, ) *PublishedApplication`

NewPublishedApplication instantiates a new PublishedApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedApplicationWithDefaults

`func NewPublishedApplicationWithDefaults() *PublishedApplication`

NewPublishedApplicationWithDefaults instantiates a new PublishedApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PublishedApplication) GetState() AppsAndDesktopsProvisionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PublishedApplication) GetStateOk() (*AppsAndDesktopsProvisionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PublishedApplication) SetState(v AppsAndDesktopsProvisionState)`

SetState sets State field to given value.


### GetStatus

`func (o *PublishedApplication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublishedApplication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublishedApplication) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublishedApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *PublishedApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublishedApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublishedApplication) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PublishedApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishedApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishedApplication) SetName(v string)`

SetName sets Name field to given value.


### GetApplicationPath

`func (o *PublishedApplication) GetApplicationPath() string`

GetApplicationPath returns the ApplicationPath field if non-nil, zero value otherwise.

### GetApplicationPathOk

`func (o *PublishedApplication) GetApplicationPathOk() (*string, bool)`

GetApplicationPathOk returns a tuple with the ApplicationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPath

`func (o *PublishedApplication) SetApplicationPath(v string)`

SetApplicationPath sets ApplicationPath field to given value.


### GetWorkingDirectory

`func (o *PublishedApplication) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *PublishedApplication) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *PublishedApplication) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *PublishedApplication) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetDescription

`func (o *PublishedApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublishedApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublishedApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublishedApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCompressedb64Icon

`func (o *PublishedApplication) GetCompressedb64Icon() string`

GetCompressedb64Icon returns the Compressedb64Icon field if non-nil, zero value otherwise.

### GetCompressedb64IconOk

`func (o *PublishedApplication) GetCompressedb64IconOk() (*string, bool)`

GetCompressedb64IconOk returns a tuple with the Compressedb64Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedb64Icon

`func (o *PublishedApplication) SetCompressedb64Icon(v string)`

SetCompressedb64Icon sets Compressedb64Icon field to given value.

### HasCompressedb64Icon

`func (o *PublishedApplication) HasCompressedb64Icon() bool`

HasCompressedb64Icon returns a boolean if a field has been set.

### GetCommandLineParams

`func (o *PublishedApplication) GetCommandLineParams() string`

GetCommandLineParams returns the CommandLineParams field if non-nil, zero value otherwise.

### GetCommandLineParamsOk

`func (o *PublishedApplication) GetCommandLineParamsOk() (*string, bool)`

GetCommandLineParamsOk returns a tuple with the CommandLineParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineParams

`func (o *PublishedApplication) SetCommandLineParams(v string)`

SetCommandLineParams sets CommandLineParams field to given value.

### HasCommandLineParams

`func (o *PublishedApplication) HasCommandLineParams() bool`

HasCommandLineParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


