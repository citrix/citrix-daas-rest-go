# SetRecommendationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**RecommendationState**](RecommendationState.md) |  | [optional] 
**Reason** | Pointer to **NullableString** | The reason why you want to dismiss the recommendation. | [optional] 

## Methods

### NewSetRecommendationRequestModel

`func NewSetRecommendationRequestModel() *SetRecommendationRequestModel`

NewSetRecommendationRequestModel instantiates a new SetRecommendationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRecommendationRequestModelWithDefaults

`func NewSetRecommendationRequestModelWithDefaults() *SetRecommendationRequestModel`

NewSetRecommendationRequestModelWithDefaults instantiates a new SetRecommendationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SetRecommendationRequestModel) GetState() RecommendationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SetRecommendationRequestModel) GetStateOk() (*RecommendationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SetRecommendationRequestModel) SetState(v RecommendationState)`

SetState sets State field to given value.

### HasState

`func (o *SetRecommendationRequestModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *SetRecommendationRequestModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SetRecommendationRequestModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SetRecommendationRequestModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SetRecommendationRequestModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *SetRecommendationRequestModel) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *SetRecommendationRequestModel) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


