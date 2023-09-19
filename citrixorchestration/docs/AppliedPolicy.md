# AppliedPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsApplied** | Pointer to **bool** | Is policy applied | [optional] 
**PolicyName** | Pointer to **NullableString** | Policy name | [optional] 
**GpoName** | Pointer to **NullableString** | GPO name | [optional] 
**Reasons** | Pointer to **map[string][]string** | Reasons why the policy is or is not applied | [optional] 
**Details** | Pointer to [**map[string][]FilterDetail**](array.md) | Reasons why the policy is or is not applied. | [optional] 

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

### GetIsApplied

`func (o *AppliedPolicy) GetIsApplied() bool`

GetIsApplied returns the IsApplied field if non-nil, zero value otherwise.

### GetIsAppliedOk

`func (o *AppliedPolicy) GetIsAppliedOk() (*bool, bool)`

GetIsAppliedOk returns a tuple with the IsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApplied

`func (o *AppliedPolicy) SetIsApplied(v bool)`

SetIsApplied sets IsApplied field to given value.

### HasIsApplied

`func (o *AppliedPolicy) HasIsApplied() bool`

HasIsApplied returns a boolean if a field has been set.

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
### GetDetails

`func (o *AppliedPolicy) GetDetails() map[string][]FilterDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AppliedPolicy) GetDetailsOk() (*map[string][]FilterDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AppliedPolicy) SetDetails(v map[string][]FilterDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AppliedPolicy) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *AppliedPolicy) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *AppliedPolicy) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


