# MachineResponseModel

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
**InMaintenanceMode** | Pointer to **bool** | Denotes if the machine is in maintenance mode. Machines in maintenance mode will not accept new sessions. | [optional] 
**MaintenanceModeReason** | Pointer to [**MaintenanceModeReason**](MaintenanceModeReason.md) |  | [optional] 
**DrainingUntilShutdown** | Pointer to **bool** | Denotes if the machine is placed to drain until shutdown | [optional] 
**IPAddress** | Pointer to **NullableString** | The IP address of the machine. | [optional] 
**IsAssigned** | Pointer to **NullableBool** | Denotes whether a private desktop has been assigned to a user/users, or a client name/address. Users can be assigned explicitly or by assigning on first use of the machine. Only relevant for privately assigned machines. | [optional] 
**MachineType** | Pointer to [**MachineType**](MachineType.md) |  | [optional] 
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
**PowerState** | Pointer to [**PowerState**](PowerState.md) |  | [optional] 
**ProvisioningType** | Pointer to [**ProvisioningType**](ProvisioningType.md) |  | [optional] 
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
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) |  | [optional] 
**SessionUserName** | Pointer to **NullableString** | The session user name. | [optional] 
**Sid** | Pointer to **NullableString** | The SID of the machine. Used to be: DesktopSid or SID (based on the context) | [optional] 
**SummaryState** | Pointer to [**SummaryState**](SummaryState.md) |  | [optional] 
**WillShutdownAfterUse** | Pointer to **NullableBool** | Flag indicating if this machine is tainted and will be shut down after all sessions on the machine have ended. This flag is only ever non-null on power-managed, single-session machines. | [optional] 
**WindowsConnectionSetting** | Pointer to [**WindowsConnectionSetting**](WindowsConnectionSetting.md) |  | [optional] 
**WindowsActivationStatus** | Pointer to **NullableString** | Windows Activation Status of the provisioned virtual machine. This is by default set to unknown and is updated when the vda registers with the DDC. | [optional] 
**WindowsActivationStatusErrorCode** | Pointer to **NullableString** | The default value is Unknown. If successful, it displays as Success, else it displays with the associated error code. | [optional] 
**WindowsActivationStatusVirtualMachineError** | Pointer to **NullableString** | Provides a detailed information of potential error message seen while activating the windows system. | [optional] 
**WindowsActivationTypeProvisionedVirtualMachine** | Pointer to **NullableString** | Windows Activation Type of the Provisioned Virtual Machine. The default value is unknown and the field is populated once the virtual machine registers with DDC. This is supported only on 2303 and successive versions. | [optional] 
**Zone** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**SupportedPowerActions** | Pointer to [**[]SupportedPowerAction**](SupportedPowerAction.md) | A list of power actions supported by this machine. | [optional] 
**FaultState** | Pointer to [**FaultState**](FaultState.md) |  | [optional] 
**ContainerMetadata** | Pointer to [**ContainerMetadataModel**](ContainerMetadataModel.md) |  | [optional] 
**Tags** | Pointer to **[]string** | The tags for this machine. | [optional] 
**UpgradeType** | Pointer to [**VdaUpgradeType**](VdaUpgradeType.md) |  | [optional] 
**UpgradeState** | Pointer to [**VdaUpgradeState**](VdaUpgradeState.md) |  | [optional] 
**MachineConfigurationOutOfSync** | Pointer to **NullableBool** | Flag indicating whether the machine&#39;s configuration is out of sync with the catalog&#39;s latest configuration | [optional] 
**UpgradeDetail** | Pointer to [**MachineUpgradeDetail**](MachineUpgradeDetail.md) |  | [optional] 

## Methods

### NewMachineResponseModel

`func NewMachineResponseModel(id string, ) *MachineResponseModel`

NewMachineResponseModel instantiates a new MachineResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineResponseModelWithDefaults

`func NewMachineResponseModelWithDefaults() *MachineResponseModel`

NewMachineResponseModelWithDefaults instantiates a new MachineResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetMachineCatalog

`func (o *MachineResponseModel) GetMachineCatalog() RefResponseModel`

GetMachineCatalog returns the MachineCatalog field if non-nil, zero value otherwise.

### GetMachineCatalogOk

`func (o *MachineResponseModel) GetMachineCatalogOk() (*RefResponseModel, bool)`

GetMachineCatalogOk returns a tuple with the MachineCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCatalog

`func (o *MachineResponseModel) SetMachineCatalog(v RefResponseModel)`

SetMachineCatalog sets MachineCatalog field to given value.

### HasMachineCatalog

`func (o *MachineResponseModel) HasMachineCatalog() bool`

HasMachineCatalog returns a boolean if a field has been set.

### GetName

`func (o *MachineResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MachineResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MachineResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUid

`func (o *MachineResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MachineResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MachineResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MachineResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *MachineResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *MachineResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetAgentVersion

`func (o *MachineResponseModel) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *MachineResponseModel) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *MachineResponseModel) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *MachineResponseModel) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *MachineResponseModel) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *MachineResponseModel) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil
### GetAllocationType

`func (o *MachineResponseModel) GetAllocationType() AllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *MachineResponseModel) GetAllocationTypeOk() (*AllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *MachineResponseModel) SetAllocationType(v AllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *MachineResponseModel) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetApplicationsInUse

`func (o *MachineResponseModel) GetApplicationsInUse() []RefResponseModel`

GetApplicationsInUse returns the ApplicationsInUse field if non-nil, zero value otherwise.

### GetApplicationsInUseOk

`func (o *MachineResponseModel) GetApplicationsInUseOk() (*[]RefResponseModel, bool)`

GetApplicationsInUseOk returns a tuple with the ApplicationsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsInUse

`func (o *MachineResponseModel) SetApplicationsInUse(v []RefResponseModel)`

SetApplicationsInUse sets ApplicationsInUse field to given value.

### HasApplicationsInUse

`func (o *MachineResponseModel) HasApplicationsInUse() bool`

HasApplicationsInUse returns a boolean if a field has been set.

### SetApplicationsInUseNil

`func (o *MachineResponseModel) SetApplicationsInUseNil(b bool)`

 SetApplicationsInUseNil sets the value for ApplicationsInUse to be an explicit nil

### UnsetApplicationsInUse
`func (o *MachineResponseModel) UnsetApplicationsInUse()`

UnsetApplicationsInUse ensures that no value is present for ApplicationsInUse, not even an explicit nil
### GetAssignedUsers

`func (o *MachineResponseModel) GetAssignedUsers() []IdentityUserResponseModel`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *MachineResponseModel) GetAssignedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *MachineResponseModel) SetAssignedUsers(v []IdentityUserResponseModel)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *MachineResponseModel) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.

### SetAssignedUsersNil

`func (o *MachineResponseModel) SetAssignedUsersNil(b bool)`

 SetAssignedUsersNil sets the value for AssignedUsers to be an explicit nil

### UnsetAssignedUsers
`func (o *MachineResponseModel) UnsetAssignedUsers()`

UnsetAssignedUsers ensures that no value is present for AssignedUsers, not even an explicit nil
### GetAssociatedUsers

`func (o *MachineResponseModel) GetAssociatedUsers() []IdentityUserResponseModel`

GetAssociatedUsers returns the AssociatedUsers field if non-nil, zero value otherwise.

### GetAssociatedUsersOk

`func (o *MachineResponseModel) GetAssociatedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetAssociatedUsersOk returns a tuple with the AssociatedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedUsers

`func (o *MachineResponseModel) SetAssociatedUsers(v []IdentityUserResponseModel)`

SetAssociatedUsers sets AssociatedUsers field to given value.

### HasAssociatedUsers

`func (o *MachineResponseModel) HasAssociatedUsers() bool`

HasAssociatedUsers returns a boolean if a field has been set.

### SetAssociatedUsersNil

`func (o *MachineResponseModel) SetAssociatedUsersNil(b bool)`

 SetAssociatedUsersNil sets the value for AssociatedUsers to be an explicit nil

### UnsetAssociatedUsers
`func (o *MachineResponseModel) UnsetAssociatedUsers()`

UnsetAssociatedUsers ensures that no value is present for AssociatedUsers, not even an explicit nil
### GetAzureAdJoinedMode

`func (o *MachineResponseModel) GetAzureAdJoinedMode() AzureAdJoinedMode`

GetAzureAdJoinedMode returns the AzureAdJoinedMode field if non-nil, zero value otherwise.

### GetAzureAdJoinedModeOk

`func (o *MachineResponseModel) GetAzureAdJoinedModeOk() (*AzureAdJoinedMode, bool)`

GetAzureAdJoinedModeOk returns a tuple with the AzureAdJoinedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdJoinedMode

`func (o *MachineResponseModel) SetAzureAdJoinedMode(v AzureAdJoinedMode)`

SetAzureAdJoinedMode sets AzureAdJoinedMode field to given value.

### HasAzureAdJoinedMode

`func (o *MachineResponseModel) HasAzureAdJoinedMode() bool`

HasAzureAdJoinedMode returns a boolean if a field has been set.

### GetContainerScopes

`func (o *MachineResponseModel) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *MachineResponseModel) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *MachineResponseModel) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.

### HasContainerScopes

`func (o *MachineResponseModel) HasContainerScopes() bool`

HasContainerScopes returns a boolean if a field has been set.

### SetContainerScopesNil

`func (o *MachineResponseModel) SetContainerScopesNil(b bool)`

 SetContainerScopesNil sets the value for ContainerScopes to be an explicit nil

### UnsetContainerScopes
`func (o *MachineResponseModel) UnsetContainerScopes()`

UnsetContainerScopes ensures that no value is present for ContainerScopes, not even an explicit nil
### GetControllerDnsName

`func (o *MachineResponseModel) GetControllerDnsName() string`

GetControllerDnsName returns the ControllerDnsName field if non-nil, zero value otherwise.

### GetControllerDnsNameOk

`func (o *MachineResponseModel) GetControllerDnsNameOk() (*string, bool)`

GetControllerDnsNameOk returns a tuple with the ControllerDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDnsName

`func (o *MachineResponseModel) SetControllerDnsName(v string)`

SetControllerDnsName sets ControllerDnsName field to given value.

### HasControllerDnsName

`func (o *MachineResponseModel) HasControllerDnsName() bool`

HasControllerDnsName returns a boolean if a field has been set.

### SetControllerDnsNameNil

`func (o *MachineResponseModel) SetControllerDnsNameNil(b bool)`

 SetControllerDnsNameNil sets the value for ControllerDnsName to be an explicit nil

### UnsetControllerDnsName
`func (o *MachineResponseModel) UnsetControllerDnsName()`

UnsetControllerDnsName ensures that no value is present for ControllerDnsName, not even an explicit nil
### GetDeliveryGroup

`func (o *MachineResponseModel) GetDeliveryGroup() RefResponseModel`

GetDeliveryGroup returns the DeliveryGroup field if non-nil, zero value otherwise.

### GetDeliveryGroupOk

`func (o *MachineResponseModel) GetDeliveryGroupOk() (*RefResponseModel, bool)`

GetDeliveryGroupOk returns a tuple with the DeliveryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroup

`func (o *MachineResponseModel) SetDeliveryGroup(v RefResponseModel)`

SetDeliveryGroup sets DeliveryGroup field to given value.

### HasDeliveryGroup

`func (o *MachineResponseModel) HasDeliveryGroup() bool`

HasDeliveryGroup returns a boolean if a field has been set.

### GetDeliveryType

`func (o *MachineResponseModel) GetDeliveryType() DeliveryKind`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *MachineResponseModel) GetDeliveryTypeOk() (*DeliveryKind, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *MachineResponseModel) SetDeliveryType(v DeliveryKind)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *MachineResponseModel) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDescription

`func (o *MachineResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MachineResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MachineResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MachineResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MachineResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MachineResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesktopConditions

`func (o *MachineResponseModel) GetDesktopConditions() []DesktopCondition`

GetDesktopConditions returns the DesktopConditions field if non-nil, zero value otherwise.

### GetDesktopConditionsOk

`func (o *MachineResponseModel) GetDesktopConditionsOk() (*[]DesktopCondition, bool)`

GetDesktopConditionsOk returns a tuple with the DesktopConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopConditions

`func (o *MachineResponseModel) SetDesktopConditions(v []DesktopCondition)`

SetDesktopConditions sets DesktopConditions field to given value.

### HasDesktopConditions

`func (o *MachineResponseModel) HasDesktopConditions() bool`

HasDesktopConditions returns a boolean if a field has been set.

### SetDesktopConditionsNil

`func (o *MachineResponseModel) SetDesktopConditionsNil(b bool)`

 SetDesktopConditionsNil sets the value for DesktopConditions to be an explicit nil

### UnsetDesktopConditions
`func (o *MachineResponseModel) UnsetDesktopConditions()`

UnsetDesktopConditions ensures that no value is present for DesktopConditions, not even an explicit nil
### GetDnsName

`func (o *MachineResponseModel) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *MachineResponseModel) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *MachineResponseModel) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *MachineResponseModel) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### SetDnsNameNil

`func (o *MachineResponseModel) SetDnsNameNil(b bool)`

 SetDnsNameNil sets the value for DnsName to be an explicit nil

### UnsetDnsName
`func (o *MachineResponseModel) UnsetDnsName()`

UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
### GetFunctionalLevel

`func (o *MachineResponseModel) GetFunctionalLevel() FunctionalLevel`

GetFunctionalLevel returns the FunctionalLevel field if non-nil, zero value otherwise.

### GetFunctionalLevelOk

`func (o *MachineResponseModel) GetFunctionalLevelOk() (*FunctionalLevel, bool)`

GetFunctionalLevelOk returns a tuple with the FunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalLevel

`func (o *MachineResponseModel) SetFunctionalLevel(v FunctionalLevel)`

SetFunctionalLevel sets FunctionalLevel field to given value.

### HasFunctionalLevel

`func (o *MachineResponseModel) HasFunctionalLevel() bool`

HasFunctionalLevel returns a boolean if a field has been set.

### GetHosting

`func (o *MachineResponseModel) GetHosting() MachineHostingResponseModel`

GetHosting returns the Hosting field if non-nil, zero value otherwise.

### GetHostingOk

`func (o *MachineResponseModel) GetHostingOk() (*MachineHostingResponseModel, bool)`

GetHostingOk returns a tuple with the Hosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosting

`func (o *MachineResponseModel) SetHosting(v MachineHostingResponseModel)`

SetHosting sets Hosting field to given value.

### HasHosting

`func (o *MachineResponseModel) HasHosting() bool`

HasHosting returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *MachineResponseModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *MachineResponseModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *MachineResponseModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *MachineResponseModel) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### GetMaintenanceModeReason

`func (o *MachineResponseModel) GetMaintenanceModeReason() MaintenanceModeReason`

GetMaintenanceModeReason returns the MaintenanceModeReason field if non-nil, zero value otherwise.

### GetMaintenanceModeReasonOk

`func (o *MachineResponseModel) GetMaintenanceModeReasonOk() (*MaintenanceModeReason, bool)`

GetMaintenanceModeReasonOk returns a tuple with the MaintenanceModeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceModeReason

`func (o *MachineResponseModel) SetMaintenanceModeReason(v MaintenanceModeReason)`

SetMaintenanceModeReason sets MaintenanceModeReason field to given value.

### HasMaintenanceModeReason

`func (o *MachineResponseModel) HasMaintenanceModeReason() bool`

HasMaintenanceModeReason returns a boolean if a field has been set.

### GetDrainingUntilShutdown

`func (o *MachineResponseModel) GetDrainingUntilShutdown() bool`

GetDrainingUntilShutdown returns the DrainingUntilShutdown field if non-nil, zero value otherwise.

### GetDrainingUntilShutdownOk

`func (o *MachineResponseModel) GetDrainingUntilShutdownOk() (*bool, bool)`

GetDrainingUntilShutdownOk returns a tuple with the DrainingUntilShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainingUntilShutdown

`func (o *MachineResponseModel) SetDrainingUntilShutdown(v bool)`

SetDrainingUntilShutdown sets DrainingUntilShutdown field to given value.

### HasDrainingUntilShutdown

`func (o *MachineResponseModel) HasDrainingUntilShutdown() bool`

HasDrainingUntilShutdown returns a boolean if a field has been set.

### GetIPAddress

`func (o *MachineResponseModel) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *MachineResponseModel) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *MachineResponseModel) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *MachineResponseModel) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### SetIPAddressNil

`func (o *MachineResponseModel) SetIPAddressNil(b bool)`

 SetIPAddressNil sets the value for IPAddress to be an explicit nil

### UnsetIPAddress
`func (o *MachineResponseModel) UnsetIPAddress()`

UnsetIPAddress ensures that no value is present for IPAddress, not even an explicit nil
### GetIsAssigned

`func (o *MachineResponseModel) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MachineResponseModel) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *MachineResponseModel) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *MachineResponseModel) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### SetIsAssignedNil

`func (o *MachineResponseModel) SetIsAssignedNil(b bool)`

 SetIsAssignedNil sets the value for IsAssigned to be an explicit nil

### UnsetIsAssigned
`func (o *MachineResponseModel) UnsetIsAssigned()`

UnsetIsAssigned ensures that no value is present for IsAssigned, not even an explicit nil
### GetMachineType

`func (o *MachineResponseModel) GetMachineType() MachineType`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *MachineResponseModel) GetMachineTypeOk() (*MachineType, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *MachineResponseModel) SetMachineType(v MachineType)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *MachineResponseModel) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetLastConnectionFailure

`func (o *MachineResponseModel) GetLastConnectionFailure() ConnectionFailureReason`

GetLastConnectionFailure returns the LastConnectionFailure field if non-nil, zero value otherwise.

### GetLastConnectionFailureOk

`func (o *MachineResponseModel) GetLastConnectionFailureOk() (*ConnectionFailureReason, bool)`

GetLastConnectionFailureOk returns a tuple with the LastConnectionFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionFailure

`func (o *MachineResponseModel) SetLastConnectionFailure(v ConnectionFailureReason)`

SetLastConnectionFailure sets LastConnectionFailure field to given value.

### HasLastConnectionFailure

`func (o *MachineResponseModel) HasLastConnectionFailure() bool`

HasLastConnectionFailure returns a boolean if a field has been set.

### GetLastConnectionTime

`func (o *MachineResponseModel) GetLastConnectionTime() string`

GetLastConnectionTime returns the LastConnectionTime field if non-nil, zero value otherwise.

### GetLastConnectionTimeOk

`func (o *MachineResponseModel) GetLastConnectionTimeOk() (*string, bool)`

GetLastConnectionTimeOk returns a tuple with the LastConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionTime

`func (o *MachineResponseModel) SetLastConnectionTime(v string)`

SetLastConnectionTime sets LastConnectionTime field to given value.

### HasLastConnectionTime

`func (o *MachineResponseModel) HasLastConnectionTime() bool`

HasLastConnectionTime returns a boolean if a field has been set.

### SetLastConnectionTimeNil

`func (o *MachineResponseModel) SetLastConnectionTimeNil(b bool)`

 SetLastConnectionTimeNil sets the value for LastConnectionTime to be an explicit nil

### UnsetLastConnectionTime
`func (o *MachineResponseModel) UnsetLastConnectionTime()`

UnsetLastConnectionTime ensures that no value is present for LastConnectionTime, not even an explicit nil
### GetFormattedLastConnectionTime

`func (o *MachineResponseModel) GetFormattedLastConnectionTime() string`

GetFormattedLastConnectionTime returns the FormattedLastConnectionTime field if non-nil, zero value otherwise.

### GetFormattedLastConnectionTimeOk

`func (o *MachineResponseModel) GetFormattedLastConnectionTimeOk() (*string, bool)`

GetFormattedLastConnectionTimeOk returns a tuple with the FormattedLastConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastConnectionTime

`func (o *MachineResponseModel) SetFormattedLastConnectionTime(v string)`

SetFormattedLastConnectionTime sets FormattedLastConnectionTime field to given value.

### HasFormattedLastConnectionTime

`func (o *MachineResponseModel) HasFormattedLastConnectionTime() bool`

HasFormattedLastConnectionTime returns a boolean if a field has been set.

### SetFormattedLastConnectionTimeNil

`func (o *MachineResponseModel) SetFormattedLastConnectionTimeNil(b bool)`

 SetFormattedLastConnectionTimeNil sets the value for FormattedLastConnectionTime to be an explicit nil

### UnsetFormattedLastConnectionTime
`func (o *MachineResponseModel) UnsetFormattedLastConnectionTime()`

UnsetFormattedLastConnectionTime ensures that no value is present for FormattedLastConnectionTime, not even an explicit nil
### GetLastConnectionUser

`func (o *MachineResponseModel) GetLastConnectionUser() IdentityUserResponseModel`

GetLastConnectionUser returns the LastConnectionUser field if non-nil, zero value otherwise.

### GetLastConnectionUserOk

`func (o *MachineResponseModel) GetLastConnectionUserOk() (*IdentityUserResponseModel, bool)`

GetLastConnectionUserOk returns a tuple with the LastConnectionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionUser

`func (o *MachineResponseModel) SetLastConnectionUser(v IdentityUserResponseModel)`

SetLastConnectionUser sets LastConnectionUser field to given value.

### HasLastConnectionUser

`func (o *MachineResponseModel) HasLastConnectionUser() bool`

HasLastConnectionUser returns a boolean if a field has been set.

### GetLastDeregistrationReason

`func (o *MachineResponseModel) GetLastDeregistrationReason() DeregistrationReason`

GetLastDeregistrationReason returns the LastDeregistrationReason field if non-nil, zero value otherwise.

### GetLastDeregistrationReasonOk

`func (o *MachineResponseModel) GetLastDeregistrationReasonOk() (*DeregistrationReason, bool)`

GetLastDeregistrationReasonOk returns a tuple with the LastDeregistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationReason

`func (o *MachineResponseModel) SetLastDeregistrationReason(v DeregistrationReason)`

SetLastDeregistrationReason sets LastDeregistrationReason field to given value.

### HasLastDeregistrationReason

`func (o *MachineResponseModel) HasLastDeregistrationReason() bool`

HasLastDeregistrationReason returns a boolean if a field has been set.

### GetLastDeregistrationTime

`func (o *MachineResponseModel) GetLastDeregistrationTime() string`

GetLastDeregistrationTime returns the LastDeregistrationTime field if non-nil, zero value otherwise.

### GetLastDeregistrationTimeOk

`func (o *MachineResponseModel) GetLastDeregistrationTimeOk() (*string, bool)`

GetLastDeregistrationTimeOk returns a tuple with the LastDeregistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationTime

`func (o *MachineResponseModel) SetLastDeregistrationTime(v string)`

SetLastDeregistrationTime sets LastDeregistrationTime field to given value.

### HasLastDeregistrationTime

`func (o *MachineResponseModel) HasLastDeregistrationTime() bool`

HasLastDeregistrationTime returns a boolean if a field has been set.

### SetLastDeregistrationTimeNil

`func (o *MachineResponseModel) SetLastDeregistrationTimeNil(b bool)`

 SetLastDeregistrationTimeNil sets the value for LastDeregistrationTime to be an explicit nil

### UnsetLastDeregistrationTime
`func (o *MachineResponseModel) UnsetLastDeregistrationTime()`

UnsetLastDeregistrationTime ensures that no value is present for LastDeregistrationTime, not even an explicit nil
### GetFormattedLastDeregistrationTime

`func (o *MachineResponseModel) GetFormattedLastDeregistrationTime() string`

GetFormattedLastDeregistrationTime returns the FormattedLastDeregistrationTime field if non-nil, zero value otherwise.

### GetFormattedLastDeregistrationTimeOk

`func (o *MachineResponseModel) GetFormattedLastDeregistrationTimeOk() (*string, bool)`

GetFormattedLastDeregistrationTimeOk returns a tuple with the FormattedLastDeregistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastDeregistrationTime

`func (o *MachineResponseModel) SetFormattedLastDeregistrationTime(v string)`

SetFormattedLastDeregistrationTime sets FormattedLastDeregistrationTime field to given value.

### HasFormattedLastDeregistrationTime

`func (o *MachineResponseModel) HasFormattedLastDeregistrationTime() bool`

HasFormattedLastDeregistrationTime returns a boolean if a field has been set.

### SetFormattedLastDeregistrationTimeNil

`func (o *MachineResponseModel) SetFormattedLastDeregistrationTimeNil(b bool)`

 SetFormattedLastDeregistrationTimeNil sets the value for FormattedLastDeregistrationTime to be an explicit nil

### UnsetFormattedLastDeregistrationTime
`func (o *MachineResponseModel) UnsetFormattedLastDeregistrationTime()`

UnsetFormattedLastDeregistrationTime ensures that no value is present for FormattedLastDeregistrationTime, not even an explicit nil
### GetLastErrorReason

`func (o *MachineResponseModel) GetLastErrorReason() string`

GetLastErrorReason returns the LastErrorReason field if non-nil, zero value otherwise.

### GetLastErrorReasonOk

`func (o *MachineResponseModel) GetLastErrorReasonOk() (*string, bool)`

GetLastErrorReasonOk returns a tuple with the LastErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorReason

`func (o *MachineResponseModel) SetLastErrorReason(v string)`

SetLastErrorReason sets LastErrorReason field to given value.

### HasLastErrorReason

`func (o *MachineResponseModel) HasLastErrorReason() bool`

HasLastErrorReason returns a boolean if a field has been set.

### SetLastErrorReasonNil

`func (o *MachineResponseModel) SetLastErrorReasonNil(b bool)`

 SetLastErrorReasonNil sets the value for LastErrorReason to be an explicit nil

### UnsetLastErrorReason
`func (o *MachineResponseModel) UnsetLastErrorReason()`

UnsetLastErrorReason ensures that no value is present for LastErrorReason, not even an explicit nil
### GetLastErrorTime

`func (o *MachineResponseModel) GetLastErrorTime() string`

GetLastErrorTime returns the LastErrorTime field if non-nil, zero value otherwise.

### GetLastErrorTimeOk

`func (o *MachineResponseModel) GetLastErrorTimeOk() (*string, bool)`

GetLastErrorTimeOk returns a tuple with the LastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorTime

`func (o *MachineResponseModel) SetLastErrorTime(v string)`

SetLastErrorTime sets LastErrorTime field to given value.

### HasLastErrorTime

`func (o *MachineResponseModel) HasLastErrorTime() bool`

HasLastErrorTime returns a boolean if a field has been set.

### SetLastErrorTimeNil

`func (o *MachineResponseModel) SetLastErrorTimeNil(b bool)`

 SetLastErrorTimeNil sets the value for LastErrorTime to be an explicit nil

### UnsetLastErrorTime
`func (o *MachineResponseModel) UnsetLastErrorTime()`

UnsetLastErrorTime ensures that no value is present for LastErrorTime, not even an explicit nil
### GetFormattedLastErrorTime

`func (o *MachineResponseModel) GetFormattedLastErrorTime() string`

GetFormattedLastErrorTime returns the FormattedLastErrorTime field if non-nil, zero value otherwise.

### GetFormattedLastErrorTimeOk

`func (o *MachineResponseModel) GetFormattedLastErrorTimeOk() (*string, bool)`

GetFormattedLastErrorTimeOk returns a tuple with the FormattedLastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastErrorTime

`func (o *MachineResponseModel) SetFormattedLastErrorTime(v string)`

SetFormattedLastErrorTime sets FormattedLastErrorTime field to given value.

### HasFormattedLastErrorTime

`func (o *MachineResponseModel) HasFormattedLastErrorTime() bool`

HasFormattedLastErrorTime returns a boolean if a field has been set.

### SetFormattedLastErrorTimeNil

`func (o *MachineResponseModel) SetFormattedLastErrorTimeNil(b bool)`

 SetFormattedLastErrorTimeNil sets the value for FormattedLastErrorTime to be an explicit nil

### UnsetFormattedLastErrorTime
`func (o *MachineResponseModel) UnsetFormattedLastErrorTime()`

UnsetFormattedLastErrorTime ensures that no value is present for FormattedLastErrorTime, not even an explicit nil
### GetLoadIndex

`func (o *MachineResponseModel) GetLoadIndex() int32`

GetLoadIndex returns the LoadIndex field if non-nil, zero value otherwise.

### GetLoadIndexOk

`func (o *MachineResponseModel) GetLoadIndexOk() (*int32, bool)`

GetLoadIndexOk returns a tuple with the LoadIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadIndex

`func (o *MachineResponseModel) SetLoadIndex(v int32)`

SetLoadIndex sets LoadIndex field to given value.

### HasLoadIndex

`func (o *MachineResponseModel) HasLoadIndex() bool`

HasLoadIndex returns a boolean if a field has been set.

### SetLoadIndexNil

`func (o *MachineResponseModel) SetLoadIndexNil(b bool)`

 SetLoadIndexNil sets the value for LoadIndex to be an explicit nil

### UnsetLoadIndex
`func (o *MachineResponseModel) UnsetLoadIndex()`

UnsetLoadIndex ensures that no value is present for LoadIndex, not even an explicit nil
### GetMachineUnavailableReason

`func (o *MachineResponseModel) GetMachineUnavailableReason() MachineUnavailableReason`

GetMachineUnavailableReason returns the MachineUnavailableReason field if non-nil, zero value otherwise.

### GetMachineUnavailableReasonOk

`func (o *MachineResponseModel) GetMachineUnavailableReasonOk() (*MachineUnavailableReason, bool)`

GetMachineUnavailableReasonOk returns a tuple with the MachineUnavailableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineUnavailableReason

`func (o *MachineResponseModel) SetMachineUnavailableReason(v MachineUnavailableReason)`

SetMachineUnavailableReason sets MachineUnavailableReason field to given value.

### HasMachineUnavailableReason

`func (o *MachineResponseModel) HasMachineUnavailableReason() bool`

HasMachineUnavailableReason returns a boolean if a field has been set.

### GetOSType

`func (o *MachineResponseModel) GetOSType() string`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *MachineResponseModel) GetOSTypeOk() (*string, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *MachineResponseModel) SetOSType(v string)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *MachineResponseModel) HasOSType() bool`

HasOSType returns a boolean if a field has been set.

### SetOSTypeNil

`func (o *MachineResponseModel) SetOSTypeNil(b bool)`

 SetOSTypeNil sets the value for OSType to be an explicit nil

### UnsetOSType
`func (o *MachineResponseModel) UnsetOSType()`

UnsetOSType ensures that no value is present for OSType, not even an explicit nil
### GetOSVersion

`func (o *MachineResponseModel) GetOSVersion() string`

GetOSVersion returns the OSVersion field if non-nil, zero value otherwise.

### GetOSVersionOk

`func (o *MachineResponseModel) GetOSVersionOk() (*string, bool)`

GetOSVersionOk returns a tuple with the OSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSVersion

`func (o *MachineResponseModel) SetOSVersion(v string)`

SetOSVersion sets OSVersion field to given value.

### HasOSVersion

`func (o *MachineResponseModel) HasOSVersion() bool`

HasOSVersion returns a boolean if a field has been set.

### SetOSVersionNil

`func (o *MachineResponseModel) SetOSVersionNil(b bool)`

 SetOSVersionNil sets the value for OSVersion to be an explicit nil

### UnsetOSVersion
`func (o *MachineResponseModel) UnsetOSVersion()`

UnsetOSVersion ensures that no value is present for OSVersion, not even an explicit nil
### GetPersistUserChanges

`func (o *MachineResponseModel) GetPersistUserChanges() PersistChanges`

GetPersistUserChanges returns the PersistUserChanges field if non-nil, zero value otherwise.

### GetPersistUserChangesOk

`func (o *MachineResponseModel) GetPersistUserChangesOk() (*PersistChanges, bool)`

GetPersistUserChangesOk returns a tuple with the PersistUserChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistUserChanges

`func (o *MachineResponseModel) SetPersistUserChanges(v PersistChanges)`

SetPersistUserChanges sets PersistUserChanges field to given value.

### HasPersistUserChanges

`func (o *MachineResponseModel) HasPersistUserChanges() bool`

HasPersistUserChanges returns a boolean if a field has been set.

### GetPowerActionPending

`func (o *MachineResponseModel) GetPowerActionPending() bool`

GetPowerActionPending returns the PowerActionPending field if non-nil, zero value otherwise.

### GetPowerActionPendingOk

`func (o *MachineResponseModel) GetPowerActionPendingOk() (*bool, bool)`

GetPowerActionPendingOk returns a tuple with the PowerActionPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerActionPending

`func (o *MachineResponseModel) SetPowerActionPending(v bool)`

SetPowerActionPending sets PowerActionPending field to given value.

### HasPowerActionPending

`func (o *MachineResponseModel) HasPowerActionPending() bool`

HasPowerActionPending returns a boolean if a field has been set.

### SetPowerActionPendingNil

`func (o *MachineResponseModel) SetPowerActionPendingNil(b bool)`

 SetPowerActionPendingNil sets the value for PowerActionPending to be an explicit nil

### UnsetPowerActionPending
`func (o *MachineResponseModel) UnsetPowerActionPending()`

UnsetPowerActionPending ensures that no value is present for PowerActionPending, not even an explicit nil
### GetPowerState

`func (o *MachineResponseModel) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *MachineResponseModel) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *MachineResponseModel) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *MachineResponseModel) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetProvisioningType

`func (o *MachineResponseModel) GetProvisioningType() ProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *MachineResponseModel) GetProvisioningTypeOk() (*ProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *MachineResponseModel) SetProvisioningType(v ProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.

### HasProvisioningType

`func (o *MachineResponseModel) HasProvisioningType() bool`

HasProvisioningType returns a boolean if a field has been set.

### GetPublishedApplications

`func (o *MachineResponseModel) GetPublishedApplications() []string`

GetPublishedApplications returns the PublishedApplications field if non-nil, zero value otherwise.

### GetPublishedApplicationsOk

`func (o *MachineResponseModel) GetPublishedApplicationsOk() (*[]string, bool)`

GetPublishedApplicationsOk returns a tuple with the PublishedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedApplications

`func (o *MachineResponseModel) SetPublishedApplications(v []string)`

SetPublishedApplications sets PublishedApplications field to given value.

### HasPublishedApplications

`func (o *MachineResponseModel) HasPublishedApplications() bool`

HasPublishedApplications returns a boolean if a field has been set.

### SetPublishedApplicationsNil

`func (o *MachineResponseModel) SetPublishedApplicationsNil(b bool)`

 SetPublishedApplicationsNil sets the value for PublishedApplications to be an explicit nil

### UnsetPublishedApplications
`func (o *MachineResponseModel) UnsetPublishedApplications()`

UnsetPublishedApplications ensures that no value is present for PublishedApplications, not even an explicit nil
### GetPublishedName

`func (o *MachineResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *MachineResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *MachineResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *MachineResponseModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *MachineResponseModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *MachineResponseModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetRegistrationState

`func (o *MachineResponseModel) GetRegistrationState() RegistrationState`

GetRegistrationState returns the RegistrationState field if non-nil, zero value otherwise.

### GetRegistrationStateOk

`func (o *MachineResponseModel) GetRegistrationStateOk() (*RegistrationState, bool)`

GetRegistrationStateOk returns a tuple with the RegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationState

`func (o *MachineResponseModel) SetRegistrationState(v RegistrationState)`

SetRegistrationState sets RegistrationState field to given value.

### HasRegistrationState

`func (o *MachineResponseModel) HasRegistrationState() bool`

HasRegistrationState returns a boolean if a field has been set.

### GetScheduledReboot

`func (o *MachineResponseModel) GetScheduledReboot() ScheduledReboot`

GetScheduledReboot returns the ScheduledReboot field if non-nil, zero value otherwise.

### GetScheduledRebootOk

`func (o *MachineResponseModel) GetScheduledRebootOk() (*ScheduledReboot, bool)`

GetScheduledRebootOk returns a tuple with the ScheduledReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReboot

`func (o *MachineResponseModel) SetScheduledReboot(v ScheduledReboot)`

SetScheduledReboot sets ScheduledReboot field to given value.

### HasScheduledReboot

`func (o *MachineResponseModel) HasScheduledReboot() bool`

HasScheduledReboot returns a boolean if a field has been set.

### GetSessionClientAddress

`func (o *MachineResponseModel) GetSessionClientAddress() string`

GetSessionClientAddress returns the SessionClientAddress field if non-nil, zero value otherwise.

### GetSessionClientAddressOk

`func (o *MachineResponseModel) GetSessionClientAddressOk() (*string, bool)`

GetSessionClientAddressOk returns a tuple with the SessionClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientAddress

`func (o *MachineResponseModel) SetSessionClientAddress(v string)`

SetSessionClientAddress sets SessionClientAddress field to given value.

### HasSessionClientAddress

`func (o *MachineResponseModel) HasSessionClientAddress() bool`

HasSessionClientAddress returns a boolean if a field has been set.

### SetSessionClientAddressNil

`func (o *MachineResponseModel) SetSessionClientAddressNil(b bool)`

 SetSessionClientAddressNil sets the value for SessionClientAddress to be an explicit nil

### UnsetSessionClientAddress
`func (o *MachineResponseModel) UnsetSessionClientAddress()`

UnsetSessionClientAddress ensures that no value is present for SessionClientAddress, not even an explicit nil
### GetSessionClientName

`func (o *MachineResponseModel) GetSessionClientName() string`

GetSessionClientName returns the SessionClientName field if non-nil, zero value otherwise.

### GetSessionClientNameOk

`func (o *MachineResponseModel) GetSessionClientNameOk() (*string, bool)`

GetSessionClientNameOk returns a tuple with the SessionClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientName

`func (o *MachineResponseModel) SetSessionClientName(v string)`

SetSessionClientName sets SessionClientName field to given value.

### HasSessionClientName

`func (o *MachineResponseModel) HasSessionClientName() bool`

HasSessionClientName returns a boolean if a field has been set.

### SetSessionClientNameNil

`func (o *MachineResponseModel) SetSessionClientNameNil(b bool)`

 SetSessionClientNameNil sets the value for SessionClientName to be an explicit nil

### UnsetSessionClientName
`func (o *MachineResponseModel) UnsetSessionClientName()`

UnsetSessionClientName ensures that no value is present for SessionClientName, not even an explicit nil
### GetSessionClientVersion

`func (o *MachineResponseModel) GetSessionClientVersion() string`

GetSessionClientVersion returns the SessionClientVersion field if non-nil, zero value otherwise.

### GetSessionClientVersionOk

`func (o *MachineResponseModel) GetSessionClientVersionOk() (*string, bool)`

GetSessionClientVersionOk returns a tuple with the SessionClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientVersion

`func (o *MachineResponseModel) SetSessionClientVersion(v string)`

SetSessionClientVersion sets SessionClientVersion field to given value.

### HasSessionClientVersion

`func (o *MachineResponseModel) HasSessionClientVersion() bool`

HasSessionClientVersion returns a boolean if a field has been set.

### SetSessionClientVersionNil

`func (o *MachineResponseModel) SetSessionClientVersionNil(b bool)`

 SetSessionClientVersionNil sets the value for SessionClientVersion to be an explicit nil

### UnsetSessionClientVersion
`func (o *MachineResponseModel) UnsetSessionClientVersion()`

UnsetSessionClientVersion ensures that no value is present for SessionClientVersion, not even an explicit nil
### GetSessionConnectedViaHostName

`func (o *MachineResponseModel) GetSessionConnectedViaHostName() string`

GetSessionConnectedViaHostName returns the SessionConnectedViaHostName field if non-nil, zero value otherwise.

### GetSessionConnectedViaHostNameOk

`func (o *MachineResponseModel) GetSessionConnectedViaHostNameOk() (*string, bool)`

GetSessionConnectedViaHostNameOk returns a tuple with the SessionConnectedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConnectedViaHostName

`func (o *MachineResponseModel) SetSessionConnectedViaHostName(v string)`

SetSessionConnectedViaHostName sets SessionConnectedViaHostName field to given value.

### HasSessionConnectedViaHostName

`func (o *MachineResponseModel) HasSessionConnectedViaHostName() bool`

HasSessionConnectedViaHostName returns a boolean if a field has been set.

### SetSessionConnectedViaHostNameNil

`func (o *MachineResponseModel) SetSessionConnectedViaHostNameNil(b bool)`

 SetSessionConnectedViaHostNameNil sets the value for SessionConnectedViaHostName to be an explicit nil

### UnsetSessionConnectedViaHostName
`func (o *MachineResponseModel) UnsetSessionConnectedViaHostName()`

UnsetSessionConnectedViaHostName ensures that no value is present for SessionConnectedViaHostName, not even an explicit nil
### GetSessionConnectedViaIP

`func (o *MachineResponseModel) GetSessionConnectedViaIP() string`

GetSessionConnectedViaIP returns the SessionConnectedViaIP field if non-nil, zero value otherwise.

### GetSessionConnectedViaIPOk

`func (o *MachineResponseModel) GetSessionConnectedViaIPOk() (*string, bool)`

GetSessionConnectedViaIPOk returns a tuple with the SessionConnectedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConnectedViaIP

`func (o *MachineResponseModel) SetSessionConnectedViaIP(v string)`

SetSessionConnectedViaIP sets SessionConnectedViaIP field to given value.

### HasSessionConnectedViaIP

`func (o *MachineResponseModel) HasSessionConnectedViaIP() bool`

HasSessionConnectedViaIP returns a boolean if a field has been set.

### SetSessionConnectedViaIPNil

`func (o *MachineResponseModel) SetSessionConnectedViaIPNil(b bool)`

 SetSessionConnectedViaIPNil sets the value for SessionConnectedViaIP to be an explicit nil

### UnsetSessionConnectedViaIP
`func (o *MachineResponseModel) UnsetSessionConnectedViaIP()`

UnsetSessionConnectedViaIP ensures that no value is present for SessionConnectedViaIP, not even an explicit nil
### GetSessionCount

`func (o *MachineResponseModel) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *MachineResponseModel) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *MachineResponseModel) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *MachineResponseModel) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### SetSessionCountNil

`func (o *MachineResponseModel) SetSessionCountNil(b bool)`

 SetSessionCountNil sets the value for SessionCount to be an explicit nil

### UnsetSessionCount
`func (o *MachineResponseModel) UnsetSessionCount()`

UnsetSessionCount ensures that no value is present for SessionCount, not even an explicit nil
### GetSessionLaunchedViaHostName

`func (o *MachineResponseModel) GetSessionLaunchedViaHostName() string`

GetSessionLaunchedViaHostName returns the SessionLaunchedViaHostName field if non-nil, zero value otherwise.

### GetSessionLaunchedViaHostNameOk

`func (o *MachineResponseModel) GetSessionLaunchedViaHostNameOk() (*string, bool)`

GetSessionLaunchedViaHostNameOk returns a tuple with the SessionLaunchedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLaunchedViaHostName

`func (o *MachineResponseModel) SetSessionLaunchedViaHostName(v string)`

SetSessionLaunchedViaHostName sets SessionLaunchedViaHostName field to given value.

### HasSessionLaunchedViaHostName

`func (o *MachineResponseModel) HasSessionLaunchedViaHostName() bool`

HasSessionLaunchedViaHostName returns a boolean if a field has been set.

### SetSessionLaunchedViaHostNameNil

`func (o *MachineResponseModel) SetSessionLaunchedViaHostNameNil(b bool)`

 SetSessionLaunchedViaHostNameNil sets the value for SessionLaunchedViaHostName to be an explicit nil

### UnsetSessionLaunchedViaHostName
`func (o *MachineResponseModel) UnsetSessionLaunchedViaHostName()`

UnsetSessionLaunchedViaHostName ensures that no value is present for SessionLaunchedViaHostName, not even an explicit nil
### GetSessionLaunchedViaIP

`func (o *MachineResponseModel) GetSessionLaunchedViaIP() string`

GetSessionLaunchedViaIP returns the SessionLaunchedViaIP field if non-nil, zero value otherwise.

### GetSessionLaunchedViaIPOk

`func (o *MachineResponseModel) GetSessionLaunchedViaIPOk() (*string, bool)`

GetSessionLaunchedViaIPOk returns a tuple with the SessionLaunchedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLaunchedViaIP

`func (o *MachineResponseModel) SetSessionLaunchedViaIP(v string)`

SetSessionLaunchedViaIP sets SessionLaunchedViaIP field to given value.

### HasSessionLaunchedViaIP

`func (o *MachineResponseModel) HasSessionLaunchedViaIP() bool`

HasSessionLaunchedViaIP returns a boolean if a field has been set.

### SetSessionLaunchedViaIPNil

`func (o *MachineResponseModel) SetSessionLaunchedViaIPNil(b bool)`

 SetSessionLaunchedViaIPNil sets the value for SessionLaunchedViaIP to be an explicit nil

### UnsetSessionLaunchedViaIP
`func (o *MachineResponseModel) UnsetSessionLaunchedViaIP()`

UnsetSessionLaunchedViaIP ensures that no value is present for SessionLaunchedViaIP, not even an explicit nil
### GetSessionProtocol

`func (o *MachineResponseModel) GetSessionProtocol() ProtocolType`

GetSessionProtocol returns the SessionProtocol field if non-nil, zero value otherwise.

### GetSessionProtocolOk

`func (o *MachineResponseModel) GetSessionProtocolOk() (*ProtocolType, bool)`

GetSessionProtocolOk returns a tuple with the SessionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProtocol

`func (o *MachineResponseModel) SetSessionProtocol(v ProtocolType)`

SetSessionProtocol sets SessionProtocol field to given value.

### HasSessionProtocol

`func (o *MachineResponseModel) HasSessionProtocol() bool`

HasSessionProtocol returns a boolean if a field has been set.

### GetSessionSecureIcaActive

`func (o *MachineResponseModel) GetSessionSecureIcaActive() bool`

GetSessionSecureIcaActive returns the SessionSecureIcaActive field if non-nil, zero value otherwise.

### GetSessionSecureIcaActiveOk

`func (o *MachineResponseModel) GetSessionSecureIcaActiveOk() (*bool, bool)`

GetSessionSecureIcaActiveOk returns a tuple with the SessionSecureIcaActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSecureIcaActive

`func (o *MachineResponseModel) SetSessionSecureIcaActive(v bool)`

SetSessionSecureIcaActive sets SessionSecureIcaActive field to given value.

### HasSessionSecureIcaActive

`func (o *MachineResponseModel) HasSessionSecureIcaActive() bool`

HasSessionSecureIcaActive returns a boolean if a field has been set.

### SetSessionSecureIcaActiveNil

`func (o *MachineResponseModel) SetSessionSecureIcaActiveNil(b bool)`

 SetSessionSecureIcaActiveNil sets the value for SessionSecureIcaActive to be an explicit nil

### UnsetSessionSecureIcaActive
`func (o *MachineResponseModel) UnsetSessionSecureIcaActive()`

UnsetSessionSecureIcaActive ensures that no value is present for SessionSecureIcaActive, not even an explicit nil
### GetSessionSmartAccessTags

`func (o *MachineResponseModel) GetSessionSmartAccessTags() []string`

GetSessionSmartAccessTags returns the SessionSmartAccessTags field if non-nil, zero value otherwise.

### GetSessionSmartAccessTagsOk

`func (o *MachineResponseModel) GetSessionSmartAccessTagsOk() (*[]string, bool)`

GetSessionSmartAccessTagsOk returns a tuple with the SessionSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSmartAccessTags

`func (o *MachineResponseModel) SetSessionSmartAccessTags(v []string)`

SetSessionSmartAccessTags sets SessionSmartAccessTags field to given value.

### HasSessionSmartAccessTags

`func (o *MachineResponseModel) HasSessionSmartAccessTags() bool`

HasSessionSmartAccessTags returns a boolean if a field has been set.

### SetSessionSmartAccessTagsNil

`func (o *MachineResponseModel) SetSessionSmartAccessTagsNil(b bool)`

 SetSessionSmartAccessTagsNil sets the value for SessionSmartAccessTags to be an explicit nil

### UnsetSessionSmartAccessTags
`func (o *MachineResponseModel) UnsetSessionSmartAccessTags()`

UnsetSessionSmartAccessTags ensures that no value is present for SessionSmartAccessTags, not even an explicit nil
### GetSessionStartTime

`func (o *MachineResponseModel) GetSessionStartTime() string`

GetSessionStartTime returns the SessionStartTime field if non-nil, zero value otherwise.

### GetSessionStartTimeOk

`func (o *MachineResponseModel) GetSessionStartTimeOk() (*string, bool)`

GetSessionStartTimeOk returns a tuple with the SessionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStartTime

`func (o *MachineResponseModel) SetSessionStartTime(v string)`

SetSessionStartTime sets SessionStartTime field to given value.

### HasSessionStartTime

`func (o *MachineResponseModel) HasSessionStartTime() bool`

HasSessionStartTime returns a boolean if a field has been set.

### SetSessionStartTimeNil

`func (o *MachineResponseModel) SetSessionStartTimeNil(b bool)`

 SetSessionStartTimeNil sets the value for SessionStartTime to be an explicit nil

### UnsetSessionStartTime
`func (o *MachineResponseModel) UnsetSessionStartTime()`

UnsetSessionStartTime ensures that no value is present for SessionStartTime, not even an explicit nil
### GetFormattedSessionStartTime

`func (o *MachineResponseModel) GetFormattedSessionStartTime() string`

GetFormattedSessionStartTime returns the FormattedSessionStartTime field if non-nil, zero value otherwise.

### GetFormattedSessionStartTimeOk

`func (o *MachineResponseModel) GetFormattedSessionStartTimeOk() (*string, bool)`

GetFormattedSessionStartTimeOk returns a tuple with the FormattedSessionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSessionStartTime

`func (o *MachineResponseModel) SetFormattedSessionStartTime(v string)`

SetFormattedSessionStartTime sets FormattedSessionStartTime field to given value.

### HasFormattedSessionStartTime

`func (o *MachineResponseModel) HasFormattedSessionStartTime() bool`

HasFormattedSessionStartTime returns a boolean if a field has been set.

### SetFormattedSessionStartTimeNil

`func (o *MachineResponseModel) SetFormattedSessionStartTimeNil(b bool)`

 SetFormattedSessionStartTimeNil sets the value for FormattedSessionStartTime to be an explicit nil

### UnsetFormattedSessionStartTime
`func (o *MachineResponseModel) UnsetFormattedSessionStartTime()`

UnsetFormattedSessionStartTime ensures that no value is present for FormattedSessionStartTime, not even an explicit nil
### GetSessionState

`func (o *MachineResponseModel) GetSessionState() SessionState`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *MachineResponseModel) GetSessionStateOk() (*SessionState, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *MachineResponseModel) SetSessionState(v SessionState)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *MachineResponseModel) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### GetSessionStateChangeTime

`func (o *MachineResponseModel) GetSessionStateChangeTime() string`

GetSessionStateChangeTime returns the SessionStateChangeTime field if non-nil, zero value otherwise.

### GetSessionStateChangeTimeOk

`func (o *MachineResponseModel) GetSessionStateChangeTimeOk() (*string, bool)`

GetSessionStateChangeTimeOk returns a tuple with the SessionStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStateChangeTime

`func (o *MachineResponseModel) SetSessionStateChangeTime(v string)`

SetSessionStateChangeTime sets SessionStateChangeTime field to given value.

### HasSessionStateChangeTime

`func (o *MachineResponseModel) HasSessionStateChangeTime() bool`

HasSessionStateChangeTime returns a boolean if a field has been set.

### SetSessionStateChangeTimeNil

`func (o *MachineResponseModel) SetSessionStateChangeTimeNil(b bool)`

 SetSessionStateChangeTimeNil sets the value for SessionStateChangeTime to be an explicit nil

### UnsetSessionStateChangeTime
`func (o *MachineResponseModel) UnsetSessionStateChangeTime()`

UnsetSessionStateChangeTime ensures that no value is present for SessionStateChangeTime, not even an explicit nil
### GetFormattedSessionStateChangeTime

`func (o *MachineResponseModel) GetFormattedSessionStateChangeTime() string`

GetFormattedSessionStateChangeTime returns the FormattedSessionStateChangeTime field if non-nil, zero value otherwise.

### GetFormattedSessionStateChangeTimeOk

`func (o *MachineResponseModel) GetFormattedSessionStateChangeTimeOk() (*string, bool)`

GetFormattedSessionStateChangeTimeOk returns a tuple with the FormattedSessionStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSessionStateChangeTime

`func (o *MachineResponseModel) SetFormattedSessionStateChangeTime(v string)`

SetFormattedSessionStateChangeTime sets FormattedSessionStateChangeTime field to given value.

### HasFormattedSessionStateChangeTime

`func (o *MachineResponseModel) HasFormattedSessionStateChangeTime() bool`

HasFormattedSessionStateChangeTime returns a boolean if a field has been set.

### SetFormattedSessionStateChangeTimeNil

`func (o *MachineResponseModel) SetFormattedSessionStateChangeTimeNil(b bool)`

 SetFormattedSessionStateChangeTimeNil sets the value for FormattedSessionStateChangeTime to be an explicit nil

### UnsetFormattedSessionStateChangeTime
`func (o *MachineResponseModel) UnsetFormattedSessionStateChangeTime()`

UnsetFormattedSessionStateChangeTime ensures that no value is present for FormattedSessionStateChangeTime, not even an explicit nil
### GetSessionSupport

`func (o *MachineResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *MachineResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *MachineResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *MachineResponseModel) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### GetSessionUserName

`func (o *MachineResponseModel) GetSessionUserName() string`

GetSessionUserName returns the SessionUserName field if non-nil, zero value otherwise.

### GetSessionUserNameOk

`func (o *MachineResponseModel) GetSessionUserNameOk() (*string, bool)`

GetSessionUserNameOk returns a tuple with the SessionUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUserName

`func (o *MachineResponseModel) SetSessionUserName(v string)`

SetSessionUserName sets SessionUserName field to given value.

### HasSessionUserName

`func (o *MachineResponseModel) HasSessionUserName() bool`

HasSessionUserName returns a boolean if a field has been set.

### SetSessionUserNameNil

`func (o *MachineResponseModel) SetSessionUserNameNil(b bool)`

 SetSessionUserNameNil sets the value for SessionUserName to be an explicit nil

### UnsetSessionUserName
`func (o *MachineResponseModel) UnsetSessionUserName()`

UnsetSessionUserName ensures that no value is present for SessionUserName, not even an explicit nil
### GetSid

`func (o *MachineResponseModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MachineResponseModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MachineResponseModel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *MachineResponseModel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *MachineResponseModel) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *MachineResponseModel) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSummaryState

`func (o *MachineResponseModel) GetSummaryState() SummaryState`

GetSummaryState returns the SummaryState field if non-nil, zero value otherwise.

### GetSummaryStateOk

`func (o *MachineResponseModel) GetSummaryStateOk() (*SummaryState, bool)`

GetSummaryStateOk returns a tuple with the SummaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryState

`func (o *MachineResponseModel) SetSummaryState(v SummaryState)`

SetSummaryState sets SummaryState field to given value.

### HasSummaryState

`func (o *MachineResponseModel) HasSummaryState() bool`

HasSummaryState returns a boolean if a field has been set.

### GetWillShutdownAfterUse

`func (o *MachineResponseModel) GetWillShutdownAfterUse() bool`

GetWillShutdownAfterUse returns the WillShutdownAfterUse field if non-nil, zero value otherwise.

### GetWillShutdownAfterUseOk

`func (o *MachineResponseModel) GetWillShutdownAfterUseOk() (*bool, bool)`

GetWillShutdownAfterUseOk returns a tuple with the WillShutdownAfterUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillShutdownAfterUse

`func (o *MachineResponseModel) SetWillShutdownAfterUse(v bool)`

SetWillShutdownAfterUse sets WillShutdownAfterUse field to given value.

### HasWillShutdownAfterUse

`func (o *MachineResponseModel) HasWillShutdownAfterUse() bool`

HasWillShutdownAfterUse returns a boolean if a field has been set.

### SetWillShutdownAfterUseNil

`func (o *MachineResponseModel) SetWillShutdownAfterUseNil(b bool)`

 SetWillShutdownAfterUseNil sets the value for WillShutdownAfterUse to be an explicit nil

### UnsetWillShutdownAfterUse
`func (o *MachineResponseModel) UnsetWillShutdownAfterUse()`

UnsetWillShutdownAfterUse ensures that no value is present for WillShutdownAfterUse, not even an explicit nil
### GetWindowsConnectionSetting

`func (o *MachineResponseModel) GetWindowsConnectionSetting() WindowsConnectionSetting`

GetWindowsConnectionSetting returns the WindowsConnectionSetting field if non-nil, zero value otherwise.

### GetWindowsConnectionSettingOk

`func (o *MachineResponseModel) GetWindowsConnectionSettingOk() (*WindowsConnectionSetting, bool)`

GetWindowsConnectionSettingOk returns a tuple with the WindowsConnectionSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsConnectionSetting

`func (o *MachineResponseModel) SetWindowsConnectionSetting(v WindowsConnectionSetting)`

SetWindowsConnectionSetting sets WindowsConnectionSetting field to given value.

### HasWindowsConnectionSetting

`func (o *MachineResponseModel) HasWindowsConnectionSetting() bool`

HasWindowsConnectionSetting returns a boolean if a field has been set.

### GetWindowsActivationStatus

`func (o *MachineResponseModel) GetWindowsActivationStatus() string`

GetWindowsActivationStatus returns the WindowsActivationStatus field if non-nil, zero value otherwise.

### GetWindowsActivationStatusOk

`func (o *MachineResponseModel) GetWindowsActivationStatusOk() (*string, bool)`

GetWindowsActivationStatusOk returns a tuple with the WindowsActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsActivationStatus

`func (o *MachineResponseModel) SetWindowsActivationStatus(v string)`

SetWindowsActivationStatus sets WindowsActivationStatus field to given value.

### HasWindowsActivationStatus

`func (o *MachineResponseModel) HasWindowsActivationStatus() bool`

HasWindowsActivationStatus returns a boolean if a field has been set.

### SetWindowsActivationStatusNil

`func (o *MachineResponseModel) SetWindowsActivationStatusNil(b bool)`

 SetWindowsActivationStatusNil sets the value for WindowsActivationStatus to be an explicit nil

### UnsetWindowsActivationStatus
`func (o *MachineResponseModel) UnsetWindowsActivationStatus()`

UnsetWindowsActivationStatus ensures that no value is present for WindowsActivationStatus, not even an explicit nil
### GetWindowsActivationStatusErrorCode

`func (o *MachineResponseModel) GetWindowsActivationStatusErrorCode() string`

GetWindowsActivationStatusErrorCode returns the WindowsActivationStatusErrorCode field if non-nil, zero value otherwise.

### GetWindowsActivationStatusErrorCodeOk

`func (o *MachineResponseModel) GetWindowsActivationStatusErrorCodeOk() (*string, bool)`

GetWindowsActivationStatusErrorCodeOk returns a tuple with the WindowsActivationStatusErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsActivationStatusErrorCode

`func (o *MachineResponseModel) SetWindowsActivationStatusErrorCode(v string)`

SetWindowsActivationStatusErrorCode sets WindowsActivationStatusErrorCode field to given value.

### HasWindowsActivationStatusErrorCode

`func (o *MachineResponseModel) HasWindowsActivationStatusErrorCode() bool`

HasWindowsActivationStatusErrorCode returns a boolean if a field has been set.

### SetWindowsActivationStatusErrorCodeNil

`func (o *MachineResponseModel) SetWindowsActivationStatusErrorCodeNil(b bool)`

 SetWindowsActivationStatusErrorCodeNil sets the value for WindowsActivationStatusErrorCode to be an explicit nil

### UnsetWindowsActivationStatusErrorCode
`func (o *MachineResponseModel) UnsetWindowsActivationStatusErrorCode()`

UnsetWindowsActivationStatusErrorCode ensures that no value is present for WindowsActivationStatusErrorCode, not even an explicit nil
### GetWindowsActivationStatusVirtualMachineError

`func (o *MachineResponseModel) GetWindowsActivationStatusVirtualMachineError() string`

GetWindowsActivationStatusVirtualMachineError returns the WindowsActivationStatusVirtualMachineError field if non-nil, zero value otherwise.

### GetWindowsActivationStatusVirtualMachineErrorOk

`func (o *MachineResponseModel) GetWindowsActivationStatusVirtualMachineErrorOk() (*string, bool)`

GetWindowsActivationStatusVirtualMachineErrorOk returns a tuple with the WindowsActivationStatusVirtualMachineError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsActivationStatusVirtualMachineError

`func (o *MachineResponseModel) SetWindowsActivationStatusVirtualMachineError(v string)`

SetWindowsActivationStatusVirtualMachineError sets WindowsActivationStatusVirtualMachineError field to given value.

### HasWindowsActivationStatusVirtualMachineError

`func (o *MachineResponseModel) HasWindowsActivationStatusVirtualMachineError() bool`

HasWindowsActivationStatusVirtualMachineError returns a boolean if a field has been set.

### SetWindowsActivationStatusVirtualMachineErrorNil

`func (o *MachineResponseModel) SetWindowsActivationStatusVirtualMachineErrorNil(b bool)`

 SetWindowsActivationStatusVirtualMachineErrorNil sets the value for WindowsActivationStatusVirtualMachineError to be an explicit nil

### UnsetWindowsActivationStatusVirtualMachineError
`func (o *MachineResponseModel) UnsetWindowsActivationStatusVirtualMachineError()`

UnsetWindowsActivationStatusVirtualMachineError ensures that no value is present for WindowsActivationStatusVirtualMachineError, not even an explicit nil
### GetWindowsActivationTypeProvisionedVirtualMachine

`func (o *MachineResponseModel) GetWindowsActivationTypeProvisionedVirtualMachine() string`

GetWindowsActivationTypeProvisionedVirtualMachine returns the WindowsActivationTypeProvisionedVirtualMachine field if non-nil, zero value otherwise.

### GetWindowsActivationTypeProvisionedVirtualMachineOk

`func (o *MachineResponseModel) GetWindowsActivationTypeProvisionedVirtualMachineOk() (*string, bool)`

GetWindowsActivationTypeProvisionedVirtualMachineOk returns a tuple with the WindowsActivationTypeProvisionedVirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsActivationTypeProvisionedVirtualMachine

`func (o *MachineResponseModel) SetWindowsActivationTypeProvisionedVirtualMachine(v string)`

SetWindowsActivationTypeProvisionedVirtualMachine sets WindowsActivationTypeProvisionedVirtualMachine field to given value.

### HasWindowsActivationTypeProvisionedVirtualMachine

`func (o *MachineResponseModel) HasWindowsActivationTypeProvisionedVirtualMachine() bool`

HasWindowsActivationTypeProvisionedVirtualMachine returns a boolean if a field has been set.

### SetWindowsActivationTypeProvisionedVirtualMachineNil

`func (o *MachineResponseModel) SetWindowsActivationTypeProvisionedVirtualMachineNil(b bool)`

 SetWindowsActivationTypeProvisionedVirtualMachineNil sets the value for WindowsActivationTypeProvisionedVirtualMachine to be an explicit nil

### UnsetWindowsActivationTypeProvisionedVirtualMachine
`func (o *MachineResponseModel) UnsetWindowsActivationTypeProvisionedVirtualMachine()`

UnsetWindowsActivationTypeProvisionedVirtualMachine ensures that no value is present for WindowsActivationTypeProvisionedVirtualMachine, not even an explicit nil
### GetZone

`func (o *MachineResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *MachineResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *MachineResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.

### HasZone

`func (o *MachineResponseModel) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSupportedPowerActions

`func (o *MachineResponseModel) GetSupportedPowerActions() []SupportedPowerAction`

GetSupportedPowerActions returns the SupportedPowerActions field if non-nil, zero value otherwise.

### GetSupportedPowerActionsOk

`func (o *MachineResponseModel) GetSupportedPowerActionsOk() (*[]SupportedPowerAction, bool)`

GetSupportedPowerActionsOk returns a tuple with the SupportedPowerActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPowerActions

`func (o *MachineResponseModel) SetSupportedPowerActions(v []SupportedPowerAction)`

SetSupportedPowerActions sets SupportedPowerActions field to given value.

### HasSupportedPowerActions

`func (o *MachineResponseModel) HasSupportedPowerActions() bool`

HasSupportedPowerActions returns a boolean if a field has been set.

### SetSupportedPowerActionsNil

`func (o *MachineResponseModel) SetSupportedPowerActionsNil(b bool)`

 SetSupportedPowerActionsNil sets the value for SupportedPowerActions to be an explicit nil

### UnsetSupportedPowerActions
`func (o *MachineResponseModel) UnsetSupportedPowerActions()`

UnsetSupportedPowerActions ensures that no value is present for SupportedPowerActions, not even an explicit nil
### GetFaultState

`func (o *MachineResponseModel) GetFaultState() FaultState`

GetFaultState returns the FaultState field if non-nil, zero value otherwise.

### GetFaultStateOk

`func (o *MachineResponseModel) GetFaultStateOk() (*FaultState, bool)`

GetFaultStateOk returns a tuple with the FaultState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultState

`func (o *MachineResponseModel) SetFaultState(v FaultState)`

SetFaultState sets FaultState field to given value.

### HasFaultState

`func (o *MachineResponseModel) HasFaultState() bool`

HasFaultState returns a boolean if a field has been set.

### GetContainerMetadata

`func (o *MachineResponseModel) GetContainerMetadata() ContainerMetadataModel`

GetContainerMetadata returns the ContainerMetadata field if non-nil, zero value otherwise.

### GetContainerMetadataOk

`func (o *MachineResponseModel) GetContainerMetadataOk() (*ContainerMetadataModel, bool)`

GetContainerMetadataOk returns a tuple with the ContainerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMetadata

`func (o *MachineResponseModel) SetContainerMetadata(v ContainerMetadataModel)`

SetContainerMetadata sets ContainerMetadata field to given value.

### HasContainerMetadata

`func (o *MachineResponseModel) HasContainerMetadata() bool`

HasContainerMetadata returns a boolean if a field has been set.

### GetTags

`func (o *MachineResponseModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MachineResponseModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MachineResponseModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MachineResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MachineResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MachineResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUpgradeType

`func (o *MachineResponseModel) GetUpgradeType() VdaUpgradeType`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *MachineResponseModel) GetUpgradeTypeOk() (*VdaUpgradeType, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *MachineResponseModel) SetUpgradeType(v VdaUpgradeType)`

SetUpgradeType sets UpgradeType field to given value.

### HasUpgradeType

`func (o *MachineResponseModel) HasUpgradeType() bool`

HasUpgradeType returns a boolean if a field has been set.

### GetUpgradeState

`func (o *MachineResponseModel) GetUpgradeState() VdaUpgradeState`

GetUpgradeState returns the UpgradeState field if non-nil, zero value otherwise.

### GetUpgradeStateOk

`func (o *MachineResponseModel) GetUpgradeStateOk() (*VdaUpgradeState, bool)`

GetUpgradeStateOk returns a tuple with the UpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeState

`func (o *MachineResponseModel) SetUpgradeState(v VdaUpgradeState)`

SetUpgradeState sets UpgradeState field to given value.

### HasUpgradeState

`func (o *MachineResponseModel) HasUpgradeState() bool`

HasUpgradeState returns a boolean if a field has been set.

### GetMachineConfigurationOutOfSync

`func (o *MachineResponseModel) GetMachineConfigurationOutOfSync() bool`

GetMachineConfigurationOutOfSync returns the MachineConfigurationOutOfSync field if non-nil, zero value otherwise.

### GetMachineConfigurationOutOfSyncOk

`func (o *MachineResponseModel) GetMachineConfigurationOutOfSyncOk() (*bool, bool)`

GetMachineConfigurationOutOfSyncOk returns a tuple with the MachineConfigurationOutOfSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineConfigurationOutOfSync

`func (o *MachineResponseModel) SetMachineConfigurationOutOfSync(v bool)`

SetMachineConfigurationOutOfSync sets MachineConfigurationOutOfSync field to given value.

### HasMachineConfigurationOutOfSync

`func (o *MachineResponseModel) HasMachineConfigurationOutOfSync() bool`

HasMachineConfigurationOutOfSync returns a boolean if a field has been set.

### SetMachineConfigurationOutOfSyncNil

`func (o *MachineResponseModel) SetMachineConfigurationOutOfSyncNil(b bool)`

 SetMachineConfigurationOutOfSyncNil sets the value for MachineConfigurationOutOfSync to be an explicit nil

### UnsetMachineConfigurationOutOfSync
`func (o *MachineResponseModel) UnsetMachineConfigurationOutOfSync()`

UnsetMachineConfigurationOutOfSync ensures that no value is present for MachineConfigurationOutOfSync, not even an explicit nil
### GetUpgradeDetail

`func (o *MachineResponseModel) GetUpgradeDetail() MachineUpgradeDetail`

GetUpgradeDetail returns the UpgradeDetail field if non-nil, zero value otherwise.

### GetUpgradeDetailOk

`func (o *MachineResponseModel) GetUpgradeDetailOk() (*MachineUpgradeDetail, bool)`

GetUpgradeDetailOk returns a tuple with the UpgradeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDetail

`func (o *MachineResponseModel) SetUpgradeDetail(v MachineUpgradeDetail)`

SetUpgradeDetail sets UpgradeDetail field to given value.

### HasUpgradeDetail

`func (o *MachineResponseModel) HasUpgradeDetail() bool`

HasUpgradeDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


