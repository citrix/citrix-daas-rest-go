// Copyright © 2025. Citrix Systems, Inc.
package models

// Update Model
var _ MappedNullable = &UpdateSTFWebReceiverAuthenticationMethodsRequestModel{}

type UpdateSTFWebReceiverAuthenticationMethodsRequestModel struct {
	AuthenticationMethods []string       `json:"AuthenticationMethods,omitempty"` // Authentication methods to support
	TokenLifeTime         NullableString `json:"TokenLifeTime,omitempty"`         // The lifetime of the authentication token before it expiries.
}

// ToMap implements MappedNullable.
func (o UpdateSTFWebReceiverAuthenticationMethodsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationMethods != nil {
		toSerialize["AuthenticationMethods"] = o.AuthenticationMethods
	}
	if o.TokenLifeTime.IsSet() {
		toSerialize["TokenLifeTime"] = o.TokenLifeTime.Get()
	}
	return toSerialize, nil
}

func (o *UpdateSTFWebReceiverAuthenticationMethodsRequestModel) SetAuthenticationMethods(v []string) {
	o.AuthenticationMethods = v
}

func (o *UpdateSTFWebReceiverAuthenticationMethodsRequestModel) SetTokenLifeTime(v string) {
	o.TokenLifeTime.Set(&v)
}
