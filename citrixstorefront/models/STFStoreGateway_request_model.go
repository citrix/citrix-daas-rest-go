package models

// Register Model
var _ MappedNullable = &RegisterSTFStoreGatewayRequestModel{}

type RegisterSTFStoreGatewayRequestModel struct {
	Gateway        NullableString `json:"Gateway,omitempty"`      // TODO: Add a model for this type.
	StoreService   NullableString `json:"StoreService,omitempty"` // TODO: Add a model for this type.
	DefaultGateway NullableBool   `json:"DefaultGateway,omitempty"`
	UseFullVpn     NullableBool   `json:"UseFullVpn,omitempty"`
}

// ToMap implements MappedNullable.
func (o RegisterSTFStoreGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gateway.IsSet() {
		toSerialize["Gateway"] = o.Gateway.Get()
	}
	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}
	if o.DefaultGateway.IsSet() {
		toSerialize["DefaultGateway"] = o.DefaultGateway.Get()
	}
	if o.UseFullVpn.IsSet() {
		toSerialize["UseFullVpn"] = o.UseFullVpn.Get()
	}
	return toSerialize, nil
}

func (o *RegisterSTFStoreGatewayRequestModel) SetGateway(v string) {
	o.Gateway.Set(&v)
}

func (o *RegisterSTFStoreGatewayRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}

func (o *RegisterSTFStoreGatewayRequestModel) SetDefaultGateway(v bool) {
	o.DefaultGateway.Set(&v)
}

func (o *RegisterSTFStoreGatewayRequestModel) SetUseFullVpn(v bool) {
	o.UseFullVpn.Set(&v)
}

// Unregister Model
var _ MappedNullable = &UnregisterSTFStoreGatewayRequestModel{}

type UnregisterSTFStoreGatewayRequestModel struct {
	Gateway      NullableString `json:"Gateway,omitempty"`      // The Gateway object to remove, either a Roaming or Store gateway
	Name         NullableString `json:"Name,omitempty"`         //Name of the Gateway to remove
	AllGateways  NullableBool   `json:"AllGateways,omitempty"`  //Remove all Gateways
	StoreService NullableString `json:"StoreService,omitempty"` // Store service to unregister from
}

// ToMap implements MappedNullable.
func (o UnregisterSTFStoreGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gateway.IsSet() {
		toSerialize["Gateway"] = o.Gateway.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.AllGateways.IsSet() {
		toSerialize["AllGateways"] = o.AllGateways.Get()
	}
	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}
	return toSerialize, nil
}

func (o *UnregisterSTFStoreGatewayRequestModel) SetGateway(v string) {
	o.Gateway.Set(&v)
}

func (o *UnregisterSTFStoreGatewayRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

func (o *UnregisterSTFStoreGatewayRequestModel) SetAllGateways(v bool) {
	o.AllGateways.Set(&v)
}

func (o *UnregisterSTFStoreGatewayRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}
