# ZonedItemSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | Pointer to **string** | Id of the zone in which to search for the zoned item. | [optional] 
**ItemType** | Pointer to [**ZonableItemType**](ZonableItemType.md) |  | [optional] 
**BasicSearchString** | Pointer to **string** | Basic search string. Specify a string which will match if contained within name of the zoned item. | [optional] 

## Methods

### NewZonedItemSearchRequestModel

`func NewZonedItemSearchRequestModel() *ZonedItemSearchRequestModel`

NewZonedItemSearchRequestModel instantiates a new ZonedItemSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonedItemSearchRequestModelWithDefaults

`func NewZonedItemSearchRequestModelWithDefaults() *ZonedItemSearchRequestModel`

NewZonedItemSearchRequestModelWithDefaults instantiates a new ZonedItemSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *ZonedItemSearchRequestModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ZonedItemSearchRequestModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ZonedItemSearchRequestModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ZonedItemSearchRequestModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetItemType

`func (o *ZonedItemSearchRequestModel) GetItemType() ZonableItemType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ZonedItemSearchRequestModel) GetItemTypeOk() (*ZonableItemType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ZonedItemSearchRequestModel) SetItemType(v ZonableItemType)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ZonedItemSearchRequestModel) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetBasicSearchString

`func (o *ZonedItemSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *ZonedItemSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *ZonedItemSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *ZonedItemSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


