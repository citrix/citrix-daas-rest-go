# HypervisorAWSConnectionDetailRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | The API key used to authenticate with the AWS APIs.  Required. | 
**SecretKey** | **string** | The secret key used to authenticate with the AWS APIs.  Required. Must be specified in the format indicated by SecretKeyFormat. | 
**SecretKeyFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**Region** | Pointer to **string** | AWS region to connect to.  Optional.  If not specified, will connect to the global AWS APIs.  This can be used to discover the regions available within AWS.  Access to all other AWS resources requires the region to be set explicitly. | [optional] 

## Methods

### NewHypervisorAWSConnectionDetailRequestModelAllOf

`func NewHypervisorAWSConnectionDetailRequestModelAllOf(apiKey string, secretKey string, ) *HypervisorAWSConnectionDetailRequestModelAllOf`

NewHypervisorAWSConnectionDetailRequestModelAllOf instantiates a new HypervisorAWSConnectionDetailRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorAWSConnectionDetailRequestModelAllOfWithDefaults

`func NewHypervisorAWSConnectionDetailRequestModelAllOfWithDefaults() *HypervisorAWSConnectionDetailRequestModelAllOf`

NewHypervisorAWSConnectionDetailRequestModelAllOfWithDefaults instantiates a new HypervisorAWSConnectionDetailRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetSecretKey

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetSecretKeyFormat

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) GetSecretKeyFormat() IdentityPasswordFormat`

GetSecretKeyFormat returns the SecretKeyFormat field if non-nil, zero value otherwise.

### GetSecretKeyFormatOk

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) GetSecretKeyFormatOk() (*IdentityPasswordFormat, bool)`

GetSecretKeyFormatOk returns a tuple with the SecretKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyFormat

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) SetSecretKeyFormat(v IdentityPasswordFormat)`

SetSecretKeyFormat sets SecretKeyFormat field to given value.

### HasSecretKeyFormat

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) HasSecretKeyFormat() bool`

HasSecretKeyFormat returns a boolean if a field has been set.

### GetRegion

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HypervisorAWSConnectionDetailRequestModelAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


