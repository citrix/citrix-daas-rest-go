// Copyright © 2024. Citrix Systems, Inc.
package models

// Create Model
var _ MappedNullable = &CreateSTFWebReceiverRequestModel{}

type CreateSTFWebReceiverRequestModel struct {
	VirtualPath               NullableString `json:"VirtualPath,omitempty"`               // 	Site virtual path
	SiteId                    NullableInt64  `json:"SiteId,omitempty"`                    //  IIS site id
	StoreService              NullableString `json:"StoreService,omitempty"`              // Store service to use
	ClassicReceiverExperience NullableBool   `json:"ClassicReceiverExperience,omitempty"` // Enable the classic Receiver experience. The classic experience is no longer supported. Defaulting to the unified experience to maintain compatibility with legacy scripting.
	FriendlyName              NullableString `json:"FriendlyName,omitempty"`              // Friendly name for the WebReceiver
	WebUIExperence            NullableString `json:"WebUIExperence,omitempty"`            // Enable the Receiver experience. This parameter allows for selecting between all UIs that are available
}

// ToMap implements MappedNullable.
func (o CreateSTFWebReceiverRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	if o.VirtualPath.IsSet() {
		toSerialize["VirtualPath"] = o.VirtualPath.Get()
	}
	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}
	if o.ClassicReceiverExperience.IsSet() {
		toSerialize["ClassicReceiverExperience"] = o.ClassicReceiverExperience.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["FriendlyName"] = o.FriendlyName.Get()
	}
	if o.WebUIExperence.IsSet() {
		toSerialize["WebUIExperence"] = o.WebUIExperence.Get()
	}
	return toSerialize, nil
}

func (o *CreateSTFWebReceiverRequestModel) SetVirtualPath(v string) {
	o.VirtualPath.Set(&v)
}
func (o *CreateSTFWebReceiverRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}
func (o *CreateSTFWebReceiverRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}
func (o *CreateSTFWebReceiverRequestModel) SetClassicReceiverExperience(v bool) {
	o.ClassicReceiverExperience.Set(&v)
}
func (o *CreateSTFWebReceiverRequestModel) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}
func (o *CreateSTFWebReceiverRequestModel) SetWebUIExperence(v string) {
	o.WebUIExperence.Set(&v)
}

// Get Model
var _ MappedNullable = &GetSTFWebReceiverRequestModel{}

type GetSTFWebReceiverRequestModel struct {
	VirtualPath           NullableString `json:"VirtualPath,omitempty"`           // Virtual path filter
	SiteId                NullableInt64  `json:"SiteId,omitempty"`                // IIS site id filter
	StoreService          NullableString `json:"StoreService,omitempty"`          // Store service filter
	AuthenticationService NullableString `json:"AuthenticationService,omitempty"` // Authentication service filter
}

// ToMap implements MappedNullable.
func (o GetSTFWebReceiverRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	if o.VirtualPath.IsSet() {
		toSerialize["VirtualPath"] = o.VirtualPath.Get()
	}
	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}
	if o.AuthenticationService.IsSet() {
		toSerialize["AuthenticationService"] = o.AuthenticationService.Get()
	}
	return toSerialize, nil
}

func (o *GetSTFWebReceiverRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}
func (o *GetSTFWebReceiverRequestModel) SetVirtualPath(v string) {
	o.VirtualPath.Set(&v)
}
func (o *GetSTFWebReceiverRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}
func (o *GetSTFWebReceiverRequestModel) SetAuthenticationService(v string) {
	o.AuthenticationService.Set(&v)
}

// Set Model
var _ MappedNullable = &GetSTFWebReceiverRequestModel{}

type SetSTFWebReceiverRequestModel struct {
	WebReceiverService        NullableString `json:"WebReceiverService,omitempty"`        // Receiver for Web service
	ClassicReceiverExperience NullableBool   `json:"ClassicReceiverExperience,omitempty"` // Enable or disable the classic or unified Receiver experience. The classic experience is no longer supported. Defaulting to the unified experience to maintain compatibility with legacy scripting.
	SessionStateTimeout       NullableInt    `json:"SessionStateTimeout,omitempty"`       // Set the session state timeout, in minutes.
	DefaultIISSite            NullableBool   `json:"DefaultIISSite,omitempty"`            // Set the Receiver for Web site as the default page in IIS
	PassThru                  NullableBool   `json:"PassThru,omitempty"`                  // Output the updated WebReceiver object
	WebUIExperence            NullableString `json:"WebUIExperence,omitempty"`            // Enable the Receiver experience. This parameter allows for selecting between all UIs that are available
}

// ToMap implements MappedNullable.
func (o SetSTFWebReceiverRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WebReceiverService.IsSet() {
		toSerialize["WebReceiverService"] = o.WebReceiverService.Get()
	}
	if o.ClassicReceiverExperience.IsSet() {
		toSerialize["ClassicReceiverExperience"] = o.ClassicReceiverExperience.Get()
	}
	if o.SessionStateTimeout.IsSet() {
		toSerialize["SessionStateTimeout"] = o.SessionStateTimeout.Get()
	}
	if o.DefaultIISSite.IsSet() {
		toSerialize["DefaultIISSite"] = o.DefaultIISSite.Get()
	}
	if o.PassThru.IsSet() {
		toSerialize["PassThru"] = o.PassThru.Get()
	}

	return toSerialize, nil
}

func (o *SetSTFWebReceiverRequestModel) SetWebReceiverService(v string) {
	o.WebReceiverService.Set(&v)
}
func (o *SetSTFWebReceiverRequestModel) SetClassicReceiverExperience(v bool) {
	o.ClassicReceiverExperience.Set(&v)
}
func (o *SetSTFWebReceiverRequestModel) SetSessionStateTimeout(v int) {
	o.SessionStateTimeout.Set(&v)
}
func (o *SetSTFWebReceiverRequestModel) SetDefaultIISSite(v bool) {
	o.DefaultIISSite.Set(&v)
}
func (o *SetSTFWebReceiverRequestModel) SetPassThru(v bool) {
	o.PassThru.Set(&v)
}

// Delete Model
var _ MappedNullable = &ClearSTFWebReceiverRequestModel{}

type ClearSTFWebReceiverRequestModel struct {
	WebReceiverService NullableString `json:"WebReceiverService,omitempty"` //Receiver for Web service
}

// ToMap implements MappedNullable.
func (o ClearSTFWebReceiverRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WebReceiverService.IsSet() {
		toSerialize["WebReceiverService"] = o.WebReceiverService.Get()
	}
	return toSerialize, nil
}

func (o *ClearSTFWebReceiverRequestModel) SetWebReceiverService(v string) {
	o.WebReceiverService.Set(&v)
}
