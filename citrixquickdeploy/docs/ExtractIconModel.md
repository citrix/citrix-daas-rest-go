# ExtractIconModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** |  | [optional] 
**IconFileBytes** | Pointer to **string** | The bytes of the file that contains the icon to be extracted | [optional] 

## Methods

### NewExtractIconModel

`func NewExtractIconModel() *ExtractIconModel`

NewExtractIconModel instantiates a new ExtractIconModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractIconModelWithDefaults

`func NewExtractIconModelWithDefaults() *ExtractIconModel`

NewExtractIconModelWithDefaults instantiates a new ExtractIconModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *ExtractIconModel) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ExtractIconModel) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ExtractIconModel) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ExtractIconModel) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetIconFileBytes

`func (o *ExtractIconModel) GetIconFileBytes() string`

GetIconFileBytes returns the IconFileBytes field if non-nil, zero value otherwise.

### GetIconFileBytesOk

`func (o *ExtractIconModel) GetIconFileBytesOk() (*string, bool)`

GetIconFileBytesOk returns a tuple with the IconFileBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFileBytes

`func (o *ExtractIconModel) SetIconFileBytes(v string)`

SetIconFileBytes sets IconFileBytes field to given value.

### HasIconFileBytes

`func (o *ExtractIconModel) HasIconFileBytes() bool`

HasIconFileBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


