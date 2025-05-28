# IPAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to [**AddressFamily**](AddressFamily.md) |  | [optional] 
**ScopeId** | Pointer to **int64** |  | [optional] 
**IsIPv6Multicast** | Pointer to **bool** |  | [optional] [readonly] 
**IsIPv6LinkLocal** | Pointer to **bool** |  | [optional] [readonly] 
**IsIPv6SiteLocal** | Pointer to **bool** |  | [optional] [readonly] 
**IsIPv6Teredo** | Pointer to **bool** |  | [optional] [readonly] 
**IsIPv6UniqueLocal** | Pointer to **bool** |  | [optional] [readonly] 
**IsIPv4MappedToIPv6** | Pointer to **bool** |  | [optional] [readonly] 
**Address** | Pointer to **int64** |  | [optional] 

## Methods

### NewIPAddress

`func NewIPAddress() *IPAddress`

NewIPAddress instantiates a new IPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressWithDefaults

`func NewIPAddressWithDefaults() *IPAddress`

NewIPAddressWithDefaults instantiates a new IPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *IPAddress) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *IPAddress) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *IPAddress) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *IPAddress) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetScopeId

`func (o *IPAddress) GetScopeId() int64`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *IPAddress) GetScopeIdOk() (*int64, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *IPAddress) SetScopeId(v int64)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *IPAddress) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetIsIPv6Multicast

`func (o *IPAddress) GetIsIPv6Multicast() bool`

GetIsIPv6Multicast returns the IsIPv6Multicast field if non-nil, zero value otherwise.

### GetIsIPv6MulticastOk

`func (o *IPAddress) GetIsIPv6MulticastOk() (*bool, bool)`

GetIsIPv6MulticastOk returns a tuple with the IsIPv6Multicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIPv6Multicast

`func (o *IPAddress) SetIsIPv6Multicast(v bool)`

SetIsIPv6Multicast sets IsIPv6Multicast field to given value.

### HasIsIPv6Multicast

`func (o *IPAddress) HasIsIPv6Multicast() bool`

HasIsIPv6Multicast returns a boolean if a field has been set.

### GetIsIPv6LinkLocal

`func (o *IPAddress) GetIsIPv6LinkLocal() bool`

GetIsIPv6LinkLocal returns the IsIPv6LinkLocal field if non-nil, zero value otherwise.

### GetIsIPv6LinkLocalOk

`func (o *IPAddress) GetIsIPv6LinkLocalOk() (*bool, bool)`

GetIsIPv6LinkLocalOk returns a tuple with the IsIPv6LinkLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIPv6LinkLocal

`func (o *IPAddress) SetIsIPv6LinkLocal(v bool)`

SetIsIPv6LinkLocal sets IsIPv6LinkLocal field to given value.

### HasIsIPv6LinkLocal

`func (o *IPAddress) HasIsIPv6LinkLocal() bool`

HasIsIPv6LinkLocal returns a boolean if a field has been set.

### GetIsIPv6SiteLocal

`func (o *IPAddress) GetIsIPv6SiteLocal() bool`

GetIsIPv6SiteLocal returns the IsIPv6SiteLocal field if non-nil, zero value otherwise.

### GetIsIPv6SiteLocalOk

`func (o *IPAddress) GetIsIPv6SiteLocalOk() (*bool, bool)`

GetIsIPv6SiteLocalOk returns a tuple with the IsIPv6SiteLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIPv6SiteLocal

`func (o *IPAddress) SetIsIPv6SiteLocal(v bool)`

SetIsIPv6SiteLocal sets IsIPv6SiteLocal field to given value.

### HasIsIPv6SiteLocal

`func (o *IPAddress) HasIsIPv6SiteLocal() bool`

HasIsIPv6SiteLocal returns a boolean if a field has been set.

### GetIsIPv6Teredo

`func (o *IPAddress) GetIsIPv6Teredo() bool`

GetIsIPv6Teredo returns the IsIPv6Teredo field if non-nil, zero value otherwise.

### GetIsIPv6TeredoOk

`func (o *IPAddress) GetIsIPv6TeredoOk() (*bool, bool)`

GetIsIPv6TeredoOk returns a tuple with the IsIPv6Teredo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIPv6Teredo

`func (o *IPAddress) SetIsIPv6Teredo(v bool)`

SetIsIPv6Teredo sets IsIPv6Teredo field to given value.

### HasIsIPv6Teredo

`func (o *IPAddress) HasIsIPv6Teredo() bool`

HasIsIPv6Teredo returns a boolean if a field has been set.

### GetIsIPv6UniqueLocal

`func (o *IPAddress) GetIsIPv6UniqueLocal() bool`

GetIsIPv6UniqueLocal returns the IsIPv6UniqueLocal field if non-nil, zero value otherwise.

### GetIsIPv6UniqueLocalOk

`func (o *IPAddress) GetIsIPv6UniqueLocalOk() (*bool, bool)`

GetIsIPv6UniqueLocalOk returns a tuple with the IsIPv6UniqueLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIPv6UniqueLocal

`func (o *IPAddress) SetIsIPv6UniqueLocal(v bool)`

SetIsIPv6UniqueLocal sets IsIPv6UniqueLocal field to given value.

### HasIsIPv6UniqueLocal

`func (o *IPAddress) HasIsIPv6UniqueLocal() bool`

HasIsIPv6UniqueLocal returns a boolean if a field has been set.

### GetIsIPv4MappedToIPv6

`func (o *IPAddress) GetIsIPv4MappedToIPv6() bool`

GetIsIPv4MappedToIPv6 returns the IsIPv4MappedToIPv6 field if non-nil, zero value otherwise.

### GetIsIPv4MappedToIPv6Ok

`func (o *IPAddress) GetIsIPv4MappedToIPv6Ok() (*bool, bool)`

GetIsIPv4MappedToIPv6Ok returns a tuple with the IsIPv4MappedToIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIPv4MappedToIPv6

`func (o *IPAddress) SetIsIPv4MappedToIPv6(v bool)`

SetIsIPv4MappedToIPv6 sets IsIPv4MappedToIPv6 field to given value.

### HasIsIPv4MappedToIPv6

`func (o *IPAddress) HasIsIPv4MappedToIPv6() bool`

HasIsIPv4MappedToIPv6 returns a boolean if a field has been set.

### GetAddress

`func (o *IPAddress) GetAddress() int64`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAddress) GetAddressOk() (*int64, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAddress) SetAddress(v int64)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IPAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


