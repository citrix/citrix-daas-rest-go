# DeliveryGroupTestResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryGroup** | [**RefResponseModel**](RefResponseModel.md) |  | 
**NumPassed** | **int32** | The number of tests that passed. | 
**NumWarnings** | **int32** | The number of warnings that were found. | 
**NumFailures** | **int32** | The number of tests that failed. | 

## Methods

### NewDeliveryGroupTestResponseModel

`func NewDeliveryGroupTestResponseModel(deliveryGroup RefResponseModel, numPassed int32, numWarnings int32, numFailures int32, ) *DeliveryGroupTestResponseModel`

NewDeliveryGroupTestResponseModel instantiates a new DeliveryGroupTestResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupTestResponseModelWithDefaults

`func NewDeliveryGroupTestResponseModelWithDefaults() *DeliveryGroupTestResponseModel`

NewDeliveryGroupTestResponseModelWithDefaults instantiates a new DeliveryGroupTestResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryGroup

`func (o *DeliveryGroupTestResponseModel) GetDeliveryGroup() RefResponseModel`

GetDeliveryGroup returns the DeliveryGroup field if non-nil, zero value otherwise.

### GetDeliveryGroupOk

`func (o *DeliveryGroupTestResponseModel) GetDeliveryGroupOk() (*RefResponseModel, bool)`

GetDeliveryGroupOk returns a tuple with the DeliveryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroup

`func (o *DeliveryGroupTestResponseModel) SetDeliveryGroup(v RefResponseModel)`

SetDeliveryGroup sets DeliveryGroup field to given value.


### GetNumPassed

`func (o *DeliveryGroupTestResponseModel) GetNumPassed() int32`

GetNumPassed returns the NumPassed field if non-nil, zero value otherwise.

### GetNumPassedOk

`func (o *DeliveryGroupTestResponseModel) GetNumPassedOk() (*int32, bool)`

GetNumPassedOk returns a tuple with the NumPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPassed

`func (o *DeliveryGroupTestResponseModel) SetNumPassed(v int32)`

SetNumPassed sets NumPassed field to given value.


### GetNumWarnings

`func (o *DeliveryGroupTestResponseModel) GetNumWarnings() int32`

GetNumWarnings returns the NumWarnings field if non-nil, zero value otherwise.

### GetNumWarningsOk

`func (o *DeliveryGroupTestResponseModel) GetNumWarningsOk() (*int32, bool)`

GetNumWarningsOk returns a tuple with the NumWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWarnings

`func (o *DeliveryGroupTestResponseModel) SetNumWarnings(v int32)`

SetNumWarnings sets NumWarnings field to given value.


### GetNumFailures

`func (o *DeliveryGroupTestResponseModel) GetNumFailures() int32`

GetNumFailures returns the NumFailures field if non-nil, zero value otherwise.

### GetNumFailuresOk

`func (o *DeliveryGroupTestResponseModel) GetNumFailuresOk() (*int32, bool)`

GetNumFailuresOk returns a tuple with the NumFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailures

`func (o *DeliveryGroupTestResponseModel) SetNumFailures(v int32)`

SetNumFailures sets NumFailures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


