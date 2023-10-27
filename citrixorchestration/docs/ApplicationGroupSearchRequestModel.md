# ApplicationGroupSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFolder** | Pointer to **NullableString** | Application folder in which to search for the application. | [optional] 
**IncludeSubFolders** | Pointer to **NullableBool** | Whether to include subfolders of the ApplicationFolder in the search for the application. | [optional] [default to true]
**ApplicationGroup** | Pointer to **NullableString** | Application group in which to search for the application. | [optional] 
**BasicSearchString** | Pointer to **NullableString** | Basic search string.  Specify a string which will match if contained within the application name. | [optional] 
**SearchFilters** | Pointer to [**[]AppSearchFilterRequestModel**](AppSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**AppSortCriteriaRequestModel**](AppSortCriteriaRequestModel.md) |  | [optional] 
**SearchFilterGroups** | Pointer to [**[]AppSearchFilterGroupRequestModel**](AppSearchFilterGroupRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**AppSearchFilterGroupsType**](AppSearchFilterGroupsType.md) |  | [optional] 
**ApplicationGroupFolder** | Pointer to **NullableString** | Application group folder in which to search for the application group. | [optional] 

## Methods

### NewApplicationGroupSearchRequestModel

`func NewApplicationGroupSearchRequestModel() *ApplicationGroupSearchRequestModel`

NewApplicationGroupSearchRequestModel instantiates a new ApplicationGroupSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupSearchRequestModelWithDefaults

`func NewApplicationGroupSearchRequestModelWithDefaults() *ApplicationGroupSearchRequestModel`

NewApplicationGroupSearchRequestModelWithDefaults instantiates a new ApplicationGroupSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFolder

`func (o *ApplicationGroupSearchRequestModel) GetApplicationFolder() string`

GetApplicationFolder returns the ApplicationFolder field if non-nil, zero value otherwise.

### GetApplicationFolderOk

`func (o *ApplicationGroupSearchRequestModel) GetApplicationFolderOk() (*string, bool)`

GetApplicationFolderOk returns a tuple with the ApplicationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFolder

`func (o *ApplicationGroupSearchRequestModel) SetApplicationFolder(v string)`

SetApplicationFolder sets ApplicationFolder field to given value.

### HasApplicationFolder

`func (o *ApplicationGroupSearchRequestModel) HasApplicationFolder() bool`

HasApplicationFolder returns a boolean if a field has been set.

### SetApplicationFolderNil

`func (o *ApplicationGroupSearchRequestModel) SetApplicationFolderNil(b bool)`

 SetApplicationFolderNil sets the value for ApplicationFolder to be an explicit nil

### UnsetApplicationFolder
`func (o *ApplicationGroupSearchRequestModel) UnsetApplicationFolder()`

UnsetApplicationFolder ensures that no value is present for ApplicationFolder, not even an explicit nil
### GetIncludeSubFolders

`func (o *ApplicationGroupSearchRequestModel) GetIncludeSubFolders() bool`

GetIncludeSubFolders returns the IncludeSubFolders field if non-nil, zero value otherwise.

### GetIncludeSubFoldersOk

`func (o *ApplicationGroupSearchRequestModel) GetIncludeSubFoldersOk() (*bool, bool)`

GetIncludeSubFoldersOk returns a tuple with the IncludeSubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubFolders

`func (o *ApplicationGroupSearchRequestModel) SetIncludeSubFolders(v bool)`

SetIncludeSubFolders sets IncludeSubFolders field to given value.

### HasIncludeSubFolders

`func (o *ApplicationGroupSearchRequestModel) HasIncludeSubFolders() bool`

HasIncludeSubFolders returns a boolean if a field has been set.

### SetIncludeSubFoldersNil

`func (o *ApplicationGroupSearchRequestModel) SetIncludeSubFoldersNil(b bool)`

 SetIncludeSubFoldersNil sets the value for IncludeSubFolders to be an explicit nil

### UnsetIncludeSubFolders
`func (o *ApplicationGroupSearchRequestModel) UnsetIncludeSubFolders()`

UnsetIncludeSubFolders ensures that no value is present for IncludeSubFolders, not even an explicit nil
### GetApplicationGroup

`func (o *ApplicationGroupSearchRequestModel) GetApplicationGroup() string`

GetApplicationGroup returns the ApplicationGroup field if non-nil, zero value otherwise.

### GetApplicationGroupOk

`func (o *ApplicationGroupSearchRequestModel) GetApplicationGroupOk() (*string, bool)`

GetApplicationGroupOk returns a tuple with the ApplicationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroup

`func (o *ApplicationGroupSearchRequestModel) SetApplicationGroup(v string)`

SetApplicationGroup sets ApplicationGroup field to given value.

### HasApplicationGroup

`func (o *ApplicationGroupSearchRequestModel) HasApplicationGroup() bool`

HasApplicationGroup returns a boolean if a field has been set.

### SetApplicationGroupNil

`func (o *ApplicationGroupSearchRequestModel) SetApplicationGroupNil(b bool)`

 SetApplicationGroupNil sets the value for ApplicationGroup to be an explicit nil

### UnsetApplicationGroup
`func (o *ApplicationGroupSearchRequestModel) UnsetApplicationGroup()`

UnsetApplicationGroup ensures that no value is present for ApplicationGroup, not even an explicit nil
### GetBasicSearchString

`func (o *ApplicationGroupSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *ApplicationGroupSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *ApplicationGroupSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *ApplicationGroupSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *ApplicationGroupSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *ApplicationGroupSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchFilters

`func (o *ApplicationGroupSearchRequestModel) GetSearchFilters() []AppSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *ApplicationGroupSearchRequestModel) GetSearchFiltersOk() (*[]AppSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *ApplicationGroupSearchRequestModel) SetSearchFilters(v []AppSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *ApplicationGroupSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *ApplicationGroupSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *ApplicationGroupSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *ApplicationGroupSearchRequestModel) GetSortCriteria() AppSortCriteriaRequestModel`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *ApplicationGroupSearchRequestModel) GetSortCriteriaOk() (*AppSortCriteriaRequestModel, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *ApplicationGroupSearchRequestModel) SetSortCriteria(v AppSortCriteriaRequestModel)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *ApplicationGroupSearchRequestModel) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.

### GetSearchFilterGroups

`func (o *ApplicationGroupSearchRequestModel) GetSearchFilterGroups() []AppSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *ApplicationGroupSearchRequestModel) GetSearchFilterGroupsOk() (*[]AppSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *ApplicationGroupSearchRequestModel) SetSearchFilterGroups(v []AppSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *ApplicationGroupSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *ApplicationGroupSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *ApplicationGroupSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *ApplicationGroupSearchRequestModel) GetSearchFilterGroupsType() AppSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *ApplicationGroupSearchRequestModel) GetSearchFilterGroupsTypeOk() (*AppSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *ApplicationGroupSearchRequestModel) SetSearchFilterGroupsType(v AppSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *ApplicationGroupSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetApplicationGroupFolder

`func (o *ApplicationGroupSearchRequestModel) GetApplicationGroupFolder() string`

GetApplicationGroupFolder returns the ApplicationGroupFolder field if non-nil, zero value otherwise.

### GetApplicationGroupFolderOk

`func (o *ApplicationGroupSearchRequestModel) GetApplicationGroupFolderOk() (*string, bool)`

GetApplicationGroupFolderOk returns a tuple with the ApplicationGroupFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroupFolder

`func (o *ApplicationGroupSearchRequestModel) SetApplicationGroupFolder(v string)`

SetApplicationGroupFolder sets ApplicationGroupFolder field to given value.

### HasApplicationGroupFolder

`func (o *ApplicationGroupSearchRequestModel) HasApplicationGroupFolder() bool`

HasApplicationGroupFolder returns a boolean if a field has been set.

### SetApplicationGroupFolderNil

`func (o *ApplicationGroupSearchRequestModel) SetApplicationGroupFolderNil(b bool)`

 SetApplicationGroupFolderNil sets the value for ApplicationGroupFolder to be an explicit nil

### UnsetApplicationGroupFolder
`func (o *ApplicationGroupSearchRequestModel) UnsetApplicationGroupFolder()`

UnsetApplicationGroupFolder ensures that no value is present for ApplicationGroupFolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


