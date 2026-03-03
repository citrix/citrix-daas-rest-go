# ServiceEndpointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceLocationId** | Pointer to **NullableString** | The resource location ID | [optional] 
**SubnetName** | Pointer to **NullableString** | The subnet name | [optional] 
**ServiceEndpoints** | Pointer to [**[]ServiceEndpointInfo**](ServiceEndpointInfo.md) | List of all available service endpoints with their current states | [optional] 

## Methods

### NewServiceEndpointsResponse

`func NewServiceEndpointsResponse() *ServiceEndpointsResponse`

NewServiceEndpointsResponse instantiates a new ServiceEndpointsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceEndpointsResponseWithDefaults

`func NewServiceEndpointsResponseWithDefaults() *ServiceEndpointsResponse`

NewServiceEndpointsResponseWithDefaults instantiates a new ServiceEndpointsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceLocationId

`func (o *ServiceEndpointsResponse) GetResourceLocationId() string`

GetResourceLocationId returns the ResourceLocationId field if non-nil, zero value otherwise.

### GetResourceLocationIdOk

`func (o *ServiceEndpointsResponse) GetResourceLocationIdOk() (*string, bool)`

GetResourceLocationIdOk returns a tuple with the ResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationId

`func (o *ServiceEndpointsResponse) SetResourceLocationId(v string)`

SetResourceLocationId sets ResourceLocationId field to given value.

### HasResourceLocationId

`func (o *ServiceEndpointsResponse) HasResourceLocationId() bool`

HasResourceLocationId returns a boolean if a field has been set.

### SetResourceLocationIdNil

`func (o *ServiceEndpointsResponse) SetResourceLocationIdNil(b bool)`

 SetResourceLocationIdNil sets the value for ResourceLocationId to be an explicit nil

### UnsetResourceLocationId
`func (o *ServiceEndpointsResponse) UnsetResourceLocationId()`

UnsetResourceLocationId ensures that no value is present for ResourceLocationId, not even an explicit nil
### GetSubnetName

`func (o *ServiceEndpointsResponse) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *ServiceEndpointsResponse) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *ServiceEndpointsResponse) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *ServiceEndpointsResponse) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *ServiceEndpointsResponse) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *ServiceEndpointsResponse) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetServiceEndpoints

`func (o *ServiceEndpointsResponse) GetServiceEndpoints() []ServiceEndpointInfo`

GetServiceEndpoints returns the ServiceEndpoints field if non-nil, zero value otherwise.

### GetServiceEndpointsOk

`func (o *ServiceEndpointsResponse) GetServiceEndpointsOk() (*[]ServiceEndpointInfo, bool)`

GetServiceEndpointsOk returns a tuple with the ServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoints

`func (o *ServiceEndpointsResponse) SetServiceEndpoints(v []ServiceEndpointInfo)`

SetServiceEndpoints sets ServiceEndpoints field to given value.

### HasServiceEndpoints

`func (o *ServiceEndpointsResponse) HasServiceEndpoints() bool`

HasServiceEndpoints returns a boolean if a field has been set.

### SetServiceEndpointsNil

`func (o *ServiceEndpointsResponse) SetServiceEndpointsNil(b bool)`

 SetServiceEndpointsNil sets the value for ServiceEndpoints to be an explicit nil

### UnsetServiceEndpoints
`func (o *ServiceEndpointsResponse) UnsetServiceEndpoints()`

UnsetServiceEndpoints ensures that no value is present for ServiceEndpoints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


