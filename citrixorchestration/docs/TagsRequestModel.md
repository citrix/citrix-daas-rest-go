# TagsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **[]string** | List of references to tags. | [optional] 
**CreateTags** | Pointer to [**[]TagRequestModel**](TagRequestModel.md) | List of tags to both create and associate with the object. | [optional] 
**DeleteTags** | Pointer to **[]string** | List of tags to delete from the site.  This removes the tag not only from the referenced object, but from *all* objects within the site, and also deletes the tag itself. | [optional] 

## Methods

### NewTagsRequestModel

`func NewTagsRequestModel() *TagsRequestModel`

NewTagsRequestModel instantiates a new TagsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsRequestModelWithDefaults

`func NewTagsRequestModelWithDefaults() *TagsRequestModel`

NewTagsRequestModelWithDefaults instantiates a new TagsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TagsRequestModel) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TagsRequestModel) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TagsRequestModel) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *TagsRequestModel) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCreateTags

`func (o *TagsRequestModel) GetCreateTags() []TagRequestModel`

GetCreateTags returns the CreateTags field if non-nil, zero value otherwise.

### GetCreateTagsOk

`func (o *TagsRequestModel) GetCreateTagsOk() (*[]TagRequestModel, bool)`

GetCreateTagsOk returns a tuple with the CreateTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTags

`func (o *TagsRequestModel) SetCreateTags(v []TagRequestModel)`

SetCreateTags sets CreateTags field to given value.

### HasCreateTags

`func (o *TagsRequestModel) HasCreateTags() bool`

HasCreateTags returns a boolean if a field has been set.

### GetDeleteTags

`func (o *TagsRequestModel) GetDeleteTags() []string`

GetDeleteTags returns the DeleteTags field if non-nil, zero value otherwise.

### GetDeleteTagsOk

`func (o *TagsRequestModel) GetDeleteTagsOk() (*[]string, bool)`

GetDeleteTagsOk returns a tuple with the DeleteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteTags

`func (o *TagsRequestModel) SetDeleteTags(v []string)`

SetDeleteTags sets DeleteTags field to given value.

### HasDeleteTags

`func (o *TagsRequestModel) HasDeleteTags() bool`

HasDeleteTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


