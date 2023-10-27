# SettingDefinitionContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditorName** | Pointer to **NullableString** | The setting editor name. | [optional] 
**SettingName** | Pointer to **NullableString** | Setting name; Is globally unique | [optional] 
**SettingType** | Pointer to **NullableString** | The setting type as defined by its value. | [optional] 
**DisplayName** | Pointer to **NullableString** | Translated setting name. Is globally unique. | [optional] 
**Category** | Pointer to **NullableString** | Setting category, e.g., ICA\\Printing. | [optional] 
**Explanation** | Pointer to **NullableString** | Setting detailed description | [optional] 
**IsUserSetting** | Pointer to **bool** | true &#x3D; user, false &#x3D; machine | [optional] 
**ValueType** | Pointer to **NullableString** | Type of setting value | [optional] 
**IsEnableDisable** | Pointer to **bool** | Label test for a boolean setting, display enable/disable if true, otherwise allow/prohibit | [optional] 
**IsToggled** | Pointer to **bool** | Indicate if the setting is toggled. | [optional] 
**HideValueHint** | Pointer to **bool** | Indicate if the detailed validation message about the value should be shown. | [optional] 
**DefaultValue** | Pointer to **NullableString** | Setting default value | [optional] 
**DisabledValue** | Pointer to **NullableString** | Disabled value for toggled settings | [optional] 
**InitialValue** | Pointer to **NullableString** | Initial value for toggled settings when value is not disabled. | [optional] 
**ValueMinimum** | Pointer to **NullableString** | Minimal value for integer setting | [optional] 
**ValueMaximum** | Pointer to **NullableString** | Maximal value for integer setting | [optional] 
**ValueValidator** | Pointer to **NullableString** | An expression executed to validate the value | [optional] 
**ValidatorError** | Pointer to **NullableString** | Warning message issued when validator fails (throws exception) | [optional] 
**RelatedSettings** | Pointer to **map[string]string** | Related settings. | [optional] 
**ValueUnit** | Pointer to **NullableString** | Unit of value in translated text | [optional] 
**AllowedText** | Pointer to **NullableString** | Custom text for the allowed explanation for some boolean settings. | [optional] 
**DeniedText** | Pointer to **NullableString** | Custom text for the denied explanation for some boolean settings. | [optional] 
**EnumType** | Pointer to [**EnumerationTypeContract**](EnumerationTypeContract.md) |  | [optional] 
**VdaVersions** | Pointer to [**[]VdaVersionContract**](VdaVersionContract.md) | VDA versions honoring this setting. | [optional] 
**GpoScope** | Pointer to **NullableString** | Scope, e.g. ConfigSlot | [optional] 
**ProductGroup** | Pointer to **NullableString** | Group, e.g., UPM. | [optional] 

## Methods

### NewSettingDefinitionContract

`func NewSettingDefinitionContract() *SettingDefinitionContract`

NewSettingDefinitionContract instantiates a new SettingDefinitionContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingDefinitionContractWithDefaults

`func NewSettingDefinitionContractWithDefaults() *SettingDefinitionContract`

NewSettingDefinitionContractWithDefaults instantiates a new SettingDefinitionContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditorName

`func (o *SettingDefinitionContract) GetEditorName() string`

GetEditorName returns the EditorName field if non-nil, zero value otherwise.

### GetEditorNameOk

`func (o *SettingDefinitionContract) GetEditorNameOk() (*string, bool)`

GetEditorNameOk returns a tuple with the EditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorName

`func (o *SettingDefinitionContract) SetEditorName(v string)`

SetEditorName sets EditorName field to given value.

### HasEditorName

`func (o *SettingDefinitionContract) HasEditorName() bool`

HasEditorName returns a boolean if a field has been set.

### SetEditorNameNil

`func (o *SettingDefinitionContract) SetEditorNameNil(b bool)`

 SetEditorNameNil sets the value for EditorName to be an explicit nil

### UnsetEditorName
`func (o *SettingDefinitionContract) UnsetEditorName()`

UnsetEditorName ensures that no value is present for EditorName, not even an explicit nil
### GetSettingName

`func (o *SettingDefinitionContract) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingDefinitionContract) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingDefinitionContract) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *SettingDefinitionContract) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *SettingDefinitionContract) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *SettingDefinitionContract) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetSettingType

`func (o *SettingDefinitionContract) GetSettingType() string`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *SettingDefinitionContract) GetSettingTypeOk() (*string, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *SettingDefinitionContract) SetSettingType(v string)`

SetSettingType sets SettingType field to given value.

### HasSettingType

`func (o *SettingDefinitionContract) HasSettingType() bool`

HasSettingType returns a boolean if a field has been set.

### SetSettingTypeNil

`func (o *SettingDefinitionContract) SetSettingTypeNil(b bool)`

 SetSettingTypeNil sets the value for SettingType to be an explicit nil

### UnsetSettingType
`func (o *SettingDefinitionContract) UnsetSettingType()`

UnsetSettingType ensures that no value is present for SettingType, not even an explicit nil
### GetDisplayName

`func (o *SettingDefinitionContract) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SettingDefinitionContract) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SettingDefinitionContract) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SettingDefinitionContract) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SettingDefinitionContract) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SettingDefinitionContract) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetCategory

`func (o *SettingDefinitionContract) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SettingDefinitionContract) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SettingDefinitionContract) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SettingDefinitionContract) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *SettingDefinitionContract) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *SettingDefinitionContract) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetExplanation

`func (o *SettingDefinitionContract) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *SettingDefinitionContract) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *SettingDefinitionContract) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *SettingDefinitionContract) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### SetExplanationNil

`func (o *SettingDefinitionContract) SetExplanationNil(b bool)`

 SetExplanationNil sets the value for Explanation to be an explicit nil

### UnsetExplanation
`func (o *SettingDefinitionContract) UnsetExplanation()`

UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil
### GetIsUserSetting

`func (o *SettingDefinitionContract) GetIsUserSetting() bool`

GetIsUserSetting returns the IsUserSetting field if non-nil, zero value otherwise.

### GetIsUserSettingOk

`func (o *SettingDefinitionContract) GetIsUserSettingOk() (*bool, bool)`

GetIsUserSettingOk returns a tuple with the IsUserSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserSetting

`func (o *SettingDefinitionContract) SetIsUserSetting(v bool)`

SetIsUserSetting sets IsUserSetting field to given value.

### HasIsUserSetting

`func (o *SettingDefinitionContract) HasIsUserSetting() bool`

HasIsUserSetting returns a boolean if a field has been set.

### GetValueType

`func (o *SettingDefinitionContract) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *SettingDefinitionContract) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *SettingDefinitionContract) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *SettingDefinitionContract) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueTypeNil

`func (o *SettingDefinitionContract) SetValueTypeNil(b bool)`

 SetValueTypeNil sets the value for ValueType to be an explicit nil

### UnsetValueType
`func (o *SettingDefinitionContract) UnsetValueType()`

UnsetValueType ensures that no value is present for ValueType, not even an explicit nil
### GetIsEnableDisable

`func (o *SettingDefinitionContract) GetIsEnableDisable() bool`

GetIsEnableDisable returns the IsEnableDisable field if non-nil, zero value otherwise.

### GetIsEnableDisableOk

`func (o *SettingDefinitionContract) GetIsEnableDisableOk() (*bool, bool)`

GetIsEnableDisableOk returns a tuple with the IsEnableDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableDisable

`func (o *SettingDefinitionContract) SetIsEnableDisable(v bool)`

SetIsEnableDisable sets IsEnableDisable field to given value.

### HasIsEnableDisable

`func (o *SettingDefinitionContract) HasIsEnableDisable() bool`

HasIsEnableDisable returns a boolean if a field has been set.

### GetIsToggled

`func (o *SettingDefinitionContract) GetIsToggled() bool`

GetIsToggled returns the IsToggled field if non-nil, zero value otherwise.

### GetIsToggledOk

`func (o *SettingDefinitionContract) GetIsToggledOk() (*bool, bool)`

GetIsToggledOk returns a tuple with the IsToggled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsToggled

`func (o *SettingDefinitionContract) SetIsToggled(v bool)`

SetIsToggled sets IsToggled field to given value.

### HasIsToggled

`func (o *SettingDefinitionContract) HasIsToggled() bool`

HasIsToggled returns a boolean if a field has been set.

### GetHideValueHint

`func (o *SettingDefinitionContract) GetHideValueHint() bool`

GetHideValueHint returns the HideValueHint field if non-nil, zero value otherwise.

### GetHideValueHintOk

`func (o *SettingDefinitionContract) GetHideValueHintOk() (*bool, bool)`

GetHideValueHintOk returns a tuple with the HideValueHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideValueHint

`func (o *SettingDefinitionContract) SetHideValueHint(v bool)`

SetHideValueHint sets HideValueHint field to given value.

### HasHideValueHint

`func (o *SettingDefinitionContract) HasHideValueHint() bool`

HasHideValueHint returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SettingDefinitionContract) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SettingDefinitionContract) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SettingDefinitionContract) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SettingDefinitionContract) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *SettingDefinitionContract) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *SettingDefinitionContract) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDisabledValue

`func (o *SettingDefinitionContract) GetDisabledValue() string`

GetDisabledValue returns the DisabledValue field if non-nil, zero value otherwise.

### GetDisabledValueOk

`func (o *SettingDefinitionContract) GetDisabledValueOk() (*string, bool)`

GetDisabledValueOk returns a tuple with the DisabledValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledValue

`func (o *SettingDefinitionContract) SetDisabledValue(v string)`

SetDisabledValue sets DisabledValue field to given value.

### HasDisabledValue

`func (o *SettingDefinitionContract) HasDisabledValue() bool`

HasDisabledValue returns a boolean if a field has been set.

### SetDisabledValueNil

`func (o *SettingDefinitionContract) SetDisabledValueNil(b bool)`

 SetDisabledValueNil sets the value for DisabledValue to be an explicit nil

### UnsetDisabledValue
`func (o *SettingDefinitionContract) UnsetDisabledValue()`

UnsetDisabledValue ensures that no value is present for DisabledValue, not even an explicit nil
### GetInitialValue

`func (o *SettingDefinitionContract) GetInitialValue() string`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *SettingDefinitionContract) GetInitialValueOk() (*string, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *SettingDefinitionContract) SetInitialValue(v string)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *SettingDefinitionContract) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### SetInitialValueNil

`func (o *SettingDefinitionContract) SetInitialValueNil(b bool)`

 SetInitialValueNil sets the value for InitialValue to be an explicit nil

### UnsetInitialValue
`func (o *SettingDefinitionContract) UnsetInitialValue()`

UnsetInitialValue ensures that no value is present for InitialValue, not even an explicit nil
### GetValueMinimum

`func (o *SettingDefinitionContract) GetValueMinimum() string`

GetValueMinimum returns the ValueMinimum field if non-nil, zero value otherwise.

### GetValueMinimumOk

`func (o *SettingDefinitionContract) GetValueMinimumOk() (*string, bool)`

GetValueMinimumOk returns a tuple with the ValueMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMinimum

`func (o *SettingDefinitionContract) SetValueMinimum(v string)`

SetValueMinimum sets ValueMinimum field to given value.

### HasValueMinimum

`func (o *SettingDefinitionContract) HasValueMinimum() bool`

HasValueMinimum returns a boolean if a field has been set.

### SetValueMinimumNil

`func (o *SettingDefinitionContract) SetValueMinimumNil(b bool)`

 SetValueMinimumNil sets the value for ValueMinimum to be an explicit nil

### UnsetValueMinimum
`func (o *SettingDefinitionContract) UnsetValueMinimum()`

UnsetValueMinimum ensures that no value is present for ValueMinimum, not even an explicit nil
### GetValueMaximum

`func (o *SettingDefinitionContract) GetValueMaximum() string`

GetValueMaximum returns the ValueMaximum field if non-nil, zero value otherwise.

### GetValueMaximumOk

`func (o *SettingDefinitionContract) GetValueMaximumOk() (*string, bool)`

GetValueMaximumOk returns a tuple with the ValueMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMaximum

`func (o *SettingDefinitionContract) SetValueMaximum(v string)`

SetValueMaximum sets ValueMaximum field to given value.

### HasValueMaximum

`func (o *SettingDefinitionContract) HasValueMaximum() bool`

HasValueMaximum returns a boolean if a field has been set.

### SetValueMaximumNil

`func (o *SettingDefinitionContract) SetValueMaximumNil(b bool)`

 SetValueMaximumNil sets the value for ValueMaximum to be an explicit nil

### UnsetValueMaximum
`func (o *SettingDefinitionContract) UnsetValueMaximum()`

UnsetValueMaximum ensures that no value is present for ValueMaximum, not even an explicit nil
### GetValueValidator

`func (o *SettingDefinitionContract) GetValueValidator() string`

GetValueValidator returns the ValueValidator field if non-nil, zero value otherwise.

### GetValueValidatorOk

`func (o *SettingDefinitionContract) GetValueValidatorOk() (*string, bool)`

GetValueValidatorOk returns a tuple with the ValueValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueValidator

`func (o *SettingDefinitionContract) SetValueValidator(v string)`

SetValueValidator sets ValueValidator field to given value.

### HasValueValidator

`func (o *SettingDefinitionContract) HasValueValidator() bool`

HasValueValidator returns a boolean if a field has been set.

### SetValueValidatorNil

`func (o *SettingDefinitionContract) SetValueValidatorNil(b bool)`

 SetValueValidatorNil sets the value for ValueValidator to be an explicit nil

### UnsetValueValidator
`func (o *SettingDefinitionContract) UnsetValueValidator()`

UnsetValueValidator ensures that no value is present for ValueValidator, not even an explicit nil
### GetValidatorError

`func (o *SettingDefinitionContract) GetValidatorError() string`

GetValidatorError returns the ValidatorError field if non-nil, zero value otherwise.

### GetValidatorErrorOk

`func (o *SettingDefinitionContract) GetValidatorErrorOk() (*string, bool)`

GetValidatorErrorOk returns a tuple with the ValidatorError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorError

`func (o *SettingDefinitionContract) SetValidatorError(v string)`

SetValidatorError sets ValidatorError field to given value.

### HasValidatorError

`func (o *SettingDefinitionContract) HasValidatorError() bool`

HasValidatorError returns a boolean if a field has been set.

### SetValidatorErrorNil

`func (o *SettingDefinitionContract) SetValidatorErrorNil(b bool)`

 SetValidatorErrorNil sets the value for ValidatorError to be an explicit nil

### UnsetValidatorError
`func (o *SettingDefinitionContract) UnsetValidatorError()`

UnsetValidatorError ensures that no value is present for ValidatorError, not even an explicit nil
### GetRelatedSettings

`func (o *SettingDefinitionContract) GetRelatedSettings() map[string]string`

GetRelatedSettings returns the RelatedSettings field if non-nil, zero value otherwise.

### GetRelatedSettingsOk

`func (o *SettingDefinitionContract) GetRelatedSettingsOk() (*map[string]string, bool)`

GetRelatedSettingsOk returns a tuple with the RelatedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedSettings

`func (o *SettingDefinitionContract) SetRelatedSettings(v map[string]string)`

SetRelatedSettings sets RelatedSettings field to given value.

### HasRelatedSettings

`func (o *SettingDefinitionContract) HasRelatedSettings() bool`

HasRelatedSettings returns a boolean if a field has been set.

### SetRelatedSettingsNil

`func (o *SettingDefinitionContract) SetRelatedSettingsNil(b bool)`

 SetRelatedSettingsNil sets the value for RelatedSettings to be an explicit nil

### UnsetRelatedSettings
`func (o *SettingDefinitionContract) UnsetRelatedSettings()`

UnsetRelatedSettings ensures that no value is present for RelatedSettings, not even an explicit nil
### GetValueUnit

`func (o *SettingDefinitionContract) GetValueUnit() string`

GetValueUnit returns the ValueUnit field if non-nil, zero value otherwise.

### GetValueUnitOk

`func (o *SettingDefinitionContract) GetValueUnitOk() (*string, bool)`

GetValueUnitOk returns a tuple with the ValueUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueUnit

`func (o *SettingDefinitionContract) SetValueUnit(v string)`

SetValueUnit sets ValueUnit field to given value.

### HasValueUnit

`func (o *SettingDefinitionContract) HasValueUnit() bool`

HasValueUnit returns a boolean if a field has been set.

### SetValueUnitNil

`func (o *SettingDefinitionContract) SetValueUnitNil(b bool)`

 SetValueUnitNil sets the value for ValueUnit to be an explicit nil

### UnsetValueUnit
`func (o *SettingDefinitionContract) UnsetValueUnit()`

UnsetValueUnit ensures that no value is present for ValueUnit, not even an explicit nil
### GetAllowedText

`func (o *SettingDefinitionContract) GetAllowedText() string`

GetAllowedText returns the AllowedText field if non-nil, zero value otherwise.

### GetAllowedTextOk

`func (o *SettingDefinitionContract) GetAllowedTextOk() (*string, bool)`

GetAllowedTextOk returns a tuple with the AllowedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedText

`func (o *SettingDefinitionContract) SetAllowedText(v string)`

SetAllowedText sets AllowedText field to given value.

### HasAllowedText

`func (o *SettingDefinitionContract) HasAllowedText() bool`

HasAllowedText returns a boolean if a field has been set.

### SetAllowedTextNil

`func (o *SettingDefinitionContract) SetAllowedTextNil(b bool)`

 SetAllowedTextNil sets the value for AllowedText to be an explicit nil

### UnsetAllowedText
`func (o *SettingDefinitionContract) UnsetAllowedText()`

UnsetAllowedText ensures that no value is present for AllowedText, not even an explicit nil
### GetDeniedText

`func (o *SettingDefinitionContract) GetDeniedText() string`

GetDeniedText returns the DeniedText field if non-nil, zero value otherwise.

### GetDeniedTextOk

`func (o *SettingDefinitionContract) GetDeniedTextOk() (*string, bool)`

GetDeniedTextOk returns a tuple with the DeniedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedText

`func (o *SettingDefinitionContract) SetDeniedText(v string)`

SetDeniedText sets DeniedText field to given value.

### HasDeniedText

`func (o *SettingDefinitionContract) HasDeniedText() bool`

HasDeniedText returns a boolean if a field has been set.

### SetDeniedTextNil

`func (o *SettingDefinitionContract) SetDeniedTextNil(b bool)`

 SetDeniedTextNil sets the value for DeniedText to be an explicit nil

### UnsetDeniedText
`func (o *SettingDefinitionContract) UnsetDeniedText()`

UnsetDeniedText ensures that no value is present for DeniedText, not even an explicit nil
### GetEnumType

`func (o *SettingDefinitionContract) GetEnumType() EnumerationTypeContract`

GetEnumType returns the EnumType field if non-nil, zero value otherwise.

### GetEnumTypeOk

`func (o *SettingDefinitionContract) GetEnumTypeOk() (*EnumerationTypeContract, bool)`

GetEnumTypeOk returns a tuple with the EnumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumType

`func (o *SettingDefinitionContract) SetEnumType(v EnumerationTypeContract)`

SetEnumType sets EnumType field to given value.

### HasEnumType

`func (o *SettingDefinitionContract) HasEnumType() bool`

HasEnumType returns a boolean if a field has been set.

### GetVdaVersions

`func (o *SettingDefinitionContract) GetVdaVersions() []VdaVersionContract`

GetVdaVersions returns the VdaVersions field if non-nil, zero value otherwise.

### GetVdaVersionsOk

`func (o *SettingDefinitionContract) GetVdaVersionsOk() (*[]VdaVersionContract, bool)`

GetVdaVersionsOk returns a tuple with the VdaVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaVersions

`func (o *SettingDefinitionContract) SetVdaVersions(v []VdaVersionContract)`

SetVdaVersions sets VdaVersions field to given value.

### HasVdaVersions

`func (o *SettingDefinitionContract) HasVdaVersions() bool`

HasVdaVersions returns a boolean if a field has been set.

### SetVdaVersionsNil

`func (o *SettingDefinitionContract) SetVdaVersionsNil(b bool)`

 SetVdaVersionsNil sets the value for VdaVersions to be an explicit nil

### UnsetVdaVersions
`func (o *SettingDefinitionContract) UnsetVdaVersions()`

UnsetVdaVersions ensures that no value is present for VdaVersions, not even an explicit nil
### GetGpoScope

`func (o *SettingDefinitionContract) GetGpoScope() string`

GetGpoScope returns the GpoScope field if non-nil, zero value otherwise.

### GetGpoScopeOk

`func (o *SettingDefinitionContract) GetGpoScopeOk() (*string, bool)`

GetGpoScopeOk returns a tuple with the GpoScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpoScope

`func (o *SettingDefinitionContract) SetGpoScope(v string)`

SetGpoScope sets GpoScope field to given value.

### HasGpoScope

`func (o *SettingDefinitionContract) HasGpoScope() bool`

HasGpoScope returns a boolean if a field has been set.

### SetGpoScopeNil

`func (o *SettingDefinitionContract) SetGpoScopeNil(b bool)`

 SetGpoScopeNil sets the value for GpoScope to be an explicit nil

### UnsetGpoScope
`func (o *SettingDefinitionContract) UnsetGpoScope()`

UnsetGpoScope ensures that no value is present for GpoScope, not even an explicit nil
### GetProductGroup

`func (o *SettingDefinitionContract) GetProductGroup() string`

GetProductGroup returns the ProductGroup field if non-nil, zero value otherwise.

### GetProductGroupOk

`func (o *SettingDefinitionContract) GetProductGroupOk() (*string, bool)`

GetProductGroupOk returns a tuple with the ProductGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroup

`func (o *SettingDefinitionContract) SetProductGroup(v string)`

SetProductGroup sets ProductGroup field to given value.

### HasProductGroup

`func (o *SettingDefinitionContract) HasProductGroup() bool`

HasProductGroup returns a boolean if a field has been set.

### SetProductGroupNil

`func (o *SettingDefinitionContract) SetProductGroupNil(b bool)`

 SetProductGroupNil sets the value for ProductGroup to be an explicit nil

### UnsetProductGroup
`func (o *SettingDefinitionContract) UnsetProductGroup()`

UnsetProductGroup ensures that no value is present for ProductGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


