# ActionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionId** | Pointer to **string** | The action id. | [optional] 
**ActionTargetName** | Pointer to **string** | The action target name, it&#39;s the related catalog name. | [optional] 
**ActionTargetUid** | Pointer to **string** | The action target uid, it&#39;s the related catalog uid. | [optional] 
**ActionTargetId** | Pointer to **string** | The action target id, it&#39;s related on catalog id. | [optional] 
**ActionType** | Pointer to [**ActionType**](ActionType.md) |  | [optional] 
**CreationTime** | Pointer to **string** | The action creation time. | [optional] 
**StartTime** | Pointer to **string** | The action start time. | [optional] 
**FinishTime** | Pointer to **string** | The action finish time. | [optional] 
**State** | Pointer to [**ActionState**](ActionState.md) |  | [optional] 
**ErrorState** | Pointer to [**ActionErrorStatus**](ActionErrorStatus.md) |  | [optional] 
**Progress** | Pointer to **float64** | The progress of the action. | [optional] 
**ProgressMessage** | Pointer to **string** | The progress message. | [optional] 
**TerminatingError** | Pointer to [**ActionResponseModelTerminatingError**](ActionResponseModelTerminatingError.md) |  | [optional] 
**NonTerminatingErrors** | Pointer to [**[]ActionError**](ActionError.md) | The non terminating errors. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of the action. | [optional] 
**Scopes** | Pointer to [**[]ScopeResponseModel**](ScopeResponseModel.md) | The scope of the catalog. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the hypervisor is assigned to.  If &#x60;null&#x60;, the hypervisor is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**IsRunning** | Pointer to **bool** | If the task is running. | [optional] 
**MasterImage** | Pointer to **string** | The master image. | [optional] 
**MachineCreationData** | Pointer to [**ActionCatalogCreationResponseModelAllOfMachineCreationData**](ActionCatalogCreationResponseModelAllOfMachineCreationData.md) |  | [optional] 
**Snapshot** | Pointer to **string** | The snapshot. | [optional] 
**MachineRemovalData** | Pointer to [**ActionMachineRemovalResponseModelAllOfMachineRemovalData**](ActionMachineRemovalResponseModelAllOfMachineRemovalData.md) |  | [optional] 
**DeleteVirtualMachines** | Pointer to **string** | The delete virtual machines.  | [optional] 

## Methods

### NewActionResponseModel

`func NewActionResponseModel() *ActionResponseModel`

NewActionResponseModel instantiates a new ActionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResponseModelWithDefaults

`func NewActionResponseModelWithDefaults() *ActionResponseModel`

NewActionResponseModelWithDefaults instantiates a new ActionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *ActionResponseModel) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ActionResponseModel) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ActionResponseModel) SetActionId(v string)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *ActionResponseModel) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### GetActionTargetName

`func (o *ActionResponseModel) GetActionTargetName() string`

GetActionTargetName returns the ActionTargetName field if non-nil, zero value otherwise.

### GetActionTargetNameOk

`func (o *ActionResponseModel) GetActionTargetNameOk() (*string, bool)`

GetActionTargetNameOk returns a tuple with the ActionTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetName

`func (o *ActionResponseModel) SetActionTargetName(v string)`

SetActionTargetName sets ActionTargetName field to given value.

### HasActionTargetName

`func (o *ActionResponseModel) HasActionTargetName() bool`

HasActionTargetName returns a boolean if a field has been set.

### GetActionTargetUid

`func (o *ActionResponseModel) GetActionTargetUid() string`

GetActionTargetUid returns the ActionTargetUid field if non-nil, zero value otherwise.

### GetActionTargetUidOk

`func (o *ActionResponseModel) GetActionTargetUidOk() (*string, bool)`

GetActionTargetUidOk returns a tuple with the ActionTargetUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetUid

`func (o *ActionResponseModel) SetActionTargetUid(v string)`

SetActionTargetUid sets ActionTargetUid field to given value.

### HasActionTargetUid

`func (o *ActionResponseModel) HasActionTargetUid() bool`

HasActionTargetUid returns a boolean if a field has been set.

### GetActionTargetId

`func (o *ActionResponseModel) GetActionTargetId() string`

GetActionTargetId returns the ActionTargetId field if non-nil, zero value otherwise.

### GetActionTargetIdOk

`func (o *ActionResponseModel) GetActionTargetIdOk() (*string, bool)`

GetActionTargetIdOk returns a tuple with the ActionTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetId

`func (o *ActionResponseModel) SetActionTargetId(v string)`

SetActionTargetId sets ActionTargetId field to given value.

### HasActionTargetId

`func (o *ActionResponseModel) HasActionTargetId() bool`

HasActionTargetId returns a boolean if a field has been set.

### GetActionType

`func (o *ActionResponseModel) GetActionType() ActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ActionResponseModel) GetActionTypeOk() (*ActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ActionResponseModel) SetActionType(v ActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ActionResponseModel) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetCreationTime

`func (o *ActionResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ActionResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ActionResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ActionResponseModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetStartTime

`func (o *ActionResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ActionResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ActionResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ActionResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetFinishTime

`func (o *ActionResponseModel) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *ActionResponseModel) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *ActionResponseModel) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *ActionResponseModel) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetState

`func (o *ActionResponseModel) GetState() ActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActionResponseModel) GetStateOk() (*ActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActionResponseModel) SetState(v ActionState)`

SetState sets State field to given value.

### HasState

`func (o *ActionResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetErrorState

`func (o *ActionResponseModel) GetErrorState() ActionErrorStatus`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *ActionResponseModel) GetErrorStateOk() (*ActionErrorStatus, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *ActionResponseModel) SetErrorState(v ActionErrorStatus)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *ActionResponseModel) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetProgress

`func (o *ActionResponseModel) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ActionResponseModel) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ActionResponseModel) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ActionResponseModel) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressMessage

`func (o *ActionResponseModel) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *ActionResponseModel) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *ActionResponseModel) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *ActionResponseModel) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### GetTerminatingError

`func (o *ActionResponseModel) GetTerminatingError() ActionResponseModelTerminatingError`

GetTerminatingError returns the TerminatingError field if non-nil, zero value otherwise.

### GetTerminatingErrorOk

`func (o *ActionResponseModel) GetTerminatingErrorOk() (*ActionResponseModelTerminatingError, bool)`

GetTerminatingErrorOk returns a tuple with the TerminatingError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatingError

`func (o *ActionResponseModel) SetTerminatingError(v ActionResponseModelTerminatingError)`

SetTerminatingError sets TerminatingError field to given value.

### HasTerminatingError

`func (o *ActionResponseModel) HasTerminatingError() bool`

HasTerminatingError returns a boolean if a field has been set.

### GetNonTerminatingErrors

`func (o *ActionResponseModel) GetNonTerminatingErrors() []ActionError`

GetNonTerminatingErrors returns the NonTerminatingErrors field if non-nil, zero value otherwise.

### GetNonTerminatingErrorsOk

`func (o *ActionResponseModel) GetNonTerminatingErrorsOk() (*[]ActionError, bool)`

GetNonTerminatingErrorsOk returns a tuple with the NonTerminatingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTerminatingErrors

`func (o *ActionResponseModel) SetNonTerminatingErrors(v []ActionError)`

SetNonTerminatingErrors sets NonTerminatingErrors field to given value.

### HasNonTerminatingErrors

`func (o *ActionResponseModel) HasNonTerminatingErrors() bool`

HasNonTerminatingErrors returns a boolean if a field has been set.

### GetMetadata

`func (o *ActionResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ActionResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScopes

`func (o *ActionResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ActionResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ActionResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ActionResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTenants

`func (o *ActionResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ActionResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ActionResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ActionResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetIsRunning

`func (o *ActionResponseModel) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *ActionResponseModel) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *ActionResponseModel) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.

### HasIsRunning

`func (o *ActionResponseModel) HasIsRunning() bool`

HasIsRunning returns a boolean if a field has been set.

### GetMasterImage

`func (o *ActionResponseModel) GetMasterImage() string`

GetMasterImage returns the MasterImage field if non-nil, zero value otherwise.

### GetMasterImageOk

`func (o *ActionResponseModel) GetMasterImageOk() (*string, bool)`

GetMasterImageOk returns a tuple with the MasterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImage

`func (o *ActionResponseModel) SetMasterImage(v string)`

SetMasterImage sets MasterImage field to given value.

### HasMasterImage

`func (o *ActionResponseModel) HasMasterImage() bool`

HasMasterImage returns a boolean if a field has been set.

### GetMachineCreationData

`func (o *ActionResponseModel) GetMachineCreationData() ActionCatalogCreationResponseModelAllOfMachineCreationData`

GetMachineCreationData returns the MachineCreationData field if non-nil, zero value otherwise.

### GetMachineCreationDataOk

`func (o *ActionResponseModel) GetMachineCreationDataOk() (*ActionCatalogCreationResponseModelAllOfMachineCreationData, bool)`

GetMachineCreationDataOk returns a tuple with the MachineCreationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCreationData

`func (o *ActionResponseModel) SetMachineCreationData(v ActionCatalogCreationResponseModelAllOfMachineCreationData)`

SetMachineCreationData sets MachineCreationData field to given value.

### HasMachineCreationData

`func (o *ActionResponseModel) HasMachineCreationData() bool`

HasMachineCreationData returns a boolean if a field has been set.

### GetSnapshot

`func (o *ActionResponseModel) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *ActionResponseModel) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *ActionResponseModel) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *ActionResponseModel) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetMachineRemovalData

`func (o *ActionResponseModel) GetMachineRemovalData() ActionMachineRemovalResponseModelAllOfMachineRemovalData`

GetMachineRemovalData returns the MachineRemovalData field if non-nil, zero value otherwise.

### GetMachineRemovalDataOk

`func (o *ActionResponseModel) GetMachineRemovalDataOk() (*ActionMachineRemovalResponseModelAllOfMachineRemovalData, bool)`

GetMachineRemovalDataOk returns a tuple with the MachineRemovalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineRemovalData

`func (o *ActionResponseModel) SetMachineRemovalData(v ActionMachineRemovalResponseModelAllOfMachineRemovalData)`

SetMachineRemovalData sets MachineRemovalData field to given value.

### HasMachineRemovalData

`func (o *ActionResponseModel) HasMachineRemovalData() bool`

HasMachineRemovalData returns a boolean if a field has been set.

### GetDeleteVirtualMachines

`func (o *ActionResponseModel) GetDeleteVirtualMachines() string`

GetDeleteVirtualMachines returns the DeleteVirtualMachines field if non-nil, zero value otherwise.

### GetDeleteVirtualMachinesOk

`func (o *ActionResponseModel) GetDeleteVirtualMachinesOk() (*string, bool)`

GetDeleteVirtualMachinesOk returns a tuple with the DeleteVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVirtualMachines

`func (o *ActionResponseModel) SetDeleteVirtualMachines(v string)`

SetDeleteVirtualMachines sets DeleteVirtualMachines field to given value.

### HasDeleteVirtualMachines

`func (o *ActionResponseModel) HasDeleteVirtualMachines() bool`

HasDeleteVirtualMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


