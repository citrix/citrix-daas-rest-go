# MaintenanceCycleVMOperationJobSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicSearchString** | Pointer to **NullableString** | Basic search string. Specify a string which will match if contained within some string property of the catalog. | [optional] 
**SearchFilters** | Pointer to [**[]MaintenanceCycleVMOperationJobSearchFilterRequestModel**](MaintenanceCycleVMOperationJobSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**MaintenanceCycleVMOperationJobSortCriteriaRequestModel**](MaintenanceCycleVMOperationJobSortCriteriaRequestModel.md) |  | [optional] 
**SearchFilterGroups** | Pointer to [**[]MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel**](MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**MaintenanceCycleVmOperationJobSearchFilterGroupsType**](MaintenanceCycleVmOperationJobSearchFilterGroupsType.md) |  | [optional] 
**SortCriteriaItems** | Pointer to [**[]MaintenanceCycleVMOperationJobSortCriteriaRequestModel**](MaintenanceCycleVMOperationJobSortCriteriaRequestModel.md) | Sort criteria for the results, multiple sorting criteria can be specified here. | [optional] 

## Methods

### NewMaintenanceCycleVMOperationJobSearchRequestModel

`func NewMaintenanceCycleVMOperationJobSearchRequestModel() *MaintenanceCycleVMOperationJobSearchRequestModel`

NewMaintenanceCycleVMOperationJobSearchRequestModel instantiates a new MaintenanceCycleVMOperationJobSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceCycleVMOperationJobSearchRequestModelWithDefaults

`func NewMaintenanceCycleVMOperationJobSearchRequestModelWithDefaults() *MaintenanceCycleVMOperationJobSearchRequestModel`

NewMaintenanceCycleVMOperationJobSearchRequestModelWithDefaults instantiates a new MaintenanceCycleVMOperationJobSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicSearchString

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchFilters

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSearchFilters() []MaintenanceCycleVMOperationJobSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSearchFiltersOk() (*[]MaintenanceCycleVMOperationJobSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetSearchFilters(v []MaintenanceCycleVMOperationJobSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSortCriteria() MaintenanceCycleVMOperationJobSortCriteriaRequestModel`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSortCriteriaOk() (*MaintenanceCycleVMOperationJobSortCriteriaRequestModel, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetSortCriteria(v MaintenanceCycleVMOperationJobSortCriteriaRequestModel)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.

### GetSearchFilterGroups

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSearchFilterGroups() []MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSearchFilterGroupsOk() (*[]MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetSearchFilterGroups(v []MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSearchFilterGroupsType() MaintenanceCycleVmOperationJobSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSearchFilterGroupsTypeOk() (*MaintenanceCycleVmOperationJobSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetSearchFilterGroupsType(v MaintenanceCycleVmOperationJobSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetSortCriteriaItems

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSortCriteriaItems() []MaintenanceCycleVMOperationJobSortCriteriaRequestModel`

GetSortCriteriaItems returns the SortCriteriaItems field if non-nil, zero value otherwise.

### GetSortCriteriaItemsOk

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) GetSortCriteriaItemsOk() (*[]MaintenanceCycleVMOperationJobSortCriteriaRequestModel, bool)`

GetSortCriteriaItemsOk returns a tuple with the SortCriteriaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteriaItems

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetSortCriteriaItems(v []MaintenanceCycleVMOperationJobSortCriteriaRequestModel)`

SetSortCriteriaItems sets SortCriteriaItems field to given value.

### HasSortCriteriaItems

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) HasSortCriteriaItems() bool`

HasSortCriteriaItems returns a boolean if a field has been set.

### SetSortCriteriaItemsNil

`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) SetSortCriteriaItemsNil(b bool)`

 SetSortCriteriaItemsNil sets the value for SortCriteriaItems to be an explicit nil

### UnsetSortCriteriaItems
`func (o *MaintenanceCycleVMOperationJobSearchRequestModel) UnsetSortCriteriaItems()`

UnsetSortCriteriaItems ensures that no value is present for SortCriteriaItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


