# FtaRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | Content-type.  (Mime type)  Optional.  If not specified, and an imported file type exists for the application with the same ExtensionName, this value will be retained.  Otherwise the default is an empty string. | [optional] 
**ExtensionName** | **string** | The extension name for the file type association. For example, \&quot;.txt\&quot; or \&quot;.doc\&quot;.   Required.  When enabling a file type association for an application, the extension name must not already be enabled for the same application.  To change properties of an already-enabled file type association, disable the file type association first. | 
**Description** | Pointer to **string** | The description of the handler for the file type association.   Optional.  If not specified, and an imported file type exists for the application with the same , this value will be retained.  Otherwise the default is an empty string. | [optional] 
**HandlerName** | Pointer to **string** | The name of the handler for the file type association (as seen in the Registry).   Optional.  If not specified, a handler name will be chosen automatically based on the application&#39;s . In most cases, callers should NOT override the default behavior and should not pass a value for this property. | [optional] 
**OpenArguments** | Pointer to **string** | The arguments for the open command that the handler should use. For example, \&quot;%1\&quot;.   Required.  When enabling a file type association for an application, the extension name must not already be enabled for the same application.  To change properties of an already-enabled file type association, you must disable the file type association first. | [optional] 

## Methods

### NewFtaRequestModel

`func NewFtaRequestModel(extensionName string, ) *FtaRequestModel`

NewFtaRequestModel instantiates a new FtaRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFtaRequestModelWithDefaults

`func NewFtaRequestModelWithDefaults() *FtaRequestModel`

NewFtaRequestModelWithDefaults instantiates a new FtaRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *FtaRequestModel) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FtaRequestModel) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FtaRequestModel) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FtaRequestModel) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExtensionName

`func (o *FtaRequestModel) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *FtaRequestModel) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *FtaRequestModel) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.


### GetDescription

`func (o *FtaRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FtaRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FtaRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FtaRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHandlerName

`func (o *FtaRequestModel) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *FtaRequestModel) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *FtaRequestModel) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.

### HasHandlerName

`func (o *FtaRequestModel) HasHandlerName() bool`

HasHandlerName returns a boolean if a field has been set.

### GetOpenArguments

`func (o *FtaRequestModel) GetOpenArguments() string`

GetOpenArguments returns the OpenArguments field if non-nil, zero value otherwise.

### GetOpenArgumentsOk

`func (o *FtaRequestModel) GetOpenArgumentsOk() (*string, bool)`

GetOpenArgumentsOk returns a tuple with the OpenArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenArguments

`func (o *FtaRequestModel) SetOpenArguments(v string)`

SetOpenArguments sets OpenArguments field to given value.

### HasOpenArguments

`func (o *FtaRequestModel) HasOpenArguments() bool`

HasOpenArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


