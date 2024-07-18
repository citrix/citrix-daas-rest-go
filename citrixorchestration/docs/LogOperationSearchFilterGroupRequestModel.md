# LogOperationSearchFilterGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchType** | Pointer to [**LogOperationSearchFilterGroupType**](LogOperationSearchFilterGroupType.md) |  | [optional] 
**SearchFilters** | Pointer to [**[]LogOperationSearchFilterRequestModel**](LogOperationSearchFilterRequestModel.md) | Advanced search filters. | [optional] 
**SearchFilterGroups** | Pointer to [**[]LogOperationSearchFilterGroupRequestModel**](LogOperationSearchFilterGroupRequestModel.md) | Advanced search filter groups. | [optional] 

## Methods

### NewLogOperationSearchFilterGroupRequestModel

`func NewLogOperationSearchFilterGroupRequestModel() *LogOperationSearchFilterGroupRequestModel`

NewLogOperationSearchFilterGroupRequestModel instantiates a new LogOperationSearchFilterGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogOperationSearchFilterGroupRequestModelWithDefaults

`func NewLogOperationSearchFilterGroupRequestModelWithDefaults() *LogOperationSearchFilterGroupRequestModel`

NewLogOperationSearchFilterGroupRequestModelWithDefaults instantiates a new LogOperationSearchFilterGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchType

`func (o *LogOperationSearchFilterGroupRequestModel) GetSearchType() LogOperationSearchFilterGroupType`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *LogOperationSearchFilterGroupRequestModel) GetSearchTypeOk() (*LogOperationSearchFilterGroupType, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *LogOperationSearchFilterGroupRequestModel) SetSearchType(v LogOperationSearchFilterGroupType)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *LogOperationSearchFilterGroupRequestModel) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

### GetSearchFilters

`func (o *LogOperationSearchFilterGroupRequestModel) GetSearchFilters() []LogOperationSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *LogOperationSearchFilterGroupRequestModel) GetSearchFiltersOk() (*[]LogOperationSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *LogOperationSearchFilterGroupRequestModel) SetSearchFilters(v []LogOperationSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *LogOperationSearchFilterGroupRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *LogOperationSearchFilterGroupRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *LogOperationSearchFilterGroupRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSearchFilterGroups

`func (o *LogOperationSearchFilterGroupRequestModel) GetSearchFilterGroups() []LogOperationSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *LogOperationSearchFilterGroupRequestModel) GetSearchFilterGroupsOk() (*[]LogOperationSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *LogOperationSearchFilterGroupRequestModel) SetSearchFilterGroups(v []LogOperationSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *LogOperationSearchFilterGroupRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *LogOperationSearchFilterGroupRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *LogOperationSearchFilterGroupRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


