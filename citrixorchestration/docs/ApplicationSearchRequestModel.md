# ApplicationSearchRequestModel

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

## Methods

### NewApplicationSearchRequestModel

`func NewApplicationSearchRequestModel() *ApplicationSearchRequestModel`

NewApplicationSearchRequestModel instantiates a new ApplicationSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSearchRequestModelWithDefaults

`func NewApplicationSearchRequestModelWithDefaults() *ApplicationSearchRequestModel`

NewApplicationSearchRequestModelWithDefaults instantiates a new ApplicationSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFolder

`func (o *ApplicationSearchRequestModel) GetApplicationFolder() string`

GetApplicationFolder returns the ApplicationFolder field if non-nil, zero value otherwise.

### GetApplicationFolderOk

`func (o *ApplicationSearchRequestModel) GetApplicationFolderOk() (*string, bool)`

GetApplicationFolderOk returns a tuple with the ApplicationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFolder

`func (o *ApplicationSearchRequestModel) SetApplicationFolder(v string)`

SetApplicationFolder sets ApplicationFolder field to given value.

### HasApplicationFolder

`func (o *ApplicationSearchRequestModel) HasApplicationFolder() bool`

HasApplicationFolder returns a boolean if a field has been set.

### SetApplicationFolderNil

`func (o *ApplicationSearchRequestModel) SetApplicationFolderNil(b bool)`

 SetApplicationFolderNil sets the value for ApplicationFolder to be an explicit nil

### UnsetApplicationFolder
`func (o *ApplicationSearchRequestModel) UnsetApplicationFolder()`

UnsetApplicationFolder ensures that no value is present for ApplicationFolder, not even an explicit nil
### GetIncludeSubFolders

`func (o *ApplicationSearchRequestModel) GetIncludeSubFolders() bool`

GetIncludeSubFolders returns the IncludeSubFolders field if non-nil, zero value otherwise.

### GetIncludeSubFoldersOk

`func (o *ApplicationSearchRequestModel) GetIncludeSubFoldersOk() (*bool, bool)`

GetIncludeSubFoldersOk returns a tuple with the IncludeSubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubFolders

`func (o *ApplicationSearchRequestModel) SetIncludeSubFolders(v bool)`

SetIncludeSubFolders sets IncludeSubFolders field to given value.

### HasIncludeSubFolders

`func (o *ApplicationSearchRequestModel) HasIncludeSubFolders() bool`

HasIncludeSubFolders returns a boolean if a field has been set.

### SetIncludeSubFoldersNil

`func (o *ApplicationSearchRequestModel) SetIncludeSubFoldersNil(b bool)`

 SetIncludeSubFoldersNil sets the value for IncludeSubFolders to be an explicit nil

### UnsetIncludeSubFolders
`func (o *ApplicationSearchRequestModel) UnsetIncludeSubFolders()`

UnsetIncludeSubFolders ensures that no value is present for IncludeSubFolders, not even an explicit nil
### GetApplicationGroup

`func (o *ApplicationSearchRequestModel) GetApplicationGroup() string`

GetApplicationGroup returns the ApplicationGroup field if non-nil, zero value otherwise.

### GetApplicationGroupOk

`func (o *ApplicationSearchRequestModel) GetApplicationGroupOk() (*string, bool)`

GetApplicationGroupOk returns a tuple with the ApplicationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroup

`func (o *ApplicationSearchRequestModel) SetApplicationGroup(v string)`

SetApplicationGroup sets ApplicationGroup field to given value.

### HasApplicationGroup

`func (o *ApplicationSearchRequestModel) HasApplicationGroup() bool`

HasApplicationGroup returns a boolean if a field has been set.

### SetApplicationGroupNil

`func (o *ApplicationSearchRequestModel) SetApplicationGroupNil(b bool)`

 SetApplicationGroupNil sets the value for ApplicationGroup to be an explicit nil

### UnsetApplicationGroup
`func (o *ApplicationSearchRequestModel) UnsetApplicationGroup()`

UnsetApplicationGroup ensures that no value is present for ApplicationGroup, not even an explicit nil
### GetBasicSearchString

`func (o *ApplicationSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *ApplicationSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *ApplicationSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *ApplicationSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *ApplicationSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *ApplicationSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchFilters

`func (o *ApplicationSearchRequestModel) GetSearchFilters() []AppSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *ApplicationSearchRequestModel) GetSearchFiltersOk() (*[]AppSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *ApplicationSearchRequestModel) SetSearchFilters(v []AppSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *ApplicationSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *ApplicationSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *ApplicationSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *ApplicationSearchRequestModel) GetSortCriteria() AppSortCriteriaRequestModel`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *ApplicationSearchRequestModel) GetSortCriteriaOk() (*AppSortCriteriaRequestModel, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *ApplicationSearchRequestModel) SetSortCriteria(v AppSortCriteriaRequestModel)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *ApplicationSearchRequestModel) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.

### GetSearchFilterGroups

`func (o *ApplicationSearchRequestModel) GetSearchFilterGroups() []AppSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *ApplicationSearchRequestModel) GetSearchFilterGroupsOk() (*[]AppSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *ApplicationSearchRequestModel) SetSearchFilterGroups(v []AppSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *ApplicationSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *ApplicationSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *ApplicationSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *ApplicationSearchRequestModel) GetSearchFilterGroupsType() AppSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *ApplicationSearchRequestModel) GetSearchFilterGroupsTypeOk() (*AppSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *ApplicationSearchRequestModel) SetSearchFilterGroupsType(v AppSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *ApplicationSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


