# ApiGetComputersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **NullableString** |  | [optional] 
**SecurityId** | Pointer to **NullableString** |  | [optional] 
**DnsHostName** | Pointer to **NullableString** |  | [optional] 
**ServicePrincipalNames** | Pointer to **[]string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsAccountLocked** | Pointer to **bool** |  | [optional] 
**ComputerGroupList** | Pointer to **[]string** |  | [optional] 
**UserCertificate** | Pointer to **NullableString** |  | [optional] 
**DomainName** | Pointer to **NullableString** |  | [optional] 
**ForestName** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**DistinguishedName** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CanonicalName** | Pointer to **NullableString** |  | [optional] 

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

### SetAccountNameNil

`func (o *ApiGetComputersResult) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *ApiGetComputersResult) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
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

### SetSecurityIdNil

`func (o *ApiGetComputersResult) SetSecurityIdNil(b bool)`

 SetSecurityIdNil sets the value for SecurityId to be an explicit nil

### UnsetSecurityId
`func (o *ApiGetComputersResult) UnsetSecurityId()`

UnsetSecurityId ensures that no value is present for SecurityId, not even an explicit nil
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

### SetDnsHostNameNil

`func (o *ApiGetComputersResult) SetDnsHostNameNil(b bool)`

 SetDnsHostNameNil sets the value for DnsHostName to be an explicit nil

### UnsetDnsHostName
`func (o *ApiGetComputersResult) UnsetDnsHostName()`

UnsetDnsHostName ensures that no value is present for DnsHostName, not even an explicit nil
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

### SetServicePrincipalNamesNil

`func (o *ApiGetComputersResult) SetServicePrincipalNamesNil(b bool)`

 SetServicePrincipalNamesNil sets the value for ServicePrincipalNames to be an explicit nil

### UnsetServicePrincipalNames
`func (o *ApiGetComputersResult) UnsetServicePrincipalNames()`

UnsetServicePrincipalNames ensures that no value is present for ServicePrincipalNames, not even an explicit nil
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

### SetComputerGroupListNil

`func (o *ApiGetComputersResult) SetComputerGroupListNil(b bool)`

 SetComputerGroupListNil sets the value for ComputerGroupList to be an explicit nil

### UnsetComputerGroupList
`func (o *ApiGetComputersResult) UnsetComputerGroupList()`

UnsetComputerGroupList ensures that no value is present for ComputerGroupList, not even an explicit nil
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

### SetUserCertificateNil

`func (o *ApiGetComputersResult) SetUserCertificateNil(b bool)`

 SetUserCertificateNil sets the value for UserCertificate to be an explicit nil

### UnsetUserCertificate
`func (o *ApiGetComputersResult) UnsetUserCertificate()`

UnsetUserCertificate ensures that no value is present for UserCertificate, not even an explicit nil
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

### SetDomainNameNil

`func (o *ApiGetComputersResult) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *ApiGetComputersResult) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
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

### SetForestNameNil

`func (o *ApiGetComputersResult) SetForestNameNil(b bool)`

 SetForestNameNil sets the value for ForestName to be an explicit nil

### UnsetForestName
`func (o *ApiGetComputersResult) UnsetForestName()`

UnsetForestName ensures that no value is present for ForestName, not even an explicit nil
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

### SetDistinguishedNameNil

`func (o *ApiGetComputersResult) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *ApiGetComputersResult) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
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

### SetNameNil

`func (o *ApiGetComputersResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiGetComputersResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### SetCanonicalNameNil

`func (o *ApiGetComputersResult) SetCanonicalNameNil(b bool)`

 SetCanonicalNameNil sets the value for CanonicalName to be an explicit nil

### UnsetCanonicalName
`func (o *ApiGetComputersResult) UnsetCanonicalName()`

UnsetCanonicalName ensures that no value is present for CanonicalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


