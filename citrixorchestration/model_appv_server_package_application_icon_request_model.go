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

// checks if the AppvServerPackageApplicationIconRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppvServerPackageApplicationIconRequestModel{}

// AppvServerPackageApplicationIconRequestModel struct for AppvServerPackageApplicationIconRequestModel
type AppvServerPackageApplicationIconRequestModel struct {
	// ManagementServeraddress of the App-V server
	Server string `json:"Server"`
	// Id of the App-V package
	PackageId string `json:"PackageId"`
	// Id of the App-V application within the package.
	ApplicationId string `json:"ApplicationId"`
	// Icon format.  Must be: `{mime-type};{width}x{height}x{colordepth}`  where:   * _mime-type_ must be `image/png`.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either `8` or `24`.  Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly.
	IconFormat NullableString `json:"IconFormat,omitempty"`
}

// NewAppvServerPackageApplicationIconRequestModel instantiates a new AppvServerPackageApplicationIconRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppvServerPackageApplicationIconRequestModel(server string, packageId string, applicationId string) *AppvServerPackageApplicationIconRequestModel {
	this := AppvServerPackageApplicationIconRequestModel{}
	this.Server = server
	this.PackageId = packageId
	this.ApplicationId = applicationId
	return &this
}

// NewAppvServerPackageApplicationIconRequestModelWithDefaults instantiates a new AppvServerPackageApplicationIconRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppvServerPackageApplicationIconRequestModelWithDefaults() *AppvServerPackageApplicationIconRequestModel {
	this := AppvServerPackageApplicationIconRequestModel{}
	return &this
}

// GetServer returns the Server field value
func (o *AppvServerPackageApplicationIconRequestModel) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *AppvServerPackageApplicationIconRequestModel) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *AppvServerPackageApplicationIconRequestModel) SetServer(v string) {
	o.Server = v
}

// GetPackageId returns the PackageId field value
func (o *AppvServerPackageApplicationIconRequestModel) GetPackageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value
// and a boolean to check if the value has been set.
func (o *AppvServerPackageApplicationIconRequestModel) GetPackageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageId, true
}

// SetPackageId sets field value
func (o *AppvServerPackageApplicationIconRequestModel) SetPackageId(v string) {
	o.PackageId = v
}

// GetApplicationId returns the ApplicationId field value
func (o *AppvServerPackageApplicationIconRequestModel) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *AppvServerPackageApplicationIconRequestModel) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *AppvServerPackageApplicationIconRequestModel) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetIconFormat returns the IconFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppvServerPackageApplicationIconRequestModel) GetIconFormat() string {
	if o == nil || IsNil(o.IconFormat.Get()) {
		var ret string
		return ret
	}
	return *o.IconFormat.Get()
}

// GetIconFormatOk returns a tuple with the IconFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppvServerPackageApplicationIconRequestModel) GetIconFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconFormat.Get(), o.IconFormat.IsSet()
}

// HasIconFormat returns a boolean if a field has been set.
func (o *AppvServerPackageApplicationIconRequestModel) HasIconFormat() bool {
	if o != nil && o.IconFormat.IsSet() {
		return true
	}

	return false
}

// SetIconFormat gets a reference to the given NullableString and assigns it to the IconFormat field.
func (o *AppvServerPackageApplicationIconRequestModel) SetIconFormat(v string) {
	o.IconFormat.Set(&v)
}

// SetIconFormatNil sets the value for IconFormat to be an explicit nil
func (o *AppvServerPackageApplicationIconRequestModel) SetIconFormatNil() {
	o.IconFormat.Set(nil)
}

// UnsetIconFormat ensures that no value is present for IconFormat, not even an explicit nil
func (o *AppvServerPackageApplicationIconRequestModel) UnsetIconFormat() {
	o.IconFormat.Unset()
}

func (o AppvServerPackageApplicationIconRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppvServerPackageApplicationIconRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Server"] = o.Server
	toSerialize["PackageId"] = o.PackageId
	toSerialize["ApplicationId"] = o.ApplicationId
	if o.IconFormat.IsSet() {
		toSerialize["IconFormat"] = o.IconFormat.Get()
	}
	return toSerialize, nil
}

type NullableAppvServerPackageApplicationIconRequestModel struct {
	value *AppvServerPackageApplicationIconRequestModel
	isSet bool
}

func (v NullableAppvServerPackageApplicationIconRequestModel) Get() *AppvServerPackageApplicationIconRequestModel {
	return v.value
}

func (v *NullableAppvServerPackageApplicationIconRequestModel) Set(val *AppvServerPackageApplicationIconRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAppvServerPackageApplicationIconRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAppvServerPackageApplicationIconRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppvServerPackageApplicationIconRequestModel(val *AppvServerPackageApplicationIconRequestModel) *NullableAppvServerPackageApplicationIconRequestModel {
	return &NullableAppvServerPackageApplicationIconRequestModel{value: val, isSet: true}
}

func (v NullableAppvServerPackageApplicationIconRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppvServerPackageApplicationIconRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
