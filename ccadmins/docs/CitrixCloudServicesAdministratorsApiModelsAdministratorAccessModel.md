# CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType**](CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType.md) |  | 
**Policies** | Pointer to [**[]CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel**](CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel.md) |  | [optional] 

## Methods

### NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel

`func NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel(accessType CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType, ) *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel`

NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessModelWithDefaults

`func NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessModelWithDefaults() *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel`

NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessModelWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel) GetAccessType() CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel) GetAccessTypeOk() (*CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel) SetAccessType(v CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType)`

SetAccessType sets AccessType field to given value.


### GetPolicies

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel) GetPolicies() []CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel) GetPoliciesOk() (*[]CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel) SetPolicies(v []CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


