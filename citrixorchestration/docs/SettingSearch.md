# SettingSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFilters** | Pointer to [**[]SearchFilter3**](SearchFilter3.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**SortingMethod3**](SortingMethod3.md) |  | [optional] 

## Methods

### NewSettingSearch

`func NewSettingSearch() *SettingSearch`

NewSettingSearch instantiates a new SettingSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingSearchWithDefaults

`func NewSettingSearchWithDefaults() *SettingSearch`

NewSettingSearchWithDefaults instantiates a new SettingSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFilters

`func (o *SettingSearch) GetSearchFilters() []SearchFilter3`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *SettingSearch) GetSearchFiltersOk() (*[]SearchFilter3, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *SettingSearch) SetSearchFilters(v []SearchFilter3)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *SettingSearch) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *SettingSearch) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *SettingSearch) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *SettingSearch) GetSortCriteria() SortingMethod3`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *SettingSearch) GetSortCriteriaOk() (*SortingMethod3, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *SettingSearch) SetSortCriteria(v SortingMethod3)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *SettingSearch) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


