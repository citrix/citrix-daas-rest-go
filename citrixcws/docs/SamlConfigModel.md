# SamlConfigModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamlEntityId** | Pointer to **NullableString** |  | [optional] 
**SamlSpEntityIdSuffix** | Pointer to **NullableString** |  | [optional] 
**SamlSignAuthRequest** | Pointer to **string** |  | [optional] 
**SamlSingleSignOnServiceUrl** | Pointer to **NullableString** |  | [optional] 
**SamlSingleSignOnServiceBinding** | Pointer to **NullableString** |  | [optional] 
**SamlResponse** | Pointer to **NullableString** |  | [optional] 
**SamlAuthenticationContext** | Pointer to **NullableString** |  | [optional] 
**SamlAuthenticationContextComparison** | Pointer to **NullableString** |  | [optional] 
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
**SamlCertPublicKey** | Pointer to **NullableString** |  | [optional] 
**SamlCertCN** | Pointer to **NullableString** |  | [optional] 
**SamlCertFileName** | Pointer to **NullableString** |  | [optional] 
**SamlCertExpiration** | Pointer to **NullableString** |  | [optional] 
**SamlCertData** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSamlConfigModel

`func NewSamlConfigModel() *SamlConfigModel`

NewSamlConfigModel instantiates a new SamlConfigModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlConfigModelWithDefaults

`func NewSamlConfigModelWithDefaults() *SamlConfigModel`

NewSamlConfigModelWithDefaults instantiates a new SamlConfigModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamlEntityId

`func (o *SamlConfigModel) GetSamlEntityId() string`

GetSamlEntityId returns the SamlEntityId field if non-nil, zero value otherwise.

### GetSamlEntityIdOk

`func (o *SamlConfigModel) GetSamlEntityIdOk() (*string, bool)`

GetSamlEntityIdOk returns a tuple with the SamlEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlEntityId

`func (o *SamlConfigModel) SetSamlEntityId(v string)`

SetSamlEntityId sets SamlEntityId field to given value.

### HasSamlEntityId

`func (o *SamlConfigModel) HasSamlEntityId() bool`

HasSamlEntityId returns a boolean if a field has been set.

### SetSamlEntityIdNil

`func (o *SamlConfigModel) SetSamlEntityIdNil(b bool)`

 SetSamlEntityIdNil sets the value for SamlEntityId to be an explicit nil

### UnsetSamlEntityId
`func (o *SamlConfigModel) UnsetSamlEntityId()`

UnsetSamlEntityId ensures that no value is present for SamlEntityId, not even an explicit nil
### GetSamlSpEntityIdSuffix

`func (o *SamlConfigModel) GetSamlSpEntityIdSuffix() string`

GetSamlSpEntityIdSuffix returns the SamlSpEntityIdSuffix field if non-nil, zero value otherwise.

### GetSamlSpEntityIdSuffixOk

`func (o *SamlConfigModel) GetSamlSpEntityIdSuffixOk() (*string, bool)`

GetSamlSpEntityIdSuffixOk returns a tuple with the SamlSpEntityIdSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSpEntityIdSuffix

`func (o *SamlConfigModel) SetSamlSpEntityIdSuffix(v string)`

SetSamlSpEntityIdSuffix sets SamlSpEntityIdSuffix field to given value.

### HasSamlSpEntityIdSuffix

`func (o *SamlConfigModel) HasSamlSpEntityIdSuffix() bool`

HasSamlSpEntityIdSuffix returns a boolean if a field has been set.

### SetSamlSpEntityIdSuffixNil

`func (o *SamlConfigModel) SetSamlSpEntityIdSuffixNil(b bool)`

 SetSamlSpEntityIdSuffixNil sets the value for SamlSpEntityIdSuffix to be an explicit nil

### UnsetSamlSpEntityIdSuffix
`func (o *SamlConfigModel) UnsetSamlSpEntityIdSuffix()`

UnsetSamlSpEntityIdSuffix ensures that no value is present for SamlSpEntityIdSuffix, not even an explicit nil
### GetSamlSignAuthRequest

`func (o *SamlConfigModel) GetSamlSignAuthRequest() string`

GetSamlSignAuthRequest returns the SamlSignAuthRequest field if non-nil, zero value otherwise.

### GetSamlSignAuthRequestOk

`func (o *SamlConfigModel) GetSamlSignAuthRequestOk() (*string, bool)`

GetSamlSignAuthRequestOk returns a tuple with the SamlSignAuthRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSignAuthRequest

`func (o *SamlConfigModel) SetSamlSignAuthRequest(v string)`

SetSamlSignAuthRequest sets SamlSignAuthRequest field to given value.

### HasSamlSignAuthRequest

`func (o *SamlConfigModel) HasSamlSignAuthRequest() bool`

HasSamlSignAuthRequest returns a boolean if a field has been set.

### GetSamlSingleSignOnServiceUrl

`func (o *SamlConfigModel) GetSamlSingleSignOnServiceUrl() string`

GetSamlSingleSignOnServiceUrl returns the SamlSingleSignOnServiceUrl field if non-nil, zero value otherwise.

### GetSamlSingleSignOnServiceUrlOk

`func (o *SamlConfigModel) GetSamlSingleSignOnServiceUrlOk() (*string, bool)`

GetSamlSingleSignOnServiceUrlOk returns a tuple with the SamlSingleSignOnServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSingleSignOnServiceUrl

`func (o *SamlConfigModel) SetSamlSingleSignOnServiceUrl(v string)`

SetSamlSingleSignOnServiceUrl sets SamlSingleSignOnServiceUrl field to given value.

### HasSamlSingleSignOnServiceUrl

`func (o *SamlConfigModel) HasSamlSingleSignOnServiceUrl() bool`

HasSamlSingleSignOnServiceUrl returns a boolean if a field has been set.

### SetSamlSingleSignOnServiceUrlNil

`func (o *SamlConfigModel) SetSamlSingleSignOnServiceUrlNil(b bool)`

 SetSamlSingleSignOnServiceUrlNil sets the value for SamlSingleSignOnServiceUrl to be an explicit nil

### UnsetSamlSingleSignOnServiceUrl
`func (o *SamlConfigModel) UnsetSamlSingleSignOnServiceUrl()`

UnsetSamlSingleSignOnServiceUrl ensures that no value is present for SamlSingleSignOnServiceUrl, not even an explicit nil
### GetSamlSingleSignOnServiceBinding

`func (o *SamlConfigModel) GetSamlSingleSignOnServiceBinding() string`

GetSamlSingleSignOnServiceBinding returns the SamlSingleSignOnServiceBinding field if non-nil, zero value otherwise.

### GetSamlSingleSignOnServiceBindingOk

`func (o *SamlConfigModel) GetSamlSingleSignOnServiceBindingOk() (*string, bool)`

GetSamlSingleSignOnServiceBindingOk returns a tuple with the SamlSingleSignOnServiceBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSingleSignOnServiceBinding

`func (o *SamlConfigModel) SetSamlSingleSignOnServiceBinding(v string)`

SetSamlSingleSignOnServiceBinding sets SamlSingleSignOnServiceBinding field to given value.

### HasSamlSingleSignOnServiceBinding

`func (o *SamlConfigModel) HasSamlSingleSignOnServiceBinding() bool`

HasSamlSingleSignOnServiceBinding returns a boolean if a field has been set.

### SetSamlSingleSignOnServiceBindingNil

`func (o *SamlConfigModel) SetSamlSingleSignOnServiceBindingNil(b bool)`

 SetSamlSingleSignOnServiceBindingNil sets the value for SamlSingleSignOnServiceBinding to be an explicit nil

### UnsetSamlSingleSignOnServiceBinding
`func (o *SamlConfigModel) UnsetSamlSingleSignOnServiceBinding()`

UnsetSamlSingleSignOnServiceBinding ensures that no value is present for SamlSingleSignOnServiceBinding, not even an explicit nil
### GetSamlResponse

`func (o *SamlConfigModel) GetSamlResponse() string`

GetSamlResponse returns the SamlResponse field if non-nil, zero value otherwise.

### GetSamlResponseOk

`func (o *SamlConfigModel) GetSamlResponseOk() (*string, bool)`

GetSamlResponseOk returns a tuple with the SamlResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlResponse

`func (o *SamlConfigModel) SetSamlResponse(v string)`

SetSamlResponse sets SamlResponse field to given value.

### HasSamlResponse

`func (o *SamlConfigModel) HasSamlResponse() bool`

HasSamlResponse returns a boolean if a field has been set.

### SetSamlResponseNil

`func (o *SamlConfigModel) SetSamlResponseNil(b bool)`

 SetSamlResponseNil sets the value for SamlResponse to be an explicit nil

### UnsetSamlResponse
`func (o *SamlConfigModel) UnsetSamlResponse()`

UnsetSamlResponse ensures that no value is present for SamlResponse, not even an explicit nil
### GetSamlAuthenticationContext

`func (o *SamlConfigModel) GetSamlAuthenticationContext() string`

GetSamlAuthenticationContext returns the SamlAuthenticationContext field if non-nil, zero value otherwise.

### GetSamlAuthenticationContextOk

`func (o *SamlConfigModel) GetSamlAuthenticationContextOk() (*string, bool)`

GetSamlAuthenticationContextOk returns a tuple with the SamlAuthenticationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAuthenticationContext

`func (o *SamlConfigModel) SetSamlAuthenticationContext(v string)`

SetSamlAuthenticationContext sets SamlAuthenticationContext field to given value.

### HasSamlAuthenticationContext

`func (o *SamlConfigModel) HasSamlAuthenticationContext() bool`

HasSamlAuthenticationContext returns a boolean if a field has been set.

### SetSamlAuthenticationContextNil

`func (o *SamlConfigModel) SetSamlAuthenticationContextNil(b bool)`

 SetSamlAuthenticationContextNil sets the value for SamlAuthenticationContext to be an explicit nil

### UnsetSamlAuthenticationContext
`func (o *SamlConfigModel) UnsetSamlAuthenticationContext()`

UnsetSamlAuthenticationContext ensures that no value is present for SamlAuthenticationContext, not even an explicit nil
### GetSamlAuthenticationContextComparison

`func (o *SamlConfigModel) GetSamlAuthenticationContextComparison() string`

GetSamlAuthenticationContextComparison returns the SamlAuthenticationContextComparison field if non-nil, zero value otherwise.

### GetSamlAuthenticationContextComparisonOk

`func (o *SamlConfigModel) GetSamlAuthenticationContextComparisonOk() (*string, bool)`

GetSamlAuthenticationContextComparisonOk returns a tuple with the SamlAuthenticationContextComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAuthenticationContextComparison

`func (o *SamlConfigModel) SetSamlAuthenticationContextComparison(v string)`

SetSamlAuthenticationContextComparison sets SamlAuthenticationContextComparison field to given value.

### HasSamlAuthenticationContextComparison

`func (o *SamlConfigModel) HasSamlAuthenticationContextComparison() bool`

HasSamlAuthenticationContextComparison returns a boolean if a field has been set.

### SetSamlAuthenticationContextComparisonNil

`func (o *SamlConfigModel) SetSamlAuthenticationContextComparisonNil(b bool)`

 SetSamlAuthenticationContextComparisonNil sets the value for SamlAuthenticationContextComparison to be an explicit nil

### UnsetSamlAuthenticationContextComparison
`func (o *SamlConfigModel) UnsetSamlAuthenticationContextComparison()`

UnsetSamlAuthenticationContextComparison ensures that no value is present for SamlAuthenticationContextComparison, not even an explicit nil
### GetSamlLogoutUrl

`func (o *SamlConfigModel) GetSamlLogoutUrl() string`

GetSamlLogoutUrl returns the SamlLogoutUrl field if non-nil, zero value otherwise.

### GetSamlLogoutUrlOk

`func (o *SamlConfigModel) GetSamlLogoutUrlOk() (*string, bool)`

GetSamlLogoutUrlOk returns a tuple with the SamlLogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLogoutUrl

`func (o *SamlConfigModel) SetSamlLogoutUrl(v string)`

SetSamlLogoutUrl sets SamlLogoutUrl field to given value.

### HasSamlLogoutUrl

`func (o *SamlConfigModel) HasSamlLogoutUrl() bool`

HasSamlLogoutUrl returns a boolean if a field has been set.

### SetSamlLogoutUrlNil

`func (o *SamlConfigModel) SetSamlLogoutUrlNil(b bool)`

 SetSamlLogoutUrlNil sets the value for SamlLogoutUrl to be an explicit nil

### UnsetSamlLogoutUrl
`func (o *SamlConfigModel) UnsetSamlLogoutUrl()`

UnsetSamlLogoutUrl ensures that no value is present for SamlLogoutUrl, not even an explicit nil
### GetSamlLogoutRequestBinding

`func (o *SamlConfigModel) GetSamlLogoutRequestBinding() string`

GetSamlLogoutRequestBinding returns the SamlLogoutRequestBinding field if non-nil, zero value otherwise.

### GetSamlLogoutRequestBindingOk

`func (o *SamlConfigModel) GetSamlLogoutRequestBindingOk() (*string, bool)`

GetSamlLogoutRequestBindingOk returns a tuple with the SamlLogoutRequestBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLogoutRequestBinding

`func (o *SamlConfigModel) SetSamlLogoutRequestBinding(v string)`

SetSamlLogoutRequestBinding sets SamlLogoutRequestBinding field to given value.

### HasSamlLogoutRequestBinding

`func (o *SamlConfigModel) HasSamlLogoutRequestBinding() bool`

HasSamlLogoutRequestBinding returns a boolean if a field has been set.

### SetSamlLogoutRequestBindingNil

`func (o *SamlConfigModel) SetSamlLogoutRequestBindingNil(b bool)`

 SetSamlLogoutRequestBindingNil sets the value for SamlLogoutRequestBinding to be an explicit nil

### UnsetSamlLogoutRequestBinding
`func (o *SamlConfigModel) UnsetSamlLogoutRequestBinding()`

UnsetSamlLogoutRequestBinding ensures that no value is present for SamlLogoutRequestBinding, not even an explicit nil
### GetSamlSignLogoutRequest

`func (o *SamlConfigModel) GetSamlSignLogoutRequest() string`

GetSamlSignLogoutRequest returns the SamlSignLogoutRequest field if non-nil, zero value otherwise.

### GetSamlSignLogoutRequestOk

`func (o *SamlConfigModel) GetSamlSignLogoutRequestOk() (*string, bool)`

GetSamlSignLogoutRequestOk returns a tuple with the SamlSignLogoutRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSignLogoutRequest

`func (o *SamlConfigModel) SetSamlSignLogoutRequest(v string)`

SetSamlSignLogoutRequest sets SamlSignLogoutRequest field to given value.

### HasSamlSignLogoutRequest

`func (o *SamlConfigModel) HasSamlSignLogoutRequest() bool`

HasSamlSignLogoutRequest returns a boolean if a field has been set.

### GetSamlAttributeNameForSid

`func (o *SamlConfigModel) GetSamlAttributeNameForSid() string`

GetSamlAttributeNameForSid returns the SamlAttributeNameForSid field if non-nil, zero value otherwise.

### GetSamlAttributeNameForSidOk

`func (o *SamlConfigModel) GetSamlAttributeNameForSidOk() (*string, bool)`

GetSamlAttributeNameForSidOk returns a tuple with the SamlAttributeNameForSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForSid

`func (o *SamlConfigModel) SetSamlAttributeNameForSid(v string)`

SetSamlAttributeNameForSid sets SamlAttributeNameForSid field to given value.

### HasSamlAttributeNameForSid

`func (o *SamlConfigModel) HasSamlAttributeNameForSid() bool`

HasSamlAttributeNameForSid returns a boolean if a field has been set.

### SetSamlAttributeNameForSidNil

`func (o *SamlConfigModel) SetSamlAttributeNameForSidNil(b bool)`

 SetSamlAttributeNameForSidNil sets the value for SamlAttributeNameForSid to be an explicit nil

### UnsetSamlAttributeNameForSid
`func (o *SamlConfigModel) UnsetSamlAttributeNameForSid()`

UnsetSamlAttributeNameForSid ensures that no value is present for SamlAttributeNameForSid, not even an explicit nil
### GetSamlAttributeNameForUpn

`func (o *SamlConfigModel) GetSamlAttributeNameForUpn() string`

GetSamlAttributeNameForUpn returns the SamlAttributeNameForUpn field if non-nil, zero value otherwise.

### GetSamlAttributeNameForUpnOk

`func (o *SamlConfigModel) GetSamlAttributeNameForUpnOk() (*string, bool)`

GetSamlAttributeNameForUpnOk returns a tuple with the SamlAttributeNameForUpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForUpn

`func (o *SamlConfigModel) SetSamlAttributeNameForUpn(v string)`

SetSamlAttributeNameForUpn sets SamlAttributeNameForUpn field to given value.

### HasSamlAttributeNameForUpn

`func (o *SamlConfigModel) HasSamlAttributeNameForUpn() bool`

HasSamlAttributeNameForUpn returns a boolean if a field has been set.

### SetSamlAttributeNameForUpnNil

`func (o *SamlConfigModel) SetSamlAttributeNameForUpnNil(b bool)`

 SetSamlAttributeNameForUpnNil sets the value for SamlAttributeNameForUpn to be an explicit nil

### UnsetSamlAttributeNameForUpn
`func (o *SamlConfigModel) UnsetSamlAttributeNameForUpn()`

UnsetSamlAttributeNameForUpn ensures that no value is present for SamlAttributeNameForUpn, not even an explicit nil
### GetSamlAttributeNameForEmail

`func (o *SamlConfigModel) GetSamlAttributeNameForEmail() string`

GetSamlAttributeNameForEmail returns the SamlAttributeNameForEmail field if non-nil, zero value otherwise.

### GetSamlAttributeNameForEmailOk

`func (o *SamlConfigModel) GetSamlAttributeNameForEmailOk() (*string, bool)`

GetSamlAttributeNameForEmailOk returns a tuple with the SamlAttributeNameForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForEmail

`func (o *SamlConfigModel) SetSamlAttributeNameForEmail(v string)`

SetSamlAttributeNameForEmail sets SamlAttributeNameForEmail field to given value.

### HasSamlAttributeNameForEmail

`func (o *SamlConfigModel) HasSamlAttributeNameForEmail() bool`

HasSamlAttributeNameForEmail returns a boolean if a field has been set.

### SetSamlAttributeNameForEmailNil

`func (o *SamlConfigModel) SetSamlAttributeNameForEmailNil(b bool)`

 SetSamlAttributeNameForEmailNil sets the value for SamlAttributeNameForEmail to be an explicit nil

### UnsetSamlAttributeNameForEmail
`func (o *SamlConfigModel) UnsetSamlAttributeNameForEmail()`

UnsetSamlAttributeNameForEmail ensures that no value is present for SamlAttributeNameForEmail, not even an explicit nil
### GetSamlAttributeNameForAdOid

`func (o *SamlConfigModel) GetSamlAttributeNameForAdOid() string`

GetSamlAttributeNameForAdOid returns the SamlAttributeNameForAdOid field if non-nil, zero value otherwise.

### GetSamlAttributeNameForAdOidOk

`func (o *SamlConfigModel) GetSamlAttributeNameForAdOidOk() (*string, bool)`

GetSamlAttributeNameForAdOidOk returns a tuple with the SamlAttributeNameForAdOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForAdOid

`func (o *SamlConfigModel) SetSamlAttributeNameForAdOid(v string)`

SetSamlAttributeNameForAdOid sets SamlAttributeNameForAdOid field to given value.

### HasSamlAttributeNameForAdOid

`func (o *SamlConfigModel) HasSamlAttributeNameForAdOid() bool`

HasSamlAttributeNameForAdOid returns a boolean if a field has been set.

### SetSamlAttributeNameForAdOidNil

`func (o *SamlConfigModel) SetSamlAttributeNameForAdOidNil(b bool)`

 SetSamlAttributeNameForAdOidNil sets the value for SamlAttributeNameForAdOid to be an explicit nil

### UnsetSamlAttributeNameForAdOid
`func (o *SamlConfigModel) UnsetSamlAttributeNameForAdOid()`

UnsetSamlAttributeNameForAdOid ensures that no value is present for SamlAttributeNameForAdOid, not even an explicit nil
### GetSamlAttributeNameForAdForest

`func (o *SamlConfigModel) GetSamlAttributeNameForAdForest() string`

GetSamlAttributeNameForAdForest returns the SamlAttributeNameForAdForest field if non-nil, zero value otherwise.

### GetSamlAttributeNameForAdForestOk

`func (o *SamlConfigModel) GetSamlAttributeNameForAdForestOk() (*string, bool)`

GetSamlAttributeNameForAdForestOk returns a tuple with the SamlAttributeNameForAdForest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForAdForest

`func (o *SamlConfigModel) SetSamlAttributeNameForAdForest(v string)`

SetSamlAttributeNameForAdForest sets SamlAttributeNameForAdForest field to given value.

### HasSamlAttributeNameForAdForest

`func (o *SamlConfigModel) HasSamlAttributeNameForAdForest() bool`

HasSamlAttributeNameForAdForest returns a boolean if a field has been set.

### SetSamlAttributeNameForAdForestNil

`func (o *SamlConfigModel) SetSamlAttributeNameForAdForestNil(b bool)`

 SetSamlAttributeNameForAdForestNil sets the value for SamlAttributeNameForAdForest to be an explicit nil

### UnsetSamlAttributeNameForAdForest
`func (o *SamlConfigModel) UnsetSamlAttributeNameForAdForest()`

UnsetSamlAttributeNameForAdForest ensures that no value is present for SamlAttributeNameForAdForest, not even an explicit nil
### GetSamlAttributeNameForAdDomain

`func (o *SamlConfigModel) GetSamlAttributeNameForAdDomain() string`

GetSamlAttributeNameForAdDomain returns the SamlAttributeNameForAdDomain field if non-nil, zero value otherwise.

### GetSamlAttributeNameForAdDomainOk

`func (o *SamlConfigModel) GetSamlAttributeNameForAdDomainOk() (*string, bool)`

GetSamlAttributeNameForAdDomainOk returns a tuple with the SamlAttributeNameForAdDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForAdDomain

`func (o *SamlConfigModel) SetSamlAttributeNameForAdDomain(v string)`

SetSamlAttributeNameForAdDomain sets SamlAttributeNameForAdDomain field to given value.

### HasSamlAttributeNameForAdDomain

`func (o *SamlConfigModel) HasSamlAttributeNameForAdDomain() bool`

HasSamlAttributeNameForAdDomain returns a boolean if a field has been set.

### SetSamlAttributeNameForAdDomainNil

`func (o *SamlConfigModel) SetSamlAttributeNameForAdDomainNil(b bool)`

 SetSamlAttributeNameForAdDomainNil sets the value for SamlAttributeNameForAdDomain to be an explicit nil

### UnsetSamlAttributeNameForAdDomain
`func (o *SamlConfigModel) UnsetSamlAttributeNameForAdDomain()`

UnsetSamlAttributeNameForAdDomain ensures that no value is present for SamlAttributeNameForAdDomain, not even an explicit nil
### GetSamlAttributeNameForUserDisplayName

`func (o *SamlConfigModel) GetSamlAttributeNameForUserDisplayName() string`

GetSamlAttributeNameForUserDisplayName returns the SamlAttributeNameForUserDisplayName field if non-nil, zero value otherwise.

### GetSamlAttributeNameForUserDisplayNameOk

`func (o *SamlConfigModel) GetSamlAttributeNameForUserDisplayNameOk() (*string, bool)`

GetSamlAttributeNameForUserDisplayNameOk returns a tuple with the SamlAttributeNameForUserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForUserDisplayName

`func (o *SamlConfigModel) SetSamlAttributeNameForUserDisplayName(v string)`

SetSamlAttributeNameForUserDisplayName sets SamlAttributeNameForUserDisplayName field to given value.

### HasSamlAttributeNameForUserDisplayName

`func (o *SamlConfigModel) HasSamlAttributeNameForUserDisplayName() bool`

HasSamlAttributeNameForUserDisplayName returns a boolean if a field has been set.

### SetSamlAttributeNameForUserDisplayNameNil

`func (o *SamlConfigModel) SetSamlAttributeNameForUserDisplayNameNil(b bool)`

 SetSamlAttributeNameForUserDisplayNameNil sets the value for SamlAttributeNameForUserDisplayName to be an explicit nil

### UnsetSamlAttributeNameForUserDisplayName
`func (o *SamlConfigModel) UnsetSamlAttributeNameForUserDisplayName()`

UnsetSamlAttributeNameForUserDisplayName ensures that no value is present for SamlAttributeNameForUserDisplayName, not even an explicit nil
### GetSamlAttributeNameForUserGivenName

`func (o *SamlConfigModel) GetSamlAttributeNameForUserGivenName() string`

GetSamlAttributeNameForUserGivenName returns the SamlAttributeNameForUserGivenName field if non-nil, zero value otherwise.

### GetSamlAttributeNameForUserGivenNameOk

`func (o *SamlConfigModel) GetSamlAttributeNameForUserGivenNameOk() (*string, bool)`

GetSamlAttributeNameForUserGivenNameOk returns a tuple with the SamlAttributeNameForUserGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForUserGivenName

`func (o *SamlConfigModel) SetSamlAttributeNameForUserGivenName(v string)`

SetSamlAttributeNameForUserGivenName sets SamlAttributeNameForUserGivenName field to given value.

### HasSamlAttributeNameForUserGivenName

`func (o *SamlConfigModel) HasSamlAttributeNameForUserGivenName() bool`

HasSamlAttributeNameForUserGivenName returns a boolean if a field has been set.

### SetSamlAttributeNameForUserGivenNameNil

`func (o *SamlConfigModel) SetSamlAttributeNameForUserGivenNameNil(b bool)`

 SetSamlAttributeNameForUserGivenNameNil sets the value for SamlAttributeNameForUserGivenName to be an explicit nil

### UnsetSamlAttributeNameForUserGivenName
`func (o *SamlConfigModel) UnsetSamlAttributeNameForUserGivenName()`

UnsetSamlAttributeNameForUserGivenName ensures that no value is present for SamlAttributeNameForUserGivenName, not even an explicit nil
### GetSamlAttributeNameForUserFamilyName

`func (o *SamlConfigModel) GetSamlAttributeNameForUserFamilyName() string`

GetSamlAttributeNameForUserFamilyName returns the SamlAttributeNameForUserFamilyName field if non-nil, zero value otherwise.

### GetSamlAttributeNameForUserFamilyNameOk

`func (o *SamlConfigModel) GetSamlAttributeNameForUserFamilyNameOk() (*string, bool)`

GetSamlAttributeNameForUserFamilyNameOk returns a tuple with the SamlAttributeNameForUserFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAttributeNameForUserFamilyName

`func (o *SamlConfigModel) SetSamlAttributeNameForUserFamilyName(v string)`

SetSamlAttributeNameForUserFamilyName sets SamlAttributeNameForUserFamilyName field to given value.

### HasSamlAttributeNameForUserFamilyName

`func (o *SamlConfigModel) HasSamlAttributeNameForUserFamilyName() bool`

HasSamlAttributeNameForUserFamilyName returns a boolean if a field has been set.

### SetSamlAttributeNameForUserFamilyNameNil

`func (o *SamlConfigModel) SetSamlAttributeNameForUserFamilyNameNil(b bool)`

 SetSamlAttributeNameForUserFamilyNameNil sets the value for SamlAttributeNameForUserFamilyName to be an explicit nil

### UnsetSamlAttributeNameForUserFamilyName
`func (o *SamlConfigModel) UnsetSamlAttributeNameForUserFamilyName()`

UnsetSamlAttributeNameForUserFamilyName ensures that no value is present for SamlAttributeNameForUserFamilyName, not even an explicit nil
### GetSamlCertPublicKey

`func (o *SamlConfigModel) GetSamlCertPublicKey() string`

GetSamlCertPublicKey returns the SamlCertPublicKey field if non-nil, zero value otherwise.

### GetSamlCertPublicKeyOk

`func (o *SamlConfigModel) GetSamlCertPublicKeyOk() (*string, bool)`

GetSamlCertPublicKeyOk returns a tuple with the SamlCertPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertPublicKey

`func (o *SamlConfigModel) SetSamlCertPublicKey(v string)`

SetSamlCertPublicKey sets SamlCertPublicKey field to given value.

### HasSamlCertPublicKey

`func (o *SamlConfigModel) HasSamlCertPublicKey() bool`

HasSamlCertPublicKey returns a boolean if a field has been set.

### SetSamlCertPublicKeyNil

`func (o *SamlConfigModel) SetSamlCertPublicKeyNil(b bool)`

 SetSamlCertPublicKeyNil sets the value for SamlCertPublicKey to be an explicit nil

### UnsetSamlCertPublicKey
`func (o *SamlConfigModel) UnsetSamlCertPublicKey()`

UnsetSamlCertPublicKey ensures that no value is present for SamlCertPublicKey, not even an explicit nil
### GetSamlCertCN

`func (o *SamlConfigModel) GetSamlCertCN() string`

GetSamlCertCN returns the SamlCertCN field if non-nil, zero value otherwise.

### GetSamlCertCNOk

`func (o *SamlConfigModel) GetSamlCertCNOk() (*string, bool)`

GetSamlCertCNOk returns a tuple with the SamlCertCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertCN

`func (o *SamlConfigModel) SetSamlCertCN(v string)`

SetSamlCertCN sets SamlCertCN field to given value.

### HasSamlCertCN

`func (o *SamlConfigModel) HasSamlCertCN() bool`

HasSamlCertCN returns a boolean if a field has been set.

### SetSamlCertCNNil

`func (o *SamlConfigModel) SetSamlCertCNNil(b bool)`

 SetSamlCertCNNil sets the value for SamlCertCN to be an explicit nil

### UnsetSamlCertCN
`func (o *SamlConfigModel) UnsetSamlCertCN()`

UnsetSamlCertCN ensures that no value is present for SamlCertCN, not even an explicit nil
### GetSamlCertFileName

`func (o *SamlConfigModel) GetSamlCertFileName() string`

GetSamlCertFileName returns the SamlCertFileName field if non-nil, zero value otherwise.

### GetSamlCertFileNameOk

`func (o *SamlConfigModel) GetSamlCertFileNameOk() (*string, bool)`

GetSamlCertFileNameOk returns a tuple with the SamlCertFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertFileName

`func (o *SamlConfigModel) SetSamlCertFileName(v string)`

SetSamlCertFileName sets SamlCertFileName field to given value.

### HasSamlCertFileName

`func (o *SamlConfigModel) HasSamlCertFileName() bool`

HasSamlCertFileName returns a boolean if a field has been set.

### SetSamlCertFileNameNil

`func (o *SamlConfigModel) SetSamlCertFileNameNil(b bool)`

 SetSamlCertFileNameNil sets the value for SamlCertFileName to be an explicit nil

### UnsetSamlCertFileName
`func (o *SamlConfigModel) UnsetSamlCertFileName()`

UnsetSamlCertFileName ensures that no value is present for SamlCertFileName, not even an explicit nil
### GetSamlCertExpiration

`func (o *SamlConfigModel) GetSamlCertExpiration() string`

GetSamlCertExpiration returns the SamlCertExpiration field if non-nil, zero value otherwise.

### GetSamlCertExpirationOk

`func (o *SamlConfigModel) GetSamlCertExpirationOk() (*string, bool)`

GetSamlCertExpirationOk returns a tuple with the SamlCertExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertExpiration

`func (o *SamlConfigModel) SetSamlCertExpiration(v string)`

SetSamlCertExpiration sets SamlCertExpiration field to given value.

### HasSamlCertExpiration

`func (o *SamlConfigModel) HasSamlCertExpiration() bool`

HasSamlCertExpiration returns a boolean if a field has been set.

### SetSamlCertExpirationNil

`func (o *SamlConfigModel) SetSamlCertExpirationNil(b bool)`

 SetSamlCertExpirationNil sets the value for SamlCertExpiration to be an explicit nil

### UnsetSamlCertExpiration
`func (o *SamlConfigModel) UnsetSamlCertExpiration()`

UnsetSamlCertExpiration ensures that no value is present for SamlCertExpiration, not even an explicit nil
### GetSamlCertData

`func (o *SamlConfigModel) GetSamlCertData() string`

GetSamlCertData returns the SamlCertData field if non-nil, zero value otherwise.

### GetSamlCertDataOk

`func (o *SamlConfigModel) GetSamlCertDataOk() (*string, bool)`

GetSamlCertDataOk returns a tuple with the SamlCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertData

`func (o *SamlConfigModel) SetSamlCertData(v string)`

SetSamlCertData sets SamlCertData field to given value.

### HasSamlCertData

`func (o *SamlConfigModel) HasSamlCertData() bool`

HasSamlCertData returns a boolean if a field has been set.

### SetSamlCertDataNil

`func (o *SamlConfigModel) SetSamlCertDataNil(b bool)`

 SetSamlCertDataNil sets the value for SamlCertData to be an explicit nil

### UnsetSamlCertData
`func (o *SamlConfigModel) UnsetSamlCertData()`

UnsetSamlCertData ensures that no value is present for SamlCertData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


