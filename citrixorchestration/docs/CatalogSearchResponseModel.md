# CatalogSearchResponseModel

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
**MachinesArePhysical** | Pointer to **bool** | Indicates whether or not machines in the machine catalog are Physical. | [optional] 
**HypervisorPluginResponse** | Pointer to [**HypervisorPluginResponseModel**](HypervisorPluginResponseModel.md) |  | [optional] 
**HypervisorConnection** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**HypervisorConnectionUid** | Pointer to **int32** | The Uid of the hypervisor connection that is associated with the machines in the catalog. This property only applies to MCS provisioned catalogs. For other provisioning types machines can be from one or more different hypervisor connections. | [optional] 

## Methods

### NewCatalogSearchResponseModel

`func NewCatalogSearchResponseModel() *CatalogSearchResponseModel`

NewCatalogSearchResponseModel instantiates a new CatalogSearchResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogSearchResponseModelWithDefaults

`func NewCatalogSearchResponseModelWithDefaults() *CatalogSearchResponseModel`

NewCatalogSearchResponseModelWithDefaults instantiates a new CatalogSearchResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogSearchResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogSearchResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogSearchResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogSearchResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CatalogSearchResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CatalogSearchResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *CatalogSearchResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CatalogSearchResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CatalogSearchResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *CatalogSearchResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *CatalogSearchResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *CatalogSearchResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetId

`func (o *CatalogSearchResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogSearchResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogSearchResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogSearchResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CatalogSearchResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CatalogSearchResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *CatalogSearchResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CatalogSearchResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CatalogSearchResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *CatalogSearchResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetAllocationType

`func (o *CatalogSearchResponseModel) GetAllocationType() AllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *CatalogSearchResponseModel) GetAllocationTypeOk() (*AllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *CatalogSearchResponseModel) SetAllocationType(v AllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *CatalogSearchResponseModel) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetAssignedCount

`func (o *CatalogSearchResponseModel) GetAssignedCount() int32`

GetAssignedCount returns the AssignedCount field if non-nil, zero value otherwise.

### GetAssignedCountOk

`func (o *CatalogSearchResponseModel) GetAssignedCountOk() (*int32, bool)`

GetAssignedCountOk returns a tuple with the AssignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCount

`func (o *CatalogSearchResponseModel) SetAssignedCount(v int32)`

SetAssignedCount sets AssignedCount field to given value.

### HasAssignedCount

`func (o *CatalogSearchResponseModel) HasAssignedCount() bool`

HasAssignedCount returns a boolean if a field has been set.

### SetAssignedCountNil

`func (o *CatalogSearchResponseModel) SetAssignedCountNil(b bool)`

 SetAssignedCountNil sets the value for AssignedCount to be an explicit nil

### UnsetAssignedCount
`func (o *CatalogSearchResponseModel) UnsetAssignedCount()`

UnsetAssignedCount ensures that no value is present for AssignedCount, not even an explicit nil
### GetAvailableAssignedCount

`func (o *CatalogSearchResponseModel) GetAvailableAssignedCount() int32`

GetAvailableAssignedCount returns the AvailableAssignedCount field if non-nil, zero value otherwise.

### GetAvailableAssignedCountOk

`func (o *CatalogSearchResponseModel) GetAvailableAssignedCountOk() (*int32, bool)`

GetAvailableAssignedCountOk returns a tuple with the AvailableAssignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAssignedCount

`func (o *CatalogSearchResponseModel) SetAvailableAssignedCount(v int32)`

SetAvailableAssignedCount sets AvailableAssignedCount field to given value.

### HasAvailableAssignedCount

`func (o *CatalogSearchResponseModel) HasAvailableAssignedCount() bool`

HasAvailableAssignedCount returns a boolean if a field has been set.

### SetAvailableAssignedCountNil

`func (o *CatalogSearchResponseModel) SetAvailableAssignedCountNil(b bool)`

 SetAvailableAssignedCountNil sets the value for AvailableAssignedCount to be an explicit nil

### UnsetAvailableAssignedCount
`func (o *CatalogSearchResponseModel) UnsetAvailableAssignedCount()`

UnsetAvailableAssignedCount ensures that no value is present for AvailableAssignedCount, not even an explicit nil
### GetAvailableCount

`func (o *CatalogSearchResponseModel) GetAvailableCount() int32`

GetAvailableCount returns the AvailableCount field if non-nil, zero value otherwise.

### GetAvailableCountOk

`func (o *CatalogSearchResponseModel) GetAvailableCountOk() (*int32, bool)`

GetAvailableCountOk returns a tuple with the AvailableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCount

`func (o *CatalogSearchResponseModel) SetAvailableCount(v int32)`

SetAvailableCount sets AvailableCount field to given value.

### HasAvailableCount

`func (o *CatalogSearchResponseModel) HasAvailableCount() bool`

HasAvailableCount returns a boolean if a field has been set.

### GetAvailableUnassignedCount

`func (o *CatalogSearchResponseModel) GetAvailableUnassignedCount() int32`

GetAvailableUnassignedCount returns the AvailableUnassignedCount field if non-nil, zero value otherwise.

### GetAvailableUnassignedCountOk

`func (o *CatalogSearchResponseModel) GetAvailableUnassignedCountOk() (*int32, bool)`

GetAvailableUnassignedCountOk returns a tuple with the AvailableUnassignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUnassignedCount

`func (o *CatalogSearchResponseModel) SetAvailableUnassignedCount(v int32)`

SetAvailableUnassignedCount sets AvailableUnassignedCount field to given value.

### HasAvailableUnassignedCount

`func (o *CatalogSearchResponseModel) HasAvailableUnassignedCount() bool`

HasAvailableUnassignedCount returns a boolean if a field has been set.

### SetAvailableUnassignedCountNil

`func (o *CatalogSearchResponseModel) SetAvailableUnassignedCountNil(b bool)`

 SetAvailableUnassignedCountNil sets the value for AvailableUnassignedCount to be an explicit nil

### UnsetAvailableUnassignedCount
`func (o *CatalogSearchResponseModel) UnsetAvailableUnassignedCount()`

UnsetAvailableUnassignedCount ensures that no value is present for AvailableUnassignedCount, not even an explicit nil
### GetDescription

`func (o *CatalogSearchResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogSearchResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogSearchResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogSearchResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogSearchResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogSearchResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsPowerManaged

`func (o *CatalogSearchResponseModel) GetIsPowerManaged() bool`

GetIsPowerManaged returns the IsPowerManaged field if non-nil, zero value otherwise.

### GetIsPowerManagedOk

`func (o *CatalogSearchResponseModel) GetIsPowerManagedOk() (*bool, bool)`

GetIsPowerManagedOk returns a tuple with the IsPowerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPowerManaged

`func (o *CatalogSearchResponseModel) SetIsPowerManaged(v bool)`

SetIsPowerManaged sets IsPowerManaged field to given value.

### HasIsPowerManaged

`func (o *CatalogSearchResponseModel) HasIsPowerManaged() bool`

HasIsPowerManaged returns a boolean if a field has been set.

### GetIsRemotePC

`func (o *CatalogSearchResponseModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *CatalogSearchResponseModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *CatalogSearchResponseModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.

### HasIsRemotePC

`func (o *CatalogSearchResponseModel) HasIsRemotePC() bool`

HasIsRemotePC returns a boolean if a field has been set.

### GetJobsInProgress

`func (o *CatalogSearchResponseModel) GetJobsInProgress() []RefResponseModel`

GetJobsInProgress returns the JobsInProgress field if non-nil, zero value otherwise.

### GetJobsInProgressOk

`func (o *CatalogSearchResponseModel) GetJobsInProgressOk() (*[]RefResponseModel, bool)`

GetJobsInProgressOk returns a tuple with the JobsInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsInProgress

`func (o *CatalogSearchResponseModel) SetJobsInProgress(v []RefResponseModel)`

SetJobsInProgress sets JobsInProgress field to given value.

### HasJobsInProgress

`func (o *CatalogSearchResponseModel) HasJobsInProgress() bool`

HasJobsInProgress returns a boolean if a field has been set.

### SetJobsInProgressNil

`func (o *CatalogSearchResponseModel) SetJobsInProgressNil(b bool)`

 SetJobsInProgressNil sets the value for JobsInProgress to be an explicit nil

### UnsetJobsInProgress
`func (o *CatalogSearchResponseModel) UnsetJobsInProgress()`

UnsetJobsInProgress ensures that no value is present for JobsInProgress, not even an explicit nil
### GetMachineType

`func (o *CatalogSearchResponseModel) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *CatalogSearchResponseModel) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *CatalogSearchResponseModel) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *CatalogSearchResponseModel) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetMetadata

`func (o *CatalogSearchResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CatalogSearchResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CatalogSearchResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CatalogSearchResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CatalogSearchResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CatalogSearchResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *CatalogSearchResponseModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *CatalogSearchResponseModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *CatalogSearchResponseModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.

### HasMinimumFunctionalLevel

`func (o *CatalogSearchResponseModel) HasMinimumFunctionalLevel() bool`

HasMinimumFunctionalLevel returns a boolean if a field has been set.

### GetHasBeenPromoted

`func (o *CatalogSearchResponseModel) GetHasBeenPromoted() bool`

GetHasBeenPromoted returns the HasBeenPromoted field if non-nil, zero value otherwise.

### GetHasBeenPromotedOk

`func (o *CatalogSearchResponseModel) GetHasBeenPromotedOk() (*bool, bool)`

GetHasBeenPromotedOk returns a tuple with the HasBeenPromoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromoted

`func (o *CatalogSearchResponseModel) SetHasBeenPromoted(v bool)`

SetHasBeenPromoted sets HasBeenPromoted field to given value.

### HasHasBeenPromoted

`func (o *CatalogSearchResponseModel) HasHasBeenPromoted() bool`

HasHasBeenPromoted returns a boolean if a field has been set.

### GetHasBeenPromotedFrom

`func (o *CatalogSearchResponseModel) GetHasBeenPromotedFrom() FunctionalLevel`

GetHasBeenPromotedFrom returns the HasBeenPromotedFrom field if non-nil, zero value otherwise.

### GetHasBeenPromotedFromOk

`func (o *CatalogSearchResponseModel) GetHasBeenPromotedFromOk() (*FunctionalLevel, bool)`

GetHasBeenPromotedFromOk returns a tuple with the HasBeenPromotedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromotedFrom

`func (o *CatalogSearchResponseModel) SetHasBeenPromotedFrom(v FunctionalLevel)`

SetHasBeenPromotedFrom sets HasBeenPromotedFrom field to given value.

### HasHasBeenPromotedFrom

`func (o *CatalogSearchResponseModel) HasHasBeenPromotedFrom() bool`

HasHasBeenPromotedFrom returns a boolean if a field has been set.

### GetCanRollbackVMImage

`func (o *CatalogSearchResponseModel) GetCanRollbackVMImage() bool`

GetCanRollbackVMImage returns the CanRollbackVMImage field if non-nil, zero value otherwise.

### GetCanRollbackVMImageOk

`func (o *CatalogSearchResponseModel) GetCanRollbackVMImageOk() (*bool, bool)`

GetCanRollbackVMImageOk returns a tuple with the CanRollbackVMImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRollbackVMImage

`func (o *CatalogSearchResponseModel) SetCanRollbackVMImage(v bool)`

SetCanRollbackVMImage sets CanRollbackVMImage field to given value.

### HasCanRollbackVMImage

`func (o *CatalogSearchResponseModel) HasCanRollbackVMImage() bool`

HasCanRollbackVMImage returns a boolean if a field has been set.

### GetCanRecreateCatalog

`func (o *CatalogSearchResponseModel) GetCanRecreateCatalog() bool`

GetCanRecreateCatalog returns the CanRecreateCatalog field if non-nil, zero value otherwise.

### GetCanRecreateCatalogOk

`func (o *CatalogSearchResponseModel) GetCanRecreateCatalogOk() (*bool, bool)`

GetCanRecreateCatalogOk returns a tuple with the CanRecreateCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRecreateCatalog

`func (o *CatalogSearchResponseModel) SetCanRecreateCatalog(v bool)`

SetCanRecreateCatalog sets CanRecreateCatalog field to given value.

### HasCanRecreateCatalog

`func (o *CatalogSearchResponseModel) HasCanRecreateCatalog() bool`

HasCanRecreateCatalog returns a boolean if a field has been set.

### GetPersistChanges

`func (o *CatalogSearchResponseModel) GetPersistChanges() PersistChanges`

GetPersistChanges returns the PersistChanges field if non-nil, zero value otherwise.

### GetPersistChangesOk

`func (o *CatalogSearchResponseModel) GetPersistChangesOk() (*PersistChanges, bool)`

GetPersistChangesOk returns a tuple with the PersistChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistChanges

`func (o *CatalogSearchResponseModel) SetPersistChanges(v PersistChanges)`

SetPersistChanges sets PersistChanges field to given value.

### HasPersistChanges

`func (o *CatalogSearchResponseModel) HasPersistChanges() bool`

HasPersistChanges returns a boolean if a field has been set.

### GetProvisioningScheme

`func (o *CatalogSearchResponseModel) GetProvisioningScheme() ProvisioningSchemeResponseModel`

GetProvisioningScheme returns the ProvisioningScheme field if non-nil, zero value otherwise.

### GetProvisioningSchemeOk

`func (o *CatalogSearchResponseModel) GetProvisioningSchemeOk() (*ProvisioningSchemeResponseModel, bool)`

GetProvisioningSchemeOk returns a tuple with the ProvisioningScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningScheme

`func (o *CatalogSearchResponseModel) SetProvisioningScheme(v ProvisioningSchemeResponseModel)`

SetProvisioningScheme sets ProvisioningScheme field to given value.

### HasProvisioningScheme

`func (o *CatalogSearchResponseModel) HasProvisioningScheme() bool`

HasProvisioningScheme returns a boolean if a field has been set.

### GetProvisioningType

`func (o *CatalogSearchResponseModel) GetProvisioningType() ProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *CatalogSearchResponseModel) GetProvisioningTypeOk() (*ProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *CatalogSearchResponseModel) SetProvisioningType(v ProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.

### HasProvisioningType

`func (o *CatalogSearchResponseModel) HasProvisioningType() bool`

HasProvisioningType returns a boolean if a field has been set.

### GetProvisioningProgress

`func (o *CatalogSearchResponseModel) GetProvisioningProgress() ProvisioningProgressResponseModel`

GetProvisioningProgress returns the ProvisioningProgress field if non-nil, zero value otherwise.

### GetProvisioningProgressOk

`func (o *CatalogSearchResponseModel) GetProvisioningProgressOk() (*ProvisioningProgressResponseModel, bool)`

GetProvisioningProgressOk returns a tuple with the ProvisioningProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProgress

`func (o *CatalogSearchResponseModel) SetProvisioningProgress(v ProvisioningProgressResponseModel)`

SetProvisioningProgress sets ProvisioningProgress field to given value.

### HasProvisioningProgress

`func (o *CatalogSearchResponseModel) HasProvisioningProgress() bool`

HasProvisioningProgress returns a boolean if a field has been set.

### GetPvsAddress

`func (o *CatalogSearchResponseModel) GetPvsAddress() string`

GetPvsAddress returns the PvsAddress field if non-nil, zero value otherwise.

### GetPvsAddressOk

`func (o *CatalogSearchResponseModel) GetPvsAddressOk() (*string, bool)`

GetPvsAddressOk returns a tuple with the PvsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsAddress

`func (o *CatalogSearchResponseModel) SetPvsAddress(v string)`

SetPvsAddress sets PvsAddress field to given value.

### HasPvsAddress

`func (o *CatalogSearchResponseModel) HasPvsAddress() bool`

HasPvsAddress returns a boolean if a field has been set.

### SetPvsAddressNil

`func (o *CatalogSearchResponseModel) SetPvsAddressNil(b bool)`

 SetPvsAddressNil sets the value for PvsAddress to be an explicit nil

### UnsetPvsAddress
`func (o *CatalogSearchResponseModel) UnsetPvsAddress()`

UnsetPvsAddress ensures that no value is present for PvsAddress, not even an explicit nil
### GetPvsDomain

`func (o *CatalogSearchResponseModel) GetPvsDomain() string`

GetPvsDomain returns the PvsDomain field if non-nil, zero value otherwise.

### GetPvsDomainOk

`func (o *CatalogSearchResponseModel) GetPvsDomainOk() (*string, bool)`

GetPvsDomainOk returns a tuple with the PvsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsDomain

`func (o *CatalogSearchResponseModel) SetPvsDomain(v string)`

SetPvsDomain sets PvsDomain field to given value.

### HasPvsDomain

`func (o *CatalogSearchResponseModel) HasPvsDomain() bool`

HasPvsDomain returns a boolean if a field has been set.

### SetPvsDomainNil

`func (o *CatalogSearchResponseModel) SetPvsDomainNil(b bool)`

 SetPvsDomainNil sets the value for PvsDomain to be an explicit nil

### UnsetPvsDomain
`func (o *CatalogSearchResponseModel) UnsetPvsDomain()`

UnsetPvsDomain ensures that no value is present for PvsDomain, not even an explicit nil
### GetRemotePCEnrollmentScopes

`func (o *CatalogSearchResponseModel) GetRemotePCEnrollmentScopes() []RemotePCEnrollmentScopeResponseModel`

GetRemotePCEnrollmentScopes returns the RemotePCEnrollmentScopes field if non-nil, zero value otherwise.

### GetRemotePCEnrollmentScopesOk

`func (o *CatalogSearchResponseModel) GetRemotePCEnrollmentScopesOk() (*[]RemotePCEnrollmentScopeResponseModel, bool)`

GetRemotePCEnrollmentScopesOk returns a tuple with the RemotePCEnrollmentScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePCEnrollmentScopes

`func (o *CatalogSearchResponseModel) SetRemotePCEnrollmentScopes(v []RemotePCEnrollmentScopeResponseModel)`

SetRemotePCEnrollmentScopes sets RemotePCEnrollmentScopes field to given value.

### HasRemotePCEnrollmentScopes

`func (o *CatalogSearchResponseModel) HasRemotePCEnrollmentScopes() bool`

HasRemotePCEnrollmentScopes returns a boolean if a field has been set.

### SetRemotePCEnrollmentScopesNil

`func (o *CatalogSearchResponseModel) SetRemotePCEnrollmentScopesNil(b bool)`

 SetRemotePCEnrollmentScopesNil sets the value for RemotePCEnrollmentScopes to be an explicit nil

### UnsetRemotePCEnrollmentScopes
`func (o *CatalogSearchResponseModel) UnsetRemotePCEnrollmentScopes()`

UnsetRemotePCEnrollmentScopes ensures that no value is present for RemotePCEnrollmentScopes, not even an explicit nil
### GetScopes

`func (o *CatalogSearchResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CatalogSearchResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CatalogSearchResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CatalogSearchResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *CatalogSearchResponseModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *CatalogSearchResponseModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *CatalogSearchResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CatalogSearchResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CatalogSearchResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CatalogSearchResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *CatalogSearchResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *CatalogSearchResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetSessionSupport

`func (o *CatalogSearchResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *CatalogSearchResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *CatalogSearchResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *CatalogSearchResponseModel) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### GetSharingKind

`func (o *CatalogSearchResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *CatalogSearchResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *CatalogSearchResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.

### HasSharingKind

`func (o *CatalogSearchResponseModel) HasSharingKind() bool`

HasSharingKind returns a boolean if a field has been set.

### GetTotalCount

`func (o *CatalogSearchResponseModel) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CatalogSearchResponseModel) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CatalogSearchResponseModel) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CatalogSearchResponseModel) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetIsBroken

`func (o *CatalogSearchResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *CatalogSearchResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *CatalogSearchResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.

### HasIsBroken

`func (o *CatalogSearchResponseModel) HasIsBroken() bool`

HasIsBroken returns a boolean if a field has been set.

### GetIsMasterImageAssociated

`func (o *CatalogSearchResponseModel) GetIsMasterImageAssociated() bool`

GetIsMasterImageAssociated returns the IsMasterImageAssociated field if non-nil, zero value otherwise.

### GetIsMasterImageAssociatedOk

`func (o *CatalogSearchResponseModel) GetIsMasterImageAssociatedOk() (*bool, bool)`

GetIsMasterImageAssociatedOk returns a tuple with the IsMasterImageAssociated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterImageAssociated

`func (o *CatalogSearchResponseModel) SetIsMasterImageAssociated(v bool)`

SetIsMasterImageAssociated sets IsMasterImageAssociated field to given value.

### HasIsMasterImageAssociated

`func (o *CatalogSearchResponseModel) HasIsMasterImageAssociated() bool`

HasIsMasterImageAssociated returns a boolean if a field has been set.

### SetIsMasterImageAssociatedNil

`func (o *CatalogSearchResponseModel) SetIsMasterImageAssociatedNil(b bool)`

 SetIsMasterImageAssociatedNil sets the value for IsMasterImageAssociated to be an explicit nil

### UnsetIsMasterImageAssociated
`func (o *CatalogSearchResponseModel) UnsetIsMasterImageAssociated()`

UnsetIsMasterImageAssociated ensures that no value is present for IsMasterImageAssociated, not even an explicit nil
### GetImageUpdateStatus

`func (o *CatalogSearchResponseModel) GetImageUpdateStatus() ImageUpdateStatus`

GetImageUpdateStatus returns the ImageUpdateStatus field if non-nil, zero value otherwise.

### GetImageUpdateStatusOk

`func (o *CatalogSearchResponseModel) GetImageUpdateStatusOk() (*ImageUpdateStatus, bool)`

GetImageUpdateStatusOk returns a tuple with the ImageUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUpdateStatus

`func (o *CatalogSearchResponseModel) SetImageUpdateStatus(v ImageUpdateStatus)`

SetImageUpdateStatus sets ImageUpdateStatus field to given value.

### HasImageUpdateStatus

`func (o *CatalogSearchResponseModel) HasImageUpdateStatus() bool`

HasImageUpdateStatus returns a boolean if a field has been set.

### GetErrors

`func (o *CatalogSearchResponseModel) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CatalogSearchResponseModel) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CatalogSearchResponseModel) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CatalogSearchResponseModel) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *CatalogSearchResponseModel) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *CatalogSearchResponseModel) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *CatalogSearchResponseModel) GetWarnings() []MachineCatalogWarningResponseModel`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CatalogSearchResponseModel) GetWarningsOk() (*[]MachineCatalogWarningResponseModel, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CatalogSearchResponseModel) SetWarnings(v []MachineCatalogWarningResponseModel)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CatalogSearchResponseModel) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *CatalogSearchResponseModel) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *CatalogSearchResponseModel) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetUnassignedCount

`func (o *CatalogSearchResponseModel) GetUnassignedCount() int32`

GetUnassignedCount returns the UnassignedCount field if non-nil, zero value otherwise.

### GetUnassignedCountOk

`func (o *CatalogSearchResponseModel) GetUnassignedCountOk() (*int32, bool)`

GetUnassignedCountOk returns a tuple with the UnassignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedCount

`func (o *CatalogSearchResponseModel) SetUnassignedCount(v int32)`

SetUnassignedCount sets UnassignedCount field to given value.

### HasUnassignedCount

`func (o *CatalogSearchResponseModel) HasUnassignedCount() bool`

HasUnassignedCount returns a boolean if a field has been set.

### GetUsedCount

`func (o *CatalogSearchResponseModel) GetUsedCount() int32`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *CatalogSearchResponseModel) GetUsedCountOk() (*int32, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *CatalogSearchResponseModel) SetUsedCount(v int32)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *CatalogSearchResponseModel) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetAvailableCountOfSuspend

`func (o *CatalogSearchResponseModel) GetAvailableCountOfSuspend() int32`

GetAvailableCountOfSuspend returns the AvailableCountOfSuspend field if non-nil, zero value otherwise.

### GetAvailableCountOfSuspendOk

`func (o *CatalogSearchResponseModel) GetAvailableCountOfSuspendOk() (*int32, bool)`

GetAvailableCountOfSuspendOk returns a tuple with the AvailableCountOfSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCountOfSuspend

`func (o *CatalogSearchResponseModel) SetAvailableCountOfSuspend(v int32)`

SetAvailableCountOfSuspend sets AvailableCountOfSuspend field to given value.

### HasAvailableCountOfSuspend

`func (o *CatalogSearchResponseModel) HasAvailableCountOfSuspend() bool`

HasAvailableCountOfSuspend returns a boolean if a field has been set.

### SetAvailableCountOfSuspendNil

`func (o *CatalogSearchResponseModel) SetAvailableCountOfSuspendNil(b bool)`

 SetAvailableCountOfSuspendNil sets the value for AvailableCountOfSuspend to be an explicit nil

### UnsetAvailableCountOfSuspend
`func (o *CatalogSearchResponseModel) UnsetAvailableCountOfSuspend()`

UnsetAvailableCountOfSuspend ensures that no value is present for AvailableCountOfSuspend, not even an explicit nil
### GetAvailableAssignedCountOfSuspend

`func (o *CatalogSearchResponseModel) GetAvailableAssignedCountOfSuspend() int32`

GetAvailableAssignedCountOfSuspend returns the AvailableAssignedCountOfSuspend field if non-nil, zero value otherwise.

### GetAvailableAssignedCountOfSuspendOk

`func (o *CatalogSearchResponseModel) GetAvailableAssignedCountOfSuspendOk() (*int32, bool)`

GetAvailableAssignedCountOfSuspendOk returns a tuple with the AvailableAssignedCountOfSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAssignedCountOfSuspend

`func (o *CatalogSearchResponseModel) SetAvailableAssignedCountOfSuspend(v int32)`

SetAvailableAssignedCountOfSuspend sets AvailableAssignedCountOfSuspend field to given value.

### HasAvailableAssignedCountOfSuspend

`func (o *CatalogSearchResponseModel) HasAvailableAssignedCountOfSuspend() bool`

HasAvailableAssignedCountOfSuspend returns a boolean if a field has been set.

### SetAvailableAssignedCountOfSuspendNil

`func (o *CatalogSearchResponseModel) SetAvailableAssignedCountOfSuspendNil(b bool)`

 SetAvailableAssignedCountOfSuspendNil sets the value for AvailableAssignedCountOfSuspend to be an explicit nil

### UnsetAvailableAssignedCountOfSuspend
`func (o *CatalogSearchResponseModel) UnsetAvailableAssignedCountOfSuspend()`

UnsetAvailableAssignedCountOfSuspend ensures that no value is present for AvailableAssignedCountOfSuspend, not even an explicit nil
### GetUpgradeInfo

`func (o *CatalogSearchResponseModel) GetUpgradeInfo() MachineCatalogUpgradeInfo`

GetUpgradeInfo returns the UpgradeInfo field if non-nil, zero value otherwise.

### GetUpgradeInfoOk

`func (o *CatalogSearchResponseModel) GetUpgradeInfoOk() (*MachineCatalogUpgradeInfo, bool)`

GetUpgradeInfoOk returns a tuple with the UpgradeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeInfo

`func (o *CatalogSearchResponseModel) SetUpgradeInfo(v MachineCatalogUpgradeInfo)`

SetUpgradeInfo sets UpgradeInfo field to given value.

### HasUpgradeInfo

`func (o *CatalogSearchResponseModel) HasUpgradeInfo() bool`

HasUpgradeInfo returns a boolean if a field has been set.

### GetZone

`func (o *CatalogSearchResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CatalogSearchResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CatalogSearchResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CatalogSearchResponseModel) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetAdminFolder

`func (o *CatalogSearchResponseModel) GetAdminFolder() RefResponseModel`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *CatalogSearchResponseModel) GetAdminFolderOk() (*RefResponseModel, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *CatalogSearchResponseModel) SetAdminFolder(v RefResponseModel)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *CatalogSearchResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### GetHypervisorVMTagging

`func (o *CatalogSearchResponseModel) GetHypervisorVMTagging() bool`

GetHypervisorVMTagging returns the HypervisorVMTagging field if non-nil, zero value otherwise.

### GetHypervisorVMTaggingOk

`func (o *CatalogSearchResponseModel) GetHypervisorVMTaggingOk() (*bool, bool)`

GetHypervisorVMTaggingOk returns a tuple with the HypervisorVMTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVMTagging

`func (o *CatalogSearchResponseModel) SetHypervisorVMTagging(v bool)`

SetHypervisorVMTagging sets HypervisorVMTagging field to given value.

### HasHypervisorVMTagging

`func (o *CatalogSearchResponseModel) HasHypervisorVMTagging() bool`

HasHypervisorVMTagging returns a boolean if a field has been set.

### SetHypervisorVMTaggingNil

`func (o *CatalogSearchResponseModel) SetHypervisorVMTaggingNil(b bool)`

 SetHypervisorVMTaggingNil sets the value for HypervisorVMTagging to be an explicit nil

### UnsetHypervisorVMTagging
`func (o *CatalogSearchResponseModel) UnsetHypervisorVMTagging()`

UnsetHypervisorVMTagging ensures that no value is present for HypervisorVMTagging, not even an explicit nil
### GetMachinesArePhysical

`func (o *CatalogSearchResponseModel) GetMachinesArePhysical() bool`

GetMachinesArePhysical returns the MachinesArePhysical field if non-nil, zero value otherwise.

### GetMachinesArePhysicalOk

`func (o *CatalogSearchResponseModel) GetMachinesArePhysicalOk() (*bool, bool)`

GetMachinesArePhysicalOk returns a tuple with the MachinesArePhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinesArePhysical

`func (o *CatalogSearchResponseModel) SetMachinesArePhysical(v bool)`

SetMachinesArePhysical sets MachinesArePhysical field to given value.

### HasMachinesArePhysical

`func (o *CatalogSearchResponseModel) HasMachinesArePhysical() bool`

HasMachinesArePhysical returns a boolean if a field has been set.

### GetHypervisorPluginResponse

`func (o *CatalogSearchResponseModel) GetHypervisorPluginResponse() HypervisorPluginResponseModel`

GetHypervisorPluginResponse returns the HypervisorPluginResponse field if non-nil, zero value otherwise.

### GetHypervisorPluginResponseOk

`func (o *CatalogSearchResponseModel) GetHypervisorPluginResponseOk() (*HypervisorPluginResponseModel, bool)`

GetHypervisorPluginResponseOk returns a tuple with the HypervisorPluginResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorPluginResponse

`func (o *CatalogSearchResponseModel) SetHypervisorPluginResponse(v HypervisorPluginResponseModel)`

SetHypervisorPluginResponse sets HypervisorPluginResponse field to given value.

### HasHypervisorPluginResponse

`func (o *CatalogSearchResponseModel) HasHypervisorPluginResponse() bool`

HasHypervisorPluginResponse returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *CatalogSearchResponseModel) GetHypervisorConnection() RefResponseModel`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *CatalogSearchResponseModel) GetHypervisorConnectionOk() (*RefResponseModel, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *CatalogSearchResponseModel) SetHypervisorConnection(v RefResponseModel)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *CatalogSearchResponseModel) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### GetHypervisorConnectionUid

`func (o *CatalogSearchResponseModel) GetHypervisorConnectionUid() int32`

GetHypervisorConnectionUid returns the HypervisorConnectionUid field if non-nil, zero value otherwise.

### GetHypervisorConnectionUidOk

`func (o *CatalogSearchResponseModel) GetHypervisorConnectionUidOk() (*int32, bool)`

GetHypervisorConnectionUidOk returns a tuple with the HypervisorConnectionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnectionUid

`func (o *CatalogSearchResponseModel) SetHypervisorConnectionUid(v int32)`

SetHypervisorConnectionUid sets HypervisorConnectionUid field to given value.

### HasHypervisorConnectionUid

`func (o *CatalogSearchResponseModel) HasHypervisorConnectionUid() bool`

HasHypervisorConnectionUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


