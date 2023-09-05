# HypervisorOciConnectionDetailRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenancyOcid** | **string** | Oracle Cloud Infrastructure tenancy to connect to. Required. | 
**Region** | **string** | Oracle Cloud Infrastructure region to connect to. Required. | 
**ServiceAccountId** | **string** | The service account ID used to access the Oracle Cloud Infrastructure APIs. Required. | 
**ServiceAccountCredential** | **string** | The private key string to access the Oracle Cloud Infrastructure APIs. Required. Must be specified in the format indicated by ServiceAccountCredentialFormat. | 
**ServiceAccountCredentialFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**ServiceAccountFingerprint** | **string** | The fingerprint of the public key associate with the ServiceAccountCredential. | 
**OciEnvironment** | Pointer to [**OciEnvironment**](OciEnvironment.md) |  | [optional] 

## Methods

### NewHypervisorOciConnectionDetailRequestModelAllOf

`func NewHypervisorOciConnectionDetailRequestModelAllOf(tenancyOcid string, region string, serviceAccountId string, serviceAccountCredential string, serviceAccountFingerprint string, ) *HypervisorOciConnectionDetailRequestModelAllOf`

NewHypervisorOciConnectionDetailRequestModelAllOf instantiates a new HypervisorOciConnectionDetailRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorOciConnectionDetailRequestModelAllOfWithDefaults

`func NewHypervisorOciConnectionDetailRequestModelAllOfWithDefaults() *HypervisorOciConnectionDetailRequestModelAllOf`

NewHypervisorOciConnectionDetailRequestModelAllOfWithDefaults instantiates a new HypervisorOciConnectionDetailRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenancyOcid

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetTenancyOcid() string`

GetTenancyOcid returns the TenancyOcid field if non-nil, zero value otherwise.

### GetTenancyOcidOk

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetTenancyOcidOk() (*string, bool)`

GetTenancyOcidOk returns a tuple with the TenancyOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyOcid

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) SetTenancyOcid(v string)`

SetTenancyOcid sets TenancyOcid field to given value.


### GetRegion

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetServiceAccountId

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetServiceAccountCredential

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetServiceAccountCredential() string`

GetServiceAccountCredential returns the ServiceAccountCredential field if non-nil, zero value otherwise.

### GetServiceAccountCredentialOk

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetServiceAccountCredentialOk() (*string, bool)`

GetServiceAccountCredentialOk returns a tuple with the ServiceAccountCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredential

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) SetServiceAccountCredential(v string)`

SetServiceAccountCredential sets ServiceAccountCredential field to given value.


### GetServiceAccountCredentialFormat

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetServiceAccountCredentialFormat() IdentityPasswordFormat`

GetServiceAccountCredentialFormat returns the ServiceAccountCredentialFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialFormatOk

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetServiceAccountCredentialFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialFormatOk returns a tuple with the ServiceAccountCredentialFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialFormat

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) SetServiceAccountCredentialFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialFormat sets ServiceAccountCredentialFormat field to given value.

### HasServiceAccountCredentialFormat

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) HasServiceAccountCredentialFormat() bool`

HasServiceAccountCredentialFormat returns a boolean if a field has been set.

### GetServiceAccountFingerprint

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetServiceAccountFingerprint() string`

GetServiceAccountFingerprint returns the ServiceAccountFingerprint field if non-nil, zero value otherwise.

### GetServiceAccountFingerprintOk

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetServiceAccountFingerprintOk() (*string, bool)`

GetServiceAccountFingerprintOk returns a tuple with the ServiceAccountFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountFingerprint

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) SetServiceAccountFingerprint(v string)`

SetServiceAccountFingerprint sets ServiceAccountFingerprint field to given value.


### GetOciEnvironment

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetOciEnvironment() OciEnvironment`

GetOciEnvironment returns the OciEnvironment field if non-nil, zero value otherwise.

### GetOciEnvironmentOk

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) GetOciEnvironmentOk() (*OciEnvironment, bool)`

GetOciEnvironmentOk returns a tuple with the OciEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciEnvironment

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) SetOciEnvironment(v OciEnvironment)`

SetOciEnvironment sets OciEnvironment field to given value.

### HasOciEnvironment

`func (o *HypervisorOciConnectionDetailRequestModelAllOf) HasOciEnvironment() bool`

HasOciEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


