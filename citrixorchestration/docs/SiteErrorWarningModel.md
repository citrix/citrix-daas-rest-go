# SiteErrorWarningModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**MisConfigurationCategory**](MisConfigurationCategory.md) |  | [optional] 
**Type** | Pointer to **NullableString** | The error or warning message. | [optional] 
**TimeInUtc** | Pointer to **NullableString** | The time of the error or warning in UTC. | [optional] 
**IsHighPriority** | Pointer to **bool** | Indicate whether the error or warning is high priority. | [optional] 

## Methods

### NewSiteErrorWarningModel

`func NewSiteErrorWarningModel() *SiteErrorWarningModel`

NewSiteErrorWarningModel instantiates a new SiteErrorWarningModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteErrorWarningModelWithDefaults

`func NewSiteErrorWarningModelWithDefaults() *SiteErrorWarningModel`

NewSiteErrorWarningModelWithDefaults instantiates a new SiteErrorWarningModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *SiteErrorWarningModel) GetCategory() MisConfigurationCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SiteErrorWarningModel) GetCategoryOk() (*MisConfigurationCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SiteErrorWarningModel) SetCategory(v MisConfigurationCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SiteErrorWarningModel) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *SiteErrorWarningModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteErrorWarningModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteErrorWarningModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SiteErrorWarningModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SiteErrorWarningModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SiteErrorWarningModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTimeInUtc

`func (o *SiteErrorWarningModel) GetTimeInUtc() string`

GetTimeInUtc returns the TimeInUtc field if non-nil, zero value otherwise.

### GetTimeInUtcOk

`func (o *SiteErrorWarningModel) GetTimeInUtcOk() (*string, bool)`

GetTimeInUtcOk returns a tuple with the TimeInUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInUtc

`func (o *SiteErrorWarningModel) SetTimeInUtc(v string)`

SetTimeInUtc sets TimeInUtc field to given value.

### HasTimeInUtc

`func (o *SiteErrorWarningModel) HasTimeInUtc() bool`

HasTimeInUtc returns a boolean if a field has been set.

### SetTimeInUtcNil

`func (o *SiteErrorWarningModel) SetTimeInUtcNil(b bool)`

 SetTimeInUtcNil sets the value for TimeInUtc to be an explicit nil

### UnsetTimeInUtc
`func (o *SiteErrorWarningModel) UnsetTimeInUtc()`

UnsetTimeInUtc ensures that no value is present for TimeInUtc, not even an explicit nil
### GetIsHighPriority

`func (o *SiteErrorWarningModel) GetIsHighPriority() bool`

GetIsHighPriority returns the IsHighPriority field if non-nil, zero value otherwise.

### GetIsHighPriorityOk

`func (o *SiteErrorWarningModel) GetIsHighPriorityOk() (*bool, bool)`

GetIsHighPriorityOk returns a tuple with the IsHighPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighPriority

`func (o *SiteErrorWarningModel) SetIsHighPriority(v bool)`

SetIsHighPriority sets IsHighPriority field to given value.

### HasIsHighPriority

`func (o *SiteErrorWarningModel) HasIsHighPriority() bool`

HasIsHighPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


