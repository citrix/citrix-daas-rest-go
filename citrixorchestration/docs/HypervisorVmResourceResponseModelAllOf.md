# HypervisorVmResourceResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VMId** | **string** | Id of the VM, as defined by the hypervisor. | 
**MacAddress** | Pointer to **string** | MAC address of the VM. | [optional] 

## Methods

### NewHypervisorVmResourceResponseModelAllOf

`func NewHypervisorVmResourceResponseModelAllOf(vMId string, ) *HypervisorVmResourceResponseModelAllOf`

NewHypervisorVmResourceResponseModelAllOf instantiates a new HypervisorVmResourceResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorVmResourceResponseModelAllOfWithDefaults

`func NewHypervisorVmResourceResponseModelAllOfWithDefaults() *HypervisorVmResourceResponseModelAllOf`

NewHypervisorVmResourceResponseModelAllOfWithDefaults instantiates a new HypervisorVmResourceResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVMId

`func (o *HypervisorVmResourceResponseModelAllOf) GetVMId() string`

GetVMId returns the VMId field if non-nil, zero value otherwise.

### GetVMIdOk

`func (o *HypervisorVmResourceResponseModelAllOf) GetVMIdOk() (*string, bool)`

GetVMIdOk returns a tuple with the VMId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMId

`func (o *HypervisorVmResourceResponseModelAllOf) SetVMId(v string)`

SetVMId sets VMId field to given value.


### GetMacAddress

`func (o *HypervisorVmResourceResponseModelAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *HypervisorVmResourceResponseModelAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *HypervisorVmResourceResponseModelAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *HypervisorVmResourceResponseModelAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


