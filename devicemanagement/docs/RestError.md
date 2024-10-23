# RestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Detail** | **string** |  | 
**Parameters** | Pointer to [**[]RestErrorParametersInner**](RestErrorParametersInner.md) |  | [optional] 

## Methods

### NewRestError

`func NewRestError(type_ string, detail string, ) *RestError`

NewRestError instantiates a new RestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestErrorWithDefaults

`func NewRestErrorWithDefaults() *RestError`

NewRestErrorWithDefaults instantiates a new RestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RestError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestError) SetType(v string)`

SetType sets Type field to given value.


### GetDetail

`func (o *RestError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RestError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RestError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetParameters

`func (o *RestError) GetParameters() []RestErrorParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RestError) GetParametersOk() (*[]RestErrorParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RestError) SetParameters(v []RestErrorParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *RestError) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


