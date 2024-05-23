# ImageVersionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Id of the image version. | 
**Number** | **int32** | The version number associated with the image version. | 
**CreationTime** | Pointer to **NullableString** | Time when the image version was created. | [optional] 
**Description** | Pointer to **NullableString** | The Description of the image version. | [optional] 
**ImageVersionStatus** | Pointer to [**ImageVersionStatus**](ImageVersionStatus.md) |  | [optional] 
**ImageDefinition** | [**RefResponseModel**](RefResponseModel.md) |  | 
**ImageVersionSpecs** | Pointer to [**[]ImageVersionSpecResponseModel**](ImageVersionSpecResponseModel.md) | The image version specifications associated with this image version. | [optional] 
**ProvisioningSchemeCount** | Pointer to **int32** | The count of provisioning schemes | [optional] 

## Methods

### NewImageVersionResponseModel

`func NewImageVersionResponseModel(id string, number int32, imageDefinition RefResponseModel, ) *ImageVersionResponseModel`

NewImageVersionResponseModel instantiates a new ImageVersionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionResponseModelWithDefaults

`func NewImageVersionResponseModelWithDefaults() *ImageVersionResponseModel`

NewImageVersionResponseModelWithDefaults instantiates a new ImageVersionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageVersionResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *ImageVersionResponseModel) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ImageVersionResponseModel) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ImageVersionResponseModel) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetCreationTime

`func (o *ImageVersionResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ImageVersionResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ImageVersionResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ImageVersionResponseModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### SetCreationTimeNil

`func (o *ImageVersionResponseModel) SetCreationTimeNil(b bool)`

 SetCreationTimeNil sets the value for CreationTime to be an explicit nil

### UnsetCreationTime
`func (o *ImageVersionResponseModel) UnsetCreationTime()`

UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
### GetDescription

`func (o *ImageVersionResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageVersionResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageVersionResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageVersionResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ImageVersionResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ImageVersionResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageVersionStatus

`func (o *ImageVersionResponseModel) GetImageVersionStatus() ImageVersionStatus`

GetImageVersionStatus returns the ImageVersionStatus field if non-nil, zero value otherwise.

### GetImageVersionStatusOk

`func (o *ImageVersionResponseModel) GetImageVersionStatusOk() (*ImageVersionStatus, bool)`

GetImageVersionStatusOk returns a tuple with the ImageVersionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersionStatus

`func (o *ImageVersionResponseModel) SetImageVersionStatus(v ImageVersionStatus)`

SetImageVersionStatus sets ImageVersionStatus field to given value.

### HasImageVersionStatus

`func (o *ImageVersionResponseModel) HasImageVersionStatus() bool`

HasImageVersionStatus returns a boolean if a field has been set.

### GetImageDefinition

`func (o *ImageVersionResponseModel) GetImageDefinition() RefResponseModel`

GetImageDefinition returns the ImageDefinition field if non-nil, zero value otherwise.

### GetImageDefinitionOk

`func (o *ImageVersionResponseModel) GetImageDefinitionOk() (*RefResponseModel, bool)`

GetImageDefinitionOk returns a tuple with the ImageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinition

`func (o *ImageVersionResponseModel) SetImageDefinition(v RefResponseModel)`

SetImageDefinition sets ImageDefinition field to given value.


### GetImageVersionSpecs

`func (o *ImageVersionResponseModel) GetImageVersionSpecs() []ImageVersionSpecResponseModel`

GetImageVersionSpecs returns the ImageVersionSpecs field if non-nil, zero value otherwise.

### GetImageVersionSpecsOk

`func (o *ImageVersionResponseModel) GetImageVersionSpecsOk() (*[]ImageVersionSpecResponseModel, bool)`

GetImageVersionSpecsOk returns a tuple with the ImageVersionSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersionSpecs

`func (o *ImageVersionResponseModel) SetImageVersionSpecs(v []ImageVersionSpecResponseModel)`

SetImageVersionSpecs sets ImageVersionSpecs field to given value.

### HasImageVersionSpecs

`func (o *ImageVersionResponseModel) HasImageVersionSpecs() bool`

HasImageVersionSpecs returns a boolean if a field has been set.

### SetImageVersionSpecsNil

`func (o *ImageVersionResponseModel) SetImageVersionSpecsNil(b bool)`

 SetImageVersionSpecsNil sets the value for ImageVersionSpecs to be an explicit nil

### UnsetImageVersionSpecs
`func (o *ImageVersionResponseModel) UnsetImageVersionSpecs()`

UnsetImageVersionSpecs ensures that no value is present for ImageVersionSpecs, not even an explicit nil
### GetProvisioningSchemeCount

`func (o *ImageVersionResponseModel) GetProvisioningSchemeCount() int32`

GetProvisioningSchemeCount returns the ProvisioningSchemeCount field if non-nil, zero value otherwise.

### GetProvisioningSchemeCountOk

`func (o *ImageVersionResponseModel) GetProvisioningSchemeCountOk() (*int32, bool)`

GetProvisioningSchemeCountOk returns a tuple with the ProvisioningSchemeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeCount

`func (o *ImageVersionResponseModel) SetProvisioningSchemeCount(v int32)`

SetProvisioningSchemeCount sets ProvisioningSchemeCount field to given value.

### HasProvisioningSchemeCount

`func (o *ImageVersionResponseModel) HasProvisioningSchemeCount() bool`

HasProvisioningSchemeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


