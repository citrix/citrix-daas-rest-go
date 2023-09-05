# HypervisorResourcePoolRefResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**XDPath** | Pointer to **string** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**FullRelativePath** | **string** | Full path to the resources within the resource pool, including the hypervisor, relative to the root of the API. Example: &#x60;Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources&#x60; | 
**Hypervisor** | [**HypervisorResourcePoolRefResponseModelAllOfHypervisor**](HypervisorResourcePoolRefResponseModelAllOfHypervisor.md) |  | 

## Methods

### NewHypervisorResourcePoolRefResponseModel

`func NewHypervisorResourcePoolRefResponseModel(fullRelativePath string, hypervisor HypervisorResourcePoolRefResponseModelAllOfHypervisor, ) *HypervisorResourcePoolRefResponseModel`

NewHypervisorResourcePoolRefResponseModel instantiates a new HypervisorResourcePoolRefResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolRefResponseModelWithDefaults

`func NewHypervisorResourcePoolRefResponseModelWithDefaults() *HypervisorResourcePoolRefResponseModel`

NewHypervisorResourcePoolRefResponseModelWithDefaults instantiates a new HypervisorResourcePoolRefResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorResourcePoolRefResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorResourcePoolRefResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorResourcePoolRefResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorResourcePoolRefResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorResourcePoolRefResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorResourcePoolRefResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorResourcePoolRefResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorResourcePoolRefResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorResourcePoolRefResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorResourcePoolRefResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorResourcePoolRefResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorResourcePoolRefResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetFullRelativePath

`func (o *HypervisorResourcePoolRefResponseModel) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorResourcePoolRefResponseModel) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorResourcePoolRefResponseModel) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.


### GetHypervisor

`func (o *HypervisorResourcePoolRefResponseModel) GetHypervisor() HypervisorResourcePoolRefResponseModelAllOfHypervisor`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *HypervisorResourcePoolRefResponseModel) GetHypervisorOk() (*HypervisorResourcePoolRefResponseModelAllOfHypervisor, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *HypervisorResourcePoolRefResponseModel) SetHypervisor(v HypervisorResourcePoolRefResponseModelAllOfHypervisor)`

SetHypervisor sets Hypervisor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


