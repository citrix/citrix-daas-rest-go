package models

import (
	"fmt"
)

type GetWebReceiverCommunicationRawResponseModel struct {
	Attempts              NullableInt `json:"Attempts,omitempty"`
	Timeout               TimeModel   `json:"Timeout,omitempty"`
	Proxy                 ProxyModel  `json:"Proxy,omitempty"`
	Loopback              NullableInt `json:"Loopback,omitempty"`
	LoopbackPortUsingHttp NullableInt `json:"LoopbackPortUsingHttp,omitempty"`
}

func (o GetWebReceiverCommunicationRawResponseModel) ConvertToResponseModel() GetWebReceiverCommunicationResponseModel {
	response := GetWebReceiverCommunicationResponseModel{
		Attempts:              o.Attempts,
		Timeout:               fmt.Sprintf("%d.%d:%d:%d", *o.Timeout.Days.Get(), *o.Timeout.Hours.Get(), *o.Timeout.Minutes.Get(), *o.Timeout.Seconds.Get()),
		Proxy:                 o.Proxy,
		LoopbackPortUsingHttp: o.LoopbackPortUsingHttp,
	}
	if o.Loopback.IsSet() {
		switch *o.Loopback.Get() {
		case 0:
			response.Loopback = "On"
		case 1:
			response.Loopback = "Off"
		case 2:
			response.Loopback = "OnUsingHttp"
		}
	}
	return response
}

type GetWebReceiverCommunicationResponseModel struct {
	Attempts              NullableInt `json:"Attempts,omitempty"`
	Timeout               string      `json:"Timeout,omitempty"`
	Proxy                 ProxyModel  `json:"Proxy,omitempty"`
	Loopback              string      `json:"Loopback,omitempty"`
	LoopbackPortUsingHttp NullableInt `json:"LoopbackPortUsingHttp,omitempty"`
}

type ProxyModel struct {
	Enabled     NullableBool   `json:"Enabled,omitempty"`
	ProcessName NullableString `json:"ProcessName,omitempty"`
	Port        NullableInt    `json:"Port,omitempty"`
}
