# ImageVersionSpecContextResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Additional data about the object in the form of key-value pairs. | [optional] 
**ImageScheme** | Pointer to [**ImageSchemeResponseModel**](ImageSchemeResponseModel.md) |  | [optional] 
**MachineProfileMetadata** | Pointer to [**MachineProfileMetadataResponseModel**](MachineProfileMetadataResponseModel.md) |  | [optional] 
**PluginFactoryName** | Pointer to **NullableString** | Custom hypervisor plugin. | [optional] 
**MachineAgentOptions** | Pointer to [**MachineAgentOptions**](MachineAgentOptions.md) |  | [optional] 

## Methods

### NewImageVersionSpecContextResponseModel

`func NewImageVersionSpecContextResponseModel() *ImageVersionSpecContextResponseModel`

NewImageVersionSpecContextResponseModel instantiates a new ImageVersionSpecContextResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionSpecContextResponseModelWithDefaults

`func NewImageVersionSpecContextResponseModelWithDefaults() *ImageVersionSpecContextResponseModel`

NewImageVersionSpecContextResponseModelWithDefaults instantiates a new ImageVersionSpecContextResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ImageVersionSpecContextResponseModel) GetAdditionalData() []NameValueStringPairModel`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ImageVersionSpecContextResponseModel) GetAdditionalDataOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ImageVersionSpecContextResponseModel) SetAdditionalData(v []NameValueStringPairModel)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ImageVersionSpecContextResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *ImageVersionSpecContextResponseModel) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *ImageVersionSpecContextResponseModel) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil
### GetImageScheme

`func (o *ImageVersionSpecContextResponseModel) GetImageScheme() ImageSchemeResponseModel`

GetImageScheme returns the ImageScheme field if non-nil, zero value otherwise.

### GetImageSchemeOk

`func (o *ImageVersionSpecContextResponseModel) GetImageSchemeOk() (*ImageSchemeResponseModel, bool)`

GetImageSchemeOk returns a tuple with the ImageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScheme

`func (o *ImageVersionSpecContextResponseModel) SetImageScheme(v ImageSchemeResponseModel)`

SetImageScheme sets ImageScheme field to given value.

### HasImageScheme

`func (o *ImageVersionSpecContextResponseModel) HasImageScheme() bool`

HasImageScheme returns a boolean if a field has been set.

### GetMachineProfileMetadata

`func (o *ImageVersionSpecContextResponseModel) GetMachineProfileMetadata() MachineProfileMetadataResponseModel`

GetMachineProfileMetadata returns the MachineProfileMetadata field if non-nil, zero value otherwise.

### GetMachineProfileMetadataOk

`func (o *ImageVersionSpecContextResponseModel) GetMachineProfileMetadataOk() (*MachineProfileMetadataResponseModel, bool)`

GetMachineProfileMetadataOk returns a tuple with the MachineProfileMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfileMetadata

`func (o *ImageVersionSpecContextResponseModel) SetMachineProfileMetadata(v MachineProfileMetadataResponseModel)`

SetMachineProfileMetadata sets MachineProfileMetadata field to given value.

### HasMachineProfileMetadata

`func (o *ImageVersionSpecContextResponseModel) HasMachineProfileMetadata() bool`

HasMachineProfileMetadata returns a boolean if a field has been set.

### GetPluginFactoryName

`func (o *ImageVersionSpecContextResponseModel) GetPluginFactoryName() string`

GetPluginFactoryName returns the PluginFactoryName field if non-nil, zero value otherwise.

### GetPluginFactoryNameOk

`func (o *ImageVersionSpecContextResponseModel) GetPluginFactoryNameOk() (*string, bool)`

GetPluginFactoryNameOk returns a tuple with the PluginFactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginFactoryName

`func (o *ImageVersionSpecContextResponseModel) SetPluginFactoryName(v string)`

SetPluginFactoryName sets PluginFactoryName field to given value.

### HasPluginFactoryName

`func (o *ImageVersionSpecContextResponseModel) HasPluginFactoryName() bool`

HasPluginFactoryName returns a boolean if a field has been set.

### SetPluginFactoryNameNil

`func (o *ImageVersionSpecContextResponseModel) SetPluginFactoryNameNil(b bool)`

 SetPluginFactoryNameNil sets the value for PluginFactoryName to be an explicit nil

### UnsetPluginFactoryName
`func (o *ImageVersionSpecContextResponseModel) UnsetPluginFactoryName()`

UnsetPluginFactoryName ensures that no value is present for PluginFactoryName, not even an explicit nil
### GetMachineAgentOptions

`func (o *ImageVersionSpecContextResponseModel) GetMachineAgentOptions() MachineAgentOptions`

GetMachineAgentOptions returns the MachineAgentOptions field if non-nil, zero value otherwise.

### GetMachineAgentOptionsOk

`func (o *ImageVersionSpecContextResponseModel) GetMachineAgentOptionsOk() (*MachineAgentOptions, bool)`

GetMachineAgentOptionsOk returns a tuple with the MachineAgentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAgentOptions

`func (o *ImageVersionSpecContextResponseModel) SetMachineAgentOptions(v MachineAgentOptions)`

SetMachineAgentOptions sets MachineAgentOptions field to given value.

### HasMachineAgentOptions

`func (o *ImageVersionSpecContextResponseModel) HasMachineAgentOptions() bool`

HasMachineAgentOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


