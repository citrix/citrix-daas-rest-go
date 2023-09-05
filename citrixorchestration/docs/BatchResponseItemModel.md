# BatchResponseItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | **string** | Reference that allows connecting the response item to the corresponding request item. | 
**Code** | **int32** | HTTP status code as a result of the request. | 
**Headers** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Response headers. | [optional] 
**Body** | Pointer to **string** | Response body, if one was present. | [optional] 

## Methods

### NewBatchResponseItemModel

`func NewBatchResponseItemModel(reference string, code int32, ) *BatchResponseItemModel`

NewBatchResponseItemModel instantiates a new BatchResponseItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseItemModelWithDefaults

`func NewBatchResponseItemModelWithDefaults() *BatchResponseItemModel`

NewBatchResponseItemModelWithDefaults instantiates a new BatchResponseItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *BatchResponseItemModel) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BatchResponseItemModel) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BatchResponseItemModel) SetReference(v string)`

SetReference sets Reference field to given value.


### GetCode

`func (o *BatchResponseItemModel) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BatchResponseItemModel) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BatchResponseItemModel) SetCode(v int32)`

SetCode sets Code field to given value.


### GetHeaders

`func (o *BatchResponseItemModel) GetHeaders() []NameValueStringPairModel`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BatchResponseItemModel) GetHeadersOk() (*[]NameValueStringPairModel, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BatchResponseItemModel) SetHeaders(v []NameValueStringPairModel)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BatchResponseItemModel) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *BatchResponseItemModel) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BatchResponseItemModel) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BatchResponseItemModel) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *BatchResponseItemModel) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


