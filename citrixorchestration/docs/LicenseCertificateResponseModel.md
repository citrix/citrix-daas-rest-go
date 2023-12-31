# LicenseCertificateResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Thumbprint** | Pointer to **NullableString** | The thumbprint of the certificate | [optional] 
**Issuer** | Pointer to **NullableString** | The issuer of the certificate | [optional] 
**Subject** | Pointer to **NullableString** | The subject of the certificate | [optional] 
**ValidFrom** | Pointer to **NullableString** | Datetime the certificate valid from | [optional] 
**ValidTo** | Pointer to **NullableString** | Datetime the certificate valid to | [optional] 
**SerialNumber** | Pointer to **NullableString** | The serialNumber of the certificate | [optional] 
**FriendlyName** | Pointer to **NullableString** | The friendly name of the certificate | [optional] 
**Version** | Pointer to **NullableInt32** | The certificate version | [optional] 
**RawData** | Pointer to **NullableString** | The raw data for the entire X509v3 certificate as a hexadecimal string | [optional] 
**ChainStatusInformations** | Pointer to **[]string** | A description of the ChainStatus value | [optional] 
**ChainStatusFlags** | Pointer to [**[]X509ChainStatusFlags**](X509ChainStatusFlags.md) | The status of the X509 chain | [optional] 

## Methods

### NewLicenseCertificateResponseModel

`func NewLicenseCertificateResponseModel() *LicenseCertificateResponseModel`

NewLicenseCertificateResponseModel instantiates a new LicenseCertificateResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseCertificateResponseModelWithDefaults

`func NewLicenseCertificateResponseModelWithDefaults() *LicenseCertificateResponseModel`

NewLicenseCertificateResponseModelWithDefaults instantiates a new LicenseCertificateResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThumbprint

`func (o *LicenseCertificateResponseModel) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *LicenseCertificateResponseModel) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *LicenseCertificateResponseModel) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *LicenseCertificateResponseModel) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *LicenseCertificateResponseModel) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *LicenseCertificateResponseModel) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetIssuer

`func (o *LicenseCertificateResponseModel) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *LicenseCertificateResponseModel) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *LicenseCertificateResponseModel) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *LicenseCertificateResponseModel) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *LicenseCertificateResponseModel) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *LicenseCertificateResponseModel) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetSubject

`func (o *LicenseCertificateResponseModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LicenseCertificateResponseModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *LicenseCertificateResponseModel) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *LicenseCertificateResponseModel) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *LicenseCertificateResponseModel) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *LicenseCertificateResponseModel) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetValidFrom

`func (o *LicenseCertificateResponseModel) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *LicenseCertificateResponseModel) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *LicenseCertificateResponseModel) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *LicenseCertificateResponseModel) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### SetValidFromNil

`func (o *LicenseCertificateResponseModel) SetValidFromNil(b bool)`

 SetValidFromNil sets the value for ValidFrom to be an explicit nil

### UnsetValidFrom
`func (o *LicenseCertificateResponseModel) UnsetValidFrom()`

UnsetValidFrom ensures that no value is present for ValidFrom, not even an explicit nil
### GetValidTo

`func (o *LicenseCertificateResponseModel) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *LicenseCertificateResponseModel) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *LicenseCertificateResponseModel) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *LicenseCertificateResponseModel) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### SetValidToNil

`func (o *LicenseCertificateResponseModel) SetValidToNil(b bool)`

 SetValidToNil sets the value for ValidTo to be an explicit nil

### UnsetValidTo
`func (o *LicenseCertificateResponseModel) UnsetValidTo()`

UnsetValidTo ensures that no value is present for ValidTo, not even an explicit nil
### GetSerialNumber

`func (o *LicenseCertificateResponseModel) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *LicenseCertificateResponseModel) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *LicenseCertificateResponseModel) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *LicenseCertificateResponseModel) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *LicenseCertificateResponseModel) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *LicenseCertificateResponseModel) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetFriendlyName

`func (o *LicenseCertificateResponseModel) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *LicenseCertificateResponseModel) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *LicenseCertificateResponseModel) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *LicenseCertificateResponseModel) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *LicenseCertificateResponseModel) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *LicenseCertificateResponseModel) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetVersion

`func (o *LicenseCertificateResponseModel) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LicenseCertificateResponseModel) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LicenseCertificateResponseModel) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LicenseCertificateResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *LicenseCertificateResponseModel) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *LicenseCertificateResponseModel) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetRawData

`func (o *LicenseCertificateResponseModel) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *LicenseCertificateResponseModel) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *LicenseCertificateResponseModel) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *LicenseCertificateResponseModel) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *LicenseCertificateResponseModel) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *LicenseCertificateResponseModel) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil
### GetChainStatusInformations

`func (o *LicenseCertificateResponseModel) GetChainStatusInformations() []string`

GetChainStatusInformations returns the ChainStatusInformations field if non-nil, zero value otherwise.

### GetChainStatusInformationsOk

`func (o *LicenseCertificateResponseModel) GetChainStatusInformationsOk() (*[]string, bool)`

GetChainStatusInformationsOk returns a tuple with the ChainStatusInformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainStatusInformations

`func (o *LicenseCertificateResponseModel) SetChainStatusInformations(v []string)`

SetChainStatusInformations sets ChainStatusInformations field to given value.

### HasChainStatusInformations

`func (o *LicenseCertificateResponseModel) HasChainStatusInformations() bool`

HasChainStatusInformations returns a boolean if a field has been set.

### SetChainStatusInformationsNil

`func (o *LicenseCertificateResponseModel) SetChainStatusInformationsNil(b bool)`

 SetChainStatusInformationsNil sets the value for ChainStatusInformations to be an explicit nil

### UnsetChainStatusInformations
`func (o *LicenseCertificateResponseModel) UnsetChainStatusInformations()`

UnsetChainStatusInformations ensures that no value is present for ChainStatusInformations, not even an explicit nil
### GetChainStatusFlags

`func (o *LicenseCertificateResponseModel) GetChainStatusFlags() []X509ChainStatusFlags`

GetChainStatusFlags returns the ChainStatusFlags field if non-nil, zero value otherwise.

### GetChainStatusFlagsOk

`func (o *LicenseCertificateResponseModel) GetChainStatusFlagsOk() (*[]X509ChainStatusFlags, bool)`

GetChainStatusFlagsOk returns a tuple with the ChainStatusFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainStatusFlags

`func (o *LicenseCertificateResponseModel) SetChainStatusFlags(v []X509ChainStatusFlags)`

SetChainStatusFlags sets ChainStatusFlags field to given value.

### HasChainStatusFlags

`func (o *LicenseCertificateResponseModel) HasChainStatusFlags() bool`

HasChainStatusFlags returns a boolean if a field has been set.

### SetChainStatusFlagsNil

`func (o *LicenseCertificateResponseModel) SetChainStatusFlagsNil(b bool)`

 SetChainStatusFlagsNil sets the value for ChainStatusFlags to be an explicit nil

### UnsetChainStatusFlags
`func (o *LicenseCertificateResponseModel) UnsetChainStatusFlags()`

UnsetChainStatusFlags ensures that no value is present for ChainStatusFlags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


