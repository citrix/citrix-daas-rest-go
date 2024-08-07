# ZoneErrorWarningModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | Pointer to **NullableString** | Zone id. | [optional] 
**ZoneName** | Pointer to **NullableString** | Zone name. | [optional] 
**Category** | Pointer to [**MisConfigurationCategory**](MisConfigurationCategory.md) |  | [optional] 
**Type** | Pointer to **NullableString** | The error or warning message. | [optional] 
**TimeInUtc** | Pointer to **NullableString** | The time of the error or warning in UTC. | [optional] 
**IsHighPriority** | Pointer to **bool** | Indicate whether the error or warning is high priority. | [optional] 

## Methods

### NewZoneErrorWarningModel

`func NewZoneErrorWarningModel() *ZoneErrorWarningModel`

NewZoneErrorWarningModel instantiates a new ZoneErrorWarningModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneErrorWarningModelWithDefaults

`func NewZoneErrorWarningModelWithDefaults() *ZoneErrorWarningModel`

NewZoneErrorWarningModelWithDefaults instantiates a new ZoneErrorWarningModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *ZoneErrorWarningModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ZoneErrorWarningModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ZoneErrorWarningModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ZoneErrorWarningModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *ZoneErrorWarningModel) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *ZoneErrorWarningModel) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetZoneName

`func (o *ZoneErrorWarningModel) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ZoneErrorWarningModel) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ZoneErrorWarningModel) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ZoneErrorWarningModel) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### SetZoneNameNil

`func (o *ZoneErrorWarningModel) SetZoneNameNil(b bool)`

 SetZoneNameNil sets the value for ZoneName to be an explicit nil

### UnsetZoneName
`func (o *ZoneErrorWarningModel) UnsetZoneName()`

UnsetZoneName ensures that no value is present for ZoneName, not even an explicit nil
### GetCategory

`func (o *ZoneErrorWarningModel) GetCategory() MisConfigurationCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ZoneErrorWarningModel) GetCategoryOk() (*MisConfigurationCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ZoneErrorWarningModel) SetCategory(v MisConfigurationCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ZoneErrorWarningModel) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *ZoneErrorWarningModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneErrorWarningModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneErrorWarningModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneErrorWarningModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ZoneErrorWarningModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ZoneErrorWarningModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTimeInUtc

`func (o *ZoneErrorWarningModel) GetTimeInUtc() string`

GetTimeInUtc returns the TimeInUtc field if non-nil, zero value otherwise.

### GetTimeInUtcOk

`func (o *ZoneErrorWarningModel) GetTimeInUtcOk() (*string, bool)`

GetTimeInUtcOk returns a tuple with the TimeInUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInUtc

`func (o *ZoneErrorWarningModel) SetTimeInUtc(v string)`

SetTimeInUtc sets TimeInUtc field to given value.

### HasTimeInUtc

`func (o *ZoneErrorWarningModel) HasTimeInUtc() bool`

HasTimeInUtc returns a boolean if a field has been set.

### SetTimeInUtcNil

`func (o *ZoneErrorWarningModel) SetTimeInUtcNil(b bool)`

 SetTimeInUtcNil sets the value for TimeInUtc to be an explicit nil

### UnsetTimeInUtc
`func (o *ZoneErrorWarningModel) UnsetTimeInUtc()`

UnsetTimeInUtc ensures that no value is present for TimeInUtc, not even an explicit nil
### GetIsHighPriority

`func (o *ZoneErrorWarningModel) GetIsHighPriority() bool`

GetIsHighPriority returns the IsHighPriority field if non-nil, zero value otherwise.

### GetIsHighPriorityOk

`func (o *ZoneErrorWarningModel) GetIsHighPriorityOk() (*bool, bool)`

GetIsHighPriorityOk returns a tuple with the IsHighPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighPriority

`func (o *ZoneErrorWarningModel) SetIsHighPriority(v bool)`

SetIsHighPriority sets IsHighPriority field to given value.

### HasIsHighPriority

`func (o *ZoneErrorWarningModel) HasIsHighPriority() bool`

HasIsHighPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


