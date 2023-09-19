# ActionCatalogCreationResponseModel

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

### NewActionCatalogCreationResponseModel

`func NewActionCatalogCreationResponseModel() *ActionCatalogCreationResponseModel`

NewActionCatalogCreationResponseModel instantiates a new ActionCatalogCreationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionCatalogCreationResponseModelWithDefaults

`func NewActionCatalogCreationResponseModelWithDefaults() *ActionCatalogCreationResponseModel`

NewActionCatalogCreationResponseModelWithDefaults instantiates a new ActionCatalogCreationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *ActionCatalogCreationResponseModel) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ActionCatalogCreationResponseModel) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ActionCatalogCreationResponseModel) SetActionId(v string)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *ActionCatalogCreationResponseModel) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### GetActionTargetName

`func (o *ActionCatalogCreationResponseModel) GetActionTargetName() string`

GetActionTargetName returns the ActionTargetName field if non-nil, zero value otherwise.

### GetActionTargetNameOk

`func (o *ActionCatalogCreationResponseModel) GetActionTargetNameOk() (*string, bool)`

GetActionTargetNameOk returns a tuple with the ActionTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetName

`func (o *ActionCatalogCreationResponseModel) SetActionTargetName(v string)`

SetActionTargetName sets ActionTargetName field to given value.

### HasActionTargetName

`func (o *ActionCatalogCreationResponseModel) HasActionTargetName() bool`

HasActionTargetName returns a boolean if a field has been set.

### SetActionTargetNameNil

`func (o *ActionCatalogCreationResponseModel) SetActionTargetNameNil(b bool)`

 SetActionTargetNameNil sets the value for ActionTargetName to be an explicit nil

### UnsetActionTargetName
`func (o *ActionCatalogCreationResponseModel) UnsetActionTargetName()`

UnsetActionTargetName ensures that no value is present for ActionTargetName, not even an explicit nil
### GetActionTargetUid

`func (o *ActionCatalogCreationResponseModel) GetActionTargetUid() string`

GetActionTargetUid returns the ActionTargetUid field if non-nil, zero value otherwise.

### GetActionTargetUidOk

`func (o *ActionCatalogCreationResponseModel) GetActionTargetUidOk() (*string, bool)`

GetActionTargetUidOk returns a tuple with the ActionTargetUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetUid

`func (o *ActionCatalogCreationResponseModel) SetActionTargetUid(v string)`

SetActionTargetUid sets ActionTargetUid field to given value.

### HasActionTargetUid

`func (o *ActionCatalogCreationResponseModel) HasActionTargetUid() bool`

HasActionTargetUid returns a boolean if a field has been set.

### SetActionTargetUidNil

`func (o *ActionCatalogCreationResponseModel) SetActionTargetUidNil(b bool)`

 SetActionTargetUidNil sets the value for ActionTargetUid to be an explicit nil

### UnsetActionTargetUid
`func (o *ActionCatalogCreationResponseModel) UnsetActionTargetUid()`

UnsetActionTargetUid ensures that no value is present for ActionTargetUid, not even an explicit nil
### GetActionTargetId

`func (o *ActionCatalogCreationResponseModel) GetActionTargetId() string`

GetActionTargetId returns the ActionTargetId field if non-nil, zero value otherwise.

### GetActionTargetIdOk

`func (o *ActionCatalogCreationResponseModel) GetActionTargetIdOk() (*string, bool)`

GetActionTargetIdOk returns a tuple with the ActionTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTargetId

`func (o *ActionCatalogCreationResponseModel) SetActionTargetId(v string)`

SetActionTargetId sets ActionTargetId field to given value.

### HasActionTargetId

`func (o *ActionCatalogCreationResponseModel) HasActionTargetId() bool`

HasActionTargetId returns a boolean if a field has been set.

### SetActionTargetIdNil

`func (o *ActionCatalogCreationResponseModel) SetActionTargetIdNil(b bool)`

 SetActionTargetIdNil sets the value for ActionTargetId to be an explicit nil

### UnsetActionTargetId
`func (o *ActionCatalogCreationResponseModel) UnsetActionTargetId()`

UnsetActionTargetId ensures that no value is present for ActionTargetId, not even an explicit nil
### GetActionType

`func (o *ActionCatalogCreationResponseModel) GetActionType() ActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ActionCatalogCreationResponseModel) GetActionTypeOk() (*ActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ActionCatalogCreationResponseModel) SetActionType(v ActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ActionCatalogCreationResponseModel) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetCreationTime

`func (o *ActionCatalogCreationResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ActionCatalogCreationResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ActionCatalogCreationResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ActionCatalogCreationResponseModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### SetCreationTimeNil

`func (o *ActionCatalogCreationResponseModel) SetCreationTimeNil(b bool)`

 SetCreationTimeNil sets the value for CreationTime to be an explicit nil

### UnsetCreationTime
`func (o *ActionCatalogCreationResponseModel) UnsetCreationTime()`

UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
### GetStartTime

`func (o *ActionCatalogCreationResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ActionCatalogCreationResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ActionCatalogCreationResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ActionCatalogCreationResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ActionCatalogCreationResponseModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ActionCatalogCreationResponseModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetFinishTime

`func (o *ActionCatalogCreationResponseModel) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *ActionCatalogCreationResponseModel) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *ActionCatalogCreationResponseModel) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *ActionCatalogCreationResponseModel) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### SetFinishTimeNil

`func (o *ActionCatalogCreationResponseModel) SetFinishTimeNil(b bool)`

 SetFinishTimeNil sets the value for FinishTime to be an explicit nil

### UnsetFinishTime
`func (o *ActionCatalogCreationResponseModel) UnsetFinishTime()`

UnsetFinishTime ensures that no value is present for FinishTime, not even an explicit nil
### GetState

`func (o *ActionCatalogCreationResponseModel) GetState() ActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActionCatalogCreationResponseModel) GetStateOk() (*ActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActionCatalogCreationResponseModel) SetState(v ActionState)`

SetState sets State field to given value.

### HasState

`func (o *ActionCatalogCreationResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetErrorState

`func (o *ActionCatalogCreationResponseModel) GetErrorState() ActionErrorStatus`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *ActionCatalogCreationResponseModel) GetErrorStateOk() (*ActionErrorStatus, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *ActionCatalogCreationResponseModel) SetErrorState(v ActionErrorStatus)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *ActionCatalogCreationResponseModel) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetProgress

`func (o *ActionCatalogCreationResponseModel) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ActionCatalogCreationResponseModel) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ActionCatalogCreationResponseModel) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ActionCatalogCreationResponseModel) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressMessage

`func (o *ActionCatalogCreationResponseModel) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *ActionCatalogCreationResponseModel) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *ActionCatalogCreationResponseModel) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *ActionCatalogCreationResponseModel) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### SetProgressMessageNil

`func (o *ActionCatalogCreationResponseModel) SetProgressMessageNil(b bool)`

 SetProgressMessageNil sets the value for ProgressMessage to be an explicit nil

### UnsetProgressMessage
`func (o *ActionCatalogCreationResponseModel) UnsetProgressMessage()`

UnsetProgressMessage ensures that no value is present for ProgressMessage, not even an explicit nil
### GetTerminatingError

`func (o *ActionCatalogCreationResponseModel) GetTerminatingError() ActionError`

GetTerminatingError returns the TerminatingError field if non-nil, zero value otherwise.

### GetTerminatingErrorOk

`func (o *ActionCatalogCreationResponseModel) GetTerminatingErrorOk() (*ActionError, bool)`

GetTerminatingErrorOk returns a tuple with the TerminatingError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatingError

`func (o *ActionCatalogCreationResponseModel) SetTerminatingError(v ActionError)`

SetTerminatingError sets TerminatingError field to given value.

### HasTerminatingError

`func (o *ActionCatalogCreationResponseModel) HasTerminatingError() bool`

HasTerminatingError returns a boolean if a field has been set.

### GetNonTerminatingErrors

`func (o *ActionCatalogCreationResponseModel) GetNonTerminatingErrors() []ActionError`

GetNonTerminatingErrors returns the NonTerminatingErrors field if non-nil, zero value otherwise.

### GetNonTerminatingErrorsOk

`func (o *ActionCatalogCreationResponseModel) GetNonTerminatingErrorsOk() (*[]ActionError, bool)`

GetNonTerminatingErrorsOk returns a tuple with the NonTerminatingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTerminatingErrors

`func (o *ActionCatalogCreationResponseModel) SetNonTerminatingErrors(v []ActionError)`

SetNonTerminatingErrors sets NonTerminatingErrors field to given value.

### HasNonTerminatingErrors

`func (o *ActionCatalogCreationResponseModel) HasNonTerminatingErrors() bool`

HasNonTerminatingErrors returns a boolean if a field has been set.

### SetNonTerminatingErrorsNil

`func (o *ActionCatalogCreationResponseModel) SetNonTerminatingErrorsNil(b bool)`

 SetNonTerminatingErrorsNil sets the value for NonTerminatingErrors to be an explicit nil

### UnsetNonTerminatingErrors
`func (o *ActionCatalogCreationResponseModel) UnsetNonTerminatingErrors()`

UnsetNonTerminatingErrors ensures that no value is present for NonTerminatingErrors, not even an explicit nil
### GetMetadata

`func (o *ActionCatalogCreationResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionCatalogCreationResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionCatalogCreationResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ActionCatalogCreationResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ActionCatalogCreationResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ActionCatalogCreationResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetScopes

`func (o *ActionCatalogCreationResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ActionCatalogCreationResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ActionCatalogCreationResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ActionCatalogCreationResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ActionCatalogCreationResponseModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ActionCatalogCreationResponseModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *ActionCatalogCreationResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ActionCatalogCreationResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ActionCatalogCreationResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ActionCatalogCreationResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ActionCatalogCreationResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ActionCatalogCreationResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetIsRunning

`func (o *ActionCatalogCreationResponseModel) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *ActionCatalogCreationResponseModel) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *ActionCatalogCreationResponseModel) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.

### HasIsRunning

`func (o *ActionCatalogCreationResponseModel) HasIsRunning() bool`

HasIsRunning returns a boolean if a field has been set.

### GetMasterImage

`func (o *ActionCatalogCreationResponseModel) GetMasterImage() string`

GetMasterImage returns the MasterImage field if non-nil, zero value otherwise.

### GetMasterImageOk

`func (o *ActionCatalogCreationResponseModel) GetMasterImageOk() (*string, bool)`

GetMasterImageOk returns a tuple with the MasterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImage

`func (o *ActionCatalogCreationResponseModel) SetMasterImage(v string)`

SetMasterImage sets MasterImage field to given value.

### HasMasterImage

`func (o *ActionCatalogCreationResponseModel) HasMasterImage() bool`

HasMasterImage returns a boolean if a field has been set.

### SetMasterImageNil

`func (o *ActionCatalogCreationResponseModel) SetMasterImageNil(b bool)`

 SetMasterImageNil sets the value for MasterImage to be an explicit nil

### UnsetMasterImage
`func (o *ActionCatalogCreationResponseModel) UnsetMasterImage()`

UnsetMasterImage ensures that no value is present for MasterImage, not even an explicit nil
### GetMachineCreationData

`func (o *ActionCatalogCreationResponseModel) GetMachineCreationData() ActionMachineCreationDetailsResponseModel`

GetMachineCreationData returns the MachineCreationData field if non-nil, zero value otherwise.

### GetMachineCreationDataOk

`func (o *ActionCatalogCreationResponseModel) GetMachineCreationDataOk() (*ActionMachineCreationDetailsResponseModel, bool)`

GetMachineCreationDataOk returns a tuple with the MachineCreationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCreationData

`func (o *ActionCatalogCreationResponseModel) SetMachineCreationData(v ActionMachineCreationDetailsResponseModel)`

SetMachineCreationData sets MachineCreationData field to given value.

### HasMachineCreationData

`func (o *ActionCatalogCreationResponseModel) HasMachineCreationData() bool`

HasMachineCreationData returns a boolean if a field has been set.

### GetSnapshot

`func (o *ActionCatalogCreationResponseModel) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *ActionCatalogCreationResponseModel) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *ActionCatalogCreationResponseModel) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *ActionCatalogCreationResponseModel) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### SetSnapshotNil

`func (o *ActionCatalogCreationResponseModel) SetSnapshotNil(b bool)`

 SetSnapshotNil sets the value for Snapshot to be an explicit nil

### UnsetSnapshot
`func (o *ActionCatalogCreationResponseModel) UnsetSnapshot()`

UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
### GetMachineRemovalData

`func (o *ActionCatalogCreationResponseModel) GetMachineRemovalData() ActionMachineRemovalDetailsResponseModel`

GetMachineRemovalData returns the MachineRemovalData field if non-nil, zero value otherwise.

### GetMachineRemovalDataOk

`func (o *ActionCatalogCreationResponseModel) GetMachineRemovalDataOk() (*ActionMachineRemovalDetailsResponseModel, bool)`

GetMachineRemovalDataOk returns a tuple with the MachineRemovalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineRemovalData

`func (o *ActionCatalogCreationResponseModel) SetMachineRemovalData(v ActionMachineRemovalDetailsResponseModel)`

SetMachineRemovalData sets MachineRemovalData field to given value.

### HasMachineRemovalData

`func (o *ActionCatalogCreationResponseModel) HasMachineRemovalData() bool`

HasMachineRemovalData returns a boolean if a field has been set.

### GetDeleteVirtualMachines

`func (o *ActionCatalogCreationResponseModel) GetDeleteVirtualMachines() string`

GetDeleteVirtualMachines returns the DeleteVirtualMachines field if non-nil, zero value otherwise.

### GetDeleteVirtualMachinesOk

`func (o *ActionCatalogCreationResponseModel) GetDeleteVirtualMachinesOk() (*string, bool)`

GetDeleteVirtualMachinesOk returns a tuple with the DeleteVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVirtualMachines

`func (o *ActionCatalogCreationResponseModel) SetDeleteVirtualMachines(v string)`

SetDeleteVirtualMachines sets DeleteVirtualMachines field to given value.

### HasDeleteVirtualMachines

`func (o *ActionCatalogCreationResponseModel) HasDeleteVirtualMachines() bool`

HasDeleteVirtualMachines returns a boolean if a field has been set.

### SetDeleteVirtualMachinesNil

`func (o *ActionCatalogCreationResponseModel) SetDeleteVirtualMachinesNil(b bool)`

 SetDeleteVirtualMachinesNil sets the value for DeleteVirtualMachines to be an explicit nil

### UnsetDeleteVirtualMachines
`func (o *ActionCatalogCreationResponseModel) UnsetDeleteVirtualMachines()`

UnsetDeleteVirtualMachines ensures that no value is present for DeleteVirtualMachines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


