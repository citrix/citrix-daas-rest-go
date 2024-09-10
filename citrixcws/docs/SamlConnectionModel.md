# SamlConnectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamlEntityId** | **string** |  | 
**SamlSignAuthRequest** | **string** |  | 
**SamlSpEntityIdSuffix** | Pointer to **NullableString** |  | [optional] 
**SamlSingleSignOnServiceUrl** | Pointer to **NullableString** |  | [optional] 
**SamlSingleSignOnServiceBinding** | **string** |  | 
**SamlResponse** | **string** |  | 
**SamlAuthDomainName** | **string** |  | 
**SamlAuthenticationContext** | **string** |  | 
**SamlAuthenticationContextComparison** | **string** |  | 
**SamlLogoutUrl** | Pointer to **NullableString** |  | [optional] 
**SamlLogoutRequestBinding** | Pointer to **NullableString** |  | [optional] 
**SamlSignLogoutRequest** | Pointer to **string** |  | [optional] 
**SamlAttributeNameForSid** | Pointer to **NullableString** |  | [optional] 
**SamlAttributeNameForUpn** | Pointer to **NullableString** |  | [optional] 
**SamlAttributeNameForEmail** | Pointer to **NullableString** |  | [optional] 
**SamlAttributeNameForAdOid** | Pointer to **NullableString** |  | [optional] 
**SamlAttributeNameForAdForest** | Pointer to **NullableString** |  | [optional] 
**SamlAttributeNameForAdDomain** | Pointer to **NullableString** |  | [optional] 
**SamlAttributeNameForUserDisplayName** | Pointer to **NullableString** |  | [optional] 
**SamlAttributeNameForUserGivenName** | Pointer to **NullableString** |  | [optional] 
**SamlAttributeNameForUserFamilyName** | Pointer to **NullableString** |  | [optional] 
**SamlCertFileName** | **string** |  | 
**SamlCertInfo** | [**SamlCertInfoModel**](SamlCertInfoModel.md) |  | 

## Methods

### NewSamlConnectionModel

`func NewSamlConnectionModel(samlEntityId string, samlSignAuthRequest string, samlSingleSignOnServiceBinding string, samlResponse string, samlAuthDomainName string, samlAuthenticationContext string, samlAuthenticationContextComparison string, samlCertFileName string, samlCertInfo SamlCertInfoModel, ) *SamlConnectionModel`

NewSamlConnectionModel instantiates a new SamlConnectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlConnectionModelWithDefaults

`func NewSamlConnectionModelWithDefaults() *SamlConnectionModel`

NewSamlConnectionModelWithDefaults instantiates a new SamlConnectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamlEntityId

`func (o *SamlConnectionModel) GetSamlEntityId() string`

GetSamlEntityId returns the SamlEntityId field if non-nil, zero value otherwise.

### GetSamlEntityIdOk

`func (o *SamlConnectionModel) GetSamlEntityIdOk() (*string, bool)`

GetSamlEntityIdOk returns a tuple with the SamlEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlEntityId

`func (o *SamlConnectionModel) SetSamlEntityId(v string)`

SetSamlEntityId sets SamlEntityId field to given value.


### GetSamlSignAuthRequest

`func (o *SamlConnectionModel) GetSamlSignAuthRequest() string`

GetSamlSignAuthRequest returns the SamlSignAuthRequest field if non-nil, zero value otherwise.

### GetSamlSignAuthRequestOk

`func (o *SamlConnectionModel) GetSamlSignAuthRequestOk() (*string, bool)`

GetSamlSignAuthRequestOk returns a tuple with the SamlSignAuthRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSignAuthRequest

`func (o *SamlConnectionModel) SetSamlSignAuthRequest(v string)`

SetSamlSignAuthRequest sets SamlSignAuthRequest field to given value.


### GetSamlSpEntityIdSuffix

`func (o *SamlConnectionModel) GetSamlSpEntityIdSuffix() string`

GetSamlSpEntityIdSuffix returns the SamlSpEntityIdSuffix field if non-nil, zero value otherwise.

### GetSamlSpEntityIdSuffixOk

`func (o *SamlConnectionModel) GetSamlSpEntityIdSuffixOk() (*string, bool)`

GetSamlSpEntityIdSuffixOk returns a tuple with the SamlSpEntityIdSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSpEntityIdSuffix

`func (o *SamlConnectionModel) SetSamlSpEntityIdSuffix(v string)`

SetSamlSpEntityIdSuffix sets SamlSpEntityIdSuffix field to given value.

### HasSamlSpEntityIdSuffix

`func (o *SamlConnectionModel) HasSamlSpEntityIdSuffix() bool`

HasSamlSpEntityIdSuffix returns a boolean if a field has been set.

### SetSamlSpEntityIdSuffixNil

`func (o *SamlConnectionModel) SetSamlSpEntityIdSuffixNil(b bool)`

 SetSamlSpEntityIdSuffixNil sets the value for SamlSpEntityIdSuffix to be an explicit nil

### UnsetSamlSpEntityIdSuffix
`func (o *SamlConnectionModel) UnsetSamlSpEntityIdSuffix()`

UnsetSamlSpEntityIdSuffix ensures that no value is present for SamlSpEntityIdSuffix, not even an explicit nil
### GetSamlSingleSignOnServiceUrl

`func (o *SamlConnectionModel) GetSamlSingleSignOnServiceUrl() string`

GetSamlSingleSignOnServiceUrl returns the SamlSingleSignOnServiceUrl field if non-nil, zero value otherwise.

### GetSamlSingleSignOnServiceUrlOk

`func (o *SamlConnectionModel) GetSamlSingleSignOnServiceUrlOk() (*string, bool)`

GetSamlSingleSignOnServiceUrlOk returns a tuple with the SamlSingleSignOnServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSingleSignOnServiceUrl

`func (o *SamlConnectionModel) SetSamlSingleSignOnServiceUrl(v string)`

SetSamlSingleSignOnServiceUrl sets SamlSingleSignOnServiceUrl field to given value.

### HasSamlSingleSignOnServiceUrl

`func (o *SamlConnectionModel) HasSamlSingleSignOnServiceUrl() bool`

HasSamlSingleSignOnServiceUrl returns a boolean if a field has been set.

### SetSamlSingleSignOnServiceUrlNil

`func (o *SamlConnectionModel) SetSamlSingleSignOnServiceUrlNil(b bool)`

 SetSamlSingleSignOnServiceUrlNil sets the value for SamlSingleSignOnServiceUrl to be an explicit nil

### UnsetSamlSingleSignOnServiceUrl
`func (o *SamlConnectionModel) UnsetSamlSingleSignOnServiceUrl()`

UnsetSamlSingleSignOnServiceUrl ensures that no value is present for SamlSingleSignOnServiceUrl, not even an explicit nil
### GetSamlSingleSignOnServiceBinding

`func (o *SamlConnectionModel) GetSamlSingleSignOnServiceBinding() string`

GetSamlSingleSignOnServiceBinding returns the SamlSingleSignOnServiceBinding field if non-nil, zero value otherwise.

### GetSamlSingleSignOnServiceBindingOk

`func (o *SamlConnectionModel) GetSamlSingleSignOnServiceBindingOk() (*string, bool)`

GetSamlSingleSignOnServiceBindingOk returns a tuple with the SamlSingleSignOnServiceBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSingleSignOnServiceBinding

`func (o *SamlConnectionModel) SetSamlSingleSignOnServiceBinding(v string)`

SetSamlSingleSignOnServiceBinding sets SamlSingleSignOnServiceBinding field to given value.


### GetSamlResponse

`func (o *SamlConnectionModel) GetSamlResponse() string`

GetSamlResponse returns the SamlResponse field if non-nil, zero value otherwise.

### GetSamlResponseOk

`func (o *SamlConnectionModel) GetSamlResponseOk() (*string, bool)`

GetSamlResponseOk returns a tuple with the SamlResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlResponse

`func (o *SamlConnectionModel) SetSamlResponse(v string)`

SetSamlResponse sets SamlResponse field to given value.


### GetSamlAuthDomainName

`func (o *SamlConnectionModel) GetSamlAuthDomainName() string`

GetSamlAuthDomainName returns the SamlAuthDomainName field if non-nil, zero value otherwise.

### GetSamlAuthDomainNameOk

`func (o *SamlConnectionModel) GetSamlAuthDomainNameOk() (*string, bool)`

GetSamlAuthDomainNameOk returns a tuple with the SamlAuthDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAuthDomainName

`func (o *SamlConnectionModel) SetSamlAuthDomainName(v string)`

SetSamlAuthDomainName sets SamlAuthDomainName field to given value.


### GetSamlAuthenticationContext

`func (o *SamlConnectionModel) GetSamlAuthenticationContext() string`

GetSamlAuthenticationContext returns the SamlAuthenticationContext field if non-nil, zero value otherwise.

### GetSamlAuthenticationContextOk

`func (o *SamlConnectionModel) GetSamlAuthenticationContextOk() (*string, bool)`

GetSamlAuthenticationContextOk returns a tuple with the SamlAuthenticationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAuthenticationContext

`func (o *SamlConnectionModel) SetSamlAuthenticationContext(v string)`

SetSamlAuthenticationContext sets SamlAuthenticationContext field to given value.


### GetSamlAuthenticationContextComparison

`func (o *SamlConnectionModel) GetSamlAuthenticationContextComparison() string`

GetSamlAuthenticationContextComparison returns the SamlAuthenticationContextComparison field if non-nil, zero value otherwise.

### GetSamlAuthenticationContextComparisonOk

`func (o *SamlConnectionModel) GetSamlAuthenticationContextComparisonOk() (*string, bool)`

GetSamlAuthenticationContextComparisonOk returns a tuple with the SamlAuthenticationContextComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAuthenticationContextComparison

`func (o *SamlConnectionModel) SetSamlAuthenticationContextComparison(v string)`

SetSamlAuthenticationContextComparison sets SamlAuthenticationContextComparison field to given value.


### GetSamlLogoutUrl

`func (o *SamlConnectionModel) GetSamlLogoutUrl() string`

GetSamlLogoutUrl returns the SamlLogoutUrl field if non-nil, zero value otherwise.

### GetSamlLogoutUrlOk

`func (o *SamlConnectionModel) GetSamlLogoutUrlOk() (*string, bool)`

GetSamlLogoutUrlOk returns a tuple with the SamlLogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLogoutUrl

`func (o *SamlConnectionModel) SetSamlLogoutUrl(v string)`

SetSamlLogoutUrl sets SamlLogoutUrl field to given value.

### HasSamlLogoutUrl

`func (o *SamlConnectionModel) HasSamlLogoutUrl() bool`

HasSamlLogoutUrl returns a boolean if a field has been set.

### SetSamlLogoutUrlNil

`func (o *SamlConnectionModel) SetSamlLogoutUrlNil(b bool)`

 SetSamlLogoutUrlNil sets the value for SamlLogoutUrl to be an explicit nil

### UnsetSamlLogoutUrl
`func (o *SamlConnectionModel) UnsetSamlLogoutUrl()`

UnsetSamlLogoutUrl ensures that no value is present for SamlLogoutUrl, not even an explicit nil
### GetSamlLogoutRequestBinding

`func (o *SamlConnectionModel) GetSamlLogoutRequestBinding() string`

GetSamlLogoutRequestBinding returns the SamlLogoutRequestBinding field if non-nil, zero value otherwise.

### GetSamlLogoutRequestBindingOk

`func (o *SamlConnectionModel) GetSamlLogoutRequestBindingOk() (*string, bool)`

GetSamlLogoutRequestBindingOk returns a tuple with the SamlLogoutRequestBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLogoutRequestBinding

`func (o *SamlConnectionModel) SetSamlLogoutRequestBinding(v string)`

SetSamlLogoutRequestBinding sets SamlLogoutRequestBinding field to given value.

### HasSamlLogoutRequestBinding

`func (o *SamlConnectionModel) HasSamlLogoutRequestBinding() bool`

HasSamlLogoutRequestBinding returns a boolean if a field has been set.

### SetSamlLogoutRequestBindingNil

`func (o *SamlConnectionModel) SetSamlLogoutRequestBindingNil(b bool)`

 SetSamlLogoutRequestBindingNil sets the value for SamlLogoutRequestBinding to be an explicit nil

### UnsetSamlLogoutRequestBinding
`func (o *SamlConnectionModel) UnsetSamlLogoutRequestBinding()`

UnsetSamlLogoutRequestBinding ensures that no value is present for SamlLogoutRequestBinding, not even an explicit nil
### GetSamlSignLogoutRequest

`func (o *SamlConnectionModel) GetSamlSignLogoutRequest() string`

GetSamlSignLogoutRequest returns the SamlSignLogoutRequest field if non-nil, zero value otherwise.

### GetSamlSignLogoutRequestOk

`func (o *SamlConnectionModel) GetSamlSignLogoutRequestOk() (*string, bool)`

GetSamlSignLogoutRequestOk returns a tuple with the SamlSignLogoutRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSignLogoutRequest

`func (o *SamlConnectionModel) SetSamlSignLogoutRequest(v string)`

SetSamlSignLogoutRequest sets SamlSignLogoutRequest field to given value.

### HasSamlSignLogoutRequest

`func (o *SamlConnectionModel) HasSamlSignLogoutRequest() bool`

HasSamlSignLogoutRequest returns a boolean if a field has been set.

### GetSamlAttributeNameForSid

`func (o *SamlConnectionModel) GetSamlAttributeNameForSid() string`

GetSamlAttributeNameForSid returns the SamlAttributeNameForSid field if non-nil, zero value otherwise.

### GetSamlAttributeNameForSidOk

`func (o *SamlConnectionModel) GetSamlAttributeNameForSidOk() (*string, bool)`

GetSamlAttributeNameForSidOk returns a tuple with the SamlAttributeNameForSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForSid

`func (o *SamlConnectionModel) SetSamlAttributeNameForSid(v string)`

SetSamlAttributeNameForSid sets SamlAttributeNameForSid field to given value.

### HasSamlAttributeNameForSid

`func (o *SamlConnectionModel) HasSamlAttributeNameForSid() bool`

HasSamlAttributeNameForSid returns a boolean if a field has been set.

### SetSamlAttributeNameForSidNil

`func (o *SamlConnectionModel) SetSamlAttributeNameForSidNil(b bool)`

 SetSamlAttributeNameForSidNil sets the value for SamlAttributeNameForSid to be an explicit nil

### UnsetSamlAttributeNameForSid
`func (o *SamlConnectionModel) UnsetSamlAttributeNameForSid()`

UnsetSamlAttributeNameForSid ensures that no value is present for SamlAttributeNameForSid, not even an explicit nil
### GetSamlAttributeNameForUpn

`func (o *SamlConnectionModel) GetSamlAttributeNameForUpn() string`

GetSamlAttributeNameForUpn returns the SamlAttributeNameForUpn field if non-nil, zero value otherwise.

### GetSamlAttributeNameForUpnOk

`func (o *SamlConnectionModel) GetSamlAttributeNameForUpnOk() (*string, bool)`

GetSamlAttributeNameForUpnOk returns a tuple with the SamlAttributeNameForUpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForUpn

`func (o *SamlConnectionModel) SetSamlAttributeNameForUpn(v string)`

SetSamlAttributeNameForUpn sets SamlAttributeNameForUpn field to given value.

### HasSamlAttributeNameForUpn

`func (o *SamlConnectionModel) HasSamlAttributeNameForUpn() bool`

HasSamlAttributeNameForUpn returns a boolean if a field has been set.

### SetSamlAttributeNameForUpnNil

`func (o *SamlConnectionModel) SetSamlAttributeNameForUpnNil(b bool)`

 SetSamlAttributeNameForUpnNil sets the value for SamlAttributeNameForUpn to be an explicit nil

### UnsetSamlAttributeNameForUpn
`func (o *SamlConnectionModel) UnsetSamlAttributeNameForUpn()`

UnsetSamlAttributeNameForUpn ensures that no value is present for SamlAttributeNameForUpn, not even an explicit nil
### GetSamlAttributeNameForEmail

`func (o *SamlConnectionModel) GetSamlAttributeNameForEmail() string`

GetSamlAttributeNameForEmail returns the SamlAttributeNameForEmail field if non-nil, zero value otherwise.

### GetSamlAttributeNameForEmailOk

`func (o *SamlConnectionModel) GetSamlAttributeNameForEmailOk() (*string, bool)`

GetSamlAttributeNameForEmailOk returns a tuple with the SamlAttributeNameForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForEmail

`func (o *SamlConnectionModel) SetSamlAttributeNameForEmail(v string)`

SetSamlAttributeNameForEmail sets SamlAttributeNameForEmail field to given value.

### HasSamlAttributeNameForEmail

`func (o *SamlConnectionModel) HasSamlAttributeNameForEmail() bool`

HasSamlAttributeNameForEmail returns a boolean if a field has been set.

### SetSamlAttributeNameForEmailNil

`func (o *SamlConnectionModel) SetSamlAttributeNameForEmailNil(b bool)`

 SetSamlAttributeNameForEmailNil sets the value for SamlAttributeNameForEmail to be an explicit nil

### UnsetSamlAttributeNameForEmail
`func (o *SamlConnectionModel) UnsetSamlAttributeNameForEmail()`

UnsetSamlAttributeNameForEmail ensures that no value is present for SamlAttributeNameForEmail, not even an explicit nil
### GetSamlAttributeNameForAdOid

`func (o *SamlConnectionModel) GetSamlAttributeNameForAdOid() string`

GetSamlAttributeNameForAdOid returns the SamlAttributeNameForAdOid field if non-nil, zero value otherwise.

### GetSamlAttributeNameForAdOidOk

`func (o *SamlConnectionModel) GetSamlAttributeNameForAdOidOk() (*string, bool)`

GetSamlAttributeNameForAdOidOk returns a tuple with the SamlAttributeNameForAdOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForAdOid

`func (o *SamlConnectionModel) SetSamlAttributeNameForAdOid(v string)`

SetSamlAttributeNameForAdOid sets SamlAttributeNameForAdOid field to given value.

### HasSamlAttributeNameForAdOid

`func (o *SamlConnectionModel) HasSamlAttributeNameForAdOid() bool`

HasSamlAttributeNameForAdOid returns a boolean if a field has been set.

### SetSamlAttributeNameForAdOidNil

`func (o *SamlConnectionModel) SetSamlAttributeNameForAdOidNil(b bool)`

 SetSamlAttributeNameForAdOidNil sets the value for SamlAttributeNameForAdOid to be an explicit nil

### UnsetSamlAttributeNameForAdOid
`func (o *SamlConnectionModel) UnsetSamlAttributeNameForAdOid()`

UnsetSamlAttributeNameForAdOid ensures that no value is present for SamlAttributeNameForAdOid, not even an explicit nil
### GetSamlAttributeNameForAdForest

`func (o *SamlConnectionModel) GetSamlAttributeNameForAdForest() string`

GetSamlAttributeNameForAdForest returns the SamlAttributeNameForAdForest field if non-nil, zero value otherwise.

### GetSamlAttributeNameForAdForestOk

`func (o *SamlConnectionModel) GetSamlAttributeNameForAdForestOk() (*string, bool)`

GetSamlAttributeNameForAdForestOk returns a tuple with the SamlAttributeNameForAdForest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForAdForest

`func (o *SamlConnectionModel) SetSamlAttributeNameForAdForest(v string)`

SetSamlAttributeNameForAdForest sets SamlAttributeNameForAdForest field to given value.

### HasSamlAttributeNameForAdForest

`func (o *SamlConnectionModel) HasSamlAttributeNameForAdForest() bool`

HasSamlAttributeNameForAdForest returns a boolean if a field has been set.

### SetSamlAttributeNameForAdForestNil

`func (o *SamlConnectionModel) SetSamlAttributeNameForAdForestNil(b bool)`

 SetSamlAttributeNameForAdForestNil sets the value for SamlAttributeNameForAdForest to be an explicit nil

### UnsetSamlAttributeNameForAdForest
`func (o *SamlConnectionModel) UnsetSamlAttributeNameForAdForest()`

UnsetSamlAttributeNameForAdForest ensures that no value is present for SamlAttributeNameForAdForest, not even an explicit nil
### GetSamlAttributeNameForAdDomain

`func (o *SamlConnectionModel) GetSamlAttributeNameForAdDomain() string`

GetSamlAttributeNameForAdDomain returns the SamlAttributeNameForAdDomain field if non-nil, zero value otherwise.

### GetSamlAttributeNameForAdDomainOk

`func (o *SamlConnectionModel) GetSamlAttributeNameForAdDomainOk() (*string, bool)`

GetSamlAttributeNameForAdDomainOk returns a tuple with the SamlAttributeNameForAdDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForAdDomain

`func (o *SamlConnectionModel) SetSamlAttributeNameForAdDomain(v string)`

SetSamlAttributeNameForAdDomain sets SamlAttributeNameForAdDomain field to given value.

### HasSamlAttributeNameForAdDomain

`func (o *SamlConnectionModel) HasSamlAttributeNameForAdDomain() bool`

HasSamlAttributeNameForAdDomain returns a boolean if a field has been set.

### SetSamlAttributeNameForAdDomainNil

`func (o *SamlConnectionModel) SetSamlAttributeNameForAdDomainNil(b bool)`

 SetSamlAttributeNameForAdDomainNil sets the value for SamlAttributeNameForAdDomain to be an explicit nil

### UnsetSamlAttributeNameForAdDomain
`func (o *SamlConnectionModel) UnsetSamlAttributeNameForAdDomain()`

UnsetSamlAttributeNameForAdDomain ensures that no value is present for SamlAttributeNameForAdDomain, not even an explicit nil
### GetSamlAttributeNameForUserDisplayName

`func (o *SamlConnectionModel) GetSamlAttributeNameForUserDisplayName() string`

GetSamlAttributeNameForUserDisplayName returns the SamlAttributeNameForUserDisplayName field if non-nil, zero value otherwise.

### GetSamlAttributeNameForUserDisplayNameOk

`func (o *SamlConnectionModel) GetSamlAttributeNameForUserDisplayNameOk() (*string, bool)`

GetSamlAttributeNameForUserDisplayNameOk returns a tuple with the SamlAttributeNameForUserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForUserDisplayName

`func (o *SamlConnectionModel) SetSamlAttributeNameForUserDisplayName(v string)`

SetSamlAttributeNameForUserDisplayName sets SamlAttributeNameForUserDisplayName field to given value.

### HasSamlAttributeNameForUserDisplayName

`func (o *SamlConnectionModel) HasSamlAttributeNameForUserDisplayName() bool`

HasSamlAttributeNameForUserDisplayName returns a boolean if a field has been set.

### SetSamlAttributeNameForUserDisplayNameNil

`func (o *SamlConnectionModel) SetSamlAttributeNameForUserDisplayNameNil(b bool)`

 SetSamlAttributeNameForUserDisplayNameNil sets the value for SamlAttributeNameForUserDisplayName to be an explicit nil

### UnsetSamlAttributeNameForUserDisplayName
`func (o *SamlConnectionModel) UnsetSamlAttributeNameForUserDisplayName()`

UnsetSamlAttributeNameForUserDisplayName ensures that no value is present for SamlAttributeNameForUserDisplayName, not even an explicit nil
### GetSamlAttributeNameForUserGivenName

`func (o *SamlConnectionModel) GetSamlAttributeNameForUserGivenName() string`

GetSamlAttributeNameForUserGivenName returns the SamlAttributeNameForUserGivenName field if non-nil, zero value otherwise.

### GetSamlAttributeNameForUserGivenNameOk

`func (o *SamlConnectionModel) GetSamlAttributeNameForUserGivenNameOk() (*string, bool)`

GetSamlAttributeNameForUserGivenNameOk returns a tuple with the SamlAttributeNameForUserGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForUserGivenName

`func (o *SamlConnectionModel) SetSamlAttributeNameForUserGivenName(v string)`

SetSamlAttributeNameForUserGivenName sets SamlAttributeNameForUserGivenName field to given value.

### HasSamlAttributeNameForUserGivenName

`func (o *SamlConnectionModel) HasSamlAttributeNameForUserGivenName() bool`

HasSamlAttributeNameForUserGivenName returns a boolean if a field has been set.

### SetSamlAttributeNameForUserGivenNameNil

`func (o *SamlConnectionModel) SetSamlAttributeNameForUserGivenNameNil(b bool)`

 SetSamlAttributeNameForUserGivenNameNil sets the value for SamlAttributeNameForUserGivenName to be an explicit nil

### UnsetSamlAttributeNameForUserGivenName
`func (o *SamlConnectionModel) UnsetSamlAttributeNameForUserGivenName()`

UnsetSamlAttributeNameForUserGivenName ensures that no value is present for SamlAttributeNameForUserGivenName, not even an explicit nil
### GetSamlAttributeNameForUserFamilyName

`func (o *SamlConnectionModel) GetSamlAttributeNameForUserFamilyName() string`

GetSamlAttributeNameForUserFamilyName returns the SamlAttributeNameForUserFamilyName field if non-nil, zero value otherwise.

### GetSamlAttributeNameForUserFamilyNameOk

`func (o *SamlConnectionModel) GetSamlAttributeNameForUserFamilyNameOk() (*string, bool)`

GetSamlAttributeNameForUserFamilyNameOk returns a tuple with the SamlAttributeNameForUserFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForUserFamilyName

`func (o *SamlConnectionModel) SetSamlAttributeNameForUserFamilyName(v string)`

SetSamlAttributeNameForUserFamilyName sets SamlAttributeNameForUserFamilyName field to given value.

### HasSamlAttributeNameForUserFamilyName

`func (o *SamlConnectionModel) HasSamlAttributeNameForUserFamilyName() bool`

HasSamlAttributeNameForUserFamilyName returns a boolean if a field has been set.

### SetSamlAttributeNameForUserFamilyNameNil

`func (o *SamlConnectionModel) SetSamlAttributeNameForUserFamilyNameNil(b bool)`

 SetSamlAttributeNameForUserFamilyNameNil sets the value for SamlAttributeNameForUserFamilyName to be an explicit nil

### UnsetSamlAttributeNameForUserFamilyName
`func (o *SamlConnectionModel) UnsetSamlAttributeNameForUserFamilyName()`

UnsetSamlAttributeNameForUserFamilyName ensures that no value is present for SamlAttributeNameForUserFamilyName, not even an explicit nil
### GetSamlCertFileName

`func (o *SamlConnectionModel) GetSamlCertFileName() string`

GetSamlCertFileName returns the SamlCertFileName field if non-nil, zero value otherwise.

### GetSamlCertFileNameOk

`func (o *SamlConnectionModel) GetSamlCertFileNameOk() (*string, bool)`

GetSamlCertFileNameOk returns a tuple with the SamlCertFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertFileName

`func (o *SamlConnectionModel) SetSamlCertFileName(v string)`

SetSamlCertFileName sets SamlCertFileName field to given value.


### GetSamlCertInfo

`func (o *SamlConnectionModel) GetSamlCertInfo() SamlCertInfoModel`

GetSamlCertInfo returns the SamlCertInfo field if non-nil, zero value otherwise.

### GetSamlCertInfoOk

`func (o *SamlConnectionModel) GetSamlCertInfoOk() (*SamlCertInfoModel, bool)`

GetSamlCertInfoOk returns a tuple with the SamlCertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertInfo

`func (o *SamlConnectionModel) SetSamlCertInfo(v SamlCertInfoModel)`

SetSamlCertInfo sets SamlCertInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


