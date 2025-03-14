# AdvisorRecommendationResponseModelCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AdvisorRecommendationResponseModel**](AdvisorRecommendationResponseModel.md) | List of items. | 
**ContinuationToken** | Pointer to **NullableString** | If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **NullableInt32** | Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to &#x60;$search&#x60; APIs. | [optional] 
**JobsInProgress** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | List of jobs currently in progress. | [optional] 

## Methods

### NewAdvisorRecommendationResponseModelCollection

`func NewAdvisorRecommendationResponseModelCollection(items []AdvisorRecommendationResponseModel, ) *AdvisorRecommendationResponseModelCollection`

NewAdvisorRecommendationResponseModelCollection instantiates a new AdvisorRecommendationResponseModelCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorRecommendationResponseModelCollectionWithDefaults

`func NewAdvisorRecommendationResponseModelCollectionWithDefaults() *AdvisorRecommendationResponseModelCollection`

NewAdvisorRecommendationResponseModelCollectionWithDefaults instantiates a new AdvisorRecommendationResponseModelCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AdvisorRecommendationResponseModelCollection) GetItems() []AdvisorRecommendationResponseModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AdvisorRecommendationResponseModelCollection) GetItemsOk() (*[]AdvisorRecommendationResponseModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AdvisorRecommendationResponseModelCollection) SetItems(v []AdvisorRecommendationResponseModel)`

SetItems sets Items field to given value.


### GetContinuationToken

`func (o *AdvisorRecommendationResponseModelCollection) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *AdvisorRecommendationResponseModelCollection) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *AdvisorRecommendationResponseModelCollection) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *AdvisorRecommendationResponseModelCollection) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *AdvisorRecommendationResponseModelCollection) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *AdvisorRecommendationResponseModelCollection) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
### GetTotalItems

`func (o *AdvisorRecommendationResponseModelCollection) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *AdvisorRecommendationResponseModelCollection) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *AdvisorRecommendationResponseModelCollection) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *AdvisorRecommendationResponseModelCollection) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *AdvisorRecommendationResponseModelCollection) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *AdvisorRecommendationResponseModelCollection) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
### GetJobsInProgress

`func (o *AdvisorRecommendationResponseModelCollection) GetJobsInProgress() []RefResponseModel`

GetJobsInProgress returns the JobsInProgress field if non-nil, zero value otherwise.

### GetJobsInProgressOk

`func (o *AdvisorRecommendationResponseModelCollection) GetJobsInProgressOk() (*[]RefResponseModel, bool)`

GetJobsInProgressOk returns a tuple with the JobsInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsInProgress

`func (o *AdvisorRecommendationResponseModelCollection) SetJobsInProgress(v []RefResponseModel)`

SetJobsInProgress sets JobsInProgress field to given value.

### HasJobsInProgress

`func (o *AdvisorRecommendationResponseModelCollection) HasJobsInProgress() bool`

HasJobsInProgress returns a boolean if a field has been set.

### SetJobsInProgressNil

`func (o *AdvisorRecommendationResponseModelCollection) SetJobsInProgressNil(b bool)`

 SetJobsInProgressNil sets the value for JobsInProgress to be an explicit nil

### UnsetJobsInProgress
`func (o *AdvisorRecommendationResponseModelCollection) UnsetJobsInProgress()`

UnsetJobsInProgress ensures that no value is present for JobsInProgress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


