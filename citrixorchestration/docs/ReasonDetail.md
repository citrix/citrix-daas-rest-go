# ReasonDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **NullableString** | The filter type. | [optional] 
**Value** | Pointer to **NullableString** | The filter value. | [optional] 
**Match** | Pointer to **bool** | Indicate if the filter matches the given value. | [optional] 

## Methods

### NewReasonDetail

`func NewReasonDetail() *ReasonDetail`

NewReasonDetail instantiates a new ReasonDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReasonDetailWithDefaults

`func NewReasonDetailWithDefaults() *ReasonDetail`

NewReasonDetailWithDefaults instantiates a new ReasonDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ReasonDetail) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ReasonDetail) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ReasonDetail) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ReasonDetail) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ReasonDetail) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ReasonDetail) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetValue

`func (o *ReasonDetail) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReasonDetail) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReasonDetail) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ReasonDetail) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ReasonDetail) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ReasonDetail) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMatch

`func (o *ReasonDetail) GetMatch() bool`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *ReasonDetail) GetMatchOk() (*bool, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *ReasonDetail) SetMatch(v bool)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *ReasonDetail) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


