# CreateMachineCatalogRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminFolder** | Pointer to **string** | The admin folder in which the machine catalog should be created. | [optional] 
**Name** | **string** | Name of the machine catalog.  Each machine catalog within a site must have a unique name.  Required. | 
**AllocationType** | Pointer to [**AllocationType**](AllocationType.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the machine catalog.  Optional.  Default is no description. | [optional] 
**HypervisorConnection** | Pointer to **string** | Specifies the hypervisor connection to use for power management of machines in this machine catalog. | [optional] 
**IsRemotePC** | Pointer to **bool** | Indicates whether or not the catalog is a RemotePC catalog.  Default is &#x60;false&#x60;. | [optional] [default to false]
**RemotePCEnrollmentScopes** | Pointer to [**[]RemotePCEnrollmentScopeRequestModel**](RemotePCEnrollmentScopeRequestModel.md) | List of one or more remote PC enrollment scopes. | [optional] 
**MachineType** | Pointer to [**MachineType**](MachineType.md) |  | [optional] 
**MinimumFunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**PersistUserChanges** | Pointer to [**PersistChanges**](PersistChanges.md) |  | [optional] 
**ProvisioningType** | [**ProvisioningType**](ProvisioningType.md) |  | 
**ProvisioningScheme** | Pointer to [**CreateMachineCatalogRequestModelProvisioningScheme**](CreateMachineCatalogRequestModelProvisioningScheme.md) |  | [optional] 
**ProvisioningSchemeImportData** | Pointer to [**CreateMachineCatalogRequestModelProvisioningSchemeImportData**](CreateMachineCatalogRequestModelProvisioningSchemeImportData.md) |  | [optional] 
**Machines** | Pointer to [**[]AddMachineToMachineCatalogRequestModel**](AddMachineToMachineCatalogRequestModel.md) | List of machines to add to the catalog.  Must not be specified if  equals  or .  Optional otherwise.  Default is not to add any machines. | [optional] 
**PvsAddress** | Pointer to **string** | IP address of the PVS server to be used.  This only applies if the ProvisioningType is PVS. | [optional] 
**PvsDomain** | Pointer to **string** | The domain of the PVS server to be used. This only applies if the ProvisioningType is PVS. | [optional] 
**PvsCollectionIds** | Pointer to **[]string** | Collection IDs of PVS collections containing machines that should be added to the catalog.  This only applies if the ProvisioningType is PVS, and is required in that case.  Each item must be a valid PVS collection ID residing on the PVS server located at the specified . | [optional] 
**Scopes** | Pointer to **[]string** | Administrative scopes which the newly created machine catalog will be a part of. | [optional] 
**Tenants** | Pointer to **[]string** | Tenant(s) to associate the machine catalog with. | [optional] 
**SessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 
**VdaUpgradeType** | Pointer to [**VdaUpgradeType**](VdaUpgradeType.md) |  | [optional] 
**Zone** | Pointer to **string** | Zone the machine catalog is associated with.  Optional.  If not specified, the machine catalog is associated with the primary zone. See PrimaryZone. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of machine catalog. | [optional] 

## Methods

### NewCreateMachineCatalogRequestModel

`func NewCreateMachineCatalogRequestModel(name string, provisioningType ProvisioningType, sessionSupport SessionSupport, ) *CreateMachineCatalogRequestModel`

NewCreateMachineCatalogRequestModel instantiates a new CreateMachineCatalogRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMachineCatalogRequestModelWithDefaults

`func NewCreateMachineCatalogRequestModelWithDefaults() *CreateMachineCatalogRequestModel`

NewCreateMachineCatalogRequestModelWithDefaults instantiates a new CreateMachineCatalogRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminFolder

`func (o *CreateMachineCatalogRequestModel) GetAdminFolder() string`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *CreateMachineCatalogRequestModel) GetAdminFolderOk() (*string, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *CreateMachineCatalogRequestModel) SetAdminFolder(v string)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *CreateMachineCatalogRequestModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### GetName

`func (o *CreateMachineCatalogRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMachineCatalogRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMachineCatalogRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetAllocationType

`func (o *CreateMachineCatalogRequestModel) GetAllocationType() AllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *CreateMachineCatalogRequestModel) GetAllocationTypeOk() (*AllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *CreateMachineCatalogRequestModel) SetAllocationType(v AllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *CreateMachineCatalogRequestModel) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateMachineCatalogRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMachineCatalogRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMachineCatalogRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMachineCatalogRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *CreateMachineCatalogRequestModel) GetHypervisorConnection() string`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *CreateMachineCatalogRequestModel) GetHypervisorConnectionOk() (*string, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *CreateMachineCatalogRequestModel) SetHypervisorConnection(v string)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *CreateMachineCatalogRequestModel) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### GetIsRemotePC

`func (o *CreateMachineCatalogRequestModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *CreateMachineCatalogRequestModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *CreateMachineCatalogRequestModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.

### HasIsRemotePC

`func (o *CreateMachineCatalogRequestModel) HasIsRemotePC() bool`

HasIsRemotePC returns a boolean if a field has been set.

### GetRemotePCEnrollmentScopes

`func (o *CreateMachineCatalogRequestModel) GetRemotePCEnrollmentScopes() []RemotePCEnrollmentScopeRequestModel`

GetRemotePCEnrollmentScopes returns the RemotePCEnrollmentScopes field if non-nil, zero value otherwise.

### GetRemotePCEnrollmentScopesOk

`func (o *CreateMachineCatalogRequestModel) GetRemotePCEnrollmentScopesOk() (*[]RemotePCEnrollmentScopeRequestModel, bool)`

GetRemotePCEnrollmentScopesOk returns a tuple with the RemotePCEnrollmentScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePCEnrollmentScopes

`func (o *CreateMachineCatalogRequestModel) SetRemotePCEnrollmentScopes(v []RemotePCEnrollmentScopeRequestModel)`

SetRemotePCEnrollmentScopes sets RemotePCEnrollmentScopes field to given value.

### HasRemotePCEnrollmentScopes

`func (o *CreateMachineCatalogRequestModel) HasRemotePCEnrollmentScopes() bool`

HasRemotePCEnrollmentScopes returns a boolean if a field has been set.

### GetMachineType

`func (o *CreateMachineCatalogRequestModel) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *CreateMachineCatalogRequestModel) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *CreateMachineCatalogRequestModel) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *CreateMachineCatalogRequestModel) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetMinimumFunctionalLevel

`func (o *CreateMachineCatalogRequestModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *CreateMachineCatalogRequestModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *CreateMachineCatalogRequestModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.

### HasMinimumFunctionalLevel

`func (o *CreateMachineCatalogRequestModel) HasMinimumFunctionalLevel() bool`

HasMinimumFunctionalLevel returns a boolean if a field has been set.

### GetPersistUserChanges

`func (o *CreateMachineCatalogRequestModel) GetPersistUserChanges() PersistChanges`

GetPersistUserChanges returns the PersistUserChanges field if non-nil, zero value otherwise.

### GetPersistUserChangesOk

`func (o *CreateMachineCatalogRequestModel) GetPersistUserChangesOk() (*PersistChanges, bool)`

GetPersistUserChangesOk returns a tuple with the PersistUserChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistUserChanges

`func (o *CreateMachineCatalogRequestModel) SetPersistUserChanges(v PersistChanges)`

SetPersistUserChanges sets PersistUserChanges field to given value.

### HasPersistUserChanges

`func (o *CreateMachineCatalogRequestModel) HasPersistUserChanges() bool`

HasPersistUserChanges returns a boolean if a field has been set.

### GetProvisioningType

`func (o *CreateMachineCatalogRequestModel) GetProvisioningType() ProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *CreateMachineCatalogRequestModel) GetProvisioningTypeOk() (*ProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *CreateMachineCatalogRequestModel) SetProvisioningType(v ProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.


### GetProvisioningScheme

`func (o *CreateMachineCatalogRequestModel) GetProvisioningScheme() CreateMachineCatalogRequestModelProvisioningScheme`

GetProvisioningScheme returns the ProvisioningScheme field if non-nil, zero value otherwise.

### GetProvisioningSchemeOk

`func (o *CreateMachineCatalogRequestModel) GetProvisioningSchemeOk() (*CreateMachineCatalogRequestModelProvisioningScheme, bool)`

GetProvisioningSchemeOk returns a tuple with the ProvisioningScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningScheme

`func (o *CreateMachineCatalogRequestModel) SetProvisioningScheme(v CreateMachineCatalogRequestModelProvisioningScheme)`

SetProvisioningScheme sets ProvisioningScheme field to given value.

### HasProvisioningScheme

`func (o *CreateMachineCatalogRequestModel) HasProvisioningScheme() bool`

HasProvisioningScheme returns a boolean if a field has been set.

### GetProvisioningSchemeImportData

`func (o *CreateMachineCatalogRequestModel) GetProvisioningSchemeImportData() CreateMachineCatalogRequestModelProvisioningSchemeImportData`

GetProvisioningSchemeImportData returns the ProvisioningSchemeImportData field if non-nil, zero value otherwise.

### GetProvisioningSchemeImportDataOk

`func (o *CreateMachineCatalogRequestModel) GetProvisioningSchemeImportDataOk() (*CreateMachineCatalogRequestModelProvisioningSchemeImportData, bool)`

GetProvisioningSchemeImportDataOk returns a tuple with the ProvisioningSchemeImportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeImportData

`func (o *CreateMachineCatalogRequestModel) SetProvisioningSchemeImportData(v CreateMachineCatalogRequestModelProvisioningSchemeImportData)`

SetProvisioningSchemeImportData sets ProvisioningSchemeImportData field to given value.

### HasProvisioningSchemeImportData

`func (o *CreateMachineCatalogRequestModel) HasProvisioningSchemeImportData() bool`

HasProvisioningSchemeImportData returns a boolean if a field has been set.

### GetMachines

`func (o *CreateMachineCatalogRequestModel) GetMachines() []AddMachineToMachineCatalogRequestModel`

GetMachines returns the Machines field if non-nil, zero value otherwise.

### GetMachinesOk

`func (o *CreateMachineCatalogRequestModel) GetMachinesOk() (*[]AddMachineToMachineCatalogRequestModel, bool)`

GetMachinesOk returns a tuple with the Machines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachines

`func (o *CreateMachineCatalogRequestModel) SetMachines(v []AddMachineToMachineCatalogRequestModel)`

SetMachines sets Machines field to given value.

### HasMachines

`func (o *CreateMachineCatalogRequestModel) HasMachines() bool`

HasMachines returns a boolean if a field has been set.

### GetPvsAddress

`func (o *CreateMachineCatalogRequestModel) GetPvsAddress() string`

GetPvsAddress returns the PvsAddress field if non-nil, zero value otherwise.

### GetPvsAddressOk

`func (o *CreateMachineCatalogRequestModel) GetPvsAddressOk() (*string, bool)`

GetPvsAddressOk returns a tuple with the PvsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsAddress

`func (o *CreateMachineCatalogRequestModel) SetPvsAddress(v string)`

SetPvsAddress sets PvsAddress field to given value.

### HasPvsAddress

`func (o *CreateMachineCatalogRequestModel) HasPvsAddress() bool`

HasPvsAddress returns a boolean if a field has been set.

### GetPvsDomain

`func (o *CreateMachineCatalogRequestModel) GetPvsDomain() string`

GetPvsDomain returns the PvsDomain field if non-nil, zero value otherwise.

### GetPvsDomainOk

`func (o *CreateMachineCatalogRequestModel) GetPvsDomainOk() (*string, bool)`

GetPvsDomainOk returns a tuple with the PvsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsDomain

`func (o *CreateMachineCatalogRequestModel) SetPvsDomain(v string)`

SetPvsDomain sets PvsDomain field to given value.

### HasPvsDomain

`func (o *CreateMachineCatalogRequestModel) HasPvsDomain() bool`

HasPvsDomain returns a boolean if a field has been set.

### GetPvsCollectionIds

`func (o *CreateMachineCatalogRequestModel) GetPvsCollectionIds() []string`

GetPvsCollectionIds returns the PvsCollectionIds field if non-nil, zero value otherwise.

### GetPvsCollectionIdsOk

`func (o *CreateMachineCatalogRequestModel) GetPvsCollectionIdsOk() (*[]string, bool)`

GetPvsCollectionIdsOk returns a tuple with the PvsCollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsCollectionIds

`func (o *CreateMachineCatalogRequestModel) SetPvsCollectionIds(v []string)`

SetPvsCollectionIds sets PvsCollectionIds field to given value.

### HasPvsCollectionIds

`func (o *CreateMachineCatalogRequestModel) HasPvsCollectionIds() bool`

HasPvsCollectionIds returns a boolean if a field has been set.

### GetScopes

`func (o *CreateMachineCatalogRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateMachineCatalogRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateMachineCatalogRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateMachineCatalogRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTenants

`func (o *CreateMachineCatalogRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateMachineCatalogRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateMachineCatalogRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateMachineCatalogRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetSessionSupport

`func (o *CreateMachineCatalogRequestModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *CreateMachineCatalogRequestModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *CreateMachineCatalogRequestModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetVdaUpgradeType

`func (o *CreateMachineCatalogRequestModel) GetVdaUpgradeType() VdaUpgradeType`

GetVdaUpgradeType returns the VdaUpgradeType field if non-nil, zero value otherwise.

### GetVdaUpgradeTypeOk

`func (o *CreateMachineCatalogRequestModel) GetVdaUpgradeTypeOk() (*VdaUpgradeType, bool)`

GetVdaUpgradeTypeOk returns a tuple with the VdaUpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaUpgradeType

`func (o *CreateMachineCatalogRequestModel) SetVdaUpgradeType(v VdaUpgradeType)`

SetVdaUpgradeType sets VdaUpgradeType field to given value.

### HasVdaUpgradeType

`func (o *CreateMachineCatalogRequestModel) HasVdaUpgradeType() bool`

HasVdaUpgradeType returns a boolean if a field has been set.

### GetZone

`func (o *CreateMachineCatalogRequestModel) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CreateMachineCatalogRequestModel) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CreateMachineCatalogRequestModel) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CreateMachineCatalogRequestModel) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateMachineCatalogRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateMachineCatalogRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateMachineCatalogRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateMachineCatalogRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


