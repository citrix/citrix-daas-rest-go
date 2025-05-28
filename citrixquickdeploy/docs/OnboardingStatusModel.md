# OnboardingStatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**OnboardingState**](OnboardingState.md) | Current state of onboarind process | [optional] 
**FailureReason** | Pointer to [**OnboardingFailureReason**](OnboardingFailureReason.md) | Reason codes that caused onboarding to fail | [optional] 
**ErrorDetails** | Pointer to **string** | Failure details | [optional] 

## Methods

### NewOnboardingStatusModel

`func NewOnboardingStatusModel() *OnboardingStatusModel`

NewOnboardingStatusModel instantiates a new OnboardingStatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingStatusModelWithDefaults

`func NewOnboardingStatusModelWithDefaults() *OnboardingStatusModel`

NewOnboardingStatusModelWithDefaults instantiates a new OnboardingStatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *OnboardingStatusModel) GetState() OnboardingState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OnboardingStatusModel) GetStateOk() (*OnboardingState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OnboardingStatusModel) SetState(v OnboardingState)`

SetState sets State field to given value.

### HasState

`func (o *OnboardingStatusModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetFailureReason

`func (o *OnboardingStatusModel) GetFailureReason() OnboardingFailureReason`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *OnboardingStatusModel) GetFailureReasonOk() (*OnboardingFailureReason, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *OnboardingStatusModel) SetFailureReason(v OnboardingFailureReason)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *OnboardingStatusModel) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetErrorDetails

`func (o *OnboardingStatusModel) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *OnboardingStatusModel) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *OnboardingStatusModel) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *OnboardingStatusModel) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


