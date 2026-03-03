# EventDataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to [**NullableSenderAuthorization**](SenderAuthorization.md) |  | [optional] [readonly] 
**Claims** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Caller** | Pointer to **NullableString** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] [readonly] 
**Id** | Pointer to **NullableString** |  | [optional] [readonly] 
**EventDataId** | Pointer to **NullableString** |  | [optional] [readonly] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] [readonly] 
**EventName** | Pointer to [**NullableMonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] [readonly] 
**Category** | Pointer to [**NullableMonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] [readonly] 
**HttpRequest** | Pointer to [**NullableEventDataHttpRequestInfo**](EventDataHttpRequestInfo.md) |  | [optional] [readonly] 
**Level** | Pointer to [**NullableMonitorEventLevel**](MonitorEventLevel.md) |  | [optional] [readonly] 
**ResourceGroupName** | Pointer to **NullableString** |  | [optional] [readonly] 
**ResourceProviderName** | Pointer to [**NullableMonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] [readonly] 
**ResourceId** | Pointer to [**NullableResourceIdentifier**](ResourceIdentifier.md) |  | [optional] [readonly] 
**ResourceType** | Pointer to [**NullableMonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] [readonly] 
**OperationId** | Pointer to **NullableString** |  | [optional] [readonly] 
**OperationName** | Pointer to [**NullableMonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] [readonly] 
**Properties** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Status** | Pointer to [**NullableMonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] [readonly] 
**SubStatus** | Pointer to [**NullableMonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] [readonly] 
**EventTimestamp** | Pointer to **NullableTime** |  | [optional] [readonly] 
**SubmissionTimestamp** | Pointer to **NullableTime** |  | [optional] [readonly] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewEventDataInfo

`func NewEventDataInfo() *EventDataInfo`

NewEventDataInfo instantiates a new EventDataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataInfoWithDefaults

`func NewEventDataInfoWithDefaults() *EventDataInfo`

NewEventDataInfoWithDefaults instantiates a new EventDataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *EventDataInfo) GetAuthorization() SenderAuthorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *EventDataInfo) GetAuthorizationOk() (*SenderAuthorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *EventDataInfo) SetAuthorization(v SenderAuthorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *EventDataInfo) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### SetAuthorizationNil

`func (o *EventDataInfo) SetAuthorizationNil(b bool)`

 SetAuthorizationNil sets the value for Authorization to be an explicit nil

### UnsetAuthorization
`func (o *EventDataInfo) UnsetAuthorization()`

UnsetAuthorization ensures that no value is present for Authorization, not even an explicit nil
### GetClaims

`func (o *EventDataInfo) GetClaims() map[string]string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *EventDataInfo) GetClaimsOk() (*map[string]string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *EventDataInfo) SetClaims(v map[string]string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *EventDataInfo) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *EventDataInfo) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *EventDataInfo) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil
### GetCaller

`func (o *EventDataInfo) GetCaller() string`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *EventDataInfo) GetCallerOk() (*string, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *EventDataInfo) SetCaller(v string)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *EventDataInfo) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *EventDataInfo) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *EventDataInfo) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetDescription

`func (o *EventDataInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventDataInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventDataInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventDataInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EventDataInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EventDataInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *EventDataInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventDataInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventDataInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventDataInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EventDataInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EventDataInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetEventDataId

`func (o *EventDataInfo) GetEventDataId() string`

GetEventDataId returns the EventDataId field if non-nil, zero value otherwise.

### GetEventDataIdOk

`func (o *EventDataInfo) GetEventDataIdOk() (*string, bool)`

GetEventDataIdOk returns a tuple with the EventDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDataId

`func (o *EventDataInfo) SetEventDataId(v string)`

SetEventDataId sets EventDataId field to given value.

### HasEventDataId

`func (o *EventDataInfo) HasEventDataId() bool`

HasEventDataId returns a boolean if a field has been set.

### SetEventDataIdNil

`func (o *EventDataInfo) SetEventDataIdNil(b bool)`

 SetEventDataIdNil sets the value for EventDataId to be an explicit nil

### UnsetEventDataId
`func (o *EventDataInfo) UnsetEventDataId()`

UnsetEventDataId ensures that no value is present for EventDataId, not even an explicit nil
### GetCorrelationId

`func (o *EventDataInfo) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *EventDataInfo) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *EventDataInfo) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *EventDataInfo) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *EventDataInfo) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *EventDataInfo) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetEventName

`func (o *EventDataInfo) GetEventName() MonitorLocalizableString`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventDataInfo) GetEventNameOk() (*MonitorLocalizableString, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventDataInfo) SetEventName(v MonitorLocalizableString)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *EventDataInfo) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### SetEventNameNil

`func (o *EventDataInfo) SetEventNameNil(b bool)`

 SetEventNameNil sets the value for EventName to be an explicit nil

### UnsetEventName
`func (o *EventDataInfo) UnsetEventName()`

UnsetEventName ensures that no value is present for EventName, not even an explicit nil
### GetCategory

`func (o *EventDataInfo) GetCategory() MonitorLocalizableString`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EventDataInfo) GetCategoryOk() (*MonitorLocalizableString, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EventDataInfo) SetCategory(v MonitorLocalizableString)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EventDataInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *EventDataInfo) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *EventDataInfo) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetHttpRequest

`func (o *EventDataInfo) GetHttpRequest() EventDataHttpRequestInfo`

GetHttpRequest returns the HttpRequest field if non-nil, zero value otherwise.

### GetHttpRequestOk

`func (o *EventDataInfo) GetHttpRequestOk() (*EventDataHttpRequestInfo, bool)`

GetHttpRequestOk returns a tuple with the HttpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequest

`func (o *EventDataInfo) SetHttpRequest(v EventDataHttpRequestInfo)`

SetHttpRequest sets HttpRequest field to given value.

### HasHttpRequest

`func (o *EventDataInfo) HasHttpRequest() bool`

HasHttpRequest returns a boolean if a field has been set.

### SetHttpRequestNil

`func (o *EventDataInfo) SetHttpRequestNil(b bool)`

 SetHttpRequestNil sets the value for HttpRequest to be an explicit nil

### UnsetHttpRequest
`func (o *EventDataInfo) UnsetHttpRequest()`

UnsetHttpRequest ensures that no value is present for HttpRequest, not even an explicit nil
### GetLevel

`func (o *EventDataInfo) GetLevel() MonitorEventLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *EventDataInfo) GetLevelOk() (*MonitorEventLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *EventDataInfo) SetLevel(v MonitorEventLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *EventDataInfo) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *EventDataInfo) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *EventDataInfo) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetResourceGroupName

`func (o *EventDataInfo) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *EventDataInfo) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *EventDataInfo) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *EventDataInfo) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *EventDataInfo) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *EventDataInfo) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetResourceProviderName

`func (o *EventDataInfo) GetResourceProviderName() MonitorLocalizableString`

GetResourceProviderName returns the ResourceProviderName field if non-nil, zero value otherwise.

### GetResourceProviderNameOk

`func (o *EventDataInfo) GetResourceProviderNameOk() (*MonitorLocalizableString, bool)`

GetResourceProviderNameOk returns a tuple with the ResourceProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProviderName

`func (o *EventDataInfo) SetResourceProviderName(v MonitorLocalizableString)`

SetResourceProviderName sets ResourceProviderName field to given value.

### HasResourceProviderName

`func (o *EventDataInfo) HasResourceProviderName() bool`

HasResourceProviderName returns a boolean if a field has been set.

### SetResourceProviderNameNil

`func (o *EventDataInfo) SetResourceProviderNameNil(b bool)`

 SetResourceProviderNameNil sets the value for ResourceProviderName to be an explicit nil

### UnsetResourceProviderName
`func (o *EventDataInfo) UnsetResourceProviderName()`

UnsetResourceProviderName ensures that no value is present for ResourceProviderName, not even an explicit nil
### GetResourceId

`func (o *EventDataInfo) GetResourceId() ResourceIdentifier`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *EventDataInfo) GetResourceIdOk() (*ResourceIdentifier, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *EventDataInfo) SetResourceId(v ResourceIdentifier)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *EventDataInfo) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *EventDataInfo) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *EventDataInfo) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceType

`func (o *EventDataInfo) GetResourceType() MonitorLocalizableString`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *EventDataInfo) GetResourceTypeOk() (*MonitorLocalizableString, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *EventDataInfo) SetResourceType(v MonitorLocalizableString)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *EventDataInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *EventDataInfo) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *EventDataInfo) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetOperationId

`func (o *EventDataInfo) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *EventDataInfo) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *EventDataInfo) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *EventDataInfo) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### SetOperationIdNil

`func (o *EventDataInfo) SetOperationIdNil(b bool)`

 SetOperationIdNil sets the value for OperationId to be an explicit nil

### UnsetOperationId
`func (o *EventDataInfo) UnsetOperationId()`

UnsetOperationId ensures that no value is present for OperationId, not even an explicit nil
### GetOperationName

`func (o *EventDataInfo) GetOperationName() MonitorLocalizableString`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *EventDataInfo) GetOperationNameOk() (*MonitorLocalizableString, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *EventDataInfo) SetOperationName(v MonitorLocalizableString)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *EventDataInfo) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### SetOperationNameNil

`func (o *EventDataInfo) SetOperationNameNil(b bool)`

 SetOperationNameNil sets the value for OperationName to be an explicit nil

### UnsetOperationName
`func (o *EventDataInfo) UnsetOperationName()`

UnsetOperationName ensures that no value is present for OperationName, not even an explicit nil
### GetProperties

`func (o *EventDataInfo) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventDataInfo) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventDataInfo) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EventDataInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *EventDataInfo) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *EventDataInfo) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetStatus

`func (o *EventDataInfo) GetStatus() MonitorLocalizableString`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventDataInfo) GetStatusOk() (*MonitorLocalizableString, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventDataInfo) SetStatus(v MonitorLocalizableString)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventDataInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EventDataInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EventDataInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubStatus

`func (o *EventDataInfo) GetSubStatus() MonitorLocalizableString`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *EventDataInfo) GetSubStatusOk() (*MonitorLocalizableString, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *EventDataInfo) SetSubStatus(v MonitorLocalizableString)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *EventDataInfo) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### SetSubStatusNil

`func (o *EventDataInfo) SetSubStatusNil(b bool)`

 SetSubStatusNil sets the value for SubStatus to be an explicit nil

### UnsetSubStatus
`func (o *EventDataInfo) UnsetSubStatus()`

UnsetSubStatus ensures that no value is present for SubStatus, not even an explicit nil
### GetEventTimestamp

`func (o *EventDataInfo) GetEventTimestamp() time.Time`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *EventDataInfo) GetEventTimestampOk() (*time.Time, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *EventDataInfo) SetEventTimestamp(v time.Time)`

SetEventTimestamp sets EventTimestamp field to given value.

### HasEventTimestamp

`func (o *EventDataInfo) HasEventTimestamp() bool`

HasEventTimestamp returns a boolean if a field has been set.

### SetEventTimestampNil

`func (o *EventDataInfo) SetEventTimestampNil(b bool)`

 SetEventTimestampNil sets the value for EventTimestamp to be an explicit nil

### UnsetEventTimestamp
`func (o *EventDataInfo) UnsetEventTimestamp()`

UnsetEventTimestamp ensures that no value is present for EventTimestamp, not even an explicit nil
### GetSubmissionTimestamp

`func (o *EventDataInfo) GetSubmissionTimestamp() time.Time`

GetSubmissionTimestamp returns the SubmissionTimestamp field if non-nil, zero value otherwise.

### GetSubmissionTimestampOk

`func (o *EventDataInfo) GetSubmissionTimestampOk() (*time.Time, bool)`

GetSubmissionTimestampOk returns a tuple with the SubmissionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTimestamp

`func (o *EventDataInfo) SetSubmissionTimestamp(v time.Time)`

SetSubmissionTimestamp sets SubmissionTimestamp field to given value.

### HasSubmissionTimestamp

`func (o *EventDataInfo) HasSubmissionTimestamp() bool`

HasSubmissionTimestamp returns a boolean if a field has been set.

### SetSubmissionTimestampNil

`func (o *EventDataInfo) SetSubmissionTimestampNil(b bool)`

 SetSubmissionTimestampNil sets the value for SubmissionTimestamp to be an explicit nil

### UnsetSubmissionTimestamp
`func (o *EventDataInfo) UnsetSubmissionTimestamp()`

UnsetSubmissionTimestamp ensures that no value is present for SubmissionTimestamp, not even an explicit nil
### GetSubscriptionId

`func (o *EventDataInfo) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *EventDataInfo) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *EventDataInfo) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *EventDataInfo) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *EventDataInfo) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *EventDataInfo) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetTenantId

`func (o *EventDataInfo) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EventDataInfo) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EventDataInfo) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *EventDataInfo) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *EventDataInfo) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *EventDataInfo) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


