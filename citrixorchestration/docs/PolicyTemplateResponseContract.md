# PolicyTemplateResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]PolicyResponseContract**](PolicyResponseContract.md) | Policies | [optional] 
**Templates** | Pointer to [**[]TemplateResponseContract**](TemplateResponseContract.md) | Templates | [optional] 

## Methods

### NewPolicyTemplateResponseContract

`func NewPolicyTemplateResponseContract() *PolicyTemplateResponseContract`

NewPolicyTemplateResponseContract instantiates a new PolicyTemplateResponseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateResponseContractWithDefaults

`func NewPolicyTemplateResponseContractWithDefaults() *PolicyTemplateResponseContract`

NewPolicyTemplateResponseContractWithDefaults instantiates a new PolicyTemplateResponseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *PolicyTemplateResponseContract) GetPolicies() []PolicyResponseContract`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *PolicyTemplateResponseContract) GetPoliciesOk() (*[]PolicyResponseContract, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *PolicyTemplateResponseContract) SetPolicies(v []PolicyResponseContract)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *PolicyTemplateResponseContract) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTemplates

`func (o *PolicyTemplateResponseContract) GetTemplates() []TemplateResponseContract`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *PolicyTemplateResponseContract) GetTemplatesOk() (*[]TemplateResponseContract, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *PolicyTemplateResponseContract) SetTemplates(v []TemplateResponseContract)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *PolicyTemplateResponseContract) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


