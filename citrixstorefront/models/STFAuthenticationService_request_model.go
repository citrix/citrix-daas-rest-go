// Copyright © 2024. Citrix Systems, Inc.
package models

// checks if the AddSTFAuthenticationServiceRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSTFAuthenticationServiceRequestModel{}

type AddSTFAuthenticationServiceRequestModel struct {
	VirtualPath  NullableString `json:"VirtualPath,omitempty"`  // The IIS virtual path to use for the service
	SiteId       NullableInt64  `json:"SiteId,omitempty"`       // The ID of IIS site to configure the Authentication service for
	FriendlyName NullableString `json:"FriendlyName,omitempty"` // The friendly name the service should be known as
}

// ToMap implements MappedNullable.
func (o *AddSTFAuthenticationServiceRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualPath.IsSet() {
		toSerialize["VirtualPath"] = o.VirtualPath.Get()
	}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["FriendlyName"] = o.FriendlyName.Get()
	}
	return toSerialize, nil
}

func (o *AddSTFAuthenticationServiceRequestModel) SetVirtualPath(v string) {
	o.VirtualPath.Set(&v)
}

func (o *AddSTFAuthenticationServiceRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}

func (o *AddSTFAuthenticationServiceRequestModel) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// checks if the GetSTFAuthenticationServiceRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSTFAuthenticationServiceRequestModel{}

type GetSTFAuthenticationServiceRequestModel struct {
	VirtualPath NullableString `json:"VirtualPath,omitempty"` // Virtual path to the Authentication service website. Excluding the parameter will match on any virtual path
	SiteId      NullableInt64  `json:"SiteId,omitempty"`      // IIS site id of the web site in which the Autentication service is hosted
}

// ToMap implements MappedNullable.
func (o *GetSTFAuthenticationServiceRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualPath.IsSet() {
		toSerialize["VirtualPath"] = o.VirtualPath.Get()
	}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	return toSerialize, nil
}

func (o *GetSTFAuthenticationServiceRequestModel) SetVirtualPath(v string) {
	o.VirtualPath.Set(&v)
}

func (o *GetSTFAuthenticationServiceRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}

// Get Authentication Service Protocol
var _ MappedNullable = &GetSTFAuthenticationServiceProtocolRequestModel{}

type GetSTFAuthenticationServiceProtocolRequestModel struct {
	AuthenticationService string `json:"AuthenticationService,omitempty"` // The Authentication service from which to get the protocols installed
}

// ToMap implements MappedNullable.
func (o *GetSTFAuthenticationServiceProtocolRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationService != "" {
		toSerialize["AuthenticationService"] = o.AuthenticationService
	}
	return toSerialize, nil
}

func (o *GetSTFAuthenticationServiceProtocolRequestModel) SetAuthenticationService(v string) {
	o.AuthenticationService = v
}

// checks if the AddUpdateSTFAuthenticationServiceProtocolRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUpdateSTFAuthenticationServiceProtocolRequestModel{}

type AddUpdateSTFAuthenticationServiceProtocolRequestModel struct {
	Name                  []string       //The name of the authentication protocol.
	AuthenticationService NullableString //The Authentication service to add the feature to.
}

// ToMap implements MappedNullable.
func (o *AddUpdateSTFAuthenticationServiceProtocolRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if len(o.Name) > 0 {
		toSerialize["Name"] = o.Name
	}
	if o.AuthenticationService.IsSet() {
		toSerialize["AuthenticationService"] = o.AuthenticationService.Get()
	}
	return toSerialize, nil
}

func (o *AddUpdateSTFAuthenticationServiceProtocolRequestModel) SetProtocolsName(v []string) {
	o.Name = v
}

func (o *AddUpdateSTFAuthenticationServiceProtocolRequestModel) SetAuthenticationService(v string) {
	o.AuthenticationService.Set(&v)
}
