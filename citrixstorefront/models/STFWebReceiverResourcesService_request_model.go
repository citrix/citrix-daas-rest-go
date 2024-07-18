// Copyright © 2024. Citrix Systems, Inc.
package models

type SetSTFWebReceiverResourcesServiceRequestModel struct {
	PersistentIconCacheEnabled NullableBool `json:"PersistentIconCacheEnabled,omitempty"`
	IcaFileCacheExpiry         NullableInt  `json:"IcaFileCacheExpiry,omitempty"`
	IconSize                   NullableInt  `json:"IconSize,omitempty"`
	ShowDesktopViewer          NullableBool `json:"ShowDesktopViewer,omitempty"`
}

// ToMap implements MappedNullable
func (o *SetSTFWebReceiverResourcesServiceRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.PersistentIconCacheEnabled.IsSet() {
		toSerialize["PersistentIconCacheEnabled"] = o.PersistentIconCacheEnabled.Get()
	}
	if o.IcaFileCacheExpiry.IsSet() {
		toSerialize["IcaFileCacheExpiry"] = o.IcaFileCacheExpiry.Get()
	}
	if o.IconSize.IsSet() {
		toSerialize["IconSize"] = o.IconSize.Get()
	}
	if o.ShowDesktopViewer.IsSet() {
		toSerialize["ShowDesktopViewer"] = o.ShowDesktopViewer.Get()
	}

	return toSerialize, nil
}

func (o *SetSTFWebReceiverResourcesServiceRequestModel) SetPersistentIconCacheEnabled(v bool) {
	o.PersistentIconCacheEnabled.Set(&v)
}

func (o *SetSTFWebReceiverResourcesServiceRequestModel) SetIcaFileCacheExpiry(v int) {
	o.IcaFileCacheExpiry.Set(&v)
}

func (o *SetSTFWebReceiverResourcesServiceRequestModel) SetIconSize(v int) {
	o.IconSize.Set(&v)
}

func (o *SetSTFWebReceiverResourcesServiceRequestModel) SetShowDesktopViewer(v bool) {
	o.ShowDesktopViewer.Set(&v)
}
