# HypervisorTemplateResourceResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **string** | The account ID for the owner of this template. | [optional] 
**IsWindowsTemplate** | Pointer to **bool** | Indicates whether this is a Windows OS template, if known. | [optional] 
**Description** | Pointer to **string** | Human-readable description of this template, as supplied when the offering was created in the cloud management API or console. | [optional] 
**HasPersistentRootVolume** | **bool** | Indicates whether this template has a persistent root volume (eg. is an EBS-backed image on AWS). | 

## Methods

### NewHypervisorTemplateResourceResponseModelAllOf

`func NewHypervisorTemplateResourceResponseModelAllOf(hasPersistentRootVolume bool, ) *HypervisorTemplateResourceResponseModelAllOf`

NewHypervisorTemplateResourceResponseModelAllOf instantiates a new HypervisorTemplateResourceResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorTemplateResourceResponseModelAllOfWithDefaults

`func NewHypervisorTemplateResourceResponseModelAllOfWithDefaults() *HypervisorTemplateResourceResponseModelAllOf`

NewHypervisorTemplateResourceResponseModelAllOfWithDefaults instantiates a new HypervisorTemplateResourceResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *HypervisorTemplateResourceResponseModelAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HypervisorTemplateResourceResponseModelAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HypervisorTemplateResourceResponseModelAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *HypervisorTemplateResourceResponseModelAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetIsWindowsTemplate

`func (o *HypervisorTemplateResourceResponseModelAllOf) GetIsWindowsTemplate() bool`

GetIsWindowsTemplate returns the IsWindowsTemplate field if non-nil, zero value otherwise.

### GetIsWindowsTemplateOk

`func (o *HypervisorTemplateResourceResponseModelAllOf) GetIsWindowsTemplateOk() (*bool, bool)`

GetIsWindowsTemplateOk returns a tuple with the IsWindowsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindowsTemplate

`func (o *HypervisorTemplateResourceResponseModelAllOf) SetIsWindowsTemplate(v bool)`

SetIsWindowsTemplate sets IsWindowsTemplate field to given value.

### HasIsWindowsTemplate

`func (o *HypervisorTemplateResourceResponseModelAllOf) HasIsWindowsTemplate() bool`

HasIsWindowsTemplate returns a boolean if a field has been set.

### GetDescription

`func (o *HypervisorTemplateResourceResponseModelAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HypervisorTemplateResourceResponseModelAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HypervisorTemplateResourceResponseModelAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HypervisorTemplateResourceResponseModelAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasPersistentRootVolume

`func (o *HypervisorTemplateResourceResponseModelAllOf) GetHasPersistentRootVolume() bool`

GetHasPersistentRootVolume returns the HasPersistentRootVolume field if non-nil, zero value otherwise.

### GetHasPersistentRootVolumeOk

`func (o *HypervisorTemplateResourceResponseModelAllOf) GetHasPersistentRootVolumeOk() (*bool, bool)`

GetHasPersistentRootVolumeOk returns a tuple with the HasPersistentRootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPersistentRootVolume

`func (o *HypervisorTemplateResourceResponseModelAllOf) SetHasPersistentRootVolume(v bool)`

SetHasPersistentRootVolume sets HasPersistentRootVolume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


