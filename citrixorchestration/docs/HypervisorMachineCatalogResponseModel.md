# HypervisorMachineCatalogResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Simple administrative name of catalog within parent admin folder (if any). This property is not guaranteed unique across all catalogs. | [optional] 
**FullName** | Pointer to **NullableString** | Unique administrative name of catalog. | [optional] 
**Id** | Pointer to **NullableString** | Id of the machine catalog. | [optional] 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED.  Use Id. | [optional] 
**AllocationType** | Pointer to [**AllocationType**](AllocationType.md) |  | [optional] 
**AssignedCount** | Pointer to **NullableInt32** | The number of assigned machines (machines that have been assigned to a user/users or a client name/address). | [optional] 
**AvailableAssignedCount** | Pointer to **NullableInt32** | The number of available machines (not in a delivery group), that are also assigned to users. | [optional] 
**AvailableCount** | Pointer to **int32** | The number of available machines (those not in any delivery group). | [optional] 
**AvailableUnassignedCount** | Pointer to **NullableInt32** | The number of available machines (those not in any delivery group) that are not assigned to users. | [optional] 
**Description** | Pointer to **NullableString** | Description of the machine catalog. | [optional] 
**IsPowerManaged** | Pointer to **bool** | Indicates whether the machines in the catalog are power-managed. | [optional] 
**IsRemotePC** | Pointer to **bool** | Indicates whether or not the catalog is a RemotePC catalog. Remote PC catalogs automatically configure appropriate machines without the need for manual configuration. CHANGE: was public bool RemotePC { get; set; } | [optional] 
**JobsInProgress** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | List of jobs currently in progress that affect the machine catalog. | [optional] 
**MachineType** | Pointer to [**MachineType**](MachineType.md) |  | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of machine catalog. | [optional] 
**MinimumFunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**HasBeenPromoted** | Pointer to **bool** | Whether the machine catalog was previously promoted from a lower MinimumFunctionalLevel. | [optional] 
**HasBeenPromotedFrom** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**CanRollbackVMImage** | Pointer to **bool** | Whether the machine catalog can roll back VM image. | [optional] 
**CanRecreateCatalog** | Pointer to **bool** | Whether the machine catalog can recreate. | [optional] 
**PersistChanges** | Pointer to [**PersistChanges**](PersistChanges.md) |  | [optional] 
**ProvisioningScheme** | Pointer to [**ProvisioningSchemeResponseModel**](ProvisioningSchemeResponseModel.md) |  | [optional] 
**ProvisioningType** | Pointer to [**ProvisioningType**](ProvisioningType.md) |  | [optional] 
**ProvisioningProgress** | Pointer to [**ProvisioningProgressResponseModel**](ProvisioningProgressResponseModel.md) |  | [optional] 
**PvsAddress** | Pointer to **NullableString** | IP address of the PVS server to be used. This only applies if the ProvisioningType is . | [optional] 
**PvsDomain** | Pointer to **NullableString** | The domain of the PVS server to be used. | [optional] 
**RemotePCEnrollmentScopes** | Pointer to [**[]RemotePCEnrollmentScopeResponseModel**](RemotePCEnrollmentScopeResponseModel.md) | List of one or more remote PC enrollment scopes. | [optional] 
**Scopes** | Pointer to [**[]ScopeResponseModel**](ScopeResponseModel.md) | Administrative scopes which the machine catalog is part of. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the machine catalog is assigned to.  If &#x60;null&#x60;, the machine catalog is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) |  | [optional] 
**SharingKind** | Pointer to [**SharingKind**](SharingKind.md) |  | [optional] 
**TotalCount** | Pointer to **int32** | The total number of machines in the catalog. | [optional] 
**IsBroken** | Pointer to **bool** | Whether the machine catalog is currently in a \&quot;Broken\&quot; state. | [optional] 
**IsMasterImageAssociated** | Pointer to **NullableBool** | Whether the machine catalog is associated with a master image. | [optional] 
**ImageUpdateStatus** | Pointer to [**ImageUpdateStatus**](ImageUpdateStatus.md) |  | [optional] 
**Errors** | Pointer to **[]string** | Gets the Errors of machines in this catalog | [optional] 
**Warnings** | Pointer to [**[]MachineCatalogWarningResponseModel**](MachineCatalogWarningResponseModel.md) | List of warnings that are currently active on the machine catalog, if any.  If there are no warnings this will not be specified. | [optional] 
**UnassignedCount** | Pointer to **int32** | The number of unassigned machines (machines not assigned to users). | [optional] 
**UsedCount** | Pointer to **int32** | The number of machines in the catalog that are in a delivery group. | [optional] 
**AvailableCountOfSuspend** | Pointer to **NullableInt32** | The number of available suspend-capable machines (those not in any delivery group). | [optional] 
**AvailableAssignedCountOfSuspend** | Pointer to **NullableInt32** | The number of available suspend-capable machines (not in a delivery group), that are also assigned to users. | [optional] 
**UpgradeInfo** | Pointer to [**MachineCatalogUpgradeInfo**](MachineCatalogUpgradeInfo.md) |  | [optional] 
**Zone** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**AdminFolder** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**HypervisorVMTagging** | Pointer to **NullableBool** | Indicates that assigned VMs from this catalog will carry a hypervisor-level tag. | [optional] 
**NumProvisionedMachines** | Pointer to **int32** | Number of machines from the catalog that are provisioned on the hypervisor or resource pool. | [optional] 

## Methods

### NewHypervisorMachineCatalogResponseModel

`func NewHypervisorMachineCatalogResponseModel() *HypervisorMachineCatalogResponseModel`

NewHypervisorMachineCatalogResponseModel instantiates a new HypervisorMachineCatalogResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorMachineCatalogResponseModelWithDefaults

`func NewHypervisorMachineCatalogResponseModelWithDefaults() *HypervisorMachineCatalogResponseModel`

NewHypervisorMachineCatalogResponseModelWithDefaults instantiates a new HypervisorMachineCatalogResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HypervisorMachineCatalogResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorMachineCatalogResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorMachineCatalogResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorMachineCatalogResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorMachineCatalogResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorMachineCatalogResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *HypervisorMachineCatalogResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *HypervisorMachineCatalogResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *HypervisorMachineCatalogResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *HypervisorMachineCatalogResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *HypervisorMachineCatalogResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *HypervisorMachineCatalogResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetId

`func (o *HypervisorMachineCatalogResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorMachineCatalogResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorMachineCatalogResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorMachineCatalogResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorMachineCatalogResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorMachineCatalogResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *HypervisorMachineCatalogResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *HypervisorMachineCatalogResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *HypervisorMachineCatalogResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *HypervisorMachineCatalogResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetAllocationType

`func (o *HypervisorMachineCatalogResponseModel) GetAllocationType() AllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *HypervisorMachineCatalogResponseModel) GetAllocationTypeOk() (*AllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *HypervisorMachineCatalogResponseModel) SetAllocationType(v AllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *HypervisorMachineCatalogResponseModel) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetAssignedCount

`func (o *HypervisorMachineCatalogResponseModel) GetAssignedCount() int32`

GetAssignedCount returns the AssignedCount field if non-nil, zero value otherwise.

### GetAssignedCountOk

`func (o *HypervisorMachineCatalogResponseModel) GetAssignedCountOk() (*int32, bool)`

GetAssignedCountOk returns a tuple with the AssignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCount

`func (o *HypervisorMachineCatalogResponseModel) SetAssignedCount(v int32)`

SetAssignedCount sets AssignedCount field to given value.

### HasAssignedCount

`func (o *HypervisorMachineCatalogResponseModel) HasAssignedCount() bool`

HasAssignedCount returns a boolean if a field has been set.

### SetAssignedCountNil

`func (o *HypervisorMachineCatalogResponseModel) SetAssignedCountNil(b bool)`

 SetAssignedCountNil sets the value for AssignedCount to be an explicit nil

### UnsetAssignedCount
`func (o *HypervisorMachineCatalogResponseModel) UnsetAssignedCount()`

UnsetAssignedCount ensures that no value is present for AssignedCount, not even an explicit nil
### GetAvailableAssignedCount

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableAssignedCount() int32`

GetAvailableAssignedCount returns the AvailableAssignedCount field if non-nil, zero value otherwise.

### GetAvailableAssignedCountOk

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableAssignedCountOk() (*int32, bool)`

GetAvailableAssignedCountOk returns a tuple with the AvailableAssignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAssignedCount

`func (o *HypervisorMachineCatalogResponseModel) SetAvailableAssignedCount(v int32)`

SetAvailableAssignedCount sets AvailableAssignedCount field to given value.

### HasAvailableAssignedCount

`func (o *HypervisorMachineCatalogResponseModel) HasAvailableAssignedCount() bool`

HasAvailableAssignedCount returns a boolean if a field has been set.

### SetAvailableAssignedCountNil

`func (o *HypervisorMachineCatalogResponseModel) SetAvailableAssignedCountNil(b bool)`

 SetAvailableAssignedCountNil sets the value for AvailableAssignedCount to be an explicit nil

### UnsetAvailableAssignedCount
`func (o *HypervisorMachineCatalogResponseModel) UnsetAvailableAssignedCount()`

UnsetAvailableAssignedCount ensures that no value is present for AvailableAssignedCount, not even an explicit nil
### GetAvailableCount

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableCount() int32`

GetAvailableCount returns the AvailableCount field if non-nil, zero value otherwise.

### GetAvailableCountOk

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableCountOk() (*int32, bool)`

GetAvailableCountOk returns a tuple with the AvailableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCount

`func (o *HypervisorMachineCatalogResponseModel) SetAvailableCount(v int32)`

SetAvailableCount sets AvailableCount field to given value.

### HasAvailableCount

`func (o *HypervisorMachineCatalogResponseModel) HasAvailableCount() bool`

HasAvailableCount returns a boolean if a field has been set.

### GetAvailableUnassignedCount

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableUnassignedCount() int32`

GetAvailableUnassignedCount returns the AvailableUnassignedCount field if non-nil, zero value otherwise.

### GetAvailableUnassignedCountOk

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableUnassignedCountOk() (*int32, bool)`

GetAvailableUnassignedCountOk returns a tuple with the AvailableUnassignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUnassignedCount

`func (o *HypervisorMachineCatalogResponseModel) SetAvailableUnassignedCount(v int32)`

SetAvailableUnassignedCount sets AvailableUnassignedCount field to given value.

### HasAvailableUnassignedCount

`func (o *HypervisorMachineCatalogResponseModel) HasAvailableUnassignedCount() bool`

HasAvailableUnassignedCount returns a boolean if a field has been set.

### SetAvailableUnassignedCountNil

`func (o *HypervisorMachineCatalogResponseModel) SetAvailableUnassignedCountNil(b bool)`

 SetAvailableUnassignedCountNil sets the value for AvailableUnassignedCount to be an explicit nil

### UnsetAvailableUnassignedCount
`func (o *HypervisorMachineCatalogResponseModel) UnsetAvailableUnassignedCount()`

UnsetAvailableUnassignedCount ensures that no value is present for AvailableUnassignedCount, not even an explicit nil
### GetDescription

`func (o *HypervisorMachineCatalogResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HypervisorMachineCatalogResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HypervisorMachineCatalogResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HypervisorMachineCatalogResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HypervisorMachineCatalogResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HypervisorMachineCatalogResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsPowerManaged

`func (o *HypervisorMachineCatalogResponseModel) GetIsPowerManaged() bool`

GetIsPowerManaged returns the IsPowerManaged field if non-nil, zero value otherwise.

### GetIsPowerManagedOk

`func (o *HypervisorMachineCatalogResponseModel) GetIsPowerManagedOk() (*bool, bool)`

GetIsPowerManagedOk returns a tuple with the IsPowerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPowerManaged

`func (o *HypervisorMachineCatalogResponseModel) SetIsPowerManaged(v bool)`

SetIsPowerManaged sets IsPowerManaged field to given value.

### HasIsPowerManaged

`func (o *HypervisorMachineCatalogResponseModel) HasIsPowerManaged() bool`

HasIsPowerManaged returns a boolean if a field has been set.

### GetIsRemotePC

`func (o *HypervisorMachineCatalogResponseModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *HypervisorMachineCatalogResponseModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *HypervisorMachineCatalogResponseModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.

### HasIsRemotePC

`func (o *HypervisorMachineCatalogResponseModel) HasIsRemotePC() bool`

HasIsRemotePC returns a boolean if a field has been set.

### GetJobsInProgress

`func (o *HypervisorMachineCatalogResponseModel) GetJobsInProgress() []RefResponseModel`

GetJobsInProgress returns the JobsInProgress field if non-nil, zero value otherwise.

### GetJobsInProgressOk

`func (o *HypervisorMachineCatalogResponseModel) GetJobsInProgressOk() (*[]RefResponseModel, bool)`

GetJobsInProgressOk returns a tuple with the JobsInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsInProgress

`func (o *HypervisorMachineCatalogResponseModel) SetJobsInProgress(v []RefResponseModel)`

SetJobsInProgress sets JobsInProgress field to given value.

### HasJobsInProgress

`func (o *HypervisorMachineCatalogResponseModel) HasJobsInProgress() bool`

HasJobsInProgress returns a boolean if a field has been set.

### SetJobsInProgressNil

`func (o *HypervisorMachineCatalogResponseModel) SetJobsInProgressNil(b bool)`

 SetJobsInProgressNil sets the value for JobsInProgress to be an explicit nil

### UnsetJobsInProgress
`func (o *HypervisorMachineCatalogResponseModel) UnsetJobsInProgress()`

UnsetJobsInProgress ensures that no value is present for JobsInProgress, not even an explicit nil
### GetMachineType

`func (o *HypervisorMachineCatalogResponseModel) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *HypervisorMachineCatalogResponseModel) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *HypervisorMachineCatalogResponseModel) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *HypervisorMachineCatalogResponseModel) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetMetadata

`func (o *HypervisorMachineCatalogResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorMachineCatalogResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorMachineCatalogResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorMachineCatalogResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HypervisorMachineCatalogResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HypervisorMachineCatalogResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *HypervisorMachineCatalogResponseModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *HypervisorMachineCatalogResponseModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *HypervisorMachineCatalogResponseModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.

### HasMinimumFunctionalLevel

`func (o *HypervisorMachineCatalogResponseModel) HasMinimumFunctionalLevel() bool`

HasMinimumFunctionalLevel returns a boolean if a field has been set.

### GetHasBeenPromoted

`func (o *HypervisorMachineCatalogResponseModel) GetHasBeenPromoted() bool`

GetHasBeenPromoted returns the HasBeenPromoted field if non-nil, zero value otherwise.

### GetHasBeenPromotedOk

`func (o *HypervisorMachineCatalogResponseModel) GetHasBeenPromotedOk() (*bool, bool)`

GetHasBeenPromotedOk returns a tuple with the HasBeenPromoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromoted

`func (o *HypervisorMachineCatalogResponseModel) SetHasBeenPromoted(v bool)`

SetHasBeenPromoted sets HasBeenPromoted field to given value.

### HasHasBeenPromoted

`func (o *HypervisorMachineCatalogResponseModel) HasHasBeenPromoted() bool`

HasHasBeenPromoted returns a boolean if a field has been set.

### GetHasBeenPromotedFrom

`func (o *HypervisorMachineCatalogResponseModel) GetHasBeenPromotedFrom() FunctionalLevel`

GetHasBeenPromotedFrom returns the HasBeenPromotedFrom field if non-nil, zero value otherwise.

### GetHasBeenPromotedFromOk

`func (o *HypervisorMachineCatalogResponseModel) GetHasBeenPromotedFromOk() (*FunctionalLevel, bool)`

GetHasBeenPromotedFromOk returns a tuple with the HasBeenPromotedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromotedFrom

`func (o *HypervisorMachineCatalogResponseModel) SetHasBeenPromotedFrom(v FunctionalLevel)`

SetHasBeenPromotedFrom sets HasBeenPromotedFrom field to given value.

### HasHasBeenPromotedFrom

`func (o *HypervisorMachineCatalogResponseModel) HasHasBeenPromotedFrom() bool`

HasHasBeenPromotedFrom returns a boolean if a field has been set.

### GetCanRollbackVMImage

`func (o *HypervisorMachineCatalogResponseModel) GetCanRollbackVMImage() bool`

GetCanRollbackVMImage returns the CanRollbackVMImage field if non-nil, zero value otherwise.

### GetCanRollbackVMImageOk

`func (o *HypervisorMachineCatalogResponseModel) GetCanRollbackVMImageOk() (*bool, bool)`

GetCanRollbackVMImageOk returns a tuple with the CanRollbackVMImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRollbackVMImage

`func (o *HypervisorMachineCatalogResponseModel) SetCanRollbackVMImage(v bool)`

SetCanRollbackVMImage sets CanRollbackVMImage field to given value.

### HasCanRollbackVMImage

`func (o *HypervisorMachineCatalogResponseModel) HasCanRollbackVMImage() bool`

HasCanRollbackVMImage returns a boolean if a field has been set.

### GetCanRecreateCatalog

`func (o *HypervisorMachineCatalogResponseModel) GetCanRecreateCatalog() bool`

GetCanRecreateCatalog returns the CanRecreateCatalog field if non-nil, zero value otherwise.

### GetCanRecreateCatalogOk

`func (o *HypervisorMachineCatalogResponseModel) GetCanRecreateCatalogOk() (*bool, bool)`

GetCanRecreateCatalogOk returns a tuple with the CanRecreateCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRecreateCatalog

`func (o *HypervisorMachineCatalogResponseModel) SetCanRecreateCatalog(v bool)`

SetCanRecreateCatalog sets CanRecreateCatalog field to given value.

### HasCanRecreateCatalog

`func (o *HypervisorMachineCatalogResponseModel) HasCanRecreateCatalog() bool`

HasCanRecreateCatalog returns a boolean if a field has been set.

### GetPersistChanges

`func (o *HypervisorMachineCatalogResponseModel) GetPersistChanges() PersistChanges`

GetPersistChanges returns the PersistChanges field if non-nil, zero value otherwise.

### GetPersistChangesOk

`func (o *HypervisorMachineCatalogResponseModel) GetPersistChangesOk() (*PersistChanges, bool)`

GetPersistChangesOk returns a tuple with the PersistChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistChanges

`func (o *HypervisorMachineCatalogResponseModel) SetPersistChanges(v PersistChanges)`

SetPersistChanges sets PersistChanges field to given value.

### HasPersistChanges

`func (o *HypervisorMachineCatalogResponseModel) HasPersistChanges() bool`

HasPersistChanges returns a boolean if a field has been set.

### GetProvisioningScheme

`func (o *HypervisorMachineCatalogResponseModel) GetProvisioningScheme() ProvisioningSchemeResponseModel`

GetProvisioningScheme returns the ProvisioningScheme field if non-nil, zero value otherwise.

### GetProvisioningSchemeOk

`func (o *HypervisorMachineCatalogResponseModel) GetProvisioningSchemeOk() (*ProvisioningSchemeResponseModel, bool)`

GetProvisioningSchemeOk returns a tuple with the ProvisioningScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningScheme

`func (o *HypervisorMachineCatalogResponseModel) SetProvisioningScheme(v ProvisioningSchemeResponseModel)`

SetProvisioningScheme sets ProvisioningScheme field to given value.

### HasProvisioningScheme

`func (o *HypervisorMachineCatalogResponseModel) HasProvisioningScheme() bool`

HasProvisioningScheme returns a boolean if a field has been set.

### GetProvisioningType

`func (o *HypervisorMachineCatalogResponseModel) GetProvisioningType() ProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *HypervisorMachineCatalogResponseModel) GetProvisioningTypeOk() (*ProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *HypervisorMachineCatalogResponseModel) SetProvisioningType(v ProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.

### HasProvisioningType

`func (o *HypervisorMachineCatalogResponseModel) HasProvisioningType() bool`

HasProvisioningType returns a boolean if a field has been set.

### GetProvisioningProgress

`func (o *HypervisorMachineCatalogResponseModel) GetProvisioningProgress() ProvisioningProgressResponseModel`

GetProvisioningProgress returns the ProvisioningProgress field if non-nil, zero value otherwise.

### GetProvisioningProgressOk

`func (o *HypervisorMachineCatalogResponseModel) GetProvisioningProgressOk() (*ProvisioningProgressResponseModel, bool)`

GetProvisioningProgressOk returns a tuple with the ProvisioningProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProgress

`func (o *HypervisorMachineCatalogResponseModel) SetProvisioningProgress(v ProvisioningProgressResponseModel)`

SetProvisioningProgress sets ProvisioningProgress field to given value.

### HasProvisioningProgress

`func (o *HypervisorMachineCatalogResponseModel) HasProvisioningProgress() bool`

HasProvisioningProgress returns a boolean if a field has been set.

### GetPvsAddress

`func (o *HypervisorMachineCatalogResponseModel) GetPvsAddress() string`

GetPvsAddress returns the PvsAddress field if non-nil, zero value otherwise.

### GetPvsAddressOk

`func (o *HypervisorMachineCatalogResponseModel) GetPvsAddressOk() (*string, bool)`

GetPvsAddressOk returns a tuple with the PvsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsAddress

`func (o *HypervisorMachineCatalogResponseModel) SetPvsAddress(v string)`

SetPvsAddress sets PvsAddress field to given value.

### HasPvsAddress

`func (o *HypervisorMachineCatalogResponseModel) HasPvsAddress() bool`

HasPvsAddress returns a boolean if a field has been set.

### SetPvsAddressNil

`func (o *HypervisorMachineCatalogResponseModel) SetPvsAddressNil(b bool)`

 SetPvsAddressNil sets the value for PvsAddress to be an explicit nil

### UnsetPvsAddress
`func (o *HypervisorMachineCatalogResponseModel) UnsetPvsAddress()`

UnsetPvsAddress ensures that no value is present for PvsAddress, not even an explicit nil
### GetPvsDomain

`func (o *HypervisorMachineCatalogResponseModel) GetPvsDomain() string`

GetPvsDomain returns the PvsDomain field if non-nil, zero value otherwise.

### GetPvsDomainOk

`func (o *HypervisorMachineCatalogResponseModel) GetPvsDomainOk() (*string, bool)`

GetPvsDomainOk returns a tuple with the PvsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsDomain

`func (o *HypervisorMachineCatalogResponseModel) SetPvsDomain(v string)`

SetPvsDomain sets PvsDomain field to given value.

### HasPvsDomain

`func (o *HypervisorMachineCatalogResponseModel) HasPvsDomain() bool`

HasPvsDomain returns a boolean if a field has been set.

### SetPvsDomainNil

`func (o *HypervisorMachineCatalogResponseModel) SetPvsDomainNil(b bool)`

 SetPvsDomainNil sets the value for PvsDomain to be an explicit nil

### UnsetPvsDomain
`func (o *HypervisorMachineCatalogResponseModel) UnsetPvsDomain()`

UnsetPvsDomain ensures that no value is present for PvsDomain, not even an explicit nil
### GetRemotePCEnrollmentScopes

`func (o *HypervisorMachineCatalogResponseModel) GetRemotePCEnrollmentScopes() []RemotePCEnrollmentScopeResponseModel`

GetRemotePCEnrollmentScopes returns the RemotePCEnrollmentScopes field if non-nil, zero value otherwise.

### GetRemotePCEnrollmentScopesOk

`func (o *HypervisorMachineCatalogResponseModel) GetRemotePCEnrollmentScopesOk() (*[]RemotePCEnrollmentScopeResponseModel, bool)`

GetRemotePCEnrollmentScopesOk returns a tuple with the RemotePCEnrollmentScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePCEnrollmentScopes

`func (o *HypervisorMachineCatalogResponseModel) SetRemotePCEnrollmentScopes(v []RemotePCEnrollmentScopeResponseModel)`

SetRemotePCEnrollmentScopes sets RemotePCEnrollmentScopes field to given value.

### HasRemotePCEnrollmentScopes

`func (o *HypervisorMachineCatalogResponseModel) HasRemotePCEnrollmentScopes() bool`

HasRemotePCEnrollmentScopes returns a boolean if a field has been set.

### SetRemotePCEnrollmentScopesNil

`func (o *HypervisorMachineCatalogResponseModel) SetRemotePCEnrollmentScopesNil(b bool)`

 SetRemotePCEnrollmentScopesNil sets the value for RemotePCEnrollmentScopes to be an explicit nil

### UnsetRemotePCEnrollmentScopes
`func (o *HypervisorMachineCatalogResponseModel) UnsetRemotePCEnrollmentScopes()`

UnsetRemotePCEnrollmentScopes ensures that no value is present for RemotePCEnrollmentScopes, not even an explicit nil
### GetScopes

`func (o *HypervisorMachineCatalogResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HypervisorMachineCatalogResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HypervisorMachineCatalogResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *HypervisorMachineCatalogResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *HypervisorMachineCatalogResponseModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *HypervisorMachineCatalogResponseModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *HypervisorMachineCatalogResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *HypervisorMachineCatalogResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *HypervisorMachineCatalogResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *HypervisorMachineCatalogResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *HypervisorMachineCatalogResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *HypervisorMachineCatalogResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetSessionSupport

`func (o *HypervisorMachineCatalogResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *HypervisorMachineCatalogResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *HypervisorMachineCatalogResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *HypervisorMachineCatalogResponseModel) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### GetSharingKind

`func (o *HypervisorMachineCatalogResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *HypervisorMachineCatalogResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *HypervisorMachineCatalogResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.

### HasSharingKind

`func (o *HypervisorMachineCatalogResponseModel) HasSharingKind() bool`

HasSharingKind returns a boolean if a field has been set.

### GetTotalCount

`func (o *HypervisorMachineCatalogResponseModel) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *HypervisorMachineCatalogResponseModel) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *HypervisorMachineCatalogResponseModel) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *HypervisorMachineCatalogResponseModel) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetIsBroken

`func (o *HypervisorMachineCatalogResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *HypervisorMachineCatalogResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *HypervisorMachineCatalogResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.

### HasIsBroken

`func (o *HypervisorMachineCatalogResponseModel) HasIsBroken() bool`

HasIsBroken returns a boolean if a field has been set.

### GetIsMasterImageAssociated

`func (o *HypervisorMachineCatalogResponseModel) GetIsMasterImageAssociated() bool`

GetIsMasterImageAssociated returns the IsMasterImageAssociated field if non-nil, zero value otherwise.

### GetIsMasterImageAssociatedOk

`func (o *HypervisorMachineCatalogResponseModel) GetIsMasterImageAssociatedOk() (*bool, bool)`

GetIsMasterImageAssociatedOk returns a tuple with the IsMasterImageAssociated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterImageAssociated

`func (o *HypervisorMachineCatalogResponseModel) SetIsMasterImageAssociated(v bool)`

SetIsMasterImageAssociated sets IsMasterImageAssociated field to given value.

### HasIsMasterImageAssociated

`func (o *HypervisorMachineCatalogResponseModel) HasIsMasterImageAssociated() bool`

HasIsMasterImageAssociated returns a boolean if a field has been set.

### SetIsMasterImageAssociatedNil

`func (o *HypervisorMachineCatalogResponseModel) SetIsMasterImageAssociatedNil(b bool)`

 SetIsMasterImageAssociatedNil sets the value for IsMasterImageAssociated to be an explicit nil

### UnsetIsMasterImageAssociated
`func (o *HypervisorMachineCatalogResponseModel) UnsetIsMasterImageAssociated()`

UnsetIsMasterImageAssociated ensures that no value is present for IsMasterImageAssociated, not even an explicit nil
### GetImageUpdateStatus

`func (o *HypervisorMachineCatalogResponseModel) GetImageUpdateStatus() ImageUpdateStatus`

GetImageUpdateStatus returns the ImageUpdateStatus field if non-nil, zero value otherwise.

### GetImageUpdateStatusOk

`func (o *HypervisorMachineCatalogResponseModel) GetImageUpdateStatusOk() (*ImageUpdateStatus, bool)`

GetImageUpdateStatusOk returns a tuple with the ImageUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUpdateStatus

`func (o *HypervisorMachineCatalogResponseModel) SetImageUpdateStatus(v ImageUpdateStatus)`

SetImageUpdateStatus sets ImageUpdateStatus field to given value.

### HasImageUpdateStatus

`func (o *HypervisorMachineCatalogResponseModel) HasImageUpdateStatus() bool`

HasImageUpdateStatus returns a boolean if a field has been set.

### GetErrors

`func (o *HypervisorMachineCatalogResponseModel) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HypervisorMachineCatalogResponseModel) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HypervisorMachineCatalogResponseModel) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *HypervisorMachineCatalogResponseModel) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *HypervisorMachineCatalogResponseModel) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *HypervisorMachineCatalogResponseModel) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *HypervisorMachineCatalogResponseModel) GetWarnings() []MachineCatalogWarningResponseModel`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *HypervisorMachineCatalogResponseModel) GetWarningsOk() (*[]MachineCatalogWarningResponseModel, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *HypervisorMachineCatalogResponseModel) SetWarnings(v []MachineCatalogWarningResponseModel)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *HypervisorMachineCatalogResponseModel) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *HypervisorMachineCatalogResponseModel) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *HypervisorMachineCatalogResponseModel) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetUnassignedCount

`func (o *HypervisorMachineCatalogResponseModel) GetUnassignedCount() int32`

GetUnassignedCount returns the UnassignedCount field if non-nil, zero value otherwise.

### GetUnassignedCountOk

`func (o *HypervisorMachineCatalogResponseModel) GetUnassignedCountOk() (*int32, bool)`

GetUnassignedCountOk returns a tuple with the UnassignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedCount

`func (o *HypervisorMachineCatalogResponseModel) SetUnassignedCount(v int32)`

SetUnassignedCount sets UnassignedCount field to given value.

### HasUnassignedCount

`func (o *HypervisorMachineCatalogResponseModel) HasUnassignedCount() bool`

HasUnassignedCount returns a boolean if a field has been set.

### GetUsedCount

`func (o *HypervisorMachineCatalogResponseModel) GetUsedCount() int32`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *HypervisorMachineCatalogResponseModel) GetUsedCountOk() (*int32, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *HypervisorMachineCatalogResponseModel) SetUsedCount(v int32)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *HypervisorMachineCatalogResponseModel) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetAvailableCountOfSuspend

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableCountOfSuspend() int32`

GetAvailableCountOfSuspend returns the AvailableCountOfSuspend field if non-nil, zero value otherwise.

### GetAvailableCountOfSuspendOk

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableCountOfSuspendOk() (*int32, bool)`

GetAvailableCountOfSuspendOk returns a tuple with the AvailableCountOfSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCountOfSuspend

`func (o *HypervisorMachineCatalogResponseModel) SetAvailableCountOfSuspend(v int32)`

SetAvailableCountOfSuspend sets AvailableCountOfSuspend field to given value.

### HasAvailableCountOfSuspend

`func (o *HypervisorMachineCatalogResponseModel) HasAvailableCountOfSuspend() bool`

HasAvailableCountOfSuspend returns a boolean if a field has been set.

### SetAvailableCountOfSuspendNil

`func (o *HypervisorMachineCatalogResponseModel) SetAvailableCountOfSuspendNil(b bool)`

 SetAvailableCountOfSuspendNil sets the value for AvailableCountOfSuspend to be an explicit nil

### UnsetAvailableCountOfSuspend
`func (o *HypervisorMachineCatalogResponseModel) UnsetAvailableCountOfSuspend()`

UnsetAvailableCountOfSuspend ensures that no value is present for AvailableCountOfSuspend, not even an explicit nil
### GetAvailableAssignedCountOfSuspend

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableAssignedCountOfSuspend() int32`

GetAvailableAssignedCountOfSuspend returns the AvailableAssignedCountOfSuspend field if non-nil, zero value otherwise.

### GetAvailableAssignedCountOfSuspendOk

`func (o *HypervisorMachineCatalogResponseModel) GetAvailableAssignedCountOfSuspendOk() (*int32, bool)`

GetAvailableAssignedCountOfSuspendOk returns a tuple with the AvailableAssignedCountOfSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAssignedCountOfSuspend

`func (o *HypervisorMachineCatalogResponseModel) SetAvailableAssignedCountOfSuspend(v int32)`

SetAvailableAssignedCountOfSuspend sets AvailableAssignedCountOfSuspend field to given value.

### HasAvailableAssignedCountOfSuspend

`func (o *HypervisorMachineCatalogResponseModel) HasAvailableAssignedCountOfSuspend() bool`

HasAvailableAssignedCountOfSuspend returns a boolean if a field has been set.

### SetAvailableAssignedCountOfSuspendNil

`func (o *HypervisorMachineCatalogResponseModel) SetAvailableAssignedCountOfSuspendNil(b bool)`

 SetAvailableAssignedCountOfSuspendNil sets the value for AvailableAssignedCountOfSuspend to be an explicit nil

### UnsetAvailableAssignedCountOfSuspend
`func (o *HypervisorMachineCatalogResponseModel) UnsetAvailableAssignedCountOfSuspend()`

UnsetAvailableAssignedCountOfSuspend ensures that no value is present for AvailableAssignedCountOfSuspend, not even an explicit nil
### GetUpgradeInfo

`func (o *HypervisorMachineCatalogResponseModel) GetUpgradeInfo() MachineCatalogUpgradeInfo`

GetUpgradeInfo returns the UpgradeInfo field if non-nil, zero value otherwise.

### GetUpgradeInfoOk

`func (o *HypervisorMachineCatalogResponseModel) GetUpgradeInfoOk() (*MachineCatalogUpgradeInfo, bool)`

GetUpgradeInfoOk returns a tuple with the UpgradeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeInfo

`func (o *HypervisorMachineCatalogResponseModel) SetUpgradeInfo(v MachineCatalogUpgradeInfo)`

SetUpgradeInfo sets UpgradeInfo field to given value.

### HasUpgradeInfo

`func (o *HypervisorMachineCatalogResponseModel) HasUpgradeInfo() bool`

HasUpgradeInfo returns a boolean if a field has been set.

### GetZone

`func (o *HypervisorMachineCatalogResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HypervisorMachineCatalogResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HypervisorMachineCatalogResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.

### HasZone

`func (o *HypervisorMachineCatalogResponseModel) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetAdminFolder

`func (o *HypervisorMachineCatalogResponseModel) GetAdminFolder() RefResponseModel`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *HypervisorMachineCatalogResponseModel) GetAdminFolderOk() (*RefResponseModel, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *HypervisorMachineCatalogResponseModel) SetAdminFolder(v RefResponseModel)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *HypervisorMachineCatalogResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### GetHypervisorVMTagging

`func (o *HypervisorMachineCatalogResponseModel) GetHypervisorVMTagging() bool`

GetHypervisorVMTagging returns the HypervisorVMTagging field if non-nil, zero value otherwise.

### GetHypervisorVMTaggingOk

`func (o *HypervisorMachineCatalogResponseModel) GetHypervisorVMTaggingOk() (*bool, bool)`

GetHypervisorVMTaggingOk returns a tuple with the HypervisorVMTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVMTagging

`func (o *HypervisorMachineCatalogResponseModel) SetHypervisorVMTagging(v bool)`

SetHypervisorVMTagging sets HypervisorVMTagging field to given value.

### HasHypervisorVMTagging

`func (o *HypervisorMachineCatalogResponseModel) HasHypervisorVMTagging() bool`

HasHypervisorVMTagging returns a boolean if a field has been set.

### SetHypervisorVMTaggingNil

`func (o *HypervisorMachineCatalogResponseModel) SetHypervisorVMTaggingNil(b bool)`

 SetHypervisorVMTaggingNil sets the value for HypervisorVMTagging to be an explicit nil

### UnsetHypervisorVMTagging
`func (o *HypervisorMachineCatalogResponseModel) UnsetHypervisorVMTagging()`

UnsetHypervisorVMTagging ensures that no value is present for HypervisorVMTagging, not even an explicit nil
### GetNumProvisionedMachines

`func (o *HypervisorMachineCatalogResponseModel) GetNumProvisionedMachines() int32`

GetNumProvisionedMachines returns the NumProvisionedMachines field if non-nil, zero value otherwise.

### GetNumProvisionedMachinesOk

`func (o *HypervisorMachineCatalogResponseModel) GetNumProvisionedMachinesOk() (*int32, bool)`

GetNumProvisionedMachinesOk returns a tuple with the NumProvisionedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProvisionedMachines

`func (o *HypervisorMachineCatalogResponseModel) SetNumProvisionedMachines(v int32)`

SetNumProvisionedMachines sets NumProvisionedMachines field to given value.

### HasNumProvisionedMachines

`func (o *HypervisorMachineCatalogResponseModel) HasNumProvisionedMachines() bool`

HasNumProvisionedMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


