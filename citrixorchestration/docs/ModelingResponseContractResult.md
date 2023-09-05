# ModelingResponseContractResult

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

### NewModelingResponseContractResult

`func NewModelingResponseContractResult() *ModelingResponseContractResult`

NewModelingResponseContractResult instantiates a new ModelingResponseContractResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelingResponseContractResultWithDefaults

`func NewModelingResponseContractResultWithDefaults() *ModelingResponseContractResult`

NewModelingResponseContractResultWithDefaults instantiates a new ModelingResponseContractResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMachineRsop

`func (o *ModelingResponseContractResult) GetIsMachineRsop() bool`

GetIsMachineRsop returns the IsMachineRsop field if non-nil, zero value otherwise.

### GetIsMachineRsopOk

`func (o *ModelingResponseContractResult) GetIsMachineRsopOk() (*bool, bool)`

GetIsMachineRsopOk returns a tuple with the IsMachineRsop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachineRsop

`func (o *ModelingResponseContractResult) SetIsMachineRsop(v bool)`

SetIsMachineRsop sets IsMachineRsop field to given value.

### HasIsMachineRsop

`func (o *ModelingResponseContractResult) HasIsMachineRsop() bool`

HasIsMachineRsop returns a boolean if a field has been set.

### GetLastProcessTime

`func (o *ModelingResponseContractResult) GetLastProcessTime() time.Time`

GetLastProcessTime returns the LastProcessTime field if non-nil, zero value otherwise.

### GetLastProcessTimeOk

`func (o *ModelingResponseContractResult) GetLastProcessTimeOk() (*time.Time, bool)`

GetLastProcessTimeOk returns a tuple with the LastProcessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessTime

`func (o *ModelingResponseContractResult) SetLastProcessTime(v time.Time)`

SetLastProcessTime sets LastProcessTime field to given value.

### HasLastProcessTime

`func (o *ModelingResponseContractResult) HasLastProcessTime() bool`

HasLastProcessTime returns a boolean if a field has been set.

### GetAppliedSettings

`func (o *ModelingResponseContractResult) GetAppliedSettings() []AppliedSetting`

GetAppliedSettings returns the AppliedSettings field if non-nil, zero value otherwise.

### GetAppliedSettingsOk

`func (o *ModelingResponseContractResult) GetAppliedSettingsOk() (*[]AppliedSetting, bool)`

GetAppliedSettingsOk returns a tuple with the AppliedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedSettings

`func (o *ModelingResponseContractResult) SetAppliedSettings(v []AppliedSetting)`

SetAppliedSettings sets AppliedSettings field to given value.

### HasAppliedSettings

`func (o *ModelingResponseContractResult) HasAppliedSettings() bool`

HasAppliedSettings returns a boolean if a field has been set.

### GetAppliedPolicies

`func (o *ModelingResponseContractResult) GetAppliedPolicies() []AppliedPolicy`

GetAppliedPolicies returns the AppliedPolicies field if non-nil, zero value otherwise.

### GetAppliedPoliciesOk

`func (o *ModelingResponseContractResult) GetAppliedPoliciesOk() (*[]AppliedPolicy, bool)`

GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedPolicies

`func (o *ModelingResponseContractResult) SetAppliedPolicies(v []AppliedPolicy)`

SetAppliedPolicies sets AppliedPolicies field to given value.

### HasAppliedPolicies

`func (o *ModelingResponseContractResult) HasAppliedPolicies() bool`

HasAppliedPolicies returns a boolean if a field has been set.

### GetFilterEvidence

`func (o *ModelingResponseContractResult) GetFilterEvidence() map[string]string`

GetFilterEvidence returns the FilterEvidence field if non-nil, zero value otherwise.

### GetFilterEvidenceOk

`func (o *ModelingResponseContractResult) GetFilterEvidenceOk() (*map[string]string, bool)`

GetFilterEvidenceOk returns a tuple with the FilterEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterEvidence

`func (o *ModelingResponseContractResult) SetFilterEvidence(v map[string]string)`

SetFilterEvidence sets FilterEvidence field to given value.

### HasFilterEvidence

`func (o *ModelingResponseContractResult) HasFilterEvidence() bool`

HasFilterEvidence returns a boolean if a field has been set.

### GetLosingSettings

`func (o *ModelingResponseContractResult) GetLosingSettings() []LosingSetting`

GetLosingSettings returns the LosingSettings field if non-nil, zero value otherwise.

### GetLosingSettingsOk

`func (o *ModelingResponseContractResult) GetLosingSettingsOk() (*[]LosingSetting, bool)`

GetLosingSettingsOk returns a tuple with the LosingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLosingSettings

`func (o *ModelingResponseContractResult) SetLosingSettings(v []LosingSetting)`

SetLosingSettings sets LosingSettings field to given value.

### HasLosingSettings

`func (o *ModelingResponseContractResult) HasLosingSettings() bool`

HasLosingSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


