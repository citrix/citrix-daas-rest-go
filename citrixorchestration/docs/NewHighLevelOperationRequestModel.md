# NewHighLevelOperationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | [**LogOperationType**](LogOperationType.md) |  | 
**Source** | **string** | Source of the operation. | 
**Text** | **string** | Human-readable description of the change. | 
**IsSuccessful** | **bool** | Indicates whether the operation completed successfully. | 

## Methods

### NewNewHighLevelOperationRequestModel

`func NewNewHighLevelOperationRequestModel(operationType LogOperationType, source string, text string, isSuccessful bool, ) *NewHighLevelOperationRequestModel`

NewNewHighLevelOperationRequestModel instantiates a new NewHighLevelOperationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewHighLevelOperationRequestModelWithDefaults

`func NewNewHighLevelOperationRequestModelWithDefaults() *NewHighLevelOperationRequestModel`

NewNewHighLevelOperationRequestModelWithDefaults instantiates a new NewHighLevelOperationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *NewHighLevelOperationRequestModel) GetOperationType() LogOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *NewHighLevelOperationRequestModel) GetOperationTypeOk() (*LogOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *NewHighLevelOperationRequestModel) SetOperationType(v LogOperationType)`

SetOperationType sets OperationType field to given value.


### GetSource

`func (o *NewHighLevelOperationRequestModel) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NewHighLevelOperationRequestModel) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NewHighLevelOperationRequestModel) SetSource(v string)`

SetSource sets Source field to given value.


### GetText

`func (o *NewHighLevelOperationRequestModel) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NewHighLevelOperationRequestModel) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NewHighLevelOperationRequestModel) SetText(v string)`

SetText sets Text field to given value.


### GetIsSuccessful

`func (o *NewHighLevelOperationRequestModel) GetIsSuccessful() bool`

GetIsSuccessful returns the IsSuccessful field if non-nil, zero value otherwise.

### GetIsSuccessfulOk

`func (o *NewHighLevelOperationRequestModel) GetIsSuccessfulOk() (*bool, bool)`

GetIsSuccessfulOk returns a tuple with the IsSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccessful

`func (o *NewHighLevelOperationRequestModel) SetIsSuccessful(v bool)`

SetIsSuccessful sets IsSuccessful field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


