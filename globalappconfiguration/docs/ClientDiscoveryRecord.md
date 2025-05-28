# ClientDiscoveryRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientInput** | Pointer to **string** |  | [optional] 
**Workspace** | Pointer to [**ClientWorkspaceModel**](ClientWorkspaceModel.md) |  | [optional] 

## Methods

### NewClientDiscoveryRecord

`func NewClientDiscoveryRecord() *ClientDiscoveryRecord`

NewClientDiscoveryRecord instantiates a new ClientDiscoveryRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDiscoveryRecordWithDefaults

`func NewClientDiscoveryRecordWithDefaults() *ClientDiscoveryRecord`

NewClientDiscoveryRecordWithDefaults instantiates a new ClientDiscoveryRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientInput

`func (o *ClientDiscoveryRecord) GetClientInput() string`

GetClientInput returns the ClientInput field if non-nil, zero value otherwise.

### GetClientInputOk

`func (o *ClientDiscoveryRecord) GetClientInputOk() (*string, bool)`

GetClientInputOk returns a tuple with the ClientInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientInput

`func (o *ClientDiscoveryRecord) SetClientInput(v string)`

SetClientInput sets ClientInput field to given value.

### HasClientInput

`func (o *ClientDiscoveryRecord) HasClientInput() bool`

HasClientInput returns a boolean if a field has been set.

### GetWorkspace

`func (o *ClientDiscoveryRecord) GetWorkspace() ClientWorkspaceModel`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *ClientDiscoveryRecord) GetWorkspaceOk() (*ClientWorkspaceModel, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *ClientDiscoveryRecord) SetWorkspace(v ClientWorkspaceModel)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *ClientDiscoveryRecord) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


