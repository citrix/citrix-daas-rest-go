// Copyright © 2025. Citrix Systems, Inc.
package models

var _ MappedNullable = &SetWebReceiverBlockingNotificationRequestModel{}

type SetWebReceiverBlockingNotificationRequestModel struct {
	Enabled   NullableBool   `json:"Enabled,omitempty"`
	Title     NullableString `json:"Title,omitempty"`
	Body      NullableString `json:"Body,omitempty"`
	Button    NullableString `json:"Button,omitempty"`
	Frequency NullableString `json:"Frequency,omitempty"`
}

func (o *SetWebReceiverBlockingNotificationRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled.IsSet() {
		toSerialize["Enabled"] = o.Enabled.Get()
	}
	if o.Title.IsSet() {
		toSerialize["Title"] = o.Title.Get()
	}
	if o.Body.IsSet() {
		toSerialize["Body"] = o.Body.Get()
	}
	if o.Button.IsSet() {
		toSerialize["Button"] = o.Button.Get()
	}
	if o.Frequency.IsSet() {
		toSerialize["Frequency"] = o.Frequency.Get()
	}
	return toSerialize, nil
}

func (o *SetWebReceiverBlockingNotificationRequestModel) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
func (o *SetWebReceiverBlockingNotificationRequestModel) SetTitle(v string) {
	o.Title.Set(&v)
}
func (o *SetWebReceiverBlockingNotificationRequestModel) SetBody(v string) {
	o.Body.Set(&v)
}
func (o *SetWebReceiverBlockingNotificationRequestModel) SetButton(v string) {
	o.Button.Set(&v)
}
func (o *SetWebReceiverBlockingNotificationRequestModel) SetFrequency(v string) {
	o.Frequency.Set(&v)
}
