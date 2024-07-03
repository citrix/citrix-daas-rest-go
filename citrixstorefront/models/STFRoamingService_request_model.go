// Copyright © 2024. Citrix Systems, Inc.
package models

// Get-STFRoamingGateway
var _ MappedNullable = &STFRoamingServiceRequestModel{}

type STFRoamingServiceRequestModel struct {
	SiteId NullableInt64 `json:"SiteId,omitempty"`
}

// ToMap implements MappedNullable.
func (o STFRoamingServiceRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	return toSerialize, nil
}

func (o *STFRoamingServiceRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}
