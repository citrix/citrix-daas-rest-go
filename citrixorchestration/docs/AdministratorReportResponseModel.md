# AdministratorReportResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedReport** | Pointer to **NullableString** | The encoded report as string. That is: RFC-4648-base-64(utf-8(html-report-content)) | [optional] 
**Encoding** | Pointer to **NullableString** | The encoding of the EncodedReport. | [optional] 
**MediaType** | Pointer to **NullableString** | The media type of the report content; | [optional] 

## Methods

### NewAdministratorReportResponseModel

`func NewAdministratorReportResponseModel() *AdministratorReportResponseModel`

NewAdministratorReportResponseModel instantiates a new AdministratorReportResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorReportResponseModelWithDefaults

`func NewAdministratorReportResponseModelWithDefaults() *AdministratorReportResponseModel`

NewAdministratorReportResponseModelWithDefaults instantiates a new AdministratorReportResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodedReport

`func (o *AdministratorReportResponseModel) GetEncodedReport() string`

GetEncodedReport returns the EncodedReport field if non-nil, zero value otherwise.

### GetEncodedReportOk

`func (o *AdministratorReportResponseModel) GetEncodedReportOk() (*string, bool)`

GetEncodedReportOk returns a tuple with the EncodedReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedReport

`func (o *AdministratorReportResponseModel) SetEncodedReport(v string)`

SetEncodedReport sets EncodedReport field to given value.

### HasEncodedReport

`func (o *AdministratorReportResponseModel) HasEncodedReport() bool`

HasEncodedReport returns a boolean if a field has been set.

### SetEncodedReportNil

`func (o *AdministratorReportResponseModel) SetEncodedReportNil(b bool)`

 SetEncodedReportNil sets the value for EncodedReport to be an explicit nil

### UnsetEncodedReport
`func (o *AdministratorReportResponseModel) UnsetEncodedReport()`

UnsetEncodedReport ensures that no value is present for EncodedReport, not even an explicit nil
### GetEncoding

`func (o *AdministratorReportResponseModel) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *AdministratorReportResponseModel) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *AdministratorReportResponseModel) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *AdministratorReportResponseModel) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### SetEncodingNil

`func (o *AdministratorReportResponseModel) SetEncodingNil(b bool)`

 SetEncodingNil sets the value for Encoding to be an explicit nil

### UnsetEncoding
`func (o *AdministratorReportResponseModel) UnsetEncoding()`

UnsetEncoding ensures that no value is present for Encoding, not even an explicit nil
### GetMediaType

`func (o *AdministratorReportResponseModel) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *AdministratorReportResponseModel) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *AdministratorReportResponseModel) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *AdministratorReportResponseModel) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *AdministratorReportResponseModel) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *AdministratorReportResponseModel) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


