# CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType**](CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType.md) |  | 
**Policies** | Pointer to [**[]CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel**](CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel.md) |  | [optional] 

## Methods

### NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel

`func NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel(accessType CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType, ) *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel`

NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel instantiates a new CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModelWithDefaults

`func NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModelWithDefaults() *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel`

NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModelWithDefaults instantiates a new CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel) GetAccessType() CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel) GetAccessTypeOk() (*CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel) SetAccessType(v CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType)`

SetAccessType sets AccessType field to given value.


### GetPolicies

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel) GetPolicies() []CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel) GetPoliciesOk() (*[]CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel) SetPolicies(v []CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


