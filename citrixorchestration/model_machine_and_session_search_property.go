/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// MachineAndSessionSearchProperty Properties which can be used for advanced machine and session searches.      Unsupported property when searching for machines :       Unsupported properties when searching for sessions :     ,,,,,,     ,,,,,,,     ,,,       When searching for machines, the session properties cause a match if the     machine is currently hosting a session that matches.  When searching for sessions, the machine properties cause a match if the session is hosted on a machine that matches.
type MachineAndSessionSearchProperty string

// List of MachineAndSessionSearchProperty
const (
	MACHINEANDSESSIONSEARCHPROPERTY_AGENT_VERSION              MachineAndSessionSearchProperty = "AgentVersion"
	MACHINEANDSESSIONSEARCHPROPERTY_ALLOCATION_TYPE            MachineAndSessionSearchProperty = "AllocationType"
	MACHINEANDSESSIONSEARCHPROPERTY_APP_STATE                  MachineAndSessionSearchProperty = "AppState"
	MACHINEANDSESSIONSEARCHPROPERTY_APPS_IN_USE                MachineAndSessionSearchProperty = "AppsInUse"
	MACHINEANDSESSIONSEARCHPROPERTY_AZURE_AD_JOINED_MODE       MachineAndSessionSearchProperty = "AzureAdJoinedMode"
	MACHINEANDSESSIONSEARCHPROPERTY_CONTROLLER_DNS_NAME        MachineAndSessionSearchProperty = "ControllerDnsName"
	MACHINEANDSESSIONSEARCHPROPERTY_CLIENT_IP                  MachineAndSessionSearchProperty = "ClientIP"
	MACHINEANDSESSIONSEARCHPROPERTY_CLIENT_NAME                MachineAndSessionSearchProperty = "ClientName"
	MACHINEANDSESSIONSEARCHPROPERTY_CONNECTED_VIA_HOST_NAME    MachineAndSessionSearchProperty = "ConnectedViaHostName"
	MACHINEANDSESSIONSEARCHPROPERTY_CONNECTED_VIA_IP           MachineAndSessionSearchProperty = "ConnectedViaIP"
	MACHINEANDSESSIONSEARCHPROPERTY_HYPERVISOR_CONNECTION      MachineAndSessionSearchProperty = "HypervisorConnection"
	MACHINEANDSESSIONSEARCHPROPERTY_CONNECTION_PROTOCOL        MachineAndSessionSearchProperty = "ConnectionProtocol"
	MACHINEANDSESSIONSEARCHPROPERTY_CURRENT_USER               MachineAndSessionSearchProperty = "CurrentUser"
	MACHINEANDSESSIONSEARCHPROPERTY_DELIVERY_GROUP             MachineAndSessionSearchProperty = "DeliveryGroup"
	MACHINEANDSESSIONSEARCHPROPERTY_FAULT_STATE                MachineAndSessionSearchProperty = "FaultState"
	MACHINEANDSESSIONSEARCHPROPERTY_IS_ASSIGNED                MachineAndSessionSearchProperty = "IsAssigned"
	MACHINEANDSESSIONSEARCHPROPERTY_LAST_CONNECTION_USER       MachineAndSessionSearchProperty = "LastConnectionUser"
	MACHINEANDSESSIONSEARCHPROPERTY_LAST_CONNECTION_TIME       MachineAndSessionSearchProperty = "LastConnectionTime"
	MACHINEANDSESSIONSEARCHPROPERTY_LAST_DEREGISTRATION_REASON MachineAndSessionSearchProperty = "LastDeregistrationReason"
	MACHINEANDSESSIONSEARCHPROPERTY_LAST_DEREGISTRATION_TIME   MachineAndSessionSearchProperty = "LastDeregistrationTime"
	MACHINEANDSESSIONSEARCHPROPERTY_LAUNCHED_VIA_HOST_NAME     MachineAndSessionSearchProperty = "LaunchedViaHostName"
	MACHINEANDSESSIONSEARCHPROPERTY_LAUNCHED_VIA_IP            MachineAndSessionSearchProperty = "LaunchedViaIP"
	MACHINEANDSESSIONSEARCHPROPERTY_PUBLISHED_NAME             MachineAndSessionSearchProperty = "PublishedName"
	MACHINEANDSESSIONSEARCHPROPERTY_LOAD_INDEX                 MachineAndSessionSearchProperty = "LoadIndex"
	MACHINEANDSESSIONSEARCHPROPERTY_START_TIME                 MachineAndSessionSearchProperty = "StartTime"
	MACHINEANDSESSIONSEARCHPROPERTY_MACHINE_CATALOG            MachineAndSessionSearchProperty = "MachineCatalog"
	MACHINEANDSESSIONSEARCHPROPERTY_MACHINE_UNAVAILABLE_REASON MachineAndSessionSearchProperty = "MachineUnavailableReason"
	MACHINEANDSESSIONSEARCHPROPERTY_IN_MAINTENANCE_MODE        MachineAndSessionSearchProperty = "InMaintenanceMode"
	MACHINEANDSESSIONSEARCHPROPERTY_MAINTENANCE_MODE_REASON    MachineAndSessionSearchProperty = "MaintenanceModeReason"
	MACHINEANDSESSIONSEARCHPROPERTY_DRAINING_UNTIL_SHUTDOWN    MachineAndSessionSearchProperty = "DrainingUntilShutdown"
	MACHINEANDSESSIONSEARCHPROPERTY_MACHINE_NAME               MachineAndSessionSearchProperty = "MachineName"
	MACHINEANDSESSIONSEARCHPROPERTY_OS_TYPE                    MachineAndSessionSearchProperty = "OSType"
	MACHINEANDSESSIONSEARCHPROPERTY_OS_VERSION                 MachineAndSessionSearchProperty = "OSVersion"
	MACHINEANDSESSIONSEARCHPROPERTY_IMAGE_OUT_OF_DATE          MachineAndSessionSearchProperty = "ImageOutOfDate"
	MACHINEANDSESSIONSEARCHPROPERTY_POWER_ACTION_PENDING       MachineAndSessionSearchProperty = "PowerActionPending"
	MACHINEANDSESSIONSEARCHPROPERTY_CLIENT_VERSION             MachineAndSessionSearchProperty = "ClientVersion"
	MACHINEANDSESSIONSEARCHPROPERTY_POWER_STATE                MachineAndSessionSearchProperty = "PowerState"
	MACHINEANDSESSIONSEARCHPROPERTY_SUPPORTED_POWER_ACTIONS    MachineAndSessionSearchProperty = "SupportedPowerActions"
	MACHINEANDSESSIONSEARCHPROPERTY_REGISTRATION_STATE         MachineAndSessionSearchProperty = "RegistrationState"
	MACHINEANDSESSIONSEARCHPROPERTY_SECURE_ICA_ACTIVE          MachineAndSessionSearchProperty = "SecureIcaActive"
	MACHINEANDSESSIONSEARCHPROPERTY_HOSTING_SERVER_NAME        MachineAndSessionSearchProperty = "HostingServerName"
	MACHINEANDSESSIONSEARCHPROPERTY_SESSION_COUNT              MachineAndSessionSearchProperty = "SessionCount"
	MACHINEANDSESSIONSEARCHPROPERTY_SESSION_STATE_CHANGE_TIME  MachineAndSessionSearchProperty = "SessionStateChangeTime"
	MACHINEANDSESSIONSEARCHPROPERTY_SESSION_STATE              MachineAndSessionSearchProperty = "SessionState"
	MACHINEANDSESSIONSEARCHPROPERTY_SESSION_SUPPORT            MachineAndSessionSearchProperty = "SessionSupport"
	MACHINEANDSESSIONSEARCHPROPERTY_SMART_ACCESS_FILTERS       MachineAndSessionSearchProperty = "SmartAccessFilters"
	MACHINEANDSESSIONSEARCHPROPERTY_SUMMARY_STATE              MachineAndSessionSearchProperty = "SummaryState"
	MACHINEANDSESSIONSEARCHPROPERTY_TAGS                       MachineAndSessionSearchProperty = "Tags"
	MACHINEANDSESSIONSEARCHPROPERTY_USER_PRINCIPAL_NAME        MachineAndSessionSearchProperty = "UserPrincipalName"
	MACHINEANDSESSIONSEARCHPROPERTY_USER_NAME                  MachineAndSessionSearchProperty = "UserName"
	MACHINEANDSESSIONSEARCHPROPERTY_USER_DISPLAY_NAME          MachineAndSessionSearchProperty = "UserDisplayName"
	MACHINEANDSESSIONSEARCHPROPERTY_HOSTED_MACHINE_NAME        MachineAndSessionSearchProperty = "HostedMachineName"
	MACHINEANDSESSIONSEARCHPROPERTY_WINDOWS_CONNECTION_SETTING MachineAndSessionSearchProperty = "WindowsConnectionSetting"
	MACHINEANDSESSIONSEARCHPROPERTY_FUNCTIONAL_LEVEL           MachineAndSessionSearchProperty = "FunctionalLevel"
	MACHINEANDSESSIONSEARCHPROPERTY_DNS_NAME                   MachineAndSessionSearchProperty = "DnsName"
	MACHINEANDSESSIONSEARCHPROPERTY_UID                        MachineAndSessionSearchProperty = "Uid"
	MACHINEANDSESSIONSEARCHPROPERTY_VDA_UPGRADE                MachineAndSessionSearchProperty = "VdaUpgrade"
	MACHINEANDSESSIONSEARCHPROPERTY_VDA_UPGRADE_STATE          MachineAndSessionSearchProperty = "VdaUpgradeState"
	MACHINEANDSESSIONSEARCHPROPERTY_PROVISIONING_TYPE          MachineAndSessionSearchProperty = "ProvisioningType"
	MACHINEANDSESSIONSEARCHPROPERTY_ZONE_NAME                  MachineAndSessionSearchProperty = "ZoneName"
)

// All allowed values of MachineAndSessionSearchProperty enum
var AllowedMachineAndSessionSearchPropertyEnumValues = []MachineAndSessionSearchProperty{
	"AgentVersion",
	"AllocationType",
	"AppState",
	"AppsInUse",
	"AzureAdJoinedMode",
	"ControllerDnsName",
	"ClientIP",
	"ClientName",
	"ConnectedViaHostName",
	"ConnectedViaIP",
	"HypervisorConnection",
	"ConnectionProtocol",
	"CurrentUser",
	"DeliveryGroup",
	"FaultState",
	"IsAssigned",
	"LastConnectionUser",
	"LastConnectionTime",
	"LastDeregistrationReason",
	"LastDeregistrationTime",
	"LaunchedViaHostName",
	"LaunchedViaIP",
	"PublishedName",
	"LoadIndex",
	"StartTime",
	"MachineCatalog",
	"MachineUnavailableReason",
	"InMaintenanceMode",
	"MaintenanceModeReason",
	"DrainingUntilShutdown",
	"MachineName",
	"OSType",
	"OSVersion",
	"ImageOutOfDate",
	"PowerActionPending",
	"ClientVersion",
	"PowerState",
	"SupportedPowerActions",
	"RegistrationState",
	"SecureIcaActive",
	"HostingServerName",
	"SessionCount",
	"SessionStateChangeTime",
	"SessionState",
	"SessionSupport",
	"SmartAccessFilters",
	"SummaryState",
	"Tags",
	"UserPrincipalName",
	"UserName",
	"UserDisplayName",
	"HostedMachineName",
	"WindowsConnectionSetting",
	"FunctionalLevel",
	"DnsName",
	"Uid",
	"VdaUpgrade",
	"VdaUpgradeState",
	"ProvisioningType",
	"ZoneName",
}

func (v *MachineAndSessionSearchProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = MachineAndSessionSearchProperty(value)
	return nil
}

// NewMachineAndSessionSearchPropertyFromValue returns a pointer to a valid MachineAndSessionSearchProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMachineAndSessionSearchPropertyFromValue(v string) (*MachineAndSessionSearchProperty, error) {
	ev := MachineAndSessionSearchProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MachineAndSessionSearchProperty: valid values are %v", v, AllowedMachineAndSessionSearchPropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MachineAndSessionSearchProperty) IsValid() bool {
	for _, existing := range AllowedMachineAndSessionSearchPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MachineAndSessionSearchProperty value
func (v MachineAndSessionSearchProperty) Ptr() *MachineAndSessionSearchProperty {
	return &v
}

type NullableMachineAndSessionSearchProperty struct {
	value *MachineAndSessionSearchProperty
	isSet bool
}

func (v NullableMachineAndSessionSearchProperty) Get() *MachineAndSessionSearchProperty {
	return v.value
}

func (v *NullableMachineAndSessionSearchProperty) Set(val *MachineAndSessionSearchProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineAndSessionSearchProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineAndSessionSearchProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineAndSessionSearchProperty(val *MachineAndSessionSearchProperty) *NullableMachineAndSessionSearchProperty {
	return &NullableMachineAndSessionSearchProperty{value: val, isSet: true}
}

func (v NullableMachineAndSessionSearchProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineAndSessionSearchProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
