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

// checks if the CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel{}

// CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel struct for CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel
type CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel struct {
	Emails []string `json:"emails,omitempty"`
	UserIds []string `json:"userIds,omitempty"`
	UcoIds []string `json:"ucoIds"`
}

type _CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel

// NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel instantiates a new CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel(ucoIds []string) *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel {
	this := CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel{}
	this.UcoIds = ucoIds
	return &this
}

// NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModelWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModelWithDefaults() *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel {
	this := CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) GetEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) GetEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) SetEmails(v []string) {
	o.Emails = v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) SetUserIds(v []string) {
	o.UserIds = v
}

// GetUcoIds returns the UcoIds field value
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) GetUcoIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UcoIds
}

// GetUcoIdsOk returns a tuple with the UcoIds field value
// and a boolean to check if the value has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) GetUcoIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UcoIds, true
}

// SetUcoIds sets field value
func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) SetUcoIds(v []string) {
	o.UcoIds = v
}

func (o CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	toSerialize["ucoIds"] = o.UcoIds
	return toSerialize, nil
}

func (o *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ucoIds",
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

	varCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel := _CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel)

	if err != nil {
		return err
	}

	*o = CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel(varCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel)

	return err
}

type NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel struct {
	value *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel
	isSet bool
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) Get() *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel {
	return v.value
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) Set(val *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel(val *CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) *NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel {
	return &NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

