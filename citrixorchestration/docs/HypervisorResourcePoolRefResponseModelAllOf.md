# HypervisorResourcePoolRefResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullRelativePath** | **string** | Full path to the resources within the resource pool, including the hypervisor, relative to the root of the API. Example: &#x60;Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources&#x60; | 
**Hypervisor** | [**HypervisorResourcePoolRefResponseModelAllOfHypervisor**](HypervisorResourcePoolRefResponseModelAllOfHypervisor.md) |  | 

## Methods

### NewHypervisorResourcePoolRefResponseModelAllOf

`func NewHypervisorResourcePoolRefResponseModelAllOf(fullRelativePath string, hypervisor HypervisorResourcePoolRefResponseModelAllOfHypervisor, ) *HypervisorResourcePoolRefResponseModelAllOf`

NewHypervisorResourcePoolRefResponseModelAllOf instantiates a new HypervisorResourcePoolRefResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolRefResponseModelAllOfWithDefaults

`func NewHypervisorResourcePoolRefResponseModelAllOfWithDefaults() *HypervisorResourcePoolRefResponseModelAllOf`

NewHypervisorResourcePoolRefResponseModelAllOfWithDefaults instantiates a new HypervisorResourcePoolRefResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullRelativePath

`func (o *HypervisorResourcePoolRefResponseModelAllOf) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *HypervisorResourcePoolRefResponseModelAllOf) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *HypervisorResourcePoolRefResponseModelAllOf) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.


### GetHypervisor

`func (o *HypervisorResourcePoolRefResponseModelAllOf) GetHypervisor() HypervisorResourcePoolRefResponseModelAllOfHypervisor`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *HypervisorResourcePoolRefResponseModelAllOf) GetHypervisorOk() (*HypervisorResourcePoolRefResponseModelAllOfHypervisor, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *HypervisorResourcePoolRefResponseModelAllOf) SetHypervisor(v HypervisorResourcePoolRefResponseModelAllOfHypervisor)`

SetHypervisor sets Hypervisor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


