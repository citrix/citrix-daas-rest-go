# ImageVersionDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDefinition** | Pointer to [**ImageVersionDetailResponseModelAllOfImageDefinition**](ImageVersionDetailResponseModelAllOfImageDefinition.md) |  | [optional] 
**ImageScheme** | Pointer to [**ImageVersionDetailResponseModelAllOfImageScheme**](ImageVersionDetailResponseModelAllOfImageScheme.md) |  | [optional] 
**VMMetadata** | Pointer to [**ImageVersionDetailResponseModelAllOfVMMetadata**](ImageVersionDetailResponseModelAllOfVMMetadata.md) |  | [optional] 
**MasterImage** | [**ImageVersionDetailResponseModelAllOfMasterImage**](ImageVersionDetailResponseModelAllOfMasterImage.md) |  | 

## Methods

### NewImageVersionDetailResponseModelAllOf

`func NewImageVersionDetailResponseModelAllOf(masterImage ImageVersionDetailResponseModelAllOfMasterImage, ) *ImageVersionDetailResponseModelAllOf`

NewImageVersionDetailResponseModelAllOf instantiates a new ImageVersionDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionDetailResponseModelAllOfWithDefaults

`func NewImageVersionDetailResponseModelAllOfWithDefaults() *ImageVersionDetailResponseModelAllOf`

NewImageVersionDetailResponseModelAllOfWithDefaults instantiates a new ImageVersionDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDefinition

`func (o *ImageVersionDetailResponseModelAllOf) GetImageDefinition() ImageVersionDetailResponseModelAllOfImageDefinition`

GetImageDefinition returns the ImageDefinition field if non-nil, zero value otherwise.

### GetImageDefinitionOk

`func (o *ImageVersionDetailResponseModelAllOf) GetImageDefinitionOk() (*ImageVersionDetailResponseModelAllOfImageDefinition, bool)`

GetImageDefinitionOk returns a tuple with the ImageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinition

`func (o *ImageVersionDetailResponseModelAllOf) SetImageDefinition(v ImageVersionDetailResponseModelAllOfImageDefinition)`

SetImageDefinition sets ImageDefinition field to given value.

### HasImageDefinition

`func (o *ImageVersionDetailResponseModelAllOf) HasImageDefinition() bool`

HasImageDefinition returns a boolean if a field has been set.

### GetImageScheme

`func (o *ImageVersionDetailResponseModelAllOf) GetImageScheme() ImageVersionDetailResponseModelAllOfImageScheme`

GetImageScheme returns the ImageScheme field if non-nil, zero value otherwise.

### GetImageSchemeOk

`func (o *ImageVersionDetailResponseModelAllOf) GetImageSchemeOk() (*ImageVersionDetailResponseModelAllOfImageScheme, bool)`

GetImageSchemeOk returns a tuple with the ImageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScheme

`func (o *ImageVersionDetailResponseModelAllOf) SetImageScheme(v ImageVersionDetailResponseModelAllOfImageScheme)`

SetImageScheme sets ImageScheme field to given value.

### HasImageScheme

`func (o *ImageVersionDetailResponseModelAllOf) HasImageScheme() bool`

HasImageScheme returns a boolean if a field has been set.

### GetVMMetadata

`func (o *ImageVersionDetailResponseModelAllOf) GetVMMetadata() ImageVersionDetailResponseModelAllOfVMMetadata`

GetVMMetadata returns the VMMetadata field if non-nil, zero value otherwise.

### GetVMMetadataOk

`func (o *ImageVersionDetailResponseModelAllOf) GetVMMetadataOk() (*ImageVersionDetailResponseModelAllOfVMMetadata, bool)`

GetVMMetadataOk returns a tuple with the VMMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMMetadata

`func (o *ImageVersionDetailResponseModelAllOf) SetVMMetadata(v ImageVersionDetailResponseModelAllOfVMMetadata)`

SetVMMetadata sets VMMetadata field to given value.

### HasVMMetadata

`func (o *ImageVersionDetailResponseModelAllOf) HasVMMetadata() bool`

HasVMMetadata returns a boolean if a field has been set.

### GetMasterImage

`func (o *ImageVersionDetailResponseModelAllOf) GetMasterImage() ImageVersionDetailResponseModelAllOfMasterImage`

GetMasterImage returns the MasterImage field if non-nil, zero value otherwise.

### GetMasterImageOk

`func (o *ImageVersionDetailResponseModelAllOf) GetMasterImageOk() (*ImageVersionDetailResponseModelAllOfMasterImage, bool)`

GetMasterImageOk returns a tuple with the MasterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImage

`func (o *ImageVersionDetailResponseModelAllOf) SetMasterImage(v ImageVersionDetailResponseModelAllOfMasterImage)`

SetMasterImage sets MasterImage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


