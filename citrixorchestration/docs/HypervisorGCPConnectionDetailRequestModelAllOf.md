# HypervisorGCPConnectionDetailRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountId** | **string** | The service account ID used to access the Google Cloud APIs. Required. | 
**ServiceAccountCredentials** | **string** | the JSON-encoded service account credentials used to access the Google Cloud APIs.  Required. Must be specified in the format indicated by ServiceAccountCredentialsFormat. | 
**ServiceAccountCredentialsFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 

## Methods

### NewHypervisorGCPConnectionDetailRequestModelAllOf

`func NewHypervisorGCPConnectionDetailRequestModelAllOf(serviceAccountId string, serviceAccountCredentials string, ) *HypervisorGCPConnectionDetailRequestModelAllOf`

NewHypervisorGCPConnectionDetailRequestModelAllOf instantiates a new HypervisorGCPConnectionDetailRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorGCPConnectionDetailRequestModelAllOfWithDefaults

`func NewHypervisorGCPConnectionDetailRequestModelAllOfWithDefaults() *HypervisorGCPConnectionDetailRequestModelAllOf`

NewHypervisorGCPConnectionDetailRequestModelAllOfWithDefaults instantiates a new HypervisorGCPConnectionDetailRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccountId

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetServiceAccountCredentials

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) GetServiceAccountCredentials() string`

GetServiceAccountCredentials returns the ServiceAccountCredentials field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsOk

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) GetServiceAccountCredentialsOk() (*string, bool)`

GetServiceAccountCredentialsOk returns a tuple with the ServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentials

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) SetServiceAccountCredentials(v string)`

SetServiceAccountCredentials sets ServiceAccountCredentials field to given value.


### GetServiceAccountCredentialsFormat

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) GetServiceAccountCredentialsFormat() IdentityPasswordFormat`

GetServiceAccountCredentialsFormat returns the ServiceAccountCredentialsFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsFormatOk

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) GetServiceAccountCredentialsFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialsFormatOk returns a tuple with the ServiceAccountCredentialsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialsFormat

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) SetServiceAccountCredentialsFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialsFormat sets ServiceAccountCredentialsFormat field to given value.

### HasServiceAccountCredentialsFormat

`func (o *HypervisorGCPConnectionDetailRequestModelAllOf) HasServiceAccountCredentialsFormat() bool`

HasServiceAccountCredentialsFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


