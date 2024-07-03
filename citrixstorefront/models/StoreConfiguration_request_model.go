// Copyright © 2024. Citrix Systems, Inc.
package models

// Get model

var _ MappedNullable = &GetStoreFarmConfigurationRequestModel{}

type GetStoreFarmConfigurationRequestModel struct {
	StoreService NullableString //Configuration of a StoreFront Store Service
}

// ToMap implements MappedNullable

func (o GetStoreFarmConfigurationRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}

	return toSerialize, nil
}

func (o *GetStoreFarmConfigurationRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}

// Set Model

var _ MappedNullable = &SetStoreFarmConfigurationRequestModel{}

type SetStoreFarmConfigurationRequestModel struct {
	StoreService                       NullableString // Store Service to Configure Global Farm Options
	EnableFileTypeAssociation          NullableBool   // Enable File Type Association so that content is seamlessly redirected to users subscribed applications when they open local files of the appropriate types.
	CommunicationTimeout               NullableString // Communication timeout while using Xml service
	ConnectionTimeout                  NullableString // Connection Timeout when using Xml service
	LeasingStatusExpiryFailed          NullableString // Time before XenDesktop 7 is retried and greater farm in failed leasing mode
	LeasingStatusExpiryLeasing         NullableString // Time before XenDesktop 7 is retried and greater farm in leasing mode
	LeasingStatusExpiryPending         NullableString // Time before XenDesktop 7 is retried and greater farm in pending leasing mode
	PooledSockets                      NullableBool   // Use pooled sockets so that StoreFront maintains a pool of sockets.
	ServerCommunicationAttempts        NullableInt    // Number of server connection attempts before failing
	BackgroundHealthCheckPollingPeriod NullableString // Period of time between polling XenApp/XenDesktop Server health
	AdvancedHealthCheck                NullableBool   // Indicates if advanced health check should be performed
	CertRevocationPolicy               NullableString // Certificate Revocation Policy to use when connecting XML services to HTTPS
}

// ToMap implements MappedNullable

func (o SetStoreFarmConfigurationRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}

	if o.EnableFileTypeAssociation.IsSet() {
		toSerialize["EnableFileTypeAssociation"] = o.EnableFileTypeAssociation.Get()
	}

	if o.CommunicationTimeout.IsSet() {
		toSerialize["CommunicationTimeout"] = o.CommunicationTimeout.Get()
	}

	if o.ConnectionTimeout.IsSet() {
		toSerialize["ConnectionTimeout"] = o.ConnectionTimeout.Get()
	}

	if o.LeasingStatusExpiryFailed.IsSet() {
		toSerialize["LeasingStatusExpiryFailed"] = o.LeasingStatusExpiryFailed.Get()
	}

	if o.LeasingStatusExpiryLeasing.IsSet() {
		toSerialize["LeasingStatusExpiryLeasing"] = o.LeasingStatusExpiryLeasing.Get()
	}

	if o.LeasingStatusExpiryPending.IsSet() {
		toSerialize["LeasingStatusExpiryPending"] = o.LeasingStatusExpiryPending.Get()
	}

	if o.PooledSockets.IsSet() {
		toSerialize["PooledSockets"] = o.PooledSockets.Get()
	}

	if o.ServerCommunicationAttempts.IsSet() {
		toSerialize["ServerCommunicationAttempts"] = o.ServerCommunicationAttempts.Get()
	}

	if o.BackgroundHealthCheckPollingPeriod.IsSet() {
		toSerialize["BackgroundHealthCheckPollingPeriod"] = o.BackgroundHealthCheckPollingPeriod.Get()
	}

	if o.AdvancedHealthCheck.IsSet() {
		toSerialize["AdvancedHealthCheck"] = o.AdvancedHealthCheck.Get()
	}

	if o.CertRevocationPolicy.IsSet() {
		toSerialize["CertRevocationPolicy"] = o.CertRevocationPolicy.Get()
	}

	return toSerialize, nil
}

func (o *SetStoreFarmConfigurationRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetEnableFileTypeAssociation(v bool) {
	o.EnableFileTypeAssociation.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetCommunicationTimeout(v string) {
	o.CommunicationTimeout.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetConnectionTimeout(v string) {
	o.ConnectionTimeout.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetLeasingStatusExpiryFailed(v string) {
	o.LeasingStatusExpiryFailed.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetLeasingStatusExpiryLeasing(v string) {
	o.LeasingStatusExpiryLeasing.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetLeasingStatusExpiryPending(v string) {
	o.LeasingStatusExpiryPending.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetPooledSockets(v bool) {
	o.PooledSockets.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetServerCommunicationAttempts(v int) {
	o.ServerCommunicationAttempts.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetBackgroundHealthCheckPollingPeriod(v string) {
	o.BackgroundHealthCheckPollingPeriod.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetAdvancedHealthCheck(v bool) {
	o.AdvancedHealthCheck.Set(&v)
}

func (o *SetStoreFarmConfigurationRequestModel) SetCertRevocationPolicy(v string) {
	o.CertRevocationPolicy.Set(&v)
}
