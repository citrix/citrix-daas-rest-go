// Copyright © 2025. Citrix Systems, Inc.
package models

// Set STFUserFarmMapping Model
var _ MappedNullable = &STFEquivalentFarmSetRequestModel{}

type STFEquivalentFarmSetRequestModel struct {
	Name                 NullableString `json:"Name,omitempty"`
	AggregationGroupName NullableString `json:"AggregationGroupName,omitempty"`
	PrimaryFarms         []string       `json:"PrimaryFarms,omitempty"`
	BackupFarms          []string       `json:"BackupFarms,omitempty"`
	LoadBalanceMode      NullableString `json:"LoadBalanceMode,omitempty"`
	FarmsAreIdentical    NullableBool   `json:"FarmsAreIdentical,omitempty"`
}

func (o STFEquivalentFarmSetRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.AggregationGroupName.IsSet() {
		toSerialize["AggregationGroupName"] = o.AggregationGroupName.Get()
	}
	if o.PrimaryFarms != nil {
		toSerialize["PrimaryFarms"] = o.PrimaryFarms
	}
	if o.BackupFarms != nil {
		toSerialize["BackupFarms"] = o.BackupFarms
	}
	if o.LoadBalanceMode.IsSet() {
		toSerialize["LoadBalanceMode"] = o.LoadBalanceMode.Get()
	}
	if o.FarmsAreIdentical.IsSet() {
		toSerialize["FarmsAreIdentical"] = o.FarmsAreIdentical.Get()
	}

	return toSerialize, nil

}

func (o *STFEquivalentFarmSetRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

func (o *STFEquivalentFarmSetRequestModel) SetAggregationGroupName(v string) {
	o.AggregationGroupName.Set(&v)
}

func (o *STFEquivalentFarmSetRequestModel) SetPrimaryFarms(v []string) {
	o.PrimaryFarms = v
}

func (o *STFEquivalentFarmSetRequestModel) SetBackupFarms(v []string) {
	o.BackupFarms = v
}

func (o *STFEquivalentFarmSetRequestModel) SetLoadBalanceMode(v string) {
	o.LoadBalanceMode.Set(&v)
}

func (o *STFEquivalentFarmSetRequestModel) SetFarmsAreIdentical(v bool) {
	o.FarmsAreIdentical.Set(&v)
}

var _ MappedNullable = &STFUserFarmMappingGroup{}

type STFUserFarmMappingGroup struct {
	GroupName  NullableString `json:"GroupName,omitempty"`
	AccountSid NullableString `json:"AccountSid,omitempty"`
}

// ToMap implements MappedNullable.
func (o STFUserFarmMappingGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupName.IsSet() {
		toSerialize["GroupName"] = o.GroupName.Get()
	}
	if o.AccountSid.IsSet() {
		toSerialize["AccountSid"] = o.AccountSid.Get()
	}

	return toSerialize, nil
}

func (o *STFUserFarmMappingGroup) SetGroupName(v string) {
	o.GroupName.Set(&v)
}

func (o *STFUserFarmMappingGroup) SetAccountSid(v string) {
	o.AccountSid.Set(&v)
}
