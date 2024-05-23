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

// checks if the ImageVersionSpecRefResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageVersionSpecRefResponseModel{}

// ImageVersionSpecRefResponseModel Image version specification reference.
type ImageVersionSpecRefResponseModel struct {
	// The Id of the image version specification.
	Id string `json:"Id"`
	PreparationType PreparationType `json:"PreparationType"`
	ResourcePool HypervisorResourcePoolRefResponseModel `json:"ResourcePool"`
}

// NewImageVersionSpecRefResponseModel instantiates a new ImageVersionSpecRefResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageVersionSpecRefResponseModel(id string, preparationType PreparationType, resourcePool HypervisorResourcePoolRefResponseModel) *ImageVersionSpecRefResponseModel {
	this := ImageVersionSpecRefResponseModel{}
	this.Id = id
	this.PreparationType = preparationType
	this.ResourcePool = resourcePool
	return &this
}

// NewImageVersionSpecRefResponseModelWithDefaults instantiates a new ImageVersionSpecRefResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageVersionSpecRefResponseModelWithDefaults() *ImageVersionSpecRefResponseModel {
	this := ImageVersionSpecRefResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *ImageVersionSpecRefResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecRefResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImageVersionSpecRefResponseModel) SetId(v string) {
	o.Id = v
}

// GetPreparationType returns the PreparationType field value
func (o *ImageVersionSpecRefResponseModel) GetPreparationType() PreparationType {
	if o == nil {
		var ret PreparationType
		return ret
	}

	return o.PreparationType
}

// GetPreparationTypeOk returns a tuple with the PreparationType field value
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecRefResponseModel) GetPreparationTypeOk() (*PreparationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreparationType, true
}

// SetPreparationType sets field value
func (o *ImageVersionSpecRefResponseModel) SetPreparationType(v PreparationType) {
	o.PreparationType = v
}

// GetResourcePool returns the ResourcePool field value
func (o *ImageVersionSpecRefResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel {
	if o == nil {
		var ret HypervisorResourcePoolRefResponseModel
		return ret
	}

	return o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value
// and a boolean to check if the value has been set.
func (o *ImageVersionSpecRefResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePool, true
}

// SetResourcePool sets field value
func (o *ImageVersionSpecRefResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel) {
	o.ResourcePool = v
}

func (o ImageVersionSpecRefResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageVersionSpecRefResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["PreparationType"] = o.PreparationType
	toSerialize["ResourcePool"] = o.ResourcePool
	return toSerialize, nil
}

type NullableImageVersionSpecRefResponseModel struct {
	value *ImageVersionSpecRefResponseModel
	isSet bool
}

func (v NullableImageVersionSpecRefResponseModel) Get() *ImageVersionSpecRefResponseModel {
	return v.value
}

func (v *NullableImageVersionSpecRefResponseModel) Set(val *ImageVersionSpecRefResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableImageVersionSpecRefResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableImageVersionSpecRefResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageVersionSpecRefResponseModel(val *ImageVersionSpecRefResponseModel) *NullableImageVersionSpecRefResponseModel {
	return &NullableImageVersionSpecRefResponseModel{value: val, isSet: true}
}

func (v NullableImageVersionSpecRefResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageVersionSpecRefResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

