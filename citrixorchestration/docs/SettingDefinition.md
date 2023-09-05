# SettingDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **string** | Setting name; Is globally unique | [optional] 
**EditorName** | Pointer to **string** | The setting editor name. | [optional] 
**DisplayName** | Pointer to **string** | Translated setting name. Is globally unique. | [optional] 
**Category** | Pointer to **string** | Setting category, e.g., ICA\\Printing. | [optional] 
**Explanation** | Pointer to **string** | Setting detailed description | [optional] 
**IsUserSetting** | **bool** | true &#x3D; user, false &#x3D; machine | 
**ValueType** | Pointer to **string** | Type of setting value | [optional] 
**IsEnableDisable** | **bool** | Label test for a boolean setting, display enable/disable if true, otherwise allow/prohibit | 
**DefaultValue** | Pointer to **string** | Setting default value | [optional] 
**DisabledValue** | Pointer to **string** | Disabled value for toggled settings | [optional] 
**InitialValue** | Pointer to **string** | Initial value for toggled settings when value is not disabled. | [optional] 
**ValueMinimum** | Pointer to **string** | Minimal value for integer setting | [optional] 
**ValueMaximum** | Pointer to **string** | Maximal value for integer setting | [optional] 
**ValueValidator** | Pointer to **string** | An expression executed to validate the value | [optional] 
**ValidatorError** | Pointer to **string** | Warning message issued when validator fails (throws exception) | [optional] 
**RelatedSettings** | Pointer to **[]string** | Related settings. | [optional] 
**ValueUnit** | Pointer to **string** | Unit of value in translated text | [optional] 
**AllowedText** | Pointer to **string** | Custom text for the allowed explanation for some boolean settings. | [optional] 
**DeniedText** | Pointer to **string** | Custom text for the denied explanation for some boolean settings. | [optional] 
**EnumType** | Pointer to [**SettingDefinitionEnumType**](SettingDefinitionEnumType.md) |  | [optional] 
**VersionDetails** | Pointer to [**[]VersionDetail**](VersionDetail.md) | VDA versions honoring this setting. | [optional] 
**VersionCode** | Pointer to **string** | The version string in GPFX file, e.g., XD&#x3D;7.8,*,B. | [optional] 
**VdaVersions** | Pointer to **[]string** | Array of VDA versions, e.g., 7.8S, 7.8M. This is much lighter than VersionDetails. | [optional] 
**GpoScope** | Pointer to **string** | Scope, e.g. ConfigSlot | [optional] 
**ProductGroup** | Pointer to **string** | Group, e.g., UPM. | [optional] 

## Methods

### NewSettingDefinition

`func NewSettingDefinition(isUserSetting bool, isEnableDisable bool, ) *SettingDefinition`

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

### GetEnumType

`func (o *SettingDefinition) GetEnumType() SettingDefinitionEnumType`

GetEnumType returns the EnumType field if non-nil, zero value otherwise.

### GetEnumTypeOk

`func (o *SettingDefinition) GetEnumTypeOk() (*SettingDefinitionEnumType, bool)`

GetEnumTypeOk returns a tuple with the EnumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumType

`func (o *SettingDefinition) SetEnumType(v SettingDefinitionEnumType)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


