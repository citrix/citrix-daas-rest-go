# EventDataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to [**SenderAuthorization**](SenderAuthorization.md) |  | [optional] 
**Claims** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Caller** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**EventDataId** | Pointer to **string** |  | [optional] [readonly] 
**CorrelationId** | Pointer to **string** |  | [optional] [readonly] 
**EventName** | Pointer to [**MonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] 
**Category** | Pointer to [**MonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] 
**HttpRequest** | Pointer to [**EventDataHttpRequestInfo**](EventDataHttpRequestInfo.md) |  | [optional] 
**Level** | Pointer to [**MonitorEventLevel**](MonitorEventLevel.md) |  | [optional] 
**ResourceGroupName** | Pointer to **string** |  | [optional] [readonly] 
**ResourceProviderName** | Pointer to [**MonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] 
**ResourceId** | Pointer to [**ResourceIdentifier**](ResourceIdentifier.md) |  | [optional] 
**ResourceType** | Pointer to [**MonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] 
**OperationId** | Pointer to **string** |  | [optional] [readonly] 
**OperationName** | Pointer to [**MonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Status** | Pointer to [**MonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] 
**SubStatus** | Pointer to [**MonitorLocalizableString**](MonitorLocalizableString.md) |  | [optional] 
**EventTimestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**SubmissionTimestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**SubscriptionId** | Pointer to **string** |  | [optional] [readonly] 
**TenantId** | Pointer to **string** |  | [optional] [readonly] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


