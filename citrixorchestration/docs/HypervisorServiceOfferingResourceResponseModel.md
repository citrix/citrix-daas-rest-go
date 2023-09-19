# HypervisorServiceOfferingResourceResponseModel

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
**DedicatedTenancy** | Pointer to **NullableString** | Indicates whether this service offering prescribes dedicated tenancy for the instances that use it. | [optional] 
**Description** | Pointer to **NullableString** | A human-readable description of this service offering, as supplied when the offering was created in the cloud management API or console. | [optional] 
**MemorySizeMB** | **float32** | The amount of memory that is available to instances using this service offering, measured in megabytes. | 
**NumberOfCores** | **float32** | The number of virtual cores that are available to instances using this service offering. | 
**AmazonComputeUnits** | Pointer to **NullableFloat32** | The number of Amazon EC2 compute units that are available to instances using this service offering. This property is only meaningful for EC2-based cloud connections. | [optional] 
**OptimizedForPooledDesktops** | **bool** | Indicates whether this service offering is intended specifically for usage with pooled desktops. | 
**NetworkPerformance** | Pointer to **NullableString** | The network performance available to instances using this service offering. | [optional] 
**NetworkPerformanceIsUnlimited** | **bool** | Indicates whether this service offering has no limit for it&#39;s network rate. | 
**NetworkPerformanceIsDefault** | **bool** | Indicates whether this service offering has a default limit for it&#39;s network rate. | 
**IsUsedInSite** | **bool** | Indicates whether the service offering is used by any machine catalogs in the site. | 
**SupportsAzurePremiumStorage** | Pointer to **NullableBool** | Indicates whether the service offering supports premium storage. This property is only meaningful on Azure. | [optional] 

## Methods

### NewHypervisorServiceOfferingResourceResponseModel

`func NewHypervisorServiceOfferingResourceResponseModel(resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourcePoolRefResponseModel, isSymLink bool, memorySizeMB float32, numberOfCores float32, optimizedForPooledDesktops bool, networkPerformanceIsUnlimited bool, networkPerformanceIsDefault bool, isUsedInSite bool, ) *HypervisorServiceOfferingResourceResponseModel`

NewHypervisorServiceOfferingResourceResponseModel instantiates a new HypervisorServiceOfferingResourceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorServiceOfferingResourceResponseModelWithDefaults

`func NewHypervisorServiceOfferingResourceResponseModelWithDefaults() *HypervisorServiceOfferingResourceResponseModel`

NewHypervisorServiceOfferingResourceResponseModelWithDefaults instantiates a new HypervisorServiceOfferingResourceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorServiceOfferingResourceResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorServiceOfferingResourceResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorServiceOfferingResourceResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorServiceOfferingResourceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorServiceOfferingResourceResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorServiceOfferingResourceResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorServiceOfferingResourceResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorServiceOfferingResourceResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorServiceOfferingResourceResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetRelativePath

`func (o *HypervisorServiceOfferingResourceResponseModel) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HypervisorServiceOfferingResourceResponseModel) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *HypervisorServiceOfferingResourceResponseModel) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
### GetFullRelativePath

`func (o *HypervisorServiceOfferingResourceResponseModel) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorServiceOfferingResourceResponseModel) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.

### HasFullRelativePath

`func (o *HypervisorServiceOfferingResourceResponseModel) HasFullRelativePath() bool`

HasFullRelativePath returns a boolean if a field has been set.

### SetFullRelativePathNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetFullRelativePathNil(b bool)`

 SetFullRelativePathNil sets the value for FullRelativePath to be an explicit nil

### UnsetFullRelativePath
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetFullRelativePath()`

UnsetFullRelativePath ensures that no value is present for FullRelativePath, not even an explicit nil
### GetResourceType

`func (o *HypervisorServiceOfferingResourceResponseModel) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorServiceOfferingResourceResponseModel) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetObjectTypeName

`func (o *HypervisorServiceOfferingResourceResponseModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *HypervisorServiceOfferingResourceResponseModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *HypervisorServiceOfferingResourceResponseModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### SetObjectTypeNameNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetObjectTypeNameNil(b bool)`

 SetObjectTypeNameNil sets the value for ObjectTypeName to be an explicit nil

### UnsetObjectTypeName
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetObjectTypeName()`

UnsetObjectTypeName ensures that no value is present for ObjectTypeName, not even an explicit nil
### GetResourceContainer

`func (o *HypervisorServiceOfferingResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModel`

GetResourceContainer returns the ResourceContainer field if non-nil, zero value otherwise.

### GetResourceContainerOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModel, bool)`

GetResourceContainerOk returns a tuple with the ResourceContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceContainer

`func (o *HypervisorServiceOfferingResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModel)`

SetResourceContainer sets ResourceContainer field to given value.

### HasResourceContainer

`func (o *HypervisorServiceOfferingResourceResponseModel) HasResourceContainer() bool`

HasResourceContainer returns a boolean if a field has been set.

### GetResourcePoolXDPath

`func (o *HypervisorServiceOfferingResourceResponseModel) GetResourcePoolXDPath() string`

GetResourcePoolXDPath returns the ResourcePoolXDPath field if non-nil, zero value otherwise.

### GetResourcePoolXDPathOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool)`

GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolXDPath

`func (o *HypervisorServiceOfferingResourceResponseModel) SetResourcePoolXDPath(v string)`

SetResourcePoolXDPath sets ResourcePoolXDPath field to given value.

### HasResourcePoolXDPath

`func (o *HypervisorServiceOfferingResourceResponseModel) HasResourcePoolXDPath() bool`

HasResourcePoolXDPath returns a boolean if a field has been set.

### SetResourcePoolXDPathNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetResourcePoolXDPathNil(b bool)`

 SetResourcePoolXDPathNil sets the value for ResourcePoolXDPath to be an explicit nil

### UnsetResourcePoolXDPath
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetResourcePoolXDPath()`

UnsetResourcePoolXDPath ensures that no value is present for ResourcePoolXDPath, not even an explicit nil
### GetFullName

`func (o *HypervisorServiceOfferingResourceResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *HypervisorServiceOfferingResourceResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsContainer

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *HypervisorServiceOfferingResourceResponseModel) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.


### GetChildren

`func (o *HypervisorServiceOfferingResourceResponseModel) GetChildren() []HypervisorResourceResponseModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetChildrenOk() (*[]HypervisorResourceResponseModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HypervisorServiceOfferingResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HypervisorServiceOfferingResourceResponseModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetIsMachine

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsMachine() bool`

GetIsMachine returns the IsMachine field if non-nil, zero value otherwise.

### GetIsMachineOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsMachineOk() (*bool, bool)`

GetIsMachineOk returns a tuple with the IsMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachine

`func (o *HypervisorServiceOfferingResourceResponseModel) SetIsMachine(v bool)`

SetIsMachine sets IsMachine field to given value.


### GetIsSnapshotable

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsSnapshotable() bool`

GetIsSnapshotable returns the IsSnapshotable field if non-nil, zero value otherwise.

### GetIsSnapshotableOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsSnapshotableOk() (*bool, bool)`

GetIsSnapshotableOk returns a tuple with the IsSnapshotable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotable

`func (o *HypervisorServiceOfferingResourceResponseModel) SetIsSnapshotable(v bool)`

SetIsSnapshotable sets IsSnapshotable field to given value.


### GetAllResourcesRelativePath

`func (o *HypervisorServiceOfferingResourceResponseModel) GetAllResourcesRelativePath() string`

GetAllResourcesRelativePath returns the AllResourcesRelativePath field if non-nil, zero value otherwise.

### GetAllResourcesRelativePathOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool)`

GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllResourcesRelativePath

`func (o *HypervisorServiceOfferingResourceResponseModel) SetAllResourcesRelativePath(v string)`

SetAllResourcesRelativePath sets AllResourcesRelativePath field to given value.


### GetResourcePool

`func (o *HypervisorServiceOfferingResourceResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *HypervisorServiceOfferingResourceResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel)`

SetResourcePool sets ResourcePool field to given value.


### GetIsSymLink

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsSymLink() bool`

GetIsSymLink returns the IsSymLink field if non-nil, zero value otherwise.

### GetIsSymLinkOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsSymLinkOk() (*bool, bool)`

GetIsSymLinkOk returns a tuple with the IsSymLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSymLink

`func (o *HypervisorServiceOfferingResourceResponseModel) SetIsSymLink(v bool)`

SetIsSymLink sets IsSymLink field to given value.


### GetAdditionalData

`func (o *HypervisorServiceOfferingResourceResponseModel) GetAdditionalData() []NameValueStringPairModel`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetAdditionalDataOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *HypervisorServiceOfferingResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *HypervisorServiceOfferingResourceResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil
### GetDedicatedTenancy

`func (o *HypervisorServiceOfferingResourceResponseModel) GetDedicatedTenancy() string`

GetDedicatedTenancy returns the DedicatedTenancy field if non-nil, zero value otherwise.

### GetDedicatedTenancyOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetDedicatedTenancyOk() (*string, bool)`

GetDedicatedTenancyOk returns a tuple with the DedicatedTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancy

`func (o *HypervisorServiceOfferingResourceResponseModel) SetDedicatedTenancy(v string)`

SetDedicatedTenancy sets DedicatedTenancy field to given value.

### HasDedicatedTenancy

`func (o *HypervisorServiceOfferingResourceResponseModel) HasDedicatedTenancy() bool`

HasDedicatedTenancy returns a boolean if a field has been set.

### SetDedicatedTenancyNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetDedicatedTenancyNil(b bool)`

 SetDedicatedTenancyNil sets the value for DedicatedTenancy to be an explicit nil

### UnsetDedicatedTenancy
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetDedicatedTenancy()`

UnsetDedicatedTenancy ensures that no value is present for DedicatedTenancy, not even an explicit nil
### GetDescription

`func (o *HypervisorServiceOfferingResourceResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HypervisorServiceOfferingResourceResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HypervisorServiceOfferingResourceResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMemorySizeMB

`func (o *HypervisorServiceOfferingResourceResponseModel) GetMemorySizeMB() float32`

GetMemorySizeMB returns the MemorySizeMB field if non-nil, zero value otherwise.

### GetMemorySizeMBOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetMemorySizeMBOk() (*float32, bool)`

GetMemorySizeMBOk returns a tuple with the MemorySizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeMB

`func (o *HypervisorServiceOfferingResourceResponseModel) SetMemorySizeMB(v float32)`

SetMemorySizeMB sets MemorySizeMB field to given value.


### GetNumberOfCores

`func (o *HypervisorServiceOfferingResourceResponseModel) GetNumberOfCores() float32`

GetNumberOfCores returns the NumberOfCores field if non-nil, zero value otherwise.

### GetNumberOfCoresOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetNumberOfCoresOk() (*float32, bool)`

GetNumberOfCoresOk returns a tuple with the NumberOfCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCores

`func (o *HypervisorServiceOfferingResourceResponseModel) SetNumberOfCores(v float32)`

SetNumberOfCores sets NumberOfCores field to given value.


### GetAmazonComputeUnits

`func (o *HypervisorServiceOfferingResourceResponseModel) GetAmazonComputeUnits() float32`

GetAmazonComputeUnits returns the AmazonComputeUnits field if non-nil, zero value otherwise.

### GetAmazonComputeUnitsOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetAmazonComputeUnitsOk() (*float32, bool)`

GetAmazonComputeUnitsOk returns a tuple with the AmazonComputeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonComputeUnits

`func (o *HypervisorServiceOfferingResourceResponseModel) SetAmazonComputeUnits(v float32)`

SetAmazonComputeUnits sets AmazonComputeUnits field to given value.

### HasAmazonComputeUnits

`func (o *HypervisorServiceOfferingResourceResponseModel) HasAmazonComputeUnits() bool`

HasAmazonComputeUnits returns a boolean if a field has been set.

### SetAmazonComputeUnitsNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetAmazonComputeUnitsNil(b bool)`

 SetAmazonComputeUnitsNil sets the value for AmazonComputeUnits to be an explicit nil

### UnsetAmazonComputeUnits
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetAmazonComputeUnits()`

UnsetAmazonComputeUnits ensures that no value is present for AmazonComputeUnits, not even an explicit nil
### GetOptimizedForPooledDesktops

`func (o *HypervisorServiceOfferingResourceResponseModel) GetOptimizedForPooledDesktops() bool`

GetOptimizedForPooledDesktops returns the OptimizedForPooledDesktops field if non-nil, zero value otherwise.

### GetOptimizedForPooledDesktopsOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetOptimizedForPooledDesktopsOk() (*bool, bool)`

GetOptimizedForPooledDesktopsOk returns a tuple with the OptimizedForPooledDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizedForPooledDesktops

`func (o *HypervisorServiceOfferingResourceResponseModel) SetOptimizedForPooledDesktops(v bool)`

SetOptimizedForPooledDesktops sets OptimizedForPooledDesktops field to given value.


### GetNetworkPerformance

`func (o *HypervisorServiceOfferingResourceResponseModel) GetNetworkPerformance() string`

GetNetworkPerformance returns the NetworkPerformance field if non-nil, zero value otherwise.

### GetNetworkPerformanceOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetNetworkPerformanceOk() (*string, bool)`

GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformance

`func (o *HypervisorServiceOfferingResourceResponseModel) SetNetworkPerformance(v string)`

SetNetworkPerformance sets NetworkPerformance field to given value.

### HasNetworkPerformance

`func (o *HypervisorServiceOfferingResourceResponseModel) HasNetworkPerformance() bool`

HasNetworkPerformance returns a boolean if a field has been set.

### SetNetworkPerformanceNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetNetworkPerformanceNil(b bool)`

 SetNetworkPerformanceNil sets the value for NetworkPerformance to be an explicit nil

### UnsetNetworkPerformance
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetNetworkPerformance()`

UnsetNetworkPerformance ensures that no value is present for NetworkPerformance, not even an explicit nil
### GetNetworkPerformanceIsUnlimited

`func (o *HypervisorServiceOfferingResourceResponseModel) GetNetworkPerformanceIsUnlimited() bool`

GetNetworkPerformanceIsUnlimited returns the NetworkPerformanceIsUnlimited field if non-nil, zero value otherwise.

### GetNetworkPerformanceIsUnlimitedOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetNetworkPerformanceIsUnlimitedOk() (*bool, bool)`

GetNetworkPerformanceIsUnlimitedOk returns a tuple with the NetworkPerformanceIsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformanceIsUnlimited

`func (o *HypervisorServiceOfferingResourceResponseModel) SetNetworkPerformanceIsUnlimited(v bool)`

SetNetworkPerformanceIsUnlimited sets NetworkPerformanceIsUnlimited field to given value.


### GetNetworkPerformanceIsDefault

`func (o *HypervisorServiceOfferingResourceResponseModel) GetNetworkPerformanceIsDefault() bool`

GetNetworkPerformanceIsDefault returns the NetworkPerformanceIsDefault field if non-nil, zero value otherwise.

### GetNetworkPerformanceIsDefaultOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetNetworkPerformanceIsDefaultOk() (*bool, bool)`

GetNetworkPerformanceIsDefaultOk returns a tuple with the NetworkPerformanceIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformanceIsDefault

`func (o *HypervisorServiceOfferingResourceResponseModel) SetNetworkPerformanceIsDefault(v bool)`

SetNetworkPerformanceIsDefault sets NetworkPerformanceIsDefault field to given value.


### GetIsUsedInSite

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsUsedInSite() bool`

GetIsUsedInSite returns the IsUsedInSite field if non-nil, zero value otherwise.

### GetIsUsedInSiteOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetIsUsedInSiteOk() (*bool, bool)`

GetIsUsedInSiteOk returns a tuple with the IsUsedInSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsedInSite

`func (o *HypervisorServiceOfferingResourceResponseModel) SetIsUsedInSite(v bool)`

SetIsUsedInSite sets IsUsedInSite field to given value.


### GetSupportsAzurePremiumStorage

`func (o *HypervisorServiceOfferingResourceResponseModel) GetSupportsAzurePremiumStorage() bool`

GetSupportsAzurePremiumStorage returns the SupportsAzurePremiumStorage field if non-nil, zero value otherwise.

### GetSupportsAzurePremiumStorageOk

`func (o *HypervisorServiceOfferingResourceResponseModel) GetSupportsAzurePremiumStorageOk() (*bool, bool)`

GetSupportsAzurePremiumStorageOk returns a tuple with the SupportsAzurePremiumStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAzurePremiumStorage

`func (o *HypervisorServiceOfferingResourceResponseModel) SetSupportsAzurePremiumStorage(v bool)`

SetSupportsAzurePremiumStorage sets SupportsAzurePremiumStorage field to given value.

### HasSupportsAzurePremiumStorage

`func (o *HypervisorServiceOfferingResourceResponseModel) HasSupportsAzurePremiumStorage() bool`

HasSupportsAzurePremiumStorage returns a boolean if a field has been set.

### SetSupportsAzurePremiumStorageNil

`func (o *HypervisorServiceOfferingResourceResponseModel) SetSupportsAzurePremiumStorageNil(b bool)`

 SetSupportsAzurePremiumStorageNil sets the value for SupportsAzurePremiumStorage to be an explicit nil

### UnsetSupportsAzurePremiumStorage
`func (o *HypervisorServiceOfferingResourceResponseModel) UnsetSupportsAzurePremiumStorage()`

UnsetSupportsAzurePremiumStorage ensures that no value is present for SupportsAzurePremiumStorage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


