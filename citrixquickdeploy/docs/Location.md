# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] [readonly] 
**SubscriptionId** | Pointer to **NullableString** | The subscription ID. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The location name. | [optional] [readonly] 
**Type** | Pointer to [**NullableLocationType**](LocationType.md) | The location type. | [optional] [readonly] 
**DisplayName** | Pointer to **NullableString** | The display name of the location. | [optional] [readonly] 
**RegionalDisplayName** | Pointer to **NullableString** | The display name of the location and its region. | [optional] [readonly] 
**Metadata** | Pointer to [**NullableLocationMetaData**](LocationMetaData.md) | Metadata of the location, such as lat/long, paired region, and others. | [optional] [readonly] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Location) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Location) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Location) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSubscriptionId

`func (o *Location) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Location) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Location) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Location) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *Location) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *Location) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Location) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Location) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Location) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *Location) GetType() LocationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Location) GetTypeOk() (*LocationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Location) SetType(v LocationType)`

SetType sets Type field to given value.

### HasType

`func (o *Location) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Location) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Location) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDisplayName

`func (o *Location) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Location) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Location) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Location) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Location) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Location) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRegionalDisplayName

`func (o *Location) GetRegionalDisplayName() string`

GetRegionalDisplayName returns the RegionalDisplayName field if non-nil, zero value otherwise.

### GetRegionalDisplayNameOk

`func (o *Location) GetRegionalDisplayNameOk() (*string, bool)`

GetRegionalDisplayNameOk returns a tuple with the RegionalDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalDisplayName

`func (o *Location) SetRegionalDisplayName(v string)`

SetRegionalDisplayName sets RegionalDisplayName field to given value.

### HasRegionalDisplayName

`func (o *Location) HasRegionalDisplayName() bool`

HasRegionalDisplayName returns a boolean if a field has been set.

### SetRegionalDisplayNameNil

`func (o *Location) SetRegionalDisplayNameNil(b bool)`

 SetRegionalDisplayNameNil sets the value for RegionalDisplayName to be an explicit nil

### UnsetRegionalDisplayName
`func (o *Location) UnsetRegionalDisplayName()`

UnsetRegionalDisplayName ensures that no value is present for RegionalDisplayName, not even an explicit nil
### GetMetadata

`func (o *Location) GetMetadata() LocationMetaData`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Location) GetMetadataOk() (*LocationMetaData, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Location) SetMetadata(v LocationMetaData)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Location) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Location) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Location) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


