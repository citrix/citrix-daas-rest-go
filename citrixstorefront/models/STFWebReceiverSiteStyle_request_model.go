// Copyright © 2025. Citrix Systems, Inc.
package models

// Get-STFWebReceiverSiteStyle
var _ MappedNullable = &GetSTFWebReceiverSiteStyleRequestModel{}

type GetSTFWebReceiverSiteStyleRequestModel struct {
	WebReceiverService NullableString `json:"WebReceiverService,omitempty"` // A .NET class representing the configuration of a StoreFront Web Receiver service
}

// ToMap implements MappedNullable
func (o GetSTFWebReceiverSiteStyleRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.WebReceiverService.IsSet() {
		toSerialize["WebReceiverService"] = o.WebReceiverService.Get()
	}

	return toSerialize, nil

}

func (o *GetSTFWebReceiverSiteStyleRequestModel) SetWebReceiverService(v string) {
	o.WebReceiverService.Set(&v)
}

// Set-STFWebReceiverSiteStyle
var _ MappedNullable = &SetSTFWebReceiverSiteStyleRequestModel{}

type SetSTFWebReceiverSiteStyleRequestModel struct {
	WebReceiverService     NullableString `json:"WebReceiverService,omitempty"`     // A .NET class representing the configuration of a StoreFront Web Receiver service
	HeaderLogoPath         NullableString `json:"HeaderLogoPath,omitempty"`         // Header Logo Path
	LogonLogoPath          NullableString `json:"LogonLogoPath,omitempty"`          // Logon Logo Path
	HeaderBackgroundColor  NullableString `json:"HeaderBackgroundColor,omitempty"`  // Background Color of the Header
	HeaderForegroundColor  NullableString `json:"HeaderForegroundColor,omitempty"`  // Foreground Color of the Header
	LinkColor              NullableString `json:"LinkColor,omitempty"`              // Link color of the page
	IgnoreNonExistentLogos NullableBool   `json:"IgnoreNonExistentLogos,omitempty"` // Whether to ignore non-existent logo files and continue to set colors
}

// ToMap implements MappedNullable
func (o SetSTFWebReceiverSiteStyleRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.WebReceiverService.IsSet() {
		toSerialize["WebReceiverService"] = o.WebReceiverService.Get()
	}

	if o.HeaderLogoPath.IsSet() {
		toSerialize["HeaderLogoPath"] = o.HeaderLogoPath.Get()
	}

	if o.LogonLogoPath.IsSet() {
		toSerialize["LogonLogoPath"] = o.LogonLogoPath.Get()
	}

	if o.HeaderBackgroundColor.IsSet() {
		toSerialize["HeaderBackgroundColor"] = o.HeaderBackgroundColor.Get()
	}

	if o.HeaderForegroundColor.IsSet() {
		toSerialize["HeaderForegroundColor"] = o.HeaderForegroundColor.Get()
	}

	if o.LinkColor.IsSet() {
		toSerialize["LinkColor"] = o.LinkColor.Get()
	}

	if o.IgnoreNonExistentLogos.IsSet() {
		toSerialize["IgnoreNonExistentLogos"] = o.IgnoreNonExistentLogos.Get()
	}

	return toSerialize, nil
}

func (o *SetSTFWebReceiverSiteStyleRequestModel) SetWebReceiverService(v string) {
	o.WebReceiverService.Set(&v)
}

func (o *SetSTFWebReceiverSiteStyleRequestModel) SetHeaderLogoPath(v string) {
	o.HeaderLogoPath.Set(&v)
}

func (o *SetSTFWebReceiverSiteStyleRequestModel) SetLogonLogoPath(v string) {
	o.LogonLogoPath.Set(&v)
}

func (o *SetSTFWebReceiverSiteStyleRequestModel) SetHeaderBackgroundColor(v string) {
	o.HeaderBackgroundColor.Set(&v)
}

func (o *SetSTFWebReceiverSiteStyleRequestModel) SetHeaderForegroundColor(v string) {
	o.HeaderForegroundColor.Set(&v)
}

func (o *SetSTFWebReceiverSiteStyleRequestModel) SetLinkColor(v string) {
	o.LinkColor.Set(&v)
}

func (o *SetSTFWebReceiverSiteStyleRequestModel) SetIgnoreNonExistentLogos(v bool) {
	o.IgnoreNonExistentLogos.Set(&v)
}

// Clear-STFWebReceiverSiteStyle
var _ MappedNullable = &ClearSTFWebReceiverSiteStyleRequestModel{}

type ClearSTFWebReceiverSiteStyleRequestModel struct {
	WebReceiverService NullableString // A .NET class representing the configuration of a StoreFront Web Receiver service
}

// ToMap implements MappedNullable
func (o ClearSTFWebReceiverSiteStyleRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.WebReceiverService.IsSet() {
		toSerialize["WebReceiverService"] = o.WebReceiverService.Get()
	}

	return toSerialize, nil

}

func (o *ClearSTFWebReceiverSiteStyleRequestModel) SetWebReceiverService(v string) {
	o.WebReceiverService.Set(&v)
}
