# UpdateMachineConfigurationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | New name for machine configuration. This may have the form \&quot;ConfigurationSlotName\\MachineConfigurationName\&quot; or \&quot;MachineConfigurationName\&quot;. If the \&quot;ConfigurationSlotName\&quot; is provided it must match the name of the configuration slot that the machine configuration is associated with. | [optional] 
**Policy** | Pointer to **NullableString** | New binary array of encoded settings data. Policy settings data created with the SDK snap-in that matches the SettingsGroup of the configuration slot. Base64 encrypted. | [optional] 
**Description** | Pointer to **NullableString** | New description for the machine configuration.  | [optional] 

## Methods

### NewUpdateMachineConfigurationRequestModel

`func NewUpdateMachineConfigurationRequestModel() *UpdateMachineConfigurationRequestModel`

NewUpdateMachineConfigurationRequestModel instantiates a new UpdateMachineConfigurationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMachineConfigurationRequestModelWithDefaults

`func NewUpdateMachineConfigurationRequestModelWithDefaults() *UpdateMachineConfigurationRequestModel`

NewUpdateMachineConfigurationRequestModelWithDefaults instantiates a new UpdateMachineConfigurationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMachineConfigurationRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMachineConfigurationRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMachineConfigurationRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMachineConfigurationRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateMachineConfigurationRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateMachineConfigurationRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPolicy

`func (o *UpdateMachineConfigurationRequestModel) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateMachineConfigurationRequestModel) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateMachineConfigurationRequestModel) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpdateMachineConfigurationRequestModel) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *UpdateMachineConfigurationRequestModel) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *UpdateMachineConfigurationRequestModel) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetDescription

`func (o *UpdateMachineConfigurationRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMachineConfigurationRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMachineConfigurationRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMachineConfigurationRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateMachineConfigurationRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateMachineConfigurationRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


