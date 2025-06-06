/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the ImageVersionSpecResourcePoolResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageVersionSpecResourcePoolResponseModel{}

// ImageVersionSpecResourcePoolResponseModel struct for ImageVersionSpecResourcePoolResponseModel
type ImageVersionSpecResourcePoolResponseModel struct {
	// Id of the resource.
	Id NullableString `json:"Id,omitempty"`
	// Name of the resource.
	Name NullableString `json:"Name,omitempty"`
	// XenApp & XenDesktop path to the resource on the hypervisor.  An example value is: `XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot` or `XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}`
	XDPath NullableString `json:"XDPath,omitempty"`
	// Full path to the resources within the resource pool, including the hypervisor, relative to the root of the API. Example: `Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources`
	FullRelativePath string                     `json:"FullRelativePath"`
	Hypervisor       HypervisorRefResponseModel `json:"Hypervisor"`
	// Indicates whether the resource pool is the primary resource pool for the image version. Resource pool used to prepare the image is the default primary resource pool.
	IsPrimary *bool `json:"IsPrimary,omitempty"`
	// The properties of the image that are specific to the target hosting infrastructure.
	CustomProperties []NameValueStringPairModel `json:"CustomProperties,omitempty"`
	// The properties of the image that are specific to the target hosting infrastructure in string format.
	CustomPropertiesInString NullableString `json:"CustomPropertiesInString,omitempty"`
	// The errors in this image version specification within this resource pool.
	Errors                             []string                            `json:"Errors,omitempty"`
	ImageVersionSpecResourcePoolStatus *ImageVersionSpecResourcePoolStatus `json:"ImageVersionSpecResourcePoolStatus,omitempty"`
}

// NewImageVersionSpecResourcePoolResponseModel instantiates a new ImageVersionSpecResourcePoolResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageVersionSpecResourcePoolResponseModel(fullRelativePath string, hypervisor HypervisorRefResponseModel) *ImageVersionSpecResourcePoolResponseModel {
	this := ImageVersionSpecResourcePoolResponseModel{}
	this.FullRelativePath = fullRelativePath
	this.Hypervisor = hypervisor
	return &this
}

// NewImageVersionSpecResourcePoolResponseModelWithDefaults instantiates a new ImageVersionSpecResourcePoolResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageVersionSpecResourcePoolResponseModelWithDefaults() *ImageVersionSpecResourcePoolResponseModel {
	this := ImageVersionSpecResourcePoolResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionSpecResourcePoolResponseModel) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionSpecResourcePoolResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ImageVersionSpecResourcePoolResponseModel) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *ImageVersionSpecResourcePoolResponseModel) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ImageVersionSpecResourcePoolResponseModel) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionSpecResourcePoolResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionSpecResourcePoolResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ImageVersionSpecResourcePoolResponseModel) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ImageVersionSpecResourcePoolResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ImageVersionSpecResourcePoolResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetXDPath returns the XDPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionSpecResourcePoolResponseModel) GetXDPath() string {
	if o == nil || IsNil(o.XDPath.Get()) {
		var ret string
		return ret
	}
	return *o.XDPath.Get()
}

// GetXDPathOk returns a tuple with the XDPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionSpecResourcePoolResponseModel) GetXDPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.XDPath.Get(), o.XDPath.IsSet()
}

// HasXDPath returns a boolean if a field has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) HasXDPath() bool {
	if o != nil && o.XDPath.IsSet() {
		return true
	}

	return false
}

// SetXDPath gets a reference to the given NullableString and assigns it to the XDPath field.
func (o *ImageVersionSpecResourcePoolResponseModel) SetXDPath(v string) {
	o.XDPath.Set(&v)
}

// SetXDPathNil sets the value for XDPath to be an explicit nil
func (o *ImageVersionSpecResourcePoolResponseModel) SetXDPathNil() {
	o.XDPath.Set(nil)
}

// UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
func (o *ImageVersionSpecResourcePoolResponseModel) UnsetXDPath() {
	o.XDPath.Unset()
}

// GetFullRelativePath returns the FullRelativePath field value
func (o *ImageVersionSpecResourcePoolResponseModel) GetFullRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullRelativePath
}

// GetFullRelativePathOk returns a tuple with the FullRelativePath field value
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) GetFullRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullRelativePath, true
}

// SetFullRelativePath sets field value
func (o *ImageVersionSpecResourcePoolResponseModel) SetFullRelativePath(v string) {
	o.FullRelativePath = v
}

// GetHypervisor returns the Hypervisor field value
func (o *ImageVersionSpecResourcePoolResponseModel) GetHypervisor() HypervisorRefResponseModel {
	if o == nil {
		var ret HypervisorRefResponseModel
		return ret
	}

	return o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) GetHypervisorOk() (*HypervisorRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hypervisor, true
}

// SetHypervisor sets field value
func (o *ImageVersionSpecResourcePoolResponseModel) SetHypervisor(v HypervisorRefResponseModel) {
	o.Hypervisor = v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *ImageVersionSpecResourcePoolResponseModel) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) HasIsPrimary() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *ImageVersionSpecResourcePoolResponseModel) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionSpecResourcePoolResponseModel) GetCustomProperties() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionSpecResourcePoolResponseModel) GetCustomPropertiesOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) HasCustomProperties() bool {
	if o != nil && IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []NameValueStringPairModel and assigns it to the CustomProperties field.
func (o *ImageVersionSpecResourcePoolResponseModel) SetCustomProperties(v []NameValueStringPairModel) {
	o.CustomProperties = v
}

// GetCustomPropertiesInString returns the CustomPropertiesInString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionSpecResourcePoolResponseModel) GetCustomPropertiesInString() string {
	if o == nil || IsNil(o.CustomPropertiesInString.Get()) {
		var ret string
		return ret
	}
	return *o.CustomPropertiesInString.Get()
}

// GetCustomPropertiesInStringOk returns a tuple with the CustomPropertiesInString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionSpecResourcePoolResponseModel) GetCustomPropertiesInStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomPropertiesInString.Get(), o.CustomPropertiesInString.IsSet()
}

// HasCustomPropertiesInString returns a boolean if a field has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) HasCustomPropertiesInString() bool {
	if o != nil && o.CustomPropertiesInString.IsSet() {
		return true
	}

	return false
}

// SetCustomPropertiesInString gets a reference to the given NullableString and assigns it to the CustomPropertiesInString field.
func (o *ImageVersionSpecResourcePoolResponseModel) SetCustomPropertiesInString(v string) {
	o.CustomPropertiesInString.Set(&v)
}

// SetCustomPropertiesInStringNil sets the value for CustomPropertiesInString to be an explicit nil
func (o *ImageVersionSpecResourcePoolResponseModel) SetCustomPropertiesInStringNil() {
	o.CustomPropertiesInString.Set(nil)
}

// UnsetCustomPropertiesInString ensures that no value is present for CustomPropertiesInString, not even an explicit nil
func (o *ImageVersionSpecResourcePoolResponseModel) UnsetCustomPropertiesInString() {
	o.CustomPropertiesInString.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionSpecResourcePoolResponseModel) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionSpecResourcePoolResponseModel) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) HasErrors() bool {
	if o != nil && IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ImageVersionSpecResourcePoolResponseModel) SetErrors(v []string) {
	o.Errors = v
}

// GetImageVersionSpecResourcePoolStatus returns the ImageVersionSpecResourcePoolStatus field value if set, zero value otherwise.
func (o *ImageVersionSpecResourcePoolResponseModel) GetImageVersionSpecResourcePoolStatus() ImageVersionSpecResourcePoolStatus {
	if o == nil || IsNil(o.ImageVersionSpecResourcePoolStatus) {
		var ret ImageVersionSpecResourcePoolStatus
		return ret
	}
	return *o.ImageVersionSpecResourcePoolStatus
}

// GetImageVersionSpecResourcePoolStatusOk returns a tuple with the ImageVersionSpecResourcePoolStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) GetImageVersionSpecResourcePoolStatusOk() (*ImageVersionSpecResourcePoolStatus, bool) {
	if o == nil || IsNil(o.ImageVersionSpecResourcePoolStatus) {
		return nil, false
	}
	return o.ImageVersionSpecResourcePoolStatus, true
}

// HasImageVersionSpecResourcePoolStatus returns a boolean if a field has been set.
func (o *ImageVersionSpecResourcePoolResponseModel) HasImageVersionSpecResourcePoolStatus() bool {
	if o != nil && !IsNil(o.ImageVersionSpecResourcePoolStatus) {
		return true
	}

	return false
}

// SetImageVersionSpecResourcePoolStatus gets a reference to the given ImageVersionSpecResourcePoolStatus and assigns it to the ImageVersionSpecResourcePoolStatus field.
func (o *ImageVersionSpecResourcePoolResponseModel) SetImageVersionSpecResourcePoolStatus(v ImageVersionSpecResourcePoolStatus) {
	o.ImageVersionSpecResourcePoolStatus = &v
}

func (o ImageVersionSpecResourcePoolResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageVersionSpecResourcePoolResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.XDPath.IsSet() {
		toSerialize["XDPath"] = o.XDPath.Get()
	}
	toSerialize["FullRelativePath"] = o.FullRelativePath
	toSerialize["Hypervisor"] = o.Hypervisor
	if !IsNil(o.IsPrimary) {
		toSerialize["IsPrimary"] = o.IsPrimary
	}
	if o.CustomProperties != nil {
		toSerialize["CustomProperties"] = o.CustomProperties
	}
	if o.CustomPropertiesInString.IsSet() {
		toSerialize["CustomPropertiesInString"] = o.CustomPropertiesInString.Get()
	}
	if o.Errors != nil {
		toSerialize["Errors"] = o.Errors
	}
	if !IsNil(o.ImageVersionSpecResourcePoolStatus) {
		toSerialize["ImageVersionSpecResourcePoolStatus"] = o.ImageVersionSpecResourcePoolStatus
	}
	return toSerialize, nil
}

type NullableImageVersionSpecResourcePoolResponseModel struct {
	value *ImageVersionSpecResourcePoolResponseModel
	isSet bool
}

func (v NullableImageVersionSpecResourcePoolResponseModel) Get() *ImageVersionSpecResourcePoolResponseModel {
	return v.value
}

func (v *NullableImageVersionSpecResourcePoolResponseModel) Set(val *ImageVersionSpecResourcePoolResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableImageVersionSpecResourcePoolResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableImageVersionSpecResourcePoolResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageVersionSpecResourcePoolResponseModel(val *ImageVersionSpecResourcePoolResponseModel) *NullableImageVersionSpecResourcePoolResponseModel {
	return &NullableImageVersionSpecResourcePoolResponseModel{value: val, isSet: true}
}

func (v NullableImageVersionSpecResourcePoolResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageVersionSpecResourcePoolResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
