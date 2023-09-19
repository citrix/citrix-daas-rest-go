# ContainerMetadataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) |  | [optional] 
**ContainerType** | Pointer to [**ContainerMetadataType**](ContainerMetadataType.md) |  | [optional] 

## Methods

### NewContainerMetadataModel

`func NewContainerMetadataModel() *ContainerMetadataModel`

NewContainerMetadataModel instantiates a new ContainerMetadataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerMetadataModelWithDefaults

`func NewContainerMetadataModelWithDefaults() *ContainerMetadataModel`

NewContainerMetadataModelWithDefaults instantiates a new ContainerMetadataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ContainerMetadataModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContainerMetadataModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContainerMetadataModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ContainerMetadataModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ContainerMetadataModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ContainerMetadataModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetContainerType

`func (o *ContainerMetadataModel) GetContainerType() ContainerMetadataType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ContainerMetadataModel) GetContainerTypeOk() (*ContainerMetadataType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ContainerMetadataModel) SetContainerType(v ContainerMetadataType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *ContainerMetadataModel) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


