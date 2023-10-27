# PolicySetSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFilters** | Pointer to [**[]SearchFilter**](SearchFilter.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**SortingMethod**](SortingMethod.md) |  | [optional] 

## Methods

### NewPolicySetSearch

`func NewPolicySetSearch() *PolicySetSearch`

NewPolicySetSearch instantiates a new PolicySetSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySetSearchWithDefaults

`func NewPolicySetSearchWithDefaults() *PolicySetSearch`

NewPolicySetSearchWithDefaults instantiates a new PolicySetSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFilters

`func (o *PolicySetSearch) GetSearchFilters() []SearchFilter`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *PolicySetSearch) GetSearchFiltersOk() (*[]SearchFilter, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *PolicySetSearch) SetSearchFilters(v []SearchFilter)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *PolicySetSearch) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *PolicySetSearch) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *PolicySetSearch) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *PolicySetSearch) GetSortCriteria() SortingMethod`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *PolicySetSearch) GetSortCriteriaOk() (*SortingMethod, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *PolicySetSearch) SetSortCriteria(v SortingMethod)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *PolicySetSearch) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


