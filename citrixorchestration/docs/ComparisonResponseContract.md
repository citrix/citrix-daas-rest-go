# ComparisonResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to **[]string** | A target is either a policy or template. This list stores the target names. | [optional] 
**Result** | Pointer to [**[]ComparisonRowContract**](ComparisonRowContract.md) | Rows in the report. | [optional] 
**WithDefaults** | Pointer to **bool** | If true, first column is setting default values. | [optional] 

## Methods

### NewComparisonResponseContract

`func NewComparisonResponseContract() *ComparisonResponseContract`

NewComparisonResponseContract instantiates a new ComparisonResponseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparisonResponseContractWithDefaults

`func NewComparisonResponseContractWithDefaults() *ComparisonResponseContract`

NewComparisonResponseContractWithDefaults instantiates a new ComparisonResponseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *ComparisonResponseContract) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ComparisonResponseContract) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ComparisonResponseContract) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ComparisonResponseContract) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *ComparisonResponseContract) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *ComparisonResponseContract) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetResult

`func (o *ComparisonResponseContract) GetResult() []ComparisonRowContract`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ComparisonResponseContract) GetResultOk() (*[]ComparisonRowContract, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ComparisonResponseContract) SetResult(v []ComparisonRowContract)`

SetResult sets Result field to given value.

### HasResult

`func (o *ComparisonResponseContract) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ComparisonResponseContract) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ComparisonResponseContract) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetWithDefaults

`func (o *ComparisonResponseContract) GetWithDefaults() bool`

GetWithDefaults returns the WithDefaults field if non-nil, zero value otherwise.

### GetWithDefaultsOk

`func (o *ComparisonResponseContract) GetWithDefaultsOk() (*bool, bool)`

GetWithDefaultsOk returns a tuple with the WithDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithDefaults

`func (o *ComparisonResponseContract) SetWithDefaults(v bool)`

SetWithDefaults sets WithDefaults field to given value.

### HasWithDefaults

`func (o *ComparisonResponseContract) HasWithDefaults() bool`

HasWithDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


