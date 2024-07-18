# DeploymentWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarningType** | Pointer to [**DeploymentWarningType**](DeploymentWarningType.md) |  | [optional] 
**Id** | Pointer to **NullableString** | Id of the warning | [optional] 
**Warning** | Pointer to **NullableString** | The warning message | [optional] 
**TaskId** | Pointer to **NullableString** | ID of the task the warning was created by | [optional] 
**IsError** | Pointer to **bool** | Indicates if the warning should be treated as an error | [optional] 
**IsDismissible** | Pointer to **bool** | Indicates if the warning message can be cleared out by the user | [optional] 
**TransactionId** | Pointer to **NullableString** | ID of the transaction the warning was created by | [optional] 

## Methods

### NewDeploymentWarning

`func NewDeploymentWarning() *DeploymentWarning`

NewDeploymentWarning instantiates a new DeploymentWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWarningWithDefaults

`func NewDeploymentWarningWithDefaults() *DeploymentWarning`

NewDeploymentWarningWithDefaults instantiates a new DeploymentWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarningType

`func (o *DeploymentWarning) GetWarningType() DeploymentWarningType`

GetWarningType returns the WarningType field if non-nil, zero value otherwise.

### GetWarningTypeOk

`func (o *DeploymentWarning) GetWarningTypeOk() (*DeploymentWarningType, bool)`

GetWarningTypeOk returns a tuple with the WarningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningType

`func (o *DeploymentWarning) SetWarningType(v DeploymentWarningType)`

SetWarningType sets WarningType field to given value.

### HasWarningType

`func (o *DeploymentWarning) HasWarningType() bool`

HasWarningType returns a boolean if a field has been set.

### GetId

`func (o *DeploymentWarning) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentWarning) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentWarning) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentWarning) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeploymentWarning) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeploymentWarning) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetWarning

`func (o *DeploymentWarning) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *DeploymentWarning) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *DeploymentWarning) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *DeploymentWarning) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### SetWarningNil

`func (o *DeploymentWarning) SetWarningNil(b bool)`

 SetWarningNil sets the value for Warning to be an explicit nil

### UnsetWarning
`func (o *DeploymentWarning) UnsetWarning()`

UnsetWarning ensures that no value is present for Warning, not even an explicit nil
### GetTaskId

`func (o *DeploymentWarning) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *DeploymentWarning) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *DeploymentWarning) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *DeploymentWarning) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *DeploymentWarning) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *DeploymentWarning) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetIsError

`func (o *DeploymentWarning) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *DeploymentWarning) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *DeploymentWarning) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *DeploymentWarning) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetIsDismissible

`func (o *DeploymentWarning) GetIsDismissible() bool`

GetIsDismissible returns the IsDismissible field if non-nil, zero value otherwise.

### GetIsDismissibleOk

`func (o *DeploymentWarning) GetIsDismissibleOk() (*bool, bool)`

GetIsDismissibleOk returns a tuple with the IsDismissible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDismissible

`func (o *DeploymentWarning) SetIsDismissible(v bool)`

SetIsDismissible sets IsDismissible field to given value.

### HasIsDismissible

`func (o *DeploymentWarning) HasIsDismissible() bool`

HasIsDismissible returns a boolean if a field has been set.

### GetTransactionId

`func (o *DeploymentWarning) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DeploymentWarning) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DeploymentWarning) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DeploymentWarning) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *DeploymentWarning) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *DeploymentWarning) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


