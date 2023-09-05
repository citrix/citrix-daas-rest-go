# MachineAndSessionSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicSearchString** | Pointer to **string** | Basic search string. Specify a string which will match if contained within some string property of the machine. | [optional] 
**SearchFilters** | Pointer to [**[]MachineAndSessionSearchFilterRequestModel**](MachineAndSessionSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**MachineAndSessionSearchRequestModelSortCriteria**](MachineAndSessionSearchRequestModelSortCriteria.md) |  | [optional] 
**SearchFilterGroups** | Pointer to [**[]MachineAndSessionSearchFilterGroupRequestModel**](MachineAndSessionSearchFilterGroupRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**MachineAndSessionSearchFilterGroupsType**](MachineAndSessionSearchFilterGroupsType.md) |  | [optional] 
**SortCriteriaItems** | Pointer to [**[]MachineAndSessionSortCriteriaRequestModel**](MachineAndSessionSortCriteriaRequestModel.md) | Sort criteria for the results, multiple sorting criteria can be specified here. | [optional] 

## Methods

### NewMachineAndSessionSearchRequestModel

`func NewMachineAndSessionSearchRequestModel() *MachineAndSessionSearchRequestModel`

NewMachineAndSessionSearchRequestModel instantiates a new MachineAndSessionSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAndSessionSearchRequestModelWithDefaults

`func NewMachineAndSessionSearchRequestModelWithDefaults() *MachineAndSessionSearchRequestModel`

NewMachineAndSessionSearchRequestModelWithDefaults instantiates a new MachineAndSessionSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicSearchString

`func (o *MachineAndSessionSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *MachineAndSessionSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *MachineAndSessionSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *MachineAndSessionSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### GetSearchFilters

`func (o *MachineAndSessionSearchRequestModel) GetSearchFilters() []MachineAndSessionSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *MachineAndSessionSearchRequestModel) GetSearchFiltersOk() (*[]MachineAndSessionSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *MachineAndSessionSearchRequestModel) SetSearchFilters(v []MachineAndSessionSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *MachineAndSessionSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### GetSortCriteria

`func (o *MachineAndSessionSearchRequestModel) GetSortCriteria() MachineAndSessionSearchRequestModelSortCriteria`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *MachineAndSessionSearchRequestModel) GetSortCriteriaOk() (*MachineAndSessionSearchRequestModelSortCriteria, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *MachineAndSessionSearchRequestModel) SetSortCriteria(v MachineAndSessionSearchRequestModelSortCriteria)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *MachineAndSessionSearchRequestModel) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.

### GetSearchFilterGroups

`func (o *MachineAndSessionSearchRequestModel) GetSearchFilterGroups() []MachineAndSessionSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *MachineAndSessionSearchRequestModel) GetSearchFilterGroupsOk() (*[]MachineAndSessionSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *MachineAndSessionSearchRequestModel) SetSearchFilterGroups(v []MachineAndSessionSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *MachineAndSessionSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### GetSearchFilterGroupsType

`func (o *MachineAndSessionSearchRequestModel) GetSearchFilterGroupsType() MachineAndSessionSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *MachineAndSessionSearchRequestModel) GetSearchFilterGroupsTypeOk() (*MachineAndSessionSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *MachineAndSessionSearchRequestModel) SetSearchFilterGroupsType(v MachineAndSessionSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *MachineAndSessionSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetSortCriteriaItems

`func (o *MachineAndSessionSearchRequestModel) GetSortCriteriaItems() []MachineAndSessionSortCriteriaRequestModel`

GetSortCriteriaItems returns the SortCriteriaItems field if non-nil, zero value otherwise.

### GetSortCriteriaItemsOk

`func (o *MachineAndSessionSearchRequestModel) GetSortCriteriaItemsOk() (*[]MachineAndSessionSortCriteriaRequestModel, bool)`

GetSortCriteriaItemsOk returns a tuple with the SortCriteriaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteriaItems

`func (o *MachineAndSessionSearchRequestModel) SetSortCriteriaItems(v []MachineAndSessionSortCriteriaRequestModel)`

SetSortCriteriaItems sets SortCriteriaItems field to given value.

### HasSortCriteriaItems

`func (o *MachineAndSessionSearchRequestModel) HasSortCriteriaItems() bool`

HasSortCriteriaItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


