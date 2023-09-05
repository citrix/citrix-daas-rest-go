# DeliveryGroupAddApplicationsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingApplications** | Pointer to [**[]PriorityRefRequestModel**](PriorityRefRequestModel.md) | List of existing applications to be associated with the delivery group. | [optional] 
**NewApplications** | Pointer to [**[]CreateApplicationRequestModel**](CreateApplicationRequestModel.md) | List of applications which should be created and associated with the delivery group. | [optional] 
**ExistingApplicationGroups** | Pointer to [**[]PriorityRefRequestModel**](PriorityRefRequestModel.md) | List of existing application groups to be associated with the delivery group. | [optional] 

## Methods

### NewDeliveryGroupAddApplicationsRequestModel

`func NewDeliveryGroupAddApplicationsRequestModel() *DeliveryGroupAddApplicationsRequestModel`

NewDeliveryGroupAddApplicationsRequestModel instantiates a new DeliveryGroupAddApplicationsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupAddApplicationsRequestModelWithDefaults

`func NewDeliveryGroupAddApplicationsRequestModelWithDefaults() *DeliveryGroupAddApplicationsRequestModel`

NewDeliveryGroupAddApplicationsRequestModelWithDefaults instantiates a new DeliveryGroupAddApplicationsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingApplications

`func (o *DeliveryGroupAddApplicationsRequestModel) GetExistingApplications() []PriorityRefRequestModel`

GetExistingApplications returns the ExistingApplications field if non-nil, zero value otherwise.

### GetExistingApplicationsOk

`func (o *DeliveryGroupAddApplicationsRequestModel) GetExistingApplicationsOk() (*[]PriorityRefRequestModel, bool)`

GetExistingApplicationsOk returns a tuple with the ExistingApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingApplications

`func (o *DeliveryGroupAddApplicationsRequestModel) SetExistingApplications(v []PriorityRefRequestModel)`

SetExistingApplications sets ExistingApplications field to given value.

### HasExistingApplications

`func (o *DeliveryGroupAddApplicationsRequestModel) HasExistingApplications() bool`

HasExistingApplications returns a boolean if a field has been set.

### GetNewApplications

`func (o *DeliveryGroupAddApplicationsRequestModel) GetNewApplications() []CreateApplicationRequestModel`

GetNewApplications returns the NewApplications field if non-nil, zero value otherwise.

### GetNewApplicationsOk

`func (o *DeliveryGroupAddApplicationsRequestModel) GetNewApplicationsOk() (*[]CreateApplicationRequestModel, bool)`

GetNewApplicationsOk returns a tuple with the NewApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewApplications

`func (o *DeliveryGroupAddApplicationsRequestModel) SetNewApplications(v []CreateApplicationRequestModel)`

SetNewApplications sets NewApplications field to given value.

### HasNewApplications

`func (o *DeliveryGroupAddApplicationsRequestModel) HasNewApplications() bool`

HasNewApplications returns a boolean if a field has been set.

### GetExistingApplicationGroups

`func (o *DeliveryGroupAddApplicationsRequestModel) GetExistingApplicationGroups() []PriorityRefRequestModel`

GetExistingApplicationGroups returns the ExistingApplicationGroups field if non-nil, zero value otherwise.

### GetExistingApplicationGroupsOk

`func (o *DeliveryGroupAddApplicationsRequestModel) GetExistingApplicationGroupsOk() (*[]PriorityRefRequestModel, bool)`

GetExistingApplicationGroupsOk returns a tuple with the ExistingApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingApplicationGroups

`func (o *DeliveryGroupAddApplicationsRequestModel) SetExistingApplicationGroups(v []PriorityRefRequestModel)`

SetExistingApplicationGroups sets ExistingApplicationGroups field to given value.

### HasExistingApplicationGroups

`func (o *DeliveryGroupAddApplicationsRequestModel) HasExistingApplicationGroups() bool`

HasExistingApplicationGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


