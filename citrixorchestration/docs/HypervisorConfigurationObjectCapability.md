# HypervisorConfigurationObjectCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **NullableString** | Type of hypervisor resource. | [optional] 
**TypeCapabilities** | Pointer to **[]string** | List of capabilities. | [optional] 

## Methods

### NewHypervisorConfigurationObjectCapability

`func NewHypervisorConfigurationObjectCapability() *HypervisorConfigurationObjectCapability`

NewHypervisorConfigurationObjectCapability instantiates a new HypervisorConfigurationObjectCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorConfigurationObjectCapabilityWithDefaults

`func NewHypervisorConfigurationObjectCapabilityWithDefaults() *HypervisorConfigurationObjectCapability`

NewHypervisorConfigurationObjectCapabilityWithDefaults instantiates a new HypervisorConfigurationObjectCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *HypervisorConfigurationObjectCapability) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorConfigurationObjectCapability) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorConfigurationObjectCapability) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *HypervisorConfigurationObjectCapability) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *HypervisorConfigurationObjectCapability) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *HypervisorConfigurationObjectCapability) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetTypeCapabilities

`func (o *HypervisorConfigurationObjectCapability) GetTypeCapabilities() []string`

GetTypeCapabilities returns the TypeCapabilities field if non-nil, zero value otherwise.

### GetTypeCapabilitiesOk

`func (o *HypervisorConfigurationObjectCapability) GetTypeCapabilitiesOk() (*[]string, bool)`

GetTypeCapabilitiesOk returns a tuple with the TypeCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCapabilities

`func (o *HypervisorConfigurationObjectCapability) SetTypeCapabilities(v []string)`

SetTypeCapabilities sets TypeCapabilities field to given value.

### HasTypeCapabilities

`func (o *HypervisorConfigurationObjectCapability) HasTypeCapabilities() bool`

HasTypeCapabilities returns a boolean if a field has been set.

### SetTypeCapabilitiesNil

`func (o *HypervisorConfigurationObjectCapability) SetTypeCapabilitiesNil(b bool)`

 SetTypeCapabilitiesNil sets the value for TypeCapabilities to be an explicit nil

### UnsetTypeCapabilities
`func (o *HypervisorConfigurationObjectCapability) UnsetTypeCapabilities()`

UnsetTypeCapabilities ensures that no value is present for TypeCapabilities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


