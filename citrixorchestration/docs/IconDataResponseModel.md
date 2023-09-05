# IconDataResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RawData** | **string** | Raw icon data, Base64-encoded. | 
**FormattedData** | Pointer to **string** | Icon data in the requested format for display.  For example, if the caller requested \&quot;image/png;32x32x24\&quot; image format, then this will be a 32x32, 24bpp PNG-format image, Base64-encoded.  If the caller did not request an icon format, this will not be set. | [optional] 

## Methods

### NewIconDataResponseModel

`func NewIconDataResponseModel(rawData string, ) *IconDataResponseModel`

NewIconDataResponseModel instantiates a new IconDataResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIconDataResponseModelWithDefaults

`func NewIconDataResponseModelWithDefaults() *IconDataResponseModel`

NewIconDataResponseModelWithDefaults instantiates a new IconDataResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRawData

`func (o *IconDataResponseModel) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *IconDataResponseModel) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *IconDataResponseModel) SetRawData(v string)`

SetRawData sets RawData field to given value.


### GetFormattedData

`func (o *IconDataResponseModel) GetFormattedData() string`

GetFormattedData returns the FormattedData field if non-nil, zero value otherwise.

### GetFormattedDataOk

`func (o *IconDataResponseModel) GetFormattedDataOk() (*string, bool)`

GetFormattedDataOk returns a tuple with the FormattedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedData

`func (o *IconDataResponseModel) SetFormattedData(v string)`

SetFormattedData sets FormattedData field to given value.

### HasFormattedData

`func (o *IconDataResponseModel) HasFormattedData() bool`

HasFormattedData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


