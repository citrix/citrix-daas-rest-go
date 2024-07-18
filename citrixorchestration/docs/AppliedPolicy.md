# AppliedPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | Pointer to **NullableString** | Name of the policy. | [optional] 
**GpoName** | Pointer to **NullableString** | Name of the GPO that contains the policy that uses this setting. | [optional] 
**Reasons** | Pointer to [**map[string][]ReasonDetail**](array.md) | Reasons why the policy is applied. | [optional] 

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

### GetPolicyName

`func (o *AppliedPolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AppliedPolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AppliedPolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *AppliedPolicy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *AppliedPolicy) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *AppliedPolicy) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetGpoName

`func (o *AppliedPolicy) GetGpoName() string`

GetGpoName returns the GpoName field if non-nil, zero value otherwise.

### GetGpoNameOk

`func (o *AppliedPolicy) GetGpoNameOk() (*string, bool)`

GetGpoNameOk returns a tuple with the GpoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpoName

`func (o *AppliedPolicy) SetGpoName(v string)`

SetGpoName sets GpoName field to given value.

### HasGpoName

`func (o *AppliedPolicy) HasGpoName() bool`

HasGpoName returns a boolean if a field has been set.

### SetGpoNameNil

`func (o *AppliedPolicy) SetGpoNameNil(b bool)`

 SetGpoNameNil sets the value for GpoName to be an explicit nil

### UnsetGpoName
`func (o *AppliedPolicy) UnsetGpoName()`

UnsetGpoName ensures that no value is present for GpoName, not even an explicit nil
### GetReasons

`func (o *AppliedPolicy) GetReasons() map[string][]ReasonDetail`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *AppliedPolicy) GetReasonsOk() (*map[string][]ReasonDetail, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *AppliedPolicy) SetReasons(v map[string][]ReasonDetail)`

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


