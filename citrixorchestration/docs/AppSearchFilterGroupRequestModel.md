# AppSearchFilterGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFilterGroupType** | Pointer to [**AppSearchFilterGroupType**](AppSearchFilterGroupType.md) |  | [optional] 
**SearchFilters** | Pointer to [**[]AppSearchFilterRequestModel**](AppSearchFilterRequestModel.md) | The search filters in search filter group | [optional] 
**SearchFilterGroupsType** | Pointer to [**AppSearchFilterGroupsType**](AppSearchFilterGroupsType.md) |  | [optional] 
**SearchFilterGroups** | Pointer to [**[]AppSearchFilterGroupRequestModel**](AppSearchFilterGroupRequestModel.md) | The search filter group in search filter groups | [optional] 

## Methods

### NewAppSearchFilterGroupRequestModel

`func NewAppSearchFilterGroupRequestModel() *AppSearchFilterGroupRequestModel`

NewAppSearchFilterGroupRequestModel instantiates a new AppSearchFilterGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSearchFilterGroupRequestModelWithDefaults

`func NewAppSearchFilterGroupRequestModelWithDefaults() *AppSearchFilterGroupRequestModel`

NewAppSearchFilterGroupRequestModelWithDefaults instantiates a new AppSearchFilterGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFilterGroupType

`func (o *AppSearchFilterGroupRequestModel) GetSearchFilterGroupType() AppSearchFilterGroupType`

GetSearchFilterGroupType returns the SearchFilterGroupType field if non-nil, zero value otherwise.

### GetSearchFilterGroupTypeOk

`func (o *AppSearchFilterGroupRequestModel) GetSearchFilterGroupTypeOk() (*AppSearchFilterGroupType, bool)`

GetSearchFilterGroupTypeOk returns a tuple with the SearchFilterGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupType

`func (o *AppSearchFilterGroupRequestModel) SetSearchFilterGroupType(v AppSearchFilterGroupType)`

SetSearchFilterGroupType sets SearchFilterGroupType field to given value.

### HasSearchFilterGroupType

`func (o *AppSearchFilterGroupRequestModel) HasSearchFilterGroupType() bool`

HasSearchFilterGroupType returns a boolean if a field has been set.

### GetSearchFilters

`func (o *AppSearchFilterGroupRequestModel) GetSearchFilters() []AppSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *AppSearchFilterGroupRequestModel) GetSearchFiltersOk() (*[]AppSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *AppSearchFilterGroupRequestModel) SetSearchFilters(v []AppSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *AppSearchFilterGroupRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *AppSearchFilterGroupRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *AppSearchFilterGroupRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *AppSearchFilterGroupRequestModel) GetSearchFilterGroupsType() AppSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *AppSearchFilterGroupRequestModel) GetSearchFilterGroupsTypeOk() (*AppSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *AppSearchFilterGroupRequestModel) SetSearchFilterGroupsType(v AppSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *AppSearchFilterGroupRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetSearchFilterGroups

`func (o *AppSearchFilterGroupRequestModel) GetSearchFilterGroups() []AppSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *AppSearchFilterGroupRequestModel) GetSearchFilterGroupsOk() (*[]AppSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *AppSearchFilterGroupRequestModel) SetSearchFilterGroups(v []AppSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *AppSearchFilterGroupRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *AppSearchFilterGroupRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *AppSearchFilterGroupRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


