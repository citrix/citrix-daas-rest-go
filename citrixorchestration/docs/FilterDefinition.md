# FilterDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to [**FilterType**](FilterType.md) |  | [optional] 
**FilterName** | Pointer to **NullableString** | Localized filter name | [optional] 
**Explanation** | Pointer to **NullableString** | Localized explanation | [optional] 
**IsUserFilter** | Pointer to **bool** | True &#x3D; user filter, False &#x3D; machine filter | [optional] 
**IsSingleton** | Pointer to **bool** | Is filter a singleton, only the NetScaler SD-Wan filter is singleton. | [optional] 

## Methods

### NewFilterDefinition

`func NewFilterDefinition() *FilterDefinition`

NewFilterDefinition instantiates a new FilterDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterDefinitionWithDefaults

`func NewFilterDefinitionWithDefaults() *FilterDefinition`

NewFilterDefinitionWithDefaults instantiates a new FilterDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *FilterDefinition) GetFilterType() FilterType`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *FilterDefinition) GetFilterTypeOk() (*FilterType, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *FilterDefinition) SetFilterType(v FilterType)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *FilterDefinition) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetFilterName

`func (o *FilterDefinition) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *FilterDefinition) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *FilterDefinition) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.

### HasFilterName

`func (o *FilterDefinition) HasFilterName() bool`

HasFilterName returns a boolean if a field has been set.

### SetFilterNameNil

`func (o *FilterDefinition) SetFilterNameNil(b bool)`

 SetFilterNameNil sets the value for FilterName to be an explicit nil

### UnsetFilterName
`func (o *FilterDefinition) UnsetFilterName()`

UnsetFilterName ensures that no value is present for FilterName, not even an explicit nil
### GetExplanation

`func (o *FilterDefinition) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *FilterDefinition) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *FilterDefinition) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *FilterDefinition) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### SetExplanationNil

`func (o *FilterDefinition) SetExplanationNil(b bool)`

 SetExplanationNil sets the value for Explanation to be an explicit nil

### UnsetExplanation
`func (o *FilterDefinition) UnsetExplanation()`

UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil
### GetIsUserFilter

`func (o *FilterDefinition) GetIsUserFilter() bool`

GetIsUserFilter returns the IsUserFilter field if non-nil, zero value otherwise.

### GetIsUserFilterOk

`func (o *FilterDefinition) GetIsUserFilterOk() (*bool, bool)`

GetIsUserFilterOk returns a tuple with the IsUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserFilter

`func (o *FilterDefinition) SetIsUserFilter(v bool)`

SetIsUserFilter sets IsUserFilter field to given value.

### HasIsUserFilter

`func (o *FilterDefinition) HasIsUserFilter() bool`

HasIsUserFilter returns a boolean if a field has been set.

### GetIsSingleton

`func (o *FilterDefinition) GetIsSingleton() bool`

GetIsSingleton returns the IsSingleton field if non-nil, zero value otherwise.

### GetIsSingletonOk

`func (o *FilterDefinition) GetIsSingletonOk() (*bool, bool)`

GetIsSingletonOk returns a tuple with the IsSingleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleton

`func (o *FilterDefinition) SetIsSingleton(v bool)`

SetIsSingleton sets IsSingleton field to given value.

### HasIsSingleton

`func (o *FilterDefinition) HasIsSingleton() bool`

HasIsSingleton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


