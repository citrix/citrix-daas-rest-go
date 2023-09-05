# HypervisorResourcePoolTestResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | [**HypervisorResourcePoolTestResponseModelResourcePool**](HypervisorResourcePoolTestResponseModelResourcePool.md) |  | 
**NumPassed** | **int32** | The number of tests that passed. | 
**NumWarnings** | **int32** | The number of warnings that were found. | 
**NumFailures** | **int32** | The number of tests that failed. | 

## Methods

### NewHypervisorResourcePoolTestResponseModel

`func NewHypervisorResourcePoolTestResponseModel(resourcePool HypervisorResourcePoolTestResponseModelResourcePool, numPassed int32, numWarnings int32, numFailures int32, ) *HypervisorResourcePoolTestResponseModel`

NewHypervisorResourcePoolTestResponseModel instantiates a new HypervisorResourcePoolTestResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolTestResponseModelWithDefaults

`func NewHypervisorResourcePoolTestResponseModelWithDefaults() *HypervisorResourcePoolTestResponseModel`

NewHypervisorResourcePoolTestResponseModelWithDefaults instantiates a new HypervisorResourcePoolTestResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *HypervisorResourcePoolTestResponseModel) GetResourcePool() HypervisorResourcePoolTestResponseModelResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *HypervisorResourcePoolTestResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolTestResponseModelResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *HypervisorResourcePoolTestResponseModel) SetResourcePool(v HypervisorResourcePoolTestResponseModelResourcePool)`

SetResourcePool sets ResourcePool field to given value.


### GetNumPassed

`func (o *HypervisorResourcePoolTestResponseModel) GetNumPassed() int32`

GetNumPassed returns the NumPassed field if non-nil, zero value otherwise.

### GetNumPassedOk

`func (o *HypervisorResourcePoolTestResponseModel) GetNumPassedOk() (*int32, bool)`

GetNumPassedOk returns a tuple with the NumPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPassed

`func (o *HypervisorResourcePoolTestResponseModel) SetNumPassed(v int32)`

SetNumPassed sets NumPassed field to given value.


### GetNumWarnings

`func (o *HypervisorResourcePoolTestResponseModel) GetNumWarnings() int32`

GetNumWarnings returns the NumWarnings field if non-nil, zero value otherwise.

### GetNumWarningsOk

`func (o *HypervisorResourcePoolTestResponseModel) GetNumWarningsOk() (*int32, bool)`

GetNumWarningsOk returns a tuple with the NumWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWarnings

`func (o *HypervisorResourcePoolTestResponseModel) SetNumWarnings(v int32)`

SetNumWarnings sets NumWarnings field to given value.


### GetNumFailures

`func (o *HypervisorResourcePoolTestResponseModel) GetNumFailures() int32`

GetNumFailures returns the NumFailures field if non-nil, zero value otherwise.

### GetNumFailuresOk

`func (o *HypervisorResourcePoolTestResponseModel) GetNumFailuresOk() (*int32, bool)`

GetNumFailuresOk returns a tuple with the NumFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailures

`func (o *HypervisorResourcePoolTestResponseModel) SetNumFailures(v int32)`

SetNumFailures sets NumFailures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


