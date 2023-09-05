# PolicySetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicySetGuid** | **string** | Guid of the policy set. | 
**PolicySetType** | [**SdkGpoPolicySetType**](SdkGpoPolicySetType.md) |  | 
**Name** | Pointer to **string** | Name of the policy set. | [optional] 
**Description** | Pointer to **string** | Policy set description. | [optional] 
**PolicyCount** | **int32** | Number of policies in the policy set. | 
**Policies** | Pointer to [**[]PolicyResponse**](PolicyResponse.md) | Policies in the policy set. | [optional] 
**Scopes** | Pointer to **[]string** | Delegated admin scopes. | [optional] 

## Methods

### NewPolicySetResponse

`func NewPolicySetResponse(policySetGuid string, policySetType SdkGpoPolicySetType, policyCount int32, ) *PolicySetResponse`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


