# PolicySearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFilters** | Pointer to [**[]SearchFilter2**](SearchFilter2.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**SortingMethod2**](SortingMethod2.md) |  | [optional] 

## Methods

### NewPolicySearch

`func NewPolicySearch() *PolicySearch`

NewPolicySearch instantiates a new PolicySearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySearchWithDefaults

`func NewPolicySearchWithDefaults() *PolicySearch`

NewPolicySearchWithDefaults instantiates a new PolicySearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFilters

`func (o *PolicySearch) GetSearchFilters() []SearchFilter2`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *PolicySearch) GetSearchFiltersOk() (*[]SearchFilter2, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *PolicySearch) SetSearchFilters(v []SearchFilter2)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *PolicySearch) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *PolicySearch) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *PolicySearch) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *PolicySearch) GetSortCriteria() SortingMethod2`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *PolicySearch) GetSortCriteriaOk() (*SortingMethod2, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *PolicySearch) SetSortCriteria(v SortingMethod2)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *PolicySearch) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


