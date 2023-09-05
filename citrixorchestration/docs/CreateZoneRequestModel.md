# CreateZoneRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the zone. Each zone within a site must have a unique name.  Name is the only required property in creating a new zone. | 
**Description** | Pointer to **string** | Description of the zone. | [optional] 
**ResourceLocation** | Pointer to **string** | Resource location associated with the zone. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of the Zone. | [optional] 

## Methods

### NewCreateZoneRequestModel

`func NewCreateZoneRequestModel(name string, ) *CreateZoneRequestModel`

NewCreateZoneRequestModel instantiates a new CreateZoneRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateZoneRequestModelWithDefaults

`func NewCreateZoneRequestModelWithDefaults() *CreateZoneRequestModel`

NewCreateZoneRequestModelWithDefaults instantiates a new CreateZoneRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateZoneRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateZoneRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateZoneRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateZoneRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateZoneRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateZoneRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateZoneRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceLocation

`func (o *CreateZoneRequestModel) GetResourceLocation() string`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *CreateZoneRequestModel) GetResourceLocationOk() (*string, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocation

`func (o *CreateZoneRequestModel) SetResourceLocation(v string)`

SetResourceLocation sets ResourceLocation field to given value.

### HasResourceLocation

`func (o *CreateZoneRequestModel) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateZoneRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateZoneRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateZoneRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateZoneRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


