// Copyright © 2024. Citrix Systems, Inc.

package models

var _ MappedNullable = &SetWebReceiverApplicationShortcutsRequestModel{}

type SetWebReceiverApplicationShortcutsRequestModel struct {
	PromptForUntrustedShortcuts NullableBool `json:"PromptForUntrustedShortcuts,omitempty"`
	TrustedUrls                 []string     `json:"TrustedUrls,omitempty"`
	GatewayUrls                 []string     `json:"GatewayUrls,omitempty"`
}

func (o SetWebReceiverApplicationShortcutsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PromptForUntrustedShortcuts.IsSet() {
		toSerialize["PromptForUntrustedShortcuts"] = o.PromptForUntrustedShortcuts.Get()
	}
	if o.TrustedUrls != nil {
		toSerialize["TrustedUrls"] = o.TrustedUrls
	}
	if o.GatewayUrls != nil {
		toSerialize["GatewayUrls"] = o.GatewayUrls
	}

	return toSerialize, nil
}

func (o *SetWebReceiverApplicationShortcutsRequestModel) SetPromptForUntrustedShortcuts(v bool) {
	o.PromptForUntrustedShortcuts.Set(&v)
}

func (o *SetWebReceiverApplicationShortcutsRequestModel) SetTrustedUrls(v []string) {
	o.TrustedUrls = v
}

func (o *SetWebReceiverApplicationShortcutsRequestModel) SetGatewayUrls(v []string) {
	o.GatewayUrls = v
}
