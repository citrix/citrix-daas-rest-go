# OnPremConnectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | ID of the connection | [optional] 
**Type** | Pointer to [**OnPremConnectionType**](OnPremConnectionType.md) | The type of connection | [optional] 
**Name** | Pointer to **NullableString** | Name of the connection | [optional] 
**StartedAt** | Pointer to **NullableTime** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **NullableInt32** | Estimated total time for the job to finish | [optional] 

## Methods

### NewOnPremConnectionModel

`func NewOnPremConnectionModel() *OnPremConnectionModel`

NewOnPremConnectionModel instantiates a new OnPremConnectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnPremConnectionModelWithDefaults

`func NewOnPremConnectionModelWithDefaults() *OnPremConnectionModel`

NewOnPremConnectionModelWithDefaults instantiates a new OnPremConnectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnPremConnectionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnPremConnectionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnPremConnectionModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnPremConnectionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OnPremConnectionModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OnPremConnectionModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *OnPremConnectionModel) GetType() OnPremConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnPremConnectionModel) GetTypeOk() (*OnPremConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnPremConnectionModel) SetType(v OnPremConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *OnPremConnectionModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *OnPremConnectionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnPremConnectionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnPremConnectionModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OnPremConnectionModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OnPremConnectionModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OnPremConnectionModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartedAt

`func (o *OnPremConnectionModel) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *OnPremConnectionModel) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *OnPremConnectionModel) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *OnPremConnectionModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *OnPremConnectionModel) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *OnPremConnectionModel) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEstimatedTimeInMinute

`func (o *OnPremConnectionModel) GetEstimatedTimeInMinute() int32`

GetEstimatedTimeInMinute returns the EstimatedTimeInMinute field if non-nil, zero value otherwise.

### GetEstimatedTimeInMinuteOk

`func (o *OnPremConnectionModel) GetEstimatedTimeInMinuteOk() (*int32, bool)`

GetEstimatedTimeInMinuteOk returns a tuple with the EstimatedTimeInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeInMinute

`func (o *OnPremConnectionModel) SetEstimatedTimeInMinute(v int32)`

SetEstimatedTimeInMinute sets EstimatedTimeInMinute field to given value.

### HasEstimatedTimeInMinute

`func (o *OnPremConnectionModel) HasEstimatedTimeInMinute() bool`

HasEstimatedTimeInMinute returns a boolean if a field has been set.

### SetEstimatedTimeInMinuteNil

`func (o *OnPremConnectionModel) SetEstimatedTimeInMinuteNil(b bool)`

 SetEstimatedTimeInMinuteNil sets the value for EstimatedTimeInMinute to be an explicit nil

### UnsetEstimatedTimeInMinute
`func (o *OnPremConnectionModel) UnsetEstimatedTimeInMinute()`

UnsetEstimatedTimeInMinute ensures that no value is present for EstimatedTimeInMinute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


