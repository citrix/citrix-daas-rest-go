# RouteOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the route | 
**Name** | **string** | The name of the route | 
**AddressPrefix** | **string** | The destination CIDR to which the route applies. | 
**NextHopType** | **string** | The type of hop the packet should be sent to. | 
**NextHopIpAddress** | Pointer to **string** | The IP address packets should be forwarded to | [optional] 
**Enabled** | Pointer to **bool** | Specifies if the route is enabled | [optional] 

## Methods

### NewRouteOverview

`func NewRouteOverview(id string, name string, addressPrefix string, nextHopType string, ) *RouteOverview`

NewRouteOverview instantiates a new RouteOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteOverviewWithDefaults

`func NewRouteOverviewWithDefaults() *RouteOverview`

NewRouteOverviewWithDefaults instantiates a new RouteOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteOverview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteOverview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteOverview) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RouteOverview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteOverview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteOverview) SetName(v string)`

SetName sets Name field to given value.


### GetAddressPrefix

`func (o *RouteOverview) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *RouteOverview) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *RouteOverview) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.


### GetNextHopType

`func (o *RouteOverview) GetNextHopType() string`

GetNextHopType returns the NextHopType field if non-nil, zero value otherwise.

### GetNextHopTypeOk

`func (o *RouteOverview) GetNextHopTypeOk() (*string, bool)`

GetNextHopTypeOk returns a tuple with the NextHopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopType

`func (o *RouteOverview) SetNextHopType(v string)`

SetNextHopType sets NextHopType field to given value.


### GetNextHopIpAddress

`func (o *RouteOverview) GetNextHopIpAddress() string`

GetNextHopIpAddress returns the NextHopIpAddress field if non-nil, zero value otherwise.

### GetNextHopIpAddressOk

`func (o *RouteOverview) GetNextHopIpAddressOk() (*string, bool)`

GetNextHopIpAddressOk returns a tuple with the NextHopIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIpAddress

`func (o *RouteOverview) SetNextHopIpAddress(v string)`

SetNextHopIpAddress sets NextHopIpAddress field to given value.

### HasNextHopIpAddress

`func (o *RouteOverview) HasNextHopIpAddress() bool`

HasNextHopIpAddress returns a boolean if a field has been set.

### GetEnabled

`func (o *RouteOverview) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RouteOverview) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RouteOverview) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RouteOverview) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


