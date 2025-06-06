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

// checks if the LicenseResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseResponseModel{}

// LicenseResponseModel Response object for the site license.
type LicenseResponseModel struct {
	// The product code.
	ProductCode string `json:"ProductCode"`
	// The product edition.
	ProductEdition string       `json:"ProductEdition"`
	LicenseModel   LicenseModel `json:"LicenseModel"`
}

// NewLicenseResponseModel instantiates a new LicenseResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseResponseModel(productCode string, productEdition string, licenseModel LicenseModel) *LicenseResponseModel {
	this := LicenseResponseModel{}
	this.ProductCode = productCode
	this.ProductEdition = productEdition
	this.LicenseModel = licenseModel
	return &this
}

// NewLicenseResponseModelWithDefaults instantiates a new LicenseResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseResponseModelWithDefaults() *LicenseResponseModel {
	this := LicenseResponseModel{}
	return &this
}

// GetProductCode returns the ProductCode field value
func (o *LicenseResponseModel) GetProductCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value
// and a boolean to check if the value has been set.
func (o *LicenseResponseModel) GetProductCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCode, true
}

// SetProductCode sets field value
func (o *LicenseResponseModel) SetProductCode(v string) {
	o.ProductCode = v
}

// GetProductEdition returns the ProductEdition field value
func (o *LicenseResponseModel) GetProductEdition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductEdition
}

// GetProductEditionOk returns a tuple with the ProductEdition field value
// and a boolean to check if the value has been set.
func (o *LicenseResponseModel) GetProductEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductEdition, true
}

// SetProductEdition sets field value
func (o *LicenseResponseModel) SetProductEdition(v string) {
	o.ProductEdition = v
}

// GetLicenseModel returns the LicenseModel field value
func (o *LicenseResponseModel) GetLicenseModel() LicenseModel {
	if o == nil {
		var ret LicenseModel
		return ret
	}

	return o.LicenseModel
}

// GetLicenseModelOk returns a tuple with the LicenseModel field value
// and a boolean to check if the value has been set.
func (o *LicenseResponseModel) GetLicenseModelOk() (*LicenseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseModel, true
}

// SetLicenseModel sets field value
func (o *LicenseResponseModel) SetLicenseModel(v LicenseModel) {
	o.LicenseModel = v
}

func (o LicenseResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ProductCode"] = o.ProductCode
	toSerialize["ProductEdition"] = o.ProductEdition
	toSerialize["LicenseModel"] = o.LicenseModel
	return toSerialize, nil
}

type NullableLicenseResponseModel struct {
	value *LicenseResponseModel
	isSet bool
}

func (v NullableLicenseResponseModel) Get() *LicenseResponseModel {
	return v.value
}

func (v *NullableLicenseResponseModel) Set(val *LicenseResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseResponseModel(val *LicenseResponseModel) *NullableLicenseResponseModel {
	return &NullableLicenseResponseModel{value: val, isSet: true}
}

func (v NullableLicenseResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
