# CreateConfigurationSlotRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the configuration slot to create. | 
**SettingsGroup** | **string** | Settings group of configuration slot to create. | 
**Description** | Pointer to **NullableString** | Description of configuration slot to create. | [optional] 

## Methods

### NewCreateConfigurationSlotRequestModel

`func NewCreateConfigurationSlotRequestModel(name string, settingsGroup string, ) *CreateConfigurationSlotRequestModel`

NewCreateConfigurationSlotRequestModel instantiates a new CreateConfigurationSlotRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConfigurationSlotRequestModelWithDefaults

`func NewCreateConfigurationSlotRequestModelWithDefaults() *CreateConfigurationSlotRequestModel`

NewCreateConfigurationSlotRequestModelWithDefaults instantiates a new CreateConfigurationSlotRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateConfigurationSlotRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConfigurationSlotRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConfigurationSlotRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetSettingsGroup

`func (o *CreateConfigurationSlotRequestModel) GetSettingsGroup() string`

GetSettingsGroup returns the SettingsGroup field if non-nil, zero value otherwise.

### GetSettingsGroupOk

`func (o *CreateConfigurationSlotRequestModel) GetSettingsGroupOk() (*string, bool)`

GetSettingsGroupOk returns a tuple with the SettingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsGroup

`func (o *CreateConfigurationSlotRequestModel) SetSettingsGroup(v string)`

SetSettingsGroup sets SettingsGroup field to given value.


### GetDescription

`func (o *CreateConfigurationSlotRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateConfigurationSlotRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateConfigurationSlotRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateConfigurationSlotRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateConfigurationSlotRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateConfigurationSlotRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


