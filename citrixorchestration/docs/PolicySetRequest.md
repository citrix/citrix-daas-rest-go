# PolicySetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicySetType** | Pointer to **NullableString** | The policy set type. | [optional] 
**Name** | Pointer to **NullableString** | Name of the policy set. | [optional] 
**Description** | Pointer to **NullableString** | Policy set description. | [optional] 
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

### SetPolicySetTypeNil

`func (o *PolicySetRequest) SetPolicySetTypeNil(b bool)`

 SetPolicySetTypeNil sets the value for PolicySetType to be an explicit nil

### UnsetPolicySetType
`func (o *PolicySetRequest) UnsetPolicySetType()`

UnsetPolicySetType ensures that no value is present for PolicySetType, not even an explicit nil
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

### SetNameNil

`func (o *PolicySetRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PolicySetRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### SetDescriptionNil

`func (o *PolicySetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicySetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### SetScopesNil

`func (o *PolicySetRequest) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *PolicySetRequest) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


