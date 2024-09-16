/*
Administrators APIs

APIs for managing CC administrators.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
)

// checks if the PendingAdministratorInviteExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingAdministratorInviteExternal{}

// PendingAdministratorInviteExternal struct for PendingAdministratorInviteExternal
type PendingAdministratorInviteExternal struct {
	Email string `json:"email"`
	RequestorEmail NullableString `json:"requestorEmail,omitempty"`
	FirstName NullableString `json:"firstName,omitempty"`
	LastName NullableString `json:"lastName,omitempty"`
	Access *AdministratorAccessModel `json:"access,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	ExternalProviderType AdministratorExternalProviderType `json:"externalProviderType"`
	ExternalProviderId string `json:"externalProviderId"`
	ExternalUserId string `json:"externalUserId"`
	ExternalProviderProperties map[string]interface{} `json:"externalProviderProperties,omitempty"`
	ExternalProviderAuthDomain NullableString `json:"externalProviderAuthDomain,omitempty"`
}

// NewPendingAdministratorInviteExternal instantiates a new PendingAdministratorInviteExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingAdministratorInviteExternal(email string, externalProviderType AdministratorExternalProviderType, externalProviderId string, externalUserId string) *PendingAdministratorInviteExternal {
	this := PendingAdministratorInviteExternal{}
	this.Email = email
	this.ExternalProviderType = externalProviderType
	this.ExternalProviderId = externalProviderId
	this.ExternalUserId = externalUserId
	return &this
}

// NewPendingAdministratorInviteExternalWithDefaults instantiates a new PendingAdministratorInviteExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingAdministratorInviteExternalWithDefaults() *PendingAdministratorInviteExternal {
	this := PendingAdministratorInviteExternal{}
	return &this
}

// GetEmail returns the Email field value
func (o *PendingAdministratorInviteExternal) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PendingAdministratorInviteExternal) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PendingAdministratorInviteExternal) SetEmail(v string) {
	o.Email = v
}

// GetRequestorEmail returns the RequestorEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingAdministratorInviteExternal) GetRequestorEmail() string {
	if o == nil || IsNil(o.RequestorEmail.Get()) {
		var ret string
		return ret
	}
	return *o.RequestorEmail.Get()
}

// GetRequestorEmailOk returns a tuple with the RequestorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingAdministratorInviteExternal) GetRequestorEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestorEmail.Get(), o.RequestorEmail.IsSet()
}

// HasRequestorEmail returns a boolean if a field has been set.
func (o *PendingAdministratorInviteExternal) HasRequestorEmail() bool {
	if o != nil && o.RequestorEmail.IsSet() {
		return true
	}

	return false
}

// SetRequestorEmail gets a reference to the given NullableString and assigns it to the RequestorEmail field.
func (o *PendingAdministratorInviteExternal) SetRequestorEmail(v string) {
	o.RequestorEmail.Set(&v)
}
// SetRequestorEmailNil sets the value for RequestorEmail to be an explicit nil
func (o *PendingAdministratorInviteExternal) SetRequestorEmailNil() {
	o.RequestorEmail.Set(nil)
}

// UnsetRequestorEmail ensures that no value is present for RequestorEmail, not even an explicit nil
func (o *PendingAdministratorInviteExternal) UnsetRequestorEmail() {
	o.RequestorEmail.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingAdministratorInviteExternal) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingAdministratorInviteExternal) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *PendingAdministratorInviteExternal) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *PendingAdministratorInviteExternal) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *PendingAdministratorInviteExternal) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *PendingAdministratorInviteExternal) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingAdministratorInviteExternal) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingAdministratorInviteExternal) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *PendingAdministratorInviteExternal) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *PendingAdministratorInviteExternal) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *PendingAdministratorInviteExternal) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *PendingAdministratorInviteExternal) UnsetLastName() {
	o.LastName.Unset()
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *PendingAdministratorInviteExternal) GetAccess() AdministratorAccessModel {
	if o == nil || IsNil(o.Access) {
		var ret AdministratorAccessModel
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingAdministratorInviteExternal) GetAccessOk() (*AdministratorAccessModel, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *PendingAdministratorInviteExternal) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AdministratorAccessModel and assigns it to the Access field.
func (o *PendingAdministratorInviteExternal) SetAccess(v AdministratorAccessModel) {
	o.Access = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingAdministratorInviteExternal) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingAdministratorInviteExternal) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PendingAdministratorInviteExternal) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *PendingAdministratorInviteExternal) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *PendingAdministratorInviteExternal) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *PendingAdministratorInviteExternal) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetExternalProviderType returns the ExternalProviderType field value
func (o *PendingAdministratorInviteExternal) GetExternalProviderType() AdministratorExternalProviderType {
	if o == nil {
		var ret AdministratorExternalProviderType
		return ret
	}

	return o.ExternalProviderType
}

// GetExternalProviderTypeOk returns a tuple with the ExternalProviderType field value
// and a boolean to check if the value has been set.
func (o *PendingAdministratorInviteExternal) GetExternalProviderTypeOk() (*AdministratorExternalProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalProviderType, true
}

// SetExternalProviderType sets field value
func (o *PendingAdministratorInviteExternal) SetExternalProviderType(v AdministratorExternalProviderType) {
	o.ExternalProviderType = v
}

// GetExternalProviderId returns the ExternalProviderId field value
func (o *PendingAdministratorInviteExternal) GetExternalProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalProviderId
}

// GetExternalProviderIdOk returns a tuple with the ExternalProviderId field value
// and a boolean to check if the value has been set.
func (o *PendingAdministratorInviteExternal) GetExternalProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalProviderId, true
}

// SetExternalProviderId sets field value
func (o *PendingAdministratorInviteExternal) SetExternalProviderId(v string) {
	o.ExternalProviderId = v
}

// GetExternalUserId returns the ExternalUserId field value
func (o *PendingAdministratorInviteExternal) GetExternalUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value
// and a boolean to check if the value has been set.
func (o *PendingAdministratorInviteExternal) GetExternalUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUserId, true
}

// SetExternalUserId sets field value
func (o *PendingAdministratorInviteExternal) SetExternalUserId(v string) {
	o.ExternalUserId = v
}

// GetExternalProviderProperties returns the ExternalProviderProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingAdministratorInviteExternal) GetExternalProviderProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalProviderProperties
}

// GetExternalProviderPropertiesOk returns a tuple with the ExternalProviderProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingAdministratorInviteExternal) GetExternalProviderPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExternalProviderProperties) {
		return map[string]interface{}{}, false
	}
	return o.ExternalProviderProperties, true
}

// HasExternalProviderProperties returns a boolean if a field has been set.
func (o *PendingAdministratorInviteExternal) HasExternalProviderProperties() bool {
	if o != nil && IsNil(o.ExternalProviderProperties) {
		return true
	}

	return false
}

// SetExternalProviderProperties gets a reference to the given map[string]interface{} and assigns it to the ExternalProviderProperties field.
func (o *PendingAdministratorInviteExternal) SetExternalProviderProperties(v map[string]interface{}) {
	o.ExternalProviderProperties = v
}

// GetExternalProviderAuthDomain returns the ExternalProviderAuthDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingAdministratorInviteExternal) GetExternalProviderAuthDomain() string {
	if o == nil || IsNil(o.ExternalProviderAuthDomain.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalProviderAuthDomain.Get()
}

// GetExternalProviderAuthDomainOk returns a tuple with the ExternalProviderAuthDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingAdministratorInviteExternal) GetExternalProviderAuthDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalProviderAuthDomain.Get(), o.ExternalProviderAuthDomain.IsSet()
}

// HasExternalProviderAuthDomain returns a boolean if a field has been set.
func (o *PendingAdministratorInviteExternal) HasExternalProviderAuthDomain() bool {
	if o != nil && o.ExternalProviderAuthDomain.IsSet() {
		return true
	}

	return false
}

// SetExternalProviderAuthDomain gets a reference to the given NullableString and assigns it to the ExternalProviderAuthDomain field.
func (o *PendingAdministratorInviteExternal) SetExternalProviderAuthDomain(v string) {
	o.ExternalProviderAuthDomain.Set(&v)
}
// SetExternalProviderAuthDomainNil sets the value for ExternalProviderAuthDomain to be an explicit nil
func (o *PendingAdministratorInviteExternal) SetExternalProviderAuthDomainNil() {
	o.ExternalProviderAuthDomain.Set(nil)
}

// UnsetExternalProviderAuthDomain ensures that no value is present for ExternalProviderAuthDomain, not even an explicit nil
func (o *PendingAdministratorInviteExternal) UnsetExternalProviderAuthDomain() {
	o.ExternalProviderAuthDomain.Unset()
}

func (o PendingAdministratorInviteExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingAdministratorInviteExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if o.RequestorEmail.IsSet() {
		toSerialize["requestorEmail"] = o.RequestorEmail.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	toSerialize["externalProviderType"] = o.ExternalProviderType
	toSerialize["externalProviderId"] = o.ExternalProviderId
	toSerialize["externalUserId"] = o.ExternalUserId
	if o.ExternalProviderProperties != nil {
		toSerialize["externalProviderProperties"] = o.ExternalProviderProperties
	}
	if o.ExternalProviderAuthDomain.IsSet() {
		toSerialize["externalProviderAuthDomain"] = o.ExternalProviderAuthDomain.Get()
	}
	return toSerialize, nil
}

type NullablePendingAdministratorInviteExternal struct {
	value *PendingAdministratorInviteExternal
	isSet bool
}

func (v NullablePendingAdministratorInviteExternal) Get() *PendingAdministratorInviteExternal {
	return v.value
}

func (v *NullablePendingAdministratorInviteExternal) Set(val *PendingAdministratorInviteExternal) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingAdministratorInviteExternal) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingAdministratorInviteExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingAdministratorInviteExternal(val *PendingAdministratorInviteExternal) *NullablePendingAdministratorInviteExternal {
	return &NullablePendingAdministratorInviteExternal{value: val, isSet: true}
}

func (v NullablePendingAdministratorInviteExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingAdministratorInviteExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

