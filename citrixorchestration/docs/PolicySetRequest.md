# PolicySetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicySetType** | Pointer to **string** | The policy set type. | [optional] 
**Name** | Pointer to **string** | Name of the policy set. | [optional] 
**Description** | Pointer to **string** | Policy set description. | [optional] 
**Scopes** | Pointer to **[]string** | Delegated administration scopes for the policy set. | [optional] 

## Methods

### NewPolicySetRequest

`func NewPolicySetRequest() *PolicySetRequest`

NewPolicySetRequest instantiates a new PolicySetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySetRequestWithDefaults

`func NewPolicySetRequestWithDefaults() *PolicySetRequest`

NewPolicySetRequestWithDefaults instantiates a new PolicySetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicySetType

`func (o *PolicySetRequest) GetPolicySetType() string`

GetPolicySetType returns the PolicySetType field if non-nil, zero value otherwise.

### GetPolicySetTypeOk

`func (o *PolicySetRequest) GetPolicySetTypeOk() (*string, bool)`

GetPolicySetTypeOk returns a tuple with the PolicySetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetType

`func (o *PolicySetRequest) SetPolicySetType(v string)`

SetPolicySetType sets PolicySetType field to given value.

### HasPolicySetType

`func (o *PolicySetRequest) HasPolicySetType() bool`

HasPolicySetType returns a boolean if a field has been set.

### GetName

`func (o *PolicySetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicySetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicySetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicySetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicySetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicySetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicySetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicySetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScopes

`func (o *PolicySetRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PolicySetRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PolicySetRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PolicySetRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


