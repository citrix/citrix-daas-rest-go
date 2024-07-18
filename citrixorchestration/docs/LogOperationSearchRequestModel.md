# LogOperationSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicSearchString** | Pointer to **NullableString** | Basic search string. | [optional] 
**SearchType** | Pointer to [**LogOperationSearchFilterGroupType**](LogOperationSearchFilterGroupType.md) |  | [optional] 
**SearchFilters** | Pointer to [**[]LogOperationSearchFilterRequestModel**](LogOperationSearchFilterRequestModel.md) | Advanced search filters. | [optional] 
**SearchFilterGroups** | Pointer to [**[]LogOperationSearchFilterGroupRequestModel**](LogOperationSearchFilterGroupRequestModel.md) | Advanced search filter groups. | [optional] 
**SortCriteria** | Pointer to [**LogOperationSortCriteriaRequestModel**](LogOperationSortCriteriaRequestModel.md) |  | [optional] 

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

### GetBasicSearchString

`func (o *LogOperationSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *LogOperationSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *LogOperationSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *LogOperationSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *LogOperationSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *LogOperationSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchType

`func (o *LogOperationSearchRequestModel) GetSearchType() LogOperationSearchFilterGroupType`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *LogOperationSearchRequestModel) GetSearchTypeOk() (*LogOperationSearchFilterGroupType, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *LogOperationSearchRequestModel) SetSearchType(v LogOperationSearchFilterGroupType)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *LogOperationSearchRequestModel) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

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

### SetSearchFiltersNil

`func (o *LogOperationSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *LogOperationSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSearchFilterGroups

`func (o *LogOperationSearchRequestModel) GetSearchFilterGroups() []LogOperationSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *LogOperationSearchRequestModel) GetSearchFilterGroupsOk() (*[]LogOperationSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *LogOperationSearchRequestModel) SetSearchFilterGroups(v []LogOperationSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *LogOperationSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *LogOperationSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *LogOperationSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSortCriteria

`func (o *LogOperationSearchRequestModel) GetSortCriteria() LogOperationSortCriteriaRequestModel`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *LogOperationSearchRequestModel) GetSortCriteriaOk() (*LogOperationSortCriteriaRequestModel, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *LogOperationSearchRequestModel) SetSortCriteria(v LogOperationSortCriteriaRequestModel)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *LogOperationSearchRequestModel) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


