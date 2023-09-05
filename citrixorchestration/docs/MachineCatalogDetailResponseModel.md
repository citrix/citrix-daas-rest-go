# MachineCatalogDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Simple administrative name of catalog within parent admin folder (if any). This property is not guaranteed unique across all catalogs. | 
**FullName** | Pointer to **string** | Unique administrative name of catalog. | [optional] 
**Id** | **string** | Id of the machine catalog. | 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED.  Use Id. | [optional] 
**AllocationType** | [**AllocationType**](AllocationType.md) |  | 
**AssignedCount** | Pointer to **int32** | The number of assigned machines (machines that have been assigned to a user/users or a client name/address). | [optional] 
**AvailableAssignedCount** | Pointer to **int32** | The number of available machines (not in a delivery group), that are also assigned to users. | [optional] 
**AvailableCount** | **int32** | The number of available machines (those not in any delivery group). | 
**AvailableUnassignedCount** | Pointer to **int32** | The number of available machines (those not in any delivery group) that are not assigned to users. | [optional] 
**Description** | Pointer to **string** | Description of the machine catalog. | [optional] 
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
**ProvisioningScheme** | Pointer to [**MachineCatalogResponseModelProvisioningScheme**](MachineCatalogResponseModelProvisioningScheme.md) |  | [optional] 
**ProvisioningType** | [**ProvisioningType**](ProvisioningType.md) |  | 
**ProvisioningProgress** | Pointer to [**MachineCatalogResponseModelProvisioningProgress**](MachineCatalogResponseModelProvisioningProgress.md) |  | [optional] 
**PvsAddress** | Pointer to **string** | IP address of the PVS server to be used. This only applies if the ProvisioningType is . | [optional] 
**PvsDomain** | Pointer to **string** | The domain of the PVS server to be used. | [optional] 
**RemotePCEnrollmentScopes** | Pointer to [**[]RemotePCEnrollmentScopeResponseModel**](RemotePCEnrollmentScopeResponseModel.md) | List of one or more remote PC enrollment scopes. | [optional] 
**Scopes** | Pointer to [**[]ScopeResponseModel**](ScopeResponseModel.md) | Administrative scopes which the machine catalog is part of. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the machine catalog is assigned to.  If &#x60;null&#x60;, the machine catalog is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**SessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 
**SharingKind** | [**SharingKind**](SharingKind.md) |  | 
**TotalCount** | **int32** | The total number of machines in the catalog. | 
**IsBroken** | **bool** | Whether the machine catalog is currently in a \&quot;Broken\&quot; state. | 
**IsMasterImageAssociated** | Pointer to **bool** | Whether the machine catalog is associated with a master image. | [optional] 
**Errors** | Pointer to **[]string** | Gets the Errors of machines in this catalog | [optional] 
**Warnings** | Pointer to [**[]MachineCatalogWarningResponseModel**](MachineCatalogWarningResponseModel.md) | List of warnings that are currently active on the machine catalog, if any.  If there are no warnings this will not be specified. | [optional] 
**UnassignedCount** | **int32** | The number of unassigned machines (machines not assigned to users). | 
**UsedCount** | **int32** | The number of machines in the catalog that are in a delivery group. | 
**UpgradeInfo** | Pointer to [**MachineCatalogResponseModelUpgradeInfo**](MachineCatalogResponseModelUpgradeInfo.md) |  | [optional] 
**Zone** | [**MachineCatalogResponseModelZone**](MachineCatalogResponseModelZone.md) |  | 
**AdminFolder** | Pointer to [**CatalogSearchResponseModelAdminFolder**](CatalogSearchResponseModelAdminFolder.md) |  | [optional] 
**AgentVersion** | Pointer to **string** | Version of the Citrix Virtual Delivery Agent (VDA) installed on the machine. | [optional] 
**HypervisorConnection** | Pointer to [**MachineCatalogDetailResponseModelAllOfHypervisorConnection**](MachineCatalogDetailResponseModelAllOfHypervisorConnection.md) |  | [optional] 
**OSType** | Pointer to **string** | A string that can be used to identify the operating system that is running on machines in the catalog. | [optional] 
**OSVersion** | Pointer to **string** | A string that can be used to identify the version of the operating system running on the machine, if known. | [optional] 
**DeliveryGroups** | [**[]MachineCatalogDeliveryGroupRefResponseModel**](MachineCatalogDeliveryGroupRefResponseModel.md) | Delivery groups associated with this catalog. | 
**UpgradeDetail** | Pointer to [**MachineCatalogDetailResponseModelAllOfUpgradeDetail**](MachineCatalogDetailResponseModelAllOfUpgradeDetail.md) |  | [optional] 

## Methods

### NewMachineCatalogDetailResponseModel

`func NewMachineCatalogDetailResponseModel(name string, id string, allocationType AllocationType, availableCount int32, isRemotePC bool, machineType MachineType, minimumFunctionalLevel FunctionalLevel, hasBeenPromoted bool, canRollbackVMImage bool, canRecreateCatalog bool, persistChanges PersistChanges, provisioningType ProvisioningType, sessionSupport SessionSupport, sharingKind SharingKind, totalCount int32, isBroken bool, unassignedCount int32, usedCount int32, zone MachineCatalogResponseModelZone, deliveryGroups []MachineCatalogDeliveryGroupRefResponseModel, ) *MachineCatalogDetailResponseModel`

NewMachineCatalogDetailResponseModel instantiates a new MachineCatalogDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogDetailResponseModelWithDefaults

`func NewMachineCatalogDetailResponseModelWithDefaults() *MachineCatalogDetailResponseModel`

NewMachineCatalogDetailResponseModelWithDefaults instantiates a new MachineCatalogDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MachineCatalogDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineCatalogDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineCatalogDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *MachineCatalogDetailResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MachineCatalogDetailResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MachineCatalogDetailResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MachineCatalogDetailResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetId

`func (o *MachineCatalogDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineCatalogDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineCatalogDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *MachineCatalogDetailResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MachineCatalogDetailResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MachineCatalogDetailResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MachineCatalogDetailResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetAllocationType

`func (o *MachineCatalogDetailResponseModel) GetAllocationType() AllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *MachineCatalogDetailResponseModel) GetAllocationTypeOk() (*AllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *MachineCatalogDetailResponseModel) SetAllocationType(v AllocationType)`

SetAllocationType sets AllocationType field to given value.


### GetAssignedCount

`func (o *MachineCatalogDetailResponseModel) GetAssignedCount() int32`

GetAssignedCount returns the AssignedCount field if non-nil, zero value otherwise.

### GetAssignedCountOk

`func (o *MachineCatalogDetailResponseModel) GetAssignedCountOk() (*int32, bool)`

GetAssignedCountOk returns a tuple with the AssignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCount

`func (o *MachineCatalogDetailResponseModel) SetAssignedCount(v int32)`

SetAssignedCount sets AssignedCount field to given value.

### HasAssignedCount

`func (o *MachineCatalogDetailResponseModel) HasAssignedCount() bool`

HasAssignedCount returns a boolean if a field has been set.

### GetAvailableAssignedCount

`func (o *MachineCatalogDetailResponseModel) GetAvailableAssignedCount() int32`

GetAvailableAssignedCount returns the AvailableAssignedCount field if non-nil, zero value otherwise.

### GetAvailableAssignedCountOk

`func (o *MachineCatalogDetailResponseModel) GetAvailableAssignedCountOk() (*int32, bool)`

GetAvailableAssignedCountOk returns a tuple with the AvailableAssignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAssignedCount

`func (o *MachineCatalogDetailResponseModel) SetAvailableAssignedCount(v int32)`

SetAvailableAssignedCount sets AvailableAssignedCount field to given value.

### HasAvailableAssignedCount

`func (o *MachineCatalogDetailResponseModel) HasAvailableAssignedCount() bool`

HasAvailableAssignedCount returns a boolean if a field has been set.

### GetAvailableCount

`func (o *MachineCatalogDetailResponseModel) GetAvailableCount() int32`

GetAvailableCount returns the AvailableCount field if non-nil, zero value otherwise.

### GetAvailableCountOk

`func (o *MachineCatalogDetailResponseModel) GetAvailableCountOk() (*int32, bool)`

GetAvailableCountOk returns a tuple with the AvailableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCount

`func (o *MachineCatalogDetailResponseModel) SetAvailableCount(v int32)`

SetAvailableCount sets AvailableCount field to given value.


### GetAvailableUnassignedCount

`func (o *MachineCatalogDetailResponseModel) GetAvailableUnassignedCount() int32`

GetAvailableUnassignedCount returns the AvailableUnassignedCount field if non-nil, zero value otherwise.

### GetAvailableUnassignedCountOk

`func (o *MachineCatalogDetailResponseModel) GetAvailableUnassignedCountOk() (*int32, bool)`

GetAvailableUnassignedCountOk returns a tuple with the AvailableUnassignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUnassignedCount

`func (o *MachineCatalogDetailResponseModel) SetAvailableUnassignedCount(v int32)`

SetAvailableUnassignedCount sets AvailableUnassignedCount field to given value.

### HasAvailableUnassignedCount

`func (o *MachineCatalogDetailResponseModel) HasAvailableUnassignedCount() bool`

HasAvailableUnassignedCount returns a boolean if a field has been set.

### GetDescription

`func (o *MachineCatalogDetailResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MachineCatalogDetailResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MachineCatalogDetailResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MachineCatalogDetailResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPowerManaged

`func (o *MachineCatalogDetailResponseModel) GetIsPowerManaged() bool`

GetIsPowerManaged returns the IsPowerManaged field if non-nil, zero value otherwise.

### GetIsPowerManagedOk

`func (o *MachineCatalogDetailResponseModel) GetIsPowerManagedOk() (*bool, bool)`

GetIsPowerManagedOk returns a tuple with the IsPowerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPowerManaged

`func (o *MachineCatalogDetailResponseModel) SetIsPowerManaged(v bool)`

SetIsPowerManaged sets IsPowerManaged field to given value.

### HasIsPowerManaged

`func (o *MachineCatalogDetailResponseModel) HasIsPowerManaged() bool`

HasIsPowerManaged returns a boolean if a field has been set.

### GetIsRemotePC

`func (o *MachineCatalogDetailResponseModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *MachineCatalogDetailResponseModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *MachineCatalogDetailResponseModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.


### GetJobsInProgress

`func (o *MachineCatalogDetailResponseModel) GetJobsInProgress() []RefResponseModel`

GetJobsInProgress returns the JobsInProgress field if non-nil, zero value otherwise.

### GetJobsInProgressOk

`func (o *MachineCatalogDetailResponseModel) GetJobsInProgressOk() (*[]RefResponseModel, bool)`

GetJobsInProgressOk returns a tuple with the JobsInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsInProgress

`func (o *MachineCatalogDetailResponseModel) SetJobsInProgress(v []RefResponseModel)`

SetJobsInProgress sets JobsInProgress field to given value.

### HasJobsInProgress

`func (o *MachineCatalogDetailResponseModel) HasJobsInProgress() bool`

HasJobsInProgress returns a boolean if a field has been set.

### GetMachineType

`func (o *MachineCatalogDetailResponseModel) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *MachineCatalogDetailResponseModel) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *MachineCatalogDetailResponseModel) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.


### GetMetadata

`func (o *MachineCatalogDetailResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MachineCatalogDetailResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MachineCatalogDetailResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MachineCatalogDetailResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMinimumFunctionalLevel

`func (o *MachineCatalogDetailResponseModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *MachineCatalogDetailResponseModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *MachineCatalogDetailResponseModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.


### GetHasBeenPromoted

`func (o *MachineCatalogDetailResponseModel) GetHasBeenPromoted() bool`

GetHasBeenPromoted returns the HasBeenPromoted field if non-nil, zero value otherwise.

### GetHasBeenPromotedOk

`func (o *MachineCatalogDetailResponseModel) GetHasBeenPromotedOk() (*bool, bool)`

GetHasBeenPromotedOk returns a tuple with the HasBeenPromoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromoted

`func (o *MachineCatalogDetailResponseModel) SetHasBeenPromoted(v bool)`

SetHasBeenPromoted sets HasBeenPromoted field to given value.


### GetHasBeenPromotedFrom

`func (o *MachineCatalogDetailResponseModel) GetHasBeenPromotedFrom() FunctionalLevel`

GetHasBeenPromotedFrom returns the HasBeenPromotedFrom field if non-nil, zero value otherwise.

### GetHasBeenPromotedFromOk

`func (o *MachineCatalogDetailResponseModel) GetHasBeenPromotedFromOk() (*FunctionalLevel, bool)`

GetHasBeenPromotedFromOk returns a tuple with the HasBeenPromotedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromotedFrom

`func (o *MachineCatalogDetailResponseModel) SetHasBeenPromotedFrom(v FunctionalLevel)`

SetHasBeenPromotedFrom sets HasBeenPromotedFrom field to given value.

### HasHasBeenPromotedFrom

`func (o *MachineCatalogDetailResponseModel) HasHasBeenPromotedFrom() bool`

HasHasBeenPromotedFrom returns a boolean if a field has been set.

### GetCanRollbackVMImage

`func (o *MachineCatalogDetailResponseModel) GetCanRollbackVMImage() bool`

GetCanRollbackVMImage returns the CanRollbackVMImage field if non-nil, zero value otherwise.

### GetCanRollbackVMImageOk

`func (o *MachineCatalogDetailResponseModel) GetCanRollbackVMImageOk() (*bool, bool)`

GetCanRollbackVMImageOk returns a tuple with the CanRollbackVMImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRollbackVMImage

`func (o *MachineCatalogDetailResponseModel) SetCanRollbackVMImage(v bool)`

SetCanRollbackVMImage sets CanRollbackVMImage field to given value.


### GetCanRecreateCatalog

`func (o *MachineCatalogDetailResponseModel) GetCanRecreateCatalog() bool`

GetCanRecreateCatalog returns the CanRecreateCatalog field if non-nil, zero value otherwise.

### GetCanRecreateCatalogOk

`func (o *MachineCatalogDetailResponseModel) GetCanRecreateCatalogOk() (*bool, bool)`

GetCanRecreateCatalogOk returns a tuple with the CanRecreateCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRecreateCatalog

`func (o *MachineCatalogDetailResponseModel) SetCanRecreateCatalog(v bool)`

SetCanRecreateCatalog sets CanRecreateCatalog field to given value.


### GetPersistChanges

`func (o *MachineCatalogDetailResponseModel) GetPersistChanges() PersistChanges`

GetPersistChanges returns the PersistChanges field if non-nil, zero value otherwise.

### GetPersistChangesOk

`func (o *MachineCatalogDetailResponseModel) GetPersistChangesOk() (*PersistChanges, bool)`

GetPersistChangesOk returns a tuple with the PersistChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistChanges

`func (o *MachineCatalogDetailResponseModel) SetPersistChanges(v PersistChanges)`

SetPersistChanges sets PersistChanges field to given value.


### GetProvisioningScheme

`func (o *MachineCatalogDetailResponseModel) GetProvisioningScheme() MachineCatalogResponseModelProvisioningScheme`

GetProvisioningScheme returns the ProvisioningScheme field if non-nil, zero value otherwise.

### GetProvisioningSchemeOk

`func (o *MachineCatalogDetailResponseModel) GetProvisioningSchemeOk() (*MachineCatalogResponseModelProvisioningScheme, bool)`

GetProvisioningSchemeOk returns a tuple with the ProvisioningScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningScheme

`func (o *MachineCatalogDetailResponseModel) SetProvisioningScheme(v MachineCatalogResponseModelProvisioningScheme)`

SetProvisioningScheme sets ProvisioningScheme field to given value.

### HasProvisioningScheme

`func (o *MachineCatalogDetailResponseModel) HasProvisioningScheme() bool`

HasProvisioningScheme returns a boolean if a field has been set.

### GetProvisioningType

`func (o *MachineCatalogDetailResponseModel) GetProvisioningType() ProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *MachineCatalogDetailResponseModel) GetProvisioningTypeOk() (*ProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *MachineCatalogDetailResponseModel) SetProvisioningType(v ProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.


### GetProvisioningProgress

`func (o *MachineCatalogDetailResponseModel) GetProvisioningProgress() MachineCatalogResponseModelProvisioningProgress`

GetProvisioningProgress returns the ProvisioningProgress field if non-nil, zero value otherwise.

### GetProvisioningProgressOk

`func (o *MachineCatalogDetailResponseModel) GetProvisioningProgressOk() (*MachineCatalogResponseModelProvisioningProgress, bool)`

GetProvisioningProgressOk returns a tuple with the ProvisioningProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProgress

`func (o *MachineCatalogDetailResponseModel) SetProvisioningProgress(v MachineCatalogResponseModelProvisioningProgress)`

SetProvisioningProgress sets ProvisioningProgress field to given value.

### HasProvisioningProgress

`func (o *MachineCatalogDetailResponseModel) HasProvisioningProgress() bool`

HasProvisioningProgress returns a boolean if a field has been set.

### GetPvsAddress

`func (o *MachineCatalogDetailResponseModel) GetPvsAddress() string`

GetPvsAddress returns the PvsAddress field if non-nil, zero value otherwise.

### GetPvsAddressOk

`func (o *MachineCatalogDetailResponseModel) GetPvsAddressOk() (*string, bool)`

GetPvsAddressOk returns a tuple with the PvsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsAddress

`func (o *MachineCatalogDetailResponseModel) SetPvsAddress(v string)`

SetPvsAddress sets PvsAddress field to given value.

### HasPvsAddress

`func (o *MachineCatalogDetailResponseModel) HasPvsAddress() bool`

HasPvsAddress returns a boolean if a field has been set.

### GetPvsDomain

`func (o *MachineCatalogDetailResponseModel) GetPvsDomain() string`

GetPvsDomain returns the PvsDomain field if non-nil, zero value otherwise.

### GetPvsDomainOk

`func (o *MachineCatalogDetailResponseModel) GetPvsDomainOk() (*string, bool)`

GetPvsDomainOk returns a tuple with the PvsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsDomain

`func (o *MachineCatalogDetailResponseModel) SetPvsDomain(v string)`

SetPvsDomain sets PvsDomain field to given value.

### HasPvsDomain

`func (o *MachineCatalogDetailResponseModel) HasPvsDomain() bool`

HasPvsDomain returns a boolean if a field has been set.

### GetRemotePCEnrollmentScopes

`func (o *MachineCatalogDetailResponseModel) GetRemotePCEnrollmentScopes() []RemotePCEnrollmentScopeResponseModel`

GetRemotePCEnrollmentScopes returns the RemotePCEnrollmentScopes field if non-nil, zero value otherwise.

### GetRemotePCEnrollmentScopesOk

`func (o *MachineCatalogDetailResponseModel) GetRemotePCEnrollmentScopesOk() (*[]RemotePCEnrollmentScopeResponseModel, bool)`

GetRemotePCEnrollmentScopesOk returns a tuple with the RemotePCEnrollmentScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePCEnrollmentScopes

`func (o *MachineCatalogDetailResponseModel) SetRemotePCEnrollmentScopes(v []RemotePCEnrollmentScopeResponseModel)`

SetRemotePCEnrollmentScopes sets RemotePCEnrollmentScopes field to given value.

### HasRemotePCEnrollmentScopes

`func (o *MachineCatalogDetailResponseModel) HasRemotePCEnrollmentScopes() bool`

HasRemotePCEnrollmentScopes returns a boolean if a field has been set.

### GetScopes

`func (o *MachineCatalogDetailResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *MachineCatalogDetailResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *MachineCatalogDetailResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *MachineCatalogDetailResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTenants

`func (o *MachineCatalogDetailResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *MachineCatalogDetailResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *MachineCatalogDetailResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *MachineCatalogDetailResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetSessionSupport

`func (o *MachineCatalogDetailResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *MachineCatalogDetailResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *MachineCatalogDetailResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSharingKind

`func (o *MachineCatalogDetailResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *MachineCatalogDetailResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *MachineCatalogDetailResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.


### GetTotalCount

`func (o *MachineCatalogDetailResponseModel) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *MachineCatalogDetailResponseModel) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *MachineCatalogDetailResponseModel) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetIsBroken

`func (o *MachineCatalogDetailResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *MachineCatalogDetailResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *MachineCatalogDetailResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.


### GetIsMasterImageAssociated

`func (o *MachineCatalogDetailResponseModel) GetIsMasterImageAssociated() bool`

GetIsMasterImageAssociated returns the IsMasterImageAssociated field if non-nil, zero value otherwise.

### GetIsMasterImageAssociatedOk

`func (o *MachineCatalogDetailResponseModel) GetIsMasterImageAssociatedOk() (*bool, bool)`

GetIsMasterImageAssociatedOk returns a tuple with the IsMasterImageAssociated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterImageAssociated

`func (o *MachineCatalogDetailResponseModel) SetIsMasterImageAssociated(v bool)`

SetIsMasterImageAssociated sets IsMasterImageAssociated field to given value.

### HasIsMasterImageAssociated

`func (o *MachineCatalogDetailResponseModel) HasIsMasterImageAssociated() bool`

HasIsMasterImageAssociated returns a boolean if a field has been set.

### GetErrors

`func (o *MachineCatalogDetailResponseModel) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MachineCatalogDetailResponseModel) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MachineCatalogDetailResponseModel) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MachineCatalogDetailResponseModel) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *MachineCatalogDetailResponseModel) GetWarnings() []MachineCatalogWarningResponseModel`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MachineCatalogDetailResponseModel) GetWarningsOk() (*[]MachineCatalogWarningResponseModel, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MachineCatalogDetailResponseModel) SetWarnings(v []MachineCatalogWarningResponseModel)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MachineCatalogDetailResponseModel) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetUnassignedCount

`func (o *MachineCatalogDetailResponseModel) GetUnassignedCount() int32`

GetUnassignedCount returns the UnassignedCount field if non-nil, zero value otherwise.

### GetUnassignedCountOk

`func (o *MachineCatalogDetailResponseModel) GetUnassignedCountOk() (*int32, bool)`

GetUnassignedCountOk returns a tuple with the UnassignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedCount

`func (o *MachineCatalogDetailResponseModel) SetUnassignedCount(v int32)`

SetUnassignedCount sets UnassignedCount field to given value.


### GetUsedCount

`func (o *MachineCatalogDetailResponseModel) GetUsedCount() int32`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *MachineCatalogDetailResponseModel) GetUsedCountOk() (*int32, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *MachineCatalogDetailResponseModel) SetUsedCount(v int32)`

SetUsedCount sets UsedCount field to given value.


### GetUpgradeInfo

`func (o *MachineCatalogDetailResponseModel) GetUpgradeInfo() MachineCatalogResponseModelUpgradeInfo`

GetUpgradeInfo returns the UpgradeInfo field if non-nil, zero value otherwise.

### GetUpgradeInfoOk

`func (o *MachineCatalogDetailResponseModel) GetUpgradeInfoOk() (*MachineCatalogResponseModelUpgradeInfo, bool)`

GetUpgradeInfoOk returns a tuple with the UpgradeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeInfo

`func (o *MachineCatalogDetailResponseModel) SetUpgradeInfo(v MachineCatalogResponseModelUpgradeInfo)`

SetUpgradeInfo sets UpgradeInfo field to given value.

### HasUpgradeInfo

`func (o *MachineCatalogDetailResponseModel) HasUpgradeInfo() bool`

HasUpgradeInfo returns a boolean if a field has been set.

### GetZone

`func (o *MachineCatalogDetailResponseModel) GetZone() MachineCatalogResponseModelZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *MachineCatalogDetailResponseModel) GetZoneOk() (*MachineCatalogResponseModelZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *MachineCatalogDetailResponseModel) SetZone(v MachineCatalogResponseModelZone)`

SetZone sets Zone field to given value.


### GetAdminFolder

`func (o *MachineCatalogDetailResponseModel) GetAdminFolder() CatalogSearchResponseModelAdminFolder`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *MachineCatalogDetailResponseModel) GetAdminFolderOk() (*CatalogSearchResponseModelAdminFolder, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *MachineCatalogDetailResponseModel) SetAdminFolder(v CatalogSearchResponseModelAdminFolder)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *MachineCatalogDetailResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### GetAgentVersion

`func (o *MachineCatalogDetailResponseModel) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *MachineCatalogDetailResponseModel) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *MachineCatalogDetailResponseModel) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *MachineCatalogDetailResponseModel) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *MachineCatalogDetailResponseModel) GetHypervisorConnection() MachineCatalogDetailResponseModelAllOfHypervisorConnection`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *MachineCatalogDetailResponseModel) GetHypervisorConnectionOk() (*MachineCatalogDetailResponseModelAllOfHypervisorConnection, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *MachineCatalogDetailResponseModel) SetHypervisorConnection(v MachineCatalogDetailResponseModelAllOfHypervisorConnection)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *MachineCatalogDetailResponseModel) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### GetOSType

`func (o *MachineCatalogDetailResponseModel) GetOSType() string`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *MachineCatalogDetailResponseModel) GetOSTypeOk() (*string, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *MachineCatalogDetailResponseModel) SetOSType(v string)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *MachineCatalogDetailResponseModel) HasOSType() bool`

HasOSType returns a boolean if a field has been set.

### GetOSVersion

`func (o *MachineCatalogDetailResponseModel) GetOSVersion() string`

GetOSVersion returns the OSVersion field if non-nil, zero value otherwise.

### GetOSVersionOk

`func (o *MachineCatalogDetailResponseModel) GetOSVersionOk() (*string, bool)`

GetOSVersionOk returns a tuple with the OSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSVersion

`func (o *MachineCatalogDetailResponseModel) SetOSVersion(v string)`

SetOSVersion sets OSVersion field to given value.

### HasOSVersion

`func (o *MachineCatalogDetailResponseModel) HasOSVersion() bool`

HasOSVersion returns a boolean if a field has been set.

### GetDeliveryGroups

`func (o *MachineCatalogDetailResponseModel) GetDeliveryGroups() []MachineCatalogDeliveryGroupRefResponseModel`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *MachineCatalogDetailResponseModel) GetDeliveryGroupsOk() (*[]MachineCatalogDeliveryGroupRefResponseModel, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *MachineCatalogDetailResponseModel) SetDeliveryGroups(v []MachineCatalogDeliveryGroupRefResponseModel)`

SetDeliveryGroups sets DeliveryGroups field to given value.


### GetUpgradeDetail

`func (o *MachineCatalogDetailResponseModel) GetUpgradeDetail() MachineCatalogDetailResponseModelAllOfUpgradeDetail`

GetUpgradeDetail returns the UpgradeDetail field if non-nil, zero value otherwise.

### GetUpgradeDetailOk

`func (o *MachineCatalogDetailResponseModel) GetUpgradeDetailOk() (*MachineCatalogDetailResponseModelAllOfUpgradeDetail, bool)`

GetUpgradeDetailOk returns a tuple with the UpgradeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDetail

`func (o *MachineCatalogDetailResponseModel) SetUpgradeDetail(v MachineCatalogDetailResponseModelAllOfUpgradeDetail)`

SetUpgradeDetail sets UpgradeDetail field to given value.

### HasUpgradeDetail

`func (o *MachineCatalogDetailResponseModel) HasUpgradeDetail() bool`

HasUpgradeDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


