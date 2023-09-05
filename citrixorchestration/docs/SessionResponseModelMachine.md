# SessionResponseModelMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of machine. Used to be: DesktopUid (and wasn&#39;t globally unique) OR UUID, depending on context Needs to be globally unique Might be constructed from site ID + internal Uid?  or use uuid | 
**MachineCatalog** | Pointer to [**MachineBaseResponseModelMachineCatalog**](MachineBaseResponseModelMachineCatalog.md) |  | [optional] 
**Name** | Pointer to **string** | DNS host name of the machine. Used to be: MachineName | [optional] 
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

### NewSessionResponseModelMachine

`func NewSessionResponseModelMachine(id string, inMaintenanceMode bool, drainingUntilShutdown bool, machineType MachineType, powerState PowerState, provisioningType ProvisioningType, sessionSupport SessionSupport, sid string, summaryState SummaryState, zone MachineResponseModelAllOfZone, faultState FaultState, ) *SessionResponseModelMachine`

NewSessionResponseModelMachine instantiates a new SessionResponseModelMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionResponseModelMachineWithDefaults

`func NewSessionResponseModelMachineWithDefaults() *SessionResponseModelMachine`

NewSessionResponseModelMachineWithDefaults instantiates a new SessionResponseModelMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionResponseModelMachine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionResponseModelMachine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionResponseModelMachine) SetId(v string)`

SetId sets Id field to given value.


### GetMachineCatalog

`func (o *SessionResponseModelMachine) GetMachineCatalog() MachineBaseResponseModelMachineCatalog`

GetMachineCatalog returns the MachineCatalog field if non-nil, zero value otherwise.

### GetMachineCatalogOk

`func (o *SessionResponseModelMachine) GetMachineCatalogOk() (*MachineBaseResponseModelMachineCatalog, bool)`

GetMachineCatalogOk returns a tuple with the MachineCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCatalog

`func (o *SessionResponseModelMachine) SetMachineCatalog(v MachineBaseResponseModelMachineCatalog)`

SetMachineCatalog sets MachineCatalog field to given value.

### HasMachineCatalog

`func (o *SessionResponseModelMachine) HasMachineCatalog() bool`

HasMachineCatalog returns a boolean if a field has been set.

### GetName

`func (o *SessionResponseModelMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionResponseModelMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionResponseModelMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SessionResponseModelMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUid

`func (o *SessionResponseModelMachine) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SessionResponseModelMachine) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SessionResponseModelMachine) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SessionResponseModelMachine) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetAgentVersion

`func (o *SessionResponseModelMachine) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *SessionResponseModelMachine) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *SessionResponseModelMachine) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *SessionResponseModelMachine) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetAllocationType

`func (o *SessionResponseModelMachine) GetAllocationType() AllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *SessionResponseModelMachine) GetAllocationTypeOk() (*AllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *SessionResponseModelMachine) SetAllocationType(v AllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *SessionResponseModelMachine) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetApplicationsInUse

`func (o *SessionResponseModelMachine) GetApplicationsInUse() []RefResponseModel`

GetApplicationsInUse returns the ApplicationsInUse field if non-nil, zero value otherwise.

### GetApplicationsInUseOk

`func (o *SessionResponseModelMachine) GetApplicationsInUseOk() (*[]RefResponseModel, bool)`

GetApplicationsInUseOk returns a tuple with the ApplicationsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsInUse

`func (o *SessionResponseModelMachine) SetApplicationsInUse(v []RefResponseModel)`

SetApplicationsInUse sets ApplicationsInUse field to given value.

### HasApplicationsInUse

`func (o *SessionResponseModelMachine) HasApplicationsInUse() bool`

HasApplicationsInUse returns a boolean if a field has been set.

### GetAssignedUsers

`func (o *SessionResponseModelMachine) GetAssignedUsers() []IdentityUserResponseModel`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *SessionResponseModelMachine) GetAssignedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *SessionResponseModelMachine) SetAssignedUsers(v []IdentityUserResponseModel)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *SessionResponseModelMachine) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.

### GetAssociatedUsers

`func (o *SessionResponseModelMachine) GetAssociatedUsers() []IdentityUserResponseModel`

GetAssociatedUsers returns the AssociatedUsers field if non-nil, zero value otherwise.

### GetAssociatedUsersOk

`func (o *SessionResponseModelMachine) GetAssociatedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetAssociatedUsersOk returns a tuple with the AssociatedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedUsers

`func (o *SessionResponseModelMachine) SetAssociatedUsers(v []IdentityUserResponseModel)`

SetAssociatedUsers sets AssociatedUsers field to given value.

### HasAssociatedUsers

`func (o *SessionResponseModelMachine) HasAssociatedUsers() bool`

HasAssociatedUsers returns a boolean if a field has been set.

### GetAzureAdJoinedMode

`func (o *SessionResponseModelMachine) GetAzureAdJoinedMode() AzureAdJoinedMode`

GetAzureAdJoinedMode returns the AzureAdJoinedMode field if non-nil, zero value otherwise.

### GetAzureAdJoinedModeOk

`func (o *SessionResponseModelMachine) GetAzureAdJoinedModeOk() (*AzureAdJoinedMode, bool)`

GetAzureAdJoinedModeOk returns a tuple with the AzureAdJoinedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdJoinedMode

`func (o *SessionResponseModelMachine) SetAzureAdJoinedMode(v AzureAdJoinedMode)`

SetAzureAdJoinedMode sets AzureAdJoinedMode field to given value.

### HasAzureAdJoinedMode

`func (o *SessionResponseModelMachine) HasAzureAdJoinedMode() bool`

HasAzureAdJoinedMode returns a boolean if a field has been set.

### GetContainerScopes

`func (o *SessionResponseModelMachine) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *SessionResponseModelMachine) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *SessionResponseModelMachine) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.

### HasContainerScopes

`func (o *SessionResponseModelMachine) HasContainerScopes() bool`

HasContainerScopes returns a boolean if a field has been set.

### GetControllerDnsName

`func (o *SessionResponseModelMachine) GetControllerDnsName() string`

GetControllerDnsName returns the ControllerDnsName field if non-nil, zero value otherwise.

### GetControllerDnsNameOk

`func (o *SessionResponseModelMachine) GetControllerDnsNameOk() (*string, bool)`

GetControllerDnsNameOk returns a tuple with the ControllerDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDnsName

`func (o *SessionResponseModelMachine) SetControllerDnsName(v string)`

SetControllerDnsName sets ControllerDnsName field to given value.

### HasControllerDnsName

`func (o *SessionResponseModelMachine) HasControllerDnsName() bool`

HasControllerDnsName returns a boolean if a field has been set.

### GetDeliveryGroup

`func (o *SessionResponseModelMachine) GetDeliveryGroup() MachineResponseModelAllOfDeliveryGroup`

GetDeliveryGroup returns the DeliveryGroup field if non-nil, zero value otherwise.

### GetDeliveryGroupOk

`func (o *SessionResponseModelMachine) GetDeliveryGroupOk() (*MachineResponseModelAllOfDeliveryGroup, bool)`

GetDeliveryGroupOk returns a tuple with the DeliveryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroup

`func (o *SessionResponseModelMachine) SetDeliveryGroup(v MachineResponseModelAllOfDeliveryGroup)`

SetDeliveryGroup sets DeliveryGroup field to given value.

### HasDeliveryGroup

`func (o *SessionResponseModelMachine) HasDeliveryGroup() bool`

HasDeliveryGroup returns a boolean if a field has been set.

### GetDeliveryType

`func (o *SessionResponseModelMachine) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *SessionResponseModelMachine) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *SessionResponseModelMachine) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *SessionResponseModelMachine) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDescription

`func (o *SessionResponseModelMachine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SessionResponseModelMachine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SessionResponseModelMachine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SessionResponseModelMachine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesktopConditions

`func (o *SessionResponseModelMachine) GetDesktopConditions() []DesktopCondition`

GetDesktopConditions returns the DesktopConditions field if non-nil, zero value otherwise.

### GetDesktopConditionsOk

`func (o *SessionResponseModelMachine) GetDesktopConditionsOk() (*[]DesktopCondition, bool)`

GetDesktopConditionsOk returns a tuple with the DesktopConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopConditions

`func (o *SessionResponseModelMachine) SetDesktopConditions(v []DesktopCondition)`

SetDesktopConditions sets DesktopConditions field to given value.

### HasDesktopConditions

`func (o *SessionResponseModelMachine) HasDesktopConditions() bool`

HasDesktopConditions returns a boolean if a field has been set.

### GetDnsName

`func (o *SessionResponseModelMachine) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *SessionResponseModelMachine) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *SessionResponseModelMachine) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *SessionResponseModelMachine) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetFunctionalLevel

`func (o *SessionResponseModelMachine) GetFunctionalLevel() FunctionalLevel`

GetFunctionalLevel returns the FunctionalLevel field if non-nil, zero value otherwise.

### GetFunctionalLevelOk

`func (o *SessionResponseModelMachine) GetFunctionalLevelOk() (*FunctionalLevel, bool)`

GetFunctionalLevelOk returns a tuple with the FunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalLevel

`func (o *SessionResponseModelMachine) SetFunctionalLevel(v FunctionalLevel)`

SetFunctionalLevel sets FunctionalLevel field to given value.

### HasFunctionalLevel

`func (o *SessionResponseModelMachine) HasFunctionalLevel() bool`

HasFunctionalLevel returns a boolean if a field has been set.

### GetHosting

`func (o *SessionResponseModelMachine) GetHosting() MachineResponseModelAllOfHosting`

GetHosting returns the Hosting field if non-nil, zero value otherwise.

### GetHostingOk

`func (o *SessionResponseModelMachine) GetHostingOk() (*MachineResponseModelAllOfHosting, bool)`

GetHostingOk returns a tuple with the Hosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosting

`func (o *SessionResponseModelMachine) SetHosting(v MachineResponseModelAllOfHosting)`

SetHosting sets Hosting field to given value.

### HasHosting

`func (o *SessionResponseModelMachine) HasHosting() bool`

HasHosting returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *SessionResponseModelMachine) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *SessionResponseModelMachine) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *SessionResponseModelMachine) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetDrainingUntilShutdown

`func (o *SessionResponseModelMachine) GetDrainingUntilShutdown() bool`

GetDrainingUntilShutdown returns the DrainingUntilShutdown field if non-nil, zero value otherwise.

### GetDrainingUntilShutdownOk

`func (o *SessionResponseModelMachine) GetDrainingUntilShutdownOk() (*bool, bool)`

GetDrainingUntilShutdownOk returns a tuple with the DrainingUntilShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainingUntilShutdown

`func (o *SessionResponseModelMachine) SetDrainingUntilShutdown(v bool)`

SetDrainingUntilShutdown sets DrainingUntilShutdown field to given value.


### GetIPAddress

`func (o *SessionResponseModelMachine) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *SessionResponseModelMachine) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *SessionResponseModelMachine) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *SessionResponseModelMachine) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### GetIsAssigned

`func (o *SessionResponseModelMachine) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *SessionResponseModelMachine) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *SessionResponseModelMachine) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *SessionResponseModelMachine) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetMachineType

`func (o *SessionResponseModelMachine) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *SessionResponseModelMachine) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *SessionResponseModelMachine) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.


### GetLastConnectionFailure

`func (o *SessionResponseModelMachine) GetLastConnectionFailure() ConnectionFailureReason`

GetLastConnectionFailure returns the LastConnectionFailure field if non-nil, zero value otherwise.

### GetLastConnectionFailureOk

`func (o *SessionResponseModelMachine) GetLastConnectionFailureOk() (*ConnectionFailureReason, bool)`

GetLastConnectionFailureOk returns a tuple with the LastConnectionFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionFailure

`func (o *SessionResponseModelMachine) SetLastConnectionFailure(v ConnectionFailureReason)`

SetLastConnectionFailure sets LastConnectionFailure field to given value.

### HasLastConnectionFailure

`func (o *SessionResponseModelMachine) HasLastConnectionFailure() bool`

HasLastConnectionFailure returns a boolean if a field has been set.

### GetLastConnectionTime

`func (o *SessionResponseModelMachine) GetLastConnectionTime() string`

GetLastConnectionTime returns the LastConnectionTime field if non-nil, zero value otherwise.

### GetLastConnectionTimeOk

`func (o *SessionResponseModelMachine) GetLastConnectionTimeOk() (*string, bool)`

GetLastConnectionTimeOk returns a tuple with the LastConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionTime

`func (o *SessionResponseModelMachine) SetLastConnectionTime(v string)`

SetLastConnectionTime sets LastConnectionTime field to given value.

### HasLastConnectionTime

`func (o *SessionResponseModelMachine) HasLastConnectionTime() bool`

HasLastConnectionTime returns a boolean if a field has been set.

### GetFormattedLastConnectionTime

`func (o *SessionResponseModelMachine) GetFormattedLastConnectionTime() string`

GetFormattedLastConnectionTime returns the FormattedLastConnectionTime field if non-nil, zero value otherwise.

### GetFormattedLastConnectionTimeOk

`func (o *SessionResponseModelMachine) GetFormattedLastConnectionTimeOk() (*string, bool)`

GetFormattedLastConnectionTimeOk returns a tuple with the FormattedLastConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastConnectionTime

`func (o *SessionResponseModelMachine) SetFormattedLastConnectionTime(v string)`

SetFormattedLastConnectionTime sets FormattedLastConnectionTime field to given value.

### HasFormattedLastConnectionTime

`func (o *SessionResponseModelMachine) HasFormattedLastConnectionTime() bool`

HasFormattedLastConnectionTime returns a boolean if a field has been set.

### GetLastConnectionUser

`func (o *SessionResponseModelMachine) GetLastConnectionUser() MachineResponseModelAllOfLastConnectionUser`

GetLastConnectionUser returns the LastConnectionUser field if non-nil, zero value otherwise.

### GetLastConnectionUserOk

`func (o *SessionResponseModelMachine) GetLastConnectionUserOk() (*MachineResponseModelAllOfLastConnectionUser, bool)`

GetLastConnectionUserOk returns a tuple with the LastConnectionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionUser

`func (o *SessionResponseModelMachine) SetLastConnectionUser(v MachineResponseModelAllOfLastConnectionUser)`

SetLastConnectionUser sets LastConnectionUser field to given value.

### HasLastConnectionUser

`func (o *SessionResponseModelMachine) HasLastConnectionUser() bool`

HasLastConnectionUser returns a boolean if a field has been set.

### GetLastDeregistrationReason

`func (o *SessionResponseModelMachine) GetLastDeregistrationReason() DeregistrationReason`

GetLastDeregistrationReason returns the LastDeregistrationReason field if non-nil, zero value otherwise.

### GetLastDeregistrationReasonOk

`func (o *SessionResponseModelMachine) GetLastDeregistrationReasonOk() (*DeregistrationReason, bool)`

GetLastDeregistrationReasonOk returns a tuple with the LastDeregistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationReason

`func (o *SessionResponseModelMachine) SetLastDeregistrationReason(v DeregistrationReason)`

SetLastDeregistrationReason sets LastDeregistrationReason field to given value.

### HasLastDeregistrationReason

`func (o *SessionResponseModelMachine) HasLastDeregistrationReason() bool`

HasLastDeregistrationReason returns a boolean if a field has been set.

### GetLastDeregistrationTime

`func (o *SessionResponseModelMachine) GetLastDeregistrationTime() string`

GetLastDeregistrationTime returns the LastDeregistrationTime field if non-nil, zero value otherwise.

### GetLastDeregistrationTimeOk

`func (o *SessionResponseModelMachine) GetLastDeregistrationTimeOk() (*string, bool)`

GetLastDeregistrationTimeOk returns a tuple with the LastDeregistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationTime

`func (o *SessionResponseModelMachine) SetLastDeregistrationTime(v string)`

SetLastDeregistrationTime sets LastDeregistrationTime field to given value.

### HasLastDeregistrationTime

`func (o *SessionResponseModelMachine) HasLastDeregistrationTime() bool`

HasLastDeregistrationTime returns a boolean if a field has been set.

### GetFormattedLastDeregistrationTime

`func (o *SessionResponseModelMachine) GetFormattedLastDeregistrationTime() string`

GetFormattedLastDeregistrationTime returns the FormattedLastDeregistrationTime field if non-nil, zero value otherwise.

### GetFormattedLastDeregistrationTimeOk

`func (o *SessionResponseModelMachine) GetFormattedLastDeregistrationTimeOk() (*string, bool)`

GetFormattedLastDeregistrationTimeOk returns a tuple with the FormattedLastDeregistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastDeregistrationTime

`func (o *SessionResponseModelMachine) SetFormattedLastDeregistrationTime(v string)`

SetFormattedLastDeregistrationTime sets FormattedLastDeregistrationTime field to given value.

### HasFormattedLastDeregistrationTime

`func (o *SessionResponseModelMachine) HasFormattedLastDeregistrationTime() bool`

HasFormattedLastDeregistrationTime returns a boolean if a field has been set.

### GetLastErrorReason

`func (o *SessionResponseModelMachine) GetLastErrorReason() string`

GetLastErrorReason returns the LastErrorReason field if non-nil, zero value otherwise.

### GetLastErrorReasonOk

`func (o *SessionResponseModelMachine) GetLastErrorReasonOk() (*string, bool)`

GetLastErrorReasonOk returns a tuple with the LastErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorReason

`func (o *SessionResponseModelMachine) SetLastErrorReason(v string)`

SetLastErrorReason sets LastErrorReason field to given value.

### HasLastErrorReason

`func (o *SessionResponseModelMachine) HasLastErrorReason() bool`

HasLastErrorReason returns a boolean if a field has been set.

### GetLastErrorTime

`func (o *SessionResponseModelMachine) GetLastErrorTime() string`

GetLastErrorTime returns the LastErrorTime field if non-nil, zero value otherwise.

### GetLastErrorTimeOk

`func (o *SessionResponseModelMachine) GetLastErrorTimeOk() (*string, bool)`

GetLastErrorTimeOk returns a tuple with the LastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorTime

`func (o *SessionResponseModelMachine) SetLastErrorTime(v string)`

SetLastErrorTime sets LastErrorTime field to given value.

### HasLastErrorTime

`func (o *SessionResponseModelMachine) HasLastErrorTime() bool`

HasLastErrorTime returns a boolean if a field has been set.

### GetFormattedLastErrorTime

`func (o *SessionResponseModelMachine) GetFormattedLastErrorTime() string`

GetFormattedLastErrorTime returns the FormattedLastErrorTime field if non-nil, zero value otherwise.

### GetFormattedLastErrorTimeOk

`func (o *SessionResponseModelMachine) GetFormattedLastErrorTimeOk() (*string, bool)`

GetFormattedLastErrorTimeOk returns a tuple with the FormattedLastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastErrorTime

`func (o *SessionResponseModelMachine) SetFormattedLastErrorTime(v string)`

SetFormattedLastErrorTime sets FormattedLastErrorTime field to given value.

### HasFormattedLastErrorTime

`func (o *SessionResponseModelMachine) HasFormattedLastErrorTime() bool`

HasFormattedLastErrorTime returns a boolean if a field has been set.

### GetLoadIndex

`func (o *SessionResponseModelMachine) GetLoadIndex() int32`

GetLoadIndex returns the LoadIndex field if non-nil, zero value otherwise.

### GetLoadIndexOk

`func (o *SessionResponseModelMachine) GetLoadIndexOk() (*int32, bool)`

GetLoadIndexOk returns a tuple with the LoadIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadIndex

`func (o *SessionResponseModelMachine) SetLoadIndex(v int32)`

SetLoadIndex sets LoadIndex field to given value.

### HasLoadIndex

`func (o *SessionResponseModelMachine) HasLoadIndex() bool`

HasLoadIndex returns a boolean if a field has been set.

### GetMachineUnavailableReason

`func (o *SessionResponseModelMachine) GetMachineUnavailableReason() MachineUnavailableReason`

GetMachineUnavailableReason returns the MachineUnavailableReason field if non-nil, zero value otherwise.

### GetMachineUnavailableReasonOk

`func (o *SessionResponseModelMachine) GetMachineUnavailableReasonOk() (*MachineUnavailableReason, bool)`

GetMachineUnavailableReasonOk returns a tuple with the MachineUnavailableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineUnavailableReason

`func (o *SessionResponseModelMachine) SetMachineUnavailableReason(v MachineUnavailableReason)`

SetMachineUnavailableReason sets MachineUnavailableReason field to given value.

### HasMachineUnavailableReason

`func (o *SessionResponseModelMachine) HasMachineUnavailableReason() bool`

HasMachineUnavailableReason returns a boolean if a field has been set.

### GetOSType

`func (o *SessionResponseModelMachine) GetOSType() string`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *SessionResponseModelMachine) GetOSTypeOk() (*string, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *SessionResponseModelMachine) SetOSType(v string)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *SessionResponseModelMachine) HasOSType() bool`

HasOSType returns a boolean if a field has been set.

### GetOSVersion

`func (o *SessionResponseModelMachine) GetOSVersion() string`

GetOSVersion returns the OSVersion field if non-nil, zero value otherwise.

### GetOSVersionOk

`func (o *SessionResponseModelMachine) GetOSVersionOk() (*string, bool)`

GetOSVersionOk returns a tuple with the OSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSVersion

`func (o *SessionResponseModelMachine) SetOSVersion(v string)`

SetOSVersion sets OSVersion field to given value.

### HasOSVersion

`func (o *SessionResponseModelMachine) HasOSVersion() bool`

HasOSVersion returns a boolean if a field has been set.

### GetPersistUserChanges

`func (o *SessionResponseModelMachine) GetPersistUserChanges() PersistChanges`

GetPersistUserChanges returns the PersistUserChanges field if non-nil, zero value otherwise.

### GetPersistUserChangesOk

`func (o *SessionResponseModelMachine) GetPersistUserChangesOk() (*PersistChanges, bool)`

GetPersistUserChangesOk returns a tuple with the PersistUserChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistUserChanges

`func (o *SessionResponseModelMachine) SetPersistUserChanges(v PersistChanges)`

SetPersistUserChanges sets PersistUserChanges field to given value.

### HasPersistUserChanges

`func (o *SessionResponseModelMachine) HasPersistUserChanges() bool`

HasPersistUserChanges returns a boolean if a field has been set.

### GetPowerActionPending

`func (o *SessionResponseModelMachine) GetPowerActionPending() bool`

GetPowerActionPending returns the PowerActionPending field if non-nil, zero value otherwise.

### GetPowerActionPendingOk

`func (o *SessionResponseModelMachine) GetPowerActionPendingOk() (*bool, bool)`

GetPowerActionPendingOk returns a tuple with the PowerActionPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerActionPending

`func (o *SessionResponseModelMachine) SetPowerActionPending(v bool)`

SetPowerActionPending sets PowerActionPending field to given value.

### HasPowerActionPending

`func (o *SessionResponseModelMachine) HasPowerActionPending() bool`

HasPowerActionPending returns a boolean if a field has been set.

### GetPowerState

`func (o *SessionResponseModelMachine) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *SessionResponseModelMachine) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *SessionResponseModelMachine) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.


### GetProvisioningType

`func (o *SessionResponseModelMachine) GetProvisioningType() ProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *SessionResponseModelMachine) GetProvisioningTypeOk() (*ProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *SessionResponseModelMachine) SetProvisioningType(v ProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.


### GetPublishedApplications

`func (o *SessionResponseModelMachine) GetPublishedApplications() []string`

GetPublishedApplications returns the PublishedApplications field if non-nil, zero value otherwise.

### GetPublishedApplicationsOk

`func (o *SessionResponseModelMachine) GetPublishedApplicationsOk() (*[]string, bool)`

GetPublishedApplicationsOk returns a tuple with the PublishedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedApplications

`func (o *SessionResponseModelMachine) SetPublishedApplications(v []string)`

SetPublishedApplications sets PublishedApplications field to given value.

### HasPublishedApplications

`func (o *SessionResponseModelMachine) HasPublishedApplications() bool`

HasPublishedApplications returns a boolean if a field has been set.

### GetPublishedName

`func (o *SessionResponseModelMachine) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *SessionResponseModelMachine) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *SessionResponseModelMachine) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *SessionResponseModelMachine) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### GetRegistrationState

`func (o *SessionResponseModelMachine) GetRegistrationState() RegistrationState`

GetRegistrationState returns the RegistrationState field if non-nil, zero value otherwise.

### GetRegistrationStateOk

`func (o *SessionResponseModelMachine) GetRegistrationStateOk() (*RegistrationState, bool)`

GetRegistrationStateOk returns a tuple with the RegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationState

`func (o *SessionResponseModelMachine) SetRegistrationState(v RegistrationState)`

SetRegistrationState sets RegistrationState field to given value.

### HasRegistrationState

`func (o *SessionResponseModelMachine) HasRegistrationState() bool`

HasRegistrationState returns a boolean if a field has been set.

### GetScheduledReboot

`func (o *SessionResponseModelMachine) GetScheduledReboot() ScheduledReboot`

GetScheduledReboot returns the ScheduledReboot field if non-nil, zero value otherwise.

### GetScheduledRebootOk

`func (o *SessionResponseModelMachine) GetScheduledRebootOk() (*ScheduledReboot, bool)`

GetScheduledRebootOk returns a tuple with the ScheduledReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReboot

`func (o *SessionResponseModelMachine) SetScheduledReboot(v ScheduledReboot)`

SetScheduledReboot sets ScheduledReboot field to given value.

### HasScheduledReboot

`func (o *SessionResponseModelMachine) HasScheduledReboot() bool`

HasScheduledReboot returns a boolean if a field has been set.

### GetSessionClientAddress

`func (o *SessionResponseModelMachine) GetSessionClientAddress() string`

GetSessionClientAddress returns the SessionClientAddress field if non-nil, zero value otherwise.

### GetSessionClientAddressOk

`func (o *SessionResponseModelMachine) GetSessionClientAddressOk() (*string, bool)`

GetSessionClientAddressOk returns a tuple with the SessionClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientAddress

`func (o *SessionResponseModelMachine) SetSessionClientAddress(v string)`

SetSessionClientAddress sets SessionClientAddress field to given value.

### HasSessionClientAddress

`func (o *SessionResponseModelMachine) HasSessionClientAddress() bool`

HasSessionClientAddress returns a boolean if a field has been set.

### GetSessionClientName

`func (o *SessionResponseModelMachine) GetSessionClientName() string`

GetSessionClientName returns the SessionClientName field if non-nil, zero value otherwise.

### GetSessionClientNameOk

`func (o *SessionResponseModelMachine) GetSessionClientNameOk() (*string, bool)`

GetSessionClientNameOk returns a tuple with the SessionClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientName

`func (o *SessionResponseModelMachine) SetSessionClientName(v string)`

SetSessionClientName sets SessionClientName field to given value.

### HasSessionClientName

`func (o *SessionResponseModelMachine) HasSessionClientName() bool`

HasSessionClientName returns a boolean if a field has been set.

### GetSessionClientVersion

`func (o *SessionResponseModelMachine) GetSessionClientVersion() string`

GetSessionClientVersion returns the SessionClientVersion field if non-nil, zero value otherwise.

### GetSessionClientVersionOk

`func (o *SessionResponseModelMachine) GetSessionClientVersionOk() (*string, bool)`

GetSessionClientVersionOk returns a tuple with the SessionClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientVersion

`func (o *SessionResponseModelMachine) SetSessionClientVersion(v string)`

SetSessionClientVersion sets SessionClientVersion field to given value.

### HasSessionClientVersion

`func (o *SessionResponseModelMachine) HasSessionClientVersion() bool`

HasSessionClientVersion returns a boolean if a field has been set.

### GetSessionConnectedViaHostName

`func (o *SessionResponseModelMachine) GetSessionConnectedViaHostName() string`

GetSessionConnectedViaHostName returns the SessionConnectedViaHostName field if non-nil, zero value otherwise.

### GetSessionConnectedViaHostNameOk

`func (o *SessionResponseModelMachine) GetSessionConnectedViaHostNameOk() (*string, bool)`

GetSessionConnectedViaHostNameOk returns a tuple with the SessionConnectedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConnectedViaHostName

`func (o *SessionResponseModelMachine) SetSessionConnectedViaHostName(v string)`

SetSessionConnectedViaHostName sets SessionConnectedViaHostName field to given value.

### HasSessionConnectedViaHostName

`func (o *SessionResponseModelMachine) HasSessionConnectedViaHostName() bool`

HasSessionConnectedViaHostName returns a boolean if a field has been set.

### GetSessionConnectedViaIP

`func (o *SessionResponseModelMachine) GetSessionConnectedViaIP() string`

GetSessionConnectedViaIP returns the SessionConnectedViaIP field if non-nil, zero value otherwise.

### GetSessionConnectedViaIPOk

`func (o *SessionResponseModelMachine) GetSessionConnectedViaIPOk() (*string, bool)`

GetSessionConnectedViaIPOk returns a tuple with the SessionConnectedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConnectedViaIP

`func (o *SessionResponseModelMachine) SetSessionConnectedViaIP(v string)`

SetSessionConnectedViaIP sets SessionConnectedViaIP field to given value.

### HasSessionConnectedViaIP

`func (o *SessionResponseModelMachine) HasSessionConnectedViaIP() bool`

HasSessionConnectedViaIP returns a boolean if a field has been set.

### GetSessionCount

`func (o *SessionResponseModelMachine) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *SessionResponseModelMachine) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *SessionResponseModelMachine) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *SessionResponseModelMachine) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetSessionLaunchedViaHostName

`func (o *SessionResponseModelMachine) GetSessionLaunchedViaHostName() string`

GetSessionLaunchedViaHostName returns the SessionLaunchedViaHostName field if non-nil, zero value otherwise.

### GetSessionLaunchedViaHostNameOk

`func (o *SessionResponseModelMachine) GetSessionLaunchedViaHostNameOk() (*string, bool)`

GetSessionLaunchedViaHostNameOk returns a tuple with the SessionLaunchedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLaunchedViaHostName

`func (o *SessionResponseModelMachine) SetSessionLaunchedViaHostName(v string)`

SetSessionLaunchedViaHostName sets SessionLaunchedViaHostName field to given value.

### HasSessionLaunchedViaHostName

`func (o *SessionResponseModelMachine) HasSessionLaunchedViaHostName() bool`

HasSessionLaunchedViaHostName returns a boolean if a field has been set.

### GetSessionLaunchedViaIP

`func (o *SessionResponseModelMachine) GetSessionLaunchedViaIP() string`

GetSessionLaunchedViaIP returns the SessionLaunchedViaIP field if non-nil, zero value otherwise.

### GetSessionLaunchedViaIPOk

`func (o *SessionResponseModelMachine) GetSessionLaunchedViaIPOk() (*string, bool)`

GetSessionLaunchedViaIPOk returns a tuple with the SessionLaunchedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLaunchedViaIP

`func (o *SessionResponseModelMachine) SetSessionLaunchedViaIP(v string)`

SetSessionLaunchedViaIP sets SessionLaunchedViaIP field to given value.

### HasSessionLaunchedViaIP

`func (o *SessionResponseModelMachine) HasSessionLaunchedViaIP() bool`

HasSessionLaunchedViaIP returns a boolean if a field has been set.

### GetSessionProtocol

`func (o *SessionResponseModelMachine) GetSessionProtocol() ProtocolType`

GetSessionProtocol returns the SessionProtocol field if non-nil, zero value otherwise.

### GetSessionProtocolOk

`func (o *SessionResponseModelMachine) GetSessionProtocolOk() (*ProtocolType, bool)`

GetSessionProtocolOk returns a tuple with the SessionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProtocol

`func (o *SessionResponseModelMachine) SetSessionProtocol(v ProtocolType)`

SetSessionProtocol sets SessionProtocol field to given value.

### HasSessionProtocol

`func (o *SessionResponseModelMachine) HasSessionProtocol() bool`

HasSessionProtocol returns a boolean if a field has been set.

### GetSessionSecureIcaActive

`func (o *SessionResponseModelMachine) GetSessionSecureIcaActive() bool`

GetSessionSecureIcaActive returns the SessionSecureIcaActive field if non-nil, zero value otherwise.

### GetSessionSecureIcaActiveOk

`func (o *SessionResponseModelMachine) GetSessionSecureIcaActiveOk() (*bool, bool)`

GetSessionSecureIcaActiveOk returns a tuple with the SessionSecureIcaActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSecureIcaActive

`func (o *SessionResponseModelMachine) SetSessionSecureIcaActive(v bool)`

SetSessionSecureIcaActive sets SessionSecureIcaActive field to given value.

### HasSessionSecureIcaActive

`func (o *SessionResponseModelMachine) HasSessionSecureIcaActive() bool`

HasSessionSecureIcaActive returns a boolean if a field has been set.

### GetSessionSmartAccessTags

`func (o *SessionResponseModelMachine) GetSessionSmartAccessTags() []string`

GetSessionSmartAccessTags returns the SessionSmartAccessTags field if non-nil, zero value otherwise.

### GetSessionSmartAccessTagsOk

`func (o *SessionResponseModelMachine) GetSessionSmartAccessTagsOk() (*[]string, bool)`

GetSessionSmartAccessTagsOk returns a tuple with the SessionSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSmartAccessTags

`func (o *SessionResponseModelMachine) SetSessionSmartAccessTags(v []string)`

SetSessionSmartAccessTags sets SessionSmartAccessTags field to given value.

### HasSessionSmartAccessTags

`func (o *SessionResponseModelMachine) HasSessionSmartAccessTags() bool`

HasSessionSmartAccessTags returns a boolean if a field has been set.

### GetSessionStartTime

`func (o *SessionResponseModelMachine) GetSessionStartTime() string`

GetSessionStartTime returns the SessionStartTime field if non-nil, zero value otherwise.

### GetSessionStartTimeOk

`func (o *SessionResponseModelMachine) GetSessionStartTimeOk() (*string, bool)`

GetSessionStartTimeOk returns a tuple with the SessionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStartTime

`func (o *SessionResponseModelMachine) SetSessionStartTime(v string)`

SetSessionStartTime sets SessionStartTime field to given value.

### HasSessionStartTime

`func (o *SessionResponseModelMachine) HasSessionStartTime() bool`

HasSessionStartTime returns a boolean if a field has been set.

### GetFormattedSessionStartTime

`func (o *SessionResponseModelMachine) GetFormattedSessionStartTime() string`

GetFormattedSessionStartTime returns the FormattedSessionStartTime field if non-nil, zero value otherwise.

### GetFormattedSessionStartTimeOk

`func (o *SessionResponseModelMachine) GetFormattedSessionStartTimeOk() (*string, bool)`

GetFormattedSessionStartTimeOk returns a tuple with the FormattedSessionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSessionStartTime

`func (o *SessionResponseModelMachine) SetFormattedSessionStartTime(v string)`

SetFormattedSessionStartTime sets FormattedSessionStartTime field to given value.

### HasFormattedSessionStartTime

`func (o *SessionResponseModelMachine) HasFormattedSessionStartTime() bool`

HasFormattedSessionStartTime returns a boolean if a field has been set.

### GetSessionState

`func (o *SessionResponseModelMachine) GetSessionState() SessionState`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *SessionResponseModelMachine) GetSessionStateOk() (*SessionState, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *SessionResponseModelMachine) SetSessionState(v SessionState)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *SessionResponseModelMachine) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### GetSessionStateChangeTime

`func (o *SessionResponseModelMachine) GetSessionStateChangeTime() string`

GetSessionStateChangeTime returns the SessionStateChangeTime field if non-nil, zero value otherwise.

### GetSessionStateChangeTimeOk

`func (o *SessionResponseModelMachine) GetSessionStateChangeTimeOk() (*string, bool)`

GetSessionStateChangeTimeOk returns a tuple with the SessionStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStateChangeTime

`func (o *SessionResponseModelMachine) SetSessionStateChangeTime(v string)`

SetSessionStateChangeTime sets SessionStateChangeTime field to given value.

### HasSessionStateChangeTime

`func (o *SessionResponseModelMachine) HasSessionStateChangeTime() bool`

HasSessionStateChangeTime returns a boolean if a field has been set.

### GetFormattedSessionStateChangeTime

`func (o *SessionResponseModelMachine) GetFormattedSessionStateChangeTime() string`

GetFormattedSessionStateChangeTime returns the FormattedSessionStateChangeTime field if non-nil, zero value otherwise.

### GetFormattedSessionStateChangeTimeOk

`func (o *SessionResponseModelMachine) GetFormattedSessionStateChangeTimeOk() (*string, bool)`

GetFormattedSessionStateChangeTimeOk returns a tuple with the FormattedSessionStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSessionStateChangeTime

`func (o *SessionResponseModelMachine) SetFormattedSessionStateChangeTime(v string)`

SetFormattedSessionStateChangeTime sets FormattedSessionStateChangeTime field to given value.

### HasFormattedSessionStateChangeTime

`func (o *SessionResponseModelMachine) HasFormattedSessionStateChangeTime() bool`

HasFormattedSessionStateChangeTime returns a boolean if a field has been set.

### GetSessionSupport

`func (o *SessionResponseModelMachine) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *SessionResponseModelMachine) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *SessionResponseModelMachine) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSessionUserName

`func (o *SessionResponseModelMachine) GetSessionUserName() string`

GetSessionUserName returns the SessionUserName field if non-nil, zero value otherwise.

### GetSessionUserNameOk

`func (o *SessionResponseModelMachine) GetSessionUserNameOk() (*string, bool)`

GetSessionUserNameOk returns a tuple with the SessionUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUserName

`func (o *SessionResponseModelMachine) SetSessionUserName(v string)`

SetSessionUserName sets SessionUserName field to given value.

### HasSessionUserName

`func (o *SessionResponseModelMachine) HasSessionUserName() bool`

HasSessionUserName returns a boolean if a field has been set.

### GetSid

`func (o *SessionResponseModelMachine) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SessionResponseModelMachine) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SessionResponseModelMachine) SetSid(v string)`

SetSid sets Sid field to given value.


### GetSummaryState

`func (o *SessionResponseModelMachine) GetSummaryState() SummaryState`

GetSummaryState returns the SummaryState field if non-nil, zero value otherwise.

### GetSummaryStateOk

`func (o *SessionResponseModelMachine) GetSummaryStateOk() (*SummaryState, bool)`

GetSummaryStateOk returns a tuple with the SummaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryState

`func (o *SessionResponseModelMachine) SetSummaryState(v SummaryState)`

SetSummaryState sets SummaryState field to given value.


### GetWillShutdownAfterUse

`func (o *SessionResponseModelMachine) GetWillShutdownAfterUse() bool`

GetWillShutdownAfterUse returns the WillShutdownAfterUse field if non-nil, zero value otherwise.

### GetWillShutdownAfterUseOk

`func (o *SessionResponseModelMachine) GetWillShutdownAfterUseOk() (*bool, bool)`

GetWillShutdownAfterUseOk returns a tuple with the WillShutdownAfterUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillShutdownAfterUse

`func (o *SessionResponseModelMachine) SetWillShutdownAfterUse(v bool)`

SetWillShutdownAfterUse sets WillShutdownAfterUse field to given value.

### HasWillShutdownAfterUse

`func (o *SessionResponseModelMachine) HasWillShutdownAfterUse() bool`

HasWillShutdownAfterUse returns a boolean if a field has been set.

### GetWindowsConnectionSetting

`func (o *SessionResponseModelMachine) GetWindowsConnectionSetting() WindowsConnectionSetting`

GetWindowsConnectionSetting returns the WindowsConnectionSetting field if non-nil, zero value otherwise.

### GetWindowsConnectionSettingOk

`func (o *SessionResponseModelMachine) GetWindowsConnectionSettingOk() (*WindowsConnectionSetting, bool)`

GetWindowsConnectionSettingOk returns a tuple with the WindowsConnectionSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsConnectionSetting

`func (o *SessionResponseModelMachine) SetWindowsConnectionSetting(v WindowsConnectionSetting)`

SetWindowsConnectionSetting sets WindowsConnectionSetting field to given value.

### HasWindowsConnectionSetting

`func (o *SessionResponseModelMachine) HasWindowsConnectionSetting() bool`

HasWindowsConnectionSetting returns a boolean if a field has been set.

### GetZone

`func (o *SessionResponseModelMachine) GetZone() MachineResponseModelAllOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SessionResponseModelMachine) GetZoneOk() (*MachineResponseModelAllOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SessionResponseModelMachine) SetZone(v MachineResponseModelAllOfZone)`

SetZone sets Zone field to given value.


### GetSupportedPowerActions

`func (o *SessionResponseModelMachine) GetSupportedPowerActions() []SupportedPowerAction`

GetSupportedPowerActions returns the SupportedPowerActions field if non-nil, zero value otherwise.

### GetSupportedPowerActionsOk

`func (o *SessionResponseModelMachine) GetSupportedPowerActionsOk() (*[]SupportedPowerAction, bool)`

GetSupportedPowerActionsOk returns a tuple with the SupportedPowerActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPowerActions

`func (o *SessionResponseModelMachine) SetSupportedPowerActions(v []SupportedPowerAction)`

SetSupportedPowerActions sets SupportedPowerActions field to given value.

### HasSupportedPowerActions

`func (o *SessionResponseModelMachine) HasSupportedPowerActions() bool`

HasSupportedPowerActions returns a boolean if a field has been set.

### GetFaultState

`func (o *SessionResponseModelMachine) GetFaultState() FaultState`

GetFaultState returns the FaultState field if non-nil, zero value otherwise.

### GetFaultStateOk

`func (o *SessionResponseModelMachine) GetFaultStateOk() (*FaultState, bool)`

GetFaultStateOk returns a tuple with the FaultState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultState

`func (o *SessionResponseModelMachine) SetFaultState(v FaultState)`

SetFaultState sets FaultState field to given value.


### GetContainerMetadata

`func (o *SessionResponseModelMachine) GetContainerMetadata() MachineResponseModelAllOfContainerMetadata`

GetContainerMetadata returns the ContainerMetadata field if non-nil, zero value otherwise.

### GetContainerMetadataOk

`func (o *SessionResponseModelMachine) GetContainerMetadataOk() (*MachineResponseModelAllOfContainerMetadata, bool)`

GetContainerMetadataOk returns a tuple with the ContainerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMetadata

`func (o *SessionResponseModelMachine) SetContainerMetadata(v MachineResponseModelAllOfContainerMetadata)`

SetContainerMetadata sets ContainerMetadata field to given value.

### HasContainerMetadata

`func (o *SessionResponseModelMachine) HasContainerMetadata() bool`

HasContainerMetadata returns a boolean if a field has been set.

### GetTags

`func (o *SessionResponseModelMachine) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SessionResponseModelMachine) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SessionResponseModelMachine) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SessionResponseModelMachine) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpgradeType

`func (o *SessionResponseModelMachine) GetUpgradeType() VdaUpgradeType`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *SessionResponseModelMachine) GetUpgradeTypeOk() (*VdaUpgradeType, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *SessionResponseModelMachine) SetUpgradeType(v VdaUpgradeType)`

SetUpgradeType sets UpgradeType field to given value.

### HasUpgradeType

`func (o *SessionResponseModelMachine) HasUpgradeType() bool`

HasUpgradeType returns a boolean if a field has been set.

### GetUpgradeState

`func (o *SessionResponseModelMachine) GetUpgradeState() VdaUpgradeState`

GetUpgradeState returns the UpgradeState field if non-nil, zero value otherwise.

### GetUpgradeStateOk

`func (o *SessionResponseModelMachine) GetUpgradeStateOk() (*VdaUpgradeState, bool)`

GetUpgradeStateOk returns a tuple with the UpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeState

`func (o *SessionResponseModelMachine) SetUpgradeState(v VdaUpgradeState)`

SetUpgradeState sets UpgradeState field to given value.

### HasUpgradeState

`func (o *SessionResponseModelMachine) HasUpgradeState() bool`

HasUpgradeState returns a boolean if a field has been set.

### GetMachineConfigurationOutOfSync

`func (o *SessionResponseModelMachine) GetMachineConfigurationOutOfSync() bool`

GetMachineConfigurationOutOfSync returns the MachineConfigurationOutOfSync field if non-nil, zero value otherwise.

### GetMachineConfigurationOutOfSyncOk

`func (o *SessionResponseModelMachine) GetMachineConfigurationOutOfSyncOk() (*bool, bool)`

GetMachineConfigurationOutOfSyncOk returns a tuple with the MachineConfigurationOutOfSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineConfigurationOutOfSync

`func (o *SessionResponseModelMachine) SetMachineConfigurationOutOfSync(v bool)`

SetMachineConfigurationOutOfSync sets MachineConfigurationOutOfSync field to given value.

### HasMachineConfigurationOutOfSync

`func (o *SessionResponseModelMachine) HasMachineConfigurationOutOfSync() bool`

HasMachineConfigurationOutOfSync returns a boolean if a field has been set.

### GetUpgradeDetail

`func (o *SessionResponseModelMachine) GetUpgradeDetail() MachineResponseModelAllOfUpgradeDetail`

GetUpgradeDetail returns the UpgradeDetail field if non-nil, zero value otherwise.

### GetUpgradeDetailOk

`func (o *SessionResponseModelMachine) GetUpgradeDetailOk() (*MachineResponseModelAllOfUpgradeDetail, bool)`

GetUpgradeDetailOk returns a tuple with the UpgradeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDetail

`func (o *SessionResponseModelMachine) SetUpgradeDetail(v MachineResponseModelAllOfUpgradeDetail)`

SetUpgradeDetail sets UpgradeDetail field to given value.

### HasUpgradeDetail

`func (o *SessionResponseModelMachine) HasUpgradeDetail() bool`

HasUpgradeDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


