# CasEventRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | Cas Event Type Header | 
**Payloads** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Information that need to send to cas, which will shown in cas as payloads. | [optional] 

## Methods

### NewCasEventRequestModel

`func NewCasEventRequestModel(eventType string, ) *CasEventRequestModel`

NewCasEventRequestModel instantiates a new CasEventRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCasEventRequestModelWithDefaults

`func NewCasEventRequestModelWithDefaults() *CasEventRequestModel`

NewCasEventRequestModelWithDefaults instantiates a new CasEventRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CasEventRequestModel) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CasEventRequestModel) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CasEventRequestModel) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetPayloads

`func (o *CasEventRequestModel) GetPayloads() []NameValueStringPairModel`

GetPayloads returns the Payloads field if non-nil, zero value otherwise.

### GetPayloadsOk

`func (o *CasEventRequestModel) GetPayloadsOk() (*[]NameValueStringPairModel, bool)`

GetPayloadsOk returns a tuple with the Payloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloads

`func (o *CasEventRequestModel) SetPayloads(v []NameValueStringPairModel)`

SetPayloads sets Payloads field to given value.

### HasPayloads

`func (o *CasEventRequestModel) HasPayloads() bool`

HasPayloads returns a boolean if a field has been set.

### SetPayloadsNil

`func (o *CasEventRequestModel) SetPayloadsNil(b bool)`

 SetPayloadsNil sets the value for Payloads to be an explicit nil

### UnsetPayloads
`func (o *CasEventRequestModel) UnsetPayloads()`

UnsetPayloads ensures that no value is present for Payloads, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


