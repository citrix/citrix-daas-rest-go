# ApiGetComputersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** |  | [optional] 
**SecurityId** | Pointer to **string** |  | [optional] 
**DnsHostName** | Pointer to **string** |  | [optional] 
**ServicePrincipalNames** | Pointer to **[]string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsAccountLocked** | Pointer to **bool** |  | [optional] 
**ComputerGroupList** | Pointer to **[]string** |  | [optional] 
**UserCertificate** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**ForestName** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**DistinguishedName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CanonicalName** | Pointer to **string** |  | [optional] 

## Methods

### NewApiGetComputersResult

`func NewApiGetComputersResult() *ApiGetComputersResult`

NewApiGetComputersResult instantiates a new ApiGetComputersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiGetComputersResultWithDefaults

`func NewApiGetComputersResultWithDefaults() *ApiGetComputersResult`

NewApiGetComputersResultWithDefaults instantiates a new ApiGetComputersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *ApiGetComputersResult) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ApiGetComputersResult) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ApiGetComputersResult) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ApiGetComputersResult) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetSecurityId

`func (o *ApiGetComputersResult) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *ApiGetComputersResult) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *ApiGetComputersResult) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *ApiGetComputersResult) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetDnsHostName

`func (o *ApiGetComputersResult) GetDnsHostName() string`

GetDnsHostName returns the DnsHostName field if non-nil, zero value otherwise.

### GetDnsHostNameOk

`func (o *ApiGetComputersResult) GetDnsHostNameOk() (*string, bool)`

GetDnsHostNameOk returns a tuple with the DnsHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHostName

`func (o *ApiGetComputersResult) SetDnsHostName(v string)`

SetDnsHostName sets DnsHostName field to given value.

### HasDnsHostName

`func (o *ApiGetComputersResult) HasDnsHostName() bool`

HasDnsHostName returns a boolean if a field has been set.

### GetServicePrincipalNames

`func (o *ApiGetComputersResult) GetServicePrincipalNames() []string`

GetServicePrincipalNames returns the ServicePrincipalNames field if non-nil, zero value otherwise.

### GetServicePrincipalNamesOk

`func (o *ApiGetComputersResult) GetServicePrincipalNamesOk() (*[]string, bool)`

GetServicePrincipalNamesOk returns a tuple with the ServicePrincipalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalNames

`func (o *ApiGetComputersResult) SetServicePrincipalNames(v []string)`

SetServicePrincipalNames sets ServicePrincipalNames field to given value.

### HasServicePrincipalNames

`func (o *ApiGetComputersResult) HasServicePrincipalNames() bool`

HasServicePrincipalNames returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ApiGetComputersResult) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiGetComputersResult) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiGetComputersResult) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ApiGetComputersResult) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsAccountLocked

`func (o *ApiGetComputersResult) GetIsAccountLocked() bool`

GetIsAccountLocked returns the IsAccountLocked field if non-nil, zero value otherwise.

### GetIsAccountLockedOk

`func (o *ApiGetComputersResult) GetIsAccountLockedOk() (*bool, bool)`

GetIsAccountLockedOk returns a tuple with the IsAccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountLocked

`func (o *ApiGetComputersResult) SetIsAccountLocked(v bool)`

SetIsAccountLocked sets IsAccountLocked field to given value.

### HasIsAccountLocked

`func (o *ApiGetComputersResult) HasIsAccountLocked() bool`

HasIsAccountLocked returns a boolean if a field has been set.

### GetComputerGroupList

`func (o *ApiGetComputersResult) GetComputerGroupList() []string`

GetComputerGroupList returns the ComputerGroupList field if non-nil, zero value otherwise.

### GetComputerGroupListOk

`func (o *ApiGetComputersResult) GetComputerGroupListOk() (*[]string, bool)`

GetComputerGroupListOk returns a tuple with the ComputerGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerGroupList

`func (o *ApiGetComputersResult) SetComputerGroupList(v []string)`

SetComputerGroupList sets ComputerGroupList field to given value.

### HasComputerGroupList

`func (o *ApiGetComputersResult) HasComputerGroupList() bool`

HasComputerGroupList returns a boolean if a field has been set.

### GetUserCertificate

`func (o *ApiGetComputersResult) GetUserCertificate() string`

GetUserCertificate returns the UserCertificate field if non-nil, zero value otherwise.

### GetUserCertificateOk

`func (o *ApiGetComputersResult) GetUserCertificateOk() (*string, bool)`

GetUserCertificateOk returns a tuple with the UserCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCertificate

`func (o *ApiGetComputersResult) SetUserCertificate(v string)`

SetUserCertificate sets UserCertificate field to given value.

### HasUserCertificate

`func (o *ApiGetComputersResult) HasUserCertificate() bool`

HasUserCertificate returns a boolean if a field has been set.

### GetDomainName

`func (o *ApiGetComputersResult) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ApiGetComputersResult) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ApiGetComputersResult) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ApiGetComputersResult) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetForestName

`func (o *ApiGetComputersResult) GetForestName() string`

GetForestName returns the ForestName field if non-nil, zero value otherwise.

### GetForestNameOk

`func (o *ApiGetComputersResult) GetForestNameOk() (*string, bool)`

GetForestNameOk returns a tuple with the ForestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestName

`func (o *ApiGetComputersResult) SetForestName(v string)`

SetForestName sets ForestName field to given value.

### HasForestName

`func (o *ApiGetComputersResult) HasForestName() bool`

HasForestName returns a boolean if a field has been set.

### GetGuid

`func (o *ApiGetComputersResult) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ApiGetComputersResult) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ApiGetComputersResult) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ApiGetComputersResult) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *ApiGetComputersResult) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *ApiGetComputersResult) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *ApiGetComputersResult) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *ApiGetComputersResult) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetName

`func (o *ApiGetComputersResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiGetComputersResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiGetComputersResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiGetComputersResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCanonicalName

`func (o *ApiGetComputersResult) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *ApiGetComputersResult) GetCanonicalNameOk() (*string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalName

`func (o *ApiGetComputersResult) SetCanonicalName(v string)`

SetCanonicalName sets CanonicalName field to given value.

### HasCanonicalName

`func (o *ApiGetComputersResult) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


