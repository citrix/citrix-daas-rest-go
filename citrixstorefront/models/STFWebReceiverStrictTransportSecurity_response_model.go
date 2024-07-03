package models

import "fmt"

type GetWebReceiverStrictTransportSecurityRawResponseModel struct {
	Enabled        NullableBool `json:"Enabled,omitempty"`
	PolicyDuration TimeModel    `json:"PolicyDuration,omitempty"`
}

func (o GetWebReceiverStrictTransportSecurityRawResponseModel) ConvertToResponseModel() GetWebReceiverStrictTransportSecurityResponseModel {
	response := GetWebReceiverStrictTransportSecurityResponseModel{
		Enabled:        *o.Enabled.Get(),
		PolicyDuration: fmt.Sprintf("%d.%d:%d:%d", *o.PolicyDuration.Days.Get(), *o.PolicyDuration.Hours.Get(), *o.PolicyDuration.Minutes.Get(), *o.PolicyDuration.Seconds.Get()),
	}
	return response

}

type GetWebReceiverStrictTransportSecurityResponseModel struct {
	Enabled        bool   `json:"Enabled,omitempty"`
	PolicyDuration string `json:"PolicyDuration,omitempty"`
}
