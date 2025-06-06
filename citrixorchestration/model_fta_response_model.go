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

// checks if the FtaResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FtaResponseModel{}

// FtaResponseModel Response object for a file-type association.
type FtaResponseModel struct {
	// Content-type.  (Mime type)
	ContentType NullableString `json:"ContentType,omitempty"`
	// The extension name for the file type association. For example, \".txt\" or \".doc\".
	ExtensionName string `json:"ExtensionName"`
	// The description of the handler for the file type association.
	Description NullableString `json:"Description,omitempty"`
	// The name of the handler for the file type association (as seen in the Registry). For example, \"TXTFILE\" or \"Word.Document.8\".
	HandlerName string `json:"HandlerName"`
	// The arguments for the open command that the handler should use. For example, \"%1\".
	OpenArguments NullableString `json:"OpenArguments,omitempty"`
	// Indicates if the file type association is enabled; i.e. whether users who receive the application will launch files of this type using the published application.
	Enabled bool `json:"Enabled"`
	// Indicates if the file type association was imported, as opposed to explicitly added.  File type associations which are not imported will be deleted when disabled.
	IsImported bool `json:"IsImported"`
}

// NewFtaResponseModel instantiates a new FtaResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFtaResponseModel(extensionName string, handlerName string, enabled bool, isImported bool) *FtaResponseModel {
	this := FtaResponseModel{}
	this.ExtensionName = extensionName
	this.HandlerName = handlerName
	this.Enabled = enabled
	this.IsImported = isImported
	return &this
}

// NewFtaResponseModelWithDefaults instantiates a new FtaResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFtaResponseModelWithDefaults() *FtaResponseModel {
	this := FtaResponseModel{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FtaResponseModel) GetContentType() string {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FtaResponseModel) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *FtaResponseModel) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *FtaResponseModel) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *FtaResponseModel) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *FtaResponseModel) UnsetContentType() {
	o.ContentType.Unset()
}

// GetExtensionName returns the ExtensionName field value
func (o *FtaResponseModel) GetExtensionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionName
}

// GetExtensionNameOk returns a tuple with the ExtensionName field value
// and a boolean to check if the value has been set.
func (o *FtaResponseModel) GetExtensionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionName, true
}

// SetExtensionName sets field value
func (o *FtaResponseModel) SetExtensionName(v string) {
	o.ExtensionName = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FtaResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FtaResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *FtaResponseModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *FtaResponseModel) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *FtaResponseModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *FtaResponseModel) UnsetDescription() {
	o.Description.Unset()
}

// GetHandlerName returns the HandlerName field value
func (o *FtaResponseModel) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *FtaResponseModel) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *FtaResponseModel) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetOpenArguments returns the OpenArguments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FtaResponseModel) GetOpenArguments() string {
	if o == nil || IsNil(o.OpenArguments.Get()) {
		var ret string
		return ret
	}
	return *o.OpenArguments.Get()
}

// GetOpenArgumentsOk returns a tuple with the OpenArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FtaResponseModel) GetOpenArgumentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenArguments.Get(), o.OpenArguments.IsSet()
}

// HasOpenArguments returns a boolean if a field has been set.
func (o *FtaResponseModel) HasOpenArguments() bool {
	if o != nil && o.OpenArguments.IsSet() {
		return true
	}

	return false
}

// SetOpenArguments gets a reference to the given NullableString and assigns it to the OpenArguments field.
func (o *FtaResponseModel) SetOpenArguments(v string) {
	o.OpenArguments.Set(&v)
}

// SetOpenArgumentsNil sets the value for OpenArguments to be an explicit nil
func (o *FtaResponseModel) SetOpenArgumentsNil() {
	o.OpenArguments.Set(nil)
}

// UnsetOpenArguments ensures that no value is present for OpenArguments, not even an explicit nil
func (o *FtaResponseModel) UnsetOpenArguments() {
	o.OpenArguments.Unset()
}

// GetEnabled returns the Enabled field value
func (o *FtaResponseModel) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FtaResponseModel) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FtaResponseModel) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIsImported returns the IsImported field value
func (o *FtaResponseModel) GetIsImported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsImported
}

// GetIsImportedOk returns a tuple with the IsImported field value
// and a boolean to check if the value has been set.
func (o *FtaResponseModel) GetIsImportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsImported, true
}

// SetIsImported sets field value
func (o *FtaResponseModel) SetIsImported(v bool) {
	o.IsImported = v
}

func (o FtaResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FtaResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentType.IsSet() {
		toSerialize["ContentType"] = o.ContentType.Get()
	}
	toSerialize["ExtensionName"] = o.ExtensionName
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	toSerialize["HandlerName"] = o.HandlerName
	if o.OpenArguments.IsSet() {
		toSerialize["OpenArguments"] = o.OpenArguments.Get()
	}
	toSerialize["Enabled"] = o.Enabled
	toSerialize["IsImported"] = o.IsImported
	return toSerialize, nil
}

type NullableFtaResponseModel struct {
	value *FtaResponseModel
	isSet bool
}

func (v NullableFtaResponseModel) Get() *FtaResponseModel {
	return v.value
}

func (v *NullableFtaResponseModel) Set(val *FtaResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFtaResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFtaResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFtaResponseModel(val *FtaResponseModel) *NullableFtaResponseModel {
	return &NullableFtaResponseModel{value: val, isSet: true}
}

func (v NullableFtaResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFtaResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
