# EditHypervisorGCPConnectionRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountId** | Pointer to **string** | The service account ID used to access the Google Cloud APIs. Optional.  If not specified, will not be changed.  If specified, ServiceAccountCredentials must also be specified. | [optional] 
**ServiceAccountCredentials** | Pointer to **string** | the JSON-encoded service account credentials used to access the Google Cloud APIs.  Optional.  If not specified, will not be changed.  If specified, must be in the format indicated by ServiceAccountCredentialsFormat. | [optional] 
**ServiceAccountCredentialsFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 

## Methods

### NewEditHypervisorGCPConnectionRequestModelAllOf

`func NewEditHypervisorGCPConnectionRequestModelAllOf() *EditHypervisorGCPConnectionRequestModelAllOf`

NewEditHypervisorGCPConnectionRequestModelAllOf instantiates a new EditHypervisorGCPConnectionRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHypervisorGCPConnectionRequestModelAllOfWithDefaults

`func NewEditHypervisorGCPConnectionRequestModelAllOfWithDefaults() *EditHypervisorGCPConnectionRequestModelAllOf`

NewEditHypervisorGCPConnectionRequestModelAllOfWithDefaults instantiates a new EditHypervisorGCPConnectionRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccountId

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetServiceAccountCredentials

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) GetServiceAccountCredentials() string`

GetServiceAccountCredentials returns the ServiceAccountCredentials field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsOk

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) GetServiceAccountCredentialsOk() (*string, bool)`

GetServiceAccountCredentialsOk returns a tuple with the ServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentials

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) SetServiceAccountCredentials(v string)`

SetServiceAccountCredentials sets ServiceAccountCredentials field to given value.

### HasServiceAccountCredentials

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) HasServiceAccountCredentials() bool`

HasServiceAccountCredentials returns a boolean if a field has been set.

### GetServiceAccountCredentialsFormat

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) GetServiceAccountCredentialsFormat() IdentityPasswordFormat`

GetServiceAccountCredentialsFormat returns the ServiceAccountCredentialsFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsFormatOk

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) GetServiceAccountCredentialsFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialsFormatOk returns a tuple with the ServiceAccountCredentialsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialsFormat

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) SetServiceAccountCredentialsFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialsFormat sets ServiceAccountCredentialsFormat field to given value.

### HasServiceAccountCredentialsFormat

`func (o *EditHypervisorGCPConnectionRequestModelAllOf) HasServiceAccountCredentialsFormat() bool`

HasServiceAccountCredentialsFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


