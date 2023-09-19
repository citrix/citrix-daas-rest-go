# CatalogSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminFolder** | Pointer to **NullableString** | Admin folder in which to search for the machine catalog. | [optional] 
**IncludeSubFolders** | Pointer to **NullableBool** | Whether to include subfolders of the AdminFolder in the search for the machine catalogs. | [optional] [default to true]
**BasicSearchString** | Pointer to **NullableString** | Basic search string. Specify a string which will match if contained within some string property of the catalog. | [optional] 
**SearchFilters** | Pointer to [**[]CatalogSearchFilterRequestModel**](CatalogSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**CatalogSortCriteriaRequestModel**](CatalogSortCriteriaRequestModel.md) |  | [optional] 
**SearchFilterGroups** | Pointer to [**[]CatalogSearchFilterGroupRequestModel**](CatalogSearchFilterGroupRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**CatalogSearchFilterGroupsType**](CatalogSearchFilterGroupsType.md) |  | [optional] 

## Methods

### NewCatalogSearchRequestModel

`func NewCatalogSearchRequestModel() *CatalogSearchRequestModel`

NewCatalogSearchRequestModel instantiates a new CatalogSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogSearchRequestModelWithDefaults

`func NewCatalogSearchRequestModelWithDefaults() *CatalogSearchRequestModel`

NewCatalogSearchRequestModelWithDefaults instantiates a new CatalogSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminFolder

`func (o *CatalogSearchRequestModel) GetAdminFolder() string`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *CatalogSearchRequestModel) GetAdminFolderOk() (*string, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *CatalogSearchRequestModel) SetAdminFolder(v string)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *CatalogSearchRequestModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### SetAdminFolderNil

`func (o *CatalogSearchRequestModel) SetAdminFolderNil(b bool)`

 SetAdminFolderNil sets the value for AdminFolder to be an explicit nil

### UnsetAdminFolder
`func (o *CatalogSearchRequestModel) UnsetAdminFolder()`

UnsetAdminFolder ensures that no value is present for AdminFolder, not even an explicit nil
### GetIncludeSubFolders

`func (o *CatalogSearchRequestModel) GetIncludeSubFolders() bool`

GetIncludeSubFolders returns the IncludeSubFolders field if non-nil, zero value otherwise.

### GetIncludeSubFoldersOk

`func (o *CatalogSearchRequestModel) GetIncludeSubFoldersOk() (*bool, bool)`

GetIncludeSubFoldersOk returns a tuple with the IncludeSubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubFolders

`func (o *CatalogSearchRequestModel) SetIncludeSubFolders(v bool)`

SetIncludeSubFolders sets IncludeSubFolders field to given value.

### HasIncludeSubFolders

`func (o *CatalogSearchRequestModel) HasIncludeSubFolders() bool`

HasIncludeSubFolders returns a boolean if a field has been set.

### SetIncludeSubFoldersNil

`func (o *CatalogSearchRequestModel) SetIncludeSubFoldersNil(b bool)`

 SetIncludeSubFoldersNil sets the value for IncludeSubFolders to be an explicit nil

### UnsetIncludeSubFolders
`func (o *CatalogSearchRequestModel) UnsetIncludeSubFolders()`

UnsetIncludeSubFolders ensures that no value is present for IncludeSubFolders, not even an explicit nil
### GetBasicSearchString

`func (o *CatalogSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *CatalogSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *CatalogSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *CatalogSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *CatalogSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *CatalogSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchFilters

`func (o *CatalogSearchRequestModel) GetSearchFilters() []CatalogSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *CatalogSearchRequestModel) GetSearchFiltersOk() (*[]CatalogSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *CatalogSearchRequestModel) SetSearchFilters(v []CatalogSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *CatalogSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *CatalogSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *CatalogSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *CatalogSearchRequestModel) GetSortCriteria() CatalogSortCriteriaRequestModel`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *CatalogSearchRequestModel) GetSortCriteriaOk() (*CatalogSortCriteriaRequestModel, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *CatalogSearchRequestModel) SetSortCriteria(v CatalogSortCriteriaRequestModel)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *CatalogSearchRequestModel) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.

### GetSearchFilterGroups

`func (o *CatalogSearchRequestModel) GetSearchFilterGroups() []CatalogSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *CatalogSearchRequestModel) GetSearchFilterGroupsOk() (*[]CatalogSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *CatalogSearchRequestModel) SetSearchFilterGroups(v []CatalogSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *CatalogSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *CatalogSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *CatalogSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *CatalogSearchRequestModel) GetSearchFilterGroupsType() CatalogSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *CatalogSearchRequestModel) GetSearchFilterGroupsTypeOk() (*CatalogSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *CatalogSearchRequestModel) SetSearchFilterGroupsType(v CatalogSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *CatalogSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


