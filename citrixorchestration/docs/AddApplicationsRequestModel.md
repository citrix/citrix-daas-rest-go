# AddApplicationsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingApplications** | Pointer to **[]string** | List of existing applications to be added to the DeliveryGroups and/or ApplicationGroups. | [optional] 
**NewApplications** | Pointer to [**[]CreateApplicationRequestModel**](CreateApplicationRequestModel.md) | List of applications which should be created and added to the DeliveryGroups and/or ApplicationGroups. | [optional] 
**ApplicationGroups** | Pointer to **[]string** | List of application groups to add the ExistingApplications and NewApplications to. | [optional] 
**DeliveryGroups** | Pointer to [**[]PriorityRefRequestModel**](PriorityRefRequestModel.md) | List of delivery groups to add the ExistingApplications and NewApplications to. | [optional] 

## Methods

### NewAddApplicationsRequestModel

`func NewAddApplicationsRequestModel() *AddApplicationsRequestModel`

NewAddApplicationsRequestModel instantiates a new AddApplicationsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddApplicationsRequestModelWithDefaults

`func NewAddApplicationsRequestModelWithDefaults() *AddApplicationsRequestModel`

NewAddApplicationsRequestModelWithDefaults instantiates a new AddApplicationsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingApplications

`func (o *AddApplicationsRequestModel) GetExistingApplications() []string`

GetExistingApplications returns the ExistingApplications field if non-nil, zero value otherwise.

### GetExistingApplicationsOk

`func (o *AddApplicationsRequestModel) GetExistingApplicationsOk() (*[]string, bool)`

GetExistingApplicationsOk returns a tuple with the ExistingApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingApplications

`func (o *AddApplicationsRequestModel) SetExistingApplications(v []string)`

SetExistingApplications sets ExistingApplications field to given value.

### HasExistingApplications

`func (o *AddApplicationsRequestModel) HasExistingApplications() bool`

HasExistingApplications returns a boolean if a field has been set.

### GetNewApplications

`func (o *AddApplicationsRequestModel) GetNewApplications() []CreateApplicationRequestModel`

GetNewApplications returns the NewApplications field if non-nil, zero value otherwise.

### GetNewApplicationsOk

`func (o *AddApplicationsRequestModel) GetNewApplicationsOk() (*[]CreateApplicationRequestModel, bool)`

GetNewApplicationsOk returns a tuple with the NewApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewApplications

`func (o *AddApplicationsRequestModel) SetNewApplications(v []CreateApplicationRequestModel)`

SetNewApplications sets NewApplications field to given value.

### HasNewApplications

`func (o *AddApplicationsRequestModel) HasNewApplications() bool`

HasNewApplications returns a boolean if a field has been set.

### GetApplicationGroups

`func (o *AddApplicationsRequestModel) GetApplicationGroups() []string`

GetApplicationGroups returns the ApplicationGroups field if non-nil, zero value otherwise.

### GetApplicationGroupsOk

`func (o *AddApplicationsRequestModel) GetApplicationGroupsOk() (*[]string, bool)`

GetApplicationGroupsOk returns a tuple with the ApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroups

`func (o *AddApplicationsRequestModel) SetApplicationGroups(v []string)`

SetApplicationGroups sets ApplicationGroups field to given value.

### HasApplicationGroups

`func (o *AddApplicationsRequestModel) HasApplicationGroups() bool`

HasApplicationGroups returns a boolean if a field has been set.

### GetDeliveryGroups

`func (o *AddApplicationsRequestModel) GetDeliveryGroups() []PriorityRefRequestModel`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *AddApplicationsRequestModel) GetDeliveryGroupsOk() (*[]PriorityRefRequestModel, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *AddApplicationsRequestModel) SetDeliveryGroups(v []PriorityRefRequestModel)`

SetDeliveryGroups sets DeliveryGroups field to given value.

### HasDeliveryGroups

`func (o *AddApplicationsRequestModel) HasDeliveryGroups() bool`

HasDeliveryGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


