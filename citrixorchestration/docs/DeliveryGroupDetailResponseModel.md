# DeliveryGroupDetailResponseModel

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
**AppAccessPolicy** | Pointer to [**AppAccessPolicyResponseModel**](AppAccessPolicyResponseModel.md) |  | [optional] 
**AppProtectionKeyLoggingRequired** | Pointer to **NullableBool** | Specifies whether key logging app protection is required. | [optional] 
**AppProtectionScreenCaptureRequired** | Pointer to **NullableBool** | Specifies whether screen capture app protection is required. | [optional] 
**AutomaticPowerOnForAssigned** | Pointer to **bool** | Whether assigned (Private) machines in the delivery group should be automatically powered-on at the start of peak time periods. | [optional] 
**AutomaticPowerOnForAssignedDuringPeak** | Pointer to **bool** | Whether assigned (Private) machines in the delivery group should be automatically powered-on throughout peak time periods. | [optional] 
**AutoScaleEnabled** | Pointer to **NullableBool** | Indicates whether auto-scale is enabled for the delivery group. | [optional] 
**RestrictAutoscaleTag** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**ColorDepth** | Pointer to [**ColorDepth**](ColorDepth.md) |  | [optional] 
**DefaultDesktopIconId** | Pointer to **NullableString** | Default icon to use for desktops published from the delivery group. was IconUid | [optional] 
**DefaultDesktopPublishedName** | Pointer to **NullableString** | Default published name to use for desktops published from the delivery group. Change: Add | [optional] 
**DesktopsInUse** | **int32** | Number of machines in the delivery group which are currently in-use. | 
**DesktopsNeverRegistered** | **int32** | Number of machines in the delivery group which have never successfully registered. | 
**DesktopsPreparing** | **int32** | Number of machines in the delivery group which are preparing for a connection. | 
**HdxSslEnabled** | Pointer to **NullableBool** | Whether connections to machines in the delivery group will use SSL. CHANGE: add: | [optional] 
**IsPowerManaged** | Pointer to **bool** | Indicates whether the machines in the delivery group are power-managed. NOTE: I used to think that MachineType&#x3D;&#x3D;Virtual meant the same thing as \&quot;power-managed\&quot;; however that&#39;s not the case.  A machine is power- managed if it is Virtual OR if it is RemotePC with a hypervisor connection (which will still have MachineType&#x3D;&#x3D;Physical). | [optional] 
**LingerSettings** | Pointer to [**FastApplicationSettingsResponseModel**](FastApplicationSettingsResponseModel.md) |  | [optional] 
**PowerOffDelayMinutes** | Pointer to **NullableInt32** | Delay before machines are powered off, when scaling down.  Specified in minutes.  Applies only to multi-session machines. | [optional] 
**MachineCost** | Pointer to **NullableFloat64** | Indicates the estimated per-hour cost for machines in the delivery group, as set by the administrator. | [optional] 
**MachinesInMaintenanceMode** | **int32** | Number of machines in the delivery group which are in maintenance mode. | 
**MachineOperatingSystems** | [**[]NameValueIntPairModel**](NameValueIntPairModel.md) | List of operating systems of machines in the delivery group, along with a count of how many machines use each OS. CHANGE: was public IDictionary&lt;string,int&gt; MachineOperatingSystems { get; set; } | 
**MachineType** | [**MachineType**](MachineType.md) |  | 
**OffMachines** | **int32** | Number of machines in the delivery group which are currently powered off. | 
**OffPeakBufferSizePercent** | Pointer to **int32** | The percentage of machines in the delivery group that should be kept available in an idle state outside peak hours. | [optional] 
**OffPeakDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakDisconnectTimeoutMinutes** | Pointer to **int32** | The number of minutes before the configured action should be performed after a user session disconnects outside peak hours. | [optional] 
**OffPeakExtendedDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakExtendedDisconnectTimeoutMinutes** | Pointer to **int32** | The number of minutes before the second configured action should be performed after a user session disconnects outside peak hours. | [optional] 
**OffPeakLogOffAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakLogOffTimeoutMinutes** | Pointer to **int32** | The number of minutes before the configured action should be performed after a user session ends outside peak hours. | [optional] 
**PeakBufferSizePercent** | Pointer to **int32** | The percentage of machines in the delivery group that should be kept available in an idle state in peak hours. | [optional] 
**PeakDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**PeakDisconnectTimeoutMinutes** | Pointer to **int32** | The number of minutes before the second configured action should be performed after a user session disconnects in peak hours. | [optional] 
**PeakExtendedDisconnectAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**PeakExtendedDisconnectTimeoutMinutes** | Pointer to **int32** | The number of minutes before the second configured action should be performed after a user session disconnects in peak hours. | [optional] 
**PeakLogOffAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**PeakLogOffTimeoutMinutes** | Pointer to **int32** | The number of minutes before the configured action should be performed after a user session ends in peak hours. | [optional] 
**DisconnectPeakIdleSessionAfterSeconds** | Pointer to **int32** | Specifies the time in seconds after which an idle session belonging to the delivery group is disconnected during peak time. | [optional] 
**DisconnectOffPeakIdleSessionAfterSeconds** | Pointer to **int32** | Specifies the time in seconds after which an idle session belonging to the delivery group is disconnected during off-peak time. | [optional] 
**LogoffPeakDisconnectedSessionAfterSeconds** | Pointer to **int32** | Specifies the time in seconds after which a disconnected session belonging to the delivery group is terminated during peak time. | [optional] 
**LogoffOffPeakDisconnectedSessionAfterSeconds** | Pointer to **int32** | Specifies the time in seconds after which a disconnected session belonging to the delivery group is terminated during off-peak time. | [optional] 
**PrelaunchSettings** | Pointer to [**FastApplicationSettingsResponseModel**](FastApplicationSettingsResponseModel.md) |  | [optional] 
**PowerTimeSchemes** | Pointer to [**[]PowerTimeSchemeResponseModel**](PowerTimeSchemeResponseModel.md) | Power management time schemes.  No two schemes will cover the same day of the week. | [optional] 
**ProtocolPriority** | Pointer to [**[]ProtocolType**](ProtocolType.md) | A list of protocols in the order in which they will be attempted for use during connection. | [optional] 
**RegisteredMachines** | **int32** | Number of machines in the delivery group that have registered. | 
**ReuseMachinesWithoutShutdownInOutage** | Pointer to **bool** | Whether machines will be reused without a shutdown in case of an outage. | [optional] 
**SecureIcaRequired** | **bool** | Whether HDX connections to machines in the delivery group require the use of the SecureICA protocol. | 
**SettlementPeriodBeforeAutoShutdownSeconds** | Pointer to **int32** | Time after a session ends during which automatic shutdown requests (for example, shutdown after use, idle pool management) are deferred. Any outstanding shutdown request takes effect after the settlement period expires. This is typically used to configure time to allow logoff scripts to complete. | [optional] 
**SettlementPeriodBeforeUseSeconds** | Pointer to **int32** | Idle period before a machine can be selected to host a new session after registration or the end of a previous session. This is typically used to allow a machine to become idle following processing associated with start-up or logoff actions. A machine may still be selected during the idle period if no other machine is available for immediate use. | [optional] 
**ShutdownDesktopsAfterUse** | Pointer to **NullableBool** | Whether machines in this delivery group should be automatically shut down when each user session completes. | [optional] 
**SimpleAccessPolicy** | Pointer to [**SimplifiedAccessPolicyResponseModel**](SimplifiedAccessPolicyResponseModel.md) |  | [optional] 
**AdvancedAccessPolicy** | Pointer to [**[]AdvancedAccessPolicyResponseModel**](AdvancedAccessPolicyResponseModel.md) | List of Advanced access policy for connections to the delivery group. | [optional] 
**StoreFrontServersForHostedReceiver** | Pointer to [**[]StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md) | List of StoreFront server addresses to configure within hosted receivers that are delivered from the delivery group. CHANGE: was public string[] StoreFronts | [optional] 
**TimeZone** | Pointer to **NullableString** | The time zone in which this delivery group&#39;s machines reside. | [optional] 
**TotalApplicationGroups** | **int32** | Total number of application groups associated with the delivery group. | 
**TotalMachines** | **int32** | Total number of machines in the delivery group. | 
**TurnOnAddedMachines** | Pointer to **NullableBool** | Whether to attempt to power on machines when they are added to the delivery group. was TurnOnAddedMachine | [optional] 
**UnassignedMachines** | Pointer to **NullableInt32** | Number of unassigned machines in the delivery group. | [optional] 
**UsedByGroupPolicy** | Pointer to **NullableBool** | Check if desktop group is used by group policy. | [optional] 
**VdaVersions** | Pointer to [**[]NameValueIntPairModel**](NameValueIntPairModel.md) | List of the versions, and count of each version, of VDAs within the delivery group. CHANGE: was public IDictionary&lt;string,int&gt; VdaVersions { get; set; } | [optional] 
**ZonePreferences** | Pointer to [**[]ZonePreference**](ZonePreference.md) | Ordered list of zone preferences to be applied when launching resources from this delivery group. | [optional] 
**LimitSecondsToForceLogOffUserDuringPeak** | Pointer to **NullableInt32** | Represents the number of seconds that must elapsed before the active sessions on the draining machines belonging to the delivery group are logged off, during peak time. | [optional] 
**LimitSecondsToForceLogOffUserDuringOffPeak** | Pointer to **NullableInt32** | Represents the number of seconds that must elapsed before the active sessions on the draining machines belonging to the delivery group are logged off, during off-peak. | [optional] 
**RestrictAutoscaleMinIdleUntaggedPercentDuringPeak** | Pointer to **NullableInt32** | Represents the percentage that the number of untagged single-session machines in an idle state, or for multi-session machines, the untagged available load capacity must fall below before Autoscale powers on and manages &#39;tagged&#39; machines, as per policy, in peak time. If the number of untagged machines in an idle state, or the untagged available load capacity goes above this threshold value, Autoscale will attempt to shut down &#39;tagged&#39; machines. | [optional] 
**RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak** | Pointer to **NullableInt32** | Represents the percentage that the number of untagged single-session machines in an idle state, or for multi-session machines, the untagged available load capacity must fall below before Autoscale powers on and manages &#39;tagged&#39; machines, as per policy, in off-peak. If the number of untagged machines in an idle state, or the untagged available load capacity goes above this threshold value, Autoscale will attempt to shut down &#39;tagged&#39; machines. | [optional] 
**LogOffWarningMessage** | Pointer to **NullableString** | The warning message to display to users in active sessions prior to logging off users, whether in peak time or off-peak. | [optional] 
**LogOffWarningTitle** | Pointer to **NullableString** | The title of the warning message dialog. | [optional] 
**AutoscaleLogOffReminderEnabled** | Pointer to **NullableBool** | Boolean value indicating whether the warning messages should be sent on an interval to nudge a logoff should be sent on an interval when autoscale is enabled. | [optional] 
**AutoscaleLogOffReminderIntervalSecondsOffPeak** | Pointer to **NullableInt32** | Represents the time interval at which messages are  sent to the user during off peak time when Autoscale is enabled. This message will nudge users to log off instead of forcibly logging them off. | [optional] 
**AutoscaleLogOffReminderIntervalSecondsPeak** | Pointer to **NullableInt32** | Represents the time interval at which messages are  sent to the user during peak time when autoscale is enabled. This message will nudge users to log off instead of forcibly logging them off. | [optional] 
**AutoscaleLogOffReminderMessage** | Pointer to **NullableString** | Notification message to display to users in active sessions belonging to machines needed by Autoscale for shutdown. | [optional] 
**AutoscaleLogOffReminderTitle** | Pointer to **NullableString** | Notification message dialog title displayed when Autoscale issues a logoff reminder request. | [optional] 

## Methods

### NewDeliveryGroupDetailResponseModel

`func NewDeliveryGroupDetailResponseModel(id string, delivering DeliveryKind, deliveryType DeliveryKind, desktopsAvailable int32, desktopsDisconnected int32, desktopsFaulted int32, desktopsUnregistered int32, enabled bool, hasBeenPromoted bool, inMaintenanceMode bool, isBroken bool, isRemotePC bool, minimumFunctionalLevel FunctionalLevel, name string, requireUserHomeZone bool, scopes []ScopeResponseModel, sessionCount int32, sessionSupport SessionSupport, sharingKind SharingKind, totalApplications int32, totalDesktops int32, applicationGroupCompatibility AppGroupCompatibility, applicationCompatibility AppOrDesktopCompatibility, desktopCompatibility AppOrDesktopCompatibility, desktopsInUse int32, desktopsNeverRegistered int32, desktopsPreparing int32, machinesInMaintenanceMode int32, machineOperatingSystems []NameValueIntPairModel, machineType MachineType, offMachines int32, registeredMachines int32, secureIcaRequired bool, totalApplicationGroups int32, totalMachines int32, ) *DeliveryGroupDetailResponseModel`

NewDeliveryGroupDetailResponseModel instantiates a new DeliveryGroupDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupDetailResponseModelWithDefaults

`func NewDeliveryGroupDetailResponseModelWithDefaults() *DeliveryGroupDetailResponseModel`

NewDeliveryGroupDetailResponseModelWithDefaults instantiates a new DeliveryGroupDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryGroupDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryGroupDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryGroupDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *DeliveryGroupDetailResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DeliveryGroupDetailResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DeliveryGroupDetailResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DeliveryGroupDetailResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUserManagement

`func (o *DeliveryGroupDetailResponseModel) GetUserManagement() UserManagementModel`

GetUserManagement returns the UserManagement field if non-nil, zero value otherwise.

### GetUserManagementOk

`func (o *DeliveryGroupDetailResponseModel) GetUserManagementOk() (*UserManagementModel, bool)`

GetUserManagementOk returns a tuple with the UserManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagement

`func (o *DeliveryGroupDetailResponseModel) SetUserManagement(v UserManagementModel)`

SetUserManagement sets UserManagement field to given value.

### HasUserManagement

`func (o *DeliveryGroupDetailResponseModel) HasUserManagement() bool`

HasUserManagement returns a boolean if a field has been set.

### GetDelivering

`func (o *DeliveryGroupDetailResponseModel) GetDelivering() DeliveryKind`

GetDelivering returns the Delivering field if non-nil, zero value otherwise.

### GetDeliveringOk

`func (o *DeliveryGroupDetailResponseModel) GetDeliveringOk() (*DeliveryKind, bool)`

GetDeliveringOk returns a tuple with the Delivering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivering

`func (o *DeliveryGroupDetailResponseModel) SetDelivering(v DeliveryKind)`

SetDelivering sets Delivering field to given value.


### GetDeliveryType

`func (o *DeliveryGroupDetailResponseModel) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *DeliveryGroupDetailResponseModel) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *DeliveryGroupDetailResponseModel) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.


### GetDescription

`func (o *DeliveryGroupDetailResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeliveryGroupDetailResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeliveryGroupDetailResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeliveryGroupDetailResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DeliveryGroupDetailResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeliveryGroupDetailResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesktopsAvailable

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsAvailable() int32`

GetDesktopsAvailable returns the DesktopsAvailable field if non-nil, zero value otherwise.

### GetDesktopsAvailableOk

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsAvailableOk() (*int32, bool)`

GetDesktopsAvailableOk returns a tuple with the DesktopsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsAvailable

`func (o *DeliveryGroupDetailResponseModel) SetDesktopsAvailable(v int32)`

SetDesktopsAvailable sets DesktopsAvailable field to given value.


### GetDesktopsDisconnected

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsDisconnected() int32`

GetDesktopsDisconnected returns the DesktopsDisconnected field if non-nil, zero value otherwise.

### GetDesktopsDisconnectedOk

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsDisconnectedOk() (*int32, bool)`

GetDesktopsDisconnectedOk returns a tuple with the DesktopsDisconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsDisconnected

`func (o *DeliveryGroupDetailResponseModel) SetDesktopsDisconnected(v int32)`

SetDesktopsDisconnected sets DesktopsDisconnected field to given value.


### GetDesktopsFaulted

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsFaulted() int32`

GetDesktopsFaulted returns the DesktopsFaulted field if non-nil, zero value otherwise.

### GetDesktopsFaultedOk

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsFaultedOk() (*int32, bool)`

GetDesktopsFaultedOk returns a tuple with the DesktopsFaulted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsFaulted

`func (o *DeliveryGroupDetailResponseModel) SetDesktopsFaulted(v int32)`

SetDesktopsFaulted sets DesktopsFaulted field to given value.


### GetDesktopsUnregistered

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsUnregistered() int32`

GetDesktopsUnregistered returns the DesktopsUnregistered field if non-nil, zero value otherwise.

### GetDesktopsUnregisteredOk

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsUnregisteredOk() (*int32, bool)`

GetDesktopsUnregisteredOk returns a tuple with the DesktopsUnregistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsUnregistered

`func (o *DeliveryGroupDetailResponseModel) SetDesktopsUnregistered(v int32)`

SetDesktopsUnregistered sets DesktopsUnregistered field to given value.


### GetEnabled

`func (o *DeliveryGroupDetailResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeliveryGroupDetailResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeliveryGroupDetailResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasBeenPromoted

`func (o *DeliveryGroupDetailResponseModel) GetHasBeenPromoted() bool`

GetHasBeenPromoted returns the HasBeenPromoted field if non-nil, zero value otherwise.

### GetHasBeenPromotedOk

`func (o *DeliveryGroupDetailResponseModel) GetHasBeenPromotedOk() (*bool, bool)`

GetHasBeenPromotedOk returns a tuple with the HasBeenPromoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromoted

`func (o *DeliveryGroupDetailResponseModel) SetHasBeenPromoted(v bool)`

SetHasBeenPromoted sets HasBeenPromoted field to given value.


### GetHasBeenPromotedFrom

`func (o *DeliveryGroupDetailResponseModel) GetHasBeenPromotedFrom() FunctionalLevel`

GetHasBeenPromotedFrom returns the HasBeenPromotedFrom field if non-nil, zero value otherwise.

### GetHasBeenPromotedFromOk

`func (o *DeliveryGroupDetailResponseModel) GetHasBeenPromotedFromOk() (*FunctionalLevel, bool)`

GetHasBeenPromotedFromOk returns a tuple with the HasBeenPromotedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenPromotedFrom

`func (o *DeliveryGroupDetailResponseModel) SetHasBeenPromotedFrom(v FunctionalLevel)`

SetHasBeenPromotedFrom sets HasBeenPromotedFrom field to given value.

### HasHasBeenPromotedFrom

`func (o *DeliveryGroupDetailResponseModel) HasHasBeenPromotedFrom() bool`

HasHasBeenPromotedFrom returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *DeliveryGroupDetailResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *DeliveryGroupDetailResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *DeliveryGroupDetailResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetIsBroken

`func (o *DeliveryGroupDetailResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *DeliveryGroupDetailResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *DeliveryGroupDetailResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.


### GetIsRemotePC

`func (o *DeliveryGroupDetailResponseModel) GetIsRemotePC() bool`

GetIsRemotePC returns the IsRemotePC field if non-nil, zero value otherwise.

### GetIsRemotePCOk

`func (o *DeliveryGroupDetailResponseModel) GetIsRemotePCOk() (*bool, bool)`

GetIsRemotePCOk returns a tuple with the IsRemotePC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePC

`func (o *DeliveryGroupDetailResponseModel) SetIsRemotePC(v bool)`

SetIsRemotePC sets IsRemotePC field to given value.


### GetMachineLogOnType

`func (o *DeliveryGroupDetailResponseModel) GetMachineLogOnType() MachineLogOnType`

GetMachineLogOnType returns the MachineLogOnType field if non-nil, zero value otherwise.

### GetMachineLogOnTypeOk

`func (o *DeliveryGroupDetailResponseModel) GetMachineLogOnTypeOk() (*MachineLogOnType, bool)`

GetMachineLogOnTypeOk returns a tuple with the MachineLogOnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLogOnType

`func (o *DeliveryGroupDetailResponseModel) SetMachineLogOnType(v MachineLogOnType)`

SetMachineLogOnType sets MachineLogOnType field to given value.

### HasMachineLogOnType

`func (o *DeliveryGroupDetailResponseModel) HasMachineLogOnType() bool`

HasMachineLogOnType returns a boolean if a field has been set.

### GetMetadata

`func (o *DeliveryGroupDetailResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeliveryGroupDetailResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeliveryGroupDetailResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeliveryGroupDetailResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DeliveryGroupDetailResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DeliveryGroupDetailResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *DeliveryGroupDetailResponseModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *DeliveryGroupDetailResponseModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *DeliveryGroupDetailResponseModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.


### GetName

`func (o *DeliveryGroupDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryGroupDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryGroupDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *DeliveryGroupDetailResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DeliveryGroupDetailResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DeliveryGroupDetailResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *DeliveryGroupDetailResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *DeliveryGroupDetailResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *DeliveryGroupDetailResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPublishedName

`func (o *DeliveryGroupDetailResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *DeliveryGroupDetailResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *DeliveryGroupDetailResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *DeliveryGroupDetailResponseModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *DeliveryGroupDetailResponseModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *DeliveryGroupDetailResponseModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetPolicySetGuid

`func (o *DeliveryGroupDetailResponseModel) GetPolicySetGuid() string`

GetPolicySetGuid returns the PolicySetGuid field if non-nil, zero value otherwise.

### GetPolicySetGuidOk

`func (o *DeliveryGroupDetailResponseModel) GetPolicySetGuidOk() (*string, bool)`

GetPolicySetGuidOk returns a tuple with the PolicySetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetGuid

`func (o *DeliveryGroupDetailResponseModel) SetPolicySetGuid(v string)`

SetPolicySetGuid sets PolicySetGuid field to given value.

### HasPolicySetGuid

`func (o *DeliveryGroupDetailResponseModel) HasPolicySetGuid() bool`

HasPolicySetGuid returns a boolean if a field has been set.

### GetProductCode

`func (o *DeliveryGroupDetailResponseModel) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *DeliveryGroupDetailResponseModel) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *DeliveryGroupDetailResponseModel) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *DeliveryGroupDetailResponseModel) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCodeNil

`func (o *DeliveryGroupDetailResponseModel) SetProductCodeNil(b bool)`

 SetProductCodeNil sets the value for ProductCode to be an explicit nil

### UnsetProductCode
`func (o *DeliveryGroupDetailResponseModel) UnsetProductCode()`

UnsetProductCode ensures that no value is present for ProductCode, not even an explicit nil
### GetLicenseModel

`func (o *DeliveryGroupDetailResponseModel) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *DeliveryGroupDetailResponseModel) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *DeliveryGroupDetailResponseModel) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *DeliveryGroupDetailResponseModel) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetRequireUserHomeZone

`func (o *DeliveryGroupDetailResponseModel) GetRequireUserHomeZone() bool`

GetRequireUserHomeZone returns the RequireUserHomeZone field if non-nil, zero value otherwise.

### GetRequireUserHomeZoneOk

`func (o *DeliveryGroupDetailResponseModel) GetRequireUserHomeZoneOk() (*bool, bool)`

GetRequireUserHomeZoneOk returns a tuple with the RequireUserHomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireUserHomeZone

`func (o *DeliveryGroupDetailResponseModel) SetRequireUserHomeZone(v bool)`

SetRequireUserHomeZone sets RequireUserHomeZone field to given value.


### GetScopes

`func (o *DeliveryGroupDetailResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DeliveryGroupDetailResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DeliveryGroupDetailResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTenants

`func (o *DeliveryGroupDetailResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *DeliveryGroupDetailResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *DeliveryGroupDetailResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *DeliveryGroupDetailResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *DeliveryGroupDetailResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *DeliveryGroupDetailResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetSessionCount

`func (o *DeliveryGroupDetailResponseModel) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *DeliveryGroupDetailResponseModel) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *DeliveryGroupDetailResponseModel) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.


### GetSessionSupport

`func (o *DeliveryGroupDetailResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *DeliveryGroupDetailResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *DeliveryGroupDetailResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSharingKind

`func (o *DeliveryGroupDetailResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *DeliveryGroupDetailResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *DeliveryGroupDetailResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.


### GetTotalApplications

`func (o *DeliveryGroupDetailResponseModel) GetTotalApplications() int32`

GetTotalApplications returns the TotalApplications field if non-nil, zero value otherwise.

### GetTotalApplicationsOk

`func (o *DeliveryGroupDetailResponseModel) GetTotalApplicationsOk() (*int32, bool)`

GetTotalApplicationsOk returns a tuple with the TotalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplications

`func (o *DeliveryGroupDetailResponseModel) SetTotalApplications(v int32)`

SetTotalApplications sets TotalApplications field to given value.


### GetTotalDesktops

`func (o *DeliveryGroupDetailResponseModel) GetTotalDesktops() int32`

GetTotalDesktops returns the TotalDesktops field if non-nil, zero value otherwise.

### GetTotalDesktopsOk

`func (o *DeliveryGroupDetailResponseModel) GetTotalDesktopsOk() (*int32, bool)`

GetTotalDesktopsOk returns a tuple with the TotalDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDesktops

`func (o *DeliveryGroupDetailResponseModel) SetTotalDesktops(v int32)`

SetTotalDesktops sets TotalDesktops field to given value.


### GetApplicationGroupCompatibility

`func (o *DeliveryGroupDetailResponseModel) GetApplicationGroupCompatibility() AppGroupCompatibility`

GetApplicationGroupCompatibility returns the ApplicationGroupCompatibility field if non-nil, zero value otherwise.

### GetApplicationGroupCompatibilityOk

`func (o *DeliveryGroupDetailResponseModel) GetApplicationGroupCompatibilityOk() (*AppGroupCompatibility, bool)`

GetApplicationGroupCompatibilityOk returns a tuple with the ApplicationGroupCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroupCompatibility

`func (o *DeliveryGroupDetailResponseModel) SetApplicationGroupCompatibility(v AppGroupCompatibility)`

SetApplicationGroupCompatibility sets ApplicationGroupCompatibility field to given value.


### GetApplicationCompatibility

`func (o *DeliveryGroupDetailResponseModel) GetApplicationCompatibility() AppOrDesktopCompatibility`

GetApplicationCompatibility returns the ApplicationCompatibility field if non-nil, zero value otherwise.

### GetApplicationCompatibilityOk

`func (o *DeliveryGroupDetailResponseModel) GetApplicationCompatibilityOk() (*AppOrDesktopCompatibility, bool)`

GetApplicationCompatibilityOk returns a tuple with the ApplicationCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCompatibility

`func (o *DeliveryGroupDetailResponseModel) SetApplicationCompatibility(v AppOrDesktopCompatibility)`

SetApplicationCompatibility sets ApplicationCompatibility field to given value.


### GetDesktopCompatibility

`func (o *DeliveryGroupDetailResponseModel) GetDesktopCompatibility() AppOrDesktopCompatibility`

GetDesktopCompatibility returns the DesktopCompatibility field if non-nil, zero value otherwise.

### GetDesktopCompatibilityOk

`func (o *DeliveryGroupDetailResponseModel) GetDesktopCompatibilityOk() (*AppOrDesktopCompatibility, bool)`

GetDesktopCompatibilityOk returns a tuple with the DesktopCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopCompatibility

`func (o *DeliveryGroupDetailResponseModel) SetDesktopCompatibility(v AppOrDesktopCompatibility)`

SetDesktopCompatibility sets DesktopCompatibility field to given value.


### GetAdminFolder

`func (o *DeliveryGroupDetailResponseModel) GetAdminFolder() RefResponseModel`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *DeliveryGroupDetailResponseModel) GetAdminFolderOk() (*RefResponseModel, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *DeliveryGroupDetailResponseModel) SetAdminFolder(v RefResponseModel)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *DeliveryGroupDetailResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### GetAppAccessPolicy

`func (o *DeliveryGroupDetailResponseModel) GetAppAccessPolicy() AppAccessPolicyResponseModel`

GetAppAccessPolicy returns the AppAccessPolicy field if non-nil, zero value otherwise.

### GetAppAccessPolicyOk

`func (o *DeliveryGroupDetailResponseModel) GetAppAccessPolicyOk() (*AppAccessPolicyResponseModel, bool)`

GetAppAccessPolicyOk returns a tuple with the AppAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAccessPolicy

`func (o *DeliveryGroupDetailResponseModel) SetAppAccessPolicy(v AppAccessPolicyResponseModel)`

SetAppAccessPolicy sets AppAccessPolicy field to given value.

### HasAppAccessPolicy

`func (o *DeliveryGroupDetailResponseModel) HasAppAccessPolicy() bool`

HasAppAccessPolicy returns a boolean if a field has been set.

### GetAppProtectionKeyLoggingRequired

`func (o *DeliveryGroupDetailResponseModel) GetAppProtectionKeyLoggingRequired() bool`

GetAppProtectionKeyLoggingRequired returns the AppProtectionKeyLoggingRequired field if non-nil, zero value otherwise.

### GetAppProtectionKeyLoggingRequiredOk

`func (o *DeliveryGroupDetailResponseModel) GetAppProtectionKeyLoggingRequiredOk() (*bool, bool)`

GetAppProtectionKeyLoggingRequiredOk returns a tuple with the AppProtectionKeyLoggingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtectionKeyLoggingRequired

`func (o *DeliveryGroupDetailResponseModel) SetAppProtectionKeyLoggingRequired(v bool)`

SetAppProtectionKeyLoggingRequired sets AppProtectionKeyLoggingRequired field to given value.

### HasAppProtectionKeyLoggingRequired

`func (o *DeliveryGroupDetailResponseModel) HasAppProtectionKeyLoggingRequired() bool`

HasAppProtectionKeyLoggingRequired returns a boolean if a field has been set.

### SetAppProtectionKeyLoggingRequiredNil

`func (o *DeliveryGroupDetailResponseModel) SetAppProtectionKeyLoggingRequiredNil(b bool)`

 SetAppProtectionKeyLoggingRequiredNil sets the value for AppProtectionKeyLoggingRequired to be an explicit nil

### UnsetAppProtectionKeyLoggingRequired
`func (o *DeliveryGroupDetailResponseModel) UnsetAppProtectionKeyLoggingRequired()`

UnsetAppProtectionKeyLoggingRequired ensures that no value is present for AppProtectionKeyLoggingRequired, not even an explicit nil
### GetAppProtectionScreenCaptureRequired

`func (o *DeliveryGroupDetailResponseModel) GetAppProtectionScreenCaptureRequired() bool`

GetAppProtectionScreenCaptureRequired returns the AppProtectionScreenCaptureRequired field if non-nil, zero value otherwise.

### GetAppProtectionScreenCaptureRequiredOk

`func (o *DeliveryGroupDetailResponseModel) GetAppProtectionScreenCaptureRequiredOk() (*bool, bool)`

GetAppProtectionScreenCaptureRequiredOk returns a tuple with the AppProtectionScreenCaptureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtectionScreenCaptureRequired

`func (o *DeliveryGroupDetailResponseModel) SetAppProtectionScreenCaptureRequired(v bool)`

SetAppProtectionScreenCaptureRequired sets AppProtectionScreenCaptureRequired field to given value.

### HasAppProtectionScreenCaptureRequired

`func (o *DeliveryGroupDetailResponseModel) HasAppProtectionScreenCaptureRequired() bool`

HasAppProtectionScreenCaptureRequired returns a boolean if a field has been set.

### SetAppProtectionScreenCaptureRequiredNil

`func (o *DeliveryGroupDetailResponseModel) SetAppProtectionScreenCaptureRequiredNil(b bool)`

 SetAppProtectionScreenCaptureRequiredNil sets the value for AppProtectionScreenCaptureRequired to be an explicit nil

### UnsetAppProtectionScreenCaptureRequired
`func (o *DeliveryGroupDetailResponseModel) UnsetAppProtectionScreenCaptureRequired()`

UnsetAppProtectionScreenCaptureRequired ensures that no value is present for AppProtectionScreenCaptureRequired, not even an explicit nil
### GetAutomaticPowerOnForAssigned

`func (o *DeliveryGroupDetailResponseModel) GetAutomaticPowerOnForAssigned() bool`

GetAutomaticPowerOnForAssigned returns the AutomaticPowerOnForAssigned field if non-nil, zero value otherwise.

### GetAutomaticPowerOnForAssignedOk

`func (o *DeliveryGroupDetailResponseModel) GetAutomaticPowerOnForAssignedOk() (*bool, bool)`

GetAutomaticPowerOnForAssignedOk returns a tuple with the AutomaticPowerOnForAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPowerOnForAssigned

`func (o *DeliveryGroupDetailResponseModel) SetAutomaticPowerOnForAssigned(v bool)`

SetAutomaticPowerOnForAssigned sets AutomaticPowerOnForAssigned field to given value.

### HasAutomaticPowerOnForAssigned

`func (o *DeliveryGroupDetailResponseModel) HasAutomaticPowerOnForAssigned() bool`

HasAutomaticPowerOnForAssigned returns a boolean if a field has been set.

### GetAutomaticPowerOnForAssignedDuringPeak

`func (o *DeliveryGroupDetailResponseModel) GetAutomaticPowerOnForAssignedDuringPeak() bool`

GetAutomaticPowerOnForAssignedDuringPeak returns the AutomaticPowerOnForAssignedDuringPeak field if non-nil, zero value otherwise.

### GetAutomaticPowerOnForAssignedDuringPeakOk

`func (o *DeliveryGroupDetailResponseModel) GetAutomaticPowerOnForAssignedDuringPeakOk() (*bool, bool)`

GetAutomaticPowerOnForAssignedDuringPeakOk returns a tuple with the AutomaticPowerOnForAssignedDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPowerOnForAssignedDuringPeak

`func (o *DeliveryGroupDetailResponseModel) SetAutomaticPowerOnForAssignedDuringPeak(v bool)`

SetAutomaticPowerOnForAssignedDuringPeak sets AutomaticPowerOnForAssignedDuringPeak field to given value.

### HasAutomaticPowerOnForAssignedDuringPeak

`func (o *DeliveryGroupDetailResponseModel) HasAutomaticPowerOnForAssignedDuringPeak() bool`

HasAutomaticPowerOnForAssignedDuringPeak returns a boolean if a field has been set.

### GetAutoScaleEnabled

`func (o *DeliveryGroupDetailResponseModel) GetAutoScaleEnabled() bool`

GetAutoScaleEnabled returns the AutoScaleEnabled field if non-nil, zero value otherwise.

### GetAutoScaleEnabledOk

`func (o *DeliveryGroupDetailResponseModel) GetAutoScaleEnabledOk() (*bool, bool)`

GetAutoScaleEnabledOk returns a tuple with the AutoScaleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaleEnabled

`func (o *DeliveryGroupDetailResponseModel) SetAutoScaleEnabled(v bool)`

SetAutoScaleEnabled sets AutoScaleEnabled field to given value.

### HasAutoScaleEnabled

`func (o *DeliveryGroupDetailResponseModel) HasAutoScaleEnabled() bool`

HasAutoScaleEnabled returns a boolean if a field has been set.

### SetAutoScaleEnabledNil

`func (o *DeliveryGroupDetailResponseModel) SetAutoScaleEnabledNil(b bool)`

 SetAutoScaleEnabledNil sets the value for AutoScaleEnabled to be an explicit nil

### UnsetAutoScaleEnabled
`func (o *DeliveryGroupDetailResponseModel) UnsetAutoScaleEnabled()`

UnsetAutoScaleEnabled ensures that no value is present for AutoScaleEnabled, not even an explicit nil
### GetRestrictAutoscaleTag

`func (o *DeliveryGroupDetailResponseModel) GetRestrictAutoscaleTag() RefResponseModel`

GetRestrictAutoscaleTag returns the RestrictAutoscaleTag field if non-nil, zero value otherwise.

### GetRestrictAutoscaleTagOk

`func (o *DeliveryGroupDetailResponseModel) GetRestrictAutoscaleTagOk() (*RefResponseModel, bool)`

GetRestrictAutoscaleTagOk returns a tuple with the RestrictAutoscaleTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleTag

`func (o *DeliveryGroupDetailResponseModel) SetRestrictAutoscaleTag(v RefResponseModel)`

SetRestrictAutoscaleTag sets RestrictAutoscaleTag field to given value.

### HasRestrictAutoscaleTag

`func (o *DeliveryGroupDetailResponseModel) HasRestrictAutoscaleTag() bool`

HasRestrictAutoscaleTag returns a boolean if a field has been set.

### GetColorDepth

`func (o *DeliveryGroupDetailResponseModel) GetColorDepth() ColorDepth`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *DeliveryGroupDetailResponseModel) GetColorDepthOk() (*ColorDepth, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *DeliveryGroupDetailResponseModel) SetColorDepth(v ColorDepth)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *DeliveryGroupDetailResponseModel) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetDefaultDesktopIconId

`func (o *DeliveryGroupDetailResponseModel) GetDefaultDesktopIconId() string`

GetDefaultDesktopIconId returns the DefaultDesktopIconId field if non-nil, zero value otherwise.

### GetDefaultDesktopIconIdOk

`func (o *DeliveryGroupDetailResponseModel) GetDefaultDesktopIconIdOk() (*string, bool)`

GetDefaultDesktopIconIdOk returns a tuple with the DefaultDesktopIconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDesktopIconId

`func (o *DeliveryGroupDetailResponseModel) SetDefaultDesktopIconId(v string)`

SetDefaultDesktopIconId sets DefaultDesktopIconId field to given value.

### HasDefaultDesktopIconId

`func (o *DeliveryGroupDetailResponseModel) HasDefaultDesktopIconId() bool`

HasDefaultDesktopIconId returns a boolean if a field has been set.

### SetDefaultDesktopIconIdNil

`func (o *DeliveryGroupDetailResponseModel) SetDefaultDesktopIconIdNil(b bool)`

 SetDefaultDesktopIconIdNil sets the value for DefaultDesktopIconId to be an explicit nil

### UnsetDefaultDesktopIconId
`func (o *DeliveryGroupDetailResponseModel) UnsetDefaultDesktopIconId()`

UnsetDefaultDesktopIconId ensures that no value is present for DefaultDesktopIconId, not even an explicit nil
### GetDefaultDesktopPublishedName

`func (o *DeliveryGroupDetailResponseModel) GetDefaultDesktopPublishedName() string`

GetDefaultDesktopPublishedName returns the DefaultDesktopPublishedName field if non-nil, zero value otherwise.

### GetDefaultDesktopPublishedNameOk

`func (o *DeliveryGroupDetailResponseModel) GetDefaultDesktopPublishedNameOk() (*string, bool)`

GetDefaultDesktopPublishedNameOk returns a tuple with the DefaultDesktopPublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDesktopPublishedName

`func (o *DeliveryGroupDetailResponseModel) SetDefaultDesktopPublishedName(v string)`

SetDefaultDesktopPublishedName sets DefaultDesktopPublishedName field to given value.

### HasDefaultDesktopPublishedName

`func (o *DeliveryGroupDetailResponseModel) HasDefaultDesktopPublishedName() bool`

HasDefaultDesktopPublishedName returns a boolean if a field has been set.

### SetDefaultDesktopPublishedNameNil

`func (o *DeliveryGroupDetailResponseModel) SetDefaultDesktopPublishedNameNil(b bool)`

 SetDefaultDesktopPublishedNameNil sets the value for DefaultDesktopPublishedName to be an explicit nil

### UnsetDefaultDesktopPublishedName
`func (o *DeliveryGroupDetailResponseModel) UnsetDefaultDesktopPublishedName()`

UnsetDefaultDesktopPublishedName ensures that no value is present for DefaultDesktopPublishedName, not even an explicit nil
### GetDesktopsInUse

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsInUse() int32`

GetDesktopsInUse returns the DesktopsInUse field if non-nil, zero value otherwise.

### GetDesktopsInUseOk

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsInUseOk() (*int32, bool)`

GetDesktopsInUseOk returns a tuple with the DesktopsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsInUse

`func (o *DeliveryGroupDetailResponseModel) SetDesktopsInUse(v int32)`

SetDesktopsInUse sets DesktopsInUse field to given value.


### GetDesktopsNeverRegistered

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsNeverRegistered() int32`

GetDesktopsNeverRegistered returns the DesktopsNeverRegistered field if non-nil, zero value otherwise.

### GetDesktopsNeverRegisteredOk

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsNeverRegisteredOk() (*int32, bool)`

GetDesktopsNeverRegisteredOk returns a tuple with the DesktopsNeverRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsNeverRegistered

`func (o *DeliveryGroupDetailResponseModel) SetDesktopsNeverRegistered(v int32)`

SetDesktopsNeverRegistered sets DesktopsNeverRegistered field to given value.


### GetDesktopsPreparing

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsPreparing() int32`

GetDesktopsPreparing returns the DesktopsPreparing field if non-nil, zero value otherwise.

### GetDesktopsPreparingOk

`func (o *DeliveryGroupDetailResponseModel) GetDesktopsPreparingOk() (*int32, bool)`

GetDesktopsPreparingOk returns a tuple with the DesktopsPreparing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsPreparing

`func (o *DeliveryGroupDetailResponseModel) SetDesktopsPreparing(v int32)`

SetDesktopsPreparing sets DesktopsPreparing field to given value.


### GetHdxSslEnabled

`func (o *DeliveryGroupDetailResponseModel) GetHdxSslEnabled() bool`

GetHdxSslEnabled returns the HdxSslEnabled field if non-nil, zero value otherwise.

### GetHdxSslEnabledOk

`func (o *DeliveryGroupDetailResponseModel) GetHdxSslEnabledOk() (*bool, bool)`

GetHdxSslEnabledOk returns a tuple with the HdxSslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdxSslEnabled

`func (o *DeliveryGroupDetailResponseModel) SetHdxSslEnabled(v bool)`

SetHdxSslEnabled sets HdxSslEnabled field to given value.

### HasHdxSslEnabled

`func (o *DeliveryGroupDetailResponseModel) HasHdxSslEnabled() bool`

HasHdxSslEnabled returns a boolean if a field has been set.

### SetHdxSslEnabledNil

`func (o *DeliveryGroupDetailResponseModel) SetHdxSslEnabledNil(b bool)`

 SetHdxSslEnabledNil sets the value for HdxSslEnabled to be an explicit nil

### UnsetHdxSslEnabled
`func (o *DeliveryGroupDetailResponseModel) UnsetHdxSslEnabled()`

UnsetHdxSslEnabled ensures that no value is present for HdxSslEnabled, not even an explicit nil
### GetIsPowerManaged

`func (o *DeliveryGroupDetailResponseModel) GetIsPowerManaged() bool`

GetIsPowerManaged returns the IsPowerManaged field if non-nil, zero value otherwise.

### GetIsPowerManagedOk

`func (o *DeliveryGroupDetailResponseModel) GetIsPowerManagedOk() (*bool, bool)`

GetIsPowerManagedOk returns a tuple with the IsPowerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPowerManaged

`func (o *DeliveryGroupDetailResponseModel) SetIsPowerManaged(v bool)`

SetIsPowerManaged sets IsPowerManaged field to given value.

### HasIsPowerManaged

`func (o *DeliveryGroupDetailResponseModel) HasIsPowerManaged() bool`

HasIsPowerManaged returns a boolean if a field has been set.

### GetLingerSettings

`func (o *DeliveryGroupDetailResponseModel) GetLingerSettings() FastApplicationSettingsResponseModel`

GetLingerSettings returns the LingerSettings field if non-nil, zero value otherwise.

### GetLingerSettingsOk

`func (o *DeliveryGroupDetailResponseModel) GetLingerSettingsOk() (*FastApplicationSettingsResponseModel, bool)`

GetLingerSettingsOk returns a tuple with the LingerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLingerSettings

`func (o *DeliveryGroupDetailResponseModel) SetLingerSettings(v FastApplicationSettingsResponseModel)`

SetLingerSettings sets LingerSettings field to given value.

### HasLingerSettings

`func (o *DeliveryGroupDetailResponseModel) HasLingerSettings() bool`

HasLingerSettings returns a boolean if a field has been set.

### GetPowerOffDelayMinutes

`func (o *DeliveryGroupDetailResponseModel) GetPowerOffDelayMinutes() int32`

GetPowerOffDelayMinutes returns the PowerOffDelayMinutes field if non-nil, zero value otherwise.

### GetPowerOffDelayMinutesOk

`func (o *DeliveryGroupDetailResponseModel) GetPowerOffDelayMinutesOk() (*int32, bool)`

GetPowerOffDelayMinutesOk returns a tuple with the PowerOffDelayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffDelayMinutes

`func (o *DeliveryGroupDetailResponseModel) SetPowerOffDelayMinutes(v int32)`

SetPowerOffDelayMinutes sets PowerOffDelayMinutes field to given value.

### HasPowerOffDelayMinutes

`func (o *DeliveryGroupDetailResponseModel) HasPowerOffDelayMinutes() bool`

HasPowerOffDelayMinutes returns a boolean if a field has been set.

### SetPowerOffDelayMinutesNil

`func (o *DeliveryGroupDetailResponseModel) SetPowerOffDelayMinutesNil(b bool)`

 SetPowerOffDelayMinutesNil sets the value for PowerOffDelayMinutes to be an explicit nil

### UnsetPowerOffDelayMinutes
`func (o *DeliveryGroupDetailResponseModel) UnsetPowerOffDelayMinutes()`

UnsetPowerOffDelayMinutes ensures that no value is present for PowerOffDelayMinutes, not even an explicit nil
### GetMachineCost

`func (o *DeliveryGroupDetailResponseModel) GetMachineCost() float64`

GetMachineCost returns the MachineCost field if non-nil, zero value otherwise.

### GetMachineCostOk

`func (o *DeliveryGroupDetailResponseModel) GetMachineCostOk() (*float64, bool)`

GetMachineCostOk returns a tuple with the MachineCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCost

`func (o *DeliveryGroupDetailResponseModel) SetMachineCost(v float64)`

SetMachineCost sets MachineCost field to given value.

### HasMachineCost

`func (o *DeliveryGroupDetailResponseModel) HasMachineCost() bool`

HasMachineCost returns a boolean if a field has been set.

### SetMachineCostNil

`func (o *DeliveryGroupDetailResponseModel) SetMachineCostNil(b bool)`

 SetMachineCostNil sets the value for MachineCost to be an explicit nil

### UnsetMachineCost
`func (o *DeliveryGroupDetailResponseModel) UnsetMachineCost()`

UnsetMachineCost ensures that no value is present for MachineCost, not even an explicit nil
### GetMachinesInMaintenanceMode

`func (o *DeliveryGroupDetailResponseModel) GetMachinesInMaintenanceMode() int32`

GetMachinesInMaintenanceMode returns the MachinesInMaintenanceMode field if non-nil, zero value otherwise.

### GetMachinesInMaintenanceModeOk

`func (o *DeliveryGroupDetailResponseModel) GetMachinesInMaintenanceModeOk() (*int32, bool)`

GetMachinesInMaintenanceModeOk returns a tuple with the MachinesInMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinesInMaintenanceMode

`func (o *DeliveryGroupDetailResponseModel) SetMachinesInMaintenanceMode(v int32)`

SetMachinesInMaintenanceMode sets MachinesInMaintenanceMode field to given value.


### GetMachineOperatingSystems

`func (o *DeliveryGroupDetailResponseModel) GetMachineOperatingSystems() []NameValueIntPairModel`

GetMachineOperatingSystems returns the MachineOperatingSystems field if non-nil, zero value otherwise.

### GetMachineOperatingSystemsOk

`func (o *DeliveryGroupDetailResponseModel) GetMachineOperatingSystemsOk() (*[]NameValueIntPairModel, bool)`

GetMachineOperatingSystemsOk returns a tuple with the MachineOperatingSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineOperatingSystems

`func (o *DeliveryGroupDetailResponseModel) SetMachineOperatingSystems(v []NameValueIntPairModel)`

SetMachineOperatingSystems sets MachineOperatingSystems field to given value.


### GetMachineType

`func (o *DeliveryGroupDetailResponseModel) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *DeliveryGroupDetailResponseModel) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *DeliveryGroupDetailResponseModel) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.


### GetOffMachines

`func (o *DeliveryGroupDetailResponseModel) GetOffMachines() int32`

GetOffMachines returns the OffMachines field if non-nil, zero value otherwise.

### GetOffMachinesOk

`func (o *DeliveryGroupDetailResponseModel) GetOffMachinesOk() (*int32, bool)`

GetOffMachinesOk returns a tuple with the OffMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffMachines

`func (o *DeliveryGroupDetailResponseModel) SetOffMachines(v int32)`

SetOffMachines sets OffMachines field to given value.


### GetOffPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakBufferSizePercent() int32`

GetOffPeakBufferSizePercent returns the OffPeakBufferSizePercent field if non-nil, zero value otherwise.

### GetOffPeakBufferSizePercentOk

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakBufferSizePercentOk() (*int32, bool)`

GetOffPeakBufferSizePercentOk returns a tuple with the OffPeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModel) SetOffPeakBufferSizePercent(v int32)`

SetOffPeakBufferSizePercent sets OffPeakBufferSizePercent field to given value.

### HasOffPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModel) HasOffPeakBufferSizePercent() bool`

HasOffPeakBufferSizePercent returns a boolean if a field has been set.

### GetOffPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakDisconnectAction() SessionChangeHostingAction`

GetOffPeakDisconnectAction returns the OffPeakDisconnectAction field if non-nil, zero value otherwise.

### GetOffPeakDisconnectActionOk

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakDisconnectActionOk returns a tuple with the OffPeakDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) SetOffPeakDisconnectAction(v SessionChangeHostingAction)`

SetOffPeakDisconnectAction sets OffPeakDisconnectAction field to given value.

### HasOffPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) HasOffPeakDisconnectAction() bool`

HasOffPeakDisconnectAction returns a boolean if a field has been set.

### GetOffPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakDisconnectTimeoutMinutes() int32`

GetOffPeakDisconnectTimeoutMinutes returns the OffPeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakDisconnectTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakDisconnectTimeoutMinutesOk returns a tuple with the OffPeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) SetOffPeakDisconnectTimeoutMinutes(v int32)`

SetOffPeakDisconnectTimeoutMinutes sets OffPeakDisconnectTimeoutMinutes field to given value.

### HasOffPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) HasOffPeakDisconnectTimeoutMinutes() bool`

HasOffPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetOffPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakExtendedDisconnectAction() SessionChangeHostingAction`

GetOffPeakExtendedDisconnectAction returns the OffPeakExtendedDisconnectAction field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectActionOk

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakExtendedDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakExtendedDisconnectActionOk returns a tuple with the OffPeakExtendedDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) SetOffPeakExtendedDisconnectAction(v SessionChangeHostingAction)`

SetOffPeakExtendedDisconnectAction sets OffPeakExtendedDisconnectAction field to given value.

### HasOffPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) HasOffPeakExtendedDisconnectAction() bool`

HasOffPeakExtendedDisconnectAction returns a boolean if a field has been set.

### GetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakExtendedDisconnectTimeoutMinutes() int32`

GetOffPeakExtendedDisconnectTimeoutMinutes returns the OffPeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the OffPeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) SetOffPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetOffPeakExtendedDisconnectTimeoutMinutes sets OffPeakExtendedDisconnectTimeoutMinutes field to given value.

### HasOffPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) HasOffPeakExtendedDisconnectTimeoutMinutes() bool`

HasOffPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetOffPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakLogOffAction() SessionChangeHostingAction`

GetOffPeakLogOffAction returns the OffPeakLogOffAction field if non-nil, zero value otherwise.

### GetOffPeakLogOffActionOk

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakLogOffActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakLogOffActionOk returns a tuple with the OffPeakLogOffAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModel) SetOffPeakLogOffAction(v SessionChangeHostingAction)`

SetOffPeakLogOffAction sets OffPeakLogOffAction field to given value.

### HasOffPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModel) HasOffPeakLogOffAction() bool`

HasOffPeakLogOffAction returns a boolean if a field has been set.

### GetOffPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakLogOffTimeoutMinutes() int32`

GetOffPeakLogOffTimeoutMinutes returns the OffPeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakLogOffTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModel) GetOffPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetOffPeakLogOffTimeoutMinutesOk returns a tuple with the OffPeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) SetOffPeakLogOffTimeoutMinutes(v int32)`

SetOffPeakLogOffTimeoutMinutes sets OffPeakLogOffTimeoutMinutes field to given value.

### HasOffPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) HasOffPeakLogOffTimeoutMinutes() bool`

HasOffPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### GetPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModel) GetPeakBufferSizePercent() int32`

GetPeakBufferSizePercent returns the PeakBufferSizePercent field if non-nil, zero value otherwise.

### GetPeakBufferSizePercentOk

`func (o *DeliveryGroupDetailResponseModel) GetPeakBufferSizePercentOk() (*int32, bool)`

GetPeakBufferSizePercentOk returns a tuple with the PeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModel) SetPeakBufferSizePercent(v int32)`

SetPeakBufferSizePercent sets PeakBufferSizePercent field to given value.

### HasPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModel) HasPeakBufferSizePercent() bool`

HasPeakBufferSizePercent returns a boolean if a field has been set.

### GetPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) GetPeakDisconnectAction() SessionChangeHostingAction`

GetPeakDisconnectAction returns the PeakDisconnectAction field if non-nil, zero value otherwise.

### GetPeakDisconnectActionOk

`func (o *DeliveryGroupDetailResponseModel) GetPeakDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetPeakDisconnectActionOk returns a tuple with the PeakDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) SetPeakDisconnectAction(v SessionChangeHostingAction)`

SetPeakDisconnectAction sets PeakDisconnectAction field to given value.

### HasPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) HasPeakDisconnectAction() bool`

HasPeakDisconnectAction returns a boolean if a field has been set.

### GetPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) GetPeakDisconnectTimeoutMinutes() int32`

GetPeakDisconnectTimeoutMinutes returns the PeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakDisconnectTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModel) GetPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakDisconnectTimeoutMinutesOk returns a tuple with the PeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) SetPeakDisconnectTimeoutMinutes(v int32)`

SetPeakDisconnectTimeoutMinutes sets PeakDisconnectTimeoutMinutes field to given value.

### HasPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) HasPeakDisconnectTimeoutMinutes() bool`

HasPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) GetPeakExtendedDisconnectAction() SessionChangeHostingAction`

GetPeakExtendedDisconnectAction returns the PeakExtendedDisconnectAction field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectActionOk

`func (o *DeliveryGroupDetailResponseModel) GetPeakExtendedDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetPeakExtendedDisconnectActionOk returns a tuple with the PeakExtendedDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) SetPeakExtendedDisconnectAction(v SessionChangeHostingAction)`

SetPeakExtendedDisconnectAction sets PeakExtendedDisconnectAction field to given value.

### HasPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModel) HasPeakExtendedDisconnectAction() bool`

HasPeakExtendedDisconnectAction returns a boolean if a field has been set.

### GetPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) GetPeakExtendedDisconnectTimeoutMinutes() int32`

GetPeakExtendedDisconnectTimeoutMinutes returns the PeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModel) GetPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the PeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) SetPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetPeakExtendedDisconnectTimeoutMinutes sets PeakExtendedDisconnectTimeoutMinutes field to given value.

### HasPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) HasPeakExtendedDisconnectTimeoutMinutes() bool`

HasPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModel) GetPeakLogOffAction() SessionChangeHostingAction`

GetPeakLogOffAction returns the PeakLogOffAction field if non-nil, zero value otherwise.

### GetPeakLogOffActionOk

`func (o *DeliveryGroupDetailResponseModel) GetPeakLogOffActionOk() (*SessionChangeHostingAction, bool)`

GetPeakLogOffActionOk returns a tuple with the PeakLogOffAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModel) SetPeakLogOffAction(v SessionChangeHostingAction)`

SetPeakLogOffAction sets PeakLogOffAction field to given value.

### HasPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModel) HasPeakLogOffAction() bool`

HasPeakLogOffAction returns a boolean if a field has been set.

### GetPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) GetPeakLogOffTimeoutMinutes() int32`

GetPeakLogOffTimeoutMinutes returns the PeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakLogOffTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModel) GetPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetPeakLogOffTimeoutMinutesOk returns a tuple with the PeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) SetPeakLogOffTimeoutMinutes(v int32)`

SetPeakLogOffTimeoutMinutes sets PeakLogOffTimeoutMinutes field to given value.

### HasPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModel) HasPeakLogOffTimeoutMinutes() bool`

HasPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### GetDisconnectPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) GetDisconnectPeakIdleSessionAfterSeconds() int32`

GetDisconnectPeakIdleSessionAfterSeconds returns the DisconnectPeakIdleSessionAfterSeconds field if non-nil, zero value otherwise.

### GetDisconnectPeakIdleSessionAfterSecondsOk

`func (o *DeliveryGroupDetailResponseModel) GetDisconnectPeakIdleSessionAfterSecondsOk() (*int32, bool)`

GetDisconnectPeakIdleSessionAfterSecondsOk returns a tuple with the DisconnectPeakIdleSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) SetDisconnectPeakIdleSessionAfterSeconds(v int32)`

SetDisconnectPeakIdleSessionAfterSeconds sets DisconnectPeakIdleSessionAfterSeconds field to given value.

### HasDisconnectPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) HasDisconnectPeakIdleSessionAfterSeconds() bool`

HasDisconnectPeakIdleSessionAfterSeconds returns a boolean if a field has been set.

### GetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) GetDisconnectOffPeakIdleSessionAfterSeconds() int32`

GetDisconnectOffPeakIdleSessionAfterSeconds returns the DisconnectOffPeakIdleSessionAfterSeconds field if non-nil, zero value otherwise.

### GetDisconnectOffPeakIdleSessionAfterSecondsOk

`func (o *DeliveryGroupDetailResponseModel) GetDisconnectOffPeakIdleSessionAfterSecondsOk() (*int32, bool)`

GetDisconnectOffPeakIdleSessionAfterSecondsOk returns a tuple with the DisconnectOffPeakIdleSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) SetDisconnectOffPeakIdleSessionAfterSeconds(v int32)`

SetDisconnectOffPeakIdleSessionAfterSeconds sets DisconnectOffPeakIdleSessionAfterSeconds field to given value.

### HasDisconnectOffPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) HasDisconnectOffPeakIdleSessionAfterSeconds() bool`

HasDisconnectOffPeakIdleSessionAfterSeconds returns a boolean if a field has been set.

### GetLogoffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) GetLogoffPeakDisconnectedSessionAfterSeconds() int32`

GetLogoffPeakDisconnectedSessionAfterSeconds returns the LogoffPeakDisconnectedSessionAfterSeconds field if non-nil, zero value otherwise.

### GetLogoffPeakDisconnectedSessionAfterSecondsOk

`func (o *DeliveryGroupDetailResponseModel) GetLogoffPeakDisconnectedSessionAfterSecondsOk() (*int32, bool)`

GetLogoffPeakDisconnectedSessionAfterSecondsOk returns a tuple with the LogoffPeakDisconnectedSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) SetLogoffPeakDisconnectedSessionAfterSeconds(v int32)`

SetLogoffPeakDisconnectedSessionAfterSeconds sets LogoffPeakDisconnectedSessionAfterSeconds field to given value.

### HasLogoffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) HasLogoffPeakDisconnectedSessionAfterSeconds() bool`

HasLogoffPeakDisconnectedSessionAfterSeconds returns a boolean if a field has been set.

### GetLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) GetLogoffOffPeakDisconnectedSessionAfterSeconds() int32`

GetLogoffOffPeakDisconnectedSessionAfterSeconds returns the LogoffOffPeakDisconnectedSessionAfterSeconds field if non-nil, zero value otherwise.

### GetLogoffOffPeakDisconnectedSessionAfterSecondsOk

`func (o *DeliveryGroupDetailResponseModel) GetLogoffOffPeakDisconnectedSessionAfterSecondsOk() (*int32, bool)`

GetLogoffOffPeakDisconnectedSessionAfterSecondsOk returns a tuple with the LogoffOffPeakDisconnectedSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) SetLogoffOffPeakDisconnectedSessionAfterSeconds(v int32)`

SetLogoffOffPeakDisconnectedSessionAfterSeconds sets LogoffOffPeakDisconnectedSessionAfterSeconds field to given value.

### HasLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModel) HasLogoffOffPeakDisconnectedSessionAfterSeconds() bool`

HasLogoffOffPeakDisconnectedSessionAfterSeconds returns a boolean if a field has been set.

### GetPrelaunchSettings

`func (o *DeliveryGroupDetailResponseModel) GetPrelaunchSettings() FastApplicationSettingsResponseModel`

GetPrelaunchSettings returns the PrelaunchSettings field if non-nil, zero value otherwise.

### GetPrelaunchSettingsOk

`func (o *DeliveryGroupDetailResponseModel) GetPrelaunchSettingsOk() (*FastApplicationSettingsResponseModel, bool)`

GetPrelaunchSettingsOk returns a tuple with the PrelaunchSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrelaunchSettings

`func (o *DeliveryGroupDetailResponseModel) SetPrelaunchSettings(v FastApplicationSettingsResponseModel)`

SetPrelaunchSettings sets PrelaunchSettings field to given value.

### HasPrelaunchSettings

`func (o *DeliveryGroupDetailResponseModel) HasPrelaunchSettings() bool`

HasPrelaunchSettings returns a boolean if a field has been set.

### GetPowerTimeSchemes

`func (o *DeliveryGroupDetailResponseModel) GetPowerTimeSchemes() []PowerTimeSchemeResponseModel`

GetPowerTimeSchemes returns the PowerTimeSchemes field if non-nil, zero value otherwise.

### GetPowerTimeSchemesOk

`func (o *DeliveryGroupDetailResponseModel) GetPowerTimeSchemesOk() (*[]PowerTimeSchemeResponseModel, bool)`

GetPowerTimeSchemesOk returns a tuple with the PowerTimeSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerTimeSchemes

`func (o *DeliveryGroupDetailResponseModel) SetPowerTimeSchemes(v []PowerTimeSchemeResponseModel)`

SetPowerTimeSchemes sets PowerTimeSchemes field to given value.

### HasPowerTimeSchemes

`func (o *DeliveryGroupDetailResponseModel) HasPowerTimeSchemes() bool`

HasPowerTimeSchemes returns a boolean if a field has been set.

### SetPowerTimeSchemesNil

`func (o *DeliveryGroupDetailResponseModel) SetPowerTimeSchemesNil(b bool)`

 SetPowerTimeSchemesNil sets the value for PowerTimeSchemes to be an explicit nil

### UnsetPowerTimeSchemes
`func (o *DeliveryGroupDetailResponseModel) UnsetPowerTimeSchemes()`

UnsetPowerTimeSchemes ensures that no value is present for PowerTimeSchemes, not even an explicit nil
### GetProtocolPriority

`func (o *DeliveryGroupDetailResponseModel) GetProtocolPriority() []ProtocolType`

GetProtocolPriority returns the ProtocolPriority field if non-nil, zero value otherwise.

### GetProtocolPriorityOk

`func (o *DeliveryGroupDetailResponseModel) GetProtocolPriorityOk() (*[]ProtocolType, bool)`

GetProtocolPriorityOk returns a tuple with the ProtocolPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPriority

`func (o *DeliveryGroupDetailResponseModel) SetProtocolPriority(v []ProtocolType)`

SetProtocolPriority sets ProtocolPriority field to given value.

### HasProtocolPriority

`func (o *DeliveryGroupDetailResponseModel) HasProtocolPriority() bool`

HasProtocolPriority returns a boolean if a field has been set.

### SetProtocolPriorityNil

`func (o *DeliveryGroupDetailResponseModel) SetProtocolPriorityNil(b bool)`

 SetProtocolPriorityNil sets the value for ProtocolPriority to be an explicit nil

### UnsetProtocolPriority
`func (o *DeliveryGroupDetailResponseModel) UnsetProtocolPriority()`

UnsetProtocolPriority ensures that no value is present for ProtocolPriority, not even an explicit nil
### GetRegisteredMachines

`func (o *DeliveryGroupDetailResponseModel) GetRegisteredMachines() int32`

GetRegisteredMachines returns the RegisteredMachines field if non-nil, zero value otherwise.

### GetRegisteredMachinesOk

`func (o *DeliveryGroupDetailResponseModel) GetRegisteredMachinesOk() (*int32, bool)`

GetRegisteredMachinesOk returns a tuple with the RegisteredMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredMachines

`func (o *DeliveryGroupDetailResponseModel) SetRegisteredMachines(v int32)`

SetRegisteredMachines sets RegisteredMachines field to given value.


### GetReuseMachinesWithoutShutdownInOutage

`func (o *DeliveryGroupDetailResponseModel) GetReuseMachinesWithoutShutdownInOutage() bool`

GetReuseMachinesWithoutShutdownInOutage returns the ReuseMachinesWithoutShutdownInOutage field if non-nil, zero value otherwise.

### GetReuseMachinesWithoutShutdownInOutageOk

`func (o *DeliveryGroupDetailResponseModel) GetReuseMachinesWithoutShutdownInOutageOk() (*bool, bool)`

GetReuseMachinesWithoutShutdownInOutageOk returns a tuple with the ReuseMachinesWithoutShutdownInOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseMachinesWithoutShutdownInOutage

`func (o *DeliveryGroupDetailResponseModel) SetReuseMachinesWithoutShutdownInOutage(v bool)`

SetReuseMachinesWithoutShutdownInOutage sets ReuseMachinesWithoutShutdownInOutage field to given value.

### HasReuseMachinesWithoutShutdownInOutage

`func (o *DeliveryGroupDetailResponseModel) HasReuseMachinesWithoutShutdownInOutage() bool`

HasReuseMachinesWithoutShutdownInOutage returns a boolean if a field has been set.

### GetSecureIcaRequired

`func (o *DeliveryGroupDetailResponseModel) GetSecureIcaRequired() bool`

GetSecureIcaRequired returns the SecureIcaRequired field if non-nil, zero value otherwise.

### GetSecureIcaRequiredOk

`func (o *DeliveryGroupDetailResponseModel) GetSecureIcaRequiredOk() (*bool, bool)`

GetSecureIcaRequiredOk returns a tuple with the SecureIcaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureIcaRequired

`func (o *DeliveryGroupDetailResponseModel) SetSecureIcaRequired(v bool)`

SetSecureIcaRequired sets SecureIcaRequired field to given value.


### GetSettlementPeriodBeforeAutoShutdownSeconds

`func (o *DeliveryGroupDetailResponseModel) GetSettlementPeriodBeforeAutoShutdownSeconds() int32`

GetSettlementPeriodBeforeAutoShutdownSeconds returns the SettlementPeriodBeforeAutoShutdownSeconds field if non-nil, zero value otherwise.

### GetSettlementPeriodBeforeAutoShutdownSecondsOk

`func (o *DeliveryGroupDetailResponseModel) GetSettlementPeriodBeforeAutoShutdownSecondsOk() (*int32, bool)`

GetSettlementPeriodBeforeAutoShutdownSecondsOk returns a tuple with the SettlementPeriodBeforeAutoShutdownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementPeriodBeforeAutoShutdownSeconds

`func (o *DeliveryGroupDetailResponseModel) SetSettlementPeriodBeforeAutoShutdownSeconds(v int32)`

SetSettlementPeriodBeforeAutoShutdownSeconds sets SettlementPeriodBeforeAutoShutdownSeconds field to given value.

### HasSettlementPeriodBeforeAutoShutdownSeconds

`func (o *DeliveryGroupDetailResponseModel) HasSettlementPeriodBeforeAutoShutdownSeconds() bool`

HasSettlementPeriodBeforeAutoShutdownSeconds returns a boolean if a field has been set.

### GetSettlementPeriodBeforeUseSeconds

`func (o *DeliveryGroupDetailResponseModel) GetSettlementPeriodBeforeUseSeconds() int32`

GetSettlementPeriodBeforeUseSeconds returns the SettlementPeriodBeforeUseSeconds field if non-nil, zero value otherwise.

### GetSettlementPeriodBeforeUseSecondsOk

`func (o *DeliveryGroupDetailResponseModel) GetSettlementPeriodBeforeUseSecondsOk() (*int32, bool)`

GetSettlementPeriodBeforeUseSecondsOk returns a tuple with the SettlementPeriodBeforeUseSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementPeriodBeforeUseSeconds

`func (o *DeliveryGroupDetailResponseModel) SetSettlementPeriodBeforeUseSeconds(v int32)`

SetSettlementPeriodBeforeUseSeconds sets SettlementPeriodBeforeUseSeconds field to given value.

### HasSettlementPeriodBeforeUseSeconds

`func (o *DeliveryGroupDetailResponseModel) HasSettlementPeriodBeforeUseSeconds() bool`

HasSettlementPeriodBeforeUseSeconds returns a boolean if a field has been set.

### GetShutdownDesktopsAfterUse

`func (o *DeliveryGroupDetailResponseModel) GetShutdownDesktopsAfterUse() bool`

GetShutdownDesktopsAfterUse returns the ShutdownDesktopsAfterUse field if non-nil, zero value otherwise.

### GetShutdownDesktopsAfterUseOk

`func (o *DeliveryGroupDetailResponseModel) GetShutdownDesktopsAfterUseOk() (*bool, bool)`

GetShutdownDesktopsAfterUseOk returns a tuple with the ShutdownDesktopsAfterUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDesktopsAfterUse

`func (o *DeliveryGroupDetailResponseModel) SetShutdownDesktopsAfterUse(v bool)`

SetShutdownDesktopsAfterUse sets ShutdownDesktopsAfterUse field to given value.

### HasShutdownDesktopsAfterUse

`func (o *DeliveryGroupDetailResponseModel) HasShutdownDesktopsAfterUse() bool`

HasShutdownDesktopsAfterUse returns a boolean if a field has been set.

### SetShutdownDesktopsAfterUseNil

`func (o *DeliveryGroupDetailResponseModel) SetShutdownDesktopsAfterUseNil(b bool)`

 SetShutdownDesktopsAfterUseNil sets the value for ShutdownDesktopsAfterUse to be an explicit nil

### UnsetShutdownDesktopsAfterUse
`func (o *DeliveryGroupDetailResponseModel) UnsetShutdownDesktopsAfterUse()`

UnsetShutdownDesktopsAfterUse ensures that no value is present for ShutdownDesktopsAfterUse, not even an explicit nil
### GetSimpleAccessPolicy

`func (o *DeliveryGroupDetailResponseModel) GetSimpleAccessPolicy() SimplifiedAccessPolicyResponseModel`

GetSimpleAccessPolicy returns the SimpleAccessPolicy field if non-nil, zero value otherwise.

### GetSimpleAccessPolicyOk

`func (o *DeliveryGroupDetailResponseModel) GetSimpleAccessPolicyOk() (*SimplifiedAccessPolicyResponseModel, bool)`

GetSimpleAccessPolicyOk returns a tuple with the SimpleAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAccessPolicy

`func (o *DeliveryGroupDetailResponseModel) SetSimpleAccessPolicy(v SimplifiedAccessPolicyResponseModel)`

SetSimpleAccessPolicy sets SimpleAccessPolicy field to given value.

### HasSimpleAccessPolicy

`func (o *DeliveryGroupDetailResponseModel) HasSimpleAccessPolicy() bool`

HasSimpleAccessPolicy returns a boolean if a field has been set.

### GetAdvancedAccessPolicy

`func (o *DeliveryGroupDetailResponseModel) GetAdvancedAccessPolicy() []AdvancedAccessPolicyResponseModel`

GetAdvancedAccessPolicy returns the AdvancedAccessPolicy field if non-nil, zero value otherwise.

### GetAdvancedAccessPolicyOk

`func (o *DeliveryGroupDetailResponseModel) GetAdvancedAccessPolicyOk() (*[]AdvancedAccessPolicyResponseModel, bool)`

GetAdvancedAccessPolicyOk returns a tuple with the AdvancedAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedAccessPolicy

`func (o *DeliveryGroupDetailResponseModel) SetAdvancedAccessPolicy(v []AdvancedAccessPolicyResponseModel)`

SetAdvancedAccessPolicy sets AdvancedAccessPolicy field to given value.

### HasAdvancedAccessPolicy

`func (o *DeliveryGroupDetailResponseModel) HasAdvancedAccessPolicy() bool`

HasAdvancedAccessPolicy returns a boolean if a field has been set.

### SetAdvancedAccessPolicyNil

`func (o *DeliveryGroupDetailResponseModel) SetAdvancedAccessPolicyNil(b bool)`

 SetAdvancedAccessPolicyNil sets the value for AdvancedAccessPolicy to be an explicit nil

### UnsetAdvancedAccessPolicy
`func (o *DeliveryGroupDetailResponseModel) UnsetAdvancedAccessPolicy()`

UnsetAdvancedAccessPolicy ensures that no value is present for AdvancedAccessPolicy, not even an explicit nil
### GetStoreFrontServersForHostedReceiver

`func (o *DeliveryGroupDetailResponseModel) GetStoreFrontServersForHostedReceiver() []StoreFrontServerResponseModel`

GetStoreFrontServersForHostedReceiver returns the StoreFrontServersForHostedReceiver field if non-nil, zero value otherwise.

### GetStoreFrontServersForHostedReceiverOk

`func (o *DeliveryGroupDetailResponseModel) GetStoreFrontServersForHostedReceiverOk() (*[]StoreFrontServerResponseModel, bool)`

GetStoreFrontServersForHostedReceiverOk returns a tuple with the StoreFrontServersForHostedReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFrontServersForHostedReceiver

`func (o *DeliveryGroupDetailResponseModel) SetStoreFrontServersForHostedReceiver(v []StoreFrontServerResponseModel)`

SetStoreFrontServersForHostedReceiver sets StoreFrontServersForHostedReceiver field to given value.

### HasStoreFrontServersForHostedReceiver

`func (o *DeliveryGroupDetailResponseModel) HasStoreFrontServersForHostedReceiver() bool`

HasStoreFrontServersForHostedReceiver returns a boolean if a field has been set.

### SetStoreFrontServersForHostedReceiverNil

`func (o *DeliveryGroupDetailResponseModel) SetStoreFrontServersForHostedReceiverNil(b bool)`

 SetStoreFrontServersForHostedReceiverNil sets the value for StoreFrontServersForHostedReceiver to be an explicit nil

### UnsetStoreFrontServersForHostedReceiver
`func (o *DeliveryGroupDetailResponseModel) UnsetStoreFrontServersForHostedReceiver()`

UnsetStoreFrontServersForHostedReceiver ensures that no value is present for StoreFrontServersForHostedReceiver, not even an explicit nil
### GetTimeZone

`func (o *DeliveryGroupDetailResponseModel) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *DeliveryGroupDetailResponseModel) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *DeliveryGroupDetailResponseModel) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *DeliveryGroupDetailResponseModel) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *DeliveryGroupDetailResponseModel) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *DeliveryGroupDetailResponseModel) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetTotalApplicationGroups

`func (o *DeliveryGroupDetailResponseModel) GetTotalApplicationGroups() int32`

GetTotalApplicationGroups returns the TotalApplicationGroups field if non-nil, zero value otherwise.

### GetTotalApplicationGroupsOk

`func (o *DeliveryGroupDetailResponseModel) GetTotalApplicationGroupsOk() (*int32, bool)`

GetTotalApplicationGroupsOk returns a tuple with the TotalApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplicationGroups

`func (o *DeliveryGroupDetailResponseModel) SetTotalApplicationGroups(v int32)`

SetTotalApplicationGroups sets TotalApplicationGroups field to given value.


### GetTotalMachines

`func (o *DeliveryGroupDetailResponseModel) GetTotalMachines() int32`

GetTotalMachines returns the TotalMachines field if non-nil, zero value otherwise.

### GetTotalMachinesOk

`func (o *DeliveryGroupDetailResponseModel) GetTotalMachinesOk() (*int32, bool)`

GetTotalMachinesOk returns a tuple with the TotalMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachines

`func (o *DeliveryGroupDetailResponseModel) SetTotalMachines(v int32)`

SetTotalMachines sets TotalMachines field to given value.


### GetTurnOnAddedMachines

`func (o *DeliveryGroupDetailResponseModel) GetTurnOnAddedMachines() bool`

GetTurnOnAddedMachines returns the TurnOnAddedMachines field if non-nil, zero value otherwise.

### GetTurnOnAddedMachinesOk

`func (o *DeliveryGroupDetailResponseModel) GetTurnOnAddedMachinesOk() (*bool, bool)`

GetTurnOnAddedMachinesOk returns a tuple with the TurnOnAddedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOnAddedMachines

`func (o *DeliveryGroupDetailResponseModel) SetTurnOnAddedMachines(v bool)`

SetTurnOnAddedMachines sets TurnOnAddedMachines field to given value.

### HasTurnOnAddedMachines

`func (o *DeliveryGroupDetailResponseModel) HasTurnOnAddedMachines() bool`

HasTurnOnAddedMachines returns a boolean if a field has been set.

### SetTurnOnAddedMachinesNil

`func (o *DeliveryGroupDetailResponseModel) SetTurnOnAddedMachinesNil(b bool)`

 SetTurnOnAddedMachinesNil sets the value for TurnOnAddedMachines to be an explicit nil

### UnsetTurnOnAddedMachines
`func (o *DeliveryGroupDetailResponseModel) UnsetTurnOnAddedMachines()`

UnsetTurnOnAddedMachines ensures that no value is present for TurnOnAddedMachines, not even an explicit nil
### GetUnassignedMachines

`func (o *DeliveryGroupDetailResponseModel) GetUnassignedMachines() int32`

GetUnassignedMachines returns the UnassignedMachines field if non-nil, zero value otherwise.

### GetUnassignedMachinesOk

`func (o *DeliveryGroupDetailResponseModel) GetUnassignedMachinesOk() (*int32, bool)`

GetUnassignedMachinesOk returns a tuple with the UnassignedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedMachines

`func (o *DeliveryGroupDetailResponseModel) SetUnassignedMachines(v int32)`

SetUnassignedMachines sets UnassignedMachines field to given value.

### HasUnassignedMachines

`func (o *DeliveryGroupDetailResponseModel) HasUnassignedMachines() bool`

HasUnassignedMachines returns a boolean if a field has been set.

### SetUnassignedMachinesNil

`func (o *DeliveryGroupDetailResponseModel) SetUnassignedMachinesNil(b bool)`

 SetUnassignedMachinesNil sets the value for UnassignedMachines to be an explicit nil

### UnsetUnassignedMachines
`func (o *DeliveryGroupDetailResponseModel) UnsetUnassignedMachines()`

UnsetUnassignedMachines ensures that no value is present for UnassignedMachines, not even an explicit nil
### GetUsedByGroupPolicy

`func (o *DeliveryGroupDetailResponseModel) GetUsedByGroupPolicy() bool`

GetUsedByGroupPolicy returns the UsedByGroupPolicy field if non-nil, zero value otherwise.

### GetUsedByGroupPolicyOk

`func (o *DeliveryGroupDetailResponseModel) GetUsedByGroupPolicyOk() (*bool, bool)`

GetUsedByGroupPolicyOk returns a tuple with the UsedByGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedByGroupPolicy

`func (o *DeliveryGroupDetailResponseModel) SetUsedByGroupPolicy(v bool)`

SetUsedByGroupPolicy sets UsedByGroupPolicy field to given value.

### HasUsedByGroupPolicy

`func (o *DeliveryGroupDetailResponseModel) HasUsedByGroupPolicy() bool`

HasUsedByGroupPolicy returns a boolean if a field has been set.

### SetUsedByGroupPolicyNil

`func (o *DeliveryGroupDetailResponseModel) SetUsedByGroupPolicyNil(b bool)`

 SetUsedByGroupPolicyNil sets the value for UsedByGroupPolicy to be an explicit nil

### UnsetUsedByGroupPolicy
`func (o *DeliveryGroupDetailResponseModel) UnsetUsedByGroupPolicy()`

UnsetUsedByGroupPolicy ensures that no value is present for UsedByGroupPolicy, not even an explicit nil
### GetVdaVersions

`func (o *DeliveryGroupDetailResponseModel) GetVdaVersions() []NameValueIntPairModel`

GetVdaVersions returns the VdaVersions field if non-nil, zero value otherwise.

### GetVdaVersionsOk

`func (o *DeliveryGroupDetailResponseModel) GetVdaVersionsOk() (*[]NameValueIntPairModel, bool)`

GetVdaVersionsOk returns a tuple with the VdaVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaVersions

`func (o *DeliveryGroupDetailResponseModel) SetVdaVersions(v []NameValueIntPairModel)`

SetVdaVersions sets VdaVersions field to given value.

### HasVdaVersions

`func (o *DeliveryGroupDetailResponseModel) HasVdaVersions() bool`

HasVdaVersions returns a boolean if a field has been set.

### SetVdaVersionsNil

`func (o *DeliveryGroupDetailResponseModel) SetVdaVersionsNil(b bool)`

 SetVdaVersionsNil sets the value for VdaVersions to be an explicit nil

### UnsetVdaVersions
`func (o *DeliveryGroupDetailResponseModel) UnsetVdaVersions()`

UnsetVdaVersions ensures that no value is present for VdaVersions, not even an explicit nil
### GetZonePreferences

`func (o *DeliveryGroupDetailResponseModel) GetZonePreferences() []ZonePreference`

GetZonePreferences returns the ZonePreferences field if non-nil, zero value otherwise.

### GetZonePreferencesOk

`func (o *DeliveryGroupDetailResponseModel) GetZonePreferencesOk() (*[]ZonePreference, bool)`

GetZonePreferencesOk returns a tuple with the ZonePreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePreferences

`func (o *DeliveryGroupDetailResponseModel) SetZonePreferences(v []ZonePreference)`

SetZonePreferences sets ZonePreferences field to given value.

### HasZonePreferences

`func (o *DeliveryGroupDetailResponseModel) HasZonePreferences() bool`

HasZonePreferences returns a boolean if a field has been set.

### SetZonePreferencesNil

`func (o *DeliveryGroupDetailResponseModel) SetZonePreferencesNil(b bool)`

 SetZonePreferencesNil sets the value for ZonePreferences to be an explicit nil

### UnsetZonePreferences
`func (o *DeliveryGroupDetailResponseModel) UnsetZonePreferences()`

UnsetZonePreferences ensures that no value is present for ZonePreferences, not even an explicit nil
### GetLimitSecondsToForceLogOffUserDuringPeak

`func (o *DeliveryGroupDetailResponseModel) GetLimitSecondsToForceLogOffUserDuringPeak() int32`

GetLimitSecondsToForceLogOffUserDuringPeak returns the LimitSecondsToForceLogOffUserDuringPeak field if non-nil, zero value otherwise.

### GetLimitSecondsToForceLogOffUserDuringPeakOk

`func (o *DeliveryGroupDetailResponseModel) GetLimitSecondsToForceLogOffUserDuringPeakOk() (*int32, bool)`

GetLimitSecondsToForceLogOffUserDuringPeakOk returns a tuple with the LimitSecondsToForceLogOffUserDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSecondsToForceLogOffUserDuringPeak

`func (o *DeliveryGroupDetailResponseModel) SetLimitSecondsToForceLogOffUserDuringPeak(v int32)`

SetLimitSecondsToForceLogOffUserDuringPeak sets LimitSecondsToForceLogOffUserDuringPeak field to given value.

### HasLimitSecondsToForceLogOffUserDuringPeak

`func (o *DeliveryGroupDetailResponseModel) HasLimitSecondsToForceLogOffUserDuringPeak() bool`

HasLimitSecondsToForceLogOffUserDuringPeak returns a boolean if a field has been set.

### SetLimitSecondsToForceLogOffUserDuringPeakNil

`func (o *DeliveryGroupDetailResponseModel) SetLimitSecondsToForceLogOffUserDuringPeakNil(b bool)`

 SetLimitSecondsToForceLogOffUserDuringPeakNil sets the value for LimitSecondsToForceLogOffUserDuringPeak to be an explicit nil

### UnsetLimitSecondsToForceLogOffUserDuringPeak
`func (o *DeliveryGroupDetailResponseModel) UnsetLimitSecondsToForceLogOffUserDuringPeak()`

UnsetLimitSecondsToForceLogOffUserDuringPeak ensures that no value is present for LimitSecondsToForceLogOffUserDuringPeak, not even an explicit nil
### GetLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *DeliveryGroupDetailResponseModel) GetLimitSecondsToForceLogOffUserDuringOffPeak() int32`

GetLimitSecondsToForceLogOffUserDuringOffPeak returns the LimitSecondsToForceLogOffUserDuringOffPeak field if non-nil, zero value otherwise.

### GetLimitSecondsToForceLogOffUserDuringOffPeakOk

`func (o *DeliveryGroupDetailResponseModel) GetLimitSecondsToForceLogOffUserDuringOffPeakOk() (*int32, bool)`

GetLimitSecondsToForceLogOffUserDuringOffPeakOk returns a tuple with the LimitSecondsToForceLogOffUserDuringOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *DeliveryGroupDetailResponseModel) SetLimitSecondsToForceLogOffUserDuringOffPeak(v int32)`

SetLimitSecondsToForceLogOffUserDuringOffPeak sets LimitSecondsToForceLogOffUserDuringOffPeak field to given value.

### HasLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *DeliveryGroupDetailResponseModel) HasLimitSecondsToForceLogOffUserDuringOffPeak() bool`

HasLimitSecondsToForceLogOffUserDuringOffPeak returns a boolean if a field has been set.

### SetLimitSecondsToForceLogOffUserDuringOffPeakNil

`func (o *DeliveryGroupDetailResponseModel) SetLimitSecondsToForceLogOffUserDuringOffPeakNil(b bool)`

 SetLimitSecondsToForceLogOffUserDuringOffPeakNil sets the value for LimitSecondsToForceLogOffUserDuringOffPeak to be an explicit nil

### UnsetLimitSecondsToForceLogOffUserDuringOffPeak
`func (o *DeliveryGroupDetailResponseModel) UnsetLimitSecondsToForceLogOffUserDuringOffPeak()`

UnsetLimitSecondsToForceLogOffUserDuringOffPeak ensures that no value is present for LimitSecondsToForceLogOffUserDuringOffPeak, not even an explicit nil
### GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *DeliveryGroupDetailResponseModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak() int32`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak returns the RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field if non-nil, zero value otherwise.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk

`func (o *DeliveryGroupDetailResponseModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk() (*int32, bool)`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk returns a tuple with the RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *DeliveryGroupDetailResponseModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak(v int32)`

SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak sets RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field to given value.

### HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *DeliveryGroupDetailResponseModel) HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak() bool`

HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak returns a boolean if a field has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakNil

`func (o *DeliveryGroupDetailResponseModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakNil(b bool)`

 SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakNil sets the value for RestrictAutoscaleMinIdleUntaggedPercentDuringPeak to be an explicit nil

### UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak
`func (o *DeliveryGroupDetailResponseModel) UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak()`

UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak ensures that no value is present for RestrictAutoscaleMinIdleUntaggedPercentDuringPeak, not even an explicit nil
### GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *DeliveryGroupDetailResponseModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak() int32`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak returns the RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field if non-nil, zero value otherwise.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk

`func (o *DeliveryGroupDetailResponseModel) GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk() (*int32, bool)`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk returns a tuple with the RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *DeliveryGroupDetailResponseModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak(v int32)`

SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak sets RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field to given value.

### HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *DeliveryGroupDetailResponseModel) HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak() bool`

HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak returns a boolean if a field has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakNil

`func (o *DeliveryGroupDetailResponseModel) SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakNil(b bool)`

 SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakNil sets the value for RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak to be an explicit nil

### UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak
`func (o *DeliveryGroupDetailResponseModel) UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak()`

UnsetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak ensures that no value is present for RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak, not even an explicit nil
### GetLogOffWarningMessage

`func (o *DeliveryGroupDetailResponseModel) GetLogOffWarningMessage() string`

GetLogOffWarningMessage returns the LogOffWarningMessage field if non-nil, zero value otherwise.

### GetLogOffWarningMessageOk

`func (o *DeliveryGroupDetailResponseModel) GetLogOffWarningMessageOk() (*string, bool)`

GetLogOffWarningMessageOk returns a tuple with the LogOffWarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOffWarningMessage

`func (o *DeliveryGroupDetailResponseModel) SetLogOffWarningMessage(v string)`

SetLogOffWarningMessage sets LogOffWarningMessage field to given value.

### HasLogOffWarningMessage

`func (o *DeliveryGroupDetailResponseModel) HasLogOffWarningMessage() bool`

HasLogOffWarningMessage returns a boolean if a field has been set.

### SetLogOffWarningMessageNil

`func (o *DeliveryGroupDetailResponseModel) SetLogOffWarningMessageNil(b bool)`

 SetLogOffWarningMessageNil sets the value for LogOffWarningMessage to be an explicit nil

### UnsetLogOffWarningMessage
`func (o *DeliveryGroupDetailResponseModel) UnsetLogOffWarningMessage()`

UnsetLogOffWarningMessage ensures that no value is present for LogOffWarningMessage, not even an explicit nil
### GetLogOffWarningTitle

`func (o *DeliveryGroupDetailResponseModel) GetLogOffWarningTitle() string`

GetLogOffWarningTitle returns the LogOffWarningTitle field if non-nil, zero value otherwise.

### GetLogOffWarningTitleOk

`func (o *DeliveryGroupDetailResponseModel) GetLogOffWarningTitleOk() (*string, bool)`

GetLogOffWarningTitleOk returns a tuple with the LogOffWarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOffWarningTitle

`func (o *DeliveryGroupDetailResponseModel) SetLogOffWarningTitle(v string)`

SetLogOffWarningTitle sets LogOffWarningTitle field to given value.

### HasLogOffWarningTitle

`func (o *DeliveryGroupDetailResponseModel) HasLogOffWarningTitle() bool`

HasLogOffWarningTitle returns a boolean if a field has been set.

### SetLogOffWarningTitleNil

`func (o *DeliveryGroupDetailResponseModel) SetLogOffWarningTitleNil(b bool)`

 SetLogOffWarningTitleNil sets the value for LogOffWarningTitle to be an explicit nil

### UnsetLogOffWarningTitle
`func (o *DeliveryGroupDetailResponseModel) UnsetLogOffWarningTitle()`

UnsetLogOffWarningTitle ensures that no value is present for LogOffWarningTitle, not even an explicit nil
### GetAutoscaleLogOffReminderEnabled

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderEnabled() bool`

GetAutoscaleLogOffReminderEnabled returns the AutoscaleLogOffReminderEnabled field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderEnabledOk

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderEnabledOk() (*bool, bool)`

GetAutoscaleLogOffReminderEnabledOk returns a tuple with the AutoscaleLogOffReminderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderEnabled

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderEnabled(v bool)`

SetAutoscaleLogOffReminderEnabled sets AutoscaleLogOffReminderEnabled field to given value.

### HasAutoscaleLogOffReminderEnabled

`func (o *DeliveryGroupDetailResponseModel) HasAutoscaleLogOffReminderEnabled() bool`

HasAutoscaleLogOffReminderEnabled returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderEnabledNil

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderEnabledNil(b bool)`

 SetAutoscaleLogOffReminderEnabledNil sets the value for AutoscaleLogOffReminderEnabled to be an explicit nil

### UnsetAutoscaleLogOffReminderEnabled
`func (o *DeliveryGroupDetailResponseModel) UnsetAutoscaleLogOffReminderEnabled()`

UnsetAutoscaleLogOffReminderEnabled ensures that no value is present for AutoscaleLogOffReminderEnabled, not even an explicit nil
### GetAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderIntervalSecondsOffPeak() int32`

GetAutoscaleLogOffReminderIntervalSecondsOffPeak returns the AutoscaleLogOffReminderIntervalSecondsOffPeak field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk() (*int32, bool)`

GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk returns a tuple with the AutoscaleLogOffReminderIntervalSecondsOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderIntervalSecondsOffPeak(v int32)`

SetAutoscaleLogOffReminderIntervalSecondsOffPeak sets AutoscaleLogOffReminderIntervalSecondsOffPeak field to given value.

### HasAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *DeliveryGroupDetailResponseModel) HasAutoscaleLogOffReminderIntervalSecondsOffPeak() bool`

HasAutoscaleLogOffReminderIntervalSecondsOffPeak returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderIntervalSecondsOffPeakNil

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderIntervalSecondsOffPeakNil(b bool)`

 SetAutoscaleLogOffReminderIntervalSecondsOffPeakNil sets the value for AutoscaleLogOffReminderIntervalSecondsOffPeak to be an explicit nil

### UnsetAutoscaleLogOffReminderIntervalSecondsOffPeak
`func (o *DeliveryGroupDetailResponseModel) UnsetAutoscaleLogOffReminderIntervalSecondsOffPeak()`

UnsetAutoscaleLogOffReminderIntervalSecondsOffPeak ensures that no value is present for AutoscaleLogOffReminderIntervalSecondsOffPeak, not even an explicit nil
### GetAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderIntervalSecondsPeak() int32`

GetAutoscaleLogOffReminderIntervalSecondsPeak returns the AutoscaleLogOffReminderIntervalSecondsPeak field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderIntervalSecondsPeakOk

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderIntervalSecondsPeakOk() (*int32, bool)`

GetAutoscaleLogOffReminderIntervalSecondsPeakOk returns a tuple with the AutoscaleLogOffReminderIntervalSecondsPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderIntervalSecondsPeak(v int32)`

SetAutoscaleLogOffReminderIntervalSecondsPeak sets AutoscaleLogOffReminderIntervalSecondsPeak field to given value.

### HasAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *DeliveryGroupDetailResponseModel) HasAutoscaleLogOffReminderIntervalSecondsPeak() bool`

HasAutoscaleLogOffReminderIntervalSecondsPeak returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderIntervalSecondsPeakNil

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderIntervalSecondsPeakNil(b bool)`

 SetAutoscaleLogOffReminderIntervalSecondsPeakNil sets the value for AutoscaleLogOffReminderIntervalSecondsPeak to be an explicit nil

### UnsetAutoscaleLogOffReminderIntervalSecondsPeak
`func (o *DeliveryGroupDetailResponseModel) UnsetAutoscaleLogOffReminderIntervalSecondsPeak()`

UnsetAutoscaleLogOffReminderIntervalSecondsPeak ensures that no value is present for AutoscaleLogOffReminderIntervalSecondsPeak, not even an explicit nil
### GetAutoscaleLogOffReminderMessage

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderMessage() string`

GetAutoscaleLogOffReminderMessage returns the AutoscaleLogOffReminderMessage field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderMessageOk

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderMessageOk() (*string, bool)`

GetAutoscaleLogOffReminderMessageOk returns a tuple with the AutoscaleLogOffReminderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderMessage

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderMessage(v string)`

SetAutoscaleLogOffReminderMessage sets AutoscaleLogOffReminderMessage field to given value.

### HasAutoscaleLogOffReminderMessage

`func (o *DeliveryGroupDetailResponseModel) HasAutoscaleLogOffReminderMessage() bool`

HasAutoscaleLogOffReminderMessage returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderMessageNil

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderMessageNil(b bool)`

 SetAutoscaleLogOffReminderMessageNil sets the value for AutoscaleLogOffReminderMessage to be an explicit nil

### UnsetAutoscaleLogOffReminderMessage
`func (o *DeliveryGroupDetailResponseModel) UnsetAutoscaleLogOffReminderMessage()`

UnsetAutoscaleLogOffReminderMessage ensures that no value is present for AutoscaleLogOffReminderMessage, not even an explicit nil
### GetAutoscaleLogOffReminderTitle

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderTitle() string`

GetAutoscaleLogOffReminderTitle returns the AutoscaleLogOffReminderTitle field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderTitleOk

`func (o *DeliveryGroupDetailResponseModel) GetAutoscaleLogOffReminderTitleOk() (*string, bool)`

GetAutoscaleLogOffReminderTitleOk returns a tuple with the AutoscaleLogOffReminderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderTitle

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderTitle(v string)`

SetAutoscaleLogOffReminderTitle sets AutoscaleLogOffReminderTitle field to given value.

### HasAutoscaleLogOffReminderTitle

`func (o *DeliveryGroupDetailResponseModel) HasAutoscaleLogOffReminderTitle() bool`

HasAutoscaleLogOffReminderTitle returns a boolean if a field has been set.

### SetAutoscaleLogOffReminderTitleNil

`func (o *DeliveryGroupDetailResponseModel) SetAutoscaleLogOffReminderTitleNil(b bool)`

 SetAutoscaleLogOffReminderTitleNil sets the value for AutoscaleLogOffReminderTitle to be an explicit nil

### UnsetAutoscaleLogOffReminderTitle
`func (o *DeliveryGroupDetailResponseModel) UnsetAutoscaleLogOffReminderTitle()`

UnsetAutoscaleLogOffReminderTitle ensures that no value is present for AutoscaleLogOffReminderTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


