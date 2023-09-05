# AttachedDiskResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the attached disk. | [optional] 
**IsBootDisk** | Pointer to **bool** | Indicates whether the disk is the boot disk of the VM. | [optional] 
**SizeGB** | Pointer to **int32** | Size of the disk, in whole GB. | [optional] 
**InfrastructureUid** | Pointer to **string** | Opaque identifier used by the underlying infrastructure to identify the disk. | [optional] 

## Methods

### NewAttachedDiskResponseModel

`func NewAttachedDiskResponseModel() *AttachedDiskResponseModel`

NewAttachedDiskResponseModel instantiates a new AttachedDiskResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachedDiskResponseModelWithDefaults

`func NewAttachedDiskResponseModelWithDefaults() *AttachedDiskResponseModel`

NewAttachedDiskResponseModelWithDefaults instantiates a new AttachedDiskResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AttachedDiskResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachedDiskResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachedDiskResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttachedDiskResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsBootDisk

`func (o *AttachedDiskResponseModel) GetIsBootDisk() bool`

GetIsBootDisk returns the IsBootDisk field if non-nil, zero value otherwise.

### GetIsBootDiskOk

`func (o *AttachedDiskResponseModel) GetIsBootDiskOk() (*bool, bool)`

GetIsBootDiskOk returns a tuple with the IsBootDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootDisk

`func (o *AttachedDiskResponseModel) SetIsBootDisk(v bool)`

SetIsBootDisk sets IsBootDisk field to given value.

### HasIsBootDisk

`func (o *AttachedDiskResponseModel) HasIsBootDisk() bool`

HasIsBootDisk returns a boolean if a field has been set.

### GetSizeGB

`func (o *AttachedDiskResponseModel) GetSizeGB() int32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *AttachedDiskResponseModel) GetSizeGBOk() (*int32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *AttachedDiskResponseModel) SetSizeGB(v int32)`

SetSizeGB sets SizeGB field to given value.

### HasSizeGB

`func (o *AttachedDiskResponseModel) HasSizeGB() bool`

HasSizeGB returns a boolean if a field has been set.

### GetInfrastructureUid

`func (o *AttachedDiskResponseModel) GetInfrastructureUid() string`

GetInfrastructureUid returns the InfrastructureUid field if non-nil, zero value otherwise.

### GetInfrastructureUidOk

`func (o *AttachedDiskResponseModel) GetInfrastructureUidOk() (*string, bool)`

GetInfrastructureUidOk returns a tuple with the InfrastructureUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureUid

`func (o *AttachedDiskResponseModel) SetInfrastructureUid(v string)`

SetInfrastructureUid sets InfrastructureUid field to given value.

### HasInfrastructureUid

`func (o *AttachedDiskResponseModel) HasInfrastructureUid() bool`

HasInfrastructureUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


