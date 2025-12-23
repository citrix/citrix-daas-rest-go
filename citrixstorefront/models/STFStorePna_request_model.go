// Copyright © 2025. Citrix Systems, Inc.
package models

// Set STFStorePnaSetRequestModel Model
var _ MappedNullable = &STFStorePnaSetRequestModel{}

type STFStorePnaSetRequestModel struct {
	LogonMethod       NullableString `json:"LogonMethod,omitempty"`
	DefaultPnaService NullableBool   `json:"DefaultPnaService,omitempty"`
}

func (o STFStorePnaSetRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LogonMethod.IsSet() {
		toSerialize["LogonMethod"] = o.LogonMethod.Get()
	}
	if o.DefaultPnaService.IsSet() {
		toSerialize["DefaultPnaService"] = o.DefaultPnaService.Get()
	}
	return toSerialize, nil
}

func (o *STFStorePnaSetRequestModel) SetLogonMethod(v string) {
	o.LogonMethod.Set(&v)
}

func (o *STFStorePnaSetRequestModel) SetDefaultPnaService(v bool) {
	o.DefaultPnaService.Set(&v)
}
