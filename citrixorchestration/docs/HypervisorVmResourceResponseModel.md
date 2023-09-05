# HypervisorVmResourceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**XDPath** | Pointer to **string** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**RelativePath** | Pointer to **string** | Path to the object, relative to the resource pool from which it was queried. &#x60;{{vm name}}.vm/{{snapshot name}}.snapshot&#x60; | [optional] 
**FullRelativePath** | Pointer to **string** | Full path to the resource, including the hypervisor, relative to the root of the API. Example: &#x60;Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources/{{RelativePath}}&#x60; | [optional] 
**ResourceType** | **string** | Type of resource. | 
**ObjectTypeName** | Pointer to **string** | The type name of the hypervisor resource object. | [optional] 
**ResourceContainer** | Pointer to [**HypervisorResourceRefResponseModelAllOfResourceContainer**](HypervisorResourceRefResponseModelAllOfResourceContainer.md) |  | [optional] 
**ResourcePoolXDPath** | Pointer to **string** | Citrix Apps and Desktops path to the resource on the ResourcePool.  An example value is: &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; This value  | [optional] 
**FullName** | **string** | Name of the resource, with the type concatenated. i.e. \&quot;name.type\&quot;. | 
**IsContainer** | **bool** | Indicates whether the resource is a container. | 
**Children** | Pointer to [**[]HypervisorResourceResponseModel**](HypervisorResourceResponseModel.md) | Child resources, when the resource is a container. | [optional] 
**IsMachine** | **bool** | Indicates whether the resource is a machine. | 
**IsSnapshotable** | **bool** | Indicates whether the resource can have a snapshot taken. | 
**AllResourcesRelativePath** | **string** | Path to the resource, relative to the special \&quot;AllResources\&quot; resource pool associated with the hypervisor. | 
**ResourcePool** | [**HypervisorResourceResponseModelAllOfResourcePool**](HypervisorResourceResponseModelAllOfResourcePool.md) |  | 
**IsSymLink** | **bool** | Indicates whether the object is a valid symbolic link. | 
**AdditionalData** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Additional data about the object in the form of key-value pairs. | [optional] 
**CpuCount** | Pointer to **int32** | Number of CPUs, if known. | [optional] 
**MemoryMB** | Pointer to **int32** | Memory in megabytes, if known. | [optional] 
**HardDiskSizeGB** | Pointer to **int32** | Hard disk size in gigabytes, if known. | [optional] 
**MinMemoryMB** | Pointer to **int32** | Minimum memory required to run this VM or snapshot, in megabytes, if known. | [optional] 
**NetworkMappings** | Pointer to [**[]NetworkMapResponseModel**](NetworkMapResponseModel.md) | Network mappings associated with the VM, if known. | [optional] 
**AttachedDisks** | Pointer to [**[]AttachedDiskResponseModel**](AttachedDiskResponseModel.md) | Disks attached to the VM, if known. | [optional] 
**VMId** | **string** | Id of the VM, as defined by the hypervisor. | 
**MacAddress** | Pointer to **string** | MAC address of the VM. | [optional] 

## Methods

### NewHypervisorVmResourceResponseModel

`func NewHypervisorVmResourceResponseModel(resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourceResponseModelAllOfResourcePool, isSymLink bool, vMId string, ) *HypervisorVmResourceResponseModel`

NewHypervisorVmResourceResponseModel instantiates a new HypervisorVmResourceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorVmResourceResponseModelWithDefaults

`func NewHypervisorVmResourceResponseModelWithDefaults() *HypervisorVmResourceResponseModel`

NewHypervisorVmResourceResponseModelWithDefaults instantiates a new HypervisorVmResourceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorVmResourceResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorVmResourceResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorVmResourceResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorVmResourceResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorVmResourceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorVmResourceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorVmResourceResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorVmResourceResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorVmResourceResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorVmResourceResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorVmResourceResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorVmResourceResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetRelativePath

`func (o *HypervisorVmResourceResponseModel) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HypervisorVmResourceResponseModel) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HypervisorVmResourceResponseModel) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *HypervisorVmResourceResponseModel) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetFullRelativePath

`func (o *HypervisorVmResourceResponseModel) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorVmResourceResponseModel) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorVmResourceResponseModel) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.

### HasFullRelativePath

`func (o *HypervisorVmResourceResponseModel) HasFullRelativePath() bool`

HasFullRelativePath returns a boolean if a field has been set.

### GetResourceType

`func (o *HypervisorVmResourceResponseModel) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorVmResourceResponseModel) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorVmResourceResponseModel) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetObjectTypeName

`func (o *HypervisorVmResourceResponseModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *HypervisorVmResourceResponseModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *HypervisorVmResourceResponseModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *HypervisorVmResourceResponseModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetResourceContainer

`func (o *HypervisorVmResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModelAllOfResourceContainer`

GetResourceContainer returns the ResourceContainer field if non-nil, zero value otherwise.

### GetResourceContainerOk

`func (o *HypervisorVmResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModelAllOfResourceContainer, bool)`

GetResourceContainerOk returns a tuple with the ResourceContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceContainer

`func (o *HypervisorVmResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModelAllOfResourceContainer)`

SetResourceContainer sets ResourceContainer field to given value.

### HasResourceContainer

`func (o *HypervisorVmResourceResponseModel) HasResourceContainer() bool`

HasResourceContainer returns a boolean if a field has been set.

### GetResourcePoolXDPath

`func (o *HypervisorVmResourceResponseModel) GetResourcePoolXDPath() string`

GetResourcePoolXDPath returns the ResourcePoolXDPath field if non-nil, zero value otherwise.

### GetResourcePoolXDPathOk

`func (o *HypervisorVmResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool)`

GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolXDPath

`func (o *HypervisorVmResourceResponseModel) SetResourcePoolXDPath(v string)`

SetResourcePoolXDPath sets ResourcePoolXDPath field to given value.

### HasResourcePoolXDPath

`func (o *HypervisorVmResourceResponseModel) HasResourcePoolXDPath() bool`

HasResourcePoolXDPath returns a boolean if a field has been set.

### GetFullName

`func (o *HypervisorVmResourceResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *HypervisorVmResourceResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *HypervisorVmResourceResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsContainer

`func (o *HypervisorVmResourceResponseModel) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *HypervisorVmResourceResponseModel) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *HypervisorVmResourceResponseModel) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.


### GetChildren

`func (o *HypervisorVmResourceResponseModel) GetChildren() []HypervisorResourceResponseModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HypervisorVmResourceResponseModel) GetChildrenOk() (*[]HypervisorResourceResponseModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HypervisorVmResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HypervisorVmResourceResponseModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetIsMachine

`func (o *HypervisorVmResourceResponseModel) GetIsMachine() bool`

GetIsMachine returns the IsMachine field if non-nil, zero value otherwise.

### GetIsMachineOk

`func (o *HypervisorVmResourceResponseModel) GetIsMachineOk() (*bool, bool)`

GetIsMachineOk returns a tuple with the IsMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachine

`func (o *HypervisorVmResourceResponseModel) SetIsMachine(v bool)`

SetIsMachine sets IsMachine field to given value.


### GetIsSnapshotable

`func (o *HypervisorVmResourceResponseModel) GetIsSnapshotable() bool`

GetIsSnapshotable returns the IsSnapshotable field if non-nil, zero value otherwise.

### GetIsSnapshotableOk

`func (o *HypervisorVmResourceResponseModel) GetIsSnapshotableOk() (*bool, bool)`

GetIsSnapshotableOk returns a tuple with the IsSnapshotable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotable

`func (o *HypervisorVmResourceResponseModel) SetIsSnapshotable(v bool)`

SetIsSnapshotable sets IsSnapshotable field to given value.


### GetAllResourcesRelativePath

`func (o *HypervisorVmResourceResponseModel) GetAllResourcesRelativePath() string`

GetAllResourcesRelativePath returns the AllResourcesRelativePath field if non-nil, zero value otherwise.

### GetAllResourcesRelativePathOk

`func (o *HypervisorVmResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool)`

GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllResourcesRelativePath

`func (o *HypervisorVmResourceResponseModel) SetAllResourcesRelativePath(v string)`

SetAllResourcesRelativePath sets AllResourcesRelativePath field to given value.


### GetResourcePool

`func (o *HypervisorVmResourceResponseModel) GetResourcePool() HypervisorResourceResponseModelAllOfResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *HypervisorVmResourceResponseModel) GetResourcePoolOk() (*HypervisorResourceResponseModelAllOfResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *HypervisorVmResourceResponseModel) SetResourcePool(v HypervisorResourceResponseModelAllOfResourcePool)`

SetResourcePool sets ResourcePool field to given value.


### GetIsSymLink

`func (o *HypervisorVmResourceResponseModel) GetIsSymLink() bool`

GetIsSymLink returns the IsSymLink field if non-nil, zero value otherwise.

### GetIsSymLinkOk

`func (o *HypervisorVmResourceResponseModel) GetIsSymLinkOk() (*bool, bool)`

GetIsSymLinkOk returns a tuple with the IsSymLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSymLink

`func (o *HypervisorVmResourceResponseModel) SetIsSymLink(v bool)`

SetIsSymLink sets IsSymLink field to given value.


### GetAdditionalData

`func (o *HypervisorVmResourceResponseModel) GetAdditionalData() []NameValueStringPairModel`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *HypervisorVmResourceResponseModel) GetAdditionalDataOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *HypervisorVmResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *HypervisorVmResourceResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetCpuCount

`func (o *HypervisorVmResourceResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *HypervisorVmResourceResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *HypervisorVmResourceResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *HypervisorVmResourceResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetMemoryMB

`func (o *HypervisorVmResourceResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *HypervisorVmResourceResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *HypervisorVmResourceResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *HypervisorVmResourceResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetHardDiskSizeGB

`func (o *HypervisorVmResourceResponseModel) GetHardDiskSizeGB() int32`

GetHardDiskSizeGB returns the HardDiskSizeGB field if non-nil, zero value otherwise.

### GetHardDiskSizeGBOk

`func (o *HypervisorVmResourceResponseModel) GetHardDiskSizeGBOk() (*int32, bool)`

GetHardDiskSizeGBOk returns a tuple with the HardDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDiskSizeGB

`func (o *HypervisorVmResourceResponseModel) SetHardDiskSizeGB(v int32)`

SetHardDiskSizeGB sets HardDiskSizeGB field to given value.

### HasHardDiskSizeGB

`func (o *HypervisorVmResourceResponseModel) HasHardDiskSizeGB() bool`

HasHardDiskSizeGB returns a boolean if a field has been set.

### GetMinMemoryMB

`func (o *HypervisorVmResourceResponseModel) GetMinMemoryMB() int32`

GetMinMemoryMB returns the MinMemoryMB field if non-nil, zero value otherwise.

### GetMinMemoryMBOk

`func (o *HypervisorVmResourceResponseModel) GetMinMemoryMBOk() (*int32, bool)`

GetMinMemoryMBOk returns a tuple with the MinMemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemoryMB

`func (o *HypervisorVmResourceResponseModel) SetMinMemoryMB(v int32)`

SetMinMemoryMB sets MinMemoryMB field to given value.

### HasMinMemoryMB

`func (o *HypervisorVmResourceResponseModel) HasMinMemoryMB() bool`

HasMinMemoryMB returns a boolean if a field has been set.

### GetNetworkMappings

`func (o *HypervisorVmResourceResponseModel) GetNetworkMappings() []NetworkMapResponseModel`

GetNetworkMappings returns the NetworkMappings field if non-nil, zero value otherwise.

### GetNetworkMappingsOk

`func (o *HypervisorVmResourceResponseModel) GetNetworkMappingsOk() (*[]NetworkMapResponseModel, bool)`

GetNetworkMappingsOk returns a tuple with the NetworkMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappings

`func (o *HypervisorVmResourceResponseModel) SetNetworkMappings(v []NetworkMapResponseModel)`

SetNetworkMappings sets NetworkMappings field to given value.

### HasNetworkMappings

`func (o *HypervisorVmResourceResponseModel) HasNetworkMappings() bool`

HasNetworkMappings returns a boolean if a field has been set.

### GetAttachedDisks

`func (o *HypervisorVmResourceResponseModel) GetAttachedDisks() []AttachedDiskResponseModel`

GetAttachedDisks returns the AttachedDisks field if non-nil, zero value otherwise.

### GetAttachedDisksOk

`func (o *HypervisorVmResourceResponseModel) GetAttachedDisksOk() (*[]AttachedDiskResponseModel, bool)`

GetAttachedDisksOk returns a tuple with the AttachedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedDisks

`func (o *HypervisorVmResourceResponseModel) SetAttachedDisks(v []AttachedDiskResponseModel)`

SetAttachedDisks sets AttachedDisks field to given value.

### HasAttachedDisks

`func (o *HypervisorVmResourceResponseModel) HasAttachedDisks() bool`

HasAttachedDisks returns a boolean if a field has been set.

### GetVMId

`func (o *HypervisorVmResourceResponseModel) GetVMId() string`

GetVMId returns the VMId field if non-nil, zero value otherwise.

### GetVMIdOk

`func (o *HypervisorVmResourceResponseModel) GetVMIdOk() (*string, bool)`

GetVMIdOk returns a tuple with the VMId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMId

`func (o *HypervisorVmResourceResponseModel) SetVMId(v string)`

SetVMId sets VMId field to given value.


### GetMacAddress

`func (o *HypervisorVmResourceResponseModel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *HypervisorVmResourceResponseModel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *HypervisorVmResourceResponseModel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *HypervisorVmResourceResponseModel) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


