package models

//Base Model shared by Add and Set StoreFarm Request Models
type BaseSTFStoreFarmRequestModel struct {
	FarmName                   NullableString `json:"FarmName"`
	Port                       NullableInt64  `json:"Port"`
	SSLRelayPort               NullableInt64  `json:"SSLRelayPort"`
	TransportType              NullableInt64  `json:"TransportType"`
	LoadBalance                NullableBool   `json:"LoadBalance"`
	XMLValidationEnabled       NullableBool   `json:"XMLValidationEnabled"`
	XMLValidationSecret        NullableString `json:"XMLValidationSecret"`
	FarmType                   NullableString `json:"FarmType"`
	Servers                    []string       `json:"Servers"`
	ServiceUrls                []string       `json:"ServiceUrls"`
	AllFailedBypassDuration    NullableInt64  `json:"AllFailedBypassDuration"`
	BypassDuration             NullableInt64  `json:"BypassDuration"`
	TicketTimeToLive           NullableInt64  `json:"TicketTimeToLive"`
	RadeTicketTimeToLive       NullableInt64  `json:"RadeTicketTimeToLive"`
	MaxFailedServersPerRequest NullableInt64  `json:"MaxFailedServersPerRequest"`
	Zones                      []string       `json:"Zones"`
	Product                    NullableString `json:"Product"`
	RestrictPoPs               NullableString `json:"RestrictPoPs"`
	FarmGuid                   NullableString `json:"FarmGuid"`
}

func (o *BaseSTFStoreFarmRequestModel) SetFarmName(v string) {
	o.FarmName.Set(&v)
}
func (o *BaseSTFStoreFarmRequestModel) SetPort(v int64) {
	o.Port.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetSSLRelayPort(v int64) {
	o.SSLRelayPort.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetTransportType(v int64) {
	o.TransportType.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetLoadBalance(v bool) {
	o.LoadBalance.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetXMLValidationEnabled(v bool) {
	o.XMLValidationEnabled.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetXMLValidationSecret(v string) {
	o.XMLValidationSecret.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetFarmType(v string) {
	o.FarmType.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetServers(v []string) {
	o.Servers = v
}

func (o *BaseSTFStoreFarmRequestModel) SetServiceUrls(v []string) {
	o.ServiceUrls = v
}

func (o *BaseSTFStoreFarmRequestModel) SetAllFailedBypassDuration(v int64) {
	o.AllFailedBypassDuration.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetBypassDuration(v int64) {
	o.BypassDuration.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetTicketTimeToLive(v int64) {
	o.TicketTimeToLive.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetRadeTicketTimeToLive(v int64) {
	o.RadeTicketTimeToLive.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetMaxFailedServersPerRequest(v int64) {
	o.MaxFailedServersPerRequest.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetZones(v []string) {
	o.Zones = v
}

func (o *BaseSTFStoreFarmRequestModel) SetProduct(v string) {
	o.Product.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetRestrictPoPs(v string) {
	o.RestrictPoPs.Set(&v)
}

func (o *BaseSTFStoreFarmRequestModel) SetFarmGuid(v string) {
	o.FarmGuid.Set(&v)
}

// Create Model
var _ MappedNullable = &AddSTFStoreFarmRequestModel{}

type AddSTFStoreFarmRequestModel struct {
	BaseSTFStoreFarmRequestModel
}

// ToMap implements MappedNullable.
func (o BaseSTFStoreFarmRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FarmName.IsSet() {
		toSerialize["SiteId"] = o.FarmName.Get()
	}
	if o.Port.IsSet() {
		toSerialize["Port"] = o.Port.Get()
	}

	if o.SSLRelayPort.IsSet() {
		toSerialize["SSLRelayPort"] = o.SSLRelayPort.Get()
	}
	if o.TransportType.IsSet() {
		toSerialize["TransportType"] = o.TransportType.Get()
	}
	if o.LoadBalance.IsSet() {
		toSerialize["LoadBalance"] = o.LoadBalance.Get()
	}

	if o.XMLValidationEnabled.IsSet() {
		toSerialize["XMLValidationEnabled"] = o.XMLValidationEnabled.Get()
	}

	if o.XMLValidationSecret.IsSet() {
		toSerialize["XMLValidationSecret"] = o.XMLValidationSecret.Get()
	}

	if o.FarmType.IsSet() {
		toSerialize["FarmType"] = o.FarmType.Get()
	}

	if len(o.Servers) > 0 {
		toSerialize["Servers"] = o.Servers
	}

	if len(o.ServiceUrls) > 0 {
		toSerialize["ServiceUrls"] = o.ServiceUrls
	}

	if o.AllFailedBypassDuration.IsSet() {
		toSerialize["AllFailedBypassDuration"] = o.AllFailedBypassDuration.Get()
	}

	if o.BypassDuration.IsSet() {
		toSerialize["BypassDuration"] = o.BypassDuration.Get()
	}

	if o.TicketTimeToLive.IsSet() {
		toSerialize["TicketTimeToLive"] = o.TicketTimeToLive.Get()
	}

	if o.RadeTicketTimeToLive.IsSet() {
		toSerialize["RadeTicketTimeToLive"] = o.RadeTicketTimeToLive.Get()
	}

	if o.MaxFailedServersPerRequest.IsSet() {
		toSerialize["MaxFailedServersPerRequest"] = o.MaxFailedServersPerRequest.Get()
	}

	if len(o.Zones) > 0 {
		toSerialize["Zones"] = o.Zones
	}

	if o.Product.IsSet() {
		toSerialize["Product"] = o.Product.Get()
	}

	if o.RestrictPoPs.IsSet() {
		toSerialize["RestrictPoPs"] = o.RestrictPoPs.Get()
	}

	if o.FarmGuid.IsSet() {
		toSerialize["FarmGuid"] = o.FarmGuid.Get()
	}

	return toSerialize, nil
}

// Get Model
var _ MappedNullable = &SetSTFStoreFarmRequestModel{}

type SetSTFStoreFarmRequestModel struct {
	AddSTFStoreFarmRequestModel
}

// Get Model
var _ MappedNullable = &GetSTFStoreFarmRequestModel{}

type GetSTFStoreFarmRequestModel struct {
	FarmName NullableString `json:"FarmName"`
}

// ToMap implements MappedNullable.
func (o GetSTFStoreFarmRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.FarmName.IsSet() {
		toSerialize["FarmName"] = o.FarmName.Get()
	}

	return toSerialize, nil
}

func (o *GetSTFStoreFarmRequestModel) SetFarmName(v string) {
	o.FarmName.Set(&v)
}
