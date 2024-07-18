# EditOperationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | The new labels of the operation. NOTE: For each label, it is case-insensitive with a max length of 60 and any leading or trailing whitespaces will be trimmed. | [optional] 

## Methods

### NewEditOperationRequestModel

`func NewEditOperationRequestModel() *EditOperationRequestModel`

NewEditOperationRequestModel instantiates a new EditOperationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditOperationRequestModelWithDefaults

`func NewEditOperationRequestModelWithDefaults() *EditOperationRequestModel`

NewEditOperationRequestModelWithDefaults instantiates a new EditOperationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *EditOperationRequestModel) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EditOperationRequestModel) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EditOperationRequestModel) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EditOperationRequestModel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *EditOperationRequestModel) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *EditOperationRequestModel) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


