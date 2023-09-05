# HypervisorResourcePoolGcpDetailResponseModelAllOfProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**XDPath** | Pointer to **string** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**RelativePath** | Pointer to **string** | Path to the object, relative to the resource pool from which it was queried. &#x60;{{vm name}}.vm/{{snapshot name}}.snapshot&#x60; | [optional] 
**FullRelativePath** | Pointer to **string** | Full path to the resource, including the hypervisor, relative to the root of the API. Example: &#x60;Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources/{{RelativePath}}&#x60; | [optional] 
**ResourceType** | Pointer to **string** | Type of resource. | [optional] 
**ObjectTypeName** | Pointer to **string** | The type name of the hypervisor resource object. | [optional] 
**ResourceContainer** | Pointer to [**HypervisorResourceRefResponseModelAllOfResourceContainer**](HypervisorResourceRefResponseModelAllOfResourceContainer.md) |  | [optional] 
**ResourcePoolXDPath** | Pointer to **string** | Citrix Apps and Desktops path to the resource on the ResourcePool.  An example value is: &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; This value  | [optional] 

## Methods

### NewHypervisorResourcePoolGcpDetailResponseModelAllOfProject

`func NewHypervisorResourcePoolGcpDetailResponseModelAllOfProject() *HypervisorResourcePoolGcpDetailResponseModelAllOfProject`

NewHypervisorResourcePoolGcpDetailResponseModelAllOfProject instantiates a new HypervisorResourcePoolGcpDetailResponseModelAllOfProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolGcpDetailResponseModelAllOfProjectWithDefaults

`func NewHypervisorResourcePoolGcpDetailResponseModelAllOfProjectWithDefaults() *HypervisorResourcePoolGcpDetailResponseModelAllOfProject`

NewHypervisorResourcePoolGcpDetailResponseModelAllOfProjectWithDefaults instantiates a new HypervisorResourcePoolGcpDetailResponseModelAllOfProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetRelativePath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetFullRelativePath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.

### HasFullRelativePath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) HasFullRelativePath() bool`

HasFullRelativePath returns a boolean if a field has been set.

### GetResourceType

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetResourceContainer

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetResourceContainer() HypervisorResourceRefResponseModelAllOfResourceContainer`

GetResourceContainer returns the ResourceContainer field if non-nil, zero value otherwise.

### GetResourceContainerOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetResourceContainerOk() (*HypervisorResourceRefResponseModelAllOfResourceContainer, bool)`

GetResourceContainerOk returns a tuple with the ResourceContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceContainer

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) SetResourceContainer(v HypervisorResourceRefResponseModelAllOfResourceContainer)`

SetResourceContainer sets ResourceContainer field to given value.

### HasResourceContainer

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) HasResourceContainer() bool`

HasResourceContainer returns a boolean if a field has been set.

### GetResourcePoolXDPath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetResourcePoolXDPath() string`

GetResourcePoolXDPath returns the ResourcePoolXDPath field if non-nil, zero value otherwise.

### GetResourcePoolXDPathOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) GetResourcePoolXDPathOk() (*string, bool)`

GetResourcePoolXDPathOk returns a tuple with the ResourcePoolXDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolXDPath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) SetResourcePoolXDPath(v string)`

SetResourcePoolXDPath sets ResourcePoolXDPath field to given value.

### HasResourcePoolXDPath

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOfProject) HasResourcePoolXDPath() bool`

HasResourcePoolXDPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


