# HypervisorResponseModelAllOfFault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | Fault state associated with connection, or &#39;None&#39; if OK. | [optional] 
**Reason** | Pointer to **string** | Exception error text associated with any fault state. | [optional] 
**TimeEntered** | Pointer to **string** | Time at which the hypervisor entered fault state | [optional] 
**DurationInSeconds** | Pointer to **float64** | Period for which the hypervisor has been in fault state | [optional] 

## Methods

### NewHypervisorResponseModelAllOfFault

`func NewHypervisorResponseModelAllOfFault() *HypervisorResponseModelAllOfFault`

NewHypervisorResponseModelAllOfFault instantiates a new HypervisorResponseModelAllOfFault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResponseModelAllOfFaultWithDefaults

`func NewHypervisorResponseModelAllOfFaultWithDefaults() *HypervisorResponseModelAllOfFault`

NewHypervisorResponseModelAllOfFaultWithDefaults instantiates a new HypervisorResponseModelAllOfFault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *HypervisorResponseModelAllOfFault) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HypervisorResponseModelAllOfFault) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HypervisorResponseModelAllOfFault) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HypervisorResponseModelAllOfFault) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *HypervisorResponseModelAllOfFault) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HypervisorResponseModelAllOfFault) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HypervisorResponseModelAllOfFault) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HypervisorResponseModelAllOfFault) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTimeEntered

`func (o *HypervisorResponseModelAllOfFault) GetTimeEntered() string`

GetTimeEntered returns the TimeEntered field if non-nil, zero value otherwise.

### GetTimeEnteredOk

`func (o *HypervisorResponseModelAllOfFault) GetTimeEnteredOk() (*string, bool)`

GetTimeEnteredOk returns a tuple with the TimeEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntered

`func (o *HypervisorResponseModelAllOfFault) SetTimeEntered(v string)`

SetTimeEntered sets TimeEntered field to given value.

### HasTimeEntered

`func (o *HypervisorResponseModelAllOfFault) HasTimeEntered() bool`

HasTimeEntered returns a boolean if a field has been set.

### GetDurationInSeconds

`func (o *HypervisorResponseModelAllOfFault) GetDurationInSeconds() float64`

GetDurationInSeconds returns the DurationInSeconds field if non-nil, zero value otherwise.

### GetDurationInSecondsOk

`func (o *HypervisorResponseModelAllOfFault) GetDurationInSecondsOk() (*float64, bool)`

GetDurationInSecondsOk returns a tuple with the DurationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInSeconds

`func (o *HypervisorResponseModelAllOfFault) SetDurationInSeconds(v float64)`

SetDurationInSeconds sets DurationInSeconds field to given value.

### HasDurationInSeconds

`func (o *HypervisorResponseModelAllOfFault) HasDurationInSeconds() bool`

HasDurationInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


