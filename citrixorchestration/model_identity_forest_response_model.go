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

// checks if the IdentityForestResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityForestResponseModel{}

// IdentityForestResponseModel The Forest identity object.
type IdentityForestResponseModel struct {
	// The name associated with the forest object.
	Name NullableString `json:"Name,omitempty"`
	// The root domain of the forest.
	RootDomain NullableString `json:"RootDomain,omitempty"`
	// The domains under the forest.
	Domains []string `json:"Domains,omitempty"`
	// The sites under the forest.
	Sites []string `json:"Sites,omitempty"`
	// Other forests that have trust with this forest.
	TrustedForests []string `json:"TrustedForests,omitempty"`
	// Set to True during a bulk identity lookup operation. For individual identity lookup, usually an exception will be throw.
	PossibleLookupFailure *bool `json:"PossibleLookupFailure,omitempty"`
	// Properties fetched and populated in the Forest object.  This is a bitfield indicating the fetched properties.
	PropertiesFetched int32 `json:"PropertiesFetched"`
}

// NewIdentityForestResponseModel instantiates a new IdentityForestResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityForestResponseModel(propertiesFetched int32) *IdentityForestResponseModel {
	this := IdentityForestResponseModel{}
	this.PropertiesFetched = propertiesFetched
	return &this
}

// NewIdentityForestResponseModelWithDefaults instantiates a new IdentityForestResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityForestResponseModelWithDefaults() *IdentityForestResponseModel {
	this := IdentityForestResponseModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityForestResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityForestResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IdentityForestResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IdentityForestResponseModel) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *IdentityForestResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IdentityForestResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetRootDomain returns the RootDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityForestResponseModel) GetRootDomain() string {
	if o == nil || IsNil(o.RootDomain.Get()) {
		var ret string
		return ret
	}
	return *o.RootDomain.Get()
}

// GetRootDomainOk returns a tuple with the RootDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityForestResponseModel) GetRootDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootDomain.Get(), o.RootDomain.IsSet()
}

// HasRootDomain returns a boolean if a field has been set.
func (o *IdentityForestResponseModel) HasRootDomain() bool {
	if o != nil && o.RootDomain.IsSet() {
		return true
	}

	return false
}

// SetRootDomain gets a reference to the given NullableString and assigns it to the RootDomain field.
func (o *IdentityForestResponseModel) SetRootDomain(v string) {
	o.RootDomain.Set(&v)
}

// SetRootDomainNil sets the value for RootDomain to be an explicit nil
func (o *IdentityForestResponseModel) SetRootDomainNil() {
	o.RootDomain.Set(nil)
}

// UnsetRootDomain ensures that no value is present for RootDomain, not even an explicit nil
func (o *IdentityForestResponseModel) UnsetRootDomain() {
	o.RootDomain.Unset()
}

// GetDomains returns the Domains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityForestResponseModel) GetDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityForestResponseModel) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *IdentityForestResponseModel) HasDomains() bool {
	if o != nil && IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *IdentityForestResponseModel) SetDomains(v []string) {
	o.Domains = v
}

// GetSites returns the Sites field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityForestResponseModel) GetSites() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityForestResponseModel) GetSitesOk() ([]string, bool) {
	if o == nil || IsNil(o.Sites) {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *IdentityForestResponseModel) HasSites() bool {
	if o != nil && IsNil(o.Sites) {
		return true
	}

	return false
}

// SetSites gets a reference to the given []string and assigns it to the Sites field.
func (o *IdentityForestResponseModel) SetSites(v []string) {
	o.Sites = v
}

// GetTrustedForests returns the TrustedForests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityForestResponseModel) GetTrustedForests() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TrustedForests
}

// GetTrustedForestsOk returns a tuple with the TrustedForests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityForestResponseModel) GetTrustedForestsOk() ([]string, bool) {
	if o == nil || IsNil(o.TrustedForests) {
		return nil, false
	}
	return o.TrustedForests, true
}

// HasTrustedForests returns a boolean if a field has been set.
func (o *IdentityForestResponseModel) HasTrustedForests() bool {
	if o != nil && IsNil(o.TrustedForests) {
		return true
	}

	return false
}

// SetTrustedForests gets a reference to the given []string and assigns it to the TrustedForests field.
func (o *IdentityForestResponseModel) SetTrustedForests(v []string) {
	o.TrustedForests = v
}

// GetPossibleLookupFailure returns the PossibleLookupFailure field value if set, zero value otherwise.
func (o *IdentityForestResponseModel) GetPossibleLookupFailure() bool {
	if o == nil || IsNil(o.PossibleLookupFailure) {
		var ret bool
		return ret
	}
	return *o.PossibleLookupFailure
}

// GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityForestResponseModel) GetPossibleLookupFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.PossibleLookupFailure) {
		return nil, false
	}
	return o.PossibleLookupFailure, true
}

// HasPossibleLookupFailure returns a boolean if a field has been set.
func (o *IdentityForestResponseModel) HasPossibleLookupFailure() bool {
	if o != nil && !IsNil(o.PossibleLookupFailure) {
		return true
	}

	return false
}

// SetPossibleLookupFailure gets a reference to the given bool and assigns it to the PossibleLookupFailure field.
func (o *IdentityForestResponseModel) SetPossibleLookupFailure(v bool) {
	o.PossibleLookupFailure = &v
}

// GetPropertiesFetched returns the PropertiesFetched field value
func (o *IdentityForestResponseModel) GetPropertiesFetched() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PropertiesFetched
}

// GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field value
// and a boolean to check if the value has been set.
func (o *IdentityForestResponseModel) GetPropertiesFetchedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertiesFetched, true
}

// SetPropertiesFetched sets field value
func (o *IdentityForestResponseModel) SetPropertiesFetched(v int32) {
	o.PropertiesFetched = v
}

func (o IdentityForestResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityForestResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.RootDomain.IsSet() {
		toSerialize["RootDomain"] = o.RootDomain.Get()
	}
	if o.Domains != nil {
		toSerialize["Domains"] = o.Domains
	}
	if o.Sites != nil {
		toSerialize["Sites"] = o.Sites
	}
	if o.TrustedForests != nil {
		toSerialize["TrustedForests"] = o.TrustedForests
	}
	if !IsNil(o.PossibleLookupFailure) {
		toSerialize["PossibleLookupFailure"] = o.PossibleLookupFailure
	}
	toSerialize["PropertiesFetched"] = o.PropertiesFetched
	return toSerialize, nil
}

type NullableIdentityForestResponseModel struct {
	value *IdentityForestResponseModel
	isSet bool
}

func (v NullableIdentityForestResponseModel) Get() *IdentityForestResponseModel {
	return v.value
}

func (v *NullableIdentityForestResponseModel) Set(val *IdentityForestResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityForestResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityForestResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityForestResponseModel(val *IdentityForestResponseModel) *NullableIdentityForestResponseModel {
	return &NullableIdentityForestResponseModel{value: val, isSet: true}
}

func (v NullableIdentityForestResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityForestResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
