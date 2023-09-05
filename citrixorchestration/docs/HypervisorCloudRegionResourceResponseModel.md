# HypervisorCloudRegionResourceResponseModel

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
**Endpoint** | Pointer to **string** | The API address with the region. | [optional] 

## Methods

### NewHypervisorCloudRegionResourceResponseModel

`func NewHypervisorCloudRegionResourceResponseModel(resourceType string, fullName string, isContainer bool, isMachine bool, isSnapshotable bool, allResourcesRelativePath string, resourcePool HypervisorResourceResponseModelAllOfResourcePool, isSymLink bool, ) *HypervisorCloudRegionResourceResponseModel`

NewHypervisorCloudRegionResourceResponseModel instantiates a new HypervisorCloudRegionResourceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorCloudRegionResourceResponseModelWithDefaults

`func NewHypervisorCloudRegionResourceResponseModelWithDefaults() *HypervisorCloudRegionResourceResponseModel`

NewHypervisorCloudRegionResourceResponseModelWithDefaults instantiates a new HypervisorCloudRegionResourceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorCloudRegionResourceResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorCloudRegionResourceResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorCloudRegionResourceResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorCloudRegionResourceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorCloudRegionResourceResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorCloudRegionResourceResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorCloudRegionResourceResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorCloudRegionResourceResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorCloudRegionResourceResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetRelativePath

`func (o *HypervisorCloudRegionResourceResponseModel) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HypervisorCloudRegionResourceResponseModel) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *HypervisorCloudRegionResourceResponseModel) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetFullRelativePath

`func (o *HypervisorCloudRegionResourceResponseModel) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorCloudRegionResourceResponseModel) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.

### HasFullRelativePath

`func (o *HypervisorCloudRegionResourceResponseModel) HasFullRelativePath() bool`

HasFullRelativePath returns a boolean if a field has been set.

### GetResourceType

`func (o *HypervisorCloudRegionResourceResponseModel) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorCloudRegionResourceResponseModel) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetObjectTypeName

`func (o *HypervisorCloudRegionResourceResponseModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *HypervisorCloudRegionResourceResponseModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *HypervisorCloudRegionResourceResponseModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetResourceContainer

`func (o *HypervisorCloudRegionResourceResponseModel) GetResourceContainer() HypervisorResourceRefResponseModelAllOfResourceContainer`

GetResourceContainer returns the ResourceContainer field if non-nil, zero value otherwise.

### GetResourceContainerOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetResourceContainerOk() (*HypervisorResourceRefResponseModelAllOfResourceContainer, bool)`

GetResourceContainerOk returns a tuple with the ResourceContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceContainer

`func (o *HypervisorCloudRegionResourceResponseModel) SetResourceContainer(v HypervisorResourceRefResponseModelAllOfResourceContainer)`

SetResourceContainer sets ResourceContainer field to given value.

### HasResourceContainer

`func (o *HypervisorCloudRegionResourceResponseModel) HasResourceContainer() bool`

HasResourceContainer returns a boolean if a field has been set.

### GetResourcePoolXDPath

`func (o *HypervisorCloudRegionResourceResponseModel) GetResourcePoolXDPath() string`

GetResourcePoolXDPath returns the ResourcePoolXDPath field if non-nil, zero value otherwise.

### GetResourcePoolXDPathOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetResourcePoolXDPathOk() (*string, bool)`

GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolXDPath

`func (o *HypervisorCloudRegionResourceResponseModel) SetResourcePoolXDPath(v string)`

SetResourcePoolXDPath sets ResourcePoolXDPath field to given value.

### HasResourcePoolXDPath

`func (o *HypervisorCloudRegionResourceResponseModel) HasResourcePoolXDPath() bool`

HasResourcePoolXDPath returns a boolean if a field has been set.

### GetFullName

`func (o *HypervisorCloudRegionResourceResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *HypervisorCloudRegionResourceResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsContainer

`func (o *HypervisorCloudRegionResourceResponseModel) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *HypervisorCloudRegionResourceResponseModel) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.


### GetChildren

`func (o *HypervisorCloudRegionResourceResponseModel) GetChildren() []HypervisorResourceResponseModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetChildrenOk() (*[]HypervisorResourceResponseModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HypervisorCloudRegionResourceResponseModel) SetChildren(v []HypervisorResourceResponseModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HypervisorCloudRegionResourceResponseModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetIsMachine

`func (o *HypervisorCloudRegionResourceResponseModel) GetIsMachine() bool`

GetIsMachine returns the IsMachine field if non-nil, zero value otherwise.

### GetIsMachineOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetIsMachineOk() (*bool, bool)`

GetIsMachineOk returns a tuple with the IsMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachine

`func (o *HypervisorCloudRegionResourceResponseModel) SetIsMachine(v bool)`

SetIsMachine sets IsMachine field to given value.


### GetIsSnapshotable

`func (o *HypervisorCloudRegionResourceResponseModel) GetIsSnapshotable() bool`

GetIsSnapshotable returns the IsSnapshotable field if non-nil, zero value otherwise.

### GetIsSnapshotableOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetIsSnapshotableOk() (*bool, bool)`

GetIsSnapshotableOk returns a tuple with the IsSnapshotable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotable

`func (o *HypervisorCloudRegionResourceResponseModel) SetIsSnapshotable(v bool)`

SetIsSnapshotable sets IsSnapshotable field to given value.


### GetAllResourcesRelativePath

`func (o *HypervisorCloudRegionResourceResponseModel) GetAllResourcesRelativePath() string`

GetAllResourcesRelativePath returns the AllResourcesRelativePath field if non-nil, zero value otherwise.

### GetAllResourcesRelativePathOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetAllResourcesRelativePathOk() (*string, bool)`

GetAllResourcesRelativePathOk returns a tuple with the AllResourcesRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllResourcesRelativePath

`func (o *HypervisorCloudRegionResourceResponseModel) SetAllResourcesRelativePath(v string)`

SetAllResourcesRelativePath sets AllResourcesRelativePath field to given value.


### GetResourcePool

`func (o *HypervisorCloudRegionResourceResponseModel) GetResourcePool() HypervisorResourceResponseModelAllOfResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetResourcePoolOk() (*HypervisorResourceResponseModelAllOfResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *HypervisorCloudRegionResourceResponseModel) SetResourcePool(v HypervisorResourceResponseModelAllOfResourcePool)`

SetResourcePool sets ResourcePool field to given value.


### GetIsSymLink

`func (o *HypervisorCloudRegionResourceResponseModel) GetIsSymLink() bool`

GetIsSymLink returns the IsSymLink field if non-nil, zero value otherwise.

### GetIsSymLinkOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetIsSymLinkOk() (*bool, bool)`

GetIsSymLinkOk returns a tuple with the IsSymLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSymLink

`func (o *HypervisorCloudRegionResourceResponseModel) SetIsSymLink(v bool)`

SetIsSymLink sets IsSymLink field to given value.


### GetAdditionalData

`func (o *HypervisorCloudRegionResourceResponseModel) GetAdditionalData() []NameValueStringPairModel`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetAdditionalDataOk() (*[]NameValueStringPairModel, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *HypervisorCloudRegionResourceResponseModel) SetAdditionalData(v []NameValueStringPairModel)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *HypervisorCloudRegionResourceResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetEndpoint

`func (o *HypervisorCloudRegionResourceResponseModel) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HypervisorCloudRegionResourceResponseModel) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HypervisorCloudRegionResourceResponseModel) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HypervisorCloudRegionResourceResponseModel) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


