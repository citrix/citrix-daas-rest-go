# DeliveryGroupResponseModel

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
**AdminFolder** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 

## Methods

### NewDeliveryGroupResponseModel

`func NewDeliveryGroupResponseModel(id string, delivering DeliveryKind, deliveryType DeliveryKind, desktopsAvailable int32, desktopsDisconnected int32, desktopsFaulted int32, desktopsUnregistered int32, enabled bool, hasBeenPromoted bool, inMaintenanceMode bool, isBroken bool, isRemotePC bool, minimumFunctionalLevel FunctionalLevel, name string, requireUserHomeZone bool, scopes []ScopeResponseModel, sessionCount int32, sessionSupport SessionSupport, sharingKind SharingKind, totalApplications int32, totalDesktops int32, applicationGroupCompatibility AppGroupCompatibility, applicationCompatibility AppOrDesktopCompatibility, desktopCompatibility AppOrDesktopCompatibility, ) *DeliveryGroupResponseModel`

NewDeliveryGroupResponseModel instantiates a new DeliveryGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupResponseModelWithDefaults

`func NewDeliveryGroupResponseModelWithDefaults() *DeliveryGroupResponseModel`

NewDeliveryGroupResponseModelWithDefaults instantiates a new DeliveryGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *DeliveryGroupResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DeliveryGroupResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DeliveryGroupResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DeliveryGroupResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUserManagement

`func (o *DeliveryGroupResponseModel) GetUserManagement() UserManagementModel`

GetUserManagement returns the UserManagement field if non-nil, zero value otherwise.

### GetUserManagementOk

`func (o *DeliveryGroupResponseModel) GetUserManagementOk() (*UserManagementModel, bool)`

GetUserManagementOk returns a tuple with the UserManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagement

`func (o *DeliveryGroupResponseModel) SetUserManagement(v UserManagementModel)`

SetUserManagement sets UserManagement field to given value.

### HasUserManagement

`func (o *DeliveryGroupResponseModel) HasUserManagement() bool`

HasUserManagement returns a boolean if a field has been set.

### GetDelivering

`func (o *DeliveryGroupResponseModel) GetDelivering() DeliveryKind`

GetDelivering returns the Delivering field if non-nil, zero value otherwise.

### GetDeliveringOk

`func (o *DeliveryGroupResponseModel) GetDeliveringOk() (*DeliveryKind, bool)`

GetDeliveringOk returns a tuple with the Delivering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivering

`func (o *DeliveryGroupResponseModel) SetDelivering(v DeliveryKind)`

SetDelivering sets Delivering field to given value.


### GetDeliveryType

`func (o *DeliveryGroupResponseModel) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *DeliveryGroupResponseModel) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *DeliveryGroupResponseModel) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.


### GetDescription

`func (o *DeliveryGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeliveryGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeliveryGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeliveryGroupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DeliveryGroupResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeliveryGroupResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesktopsAvailable

`func (o *DeliveryGroupResponseModel) GetDesktopsAvailable() int32`

GetDesktopsAvailable returns the DesktopsAvailable field if non-nil, zero value otherwise.

### GetDesktopsAvailableOk

`func (o *DeliveryGroupResponseModel) GetDesktopsAvailableOk() (*int32, bool)`

GetDesktopsAvailableOk returns a tuple with the DesktopsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsAvailable

`func (o *DeliveryGroupResponseModel) SetDesktopsAvailable(v int32)`

SetDesktopsAvailable sets DesktopsAvailable field to given value.


### GetDesktopsDisconnected

`func (o *DeliveryGroupResponseModel) GetDesktopsDisconnected() int32`

GetDesktopsDisconnected returns the DesktopsDisconnected field if non-nil, zero value otherwise.

### GetDesktopsDisconnectedOk

`func (o *DeliveryGroupResponseModel) GetDesktopsDisconnectedOk() (*int32, bool)`

GetDesktopsDisconnectedOk returns a tuple with the DesktopsDisconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsDisconnected

`func (o *DeliveryGroupResponseModel) SetDesktopsDisconnected(v int32)`

SetDesktopsDisconnected sets DesktopsDisconnected field to given value.


### GetDesktopsFaulted

`func (o *DeliveryGroupResponseModel) GetDesktopsFaulted() int32`

GetDesktopsFaulted returns the DesktopsFaulted field if non-nil, zero value otherwise.

### GetDesktopsFaultedOk

`func (o *DeliveryGroupResponseModel) GetDesktopsFaultedOk() (*int32, bool)`

GetDesktopsFaultedOk returns a tuple with the DesktopsFaulted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsFaulted

`func (o *DeliveryGroupResponseModel) SetDesktopsFaulted(v int32)`

SetDesktopsFaulted sets DesktopsFaulted field to given value.


### GetDesktopsUnregistered

`func (o *DeliveryGroupResponseModel) GetDesktopsUnregistered() int32`

GetDesktopsUnregistered returns the DesktopsUnregistered field if non-nil, zero value otherwise.

### GetDesktopsUnregisteredOk

`func (o *DeliveryGroupResponseModel) GetDesktopsUnregisteredOk() (*int32, bool)`

GetDesktopsUnregisteredOk returns a tuple with the DesktopsUnregistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsUnregistered

`func (o *DeliveryGroupResponseModel) SetDesktopsUnregistered(v int32)`

SetDesktopsUnregistered sets DesktopsUnregistered field to given value.


### GetEnabled

`func (o *DeliveryGroupResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeliveryGroupResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeliveryGroupResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasBeenPromoted

`func (o *DeliveryGroupResponseModel) GetHasBeenPromoted() bool`

GetHasBeenPromoted returns the HasBeenPromoted field if non-nil, zero value otherwise.

### GetHasBeenPromotedOk

`func (o *DeliveryGroupResponseModel) GetHasBeenPromotedOk() (*bool, bool)`

GetHasBeenPromotedOk returns a tuple with the HasBeenPromoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromoted

`func (o *DeliveryGroupResponseModel) SetHasBeenPromoted(v bool)`

SetHasBeenPromoted sets HasBeenPromoted field to given value.


### GetHasBeenPromotedFrom

`func (o *DeliveryGroupResponseModel) GetHasBeenPromotedFrom() FunctionalLevel`

GetHasBeenPromotedFrom returns the HasBeenPromotedFrom field if non-nil, zero value otherwise.

### GetHasBeenPromotedFromOk

`func (o *DeliveryGroupResponseModel) GetHasBeenPromotedFromOk() (*FunctionalLevel, bool)`

GetHasBeenPromotedFromOk returns a tuple with the HasBeenPromotedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromotedFrom

`func (o *DeliveryGroupResponseModel) SetHasBeenPromotedFrom(v FunctionalLevel)`

SetHasBeenPromotedFrom sets HasBeenPromotedFrom field to given value.

### HasHasBeenPromotedFrom

`func (o *DeliveryGroupResponseModel) HasHasBeenPromotedFrom() bool`

HasHasBeenPromotedFrom returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *DeliveryGroupResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *DeliveryGroupResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *DeliveryGroupResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetIsBroken

`func (o *DeliveryGroupResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *DeliveryGroupResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *DeliveryGroupResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.


### GetIsRemotePC

`func (o *DeliveryGroupResponseModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *DeliveryGroupResponseModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *DeliveryGroupResponseModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.


### GetMachineLogOnType

`func (o *DeliveryGroupResponseModel) GetMachineLogOnType() MachineLogOnType`

GetMachineLogOnType returns the MachineLogOnType field if non-nil, zero value otherwise.

### GetMachineLogOnTypeOk

`func (o *DeliveryGroupResponseModel) GetMachineLogOnTypeOk() (*MachineLogOnType, bool)`

GetMachineLogOnTypeOk returns a tuple with the MachineLogOnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLogOnType

`func (o *DeliveryGroupResponseModel) SetMachineLogOnType(v MachineLogOnType)`

SetMachineLogOnType sets MachineLogOnType field to given value.

### HasMachineLogOnType

`func (o *DeliveryGroupResponseModel) HasMachineLogOnType() bool`

HasMachineLogOnType returns a boolean if a field has been set.

### GetMetadata

`func (o *DeliveryGroupResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeliveryGroupResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeliveryGroupResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeliveryGroupResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DeliveryGroupResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DeliveryGroupResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *DeliveryGroupResponseModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *DeliveryGroupResponseModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *DeliveryGroupResponseModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.


### GetName

`func (o *DeliveryGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *DeliveryGroupResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DeliveryGroupResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DeliveryGroupResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *DeliveryGroupResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *DeliveryGroupResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *DeliveryGroupResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPublishedName

`func (o *DeliveryGroupResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *DeliveryGroupResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *DeliveryGroupResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *DeliveryGroupResponseModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *DeliveryGroupResponseModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *DeliveryGroupResponseModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetPolicySetGuid

`func (o *DeliveryGroupResponseModel) GetPolicySetGuid() string`

GetPolicySetGuid returns the PolicySetGuid field if non-nil, zero value otherwise.

### GetPolicySetGuidOk

`func (o *DeliveryGroupResponseModel) GetPolicySetGuidOk() (*string, bool)`

GetPolicySetGuidOk returns a tuple with the PolicySetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetGuid

`func (o *DeliveryGroupResponseModel) SetPolicySetGuid(v string)`

SetPolicySetGuid sets PolicySetGuid field to given value.

### HasPolicySetGuid

`func (o *DeliveryGroupResponseModel) HasPolicySetGuid() bool`

HasPolicySetGuid returns a boolean if a field has been set.

### GetProductCode

`func (o *DeliveryGroupResponseModel) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *DeliveryGroupResponseModel) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *DeliveryGroupResponseModel) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *DeliveryGroupResponseModel) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCodeNil

`func (o *DeliveryGroupResponseModel) SetProductCodeNil(b bool)`

 SetProductCodeNil sets the value for ProductCode to be an explicit nil

### UnsetProductCode
`func (o *DeliveryGroupResponseModel) UnsetProductCode()`

UnsetProductCode ensures that no value is present for ProductCode, not even an explicit nil
### GetLicenseModel

`func (o *DeliveryGroupResponseModel) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *DeliveryGroupResponseModel) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *DeliveryGroupResponseModel) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *DeliveryGroupResponseModel) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetRequireUserHomeZone

`func (o *DeliveryGroupResponseModel) GetRequireUserHomeZone() bool`

GetRequireUserHomeZone returns the RequireUserHomeZone field if non-nil, zero value otherwise.

### GetRequireUserHomeZoneOk

`func (o *DeliveryGroupResponseModel) GetRequireUserHomeZoneOk() (*bool, bool)`

GetRequireUserHomeZoneOk returns a tuple with the RequireUserHomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireUserHomeZone

`func (o *DeliveryGroupResponseModel) SetRequireUserHomeZone(v bool)`

SetRequireUserHomeZone sets RequireUserHomeZone field to given value.


### GetScopes

`func (o *DeliveryGroupResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DeliveryGroupResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DeliveryGroupResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *DeliveryGroupResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *DeliveryGroupResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *DeliveryGroupResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *DeliveryGroupResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *DeliveryGroupResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *DeliveryGroupResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetSessionCount

`func (o *DeliveryGroupResponseModel) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *DeliveryGroupResponseModel) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *DeliveryGroupResponseModel) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.


### GetSessionSupport

`func (o *DeliveryGroupResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *DeliveryGroupResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *DeliveryGroupResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSharingKind

`func (o *DeliveryGroupResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *DeliveryGroupResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *DeliveryGroupResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.


### GetTotalApplications

`func (o *DeliveryGroupResponseModel) GetTotalApplications() int32`

GetTotalApplications returns the TotalApplications field if non-nil, zero value otherwise.

### GetTotalApplicationsOk

`func (o *DeliveryGroupResponseModel) GetTotalApplicationsOk() (*int32, bool)`

GetTotalApplicationsOk returns a tuple with the TotalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplications

`func (o *DeliveryGroupResponseModel) SetTotalApplications(v int32)`

SetTotalApplications sets TotalApplications field to given value.


### GetTotalDesktops

`func (o *DeliveryGroupResponseModel) GetTotalDesktops() int32`

GetTotalDesktops returns the TotalDesktops field if non-nil, zero value otherwise.

### GetTotalDesktopsOk

`func (o *DeliveryGroupResponseModel) GetTotalDesktopsOk() (*int32, bool)`

GetTotalDesktopsOk returns a tuple with the TotalDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDesktops

`func (o *DeliveryGroupResponseModel) SetTotalDesktops(v int32)`

SetTotalDesktops sets TotalDesktops field to given value.


### GetApplicationGroupCompatibility

`func (o *DeliveryGroupResponseModel) GetApplicationGroupCompatibility() AppGroupCompatibility`

GetApplicationGroupCompatibility returns the ApplicationGroupCompatibility field if non-nil, zero value otherwise.

### GetApplicationGroupCompatibilityOk

`func (o *DeliveryGroupResponseModel) GetApplicationGroupCompatibilityOk() (*AppGroupCompatibility, bool)`

GetApplicationGroupCompatibilityOk returns a tuple with the ApplicationGroupCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroupCompatibility

`func (o *DeliveryGroupResponseModel) SetApplicationGroupCompatibility(v AppGroupCompatibility)`

SetApplicationGroupCompatibility sets ApplicationGroupCompatibility field to given value.


### GetApplicationCompatibility

`func (o *DeliveryGroupResponseModel) GetApplicationCompatibility() AppOrDesktopCompatibility`

GetApplicationCompatibility returns the ApplicationCompatibility field if non-nil, zero value otherwise.

### GetApplicationCompatibilityOk

`func (o *DeliveryGroupResponseModel) GetApplicationCompatibilityOk() (*AppOrDesktopCompatibility, bool)`

GetApplicationCompatibilityOk returns a tuple with the ApplicationCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCompatibility

`func (o *DeliveryGroupResponseModel) SetApplicationCompatibility(v AppOrDesktopCompatibility)`

SetApplicationCompatibility sets ApplicationCompatibility field to given value.


### GetDesktopCompatibility

`func (o *DeliveryGroupResponseModel) GetDesktopCompatibility() AppOrDesktopCompatibility`

GetDesktopCompatibility returns the DesktopCompatibility field if non-nil, zero value otherwise.

### GetDesktopCompatibilityOk

`func (o *DeliveryGroupResponseModel) GetDesktopCompatibilityOk() (*AppOrDesktopCompatibility, bool)`

GetDesktopCompatibilityOk returns a tuple with the DesktopCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopCompatibility

`func (o *DeliveryGroupResponseModel) SetDesktopCompatibility(v AppOrDesktopCompatibility)`

SetDesktopCompatibility sets DesktopCompatibility field to given value.


### GetAdminFolder

`func (o *DeliveryGroupResponseModel) GetAdminFolder() RefResponseModel`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *DeliveryGroupResponseModel) GetAdminFolderOk() (*RefResponseModel, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *DeliveryGroupResponseModel) SetAdminFolder(v RefResponseModel)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *DeliveryGroupResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


