# SimulationResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUserRsop** | Pointer to **bool** | The report is for machine or user. | [optional] 
**FilterEvidence** | Pointer to **map[string]string** | Filter evidence used for the simulation. | [optional] 
**AppliedSettings** | Pointer to [**[]AppliedSetting**](AppliedSetting.md) | Applied settings. | [optional] 
**AppliedPolicies** | Pointer to [**[]AppliedPolicy**](AppliedPolicy.md) | Applied policies. | [optional] 
**LosingSettings** | Pointer to [**[]LosingSetting**](LosingSetting.md) | Settings that did not get applied. | [optional] 
**LosingPolicies** | Pointer to [**[]LosingPolicy**](LosingPolicy.md) | Policies that did not get applied. Some of them should be applied in theory but not applied in practice. | [optional] 

## Methods

### NewSimulationResponseContract

`func NewSimulationResponseContract() *SimulationResponseContract`

NewSimulationResponseContract instantiates a new SimulationResponseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationResponseContractWithDefaults

`func NewSimulationResponseContractWithDefaults() *SimulationResponseContract`

NewSimulationResponseContractWithDefaults instantiates a new SimulationResponseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUserRsop

`func (o *SimulationResponseContract) GetIsUserRsop() bool`

GetIsUserRsop returns the IsUserRsop field if non-nil, zero value otherwise.

### GetIsUserRsopOk

`func (o *SimulationResponseContract) GetIsUserRsopOk() (*bool, bool)`

GetIsUserRsopOk returns a tuple with the IsUserRsop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserRsop

`func (o *SimulationResponseContract) SetIsUserRsop(v bool)`

SetIsUserRsop sets IsUserRsop field to given value.

### HasIsUserRsop

`func (o *SimulationResponseContract) HasIsUserRsop() bool`

HasIsUserRsop returns a boolean if a field has been set.

### GetFilterEvidence

`func (o *SimulationResponseContract) GetFilterEvidence() map[string]string`

GetFilterEvidence returns the FilterEvidence field if non-nil, zero value otherwise.

### GetFilterEvidenceOk

`func (o *SimulationResponseContract) GetFilterEvidenceOk() (*map[string]string, bool)`

GetFilterEvidenceOk returns a tuple with the FilterEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterEvidence

`func (o *SimulationResponseContract) SetFilterEvidence(v map[string]string)`

SetFilterEvidence sets FilterEvidence field to given value.

### HasFilterEvidence

`func (o *SimulationResponseContract) HasFilterEvidence() bool`

HasFilterEvidence returns a boolean if a field has been set.

### SetFilterEvidenceNil

`func (o *SimulationResponseContract) SetFilterEvidenceNil(b bool)`

 SetFilterEvidenceNil sets the value for FilterEvidence to be an explicit nil

### UnsetFilterEvidence
`func (o *SimulationResponseContract) UnsetFilterEvidence()`

UnsetFilterEvidence ensures that no value is present for FilterEvidence, not even an explicit nil
### GetAppliedSettings

`func (o *SimulationResponseContract) GetAppliedSettings() []AppliedSetting`

GetAppliedSettings returns the AppliedSettings field if non-nil, zero value otherwise.

### GetAppliedSettingsOk

`func (o *SimulationResponseContract) GetAppliedSettingsOk() (*[]AppliedSetting, bool)`

GetAppliedSettingsOk returns a tuple with the AppliedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedSettings

`func (o *SimulationResponseContract) SetAppliedSettings(v []AppliedSetting)`

SetAppliedSettings sets AppliedSettings field to given value.

### HasAppliedSettings

`func (o *SimulationResponseContract) HasAppliedSettings() bool`

HasAppliedSettings returns a boolean if a field has been set.

### SetAppliedSettingsNil

`func (o *SimulationResponseContract) SetAppliedSettingsNil(b bool)`

 SetAppliedSettingsNil sets the value for AppliedSettings to be an explicit nil

### UnsetAppliedSettings
`func (o *SimulationResponseContract) UnsetAppliedSettings()`

UnsetAppliedSettings ensures that no value is present for AppliedSettings, not even an explicit nil
### GetAppliedPolicies

`func (o *SimulationResponseContract) GetAppliedPolicies() []AppliedPolicy`

GetAppliedPolicies returns the AppliedPolicies field if non-nil, zero value otherwise.

### GetAppliedPoliciesOk

`func (o *SimulationResponseContract) GetAppliedPoliciesOk() (*[]AppliedPolicy, bool)`

GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedPolicies

`func (o *SimulationResponseContract) SetAppliedPolicies(v []AppliedPolicy)`

SetAppliedPolicies sets AppliedPolicies field to given value.

### HasAppliedPolicies

`func (o *SimulationResponseContract) HasAppliedPolicies() bool`

HasAppliedPolicies returns a boolean if a field has been set.

### SetAppliedPoliciesNil

`func (o *SimulationResponseContract) SetAppliedPoliciesNil(b bool)`

 SetAppliedPoliciesNil sets the value for AppliedPolicies to be an explicit nil

### UnsetAppliedPolicies
`func (o *SimulationResponseContract) UnsetAppliedPolicies()`

UnsetAppliedPolicies ensures that no value is present for AppliedPolicies, not even an explicit nil
### GetLosingSettings

`func (o *SimulationResponseContract) GetLosingSettings() []LosingSetting`

GetLosingSettings returns the LosingSettings field if non-nil, zero value otherwise.

### GetLosingSettingsOk

`func (o *SimulationResponseContract) GetLosingSettingsOk() (*[]LosingSetting, bool)`

GetLosingSettingsOk returns a tuple with the LosingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLosingSettings

`func (o *SimulationResponseContract) SetLosingSettings(v []LosingSetting)`

SetLosingSettings sets LosingSettings field to given value.

### HasLosingSettings

`func (o *SimulationResponseContract) HasLosingSettings() bool`

HasLosingSettings returns a boolean if a field has been set.

### SetLosingSettingsNil

`func (o *SimulationResponseContract) SetLosingSettingsNil(b bool)`

 SetLosingSettingsNil sets the value for LosingSettings to be an explicit nil

### UnsetLosingSettings
`func (o *SimulationResponseContract) UnsetLosingSettings()`

UnsetLosingSettings ensures that no value is present for LosingSettings, not even an explicit nil
### GetLosingPolicies

`func (o *SimulationResponseContract) GetLosingPolicies() []LosingPolicy`

GetLosingPolicies returns the LosingPolicies field if non-nil, zero value otherwise.

### GetLosingPoliciesOk

`func (o *SimulationResponseContract) GetLosingPoliciesOk() (*[]LosingPolicy, bool)`

GetLosingPoliciesOk returns a tuple with the LosingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLosingPolicies

`func (o *SimulationResponseContract) SetLosingPolicies(v []LosingPolicy)`

SetLosingPolicies sets LosingPolicies field to given value.

### HasLosingPolicies

`func (o *SimulationResponseContract) HasLosingPolicies() bool`

HasLosingPolicies returns a boolean if a field has been set.

### SetLosingPoliciesNil

`func (o *SimulationResponseContract) SetLosingPoliciesNil(b bool)`

 SetLosingPoliciesNil sets the value for LosingPolicies to be an explicit nil

### UnsetLosingPolicies
`func (o *SimulationResponseContract) UnsetLosingPolicies()`

UnsetLosingPolicies ensures that no value is present for LosingPolicies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


