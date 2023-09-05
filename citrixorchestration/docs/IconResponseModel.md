# IconResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RawData** | **string** | Raw icon data, Base64-encoded. | 
**FormattedData** | Pointer to **string** | Icon data in the requested format for display.  For example, if the caller requested \&quot;image/png;32x32x24\&quot; image format, then this will be a 32x32, 24bpp PNG-format image, Base64-encoded.  If the caller did not request an icon format, this will not be set. | [optional] 
**Id** | **string** | Id of the icon. | 
**IsBuiltIn** | **bool** | Indicates if the icon is built-in (i.e. Citrix-provided).  The built-in icons can be used for applications or desktops that don&#39;t have their own icons. | 
**Uid** | Pointer to **int32** | DEPRECATED.  Use Id. | [optional] 

## Methods

### NewIconResponseModel

`func NewIconResponseModel(rawData string, id string, isBuiltIn bool, ) *IconResponseModel`

NewIconResponseModel instantiates a new IconResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIconResponseModelWithDefaults

`func NewIconResponseModelWithDefaults() *IconResponseModel`

NewIconResponseModelWithDefaults instantiates a new IconResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRawData

`func (o *IconResponseModel) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *IconResponseModel) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *IconResponseModel) SetRawData(v string)`

SetRawData sets RawData field to given value.


### GetFormattedData

`func (o *IconResponseModel) GetFormattedData() string`

GetFormattedData returns the FormattedData field if non-nil, zero value otherwise.

### GetFormattedDataOk

`func (o *IconResponseModel) GetFormattedDataOk() (*string, bool)`

GetFormattedDataOk returns a tuple with the FormattedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedData

`func (o *IconResponseModel) SetFormattedData(v string)`

SetFormattedData sets FormattedData field to given value.

### HasFormattedData

`func (o *IconResponseModel) HasFormattedData() bool`

HasFormattedData returns a boolean if a field has been set.

### GetId

`func (o *IconResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IconResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IconResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetIsBuiltIn

`func (o *IconResponseModel) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *IconResponseModel) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *IconResponseModel) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.


### GetUid

`func (o *IconResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IconResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IconResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IconResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


