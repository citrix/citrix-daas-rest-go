# ObjectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Formatters** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**DeclaredType** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewObjectResult

`func NewObjectResult() *ObjectResult`

NewObjectResult instantiates a new ObjectResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectResultWithDefaults

`func NewObjectResultWithDefaults() *ObjectResult`

NewObjectResultWithDefaults instantiates a new ObjectResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ObjectResult) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ObjectResult) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ObjectResult) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ObjectResult) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFormatters

`func (o *ObjectResult) GetFormatters() []map[string]interface{}`

GetFormatters returns the Formatters field if non-nil, zero value otherwise.

### GetFormattersOk

`func (o *ObjectResult) GetFormattersOk() (*[]map[string]interface{}, bool)`

GetFormattersOk returns a tuple with the Formatters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatters

`func (o *ObjectResult) SetFormatters(v []map[string]interface{})`

SetFormatters sets Formatters field to given value.

### HasFormatters

`func (o *ObjectResult) HasFormatters() bool`

HasFormatters returns a boolean if a field has been set.

### GetContentTypes

`func (o *ObjectResult) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *ObjectResult) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *ObjectResult) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *ObjectResult) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetDeclaredType

`func (o *ObjectResult) GetDeclaredType() string`

GetDeclaredType returns the DeclaredType field if non-nil, zero value otherwise.

### GetDeclaredTypeOk

`func (o *ObjectResult) GetDeclaredTypeOk() (*string, bool)`

GetDeclaredTypeOk returns a tuple with the DeclaredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredType

`func (o *ObjectResult) SetDeclaredType(v string)`

SetDeclaredType sets DeclaredType field to given value.

### HasDeclaredType

`func (o *ObjectResult) HasDeclaredType() bool`

HasDeclaredType returns a boolean if a field has been set.

### GetStatusCode

`func (o *ObjectResult) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ObjectResult) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ObjectResult) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ObjectResult) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


