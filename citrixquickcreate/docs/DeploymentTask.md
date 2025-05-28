# DeploymentTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to [**NullableDeploymentTaskOperationType**](DeploymentTaskOperationType.md) | Task Type | [optional] 
**DeploymentId** | Pointer to **NullableString** | Deployment Id this task is working on | [optional] 

## Methods

### NewDeploymentTask

`func NewDeploymentTask() *DeploymentTask`

NewDeploymentTask instantiates a new DeploymentTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentTaskWithDefaults

`func NewDeploymentTaskWithDefaults() *DeploymentTask`

NewDeploymentTaskWithDefaults instantiates a new DeploymentTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *DeploymentTask) GetOperation() DeploymentTaskOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *DeploymentTask) GetOperationOk() (*DeploymentTaskOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *DeploymentTask) SetOperation(v DeploymentTaskOperationType)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *DeploymentTask) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *DeploymentTask) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *DeploymentTask) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetDeploymentId

`func (o *DeploymentTask) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DeploymentTask) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DeploymentTask) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DeploymentTask) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *DeploymentTask) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *DeploymentTask) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


