package models

var _ MappedNullable = &SetWebReceiverStrictTransportSecurityRequestModel{}

type SetWebReceiverStrictTransportSecurityRequestModel struct {
	Enabled        NullableBool   `json:"Enabled,omitempty"`
	PolicyDuration NullableString `json:"PolicyDuration,omitempty"`
}

func (o SetWebReceiverStrictTransportSecurityRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled.IsSet() {
		toSerialize["Enabled"] = o.Enabled.Get()
	}
	if o.PolicyDuration.IsSet() {
		toSerialize["PolicyDuration"] = o.PolicyDuration.Get()
	}
	return toSerialize, nil
}

func (o *SetWebReceiverStrictTransportSecurityRequestModel) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

func (o *SetWebReceiverStrictTransportSecurityRequestModel) SetPolicyDuration(v string) {
	o.PolicyDuration.Set(&v)
}
