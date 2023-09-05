# DeliveryGroupDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAccessPolicy** | Pointer to [**DeliveryGroupDetailResponseModelAllOfAppAccessPolicy**](DeliveryGroupDetailResponseModelAllOfAppAccessPolicy.md) |  | [optional] 
**AppProtectionKeyLoggingRequired** | Pointer to **bool** | Specifies whether key logging app protection is required. | [optional] 
**AppProtectionScreenCaptureRequired** | Pointer to **bool** | Specifies whether screen capture app protection is required. | [optional] 
**AutomaticPowerOnForAssigned** | Pointer to **bool** | Whether assigned (Private) machines in the delivery group should be automatically powered-on at the start of peak time periods. | [optional] 
**AutomaticPowerOnForAssignedDuringPeak** | Pointer to **bool** | Whether assigned (Private) machines in the delivery group should be automatically powered-on throughout peak time periods. | [optional] 
**AutoScaleEnabled** | Pointer to **bool** | Indicates whether auto-scale is enabled for the delivery group. | [optional] 
**RestrictAutoscaleTag** | Pointer to [**DeliveryGroupDetailResponseModelAllOfRestrictAutoscaleTag**](DeliveryGroupDetailResponseModelAllOfRestrictAutoscaleTag.md) |  | [optional] 
**ColorDepth** | Pointer to [**ColorDepth**](ColorDepth.md) |  | [optional] 
**DefaultDesktopIconId** | Pointer to **string** | Default icon to use for desktops published from the delivery group. was IconUid | [optional] 
**DefaultDesktopPublishedName** | Pointer to **string** | Default published name to use for desktops published from the delivery group. Change: Add | [optional] 
**DesktopsInUse** | **int32** | Number of machines in the delivery group which are currently in-use. | 
**DesktopsNeverRegistered** | **int32** | Number of machines in the delivery group which have never successfully registered. | 
**DesktopsPreparing** | **int32** | Number of machines in the delivery group which are preparing for a connection. | 
**HdxSslEnabled** | Pointer to **bool** | Whether connections to machines in the delivery group will use SSL. CHANGE: add: | [optional] 
**IsPowerManaged** | Pointer to **bool** | Indicates whether the machines in the delivery group are power-managed. NOTE: I used to think that MachineType&#x3D;&#x3D;Virtual meant the same thing as \&quot;power-managed\&quot;; however that&#39;s not the case.  A machine is power- managed if it is Virtual OR if it is RemotePC with a hypervisor connection (which will still have MachineType&#x3D;&#x3D;Physical). | [optional] 
**LingerSettings** | Pointer to [**DeliveryGroupDetailResponseModelAllOfLingerSettings**](DeliveryGroupDetailResponseModelAllOfLingerSettings.md) |  | [optional] 
**PowerOffDelayMinutes** | Pointer to **int32** | Delay before machines are powered off, when scaling down.  Specified in minutes.  Applies only to multi-session machines. | [optional] 
**MachineCost** | Pointer to **float64** | Indicates the estimated per-hour cost for machines in the delivery group, as set by the administrator. | [optional] 
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
**PrelaunchSettings** | Pointer to [**DeliveryGroupDetailResponseModelAllOfPrelaunchSettings**](DeliveryGroupDetailResponseModelAllOfPrelaunchSettings.md) |  | [optional] 
**PowerTimeSchemes** | Pointer to [**[]PowerTimeSchemeResponseModel**](PowerTimeSchemeResponseModel.md) | Power management time schemes.  No two schemes will cover the same day of the week. | [optional] 
**ProtocolPriority** | Pointer to [**[]ProtocolType**](ProtocolType.md) | A list of protocols in the order in which they will be attempted for use during connection. | [optional] 
**RegisteredMachines** | **int32** | Number of machines in the delivery group that have registered. | 
**ReuseMachinesWithoutShutdownInOutage** | Pointer to **bool** | Whether machines will be reused without a shutdown in case of an outage. | [optional] 
**SecureIcaRequired** | **bool** | Whether HDX connections to machines in the delivery group require the use of the SecureICA protocol. | 
**SettlementPeriodBeforeAutoShutdownSeconds** | Pointer to **int32** | Time after a session ends during which automatic shutdown requests (for example, shutdown after use, idle pool management) are deferred. Any outstanding shutdown request takes effect after the settlement period expires. This is typically used to configure time to allow logoff scripts to complete. | [optional] 
**SettlementPeriodBeforeUseSeconds** | Pointer to **int32** | Idle period before a machine can be selected to host a new session after registration or the end of a previous session. This is typically used to allow a machine to become idle following processing associated with start-up or logoff actions. A machine may still be selected during the idle period if no other machine is available for immediate use. | [optional] 
**ShutdownDesktopsAfterUse** | Pointer to **bool** | Whether machines in this delivery group should be automatically shut down when each user session completes. | [optional] 
**SimpleAccessPolicy** | Pointer to [**DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy**](DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy.md) |  | [optional] 
**StoreFrontServersForHostedReceiver** | Pointer to [**[]StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md) | List of StoreFront server addresses to configure within hosted receivers that are delivered from the delivery group. CHANGE: was public string[] StoreFronts | [optional] 
**TimeZone** | Pointer to **string** | The time zone in which this delivery group&#39;s machines reside. | [optional] 
**TotalApplicationGroups** | **int32** | Total number of application groups associated with the delivery group. | 
**TotalMachines** | **int32** | Total number of machines in the delivery group. | 
**TurnOnAddedMachines** | Pointer to **bool** | Whether to attempt to power on machines when they are added to the delivery group. was TurnOnAddedMachine | [optional] 
**UnassignedMachines** | Pointer to **int32** | Number of unassigned machines in the delivery group. | [optional] 
**UsedByGroupPolicy** | Pointer to **bool** | Check if desktop group is used by group policy. | [optional] 
**VdaVersions** | Pointer to [**[]NameValueIntPairModel**](NameValueIntPairModel.md) | List of the versions, and count of each version, of VDAs within the delivery group. CHANGE: was public IDictionary&lt;string,int&gt; VdaVersions { get; set; } | [optional] 
**ZonePreferences** | Pointer to [**[]ZonePreference**](ZonePreference.md) | Ordered list of zone preferences to be applied when launching resources from this delivery group. | [optional] 
**LimitSecondsToForceLogOffUserDuringPeak** | Pointer to **int32** | Represents the number of seconds that must elapsed before the active sessions on the draining machines belonging to the delivery group are logged off, during peak time. | [optional] 
**LimitSecondsToForceLogOffUserDuringOffPeak** | Pointer to **int32** | Represents the number of seconds that must elapsed before the active sessions on the draining machines belonging to the delivery group are logged off, during off-peak. | [optional] 
**RestrictAutoscaleMinIdleUntaggedPercentDuringPeak** | Pointer to **int32** | Represents the percentage that the number of untagged single-session machines in an idle state, or for multi-session machines, the untagged available load capacity must fall below before Autoscale powers on and manages &#39;tagged&#39; machines, as per policy, in peak time. If the number of untagged machines in an idle state, or the untagged available load capacity goes above this threshold value, Autoscale will attempt to shut down &#39;tagged&#39; machines. | [optional] 
**RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak** | Pointer to **int32** | Represents the percentage that the number of untagged single-session machines in an idle state, or for multi-session machines, the untagged available load capacity must fall below before Autoscale powers on and manages &#39;tagged&#39; machines, as per policy, in off-peak. If the number of untagged machines in an idle state, or the untagged available load capacity goes above this threshold value, Autoscale will attempt to shut down &#39;tagged&#39; machines. | [optional] 
**LogOffWarningMessage** | Pointer to **string** | The warning message to display to users in active sessions prior to logging off users, whether in peak time or off-peak. | [optional] 
**LogOffWarningTitle** | Pointer to **string** | The title of the warning message dialog. | [optional] 
**AutoscaleLogOffReminderEnabled** | Pointer to **bool** | Boolean value indicating whether the warning messages should be sent on an interval to nudge a logoff should be sent on an interval when autoscale is enabled. | [optional] 
**AutoscaleLogOffReminderIntervalSecondsOffPeak** | Pointer to **int32** | Represents the time interval at which messages are  sent to the user during off peak time when Autoscale is enabled. This message will nudge users to log off instead of forcibly logging them off. | [optional] 
**AutoscaleLogOffReminderIntervalSecondsPeak** | Pointer to **int32** | Represents the time interval at which messages are  sent to the user during peak time when autoscale is enabled. This message will nudge users to log off instead of forcibly logging them off. | [optional] 
**AutoscaleLogOffReminderMessage** | Pointer to **string** | Notification message to display to users in active sessions belonging to machines needed by Autoscale for shutdown. | [optional] 
**AutoscaleLogOffReminderTitle** | Pointer to **string** | Notification message dialog title displayed when Autoscale issues a logoff reminder request. | [optional] 

## Methods

### NewDeliveryGroupDetailResponseModelAllOf

`func NewDeliveryGroupDetailResponseModelAllOf(desktopsInUse int32, desktopsNeverRegistered int32, desktopsPreparing int32, machinesInMaintenanceMode int32, machineOperatingSystems []NameValueIntPairModel, machineType MachineType, offMachines int32, registeredMachines int32, secureIcaRequired bool, totalApplicationGroups int32, totalMachines int32, ) *DeliveryGroupDetailResponseModelAllOf`

NewDeliveryGroupDetailResponseModelAllOf instantiates a new DeliveryGroupDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupDetailResponseModelAllOfWithDefaults

`func NewDeliveryGroupDetailResponseModelAllOfWithDefaults() *DeliveryGroupDetailResponseModelAllOf`

NewDeliveryGroupDetailResponseModelAllOfWithDefaults instantiates a new DeliveryGroupDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppAccessPolicy

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAppAccessPolicy() DeliveryGroupDetailResponseModelAllOfAppAccessPolicy`

GetAppAccessPolicy returns the AppAccessPolicy field if non-nil, zero value otherwise.

### GetAppAccessPolicyOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAppAccessPolicyOk() (*DeliveryGroupDetailResponseModelAllOfAppAccessPolicy, bool)`

GetAppAccessPolicyOk returns a tuple with the AppAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAccessPolicy

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAppAccessPolicy(v DeliveryGroupDetailResponseModelAllOfAppAccessPolicy)`

SetAppAccessPolicy sets AppAccessPolicy field to given value.

### HasAppAccessPolicy

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAppAccessPolicy() bool`

HasAppAccessPolicy returns a boolean if a field has been set.

### GetAppProtectionKeyLoggingRequired

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAppProtectionKeyLoggingRequired() bool`

GetAppProtectionKeyLoggingRequired returns the AppProtectionKeyLoggingRequired field if non-nil, zero value otherwise.

### GetAppProtectionKeyLoggingRequiredOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAppProtectionKeyLoggingRequiredOk() (*bool, bool)`

GetAppProtectionKeyLoggingRequiredOk returns a tuple with the AppProtectionKeyLoggingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtectionKeyLoggingRequired

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAppProtectionKeyLoggingRequired(v bool)`

SetAppProtectionKeyLoggingRequired sets AppProtectionKeyLoggingRequired field to given value.

### HasAppProtectionKeyLoggingRequired

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAppProtectionKeyLoggingRequired() bool`

HasAppProtectionKeyLoggingRequired returns a boolean if a field has been set.

### GetAppProtectionScreenCaptureRequired

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAppProtectionScreenCaptureRequired() bool`

GetAppProtectionScreenCaptureRequired returns the AppProtectionScreenCaptureRequired field if non-nil, zero value otherwise.

### GetAppProtectionScreenCaptureRequiredOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAppProtectionScreenCaptureRequiredOk() (*bool, bool)`

GetAppProtectionScreenCaptureRequiredOk returns a tuple with the AppProtectionScreenCaptureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtectionScreenCaptureRequired

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAppProtectionScreenCaptureRequired(v bool)`

SetAppProtectionScreenCaptureRequired sets AppProtectionScreenCaptureRequired field to given value.

### HasAppProtectionScreenCaptureRequired

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAppProtectionScreenCaptureRequired() bool`

HasAppProtectionScreenCaptureRequired returns a boolean if a field has been set.

### GetAutomaticPowerOnForAssigned

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutomaticPowerOnForAssigned() bool`

GetAutomaticPowerOnForAssigned returns the AutomaticPowerOnForAssigned field if non-nil, zero value otherwise.

### GetAutomaticPowerOnForAssignedOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutomaticPowerOnForAssignedOk() (*bool, bool)`

GetAutomaticPowerOnForAssignedOk returns a tuple with the AutomaticPowerOnForAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPowerOnForAssigned

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAutomaticPowerOnForAssigned(v bool)`

SetAutomaticPowerOnForAssigned sets AutomaticPowerOnForAssigned field to given value.

### HasAutomaticPowerOnForAssigned

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAutomaticPowerOnForAssigned() bool`

HasAutomaticPowerOnForAssigned returns a boolean if a field has been set.

### GetAutomaticPowerOnForAssignedDuringPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutomaticPowerOnForAssignedDuringPeak() bool`

GetAutomaticPowerOnForAssignedDuringPeak returns the AutomaticPowerOnForAssignedDuringPeak field if non-nil, zero value otherwise.

### GetAutomaticPowerOnForAssignedDuringPeakOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutomaticPowerOnForAssignedDuringPeakOk() (*bool, bool)`

GetAutomaticPowerOnForAssignedDuringPeakOk returns a tuple with the AutomaticPowerOnForAssignedDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPowerOnForAssignedDuringPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAutomaticPowerOnForAssignedDuringPeak(v bool)`

SetAutomaticPowerOnForAssignedDuringPeak sets AutomaticPowerOnForAssignedDuringPeak field to given value.

### HasAutomaticPowerOnForAssignedDuringPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAutomaticPowerOnForAssignedDuringPeak() bool`

HasAutomaticPowerOnForAssignedDuringPeak returns a boolean if a field has been set.

### GetAutoScaleEnabled

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoScaleEnabled() bool`

GetAutoScaleEnabled returns the AutoScaleEnabled field if non-nil, zero value otherwise.

### GetAutoScaleEnabledOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoScaleEnabledOk() (*bool, bool)`

GetAutoScaleEnabledOk returns a tuple with the AutoScaleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaleEnabled

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAutoScaleEnabled(v bool)`

SetAutoScaleEnabled sets AutoScaleEnabled field to given value.

### HasAutoScaleEnabled

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAutoScaleEnabled() bool`

HasAutoScaleEnabled returns a boolean if a field has been set.

### GetRestrictAutoscaleTag

`func (o *DeliveryGroupDetailResponseModelAllOf) GetRestrictAutoscaleTag() DeliveryGroupDetailResponseModelAllOfRestrictAutoscaleTag`

GetRestrictAutoscaleTag returns the RestrictAutoscaleTag field if non-nil, zero value otherwise.

### GetRestrictAutoscaleTagOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetRestrictAutoscaleTagOk() (*DeliveryGroupDetailResponseModelAllOfRestrictAutoscaleTag, bool)`

GetRestrictAutoscaleTagOk returns a tuple with the RestrictAutoscaleTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleTag

`func (o *DeliveryGroupDetailResponseModelAllOf) SetRestrictAutoscaleTag(v DeliveryGroupDetailResponseModelAllOfRestrictAutoscaleTag)`

SetRestrictAutoscaleTag sets RestrictAutoscaleTag field to given value.

### HasRestrictAutoscaleTag

`func (o *DeliveryGroupDetailResponseModelAllOf) HasRestrictAutoscaleTag() bool`

HasRestrictAutoscaleTag returns a boolean if a field has been set.

### GetColorDepth

`func (o *DeliveryGroupDetailResponseModelAllOf) GetColorDepth() ColorDepth`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetColorDepthOk() (*ColorDepth, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *DeliveryGroupDetailResponseModelAllOf) SetColorDepth(v ColorDepth)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *DeliveryGroupDetailResponseModelAllOf) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetDefaultDesktopIconId

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDefaultDesktopIconId() string`

GetDefaultDesktopIconId returns the DefaultDesktopIconId field if non-nil, zero value otherwise.

### GetDefaultDesktopIconIdOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDefaultDesktopIconIdOk() (*string, bool)`

GetDefaultDesktopIconIdOk returns a tuple with the DefaultDesktopIconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDesktopIconId

`func (o *DeliveryGroupDetailResponseModelAllOf) SetDefaultDesktopIconId(v string)`

SetDefaultDesktopIconId sets DefaultDesktopIconId field to given value.

### HasDefaultDesktopIconId

`func (o *DeliveryGroupDetailResponseModelAllOf) HasDefaultDesktopIconId() bool`

HasDefaultDesktopIconId returns a boolean if a field has been set.

### GetDefaultDesktopPublishedName

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDefaultDesktopPublishedName() string`

GetDefaultDesktopPublishedName returns the DefaultDesktopPublishedName field if non-nil, zero value otherwise.

### GetDefaultDesktopPublishedNameOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDefaultDesktopPublishedNameOk() (*string, bool)`

GetDefaultDesktopPublishedNameOk returns a tuple with the DefaultDesktopPublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDesktopPublishedName

`func (o *DeliveryGroupDetailResponseModelAllOf) SetDefaultDesktopPublishedName(v string)`

SetDefaultDesktopPublishedName sets DefaultDesktopPublishedName field to given value.

### HasDefaultDesktopPublishedName

`func (o *DeliveryGroupDetailResponseModelAllOf) HasDefaultDesktopPublishedName() bool`

HasDefaultDesktopPublishedName returns a boolean if a field has been set.

### GetDesktopsInUse

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDesktopsInUse() int32`

GetDesktopsInUse returns the DesktopsInUse field if non-nil, zero value otherwise.

### GetDesktopsInUseOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDesktopsInUseOk() (*int32, bool)`

GetDesktopsInUseOk returns a tuple with the DesktopsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsInUse

`func (o *DeliveryGroupDetailResponseModelAllOf) SetDesktopsInUse(v int32)`

SetDesktopsInUse sets DesktopsInUse field to given value.


### GetDesktopsNeverRegistered

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDesktopsNeverRegistered() int32`

GetDesktopsNeverRegistered returns the DesktopsNeverRegistered field if non-nil, zero value otherwise.

### GetDesktopsNeverRegisteredOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDesktopsNeverRegisteredOk() (*int32, bool)`

GetDesktopsNeverRegisteredOk returns a tuple with the DesktopsNeverRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsNeverRegistered

`func (o *DeliveryGroupDetailResponseModelAllOf) SetDesktopsNeverRegistered(v int32)`

SetDesktopsNeverRegistered sets DesktopsNeverRegistered field to given value.


### GetDesktopsPreparing

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDesktopsPreparing() int32`

GetDesktopsPreparing returns the DesktopsPreparing field if non-nil, zero value otherwise.

### GetDesktopsPreparingOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDesktopsPreparingOk() (*int32, bool)`

GetDesktopsPreparingOk returns a tuple with the DesktopsPreparing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsPreparing

`func (o *DeliveryGroupDetailResponseModelAllOf) SetDesktopsPreparing(v int32)`

SetDesktopsPreparing sets DesktopsPreparing field to given value.


### GetHdxSslEnabled

`func (o *DeliveryGroupDetailResponseModelAllOf) GetHdxSslEnabled() bool`

GetHdxSslEnabled returns the HdxSslEnabled field if non-nil, zero value otherwise.

### GetHdxSslEnabledOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetHdxSslEnabledOk() (*bool, bool)`

GetHdxSslEnabledOk returns a tuple with the HdxSslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdxSslEnabled

`func (o *DeliveryGroupDetailResponseModelAllOf) SetHdxSslEnabled(v bool)`

SetHdxSslEnabled sets HdxSslEnabled field to given value.

### HasHdxSslEnabled

`func (o *DeliveryGroupDetailResponseModelAllOf) HasHdxSslEnabled() bool`

HasHdxSslEnabled returns a boolean if a field has been set.

### GetIsPowerManaged

`func (o *DeliveryGroupDetailResponseModelAllOf) GetIsPowerManaged() bool`

GetIsPowerManaged returns the IsPowerManaged field if non-nil, zero value otherwise.

### GetIsPowerManagedOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetIsPowerManagedOk() (*bool, bool)`

GetIsPowerManagedOk returns a tuple with the IsPowerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPowerManaged

`func (o *DeliveryGroupDetailResponseModelAllOf) SetIsPowerManaged(v bool)`

SetIsPowerManaged sets IsPowerManaged field to given value.

### HasIsPowerManaged

`func (o *DeliveryGroupDetailResponseModelAllOf) HasIsPowerManaged() bool`

HasIsPowerManaged returns a boolean if a field has been set.

### GetLingerSettings

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLingerSettings() DeliveryGroupDetailResponseModelAllOfLingerSettings`

GetLingerSettings returns the LingerSettings field if non-nil, zero value otherwise.

### GetLingerSettingsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLingerSettingsOk() (*DeliveryGroupDetailResponseModelAllOfLingerSettings, bool)`

GetLingerSettingsOk returns a tuple with the LingerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLingerSettings

`func (o *DeliveryGroupDetailResponseModelAllOf) SetLingerSettings(v DeliveryGroupDetailResponseModelAllOfLingerSettings)`

SetLingerSettings sets LingerSettings field to given value.

### HasLingerSettings

`func (o *DeliveryGroupDetailResponseModelAllOf) HasLingerSettings() bool`

HasLingerSettings returns a boolean if a field has been set.

### GetPowerOffDelayMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPowerOffDelayMinutes() int32`

GetPowerOffDelayMinutes returns the PowerOffDelayMinutes field if non-nil, zero value otherwise.

### GetPowerOffDelayMinutesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPowerOffDelayMinutesOk() (*int32, bool)`

GetPowerOffDelayMinutesOk returns a tuple with the PowerOffDelayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffDelayMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPowerOffDelayMinutes(v int32)`

SetPowerOffDelayMinutes sets PowerOffDelayMinutes field to given value.

### HasPowerOffDelayMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPowerOffDelayMinutes() bool`

HasPowerOffDelayMinutes returns a boolean if a field has been set.

### GetMachineCost

`func (o *DeliveryGroupDetailResponseModelAllOf) GetMachineCost() float64`

GetMachineCost returns the MachineCost field if non-nil, zero value otherwise.

### GetMachineCostOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetMachineCostOk() (*float64, bool)`

GetMachineCostOk returns a tuple with the MachineCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCost

`func (o *DeliveryGroupDetailResponseModelAllOf) SetMachineCost(v float64)`

SetMachineCost sets MachineCost field to given value.

### HasMachineCost

`func (o *DeliveryGroupDetailResponseModelAllOf) HasMachineCost() bool`

HasMachineCost returns a boolean if a field has been set.

### GetMachinesInMaintenanceMode

`func (o *DeliveryGroupDetailResponseModelAllOf) GetMachinesInMaintenanceMode() int32`

GetMachinesInMaintenanceMode returns the MachinesInMaintenanceMode field if non-nil, zero value otherwise.

### GetMachinesInMaintenanceModeOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetMachinesInMaintenanceModeOk() (*int32, bool)`

GetMachinesInMaintenanceModeOk returns a tuple with the MachinesInMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinesInMaintenanceMode

`func (o *DeliveryGroupDetailResponseModelAllOf) SetMachinesInMaintenanceMode(v int32)`

SetMachinesInMaintenanceMode sets MachinesInMaintenanceMode field to given value.


### GetMachineOperatingSystems

`func (o *DeliveryGroupDetailResponseModelAllOf) GetMachineOperatingSystems() []NameValueIntPairModel`

GetMachineOperatingSystems returns the MachineOperatingSystems field if non-nil, zero value otherwise.

### GetMachineOperatingSystemsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetMachineOperatingSystemsOk() (*[]NameValueIntPairModel, bool)`

GetMachineOperatingSystemsOk returns a tuple with the MachineOperatingSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineOperatingSystems

`func (o *DeliveryGroupDetailResponseModelAllOf) SetMachineOperatingSystems(v []NameValueIntPairModel)`

SetMachineOperatingSystems sets MachineOperatingSystems field to given value.


### GetMachineType

`func (o *DeliveryGroupDetailResponseModelAllOf) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *DeliveryGroupDetailResponseModelAllOf) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.


### GetOffMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffMachines() int32`

GetOffMachines returns the OffMachines field if non-nil, zero value otherwise.

### GetOffMachinesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffMachinesOk() (*int32, bool)`

GetOffMachinesOk returns a tuple with the OffMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) SetOffMachines(v int32)`

SetOffMachines sets OffMachines field to given value.


### GetOffPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakBufferSizePercent() int32`

GetOffPeakBufferSizePercent returns the OffPeakBufferSizePercent field if non-nil, zero value otherwise.

### GetOffPeakBufferSizePercentOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakBufferSizePercentOk() (*int32, bool)`

GetOffPeakBufferSizePercentOk returns a tuple with the OffPeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModelAllOf) SetOffPeakBufferSizePercent(v int32)`

SetOffPeakBufferSizePercent sets OffPeakBufferSizePercent field to given value.

### HasOffPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModelAllOf) HasOffPeakBufferSizePercent() bool`

HasOffPeakBufferSizePercent returns a boolean if a field has been set.

### GetOffPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakDisconnectAction() SessionChangeHostingAction`

GetOffPeakDisconnectAction returns the OffPeakDisconnectAction field if non-nil, zero value otherwise.

### GetOffPeakDisconnectActionOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakDisconnectActionOk returns a tuple with the OffPeakDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) SetOffPeakDisconnectAction(v SessionChangeHostingAction)`

SetOffPeakDisconnectAction sets OffPeakDisconnectAction field to given value.

### HasOffPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) HasOffPeakDisconnectAction() bool`

HasOffPeakDisconnectAction returns a boolean if a field has been set.

### GetOffPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakDisconnectTimeoutMinutes() int32`

GetOffPeakDisconnectTimeoutMinutes returns the OffPeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakDisconnectTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakDisconnectTimeoutMinutesOk returns a tuple with the OffPeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) SetOffPeakDisconnectTimeoutMinutes(v int32)`

SetOffPeakDisconnectTimeoutMinutes sets OffPeakDisconnectTimeoutMinutes field to given value.

### HasOffPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) HasOffPeakDisconnectTimeoutMinutes() bool`

HasOffPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetOffPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakExtendedDisconnectAction() SessionChangeHostingAction`

GetOffPeakExtendedDisconnectAction returns the OffPeakExtendedDisconnectAction field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectActionOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakExtendedDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakExtendedDisconnectActionOk returns a tuple with the OffPeakExtendedDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) SetOffPeakExtendedDisconnectAction(v SessionChangeHostingAction)`

SetOffPeakExtendedDisconnectAction sets OffPeakExtendedDisconnectAction field to given value.

### HasOffPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) HasOffPeakExtendedDisconnectAction() bool`

HasOffPeakExtendedDisconnectAction returns a boolean if a field has been set.

### GetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakExtendedDisconnectTimeoutMinutes() int32`

GetOffPeakExtendedDisconnectTimeoutMinutes returns the OffPeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the OffPeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) SetOffPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetOffPeakExtendedDisconnectTimeoutMinutes sets OffPeakExtendedDisconnectTimeoutMinutes field to given value.

### HasOffPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) HasOffPeakExtendedDisconnectTimeoutMinutes() bool`

HasOffPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetOffPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakLogOffAction() SessionChangeHostingAction`

GetOffPeakLogOffAction returns the OffPeakLogOffAction field if non-nil, zero value otherwise.

### GetOffPeakLogOffActionOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakLogOffActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakLogOffActionOk returns a tuple with the OffPeakLogOffAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModelAllOf) SetOffPeakLogOffAction(v SessionChangeHostingAction)`

SetOffPeakLogOffAction sets OffPeakLogOffAction field to given value.

### HasOffPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModelAllOf) HasOffPeakLogOffAction() bool`

HasOffPeakLogOffAction returns a boolean if a field has been set.

### GetOffPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakLogOffTimeoutMinutes() int32`

GetOffPeakLogOffTimeoutMinutes returns the OffPeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakLogOffTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetOffPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetOffPeakLogOffTimeoutMinutesOk returns a tuple with the OffPeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) SetOffPeakLogOffTimeoutMinutes(v int32)`

SetOffPeakLogOffTimeoutMinutes sets OffPeakLogOffTimeoutMinutes field to given value.

### HasOffPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) HasOffPeakLogOffTimeoutMinutes() bool`

HasOffPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### GetPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakBufferSizePercent() int32`

GetPeakBufferSizePercent returns the PeakBufferSizePercent field if non-nil, zero value otherwise.

### GetPeakBufferSizePercentOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakBufferSizePercentOk() (*int32, bool)`

GetPeakBufferSizePercentOk returns a tuple with the PeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPeakBufferSizePercent(v int32)`

SetPeakBufferSizePercent sets PeakBufferSizePercent field to given value.

### HasPeakBufferSizePercent

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPeakBufferSizePercent() bool`

HasPeakBufferSizePercent returns a boolean if a field has been set.

### GetPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakDisconnectAction() SessionChangeHostingAction`

GetPeakDisconnectAction returns the PeakDisconnectAction field if non-nil, zero value otherwise.

### GetPeakDisconnectActionOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetPeakDisconnectActionOk returns a tuple with the PeakDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPeakDisconnectAction(v SessionChangeHostingAction)`

SetPeakDisconnectAction sets PeakDisconnectAction field to given value.

### HasPeakDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPeakDisconnectAction() bool`

HasPeakDisconnectAction returns a boolean if a field has been set.

### GetPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakDisconnectTimeoutMinutes() int32`

GetPeakDisconnectTimeoutMinutes returns the PeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakDisconnectTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakDisconnectTimeoutMinutesOk returns a tuple with the PeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPeakDisconnectTimeoutMinutes(v int32)`

SetPeakDisconnectTimeoutMinutes sets PeakDisconnectTimeoutMinutes field to given value.

### HasPeakDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPeakDisconnectTimeoutMinutes() bool`

HasPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakExtendedDisconnectAction() SessionChangeHostingAction`

GetPeakExtendedDisconnectAction returns the PeakExtendedDisconnectAction field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectActionOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakExtendedDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetPeakExtendedDisconnectActionOk returns a tuple with the PeakExtendedDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPeakExtendedDisconnectAction(v SessionChangeHostingAction)`

SetPeakExtendedDisconnectAction sets PeakExtendedDisconnectAction field to given value.

### HasPeakExtendedDisconnectAction

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPeakExtendedDisconnectAction() bool`

HasPeakExtendedDisconnectAction returns a boolean if a field has been set.

### GetPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakExtendedDisconnectTimeoutMinutes() int32`

GetPeakExtendedDisconnectTimeoutMinutes returns the PeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the PeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetPeakExtendedDisconnectTimeoutMinutes sets PeakExtendedDisconnectTimeoutMinutes field to given value.

### HasPeakExtendedDisconnectTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPeakExtendedDisconnectTimeoutMinutes() bool`

HasPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakLogOffAction() SessionChangeHostingAction`

GetPeakLogOffAction returns the PeakLogOffAction field if non-nil, zero value otherwise.

### GetPeakLogOffActionOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakLogOffActionOk() (*SessionChangeHostingAction, bool)`

GetPeakLogOffActionOk returns a tuple with the PeakLogOffAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPeakLogOffAction(v SessionChangeHostingAction)`

SetPeakLogOffAction sets PeakLogOffAction field to given value.

### HasPeakLogOffAction

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPeakLogOffAction() bool`

HasPeakLogOffAction returns a boolean if a field has been set.

### GetPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakLogOffTimeoutMinutes() int32`

GetPeakLogOffTimeoutMinutes returns the PeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakLogOffTimeoutMinutesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetPeakLogOffTimeoutMinutesOk returns a tuple with the PeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPeakLogOffTimeoutMinutes(v int32)`

SetPeakLogOffTimeoutMinutes sets PeakLogOffTimeoutMinutes field to given value.

### HasPeakLogOffTimeoutMinutes

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPeakLogOffTimeoutMinutes() bool`

HasPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### GetDisconnectPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDisconnectPeakIdleSessionAfterSeconds() int32`

GetDisconnectPeakIdleSessionAfterSeconds returns the DisconnectPeakIdleSessionAfterSeconds field if non-nil, zero value otherwise.

### GetDisconnectPeakIdleSessionAfterSecondsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDisconnectPeakIdleSessionAfterSecondsOk() (*int32, bool)`

GetDisconnectPeakIdleSessionAfterSecondsOk returns a tuple with the DisconnectPeakIdleSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) SetDisconnectPeakIdleSessionAfterSeconds(v int32)`

SetDisconnectPeakIdleSessionAfterSeconds sets DisconnectPeakIdleSessionAfterSeconds field to given value.

### HasDisconnectPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) HasDisconnectPeakIdleSessionAfterSeconds() bool`

HasDisconnectPeakIdleSessionAfterSeconds returns a boolean if a field has been set.

### GetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDisconnectOffPeakIdleSessionAfterSeconds() int32`

GetDisconnectOffPeakIdleSessionAfterSeconds returns the DisconnectOffPeakIdleSessionAfterSeconds field if non-nil, zero value otherwise.

### GetDisconnectOffPeakIdleSessionAfterSecondsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetDisconnectOffPeakIdleSessionAfterSecondsOk() (*int32, bool)`

GetDisconnectOffPeakIdleSessionAfterSecondsOk returns a tuple with the DisconnectOffPeakIdleSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) SetDisconnectOffPeakIdleSessionAfterSeconds(v int32)`

SetDisconnectOffPeakIdleSessionAfterSeconds sets DisconnectOffPeakIdleSessionAfterSeconds field to given value.

### HasDisconnectOffPeakIdleSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) HasDisconnectOffPeakIdleSessionAfterSeconds() bool`

HasDisconnectOffPeakIdleSessionAfterSeconds returns a boolean if a field has been set.

### GetLogoffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLogoffPeakDisconnectedSessionAfterSeconds() int32`

GetLogoffPeakDisconnectedSessionAfterSeconds returns the LogoffPeakDisconnectedSessionAfterSeconds field if non-nil, zero value otherwise.

### GetLogoffPeakDisconnectedSessionAfterSecondsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLogoffPeakDisconnectedSessionAfterSecondsOk() (*int32, bool)`

GetLogoffPeakDisconnectedSessionAfterSecondsOk returns a tuple with the LogoffPeakDisconnectedSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) SetLogoffPeakDisconnectedSessionAfterSeconds(v int32)`

SetLogoffPeakDisconnectedSessionAfterSeconds sets LogoffPeakDisconnectedSessionAfterSeconds field to given value.

### HasLogoffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) HasLogoffPeakDisconnectedSessionAfterSeconds() bool`

HasLogoffPeakDisconnectedSessionAfterSeconds returns a boolean if a field has been set.

### GetLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLogoffOffPeakDisconnectedSessionAfterSeconds() int32`

GetLogoffOffPeakDisconnectedSessionAfterSeconds returns the LogoffOffPeakDisconnectedSessionAfterSeconds field if non-nil, zero value otherwise.

### GetLogoffOffPeakDisconnectedSessionAfterSecondsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLogoffOffPeakDisconnectedSessionAfterSecondsOk() (*int32, bool)`

GetLogoffOffPeakDisconnectedSessionAfterSecondsOk returns a tuple with the LogoffOffPeakDisconnectedSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) SetLogoffOffPeakDisconnectedSessionAfterSeconds(v int32)`

SetLogoffOffPeakDisconnectedSessionAfterSeconds sets LogoffOffPeakDisconnectedSessionAfterSeconds field to given value.

### HasLogoffOffPeakDisconnectedSessionAfterSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) HasLogoffOffPeakDisconnectedSessionAfterSeconds() bool`

HasLogoffOffPeakDisconnectedSessionAfterSeconds returns a boolean if a field has been set.

### GetPrelaunchSettings

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPrelaunchSettings() DeliveryGroupDetailResponseModelAllOfPrelaunchSettings`

GetPrelaunchSettings returns the PrelaunchSettings field if non-nil, zero value otherwise.

### GetPrelaunchSettingsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPrelaunchSettingsOk() (*DeliveryGroupDetailResponseModelAllOfPrelaunchSettings, bool)`

GetPrelaunchSettingsOk returns a tuple with the PrelaunchSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrelaunchSettings

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPrelaunchSettings(v DeliveryGroupDetailResponseModelAllOfPrelaunchSettings)`

SetPrelaunchSettings sets PrelaunchSettings field to given value.

### HasPrelaunchSettings

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPrelaunchSettings() bool`

HasPrelaunchSettings returns a boolean if a field has been set.

### GetPowerTimeSchemes

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPowerTimeSchemes() []PowerTimeSchemeResponseModel`

GetPowerTimeSchemes returns the PowerTimeSchemes field if non-nil, zero value otherwise.

### GetPowerTimeSchemesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetPowerTimeSchemesOk() (*[]PowerTimeSchemeResponseModel, bool)`

GetPowerTimeSchemesOk returns a tuple with the PowerTimeSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerTimeSchemes

`func (o *DeliveryGroupDetailResponseModelAllOf) SetPowerTimeSchemes(v []PowerTimeSchemeResponseModel)`

SetPowerTimeSchemes sets PowerTimeSchemes field to given value.

### HasPowerTimeSchemes

`func (o *DeliveryGroupDetailResponseModelAllOf) HasPowerTimeSchemes() bool`

HasPowerTimeSchemes returns a boolean if a field has been set.

### GetProtocolPriority

`func (o *DeliveryGroupDetailResponseModelAllOf) GetProtocolPriority() []ProtocolType`

GetProtocolPriority returns the ProtocolPriority field if non-nil, zero value otherwise.

### GetProtocolPriorityOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetProtocolPriorityOk() (*[]ProtocolType, bool)`

GetProtocolPriorityOk returns a tuple with the ProtocolPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPriority

`func (o *DeliveryGroupDetailResponseModelAllOf) SetProtocolPriority(v []ProtocolType)`

SetProtocolPriority sets ProtocolPriority field to given value.

### HasProtocolPriority

`func (o *DeliveryGroupDetailResponseModelAllOf) HasProtocolPriority() bool`

HasProtocolPriority returns a boolean if a field has been set.

### GetRegisteredMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) GetRegisteredMachines() int32`

GetRegisteredMachines returns the RegisteredMachines field if non-nil, zero value otherwise.

### GetRegisteredMachinesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetRegisteredMachinesOk() (*int32, bool)`

GetRegisteredMachinesOk returns a tuple with the RegisteredMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) SetRegisteredMachines(v int32)`

SetRegisteredMachines sets RegisteredMachines field to given value.


### GetReuseMachinesWithoutShutdownInOutage

`func (o *DeliveryGroupDetailResponseModelAllOf) GetReuseMachinesWithoutShutdownInOutage() bool`

GetReuseMachinesWithoutShutdownInOutage returns the ReuseMachinesWithoutShutdownInOutage field if non-nil, zero value otherwise.

### GetReuseMachinesWithoutShutdownInOutageOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetReuseMachinesWithoutShutdownInOutageOk() (*bool, bool)`

GetReuseMachinesWithoutShutdownInOutageOk returns a tuple with the ReuseMachinesWithoutShutdownInOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseMachinesWithoutShutdownInOutage

`func (o *DeliveryGroupDetailResponseModelAllOf) SetReuseMachinesWithoutShutdownInOutage(v bool)`

SetReuseMachinesWithoutShutdownInOutage sets ReuseMachinesWithoutShutdownInOutage field to given value.

### HasReuseMachinesWithoutShutdownInOutage

`func (o *DeliveryGroupDetailResponseModelAllOf) HasReuseMachinesWithoutShutdownInOutage() bool`

HasReuseMachinesWithoutShutdownInOutage returns a boolean if a field has been set.

### GetSecureIcaRequired

`func (o *DeliveryGroupDetailResponseModelAllOf) GetSecureIcaRequired() bool`

GetSecureIcaRequired returns the SecureIcaRequired field if non-nil, zero value otherwise.

### GetSecureIcaRequiredOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetSecureIcaRequiredOk() (*bool, bool)`

GetSecureIcaRequiredOk returns a tuple with the SecureIcaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureIcaRequired

`func (o *DeliveryGroupDetailResponseModelAllOf) SetSecureIcaRequired(v bool)`

SetSecureIcaRequired sets SecureIcaRequired field to given value.


### GetSettlementPeriodBeforeAutoShutdownSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) GetSettlementPeriodBeforeAutoShutdownSeconds() int32`

GetSettlementPeriodBeforeAutoShutdownSeconds returns the SettlementPeriodBeforeAutoShutdownSeconds field if non-nil, zero value otherwise.

### GetSettlementPeriodBeforeAutoShutdownSecondsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetSettlementPeriodBeforeAutoShutdownSecondsOk() (*int32, bool)`

GetSettlementPeriodBeforeAutoShutdownSecondsOk returns a tuple with the SettlementPeriodBeforeAutoShutdownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementPeriodBeforeAutoShutdownSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) SetSettlementPeriodBeforeAutoShutdownSeconds(v int32)`

SetSettlementPeriodBeforeAutoShutdownSeconds sets SettlementPeriodBeforeAutoShutdownSeconds field to given value.

### HasSettlementPeriodBeforeAutoShutdownSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) HasSettlementPeriodBeforeAutoShutdownSeconds() bool`

HasSettlementPeriodBeforeAutoShutdownSeconds returns a boolean if a field has been set.

### GetSettlementPeriodBeforeUseSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) GetSettlementPeriodBeforeUseSeconds() int32`

GetSettlementPeriodBeforeUseSeconds returns the SettlementPeriodBeforeUseSeconds field if non-nil, zero value otherwise.

### GetSettlementPeriodBeforeUseSecondsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetSettlementPeriodBeforeUseSecondsOk() (*int32, bool)`

GetSettlementPeriodBeforeUseSecondsOk returns a tuple with the SettlementPeriodBeforeUseSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementPeriodBeforeUseSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) SetSettlementPeriodBeforeUseSeconds(v int32)`

SetSettlementPeriodBeforeUseSeconds sets SettlementPeriodBeforeUseSeconds field to given value.

### HasSettlementPeriodBeforeUseSeconds

`func (o *DeliveryGroupDetailResponseModelAllOf) HasSettlementPeriodBeforeUseSeconds() bool`

HasSettlementPeriodBeforeUseSeconds returns a boolean if a field has been set.

### GetShutdownDesktopsAfterUse

`func (o *DeliveryGroupDetailResponseModelAllOf) GetShutdownDesktopsAfterUse() bool`

GetShutdownDesktopsAfterUse returns the ShutdownDesktopsAfterUse field if non-nil, zero value otherwise.

### GetShutdownDesktopsAfterUseOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetShutdownDesktopsAfterUseOk() (*bool, bool)`

GetShutdownDesktopsAfterUseOk returns a tuple with the ShutdownDesktopsAfterUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDesktopsAfterUse

`func (o *DeliveryGroupDetailResponseModelAllOf) SetShutdownDesktopsAfterUse(v bool)`

SetShutdownDesktopsAfterUse sets ShutdownDesktopsAfterUse field to given value.

### HasShutdownDesktopsAfterUse

`func (o *DeliveryGroupDetailResponseModelAllOf) HasShutdownDesktopsAfterUse() bool`

HasShutdownDesktopsAfterUse returns a boolean if a field has been set.

### GetSimpleAccessPolicy

`func (o *DeliveryGroupDetailResponseModelAllOf) GetSimpleAccessPolicy() DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy`

GetSimpleAccessPolicy returns the SimpleAccessPolicy field if non-nil, zero value otherwise.

### GetSimpleAccessPolicyOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetSimpleAccessPolicyOk() (*DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy, bool)`

GetSimpleAccessPolicyOk returns a tuple with the SimpleAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAccessPolicy

`func (o *DeliveryGroupDetailResponseModelAllOf) SetSimpleAccessPolicy(v DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy)`

SetSimpleAccessPolicy sets SimpleAccessPolicy field to given value.

### HasSimpleAccessPolicy

`func (o *DeliveryGroupDetailResponseModelAllOf) HasSimpleAccessPolicy() bool`

HasSimpleAccessPolicy returns a boolean if a field has been set.

### GetStoreFrontServersForHostedReceiver

`func (o *DeliveryGroupDetailResponseModelAllOf) GetStoreFrontServersForHostedReceiver() []StoreFrontServerResponseModel`

GetStoreFrontServersForHostedReceiver returns the StoreFrontServersForHostedReceiver field if non-nil, zero value otherwise.

### GetStoreFrontServersForHostedReceiverOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetStoreFrontServersForHostedReceiverOk() (*[]StoreFrontServerResponseModel, bool)`

GetStoreFrontServersForHostedReceiverOk returns a tuple with the StoreFrontServersForHostedReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFrontServersForHostedReceiver

`func (o *DeliveryGroupDetailResponseModelAllOf) SetStoreFrontServersForHostedReceiver(v []StoreFrontServerResponseModel)`

SetStoreFrontServersForHostedReceiver sets StoreFrontServersForHostedReceiver field to given value.

### HasStoreFrontServersForHostedReceiver

`func (o *DeliveryGroupDetailResponseModelAllOf) HasStoreFrontServersForHostedReceiver() bool`

HasStoreFrontServersForHostedReceiver returns a boolean if a field has been set.

### GetTimeZone

`func (o *DeliveryGroupDetailResponseModelAllOf) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *DeliveryGroupDetailResponseModelAllOf) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *DeliveryGroupDetailResponseModelAllOf) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTotalApplicationGroups

`func (o *DeliveryGroupDetailResponseModelAllOf) GetTotalApplicationGroups() int32`

GetTotalApplicationGroups returns the TotalApplicationGroups field if non-nil, zero value otherwise.

### GetTotalApplicationGroupsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetTotalApplicationGroupsOk() (*int32, bool)`

GetTotalApplicationGroupsOk returns a tuple with the TotalApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplicationGroups

`func (o *DeliveryGroupDetailResponseModelAllOf) SetTotalApplicationGroups(v int32)`

SetTotalApplicationGroups sets TotalApplicationGroups field to given value.


### GetTotalMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) GetTotalMachines() int32`

GetTotalMachines returns the TotalMachines field if non-nil, zero value otherwise.

### GetTotalMachinesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetTotalMachinesOk() (*int32, bool)`

GetTotalMachinesOk returns a tuple with the TotalMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) SetTotalMachines(v int32)`

SetTotalMachines sets TotalMachines field to given value.


### GetTurnOnAddedMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) GetTurnOnAddedMachines() bool`

GetTurnOnAddedMachines returns the TurnOnAddedMachines field if non-nil, zero value otherwise.

### GetTurnOnAddedMachinesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetTurnOnAddedMachinesOk() (*bool, bool)`

GetTurnOnAddedMachinesOk returns a tuple with the TurnOnAddedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOnAddedMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) SetTurnOnAddedMachines(v bool)`

SetTurnOnAddedMachines sets TurnOnAddedMachines field to given value.

### HasTurnOnAddedMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) HasTurnOnAddedMachines() bool`

HasTurnOnAddedMachines returns a boolean if a field has been set.

### GetUnassignedMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) GetUnassignedMachines() int32`

GetUnassignedMachines returns the UnassignedMachines field if non-nil, zero value otherwise.

### GetUnassignedMachinesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetUnassignedMachinesOk() (*int32, bool)`

GetUnassignedMachinesOk returns a tuple with the UnassignedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) SetUnassignedMachines(v int32)`

SetUnassignedMachines sets UnassignedMachines field to given value.

### HasUnassignedMachines

`func (o *DeliveryGroupDetailResponseModelAllOf) HasUnassignedMachines() bool`

HasUnassignedMachines returns a boolean if a field has been set.

### GetUsedByGroupPolicy

`func (o *DeliveryGroupDetailResponseModelAllOf) GetUsedByGroupPolicy() bool`

GetUsedByGroupPolicy returns the UsedByGroupPolicy field if non-nil, zero value otherwise.

### GetUsedByGroupPolicyOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetUsedByGroupPolicyOk() (*bool, bool)`

GetUsedByGroupPolicyOk returns a tuple with the UsedByGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedByGroupPolicy

`func (o *DeliveryGroupDetailResponseModelAllOf) SetUsedByGroupPolicy(v bool)`

SetUsedByGroupPolicy sets UsedByGroupPolicy field to given value.

### HasUsedByGroupPolicy

`func (o *DeliveryGroupDetailResponseModelAllOf) HasUsedByGroupPolicy() bool`

HasUsedByGroupPolicy returns a boolean if a field has been set.

### GetVdaVersions

`func (o *DeliveryGroupDetailResponseModelAllOf) GetVdaVersions() []NameValueIntPairModel`

GetVdaVersions returns the VdaVersions field if non-nil, zero value otherwise.

### GetVdaVersionsOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetVdaVersionsOk() (*[]NameValueIntPairModel, bool)`

GetVdaVersionsOk returns a tuple with the VdaVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaVersions

`func (o *DeliveryGroupDetailResponseModelAllOf) SetVdaVersions(v []NameValueIntPairModel)`

SetVdaVersions sets VdaVersions field to given value.

### HasVdaVersions

`func (o *DeliveryGroupDetailResponseModelAllOf) HasVdaVersions() bool`

HasVdaVersions returns a boolean if a field has been set.

### GetZonePreferences

`func (o *DeliveryGroupDetailResponseModelAllOf) GetZonePreferences() []ZonePreference`

GetZonePreferences returns the ZonePreferences field if non-nil, zero value otherwise.

### GetZonePreferencesOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetZonePreferencesOk() (*[]ZonePreference, bool)`

GetZonePreferencesOk returns a tuple with the ZonePreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePreferences

`func (o *DeliveryGroupDetailResponseModelAllOf) SetZonePreferences(v []ZonePreference)`

SetZonePreferences sets ZonePreferences field to given value.

### HasZonePreferences

`func (o *DeliveryGroupDetailResponseModelAllOf) HasZonePreferences() bool`

HasZonePreferences returns a boolean if a field has been set.

### GetLimitSecondsToForceLogOffUserDuringPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLimitSecondsToForceLogOffUserDuringPeak() int32`

GetLimitSecondsToForceLogOffUserDuringPeak returns the LimitSecondsToForceLogOffUserDuringPeak field if non-nil, zero value otherwise.

### GetLimitSecondsToForceLogOffUserDuringPeakOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLimitSecondsToForceLogOffUserDuringPeakOk() (*int32, bool)`

GetLimitSecondsToForceLogOffUserDuringPeakOk returns a tuple with the LimitSecondsToForceLogOffUserDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSecondsToForceLogOffUserDuringPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) SetLimitSecondsToForceLogOffUserDuringPeak(v int32)`

SetLimitSecondsToForceLogOffUserDuringPeak sets LimitSecondsToForceLogOffUserDuringPeak field to given value.

### HasLimitSecondsToForceLogOffUserDuringPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) HasLimitSecondsToForceLogOffUserDuringPeak() bool`

HasLimitSecondsToForceLogOffUserDuringPeak returns a boolean if a field has been set.

### GetLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLimitSecondsToForceLogOffUserDuringOffPeak() int32`

GetLimitSecondsToForceLogOffUserDuringOffPeak returns the LimitSecondsToForceLogOffUserDuringOffPeak field if non-nil, zero value otherwise.

### GetLimitSecondsToForceLogOffUserDuringOffPeakOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLimitSecondsToForceLogOffUserDuringOffPeakOk() (*int32, bool)`

GetLimitSecondsToForceLogOffUserDuringOffPeakOk returns a tuple with the LimitSecondsToForceLogOffUserDuringOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) SetLimitSecondsToForceLogOffUserDuringOffPeak(v int32)`

SetLimitSecondsToForceLogOffUserDuringOffPeak sets LimitSecondsToForceLogOffUserDuringOffPeak field to given value.

### HasLimitSecondsToForceLogOffUserDuringOffPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) HasLimitSecondsToForceLogOffUserDuringOffPeak() bool`

HasLimitSecondsToForceLogOffUserDuringOffPeak returns a boolean if a field has been set.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak() int32`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak returns the RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field if non-nil, zero value otherwise.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk() (*int32, bool)`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringPeakOk returns a tuple with the RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak(v int32)`

SetRestrictAutoscaleMinIdleUntaggedPercentDuringPeak sets RestrictAutoscaleMinIdleUntaggedPercentDuringPeak field to given value.

### HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak() bool`

HasRestrictAutoscaleMinIdleUntaggedPercentDuringPeak returns a boolean if a field has been set.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak() int32`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak returns the RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field if non-nil, zero value otherwise.

### GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk() (*int32, bool)`

GetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeakOk returns a tuple with the RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak(v int32)`

SetRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak sets RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak field to given value.

### HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak() bool`

HasRestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak returns a boolean if a field has been set.

### GetLogOffWarningMessage

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLogOffWarningMessage() string`

GetLogOffWarningMessage returns the LogOffWarningMessage field if non-nil, zero value otherwise.

### GetLogOffWarningMessageOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLogOffWarningMessageOk() (*string, bool)`

GetLogOffWarningMessageOk returns a tuple with the LogOffWarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOffWarningMessage

`func (o *DeliveryGroupDetailResponseModelAllOf) SetLogOffWarningMessage(v string)`

SetLogOffWarningMessage sets LogOffWarningMessage field to given value.

### HasLogOffWarningMessage

`func (o *DeliveryGroupDetailResponseModelAllOf) HasLogOffWarningMessage() bool`

HasLogOffWarningMessage returns a boolean if a field has been set.

### GetLogOffWarningTitle

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLogOffWarningTitle() string`

GetLogOffWarningTitle returns the LogOffWarningTitle field if non-nil, zero value otherwise.

### GetLogOffWarningTitleOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetLogOffWarningTitleOk() (*string, bool)`

GetLogOffWarningTitleOk returns a tuple with the LogOffWarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOffWarningTitle

`func (o *DeliveryGroupDetailResponseModelAllOf) SetLogOffWarningTitle(v string)`

SetLogOffWarningTitle sets LogOffWarningTitle field to given value.

### HasLogOffWarningTitle

`func (o *DeliveryGroupDetailResponseModelAllOf) HasLogOffWarningTitle() bool`

HasLogOffWarningTitle returns a boolean if a field has been set.

### GetAutoscaleLogOffReminderEnabled

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderEnabled() bool`

GetAutoscaleLogOffReminderEnabled returns the AutoscaleLogOffReminderEnabled field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderEnabledOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderEnabledOk() (*bool, bool)`

GetAutoscaleLogOffReminderEnabledOk returns a tuple with the AutoscaleLogOffReminderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderEnabled

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAutoscaleLogOffReminderEnabled(v bool)`

SetAutoscaleLogOffReminderEnabled sets AutoscaleLogOffReminderEnabled field to given value.

### HasAutoscaleLogOffReminderEnabled

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAutoscaleLogOffReminderEnabled() bool`

HasAutoscaleLogOffReminderEnabled returns a boolean if a field has been set.

### GetAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderIntervalSecondsOffPeak() int32`

GetAutoscaleLogOffReminderIntervalSecondsOffPeak returns the AutoscaleLogOffReminderIntervalSecondsOffPeak field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk() (*int32, bool)`

GetAutoscaleLogOffReminderIntervalSecondsOffPeakOk returns a tuple with the AutoscaleLogOffReminderIntervalSecondsOffPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAutoscaleLogOffReminderIntervalSecondsOffPeak(v int32)`

SetAutoscaleLogOffReminderIntervalSecondsOffPeak sets AutoscaleLogOffReminderIntervalSecondsOffPeak field to given value.

### HasAutoscaleLogOffReminderIntervalSecondsOffPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAutoscaleLogOffReminderIntervalSecondsOffPeak() bool`

HasAutoscaleLogOffReminderIntervalSecondsOffPeak returns a boolean if a field has been set.

### GetAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderIntervalSecondsPeak() int32`

GetAutoscaleLogOffReminderIntervalSecondsPeak returns the AutoscaleLogOffReminderIntervalSecondsPeak field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderIntervalSecondsPeakOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderIntervalSecondsPeakOk() (*int32, bool)`

GetAutoscaleLogOffReminderIntervalSecondsPeakOk returns a tuple with the AutoscaleLogOffReminderIntervalSecondsPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAutoscaleLogOffReminderIntervalSecondsPeak(v int32)`

SetAutoscaleLogOffReminderIntervalSecondsPeak sets AutoscaleLogOffReminderIntervalSecondsPeak field to given value.

### HasAutoscaleLogOffReminderIntervalSecondsPeak

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAutoscaleLogOffReminderIntervalSecondsPeak() bool`

HasAutoscaleLogOffReminderIntervalSecondsPeak returns a boolean if a field has been set.

### GetAutoscaleLogOffReminderMessage

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderMessage() string`

GetAutoscaleLogOffReminderMessage returns the AutoscaleLogOffReminderMessage field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderMessageOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderMessageOk() (*string, bool)`

GetAutoscaleLogOffReminderMessageOk returns a tuple with the AutoscaleLogOffReminderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderMessage

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAutoscaleLogOffReminderMessage(v string)`

SetAutoscaleLogOffReminderMessage sets AutoscaleLogOffReminderMessage field to given value.

### HasAutoscaleLogOffReminderMessage

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAutoscaleLogOffReminderMessage() bool`

HasAutoscaleLogOffReminderMessage returns a boolean if a field has been set.

### GetAutoscaleLogOffReminderTitle

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderTitle() string`

GetAutoscaleLogOffReminderTitle returns the AutoscaleLogOffReminderTitle field if non-nil, zero value otherwise.

### GetAutoscaleLogOffReminderTitleOk

`func (o *DeliveryGroupDetailResponseModelAllOf) GetAutoscaleLogOffReminderTitleOk() (*string, bool)`

GetAutoscaleLogOffReminderTitleOk returns a tuple with the AutoscaleLogOffReminderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaleLogOffReminderTitle

`func (o *DeliveryGroupDetailResponseModelAllOf) SetAutoscaleLogOffReminderTitle(v string)`

SetAutoscaleLogOffReminderTitle sets AutoscaleLogOffReminderTitle field to given value.

### HasAutoscaleLogOffReminderTitle

`func (o *DeliveryGroupDetailResponseModelAllOf) HasAutoscaleLogOffReminderTitle() bool`

HasAutoscaleLogOffReminderTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


