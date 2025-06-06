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

// checks if the RecommendationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationModel{}

// RecommendationModel struct for RecommendationModel
type RecommendationModel struct {
	// The unique identifier of this recommendation
	RecommendationId *string `json:"RecommendationId,omitempty"`
	// The full description of this recommendation
	RecommendationDescription NullableString `json:"RecommendationDescription,omitempty"`
	// The source of this recommendation
	RecommendationSource NullableString `json:"RecommendationSource,omitempty"`
	// The type of this recommendation.
	RecommendationType NullableString `json:"RecommendationType,omitempty"`
	// The state of this recommendation
	RecommendationState NullableString `json:"RecommendationState,omitempty"`
	// The metadata associated with this recommendation
	RecommendationMetadata NullableString `json:"RecommendationMetadata,omitempty"`
	// The reason why this recommendation was dismissed or rejected by the user
	Reason NullableString `json:"Reason,omitempty"`
	// The unique identifier of the target associated with this recommendation
	TargetUid *string `json:"TargetUid,omitempty"`
	// The name of the target associated with this recommendation. E.g. Provisioning Scheme Name
	TargetName NullableString `json:"TargetName,omitempty"`
	// The type of target associated with this recommendation. E.g. Provisioning Scheme
	TargetType NullableString `json:"TargetType,omitempty"`
	// The datetime at which this recommendation was generated by the system
	CreatedDateTimeInUtc NullableString `json:"CreatedDateTimeInUtc,omitempty"`
	// The datetime at which this recommendation was updated
	UpdatedDateTimeInUtc NullableString `json:"UpdatedDateTimeInUtc,omitempty"`
	// The datetime at which this recommendation needs to be reminded
	RemindInDateTimeInUtc NullableString `json:"RemindInDateTimeInUtc,omitempty"`
	// The machine catalog unique identifier associated with this recommendation
	MachineCatalogId NullableString `json:"MachineCatalogId,omitempty"`
	// The machine catalog name associated with this recommendation
	MachineCatalogName NullableString `json:"MachineCatalogName,omitempty"`
}

// NewRecommendationModel instantiates a new RecommendationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationModel() *RecommendationModel {
	this := RecommendationModel{}
	return &this
}

// NewRecommendationModelWithDefaults instantiates a new RecommendationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationModelWithDefaults() *RecommendationModel {
	this := RecommendationModel{}
	return &this
}

// GetRecommendationId returns the RecommendationId field value if set, zero value otherwise.
func (o *RecommendationModel) GetRecommendationId() string {
	if o == nil || IsNil(o.RecommendationId) {
		var ret string
		return ret
	}
	return *o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationModel) GetRecommendationIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationId) {
		return nil, false
	}
	return o.RecommendationId, true
}

// HasRecommendationId returns a boolean if a field has been set.
func (o *RecommendationModel) HasRecommendationId() bool {
	if o != nil && !IsNil(o.RecommendationId) {
		return true
	}

	return false
}

// SetRecommendationId gets a reference to the given string and assigns it to the RecommendationId field.
func (o *RecommendationModel) SetRecommendationId(v string) {
	o.RecommendationId = &v
}

// GetRecommendationDescription returns the RecommendationDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetRecommendationDescription() string {
	if o == nil || IsNil(o.RecommendationDescription.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendationDescription.Get()
}

// GetRecommendationDescriptionOk returns a tuple with the RecommendationDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetRecommendationDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendationDescription.Get(), o.RecommendationDescription.IsSet()
}

// HasRecommendationDescription returns a boolean if a field has been set.
func (o *RecommendationModel) HasRecommendationDescription() bool {
	if o != nil && o.RecommendationDescription.IsSet() {
		return true
	}

	return false
}

// SetRecommendationDescription gets a reference to the given NullableString and assigns it to the RecommendationDescription field.
func (o *RecommendationModel) SetRecommendationDescription(v string) {
	o.RecommendationDescription.Set(&v)
}

// SetRecommendationDescriptionNil sets the value for RecommendationDescription to be an explicit nil
func (o *RecommendationModel) SetRecommendationDescriptionNil() {
	o.RecommendationDescription.Set(nil)
}

// UnsetRecommendationDescription ensures that no value is present for RecommendationDescription, not even an explicit nil
func (o *RecommendationModel) UnsetRecommendationDescription() {
	o.RecommendationDescription.Unset()
}

// GetRecommendationSource returns the RecommendationSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetRecommendationSource() string {
	if o == nil || IsNil(o.RecommendationSource.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendationSource.Get()
}

// GetRecommendationSourceOk returns a tuple with the RecommendationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetRecommendationSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendationSource.Get(), o.RecommendationSource.IsSet()
}

// HasRecommendationSource returns a boolean if a field has been set.
func (o *RecommendationModel) HasRecommendationSource() bool {
	if o != nil && o.RecommendationSource.IsSet() {
		return true
	}

	return false
}

// SetRecommendationSource gets a reference to the given NullableString and assigns it to the RecommendationSource field.
func (o *RecommendationModel) SetRecommendationSource(v string) {
	o.RecommendationSource.Set(&v)
}

// SetRecommendationSourceNil sets the value for RecommendationSource to be an explicit nil
func (o *RecommendationModel) SetRecommendationSourceNil() {
	o.RecommendationSource.Set(nil)
}

// UnsetRecommendationSource ensures that no value is present for RecommendationSource, not even an explicit nil
func (o *RecommendationModel) UnsetRecommendationSource() {
	o.RecommendationSource.Unset()
}

// GetRecommendationType returns the RecommendationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetRecommendationType() string {
	if o == nil || IsNil(o.RecommendationType.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendationType.Get()
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendationType.Get(), o.RecommendationType.IsSet()
}

// HasRecommendationType returns a boolean if a field has been set.
func (o *RecommendationModel) HasRecommendationType() bool {
	if o != nil && o.RecommendationType.IsSet() {
		return true
	}

	return false
}

// SetRecommendationType gets a reference to the given NullableString and assigns it to the RecommendationType field.
func (o *RecommendationModel) SetRecommendationType(v string) {
	o.RecommendationType.Set(&v)
}

// SetRecommendationTypeNil sets the value for RecommendationType to be an explicit nil
func (o *RecommendationModel) SetRecommendationTypeNil() {
	o.RecommendationType.Set(nil)
}

// UnsetRecommendationType ensures that no value is present for RecommendationType, not even an explicit nil
func (o *RecommendationModel) UnsetRecommendationType() {
	o.RecommendationType.Unset()
}

// GetRecommendationState returns the RecommendationState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetRecommendationState() string {
	if o == nil || IsNil(o.RecommendationState.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendationState.Get()
}

// GetRecommendationStateOk returns a tuple with the RecommendationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetRecommendationStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendationState.Get(), o.RecommendationState.IsSet()
}

// HasRecommendationState returns a boolean if a field has been set.
func (o *RecommendationModel) HasRecommendationState() bool {
	if o != nil && o.RecommendationState.IsSet() {
		return true
	}

	return false
}

// SetRecommendationState gets a reference to the given NullableString and assigns it to the RecommendationState field.
func (o *RecommendationModel) SetRecommendationState(v string) {
	o.RecommendationState.Set(&v)
}

// SetRecommendationStateNil sets the value for RecommendationState to be an explicit nil
func (o *RecommendationModel) SetRecommendationStateNil() {
	o.RecommendationState.Set(nil)
}

// UnsetRecommendationState ensures that no value is present for RecommendationState, not even an explicit nil
func (o *RecommendationModel) UnsetRecommendationState() {
	o.RecommendationState.Unset()
}

// GetRecommendationMetadata returns the RecommendationMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetRecommendationMetadata() string {
	if o == nil || IsNil(o.RecommendationMetadata.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendationMetadata.Get()
}

// GetRecommendationMetadataOk returns a tuple with the RecommendationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetRecommendationMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendationMetadata.Get(), o.RecommendationMetadata.IsSet()
}

// HasRecommendationMetadata returns a boolean if a field has been set.
func (o *RecommendationModel) HasRecommendationMetadata() bool {
	if o != nil && o.RecommendationMetadata.IsSet() {
		return true
	}

	return false
}

// SetRecommendationMetadata gets a reference to the given NullableString and assigns it to the RecommendationMetadata field.
func (o *RecommendationModel) SetRecommendationMetadata(v string) {
	o.RecommendationMetadata.Set(&v)
}

// SetRecommendationMetadataNil sets the value for RecommendationMetadata to be an explicit nil
func (o *RecommendationModel) SetRecommendationMetadataNil() {
	o.RecommendationMetadata.Set(nil)
}

// UnsetRecommendationMetadata ensures that no value is present for RecommendationMetadata, not even an explicit nil
func (o *RecommendationModel) UnsetRecommendationMetadata() {
	o.RecommendationMetadata.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *RecommendationModel) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *RecommendationModel) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil
func (o *RecommendationModel) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *RecommendationModel) UnsetReason() {
	o.Reason.Unset()
}

// GetTargetUid returns the TargetUid field value if set, zero value otherwise.
func (o *RecommendationModel) GetTargetUid() string {
	if o == nil || IsNil(o.TargetUid) {
		var ret string
		return ret
	}
	return *o.TargetUid
}

// GetTargetUidOk returns a tuple with the TargetUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationModel) GetTargetUidOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUid) {
		return nil, false
	}
	return o.TargetUid, true
}

// HasTargetUid returns a boolean if a field has been set.
func (o *RecommendationModel) HasTargetUid() bool {
	if o != nil && !IsNil(o.TargetUid) {
		return true
	}

	return false
}

// SetTargetUid gets a reference to the given string and assigns it to the TargetUid field.
func (o *RecommendationModel) SetTargetUid(v string) {
	o.TargetUid = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetTargetName() string {
	if o == nil || IsNil(o.TargetName.Get()) {
		var ret string
		return ret
	}
	return *o.TargetName.Get()
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetTargetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetName.Get(), o.TargetName.IsSet()
}

// HasTargetName returns a boolean if a field has been set.
func (o *RecommendationModel) HasTargetName() bool {
	if o != nil && o.TargetName.IsSet() {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given NullableString and assigns it to the TargetName field.
func (o *RecommendationModel) SetTargetName(v string) {
	o.TargetName.Set(&v)
}

// SetTargetNameNil sets the value for TargetName to be an explicit nil
func (o *RecommendationModel) SetTargetNameNil() {
	o.TargetName.Set(nil)
}

// UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
func (o *RecommendationModel) UnsetTargetName() {
	o.TargetName.Unset()
}

// GetTargetType returns the TargetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetTargetType() string {
	if o == nil || IsNil(o.TargetType.Get()) {
		var ret string
		return ret
	}
	return *o.TargetType.Get()
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetTargetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetType.Get(), o.TargetType.IsSet()
}

// HasTargetType returns a boolean if a field has been set.
func (o *RecommendationModel) HasTargetType() bool {
	if o != nil && o.TargetType.IsSet() {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given NullableString and assigns it to the TargetType field.
func (o *RecommendationModel) SetTargetType(v string) {
	o.TargetType.Set(&v)
}

// SetTargetTypeNil sets the value for TargetType to be an explicit nil
func (o *RecommendationModel) SetTargetTypeNil() {
	o.TargetType.Set(nil)
}

// UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
func (o *RecommendationModel) UnsetTargetType() {
	o.TargetType.Unset()
}

// GetCreatedDateTimeInUtc returns the CreatedDateTimeInUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetCreatedDateTimeInUtc() string {
	if o == nil || IsNil(o.CreatedDateTimeInUtc.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedDateTimeInUtc.Get()
}

// GetCreatedDateTimeInUtcOk returns a tuple with the CreatedDateTimeInUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetCreatedDateTimeInUtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDateTimeInUtc.Get(), o.CreatedDateTimeInUtc.IsSet()
}

// HasCreatedDateTimeInUtc returns a boolean if a field has been set.
func (o *RecommendationModel) HasCreatedDateTimeInUtc() bool {
	if o != nil && o.CreatedDateTimeInUtc.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTimeInUtc gets a reference to the given NullableString and assigns it to the CreatedDateTimeInUtc field.
func (o *RecommendationModel) SetCreatedDateTimeInUtc(v string) {
	o.CreatedDateTimeInUtc.Set(&v)
}

// SetCreatedDateTimeInUtcNil sets the value for CreatedDateTimeInUtc to be an explicit nil
func (o *RecommendationModel) SetCreatedDateTimeInUtcNil() {
	o.CreatedDateTimeInUtc.Set(nil)
}

// UnsetCreatedDateTimeInUtc ensures that no value is present for CreatedDateTimeInUtc, not even an explicit nil
func (o *RecommendationModel) UnsetCreatedDateTimeInUtc() {
	o.CreatedDateTimeInUtc.Unset()
}

// GetUpdatedDateTimeInUtc returns the UpdatedDateTimeInUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetUpdatedDateTimeInUtc() string {
	if o == nil || IsNil(o.UpdatedDateTimeInUtc.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedDateTimeInUtc.Get()
}

// GetUpdatedDateTimeInUtcOk returns a tuple with the UpdatedDateTimeInUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetUpdatedDateTimeInUtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedDateTimeInUtc.Get(), o.UpdatedDateTimeInUtc.IsSet()
}

// HasUpdatedDateTimeInUtc returns a boolean if a field has been set.
func (o *RecommendationModel) HasUpdatedDateTimeInUtc() bool {
	if o != nil && o.UpdatedDateTimeInUtc.IsSet() {
		return true
	}

	return false
}

// SetUpdatedDateTimeInUtc gets a reference to the given NullableString and assigns it to the UpdatedDateTimeInUtc field.
func (o *RecommendationModel) SetUpdatedDateTimeInUtc(v string) {
	o.UpdatedDateTimeInUtc.Set(&v)
}

// SetUpdatedDateTimeInUtcNil sets the value for UpdatedDateTimeInUtc to be an explicit nil
func (o *RecommendationModel) SetUpdatedDateTimeInUtcNil() {
	o.UpdatedDateTimeInUtc.Set(nil)
}

// UnsetUpdatedDateTimeInUtc ensures that no value is present for UpdatedDateTimeInUtc, not even an explicit nil
func (o *RecommendationModel) UnsetUpdatedDateTimeInUtc() {
	o.UpdatedDateTimeInUtc.Unset()
}

// GetRemindInDateTimeInUtc returns the RemindInDateTimeInUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetRemindInDateTimeInUtc() string {
	if o == nil || IsNil(o.RemindInDateTimeInUtc.Get()) {
		var ret string
		return ret
	}
	return *o.RemindInDateTimeInUtc.Get()
}

// GetRemindInDateTimeInUtcOk returns a tuple with the RemindInDateTimeInUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetRemindInDateTimeInUtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemindInDateTimeInUtc.Get(), o.RemindInDateTimeInUtc.IsSet()
}

// HasRemindInDateTimeInUtc returns a boolean if a field has been set.
func (o *RecommendationModel) HasRemindInDateTimeInUtc() bool {
	if o != nil && o.RemindInDateTimeInUtc.IsSet() {
		return true
	}

	return false
}

// SetRemindInDateTimeInUtc gets a reference to the given NullableString and assigns it to the RemindInDateTimeInUtc field.
func (o *RecommendationModel) SetRemindInDateTimeInUtc(v string) {
	o.RemindInDateTimeInUtc.Set(&v)
}

// SetRemindInDateTimeInUtcNil sets the value for RemindInDateTimeInUtc to be an explicit nil
func (o *RecommendationModel) SetRemindInDateTimeInUtcNil() {
	o.RemindInDateTimeInUtc.Set(nil)
}

// UnsetRemindInDateTimeInUtc ensures that no value is present for RemindInDateTimeInUtc, not even an explicit nil
func (o *RecommendationModel) UnsetRemindInDateTimeInUtc() {
	o.RemindInDateTimeInUtc.Unset()
}

// GetMachineCatalogId returns the MachineCatalogId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetMachineCatalogId() string {
	if o == nil || IsNil(o.MachineCatalogId.Get()) {
		var ret string
		return ret
	}
	return *o.MachineCatalogId.Get()
}

// GetMachineCatalogIdOk returns a tuple with the MachineCatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetMachineCatalogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineCatalogId.Get(), o.MachineCatalogId.IsSet()
}

// HasMachineCatalogId returns a boolean if a field has been set.
func (o *RecommendationModel) HasMachineCatalogId() bool {
	if o != nil && o.MachineCatalogId.IsSet() {
		return true
	}

	return false
}

// SetMachineCatalogId gets a reference to the given NullableString and assigns it to the MachineCatalogId field.
func (o *RecommendationModel) SetMachineCatalogId(v string) {
	o.MachineCatalogId.Set(&v)
}

// SetMachineCatalogIdNil sets the value for MachineCatalogId to be an explicit nil
func (o *RecommendationModel) SetMachineCatalogIdNil() {
	o.MachineCatalogId.Set(nil)
}

// UnsetMachineCatalogId ensures that no value is present for MachineCatalogId, not even an explicit nil
func (o *RecommendationModel) UnsetMachineCatalogId() {
	o.MachineCatalogId.Unset()
}

// GetMachineCatalogName returns the MachineCatalogName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationModel) GetMachineCatalogName() string {
	if o == nil || IsNil(o.MachineCatalogName.Get()) {
		var ret string
		return ret
	}
	return *o.MachineCatalogName.Get()
}

// GetMachineCatalogNameOk returns a tuple with the MachineCatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationModel) GetMachineCatalogNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineCatalogName.Get(), o.MachineCatalogName.IsSet()
}

// HasMachineCatalogName returns a boolean if a field has been set.
func (o *RecommendationModel) HasMachineCatalogName() bool {
	if o != nil && o.MachineCatalogName.IsSet() {
		return true
	}

	return false
}

// SetMachineCatalogName gets a reference to the given NullableString and assigns it to the MachineCatalogName field.
func (o *RecommendationModel) SetMachineCatalogName(v string) {
	o.MachineCatalogName.Set(&v)
}

// SetMachineCatalogNameNil sets the value for MachineCatalogName to be an explicit nil
func (o *RecommendationModel) SetMachineCatalogNameNil() {
	o.MachineCatalogName.Set(nil)
}

// UnsetMachineCatalogName ensures that no value is present for MachineCatalogName, not even an explicit nil
func (o *RecommendationModel) UnsetMachineCatalogName() {
	o.MachineCatalogName.Unset()
}

func (o RecommendationModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecommendationId) {
		toSerialize["RecommendationId"] = o.RecommendationId
	}
	if o.RecommendationDescription.IsSet() {
		toSerialize["RecommendationDescription"] = o.RecommendationDescription.Get()
	}
	if o.RecommendationSource.IsSet() {
		toSerialize["RecommendationSource"] = o.RecommendationSource.Get()
	}
	if o.RecommendationType.IsSet() {
		toSerialize["RecommendationType"] = o.RecommendationType.Get()
	}
	if o.RecommendationState.IsSet() {
		toSerialize["RecommendationState"] = o.RecommendationState.Get()
	}
	if o.RecommendationMetadata.IsSet() {
		toSerialize["RecommendationMetadata"] = o.RecommendationMetadata.Get()
	}
	if o.Reason.IsSet() {
		toSerialize["Reason"] = o.Reason.Get()
	}
	if !IsNil(o.TargetUid) {
		toSerialize["TargetUid"] = o.TargetUid
	}
	if o.TargetName.IsSet() {
		toSerialize["TargetName"] = o.TargetName.Get()
	}
	if o.TargetType.IsSet() {
		toSerialize["TargetType"] = o.TargetType.Get()
	}
	if o.CreatedDateTimeInUtc.IsSet() {
		toSerialize["CreatedDateTimeInUtc"] = o.CreatedDateTimeInUtc.Get()
	}
	if o.UpdatedDateTimeInUtc.IsSet() {
		toSerialize["UpdatedDateTimeInUtc"] = o.UpdatedDateTimeInUtc.Get()
	}
	if o.RemindInDateTimeInUtc.IsSet() {
		toSerialize["RemindInDateTimeInUtc"] = o.RemindInDateTimeInUtc.Get()
	}
	if o.MachineCatalogId.IsSet() {
		toSerialize["MachineCatalogId"] = o.MachineCatalogId.Get()
	}
	if o.MachineCatalogName.IsSet() {
		toSerialize["MachineCatalogName"] = o.MachineCatalogName.Get()
	}
	return toSerialize, nil
}

type NullableRecommendationModel struct {
	value *RecommendationModel
	isSet bool
}

func (v NullableRecommendationModel) Get() *RecommendationModel {
	return v.value
}

func (v *NullableRecommendationModel) Set(val *RecommendationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationModel(val *RecommendationModel) *NullableRecommendationModel {
	return &NullableRecommendationModel{value: val, isSet: true}
}

func (v NullableRecommendationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
