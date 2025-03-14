# MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFilterGroupType** | Pointer to [**MaintenanceCycleVmOperationJobSearchFilterGroupType**](MaintenanceCycleVmOperationJobSearchFilterGroupType.md) |  | [optional] 
**SearchFilters** | Pointer to [**[]MaintenanceCycleVMOperationJobSearchFilterRequestModel**](MaintenanceCycleVMOperationJobSearchFilterRequestModel.md) | The search filters in search filter group | [optional] 
**SearchFilterGroupsType** | Pointer to [**MaintenanceCycleVmOperationJobSearchFilterGroupsType**](MaintenanceCycleVmOperationJobSearchFilterGroupsType.md) |  | [optional] 
**SearchFilterGroups** | Pointer to [**[]MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel**](MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel.md) | The search filter group in search filter groups | [optional] 

## Methods

### NewMaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel

`func NewMaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel() *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel`

NewMaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel instantiates a new MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceCycleVmOperationJobSearchFilterGroupsRequestModelWithDefaults

`func NewMaintenanceCycleVmOperationJobSearchFilterGroupsRequestModelWithDefaults() *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel`

NewMaintenanceCycleVmOperationJobSearchFilterGroupsRequestModelWithDefaults instantiates a new MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFilterGroupType

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) GetSearchFilterGroupType() MaintenanceCycleVmOperationJobSearchFilterGroupType`

GetSearchFilterGroupType returns the SearchFilterGroupType field if non-nil, zero value otherwise.

### GetSearchFilterGroupTypeOk

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) GetSearchFilterGroupTypeOk() (*MaintenanceCycleVmOperationJobSearchFilterGroupType, bool)`

GetSearchFilterGroupTypeOk returns a tuple with the SearchFilterGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupType

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) SetSearchFilterGroupType(v MaintenanceCycleVmOperationJobSearchFilterGroupType)`

SetSearchFilterGroupType sets SearchFilterGroupType field to given value.

### HasSearchFilterGroupType

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) HasSearchFilterGroupType() bool`

HasSearchFilterGroupType returns a boolean if a field has been set.

### GetSearchFilters

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) GetSearchFilters() []MaintenanceCycleVMOperationJobSearchFilterRequestModel`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) GetSearchFiltersOk() (*[]MaintenanceCycleVMOperationJobSearchFilterRequestModel, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) SetSearchFilters(v []MaintenanceCycleVMOperationJobSearchFilterRequestModel)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### SetSearchFiltersNil

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) SetSearchFiltersNil(b bool)`

 SetSearchFiltersNil sets the value for SearchFilters to be an explicit nil

### UnsetSearchFilters
`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) UnsetSearchFilters()`

UnsetSearchFilters ensures that no value is present for SearchFilters, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) GetSearchFilterGroupsType() MaintenanceCycleVmOperationJobSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) GetSearchFilterGroupsTypeOk() (*MaintenanceCycleVmOperationJobSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) SetSearchFilterGroupsType(v MaintenanceCycleVmOperationJobSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetSearchFilterGroups

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) GetSearchFilterGroups() []MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) GetSearchFilterGroupsOk() (*[]MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) SetSearchFilterGroups(v []MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *MaintenanceCycleVmOperationJobSearchFilterGroupsRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


