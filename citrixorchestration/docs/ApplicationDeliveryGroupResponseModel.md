# ApplicationDeliveryGroupResponseModel

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
**SessionCount** | **int32** | Number of sessions currently running on machines in the delivery group. | 
**SessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 
**SharingKind** | [**SharingKind**](SharingKind.md) |  | 
**TotalApplications** | **int32** | Total number of applications published on the delivery group. | 
**TotalDesktops** | **int32** | Total number of desktops in the delivery group. | 
**ApplicationGroupCompatibility** | [**AppGroupCompatibility**](AppGroupCompatibility.md) |  | 
**ApplicationCompatibility** | [**AppOrDesktopCompatibility**](AppOrDesktopCompatibility.md) |  | 
**DesktopCompatibility** | [**AppOrDesktopCompatibility**](AppOrDesktopCompatibility.md) |  | 
**RequiredSleepCapability** | Pointer to [**RequiredSleepCapability**](RequiredSleepCapability.md) |  | [optional] 
**AdminFolder** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**IsPowerManaged** | Pointer to **bool** | Indicates whether the machines in the delivery group are power-managed. NOTE: I used to think that MachineType&#x3D;&#x3D;Virtual meant the same thing as \&quot;power-managed\&quot;; however that&#39;s not the case.  A machine is power- managed if it is Virtual OR if it is RemotePC with a hypervisor connection (which will still have MachineType&#x3D;&#x3D;Physical). | [optional] 
**AutoscalingEnabled** | Pointer to **bool** | Specifies whether machines in this desktop group can be Autoscaled. | [optional] 
**ReuseMachinesWithoutShutdownInOutage** | Pointer to **bool** | Whether machines will be reused without a shutdown in case of an outage. | [optional] 
**Priority** | **int32** | Specifies the priority of the mapping between application and delivery group where lower numbers imply higher priority with zero being highest. | 

## Methods

### NewApplicationDeliveryGroupResponseModel

`func NewApplicationDeliveryGroupResponseModel(id string, delivering DeliveryKind, deliveryType DeliveryKind, desktopsAvailable int32, desktopsDisconnected int32, desktopsFaulted int32, desktopsUnregistered int32, enabled bool, hasBeenPromoted bool, inMaintenanceMode bool, isBroken bool, isRemotePC bool, minimumFunctionalLevel FunctionalLevel, name string, requireUserHomeZone bool, scopes []ScopeResponseModel, sessionCount int32, sessionSupport SessionSupport, sharingKind SharingKind, totalApplications int32, totalDesktops int32, applicationGroupCompatibility AppGroupCompatibility, applicationCompatibility AppOrDesktopCompatibility, desktopCompatibility AppOrDesktopCompatibility, priority int32, ) *ApplicationDeliveryGroupResponseModel`

NewApplicationDeliveryGroupResponseModel instantiates a new ApplicationDeliveryGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDeliveryGroupResponseModelWithDefaults

`func NewApplicationDeliveryGroupResponseModelWithDefaults() *ApplicationDeliveryGroupResponseModel`

NewApplicationDeliveryGroupResponseModelWithDefaults instantiates a new ApplicationDeliveryGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationDeliveryGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationDeliveryGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationDeliveryGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *ApplicationDeliveryGroupResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApplicationDeliveryGroupResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApplicationDeliveryGroupResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApplicationDeliveryGroupResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUserManagement

`func (o *ApplicationDeliveryGroupResponseModel) GetUserManagement() UserManagementModel`

GetUserManagement returns the UserManagement field if non-nil, zero value otherwise.

### GetUserManagementOk

`func (o *ApplicationDeliveryGroupResponseModel) GetUserManagementOk() (*UserManagementModel, bool)`

GetUserManagementOk returns a tuple with the UserManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagement

`func (o *ApplicationDeliveryGroupResponseModel) SetUserManagement(v UserManagementModel)`

SetUserManagement sets UserManagement field to given value.

### HasUserManagement

`func (o *ApplicationDeliveryGroupResponseModel) HasUserManagement() bool`

HasUserManagement returns a boolean if a field has been set.

### GetDelivering

`func (o *ApplicationDeliveryGroupResponseModel) GetDelivering() DeliveryKind`

GetDelivering returns the Delivering field if non-nil, zero value otherwise.

### GetDeliveringOk

`func (o *ApplicationDeliveryGroupResponseModel) GetDeliveringOk() (*DeliveryKind, bool)`

GetDeliveringOk returns a tuple with the Delivering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivering

`func (o *ApplicationDeliveryGroupResponseModel) SetDelivering(v DeliveryKind)`

SetDelivering sets Delivering field to given value.


### GetDeliveryType

`func (o *ApplicationDeliveryGroupResponseModel) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ApplicationDeliveryGroupResponseModel) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ApplicationDeliveryGroupResponseModel) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.


### GetDescription

`func (o *ApplicationDeliveryGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationDeliveryGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationDeliveryGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationDeliveryGroupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationDeliveryGroupResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationDeliveryGroupResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesktopsAvailable

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopsAvailable() int32`

GetDesktopsAvailable returns the DesktopsAvailable field if non-nil, zero value otherwise.

### GetDesktopsAvailableOk

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopsAvailableOk() (*int32, bool)`

GetDesktopsAvailableOk returns a tuple with the DesktopsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsAvailable

`func (o *ApplicationDeliveryGroupResponseModel) SetDesktopsAvailable(v int32)`

SetDesktopsAvailable sets DesktopsAvailable field to given value.


### GetDesktopsDisconnected

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopsDisconnected() int32`

GetDesktopsDisconnected returns the DesktopsDisconnected field if non-nil, zero value otherwise.

### GetDesktopsDisconnectedOk

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopsDisconnectedOk() (*int32, bool)`

GetDesktopsDisconnectedOk returns a tuple with the DesktopsDisconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsDisconnected

`func (o *ApplicationDeliveryGroupResponseModel) SetDesktopsDisconnected(v int32)`

SetDesktopsDisconnected sets DesktopsDisconnected field to given value.


### GetDesktopsFaulted

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopsFaulted() int32`

GetDesktopsFaulted returns the DesktopsFaulted field if non-nil, zero value otherwise.

### GetDesktopsFaultedOk

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopsFaultedOk() (*int32, bool)`

GetDesktopsFaultedOk returns a tuple with the DesktopsFaulted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsFaulted

`func (o *ApplicationDeliveryGroupResponseModel) SetDesktopsFaulted(v int32)`

SetDesktopsFaulted sets DesktopsFaulted field to given value.


### GetDesktopsUnregistered

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopsUnregistered() int32`

GetDesktopsUnregistered returns the DesktopsUnregistered field if non-nil, zero value otherwise.

### GetDesktopsUnregisteredOk

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopsUnregisteredOk() (*int32, bool)`

GetDesktopsUnregisteredOk returns a tuple with the DesktopsUnregistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsUnregistered

`func (o *ApplicationDeliveryGroupResponseModel) SetDesktopsUnregistered(v int32)`

SetDesktopsUnregistered sets DesktopsUnregistered field to given value.


### GetEnabled

`func (o *ApplicationDeliveryGroupResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationDeliveryGroupResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationDeliveryGroupResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasBeenPromoted

`func (o *ApplicationDeliveryGroupResponseModel) GetHasBeenPromoted() bool`

GetHasBeenPromoted returns the HasBeenPromoted field if non-nil, zero value otherwise.

### GetHasBeenPromotedOk

`func (o *ApplicationDeliveryGroupResponseModel) GetHasBeenPromotedOk() (*bool, bool)`

GetHasBeenPromotedOk returns a tuple with the HasBeenPromoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromoted

`func (o *ApplicationDeliveryGroupResponseModel) SetHasBeenPromoted(v bool)`

SetHasBeenPromoted sets HasBeenPromoted field to given value.


### GetHasBeenPromotedFrom

`func (o *ApplicationDeliveryGroupResponseModel) GetHasBeenPromotedFrom() FunctionalLevel`

GetHasBeenPromotedFrom returns the HasBeenPromotedFrom field if non-nil, zero value otherwise.

### GetHasBeenPromotedFromOk

`func (o *ApplicationDeliveryGroupResponseModel) GetHasBeenPromotedFromOk() (*FunctionalLevel, bool)`

GetHasBeenPromotedFromOk returns a tuple with the HasBeenPromotedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromotedFrom

`func (o *ApplicationDeliveryGroupResponseModel) SetHasBeenPromotedFrom(v FunctionalLevel)`

SetHasBeenPromotedFrom sets HasBeenPromotedFrom field to given value.

### HasHasBeenPromotedFrom

`func (o *ApplicationDeliveryGroupResponseModel) HasHasBeenPromotedFrom() bool`

HasHasBeenPromotedFrom returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *ApplicationDeliveryGroupResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *ApplicationDeliveryGroupResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *ApplicationDeliveryGroupResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetIsBroken

`func (o *ApplicationDeliveryGroupResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *ApplicationDeliveryGroupResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *ApplicationDeliveryGroupResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.


### GetIsRemotePC

`func (o *ApplicationDeliveryGroupResponseModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *ApplicationDeliveryGroupResponseModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *ApplicationDeliveryGroupResponseModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.


### GetMachineLogOnType

`func (o *ApplicationDeliveryGroupResponseModel) GetMachineLogOnType() MachineLogOnType`

GetMachineLogOnType returns the MachineLogOnType field if non-nil, zero value otherwise.

### GetMachineLogOnTypeOk

`func (o *ApplicationDeliveryGroupResponseModel) GetMachineLogOnTypeOk() (*MachineLogOnType, bool)`

GetMachineLogOnTypeOk returns a tuple with the MachineLogOnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLogOnType

`func (o *ApplicationDeliveryGroupResponseModel) SetMachineLogOnType(v MachineLogOnType)`

SetMachineLogOnType sets MachineLogOnType field to given value.

### HasMachineLogOnType

`func (o *ApplicationDeliveryGroupResponseModel) HasMachineLogOnType() bool`

HasMachineLogOnType returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationDeliveryGroupResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationDeliveryGroupResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationDeliveryGroupResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationDeliveryGroupResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ApplicationDeliveryGroupResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ApplicationDeliveryGroupResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *ApplicationDeliveryGroupResponseModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *ApplicationDeliveryGroupResponseModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *ApplicationDeliveryGroupResponseModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.


### GetName

`func (o *ApplicationDeliveryGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationDeliveryGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationDeliveryGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *ApplicationDeliveryGroupResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ApplicationDeliveryGroupResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ApplicationDeliveryGroupResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ApplicationDeliveryGroupResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ApplicationDeliveryGroupResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ApplicationDeliveryGroupResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPublishedName

`func (o *ApplicationDeliveryGroupResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *ApplicationDeliveryGroupResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *ApplicationDeliveryGroupResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *ApplicationDeliveryGroupResponseModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *ApplicationDeliveryGroupResponseModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *ApplicationDeliveryGroupResponseModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetPolicySetGuid

`func (o *ApplicationDeliveryGroupResponseModel) GetPolicySetGuid() string`

GetPolicySetGuid returns the PolicySetGuid field if non-nil, zero value otherwise.

### GetPolicySetGuidOk

`func (o *ApplicationDeliveryGroupResponseModel) GetPolicySetGuidOk() (*string, bool)`

GetPolicySetGuidOk returns a tuple with the PolicySetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetGuid

`func (o *ApplicationDeliveryGroupResponseModel) SetPolicySetGuid(v string)`

SetPolicySetGuid sets PolicySetGuid field to given value.

### HasPolicySetGuid

`func (o *ApplicationDeliveryGroupResponseModel) HasPolicySetGuid() bool`

HasPolicySetGuid returns a boolean if a field has been set.

### GetProductCode

`func (o *ApplicationDeliveryGroupResponseModel) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ApplicationDeliveryGroupResponseModel) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ApplicationDeliveryGroupResponseModel) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *ApplicationDeliveryGroupResponseModel) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCodeNil

`func (o *ApplicationDeliveryGroupResponseModel) SetProductCodeNil(b bool)`

 SetProductCodeNil sets the value for ProductCode to be an explicit nil

### UnsetProductCode
`func (o *ApplicationDeliveryGroupResponseModel) UnsetProductCode()`

UnsetProductCode ensures that no value is present for ProductCode, not even an explicit nil
### GetLicenseModel

`func (o *ApplicationDeliveryGroupResponseModel) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *ApplicationDeliveryGroupResponseModel) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *ApplicationDeliveryGroupResponseModel) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *ApplicationDeliveryGroupResponseModel) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetRequireUserHomeZone

`func (o *ApplicationDeliveryGroupResponseModel) GetRequireUserHomeZone() bool`

GetRequireUserHomeZone returns the RequireUserHomeZone field if non-nil, zero value otherwise.

### GetRequireUserHomeZoneOk

`func (o *ApplicationDeliveryGroupResponseModel) GetRequireUserHomeZoneOk() (*bool, bool)`

GetRequireUserHomeZoneOk returns a tuple with the RequireUserHomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireUserHomeZone

`func (o *ApplicationDeliveryGroupResponseModel) SetRequireUserHomeZone(v bool)`

SetRequireUserHomeZone sets RequireUserHomeZone field to given value.


### GetScopes

`func (o *ApplicationDeliveryGroupResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationDeliveryGroupResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationDeliveryGroupResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *ApplicationDeliveryGroupResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApplicationDeliveryGroupResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApplicationDeliveryGroupResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApplicationDeliveryGroupResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ApplicationDeliveryGroupResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ApplicationDeliveryGroupResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetSessionCount

`func (o *ApplicationDeliveryGroupResponseModel) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *ApplicationDeliveryGroupResponseModel) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *ApplicationDeliveryGroupResponseModel) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.


### GetSessionSupport

`func (o *ApplicationDeliveryGroupResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *ApplicationDeliveryGroupResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *ApplicationDeliveryGroupResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSharingKind

`func (o *ApplicationDeliveryGroupResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *ApplicationDeliveryGroupResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *ApplicationDeliveryGroupResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.


### GetTotalApplications

`func (o *ApplicationDeliveryGroupResponseModel) GetTotalApplications() int32`

GetTotalApplications returns the TotalApplications field if non-nil, zero value otherwise.

### GetTotalApplicationsOk

`func (o *ApplicationDeliveryGroupResponseModel) GetTotalApplicationsOk() (*int32, bool)`

GetTotalApplicationsOk returns a tuple with the TotalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplications

`func (o *ApplicationDeliveryGroupResponseModel) SetTotalApplications(v int32)`

SetTotalApplications sets TotalApplications field to given value.


### GetTotalDesktops

`func (o *ApplicationDeliveryGroupResponseModel) GetTotalDesktops() int32`

GetTotalDesktops returns the TotalDesktops field if non-nil, zero value otherwise.

### GetTotalDesktopsOk

`func (o *ApplicationDeliveryGroupResponseModel) GetTotalDesktopsOk() (*int32, bool)`

GetTotalDesktopsOk returns a tuple with the TotalDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDesktops

`func (o *ApplicationDeliveryGroupResponseModel) SetTotalDesktops(v int32)`

SetTotalDesktops sets TotalDesktops field to given value.


### GetApplicationGroupCompatibility

`func (o *ApplicationDeliveryGroupResponseModel) GetApplicationGroupCompatibility() AppGroupCompatibility`

GetApplicationGroupCompatibility returns the ApplicationGroupCompatibility field if non-nil, zero value otherwise.

### GetApplicationGroupCompatibilityOk

`func (o *ApplicationDeliveryGroupResponseModel) GetApplicationGroupCompatibilityOk() (*AppGroupCompatibility, bool)`

GetApplicationGroupCompatibilityOk returns a tuple with the ApplicationGroupCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroupCompatibility

`func (o *ApplicationDeliveryGroupResponseModel) SetApplicationGroupCompatibility(v AppGroupCompatibility)`

SetApplicationGroupCompatibility sets ApplicationGroupCompatibility field to given value.


### GetApplicationCompatibility

`func (o *ApplicationDeliveryGroupResponseModel) GetApplicationCompatibility() AppOrDesktopCompatibility`

GetApplicationCompatibility returns the ApplicationCompatibility field if non-nil, zero value otherwise.

### GetApplicationCompatibilityOk

`func (o *ApplicationDeliveryGroupResponseModel) GetApplicationCompatibilityOk() (*AppOrDesktopCompatibility, bool)`

GetApplicationCompatibilityOk returns a tuple with the ApplicationCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCompatibility

`func (o *ApplicationDeliveryGroupResponseModel) SetApplicationCompatibility(v AppOrDesktopCompatibility)`

SetApplicationCompatibility sets ApplicationCompatibility field to given value.


### GetDesktopCompatibility

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopCompatibility() AppOrDesktopCompatibility`

GetDesktopCompatibility returns the DesktopCompatibility field if non-nil, zero value otherwise.

### GetDesktopCompatibilityOk

`func (o *ApplicationDeliveryGroupResponseModel) GetDesktopCompatibilityOk() (*AppOrDesktopCompatibility, bool)`

GetDesktopCompatibilityOk returns a tuple with the DesktopCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopCompatibility

`func (o *ApplicationDeliveryGroupResponseModel) SetDesktopCompatibility(v AppOrDesktopCompatibility)`

SetDesktopCompatibility sets DesktopCompatibility field to given value.


### GetRequiredSleepCapability

`func (o *ApplicationDeliveryGroupResponseModel) GetRequiredSleepCapability() RequiredSleepCapability`

GetRequiredSleepCapability returns the RequiredSleepCapability field if non-nil, zero value otherwise.

### GetRequiredSleepCapabilityOk

`func (o *ApplicationDeliveryGroupResponseModel) GetRequiredSleepCapabilityOk() (*RequiredSleepCapability, bool)`

GetRequiredSleepCapabilityOk returns a tuple with the RequiredSleepCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSleepCapability

`func (o *ApplicationDeliveryGroupResponseModel) SetRequiredSleepCapability(v RequiredSleepCapability)`

SetRequiredSleepCapability sets RequiredSleepCapability field to given value.

### HasRequiredSleepCapability

`func (o *ApplicationDeliveryGroupResponseModel) HasRequiredSleepCapability() bool`

HasRequiredSleepCapability returns a boolean if a field has been set.

### GetAdminFolder

`func (o *ApplicationDeliveryGroupResponseModel) GetAdminFolder() RefResponseModel`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *ApplicationDeliveryGroupResponseModel) GetAdminFolderOk() (*RefResponseModel, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *ApplicationDeliveryGroupResponseModel) SetAdminFolder(v RefResponseModel)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *ApplicationDeliveryGroupResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### GetIsPowerManaged

`func (o *ApplicationDeliveryGroupResponseModel) GetIsPowerManaged() bool`

GetIsPowerManaged returns the IsPowerManaged field if non-nil, zero value otherwise.

### GetIsPowerManagedOk

`func (o *ApplicationDeliveryGroupResponseModel) GetIsPowerManagedOk() (*bool, bool)`

GetIsPowerManagedOk returns a tuple with the IsPowerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPowerManaged

`func (o *ApplicationDeliveryGroupResponseModel) SetIsPowerManaged(v bool)`

SetIsPowerManaged sets IsPowerManaged field to given value.

### HasIsPowerManaged

`func (o *ApplicationDeliveryGroupResponseModel) HasIsPowerManaged() bool`

HasIsPowerManaged returns a boolean if a field has been set.

### GetAutoscalingEnabled

`func (o *ApplicationDeliveryGroupResponseModel) GetAutoscalingEnabled() bool`

GetAutoscalingEnabled returns the AutoscalingEnabled field if non-nil, zero value otherwise.

### GetAutoscalingEnabledOk

`func (o *ApplicationDeliveryGroupResponseModel) GetAutoscalingEnabledOk() (*bool, bool)`

GetAutoscalingEnabledOk returns a tuple with the AutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingEnabled

`func (o *ApplicationDeliveryGroupResponseModel) SetAutoscalingEnabled(v bool)`

SetAutoscalingEnabled sets AutoscalingEnabled field to given value.

### HasAutoscalingEnabled

`func (o *ApplicationDeliveryGroupResponseModel) HasAutoscalingEnabled() bool`

HasAutoscalingEnabled returns a boolean if a field has been set.

### GetReuseMachinesWithoutShutdownInOutage

`func (o *ApplicationDeliveryGroupResponseModel) GetReuseMachinesWithoutShutdownInOutage() bool`

GetReuseMachinesWithoutShutdownInOutage returns the ReuseMachinesWithoutShutdownInOutage field if non-nil, zero value otherwise.

### GetReuseMachinesWithoutShutdownInOutageOk

`func (o *ApplicationDeliveryGroupResponseModel) GetReuseMachinesWithoutShutdownInOutageOk() (*bool, bool)`

GetReuseMachinesWithoutShutdownInOutageOk returns a tuple with the ReuseMachinesWithoutShutdownInOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseMachinesWithoutShutdownInOutage

`func (o *ApplicationDeliveryGroupResponseModel) SetReuseMachinesWithoutShutdownInOutage(v bool)`

SetReuseMachinesWithoutShutdownInOutage sets ReuseMachinesWithoutShutdownInOutage field to given value.

### HasReuseMachinesWithoutShutdownInOutage

`func (o *ApplicationDeliveryGroupResponseModel) HasReuseMachinesWithoutShutdownInOutage() bool`

HasReuseMachinesWithoutShutdownInOutage returns a boolean if a field has been set.

### GetPriority

`func (o *ApplicationDeliveryGroupResponseModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApplicationDeliveryGroupResponseModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApplicationDeliveryGroupResponseModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


