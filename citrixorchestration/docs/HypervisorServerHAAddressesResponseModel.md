# HypervisorServerHAAddressesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**XDPath** | Pointer to **string** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**HAAddresses** | Pointer to **[]string** | The HA addresses array. | [optional] 

## Methods

### NewHypervisorServerHAAddressesResponseModel

`func NewHypervisorServerHAAddressesResponseModel() *HypervisorServerHAAddressesResponseModel`

NewHypervisorServerHAAddressesResponseModel instantiates a new HypervisorServerHAAddressesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorServerHAAddressesResponseModelWithDefaults

`func NewHypervisorServerHAAddressesResponseModelWithDefaults() *HypervisorServerHAAddressesResponseModel`

NewHypervisorServerHAAddressesResponseModelWithDefaults instantiates a new HypervisorServerHAAddressesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorServerHAAddressesResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorServerHAAddressesResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorServerHAAddressesResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorServerHAAddressesResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorServerHAAddressesResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorServerHAAddressesResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorServerHAAddressesResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorServerHAAddressesResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorServerHAAddressesResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorServerHAAddressesResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorServerHAAddressesResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorServerHAAddressesResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetHAAddresses

`func (o *HypervisorServerHAAddressesResponseModel) GetHAAddresses() []string`

GetHAAddresses returns the HAAddresses field if non-nil, zero value otherwise.

### GetHAAddressesOk

`func (o *HypervisorServerHAAddressesResponseModel) GetHAAddressesOk() (*[]string, bool)`

GetHAAddressesOk returns a tuple with the HAAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHAAddresses

`func (o *HypervisorServerHAAddressesResponseModel) SetHAAddresses(v []string)`

SetHAAddresses sets HAAddresses field to given value.

### HasHAAddresses

`func (o *HypervisorServerHAAddressesResponseModel) HasHAAddresses() bool`

HasHAAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


