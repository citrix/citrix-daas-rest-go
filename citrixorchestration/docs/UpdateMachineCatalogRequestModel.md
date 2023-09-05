# UpdateMachineCatalogRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminFolder** | Pointer to **string** | The folder in which the machine catalog resides. If not specified, the value will not be changed. May be specified as either the folder Id or Path. If specified as a path, and the path does not exist, it will be automatically created. | [optional] 
**Name** | Pointer to **string** | Name of the machine catalog.  Each machine catalog within a site must have a unique name.  If not specified, will not be changed. | [optional] 
**Description** | Pointer to **string** | Description of the machine catalog.  If not specified, will not be changed. | [optional] 
**HypervisorConnection** | Pointer to **string** | Specifies the hypervisor connection to use for power management of machines in this machine catalog.  If not specified, will not be changed. | [optional] 
**RemotePCEnrollmentScopes** | Pointer to [**[]RemotePCEnrollmentScopeRequestModel**](RemotePCEnrollmentScopeRequestModel.md) | List of one or more remote PC enrollment scopes.  If not specified, will not be changed. If specified, *all* enrollment scopes must be listed.  Existing enrollment scopes that are not listed will be removed.   Removing an enrollment scope will *not* remove existing remote PCs that were enrolled via that scope. | [optional] 
**MinimumFunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**PvsAddress** | Pointer to **string** | IP address of the PVS server to be used.  Can only be specified if  is . If not specified, will not be changed. | [optional] 
**PvsDomain** | Pointer to **string** | The domain of the PVS server to be used.  Can only be specified if  is . If not specified, will not be changed. | [optional] 
**Scopes** | Pointer to **[]string** | Administrative scopes which the machine catalog should be a part of.  The \&quot;All\&quot; scope, and any tenant scopes, are implicit and cannot be removed.  To remove from all non-implicit scopes, specify an empty array ([]).   Specifying tenant scopes is equivalent to specifying the  property and is subject to the same constraints. | [optional] 
**Tenants** | Pointer to **[]string** | Tenants to associate with the machine catalog. | [optional] 
**ShowRdsLicenseWarning** | Pointer to **bool** | Indicate whether the RDS license warning should appear for the machine catalog. | [optional] 
**Zone** | Pointer to **string** | Zone the machine catalog is associated with.  If not specified, will not be changed. | [optional] 
**CpuCount** | Pointer to **int32** | The number of processors that virtual machines created from the provisioning scheme should use. | [optional] 
**MemoryMB** | Pointer to **int32** | The maximum amount of memory that virtual machines created from the provisioning scheme should use. | [optional] 
**ServiceOfferingPath** | Pointer to **string** | The hypervisor resource path of the Cloud service offering to use when creating machines. | [optional] 
**NetworkMapping** | Pointer to [**[]NetworkMapRequestModel**](NetworkMapRequestModel.md) | Specifies how the attached NICs are mapped to networks. If this parameter is omitted, the current NICs setting is not updated. If an empty array is specified, new VMs will be created with a single NIC, which is mapped to the default network in the hosting unit. If an non-empty array is supplied, the NICs setting is updated, and new machines will be created with the number of NICs specified in the array, with each NIC attached to the specified network. | [optional] 
**VdaUpgradeType** | Pointer to [**VdaUpgradeType**](VdaUpgradeType.md) |  | [optional] 
**MachineProfilePath** | Pointer to **string** | The path in the resource pool to the virtual machine template that will be used. This identifies the VM template to be used and the default values for the tags, virtual machine size, boot diagnostics, host cache property of OS disk, accelerated networking and availability zone. This must be a path to a Virtual machine or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the provisioning scheme that are specific to the target hosting infrastructure. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of machine catalog. Set the value of the NameValueStringPairModel is null or empty will be remove this metadata. Not existing Name and Value NameValueStringPairModel object will be added. The same Name but different value object will be updated. | [optional] 
**AzureADSecurityGroupName** | Pointer to **string** | Name of Azure AD security group to be created. | [optional] 
**AzureADTenantId** | Pointer to **string** | TenantId of Azure Directory. | [optional] 

## Methods

### NewUpdateMachineCatalogRequestModel

`func NewUpdateMachineCatalogRequestModel() *UpdateMachineCatalogRequestModel`

NewUpdateMachineCatalogRequestModel instantiates a new UpdateMachineCatalogRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMachineCatalogRequestModelWithDefaults

`func NewUpdateMachineCatalogRequestModelWithDefaults() *UpdateMachineCatalogRequestModel`

NewUpdateMachineCatalogRequestModelWithDefaults instantiates a new UpdateMachineCatalogRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminFolder

`func (o *UpdateMachineCatalogRequestModel) GetAdminFolder() string`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *UpdateMachineCatalogRequestModel) GetAdminFolderOk() (*string, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *UpdateMachineCatalogRequestModel) SetAdminFolder(v string)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *UpdateMachineCatalogRequestModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### GetName

`func (o *UpdateMachineCatalogRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMachineCatalogRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMachineCatalogRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMachineCatalogRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateMachineCatalogRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMachineCatalogRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMachineCatalogRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMachineCatalogRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *UpdateMachineCatalogRequestModel) GetHypervisorConnection() string`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *UpdateMachineCatalogRequestModel) GetHypervisorConnectionOk() (*string, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *UpdateMachineCatalogRequestModel) SetHypervisorConnection(v string)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *UpdateMachineCatalogRequestModel) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### GetRemotePCEnrollmentScopes

`func (o *UpdateMachineCatalogRequestModel) GetRemotePCEnrollmentScopes() []RemotePCEnrollmentScopeRequestModel`

GetRemotePCEnrollmentScopes returns the RemotePCEnrollmentScopes field if non-nil, zero value otherwise.

### GetRemotePCEnrollmentScopesOk

`func (o *UpdateMachineCatalogRequestModel) GetRemotePCEnrollmentScopesOk() (*[]RemotePCEnrollmentScopeRequestModel, bool)`

GetRemotePCEnrollmentScopesOk returns a tuple with the RemotePCEnrollmentScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePCEnrollmentScopes

`func (o *UpdateMachineCatalogRequestModel) SetRemotePCEnrollmentScopes(v []RemotePCEnrollmentScopeRequestModel)`

SetRemotePCEnrollmentScopes sets RemotePCEnrollmentScopes field to given value.

### HasRemotePCEnrollmentScopes

`func (o *UpdateMachineCatalogRequestModel) HasRemotePCEnrollmentScopes() bool`

HasRemotePCEnrollmentScopes returns a boolean if a field has been set.

### GetMinimumFunctionalLevel

`func (o *UpdateMachineCatalogRequestModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *UpdateMachineCatalogRequestModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *UpdateMachineCatalogRequestModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.

### HasMinimumFunctionalLevel

`func (o *UpdateMachineCatalogRequestModel) HasMinimumFunctionalLevel() bool`

HasMinimumFunctionalLevel returns a boolean if a field has been set.

### GetPvsAddress

`func (o *UpdateMachineCatalogRequestModel) GetPvsAddress() string`

GetPvsAddress returns the PvsAddress field if non-nil, zero value otherwise.

### GetPvsAddressOk

`func (o *UpdateMachineCatalogRequestModel) GetPvsAddressOk() (*string, bool)`

GetPvsAddressOk returns a tuple with the PvsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsAddress

`func (o *UpdateMachineCatalogRequestModel) SetPvsAddress(v string)`

SetPvsAddress sets PvsAddress field to given value.

### HasPvsAddress

`func (o *UpdateMachineCatalogRequestModel) HasPvsAddress() bool`

HasPvsAddress returns a boolean if a field has been set.

### GetPvsDomain

`func (o *UpdateMachineCatalogRequestModel) GetPvsDomain() string`

GetPvsDomain returns the PvsDomain field if non-nil, zero value otherwise.

### GetPvsDomainOk

`func (o *UpdateMachineCatalogRequestModel) GetPvsDomainOk() (*string, bool)`

GetPvsDomainOk returns a tuple with the PvsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsDomain

`func (o *UpdateMachineCatalogRequestModel) SetPvsDomain(v string)`

SetPvsDomain sets PvsDomain field to given value.

### HasPvsDomain

`func (o *UpdateMachineCatalogRequestModel) HasPvsDomain() bool`

HasPvsDomain returns a boolean if a field has been set.

### GetScopes

`func (o *UpdateMachineCatalogRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateMachineCatalogRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateMachineCatalogRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateMachineCatalogRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateMachineCatalogRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateMachineCatalogRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateMachineCatalogRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateMachineCatalogRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetShowRdsLicenseWarning

`func (o *UpdateMachineCatalogRequestModel) GetShowRdsLicenseWarning() bool`

GetShowRdsLicenseWarning returns the ShowRdsLicenseWarning field if non-nil, zero value otherwise.

### GetShowRdsLicenseWarningOk

`func (o *UpdateMachineCatalogRequestModel) GetShowRdsLicenseWarningOk() (*bool, bool)`

GetShowRdsLicenseWarningOk returns a tuple with the ShowRdsLicenseWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRdsLicenseWarning

`func (o *UpdateMachineCatalogRequestModel) SetShowRdsLicenseWarning(v bool)`

SetShowRdsLicenseWarning sets ShowRdsLicenseWarning field to given value.

### HasShowRdsLicenseWarning

`func (o *UpdateMachineCatalogRequestModel) HasShowRdsLicenseWarning() bool`

HasShowRdsLicenseWarning returns a boolean if a field has been set.

### GetZone

`func (o *UpdateMachineCatalogRequestModel) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateMachineCatalogRequestModel) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateMachineCatalogRequestModel) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateMachineCatalogRequestModel) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetCpuCount

`func (o *UpdateMachineCatalogRequestModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *UpdateMachineCatalogRequestModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *UpdateMachineCatalogRequestModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *UpdateMachineCatalogRequestModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetMemoryMB

`func (o *UpdateMachineCatalogRequestModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *UpdateMachineCatalogRequestModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *UpdateMachineCatalogRequestModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *UpdateMachineCatalogRequestModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetServiceOfferingPath

`func (o *UpdateMachineCatalogRequestModel) GetServiceOfferingPath() string`

GetServiceOfferingPath returns the ServiceOfferingPath field if non-nil, zero value otherwise.

### GetServiceOfferingPathOk

`func (o *UpdateMachineCatalogRequestModel) GetServiceOfferingPathOk() (*string, bool)`

GetServiceOfferingPathOk returns a tuple with the ServiceOfferingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOfferingPath

`func (o *UpdateMachineCatalogRequestModel) SetServiceOfferingPath(v string)`

SetServiceOfferingPath sets ServiceOfferingPath field to given value.

### HasServiceOfferingPath

`func (o *UpdateMachineCatalogRequestModel) HasServiceOfferingPath() bool`

HasServiceOfferingPath returns a boolean if a field has been set.

### GetNetworkMapping

`func (o *UpdateMachineCatalogRequestModel) GetNetworkMapping() []NetworkMapRequestModel`

GetNetworkMapping returns the NetworkMapping field if non-nil, zero value otherwise.

### GetNetworkMappingOk

`func (o *UpdateMachineCatalogRequestModel) GetNetworkMappingOk() (*[]NetworkMapRequestModel, bool)`

GetNetworkMappingOk returns a tuple with the NetworkMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMapping

`func (o *UpdateMachineCatalogRequestModel) SetNetworkMapping(v []NetworkMapRequestModel)`

SetNetworkMapping sets NetworkMapping field to given value.

### HasNetworkMapping

`func (o *UpdateMachineCatalogRequestModel) HasNetworkMapping() bool`

HasNetworkMapping returns a boolean if a field has been set.

### GetVdaUpgradeType

`func (o *UpdateMachineCatalogRequestModel) GetVdaUpgradeType() VdaUpgradeType`

GetVdaUpgradeType returns the VdaUpgradeType field if non-nil, zero value otherwise.

### GetVdaUpgradeTypeOk

`func (o *UpdateMachineCatalogRequestModel) GetVdaUpgradeTypeOk() (*VdaUpgradeType, bool)`

GetVdaUpgradeTypeOk returns a tuple with the VdaUpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaUpgradeType

`func (o *UpdateMachineCatalogRequestModel) SetVdaUpgradeType(v VdaUpgradeType)`

SetVdaUpgradeType sets VdaUpgradeType field to given value.

### HasVdaUpgradeType

`func (o *UpdateMachineCatalogRequestModel) HasVdaUpgradeType() bool`

HasVdaUpgradeType returns a boolean if a field has been set.

### GetMachineProfilePath

`func (o *UpdateMachineCatalogRequestModel) GetMachineProfilePath() string`

GetMachineProfilePath returns the MachineProfilePath field if non-nil, zero value otherwise.

### GetMachineProfilePathOk

`func (o *UpdateMachineCatalogRequestModel) GetMachineProfilePathOk() (*string, bool)`

GetMachineProfilePathOk returns a tuple with the MachineProfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfilePath

`func (o *UpdateMachineCatalogRequestModel) SetMachineProfilePath(v string)`

SetMachineProfilePath sets MachineProfilePath field to given value.

### HasMachineProfilePath

`func (o *UpdateMachineCatalogRequestModel) HasMachineProfilePath() bool`

HasMachineProfilePath returns a boolean if a field has been set.

### GetCustomProperties

`func (o *UpdateMachineCatalogRequestModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *UpdateMachineCatalogRequestModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *UpdateMachineCatalogRequestModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *UpdateMachineCatalogRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateMachineCatalogRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateMachineCatalogRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateMachineCatalogRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateMachineCatalogRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAzureADSecurityGroupName

`func (o *UpdateMachineCatalogRequestModel) GetAzureADSecurityGroupName() string`

GetAzureADSecurityGroupName returns the AzureADSecurityGroupName field if non-nil, zero value otherwise.

### GetAzureADSecurityGroupNameOk

`func (o *UpdateMachineCatalogRequestModel) GetAzureADSecurityGroupNameOk() (*string, bool)`

GetAzureADSecurityGroupNameOk returns a tuple with the AzureADSecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADSecurityGroupName

`func (o *UpdateMachineCatalogRequestModel) SetAzureADSecurityGroupName(v string)`

SetAzureADSecurityGroupName sets AzureADSecurityGroupName field to given value.

### HasAzureADSecurityGroupName

`func (o *UpdateMachineCatalogRequestModel) HasAzureADSecurityGroupName() bool`

HasAzureADSecurityGroupName returns a boolean if a field has been set.

### GetAzureADTenantId

`func (o *UpdateMachineCatalogRequestModel) GetAzureADTenantId() string`

GetAzureADTenantId returns the AzureADTenantId field if non-nil, zero value otherwise.

### GetAzureADTenantIdOk

`func (o *UpdateMachineCatalogRequestModel) GetAzureADTenantIdOk() (*string, bool)`

GetAzureADTenantIdOk returns a tuple with the AzureADTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADTenantId

`func (o *UpdateMachineCatalogRequestModel) SetAzureADTenantId(v string)`

SetAzureADTenantId sets AzureADTenantId field to given value.

### HasAzureADTenantId

`func (o *UpdateMachineCatalogRequestModel) HasAzureADTenantId() bool`

HasAzureADTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


