# CreateDeliveryGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminFolder** | Pointer to **NullableString** | The admin folder in which the delivery group should be created. | [optional] 
**Applications** | Pointer to [**DeliveryGroupAddApplicationsRequestModel**](DeliveryGroupAddApplicationsRequestModel.md) |  | [optional] 
**UserManagement** | Pointer to [**UserManagementModel**](UserManagementModel.md) |  | [optional] 
**DeliveryType** | Pointer to [**DeliveryKind**](DeliveryKind.md) |  | [optional] 
**Description** | Pointer to **NullableString** | A description for this delivery group useful for administrators of the site. | [optional] 
**Desktops** | Pointer to [**[]DesktopRequestModel**](DesktopRequestModel.md) | List of desktop resources to publish on the delivery group. | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the delivery group should be in the enabled state; all resources published on disabled delivery groups do not appear to users. | [optional] [default to true]
**InMaintenanceMode** | Pointer to **NullableBool** | Whether the delivery group should be created in maintenance mode; a delivery group in maintenance mode will not allow users to connect or reconnect to machines in the delivery group. | [optional] [default to false]
**MachineCatalogs** | Pointer to [**[]DeliveryGroupAddMachinesRequestModel**](DeliveryGroupAddMachinesRequestModel.md) | List of machine catalogs from which to assign machines to the newly created delivery group. | [optional] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) |  | [optional] 
**SharingKind** | Pointer to [**SharingKind**](SharingKind.md) |  | [optional] 
**IsRemotePC** | Pointer to **NullableBool** | Optional, specifies whether this is to be a Remote PC delivery group. Default comes from the machine catalog, only needs to be specified when creating a delivery group without machine catalog. | [optional] 
**MinimumFunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**Name** | **string** | The name of the new delivery group. | 
**RequireUserHomeZone** | Pointer to **NullableBool** | Whether to require the resources from this delivery group to launch within the user&#39;s home zone. | [optional] [default to false]
**Scopes** | Pointer to **[]string** | Administrative scopes which the newly created delivery group will be a part of. | [optional] 
**Tenants** | Pointer to **[]string** | Tenant(s) to associate the delivery group with. | [optional] 
**AppProtectionKeyLoggingRequired** | Pointer to **NullableBool** | Specifies whether key logging app protection is required. | [optional] 
**AppProtectionScreenCaptureRequired** | Pointer to **NullableBool** | Specifies whether screen capture app protection is required. | [optional] 
**AppAccessPolicy** | Pointer to [**AppAccessPolicyRequestModel**](AppAccessPolicyRequestModel.md) |  | [optional] 
**AutomaticPowerOnForAssigned** | Pointer to **NullableBool** | Whether assigned (Private) machines in the delivery group should be automatically powered-on at the start of peak time periods. | [optional] 
**AutomaticPowerOnForAssignedDuringPeak** | Pointer to **NullableBool** | Whether assigned (Private) machines in the delivery group should be automatically powered-on throughout peak time periods. | [optional] 
**AutoScaleEnabled** | Pointer to **NullableBool** | Whether auto-scale is enabled for the delivery group. | [optional] 
**RestrictAutoscaleTag** | Pointer to **NullableString** | Specifies the tag used to restrict autoscale. | [optional] 
**ColorDepth** | Pointer to [**ColorDepth**](ColorDepth.md) |  | [optional] 
**DefaultDesktopIcon** | Pointer to **NullableString** | Default icon to use for desktops published from the delivery group. | [optional] 
**DefaultDesktopPublishedName** | Pointer to **NullableString** | Default published name to use for desktops published from the delivery group. | [optional] 
**HdxSslEnabled** | Pointer to **NullableBool** | Whether connections to machines in the delivery group will use SSL. | [optional] [default to false]
**LingerSettings** | Pointer to [**FastApplicationSettingsRequestModel**](FastApplicationSettingsRequestModel.md) |  | [optional] 
**OffPeakBufferSizePercent** | Pointer to **NullableInt32** | The percentage of machines in the delivery group that should be kept available in an idle state outside peak hours. | [optional] [default to 0]
**OffPeakDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be performed after a user session disconnects outside peak hours. | [optional] [default to 0]
**OffPeakExtendedDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakExtendedDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the second configured action should be performed after a user session disconnects outside peak hours. | [optional] [default to 0]
**OffPeakLogOffAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakLogOffTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be performed after a user session ends outside peak hours. | [optional] [default to 0]
**LimitSecondsToForceLogOffUserDuringPeak** | Pointer to **NullableInt32** | Represents the number of seconds that must elapsed before the active sessions on the draining machines belonging to the delivery group are logged off, during peak time. | [optional] [default to 0]
**LimitSecondsToForceLogOffUserDuringOffPeak** | Pointer to **NullableInt32** | represents the number of seconds that must elapsed before the active sessions on the draining machines belonging to the delivery group are logged off, during off-peak. | [optional] [default to 0]
**LogOffWarningMessage** | Pointer to **NullableString** | The warning message to display to users in active sessions prior to logging off users, whether in peak time or off-peak. | [optional] 
**LogOffWarningTitle** | Pointer to **NullableString** | The title of the warning message dialog. | [optional] 
**LoadBalanceType** | Pointer to [**LoadBalanceType**](LoadBalanceType.md) |  | [optional] 
**RestrictAutoscaleMinIdleUntaggedPercentDuringPeak** | Pointer to **NullableInt32** | The percentage that the number of untagged single-session machines in an idle state, or for multi-session machines, the untagged available load capacity must fall below before Autoscale powers on and manages &#39;tagged&#39; machines, as per policy, in peak time. If the number of untagged machines in an idle state, or the untagged available load capacity goes above this threshold value, Autoscale will attempt to shut down &#39;tagged&#39; machines. | [optional] [default to -1]
**RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak** | Pointer to **NullableInt32** | The percentage that the number of untagged single-session machines in an idle state, or for multi-session machines, the untagged available load capacity must fall below before Autoscale powers on and manages &#39;tagged&#39; machines, as per policy, in off-peak. If the number of untagged machines in an idle state, or the untagged available load capacity goes above this threshold value, Autoscale will attempt to shut down &#39;tagged&#39; machines. | [optional] [default to -1]
**AutoscaleLogOffReminderEnabled** | Pointer to **NullableBool** | Boolean value indicating whether the warning messages should be sent on an interval to nudge a logoff should be sent on an interval when autoscale is enabled. | [optional] [default to false]
**AutoscaleLogOffReminderIntervalSecondsOffPeak** | Pointer to **NullableInt32** | Represents the time interval at which messages are  sent to the user during off peak time when autoscale is enabled. This message will nudge users to log off instead of forcibly logging them off. | [optional] [default to 0]
**AutoscaleLogOffReminderIntervalSecondsPeak** | Pointer to **NullableInt32** | Represents the time interval at which messages are  sent to the user during peak time when autoscale is enabled. This message will nudge users to log off instead of forcibly logging them off. | [optional] [default to 0]
**PeakAutoscaleAssignedPowerOnIdleTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action is performed on an assigned machine previously started by autoscale that subsequently remains unused. | [optional] 
**PeakAutoscaleAssignedPowerOnIdleAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**AutoscaleLogOffReminderMessage** | Pointer to **NullableString** | Notification message to display to users in active sessions belonging to machines needed by Autoscale for shutdown. | [optional] 
**AutoscaleLogOffReminderTitle** | Pointer to **NullableString** | Notification message dialog title displayed when Autoscale issues a logoff reminder request. | [optional] 
**PeakBufferSizePercent** | Pointer to **NullableInt32** | The percentage of machines in the delivery group that should be kept available in an idle state in peak hours. | [optional] [default to 0]
**PeakDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**PeakDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be performed after a user session disconnects in peak hours. | [optional] [default to 0]
**PeakExtendedDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**PeakExtendedDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the second configured action should be performed after a user session disconnects in peak hours. | [optional] [default to 0]
**PeakLogOffAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**PeakLogOffTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be performed after a user session ends in peak hours. | [optional] [default to 0]
**DisconnectPeakIdleSessionAfterSeconds** | Pointer to **NullableInt32** | Specifies the time in seconds after which an idle session belonging to the delivery group is disconnected during peak time. | [optional] [default to 0]
**DisconnectOffPeakIdleSessionAfterSeconds** | Pointer to **NullableInt32** | Specifies the time in seconds after which an idle session belonging to the delivery group is disconnected during off-peak time. | [optional] [default to 0]
**LogoffPeakDisconnectedSessionAfterSeconds** | Pointer to **NullableInt32** | Specifies the time in seconds after which a disconnected session belonging to the delivery group is terminated during peak time. | [optional] [default to 0]
**LogoffOffPeakDisconnectedSessionAfterSeconds** | Pointer to **NullableInt32** | Specifies the time in seconds after which a disconnected session belonging to the delivery group is terminated during off-peak time. | [optional] [default to 0]
**AutoscaleScaleDownActionDuringPeak** | Pointer to [**AutoscaleScaleDownAction**](AutoscaleScaleDownAction.md) |  | [optional] 
**AutoscaleScaleDownActionDuringOffPeak** | Pointer to [**AutoscaleScaleDownAction**](AutoscaleScaleDownAction.md) |  | [optional] 
**PrelaunchSettings** | Pointer to [**FastApplicationSettingsRequestModel**](FastApplicationSettingsRequestModel.md) |  | [optional] 
**PowerTimeSchemes** | Pointer to [**[]PowerTimeSchemeRequestModel**](PowerTimeSchemeRequestModel.md) | Power management time schemes.  No two schemes for the same delivery group may cover the same day of the week. | [optional] 
**PowerOffDelayMinutes** | Pointer to **NullableInt32** | Delay before machines are powered off, when scaling down.  Specified in minutes.  Applies only to multi-session machines. | [optional] [default to 30]
**MachineCost** | Pointer to **NullableFloat64** | Indicates the estimated per-hour cost for machines in the delivery group, as set by the administrator. | [optional] 
**MachineLogOnType** | Pointer to [**MachineLogOnType**](MachineLogOnType.md) |  | [optional] 
**ProductCode** | Pointer to **NullableString** | The product code of the delivery group. | [optional] 
**ReuseMachinesWithoutShutdownInOutage** | Pointer to **NullableBool** | Whether machines will be reused without a shutdown in case of an outage. | [optional] 
**LicenseModel** | Pointer to [**LicenseModel**](LicenseModel.md) |  | [optional] 
**ProtocolPriority** | Pointer to [**[]ProtocolType**](ProtocolType.md) | A list of protocols in the order in which they should be attempted for use during connection. | [optional] 
**RebootSchedules** | Pointer to [**[]RebootScheduleRequestModel**](RebootScheduleRequestModel.md) | Reboot schedules.  No two schedules for the same delivery group may cover the same day of the week. | [optional] 
**SecureIcaRequired** | Pointer to **NullableBool** | Whether HDX connections to machines in the delivery group require the use of the SecureICA protocol. | [optional] 
**SettlementPeriodBeforeAutoShutdownSeconds** | Pointer to **NullableInt32** | Time after a session ends during which automatic shutdown requests (for example, shutdown after use, idle pool management) are deferred. Any outstanding shutdown request takes effect after the settlement period expires. This is typically used to configure time to allow logoff scripts to complete. | [optional] [default to 0]
**SettlementPeriodBeforeUseSeconds** | Pointer to **NullableInt32** | Idle period before a machine can be selected to host a new session after registration or the end of a previous session. This is typically used to allow a machine to become idle following processing associated with start-up or logoff actions. A machine may still be selected during the idle period if no other machine is available for immediate use. | [optional] [default to 0]
**ShutdownDesktopsAfterUse** | Pointer to **NullableBool** | Whether machines in this delivery group should be automatically shut down when each user session completes. | [optional] [default to false]
**SimpleAccessPolicy** | Pointer to [**SimplifiedAccessPolicyRequestModel**](SimplifiedAccessPolicyRequestModel.md) |  | [optional] 
**StoreFrontServersForHostedReceiver** | Pointer to [**[]StoreFrontServerRequestModel**](StoreFrontServerRequestModel.md) | List of StoreFront server addresses to configure within hosted receivers that are delivered from the delivery group. | [optional] 
**TimeZone** | Pointer to **NullableString** | The time zone in which this delivery group&#39;s machines reside. | [optional] 
**TurnOnAddedMachines** | Pointer to **NullableBool** | Whether to attempt to power on machines when they are added to the delivery group. | [optional] 
**ZonePreferences** | Pointer to [**[]ZonePreference**](ZonePreference.md) | Ordered list of zone preferences to be applied when launching resources from this delivery group. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of the new delivery group. | [optional] 
**PolicySetGuid** | Pointer to **string** | The GUID of the policy set assigned to this desktop group. | [optional] 
**RequiredSleepCapability** | Pointer to [**RequiredSleepCapability**](RequiredSleepCapability.md) |  | [optional] 

## Methods

### NewCreateDeliveryGroupRequestModel

`func NewCreateDeliveryGroupRequestModel(name string, ) *CreateDeliveryGroupRequestModel`

NewCreateDeliveryGroupRequestModel instantiates a new CreateDeliveryGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeliveryGroupRequestModelWithDefaults

`func NewCreateDeliveryGroupRequestModelWithDefaults() *CreateDeliveryGroupRequestModel`

NewCreateDeliveryGroupRequestModelWithDefaults instantiates a new CreateDeliveryGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminFolder

`func (o *CreateDeliveryGroupRequestModel) GetAdminFolder() string`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *CreateDeliveryGroupRequestModel) GetAdminFolderOk() (*string, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *CreateDeliveryGroupRequestModel) SetAdminFolder(v string)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *CreateDeliveryGroupRequestModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### SetAdminFolderNil

`func (o *CreateDeliveryGroupRequestModel) SetAdminFolderNil(b bool)`

 SetAdminFolderNil sets the value for AdminFolder to be an explicit nil

### UnsetAdminFolder
`func (o *CreateDeliveryGroupRequestModel) UnsetAdminFolder()`

UnsetAdminFolder ensures that no value is present for AdminFolder, not even an explicit nil
### GetApplications

`func (o *CreateDeliveryGroupRequestModel) GetApplications() DeliveryGroupAddApplicationsRequestModel`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *CreateDeliveryGroupRequestModel) GetApplicationsOk() (*DeliveryGroupAddApplicationsRequestModel, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *CreateDeliveryGroupRequestModel) SetApplications(v DeliveryGroupAddApplicationsRequestModel)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *CreateDeliveryGroupRequestModel) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetUserManagement

`func (o *CreateDeliveryGroupRequestModel) GetUserManagement() UserManagementModel`

GetUserManagement returns the UserManagement field if non-nil, zero value otherwise.

### GetUserManagementOk

`func (o *CreateDeliveryGroupRequestModel) GetUserManagementOk() (*UserManagementModel, bool)`

GetUserManagementOk returns a tuple with the UserManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagement

`func (o *CreateDeliveryGroupRequestModel) SetUserManagement(v UserManagementModel)`

SetUserManagement sets UserManagement field to given value.

### HasUserManagement

`func (o *CreateDeliveryGroupRequestModel) HasUserManagement() bool`

HasUserManagement returns a boolean if a field has been set.

### GetDeliveryType

`func (o *CreateDeliveryGroupRequestModel) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *CreateDeliveryGroupRequestModel) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *CreateDeliveryGroupRequestModel) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *CreateDeliveryGroupRequestModel) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDeliveryGroupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDeliveryGroupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDeliveryGroupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDeliveryGroupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateDeliveryGroupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateDeliveryGroupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesktops

`func (o *CreateDeliveryGroupRequestModel) GetDesktops() []DesktopRequestModel`

GetDesktops returns the Desktops field if non-nil, zero value otherwise.

### GetDesktopsOk

`func (o *CreateDeliveryGroupRequestModel) GetDesktopsOk() (*[]DesktopRequestModel, bool)`

GetDesktopsOk returns a tuple with the Desktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktops

`func (o *CreateDeliveryGroupRequestModel) SetDesktops(v []DesktopRequestModel)`

SetDesktops sets Desktops field to given value.

### HasDesktops

`func (o *CreateDeliveryGroupRequestModel) HasDesktops() bool`

HasDesktops returns a boolean if a field has been set.

### SetDesktopsNil

`func (o *CreateDeliveryGroupRequestModel) SetDesktopsNil(b bool)`

 SetDesktopsNil sets the value for Desktops to be an explicit nil

### UnsetDesktops
`func (o *CreateDeliveryGroupRequestModel) UnsetDesktops()`

UnsetDesktops ensures that no value is present for Desktops, not even an explicit nil
### GetEnabled

`func (o *CreateDeliveryGroupRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateDeliveryGroupRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateDeliveryGroupRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateDeliveryGroupRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *CreateDeliveryGroupRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *CreateDeliveryGroupRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetInMaintenanceMode

`func (o *CreateDeliveryGroupRequestModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *CreateDeliveryGroupRequestModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *CreateDeliveryGroupRequestModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *CreateDeliveryGroupRequestModel) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### SetInMaintenanceModeNil

`func (o *CreateDeliveryGroupRequestModel) SetInMaintenanceModeNil(b bool)`

 SetInMaintenanceModeNil sets the value for InMaintenanceMode to be an explicit nil

### UnsetInMaintenanceMode
`func (o *CreateDeliveryGroupRequestModel) UnsetInMaintenanceMode()`

UnsetInMaintenanceMode ensures that no value is present for InMaintenanceMode, not even an explicit nil
### GetMachineCatalogs

`func (o *CreateDeliveryGroupRequestModel) GetMachineCatalogs() []DeliveryGroupAddMachinesRequestModel`

GetMachineCatalogs returns the MachineCatalogs field if non-nil, zero value otherwise.

### GetMachineCatalogsOk

`func (o *CreateDeliveryGroupRequestModel) GetMachineCatalogsOk() (*[]DeliveryGroupAddMachinesRequestModel, bool)`

GetMachineCatalogsOk returns a tuple with the MachineCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCatalogs

`func (o *CreateDeliveryGroupRequestModel) SetMachineCatalogs(v []DeliveryGroupAddMachinesRequestModel)`

SetMachineCatalogs sets MachineCatalogs field to given value.

### HasMachineCatalogs

`func (o *CreateDeliveryGroupRequestModel) HasMachineCatalogs() bool`

HasMachineCatalogs returns a boolean if a field has been set.

### SetMachineCatalogsNil

`func (o *CreateDeliveryGroupRequestModel) SetMachineCatalogsNil(b bool)`

 SetMachineCatalogsNil sets the value for MachineCatalogs to be an explicit nil

### UnsetMachineCatalogs
`func (o *CreateDeliveryGroupRequestModel) UnsetMachineCatalogs()`

UnsetMachineCatalogs ensures that no value is present for MachineCatalogs, not even an explicit nil
### GetSessionSupport

`func (o *CreateDeliveryGroupRequestModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *CreateDeliveryGroupRequestModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *CreateDeliveryGroupRequestModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *CreateDeliveryGroupRequestModel) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### GetSharingKind

`func (o *CreateDeliveryGroupRequestModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *CreateDeliveryGroupRequestModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *CreateDeliveryGroupRequestModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.

### HasSharingKind

`func (o *CreateDeliveryGroupRequestModel) HasSharingKind() bool`

HasSharingKind returns a boolean if a field has been set.

### GetIsRemotePC

`func (o *CreateDeliveryGroupRequestModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *CreateDeliveryGroupRequestModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *CreateDeliveryGroupRequestModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.

### HasIsRemotePC

`func (o *CreateDeliveryGroupRequestModel) HasIsRemotePC() bool`

HasIsRemotePC returns a boolean if a field has been set.

### SetIsRemotePCNil

`func (o *CreateDeliveryGroupRequestModel) SetIsRemotePCNil(b bool)`

 SetIsRemotePCNil sets the value for IsRemotePC to be an explicit nil

### UnsetIsRemotePC
`func (o *CreateDeliveryGroupRequestModel) UnsetIsRemotePC()`

UnsetIsRemotePC ensures that no value is present for IsRemotePC, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *CreateDeliveryGroupRequestModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *CreateDeliveryGroupRequestModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *CreateDeliveryGroupRequestModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.

### HasMinimumFunctionalLevel

`func (o *CreateDeliveryGroupRequestModel) HasMinimumFunctionalLevel() bool`

HasMinimumFunctionalLevel returns a boolean if a field has been set.

### GetName

`func (o *CreateDeliveryGroupRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDeliveryGroupRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDeliveryGroupRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetRequireUserHomeZone

`func (o *CreateDeliveryGroupRequestModel) GetRequireUserHomeZone() bool`

GetRequireUserHomeZone returns the RequireUserHomeZone field if non-nil, zero value otherwise.

### GetRequireUserHomeZoneOk

`func (o *CreateDeliveryGroupRequestModel) GetRequireUserHomeZoneOk() (*bool, bool)`

GetRequireUserHomeZoneOk returns a tuple with the RequireUserHomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireUserHomeZone

`func (o *CreateDeliveryGroupRequestModel) SetRequireUserHomeZone(v bool)`

SetRequireUserHomeZone sets RequireUserHomeZone field to given value.

### HasRequireUserHomeZone

`func (o *CreateDeliveryGroupRequestModel) HasRequireUserHomeZone() bool`

HasRequireUserHomeZone returns a boolean if a field has been set.

### SetRequireUserHomeZoneNil

`func (o *CreateDeliveryGroupRequestModel) SetRequireUserHomeZoneNil(b bool)`

 SetRequireUserHomeZoneNil sets the value for RequireUserHomeZone to be an explicit nil

### UnsetRequireUserHomeZone
`func (o *CreateDeliveryGroupRequestModel) UnsetRequireUserHomeZone()`

UnsetRequireUserHomeZone ensures that no value is present for RequireUserHomeZone, not even an explicit nil
### GetScopes

`func (o *CreateDeliveryGroupRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateDeliveryGroupRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateDeliveryGroupRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateDeliveryGroupRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *CreateDeliveryGroupRequestModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *CreateDeliveryGroupRequestModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *CreateDeliveryGroupRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateDeliveryGroupRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateDeliveryGroupRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateDeliveryGroupRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *CreateDeliveryGroupRequestModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *CreateDeliveryGroupRequestModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetAppProtectionKeyLoggingRequired

`func (o *CreateDeliveryGroupRequestModel) GetAppProtectionKeyLoggingRequired() bool`

GetAppProtectionKeyLoggingRequired returns the AppProtectionKeyLoggingRequired field if non-nil, zero value otherwise.

### GetAppProtectionKeyLoggingRequiredOk

`func (o *CreateDeliveryGroupRequestModel) GetAppProtectionKeyLoggingRequiredOk() (*bool, bool)`

GetAppProtectionKeyLoggingRequiredOk returns a tuple with the AppProtectionKeyLoggingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtectionKeyLoggingRequired

`func (o *CreateDeliveryGroupRequestModel) SetAppProtectionKeyLoggingRequired(v bool)`

SetAppProtectionKeyLoggingRequired sets AppProtectionKeyLoggingRequired field to given value.

### HasAppProtectionKeyLoggingRequired

`func (o *CreateDeliveryGroupRequestModel) HasAppProtectionKeyLoggingRequired() bool`

HasAppProtectionKeyLoggingRequired returns a boolean if a field has been set.

### SetAppProtectionKeyLoggingRequiredNil

`func (o *CreateDeliveryGroupRequestModel) SetAppProtectionKeyLoggingRequiredNil(b bool)`

 SetAppProtectionKeyLoggingRequiredNil sets the value for AppProtectionKeyLoggingRequired to be an explicit nil

### UnsetAppProtectionKeyLoggingRequired
`func (o *CreateDeliveryGroupRequestModel) UnsetAppProtectionKeyLoggingRequired()`

UnsetAppProtectionKeyLoggingRequired ensures that no value is present for AppProtectionKeyLoggingRequired, not even an explicit nil
### GetAppProtectionScreenCaptureRequired

`func (o *CreateDeliveryGroupRequestModel) GetAppProtectionScreenCaptureRequired() bool`

GetAppProtectionScreenCaptureRequired returns the AppProtectionScreenCaptureRequired field if non-nil, zero value otherwise.

### GetAppProtectionScreenCaptureRequiredOk

`func (o *CreateDeliveryGroupRequestModel) GetAppProtectionScreenCaptureRequiredOk() (*bool, bool)`

GetAppProtectionScreenCaptureRequiredOk returns a tuple with the AppProtectionScreenCaptureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtectionScreenCaptureRequired

`func (o *CreateDeliveryGroupRequestModel) SetAppProtectionScreenCaptureRequired(v bool)`

SetAppProtectionScreenCaptureRequired sets AppProtectionScreenCaptureRequired field to given value.

### HasAppProtectionScreenCaptureRequired

`func (o *CreateDeliveryGroupRequestModel) HasAppProtectionScreenCaptureRequired() bool`

HasAppProtectionScreenCaptureRequired returns a boolean if a field has been set.

### SetAppProtectionScreenCaptureRequiredNil

`func (o *CreateDeliveryGroupRequestModel) SetAppProtectionScreenCaptureRequiredNil(b bool)`

 SetAppProtectionScreenCaptureRequiredNil sets the value for AppProtectionScreenCaptureRequired to be an explicit nil

### UnsetAppProtectionScreenCaptureRequired
`func (o *CreateDeliveryGroupRequestModel) UnsetAppProtectionScreenCaptureRequired()`

UnsetAppProtectionScreenCaptureRequired ensures that no value is present for AppProtectionScreenCaptureRequired, not even an explicit nil
### GetAppAccessPolicy

`func (o *CreateDeliveryGroupRequestModel) GetAppAccessPolicy() AppAccessPolicyRequestModel`

GetAppAccessPolicy returns the AppAccessPolicy field if non-nil, zero value otherwise.

### GetAppAccessPolicyOk

`func (o *CreateDeliveryGroupRequestModel) GetAppAccessPolicyOk() (*AppAccessPolicyRequestModel, bool)`

GetAppAccessPolicyOk returns a tuple with the AppAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAccessPolicy

`func (o *CreateDeliveryGroupRequestModel) SetAppAccessPolicy(v AppAccessPolicyRequestModel)`

SetAppAccessPolicy sets AppAccessPolicy field to given value.

### HasAppAccessPolicy

`func (o *CreateDeliveryGroupRequestModel) HasAppAccessPolicy() bool`

HasAppAccessPolicy returns a boolean if a field has been set.

### GetAutomaticPowerOnForAssigned

`func (o *CreateDeliveryGroupRequestModel) GetAutomaticPowerOnForAssigned() bool`

GetAutomaticPowerOnForAssigned returns the AutomaticPowerOnForAssigned field if non-nil, zero value otherwise.

### GetAutomaticPowerOnForAssignedOk

`func (o *CreateDeliveryGroupRequestModel) GetAutomaticPowerOnForAssignedOk() (*bool, bool)`

GetAutomaticPowerOnForAssignedOk returns a tuple with the AutomaticPowerOnForAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPowerOnForAssigned

`func (o *CreateDeliveryGroupRequestModel) SetAutomaticPowerOnForAssigned(v bool)`

SetAutomaticPowerOnForAssigned sets AutomaticPowerOnForAssigned field to given value.

### HasAutomaticPowerOnForAssigned

`func (o *CreateDeliveryGroupRequestModel) HasAutomaticPowerOnForAssigned() bool`

HasAutomaticPowerOnForAssigned returns a boolean if a field has been set.

### SetAutomaticPowerOnForAssignedNil

`func (o *CreateDeliveryGroupRequestModel) SetAutomaticPowerOnForAssignedNil(b bool)`

 SetAutomaticPowerOnForAssignedNil sets the value for AutomaticPowerOnForAssigned to be an explicit nil

### UnsetAutomaticPowerOnForAssigned
`func (o *CreateDeliveryGroupRequestModel) UnsetAutomaticPowerOnForAssigned()`

UnsetAutomaticPowerOnForAssigned ensures that no value is present for AutomaticPowerOnForAssigned, not even an explicit nil
### GetAutomaticPowerOnForAssignedDuringPeak

`func (o *CreateDeliveryGroupRequestModel) GetAutomaticPowerOnForAssignedDuringPeak() bool`

GetAutomaticPowerOnForAssignedDuringPeak returns the AutomaticPowerOnForAssignedDuringPeak field if non-nil, zero value otherwise.

### GetAutomaticPowerOnForAssignedDuringPeakOk

`func (o *CreateDeliveryGroupRequestModel) GetAutomaticPowerOnForAssignedDuringPeakOk() (*bool, bool)`

GetAutomaticPowerOnForAssignedDuringPeakOk returns a tuple with the AutomaticPowerOnForAssignedDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPowerOnForAssignedDuringPeak

`func (o *CreateDeliveryGroupRequestModel) SetAutomaticPowerOnForAssignedDuringPeak(v bool)`

SetAutomaticPowerOnForAssignedDuringPeak sets AutomaticPowerOnForAssignedDuringPeak field to given value.

### HasAutomaticPowerOnForAssignedDuringPeak

`func (o *CreateDeliveryGroupRequestModel) HasAutomaticPowerOnForAssignedDuringPeak() bool`

HasAutomaticPowerOnForAssignedDuringPeak returns a boolean if a field has been set.

### SetAutomaticPowerOnForAssignedDuringPeakNil

`func (o *CreateDeliveryGroupRequestModel) SetAutomaticPowerOnForAssignedDuringPeakNil(b bool)`

 SetAutomaticPowerOnForAssignedDuringPeakNil sets the value for AutomaticPowerOnForAssignedDuringPeak to be an explicit nil

### UnsetAutomaticPowerOnForAssignedDuringPeak
`func (o *CreateDeliveryGroupRequestModel) UnsetAutomaticPowerOnForAssignedDuringPeak()`

UnsetAutomaticPowerOnForAssignedDuringPeak ensures that no value is present for AutomaticPowerOnForAssignedDuringPeak, not even an explicit nil
### GetAutoScaleEnabled

`func (o *CreateDeliveryGroupRequestModel) GetAutoScaleEnabled() bool`

GetAutoScaleEnabled returns the AutoScaleEnabled field if non-nil, zero value otherwise.

### GetAutoScaleEnabledOk

`func (o *CreateDeliveryGroupRequestModel) GetAutoScaleEnabledOk() (*bool, bool)`

GetAutoScaleEnabledOk returns a tuple with the AutoScaleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaleEnabled

`func (o *CreateDeliveryGroupRequestModel) SetAutoScaleEnabled(v bool)`

SetAutoScaleEnabled sets AutoScaleEnabled field to given value.

### HasAutoScaleEnabled

`func (o *CreateDeliveryGroupRequestModel) HasAutoScaleEnabled() bool`

HasAutoScaleEnabled returns a boolean if a field has been set.

### SetAutoScaleEnabledNil

`func (o *CreateDeliveryGroupRequestModel) SetAutoScaleEnabledNil(b bool)`

 SetAutoScaleEnabledNil sets the value for AutoScaleEnabled to be an explicit nil

### UnsetAutoScaleEnabled
`func (o *CreateDeliveryGroupRequestModel) UnsetAutoScaleEnabled()`

UnsetAutoScaleEnabled ensures that no value is present for AutoScaleEnabled, not even an explicit nil
### GetRestrictAutoscaleTag

`func (o *CreateDeliveryGroupRequestModel) GetRestrictAutoscaleTag() string`

GetRestrictAutoscaleTag returns the RestrictAutoscaleTag field if non-nil, zero value otherwise.

### GetRestrictAutoscaleTagOk

`func (o *CreateDeliveryGroupRequestModel) GetRestrictAutoscaleTagOk() (*string, bool)`

GetRestrictAutoscaleTagOk returns a tuple with the RestrictAutoscaleTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleTag

`func (o *CreateDeliveryGroupRequestModel) SetRestrictAutoscaleTag(v string)`

SetRestrictAutoscaleTag sets RestrictAutoscaleTag field to given value.

### HasRestrictAutoscaleTag

`func (o *CreateDeliveryGroupRequestModel) HasRestrictAutoscaleTag() bool`

HasRestrictAutoscaleTag returns a boolean if a field has been set.

### SetRestrictAutoscaleTagNil

`func (o *CreateDeliveryGroupRequestModel) SetRestrictAutoscaleTagNil(b bool)`

 SetRestrictAutoscaleTagNil sets the value for RestrictAutoscaleTag to be an explicit nil

### UnsetRestrictAutoscaleTag
`func (o *CreateDeliveryGroupRequestModel) UnsetRestrictAutoscaleTag()`

UnsetRestrictAutoscaleTag ensures that no value is present for RestrictAutoscaleTag, not even an explicit nil
### GetColorDepth

`func (o *CreateDeliveryGroupRequestModel) GetColorDepth() ColorDepth`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *CreateDeliveryGroupRequestModel) GetColorDepthOk() (*ColorDepth, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *CreateDeliveryGroupRequestModel) SetColorDepth(v ColorDepth)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *CreateDeliveryGroupRequestModel) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetDefaultDesktopIcon

`func (o *CreateDeliveryGroupRequestModel) GetDefaultDesktopIcon() string`

GetDefaultDesktopIcon returns the DefaultDesktopIcon field if non-nil, zero value otherwise.

### GetDefaultDesktopIconOk

`func (o *CreateDeliveryGroupRequestModel) GetDefaultDesktopIconOk() (*string, bool)`

GetDefaultDesktopIconOk returns a tuple with the DefaultDesktopIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDesktopIcon

`func (o *CreateDeliveryGroupRequestModel) SetDefaultDesktopIcon(v string)`

SetDefaultDesktopIcon sets DefaultDesktopIcon field to given value.

### HasDefaultDesktopIcon

`func (o *CreateDeliveryGroupRequestModel) HasDefaultDesktopIcon() bool`

HasDefaultDesktopIcon returns a boolean if a field has been set.

### SetDefaultDesktopIconNil

`func (o *CreateDeliveryGroupRequestModel) SetDefaultDesktopIconNil(b bool)`

 SetDefaultDesktopIconNil sets the value for DefaultDesktopIcon to be an explicit nil

### UnsetDefaultDesktopIcon
`func (o *CreateDeliveryGroupRequestModel) UnsetDefaultDesktopIcon()`

UnsetDefaultDesktopIcon ensures that no value is present for DefaultDesktopIcon, not even an explicit nil
### GetDefaultDesktopPublishedName

`func (o *CreateDeliveryGroupRequestModel) GetDefaultDesktopPublishedName() string`

GetDefaultDesktopPublishedName returns the DefaultDesktopPublishedName field if non-nil, zero value otherwise.

### GetDefaultDesktopPublishedNameOk

`func (o *CreateDeliveryGroupRequestModel) GetDefaultDesktopPublishedNameOk() (*string, bool)`

GetDefaultDesktopPublishedNameOk returns a tuple with the DefaultDesktopPublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDesktopPublishedName

`func (o *CreateDeliveryGroupRequestModel) SetDefaultDesktopPublishedName(v string)`

SetDefaultDesktopPublishedName sets DefaultDesktopPublishedName field to given value.

### HasDefaultDesktopPublishedName

`func (o *CreateDeliveryGroupRequestModel) HasDefaultDesktopPublishedName() bool`

HasDefaultDesktopPublishedName returns a boolean if a field has been set.

### SetDefaultDesktopPublishedNameNil

`func (o *CreateDeliveryGroupRequestModel) SetDefaultDesktopPublishedNameNil(b bool)`

 SetDefaultDesktopPublishedNameNil sets the value for DefaultDesktopPublishedName to be an explicit nil

### UnsetDefaultDesktopPublishedName
`func (o *CreateDeliveryGroupRequestModel) UnsetDefaultDesktopPublishedName()`

UnsetDefaultDesktopPublishedName ensures that no value is present for DefaultDesktopPublishedName, not even an explicit nil
### GetHdxSslEnabled

`func (o *CreateDeliveryGroupRequestModel) GetHdxSslEnabled() bool`

GetHdxSslEnabled returns the HdxSslEnabled field if non-nil, zero value otherwise.

### GetHdxSslEnabledOk

`func (o *CreateDeliveryGroupRequestModel) GetHdxSslEnabledOk() (*bool, bool)`

GetHdxSslEnabledOk returns a tuple with the HdxSslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdxSslEnabled

`func (o *CreateDeliveryGroupRequestModel) SetHdxSslEnabled(v bool)`

SetHdxSslEnabled sets HdxSslEnabled field to given value.

### HasHdxSslEnabled

`func (o *CreateDeliveryGroupRequestModel) HasHdxSslEnabled() bool`

HasHdxSslEnabled returns a boolean if a field has been set.

### SetHdxSslEnabledNil

`func (o *CreateDeliveryGroupRequestModel) SetHdxSslEnabledNil(b bool)`

 SetHdxSslEnabledNil sets the value for HdxSslEnabled to be an explicit nil

### UnsetHdxSslEnabled
`func (o *CreateDeliveryGroupRequestModel) UnsetHdxSslEnabled()`

UnsetHdxSslEnabled ensures that no value is present for HdxSslEnabled, not even an explicit nil
### GetLingerSettings

`func (o *CreateDeliveryGroupRequestModel) GetLingerSettings() FastApplicationSettingsRequestModel`

GetLingerSettings returns the LingerSettings field if non-nil, zero value otherwise.

### GetLingerSettingsOk

`func (o *CreateDeliveryGroupRequestModel) GetLingerSettingsOk() (*FastApplicationSettingsRequestModel, bool)`

GetLingerSettingsOk returns a tuple with the LingerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLingerSettings

`func (o *CreateDeliveryGroupRequestModel) SetLingerSettings(v FastApplicationSettingsRequestModel)`

SetLingerSettings sets LingerSettings field to given value.

### HasLingerSettings

`func (o *CreateDeliveryGroupRequestModel) HasLingerSettings() bool`

HasLingerSettings returns a boolean if a field has been set.

### GetOffPeakBufferSizePercent

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakBufferSizePercent() int32`

GetOffPeakBufferSizePercent returns the OffPeakBufferSizePercent field if non-nil, zero value otherwise.

### GetOffPeakBufferSizePercentOk

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakBufferSizePercentOk() (*int32, bool)`

GetOffPeakBufferSizePercentOk returns a tuple with the OffPeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakBufferSizePercent

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakBufferSizePercent(v int32)`

SetOffPeakBufferSizePercent sets OffPeakBufferSizePercent field to given value.

### HasOffPeakBufferSizePercent

`func (o *CreateDeliveryGroupRequestModel) HasOffPeakBufferSizePercent() bool`

HasOffPeakBufferSizePercent returns a boolean if a field has been set.

### SetOffPeakBufferSizePercentNil

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakBufferSizePercentNil(b bool)`

 SetOffPeakBufferSizePercentNil sets the value for OffPeakBufferSizePercent to be an explicit nil

### UnsetOffPeakBufferSizePercent
`func (o *CreateDeliveryGroupRequestModel) UnsetOffPeakBufferSizePercent()`

UnsetOffPeakBufferSizePercent ensures that no value is present for OffPeakBufferSizePercent, not even an explicit nil
### GetOffPeakDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakDisconnectAction() SessionChangeHostingAction`

GetOffPeakDisconnectAction returns the OffPeakDisconnectAction field if non-nil, zero value otherwise.

### GetOffPeakDisconnectActionOk

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakDisconnectActionOk returns a tuple with the OffPeakDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakDisconnectAction(v SessionChangeHostingAction)`

SetOffPeakDisconnectAction sets OffPeakDisconnectAction field to given value.

### HasOffPeakDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) HasOffPeakDisconnectAction() bool`

HasOffPeakDisconnectAction returns a boolean if a field has been set.

### GetOffPeakDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakDisconnectTimeoutMinutes() int32`

GetOffPeakDisconnectTimeoutMinutes returns the OffPeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakDisconnectTimeoutMinutesOk

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakDisconnectTimeoutMinutesOk returns a tuple with the OffPeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakDisconnectTimeoutMinutes(v int32)`

SetOffPeakDisconnectTimeoutMinutes sets OffPeakDisconnectTimeoutMinutes field to given value.

### HasOffPeakDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) HasOffPeakDisconnectTimeoutMinutes() bool`

HasOffPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakDisconnectTimeoutMinutesNil

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakDisconnectTimeoutMinutesNil(b bool)`

 SetOffPeakDisconnectTimeoutMinutesNil sets the value for OffPeakDisconnectTimeoutMinutes to be an explicit nil

### UnsetOffPeakDisconnectTimeoutMinutes
`func (o *CreateDeliveryGroupRequestModel) UnsetOffPeakDisconnectTimeoutMinutes()`

UnsetOffPeakDisconnectTimeoutMinutes ensures that no value is present for OffPeakDisconnectTimeoutMinutes, not even an explicit nil
### GetOffPeakExtendedDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakExtendedDisconnectAction() SessionChangeHostingAction`

GetOffPeakExtendedDisconnectAction returns the OffPeakExtendedDisconnectAction field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectActionOk

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakExtendedDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakExtendedDisconnectActionOk returns a tuple with the OffPeakExtendedDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakExtendedDisconnectAction(v SessionChangeHostingAction)`

SetOffPeakExtendedDisconnectAction sets OffPeakExtendedDisconnectAction field to given value.

### HasOffPeakExtendedDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) HasOffPeakExtendedDisconnectAction() bool`

HasOffPeakExtendedDisconnectAction returns a boolean if a field has been set.

### GetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakExtendedDisconnectTimeoutMinutes() int32`

GetOffPeakExtendedDisconnectTimeoutMinutes returns the OffPeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectTimeoutMinutesOk

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the OffPeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetOffPeakExtendedDisconnectTimeoutMinutes sets OffPeakExtendedDisconnectTimeoutMinutes field to given value.

### HasOffPeakExtendedDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) HasOffPeakExtendedDisconnectTimeoutMinutes() bool`

HasOffPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakExtendedDisconnectTimeoutMinutesNil

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakExtendedDisconnectTimeoutMinutesNil(b bool)`

 SetOffPeakExtendedDisconnectTimeoutMinutesNil sets the value for OffPeakExtendedDisconnectTimeoutMinutes to be an explicit nil

### UnsetOffPeakExtendedDisconnectTimeoutMinutes
`func (o *CreateDeliveryGroupRequestModel) UnsetOffPeakExtendedDisconnectTimeoutMinutes()`

UnsetOffPeakExtendedDisconnectTimeoutMinutes ensures that no value is present for OffPeakExtendedDisconnectTimeoutMinutes, not even an explicit nil
### GetOffPeakLogOffAction

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakLogOffAction() SessionChangeHostingAction`

GetOffPeakLogOffAction returns the OffPeakLogOffAction field if non-nil, zero value otherwise.

### GetOffPeakLogOffActionOk

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakLogOffActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakLogOffActionOk returns a tuple with the OffPeakLogOffAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffAction

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakLogOffAction(v SessionChangeHostingAction)`

SetOffPeakLogOffAction sets OffPeakLogOffAction field to given value.

### HasOffPeakLogOffAction

`func (o *CreateDeliveryGroupRequestModel) HasOffPeakLogOffAction() bool`

HasOffPeakLogOffAction returns a boolean if a field has been set.

### GetOffPeakLogOffTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakLogOffTimeoutMinutes() int32`

GetOffPeakLogOffTimeoutMinutes returns the OffPeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakLogOffTimeoutMinutesOk

`func (o *CreateDeliveryGroupRequestModel) GetOffPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetOffPeakLogOffTimeoutMinutesOk returns a tuple with the OffPeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakLogOffTimeoutMinutes(v int32)`

SetOffPeakLogOffTimeoutMinutes sets OffPeakLogOffTimeoutMinutes field to given value.

### HasOffPeakLogOffTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) HasOffPeakLogOffTimeoutMinutes() bool`

HasOffPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakLogOffTimeoutMinutesNil

`func (o *CreateDeliveryGroupRequestModel) SetOffPeakLogOffTimeoutMinutesNil(b bool)`

 SetOffPeakLogOffTimeoutMinutesNil sets the value for OffPeakLogOffTimeoutMinutes to be an explicit nil

### UnsetOffPeakLogOffTimeoutMinutes
`func (o *CreateDeliveryGroupRequestModel) UnsetOffPeakLogOffTimeoutMinutes()`

UnsetOffPeakLogOffTimeoutMinutes ensures that no value is present for OffPeakLogOffTimeoutMinutes, not even an explicit nil
### GetLimitSecondsToForceLogOffUserDuringPeak

`func (o *CreateDeliveryGroupRequestModel) GetLimitSecondsToForceLogOffUserDuringPeak() int32`

GetLimitSecondsToForceLogOffUserDuringPeak returns the LimitSecondsToForceLogOffUserDuringPeak field if non-nil, zero value otherwise.

### GetLimitSecondsToForceLogOffUserDuringPeakOk

`func (o *CreateDeliveryGroupRequestModel) GetLimitSecondsToForceLogOffUserDuringPeakOk() (*int32, bool)`

GetLimitSecondsToForceLogOffUserDuringPeakOk returns a tuple with the LimitSecondsToForceLogOffUserDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSecondsToForceLogOffUserDuringPeak

`func (o *CreateDeliveryGroupRequestModel) SetLimitSecondsToForceLogOffUserDuringPeak(v int32)`

SetLimitSecondsToForceLogOffUserDuringPeak sets LimitSecondsToForceLogOffUserDuringPeak field to given value.

### HasLimitSecondsToForceLogOffUserDuringPeak

`func (o *CreateDeliveryGroupRequestModel) HasLimitSecondsToForceLogOffUserDuringPeak() bool`

HasLimitSecondsToForceLogOffUserDuringPeak returns a boolean if a field has been set.

### SetLimitSecondsToForceLogOffUserDuringPeakNil

`func (o *CreateDeliveryGroupRequestModel) SetLimitSecondsToForceLogOffUserDuringPeakNil(b bool)`

 SetLimitSecondsToForceLogOffUserDuringPeakNil sets the value for LimitSecondsToForceLogOffUserDuringPeak to be an explicit nil

### UnsetLimitSecondsToForceLogOffUserDuringPeak
`func (o *CreateDeliveryGroupRequestModel) UnsetLimitSecondsToForceLogOffUserDuringPeak()`

UnsetLimitSecondsToForceLogOffUserDuringPeak ensures that no value is present for LimitSecondsToForceLogOffUserDuringPeak, not even an explicit nil
### GetLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *CreateDeliveryGroupRequestModel) GetLimitSecondsToForceLogOffUserDuringOffPeak() int32`

GetLimitSecondsToForceLogOffUserDuringOffPeak returns the LimitSecondsToForceLogOffUserDuringOffPeak field if non-nil, zero value otherwise.

### GetLimitSecondsToForceLogOffUserDuringOffPeakOk

`func (o *CreateDeliveryGroupRequestModel) GetLimitSecondsToForceLogOffUserDuringOffPeakOk() (*int32, bool)`

GetLimitSecondsToForceLogOffUserDuringOffPeakOk returns a tuple with the LimitSecondsToForceLogOffUserDuringOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *CreateDeliveryGroupRequestModel) SetLimitSecondsToForceLogOffUserDuringOffPeak(v int32)`

SetLimitSecondsToForceLogOffUserDuringOffPeak sets LimitSecondsToForceLogOffUserDuringOffPeak field to given value.

### HasLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *CreateDeliveryGroupRequestModel) HasLimitSecondsToForceLogOffUserDuringOffPeak() bool`

HasLimitSecondsToForceLogOffUserDuringOffPeak returns a boolean if a field has been set.

### SetLimitSecondsToForceLogOffUserDuringOffPeakNil

`func (o *CreateDeliveryGroupRequestModel) SetLimitSecondsToForceLogOffUserDuringOffPeakNil(b bool)`

 SetLimitSecondsToForceLogOffUserDuringOffPeakNil sets the value for LimitSecondsToForceLogOffUserDuringOffPeak to be an explicit nil

### UnsetLimitSecondsToForceLogOffUserDuringOffPeak
`func (o *CreateDeliveryGroupRequestModel) UnsetLimitSecondsToForceLogOffUserDuringOffPeak()`

UnsetLimitSecondsToForceLogOffUserDuringOffPeak ensures that no value is present for LimitSecondsToForceLogOffUserDuringOffPeak, not even an explicit nil
### GetLogOffWarningMessage

`func (o *CreateDeliveryGroupRequestModel) GetLogOffWarningMessage() string`

GetLogOffWarningMessage returns the LogOffWarningMessage field if non-nil, zero value otherwise.

### GetLogOffWarningMessageOk

`func (o *CreateDeliveryGroupRequestModel) GetLogOffWarningMessageOk() (*string, bool)`

GetLogOffWarningMessageOk returns a tuple with the LogOffWarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOffWarningMessage

`func (o *CreateDeliveryGroupRequestModel) SetLogOffWarningMessage(v string)`

SetLogOffWarningMessage sets LogOffWarningMessage field to given value.

### HasLogOffWarningMessage

`func (o *CreateDeliveryGroupRequestModel) HasLogOffWarningMessage() bool`

HasLogOffWarningMessage returns a boolean if a field has been set.

### SetLogOffWarningMessageNil

`func (o *CreateDeliveryGroupRequestModel) SetLogOffWarningMessageNil(b bool)`

 SetLogOffWarningMessageNil sets the value for LogOffWarningMessage to be an explicit nil

### UnsetLogOffWarningMessage
`func (o *CreateDeliveryGroupRequestModel) UnsetLogOffWarningMessage()`

UnsetLogOffWarningMessage ensures that no value is present for LogOffWarningMessage, not even an explicit nil
### GetLogOffWarningTitle

`func (o *CreateDeliveryGroupRequestModel) GetLogOffWarningTitle() string`

GetLogOffWarningTitle returns the LogOffWarningTitle field if non-nil, zero value otherwise.

### GetLogOffWarningTitleOk

`func (o *CreateDeliveryGroupRequestModel) GetLogOffWarningTitleOk() (*string, bool)`

GetLogOffWarningTitleOk returns a tuple with the LogOffWarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOffWarningTitle

`func (o *CreateDeliveryGroupRequestModel) SetLogOffWarningTitle(v string)`

SetLogOffWarningTitle sets LogOffWarningTitle field to given value.

### HasLogOffWarningTitle

`func (o *CreateDeliveryGroupRequestModel) HasLogOffWarningTitle() bool`

HasLogOffWarningTitle returns a boolean if a field has been set.

### SetLogOffWarningTitleNil

`func (o *CreateDeliveryGroupRequestModel) SetLogOffWarningTitleNil(b bool)`

 SetLogOffWarningTitleNil sets the value for LogOffWarningTitle to be an explicit nil

### UnsetLogOffWarningTitle
`func (o *CreateDeliveryGroupRequestModel) UnsetLogOffWarningTitle()`

UnsetLogOffWarningTitle ensures that no value is present for LogOffWarningTitle, not even an explicit nil
### GetLoadBalanceType

`func (o *CreateDeliveryGroupRequestModel) GetLoadBalanceType() LoadBalanceType`

GetLoadBalanceType returns the LoadBalanceType field if non-nil, zero value otherwise.

### GetLoadBalanceTypeOk

`func (o *CreateDeliveryGroupRequestModel) GetLoadBalanceTypeOk() (*LoadBalanceType, bool)`

GetLoadBalanceTypeOk returns a tuple with the LoadBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceType

`func (o *CreateDeliveryGroupRequestModel) SetLoadBalanceType(v LoadBalanceType)`

SetLoadBalanceType sets LoadBalanceType field to given value.

### HasLoadBalanceType

`func (o *CreateDeliveryGroupRequestModel) HasLoadBalanceType() bool`

HasLoadBalanceType returns a boolean if a field has been set.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *CreateDeliveryGroupRequestModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak() int32`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak returns the RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field if non-nil, zero value otherwise.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk

`func (o *CreateDeliveryGroupRequestModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk() (*int32, bool)`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk returns a tuple with the RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *CreateDeliveryGroupRequestModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak(v int32)`

SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak sets RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field to given value.

### HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *CreateDeliveryGroupRequestModel) HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak() bool`

HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak returns a boolean if a field has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakNil

`func (o *CreateDeliveryGroupRequestModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakNil(b bool)`

 SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakNil sets the value for RestrictAutoscaleMinIdleUntaggedPercentDuringPeak to be an explicit nil

### UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak
`func (o *CreateDeliveryGroupRequestModel) UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak()`

UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak ensures that no value is present for RestrictAutoscaleMinIdleUntaggedPercentDuringPeak, not even an explicit nil
### GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *CreateDeliveryGroupRequestModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak() int32`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak returns the RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field if non-nil, zero value otherwise.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk

`func (o *CreateDeliveryGroupRequestModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk() (*int32, bool)`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk returns a tuple with the RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *CreateDeliveryGroupRequestModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak(v int32)`

SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak sets RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field to given value.

### HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *CreateDeliveryGroupRequestModel) HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak() bool`

HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak returns a boolean if a field has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakNil

`func (o *CreateDeliveryGroupRequestModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakNil(b bool)`

 SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakNil sets the value for RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak to be an explicit nil

### UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak
`func (o *CreateDeliveryGroupRequestModel) UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak()`

UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak ensures that no value is present for RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak, not even an explicit nil
### GetAutoscaleLogOffReminderEnabled

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderEnabled() bool`

GetAutoscaleLogOffReminderEnabled returns the AutoscaleLogOffReminderEnabled field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderEnabledOk

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderEnabledOk() (*bool, bool)`

GetAutoscaleLogOffReminderEnabledOk returns a tuple with the AutoscaleLogOffReminderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderEnabled

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderEnabled(v bool)`

SetAutoscaleLogOffReminderEnabled sets AutoscaleLogOffReminderEnabled field to given value.

### HasAutoscaleLogOffReminderEnabled

`func (o *CreateDeliveryGroupRequestModel) HasAutoscaleLogOffReminderEnabled() bool`

HasAutoscaleLogOffReminderEnabled returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderEnabledNil

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderEnabledNil(b bool)`

 SetAutoscaleLogOffReminderEnabledNil sets the value for AutoscaleLogOffReminderEnabled to be an explicit nil

### UnsetAutoscaleLogOffReminderEnabled
`func (o *CreateDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderEnabled()`

UnsetAutoscaleLogOffReminderEnabled ensures that no value is present for AutoscaleLogOffReminderEnabled, not even an explicit nil
### GetAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderIntervalSecondsOffPeak() int32`

GetAutoscaleLogOffReminderIntervalSecondsOffPeak returns the AutoscaleLogOffReminderIntervalSecondsOffPeak field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk() (*int32, bool)`

GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk returns a tuple with the AutoscaleLogOffReminderIntervalSecondsOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderIntervalSecondsOffPeak(v int32)`

SetAutoscaleLogOffReminderIntervalSecondsOffPeak sets AutoscaleLogOffReminderIntervalSecondsOffPeak field to given value.

### HasAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *CreateDeliveryGroupRequestModel) HasAutoscaleLogOffReminderIntervalSecondsOffPeak() bool`

HasAutoscaleLogOffReminderIntervalSecondsOffPeak returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderIntervalSecondsOffPeakNil

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderIntervalSecondsOffPeakNil(b bool)`

 SetAutoscaleLogOffReminderIntervalSecondsOffPeakNil sets the value for AutoscaleLogOffReminderIntervalSecondsOffPeak to be an explicit nil

### UnsetAutoscaleLogOffReminderIntervalSecondsOffPeak
`func (o *CreateDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderIntervalSecondsOffPeak()`

UnsetAutoscaleLogOffReminderIntervalSecondsOffPeak ensures that no value is present for AutoscaleLogOffReminderIntervalSecondsOffPeak, not even an explicit nil
### GetAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderIntervalSecondsPeak() int32`

GetAutoscaleLogOffReminderIntervalSecondsPeak returns the AutoscaleLogOffReminderIntervalSecondsPeak field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderIntervalSecondsPeakOk

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderIntervalSecondsPeakOk() (*int32, bool)`

GetAutoscaleLogOffReminderIntervalSecondsPeakOk returns a tuple with the AutoscaleLogOffReminderIntervalSecondsPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderIntervalSecondsPeak(v int32)`

SetAutoscaleLogOffReminderIntervalSecondsPeak sets AutoscaleLogOffReminderIntervalSecondsPeak field to given value.

### HasAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *CreateDeliveryGroupRequestModel) HasAutoscaleLogOffReminderIntervalSecondsPeak() bool`

HasAutoscaleLogOffReminderIntervalSecondsPeak returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderIntervalSecondsPeakNil

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderIntervalSecondsPeakNil(b bool)`

 SetAutoscaleLogOffReminderIntervalSecondsPeakNil sets the value for AutoscaleLogOffReminderIntervalSecondsPeak to be an explicit nil

### UnsetAutoscaleLogOffReminderIntervalSecondsPeak
`func (o *CreateDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderIntervalSecondsPeak()`

UnsetAutoscaleLogOffReminderIntervalSecondsPeak ensures that no value is present for AutoscaleLogOffReminderIntervalSecondsPeak, not even an explicit nil
### GetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) GetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes() int32`

GetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes returns the PeakAutoscaleAssignedPowerOnIdleTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutesOk

`func (o *CreateDeliveryGroupRequestModel) GetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutesOk() (*int32, bool)`

GetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutesOk returns a tuple with the PeakAutoscaleAssignedPowerOnIdleTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) SetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes(v int32)`

SetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes sets PeakAutoscaleAssignedPowerOnIdleTimeoutMinutes field to given value.

### HasPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) HasPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes() bool`

HasPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes returns a boolean if a field has been set.

### SetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutesNil

`func (o *CreateDeliveryGroupRequestModel) SetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutesNil(b bool)`

 SetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutesNil sets the value for PeakAutoscaleAssignedPowerOnIdleTimeoutMinutes to be an explicit nil

### UnsetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes
`func (o *CreateDeliveryGroupRequestModel) UnsetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes()`

UnsetPeakAutoscaleAssignedPowerOnIdleTimeoutMinutes ensures that no value is present for PeakAutoscaleAssignedPowerOnIdleTimeoutMinutes, not even an explicit nil
### GetPeakAutoscaleAssignedPowerOnIdleAction

`func (o *CreateDeliveryGroupRequestModel) GetPeakAutoscaleAssignedPowerOnIdleAction() SessionChangeHostingAction`

GetPeakAutoscaleAssignedPowerOnIdleAction returns the PeakAutoscaleAssignedPowerOnIdleAction field if non-nil, zero value otherwise.

### GetPeakAutoscaleAssignedPowerOnIdleActionOk

`func (o *CreateDeliveryGroupRequestModel) GetPeakAutoscaleAssignedPowerOnIdleActionOk() (*SessionChangeHostingAction, bool)`

GetPeakAutoscaleAssignedPowerOnIdleActionOk returns a tuple with the PeakAutoscaleAssignedPowerOnIdleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakAutoscaleAssignedPowerOnIdleAction

`func (o *CreateDeliveryGroupRequestModel) SetPeakAutoscaleAssignedPowerOnIdleAction(v SessionChangeHostingAction)`

SetPeakAutoscaleAssignedPowerOnIdleAction sets PeakAutoscaleAssignedPowerOnIdleAction field to given value.

### HasPeakAutoscaleAssignedPowerOnIdleAction

`func (o *CreateDeliveryGroupRequestModel) HasPeakAutoscaleAssignedPowerOnIdleAction() bool`

HasPeakAutoscaleAssignedPowerOnIdleAction returns a boolean if a field has been set.

### GetAutoscaleLogOffReminderMessage

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderMessage() string`

GetAutoscaleLogOffReminderMessage returns the AutoscaleLogOffReminderMessage field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderMessageOk

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderMessageOk() (*string, bool)`

GetAutoscaleLogOffReminderMessageOk returns a tuple with the AutoscaleLogOffReminderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderMessage

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderMessage(v string)`

SetAutoscaleLogOffReminderMessage sets AutoscaleLogOffReminderMessage field to given value.

### HasAutoscaleLogOffReminderMessage

`func (o *CreateDeliveryGroupRequestModel) HasAutoscaleLogOffReminderMessage() bool`

HasAutoscaleLogOffReminderMessage returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderMessageNil

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderMessageNil(b bool)`

 SetAutoscaleLogOffReminderMessageNil sets the value for AutoscaleLogOffReminderMessage to be an explicit nil

### UnsetAutoscaleLogOffReminderMessage
`func (o *CreateDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderMessage()`

UnsetAutoscaleLogOffReminderMessage ensures that no value is present for AutoscaleLogOffReminderMessage, not even an explicit nil
### GetAutoscaleLogOffReminderTitle

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderTitle() string`

GetAutoscaleLogOffReminderTitle returns the AutoscaleLogOffReminderTitle field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderTitleOk

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleLogOffReminderTitleOk() (*string, bool)`

GetAutoscaleLogOffReminderTitleOk returns a tuple with the AutoscaleLogOffReminderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderTitle

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderTitle(v string)`

SetAutoscaleLogOffReminderTitle sets AutoscaleLogOffReminderTitle field to given value.

### HasAutoscaleLogOffReminderTitle

`func (o *CreateDeliveryGroupRequestModel) HasAutoscaleLogOffReminderTitle() bool`

HasAutoscaleLogOffReminderTitle returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderTitleNil

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleLogOffReminderTitleNil(b bool)`

 SetAutoscaleLogOffReminderTitleNil sets the value for AutoscaleLogOffReminderTitle to be an explicit nil

### UnsetAutoscaleLogOffReminderTitle
`func (o *CreateDeliveryGroupRequestModel) UnsetAutoscaleLogOffReminderTitle()`

UnsetAutoscaleLogOffReminderTitle ensures that no value is present for AutoscaleLogOffReminderTitle, not even an explicit nil
### GetPeakBufferSizePercent

`func (o *CreateDeliveryGroupRequestModel) GetPeakBufferSizePercent() int32`

GetPeakBufferSizePercent returns the PeakBufferSizePercent field if non-nil, zero value otherwise.

### GetPeakBufferSizePercentOk

`func (o *CreateDeliveryGroupRequestModel) GetPeakBufferSizePercentOk() (*int32, bool)`

GetPeakBufferSizePercentOk returns a tuple with the PeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBufferSizePercent

`func (o *CreateDeliveryGroupRequestModel) SetPeakBufferSizePercent(v int32)`

SetPeakBufferSizePercent sets PeakBufferSizePercent field to given value.

### HasPeakBufferSizePercent

`func (o *CreateDeliveryGroupRequestModel) HasPeakBufferSizePercent() bool`

HasPeakBufferSizePercent returns a boolean if a field has been set.

### SetPeakBufferSizePercentNil

`func (o *CreateDeliveryGroupRequestModel) SetPeakBufferSizePercentNil(b bool)`

 SetPeakBufferSizePercentNil sets the value for PeakBufferSizePercent to be an explicit nil

### UnsetPeakBufferSizePercent
`func (o *CreateDeliveryGroupRequestModel) UnsetPeakBufferSizePercent()`

UnsetPeakBufferSizePercent ensures that no value is present for PeakBufferSizePercent, not even an explicit nil
### GetPeakDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) GetPeakDisconnectAction() SessionChangeHostingAction`

GetPeakDisconnectAction returns the PeakDisconnectAction field if non-nil, zero value otherwise.

### GetPeakDisconnectActionOk

`func (o *CreateDeliveryGroupRequestModel) GetPeakDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetPeakDisconnectActionOk returns a tuple with the PeakDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) SetPeakDisconnectAction(v SessionChangeHostingAction)`

SetPeakDisconnectAction sets PeakDisconnectAction field to given value.

### HasPeakDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) HasPeakDisconnectAction() bool`

HasPeakDisconnectAction returns a boolean if a field has been set.

### GetPeakDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) GetPeakDisconnectTimeoutMinutes() int32`

GetPeakDisconnectTimeoutMinutes returns the PeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakDisconnectTimeoutMinutesOk

`func (o *CreateDeliveryGroupRequestModel) GetPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakDisconnectTimeoutMinutesOk returns a tuple with the PeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) SetPeakDisconnectTimeoutMinutes(v int32)`

SetPeakDisconnectTimeoutMinutes sets PeakDisconnectTimeoutMinutes field to given value.

### HasPeakDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) HasPeakDisconnectTimeoutMinutes() bool`

HasPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetPeakDisconnectTimeoutMinutesNil

`func (o *CreateDeliveryGroupRequestModel) SetPeakDisconnectTimeoutMinutesNil(b bool)`

 SetPeakDisconnectTimeoutMinutesNil sets the value for PeakDisconnectTimeoutMinutes to be an explicit nil

### UnsetPeakDisconnectTimeoutMinutes
`func (o *CreateDeliveryGroupRequestModel) UnsetPeakDisconnectTimeoutMinutes()`

UnsetPeakDisconnectTimeoutMinutes ensures that no value is present for PeakDisconnectTimeoutMinutes, not even an explicit nil
### GetPeakExtendedDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) GetPeakExtendedDisconnectAction() SessionChangeHostingAction`

GetPeakExtendedDisconnectAction returns the PeakExtendedDisconnectAction field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectActionOk

`func (o *CreateDeliveryGroupRequestModel) GetPeakExtendedDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetPeakExtendedDisconnectActionOk returns a tuple with the PeakExtendedDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) SetPeakExtendedDisconnectAction(v SessionChangeHostingAction)`

SetPeakExtendedDisconnectAction sets PeakExtendedDisconnectAction field to given value.

### HasPeakExtendedDisconnectAction

`func (o *CreateDeliveryGroupRequestModel) HasPeakExtendedDisconnectAction() bool`

HasPeakExtendedDisconnectAction returns a boolean if a field has been set.

### GetPeakExtendedDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) GetPeakExtendedDisconnectTimeoutMinutes() int32`

GetPeakExtendedDisconnectTimeoutMinutes returns the PeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectTimeoutMinutesOk

`func (o *CreateDeliveryGroupRequestModel) GetPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the PeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) SetPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetPeakExtendedDisconnectTimeoutMinutes sets PeakExtendedDisconnectTimeoutMinutes field to given value.

### HasPeakExtendedDisconnectTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) HasPeakExtendedDisconnectTimeoutMinutes() bool`

HasPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetPeakExtendedDisconnectTimeoutMinutesNil

`func (o *CreateDeliveryGroupRequestModel) SetPeakExtendedDisconnectTimeoutMinutesNil(b bool)`

 SetPeakExtendedDisconnectTimeoutMinutesNil sets the value for PeakExtendedDisconnectTimeoutMinutes to be an explicit nil

### UnsetPeakExtendedDisconnectTimeoutMinutes
`func (o *CreateDeliveryGroupRequestModel) UnsetPeakExtendedDisconnectTimeoutMinutes()`

UnsetPeakExtendedDisconnectTimeoutMinutes ensures that no value is present for PeakExtendedDisconnectTimeoutMinutes, not even an explicit nil
### GetPeakLogOffAction

`func (o *CreateDeliveryGroupRequestModel) GetPeakLogOffAction() SessionChangeHostingAction`

GetPeakLogOffAction returns the PeakLogOffAction field if non-nil, zero value otherwise.

### GetPeakLogOffActionOk

`func (o *CreateDeliveryGroupRequestModel) GetPeakLogOffActionOk() (*SessionChangeHostingAction, bool)`

GetPeakLogOffActionOk returns a tuple with the PeakLogOffAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLogOffAction

`func (o *CreateDeliveryGroupRequestModel) SetPeakLogOffAction(v SessionChangeHostingAction)`

SetPeakLogOffAction sets PeakLogOffAction field to given value.

### HasPeakLogOffAction

`func (o *CreateDeliveryGroupRequestModel) HasPeakLogOffAction() bool`

HasPeakLogOffAction returns a boolean if a field has been set.

### GetPeakLogOffTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) GetPeakLogOffTimeoutMinutes() int32`

GetPeakLogOffTimeoutMinutes returns the PeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakLogOffTimeoutMinutesOk

`func (o *CreateDeliveryGroupRequestModel) GetPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetPeakLogOffTimeoutMinutesOk returns a tuple with the PeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLogOffTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) SetPeakLogOffTimeoutMinutes(v int32)`

SetPeakLogOffTimeoutMinutes sets PeakLogOffTimeoutMinutes field to given value.

### HasPeakLogOffTimeoutMinutes

`func (o *CreateDeliveryGroupRequestModel) HasPeakLogOffTimeoutMinutes() bool`

HasPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### SetPeakLogOffTimeoutMinutesNil

`func (o *CreateDeliveryGroupRequestModel) SetPeakLogOffTimeoutMinutesNil(b bool)`

 SetPeakLogOffTimeoutMinutesNil sets the value for PeakLogOffTimeoutMinutes to be an explicit nil

### UnsetPeakLogOffTimeoutMinutes
`func (o *CreateDeliveryGroupRequestModel) UnsetPeakLogOffTimeoutMinutes()`

UnsetPeakLogOffTimeoutMinutes ensures that no value is present for PeakLogOffTimeoutMinutes, not even an explicit nil
### GetDisconnectPeakIdleSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) GetDisconnectPeakIdleSessionAfterSeconds() int32`

GetDisconnectPeakIdleSessionAfterSeconds returns the DisconnectPeakIdleSessionAfterSeconds field if non-nil, zero value otherwise.

### GetDisconnectPeakIdleSessionAfterSecondsOk

`func (o *CreateDeliveryGroupRequestModel) GetDisconnectPeakIdleSessionAfterSecondsOk() (*int32, bool)`

GetDisconnectPeakIdleSessionAfterSecondsOk returns a tuple with the DisconnectPeakIdleSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectPeakIdleSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) SetDisconnectPeakIdleSessionAfterSeconds(v int32)`

SetDisconnectPeakIdleSessionAfterSeconds sets DisconnectPeakIdleSessionAfterSeconds field to given value.

### HasDisconnectPeakIdleSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) HasDisconnectPeakIdleSessionAfterSeconds() bool`

HasDisconnectPeakIdleSessionAfterSeconds returns a boolean if a field has been set.

### SetDisconnectPeakIdleSessionAfterSecondsNil

`func (o *CreateDeliveryGroupRequestModel) SetDisconnectPeakIdleSessionAfterSecondsNil(b bool)`

 SetDisconnectPeakIdleSessionAfterSecondsNil sets the value for DisconnectPeakIdleSessionAfterSeconds to be an explicit nil

### UnsetDisconnectPeakIdleSessionAfterSeconds
`func (o *CreateDeliveryGroupRequestModel) UnsetDisconnectPeakIdleSessionAfterSeconds()`

UnsetDisconnectPeakIdleSessionAfterSeconds ensures that no value is present for DisconnectPeakIdleSessionAfterSeconds, not even an explicit nil
### GetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) GetDisconnectOffPeakIdleSessionAfterSeconds() int32`

GetDisconnectOffPeakIdleSessionAfterSeconds returns the DisconnectOffPeakIdleSessionAfterSeconds field if non-nil, zero value otherwise.

### GetDisconnectOffPeakIdleSessionAfterSecondsOk

`func (o *CreateDeliveryGroupRequestModel) GetDisconnectOffPeakIdleSessionAfterSecondsOk() (*int32, bool)`

GetDisconnectOffPeakIdleSessionAfterSecondsOk returns a tuple with the DisconnectOffPeakIdleSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) SetDisconnectOffPeakIdleSessionAfterSeconds(v int32)`

SetDisconnectOffPeakIdleSessionAfterSeconds sets DisconnectOffPeakIdleSessionAfterSeconds field to given value.

### HasDisconnectOffPeakIdleSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) HasDisconnectOffPeakIdleSessionAfterSeconds() bool`

HasDisconnectOffPeakIdleSessionAfterSeconds returns a boolean if a field has been set.

### SetDisconnectOffPeakIdleSessionAfterSecondsNil

`func (o *CreateDeliveryGroupRequestModel) SetDisconnectOffPeakIdleSessionAfterSecondsNil(b bool)`

 SetDisconnectOffPeakIdleSessionAfterSecondsNil sets the value for DisconnectOffPeakIdleSessionAfterSeconds to be an explicit nil

### UnsetDisconnectOffPeakIdleSessionAfterSeconds
`func (o *CreateDeliveryGroupRequestModel) UnsetDisconnectOffPeakIdleSessionAfterSeconds()`

UnsetDisconnectOffPeakIdleSessionAfterSeconds ensures that no value is present for DisconnectOffPeakIdleSessionAfterSeconds, not even an explicit nil
### GetLogoffPeakDisconnectedSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) GetLogoffPeakDisconnectedSessionAfterSeconds() int32`

GetLogoffPeakDisconnectedSessionAfterSeconds returns the LogoffPeakDisconnectedSessionAfterSeconds field if non-nil, zero value otherwise.

### GetLogoffPeakDisconnectedSessionAfterSecondsOk

`func (o *CreateDeliveryGroupRequestModel) GetLogoffPeakDisconnectedSessionAfterSecondsOk() (*int32, bool)`

GetLogoffPeakDisconnectedSessionAfterSecondsOk returns a tuple with the LogoffPeakDisconnectedSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPeakDisconnectedSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) SetLogoffPeakDisconnectedSessionAfterSeconds(v int32)`

SetLogoffPeakDisconnectedSessionAfterSeconds sets LogoffPeakDisconnectedSessionAfterSeconds field to given value.

### HasLogoffPeakDisconnectedSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) HasLogoffPeakDisconnectedSessionAfterSeconds() bool`

HasLogoffPeakDisconnectedSessionAfterSeconds returns a boolean if a field has been set.

### SetLogoffPeakDisconnectedSessionAfterSecondsNil

`func (o *CreateDeliveryGroupRequestModel) SetLogoffPeakDisconnectedSessionAfterSecondsNil(b bool)`

 SetLogoffPeakDisconnectedSessionAfterSecondsNil sets the value for LogoffPeakDisconnectedSessionAfterSeconds to be an explicit nil

### UnsetLogoffPeakDisconnectedSessionAfterSeconds
`func (o *CreateDeliveryGroupRequestModel) UnsetLogoffPeakDisconnectedSessionAfterSeconds()`

UnsetLogoffPeakDisconnectedSessionAfterSeconds ensures that no value is present for LogoffPeakDisconnectedSessionAfterSeconds, not even an explicit nil
### GetLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) GetLogoffOffPeakDisconnectedSessionAfterSeconds() int32`

GetLogoffOffPeakDisconnectedSessionAfterSeconds returns the LogoffOffPeakDisconnectedSessionAfterSeconds field if non-nil, zero value otherwise.

### GetLogoffOffPeakDisconnectedSessionAfterSecondsOk

`func (o *CreateDeliveryGroupRequestModel) GetLogoffOffPeakDisconnectedSessionAfterSecondsOk() (*int32, bool)`

GetLogoffOffPeakDisconnectedSessionAfterSecondsOk returns a tuple with the LogoffOffPeakDisconnectedSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) SetLogoffOffPeakDisconnectedSessionAfterSeconds(v int32)`

SetLogoffOffPeakDisconnectedSessionAfterSeconds sets LogoffOffPeakDisconnectedSessionAfterSeconds field to given value.

### HasLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *CreateDeliveryGroupRequestModel) HasLogoffOffPeakDisconnectedSessionAfterSeconds() bool`

HasLogoffOffPeakDisconnectedSessionAfterSeconds returns a boolean if a field has been set.

### SetLogoffOffPeakDisconnectedSessionAfterSecondsNil

`func (o *CreateDeliveryGroupRequestModel) SetLogoffOffPeakDisconnectedSessionAfterSecondsNil(b bool)`

 SetLogoffOffPeakDisconnectedSessionAfterSecondsNil sets the value for LogoffOffPeakDisconnectedSessionAfterSeconds to be an explicit nil

### UnsetLogoffOffPeakDisconnectedSessionAfterSeconds
`func (o *CreateDeliveryGroupRequestModel) UnsetLogoffOffPeakDisconnectedSessionAfterSeconds()`

UnsetLogoffOffPeakDisconnectedSessionAfterSeconds ensures that no value is present for LogoffOffPeakDisconnectedSessionAfterSeconds, not even an explicit nil
### GetAutoscaleScaleDownActionDuringPeak

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleScaleDownActionDuringPeak() AutoscaleScaleDownAction`

GetAutoscaleScaleDownActionDuringPeak returns the AutoscaleScaleDownActionDuringPeak field if non-nil, zero value otherwise.

### GetAutoscaleScaleDownActionDuringPeakOk

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleScaleDownActionDuringPeakOk() (*AutoscaleScaleDownAction, bool)`

GetAutoscaleScaleDownActionDuringPeakOk returns a tuple with the AutoscaleScaleDownActionDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleScaleDownActionDuringPeak

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleScaleDownActionDuringPeak(v AutoscaleScaleDownAction)`

SetAutoscaleScaleDownActionDuringPeak sets AutoscaleScaleDownActionDuringPeak field to given value.

### HasAutoscaleScaleDownActionDuringPeak

`func (o *CreateDeliveryGroupRequestModel) HasAutoscaleScaleDownActionDuringPeak() bool`

HasAutoscaleScaleDownActionDuringPeak returns a boolean if a field has been set.

### GetAutoscaleScaleDownActionDuringOffPeak

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleScaleDownActionDuringOffPeak() AutoscaleScaleDownAction`

GetAutoscaleScaleDownActionDuringOffPeak returns the AutoscaleScaleDownActionDuringOffPeak field if non-nil, zero value otherwise.

### GetAutoscaleScaleDownActionDuringOffPeakOk

`func (o *CreateDeliveryGroupRequestModel) GetAutoscaleScaleDownActionDuringOffPeakOk() (*AutoscaleScaleDownAction, bool)`

GetAutoscaleScaleDownActionDuringOffPeakOk returns a tuple with the AutoscaleScaleDownActionDuringOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleScaleDownActionDuringOffPeak

`func (o *CreateDeliveryGroupRequestModel) SetAutoscaleScaleDownActionDuringOffPeak(v AutoscaleScaleDownAction)`

SetAutoscaleScaleDownActionDuringOffPeak sets AutoscaleScaleDownActionDuringOffPeak field to given value.

### HasAutoscaleScaleDownActionDuringOffPeak

`func (o *CreateDeliveryGroupRequestModel) HasAutoscaleScaleDownActionDuringOffPeak() bool`

HasAutoscaleScaleDownActionDuringOffPeak returns a boolean if a field has been set.

### GetPrelaunchSettings

`func (o *CreateDeliveryGroupRequestModel) GetPrelaunchSettings() FastApplicationSettingsRequestModel`

GetPrelaunchSettings returns the PrelaunchSettings field if non-nil, zero value otherwise.

### GetPrelaunchSettingsOk

`func (o *CreateDeliveryGroupRequestModel) GetPrelaunchSettingsOk() (*FastApplicationSettingsRequestModel, bool)`

GetPrelaunchSettingsOk returns a tuple with the PrelaunchSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrelaunchSettings

`func (o *CreateDeliveryGroupRequestModel) SetPrelaunchSettings(v FastApplicationSettingsRequestModel)`

SetPrelaunchSettings sets PrelaunchSettings field to given value.

### HasPrelaunchSettings

`func (o *CreateDeliveryGroupRequestModel) HasPrelaunchSettings() bool`

HasPrelaunchSettings returns a boolean if a field has been set.

### GetPowerTimeSchemes

`func (o *CreateDeliveryGroupRequestModel) GetPowerTimeSchemes() []PowerTimeSchemeRequestModel`

GetPowerTimeSchemes returns the PowerTimeSchemes field if non-nil, zero value otherwise.

### GetPowerTimeSchemesOk

`func (o *CreateDeliveryGroupRequestModel) GetPowerTimeSchemesOk() (*[]PowerTimeSchemeRequestModel, bool)`

GetPowerTimeSchemesOk returns a tuple with the PowerTimeSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerTimeSchemes

`func (o *CreateDeliveryGroupRequestModel) SetPowerTimeSchemes(v []PowerTimeSchemeRequestModel)`

SetPowerTimeSchemes sets PowerTimeSchemes field to given value.

### HasPowerTimeSchemes

`func (o *CreateDeliveryGroupRequestModel) HasPowerTimeSchemes() bool`

HasPowerTimeSchemes returns a boolean if a field has been set.

### SetPowerTimeSchemesNil

`func (o *CreateDeliveryGroupRequestModel) SetPowerTimeSchemesNil(b bool)`

 SetPowerTimeSchemesNil sets the value for PowerTimeSchemes to be an explicit nil

### UnsetPowerTimeSchemes
`func (o *CreateDeliveryGroupRequestModel) UnsetPowerTimeSchemes()`

UnsetPowerTimeSchemes ensures that no value is present for PowerTimeSchemes, not even an explicit nil
### GetPowerOffDelayMinutes

`func (o *CreateDeliveryGroupRequestModel) GetPowerOffDelayMinutes() int32`

GetPowerOffDelayMinutes returns the PowerOffDelayMinutes field if non-nil, zero value otherwise.

### GetPowerOffDelayMinutesOk

`func (o *CreateDeliveryGroupRequestModel) GetPowerOffDelayMinutesOk() (*int32, bool)`

GetPowerOffDelayMinutesOk returns a tuple with the PowerOffDelayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffDelayMinutes

`func (o *CreateDeliveryGroupRequestModel) SetPowerOffDelayMinutes(v int32)`

SetPowerOffDelayMinutes sets PowerOffDelayMinutes field to given value.

### HasPowerOffDelayMinutes

`func (o *CreateDeliveryGroupRequestModel) HasPowerOffDelayMinutes() bool`

HasPowerOffDelayMinutes returns a boolean if a field has been set.

### SetPowerOffDelayMinutesNil

`func (o *CreateDeliveryGroupRequestModel) SetPowerOffDelayMinutesNil(b bool)`

 SetPowerOffDelayMinutesNil sets the value for PowerOffDelayMinutes to be an explicit nil

### UnsetPowerOffDelayMinutes
`func (o *CreateDeliveryGroupRequestModel) UnsetPowerOffDelayMinutes()`

UnsetPowerOffDelayMinutes ensures that no value is present for PowerOffDelayMinutes, not even an explicit nil
### GetMachineCost

`func (o *CreateDeliveryGroupRequestModel) GetMachineCost() float64`

GetMachineCost returns the MachineCost field if non-nil, zero value otherwise.

### GetMachineCostOk

`func (o *CreateDeliveryGroupRequestModel) GetMachineCostOk() (*float64, bool)`

GetMachineCostOk returns a tuple with the MachineCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCost

`func (o *CreateDeliveryGroupRequestModel) SetMachineCost(v float64)`

SetMachineCost sets MachineCost field to given value.

### HasMachineCost

`func (o *CreateDeliveryGroupRequestModel) HasMachineCost() bool`

HasMachineCost returns a boolean if a field has been set.

### SetMachineCostNil

`func (o *CreateDeliveryGroupRequestModel) SetMachineCostNil(b bool)`

 SetMachineCostNil sets the value for MachineCost to be an explicit nil

### UnsetMachineCost
`func (o *CreateDeliveryGroupRequestModel) UnsetMachineCost()`

UnsetMachineCost ensures that no value is present for MachineCost, not even an explicit nil
### GetMachineLogOnType

`func (o *CreateDeliveryGroupRequestModel) GetMachineLogOnType() MachineLogOnType`

GetMachineLogOnType returns the MachineLogOnType field if non-nil, zero value otherwise.

### GetMachineLogOnTypeOk

`func (o *CreateDeliveryGroupRequestModel) GetMachineLogOnTypeOk() (*MachineLogOnType, bool)`

GetMachineLogOnTypeOk returns a tuple with the MachineLogOnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLogOnType

`func (o *CreateDeliveryGroupRequestModel) SetMachineLogOnType(v MachineLogOnType)`

SetMachineLogOnType sets MachineLogOnType field to given value.

### HasMachineLogOnType

`func (o *CreateDeliveryGroupRequestModel) HasMachineLogOnType() bool`

HasMachineLogOnType returns a boolean if a field has been set.

### GetProductCode

`func (o *CreateDeliveryGroupRequestModel) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *CreateDeliveryGroupRequestModel) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *CreateDeliveryGroupRequestModel) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *CreateDeliveryGroupRequestModel) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCodeNil

`func (o *CreateDeliveryGroupRequestModel) SetProductCodeNil(b bool)`

 SetProductCodeNil sets the value for ProductCode to be an explicit nil

### UnsetProductCode
`func (o *CreateDeliveryGroupRequestModel) UnsetProductCode()`

UnsetProductCode ensures that no value is present for ProductCode, not even an explicit nil
### GetReuseMachinesWithoutShutdownInOutage

`func (o *CreateDeliveryGroupRequestModel) GetReuseMachinesWithoutShutdownInOutage() bool`

GetReuseMachinesWithoutShutdownInOutage returns the ReuseMachinesWithoutShutdownInOutage field if non-nil, zero value otherwise.

### GetReuseMachinesWithoutShutdownInOutageOk

`func (o *CreateDeliveryGroupRequestModel) GetReuseMachinesWithoutShutdownInOutageOk() (*bool, bool)`

GetReuseMachinesWithoutShutdownInOutageOk returns a tuple with the ReuseMachinesWithoutShutdownInOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseMachinesWithoutShutdownInOutage

`func (o *CreateDeliveryGroupRequestModel) SetReuseMachinesWithoutShutdownInOutage(v bool)`

SetReuseMachinesWithoutShutdownInOutage sets ReuseMachinesWithoutShutdownInOutage field to given value.

### HasReuseMachinesWithoutShutdownInOutage

`func (o *CreateDeliveryGroupRequestModel) HasReuseMachinesWithoutShutdownInOutage() bool`

HasReuseMachinesWithoutShutdownInOutage returns a boolean if a field has been set.

### SetReuseMachinesWithoutShutdownInOutageNil

`func (o *CreateDeliveryGroupRequestModel) SetReuseMachinesWithoutShutdownInOutageNil(b bool)`

 SetReuseMachinesWithoutShutdownInOutageNil sets the value for ReuseMachinesWithoutShutdownInOutage to be an explicit nil

### UnsetReuseMachinesWithoutShutdownInOutage
`func (o *CreateDeliveryGroupRequestModel) UnsetReuseMachinesWithoutShutdownInOutage()`

UnsetReuseMachinesWithoutShutdownInOutage ensures that no value is present for ReuseMachinesWithoutShutdownInOutage, not even an explicit nil
### GetLicenseModel

`func (o *CreateDeliveryGroupRequestModel) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *CreateDeliveryGroupRequestModel) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *CreateDeliveryGroupRequestModel) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *CreateDeliveryGroupRequestModel) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetProtocolPriority

`func (o *CreateDeliveryGroupRequestModel) GetProtocolPriority() []ProtocolType`

GetProtocolPriority returns the ProtocolPriority field if non-nil, zero value otherwise.

### GetProtocolPriorityOk

`func (o *CreateDeliveryGroupRequestModel) GetProtocolPriorityOk() (*[]ProtocolType, bool)`

GetProtocolPriorityOk returns a tuple with the ProtocolPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPriority

`func (o *CreateDeliveryGroupRequestModel) SetProtocolPriority(v []ProtocolType)`

SetProtocolPriority sets ProtocolPriority field to given value.

### HasProtocolPriority

`func (o *CreateDeliveryGroupRequestModel) HasProtocolPriority() bool`

HasProtocolPriority returns a boolean if a field has been set.

### SetProtocolPriorityNil

`func (o *CreateDeliveryGroupRequestModel) SetProtocolPriorityNil(b bool)`

 SetProtocolPriorityNil sets the value for ProtocolPriority to be an explicit nil

### UnsetProtocolPriority
`func (o *CreateDeliveryGroupRequestModel) UnsetProtocolPriority()`

UnsetProtocolPriority ensures that no value is present for ProtocolPriority, not even an explicit nil
### GetRebootSchedules

`func (o *CreateDeliveryGroupRequestModel) GetRebootSchedules() []RebootScheduleRequestModel`

GetRebootSchedules returns the RebootSchedules field if non-nil, zero value otherwise.

### GetRebootSchedulesOk

`func (o *CreateDeliveryGroupRequestModel) GetRebootSchedulesOk() (*[]RebootScheduleRequestModel, bool)`

GetRebootSchedulesOk returns a tuple with the RebootSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootSchedules

`func (o *CreateDeliveryGroupRequestModel) SetRebootSchedules(v []RebootScheduleRequestModel)`

SetRebootSchedules sets RebootSchedules field to given value.

### HasRebootSchedules

`func (o *CreateDeliveryGroupRequestModel) HasRebootSchedules() bool`

HasRebootSchedules returns a boolean if a field has been set.

### SetRebootSchedulesNil

`func (o *CreateDeliveryGroupRequestModel) SetRebootSchedulesNil(b bool)`

 SetRebootSchedulesNil sets the value for RebootSchedules to be an explicit nil

### UnsetRebootSchedules
`func (o *CreateDeliveryGroupRequestModel) UnsetRebootSchedules()`

UnsetRebootSchedules ensures that no value is present for RebootSchedules, not even an explicit nil
### GetSecureIcaRequired

`func (o *CreateDeliveryGroupRequestModel) GetSecureIcaRequired() bool`

GetSecureIcaRequired returns the SecureIcaRequired field if non-nil, zero value otherwise.

### GetSecureIcaRequiredOk

`func (o *CreateDeliveryGroupRequestModel) GetSecureIcaRequiredOk() (*bool, bool)`

GetSecureIcaRequiredOk returns a tuple with the SecureIcaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureIcaRequired

`func (o *CreateDeliveryGroupRequestModel) SetSecureIcaRequired(v bool)`

SetSecureIcaRequired sets SecureIcaRequired field to given value.

### HasSecureIcaRequired

`func (o *CreateDeliveryGroupRequestModel) HasSecureIcaRequired() bool`

HasSecureIcaRequired returns a boolean if a field has been set.

### SetSecureIcaRequiredNil

`func (o *CreateDeliveryGroupRequestModel) SetSecureIcaRequiredNil(b bool)`

 SetSecureIcaRequiredNil sets the value for SecureIcaRequired to be an explicit nil

### UnsetSecureIcaRequired
`func (o *CreateDeliveryGroupRequestModel) UnsetSecureIcaRequired()`

UnsetSecureIcaRequired ensures that no value is present for SecureIcaRequired, not even an explicit nil
### GetSettlementPeriodBeforeAutoShutdownSeconds

`func (o *CreateDeliveryGroupRequestModel) GetSettlementPeriodBeforeAutoShutdownSeconds() int32`

GetSettlementPeriodBeforeAutoShutdownSeconds returns the SettlementPeriodBeforeAutoShutdownSeconds field if non-nil, zero value otherwise.

### GetSettlementPeriodBeforeAutoShutdownSecondsOk

`func (o *CreateDeliveryGroupRequestModel) GetSettlementPeriodBeforeAutoShutdownSecondsOk() (*int32, bool)`

GetSettlementPeriodBeforeAutoShutdownSecondsOk returns a tuple with the SettlementPeriodBeforeAutoShutdownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementPeriodBeforeAutoShutdownSeconds

`func (o *CreateDeliveryGroupRequestModel) SetSettlementPeriodBeforeAutoShutdownSeconds(v int32)`

SetSettlementPeriodBeforeAutoShutdownSeconds sets SettlementPeriodBeforeAutoShutdownSeconds field to given value.

### HasSettlementPeriodBeforeAutoShutdownSeconds

`func (o *CreateDeliveryGroupRequestModel) HasSettlementPeriodBeforeAutoShutdownSeconds() bool`

HasSettlementPeriodBeforeAutoShutdownSeconds returns a boolean if a field has been set.

### SetSettlementPeriodBeforeAutoShutdownSecondsNil

`func (o *CreateDeliveryGroupRequestModel) SetSettlementPeriodBeforeAutoShutdownSecondsNil(b bool)`

 SetSettlementPeriodBeforeAutoShutdownSecondsNil sets the value for SettlementPeriodBeforeAutoShutdownSeconds to be an explicit nil

### UnsetSettlementPeriodBeforeAutoShutdownSeconds
`func (o *CreateDeliveryGroupRequestModel) UnsetSettlementPeriodBeforeAutoShutdownSeconds()`

UnsetSettlementPeriodBeforeAutoShutdownSeconds ensures that no value is present for SettlementPeriodBeforeAutoShutdownSeconds, not even an explicit nil
### GetSettlementPeriodBeforeUseSeconds

`func (o *CreateDeliveryGroupRequestModel) GetSettlementPeriodBeforeUseSeconds() int32`

GetSettlementPeriodBeforeUseSeconds returns the SettlementPeriodBeforeUseSeconds field if non-nil, zero value otherwise.

### GetSettlementPeriodBeforeUseSecondsOk

`func (o *CreateDeliveryGroupRequestModel) GetSettlementPeriodBeforeUseSecondsOk() (*int32, bool)`

GetSettlementPeriodBeforeUseSecondsOk returns a tuple with the SettlementPeriodBeforeUseSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementPeriodBeforeUseSeconds

`func (o *CreateDeliveryGroupRequestModel) SetSettlementPeriodBeforeUseSeconds(v int32)`

SetSettlementPeriodBeforeUseSeconds sets SettlementPeriodBeforeUseSeconds field to given value.

### HasSettlementPeriodBeforeUseSeconds

`func (o *CreateDeliveryGroupRequestModel) HasSettlementPeriodBeforeUseSeconds() bool`

HasSettlementPeriodBeforeUseSeconds returns a boolean if a field has been set.

### SetSettlementPeriodBeforeUseSecondsNil

`func (o *CreateDeliveryGroupRequestModel) SetSettlementPeriodBeforeUseSecondsNil(b bool)`

 SetSettlementPeriodBeforeUseSecondsNil sets the value for SettlementPeriodBeforeUseSeconds to be an explicit nil

### UnsetSettlementPeriodBeforeUseSeconds
`func (o *CreateDeliveryGroupRequestModel) UnsetSettlementPeriodBeforeUseSeconds()`

UnsetSettlementPeriodBeforeUseSeconds ensures that no value is present for SettlementPeriodBeforeUseSeconds, not even an explicit nil
### GetShutdownDesktopsAfterUse

`func (o *CreateDeliveryGroupRequestModel) GetShutdownDesktopsAfterUse() bool`

GetShutdownDesktopsAfterUse returns the ShutdownDesktopsAfterUse field if non-nil, zero value otherwise.

### GetShutdownDesktopsAfterUseOk

`func (o *CreateDeliveryGroupRequestModel) GetShutdownDesktopsAfterUseOk() (*bool, bool)`

GetShutdownDesktopsAfterUseOk returns a tuple with the ShutdownDesktopsAfterUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDesktopsAfterUse

`func (o *CreateDeliveryGroupRequestModel) SetShutdownDesktopsAfterUse(v bool)`

SetShutdownDesktopsAfterUse sets ShutdownDesktopsAfterUse field to given value.

### HasShutdownDesktopsAfterUse

`func (o *CreateDeliveryGroupRequestModel) HasShutdownDesktopsAfterUse() bool`

HasShutdownDesktopsAfterUse returns a boolean if a field has been set.

### SetShutdownDesktopsAfterUseNil

`func (o *CreateDeliveryGroupRequestModel) SetShutdownDesktopsAfterUseNil(b bool)`

 SetShutdownDesktopsAfterUseNil sets the value for ShutdownDesktopsAfterUse to be an explicit nil

### UnsetShutdownDesktopsAfterUse
`func (o *CreateDeliveryGroupRequestModel) UnsetShutdownDesktopsAfterUse()`

UnsetShutdownDesktopsAfterUse ensures that no value is present for ShutdownDesktopsAfterUse, not even an explicit nil
### GetSimpleAccessPolicy

`func (o *CreateDeliveryGroupRequestModel) GetSimpleAccessPolicy() SimplifiedAccessPolicyRequestModel`

GetSimpleAccessPolicy returns the SimpleAccessPolicy field if non-nil, zero value otherwise.

### GetSimpleAccessPolicyOk

`func (o *CreateDeliveryGroupRequestModel) GetSimpleAccessPolicyOk() (*SimplifiedAccessPolicyRequestModel, bool)`

GetSimpleAccessPolicyOk returns a tuple with the SimpleAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAccessPolicy

`func (o *CreateDeliveryGroupRequestModel) SetSimpleAccessPolicy(v SimplifiedAccessPolicyRequestModel)`

SetSimpleAccessPolicy sets SimpleAccessPolicy field to given value.

### HasSimpleAccessPolicy

`func (o *CreateDeliveryGroupRequestModel) HasSimpleAccessPolicy() bool`

HasSimpleAccessPolicy returns a boolean if a field has been set.

### GetStoreFrontServersForHostedReceiver

`func (o *CreateDeliveryGroupRequestModel) GetStoreFrontServersForHostedReceiver() []StoreFrontServerRequestModel`

GetStoreFrontServersForHostedReceiver returns the StoreFrontServersForHostedReceiver field if non-nil, zero value otherwise.

### GetStoreFrontServersForHostedReceiverOk

`func (o *CreateDeliveryGroupRequestModel) GetStoreFrontServersForHostedReceiverOk() (*[]StoreFrontServerRequestModel, bool)`

GetStoreFrontServersForHostedReceiverOk returns a tuple with the StoreFrontServersForHostedReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFrontServersForHostedReceiver

`func (o *CreateDeliveryGroupRequestModel) SetStoreFrontServersForHostedReceiver(v []StoreFrontServerRequestModel)`

SetStoreFrontServersForHostedReceiver sets StoreFrontServersForHostedReceiver field to given value.

### HasStoreFrontServersForHostedReceiver

`func (o *CreateDeliveryGroupRequestModel) HasStoreFrontServersForHostedReceiver() bool`

HasStoreFrontServersForHostedReceiver returns a boolean if a field has been set.

### SetStoreFrontServersForHostedReceiverNil

`func (o *CreateDeliveryGroupRequestModel) SetStoreFrontServersForHostedReceiverNil(b bool)`

 SetStoreFrontServersForHostedReceiverNil sets the value for StoreFrontServersForHostedReceiver to be an explicit nil

### UnsetStoreFrontServersForHostedReceiver
`func (o *CreateDeliveryGroupRequestModel) UnsetStoreFrontServersForHostedReceiver()`

UnsetStoreFrontServersForHostedReceiver ensures that no value is present for StoreFrontServersForHostedReceiver, not even an explicit nil
### GetTimeZone

`func (o *CreateDeliveryGroupRequestModel) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CreateDeliveryGroupRequestModel) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CreateDeliveryGroupRequestModel) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CreateDeliveryGroupRequestModel) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *CreateDeliveryGroupRequestModel) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *CreateDeliveryGroupRequestModel) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetTurnOnAddedMachines

`func (o *CreateDeliveryGroupRequestModel) GetTurnOnAddedMachines() bool`

GetTurnOnAddedMachines returns the TurnOnAddedMachines field if non-nil, zero value otherwise.

### GetTurnOnAddedMachinesOk

`func (o *CreateDeliveryGroupRequestModel) GetTurnOnAddedMachinesOk() (*bool, bool)`

GetTurnOnAddedMachinesOk returns a tuple with the TurnOnAddedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOnAddedMachines

`func (o *CreateDeliveryGroupRequestModel) SetTurnOnAddedMachines(v bool)`

SetTurnOnAddedMachines sets TurnOnAddedMachines field to given value.

### HasTurnOnAddedMachines

`func (o *CreateDeliveryGroupRequestModel) HasTurnOnAddedMachines() bool`

HasTurnOnAddedMachines returns a boolean if a field has been set.

### SetTurnOnAddedMachinesNil

`func (o *CreateDeliveryGroupRequestModel) SetTurnOnAddedMachinesNil(b bool)`

 SetTurnOnAddedMachinesNil sets the value for TurnOnAddedMachines to be an explicit nil

### UnsetTurnOnAddedMachines
`func (o *CreateDeliveryGroupRequestModel) UnsetTurnOnAddedMachines()`

UnsetTurnOnAddedMachines ensures that no value is present for TurnOnAddedMachines, not even an explicit nil
### GetZonePreferences

`func (o *CreateDeliveryGroupRequestModel) GetZonePreferences() []ZonePreference`

GetZonePreferences returns the ZonePreferences field if non-nil, zero value otherwise.

### GetZonePreferencesOk

`func (o *CreateDeliveryGroupRequestModel) GetZonePreferencesOk() (*[]ZonePreference, bool)`

GetZonePreferencesOk returns a tuple with the ZonePreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePreferences

`func (o *CreateDeliveryGroupRequestModel) SetZonePreferences(v []ZonePreference)`

SetZonePreferences sets ZonePreferences field to given value.

### HasZonePreferences

`func (o *CreateDeliveryGroupRequestModel) HasZonePreferences() bool`

HasZonePreferences returns a boolean if a field has been set.

### SetZonePreferencesNil

`func (o *CreateDeliveryGroupRequestModel) SetZonePreferencesNil(b bool)`

 SetZonePreferencesNil sets the value for ZonePreferences to be an explicit nil

### UnsetZonePreferences
`func (o *CreateDeliveryGroupRequestModel) UnsetZonePreferences()`

UnsetZonePreferences ensures that no value is present for ZonePreferences, not even an explicit nil
### GetMetadata

`func (o *CreateDeliveryGroupRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateDeliveryGroupRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateDeliveryGroupRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateDeliveryGroupRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateDeliveryGroupRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateDeliveryGroupRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPolicySetGuid

`func (o *CreateDeliveryGroupRequestModel) GetPolicySetGuid() string`

GetPolicySetGuid returns the PolicySetGuid field if non-nil, zero value otherwise.

### GetPolicySetGuidOk

`func (o *CreateDeliveryGroupRequestModel) GetPolicySetGuidOk() (*string, bool)`

GetPolicySetGuidOk returns a tuple with the PolicySetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetGuid

`func (o *CreateDeliveryGroupRequestModel) SetPolicySetGuid(v string)`

SetPolicySetGuid sets PolicySetGuid field to given value.

### HasPolicySetGuid

`func (o *CreateDeliveryGroupRequestModel) HasPolicySetGuid() bool`

HasPolicySetGuid returns a boolean if a field has been set.

### GetRequiredSleepCapability

`func (o *CreateDeliveryGroupRequestModel) GetRequiredSleepCapability() RequiredSleepCapability`

GetRequiredSleepCapability returns the RequiredSleepCapability field if non-nil, zero value otherwise.

### GetRequiredSleepCapabilityOk

`func (o *CreateDeliveryGroupRequestModel) GetRequiredSleepCapabilityOk() (*RequiredSleepCapability, bool)`

GetRequiredSleepCapabilityOk returns a tuple with the RequiredSleepCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSleepCapability

`func (o *CreateDeliveryGroupRequestModel) SetRequiredSleepCapability(v RequiredSleepCapability)`

SetRequiredSleepCapability sets RequiredSleepCapability field to given value.

### HasRequiredSleepCapability

`func (o *CreateDeliveryGroupRequestModel) HasRequiredSleepCapability() bool`

HasRequiredSleepCapability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


