# AllResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]PolicyResponseContract**](PolicyResponseContract.md) | Policies | [optional] 
**Templates** | Pointer to [**[]TemplateResponseContract**](TemplateResponseContract.md) | Templates | [optional] 
**SettingDefinitions** | Pointer to [**[]SettingDefinitionContract**](SettingDefinitionContract.md) | Setting Definitions | [optional] 
**FilterDefinitions** | Pointer to [**[]FilterDefinitionContract**](FilterDefinitionContract.md) | Filter Definitions | [optional] 

## Methods

### NewAllResponseContract

`func NewAllResponseContract() *AllResponseContract`

NewAllResponseContract instantiates a new AllResponseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllResponseContractWithDefaults

`func NewAllResponseContractWithDefaults() *AllResponseContract`

NewAllResponseContractWithDefaults instantiates a new AllResponseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *AllResponseContract) GetPolicies() []PolicyResponseContract`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AllResponseContract) GetPoliciesOk() (*[]PolicyResponseContract, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AllResponseContract) SetPolicies(v []PolicyResponseContract)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AllResponseContract) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *AllResponseContract) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *AllResponseContract) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil
### GetTemplates

`func (o *AllResponseContract) GetTemplates() []TemplateResponseContract`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *AllResponseContract) GetTemplatesOk() (*[]TemplateResponseContract, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *AllResponseContract) SetTemplates(v []TemplateResponseContract)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *AllResponseContract) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### SetTemplatesNil

`func (o *AllResponseContract) SetTemplatesNil(b bool)`

 SetTemplatesNil sets the value for Templates to be an explicit nil

### UnsetTemplates
`func (o *AllResponseContract) UnsetTemplates()`

UnsetTemplates ensures that no value is present for Templates, not even an explicit nil
### GetSettingDefinitions

`func (o *AllResponseContract) GetSettingDefinitions() []SettingDefinitionContract`

GetSettingDefinitions returns the SettingDefinitions field if non-nil, zero value otherwise.

### GetSettingDefinitionsOk

`func (o *AllResponseContract) GetSettingDefinitionsOk() (*[]SettingDefinitionContract, bool)`

GetSettingDefinitionsOk returns a tuple with the SettingDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingDefinitions

`func (o *AllResponseContract) SetSettingDefinitions(v []SettingDefinitionContract)`

SetSettingDefinitions sets SettingDefinitions field to given value.

### HasSettingDefinitions

`func (o *AllResponseContract) HasSettingDefinitions() bool`

HasSettingDefinitions returns a boolean if a field has been set.

### SetSettingDefinitionsNil

`func (o *AllResponseContract) SetSettingDefinitionsNil(b bool)`

 SetSettingDefinitionsNil sets the value for SettingDefinitions to be an explicit nil

### UnsetSettingDefinitions
`func (o *AllResponseContract) UnsetSettingDefinitions()`

UnsetSettingDefinitions ensures that no value is present for SettingDefinitions, not even an explicit nil
### GetFilterDefinitions

`func (o *AllResponseContract) GetFilterDefinitions() []FilterDefinitionContract`

GetFilterDefinitions returns the FilterDefinitions field if non-nil, zero value otherwise.

### GetFilterDefinitionsOk

`func (o *AllResponseContract) GetFilterDefinitionsOk() (*[]FilterDefinitionContract, bool)`

GetFilterDefinitionsOk returns a tuple with the FilterDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDefinitions

`func (o *AllResponseContract) SetFilterDefinitions(v []FilterDefinitionContract)`

SetFilterDefinitions sets FilterDefinitions field to given value.

### HasFilterDefinitions

`func (o *AllResponseContract) HasFilterDefinitions() bool`

HasFilterDefinitions returns a boolean if a field has been set.

### SetFilterDefinitionsNil

`func (o *AllResponseContract) SetFilterDefinitionsNil(b bool)`

 SetFilterDefinitionsNil sets the value for FilterDefinitions to be an explicit nil

### UnsetFilterDefinitions
`func (o *AllResponseContract) UnsetFilterDefinitions()`

UnsetFilterDefinitions ensures that no value is present for FilterDefinitions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


