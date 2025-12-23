// Copyright © 2025. Citrix Systems, Inc.
package models

import (
	"fmt"
)

type StoreFarmConfigurationRawResponseModel struct {
	Name                               NullableString `json:"Name,omitempty"`
	EnableFileTypeAssociation          NullableBool   `json:"EnableFileTypeAssociation,omitempty"`
	PooledSockets                      NullableBool   `json:"PooledSockets,omitempty"`
	ServerCommunicationAttempts        NullableInt    `json:"ServerCommunicationAttempts,omitempty"`
	CommunicationTimeout               TimeModel      `json:"CommunicationTimeout,omitempty"`
	ConnectionTimeout                  TimeModel      `json:"ConnectionTimeout,omitempty"`
	CertRevocationPolicy               NullableString `json:"CertRevocationPolicy,omitempty"`
	MultiFarmAuthenticationMode        NullableString `json:"MultiFarmAuthenticationMode,omitempty"`
	LeasingStatusExpiryLeasing         TimeModel      `json:"LeasingStatusExpiryLeasing,omitempty"`
	LeasingStatusExpiryFailed          TimeModel      `json:"LeasingStatusExpiryFailed,omitempty"`
	LeasingStatusExpiryPending         TimeModel      `json:"LeasingStatusExpiryPending,omitempty"`
	BackgroundHealthCheckPollingPeriod TimeModel      `json:"BackgroundHealthCheckPollingPeriod,omitempty"`
	AdvancedHealthCheck                NullableBool   `json:"AdvancedHealthCheck,omitempty"`
	//Farms                              []StoreFarmModel `json:"Farms,omitempty"`
}

type StoreFarmConfigurationResponseModel struct {
	Name                               NullableString `json:"Name,omitempty"`
	EnableFileTypeAssociation          NullableBool   `json:"EnableFileTypeAssociation,omitempty"`
	PooledSockets                      NullableBool   `json:"PooledSockets,omitempty"`
	ServerCommunicationAttempts        NullableInt    `json:"ServerCommunicationAttempts,omitempty"`
	CommunicationTimeout               string         `json:"CommunicationTimeout,omitempty"`
	ConnectionTimeout                  string         `json:"ConnectionTimeout,omitempty"`
	CertRevocationPolicy               NullableString `json:"CertRevocationPolicy,omitempty"`
	MultiFarmAuthenticationMode        NullableString `json:"MultiFarmAuthenticationMode,omitempty"`
	LeasingStatusExpiryLeasing         string         `json:"LeasingStatusExpiryLeasing,omitempty"`
	LeasingStatusExpiryFailed          string         `json:"LeasingStatusExpiryFailed,omitempty"`
	LeasingStatusExpiryPending         string         `json:"LeasingStatusExpiryPending,omitempty"`
	BackgroundHealthCheckPollingPeriod string         `json:"BackgroundHealthCheckPollingPeriod,omitempty"`
	AdvancedHealthCheck                NullableBool   `json:"AdvancedHealthCheck,omitempty"`
	//Farms                              []StoreFarmModel `json:"Farms,omitempty"`
}

func (rawResponse StoreFarmConfigurationRawResponseModel) ConvertToResponseModel() StoreFarmConfigurationResponseModel {
	response := StoreFarmConfigurationResponseModel{

		Name: rawResponse.Name,

		EnableFileTypeAssociation: rawResponse.EnableFileTypeAssociation,

		PooledSockets: rawResponse.PooledSockets,

		ServerCommunicationAttempts: rawResponse.ServerCommunicationAttempts,

		CommunicationTimeout: fmt.Sprintf("%d.%d:%d:%d", *rawResponse.CommunicationTimeout.Days.Get(), *rawResponse.CommunicationTimeout.Hours.Get(), *rawResponse.CommunicationTimeout.Minutes.Get(), *rawResponse.CommunicationTimeout.Seconds.Get()),
		ConnectionTimeout:    fmt.Sprintf("%d.%d:%d:%d", *rawResponse.ConnectionTimeout.Days.Get(), *rawResponse.ConnectionTimeout.Hours.Get(), *rawResponse.ConnectionTimeout.Minutes.Get(), *rawResponse.ConnectionTimeout.Seconds.Get()),

		CertRevocationPolicy:        rawResponse.CertRevocationPolicy,
		MultiFarmAuthenticationMode: rawResponse.MultiFarmAuthenticationMode,

		LeasingStatusExpiryFailed:  fmt.Sprintf("%d.%d:%d:%d", *rawResponse.LeasingStatusExpiryFailed.Days.Get(), *rawResponse.LeasingStatusExpiryFailed.Hours.Get(), *rawResponse.LeasingStatusExpiryFailed.Minutes.Get(), *rawResponse.LeasingStatusExpiryFailed.Seconds.Get()),
		LeasingStatusExpiryLeasing: fmt.Sprintf("%d.%d:%d:%d", *rawResponse.LeasingStatusExpiryLeasing.Days.Get(), *rawResponse.LeasingStatusExpiryLeasing.Hours.Get(), *rawResponse.LeasingStatusExpiryLeasing.Minutes.Get(), *rawResponse.LeasingStatusExpiryLeasing.Seconds.Get()),
		LeasingStatusExpiryPending: fmt.Sprintf("%d.%d:%d:%d", *rawResponse.LeasingStatusExpiryPending.Days.Get(), *rawResponse.LeasingStatusExpiryPending.Hours.Get(), *rawResponse.LeasingStatusExpiryPending.Minutes.Get(), *rawResponse.LeasingStatusExpiryPending.Seconds.Get()),

		BackgroundHealthCheckPollingPeriod: fmt.Sprintf("%d.%d:%d:%d", *rawResponse.BackgroundHealthCheckPollingPeriod.Days.Get(), *rawResponse.BackgroundHealthCheckPollingPeriod.Hours.Get(), *rawResponse.BackgroundHealthCheckPollingPeriod.Minutes.Get(), *rawResponse.BackgroundHealthCheckPollingPeriod.Seconds.Get()),
		AdvancedHealthCheck:                rawResponse.AdvancedHealthCheck,
		//Farms:                              rawResponse.Farms,
	}
	return response
}
