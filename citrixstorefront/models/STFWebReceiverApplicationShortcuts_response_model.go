// Copyright © 2024. Citrix Systems, Inc.

package models

type GetWebReceiverApplicationShortcutsResponseModel struct {
	PromptForUntrustedShortcuts NullableBool     `json:"PromptForUntrustedShortcuts,omitempty"`
	TrustedUrls                 []NullableString `json:"TrustedUrls,omitempty"`
	GatewayUrls                 []NullableString `json:"GatewayUrls,omitempty"`
}

func (o *GetWebReceiverApplicationShortcutsResponseModel) GetPromptForUntrustedShortcuts() bool {
	return *o.PromptForUntrustedShortcuts.Get()
}

func (o *GetWebReceiverApplicationShortcutsResponseModel) GetTrustedUrls() []string {
	var trustedUrls []string
	for _, trustedUrl := range o.TrustedUrls {
		trustedUrls = append(trustedUrls, *trustedUrl.Get())
	}
	return trustedUrls
}

func (o *GetWebReceiverApplicationShortcutsResponseModel) GetGatewayUrls() []string {
	var gatewayUrls []string
	for _, gatewayUrl := range o.GatewayUrls {
		gatewayUrls = append(gatewayUrls, *gatewayUrl.Get())
	}
	return gatewayUrls
}
