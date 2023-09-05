# ActionUpdateImageResponseModel

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

### NewActionUpdateImageResponseModel

`func NewActionUpdateImageResponseModel() *ActionUpdateImageResponseModel`

NewActionUpdateImageResponseModel instantiates a new ActionUpdateImageResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionUpdateImageResponseModelWithDefaults

`func NewActionUpdateImageResponseModelWithDefaults() *ActionUpdateImageResponseModel`

NewActionUpdateImageResponseModelWithDefaults instantiates a new ActionUpdateImageResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *ActionUpdateImageResponseModel) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ActionUpdateImageResponseModel) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ActionUpdateImageResponseModel) SetActionId(v string)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *ActionUpdateImageResponseModel) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### GetActionTargetName

`func (o *ActionUpdateImageResponseModel) GetActionTargetName() string`

GetActionTargetName returns the ActionTargetName field if non-nil, zero value otherwise.

### GetActionTargetNameOk

`func (o *ActionUpdateImageResponseModel) GetActionTargetNameOk() (*string, bool)`

GetActionTargetNameOk returns a tuple with the ActionTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetName

`func (o *ActionUpdateImageResponseModel) SetActionTargetName(v string)`

SetActionTargetName sets ActionTargetName field to given value.

### HasActionTargetName

`func (o *ActionUpdateImageResponseModel) HasActionTargetName() bool`

HasActionTargetName returns a boolean if a field has been set.

### GetActionTargetUid

`func (o *ActionUpdateImageResponseModel) GetActionTargetUid() string`

GetActionTargetUid returns the ActionTargetUid field if non-nil, zero value otherwise.

### GetActionTargetUidOk

`func (o *ActionUpdateImageResponseModel) GetActionTargetUidOk() (*string, bool)`

GetActionTargetUidOk returns a tuple with the ActionTargetUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetUid

`func (o *ActionUpdateImageResponseModel) SetActionTargetUid(v string)`

SetActionTargetUid sets ActionTargetUid field to given value.

### HasActionTargetUid

`func (o *ActionUpdateImageResponseModel) HasActionTargetUid() bool`

HasActionTargetUid returns a boolean if a field has been set.

### GetActionTargetId

`func (o *ActionUpdateImageResponseModel) GetActionTargetId() string`

GetActionTargetId returns the ActionTargetId field if non-nil, zero value otherwise.

### GetActionTargetIdOk

`func (o *ActionUpdateImageResponseModel) GetActionTargetIdOk() (*string, bool)`

GetActionTargetIdOk returns a tuple with the ActionTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetId

`func (o *ActionUpdateImageResponseModel) SetActionTargetId(v string)`

SetActionTargetId sets ActionTargetId field to given value.

### HasActionTargetId

`func (o *ActionUpdateImageResponseModel) HasActionTargetId() bool`

HasActionTargetId returns a boolean if a field has been set.

### GetActionType

`func (o *ActionUpdateImageResponseModel) GetActionType() ActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ActionUpdateImageResponseModel) GetActionTypeOk() (*ActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ActionUpdateImageResponseModel) SetActionType(v ActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ActionUpdateImageResponseModel) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetCreationTime

`func (o *ActionUpdateImageResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ActionUpdateImageResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ActionUpdateImageResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ActionUpdateImageResponseModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetStartTime

`func (o *ActionUpdateImageResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ActionUpdateImageResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ActionUpdateImageResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ActionUpdateImageResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetFinishTime

`func (o *ActionUpdateImageResponseModel) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *ActionUpdateImageResponseModel) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *ActionUpdateImageResponseModel) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *ActionUpdateImageResponseModel) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetState

`func (o *ActionUpdateImageResponseModel) GetState() ActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActionUpdateImageResponseModel) GetStateOk() (*ActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActionUpdateImageResponseModel) SetState(v ActionState)`

SetState sets State field to given value.

### HasState

`func (o *ActionUpdateImageResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetErrorState

`func (o *ActionUpdateImageResponseModel) GetErrorState() ActionErrorStatus`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *ActionUpdateImageResponseModel) GetErrorStateOk() (*ActionErrorStatus, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *ActionUpdateImageResponseModel) SetErrorState(v ActionErrorStatus)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *ActionUpdateImageResponseModel) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetProgress

`func (o *ActionUpdateImageResponseModel) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ActionUpdateImageResponseModel) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ActionUpdateImageResponseModel) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ActionUpdateImageResponseModel) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressMessage

`func (o *ActionUpdateImageResponseModel) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *ActionUpdateImageResponseModel) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *ActionUpdateImageResponseModel) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *ActionUpdateImageResponseModel) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### GetTerminatingError

`func (o *ActionUpdateImageResponseModel) GetTerminatingError() ActionResponseModelTerminatingError`

GetTerminatingError returns the TerminatingError field if non-nil, zero value otherwise.

### GetTerminatingErrorOk

`func (o *ActionUpdateImageResponseModel) GetTerminatingErrorOk() (*ActionResponseModelTerminatingError, bool)`

GetTerminatingErrorOk returns a tuple with the TerminatingError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatingError

`func (o *ActionUpdateImageResponseModel) SetTerminatingError(v ActionResponseModelTerminatingError)`

SetTerminatingError sets TerminatingError field to given value.

### HasTerminatingError

`func (o *ActionUpdateImageResponseModel) HasTerminatingError() bool`

HasTerminatingError returns a boolean if a field has been set.

### GetNonTerminatingErrors

`func (o *ActionUpdateImageResponseModel) GetNonTerminatingErrors() []ActionError`

GetNonTerminatingErrors returns the NonTerminatingErrors field if non-nil, zero value otherwise.

### GetNonTerminatingErrorsOk

`func (o *ActionUpdateImageResponseModel) GetNonTerminatingErrorsOk() (*[]ActionError, bool)`

GetNonTerminatingErrorsOk returns a tuple with the NonTerminatingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTerminatingErrors

`func (o *ActionUpdateImageResponseModel) SetNonTerminatingErrors(v []ActionError)`

SetNonTerminatingErrors sets NonTerminatingErrors field to given value.

### HasNonTerminatingErrors

`func (o *ActionUpdateImageResponseModel) HasNonTerminatingErrors() bool`

HasNonTerminatingErrors returns a boolean if a field has been set.

### GetMetadata

`func (o *ActionUpdateImageResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionUpdateImageResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionUpdateImageResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ActionUpdateImageResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScopes

`func (o *ActionUpdateImageResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ActionUpdateImageResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ActionUpdateImageResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ActionUpdateImageResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTenants

`func (o *ActionUpdateImageResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ActionUpdateImageResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ActionUpdateImageResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ActionUpdateImageResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetIsRunning

`func (o *ActionUpdateImageResponseModel) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *ActionUpdateImageResponseModel) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *ActionUpdateImageResponseModel) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.

### HasIsRunning

`func (o *ActionUpdateImageResponseModel) HasIsRunning() bool`

HasIsRunning returns a boolean if a field has been set.

### GetMasterImage

`func (o *ActionUpdateImageResponseModel) GetMasterImage() string`

GetMasterImage returns the MasterImage field if non-nil, zero value otherwise.

### GetMasterImageOk

`func (o *ActionUpdateImageResponseModel) GetMasterImageOk() (*string, bool)`

GetMasterImageOk returns a tuple with the MasterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImage

`func (o *ActionUpdateImageResponseModel) SetMasterImage(v string)`

SetMasterImage sets MasterImage field to given value.

### HasMasterImage

`func (o *ActionUpdateImageResponseModel) HasMasterImage() bool`

HasMasterImage returns a boolean if a field has been set.

### GetMachineCreationData

`func (o *ActionUpdateImageResponseModel) GetMachineCreationData() ActionCatalogCreationResponseModelAllOfMachineCreationData`

GetMachineCreationData returns the MachineCreationData field if non-nil, zero value otherwise.

### GetMachineCreationDataOk

`func (o *ActionUpdateImageResponseModel) GetMachineCreationDataOk() (*ActionCatalogCreationResponseModelAllOfMachineCreationData, bool)`

GetMachineCreationDataOk returns a tuple with the MachineCreationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCreationData

`func (o *ActionUpdateImageResponseModel) SetMachineCreationData(v ActionCatalogCreationResponseModelAllOfMachineCreationData)`

SetMachineCreationData sets MachineCreationData field to given value.

### HasMachineCreationData

`func (o *ActionUpdateImageResponseModel) HasMachineCreationData() bool`

HasMachineCreationData returns a boolean if a field has been set.

### GetSnapshot

`func (o *ActionUpdateImageResponseModel) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *ActionUpdateImageResponseModel) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *ActionUpdateImageResponseModel) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *ActionUpdateImageResponseModel) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetMachineRemovalData

`func (o *ActionUpdateImageResponseModel) GetMachineRemovalData() ActionMachineRemovalResponseModelAllOfMachineRemovalData`

GetMachineRemovalData returns the MachineRemovalData field if non-nil, zero value otherwise.

### GetMachineRemovalDataOk

`func (o *ActionUpdateImageResponseModel) GetMachineRemovalDataOk() (*ActionMachineRemovalResponseModelAllOfMachineRemovalData, bool)`

GetMachineRemovalDataOk returns a tuple with the MachineRemovalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineRemovalData

`func (o *ActionUpdateImageResponseModel) SetMachineRemovalData(v ActionMachineRemovalResponseModelAllOfMachineRemovalData)`

SetMachineRemovalData sets MachineRemovalData field to given value.

### HasMachineRemovalData

`func (o *ActionUpdateImageResponseModel) HasMachineRemovalData() bool`

HasMachineRemovalData returns a boolean if a field has been set.

### GetDeleteVirtualMachines

`func (o *ActionUpdateImageResponseModel) GetDeleteVirtualMachines() string`

GetDeleteVirtualMachines returns the DeleteVirtualMachines field if non-nil, zero value otherwise.

### GetDeleteVirtualMachinesOk

`func (o *ActionUpdateImageResponseModel) GetDeleteVirtualMachinesOk() (*string, bool)`

GetDeleteVirtualMachinesOk returns a tuple with the DeleteVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVirtualMachines

`func (o *ActionUpdateImageResponseModel) SetDeleteVirtualMachines(v string)`

SetDeleteVirtualMachines sets DeleteVirtualMachines field to given value.

### HasDeleteVirtualMachines

`func (o *ActionUpdateImageResponseModel) HasDeleteVirtualMachines() bool`

HasDeleteVirtualMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


