# SystemNetHttpHttpRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**SystemVersion**](SystemVersion.md) |  | [optional] 
**VersionPolicy** | Pointer to [**SystemNetHttpHttpVersionPolicy**](SystemNetHttpHttpVersionPolicy.md) |  | [optional] 
**Content** | Pointer to [**SystemNetHttpHttpContent**](SystemNetHttpHttpContent.md) |  | [optional] 
**Method** | Pointer to [**SystemNetHttpHttpMethod**](SystemNetHttpHttpMethod.md) |  | [optional] 
**RequestUri** | Pointer to **NullableString** |  | [optional] 
**Headers** | Pointer to [**[]SystemCollectionsGenericKeyValueEnumerablePair**](SystemCollectionsGenericKeyValueEnumerablePair.md) |  | [optional] [readonly] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewSystemNetHttpHttpRequestMessage

`func NewSystemNetHttpHttpRequestMessage() *SystemNetHttpHttpRequestMessage`

NewSystemNetHttpHttpRequestMessage instantiates a new SystemNetHttpHttpRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemNetHttpHttpRequestMessageWithDefaults

`func NewSystemNetHttpHttpRequestMessageWithDefaults() *SystemNetHttpHttpRequestMessage`

NewSystemNetHttpHttpRequestMessageWithDefaults instantiates a new SystemNetHttpHttpRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SystemNetHttpHttpRequestMessage) GetVersion() SystemVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemNetHttpHttpRequestMessage) GetVersionOk() (*SystemVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemNetHttpHttpRequestMessage) SetVersion(v SystemVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemNetHttpHttpRequestMessage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionPolicy

`func (o *SystemNetHttpHttpRequestMessage) GetVersionPolicy() SystemNetHttpHttpVersionPolicy`

GetVersionPolicy returns the VersionPolicy field if non-nil, zero value otherwise.

### GetVersionPolicyOk

`func (o *SystemNetHttpHttpRequestMessage) GetVersionPolicyOk() (*SystemNetHttpHttpVersionPolicy, bool)`

GetVersionPolicyOk returns a tuple with the VersionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPolicy

`func (o *SystemNetHttpHttpRequestMessage) SetVersionPolicy(v SystemNetHttpHttpVersionPolicy)`

SetVersionPolicy sets VersionPolicy field to given value.

### HasVersionPolicy

`func (o *SystemNetHttpHttpRequestMessage) HasVersionPolicy() bool`

HasVersionPolicy returns a boolean if a field has been set.

### GetContent

`func (o *SystemNetHttpHttpRequestMessage) GetContent() SystemNetHttpHttpContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SystemNetHttpHttpRequestMessage) GetContentOk() (*SystemNetHttpHttpContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SystemNetHttpHttpRequestMessage) SetContent(v SystemNetHttpHttpContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *SystemNetHttpHttpRequestMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMethod

`func (o *SystemNetHttpHttpRequestMessage) GetMethod() SystemNetHttpHttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SystemNetHttpHttpRequestMessage) GetMethodOk() (*SystemNetHttpHttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SystemNetHttpHttpRequestMessage) SetMethod(v SystemNetHttpHttpMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SystemNetHttpHttpRequestMessage) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRequestUri

`func (o *SystemNetHttpHttpRequestMessage) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *SystemNetHttpHttpRequestMessage) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *SystemNetHttpHttpRequestMessage) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *SystemNetHttpHttpRequestMessage) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### SetRequestUriNil

`func (o *SystemNetHttpHttpRequestMessage) SetRequestUriNil(b bool)`

 SetRequestUriNil sets the value for RequestUri to be an explicit nil

### UnsetRequestUri
`func (o *SystemNetHttpHttpRequestMessage) UnsetRequestUri()`

UnsetRequestUri ensures that no value is present for RequestUri, not even an explicit nil
### GetHeaders

`func (o *SystemNetHttpHttpRequestMessage) GetHeaders() []SystemCollectionsGenericKeyValueEnumerablePair`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SystemNetHttpHttpRequestMessage) GetHeadersOk() (*[]SystemCollectionsGenericKeyValueEnumerablePair, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SystemNetHttpHttpRequestMessage) SetHeaders(v []SystemCollectionsGenericKeyValueEnumerablePair)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SystemNetHttpHttpRequestMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *SystemNetHttpHttpRequestMessage) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *SystemNetHttpHttpRequestMessage) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetProperties

`func (o *SystemNetHttpHttpRequestMessage) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SystemNetHttpHttpRequestMessage) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SystemNetHttpHttpRequestMessage) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SystemNetHttpHttpRequestMessage) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *SystemNetHttpHttpRequestMessage) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *SystemNetHttpHttpRequestMessage) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetOptions

`func (o *SystemNetHttpHttpRequestMessage) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SystemNetHttpHttpRequestMessage) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SystemNetHttpHttpRequestMessage) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SystemNetHttpHttpRequestMessage) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *SystemNetHttpHttpRequestMessage) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *SystemNetHttpHttpRequestMessage) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


