# FilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **NullableString** | The filter type. | [optional] 
**IsAllowed** | Pointer to **bool** | Filtered policy is allowed or denied if the filter condition is met. | [optional] 
**IsEnabled** | Pointer to **bool** | Filter is enabled or disabled. | [optional] 
**FilterData** | Pointer to **NullableString** | Filter data as a serialized JSON string. This is ignored if the filter is a Citrix SD-Wan filter. | [optional] 

## Methods

### NewFilterRequest

`func NewFilterRequest() *FilterRequest`

NewFilterRequest instantiates a new FilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterRequestWithDefaults

`func NewFilterRequestWithDefaults() *FilterRequest`

NewFilterRequestWithDefaults instantiates a new FilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *FilterRequest) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *FilterRequest) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *FilterRequest) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *FilterRequest) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### SetFilterTypeNil

`func (o *FilterRequest) SetFilterTypeNil(b bool)`

 SetFilterTypeNil sets the value for FilterType to be an explicit nil

### UnsetFilterType
`func (o *FilterRequest) UnsetFilterType()`

UnsetFilterType ensures that no value is present for FilterType, not even an explicit nil
### GetIsAllowed

`func (o *FilterRequest) GetIsAllowed() bool`

GetIsAllowed returns the IsAllowed field if non-nil, zero value otherwise.

### GetIsAllowedOk

`func (o *FilterRequest) GetIsAllowedOk() (*bool, bool)`

GetIsAllowedOk returns a tuple with the IsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowed

`func (o *FilterRequest) SetIsAllowed(v bool)`

SetIsAllowed sets IsAllowed field to given value.

### HasIsAllowed

`func (o *FilterRequest) HasIsAllowed() bool`

HasIsAllowed returns a boolean if a field has been set.

### GetIsEnabled

`func (o *FilterRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FilterRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FilterRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FilterRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetFilterData

`func (o *FilterRequest) GetFilterData() string`

GetFilterData returns the FilterData field if non-nil, zero value otherwise.

### GetFilterDataOk

`func (o *FilterRequest) GetFilterDataOk() (*string, bool)`

GetFilterDataOk returns a tuple with the FilterData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterData

`func (o *FilterRequest) SetFilterData(v string)`

SetFilterData sets FilterData field to given value.

### HasFilterData

`func (o *FilterRequest) HasFilterData() bool`

HasFilterData returns a boolean if a field has been set.

### SetFilterDataNil

`func (o *FilterRequest) SetFilterDataNil(b bool)`

 SetFilterDataNil sets the value for FilterData to be an explicit nil

### UnsetFilterData
`func (o *FilterRequest) UnsetFilterData()`

UnsetFilterData ensures that no value is present for FilterData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


