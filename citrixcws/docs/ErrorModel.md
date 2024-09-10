# ErrorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] 
**Detail** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to **[]map[string]string** |  | [optional] 

## Methods

### NewErrorModel

`func NewErrorModel() *ErrorModel`

NewErrorModel instantiates a new ErrorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorModelWithDefaults

`func NewErrorModelWithDefaults() *ErrorModel`

NewErrorModelWithDefaults instantiates a new ErrorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ErrorModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ErrorModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDetail

`func (o *ErrorModel) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorModel) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorModel) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorModel) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ErrorModel) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ErrorModel) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetParameters

`func (o *ErrorModel) GetParameters() []map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ErrorModel) GetParametersOk() (*[]map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ErrorModel) SetParameters(v []map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ErrorModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *ErrorModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *ErrorModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


