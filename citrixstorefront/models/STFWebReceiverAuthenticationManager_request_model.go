package models

var _ MappedNullable = &SetWebReceiverAuthenticationManagerRequestModel{}

type SetWebReceiverAuthenticationManagerRequestModel struct {
	LoginFormTimeout NullableInt `json:"LoginFormTimeout,omitempty"`
}

func (o SetWebReceiverAuthenticationManagerRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LoginFormTimeout.IsSet() {
		toSerialize["LoginFormTimeout"] = o.LoginFormTimeout.Get()
	}
	return toSerialize, nil
}

func (o *SetWebReceiverAuthenticationManagerRequestModel) SetLoginFormTimeout(v int) {
	o.LoginFormTimeout.Set(&v)
}
