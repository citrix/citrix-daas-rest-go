# FilterDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to [**FilterType2**](FilterType2.md) |  | [optional] 
**FilterValue** | Pointer to **NullableString** | Filter value, if available. | [optional] 
**IsMatched** | Pointer to **bool** | Indicate if the filter is a match. | [optional] 
**FilterData** | Pointer to **NullableString** | The filter data. This may need to get translated to a user readable string. | [optional] 
**IsWithout** | Pointer to **NullableBool** | Valid only for the branch repeater filter and access control filter. Indicate the mode or condition. | [optional] 
**Gateway** | Pointer to **NullableString** | Valid only for the access control filter. The access gateway. | [optional] 

## Methods

### NewFilterDetail

`func NewFilterDetail() *FilterDetail`

NewFilterDetail instantiates a new FilterDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterDetailWithDefaults

`func NewFilterDetailWithDefaults() *FilterDetail`

NewFilterDetailWithDefaults instantiates a new FilterDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *FilterDetail) GetFilterType() FilterType2`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *FilterDetail) GetFilterTypeOk() (*FilterType2, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *FilterDetail) SetFilterType(v FilterType2)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *FilterDetail) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetFilterValue

`func (o *FilterDetail) GetFilterValue() string`

GetFilterValue returns the FilterValue field if non-nil, zero value otherwise.

### GetFilterValueOk

`func (o *FilterDetail) GetFilterValueOk() (*string, bool)`

GetFilterValueOk returns a tuple with the FilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValue

`func (o *FilterDetail) SetFilterValue(v string)`

SetFilterValue sets FilterValue field to given value.

### HasFilterValue

`func (o *FilterDetail) HasFilterValue() bool`

HasFilterValue returns a boolean if a field has been set.

### SetFilterValueNil

`func (o *FilterDetail) SetFilterValueNil(b bool)`

 SetFilterValueNil sets the value for FilterValue to be an explicit nil

### UnsetFilterValue
`func (o *FilterDetail) UnsetFilterValue()`

UnsetFilterValue ensures that no value is present for FilterValue, not even an explicit nil
### GetIsMatched

`func (o *FilterDetail) GetIsMatched() bool`

GetIsMatched returns the IsMatched field if non-nil, zero value otherwise.

### GetIsMatchedOk

`func (o *FilterDetail) GetIsMatchedOk() (*bool, bool)`

GetIsMatchedOk returns a tuple with the IsMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatched

`func (o *FilterDetail) SetIsMatched(v bool)`

SetIsMatched sets IsMatched field to given value.

### HasIsMatched

`func (o *FilterDetail) HasIsMatched() bool`

HasIsMatched returns a boolean if a field has been set.

### GetFilterData

`func (o *FilterDetail) GetFilterData() string`

GetFilterData returns the FilterData field if non-nil, zero value otherwise.

### GetFilterDataOk

`func (o *FilterDetail) GetFilterDataOk() (*string, bool)`

GetFilterDataOk returns a tuple with the FilterData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterData

`func (o *FilterDetail) SetFilterData(v string)`

SetFilterData sets FilterData field to given value.

### HasFilterData

`func (o *FilterDetail) HasFilterData() bool`

HasFilterData returns a boolean if a field has been set.

### SetFilterDataNil

`func (o *FilterDetail) SetFilterDataNil(b bool)`

 SetFilterDataNil sets the value for FilterData to be an explicit nil

### UnsetFilterData
`func (o *FilterDetail) UnsetFilterData()`

UnsetFilterData ensures that no value is present for FilterData, not even an explicit nil
### GetIsWithout

`func (o *FilterDetail) GetIsWithout() bool`

GetIsWithout returns the IsWithout field if non-nil, zero value otherwise.

### GetIsWithoutOk

`func (o *FilterDetail) GetIsWithoutOk() (*bool, bool)`

GetIsWithoutOk returns a tuple with the IsWithout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWithout

`func (o *FilterDetail) SetIsWithout(v bool)`

SetIsWithout sets IsWithout field to given value.

### HasIsWithout

`func (o *FilterDetail) HasIsWithout() bool`

HasIsWithout returns a boolean if a field has been set.

### SetIsWithoutNil

`func (o *FilterDetail) SetIsWithoutNil(b bool)`

 SetIsWithoutNil sets the value for IsWithout to be an explicit nil

### UnsetIsWithout
`func (o *FilterDetail) UnsetIsWithout()`

UnsetIsWithout ensures that no value is present for IsWithout, not even an explicit nil
### GetGateway

`func (o *FilterDetail) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *FilterDetail) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *FilterDetail) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *FilterDetail) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *FilterDetail) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *FilterDetail) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


