package models

var _ MappedNullable = &SetWebReceiverCommunicationRequestModel{}

type SetWebReceiverCommunicationRequestModel struct {
	Attempts              NullableInt    `json:"Attempts,omitempty"`
	Timeout               NullableString `json:"Timeout,omitempty"`
	Loopback              NullableString `json:"Loopback,omitempty"`
	LoopbackPortUsingHttp NullableInt    `json:"LoopbackPortUsingHttp,omitempty"`
	ProxyEnabled          NullableBool   `json:"ProxyEnabled,omitempty"`
	ProxyPort             NullableInt    `json:"ProxyPort,omitempty"`
	ProxyProcessName      NullableString `json:"ProxyProcessName,omitempty"`
}

func (o *SetWebReceiverCommunicationRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attempts.IsSet() {
		toSerialize["Attempts"] = o.Attempts.Get()
	}
	if o.Timeout.IsSet() {
		toSerialize["Timeout"] = o.Timeout.Get()
	}
	if o.Loopback.IsSet() {
		toSerialize["Loopback"] = o.Loopback.Get()
	}
	if o.LoopbackPortUsingHttp.IsSet() {
		toSerialize["LoopbackPortUsingHttp"] = o.LoopbackPortUsingHttp.Get()
	}
	if o.ProxyEnabled.IsSet() {
		toSerialize["ProxyEnabled"] = o.ProxyEnabled.Get()
	}
	if o.ProxyPort.IsSet() {
		toSerialize["ProxyPort"] = o.ProxyPort.Get()
	}
	if o.ProxyProcessName.IsSet() {
		toSerialize["ProxyProcessName"] = o.ProxyProcessName.Get()
	}
	return toSerialize, nil
}

func (o *SetWebReceiverCommunicationRequestModel) SetAttempts(v int) {
	o.Attempts.Set(&v)
}

func (o *SetWebReceiverCommunicationRequestModel) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

func (o *SetWebReceiverCommunicationRequestModel) SetLoopback(v string) {
	o.Loopback.Set(&v)
}

func (o *SetWebReceiverCommunicationRequestModel) SetLoopbackPortUsingHttp(v int) {
	o.LoopbackPortUsingHttp.Set(&v)
}

func (o *SetWebReceiverCommunicationRequestModel) SetProxyEnabled(v bool) {
	o.ProxyEnabled.Set(&v)
}

func (o *SetWebReceiverCommunicationRequestModel) SetProxyPort(v int) {
	o.ProxyPort.Set(&v)
}

func (o *SetWebReceiverCommunicationRequestModel) SetProxyProcessName(v string) {
	o.ProxyProcessName.Set(&v)
}
