# ZonedItemRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the zoned item. | 
**Name** | **string** | Name of the zoned item. | 
**ItemType** | [**ZonableItemType**](ZonableItemType.md) |  | 
**AdditionalPropertiesField** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Additional properties of the zoned item. | [optional] 

## Methods

### NewZonedItemRequestModel

`func NewZonedItemRequestModel(id string, name string, itemType ZonableItemType, ) *ZonedItemRequestModel`

NewZonedItemRequestModel instantiates a new ZonedItemRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonedItemRequestModelWithDefaults

`func NewZonedItemRequestModelWithDefaults() *ZonedItemRequestModel`

NewZonedItemRequestModelWithDefaults instantiates a new ZonedItemRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZonedItemRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZonedItemRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZonedItemRequestModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ZonedItemRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZonedItemRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZonedItemRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetItemType

`func (o *ZonedItemRequestModel) GetItemType() ZonableItemType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ZonedItemRequestModel) GetItemTypeOk() (*ZonableItemType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ZonedItemRequestModel) SetItemType(v ZonableItemType)`

SetItemType sets ItemType field to given value.


### GetAdditionalPropertiesField

`func (o *ZonedItemRequestModel) GetAdditionalPropertiesField() []NameValueStringPairModel`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *ZonedItemRequestModel) GetAdditionalPropertiesFieldOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *ZonedItemRequestModel) SetAdditionalPropertiesField(v []NameValueStringPairModel)`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *ZonedItemRequestModel) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### SetAdditionalPropertiesFieldNil

`func (o *ZonedItemRequestModel) SetAdditionalPropertiesFieldNil(b bool)`

 SetAdditionalPropertiesFieldNil sets the value for AdditionalPropertiesField to be an explicit nil

### UnsetAdditionalPropertiesField
`func (o *ZonedItemRequestModel) UnsetAdditionalPropertiesField()`

UnsetAdditionalPropertiesField ensures that no value is present for AdditionalPropertiesField, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


