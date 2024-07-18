# ProvisioningOperationEventSearchResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedObjectType** | Pointer to **NullableString** | Linked Object Type. | [optional] 
**LinkedObjectUid** | Pointer to **NullableString** | Linked Object Uid. | [optional] 
**EventRecordId** | Pointer to **int64** | Event Record Id. | [optional] 
**EventCategory** | Pointer to **NullableString** | Event Categaory. | [optional] 
**EventSeverity** | Pointer to **NullableString** | Event Severity. | [optional] 
**EventSource** | Pointer to **NullableString** | Event Source. | [optional] 
**EventState** | Pointer to **NullableString** | Event State. | [optional] 
**EventMessage** | Pointer to **NullableString** | Event Message. | [optional] 
**EventDateTime** | Pointer to **NullableString** | Event Date Time. | [optional] 
**EventAdditionalData** | Pointer to **NullableString** | Event Additional Data. | [optional] 
**OperationName** | Pointer to **NullableString** | Operation Name. | [optional] 
**OperationTargetType** | Pointer to **NullableString** | Operation Target Type. | [optional] 
**OperationTargetName** | Pointer to **NullableString** | Operation Target Name. | [optional] 
**OperationType** | Pointer to **NullableString** | Operation Type. | [optional] 
**Recommendation** | Pointer to **NullableString** | Recommendation. | [optional] 

## Methods

### NewProvisioningOperationEventSearchResponseModel

`func NewProvisioningOperationEventSearchResponseModel() *ProvisioningOperationEventSearchResponseModel`

NewProvisioningOperationEventSearchResponseModel instantiates a new ProvisioningOperationEventSearchResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningOperationEventSearchResponseModelWithDefaults

`func NewProvisioningOperationEventSearchResponseModelWithDefaults() *ProvisioningOperationEventSearchResponseModel`

NewProvisioningOperationEventSearchResponseModelWithDefaults instantiates a new ProvisioningOperationEventSearchResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedObjectType

`func (o *ProvisioningOperationEventSearchResponseModel) GetLinkedObjectType() string`

GetLinkedObjectType returns the LinkedObjectType field if non-nil, zero value otherwise.

### GetLinkedObjectTypeOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetLinkedObjectTypeOk() (*string, bool)`

GetLinkedObjectTypeOk returns a tuple with the LinkedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedObjectType

`func (o *ProvisioningOperationEventSearchResponseModel) SetLinkedObjectType(v string)`

SetLinkedObjectType sets LinkedObjectType field to given value.

### HasLinkedObjectType

`func (o *ProvisioningOperationEventSearchResponseModel) HasLinkedObjectType() bool`

HasLinkedObjectType returns a boolean if a field has been set.

### SetLinkedObjectTypeNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetLinkedObjectTypeNil(b bool)`

 SetLinkedObjectTypeNil sets the value for LinkedObjectType to be an explicit nil

### UnsetLinkedObjectType
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetLinkedObjectType()`

UnsetLinkedObjectType ensures that no value is present for LinkedObjectType, not even an explicit nil
### GetLinkedObjectUid

`func (o *ProvisioningOperationEventSearchResponseModel) GetLinkedObjectUid() string`

GetLinkedObjectUid returns the LinkedObjectUid field if non-nil, zero value otherwise.

### GetLinkedObjectUidOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetLinkedObjectUidOk() (*string, bool)`

GetLinkedObjectUidOk returns a tuple with the LinkedObjectUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedObjectUid

`func (o *ProvisioningOperationEventSearchResponseModel) SetLinkedObjectUid(v string)`

SetLinkedObjectUid sets LinkedObjectUid field to given value.

### HasLinkedObjectUid

`func (o *ProvisioningOperationEventSearchResponseModel) HasLinkedObjectUid() bool`

HasLinkedObjectUid returns a boolean if a field has been set.

### SetLinkedObjectUidNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetLinkedObjectUidNil(b bool)`

 SetLinkedObjectUidNil sets the value for LinkedObjectUid to be an explicit nil

### UnsetLinkedObjectUid
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetLinkedObjectUid()`

UnsetLinkedObjectUid ensures that no value is present for LinkedObjectUid, not even an explicit nil
### GetEventRecordId

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventRecordId() int64`

GetEventRecordId returns the EventRecordId field if non-nil, zero value otherwise.

### GetEventRecordIdOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventRecordIdOk() (*int64, bool)`

GetEventRecordIdOk returns a tuple with the EventRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRecordId

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventRecordId(v int64)`

SetEventRecordId sets EventRecordId field to given value.

### HasEventRecordId

`func (o *ProvisioningOperationEventSearchResponseModel) HasEventRecordId() bool`

HasEventRecordId returns a boolean if a field has been set.

### GetEventCategory

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventCategory() string`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventCategoryOk() (*string, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventCategory(v string)`

SetEventCategory sets EventCategory field to given value.

### HasEventCategory

`func (o *ProvisioningOperationEventSearchResponseModel) HasEventCategory() bool`

HasEventCategory returns a boolean if a field has been set.

### SetEventCategoryNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventCategoryNil(b bool)`

 SetEventCategoryNil sets the value for EventCategory to be an explicit nil

### UnsetEventCategory
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetEventCategory()`

UnsetEventCategory ensures that no value is present for EventCategory, not even an explicit nil
### GetEventSeverity

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventSeverity() string`

GetEventSeverity returns the EventSeverity field if non-nil, zero value otherwise.

### GetEventSeverityOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventSeverityOk() (*string, bool)`

GetEventSeverityOk returns a tuple with the EventSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSeverity

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventSeverity(v string)`

SetEventSeverity sets EventSeverity field to given value.

### HasEventSeverity

`func (o *ProvisioningOperationEventSearchResponseModel) HasEventSeverity() bool`

HasEventSeverity returns a boolean if a field has been set.

### SetEventSeverityNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventSeverityNil(b bool)`

 SetEventSeverityNil sets the value for EventSeverity to be an explicit nil

### UnsetEventSeverity
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetEventSeverity()`

UnsetEventSeverity ensures that no value is present for EventSeverity, not even an explicit nil
### GetEventSource

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *ProvisioningOperationEventSearchResponseModel) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### SetEventSourceNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventSourceNil(b bool)`

 SetEventSourceNil sets the value for EventSource to be an explicit nil

### UnsetEventSource
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetEventSource()`

UnsetEventSource ensures that no value is present for EventSource, not even an explicit nil
### GetEventState

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventState() string`

GetEventState returns the EventState field if non-nil, zero value otherwise.

### GetEventStateOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventStateOk() (*string, bool)`

GetEventStateOk returns a tuple with the EventState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventState

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventState(v string)`

SetEventState sets EventState field to given value.

### HasEventState

`func (o *ProvisioningOperationEventSearchResponseModel) HasEventState() bool`

HasEventState returns a boolean if a field has been set.

### SetEventStateNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventStateNil(b bool)`

 SetEventStateNil sets the value for EventState to be an explicit nil

### UnsetEventState
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetEventState()`

UnsetEventState ensures that no value is present for EventState, not even an explicit nil
### GetEventMessage

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventMessage() string`

GetEventMessage returns the EventMessage field if non-nil, zero value otherwise.

### GetEventMessageOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventMessageOk() (*string, bool)`

GetEventMessageOk returns a tuple with the EventMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessage

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventMessage(v string)`

SetEventMessage sets EventMessage field to given value.

### HasEventMessage

`func (o *ProvisioningOperationEventSearchResponseModel) HasEventMessage() bool`

HasEventMessage returns a boolean if a field has been set.

### SetEventMessageNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventMessageNil(b bool)`

 SetEventMessageNil sets the value for EventMessage to be an explicit nil

### UnsetEventMessage
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetEventMessage()`

UnsetEventMessage ensures that no value is present for EventMessage, not even an explicit nil
### GetEventDateTime

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventDateTime() string`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventDateTimeOk() (*string, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventDateTime(v string)`

SetEventDateTime sets EventDateTime field to given value.

### HasEventDateTime

`func (o *ProvisioningOperationEventSearchResponseModel) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.

### SetEventDateTimeNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventDateTimeNil(b bool)`

 SetEventDateTimeNil sets the value for EventDateTime to be an explicit nil

### UnsetEventDateTime
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetEventDateTime()`

UnsetEventDateTime ensures that no value is present for EventDateTime, not even an explicit nil
### GetEventAdditionalData

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventAdditionalData() string`

GetEventAdditionalData returns the EventAdditionalData field if non-nil, zero value otherwise.

### GetEventAdditionalDataOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetEventAdditionalDataOk() (*string, bool)`

GetEventAdditionalDataOk returns a tuple with the EventAdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAdditionalData

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventAdditionalData(v string)`

SetEventAdditionalData sets EventAdditionalData field to given value.

### HasEventAdditionalData

`func (o *ProvisioningOperationEventSearchResponseModel) HasEventAdditionalData() bool`

HasEventAdditionalData returns a boolean if a field has been set.

### SetEventAdditionalDataNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetEventAdditionalDataNil(b bool)`

 SetEventAdditionalDataNil sets the value for EventAdditionalData to be an explicit nil

### UnsetEventAdditionalData
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetEventAdditionalData()`

UnsetEventAdditionalData ensures that no value is present for EventAdditionalData, not even an explicit nil
### GetOperationName

`func (o *ProvisioningOperationEventSearchResponseModel) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *ProvisioningOperationEventSearchResponseModel) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *ProvisioningOperationEventSearchResponseModel) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### SetOperationNameNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetOperationNameNil(b bool)`

 SetOperationNameNil sets the value for OperationName to be an explicit nil

### UnsetOperationName
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetOperationName()`

UnsetOperationName ensures that no value is present for OperationName, not even an explicit nil
### GetOperationTargetType

`func (o *ProvisioningOperationEventSearchResponseModel) GetOperationTargetType() string`

GetOperationTargetType returns the OperationTargetType field if non-nil, zero value otherwise.

### GetOperationTargetTypeOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetOperationTargetTypeOk() (*string, bool)`

GetOperationTargetTypeOk returns a tuple with the OperationTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTargetType

`func (o *ProvisioningOperationEventSearchResponseModel) SetOperationTargetType(v string)`

SetOperationTargetType sets OperationTargetType field to given value.

### HasOperationTargetType

`func (o *ProvisioningOperationEventSearchResponseModel) HasOperationTargetType() bool`

HasOperationTargetType returns a boolean if a field has been set.

### SetOperationTargetTypeNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetOperationTargetTypeNil(b bool)`

 SetOperationTargetTypeNil sets the value for OperationTargetType to be an explicit nil

### UnsetOperationTargetType
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetOperationTargetType()`

UnsetOperationTargetType ensures that no value is present for OperationTargetType, not even an explicit nil
### GetOperationTargetName

`func (o *ProvisioningOperationEventSearchResponseModel) GetOperationTargetName() string`

GetOperationTargetName returns the OperationTargetName field if non-nil, zero value otherwise.

### GetOperationTargetNameOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetOperationTargetNameOk() (*string, bool)`

GetOperationTargetNameOk returns a tuple with the OperationTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTargetName

`func (o *ProvisioningOperationEventSearchResponseModel) SetOperationTargetName(v string)`

SetOperationTargetName sets OperationTargetName field to given value.

### HasOperationTargetName

`func (o *ProvisioningOperationEventSearchResponseModel) HasOperationTargetName() bool`

HasOperationTargetName returns a boolean if a field has been set.

### SetOperationTargetNameNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetOperationTargetNameNil(b bool)`

 SetOperationTargetNameNil sets the value for OperationTargetName to be an explicit nil

### UnsetOperationTargetName
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetOperationTargetName()`

UnsetOperationTargetName ensures that no value is present for OperationTargetName, not even an explicit nil
### GetOperationType

`func (o *ProvisioningOperationEventSearchResponseModel) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ProvisioningOperationEventSearchResponseModel) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *ProvisioningOperationEventSearchResponseModel) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationTypeNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetOperationTypeNil(b bool)`

 SetOperationTypeNil sets the value for OperationType to be an explicit nil

### UnsetOperationType
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetOperationType()`

UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
### GetRecommendation

`func (o *ProvisioningOperationEventSearchResponseModel) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *ProvisioningOperationEventSearchResponseModel) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *ProvisioningOperationEventSearchResponseModel) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *ProvisioningOperationEventSearchResponseModel) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### SetRecommendationNil

`func (o *ProvisioningOperationEventSearchResponseModel) SetRecommendationNil(b bool)`

 SetRecommendationNil sets the value for Recommendation to be an explicit nil

### UnsetRecommendation
`func (o *ProvisioningOperationEventSearchResponseModel) UnsetRecommendation()`

UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


