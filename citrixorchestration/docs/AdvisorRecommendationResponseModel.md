# AdvisorRecommendationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Gets or sets the unique identifier for the recommendation. | [optional] 
**Recommendation** | Pointer to **NullableString** | Gets or sets the recommendation text. | [optional] 
**Details** | Pointer to **NullableString** | Gets or sets the details of the recommendation. | [optional] 
**Impact** | Pointer to **NullableString** | Gets or sets the impact of the recommendation. | [optional] 
**Aspect** | Pointer to **NullableString** | Gets or sets the aspect of the recommendation. | [optional] 
**Component** | Pointer to **NullableString** | Gets or sets the component related to the recommendation. | [optional] 
**IsRecommendationNeeded** | Pointer to **bool** | Gets or sets if recommendation is triggered. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata | [optional] 
**LastReportedTime** | Pointer to **NullableString** | Gets or sets the last reported time of the recommendation. | [optional] 
**AffectedResources** | Pointer to **map[string]interface{}** | Gets or sets the affected resources by the recommendation. | [optional] 

## Methods

### NewAdvisorRecommendationResponseModel

`func NewAdvisorRecommendationResponseModel() *AdvisorRecommendationResponseModel`

NewAdvisorRecommendationResponseModel instantiates a new AdvisorRecommendationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorRecommendationResponseModelWithDefaults

`func NewAdvisorRecommendationResponseModelWithDefaults() *AdvisorRecommendationResponseModel`

NewAdvisorRecommendationResponseModelWithDefaults instantiates a new AdvisorRecommendationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdvisorRecommendationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisorRecommendationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisorRecommendationResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisorRecommendationResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AdvisorRecommendationResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AdvisorRecommendationResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRecommendation

`func (o *AdvisorRecommendationResponseModel) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *AdvisorRecommendationResponseModel) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *AdvisorRecommendationResponseModel) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *AdvisorRecommendationResponseModel) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### SetRecommendationNil

`func (o *AdvisorRecommendationResponseModel) SetRecommendationNil(b bool)`

 SetRecommendationNil sets the value for Recommendation to be an explicit nil

### UnsetRecommendation
`func (o *AdvisorRecommendationResponseModel) UnsetRecommendation()`

UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
### GetDetails

`func (o *AdvisorRecommendationResponseModel) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisorRecommendationResponseModel) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisorRecommendationResponseModel) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisorRecommendationResponseModel) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *AdvisorRecommendationResponseModel) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *AdvisorRecommendationResponseModel) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetImpact

`func (o *AdvisorRecommendationResponseModel) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *AdvisorRecommendationResponseModel) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *AdvisorRecommendationResponseModel) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *AdvisorRecommendationResponseModel) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *AdvisorRecommendationResponseModel) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *AdvisorRecommendationResponseModel) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetAspect

`func (o *AdvisorRecommendationResponseModel) GetAspect() string`

GetAspect returns the Aspect field if non-nil, zero value otherwise.

### GetAspectOk

`func (o *AdvisorRecommendationResponseModel) GetAspectOk() (*string, bool)`

GetAspectOk returns a tuple with the Aspect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspect

`func (o *AdvisorRecommendationResponseModel) SetAspect(v string)`

SetAspect sets Aspect field to given value.

### HasAspect

`func (o *AdvisorRecommendationResponseModel) HasAspect() bool`

HasAspect returns a boolean if a field has been set.

### SetAspectNil

`func (o *AdvisorRecommendationResponseModel) SetAspectNil(b bool)`

 SetAspectNil sets the value for Aspect to be an explicit nil

### UnsetAspect
`func (o *AdvisorRecommendationResponseModel) UnsetAspect()`

UnsetAspect ensures that no value is present for Aspect, not even an explicit nil
### GetComponent

`func (o *AdvisorRecommendationResponseModel) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AdvisorRecommendationResponseModel) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AdvisorRecommendationResponseModel) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AdvisorRecommendationResponseModel) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### SetComponentNil

`func (o *AdvisorRecommendationResponseModel) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *AdvisorRecommendationResponseModel) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil
### GetIsRecommendationNeeded

`func (o *AdvisorRecommendationResponseModel) GetIsRecommendationNeeded() bool`

GetIsRecommendationNeeded returns the IsRecommendationNeeded field if non-nil, zero value otherwise.

### GetIsRecommendationNeededOk

`func (o *AdvisorRecommendationResponseModel) GetIsRecommendationNeededOk() (*bool, bool)`

GetIsRecommendationNeededOk returns a tuple with the IsRecommendationNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecommendationNeeded

`func (o *AdvisorRecommendationResponseModel) SetIsRecommendationNeeded(v bool)`

SetIsRecommendationNeeded sets IsRecommendationNeeded field to given value.

### HasIsRecommendationNeeded

`func (o *AdvisorRecommendationResponseModel) HasIsRecommendationNeeded() bool`

HasIsRecommendationNeeded returns a boolean if a field has been set.

### GetMetadata

`func (o *AdvisorRecommendationResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AdvisorRecommendationResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AdvisorRecommendationResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AdvisorRecommendationResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AdvisorRecommendationResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AdvisorRecommendationResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetLastReportedTime

`func (o *AdvisorRecommendationResponseModel) GetLastReportedTime() string`

GetLastReportedTime returns the LastReportedTime field if non-nil, zero value otherwise.

### GetLastReportedTimeOk

`func (o *AdvisorRecommendationResponseModel) GetLastReportedTimeOk() (*string, bool)`

GetLastReportedTimeOk returns a tuple with the LastReportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedTime

`func (o *AdvisorRecommendationResponseModel) SetLastReportedTime(v string)`

SetLastReportedTime sets LastReportedTime field to given value.

### HasLastReportedTime

`func (o *AdvisorRecommendationResponseModel) HasLastReportedTime() bool`

HasLastReportedTime returns a boolean if a field has been set.

### SetLastReportedTimeNil

`func (o *AdvisorRecommendationResponseModel) SetLastReportedTimeNil(b bool)`

 SetLastReportedTimeNil sets the value for LastReportedTime to be an explicit nil

### UnsetLastReportedTime
`func (o *AdvisorRecommendationResponseModel) UnsetLastReportedTime()`

UnsetLastReportedTime ensures that no value is present for LastReportedTime, not even an explicit nil
### GetAffectedResources

`func (o *AdvisorRecommendationResponseModel) GetAffectedResources() map[string]interface{}`

GetAffectedResources returns the AffectedResources field if non-nil, zero value otherwise.

### GetAffectedResourcesOk

`func (o *AdvisorRecommendationResponseModel) GetAffectedResourcesOk() (*map[string]interface{}, bool)`

GetAffectedResourcesOk returns a tuple with the AffectedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedResources

`func (o *AdvisorRecommendationResponseModel) SetAffectedResources(v map[string]interface{})`

SetAffectedResources sets AffectedResources field to given value.

### HasAffectedResources

`func (o *AdvisorRecommendationResponseModel) HasAffectedResources() bool`

HasAffectedResources returns a boolean if a field has been set.

### SetAffectedResourcesNil

`func (o *AdvisorRecommendationResponseModel) SetAffectedResourcesNil(b bool)`

 SetAffectedResourcesNil sets the value for AffectedResources to be an explicit nil

### UnsetAffectedResources
`func (o *AdvisorRecommendationResponseModel) UnsetAffectedResources()`

UnsetAffectedResources ensures that no value is present for AffectedResources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


