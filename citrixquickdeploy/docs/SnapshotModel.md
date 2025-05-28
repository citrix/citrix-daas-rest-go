# SnapshotModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Type** | Pointer to [**SnapshotType**](SnapshotType.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**SnapshotName** | Pointer to **string** |  | [optional] 

## Methods

### NewSnapshotModel

`func NewSnapshotModel() *SnapshotModel`

NewSnapshotModel instantiates a new SnapshotModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotModelWithDefaults

`func NewSnapshotModelWithDefaults() *SnapshotModel`

NewSnapshotModelWithDefaults instantiates a new SnapshotModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStamp

`func (o *SnapshotModel) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *SnapshotModel) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *SnapshotModel) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *SnapshotModel) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetType

`func (o *SnapshotModel) GetType() SnapshotType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotModel) GetTypeOk() (*SnapshotType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotModel) SetType(v SnapshotType)`

SetType sets Type field to given value.

### HasType

`func (o *SnapshotModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *SnapshotModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SnapshotModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SnapshotModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SnapshotModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSnapshotName

`func (o *SnapshotModel) GetSnapshotName() string`

GetSnapshotName returns the SnapshotName field if non-nil, zero value otherwise.

### GetSnapshotNameOk

`func (o *SnapshotModel) GetSnapshotNameOk() (*string, bool)`

GetSnapshotNameOk returns a tuple with the SnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotName

`func (o *SnapshotModel) SetSnapshotName(v string)`

SetSnapshotName sets SnapshotName field to given value.

### HasSnapshotName

`func (o *SnapshotModel) HasSnapshotName() bool`

HasSnapshotName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


