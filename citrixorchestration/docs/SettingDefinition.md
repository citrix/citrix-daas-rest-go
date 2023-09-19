# SettingDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **NullableString** | Setting name; Is globally unique | [optional] 
**EditorName** | Pointer to **NullableString** | The setting editor name. | [optional] 
**DisplayName** | Pointer to **NullableString** | Translated setting name. Is globally unique. | [optional] 
**Category** | Pointer to **NullableString** | Setting category, e.g., ICA\\Printing. | [optional] 
**Explanation** | Pointer to **NullableString** | Setting detailed description | [optional] 
**IsUserSetting** | Pointer to **bool** | true &#x3D; user, false &#x3D; machine | [optional] 
**ValueType** | Pointer to **NullableString** | Type of setting value | [optional] 
**IsEnableDisable** | Pointer to **bool** | Label test for a boolean setting, display enable/disable if true, otherwise allow/prohibit | [optional] 
**IsToggled** | Pointer to **bool** | Indicate if the integer, string, enum or multiline string is toggled. | [optional] 
**DefaultValue** | Pointer to **NullableString** | Setting default value | [optional] 
**DisabledValue** | Pointer to **NullableString** | Disabled value for toggled settings | [optional] 
**InitialValue** | Pointer to **NullableString** | Initial value for toggled settings when value is not disabled. | [optional] 
**ValueMinimum** | Pointer to **NullableString** | Minimal value for integer setting | [optional] 
**ValueMaximum** | Pointer to **NullableString** | Maximal value for integer setting | [optional] 
**ValueValidator** | Pointer to **NullableString** | An expression executed to validate the value | [optional] 
**ValidatorError** | Pointer to **NullableString** | Warning message issued when validator fails (throws exception) | [optional] 
**RelatedSettings** | Pointer to **[]string** | Related settings. | [optional] 
**ValueUnit** | Pointer to **NullableString** | Unit of value in translated text | [optional] 
**AllowedText** | Pointer to **NullableString** | Custom text for the allowed explanation for some boolean settings. | [optional] 
**DeniedText** | Pointer to **NullableString** | Custom text for the denied explanation for some boolean settings. | [optional] 
**EnumType** | Pointer to [**EnumerationType**](EnumerationType.md) |  | [optional] 
**VersionDetails** | Pointer to [**[]VersionDetail**](VersionDetail.md) | VDA versions honoring this setting. | [optional] 
**VersionCode** | Pointer to **NullableString** | The version string in GPFX file, e.g., XD&#x3D;7.8,*,B. | [optional] 
**VdaVersions** | Pointer to **[]string** | Array of VDA versions, e.g., 7.8S, 7.8M. This is much lighter than VersionDetails. | [optional] 
**GpoScope** | Pointer to **NullableString** | Scope, e.g. ConfigSlot | [optional] 
**ProductGroup** | Pointer to **NullableString** | Group, e.g., UPM. | [optional] 

## Methods

### NewSettingDefinition

`func NewSettingDefinition() *SettingDefinition`

NewSettingDefinition instantiates a new SettingDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingDefinitionWithDefaults

`func NewSettingDefinitionWithDefaults() *SettingDefinition`

NewSettingDefinitionWithDefaults instantiates a new SettingDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingName

`func (o *SettingDefinition) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingDefinition) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingDefinition) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *SettingDefinition) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *SettingDefinition) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *SettingDefinition) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetEditorName

`func (o *SettingDefinition) GetEditorName() string`

GetEditorName returns the EditorName field if non-nil, zero value otherwise.

### GetEditorNameOk

`func (o *SettingDefinition) GetEditorNameOk() (*string, bool)`

GetEditorNameOk returns a tuple with the EditorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorName

`func (o *SettingDefinition) SetEditorName(v string)`

SetEditorName sets EditorName field to given value.

### HasEditorName

`func (o *SettingDefinition) HasEditorName() bool`

HasEditorName returns a boolean if a field has been set.

### SetEditorNameNil

`func (o *SettingDefinition) SetEditorNameNil(b bool)`

 SetEditorNameNil sets the value for EditorName to be an explicit nil

### UnsetEditorName
`func (o *SettingDefinition) UnsetEditorName()`

UnsetEditorName ensures that no value is present for EditorName, not even an explicit nil
### GetDisplayName

`func (o *SettingDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SettingDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SettingDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SettingDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SettingDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SettingDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetCategory

`func (o *SettingDefinition) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SettingDefinition) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SettingDefinition) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SettingDefinition) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *SettingDefinition) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *SettingDefinition) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetExplanation

`func (o *SettingDefinition) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *SettingDefinition) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *SettingDefinition) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *SettingDefinition) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### SetExplanationNil

`func (o *SettingDefinition) SetExplanationNil(b bool)`

 SetExplanationNil sets the value for Explanation to be an explicit nil

### UnsetExplanation
`func (o *SettingDefinition) UnsetExplanation()`

UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil
### GetIsUserSetting

`func (o *SettingDefinition) GetIsUserSetting() bool`

GetIsUserSetting returns the IsUserSetting field if non-nil, zero value otherwise.

### GetIsUserSettingOk

`func (o *SettingDefinition) GetIsUserSettingOk() (*bool, bool)`

GetIsUserSettingOk returns a tuple with the IsUserSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserSetting

`func (o *SettingDefinition) SetIsUserSetting(v bool)`

SetIsUserSetting sets IsUserSetting field to given value.

### HasIsUserSetting

`func (o *SettingDefinition) HasIsUserSetting() bool`

HasIsUserSetting returns a boolean if a field has been set.

### GetValueType

`func (o *SettingDefinition) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *SettingDefinition) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *SettingDefinition) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *SettingDefinition) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueTypeNil

`func (o *SettingDefinition) SetValueTypeNil(b bool)`

 SetValueTypeNil sets the value for ValueType to be an explicit nil

### UnsetValueType
`func (o *SettingDefinition) UnsetValueType()`

UnsetValueType ensures that no value is present for ValueType, not even an explicit nil
### GetIsEnableDisable

`func (o *SettingDefinition) GetIsEnableDisable() bool`

GetIsEnableDisable returns the IsEnableDisable field if non-nil, zero value otherwise.

### GetIsEnableDisableOk

`func (o *SettingDefinition) GetIsEnableDisableOk() (*bool, bool)`

GetIsEnableDisableOk returns a tuple with the IsEnableDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableDisable

`func (o *SettingDefinition) SetIsEnableDisable(v bool)`

SetIsEnableDisable sets IsEnableDisable field to given value.

### HasIsEnableDisable

`func (o *SettingDefinition) HasIsEnableDisable() bool`

HasIsEnableDisable returns a boolean if a field has been set.

### GetIsToggled

`func (o *SettingDefinition) GetIsToggled() bool`

GetIsToggled returns the IsToggled field if non-nil, zero value otherwise.

### GetIsToggledOk

`func (o *SettingDefinition) GetIsToggledOk() (*bool, bool)`

GetIsToggledOk returns a tuple with the IsToggled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsToggled

`func (o *SettingDefinition) SetIsToggled(v bool)`

SetIsToggled sets IsToggled field to given value.

### HasIsToggled

`func (o *SettingDefinition) HasIsToggled() bool`

HasIsToggled returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SettingDefinition) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SettingDefinition) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SettingDefinition) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SettingDefinition) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *SettingDefinition) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *SettingDefinition) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDisabledValue

`func (o *SettingDefinition) GetDisabledValue() string`

GetDisabledValue returns the DisabledValue field if non-nil, zero value otherwise.

### GetDisabledValueOk

`func (o *SettingDefinition) GetDisabledValueOk() (*string, bool)`

GetDisabledValueOk returns a tuple with the DisabledValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledValue

`func (o *SettingDefinition) SetDisabledValue(v string)`

SetDisabledValue sets DisabledValue field to given value.

### HasDisabledValue

`func (o *SettingDefinition) HasDisabledValue() bool`

HasDisabledValue returns a boolean if a field has been set.

### SetDisabledValueNil

`func (o *SettingDefinition) SetDisabledValueNil(b bool)`

 SetDisabledValueNil sets the value for DisabledValue to be an explicit nil

### UnsetDisabledValue
`func (o *SettingDefinition) UnsetDisabledValue()`

UnsetDisabledValue ensures that no value is present for DisabledValue, not even an explicit nil
### GetInitialValue

`func (o *SettingDefinition) GetInitialValue() string`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *SettingDefinition) GetInitialValueOk() (*string, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *SettingDefinition) SetInitialValue(v string)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *SettingDefinition) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### SetInitialValueNil

`func (o *SettingDefinition) SetInitialValueNil(b bool)`

 SetInitialValueNil sets the value for InitialValue to be an explicit nil

### UnsetInitialValue
`func (o *SettingDefinition) UnsetInitialValue()`

UnsetInitialValue ensures that no value is present for InitialValue, not even an explicit nil
### GetValueMinimum

`func (o *SettingDefinition) GetValueMinimum() string`

GetValueMinimum returns the ValueMinimum field if non-nil, zero value otherwise.

### GetValueMinimumOk

`func (o *SettingDefinition) GetValueMinimumOk() (*string, bool)`

GetValueMinimumOk returns a tuple with the ValueMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMinimum

`func (o *SettingDefinition) SetValueMinimum(v string)`

SetValueMinimum sets ValueMinimum field to given value.

### HasValueMinimum

`func (o *SettingDefinition) HasValueMinimum() bool`

HasValueMinimum returns a boolean if a field has been set.

### SetValueMinimumNil

`func (o *SettingDefinition) SetValueMinimumNil(b bool)`

 SetValueMinimumNil sets the value for ValueMinimum to be an explicit nil

### UnsetValueMinimum
`func (o *SettingDefinition) UnsetValueMinimum()`

UnsetValueMinimum ensures that no value is present for ValueMinimum, not even an explicit nil
### GetValueMaximum

`func (o *SettingDefinition) GetValueMaximum() string`

GetValueMaximum returns the ValueMaximum field if non-nil, zero value otherwise.

### GetValueMaximumOk

`func (o *SettingDefinition) GetValueMaximumOk() (*string, bool)`

GetValueMaximumOk returns a tuple with the ValueMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMaximum

`func (o *SettingDefinition) SetValueMaximum(v string)`

SetValueMaximum sets ValueMaximum field to given value.

### HasValueMaximum

`func (o *SettingDefinition) HasValueMaximum() bool`

HasValueMaximum returns a boolean if a field has been set.

### SetValueMaximumNil

`func (o *SettingDefinition) SetValueMaximumNil(b bool)`

 SetValueMaximumNil sets the value for ValueMaximum to be an explicit nil

### UnsetValueMaximum
`func (o *SettingDefinition) UnsetValueMaximum()`

UnsetValueMaximum ensures that no value is present for ValueMaximum, not even an explicit nil
### GetValueValidator

`func (o *SettingDefinition) GetValueValidator() string`

GetValueValidator returns the ValueValidator field if non-nil, zero value otherwise.

### GetValueValidatorOk

`func (o *SettingDefinition) GetValueValidatorOk() (*string, bool)`

GetValueValidatorOk returns a tuple with the ValueValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueValidator

`func (o *SettingDefinition) SetValueValidator(v string)`

SetValueValidator sets ValueValidator field to given value.

### HasValueValidator

`func (o *SettingDefinition) HasValueValidator() bool`

HasValueValidator returns a boolean if a field has been set.

### SetValueValidatorNil

`func (o *SettingDefinition) SetValueValidatorNil(b bool)`

 SetValueValidatorNil sets the value for ValueValidator to be an explicit nil

### UnsetValueValidator
`func (o *SettingDefinition) UnsetValueValidator()`

UnsetValueValidator ensures that no value is present for ValueValidator, not even an explicit nil
### GetValidatorError

`func (o *SettingDefinition) GetValidatorError() string`

GetValidatorError returns the ValidatorError field if non-nil, zero value otherwise.

### GetValidatorErrorOk

`func (o *SettingDefinition) GetValidatorErrorOk() (*string, bool)`

GetValidatorErrorOk returns a tuple with the ValidatorError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorError

`func (o *SettingDefinition) SetValidatorError(v string)`

SetValidatorError sets ValidatorError field to given value.

### HasValidatorError

`func (o *SettingDefinition) HasValidatorError() bool`

HasValidatorError returns a boolean if a field has been set.

### SetValidatorErrorNil

`func (o *SettingDefinition) SetValidatorErrorNil(b bool)`

 SetValidatorErrorNil sets the value for ValidatorError to be an explicit nil

### UnsetValidatorError
`func (o *SettingDefinition) UnsetValidatorError()`

UnsetValidatorError ensures that no value is present for ValidatorError, not even an explicit nil
### GetRelatedSettings

`func (o *SettingDefinition) GetRelatedSettings() []string`

GetRelatedSettings returns the RelatedSettings field if non-nil, zero value otherwise.

### GetRelatedSettingsOk

`func (o *SettingDefinition) GetRelatedSettingsOk() (*[]string, bool)`

GetRelatedSettingsOk returns a tuple with the RelatedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedSettings

`func (o *SettingDefinition) SetRelatedSettings(v []string)`

SetRelatedSettings sets RelatedSettings field to given value.

### HasRelatedSettings

`func (o *SettingDefinition) HasRelatedSettings() bool`

HasRelatedSettings returns a boolean if a field has been set.

### SetRelatedSettingsNil

`func (o *SettingDefinition) SetRelatedSettingsNil(b bool)`

 SetRelatedSettingsNil sets the value for RelatedSettings to be an explicit nil

### UnsetRelatedSettings
`func (o *SettingDefinition) UnsetRelatedSettings()`

UnsetRelatedSettings ensures that no value is present for RelatedSettings, not even an explicit nil
### GetValueUnit

`func (o *SettingDefinition) GetValueUnit() string`

GetValueUnit returns the ValueUnit field if non-nil, zero value otherwise.

### GetValueUnitOk

`func (o *SettingDefinition) GetValueUnitOk() (*string, bool)`

GetValueUnitOk returns a tuple with the ValueUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueUnit

`func (o *SettingDefinition) SetValueUnit(v string)`

SetValueUnit sets ValueUnit field to given value.

### HasValueUnit

`func (o *SettingDefinition) HasValueUnit() bool`

HasValueUnit returns a boolean if a field has been set.

### SetValueUnitNil

`func (o *SettingDefinition) SetValueUnitNil(b bool)`

 SetValueUnitNil sets the value for ValueUnit to be an explicit nil

### UnsetValueUnit
`func (o *SettingDefinition) UnsetValueUnit()`

UnsetValueUnit ensures that no value is present for ValueUnit, not even an explicit nil
### GetAllowedText

`func (o *SettingDefinition) GetAllowedText() string`

GetAllowedText returns the AllowedText field if non-nil, zero value otherwise.

### GetAllowedTextOk

`func (o *SettingDefinition) GetAllowedTextOk() (*string, bool)`

GetAllowedTextOk returns a tuple with the AllowedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedText

`func (o *SettingDefinition) SetAllowedText(v string)`

SetAllowedText sets AllowedText field to given value.

### HasAllowedText

`func (o *SettingDefinition) HasAllowedText() bool`

HasAllowedText returns a boolean if a field has been set.

### SetAllowedTextNil

`func (o *SettingDefinition) SetAllowedTextNil(b bool)`

 SetAllowedTextNil sets the value for AllowedText to be an explicit nil

### UnsetAllowedText
`func (o *SettingDefinition) UnsetAllowedText()`

UnsetAllowedText ensures that no value is present for AllowedText, not even an explicit nil
### GetDeniedText

`func (o *SettingDefinition) GetDeniedText() string`

GetDeniedText returns the DeniedText field if non-nil, zero value otherwise.

### GetDeniedTextOk

`func (o *SettingDefinition) GetDeniedTextOk() (*string, bool)`

GetDeniedTextOk returns a tuple with the DeniedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedText

`func (o *SettingDefinition) SetDeniedText(v string)`

SetDeniedText sets DeniedText field to given value.

### HasDeniedText

`func (o *SettingDefinition) HasDeniedText() bool`

HasDeniedText returns a boolean if a field has been set.

### SetDeniedTextNil

`func (o *SettingDefinition) SetDeniedTextNil(b bool)`

 SetDeniedTextNil sets the value for DeniedText to be an explicit nil

### UnsetDeniedText
`func (o *SettingDefinition) UnsetDeniedText()`

UnsetDeniedText ensures that no value is present for DeniedText, not even an explicit nil
### GetEnumType

`func (o *SettingDefinition) GetEnumType() EnumerationType`

GetEnumType returns the EnumType field if non-nil, zero value otherwise.

### GetEnumTypeOk

`func (o *SettingDefinition) GetEnumTypeOk() (*EnumerationType, bool)`

GetEnumTypeOk returns a tuple with the EnumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumType

`func (o *SettingDefinition) SetEnumType(v EnumerationType)`

SetEnumType sets EnumType field to given value.

### HasEnumType

`func (o *SettingDefinition) HasEnumType() bool`

HasEnumType returns a boolean if a field has been set.

### GetVersionDetails

`func (o *SettingDefinition) GetVersionDetails() []VersionDetail`

GetVersionDetails returns the VersionDetails field if non-nil, zero value otherwise.

### GetVersionDetailsOk

`func (o *SettingDefinition) GetVersionDetailsOk() (*[]VersionDetail, bool)`

GetVersionDetailsOk returns a tuple with the VersionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionDetails

`func (o *SettingDefinition) SetVersionDetails(v []VersionDetail)`

SetVersionDetails sets VersionDetails field to given value.

### HasVersionDetails

`func (o *SettingDefinition) HasVersionDetails() bool`

HasVersionDetails returns a boolean if a field has been set.

### SetVersionDetailsNil

`func (o *SettingDefinition) SetVersionDetailsNil(b bool)`

 SetVersionDetailsNil sets the value for VersionDetails to be an explicit nil

### UnsetVersionDetails
`func (o *SettingDefinition) UnsetVersionDetails()`

UnsetVersionDetails ensures that no value is present for VersionDetails, not even an explicit nil
### GetVersionCode

`func (o *SettingDefinition) GetVersionCode() string`

GetVersionCode returns the VersionCode field if non-nil, zero value otherwise.

### GetVersionCodeOk

`func (o *SettingDefinition) GetVersionCodeOk() (*string, bool)`

GetVersionCodeOk returns a tuple with the VersionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCode

`func (o *SettingDefinition) SetVersionCode(v string)`

SetVersionCode sets VersionCode field to given value.

### HasVersionCode

`func (o *SettingDefinition) HasVersionCode() bool`

HasVersionCode returns a boolean if a field has been set.

### SetVersionCodeNil

`func (o *SettingDefinition) SetVersionCodeNil(b bool)`

 SetVersionCodeNil sets the value for VersionCode to be an explicit nil

### UnsetVersionCode
`func (o *SettingDefinition) UnsetVersionCode()`

UnsetVersionCode ensures that no value is present for VersionCode, not even an explicit nil
### GetVdaVersions

`func (o *SettingDefinition) GetVdaVersions() []string`

GetVdaVersions returns the VdaVersions field if non-nil, zero value otherwise.

### GetVdaVersionsOk

`func (o *SettingDefinition) GetVdaVersionsOk() (*[]string, bool)`

GetVdaVersionsOk returns a tuple with the VdaVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaVersions

`func (o *SettingDefinition) SetVdaVersions(v []string)`

SetVdaVersions sets VdaVersions field to given value.

### HasVdaVersions

`func (o *SettingDefinition) HasVdaVersions() bool`

HasVdaVersions returns a boolean if a field has been set.

### SetVdaVersionsNil

`func (o *SettingDefinition) SetVdaVersionsNil(b bool)`

 SetVdaVersionsNil sets the value for VdaVersions to be an explicit nil

### UnsetVdaVersions
`func (o *SettingDefinition) UnsetVdaVersions()`

UnsetVdaVersions ensures that no value is present for VdaVersions, not even an explicit nil
### GetGpoScope

`func (o *SettingDefinition) GetGpoScope() string`

GetGpoScope returns the GpoScope field if non-nil, zero value otherwise.

### GetGpoScopeOk

`func (o *SettingDefinition) GetGpoScopeOk() (*string, bool)`

GetGpoScopeOk returns a tuple with the GpoScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpoScope

`func (o *SettingDefinition) SetGpoScope(v string)`

SetGpoScope sets GpoScope field to given value.

### HasGpoScope

`func (o *SettingDefinition) HasGpoScope() bool`

HasGpoScope returns a boolean if a field has been set.

### SetGpoScopeNil

`func (o *SettingDefinition) SetGpoScopeNil(b bool)`

 SetGpoScopeNil sets the value for GpoScope to be an explicit nil

### UnsetGpoScope
`func (o *SettingDefinition) UnsetGpoScope()`

UnsetGpoScope ensures that no value is present for GpoScope, not even an explicit nil
### GetProductGroup

`func (o *SettingDefinition) GetProductGroup() string`

GetProductGroup returns the ProductGroup field if non-nil, zero value otherwise.

### GetProductGroupOk

`func (o *SettingDefinition) GetProductGroupOk() (*string, bool)`

GetProductGroupOk returns a tuple with the ProductGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroup

`func (o *SettingDefinition) SetProductGroup(v string)`

SetProductGroup sets ProductGroup field to given value.

### HasProductGroup

`func (o *SettingDefinition) HasProductGroup() bool`

HasProductGroup returns a boolean if a field has been set.

### SetProductGroupNil

`func (o *SettingDefinition) SetProductGroupNil(b bool)`

 SetProductGroupNil sets the value for ProductGroup to be an explicit nil

### UnsetProductGroup
`func (o *SettingDefinition) UnsetProductGroup()`

UnsetProductGroup ensures that no value is present for ProductGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


