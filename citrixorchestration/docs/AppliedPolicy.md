# AppliedPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyIdentity** | Pointer to [**PolicyIdentity**](PolicyIdentity.md) |  | [optional] 
**Reasons** | Pointer to **map[string][]string** | Reasons why the policy is applied. | [optional] 

## Methods

### NewAppliedPolicy

`func NewAppliedPolicy() *AppliedPolicy`

NewAppliedPolicy instantiates a new AppliedPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedPolicyWithDefaults

`func NewAppliedPolicyWithDefaults() *AppliedPolicy`

NewAppliedPolicyWithDefaults instantiates a new AppliedPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyIdentity

`func (o *AppliedPolicy) GetPolicyIdentity() PolicyIdentity`

GetPolicyIdentity returns the PolicyIdentity field if non-nil, zero value otherwise.

### GetPolicyIdentityOk

`func (o *AppliedPolicy) GetPolicyIdentityOk() (*PolicyIdentity, bool)`

GetPolicyIdentityOk returns a tuple with the PolicyIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentity

`func (o *AppliedPolicy) SetPolicyIdentity(v PolicyIdentity)`

SetPolicyIdentity sets PolicyIdentity field to given value.

### HasPolicyIdentity

`func (o *AppliedPolicy) HasPolicyIdentity() bool`

HasPolicyIdentity returns a boolean if a field has been set.

### GetReasons

`func (o *AppliedPolicy) GetReasons() map[string][]string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *AppliedPolicy) GetReasonsOk() (*map[string][]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *AppliedPolicy) SetReasons(v map[string][]string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *AppliedPolicy) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### SetReasonsNil

`func (o *AppliedPolicy) SetReasonsNil(b bool)`

 SetReasonsNil sets the value for Reasons to be an explicit nil

### UnsetReasons
`func (o *AppliedPolicy) UnsetReasons()`

UnsetReasons ensures that no value is present for Reasons, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


