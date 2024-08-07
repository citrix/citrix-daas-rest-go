# ServiceAccountTestResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccount** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**NumPassed** | Pointer to **int32** | The number of tests that passed. | [optional] 
**NumWarnings** | Pointer to **int32** | The number of warnings that were found. | [optional] 
**NumFailures** | Pointer to **int32** | The number of tests that failed. | [optional] 

## Methods

### NewServiceAccountTestResponseModel

`func NewServiceAccountTestResponseModel() *ServiceAccountTestResponseModel`

NewServiceAccountTestResponseModel instantiates a new ServiceAccountTestResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountTestResponseModelWithDefaults

`func NewServiceAccountTestResponseModelWithDefaults() *ServiceAccountTestResponseModel`

NewServiceAccountTestResponseModelWithDefaults instantiates a new ServiceAccountTestResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccount

`func (o *ServiceAccountTestResponseModel) GetServiceAccount() RefResponseModel`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ServiceAccountTestResponseModel) GetServiceAccountOk() (*RefResponseModel, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ServiceAccountTestResponseModel) SetServiceAccount(v RefResponseModel)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *ServiceAccountTestResponseModel) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetNumPassed

`func (o *ServiceAccountTestResponseModel) GetNumPassed() int32`

GetNumPassed returns the NumPassed field if non-nil, zero value otherwise.

### GetNumPassedOk

`func (o *ServiceAccountTestResponseModel) GetNumPassedOk() (*int32, bool)`

GetNumPassedOk returns a tuple with the NumPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPassed

`func (o *ServiceAccountTestResponseModel) SetNumPassed(v int32)`

SetNumPassed sets NumPassed field to given value.

### HasNumPassed

`func (o *ServiceAccountTestResponseModel) HasNumPassed() bool`

HasNumPassed returns a boolean if a field has been set.

### GetNumWarnings

`func (o *ServiceAccountTestResponseModel) GetNumWarnings() int32`

GetNumWarnings returns the NumWarnings field if non-nil, zero value otherwise.

### GetNumWarningsOk

`func (o *ServiceAccountTestResponseModel) GetNumWarningsOk() (*int32, bool)`

GetNumWarningsOk returns a tuple with the NumWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWarnings

`func (o *ServiceAccountTestResponseModel) SetNumWarnings(v int32)`

SetNumWarnings sets NumWarnings field to given value.

### HasNumWarnings

`func (o *ServiceAccountTestResponseModel) HasNumWarnings() bool`

HasNumWarnings returns a boolean if a field has been set.

### GetNumFailures

`func (o *ServiceAccountTestResponseModel) GetNumFailures() int32`

GetNumFailures returns the NumFailures field if non-nil, zero value otherwise.

### GetNumFailuresOk

`func (o *ServiceAccountTestResponseModel) GetNumFailuresOk() (*int32, bool)`

GetNumFailuresOk returns a tuple with the NumFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailures

`func (o *ServiceAccountTestResponseModel) SetNumFailures(v int32)`

SetNumFailures sets NumFailures field to given value.

### HasNumFailures

`func (o *ServiceAccountTestResponseModel) HasNumFailures() bool`

HasNumFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


