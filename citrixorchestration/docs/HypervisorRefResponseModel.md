# HypervisorRefResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **NullableInt32** | DEPRECATED. Use Id. | [optional] 
**Name** | Pointer to **NullableString** | Name of the object. | [optional] 
**ConnectionType** | Pointer to [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | [optional] 
**PluginFactoryName** | Pointer to **NullableString** | The class name for the Citrix Managed Machine library that is used to access the hypervisor. | [optional] 

## Methods

### NewHypervisorRefResponseModel

`func NewHypervisorRefResponseModel() *HypervisorRefResponseModel`

NewHypervisorRefResponseModel instantiates a new HypervisorRefResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorRefResponseModelWithDefaults

`func NewHypervisorRefResponseModelWithDefaults() *HypervisorRefResponseModel`

NewHypervisorRefResponseModelWithDefaults instantiates a new HypervisorRefResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorRefResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorRefResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorRefResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorRefResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorRefResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorRefResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *HypervisorRefResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *HypervisorRefResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *HypervisorRefResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *HypervisorRefResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *HypervisorRefResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *HypervisorRefResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetName

`func (o *HypervisorRefResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorRefResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorRefResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorRefResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorRefResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorRefResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConnectionType

`func (o *HypervisorRefResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorRefResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorRefResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *HypervisorRefResponseModel) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetPluginFactoryName

`func (o *HypervisorRefResponseModel) GetPluginFactoryName() string`

GetPluginFactoryName returns the PluginFactoryName field if non-nil, zero value otherwise.

### GetPluginFactoryNameOk

`func (o *HypervisorRefResponseModel) GetPluginFactoryNameOk() (*string, bool)`

GetPluginFactoryNameOk returns a tuple with the PluginFactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginFactoryName

`func (o *HypervisorRefResponseModel) SetPluginFactoryName(v string)`

SetPluginFactoryName sets PluginFactoryName field to given value.

### HasPluginFactoryName

`func (o *HypervisorRefResponseModel) HasPluginFactoryName() bool`

HasPluginFactoryName returns a boolean if a field has been set.

### SetPluginFactoryNameNil

`func (o *HypervisorRefResponseModel) SetPluginFactoryNameNil(b bool)`

 SetPluginFactoryNameNil sets the value for PluginFactoryName to be an explicit nil

### UnsetPluginFactoryName
`func (o *HypervisorRefResponseModel) UnsetPluginFactoryName()`

UnsetPluginFactoryName ensures that no value is present for PluginFactoryName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


