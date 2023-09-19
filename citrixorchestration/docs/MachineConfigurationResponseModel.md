# MachineConfigurationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The Id of a machine configuration. | [optional] 
**ConfigurationSlot** | Pointer to [**ConfigurationSlotResponseModel**](ConfigurationSlotResponseModel.md) |  | [optional] 
**Name** | Pointer to **NullableString** | The name of a machine configuration. | [optional] 
**LeafName** | Pointer to **NullableString** | The name of a machine configuration. | [optional] 
**Policy** | Pointer to **NullableString** | The policy of a configuration slot. | [optional] 
**Description** | Pointer to **NullableString** | The Description of a machine configuration. | [optional] 
**DeliveryGroups** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The associated desktop grouds ref. | [optional] 

## Methods

### NewMachineConfigurationResponseModel

`func NewMachineConfigurationResponseModel() *MachineConfigurationResponseModel`

NewMachineConfigurationResponseModel instantiates a new MachineConfigurationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineConfigurationResponseModelWithDefaults

`func NewMachineConfigurationResponseModelWithDefaults() *MachineConfigurationResponseModel`

NewMachineConfigurationResponseModelWithDefaults instantiates a new MachineConfigurationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineConfigurationResponseModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineConfigurationResponseModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineConfigurationResponseModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MachineConfigurationResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationSlot

`func (o *MachineConfigurationResponseModel) GetConfigurationSlot() ConfigurationSlotResponseModel`

GetConfigurationSlot returns the ConfigurationSlot field if non-nil, zero value otherwise.

### GetConfigurationSlotOk

`func (o *MachineConfigurationResponseModel) GetConfigurationSlotOk() (*ConfigurationSlotResponseModel, bool)`

GetConfigurationSlotOk returns a tuple with the ConfigurationSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSlot

`func (o *MachineConfigurationResponseModel) SetConfigurationSlot(v ConfigurationSlotResponseModel)`

SetConfigurationSlot sets ConfigurationSlot field to given value.

### HasConfigurationSlot

`func (o *MachineConfigurationResponseModel) HasConfigurationSlot() bool`

HasConfigurationSlot returns a boolean if a field has been set.

### GetName

`func (o *MachineConfigurationResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineConfigurationResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineConfigurationResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineConfigurationResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MachineConfigurationResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MachineConfigurationResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLeafName

`func (o *MachineConfigurationResponseModel) GetLeafName() string`

GetLeafName returns the LeafName field if non-nil, zero value otherwise.

### GetLeafNameOk

`func (o *MachineConfigurationResponseModel) GetLeafNameOk() (*string, bool)`

GetLeafNameOk returns a tuple with the LeafName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafName

`func (o *MachineConfigurationResponseModel) SetLeafName(v string)`

SetLeafName sets LeafName field to given value.

### HasLeafName

`func (o *MachineConfigurationResponseModel) HasLeafName() bool`

HasLeafName returns a boolean if a field has been set.

### SetLeafNameNil

`func (o *MachineConfigurationResponseModel) SetLeafNameNil(b bool)`

 SetLeafNameNil sets the value for LeafName to be an explicit nil

### UnsetLeafName
`func (o *MachineConfigurationResponseModel) UnsetLeafName()`

UnsetLeafName ensures that no value is present for LeafName, not even an explicit nil
### GetPolicy

`func (o *MachineConfigurationResponseModel) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *MachineConfigurationResponseModel) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *MachineConfigurationResponseModel) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *MachineConfigurationResponseModel) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *MachineConfigurationResponseModel) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *MachineConfigurationResponseModel) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetDescription

`func (o *MachineConfigurationResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MachineConfigurationResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MachineConfigurationResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MachineConfigurationResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MachineConfigurationResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MachineConfigurationResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDeliveryGroups

`func (o *MachineConfigurationResponseModel) GetDeliveryGroups() []RefResponseModel`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *MachineConfigurationResponseModel) GetDeliveryGroupsOk() (*[]RefResponseModel, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *MachineConfigurationResponseModel) SetDeliveryGroups(v []RefResponseModel)`

SetDeliveryGroups sets DeliveryGroups field to given value.

### HasDeliveryGroups

`func (o *MachineConfigurationResponseModel) HasDeliveryGroups() bool`

HasDeliveryGroups returns a boolean if a field has been set.

### SetDeliveryGroupsNil

`func (o *MachineConfigurationResponseModel) SetDeliveryGroupsNil(b bool)`

 SetDeliveryGroupsNil sets the value for DeliveryGroups to be an explicit nil

### UnsetDeliveryGroups
`func (o *MachineConfigurationResponseModel) UnsetDeliveryGroups()`

UnsetDeliveryGroups ensures that no value is present for DeliveryGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


