# FilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyGuid** | Pointer to **string** | Policy GUID. | [optional] 
**FilterType** | Pointer to **NullableString** | The filter type. | [optional] 
**FilterGuid** | Pointer to **string** | The filter Guid. | [optional] 
**IsAllowed** | Pointer to **bool** | Allow or deny | [optional] 
**IsEnabled** | Pointer to **bool** | Enabled or disabled | [optional] 
**FilterData** | Pointer to **NullableString** | Serialized JSON string. | [optional] 

## Methods

### NewFilterResponse

`func NewFilterResponse() *FilterResponse`

NewFilterResponse instantiates a new FilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterResponseWithDefaults

`func NewFilterResponseWithDefaults() *FilterResponse`

NewFilterResponseWithDefaults instantiates a new FilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyGuid

`func (o *FilterResponse) GetPolicyGuid() string`

GetPolicyGuid returns the PolicyGuid field if non-nil, zero value otherwise.

### GetPolicyGuidOk

`func (o *FilterResponse) GetPolicyGuidOk() (*string, bool)`

GetPolicyGuidOk returns a tuple with the PolicyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGuid

`func (o *FilterResponse) SetPolicyGuid(v string)`

SetPolicyGuid sets PolicyGuid field to given value.

### HasPolicyGuid

`func (o *FilterResponse) HasPolicyGuid() bool`

HasPolicyGuid returns a boolean if a field has been set.

### GetFilterType

`func (o *FilterResponse) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *FilterResponse) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *FilterResponse) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *FilterResponse) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### SetFilterTypeNil

`func (o *FilterResponse) SetFilterTypeNil(b bool)`

 SetFilterTypeNil sets the value for FilterType to be an explicit nil

### UnsetFilterType
`func (o *FilterResponse) UnsetFilterType()`

UnsetFilterType ensures that no value is present for FilterType, not even an explicit nil
### GetFilterGuid

`func (o *FilterResponse) GetFilterGuid() string`

GetFilterGuid returns the FilterGuid field if non-nil, zero value otherwise.

### GetFilterGuidOk

`func (o *FilterResponse) GetFilterGuidOk() (*string, bool)`

GetFilterGuidOk returns a tuple with the FilterGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGuid

`func (o *FilterResponse) SetFilterGuid(v string)`

SetFilterGuid sets FilterGuid field to given value.

### HasFilterGuid

`func (o *FilterResponse) HasFilterGuid() bool`

HasFilterGuid returns a boolean if a field has been set.

### GetIsAllowed

`func (o *FilterResponse) GetIsAllowed() bool`

GetIsAllowed returns the IsAllowed field if non-nil, zero value otherwise.

### GetIsAllowedOk

`func (o *FilterResponse) GetIsAllowedOk() (*bool, bool)`

GetIsAllowedOk returns a tuple with the IsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowed

`func (o *FilterResponse) SetIsAllowed(v bool)`

SetIsAllowed sets IsAllowed field to given value.

### HasIsAllowed

`func (o *FilterResponse) HasIsAllowed() bool`

HasIsAllowed returns a boolean if a field has been set.

### GetIsEnabled

`func (o *FilterResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FilterResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FilterResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FilterResponse) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetFilterData

`func (o *FilterResponse) GetFilterData() string`

GetFilterData returns the FilterData field if non-nil, zero value otherwise.

### GetFilterDataOk

`func (o *FilterResponse) GetFilterDataOk() (*string, bool)`

GetFilterDataOk returns a tuple with the FilterData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterData

`func (o *FilterResponse) SetFilterData(v string)`

SetFilterData sets FilterData field to given value.

### HasFilterData

`func (o *FilterResponse) HasFilterData() bool`

HasFilterData returns a boolean if a field has been set.

### SetFilterDataNil

`func (o *FilterResponse) SetFilterDataNil(b bool)`

 SetFilterDataNil sets the value for FilterData to be an explicit nil

### UnsetFilterData
`func (o *FilterResponse) UnsetFilterData()`

UnsetFilterData ensures that no value is present for FilterData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


