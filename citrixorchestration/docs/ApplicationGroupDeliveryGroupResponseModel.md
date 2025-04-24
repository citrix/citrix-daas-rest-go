# ApplicationGroupDeliveryGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Globally unique identifier of the delivery group. | 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED.&#x60; DEPRECATED.  Use Id. | [optional] 
**UserManagement** | Pointer to [**UserManagementModel**](UserManagementModel.md) |  | [optional] 
**Delivering** | [**DeliveryKind**](DeliveryKind.md) |  | 
**DeliveryType** | [**DeliveryKind**](DeliveryKind.md) |  | 
**Description** | Pointer to **NullableString** | A description for this delivery group useful for administrators of the site. | [optional] 
**DesktopsAvailable** | **int32** | Number of machines in the delivery group which are available to launch sessions on. | 
**DesktopsDisconnected** | **int32** | Number of desktops in the delivery group which have disconnected sessions. | 
**DesktopsFaulted** | **int32** | Number of desktops in the delivery group which are in a faulted state. | 
**DesktopsUnregistered** | **int32** | Number of desktops in the delivery group which have failed to register. | 
**Enabled** | **bool** | Whether the delivery group is enabled.  All resources published on a disabled delivery group do not appear to users. | 
**HasBeenPromoted** | **bool** | Whether the delivery group was previously promoted from a lower MinimumFunctionalLevel. | 
**HasBeenPromotedFrom** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**InMaintenanceMode** | **bool** | Whether the delivery group is in maintenance mode; a delivery group in maintenance mode will not allow users to connect or reconnect to machines in the delivery group. | 
**IsBroken** | **bool** | Whether the delivery group is currently in a \&quot;Broken\&quot; state. | 
**IsRemotePC** | **bool** | Whether the delivery group is comprised of RemotePC VDAs. | 
**MachineLogOnType** | Pointer to [**MachineLogOnType**](MachineLogOnType.md) |  | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of delivery group. | [optional] 
**MinimumFunctionalLevel** | [**FunctionalLevel**](FunctionalLevel.md) |  | 
**Name** | **string** | Simple administrative name of delivery group within parent admin folder (if any). This property is not guaranteed unique across all catalogs. | 
**FullName** | Pointer to **NullableString** | Unique administrative name of delivery group. | [optional] 
**PublishedName** | Pointer to **NullableString** | The name of the desktop group as it is to appear to the user in StoreFront. | [optional] 
**PolicySetGuid** | Pointer to **string** | The Guid of the policy set that is assigned to this desktop group. If the value is Guid.Empty, there is no policy set assigned to the desktop group. | [optional] 
**ProductCode** | Pointer to **NullableString** | The product code of the delivery group. | [optional] 
**LicenseModel** | Pointer to [**LicenseModel**](LicenseModel.md) |  | [optional] 
**RequireUserHomeZone** | **bool** | Whether the resources from this delivery group are required to launch within the user&#39;s home zone. | 
**Scopes** | [**[]ScopeResponseModel**](ScopeResponseModel.md) | Administrative scopes which the delivery group is part of. | 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the delivery group is assigned to.  If &#x60;null&#x60;, the delivery group is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**Tags** | Pointer to **[]string** | The tags directly associated with the delivery group. | [optional] 
**SessionCount** | **int32** | Number of sessions currently running on machines in the delivery group. | 
**SessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 
**SharingKind** | [**SharingKind**](SharingKind.md) |  | 
**TotalApplications** | **int32** | Total number of applications published on the delivery group. | 
**TotalDesktops** | **int32** | Total number of desktops in the delivery group. | 
**ApplicationGroupCompatibility** | [**AppGroupCompatibility**](AppGroupCompatibility.md) |  | 
**ApplicationCompatibility** | Pointer to [**AppOrDesktopCompatibility**](AppOrDesktopCompatibility.md) |  | [optional] 
**DesktopCompatibility** | Pointer to [**AppOrDesktopCompatibility**](AppOrDesktopCompatibility.md) |  | [optional] 
**RequiredSleepCapability** | Pointer to [**RequiredSleepCapability**](RequiredSleepCapability.md) |  | [optional] 
**AdminFolder** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**IsPowerManaged** | Pointer to **bool** | Indicates whether the machines in the delivery group are power-managed. NOTE: I used to think that MachineType&#x3D;&#x3D;Virtual meant the same thing as \&quot;power-managed\&quot;; however that&#39;s not the case.  A machine is power- managed if it is Virtual OR if it is RemotePC with a hypervisor connection (which will still have MachineType&#x3D;&#x3D;Physical). | [optional] 
**AutoscalingEnabled** | Pointer to **bool** | Specifies whether machines in this desktop group can be Autoscaled. | [optional] 
**ReuseMachinesWithoutShutdownInOutage** | Pointer to **bool** | Whether machines will be reused without a shutdown in case of an outage. | [optional] 
**TotalDesktopsOfSuspend** | Pointer to **NullableInt32** | Total number of suspend-capable desktops in the delivery group. | [optional] 
**DefaultDesktopIconId** | Pointer to **NullableString** | Default icon to use for desktops published from the delivery group. was IconUid | [optional] 
**Priority** | **int32** | Specifies the priority of the mapping between application and delivery group where lower numbers imply higher priority with zero being highest. | 
**NumMachines** | **int32** | Number of machines within the delivery group that are capable of hosting the applications in the application group. | 

## Methods

### NewApplicationGroupDeliveryGroupResponseModel

`func NewApplicationGroupDeliveryGroupResponseModel(id string, delivering DeliveryKind, deliveryType DeliveryKind, desktopsAvailable int32, desktopsDisconnected int32, desktopsFaulted int32, desktopsUnregistered int32, enabled bool, hasBeenPromoted bool, inMaintenanceMode bool, isBroken bool, isRemotePC bool, minimumFunctionalLevel FunctionalLevel, name string, requireUserHomeZone bool, scopes []ScopeResponseModel, sessionCount int32, sessionSupport SessionSupport, sharingKind SharingKind, totalApplications int32, totalDesktops int32, applicationGroupCompatibility AppGroupCompatibility, priority int32, numMachines int32, ) *ApplicationGroupDeliveryGroupResponseModel`

NewApplicationGroupDeliveryGroupResponseModel instantiates a new ApplicationGroupDeliveryGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupDeliveryGroupResponseModelWithDefaults

`func NewApplicationGroupDeliveryGroupResponseModelWithDefaults() *ApplicationGroupDeliveryGroupResponseModel`

NewApplicationGroupDeliveryGroupResponseModelWithDefaults instantiates a new ApplicationGroupDeliveryGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUserManagement

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetUserManagement() UserManagementModel`

GetUserManagement returns the UserManagement field if non-nil, zero value otherwise.

### GetUserManagementOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetUserManagementOk() (*UserManagementModel, bool)`

GetUserManagementOk returns a tuple with the UserManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagement

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetUserManagement(v UserManagementModel)`

SetUserManagement sets UserManagement field to given value.

### HasUserManagement

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasUserManagement() bool`

HasUserManagement returns a boolean if a field has been set.

### GetDelivering

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDelivering() DeliveryKind`

GetDelivering returns the Delivering field if non-nil, zero value otherwise.

### GetDeliveringOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDeliveringOk() (*DeliveryKind, bool)`

GetDeliveringOk returns a tuple with the Delivering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivering

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDelivering(v DeliveryKind)`

SetDelivering sets Delivering field to given value.


### GetDeliveryType

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.


### GetDescription

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationGroupDeliveryGroupResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesktopsAvailable

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopsAvailable() int32`

GetDesktopsAvailable returns the DesktopsAvailable field if non-nil, zero value otherwise.

### GetDesktopsAvailableOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopsAvailableOk() (*int32, bool)`

GetDesktopsAvailableOk returns a tuple with the DesktopsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsAvailable

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDesktopsAvailable(v int32)`

SetDesktopsAvailable sets DesktopsAvailable field to given value.


### GetDesktopsDisconnected

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopsDisconnected() int32`

GetDesktopsDisconnected returns the DesktopsDisconnected field if non-nil, zero value otherwise.

### GetDesktopsDisconnectedOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopsDisconnectedOk() (*int32, bool)`

GetDesktopsDisconnectedOk returns a tuple with the DesktopsDisconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsDisconnected

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDesktopsDisconnected(v int32)`

SetDesktopsDisconnected sets DesktopsDisconnected field to given value.


### GetDesktopsFaulted

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopsFaulted() int32`

GetDesktopsFaulted returns the DesktopsFaulted field if non-nil, zero value otherwise.

### GetDesktopsFaultedOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopsFaultedOk() (*int32, bool)`

GetDesktopsFaultedOk returns a tuple with the DesktopsFaulted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsFaulted

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDesktopsFaulted(v int32)`

SetDesktopsFaulted sets DesktopsFaulted field to given value.


### GetDesktopsUnregistered

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopsUnregistered() int32`

GetDesktopsUnregistered returns the DesktopsUnregistered field if non-nil, zero value otherwise.

### GetDesktopsUnregisteredOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopsUnregisteredOk() (*int32, bool)`

GetDesktopsUnregisteredOk returns a tuple with the DesktopsUnregistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsUnregistered

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDesktopsUnregistered(v int32)`

SetDesktopsUnregistered sets DesktopsUnregistered field to given value.


### GetEnabled

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasBeenPromoted

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetHasBeenPromoted() bool`

GetHasBeenPromoted returns the HasBeenPromoted field if non-nil, zero value otherwise.

### GetHasBeenPromotedOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetHasBeenPromotedOk() (*bool, bool)`

GetHasBeenPromotedOk returns a tuple with the HasBeenPromoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromoted

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetHasBeenPromoted(v bool)`

SetHasBeenPromoted sets HasBeenPromoted field to given value.


### GetHasBeenPromotedFrom

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetHasBeenPromotedFrom() FunctionalLevel`

GetHasBeenPromotedFrom returns the HasBeenPromotedFrom field if non-nil, zero value otherwise.

### GetHasBeenPromotedFromOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetHasBeenPromotedFromOk() (*FunctionalLevel, bool)`

GetHasBeenPromotedFromOk returns a tuple with the HasBeenPromotedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromotedFrom

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetHasBeenPromotedFrom(v FunctionalLevel)`

SetHasBeenPromotedFrom sets HasBeenPromotedFrom field to given value.

### HasHasBeenPromotedFrom

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasHasBeenPromotedFrom() bool`

HasHasBeenPromotedFrom returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetIsBroken

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.


### GetIsRemotePC

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.


### GetMachineLogOnType

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetMachineLogOnType() MachineLogOnType`

GetMachineLogOnType returns the MachineLogOnType field if non-nil, zero value otherwise.

### GetMachineLogOnTypeOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetMachineLogOnTypeOk() (*MachineLogOnType, bool)`

GetMachineLogOnTypeOk returns a tuple with the MachineLogOnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLogOnType

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetMachineLogOnType(v MachineLogOnType)`

SetMachineLogOnType sets MachineLogOnType field to given value.

### HasMachineLogOnType

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasMachineLogOnType() bool`

HasMachineLogOnType returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ApplicationGroupDeliveryGroupResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.


### GetName

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ApplicationGroupDeliveryGroupResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPublishedName

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *ApplicationGroupDeliveryGroupResponseModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetPolicySetGuid

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetPolicySetGuid() string`

GetPolicySetGuid returns the PolicySetGuid field if non-nil, zero value otherwise.

### GetPolicySetGuidOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetPolicySetGuidOk() (*string, bool)`

GetPolicySetGuidOk returns a tuple with the PolicySetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetGuid

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetPolicySetGuid(v string)`

SetPolicySetGuid sets PolicySetGuid field to given value.

### HasPolicySetGuid

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasPolicySetGuid() bool`

HasPolicySetGuid returns a boolean if a field has been set.

### GetProductCode

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCodeNil

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetProductCodeNil(b bool)`

 SetProductCodeNil sets the value for ProductCode to be an explicit nil

### UnsetProductCode
`func (o *ApplicationGroupDeliveryGroupResponseModel) UnsetProductCode()`

UnsetProductCode ensures that no value is present for ProductCode, not even an explicit nil
### GetLicenseModel

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetRequireUserHomeZone

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetRequireUserHomeZone() bool`

GetRequireUserHomeZone returns the RequireUserHomeZone field if non-nil, zero value otherwise.

### GetRequireUserHomeZoneOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetRequireUserHomeZoneOk() (*bool, bool)`

GetRequireUserHomeZoneOk returns a tuple with the RequireUserHomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireUserHomeZone

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetRequireUserHomeZone(v bool)`

SetRequireUserHomeZone sets RequireUserHomeZone field to given value.


### GetScopes

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ApplicationGroupDeliveryGroupResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetTags

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApplicationGroupDeliveryGroupResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSessionCount

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.


### GetSessionSupport

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSharingKind

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.


### GetTotalApplications

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTotalApplications() int32`

GetTotalApplications returns the TotalApplications field if non-nil, zero value otherwise.

### GetTotalApplicationsOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTotalApplicationsOk() (*int32, bool)`

GetTotalApplicationsOk returns a tuple with the TotalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplications

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetTotalApplications(v int32)`

SetTotalApplications sets TotalApplications field to given value.


### GetTotalDesktops

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTotalDesktops() int32`

GetTotalDesktops returns the TotalDesktops field if non-nil, zero value otherwise.

### GetTotalDesktopsOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTotalDesktopsOk() (*int32, bool)`

GetTotalDesktopsOk returns a tuple with the TotalDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDesktops

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetTotalDesktops(v int32)`

SetTotalDesktops sets TotalDesktops field to given value.


### GetApplicationGroupCompatibility

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetApplicationGroupCompatibility() AppGroupCompatibility`

GetApplicationGroupCompatibility returns the ApplicationGroupCompatibility field if non-nil, zero value otherwise.

### GetApplicationGroupCompatibilityOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetApplicationGroupCompatibilityOk() (*AppGroupCompatibility, bool)`

GetApplicationGroupCompatibilityOk returns a tuple with the ApplicationGroupCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroupCompatibility

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetApplicationGroupCompatibility(v AppGroupCompatibility)`

SetApplicationGroupCompatibility sets ApplicationGroupCompatibility field to given value.


### GetApplicationCompatibility

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetApplicationCompatibility() AppOrDesktopCompatibility`

GetApplicationCompatibility returns the ApplicationCompatibility field if non-nil, zero value otherwise.

### GetApplicationCompatibilityOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetApplicationCompatibilityOk() (*AppOrDesktopCompatibility, bool)`

GetApplicationCompatibilityOk returns a tuple with the ApplicationCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCompatibility

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetApplicationCompatibility(v AppOrDesktopCompatibility)`

SetApplicationCompatibility sets ApplicationCompatibility field to given value.

### HasApplicationCompatibility

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasApplicationCompatibility() bool`

HasApplicationCompatibility returns a boolean if a field has been set.

### GetDesktopCompatibility

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopCompatibility() AppOrDesktopCompatibility`

GetDesktopCompatibility returns the DesktopCompatibility field if non-nil, zero value otherwise.

### GetDesktopCompatibilityOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDesktopCompatibilityOk() (*AppOrDesktopCompatibility, bool)`

GetDesktopCompatibilityOk returns a tuple with the DesktopCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopCompatibility

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDesktopCompatibility(v AppOrDesktopCompatibility)`

SetDesktopCompatibility sets DesktopCompatibility field to given value.

### HasDesktopCompatibility

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasDesktopCompatibility() bool`

HasDesktopCompatibility returns a boolean if a field has been set.

### GetRequiredSleepCapability

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetRequiredSleepCapability() RequiredSleepCapability`

GetRequiredSleepCapability returns the RequiredSleepCapability field if non-nil, zero value otherwise.

### GetRequiredSleepCapabilityOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetRequiredSleepCapabilityOk() (*RequiredSleepCapability, bool)`

GetRequiredSleepCapabilityOk returns a tuple with the RequiredSleepCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSleepCapability

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetRequiredSleepCapability(v RequiredSleepCapability)`

SetRequiredSleepCapability sets RequiredSleepCapability field to given value.

### HasRequiredSleepCapability

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasRequiredSleepCapability() bool`

HasRequiredSleepCapability returns a boolean if a field has been set.

### GetAdminFolder

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetAdminFolder() RefResponseModel`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetAdminFolderOk() (*RefResponseModel, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetAdminFolder(v RefResponseModel)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### GetIsPowerManaged

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetIsPowerManaged() bool`

GetIsPowerManaged returns the IsPowerManaged field if non-nil, zero value otherwise.

### GetIsPowerManagedOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetIsPowerManagedOk() (*bool, bool)`

GetIsPowerManagedOk returns a tuple with the IsPowerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPowerManaged

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetIsPowerManaged(v bool)`

SetIsPowerManaged sets IsPowerManaged field to given value.

### HasIsPowerManaged

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasIsPowerManaged() bool`

HasIsPowerManaged returns a boolean if a field has been set.

### GetAutoscalingEnabled

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetAutoscalingEnabled() bool`

GetAutoscalingEnabled returns the AutoscalingEnabled field if non-nil, zero value otherwise.

### GetAutoscalingEnabledOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetAutoscalingEnabledOk() (*bool, bool)`

GetAutoscalingEnabledOk returns a tuple with the AutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingEnabled

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetAutoscalingEnabled(v bool)`

SetAutoscalingEnabled sets AutoscalingEnabled field to given value.

### HasAutoscalingEnabled

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasAutoscalingEnabled() bool`

HasAutoscalingEnabled returns a boolean if a field has been set.

### GetReuseMachinesWithoutShutdownInOutage

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetReuseMachinesWithoutShutdownInOutage() bool`

GetReuseMachinesWithoutShutdownInOutage returns the ReuseMachinesWithoutShutdownInOutage field if non-nil, zero value otherwise.

### GetReuseMachinesWithoutShutdownInOutageOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetReuseMachinesWithoutShutdownInOutageOk() (*bool, bool)`

GetReuseMachinesWithoutShutdownInOutageOk returns a tuple with the ReuseMachinesWithoutShutdownInOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseMachinesWithoutShutdownInOutage

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetReuseMachinesWithoutShutdownInOutage(v bool)`

SetReuseMachinesWithoutShutdownInOutage sets ReuseMachinesWithoutShutdownInOutage field to given value.

### HasReuseMachinesWithoutShutdownInOutage

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasReuseMachinesWithoutShutdownInOutage() bool`

HasReuseMachinesWithoutShutdownInOutage returns a boolean if a field has been set.

### GetTotalDesktopsOfSuspend

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTotalDesktopsOfSuspend() int32`

GetTotalDesktopsOfSuspend returns the TotalDesktopsOfSuspend field if non-nil, zero value otherwise.

### GetTotalDesktopsOfSuspendOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetTotalDesktopsOfSuspendOk() (*int32, bool)`

GetTotalDesktopsOfSuspendOk returns a tuple with the TotalDesktopsOfSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDesktopsOfSuspend

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetTotalDesktopsOfSuspend(v int32)`

SetTotalDesktopsOfSuspend sets TotalDesktopsOfSuspend field to given value.

### HasTotalDesktopsOfSuspend

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasTotalDesktopsOfSuspend() bool`

HasTotalDesktopsOfSuspend returns a boolean if a field has been set.

### SetTotalDesktopsOfSuspendNil

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetTotalDesktopsOfSuspendNil(b bool)`

 SetTotalDesktopsOfSuspendNil sets the value for TotalDesktopsOfSuspend to be an explicit nil

### UnsetTotalDesktopsOfSuspend
`func (o *ApplicationGroupDeliveryGroupResponseModel) UnsetTotalDesktopsOfSuspend()`

UnsetTotalDesktopsOfSuspend ensures that no value is present for TotalDesktopsOfSuspend, not even an explicit nil
### GetDefaultDesktopIconId

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDefaultDesktopIconId() string`

GetDefaultDesktopIconId returns the DefaultDesktopIconId field if non-nil, zero value otherwise.

### GetDefaultDesktopIconIdOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetDefaultDesktopIconIdOk() (*string, bool)`

GetDefaultDesktopIconIdOk returns a tuple with the DefaultDesktopIconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDesktopIconId

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDefaultDesktopIconId(v string)`

SetDefaultDesktopIconId sets DefaultDesktopIconId field to given value.

### HasDefaultDesktopIconId

`func (o *ApplicationGroupDeliveryGroupResponseModel) HasDefaultDesktopIconId() bool`

HasDefaultDesktopIconId returns a boolean if a field has been set.

### SetDefaultDesktopIconIdNil

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetDefaultDesktopIconIdNil(b bool)`

 SetDefaultDesktopIconIdNil sets the value for DefaultDesktopIconId to be an explicit nil

### UnsetDefaultDesktopIconId
`func (o *ApplicationGroupDeliveryGroupResponseModel) UnsetDefaultDesktopIconId()`

UnsetDefaultDesktopIconId ensures that no value is present for DefaultDesktopIconId, not even an explicit nil
### GetPriority

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetNumMachines

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *ApplicationGroupDeliveryGroupResponseModel) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *ApplicationGroupDeliveryGroupResponseModel) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


