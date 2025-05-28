# AzureRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressPrefix** | **string** | The destination CIDR to which the route applies. | 
**NextHopType** | **string** | The type of hop the packet should be sent to. | 
**NextHopIpAddress** | Pointer to **string** | The IP address packets should be forwarded to | [optional] 
**Name** | **string** | Name of the route | 
**Enabled** | Pointer to **bool** | Specifying if the route should be applied | [optional] 

## Methods

### NewAzureRoute

`func NewAzureRoute(addressPrefix string, nextHopType string, name string, ) *AzureRoute`

NewAzureRoute instantiates a new AzureRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRouteWithDefaults

`func NewAzureRouteWithDefaults() *AzureRoute`

NewAzureRouteWithDefaults instantiates a new AzureRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressPrefix

`func (o *AzureRoute) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *AzureRoute) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *AzureRoute) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.


### GetNextHopType

`func (o *AzureRoute) GetNextHopType() string`

GetNextHopType returns the NextHopType field if non-nil, zero value otherwise.

### GetNextHopTypeOk

`func (o *AzureRoute) GetNextHopTypeOk() (*string, bool)`

GetNextHopTypeOk returns a tuple with the NextHopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopType

`func (o *AzureRoute) SetNextHopType(v string)`

SetNextHopType sets NextHopType field to given value.


### GetNextHopIpAddress

`func (o *AzureRoute) GetNextHopIpAddress() string`

GetNextHopIpAddress returns the NextHopIpAddress field if non-nil, zero value otherwise.

### GetNextHopIpAddressOk

`func (o *AzureRoute) GetNextHopIpAddressOk() (*string, bool)`

GetNextHopIpAddressOk returns a tuple with the NextHopIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIpAddress

`func (o *AzureRoute) SetNextHopIpAddress(v string)`

SetNextHopIpAddress sets NextHopIpAddress field to given value.

### HasNextHopIpAddress

`func (o *AzureRoute) HasNextHopIpAddress() bool`

HasNextHopIpAddress returns a boolean if a field has been set.

### GetName

`func (o *AzureRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureRoute) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *AzureRoute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AzureRoute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AzureRoute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AzureRoute) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


