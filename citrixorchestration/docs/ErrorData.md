# ErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **NullableString** | The message that describes the current exception. | [optional] 
**SdkErrorId** | **string** | The developer-defined identifier of the error. | 
**ErrorType** | Pointer to **NullableString** | The runtime type of the current exception. | [optional] 
**ErrorDetails** | Pointer to **NullableString** | The additional details. | [optional] 
**SupportLink** | Pointer to **NullableString** | The support url. | [optional] 
**ErrorReportXml** | Pointer to **NullableString** | Represents a citrix.com error report. | [optional] 
**ErrorReport** | Pointer to [**ErrorReport**](ErrorReport.md) |  | [optional] 

## Methods

### NewErrorData

`func NewErrorData(sdkErrorId string, ) *ErrorData`

NewErrorData instantiates a new ErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDataWithDefaults

`func NewErrorDataWithDefaults() *ErrorData`

NewErrorDataWithDefaults instantiates a new ErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *ErrorData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ErrorData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ErrorData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ErrorData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ErrorData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ErrorData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetSdkErrorId

`func (o *ErrorData) GetSdkErrorId() string`

GetSdkErrorId returns the SdkErrorId field if non-nil, zero value otherwise.

### GetSdkErrorIdOk

`func (o *ErrorData) GetSdkErrorIdOk() (*string, bool)`

GetSdkErrorIdOk returns a tuple with the SdkErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkErrorId

`func (o *ErrorData) SetSdkErrorId(v string)`

SetSdkErrorId sets SdkErrorId field to given value.


### GetErrorType

`func (o *ErrorData) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ErrorData) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ErrorData) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *ErrorData) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### SetErrorTypeNil

`func (o *ErrorData) SetErrorTypeNil(b bool)`

 SetErrorTypeNil sets the value for ErrorType to be an explicit nil

### UnsetErrorType
`func (o *ErrorData) UnsetErrorType()`

UnsetErrorType ensures that no value is present for ErrorType, not even an explicit nil
### GetErrorDetails

`func (o *ErrorData) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ErrorData) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ErrorData) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ErrorData) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### SetErrorDetailsNil

`func (o *ErrorData) SetErrorDetailsNil(b bool)`

 SetErrorDetailsNil sets the value for ErrorDetails to be an explicit nil

### UnsetErrorDetails
`func (o *ErrorData) UnsetErrorDetails()`

UnsetErrorDetails ensures that no value is present for ErrorDetails, not even an explicit nil
### GetSupportLink

`func (o *ErrorData) GetSupportLink() string`

GetSupportLink returns the SupportLink field if non-nil, zero value otherwise.

### GetSupportLinkOk

`func (o *ErrorData) GetSupportLinkOk() (*string, bool)`

GetSupportLinkOk returns a tuple with the SupportLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLink

`func (o *ErrorData) SetSupportLink(v string)`

SetSupportLink sets SupportLink field to given value.

### HasSupportLink

`func (o *ErrorData) HasSupportLink() bool`

HasSupportLink returns a boolean if a field has been set.

### SetSupportLinkNil

`func (o *ErrorData) SetSupportLinkNil(b bool)`

 SetSupportLinkNil sets the value for SupportLink to be an explicit nil

### UnsetSupportLink
`func (o *ErrorData) UnsetSupportLink()`

UnsetSupportLink ensures that no value is present for SupportLink, not even an explicit nil
### GetErrorReportXml

`func (o *ErrorData) GetErrorReportXml() string`

GetErrorReportXml returns the ErrorReportXml field if non-nil, zero value otherwise.

### GetErrorReportXmlOk

`func (o *ErrorData) GetErrorReportXmlOk() (*string, bool)`

GetErrorReportXmlOk returns a tuple with the ErrorReportXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReportXml

`func (o *ErrorData) SetErrorReportXml(v string)`

SetErrorReportXml sets ErrorReportXml field to given value.

### HasErrorReportXml

`func (o *ErrorData) HasErrorReportXml() bool`

HasErrorReportXml returns a boolean if a field has been set.

### SetErrorReportXmlNil

`func (o *ErrorData) SetErrorReportXmlNil(b bool)`

 SetErrorReportXmlNil sets the value for ErrorReportXml to be an explicit nil

### UnsetErrorReportXml
`func (o *ErrorData) UnsetErrorReportXml()`

UnsetErrorReportXml ensures that no value is present for ErrorReportXml, not even an explicit nil
### GetErrorReport

`func (o *ErrorData) GetErrorReport() ErrorReport`

GetErrorReport returns the ErrorReport field if non-nil, zero value otherwise.

### GetErrorReportOk

`func (o *ErrorData) GetErrorReportOk() (*ErrorReport, bool)`

GetErrorReportOk returns a tuple with the ErrorReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReport

`func (o *ErrorData) SetErrorReport(v ErrorReport)`

SetErrorReport sets ErrorReport field to given value.

### HasErrorReport

`func (o *ErrorData) HasErrorReport() bool`

HasErrorReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


