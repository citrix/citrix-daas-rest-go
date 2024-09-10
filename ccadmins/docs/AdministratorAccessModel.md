# AdministratorAccessModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**AdministratorAccessType**](AdministratorAccessType.md) |  | 
**Policies** | Pointer to [**[]AdministratorAccessPolicyModel**](AdministratorAccessPolicyModel.md) |  | [optional] 

## Methods

### NewAdministratorAccessModel

`func NewAdministratorAccessModel(accessType AdministratorAccessType, ) *AdministratorAccessModel`

NewAdministratorAccessModel instantiates a new AdministratorAccessModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorAccessModelWithDefaults

`func NewAdministratorAccessModelWithDefaults() *AdministratorAccessModel`

NewAdministratorAccessModelWithDefaults instantiates a new AdministratorAccessModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *AdministratorAccessModel) GetAccessType() AdministratorAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AdministratorAccessModel) GetAccessTypeOk() (*AdministratorAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AdministratorAccessModel) SetAccessType(v AdministratorAccessType)`

SetAccessType sets AccessType field to given value.


### GetPolicies

`func (o *AdministratorAccessModel) GetPolicies() []AdministratorAccessPolicyModel`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AdministratorAccessModel) GetPoliciesOk() (*[]AdministratorAccessPolicyModel, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AdministratorAccessModel) SetPolicies(v []AdministratorAccessPolicyModel)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AdministratorAccessModel) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *AdministratorAccessModel) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *AdministratorAccessModel) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


