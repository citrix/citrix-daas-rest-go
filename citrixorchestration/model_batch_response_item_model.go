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

// checks if the BatchResponseItemModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchResponseItemModel{}

// BatchResponseItemModel Response object for a single item in a batch request.
type BatchResponseItemModel struct {
	// Reference that allows connecting the response item to the corresponding request item.
	Reference string `json:"Reference"`
	// HTTP status code as a result of the request.
	Code int32 `json:"Code"`
	// Response headers.
	Headers []NameValueStringPairModel `json:"Headers,omitempty"`
	// Response body, if one was present.
	Body NullableString `json:"Body,omitempty"`
}

// NewBatchResponseItemModel instantiates a new BatchResponseItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseItemModel(reference string, code int32) *BatchResponseItemModel {
	this := BatchResponseItemModel{}
	this.Reference = reference
	this.Code = code
	return &this
}

// NewBatchResponseItemModelWithDefaults instantiates a new BatchResponseItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseItemModelWithDefaults() *BatchResponseItemModel {
	this := BatchResponseItemModel{}
	return &this
}

// GetReference returns the Reference field value
func (o *BatchResponseItemModel) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *BatchResponseItemModel) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *BatchResponseItemModel) SetReference(v string) {
	o.Reference = v
}

// GetCode returns the Code field value
func (o *BatchResponseItemModel) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *BatchResponseItemModel) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *BatchResponseItemModel) SetCode(v int32) {
	o.Code = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchResponseItemModel) GetHeaders() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchResponseItemModel) GetHeadersOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BatchResponseItemModel) HasHeaders() bool {
	if o != nil && IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []NameValueStringPairModel and assigns it to the Headers field.
func (o *BatchResponseItemModel) SetHeaders(v []NameValueStringPairModel) {
	o.Headers = v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchResponseItemModel) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchResponseItemModel) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *BatchResponseItemModel) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *BatchResponseItemModel) SetBody(v string) {
	o.Body.Set(&v)
}

// SetBodyNil sets the value for Body to be an explicit nil
func (o *BatchResponseItemModel) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *BatchResponseItemModel) UnsetBody() {
	o.Body.Unset()
}

func (o BatchResponseItemModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchResponseItemModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Reference"] = o.Reference
	toSerialize["Code"] = o.Code
	if o.Headers != nil {
		toSerialize["Headers"] = o.Headers
	}
	if o.Body.IsSet() {
		toSerialize["Body"] = o.Body.Get()
	}
	return toSerialize, nil
}

type NullableBatchResponseItemModel struct {
	value *BatchResponseItemModel
	isSet bool
}

func (v NullableBatchResponseItemModel) Get() *BatchResponseItemModel {
	return v.value
}

func (v *NullableBatchResponseItemModel) Set(val *BatchResponseItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseItemModel(val *BatchResponseItemModel) *NullableBatchResponseItemModel {
	return &NullableBatchResponseItemModel{value: val, isSet: true}
}

func (v NullableBatchResponseItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
