// Copyright © 2025. Citrix Systems, Inc.

package models

var _ MappedNullable = &SetWebReceiverAppProtectionRequestModel{}

type SetWebReceiverAppProtectionRequestModel struct {
	Enabled NullableString `json:"Enabled,omitempty"`
}

func (o *SetWebReceiverAppProtectionRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled.IsSet() {
		toSerialize["Enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}
func (o *SetWebReceiverAppProtectionRequestModel) SetEnabled(v string) {
	o.Enabled.Set(&v)
}
