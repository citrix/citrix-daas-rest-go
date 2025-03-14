# MaintenanceCycleSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicSearchString** | Pointer to **NullableString** | Basic search string. Specify a string which will match if contained within some string property of the catalog. | [optional] 
**SearchFilters** | Pointer to [**[]MaintenanceCycleSearchFilterRequestModel**](MaintenanceCycleSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**MaintenanceCycleSortCriteriaRequestModel**](MaintenanceCycleSortCriteriaRequestModel.md) |  | [optional] 
**SearchFilterGroups** | Pointer to [**[]MaintenanceCycleSearchFilterGroupsRequestModel**](MaintenanceCycleSearchFilterGroupsRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**MaintenanceCycleSearchFilterGroupsType**](MaintenanceCycleSearchFilterGroupsType.md) |  | [optional] 
**SortCriteriaItems** | Pointer to [**[]MaintenanceCycleSortCriteriaRequestModel**](MaintenanceCycleSortCriteriaRequestModel.md) | Sort criteria for the results, multiple sorting criteria can be specified here. | [optional] 

## Methods

### NewMaintenanceCycleSearchRequestModel

`func NewMaintenanceCycleSearchRequestModel() *MaintenanceCycleSearchRequestModel`

NewMaintenanceCycleSearchRequestModel instantiates a new MaintenanceCycleSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceCycleSearchRequestModelWithDefaults

`func NewMaintenanceCycleSearchRequestModelWithDefaults() *MaintenanceCycleSearchRequestModel`

NewMaintenanceCycleSearchRequestModelWithDefaults instantiates a new MaintenanceCycleSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicSearchString

`func (o *MaintenanceCycleSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *MaintenanceCycleSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *MaintenanceCycleSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *MaintenanceCycleSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *MaintenanceCycleSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *MaintenanceCycleSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchFilters

`func (o *MaintenanceCycleSearchRequestModel) GetSearchFilters() []MaintenanceCycleSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *MaintenanceCycleSearchRequestModel) GetSearchFiltersOk() (*[]MaintenanceCycleSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *MaintenanceCycleSearchRequestModel) SetSearchFilters(v []MaintenanceCycleSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *MaintenanceCycleSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *MaintenanceCycleSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *MaintenanceCycleSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *MaintenanceCycleSearchRequestModel) GetSortCriteria() MaintenanceCycleSortCriteriaRequestModel`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *MaintenanceCycleSearchRequestModel) GetSortCriteriaOk() (*MaintenanceCycleSortCriteriaRequestModel, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *MaintenanceCycleSearchRequestModel) SetSortCriteria(v MaintenanceCycleSortCriteriaRequestModel)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *MaintenanceCycleSearchRequestModel) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.

### GetSearchFilterGroups

`func (o *MaintenanceCycleSearchRequestModel) GetSearchFilterGroups() []MaintenanceCycleSearchFilterGroupsRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *MaintenanceCycleSearchRequestModel) GetSearchFilterGroupsOk() (*[]MaintenanceCycleSearchFilterGroupsRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *MaintenanceCycleSearchRequestModel) SetSearchFilterGroups(v []MaintenanceCycleSearchFilterGroupsRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *MaintenanceCycleSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *MaintenanceCycleSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *MaintenanceCycleSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *MaintenanceCycleSearchRequestModel) GetSearchFilterGroupsType() MaintenanceCycleSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *MaintenanceCycleSearchRequestModel) GetSearchFilterGroupsTypeOk() (*MaintenanceCycleSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *MaintenanceCycleSearchRequestModel) SetSearchFilterGroupsType(v MaintenanceCycleSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *MaintenanceCycleSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetSortCriteriaItems

`func (o *MaintenanceCycleSearchRequestModel) GetSortCriteriaItems() []MaintenanceCycleSortCriteriaRequestModel`

GetSortCriteriaItems returns the SortCriteriaItems field if non-nil, zero value otherwise.

### GetSortCriteriaItemsOk

`func (o *MaintenanceCycleSearchRequestModel) GetSortCriteriaItemsOk() (*[]MaintenanceCycleSortCriteriaRequestModel, bool)`

GetSortCriteriaItemsOk returns a tuple with the SortCriteriaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteriaItems

`func (o *MaintenanceCycleSearchRequestModel) SetSortCriteriaItems(v []MaintenanceCycleSortCriteriaRequestModel)`

SetSortCriteriaItems sets SortCriteriaItems field to given value.

### HasSortCriteriaItems

`func (o *MaintenanceCycleSearchRequestModel) HasSortCriteriaItems() bool`

HasSortCriteriaItems returns a boolean if a field has been set.

### SetSortCriteriaItemsNil

`func (o *MaintenanceCycleSearchRequestModel) SetSortCriteriaItemsNil(b bool)`

 SetSortCriteriaItemsNil sets the value for SortCriteriaItems to be an explicit nil

### UnsetSortCriteriaItems
`func (o *MaintenanceCycleSearchRequestModel) UnsetSortCriteriaItems()`

UnsetSortCriteriaItems ensures that no value is present for SortCriteriaItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


