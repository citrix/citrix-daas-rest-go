# ConfigurationSlotResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The Id of a configuration slot. | [optional] 
**Name** | Pointer to **NullableString** | The name of a configuration slot. | [optional] 
**SettingsGroup** | Pointer to **NullableString** | The SettingsGroup of a configuration slot. | [optional] 
**Description** | Pointer to **NullableString** | The Description of a configuration slot. | [optional] 

## Methods

### NewConfigurationSlotResponseModel

`func NewConfigurationSlotResponseModel() *ConfigurationSlotResponseModel`

NewConfigurationSlotResponseModel instantiates a new ConfigurationSlotResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationSlotResponseModelWithDefaults

`func NewConfigurationSlotResponseModelWithDefaults() *ConfigurationSlotResponseModel`

NewConfigurationSlotResponseModelWithDefaults instantiates a new ConfigurationSlotResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationSlotResponseModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationSlotResponseModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationSlotResponseModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationSlotResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConfigurationSlotResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationSlotResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationSlotResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigurationSlotResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConfigurationSlotResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfigurationSlotResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSettingsGroup

`func (o *ConfigurationSlotResponseModel) GetSettingsGroup() string`

GetSettingsGroup returns the SettingsGroup field if non-nil, zero value otherwise.

### GetSettingsGroupOk

`func (o *ConfigurationSlotResponseModel) GetSettingsGroupOk() (*string, bool)`

GetSettingsGroupOk returns a tuple with the SettingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsGroup

`func (o *ConfigurationSlotResponseModel) SetSettingsGroup(v string)`

SetSettingsGroup sets SettingsGroup field to given value.

### HasSettingsGroup

`func (o *ConfigurationSlotResponseModel) HasSettingsGroup() bool`

HasSettingsGroup returns a boolean if a field has been set.

### SetSettingsGroupNil

`func (o *ConfigurationSlotResponseModel) SetSettingsGroupNil(b bool)`

 SetSettingsGroupNil sets the value for SettingsGroup to be an explicit nil

### UnsetSettingsGroup
`func (o *ConfigurationSlotResponseModel) UnsetSettingsGroup()`

UnsetSettingsGroup ensures that no value is present for SettingsGroup, not even an explicit nil
### GetDescription

`func (o *ConfigurationSlotResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationSlotResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationSlotResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationSlotResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigurationSlotResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigurationSlotResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


