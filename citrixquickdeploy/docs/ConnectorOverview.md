# ConnectorOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorName** | Pointer to **string** | Name of the Edge Server VM | [optional] 
**AzureResourceGroup** | Pointer to **string** | Resource group of the Edge Server VM | [optional] 
**State** | [**ConnectorJobState**](ConnectorJobState.md) | The current state of the connector install job | 
**Status** | Pointer to **string** | Any status message that needs to be shown to the user | [optional] 
**LastStatusModified** | Pointer to **time.Time** | Last time the status was modified | [optional] 
**Error** | Pointer to **string** | Error message in case of failures | [optional] 
**CompletionPercentage** | Pointer to **int32** | The completion percentage of the current job | [optional] 
**IsRebootInProgress** | Pointer to **bool** | Whether the connector is currently being rebooted or not. | [optional] 

## Methods

### NewConnectorOverview

`func NewConnectorOverview(state ConnectorJobState, ) *ConnectorOverview`

NewConnectorOverview instantiates a new ConnectorOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorOverviewWithDefaults

`func NewConnectorOverviewWithDefaults() *ConnectorOverview`

NewConnectorOverviewWithDefaults instantiates a new ConnectorOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorName

`func (o *ConnectorOverview) GetConnectorName() string`

GetConnectorName returns the ConnectorName field if non-nil, zero value otherwise.

### GetConnectorNameOk

`func (o *ConnectorOverview) GetConnectorNameOk() (*string, bool)`

GetConnectorNameOk returns a tuple with the ConnectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorName

`func (o *ConnectorOverview) SetConnectorName(v string)`

SetConnectorName sets ConnectorName field to given value.

### HasConnectorName

`func (o *ConnectorOverview) HasConnectorName() bool`

HasConnectorName returns a boolean if a field has been set.

### GetAzureResourceGroup

`func (o *ConnectorOverview) GetAzureResourceGroup() string`

GetAzureResourceGroup returns the AzureResourceGroup field if non-nil, zero value otherwise.

### GetAzureResourceGroupOk

`func (o *ConnectorOverview) GetAzureResourceGroupOk() (*string, bool)`

GetAzureResourceGroupOk returns a tuple with the AzureResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroup

`func (o *ConnectorOverview) SetAzureResourceGroup(v string)`

SetAzureResourceGroup sets AzureResourceGroup field to given value.

### HasAzureResourceGroup

`func (o *ConnectorOverview) HasAzureResourceGroup() bool`

HasAzureResourceGroup returns a boolean if a field has been set.

### GetState

`func (o *ConnectorOverview) GetState() ConnectorJobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectorOverview) GetStateOk() (*ConnectorJobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectorOverview) SetState(v ConnectorJobState)`

SetState sets State field to given value.


### GetStatus

`func (o *ConnectorOverview) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorOverview) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorOverview) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectorOverview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastStatusModified

`func (o *ConnectorOverview) GetLastStatusModified() time.Time`

GetLastStatusModified returns the LastStatusModified field if non-nil, zero value otherwise.

### GetLastStatusModifiedOk

`func (o *ConnectorOverview) GetLastStatusModifiedOk() (*time.Time, bool)`

GetLastStatusModifiedOk returns a tuple with the LastStatusModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusModified

`func (o *ConnectorOverview) SetLastStatusModified(v time.Time)`

SetLastStatusModified sets LastStatusModified field to given value.

### HasLastStatusModified

`func (o *ConnectorOverview) HasLastStatusModified() bool`

HasLastStatusModified returns a boolean if a field has been set.

### GetError

`func (o *ConnectorOverview) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConnectorOverview) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConnectorOverview) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConnectorOverview) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCompletionPercentage

`func (o *ConnectorOverview) GetCompletionPercentage() int32`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *ConnectorOverview) GetCompletionPercentageOk() (*int32, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *ConnectorOverview) SetCompletionPercentage(v int32)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *ConnectorOverview) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### GetIsRebootInProgress

`func (o *ConnectorOverview) GetIsRebootInProgress() bool`

GetIsRebootInProgress returns the IsRebootInProgress field if non-nil, zero value otherwise.

### GetIsRebootInProgressOk

`func (o *ConnectorOverview) GetIsRebootInProgressOk() (*bool, bool)`

GetIsRebootInProgressOk returns a tuple with the IsRebootInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRebootInProgress

`func (o *ConnectorOverview) SetIsRebootInProgress(v bool)`

SetIsRebootInProgress sets IsRebootInProgress field to given value.

### HasIsRebootInProgress

`func (o *ConnectorOverview) HasIsRebootInProgress() bool`

HasIsRebootInProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


