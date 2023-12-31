# HypervisorFaultResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **NullableString** | Fault state associated with connection, or &#39;None&#39; if OK. | [optional] 
**Reason** | Pointer to **NullableString** | Exception error text associated with any fault state. | [optional] 
**TimeEntered** | Pointer to **NullableString** | Time at which the hypervisor entered fault state | [optional] 
**DurationInSeconds** | Pointer to **NullableFloat64** | Period for which the hypervisor has been in fault state | [optional] 

## Methods

### NewHypervisorFaultResponseModel

`func NewHypervisorFaultResponseModel() *HypervisorFaultResponseModel`

NewHypervisorFaultResponseModel instantiates a new HypervisorFaultResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorFaultResponseModelWithDefaults

`func NewHypervisorFaultResponseModelWithDefaults() *HypervisorFaultResponseModel`

NewHypervisorFaultResponseModelWithDefaults instantiates a new HypervisorFaultResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *HypervisorFaultResponseModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HypervisorFaultResponseModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HypervisorFaultResponseModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HypervisorFaultResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *HypervisorFaultResponseModel) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *HypervisorFaultResponseModel) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetReason

`func (o *HypervisorFaultResponseModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HypervisorFaultResponseModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HypervisorFaultResponseModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HypervisorFaultResponseModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *HypervisorFaultResponseModel) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *HypervisorFaultResponseModel) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetTimeEntered

`func (o *HypervisorFaultResponseModel) GetTimeEntered() string`

GetTimeEntered returns the TimeEntered field if non-nil, zero value otherwise.

### GetTimeEnteredOk

`func (o *HypervisorFaultResponseModel) GetTimeEnteredOk() (*string, bool)`

GetTimeEnteredOk returns a tuple with the TimeEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntered

`func (o *HypervisorFaultResponseModel) SetTimeEntered(v string)`

SetTimeEntered sets TimeEntered field to given value.

### HasTimeEntered

`func (o *HypervisorFaultResponseModel) HasTimeEntered() bool`

HasTimeEntered returns a boolean if a field has been set.

### SetTimeEnteredNil

`func (o *HypervisorFaultResponseModel) SetTimeEnteredNil(b bool)`

 SetTimeEnteredNil sets the value for TimeEntered to be an explicit nil

### UnsetTimeEntered
`func (o *HypervisorFaultResponseModel) UnsetTimeEntered()`

UnsetTimeEntered ensures that no value is present for TimeEntered, not even an explicit nil
### GetDurationInSeconds

`func (o *HypervisorFaultResponseModel) GetDurationInSeconds() float64`

GetDurationInSeconds returns the DurationInSeconds field if non-nil, zero value otherwise.

### GetDurationInSecondsOk

`func (o *HypervisorFaultResponseModel) GetDurationInSecondsOk() (*float64, bool)`

GetDurationInSecondsOk returns a tuple with the DurationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInSeconds

`func (o *HypervisorFaultResponseModel) SetDurationInSeconds(v float64)`

SetDurationInSeconds sets DurationInSeconds field to given value.

### HasDurationInSeconds

`func (o *HypervisorFaultResponseModel) HasDurationInSeconds() bool`

HasDurationInSeconds returns a boolean if a field has been set.

### SetDurationInSecondsNil

`func (o *HypervisorFaultResponseModel) SetDurationInSecondsNil(b bool)`

 SetDurationInSecondsNil sets the value for DurationInSeconds to be an explicit nil

### UnsetDurationInSeconds
`func (o *HypervisorFaultResponseModel) UnsetDurationInSeconds()`

UnsetDurationInSeconds ensures that no value is present for DurationInSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


