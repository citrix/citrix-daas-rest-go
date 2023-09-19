# HypervisorTemplateResourceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the resource. | [optional] 
**Name** | Pointer to **NullableString** | Name of the resource. | [optional] 
**XDPath** | Pointer to **NullableString** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**RelativePath** | Pointer to **NullableString** | Path to the object, relative to the resource pool from which it was queried. &#x60;{{vm name}}.vm/{{snapshot name}}.snapshot&#x60; | [optional] 
**FullRelativePath** | Pointer to **NullableString** | Full path to the resource, including the hypervisor, relative to the root of the API. Example: &#x60;Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources/{{RelativePath}}&#x60; | [optional] 
**ResourceType** | **string** | Type of resource. | 
**ObjectTypeName** | Pointer to **NullableString** | The type name of the hypervisor resource object. | [optional] 
**ResourceContainer** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**ResourcePoolXDPath** | Pointer to **NullableString** | Citrix Apps and Desktops path to the resource on the ResourcePool.  An example value is: &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; This value  | [optional] 
**FullName** | **string** | Name of the resource, with the type concatenated. i.e. \&quot;name.type\&quot;. | 
**IsContainer** | **bool** | Indicates whether the resource is a container. | 
**Children** | Pointer to [**[]HypervisorResourceResponseModel**](HypervisorResourceResponseModel.md) | Child resources, when the resource is a container. | [optional] 
**IsMachine** | **bool** | Indicates whether the resource is a machine. | 
**IsSnapshotable** | **bool** | Indicates whether the resource can have a snapshot taken. | 
**AllResourcesRelativePath** | **string** | Path to the resource, relative to the special \&quot;AllResources\&quot; resource pool associated with the hypervisor. | 
**ResourcePool** | [**HypervisorResourcePoolRefResponseModel**](HypervisorResourcePoolRefResponseModel.md) |  | 
**IsSymLink** | **bool** | Indicates whether the object is a valid symbolic link. | 
**AdditionalData** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Additional data about the object in the form of key-value pairs. | [optional] 
**CpuCount** | Pointer to **NullableInt32** | Number of CPUs, if known. | [optional] 
**MemoryMB** | Pointer to **NullableInt32** | Memory in megabytes, if known. | [optional] 
**HardDiskSizeGB** | Pointer to **NullableInt32** | Hard disk size in gigabytes, if known. | [optional] 
**MinMemoryMB** | Pointer to **NullableInt32** | Minimum memory required to run this VM or snapshot, in megabytes, if known. | [optional] 
**NetworkMappings** | Pointer to [**[]NetworkMapResponseModel**](NetworkMapResponseModel.md) | Network mappings associated with the VM, if known. | [optional] 
**AttachedDisks** | Pointer to [**[]AttachedDiskResponseModel**](AttachedDiskResponseModel.md) | Disks attached to the VM, if known. | [optional] 
**Owner** | Pointer to **NullableString** | The account ID for the owner of this template. | [optional] 
**IsWindowsTemplate** | Pointer to **NullableBool** | Indicates whether this is a Windows OS template, if known. | [optional] 
**Description** | Pointer to **NullableString** | Human-readable description of this template, as supplied when the offering was created in the cloud management API or console. | [optional] 
**HasPersistentRootVolume** | **bool** | Indicates whether this template has a persistent root volume (eg. is an EBS-backed image on AWS). | 

## Methods

### NewHypervisorTemplateResourceResponseModel

`func NewHypervisorTemplateResourceResponseModel(resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourcePoolRefResponseModel, isSymLink bool, hasPersistentRootVolume bool, ) *HypervisorTemplateResourceResponseModel`

NewHypervisorTemplateResourceResponseModel instantiates a new HypervisorTemplateResourceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorTemplateResourceResponseModelWithDefaults

`func NewHypervisorTemplateResourceResponseModelWithDefaults() *HypervisorTemplateResourceResponseModel`

NewHypervisorTemplateResourceResponseModelWithDefaults instantiates a new HypervisorTemplateResourceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorTemplateResourceResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorTemplateResourceResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorTemplateResourceResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorTemplateResourceResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorTemplateResourceResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorTemplateResourceResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorTemplateResourceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorTemplateResourceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorTemplateResourceResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorTemplateResourceResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorTemplateResourceResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorTemplateResourceResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorTemplateResourceResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorTemplateResourceResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorTemplateResourceResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorTemplateResourceResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorTemplateResourceResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorTemplateResourceResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetRelativePath

`func (o *HypervisorTemplateResourceResponseModel) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HypervisorTemplateResourceResponseModel) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HypervisorTemplateResourceResponseModel) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *HypervisorTemplateResourceResponseModel) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *HypervisorTemplateResourceResponseModel) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *HypervisorTemplateResourceResponseModel) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
### GetFullRelativePath

`func (o *HypervisorTemplateResourceResponseModel) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorTemplateResourceResponseModel) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorTemplateResourceResponseModel) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.

### HasFullRelativePath

`func (o *HypervisorTemplateResourceResponseModel) HasFullRelativePath() bool`

HasFullRelativePath returns a boolean if a field has been set.

### SetFullRelativePathNil

`func (o *HypervisorTemplateResourceResponseModel) SetFullRelativePathNil(b bool)`

 SetFullRelativePathNil sets the value for FullRelativePath to be an explicit nil

### UnsetFullRelativePath
`func (o *HypervisorTemplateResourceResponseModel) UnsetFullRelativePath()`

UnsetFullRelativePath ensures that no value is present for FullRelativePath, not even an explicit nil
### GetResourceType

`func (o *HypervisorTemplateResourceResponseModel) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorTemplateResourceResponseModel) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorTemplateResourceResponseModel) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetObjectTypeName

`func (o *HypervisorTemplateResourceResponseModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *HypervisorTemplateResourceResponseModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *HypervisorTemplateResourceResponseModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *HypervisorTemplateResourceResponseModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### SetObjectTypeNameNil

`func (o *HypervisorTemplateResourceResponseModel) SetObjectTypeNameNil(b bool)`

 SetObjectTypeNameNil sets the value for ObjectTypeName to be an explicit nil

### UnsetObjectTypeName
`func (o *HypervisorTemplateResourceResponseModel) UnsetObjectTypeName()`

UnsetObjectTypeName ensures that no value is present for ObjectTypeName, not even an explicit nil
### GetResourceContainer

`func (o *HypervisorTemplateResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModel`

GetResourceContainer returns the ResourceContainer field if non-nil, zero value otherwise.

### GetResourceContainerOk

`func (o *HypervisorTemplateResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModel, bool)`

GetResourceContainerOk returns a tuple with the ResourceContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceContainer

`func (o *HypervisorTemplateResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModel)`

SetResourceContainer sets ResourceContainer field to given value.

### HasResourceContainer

`func (o *HypervisorTemplateResourceResponseModel) HasResourceContainer() bool`

HasResourceContainer returns a boolean if a field has been set.

### GetResourcePoolXDPath

`func (o *HypervisorTemplateResourceResponseModel) GetResourcePoolXDPath() string`

GetResourcePoolXDPath returns the ResourcePoolXDPath field if non-nil, zero value otherwise.

### GetResourcePoolXDPathOk

`func (o *HypervisorTemplateResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool)`

GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolXDPath

`func (o *HypervisorTemplateResourceResponseModel) SetResourcePoolXDPath(v string)`

SetResourcePoolXDPath sets ResourcePoolXDPath field to given value.

### HasResourcePoolXDPath

`func (o *HypervisorTemplateResourceResponseModel) HasResourcePoolXDPath() bool`

HasResourcePoolXDPath returns a boolean if a field has been set.

### SetResourcePoolXDPathNil

`func (o *HypervisorTemplateResourceResponseModel) SetResourcePoolXDPathNil(b bool)`

 SetResourcePoolXDPathNil sets the value for ResourcePoolXDPath to be an explicit nil

### UnsetResourcePoolXDPath
`func (o *HypervisorTemplateResourceResponseModel) UnsetResourcePoolXDPath()`

UnsetResourcePoolXDPath ensures that no value is present for ResourcePoolXDPath, not even an explicit nil
### GetFullName

`func (o *HypervisorTemplateResourceResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *HypervisorTemplateResourceResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *HypervisorTemplateResourceResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsContainer

`func (o *HypervisorTemplateResourceResponseModel) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *HypervisorTemplateResourceResponseModel) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *HypervisorTemplateResourceResponseModel) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.


### GetChildren

`func (o *HypervisorTemplateResourceResponseModel) GetChildren() []HypervisorResourceResponseModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HypervisorTemplateResourceResponseModel) GetChildrenOk() (*[]HypervisorResourceResponseModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HypervisorTemplateResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HypervisorTemplateResourceResponseModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *HypervisorTemplateResourceResponseModel) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *HypervisorTemplateResourceResponseModel) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetIsMachine

`func (o *HypervisorTemplateResourceResponseModel) GetIsMachine() bool`

GetIsMachine returns the IsMachine field if non-nil, zero value otherwise.

### GetIsMachineOk

`func (o *HypervisorTemplateResourceResponseModel) GetIsMachineOk() (*bool, bool)`

GetIsMachineOk returns a tuple with the IsMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachine

`func (o *HypervisorTemplateResourceResponseModel) SetIsMachine(v bool)`

SetIsMachine sets IsMachine field to given value.


### GetIsSnapshotable

`func (o *HypervisorTemplateResourceResponseModel) GetIsSnapshotable() bool`

GetIsSnapshotable returns the IsSnapshotable field if non-nil, zero value otherwise.

### GetIsSnapshotableOk

`func (o *HypervisorTemplateResourceResponseModel) GetIsSnapshotableOk() (*bool, bool)`

GetIsSnapshotableOk returns a tuple with the IsSnapshotable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotable

`func (o *HypervisorTemplateResourceResponseModel) SetIsSnapshotable(v bool)`

SetIsSnapshotable sets IsSnapshotable field to given value.


### GetAllResourcesRelativePath

`func (o *HypervisorTemplateResourceResponseModel) GetAllResourcesRelativePath() string`

GetAllResourcesRelativePath returns the AllResourcesRelativePath field if non-nil, zero value otherwise.

### GetAllResourcesRelativePathOk

`func (o *HypervisorTemplateResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool)`

GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllResourcesRelativePath

`func (o *HypervisorTemplateResourceResponseModel) SetAllResourcesRelativePath(v string)`

SetAllResourcesRelativePath sets AllResourcesRelativePath field to given value.


### GetResourcePool

`func (o *HypervisorTemplateResourceResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *HypervisorTemplateResourceResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *HypervisorTemplateResourceResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel)`

SetResourcePool sets ResourcePool field to given value.


### GetIsSymLink

`func (o *HypervisorTemplateResourceResponseModel) GetIsSymLink() bool`

GetIsSymLink returns the IsSymLink field if non-nil, zero value otherwise.

### GetIsSymLinkOk

`func (o *HypervisorTemplateResourceResponseModel) GetIsSymLinkOk() (*bool, bool)`

GetIsSymLinkOk returns a tuple with the IsSymLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSymLink

`func (o *HypervisorTemplateResourceResponseModel) SetIsSymLink(v bool)`

SetIsSymLink sets IsSymLink field to given value.


### GetAdditionalData

`func (o *HypervisorTemplateResourceResponseModel) GetAdditionalData() []NameValueStringPairModel`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *HypervisorTemplateResourceResponseModel) GetAdditionalDataOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *HypervisorTemplateResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *HypervisorTemplateResourceResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *HypervisorTemplateResourceResponseModel) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *HypervisorTemplateResourceResponseModel) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil
### GetCpuCount

`func (o *HypervisorTemplateResourceResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *HypervisorTemplateResourceResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *HypervisorTemplateResourceResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *HypervisorTemplateResourceResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *HypervisorTemplateResourceResponseModel) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *HypervisorTemplateResourceResponseModel) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetMemoryMB

`func (o *HypervisorTemplateResourceResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *HypervisorTemplateResourceResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *HypervisorTemplateResourceResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *HypervisorTemplateResourceResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### SetMemoryMBNil

`func (o *HypervisorTemplateResourceResponseModel) SetMemoryMBNil(b bool)`

 SetMemoryMBNil sets the value for MemoryMB to be an explicit nil

### UnsetMemoryMB
`func (o *HypervisorTemplateResourceResponseModel) UnsetMemoryMB()`

UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
### GetHardDiskSizeGB

`func (o *HypervisorTemplateResourceResponseModel) GetHardDiskSizeGB() int32`

GetHardDiskSizeGB returns the HardDiskSizeGB field if non-nil, zero value otherwise.

### GetHardDiskSizeGBOk

`func (o *HypervisorTemplateResourceResponseModel) GetHardDiskSizeGBOk() (*int32, bool)`

GetHardDiskSizeGBOk returns a tuple with the HardDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDiskSizeGB

`func (o *HypervisorTemplateResourceResponseModel) SetHardDiskSizeGB(v int32)`

SetHardDiskSizeGB sets HardDiskSizeGB field to given value.

### HasHardDiskSizeGB

`func (o *HypervisorTemplateResourceResponseModel) HasHardDiskSizeGB() bool`

HasHardDiskSizeGB returns a boolean if a field has been set.

### SetHardDiskSizeGBNil

`func (o *HypervisorTemplateResourceResponseModel) SetHardDiskSizeGBNil(b bool)`

 SetHardDiskSizeGBNil sets the value for HardDiskSizeGB to be an explicit nil

### UnsetHardDiskSizeGB
`func (o *HypervisorTemplateResourceResponseModel) UnsetHardDiskSizeGB()`

UnsetHardDiskSizeGB ensures that no value is present for HardDiskSizeGB, not even an explicit nil
### GetMinMemoryMB

`func (o *HypervisorTemplateResourceResponseModel) GetMinMemoryMB() int32`

GetMinMemoryMB returns the MinMemoryMB field if non-nil, zero value otherwise.

### GetMinMemoryMBOk

`func (o *HypervisorTemplateResourceResponseModel) GetMinMemoryMBOk() (*int32, bool)`

GetMinMemoryMBOk returns a tuple with the MinMemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemoryMB

`func (o *HypervisorTemplateResourceResponseModel) SetMinMemoryMB(v int32)`

SetMinMemoryMB sets MinMemoryMB field to given value.

### HasMinMemoryMB

`func (o *HypervisorTemplateResourceResponseModel) HasMinMemoryMB() bool`

HasMinMemoryMB returns a boolean if a field has been set.

### SetMinMemoryMBNil

`func (o *HypervisorTemplateResourceResponseModel) SetMinMemoryMBNil(b bool)`

 SetMinMemoryMBNil sets the value for MinMemoryMB to be an explicit nil

### UnsetMinMemoryMB
`func (o *HypervisorTemplateResourceResponseModel) UnsetMinMemoryMB()`

UnsetMinMemoryMB ensures that no value is present for MinMemoryMB, not even an explicit nil
### GetNetworkMappings

`func (o *HypervisorTemplateResourceResponseModel) GetNetworkMappings() []NetworkMapResponseModel`

GetNetworkMappings returns the NetworkMappings field if non-nil, zero value otherwise.

### GetNetworkMappingsOk

`func (o *HypervisorTemplateResourceResponseModel) GetNetworkMappingsOk() (*[]NetworkMapResponseModel, bool)`

GetNetworkMappingsOk returns a tuple with the NetworkMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappings

`func (o *HypervisorTemplateResourceResponseModel) SetNetworkMappings(v []NetworkMapResponseModel)`

SetNetworkMappings sets NetworkMappings field to given value.

### HasNetworkMappings

`func (o *HypervisorTemplateResourceResponseModel) HasNetworkMappings() bool`

HasNetworkMappings returns a boolean if a field has been set.

### SetNetworkMappingsNil

`func (o *HypervisorTemplateResourceResponseModel) SetNetworkMappingsNil(b bool)`

 SetNetworkMappingsNil sets the value for NetworkMappings to be an explicit nil

### UnsetNetworkMappings
`func (o *HypervisorTemplateResourceResponseModel) UnsetNetworkMappings()`

UnsetNetworkMappings ensures that no value is present for NetworkMappings, not even an explicit nil
### GetAttachedDisks

`func (o *HypervisorTemplateResourceResponseModel) GetAttachedDisks() []AttachedDiskResponseModel`

GetAttachedDisks returns the AttachedDisks field if non-nil, zero value otherwise.

### GetAttachedDisksOk

`func (o *HypervisorTemplateResourceResponseModel) GetAttachedDisksOk() (*[]AttachedDiskResponseModel, bool)`

GetAttachedDisksOk returns a tuple with the AttachedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedDisks

`func (o *HypervisorTemplateResourceResponseModel) SetAttachedDisks(v []AttachedDiskResponseModel)`

SetAttachedDisks sets AttachedDisks field to given value.

### HasAttachedDisks

`func (o *HypervisorTemplateResourceResponseModel) HasAttachedDisks() bool`

HasAttachedDisks returns a boolean if a field has been set.

### SetAttachedDisksNil

`func (o *HypervisorTemplateResourceResponseModel) SetAttachedDisksNil(b bool)`

 SetAttachedDisksNil sets the value for AttachedDisks to be an explicit nil

### UnsetAttachedDisks
`func (o *HypervisorTemplateResourceResponseModel) UnsetAttachedDisks()`

UnsetAttachedDisks ensures that no value is present for AttachedDisks, not even an explicit nil
### GetOwner

`func (o *HypervisorTemplateResourceResponseModel) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HypervisorTemplateResourceResponseModel) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HypervisorTemplateResourceResponseModel) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *HypervisorTemplateResourceResponseModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *HypervisorTemplateResourceResponseModel) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *HypervisorTemplateResourceResponseModel) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetIsWindowsTemplate

`func (o *HypervisorTemplateResourceResponseModel) GetIsWindowsTemplate() bool`

GetIsWindowsTemplate returns the IsWindowsTemplate field if non-nil, zero value otherwise.

### GetIsWindowsTemplateOk

`func (o *HypervisorTemplateResourceResponseModel) GetIsWindowsTemplateOk() (*bool, bool)`

GetIsWindowsTemplateOk returns a tuple with the IsWindowsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindowsTemplate

`func (o *HypervisorTemplateResourceResponseModel) SetIsWindowsTemplate(v bool)`

SetIsWindowsTemplate sets IsWindowsTemplate field to given value.

### HasIsWindowsTemplate

`func (o *HypervisorTemplateResourceResponseModel) HasIsWindowsTemplate() bool`

HasIsWindowsTemplate returns a boolean if a field has been set.

### SetIsWindowsTemplateNil

`func (o *HypervisorTemplateResourceResponseModel) SetIsWindowsTemplateNil(b bool)`

 SetIsWindowsTemplateNil sets the value for IsWindowsTemplate to be an explicit nil

### UnsetIsWindowsTemplate
`func (o *HypervisorTemplateResourceResponseModel) UnsetIsWindowsTemplate()`

UnsetIsWindowsTemplate ensures that no value is present for IsWindowsTemplate, not even an explicit nil
### GetDescription

`func (o *HypervisorTemplateResourceResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HypervisorTemplateResourceResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HypervisorTemplateResourceResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HypervisorTemplateResourceResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HypervisorTemplateResourceResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HypervisorTemplateResourceResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHasPersistentRootVolume

`func (o *HypervisorTemplateResourceResponseModel) GetHasPersistentRootVolume() bool`

GetHasPersistentRootVolume returns the HasPersistentRootVolume field if non-nil, zero value otherwise.

### GetHasPersistentRootVolumeOk

`func (o *HypervisorTemplateResourceResponseModel) GetHasPersistentRootVolumeOk() (*bool, bool)`

GetHasPersistentRootVolumeOk returns a tuple with the HasPersistentRootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPersistentRootVolume

`func (o *HypervisorTemplateResourceResponseModel) SetHasPersistentRootVolume(v bool)`

SetHasPersistentRootVolume sets HasPersistentRootVolume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


