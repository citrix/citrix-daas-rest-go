# DeploySecretsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountPassword** | Pointer to **string** | Service account password for required in domain joining. This will be stored in a vault. | [optional] 

## Methods

### NewDeploySecretsModel

`func NewDeploySecretsModel() *DeploySecretsModel`

NewDeploySecretsModel instantiates a new DeploySecretsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploySecretsModelWithDefaults

`func NewDeploySecretsModelWithDefaults() *DeploySecretsModel`

NewDeploySecretsModelWithDefaults instantiates a new DeploySecretsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccountPassword

`func (o *DeploySecretsModel) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *DeploySecretsModel) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *DeploySecretsModel) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.

### HasServiceAccountPassword

`func (o *DeploySecretsModel) HasServiceAccountPassword() bool`

HasServiceAccountPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


