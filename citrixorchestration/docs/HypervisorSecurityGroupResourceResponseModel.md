# HypervisorSecurityGroupResourceResponseModel

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
**Description** | Pointer to **string** | Description of the security group. | [optional] 
**EgressRules** | [**[]HypervisorSecurityGroupRuleResponseModel**](HypervisorSecurityGroupRuleResponseModel.md) | Egress rules governing outbound network traffic. | 
**IngressRules** | [**[]HypervisorSecurityGroupRuleResponseModel**](HypervisorSecurityGroupRuleResponseModel.md) | Ingress rules governing inbound network traffic. | 
**VirtualPrivateCloudId** | Pointer to **string** | The identifier of the associated VPC. | [optional] 

## Methods

### NewHypervisorSecurityGroupResourceResponseModel

`func NewHypervisorSecurityGroupResourceResponseModel(resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourceResponseModelAllOfResourcePool, isSymLink bool, egressRules []HypervisorSecurityGroupRuleResponseModel, ingressRules []HypervisorSecurityGroupRuleResponseModel, ) *HypervisorSecurityGroupResourceResponseModel`

NewHypervisorSecurityGroupResourceResponseModel instantiates a new HypervisorSecurityGroupResourceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorSecurityGroupResourceResponseModelWithDefaults

`func NewHypervisorSecurityGroupResourceResponseModelWithDefaults() *HypervisorSecurityGroupResourceResponseModel`

NewHypervisorSecurityGroupResourceResponseModelWithDefaults instantiates a new HypervisorSecurityGroupResourceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorSecurityGroupResourceResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorSecurityGroupResourceResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorSecurityGroupResourceResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorSecurityGroupResourceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorSecurityGroupResourceResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorSecurityGroupResourceResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorSecurityGroupResourceResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorSecurityGroupResourceResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorSecurityGroupResourceResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetRelativePath

`func (o *HypervisorSecurityGroupResourceResponseModel) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HypervisorSecurityGroupResourceResponseModel) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *HypervisorSecurityGroupResourceResponseModel) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetFullRelativePath

`func (o *HypervisorSecurityGroupResourceResponseModel) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorSecurityGroupResourceResponseModel) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.

### HasFullRelativePath

`func (o *HypervisorSecurityGroupResourceResponseModel) HasFullRelativePath() bool`

HasFullRelativePath returns a boolean if a field has been set.

### GetResourceType

`func (o *HypervisorSecurityGroupResourceResponseModel) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorSecurityGroupResourceResponseModel) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetObjectTypeName

`func (o *HypervisorSecurityGroupResourceResponseModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *HypervisorSecurityGroupResourceResponseModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *HypervisorSecurityGroupResourceResponseModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetResourceContainer

`func (o *HypervisorSecurityGroupResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModelAllOfResourceContainer`

GetResourceContainer returns the ResourceContainer field if non-nil, zero value otherwise.

### GetResourceContainerOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModelAllOfResourceContainer, bool)`

GetResourceContainerOk returns a tuple with the ResourceContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceContainer

`func (o *HypervisorSecurityGroupResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModelAllOfResourceContainer)`

SetResourceContainer sets ResourceContainer field to given value.

### HasResourceContainer

`func (o *HypervisorSecurityGroupResourceResponseModel) HasResourceContainer() bool`

HasResourceContainer returns a boolean if a field has been set.

### GetResourcePoolXDPath

`func (o *HypervisorSecurityGroupResourceResponseModel) GetResourcePoolXDPath() string`

GetResourcePoolXDPath returns the ResourcePoolXDPath field if non-nil, zero value otherwise.

### GetResourcePoolXDPathOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool)`

GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolXDPath

`func (o *HypervisorSecurityGroupResourceResponseModel) SetResourcePoolXDPath(v string)`

SetResourcePoolXDPath sets ResourcePoolXDPath field to given value.

### HasResourcePoolXDPath

`func (o *HypervisorSecurityGroupResourceResponseModel) HasResourcePoolXDPath() bool`

HasResourcePoolXDPath returns a boolean if a field has been set.

### GetFullName

`func (o *HypervisorSecurityGroupResourceResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *HypervisorSecurityGroupResourceResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsContainer

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *HypervisorSecurityGroupResourceResponseModel) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.


### GetChildren

`func (o *HypervisorSecurityGroupResourceResponseModel) GetChildren() []HypervisorResourceResponseModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetChildrenOk() (*[]HypervisorResourceResponseModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HypervisorSecurityGroupResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HypervisorSecurityGroupResourceResponseModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetIsMachine

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIsMachine() bool`

GetIsMachine returns the IsMachine field if non-nil, zero value otherwise.

### GetIsMachineOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIsMachineOk() (*bool, bool)`

GetIsMachineOk returns a tuple with the IsMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachine

`func (o *HypervisorSecurityGroupResourceResponseModel) SetIsMachine(v bool)`

SetIsMachine sets IsMachine field to given value.


### GetIsSnapshotable

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIsSnapshotable() bool`

GetIsSnapshotable returns the IsSnapshotable field if non-nil, zero value otherwise.

### GetIsSnapshotableOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIsSnapshotableOk() (*bool, bool)`

GetIsSnapshotableOk returns a tuple with the IsSnapshotable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotable

`func (o *HypervisorSecurityGroupResourceResponseModel) SetIsSnapshotable(v bool)`

SetIsSnapshotable sets IsSnapshotable field to given value.


### GetAllResourcesRelativePath

`func (o *HypervisorSecurityGroupResourceResponseModel) GetAllResourcesRelativePath() string`

GetAllResourcesRelativePath returns the AllResourcesRelativePath field if non-nil, zero value otherwise.

### GetAllResourcesRelativePathOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool)`

GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllResourcesRelativePath

`func (o *HypervisorSecurityGroupResourceResponseModel) SetAllResourcesRelativePath(v string)`

SetAllResourcesRelativePath sets AllResourcesRelativePath field to given value.


### GetResourcePool

`func (o *HypervisorSecurityGroupResourceResponseModel) GetResourcePool() HypervisorResourceResponseModelAllOfResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetResourcePoolOk() (*HypervisorResourceResponseModelAllOfResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *HypervisorSecurityGroupResourceResponseModel) SetResourcePool(v HypervisorResourceResponseModelAllOfResourcePool)`

SetResourcePool sets ResourcePool field to given value.


### GetIsSymLink

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIsSymLink() bool`

GetIsSymLink returns the IsSymLink field if non-nil, zero value otherwise.

### GetIsSymLinkOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIsSymLinkOk() (*bool, bool)`

GetIsSymLinkOk returns a tuple with the IsSymLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSymLink

`func (o *HypervisorSecurityGroupResourceResponseModel) SetIsSymLink(v bool)`

SetIsSymLink sets IsSymLink field to given value.


### GetAdditionalData

`func (o *HypervisorSecurityGroupResourceResponseModel) GetAdditionalData() []NameValueStringPairModel`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetAdditionalDataOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *HypervisorSecurityGroupResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *HypervisorSecurityGroupResourceResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetDescription

`func (o *HypervisorSecurityGroupResourceResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HypervisorSecurityGroupResourceResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HypervisorSecurityGroupResourceResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEgressRules

`func (o *HypervisorSecurityGroupResourceResponseModel) GetEgressRules() []HypervisorSecurityGroupRuleResponseModel`

GetEgressRules returns the EgressRules field if non-nil, zero value otherwise.

### GetEgressRulesOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetEgressRulesOk() (*[]HypervisorSecurityGroupRuleResponseModel, bool)`

GetEgressRulesOk returns a tuple with the EgressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRules

`func (o *HypervisorSecurityGroupResourceResponseModel) SetEgressRules(v []HypervisorSecurityGroupRuleResponseModel)`

SetEgressRules sets EgressRules field to given value.


### GetIngressRules

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIngressRules() []HypervisorSecurityGroupRuleResponseModel`

GetIngressRules returns the IngressRules field if non-nil, zero value otherwise.

### GetIngressRulesOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetIngressRulesOk() (*[]HypervisorSecurityGroupRuleResponseModel, bool)`

GetIngressRulesOk returns a tuple with the IngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRules

`func (o *HypervisorSecurityGroupResourceResponseModel) SetIngressRules(v []HypervisorSecurityGroupRuleResponseModel)`

SetIngressRules sets IngressRules field to given value.


### GetVirtualPrivateCloudId

`func (o *HypervisorSecurityGroupResourceResponseModel) GetVirtualPrivateCloudId() string`

GetVirtualPrivateCloudId returns the VirtualPrivateCloudId field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudIdOk

`func (o *HypervisorSecurityGroupResourceResponseModel) GetVirtualPrivateCloudIdOk() (*string, bool)`

GetVirtualPrivateCloudIdOk returns a tuple with the VirtualPrivateCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloudId

`func (o *HypervisorSecurityGroupResourceResponseModel) SetVirtualPrivateCloudId(v string)`

SetVirtualPrivateCloudId sets VirtualPrivateCloudId field to given value.

### HasVirtualPrivateCloudId

`func (o *HypervisorSecurityGroupResourceResponseModel) HasVirtualPrivateCloudId() bool`

HasVirtualPrivateCloudId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


