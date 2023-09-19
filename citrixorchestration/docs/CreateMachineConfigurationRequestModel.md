# CreateMachineConfigurationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationSlotId** | **int32** | Id of the configuration slot to associate with this machine configuration. | 
**LeafName** | **string** | Name of the new machine configuration. | 
**Policy** | **string** | Policy settings data created with the SDK snap-in that matches the SettingsGroup of the configuration slot. Base64 encrypted. | 
**Description** | Pointer to **NullableString** | Description of the new machine configuration. | [optional] 

## Methods

### NewCreateMachineConfigurationRequestModel

`func NewCreateMachineConfigurationRequestModel(configurationSlotId int32, leafName string, policy string, ) *CreateMachineConfigurationRequestModel`

NewCreateMachineConfigurationRequestModel instantiates a new CreateMachineConfigurationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMachineConfigurationRequestModelWithDefaults

`func NewCreateMachineConfigurationRequestModelWithDefaults() *CreateMachineConfigurationRequestModel`

NewCreateMachineConfigurationRequestModelWithDefaults instantiates a new CreateMachineConfigurationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationSlotId

`func (o *CreateMachineConfigurationRequestModel) GetConfigurationSlotId() int32`

GetConfigurationSlotId returns the ConfigurationSlotId field if non-nil, zero value otherwise.

### GetConfigurationSlotIdOk

`func (o *CreateMachineConfigurationRequestModel) GetConfigurationSlotIdOk() (*int32, bool)`

GetConfigurationSlotIdOk returns a tuple with the ConfigurationSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSlotId

`func (o *CreateMachineConfigurationRequestModel) SetConfigurationSlotId(v int32)`

SetConfigurationSlotId sets ConfigurationSlotId field to given value.


### GetLeafName

`func (o *CreateMachineConfigurationRequestModel) GetLeafName() string`

GetLeafName returns the LeafName field if non-nil, zero value otherwise.

### GetLeafNameOk

`func (o *CreateMachineConfigurationRequestModel) GetLeafNameOk() (*string, bool)`

GetLeafNameOk returns a tuple with the LeafName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafName

`func (o *CreateMachineConfigurationRequestModel) SetLeafName(v string)`

SetLeafName sets LeafName field to given value.


### GetPolicy

`func (o *CreateMachineConfigurationRequestModel) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreateMachineConfigurationRequestModel) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreateMachineConfigurationRequestModel) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetDescription

`func (o *CreateMachineConfigurationRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMachineConfigurationRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMachineConfigurationRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMachineConfigurationRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateMachineConfigurationRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateMachineConfigurationRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


