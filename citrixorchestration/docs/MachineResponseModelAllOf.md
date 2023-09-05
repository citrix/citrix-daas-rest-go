# MachineResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** | DEPRECATED. Use Id. Used to be: DesktopUid | [optional] 
**AgentVersion** | Pointer to **string** | Version of the Citrix Virtual Delivery Agent (VDA) installed on the machine. | [optional] 
**AllocationType** | Pointer to [**AllocationType**](AllocationType.md) |  | [optional] 
**ApplicationsInUse** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | List of applications in use in the session. | [optional] 
**AssignedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | List of one or more users to whom the machine is assigned. Only used when AllocationType is equal to Static. | [optional] 
**AssociatedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The current user(s) for shared machines and the assigned users for private machines. | [optional] 
**AzureAdJoinedMode** | Pointer to [**AzureAdJoinedMode**](AzureAdJoinedMode.md) |  | [optional] 
**ContainerScopes** | Pointer to [**[]ContainerScopeResponseModel**](ContainerScopeResponseModel.md) | Delegated admin scopes in which the containers of the machine reside. | [optional] 
**ControllerDnsName** | Pointer to **string** | The DNS host name of the controller that the machine is registered to. | [optional] 
**DeliveryGroup** | Pointer to [**MachineResponseModelAllOfDeliveryGroup**](MachineResponseModelAllOfDeliveryGroup.md) |  | [optional] 
**DeliveryType** | Pointer to [**DeliveryKind**](DeliveryKind.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the machine. | [optional] 
**DesktopConditions** | Pointer to [**[]DesktopCondition**](DesktopCondition.md) | List of outstanding desktop conditions for the machine. | [optional] 
**DnsName** | Pointer to **string** | The DNS host name of the machine. | [optional] 
**FunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**Hosting** | Pointer to [**MachineResponseModelAllOfHosting**](MachineResponseModelAllOfHosting.md) |  | [optional] 
**InMaintenanceMode** | **bool** | Denotes if the machine is in maintenance mode. Machines in maintenance mode will not accept new sessions. | 
**DrainingUntilShutdown** | **bool** | Denotes if the machine is placed to drain until shutdown | 
**IPAddress** | Pointer to **string** | The IP address of the machine. | [optional] 
**IsAssigned** | Pointer to **bool** | Denotes whether a private desktop has been assigned to a user/users, or a client name/address. Users can be assigned explicitly or by assigning on first use of the machine. Only relevant for privately assigned machines. | [optional] 
**MachineType** | [**MachineType**](MachineType.md) |  | 
**LastConnectionFailure** | Pointer to [**ConnectionFailureReason**](ConnectionFailureReason.md) |  | [optional] 
**LastConnectionTime** | Pointer to **string** | Time of the last detected connection attempt that either failed or succeeded. | [optional] 
**FormattedLastConnectionTime** | Pointer to **string** | Formatted time of the last detected connection attempt that either failed or succeeded. RFC 3339 compatible format. | [optional] 
**LastConnectionUser** | Pointer to [**MachineResponseModelAllOfLastConnectionUser**](MachineResponseModelAllOfLastConnectionUser.md) |  | [optional] 
**LastDeregistrationReason** | Pointer to [**DeregistrationReason**](DeregistrationReason.md) |  | [optional] 
**LastDeregistrationTime** | Pointer to **string** | Time of the last deregistration of the machine from the controller. | [optional] 
**FormattedLastDeregistrationTime** | Pointer to **string** | Formatted time of the last deregistration of the machine from the controller. RFC 3339 compatible format. | [optional] 
**LastErrorReason** | Pointer to **string** | The reason for the last error detected in the machine. | [optional] 
**LastErrorTime** | Pointer to **string** | The time of the last detected error. | [optional] 
**FormattedLastErrorTime** | Pointer to **string** | The formatted time of the last detected error. RFC 3339 compatible format. | [optional] 
**LoadIndex** | Pointer to **int32** | Gives current effective load index. Only used when SessionSupport is equal to MultiSession. | [optional] 
**MachineUnavailableReason** | Pointer to [**MachineUnavailableReason**](MachineUnavailableReason.md) |  | [optional] 
**OSType** | Pointer to **string** | A string that can be used to identify the operating system that is running on the machine. | [optional] 
**OSVersion** | Pointer to **string** | A string that can be used to identify the version of the operating system running on the machine, if known. | [optional] 
**PersistUserChanges** | Pointer to [**PersistChanges**](PersistChanges.md) |  | [optional] 
**PowerActionPending** | Pointer to **bool** | Indicates if there are any pending power actions for the machine. Only relevant for power-managed machines. | [optional] 
**PowerState** | [**PowerState**](PowerState.md) |  | 
**ProvisioningType** | [**ProvisioningType**](ProvisioningType.md) |  | 
**PublishedApplications** | Pointer to **[]string** | Indicates the published applications. | [optional] 
**PublishedName** | Pointer to **string** | The name of the machine that is displayed in Receiver, if the machine has been published. | [optional] 
**RegistrationState** | Pointer to [**RegistrationState**](RegistrationState.md) |  | [optional] 
**ScheduledReboot** | Pointer to [**ScheduledReboot**](ScheduledReboot.md) |  | [optional] 
**SessionClientAddress** | Pointer to **string** | The IP address of the client connected to the session. | [optional] 
**SessionClientName** | Pointer to **string** | The host name of the client connected to the session. | [optional] 
**SessionClientVersion** | Pointer to **string** | The version of the Citrix Receiver running on the client connected to the session. | [optional] 
**SessionConnectedViaHostName** | Pointer to **string** | The host name of the incoming connection. This is usually a gateway, router or client. | [optional] 
**SessionConnectedViaIP** | Pointer to **string** | The IP address of the incoming connection This is usually a gateway, router or client. | [optional] 
**SessionCount** | Pointer to **int32** | Number of sessions running on the machine. | [optional] 
**SessionLaunchedViaHostName** | Pointer to **string** | The host name of the StoreFront server used to launch the session. | [optional] 
**SessionLaunchedViaIP** | Pointer to **string** | The IP address of the StoreFront server used to launch the session. | [optional] 
**SessionProtocol** | Pointer to [**ProtocolType**](ProtocolType.md) |  | [optional] 
**SessionSecureIcaActive** | Pointer to **bool** | Indicates whether SecureICA is active on the session. | [optional] 
**SessionSmartAccessTags** | Pointer to **[]string** | The Smart Access tags for this session. | [optional] 
**SessionStartTime** | Pointer to **string** | The time indicates when the session was started. | [optional] 
**FormattedSessionStartTime** | Pointer to **string** | The formatted time indicates when the session was started. RFC 3339 compatible format. | [optional] 
**SessionState** | Pointer to [**SessionState**](SessionState.md) |  | [optional] 
**SessionStateChangeTime** | Pointer to **string** | The time of the most recent state change for the session. | [optional] 
**FormattedSessionStateChangeTime** | Pointer to **string** | The formatted time of the most recent state change for the session. RFC 3339 compatible format. | [optional] 
**SessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 
**SessionUserName** | Pointer to **string** | The session user name. | [optional] 
**Sid** | **string** | The SID of the machine. Used to be: DesktopSid or SID (based on the context) | 
**SummaryState** | [**SummaryState**](SummaryState.md) |  | 
**WillShutdownAfterUse** | Pointer to **bool** | Flag indicating if this machine is tainted and will be shut down after all sessions on the machine have ended. This flag is only ever non-null on power-managed, single-session machines. | [optional] 
**WindowsConnectionSetting** | Pointer to [**WindowsConnectionSetting**](WindowsConnectionSetting.md) |  | [optional] 
**Zone** | [**MachineResponseModelAllOfZone**](MachineResponseModelAllOfZone.md) |  | 
**SupportedPowerActions** | Pointer to [**[]SupportedPowerAction**](SupportedPowerAction.md) | A list of power actions supported by this machine. | [optional] 
**FaultState** | [**FaultState**](FaultState.md) |  | 
**ContainerMetadata** | Pointer to [**MachineResponseModelAllOfContainerMetadata**](MachineResponseModelAllOfContainerMetadata.md) |  | [optional] 
**Tags** | Pointer to **[]string** | The tags for this machine. | [optional] 
**UpgradeType** | Pointer to [**VdaUpgradeType**](VdaUpgradeType.md) |  | [optional] 
**UpgradeState** | Pointer to [**VdaUpgradeState**](VdaUpgradeState.md) |  | [optional] 
**MachineConfigurationOutOfSync** | Pointer to **bool** | Flag indicating whether the machine&#39;s configuration is out of sync with the catalog&#39;s latest configuration | [optional] 
**UpgradeDetail** | Pointer to [**MachineResponseModelAllOfUpgradeDetail**](MachineResponseModelAllOfUpgradeDetail.md) |  | [optional] 

## Methods

### NewMachineResponseModelAllOf

`func NewMachineResponseModelAllOf(inMaintenanceMode bool, drainingUntilShutdown bool, machineType MachineType, powerState PowerState, provisioningType ProvisioningType, sessionSupport SessionSupport, sid string, summaryState SummaryState, zone MachineResponseModelAllOfZone, faultState FaultState, ) *MachineResponseModelAllOf`

NewMachineResponseModelAllOf instantiates a new MachineResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineResponseModelAllOfWithDefaults

`func NewMachineResponseModelAllOfWithDefaults() *MachineResponseModelAllOf`

NewMachineResponseModelAllOfWithDefaults instantiates a new MachineResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *MachineResponseModelAllOf) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MachineResponseModelAllOf) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MachineResponseModelAllOf) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MachineResponseModelAllOf) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetAgentVersion

`func (o *MachineResponseModelAllOf) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *MachineResponseModelAllOf) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *MachineResponseModelAllOf) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *MachineResponseModelAllOf) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetAllocationType

`func (o *MachineResponseModelAllOf) GetAllocationType() AllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *MachineResponseModelAllOf) GetAllocationTypeOk() (*AllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *MachineResponseModelAllOf) SetAllocationType(v AllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *MachineResponseModelAllOf) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetApplicationsInUse

`func (o *MachineResponseModelAllOf) GetApplicationsInUse() []RefResponseModel`

GetApplicationsInUse returns the ApplicationsInUse field if non-nil, zero value otherwise.

### GetApplicationsInUseOk

`func (o *MachineResponseModelAllOf) GetApplicationsInUseOk() (*[]RefResponseModel, bool)`

GetApplicationsInUseOk returns a tuple with the ApplicationsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsInUse

`func (o *MachineResponseModelAllOf) SetApplicationsInUse(v []RefResponseModel)`

SetApplicationsInUse sets ApplicationsInUse field to given value.

### HasApplicationsInUse

`func (o *MachineResponseModelAllOf) HasApplicationsInUse() bool`

HasApplicationsInUse returns a boolean if a field has been set.

### GetAssignedUsers

`func (o *MachineResponseModelAllOf) GetAssignedUsers() []IdentityUserResponseModel`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *MachineResponseModelAllOf) GetAssignedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *MachineResponseModelAllOf) SetAssignedUsers(v []IdentityUserResponseModel)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *MachineResponseModelAllOf) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.

### GetAssociatedUsers

`func (o *MachineResponseModelAllOf) GetAssociatedUsers() []IdentityUserResponseModel`

GetAssociatedUsers returns the AssociatedUsers field if non-nil, zero value otherwise.

### GetAssociatedUsersOk

`func (o *MachineResponseModelAllOf) GetAssociatedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetAssociatedUsersOk returns a tuple with the AssociatedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedUsers

`func (o *MachineResponseModelAllOf) SetAssociatedUsers(v []IdentityUserResponseModel)`

SetAssociatedUsers sets AssociatedUsers field to given value.

### HasAssociatedUsers

`func (o *MachineResponseModelAllOf) HasAssociatedUsers() bool`

HasAssociatedUsers returns a boolean if a field has been set.

### GetAzureAdJoinedMode

`func (o *MachineResponseModelAllOf) GetAzureAdJoinedMode() AzureAdJoinedMode`

GetAzureAdJoinedMode returns the AzureAdJoinedMode field if non-nil, zero value otherwise.

### GetAzureAdJoinedModeOk

`func (o *MachineResponseModelAllOf) GetAzureAdJoinedModeOk() (*AzureAdJoinedMode, bool)`

GetAzureAdJoinedModeOk returns a tuple with the AzureAdJoinedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdJoinedMode

`func (o *MachineResponseModelAllOf) SetAzureAdJoinedMode(v AzureAdJoinedMode)`

SetAzureAdJoinedMode sets AzureAdJoinedMode field to given value.

### HasAzureAdJoinedMode

`func (o *MachineResponseModelAllOf) HasAzureAdJoinedMode() bool`

HasAzureAdJoinedMode returns a boolean if a field has been set.

### GetContainerScopes

`func (o *MachineResponseModelAllOf) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *MachineResponseModelAllOf) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *MachineResponseModelAllOf) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.

### HasContainerScopes

`func (o *MachineResponseModelAllOf) HasContainerScopes() bool`

HasContainerScopes returns a boolean if a field has been set.

### GetControllerDnsName

`func (o *MachineResponseModelAllOf) GetControllerDnsName() string`

GetControllerDnsName returns the ControllerDnsName field if non-nil, zero value otherwise.

### GetControllerDnsNameOk

`func (o *MachineResponseModelAllOf) GetControllerDnsNameOk() (*string, bool)`

GetControllerDnsNameOk returns a tuple with the ControllerDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDnsName

`func (o *MachineResponseModelAllOf) SetControllerDnsName(v string)`

SetControllerDnsName sets ControllerDnsName field to given value.

### HasControllerDnsName

`func (o *MachineResponseModelAllOf) HasControllerDnsName() bool`

HasControllerDnsName returns a boolean if a field has been set.

### GetDeliveryGroup

`func (o *MachineResponseModelAllOf) GetDeliveryGroup() MachineResponseModelAllOfDeliveryGroup`

GetDeliveryGroup returns the DeliveryGroup field if non-nil, zero value otherwise.

### GetDeliveryGroupOk

`func (o *MachineResponseModelAllOf) GetDeliveryGroupOk() (*MachineResponseModelAllOfDeliveryGroup, bool)`

GetDeliveryGroupOk returns a tuple with the DeliveryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroup

`func (o *MachineResponseModelAllOf) SetDeliveryGroup(v MachineResponseModelAllOfDeliveryGroup)`

SetDeliveryGroup sets DeliveryGroup field to given value.

### HasDeliveryGroup

`func (o *MachineResponseModelAllOf) HasDeliveryGroup() bool`

HasDeliveryGroup returns a boolean if a field has been set.

### GetDeliveryType

`func (o *MachineResponseModelAllOf) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *MachineResponseModelAllOf) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *MachineResponseModelAllOf) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *MachineResponseModelAllOf) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDescription

`func (o *MachineResponseModelAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MachineResponseModelAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MachineResponseModelAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MachineResponseModelAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesktopConditions

`func (o *MachineResponseModelAllOf) GetDesktopConditions() []DesktopCondition`

GetDesktopConditions returns the DesktopConditions field if non-nil, zero value otherwise.

### GetDesktopConditionsOk

`func (o *MachineResponseModelAllOf) GetDesktopConditionsOk() (*[]DesktopCondition, bool)`

GetDesktopConditionsOk returns a tuple with the DesktopConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopConditions

`func (o *MachineResponseModelAllOf) SetDesktopConditions(v []DesktopCondition)`

SetDesktopConditions sets DesktopConditions field to given value.

### HasDesktopConditions

`func (o *MachineResponseModelAllOf) HasDesktopConditions() bool`

HasDesktopConditions returns a boolean if a field has been set.

### GetDnsName

`func (o *MachineResponseModelAllOf) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *MachineResponseModelAllOf) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *MachineResponseModelAllOf) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *MachineResponseModelAllOf) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetFunctionalLevel

`func (o *MachineResponseModelAllOf) GetFunctionalLevel() FunctionalLevel`

GetFunctionalLevel returns the FunctionalLevel field if non-nil, zero value otherwise.

### GetFunctionalLevelOk

`func (o *MachineResponseModelAllOf) GetFunctionalLevelOk() (*FunctionalLevel, bool)`

GetFunctionalLevelOk returns a tuple with the FunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalLevel

`func (o *MachineResponseModelAllOf) SetFunctionalLevel(v FunctionalLevel)`

SetFunctionalLevel sets FunctionalLevel field to given value.

### HasFunctionalLevel

`func (o *MachineResponseModelAllOf) HasFunctionalLevel() bool`

HasFunctionalLevel returns a boolean if a field has been set.

### GetHosting

`func (o *MachineResponseModelAllOf) GetHosting() MachineResponseModelAllOfHosting`

GetHosting returns the Hosting field if non-nil, zero value otherwise.

### GetHostingOk

`func (o *MachineResponseModelAllOf) GetHostingOk() (*MachineResponseModelAllOfHosting, bool)`

GetHostingOk returns a tuple with the Hosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosting

`func (o *MachineResponseModelAllOf) SetHosting(v MachineResponseModelAllOfHosting)`

SetHosting sets Hosting field to given value.

### HasHosting

`func (o *MachineResponseModelAllOf) HasHosting() bool`

HasHosting returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *MachineResponseModelAllOf) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *MachineResponseModelAllOf) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *MachineResponseModelAllOf) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetDrainingUntilShutdown

`func (o *MachineResponseModelAllOf) GetDrainingUntilShutdown() bool`

GetDrainingUntilShutdown returns the DrainingUntilShutdown field if non-nil, zero value otherwise.

### GetDrainingUntilShutdownOk

`func (o *MachineResponseModelAllOf) GetDrainingUntilShutdownOk() (*bool, bool)`

GetDrainingUntilShutdownOk returns a tuple with the DrainingUntilShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainingUntilShutdown

`func (o *MachineResponseModelAllOf) SetDrainingUntilShutdown(v bool)`

SetDrainingUntilShutdown sets DrainingUntilShutdown field to given value.


### GetIPAddress

`func (o *MachineResponseModelAllOf) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *MachineResponseModelAllOf) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *MachineResponseModelAllOf) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *MachineResponseModelAllOf) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### GetIsAssigned

`func (o *MachineResponseModelAllOf) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MachineResponseModelAllOf) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *MachineResponseModelAllOf) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *MachineResponseModelAllOf) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetMachineType

`func (o *MachineResponseModelAllOf) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *MachineResponseModelAllOf) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *MachineResponseModelAllOf) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.


### GetLastConnectionFailure

`func (o *MachineResponseModelAllOf) GetLastConnectionFailure() ConnectionFailureReason`

GetLastConnectionFailure returns the LastConnectionFailure field if non-nil, zero value otherwise.

### GetLastConnectionFailureOk

`func (o *MachineResponseModelAllOf) GetLastConnectionFailureOk() (*ConnectionFailureReason, bool)`

GetLastConnectionFailureOk returns a tuple with the LastConnectionFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionFailure

`func (o *MachineResponseModelAllOf) SetLastConnectionFailure(v ConnectionFailureReason)`

SetLastConnectionFailure sets LastConnectionFailure field to given value.

### HasLastConnectionFailure

`func (o *MachineResponseModelAllOf) HasLastConnectionFailure() bool`

HasLastConnectionFailure returns a boolean if a field has been set.

### GetLastConnectionTime

`func (o *MachineResponseModelAllOf) GetLastConnectionTime() string`

GetLastConnectionTime returns the LastConnectionTime field if non-nil, zero value otherwise.

### GetLastConnectionTimeOk

`func (o *MachineResponseModelAllOf) GetLastConnectionTimeOk() (*string, bool)`

GetLastConnectionTimeOk returns a tuple with the LastConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionTime

`func (o *MachineResponseModelAllOf) SetLastConnectionTime(v string)`

SetLastConnectionTime sets LastConnectionTime field to given value.

### HasLastConnectionTime

`func (o *MachineResponseModelAllOf) HasLastConnectionTime() bool`

HasLastConnectionTime returns a boolean if a field has been set.

### GetFormattedLastConnectionTime

`func (o *MachineResponseModelAllOf) GetFormattedLastConnectionTime() string`

GetFormattedLastConnectionTime returns the FormattedLastConnectionTime field if non-nil, zero value otherwise.

### GetFormattedLastConnectionTimeOk

`func (o *MachineResponseModelAllOf) GetFormattedLastConnectionTimeOk() (*string, bool)`

GetFormattedLastConnectionTimeOk returns a tuple with the FormattedLastConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastConnectionTime

`func (o *MachineResponseModelAllOf) SetFormattedLastConnectionTime(v string)`

SetFormattedLastConnectionTime sets FormattedLastConnectionTime field to given value.

### HasFormattedLastConnectionTime

`func (o *MachineResponseModelAllOf) HasFormattedLastConnectionTime() bool`

HasFormattedLastConnectionTime returns a boolean if a field has been set.

### GetLastConnectionUser

`func (o *MachineResponseModelAllOf) GetLastConnectionUser() MachineResponseModelAllOfLastConnectionUser`

GetLastConnectionUser returns the LastConnectionUser field if non-nil, zero value otherwise.

### GetLastConnectionUserOk

`func (o *MachineResponseModelAllOf) GetLastConnectionUserOk() (*MachineResponseModelAllOfLastConnectionUser, bool)`

GetLastConnectionUserOk returns a tuple with the LastConnectionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionUser

`func (o *MachineResponseModelAllOf) SetLastConnectionUser(v MachineResponseModelAllOfLastConnectionUser)`

SetLastConnectionUser sets LastConnectionUser field to given value.

### HasLastConnectionUser

`func (o *MachineResponseModelAllOf) HasLastConnectionUser() bool`

HasLastConnectionUser returns a boolean if a field has been set.

### GetLastDeregistrationReason

`func (o *MachineResponseModelAllOf) GetLastDeregistrationReason() DeregistrationReason`

GetLastDeregistrationReason returns the LastDeregistrationReason field if non-nil, zero value otherwise.

### GetLastDeregistrationReasonOk

`func (o *MachineResponseModelAllOf) GetLastDeregistrationReasonOk() (*DeregistrationReason, bool)`

GetLastDeregistrationReasonOk returns a tuple with the LastDeregistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationReason

`func (o *MachineResponseModelAllOf) SetLastDeregistrationReason(v DeregistrationReason)`

SetLastDeregistrationReason sets LastDeregistrationReason field to given value.

### HasLastDeregistrationReason

`func (o *MachineResponseModelAllOf) HasLastDeregistrationReason() bool`

HasLastDeregistrationReason returns a boolean if a field has been set.

### GetLastDeregistrationTime

`func (o *MachineResponseModelAllOf) GetLastDeregistrationTime() string`

GetLastDeregistrationTime returns the LastDeregistrationTime field if non-nil, zero value otherwise.

### GetLastDeregistrationTimeOk

`func (o *MachineResponseModelAllOf) GetLastDeregistrationTimeOk() (*string, bool)`

GetLastDeregistrationTimeOk returns a tuple with the LastDeregistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationTime

`func (o *MachineResponseModelAllOf) SetLastDeregistrationTime(v string)`

SetLastDeregistrationTime sets LastDeregistrationTime field to given value.

### HasLastDeregistrationTime

`func (o *MachineResponseModelAllOf) HasLastDeregistrationTime() bool`

HasLastDeregistrationTime returns a boolean if a field has been set.

### GetFormattedLastDeregistrationTime

`func (o *MachineResponseModelAllOf) GetFormattedLastDeregistrationTime() string`

GetFormattedLastDeregistrationTime returns the FormattedLastDeregistrationTime field if non-nil, zero value otherwise.

### GetFormattedLastDeregistrationTimeOk

`func (o *MachineResponseModelAllOf) GetFormattedLastDeregistrationTimeOk() (*string, bool)`

GetFormattedLastDeregistrationTimeOk returns a tuple with the FormattedLastDeregistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastDeregistrationTime

`func (o *MachineResponseModelAllOf) SetFormattedLastDeregistrationTime(v string)`

SetFormattedLastDeregistrationTime sets FormattedLastDeregistrationTime field to given value.

### HasFormattedLastDeregistrationTime

`func (o *MachineResponseModelAllOf) HasFormattedLastDeregistrationTime() bool`

HasFormattedLastDeregistrationTime returns a boolean if a field has been set.

### GetLastErrorReason

`func (o *MachineResponseModelAllOf) GetLastErrorReason() string`

GetLastErrorReason returns the LastErrorReason field if non-nil, zero value otherwise.

### GetLastErrorReasonOk

`func (o *MachineResponseModelAllOf) GetLastErrorReasonOk() (*string, bool)`

GetLastErrorReasonOk returns a tuple with the LastErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorReason

`func (o *MachineResponseModelAllOf) SetLastErrorReason(v string)`

SetLastErrorReason sets LastErrorReason field to given value.

### HasLastErrorReason

`func (o *MachineResponseModelAllOf) HasLastErrorReason() bool`

HasLastErrorReason returns a boolean if a field has been set.

### GetLastErrorTime

`func (o *MachineResponseModelAllOf) GetLastErrorTime() string`

GetLastErrorTime returns the LastErrorTime field if non-nil, zero value otherwise.

### GetLastErrorTimeOk

`func (o *MachineResponseModelAllOf) GetLastErrorTimeOk() (*string, bool)`

GetLastErrorTimeOk returns a tuple with the LastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorTime

`func (o *MachineResponseModelAllOf) SetLastErrorTime(v string)`

SetLastErrorTime sets LastErrorTime field to given value.

### HasLastErrorTime

`func (o *MachineResponseModelAllOf) HasLastErrorTime() bool`

HasLastErrorTime returns a boolean if a field has been set.

### GetFormattedLastErrorTime

`func (o *MachineResponseModelAllOf) GetFormattedLastErrorTime() string`

GetFormattedLastErrorTime returns the FormattedLastErrorTime field if non-nil, zero value otherwise.

### GetFormattedLastErrorTimeOk

`func (o *MachineResponseModelAllOf) GetFormattedLastErrorTimeOk() (*string, bool)`

GetFormattedLastErrorTimeOk returns a tuple with the FormattedLastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastErrorTime

`func (o *MachineResponseModelAllOf) SetFormattedLastErrorTime(v string)`

SetFormattedLastErrorTime sets FormattedLastErrorTime field to given value.

### HasFormattedLastErrorTime

`func (o *MachineResponseModelAllOf) HasFormattedLastErrorTime() bool`

HasFormattedLastErrorTime returns a boolean if a field has been set.

### GetLoadIndex

`func (o *MachineResponseModelAllOf) GetLoadIndex() int32`

GetLoadIndex returns the LoadIndex field if non-nil, zero value otherwise.

### GetLoadIndexOk

`func (o *MachineResponseModelAllOf) GetLoadIndexOk() (*int32, bool)`

GetLoadIndexOk returns a tuple with the LoadIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadIndex

`func (o *MachineResponseModelAllOf) SetLoadIndex(v int32)`

SetLoadIndex sets LoadIndex field to given value.

### HasLoadIndex

`func (o *MachineResponseModelAllOf) HasLoadIndex() bool`

HasLoadIndex returns a boolean if a field has been set.

### GetMachineUnavailableReason

`func (o *MachineResponseModelAllOf) GetMachineUnavailableReason() MachineUnavailableReason`

GetMachineUnavailableReason returns the MachineUnavailableReason field if non-nil, zero value otherwise.

### GetMachineUnavailableReasonOk

`func (o *MachineResponseModelAllOf) GetMachineUnavailableReasonOk() (*MachineUnavailableReason, bool)`

GetMachineUnavailableReasonOk returns a tuple with the MachineUnavailableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineUnavailableReason

`func (o *MachineResponseModelAllOf) SetMachineUnavailableReason(v MachineUnavailableReason)`

SetMachineUnavailableReason sets MachineUnavailableReason field to given value.

### HasMachineUnavailableReason

`func (o *MachineResponseModelAllOf) HasMachineUnavailableReason() bool`

HasMachineUnavailableReason returns a boolean if a field has been set.

### GetOSType

`func (o *MachineResponseModelAllOf) GetOSType() string`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *MachineResponseModelAllOf) GetOSTypeOk() (*string, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *MachineResponseModelAllOf) SetOSType(v string)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *MachineResponseModelAllOf) HasOSType() bool`

HasOSType returns a boolean if a field has been set.

### GetOSVersion

`func (o *MachineResponseModelAllOf) GetOSVersion() string`

GetOSVersion returns the OSVersion field if non-nil, zero value otherwise.

### GetOSVersionOk

`func (o *MachineResponseModelAllOf) GetOSVersionOk() (*string, bool)`

GetOSVersionOk returns a tuple with the OSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSVersion

`func (o *MachineResponseModelAllOf) SetOSVersion(v string)`

SetOSVersion sets OSVersion field to given value.

### HasOSVersion

`func (o *MachineResponseModelAllOf) HasOSVersion() bool`

HasOSVersion returns a boolean if a field has been set.

### GetPersistUserChanges

`func (o *MachineResponseModelAllOf) GetPersistUserChanges() PersistChanges`

GetPersistUserChanges returns the PersistUserChanges field if non-nil, zero value otherwise.

### GetPersistUserChangesOk

`func (o *MachineResponseModelAllOf) GetPersistUserChangesOk() (*PersistChanges, bool)`

GetPersistUserChangesOk returns a tuple with the PersistUserChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistUserChanges

`func (o *MachineResponseModelAllOf) SetPersistUserChanges(v PersistChanges)`

SetPersistUserChanges sets PersistUserChanges field to given value.

### HasPersistUserChanges

`func (o *MachineResponseModelAllOf) HasPersistUserChanges() bool`

HasPersistUserChanges returns a boolean if a field has been set.

### GetPowerActionPending

`func (o *MachineResponseModelAllOf) GetPowerActionPending() bool`

GetPowerActionPending returns the PowerActionPending field if non-nil, zero value otherwise.

### GetPowerActionPendingOk

`func (o *MachineResponseModelAllOf) GetPowerActionPendingOk() (*bool, bool)`

GetPowerActionPendingOk returns a tuple with the PowerActionPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerActionPending

`func (o *MachineResponseModelAllOf) SetPowerActionPending(v bool)`

SetPowerActionPending sets PowerActionPending field to given value.

### HasPowerActionPending

`func (o *MachineResponseModelAllOf) HasPowerActionPending() bool`

HasPowerActionPending returns a boolean if a field has been set.

### GetPowerState

`func (o *MachineResponseModelAllOf) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *MachineResponseModelAllOf) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *MachineResponseModelAllOf) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.


### GetProvisioningType

`func (o *MachineResponseModelAllOf) GetProvisioningType() ProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *MachineResponseModelAllOf) GetProvisioningTypeOk() (*ProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *MachineResponseModelAllOf) SetProvisioningType(v ProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.


### GetPublishedApplications

`func (o *MachineResponseModelAllOf) GetPublishedApplications() []string`

GetPublishedApplications returns the PublishedApplications field if non-nil, zero value otherwise.

### GetPublishedApplicationsOk

`func (o *MachineResponseModelAllOf) GetPublishedApplicationsOk() (*[]string, bool)`

GetPublishedApplicationsOk returns a tuple with the PublishedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedApplications

`func (o *MachineResponseModelAllOf) SetPublishedApplications(v []string)`

SetPublishedApplications sets PublishedApplications field to given value.

### HasPublishedApplications

`func (o *MachineResponseModelAllOf) HasPublishedApplications() bool`

HasPublishedApplications returns a boolean if a field has been set.

### GetPublishedName

`func (o *MachineResponseModelAllOf) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *MachineResponseModelAllOf) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *MachineResponseModelAllOf) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *MachineResponseModelAllOf) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### GetRegistrationState

`func (o *MachineResponseModelAllOf) GetRegistrationState() RegistrationState`

GetRegistrationState returns the RegistrationState field if non-nil, zero value otherwise.

### GetRegistrationStateOk

`func (o *MachineResponseModelAllOf) GetRegistrationStateOk() (*RegistrationState, bool)`

GetRegistrationStateOk returns a tuple with the RegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationState

`func (o *MachineResponseModelAllOf) SetRegistrationState(v RegistrationState)`

SetRegistrationState sets RegistrationState field to given value.

### HasRegistrationState

`func (o *MachineResponseModelAllOf) HasRegistrationState() bool`

HasRegistrationState returns a boolean if a field has been set.

### GetScheduledReboot

`func (o *MachineResponseModelAllOf) GetScheduledReboot() ScheduledReboot`

GetScheduledReboot returns the ScheduledReboot field if non-nil, zero value otherwise.

### GetScheduledRebootOk

`func (o *MachineResponseModelAllOf) GetScheduledRebootOk() (*ScheduledReboot, bool)`

GetScheduledRebootOk returns a tuple with the ScheduledReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReboot

`func (o *MachineResponseModelAllOf) SetScheduledReboot(v ScheduledReboot)`

SetScheduledReboot sets ScheduledReboot field to given value.

### HasScheduledReboot

`func (o *MachineResponseModelAllOf) HasScheduledReboot() bool`

HasScheduledReboot returns a boolean if a field has been set.

### GetSessionClientAddress

`func (o *MachineResponseModelAllOf) GetSessionClientAddress() string`

GetSessionClientAddress returns the SessionClientAddress field if non-nil, zero value otherwise.

### GetSessionClientAddressOk

`func (o *MachineResponseModelAllOf) GetSessionClientAddressOk() (*string, bool)`

GetSessionClientAddressOk returns a tuple with the SessionClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientAddress

`func (o *MachineResponseModelAllOf) SetSessionClientAddress(v string)`

SetSessionClientAddress sets SessionClientAddress field to given value.

### HasSessionClientAddress

`func (o *MachineResponseModelAllOf) HasSessionClientAddress() bool`

HasSessionClientAddress returns a boolean if a field has been set.

### GetSessionClientName

`func (o *MachineResponseModelAllOf) GetSessionClientName() string`

GetSessionClientName returns the SessionClientName field if non-nil, zero value otherwise.

### GetSessionClientNameOk

`func (o *MachineResponseModelAllOf) GetSessionClientNameOk() (*string, bool)`

GetSessionClientNameOk returns a tuple with the SessionClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientName

`func (o *MachineResponseModelAllOf) SetSessionClientName(v string)`

SetSessionClientName sets SessionClientName field to given value.

### HasSessionClientName

`func (o *MachineResponseModelAllOf) HasSessionClientName() bool`

HasSessionClientName returns a boolean if a field has been set.

### GetSessionClientVersion

`func (o *MachineResponseModelAllOf) GetSessionClientVersion() string`

GetSessionClientVersion returns the SessionClientVersion field if non-nil, zero value otherwise.

### GetSessionClientVersionOk

`func (o *MachineResponseModelAllOf) GetSessionClientVersionOk() (*string, bool)`

GetSessionClientVersionOk returns a tuple with the SessionClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientVersion

`func (o *MachineResponseModelAllOf) SetSessionClientVersion(v string)`

SetSessionClientVersion sets SessionClientVersion field to given value.

### HasSessionClientVersion

`func (o *MachineResponseModelAllOf) HasSessionClientVersion() bool`

HasSessionClientVersion returns a boolean if a field has been set.

### GetSessionConnectedViaHostName

`func (o *MachineResponseModelAllOf) GetSessionConnectedViaHostName() string`

GetSessionConnectedViaHostName returns the SessionConnectedViaHostName field if non-nil, zero value otherwise.

### GetSessionConnectedViaHostNameOk

`func (o *MachineResponseModelAllOf) GetSessionConnectedViaHostNameOk() (*string, bool)`

GetSessionConnectedViaHostNameOk returns a tuple with the SessionConnectedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConnectedViaHostName

`func (o *MachineResponseModelAllOf) SetSessionConnectedViaHostName(v string)`

SetSessionConnectedViaHostName sets SessionConnectedViaHostName field to given value.

### HasSessionConnectedViaHostName

`func (o *MachineResponseModelAllOf) HasSessionConnectedViaHostName() bool`

HasSessionConnectedViaHostName returns a boolean if a field has been set.

### GetSessionConnectedViaIP

`func (o *MachineResponseModelAllOf) GetSessionConnectedViaIP() string`

GetSessionConnectedViaIP returns the SessionConnectedViaIP field if non-nil, zero value otherwise.

### GetSessionConnectedViaIPOk

`func (o *MachineResponseModelAllOf) GetSessionConnectedViaIPOk() (*string, bool)`

GetSessionConnectedViaIPOk returns a tuple with the SessionConnectedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConnectedViaIP

`func (o *MachineResponseModelAllOf) SetSessionConnectedViaIP(v string)`

SetSessionConnectedViaIP sets SessionConnectedViaIP field to given value.

### HasSessionConnectedViaIP

`func (o *MachineResponseModelAllOf) HasSessionConnectedViaIP() bool`

HasSessionConnectedViaIP returns a boolean if a field has been set.

### GetSessionCount

`func (o *MachineResponseModelAllOf) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *MachineResponseModelAllOf) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *MachineResponseModelAllOf) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *MachineResponseModelAllOf) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetSessionLaunchedViaHostName

`func (o *MachineResponseModelAllOf) GetSessionLaunchedViaHostName() string`

GetSessionLaunchedViaHostName returns the SessionLaunchedViaHostName field if non-nil, zero value otherwise.

### GetSessionLaunchedViaHostNameOk

`func (o *MachineResponseModelAllOf) GetSessionLaunchedViaHostNameOk() (*string, bool)`

GetSessionLaunchedViaHostNameOk returns a tuple with the SessionLaunchedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLaunchedViaHostName

`func (o *MachineResponseModelAllOf) SetSessionLaunchedViaHostName(v string)`

SetSessionLaunchedViaHostName sets SessionLaunchedViaHostName field to given value.

### HasSessionLaunchedViaHostName

`func (o *MachineResponseModelAllOf) HasSessionLaunchedViaHostName() bool`

HasSessionLaunchedViaHostName returns a boolean if a field has been set.

### GetSessionLaunchedViaIP

`func (o *MachineResponseModelAllOf) GetSessionLaunchedViaIP() string`

GetSessionLaunchedViaIP returns the SessionLaunchedViaIP field if non-nil, zero value otherwise.

### GetSessionLaunchedViaIPOk

`func (o *MachineResponseModelAllOf) GetSessionLaunchedViaIPOk() (*string, bool)`

GetSessionLaunchedViaIPOk returns a tuple with the SessionLaunchedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLaunchedViaIP

`func (o *MachineResponseModelAllOf) SetSessionLaunchedViaIP(v string)`

SetSessionLaunchedViaIP sets SessionLaunchedViaIP field to given value.

### HasSessionLaunchedViaIP

`func (o *MachineResponseModelAllOf) HasSessionLaunchedViaIP() bool`

HasSessionLaunchedViaIP returns a boolean if a field has been set.

### GetSessionProtocol

`func (o *MachineResponseModelAllOf) GetSessionProtocol() ProtocolType`

GetSessionProtocol returns the SessionProtocol field if non-nil, zero value otherwise.

### GetSessionProtocolOk

`func (o *MachineResponseModelAllOf) GetSessionProtocolOk() (*ProtocolType, bool)`

GetSessionProtocolOk returns a tuple with the SessionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProtocol

`func (o *MachineResponseModelAllOf) SetSessionProtocol(v ProtocolType)`

SetSessionProtocol sets SessionProtocol field to given value.

### HasSessionProtocol

`func (o *MachineResponseModelAllOf) HasSessionProtocol() bool`

HasSessionProtocol returns a boolean if a field has been set.

### GetSessionSecureIcaActive

`func (o *MachineResponseModelAllOf) GetSessionSecureIcaActive() bool`

GetSessionSecureIcaActive returns the SessionSecureIcaActive field if non-nil, zero value otherwise.

### GetSessionSecureIcaActiveOk

`func (o *MachineResponseModelAllOf) GetSessionSecureIcaActiveOk() (*bool, bool)`

GetSessionSecureIcaActiveOk returns a tuple with the SessionSecureIcaActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSecureIcaActive

`func (o *MachineResponseModelAllOf) SetSessionSecureIcaActive(v bool)`

SetSessionSecureIcaActive sets SessionSecureIcaActive field to given value.

### HasSessionSecureIcaActive

`func (o *MachineResponseModelAllOf) HasSessionSecureIcaActive() bool`

HasSessionSecureIcaActive returns a boolean if a field has been set.

### GetSessionSmartAccessTags

`func (o *MachineResponseModelAllOf) GetSessionSmartAccessTags() []string`

GetSessionSmartAccessTags returns the SessionSmartAccessTags field if non-nil, zero value otherwise.

### GetSessionSmartAccessTagsOk

`func (o *MachineResponseModelAllOf) GetSessionSmartAccessTagsOk() (*[]string, bool)`

GetSessionSmartAccessTagsOk returns a tuple with the SessionSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSmartAccessTags

`func (o *MachineResponseModelAllOf) SetSessionSmartAccessTags(v []string)`

SetSessionSmartAccessTags sets SessionSmartAccessTags field to given value.

### HasSessionSmartAccessTags

`func (o *MachineResponseModelAllOf) HasSessionSmartAccessTags() bool`

HasSessionSmartAccessTags returns a boolean if a field has been set.

### GetSessionStartTime

`func (o *MachineResponseModelAllOf) GetSessionStartTime() string`

GetSessionStartTime returns the SessionStartTime field if non-nil, zero value otherwise.

### GetSessionStartTimeOk

`func (o *MachineResponseModelAllOf) GetSessionStartTimeOk() (*string, bool)`

GetSessionStartTimeOk returns a tuple with the SessionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStartTime

`func (o *MachineResponseModelAllOf) SetSessionStartTime(v string)`

SetSessionStartTime sets SessionStartTime field to given value.

### HasSessionStartTime

`func (o *MachineResponseModelAllOf) HasSessionStartTime() bool`

HasSessionStartTime returns a boolean if a field has been set.

### GetFormattedSessionStartTime

`func (o *MachineResponseModelAllOf) GetFormattedSessionStartTime() string`

GetFormattedSessionStartTime returns the FormattedSessionStartTime field if non-nil, zero value otherwise.

### GetFormattedSessionStartTimeOk

`func (o *MachineResponseModelAllOf) GetFormattedSessionStartTimeOk() (*string, bool)`

GetFormattedSessionStartTimeOk returns a tuple with the FormattedSessionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSessionStartTime

`func (o *MachineResponseModelAllOf) SetFormattedSessionStartTime(v string)`

SetFormattedSessionStartTime sets FormattedSessionStartTime field to given value.

### HasFormattedSessionStartTime

`func (o *MachineResponseModelAllOf) HasFormattedSessionStartTime() bool`

HasFormattedSessionStartTime returns a boolean if a field has been set.

### GetSessionState

`func (o *MachineResponseModelAllOf) GetSessionState() SessionState`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *MachineResponseModelAllOf) GetSessionStateOk() (*SessionState, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *MachineResponseModelAllOf) SetSessionState(v SessionState)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *MachineResponseModelAllOf) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### GetSessionStateChangeTime

`func (o *MachineResponseModelAllOf) GetSessionStateChangeTime() string`

GetSessionStateChangeTime returns the SessionStateChangeTime field if non-nil, zero value otherwise.

### GetSessionStateChangeTimeOk

`func (o *MachineResponseModelAllOf) GetSessionStateChangeTimeOk() (*string, bool)`

GetSessionStateChangeTimeOk returns a tuple with the SessionStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStateChangeTime

`func (o *MachineResponseModelAllOf) SetSessionStateChangeTime(v string)`

SetSessionStateChangeTime sets SessionStateChangeTime field to given value.

### HasSessionStateChangeTime

`func (o *MachineResponseModelAllOf) HasSessionStateChangeTime() bool`

HasSessionStateChangeTime returns a boolean if a field has been set.

### GetFormattedSessionStateChangeTime

`func (o *MachineResponseModelAllOf) GetFormattedSessionStateChangeTime() string`

GetFormattedSessionStateChangeTime returns the FormattedSessionStateChangeTime field if non-nil, zero value otherwise.

### GetFormattedSessionStateChangeTimeOk

`func (o *MachineResponseModelAllOf) GetFormattedSessionStateChangeTimeOk() (*string, bool)`

GetFormattedSessionStateChangeTimeOk returns a tuple with the FormattedSessionStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSessionStateChangeTime

`func (o *MachineResponseModelAllOf) SetFormattedSessionStateChangeTime(v string)`

SetFormattedSessionStateChangeTime sets FormattedSessionStateChangeTime field to given value.

### HasFormattedSessionStateChangeTime

`func (o *MachineResponseModelAllOf) HasFormattedSessionStateChangeTime() bool`

HasFormattedSessionStateChangeTime returns a boolean if a field has been set.

### GetSessionSupport

`func (o *MachineResponseModelAllOf) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *MachineResponseModelAllOf) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *MachineResponseModelAllOf) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSessionUserName

`func (o *MachineResponseModelAllOf) GetSessionUserName() string`

GetSessionUserName returns the SessionUserName field if non-nil, zero value otherwise.

### GetSessionUserNameOk

`func (o *MachineResponseModelAllOf) GetSessionUserNameOk() (*string, bool)`

GetSessionUserNameOk returns a tuple with the SessionUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUserName

`func (o *MachineResponseModelAllOf) SetSessionUserName(v string)`

SetSessionUserName sets SessionUserName field to given value.

### HasSessionUserName

`func (o *MachineResponseModelAllOf) HasSessionUserName() bool`

HasSessionUserName returns a boolean if a field has been set.

### GetSid

`func (o *MachineResponseModelAllOf) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MachineResponseModelAllOf) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MachineResponseModelAllOf) SetSid(v string)`

SetSid sets Sid field to given value.


### GetSummaryState

`func (o *MachineResponseModelAllOf) GetSummaryState() SummaryState`

GetSummaryState returns the SummaryState field if non-nil, zero value otherwise.

### GetSummaryStateOk

`func (o *MachineResponseModelAllOf) GetSummaryStateOk() (*SummaryState, bool)`

GetSummaryStateOk returns a tuple with the SummaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryState

`func (o *MachineResponseModelAllOf) SetSummaryState(v SummaryState)`

SetSummaryState sets SummaryState field to given value.


### GetWillShutdownAfterUse

`func (o *MachineResponseModelAllOf) GetWillShutdownAfterUse() bool`

GetWillShutdownAfterUse returns the WillShutdownAfterUse field if non-nil, zero value otherwise.

### GetWillShutdownAfterUseOk

`func (o *MachineResponseModelAllOf) GetWillShutdownAfterUseOk() (*bool, bool)`

GetWillShutdownAfterUseOk returns a tuple with the WillShutdownAfterUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillShutdownAfterUse

`func (o *MachineResponseModelAllOf) SetWillShutdownAfterUse(v bool)`

SetWillShutdownAfterUse sets WillShutdownAfterUse field to given value.

### HasWillShutdownAfterUse

`func (o *MachineResponseModelAllOf) HasWillShutdownAfterUse() bool`

HasWillShutdownAfterUse returns a boolean if a field has been set.

### GetWindowsConnectionSetting

`func (o *MachineResponseModelAllOf) GetWindowsConnectionSetting() WindowsConnectionSetting`

GetWindowsConnectionSetting returns the WindowsConnectionSetting field if non-nil, zero value otherwise.

### GetWindowsConnectionSettingOk

`func (o *MachineResponseModelAllOf) GetWindowsConnectionSettingOk() (*WindowsConnectionSetting, bool)`

GetWindowsConnectionSettingOk returns a tuple with the WindowsConnectionSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsConnectionSetting

`func (o *MachineResponseModelAllOf) SetWindowsConnectionSetting(v WindowsConnectionSetting)`

SetWindowsConnectionSetting sets WindowsConnectionSetting field to given value.

### HasWindowsConnectionSetting

`func (o *MachineResponseModelAllOf) HasWindowsConnectionSetting() bool`

HasWindowsConnectionSetting returns a boolean if a field has been set.

### GetZone

`func (o *MachineResponseModelAllOf) GetZone() MachineResponseModelAllOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *MachineResponseModelAllOf) GetZoneOk() (*MachineResponseModelAllOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *MachineResponseModelAllOf) SetZone(v MachineResponseModelAllOfZone)`

SetZone sets Zone field to given value.


### GetSupportedPowerActions

`func (o *MachineResponseModelAllOf) GetSupportedPowerActions() []SupportedPowerAction`

GetSupportedPowerActions returns the SupportedPowerActions field if non-nil, zero value otherwise.

### GetSupportedPowerActionsOk

`func (o *MachineResponseModelAllOf) GetSupportedPowerActionsOk() (*[]SupportedPowerAction, bool)`

GetSupportedPowerActionsOk returns a tuple with the SupportedPowerActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPowerActions

`func (o *MachineResponseModelAllOf) SetSupportedPowerActions(v []SupportedPowerAction)`

SetSupportedPowerActions sets SupportedPowerActions field to given value.

### HasSupportedPowerActions

`func (o *MachineResponseModelAllOf) HasSupportedPowerActions() bool`

HasSupportedPowerActions returns a boolean if a field has been set.

### GetFaultState

`func (o *MachineResponseModelAllOf) GetFaultState() FaultState`

GetFaultState returns the FaultState field if non-nil, zero value otherwise.

### GetFaultStateOk

`func (o *MachineResponseModelAllOf) GetFaultStateOk() (*FaultState, bool)`

GetFaultStateOk returns a tuple with the FaultState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultState

`func (o *MachineResponseModelAllOf) SetFaultState(v FaultState)`

SetFaultState sets FaultState field to given value.


### GetContainerMetadata

`func (o *MachineResponseModelAllOf) GetContainerMetadata() MachineResponseModelAllOfContainerMetadata`

GetContainerMetadata returns the ContainerMetadata field if non-nil, zero value otherwise.

### GetContainerMetadataOk

`func (o *MachineResponseModelAllOf) GetContainerMetadataOk() (*MachineResponseModelAllOfContainerMetadata, bool)`

GetContainerMetadataOk returns a tuple with the ContainerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMetadata

`func (o *MachineResponseModelAllOf) SetContainerMetadata(v MachineResponseModelAllOfContainerMetadata)`

SetContainerMetadata sets ContainerMetadata field to given value.

### HasContainerMetadata

`func (o *MachineResponseModelAllOf) HasContainerMetadata() bool`

HasContainerMetadata returns a boolean if a field has been set.

### GetTags

`func (o *MachineResponseModelAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MachineResponseModelAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MachineResponseModelAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MachineResponseModelAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpgradeType

`func (o *MachineResponseModelAllOf) GetUpgradeType() VdaUpgradeType`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *MachineResponseModelAllOf) GetUpgradeTypeOk() (*VdaUpgradeType, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *MachineResponseModelAllOf) SetUpgradeType(v VdaUpgradeType)`

SetUpgradeType sets UpgradeType field to given value.

### HasUpgradeType

`func (o *MachineResponseModelAllOf) HasUpgradeType() bool`

HasUpgradeType returns a boolean if a field has been set.

### GetUpgradeState

`func (o *MachineResponseModelAllOf) GetUpgradeState() VdaUpgradeState`

GetUpgradeState returns the UpgradeState field if non-nil, zero value otherwise.

### GetUpgradeStateOk

`func (o *MachineResponseModelAllOf) GetUpgradeStateOk() (*VdaUpgradeState, bool)`

GetUpgradeStateOk returns a tuple with the UpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeState

`func (o *MachineResponseModelAllOf) SetUpgradeState(v VdaUpgradeState)`

SetUpgradeState sets UpgradeState field to given value.

### HasUpgradeState

`func (o *MachineResponseModelAllOf) HasUpgradeState() bool`

HasUpgradeState returns a boolean if a field has been set.

### GetMachineConfigurationOutOfSync

`func (o *MachineResponseModelAllOf) GetMachineConfigurationOutOfSync() bool`

GetMachineConfigurationOutOfSync returns the MachineConfigurationOutOfSync field if non-nil, zero value otherwise.

### GetMachineConfigurationOutOfSyncOk

`func (o *MachineResponseModelAllOf) GetMachineConfigurationOutOfSyncOk() (*bool, bool)`

GetMachineConfigurationOutOfSyncOk returns a tuple with the MachineConfigurationOutOfSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineConfigurationOutOfSync

`func (o *MachineResponseModelAllOf) SetMachineConfigurationOutOfSync(v bool)`

SetMachineConfigurationOutOfSync sets MachineConfigurationOutOfSync field to given value.

### HasMachineConfigurationOutOfSync

`func (o *MachineResponseModelAllOf) HasMachineConfigurationOutOfSync() bool`

HasMachineConfigurationOutOfSync returns a boolean if a field has been set.

### GetUpgradeDetail

`func (o *MachineResponseModelAllOf) GetUpgradeDetail() MachineResponseModelAllOfUpgradeDetail`

GetUpgradeDetail returns the UpgradeDetail field if non-nil, zero value otherwise.

### GetUpgradeDetailOk

`func (o *MachineResponseModelAllOf) GetUpgradeDetailOk() (*MachineResponseModelAllOfUpgradeDetail, bool)`

GetUpgradeDetailOk returns a tuple with the UpgradeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDetail

`func (o *MachineResponseModelAllOf) SetUpgradeDetail(v MachineResponseModelAllOfUpgradeDetail)`

SetUpgradeDetail sets UpgradeDetail field to given value.

### HasUpgradeDetail

`func (o *MachineResponseModelAllOf) HasUpgradeDetail() bool`

HasUpgradeDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


