# CopyPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyGuids** | Pointer to **[]string** | GUIDs of some of the policies of a policy set. | [optional] 
**ToPolicySets** | Pointer to **[]string** | GUIDs of the destination policy sets. | [optional] 

## Methods

### NewCopyPoliciesRequest

`func NewCopyPoliciesRequest() *CopyPoliciesRequest`

NewCopyPoliciesRequest instantiates a new CopyPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyPoliciesRequestWithDefaults

`func NewCopyPoliciesRequestWithDefaults() *CopyPoliciesRequest`

NewCopyPoliciesRequestWithDefaults instantiates a new CopyPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyGuids

`func (o *CopyPoliciesRequest) GetPolicyGuids() []string`

GetPolicyGuids returns the PolicyGuids field if non-nil, zero value otherwise.

### GetPolicyGuidsOk

`func (o *CopyPoliciesRequest) GetPolicyGuidsOk() (*[]string, bool)`

GetPolicyGuidsOk returns a tuple with the PolicyGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGuids

`func (o *CopyPoliciesRequest) SetPolicyGuids(v []string)`

SetPolicyGuids sets PolicyGuids field to given value.

### HasPolicyGuids

`func (o *CopyPoliciesRequest) HasPolicyGuids() bool`

HasPolicyGuids returns a boolean if a field has been set.

### SetPolicyGuidsNil

`func (o *CopyPoliciesRequest) SetPolicyGuidsNil(b bool)`

 SetPolicyGuidsNil sets the value for PolicyGuids to be an explicit nil

### UnsetPolicyGuids
`func (o *CopyPoliciesRequest) UnsetPolicyGuids()`

UnsetPolicyGuids ensures that no value is present for PolicyGuids, not even an explicit nil
### GetToPolicySets

`func (o *CopyPoliciesRequest) GetToPolicySets() []string`

GetToPolicySets returns the ToPolicySets field if non-nil, zero value otherwise.

### GetToPolicySetsOk

`func (o *CopyPoliciesRequest) GetToPolicySetsOk() (*[]string, bool)`

GetToPolicySetsOk returns a tuple with the ToPolicySets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPolicySets

`func (o *CopyPoliciesRequest) SetToPolicySets(v []string)`

SetToPolicySets sets ToPolicySets field to given value.

### HasToPolicySets

`func (o *CopyPoliciesRequest) HasToPolicySets() bool`

HasToPolicySets returns a boolean if a field has been set.

### SetToPolicySetsNil

`func (o *CopyPoliciesRequest) SetToPolicySetsNil(b bool)`

 SetToPolicySetsNil sets the value for ToPolicySets to be an explicit nil

### UnsetToPolicySets
`func (o *CopyPoliciesRequest) UnsetToPolicySets()`

UnsetToPolicySets ensures that no value is present for ToPolicySets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


