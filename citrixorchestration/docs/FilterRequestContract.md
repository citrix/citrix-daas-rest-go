# FilterRequestContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **NullableString** | The filter name, or filter type. | [optional] 
**FilterName** | Pointer to **NullableString** | The filter name. | [optional] 
**IsAllowed** | Pointer to **NullableBool** | Allow or deny | [optional] [default to false]
**IsEnabled** | Pointer to **NullableBool** | Enabled or disabled | [optional] [default to false]
**FilterValue** | Pointer to **NullableString** | Serialized JSON string. | [optional] 

## Methods

### NewFilterRequestContract

`func NewFilterRequestContract() *FilterRequestContract`

NewFilterRequestContract instantiates a new FilterRequestContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterRequestContractWithDefaults

`func NewFilterRequestContractWithDefaults() *FilterRequestContract`

NewFilterRequestContractWithDefaults instantiates a new FilterRequestContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *FilterRequestContract) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *FilterRequestContract) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *FilterRequestContract) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *FilterRequestContract) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### SetFilterTypeNil

`func (o *FilterRequestContract) SetFilterTypeNil(b bool)`

 SetFilterTypeNil sets the value for FilterType to be an explicit nil

### UnsetFilterType
`func (o *FilterRequestContract) UnsetFilterType()`

UnsetFilterType ensures that no value is present for FilterType, not even an explicit nil
### GetFilterName

`func (o *FilterRequestContract) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *FilterRequestContract) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *FilterRequestContract) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.

### HasFilterName

`func (o *FilterRequestContract) HasFilterName() bool`

HasFilterName returns a boolean if a field has been set.

### SetFilterNameNil

`func (o *FilterRequestContract) SetFilterNameNil(b bool)`

 SetFilterNameNil sets the value for FilterName to be an explicit nil

### UnsetFilterName
`func (o *FilterRequestContract) UnsetFilterName()`

UnsetFilterName ensures that no value is present for FilterName, not even an explicit nil
### GetIsAllowed

`func (o *FilterRequestContract) GetIsAllowed() bool`

GetIsAllowed returns the IsAllowed field if non-nil, zero value otherwise.

### GetIsAllowedOk

`func (o *FilterRequestContract) GetIsAllowedOk() (*bool, bool)`

GetIsAllowedOk returns a tuple with the IsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowed

`func (o *FilterRequestContract) SetIsAllowed(v bool)`

SetIsAllowed sets IsAllowed field to given value.

### HasIsAllowed

`func (o *FilterRequestContract) HasIsAllowed() bool`

HasIsAllowed returns a boolean if a field has been set.

### SetIsAllowedNil

`func (o *FilterRequestContract) SetIsAllowedNil(b bool)`

 SetIsAllowedNil sets the value for IsAllowed to be an explicit nil

### UnsetIsAllowed
`func (o *FilterRequestContract) UnsetIsAllowed()`

UnsetIsAllowed ensures that no value is present for IsAllowed, not even an explicit nil
### GetIsEnabled

`func (o *FilterRequestContract) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FilterRequestContract) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FilterRequestContract) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FilterRequestContract) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *FilterRequestContract) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *FilterRequestContract) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetFilterValue

`func (o *FilterRequestContract) GetFilterValue() string`

GetFilterValue returns the FilterValue field if non-nil, zero value otherwise.

### GetFilterValueOk

`func (o *FilterRequestContract) GetFilterValueOk() (*string, bool)`

GetFilterValueOk returns a tuple with the FilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValue

`func (o *FilterRequestContract) SetFilterValue(v string)`

SetFilterValue sets FilterValue field to given value.

### HasFilterValue

`func (o *FilterRequestContract) HasFilterValue() bool`

HasFilterValue returns a boolean if a field has been set.

### SetFilterValueNil

`func (o *FilterRequestContract) SetFilterValueNil(b bool)`

 SetFilterValueNil sets the value for FilterValue to be an explicit nil

### UnsetFilterValue
`func (o *FilterRequestContract) UnsetFilterValue()`

UnsetFilterValue ensures that no value is present for FilterValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


