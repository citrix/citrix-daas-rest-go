# EditDeliveryGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminFolder** | Pointer to **NullableString** | The folder in which the delivery group resides. If not specified, the value will not be changed. May be specified as either the folder Id or Path. If specified as a path, and the path does not exist, it will be automatically created. | [optional] 
**AssignMachinesToUsers** | Pointer to [**[]AssignMachineToUserRequestModel**](AssignMachineToUserRequestModel.md) | Machine assignments to users. | [optional] 
**DeliveryType** | Pointer to [**DeliveryKind**](DeliveryKind.md) |  | [optional] 
**Description** | Pointer to **NullableString** | A description for this delivery group useful for administrators of the site. | [optional] 
**Desktops** | Pointer to [**[]DesktopRequestModel**](DesktopRequestModel.md) | List of desktop resources to publish on the delivery group. | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the delivery group should be in the enabled state; all resources published on disabled delivery groups do not appear to users. | [optional] 
**InMaintenanceMode** | Pointer to **NullableBool** | Whether the delivery group should be created in maintenance mode; a delivery group in maintenance mode will not allow users to connect or reconnect to machines in the delivery group. | [optional] 
**MinimumFunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**Name** | Pointer to **NullableString** | The name of the new delivery group. | [optional] 
**PublishedName** | Pointer to **NullableString** | The name of the desktop group as it is to appear to the user in StoreFront. | [optional] 
**RequireUserHomeZone** | Pointer to **NullableBool** | Whether to require the resources from this delivery group to launch within the user&#39;s home zone. | [optional] 
**Scopes** | Pointer to **[]string** | Administrative scopes which the delivery group should be a part of. If not specified, don&#39;t change. If specified, *all* desired scopes must be listed.  Any scope that the delivery group was part of previously, but which is not passed, will be removed. | [optional] 
**Tenants** | Pointer to **[]string** | Tenants to associate with the delivery group. | [optional] 
**AppProtectionKeyLoggingRequired** | Pointer to **NullableBool** | Specifies whether key logging app protection is required. | [optional] 
**AppProtectionScreenCaptureRequired** | Pointer to **NullableBool** | Specifies whether screen capture app protection is required. | [optional] 
**AppAccessPolicy** | Pointer to [**AppAccessPolicyRequestModel**](AppAccessPolicyRequestModel.md) |  | [optional] 
**AutomaticPowerOnForAssigned** | Pointer to **NullableBool** | Whether assigned (Private) machines in the delivery group should be automatically powered-on at the start of peak time periods. | [optional] 
**AutomaticPowerOnForAssignedDuringPeak** | Pointer to **NullableBool** | Whether assigned (Private) machines in the delivery group should be automatically powered-on throughout peak time periods. | [optional] 
**AutoScaleEnabled** | Pointer to **NullableBool** | Whether auto-scale is enabled for the delivery group. | [optional] 
**RestrictAutoscaleTag** | Pointer to **NullableString** | Specifies the tag used to restrict autoscale. | [optional] 
**ColorDepth** | Pointer to [**ColorDepth**](ColorDepth.md) |  | [optional] 
**ProductCode** | Pointer to **NullableString** | Specifies the product code of the delivery group. | [optional] 
**LicenseModel** | Pointer to [**LicenseModel**](LicenseModel.md) |  | [optional] 
**SecureIcaEnabled** | Pointer to **NullableBool** | Whether connections to machines in the delivery group will use SecureIca to encrypt the ICA protocol. | [optional] 
**DefaultDesktopIcon** | Pointer to **NullableString** | Default icon to use for desktops published from the delivery group. | [optional] 
**HdxSslEnabled** | Pointer to **NullableBool** | Whether connections to machines in the delivery group will use SSL. | [optional] 
**LingerSettings** | Pointer to [**FastApplicationSettingsRequestModel**](FastApplicationSettingsRequestModel.md) |  | [optional] 
**RestrictAutoscaleMinIdleUntaggedPercentDuringPeak** | Pointer to **NullableInt32** | The percentage that the number of untagged single-session machines in an idle state, or for multi-session machines, the untagged available load capacity must fall below before Autoscale powers on and manages &#39;tagged&#39; machines, as per policy, in peak time. If the number of untagged machines in an idle state, or the untagged available load capacity goes above this threshold value, Autoscale will attempt to shut down &#39;tagged&#39; machines. | [optional] 
**RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak** | Pointer to **NullableInt32** | The percentage that the number of untagged single-session machines in an idle state, or for multi-session machines, the untagged available load capacity must fall below before Autoscale powers on and manages &#39;tagged&#39; machines, as per policy, in off-peak. If the number of untagged machines in an idle state, or the untagged available load capacity goes above this threshold value, Autoscale will attempt to shut down &#39;tagged&#39; machines. | [optional] 
**OffPeakBufferSizePercent** | Pointer to **NullableInt32** | The percentage of machines in the delivery group that should be kept available in an idle state outside peak hours. | [optional] 
**OffPeakDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be performed after a user session disconnects outside peak hours. | [optional] 
**OffPeakExtendedDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakExtendedDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the second configured action should be performed after a user session disconnects outside peak hours. | [optional] 
**OffPeakLogOffAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakLogOffTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be performed after a user session ends outside peak hours. | [optional] 
**PeakBufferSizePercent** | Pointer to **NullableInt32** | The percentage of machines in the delivery group that should be kept available in an idle state in peak hours. | [optional] 
**PeakDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**PeakDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be performed after a user session disconnects in peak hours. | [optional] 
**PeakExtendedDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**PeakExtendedDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the second configured action should be performed after a user session disconnects in peak hours. | [optional] 
**LimitSecondsToForceLogOffUserDuringPeak** | Pointer to **NullableInt32** | Represents the number of seconds that must elapsed before the active sessions on the draining machines belonging to the delivery group are logged off, during peak time. | [optional] 
**LimitSecondsToForceLogOffUserDuringOffPeak** | Pointer to **NullableInt32** | represents the number of seconds that must elapsed before the active sessions on the draining machines belonging to the delivery group are logged off, during off-peak. | [optional] 
**LogOffWarningMessage** | Pointer to **NullableString** | The warning message to display to users in active sessions prior to logging off users, whether in peak time or off-peak. | [optional] 
**LogOffWarningTitle** | Pointer to **NullableString** | The title of the warning message dialog. | [optional] 
**PeakLogOffAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**PeakLogOffTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be performed after a user session ends in peak hours. | [optional] 
**DisconnectPeakIdleSessionAfterSeconds** | Pointer to **NullableInt32** | Specifies the time in seconds after which an idle session belonging to the delivery group is disconnected during peak time. | [optional] 
**DisconnectOffPeakIdleSessionAfterSeconds** | Pointer to **NullableInt32** | Specifies the time in seconds after which an idle session belonging to the delivery group is disconnected during off-peak time. | [optional] 
**LogoffPeakDisconnectedSessionAfterSeconds** | Pointer to **NullableInt32** | Specifies the time in seconds after which a disconnected session belonging to the delivery group is terminated during peak time. | [optional] 
**LogoffOffPeakDisconnectedSessionAfterSeconds** | Pointer to **NullableInt32** | Specifies the time in seconds after which a disconnected session belonging to the delivery group is terminated during off-peak time. | [optional] 
**PowerTimeSchemes** | Pointer to [**[]PowerTimeSchemeRequestModel**](PowerTimeSchemeRequestModel.md) | Power management time schemes.  No two schemes may cover the same day of the week. | [optional] 
**PowerOffDelayMinutes** | Pointer to **NullableInt32** | Delay before machines are powered off, when scaling down.  Specified in minutes.  Applies only to multi-session machines. | [optional] 
**AutoscaleLogOffReminderEnabled** | Pointer to **NullableBool** | Boolean value indicating whether the warning messages should be sent on an interval to nudge a logoff should be sent on an interval when autoscale is enabled. | [optional] 
**AutoscaleLogOffReminderIntervalSecondsOffPeak** | Pointer to **NullableInt32** | Represents the time interval at which messages are  sent to the user during off peak time when Autoscale is enabled. This message will nudge users to log off instead of forcibly logging them off. | [optional] 
**AutoscaleLogOffReminderIntervalSecondsPeak** | Pointer to **NullableInt32** | Represents the time interval at which messages are  sent to the user during peak time when autoscale is enabled. This message will nudge users to log off instead of forcibly logging them off. | [optional] 
**AutoscaleLogOffReminderMessage** | Pointer to **NullableString** | Notification message to display to users in active sessions belonging to machines needed by Autoscale for shutdown. | [optional] 
**AutoscaleLogOffReminderTitle** | Pointer to **NullableString** | Notification message dialog title displayed when Autoscale issues a logoff reminder request. | [optional] 
**MachineCost** | Pointer to **NullableFloat64** | Indicates the estimated per-hour cost for machines in the delivery group, as set by the administrator. | [optional] 
**MachineLogOnType** | Pointer to [**MachineLogOnType**](MachineLogOnType.md) |  | [optional] 
**PrelaunchSettings** | Pointer to [**FastApplicationSettingsRequestModel**](FastApplicationSettingsRequestModel.md) |  | [optional] 
**ProtocolPriority** | Pointer to [**[]ProtocolType**](ProtocolType.md) | A list of protocols in the order in which they should be attempted for use during connection. | [optional] 
**RebootSchedules** | Pointer to [**[]RebootScheduleRequestModel**](RebootScheduleRequestModel.md) | A list of reboot schedules. | [optional] 
**SettlementPeriodBeforeAutoShutdownSeconds** | Pointer to **NullableInt32** | Time after a session ends during which automatic shutdown requests (for example, shutdown after use, idle pool management) are deferred. Any outstanding shutdown request takes effect after the settlement period expires. This is typically used to configure time to allow logoff scripts to complete. | [optional] 
**SettlementPeriodBeforeUseSeconds** | Pointer to **NullableInt32** | Idle period before a machine can be selected to host a new session after registration or the end of a previous session. This is typically used to allow a machine to become idle following processing associated with start-up or logoff actions. A machine may still be selected during the idle period if no other machine is available for immediate use. | [optional] 
**ShutdownDesktopsAfterUse** | Pointer to **NullableBool** | Whether machines in this delivery group should be automatically shut down when each user session completes. | [optional] 
**SimpleAccessPolicy** | Pointer to [**SimplifiedAccessPolicyRequestModel**](SimplifiedAccessPolicyRequestModel.md) |  | [optional] 
**AdvancedAccessPolicy** | Pointer to [**[]AdvancedAccessPolicyRequestModel**](AdvancedAccessPolicyRequestModel.md) | Advanced access policy for connections to the delivery group. | [optional] 
**StoreFrontServersForHostedReceiver** | Pointer to [**[]StoreFrontServerRequestModel**](StoreFrontServerRequestModel.md) | List of StoreFront server addresses to configure within hosted receivers that are delivered from the delivery group. | [optional] 
**TimeZone** | Pointer to **NullableString** | The time zone in which this delivery group&#39;s machines reside. | [optional] 
**TurnOnAddedMachines** | Pointer to **NullableBool** | Whether to attempt to power on machines when they are added to the delivery group. | [optional] 
**ZonePreferences** | Pointer to [**[]ZonePreference**](ZonePreference.md) | Ordered list of zone preferences to be applied when launching resources from this delivery group. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of the delivery group. | [optional] 
**PolicySetGuid** | Pointer to **NullableString** | The GUID of the policy set assigned to this desktop group. Change if not null. Use Guid.Empty to clear the value stored in the database. A non-null and non-empty GUID assigns the policy set to this desktop group. | [optional] 

## Methods

### NewEditDeliveryGroupRequestModel

`func NewEditDeliveryGroupRequestModel() *EditDeliveryGroupRequestModel`

NewEditDeliveryGroupRequestModel instantiates a new EditDeliveryGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditDeliveryGroupRequestModelWithDefaults

`func NewEditDeliveryGroupRequestModelWithDefaults() *EditDeliveryGroupRequestModel`

NewEditDeliveryGroupRequestModelWithDefaults instantiates a new EditDeliveryGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminFolder

`func (o *EditDeliveryGroupRequestModel) GetAdminFolder() string`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *EditDeliveryGroupRequestModel) GetAdminFolderOk() (*string, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *EditDeliveryGroupRequestModel) SetAdminFolder(v string)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *EditDeliveryGroupRequestModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### SetAdminFolderNil

`func (o *EditDeliveryGroupRequestModel) SetAdminFolderNil(b bool)`

 SetAdminFolderNil sets the value for AdminFolder to be an explicit nil

### UnsetAdminFolder
`func (o *EditDeliveryGroupRequestModel) UnsetAdminFolder()`

UnsetAdminFolder ensures that no value is present for AdminFolder, not even an explicit nil
### GetAssignMachinesToUsers

`func (o *EditDeliveryGroupRequestModel) GetAssignMachinesToUsers() []AssignMachineToUserRequestModel`

GetAssignMachinesToUsers returns the AssignMachinesToUsers field if non-nil, zero value otherwise.

### GetAssignMachinesToUsersOk

`func (o *EditDeliveryGroupRequestModel) GetAssignMachinesToUsersOk() (*[]AssignMachineToUserRequestModel, bool)`

GetAssignMachinesToUsersOk returns a tuple with the AssignMachinesToUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignMachinesToUsers

`func (o *EditDeliveryGroupRequestModel) SetAssignMachinesToUsers(v []AssignMachineToUserRequestModel)`

SetAssignMachinesToUsers sets AssignMachinesToUsers field to given value.

### HasAssignMachinesToUsers

`func (o *EditDeliveryGroupRequestModel) HasAssignMachinesToUsers() bool`

HasAssignMachinesToUsers returns a boolean if a field has been set.

### SetAssignMachinesToUsersNil

`func (o *EditDeliveryGroupRequestModel) SetAssignMachinesToUsersNil(b bool)`

 SetAssignMachinesToUsersNil sets the value for AssignMachinesToUsers to be an explicit nil

### UnsetAssignMachinesToUsers
`func (o *EditDeliveryGroupRequestModel) UnsetAssignMachinesToUsers()`

UnsetAssignMachinesToUsers ensures that no value is present for AssignMachinesToUsers, not even an explicit nil
### GetDeliveryType

`func (o *EditDeliveryGroupRequestModel) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *EditDeliveryGroupRequestModel) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *EditDeliveryGroupRequestModel) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *EditDeliveryGroupRequestModel) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDescription

`func (o *EditDeliveryGroupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditDeliveryGroupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditDeliveryGroupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditDeliveryGroupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EditDeliveryGroupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditDeliveryGroupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesktops

`func (o *EditDeliveryGroupRequestModel) GetDesktops() []DesktopRequestModel`

GetDesktops returns the Desktops field if non-nil, zero value otherwise.

### GetDesktopsOk

`func (o *EditDeliveryGroupRequestModel) GetDesktopsOk() (*[]DesktopRequestModel, bool)`

GetDesktopsOk returns a tuple with the Desktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktops

`func (o *EditDeliveryGroupRequestModel) SetDesktops(v []DesktopRequestModel)`

SetDesktops sets Desktops field to given value.

### HasDesktops

`func (o *EditDeliveryGroupRequestModel) HasDesktops() bool`

HasDesktops returns a boolean if a field has been set.

### SetDesktopsNil

`func (o *EditDeliveryGroupRequestModel) SetDesktopsNil(b bool)`

 SetDesktopsNil sets the value for Desktops to be an explicit nil

### UnsetDesktops
`func (o *EditDeliveryGroupRequestModel) UnsetDesktops()`

UnsetDesktops ensures that no value is present for Desktops, not even an explicit nil
### GetEnabled

`func (o *EditDeliveryGroupRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EditDeliveryGroupRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EditDeliveryGroupRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EditDeliveryGroupRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *EditDeliveryGroupRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *EditDeliveryGroupRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetInMaintenanceMode

`func (o *EditDeliveryGroupRequestModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *EditDeliveryGroupRequestModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *EditDeliveryGroupRequestModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *EditDeliveryGroupRequestModel) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### SetInMaintenanceModeNil

`func (o *EditDeliveryGroupRequestModel) SetInMaintenanceModeNil(b bool)`

 SetInMaintenanceModeNil sets the value for InMaintenanceMode to be an explicit nil

### UnsetInMaintenanceMode
`func (o *EditDeliveryGroupRequestModel) UnsetInMaintenanceMode()`

UnsetInMaintenanceMode ensures that no value is present for InMaintenanceMode, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *EditDeliveryGroupRequestModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *EditDeliveryGroupRequestModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *EditDeliveryGroupRequestModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.

### HasMinimumFunctionalLevel

`func (o *EditDeliveryGroupRequestModel) HasMinimumFunctionalLevel() bool`

HasMinimumFunctionalLevel returns a boolean if a field has been set.

### GetName

`func (o *EditDeliveryGroupRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditDeliveryGroupRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditDeliveryGroupRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditDeliveryGroupRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditDeliveryGroupRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditDeliveryGroupRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPublishedName

`func (o *EditDeliveryGroupRequestModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *EditDeliveryGroupRequestModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *EditDeliveryGroupRequestModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *EditDeliveryGroupRequestModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *EditDeliveryGroupRequestModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *EditDeliveryGroupRequestModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetRequireUserHomeZone

`func (o *EditDeliveryGroupRequestModel) GetRequireUserHomeZone() bool`

GetRequireUserHomeZone returns the RequireUserHomeZone field if non-nil, zero value otherwise.

### GetRequireUserHomeZoneOk

`func (o *EditDeliveryGroupRequestModel) GetRequireUserHomeZoneOk() (*bool, bool)`

GetRequireUserHomeZoneOk returns a tuple with the RequireUserHomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireUserHomeZone

`func (o *EditDeliveryGroupRequestModel) SetRequireUserHomeZone(v bool)`

SetRequireUserHomeZone sets RequireUserHomeZone field to given value.

### HasRequireUserHomeZone

`func (o *EditDeliveryGroupRequestModel) HasRequireUserHomeZone() bool`

HasRequireUserHomeZone returns a boolean if a field has been set.

### SetRequireUserHomeZoneNil

`func (o *EditDeliveryGroupRequestModel) SetRequireUserHomeZoneNil(b bool)`

 SetRequireUserHomeZoneNil sets the value for RequireUserHomeZone to be an explicit nil

### UnsetRequireUserHomeZone
`func (o *EditDeliveryGroupRequestModel) UnsetRequireUserHomeZone()`

UnsetRequireUserHomeZone ensures that no value is present for RequireUserHomeZone, not even an explicit nil
### GetScopes

`func (o *EditDeliveryGroupRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *EditDeliveryGroupRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *EditDeliveryGroupRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *EditDeliveryGroupRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *EditDeliveryGroupRequestModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *EditDeliveryGroupRequestModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *EditDeliveryGroupRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *EditDeliveryGroupRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *EditDeliveryGroupRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *EditDeliveryGroupRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *EditDeliveryGroupRequestModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *EditDeliveryGroupRequestModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetAppProtectionKeyLoggingRequired

`func (o *EditDeliveryGroupRequestModel) GetAppProtectionKeyLoggingRequired() bool`

GetAppProtectionKeyLoggingRequired returns the AppProtectionKeyLoggingRequired field if non-nil, zero value otherwise.

### GetAppProtectionKeyLoggingRequiredOk

`func (o *EditDeliveryGroupRequestModel) GetAppProtectionKeyLoggingRequiredOk() (*bool, bool)`

GetAppProtectionKeyLoggingRequiredOk returns a tuple with the AppProtectionKeyLoggingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtectionKeyLoggingRequired

`func (o *EditDeliveryGroupRequestModel) SetAppProtectionKeyLoggingRequired(v bool)`

SetAppProtectionKeyLoggingRequired sets AppProtectionKeyLoggingRequired field to given value.

### HasAppProtectionKeyLoggingRequired

`func (o *EditDeliveryGroupRequestModel) HasAppProtectionKeyLoggingRequired() bool`

HasAppProtectionKeyLoggingRequired returns a boolean if a field has been set.

### SetAppProtectionKeyLoggingRequiredNil

`func (o *EditDeliveryGroupRequestModel) SetAppProtectionKeyLoggingRequiredNil(b bool)`

 SetAppProtectionKeyLoggingRequiredNil sets the value for AppProtectionKeyLoggingRequired to be an explicit nil

### UnsetAppProtectionKeyLoggingRequired
`func (o *EditDeliveryGroupRequestModel) UnsetAppProtectionKeyLoggingRequired()`

UnsetAppProtectionKeyLoggingRequired ensures that no value is present for AppProtectionKeyLoggingRequired, not even an explicit nil
### GetAppProtectionScreenCaptureRequired

`func (o *EditDeliveryGroupRequestModel) GetAppProtectionScreenCaptureRequired() bool`

GetAppProtectionScreenCaptureRequired returns the AppProtectionScreenCaptureRequired field if non-nil, zero value otherwise.

### GetAppProtectionScreenCaptureRequiredOk

`func (o *EditDeliveryGroupRequestModel) GetAppProtectionScreenCaptureRequiredOk() (*bool, bool)`

GetAppProtectionScreenCaptureRequiredOk returns a tuple with the AppProtectionScreenCaptureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtectionScreenCaptureRequired

`func (o *EditDeliveryGroupRequestModel) SetAppProtectionScreenCaptureRequired(v bool)`

SetAppProtectionScreenCaptureRequired sets AppProtectionScreenCaptureRequired field to given value.

### HasAppProtectionScreenCaptureRequired

`func (o *EditDeliveryGroupRequestModel) HasAppProtectionScreenCaptureRequired() bool`

HasAppProtectionScreenCaptureRequired returns a boolean if a field has been set.

### SetAppProtectionScreenCaptureRequiredNil

`func (o *EditDeliveryGroupRequestModel) SetAppProtectionScreenCaptureRequiredNil(b bool)`

 SetAppProtectionScreenCaptureRequiredNil sets the value for AppProtectionScreenCaptureRequired to be an explicit nil

### UnsetAppProtectionScreenCaptureRequired
`func (o *EditDeliveryGroupRequestModel) UnsetAppProtectionScreenCaptureRequired()`

UnsetAppProtectionScreenCaptureRequired ensures that no value is present for AppProtectionScreenCaptureRequired, not even an explicit nil
### GetAppAccessPolicy

`func (o *EditDeliveryGroupRequestModel) GetAppAccessPolicy() AppAccessPolicyRequestModel`

GetAppAccessPolicy returns the AppAccessPolicy field if non-nil, zero value otherwise.

### GetAppAccessPolicyOk

`func (o *EditDeliveryGroupRequestModel) GetAppAccessPolicyOk() (*AppAccessPolicyRequestModel, bool)`

GetAppAccessPolicyOk returns a tuple with the AppAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAccessPolicy

`func (o *EditDeliveryGroupRequestModel) SetAppAccessPolicy(v AppAccessPolicyRequestModel)`

SetAppAccessPolicy sets AppAccessPolicy field to given value.

### HasAppAccessPolicy

`func (o *EditDeliveryGroupRequestModel) HasAppAccessPolicy() bool`

HasAppAccessPolicy returns a boolean if a field has been set.

### GetAutomaticPowerOnForAssigned

`func (o *EditDeliveryGroupRequestModel) GetAutomaticPowerOnForAssigned() bool`

GetAutomaticPowerOnForAssigned returns the AutomaticPowerOnForAssigned field if non-nil, zero value otherwise.

### GetAutomaticPowerOnForAssignedOk

`func (o *EditDeliveryGroupRequestModel) GetAutomaticPowerOnForAssignedOk() (*bool, bool)`

GetAutomaticPowerOnForAssignedOk returns a tuple with the AutomaticPowerOnForAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPowerOnForAssigned

`func (o *EditDeliveryGroupRequestModel) SetAutomaticPowerOnForAssigned(v bool)`

SetAutomaticPowerOnForAssigned sets AutomaticPowerOnForAssigned field to given value.

### HasAutomaticPowerOnForAssigned

`func (o *EditDeliveryGroupRequestModel) HasAutomaticPowerOnForAssigned() bool`

HasAutomaticPowerOnForAssigned returns a boolean if a field has been set.

### SetAutomaticPowerOnForAssignedNil

`func (o *EditDeliveryGroupRequestModel) SetAutomaticPowerOnForAssignedNil(b bool)`

 SetAutomaticPowerOnForAssignedNil sets the value for AutomaticPowerOnForAssigned to be an explicit nil

### UnsetAutomaticPowerOnForAssigned
`func (o *EditDeliveryGroupRequestModel) UnsetAutomaticPowerOnForAssigned()`

UnsetAutomaticPowerOnForAssigned ensures that no value is present for AutomaticPowerOnForAssigned, not even an explicit nil
### GetAutomaticPowerOnForAssignedDuringPeak

`func (o *EditDeliveryGroupRequestModel) GetAutomaticPowerOnForAssignedDuringPeak() bool`

GetAutomaticPowerOnForAssignedDuringPeak returns the AutomaticPowerOnForAssignedDuringPeak field if non-nil, zero value otherwise.

### GetAutomaticPowerOnForAssignedDuringPeakOk

`func (o *EditDeliveryGroupRequestModel) GetAutomaticPowerOnForAssignedDuringPeakOk() (*bool, bool)`

GetAutomaticPowerOnForAssignedDuringPeakOk returns a tuple with the AutomaticPowerOnForAssignedDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPowerOnForAssignedDuringPeak

`func (o *EditDeliveryGroupRequestModel) SetAutomaticPowerOnForAssignedDuringPeak(v bool)`

SetAutomaticPowerOnForAssignedDuringPeak sets AutomaticPowerOnForAssignedDuringPeak field to given value.

### HasAutomaticPowerOnForAssignedDuringPeak

`func (o *EditDeliveryGroupRequestModel) HasAutomaticPowerOnForAssignedDuringPeak() bool`

HasAutomaticPowerOnForAssignedDuringPeak returns a boolean if a field has been set.

### SetAutomaticPowerOnForAssignedDuringPeakNil

`func (o *EditDeliveryGroupRequestModel) SetAutomaticPowerOnForAssignedDuringPeakNil(b bool)`

 SetAutomaticPowerOnForAssignedDuringPeakNil sets the value for AutomaticPowerOnForAssignedDuringPeak to be an explicit nil

### UnsetAutomaticPowerOnForAssignedDuringPeak
`func (o *EditDeliveryGroupRequestModel) UnsetAutomaticPowerOnForAssignedDuringPeak()`

UnsetAutomaticPowerOnForAssignedDuringPeak ensures that no value is present for AutomaticPowerOnForAssignedDuringPeak, not even an explicit nil
### GetAutoScaleEnabled

`func (o *EditDeliveryGroupRequestModel) GetAutoScaleEnabled() bool`

GetAutoScaleEnabled returns the AutoScaleEnabled field if non-nil, zero value otherwise.

### GetAutoScaleEnabledOk

`func (o *EditDeliveryGroupRequestModel) GetAutoScaleEnabledOk() (*bool, bool)`

GetAutoScaleEnabledOk returns a tuple with the AutoScaleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaleEnabled

`func (o *EditDeliveryGroupRequestModel) SetAutoScaleEnabled(v bool)`

SetAutoScaleEnabled sets AutoScaleEnabled field to given value.

### HasAutoScaleEnabled

`func (o *EditDeliveryGroupRequestModel) HasAutoScaleEnabled() bool`

HasAutoScaleEnabled returns a boolean if a field has been set.

### SetAutoScaleEnabledNil

`func (o *EditDeliveryGroupRequestModel) SetAutoScaleEnabledNil(b bool)`

 SetAutoScaleEnabledNil sets the value for AutoScaleEnabled to be an explicit nil

### UnsetAutoScaleEnabled
`func (o *EditDeliveryGroupRequestModel) UnsetAutoScaleEnabled()`

UnsetAutoScaleEnabled ensures that no value is present for AutoScaleEnabled, not even an explicit nil
### GetRestrictAutoscaleTag

`func (o *EditDeliveryGroupRequestModel) GetRestrictAutoscaleTag() string`

GetRestrictAutoscaleTag returns the RestrictAutoscaleTag field if non-nil, zero value otherwise.

### GetRestrictAutoscaleTagOk

`func (o *EditDeliveryGroupRequestModel) GetRestrictAutoscaleTagOk() (*string, bool)`

GetRestrictAutoscaleTagOk returns a tuple with the RestrictAutoscaleTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleTag

`func (o *EditDeliveryGroupRequestModel) SetRestrictAutoscaleTag(v string)`

SetRestrictAutoscaleTag sets RestrictAutoscaleTag field to given value.

### HasRestrictAutoscaleTag

`func (o *EditDeliveryGroupRequestModel) HasRestrictAutoscaleTag() bool`

HasRestrictAutoscaleTag returns a boolean if a field has been set.

### SetRestrictAutoscaleTagNil

`func (o *EditDeliveryGroupRequestModel) SetRestrictAutoscaleTagNil(b bool)`

 SetRestrictAutoscaleTagNil sets the value for RestrictAutoscaleTag to be an explicit nil

### UnsetRestrictAutoscaleTag
`func (o *EditDeliveryGroupRequestModel) UnsetRestrictAutoscaleTag()`

UnsetRestrictAutoscaleTag ensures that no value is present for RestrictAutoscaleTag, not even an explicit nil
### GetColorDepth

`func (o *EditDeliveryGroupRequestModel) GetColorDepth() ColorDepth`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *EditDeliveryGroupRequestModel) GetColorDepthOk() (*ColorDepth, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *EditDeliveryGroupRequestModel) SetColorDepth(v ColorDepth)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *EditDeliveryGroupRequestModel) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetProductCode

`func (o *EditDeliveryGroupRequestModel) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *EditDeliveryGroupRequestModel) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *EditDeliveryGroupRequestModel) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *EditDeliveryGroupRequestModel) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCodeNil

`func (o *EditDeliveryGroupRequestModel) SetProductCodeNil(b bool)`

 SetProductCodeNil sets the value for ProductCode to be an explicit nil

### UnsetProductCode
`func (o *EditDeliveryGroupRequestModel) UnsetProductCode()`

UnsetProductCode ensures that no value is present for ProductCode, not even an explicit nil
### GetLicenseModel

`func (o *EditDeliveryGroupRequestModel) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *EditDeliveryGroupRequestModel) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *EditDeliveryGroupRequestModel) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *EditDeliveryGroupRequestModel) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetSecureIcaEnabled

`func (o *EditDeliveryGroupRequestModel) GetSecureIcaEnabled() bool`

GetSecureIcaEnabled returns the SecureIcaEnabled field if non-nil, zero value otherwise.

### GetSecureIcaEnabledOk

`func (o *EditDeliveryGroupRequestModel) GetSecureIcaEnabledOk() (*bool, bool)`

GetSecureIcaEnabledOk returns a tuple with the SecureIcaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureIcaEnabled

`func (o *EditDeliveryGroupRequestModel) SetSecureIcaEnabled(v bool)`

SetSecureIcaEnabled sets SecureIcaEnabled field to given value.

### HasSecureIcaEnabled

`func (o *EditDeliveryGroupRequestModel) HasSecureIcaEnabled() bool`

HasSecureIcaEnabled returns a boolean if a field has been set.

### SetSecureIcaEnabledNil

`func (o *EditDeliveryGroupRequestModel) SetSecureIcaEnabledNil(b bool)`

 SetSecureIcaEnabledNil sets the value for SecureIcaEnabled to be an explicit nil

### UnsetSecureIcaEnabled
`func (o *EditDeliveryGroupRequestModel) UnsetSecureIcaEnabled()`

UnsetSecureIcaEnabled ensures that no value is present for SecureIcaEnabled, not even an explicit nil
### GetDefaultDesktopIcon

`func (o *EditDeliveryGroupRequestModel) GetDefaultDesktopIcon() string`

GetDefaultDesktopIcon returns the DefaultDesktopIcon field if non-nil, zero value otherwise.

### GetDefaultDesktopIconOk

`func (o *EditDeliveryGroupRequestModel) GetDefaultDesktopIconOk() (*string, bool)`

GetDefaultDesktopIconOk returns a tuple with the DefaultDesktopIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDesktopIcon

`func (o *EditDeliveryGroupRequestModel) SetDefaultDesktopIcon(v string)`

SetDefaultDesktopIcon sets DefaultDesktopIcon field to given value.

### HasDefaultDesktopIcon

`func (o *EditDeliveryGroupRequestModel) HasDefaultDesktopIcon() bool`

HasDefaultDesktopIcon returns a boolean if a field has been set.

### SetDefaultDesktopIconNil

`func (o *EditDeliveryGroupRequestModel) SetDefaultDesktopIconNil(b bool)`

 SetDefaultDesktopIconNil sets the value for DefaultDesktopIcon to be an explicit nil

### UnsetDefaultDesktopIcon
`func (o *EditDeliveryGroupRequestModel) UnsetDefaultDesktopIcon()`

UnsetDefaultDesktopIcon ensures that no value is present for DefaultDesktopIcon, not even an explicit nil
### GetHdxSslEnabled

`func (o *EditDeliveryGroupRequestModel) GetHdxSslEnabled() bool`

GetHdxSslEnabled returns the HdxSslEnabled field if non-nil, zero value otherwise.

### GetHdxSslEnabledOk

`func (o *EditDeliveryGroupRequestModel) GetHdxSslEnabledOk() (*bool, bool)`

GetHdxSslEnabledOk returns a tuple with the HdxSslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdxSslEnabled

`func (o *EditDeliveryGroupRequestModel) SetHdxSslEnabled(v bool)`

SetHdxSslEnabled sets HdxSslEnabled field to given value.

### HasHdxSslEnabled

`func (o *EditDeliveryGroupRequestModel) HasHdxSslEnabled() bool`

HasHdxSslEnabled returns a boolean if a field has been set.

### SetHdxSslEnabledNil

`func (o *EditDeliveryGroupRequestModel) SetHdxSslEnabledNil(b bool)`

 SetHdxSslEnabledNil sets the value for HdxSslEnabled to be an explicit nil

### UnsetHdxSslEnabled
`func (o *EditDeliveryGroupRequestModel) UnsetHdxSslEnabled()`

UnsetHdxSslEnabled ensures that no value is present for HdxSslEnabled, not even an explicit nil
### GetLingerSettings

`func (o *EditDeliveryGroupRequestModel) GetLingerSettings() FastApplicationSettingsRequestModel`

GetLingerSettings returns the LingerSettings field if non-nil, zero value otherwise.

### GetLingerSettingsOk

`func (o *EditDeliveryGroupRequestModel) GetLingerSettingsOk() (*FastApplicationSettingsRequestModel, bool)`

GetLingerSettingsOk returns a tuple with the LingerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLingerSettings

`func (o *EditDeliveryGroupRequestModel) SetLingerSettings(v FastApplicationSettingsRequestModel)`

SetLingerSettings sets LingerSettings field to given value.

### HasLingerSettings

`func (o *EditDeliveryGroupRequestModel) HasLingerSettings() bool`

HasLingerSettings returns a boolean if a field has been set.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *EditDeliveryGroupRequestModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak() int32`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak returns the RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field if non-nil, zero value otherwise.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk

`func (o *EditDeliveryGroupRequestModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk() (*int32, bool)`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk returns a tuple with the RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *EditDeliveryGroupRequestModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak(v int32)`

SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak sets RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field to given value.

### HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *EditDeliveryGroupRequestModel) HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak() bool`

HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak returns a boolean if a field has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakNil

`func (o *EditDeliveryGroupRequestModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakNil(b bool)`

 SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakNil sets the value for RestrictAutoscaleMinIdleUntaggedPercentDuringPeak to be an explicit nil

### UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak
`func (o *EditDeliveryGroupRequestModel) UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak()`

UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak ensures that no value is present for RestrictAutoscaleMinIdleUntaggedPercentDuringPeak, not even an explicit nil
### GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *EditDeliveryGroupRequestModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak() int32`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak returns the RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field if non-nil, zero value otherwise.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk

`func (o *EditDeliveryGroupRequestModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk() (*int32, bool)`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk returns a tuple with the RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *EditDeliveryGroupRequestModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak(v int32)`

SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak sets RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field to given value.

### HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *EditDeliveryGroupRequestModel) HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak() bool`

HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak returns a boolean if a field has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakNil

`func (o *EditDeliveryGroupRequestModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakNil(b bool)`

 SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakNil sets the value for RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak to be an explicit nil

### UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak
`func (o *EditDeliveryGroupRequestModel) UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak()`

UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak ensures that no value is present for RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak, not even an explicit nil
### GetOffPeakBufferSizePercent

`func (o *EditDeliveryGroupRequestModel) GetOffPeakBufferSizePercent() int32`

GetOffPeakBufferSizePercent returns the OffPeakBufferSizePercent field if non-nil, zero value otherwise.

### GetOffPeakBufferSizePercentOk

`func (o *EditDeliveryGroupRequestModel) GetOffPeakBufferSizePercentOk() (*int32, bool)`

GetOffPeakBufferSizePercentOk returns a tuple with the OffPeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakBufferSizePercent

`func (o *EditDeliveryGroupRequestModel) SetOffPeakBufferSizePercent(v int32)`

SetOffPeakBufferSizePercent sets OffPeakBufferSizePercent field to given value.

### HasOffPeakBufferSizePercent

`func (o *EditDeliveryGroupRequestModel) HasOffPeakBufferSizePercent() bool`

HasOffPeakBufferSizePercent returns a boolean if a field has been set.

### SetOffPeakBufferSizePercentNil

`func (o *EditDeliveryGroupRequestModel) SetOffPeakBufferSizePercentNil(b bool)`

 SetOffPeakBufferSizePercentNil sets the value for OffPeakBufferSizePercent to be an explicit nil

### UnsetOffPeakBufferSizePercent
`func (o *EditDeliveryGroupRequestModel) UnsetOffPeakBufferSizePercent()`

UnsetOffPeakBufferSizePercent ensures that no value is present for OffPeakBufferSizePercent, not even an explicit nil
### GetOffPeakDisconnectAction

`func (o *EditDeliveryGroupRequestModel) GetOffPeakDisconnectAction() SessionChangeHostingAction`

GetOffPeakDisconnectAction returns the OffPeakDisconnectAction field if non-nil, zero value otherwise.

### GetOffPeakDisconnectActionOk

`func (o *EditDeliveryGroupRequestModel) GetOffPeakDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakDisconnectActionOk returns a tuple with the OffPeakDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectAction

`func (o *EditDeliveryGroupRequestModel) SetOffPeakDisconnectAction(v SessionChangeHostingAction)`

SetOffPeakDisconnectAction sets OffPeakDisconnectAction field to given value.

### HasOffPeakDisconnectAction

`func (o *EditDeliveryGroupRequestModel) HasOffPeakDisconnectAction() bool`

HasOffPeakDisconnectAction returns a boolean if a field has been set.

### GetOffPeakDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) GetOffPeakDisconnectTimeoutMinutes() int32`

GetOffPeakDisconnectTimeoutMinutes returns the OffPeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakDisconnectTimeoutMinutesOk

`func (o *EditDeliveryGroupRequestModel) GetOffPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakDisconnectTimeoutMinutesOk returns a tuple with the OffPeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) SetOffPeakDisconnectTimeoutMinutes(v int32)`

SetOffPeakDisconnectTimeoutMinutes sets OffPeakDisconnectTimeoutMinutes field to given value.

### HasOffPeakDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) HasOffPeakDisconnectTimeoutMinutes() bool`

HasOffPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakDisconnectTimeoutMinutesNil

`func (o *EditDeliveryGroupRequestModel) SetOffPeakDisconnectTimeoutMinutesNil(b bool)`

 SetOffPeakDisconnectTimeoutMinutesNil sets the value for OffPeakDisconnectTimeoutMinutes to be an explicit nil

### UnsetOffPeakDisconnectTimeoutMinutes
`func (o *EditDeliveryGroupRequestModel) UnsetOffPeakDisconnectTimeoutMinutes()`

UnsetOffPeakDisconnectTimeoutMinutes ensures that no value is present for OffPeakDisconnectTimeoutMinutes, not even an explicit nil
### GetOffPeakExtendedDisconnectAction

`func (o *EditDeliveryGroupRequestModel) GetOffPeakExtendedDisconnectAction() SessionChangeHostingAction`

GetOffPeakExtendedDisconnectAction returns the OffPeakExtendedDisconnectAction field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectActionOk

`func (o *EditDeliveryGroupRequestModel) GetOffPeakExtendedDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakExtendedDisconnectActionOk returns a tuple with the OffPeakExtendedDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectAction

`func (o *EditDeliveryGroupRequestModel) SetOffPeakExtendedDisconnectAction(v SessionChangeHostingAction)`

SetOffPeakExtendedDisconnectAction sets OffPeakExtendedDisconnectAction field to given value.

### HasOffPeakExtendedDisconnectAction

`func (o *EditDeliveryGroupRequestModel) HasOffPeakExtendedDisconnectAction() bool`

HasOffPeakExtendedDisconnectAction returns a boolean if a field has been set.

### GetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) GetOffPeakExtendedDisconnectTimeoutMinutes() int32`

GetOffPeakExtendedDisconnectTimeoutMinutes returns the OffPeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectTimeoutMinutesOk

`func (o *EditDeliveryGroupRequestModel) GetOffPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the OffPeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) SetOffPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetOffPeakExtendedDisconnectTimeoutMinutes sets OffPeakExtendedDisconnectTimeoutMinutes field to given value.

### HasOffPeakExtendedDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) HasOffPeakExtendedDisconnectTimeoutMinutes() bool`

HasOffPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakExtendedDisconnectTimeoutMinutesNil

`func (o *EditDeliveryGroupRequestModel) SetOffPeakExtendedDisconnectTimeoutMinutesNil(b bool)`

 SetOffPeakExtendedDisconnectTimeoutMinutesNil sets the value for OffPeakExtendedDisconnectTimeoutMinutes to be an explicit nil

### UnsetOffPeakExtendedDisconnectTimeoutMinutes
`func (o *EditDeliveryGroupRequestModel) UnsetOffPeakExtendedDisconnectTimeoutMinutes()`

UnsetOffPeakExtendedDisconnectTimeoutMinutes ensures that no value is present for OffPeakExtendedDisconnectTimeoutMinutes, not even an explicit nil
### GetOffPeakLogOffAction

`func (o *EditDeliveryGroupRequestModel) GetOffPeakLogOffAction() SessionChangeHostingAction`

GetOffPeakLogOffAction returns the OffPeakLogOffAction field if non-nil, zero value otherwise.

### GetOffPeakLogOffActionOk

`func (o *EditDeliveryGroupRequestModel) GetOffPeakLogOffActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakLogOffActionOk returns a tuple with the OffPeakLogOffAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffAction

`func (o *EditDeliveryGroupRequestModel) SetOffPeakLogOffAction(v SessionChangeHostingAction)`

SetOffPeakLogOffAction sets OffPeakLogOffAction field to given value.

### HasOffPeakLogOffAction

`func (o *EditDeliveryGroupRequestModel) HasOffPeakLogOffAction() bool`

HasOffPeakLogOffAction returns a boolean if a field has been set.

### GetOffPeakLogOffTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) GetOffPeakLogOffTimeoutMinutes() int32`

GetOffPeakLogOffTimeoutMinutes returns the OffPeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakLogOffTimeoutMinutesOk

`func (o *EditDeliveryGroupRequestModel) GetOffPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetOffPeakLogOffTimeoutMinutesOk returns a tuple with the OffPeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) SetOffPeakLogOffTimeoutMinutes(v int32)`

SetOffPeakLogOffTimeoutMinutes sets OffPeakLogOffTimeoutMinutes field to given value.

### HasOffPeakLogOffTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) HasOffPeakLogOffTimeoutMinutes() bool`

HasOffPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakLogOffTimeoutMinutesNil

`func (o *EditDeliveryGroupRequestModel) SetOffPeakLogOffTimeoutMinutesNil(b bool)`

 SetOffPeakLogOffTimeoutMinutesNil sets the value for OffPeakLogOffTimeoutMinutes to be an explicit nil

### UnsetOffPeakLogOffTimeoutMinutes
`func (o *EditDeliveryGroupRequestModel) UnsetOffPeakLogOffTimeoutMinutes()`

UnsetOffPeakLogOffTimeoutMinutes ensures that no value is present for OffPeakLogOffTimeoutMinutes, not even an explicit nil
### GetPeakBufferSizePercent

`func (o *EditDeliveryGroupRequestModel) GetPeakBufferSizePercent() int32`

GetPeakBufferSizePercent returns the PeakBufferSizePercent field if non-nil, zero value otherwise.

### GetPeakBufferSizePercentOk

`func (o *EditDeliveryGroupRequestModel) GetPeakBufferSizePercentOk() (*int32, bool)`

GetPeakBufferSizePercentOk returns a tuple with the PeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBufferSizePercent

`func (o *EditDeliveryGroupRequestModel) SetPeakBufferSizePercent(v int32)`

SetPeakBufferSizePercent sets PeakBufferSizePercent field to given value.

### HasPeakBufferSizePercent

`func (o *EditDeliveryGroupRequestModel) HasPeakBufferSizePercent() bool`

HasPeakBufferSizePercent returns a boolean if a field has been set.

### SetPeakBufferSizePercentNil

`func (o *EditDeliveryGroupRequestModel) SetPeakBufferSizePercentNil(b bool)`

 SetPeakBufferSizePercentNil sets the value for PeakBufferSizePercent to be an explicit nil

### UnsetPeakBufferSizePercent
`func (o *EditDeliveryGroupRequestModel) UnsetPeakBufferSizePercent()`

UnsetPeakBufferSizePercent ensures that no value is present for PeakBufferSizePercent, not even an explicit nil
### GetPeakDisconnectAction

`func (o *EditDeliveryGroupRequestModel) GetPeakDisconnectAction() SessionChangeHostingAction`

GetPeakDisconnectAction returns the PeakDisconnectAction field if non-nil, zero value otherwise.

### GetPeakDisconnectActionOk

`func (o *EditDeliveryGroupRequestModel) GetPeakDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetPeakDisconnectActionOk returns a tuple with the PeakDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectAction

`func (o *EditDeliveryGroupRequestModel) SetPeakDisconnectAction(v SessionChangeHostingAction)`

SetPeakDisconnectAction sets PeakDisconnectAction field to given value.

### HasPeakDisconnectAction

`func (o *EditDeliveryGroupRequestModel) HasPeakDisconnectAction() bool`

HasPeakDisconnectAction returns a boolean if a field has been set.

### GetPeakDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) GetPeakDisconnectTimeoutMinutes() int32`

GetPeakDisconnectTimeoutMinutes returns the PeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakDisconnectTimeoutMinutesOk

`func (o *EditDeliveryGroupRequestModel) GetPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakDisconnectTimeoutMinutesOk returns a tuple with the PeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) SetPeakDisconnectTimeoutMinutes(v int32)`

SetPeakDisconnectTimeoutMinutes sets PeakDisconnectTimeoutMinutes field to given value.

### HasPeakDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) HasPeakDisconnectTimeoutMinutes() bool`

HasPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetPeakDisconnectTimeoutMinutesNil

`func (o *EditDeliveryGroupRequestModel) SetPeakDisconnectTimeoutMinutesNil(b bool)`

 SetPeakDisconnectTimeoutMinutesNil sets the value for PeakDisconnectTimeoutMinutes to be an explicit nil

### UnsetPeakDisconnectTimeoutMinutes
`func (o *EditDeliveryGroupRequestModel) UnsetPeakDisconnectTimeoutMinutes()`

UnsetPeakDisconnectTimeoutMinutes ensures that no value is present for PeakDisconnectTimeoutMinutes, not even an explicit nil
### GetPeakExtendedDisconnectAction

`func (o *EditDeliveryGroupRequestModel) GetPeakExtendedDisconnectAction() SessionChangeHostingAction`

GetPeakExtendedDisconnectAction returns the PeakExtendedDisconnectAction field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectActionOk

`func (o *EditDeliveryGroupRequestModel) GetPeakExtendedDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetPeakExtendedDisconnectActionOk returns a tuple with the PeakExtendedDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectAction

`func (o *EditDeliveryGroupRequestModel) SetPeakExtendedDisconnectAction(v SessionChangeHostingAction)`

SetPeakExtendedDisconnectAction sets PeakExtendedDisconnectAction field to given value.

### HasPeakExtendedDisconnectAction

`func (o *EditDeliveryGroupRequestModel) HasPeakExtendedDisconnectAction() bool`

HasPeakExtendedDisconnectAction returns a boolean if a field has been set.

### GetPeakExtendedDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) GetPeakExtendedDisconnectTimeoutMinutes() int32`

GetPeakExtendedDisconnectTimeoutMinutes returns the PeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectTimeoutMinutesOk

`func (o *EditDeliveryGroupRequestModel) GetPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the PeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) SetPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetPeakExtendedDisconnectTimeoutMinutes sets PeakExtendedDisconnectTimeoutMinutes field to given value.

### HasPeakExtendedDisconnectTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) HasPeakExtendedDisconnectTimeoutMinutes() bool`

HasPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetPeakExtendedDisconnectTimeoutMinutesNil

`func (o *EditDeliveryGroupRequestModel) SetPeakExtendedDisconnectTimeoutMinutesNil(b bool)`

 SetPeakExtendedDisconnectTimeoutMinutesNil sets the value for PeakExtendedDisconnectTimeoutMinutes to be an explicit nil

### UnsetPeakExtendedDisconnectTimeoutMinutes
`func (o *EditDeliveryGroupRequestModel) UnsetPeakExtendedDisconnectTimeoutMinutes()`

UnsetPeakExtendedDisconnectTimeoutMinutes ensures that no value is present for PeakExtendedDisconnectTimeoutMinutes, not even an explicit nil
### GetLimitSecondsToForceLogOffUserDuringPeak

`func (o *EditDeliveryGroupRequestModel) GetLimitSecondsToForceLogOffUserDuringPeak() int32`

GetLimitSecondsToForceLogOffUserDuringPeak returns the LimitSecondsToForceLogOffUserDuringPeak field if non-nil, zero value otherwise.

### GetLimitSecondsToForceLogOffUserDuringPeakOk

`func (o *EditDeliveryGroupRequestModel) GetLimitSecondsToForceLogOffUserDuringPeakOk() (*int32, bool)`

GetLimitSecondsToForceLogOffUserDuringPeakOk returns a tuple with the LimitSecondsToForceLogOffUserDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSecondsToForceLogOffUserDuringPeak

`func (o *EditDeliveryGroupRequestModel) SetLimitSecondsToForceLogOffUserDuringPeak(v int32)`

SetLimitSecondsToForceLogOffUserDuringPeak sets LimitSecondsToForceLogOffUserDuringPeak field to given value.

### HasLimitSecondsToForceLogOffUserDuringPeak

`func (o *EditDeliveryGroupRequestModel) HasLimitSecondsToForceLogOffUserDuringPeak() bool`

HasLimitSecondsToForceLogOffUserDuringPeak returns a boolean if a field has been set.

### SetLimitSecondsToForceLogOffUserDuringPeakNil

`func (o *EditDeliveryGroupRequestModel) SetLimitSecondsToForceLogOffUserDuringPeakNil(b bool)`

 SetLimitSecondsToForceLogOffUserDuringPeakNil sets the value for LimitSecondsToForceLogOffUserDuringPeak to be an explicit nil

### UnsetLimitSecondsToForceLogOffUserDuringPeak
`func (o *EditDeliveryGroupRequestModel) UnsetLimitSecondsToForceLogOffUserDuringPeak()`

UnsetLimitSecondsToForceLogOffUserDuringPeak ensures that no value is present for LimitSecondsToForceLogOffUserDuringPeak, not even an explicit nil
### GetLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *EditDeliveryGroupRequestModel) GetLimitSecondsToForceLogOffUserDuringOffPeak() int32`

GetLimitSecondsToForceLogOffUserDuringOffPeak returns the LimitSecondsToForceLogOffUserDuringOffPeak field if non-nil, zero value otherwise.

### GetLimitSecondsToForceLogOffUserDuringOffPeakOk

`func (o *EditDeliveryGroupRequestModel) GetLimitSecondsToForceLogOffUserDuringOffPeakOk() (*int32, bool)`

GetLimitSecondsToForceLogOffUserDuringOffPeakOk returns a tuple with the LimitSecondsToForceLogOffUserDuringOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *EditDeliveryGroupRequestModel) SetLimitSecondsToForceLogOffUserDuringOffPeak(v int32)`

SetLimitSecondsToForceLogOffUserDuringOffPeak sets LimitSecondsToForceLogOffUserDuringOffPeak field to given value.

### HasLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *EditDeliveryGroupRequestModel) HasLimitSecondsToForceLogOffUserDuringOffPeak() bool`

HasLimitSecondsToForceLogOffUserDuringOffPeak returns a boolean if a field has been set.

### SetLimitSecondsToForceLogOffUserDuringOffPeakNil

`func (o *EditDeliveryGroupRequestModel) SetLimitSecondsToForceLogOffUserDuringOffPeakNil(b bool)`

 SetLimitSecondsToForceLogOffUserDuringOffPeakNil sets the value for LimitSecondsToForceLogOffUserDuringOffPeak to be an explicit nil

### UnsetLimitSecondsToForceLogOffUserDuringOffPeak
`func (o *EditDeliveryGroupRequestModel) UnsetLimitSecondsToForceLogOffUserDuringOffPeak()`

UnsetLimitSecondsToForceLogOffUserDuringOffPeak ensures that no value is present for LimitSecondsToForceLogOffUserDuringOffPeak, not even an explicit nil
### GetLogOffWarningMessage

`func (o *EditDeliveryGroupRequestModel) GetLogOffWarningMessage() string`

GetLogOffWarningMessage returns the LogOffWarningMessage field if non-nil, zero value otherwise.

### GetLogOffWarningMessageOk

`func (o *EditDeliveryGroupRequestModel) GetLogOffWarningMessageOk() (*string, bool)`

GetLogOffWarningMessageOk returns a tuple with the LogOffWarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOffWarningMessage

`func (o *EditDeliveryGroupRequestModel) SetLogOffWarningMessage(v string)`

SetLogOffWarningMessage sets LogOffWarningMessage field to given value.

### HasLogOffWarningMessage

`func (o *EditDeliveryGroupRequestModel) HasLogOffWarningMessage() bool`

HasLogOffWarningMessage returns a boolean if a field has been set.

### SetLogOffWarningMessageNil

`func (o *EditDeliveryGroupRequestModel) SetLogOffWarningMessageNil(b bool)`

 SetLogOffWarningMessageNil sets the value for LogOffWarningMessage to be an explicit nil

### UnsetLogOffWarningMessage
`func (o *EditDeliveryGroupRequestModel) UnsetLogOffWarningMessage()`

UnsetLogOffWarningMessage ensures that no value is present for LogOffWarningMessage, not even an explicit nil
### GetLogOffWarningTitle

`func (o *EditDeliveryGroupRequestModel) GetLogOffWarningTitle() string`

GetLogOffWarningTitle returns the LogOffWarningTitle field if non-nil, zero value otherwise.

### GetLogOffWarningTitleOk

`func (o *EditDeliveryGroupRequestModel) GetLogOffWarningTitleOk() (*string, bool)`

GetLogOffWarningTitleOk returns a tuple with the LogOffWarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOffWarningTitle

`func (o *EditDeliveryGroupRequestModel) SetLogOffWarningTitle(v string)`

SetLogOffWarningTitle sets LogOffWarningTitle field to given value.

### HasLogOffWarningTitle

`func (o *EditDeliveryGroupRequestModel) HasLogOffWarningTitle() bool`

HasLogOffWarningTitle returns a boolean if a field has been set.

### SetLogOffWarningTitleNil

`func (o *EditDeliveryGroupRequestModel) SetLogOffWarningTitleNil(b bool)`

 SetLogOffWarningTitleNil sets the value for LogOffWarningTitle to be an explicit nil

### UnsetLogOffWarningTitle
`func (o *EditDeliveryGroupRequestModel) UnsetLogOffWarningTitle()`

UnsetLogOffWarningTitle ensures that no value is present for LogOffWarningTitle, not even an explicit nil
### GetPeakLogOffAction

`func (o *EditDeliveryGroupRequestModel) GetPeakLogOffAction() SessionChangeHostingAction`

GetPeakLogOffAction returns the PeakLogOffAction field if non-nil, zero value otherwise.

### GetPeakLogOffActionOk

`func (o *EditDeliveryGroupRequestModel) GetPeakLogOffActionOk() (*SessionChangeHostingAction, bool)`

GetPeakLogOffActionOk returns a tuple with the PeakLogOffAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLogOffAction

`func (o *EditDeliveryGroupRequestModel) SetPeakLogOffAction(v SessionChangeHostingAction)`

SetPeakLogOffAction sets PeakLogOffAction field to given value.

### HasPeakLogOffAction

`func (o *EditDeliveryGroupRequestModel) HasPeakLogOffAction() bool`

HasPeakLogOffAction returns a boolean if a field has been set.

### GetPeakLogOffTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) GetPeakLogOffTimeoutMinutes() int32`

GetPeakLogOffTimeoutMinutes returns the PeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakLogOffTimeoutMinutesOk

`func (o *EditDeliveryGroupRequestModel) GetPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetPeakLogOffTimeoutMinutesOk returns a tuple with the PeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLogOffTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) SetPeakLogOffTimeoutMinutes(v int32)`

SetPeakLogOffTimeoutMinutes sets PeakLogOffTimeoutMinutes field to given value.

### HasPeakLogOffTimeoutMinutes

`func (o *EditDeliveryGroupRequestModel) HasPeakLogOffTimeoutMinutes() bool`

HasPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### SetPeakLogOffTimeoutMinutesNil

`func (o *EditDeliveryGroupRequestModel) SetPeakLogOffTimeoutMinutesNil(b bool)`

 SetPeakLogOffTimeoutMinutesNil sets the value for PeakLogOffTimeoutMinutes to be an explicit nil

### UnsetPeakLogOffTimeoutMinutes
`func (o *EditDeliveryGroupRequestModel) UnsetPeakLogOffTimeoutMinutes()`

UnsetPeakLogOffTimeoutMinutes ensures that no value is present for PeakLogOffTimeoutMinutes, not even an explicit nil
### GetDisconnectPeakIdleSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) GetDisconnectPeakIdleSessionAfterSeconds() int32`

GetDisconnectPeakIdleSessionAfterSeconds returns the DisconnectPeakIdleSessionAfterSeconds field if non-nil, zero value otherwise.

### GetDisconnectPeakIdleSessionAfterSecondsOk

`func (o *EditDeliveryGroupRequestModel) GetDisconnectPeakIdleSessionAfterSecondsOk() (*int32, bool)`

GetDisconnectPeakIdleSessionAfterSecondsOk returns a tuple with the DisconnectPeakIdleSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectPeakIdleSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) SetDisconnectPeakIdleSessionAfterSeconds(v int32)`

SetDisconnectPeakIdleSessionAfterSeconds sets DisconnectPeakIdleSessionAfterSeconds field to given value.

### HasDisconnectPeakIdleSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) HasDisconnectPeakIdleSessionAfterSeconds() bool`

HasDisconnectPeakIdleSessionAfterSeconds returns a boolean if a field has been set.

### SetDisconnectPeakIdleSessionAfterSecondsNil

`func (o *EditDeliveryGroupRequestModel) SetDisconnectPeakIdleSessionAfterSecondsNil(b bool)`

 SetDisconnectPeakIdleSessionAfterSecondsNil sets the value for DisconnectPeakIdleSessionAfterSeconds to be an explicit nil

### UnsetDisconnectPeakIdleSessionAfterSeconds
`func (o *EditDeliveryGroupRequestModel) UnsetDisconnectPeakIdleSessionAfterSeconds()`

UnsetDisconnectPeakIdleSessionAfterSeconds ensures that no value is present for DisconnectPeakIdleSessionAfterSeconds, not even an explicit nil
### GetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) GetDisconnectOffPeakIdleSessionAfterSeconds() int32`

GetDisconnectOffPeakIdleSessionAfterSeconds returns the DisconnectOffPeakIdleSessionAfterSeconds field if non-nil, zero value otherwise.

### GetDisconnectOffPeakIdleSessionAfterSecondsOk

`func (o *EditDeliveryGroupRequestModel) GetDisconnectOffPeakIdleSessionAfterSecondsOk() (*int32, bool)`

GetDisconnectOffPeakIdleSessionAfterSecondsOk returns a tuple with the DisconnectOffPeakIdleSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) SetDisconnectOffPeakIdleSessionAfterSeconds(v int32)`

SetDisconnectOffPeakIdleSessionAfterSeconds sets DisconnectOffPeakIdleSessionAfterSeconds field to given value.

### HasDisconnectOffPeakIdleSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) HasDisconnectOffPeakIdleSessionAfterSeconds() bool`

HasDisconnectOffPeakIdleSessionAfterSeconds returns a boolean if a field has been set.

### SetDisconnectOffPeakIdleSessionAfterSecondsNil

`func (o *EditDeliveryGroupRequestModel) SetDisconnectOffPeakIdleSessionAfterSecondsNil(b bool)`

 SetDisconnectOffPeakIdleSessionAfterSecondsNil sets the value for DisconnectOffPeakIdleSessionAfterSeconds to be an explicit nil

### UnsetDisconnectOffPeakIdleSessionAfterSeconds
`func (o *EditDeliveryGroupRequestModel) UnsetDisconnectOffPeakIdleSessionAfterSeconds()`

UnsetDisconnectOffPeakIdleSessionAfterSeconds ensures that no value is present for DisconnectOffPeakIdleSessionAfterSeconds, not even an explicit nil
### GetLogoffPeakDisconnectedSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) GetLogoffPeakDisconnectedSessionAfterSeconds() int32`

GetLogoffPeakDisconnectedSessionAfterSeconds returns the LogoffPeakDisconnectedSessionAfterSeconds field if non-nil, zero value otherwise.

### GetLogoffPeakDisconnectedSessionAfterSecondsOk

`func (o *EditDeliveryGroupRequestModel) GetLogoffPeakDisconnectedSessionAfterSecondsOk() (*int32, bool)`

GetLogoffPeakDisconnectedSessionAfterSecondsOk returns a tuple with the LogoffPeakDisconnectedSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPeakDisconnectedSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) SetLogoffPeakDisconnectedSessionAfterSeconds(v int32)`

SetLogoffPeakDisconnectedSessionAfterSeconds sets LogoffPeakDisconnectedSessionAfterSeconds field to given value.

### HasLogoffPeakDisconnectedSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) HasLogoffPeakDisconnectedSessionAfterSeconds() bool`

HasLogoffPeakDisconnectedSessionAfterSeconds returns a boolean if a field has been set.

### SetLogoffPeakDisconnectedSessionAfterSecondsNil

`func (o *EditDeliveryGroupRequestModel) SetLogoffPeakDisconnectedSessionAfterSecondsNil(b bool)`

 SetLogoffPeakDisconnectedSessionAfterSecondsNil sets the value for LogoffPeakDisconnectedSessionAfterSeconds to be an explicit nil

### UnsetLogoffPeakDisconnectedSessionAfterSeconds
`func (o *EditDeliveryGroupRequestModel) UnsetLogoffPeakDisconnectedSessionAfterSeconds()`

UnsetLogoffPeakDisconnectedSessionAfterSeconds ensures that no value is present for LogoffPeakDisconnectedSessionAfterSeconds, not even an explicit nil
### GetLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) GetLogoffOffPeakDisconnectedSessionAfterSeconds() int32`

GetLogoffOffPeakDisconnectedSessionAfterSeconds returns the LogoffOffPeakDisconnectedSessionAfterSeconds field if non-nil, zero value otherwise.

### GetLogoffOffPeakDisconnectedSessionAfterSecondsOk

`func (o *EditDeliveryGroupRequestModel) GetLogoffOffPeakDisconnectedSessionAfterSecondsOk() (*int32, bool)`

GetLogoffOffPeakDisconnectedSessionAfterSecondsOk returns a tuple with the LogoffOffPeakDisconnectedSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) SetLogoffOffPeakDisconnectedSessionAfterSeconds(v int32)`

SetLogoffOffPeakDisconnectedSessionAfterSeconds sets LogoffOffPeakDisconnectedSessionAfterSeconds field to given value.

### HasLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *EditDeliveryGroupRequestModel) HasLogoffOffPeakDisconnectedSessionAfterSeconds() bool`

HasLogoffOffPeakDisconnectedSessionAfterSeconds returns a boolean if a field has been set.

### SetLogoffOffPeakDisconnectedSessionAfterSecondsNil

`func (o *EditDeliveryGroupRequestModel) SetLogoffOffPeakDisconnectedSessionAfterSecondsNil(b bool)`

 SetLogoffOffPeakDisconnectedSessionAfterSecondsNil sets the value for LogoffOffPeakDisconnectedSessionAfterSeconds to be an explicit nil

### UnsetLogoffOffPeakDisconnectedSessionAfterSeconds
`func (o *EditDeliveryGroupRequestModel) UnsetLogoffOffPeakDisconnectedSessionAfterSeconds()`

UnsetLogoffOffPeakDisconnectedSessionAfterSeconds ensures that no value is present for LogoffOffPeakDisconnectedSessionAfterSeconds, not even an explicit nil
### GetPowerTimeSchemes

`func (o *EditDeliveryGroupRequestModel) GetPowerTimeSchemes() []PowerTimeSchemeRequestModel`

GetPowerTimeSchemes returns the PowerTimeSchemes field if non-nil, zero value otherwise.

### GetPowerTimeSchemesOk

`func (o *EditDeliveryGroupRequestModel) GetPowerTimeSchemesOk() (*[]PowerTimeSchemeRequestModel, bool)`

GetPowerTimeSchemesOk returns a tuple with the PowerTimeSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerTimeSchemes

`func (o *EditDeliveryGroupRequestModel) SetPowerTimeSchemes(v []PowerTimeSchemeRequestModel)`

SetPowerTimeSchemes sets PowerTimeSchemes field to given value.

### HasPowerTimeSchemes

`func (o *EditDeliveryGroupRequestModel) HasPowerTimeSchemes() bool`

HasPowerTimeSchemes returns a boolean if a field has been set.

### SetPowerTimeSchemesNil

`func (o *EditDeliveryGroupRequestModel) SetPowerTimeSchemesNil(b bool)`

 SetPowerTimeSchemesNil sets the value for PowerTimeSchemes to be an explicit nil

### UnsetPowerTimeSchemes
`func (o *EditDeliveryGroupRequestModel) UnsetPowerTimeSchemes()`

UnsetPowerTimeSchemes ensures that no value is present for PowerTimeSchemes, not even an explicit nil
### GetPowerOffDelayMinutes

`func (o *EditDeliveryGroupRequestModel) GetPowerOffDelayMinutes() int32`

GetPowerOffDelayMinutes returns the PowerOffDelayMinutes field if non-nil, zero value otherwise.

### GetPowerOffDelayMinutesOk

`func (o *EditDeliveryGroupRequestModel) GetPowerOffDelayMinutesOk() (*int32, bool)`

GetPowerOffDelayMinutesOk returns a tuple with the PowerOffDelayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffDelayMinutes

`func (o *EditDeliveryGroupRequestModel) SetPowerOffDelayMinutes(v int32)`

SetPowerOffDelayMinutes sets PowerOffDelayMinutes field to given value.

### HasPowerOffDelayMinutes

`func (o *EditDeliveryGroupRequestModel) HasPowerOffDelayMinutes() bool`

HasPowerOffDelayMinutes returns a boolean if a field has been set.

### SetPowerOffDelayMinutesNil

`func (o *EditDeliveryGroupRequestModel) SetPowerOffDelayMinutesNil(b bool)`

 SetPowerOffDelayMinutesNil sets the value for PowerOffDelayMinutes to be an explicit nil

### UnsetPowerOffDelayMinutes
`func (o *EditDeliveryGroupRequestModel) UnsetPowerOffDelayMinutes()`

UnsetPowerOffDelayMinutes ensures that no value is present for PowerOffDelayMinutes, not even an explicit nil
### GetAutoscaleLogOffReminderEnabled

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderEnabled() bool`

GetAutoscaleLogOffReminderEnabled returns the AutoscaleLogOffReminderEnabled field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderEnabledOk

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderEnabledOk() (*bool, bool)`

GetAutoscaleLogOffReminderEnabledOk returns a tuple with the AutoscaleLogOffReminderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderEnabled

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderEnabled(v bool)`

SetAutoscaleLogOffReminderEnabled sets AutoscaleLogOffReminderEnabled field to given value.

### HasAutoscaleLogOffReminderEnabled

`func (o *EditDeliveryGroupRequestModel) HasAutoscaleLogOffReminderEnabled() bool`

HasAutoscaleLogOffReminderEnabled returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderEnabledNil

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderEnabledNil(b bool)`

 SetAutoscaleLogOffReminderEnabledNil sets the value for AutoscaleLogOffReminderEnabled to be an explicit nil

### UnsetAutoscaleLogOffReminderEnabled
`func (o *EditDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderEnabled()`

UnsetAutoscaleLogOffReminderEnabled ensures that no value is present for AutoscaleLogOffReminderEnabled, not even an explicit nil
### GetAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderIntervalSecondsOffPeak() int32`

GetAutoscaleLogOffReminderIntervalSecondsOffPeak returns the AutoscaleLogOffReminderIntervalSecondsOffPeak field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk() (*int32, bool)`

GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk returns a tuple with the AutoscaleLogOffReminderIntervalSecondsOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderIntervalSecondsOffPeak(v int32)`

SetAutoscaleLogOffReminderIntervalSecondsOffPeak sets AutoscaleLogOffReminderIntervalSecondsOffPeak field to given value.

### HasAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *EditDeliveryGroupRequestModel) HasAutoscaleLogOffReminderIntervalSecondsOffPeak() bool`

HasAutoscaleLogOffReminderIntervalSecondsOffPeak returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderIntervalSecondsOffPeakNil

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderIntervalSecondsOffPeakNil(b bool)`

 SetAutoscaleLogOffReminderIntervalSecondsOffPeakNil sets the value for AutoscaleLogOffReminderIntervalSecondsOffPeak to be an explicit nil

### UnsetAutoscaleLogOffReminderIntervalSecondsOffPeak
`func (o *EditDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderIntervalSecondsOffPeak()`

UnsetAutoscaleLogOffReminderIntervalSecondsOffPeak ensures that no value is present for AutoscaleLogOffReminderIntervalSecondsOffPeak, not even an explicit nil
### GetAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderIntervalSecondsPeak() int32`

GetAutoscaleLogOffReminderIntervalSecondsPeak returns the AutoscaleLogOffReminderIntervalSecondsPeak field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderIntervalSecondsPeakOk

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderIntervalSecondsPeakOk() (*int32, bool)`

GetAutoscaleLogOffReminderIntervalSecondsPeakOk returns a tuple with the AutoscaleLogOffReminderIntervalSecondsPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderIntervalSecondsPeak(v int32)`

SetAutoscaleLogOffReminderIntervalSecondsPeak sets AutoscaleLogOffReminderIntervalSecondsPeak field to given value.

### HasAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *EditDeliveryGroupRequestModel) HasAutoscaleLogOffReminderIntervalSecondsPeak() bool`

HasAutoscaleLogOffReminderIntervalSecondsPeak returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderIntervalSecondsPeakNil

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderIntervalSecondsPeakNil(b bool)`

 SetAutoscaleLogOffReminderIntervalSecondsPeakNil sets the value for AutoscaleLogOffReminderIntervalSecondsPeak to be an explicit nil

### UnsetAutoscaleLogOffReminderIntervalSecondsPeak
`func (o *EditDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderIntervalSecondsPeak()`

UnsetAutoscaleLogOffReminderIntervalSecondsPeak ensures that no value is present for AutoscaleLogOffReminderIntervalSecondsPeak, not even an explicit nil
### GetAutoscaleLogOffReminderMessage

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderMessage() string`

GetAutoscaleLogOffReminderMessage returns the AutoscaleLogOffReminderMessage field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderMessageOk

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderMessageOk() (*string, bool)`

GetAutoscaleLogOffReminderMessageOk returns a tuple with the AutoscaleLogOffReminderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderMessage

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderMessage(v string)`

SetAutoscaleLogOffReminderMessage sets AutoscaleLogOffReminderMessage field to given value.

### HasAutoscaleLogOffReminderMessage

`func (o *EditDeliveryGroupRequestModel) HasAutoscaleLogOffReminderMessage() bool`

HasAutoscaleLogOffReminderMessage returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderMessageNil

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderMessageNil(b bool)`

 SetAutoscaleLogOffReminderMessageNil sets the value for AutoscaleLogOffReminderMessage to be an explicit nil

### UnsetAutoscaleLogOffReminderMessage
`func (o *EditDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderMessage()`

UnsetAutoscaleLogOffReminderMessage ensures that no value is present for AutoscaleLogOffReminderMessage, not even an explicit nil
### GetAutoscaleLogOffReminderTitle

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderTitle() string`

GetAutoscaleLogOffReminderTitle returns the AutoscaleLogOffReminderTitle field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderTitleOk

`func (o *EditDeliveryGroupRequestModel) GetAutoscaleLogOffReminderTitleOk() (*string, bool)`

GetAutoscaleLogOffReminderTitleOk returns a tuple with the AutoscaleLogOffReminderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderTitle

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderTitle(v string)`

SetAutoscaleLogOffReminderTitle sets AutoscaleLogOffReminderTitle field to given value.

### HasAutoscaleLogOffReminderTitle

`func (o *EditDeliveryGroupRequestModel) HasAutoscaleLogOffReminderTitle() bool`

HasAutoscaleLogOffReminderTitle returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderTitleNil

`func (o *EditDeliveryGroupRequestModel) SetAutoscaleLogOffReminderTitleNil(b bool)`

 SetAutoscaleLogOffReminderTitleNil sets the value for AutoscaleLogOffReminderTitle to be an explicit nil

### UnsetAutoscaleLogOffReminderTitle
`func (o *EditDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderTitle()`

UnsetAutoscaleLogOffReminderTitle ensures that no value is present for AutoscaleLogOffReminderTitle, not even an explicit nil
### GetMachineCost

`func (o *EditDeliveryGroupRequestModel) GetMachineCost() float64`

GetMachineCost returns the MachineCost field if non-nil, zero value otherwise.

### GetMachineCostOk

`func (o *EditDeliveryGroupRequestModel) GetMachineCostOk() (*float64, bool)`

GetMachineCostOk returns a tuple with the MachineCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCost

`func (o *EditDeliveryGroupRequestModel) SetMachineCost(v float64)`

SetMachineCost sets MachineCost field to given value.

### HasMachineCost

`func (o *EditDeliveryGroupRequestModel) HasMachineCost() bool`

HasMachineCost returns a boolean if a field has been set.

### SetMachineCostNil

`func (o *EditDeliveryGroupRequestModel) SetMachineCostNil(b bool)`

 SetMachineCostNil sets the value for MachineCost to be an explicit nil

### UnsetMachineCost
`func (o *EditDeliveryGroupRequestModel) UnsetMachineCost()`

UnsetMachineCost ensures that no value is present for MachineCost, not even an explicit nil
### GetMachineLogOnType

`func (o *EditDeliveryGroupRequestModel) GetMachineLogOnType() MachineLogOnType`

GetMachineLogOnType returns the MachineLogOnType field if non-nil, zero value otherwise.

### GetMachineLogOnTypeOk

`func (o *EditDeliveryGroupRequestModel) GetMachineLogOnTypeOk() (*MachineLogOnType, bool)`

GetMachineLogOnTypeOk returns a tuple with the MachineLogOnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLogOnType

`func (o *EditDeliveryGroupRequestModel) SetMachineLogOnType(v MachineLogOnType)`

SetMachineLogOnType sets MachineLogOnType field to given value.

### HasMachineLogOnType

`func (o *EditDeliveryGroupRequestModel) HasMachineLogOnType() bool`

HasMachineLogOnType returns a boolean if a field has been set.

### GetPrelaunchSettings

`func (o *EditDeliveryGroupRequestModel) GetPrelaunchSettings() FastApplicationSettingsRequestModel`

GetPrelaunchSettings returns the PrelaunchSettings field if non-nil, zero value otherwise.

### GetPrelaunchSettingsOk

`func (o *EditDeliveryGroupRequestModel) GetPrelaunchSettingsOk() (*FastApplicationSettingsRequestModel, bool)`

GetPrelaunchSettingsOk returns a tuple with the PrelaunchSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrelaunchSettings

`func (o *EditDeliveryGroupRequestModel) SetPrelaunchSettings(v FastApplicationSettingsRequestModel)`

SetPrelaunchSettings sets PrelaunchSettings field to given value.

### HasPrelaunchSettings

`func (o *EditDeliveryGroupRequestModel) HasPrelaunchSettings() bool`

HasPrelaunchSettings returns a boolean if a field has been set.

### GetProtocolPriority

`func (o *EditDeliveryGroupRequestModel) GetProtocolPriority() []ProtocolType`

GetProtocolPriority returns the ProtocolPriority field if non-nil, zero value otherwise.

### GetProtocolPriorityOk

`func (o *EditDeliveryGroupRequestModel) GetProtocolPriorityOk() (*[]ProtocolType, bool)`

GetProtocolPriorityOk returns a tuple with the ProtocolPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPriority

`func (o *EditDeliveryGroupRequestModel) SetProtocolPriority(v []ProtocolType)`

SetProtocolPriority sets ProtocolPriority field to given value.

### HasProtocolPriority

`func (o *EditDeliveryGroupRequestModel) HasProtocolPriority() bool`

HasProtocolPriority returns a boolean if a field has been set.

### SetProtocolPriorityNil

`func (o *EditDeliveryGroupRequestModel) SetProtocolPriorityNil(b bool)`

 SetProtocolPriorityNil sets the value for ProtocolPriority to be an explicit nil

### UnsetProtocolPriority
`func (o *EditDeliveryGroupRequestModel) UnsetProtocolPriority()`

UnsetProtocolPriority ensures that no value is present for ProtocolPriority, not even an explicit nil
### GetRebootSchedules

`func (o *EditDeliveryGroupRequestModel) GetRebootSchedules() []RebootScheduleRequestModel`

GetRebootSchedules returns the RebootSchedules field if non-nil, zero value otherwise.

### GetRebootSchedulesOk

`func (o *EditDeliveryGroupRequestModel) GetRebootSchedulesOk() (*[]RebootScheduleRequestModel, bool)`

GetRebootSchedulesOk returns a tuple with the RebootSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootSchedules

`func (o *EditDeliveryGroupRequestModel) SetRebootSchedules(v []RebootScheduleRequestModel)`

SetRebootSchedules sets RebootSchedules field to given value.

### HasRebootSchedules

`func (o *EditDeliveryGroupRequestModel) HasRebootSchedules() bool`

HasRebootSchedules returns a boolean if a field has been set.

### SetRebootSchedulesNil

`func (o *EditDeliveryGroupRequestModel) SetRebootSchedulesNil(b bool)`

 SetRebootSchedulesNil sets the value for RebootSchedules to be an explicit nil

### UnsetRebootSchedules
`func (o *EditDeliveryGroupRequestModel) UnsetRebootSchedules()`

UnsetRebootSchedules ensures that no value is present for RebootSchedules, not even an explicit nil
### GetSettlementPeriodBeforeAutoShutdownSeconds

`func (o *EditDeliveryGroupRequestModel) GetSettlementPeriodBeforeAutoShutdownSeconds() int32`

GetSettlementPeriodBeforeAutoShutdownSeconds returns the SettlementPeriodBeforeAutoShutdownSeconds field if non-nil, zero value otherwise.

### GetSettlementPeriodBeforeAutoShutdownSecondsOk

`func (o *EditDeliveryGroupRequestModel) GetSettlementPeriodBeforeAutoShutdownSecondsOk() (*int32, bool)`

GetSettlementPeriodBeforeAutoShutdownSecondsOk returns a tuple with the SettlementPeriodBeforeAutoShutdownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementPeriodBeforeAutoShutdownSeconds

`func (o *EditDeliveryGroupRequestModel) SetSettlementPeriodBeforeAutoShutdownSeconds(v int32)`

SetSettlementPeriodBeforeAutoShutdownSeconds sets SettlementPeriodBeforeAutoShutdownSeconds field to given value.

### HasSettlementPeriodBeforeAutoShutdownSeconds

`func (o *EditDeliveryGroupRequestModel) HasSettlementPeriodBeforeAutoShutdownSeconds() bool`

HasSettlementPeriodBeforeAutoShutdownSeconds returns a boolean if a field has been set.

### SetSettlementPeriodBeforeAutoShutdownSecondsNil

`func (o *EditDeliveryGroupRequestModel) SetSettlementPeriodBeforeAutoShutdownSecondsNil(b bool)`

 SetSettlementPeriodBeforeAutoShutdownSecondsNil sets the value for SettlementPeriodBeforeAutoShutdownSeconds to be an explicit nil

### UnsetSettlementPeriodBeforeAutoShutdownSeconds
`func (o *EditDeliveryGroupRequestModel) UnsetSettlementPeriodBeforeAutoShutdownSeconds()`

UnsetSettlementPeriodBeforeAutoShutdownSeconds ensures that no value is present for SettlementPeriodBeforeAutoShutdownSeconds, not even an explicit nil
### GetSettlementPeriodBeforeUseSeconds

`func (o *EditDeliveryGroupRequestModel) GetSettlementPeriodBeforeUseSeconds() int32`

GetSettlementPeriodBeforeUseSeconds returns the SettlementPeriodBeforeUseSeconds field if non-nil, zero value otherwise.

### GetSettlementPeriodBeforeUseSecondsOk

`func (o *EditDeliveryGroupRequestModel) GetSettlementPeriodBeforeUseSecondsOk() (*int32, bool)`

GetSettlementPeriodBeforeUseSecondsOk returns a tuple with the SettlementPeriodBeforeUseSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementPeriodBeforeUseSeconds

`func (o *EditDeliveryGroupRequestModel) SetSettlementPeriodBeforeUseSeconds(v int32)`

SetSettlementPeriodBeforeUseSeconds sets SettlementPeriodBeforeUseSeconds field to given value.

### HasSettlementPeriodBeforeUseSeconds

`func (o *EditDeliveryGroupRequestModel) HasSettlementPeriodBeforeUseSeconds() bool`

HasSettlementPeriodBeforeUseSeconds returns a boolean if a field has been set.

### SetSettlementPeriodBeforeUseSecondsNil

`func (o *EditDeliveryGroupRequestModel) SetSettlementPeriodBeforeUseSecondsNil(b bool)`

 SetSettlementPeriodBeforeUseSecondsNil sets the value for SettlementPeriodBeforeUseSeconds to be an explicit nil

### UnsetSettlementPeriodBeforeUseSeconds
`func (o *EditDeliveryGroupRequestModel) UnsetSettlementPeriodBeforeUseSeconds()`

UnsetSettlementPeriodBeforeUseSeconds ensures that no value is present for SettlementPeriodBeforeUseSeconds, not even an explicit nil
### GetShutdownDesktopsAfterUse

`func (o *EditDeliveryGroupRequestModel) GetShutdownDesktopsAfterUse() bool`

GetShutdownDesktopsAfterUse returns the ShutdownDesktopsAfterUse field if non-nil, zero value otherwise.

### GetShutdownDesktopsAfterUseOk

`func (o *EditDeliveryGroupRequestModel) GetShutdownDesktopsAfterUseOk() (*bool, bool)`

GetShutdownDesktopsAfterUseOk returns a tuple with the ShutdownDesktopsAfterUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDesktopsAfterUse

`func (o *EditDeliveryGroupRequestModel) SetShutdownDesktopsAfterUse(v bool)`

SetShutdownDesktopsAfterUse sets ShutdownDesktopsAfterUse field to given value.

### HasShutdownDesktopsAfterUse

`func (o *EditDeliveryGroupRequestModel) HasShutdownDesktopsAfterUse() bool`

HasShutdownDesktopsAfterUse returns a boolean if a field has been set.

### SetShutdownDesktopsAfterUseNil

`func (o *EditDeliveryGroupRequestModel) SetShutdownDesktopsAfterUseNil(b bool)`

 SetShutdownDesktopsAfterUseNil sets the value for ShutdownDesktopsAfterUse to be an explicit nil

### UnsetShutdownDesktopsAfterUse
`func (o *EditDeliveryGroupRequestModel) UnsetShutdownDesktopsAfterUse()`

UnsetShutdownDesktopsAfterUse ensures that no value is present for ShutdownDesktopsAfterUse, not even an explicit nil
### GetSimpleAccessPolicy

`func (o *EditDeliveryGroupRequestModel) GetSimpleAccessPolicy() SimplifiedAccessPolicyRequestModel`

GetSimpleAccessPolicy returns the SimpleAccessPolicy field if non-nil, zero value otherwise.

### GetSimpleAccessPolicyOk

`func (o *EditDeliveryGroupRequestModel) GetSimpleAccessPolicyOk() (*SimplifiedAccessPolicyRequestModel, bool)`

GetSimpleAccessPolicyOk returns a tuple with the SimpleAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAccessPolicy

`func (o *EditDeliveryGroupRequestModel) SetSimpleAccessPolicy(v SimplifiedAccessPolicyRequestModel)`

SetSimpleAccessPolicy sets SimpleAccessPolicy field to given value.

### HasSimpleAccessPolicy

`func (o *EditDeliveryGroupRequestModel) HasSimpleAccessPolicy() bool`

HasSimpleAccessPolicy returns a boolean if a field has been set.

### GetAdvancedAccessPolicy

`func (o *EditDeliveryGroupRequestModel) GetAdvancedAccessPolicy() []AdvancedAccessPolicyRequestModel`

GetAdvancedAccessPolicy returns the AdvancedAccessPolicy field if non-nil, zero value otherwise.

### GetAdvancedAccessPolicyOk

`func (o *EditDeliveryGroupRequestModel) GetAdvancedAccessPolicyOk() (*[]AdvancedAccessPolicyRequestModel, bool)`

GetAdvancedAccessPolicyOk returns a tuple with the AdvancedAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedAccessPolicy

`func (o *EditDeliveryGroupRequestModel) SetAdvancedAccessPolicy(v []AdvancedAccessPolicyRequestModel)`

SetAdvancedAccessPolicy sets AdvancedAccessPolicy field to given value.

### HasAdvancedAccessPolicy

`func (o *EditDeliveryGroupRequestModel) HasAdvancedAccessPolicy() bool`

HasAdvancedAccessPolicy returns a boolean if a field has been set.

### SetAdvancedAccessPolicyNil

`func (o *EditDeliveryGroupRequestModel) SetAdvancedAccessPolicyNil(b bool)`

 SetAdvancedAccessPolicyNil sets the value for AdvancedAccessPolicy to be an explicit nil

### UnsetAdvancedAccessPolicy
`func (o *EditDeliveryGroupRequestModel) UnsetAdvancedAccessPolicy()`

UnsetAdvancedAccessPolicy ensures that no value is present for AdvancedAccessPolicy, not even an explicit nil
### GetStoreFrontServersForHostedReceiver

`func (o *EditDeliveryGroupRequestModel) GetStoreFrontServersForHostedReceiver() []StoreFrontServerRequestModel`

GetStoreFrontServersForHostedReceiver returns the StoreFrontServersForHostedReceiver field if non-nil, zero value otherwise.

### GetStoreFrontServersForHostedReceiverOk

`func (o *EditDeliveryGroupRequestModel) GetStoreFrontServersForHostedReceiverOk() (*[]StoreFrontServerRequestModel, bool)`

GetStoreFrontServersForHostedReceiverOk returns a tuple with the StoreFrontServersForHostedReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFrontServersForHostedReceiver

`func (o *EditDeliveryGroupRequestModel) SetStoreFrontServersForHostedReceiver(v []StoreFrontServerRequestModel)`

SetStoreFrontServersForHostedReceiver sets StoreFrontServersForHostedReceiver field to given value.

### HasStoreFrontServersForHostedReceiver

`func (o *EditDeliveryGroupRequestModel) HasStoreFrontServersForHostedReceiver() bool`

HasStoreFrontServersForHostedReceiver returns a boolean if a field has been set.

### SetStoreFrontServersForHostedReceiverNil

`func (o *EditDeliveryGroupRequestModel) SetStoreFrontServersForHostedReceiverNil(b bool)`

 SetStoreFrontServersForHostedReceiverNil sets the value for StoreFrontServersForHostedReceiver to be an explicit nil

### UnsetStoreFrontServersForHostedReceiver
`func (o *EditDeliveryGroupRequestModel) UnsetStoreFrontServersForHostedReceiver()`

UnsetStoreFrontServersForHostedReceiver ensures that no value is present for StoreFrontServersForHostedReceiver, not even an explicit nil
### GetTimeZone

`func (o *EditDeliveryGroupRequestModel) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *EditDeliveryGroupRequestModel) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *EditDeliveryGroupRequestModel) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *EditDeliveryGroupRequestModel) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *EditDeliveryGroupRequestModel) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *EditDeliveryGroupRequestModel) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetTurnOnAddedMachines

`func (o *EditDeliveryGroupRequestModel) GetTurnOnAddedMachines() bool`

GetTurnOnAddedMachines returns the TurnOnAddedMachines field if non-nil, zero value otherwise.

### GetTurnOnAddedMachinesOk

`func (o *EditDeliveryGroupRequestModel) GetTurnOnAddedMachinesOk() (*bool, bool)`

GetTurnOnAddedMachinesOk returns a tuple with the TurnOnAddedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOnAddedMachines

`func (o *EditDeliveryGroupRequestModel) SetTurnOnAddedMachines(v bool)`

SetTurnOnAddedMachines sets TurnOnAddedMachines field to given value.

### HasTurnOnAddedMachines

`func (o *EditDeliveryGroupRequestModel) HasTurnOnAddedMachines() bool`

HasTurnOnAddedMachines returns a boolean if a field has been set.

### SetTurnOnAddedMachinesNil

`func (o *EditDeliveryGroupRequestModel) SetTurnOnAddedMachinesNil(b bool)`

 SetTurnOnAddedMachinesNil sets the value for TurnOnAddedMachines to be an explicit nil

### UnsetTurnOnAddedMachines
`func (o *EditDeliveryGroupRequestModel) UnsetTurnOnAddedMachines()`

UnsetTurnOnAddedMachines ensures that no value is present for TurnOnAddedMachines, not even an explicit nil
### GetZonePreferences

`func (o *EditDeliveryGroupRequestModel) GetZonePreferences() []ZonePreference`

GetZonePreferences returns the ZonePreferences field if non-nil, zero value otherwise.

### GetZonePreferencesOk

`func (o *EditDeliveryGroupRequestModel) GetZonePreferencesOk() (*[]ZonePreference, bool)`

GetZonePreferencesOk returns a tuple with the ZonePreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePreferences

`func (o *EditDeliveryGroupRequestModel) SetZonePreferences(v []ZonePreference)`

SetZonePreferences sets ZonePreferences field to given value.

### HasZonePreferences

`func (o *EditDeliveryGroupRequestModel) HasZonePreferences() bool`

HasZonePreferences returns a boolean if a field has been set.

### SetZonePreferencesNil

`func (o *EditDeliveryGroupRequestModel) SetZonePreferencesNil(b bool)`

 SetZonePreferencesNil sets the value for ZonePreferences to be an explicit nil

### UnsetZonePreferences
`func (o *EditDeliveryGroupRequestModel) UnsetZonePreferences()`

UnsetZonePreferences ensures that no value is present for ZonePreferences, not even an explicit nil
### GetMetadata

`func (o *EditDeliveryGroupRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EditDeliveryGroupRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EditDeliveryGroupRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EditDeliveryGroupRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EditDeliveryGroupRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EditDeliveryGroupRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPolicySetGuid

`func (o *EditDeliveryGroupRequestModel) GetPolicySetGuid() string`

GetPolicySetGuid returns the PolicySetGuid field if non-nil, zero value otherwise.

### GetPolicySetGuidOk

`func (o *EditDeliveryGroupRequestModel) GetPolicySetGuidOk() (*string, bool)`

GetPolicySetGuidOk returns a tuple with the PolicySetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetGuid

`func (o *EditDeliveryGroupRequestModel) SetPolicySetGuid(v string)`

SetPolicySetGuid sets PolicySetGuid field to given value.

### HasPolicySetGuid

`func (o *EditDeliveryGroupRequestModel) HasPolicySetGuid() bool`

HasPolicySetGuid returns a boolean if a field has been set.

### SetPolicySetGuidNil

`func (o *EditDeliveryGroupRequestModel) SetPolicySetGuidNil(b bool)`

 SetPolicySetGuidNil sets the value for PolicySetGuid to be an explicit nil

### UnsetPolicySetGuid
`func (o *EditDeliveryGroupRequestModel) UnsetPolicySetGuid()`

UnsetPolicySetGuid ensures that no value is present for PolicySetGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


