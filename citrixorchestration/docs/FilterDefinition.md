# FilterDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | [**FilterType**](FilterType.md) |  | 
**FilterName** | Pointer to **string** | Localized filter name | [optional] 
**Explanation** | Pointer to **string** | Localized explanation | [optional] 
**IsUserFilter** | **bool** | True &#x3D; user filter, False &#x3D; machine filter | 
**IsSingleton** | **bool** | Is filter a singleton, only the NetScaler SD-Wan filter is singleton. | 

## Methods

### NewFilterDefinition

`func NewFilterDefinition(filterType FilterType, isUserFilter bool, isSingleton bool, ) *FilterDefinition`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


