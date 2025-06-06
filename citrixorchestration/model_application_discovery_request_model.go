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

// checks if the ApplicationDiscoveryRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationDiscoveryRequestModel{}

// ApplicationDiscoveryRequestModel Specify ApplicationGroups, DeliverGroups or MachineCatalogs to find the best machine to browse application; Or specify machine IP/DnsName to browse application from the machine. If both ApplicationGroups, DeliveryGroups, MachineCatalogs and ServerOrPath are specified, ServerOrPath will take priority.
type ApplicationDiscoveryRequestModel struct {
	// The machines associated with Application Groups will be browsed.
	ApplicationGroups []string `json:"ApplicationGroups,omitempty"`
	// The machines associated with Delivery Groups will be browsed.
	DeliveryGroups []string `json:"DeliveryGroups,omitempty"`
	// The machines in the catalogs will be browsed.
	MachineCatalogs []string `json:"MachineCatalogs,omitempty"`
	// Specify machine IP or DNS to get machine shares; Specify local path to get directory/files.
	ServerOrPath NullableString `json:"ServerOrPath,omitempty"`
}

// NewApplicationDiscoveryRequestModel instantiates a new ApplicationDiscoveryRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationDiscoveryRequestModel() *ApplicationDiscoveryRequestModel {
	this := ApplicationDiscoveryRequestModel{}
	return &this
}

// NewApplicationDiscoveryRequestModelWithDefaults instantiates a new ApplicationDiscoveryRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationDiscoveryRequestModelWithDefaults() *ApplicationDiscoveryRequestModel {
	this := ApplicationDiscoveryRequestModel{}
	return &this
}

// GetApplicationGroups returns the ApplicationGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationDiscoveryRequestModel) GetApplicationGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ApplicationGroups
}

// GetApplicationGroupsOk returns a tuple with the ApplicationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationDiscoveryRequestModel) GetApplicationGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.ApplicationGroups) {
		return nil, false
	}
	return o.ApplicationGroups, true
}

// HasApplicationGroups returns a boolean if a field has been set.
func (o *ApplicationDiscoveryRequestModel) HasApplicationGroups() bool {
	if o != nil && IsNil(o.ApplicationGroups) {
		return true
	}

	return false
}

// SetApplicationGroups gets a reference to the given []string and assigns it to the ApplicationGroups field.
func (o *ApplicationDiscoveryRequestModel) SetApplicationGroups(v []string) {
	o.ApplicationGroups = v
}

// GetDeliveryGroups returns the DeliveryGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationDiscoveryRequestModel) GetDeliveryGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DeliveryGroups
}

// GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationDiscoveryRequestModel) GetDeliveryGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.DeliveryGroups) {
		return nil, false
	}
	return o.DeliveryGroups, true
}

// HasDeliveryGroups returns a boolean if a field has been set.
func (o *ApplicationDiscoveryRequestModel) HasDeliveryGroups() bool {
	if o != nil && IsNil(o.DeliveryGroups) {
		return true
	}

	return false
}

// SetDeliveryGroups gets a reference to the given []string and assigns it to the DeliveryGroups field.
func (o *ApplicationDiscoveryRequestModel) SetDeliveryGroups(v []string) {
	o.DeliveryGroups = v
}

// GetMachineCatalogs returns the MachineCatalogs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationDiscoveryRequestModel) GetMachineCatalogs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MachineCatalogs
}

// GetMachineCatalogsOk returns a tuple with the MachineCatalogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationDiscoveryRequestModel) GetMachineCatalogsOk() ([]string, bool) {
	if o == nil || IsNil(o.MachineCatalogs) {
		return nil, false
	}
	return o.MachineCatalogs, true
}

// HasMachineCatalogs returns a boolean if a field has been set.
func (o *ApplicationDiscoveryRequestModel) HasMachineCatalogs() bool {
	if o != nil && IsNil(o.MachineCatalogs) {
		return true
	}

	return false
}

// SetMachineCatalogs gets a reference to the given []string and assigns it to the MachineCatalogs field.
func (o *ApplicationDiscoveryRequestModel) SetMachineCatalogs(v []string) {
	o.MachineCatalogs = v
}

// GetServerOrPath returns the ServerOrPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationDiscoveryRequestModel) GetServerOrPath() string {
	if o == nil || IsNil(o.ServerOrPath.Get()) {
		var ret string
		return ret
	}
	return *o.ServerOrPath.Get()
}

// GetServerOrPathOk returns a tuple with the ServerOrPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationDiscoveryRequestModel) GetServerOrPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerOrPath.Get(), o.ServerOrPath.IsSet()
}

// HasServerOrPath returns a boolean if a field has been set.
func (o *ApplicationDiscoveryRequestModel) HasServerOrPath() bool {
	if o != nil && o.ServerOrPath.IsSet() {
		return true
	}

	return false
}

// SetServerOrPath gets a reference to the given NullableString and assigns it to the ServerOrPath field.
func (o *ApplicationDiscoveryRequestModel) SetServerOrPath(v string) {
	o.ServerOrPath.Set(&v)
}

// SetServerOrPathNil sets the value for ServerOrPath to be an explicit nil
func (o *ApplicationDiscoveryRequestModel) SetServerOrPathNil() {
	o.ServerOrPath.Set(nil)
}

// UnsetServerOrPath ensures that no value is present for ServerOrPath, not even an explicit nil
func (o *ApplicationDiscoveryRequestModel) UnsetServerOrPath() {
	o.ServerOrPath.Unset()
}

func (o ApplicationDiscoveryRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationDiscoveryRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationGroups != nil {
		toSerialize["ApplicationGroups"] = o.ApplicationGroups
	}
	if o.DeliveryGroups != nil {
		toSerialize["DeliveryGroups"] = o.DeliveryGroups
	}
	if o.MachineCatalogs != nil {
		toSerialize["MachineCatalogs"] = o.MachineCatalogs
	}
	if o.ServerOrPath.IsSet() {
		toSerialize["ServerOrPath"] = o.ServerOrPath.Get()
	}
	return toSerialize, nil
}

type NullableApplicationDiscoveryRequestModel struct {
	value *ApplicationDiscoveryRequestModel
	isSet bool
}

func (v NullableApplicationDiscoveryRequestModel) Get() *ApplicationDiscoveryRequestModel {
	return v.value
}

func (v *NullableApplicationDiscoveryRequestModel) Set(val *ApplicationDiscoveryRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationDiscoveryRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationDiscoveryRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationDiscoveryRequestModel(val *ApplicationDiscoveryRequestModel) *NullableApplicationDiscoveryRequestModel {
	return &NullableApplicationDiscoveryRequestModel{value: val, isSet: true}
}

func (v NullableApplicationDiscoveryRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationDiscoveryRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
