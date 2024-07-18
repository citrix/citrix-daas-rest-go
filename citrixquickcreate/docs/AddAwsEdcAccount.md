# AddAwsEdcAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKeyId** | Pointer to **NullableString** | The ID of Access Key associated with the account | [optional] 
**AwsSecretAccessKey** | Pointer to **NullableString** | The secret associated with access key for the account | [optional] 
**AwsRegion** | Pointer to **string** | The AWS region the account is associated with | [optional] 
**AwsRoleArn** | Pointer to **NullableString** | The ARN of the role to assume when making requests in this account | [optional] 
**AwsExternalId** | Pointer to **NullableString** | The ARN external ID to use when assuming the role | [optional] 
**AwsByolFeatureEnabled** | Pointer to **NullableBool** | Indicates if the associated AWS EDC account has BYOL support is enabled\\r\\nBy default we expect customer to use BYOL configuration to be enabled at this moment | [optional] 

## Methods

### NewAddAwsEdcAccount

`func NewAddAwsEdcAccount() *AddAwsEdcAccount`

NewAddAwsEdcAccount instantiates a new AddAwsEdcAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAwsEdcAccountWithDefaults

`func NewAddAwsEdcAccountWithDefaults() *AddAwsEdcAccount`

NewAddAwsEdcAccountWithDefaults instantiates a new AddAwsEdcAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKeyId

`func (o *AddAwsEdcAccount) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *AddAwsEdcAccount) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *AddAwsEdcAccount) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *AddAwsEdcAccount) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### SetAwsAccessKeyIdNil

`func (o *AddAwsEdcAccount) SetAwsAccessKeyIdNil(b bool)`

 SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil

### UnsetAwsAccessKeyId
`func (o *AddAwsEdcAccount) UnsetAwsAccessKeyId()`

UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *AddAwsEdcAccount) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AddAwsEdcAccount) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AddAwsEdcAccount) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *AddAwsEdcAccount) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### SetAwsSecretAccessKeyNil

`func (o *AddAwsEdcAccount) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *AddAwsEdcAccount) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
### GetAwsRegion

`func (o *AddAwsEdcAccount) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AddAwsEdcAccount) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AddAwsEdcAccount) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AddAwsEdcAccount) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAwsRoleArn

`func (o *AddAwsEdcAccount) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *AddAwsEdcAccount) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *AddAwsEdcAccount) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.

### HasAwsRoleArn

`func (o *AddAwsEdcAccount) HasAwsRoleArn() bool`

HasAwsRoleArn returns a boolean if a field has been set.

### SetAwsRoleArnNil

`func (o *AddAwsEdcAccount) SetAwsRoleArnNil(b bool)`

 SetAwsRoleArnNil sets the value for AwsRoleArn to be an explicit nil

### UnsetAwsRoleArn
`func (o *AddAwsEdcAccount) UnsetAwsRoleArn()`

UnsetAwsRoleArn ensures that no value is present for AwsRoleArn, not even an explicit nil
### GetAwsExternalId

`func (o *AddAwsEdcAccount) GetAwsExternalId() string`

GetAwsExternalId returns the AwsExternalId field if non-nil, zero value otherwise.

### GetAwsExternalIdOk

`func (o *AddAwsEdcAccount) GetAwsExternalIdOk() (*string, bool)`

GetAwsExternalIdOk returns a tuple with the AwsExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalId

`func (o *AddAwsEdcAccount) SetAwsExternalId(v string)`

SetAwsExternalId sets AwsExternalId field to given value.

### HasAwsExternalId

`func (o *AddAwsEdcAccount) HasAwsExternalId() bool`

HasAwsExternalId returns a boolean if a field has been set.

### SetAwsExternalIdNil

`func (o *AddAwsEdcAccount) SetAwsExternalIdNil(b bool)`

 SetAwsExternalIdNil sets the value for AwsExternalId to be an explicit nil

### UnsetAwsExternalId
`func (o *AddAwsEdcAccount) UnsetAwsExternalId()`

UnsetAwsExternalId ensures that no value is present for AwsExternalId, not even an explicit nil
### GetAwsByolFeatureEnabled

`func (o *AddAwsEdcAccount) GetAwsByolFeatureEnabled() bool`

GetAwsByolFeatureEnabled returns the AwsByolFeatureEnabled field if non-nil, zero value otherwise.

### GetAwsByolFeatureEnabledOk

`func (o *AddAwsEdcAccount) GetAwsByolFeatureEnabledOk() (*bool, bool)`

GetAwsByolFeatureEnabledOk returns a tuple with the AwsByolFeatureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsByolFeatureEnabled

`func (o *AddAwsEdcAccount) SetAwsByolFeatureEnabled(v bool)`

SetAwsByolFeatureEnabled sets AwsByolFeatureEnabled field to given value.

### HasAwsByolFeatureEnabled

`func (o *AddAwsEdcAccount) HasAwsByolFeatureEnabled() bool`

HasAwsByolFeatureEnabled returns a boolean if a field has been set.

### SetAwsByolFeatureEnabledNil

`func (o *AddAwsEdcAccount) SetAwsByolFeatureEnabledNil(b bool)`

 SetAwsByolFeatureEnabledNil sets the value for AwsByolFeatureEnabled to be an explicit nil

### UnsetAwsByolFeatureEnabled
`func (o *AddAwsEdcAccount) UnsetAwsByolFeatureEnabled()`

UnsetAwsByolFeatureEnabled ensures that no value is present for AwsByolFeatureEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


