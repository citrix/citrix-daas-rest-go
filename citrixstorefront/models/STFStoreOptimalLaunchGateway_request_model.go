// Copyright © 2025. Citrix Systems, Inc.

package models

// Register-STFStoreOptimalLaunchGateway

var _ MappedNullable = &RegisterSTFStoreOptimalLaunchGatewayRequestModel{}

type RegisterSTFStoreOptimalLaunchGatewayRequestModel struct {
	FarmName     []string       `json:"FarmName,omitempty"`     // Farms to apply
	ZoneName     []string       `json:"ZoneName,omitempty"`     // Zones to apply
	StoreService NullableString `json:"StoreService,omitempty"` // Store service reference (not set by user)
}

// ToMap implements MappedNullable.
func (o RegisterSTFStoreOptimalLaunchGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if len(o.FarmName) > 0 {
		toSerialize["FarmName"] = o.FarmName
	}
	if len(o.ZoneName) > 0 {
		toSerialize["ZoneName"] = o.ZoneName
	}
	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}
	return toSerialize, nil
}

func (o *RegisterSTFStoreOptimalLaunchGatewayRequestModel) SetFarmName(v []string) {
	o.FarmName = v
}

func (o *RegisterSTFStoreOptimalLaunchGatewayRequestModel) SetZoneName(v []string) {
	o.ZoneName = v
}

func (o *RegisterSTFStoreOptimalLaunchGatewayRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}

// Unregister model for Unregister-STFStoreOptimalLaunchGateway
var _ MappedNullable = &UnregisterSTFStoreOptimalLaunchGatewayRequestModel{}

type UnregisterSTFStoreOptimalLaunchGatewayRequestModel struct {
	GatewayName  NullableString `json:"GatewayName,omitempty"`  // Name of the gateway to remove
	StoreService NullableString `json:"StoreService,omitempty"` // Store service reference
}

func (o UnregisterSTFStoreOptimalLaunchGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayName.IsSet() {
		toSerialize["GatewayName"] = o.GatewayName.Get()
	}
	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}
	return toSerialize, nil
}

func (o *UnregisterSTFStoreOptimalLaunchGatewayRequestModel) SetGatewayName(v string) {
	o.GatewayName.Set(&v)
}

func (o *UnregisterSTFStoreOptimalLaunchGatewayRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}
