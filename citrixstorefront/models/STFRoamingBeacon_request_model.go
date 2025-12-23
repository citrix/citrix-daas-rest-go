// Copyright © 2025. Citrix Systems, Inc.
package models

// Set Model for Internal Beacon

var _ MappedNullable = &SetSTFRoamingInternalBeaconRequestModel{}

type SetSTFRoamingInternalBeaconRequestModel struct {
	Internal string `json:"Internal,omitempty"` // Internal
}

// ToMap implements MappedNullable.

func (o SetSTFRoamingInternalBeaconRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Internal != "" {
		toSerialize["Internal"] = o.Internal
	}

	return toSerialize, nil
}

func (o *SetSTFRoamingInternalBeaconRequestModel) SetInternal(v string) {
	o.Internal = v
}

// Set Model for External Beacon

var _ MappedNullable = &SetSTFRoamingExternalBeaconRequestModel{}

type SetSTFRoamingExternalBeaconRequestModel struct {
	External []string `json:"External,omitempty"` // External
}

// ToMap implements MappedNullable.

func (o SetSTFRoamingExternalBeaconRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.External == nil {
		toSerialize["External"] = o.External
	}

	return toSerialize, nil
}

func (o *SetSTFRoamingExternalBeaconRequestModel) SetExternal(v []string) {
	o.External = v
}
