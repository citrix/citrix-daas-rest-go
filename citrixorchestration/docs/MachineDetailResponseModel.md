# MachineDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of machine. Used to be: DesktopUid (and wasn&#39;t globally unique) OR UUID, depending on context Needs to be globally unique Might be constructed from site ID + internal Uid?  or use uuid | 
**MachineCatalog** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**Name** | Pointer to **NullableString** | DNS host name of the machine. Used to be: MachineName | [optional] 
**Uid** | Pointer to **NullableInt32** | DEPRECATED. Use Id. Used to be: DesktopUid | [optional] 
**AgentVersion** | Pointer to **NullableString** | Version of the Citrix Virtual Delivery Agent (VDA) installed on the machine. | [optional] 
**AllocationType** | Pointer to [**AllocationType**](AllocationType.md) |  | [optional] 
**ApplicationsInUse** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | List of applications in use in the session. | [optional] 
**AssignedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | List of one or more users to whom the machine is assigned. Only used when AllocationType is equal to Static. | [optional] 
**AssociatedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The current user(s) for shared machines and the assigned users for private machines. | [optional] 
**AzureAdJoinedMode** | Pointer to [**AzureAdJoinedMode**](AzureAdJoinedMode.md) |  | [optional] 
**ContainerScopes** | Pointer to [**[]ContainerScopeResponseModel**](ContainerScopeResponseModel.md) | Delegated admin scopes in which the containers of the machine reside. | [optional] 
**ControllerDnsName** | Pointer to **NullableString** | The DNS host name of the controller that the machine is registered to. | [optional] 
**DeliveryGroup** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**DeliveryType** | Pointer to [**DeliveryKind**](DeliveryKind.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Description of the machine. | [optional] 
**DesktopConditions** | Pointer to [**[]DesktopCondition**](DesktopCondition.md) | List of outstanding desktop conditions for the machine. | [optional] 
**DnsName** | Pointer to **NullableString** | The DNS host name of the machine. | [optional] 
**FunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**Hosting** | Pointer to [**MachineHostingResponseModel**](MachineHostingResponseModel.md) |  | [optional] 
**InMaintenanceMode** | **bool** | Denotes if the machine is in maintenance mode. Machines in maintenance mode will not accept new sessions. | 
**MaintenanceModeReason** | Pointer to [**MaintenanceModeReason**](MaintenanceModeReason.md) |  | [optional] 
**DrainingUntilShutdown** | **bool** | Denotes if the machine is placed to drain until shutdown | 
**IPAddress** | Pointer to **NullableString** | The IP address of the machine. | [optional] 
**IsAssigned** | Pointer to **NullableBool** | Denotes whether a private desktop has been assigned to a user/users, or a client name/address. Users can be assigned explicitly or by assigning on first use of the machine. Only relevant for privately assigned machines. | [optional] 
**MachineType** | [**MachineType**](MachineType.md) |  | 
**LastConnectionFailure** | Pointer to [**ConnectionFailureReason**](ConnectionFailureReason.md) |  | [optional] 
**LastConnectionTime** | Pointer to **NullableString** | Time of the last detected connection attempt that either failed or succeeded. | [optional] 
**FormattedLastConnectionTime** | Pointer to **NullableString** | Formatted time of the last detected connection attempt that either failed or succeeded. RFC 3339 compatible format. | [optional] 
**LastConnectionUser** | Pointer to [**IdentityUserResponseModel**](IdentityUserResponseModel.md) |  | [optional] 
**LastDeregistrationReason** | Pointer to [**DeregistrationReason**](DeregistrationReason.md) |  | [optional] 
**LastDeregistrationTime** | Pointer to **NullableString** | Time of the last deregistration of the machine from the controller. | [optional] 
**FormattedLastDeregistrationTime** | Pointer to **NullableString** | Formatted time of the last deregistration of the machine from the controller. RFC 3339 compatible format. | [optional] 
**LastErrorReason** | Pointer to **NullableString** | The reason for the last error detected in the machine. | [optional] 
**LastErrorTime** | Pointer to **NullableString** | The time of the last detected error. | [optional] 
**FormattedLastErrorTime** | Pointer to **NullableString** | The formatted time of the last detected error. RFC 3339 compatible format. | [optional] 
**LoadIndex** | Pointer to **NullableInt32** | Gives current effective load index. Only used when SessionSupport is equal to MultiSession. | [optional] 
**MachineUnavailableReason** | Pointer to [**MachineUnavailableReason**](MachineUnavailableReason.md) |  | [optional] 
**OSType** | Pointer to **NullableString** | A string that can be used to identify the operating system that is running on the machine. | [optional] 
**OSVersion** | Pointer to **NullableString** | A string that can be used to identify the version of the operating system running on the machine, if known. | [optional] 
**PersistUserChanges** | Pointer to [**PersistChanges**](PersistChanges.md) |  | [optional] 
**PowerActionPending** | Pointer to **NullableBool** | Indicates if there are any pending power actions for the machine. Only relevant for power-managed machines. | [optional] 
**PowerState** | [**PowerState**](PowerState.md) |  | 
**ProvisioningType** | [**ProvisioningType**](ProvisioningType.md) |  | 
**PublishedApplications** | Pointer to **[]string** | Indicates the published applications. | [optional] 
**PublishedName** | Pointer to **NullableString** | The name of the machine that is displayed in Receiver, if the machine has been published. | [optional] 
**RegistrationState** | Pointer to [**RegistrationState**](RegistrationState.md) |  | [optional] 
**ScheduledReboot** | Pointer to [**ScheduledReboot**](ScheduledReboot.md) |  | [optional] 
**SessionClientAddress** | Pointer to **NullableString** | The IP address of the client connected to the session. | [optional] 
**SessionClientName** | Pointer to **NullableString** | The host name of the client connected to the session. | [optional] 
**SessionClientVersion** | Pointer to **NullableString** | The version of the Citrix Receiver running on the client connected to the session. | [optional] 
**SessionConnectedViaHostName** | Pointer to **NullableString** | The host name of the incoming connection. This is usually a gateway, router or client. | [optional] 
**SessionConnectedViaIP** | Pointer to **NullableString** | The IP address of the incoming connection This is usually a gateway, router or client. | [optional] 
**SessionCount** | Pointer to **NullableInt32** | Number of sessions running on the machine. | [optional] 
**SessionLaunchedViaHostName** | Pointer to **NullableString** | The host name of the StoreFront server used to launch the session. | [optional] 
**SessionLaunchedViaIP** | Pointer to **NullableString** | The IP address of the StoreFront server used to launch the session. | [optional] 
**SessionProtocol** | Pointer to [**ProtocolType**](ProtocolType.md) |  | [optional] 
**SessionSecureIcaActive** | Pointer to **NullableBool** | Indicates whether SecureICA is active on the session. | [optional] 
**SessionSmartAccessTags** | Pointer to **[]string** | The Smart Access tags for this session. | [optional] 
**SessionStartTime** | Pointer to **NullableString** | The time indicates when the session was started. | [optional] 
**FormattedSessionStartTime** | Pointer to **NullableString** | The formatted time indicates when the session was started. RFC 3339 compatible format. | [optional] 
**SessionState** | Pointer to [**SessionState**](SessionState.md) |  | [optional] 
**SessionStateChangeTime** | Pointer to **NullableString** | The time of the most recent state change for the session. | [optional] 
**FormattedSessionStateChangeTime** | Pointer to **NullableString** | The formatted time of the most recent state change for the session. RFC 3339 compatible format. | [optional] 
**SessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 
**SessionUserName** | Pointer to **NullableString** | The session user name. | [optional] 
**Sid** | **string** | The SID of the machine. Used to be: DesktopSid or SID (based on the context) | 
**SummaryState** | [**SummaryState**](SummaryState.md) |  | 
**WillShutdownAfterUse** | Pointer to **NullableBool** | Flag indicating if this machine is tainted and will be shut down after all sessions on the machine have ended. This flag is only ever non-null on power-managed, single-session machines. | [optional] 
**WindowsConnectionSetting** | Pointer to [**WindowsConnectionSetting**](WindowsConnectionSetting.md) |  | [optional] 
**Zone** | [**RefResponseModel**](RefResponseModel.md) |  | 
**SupportedPowerActions** | Pointer to [**[]SupportedPowerAction**](SupportedPowerAction.md) | A list of power actions supported by this machine. | [optional] 
**FaultState** | [**FaultState**](FaultState.md) |  | 
**ContainerMetadata** | Pointer to [**ContainerMetadataModel**](ContainerMetadataModel.md) |  | [optional] 
**Tags** | Pointer to **[]string** | The tags for this machine. | [optional] 
**UpgradeType** | Pointer to [**VdaUpgradeType**](VdaUpgradeType.md) |  | [optional] 
**UpgradeState** | Pointer to [**VdaUpgradeState**](VdaUpgradeState.md) |  | [optional] 
**MachineConfigurationOutOfSync** | Pointer to **NullableBool** | Flag indicating whether the machine&#39;s configuration is out of sync with the catalog&#39;s latest configuration | [optional] 
**UpgradeDetail** | Pointer to [**MachineUpgradeDetail**](MachineUpgradeDetail.md) |  | [optional] 
**AssignedClientName** | Pointer to **NullableString** | The name of the endpoint client device that the machine has been assigned to. | [optional] 
**AssignedIPAddress** | Pointer to **NullableString** | The IP address of the endpoint client device that the machine has been assigned to. | [optional] 
**BrowserName** | Pointer to **NullableString** | Site-wide unique name identifying associated desktop to other components (for example StoreFront). This is typically non-null only for machines backing assigned private desktops. | [optional] 
**ColorDepth** | Pointer to [**ColorDepth**](ColorDepth.md) |  | [optional] 
**IconId** | Pointer to **NullableString** | The machine&#39;s icon that is displayed in Receiver. | [optional] 
**IsReserved** | **bool** | Indicates if machine is reserved for special use, for example for AppDisk preparation. A reserved machine cannot be a member of a delivery group. | 
**LoadIndexes** | Pointer to **[]int32** | Gives the last reported individual load indexes that were used in the calculation of the LoadIndex value. Note that the LoadIndex value may have been subsequently adjusted due to session brokering operations. This value is only set when SessionSupport is equal to MultiSession. | [optional] 
**SecureIcaRequired** | Pointer to **NullableBool** | Flag indicating whether SecureICA is required or not when starting a session on the machine. | [optional] 
**SessionsEstablished** | **int32** | Number of established sessions on this machine.  When SessionSupport is equal to MultiSession, this excludes established sessions which have not yet completed their logon processing. | 
**SessionsPending** | **int32** | Number of pending (brokered but not yet established) sessions on this machine.  When SessionSupport is equal to MultiSession, this also includes established sessions which have not yet completed their logon processing. | 
**StoreFrontServersForHostedReceiver** | Pointer to [**[]StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md) | StoreFront servers to use for Receiver when it is hosted on the machine. | [optional] 
**VMToolsState** | Pointer to [**VMToolsState**](VMToolsState.md) |  | [optional] 
**FailureReason** | Pointer to **NullableString** | Failure reason of power action. | [optional] 

## Methods

### NewMachineDetailResponseModel

`func NewMachineDetailResponseModel(id string, inMaintenanceMode bool, drainingUntilShutdown bool, machineType MachineType, powerState PowerState, provisioningType ProvisioningType, sessionSupport SessionSupport, sid string, summaryState SummaryState, zone RefResponseModel, faultState FaultState, isReserved bool, sessionsEstablished int32, sessionsPending int32, ) *MachineDetailResponseModel`

NewMachineDetailResponseModel instantiates a new MachineDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineDetailResponseModelWithDefaults

`func NewMachineDetailResponseModelWithDefaults() *MachineDetailResponseModel`

NewMachineDetailResponseModelWithDefaults instantiates a new MachineDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetMachineCatalog

`func (o *MachineDetailResponseModel) GetMachineCatalog() RefResponseModel`

GetMachineCatalog returns the MachineCatalog field if non-nil, zero value otherwise.

### GetMachineCatalogOk

`func (o *MachineDetailResponseModel) GetMachineCatalogOk() (*RefResponseModel, bool)`

GetMachineCatalogOk returns a tuple with the MachineCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCatalog

`func (o *MachineDetailResponseModel) SetMachineCatalog(v RefResponseModel)`

SetMachineCatalog sets MachineCatalog field to given value.

### HasMachineCatalog

`func (o *MachineDetailResponseModel) HasMachineCatalog() bool`

HasMachineCatalog returns a boolean if a field has been set.

### GetName

`func (o *MachineDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineDetailResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MachineDetailResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MachineDetailResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUid

`func (o *MachineDetailResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MachineDetailResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MachineDetailResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MachineDetailResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *MachineDetailResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *MachineDetailResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetAgentVersion

`func (o *MachineDetailResponseModel) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *MachineDetailResponseModel) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *MachineDetailResponseModel) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *MachineDetailResponseModel) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *MachineDetailResponseModel) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *MachineDetailResponseModel) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil
### GetAllocationType

`func (o *MachineDetailResponseModel) GetAllocationType() AllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *MachineDetailResponseModel) GetAllocationTypeOk() (*AllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *MachineDetailResponseModel) SetAllocationType(v AllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *MachineDetailResponseModel) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetApplicationsInUse

`func (o *MachineDetailResponseModel) GetApplicationsInUse() []RefResponseModel`

GetApplicationsInUse returns the ApplicationsInUse field if non-nil, zero value otherwise.

### GetApplicationsInUseOk

`func (o *MachineDetailResponseModel) GetApplicationsInUseOk() (*[]RefResponseModel, bool)`

GetApplicationsInUseOk returns a tuple with the ApplicationsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsInUse

`func (o *MachineDetailResponseModel) SetApplicationsInUse(v []RefResponseModel)`

SetApplicationsInUse sets ApplicationsInUse field to given value.

### HasApplicationsInUse

`func (o *MachineDetailResponseModel) HasApplicationsInUse() bool`

HasApplicationsInUse returns a boolean if a field has been set.

### SetApplicationsInUseNil

`func (o *MachineDetailResponseModel) SetApplicationsInUseNil(b bool)`

 SetApplicationsInUseNil sets the value for ApplicationsInUse to be an explicit nil

### UnsetApplicationsInUse
`func (o *MachineDetailResponseModel) UnsetApplicationsInUse()`

UnsetApplicationsInUse ensures that no value is present for ApplicationsInUse, not even an explicit nil
### GetAssignedUsers

`func (o *MachineDetailResponseModel) GetAssignedUsers() []IdentityUserResponseModel`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *MachineDetailResponseModel) GetAssignedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *MachineDetailResponseModel) SetAssignedUsers(v []IdentityUserResponseModel)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *MachineDetailResponseModel) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.

### SetAssignedUsersNil

`func (o *MachineDetailResponseModel) SetAssignedUsersNil(b bool)`

 SetAssignedUsersNil sets the value for AssignedUsers to be an explicit nil

### UnsetAssignedUsers
`func (o *MachineDetailResponseModel) UnsetAssignedUsers()`

UnsetAssignedUsers ensures that no value is present for AssignedUsers, not even an explicit nil
### GetAssociatedUsers

`func (o *MachineDetailResponseModel) GetAssociatedUsers() []IdentityUserResponseModel`

GetAssociatedUsers returns the AssociatedUsers field if non-nil, zero value otherwise.

### GetAssociatedUsersOk

`func (o *MachineDetailResponseModel) GetAssociatedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetAssociatedUsersOk returns a tuple with the AssociatedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedUsers

`func (o *MachineDetailResponseModel) SetAssociatedUsers(v []IdentityUserResponseModel)`

SetAssociatedUsers sets AssociatedUsers field to given value.

### HasAssociatedUsers

`func (o *MachineDetailResponseModel) HasAssociatedUsers() bool`

HasAssociatedUsers returns a boolean if a field has been set.

### SetAssociatedUsersNil

`func (o *MachineDetailResponseModel) SetAssociatedUsersNil(b bool)`

 SetAssociatedUsersNil sets the value for AssociatedUsers to be an explicit nil

### UnsetAssociatedUsers
`func (o *MachineDetailResponseModel) UnsetAssociatedUsers()`

UnsetAssociatedUsers ensures that no value is present for AssociatedUsers, not even an explicit nil
### GetAzureAdJoinedMode

`func (o *MachineDetailResponseModel) GetAzureAdJoinedMode() AzureAdJoinedMode`

GetAzureAdJoinedMode returns the AzureAdJoinedMode field if non-nil, zero value otherwise.

### GetAzureAdJoinedModeOk

`func (o *MachineDetailResponseModel) GetAzureAdJoinedModeOk() (*AzureAdJoinedMode, bool)`

GetAzureAdJoinedModeOk returns a tuple with the AzureAdJoinedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdJoinedMode

`func (o *MachineDetailResponseModel) SetAzureAdJoinedMode(v AzureAdJoinedMode)`

SetAzureAdJoinedMode sets AzureAdJoinedMode field to given value.

### HasAzureAdJoinedMode

`func (o *MachineDetailResponseModel) HasAzureAdJoinedMode() bool`

HasAzureAdJoinedMode returns a boolean if a field has been set.

### GetContainerScopes

`func (o *MachineDetailResponseModel) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *MachineDetailResponseModel) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *MachineDetailResponseModel) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.

### HasContainerScopes

`func (o *MachineDetailResponseModel) HasContainerScopes() bool`

HasContainerScopes returns a boolean if a field has been set.

### SetContainerScopesNil

`func (o *MachineDetailResponseModel) SetContainerScopesNil(b bool)`

 SetContainerScopesNil sets the value for ContainerScopes to be an explicit nil

### UnsetContainerScopes
`func (o *MachineDetailResponseModel) UnsetContainerScopes()`

UnsetContainerScopes ensures that no value is present for ContainerScopes, not even an explicit nil
### GetControllerDnsName

`func (o *MachineDetailResponseModel) GetControllerDnsName() string`

GetControllerDnsName returns the ControllerDnsName field if non-nil, zero value otherwise.

### GetControllerDnsNameOk

`func (o *MachineDetailResponseModel) GetControllerDnsNameOk() (*string, bool)`

GetControllerDnsNameOk returns a tuple with the ControllerDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDnsName

`func (o *MachineDetailResponseModel) SetControllerDnsName(v string)`

SetControllerDnsName sets ControllerDnsName field to given value.

### HasControllerDnsName

`func (o *MachineDetailResponseModel) HasControllerDnsName() bool`

HasControllerDnsName returns a boolean if a field has been set.

### SetControllerDnsNameNil

`func (o *MachineDetailResponseModel) SetControllerDnsNameNil(b bool)`

 SetControllerDnsNameNil sets the value for ControllerDnsName to be an explicit nil

### UnsetControllerDnsName
`func (o *MachineDetailResponseModel) UnsetControllerDnsName()`

UnsetControllerDnsName ensures that no value is present for ControllerDnsName, not even an explicit nil
### GetDeliveryGroup

`func (o *MachineDetailResponseModel) GetDeliveryGroup() RefResponseModel`

GetDeliveryGroup returns the DeliveryGroup field if non-nil, zero value otherwise.

### GetDeliveryGroupOk

`func (o *MachineDetailResponseModel) GetDeliveryGroupOk() (*RefResponseModel, bool)`

GetDeliveryGroupOk returns a tuple with the DeliveryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroup

`func (o *MachineDetailResponseModel) SetDeliveryGroup(v RefResponseModel)`

SetDeliveryGroup sets DeliveryGroup field to given value.

### HasDeliveryGroup

`func (o *MachineDetailResponseModel) HasDeliveryGroup() bool`

HasDeliveryGroup returns a boolean if a field has been set.

### GetDeliveryType

`func (o *MachineDetailResponseModel) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *MachineDetailResponseModel) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *MachineDetailResponseModel) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *MachineDetailResponseModel) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDescription

`func (o *MachineDetailResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MachineDetailResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MachineDetailResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MachineDetailResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MachineDetailResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MachineDetailResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesktopConditions

`func (o *MachineDetailResponseModel) GetDesktopConditions() []DesktopCondition`

GetDesktopConditions returns the DesktopConditions field if non-nil, zero value otherwise.

### GetDesktopConditionsOk

`func (o *MachineDetailResponseModel) GetDesktopConditionsOk() (*[]DesktopCondition, bool)`

GetDesktopConditionsOk returns a tuple with the DesktopConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopConditions

`func (o *MachineDetailResponseModel) SetDesktopConditions(v []DesktopCondition)`

SetDesktopConditions sets DesktopConditions field to given value.

### HasDesktopConditions

`func (o *MachineDetailResponseModel) HasDesktopConditions() bool`

HasDesktopConditions returns a boolean if a field has been set.

### SetDesktopConditionsNil

`func (o *MachineDetailResponseModel) SetDesktopConditionsNil(b bool)`

 SetDesktopConditionsNil sets the value for DesktopConditions to be an explicit nil

### UnsetDesktopConditions
`func (o *MachineDetailResponseModel) UnsetDesktopConditions()`

UnsetDesktopConditions ensures that no value is present for DesktopConditions, not even an explicit nil
### GetDnsName

`func (o *MachineDetailResponseModel) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *MachineDetailResponseModel) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *MachineDetailResponseModel) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *MachineDetailResponseModel) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### SetDnsNameNil

`func (o *MachineDetailResponseModel) SetDnsNameNil(b bool)`

 SetDnsNameNil sets the value for DnsName to be an explicit nil

### UnsetDnsName
`func (o *MachineDetailResponseModel) UnsetDnsName()`

UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
### GetFunctionalLevel

`func (o *MachineDetailResponseModel) GetFunctionalLevel() FunctionalLevel`

GetFunctionalLevel returns the FunctionalLevel field if non-nil, zero value otherwise.

### GetFunctionalLevelOk

`func (o *MachineDetailResponseModel) GetFunctionalLevelOk() (*FunctionalLevel, bool)`

GetFunctionalLevelOk returns a tuple with the FunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalLevel

`func (o *MachineDetailResponseModel) SetFunctionalLevel(v FunctionalLevel)`

SetFunctionalLevel sets FunctionalLevel field to given value.

### HasFunctionalLevel

`func (o *MachineDetailResponseModel) HasFunctionalLevel() bool`

HasFunctionalLevel returns a boolean if a field has been set.

### GetHosting

`func (o *MachineDetailResponseModel) GetHosting() MachineHostingResponseModel`

GetHosting returns the Hosting field if non-nil, zero value otherwise.

### GetHostingOk

`func (o *MachineDetailResponseModel) GetHostingOk() (*MachineHostingResponseModel, bool)`

GetHostingOk returns a tuple with the Hosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosting

`func (o *MachineDetailResponseModel) SetHosting(v MachineHostingResponseModel)`

SetHosting sets Hosting field to given value.

### HasHosting

`func (o *MachineDetailResponseModel) HasHosting() bool`

HasHosting returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *MachineDetailResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *MachineDetailResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *MachineDetailResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.


### GetMaintenanceModeReason

`func (o *MachineDetailResponseModel) GetMaintenanceModeReason() MaintenanceModeReason`

GetMaintenanceModeReason returns the MaintenanceModeReason field if non-nil, zero value otherwise.

### GetMaintenanceModeReasonOk

`func (o *MachineDetailResponseModel) GetMaintenanceModeReasonOk() (*MaintenanceModeReason, bool)`

GetMaintenanceModeReasonOk returns a tuple with the MaintenanceModeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceModeReason

`func (o *MachineDetailResponseModel) SetMaintenanceModeReason(v MaintenanceModeReason)`

SetMaintenanceModeReason sets MaintenanceModeReason field to given value.

### HasMaintenanceModeReason

`func (o *MachineDetailResponseModel) HasMaintenanceModeReason() bool`

HasMaintenanceModeReason returns a boolean if a field has been set.

### GetDrainingUntilShutdown

`func (o *MachineDetailResponseModel) GetDrainingUntilShutdown() bool`

GetDrainingUntilShutdown returns the DrainingUntilShutdown field if non-nil, zero value otherwise.

### GetDrainingUntilShutdownOk

`func (o *MachineDetailResponseModel) GetDrainingUntilShutdownOk() (*bool, bool)`

GetDrainingUntilShutdownOk returns a tuple with the DrainingUntilShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainingUntilShutdown

`func (o *MachineDetailResponseModel) SetDrainingUntilShutdown(v bool)`

SetDrainingUntilShutdown sets DrainingUntilShutdown field to given value.


### GetIPAddress

`func (o *MachineDetailResponseModel) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *MachineDetailResponseModel) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *MachineDetailResponseModel) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *MachineDetailResponseModel) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### SetIPAddressNil

`func (o *MachineDetailResponseModel) SetIPAddressNil(b bool)`

 SetIPAddressNil sets the value for IPAddress to be an explicit nil

### UnsetIPAddress
`func (o *MachineDetailResponseModel) UnsetIPAddress()`

UnsetIPAddress ensures that no value is present for IPAddress, not even an explicit nil
### GetIsAssigned

`func (o *MachineDetailResponseModel) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MachineDetailResponseModel) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *MachineDetailResponseModel) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *MachineDetailResponseModel) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### SetIsAssignedNil

`func (o *MachineDetailResponseModel) SetIsAssignedNil(b bool)`

 SetIsAssignedNil sets the value for IsAssigned to be an explicit nil

### UnsetIsAssigned
`func (o *MachineDetailResponseModel) UnsetIsAssigned()`

UnsetIsAssigned ensures that no value is present for IsAssigned, not even an explicit nil
### GetMachineType

`func (o *MachineDetailResponseModel) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *MachineDetailResponseModel) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *MachineDetailResponseModel) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.


### GetLastConnectionFailure

`func (o *MachineDetailResponseModel) GetLastConnectionFailure() ConnectionFailureReason`

GetLastConnectionFailure returns the LastConnectionFailure field if non-nil, zero value otherwise.

### GetLastConnectionFailureOk

`func (o *MachineDetailResponseModel) GetLastConnectionFailureOk() (*ConnectionFailureReason, bool)`

GetLastConnectionFailureOk returns a tuple with the LastConnectionFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionFailure

`func (o *MachineDetailResponseModel) SetLastConnectionFailure(v ConnectionFailureReason)`

SetLastConnectionFailure sets LastConnectionFailure field to given value.

### HasLastConnectionFailure

`func (o *MachineDetailResponseModel) HasLastConnectionFailure() bool`

HasLastConnectionFailure returns a boolean if a field has been set.

### GetLastConnectionTime

`func (o *MachineDetailResponseModel) GetLastConnectionTime() string`

GetLastConnectionTime returns the LastConnectionTime field if non-nil, zero value otherwise.

### GetLastConnectionTimeOk

`func (o *MachineDetailResponseModel) GetLastConnectionTimeOk() (*string, bool)`

GetLastConnectionTimeOk returns a tuple with the LastConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionTime

`func (o *MachineDetailResponseModel) SetLastConnectionTime(v string)`

SetLastConnectionTime sets LastConnectionTime field to given value.

### HasLastConnectionTime

`func (o *MachineDetailResponseModel) HasLastConnectionTime() bool`

HasLastConnectionTime returns a boolean if a field has been set.

### SetLastConnectionTimeNil

`func (o *MachineDetailResponseModel) SetLastConnectionTimeNil(b bool)`

 SetLastConnectionTimeNil sets the value for LastConnectionTime to be an explicit nil

### UnsetLastConnectionTime
`func (o *MachineDetailResponseModel) UnsetLastConnectionTime()`

UnsetLastConnectionTime ensures that no value is present for LastConnectionTime, not even an explicit nil
### GetFormattedLastConnectionTime

`func (o *MachineDetailResponseModel) GetFormattedLastConnectionTime() string`

GetFormattedLastConnectionTime returns the FormattedLastConnectionTime field if non-nil, zero value otherwise.

### GetFormattedLastConnectionTimeOk

`func (o *MachineDetailResponseModel) GetFormattedLastConnectionTimeOk() (*string, bool)`

GetFormattedLastConnectionTimeOk returns a tuple with the FormattedLastConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastConnectionTime

`func (o *MachineDetailResponseModel) SetFormattedLastConnectionTime(v string)`

SetFormattedLastConnectionTime sets FormattedLastConnectionTime field to given value.

### HasFormattedLastConnectionTime

`func (o *MachineDetailResponseModel) HasFormattedLastConnectionTime() bool`

HasFormattedLastConnectionTime returns a boolean if a field has been set.

### SetFormattedLastConnectionTimeNil

`func (o *MachineDetailResponseModel) SetFormattedLastConnectionTimeNil(b bool)`

 SetFormattedLastConnectionTimeNil sets the value for FormattedLastConnectionTime to be an explicit nil

### UnsetFormattedLastConnectionTime
`func (o *MachineDetailResponseModel) UnsetFormattedLastConnectionTime()`

UnsetFormattedLastConnectionTime ensures that no value is present for FormattedLastConnectionTime, not even an explicit nil
### GetLastConnectionUser

`func (o *MachineDetailResponseModel) GetLastConnectionUser() IdentityUserResponseModel`

GetLastConnectionUser returns the LastConnectionUser field if non-nil, zero value otherwise.

### GetLastConnectionUserOk

`func (o *MachineDetailResponseModel) GetLastConnectionUserOk() (*IdentityUserResponseModel, bool)`

GetLastConnectionUserOk returns a tuple with the LastConnectionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionUser

`func (o *MachineDetailResponseModel) SetLastConnectionUser(v IdentityUserResponseModel)`

SetLastConnectionUser sets LastConnectionUser field to given value.

### HasLastConnectionUser

`func (o *MachineDetailResponseModel) HasLastConnectionUser() bool`

HasLastConnectionUser returns a boolean if a field has been set.

### GetLastDeregistrationReason

`func (o *MachineDetailResponseModel) GetLastDeregistrationReason() DeregistrationReason`

GetLastDeregistrationReason returns the LastDeregistrationReason field if non-nil, zero value otherwise.

### GetLastDeregistrationReasonOk

`func (o *MachineDetailResponseModel) GetLastDeregistrationReasonOk() (*DeregistrationReason, bool)`

GetLastDeregistrationReasonOk returns a tuple with the LastDeregistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationReason

`func (o *MachineDetailResponseModel) SetLastDeregistrationReason(v DeregistrationReason)`

SetLastDeregistrationReason sets LastDeregistrationReason field to given value.

### HasLastDeregistrationReason

`func (o *MachineDetailResponseModel) HasLastDeregistrationReason() bool`

HasLastDeregistrationReason returns a boolean if a field has been set.

### GetLastDeregistrationTime

`func (o *MachineDetailResponseModel) GetLastDeregistrationTime() string`

GetLastDeregistrationTime returns the LastDeregistrationTime field if non-nil, zero value otherwise.

### GetLastDeregistrationTimeOk

`func (o *MachineDetailResponseModel) GetLastDeregistrationTimeOk() (*string, bool)`

GetLastDeregistrationTimeOk returns a tuple with the LastDeregistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationTime

`func (o *MachineDetailResponseModel) SetLastDeregistrationTime(v string)`

SetLastDeregistrationTime sets LastDeregistrationTime field to given value.

### HasLastDeregistrationTime

`func (o *MachineDetailResponseModel) HasLastDeregistrationTime() bool`

HasLastDeregistrationTime returns a boolean if a field has been set.

### SetLastDeregistrationTimeNil

`func (o *MachineDetailResponseModel) SetLastDeregistrationTimeNil(b bool)`

 SetLastDeregistrationTimeNil sets the value for LastDeregistrationTime to be an explicit nil

### UnsetLastDeregistrationTime
`func (o *MachineDetailResponseModel) UnsetLastDeregistrationTime()`

UnsetLastDeregistrationTime ensures that no value is present for LastDeregistrationTime, not even an explicit nil
### GetFormattedLastDeregistrationTime

`func (o *MachineDetailResponseModel) GetFormattedLastDeregistrationTime() string`

GetFormattedLastDeregistrationTime returns the FormattedLastDeregistrationTime field if non-nil, zero value otherwise.

### GetFormattedLastDeregistrationTimeOk

`func (o *MachineDetailResponseModel) GetFormattedLastDeregistrationTimeOk() (*string, bool)`

GetFormattedLastDeregistrationTimeOk returns a tuple with the FormattedLastDeregistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastDeregistrationTime

`func (o *MachineDetailResponseModel) SetFormattedLastDeregistrationTime(v string)`

SetFormattedLastDeregistrationTime sets FormattedLastDeregistrationTime field to given value.

### HasFormattedLastDeregistrationTime

`func (o *MachineDetailResponseModel) HasFormattedLastDeregistrationTime() bool`

HasFormattedLastDeregistrationTime returns a boolean if a field has been set.

### SetFormattedLastDeregistrationTimeNil

`func (o *MachineDetailResponseModel) SetFormattedLastDeregistrationTimeNil(b bool)`

 SetFormattedLastDeregistrationTimeNil sets the value for FormattedLastDeregistrationTime to be an explicit nil

### UnsetFormattedLastDeregistrationTime
`func (o *MachineDetailResponseModel) UnsetFormattedLastDeregistrationTime()`

UnsetFormattedLastDeregistrationTime ensures that no value is present for FormattedLastDeregistrationTime, not even an explicit nil
### GetLastErrorReason

`func (o *MachineDetailResponseModel) GetLastErrorReason() string`

GetLastErrorReason returns the LastErrorReason field if non-nil, zero value otherwise.

### GetLastErrorReasonOk

`func (o *MachineDetailResponseModel) GetLastErrorReasonOk() (*string, bool)`

GetLastErrorReasonOk returns a tuple with the LastErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorReason

`func (o *MachineDetailResponseModel) SetLastErrorReason(v string)`

SetLastErrorReason sets LastErrorReason field to given value.

### HasLastErrorReason

`func (o *MachineDetailResponseModel) HasLastErrorReason() bool`

HasLastErrorReason returns a boolean if a field has been set.

### SetLastErrorReasonNil

`func (o *MachineDetailResponseModel) SetLastErrorReasonNil(b bool)`

 SetLastErrorReasonNil sets the value for LastErrorReason to be an explicit nil

### UnsetLastErrorReason
`func (o *MachineDetailResponseModel) UnsetLastErrorReason()`

UnsetLastErrorReason ensures that no value is present for LastErrorReason, not even an explicit nil
### GetLastErrorTime

`func (o *MachineDetailResponseModel) GetLastErrorTime() string`

GetLastErrorTime returns the LastErrorTime field if non-nil, zero value otherwise.

### GetLastErrorTimeOk

`func (o *MachineDetailResponseModel) GetLastErrorTimeOk() (*string, bool)`

GetLastErrorTimeOk returns a tuple with the LastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorTime

`func (o *MachineDetailResponseModel) SetLastErrorTime(v string)`

SetLastErrorTime sets LastErrorTime field to given value.

### HasLastErrorTime

`func (o *MachineDetailResponseModel) HasLastErrorTime() bool`

HasLastErrorTime returns a boolean if a field has been set.

### SetLastErrorTimeNil

`func (o *MachineDetailResponseModel) SetLastErrorTimeNil(b bool)`

 SetLastErrorTimeNil sets the value for LastErrorTime to be an explicit nil

### UnsetLastErrorTime
`func (o *MachineDetailResponseModel) UnsetLastErrorTime()`

UnsetLastErrorTime ensures that no value is present for LastErrorTime, not even an explicit nil
### GetFormattedLastErrorTime

`func (o *MachineDetailResponseModel) GetFormattedLastErrorTime() string`

GetFormattedLastErrorTime returns the FormattedLastErrorTime field if non-nil, zero value otherwise.

### GetFormattedLastErrorTimeOk

`func (o *MachineDetailResponseModel) GetFormattedLastErrorTimeOk() (*string, bool)`

GetFormattedLastErrorTimeOk returns a tuple with the FormattedLastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastErrorTime

`func (o *MachineDetailResponseModel) SetFormattedLastErrorTime(v string)`

SetFormattedLastErrorTime sets FormattedLastErrorTime field to given value.

### HasFormattedLastErrorTime

`func (o *MachineDetailResponseModel) HasFormattedLastErrorTime() bool`

HasFormattedLastErrorTime returns a boolean if a field has been set.

### SetFormattedLastErrorTimeNil

`func (o *MachineDetailResponseModel) SetFormattedLastErrorTimeNil(b bool)`

 SetFormattedLastErrorTimeNil sets the value for FormattedLastErrorTime to be an explicit nil

### UnsetFormattedLastErrorTime
`func (o *MachineDetailResponseModel) UnsetFormattedLastErrorTime()`

UnsetFormattedLastErrorTime ensures that no value is present for FormattedLastErrorTime, not even an explicit nil
### GetLoadIndex

`func (o *MachineDetailResponseModel) GetLoadIndex() int32`

GetLoadIndex returns the LoadIndex field if non-nil, zero value otherwise.

### GetLoadIndexOk

`func (o *MachineDetailResponseModel) GetLoadIndexOk() (*int32, bool)`

GetLoadIndexOk returns a tuple with the LoadIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadIndex

`func (o *MachineDetailResponseModel) SetLoadIndex(v int32)`

SetLoadIndex sets LoadIndex field to given value.

### HasLoadIndex

`func (o *MachineDetailResponseModel) HasLoadIndex() bool`

HasLoadIndex returns a boolean if a field has been set.

### SetLoadIndexNil

`func (o *MachineDetailResponseModel) SetLoadIndexNil(b bool)`

 SetLoadIndexNil sets the value for LoadIndex to be an explicit nil

### UnsetLoadIndex
`func (o *MachineDetailResponseModel) UnsetLoadIndex()`

UnsetLoadIndex ensures that no value is present for LoadIndex, not even an explicit nil
### GetMachineUnavailableReason

`func (o *MachineDetailResponseModel) GetMachineUnavailableReason() MachineUnavailableReason`

GetMachineUnavailableReason returns the MachineUnavailableReason field if non-nil, zero value otherwise.

### GetMachineUnavailableReasonOk

`func (o *MachineDetailResponseModel) GetMachineUnavailableReasonOk() (*MachineUnavailableReason, bool)`

GetMachineUnavailableReasonOk returns a tuple with the MachineUnavailableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineUnavailableReason

`func (o *MachineDetailResponseModel) SetMachineUnavailableReason(v MachineUnavailableReason)`

SetMachineUnavailableReason sets MachineUnavailableReason field to given value.

### HasMachineUnavailableReason

`func (o *MachineDetailResponseModel) HasMachineUnavailableReason() bool`

HasMachineUnavailableReason returns a boolean if a field has been set.

### GetOSType

`func (o *MachineDetailResponseModel) GetOSType() string`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *MachineDetailResponseModel) GetOSTypeOk() (*string, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *MachineDetailResponseModel) SetOSType(v string)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *MachineDetailResponseModel) HasOSType() bool`

HasOSType returns a boolean if a field has been set.

### SetOSTypeNil

`func (o *MachineDetailResponseModel) SetOSTypeNil(b bool)`

 SetOSTypeNil sets the value for OSType to be an explicit nil

### UnsetOSType
`func (o *MachineDetailResponseModel) UnsetOSType()`

UnsetOSType ensures that no value is present for OSType, not even an explicit nil
### GetOSVersion

`func (o *MachineDetailResponseModel) GetOSVersion() string`

GetOSVersion returns the OSVersion field if non-nil, zero value otherwise.

### GetOSVersionOk

`func (o *MachineDetailResponseModel) GetOSVersionOk() (*string, bool)`

GetOSVersionOk returns a tuple with the OSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSVersion

`func (o *MachineDetailResponseModel) SetOSVersion(v string)`

SetOSVersion sets OSVersion field to given value.

### HasOSVersion

`func (o *MachineDetailResponseModel) HasOSVersion() bool`

HasOSVersion returns a boolean if a field has been set.

### SetOSVersionNil

`func (o *MachineDetailResponseModel) SetOSVersionNil(b bool)`

 SetOSVersionNil sets the value for OSVersion to be an explicit nil

### UnsetOSVersion
`func (o *MachineDetailResponseModel) UnsetOSVersion()`

UnsetOSVersion ensures that no value is present for OSVersion, not even an explicit nil
### GetPersistUserChanges

`func (o *MachineDetailResponseModel) GetPersistUserChanges() PersistChanges`

GetPersistUserChanges returns the PersistUserChanges field if non-nil, zero value otherwise.

### GetPersistUserChangesOk

`func (o *MachineDetailResponseModel) GetPersistUserChangesOk() (*PersistChanges, bool)`

GetPersistUserChangesOk returns a tuple with the PersistUserChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistUserChanges

`func (o *MachineDetailResponseModel) SetPersistUserChanges(v PersistChanges)`

SetPersistUserChanges sets PersistUserChanges field to given value.

### HasPersistUserChanges

`func (o *MachineDetailResponseModel) HasPersistUserChanges() bool`

HasPersistUserChanges returns a boolean if a field has been set.

### GetPowerActionPending

`func (o *MachineDetailResponseModel) GetPowerActionPending() bool`

GetPowerActionPending returns the PowerActionPending field if non-nil, zero value otherwise.

### GetPowerActionPendingOk

`func (o *MachineDetailResponseModel) GetPowerActionPendingOk() (*bool, bool)`

GetPowerActionPendingOk returns a tuple with the PowerActionPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerActionPending

`func (o *MachineDetailResponseModel) SetPowerActionPending(v bool)`

SetPowerActionPending sets PowerActionPending field to given value.

### HasPowerActionPending

`func (o *MachineDetailResponseModel) HasPowerActionPending() bool`

HasPowerActionPending returns a boolean if a field has been set.

### SetPowerActionPendingNil

`func (o *MachineDetailResponseModel) SetPowerActionPendingNil(b bool)`

 SetPowerActionPendingNil sets the value for PowerActionPending to be an explicit nil

### UnsetPowerActionPending
`func (o *MachineDetailResponseModel) UnsetPowerActionPending()`

UnsetPowerActionPending ensures that no value is present for PowerActionPending, not even an explicit nil
### GetPowerState

`func (o *MachineDetailResponseModel) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *MachineDetailResponseModel) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *MachineDetailResponseModel) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.


### GetProvisioningType

`func (o *MachineDetailResponseModel) GetProvisioningType() ProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *MachineDetailResponseModel) GetProvisioningTypeOk() (*ProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *MachineDetailResponseModel) SetProvisioningType(v ProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.


### GetPublishedApplications

`func (o *MachineDetailResponseModel) GetPublishedApplications() []string`

GetPublishedApplications returns the PublishedApplications field if non-nil, zero value otherwise.

### GetPublishedApplicationsOk

`func (o *MachineDetailResponseModel) GetPublishedApplicationsOk() (*[]string, bool)`

GetPublishedApplicationsOk returns a tuple with the PublishedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedApplications

`func (o *MachineDetailResponseModel) SetPublishedApplications(v []string)`

SetPublishedApplications sets PublishedApplications field to given value.

### HasPublishedApplications

`func (o *MachineDetailResponseModel) HasPublishedApplications() bool`

HasPublishedApplications returns a boolean if a field has been set.

### SetPublishedApplicationsNil

`func (o *MachineDetailResponseModel) SetPublishedApplicationsNil(b bool)`

 SetPublishedApplicationsNil sets the value for PublishedApplications to be an explicit nil

### UnsetPublishedApplications
`func (o *MachineDetailResponseModel) UnsetPublishedApplications()`

UnsetPublishedApplications ensures that no value is present for PublishedApplications, not even an explicit nil
### GetPublishedName

`func (o *MachineDetailResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *MachineDetailResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *MachineDetailResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *MachineDetailResponseModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *MachineDetailResponseModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *MachineDetailResponseModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetRegistrationState

`func (o *MachineDetailResponseModel) GetRegistrationState() RegistrationState`

GetRegistrationState returns the RegistrationState field if non-nil, zero value otherwise.

### GetRegistrationStateOk

`func (o *MachineDetailResponseModel) GetRegistrationStateOk() (*RegistrationState, bool)`

GetRegistrationStateOk returns a tuple with the RegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationState

`func (o *MachineDetailResponseModel) SetRegistrationState(v RegistrationState)`

SetRegistrationState sets RegistrationState field to given value.

### HasRegistrationState

`func (o *MachineDetailResponseModel) HasRegistrationState() bool`

HasRegistrationState returns a boolean if a field has been set.

### GetScheduledReboot

`func (o *MachineDetailResponseModel) GetScheduledReboot() ScheduledReboot`

GetScheduledReboot returns the ScheduledReboot field if non-nil, zero value otherwise.

### GetScheduledRebootOk

`func (o *MachineDetailResponseModel) GetScheduledRebootOk() (*ScheduledReboot, bool)`

GetScheduledRebootOk returns a tuple with the ScheduledReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReboot

`func (o *MachineDetailResponseModel) SetScheduledReboot(v ScheduledReboot)`

SetScheduledReboot sets ScheduledReboot field to given value.

### HasScheduledReboot

`func (o *MachineDetailResponseModel) HasScheduledReboot() bool`

HasScheduledReboot returns a boolean if a field has been set.

### GetSessionClientAddress

`func (o *MachineDetailResponseModel) GetSessionClientAddress() string`

GetSessionClientAddress returns the SessionClientAddress field if non-nil, zero value otherwise.

### GetSessionClientAddressOk

`func (o *MachineDetailResponseModel) GetSessionClientAddressOk() (*string, bool)`

GetSessionClientAddressOk returns a tuple with the SessionClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientAddress

`func (o *MachineDetailResponseModel) SetSessionClientAddress(v string)`

SetSessionClientAddress sets SessionClientAddress field to given value.

### HasSessionClientAddress

`func (o *MachineDetailResponseModel) HasSessionClientAddress() bool`

HasSessionClientAddress returns a boolean if a field has been set.

### SetSessionClientAddressNil

`func (o *MachineDetailResponseModel) SetSessionClientAddressNil(b bool)`

 SetSessionClientAddressNil sets the value for SessionClientAddress to be an explicit nil

### UnsetSessionClientAddress
`func (o *MachineDetailResponseModel) UnsetSessionClientAddress()`

UnsetSessionClientAddress ensures that no value is present for SessionClientAddress, not even an explicit nil
### GetSessionClientName

`func (o *MachineDetailResponseModel) GetSessionClientName() string`

GetSessionClientName returns the SessionClientName field if non-nil, zero value otherwise.

### GetSessionClientNameOk

`func (o *MachineDetailResponseModel) GetSessionClientNameOk() (*string, bool)`

GetSessionClientNameOk returns a tuple with the SessionClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientName

`func (o *MachineDetailResponseModel) SetSessionClientName(v string)`

SetSessionClientName sets SessionClientName field to given value.

### HasSessionClientName

`func (o *MachineDetailResponseModel) HasSessionClientName() bool`

HasSessionClientName returns a boolean if a field has been set.

### SetSessionClientNameNil

`func (o *MachineDetailResponseModel) SetSessionClientNameNil(b bool)`

 SetSessionClientNameNil sets the value for SessionClientName to be an explicit nil

### UnsetSessionClientName
`func (o *MachineDetailResponseModel) UnsetSessionClientName()`

UnsetSessionClientName ensures that no value is present for SessionClientName, not even an explicit nil
### GetSessionClientVersion

`func (o *MachineDetailResponseModel) GetSessionClientVersion() string`

GetSessionClientVersion returns the SessionClientVersion field if non-nil, zero value otherwise.

### GetSessionClientVersionOk

`func (o *MachineDetailResponseModel) GetSessionClientVersionOk() (*string, bool)`

GetSessionClientVersionOk returns a tuple with the SessionClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientVersion

`func (o *MachineDetailResponseModel) SetSessionClientVersion(v string)`

SetSessionClientVersion sets SessionClientVersion field to given value.

### HasSessionClientVersion

`func (o *MachineDetailResponseModel) HasSessionClientVersion() bool`

HasSessionClientVersion returns a boolean if a field has been set.

### SetSessionClientVersionNil

`func (o *MachineDetailResponseModel) SetSessionClientVersionNil(b bool)`

 SetSessionClientVersionNil sets the value for SessionClientVersion to be an explicit nil

### UnsetSessionClientVersion
`func (o *MachineDetailResponseModel) UnsetSessionClientVersion()`

UnsetSessionClientVersion ensures that no value is present for SessionClientVersion, not even an explicit nil
### GetSessionConnectedViaHostName

`func (o *MachineDetailResponseModel) GetSessionConnectedViaHostName() string`

GetSessionConnectedViaHostName returns the SessionConnectedViaHostName field if non-nil, zero value otherwise.

### GetSessionConnectedViaHostNameOk

`func (o *MachineDetailResponseModel) GetSessionConnectedViaHostNameOk() (*string, bool)`

GetSessionConnectedViaHostNameOk returns a tuple with the SessionConnectedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConnectedViaHostName

`func (o *MachineDetailResponseModel) SetSessionConnectedViaHostName(v string)`

SetSessionConnectedViaHostName sets SessionConnectedViaHostName field to given value.

### HasSessionConnectedViaHostName

`func (o *MachineDetailResponseModel) HasSessionConnectedViaHostName() bool`

HasSessionConnectedViaHostName returns a boolean if a field has been set.

### SetSessionConnectedViaHostNameNil

`func (o *MachineDetailResponseModel) SetSessionConnectedViaHostNameNil(b bool)`

 SetSessionConnectedViaHostNameNil sets the value for SessionConnectedViaHostName to be an explicit nil

### UnsetSessionConnectedViaHostName
`func (o *MachineDetailResponseModel) UnsetSessionConnectedViaHostName()`

UnsetSessionConnectedViaHostName ensures that no value is present for SessionConnectedViaHostName, not even an explicit nil
### GetSessionConnectedViaIP

`func (o *MachineDetailResponseModel) GetSessionConnectedViaIP() string`

GetSessionConnectedViaIP returns the SessionConnectedViaIP field if non-nil, zero value otherwise.

### GetSessionConnectedViaIPOk

`func (o *MachineDetailResponseModel) GetSessionConnectedViaIPOk() (*string, bool)`

GetSessionConnectedViaIPOk returns a tuple with the SessionConnectedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConnectedViaIP

`func (o *MachineDetailResponseModel) SetSessionConnectedViaIP(v string)`

SetSessionConnectedViaIP sets SessionConnectedViaIP field to given value.

### HasSessionConnectedViaIP

`func (o *MachineDetailResponseModel) HasSessionConnectedViaIP() bool`

HasSessionConnectedViaIP returns a boolean if a field has been set.

### SetSessionConnectedViaIPNil

`func (o *MachineDetailResponseModel) SetSessionConnectedViaIPNil(b bool)`

 SetSessionConnectedViaIPNil sets the value for SessionConnectedViaIP to be an explicit nil

### UnsetSessionConnectedViaIP
`func (o *MachineDetailResponseModel) UnsetSessionConnectedViaIP()`

UnsetSessionConnectedViaIP ensures that no value is present for SessionConnectedViaIP, not even an explicit nil
### GetSessionCount

`func (o *MachineDetailResponseModel) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *MachineDetailResponseModel) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *MachineDetailResponseModel) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *MachineDetailResponseModel) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### SetSessionCountNil

`func (o *MachineDetailResponseModel) SetSessionCountNil(b bool)`

 SetSessionCountNil sets the value for SessionCount to be an explicit nil

### UnsetSessionCount
`func (o *MachineDetailResponseModel) UnsetSessionCount()`

UnsetSessionCount ensures that no value is present for SessionCount, not even an explicit nil
### GetSessionLaunchedViaHostName

`func (o *MachineDetailResponseModel) GetSessionLaunchedViaHostName() string`

GetSessionLaunchedViaHostName returns the SessionLaunchedViaHostName field if non-nil, zero value otherwise.

### GetSessionLaunchedViaHostNameOk

`func (o *MachineDetailResponseModel) GetSessionLaunchedViaHostNameOk() (*string, bool)`

GetSessionLaunchedViaHostNameOk returns a tuple with the SessionLaunchedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLaunchedViaHostName

`func (o *MachineDetailResponseModel) SetSessionLaunchedViaHostName(v string)`

SetSessionLaunchedViaHostName sets SessionLaunchedViaHostName field to given value.

### HasSessionLaunchedViaHostName

`func (o *MachineDetailResponseModel) HasSessionLaunchedViaHostName() bool`

HasSessionLaunchedViaHostName returns a boolean if a field has been set.

### SetSessionLaunchedViaHostNameNil

`func (o *MachineDetailResponseModel) SetSessionLaunchedViaHostNameNil(b bool)`

 SetSessionLaunchedViaHostNameNil sets the value for SessionLaunchedViaHostName to be an explicit nil

### UnsetSessionLaunchedViaHostName
`func (o *MachineDetailResponseModel) UnsetSessionLaunchedViaHostName()`

UnsetSessionLaunchedViaHostName ensures that no value is present for SessionLaunchedViaHostName, not even an explicit nil
### GetSessionLaunchedViaIP

`func (o *MachineDetailResponseModel) GetSessionLaunchedViaIP() string`

GetSessionLaunchedViaIP returns the SessionLaunchedViaIP field if non-nil, zero value otherwise.

### GetSessionLaunchedViaIPOk

`func (o *MachineDetailResponseModel) GetSessionLaunchedViaIPOk() (*string, bool)`

GetSessionLaunchedViaIPOk returns a tuple with the SessionLaunchedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLaunchedViaIP

`func (o *MachineDetailResponseModel) SetSessionLaunchedViaIP(v string)`

SetSessionLaunchedViaIP sets SessionLaunchedViaIP field to given value.

### HasSessionLaunchedViaIP

`func (o *MachineDetailResponseModel) HasSessionLaunchedViaIP() bool`

HasSessionLaunchedViaIP returns a boolean if a field has been set.

### SetSessionLaunchedViaIPNil

`func (o *MachineDetailResponseModel) SetSessionLaunchedViaIPNil(b bool)`

 SetSessionLaunchedViaIPNil sets the value for SessionLaunchedViaIP to be an explicit nil

### UnsetSessionLaunchedViaIP
`func (o *MachineDetailResponseModel) UnsetSessionLaunchedViaIP()`

UnsetSessionLaunchedViaIP ensures that no value is present for SessionLaunchedViaIP, not even an explicit nil
### GetSessionProtocol

`func (o *MachineDetailResponseModel) GetSessionProtocol() ProtocolType`

GetSessionProtocol returns the SessionProtocol field if non-nil, zero value otherwise.

### GetSessionProtocolOk

`func (o *MachineDetailResponseModel) GetSessionProtocolOk() (*ProtocolType, bool)`

GetSessionProtocolOk returns a tuple with the SessionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProtocol

`func (o *MachineDetailResponseModel) SetSessionProtocol(v ProtocolType)`

SetSessionProtocol sets SessionProtocol field to given value.

### HasSessionProtocol

`func (o *MachineDetailResponseModel) HasSessionProtocol() bool`

HasSessionProtocol returns a boolean if a field has been set.

### GetSessionSecureIcaActive

`func (o *MachineDetailResponseModel) GetSessionSecureIcaActive() bool`

GetSessionSecureIcaActive returns the SessionSecureIcaActive field if non-nil, zero value otherwise.

### GetSessionSecureIcaActiveOk

`func (o *MachineDetailResponseModel) GetSessionSecureIcaActiveOk() (*bool, bool)`

GetSessionSecureIcaActiveOk returns a tuple with the SessionSecureIcaActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSecureIcaActive

`func (o *MachineDetailResponseModel) SetSessionSecureIcaActive(v bool)`

SetSessionSecureIcaActive sets SessionSecureIcaActive field to given value.

### HasSessionSecureIcaActive

`func (o *MachineDetailResponseModel) HasSessionSecureIcaActive() bool`

HasSessionSecureIcaActive returns a boolean if a field has been set.

### SetSessionSecureIcaActiveNil

`func (o *MachineDetailResponseModel) SetSessionSecureIcaActiveNil(b bool)`

 SetSessionSecureIcaActiveNil sets the value for SessionSecureIcaActive to be an explicit nil

### UnsetSessionSecureIcaActive
`func (o *MachineDetailResponseModel) UnsetSessionSecureIcaActive()`

UnsetSessionSecureIcaActive ensures that no value is present for SessionSecureIcaActive, not even an explicit nil
### GetSessionSmartAccessTags

`func (o *MachineDetailResponseModel) GetSessionSmartAccessTags() []string`

GetSessionSmartAccessTags returns the SessionSmartAccessTags field if non-nil, zero value otherwise.

### GetSessionSmartAccessTagsOk

`func (o *MachineDetailResponseModel) GetSessionSmartAccessTagsOk() (*[]string, bool)`

GetSessionSmartAccessTagsOk returns a tuple with the SessionSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSmartAccessTags

`func (o *MachineDetailResponseModel) SetSessionSmartAccessTags(v []string)`

SetSessionSmartAccessTags sets SessionSmartAccessTags field to given value.

### HasSessionSmartAccessTags

`func (o *MachineDetailResponseModel) HasSessionSmartAccessTags() bool`

HasSessionSmartAccessTags returns a boolean if a field has been set.

### SetSessionSmartAccessTagsNil

`func (o *MachineDetailResponseModel) SetSessionSmartAccessTagsNil(b bool)`

 SetSessionSmartAccessTagsNil sets the value for SessionSmartAccessTags to be an explicit nil

### UnsetSessionSmartAccessTags
`func (o *MachineDetailResponseModel) UnsetSessionSmartAccessTags()`

UnsetSessionSmartAccessTags ensures that no value is present for SessionSmartAccessTags, not even an explicit nil
### GetSessionStartTime

`func (o *MachineDetailResponseModel) GetSessionStartTime() string`

GetSessionStartTime returns the SessionStartTime field if non-nil, zero value otherwise.

### GetSessionStartTimeOk

`func (o *MachineDetailResponseModel) GetSessionStartTimeOk() (*string, bool)`

GetSessionStartTimeOk returns a tuple with the SessionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStartTime

`func (o *MachineDetailResponseModel) SetSessionStartTime(v string)`

SetSessionStartTime sets SessionStartTime field to given value.

### HasSessionStartTime

`func (o *MachineDetailResponseModel) HasSessionStartTime() bool`

HasSessionStartTime returns a boolean if a field has been set.

### SetSessionStartTimeNil

`func (o *MachineDetailResponseModel) SetSessionStartTimeNil(b bool)`

 SetSessionStartTimeNil sets the value for SessionStartTime to be an explicit nil

### UnsetSessionStartTime
`func (o *MachineDetailResponseModel) UnsetSessionStartTime()`

UnsetSessionStartTime ensures that no value is present for SessionStartTime, not even an explicit nil
### GetFormattedSessionStartTime

`func (o *MachineDetailResponseModel) GetFormattedSessionStartTime() string`

GetFormattedSessionStartTime returns the FormattedSessionStartTime field if non-nil, zero value otherwise.

### GetFormattedSessionStartTimeOk

`func (o *MachineDetailResponseModel) GetFormattedSessionStartTimeOk() (*string, bool)`

GetFormattedSessionStartTimeOk returns a tuple with the FormattedSessionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSessionStartTime

`func (o *MachineDetailResponseModel) SetFormattedSessionStartTime(v string)`

SetFormattedSessionStartTime sets FormattedSessionStartTime field to given value.

### HasFormattedSessionStartTime

`func (o *MachineDetailResponseModel) HasFormattedSessionStartTime() bool`

HasFormattedSessionStartTime returns a boolean if a field has been set.

### SetFormattedSessionStartTimeNil

`func (o *MachineDetailResponseModel) SetFormattedSessionStartTimeNil(b bool)`

 SetFormattedSessionStartTimeNil sets the value for FormattedSessionStartTime to be an explicit nil

### UnsetFormattedSessionStartTime
`func (o *MachineDetailResponseModel) UnsetFormattedSessionStartTime()`

UnsetFormattedSessionStartTime ensures that no value is present for FormattedSessionStartTime, not even an explicit nil
### GetSessionState

`func (o *MachineDetailResponseModel) GetSessionState() SessionState`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *MachineDetailResponseModel) GetSessionStateOk() (*SessionState, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *MachineDetailResponseModel) SetSessionState(v SessionState)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *MachineDetailResponseModel) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### GetSessionStateChangeTime

`func (o *MachineDetailResponseModel) GetSessionStateChangeTime() string`

GetSessionStateChangeTime returns the SessionStateChangeTime field if non-nil, zero value otherwise.

### GetSessionStateChangeTimeOk

`func (o *MachineDetailResponseModel) GetSessionStateChangeTimeOk() (*string, bool)`

GetSessionStateChangeTimeOk returns a tuple with the SessionStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStateChangeTime

`func (o *MachineDetailResponseModel) SetSessionStateChangeTime(v string)`

SetSessionStateChangeTime sets SessionStateChangeTime field to given value.

### HasSessionStateChangeTime

`func (o *MachineDetailResponseModel) HasSessionStateChangeTime() bool`

HasSessionStateChangeTime returns a boolean if a field has been set.

### SetSessionStateChangeTimeNil

`func (o *MachineDetailResponseModel) SetSessionStateChangeTimeNil(b bool)`

 SetSessionStateChangeTimeNil sets the value for SessionStateChangeTime to be an explicit nil

### UnsetSessionStateChangeTime
`func (o *MachineDetailResponseModel) UnsetSessionStateChangeTime()`

UnsetSessionStateChangeTime ensures that no value is present for SessionStateChangeTime, not even an explicit nil
### GetFormattedSessionStateChangeTime

`func (o *MachineDetailResponseModel) GetFormattedSessionStateChangeTime() string`

GetFormattedSessionStateChangeTime returns the FormattedSessionStateChangeTime field if non-nil, zero value otherwise.

### GetFormattedSessionStateChangeTimeOk

`func (o *MachineDetailResponseModel) GetFormattedSessionStateChangeTimeOk() (*string, bool)`

GetFormattedSessionStateChangeTimeOk returns a tuple with the FormattedSessionStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSessionStateChangeTime

`func (o *MachineDetailResponseModel) SetFormattedSessionStateChangeTime(v string)`

SetFormattedSessionStateChangeTime sets FormattedSessionStateChangeTime field to given value.

### HasFormattedSessionStateChangeTime

`func (o *MachineDetailResponseModel) HasFormattedSessionStateChangeTime() bool`

HasFormattedSessionStateChangeTime returns a boolean if a field has been set.

### SetFormattedSessionStateChangeTimeNil

`func (o *MachineDetailResponseModel) SetFormattedSessionStateChangeTimeNil(b bool)`

 SetFormattedSessionStateChangeTimeNil sets the value for FormattedSessionStateChangeTime to be an explicit nil

### UnsetFormattedSessionStateChangeTime
`func (o *MachineDetailResponseModel) UnsetFormattedSessionStateChangeTime()`

UnsetFormattedSessionStateChangeTime ensures that no value is present for FormattedSessionStateChangeTime, not even an explicit nil
### GetSessionSupport

`func (o *MachineDetailResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *MachineDetailResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *MachineDetailResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetSessionUserName

`func (o *MachineDetailResponseModel) GetSessionUserName() string`

GetSessionUserName returns the SessionUserName field if non-nil, zero value otherwise.

### GetSessionUserNameOk

`func (o *MachineDetailResponseModel) GetSessionUserNameOk() (*string, bool)`

GetSessionUserNameOk returns a tuple with the SessionUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUserName

`func (o *MachineDetailResponseModel) SetSessionUserName(v string)`

SetSessionUserName sets SessionUserName field to given value.

### HasSessionUserName

`func (o *MachineDetailResponseModel) HasSessionUserName() bool`

HasSessionUserName returns a boolean if a field has been set.

### SetSessionUserNameNil

`func (o *MachineDetailResponseModel) SetSessionUserNameNil(b bool)`

 SetSessionUserNameNil sets the value for SessionUserName to be an explicit nil

### UnsetSessionUserName
`func (o *MachineDetailResponseModel) UnsetSessionUserName()`

UnsetSessionUserName ensures that no value is present for SessionUserName, not even an explicit nil
### GetSid

`func (o *MachineDetailResponseModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MachineDetailResponseModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MachineDetailResponseModel) SetSid(v string)`

SetSid sets Sid field to given value.


### GetSummaryState

`func (o *MachineDetailResponseModel) GetSummaryState() SummaryState`

GetSummaryState returns the SummaryState field if non-nil, zero value otherwise.

### GetSummaryStateOk

`func (o *MachineDetailResponseModel) GetSummaryStateOk() (*SummaryState, bool)`

GetSummaryStateOk returns a tuple with the SummaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryState

`func (o *MachineDetailResponseModel) SetSummaryState(v SummaryState)`

SetSummaryState sets SummaryState field to given value.


### GetWillShutdownAfterUse

`func (o *MachineDetailResponseModel) GetWillShutdownAfterUse() bool`

GetWillShutdownAfterUse returns the WillShutdownAfterUse field if non-nil, zero value otherwise.

### GetWillShutdownAfterUseOk

`func (o *MachineDetailResponseModel) GetWillShutdownAfterUseOk() (*bool, bool)`

GetWillShutdownAfterUseOk returns a tuple with the WillShutdownAfterUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillShutdownAfterUse

`func (o *MachineDetailResponseModel) SetWillShutdownAfterUse(v bool)`

SetWillShutdownAfterUse sets WillShutdownAfterUse field to given value.

### HasWillShutdownAfterUse

`func (o *MachineDetailResponseModel) HasWillShutdownAfterUse() bool`

HasWillShutdownAfterUse returns a boolean if a field has been set.

### SetWillShutdownAfterUseNil

`func (o *MachineDetailResponseModel) SetWillShutdownAfterUseNil(b bool)`

 SetWillShutdownAfterUseNil sets the value for WillShutdownAfterUse to be an explicit nil

### UnsetWillShutdownAfterUse
`func (o *MachineDetailResponseModel) UnsetWillShutdownAfterUse()`

UnsetWillShutdownAfterUse ensures that no value is present for WillShutdownAfterUse, not even an explicit nil
### GetWindowsConnectionSetting

`func (o *MachineDetailResponseModel) GetWindowsConnectionSetting() WindowsConnectionSetting`

GetWindowsConnectionSetting returns the WindowsConnectionSetting field if non-nil, zero value otherwise.

### GetWindowsConnectionSettingOk

`func (o *MachineDetailResponseModel) GetWindowsConnectionSettingOk() (*WindowsConnectionSetting, bool)`

GetWindowsConnectionSettingOk returns a tuple with the WindowsConnectionSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsConnectionSetting

`func (o *MachineDetailResponseModel) SetWindowsConnectionSetting(v WindowsConnectionSetting)`

SetWindowsConnectionSetting sets WindowsConnectionSetting field to given value.

### HasWindowsConnectionSetting

`func (o *MachineDetailResponseModel) HasWindowsConnectionSetting() bool`

HasWindowsConnectionSetting returns a boolean if a field has been set.

### GetZone

`func (o *MachineDetailResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *MachineDetailResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *MachineDetailResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.


### GetSupportedPowerActions

`func (o *MachineDetailResponseModel) GetSupportedPowerActions() []SupportedPowerAction`

GetSupportedPowerActions returns the SupportedPowerActions field if non-nil, zero value otherwise.

### GetSupportedPowerActionsOk

`func (o *MachineDetailResponseModel) GetSupportedPowerActionsOk() (*[]SupportedPowerAction, bool)`

GetSupportedPowerActionsOk returns a tuple with the SupportedPowerActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPowerActions

`func (o *MachineDetailResponseModel) SetSupportedPowerActions(v []SupportedPowerAction)`

SetSupportedPowerActions sets SupportedPowerActions field to given value.

### HasSupportedPowerActions

`func (o *MachineDetailResponseModel) HasSupportedPowerActions() bool`

HasSupportedPowerActions returns a boolean if a field has been set.

### SetSupportedPowerActionsNil

`func (o *MachineDetailResponseModel) SetSupportedPowerActionsNil(b bool)`

 SetSupportedPowerActionsNil sets the value for SupportedPowerActions to be an explicit nil

### UnsetSupportedPowerActions
`func (o *MachineDetailResponseModel) UnsetSupportedPowerActions()`

UnsetSupportedPowerActions ensures that no value is present for SupportedPowerActions, not even an explicit nil
### GetFaultState

`func (o *MachineDetailResponseModel) GetFaultState() FaultState`

GetFaultState returns the FaultState field if non-nil, zero value otherwise.

### GetFaultStateOk

`func (o *MachineDetailResponseModel) GetFaultStateOk() (*FaultState, bool)`

GetFaultStateOk returns a tuple with the FaultState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultState

`func (o *MachineDetailResponseModel) SetFaultState(v FaultState)`

SetFaultState sets FaultState field to given value.


### GetContainerMetadata

`func (o *MachineDetailResponseModel) GetContainerMetadata() ContainerMetadataModel`

GetContainerMetadata returns the ContainerMetadata field if non-nil, zero value otherwise.

### GetContainerMetadataOk

`func (o *MachineDetailResponseModel) GetContainerMetadataOk() (*ContainerMetadataModel, bool)`

GetContainerMetadataOk returns a tuple with the ContainerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMetadata

`func (o *MachineDetailResponseModel) SetContainerMetadata(v ContainerMetadataModel)`

SetContainerMetadata sets ContainerMetadata field to given value.

### HasContainerMetadata

`func (o *MachineDetailResponseModel) HasContainerMetadata() bool`

HasContainerMetadata returns a boolean if a field has been set.

### GetTags

`func (o *MachineDetailResponseModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MachineDetailResponseModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MachineDetailResponseModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MachineDetailResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MachineDetailResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MachineDetailResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUpgradeType

`func (o *MachineDetailResponseModel) GetUpgradeType() VdaUpgradeType`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *MachineDetailResponseModel) GetUpgradeTypeOk() (*VdaUpgradeType, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *MachineDetailResponseModel) SetUpgradeType(v VdaUpgradeType)`

SetUpgradeType sets UpgradeType field to given value.

### HasUpgradeType

`func (o *MachineDetailResponseModel) HasUpgradeType() bool`

HasUpgradeType returns a boolean if a field has been set.

### GetUpgradeState

`func (o *MachineDetailResponseModel) GetUpgradeState() VdaUpgradeState`

GetUpgradeState returns the UpgradeState field if non-nil, zero value otherwise.

### GetUpgradeStateOk

`func (o *MachineDetailResponseModel) GetUpgradeStateOk() (*VdaUpgradeState, bool)`

GetUpgradeStateOk returns a tuple with the UpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeState

`func (o *MachineDetailResponseModel) SetUpgradeState(v VdaUpgradeState)`

SetUpgradeState sets UpgradeState field to given value.

### HasUpgradeState

`func (o *MachineDetailResponseModel) HasUpgradeState() bool`

HasUpgradeState returns a boolean if a field has been set.

### GetMachineConfigurationOutOfSync

`func (o *MachineDetailResponseModel) GetMachineConfigurationOutOfSync() bool`

GetMachineConfigurationOutOfSync returns the MachineConfigurationOutOfSync field if non-nil, zero value otherwise.

### GetMachineConfigurationOutOfSyncOk

`func (o *MachineDetailResponseModel) GetMachineConfigurationOutOfSyncOk() (*bool, bool)`

GetMachineConfigurationOutOfSyncOk returns a tuple with the MachineConfigurationOutOfSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineConfigurationOutOfSync

`func (o *MachineDetailResponseModel) SetMachineConfigurationOutOfSync(v bool)`

SetMachineConfigurationOutOfSync sets MachineConfigurationOutOfSync field to given value.

### HasMachineConfigurationOutOfSync

`func (o *MachineDetailResponseModel) HasMachineConfigurationOutOfSync() bool`

HasMachineConfigurationOutOfSync returns a boolean if a field has been set.

### SetMachineConfigurationOutOfSyncNil

`func (o *MachineDetailResponseModel) SetMachineConfigurationOutOfSyncNil(b bool)`

 SetMachineConfigurationOutOfSyncNil sets the value for MachineConfigurationOutOfSync to be an explicit nil

### UnsetMachineConfigurationOutOfSync
`func (o *MachineDetailResponseModel) UnsetMachineConfigurationOutOfSync()`

UnsetMachineConfigurationOutOfSync ensures that no value is present for MachineConfigurationOutOfSync, not even an explicit nil
### GetUpgradeDetail

`func (o *MachineDetailResponseModel) GetUpgradeDetail() MachineUpgradeDetail`

GetUpgradeDetail returns the UpgradeDetail field if non-nil, zero value otherwise.

### GetUpgradeDetailOk

`func (o *MachineDetailResponseModel) GetUpgradeDetailOk() (*MachineUpgradeDetail, bool)`

GetUpgradeDetailOk returns a tuple with the UpgradeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDetail

`func (o *MachineDetailResponseModel) SetUpgradeDetail(v MachineUpgradeDetail)`

SetUpgradeDetail sets UpgradeDetail field to given value.

### HasUpgradeDetail

`func (o *MachineDetailResponseModel) HasUpgradeDetail() bool`

HasUpgradeDetail returns a boolean if a field has been set.

### GetAssignedClientName

`func (o *MachineDetailResponseModel) GetAssignedClientName() string`

GetAssignedClientName returns the AssignedClientName field if non-nil, zero value otherwise.

### GetAssignedClientNameOk

`func (o *MachineDetailResponseModel) GetAssignedClientNameOk() (*string, bool)`

GetAssignedClientNameOk returns a tuple with the AssignedClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedClientName

`func (o *MachineDetailResponseModel) SetAssignedClientName(v string)`

SetAssignedClientName sets AssignedClientName field to given value.

### HasAssignedClientName

`func (o *MachineDetailResponseModel) HasAssignedClientName() bool`

HasAssignedClientName returns a boolean if a field has been set.

### SetAssignedClientNameNil

`func (o *MachineDetailResponseModel) SetAssignedClientNameNil(b bool)`

 SetAssignedClientNameNil sets the value for AssignedClientName to be an explicit nil

### UnsetAssignedClientName
`func (o *MachineDetailResponseModel) UnsetAssignedClientName()`

UnsetAssignedClientName ensures that no value is present for AssignedClientName, not even an explicit nil
### GetAssignedIPAddress

`func (o *MachineDetailResponseModel) GetAssignedIPAddress() string`

GetAssignedIPAddress returns the AssignedIPAddress field if non-nil, zero value otherwise.

### GetAssignedIPAddressOk

`func (o *MachineDetailResponseModel) GetAssignedIPAddressOk() (*string, bool)`

GetAssignedIPAddressOk returns a tuple with the AssignedIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedIPAddress

`func (o *MachineDetailResponseModel) SetAssignedIPAddress(v string)`

SetAssignedIPAddress sets AssignedIPAddress field to given value.

### HasAssignedIPAddress

`func (o *MachineDetailResponseModel) HasAssignedIPAddress() bool`

HasAssignedIPAddress returns a boolean if a field has been set.

### SetAssignedIPAddressNil

`func (o *MachineDetailResponseModel) SetAssignedIPAddressNil(b bool)`

 SetAssignedIPAddressNil sets the value for AssignedIPAddress to be an explicit nil

### UnsetAssignedIPAddress
`func (o *MachineDetailResponseModel) UnsetAssignedIPAddress()`

UnsetAssignedIPAddress ensures that no value is present for AssignedIPAddress, not even an explicit nil
### GetBrowserName

`func (o *MachineDetailResponseModel) GetBrowserName() string`

GetBrowserName returns the BrowserName field if non-nil, zero value otherwise.

### GetBrowserNameOk

`func (o *MachineDetailResponseModel) GetBrowserNameOk() (*string, bool)`

GetBrowserNameOk returns a tuple with the BrowserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserName

`func (o *MachineDetailResponseModel) SetBrowserName(v string)`

SetBrowserName sets BrowserName field to given value.

### HasBrowserName

`func (o *MachineDetailResponseModel) HasBrowserName() bool`

HasBrowserName returns a boolean if a field has been set.

### SetBrowserNameNil

`func (o *MachineDetailResponseModel) SetBrowserNameNil(b bool)`

 SetBrowserNameNil sets the value for BrowserName to be an explicit nil

### UnsetBrowserName
`func (o *MachineDetailResponseModel) UnsetBrowserName()`

UnsetBrowserName ensures that no value is present for BrowserName, not even an explicit nil
### GetColorDepth

`func (o *MachineDetailResponseModel) GetColorDepth() ColorDepth`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *MachineDetailResponseModel) GetColorDepthOk() (*ColorDepth, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *MachineDetailResponseModel) SetColorDepth(v ColorDepth)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *MachineDetailResponseModel) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetIconId

`func (o *MachineDetailResponseModel) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *MachineDetailResponseModel) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *MachineDetailResponseModel) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *MachineDetailResponseModel) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### SetIconIdNil

`func (o *MachineDetailResponseModel) SetIconIdNil(b bool)`

 SetIconIdNil sets the value for IconId to be an explicit nil

### UnsetIconId
`func (o *MachineDetailResponseModel) UnsetIconId()`

UnsetIconId ensures that no value is present for IconId, not even an explicit nil
### GetIsReserved

`func (o *MachineDetailResponseModel) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *MachineDetailResponseModel) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *MachineDetailResponseModel) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.


### GetLoadIndexes

`func (o *MachineDetailResponseModel) GetLoadIndexes() []int32`

GetLoadIndexes returns the LoadIndexes field if non-nil, zero value otherwise.

### GetLoadIndexesOk

`func (o *MachineDetailResponseModel) GetLoadIndexesOk() (*[]int32, bool)`

GetLoadIndexesOk returns a tuple with the LoadIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadIndexes

`func (o *MachineDetailResponseModel) SetLoadIndexes(v []int32)`

SetLoadIndexes sets LoadIndexes field to given value.

### HasLoadIndexes

`func (o *MachineDetailResponseModel) HasLoadIndexes() bool`

HasLoadIndexes returns a boolean if a field has been set.

### SetLoadIndexesNil

`func (o *MachineDetailResponseModel) SetLoadIndexesNil(b bool)`

 SetLoadIndexesNil sets the value for LoadIndexes to be an explicit nil

### UnsetLoadIndexes
`func (o *MachineDetailResponseModel) UnsetLoadIndexes()`

UnsetLoadIndexes ensures that no value is present for LoadIndexes, not even an explicit nil
### GetSecureIcaRequired

`func (o *MachineDetailResponseModel) GetSecureIcaRequired() bool`

GetSecureIcaRequired returns the SecureIcaRequired field if non-nil, zero value otherwise.

### GetSecureIcaRequiredOk

`func (o *MachineDetailResponseModel) GetSecureIcaRequiredOk() (*bool, bool)`

GetSecureIcaRequiredOk returns a tuple with the SecureIcaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureIcaRequired

`func (o *MachineDetailResponseModel) SetSecureIcaRequired(v bool)`

SetSecureIcaRequired sets SecureIcaRequired field to given value.

### HasSecureIcaRequired

`func (o *MachineDetailResponseModel) HasSecureIcaRequired() bool`

HasSecureIcaRequired returns a boolean if a field has been set.

### SetSecureIcaRequiredNil

`func (o *MachineDetailResponseModel) SetSecureIcaRequiredNil(b bool)`

 SetSecureIcaRequiredNil sets the value for SecureIcaRequired to be an explicit nil

### UnsetSecureIcaRequired
`func (o *MachineDetailResponseModel) UnsetSecureIcaRequired()`

UnsetSecureIcaRequired ensures that no value is present for SecureIcaRequired, not even an explicit nil
### GetSessionsEstablished

`func (o *MachineDetailResponseModel) GetSessionsEstablished() int32`

GetSessionsEstablished returns the SessionsEstablished field if non-nil, zero value otherwise.

### GetSessionsEstablishedOk

`func (o *MachineDetailResponseModel) GetSessionsEstablishedOk() (*int32, bool)`

GetSessionsEstablishedOk returns a tuple with the SessionsEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsEstablished

`func (o *MachineDetailResponseModel) SetSessionsEstablished(v int32)`

SetSessionsEstablished sets SessionsEstablished field to given value.


### GetSessionsPending

`func (o *MachineDetailResponseModel) GetSessionsPending() int32`

GetSessionsPending returns the SessionsPending field if non-nil, zero value otherwise.

### GetSessionsPendingOk

`func (o *MachineDetailResponseModel) GetSessionsPendingOk() (*int32, bool)`

GetSessionsPendingOk returns a tuple with the SessionsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsPending

`func (o *MachineDetailResponseModel) SetSessionsPending(v int32)`

SetSessionsPending sets SessionsPending field to given value.


### GetStoreFrontServersForHostedReceiver

`func (o *MachineDetailResponseModel) GetStoreFrontServersForHostedReceiver() []StoreFrontServerResponseModel`

GetStoreFrontServersForHostedReceiver returns the StoreFrontServersForHostedReceiver field if non-nil, zero value otherwise.

### GetStoreFrontServersForHostedReceiverOk

`func (o *MachineDetailResponseModel) GetStoreFrontServersForHostedReceiverOk() (*[]StoreFrontServerResponseModel, bool)`

GetStoreFrontServersForHostedReceiverOk returns a tuple with the StoreFrontServersForHostedReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFrontServersForHostedReceiver

`func (o *MachineDetailResponseModel) SetStoreFrontServersForHostedReceiver(v []StoreFrontServerResponseModel)`

SetStoreFrontServersForHostedReceiver sets StoreFrontServersForHostedReceiver field to given value.

### HasStoreFrontServersForHostedReceiver

`func (o *MachineDetailResponseModel) HasStoreFrontServersForHostedReceiver() bool`

HasStoreFrontServersForHostedReceiver returns a boolean if a field has been set.

### SetStoreFrontServersForHostedReceiverNil

`func (o *MachineDetailResponseModel) SetStoreFrontServersForHostedReceiverNil(b bool)`

 SetStoreFrontServersForHostedReceiverNil sets the value for StoreFrontServersForHostedReceiver to be an explicit nil

### UnsetStoreFrontServersForHostedReceiver
`func (o *MachineDetailResponseModel) UnsetStoreFrontServersForHostedReceiver()`

UnsetStoreFrontServersForHostedReceiver ensures that no value is present for StoreFrontServersForHostedReceiver, not even an explicit nil
### GetVMToolsState

`func (o *MachineDetailResponseModel) GetVMToolsState() VMToolsState`

GetVMToolsState returns the VMToolsState field if non-nil, zero value otherwise.

### GetVMToolsStateOk

`func (o *MachineDetailResponseModel) GetVMToolsStateOk() (*VMToolsState, bool)`

GetVMToolsStateOk returns a tuple with the VMToolsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMToolsState

`func (o *MachineDetailResponseModel) SetVMToolsState(v VMToolsState)`

SetVMToolsState sets VMToolsState field to given value.

### HasVMToolsState

`func (o *MachineDetailResponseModel) HasVMToolsState() bool`

HasVMToolsState returns a boolean if a field has been set.

### GetFailureReason

`func (o *MachineDetailResponseModel) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MachineDetailResponseModel) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MachineDetailResponseModel) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MachineDetailResponseModel) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *MachineDetailResponseModel) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *MachineDetailResponseModel) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


