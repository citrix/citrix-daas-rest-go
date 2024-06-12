// Copyright © 2024. Citrix Systems, Inc.

package models

type STFUserFarmMappingRawResponseModel struct {
	Name         NullableString               `json:"Name,omitempty"`
	GroupMembers map[string]string            `json:"GroupMembers,omitempty"`
	FarmSets     []STFFarmSetRawResponseModel `json:"FarmSets,omitempty"`
}

type STFUserFarmMappingResponseModel struct {
	Name         NullableString                `json:"Name,omitempty"`
	VirtualPath  NullableString                `json:"VirtualPath,omitempty"`
	GroupMembers []STFGroupMemberResponseModel `json:"GroupMembers,omitempty"`
	FarmSets     []STFFarmSetResponseModel     `json:"FarmSets,omitempty"`
}

type STFFarmSetRawResponseModel struct {
	Name                 NullableString `json:"Name,omitempty"`
	LoadBalanceMode      NullableInt    `json:"LoadBalanceMode,omitempty"` // Failover = 0, LoadBalanced = 1
	FarmsAreIdentical    NullableBool   `json:"FarmsAreIdentical,omitempty"`
	AggregationGroupName NullableString `json:"AggregationGroupName,omitempty"`
	PrimaryFarms         NullableString `json:"PrimaryFarms,omitempty"`
	BackupFarms          NullableString `json:"BackupFarms,omitempty"`
}

type STFFarmSetResponseModel struct {
	Name                 NullableString `json:"Name,omitempty"`
	LoadBalanceMode      NullableString `json:"LoadBalanceMode,omitempty"` // Failover = 0, LoadBalanced = 1
	FarmsAreIdentical    NullableBool   `json:"FarmsAreIdentical,omitempty"`
	AggregationGroupName NullableString `json:"AggregationGroupName,omitempty"`
	PrimaryFarms         []string       `json:"PrimaryFarms,omitempty"`
	BackupFarms          []string       `json:"BackupFarms,omitempty"`
}

func (o *STFFarmSetResponseModel) SetName(v string) {
	o.Name.Set(&v)
}

func (o *STFFarmSetResponseModel) SetLoadBalanceMode(v string) {
	o.LoadBalanceMode.Set(&v)
}

func (o *STFFarmSetResponseModel) SetFarmsAreIdentical(v bool) {
	o.FarmsAreIdentical.Set(&v)
}

func (o *STFFarmSetResponseModel) SetAggregationGroupName(v string) {
	o.AggregationGroupName.Set(&v)
}

func (o *STFFarmSetResponseModel) SetPrimaryFarms(v []string) {
	o.PrimaryFarms = v
}

func (o *STFFarmSetResponseModel) SetBackupFarms(v []string) {
	o.BackupFarms = v
}

type STFGroupMemberResponseModel struct {
	GroupName  NullableString `json:"GroupName,omitempty"`
	AccountSid NullableString `json:"AccountSid,omitempty"`
}

func (o *STFGroupMemberResponseModel) SetGroupName(v string) {
	o.GroupName.Set(&v)
}

func (o *STFGroupMemberResponseModel) SetAccountSid(v string) {
	o.AccountSid.Set(&v)
}
