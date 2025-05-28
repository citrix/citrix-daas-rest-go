# UpdateCustomerAccountAsyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The new name of account | [optional] 
**AccountOperationType** | [**UpdateAccountOperationType**](UpdateAccountOperationType.md) | The type of provider associated with the account | 
**AwsAccessKeyId** | Pointer to **NullableString** | The ID of Access Key associated with the account | [optional] 
**AwsSecretAccessKey** | Pointer to **NullableString** | The secret associated with access key for the account | [optional] 
**AwsRoleArn** | Pointer to **NullableString** | The ARN of the role to assume | [optional] 

## Methods

### NewUpdateCustomerAccountAsyncRequest

`func NewUpdateCustomerAccountAsyncRequest(accountOperationType UpdateAccountOperationType, ) *UpdateCustomerAccountAsyncRequest`

NewUpdateCustomerAccountAsyncRequest instantiates a new UpdateCustomerAccountAsyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerAccountAsyncRequestWithDefaults

`func NewUpdateCustomerAccountAsyncRequestWithDefaults() *UpdateCustomerAccountAsyncRequest`

NewUpdateCustomerAccountAsyncRequestWithDefaults instantiates a new UpdateCustomerAccountAsyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCustomerAccountAsyncRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomerAccountAsyncRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomerAccountAsyncRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomerAccountAsyncRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateCustomerAccountAsyncRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateCustomerAccountAsyncRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccountOperationType

`func (o *UpdateCustomerAccountAsyncRequest) GetAccountOperationType() UpdateAccountOperationType`

GetAccountOperationType returns the AccountOperationType field if non-nil, zero value otherwise.

### GetAccountOperationTypeOk

`func (o *UpdateCustomerAccountAsyncRequest) GetAccountOperationTypeOk() (*UpdateAccountOperationType, bool)`

GetAccountOperationTypeOk returns a tuple with the AccountOperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOperationType

`func (o *UpdateCustomerAccountAsyncRequest) SetAccountOperationType(v UpdateAccountOperationType)`

SetAccountOperationType sets AccountOperationType field to given value.


### GetAwsAccessKeyId

`func (o *UpdateCustomerAccountAsyncRequest) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *UpdateCustomerAccountAsyncRequest) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *UpdateCustomerAccountAsyncRequest) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *UpdateCustomerAccountAsyncRequest) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### SetAwsAccessKeyIdNil

`func (o *UpdateCustomerAccountAsyncRequest) SetAwsAccessKeyIdNil(b bool)`

 SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil

### UnsetAwsAccessKeyId
`func (o *UpdateCustomerAccountAsyncRequest) UnsetAwsAccessKeyId()`

UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *UpdateCustomerAccountAsyncRequest) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *UpdateCustomerAccountAsyncRequest) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *UpdateCustomerAccountAsyncRequest) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *UpdateCustomerAccountAsyncRequest) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### SetAwsSecretAccessKeyNil

`func (o *UpdateCustomerAccountAsyncRequest) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *UpdateCustomerAccountAsyncRequest) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
### GetAwsRoleArn

`func (o *UpdateCustomerAccountAsyncRequest) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *UpdateCustomerAccountAsyncRequest) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *UpdateCustomerAccountAsyncRequest) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.

### HasAwsRoleArn

`func (o *UpdateCustomerAccountAsyncRequest) HasAwsRoleArn() bool`

HasAwsRoleArn returns a boolean if a field has been set.

### SetAwsRoleArnNil

`func (o *UpdateCustomerAccountAsyncRequest) SetAwsRoleArnNil(b bool)`

 SetAwsRoleArnNil sets the value for AwsRoleArn to be an explicit nil

### UnsetAwsRoleArn
`func (o *UpdateCustomerAccountAsyncRequest) UnsetAwsRoleArn()`

UnsetAwsRoleArn ensures that no value is present for AwsRoleArn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


