# ProvisionedVirtualMachineSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicSearchString** | Pointer to **NullableString** | Basic search string. Specify a string which will match if contained within some string property of the machine. | [optional] 
**SearchFilters** | Pointer to [**[]ProvisionedVirtualMachineSearchFilterRequestModel**](ProvisionedVirtualMachineSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SearchFilterGroups** | Pointer to [**[]ProvisionedVirtualMachineSearchFilterGroupRequestModel**](ProvisionedVirtualMachineSearchFilterGroupRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**ProvisionedVirtualMachineSearchFilterGroupsType**](ProvisionedVirtualMachineSearchFilterGroupsType.md) |  | [optional] 
**SortCriteriaItems** | Pointer to [**[]ProvisionedVirtualMachineSortCriteriaRequestModel**](ProvisionedVirtualMachineSortCriteriaRequestModel.md) | Sort criteria for the results, multiple sorting criteria can be specified here. | [optional] 

## Methods

### NewProvisionedVirtualMachineSearchRequestModel

`func NewProvisionedVirtualMachineSearchRequestModel() *ProvisionedVirtualMachineSearchRequestModel`

NewProvisionedVirtualMachineSearchRequestModel instantiates a new ProvisionedVirtualMachineSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedVirtualMachineSearchRequestModelWithDefaults

`func NewProvisionedVirtualMachineSearchRequestModelWithDefaults() *ProvisionedVirtualMachineSearchRequestModel`

NewProvisionedVirtualMachineSearchRequestModelWithDefaults instantiates a new ProvisionedVirtualMachineSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicSearchString

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *ProvisionedVirtualMachineSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *ProvisionedVirtualMachineSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *ProvisionedVirtualMachineSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *ProvisionedVirtualMachineSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchFilters

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetSearchFilters() []ProvisionedVirtualMachineSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetSearchFiltersOk() (*[]ProvisionedVirtualMachineSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *ProvisionedVirtualMachineSearchRequestModel) SetSearchFilters(v []ProvisionedVirtualMachineSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *ProvisionedVirtualMachineSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *ProvisionedVirtualMachineSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *ProvisionedVirtualMachineSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSearchFilterGroups

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetSearchFilterGroups() []ProvisionedVirtualMachineSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetSearchFilterGroupsOk() (*[]ProvisionedVirtualMachineSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *ProvisionedVirtualMachineSearchRequestModel) SetSearchFilterGroups(v []ProvisionedVirtualMachineSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *ProvisionedVirtualMachineSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *ProvisionedVirtualMachineSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *ProvisionedVirtualMachineSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetSearchFilterGroupsType() ProvisionedVirtualMachineSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetSearchFilterGroupsTypeOk() (*ProvisionedVirtualMachineSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *ProvisionedVirtualMachineSearchRequestModel) SetSearchFilterGroupsType(v ProvisionedVirtualMachineSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *ProvisionedVirtualMachineSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetSortCriteriaItems

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetSortCriteriaItems() []ProvisionedVirtualMachineSortCriteriaRequestModel`

GetSortCriteriaItems returns the SortCriteriaItems field if non-nil, zero value otherwise.

### GetSortCriteriaItemsOk

`func (o *ProvisionedVirtualMachineSearchRequestModel) GetSortCriteriaItemsOk() (*[]ProvisionedVirtualMachineSortCriteriaRequestModel, bool)`

GetSortCriteriaItemsOk returns a tuple with the SortCriteriaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteriaItems

`func (o *ProvisionedVirtualMachineSearchRequestModel) SetSortCriteriaItems(v []ProvisionedVirtualMachineSortCriteriaRequestModel)`

SetSortCriteriaItems sets SortCriteriaItems field to given value.

### HasSortCriteriaItems

`func (o *ProvisionedVirtualMachineSearchRequestModel) HasSortCriteriaItems() bool`

HasSortCriteriaItems returns a boolean if a field has been set.

### SetSortCriteriaItemsNil

`func (o *ProvisionedVirtualMachineSearchRequestModel) SetSortCriteriaItemsNil(b bool)`

 SetSortCriteriaItemsNil sets the value for SortCriteriaItems to be an explicit nil

### UnsetSortCriteriaItems
`func (o *ProvisionedVirtualMachineSearchRequestModel) UnsetSortCriteriaItems()`

UnsetSortCriteriaItems ensures that no value is present for SortCriteriaItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


