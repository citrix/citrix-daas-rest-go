package models

// Get Model
var _ MappedNullable = &GetSTFWebReceiverAuthenticationMethodsRequestModel{}

type GetSTFWebReceiverAuthenticationMethodsRequestModel struct {
	WebReceiverService NullableString `json:"WebReceiverService,omitempty"` // The WebReceiver service.
}

// ToMap implements MappedNullable.
func (o GetSTFWebReceiverAuthenticationMethodsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WebReceiverService.IsSet() {
		toSerialize["WebReceiverService"] = o.WebReceiverService.Get()
	}
	return toSerialize, nil
}

func (o *GetSTFWebReceiverAuthenticationMethodsRequestModel) SetWebReceiverService(v string) {
	o.WebReceiverService.Set(&v)
}

// Update Model
var _ MappedNullable = &UpdateSTFWebReceiverAuthenticationMethodsRequestModel{}

type UpdateSTFWebReceiverAuthenticationMethodsRequestModel struct {
	WebReceiverService    NullableString `json:"WebReceiverService,omitempty"`    // The WebReceiver service.
	AuthenticationMethods []string       `json:"AuthenticationMethods,omitempty"` // Authentication methods to support
	TokenLifeTime         NullableString `json:"TokenLifeTime,omitempty"`         // The lifetime of the authentication token before it expiries.
}

// ToMap implements MappedNullable.
func (o UpdateSTFWebReceiverAuthenticationMethodsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WebReceiverService.IsSet() {
		toSerialize["WebReceiverService"] = o.WebReceiverService.Get()
	}
	if o.AuthenticationMethods != nil {
		toSerialize["AuthenticationMethods"] = o.AuthenticationMethods
	}
	if o.TokenLifeTime.IsSet() {
		toSerialize["TokenLifeTime"] = o.TokenLifeTime.Get()
	}
	return toSerialize, nil
}

func (o *UpdateSTFWebReceiverAuthenticationMethodsRequestModel) SetWebReceiverService(v string) {
	o.WebReceiverService.Set(&v)
}

func (o *UpdateSTFWebReceiverAuthenticationMethodsRequestModel) SetAuthenticationMethods(v []string) {
	o.AuthenticationMethods = v
}

func (o *UpdateSTFWebReceiverAuthenticationMethodsRequestModel) SetTokenLifeTime(v string) {
	o.TokenLifeTime.Set(&v)
}
