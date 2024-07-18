# SystemNetHttpHttpResponseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**SystemVersion**](SystemVersion.md) |  | [optional] 
**Content** | Pointer to [**SystemNetHttpHttpContent**](SystemNetHttpHttpContent.md) |  | [optional] 
**StatusCode** | Pointer to [**SystemNetHttpStatusCode**](SystemNetHttpStatusCode.md) |  | [optional] 
**ReasonPhrase** | Pointer to **NullableString** |  | [optional] 
**Headers** | Pointer to [**[]SystemCollectionsGenericKeyValueEnumerablePair**](SystemCollectionsGenericKeyValueEnumerablePair.md) |  | [optional] [readonly] 
**TrailingHeaders** | Pointer to [**[]SystemCollectionsGenericKeyValueEnumerablePair**](SystemCollectionsGenericKeyValueEnumerablePair.md) |  | [optional] [readonly] 
**RequestMessage** | Pointer to [**SystemNetHttpHttpRequestMessage**](SystemNetHttpHttpRequestMessage.md) |  | [optional] 
**IsSuccessStatusCode** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewSystemNetHttpHttpResponseMessage

`func NewSystemNetHttpHttpResponseMessage() *SystemNetHttpHttpResponseMessage`

NewSystemNetHttpHttpResponseMessage instantiates a new SystemNetHttpHttpResponseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemNetHttpHttpResponseMessageWithDefaults

`func NewSystemNetHttpHttpResponseMessageWithDefaults() *SystemNetHttpHttpResponseMessage`

NewSystemNetHttpHttpResponseMessageWithDefaults instantiates a new SystemNetHttpHttpResponseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SystemNetHttpHttpResponseMessage) GetVersion() SystemVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemNetHttpHttpResponseMessage) GetVersionOk() (*SystemVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemNetHttpHttpResponseMessage) SetVersion(v SystemVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemNetHttpHttpResponseMessage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetContent

`func (o *SystemNetHttpHttpResponseMessage) GetContent() SystemNetHttpHttpContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SystemNetHttpHttpResponseMessage) GetContentOk() (*SystemNetHttpHttpContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SystemNetHttpHttpResponseMessage) SetContent(v SystemNetHttpHttpContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *SystemNetHttpHttpResponseMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetStatusCode

`func (o *SystemNetHttpHttpResponseMessage) GetStatusCode() SystemNetHttpStatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SystemNetHttpHttpResponseMessage) GetStatusCodeOk() (*SystemNetHttpStatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SystemNetHttpHttpResponseMessage) SetStatusCode(v SystemNetHttpStatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *SystemNetHttpHttpResponseMessage) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetReasonPhrase

`func (o *SystemNetHttpHttpResponseMessage) GetReasonPhrase() string`

GetReasonPhrase returns the ReasonPhrase field if non-nil, zero value otherwise.

### GetReasonPhraseOk

`func (o *SystemNetHttpHttpResponseMessage) GetReasonPhraseOk() (*string, bool)`

GetReasonPhraseOk returns a tuple with the ReasonPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonPhrase

`func (o *SystemNetHttpHttpResponseMessage) SetReasonPhrase(v string)`

SetReasonPhrase sets ReasonPhrase field to given value.

### HasReasonPhrase

`func (o *SystemNetHttpHttpResponseMessage) HasReasonPhrase() bool`

HasReasonPhrase returns a boolean if a field has been set.

### SetReasonPhraseNil

`func (o *SystemNetHttpHttpResponseMessage) SetReasonPhraseNil(b bool)`

 SetReasonPhraseNil sets the value for ReasonPhrase to be an explicit nil

### UnsetReasonPhrase
`func (o *SystemNetHttpHttpResponseMessage) UnsetReasonPhrase()`

UnsetReasonPhrase ensures that no value is present for ReasonPhrase, not even an explicit nil
### GetHeaders

`func (o *SystemNetHttpHttpResponseMessage) GetHeaders() []SystemCollectionsGenericKeyValueEnumerablePair`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SystemNetHttpHttpResponseMessage) GetHeadersOk() (*[]SystemCollectionsGenericKeyValueEnumerablePair, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SystemNetHttpHttpResponseMessage) SetHeaders(v []SystemCollectionsGenericKeyValueEnumerablePair)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SystemNetHttpHttpResponseMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *SystemNetHttpHttpResponseMessage) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *SystemNetHttpHttpResponseMessage) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetTrailingHeaders

`func (o *SystemNetHttpHttpResponseMessage) GetTrailingHeaders() []SystemCollectionsGenericKeyValueEnumerablePair`

GetTrailingHeaders returns the TrailingHeaders field if non-nil, zero value otherwise.

### GetTrailingHeadersOk

`func (o *SystemNetHttpHttpResponseMessage) GetTrailingHeadersOk() (*[]SystemCollectionsGenericKeyValueEnumerablePair, bool)`

GetTrailingHeadersOk returns a tuple with the TrailingHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingHeaders

`func (o *SystemNetHttpHttpResponseMessage) SetTrailingHeaders(v []SystemCollectionsGenericKeyValueEnumerablePair)`

SetTrailingHeaders sets TrailingHeaders field to given value.

### HasTrailingHeaders

`func (o *SystemNetHttpHttpResponseMessage) HasTrailingHeaders() bool`

HasTrailingHeaders returns a boolean if a field has been set.

### SetTrailingHeadersNil

`func (o *SystemNetHttpHttpResponseMessage) SetTrailingHeadersNil(b bool)`

 SetTrailingHeadersNil sets the value for TrailingHeaders to be an explicit nil

### UnsetTrailingHeaders
`func (o *SystemNetHttpHttpResponseMessage) UnsetTrailingHeaders()`

UnsetTrailingHeaders ensures that no value is present for TrailingHeaders, not even an explicit nil
### GetRequestMessage

`func (o *SystemNetHttpHttpResponseMessage) GetRequestMessage() SystemNetHttpHttpRequestMessage`

GetRequestMessage returns the RequestMessage field if non-nil, zero value otherwise.

### GetRequestMessageOk

`func (o *SystemNetHttpHttpResponseMessage) GetRequestMessageOk() (*SystemNetHttpHttpRequestMessage, bool)`

GetRequestMessageOk returns a tuple with the RequestMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessage

`func (o *SystemNetHttpHttpResponseMessage) SetRequestMessage(v SystemNetHttpHttpRequestMessage)`

SetRequestMessage sets RequestMessage field to given value.

### HasRequestMessage

`func (o *SystemNetHttpHttpResponseMessage) HasRequestMessage() bool`

HasRequestMessage returns a boolean if a field has been set.

### GetIsSuccessStatusCode

`func (o *SystemNetHttpHttpResponseMessage) GetIsSuccessStatusCode() bool`

GetIsSuccessStatusCode returns the IsSuccessStatusCode field if non-nil, zero value otherwise.

### GetIsSuccessStatusCodeOk

`func (o *SystemNetHttpHttpResponseMessage) GetIsSuccessStatusCodeOk() (*bool, bool)`

GetIsSuccessStatusCodeOk returns a tuple with the IsSuccessStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccessStatusCode

`func (o *SystemNetHttpHttpResponseMessage) SetIsSuccessStatusCode(v bool)`

SetIsSuccessStatusCode sets IsSuccessStatusCode field to given value.

### HasIsSuccessStatusCode

`func (o *SystemNetHttpHttpResponseMessage) HasIsSuccessStatusCode() bool`

HasIsSuccessStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


