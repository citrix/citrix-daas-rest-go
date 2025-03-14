# PolicySetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicySetGuid** | Pointer to **string** | Guid of the policy set. | [optional] 
**PolicySetType** | Pointer to [**SdkGpoPolicySetType**](SdkGpoPolicySetType.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Name of the policy set. | [optional] 
**Description** | Pointer to **NullableString** | Policy set description. | [optional] 
**PolicyCount** | Pointer to **int32** | Number of policies in the policy set. | [optional] 
**Policies** | Pointer to [**[]PolicyResponse**](PolicyResponse.md) | Policies in the policy set. | [optional] 
**Scopes** | Pointer to **[]string** | Delegated admin scopes. | [optional] 
**IsAssigned** | Pointer to **bool** | Indicate if policy set is assigned to at least one delivery group. | [optional] 
**LastError** | Pointer to **NullableString** | The most recent error that occurred when the policy data was converted to/from policy set. | [optional] 

## Methods

### NewPolicySetResponse

`func NewPolicySetResponse() *PolicySetResponse`

NewPolicySetResponse instantiates a new PolicySetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySetResponseWithDefaults

`func NewPolicySetResponseWithDefaults() *PolicySetResponse`

NewPolicySetResponseWithDefaults instantiates a new PolicySetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicySetGuid

`func (o *PolicySetResponse) GetPolicySetGuid() string`

GetPolicySetGuid returns the PolicySetGuid field if non-nil, zero value otherwise.

### GetPolicySetGuidOk

`func (o *PolicySetResponse) GetPolicySetGuidOk() (*string, bool)`

GetPolicySetGuidOk returns a tuple with the PolicySetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetGuid

`func (o *PolicySetResponse) SetPolicySetGuid(v string)`

SetPolicySetGuid sets PolicySetGuid field to given value.

### HasPolicySetGuid

`func (o *PolicySetResponse) HasPolicySetGuid() bool`

HasPolicySetGuid returns a boolean if a field has been set.

### GetPolicySetType

`func (o *PolicySetResponse) GetPolicySetType() SdkGpoPolicySetType`

GetPolicySetType returns the PolicySetType field if non-nil, zero value otherwise.

### GetPolicySetTypeOk

`func (o *PolicySetResponse) GetPolicySetTypeOk() (*SdkGpoPolicySetType, bool)`

GetPolicySetTypeOk returns a tuple with the PolicySetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetType

`func (o *PolicySetResponse) SetPolicySetType(v SdkGpoPolicySetType)`

SetPolicySetType sets PolicySetType field to given value.

### HasPolicySetType

`func (o *PolicySetResponse) HasPolicySetType() bool`

HasPolicySetType returns a boolean if a field has been set.

### GetName

`func (o *PolicySetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicySetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicySetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicySetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PolicySetResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PolicySetResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PolicySetResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicySetResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicySetResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicySetResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicySetResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicySetResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyCount

`func (o *PolicySetResponse) GetPolicyCount() int32`

GetPolicyCount returns the PolicyCount field if non-nil, zero value otherwise.

### GetPolicyCountOk

`func (o *PolicySetResponse) GetPolicyCountOk() (*int32, bool)`

GetPolicyCountOk returns a tuple with the PolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCount

`func (o *PolicySetResponse) SetPolicyCount(v int32)`

SetPolicyCount sets PolicyCount field to given value.

### HasPolicyCount

`func (o *PolicySetResponse) HasPolicyCount() bool`

HasPolicyCount returns a boolean if a field has been set.

### GetPolicies

`func (o *PolicySetResponse) GetPolicies() []PolicyResponse`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *PolicySetResponse) GetPoliciesOk() (*[]PolicyResponse, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *PolicySetResponse) SetPolicies(v []PolicyResponse)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *PolicySetResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *PolicySetResponse) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *PolicySetResponse) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil
### GetScopes

`func (o *PolicySetResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PolicySetResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PolicySetResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PolicySetResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *PolicySetResponse) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *PolicySetResponse) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetIsAssigned

`func (o *PolicySetResponse) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *PolicySetResponse) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *PolicySetResponse) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *PolicySetResponse) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetLastError

`func (o *PolicySetResponse) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *PolicySetResponse) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *PolicySetResponse) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *PolicySetResponse) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *PolicySetResponse) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *PolicySetResponse) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


