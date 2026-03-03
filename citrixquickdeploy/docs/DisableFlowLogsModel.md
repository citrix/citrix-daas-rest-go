# DisableFlowLogsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFlowLog** | Pointer to **bool** | Whether to delete the flow log resource from Azure.  Default is false - flow log is disabled (Enabled &#x3D; false) but not deleted.  Set to true to permanently delete the flow log resource. | [optional] 
**DeleteStorageAccount** | Pointer to **bool** | Whether to delete the storage account when disabling flow logs.  Default is false - storage account is retained for other VNets in the region. | [optional] 
**Force** | Pointer to **bool** | Force disable by checking Azure for flow logs even if not found in DataStore.  Useful for cleanup when DataStore is out of sync with Azure.  Default is false. | [optional] 

## Methods

### NewDisableFlowLogsModel

`func NewDisableFlowLogsModel() *DisableFlowLogsModel`

NewDisableFlowLogsModel instantiates a new DisableFlowLogsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableFlowLogsModelWithDefaults

`func NewDisableFlowLogsModelWithDefaults() *DisableFlowLogsModel`

NewDisableFlowLogsModelWithDefaults instantiates a new DisableFlowLogsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFlowLog

`func (o *DisableFlowLogsModel) GetDeleteFlowLog() bool`

GetDeleteFlowLog returns the DeleteFlowLog field if non-nil, zero value otherwise.

### GetDeleteFlowLogOk

`func (o *DisableFlowLogsModel) GetDeleteFlowLogOk() (*bool, bool)`

GetDeleteFlowLogOk returns a tuple with the DeleteFlowLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFlowLog

`func (o *DisableFlowLogsModel) SetDeleteFlowLog(v bool)`

SetDeleteFlowLog sets DeleteFlowLog field to given value.

### HasDeleteFlowLog

`func (o *DisableFlowLogsModel) HasDeleteFlowLog() bool`

HasDeleteFlowLog returns a boolean if a field has been set.

### GetDeleteStorageAccount

`func (o *DisableFlowLogsModel) GetDeleteStorageAccount() bool`

GetDeleteStorageAccount returns the DeleteStorageAccount field if non-nil, zero value otherwise.

### GetDeleteStorageAccountOk

`func (o *DisableFlowLogsModel) GetDeleteStorageAccountOk() (*bool, bool)`

GetDeleteStorageAccountOk returns a tuple with the DeleteStorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStorageAccount

`func (o *DisableFlowLogsModel) SetDeleteStorageAccount(v bool)`

SetDeleteStorageAccount sets DeleteStorageAccount field to given value.

### HasDeleteStorageAccount

`func (o *DisableFlowLogsModel) HasDeleteStorageAccount() bool`

HasDeleteStorageAccount returns a boolean if a field has been set.

### GetForce

`func (o *DisableFlowLogsModel) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *DisableFlowLogsModel) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *DisableFlowLogsModel) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *DisableFlowLogsModel) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


