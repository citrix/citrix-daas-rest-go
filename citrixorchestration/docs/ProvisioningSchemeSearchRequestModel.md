# ProvisioningSchemeSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicSearchString** | Pointer to **NullableString** | Basic search string. Specify a string which will match if contained within some string property of the provisioning scheme. | [optional] 
**SearchFilters** | Pointer to [**[]ProvisioningSchemeSearchFilterRequestModel**](ProvisioningSchemeSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SearchFilterGroups** | Pointer to [**[]ProvisioningSchemeSearchFilterGroupRequestModel**](ProvisioningSchemeSearchFilterGroupRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**ProvisioningSchemeSearchFilterGroupsType**](ProvisioningSchemeSearchFilterGroupsType.md) |  | [optional] 
**SortCriteriaItems** | Pointer to [**[]ProvisioningSchemeSortCriteriaRequestModel**](ProvisioningSchemeSortCriteriaRequestModel.md) | Sort criteria for the results, multiple sorting criteria can be specified here. | [optional] 

## Methods

### NewProvisioningSchemeSearchRequestModel

`func NewProvisioningSchemeSearchRequestModel() *ProvisioningSchemeSearchRequestModel`

NewProvisioningSchemeSearchRequestModel instantiates a new ProvisioningSchemeSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeSearchRequestModelWithDefaults

`func NewProvisioningSchemeSearchRequestModelWithDefaults() *ProvisioningSchemeSearchRequestModel`

NewProvisioningSchemeSearchRequestModelWithDefaults instantiates a new ProvisioningSchemeSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicSearchString

`func (o *ProvisioningSchemeSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *ProvisioningSchemeSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *ProvisioningSchemeSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *ProvisioningSchemeSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *ProvisioningSchemeSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *ProvisioningSchemeSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchFilters

`func (o *ProvisioningSchemeSearchRequestModel) GetSearchFilters() []ProvisioningSchemeSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *ProvisioningSchemeSearchRequestModel) GetSearchFiltersOk() (*[]ProvisioningSchemeSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *ProvisioningSchemeSearchRequestModel) SetSearchFilters(v []ProvisioningSchemeSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *ProvisioningSchemeSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *ProvisioningSchemeSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *ProvisioningSchemeSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSearchFilterGroups

`func (o *ProvisioningSchemeSearchRequestModel) GetSearchFilterGroups() []ProvisioningSchemeSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *ProvisioningSchemeSearchRequestModel) GetSearchFilterGroupsOk() (*[]ProvisioningSchemeSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *ProvisioningSchemeSearchRequestModel) SetSearchFilterGroups(v []ProvisioningSchemeSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *ProvisioningSchemeSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *ProvisioningSchemeSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *ProvisioningSchemeSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *ProvisioningSchemeSearchRequestModel) GetSearchFilterGroupsType() ProvisioningSchemeSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *ProvisioningSchemeSearchRequestModel) GetSearchFilterGroupsTypeOk() (*ProvisioningSchemeSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *ProvisioningSchemeSearchRequestModel) SetSearchFilterGroupsType(v ProvisioningSchemeSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *ProvisioningSchemeSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetSortCriteriaItems

`func (o *ProvisioningSchemeSearchRequestModel) GetSortCriteriaItems() []ProvisioningSchemeSortCriteriaRequestModel`

GetSortCriteriaItems returns the SortCriteriaItems field if non-nil, zero value otherwise.

### GetSortCriteriaItemsOk

`func (o *ProvisioningSchemeSearchRequestModel) GetSortCriteriaItemsOk() (*[]ProvisioningSchemeSortCriteriaRequestModel, bool)`

GetSortCriteriaItemsOk returns a tuple with the SortCriteriaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteriaItems

`func (o *ProvisioningSchemeSearchRequestModel) SetSortCriteriaItems(v []ProvisioningSchemeSortCriteriaRequestModel)`

SetSortCriteriaItems sets SortCriteriaItems field to given value.

### HasSortCriteriaItems

`func (o *ProvisioningSchemeSearchRequestModel) HasSortCriteriaItems() bool`

HasSortCriteriaItems returns a boolean if a field has been set.

### SetSortCriteriaItemsNil

`func (o *ProvisioningSchemeSearchRequestModel) SetSortCriteriaItemsNil(b bool)`

 SetSortCriteriaItemsNil sets the value for SortCriteriaItems to be an explicit nil

### UnsetSortCriteriaItems
`func (o *ProvisioningSchemeSearchRequestModel) UnsetSortCriteriaItems()`

UnsetSortCriteriaItems ensures that no value is present for SortCriteriaItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


