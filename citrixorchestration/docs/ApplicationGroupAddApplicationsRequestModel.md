# ApplicationGroupAddApplicationsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingApplications** | Pointer to **[]string** | List of existing applications to be associated with the application group.  Optional.  Default is not to create any new associations of existing applications to the application group.   Each item can be specified by  or .   All applications associated with a application group must SharingKind equal to Shared. | [optional] 
**NewApplications** | Pointer to [**[]CreateApplicationRequestModel**](CreateApplicationRequestModel.md) | List of applications which should be created and associated with the application group. Optional.  Default is not to create any new applications associated with the application group. | [optional] 

## Methods

### NewApplicationGroupAddApplicationsRequestModel

`func NewApplicationGroupAddApplicationsRequestModel() *ApplicationGroupAddApplicationsRequestModel`

NewApplicationGroupAddApplicationsRequestModel instantiates a new ApplicationGroupAddApplicationsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupAddApplicationsRequestModelWithDefaults

`func NewApplicationGroupAddApplicationsRequestModelWithDefaults() *ApplicationGroupAddApplicationsRequestModel`

NewApplicationGroupAddApplicationsRequestModelWithDefaults instantiates a new ApplicationGroupAddApplicationsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingApplications

`func (o *ApplicationGroupAddApplicationsRequestModel) GetExistingApplications() []string`

GetExistingApplications returns the ExistingApplications field if non-nil, zero value otherwise.

### GetExistingApplicationsOk

`func (o *ApplicationGroupAddApplicationsRequestModel) GetExistingApplicationsOk() (*[]string, bool)`

GetExistingApplicationsOk returns a tuple with the ExistingApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingApplications

`func (o *ApplicationGroupAddApplicationsRequestModel) SetExistingApplications(v []string)`

SetExistingApplications sets ExistingApplications field to given value.

### HasExistingApplications

`func (o *ApplicationGroupAddApplicationsRequestModel) HasExistingApplications() bool`

HasExistingApplications returns a boolean if a field has been set.

### SetExistingApplicationsNil

`func (o *ApplicationGroupAddApplicationsRequestModel) SetExistingApplicationsNil(b bool)`

 SetExistingApplicationsNil sets the value for ExistingApplications to be an explicit nil

### UnsetExistingApplications
`func (o *ApplicationGroupAddApplicationsRequestModel) UnsetExistingApplications()`

UnsetExistingApplications ensures that no value is present for ExistingApplications, not even an explicit nil
### GetNewApplications

`func (o *ApplicationGroupAddApplicationsRequestModel) GetNewApplications() []CreateApplicationRequestModel`

GetNewApplications returns the NewApplications field if non-nil, zero value otherwise.

### GetNewApplicationsOk

`func (o *ApplicationGroupAddApplicationsRequestModel) GetNewApplicationsOk() (*[]CreateApplicationRequestModel, bool)`

GetNewApplicationsOk returns a tuple with the NewApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewApplications

`func (o *ApplicationGroupAddApplicationsRequestModel) SetNewApplications(v []CreateApplicationRequestModel)`

SetNewApplications sets NewApplications field to given value.

### HasNewApplications

`func (o *ApplicationGroupAddApplicationsRequestModel) HasNewApplications() bool`

HasNewApplications returns a boolean if a field has been set.

### SetNewApplicationsNil

`func (o *ApplicationGroupAddApplicationsRequestModel) SetNewApplicationsNil(b bool)`

 SetNewApplicationsNil sets the value for NewApplications to be an explicit nil

### UnsetNewApplications
`func (o *ApplicationGroupAddApplicationsRequestModel) UnsetNewApplications()`

UnsetNewApplications ensures that no value is present for NewApplications, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


