# IdentityMachineResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The Domain the object belongs to. NOTE: When doing a forest level search, the domain property is populated only if the object contains a SID and the SID was requested as property to be fetched. | [optional] 
**Forest** | Pointer to **string** | The forest the object belongs to. | [optional] 
**Guid** | Pointer to **string** | The GUID of the object, the GUID is unique across the enterprise and anywhere else. | [optional] 
**DistinguishedName** | Pointer to **string** | The Distinguished name of the object. Distinguished names (DNs) are unique and they unambiguously identify objects in the directory. | [optional] 
**CanonicalName** | Pointer to **string** | The unique distinguished name of the object in canonical format. | [optional] 
**PossibleLookupFailure** | Pointer to **bool** | Set to True during a bulk identity lookup operation. For individual identity lookup, usually an exception will be thrown. | [optional] 
**DirectoryServer** | Pointer to **string** | Domain controller that performed the add/update operation on this object | [optional] 
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

### NewIdentityMachineResponseModel

`func NewIdentityMachineResponseModel(propertiesFetched int32, ) *IdentityMachineResponseModel`

NewIdentityMachineResponseModel instantiates a new IdentityMachineResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityMachineResponseModelWithDefaults

`func NewIdentityMachineResponseModelWithDefaults() *IdentityMachineResponseModel`

NewIdentityMachineResponseModelWithDefaults instantiates a new IdentityMachineResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *IdentityMachineResponseModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityMachineResponseModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityMachineResponseModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentityMachineResponseModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetForest

`func (o *IdentityMachineResponseModel) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *IdentityMachineResponseModel) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *IdentityMachineResponseModel) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *IdentityMachineResponseModel) HasForest() bool`

HasForest returns a boolean if a field has been set.

### GetGuid

`func (o *IdentityMachineResponseModel) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IdentityMachineResponseModel) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IdentityMachineResponseModel) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IdentityMachineResponseModel) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *IdentityMachineResponseModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *IdentityMachineResponseModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *IdentityMachineResponseModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *IdentityMachineResponseModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetCanonicalName

`func (o *IdentityMachineResponseModel) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *IdentityMachineResponseModel) GetCanonicalNameOk() (*string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalName

`func (o *IdentityMachineResponseModel) SetCanonicalName(v string)`

SetCanonicalName sets CanonicalName field to given value.

### HasCanonicalName

`func (o *IdentityMachineResponseModel) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.

### GetPossibleLookupFailure

`func (o *IdentityMachineResponseModel) GetPossibleLookupFailure() bool`

GetPossibleLookupFailure returns the PossibleLookupFailure field if non-nil, zero value otherwise.

### GetPossibleLookupFailureOk

`func (o *IdentityMachineResponseModel) GetPossibleLookupFailureOk() (*bool, bool)`

GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleLookupFailure

`func (o *IdentityMachineResponseModel) SetPossibleLookupFailure(v bool)`

SetPossibleLookupFailure sets PossibleLookupFailure field to given value.

### HasPossibleLookupFailure

`func (o *IdentityMachineResponseModel) HasPossibleLookupFailure() bool`

HasPossibleLookupFailure returns a boolean if a field has been set.

### GetDirectoryServer

`func (o *IdentityMachineResponseModel) GetDirectoryServer() string`

GetDirectoryServer returns the DirectoryServer field if non-nil, zero value otherwise.

### GetDirectoryServerOk

`func (o *IdentityMachineResponseModel) GetDirectoryServerOk() (*string, bool)`

GetDirectoryServerOk returns a tuple with the DirectoryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServer

`func (o *IdentityMachineResponseModel) SetDirectoryServer(v string)`

SetDirectoryServer sets DirectoryServer field to given value.

### HasDirectoryServer

`func (o *IdentityMachineResponseModel) HasDirectoryServer() bool`

HasDirectoryServer returns a boolean if a field has been set.

### GetSamName

`func (o *IdentityMachineResponseModel) GetSamName() string`

GetSamName returns the SamName field if non-nil, zero value otherwise.

### GetSamNameOk

`func (o *IdentityMachineResponseModel) GetSamNameOk() (*string, bool)`

GetSamNameOk returns a tuple with the SamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamName

`func (o *IdentityMachineResponseModel) SetSamName(v string)`

SetSamName sets SamName field to given value.

### HasSamName

`func (o *IdentityMachineResponseModel) HasSamName() bool`

HasSamName returns a boolean if a field has been set.

### GetSamAccountName

`func (o *IdentityMachineResponseModel) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *IdentityMachineResponseModel) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *IdentityMachineResponseModel) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *IdentityMachineResponseModel) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.

### GetDnsName

`func (o *IdentityMachineResponseModel) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *IdentityMachineResponseModel) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *IdentityMachineResponseModel) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *IdentityMachineResponseModel) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetSid

`func (o *IdentityMachineResponseModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IdentityMachineResponseModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IdentityMachineResponseModel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IdentityMachineResponseModel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetServicePrincipalNames

`func (o *IdentityMachineResponseModel) GetServicePrincipalNames() []string`

GetServicePrincipalNames returns the ServicePrincipalNames field if non-nil, zero value otherwise.

### GetServicePrincipalNamesOk

`func (o *IdentityMachineResponseModel) GetServicePrincipalNamesOk() (*[]string, bool)`

GetServicePrincipalNamesOk returns a tuple with the ServicePrincipalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalNames

`func (o *IdentityMachineResponseModel) SetServicePrincipalNames(v []string)`

SetServicePrincipalNames sets ServicePrincipalNames field to given value.

### HasServicePrincipalNames

`func (o *IdentityMachineResponseModel) HasServicePrincipalNames() bool`

HasServicePrincipalNames returns a boolean if a field has been set.

### GetIPAddress

`func (o *IdentityMachineResponseModel) GetIPAddress() []string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *IdentityMachineResponseModel) GetIPAddressOk() (*[]string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *IdentityMachineResponseModel) SetIPAddress(v []string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *IdentityMachineResponseModel) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### GetIPAddressResolveMethod

`func (o *IdentityMachineResponseModel) GetIPAddressResolveMethod() IdentityMachineIPAddressResolveMethod`

GetIPAddressResolveMethod returns the IPAddressResolveMethod field if non-nil, zero value otherwise.

### GetIPAddressResolveMethodOk

`func (o *IdentityMachineResponseModel) GetIPAddressResolveMethodOk() (*IdentityMachineIPAddressResolveMethod, bool)`

GetIPAddressResolveMethodOk returns a tuple with the IPAddressResolveMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddressResolveMethod

`func (o *IdentityMachineResponseModel) SetIPAddressResolveMethod(v IdentityMachineIPAddressResolveMethod)`

SetIPAddressResolveMethod sets IPAddressResolveMethod field to given value.

### HasIPAddressResolveMethod

`func (o *IdentityMachineResponseModel) HasIPAddressResolveMethod() bool`

HasIPAddressResolveMethod returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityMachineResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityMachineResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityMachineResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentityMachineResponseModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocked

`func (o *IdentityMachineResponseModel) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *IdentityMachineResponseModel) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *IdentityMachineResponseModel) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *IdentityMachineResponseModel) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *IdentityMachineResponseModel) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *IdentityMachineResponseModel) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *IdentityMachineResponseModel) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


