# ActionMachineCreationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionId** | Pointer to **string** | The action id. | [optional] 
**ActionTargetName** | Pointer to **NullableString** | The action target name, it&#39;s the related catalog name. | [optional] 
**ActionTargetUid** | Pointer to **NullableString** | The action target uid, it&#39;s the related catalog uid. | [optional] 
**ActionTargetId** | Pointer to **NullableString** | The action target id, it&#39;s related on catalog id. | [optional] 
**ActionType** | Pointer to [**ActionType**](ActionType.md) |  | [optional] 
**CreationTime** | Pointer to **NullableString** | The action creation time. | [optional] 
**StartTime** | Pointer to **NullableString** | The action start time. | [optional] 
**FinishTime** | Pointer to **NullableString** | The action finish time. | [optional] 
**State** | Pointer to [**ActionState**](ActionState.md) |  | [optional] 
**ErrorState** | Pointer to [**ActionErrorStatus**](ActionErrorStatus.md) |  | [optional] 
**Progress** | Pointer to **float64** | The progress of the action. | [optional] 
**ProgressMessage** | Pointer to **NullableString** | The progress message. | [optional] 
**TerminatingError** | Pointer to [**ActionError**](ActionError.md) |  | [optional] 
**NonTerminatingErrors** | Pointer to [**[]ActionError**](ActionError.md) | The non terminating errors. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of the action. | [optional] 
**Scopes** | Pointer to [**[]ScopeResponseModel**](ScopeResponseModel.md) | The scope of the catalog. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the hypervisor is assigned to.  If &#x60;null&#x60;, the hypervisor is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**IsRunning** | Pointer to **bool** | If the task is running. | [optional] 
**MasterImage** | Pointer to **NullableString** | The master image. | [optional] 
**MachineCreationData** | Pointer to [**ActionMachineCreationDetailsResponseModel**](ActionMachineCreationDetailsResponseModel.md) |  | [optional] 
**Snapshot** | Pointer to **NullableString** | The snapshot. | [optional] 
**MachineRemovalData** | Pointer to [**ActionMachineRemovalDetailsResponseModel**](ActionMachineRemovalDetailsResponseModel.md) |  | [optional] 
**DeleteVirtualMachines** | Pointer to **NullableString** | The delete virtual machines.  | [optional] 

## Methods

### NewActionMachineCreationResponseModel

`func NewActionMachineCreationResponseModel() *ActionMachineCreationResponseModel`

NewActionMachineCreationResponseModel instantiates a new ActionMachineCreationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionMachineCreationResponseModelWithDefaults

`func NewActionMachineCreationResponseModelWithDefaults() *ActionMachineCreationResponseModel`

NewActionMachineCreationResponseModelWithDefaults instantiates a new ActionMachineCreationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *ActionMachineCreationResponseModel) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ActionMachineCreationResponseModel) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ActionMachineCreationResponseModel) SetActionId(v string)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *ActionMachineCreationResponseModel) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### GetActionTargetName

`func (o *ActionMachineCreationResponseModel) GetActionTargetName() string`

GetActionTargetName returns the ActionTargetName field if non-nil, zero value otherwise.

### GetActionTargetNameOk

`func (o *ActionMachineCreationResponseModel) GetActionTargetNameOk() (*string, bool)`

GetActionTargetNameOk returns a tuple with the ActionTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetName

`func (o *ActionMachineCreationResponseModel) SetActionTargetName(v string)`

SetActionTargetName sets ActionTargetName field to given value.

### HasActionTargetName

`func (o *ActionMachineCreationResponseModel) HasActionTargetName() bool`

HasActionTargetName returns a boolean if a field has been set.

### SetActionTargetNameNil

`func (o *ActionMachineCreationResponseModel) SetActionTargetNameNil(b bool)`

 SetActionTargetNameNil sets the value for ActionTargetName to be an explicit nil

### UnsetActionTargetName
`func (o *ActionMachineCreationResponseModel) UnsetActionTargetName()`

UnsetActionTargetName ensures that no value is present for ActionTargetName, not even an explicit nil
### GetActionTargetUid

`func (o *ActionMachineCreationResponseModel) GetActionTargetUid() string`

GetActionTargetUid returns the ActionTargetUid field if non-nil, zero value otherwise.

### GetActionTargetUidOk

`func (o *ActionMachineCreationResponseModel) GetActionTargetUidOk() (*string, bool)`

GetActionTargetUidOk returns a tuple with the ActionTargetUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetUid

`func (o *ActionMachineCreationResponseModel) SetActionTargetUid(v string)`

SetActionTargetUid sets ActionTargetUid field to given value.

### HasActionTargetUid

`func (o *ActionMachineCreationResponseModel) HasActionTargetUid() bool`

HasActionTargetUid returns a boolean if a field has been set.

### SetActionTargetUidNil

`func (o *ActionMachineCreationResponseModel) SetActionTargetUidNil(b bool)`

 SetActionTargetUidNil sets the value for ActionTargetUid to be an explicit nil

### UnsetActionTargetUid
`func (o *ActionMachineCreationResponseModel) UnsetActionTargetUid()`

UnsetActionTargetUid ensures that no value is present for ActionTargetUid, not even an explicit nil
### GetActionTargetId

`func (o *ActionMachineCreationResponseModel) GetActionTargetId() string`

GetActionTargetId returns the ActionTargetId field if non-nil, zero value otherwise.

### GetActionTargetIdOk

`func (o *ActionMachineCreationResponseModel) GetActionTargetIdOk() (*string, bool)`

GetActionTargetIdOk returns a tuple with the ActionTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetId

`func (o *ActionMachineCreationResponseModel) SetActionTargetId(v string)`

SetActionTargetId sets ActionTargetId field to given value.

### HasActionTargetId

`func (o *ActionMachineCreationResponseModel) HasActionTargetId() bool`

HasActionTargetId returns a boolean if a field has been set.

### SetActionTargetIdNil

`func (o *ActionMachineCreationResponseModel) SetActionTargetIdNil(b bool)`

 SetActionTargetIdNil sets the value for ActionTargetId to be an explicit nil

### UnsetActionTargetId
`func (o *ActionMachineCreationResponseModel) UnsetActionTargetId()`

UnsetActionTargetId ensures that no value is present for ActionTargetId, not even an explicit nil
### GetActionType

`func (o *ActionMachineCreationResponseModel) GetActionType() ActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ActionMachineCreationResponseModel) GetActionTypeOk() (*ActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ActionMachineCreationResponseModel) SetActionType(v ActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ActionMachineCreationResponseModel) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetCreationTime

`func (o *ActionMachineCreationResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ActionMachineCreationResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ActionMachineCreationResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ActionMachineCreationResponseModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### SetCreationTimeNil

`func (o *ActionMachineCreationResponseModel) SetCreationTimeNil(b bool)`

 SetCreationTimeNil sets the value for CreationTime to be an explicit nil

### UnsetCreationTime
`func (o *ActionMachineCreationResponseModel) UnsetCreationTime()`

UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
### GetStartTime

`func (o *ActionMachineCreationResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ActionMachineCreationResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ActionMachineCreationResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ActionMachineCreationResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ActionMachineCreationResponseModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ActionMachineCreationResponseModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetFinishTime

`func (o *ActionMachineCreationResponseModel) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *ActionMachineCreationResponseModel) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *ActionMachineCreationResponseModel) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *ActionMachineCreationResponseModel) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### SetFinishTimeNil

`func (o *ActionMachineCreationResponseModel) SetFinishTimeNil(b bool)`

 SetFinishTimeNil sets the value for FinishTime to be an explicit nil

### UnsetFinishTime
`func (o *ActionMachineCreationResponseModel) UnsetFinishTime()`

UnsetFinishTime ensures that no value is present for FinishTime, not even an explicit nil
### GetState

`func (o *ActionMachineCreationResponseModel) GetState() ActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActionMachineCreationResponseModel) GetStateOk() (*ActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActionMachineCreationResponseModel) SetState(v ActionState)`

SetState sets State field to given value.

### HasState

`func (o *ActionMachineCreationResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetErrorState

`func (o *ActionMachineCreationResponseModel) GetErrorState() ActionErrorStatus`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *ActionMachineCreationResponseModel) GetErrorStateOk() (*ActionErrorStatus, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *ActionMachineCreationResponseModel) SetErrorState(v ActionErrorStatus)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *ActionMachineCreationResponseModel) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetProgress

`func (o *ActionMachineCreationResponseModel) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ActionMachineCreationResponseModel) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ActionMachineCreationResponseModel) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ActionMachineCreationResponseModel) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressMessage

`func (o *ActionMachineCreationResponseModel) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *ActionMachineCreationResponseModel) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *ActionMachineCreationResponseModel) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *ActionMachineCreationResponseModel) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### SetProgressMessageNil

`func (o *ActionMachineCreationResponseModel) SetProgressMessageNil(b bool)`

 SetProgressMessageNil sets the value for ProgressMessage to be an explicit nil

### UnsetProgressMessage
`func (o *ActionMachineCreationResponseModel) UnsetProgressMessage()`

UnsetProgressMessage ensures that no value is present for ProgressMessage, not even an explicit nil
### GetTerminatingError

`func (o *ActionMachineCreationResponseModel) GetTerminatingError() ActionError`

GetTerminatingError returns the TerminatingError field if non-nil, zero value otherwise.

### GetTerminatingErrorOk

`func (o *ActionMachineCreationResponseModel) GetTerminatingErrorOk() (*ActionError, bool)`

GetTerminatingErrorOk returns a tuple with the TerminatingError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatingError

`func (o *ActionMachineCreationResponseModel) SetTerminatingError(v ActionError)`

SetTerminatingError sets TerminatingError field to given value.

### HasTerminatingError

`func (o *ActionMachineCreationResponseModel) HasTerminatingError() bool`

HasTerminatingError returns a boolean if a field has been set.

### GetNonTerminatingErrors

`func (o *ActionMachineCreationResponseModel) GetNonTerminatingErrors() []ActionError`

GetNonTerminatingErrors returns the NonTerminatingErrors field if non-nil, zero value otherwise.

### GetNonTerminatingErrorsOk

`func (o *ActionMachineCreationResponseModel) GetNonTerminatingErrorsOk() (*[]ActionError, bool)`

GetNonTerminatingErrorsOk returns a tuple with the NonTerminatingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTerminatingErrors

`func (o *ActionMachineCreationResponseModel) SetNonTerminatingErrors(v []ActionError)`

SetNonTerminatingErrors sets NonTerminatingErrors field to given value.

### HasNonTerminatingErrors

`func (o *ActionMachineCreationResponseModel) HasNonTerminatingErrors() bool`

HasNonTerminatingErrors returns a boolean if a field has been set.

### SetNonTerminatingErrorsNil

`func (o *ActionMachineCreationResponseModel) SetNonTerminatingErrorsNil(b bool)`

 SetNonTerminatingErrorsNil sets the value for NonTerminatingErrors to be an explicit nil

### UnsetNonTerminatingErrors
`func (o *ActionMachineCreationResponseModel) UnsetNonTerminatingErrors()`

UnsetNonTerminatingErrors ensures that no value is present for NonTerminatingErrors, not even an explicit nil
### GetMetadata

`func (o *ActionMachineCreationResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionMachineCreationResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionMachineCreationResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ActionMachineCreationResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ActionMachineCreationResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ActionMachineCreationResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetScopes

`func (o *ActionMachineCreationResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ActionMachineCreationResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ActionMachineCreationResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ActionMachineCreationResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ActionMachineCreationResponseModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ActionMachineCreationResponseModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *ActionMachineCreationResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ActionMachineCreationResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ActionMachineCreationResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ActionMachineCreationResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ActionMachineCreationResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ActionMachineCreationResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetIsRunning

`func (o *ActionMachineCreationResponseModel) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *ActionMachineCreationResponseModel) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *ActionMachineCreationResponseModel) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.

### HasIsRunning

`func (o *ActionMachineCreationResponseModel) HasIsRunning() bool`

HasIsRunning returns a boolean if a field has been set.

### GetMasterImage

`func (o *ActionMachineCreationResponseModel) GetMasterImage() string`

GetMasterImage returns the MasterImage field if non-nil, zero value otherwise.

### GetMasterImageOk

`func (o *ActionMachineCreationResponseModel) GetMasterImageOk() (*string, bool)`

GetMasterImageOk returns a tuple with the MasterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImage

`func (o *ActionMachineCreationResponseModel) SetMasterImage(v string)`

SetMasterImage sets MasterImage field to given value.

### HasMasterImage

`func (o *ActionMachineCreationResponseModel) HasMasterImage() bool`

HasMasterImage returns a boolean if a field has been set.

### SetMasterImageNil

`func (o *ActionMachineCreationResponseModel) SetMasterImageNil(b bool)`

 SetMasterImageNil sets the value for MasterImage to be an explicit nil

### UnsetMasterImage
`func (o *ActionMachineCreationResponseModel) UnsetMasterImage()`

UnsetMasterImage ensures that no value is present for MasterImage, not even an explicit nil
### GetMachineCreationData

`func (o *ActionMachineCreationResponseModel) GetMachineCreationData() ActionMachineCreationDetailsResponseModel`

GetMachineCreationData returns the MachineCreationData field if non-nil, zero value otherwise.

### GetMachineCreationDataOk

`func (o *ActionMachineCreationResponseModel) GetMachineCreationDataOk() (*ActionMachineCreationDetailsResponseModel, bool)`

GetMachineCreationDataOk returns a tuple with the MachineCreationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCreationData

`func (o *ActionMachineCreationResponseModel) SetMachineCreationData(v ActionMachineCreationDetailsResponseModel)`

SetMachineCreationData sets MachineCreationData field to given value.

### HasMachineCreationData

`func (o *ActionMachineCreationResponseModel) HasMachineCreationData() bool`

HasMachineCreationData returns a boolean if a field has been set.

### GetSnapshot

`func (o *ActionMachineCreationResponseModel) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *ActionMachineCreationResponseModel) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *ActionMachineCreationResponseModel) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *ActionMachineCreationResponseModel) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### SetSnapshotNil

`func (o *ActionMachineCreationResponseModel) SetSnapshotNil(b bool)`

 SetSnapshotNil sets the value for Snapshot to be an explicit nil

### UnsetSnapshot
`func (o *ActionMachineCreationResponseModel) UnsetSnapshot()`

UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
### GetMachineRemovalData

`func (o *ActionMachineCreationResponseModel) GetMachineRemovalData() ActionMachineRemovalDetailsResponseModel`

GetMachineRemovalData returns the MachineRemovalData field if non-nil, zero value otherwise.

### GetMachineRemovalDataOk

`func (o *ActionMachineCreationResponseModel) GetMachineRemovalDataOk() (*ActionMachineRemovalDetailsResponseModel, bool)`

GetMachineRemovalDataOk returns a tuple with the MachineRemovalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineRemovalData

`func (o *ActionMachineCreationResponseModel) SetMachineRemovalData(v ActionMachineRemovalDetailsResponseModel)`

SetMachineRemovalData sets MachineRemovalData field to given value.

### HasMachineRemovalData

`func (o *ActionMachineCreationResponseModel) HasMachineRemovalData() bool`

HasMachineRemovalData returns a boolean if a field has been set.

### GetDeleteVirtualMachines

`func (o *ActionMachineCreationResponseModel) GetDeleteVirtualMachines() string`

GetDeleteVirtualMachines returns the DeleteVirtualMachines field if non-nil, zero value otherwise.

### GetDeleteVirtualMachinesOk

`func (o *ActionMachineCreationResponseModel) GetDeleteVirtualMachinesOk() (*string, bool)`

GetDeleteVirtualMachinesOk returns a tuple with the DeleteVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVirtualMachines

`func (o *ActionMachineCreationResponseModel) SetDeleteVirtualMachines(v string)`

SetDeleteVirtualMachines sets DeleteVirtualMachines field to given value.

### HasDeleteVirtualMachines

`func (o *ActionMachineCreationResponseModel) HasDeleteVirtualMachines() bool`

HasDeleteVirtualMachines returns a boolean if a field has been set.

### SetDeleteVirtualMachinesNil

`func (o *ActionMachineCreationResponseModel) SetDeleteVirtualMachinesNil(b bool)`

 SetDeleteVirtualMachinesNil sets the value for DeleteVirtualMachines to be an explicit nil

### UnsetDeleteVirtualMachines
`func (o *ActionMachineCreationResponseModel) UnsetDeleteVirtualMachines()`

UnsetDeleteVirtualMachines ensures that no value is present for DeleteVirtualMachines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


