# ZonedItemResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the zoned item. | 
**Name** | **string** | Name of the zoned item. | 
**Description** | Pointer to **string** | Description of the zoned item. | [optional] 
**ItemType** | [**ZonableItemType**](ZonableItemType.md) |  | 
**Scopes** | Pointer to [**[]ScopeResponseModel**](ScopeResponseModel.md) | Administrative scopes of the zoned item. | [optional] 
**AdditionalPropertiesField** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Additional properties of the zoned item. | [optional] 
**ZoneId** | Pointer to **string** | ID of the zone which this item is in. | [optional] 
**ZoneName** | Pointer to **string** | Name of the zone which this item is in. | [optional] 

## Methods

### NewZonedItemResponseModel

`func NewZonedItemResponseModel(id string, name string, itemType ZonableItemType, ) *ZonedItemResponseModel`

NewZonedItemResponseModel instantiates a new ZonedItemResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonedItemResponseModelWithDefaults

`func NewZonedItemResponseModelWithDefaults() *ZonedItemResponseModel`

NewZonedItemResponseModelWithDefaults instantiates a new ZonedItemResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZonedItemResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZonedItemResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZonedItemResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ZonedItemResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZonedItemResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZonedItemResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ZonedItemResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZonedItemResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZonedItemResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ZonedItemResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItemType

`func (o *ZonedItemResponseModel) GetItemType() ZonableItemType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ZonedItemResponseModel) GetItemTypeOk() (*ZonableItemType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ZonedItemResponseModel) SetItemType(v ZonableItemType)`

SetItemType sets ItemType field to given value.


### GetScopes

`func (o *ZonedItemResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ZonedItemResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ZonedItemResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ZonedItemResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetAdditionalPropertiesField

`func (o *ZonedItemResponseModel) GetAdditionalPropertiesField() []NameValueStringPairModel`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *ZonedItemResponseModel) GetAdditionalPropertiesFieldOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *ZonedItemResponseModel) SetAdditionalPropertiesField(v []NameValueStringPairModel)`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *ZonedItemResponseModel) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetZoneId

`func (o *ZonedItemResponseModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ZonedItemResponseModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ZonedItemResponseModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ZonedItemResponseModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneName

`func (o *ZonedItemResponseModel) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ZonedItemResponseModel) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ZonedItemResponseModel) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ZonedItemResponseModel) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


