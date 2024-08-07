/*
Citrix.CloudServices.Administrators.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
)

// checks if the CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences{}

// CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences struct for CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences
type CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences struct {
	SendNotificationEmails NullableBool `json:"sendNotificationEmails,omitempty"`
	NotificationsSubscribed []CitrixCloudServicesAdministratorsApiModelsAdministratorNotification `json:"notificationsSubscribed,omitempty"`
	EnabledDate NullableString `json:"enabledDate,omitempty"`
}

// NewCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences() *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences {
	this := CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences{}
	return &this
}

// NewCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferencesWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferencesWithDefaults() *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences {
	this := CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences{}
	return &this
}

// GetSendNotificationEmails returns the SendNotificationEmails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) GetSendNotificationEmails() bool {
	if o == nil || IsNil(o.SendNotificationEmails.Get()) {
		var ret bool
		return ret
	}
	return *o.SendNotificationEmails.Get()
}

// GetSendNotificationEmailsOk returns a tuple with the SendNotificationEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) GetSendNotificationEmailsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendNotificationEmails.Get(), o.SendNotificationEmails.IsSet()
}

// HasSendNotificationEmails returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) HasSendNotificationEmails() bool {
	if o != nil && o.SendNotificationEmails.IsSet() {
		return true
	}

	return false
}

// SetSendNotificationEmails gets a reference to the given NullableBool and assigns it to the SendNotificationEmails field.
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) SetSendNotificationEmails(v bool) {
	o.SendNotificationEmails.Set(&v)
}
// SetSendNotificationEmailsNil sets the value for SendNotificationEmails to be an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) SetSendNotificationEmailsNil() {
	o.SendNotificationEmails.Set(nil)
}

// UnsetSendNotificationEmails ensures that no value is present for SendNotificationEmails, not even an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) UnsetSendNotificationEmails() {
	o.SendNotificationEmails.Unset()
}

// GetNotificationsSubscribed returns the NotificationsSubscribed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) GetNotificationsSubscribed() []CitrixCloudServicesAdministratorsApiModelsAdministratorNotification {
	if o == nil {
		var ret []CitrixCloudServicesAdministratorsApiModelsAdministratorNotification
		return ret
	}
	return o.NotificationsSubscribed
}

// GetNotificationsSubscribedOk returns a tuple with the NotificationsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) GetNotificationsSubscribedOk() ([]CitrixCloudServicesAdministratorsApiModelsAdministratorNotification, bool) {
	if o == nil || IsNil(o.NotificationsSubscribed) {
		return nil, false
	}
	return o.NotificationsSubscribed, true
}

// HasNotificationsSubscribed returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) HasNotificationsSubscribed() bool {
	if o != nil && !IsNil(o.NotificationsSubscribed) {
		return true
	}

	return false
}

// SetNotificationsSubscribed gets a reference to the given []CitrixCloudServicesAdministratorsApiModelsAdministratorNotification and assigns it to the NotificationsSubscribed field.
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) SetNotificationsSubscribed(v []CitrixCloudServicesAdministratorsApiModelsAdministratorNotification) {
	o.NotificationsSubscribed = v
}

// GetEnabledDate returns the EnabledDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) GetEnabledDate() string {
	if o == nil || IsNil(o.EnabledDate.Get()) {
		var ret string
		return ret
	}
	return *o.EnabledDate.Get()
}

// GetEnabledDateOk returns a tuple with the EnabledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) GetEnabledDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnabledDate.Get(), o.EnabledDate.IsSet()
}

// HasEnabledDate returns a boolean if a field has been set.
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) HasEnabledDate() bool {
	if o != nil && o.EnabledDate.IsSet() {
		return true
	}

	return false
}

// SetEnabledDate gets a reference to the given NullableString and assigns it to the EnabledDate field.
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) SetEnabledDate(v string) {
	o.EnabledDate.Set(&v)
}
// SetEnabledDateNil sets the value for EnabledDate to be an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) SetEnabledDateNil() {
	o.EnabledDate.Set(nil)
}

// UnsetEnabledDate ensures that no value is present for EnabledDate, not even an explicit nil
func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) UnsetEnabledDate() {
	o.EnabledDate.Unset()
}

func (o CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SendNotificationEmails.IsSet() {
		toSerialize["sendNotificationEmails"] = o.SendNotificationEmails.Get()
	}
	if o.NotificationsSubscribed != nil {
		toSerialize["notificationsSubscribed"] = o.NotificationsSubscribed
	}
	if o.EnabledDate.IsSet() {
		toSerialize["enabledDate"] = o.EnabledDate.Get()
	}
	return toSerialize, nil
}

type NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences struct {
	value *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences
	isSet bool
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) Get() *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences {
	return v.value
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) Set(val *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences(val *CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences {
	return &NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

