package models

// Replace "your-username" with your actual GitHub username or the correct import path

// Create Model
var _ MappedNullable = &CreateSTFDeploymentRequestModel{}

type CreateSTFDeploymentRequestModel struct {
	SiteId      NullableInt64  `json:"SiteId,omitempty"`      // The IIS site id of the deployment
	HostBaseUrl NullableString `json:"HostBaseUrl,omitempty"` // Url used to access the StoreFront server group
}

// ToMap implements MappedNullable.
func (o CreateSTFDeploymentRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	if o.HostBaseUrl.IsSet() {
		toSerialize["HostBaseUrl"] = o.HostBaseUrl.Get()
	}
	return toSerialize, nil
}

func (o *CreateSTFDeploymentRequestModel) SetHostBaseUrl(v string) {
	o.HostBaseUrl.Set(&v)
}

func (o *CreateSTFDeploymentRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}

// Get Model
var _ MappedNullable = &GetSTFDeploymentRequestModel{}

type GetSTFDeploymentRequestModel struct {
	SiteId NullableInt64 `json:"SiteId,omitempty"` // The IIS site id of the deployment
}

// ToMap implements MappedNullable.
func (o GetSTFDeploymentRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	return toSerialize, nil
}

func (o *GetSTFDeploymentRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}

// Set Model
var _ MappedNullable = &GetSTFDeploymentRequestModel{}

type SetSTFDeploymentRequestModel struct {
	SiteId      NullableInt64  `json:"SiteId,omitempty"`      // The IIS site id of the deployment
	HostBaseUrl NullableString `json:"HostBaseUrl,omitempty"` // Url used to access the StoreFront server group
}

// ToMap implements MappedNullable.
func (o SetSTFDeploymentRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	if o.HostBaseUrl.IsSet() {
		toSerialize["HostBaseUrl"] = o.HostBaseUrl.Get()
	}
	return toSerialize, nil
}

func (o *SetSTFDeploymentRequestModel) SetHostBaseUrl(v string) {
	o.HostBaseUrl.Set(&v)
}

func (o *SetSTFDeploymentRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}

// Remove Model
var _ MappedNullable = &ClearSTFDeploymentRequestModel{}

type ClearSTFDeploymentRequestModel struct {
	SiteId NullableInt64 `json:"SiteId,omitempty"` // The IIS site id of the deployment
}

// ToMap implements MappedNullable.
func (o ClearSTFDeploymentRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	return toSerialize, nil
}

func (o *ClearSTFDeploymentRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}
