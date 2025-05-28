# ResourceConnectionTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to [**NullableResourceConnectionTaskOperationType**](ResourceConnectionTaskOperationType.md) | Task Type | [optional] 
**ResourceConnectionId** | Pointer to **NullableString** | The ResourceConnectionId this task is working on | [optional] 

## Methods

### NewResourceConnectionTask

`func NewResourceConnectionTask() *ResourceConnectionTask`

NewResourceConnectionTask instantiates a new ResourceConnectionTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConnectionTaskWithDefaults

`func NewResourceConnectionTaskWithDefaults() *ResourceConnectionTask`

NewResourceConnectionTaskWithDefaults instantiates a new ResourceConnectionTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ResourceConnectionTask) GetOperation() ResourceConnectionTaskOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ResourceConnectionTask) GetOperationOk() (*ResourceConnectionTaskOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ResourceConnectionTask) SetOperation(v ResourceConnectionTaskOperationType)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ResourceConnectionTask) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *ResourceConnectionTask) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *ResourceConnectionTask) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetResourceConnectionId

`func (o *ResourceConnectionTask) GetResourceConnectionId() string`

GetResourceConnectionId returns the ResourceConnectionId field if non-nil, zero value otherwise.

### GetResourceConnectionIdOk

`func (o *ResourceConnectionTask) GetResourceConnectionIdOk() (*string, bool)`

GetResourceConnectionIdOk returns a tuple with the ResourceConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConnectionId

`func (o *ResourceConnectionTask) SetResourceConnectionId(v string)`

SetResourceConnectionId sets ResourceConnectionId field to given value.

### HasResourceConnectionId

`func (o *ResourceConnectionTask) HasResourceConnectionId() bool`

HasResourceConnectionId returns a boolean if a field has been set.

### SetResourceConnectionIdNil

`func (o *ResourceConnectionTask) SetResourceConnectionIdNil(b bool)`

 SetResourceConnectionIdNil sets the value for ResourceConnectionId to be an explicit nil

### UnsetResourceConnectionId
`func (o *ResourceConnectionTask) UnsetResourceConnectionId()`

UnsetResourceConnectionId ensures that no value is present for ResourceConnectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


