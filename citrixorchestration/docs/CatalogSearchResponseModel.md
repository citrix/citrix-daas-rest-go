# CatalogSearchResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Simple administrative name of catalog within parent admin folder (if any). This property is not guaranteed unique across all catalogs. | [optional] 
**FullName** | Pointer to **NullableString** | Unique administrative name of catalog. | [optional] 
**Id** | Pointer to **NullableString** | Id of the machine catalog. | [optional] 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED.  Use Id. | [optional] 
**AllocationType** | Pointer to [**AllocationType**](AllocationType.md) |  | [optional] 
**HypervisorPluginResponse** | Pointer to [**HypervisorPluginResponseModel**](HypervisorPluginResponseModel.md) |  | [optional] 
**AssignedCount** | Pointer to **NullableInt32** | The number of assigned machines (machines that have been assigned to a user/users or a client name/address). | [optional] 
**AvailableCount** | Pointer to **int32** | The number of available machines (those not in any delivery group). | [optional] 
**Description** | Pointer to **NullableString** | Description of the machine catalog. | [optional] 
**HypervisorConnectionUid** | Pointer to **int32** | The Uid of the hypervisor connection that is associated with the machines in the catalog. This property only applies to MCS provisioned catalogs. For other provisioning types machines can be from one or more different hypervisor connections. | [optional] 
**IsPowerManaged** | Pointer to **bool** | Indicates whether the machines in the catalog are power-managed. | [optional] 
**IsRemotePC** | Pointer to **bool** | Indicates whether or not the catalog is a RemotePC catalog. Remote PC catalogs automatically configure appropriate machines without the need for manual configuration. CHANGE: was public bool RemotePC { get; set; } | [optional] 
**MachinesArePhysical** | Pointer to **bool** | Indicates whether or not machines in the machine catalog are Physical. | [optional] 
**ProvisioningType** | Pointer to [**ProvisioningType**](ProvisioningType.md) |  | [optional] 
**PvsAddress** | Pointer to **NullableString** | IP address of the PVS server to be used. This only applies if the ProvisioningType is . | [optional] 
**PvsDomain** | Pointer to **NullableString** | The domain of the PVS server to be used. | [optional] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) |  | [optional] 
**UnassignedCount** | Pointer to **int32** | The number of unassigned machines (machines not assigned to users). | [optional] 
**UsedCount** | Pointer to **int32** | The number of machines in the catalog that are in a delivery group. | [optional] 
**AdminFolder** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


