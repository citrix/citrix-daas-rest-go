# ComparisonRowContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **NullableString** | Can be either a category name or setting name. | [optional] 
**IsSetting** | Pointer to **bool** | Indicate if SettingName is a setting name or category name. | [optional] 
**IsDifferent** | Pointer to **bool** | Valid only when IsSetting is false. Used to indicate if there is any difference among the settings in the category. | [optional] 
**IsInUse** | Pointer to **[]bool** | There are Targets.Length number of items in the list, stored in the same order as the items in Targets. Value is true if the setting is configured in the target. Value is true if SettingName is a category name and at least one of the settings under the category has IsInUse set to true. | [optional] 
**Values** | Pointer to **[]string** | There are Targets.Length number of items in the list, stored in the same order as the items in Targets. Value should not be used if the corresponding IsInUse value is false. If the IsInUse value is false, a - should be displayed if IsSetting is true. Don&#39;t use the values here if IsSetting is false. | [optional] 
**Category** | Pointer to **NullableString** | For internal use to keep track of setting category. | [optional] 
**DefaultValue** | Pointer to **NullableString** | The default value as specified in the GPFX file for the setting. | [optional] 

## Methods

### NewComparisonRowContract

`func NewComparisonRowContract() *ComparisonRowContract`

NewComparisonRowContract instantiates a new ComparisonRowContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparisonRowContractWithDefaults

`func NewComparisonRowContractWithDefaults() *ComparisonRowContract`

NewComparisonRowContractWithDefaults instantiates a new ComparisonRowContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingName

`func (o *ComparisonRowContract) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *ComparisonRowContract) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *ComparisonRowContract) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *ComparisonRowContract) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *ComparisonRowContract) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *ComparisonRowContract) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetIsSetting

`func (o *ComparisonRowContract) GetIsSetting() bool`

GetIsSetting returns the IsSetting field if non-nil, zero value otherwise.

### GetIsSettingOk

`func (o *ComparisonRowContract) GetIsSettingOk() (*bool, bool)`

GetIsSettingOk returns a tuple with the IsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetting

`func (o *ComparisonRowContract) SetIsSetting(v bool)`

SetIsSetting sets IsSetting field to given value.

### HasIsSetting

`func (o *ComparisonRowContract) HasIsSetting() bool`

HasIsSetting returns a boolean if a field has been set.

### GetIsDifferent

`func (o *ComparisonRowContract) GetIsDifferent() bool`

GetIsDifferent returns the IsDifferent field if non-nil, zero value otherwise.

### GetIsDifferentOk

`func (o *ComparisonRowContract) GetIsDifferentOk() (*bool, bool)`

GetIsDifferentOk returns a tuple with the IsDifferent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDifferent

`func (o *ComparisonRowContract) SetIsDifferent(v bool)`

SetIsDifferent sets IsDifferent field to given value.

### HasIsDifferent

`func (o *ComparisonRowContract) HasIsDifferent() bool`

HasIsDifferent returns a boolean if a field has been set.

### GetIsInUse

`func (o *ComparisonRowContract) GetIsInUse() []bool`

GetIsInUse returns the IsInUse field if non-nil, zero value otherwise.

### GetIsInUseOk

`func (o *ComparisonRowContract) GetIsInUseOk() (*[]bool, bool)`

GetIsInUseOk returns a tuple with the IsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInUse

`func (o *ComparisonRowContract) SetIsInUse(v []bool)`

SetIsInUse sets IsInUse field to given value.

### HasIsInUse

`func (o *ComparisonRowContract) HasIsInUse() bool`

HasIsInUse returns a boolean if a field has been set.

### SetIsInUseNil

`func (o *ComparisonRowContract) SetIsInUseNil(b bool)`

 SetIsInUseNil sets the value for IsInUse to be an explicit nil

### UnsetIsInUse
`func (o *ComparisonRowContract) UnsetIsInUse()`

UnsetIsInUse ensures that no value is present for IsInUse, not even an explicit nil
### GetValues

`func (o *ComparisonRowContract) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ComparisonRowContract) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ComparisonRowContract) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ComparisonRowContract) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ComparisonRowContract) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ComparisonRowContract) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetCategory

`func (o *ComparisonRowContract) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ComparisonRowContract) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ComparisonRowContract) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ComparisonRowContract) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ComparisonRowContract) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ComparisonRowContract) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDefaultValue

`func (o *ComparisonRowContract) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ComparisonRowContract) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ComparisonRowContract) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ComparisonRowContract) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ComparisonRowContract) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ComparisonRowContract) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


