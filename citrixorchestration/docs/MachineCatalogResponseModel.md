# MachineCatalogResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Simple administrative name of catalog within parent admin folder (if any). This property is not guaranteed unique across all catalogs. | 
**FullName** | Pointer to **NullableString** | Unique administrative name of catalog. | [optional] 
**Id** | **string** | Id of the machine catalog. | 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED.  Use Id. | [optional] 
**AllocationType** | [**AllocationType**](AllocationType.md) |  | 
**AssignedCount** | Pointer to **NullableInt32** | The number of assigned machines (machines that have been assigned to a user/users or a client name/address). | [optional] 
**AvailableAssignedCount** | Pointer to **NullableInt32** | The number of available machines (not in a delivery group), that are also assigned to users. | [optional] 
**AvailableCount** | **int32** | The number of available machines (those not in any delivery group). | 
**AvailableUnassignedCount** | Pointer to **NullableInt32** | The number of available machines (those not in any delivery group) that are not assigned to users. | [optional] 
**Description** | Pointer to **NullableString** | Description of the machine catalog. | [optional] 
**IsPowerManaged** | Pointer to **bool** | Indicates whether the machines in the catalog are power-managed. | [optional] 
**IsRemotePC** | **bool** | Indicates whether or not the catalog is a RemotePC catalog. Remote PC catalogs automatically configure appropriate machines without the need for manual configuration. CHANGE: was public bool RemotePC { get; set; } | 
**JobsInProgress** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | List of jobs currently in progress that affect the machine catalog. | [optional] 
**MachineType** | [**MachineType**](MachineType.md) |  | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of machine catalog. | [optional] 
**MinimumFunctionalLevel** | [**FunctionalLevel**](FunctionalLevel.md) |  | 
**HasBeenPromoted** | **bool** | Whether the machine catalog was previously promoted from a lower MinimumFunctionalLevel. | 
**HasBeenPromotedFrom** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**CanRollbackVMImage** | **bool** | Whether the machine catalog can roll back VM image. | 
**CanRecreateCatalog** | **bool** | Whether the machine catalog can recreate. | 
**PersistChanges** | [**PersistChanges**](PersistChanges.md) |  | 
**ProvisioningScheme** | Pointer to [**ProvisioningSchemeResponseModel**](ProvisioningSchemeResponseModel.md) |  | [optional] 
**ProvisioningType** | [**ProvisioningType**](ProvisioningType.md) |  | 
**ProvisioningProgress** | Pointer to [**ProvisioningProgressResponseModel**](ProvisioningProgressResponseModel.md) |  | [optional] 
**PvsAddress** | Pointer to **NullableString** | IP address of the PVS server to be used. This only applies if the ProvisioningType is . | [optional] 
**PvsDomain** | Pointer to **NullableString** | The domain of the PVS server to be used. | [optional] 
**RemotePCEnrollmentScopes** | Pointer to [**[]RemotePCEnrollmentScopeResponseModel**](RemotePCEnrollmentScopeResponseModel.md) | List of one or more remote PC enrollment scopes. | [optional] 
**Scopes** | Pointer to [**[]ScopeResponseModel**](ScopeResponseModel.md) | Administrative scopes which the machine catalog is part of. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the machine catalog is assigned to.  If &#x60;null&#x60;, the machine catalog is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**SessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 
**SharingKind** | [**SharingKind**](SharingKind.md) |  | 
**TotalCount** | **int32** | The total number of machines in the catalog. | 
**IsBroken** | **bool** | Whether the machine catalog is currently in a \&quot;Broken\&quot; state. | 
**IsMasterImageAssociated** | Pointer to **NullableBool** | Whether the machine catalog is associated with a master image. | [optional] 
**Errors** | Pointer to **[]string** | Gets the Errors of machines in this catalog | [optional] 
**Warnings** | Pointer to [**[]MachineCatalogWarningResponseModel**](MachineCatalogWarningResponseModel.md) | List of warnings that are currently active on the machine catalog, if any.  If there are no warnings this will not be specified. | [optional] 
**UnassignedCount** | **int32** | The number of unassigned machines (machines not assigned to users). | 
**UsedCount** | **int32** | The number of machines in the catalog that are in a delivery group. | 
**UpgradeInfo** | Pointer to [**MachineCatalogUpgradeInfo**](MachineCatalogUpgradeInfo.md) |  | [optional] 
**Zone** | [**RefResponseModel**](RefResponseModel.md) |  | 
**AdminFolder** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 

## Methods

### NewMachineCatalogResponseModel

`func NewMachineCatalogResponseModel(name string, id string, allocationType AllocationType, availableCount int32, isRemotePC bool, machineType MachineType, minimumFunctionalLevel FunctionalLevel, hasBeenPromoted bool, canRollbackVMImage bool, canRecreateCatalog bool, persistChanges PersistChanges, provisioningType ProvisioningType, sessionSupport SessionSupport, sharingKind SharingKind, totalCount int32, isBroken bool, unassignedCount int32, usedCount int32, zone RefResponseModel, ) *MachineCatalogResponseModel`

NewMachineCatalogResponseModel instantiates a new MachineCatalogResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogResponseModelWithDefaults

`func NewMachineCatalogResponseModelWithDefaults() *MachineCatalogResponseModel`

NewMachineCatalogResponseModelWithDefaults instantiates a new MachineCatalogResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MachineCatalogResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineCatalogResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineCatalogResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *MachineCatalogResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MachineCatalogResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MachineCatalogResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MachineCatalogResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *MachineCatalogResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *MachineCatalogResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetId

`func (o *MachineCatalogResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineCatalogResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineCatalogResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *MachineCatalogResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MachineCatalogResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MachineCatalogResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MachineCatalogResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetAllocationType

`func (o *MachineCatalogResponseModel) GetAllocationType() AllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *MachineCatalogResponseModel) GetAllocationTypeOk() (*AllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *MachineCatalogResponseModel) SetAllocationType(v AllocationType)`

SetAllocationType sets AllocationType field to given value.


### GetAssignedCount

`func (o *MachineCatalogResponseModel) GetAssignedCount() int32`

GetAssignedCount returns the AssignedCount field if non-nil, zero value otherwise.

### GetAssignedCountOk

`func (o *MachineCatalogResponseModel) GetAssignedCountOk() (*int32, bool)`

GetAssignedCountOk returns a tuple with the AssignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCount

`func (o *MachineCatalogResponseModel) SetAssignedCount(v int32)`

SetAssignedCount sets AssignedCount field to given value.

### HasAssignedCount

`func (o *MachineCatalogResponseModel) HasAssignedCount() bool`

HasAssignedCount returns a boolean if a field has been set.

### SetAssignedCountNil

`func (o *MachineCatalogResponseModel) SetAssignedCountNil(b bool)`

 SetAssignedCountNil sets the value for AssignedCount to be an explicit nil

### UnsetAssignedCount
`func (o *MachineCatalogResponseModel) UnsetAssignedCount()`

UnsetAssignedCount ensures that no value is present for AssignedCount, not even an explicit nil
### GetAvailableAssignedCount

`func (o *MachineCatalogResponseModel) GetAvailableAssignedCount() int32`

GetAvailableAssignedCount returns the AvailableAssignedCount field if non-nil, zero value otherwise.

### GetAvailableAssignedCountOk

`func (o *MachineCatalogResponseModel) GetAvailableAssignedCountOk() (*int32, bool)`

GetAvailableAssignedCountOk returns a tuple with the AvailableAssignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAssignedCount

`func (o *MachineCatalogResponseModel) SetAvailableAssignedCount(v int32)`

SetAvailableAssignedCount sets AvailableAssignedCount field to given value.

### HasAvailableAssignedCount

`func (o *MachineCatalogResponseModel) HasAvailableAssignedCount() bool`

HasAvailableAssignedCount returns a boolean if a field has been set.

### SetAvailableAssignedCountNil

`func (o *MachineCatalogResponseModel) SetAvailableAssignedCountNil(b bool)`

 SetAvailableAssignedCountNil sets the value for AvailableAssignedCount to be an explicit nil

### UnsetAvailableAssignedCount
`func (o *MachineCatalogResponseModel) UnsetAvailableAssignedCount()`

UnsetAvailableAssignedCount ensures that no value is present for AvailableAssignedCount, not even an explicit nil
### GetAvailableCount

`func (o *MachineCatalogResponseModel) GetAvailableCount() int32`

GetAvailableCount returns the AvailableCount field if non-nil, zero value otherwise.

### GetAvailableCountOk

`func (o *MachineCatalogResponseModel) GetAvailableCountOk() (*int32, bool)`

GetAvailableCountOk returns a tuple with the AvailableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCount

`func (o *MachineCatalogResponseModel) SetAvailableCount(v int32)`

SetAvailableCount sets AvailableCount field to given value.


### GetAvailableUnassignedCount

`func (o *MachineCatalogResponseModel) GetAvailableUnassignedCount() int32`

GetAvailableUnassignedCount returns the AvailableUnassignedCount field if non-nil, zero value otherwise.

### GetAvailableUnassignedCountOk

`func (o *MachineCatalogResponseModel) GetAvailableUnassignedCountOk() (*int32, bool)`

GetAvailableUnassignedCountOk returns a tuple with the AvailableUnassignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUnassignedCount

`func (o *MachineCatalogResponseModel) SetAvailableUnassignedCount(v int32)`

SetAvailableUnassignedCount sets AvailableUnassignedCount field to given value.

### HasAvailableUnassignedCount

`func (o *MachineCatalogResponseModel) HasAvailableUnassignedCount() bool`

HasAvailableUnassignedCount returns a boolean if a field has been set.

### SetAvailableUnassignedCountNil

`func (o *MachineCatalogResponseModel) SetAvailableUnassignedCountNil(b bool)`

 SetAvailableUnassignedCountNil sets the value for AvailableUnassignedCount to be an explicit nil

### UnsetAvailableUnassignedCount
`func (o *MachineCatalogResponseModel) UnsetAvailableUnassignedCount()`

UnsetAvailableUnassignedCount ensures that no value is present for AvailableUnassignedCount, not even an explicit nil
### GetDescription

`func (o *MachineCatalogResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MachineCatalogResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MachineCatalogResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MachineCatalogResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MachineCatalogResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MachineCatalogResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsPowerManaged

`func (o *MachineCatalogResponseModel) GetIsPowerManaged() bool`

GetIsPowerManaged returns the IsPowerManaged field if non-nil, zero value otherwise.

### GetIsPowerManagedOk

`func (o *MachineCatalogResponseModel) GetIsPowerManagedOk() (*bool, bool)`

GetIsPowerManagedOk returns a tuple with the IsPowerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPowerManaged

`func (o *MachineCatalogResponseModel) SetIsPowerManaged(v bool)`

SetIsPowerManaged sets IsPowerManaged field to given value.

### HasIsPowerManaged

`func (o *MachineCatalogResponseModel) HasIsPowerManaged() bool`

HasIsPowerManaged returns a boolean if a field has been set.

### GetIsRemotePC

`func (o *MachineCatalogResponseModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *MachineCatalogResponseModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *MachineCatalogResponseModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.


### GetJobsInProgress

`func (o *MachineCatalogResponseModel) GetJobsInProgress() []RefResponseModel`

GetJobsInProgress returns the JobsInProgress field if non-nil, zero value otherwise.

### GetJobsInProgressOk

`func (o *MachineCatalogResponseModel) GetJobsInProgressOk() (*[]RefResponseModel, bool)`

GetJobsInProgressOk returns a tuple with the JobsInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsInProgress

`func (o *MachineCatalogResponseModel) SetJobsInProgress(v []RefResponseModel)`

SetJobsInProgress sets JobsInProgress field to given value.

### HasJobsInProgress

`func (o *MachineCatalogResponseModel) HasJobsInProgress() bool`

HasJobsInProgress returns a boolean if a field has been set.

### SetJobsInProgressNil

`func (o *MachineCatalogResponseModel) SetJobsInProgressNil(b bool)`

 SetJobsInProgressNil sets the value for JobsInProgress to be an explicit nil

### UnsetJobsInProgress
`func (o *MachineCatalogResponseModel) UnsetJobsInProgress()`

UnsetJobsInProgress ensures that no value is present for JobsInProgress, not even an explicit nil
### GetMachineType

`func (o *MachineCatalogResponseModel) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *MachineCatalogResponseModel) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *MachineCatalogResponseModel) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.


### GetMetadata

`func (o *MachineCatalogResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MachineCatalogResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MachineCatalogResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MachineCatalogResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *MachineCatalogResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *MachineCatalogResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *MachineCatalogResponseModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *MachineCatalogResponseModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *MachineCatalogResponseModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.


### GetHasBeenPromoted

`func (o *MachineCatalogResponseModel) GetHasBeenPromoted() bool`

GetHasBeenPromoted returns the HasBeenPromoted field if non-nil, zero value otherwise.

### GetHasBeenPromotedOk

`func (o *MachineCatalogResponseModel) GetHasBeenPromotedOk() (*bool, bool)`

GetHasBeenPromotedOk returns a tuple with the HasBeenPromoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromoted

`func (o *MachineCatalogResponseModel) SetHasBeenPromoted(v bool)`

SetHasBeenPromoted sets HasBeenPromoted field to given value.


### GetHasBeenPromotedFrom

`func (o *MachineCatalogResponseModel) GetHasBeenPromotedFrom() FunctionalLevel`

GetHasBeenPromotedFrom returns the HasBeenPromotedFrom field if non-nil, zero value otherwise.

### GetHasBeenPromotedFromOk

`func (o *MachineCatalogResponseModel) GetHasBeenPromotedFromOk() (*FunctionalLevel, bool)`

GetHasBeenPromotedFromOk returns a tuple with the HasBeenPromotedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromotedFrom

`func (o *MachineCatalogResponseModel) SetHasBeenPromotedFrom(v FunctionalLevel)`

SetHasBeenPromotedFrom sets HasBeenPromotedFrom field to given value.

### HasHasBeenPromotedFrom

`func (o *MachineCatalogResponseModel) HasHasBeenPromotedFrom() bool`

HasHasBeenPromotedFrom returns a boolean if a field has been set.

### GetCanRollbackVMImage

`func (o *MachineCatalogResponseModel) GetCanRollbackVMImage() bool`

GetCanRollbackVMImage returns the CanRollbackVMImage field if non-nil, zero value otherwise.

### GetCanRollbackVMImageOk

`func (o *MachineCatalogResponseModel) GetCanRollbackVMImageOk() (*bool, bool)`

GetCanRollbackVMImageOk returns a tuple with the CanRollbackVMImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRollbackVMImage

`func (o *MachineCatalogResponseModel) SetCanRollbackVMImage(v bool)`

SetCanRollbackVMImage sets CanRollbackVMImage field to given value.


### GetCanRecreateCatalog

`func (o *MachineCatalogResponseModel) GetCanRecreateCatalog() bool`

GetCanRecreateCatalog returns the CanRecreateCatalog field if non-nil, zero value otherwise.

### GetCanRecreateCatalogOk

`func (o *MachineCatalogResponseModel) GetCanRecreateCatalogOk() (*bool, bool)`

GetCanRecreateCatalogOk returns a tuple with the CanRecreateCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRecreateCatalog

`func (o *MachineCatalogResponseModel) SetCanRecreateCatalog(v bool)`

SetCanRecreateCatalog sets CanRecreateCatalog field to given value.


### GetPersistChanges

`func (o *MachineCatalogResponseModel) GetPersistChanges() PersistChanges`

GetPersistChanges returns the PersistChanges field if non-nil, zero value otherwise.

### GetPersistChangesOk

`func (o *MachineCatalogResponseModel) GetPersistChangesOk() (*PersistChanges, bool)`

GetPersistChangesOk returns a tuple with the PersistChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistChanges

`func (o *MachineCatalogResponseModel) SetPersistChanges(v PersistChanges)`

SetPersistChanges sets PersistChanges field to given value.


### GetProvisioningScheme

`func (o *MachineCatalogResponseModel) GetProvisioningScheme() ProvisioningSchemeResponseModel`

GetProvisioningScheme returns the ProvisioningScheme field if non-nil, zero value otherwise.

### GetProvisioningSchemeOk

`func (o *MachineCatalogResponseModel) GetProvisioningSchemeOk() (*ProvisioningSchemeResponseModel, bool)`

GetProvisioningSchemeOk returns a tuple with the ProvisioningScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningScheme

`func (o *MachineCatalogResponseModel) SetProvisioningScheme(v ProvisioningSchemeResponseModel)`

SetProvisioningScheme sets ProvisioningScheme field to given value.

### HasProvisioningScheme

`func (o *MachineCatalogResponseModel) HasProvisioningScheme() bool`

HasProvisioningScheme returns a boolean if a field has been set.

### GetProvisioningType

`func (o *MachineCatalogResponseModel) GetProvisioningType() ProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *MachineCatalogResponseModel) GetProvisioningTypeOk() (*ProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *MachineCatalogResponseModel) SetProvisioningType(v ProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.


### GetProvisioningProgress

`func (o *MachineCatalogResponseModel) GetProvisioningProgress() ProvisioningProgressResponseModel`

GetProvisioningProgress returns the ProvisioningProgress field if non-nil, zero value otherwise.

### GetProvisioningProgressOk

`func (o *MachineCatalogResponseModel) GetProvisioningProgressOk() (*ProvisioningProgressResponseModel, bool)`

GetProvisioningProgressOk returns a tuple with the ProvisioningProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProgress

`func (o *MachineCatalogResponseModel) SetProvisioningProgress(v ProvisioningProgressResponseModel)`

SetProvisioningProgress sets ProvisioningProgress field to given value.

### HasProvisioningProgress

`func (o *MachineCatalogResponseModel) HasProvisioningProgress() bool`

HasProvisioningProgress returns a boolean if a field has been set.

### GetPvsAddress

`func (o *MachineCatalogResponseModel) GetPvsAddress() string`

GetPvsAddress returns the PvsAddress field if non-nil, zero value otherwise.

### GetPvsAddressOk

`func (o *MachineCatalogResponseModel) GetPvsAddressOk() (*string, bool)`

GetPvsAddressOk returns a tuple with the PvsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsAddress

`func (o *MachineCatalogResponseModel) SetPvsAddress(v string)`

SetPvsAddress sets PvsAddress field to given value.

### HasPvsAddress

`func (o *MachineCatalogResponseModel) HasPvsAddress() bool`

HasPvsAddress returns a boolean if a field has been set.

### SetPvsAddressNil

`func (o *MachineCatalogResponseModel) SetPvsAddressNil(b bool)`

 SetPvsAddressNil sets the value for PvsAddress to be an explicit nil

### UnsetPvsAddress
`func (o *MachineCatalogResponseModel) UnsetPvsAddress()`

UnsetPvsAddress ensures that no value is present for PvsAddress, not even an explicit nil
### GetPvsDomain

`func (o *MachineCatalogResponseModel) GetPvsDomain() string`

GetPvsDomain returns the PvsDomain field if non-nil, zero value otherwise.

### GetPvsDomainOk

`func (o *MachineCatalogResponseModel) GetPvsDomainOk() (*string, bool)`

GetPvsDomainOk returns a tuple with the PvsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsDomain

`func (o *MachineCatalogResponseModel) SetPvsDomain(v string)`

SetPvsDomain sets PvsDomain field to given value.

### HasPvsDomain

`func (o *MachineCatalogResponseModel) HasPvsDomain() bool`

HasPvsDomain returns a boolean if a field has been set.

### SetPvsDomainNil

`func (o *MachineCatalogResponseModel) SetPvsDomainNil(b bool)`

 SetPvsDomainNil sets the value for PvsDomain to be an explicit nil

### UnsetPvsDomain
`func (o *MachineCatalogResponseModel) UnsetPvsDomain()`

UnsetPvsDomain ensures that no value is present for PvsDomain, not even an explicit nil
### GetRemotePCEnrollmentScopes

`func (o *MachineCatalogResponseModel) GetRemotePCEnrollmentScopes() []RemotePCEnrollmentScopeResponseModel`

GetRemotePCEnrollmentScopes returns the RemotePCEnrollmentScopes field if non-nil, zero value otherwise.

### GetRemotePCEnrollmentScopesOk

`func (o *MachineCatalogResponseModel) GetRemotePCEnrollmentScopesOk() (*[]RemotePCEnrollmentScopeResponseModel, bool)`

GetRemotePCEnrollmentScopesOk returns a tuple with the RemotePCEnrollmentScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePCEnrollmentScopes

`func (o *MachineCatalogResponseModel) SetRemotePCEnrollmentScopes(v []RemotePCEnrollmentScopeResponseModel)`

SetRemotePCEnrollmentScopes sets RemotePCEnrollmentScopes field to given value.

### HasRemotePCEnrollmentScopes

`func (o *MachineCatalogResponseModel) HasRemotePCEnrollmentScopes() bool`

HasRemotePCEnrollmentScopes returns a boolean if a field has been set.

### SetRemotePCEnrollmentScopesNil

`func (o *MachineCatalogResponseModel) SetRemotePCEnrollmentScopesNil(b bool)`

 SetRemotePCEnrollmentScopesNil sets the value for RemotePCEnrollmentScopes to be an explicit nil

### UnsetRemotePCEnrollmentScopes
`func (o *MachineCatalogResponseModel) UnsetRemotePCEnrollmentScopes()`

UnsetRemotePCEnrollmentScopes ensures that no value is present for RemotePCEnrollmentScopes, not even an explicit nil
### GetScopes

`func (o *MachineCatalogResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *MachineCatalogResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *MachineCatalogResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *MachineCatalogResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *MachineCatalogResponseModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *MachineCatalogResponseModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *MachineCatalogResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *MachineCatalogResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *MachineCatalogResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *MachineCatalogResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *MachineCatalogResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *MachineCatalogResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetSessionSupport

`func (o *MachineCatalogResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *MachineCatalogResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *MachineCatalogResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSharingKind

`func (o *MachineCatalogResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *MachineCatalogResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *MachineCatalogResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.


### GetTotalCount

`func (o *MachineCatalogResponseModel) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *MachineCatalogResponseModel) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *MachineCatalogResponseModel) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetIsBroken

`func (o *MachineCatalogResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *MachineCatalogResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *MachineCatalogResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.


### GetIsMasterImageAssociated

`func (o *MachineCatalogResponseModel) GetIsMasterImageAssociated() bool`

GetIsMasterImageAssociated returns the IsMasterImageAssociated field if non-nil, zero value otherwise.

### GetIsMasterImageAssociatedOk

`func (o *MachineCatalogResponseModel) GetIsMasterImageAssociatedOk() (*bool, bool)`

GetIsMasterImageAssociatedOk returns a tuple with the IsMasterImageAssociated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterImageAssociated

`func (o *MachineCatalogResponseModel) SetIsMasterImageAssociated(v bool)`

SetIsMasterImageAssociated sets IsMasterImageAssociated field to given value.

### HasIsMasterImageAssociated

`func (o *MachineCatalogResponseModel) HasIsMasterImageAssociated() bool`

HasIsMasterImageAssociated returns a boolean if a field has been set.

### SetIsMasterImageAssociatedNil

`func (o *MachineCatalogResponseModel) SetIsMasterImageAssociatedNil(b bool)`

 SetIsMasterImageAssociatedNil sets the value for IsMasterImageAssociated to be an explicit nil

### UnsetIsMasterImageAssociated
`func (o *MachineCatalogResponseModel) UnsetIsMasterImageAssociated()`

UnsetIsMasterImageAssociated ensures that no value is present for IsMasterImageAssociated, not even an explicit nil
### GetErrors

`func (o *MachineCatalogResponseModel) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MachineCatalogResponseModel) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MachineCatalogResponseModel) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MachineCatalogResponseModel) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *MachineCatalogResponseModel) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *MachineCatalogResponseModel) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *MachineCatalogResponseModel) GetWarnings() []MachineCatalogWarningResponseModel`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MachineCatalogResponseModel) GetWarningsOk() (*[]MachineCatalogWarningResponseModel, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MachineCatalogResponseModel) SetWarnings(v []MachineCatalogWarningResponseModel)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MachineCatalogResponseModel) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *MachineCatalogResponseModel) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *MachineCatalogResponseModel) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetUnassignedCount

`func (o *MachineCatalogResponseModel) GetUnassignedCount() int32`

GetUnassignedCount returns the UnassignedCount field if non-nil, zero value otherwise.

### GetUnassignedCountOk

`func (o *MachineCatalogResponseModel) GetUnassignedCountOk() (*int32, bool)`

GetUnassignedCountOk returns a tuple with the UnassignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedCount

`func (o *MachineCatalogResponseModel) SetUnassignedCount(v int32)`

SetUnassignedCount sets UnassignedCount field to given value.


### GetUsedCount

`func (o *MachineCatalogResponseModel) GetUsedCount() int32`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *MachineCatalogResponseModel) GetUsedCountOk() (*int32, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *MachineCatalogResponseModel) SetUsedCount(v int32)`

SetUsedCount sets UsedCount field to given value.


### GetUpgradeInfo

`func (o *MachineCatalogResponseModel) GetUpgradeInfo() MachineCatalogUpgradeInfo`

GetUpgradeInfo returns the UpgradeInfo field if non-nil, zero value otherwise.

### GetUpgradeInfoOk

`func (o *MachineCatalogResponseModel) GetUpgradeInfoOk() (*MachineCatalogUpgradeInfo, bool)`

GetUpgradeInfoOk returns a tuple with the UpgradeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeInfo

`func (o *MachineCatalogResponseModel) SetUpgradeInfo(v MachineCatalogUpgradeInfo)`

SetUpgradeInfo sets UpgradeInfo field to given value.

### HasUpgradeInfo

`func (o *MachineCatalogResponseModel) HasUpgradeInfo() bool`

HasUpgradeInfo returns a boolean if a field has been set.

### GetZone

`func (o *MachineCatalogResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *MachineCatalogResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *MachineCatalogResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.


### GetAdminFolder

`func (o *MachineCatalogResponseModel) GetAdminFolder() RefResponseModel`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *MachineCatalogResponseModel) GetAdminFolderOk() (*RefResponseModel, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *MachineCatalogResponseModel) SetAdminFolder(v RefResponseModel)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *MachineCatalogResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


