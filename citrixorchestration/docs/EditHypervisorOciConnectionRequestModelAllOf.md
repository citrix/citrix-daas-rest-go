# EditHypervisorOciConnectionRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountId** | **string** | The service account ID used to access the Oracle Cloud Infrastructure APIs. Required. | 
**ServiceAccountCredential** | **string** | The private key string to access the Oracle Cloud Infrastructure APIs. Required. Must be specified in the format indicated by ServiceAccountCredentialFormat. | 
**ServiceAccountCredentialFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**ServiceAccountFingerprint** | **string** | The fingerprint of the public key associate with the ServiceAccountCredential. | 

## Methods

### NewEditHypervisorOciConnectionRequestModelAllOf

`func NewEditHypervisorOciConnectionRequestModelAllOf(serviceAccountId string, serviceAccountCredential string, serviceAccountFingerprint string, ) *EditHypervisorOciConnectionRequestModelAllOf`

NewEditHypervisorOciConnectionRequestModelAllOf instantiates a new EditHypervisorOciConnectionRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHypervisorOciConnectionRequestModelAllOfWithDefaults

`func NewEditHypervisorOciConnectionRequestModelAllOfWithDefaults() *EditHypervisorOciConnectionRequestModelAllOf`

NewEditHypervisorOciConnectionRequestModelAllOfWithDefaults instantiates a new EditHypervisorOciConnectionRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccountId

`func (o *EditHypervisorOciConnectionRequestModelAllOf) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *EditHypervisorOciConnectionRequestModelAllOf) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *EditHypervisorOciConnectionRequestModelAllOf) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetServiceAccountCredential

`func (o *EditHypervisorOciConnectionRequestModelAllOf) GetServiceAccountCredential() string`

GetServiceAccountCredential returns the ServiceAccountCredential field if non-nil, zero value otherwise.

### GetServiceAccountCredentialOk

`func (o *EditHypervisorOciConnectionRequestModelAllOf) GetServiceAccountCredentialOk() (*string, bool)`

GetServiceAccountCredentialOk returns a tuple with the ServiceAccountCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredential

`func (o *EditHypervisorOciConnectionRequestModelAllOf) SetServiceAccountCredential(v string)`

SetServiceAccountCredential sets ServiceAccountCredential field to given value.


### GetServiceAccountCredentialFormat

`func (o *EditHypervisorOciConnectionRequestModelAllOf) GetServiceAccountCredentialFormat() IdentityPasswordFormat`

GetServiceAccountCredentialFormat returns the ServiceAccountCredentialFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialFormatOk

`func (o *EditHypervisorOciConnectionRequestModelAllOf) GetServiceAccountCredentialFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialFormatOk returns a tuple with the ServiceAccountCredentialFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialFormat

`func (o *EditHypervisorOciConnectionRequestModelAllOf) SetServiceAccountCredentialFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialFormat sets ServiceAccountCredentialFormat field to given value.

### HasServiceAccountCredentialFormat

`func (o *EditHypervisorOciConnectionRequestModelAllOf) HasServiceAccountCredentialFormat() bool`

HasServiceAccountCredentialFormat returns a boolean if a field has been set.

### GetServiceAccountFingerprint

`func (o *EditHypervisorOciConnectionRequestModelAllOf) GetServiceAccountFingerprint() string`

GetServiceAccountFingerprint returns the ServiceAccountFingerprint field if non-nil, zero value otherwise.

### GetServiceAccountFingerprintOk

`func (o *EditHypervisorOciConnectionRequestModelAllOf) GetServiceAccountFingerprintOk() (*string, bool)`

GetServiceAccountFingerprintOk returns a tuple with the ServiceAccountFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountFingerprint

`func (o *EditHypervisorOciConnectionRequestModelAllOf) SetServiceAccountFingerprint(v string)`

SetServiceAccountFingerprint sets ServiceAccountFingerprint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


