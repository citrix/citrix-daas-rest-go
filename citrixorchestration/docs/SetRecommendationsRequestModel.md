# SetRecommendationsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**RecommendationState**](RecommendationState.md) |  | [optional] 
**Reason** | Pointer to **NullableString** | The reason why you want to dismiss the recommendation. | [optional] 
**Ids** | Pointer to **[]string** | The list of the recommendation id. The format should be GUID array. | [optional] 

## Methods

### NewSetRecommendationsRequestModel

`func NewSetRecommendationsRequestModel() *SetRecommendationsRequestModel`

NewSetRecommendationsRequestModel instantiates a new SetRecommendationsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRecommendationsRequestModelWithDefaults

`func NewSetRecommendationsRequestModelWithDefaults() *SetRecommendationsRequestModel`

NewSetRecommendationsRequestModelWithDefaults instantiates a new SetRecommendationsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SetRecommendationsRequestModel) GetState() RecommendationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SetRecommendationsRequestModel) GetStateOk() (*RecommendationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SetRecommendationsRequestModel) SetState(v RecommendationState)`

SetState sets State field to given value.

### HasState

`func (o *SetRecommendationsRequestModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *SetRecommendationsRequestModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SetRecommendationsRequestModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SetRecommendationsRequestModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SetRecommendationsRequestModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *SetRecommendationsRequestModel) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *SetRecommendationsRequestModel) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetIds

`func (o *SetRecommendationsRequestModel) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *SetRecommendationsRequestModel) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *SetRecommendationsRequestModel) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *SetRecommendationsRequestModel) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *SetRecommendationsRequestModel) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *SetRecommendationsRequestModel) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


