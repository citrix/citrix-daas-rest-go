# ModelingResultContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMachineRsop** | Pointer to **bool** | The report is for machine or user | [optional] 
**LastProcessTime** | Pointer to **time.Time** | The most recent simulation time | [optional] 
**AppliedSettings** | Pointer to [**[]AppliedSetting2**](AppliedSetting2.md) | Applied settings | [optional] 
**AppliedPolicies** | Pointer to [**[]AppliedPolicy2**](AppliedPolicy2.md) | Applied policies | [optional] 
**FilterEvidence** | Pointer to **map[string]string** | Filter evidence used for the simulation | [optional] 
**LosingSettings** | Pointer to [**[]LosingSetting2**](LosingSetting2.md) | Settings that did not get applied. | [optional] 
**LosingPolicies** | Pointer to [**[]LosingPolicy2**](LosingPolicy2.md) | Policies that did not get applied. Some of them should be applied in theory but not applied in practice. | [optional] 

## Methods

### NewModelingResultContract

`func NewModelingResultContract() *ModelingResultContract`

NewModelingResultContract instantiates a new ModelingResultContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelingResultContractWithDefaults

`func NewModelingResultContractWithDefaults() *ModelingResultContract`

NewModelingResultContractWithDefaults instantiates a new ModelingResultContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMachineRsop

`func (o *ModelingResultContract) GetIsMachineRsop() bool`

GetIsMachineRsop returns the IsMachineRsop field if non-nil, zero value otherwise.

### GetIsMachineRsopOk

`func (o *ModelingResultContract) GetIsMachineRsopOk() (*bool, bool)`

GetIsMachineRsopOk returns a tuple with the IsMachineRsop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachineRsop

`func (o *ModelingResultContract) SetIsMachineRsop(v bool)`

SetIsMachineRsop sets IsMachineRsop field to given value.

### HasIsMachineRsop

`func (o *ModelingResultContract) HasIsMachineRsop() bool`

HasIsMachineRsop returns a boolean if a field has been set.

### GetLastProcessTime

`func (o *ModelingResultContract) GetLastProcessTime() time.Time`

GetLastProcessTime returns the LastProcessTime field if non-nil, zero value otherwise.

### GetLastProcessTimeOk

`func (o *ModelingResultContract) GetLastProcessTimeOk() (*time.Time, bool)`

GetLastProcessTimeOk returns a tuple with the LastProcessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessTime

`func (o *ModelingResultContract) SetLastProcessTime(v time.Time)`

SetLastProcessTime sets LastProcessTime field to given value.

### HasLastProcessTime

`func (o *ModelingResultContract) HasLastProcessTime() bool`

HasLastProcessTime returns a boolean if a field has been set.

### GetAppliedSettings

`func (o *ModelingResultContract) GetAppliedSettings() []AppliedSetting2`

GetAppliedSettings returns the AppliedSettings field if non-nil, zero value otherwise.

### GetAppliedSettingsOk

`func (o *ModelingResultContract) GetAppliedSettingsOk() (*[]AppliedSetting2, bool)`

GetAppliedSettingsOk returns a tuple with the AppliedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedSettings

`func (o *ModelingResultContract) SetAppliedSettings(v []AppliedSetting2)`

SetAppliedSettings sets AppliedSettings field to given value.

### HasAppliedSettings

`func (o *ModelingResultContract) HasAppliedSettings() bool`

HasAppliedSettings returns a boolean if a field has been set.

### SetAppliedSettingsNil

`func (o *ModelingResultContract) SetAppliedSettingsNil(b bool)`

 SetAppliedSettingsNil sets the value for AppliedSettings to be an explicit nil

### UnsetAppliedSettings
`func (o *ModelingResultContract) UnsetAppliedSettings()`

UnsetAppliedSettings ensures that no value is present for AppliedSettings, not even an explicit nil
### GetAppliedPolicies

`func (o *ModelingResultContract) GetAppliedPolicies() []AppliedPolicy2`

GetAppliedPolicies returns the AppliedPolicies field if non-nil, zero value otherwise.

### GetAppliedPoliciesOk

`func (o *ModelingResultContract) GetAppliedPoliciesOk() (*[]AppliedPolicy2, bool)`

GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedPolicies

`func (o *ModelingResultContract) SetAppliedPolicies(v []AppliedPolicy2)`

SetAppliedPolicies sets AppliedPolicies field to given value.

### HasAppliedPolicies

`func (o *ModelingResultContract) HasAppliedPolicies() bool`

HasAppliedPolicies returns a boolean if a field has been set.

### SetAppliedPoliciesNil

`func (o *ModelingResultContract) SetAppliedPoliciesNil(b bool)`

 SetAppliedPoliciesNil sets the value for AppliedPolicies to be an explicit nil

### UnsetAppliedPolicies
`func (o *ModelingResultContract) UnsetAppliedPolicies()`

UnsetAppliedPolicies ensures that no value is present for AppliedPolicies, not even an explicit nil
### GetFilterEvidence

`func (o *ModelingResultContract) GetFilterEvidence() map[string]string`

GetFilterEvidence returns the FilterEvidence field if non-nil, zero value otherwise.

### GetFilterEvidenceOk

`func (o *ModelingResultContract) GetFilterEvidenceOk() (*map[string]string, bool)`

GetFilterEvidenceOk returns a tuple with the FilterEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterEvidence

`func (o *ModelingResultContract) SetFilterEvidence(v map[string]string)`

SetFilterEvidence sets FilterEvidence field to given value.

### HasFilterEvidence

`func (o *ModelingResultContract) HasFilterEvidence() bool`

HasFilterEvidence returns a boolean if a field has been set.

### SetFilterEvidenceNil

`func (o *ModelingResultContract) SetFilterEvidenceNil(b bool)`

 SetFilterEvidenceNil sets the value for FilterEvidence to be an explicit nil

### UnsetFilterEvidence
`func (o *ModelingResultContract) UnsetFilterEvidence()`

UnsetFilterEvidence ensures that no value is present for FilterEvidence, not even an explicit nil
### GetLosingSettings

`func (o *ModelingResultContract) GetLosingSettings() []LosingSetting2`

GetLosingSettings returns the LosingSettings field if non-nil, zero value otherwise.

### GetLosingSettingsOk

`func (o *ModelingResultContract) GetLosingSettingsOk() (*[]LosingSetting2, bool)`

GetLosingSettingsOk returns a tuple with the LosingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLosingSettings

`func (o *ModelingResultContract) SetLosingSettings(v []LosingSetting2)`

SetLosingSettings sets LosingSettings field to given value.

### HasLosingSettings

`func (o *ModelingResultContract) HasLosingSettings() bool`

HasLosingSettings returns a boolean if a field has been set.

### SetLosingSettingsNil

`func (o *ModelingResultContract) SetLosingSettingsNil(b bool)`

 SetLosingSettingsNil sets the value for LosingSettings to be an explicit nil

### UnsetLosingSettings
`func (o *ModelingResultContract) UnsetLosingSettings()`

UnsetLosingSettings ensures that no value is present for LosingSettings, not even an explicit nil
### GetLosingPolicies

`func (o *ModelingResultContract) GetLosingPolicies() []LosingPolicy2`

GetLosingPolicies returns the LosingPolicies field if non-nil, zero value otherwise.

### GetLosingPoliciesOk

`func (o *ModelingResultContract) GetLosingPoliciesOk() (*[]LosingPolicy2, bool)`

GetLosingPoliciesOk returns a tuple with the LosingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLosingPolicies

`func (o *ModelingResultContract) SetLosingPolicies(v []LosingPolicy2)`

SetLosingPolicies sets LosingPolicies field to given value.

### HasLosingPolicies

`func (o *ModelingResultContract) HasLosingPolicies() bool`

HasLosingPolicies returns a boolean if a field has been set.

### SetLosingPoliciesNil

`func (o *ModelingResultContract) SetLosingPoliciesNil(b bool)`

 SetLosingPoliciesNil sets the value for LosingPolicies to be an explicit nil

### UnsetLosingPolicies
`func (o *ModelingResultContract) UnsetLosingPolicies()`

UnsetLosingPolicies ensures that no value is present for LosingPolicies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


