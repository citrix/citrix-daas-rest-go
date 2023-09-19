# AddIconRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RawData** | **string** | Raw icon data, Base64-encoded.              | 
**IconFormat** | Pointer to **NullableString** | Icon format.             Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;             * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.)             * _width_ and _height_ are specified in pixels.             * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.              example: &#x60;\&quot;image/png;32x32x24\&quot;&#x60;             Optional.If specified, the formatted icon data  will be contained in the response.              | [optional] 

## Methods

### NewAddIconRequestModel

`func NewAddIconRequestModel(rawData string, ) *AddIconRequestModel`

NewAddIconRequestModel instantiates a new AddIconRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIconRequestModelWithDefaults

`func NewAddIconRequestModelWithDefaults() *AddIconRequestModel`

NewAddIconRequestModelWithDefaults instantiates a new AddIconRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRawData

`func (o *AddIconRequestModel) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *AddIconRequestModel) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *AddIconRequestModel) SetRawData(v string)`

SetRawData sets RawData field to given value.


### GetIconFormat

`func (o *AddIconRequestModel) GetIconFormat() string`

GetIconFormat returns the IconFormat field if non-nil, zero value otherwise.

### GetIconFormatOk

`func (o *AddIconRequestModel) GetIconFormatOk() (*string, bool)`

GetIconFormatOk returns a tuple with the IconFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFormat

`func (o *AddIconRequestModel) SetIconFormat(v string)`

SetIconFormat sets IconFormat field to given value.

### HasIconFormat

`func (o *AddIconRequestModel) HasIconFormat() bool`

HasIconFormat returns a boolean if a field has been set.

### SetIconFormatNil

`func (o *AddIconRequestModel) SetIconFormatNil(b bool)`

 SetIconFormatNil sets the value for IconFormat to be an explicit nil

### UnsetIconFormat
`func (o *AddIconRequestModel) UnsetIconFormat()`

UnsetIconFormat ensures that no value is present for IconFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


