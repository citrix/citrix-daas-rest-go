/*
Citrix.CloudServices.Administrators.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest{}

// CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest struct for CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest
type CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest struct {
	RbacRoles []string `json:"rbacRoles"`
	Operation string `json:"operation" validate:"regexp=^Delete$|^Update$"`
}

type _CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest

// NewCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest instantiates a new CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest(rbacRoles []string, operation string) *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest {
	this := CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest{}
	this.RbacRoles = rbacRoles
	this.Operation = operation
	return &this
}

// NewCitrixCloudServicesAdministratorsApiModelsRoleOperationRequestWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCitrixCloudServicesAdministratorsApiModelsRoleOperationRequestWithDefaults() *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest {
	this := CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest{}
	return &this
}

// GetRbacRoles returns the RbacRoles field value
func (o *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) GetRbacRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RbacRoles
}

// GetRbacRolesOk returns a tuple with the RbacRoles field value
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) GetRbacRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RbacRoles, true
}

// SetRbacRoles sets field value
func (o *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) SetRbacRoles(v []string) {
	o.RbacRoles = v
}

// GetOperation returns the Operation field value
func (o *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) SetOperation(v string) {
	o.Operation = v
}

func (o CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rbacRoles"] = o.RbacRoles
	toSerialize["operation"] = o.Operation
	return toSerialize, nil
}

func (o *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rbacRoles",
		"operation",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest := _CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest)

	if err != nil {
		return err
	}

	*o = CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest(varCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest)

	return err
}

type NullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest struct {
	value *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest
	isSet bool
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) Get() *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest {
	return v.value
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) Set(val *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest(val *CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) *NullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest {
	return &NullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

