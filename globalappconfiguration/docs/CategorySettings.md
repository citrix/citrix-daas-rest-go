# CategorySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**interface**](interface.md) |  | [optional] 

## Methods

### NewCategorySettings

`func NewCategorySettings() *CategorySettings`

NewCategorySettings instantiates a new CategorySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySettingsWithDefaults

`func NewCategorySettingsWithDefaults() *CategorySettings`

NewCategorySettingsWithDefaults instantiates a new CategorySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CategorySettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategorySettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategorySettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CategorySettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *CategorySettings) GetValue() interface`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CategorySettings) GetValueOk() (*interface, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CategorySettings) SetValue(v interface)`

SetValue sets Value field to given value.

### HasValue

`func (o *CategorySettings) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


