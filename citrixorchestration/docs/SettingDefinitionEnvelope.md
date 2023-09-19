# SettingDefinitionEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SettingDefinition**](SettingDefinition.md) | Items of collection. | [optional] 
**ContinuationToken** | Pointer to **NullableString** | Continuation token if paging is requested. This is a compressed JSON string of the SettingSearchFilters object. | [optional] 

## Methods

### NewSettingDefinitionEnvelope

`func NewSettingDefinitionEnvelope() *SettingDefinitionEnvelope`

NewSettingDefinitionEnvelope instantiates a new SettingDefinitionEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingDefinitionEnvelopeWithDefaults

`func NewSettingDefinitionEnvelopeWithDefaults() *SettingDefinitionEnvelope`

NewSettingDefinitionEnvelopeWithDefaults instantiates a new SettingDefinitionEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SettingDefinitionEnvelope) GetItems() []SettingDefinition`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SettingDefinitionEnvelope) GetItemsOk() (*[]SettingDefinition, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SettingDefinitionEnvelope) SetItems(v []SettingDefinition)`

SetItems sets Items field to given value.

### HasItems

`func (o *SettingDefinitionEnvelope) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *SettingDefinitionEnvelope) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *SettingDefinitionEnvelope) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetContinuationToken

`func (o *SettingDefinitionEnvelope) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *SettingDefinitionEnvelope) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *SettingDefinitionEnvelope) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *SettingDefinitionEnvelope) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *SettingDefinitionEnvelope) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *SettingDefinitionEnvelope) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


