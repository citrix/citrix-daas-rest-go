# CatalogCapacitySettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeWorker** | Pointer to [**CatalogComputeWorkerModel**](CatalogComputeWorkerModel.md) | Compute settings for the catalog | [optional] 
**ScaleSettings** | Pointer to [**CatalogScaleSettingsModel**](CatalogScaleSettingsModel.md) | Scale settings for the catalog | [optional] 
**SessionTimeout** | Pointer to **int32** | Idle timeout for session in the catalog (in mins) | [optional] 
**MultiSessionDisconnectedSessionTimeout** | Pointer to **int32** | Minutes to wait for disconnected sessions to be logged off on multi-session VMs | [optional] 
**IsActive** | Pointer to **bool** | Indicates if the capacity job is currently active | [optional] [readonly] 

## Methods

### NewCatalogCapacitySettingsModel

`func NewCatalogCapacitySettingsModel() *CatalogCapacitySettingsModel`

NewCatalogCapacitySettingsModel instantiates a new CatalogCapacitySettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogCapacitySettingsModelWithDefaults

`func NewCatalogCapacitySettingsModelWithDefaults() *CatalogCapacitySettingsModel`

NewCatalogCapacitySettingsModelWithDefaults instantiates a new CatalogCapacitySettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeWorker

`func (o *CatalogCapacitySettingsModel) GetComputeWorker() CatalogComputeWorkerModel`

GetComputeWorker returns the ComputeWorker field if non-nil, zero value otherwise.

### GetComputeWorkerOk

`func (o *CatalogCapacitySettingsModel) GetComputeWorkerOk() (*CatalogComputeWorkerModel, bool)`

GetComputeWorkerOk returns a tuple with the ComputeWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeWorker

`func (o *CatalogCapacitySettingsModel) SetComputeWorker(v CatalogComputeWorkerModel)`

SetComputeWorker sets ComputeWorker field to given value.

### HasComputeWorker

`func (o *CatalogCapacitySettingsModel) HasComputeWorker() bool`

HasComputeWorker returns a boolean if a field has been set.

### GetScaleSettings

`func (o *CatalogCapacitySettingsModel) GetScaleSettings() CatalogScaleSettingsModel`

GetScaleSettings returns the ScaleSettings field if non-nil, zero value otherwise.

### GetScaleSettingsOk

`func (o *CatalogCapacitySettingsModel) GetScaleSettingsOk() (*CatalogScaleSettingsModel, bool)`

GetScaleSettingsOk returns a tuple with the ScaleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleSettings

`func (o *CatalogCapacitySettingsModel) SetScaleSettings(v CatalogScaleSettingsModel)`

SetScaleSettings sets ScaleSettings field to given value.

### HasScaleSettings

`func (o *CatalogCapacitySettingsModel) HasScaleSettings() bool`

HasScaleSettings returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *CatalogCapacitySettingsModel) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *CatalogCapacitySettingsModel) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *CatalogCapacitySettingsModel) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *CatalogCapacitySettingsModel) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetMultiSessionDisconnectedSessionTimeout

`func (o *CatalogCapacitySettingsModel) GetMultiSessionDisconnectedSessionTimeout() int32`

GetMultiSessionDisconnectedSessionTimeout returns the MultiSessionDisconnectedSessionTimeout field if non-nil, zero value otherwise.

### GetMultiSessionDisconnectedSessionTimeoutOk

`func (o *CatalogCapacitySettingsModel) GetMultiSessionDisconnectedSessionTimeoutOk() (*int32, bool)`

GetMultiSessionDisconnectedSessionTimeoutOk returns a tuple with the MultiSessionDisconnectedSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSessionDisconnectedSessionTimeout

`func (o *CatalogCapacitySettingsModel) SetMultiSessionDisconnectedSessionTimeout(v int32)`

SetMultiSessionDisconnectedSessionTimeout sets MultiSessionDisconnectedSessionTimeout field to given value.

### HasMultiSessionDisconnectedSessionTimeout

`func (o *CatalogCapacitySettingsModel) HasMultiSessionDisconnectedSessionTimeout() bool`

HasMultiSessionDisconnectedSessionTimeout returns a boolean if a field has been set.

### GetIsActive

`func (o *CatalogCapacitySettingsModel) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CatalogCapacitySettingsModel) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CatalogCapacitySettingsModel) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CatalogCapacitySettingsModel) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


