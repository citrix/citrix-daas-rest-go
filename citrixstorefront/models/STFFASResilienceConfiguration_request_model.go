// Copyright © 2025. Citrix Systems, Inc.
package models

var _ MappedNullable = &SetFASResilienceConfigurationRequestModel{}

type SetFASResilienceConfigurationRequestModel struct {
	Enabled                       NullableBool `json:"Enabled,omitempty"`
	NumberOfExceptionsBeforeBreak NullableInt  `json:"NumberOfExceptionsBeforeBreak,omitempty"`
	DurationOfBreakInMinutes      NullableInt  `json:"DurationOfBreakInMinutes,omitempty"`
}

func (o *SetFASResilienceConfigurationRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled.IsSet() {
		toSerialize["Enabled"] = o.Enabled.Get()
	}
	if o.NumberOfExceptionsBeforeBreak.IsSet() {
		toSerialize["NumberOfExceptionsBeforeBreak"] = o.NumberOfExceptionsBeforeBreak.Get()
	}
	if o.DurationOfBreakInMinutes.IsSet() {
		toSerialize["DurationOfBreakInMinutes"] = o.DurationOfBreakInMinutes.Get()
	}
	return toSerialize, nil
}

func (o *SetFASResilienceConfigurationRequestModel) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

func (o *SetFASResilienceConfigurationRequestModel) SetNumberOfExceptionsBeforeBreak(v int) {
	o.NumberOfExceptionsBeforeBreak.Set(&v)
}

func (o *SetFASResilienceConfigurationRequestModel) SetDurationOfBreakInMinutes(v int) {
	o.DurationOfBreakInMinutes.Set(&v)
}
