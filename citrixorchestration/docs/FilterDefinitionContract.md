# FilterDefinitionContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **NullableString** | Filter type | [optional] 
**FilterName** | Pointer to **NullableString** | Localized filter name | [optional] 
**Explanation** | Pointer to **NullableString** | Localized explanation | [optional] 
**IsUserFilter** | Pointer to **bool** | True &#x3D; user filter, False &#x3D; machine filter | [optional] 

## Methods

### NewFilterDefinitionContract

`func NewFilterDefinitionContract() *FilterDefinitionContract`

NewFilterDefinitionContract instantiates a new FilterDefinitionContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterDefinitionContractWithDefaults

`func NewFilterDefinitionContractWithDefaults() *FilterDefinitionContract`

NewFilterDefinitionContractWithDefaults instantiates a new FilterDefinitionContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *FilterDefinitionContract) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *FilterDefinitionContract) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *FilterDefinitionContract) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *FilterDefinitionContract) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### SetFilterTypeNil

`func (o *FilterDefinitionContract) SetFilterTypeNil(b bool)`

 SetFilterTypeNil sets the value for FilterType to be an explicit nil

### UnsetFilterType
`func (o *FilterDefinitionContract) UnsetFilterType()`

UnsetFilterType ensures that no value is present for FilterType, not even an explicit nil
### GetFilterName

`func (o *FilterDefinitionContract) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *FilterDefinitionContract) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *FilterDefinitionContract) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.

### HasFilterName

`func (o *FilterDefinitionContract) HasFilterName() bool`

HasFilterName returns a boolean if a field has been set.

### SetFilterNameNil

`func (o *FilterDefinitionContract) SetFilterNameNil(b bool)`

 SetFilterNameNil sets the value for FilterName to be an explicit nil

### UnsetFilterName
`func (o *FilterDefinitionContract) UnsetFilterName()`

UnsetFilterName ensures that no value is present for FilterName, not even an explicit nil
### GetExplanation

`func (o *FilterDefinitionContract) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *FilterDefinitionContract) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *FilterDefinitionContract) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *FilterDefinitionContract) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### SetExplanationNil

`func (o *FilterDefinitionContract) SetExplanationNil(b bool)`

 SetExplanationNil sets the value for Explanation to be an explicit nil

### UnsetExplanation
`func (o *FilterDefinitionContract) UnsetExplanation()`

UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil
### GetIsUserFilter

`func (o *FilterDefinitionContract) GetIsUserFilter() bool`

GetIsUserFilter returns the IsUserFilter field if non-nil, zero value otherwise.

### GetIsUserFilterOk

`func (o *FilterDefinitionContract) GetIsUserFilterOk() (*bool, bool)`

GetIsUserFilterOk returns a tuple with the IsUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserFilter

`func (o *FilterDefinitionContract) SetIsUserFilter(v bool)`

SetIsUserFilter sets IsUserFilter field to given value.

### HasIsUserFilter

`func (o *FilterDefinitionContract) HasIsUserFilter() bool`

HasIsUserFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


