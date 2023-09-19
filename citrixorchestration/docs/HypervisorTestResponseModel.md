# HypervisorTestResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hypervisor** | [**RefResponseModel**](RefResponseModel.md) |  | 
**NumPassed** | **int32** | The number of tests that passed. | 
**NumWarnings** | **int32** | The number of warnings that were found. | 
**NumFailures** | **int32** | The number of tests that failed. | 

## Methods

### NewHypervisorTestResponseModel

`func NewHypervisorTestResponseModel(hypervisor RefResponseModel, numPassed int32, numWarnings int32, numFailures int32, ) *HypervisorTestResponseModel`

NewHypervisorTestResponseModel instantiates a new HypervisorTestResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorTestResponseModelWithDefaults

`func NewHypervisorTestResponseModelWithDefaults() *HypervisorTestResponseModel`

NewHypervisorTestResponseModelWithDefaults instantiates a new HypervisorTestResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypervisor

`func (o *HypervisorTestResponseModel) GetHypervisor() RefResponseModel`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *HypervisorTestResponseModel) GetHypervisorOk() (*RefResponseModel, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *HypervisorTestResponseModel) SetHypervisor(v RefResponseModel)`

SetHypervisor sets Hypervisor field to given value.


### GetNumPassed

`func (o *HypervisorTestResponseModel) GetNumPassed() int32`

GetNumPassed returns the NumPassed field if non-nil, zero value otherwise.

### GetNumPassedOk

`func (o *HypervisorTestResponseModel) GetNumPassedOk() (*int32, bool)`

GetNumPassedOk returns a tuple with the NumPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPassed

`func (o *HypervisorTestResponseModel) SetNumPassed(v int32)`

SetNumPassed sets NumPassed field to given value.


### GetNumWarnings

`func (o *HypervisorTestResponseModel) GetNumWarnings() int32`

GetNumWarnings returns the NumWarnings field if non-nil, zero value otherwise.

### GetNumWarningsOk

`func (o *HypervisorTestResponseModel) GetNumWarningsOk() (*int32, bool)`

GetNumWarningsOk returns a tuple with the NumWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWarnings

`func (o *HypervisorTestResponseModel) SetNumWarnings(v int32)`

SetNumWarnings sets NumWarnings field to given value.


### GetNumFailures

`func (o *HypervisorTestResponseModel) GetNumFailures() int32`

GetNumFailures returns the NumFailures field if non-nil, zero value otherwise.

### GetNumFailuresOk

`func (o *HypervisorTestResponseModel) GetNumFailuresOk() (*int32, bool)`

GetNumFailuresOk returns a tuple with the NumFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailures

`func (o *HypervisorTestResponseModel) SetNumFailures(v int32)`

SetNumFailures sets NumFailures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


