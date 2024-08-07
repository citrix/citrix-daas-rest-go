# ImageDefinitionHypervisorConnectionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **NullableInt32** | DEPRECATED. Use Id. | [optional] 
**Name** | Pointer to **NullableString** | Name of the object. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image definition that are specific to the target hypervisor connection. | [optional] 
**HypervisorData** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Hypervisor data about the object in the form of key-value pairs. TODO AdditionalData or HypervisorData? | [optional] 
**PluginFactoryName** | Pointer to **NullableString** | The connection plugin factory name. | [optional] 

## Methods

### NewImageDefinitionHypervisorConnectionResponseModel

`func NewImageDefinitionHypervisorConnectionResponseModel() *ImageDefinitionHypervisorConnectionResponseModel`

NewImageDefinitionHypervisorConnectionResponseModel instantiates a new ImageDefinitionHypervisorConnectionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDefinitionHypervisorConnectionResponseModelWithDefaults

`func NewImageDefinitionHypervisorConnectionResponseModelWithDefaults() *ImageDefinitionHypervisorConnectionResponseModel`

NewImageDefinitionHypervisorConnectionResponseModelWithDefaults instantiates a new ImageDefinitionHypervisorConnectionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageDefinitionHypervisorConnectionResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ImageDefinitionHypervisorConnectionResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ImageDefinitionHypervisorConnectionResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ImageDefinitionHypervisorConnectionResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetName

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageDefinitionHypervisorConnectionResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImageDefinitionHypervisorConnectionResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCustomProperties

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ImageDefinitionHypervisorConnectionResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *ImageDefinitionHypervisorConnectionResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetHypervisorData

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetHypervisorData() []NameValueStringPairModel`

GetHypervisorData returns the HypervisorData field if non-nil, zero value otherwise.

### GetHypervisorDataOk

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetHypervisorDataOk() (*[]NameValueStringPairModel, bool)`

GetHypervisorDataOk returns a tuple with the HypervisorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorData

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetHypervisorData(v []NameValueStringPairModel)`

SetHypervisorData sets HypervisorData field to given value.

### HasHypervisorData

`func (o *ImageDefinitionHypervisorConnectionResponseModel) HasHypervisorData() bool`

HasHypervisorData returns a boolean if a field has been set.

### SetHypervisorDataNil

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetHypervisorDataNil(b bool)`

 SetHypervisorDataNil sets the value for HypervisorData to be an explicit nil

### UnsetHypervisorData
`func (o *ImageDefinitionHypervisorConnectionResponseModel) UnsetHypervisorData()`

UnsetHypervisorData ensures that no value is present for HypervisorData, not even an explicit nil
### GetPluginFactoryName

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetPluginFactoryName() string`

GetPluginFactoryName returns the PluginFactoryName field if non-nil, zero value otherwise.

### GetPluginFactoryNameOk

`func (o *ImageDefinitionHypervisorConnectionResponseModel) GetPluginFactoryNameOk() (*string, bool)`

GetPluginFactoryNameOk returns a tuple with the PluginFactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginFactoryName

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetPluginFactoryName(v string)`

SetPluginFactoryName sets PluginFactoryName field to given value.

### HasPluginFactoryName

`func (o *ImageDefinitionHypervisorConnectionResponseModel) HasPluginFactoryName() bool`

HasPluginFactoryName returns a boolean if a field has been set.

### SetPluginFactoryNameNil

`func (o *ImageDefinitionHypervisorConnectionResponseModel) SetPluginFactoryNameNil(b bool)`

 SetPluginFactoryNameNil sets the value for PluginFactoryName to be an explicit nil

### UnsetPluginFactoryName
`func (o *ImageDefinitionHypervisorConnectionResponseModel) UnsetPluginFactoryName()`

UnsetPluginFactoryName ensures that no value is present for PluginFactoryName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


