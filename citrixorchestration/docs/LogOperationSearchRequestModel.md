# LogOperationSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFilters** | Pointer to [**[]LogOperationSearchFilterRequestModel**](LogOperationSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**LogOperationSearchRequestModelSortCriteria**](LogOperationSearchRequestModelSortCriteria.md) |  | [optional] 

## Methods

### NewLogOperationSearchRequestModel

`func NewLogOperationSearchRequestModel() *LogOperationSearchRequestModel`

NewLogOperationSearchRequestModel instantiates a new LogOperationSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogOperationSearchRequestModelWithDefaults

`func NewLogOperationSearchRequestModelWithDefaults() *LogOperationSearchRequestModel`

NewLogOperationSearchRequestModelWithDefaults instantiates a new LogOperationSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFilters

`func (o *LogOperationSearchRequestModel) GetSearchFilters() []LogOperationSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *LogOperationSearchRequestModel) GetSearchFiltersOk() (*[]LogOperationSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *LogOperationSearchRequestModel) SetSearchFilters(v []LogOperationSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *LogOperationSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### GetSortCriteria

`func (o *LogOperationSearchRequestModel) GetSortCriteria() LogOperationSearchRequestModelSortCriteria`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *LogOperationSearchRequestModel) GetSortCriteriaOk() (*LogOperationSearchRequestModelSortCriteria, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *LogOperationSearchRequestModel) SetSortCriteria(v LogOperationSearchRequestModelSortCriteria)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *LogOperationSearchRequestModel) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


