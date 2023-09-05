# MachineCatalogResponseModelProvisioningProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRunning** | **bool** | Indicates whether the provisioning task is running. | 
**Progress** | **int32** | Specifies the current progress of the provisioning task. | 
**ProgressMessage** | Pointer to **string** | Specifies the message of current provisioning task. | [optional] 
**ProgressMessageList** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewMachineCatalogResponseModelProvisioningProgress

`func NewMachineCatalogResponseModelProvisioningProgress(isRunning bool, progress int32, ) *MachineCatalogResponseModelProvisioningProgress`

NewMachineCatalogResponseModelProvisioningProgress instantiates a new MachineCatalogResponseModelProvisioningProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogResponseModelProvisioningProgressWithDefaults

`func NewMachineCatalogResponseModelProvisioningProgressWithDefaults() *MachineCatalogResponseModelProvisioningProgress`

NewMachineCatalogResponseModelProvisioningProgressWithDefaults instantiates a new MachineCatalogResponseModelProvisioningProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRunning

`func (o *MachineCatalogResponseModelProvisioningProgress) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *MachineCatalogResponseModelProvisioningProgress) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *MachineCatalogResponseModelProvisioningProgress) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.


### GetProgress

`func (o *MachineCatalogResponseModelProvisioningProgress) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *MachineCatalogResponseModelProvisioningProgress) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *MachineCatalogResponseModelProvisioningProgress) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetProgressMessage

`func (o *MachineCatalogResponseModelProvisioningProgress) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *MachineCatalogResponseModelProvisioningProgress) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *MachineCatalogResponseModelProvisioningProgress) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *MachineCatalogResponseModelProvisioningProgress) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### GetProgressMessageList

`func (o *MachineCatalogResponseModelProvisioningProgress) GetProgressMessageList() [][]string`

GetProgressMessageList returns the ProgressMessageList field if non-nil, zero value otherwise.

### GetProgressMessageListOk

`func (o *MachineCatalogResponseModelProvisioningProgress) GetProgressMessageListOk() (*[][]string, bool)`

GetProgressMessageListOk returns a tuple with the ProgressMessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessageList

`func (o *MachineCatalogResponseModelProvisioningProgress) SetProgressMessageList(v [][]string)`

SetProgressMessageList sets ProgressMessageList field to given value.

### HasProgressMessageList

`func (o *MachineCatalogResponseModelProvisioningProgress) HasProgressMessageList() bool`

HasProgressMessageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


