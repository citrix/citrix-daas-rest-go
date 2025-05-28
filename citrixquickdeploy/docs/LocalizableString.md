# LocalizableString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**LocalizedValue** | Pointer to **string** |  | [optional] 

## Methods

### NewLocalizableString

`func NewLocalizableString() *LocalizableString`

NewLocalizableString instantiates a new LocalizableString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalizableStringWithDefaults

`func NewLocalizableStringWithDefaults() *LocalizableString`

NewLocalizableStringWithDefaults instantiates a new LocalizableString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *LocalizableString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LocalizableString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LocalizableString) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *LocalizableString) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLocalizedValue

`func (o *LocalizableString) GetLocalizedValue() string`

GetLocalizedValue returns the LocalizedValue field if non-nil, zero value otherwise.

### GetLocalizedValueOk

`func (o *LocalizableString) GetLocalizedValueOk() (*string, bool)`

GetLocalizedValueOk returns a tuple with the LocalizedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedValue

`func (o *LocalizableString) SetLocalizedValue(v string)`

SetLocalizedValue sets LocalizedValue field to given value.

### HasLocalizedValue

`func (o *LocalizableString) HasLocalizedValue() bool`

HasLocalizedValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


