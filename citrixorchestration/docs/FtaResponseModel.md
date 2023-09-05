# FtaResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | Content-type.  (Mime type) | [optional] 
**ExtensionName** | **string** | The extension name for the file type association. For example, \&quot;.txt\&quot; or \&quot;.doc\&quot;. | 
**Description** | Pointer to **string** | The description of the handler for the file type association. | [optional] 
**HandlerName** | **string** | The name of the handler for the file type association (as seen in the Registry). For example, \&quot;TXTFILE\&quot; or \&quot;Word.Document.8\&quot;. | 
**OpenArguments** | Pointer to **string** | The arguments for the open command that the handler should use. For example, \&quot;%1\&quot;. | [optional] 
**Enabled** | **bool** | Indicates if the file type association is enabled; i.e. whether users who receive the application will launch files of this type using the published application. | 
**IsImported** | **bool** | Indicates if the file type association was imported, as opposed to explicitly added.  File type associations which are not imported will be deleted when disabled. | 

## Methods

### NewFtaResponseModel

`func NewFtaResponseModel(extensionName string, handlerName string, enabled bool, isImported bool, ) *FtaResponseModel`

NewFtaResponseModel instantiates a new FtaResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFtaResponseModelWithDefaults

`func NewFtaResponseModelWithDefaults() *FtaResponseModel`

NewFtaResponseModelWithDefaults instantiates a new FtaResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *FtaResponseModel) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FtaResponseModel) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FtaResponseModel) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FtaResponseModel) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExtensionName

`func (o *FtaResponseModel) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *FtaResponseModel) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *FtaResponseModel) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.


### GetDescription

`func (o *FtaResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FtaResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FtaResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FtaResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHandlerName

`func (o *FtaResponseModel) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *FtaResponseModel) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *FtaResponseModel) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetOpenArguments

`func (o *FtaResponseModel) GetOpenArguments() string`

GetOpenArguments returns the OpenArguments field if non-nil, zero value otherwise.

### GetOpenArgumentsOk

`func (o *FtaResponseModel) GetOpenArgumentsOk() (*string, bool)`

GetOpenArgumentsOk returns a tuple with the OpenArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenArguments

`func (o *FtaResponseModel) SetOpenArguments(v string)`

SetOpenArguments sets OpenArguments field to given value.

### HasOpenArguments

`func (o *FtaResponseModel) HasOpenArguments() bool`

HasOpenArguments returns a boolean if a field has been set.

### GetEnabled

`func (o *FtaResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FtaResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FtaResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIsImported

`func (o *FtaResponseModel) GetIsImported() bool`

GetIsImported returns the IsImported field if non-nil, zero value otherwise.

### GetIsImportedOk

`func (o *FtaResponseModel) GetIsImportedOk() (*bool, bool)`

GetIsImportedOk returns a tuple with the IsImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImported

`func (o *FtaResponseModel) SetIsImported(v bool)`

SetIsImported sets IsImported field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


