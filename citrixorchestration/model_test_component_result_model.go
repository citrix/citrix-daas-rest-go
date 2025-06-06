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

// checks if the TestComponentResultModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestComponentResultModel{}

// TestComponentResultModel A test component result is the result of running a test against a particular controller.  If a test has EnvTestScope.Site, the test result will have only one component.  If a test has EnvTestScope.Controller, the test will have as many component results as there are controllers in the site.
type TestComponentResultModel struct {
	// Gets or sets the test component target, usually the machine the test was run on.
	TestComponentTarget NullableString `json:"TestComponentTarget,omitempty"`
	// Gets or sets the test component start time. If the test has not yet been run, this may be null.
	TestComponentStartTime NullableString `json:"TestComponentStartTime,omitempty"`
	// Gets or sets the formatted test component start time. If the test has not yet been run, this may be null. RFC 3339 compatible format.
	FormattedTestComponentStartTime NullableString `json:"FormattedTestComponentStartTime,omitempty"`
	// Gets or sets the test component end time. If the test has not yet been run, this may be null.
	TestComponentEndTime NullableString `json:"TestComponentEndTime,omitempty"`
	// Gets or sets the formatted test component end time. If the test has not yet been run, this may be null. RFC 3339 compatible format.
	FormattedTestComponentEndTime NullableString `json:"FormattedTestComponentEndTime,omitempty"`
	TestComponentStatus           *TestStatus    `json:"TestComponentStatus,omitempty"`
	// Gets or sets the result details, which is a list of additional information about the state of a machine when a test succeeds, but has warnings, or fails.  There may be more than one value for each severity, for example, multiple warnings about the state of the License server. The values must be localized.
	ResultDetails []TestComponentResultDetailModel `json:"ResultDetails,omitempty"`
}

// NewTestComponentResultModel instantiates a new TestComponentResultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestComponentResultModel() *TestComponentResultModel {
	this := TestComponentResultModel{}
	return &this
}

// NewTestComponentResultModelWithDefaults instantiates a new TestComponentResultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestComponentResultModelWithDefaults() *TestComponentResultModel {
	this := TestComponentResultModel{}
	return &this
}

// GetTestComponentTarget returns the TestComponentTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestComponentResultModel) GetTestComponentTarget() string {
	if o == nil || IsNil(o.TestComponentTarget.Get()) {
		var ret string
		return ret
	}
	return *o.TestComponentTarget.Get()
}

// GetTestComponentTargetOk returns a tuple with the TestComponentTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestComponentResultModel) GetTestComponentTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestComponentTarget.Get(), o.TestComponentTarget.IsSet()
}

// HasTestComponentTarget returns a boolean if a field has been set.
func (o *TestComponentResultModel) HasTestComponentTarget() bool {
	if o != nil && o.TestComponentTarget.IsSet() {
		return true
	}

	return false
}

// SetTestComponentTarget gets a reference to the given NullableString and assigns it to the TestComponentTarget field.
func (o *TestComponentResultModel) SetTestComponentTarget(v string) {
	o.TestComponentTarget.Set(&v)
}

// SetTestComponentTargetNil sets the value for TestComponentTarget to be an explicit nil
func (o *TestComponentResultModel) SetTestComponentTargetNil() {
	o.TestComponentTarget.Set(nil)
}

// UnsetTestComponentTarget ensures that no value is present for TestComponentTarget, not even an explicit nil
func (o *TestComponentResultModel) UnsetTestComponentTarget() {
	o.TestComponentTarget.Unset()
}

// GetTestComponentStartTime returns the TestComponentStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestComponentResultModel) GetTestComponentStartTime() string {
	if o == nil || IsNil(o.TestComponentStartTime.Get()) {
		var ret string
		return ret
	}
	return *o.TestComponentStartTime.Get()
}

// GetTestComponentStartTimeOk returns a tuple with the TestComponentStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestComponentResultModel) GetTestComponentStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestComponentStartTime.Get(), o.TestComponentStartTime.IsSet()
}

// HasTestComponentStartTime returns a boolean if a field has been set.
func (o *TestComponentResultModel) HasTestComponentStartTime() bool {
	if o != nil && o.TestComponentStartTime.IsSet() {
		return true
	}

	return false
}

// SetTestComponentStartTime gets a reference to the given NullableString and assigns it to the TestComponentStartTime field.
func (o *TestComponentResultModel) SetTestComponentStartTime(v string) {
	o.TestComponentStartTime.Set(&v)
}

// SetTestComponentStartTimeNil sets the value for TestComponentStartTime to be an explicit nil
func (o *TestComponentResultModel) SetTestComponentStartTimeNil() {
	o.TestComponentStartTime.Set(nil)
}

// UnsetTestComponentStartTime ensures that no value is present for TestComponentStartTime, not even an explicit nil
func (o *TestComponentResultModel) UnsetTestComponentStartTime() {
	o.TestComponentStartTime.Unset()
}

// GetFormattedTestComponentStartTime returns the FormattedTestComponentStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestComponentResultModel) GetFormattedTestComponentStartTime() string {
	if o == nil || IsNil(o.FormattedTestComponentStartTime.Get()) {
		var ret string
		return ret
	}
	return *o.FormattedTestComponentStartTime.Get()
}

// GetFormattedTestComponentStartTimeOk returns a tuple with the FormattedTestComponentStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestComponentResultModel) GetFormattedTestComponentStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormattedTestComponentStartTime.Get(), o.FormattedTestComponentStartTime.IsSet()
}

// HasFormattedTestComponentStartTime returns a boolean if a field has been set.
func (o *TestComponentResultModel) HasFormattedTestComponentStartTime() bool {
	if o != nil && o.FormattedTestComponentStartTime.IsSet() {
		return true
	}

	return false
}

// SetFormattedTestComponentStartTime gets a reference to the given NullableString and assigns it to the FormattedTestComponentStartTime field.
func (o *TestComponentResultModel) SetFormattedTestComponentStartTime(v string) {
	o.FormattedTestComponentStartTime.Set(&v)
}

// SetFormattedTestComponentStartTimeNil sets the value for FormattedTestComponentStartTime to be an explicit nil
func (o *TestComponentResultModel) SetFormattedTestComponentStartTimeNil() {
	o.FormattedTestComponentStartTime.Set(nil)
}

// UnsetFormattedTestComponentStartTime ensures that no value is present for FormattedTestComponentStartTime, not even an explicit nil
func (o *TestComponentResultModel) UnsetFormattedTestComponentStartTime() {
	o.FormattedTestComponentStartTime.Unset()
}

// GetTestComponentEndTime returns the TestComponentEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestComponentResultModel) GetTestComponentEndTime() string {
	if o == nil || IsNil(o.TestComponentEndTime.Get()) {
		var ret string
		return ret
	}
	return *o.TestComponentEndTime.Get()
}

// GetTestComponentEndTimeOk returns a tuple with the TestComponentEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestComponentResultModel) GetTestComponentEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestComponentEndTime.Get(), o.TestComponentEndTime.IsSet()
}

// HasTestComponentEndTime returns a boolean if a field has been set.
func (o *TestComponentResultModel) HasTestComponentEndTime() bool {
	if o != nil && o.TestComponentEndTime.IsSet() {
		return true
	}

	return false
}

// SetTestComponentEndTime gets a reference to the given NullableString and assigns it to the TestComponentEndTime field.
func (o *TestComponentResultModel) SetTestComponentEndTime(v string) {
	o.TestComponentEndTime.Set(&v)
}

// SetTestComponentEndTimeNil sets the value for TestComponentEndTime to be an explicit nil
func (o *TestComponentResultModel) SetTestComponentEndTimeNil() {
	o.TestComponentEndTime.Set(nil)
}

// UnsetTestComponentEndTime ensures that no value is present for TestComponentEndTime, not even an explicit nil
func (o *TestComponentResultModel) UnsetTestComponentEndTime() {
	o.TestComponentEndTime.Unset()
}

// GetFormattedTestComponentEndTime returns the FormattedTestComponentEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestComponentResultModel) GetFormattedTestComponentEndTime() string {
	if o == nil || IsNil(o.FormattedTestComponentEndTime.Get()) {
		var ret string
		return ret
	}
	return *o.FormattedTestComponentEndTime.Get()
}

// GetFormattedTestComponentEndTimeOk returns a tuple with the FormattedTestComponentEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestComponentResultModel) GetFormattedTestComponentEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormattedTestComponentEndTime.Get(), o.FormattedTestComponentEndTime.IsSet()
}

// HasFormattedTestComponentEndTime returns a boolean if a field has been set.
func (o *TestComponentResultModel) HasFormattedTestComponentEndTime() bool {
	if o != nil && o.FormattedTestComponentEndTime.IsSet() {
		return true
	}

	return false
}

// SetFormattedTestComponentEndTime gets a reference to the given NullableString and assigns it to the FormattedTestComponentEndTime field.
func (o *TestComponentResultModel) SetFormattedTestComponentEndTime(v string) {
	o.FormattedTestComponentEndTime.Set(&v)
}

// SetFormattedTestComponentEndTimeNil sets the value for FormattedTestComponentEndTime to be an explicit nil
func (o *TestComponentResultModel) SetFormattedTestComponentEndTimeNil() {
	o.FormattedTestComponentEndTime.Set(nil)
}

// UnsetFormattedTestComponentEndTime ensures that no value is present for FormattedTestComponentEndTime, not even an explicit nil
func (o *TestComponentResultModel) UnsetFormattedTestComponentEndTime() {
	o.FormattedTestComponentEndTime.Unset()
}

// GetTestComponentStatus returns the TestComponentStatus field value if set, zero value otherwise.
func (o *TestComponentResultModel) GetTestComponentStatus() TestStatus {
	if o == nil || IsNil(o.TestComponentStatus) {
		var ret TestStatus
		return ret
	}
	return *o.TestComponentStatus
}

// GetTestComponentStatusOk returns a tuple with the TestComponentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestComponentResultModel) GetTestComponentStatusOk() (*TestStatus, bool) {
	if o == nil || IsNil(o.TestComponentStatus) {
		return nil, false
	}
	return o.TestComponentStatus, true
}

// HasTestComponentStatus returns a boolean if a field has been set.
func (o *TestComponentResultModel) HasTestComponentStatus() bool {
	if o != nil && !IsNil(o.TestComponentStatus) {
		return true
	}

	return false
}

// SetTestComponentStatus gets a reference to the given TestStatus and assigns it to the TestComponentStatus field.
func (o *TestComponentResultModel) SetTestComponentStatus(v TestStatus) {
	o.TestComponentStatus = &v
}

// GetResultDetails returns the ResultDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestComponentResultModel) GetResultDetails() []TestComponentResultDetailModel {
	if o == nil {
		var ret []TestComponentResultDetailModel
		return ret
	}
	return o.ResultDetails
}

// GetResultDetailsOk returns a tuple with the ResultDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestComponentResultModel) GetResultDetailsOk() ([]TestComponentResultDetailModel, bool) {
	if o == nil || IsNil(o.ResultDetails) {
		return nil, false
	}
	return o.ResultDetails, true
}

// HasResultDetails returns a boolean if a field has been set.
func (o *TestComponentResultModel) HasResultDetails() bool {
	if o != nil && IsNil(o.ResultDetails) {
		return true
	}

	return false
}

// SetResultDetails gets a reference to the given []TestComponentResultDetailModel and assigns it to the ResultDetails field.
func (o *TestComponentResultModel) SetResultDetails(v []TestComponentResultDetailModel) {
	o.ResultDetails = v
}

func (o TestComponentResultModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestComponentResultModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TestComponentTarget.IsSet() {
		toSerialize["TestComponentTarget"] = o.TestComponentTarget.Get()
	}
	if o.TestComponentStartTime.IsSet() {
		toSerialize["TestComponentStartTime"] = o.TestComponentStartTime.Get()
	}
	if o.FormattedTestComponentStartTime.IsSet() {
		toSerialize["FormattedTestComponentStartTime"] = o.FormattedTestComponentStartTime.Get()
	}
	if o.TestComponentEndTime.IsSet() {
		toSerialize["TestComponentEndTime"] = o.TestComponentEndTime.Get()
	}
	if o.FormattedTestComponentEndTime.IsSet() {
		toSerialize["FormattedTestComponentEndTime"] = o.FormattedTestComponentEndTime.Get()
	}
	if !IsNil(o.TestComponentStatus) {
		toSerialize["TestComponentStatus"] = o.TestComponentStatus
	}
	if o.ResultDetails != nil {
		toSerialize["ResultDetails"] = o.ResultDetails
	}
	return toSerialize, nil
}

type NullableTestComponentResultModel struct {
	value *TestComponentResultModel
	isSet bool
}

func (v NullableTestComponentResultModel) Get() *TestComponentResultModel {
	return v.value
}

func (v *NullableTestComponentResultModel) Set(val *TestComponentResultModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestComponentResultModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestComponentResultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestComponentResultModel(val *TestComponentResultModel) *NullableTestComponentResultModel {
	return &NullableTestComponentResultModel{value: val, isSet: true}
}

func (v NullableTestComponentResultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestComponentResultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
