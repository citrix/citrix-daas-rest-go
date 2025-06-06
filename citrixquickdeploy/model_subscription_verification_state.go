/*
Citrix Virtual App & Desktop Catalog Service 147.0.26651.57932

Catalog Service

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickdeploy

import (
	"encoding/json"
)

// checks if the SubscriptionVerificationState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionVerificationState{}

// SubscriptionVerificationState struct for SubscriptionVerificationState
type SubscriptionVerificationState struct {
	// ID of the subscription
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// List of assignments in the subscription that are not expected
	UnexpectedAssignments    []AzureAssignment `json:"unexpectedAssignments,omitempty"`
	HasUnexpectedAssignments *bool             `json:"hasUnexpectedAssignments,omitempty"`
}

// NewSubscriptionVerificationStateWithDefaults instantiates a new SubscriptionVerificationState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionVerificationStateWithDefaults() *SubscriptionVerificationState {
	this := SubscriptionVerificationState{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SubscriptionVerificationState) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVerificationState) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SubscriptionVerificationState) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUnexpectedAssignments returns the UnexpectedAssignments field value if set, zero value otherwise.
func (o *SubscriptionVerificationState) GetUnexpectedAssignments() []AzureAssignment {
	if o == nil || IsNil(o.UnexpectedAssignments) {
		var ret []AzureAssignment
		return ret
	}
	return o.UnexpectedAssignments
}

// GetUnexpectedAssignmentsOk returns a tuple with the UnexpectedAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVerificationState) GetUnexpectedAssignmentsOk() ([]AzureAssignment, bool) {
	if o == nil || IsNil(o.UnexpectedAssignments) {
		return nil, false
	}
	return o.UnexpectedAssignments, true
}

// SetUnexpectedAssignments gets a reference to the given []AzureAssignment and assigns it to the UnexpectedAssignments field.
func (o *SubscriptionVerificationState) SetUnexpectedAssignments(v []AzureAssignment) {
	o.UnexpectedAssignments = v
}

// GetHasUnexpectedAssignments returns the HasUnexpectedAssignments field value if set, zero value otherwise.
func (o *SubscriptionVerificationState) GetHasUnexpectedAssignments() bool {
	if o == nil || IsNil(o.HasUnexpectedAssignments) {
		var ret bool
		return ret
	}
	return *o.HasUnexpectedAssignments
}

// GetHasUnexpectedAssignmentsOk returns a tuple with the HasUnexpectedAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVerificationState) GetHasUnexpectedAssignmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUnexpectedAssignments) {
		return nil, false
	}
	return o.HasUnexpectedAssignments, true
}

// SetHasUnexpectedAssignments gets a reference to the given bool and assigns it to the HasUnexpectedAssignments field.
func (o *SubscriptionVerificationState) SetHasUnexpectedAssignments(v bool) {
	o.HasUnexpectedAssignments = &v
}

func (o SubscriptionVerificationState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionVerificationState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.UnexpectedAssignments) {
		toSerialize["unexpectedAssignments"] = o.UnexpectedAssignments
	}
	if !IsNil(o.HasUnexpectedAssignments) {
		toSerialize["hasUnexpectedAssignments"] = o.HasUnexpectedAssignments
	}
	return toSerialize, nil
}

type NullableSubscriptionVerificationState struct {
	value *SubscriptionVerificationState
	isSet bool
}

func (v NullableSubscriptionVerificationState) Get() *SubscriptionVerificationState {
	return v.value
}

func (v *NullableSubscriptionVerificationState) Set(val *SubscriptionVerificationState) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionVerificationState) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionVerificationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionVerificationState(val *SubscriptionVerificationState) *NullableSubscriptionVerificationState {
	return &NullableSubscriptionVerificationState{value: val, isSet: true}
}

func (v NullableSubscriptionVerificationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionVerificationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
