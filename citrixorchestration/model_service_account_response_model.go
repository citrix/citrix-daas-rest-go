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

// checks if the ServiceAccountResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountResponseModel{}

// ServiceAccountResponseModel Response model for service account.
type ServiceAccountResponseModel struct {
	ServiceAccountUid *string `json:"ServiceAccountUid,omitempty"`
	// Gets or sets the identity provider type for the service account
	IdentityProviderType NullableString `json:"IdentityProviderType,omitempty"`
	// Gets or sets the identity provider id for the service account
	IdentityProviderIdentifier NullableString `json:"IdentityProviderIdentifier,omitempty"`
	// Gets or sets the identifier for the service account
	AccountId NullableString `json:"AccountId,omitempty"`
	// Gets or sets the secret expiration time for the service account
	SecretExpiryTime NullableString `json:"SecretExpiryTime,omitempty"`
	// Gets or sets the capabilities for the service account
	Capabilities []ServiceAccountCapabilityReference `json:"Capabilities,omitempty"`
	// Gets or sets the healthy status for the service account
	IsHealthy *bool `json:"IsHealthy,omitempty"`
	// Gets or sets the failure reason for the service account if IsHealthy is false
	FailureReason NullableString `json:"FailureReason,omitempty"`
	// Gets or sets the array of administration scopes
	Scopes []ScopeResponseModel `json:"Scopes,omitempty"`
	// Gets or sets the tenant id
	TenantId NullableString `json:"TenantId,omitempty"`
	// Gets or sets the display name
	DisplayName NullableString `json:"DisplayName,omitempty"`
	// Gets or sets the description
	Description NullableString `json:"Description,omitempty"`
}

// NewServiceAccountResponseModel instantiates a new ServiceAccountResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountResponseModel() *ServiceAccountResponseModel {
	this := ServiceAccountResponseModel{}
	return &this
}

// NewServiceAccountResponseModelWithDefaults instantiates a new ServiceAccountResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountResponseModelWithDefaults() *ServiceAccountResponseModel {
	this := ServiceAccountResponseModel{}
	return &this
}

// GetServiceAccountUid returns the ServiceAccountUid field value if set, zero value otherwise.
func (o *ServiceAccountResponseModel) GetServiceAccountUid() string {
	if o == nil || IsNil(o.ServiceAccountUid) {
		var ret string
		return ret
	}
	return *o.ServiceAccountUid
}

// GetServiceAccountUidOk returns a tuple with the ServiceAccountUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponseModel) GetServiceAccountUidOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccountUid) {
		return nil, false
	}
	return o.ServiceAccountUid, true
}

// HasServiceAccountUid returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasServiceAccountUid() bool {
	if o != nil && !IsNil(o.ServiceAccountUid) {
		return true
	}

	return false
}

// SetServiceAccountUid gets a reference to the given string and assigns it to the ServiceAccountUid field.
func (o *ServiceAccountResponseModel) SetServiceAccountUid(v string) {
	o.ServiceAccountUid = &v
}

// GetIdentityProviderType returns the IdentityProviderType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetIdentityProviderType() string {
	if o == nil || IsNil(o.IdentityProviderType.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityProviderType.Get()
}

// GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetIdentityProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityProviderType.Get(), o.IdentityProviderType.IsSet()
}

// HasIdentityProviderType returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasIdentityProviderType() bool {
	if o != nil && o.IdentityProviderType.IsSet() {
		return true
	}

	return false
}

// SetIdentityProviderType gets a reference to the given NullableString and assigns it to the IdentityProviderType field.
func (o *ServiceAccountResponseModel) SetIdentityProviderType(v string) {
	o.IdentityProviderType.Set(&v)
}
// SetIdentityProviderTypeNil sets the value for IdentityProviderType to be an explicit nil
func (o *ServiceAccountResponseModel) SetIdentityProviderTypeNil() {
	o.IdentityProviderType.Set(nil)
}

// UnsetIdentityProviderType ensures that no value is present for IdentityProviderType, not even an explicit nil
func (o *ServiceAccountResponseModel) UnsetIdentityProviderType() {
	o.IdentityProviderType.Unset()
}

// GetIdentityProviderIdentifier returns the IdentityProviderIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetIdentityProviderIdentifier() string {
	if o == nil || IsNil(o.IdentityProviderIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityProviderIdentifier.Get()
}

// GetIdentityProviderIdentifierOk returns a tuple with the IdentityProviderIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetIdentityProviderIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityProviderIdentifier.Get(), o.IdentityProviderIdentifier.IsSet()
}

// HasIdentityProviderIdentifier returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasIdentityProviderIdentifier() bool {
	if o != nil && o.IdentityProviderIdentifier.IsSet() {
		return true
	}

	return false
}

// SetIdentityProviderIdentifier gets a reference to the given NullableString and assigns it to the IdentityProviderIdentifier field.
func (o *ServiceAccountResponseModel) SetIdentityProviderIdentifier(v string) {
	o.IdentityProviderIdentifier.Set(&v)
}
// SetIdentityProviderIdentifierNil sets the value for IdentityProviderIdentifier to be an explicit nil
func (o *ServiceAccountResponseModel) SetIdentityProviderIdentifierNil() {
	o.IdentityProviderIdentifier.Set(nil)
}

// UnsetIdentityProviderIdentifier ensures that no value is present for IdentityProviderIdentifier, not even an explicit nil
func (o *ServiceAccountResponseModel) UnsetIdentityProviderIdentifier() {
	o.IdentityProviderIdentifier.Unset()
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId.Get()) {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *ServiceAccountResponseModel) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *ServiceAccountResponseModel) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *ServiceAccountResponseModel) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetSecretExpiryTime returns the SecretExpiryTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetSecretExpiryTime() string {
	if o == nil || IsNil(o.SecretExpiryTime.Get()) {
		var ret string
		return ret
	}
	return *o.SecretExpiryTime.Get()
}

// GetSecretExpiryTimeOk returns a tuple with the SecretExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetSecretExpiryTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretExpiryTime.Get(), o.SecretExpiryTime.IsSet()
}

// HasSecretExpiryTime returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasSecretExpiryTime() bool {
	if o != nil && o.SecretExpiryTime.IsSet() {
		return true
	}

	return false
}

// SetSecretExpiryTime gets a reference to the given NullableString and assigns it to the SecretExpiryTime field.
func (o *ServiceAccountResponseModel) SetSecretExpiryTime(v string) {
	o.SecretExpiryTime.Set(&v)
}
// SetSecretExpiryTimeNil sets the value for SecretExpiryTime to be an explicit nil
func (o *ServiceAccountResponseModel) SetSecretExpiryTimeNil() {
	o.SecretExpiryTime.Set(nil)
}

// UnsetSecretExpiryTime ensures that no value is present for SecretExpiryTime, not even an explicit nil
func (o *ServiceAccountResponseModel) UnsetSecretExpiryTime() {
	o.SecretExpiryTime.Unset()
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetCapabilities() []ServiceAccountCapabilityReference {
	if o == nil {
		var ret []ServiceAccountCapabilityReference
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetCapabilitiesOk() ([]ServiceAccountCapabilityReference, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasCapabilities() bool {
	if o != nil && IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []ServiceAccountCapabilityReference and assigns it to the Capabilities field.
func (o *ServiceAccountResponseModel) SetCapabilities(v []ServiceAccountCapabilityReference) {
	o.Capabilities = v
}

// GetIsHealthy returns the IsHealthy field value if set, zero value otherwise.
func (o *ServiceAccountResponseModel) GetIsHealthy() bool {
	if o == nil || IsNil(o.IsHealthy) {
		var ret bool
		return ret
	}
	return *o.IsHealthy
}

// GetIsHealthyOk returns a tuple with the IsHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponseModel) GetIsHealthyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHealthy) {
		return nil, false
	}
	return o.IsHealthy, true
}

// HasIsHealthy returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasIsHealthy() bool {
	if o != nil && !IsNil(o.IsHealthy) {
		return true
	}

	return false
}

// SetIsHealthy gets a reference to the given bool and assigns it to the IsHealthy field.
func (o *ServiceAccountResponseModel) SetIsHealthy(v bool) {
	o.IsHealthy = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason.Get()) {
		var ret string
		return ret
	}
	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// HasFailureReason returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasFailureReason() bool {
	if o != nil && o.FailureReason.IsSet() {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given NullableString and assigns it to the FailureReason field.
func (o *ServiceAccountResponseModel) SetFailureReason(v string) {
	o.FailureReason.Set(&v)
}
// SetFailureReasonNil sets the value for FailureReason to be an explicit nil
func (o *ServiceAccountResponseModel) SetFailureReasonNil() {
	o.FailureReason.Set(nil)
}

// UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
func (o *ServiceAccountResponseModel) UnsetFailureReason() {
	o.FailureReason.Unset()
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetScopes() []ScopeResponseModel {
	if o == nil {
		var ret []ScopeResponseModel
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetScopesOk() ([]ScopeResponseModel, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasScopes() bool {
	if o != nil && IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []ScopeResponseModel and assigns it to the Scopes field.
func (o *ServiceAccountResponseModel) SetScopes(v []ScopeResponseModel) {
	o.Scopes = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *ServiceAccountResponseModel) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *ServiceAccountResponseModel) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *ServiceAccountResponseModel) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *ServiceAccountResponseModel) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *ServiceAccountResponseModel) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *ServiceAccountResponseModel) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountResponseModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ServiceAccountResponseModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ServiceAccountResponseModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ServiceAccountResponseModel) UnsetDescription() {
	o.Description.Unset()
}

func (o ServiceAccountResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceAccountUid) {
		toSerialize["ServiceAccountUid"] = o.ServiceAccountUid
	}
	if o.IdentityProviderType.IsSet() {
		toSerialize["IdentityProviderType"] = o.IdentityProviderType.Get()
	}
	if o.IdentityProviderIdentifier.IsSet() {
		toSerialize["IdentityProviderIdentifier"] = o.IdentityProviderIdentifier.Get()
	}
	if o.AccountId.IsSet() {
		toSerialize["AccountId"] = o.AccountId.Get()
	}
	if o.SecretExpiryTime.IsSet() {
		toSerialize["SecretExpiryTime"] = o.SecretExpiryTime.Get()
	}
	if o.Capabilities != nil {
		toSerialize["Capabilities"] = o.Capabilities
	}
	if !IsNil(o.IsHealthy) {
		toSerialize["IsHealthy"] = o.IsHealthy
	}
	if o.FailureReason.IsSet() {
		toSerialize["FailureReason"] = o.FailureReason.Get()
	}
	if o.Scopes != nil {
		toSerialize["Scopes"] = o.Scopes
	}
	if o.TenantId.IsSet() {
		toSerialize["TenantId"] = o.TenantId.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["DisplayName"] = o.DisplayName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableServiceAccountResponseModel struct {
	value *ServiceAccountResponseModel
	isSet bool
}

func (v NullableServiceAccountResponseModel) Get() *ServiceAccountResponseModel {
	return v.value
}

func (v *NullableServiceAccountResponseModel) Set(val *ServiceAccountResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountResponseModel(val *ServiceAccountResponseModel) *NullableServiceAccountResponseModel {
	return &NullableServiceAccountResponseModel{value: val, isSet: true}
}

func (v NullableServiceAccountResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

