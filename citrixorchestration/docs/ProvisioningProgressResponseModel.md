# ProvisioningProgressResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRunning** | **bool** | Indicates whether the provisioning task is running. | 
**Progress** | **int32** | Specifies the current progress of the provisioning task. | 
**ProgressMessage** | Pointer to **NullableString** | Specifies the message of current provisioning task. | [optional] 
**ProgressMessageList** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewProvisioningProgressResponseModel

`func NewProvisioningProgressResponseModel(isRunning bool, progress int32, ) *ProvisioningProgressResponseModel`

NewProvisioningProgressResponseModel instantiates a new ProvisioningProgressResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningProgressResponseModelWithDefaults

`func NewProvisioningProgressResponseModelWithDefaults() *ProvisioningProgressResponseModel`

NewProvisioningProgressResponseModelWithDefaults instantiates a new ProvisioningProgressResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRunning

`func (o *ProvisioningProgressResponseModel) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *ProvisioningProgressResponseModel) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *ProvisioningProgressResponseModel) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.


### GetProgress

`func (o *ProvisioningProgressResponseModel) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ProvisioningProgressResponseModel) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ProvisioningProgressResponseModel) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetProgressMessage

`func (o *ProvisioningProgressResponseModel) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *ProvisioningProgressResponseModel) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *ProvisioningProgressResponseModel) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *ProvisioningProgressResponseModel) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### SetProgressMessageNil

`func (o *ProvisioningProgressResponseModel) SetProgressMessageNil(b bool)`

 SetProgressMessageNil sets the value for ProgressMessage to be an explicit nil

### UnsetProgressMessage
`func (o *ProvisioningProgressResponseModel) UnsetProgressMessage()`

UnsetProgressMessage ensures that no value is present for ProgressMessage, not even an explicit nil
### GetProgressMessageList

`func (o *ProvisioningProgressResponseModel) GetProgressMessageList() [][]string`

GetProgressMessageList returns the ProgressMessageList field if non-nil, zero value otherwise.

### GetProgressMessageListOk

`func (o *ProvisioningProgressResponseModel) GetProgressMessageListOk() (*[][]string, bool)`

GetProgressMessageListOk returns a tuple with the ProgressMessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessageList

`func (o *ProvisioningProgressResponseModel) SetProgressMessageList(v [][]string)`

SetProgressMessageList sets ProgressMessageList field to given value.

### HasProgressMessageList

`func (o *ProvisioningProgressResponseModel) HasProgressMessageList() bool`

HasProgressMessageList returns a boolean if a field has been set.

### SetProgressMessageListNil

`func (o *ProvisioningProgressResponseModel) SetProgressMessageListNil(b bool)`

 SetProgressMessageListNil sets the value for ProgressMessageList to be an explicit nil

### UnsetProgressMessageList
`func (o *ProvisioningProgressResponseModel) UnsetProgressMessageList()`

UnsetProgressMessageList ensures that no value is present for ProgressMessageList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


