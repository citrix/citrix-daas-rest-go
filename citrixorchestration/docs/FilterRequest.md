# FilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** | The filter filter type. | [optional] 
**IsAllowed** | **bool** | Allow or deny | 
**IsEnabled** | **bool** | Enabled or disabled | 
**FilterData** | Pointer to **string** | Serialized JSON string. | [optional] 

## Methods

### NewFilterRequest

`func NewFilterRequest(isAllowed bool, isEnabled bool, ) *FilterRequest`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


