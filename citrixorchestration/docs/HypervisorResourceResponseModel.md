# HypervisorResourceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsSecurityGroups** | **bool** | Indicates whether this zone supports the use of security groups for isolation. | 
**Enabled** | **bool** | Indicates whether the GPU type is enabled on the hypervisor. | 
**FrameBufferSizeMB** | **int32** | Frame Buffer Size in MB | 
**HasDedicatedResource** | **bool** | Indicates whether the GPU type has resources dedicated to it. | 
**Description** | Pointer to **NullableString** | Description of the security group. | [optional] 
**EgressRules** | [**[]HypervisorSecurityGroupRuleResponseModel**](HypervisorSecurityGroupRuleResponseModel.md) | Egress rules governing outbound network traffic. | 
**IngressRules** | [**[]HypervisorSecurityGroupRuleResponseModel**](HypervisorSecurityGroupRuleResponseModel.md) | Ingress rules governing inbound network traffic. | 
**VirtualPrivateCloudId** | Pointer to **NullableString** | The identifier of the associated VPC. | [optional] 
**DedicatedTenancy** | Pointer to **NullableString** | Indicates whether this service offering prescribes dedicated tenancy for the instances that use it. | [optional] 
**MemorySizeMB** | **float32** | The amount of memory that is available to instances using this service offering, measured in megabytes. | 
**NumberOfCores** | **float32** | The number of virtual cores that are available to instances using this service offering. | 
**AmazonComputeUnits** | Pointer to **NullableFloat32** | The number of Amazon EC2 compute units that are available to instances using this service offering. This property is only meaningful for EC2-based cloud connections. | [optional] 
**OptimizedForPooledDesktops** | **bool** | Indicates whether this service offering is intended specifically for usage with pooled desktops. | 
**NetworkPerformance** | Pointer to **NullableString** | The network performance available to instances using this service offering. | [optional] 
**NetworkPerformanceIsUnlimited** | **bool** | Indicates whether this service offering has no limit for it&#39;s network rate. | 
**NetworkPerformanceIsDefault** | **bool** | Indicates whether this service offering has a default limit for it&#39;s network rate. | 
**IsUsedInSite** | **bool** | Indicates whether the service offering is used by any machine catalogs in the site. | 
**SupportsAzurePremiumStorage** | Pointer to **NullableBool** | Indicates whether the service offering supports premium storage. This property is only meaningful on Azure. | [optional] 
**Superseded** | **bool** | Indicates whether storage has been superseded.  Superseded storage may be used for existing virtual machines, but is not used when provisioning new virtual machines. | 
**Owner** | Pointer to **NullableString** | The account ID for the owner of this template. | [optional] 
**IsWindowsTemplate** | Pointer to **NullableBool** | Indicates whether this is a Windows OS template, if known. | [optional] 
**HasPersistentRootVolume** | **bool** | Indicates whether this template has a persistent root volume (eg. is an EBS-backed image on AWS). | 
**VMId** | **string** | Id of the VM, as defined by the hypervisor. | 
**MacAddress** | Pointer to **NullableString** | MAC address of the VM. | [optional] 
**CpuCount** | Pointer to **NullableInt32** | Number of CPUs, if known. | [optional] 
**MemoryMB** | Pointer to **NullableInt32** | Memory in megabytes, if known. | [optional] 
**HardDiskSizeGB** | Pointer to **NullableInt32** | Hard disk size in gigabytes, if known. | [optional] 
**MinMemoryMB** | Pointer to **NullableInt32** | Minimum memory required to run this VM or snapshot, in megabytes, if known. | [optional] 
**NetworkMappings** | Pointer to [**[]NetworkMapResponseModel**](NetworkMapResponseModel.md) | Network mappings associated with the VM, if known. | [optional] 
**AttachedDisks** | Pointer to [**[]AttachedDiskResponseModel**](AttachedDiskResponseModel.md) | Disks attached to the VM, if known. | [optional] 
**InstanceSecurityGroupLimit** | **int32** | Indicates the maximum number of security groups allowed per instance in this VPC | 
**Endpoint** | Pointer to **NullableString** | The API address with the region. | [optional] 
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

## Methods

### NewHypervisorResourceResponseModel

`func NewHypervisorResourceResponseModel(supportsSecurityGroups bool, enabled bool, frameBufferSizeMB int32, hasDedicatedResource bool, egressRules []HypervisorSecurityGroupRuleResponseModel, ingressRules []HypervisorSecurityGroupRuleResponseModel, memorySizeMB float32, numberOfCores float32, optimizedForPooledDesktops bool, networkPerformanceIsUnlimited bool, networkPerformanceIsDefault bool, isUsedInSite bool, superseded bool, hasPersistentRootVolume bool, vMId string, instanceSecurityGroupLimit int32, resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourcePoolRefResponseModel, isSymLink bool, ) *HypervisorResourceResponseModel`

NewHypervisorResourceResponseModel instantiates a new HypervisorResourceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourceResponseModelWithDefaults

`func NewHypervisorResourceResponseModelWithDefaults() *HypervisorResourceResponseModel`

NewHypervisorResourceResponseModelWithDefaults instantiates a new HypervisorResourceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsSecurityGroups

`func (o *HypervisorResourceResponseModel) GetSupportsSecurityGroups() bool`

GetSupportsSecurityGroups returns the SupportsSecurityGroups field if non-nil, zero value otherwise.

### GetSupportsSecurityGroupsOk

`func (o *HypervisorResourceResponseModel) GetSupportsSecurityGroupsOk() (*bool, bool)`

GetSupportsSecurityGroupsOk returns a tuple with the SupportsSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSecurityGroups

`func (o *HypervisorResourceResponseModel) SetSupportsSecurityGroups(v bool)`

SetSupportsSecurityGroups sets SupportsSecurityGroups field to given value.


### GetEnabled

`func (o *HypervisorResourceResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HypervisorResourceResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HypervisorResourceResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFrameBufferSizeMB

`func (o *HypervisorResourceResponseModel) GetFrameBufferSizeMB() int32`

GetFrameBufferSizeMB returns the FrameBufferSizeMB field if non-nil, zero value otherwise.

### GetFrameBufferSizeMBOk

`func (o *HypervisorResourceResponseModel) GetFrameBufferSizeMBOk() (*int32, bool)`

GetFrameBufferSizeMBOk returns a tuple with the FrameBufferSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameBufferSizeMB

`func (o *HypervisorResourceResponseModel) SetFrameBufferSizeMB(v int32)`

SetFrameBufferSizeMB sets FrameBufferSizeMB field to given value.


### GetHasDedicatedResource

`func (o *HypervisorResourceResponseModel) GetHasDedicatedResource() bool`

GetHasDedicatedResource returns the HasDedicatedResource field if non-nil, zero value otherwise.

### GetHasDedicatedResourceOk

`func (o *HypervisorResourceResponseModel) GetHasDedicatedResourceOk() (*bool, bool)`

GetHasDedicatedResourceOk returns a tuple with the HasDedicatedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDedicatedResource

`func (o *HypervisorResourceResponseModel) SetHasDedicatedResource(v bool)`

SetHasDedicatedResource sets HasDedicatedResource field to given value.


### GetDescription

`func (o *HypervisorResourceResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HypervisorResourceResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HypervisorResourceResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HypervisorResourceResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HypervisorResourceResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HypervisorResourceResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEgressRules

`func (o *HypervisorResourceResponseModel) GetEgressRules() []HypervisorSecurityGroupRuleResponseModel`

GetEgressRules returns the EgressRules field if non-nil, zero value otherwise.

### GetEgressRulesOk

`func (o *HypervisorResourceResponseModel) GetEgressRulesOk() (*[]HypervisorSecurityGroupRuleResponseModel, bool)`

GetEgressRulesOk returns a tuple with the EgressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRules

`func (o *HypervisorResourceResponseModel) SetEgressRules(v []HypervisorSecurityGroupRuleResponseModel)`

SetEgressRules sets EgressRules field to given value.


### GetIngressRules

`func (o *HypervisorResourceResponseModel) GetIngressRules() []HypervisorSecurityGroupRuleResponseModel`

GetIngressRules returns the IngressRules field if non-nil, zero value otherwise.

### GetIngressRulesOk

`func (o *HypervisorResourceResponseModel) GetIngressRulesOk() (*[]HypervisorSecurityGroupRuleResponseModel, bool)`

GetIngressRulesOk returns a tuple with the IngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRules

`func (o *HypervisorResourceResponseModel) SetIngressRules(v []HypervisorSecurityGroupRuleResponseModel)`

SetIngressRules sets IngressRules field to given value.


### GetVirtualPrivateCloudId

`func (o *HypervisorResourceResponseModel) GetVirtualPrivateCloudId() string`

GetVirtualPrivateCloudId returns the VirtualPrivateCloudId field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudIdOk

`func (o *HypervisorResourceResponseModel) GetVirtualPrivateCloudIdOk() (*string, bool)`

GetVirtualPrivateCloudIdOk returns a tuple with the VirtualPrivateCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloudId

`func (o *HypervisorResourceResponseModel) SetVirtualPrivateCloudId(v string)`

SetVirtualPrivateCloudId sets VirtualPrivateCloudId field to given value.

### HasVirtualPrivateCloudId

`func (o *HypervisorResourceResponseModel) HasVirtualPrivateCloudId() bool`

HasVirtualPrivateCloudId returns a boolean if a field has been set.

### SetVirtualPrivateCloudIdNil

`func (o *HypervisorResourceResponseModel) SetVirtualPrivateCloudIdNil(b bool)`

 SetVirtualPrivateCloudIdNil sets the value for VirtualPrivateCloudId to be an explicit nil

### UnsetVirtualPrivateCloudId
`func (o *HypervisorResourceResponseModel) UnsetVirtualPrivateCloudId()`

UnsetVirtualPrivateCloudId ensures that no value is present for VirtualPrivateCloudId, not even an explicit nil
### GetDedicatedTenancy

`func (o *HypervisorResourceResponseModel) GetDedicatedTenancy() string`

GetDedicatedTenancy returns the DedicatedTenancy field if non-nil, zero value otherwise.

### GetDedicatedTenancyOk

`func (o *HypervisorResourceResponseModel) GetDedicatedTenancyOk() (*string, bool)`

GetDedicatedTenancyOk returns a tuple with the DedicatedTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancy

`func (o *HypervisorResourceResponseModel) SetDedicatedTenancy(v string)`

SetDedicatedTenancy sets DedicatedTenancy field to given value.

### HasDedicatedTenancy

`func (o *HypervisorResourceResponseModel) HasDedicatedTenancy() bool`

HasDedicatedTenancy returns a boolean if a field has been set.

### SetDedicatedTenancyNil

`func (o *HypervisorResourceResponseModel) SetDedicatedTenancyNil(b bool)`

 SetDedicatedTenancyNil sets the value for DedicatedTenancy to be an explicit nil

### UnsetDedicatedTenancy
`func (o *HypervisorResourceResponseModel) UnsetDedicatedTenancy()`

UnsetDedicatedTenancy ensures that no value is present for DedicatedTenancy, not even an explicit nil
### GetMemorySizeMB

`func (o *HypervisorResourceResponseModel) GetMemorySizeMB() float32`

GetMemorySizeMB returns the MemorySizeMB field if non-nil, zero value otherwise.

### GetMemorySizeMBOk

`func (o *HypervisorResourceResponseModel) GetMemorySizeMBOk() (*float32, bool)`

GetMemorySizeMBOk returns a tuple with the MemorySizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeMB

`func (o *HypervisorResourceResponseModel) SetMemorySizeMB(v float32)`

SetMemorySizeMB sets MemorySizeMB field to given value.


### GetNumberOfCores

`func (o *HypervisorResourceResponseModel) GetNumberOfCores() float32`

GetNumberOfCores returns the NumberOfCores field if non-nil, zero value otherwise.

### GetNumberOfCoresOk

`func (o *HypervisorResourceResponseModel) GetNumberOfCoresOk() (*float32, bool)`

GetNumberOfCoresOk returns a tuple with the NumberOfCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCores

`func (o *HypervisorResourceResponseModel) SetNumberOfCores(v float32)`

SetNumberOfCores sets NumberOfCores field to given value.


### GetAmazonComputeUnits

`func (o *HypervisorResourceResponseModel) GetAmazonComputeUnits() float32`

GetAmazonComputeUnits returns the AmazonComputeUnits field if non-nil, zero value otherwise.

### GetAmazonComputeUnitsOk

`func (o *HypervisorResourceResponseModel) GetAmazonComputeUnitsOk() (*float32, bool)`

GetAmazonComputeUnitsOk returns a tuple with the AmazonComputeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonComputeUnits

`func (o *HypervisorResourceResponseModel) SetAmazonComputeUnits(v float32)`

SetAmazonComputeUnits sets AmazonComputeUnits field to given value.

### HasAmazonComputeUnits

`func (o *HypervisorResourceResponseModel) HasAmazonComputeUnits() bool`

HasAmazonComputeUnits returns a boolean if a field has been set.

### SetAmazonComputeUnitsNil

`func (o *HypervisorResourceResponseModel) SetAmazonComputeUnitsNil(b bool)`

 SetAmazonComputeUnitsNil sets the value for AmazonComputeUnits to be an explicit nil

### UnsetAmazonComputeUnits
`func (o *HypervisorResourceResponseModel) UnsetAmazonComputeUnits()`

UnsetAmazonComputeUnits ensures that no value is present for AmazonComputeUnits, not even an explicit nil
### GetOptimizedForPooledDesktops

`func (o *HypervisorResourceResponseModel) GetOptimizedForPooledDesktops() bool`

GetOptimizedForPooledDesktops returns the OptimizedForPooledDesktops field if non-nil, zero value otherwise.

### GetOptimizedForPooledDesktopsOk

`func (o *HypervisorResourceResponseModel) GetOptimizedForPooledDesktopsOk() (*bool, bool)`

GetOptimizedForPooledDesktopsOk returns a tuple with the OptimizedForPooledDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizedForPooledDesktops

`func (o *HypervisorResourceResponseModel) SetOptimizedForPooledDesktops(v bool)`

SetOptimizedForPooledDesktops sets OptimizedForPooledDesktops field to given value.


### GetNetworkPerformance

`func (o *HypervisorResourceResponseModel) GetNetworkPerformance() string`

GetNetworkPerformance returns the NetworkPerformance field if non-nil, zero value otherwise.

### GetNetworkPerformanceOk

`func (o *HypervisorResourceResponseModel) GetNetworkPerformanceOk() (*string, bool)`

GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformance

`func (o *HypervisorResourceResponseModel) SetNetworkPerformance(v string)`

SetNetworkPerformance sets NetworkPerformance field to given value.

### HasNetworkPerformance

`func (o *HypervisorResourceResponseModel) HasNetworkPerformance() bool`

HasNetworkPerformance returns a boolean if a field has been set.

### SetNetworkPerformanceNil

`func (o *HypervisorResourceResponseModel) SetNetworkPerformanceNil(b bool)`

 SetNetworkPerformanceNil sets the value for NetworkPerformance to be an explicit nil

### UnsetNetworkPerformance
`func (o *HypervisorResourceResponseModel) UnsetNetworkPerformance()`

UnsetNetworkPerformance ensures that no value is present for NetworkPerformance, not even an explicit nil
### GetNetworkPerformanceIsUnlimited

`func (o *HypervisorResourceResponseModel) GetNetworkPerformanceIsUnlimited() bool`

GetNetworkPerformanceIsUnlimited returns the NetworkPerformanceIsUnlimited field if non-nil, zero value otherwise.

### GetNetworkPerformanceIsUnlimitedOk

`func (o *HypervisorResourceResponseModel) GetNetworkPerformanceIsUnlimitedOk() (*bool, bool)`

GetNetworkPerformanceIsUnlimitedOk returns a tuple with the NetworkPerformanceIsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformanceIsUnlimited

`func (o *HypervisorResourceResponseModel) SetNetworkPerformanceIsUnlimited(v bool)`

SetNetworkPerformanceIsUnlimited sets NetworkPerformanceIsUnlimited field to given value.


### GetNetworkPerformanceIsDefault

`func (o *HypervisorResourceResponseModel) GetNetworkPerformanceIsDefault() bool`

GetNetworkPerformanceIsDefault returns the NetworkPerformanceIsDefault field if non-nil, zero value otherwise.

### GetNetworkPerformanceIsDefaultOk

`func (o *HypervisorResourceResponseModel) GetNetworkPerformanceIsDefaultOk() (*bool, bool)`

GetNetworkPerformanceIsDefaultOk returns a tuple with the NetworkPerformanceIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformanceIsDefault

`func (o *HypervisorResourceResponseModel) SetNetworkPerformanceIsDefault(v bool)`

SetNetworkPerformanceIsDefault sets NetworkPerformanceIsDefault field to given value.


### GetIsUsedInSite

`func (o *HypervisorResourceResponseModel) GetIsUsedInSite() bool`

GetIsUsedInSite returns the IsUsedInSite field if non-nil, zero value otherwise.

### GetIsUsedInSiteOk

`func (o *HypervisorResourceResponseModel) GetIsUsedInSiteOk() (*bool, bool)`

GetIsUsedInSiteOk returns a tuple with the IsUsedInSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedInSite

`func (o *HypervisorResourceResponseModel) SetIsUsedInSite(v bool)`

SetIsUsedInSite sets IsUsedInSite field to given value.


### GetSupportsAzurePremiumStorage

`func (o *HypervisorResourceResponseModel) GetSupportsAzurePremiumStorage() bool`

GetSupportsAzurePremiumStorage returns the SupportsAzurePremiumStorage field if non-nil, zero value otherwise.

### GetSupportsAzurePremiumStorageOk

`func (o *HypervisorResourceResponseModel) GetSupportsAzurePremiumStorageOk() (*bool, bool)`

GetSupportsAzurePremiumStorageOk returns a tuple with the SupportsAzurePremiumStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAzurePremiumStorage

`func (o *HypervisorResourceResponseModel) SetSupportsAzurePremiumStorage(v bool)`

SetSupportsAzurePremiumStorage sets SupportsAzurePremiumStorage field to given value.

### HasSupportsAzurePremiumStorage

`func (o *HypervisorResourceResponseModel) HasSupportsAzurePremiumStorage() bool`

HasSupportsAzurePremiumStorage returns a boolean if a field has been set.

### SetSupportsAzurePremiumStorageNil

`func (o *HypervisorResourceResponseModel) SetSupportsAzurePremiumStorageNil(b bool)`

 SetSupportsAzurePremiumStorageNil sets the value for SupportsAzurePremiumStorage to be an explicit nil

### UnsetSupportsAzurePremiumStorage
`func (o *HypervisorResourceResponseModel) UnsetSupportsAzurePremiumStorage()`

UnsetSupportsAzurePremiumStorage ensures that no value is present for SupportsAzurePremiumStorage, not even an explicit nil
### GetSuperseded

`func (o *HypervisorResourceResponseModel) GetSuperseded() bool`

GetSuperseded returns the Superseded field if non-nil, zero value otherwise.

### GetSupersededOk

`func (o *HypervisorResourceResponseModel) GetSupersededOk() (*bool, bool)`

GetSupersededOk returns a tuple with the Superseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperseded

`func (o *HypervisorResourceResponseModel) SetSuperseded(v bool)`

SetSuperseded sets Superseded field to given value.


### GetOwner

`func (o *HypervisorResourceResponseModel) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HypervisorResourceResponseModel) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HypervisorResourceResponseModel) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *HypervisorResourceResponseModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *HypervisorResourceResponseModel) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *HypervisorResourceResponseModel) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetIsWindowsTemplate

`func (o *HypervisorResourceResponseModel) GetIsWindowsTemplate() bool`

GetIsWindowsTemplate returns the IsWindowsTemplate field if non-nil, zero value otherwise.

### GetIsWindowsTemplateOk

`func (o *HypervisorResourceResponseModel) GetIsWindowsTemplateOk() (*bool, bool)`

GetIsWindowsTemplateOk returns a tuple with the IsWindowsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindowsTemplate

`func (o *HypervisorResourceResponseModel) SetIsWindowsTemplate(v bool)`

SetIsWindowsTemplate sets IsWindowsTemplate field to given value.

### HasIsWindowsTemplate

`func (o *HypervisorResourceResponseModel) HasIsWindowsTemplate() bool`

HasIsWindowsTemplate returns a boolean if a field has been set.

### SetIsWindowsTemplateNil

`func (o *HypervisorResourceResponseModel) SetIsWindowsTemplateNil(b bool)`

 SetIsWindowsTemplateNil sets the value for IsWindowsTemplate to be an explicit nil

### UnsetIsWindowsTemplate
`func (o *HypervisorResourceResponseModel) UnsetIsWindowsTemplate()`

UnsetIsWindowsTemplate ensures that no value is present for IsWindowsTemplate, not even an explicit nil
### GetHasPersistentRootVolume

`func (o *HypervisorResourceResponseModel) GetHasPersistentRootVolume() bool`

GetHasPersistentRootVolume returns the HasPersistentRootVolume field if non-nil, zero value otherwise.

### GetHasPersistentRootVolumeOk

`func (o *HypervisorResourceResponseModel) GetHasPersistentRootVolumeOk() (*bool, bool)`

GetHasPersistentRootVolumeOk returns a tuple with the HasPersistentRootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPersistentRootVolume

`func (o *HypervisorResourceResponseModel) SetHasPersistentRootVolume(v bool)`

SetHasPersistentRootVolume sets HasPersistentRootVolume field to given value.


### GetVMId

`func (o *HypervisorResourceResponseModel) GetVMId() string`

GetVMId returns the VMId field if non-nil, zero value otherwise.

### GetVMIdOk

`func (o *HypervisorResourceResponseModel) GetVMIdOk() (*string, bool)`

GetVMIdOk returns a tuple with the VMId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMId

`func (o *HypervisorResourceResponseModel) SetVMId(v string)`

SetVMId sets VMId field to given value.


### GetMacAddress

`func (o *HypervisorResourceResponseModel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *HypervisorResourceResponseModel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *HypervisorResourceResponseModel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *HypervisorResourceResponseModel) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *HypervisorResourceResponseModel) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *HypervisorResourceResponseModel) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetCpuCount

`func (o *HypervisorResourceResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *HypervisorResourceResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *HypervisorResourceResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *HypervisorResourceResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *HypervisorResourceResponseModel) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *HypervisorResourceResponseModel) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetMemoryMB

`func (o *HypervisorResourceResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *HypervisorResourceResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *HypervisorResourceResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *HypervisorResourceResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### SetMemoryMBNil

`func (o *HypervisorResourceResponseModel) SetMemoryMBNil(b bool)`

 SetMemoryMBNil sets the value for MemoryMB to be an explicit nil

### UnsetMemoryMB
`func (o *HypervisorResourceResponseModel) UnsetMemoryMB()`

UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
### GetHardDiskSizeGB

`func (o *HypervisorResourceResponseModel) GetHardDiskSizeGB() int32`

GetHardDiskSizeGB returns the HardDiskSizeGB field if non-nil, zero value otherwise.

### GetHardDiskSizeGBOk

`func (o *HypervisorResourceResponseModel) GetHardDiskSizeGBOk() (*int32, bool)`

GetHardDiskSizeGBOk returns a tuple with the HardDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDiskSizeGB

`func (o *HypervisorResourceResponseModel) SetHardDiskSizeGB(v int32)`

SetHardDiskSizeGB sets HardDiskSizeGB field to given value.

### HasHardDiskSizeGB

`func (o *HypervisorResourceResponseModel) HasHardDiskSizeGB() bool`

HasHardDiskSizeGB returns a boolean if a field has been set.

### SetHardDiskSizeGBNil

`func (o *HypervisorResourceResponseModel) SetHardDiskSizeGBNil(b bool)`

 SetHardDiskSizeGBNil sets the value for HardDiskSizeGB to be an explicit nil

### UnsetHardDiskSizeGB
`func (o *HypervisorResourceResponseModel) UnsetHardDiskSizeGB()`

UnsetHardDiskSizeGB ensures that no value is present for HardDiskSizeGB, not even an explicit nil
### GetMinMemoryMB

`func (o *HypervisorResourceResponseModel) GetMinMemoryMB() int32`

GetMinMemoryMB returns the MinMemoryMB field if non-nil, zero value otherwise.

### GetMinMemoryMBOk

`func (o *HypervisorResourceResponseModel) GetMinMemoryMBOk() (*int32, bool)`

GetMinMemoryMBOk returns a tuple with the MinMemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemoryMB

`func (o *HypervisorResourceResponseModel) SetMinMemoryMB(v int32)`

SetMinMemoryMB sets MinMemoryMB field to given value.

### HasMinMemoryMB

`func (o *HypervisorResourceResponseModel) HasMinMemoryMB() bool`

HasMinMemoryMB returns a boolean if a field has been set.

### SetMinMemoryMBNil

`func (o *HypervisorResourceResponseModel) SetMinMemoryMBNil(b bool)`

 SetMinMemoryMBNil sets the value for MinMemoryMB to be an explicit nil

### UnsetMinMemoryMB
`func (o *HypervisorResourceResponseModel) UnsetMinMemoryMB()`

UnsetMinMemoryMB ensures that no value is present for MinMemoryMB, not even an explicit nil
### GetNetworkMappings

`func (o *HypervisorResourceResponseModel) GetNetworkMappings() []NetworkMapResponseModel`

GetNetworkMappings returns the NetworkMappings field if non-nil, zero value otherwise.

### GetNetworkMappingsOk

`func (o *HypervisorResourceResponseModel) GetNetworkMappingsOk() (*[]NetworkMapResponseModel, bool)`

GetNetworkMappingsOk returns a tuple with the NetworkMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappings

`func (o *HypervisorResourceResponseModel) SetNetworkMappings(v []NetworkMapResponseModel)`

SetNetworkMappings sets NetworkMappings field to given value.

### HasNetworkMappings

`func (o *HypervisorResourceResponseModel) HasNetworkMappings() bool`

HasNetworkMappings returns a boolean if a field has been set.

### SetNetworkMappingsNil

`func (o *HypervisorResourceResponseModel) SetNetworkMappingsNil(b bool)`

 SetNetworkMappingsNil sets the value for NetworkMappings to be an explicit nil

### UnsetNetworkMappings
`func (o *HypervisorResourceResponseModel) UnsetNetworkMappings()`

UnsetNetworkMappings ensures that no value is present for NetworkMappings, not even an explicit nil
### GetAttachedDisks

`func (o *HypervisorResourceResponseModel) GetAttachedDisks() []AttachedDiskResponseModel`

GetAttachedDisks returns the AttachedDisks field if non-nil, zero value otherwise.

### GetAttachedDisksOk

`func (o *HypervisorResourceResponseModel) GetAttachedDisksOk() (*[]AttachedDiskResponseModel, bool)`

GetAttachedDisksOk returns a tuple with the AttachedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedDisks

`func (o *HypervisorResourceResponseModel) SetAttachedDisks(v []AttachedDiskResponseModel)`

SetAttachedDisks sets AttachedDisks field to given value.

### HasAttachedDisks

`func (o *HypervisorResourceResponseModel) HasAttachedDisks() bool`

HasAttachedDisks returns a boolean if a field has been set.

### SetAttachedDisksNil

`func (o *HypervisorResourceResponseModel) SetAttachedDisksNil(b bool)`

 SetAttachedDisksNil sets the value for AttachedDisks to be an explicit nil

### UnsetAttachedDisks
`func (o *HypervisorResourceResponseModel) UnsetAttachedDisks()`

UnsetAttachedDisks ensures that no value is present for AttachedDisks, not even an explicit nil
### GetInstanceSecurityGroupLimit

`func (o *HypervisorResourceResponseModel) GetInstanceSecurityGroupLimit() int32`

GetInstanceSecurityGroupLimit returns the InstanceSecurityGroupLimit field if non-nil, zero value otherwise.

### GetInstanceSecurityGroupLimitOk

`func (o *HypervisorResourceResponseModel) GetInstanceSecurityGroupLimitOk() (*int32, bool)`

GetInstanceSecurityGroupLimitOk returns a tuple with the InstanceSecurityGroupLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSecurityGroupLimit

`func (o *HypervisorResourceResponseModel) SetInstanceSecurityGroupLimit(v int32)`

SetInstanceSecurityGroupLimit sets InstanceSecurityGroupLimit field to given value.


### GetEndpoint

`func (o *HypervisorResourceResponseModel) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HypervisorResourceResponseModel) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HypervisorResourceResponseModel) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HypervisorResourceResponseModel) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *HypervisorResourceResponseModel) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *HypervisorResourceResponseModel) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetId

`func (o *HypervisorResourceResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorResourceResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorResourceResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorResourceResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorResourceResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorResourceResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorResourceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorResourceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorResourceResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorResourceResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorResourceResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorResourceResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorResourceResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorResourceResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorResourceResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorResourceResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorResourceResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorResourceResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetRelativePath

`func (o *HypervisorResourceResponseModel) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HypervisorResourceResponseModel) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HypervisorResourceResponseModel) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *HypervisorResourceResponseModel) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *HypervisorResourceResponseModel) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *HypervisorResourceResponseModel) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
### GetFullRelativePath

`func (o *HypervisorResourceResponseModel) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorResourceResponseModel) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorResourceResponseModel) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.

### HasFullRelativePath

`func (o *HypervisorResourceResponseModel) HasFullRelativePath() bool`

HasFullRelativePath returns a boolean if a field has been set.

### SetFullRelativePathNil

`func (o *HypervisorResourceResponseModel) SetFullRelativePathNil(b bool)`

 SetFullRelativePathNil sets the value for FullRelativePath to be an explicit nil

### UnsetFullRelativePath
`func (o *HypervisorResourceResponseModel) UnsetFullRelativePath()`

UnsetFullRelativePath ensures that no value is present for FullRelativePath, not even an explicit nil
### GetResourceType

`func (o *HypervisorResourceResponseModel) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorResourceResponseModel) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorResourceResponseModel) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetObjectTypeName

`func (o *HypervisorResourceResponseModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *HypervisorResourceResponseModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *HypervisorResourceResponseModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *HypervisorResourceResponseModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### SetObjectTypeNameNil

`func (o *HypervisorResourceResponseModel) SetObjectTypeNameNil(b bool)`

 SetObjectTypeNameNil sets the value for ObjectTypeName to be an explicit nil

### UnsetObjectTypeName
`func (o *HypervisorResourceResponseModel) UnsetObjectTypeName()`

UnsetObjectTypeName ensures that no value is present for ObjectTypeName, not even an explicit nil
### GetResourceContainer

`func (o *HypervisorResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModel`

GetResourceContainer returns the ResourceContainer field if non-nil, zero value otherwise.

### GetResourceContainerOk

`func (o *HypervisorResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModel, bool)`

GetResourceContainerOk returns a tuple with the ResourceContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceContainer

`func (o *HypervisorResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModel)`

SetResourceContainer sets ResourceContainer field to given value.

### HasResourceContainer

`func (o *HypervisorResourceResponseModel) HasResourceContainer() bool`

HasResourceContainer returns a boolean if a field has been set.

### GetResourcePoolXDPath

`func (o *HypervisorResourceResponseModel) GetResourcePoolXDPath() string`

GetResourcePoolXDPath returns the ResourcePoolXDPath field if non-nil, zero value otherwise.

### GetResourcePoolXDPathOk

`func (o *HypervisorResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool)`

GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolXDPath

`func (o *HypervisorResourceResponseModel) SetResourcePoolXDPath(v string)`

SetResourcePoolXDPath sets ResourcePoolXDPath field to given value.

### HasResourcePoolXDPath

`func (o *HypervisorResourceResponseModel) HasResourcePoolXDPath() bool`

HasResourcePoolXDPath returns a boolean if a field has been set.

### SetResourcePoolXDPathNil

`func (o *HypervisorResourceResponseModel) SetResourcePoolXDPathNil(b bool)`

 SetResourcePoolXDPathNil sets the value for ResourcePoolXDPath to be an explicit nil

### UnsetResourcePoolXDPath
`func (o *HypervisorResourceResponseModel) UnsetResourcePoolXDPath()`

UnsetResourcePoolXDPath ensures that no value is present for ResourcePoolXDPath, not even an explicit nil
### GetFullName

`func (o *HypervisorResourceResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *HypervisorResourceResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *HypervisorResourceResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsContainer

`func (o *HypervisorResourceResponseModel) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *HypervisorResourceResponseModel) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *HypervisorResourceResponseModel) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.


### GetChildren

`func (o *HypervisorResourceResponseModel) GetChildren() []HypervisorResourceResponseModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HypervisorResourceResponseModel) GetChildrenOk() (*[]HypervisorResourceResponseModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HypervisorResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HypervisorResourceResponseModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *HypervisorResourceResponseModel) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *HypervisorResourceResponseModel) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetIsMachine

`func (o *HypervisorResourceResponseModel) GetIsMachine() bool`

GetIsMachine returns the IsMachine field if non-nil, zero value otherwise.

### GetIsMachineOk

`func (o *HypervisorResourceResponseModel) GetIsMachineOk() (*bool, bool)`

GetIsMachineOk returns a tuple with the IsMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachine

`func (o *HypervisorResourceResponseModel) SetIsMachine(v bool)`

SetIsMachine sets IsMachine field to given value.


### GetIsSnapshotable

`func (o *HypervisorResourceResponseModel) GetIsSnapshotable() bool`

GetIsSnapshotable returns the IsSnapshotable field if non-nil, zero value otherwise.

### GetIsSnapshotableOk

`func (o *HypervisorResourceResponseModel) GetIsSnapshotableOk() (*bool, bool)`

GetIsSnapshotableOk returns a tuple with the IsSnapshotable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotable

`func (o *HypervisorResourceResponseModel) SetIsSnapshotable(v bool)`

SetIsSnapshotable sets IsSnapshotable field to given value.


### GetAllResourcesRelativePath

`func (o *HypervisorResourceResponseModel) GetAllResourcesRelativePath() string`

GetAllResourcesRelativePath returns the AllResourcesRelativePath field if non-nil, zero value otherwise.

### GetAllResourcesRelativePathOk

`func (o *HypervisorResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool)`

GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllResourcesRelativePath

`func (o *HypervisorResourceResponseModel) SetAllResourcesRelativePath(v string)`

SetAllResourcesRelativePath sets AllResourcesRelativePath field to given value.


### GetResourcePool

`func (o *HypervisorResourceResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *HypervisorResourceResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *HypervisorResourceResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel)`

SetResourcePool sets ResourcePool field to given value.


### GetIsSymLink

`func (o *HypervisorResourceResponseModel) GetIsSymLink() bool`

GetIsSymLink returns the IsSymLink field if non-nil, zero value otherwise.

### GetIsSymLinkOk

`func (o *HypervisorResourceResponseModel) GetIsSymLinkOk() (*bool, bool)`

GetIsSymLinkOk returns a tuple with the IsSymLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSymLink

`func (o *HypervisorResourceResponseModel) SetIsSymLink(v bool)`

SetIsSymLink sets IsSymLink field to given value.


### GetAdditionalData

`func (o *HypervisorResourceResponseModel) GetAdditionalData() []NameValueStringPairModel`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *HypervisorResourceResponseModel) GetAdditionalDataOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *HypervisorResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *HypervisorResourceResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *HypervisorResourceResponseModel) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *HypervisorResourceResponseModel) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


