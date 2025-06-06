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

// checks if the DetectOrphanedResourcesResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectOrphanedResourcesResponseModel{}

// DetectOrphanedResourcesResponseModel struct for DetectOrphanedResourcesResponseModel
type DetectOrphanedResourcesResponseModel struct {
	Hypervisor  *RefResponseModel `json:"Hypervisor,omitempty"`
	DetectScope *DetectScope      `json:"DetectScope,omitempty"`
	// The owner of the detect report
	CreatedBy NullableString `json:"CreatedBy,omitempty"`
	// The time of the detect report created
	ReportTime NullableString `json:"ReportTime,omitempty"`
	// The time of detect start.
	DetectStart NullableString `json:"DetectStart,omitempty"`
	// The time of detect end.
	DetectEnd NullableString `json:"DetectEnd,omitempty"`
	// The number of orphaned resources detected.
	NumDetected *int32 `json:"NumDetected,omitempty"`
	// Error message if orphaned resources detected failed.
	ErrorMsg NullableString `json:"ErrorMsg,omitempty"`
	// If orphaned resources detected failed
	IsFailed *bool `json:"IsFailed,omitempty"`
	// The detect results
	DetectResults []DetectOrphanedResourcesResultModel `json:"DetectResults,omitempty"`
}

// NewDetectOrphanedResourcesResponseModel instantiates a new DetectOrphanedResourcesResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectOrphanedResourcesResponseModel() *DetectOrphanedResourcesResponseModel {
	this := DetectOrphanedResourcesResponseModel{}
	return &this
}

// NewDetectOrphanedResourcesResponseModelWithDefaults instantiates a new DetectOrphanedResourcesResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectOrphanedResourcesResponseModelWithDefaults() *DetectOrphanedResourcesResponseModel {
	this := DetectOrphanedResourcesResponseModel{}
	return &this
}

// GetHypervisor returns the Hypervisor field value if set, zero value otherwise.
func (o *DetectOrphanedResourcesResponseModel) GetHypervisor() RefResponseModel {
	if o == nil || IsNil(o.Hypervisor) {
		var ret RefResponseModel
		return ret
	}
	return *o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectOrphanedResourcesResponseModel) GetHypervisorOk() (*RefResponseModel, bool) {
	if o == nil || IsNil(o.Hypervisor) {
		return nil, false
	}
	return o.Hypervisor, true
}

// HasHypervisor returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasHypervisor() bool {
	if o != nil && !IsNil(o.Hypervisor) {
		return true
	}

	return false
}

// SetHypervisor gets a reference to the given RefResponseModel and assigns it to the Hypervisor field.
func (o *DetectOrphanedResourcesResponseModel) SetHypervisor(v RefResponseModel) {
	o.Hypervisor = &v
}

// GetDetectScope returns the DetectScope field value if set, zero value otherwise.
func (o *DetectOrphanedResourcesResponseModel) GetDetectScope() DetectScope {
	if o == nil || IsNil(o.DetectScope) {
		var ret DetectScope
		return ret
	}
	return *o.DetectScope
}

// GetDetectScopeOk returns a tuple with the DetectScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectOrphanedResourcesResponseModel) GetDetectScopeOk() (*DetectScope, bool) {
	if o == nil || IsNil(o.DetectScope) {
		return nil, false
	}
	return o.DetectScope, true
}

// HasDetectScope returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasDetectScope() bool {
	if o != nil && !IsNil(o.DetectScope) {
		return true
	}

	return false
}

// SetDetectScope gets a reference to the given DetectScope and assigns it to the DetectScope field.
func (o *DetectOrphanedResourcesResponseModel) SetDetectScope(v DetectScope) {
	o.DetectScope = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetectOrphanedResourcesResponseModel) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetectOrphanedResourcesResponseModel) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *DetectOrphanedResourcesResponseModel) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *DetectOrphanedResourcesResponseModel) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *DetectOrphanedResourcesResponseModel) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetReportTime returns the ReportTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetectOrphanedResourcesResponseModel) GetReportTime() string {
	if o == nil || IsNil(o.ReportTime.Get()) {
		var ret string
		return ret
	}
	return *o.ReportTime.Get()
}

// GetReportTimeOk returns a tuple with the ReportTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetectOrphanedResourcesResponseModel) GetReportTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportTime.Get(), o.ReportTime.IsSet()
}

// HasReportTime returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasReportTime() bool {
	if o != nil && o.ReportTime.IsSet() {
		return true
	}

	return false
}

// SetReportTime gets a reference to the given NullableString and assigns it to the ReportTime field.
func (o *DetectOrphanedResourcesResponseModel) SetReportTime(v string) {
	o.ReportTime.Set(&v)
}

// SetReportTimeNil sets the value for ReportTime to be an explicit nil
func (o *DetectOrphanedResourcesResponseModel) SetReportTimeNil() {
	o.ReportTime.Set(nil)
}

// UnsetReportTime ensures that no value is present for ReportTime, not even an explicit nil
func (o *DetectOrphanedResourcesResponseModel) UnsetReportTime() {
	o.ReportTime.Unset()
}

// GetDetectStart returns the DetectStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetectOrphanedResourcesResponseModel) GetDetectStart() string {
	if o == nil || IsNil(o.DetectStart.Get()) {
		var ret string
		return ret
	}
	return *o.DetectStart.Get()
}

// GetDetectStartOk returns a tuple with the DetectStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetectOrphanedResourcesResponseModel) GetDetectStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetectStart.Get(), o.DetectStart.IsSet()
}

// HasDetectStart returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasDetectStart() bool {
	if o != nil && o.DetectStart.IsSet() {
		return true
	}

	return false
}

// SetDetectStart gets a reference to the given NullableString and assigns it to the DetectStart field.
func (o *DetectOrphanedResourcesResponseModel) SetDetectStart(v string) {
	o.DetectStart.Set(&v)
}

// SetDetectStartNil sets the value for DetectStart to be an explicit nil
func (o *DetectOrphanedResourcesResponseModel) SetDetectStartNil() {
	o.DetectStart.Set(nil)
}

// UnsetDetectStart ensures that no value is present for DetectStart, not even an explicit nil
func (o *DetectOrphanedResourcesResponseModel) UnsetDetectStart() {
	o.DetectStart.Unset()
}

// GetDetectEnd returns the DetectEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetectOrphanedResourcesResponseModel) GetDetectEnd() string {
	if o == nil || IsNil(o.DetectEnd.Get()) {
		var ret string
		return ret
	}
	return *o.DetectEnd.Get()
}

// GetDetectEndOk returns a tuple with the DetectEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetectOrphanedResourcesResponseModel) GetDetectEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetectEnd.Get(), o.DetectEnd.IsSet()
}

// HasDetectEnd returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasDetectEnd() bool {
	if o != nil && o.DetectEnd.IsSet() {
		return true
	}

	return false
}

// SetDetectEnd gets a reference to the given NullableString and assigns it to the DetectEnd field.
func (o *DetectOrphanedResourcesResponseModel) SetDetectEnd(v string) {
	o.DetectEnd.Set(&v)
}

// SetDetectEndNil sets the value for DetectEnd to be an explicit nil
func (o *DetectOrphanedResourcesResponseModel) SetDetectEndNil() {
	o.DetectEnd.Set(nil)
}

// UnsetDetectEnd ensures that no value is present for DetectEnd, not even an explicit nil
func (o *DetectOrphanedResourcesResponseModel) UnsetDetectEnd() {
	o.DetectEnd.Unset()
}

// GetNumDetected returns the NumDetected field value if set, zero value otherwise.
func (o *DetectOrphanedResourcesResponseModel) GetNumDetected() int32 {
	if o == nil || IsNil(o.NumDetected) {
		var ret int32
		return ret
	}
	return *o.NumDetected
}

// GetNumDetectedOk returns a tuple with the NumDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectOrphanedResourcesResponseModel) GetNumDetectedOk() (*int32, bool) {
	if o == nil || IsNil(o.NumDetected) {
		return nil, false
	}
	return o.NumDetected, true
}

// HasNumDetected returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasNumDetected() bool {
	if o != nil && !IsNil(o.NumDetected) {
		return true
	}

	return false
}

// SetNumDetected gets a reference to the given int32 and assigns it to the NumDetected field.
func (o *DetectOrphanedResourcesResponseModel) SetNumDetected(v int32) {
	o.NumDetected = &v
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetectOrphanedResourcesResponseModel) GetErrorMsg() string {
	if o == nil || IsNil(o.ErrorMsg.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMsg.Get()
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetectOrphanedResourcesResponseModel) GetErrorMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMsg.Get(), o.ErrorMsg.IsSet()
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasErrorMsg() bool {
	if o != nil && o.ErrorMsg.IsSet() {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given NullableString and assigns it to the ErrorMsg field.
func (o *DetectOrphanedResourcesResponseModel) SetErrorMsg(v string) {
	o.ErrorMsg.Set(&v)
}

// SetErrorMsgNil sets the value for ErrorMsg to be an explicit nil
func (o *DetectOrphanedResourcesResponseModel) SetErrorMsgNil() {
	o.ErrorMsg.Set(nil)
}

// UnsetErrorMsg ensures that no value is present for ErrorMsg, not even an explicit nil
func (o *DetectOrphanedResourcesResponseModel) UnsetErrorMsg() {
	o.ErrorMsg.Unset()
}

// GetIsFailed returns the IsFailed field value if set, zero value otherwise.
func (o *DetectOrphanedResourcesResponseModel) GetIsFailed() bool {
	if o == nil || IsNil(o.IsFailed) {
		var ret bool
		return ret
	}
	return *o.IsFailed
}

// GetIsFailedOk returns a tuple with the IsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectOrphanedResourcesResponseModel) GetIsFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFailed) {
		return nil, false
	}
	return o.IsFailed, true
}

// HasIsFailed returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasIsFailed() bool {
	if o != nil && !IsNil(o.IsFailed) {
		return true
	}

	return false
}

// SetIsFailed gets a reference to the given bool and assigns it to the IsFailed field.
func (o *DetectOrphanedResourcesResponseModel) SetIsFailed(v bool) {
	o.IsFailed = &v
}

// GetDetectResults returns the DetectResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetectOrphanedResourcesResponseModel) GetDetectResults() []DetectOrphanedResourcesResultModel {
	if o == nil {
		var ret []DetectOrphanedResourcesResultModel
		return ret
	}
	return o.DetectResults
}

// GetDetectResultsOk returns a tuple with the DetectResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetectOrphanedResourcesResponseModel) GetDetectResultsOk() ([]DetectOrphanedResourcesResultModel, bool) {
	if o == nil || IsNil(o.DetectResults) {
		return nil, false
	}
	return o.DetectResults, true
}

// HasDetectResults returns a boolean if a field has been set.
func (o *DetectOrphanedResourcesResponseModel) HasDetectResults() bool {
	if o != nil && IsNil(o.DetectResults) {
		return true
	}

	return false
}

// SetDetectResults gets a reference to the given []DetectOrphanedResourcesResultModel and assigns it to the DetectResults field.
func (o *DetectOrphanedResourcesResponseModel) SetDetectResults(v []DetectOrphanedResourcesResultModel) {
	o.DetectResults = v
}

func (o DetectOrphanedResourcesResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectOrphanedResourcesResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hypervisor) {
		toSerialize["Hypervisor"] = o.Hypervisor
	}
	if !IsNil(o.DetectScope) {
		toSerialize["DetectScope"] = o.DetectScope
	}
	if o.CreatedBy.IsSet() {
		toSerialize["CreatedBy"] = o.CreatedBy.Get()
	}
	if o.ReportTime.IsSet() {
		toSerialize["ReportTime"] = o.ReportTime.Get()
	}
	if o.DetectStart.IsSet() {
		toSerialize["DetectStart"] = o.DetectStart.Get()
	}
	if o.DetectEnd.IsSet() {
		toSerialize["DetectEnd"] = o.DetectEnd.Get()
	}
	if !IsNil(o.NumDetected) {
		toSerialize["NumDetected"] = o.NumDetected
	}
	if o.ErrorMsg.IsSet() {
		toSerialize["ErrorMsg"] = o.ErrorMsg.Get()
	}
	if !IsNil(o.IsFailed) {
		toSerialize["IsFailed"] = o.IsFailed
	}
	if o.DetectResults != nil {
		toSerialize["DetectResults"] = o.DetectResults
	}
	return toSerialize, nil
}

type NullableDetectOrphanedResourcesResponseModel struct {
	value *DetectOrphanedResourcesResponseModel
	isSet bool
}

func (v NullableDetectOrphanedResourcesResponseModel) Get() *DetectOrphanedResourcesResponseModel {
	return v.value
}

func (v *NullableDetectOrphanedResourcesResponseModel) Set(val *DetectOrphanedResourcesResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectOrphanedResourcesResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectOrphanedResourcesResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectOrphanedResourcesResponseModel(val *DetectOrphanedResourcesResponseModel) *NullableDetectOrphanedResourcesResponseModel {
	return &NullableDetectOrphanedResourcesResponseModel{value: val, isSet: true}
}

func (v NullableDetectOrphanedResourcesResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectOrphanedResourcesResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
