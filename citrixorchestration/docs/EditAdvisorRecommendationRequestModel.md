# EditAdvisorRecommendationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dismiss** | Pointer to **NullableBool** | Dismiss the recommendation | [optional] 
**AddedMetadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The added metadata | [optional] 

## Methods

### NewEditAdvisorRecommendationRequestModel

`func NewEditAdvisorRecommendationRequestModel() *EditAdvisorRecommendationRequestModel`

NewEditAdvisorRecommendationRequestModel instantiates a new EditAdvisorRecommendationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAdvisorRecommendationRequestModelWithDefaults

`func NewEditAdvisorRecommendationRequestModelWithDefaults() *EditAdvisorRecommendationRequestModel`

NewEditAdvisorRecommendationRequestModelWithDefaults instantiates a new EditAdvisorRecommendationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDismiss

`func (o *EditAdvisorRecommendationRequestModel) GetDismiss() bool`

GetDismiss returns the Dismiss field if non-nil, zero value otherwise.

### GetDismissOk

`func (o *EditAdvisorRecommendationRequestModel) GetDismissOk() (*bool, bool)`

GetDismissOk returns a tuple with the Dismiss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismiss

`func (o *EditAdvisorRecommendationRequestModel) SetDismiss(v bool)`

SetDismiss sets Dismiss field to given value.

### HasDismiss

`func (o *EditAdvisorRecommendationRequestModel) HasDismiss() bool`

HasDismiss returns a boolean if a field has been set.

### SetDismissNil

`func (o *EditAdvisorRecommendationRequestModel) SetDismissNil(b bool)`

 SetDismissNil sets the value for Dismiss to be an explicit nil

### UnsetDismiss
`func (o *EditAdvisorRecommendationRequestModel) UnsetDismiss()`

UnsetDismiss ensures that no value is present for Dismiss, not even an explicit nil
### GetAddedMetadata

`func (o *EditAdvisorRecommendationRequestModel) GetAddedMetadata() []NameValueStringPairModel`

GetAddedMetadata returns the AddedMetadata field if non-nil, zero value otherwise.

### GetAddedMetadataOk

`func (o *EditAdvisorRecommendationRequestModel) GetAddedMetadataOk() (*[]NameValueStringPairModel, bool)`

GetAddedMetadataOk returns a tuple with the AddedMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedMetadata

`func (o *EditAdvisorRecommendationRequestModel) SetAddedMetadata(v []NameValueStringPairModel)`

SetAddedMetadata sets AddedMetadata field to given value.

### HasAddedMetadata

`func (o *EditAdvisorRecommendationRequestModel) HasAddedMetadata() bool`

HasAddedMetadata returns a boolean if a field has been set.

### SetAddedMetadataNil

`func (o *EditAdvisorRecommendationRequestModel) SetAddedMetadataNil(b bool)`

 SetAddedMetadataNil sets the value for AddedMetadata to be an explicit nil

### UnsetAddedMetadata
`func (o *EditAdvisorRecommendationRequestModel) UnsetAddedMetadata()`

UnsetAddedMetadata ensures that no value is present for AddedMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


