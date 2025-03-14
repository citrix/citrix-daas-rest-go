# AppliedPolicy2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsApplied** | Pointer to **bool** | Is policy applied. | [optional] 
**PolicyName** | Pointer to **NullableString** | Policy name. | [optional] 
**GpoName** | Pointer to **NullableString** | GPO name. | [optional] 
**Reasons** | Pointer to **map[string][]string** | Reasons why the policy is or is not applied. | [optional] 
**Details** | Pointer to [**map[string][]FilterDetail**](array.md) | Reasons why the policy is or is not applied. | [optional] 

## Methods

### NewAppliedPolicy2

`func NewAppliedPolicy2() *AppliedPolicy2`

NewAppliedPolicy2 instantiates a new AppliedPolicy2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedPolicy2WithDefaults

`func NewAppliedPolicy2WithDefaults() *AppliedPolicy2`

NewAppliedPolicy2WithDefaults instantiates a new AppliedPolicy2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsApplied

`func (o *AppliedPolicy2) GetIsApplied() bool`

GetIsApplied returns the IsApplied field if non-nil, zero value otherwise.

### GetIsAppliedOk

`func (o *AppliedPolicy2) GetIsAppliedOk() (*bool, bool)`

GetIsAppliedOk returns a tuple with the IsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApplied

`func (o *AppliedPolicy2) SetIsApplied(v bool)`

SetIsApplied sets IsApplied field to given value.

### HasIsApplied

`func (o *AppliedPolicy2) HasIsApplied() bool`

HasIsApplied returns a boolean if a field has been set.

### GetPolicyName

`func (o *AppliedPolicy2) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AppliedPolicy2) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AppliedPolicy2) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *AppliedPolicy2) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *AppliedPolicy2) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *AppliedPolicy2) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetGpoName

`func (o *AppliedPolicy2) GetGpoName() string`

GetGpoName returns the GpoName field if non-nil, zero value otherwise.

### GetGpoNameOk

`func (o *AppliedPolicy2) GetGpoNameOk() (*string, bool)`

GetGpoNameOk returns a tuple with the GpoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpoName

`func (o *AppliedPolicy2) SetGpoName(v string)`

SetGpoName sets GpoName field to given value.

### HasGpoName

`func (o *AppliedPolicy2) HasGpoName() bool`

HasGpoName returns a boolean if a field has been set.

### SetGpoNameNil

`func (o *AppliedPolicy2) SetGpoNameNil(b bool)`

 SetGpoNameNil sets the value for GpoName to be an explicit nil

### UnsetGpoName
`func (o *AppliedPolicy2) UnsetGpoName()`

UnsetGpoName ensures that no value is present for GpoName, not even an explicit nil
### GetReasons

`func (o *AppliedPolicy2) GetReasons() map[string][]string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *AppliedPolicy2) GetReasonsOk() (*map[string][]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *AppliedPolicy2) SetReasons(v map[string][]string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *AppliedPolicy2) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### SetReasonsNil

`func (o *AppliedPolicy2) SetReasonsNil(b bool)`

 SetReasonsNil sets the value for Reasons to be an explicit nil

### UnsetReasons
`func (o *AppliedPolicy2) UnsetReasons()`

UnsetReasons ensures that no value is present for Reasons, not even an explicit nil
### GetDetails

`func (o *AppliedPolicy2) GetDetails() map[string][]FilterDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AppliedPolicy2) GetDetailsOk() (*map[string][]FilterDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AppliedPolicy2) SetDetails(v map[string][]FilterDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AppliedPolicy2) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *AppliedPolicy2) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *AppliedPolicy2) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


