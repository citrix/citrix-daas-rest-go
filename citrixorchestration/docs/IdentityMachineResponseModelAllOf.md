# IdentityMachineResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamName** | Pointer to **string** | Fully qualified sAMAccountName of the Machine object (including domain). | [optional] 
**SamAccountName** | Pointer to **string** | Unqualified SAM name (Directory property) of the Machine object.  Does not include the domain. The sAMAccountName is unique within a domain. | [optional] 
**DnsName** | Pointer to **string** | The full qualified domain name associated with the Machine object. | [optional] 
**Sid** | Pointer to **string** | The Security Identifier associated with the Machine object. | [optional] 
**ServicePrincipalNames** | Pointer to **[]string** | The service principal names (SPN) associated with the Machine object. | [optional] 
**IPAddress** | Pointer to **[]string** | The IP addresses associated with the Machine object. | [optional] 
**IPAddressResolveMethod** | Pointer to [**IdentityMachineIPAddressResolveMethod**](IdentityMachineIPAddressResolveMethod.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Machine object is enabled. NOTE that this is opposite of the low-level SDK, which has \&quot;IsDisabled\&quot;. By changing to \&quot;Enabled\&quot; it avoids a confusing double-negative.  It also matches all other SDK objects by using \&quot;Enabled\&quot; rather than \&quot;IsEnabled\&quot;. | [optional] 
**Locked** | Pointer to **bool** | Indicates whether the Machine object is locked. low-level has \&quot;IsLocked\&quot;. | [optional] 
**PropertiesFetched** | **int32** | Properties fetched and populated in the Machine object.  This is a bitfield indicating the fetched properties. | 

## Methods

### NewIdentityMachineResponseModelAllOf

`func NewIdentityMachineResponseModelAllOf(propertiesFetched int32, ) *IdentityMachineResponseModelAllOf`

NewIdentityMachineResponseModelAllOf instantiates a new IdentityMachineResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityMachineResponseModelAllOfWithDefaults

`func NewIdentityMachineResponseModelAllOfWithDefaults() *IdentityMachineResponseModelAllOf`

NewIdentityMachineResponseModelAllOfWithDefaults instantiates a new IdentityMachineResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamName

`func (o *IdentityMachineResponseModelAllOf) GetSamName() string`

GetSamName returns the SamName field if non-nil, zero value otherwise.

### GetSamNameOk

`func (o *IdentityMachineResponseModelAllOf) GetSamNameOk() (*string, bool)`

GetSamNameOk returns a tuple with the SamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamName

`func (o *IdentityMachineResponseModelAllOf) SetSamName(v string)`

SetSamName sets SamName field to given value.

### HasSamName

`func (o *IdentityMachineResponseModelAllOf) HasSamName() bool`

HasSamName returns a boolean if a field has been set.

### GetSamAccountName

`func (o *IdentityMachineResponseModelAllOf) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *IdentityMachineResponseModelAllOf) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *IdentityMachineResponseModelAllOf) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *IdentityMachineResponseModelAllOf) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.

### GetDnsName

`func (o *IdentityMachineResponseModelAllOf) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *IdentityMachineResponseModelAllOf) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *IdentityMachineResponseModelAllOf) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *IdentityMachineResponseModelAllOf) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetSid

`func (o *IdentityMachineResponseModelAllOf) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IdentityMachineResponseModelAllOf) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IdentityMachineResponseModelAllOf) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IdentityMachineResponseModelAllOf) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetServicePrincipalNames

`func (o *IdentityMachineResponseModelAllOf) GetServicePrincipalNames() []string`

GetServicePrincipalNames returns the ServicePrincipalNames field if non-nil, zero value otherwise.

### GetServicePrincipalNamesOk

`func (o *IdentityMachineResponseModelAllOf) GetServicePrincipalNamesOk() (*[]string, bool)`

GetServicePrincipalNamesOk returns a tuple with the ServicePrincipalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalNames

`func (o *IdentityMachineResponseModelAllOf) SetServicePrincipalNames(v []string)`

SetServicePrincipalNames sets ServicePrincipalNames field to given value.

### HasServicePrincipalNames

`func (o *IdentityMachineResponseModelAllOf) HasServicePrincipalNames() bool`

HasServicePrincipalNames returns a boolean if a field has been set.

### GetIPAddress

`func (o *IdentityMachineResponseModelAllOf) GetIPAddress() []string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *IdentityMachineResponseModelAllOf) GetIPAddressOk() (*[]string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *IdentityMachineResponseModelAllOf) SetIPAddress(v []string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *IdentityMachineResponseModelAllOf) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### GetIPAddressResolveMethod

`func (o *IdentityMachineResponseModelAllOf) GetIPAddressResolveMethod() IdentityMachineIPAddressResolveMethod`

GetIPAddressResolveMethod returns the IPAddressResolveMethod field if non-nil, zero value otherwise.

### GetIPAddressResolveMethodOk

`func (o *IdentityMachineResponseModelAllOf) GetIPAddressResolveMethodOk() (*IdentityMachineIPAddressResolveMethod, bool)`

GetIPAddressResolveMethodOk returns a tuple with the IPAddressResolveMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddressResolveMethod

`func (o *IdentityMachineResponseModelAllOf) SetIPAddressResolveMethod(v IdentityMachineIPAddressResolveMethod)`

SetIPAddressResolveMethod sets IPAddressResolveMethod field to given value.

### HasIPAddressResolveMethod

`func (o *IdentityMachineResponseModelAllOf) HasIPAddressResolveMethod() bool`

HasIPAddressResolveMethod returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityMachineResponseModelAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityMachineResponseModelAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityMachineResponseModelAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentityMachineResponseModelAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocked

`func (o *IdentityMachineResponseModelAllOf) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *IdentityMachineResponseModelAllOf) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *IdentityMachineResponseModelAllOf) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *IdentityMachineResponseModelAllOf) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *IdentityMachineResponseModelAllOf) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *IdentityMachineResponseModelAllOf) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *IdentityMachineResponseModelAllOf) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


