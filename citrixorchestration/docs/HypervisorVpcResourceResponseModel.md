# HypervisorVpcResourceResponseModel

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

### NewHypervisorVpcResourceResponseModel

`func NewHypervisorVpcResourceResponseModel(supportsSecurityGroups bool, enabled bool, frameBufferSizeMB int32, hasDedicatedResource bool, egressRules []HypervisorSecurityGroupRuleResponseModel, ingressRules []HypervisorSecurityGroupRuleResponseModel, memorySizeMB float32, numberOfCores float32, optimizedForPooledDesktops bool, networkPerformanceIsUnlimited bool, networkPerformanceIsDefault bool, isUsedInSite bool, superseded bool, hasPersistentRootVolume bool, vMId string, instanceSecurityGroupLimit int32, resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourcePoolRefResponseModel, isSymLink bool, ) *HypervisorVpcResourceResponseModel`

NewHypervisorVpcResourceResponseModel instantiates a new HypervisorVpcResourceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorVpcResourceResponseModelWithDefaults

`func NewHypervisorVpcResourceResponseModelWithDefaults() *HypervisorVpcResourceResponseModel`

NewHypervisorVpcResourceResponseModelWithDefaults instantiates a new HypervisorVpcResourceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsSecurityGroups

`func (o *HypervisorVpcResourceResponseModel) GetSupportsSecurityGroups() bool`

GetSupportsSecurityGroups returns the SupportsSecurityGroups field if non-nil, zero value otherwise.

### GetSupportsSecurityGroupsOk

`func (o *HypervisorVpcResourceResponseModel) GetSupportsSecurityGroupsOk() (*bool, bool)`

GetSupportsSecurityGroupsOk returns a tuple with the SupportsSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSecurityGroups

`func (o *HypervisorVpcResourceResponseModel) SetSupportsSecurityGroups(v bool)`

SetSupportsSecurityGroups sets SupportsSecurityGroups field to given value.


### GetEnabled

`func (o *HypervisorVpcResourceResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HypervisorVpcResourceResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HypervisorVpcResourceResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFrameBufferSizeMB

`func (o *HypervisorVpcResourceResponseModel) GetFrameBufferSizeMB() int32`

GetFrameBufferSizeMB returns the FrameBufferSizeMB field if non-nil, zero value otherwise.

### GetFrameBufferSizeMBOk

`func (o *HypervisorVpcResourceResponseModel) GetFrameBufferSizeMBOk() (*int32, bool)`

GetFrameBufferSizeMBOk returns a tuple with the FrameBufferSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameBufferSizeMB

`func (o *HypervisorVpcResourceResponseModel) SetFrameBufferSizeMB(v int32)`

SetFrameBufferSizeMB sets FrameBufferSizeMB field to given value.


### GetHasDedicatedResource

`func (o *HypervisorVpcResourceResponseModel) GetHasDedicatedResource() bool`

GetHasDedicatedResource returns the HasDedicatedResource field if non-nil, zero value otherwise.

### GetHasDedicatedResourceOk

`func (o *HypervisorVpcResourceResponseModel) GetHasDedicatedResourceOk() (*bool, bool)`

GetHasDedicatedResourceOk returns a tuple with the HasDedicatedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDedicatedResource

`func (o *HypervisorVpcResourceResponseModel) SetHasDedicatedResource(v bool)`

SetHasDedicatedResource sets HasDedicatedResource field to given value.


### GetDescription

`func (o *HypervisorVpcResourceResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HypervisorVpcResourceResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HypervisorVpcResourceResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HypervisorVpcResourceResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HypervisorVpcResourceResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HypervisorVpcResourceResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEgressRules

`func (o *HypervisorVpcResourceResponseModel) GetEgressRules() []HypervisorSecurityGroupRuleResponseModel`

GetEgressRules returns the EgressRules field if non-nil, zero value otherwise.

### GetEgressRulesOk

`func (o *HypervisorVpcResourceResponseModel) GetEgressRulesOk() (*[]HypervisorSecurityGroupRuleResponseModel, bool)`

GetEgressRulesOk returns a tuple with the EgressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRules

`func (o *HypervisorVpcResourceResponseModel) SetEgressRules(v []HypervisorSecurityGroupRuleResponseModel)`

SetEgressRules sets EgressRules field to given value.


### GetIngressRules

`func (o *HypervisorVpcResourceResponseModel) GetIngressRules() []HypervisorSecurityGroupRuleResponseModel`

GetIngressRules returns the IngressRules field if non-nil, zero value otherwise.

### GetIngressRulesOk

`func (o *HypervisorVpcResourceResponseModel) GetIngressRulesOk() (*[]HypervisorSecurityGroupRuleResponseModel, bool)`

GetIngressRulesOk returns a tuple with the IngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRules

`func (o *HypervisorVpcResourceResponseModel) SetIngressRules(v []HypervisorSecurityGroupRuleResponseModel)`

SetIngressRules sets IngressRules field to given value.


### GetVirtualPrivateCloudId

`func (o *HypervisorVpcResourceResponseModel) GetVirtualPrivateCloudId() string`

GetVirtualPrivateCloudId returns the VirtualPrivateCloudId field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudIdOk

`func (o *HypervisorVpcResourceResponseModel) GetVirtualPrivateCloudIdOk() (*string, bool)`

GetVirtualPrivateCloudIdOk returns a tuple with the VirtualPrivateCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloudId

`func (o *HypervisorVpcResourceResponseModel) SetVirtualPrivateCloudId(v string)`

SetVirtualPrivateCloudId sets VirtualPrivateCloudId field to given value.

### HasVirtualPrivateCloudId

`func (o *HypervisorVpcResourceResponseModel) HasVirtualPrivateCloudId() bool`

HasVirtualPrivateCloudId returns a boolean if a field has been set.

### SetVirtualPrivateCloudIdNil

`func (o *HypervisorVpcResourceResponseModel) SetVirtualPrivateCloudIdNil(b bool)`

 SetVirtualPrivateCloudIdNil sets the value for VirtualPrivateCloudId to be an explicit nil

### UnsetVirtualPrivateCloudId
`func (o *HypervisorVpcResourceResponseModel) UnsetVirtualPrivateCloudId()`

UnsetVirtualPrivateCloudId ensures that no value is present for VirtualPrivateCloudId, not even an explicit nil
### GetDedicatedTenancy

`func (o *HypervisorVpcResourceResponseModel) GetDedicatedTenancy() string`

GetDedicatedTenancy returns the DedicatedTenancy field if non-nil, zero value otherwise.

### GetDedicatedTenancyOk

`func (o *HypervisorVpcResourceResponseModel) GetDedicatedTenancyOk() (*string, bool)`

GetDedicatedTenancyOk returns a tuple with the DedicatedTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancy

`func (o *HypervisorVpcResourceResponseModel) SetDedicatedTenancy(v string)`

SetDedicatedTenancy sets DedicatedTenancy field to given value.

### HasDedicatedTenancy

`func (o *HypervisorVpcResourceResponseModel) HasDedicatedTenancy() bool`

HasDedicatedTenancy returns a boolean if a field has been set.

### SetDedicatedTenancyNil

`func (o *HypervisorVpcResourceResponseModel) SetDedicatedTenancyNil(b bool)`

 SetDedicatedTenancyNil sets the value for DedicatedTenancy to be an explicit nil

### UnsetDedicatedTenancy
`func (o *HypervisorVpcResourceResponseModel) UnsetDedicatedTenancy()`

UnsetDedicatedTenancy ensures that no value is present for DedicatedTenancy, not even an explicit nil
### GetMemorySizeMB

`func (o *HypervisorVpcResourceResponseModel) GetMemorySizeMB() float32`

GetMemorySizeMB returns the MemorySizeMB field if non-nil, zero value otherwise.

### GetMemorySizeMBOk

`func (o *HypervisorVpcResourceResponseModel) GetMemorySizeMBOk() (*float32, bool)`

GetMemorySizeMBOk returns a tuple with the MemorySizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeMB

`func (o *HypervisorVpcResourceResponseModel) SetMemorySizeMB(v float32)`

SetMemorySizeMB sets MemorySizeMB field to given value.


### GetNumberOfCores

`func (o *HypervisorVpcResourceResponseModel) GetNumberOfCores() float32`

GetNumberOfCores returns the NumberOfCores field if non-nil, zero value otherwise.

### GetNumberOfCoresOk

`func (o *HypervisorVpcResourceResponseModel) GetNumberOfCoresOk() (*float32, bool)`

GetNumberOfCoresOk returns a tuple with the NumberOfCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCores

`func (o *HypervisorVpcResourceResponseModel) SetNumberOfCores(v float32)`

SetNumberOfCores sets NumberOfCores field to given value.


### GetAmazonComputeUnits

`func (o *HypervisorVpcResourceResponseModel) GetAmazonComputeUnits() float32`

GetAmazonComputeUnits returns the AmazonComputeUnits field if non-nil, zero value otherwise.

### GetAmazonComputeUnitsOk

`func (o *HypervisorVpcResourceResponseModel) GetAmazonComputeUnitsOk() (*float32, bool)`

GetAmazonComputeUnitsOk returns a tuple with the AmazonComputeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonComputeUnits

`func (o *HypervisorVpcResourceResponseModel) SetAmazonComputeUnits(v float32)`

SetAmazonComputeUnits sets AmazonComputeUnits field to given value.

### HasAmazonComputeUnits

`func (o *HypervisorVpcResourceResponseModel) HasAmazonComputeUnits() bool`

HasAmazonComputeUnits returns a boolean if a field has been set.

### SetAmazonComputeUnitsNil

`func (o *HypervisorVpcResourceResponseModel) SetAmazonComputeUnitsNil(b bool)`

 SetAmazonComputeUnitsNil sets the value for AmazonComputeUnits to be an explicit nil

### UnsetAmazonComputeUnits
`func (o *HypervisorVpcResourceResponseModel) UnsetAmazonComputeUnits()`

UnsetAmazonComputeUnits ensures that no value is present for AmazonComputeUnits, not even an explicit nil
### GetOptimizedForPooledDesktops

`func (o *HypervisorVpcResourceResponseModel) GetOptimizedForPooledDesktops() bool`

GetOptimizedForPooledDesktops returns the OptimizedForPooledDesktops field if non-nil, zero value otherwise.

### GetOptimizedForPooledDesktopsOk

`func (o *HypervisorVpcResourceResponseModel) GetOptimizedForPooledDesktopsOk() (*bool, bool)`

GetOptimizedForPooledDesktopsOk returns a tuple with the OptimizedForPooledDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizedForPooledDesktops

`func (o *HypervisorVpcResourceResponseModel) SetOptimizedForPooledDesktops(v bool)`

SetOptimizedForPooledDesktops sets OptimizedForPooledDesktops field to given value.


### GetNetworkPerformance

`func (o *HypervisorVpcResourceResponseModel) GetNetworkPerformance() string`

GetNetworkPerformance returns the NetworkPerformance field if non-nil, zero value otherwise.

### GetNetworkPerformanceOk

`func (o *HypervisorVpcResourceResponseModel) GetNetworkPerformanceOk() (*string, bool)`

GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformance

`func (o *HypervisorVpcResourceResponseModel) SetNetworkPerformance(v string)`

SetNetworkPerformance sets NetworkPerformance field to given value.

### HasNetworkPerformance

`func (o *HypervisorVpcResourceResponseModel) HasNetworkPerformance() bool`

HasNetworkPerformance returns a boolean if a field has been set.

### SetNetworkPerformanceNil

`func (o *HypervisorVpcResourceResponseModel) SetNetworkPerformanceNil(b bool)`

 SetNetworkPerformanceNil sets the value for NetworkPerformance to be an explicit nil

### UnsetNetworkPerformance
`func (o *HypervisorVpcResourceResponseModel) UnsetNetworkPerformance()`

UnsetNetworkPerformance ensures that no value is present for NetworkPerformance, not even an explicit nil
### GetNetworkPerformanceIsUnlimited

`func (o *HypervisorVpcResourceResponseModel) GetNetworkPerformanceIsUnlimited() bool`

GetNetworkPerformanceIsUnlimited returns the NetworkPerformanceIsUnlimited field if non-nil, zero value otherwise.

### GetNetworkPerformanceIsUnlimitedOk

`func (o *HypervisorVpcResourceResponseModel) GetNetworkPerformanceIsUnlimitedOk() (*bool, bool)`

GetNetworkPerformanceIsUnlimitedOk returns a tuple with the NetworkPerformanceIsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformanceIsUnlimited

`func (o *HypervisorVpcResourceResponseModel) SetNetworkPerformanceIsUnlimited(v bool)`

SetNetworkPerformanceIsUnlimited sets NetworkPerformanceIsUnlimited field to given value.


### GetNetworkPerformanceIsDefault

`func (o *HypervisorVpcResourceResponseModel) GetNetworkPerformanceIsDefault() bool`

GetNetworkPerformanceIsDefault returns the NetworkPerformanceIsDefault field if non-nil, zero value otherwise.

### GetNetworkPerformanceIsDefaultOk

`func (o *HypervisorVpcResourceResponseModel) GetNetworkPerformanceIsDefaultOk() (*bool, bool)`

GetNetworkPerformanceIsDefaultOk returns a tuple with the NetworkPerformanceIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformanceIsDefault

`func (o *HypervisorVpcResourceResponseModel) SetNetworkPerformanceIsDefault(v bool)`

SetNetworkPerformanceIsDefault sets NetworkPerformanceIsDefault field to given value.


### GetIsUsedInSite

`func (o *HypervisorVpcResourceResponseModel) GetIsUsedInSite() bool`

GetIsUsedInSite returns the IsUsedInSite field if non-nil, zero value otherwise.

### GetIsUsedInSiteOk

`func (o *HypervisorVpcResourceResponseModel) GetIsUsedInSiteOk() (*bool, bool)`

GetIsUsedInSiteOk returns a tuple with the IsUsedInSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedInSite

`func (o *HypervisorVpcResourceResponseModel) SetIsUsedInSite(v bool)`

SetIsUsedInSite sets IsUsedInSite field to given value.


### GetSupportsAzurePremiumStorage

`func (o *HypervisorVpcResourceResponseModel) GetSupportsAzurePremiumStorage() bool`

GetSupportsAzurePremiumStorage returns the SupportsAzurePremiumStorage field if non-nil, zero value otherwise.

### GetSupportsAzurePremiumStorageOk

`func (o *HypervisorVpcResourceResponseModel) GetSupportsAzurePremiumStorageOk() (*bool, bool)`

GetSupportsAzurePremiumStorageOk returns a tuple with the SupportsAzurePremiumStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAzurePremiumStorage

`func (o *HypervisorVpcResourceResponseModel) SetSupportsAzurePremiumStorage(v bool)`

SetSupportsAzurePremiumStorage sets SupportsAzurePremiumStorage field to given value.

### HasSupportsAzurePremiumStorage

`func (o *HypervisorVpcResourceResponseModel) HasSupportsAzurePremiumStorage() bool`

HasSupportsAzurePremiumStorage returns a boolean if a field has been set.

### SetSupportsAzurePremiumStorageNil

`func (o *HypervisorVpcResourceResponseModel) SetSupportsAzurePremiumStorageNil(b bool)`

 SetSupportsAzurePremiumStorageNil sets the value for SupportsAzurePremiumStorage to be an explicit nil

### UnsetSupportsAzurePremiumStorage
`func (o *HypervisorVpcResourceResponseModel) UnsetSupportsAzurePremiumStorage()`

UnsetSupportsAzurePremiumStorage ensures that no value is present for SupportsAzurePremiumStorage, not even an explicit nil
### GetSuperseded

`func (o *HypervisorVpcResourceResponseModel) GetSuperseded() bool`

GetSuperseded returns the Superseded field if non-nil, zero value otherwise.

### GetSupersededOk

`func (o *HypervisorVpcResourceResponseModel) GetSupersededOk() (*bool, bool)`

GetSupersededOk returns a tuple with the Superseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperseded

`func (o *HypervisorVpcResourceResponseModel) SetSuperseded(v bool)`

SetSuperseded sets Superseded field to given value.


### GetOwner

`func (o *HypervisorVpcResourceResponseModel) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HypervisorVpcResourceResponseModel) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HypervisorVpcResourceResponseModel) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *HypervisorVpcResourceResponseModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *HypervisorVpcResourceResponseModel) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *HypervisorVpcResourceResponseModel) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetIsWindowsTemplate

`func (o *HypervisorVpcResourceResponseModel) GetIsWindowsTemplate() bool`

GetIsWindowsTemplate returns the IsWindowsTemplate field if non-nil, zero value otherwise.

### GetIsWindowsTemplateOk

`func (o *HypervisorVpcResourceResponseModel) GetIsWindowsTemplateOk() (*bool, bool)`

GetIsWindowsTemplateOk returns a tuple with the IsWindowsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindowsTemplate

`func (o *HypervisorVpcResourceResponseModel) SetIsWindowsTemplate(v bool)`

SetIsWindowsTemplate sets IsWindowsTemplate field to given value.

### HasIsWindowsTemplate

`func (o *HypervisorVpcResourceResponseModel) HasIsWindowsTemplate() bool`

HasIsWindowsTemplate returns a boolean if a field has been set.

### SetIsWindowsTemplateNil

`func (o *HypervisorVpcResourceResponseModel) SetIsWindowsTemplateNil(b bool)`

 SetIsWindowsTemplateNil sets the value for IsWindowsTemplate to be an explicit nil

### UnsetIsWindowsTemplate
`func (o *HypervisorVpcResourceResponseModel) UnsetIsWindowsTemplate()`

UnsetIsWindowsTemplate ensures that no value is present for IsWindowsTemplate, not even an explicit nil
### GetHasPersistentRootVolume

`func (o *HypervisorVpcResourceResponseModel) GetHasPersistentRootVolume() bool`

GetHasPersistentRootVolume returns the HasPersistentRootVolume field if non-nil, zero value otherwise.

### GetHasPersistentRootVolumeOk

`func (o *HypervisorVpcResourceResponseModel) GetHasPersistentRootVolumeOk() (*bool, bool)`

GetHasPersistentRootVolumeOk returns a tuple with the HasPersistentRootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPersistentRootVolume

`func (o *HypervisorVpcResourceResponseModel) SetHasPersistentRootVolume(v bool)`

SetHasPersistentRootVolume sets HasPersistentRootVolume field to given value.


### GetVMId

`func (o *HypervisorVpcResourceResponseModel) GetVMId() string`

GetVMId returns the VMId field if non-nil, zero value otherwise.

### GetVMIdOk

`func (o *HypervisorVpcResourceResponseModel) GetVMIdOk() (*string, bool)`

GetVMIdOk returns a tuple with the VMId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMId

`func (o *HypervisorVpcResourceResponseModel) SetVMId(v string)`

SetVMId sets VMId field to given value.


### GetMacAddress

`func (o *HypervisorVpcResourceResponseModel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *HypervisorVpcResourceResponseModel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *HypervisorVpcResourceResponseModel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *HypervisorVpcResourceResponseModel) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *HypervisorVpcResourceResponseModel) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *HypervisorVpcResourceResponseModel) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetCpuCount

`func (o *HypervisorVpcResourceResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *HypervisorVpcResourceResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *HypervisorVpcResourceResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *HypervisorVpcResourceResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *HypervisorVpcResourceResponseModel) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *HypervisorVpcResourceResponseModel) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetMemoryMB

`func (o *HypervisorVpcResourceResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *HypervisorVpcResourceResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *HypervisorVpcResourceResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *HypervisorVpcResourceResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### SetMemoryMBNil

`func (o *HypervisorVpcResourceResponseModel) SetMemoryMBNil(b bool)`

 SetMemoryMBNil sets the value for MemoryMB to be an explicit nil

### UnsetMemoryMB
`func (o *HypervisorVpcResourceResponseModel) UnsetMemoryMB()`

UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
### GetHardDiskSizeGB

`func (o *HypervisorVpcResourceResponseModel) GetHardDiskSizeGB() int32`

GetHardDiskSizeGB returns the HardDiskSizeGB field if non-nil, zero value otherwise.

### GetHardDiskSizeGBOk

`func (o *HypervisorVpcResourceResponseModel) GetHardDiskSizeGBOk() (*int32, bool)`

GetHardDiskSizeGBOk returns a tuple with the HardDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDiskSizeGB

`func (o *HypervisorVpcResourceResponseModel) SetHardDiskSizeGB(v int32)`

SetHardDiskSizeGB sets HardDiskSizeGB field to given value.

### HasHardDiskSizeGB

`func (o *HypervisorVpcResourceResponseModel) HasHardDiskSizeGB() bool`

HasHardDiskSizeGB returns a boolean if a field has been set.

### SetHardDiskSizeGBNil

`func (o *HypervisorVpcResourceResponseModel) SetHardDiskSizeGBNil(b bool)`

 SetHardDiskSizeGBNil sets the value for HardDiskSizeGB to be an explicit nil

### UnsetHardDiskSizeGB
`func (o *HypervisorVpcResourceResponseModel) UnsetHardDiskSizeGB()`

UnsetHardDiskSizeGB ensures that no value is present for HardDiskSizeGB, not even an explicit nil
### GetMinMemoryMB

`func (o *HypervisorVpcResourceResponseModel) GetMinMemoryMB() int32`

GetMinMemoryMB returns the MinMemoryMB field if non-nil, zero value otherwise.

### GetMinMemoryMBOk

`func (o *HypervisorVpcResourceResponseModel) GetMinMemoryMBOk() (*int32, bool)`

GetMinMemoryMBOk returns a tuple with the MinMemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemoryMB

`func (o *HypervisorVpcResourceResponseModel) SetMinMemoryMB(v int32)`

SetMinMemoryMB sets MinMemoryMB field to given value.

### HasMinMemoryMB

`func (o *HypervisorVpcResourceResponseModel) HasMinMemoryMB() bool`

HasMinMemoryMB returns a boolean if a field has been set.

### SetMinMemoryMBNil

`func (o *HypervisorVpcResourceResponseModel) SetMinMemoryMBNil(b bool)`

 SetMinMemoryMBNil sets the value for MinMemoryMB to be an explicit nil

### UnsetMinMemoryMB
`func (o *HypervisorVpcResourceResponseModel) UnsetMinMemoryMB()`

UnsetMinMemoryMB ensures that no value is present for MinMemoryMB, not even an explicit nil
### GetNetworkMappings

`func (o *HypervisorVpcResourceResponseModel) GetNetworkMappings() []NetworkMapResponseModel`

GetNetworkMappings returns the NetworkMappings field if non-nil, zero value otherwise.

### GetNetworkMappingsOk

`func (o *HypervisorVpcResourceResponseModel) GetNetworkMappingsOk() (*[]NetworkMapResponseModel, bool)`

GetNetworkMappingsOk returns a tuple with the NetworkMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappings

`func (o *HypervisorVpcResourceResponseModel) SetNetworkMappings(v []NetworkMapResponseModel)`

SetNetworkMappings sets NetworkMappings field to given value.

### HasNetworkMappings

`func (o *HypervisorVpcResourceResponseModel) HasNetworkMappings() bool`

HasNetworkMappings returns a boolean if a field has been set.

### SetNetworkMappingsNil

`func (o *HypervisorVpcResourceResponseModel) SetNetworkMappingsNil(b bool)`

 SetNetworkMappingsNil sets the value for NetworkMappings to be an explicit nil

### UnsetNetworkMappings
`func (o *HypervisorVpcResourceResponseModel) UnsetNetworkMappings()`

UnsetNetworkMappings ensures that no value is present for NetworkMappings, not even an explicit nil
### GetAttachedDisks

`func (o *HypervisorVpcResourceResponseModel) GetAttachedDisks() []AttachedDiskResponseModel`

GetAttachedDisks returns the AttachedDisks field if non-nil, zero value otherwise.

### GetAttachedDisksOk

`func (o *HypervisorVpcResourceResponseModel) GetAttachedDisksOk() (*[]AttachedDiskResponseModel, bool)`

GetAttachedDisksOk returns a tuple with the AttachedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedDisks

`func (o *HypervisorVpcResourceResponseModel) SetAttachedDisks(v []AttachedDiskResponseModel)`

SetAttachedDisks sets AttachedDisks field to given value.

### HasAttachedDisks

`func (o *HypervisorVpcResourceResponseModel) HasAttachedDisks() bool`

HasAttachedDisks returns a boolean if a field has been set.

### SetAttachedDisksNil

`func (o *HypervisorVpcResourceResponseModel) SetAttachedDisksNil(b bool)`

 SetAttachedDisksNil sets the value for AttachedDisks to be an explicit nil

### UnsetAttachedDisks
`func (o *HypervisorVpcResourceResponseModel) UnsetAttachedDisks()`

UnsetAttachedDisks ensures that no value is present for AttachedDisks, not even an explicit nil
### GetInstanceSecurityGroupLimit

`func (o *HypervisorVpcResourceResponseModel) GetInstanceSecurityGroupLimit() int32`

GetInstanceSecurityGroupLimit returns the InstanceSecurityGroupLimit field if non-nil, zero value otherwise.

### GetInstanceSecurityGroupLimitOk

`func (o *HypervisorVpcResourceResponseModel) GetInstanceSecurityGroupLimitOk() (*int32, bool)`

GetInstanceSecurityGroupLimitOk returns a tuple with the InstanceSecurityGroupLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSecurityGroupLimit

`func (o *HypervisorVpcResourceResponseModel) SetInstanceSecurityGroupLimit(v int32)`

SetInstanceSecurityGroupLimit sets InstanceSecurityGroupLimit field to given value.


### GetEndpoint

`func (o *HypervisorVpcResourceResponseModel) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HypervisorVpcResourceResponseModel) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HypervisorVpcResourceResponseModel) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HypervisorVpcResourceResponseModel) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *HypervisorVpcResourceResponseModel) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *HypervisorVpcResourceResponseModel) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetId

`func (o *HypervisorVpcResourceResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorVpcResourceResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorVpcResourceResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorVpcResourceResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorVpcResourceResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorVpcResourceResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorVpcResourceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorVpcResourceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorVpcResourceResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorVpcResourceResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorVpcResourceResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorVpcResourceResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorVpcResourceResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorVpcResourceResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorVpcResourceResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorVpcResourceResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorVpcResourceResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorVpcResourceResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetRelativePath

`func (o *HypervisorVpcResourceResponseModel) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HypervisorVpcResourceResponseModel) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HypervisorVpcResourceResponseModel) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *HypervisorVpcResourceResponseModel) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *HypervisorVpcResourceResponseModel) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *HypervisorVpcResourceResponseModel) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
### GetFullRelativePath

`func (o *HypervisorVpcResourceResponseModel) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorVpcResourceResponseModel) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorVpcResourceResponseModel) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.

### HasFullRelativePath

`func (o *HypervisorVpcResourceResponseModel) HasFullRelativePath() bool`

HasFullRelativePath returns a boolean if a field has been set.

### SetFullRelativePathNil

`func (o *HypervisorVpcResourceResponseModel) SetFullRelativePathNil(b bool)`

 SetFullRelativePathNil sets the value for FullRelativePath to be an explicit nil

### UnsetFullRelativePath
`func (o *HypervisorVpcResourceResponseModel) UnsetFullRelativePath()`

UnsetFullRelativePath ensures that no value is present for FullRelativePath, not even an explicit nil
### GetResourceType

`func (o *HypervisorVpcResourceResponseModel) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorVpcResourceResponseModel) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorVpcResourceResponseModel) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetObjectTypeName

`func (o *HypervisorVpcResourceResponseModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *HypervisorVpcResourceResponseModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *HypervisorVpcResourceResponseModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *HypervisorVpcResourceResponseModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### SetObjectTypeNameNil

`func (o *HypervisorVpcResourceResponseModel) SetObjectTypeNameNil(b bool)`

 SetObjectTypeNameNil sets the value for ObjectTypeName to be an explicit nil

### UnsetObjectTypeName
`func (o *HypervisorVpcResourceResponseModel) UnsetObjectTypeName()`

UnsetObjectTypeName ensures that no value is present for ObjectTypeName, not even an explicit nil
### GetResourceContainer

`func (o *HypervisorVpcResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModel`

GetResourceContainer returns the ResourceContainer field if non-nil, zero value otherwise.

### GetResourceContainerOk

`func (o *HypervisorVpcResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModel, bool)`

GetResourceContainerOk returns a tuple with the ResourceContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceContainer

`func (o *HypervisorVpcResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModel)`

SetResourceContainer sets ResourceContainer field to given value.

### HasResourceContainer

`func (o *HypervisorVpcResourceResponseModel) HasResourceContainer() bool`

HasResourceContainer returns a boolean if a field has been set.

### GetResourcePoolXDPath

`func (o *HypervisorVpcResourceResponseModel) GetResourcePoolXDPath() string`

GetResourcePoolXDPath returns the ResourcePoolXDPath field if non-nil, zero value otherwise.

### GetResourcePoolXDPathOk

`func (o *HypervisorVpcResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool)`

GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolXDPath

`func (o *HypervisorVpcResourceResponseModel) SetResourcePoolXDPath(v string)`

SetResourcePoolXDPath sets ResourcePoolXDPath field to given value.

### HasResourcePoolXDPath

`func (o *HypervisorVpcResourceResponseModel) HasResourcePoolXDPath() bool`

HasResourcePoolXDPath returns a boolean if a field has been set.

### SetResourcePoolXDPathNil

`func (o *HypervisorVpcResourceResponseModel) SetResourcePoolXDPathNil(b bool)`

 SetResourcePoolXDPathNil sets the value for ResourcePoolXDPath to be an explicit nil

### UnsetResourcePoolXDPath
`func (o *HypervisorVpcResourceResponseModel) UnsetResourcePoolXDPath()`

UnsetResourcePoolXDPath ensures that no value is present for ResourcePoolXDPath, not even an explicit nil
### GetFullName

`func (o *HypervisorVpcResourceResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *HypervisorVpcResourceResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *HypervisorVpcResourceResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsContainer

`func (o *HypervisorVpcResourceResponseModel) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *HypervisorVpcResourceResponseModel) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *HypervisorVpcResourceResponseModel) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.


### GetChildren

`func (o *HypervisorVpcResourceResponseModel) GetChildren() []HypervisorResourceResponseModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HypervisorVpcResourceResponseModel) GetChildrenOk() (*[]HypervisorResourceResponseModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HypervisorVpcResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HypervisorVpcResourceResponseModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *HypervisorVpcResourceResponseModel) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *HypervisorVpcResourceResponseModel) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetIsMachine

`func (o *HypervisorVpcResourceResponseModel) GetIsMachine() bool`

GetIsMachine returns the IsMachine field if non-nil, zero value otherwise.

### GetIsMachineOk

`func (o *HypervisorVpcResourceResponseModel) GetIsMachineOk() (*bool, bool)`

GetIsMachineOk returns a tuple with the IsMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachine

`func (o *HypervisorVpcResourceResponseModel) SetIsMachine(v bool)`

SetIsMachine sets IsMachine field to given value.


### GetIsSnapshotable

`func (o *HypervisorVpcResourceResponseModel) GetIsSnapshotable() bool`

GetIsSnapshotable returns the IsSnapshotable field if non-nil, zero value otherwise.

### GetIsSnapshotableOk

`func (o *HypervisorVpcResourceResponseModel) GetIsSnapshotableOk() (*bool, bool)`

GetIsSnapshotableOk returns a tuple with the IsSnapshotable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotable

`func (o *HypervisorVpcResourceResponseModel) SetIsSnapshotable(v bool)`

SetIsSnapshotable sets IsSnapshotable field to given value.


### GetAllResourcesRelativePath

`func (o *HypervisorVpcResourceResponseModel) GetAllResourcesRelativePath() string`

GetAllResourcesRelativePath returns the AllResourcesRelativePath field if non-nil, zero value otherwise.

### GetAllResourcesRelativePathOk

`func (o *HypervisorVpcResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool)`

GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllResourcesRelativePath

`func (o *HypervisorVpcResourceResponseModel) SetAllResourcesRelativePath(v string)`

SetAllResourcesRelativePath sets AllResourcesRelativePath field to given value.


### GetResourcePool

`func (o *HypervisorVpcResourceResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *HypervisorVpcResourceResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *HypervisorVpcResourceResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel)`

SetResourcePool sets ResourcePool field to given value.


### GetIsSymLink

`func (o *HypervisorVpcResourceResponseModel) GetIsSymLink() bool`

GetIsSymLink returns the IsSymLink field if non-nil, zero value otherwise.

### GetIsSymLinkOk

`func (o *HypervisorVpcResourceResponseModel) GetIsSymLinkOk() (*bool, bool)`

GetIsSymLinkOk returns a tuple with the IsSymLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSymLink

`func (o *HypervisorVpcResourceResponseModel) SetIsSymLink(v bool)`

SetIsSymLink sets IsSymLink field to given value.


### GetAdditionalData

`func (o *HypervisorVpcResourceResponseModel) GetAdditionalData() []NameValueStringPairModel`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *HypervisorVpcResourceResponseModel) GetAdditionalDataOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *HypervisorVpcResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *HypervisorVpcResourceResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *HypervisorVpcResourceResponseModel) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *HypervisorVpcResourceResponseModel) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


