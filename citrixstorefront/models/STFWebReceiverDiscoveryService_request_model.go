package models

var _ MappedNullable = &SetWebReceiverDiscoveryRequestModel{}

type SetWebReceiverDiscoveryRequestModel struct {
	RunDiscoveryAtStart NullableBool `json:"RunDiscoveryAtStart,omitempty"`
}

func (o *SetWebReceiverDiscoveryRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RunDiscoveryAtStart.IsSet() {
		toSerialize["RunDiscoveryAtStart"] = o.RunDiscoveryAtStart.Get()
	}
	return toSerialize, nil
}

func (o *SetWebReceiverDiscoveryRequestModel) SetRunDiscoveryAtStart(v bool) {
	o.RunDiscoveryAtStart.Set(&v)
}
