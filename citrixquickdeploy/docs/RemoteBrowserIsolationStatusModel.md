# RemoteBrowserIsolationStatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RemoteBrowserIsolationStatusModelJobState**](RemoteBrowserIsolationStatusModelJobState.md) | Current state of the job | [optional] 
**ErrorDetails** | Pointer to **string** | Error occurred in the job | [optional] 
**InUsedByRemoteBrowserIsolation** | Pointer to **bool** | Indicate if the subscription is in-used by Remote Browser Isolation | [optional] 

## Methods

### NewRemoteBrowserIsolationStatusModel

`func NewRemoteBrowserIsolationStatusModel() *RemoteBrowserIsolationStatusModel`

NewRemoteBrowserIsolationStatusModel instantiates a new RemoteBrowserIsolationStatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteBrowserIsolationStatusModelWithDefaults

`func NewRemoteBrowserIsolationStatusModelWithDefaults() *RemoteBrowserIsolationStatusModel`

NewRemoteBrowserIsolationStatusModelWithDefaults instantiates a new RemoteBrowserIsolationStatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RemoteBrowserIsolationStatusModel) GetStatus() RemoteBrowserIsolationStatusModelJobState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteBrowserIsolationStatusModel) GetStatusOk() (*RemoteBrowserIsolationStatusModelJobState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteBrowserIsolationStatusModel) SetStatus(v RemoteBrowserIsolationStatusModelJobState)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteBrowserIsolationStatusModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorDetails

`func (o *RemoteBrowserIsolationStatusModel) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *RemoteBrowserIsolationStatusModel) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *RemoteBrowserIsolationStatusModel) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *RemoteBrowserIsolationStatusModel) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetInUsedByRemoteBrowserIsolation

`func (o *RemoteBrowserIsolationStatusModel) GetInUsedByRemoteBrowserIsolation() bool`

GetInUsedByRemoteBrowserIsolation returns the InUsedByRemoteBrowserIsolation field if non-nil, zero value otherwise.

### GetInUsedByRemoteBrowserIsolationOk

`func (o *RemoteBrowserIsolationStatusModel) GetInUsedByRemoteBrowserIsolationOk() (*bool, bool)`

GetInUsedByRemoteBrowserIsolationOk returns a tuple with the InUsedByRemoteBrowserIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUsedByRemoteBrowserIsolation

`func (o *RemoteBrowserIsolationStatusModel) SetInUsedByRemoteBrowserIsolation(v bool)`

SetInUsedByRemoteBrowserIsolation sets InUsedByRemoteBrowserIsolation field to given value.

### HasInUsedByRemoteBrowserIsolation

`func (o *RemoteBrowserIsolationStatusModel) HasInUsedByRemoteBrowserIsolation() bool`

HasInUsedByRemoteBrowserIsolation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


