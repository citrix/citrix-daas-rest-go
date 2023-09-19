# BatchRequestItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | **string** | Reference.  Required, and must be unique within all items in a single batch request. | 
**Method** | Pointer to **NullableString** | HTTP method of the request endpoint.  (GET, HEAD, PUT, PATCH, POST, DELETE) | [optional] [default to "GET"]
**RelativeUrl** | **string** | Relative URL of the request endpoint vs. the site root. Must start with the API version i.e. &#x60;\&quot;v1/{customerid}/{siteid}/...\&quot;&#x60; | 
**Headers** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | List of request headers. | [optional] 
**Body** | Pointer to **NullableString** | Request body.  Optional for PUT, PATCH, and POST. Cannot be specified for GET, HEAD, and DELETE. | [optional] 

## Methods

### NewBatchRequestItemModel

`func NewBatchRequestItemModel(reference string, relativeUrl string, ) *BatchRequestItemModel`

NewBatchRequestItemModel instantiates a new BatchRequestItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestItemModelWithDefaults

`func NewBatchRequestItemModelWithDefaults() *BatchRequestItemModel`

NewBatchRequestItemModelWithDefaults instantiates a new BatchRequestItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *BatchRequestItemModel) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BatchRequestItemModel) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BatchRequestItemModel) SetReference(v string)`

SetReference sets Reference field to given value.


### GetMethod

`func (o *BatchRequestItemModel) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *BatchRequestItemModel) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *BatchRequestItemModel) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *BatchRequestItemModel) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *BatchRequestItemModel) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *BatchRequestItemModel) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetRelativeUrl

`func (o *BatchRequestItemModel) GetRelativeUrl() string`

GetRelativeUrl returns the RelativeUrl field if non-nil, zero value otherwise.

### GetRelativeUrlOk

`func (o *BatchRequestItemModel) GetRelativeUrlOk() (*string, bool)`

GetRelativeUrlOk returns a tuple with the RelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeUrl

`func (o *BatchRequestItemModel) SetRelativeUrl(v string)`

SetRelativeUrl sets RelativeUrl field to given value.


### GetHeaders

`func (o *BatchRequestItemModel) GetHeaders() []NameValueStringPairModel`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BatchRequestItemModel) GetHeadersOk() (*[]NameValueStringPairModel, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BatchRequestItemModel) SetHeaders(v []NameValueStringPairModel)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BatchRequestItemModel) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *BatchRequestItemModel) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *BatchRequestItemModel) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetBody

`func (o *BatchRequestItemModel) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BatchRequestItemModel) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BatchRequestItemModel) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *BatchRequestItemModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *BatchRequestItemModel) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *BatchRequestItemModel) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


