# HypervisorTraditionalConnectionDetailRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | Hypervisor user name.  Required. | 
**Password** | **string** | Hypervisor password.  Required. Must be specified in the format indicated by PasswordFormat. | 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**Addresses** | **[]string** | Hypervisor address(es).  At least one is required. | 
**PluginId** | Pointer to **string** | Custom hypervisor plugin ID.  Required if the ConnectionType is Custom.  Otherwise must not be specified. | [optional] 
**SslThumbprints** | Pointer to **[]string** | SSL certificate thumbprints to consider acceptable for this connection.  Optional.  If not specified, and the hypervisor uses SSL for its connection, the SSL certificate&#39;s root certification authority and any intermediate certificates must be trusted. | [optional] 

## Methods

### NewHypervisorTraditionalConnectionDetailRequestModelAllOf

`func NewHypervisorTraditionalConnectionDetailRequestModelAllOf(userName string, password string, addresses []string, ) *HypervisorTraditionalConnectionDetailRequestModelAllOf`

NewHypervisorTraditionalConnectionDetailRequestModelAllOf instantiates a new HypervisorTraditionalConnectionDetailRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorTraditionalConnectionDetailRequestModelAllOfWithDefaults

`func NewHypervisorTraditionalConnectionDetailRequestModelAllOfWithDefaults() *HypervisorTraditionalConnectionDetailRequestModelAllOf`

NewHypervisorTraditionalConnectionDetailRequestModelAllOfWithDefaults instantiates a new HypervisorTraditionalConnectionDetailRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPassword

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPasswordFormat

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.

### GetAddresses

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetPluginId

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetSslThumbprints

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetSslThumbprints() []string`

GetSslThumbprints returns the SslThumbprints field if non-nil, zero value otherwise.

### GetSslThumbprintsOk

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) GetSslThumbprintsOk() (*[]string, bool)`

GetSslThumbprintsOk returns a tuple with the SslThumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprints

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) SetSslThumbprints(v []string)`

SetSslThumbprints sets SslThumbprints field to given value.

### HasSslThumbprints

`func (o *HypervisorTraditionalConnectionDetailRequestModelAllOf) HasSslThumbprints() bool`

HasSslThumbprints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


