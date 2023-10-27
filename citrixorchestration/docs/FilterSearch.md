# FilterSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFilters** | Pointer to [**[]SearchFilter4**](SearchFilter4.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**SortingMethod4**](SortingMethod4.md) |  | [optional] 

## Methods

### NewFilterSearch

`func NewFilterSearch() *FilterSearch`

NewFilterSearch instantiates a new FilterSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterSearchWithDefaults

`func NewFilterSearchWithDefaults() *FilterSearch`

NewFilterSearchWithDefaults instantiates a new FilterSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFilters

`func (o *FilterSearch) GetSearchFilters() []SearchFilter4`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *FilterSearch) GetSearchFiltersOk() (*[]SearchFilter4, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *FilterSearch) SetSearchFilters(v []SearchFilter4)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *FilterSearch) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *FilterSearch) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *FilterSearch) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *FilterSearch) GetSortCriteria() SortingMethod4`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *FilterSearch) GetSortCriteriaOk() (*SortingMethod4, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *FilterSearch) SetSortCriteria(v SortingMethod4)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *FilterSearch) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


