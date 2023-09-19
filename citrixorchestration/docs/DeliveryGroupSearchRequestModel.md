# DeliveryGroupSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminFolder** | Pointer to **NullableString** | Admin folder in which to search for the delivery group. | [optional] 
**IncludeSubFolders** | Pointer to **NullableBool** | Whether to include subfolders of the AdminFolder in the search for the delivery groups. | [optional] [default to true]
**BasicSearchString** | Pointer to **NullableString** | Basic search string. Specify a string which will match if contained within some string property of the deliveryGroup. | [optional] 
**SearchFilters** | Pointer to [**[]DeliveryGroupSearchFilterRequestModel**](DeliveryGroupSearchFilterRequestModel.md) | List of advanced search filters. | [optional] 
**SortCriteria** | Pointer to [**DeliveryGroupSortCriteriaRequestModel**](DeliveryGroupSortCriteriaRequestModel.md) |  | [optional] 
**SearchFilterGroups** | Pointer to [**[]DeliveryGroupSearchFilterGroupRequestModel**](DeliveryGroupSearchFilterGroupRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**DeliveryGroupSearchFilterGroupsType**](DeliveryGroupSearchFilterGroupsType.md) |  | [optional] 

## Methods

### NewDeliveryGroupSearchRequestModel

`func NewDeliveryGroupSearchRequestModel() *DeliveryGroupSearchRequestModel`

NewDeliveryGroupSearchRequestModel instantiates a new DeliveryGroupSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupSearchRequestModelWithDefaults

`func NewDeliveryGroupSearchRequestModelWithDefaults() *DeliveryGroupSearchRequestModel`

NewDeliveryGroupSearchRequestModelWithDefaults instantiates a new DeliveryGroupSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminFolder

`func (o *DeliveryGroupSearchRequestModel) GetAdminFolder() string`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *DeliveryGroupSearchRequestModel) GetAdminFolderOk() (*string, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *DeliveryGroupSearchRequestModel) SetAdminFolder(v string)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *DeliveryGroupSearchRequestModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### SetAdminFolderNil

`func (o *DeliveryGroupSearchRequestModel) SetAdminFolderNil(b bool)`

 SetAdminFolderNil sets the value for AdminFolder to be an explicit nil

### UnsetAdminFolder
`func (o *DeliveryGroupSearchRequestModel) UnsetAdminFolder()`

UnsetAdminFolder ensures that no value is present for AdminFolder, not even an explicit nil
### GetIncludeSubFolders

`func (o *DeliveryGroupSearchRequestModel) GetIncludeSubFolders() bool`

GetIncludeSubFolders returns the IncludeSubFolders field if non-nil, zero value otherwise.

### GetIncludeSubFoldersOk

`func (o *DeliveryGroupSearchRequestModel) GetIncludeSubFoldersOk() (*bool, bool)`

GetIncludeSubFoldersOk returns a tuple with the IncludeSubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubFolders

`func (o *DeliveryGroupSearchRequestModel) SetIncludeSubFolders(v bool)`

SetIncludeSubFolders sets IncludeSubFolders field to given value.

### HasIncludeSubFolders

`func (o *DeliveryGroupSearchRequestModel) HasIncludeSubFolders() bool`

HasIncludeSubFolders returns a boolean if a field has been set.

### SetIncludeSubFoldersNil

`func (o *DeliveryGroupSearchRequestModel) SetIncludeSubFoldersNil(b bool)`

 SetIncludeSubFoldersNil sets the value for IncludeSubFolders to be an explicit nil

### UnsetIncludeSubFolders
`func (o *DeliveryGroupSearchRequestModel) UnsetIncludeSubFolders()`

UnsetIncludeSubFolders ensures that no value is present for IncludeSubFolders, not even an explicit nil
### GetBasicSearchString

`func (o *DeliveryGroupSearchRequestModel) GetBasicSearchString() string`

GetBasicSearchString returns the BasicSearchString field if non-nil, zero value otherwise.

### GetBasicSearchStringOk

`func (o *DeliveryGroupSearchRequestModel) GetBasicSearchStringOk() (*string, bool)`

GetBasicSearchStringOk returns a tuple with the BasicSearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSearchString

`func (o *DeliveryGroupSearchRequestModel) SetBasicSearchString(v string)`

SetBasicSearchString sets BasicSearchString field to given value.

### HasBasicSearchString

`func (o *DeliveryGroupSearchRequestModel) HasBasicSearchString() bool`

HasBasicSearchString returns a boolean if a field has been set.

### SetBasicSearchStringNil

`func (o *DeliveryGroupSearchRequestModel) SetBasicSearchStringNil(b bool)`

 SetBasicSearchStringNil sets the value for BasicSearchString to be an explicit nil

### UnsetBasicSearchString
`func (o *DeliveryGroupSearchRequestModel) UnsetBasicSearchString()`

UnsetBasicSearchString ensures that no value is present for BasicSearchString, not even an explicit nil
### GetSearchFilters

`func (o *DeliveryGroupSearchRequestModel) GetSearchFilters() []DeliveryGroupSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *DeliveryGroupSearchRequestModel) GetSearchFiltersOk() (*[]DeliveryGroupSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *DeliveryGroupSearchRequestModel) SetSearchFilters(v []DeliveryGroupSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *DeliveryGroupSearchRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *DeliveryGroupSearchRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *DeliveryGroupSearchRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSortCriteria

`func (o *DeliveryGroupSearchRequestModel) GetSortCriteria() DeliveryGroupSortCriteriaRequestModel`

GetSortCriteria returns the SortCriteria field if non-nil, zero value otherwise.

### GetSortCriteriaOk

`func (o *DeliveryGroupSearchRequestModel) GetSortCriteriaOk() (*DeliveryGroupSortCriteriaRequestModel, bool)`

GetSortCriteriaOk returns a tuple with the SortCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCriteria

`func (o *DeliveryGroupSearchRequestModel) SetSortCriteria(v DeliveryGroupSortCriteriaRequestModel)`

SetSortCriteria sets SortCriteria field to given value.

### HasSortCriteria

`func (o *DeliveryGroupSearchRequestModel) HasSortCriteria() bool`

HasSortCriteria returns a boolean if a field has been set.

### GetSearchFilterGroups

`func (o *DeliveryGroupSearchRequestModel) GetSearchFilterGroups() []DeliveryGroupSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *DeliveryGroupSearchRequestModel) GetSearchFilterGroupsOk() (*[]DeliveryGroupSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *DeliveryGroupSearchRequestModel) SetSearchFilterGroups(v []DeliveryGroupSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *DeliveryGroupSearchRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *DeliveryGroupSearchRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *DeliveryGroupSearchRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *DeliveryGroupSearchRequestModel) GetSearchFilterGroupsType() DeliveryGroupSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *DeliveryGroupSearchRequestModel) GetSearchFilterGroupsTypeOk() (*DeliveryGroupSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *DeliveryGroupSearchRequestModel) SetSearchFilterGroupsType(v DeliveryGroupSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *DeliveryGroupSearchRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


