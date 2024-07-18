# ProvisioningOperationEventSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicSearchString** | Pointer to **NullableString** | Basic search string. Specify a string which will match if contained within some string property of the operation event. | [optional] 
**SearchFilters** | Pointer to [**[]ProvisioningOperationEventSearchFilterRequestModel**](ProvisioningOperationEventSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SearchFilterGroups** | Pointer to [**[]ProvisioningOperationEventSearchFilterGroupRequestModel**](ProvisioningOperationEventSearchFilterGroupRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**ProvisioningOperationEventSearchFilterGroupsType**](ProvisioningOperationEventSearchFilterGroupsType.md) |  | [optional] 
**SortCriteriaItems** | Pointer to [**[]ProvisioningOperationEventSortCriteriaRequestModel**](ProvisioningOperationEventSortCriteriaRequestModel.md) | Sort criteria for the results, multiple sorting criteria can be specified here. | [optional] 

## Methods

### NewProvisioningOperationEventSearchRequestModel

`func NewProvisioningOperationEventSearchRequestModel() *ProvisioningOperationEventSearchRequestModel`

NewProvisioningOperationEventSearchRequestModel instantiates a new ProvisioningOperationEventSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningOperationEventSearchRequestModelWithDefaults

`func NewProvisioningOperationEventSearchRequestModelWithDefaults() *ProvisioningOperationEventSearchRequestModel`

NewProvisioningOperationEventSearchRequestModelWithDefaults instantiates a new ProvisioningOperationEventSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicSearchString

`func (o *ProvisioningOperationEventSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *ProvisioningOperationEventSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *ProvisioningOperationEventSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *ProvisioningOperationEventSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *ProvisioningOperationEventSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *ProvisioningOperationEventSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchFilters

`func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilters() []ProvisioningOperationEventSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFiltersOk() (*[]ProvisioningOperationEventSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *ProvisioningOperationEventSearchRequestModel) SetSearchFilters(v []ProvisioningOperationEventSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *ProvisioningOperationEventSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *ProvisioningOperationEventSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *ProvisioningOperationEventSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSearchFilterGroups

`func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilterGroups() []ProvisioningOperationEventSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilterGroupsOk() (*[]ProvisioningOperationEventSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *ProvisioningOperationEventSearchRequestModel) SetSearchFilterGroups(v []ProvisioningOperationEventSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *ProvisioningOperationEventSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *ProvisioningOperationEventSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *ProvisioningOperationEventSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilterGroupsType() ProvisioningOperationEventSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *ProvisioningOperationEventSearchRequestModel) GetSearchFilterGroupsTypeOk() (*ProvisioningOperationEventSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *ProvisioningOperationEventSearchRequestModel) SetSearchFilterGroupsType(v ProvisioningOperationEventSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *ProvisioningOperationEventSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetSortCriteriaItems

`func (o *ProvisioningOperationEventSearchRequestModel) GetSortCriteriaItems() []ProvisioningOperationEventSortCriteriaRequestModel`

GetSortCriteriaItems returns the SortCriteriaItems field if non-nil, zero value otherwise.

### GetSortCriteriaItemsOk

`func (o *ProvisioningOperationEventSearchRequestModel) GetSortCriteriaItemsOk() (*[]ProvisioningOperationEventSortCriteriaRequestModel, bool)`

GetSortCriteriaItemsOk returns a tuple with the SortCriteriaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteriaItems

`func (o *ProvisioningOperationEventSearchRequestModel) SetSortCriteriaItems(v []ProvisioningOperationEventSortCriteriaRequestModel)`

SetSortCriteriaItems sets SortCriteriaItems field to given value.

### HasSortCriteriaItems

`func (o *ProvisioningOperationEventSearchRequestModel) HasSortCriteriaItems() bool`

HasSortCriteriaItems returns a boolean if a field has been set.

### SetSortCriteriaItemsNil

`func (o *ProvisioningOperationEventSearchRequestModel) SetSortCriteriaItemsNil(b bool)`

 SetSortCriteriaItemsNil sets the value for SortCriteriaItems to be an explicit nil

### UnsetSortCriteriaItems
`func (o *ProvisioningOperationEventSearchRequestModel) UnsetSortCriteriaItems()`

UnsetSortCriteriaItems ensures that no value is present for SortCriteriaItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


