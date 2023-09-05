# HypervisorsAndResourcePoolsResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayConnectionType** | **string** | Hypervisor connection connection type display name. | 
**IsBroken** | **bool** | Indicates whether the hypervisor connection is broken. | 
**Errors** | **[]string** | Hypervisor connection broken errors. | 
**ErrorState** | **string** | Current hypervisor connection error status. | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of hypervisor connections. | [optional] 
**ResourcePools** | Pointer to [**[]HypervisorBaseResponseModel**](HypervisorBaseResponseModel.md) | Hypervisor resource pools ref response list. | [optional] 

## Methods

### NewHypervisorsAndResourcePoolsResponseModelAllOf

`func NewHypervisorsAndResourcePoolsResponseModelAllOf(displayConnectionType string, isBroken bool, errors []string, errorState string, ) *HypervisorsAndResourcePoolsResponseModelAllOf`

NewHypervisorsAndResourcePoolsResponseModelAllOf instantiates a new HypervisorsAndResourcePoolsResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorsAndResourcePoolsResponseModelAllOfWithDefaults

`func NewHypervisorsAndResourcePoolsResponseModelAllOfWithDefaults() *HypervisorsAndResourcePoolsResponseModelAllOf`

NewHypervisorsAndResourcePoolsResponseModelAllOfWithDefaults instantiates a new HypervisorsAndResourcePoolsResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayConnectionType

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetDisplayConnectionType() string`

GetDisplayConnectionType returns the DisplayConnectionType field if non-nil, zero value otherwise.

### GetDisplayConnectionTypeOk

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetDisplayConnectionTypeOk() (*string, bool)`

GetDisplayConnectionTypeOk returns a tuple with the DisplayConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayConnectionType

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) SetDisplayConnectionType(v string)`

SetDisplayConnectionType sets DisplayConnectionType field to given value.


### GetIsBroken

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.


### GetErrors

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### GetErrorState

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.


### GetMetadata

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResourcePools

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetResourcePools() []HypervisorBaseResponseModel`

GetResourcePools returns the ResourcePools field if non-nil, zero value otherwise.

### GetResourcePoolsOk

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) GetResourcePoolsOk() (*[]HypervisorBaseResponseModel, bool)`

GetResourcePoolsOk returns a tuple with the ResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePools

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) SetResourcePools(v []HypervisorBaseResponseModel)`

SetResourcePools sets ResourcePools field to given value.

### HasResourcePools

`func (o *HypervisorsAndResourcePoolsResponseModelAllOf) HasResourcePools() bool`

HasResourcePools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


