# ModelingResultContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMachineRsop** | Pointer to **bool** | The report is for machine or user | [optional] 
**LastProcessTime** | Pointer to **time.Time** | The most recent simulation time | [optional] 
**AppliedSettings** | Pointer to [**[]AppliedSetting**](AppliedSetting.md) | Applied settings | [optional] 
**AppliedPolicies** | Pointer to [**[]AppliedPolicy**](AppliedPolicy.md) | Applied policies | [optional] 
**FilterEvidence** | Pointer to **map[string]string** | Filter evidence used for the simulation | [optional] 
**LosingSettings** | Pointer to [**[]LosingSetting**](LosingSetting.md) | Settings that did not get applied. | [optional] 

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

`func (o *ModelingResultContract) GetAppliedSettings() []AppliedSetting`

GetAppliedSettings returns the AppliedSettings field if non-nil, zero value otherwise.

### GetAppliedSettingsOk

`func (o *ModelingResultContract) GetAppliedSettingsOk() (*[]AppliedSetting, bool)`

GetAppliedSettingsOk returns a tuple with the AppliedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedSettings

`func (o *ModelingResultContract) SetAppliedSettings(v []AppliedSetting)`

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

`func (o *ModelingResultContract) GetAppliedPolicies() []AppliedPolicy`

GetAppliedPolicies returns the AppliedPolicies field if non-nil, zero value otherwise.

### GetAppliedPoliciesOk

`func (o *ModelingResultContract) GetAppliedPoliciesOk() (*[]AppliedPolicy, bool)`

GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedPolicies

`func (o *ModelingResultContract) SetAppliedPolicies(v []AppliedPolicy)`

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

`func (o *ModelingResultContract) GetLosingSettings() []LosingSetting`

GetLosingSettings returns the LosingSettings field if non-nil, zero value otherwise.

### GetLosingSettingsOk

`func (o *ModelingResultContract) GetLosingSettingsOk() (*[]LosingSetting, bool)`

GetLosingSettingsOk returns a tuple with the LosingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLosingSettings

`func (o *ModelingResultContract) SetLosingSettings(v []LosingSetting)`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


