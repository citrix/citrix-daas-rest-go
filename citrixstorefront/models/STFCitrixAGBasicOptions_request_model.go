// Copyright © 2025. Citrix Systems, Inc.
package models

// checks if the SetSTFCitrixAGBasicOptionsRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetSTFCitrixAGBasicOptionsRequestModel{}

type SetSTFCitrixAGBasicOptionsRequestModel struct {
	AuthenticationService    NullableString `json:"AuthenticationService,omitempty"`    // A .NET class representing the configuration of a StoreFront Authentication service
	CredentialValidationMode NullableString `json:"CredentialValidationMode,omitempty"` // The credential validation mode to use for authentication
	PasswordRequired         NullableBool   `json:"PasswordRequired,omitempty"`         // Password required for authentication
}

// ToMap implements MappedNullable.
func (o *SetSTFCitrixAGBasicOptionsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationService.IsSet() {
		toSerialize["AuthenticationService"] = o.AuthenticationService.Get()
	}
	if o.CredentialValidationMode.IsSet() {
		toSerialize["CredentialValidationMode"] = o.CredentialValidationMode.Get()
	}
	if o.PasswordRequired.IsSet() {
		toSerialize["PasswordRequired"] = o.PasswordRequired.Get()
	}
	return toSerialize, nil
}

func (o *SetSTFCitrixAGBasicOptionsRequestModel) SetAuthenticationService(v string) {
	o.AuthenticationService.Set(&v)
}

func (o *SetSTFCitrixAGBasicOptionsRequestModel) SetCredentialValidationMode(v string) {
	o.CredentialValidationMode.Set(&v)
}

func (o *SetSTFCitrixAGBasicOptionsRequestModel) SetPasswordRequired(v bool) {
	o.PasswordRequired.Set(&v)
}

// Get-STFWebReceiverSiteStyle
var _ MappedNullable = &GetSTFCitrixAGBasicOptionsRequestModel{}

type GetSTFCitrixAGBasicOptionsRequestModel struct {
	AuthenticationService NullableString `json:"AuthenticationService,omitempty"` // A .NET class representing the configuration of a StoreFront Authentication service
}

// ToMap implements MappedNullable
func (o GetSTFCitrixAGBasicOptionsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.AuthenticationService.IsSet() {
		toSerialize["AuthenticationService"] = o.AuthenticationService.Get()
	}
	return toSerialize, nil
}

func (o *GetSTFCitrixAGBasicOptionsRequestModel) SetAuthenticationService(v string) {
	o.AuthenticationService.Set(&v)
}
