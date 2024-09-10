# SamlCertInfoModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotAfter** | **string** |  | 
**PublicKey** | **string** |  | 
**CommonName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSamlCertInfoModel

`func NewSamlCertInfoModel(notAfter string, publicKey string, ) *SamlCertInfoModel`

NewSamlCertInfoModel instantiates a new SamlCertInfoModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlCertInfoModelWithDefaults

`func NewSamlCertInfoModelWithDefaults() *SamlCertInfoModel`

NewSamlCertInfoModelWithDefaults instantiates a new SamlCertInfoModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotAfter

`func (o *SamlCertInfoModel) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *SamlCertInfoModel) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *SamlCertInfoModel) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.


### GetPublicKey

`func (o *SamlCertInfoModel) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SamlCertInfoModel) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SamlCertInfoModel) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetCommonName

`func (o *SamlCertInfoModel) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *SamlCertInfoModel) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *SamlCertInfoModel) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *SamlCertInfoModel) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *SamlCertInfoModel) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *SamlCertInfoModel) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


