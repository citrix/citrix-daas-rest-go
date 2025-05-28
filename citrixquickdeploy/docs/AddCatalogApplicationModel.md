# AddCatalogApplicationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The application id in the datastore. If the id is not known (for a start menu application) or for adding a custom application, the Id should be left null | [optional] 
**IsStartMenuDiscovered** | Pointer to **bool** | Indicates if the application was discovered via scan of start menu apps or manually entered | [optional] 
**Name** | **string** | Name of the image | 
**Path** | **string** | Path of the application | 
**WorkingDirectory** | Pointer to **string** | Working directory of the app at launch | [optional] 
**Description** | Pointer to **string** | Application description that show up in Workspace | [optional] 
**CommandLineParameters** | Pointer to **string** | Command line paramters to pass to the app when launching | [optional] 
**Base64Icon** | Pointer to **string** | The raw app icon represented as a base64 string. For start menu app, if you want to use the default app icon, this field should be left blank. | [optional] 

## Methods

### NewAddCatalogApplicationModel

`func NewAddCatalogApplicationModel(name string, path string, ) *AddCatalogApplicationModel`

NewAddCatalogApplicationModel instantiates a new AddCatalogApplicationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogApplicationModelWithDefaults

`func NewAddCatalogApplicationModelWithDefaults() *AddCatalogApplicationModel`

NewAddCatalogApplicationModelWithDefaults instantiates a new AddCatalogApplicationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddCatalogApplicationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCatalogApplicationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCatalogApplicationModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddCatalogApplicationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsStartMenuDiscovered

`func (o *AddCatalogApplicationModel) GetIsStartMenuDiscovered() bool`

GetIsStartMenuDiscovered returns the IsStartMenuDiscovered field if non-nil, zero value otherwise.

### GetIsStartMenuDiscoveredOk

`func (o *AddCatalogApplicationModel) GetIsStartMenuDiscoveredOk() (*bool, bool)`

GetIsStartMenuDiscoveredOk returns a tuple with the IsStartMenuDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStartMenuDiscovered

`func (o *AddCatalogApplicationModel) SetIsStartMenuDiscovered(v bool)`

SetIsStartMenuDiscovered sets IsStartMenuDiscovered field to given value.

### HasIsStartMenuDiscovered

`func (o *AddCatalogApplicationModel) HasIsStartMenuDiscovered() bool`

HasIsStartMenuDiscovered returns a boolean if a field has been set.

### GetName

`func (o *AddCatalogApplicationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCatalogApplicationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCatalogApplicationModel) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *AddCatalogApplicationModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AddCatalogApplicationModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AddCatalogApplicationModel) SetPath(v string)`

SetPath sets Path field to given value.


### GetWorkingDirectory

`func (o *AddCatalogApplicationModel) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *AddCatalogApplicationModel) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *AddCatalogApplicationModel) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *AddCatalogApplicationModel) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetDescription

`func (o *AddCatalogApplicationModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCatalogApplicationModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCatalogApplicationModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCatalogApplicationModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCommandLineParameters

`func (o *AddCatalogApplicationModel) GetCommandLineParameters() string`

GetCommandLineParameters returns the CommandLineParameters field if non-nil, zero value otherwise.

### GetCommandLineParametersOk

`func (o *AddCatalogApplicationModel) GetCommandLineParametersOk() (*string, bool)`

GetCommandLineParametersOk returns a tuple with the CommandLineParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineParameters

`func (o *AddCatalogApplicationModel) SetCommandLineParameters(v string)`

SetCommandLineParameters sets CommandLineParameters field to given value.

### HasCommandLineParameters

`func (o *AddCatalogApplicationModel) HasCommandLineParameters() bool`

HasCommandLineParameters returns a boolean if a field has been set.

### GetBase64Icon

`func (o *AddCatalogApplicationModel) GetBase64Icon() string`

GetBase64Icon returns the Base64Icon field if non-nil, zero value otherwise.

### GetBase64IconOk

`func (o *AddCatalogApplicationModel) GetBase64IconOk() (*string, bool)`

GetBase64IconOk returns a tuple with the Base64Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Icon

`func (o *AddCatalogApplicationModel) SetBase64Icon(v string)`

SetBase64Icon sets Base64Icon field to given value.

### HasBase64Icon

`func (o *AddCatalogApplicationModel) HasBase64Icon() bool`

HasBase64Icon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


