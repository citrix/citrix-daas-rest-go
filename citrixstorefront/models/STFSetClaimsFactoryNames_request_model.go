// Copyright © 2025. Citrix Systems, Inc.
package models

// Set STFCalimsFactoryNames Request Model
var _ MappedNullable = &SetSTFClaimsFactoryNamesRequestModel{}

type SetSTFClaimsFactoryNamesRequestModel struct {
	ClaimsFactoryName NullableString `json:"ClaimsFactoryName,omitempty"`
}

// ToMap implements MappedNullable.
func (o *SetSTFClaimsFactoryNamesRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClaimsFactoryName.IsSet() {
		toSerialize["ClaimsFactoryName"] = o.ClaimsFactoryName.Get()
	}
	return toSerialize, nil
}

func (o *SetSTFClaimsFactoryNamesRequestModel) SetClaimsFactoryName(v string) {
	o.ClaimsFactoryName.Set(&v)
}
