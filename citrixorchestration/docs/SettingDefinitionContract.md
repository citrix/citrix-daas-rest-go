# SettingDefinitionContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditorName** | Pointer to **string** | The setting editor name. | [optional] 
**SettingName** | Pointer to **string** | Setting name; Is globally unique | [optional] 
**SettingType** | Pointer to **string** | The setting type as defined by its value. | [optional] 
**DisplayName** | Pointer to **string** | Translated setting name. Is globally unique. | [optional] 
**Category** | Pointer to **string** | Setting category, e.g., ICA\\Printing. | [optional] 
**Explanation** | Pointer to **string** | Setting detailed description | [optional] 
**IsUserSetting** | Pointer to **bool** | true &#x3D; user, false &#x3D; machine | [optional] 
**ValueType** | Pointer to **string** | Type of setting value | [optional] 
**IsEnableDisable** | Pointer to **bool** | Label test for a boolean setting, display enable/disable if true, otherwise allow/prohibit | [optional] 
**DefaultValue** | Pointer to **string** | Setting default value | [optional] 
**DisabledValue** | Pointer to **string** | Disabled value for toggled settings | [optional] 
**InitialValue** | Pointer to **string** | Initial value for taggled settings when value is not disabled. | [optional] 
**ValueMinimum** | Pointer to **string** | Minimal value for integer setting | [optional] 
**ValueMaximum** | Pointer to **string** | Maximal value for integer settin | [optional] 
**ValueValidator** | Pointer to **string** | An expression executed to validate the value | [optional] 
**ValidatorError** | Pointer to **string** | Warning message issued when validator fails (throws exception) | [optional] 
**RelatedSettings** | Pointer to **map[string]string** | Related settings. | [optional] 
**ValueUnit** | Pointer to **string** | Unit of value in translated text | [optional] 
**AllowedText** | Pointer to **string** | Custom text for the allowed explanation for some boolean settings. | [optional] 
**DeniedText** | Pointer to **string** | Custom text for the denied explaination for some boolean settings. | [optional] 
**EnumType** | Pointer to [**SettingDefinitionContractEnumType**](SettingDefinitionContractEnumType.md) |  | [optional] 
**VdaVersions** | Pointer to [**[]VdaVersionContract**](VdaVersionContract.md) | VDA versions honoring this setting. | [optional] 
**GpoScope** | Pointer to **string** | Scope, e.g. ConfigSlot | [optional] 
**ProductGroup** | Pointer to **string** | Group, e.g., UPM. | [optional] 

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

### GetEnumType

`func (o *SettingDefinitionContract) GetEnumType() SettingDefinitionContractEnumType`

GetEnumType returns the EnumType field if non-nil, zero value otherwise.

### GetEnumTypeOk

`func (o *SettingDefinitionContract) GetEnumTypeOk() (*SettingDefinitionContractEnumType, bool)`

GetEnumTypeOk returns a tuple with the EnumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumType

`func (o *SettingDefinitionContract) SetEnumType(v SettingDefinitionContractEnumType)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


