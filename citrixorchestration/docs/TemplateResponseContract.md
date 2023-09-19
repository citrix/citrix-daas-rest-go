# TemplateResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateName** | Pointer to **NullableString** | Template name as identifier | [optional] 
**IsReadOnly** | Pointer to **bool** | Is template builtin. | [optional] 
**Description** | Pointer to **NullableString** | Template description | [optional] 
**Settings** | Pointer to [**[]SettingResponseContract**](SettingResponseContract.md) | Settings | [optional] 

## Methods

### NewTemplateResponseContract

`func NewTemplateResponseContract() *TemplateResponseContract`

NewTemplateResponseContract instantiates a new TemplateResponseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateResponseContractWithDefaults

`func NewTemplateResponseContractWithDefaults() *TemplateResponseContract`

NewTemplateResponseContractWithDefaults instantiates a new TemplateResponseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateName

`func (o *TemplateResponseContract) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *TemplateResponseContract) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *TemplateResponseContract) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *TemplateResponseContract) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *TemplateResponseContract) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *TemplateResponseContract) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetIsReadOnly

`func (o *TemplateResponseContract) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *TemplateResponseContract) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *TemplateResponseContract) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *TemplateResponseContract) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateResponseContract) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateResponseContract) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateResponseContract) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateResponseContract) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateResponseContract) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateResponseContract) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSettings

`func (o *TemplateResponseContract) GetSettings() []SettingResponseContract`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *TemplateResponseContract) GetSettingsOk() (*[]SettingResponseContract, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *TemplateResponseContract) SetSettings(v []SettingResponseContract)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *TemplateResponseContract) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *TemplateResponseContract) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *TemplateResponseContract) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


